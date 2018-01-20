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
func NewQFinalStateFromPointer(cthis unsafe.Pointer) *QFinalState {
	bcthis0 := NewQAbstractStateFromPointer(cthis)
	return &QFinalState{bcthis0}
}

// /usr/include/qt/QtCore/qfinalstate.h:52
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QFinalState) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFinalState10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfinalstate.h:54
// index:0
// Public
// void QFinalState(class QState *)
func NewQFinalState(parent unsafe.Pointer) *QFinalState {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFinalStateC2EP6QState", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQFinalStateFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qfinalstate.h:55
// index:0
// Public virtual
// void ~QFinalState()
func DeleteQFinalState(*QFinalState) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFinalStateD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfinalstate.h:58
// index:0
// Protected virtual
// void onEntry(class QEvent *)
func (this *QFinalState) OnEntry(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFinalState7onEntryEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfinalstate.h:59
// index:0
// Protected virtual
// void onExit(class QEvent *)
func (this *QFinalState) OnExit(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFinalState6onExitEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfinalstate.h:61
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QFinalState) Event(e unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFinalState5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
