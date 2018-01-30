package main

import "qt.go/qtcore"
import "qt.go/qtwidgets"

func main() {
	argv := []string{"./guiapp", "-v", "-x"}
	app := qtwidgets.NewQApplication(len(argv), argv, 0)
	btn := qtwidgets.NewQPushButton_1(qtcore.NewQString_5("abcbtn"), nil)
	btn.Show()
	// app.Exit(0)
	app.Exec()
}
