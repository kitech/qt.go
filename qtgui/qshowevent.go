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

type QShowEvent struct {
	*qtcore.QEvent
}

func (this *QShowEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QShowEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQShowEventFromPointer(cthis unsafe.Pointer) *QShowEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QShowEvent{bcthis0}
}
func (*QShowEvent) NewFromPointer(cthis unsafe.Pointer) *QShowEvent {
	return NewQShowEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:493
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QShowEvent()
func NewQShowEvent() *QShowEvent {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QShowEventC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQShowEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQShowEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:494
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QShowEvent()
func DeleteQShowEvent(this *QShowEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QShowEventD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
