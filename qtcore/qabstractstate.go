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
func (this *QAbstractState) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQAbstractStateFromPointer(cthis unsafe.Pointer) *QAbstractState {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QAbstractState{bcthis0}
}
func (*QAbstractState) NewFromPointer(cthis unsafe.Pointer) *QAbstractState {
	return NewQAbstractStateFromPointer(cthis)
}

// /usr/include/qt/QtCore/qabstractstate.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QAbstractState) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QAbstractState10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qabstractstate.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractState()
func DeleteQAbstractState(*QAbstractState) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QAbstractStateD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractstate.h:60
// index:0
// Public Visibility=Default Availability=Available
// [8] QState * parentState()
func (this *QAbstractState) ParentState() *QState /*777 QState **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QAbstractState11parentStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qabstractstate.h:61
// index:0
// Public Visibility=Default Availability=Available
// [8] QStateMachine * machine()
func (this *QAbstractState) Machine() *QStateMachine /*777 QStateMachine **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QAbstractState7machineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStateMachineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qabstractstate.h:63
// index:0
// Public Visibility=Default Availability=Available
// [1] bool active()
func (this *QAbstractState) Active() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QAbstractState6activeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractstate.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activeChanged(_Bool)
func (this *QAbstractState) ActiveChanged(active bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QAbstractState13activeChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), active)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractstate.h:71
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QAbstractState(QState *)
func NewQAbstractState(parent *QState /*777 QState **/) *QAbstractState {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QAbstractStateC1EP6QState", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractStateFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qabstractstate.h:73
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [-2] void onEntry(QEvent *)
func (this *QAbstractState) OnEntry(event *QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QAbstractState7onEntryEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractstate.h:74
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [-2] void onExit(QEvent *)
func (this *QAbstractState) OnExit(event *QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QAbstractState6onExitEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractstate.h:76
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QAbstractState) Event(e *QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QAbstractState5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
