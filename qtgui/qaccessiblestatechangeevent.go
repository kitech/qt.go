package qtgui

// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QAccessibleStateChangeEvent struct {
	*QAccessibleEvent
}
type QAccessibleStateChangeEvent_ITF interface {
	QAccessibleEvent_ITF
	QAccessibleStateChangeEvent_PTR() *QAccessibleStateChangeEvent
}

func (ptr *QAccessibleStateChangeEvent) QAccessibleStateChangeEvent_PTR() *QAccessibleStateChangeEvent {
	return ptr
}

func (this *QAccessibleStateChangeEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAccessibleEvent.GetCthis()
	}
}
func (this *QAccessibleStateChangeEvent) SetCthis(cthis unsafe.Pointer) {
	this.QAccessibleEvent = NewQAccessibleEventFromPointer(cthis)
}
func NewQAccessibleStateChangeEventFromPointer(cthis unsafe.Pointer) *QAccessibleStateChangeEvent {
	bcthis0 := NewQAccessibleEventFromPointer(cthis)
	return &QAccessibleStateChangeEvent{bcthis0}
}
func (*QAccessibleStateChangeEvent) NewFromPointer(cthis unsafe.Pointer) *QAccessibleStateChangeEvent {
	return NewQAccessibleStateChangeEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qaccessible.h:733
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAccessibleStateChangeEvent()
func DeleteQAccessibleStateChangeEvent(this *QAccessibleStateChangeEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAccessibleStateChangeEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 40)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qaccessible.h:735
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QAccessible::State changedStates()
func (this *QAccessibleStateChangeEvent) ChangedStates() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QAccessibleStateChangeEvent13changedStatesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
