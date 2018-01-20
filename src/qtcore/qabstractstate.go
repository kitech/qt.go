//  header block begin
// /usr/include/qt/QtCore/qabstractstate.h
// #include <qabstractstate.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 31
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

// /usr/include/qt/QtCore/qabstractstate.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QAbstractState) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QAbstractState10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractstate.h:58
// index:0
// virtual
// void ~QAbstractState()
func DeleteQAbstractState(*QAbstractState) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QAbstractStateD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractstate.h:60
// index:0
// QState * parentState()
func (this *QAbstractState) ParentState() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QAbstractState11parentStateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractstate.h:61
// index:0
// QStateMachine * machine()
func (this *QAbstractState) Machine() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QAbstractState7machineEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractstate.h:63
// index:0
// bool active()
func (this *QAbstractState) Active() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QAbstractState6activeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractstate.h:68
// index:0
// void activeChanged(_Bool)
func (this *QAbstractState) ActiveChanged(active bool) {
	// 0: (, active bool), (&active)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QAbstractState13activeChangedEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &active)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractstate.h:71
// index:0
// void QAbstractState(class QState *)
func NewQAbstractState(parent unsafe.Pointer) *QAbstractState {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QAbstractStateC1EP6QState", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractStateFromPointer(cthis)
	return gothis
}
func NewQAbstractStateFromPointer(cthis unsafe.Pointer) *QAbstractState {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QAbstractState{bcthis0}
}

// /usr/include/qt/QtCore/qabstractstate.h:79
// index:1
// void QAbstractState(class QAbstractStatePrivate &, class QState *)
func NewQAbstractState_1(dd unsafe.Pointer, parent unsafe.Pointer) *QAbstractState {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QAbstractStateC1ER21QAbstractStatePrivateP6QState", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractStateFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qabstractstate.h:73
// index:0
// pure virtual
// void onEntry(class QEvent *)
func (this *QAbstractState) OnEntry(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QAbstractState7onEntryEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractstate.h:74
// index:0
// pure virtual
// void onExit(class QEvent *)
func (this *QAbstractState) OnExit(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QAbstractState6onExitEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractstate.h:76
// index:0
// virtual
// bool event(class QEvent *)
func (this *QAbstractState) Event(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QAbstractState5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

//  body block end
