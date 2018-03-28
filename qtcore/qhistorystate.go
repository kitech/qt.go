package qtcore

// /usr/include/qt/QtCore/qhistorystate.h
// #include <qhistorystate.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 33
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
func (this *QHistoryState) InheritOnEntry(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "onEntry", f)
}

// void onExit(QEvent *)
func (this *QHistoryState) InheritOnExit(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "onExit", f)
}

// bool event(QEvent *)
func (this *QHistoryState) InheritEvent(f func(e *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

/*

 */
type QHistoryState struct {
	*QAbstractState
}
type QHistoryState_ITF interface {
	QAbstractState_ITF
	QHistoryState_PTR() *QHistoryState
}

func (ptr *QHistoryState) QHistoryState_PTR() *QHistoryState { return ptr }

func (this *QHistoryState) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractState.GetCthis()
	}
}
func (this *QHistoryState) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractState = NewQAbstractStateFromPointer(cthis)
}
func NewQHistoryStateFromPointer(cthis unsafe.Pointer) *QHistoryState {
	bcthis0 := NewQAbstractStateFromPointer(cthis)
	return &QHistoryState{bcthis0}
}
func (*QHistoryState) NewFromPointer(cthis unsafe.Pointer) *QHistoryState {
	return NewQHistoryStateFromPointer(cthis)
}

// /usr/include/qt/QtCore/qhistorystate.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QHistoryState) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QHistoryState10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qhistorystate.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QHistoryState(QState *)

/*
Constructs a new shallow history state with the given parent state.
*/
func NewQHistoryState(parent QState_ITF /*777 QState **/) *QHistoryState {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QState_PTR() != nil {
		convArg0 = parent.QState_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QHistoryStateC2EP6QState", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHistoryStateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QHistoryState")
	return gothis
}

// /usr/include/qt/QtCore/qhistorystate.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QHistoryState(QState *)

/*
Constructs a new shallow history state with the given parent state.
*/
func NewQHistoryState__() *QHistoryState {
	// arg: 0, QState *=Pointer, QState=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QHistoryStateC2EP6QState", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHistoryStateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QHistoryState")
	return gothis
}

// /usr/include/qt/QtCore/qhistorystate.h:65
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QHistoryState(QHistoryState::HistoryType, QState *)

/*
Constructs a new shallow history state with the given parent state.
*/
func NewQHistoryState_1(type_ int, parent QState_ITF /*777 QState **/) *QHistoryState {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QState_PTR() != nil {
		convArg1 = parent.QState_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QHistoryStateC2ENS_11HistoryTypeEP6QState", qtrt.FFI_TYPE_POINTER, type_, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHistoryStateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QHistoryState")
	return gothis
}

// /usr/include/qt/QtCore/qhistorystate.h:65
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QHistoryState(QHistoryState::HistoryType, QState *)

/*
Constructs a new shallow history state with the given parent state.
*/
func NewQHistoryState_1_(type_ int) *QHistoryState {
	// arg: 1, QState *=Pointer, QState=Record,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QHistoryStateC2ENS_11HistoryTypeEP6QState", qtrt.FFI_TYPE_POINTER, type_, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHistoryStateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QHistoryState")
	return gothis
}

// /usr/include/qt/QtCore/qhistorystate.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QHistoryState()

/*

 */
func DeleteQHistoryState(this *QHistoryState) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QHistoryStateD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qhistorystate.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractTransition * defaultTransition() const

/*
Returns this history state's default transition. The default transition is taken when the history state has never been entered before. The target states of the default transition therefore make up the default state.

This function was introduced in  Qt 5.6.

Note: Getter function for property defaultTransition.

See also setDefaultTransition().
*/
func (this *QHistoryState) DefaultTransition() *QAbstractTransition /*777 QAbstractTransition **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QHistoryState17defaultTransitionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractTransitionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qhistorystate.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultTransition(QAbstractTransition *)

/*
Sets this history state's default transition to be the given transition. This will set the source state of the transition to the history state.

Note that the eventTest method of the transition will never be called.

This function was introduced in  Qt 5.6.

Note: Setter function for property defaultTransition.

See also defaultTransition().
*/
func (this *QHistoryState) SetDefaultTransition(transition QAbstractTransition_ITF /*777 QAbstractTransition **/) {
	var convArg0 unsafe.Pointer
	if transition != nil && transition.QAbstractTransition_PTR() != nil {
		convArg0 = transition.QAbstractTransition_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QHistoryState20setDefaultTransitionEP19QAbstractTransition", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractState * defaultState() const

/*
Returns this history state's default state. The default state indicates the state to transition to if the parent state has never been entered before.

Note: Getter function for property defaultState.

See also setDefaultState().
*/
func (this *QHistoryState) DefaultState() *QAbstractState /*777 QAbstractState **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QHistoryState12defaultStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractStateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qhistorystate.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultState(QAbstractState *)

/*
Sets this history state's default state to be the given state. state must be a sibling of this history state.

Note that this function does not set state as the initial state of its parent.

Note: Setter function for property defaultState.

See also defaultState().
*/
func (this *QHistoryState) SetDefaultState(state QAbstractState_ITF /*777 QAbstractState **/) {
	var convArg0 unsafe.Pointer
	if state != nil && state.QAbstractState_PTR() != nil {
		convArg0 = state.QAbstractState_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QHistoryState15setDefaultStateEP14QAbstractState", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:74
// index:0
// Public Visibility=Default Availability=Available
// [4] QHistoryState::HistoryType historyType() const

/*
Returns the type of history that this history state records.

Note: Getter function for property historyType.

See also setHistoryType().
*/
func (this *QHistoryState) HistoryType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QHistoryState11historyTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHistoryType(QHistoryState::HistoryType)

/*
Sets the type of history that this history state records.

Note: Setter function for property historyType.

See also historyType().
*/
func (this *QHistoryState) SetHistoryType(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QHistoryState14setHistoryTypeENS_11HistoryTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:83
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void onEntry(QEvent *)

/*
Reimplemented from QAbstractState::onEntry().
*/
func (this *QHistoryState) OnEntry(event QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QHistoryState7onEntryEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:84
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void onExit(QEvent *)

/*
Reimplemented from QAbstractState::onExit().
*/
func (this *QHistoryState) OnExit(event QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QHistoryState6onExitEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:86
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QHistoryState) Event(e QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QHistoryState5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*
This enum specifies the type of history that a QHistoryState records.


*/
type QHistoryState__HistoryType = int

// Only the immediate child states of the parent state are recorded. In this case a transition with the history state as its target will end up in the immediate child state that the parent was in the last time it was exited. This is the default.
const QHistoryState__ShallowHistory QHistoryState__HistoryType = 0

// Nested states are recorded. In this case a transition with the history state as its target will end up in the most deeply nested descendant state the parent was in the last time it was exited.
const QHistoryState__DeepHistory QHistoryState__HistoryType = 1

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
