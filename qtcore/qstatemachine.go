package qtcore

// /usr/include/qt/QtCore/qstatemachine.h
// #include <qstatemachine.h>
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

// void onEntry(QEvent *)
func (this *QStateMachine) InheritOnEntry(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "onEntry", f)
}

// void onExit(QEvent *)
func (this *QStateMachine) InheritOnExit(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "onExit", f)
}

// void beginSelectTransitions(QEvent *)
func (this *QStateMachine) InheritBeginSelectTransitions(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "beginSelectTransitions", f)
}

// void endSelectTransitions(QEvent *)
func (this *QStateMachine) InheritEndSelectTransitions(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "endSelectTransitions", f)
}

// void beginMicrostep(QEvent *)
func (this *QStateMachine) InheritBeginMicrostep(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "beginMicrostep", f)
}

// void endMicrostep(QEvent *)
func (this *QStateMachine) InheritEndMicrostep(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "endMicrostep", f)
}

// bool event(QEvent *)
func (this *QStateMachine) InheritEvent(f func(e *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

/*

 */
type QStateMachine struct {
	*QState
}
type QStateMachine_ITF interface {
	QState_ITF
	QStateMachine_PTR() *QStateMachine
}

func (ptr *QStateMachine) QStateMachine_PTR() *QStateMachine { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QStateMachine) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStateMachine10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qstatemachine.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStateMachine(QObject *)

/*
Constructs a new state machine with the given parent.
*/
func NewQStateMachine(parent QObject_ITF /*777 QObject **/) *QStateMachine {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachineC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStateMachineFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QStateMachine")
	return gothis
}

// /usr/include/qt/QtCore/qstatemachine.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStateMachine(QObject *)

/*
Constructs a new state machine with the given parent.
*/
func NewQStateMachine__() *QStateMachine {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachineC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStateMachineFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QStateMachine")
	return gothis
}

// /usr/include/qt/QtCore/qstatemachine.h:113
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QStateMachine(QState::ChildMode, QObject *)

/*
Constructs a new state machine with the given parent.
*/
func NewQStateMachine_1(childMode int, parent QObject_ITF /*777 QObject **/) *QStateMachine {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachineC2EN6QState9ChildModeEP7QObject", qtrt.FFI_TYPE_POINTER, childMode, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStateMachineFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QStateMachine")
	return gothis
}

// /usr/include/qt/QtCore/qstatemachine.h:113
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QStateMachine(QState::ChildMode, QObject *)

/*
Constructs a new state machine with the given parent.
*/
func NewQStateMachine_1_(childMode int) *QStateMachine {
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachineC2EN6QState9ChildModeEP7QObject", qtrt.FFI_TYPE_POINTER, childMode, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStateMachineFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QStateMachine")
	return gothis
}

// /usr/include/qt/QtCore/qstatemachine.h:114
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QStateMachine()

/*

 */
func DeleteQStateMachine(this *QStateMachine) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachineD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qstatemachine.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addState(QAbstractState *)

