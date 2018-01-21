package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

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
type QHideEvent struct {
	*qtcore.QEvent
}

func (this *QHideEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func NewQHideEventFromPointer(cthis unsafe.Pointer) *QHideEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QHideEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:501
// index:0
// Public
// void QHideEvent()
func NewQHideEvent() *QHideEvent {
	cthis := qtrt.Calloc(1, 256) // 24
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QHideEventC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQHideEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:502
// index:0
// Public virtual
// void ~QHideEvent()
func DeleteQHideEvent(*QHideEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QHideEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

//  body block end
