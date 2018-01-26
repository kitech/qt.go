package main

import (
	"gopp"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"sort"
	"strings"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/clbanning/mxj"
)

// 使用.ui文件的转换

var file string
var filep string // file name without externsion

type convContext struct {
	topCls string
	cp     *CodePager
	djso   *gopp.Json

	curobjName string
	curobjCls  string
	curloName  string // layout, used for layout->addWidget
	topobjName string // MainWIndow/MainForm...
}

var cvctx = &convContext{}

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	file = os.Args[1]
	filep = file[:len(file)-2]
	log.Println(file, filep)

	fp, err := os.Open(file)
	gopp.ErrPrint(err, fp)

	xml, err := mxj.NewMapXmlReader(fp)
	gopp.ErrPrint(err, xml)
	bcc, err := xml.Json()
	gopp.ErrPrint(err)
	jso, err := simplejson.NewJson(bcc)
	gopp.ErrPrint(err)

	cvctx.djso = gopp.NewJsonFromObject(jso)
	cvctx.cp = NewCodePager()
	cvctx.cp.APf("header", "package main")
	cvctx.cp.APf("import", "import \"qtcore\"")
	cvctx.cp.APf("import", "// import \"qtgui\"")
	cvctx.cp.APf("import", "import \"qtwidgets\"")
	cvctx.cp.APf("import", "import \"qtmock\"")
	visitJson(jso, 0)
	cvctx.cp.APf("struct", "}")

	cvctx.cp.APf("setupUi", "  %s.SetCentralWidget(this.Centralwidget)", cvctx.topobjName)
	cvctx.cp.APf("setupUi", "  this.RetranslateUi(%s)", cvctx.topobjName) // TODO binding bug
	cvctx.cp.APf("setupUi", "} // setupUi")

	cvctx.cp.APf("retranslateUi", "} // retranslateUi")

	log.Println(cvctx.cp.ExportAll())
	ioutil.WriteFile("/tmp/ui_xxx.go", []byte(cvctx.cp.ExportAll()), 0644)
}

