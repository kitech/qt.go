package main

import (
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	gopp "github.com/kitech/goplusplus"
	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtwidgets"
)

func Test0(t *testing.T) {
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)
	_ = app
	w := qtwidgets.NewQWidget(nil, 0)
	sz := w.Size()
	log.Println(sz)
}

func Test1(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)
	_ = app
	w := qtwidgets.NewQWidget(nil, 0)
	rdh := gopp.AbsI32(rand.Int()) % 500
	rdw := gopp.AbsI32(rand.Int()) % 500
	w.SetFixedHeight(rdh)
	w.SetFixedWidth(rdw)

	szo := w.Size()
	if szo.Width() != rdw || szo.Height() != rdh {
		t.Error("size not match: want:", rdw, rdh, "ret:", szo.Width(), szo.Height())
	}
	log.Println(szo.Rheight(), szo.Height())

	w.SetObjectName(qtcore.NewQString5("hehhehhe"))
	objname := w.ObjectName()
	if objname.Length() != 8 {
		t.Error(objname)
	}

	{
		szo := qtcore.NewQSize()
		log.Println(szo.Height(), szo.Rheight())
	}

	{
		szo := qtcore.NewQSize1(0, 0)
		log.Println(szo.Height(), szo.Rheight())
	}

	{
		szo := qtcore.NewQSizeF()
		log.Println(szo.Height(), szo.Rheight(), szo.IsEmpty())
	}

	{
		szo := qtcore.NewQSizeF2(5.5, 6.5)
		log.Println(szo.Height(), szo.Rheight(), szo.IsEmpty())
	}

}
