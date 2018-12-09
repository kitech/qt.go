// +build !minimal

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
// extern C begin: 19
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// bool eventTest(QEvent *)
func (this *QEventTransition) InheritEventTest(f func(event *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventTest", f)
}

// void onTransition(QEvent *)
func (this *QEventTransition) InheritOnTransition(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "onTransition", f)
}

// bool event(QEvent *)
func (this *QEventTransition) InheritEvent(f func(e *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

/*

 */
type QEventTransition struct {
	*QAbstractTransition
}
type QEventTransition_ITF interface {
	QAbstractTransition_ITF
	QEventTransition_PTR() *QEventTransition
}

func (ptr *QEventTransition) QEventTransition_PTR() *QEventTransition { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QEventTransition) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QEventTransition10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qeventtransition.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QEventTransition(QState *)

/*
Constructs a new QEventTransition object with the given sourceState.
*/
func (*QEventTransition) NewForInherit(sourceState QState_ITF /*777 QState **/) *QEventTransition {
	return NewQEventTransition(sourceState)
}
func NewQEventTransition(sourceState QState_ITF /*777 QState **/) *QEventTransition {
	var convArg0 unsafe.Pointer
	if sourceState != nil && sourceState.QState_PTR() != nil {
		convArg0 = sourceState.QState_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QEventTransitionC2EP6QState", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQEventTransitionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QEventTransition")
	return gothis
}

// /usr/include/qt/QtCore/qeventtransition.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QEventTransition(QState *)

/*
Constructs a new QEventTransition object with the given sourceState.
*/
func (*QEventTransition) NewForInheritp() *QEventTransition {
	return NewQEventTransitionp()
}
func NewQEventTransitionp() *QEventTransition {
	// arg: 0, QState *=Pointer, QState=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN16QEventTransitionC2EP6QState", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQEventTransitionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QEventTransition")
	return gothis
}

// /usr/include/qt/QtCore/qeventtransition.h:58
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QEventTransition(QObject *, QEvent::Type, QState *)

/*
Constructs a new QEventTransition object with the given sourceState.
*/
func (*QEventTransition) NewForInherit1(object QObject_ITF /*777 QObject **/, type_ int, sourceState QState_ITF /*777 QState **/) *QEventTransition {
	return NewQEventTransition1(object, type_, sourceState)
}
func NewQEventTransition1(object QObject_ITF /*777 QObject **/, type_ int, sourceState QState_ITF /*777 QState **/) *QEventTransition {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if sourceState != nil && sourceState.QState_PTR() != nil {
		convArg2 = sourceState.QState_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QEventTransitionC2EP7QObjectN6QEvent4TypeEP6QState", qtrt.FFI_TYPE_POINTER, convArg0, type_, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQEventTransitionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QEventTransition")
	return gothis
}

// /usr/include/qt/QtCore/qeventtransition.h:58
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QEventTransition(QObject *, QEvent::Type, QState *)

/*
Constructs a new QEventTransition object with the given sourceState.
*/
func (*QEventTransition) NewForInherit1p(object QObject_ITF /*777 QObject **/, type_ int) *QEventTransition {
	return NewQEventTransition1p(object, type_)
}
func NewQEventTransition1p(object QObject_ITF /*777 QObject **/, type_ int) *QEventTransition {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	// arg: 2, QState *=Pointer, QState=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN16QEventTransitionC2EP7QObjectN6QEvent4TypeEP6QState", qtrt.FFI_TYPE_POINTER, convArg0, type_, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQEventTransitionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QEventTransition")
	return gothis
}

// /usr/include/qt/QtCore/qeventtransition.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QEventTransition()

/*

 */
func DeleteQEventTransition(this *QEventTransition) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QEventTransitionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qeventtransition.h:61
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * eventSource() const

/*
Returns the event source associated with this event transition.

Note: Getter function for property eventSource.

See also setEventSource().
*/
func (this *QEventTransition) EventSource() *QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QEventTransition11eventSourceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qeventtransition.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEventSource(QObject *)

/*
Sets the event source associated with this event transition to be the given object.

Note: Setter function for property eventSource.

See also eventSource().
*/
func (this *QEventTransition) SetEventSource(object QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QEventTransition14setEventSourceEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventtransition.h:64
// index:0
// Public Visibility=Default Availability=Available
// [4] QEvent::Type eventType() const

/*
Returns the event type that this event transition is associated with.

Note: Getter function for property eventType.

See also setEventType().
*/
func (this *QEventTransition) EventType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QEventTransition9eventTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qeventtransition.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEventType(QEvent::Type)

/*
Sets the event type that this event transition is associated with.

Note: Setter function for property eventType.

See also eventType().
*/
func (this *QEventTransition) SetEventType(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QEventTransition12setEventTypeEN6QEvent4TypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventtransition.h:68
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventTest(QEvent *)

/*
Reimplemented from QAbstractTransition::eventTest().
*/
func (this *QEventTransition) EventTest(event QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QEventTransition9eventTestEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qeventtransition.h:69
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void onTransition(QEvent *)

/*
Reimplemented from QAbstractTransition::onTransition().
*/
func (this *QEventTransition) OnTransition(event QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QEventTransition12onTransitionEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventtransition.h:71
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QEventTransition) Event(e QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
