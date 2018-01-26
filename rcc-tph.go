package main

import (
	"fmt"
	"gopp"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var file string
var filep string // file name without externsion

var cp = NewCodePager()

const (
	INSEC_NONE = iota
	INSEC_RES_DATA
	INSEC_RES_NAME
	INSEC_RES_STRUCT
	INSEC_DONE
)

// usage: go-rcc <rcc.cpp>
func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	file = os.Args[1]
	filep = file[0:strings.LastIndex(file, ".")]
	log.Println(file, filep)

	bcc, err := ioutil.ReadFile(file)
	gopp.ErrPrint(err, file)
	lines := strings.Split(string(bcc), "\n")
	log.Println("lines:", len(lines), "size:", len(bcc))

	cp.APf("header", "// import \"qtcore\"")
	cp.APf("header", "import \"qtmock\"")
	cp.APf("header", "import \"unsafe\"")

	insection := INSEC_NONE
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.Contains(line, "qt_resource_data[]") {
			insection = INSEC_RES_DATA
			cp.APf("body", "var qt_resource_data = []byte{")
			continue
		}

		if strings.Contains(line, "qt_resource_name[]") {
			insection = INSEC_RES_NAME
			cp.APf("body", "var qt_resource_name = []byte{")
			continue
		}

		if strings.Contains(line, "qt_resource_struct[]") {
			insection = INSEC_RES_STRUCT
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
	cp.APf("body", "var qt_rcc_version = 0x2 // >= 5.8.0, else 0x1")
	cp.APf("body", "func qInitResources(){")
	cp.APf("body", "  qtmock.QRegisterResourceData(qt_rcc_version, unsafe.Pointer(&qt_resource_struct[0]), unsafe.Pointer(&qt_resource_name[0]), unsafe.Pointer(&qt_resource_data[0]))")
	cp.APf("body", "}")

	cp.APf("body", "func qCleanupResources(){")
	cp.APf("body", "  qtmock.QUnregisterResourceData(qt_rcc_version, unsafe.Pointer(&qt_resource_struct[0]), unsafe.Pointer(&qt_resource_name[0]), unsafe.Pointer(&qt_resource_data[0]))")
	cp.APf("body", "}")

	cp.APf("body", "func init(){qInitResources()}")

	log.Printf("saving... %s_rc.go...\n", filep)
	code = "package main\n"
	code += cp.ExportAll()
	ioutil.WriteFile(fmt.Sprintf("%s_rc.go", filep), []byte(code), mod)
}

func colon2uline(s string) string { return strings.Replace(s, ":", "_", -1) }
func untitle(s string) string     { return strings.ToLower(s[0:1]) + s[1:] }
