//  header block begin
// /usr/include/qt/QtCore/qstate.h
// #include <qstate.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
	return this.QAbstractState.GetCthis()
}
func NewQStateFromPointer(cthis unsafe.Pointer) *QState {
	bcthis0 := NewQAbstractStateFromPointer(cthis)
	return &QState{bcthis0}
}

// /usr/include/qt/QtCore/qstate.h:57
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QState) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QState10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstate.h:74
// index:0
// Public
// void QState(class QState *)
func NewQState(parent unsafe.Pointer) *QState {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStateC2EPS_", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQStateFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstate.h:75
// index:1
// Public
// void QState(enum QState::ChildMode, class QState *)
func NewQState_1(childMode int, parent unsafe.Pointer) *QState {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStateC2ENS_9ChildModeEPS_", ffiqt.FFI_TYPE_VOID, cthis, &childMode, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQStateFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstate.h:76
// index:0
// Public virtual
// void ~QState()
func DeleteQState(*QState) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStateD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:78
// index:0
// Public
// QAbstractState * errorState()
func (this *QState) ErrorState() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QState10errorStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstate.h:79
// index:0
// Public
// void setErrorState(class QAbstractState *)
func (this *QState) SetErrorState(state unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState13setErrorStateEP14QAbstractState", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:81
// index:0
// Public
// void addTransition(class QAbstractTransition *)
func (this *QState) AddTransition(transition unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState13addTransitionEP19QAbstractTransition", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), transition)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:82
// index:1
// Public
// QSignalTransition * addTransition(const class QObject *, const char *, class QAbstractState *)
func (this *QState) AddTransition_1(sender unsafe.Pointer, signal string, target unsafe.Pointer) interface{} {
	var convArg1 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState13addTransitionEPK7QObjectPKcP14QAbstractState", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), sender, convArg1, target)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstate.h:96
// index:2
// Public
// QAbstractTransition * addTransition(class QAbstractState *)
func (this *QState) AddTransition_2(target unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState13addTransitionEP14QAbstractState", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), target)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstate.h:97
// index:0
// Public
// void removeTransition(class QAbstractTransition *)
func (this *QState) RemoveTransition(transition unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState16removeTransitionEP19QAbstractTransition", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), transition)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:98
// index:0
// Public
// QList<QAbstractTransition *> transitions()
func (this *QState) Transitions() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QState11transitionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstate.h:100
// index:0
// Public
// QAbstractState * initialState()
func (this *QState) InitialState() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QState12initialStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstate.h:101
// index:0
// Public
// void setInitialState(class QAbstractState *)
func (this *QState) SetInitialState(state unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState15setInitialStateEP14QAbstractState", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:103
// index:0
// Public
// QState::ChildMode childMode()
func (this *QState) ChildMode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QState9childModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstate.h:104
// index:0
// Public
// void setChildMode(enum QState::ChildMode)
func (this *QState) SetChildMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState12setChildModeENS_9ChildModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:107
// index:0
// Public
// void assignProperty(class QObject *, const char *, const class QVariant &)
func (this *QState) AssignProperty(object unsafe.Pointer, name string, value *QVariant) {
	var convArg1 = qtrt.CString(name)
	defer qtrt.FreeMem(convArg1)
	var convArg2 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState14assignPropertyEP7QObjectPKcRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), object, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:119
// index:0
// Protected virtual
// void onEntry(class QEvent *)
func (this *QState) OnEntry(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState7onEntryEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:120
// index:0
// Protected virtual
// void onExit(class QEvent *)
func (this *QState) OnExit(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState6onExitEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:122
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QState) Event(e unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
