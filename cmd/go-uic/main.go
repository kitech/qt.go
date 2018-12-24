package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"

	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/toolutil"
)

var class string
var wclass string // setupUi中用到的widget class
var woname string // main widget name
var file string
var filep string // file name without externsion

var cp = toolutil.NewCodePager()

const (
	INSEC_NONE = iota
	INSEC_CLASS
	INSEC_MEMBER
	INSEC_SETUPUI
	INSEC_RETRANSLATEUI
	INSEC_DONE
)

// usage: go-uic <foo.ui>
// depend on: /usr/bin/uic
func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	file = os.Args[1]

	filep = "ui_" + path.Base(file)[0:strings.LastIndex(path.Base(file), ".")]
	filep = qtrt.IfElseStr(path.Dir(file) == "", filep, path.Dir(file)+"/"+filep)
	log.Println(file, filep)

	scc, err := toolutil.RunCmdOut("uic", "-g", "cpp", file)
	qtrt.ErrPrint(err)
	if err != nil {
		os.Exit(1)
	}

	lines := strings.Split(scc, "\n")
	log.Println("lines:", len(lines), "size:", len(scc))

	cp.APf("header", "import \"github.com/kitech/qt.go/qtcore\"")
	cp.APf("header", "import \"github.com/kitech/qt.go/qtgui\"")
	cp.APf("header", "import \"github.com/kitech/qt.go/qtwidgets\"")
	cp.APf("header", "import \"github.com/kitech/qt.go/qtquickwidgets\"")
	cp.APf("header", "import \"github.com/kitech/qt.go/qtmock\"")
	cp.APf("header", "import \"github.com/kitech/qt.go/qtrt\"")
	cp.APf("header", "func init(){qtcore.KeepMe()}")
	cp.APf("header", "func init(){qtgui.KeepMe()}")
	cp.APf("header", "func init(){qtwidgets.KeepMe()}")
	cp.APf("header", "func init(){qtquickwidgets.KeepMe()}")
	cp.APf("header", "func init(){qtmock.KeepMe()}")
	cp.APf("header", "func init(){qtrt.KeepMe()}")

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
			onDone()
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

// 记录每一个member的类型。 member name => type name
var memberTypes = map[string]string{}

