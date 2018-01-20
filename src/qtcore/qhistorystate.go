//  header block begin
// /usr/include/qt/QtCore/qhistorystate.h
// #include <qhistorystate.h>
// #include <QtCore>
package qtcore

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
type QHistoryState struct {
	*QAbstractState
}

func (this *QHistoryState) GetCthis() unsafe.Pointer {
	return this.QAbstractState.GetCthis()
}
func NewQHistoryStateFromPointer(cthis unsafe.Pointer) *QHistoryState {
	bcthis0 := NewQAbstractStateFromPointer(cthis)
	return &QHistoryState{bcthis0}
}

// /usr/include/qt/QtCore/qhistorystate.h:53
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QHistoryState) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QHistoryState10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qhistorystate.h:64
// index:0
// Public
// void QHistoryState(class QState *)
func NewQHistoryState(parent unsafe.Pointer) *QHistoryState {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QHistoryStateC2EP6QState", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQHistoryStateFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qhistorystate.h:65
// index:1
// Public
// void QHistoryState(enum QHistoryState::HistoryType, class QState *)
func NewQHistoryState_1(type_ int, parent unsafe.Pointer) *QHistoryState {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QHistoryStateC2ENS_11HistoryTypeEP6QState", ffiqt.FFI_TYPE_VOID, cthis, &type_, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQHistoryStateFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qhistorystate.h:66
// index:0
// Public virtual
// void ~QHistoryState()
func DeleteQHistoryState(*QHistoryState) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QHistoryStateD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:68
// index:0
// Public
// QAbstractTransition * defaultTransition()
func (this *QHistoryState) DefaultTransition() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QHistoryState17defaultTransitionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qhistorystate.h:69
// index:0
// Public
// void setDefaultTransition(class QAbstractTransition *)
func (this *QHistoryState) SetDefaultTransition(transition unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QHistoryState20setDefaultTransitionEP19QAbstractTransition", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), transition)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:71
// index:0
// Public
// QAbstractState * defaultState()
func (this *QHistoryState) DefaultState() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QHistoryState12defaultStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qhistorystate.h:72
// index:0
// Public
// void setDefaultState(class QAbstractState *)
func (this *QHistoryState) SetDefaultState(state unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QHistoryState15setDefaultStateEP14QAbstractState", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:74
// index:0
// Public
// QHistoryState::HistoryType historyType()
func (this *QHistoryState) HistoryType() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QHistoryState11historyTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qhistorystate.h:75
// index:0
// Public
// void setHistoryType(enum QHistoryState::HistoryType)
func (this *QHistoryState) SetHistoryType(type_ int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QHistoryState14setHistoryTypeENS_11HistoryTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:83
// index:0
// Protected virtual
// void onEntry(class QEvent *)
func (this *QHistoryState) OnEntry(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QHistoryState7onEntryEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:84
// index:0
// Protected virtual
// void onExit(class QEvent *)
func (this *QHistoryState) OnExit(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QHistoryState6onExitEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:86
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QHistoryState) Event(e unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QHistoryState5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
