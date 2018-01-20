//  header block begin
// /usr/include/qt/QtCore/qsignaltransition.h
// #include <qsignaltransition.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
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
type QSignalTransition struct {
	*QAbstractTransition
}

func (this *QSignalTransition) GetCthis() unsafe.Pointer {
	return this.QAbstractTransition.GetCthis()
}
func NewQSignalTransitionFromPointer(cthis unsafe.Pointer) *QSignalTransition {
	bcthis0 := NewQAbstractTransitionFromPointer(cthis)
	return &QSignalTransition{bcthis0}
}

// /usr/include/qt/QtCore/qsignaltransition.h:53
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QSignalTransition) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSignalTransition10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsignaltransition.h:58
// index:0
// Public
// void QSignalTransition(class QState *)
func NewQSignalTransition(sourceState unsafe.Pointer) *QSignalTransition {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSignalTransitionC2EP6QState", ffiqt.FFI_TYPE_VOID, cthis, sourceState)
	gopp.ErrPrint(err, rv)
	gothis := NewQSignalTransitionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsignaltransition.h:59
// index:1
// Public
// void QSignalTransition(const class QObject *, const char *, class QState *)
func NewQSignalTransition_1(sender unsafe.Pointer, signal string, sourceState unsafe.Pointer) *QSignalTransition {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg1 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSignalTransitionC2EPK7QObjectPKcP6QState", ffiqt.FFI_TYPE_VOID, cthis, sender, convArg1, sourceState)
	gopp.ErrPrint(err, rv)
	gothis := NewQSignalTransitionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsignaltransition.h:74
// index:0
// Public virtual
// void ~QSignalTransition()
func DeleteQSignalTransition(*QSignalTransition) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSignalTransitionD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignaltransition.h:76
// index:0
// Public
// QObject * senderObject()
func (this *QSignalTransition) SenderObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSignalTransition12senderObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsignaltransition.h:77
// index:0
// Public
// void setSenderObject(const class QObject *)
func (this *QSignalTransition) SetSenderObject(sender unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSignalTransition15setSenderObjectEPK7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), sender)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignaltransition.h:79
// index:0
// Public
// QByteArray signal()
func (this *QSignalTransition) Signal() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSignalTransition6signalEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsignaltransition.h:80
// index:0
// Public
// void setSignal(const class QByteArray &)
func (this *QSignalTransition) SetSignal(signal *QByteArray) {
	var convArg0 = signal.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSignalTransition9setSignalERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignaltransition.h:83
// index:0
// Protected virtual
// bool eventTest(class QEvent *)
func (this *QSignalTransition) EventTest(event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSignalTransition9eventTestEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsignaltransition.h:84
// index:0
// Protected virtual
// void onTransition(class QEvent *)
func (this *QSignalTransition) OnTransition(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSignalTransition12onTransitionEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignaltransition.h:86
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QSignalTransition) Event(e unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSignalTransition5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
