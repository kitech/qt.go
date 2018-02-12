package main

import (
	"log"
	"os"
	"testing"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtrt"
)

func Test0(t *testing.T) {
	app := qtcore.NewQCoreApplication(len(os.Args), os.Args, 0)
	log.Println(app)

	{ // ok
		rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject18connectSlotsByNameEP7QObject", qtrt.FFI_TYPE_POINTER, app.GetCthis())
		qtrt.ErrPrint(err, rv)
	}
	{
		var nilthis *qtcore.QMetaObject
		nilthis.ConnectSlotsByName(qtcore.NewQObjectFromPointer(app.GetCthis()))
	}
	{
		qtcore.QMetaObject_ConnectSlotsByName(qtcore.NewQObjectFromPointer(app.GetCthis()))
	}
}
