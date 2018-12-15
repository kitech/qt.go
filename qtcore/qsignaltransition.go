// +build !minimal

package qtcore

// /usr/include/qt/QtCore/qsignaltransition.h
// #include <qsignaltransition.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
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
func (this *QSignalTransition) InheritEventTest(f func(event *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventTest", f)
}

// void onTransition(QEvent *)
func (this *QSignalTransition) InheritOnTransition(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "onTransition", f)
}

// bool event(QEvent *)
func (this *QSignalTransition) InheritEvent(f func(e *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

/*

 */
type QSignalTransition struct {
	*QAbstractTransition
}
type QSignalTransition_ITF interface {
	QAbstractTransition_ITF
	QSignalTransition_PTR() *QSignalTransition
}

func (ptr *QSignalTransition) QSignalTransition_PTR() *QSignalTransition { return ptr }

func (this *QSignalTransition) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractTransition.GetCthis()
	}
}
func (this *QSignalTransition) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractTransition = NewQAbstractTransitionFromPointer(cthis)
}
func NewQSignalTransitionFromPointer(cthis unsafe.Pointer) *QSignalTransition {
	bcthis0 := NewQAbstractTransitionFromPointer(cthis)
	return &QSignalTransition{bcthis0}
}
func (*QSignalTransition) NewFromPointer(cthis unsafe.Pointer) *QSignalTransition {
	return NewQSignalTransitionFromPointer(cthis)
}

// /usr/include/qt/QtCore/qsignaltransition.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QSignalTransition) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSignalTransition10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qsignaltransition.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSignalTransition(QState *)

/*
Constructs a new signal transition with the given sourceState.
*/
func (*QSignalTransition) NewForInherit(sourceState QState_ITF /*777 QState **/) *QSignalTransition {
	return NewQSignalTransition(sourceState)
}
func NewQSignalTransition(sourceState QState_ITF /*777 QState **/) *QSignalTransition {
	var convArg0 unsafe.Pointer
	if sourceState != nil && sourceState.QState_PTR() != nil {
		convArg0 = sourceState.QState_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSignalTransitionC2EP6QState", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSignalTransitionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSignalTransition")
	return gothis
}

// /usr/include/qt/QtCore/qsignaltransition.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSignalTransition(QState *)

/*
Constructs a new signal transition with the given sourceState.
*/
func (*QSignalTransition) NewForInheritp() *QSignalTransition {
	return NewQSignalTransitionp()
}
func NewQSignalTransitionp() *QSignalTransition {
	// arg: 0, QState *=Pointer, QState=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSignalTransitionC2EP6QState", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSignalTransitionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSignalTransition")
	return gothis
}

// /usr/include/qt/QtCore/qsignaltransition.h:59
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSignalTransition(const QObject *, const char *, QState *)

/*
Constructs a new signal transition with the given sourceState.
*/
func (*QSignalTransition) NewForInherit1(sender QObject_ITF /*777 const QObject **/, signal string, sourceState QState_ITF /*777 QState **/) *QSignalTransition {
	return NewQSignalTransition1(sender, signal, sourceState)
}
func NewQSignalTransition1(sender QObject_ITF /*777 const QObject **/, signal string, sourceState QState_ITF /*777 QState **/) *QSignalTransition {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if sourceState != nil && sourceState.QState_PTR() != nil {
		convArg2 = sourceState.QState_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSignalTransitionC2EPK7QObjectPKcP6QState", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSignalTransitionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSignalTransition")
	return gothis
}

// /usr/include/qt/QtCore/qsignaltransition.h:59
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSignalTransition(const QObject *, const char *, QState *)

/*
Constructs a new signal transition with the given sourceState.
*/
func (*QSignalTransition) NewForInherit1p(sender QObject_ITF /*777 const QObject **/, signal string) *QSignalTransition {
	return NewQSignalTransition1p(sender, signal)
}
func NewQSignalTransition1p(sender QObject_ITF /*777 const QObject **/, signal string) *QSignalTransition {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg1)
	// arg: 2, QState *=Pointer, QState=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSignalTransitionC2EPK7QObjectPKcP6QState", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSignalTransitionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSignalTransition")
	return gothis
}

// /usr/include/qt/QtCore/qsignaltransition.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSignalTransition()

/*

 */
func DeleteQSignalTransition(this *QSignalTransition) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSignalTransitionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qsignaltransition.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * senderObject() const

/*
Returns the sender object associated with this signal transition.

Note: Getter function for property senderObject.

See also setSenderObject().
*/
func (this *QSignalTransition) SenderObject() *QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSignalTransition12senderObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qsignaltransition.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSenderObject(const QObject *)

/*
Sets the sender object associated with this signal transition.

Note: Setter function for property senderObject.

See also senderObject().
*/
func (this *QSignalTransition) SetSenderObject(sender QObject_ITF /*777 const QObject **/) {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSignalTransition15setSenderObjectEPK7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignaltransition.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray signal() const

/*
Returns the signal associated with this signal transition.

Note: Getter function for property signal.

See also setSignal().
*/
func (this *QSignalTransition) Signal() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSignalTransition6signalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qsignaltransition.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSignal(const QByteArray &)

/*
Sets the signal associated with this signal transition.

Note: Setter function for property signal.

See also signal().
*/
func (this *QSignalTransition) SetSignal(signal QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if signal != nil && signal.QByteArray_PTR() != nil {
		convArg0 = signal.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSignalTransition9setSignalERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignaltransition.h:83
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventTest(QEvent *)

/*
Reimplemented from QAbstractTransition::eventTest().

The default implementation returns true if the event is a QStateMachine::SignalEvent object and the event's sender and signal index match this transition, and returns false otherwise.
*/
func (this *QSignalTransition) EventTest(event QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSignalTransition9eventTestEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsignaltransition.h:84
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void onTransition(QEvent *)

/*
Reimplemented from QAbstractTransition::onTransition().
*/
func (this *QSignalTransition) OnTransition(event QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSignalTransition12onTransitionEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignaltransition.h:86
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QAbstractTransition::event().
*/
func (this *QSignalTransition) Event(e QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSignalTransition5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
