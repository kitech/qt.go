package main

import (
	"os"
	"testing"

	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtwidgets"
)

// using go test tests/qxxx_test.go

func Test0(t *testing.T) {
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)
	pix := qtgui.NewQPixmap3("/home/me/oss/src/tox-homeserver/tofia/app/src/main/res/drawable/icon_avatar_40.png", "dummy123", 0)
	lab := qtwidgets.NewQLabel1("hehehe", nil, 0)
	lab.SetPixmap(pix)
	lab.Show()
	app.Exec()
}
