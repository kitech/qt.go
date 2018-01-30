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
import "gopp"
import "qt.go/cffiqt"
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
func (this *QHideEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQHideEventFromPointer(cthis unsafe.Pointer) *QHideEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QHideEvent{bcthis0}
}
func (*QHideEvent) NewFromPointer(cthis unsafe.Pointer) *QHideEvent {
	return NewQHideEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:501
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QHideEvent()
func NewQHideEvent() *QHideEvent {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QHideEventC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQHideEventFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:502
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QHideEvent()
func DeleteQHideEvent(*QHideEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QHideEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

//  body block end
