//  header block begin
// /usr/include/qt/QtCore/qfinalstate.h
// #include <qfinalstate.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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
type QFinalState struct {
	*QAbstractState
}

func (this *QFinalState) GetCthis() unsafe.Pointer {
	return this.QAbstractState.GetCthis()
}

// /usr/include/qt/QtCore/qfinalstate.h:52
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QFinalState) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFinalState10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfinalstate.h:54
// index:0
// void QFinalState(class QState *)
func NewQFinalState(parent unsafe.Pointer) *QFinalState {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFinalStateC2EP6QState", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQFinalStateFromPointer(cthis)
	return gothis
}
func NewQFinalStateFromPointer(cthis unsafe.Pointer) *QFinalState {
	bcthis0 := NewQAbstractStateFromPointer(cthis)
	return &QFinalState{bcthis0}
}

// /usr/include/qt/QtCore/qfinalstate.h:64
// index:1
// void QFinalState(class QFinalStatePrivate &, class QState *)
func NewQFinalState_1(dd unsafe.Pointer, parent unsafe.Pointer) *QFinalState {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFinalStateC2ER18QFinalStatePrivateP6QState", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQFinalStateFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qfinalstate.h:55
// index:0
// virtual
// void ~QFinalState()
func DeleteQFinalState(*QFinalState) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFinalStateD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfinalstate.h:58
// index:0
// virtual
// void onEntry(class QEvent *)
func (this *QFinalState) OnEntry(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFinalState7onEntryEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfinalstate.h:59
// index:0
// virtual
// void onExit(class QEvent *)
func (this *QFinalState) OnExit(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFinalState6onExitEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfinalstate.h:61
// index:0
// virtual
// bool event(class QEvent *)
func (this *QFinalState) Event(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFinalState5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

//  body block end
