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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qevent.h:434
// index:0
// void QExposeEvent(const class QRegion &)
func NewQExposeEvent(rgn unsafe.Pointer) *QExposeEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QExposeEventC2ERK7QRegion", ffiqt.FFI_TYPE_VOID, cthis, rgn)
	gopp.ErrPrint(err, rv)
	return &QExposeEvent{cthis}
}

// /usr/include/qt/QtGui/qevent.h:435
// index:0
// virtual
// void ~QExposeEvent()
func DeleteQExposeEvent(*QExposeEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QExposeEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:437
// index:0
// inline
// const QRegion & region()
func (this *QExposeEvent) Region() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QExposeEvent6regionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