func transformMember(line string, class string) {
	if line == "" {
		return
	}

	parts := strings.Split(line, "*")
	memname := parts[1][:len(parts[1])-1]
	memcls := strings.TrimSpace(parts[0])
	log.Println("transform member...", memcls, memname, line)
	memberTypes[strings.Title(memname)] = memcls

	switch memcls {
	case "QQuickWidget":
		cp.APf("struct", "  %s *qtquickwidgets.%s", strings.Title(memname), memcls)
	default:
		cp.APf("struct", "  %s *qtwidgets.%s", strings.Title(memname), memcls)
	}

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
		cp.APf("setupUi", "if %s.ObjectName() == \"\" {", woname)
	} else if strings.Contains(line, fmt.Sprintf("%s->setObjectName(", woname)) {
		oname := strings.Split(line, "\"")[1]
		cp.APf("setupUi", "%s.SetObjectName(\"%s\")", woname, oname)
		cp.APf("setupUi", "}")
	} else if strings.HasPrefix(line, fmt.Sprintf("%s->resize(", woname)) {
		reg := regexp.MustCompile("resize.([0-9]+), ([0-9]+).")
		mats := reg.FindAllStringSubmatch(line, -1)
		log.Println(mats)
		cp.APf("setupUi", "%s.Resize(%s, %s)", woname, mats[0][1], mats[0][2])
	} else if strings.HasPrefix(line, fmt.Sprintf("retranslateUi(%s);", woname)) {
		cp.APf("setupUi", "this.RetranslateUi(%s)", woname)
	} else {
		reg1 := regexp.MustCompile(`(.*) = new (Q.+)\(([0-9A-Za-z_]*)?\);`)
		reg2 := regexp.MustCompile(`(.*)->setObjectName\(QStringLiteral\((.+)\)\);`)
		reg2_1 := regexp.MustCompile(`(.*)->setObjectName\(QString::fromUtf8\((.+)\)\)`)
		reg3 := regexp.MustCompile(`(.*)->(set.+Size)\(QSize\((.+)\)\);`)
		reg4 := regexp.MustCompile(`([^->]+)[->.]+set([^(]+)\((.+)\);`)
		reg5 := regexp.MustCompile(`([^->]+)[->.]+(add[^(]+)\((.+)\);`)
		reg6 := regexp.MustCompile(`(Q[A-Z][iconfont]+) ([iconfont0-9]+);`)
		reg7 := regexp.MustCompile(`QSizePolicy (.+)\((.+), (.+)\);`)
		reg8 := regexp.MustCompile(`(.+) = new (QSpacerItem)\(([0-9]+), ([0-9]+), (.+), (.+)\);`)
		reg9 := regexp.MustCompile(`if \((.+)->([rowCountcolumn]+)\(\) < ([0-9]+)\)`) // if (tableWidget->rowCount() < 3)
		reg100 := regexp.MustCompile(`QMetaObject::connectSlotsByName\((.+)\);`)
		if reg100.MatchString(line) {
			mats := reg100.FindAllStringSubmatch(line, -1)
			cp.APf("setupUi", "qtcore.QMetaObject_ConnectSlotsByName(%s) // 100111", mats[0][1])
		} else if reg1.MatchString(line) {
			mats := reg1.FindAllStringSubmatch(line, -1)
			log.Println(mats)
			refmtname := strings.Title(mats[0][1])
			refmtval := strings.Title(mats[0][3])
			refmtsuf := ""
			pkgname := "qtwidgets"
			switch mats[0][2] {
			case "QAction":
				refmtval = fmt.Sprintf("this.%s", refmtval)
			case "QWidget":
				if refmtval != "" {
					refmtval = fmt.Sprintf("this.%s, 0", refmtval)
				} else {
					refmtval = fmt.Sprintf("nil, 0")
				}
			case "QVBoxLayout", "QHBoxLayout":
				if refmtval != "" {
					refmtval = fmt.Sprintf("this.%s", refmtval)
					refmtsuf = "1"
				}
			case "QGridLayout":
				if refmtval == "" {
					refmtval = "nil"
				} else {
					refmtval = "this." + refmtval
				}
			case "QLabel", "QFrame", "QToolBox":
				refmtval = "this." + refmtval + ", 0"
			case "QTableWidgetItem":
				parts := strings.Split(refmtname, "*")
				refmtname = parts[1]
				refmtval = "0"
				cp.APf("struct", "  %s *qtwidgets.%s // 215", parts[1], parts[0])
			case "QQuickWidget":
				pkgname = "qtquickwidgets"
				refmtval = "this." + refmtval
			default:
				refmtval = "this." + refmtval
			}
			cp.APf("setupUi", "this.%s = %s.New%s%s(%s) // 111",
				refmtname, pkgname, mats[0][2], refmtsuf, refmtval)
		} else if reg2.MatchString(line) {
			mats := reg2.FindAllStringSubmatch(line, -1)
			log.Println(mats)
			cp.APf("setupUi", "this.%s.SetObjectName(%s) // 112",
				strings.Title(mats[0][1]), strings.Title(mats[0][2]))
		} else if reg2_1.MatchString(line) {
			log.Println("reg2_1")
			mats := reg2_1.FindAllStringSubmatch(line, -1)
			log.Println(mats)
			log.Println(strings.Title(mats[0][2]))
			cp.APf("setupUi", "this.%s.SetObjectName(%s) // 112",
				strings.Title(mats[0][1]), strings.Title(mats[0][2]))
		} else if reg3.MatchString(line) {
			mats := reg3.FindAllStringSubmatch(line, -1)
			log.Println(mats)
			refmtname := strings.Title(mats[0][2])
			refmtval := mats[0][3]
			refmtsuf := ""
			switch mats[0][2] {
			case "setMaximumSize", "setMinimumSize":
				refmtsuf = "1"
			case "setIconSize":
				refmtval = fmt.Sprintf("qtcore.NewQSize1(%s)", refmtval)
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
				if mats[0][1] == "layout()" { // toolBox->layout()->setSpacing(60);
					mats[0][1] = line[:strings.Index(line, mats[0][1])-2] + "." + strings.Title(mats[0][1])
				}
			case "PointSize", "Weight", "ColumnCount", "RowCount": // do nothing
			case "ContentsMargins", "CurrentIndex", "LineWidth",
				"AutoScrollMargin", "MidLineWidth", "ToolTipDuration",
				"Margin", "Indent", "Value", "Minimum", "Maximum",
				"LineWrapColumnOrWidth", "TabStopWidth", "CursorWidth",
				"DigitCount": // num no change
			case "MaxVisibleItems", "ModelColumn", "VerticalSpacing":
			case "Orientation", "TextFormat", "TextElideMode", "LayoutDirection":
				refmtval = "qtcore." + strings.Replace(refmtval, ":", "_", -1)
			case "TextInteractionFlags", "InputMethodHints":
				refmtval = "qtcore." + strings.Replace(refmtval, ":", "_", -1)
				refmtval = strings.Replace(refmtval, "|", "|qtcore.", -1)
			case "AutoRaise", "WidgetResizable", "AlternatingRowColors",
				"AutoRepeat", "AutoExclusive", "DocumentMode",
				"Checked", "Flat", "AutoFillBackground":
				refmtval = strings.ToLower(refmtval[0:1]) + refmtval[1:]
			case "Bold", "OpenExternalLinks", "WordWrap", "Frame", "Editable",
				"DragDropOverwriteMode", "ReadOnly", "AcceptDrops",
				"AcceptRichText", "Checkable", "MouseTracking", "Enabled",
				"UsesScrollButtons", "AutoScroll", "CornerButtonEnabled",
				"DragEnabled", "SortingEnabled", "ShowGrid", "TabKeyNavigation",
				"TabletTracking", "InvertedAppearance", "TextVisible",
				"TabChangesFocus", "UndoRedoEnabled", "OverwriteMode",
				"OpenLinks", "SmallDecimalPoint":
				refmtval = untitle(refmtval) // True => true
			case "ToolButtonStyle", "GridStyle", "DefaultDropAction":
				refmtval = "qtcore." + strings.Replace(refmtval, ":", "_", -1)
			case "Alignment", "FocusPolicy", "ArrowType":
				refmtval = "qtcore." + strings.Replace(refmtval, ":", "_", -1)
				refmtval = strings.Replace(refmtval, "|", "|qtcore.", -1)
			case "EditTriggers", "StandardButtons", "PopupMode":
				refmtval = "qtwidgets." + strings.Replace(refmtval, ":", "_", -1)
				refmtval = strings.Replace(refmtval, "|", "|qtwidgets.", -1)
			case "Geometry":
				refmtval = strings.TrimRight(refmtval[6:], ")")
			case "HorizontalScrollBarPolicy", "VerticalScrollBarPolicy", "ContextMenuPolicy":
				refmtval = "qtcore." + strings.Replace(refmtval, ":", "_", -1)
			case "SizeAdjustPolicy", "SizeConstraint", "FrameShape", "FrameShadow",
				"SelectionBehavior", "SelectionMode", "DragDropMode",
				"VerticalScrollMode", "HorizontalScrollMode", "EchoMode",
				"TextDirection", "AutoFormatting", "LineWrapMode", "Mode",
				"SegmentStyle":
				refmtval = "qtwidgets." + strings.Replace(refmtval, ":", "_", -1)
			case "ResizeMode":
				refmtval = "qtquickwidgets." + strings.Replace(refmtval, ":", "_", -1)
			case "MaximumSize", "MinimumSize":
				refmtsuf = "1"
			case "Widget":
				// refmtname = "this." + refmtname
				refmtval = "this." + strings.Title(refmtval)
			case "Pixmap":
				refmtval = mats[0][3]
				refmtval = fmt.Sprintf("qtgui.NewQPixmap3(\"%s\", \"dummy123\", 0)", strings.Split(refmtval, "\"")[1])
			case "Source": // QQuickWidget
				refmtval = mats[0][3]
				refmtval = fmt.Sprintf("qtcore.NewQUrl1(\"%s\", 0)", strings.Split(refmtval, "\"")[1])
			case "StyleSheet": // QQuickWidget
				refmtval = mats[0][3]
				refmtval = fmt.Sprintf("\"%s\"", strings.Split(refmtval, "\"")[1])
			case "HeightForWidth":
				refmtval = "this." + strings.Replace(refmtval, "->", ".", -1)
			case "HorizontalHeaderItem", "VerticalHeaderItem":
				refmtval = strings.Replace(refmtval, ",", ", this.", 1)
				// refmtval = "false" // TODO label_x.SizePolicy().HasHeightForWidth() crash
			// case "SizePolicy": // TODO 可能值有点问题，过滤掉设置setSizePolicy
			// break
			case "Property":
				refmtval = mats[0][3]
				numreg := regexp.MustCompile(`QVariant\([e\+\d]+\)`)
				if strings.Contains(refmtval, "(false)") || strings.Contains(refmtval, "(true)") {
					refmtval = strings.Replace(refmtval, "QVariant", "qtcore.NewQVariant9", -1)
				} else if numreg.MatchString(refmtval) {
					refmtval = strings.Replace(refmtval, "QVariant", "qtcore.NewQVariant5", -1)
				}
			default:
				log.Println(line, refmtname)
				refmtval = "this." + refmtval
			}

			// cp.APf("setupUi", "// %s", line)
			cp.APf("setupUi", "this.%s.Set%s%s(%s) // 114",
				strings.Title(mats[0][1]), refmtname, refmtsuf, refmtval)

		} else if reg5.MatchString(line) {
			mats := reg5.FindAllStringSubmatch(line, -1)
			log.Println(mats)
			refmtname := strings.Title(mats[0][2])
			refmtval := strings.Title(mats[0][3])
			memty := memberTypes[strings.Title(mats[0][1])]
			switch mats[0][2] {
			case "addWidget":
				log.Println(line, memty, mats[0][1], refmtname)
				if memty == "QGridLayout" {
					refmtname += "2"
					refmtval = fmt.Sprintf("this.%s, 0", refmtval)
				} else {
					if strings.Contains(mats[0][1], "Layout") {
						refmtname = "Layout()." + refmtname
					}
					refmtval = fmt.Sprintf("this.%s", refmtval)
				}
			case "addLayout":
				refmtval = fmt.Sprintf("this.%s, 0", refmtval)
			case "addItem":
				// toolbox
				tbreg := regexp.MustCompile(`(.+)\, QStringLiteral\((.+)\)`)
				if strings.HasPrefix(refmtval, "QString") { // QComboBox
					refmtval = fmt.Sprintf("\"\", qtcore.NewQVariant12(\"wtf\")")
				} else if tbreg.MatchString(refmtval) {
					tbmats := tbreg.FindAllStringSubmatch(refmtval, -1)
					refmtval = fmt.Sprintf("this.%s, %s", tbmats[0][1], tbmats[0][2])
				} else {
					refmtval = fmt.Sprintf("this.%s", refmtval)
				}
			case "addFile":
				parts := strings.Split(refmtval, ", ")
				alst := []string{}
				log.Println(strings.Split(mats[0][3], "\""))
				icores := strings.Split(mats[0][3], "\"")[1]
				alst = append(alst, fmt.Sprintf("\"%s\"", icores))
				alst = append(alst, "qtcore.NewQSize()")
				alst = append(alst, "qtgui."+colon2uline(parts[2]))
				alst = append(alst, "qtgui."+colon2uline(parts[3]))
				refmtval = strings.Join(alst, ", ")
				refmtname = "AddFile"
			case "addTab":
				refmtval = fmt.Sprintf("this.%s, \"\"", strings.Split(refmtval, ",")[0])
			case "addAction":
				refmtval = fmt.Sprintf("this.%s", strings.Replace(refmtval, "->", ".", -1))
				refmtname = "QWidget." + refmtname
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
			cp.APf("setupUi", "  this.%s = qtwidgets.New%s1(%s, 1)", strings.Title(mats[0][1]), "QSizePolicy", refmtval)
		} else if reg8.MatchString(line) {
			mats := reg8.FindAllStringSubmatch(line, -1)
			log.Println(mats, line)
			a0 := mats[0][3]
			a1 := mats[0][4]
			a2 := "qtwidgets." + strings.Replace(mats[0][5], ":", "_", -1)
			a3 := "qtwidgets." + strings.Replace(mats[0][6], ":", "_", -1)
			cp.APf("setupUi", " this.%s = qtwidgets.NewQSpacerItem(%s, %s, %s, %s)",
				strings.Title(mats[0][1]), a0, a1, a2, a3)
			cp.APf("setupUi", " qtrt.ReleaseOwnerToQt(this.%s)", strings.Title(mats[0][1]))
		} else if reg9.MatchString(line) {
			mats := reg9.FindAllStringSubmatch(line, -1)
			log.Println(len(mats), mats, "//", line)
			cp.APf("setupUi", " // if this.%s.%s() > %s // 119",
				strings.Title(mats[0][1]), strings.Title(mats[0][2]), mats[0][3])
		} else {
			cp.APf("setupUi", fmt.Sprintf("// %s // 126", line))
		}
	}
}

