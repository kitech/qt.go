package main

import (
	"log"
	"qtwidgets"
	"unsafe"
)

func main() {
	argv := []string{"./guiapp", "-v", "-x"}
	app := qtwidgets.NewQApplication(len(argv), argv, 0)
	log.Println(app)
	var p unsafe.Pointer
	btn := qtwidgets.NewQPushButton(p)
	log.Println(btn)
	btn.Show()
	// var w = (*qtwidgets.QWidget)(btn)
	// w.Show()
	app.Exec()
}
