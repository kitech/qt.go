package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/toolutil"
)

var file string
var filep string            // file name without externsion
var pkgname string = "main" // final code's package name

var cp = toolutil.NewCodePager()

const (
	INSEC_NONE = iota
	INSEC_RES_DATA
	INSEC_RES_NAME
	INSEC_RES_STRUCT
	INSEC_DONE
)

func init() {
	flag.StringVar(&pkgname, "pkg", pkgname, "generated code's package name")
}

// cgo version rcc compile. feature: decrease go compile time RAM usage(~72%, 2.8G=>0.8G). but depend on gcc.
// usage: go-rcc -pkg [name] <rcc.qrc>
// depend on: /usr/bin/rcc
func main() {
	flag.Parse()

	log.SetFlags(log.Flags() | log.Lshortfile)
	file = flag.Arg(0)
	log.Println(flag.Arg(0))
	filep = path.Base(file)[0:strings.LastIndex(path.Base(file), ".")]
	filep = qtrt.IfElseStr(path.Dir(file) == "", filep, path.Dir(file)+"/"+filep)
	log.Println(file, filep)

	scc, err := toolutil.RunCmdOut("rcc", file)
	qtrt.ErrPrint(err, "rcc", file)
	if err != nil {
		os.Exit(1)
	}

	lines := strings.Split(scc, "\n")
	log.Println("lines:", len(lines), "size:", len(scc))

	cp.APf("header", "")
	cp.APf("header", "/*")
	cp.APf("header", "#include <string.h>")

	hasRccData, hasRccName, hasRccStruct := false, false, false
	insection := INSEC_NONE
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.Contains(line, "qt_resource_data[]") {
			insection = INSEC_RES_DATA
			hasRccData = true
			// cp.APf("body", "var qt_resource_data = []byte{")
			cp.APf("body", line[7:])
			continue
		}

		if strings.Contains(line, "qt_resource_name[]") {
			insection = INSEC_RES_NAME
			hasRccName = true
			// cp.APf("body", "var qt_resource_name = []byte{")
			cp.APf("body", line[7:])
			continue
		}

		if strings.Contains(line, "qt_resource_struct[]") {
			insection = INSEC_RES_STRUCT
			hasRccStruct = true
			// cp.APf("body", "var qt_resource_struct = []byte{")
			cp.APf("body", line[7:])
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

	cp.APf("header2", "*/")
	cp.APf("header2", "import \"C\"")
	cp.APf("header2", "import \"unsafe\"")
	cp.APf("header2", "// import \"github.com/kitech/qt.go/qtcore\"")
	cp.APf("header2", "import \"github.com/kitech/qt.go/qtmock\"")

	saveCode()
}

func transformMember(line string) {
	home := os.Getenv("HOME")
	line = strings.Replace(line, home, "/home/me", -1)
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
	cp.APf("bodygo", "var qt_rcc_version = 0x2 // >= 5.8.0, else 0x1")
	cp.APf("bodygo", "func qInitResources(){")
	cp.APf("bodygo", "  if C.sizeof_qt_resource_name > 0 {")
	cp.APf("bodygo", "  qtmock.QRegisterResourceData(qt_rcc_version, unsafe.Pointer(&C.qt_resource_struct[0]), unsafe.Pointer(&C.qt_resource_name[0]), unsafe.Pointer(&C.qt_resource_data[0]))")
	cp.APf("bodygo", "  }")
	cp.APf("bodygo", "}")

	cp.APf("bodygo", "func qCleanupResources(){")
	cp.APf("bodygo", "  if C.sizeof_qt_resource_name > 0 {")
	cp.APf("bodygo", "  qtmock.QUnregisterResourceData(qt_rcc_version, unsafe.Pointer(&C.qt_resource_struct[0]), unsafe.Pointer(&C.qt_resource_name[0]), unsafe.Pointer(&C.qt_resource_data[0]))")
	cp.APf("bodygo", "  }")
	cp.APf("bodygo", "}")

	cp.APf("bodygo", "func init(){qInitResources()}")

	log.Printf("saving... %s_rc.go...\n", filep)
	code = fmt.Sprintf("package %s\n", pkgname)
	code += cp.ExportAll()
	savefile := fmt.Sprintf("%s_rc.go", filep)
	err := ioutil.WriteFile(savefile, []byte(code), mod)
	qtrt.ErrPrint(err, savefile)

	// gofmt the code
	gofmtPath, err := exec.LookPath("gofmt")
	qtrt.ErrPrint(err)
	cmd := exec.Command(gofmtPath, "-w", savefile)
	err = cmd.Run()
	qtrt.ErrPrint(err, cmd)
}

func colon2uline(s string) string { return strings.Replace(s, ":", "_", -1) }
func untitle(s string) string     { return strings.ToLower(s[0:1]) + s[1:] }
