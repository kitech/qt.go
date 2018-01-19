//  header block begin
// /usr/include/qt/QtCore/qstatemachine.h
// #include <qstatemachine.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
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
type QStateMachine struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qstatemachine.h:59
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QStateMachine) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStateMachine10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:112
// index:0
// void QStateMachine(class QObject *)
func NewQStateMachine(parent unsafe.Pointer) *QStateMachine {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachineC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QStateMachine{cthis}
}

// /usr/include/qt/QtCore/qstatemachine.h:113
// index:1
// void QStateMachine(class QState::ChildMode, class QObject *)
func NewQStateMachine_1(childMode int, parent unsafe.Pointer) *QStateMachine {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachineC2EN6QState9ChildModeEP7QObject", ffiqt.FFI_TYPE_VOID, cthis, &childMode, parent)
	gopp.ErrPrint(err, rv)
	return &QStateMachine{cthis}
}

// /usr/include/qt/QtCore/qstatemachine.h:114
// index:0
// virtual
// void ~QStateMachine()
func DeleteQStateMachine(*QStateMachine) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachineD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:116
// index:0
// void addState(class QAbstractState *)
func (this *QStateMachine) AddState(state unsafe.Pointer) {
	// 0: (, QAbstractState * state), (state)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine8addStateEP14QAbstractState", ffiqt.FFI_TYPE_VOID, this.cthis, state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:117
// index:0
// void removeState(class QAbstractState *)
func (this *QStateMachine) RemoveState(state unsafe.Pointer) {
	// 0: (, QAbstractState * state), (state)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine11removeStateEP14QAbstractState", ffiqt.FFI_TYPE_VOID, this.cthis, state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:119
// index:0
// QStateMachine::Error error()
func (this *QStateMachine) Error() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStateMachine5errorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:120
// index:0
// QString errorString()
func (this *QStateMachine) ErrorString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStateMachine11errorStringEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:121
// index:0
// void clearError()
func (this *QStateMachine) ClearError() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine10clearErrorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:123
// index:0
// bool isRunning()
func (this *QStateMachine) IsRunning() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStateMachine9isRunningEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:126
// index:0
// bool isAnimated()
func (this *QStateMachine) IsAnimated() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStateMachine10isAnimatedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:127
// index:0
// void setAnimated(_Bool)
func (this *QStateMachine) SetAnimated(enabled bool) {
	// 0: (, bool enabled), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine11setAnimatedEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:129
// index:0
// void addDefaultAnimation(class QAbstractAnimation *)
func (this *QStateMachine) AddDefaultAnimation(animation unsafe.Pointer) {
	// 0: (, QAbstractAnimation * animation), (animation)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine19addDefaultAnimationEP18QAbstractAnimation", ffiqt.FFI_TYPE_VOID, this.cthis, animation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:130
// index:0
// QList<QAbstractAnimation *> defaultAnimations()
func (this *QStateMachine) DefaultAnimations() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStateMachine17defaultAnimationsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:131
// index:0
// void removeDefaultAnimation(class QAbstractAnimation *)
func (this *QStateMachine) RemoveDefaultAnimation(animation unsafe.Pointer) {
	// 0: (, QAbstractAnimation * animation), (animation)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine22removeDefaultAnimationEP18QAbstractAnimation", ffiqt.FFI_TYPE_VOID, this.cthis, animation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:134
// index:0
// QState::RestorePolicy globalRestorePolicy()
func (this *QStateMachine) GlobalRestorePolicy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStateMachine19globalRestorePolicyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:135
// index:0
// void setGlobalRestorePolicy(class QState::RestorePolicy)
func (this *QStateMachine) SetGlobalRestorePolicy(restorePolicy int) {
	// 0: (, QState::RestorePolicy restorePolicy), (&restorePolicy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine22setGlobalRestorePolicyEN6QState13RestorePolicyE", ffiqt.FFI_TYPE_VOID, this.cthis, &restorePolicy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:137
// index:0
// void postEvent(class QEvent *, enum QStateMachine::EventPriority)
func (this *QStateMachine) PostEvent(event unsafe.Pointer, priority int) {
	// 0: (, QEvent * event, QStateMachine::EventPriority priority), (event, &priority)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine9postEventEP6QEventNS_13EventPriorityE", ffiqt.FFI_TYPE_VOID, this.cthis, event, &priority)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:138
// index:0
// int postDelayedEvent(class QEvent *, int)
func (this *QStateMachine) PostDelayedEvent(event unsafe.Pointer, delay int) {
	// 0: (, QEvent * event, int delay), (event, &delay)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine16postDelayedEventEP6QEventi", ffiqt.FFI_TYPE_VOID, this.cthis, event, &delay)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:139
// index:0
// bool cancelDelayedEvent(int)
func (this *QStateMachine) CancelDelayedEvent(id int) {
	// 0: (, int id), (&id)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine18cancelDelayedEventEi", ffiqt.FFI_TYPE_VOID, this.cthis, &id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:141
// index:0
// QSet<QAbstractState *> configuration()
func (this *QStateMachine) Configuration() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStateMachine13configurationEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:144
// index:0
// virtual
// bool eventFilter(class QObject *, class QEvent *)
func (this *QStateMachine) EventFilter(watched unsafe.Pointer, event unsafe.Pointer) {
	// 0: (, QObject * watched, QEvent * event), (watched, event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_VOID, this.cthis, watched, event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:148
// index:0
// void start()
func (this *QStateMachine) Start() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine5startEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:149
// index:0
// void stop()
func (this *QStateMachine) Stop() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine4stopEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:150
// index:0
// void setRunning(_Bool)
func (this *QStateMachine) SetRunning(running bool) {
	// 0: (, bool running), (&running)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine10setRunningEb", ffiqt.FFI_TYPE_VOID, this.cthis, &running)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:155
// index:0
// void runningChanged(_Bool)
func (this *QStateMachine) RunningChanged(running bool) {
	// 0: (, bool running), (&running)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine14runningChangedEb", ffiqt.FFI_TYPE_VOID, this.cthis, &running)
	gopp.ErrPrint(err, rv)
}

//  body block end
