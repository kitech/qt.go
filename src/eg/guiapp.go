package main

/*
 */
import "C"
import (
	"log"
	"qtcore"
	"qtwidgets"
	"unsafe"
)

func main() {
	// rs := qtrt.CString("hehehg呵呵")
	txt := qtcore.NewQString_5("hehehg呵呵")
	log.Println(txt.Length())

	argv := []string{"./guiapp", "-v", "-x"}
	app := qtwidgets.NewQApplication(len(argv), argv, 0)
	log.Println(app)
	var p unsafe.Pointer
	btn := qtwidgets.NewQPushButton(p)
	log.Println(btn)
	// btn.SetText(txt.GetCthis())
	btn.SetText(txt)
	btn.Show()

	// var w = (*qtwidgets.QWidget)(btn)
	// w.Show()
	app.Exec()
}
