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
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
// void onEntry(class QEvent *)
func (this *QStateMachine) InheritOnEntry(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "onEntry", f)
}

// void onExit(class QEvent *)
func (this *QStateMachine) InheritOnExit(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "onExit", f)
}

// void beginSelectTransitions(class QEvent *)
func (this *QStateMachine) InheritBeginSelectTransitions(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "beginSelectTransitions", f)
}

// void endSelectTransitions(class QEvent *)
func (this *QStateMachine) InheritEndSelectTransitions(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "endSelectTransitions", f)
}

// void beginMicrostep(class QEvent *)
func (this *QStateMachine) InheritBeginMicrostep(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "beginMicrostep", f)
}

// void endMicrostep(class QEvent *)
func (this *QStateMachine) InheritEndMicrostep(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "endMicrostep", f)
}

// bool event(class QEvent *)
func (this *QStateMachine) InheritEvent(f func(e *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QStateMachine) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStateMachine10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qstatemachine.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStateMachine(QObject *)
func NewQStateMachine(parent *QObject /*777 QObject **/) *QStateMachine {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachineC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStateMachineFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qstatemachine.h:113
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QStateMachine(QState::ChildMode, QObject *)
func NewQStateMachine_1(childMode int, parent *QObject /*777 QObject **/) *QStateMachine {
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachineC2EN6QState9ChildModeEP7QObject", qtrt.FFI_TYPE_POINTER, childMode, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQStateMachineFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qstatemachine.h:114
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QStateMachine()
func DeleteQStateMachine(this *QStateMachine) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachineD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qstatemachine.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addState(QAbstractState *)
func (this *QStateMachine) AddState(state *QAbstractState /*777 QAbstractState **/) {
	var convArg0 = state.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine8addStateEP14QAbstractState", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeState(QAbstractState *)
func (this *QStateMachine) RemoveState(state *QAbstractState /*777 QAbstractState **/) {
	var convArg0 = state.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine11removeStateEP14QAbstractState", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:119
// index:0
// Public Visibility=Default Availability=Available
// [4] QStateMachine::Error error()
func (this *QStateMachine) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStateMachine5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:120
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString()
func (this *QStateMachine) ErrorString() *QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStateMachine11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qstatemachine.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearError()
func (this *QStateMachine) ClearError() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine10clearErrorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:123
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRunning()
func (this *QStateMachine) IsRunning() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStateMachine9isRunningEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstatemachine.h:126
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAnimated()
func (this *QStateMachine) IsAnimated() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStateMachine10isAnimatedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstatemachine.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAnimated(_Bool)
func (this *QStateMachine) SetAnimated(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine11setAnimatedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addDefaultAnimation(QAbstractAnimation *)
func (this *QStateMachine) AddDefaultAnimation(animation *QAbstractAnimation /*777 QAbstractAnimation **/) {
	var convArg0 = animation.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine19addDefaultAnimationEP18QAbstractAnimation", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeDefaultAnimation(QAbstractAnimation *)
func (this *QStateMachine) RemoveDefaultAnimation(animation *QAbstractAnimation /*777 QAbstractAnimation **/) {
	var convArg0 = animation.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine22removeDefaultAnimationEP18QAbstractAnimation", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:134
// index:0
// Public Visibility=Default Availability=Available
// [4] QState::RestorePolicy globalRestorePolicy()
func (this *QStateMachine) GlobalRestorePolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStateMachine19globalRestorePolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGlobalRestorePolicy(QState::RestorePolicy)
func (this *QStateMachine) SetGlobalRestorePolicy(restorePolicy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine22setGlobalRestorePolicyEN6QState13RestorePolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), restorePolicy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void postEvent(QEvent *, enum QStateMachine::EventPriority)
func (this *QStateMachine) PostEvent(event *QEvent /*777 QEvent **/, priority int) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine9postEventEP6QEventNS_13EventPriorityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, priority)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:138
// index:0
// Public Visibility=Default Availability=Available
// [4] int postDelayedEvent(QEvent *, int)
func (this *QStateMachine) PostDelayedEvent(event *QEvent /*777 QEvent **/, delay int) int {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine16postDelayedEventEP6QEventi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, delay)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstatemachine.h:139
// index:0
// Public Visibility=Default Availability=Available
// [1] bool cancelDelayedEvent(int)
func (this *QStateMachine) CancelDelayedEvent(id int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine18cancelDelayedEventEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstatemachine.h:144
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)
func (this *QStateMachine) EventFilter(watched *QObject /*777 QObject **/, event *QEvent /*777 QEvent **/) bool {
	var convArg0 = watched.GetCthis()
	var convArg1 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstatemachine.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start()
func (this *QStateMachine) Start() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine5startEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stop()
func (this *QStateMachine) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:150
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRunning(_Bool)
func (this *QStateMachine) SetRunning(running bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine10setRunningEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), running)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void runningChanged(_Bool)
func (this *QStateMachine) RunningChanged(running bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine14runningChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), running)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:159
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void onEntry(QEvent *)
func (this *QStateMachine) OnEntry(event *QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine7onEntryEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:160
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void onExit(QEvent *)
func (this *QStateMachine) OnExit(event *QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine6onExitEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:162
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void beginSelectTransitions(QEvent *)
func (this *QStateMachine) BeginSelectTransitions(event *QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine22beginSelectTransitionsEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:163
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void endSelectTransitions(QEvent *)
func (this *QStateMachine) EndSelectTransitions(event *QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine20endSelectTransitionsEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:165
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void beginMicrostep(QEvent *)
func (this *QStateMachine) BeginMicrostep(event *QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine14beginMicrostepEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:166
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void endMicrostep(QEvent *)
func (this *QStateMachine) EndMicrostep(event *QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine12endMicrostepEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:168
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QStateMachine) Event(e *QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
