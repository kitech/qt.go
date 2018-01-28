package main

import (
	"gopp"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
	"unsafe"

	ffiqt "qt.go/cffiqt"
	ffiqtalt "qt.go/cffiqtalt"
	"qt.go/qtcore"
	"qt.go/qtgui"
	"qt.go/qtrt"
)

func Test1(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	app := qtgui.NewQGuiApplication(len(os.Args), os.Args, 0)
	_ = app

	td := qtgui.NewQTextDocument(nil)
	log.Println(td)
	td.SetPlainText(qtcore.NewQString_5(`rv, err := ffiqt.InvokeQtFunc7\n("_ZNK5QSize9boundedToERKS_", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)`))
	td.AdjustSize()
	{
		szo := td.Size()
		log.Println(szo)
		log.Println(szo.IsEmpty(), szo.IsNull(), szo.IsValid())
		log.Println(szo.Width(), szo.Height())
		log.Println(szo.Rwidth(), szo.Rheight())
		log.Println(td.PageCount(), td.TextWidth())
	}

	{
		szo := td.Size()
		rv, err := ffiqt.InvokeQtFunc6("_ZN6QSizeF7rheightEv", ffiqt.FFI_TYPE_POINTER, szo.GetCthis())
		gopp.ErrPrint(err, rv)
		//  return rv
		log.Println(ffiqtalt.Cpretval2go("float64", rv).(float64))
	}

	{
		szo := td.Size()
		szo.SetHeight(6.7)
		szo.SetWidth(8.9)
		log.Println(szo.IsEmpty(), szo.IsNull(), szo.IsValid())
		log.Println(szo.Width(), szo.Height())
		log.Println(szo.Rwidth(), szo.Rheight())
	}

	{
		this := td
		mv := qtrt.Calloc(1, 256)
		log.Println(mv)
		rv, err := ffiqt.InvokeQtFunc7("_ZNK13QTextDocument4sizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
		gopp.ErrPrint(err, rv)
		//  return rv
		log.Println(unsafe.Pointer(uintptr(rv)), mv)
		rv = uint64(uintptr(mv))
		rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
		log.Println(rv2)
		log.Println(rv2.Width(), rv2.Height())

	}
}