/*
Adds the given state to this state machine. The state becomes a top-level state.

If the state is already in a different machine, it will first be removed from its old machine, and then added to this machine.

See also removeState() and setInitialState().
*/
func (this *QStateMachine) AddState(state QAbstractState_ITF /*777 QAbstractState **/) {
	var convArg0 unsafe.Pointer
	if state != nil && state.QAbstractState_PTR() != nil {
		convArg0 = state.QAbstractState_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine8addStateEP14QAbstractState", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeState(QAbstractState *)

/*
Removes the given state from this state machine. The state machine releases ownership of the state.

See also addState().
*/
func (this *QStateMachine) RemoveState(state QAbstractState_ITF /*777 QAbstractState **/) {
	var convArg0 unsafe.Pointer
	if state != nil && state.QAbstractState_PTR() != nil {
		convArg0 = state.QAbstractState_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine11removeStateEP14QAbstractState", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:119
// index:0
// Public Visibility=Default Availability=Available
// [4] QStateMachine::Error error() const

/*
Returns the error code of the last error that occurred in the state machine.
*/
func (this *QStateMachine) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStateMachine5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:120
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const

/*
Returns the error string of the last error that occurred in the state machine.

Note: Getter function for property errorString.
*/
func (this *QStateMachine) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStateMachine11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstatemachine.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearError()

/*
Clears the error string and error code of the state machine.
*/
func (this *QStateMachine) ClearError() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine10clearErrorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:123
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRunning() const

/*

 */
func (this *QStateMachine) IsRunning() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStateMachine9isRunningEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstatemachine.h:126
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAnimated() const

/*
Returns whether animations are enabled for this state machine.

Note: Getter function for property animated.
*/
func (this *QStateMachine) IsAnimated() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStateMachine10isAnimatedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstatemachine.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAnimated(bool)

/*
Sets whether animations are enabled for this state machine.

Note: Setter function for property animated.

See also isAnimated().
*/
func (this *QStateMachine) SetAnimated(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine11setAnimatedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addDefaultAnimation(QAbstractAnimation *)

/*
Adds a default animation to be considered for any transition.
*/
func (this *QStateMachine) AddDefaultAnimation(animation QAbstractAnimation_ITF /*777 QAbstractAnimation **/) {
	var convArg0 unsafe.Pointer
	if animation != nil && animation.QAbstractAnimation_PTR() != nil {
		convArg0 = animation.QAbstractAnimation_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine19addDefaultAnimationEP18QAbstractAnimation", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeDefaultAnimation(QAbstractAnimation *)

/*
Removes animation from the list of default animations.
*/
func (this *QStateMachine) RemoveDefaultAnimation(animation QAbstractAnimation_ITF /*777 QAbstractAnimation **/) {
	var convArg0 unsafe.Pointer
	if animation != nil && animation.QAbstractAnimation_PTR() != nil {
		convArg0 = animation.QAbstractAnimation_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine22removeDefaultAnimationEP18QAbstractAnimation", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:134
// index:0
// Public Visibility=Default Availability=Available
// [4] QState::RestorePolicy globalRestorePolicy() const

/*
Returns the restore policy of the state machine.

Note: Getter function for property globalRestorePolicy.

See also setGlobalRestorePolicy().
*/
func (this *QStateMachine) GlobalRestorePolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStateMachine19globalRestorePolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGlobalRestorePolicy(QState::RestorePolicy)

/*
Sets the restore policy of the state machine to restorePolicy. The default restore policy is QState::DontRestoreProperties.

Note: Setter function for property globalRestorePolicy.

See also globalRestorePolicy().
*/
func (this *QStateMachine) SetGlobalRestorePolicy(restorePolicy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine22setGlobalRestorePolicyEN6QState13RestorePolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), restorePolicy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void postEvent(QEvent *, QStateMachine::EventPriority)

/*
Posts the given event of the given priority for processing by this state machine.

This function returns immediately. The event is added to the state machine's event queue. Events are processed in the order posted. The state machine takes ownership of the event and deletes it once it has been processed.

You can only post events when the state machine is running or when it is starting up.

Note: This function is thread-safe.

See also postDelayedEvent().
*/
func (this *QStateMachine) PostEvent(event QEvent_ITF /*777 QEvent **/, priority int) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine9postEventEP6QEventNS_13EventPriorityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, priority)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void postEvent(QEvent *, QStateMachine::EventPriority)

/*
Posts the given event of the given priority for processing by this state machine.

This function returns immediately. The event is added to the state machine's event queue. Events are processed in the order posted. The state machine takes ownership of the event and deletes it once it has been processed.

You can only post events when the state machine is running or when it is starting up.

Note: This function is thread-safe.

See also postDelayedEvent().
*/
func (this *QStateMachine) PostEvent__(event QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	// arg: 1, QStateMachine::EventPriority=Enum, QStateMachine::EventPriority=Enum, , Invalid
	priority := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine9postEventEP6QEventNS_13EventPriorityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, priority)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:138
// index:0
// Public Visibility=Default Availability=Available
// [4] int postDelayedEvent(QEvent *, int)

/*
Posts the given event for processing by this state machine, with the given delay in milliseconds. Returns an identifier associated with the delayed event, or -1 if the event could not be posted.

This function returns immediately. When the delay has expired, the event will be added to the state machine's event queue for processing. The state machine takes ownership of the event and deletes it once it has been processed.

You can only post events when the state machine is running.

Note: This function is thread-safe.

See also cancelDelayedEvent() and postEvent().
*/
func (this *QStateMachine) PostDelayedEvent(event QEvent_ITF /*777 QEvent **/, delay int) int {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine16postDelayedEventEP6QEventi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, delay)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstatemachine.h:139
// index:0
// Public Visibility=Default Availability=Available
// [1] bool cancelDelayedEvent(int)

/*
Cancels the delayed event identified by the given id. The id should be a value returned by a call to postDelayedEvent(). Returns true if the event was successfully cancelled, otherwise returns false.

Note: This function is thread-safe.

See also postDelayedEvent().
*/
func (this *QStateMachine) CancelDelayedEvent(id int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine18cancelDelayedEventEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstatemachine.h:144
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)

/*
Reimplemented from QObject::eventFilter().
*/
func (this *QStateMachine) EventFilter(watched QObject_ITF /*777 QObject **/, event QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if watched != nil && watched.QObject_PTR() != nil {
		convArg0 = watched.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg1 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstatemachine.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start()

/*
Starts this state machine. The machine will reset its configuration and transition to the initial state. When a final top-level state (QFinalState) is entered, the machine will emit the finished() signal.

Note: A state machine will not run without a running event loop, such as the main application event loop started with QCoreApplication::exec() or QApplication::exec().

See also started(), finished(), stop(), initialState(), and setRunning().
*/
func (this *QStateMachine) Start() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine5startEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stop()

/*
Stops this state machine. The state machine will stop processing events and then emit the stopped() signal.

See also stopped(), start(), and setRunning().
*/
func (this *QStateMachine) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:150
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRunning(bool)

/*

 */
func (this *QStateMachine) SetRunning(running bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine10setRunningEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), running)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void runningChanged(bool)

/*
This signal is emitted when the running property is changed with running as argument.

This function was introduced in  Qt 5.4.

Note: Notifier signal for property running.

See also QStateMachine::running.
*/
func (this *QStateMachine) RunningChanged(running bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine14runningChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), running)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:159
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void onEntry(QEvent *)

/*
Reimplemented from QAbstractState::onEntry().

This function will call start() to start the state machine.
*/
func (this *QStateMachine) OnEntry(event QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine7onEntryEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:160
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void onExit(QEvent *)

/*
Reimplemented from QAbstractState::onExit().

This function will call stop() to stop the state machine and subsequently emit the stopped() signal.
*/
func (this *QStateMachine) OnExit(event QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine6onExitEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:162
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void beginSelectTransitions(QEvent *)

/*

 */
func (this *QStateMachine) BeginSelectTransitions(event QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine22beginSelectTransitionsEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:163
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void endSelectTransitions(QEvent *)

/*

 */
func (this *QStateMachine) EndSelectTransitions(event QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine20endSelectTransitionsEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:165
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void beginMicrostep(QEvent *)

/*

 */
func (this *QStateMachine) BeginMicrostep(event QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine14beginMicrostepEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:166
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void endMicrostep(QEvent *)

/*

 */
func (this *QStateMachine) EndMicrostep(event QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine12endMicrostepEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:168
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QStateMachine) Event(e QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*
This enum type specifies the priority of an event posted to the state machine using postEvent().

Events of high priority are processed before events of normal priority.


*/
type QStateMachine__EventPriority = int

// The event has normal priority.
const QStateMachine__NormalPriority QStateMachine__EventPriority = 0

// The event has high priority.
const QStateMachine__HighPriority QStateMachine__EventPriority = 1

/*
This enum type defines errors that can occur in the state machine at run time. When the state machine encounters an unrecoverable error at run time, it will set the error code returned by error(), the error message returned by errorString(), and enter an error state based on the context of the error.



See also setErrorState().

*/
type QStateMachine__Error = int

// No error has occurred.
const QStateMachine__NoError QStateMachine__Error = 0

// The machine has entered a QState with children which does not have an initial state set. The context of this error is the state which is missing an initial state.
const QStateMachine__NoInitialStateError QStateMachine__Error = 1

// The machine has entered a QHistoryState which does not have a default state set. The context of this error is the QHistoryState which is missing a default state.
const QStateMachine__NoDefaultStateInHistoryStateError QStateMachine__Error = 2

// The machine has selected a transition whose source and targets are not part of the same tree of states, and thus are not part of the same state machine. Commonly, this could mean that one of the states has not been given any parent or added to any machine. The context of this error is the source state of the transition.
const QStateMachine__NoCommonAncestorForTransitionError QStateMachine__Error = 3

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
