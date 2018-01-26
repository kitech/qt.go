package main

import (
	"fmt"
	"gopp"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

var class string
var wclass string // setupUi中用到的widget class
var woname string // main widget name
var file string
var filep string // file name without externsion

var cp = NewCodePager()

const (
	INSEC_NONE = iota
	INSEC_CLASS
	INSEC_MEMBER
	INSEC_SETUPUI
	INSEC_RETRANSLATEUI
	INSEC_DONE
)

// usage: go-uic <ui_foo.h>
func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	file = os.Args[1]
	filep = file[0:strings.LastIndex(file, ".")]
	log.Println(file, filep)

	bcc, err := ioutil.ReadFile(file)
	gopp.ErrPrint(err, file)
	lines := strings.Split(string(bcc), "\n")
	log.Println("lines:", len(lines), "size:", len(bcc))

	cp.APf("header", "import \"qtcore\"")
	cp.APf("header", "import \"qtgui\"")
	cp.APf("header", "import \"qtwidgets\"")
	cp.APf("header", "import \"qtmock\"")

	insection := INSEC_NONE
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "class ") {
			class = strings.Split(line, " ")[1]
			insection = INSEC_CLASS
			onClass(class)
			continue
		}
		if strings.HasPrefix(line, "public:") {
			insection = INSEC_MEMBER
			continue
		}
		if strings.HasPrefix(line, "void setupUi(") {
			insection = INSEC_SETUPUI
			wclass = strings.Split(strings.Split(line, "(")[1], " *")[0]
			woname = strings.Split(strings.Split(line, "(")[1], " *")[1]
			woname = strings.TrimRight(woname, ")")
			cp.APf("struct", "  %s *qtwidgets.%s", woname, wclass)
			cp.APf("setupUi", "// "+line)
			cp.APf("setupUi", "func (this *%s) SetupUi(%s *qtwidgets.%s) {", class, woname, wclass)
			cp.APf("setupUi", "  this.%s = %s", woname, woname)
			continue
		}

		if strings.HasPrefix(line, "void retranslateUi(") {
			cp.APf("struct", "}")
			cp.APf("setupUi", "}")
			insection = INSEC_RETRANSLATEUI
			cp.APf("setupUi", "// "+line)
			cp.APf("retranslateUi", "func (this *%s) RetranslateUi(%s *qtwidgets.%s) {", class, woname, wclass)
			continue
		}

		if strings.HasPrefix(line, "namespace Ui {") {
			cp.APf("retranslateUi", "}")
			insection = INSEC_DONE
			break
		}

		switch insection {
		case INSEC_MEMBER:
			// do smth...
			// log.Println(line)
			transformMember(line, class)
		case INSEC_SETUPUI:
			onSetupUi(line)
		case INSEC_RETRANSLATEUI:
			onRetranslateUi(line)
		}
	}

	saveCode()
}

func transformMember(line string, class string) {
	if line == "" {
		return
	}

	parts := strings.Split(line, "*")
	mname := parts[1][:len(parts[1])-1]
	memcls := parts[0]
	log.Println("transform member...", mname, line)

	cp.APf("struct", "  %s *qtwidgets.%s", strings.Title(mname), memcls)

	// var reline string

	// reline = fmt.Sprintf("func %s_Get_%s(p unsafe.Pointer) unsafe.Pointer {return C.%s_Get_%s(p)}",
	//	class, mname, class, mname)
	// golines = append(golines, reline)
}

func onClass(class string) {
	log.Printf("generating %s struct...\n", class)
	cp.APf("struct", "func New%s() *%s{", class, class)
	cp.APf("struct", "  return &%s{}", class)
	cp.APf("struct", "}")
	cp.APf("struct", "type %s struct{", class)
}

