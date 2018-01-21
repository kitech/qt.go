package qtcore

// /usr/include/qt/QtCore/qabstractstate.h
// #include <qabstractstate.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 28
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
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func NewQAbstractStateFromPointer(cthis unsafe.Pointer) *QAbstractState {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QAbstractState{bcthis0}
}

// /usr/include/qt/QtCore/qabstractstate.h:55
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QAbstractState) MetaObject() *QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QAbstractState10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
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
func (this *QAbstractState) ParentState() *QState /*444 QState **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QAbstractState11parentStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qabstractstate.h:61
// index:0
// Public
// QStateMachine * machine()
func (this *QAbstractState) Machine() *QStateMachine /*444 QStateMachine **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QAbstractState7machineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStateMachineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qabstractstate.h:63
// index:0
// Public
// bool active()
func (this *QAbstractState) Active() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QAbstractState6activeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
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
func NewQAbstractState(parent *QState /*444 QState **/) *QAbstractState {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QAbstractStateC1EP6QState", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractStateFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qabstractstate.h:73
// index:0
// Protected pure virtual
// void onEntry(class QEvent *)
func (this *QAbstractState) OnEntry(event *QEvent /*444 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QAbstractState7onEntryEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractstate.h:74
// index:0
// Protected pure virtual
// void onExit(class QEvent *)
func (this *QAbstractState) OnExit(event *QEvent /*444 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QAbstractState6onExitEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractstate.h:76
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QAbstractState) Event(e *QEvent /*444 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QAbstractState5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