func onRetranslateUi(line string) {
	log.Println(line)
	reg1 := regexp.MustCompile(`(.+)->(.+)\(QApplication::translate\(\"(.+)\", \"(.+)\", nullptr\)\);`)
	reg2 := regexp.MustCompile(`(.+)->(.+)\(([0-9]+), QApplication::translate\(\"(.+)\", \"(.+)\", nullptr\)\);`)
	reg3 := regexp.MustCompile(`(.+)->(.+)\((.+)->(.+)\((.+)\), QApplication::translate\(\"(.+)\", \"(.+)\", nullptr\)\);`)
	if reg1.MatchString(line) {
		mats := reg1.FindAllStringSubmatch(line, -1)
		log.Println(mats)
		if strings.HasPrefix(mats[0][1], "___") {
			mats[0][1] = mats[0][1][1:]
		}
		cp.APf("retranslateUi",
			"this.%s.%s(qtcore.QCoreApplication_Translate(\"%s\", \"%s\", \"dummy123\", 0))",
			strings.Title(mats[0][1]), strings.Title(mats[0][2]), mats[0][3], mats[0][4])
	} else if reg2.MatchString(line) {
		mats := reg2.FindAllStringSubmatch(line, -1)
		log.Println(mats)
		cp.APf("retranslateUi",
			"this.%s.%s(%s, qtcore.QCoreApplication_Translate(\"%s\", \"%s\", \"dummy123\", 0))",
			strings.Title(mats[0][1]), strings.Title(mats[0][2]), mats[0][3], mats[0][4], mats[0][5])
	} else if reg3.MatchString(line) {
		mats := reg3.FindAllStringSubmatch(line, -1)
		log.Println(mats)
		cp.APf("retranslateUi",
			"this.%s.%s(this.%s.%s(this.%s), qtcore.QCoreApplication_Translate(\"%s\", \"%s\", \"dummy123\", 0))",
			strings.Title(mats[0][1]), strings.Title(mats[0][2]),
			strings.Title(mats[0][3]), strings.Title(mats[0][4]),
			strings.Title(mats[0][5]), mats[0][6], mats[0][7])
	} else {
		log.Println("what not matched:", line)
		cp.APf("retranslateUi", "// noimpl: %s", line)
	}
	// TODO 多行tooltip的情况
}

