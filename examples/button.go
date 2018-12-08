package main

import "github.com/kitech/qt.go/qtwidgets"

func main() {
	argv := []string{"./guiapp", "-v", "-x"}
	app := qtwidgets.NewQApplication(len(argv), argv, 0)
	btn := qtwidgets.NewQPushButton1("abcbtn中", nil)
	btn.Show()

	// app.Exit(0)
	app.Exec()
}
