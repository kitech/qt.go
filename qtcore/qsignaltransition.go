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
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
// bool eventTest(class QEvent *)
func (this *QSignalTransition) InheritEventTest(f func(event *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventTest", f)
}

// void onTransition(class QEvent *)
func (this *QSignalTransition) InheritOnTransition(f func(event *QEvent /*777 QEvent **/)) {
	qtrt.SetAllInheritCallback(this, "onTransition", f)
}

// bool event(class QEvent *)
func (this *QSignalTransition) InheritEvent(f func(e *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QSignalTransition) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSignalTransition10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qsignaltransition.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSignalTransition(QState *)
func NewQSignalTransition(sourceState *QState /*777 QState **/) *QSignalTransition {
	var convArg0 = sourceState.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSignalTransitionC2EP6QState", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSignalTransitionFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qsignaltransition.h:59
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSignalTransition(const QObject *, const char *, QState *)
func NewQSignalTransition_1(sender *QObject /*777 const QObject **/, signal string, sourceState *QState /*777 QState **/) *QSignalTransition {
	var convArg0 = sender.GetCthis()
	var convArg1 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg1)
	var convArg2 = sourceState.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSignalTransitionC2EPK7QObjectPKcP6QState", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQSignalTransitionFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qsignaltransition.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSignalTransition()
func DeleteQSignalTransition(this *QSignalTransition) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSignalTransitionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qsignaltransition.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * senderObject()
func (this *QSignalTransition) SenderObject() *QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSignalTransition12senderObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qsignaltransition.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSenderObject(const QObject *)
func (this *QSignalTransition) SetSenderObject(sender *QObject /*777 const QObject **/) {
	var convArg0 = sender.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSignalTransition15setSenderObjectEPK7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignaltransition.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray signal()
func (this *QSignalTransition) Signal() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSignalTransition6signalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qsignaltransition.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSignal(const QByteArray &)
func (this *QSignalTransition) SetSignal(signal *QByteArray) {
	var convArg0 = signal.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSignalTransition9setSignalERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignaltransition.h:83
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventTest(QEvent *)
func (this *QSignalTransition) EventTest(event *QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSignalTransition9eventTestEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qsignaltransition.h:84
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void onTransition(QEvent *)
func (this *QSignalTransition) OnTransition(event *QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSignalTransition12onTransitionEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignaltransition.h:86
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QSignalTransition) Event(e *QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSignalTransition5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
