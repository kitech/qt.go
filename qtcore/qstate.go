package qtcore

// /usr/include/qt/QtCore/qstate.h
// #include <qstate.h>
// #include <QtCore>

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
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// void onEntry(class QEvent *)
func (this *QState) InheritOnEntry(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "onEntry", f)
}

// void onExit(class QEvent *)
func (this *QState) InheritOnExit(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "onExit", f)
}

// bool event(class QEvent *)
func (this *QState) InheritEvent(f func(e *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

type QState struct {
	*QAbstractState
}
type QState_ITF interface {
	QAbstractState_ITF
	QState_PTR() *QState
}

func (ptr *QState) QState_PTR() *QState { return ptr }

func (this *QState) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractState.GetCthis()
	}
}
func (this *QState) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractState = NewQAbstractStateFromPointer(cthis)
}
func NewQStateFromPointer(cthis unsafe.Pointer) *QState {
	bcthis0 := NewQAbstractStateFromPointer(cthis)
	return &QState{bcthis0}
}
func (*QState) NewFromPointer(cthis unsafe.Pointer) *QState {
	return NewQStateFromPointer(cthis)
}

// /usr/include/qt/QtCore/qstate.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QState) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QState10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qstate.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QState(QState *)
func NewQState(parent QState_ITF /*777 QState **/) *QState {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QState_PTR() != nil {
		convArg0 = parent.QState_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QStateC2EPS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QState")
	return gothis
}

// /usr/include/qt/QtCore/qstate.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QState(QState *)
func NewQState__() *QState {
	// arg: 0, QState *=Pointer, QState=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QStateC2EPS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QState")
	return gothis
}

// /usr/include/qt/QtCore/qstate.h:75
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QState(enum QState::ChildMode, QState *)
func NewQState_1(childMode int, parent QState_ITF /*777 QState **/) *QState {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QState_PTR() != nil {
		convArg1 = parent.QState_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QStateC2ENS_9ChildModeEPS_", qtrt.FFI_TYPE_POINTER, childMode, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QState")
	return gothis
}

// /usr/include/qt/QtCore/qstate.h:75
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QState(enum QState::ChildMode, QState *)
func NewQState_1_(childMode int) *QState {
	// arg: 1, QState *=Pointer, QState=Record,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QStateC2ENS_9ChildModeEPS_", qtrt.FFI_TYPE_POINTER, childMode, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QState")
	return gothis
}

// /usr/include/qt/QtCore/qstate.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QState()
func DeleteQState(this *QState) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QStateD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qstate.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractState * errorState() const
func (this *QState) ErrorState() *QAbstractState /*777 QAbstractState **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QState10errorStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractStateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qstate.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setErrorState(QAbstractState *)
func (this *QState) SetErrorState(state QAbstractState_ITF /*777 QAbstractState **/) {
	var convArg0 unsafe.Pointer
	if state != nil && state.QAbstractState_PTR() != nil {
		convArg0 = state.QAbstractState_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QState13setErrorStateEP14QAbstractState", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addTransition(QAbstractTransition *)
func (this *QState) AddTransition(transition QAbstractTransition_ITF /*777 QAbstractTransition **/) {
	var convArg0 unsafe.Pointer
	if transition != nil && transition.QAbstractTransition_PTR() != nil {
		convArg0 = transition.QAbstractTransition_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QState13addTransitionEP19QAbstractTransition", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:82
// index:1
// Public Visibility=Default Availability=Available
// [8] QSignalTransition * addTransition(const QObject *, const char *, QAbstractState *)
func (this *QState) AddTransition_1(sender QObject_ITF /*777 const QObject **/, signal string, target QAbstractState_ITF /*777 QAbstractState **/) *QSignalTransition /*777 QSignalTransition **/ {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if target != nil && target.QAbstractState_PTR() != nil {
		convArg2 = target.QAbstractState_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QState13addTransitionEPK7QObjectPKcP14QAbstractState", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSignalTransitionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qstate.h:96
// index:2
// Public Visibility=Default Availability=Available
// [8] QAbstractTransition * addTransition(QAbstractState *)
func (this *QState) AddTransition_2(target QAbstractState_ITF /*777 QAbstractState **/) *QAbstractTransition /*777 QAbstractTransition **/ {
	var convArg0 unsafe.Pointer
	if target != nil && target.QAbstractState_PTR() != nil {
		convArg0 = target.QAbstractState_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QState13addTransitionEP14QAbstractState", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractTransitionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qstate.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeTransition(QAbstractTransition *)
func (this *QState) RemoveTransition(transition QAbstractTransition_ITF /*777 QAbstractTransition **/) {
	var convArg0 unsafe.Pointer
	if transition != nil && transition.QAbstractTransition_PTR() != nil {
		convArg0 = transition.QAbstractTransition_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QState16removeTransitionEP19QAbstractTransition", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:100
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractState * initialState() const
func (this *QState) InitialState() *QAbstractState /*777 QAbstractState **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QState12initialStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractStateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qstate.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInitialState(QAbstractState *)
func (this *QState) SetInitialState(state QAbstractState_ITF /*777 QAbstractState **/) {
	var convArg0 unsafe.Pointer
	if state != nil && state.QAbstractState_PTR() != nil {
		convArg0 = state.QAbstractState_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QState15setInitialStateEP14QAbstractState", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:103
// index:0
// Public Visibility=Default Availability=Available
// [4] QState::ChildMode childMode() const
func (this *QState) ChildMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QState9childModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qstate.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setChildMode(enum QState::ChildMode)
func (this *QState) SetChildMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QState12setChildModeENS_9ChildModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void assignProperty(QObject *, const char *, const QVariant &)
func (this *QState) AssignProperty(object QObject_ITF /*777 QObject **/, name string, value QVariant_ITF) {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(name)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg2 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QState14assignPropertyEP7QObjectPKcRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:119
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void onEntry(QEvent *)
func (this *QState) OnEntry(event QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QState7onEntryEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:120
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void onExit(QEvent *)
func (this *QState) OnExit(event QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QState6onExitEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:122
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QState) Event(e QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QState5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

type QState__ChildMode = int

const QState__ExclusiveStates QState__ChildMode = 0
const QState__ParallelStates QState__ChildMode = 1

type QState__RestorePolicy = int

const QState__DontRestoreProperties QState__RestorePolicy = 0
const QState__RestoreProperties QState__RestorePolicy = 1

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
