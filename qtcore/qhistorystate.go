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

// void onEntry(class QEvent *)
func (this *QHistoryState) InheritOnEntry(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "onEntry", f)
}

// void onExit(class QEvent *)
func (this *QHistoryState) InheritOnExit(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "onExit", f)
}

// bool event(class QEvent *)
func (this *QHistoryState) InheritEvent(f func(e *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

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
// [8] const QMetaObject * metaObject()
func (this *QHistoryState) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QHistoryState10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qhistorystate.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QHistoryState(QState *)
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

// /usr/include/qt/QtCore/qhistorystate.h:65
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QHistoryState(enum QHistoryState::HistoryType, QState *)
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

// /usr/include/qt/QtCore/qhistorystate.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QHistoryState()
func DeleteQHistoryState(this *QHistoryState) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QHistoryStateD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qhistorystate.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractTransition * defaultTransition()
func (this *QHistoryState) DefaultTransition() *QAbstractTransition /*777 QAbstractTransition **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QHistoryState17defaultTransitionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractTransitionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qhistorystate.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultTransition(QAbstractTransition *)
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
// [8] QAbstractState * defaultState()
func (this *QHistoryState) DefaultState() *QAbstractState /*777 QAbstractState **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QHistoryState12defaultStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractStateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qhistorystate.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultState(QAbstractState *)
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
// [4] QHistoryState::HistoryType historyType()
func (this *QHistoryState) HistoryType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QHistoryState11historyTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHistoryType(enum QHistoryState::HistoryType)
func (this *QHistoryState) SetHistoryType(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QHistoryState14setHistoryTypeENS_11HistoryTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:83
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void onEntry(QEvent *)
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
func (this *QHistoryState) Event(e QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QHistoryState5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

type QHistoryState__HistoryType = int

const QHistoryState__ShallowHistory QHistoryState__HistoryType = 0
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
