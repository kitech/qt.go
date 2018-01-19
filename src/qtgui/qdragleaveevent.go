//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
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
type QDragLeaveEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qevent.h:671
// index:0
// void QDragLeaveEvent()
func NewQDragLeaveEvent() *QDragLeaveEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QDragLeaveEventC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QDragLeaveEvent{cthis}
}

// /usr/include/qt/QtGui/qevent.h:672
// index:0
// virtual
// void ~QDragLeaveEvent()
func DeleteQDragLeaveEvent(*QDragLeaveEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QDragLeaveEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

//  body block end
