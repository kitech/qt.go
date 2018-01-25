package qtcore

// /usr/include/qt/QtCore/qsignaltransition.h
// #include <qsignaltransition.h>
// #include <QtCore>

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
	if this == nil {
		return nil
	} else {
		return this.QAbstractTransition.GetCthis()
	}
}
func (this *QSignalTransition) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractTransition = NewQAbstractTransitionFromPointer(cthis)
}
func NewQSignalTransitionFromPointer(cthis unsafe.Pointer) *QSignalTransition {
	bcthis0 := NewQAbstractTransitionFromPointer(cthis)
	return &QSignalTransition{bcthis0}
}
func (*QSignalTransition) NewFromPointer(cthis unsafe.Pointer) *QSignalTransition {
	return NewQSignalTransitionFromPointer(cthis)
}

// /usr/include/qt/QtCore/qsignaltransition.h:53
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QSignalTransition) MetaObject() *QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSignalTransition10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qsignaltransition.h:58
// index:0
// Public
// void QSignalTransition(class QState *)
func NewQSignalTransition(sourceState *QState /*444 QState **/) *QSignalTransition {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = sourceState.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSignalTransitionC2EP6QState", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSignalTransitionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsignaltransition.h:59
// index:1
// Public
// void QSignalTransition(const class QObject *, const char *, class QState *)
func NewQSignalTransition_1(sender *QObject /*444 const QObject **/, signal string, sourceState *QState /*444 QState **/) *QSignalTransition {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = sender.GetCthis()
	var convArg1 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg1)
	var convArg2 = sourceState.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSignalTransitionC2EPK7QObjectPKcP6QState", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, convArg2)
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
func (this *QSignalTransition) SenderObject() *QObject /*444 QObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSignalTransition12senderObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qsignaltransition.h:77
// index:0
// Public
// void setSenderObject(const class QObject *)
func (this *QSignalTransition) SetSenderObject(sender *QObject /*444 const QObject **/) {
	var convArg0 = sender.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSignalTransition15setSenderObjectEPK7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignaltransition.h:79
// index:0
// Public
// QByteArray signal()
func (this *QSignalTransition) Signal() *QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSignalTransition6signalEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
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
func (this *QSignalTransition) EventTest(event *QEvent /*444 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSignalTransition9eventTestEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qsignaltransition.h:84
// index:0
// Protected virtual
// void onTransition(class QEvent *)
func (this *QSignalTransition) OnTransition(event *QEvent /*444 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSignalTransition12onTransitionEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignaltransition.h:86
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QSignalTransition) Event(e *QEvent /*444 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSignalTransition5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
