package qtcore

// /usr/include/qt/QtCore/qabstractstate.h
// #include <qabstractstate.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 29
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// void onEntry(QEvent *)
func (this *QAbstractState) InheritOnEntry(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "onEntry", f)
}

// void onExit(QEvent *)
func (this *QAbstractState) InheritOnExit(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "onExit", f)
}

// bool event(QEvent *)
func (this *QAbstractState) InheritEvent(f func(e *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

/*

 */
type QAbstractState struct {
	*QObject
}
type QAbstractState_ITF interface {
	QObject_ITF
	QAbstractState_PTR() *QAbstractState
}

func (ptr *QAbstractState) QAbstractState_PTR() *QAbstractState { return ptr }

func (this *QAbstractState) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QAbstractState) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQAbstractStateFromPointer(cthis unsafe.Pointer) *QAbstractState {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QAbstractState{bcthis0}
}
func (*QAbstractState) NewFromPointer(cthis unsafe.Pointer) *QAbstractState {
	return NewQAbstractStateFromPointer(cthis)
}

// /usr/include/qt/QtCore/qabstractstate.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAbstractState) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAbstractState10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qabstractstate.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractState()

/*

 */
func DeleteQAbstractState(this *QAbstractState) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAbstractStateD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qabstractstate.h:60
// index:0
// Public Visibility=Default Availability=Available
// [8] QState * parentState() const

/*
Returns this state's parent state, or 0 if the state has no parent state.
*/
func (this *QAbstractState) ParentState() *QState /*777 QState **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAbstractState11parentStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qabstractstate.h:61
// index:0
// Public Visibility=Default Availability=Available
// [8] QStateMachine * machine() const

/*
Returns the state machine that this state is part of, or 0 if the state is not part of a state machine.
*/
func (this *QAbstractState) Machine() *QStateMachine /*777 QStateMachine **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAbstractState7machineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStateMachineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qabstractstate.h:63
// index:0
// Public Visibility=Default Availability=Available
// [1] bool active() const

/*
Returns whether this state is active.

Note: Getter function for property active.

See also activeChanged(bool), entered(), and exited().
*/
func (this *QAbstractState) Active() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAbstractState6activeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractstate.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activeChanged(bool)

/*
This signal is emitted when the active property is changed with active as argument.

This function was introduced in  Qt 5.4.

Note: Notifier signal for property active.

See also QAbstractState::active, entered(), and exited().
*/
func (this *QAbstractState) ActiveChanged(active bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAbstractState13activeChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), active)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractstate.h:71
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QAbstractState(QState *)

/*
Constructs a new state with the given parent state.
*/
func NewQAbstractState(parent QState_ITF /*777 QState **/) *QAbstractState {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QState_PTR() != nil {
		convArg0 = parent.QState_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAbstractStateC2EP6QState", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractStateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractState")
	return gothis
}

// /usr/include/qt/QtCore/qabstractstate.h:71
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QAbstractState(QState *)

/*
Constructs a new state with the given parent state.
*/
func NewQAbstractState__() *QAbstractState {
	// arg: 0, QState *=Pointer, QState=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAbstractStateC2EP6QState", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractStateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractState")
	return gothis
}

// /usr/include/qt/QtCore/qabstractstate.h:73
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [-2] void onEntry(QEvent *)

/*
This function is called when the state is entered. The given event is what caused the state to be entered. Reimplement this function to perform custom processing when the state is entered.
*/
func (this *QAbstractState) OnEntry(event QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAbstractState7onEntryEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractstate.h:74
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [-2] void onExit(QEvent *)

/*
This function is called when the state is exited. The given event is what caused the state to be exited. Reimplement this function to perform custom processing when the state is exited.
*/
func (this *QAbstractState) OnExit(event QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAbstractState6onExitEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractstate.h:76
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QAbstractState) Event(e QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAbstractState5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
