package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

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
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
	*qtcore.QEvent
}

func (this *QDragLeaveEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QDragLeaveEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQDragLeaveEventFromPointer(cthis unsafe.Pointer) *QDragLeaveEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QDragLeaveEvent{bcthis0}
}
func (*QDragLeaveEvent) NewFromPointer(cthis unsafe.Pointer) *QDragLeaveEvent {
	return NewQDragLeaveEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:671
// index:0
// Public
// void QDragLeaveEvent()
func NewQDragLeaveEvent() *QDragLeaveEvent {
	cthis := qtrt.Calloc(1, 256) // 24
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QDragLeaveEventC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQDragLeaveEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:672
// index:0
// Public virtual
// void ~QDragLeaveEvent()
func DeleteQDragLeaveEvent(*QDragLeaveEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QDragLeaveEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

//  body block end
