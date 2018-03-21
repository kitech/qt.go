package main

import (
	"os"
	"testing"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtwidgets"
)

// f: return for hold ref
func runCoreApp(f func() interface{}) {
	app := qtcore.NewQCoreApplication(len(os.Args), os.Args, 0)
	retref := f()
	_ = retref
	app.Exec()
}

func runGuiApp(f func() interface{}) {
	app := qtgui.NewQGuiApplication(len(os.Args), os.Args, 0)
	retref := f()
	_ = retref
	app.Exec()
}

func runWidgetApp(f func() interface{}) {
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)
	retref := f()
	_ = retref
	app.Exec()
}

// using go test -v (-run RunCmdInSlot) tests/qxxx_test.go
func TestRunCmdInSlot(t *testing.T) {

}
