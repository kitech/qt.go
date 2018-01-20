//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
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
type QExposeEvent struct {
	*qtcore.QEvent
}

func (this *QExposeEvent) GetCthis() unsafe.Pointer {
	return this.QEvent.GetCthis()
}
func NewQExposeEventFromPointer(cthis unsafe.Pointer) *QExposeEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QExposeEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:434
// index:0
// Public
// void QExposeEvent(const class QRegion &)
func NewQExposeEvent(rgn *QRegion) *QExposeEvent {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = rgn.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QExposeEventC2ERK7QRegion", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQExposeEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:435
// index:0
// Public virtual
// void ~QExposeEvent()
func DeleteQExposeEvent(*QExposeEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QExposeEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:437
// index:0
// Public inline
// const QRegion & region()
func (this *QExposeEvent) Region() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QExposeEvent6regionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
