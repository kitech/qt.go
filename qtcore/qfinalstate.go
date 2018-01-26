package qtcore

// /usr/include/qt/QtCore/qfinalstate.h
// #include <qfinalstate.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
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
type QFinalState struct {
	*QAbstractState
}

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
// Public virtual
// const QMetaObject * metaObject()
func (this *QFinalState) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFinalState10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qfinalstate.h:54
// index:0
// Public
// void QFinalState(class QState *)
func NewQFinalState(parent *QState /*777 QState **/) *QFinalState {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFinalStateC2EP6QState", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
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
func (this *QFinalState) OnEntry(event *QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFinalState7onEntryEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfinalstate.h:59
// index:0
// Protected virtual
// void onExit(class QEvent *)
func (this *QFinalState) OnExit(event *QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFinalState6onExitEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfinalstate.h:61
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QFinalState) Event(e *QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFinalState5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
