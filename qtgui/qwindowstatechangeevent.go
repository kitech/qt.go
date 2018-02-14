package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QWindowStateChangeEvent struct {
	*qtcore.QEvent
}
type QWindowStateChangeEvent_ITF interface {
	qtcore.QEvent_ITF
	QWindowStateChangeEvent_PTR() *QWindowStateChangeEvent
}

func (ptr *QWindowStateChangeEvent) QWindowStateChangeEvent_PTR() *QWindowStateChangeEvent { return ptr }

func (this *QWindowStateChangeEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QWindowStateChangeEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQWindowStateChangeEventFromPointer(cthis unsafe.Pointer) *QWindowStateChangeEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QWindowStateChangeEvent{bcthis0}
}
func (*QWindowStateChangeEvent) NewFromPointer(cthis unsafe.Pointer) *QWindowStateChangeEvent {
	return NewQWindowStateChangeEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:783
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWindowStateChangeEvent(Qt::WindowStates, _Bool)
func NewQWindowStateChangeEvent(aOldState int, isOverride bool) *QWindowStateChangeEvent {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWindowStateChangeEventC2E6QFlagsIN2Qt11WindowStateEEb", qtrt.FFI_TYPE_POINTER, aOldState, isOverride)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWindowStateChangeEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWindowStateChangeEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:784
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWindowStateChangeEvent()
func DeleteQWindowStateChangeEvent(this *QWindowStateChangeEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWindowStateChangeEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:786
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::WindowStates oldState()
func (this *QWindowStateChangeEvent) OldState() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QWindowStateChangeEvent8oldStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:787
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isOverride()
func (this *QWindowStateChangeEvent) IsOverride() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QWindowStateChangeEvent10isOverrideEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
