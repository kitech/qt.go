package main

import (
	"log"

	"qt.go/qtwidgets"
)

func logp(args ...interface{}) { log.Println(args...) }

func main() {

	argv := []string{"./guiapp", "-v", "-x"}
	app := qtwidgets.NewQApplication(len(argv), argv, 0)
	log.Println(app)

	mw := qtwidgets.NewQMainWindow(nil, 0)
	uiw := NewUi_MainWindow()
	uiw.SetupUi(mw)

	mw.Show()
	app.Exec()
}
