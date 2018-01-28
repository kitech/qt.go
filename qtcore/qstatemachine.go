package qtcore

// /usr/include/qt/QtCore/qstatemachine.h
// #include <qstatemachine.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
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
type QStateMachine struct {
	*QState
}

func (this *QStateMachine) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QState.GetCthis()
	}
}
func (this *QStateMachine) SetCthis(cthis unsafe.Pointer) {
	this.QState = NewQStateFromPointer(cthis)
}
func NewQStateMachineFromPointer(cthis unsafe.Pointer) *QStateMachine {
	bcthis0 := NewQStateFromPointer(cthis)
	return &QStateMachine{bcthis0}
}
func (*QStateMachine) NewFromPointer(cthis unsafe.Pointer) *QStateMachine {
	return NewQStateMachineFromPointer(cthis)
}

// /usr/include/qt/QtCore/qstatemachine.h:59
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QStateMachine) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStateMachine10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qstatemachine.h:112
// index:0
// Public
// void QStateMachine(QObject *)
func NewQStateMachine(parent *QObject /*777 QObject **/) *QStateMachine {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachineC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStateMachineFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstatemachine.h:113
// index:1
// Public
// void QStateMachine(QState::ChildMode, QObject *)
func NewQStateMachine_1(childMode int, parent *QObject /*777 QObject **/) *QStateMachine {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachineC2EN6QState9ChildModeEP7QObject", ffiqt.FFI_TYPE_VOID, cthis, childMode, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQStateMachineFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstatemachine.h:114
// index:0
// Public virtual
// void ~QStateMachine()
func DeleteQStateMachine(*QStateMachine) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachineD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:116
// index:0
// Public
// void addState(QAbstractState *)
func (this *QStateMachine) AddState(state *QAbstractState /*777 QAbstractState **/) {
	var convArg0 = state.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine8addStateEP14QAbstractState", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:117
// index:0
// Public
// void removeState(QAbstractState *)
func (this *QStateMachine) RemoveState(state *QAbstractState /*777 QAbstractState **/) {
	var convArg0 = state.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine11removeStateEP14QAbstractState", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:119
// index:0
// Public
// QStateMachine::Error error()
func (this *QStateMachine) Error() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStateMachine5errorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:120
// index:0
// Public
// QString errorString()
func (this *QStateMachine) ErrorString() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStateMachine11errorStringEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qstatemachine.h:121
// index:0
// Public
// void clearError()
func (this *QStateMachine) ClearError() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine10clearErrorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:123
// index:0
// Public
// bool isRunning()
func (this *QStateMachine) IsRunning() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStateMachine9isRunningEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstatemachine.h:126
// index:0
// Public
// bool isAnimated()
func (this *QStateMachine) IsAnimated() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStateMachine10isAnimatedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstatemachine.h:127
// index:0
// Public
// void setAnimated(_Bool)
func (this *QStateMachine) SetAnimated(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine11setAnimatedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:129
// index:0
// Public
// void addDefaultAnimation(QAbstractAnimation *)
func (this *QStateMachine) AddDefaultAnimation(animation *QAbstractAnimation /*777 QAbstractAnimation **/) {
	var convArg0 = animation.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine19addDefaultAnimationEP18QAbstractAnimation", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:131
// index:0
// Public
// void removeDefaultAnimation(QAbstractAnimation *)
func (this *QStateMachine) RemoveDefaultAnimation(animation *QAbstractAnimation /*777 QAbstractAnimation **/) {
	var convArg0 = animation.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine22removeDefaultAnimationEP18QAbstractAnimation", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:134
// index:0
// Public
// QState::RestorePolicy globalRestorePolicy()
func (this *QStateMachine) GlobalRestorePolicy() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStateMachine19globalRestorePolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:135
// index:0
// Public
// void setGlobalRestorePolicy(QState::RestorePolicy)
func (this *QStateMachine) SetGlobalRestorePolicy(restorePolicy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine22setGlobalRestorePolicyEN6QState13RestorePolicyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), restorePolicy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:137
// index:0
// Public
// void postEvent(QEvent *, enum QStateMachine::EventPriority)
func (this *QStateMachine) PostEvent(event *QEvent /*777 QEvent **/, priority int) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine9postEventEP6QEventNS_13EventPriorityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, priority)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:138
// index:0
// Public
// int postDelayedEvent(QEvent *, int)
func (this *QStateMachine) PostDelayedEvent(event *QEvent /*777 QEvent **/, delay int) int {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine16postDelayedEventEP6QEventi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, delay)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstatemachine.h:139
// index:0
// Public
// bool cancelDelayedEvent(int)
func (this *QStateMachine) CancelDelayedEvent(id int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine18cancelDelayedEventEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), id)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstatemachine.h:144
// index:0
// Public virtual
// bool eventFilter(QObject *, QEvent *)
func (this *QStateMachine) EventFilter(watched *QObject /*777 QObject **/, event *QEvent /*777 QEvent **/) bool {
	var convArg0 = watched.GetCthis()
	var convArg1 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstatemachine.h:148
// index:0
// Public
// void start()
func (this *QStateMachine) Start() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine5startEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:149
// index:0
// Public
// void stop()
func (this *QStateMachine) Stop() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine4stopEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:150
// index:0
// Public
// void setRunning(_Bool)
func (this *QStateMachine) SetRunning(running bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine10setRunningEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), running)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:155
// index:0
// Public
// void runningChanged(_Bool)
func (this *QStateMachine) RunningChanged(running bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine14runningChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), running)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:159
// index:0
// Protected virtual
// void onEntry(QEvent *)
func (this *QStateMachine) OnEntry(event *QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine7onEntryEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:160
// index:0
// Protected virtual
// void onExit(QEvent *)
func (this *QStateMachine) OnExit(event *QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine6onExitEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:162
// index:0
// Protected virtual
// void beginSelectTransitions(QEvent *)
func (this *QStateMachine) BeginSelectTransitions(event *QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine22beginSelectTransitionsEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:163
// index:0
// Protected virtual
// void endSelectTransitions(QEvent *)
func (this *QStateMachine) EndSelectTransitions(event *QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine20endSelectTransitionsEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:165
// index:0
// Protected virtual
// void beginMicrostep(QEvent *)
func (this *QStateMachine) BeginMicrostep(event *QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine14beginMicrostepEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:166
// index:0
// Protected virtual
// void endMicrostep(QEvent *)
func (this *QStateMachine) EndMicrostep(event *QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine12endMicrostepEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:168
// index:0
// Protected virtual
// bool event(QEvent *)
func (this *QStateMachine) Event(e *QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStateMachine5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

type QStateMachine__EventPriority = int

const QStateMachine__NormalPriority QStateMachine__EventPriority = 0
const QStateMachine__HighPriority QStateMachine__EventPriority = 1

type QStateMachine__Error = int

const QStateMachine__NoError QStateMachine__Error = 0
const QStateMachine__NoInitialStateError QStateMachine__Error = 1
const QStateMachine__NoDefaultStateInHistoryStateError QStateMachine__Error = 2
const QStateMachine__NoCommonAncestorForTransitionError QStateMachine__Error = 3

//  body block end
