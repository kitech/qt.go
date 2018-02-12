package qtcore

// /usr/include/qt/QtCore/qeventtransition.h
// #include <qeventtransition.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// bool eventTest(class QEvent *)
func (this *QEventTransition) InheritEventTest(f func(event *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventTest", f)
}

// void onTransition(class QEvent *)
func (this *QEventTransition) InheritOnTransition(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "onTransition", f)
}

// bool event(class QEvent *)
func (this *QEventTransition) InheritEvent(f func(e *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

type QEventTransition struct {
	*QAbstractTransition
}

func (this *QEventTransition) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractTransition.GetCthis()
	}
}
func (this *QEventTransition) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractTransition = NewQAbstractTransitionFromPointer(cthis)
}
func NewQEventTransitionFromPointer(cthis unsafe.Pointer) *QEventTransition {
	bcthis0 := NewQAbstractTransitionFromPointer(cthis)
	return &QEventTransition{bcthis0}
}
func (*QEventTransition) NewFromPointer(cthis unsafe.Pointer) *QEventTransition {
	return NewQEventTransitionFromPointer(cthis)
}

// /usr/include/qt/QtCore/qeventtransition.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QEventTransition) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QEventTransition10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qeventtransition.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QEventTransition(QState *)
func NewQEventTransition(sourceState *QState /*777 QState **/) *QEventTransition {
	var convArg0 = sourceState.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QEventTransitionC2EP6QState", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQEventTransitionFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qeventtransition.h:58
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QEventTransition(QObject *, QEvent::Type, QState *)
func NewQEventTransition_1(object *QObject /*777 QObject **/, type_ int, sourceState *QState /*777 QState **/) *QEventTransition {
	var convArg0 = object.GetCthis()
	var convArg2 = sourceState.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QEventTransitionC2EP7QObjectN6QEvent4TypeEP6QState", qtrt.FFI_TYPE_POINTER, convArg0, type_, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQEventTransitionFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qeventtransition.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QEventTransition()
func DeleteQEventTransition(this *QEventTransition) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QEventTransitionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qeventtransition.h:61
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * eventSource()
func (this *QEventTransition) EventSource() *QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QEventTransition11eventSourceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qeventtransition.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEventSource(QObject *)
func (this *QEventTransition) SetEventSource(object *QObject /*777 QObject **/) {
	var convArg0 = object.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QEventTransition14setEventSourceEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventtransition.h:64
// index:0
// Public Visibility=Default Availability=Available
// [4] QEvent::Type eventType()
func (this *QEventTransition) EventType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QEventTransition9eventTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qeventtransition.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEventType(QEvent::Type)
func (this *QEventTransition) SetEventType(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QEventTransition12setEventTypeEN6QEvent4TypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventtransition.h:68
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventTest(QEvent *)
func (this *QEventTransition) EventTest(event *QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QEventTransition9eventTestEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qeventtransition.h:69
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void onTransition(QEvent *)
func (this *QEventTransition) OnTransition(event *QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QEventTransition12onTransitionEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventtransition.h:71
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QEventTransition) Event(e *QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QEventTransition5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
		qtrt.KeepMe()
	}
}

//  keep block end