// 思路，应该是递归遍历查找widget/layout/action对象，
// 对于属于的处理，应该立即处理，而不是递归，因为可能有顺序问题。
// 第一个Widget不在创建
// 如果是QMainWindow，
// 对于layout 的parent
func visitJson(jso *simplejson.Json, depth int) {
	cpr := cvctx.cp

	// 只处理object对象，name指当前object的名字
	process_obj := func(jso *simplejson.Json, name string) {
		log.Println(name, depth)
		switch name {
		case "ui":
			log.Println("version:", jso.GetPath("-version").MustString())
		case "class":
			cpr.APf("struct", "type Ui_%v struct {", jso.Interface())
			cpr.APf("ctor", "func NewUi_%s () *Ui_%s {", jso.Interface(), jso.Interface())
			cpr.APf("ctor", "  return &Ui_%s{}", jso.Interface())
			cpr.APf("ctor", "}")
		case "widget":
			clsn := jso.Get("-class").MustString()
			varn := jso.Get("-name").MustString()
			log.Println(clsn, varn)
			if clsn == "" && varn == "" {
				log.Println(clsn, varn)
				break
			}
			cvctx.curobjName = varn
			cvctx.curobjCls = clsn
			cpr.APf("setupUi", "")
			if depth == 1 { // top widget
				cpr.APf("setupUi", "func (this *Ui_%v) SetupUi(%s *qtwidgets.%s) {",
					varn, varn, clsn)
				cpr.APf("retranslateUi", "func (this *Ui_%v) RetranslateUi(%s *qtwidgets.%s) {",
					varn, varn, clsn)
				cpr.APf("setupUi", "  this.%s = %s", varn, varn)
				cpr.APf("setupUi", "  if this.%s.ObjectName().IsEmpty(){", varn)
				cpr.APf("setupUi", "    this.%s.SetObjectName(qtcore.NewQString_5(\"%s\"))", varn, varn)
				cpr.APf("setupUi", "  }")
				cvctx.topobjName = varn
			}

			switch clsn {
			case "QSplitter", "QToolButton", "QLineEdit", "QListWidget", "QScrollArea":
				cpr.APf("setupUi", "  this.%s = qtwidgets.New%s(qtwidgets.NewQWidgetFromPointer(%s.GetCthis()))", strings.Title(varn), clsn, cvctx.topobjName)
			default:
				if varn == cvctx.topobjName {
				} else {
					cpr.APf("setupUi", "  this.%s = qtwidgets.New%s(qtwidgets.NewQWidgetFromPointer(%s.GetCthis()), 0)", strings.Title(varn), clsn, cvctx.topobjName)
				}
			}

			if varn == cvctx.topobjName {

			} else {
				cpr.APf("setupUi", "  this.%s.SetObjectName(qtcore.NewQString_5(\"%s\"))", strings.Title(varn), varn)
			}
			if cvctx.curloName != "" {
				cpr.APf("setupUi", "  this.%s.AddWidget(qtwidgets.NewQWidgetFromPointer(this.%s.GetCthis()))",
					strings.Title(cvctx.curloName), strings.Title(varn))
			}

			if varn == cvctx.topobjName {
			} else {
			}
			cpr.APf("struct", "  %s *qtwidgets.%s", strings.Title(varn), clsn)
		case "layout":
			clsn := jso.Get("-class").MustString()
			varn := jso.Get("-name").MustString()
			log.Println(clsn, varn)
			if clsn == "" && varn == "" {
				log.Println(clsn, varn)
				break
			}
			cpr.APf("setupUi", "")
			if cvctx.curobjName == "centralwidget" {
				cpr.APf("setupUi", "  this.%s = qtwidgets.New%s_1(this.Centralwidget)", strings.Title(varn), clsn)
			} else {
				cpr.APf("setupUi", "  this.%s = qtwidgets.New%s()", strings.Title(varn), clsn)
			}
			cpr.APf("setupUi", "  this.%s.SetObjectName(qtcore.NewQString_5(\"%s\"))", strings.Title(varn), varn)

			cpr.APf("struct", "  %s *qtwidgets.%s", strings.Title(varn), clsn)

			cvctx.curobjName = varn
			cvctx.curobjCls = clsn
			cvctx.curloName = varn
		case "action":
			// TODO QToolBar, QStatusBar, QMenu
		default:
		}

		setProp := func(pname string, pval *simplejson.Json) {
			switch pname {
			case "spacing":
				log.Println(pname, pval.Get("number").MustString(), cvctx.curobjName)
				cpr.APf("setupUi", "  this.%s.Set%s(%s)",
					strings.Title(cvctx.curobjName), strings.Title(pname), pval.Get("number").MustString())
			case "geometry":
				widthx := pval.Get("rect").Get("width").MustString()
				heightx := pval.Get("rect").Get("height").MustString()
				cpr.APf("setupUi", "  this.%s.Set%s(0, 0, %s, %s)",
					strings.Title(cvctx.curobjName), strings.Title(pname), widthx, heightx)
			case "minimumSize", "maximumSize", "iconSize":
				log.Println(pname, pval.Interface(), cvctx.curobjName)
				widthx := pval.Get("size").Get("width").MustString()
				heightx := pval.Get("size").Get("height").MustString()
				cpr.APf("setupUi", "  this.%s.Set%s(qtcore.NewQSize_1(%s, %s))",
					strings.Title(cvctx.curobjName), strings.Title(pname), widthx, heightx)
			case "orientation":
				log.Println(pname, pval.Interface(), cvctx.curobjName)
				orien := pval.Get("enum").MustString()
				orien = "qtcore." + strings.Replace(orien, "::", "__", -1)
				switch cvctx.curobjCls {
				case "QSplitter":
					cpr.APf("setupUi", "  this.%s.Set%s(%s)",
						strings.Title(cvctx.curobjName), strings.Title(pname), orien)
				case "QVBoxLayout", "QHBoxLayout":
					// 这个并不需要，并且没有这个函数，即使有的话，也是direction属性
				}
			case "text", "placeholderText":
				// TODO escape special chars
				cpr.APf("retranslateUi", "  this.%s.Set%s(qtmock.QCoreApplication_Translate(\"%s\", \"%s\", \"dummy123\", 0))",
					strings.Title(cvctx.curobjName), strings.Title(pname), cvctx.topobjName,
					pval.Get("string").MustString())

			}
		}

		// process property right now
		propjsos, has := jso.CheckGet("property")
		if has {
			log.Println(reflect.TypeOf(propjsos.Interface()).Kind())
			valsx := propjsos.Interface()
			valsv := reflect.ValueOf(valsx)
			switch valsv.Kind() {
			case reflect.Map:
				vals := valsx.(map[string]interface{})
				keys_ := []string{}
				for key_, _ := range vals {
					keys_ = append(keys_, key_)
				}
				sort.Strings(keys_)
				for _, key_ := range keys_ {
					valx := vals[key_]
					log.Println(key_, valx)
					switch key_ {
					case "-name":
						setProp(valx.(string), propjsos)
					}
				}
			case reflect.Slice:
				for i := 0; i < len(propjsos.MustArray()); i++ {
					pjso := propjsos.GetIndex(i)
					pname := pjso.Get("-name").MustString()
					setProp(pname, pjso)
				}
			}
		}
	}

	_ = process_obj

	valsx := jso.Interface()
	valsv := reflect.ValueOf(valsx)
	switch valsv.Kind() {
	case reflect.Map:
		vals := valsx.(map[string]interface{}) // TODO 排序
		names := []string{}
		for name_, _ := range vals {
			names = append(names, name_)
		}
		sort.Strings(names)
		for idx, name_ := range names {
			val := vals[name_]
			log.Println(idx, name_, val == nil)
			if strings.HasPrefix(name_, "-") {
				// attr
			} else {
				process_obj(jso.Get(name_), name_)
				visitJson(jso.Get(name_), depth+1)
			}
		}
	case reflect.Slice:
		vals := valsx.([]interface{})
		for idx, val := range vals {
			log.Println(idx, depth, val == nil)
			visitJson(jso.GetIndex(idx), depth+1)
		}
	case reflect.Invalid:
		log.Println(valsv.Kind().String(), valsv)
	default:
		log.Println(valsv.Kind().String())
	}
}

func visitXml(xml mxj.Map, name, path string, depth int) {
	cpr := cvctx.cp

	if path == "ui" {
		// log.Println(xml.ValueForPath(path + ".-version"))
	} else {
		log.Println(path, depth)
		// log.Println(xml.Attributes(path))
		switch name {
		case "widget":
			clsn, err := xml.ValueForPath(path + ".-class")
			varn, err := xml.ValueForPath(path + ".-name")
			gopp.ErrPrint(err)
			log.Println(clsn, varn)
			cpr.APf("struct", "%s *qtwidgets.%s", varn, clsn)
		case "layout":
			clsn, err := xml.ValueForPath(path + ".-class")
			varn, err := xml.ValueForPath(path + ".-name")
			gopp.ErrPrint(err)
			log.Println(clsn, varn)
			cpr.APf("struct", "%s *qtwidgets.%s", varn, clsn)

		case "property":

		default:
		}
	}

	relems, err := xml.Elements(path)
	gopp.ErrPrint(err, depth)

	for _, elem := range relems {
		fpath := path + "." + elem
		vals := xml.PathsForKey(fpath)
		gopp.ErrPrint(err, vals)
		log.Println("for vals:", fpath, depth)
		// visit(xml, elem, fpath, depth+1)
	}
}
