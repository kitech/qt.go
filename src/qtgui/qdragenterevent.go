//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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
type QDragEnterEvent struct {
	*QDragMoveEvent
}

func (this *QDragEnterEvent) GetCthis() unsafe.Pointer {
	return this.QDragMoveEvent.GetCthis()
}
func NewQDragEnterEventFromPointer(cthis unsafe.Pointer) *QDragEnterEvent {
	bcthis0 := NewQDragMoveEventFromPointer(cthis)
	return &QDragEnterEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:664
// index:0
// Public virtual
// void ~QDragEnterEvent()
func DeleteQDragEnterEvent(*QDragEnterEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QDragEnterEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

//  body block end
