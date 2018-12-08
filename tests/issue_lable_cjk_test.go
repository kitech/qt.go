package main

import (
	"os"
	"testing"
	"time"

	"github.com/kitech/qt.go/qtcore"

	"github.com/kitech/qt.go/qtwidgets"
)

func Test0(t *testing.T) {
	time.Sleep(5 * time.Second)
	// qtrt.DebugFinal = false

	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)

	lab := qtwidgets.NewQLabelp1(nil)
	lab.SetText(qtcore.QCoreApplication_Translate("MainWindow", "TextLabel\345\223\210\345\226\275", "dummy123", 0))
	lab.Show()

	app.Exec()
	time.Sleep(50 * time.Second)
}
