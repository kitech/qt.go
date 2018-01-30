package main

/*
 */
import "C"
import (
	"gopp"
	"log"
	"unsafe"

	ffiqt "qt.go/cffiqt"
	"qt.go/qtcore"
	"qt.go/qtrt"
	"qt.go/qtwidgets"
)

func tLength(ptr unsafe.Pointer) int {
	rv, err := ffiqt.InvokeQtFunc6("C_ZNK7QString6lengthEv", ffiqt.FFI_TYPE_POINTER, ptr)
	gopp.ErrPrint(err, rv)
	//  return rv
	log.Println(rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

func main() {
	// rs := qtrt.CString("hehehg呵呵")
	txt := qtcore.NewQString_5("hehehg呵呵")
	log.Println(tLength(txt.GetCthis()), txt.IsEmpty())
	log.Println(txt.ToUtf8().Data())

	// os.Exit(0)
	argv := []string{"./guiapp", "-v", "-x"}
	app := qtwidgets.NewQApplication(len(argv), argv, 0)
	log.Println(app)
	// var p unsafe.Pointer
	// p := qtwidgets.NewQWidgetFromPointer(nil)
	btn := qtwidgets.NewQPushButton(nil)
	log.Println(btn, btn.GetCthis())
	btn.SetText(txt)
	btn.Show()

	// var w = (*qtwidgets.QWidget)(btn)
	// w.Show()
	app.Exec()
}
