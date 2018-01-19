//  header block begin
// /usr/include/qt/QtCore/qhistorystate.h
// #include <qhistorystate.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 40
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qhistorystate.h:53
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QHistoryState) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QHistoryState10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:64
// index:0
// void QHistoryState(class QState *)
func NewQHistoryState(parent unsafe.Pointer) *QHistoryState {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QHistoryStateC2EP6QState", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QHistoryState{cthis}
}

// /usr/include/qt/QtCore/qhistorystate.h:65
// index:1
// void QHistoryState(enum QHistoryState::HistoryType, class QState *)
func NewQHistoryState_1(type_ int, parent unsafe.Pointer) *QHistoryState {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QHistoryStateC2ENS_11HistoryTypeEP6QState", ffiqt.FFI_TYPE_VOID, cthis, &type_, parent)
	gopp.ErrPrint(err, rv)
	return &QHistoryState{cthis}
}

// /usr/include/qt/QtCore/qhistorystate.h:66
// index:0
// virtual
// void ~QHistoryState()
func DeleteQHistoryState(*QHistoryState) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QHistoryStateD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:68
// index:0
// QAbstractTransition * defaultTransition()
func (this *QHistoryState) DefaultTransition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QHistoryState17defaultTransitionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:69
// index:0
// void setDefaultTransition(class QAbstractTransition *)
func (this *QHistoryState) SetDefaultTransition(transition unsafe.Pointer) {
	// 0: (, QAbstractTransition * transition), (transition)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QHistoryState20setDefaultTransitionEP19QAbstractTransition", ffiqt.FFI_TYPE_VOID, this.cthis, transition)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:71
// index:0
// QAbstractState * defaultState()
func (this *QHistoryState) DefaultState() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QHistoryState12defaultStateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:72
// index:0
// void setDefaultState(class QAbstractState *)
func (this *QHistoryState) SetDefaultState(state unsafe.Pointer) {
	// 0: (, QAbstractState * state), (state)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QHistoryState15setDefaultStateEP14QAbstractState", ffiqt.FFI_TYPE_VOID, this.cthis, state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:74
// index:0
// QHistoryState::HistoryType historyType()
func (this *QHistoryState) HistoryType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QHistoryState11historyTypeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhistorystate.h:75
// index:0
// void setHistoryType(enum QHistoryState::HistoryType)
func (this *QHistoryState) SetHistoryType(type_ int) {
	// 0: (, QHistoryState::HistoryType type), (&type_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QHistoryState14setHistoryTypeENS_11HistoryTypeE", ffiqt.FFI_TYPE_VOID, this.cthis, &type_)
	gopp.ErrPrint(err, rv)
}

//  body block end
