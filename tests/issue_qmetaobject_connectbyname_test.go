package main

import (
	"gopp"
	"log"
	"os"
	"testing"

	ffiqt "qt.go/cffiqt"
	"qt.go/qtcore"
)

func Test0(t *testing.T) {
	app := qtcore.NewQCoreApplication(len(os.Args), os.Args, 0)
	log.Println(app)

	{ // ok
		rv, err := ffiqt.InvokeQtFunc6("_ZN11QMetaObject18connectSlotsByNameEP7QObject", ffiqt.FFI_TYPE_POINTER, app.GetCthis())
		gopp.ErrPrint(err, rv)
	}
	{
		var nilthis *qtcore.QMetaObject
		nilthis.ConnectSlotsByName(qtcore.NewQObjectFromPointer(app.GetCthis()))
	}
	{
		qtcore.QMetaObject_ConnectSlotsByName(qtcore.NewQObjectFromPointer(app.GetCthis()))
	}
}
