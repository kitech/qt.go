//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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
type QShowEvent struct {
	*qtcore.QEvent
}

func (this *QShowEvent) GetCthis() unsafe.Pointer {
	return this.QEvent.GetCthis()
}

// /usr/include/qt/QtGui/qevent.h:493
// index:0
// void QShowEvent()
func NewQShowEvent() *QShowEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QShowEventC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQShowEventFromPointer(cthis)
	return gothis
}
func NewQShowEventFromPointer(cthis unsafe.Pointer) *QShowEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QShowEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:494
// index:0
// virtual
// void ~QShowEvent()
func DeleteQShowEvent(*QShowEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QShowEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

//  body block end
