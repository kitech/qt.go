package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/kitech/qt.go/toolutil"
)

var file string
var filep string // file name without externsion

var cp = toolutil.NewCodePager()

const (
	INSEC_NONE = iota
	INSEC_RES_DATA
	INSEC_RES_NAME
	INSEC_RES_STRUCT
	INSEC_DONE
)

// usage: go-rcc <rcc.qrc>
// depend on: /usr/bin/rcc
func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	file = os.Args[1]
	filep = path.Base(file)[0:strings.LastIndex(path.Base(file), ".")]
	filep = IfElseStr(path.Dir(file) == "", filep, path.Dir(file)+"/"+filep)
	log.Println(file, filep)

	scc, err := toolutil.RunCmdOut("rcc", file)
	ErrPrint(err, "rcc", file)
	if err != nil {
		os.Exit(1)
	}

	lines := strings.Split(scc, "\n")
	log.Println("lines:", len(lines), "size:", len(scc))

	cp.APf("header", "import \"unsafe\"")
	cp.APf("header", "// import \"github.com/kitech/qt.go/qtcore\"")
	cp.APf("header", "import \"github.com/kitech/qt.go/qtmock\"")

	hasRccData, hasRccName, hasRccStruct := false, false, false
	insection := INSEC_NONE
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.Contains(line, "qt_resource_data[]") {
			insection = INSEC_RES_DATA
			hasRccData = true
			cp.APf("body", "var qt_resource_data = []byte{")
			continue
		}

		if strings.Contains(line, "qt_resource_name[]") {
			insection = INSEC_RES_NAME
			hasRccName = true
			cp.APf("body", "var qt_resource_name = []byte{")
			continue
		}

		if strings.Contains(line, "qt_resource_struct[]") {
			insection = INSEC_RES_STRUCT
			hasRccStruct = true
			cp.APf("body", "var qt_resource_struct = []byte{")
			continue
		}

		if strings.HasPrefix(line, "#ifdef QT_NAMESPACE") {
			insection = INSEC_DONE
			break
		}

		switch insection {
		case INSEC_RES_DATA:
			// do smth...
			// log.Println(line)
			transformMember(line)
		case INSEC_RES_NAME:
			onSetupUi(line)
		case INSEC_RES_STRUCT:
			onRetranslateUi(line)
		}
	}

	// fix not exists
	if !hasRccData {
		cp.APf("body", "var qt_resource_data = []byte{} // empty")
	}
	if !hasRccName {
		cp.APf("body", "var qt_resource_name = []byte{} // empty")
	}
	if !hasRccStruct {
		cp.APf("body", "var qt_resource_struct = []byte{} // empty")
	}

	saveCode()
}

func transformMember(line string) {

	home := os.Getenv("HOME")
	if home != "" {
		line = strings.Replace(line, home, "/home/me", -1)
	}
	cp.APf("body", line)
}

func onSetupUi(line string) {
	cp.APf("body", line)
}

func onRetranslateUi(line string) {
	cp.APf("body", line)
}

func saveCode() {
	var code string
	var mod os.FileMode = 0644

	// bool qRegisterResourceData (int, const unsigned char *, const unsigned char *, const unsigned char *).
	// bool qUnregisterResourceData (int, const unsigned char *, const unsigned char *, const unsigned char *)
	cp.APf("body", "var qt_rcc_version = 0x2 // >= 5.8.0, else 0x1")
	cp.APf("body", "func qInitResources(){")
	cp.APf("body", "  if len(qt_resource_name) > 0 {")
	cp.APf("body", "  qtmock.QRegisterResourceData(qt_rcc_version, unsafe.Pointer(&qt_resource_struct[0]), unsafe.Pointer(&qt_resource_name[0]), unsafe.Pointer(&qt_resource_data[0]))")
	cp.APf("body", "  }")
	cp.APf("body", "}")

	cp.APf("body", "func qCleanupResources(){")
	cp.APf("body", "  if len(qt_resource_name) > 0 {")
	cp.APf("body", "  qtmock.QUnregisterResourceData(qt_rcc_version, unsafe.Pointer(&qt_resource_struct[0]), unsafe.Pointer(&qt_resource_name[0]), unsafe.Pointer(&qt_resource_data[0]))")
	cp.APf("body", "  }")
	cp.APf("body", "}")

	cp.APf("body", "func init(){qInitResources()}")

	log.Printf("saving... %s_rc.go...\n", filep)
	code = "package main\n"
	code += cp.ExportAll()
	savefile := fmt.Sprintf("%s_rc.go", filep)
	err := ioutil.WriteFile(savefile, []byte(code), mod)
	ErrPrint(err, savefile)

	// gofmt the code
	gofmtPath, err := exec.LookPath("gofmt")
	ErrPrint(err)
	cmd := exec.Command(gofmtPath, "-w", savefile)
	err = cmd.Run()
	ErrPrint(err, cmd)
}

func colon2uline(s string) string { return strings.Replace(s, ":", "_", -1) }
func untitle(s string) string     { return strings.ToLower(s[0:1]) + s[1:] }

///
func printq(v interface{}, args ...interface{}) string {
	msg := fmt.Sprintf("%+v", v)
	for _, arg := range args {
		msg += fmt.Sprintf(" %+v", arg)
	}
	return msg
}

func ErrPrint(err error, args ...interface{}) error {
	if err != nil {
		log.Output(2, printq(err, args...))
	}
	return err
}

// TODO 要是侯选可以惰性求值就好了，否则在只能一个求值的场景则会有问题
// 简单的三元去处模拟函数
func IfElse(q bool, tv interface{}, fv interface{}) interface{} {
	if q == true {
		return tv
	} else {
		return fv
	}
}

func IfElseInt(q bool, tv int, fv int) int {
	return IfElse(q, tv, fv).(int)
}

func IfElseStr(q bool, tv string, fv string) string {
	return IfElse(q, tv, fv).(string)
}
