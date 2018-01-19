//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

func init() {
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
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QPaintEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qevent.h:405
// index:0
// void QPaintEvent(const class QRegion &)
func NewQPaintEvent(paintRegion unsafe.Pointer) *QPaintEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPaintEventC2ERK7QRegion", ffiqt.FFI_TYPE_VOID, cthis, paintRegion)
	gopp.ErrPrint(err, rv)
	return &QPaintEvent{cthis}
}

// /usr/include/qt/QtGui/qevent.h:406
// index:1
// void QPaintEvent(const class QRect &)
func NewQPaintEvent_1(paintRect unsafe.Pointer) *QPaintEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPaintEventC2ERK5QRect", ffiqt.FFI_TYPE_VOID, cthis, paintRect)
	gopp.ErrPrint(err, rv)
	return &QPaintEvent{cthis}
}

// /usr/include/qt/QtGui/qevent.h:407
// index:0
// virtual
// void ~QPaintEvent()
func DeleteQPaintEvent(*QPaintEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPaintEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:409
// index:0
// inline
// const QRect & rect()
func (this *QPaintEvent) Rect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPaintEvent4rectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:410
// index:0
// inline
// const QRegion & region()
func (this *QPaintEvent) Region() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPaintEvent6regionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
