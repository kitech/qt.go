package qtgui

// /usr/include/qt/QtGui/qpen.h
// #include <qpen.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*
 */
// size 8
type QPen struct {
	*qtrt.CObject
}
type QPen_ITF interface {
	QPen_PTR() *QPen
}

func (ptr *QPen) QPen_PTR() *QPen { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QPenFromptr(cthis Voidptr) *QPen {
	return &QPen{&qtrt.CObject{cthis}}
}
func (*QPen) Fromptr(cthis Voidptr) *QPen {
	return QPenFromptr(cthis)
}

// /usr/include/qt/QtGui/qpen.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPen()

/*
 */
func (*QPen) NewForInherit() *QPen {
	return NewQPen()
}
func NewQPen() *QPen {
	cthis := qtrt.Malloc(8)
	rv, err := qtrt.Qtcc3(46253105, "_ZN4QPenC2Ev", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, Voidptr(&cthis))
	qtrt.ErrPrint2(err, rv)
	gothis := QPenFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQPen)
	return gothis
}

func DeleteQPen(this *QPen) {
	rv, err := qtrt.Qtcc1(0, "_ZN4QPenD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10145() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