func onDone() {

	cp.APf("new2", "func New%s2() *%s{", class, class)
	cp.APf("new2", "  this := &%s{}", class)
	switch wclass {
	case "QWidget", "QMainWindow", "QDialog":
		cp.APf("new2", "  w := qtwidgets.New%s(nil, 0)", wclass)
	default:
		log.Println("wtf...", wclass)
	}
	cp.APf("new2", "  this.SetupUi(w)")
	cp.APf("new2", "  return this")
	cp.APf("new2", "}")

	cp.APf("done", "")
	cp.APf("done", "func (this *%s) QWidget_PTR() *qtwidgets.QWidget {", class)
	cp.APf("done", "  return this.%s.QWidget_PTR()", woname)
	cp.APf("done", "}")
	cp.APf("done", "")
}

func saveCode() {
	var code string
	var mod os.FileMode = 0644

	log.Printf("saving... %s.go...\n", filep)
	code = "package main\n"
	code += cp.ExportAll()
	savefile := fmt.Sprintf("%s.go", filep)
	err := ioutil.WriteFile(savefile, []byte(code), mod)
	qtrt.ErrPrint(err, savefile)

	// gofmt the code
	gofmtPath, err := exec.LookPath("gofmt")
	qtrt.ErrPrint(err)
	cmd := exec.Command(gofmtPath, []string{"-w", savefile}...)
	err = cmd.Run()
	qtrt.ErrPrint(err, cmd)
}

func colon2uline(s string) string { return strings.Replace(s, ":", "_", -1) }
func untitle(s string) string     { return strings.ToLower(s[0:1]) + s[1:] }
