package qtcore

// /usr/include/qt/QtCore/qfinalstate.h
// #include <qfinalstate.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// void onEntry(class QEvent *)
func (this *QFinalState) InheritOnEntry(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "onEntry", f)
}

// void onExit(class QEvent *)
func (this *QFinalState) InheritOnExit(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "onExit", f)
}

// bool event(class QEvent *)
func (this *QFinalState) InheritEvent(f func(e *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

type QFinalState struct {
	*QAbstractState
}
type QFinalState_ITF interface {
	QAbstractState_ITF
	QFinalState_PTR() *QFinalState
}

func (ptr *QFinalState) QFinalState_PTR() *QFinalState { return ptr }

func (this *QFinalState) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractState.GetCthis()
	}
}
func (this *QFinalState) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractState = NewQAbstractStateFromPointer(cthis)
}
func NewQFinalStateFromPointer(cthis unsafe.Pointer) *QFinalState {
	bcthis0 := NewQAbstractStateFromPointer(cthis)
	return &QFinalState{bcthis0}
}
func (*QFinalState) NewFromPointer(cthis unsafe.Pointer) *QFinalState {
	return NewQFinalStateFromPointer(cthis)
}

// /usr/include/qt/QtCore/qfinalstate.h:52
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QFinalState) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFinalState10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qfinalstate.h:54
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFinalState(QState *)
func NewQFinalState(parent QState_ITF /*777 QState **/) *QFinalState {
	var convArg0 = parent.QState_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFinalStateC2EP6QState", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFinalStateFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qfinalstate.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QFinalState()
func DeleteQFinalState(this *QFinalState) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFinalStateD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qfinalstate.h:58
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void onEntry(QEvent *)
func (this *QFinalState) OnEntry(event QEvent_ITF /*777 QEvent **/) {
	var convArg0 = event.QEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFinalState7onEntryEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfinalstate.h:59
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void onExit(QEvent *)
func (this *QFinalState) OnExit(event QEvent_ITF /*777 QEvent **/) {
	var convArg0 = event.QEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFinalState6onExitEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfinalstate.h:61
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QFinalState) Event(e QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 = e.QEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFinalState5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

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
		qtrt.KeepMe()
	}
}

//  keep block end
