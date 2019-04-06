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
// extern C begin: 3
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

/*

 */
type QActionEvent struct {
	*qtcore.QEvent
}
type QActionEvent_ITF interface {
	qtcore.QEvent_ITF
	QActionEvent_PTR() *QActionEvent
}

func (ptr *QActionEvent) QActionEvent_PTR() *QActionEvent { return ptr }

func (this *QActionEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QActionEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQActionEventFromPointer(cthis unsafe.Pointer) *QActionEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QActionEvent{bcthis0}
}
func (*QActionEvent) NewFromPointer(cthis unsafe.Pointer) *QActionEvent {
	return NewQActionEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:732
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QActionEvent(int, QAction *, QAction *)

/*

 */
func (*QActionEvent) NewForInherit(type_ int, action unsafe.Pointer /*666*/, before unsafe.Pointer /*666*/) *QActionEvent {
	return NewQActionEvent(type_, action, before)
}
func NewQActionEvent(type_ int, action unsafe.Pointer /*666*/, before unsafe.Pointer /*666*/) *QActionEvent {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QActionEventC2EiP7QActionS1_", qtrt.FFI_TYPE_POINTER, type_, action, before)
	qtrt.ErrPrint(err, rv)
	gothis := NewQActionEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQActionEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:732
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QActionEvent(int, QAction *, QAction *)

/*

 */
func (*QActionEvent) NewForInheritp(type_ int, action unsafe.Pointer /*666*/) *QActionEvent {
	return NewQActionEventp(type_, action)
}
func NewQActionEventp(type_ int, action unsafe.Pointer /*666*/) *QActionEvent {
	// arg: 2, QAction *=Pointer, QAction=Record, , Invalid
	var before unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QActionEventC2EiP7QActionS1_", qtrt.FFI_TYPE_POINTER, type_, action, before)
	qtrt.ErrPrint(err, rv)
	gothis := NewQActionEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQActionEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:733
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QActionEvent()

/*

 */
func DeleteQActionEvent(this *QActionEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QActionEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 40)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:735
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QAction * action() const

/*

 */
func (this *QActionEvent) Action() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QActionEvent6actionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qevent.h:736
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QAction * before() const

/*

 */
func (this *QActionEvent) Before() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QActionEvent6beforeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

//  body block end

//  keep block begin

func init_unused_10699() {
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
