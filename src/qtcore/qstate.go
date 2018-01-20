//  header block begin
// /usr/include/qt/QtCore/qstate.h
// #include <qstate.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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

// /usr/include/qt/QtCore/qstate.h:57
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QState) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QState10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:74
// index:0
// void QState(class QState *)
func NewQState(parent unsafe.Pointer) *QState {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStateC2EPS_", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQStateFromPointer(cthis)
	return gothis
}
func NewQStateFromPointer(cthis unsafe.Pointer) *QState {
	bcthis0 := NewQAbstractStateFromPointer(cthis)
	return &QState{bcthis0}
}

// /usr/include/qt/QtCore/qstate.h:75
// index:1
// void QState(enum QState::ChildMode, class QState *)
func NewQState_1(childMode int, parent unsafe.Pointer) *QState {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStateC2ENS_9ChildModeEPS_", ffiqt.FFI_TYPE_VOID, cthis, &childMode, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQStateFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstate.h:125
// index:2
// void QState(class QStatePrivate &, class QState *)
func NewQState_2(dd unsafe.Pointer, parent unsafe.Pointer) *QState {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStateC2ER13QStatePrivatePS_", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQStateFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstate.h:76
// index:0
// virtual
// void ~QState()
func DeleteQState(*QState) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStateD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:78
// index:0
// QAbstractState * errorState()
func (this *QState) ErrorState() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QState10errorStateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:79
// index:0
// void setErrorState(class QAbstractState *)
func (this *QState) SetErrorState(state unsafe.Pointer) {
	// 0: (, state QAbstractState *), (state)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState13setErrorStateEP14QAbstractState", ffiqt.FFI_TYPE_VOID, this.GetCthis(), state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:81
// index:0
// void addTransition(class QAbstractTransition *)
func (this *QState) AddTransition(transition unsafe.Pointer) {
	// 0: (, transition QAbstractTransition *), (transition)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState13addTransitionEP19QAbstractTransition", ffiqt.FFI_TYPE_VOID, this.GetCthis(), transition)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:82
// index:1
// QSignalTransition * addTransition(const class QObject *, const char *, class QAbstractState *)
func (this *QState) AddTransition_1(sender unsafe.Pointer, signal unsafe.Pointer, target unsafe.Pointer) {
	// 1: (, sender const QObject *, signal const char *, target QAbstractState *), (sender, signal, target)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState13addTransitionEPK7QObjectPKcP14QAbstractState", ffiqt.FFI_TYPE_VOID, this.GetCthis(), sender, signal, target)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:96
// index:2
// QAbstractTransition * addTransition(class QAbstractState *)
func (this *QState) AddTransition_2(target unsafe.Pointer) {
	// 2: (, target QAbstractState *), (target)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState13addTransitionEP14QAbstractState", ffiqt.FFI_TYPE_VOID, this.GetCthis(), target)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:97
// index:0
// void removeTransition(class QAbstractTransition *)
func (this *QState) RemoveTransition(transition unsafe.Pointer) {
	// 0: (, transition QAbstractTransition *), (transition)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState16removeTransitionEP19QAbstractTransition", ffiqt.FFI_TYPE_VOID, this.GetCthis(), transition)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:98
// index:0
// QList<QAbstractTransition *> transitions()
func (this *QState) Transitions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QState11transitionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:100
// index:0
// QAbstractState * initialState()
func (this *QState) InitialState() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QState12initialStateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:101
// index:0
// void setInitialState(class QAbstractState *)
func (this *QState) SetInitialState(state unsafe.Pointer) {
	// 0: (, state QAbstractState *), (state)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState15setInitialStateEP14QAbstractState", ffiqt.FFI_TYPE_VOID, this.GetCthis(), state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:103
// index:0
// QState::ChildMode childMode()
func (this *QState) ChildMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QState9childModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:104
// index:0
// void setChildMode(enum QState::ChildMode)
func (this *QState) SetChildMode(mode int) {
	// 0: (, mode QState::ChildMode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState12setChildModeENS_9ChildModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:107
// index:0
// void assignProperty(class QObject *, const char *, const class QVariant &)
func (this *QState) AssignProperty(object unsafe.Pointer, name unsafe.Pointer, value unsafe.Pointer) {
	// 0: (, object QObject *, name const char *, value const QVariant &), (object, name, value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState14assignPropertyEP7QObjectPKcRK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), object, name, value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:119
// index:0
// virtual
// void onEntry(class QEvent *)
func (this *QState) OnEntry(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState7onEntryEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:120
// index:0
// virtual
// void onExit(class QEvent *)
func (this *QState) OnExit(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState6onExitEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstate.h:122
// index:0
// virtual
// bool event(class QEvent *)
func (this *QState) Event(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QState5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

//  body block end