func onSetupUi(line string) {
	if line == "" {
		cp.APf("setupUi", "")
		return
	}

	if strings.Contains(line, "objectName().isEmpty()") {
		cp.APf("setupUi", "if %s.ObjectName().IsEmpty() {", woname)
	} else if strings.Contains(line, fmt.Sprintf("%s->setObjectName(", woname)) {
		oname := strings.Split(line, "\"")[1]
		cp.APf("setupUi", "%s.SetObjectName(qtcore.NewQString_5(\"%s\"))", woname, oname)
		cp.APf("setupUi", "}")
	} else if strings.HasPrefix(line, fmt.Sprintf("%s->resize(", woname)) {
		reg := regexp.MustCompile("resize.([0-9]+), ([0-9]+).")
		mats := reg.FindAllStringSubmatch(line, -1)
		log.Println(mats)
		cp.APf("setupUi", "%s.Resize(%s, %s)", woname, mats[0][1], mats[0][2])
	} else if strings.HasPrefix(line, fmt.Sprintf("retranslateUi(%s);", woname)) {
		cp.APf("setupUi", "this.RetranslateUi(%s)", woname)
	} else {
		reg1 := regexp.MustCompile(`(.*) = new (Q.*)\(([0-9A-Za-z_]*)?\);`)
		reg2 := regexp.MustCompile(`(.*)->setObjectName\(QStringLiteral\((.+)\)\);`)
		reg3 := regexp.MustCompile(`(.*)->(set.+Size)\(QSize\((.+)\)\);`)
		reg4 := regexp.MustCompile(`([^->]+)[->.]+set([^(]+)\((.+)\);`)
		reg5 := regexp.MustCompile(`([^->]+)[->.]+(add[^(]+)\((.+)\);`)
		reg6 := regexp.MustCompile(`(Q[A-Z][iconfont]+) ([iconfont0-9]+);`)
		reg7 := regexp.MustCompile(`QSizePolicy (.+)\((.+), (.+)\);`)
		reg8 := regexp.MustCompile(`(.+) = new (QSpacerItem)\(([0-9]+), ([0-9]+), (.+), (.+)\);`)
		if reg1.MatchString(line) {
			mats := reg1.FindAllStringSubmatch(line, -1)
			refmtval := strings.Title(mats[0][3])
			refmtsuf := ""
			switch mats[0][2] {
			case "QAction":
				refmtval = fmt.Sprintf("qtcore.NewQObjectFromPointer(%s.GetCthis())", refmtval)
			case "QWidget":
				if refmtval != "" {
					refmtval = fmt.Sprintf("qtwidgets.NewQWidgetFromPointer(this.%s.GetCthis()), 0", refmtval)
				} else {
					refmtval = fmt.Sprintf("nil, 0")
				}
			case "QVBoxLayout", "QHBoxLayout":
				if refmtval != "" {
					refmtval = fmt.Sprintf("qtwidgets.NewQWidgetFromPointer(this.%s.GetCthis())", refmtval)
					refmtsuf = "_1"
				}
			case "QLabel":
				refmtval = "this." + refmtval + ", 0"
			default:
				refmtval = "this." + refmtval
			}
			cp.APf("setupUi", "this.%s = qtwidgets.New%s%s(%s) // 111",
				strings.Title(mats[0][1]), mats[0][2], refmtsuf, refmtval)
		} else if reg2.MatchString(line) {
			mats := reg2.FindAllStringSubmatch(line, -1)
			log.Println(mats)
			cp.APf("setupUi", "this.%s.SetObjectName(qtcore.NewQString_5(%s)) // 112",
				strings.Title(mats[0][1]), strings.Title(mats[0][2]))
		} else if reg3.MatchString(line) {
			mats := reg3.FindAllStringSubmatch(line, -1)
			log.Println(mats)
			refmtname := strings.Title(mats[0][2])
			refmtval := mats[0][3]
			refmtsuf := ""
			switch mats[0][2] {
			case "setMaximumSize", "setMinimumSize":
				refmtsuf = "_1"
			case "setIconSize":
				refmtval = fmt.Sprintf("qtcore.NewQSize_1(%s)", refmtval)
			}
			cp.APf("setupUi", "this.%s.%s%s(%s)  // 113",
				strings.Title(mats[0][1]), refmtname, refmtsuf, refmtval)
		} else if reg4.MatchString(line) {
			mats := reg4.FindAllStringSubmatch(line, -1)
			log.Println(mats)
			refmtname := strings.Title(mats[0][2])
			refmtval := strings.Title(mats[0][3])
			refmtsuf := ""
			switch mats[0][2] {
			case "Spacing", "HorizontalStretch", "VerticalStretch":
			case "PointSize", "Weight": // do nothing
			case "ContentsMargins":
			case "Orientation":
				refmtval = "qtcore." + strings.Replace(refmtval, ":", "_", -1)
			case "TextInteractionFlags":
				refmtval = "qtcore." + strings.Replace(refmtval, ":", "_", -1)
				refmtval = strings.Replace(refmtval, "|", "|qtcore.", -1)
			case "AutoRaise", "WidgetResizable", "AlternatingRowColors":
				refmtval = strings.ToLower(refmtval[0:1]) + refmtval[1:]
			case "ToolButtonStyle":
				refmtval = "qtcore." + strings.Replace(refmtval, ":", "_", -1)
			case "Geometry":
				refmtval = strings.TrimRight(refmtval[6:], ")")
			case "HorizontalScrollBarPolicy":
				refmtval = "qtcore." + strings.Replace(refmtval, ":", "_", -1)
			case "MaximumSize", "MinimumSize":
				refmtsuf = "_1"
			case "Widget":
				// refmtname = "this." + refmtname
				refmtval = "this." + strings.Title(refmtval)
			case "Bold":
				refmtval = untitle(refmtval) // True => true
			case "Pixmap":
				refmtval = fmt.Sprintf("qtgui.NewQPixmap_3(qtcore.NewQString_5(\"%s\"), \"dummy123\", 0)", strings.Split(refmtval, "\"")[1])
			case "HeightForWidth":
				refmtval = "this." + strings.Replace(refmtval, "->", ".", -1)
				refmtval = "false" // TODO label_x.SizePolicy().HasHeightForWidth() crash
			default:
				log.Println(line, refmtname)
				refmtval = "this." + refmtval
			}
			cp.APf("setupUi", "this.%s.Set%s%s(%s) // 114",
				strings.Title(mats[0][1]), refmtname, refmtsuf, refmtval)
		} else if reg5.MatchString(line) {
			mats := reg5.FindAllStringSubmatch(line, -1)
			log.Println(mats)
			refmtname := strings.Title(mats[0][2])
			refmtval := strings.Title(mats[0][3])
			switch mats[0][2] {
			case "addWidget":
				refmtval = fmt.Sprintf("qtwidgets.NewQWidgetFromPointer(this.%s.GetCthis())", refmtval)
			case "addLayout":
				refmtval = fmt.Sprintf("qtwidgets.NewQLayoutFromPointer(this.%s.GetCthis()), 0", refmtval)
			case "addItem":
				refmtval = fmt.Sprintf("qtwidgets.NewQLayoutItemFromPointer(this.%s.GetCthis())", refmtval)
			case "addFile":
				parts := strings.Split(refmtval, ", ")
				alst := []string{}
				log.Println(strings.Split(mats[0][3], "\""))
				icores := strings.Split(mats[0][3], "\"")[1]
				alst = append(alst, fmt.Sprintf("qtcore.NewQString_5(\"%s\")", icores))
				alst = append(alst, "qtcore.NewQSize()")
				alst = append(alst, "qtgui."+colon2uline(parts[2]))
				alst = append(alst, "qtgui."+colon2uline(parts[3]))
				refmtval = strings.Join(alst, ", ")
				refmtname = "AddFile"
			default:
				refmtval = "this." + refmtval
			}
			cp.APf("setupUi", "this.%s.%s(%s) // 115",
				strings.Title(mats[0][1]), refmtname, refmtval)
		} else if reg6.MatchString(line) {
			mats := reg6.FindAllStringSubmatch(line, -1)
			log.Println(mats, line)
			cp.APf("struct", "  %s *qtgui.%s // 116", strings.Title(mats[0][2]), mats[0][1])
			cp.APf("setupUi", "  this.%s = qtgui.New%s()", strings.Title(mats[0][2]), mats[0][1])
		} else if reg7.MatchString(line) {
			mats := reg7.FindAllStringSubmatch(line, -1)
			log.Println(mats, line)
			cp.APf("struct", "  %s *qtwidgets.%s", strings.Title(mats[0][1]), "QSizePolicy")
			refmtvals := []string{}
			for _, v := range mats[0][2:] {
				refmtvals = append(refmtvals, "qtwidgets."+strings.Replace(v, ":", "_", -1))
			}
			refmtval := strings.Join(refmtvals, ", ")
			cp.APf("setupUi", "  this.%s = qtwidgets.New%s_1(%s, 1)", strings.Title(mats[0][1]), "QSizePolicy", refmtval)
		} else if reg8.MatchString(line) {
			mats := reg8.FindAllStringSubmatch(line, -1)
			log.Println(mats, line)
			a0 := mats[0][3]
			a1 := mats[0][4]
			a2 := "qtwidgets." + strings.Replace(mats[0][5], ":", "_", -1)
			a3 := "qtwidgets." + strings.Replace(mats[0][6], ":", "_", -1)
			cp.APf("setupUi", " this.%s = qtwidgets.NewQSpacerItem(%s, %s, %s, %s)",
				strings.Title(mats[0][1]), a0, a1, a2, a3)
		} else {
			cp.APf("setupUi", fmt.Sprintf("// %s // 126", line))
		}
	}
}

func onRetranslateUi(line string) {
	log.Println(line)
	reg1 := regexp.MustCompile(`(.+)->(.+)\(QApplication::translate\(\"(.+)\", \"(.+)\", nullptr\)\);`)
	if reg1.MatchString(line) {
		mats := reg1.FindAllStringSubmatch(line, -1)
		log.Println(mats)
		cp.APf("retranslateUi",
			// TODO qtcore.QCoreApplication_Translate crash
			"this.%s.%s(qtmock.QCoreApplication_Translate(\"%s\", \"%s\", \"dummy123\", 0))",
			strings.Title(mats[0][1]), strings.Title(mats[0][2]), mats[0][3], mats[0][4])
	}
}

func saveCode() {
	var code string
	var mod os.FileMode = 0644

	log.Printf("saving... %s.go...\n", filep)
	code = "package main\n"
	code += cp.ExportAll()
	ioutil.WriteFile(fmt.Sprintf("%s.go", filep), []byte(code), mod)
}

func colon2uline(s string) string { return strings.Replace(s, ":", "_", -1) }
func untitle(s string) string     { return strings.ToLower(s[0:1]) + s[1:] }
