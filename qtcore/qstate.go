package qtcore

// /usr/include/qt/QtCore/qstate.h
// #include <qstate.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

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
}

//  ext block end

//  body block begin
type QState struct {
	*QAbstractState
}

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
// [8] const QMetaObject * metaObject()
func (this *QState) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QState10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qstate.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QState(QState *)
func NewQState(parent *QState /*777 QState **/) *QState {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStateC2EPS_", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStateFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qstate.h:75
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QState(enum QState::ChildMode, QState *)
func NewQState_1(childMode int, parent *QState /*777 QState **/) *QState {
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStateC2ENS_9ChildModeEPS_", ffiqt.FFI_TYPE_POINTER, childMode, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQStateFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qstate.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QState()
func DeleteQState(*QState) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStateD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractState * errorState()
func (this *QState) ErrorState() *QAbstractState /*777 QAbstractState **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QState10errorStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAbstractStateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qstate.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setErrorState(QAbstractState *)
func (this *QState) SetErrorState(state *QAbstractState /*777 QAbstractState **/) {
	var convArg0 = state.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState13setErrorStateEP14QAbstractState", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addTransition(QAbstractTransition *)
func (this *QState) AddTransition(transition *QAbstractTransition /*777 QAbstractTransition **/) {
	var convArg0 = transition.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState13addTransitionEP19QAbstractTransition", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:82
// index:1
// Public Visibility=Default Availability=Available
// [8] QSignalTransition * addTransition(const QObject *, const char *, QAbstractState *)
func (this *QState) AddTransition_1(sender *QObject /*777 const QObject **/, signal string, target *QAbstractState /*777 QAbstractState **/) *QSignalTransition /*777 QSignalTransition **/ {
	var convArg0 = sender.GetCthis()
	var convArg1 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg1)
	var convArg2 = target.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState13addTransitionEPK7QObjectPKcP14QAbstractState", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSignalTransitionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qstate.h:96
// index:2
// Public Visibility=Default Availability=Available
// [8] QAbstractTransition * addTransition(QAbstractState *)
func (this *QState) AddTransition_2(target *QAbstractState /*777 QAbstractState **/) *QAbstractTransition /*777 QAbstractTransition **/ {
	var convArg0 = target.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState13addTransitionEP14QAbstractState", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAbstractTransitionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qstate.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeTransition(QAbstractTransition *)
func (this *QState) RemoveTransition(transition *QAbstractTransition /*777 QAbstractTransition **/) {
	var convArg0 = transition.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState16removeTransitionEP19QAbstractTransition", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:100
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractState * initialState()
func (this *QState) InitialState() *QAbstractState /*777 QAbstractState **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QState12initialStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAbstractStateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qstate.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInitialState(QAbstractState *)
func (this *QState) SetInitialState(state *QAbstractState /*777 QAbstractState **/) {
	var convArg0 = state.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState15setInitialStateEP14QAbstractState", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:103
// index:0
// Public Visibility=Default Availability=Available
// [4] QState::ChildMode childMode()
func (this *QState) ChildMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QState9childModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qstate.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setChildMode(enum QState::ChildMode)
func (this *QState) SetChildMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState12setChildModeENS_9ChildModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void assignProperty(QObject *, const char *, const QVariant &)
func (this *QState) AssignProperty(object *QObject /*777 QObject **/, name string, value *QVariant) {
	var convArg0 = object.GetCthis()
	var convArg1 = qtrt.CString(name)
	defer qtrt.FreeMem(convArg1)
	var convArg2 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState14assignPropertyEP7QObjectPKcRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:119
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void onEntry(QEvent *)
func (this *QState) OnEntry(event *QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState7onEntryEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:120
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void onExit(QEvent *)
func (this *QState) OnExit(event *QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState6onExitEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:122
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QState) Event(e *QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

type QState__ChildMode = int

const QState__ExclusiveStates QState__ChildMode = 0
const QState__ParallelStates QState__ChildMode = 1

type QState__RestorePolicy = int

const QState__DontRestoreProperties QState__RestorePolicy = 0
const QState__RestoreProperties QState__RestorePolicy = 1

//  body block end
