//  header block begin
// /usr/include/qt/QtCore/qabstractstate.h
// #include <qabstractstate.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 30
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
type QAbstractState struct {
	*QObject
}

func (this *QAbstractState) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQAbstractStateFromPointer(cthis unsafe.Pointer) *QAbstractState {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QAbstractState{bcthis0}
}

// /usr/include/qt/QtCore/qabstractstate.h:55
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QAbstractState) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QAbstractState10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractstate.h:58
// index:0
// Public virtual
// void ~QAbstractState()
func DeleteQAbstractState(*QAbstractState) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QAbstractStateD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractstate.h:60
// index:0
// Public
// QState * parentState()
func (this *QAbstractState) ParentState() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QAbstractState11parentStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractstate.h:61
// index:0
// Public
// QStateMachine * machine()
func (this *QAbstractState) Machine() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QAbstractState7machineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractstate.h:63
// index:0
// Public
// bool active()
func (this *QAbstractState) Active() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QAbstractState6activeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractstate.h:68
// index:0
// Public
// void activeChanged(_Bool)
func (this *QAbstractState) ActiveChanged(active bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QAbstractState13activeChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &active)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractstate.h:71
// index:0
// Protected
// void QAbstractState(class QState *)
func NewQAbstractState(parent unsafe.Pointer) *QAbstractState {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QAbstractStateC1EP6QState", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractStateFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qabstractstate.h:73
// index:0
// Protected pure virtual
// void onEntry(class QEvent *)
func (this *QAbstractState) OnEntry(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QAbstractState7onEntryEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractstate.h:74
// index:0
// Protected pure virtual
// void onExit(class QEvent *)
func (this *QAbstractState) OnExit(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QAbstractState6onExitEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractstate.h:76
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QAbstractState) Event(e unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QAbstractState5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
