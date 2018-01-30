package qtcore

// /usr/include/qt/QtCore/qhistorystate.h
// #include <qhistorystate.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 33
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
type QHistoryState struct {
	*QAbstractState
}

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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QHistoryState10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qhistorystate.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QHistoryState(QState *)
func NewQHistoryState(parent *QState /*777 QState **/) *QHistoryState {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QHistoryStateC2EP6QState", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQHistoryStateFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qhistorystate.h:65
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QHistoryState(enum QHistoryState::HistoryType, QState *)
func NewQHistoryState_1(type_ int, parent *QState /*777 QState **/) *QHistoryState {
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QHistoryStateC2ENS_11HistoryTypeEP6QState", ffiqt.FFI_TYPE_POINTER, type_, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQHistoryStateFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qhistorystate.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QHistoryState()
func DeleteQHistoryState(*QHistoryState) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QHistoryStateD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractTransition * defaultTransition()
func (this *QHistoryState) DefaultTransition() *QAbstractTransition /*777 QAbstractTransition **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QHistoryState17defaultTransitionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAbstractTransitionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qhistorystate.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultTransition(QAbstractTransition *)
func (this *QHistoryState) SetDefaultTransition(transition *QAbstractTransition /*777 QAbstractTransition **/) {
	var convArg0 = transition.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QHistoryState20setDefaultTransitionEP19QAbstractTransition", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractState * defaultState()
func (this *QHistoryState) DefaultState() *QAbstractState /*777 QAbstractState **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QHistoryState12defaultStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAbstractStateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qhistorystate.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultState(QAbstractState *)
func (this *QHistoryState) SetDefaultState(state *QAbstractState /*777 QAbstractState **/) {
	var convArg0 = state.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QHistoryState15setDefaultStateEP14QAbstractState", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:74
// index:0
// Public Visibility=Default Availability=Available
// [4] QHistoryState::HistoryType historyType()
func (this *QHistoryState) HistoryType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QHistoryState11historyTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHistoryType(enum QHistoryState::HistoryType)
func (this *QHistoryState) SetHistoryType(type_ int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QHistoryState14setHistoryTypeENS_11HistoryTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:83
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void onEntry(QEvent *)
func (this *QHistoryState) OnEntry(event *QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QHistoryState7onEntryEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:84
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void onExit(QEvent *)
func (this *QHistoryState) OnExit(event *QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QHistoryState6onExitEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:86
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QHistoryState) Event(e *QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QHistoryState5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

type QHistoryState__HistoryType = int

const QHistoryState__ShallowHistory QHistoryState__HistoryType = 0
const QHistoryState__DeepHistory QHistoryState__HistoryType = 1

//  body block end
