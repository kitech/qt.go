//  header block begin
// /usr/include/qt/QtCore/qeventtransition.h
// #include <qeventtransition.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
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
type QEventTransition struct {
	*QAbstractTransition
}

func (this *QEventTransition) GetCthis() unsafe.Pointer {
	return this.QAbstractTransition.GetCthis()
}
func NewQEventTransitionFromPointer(cthis unsafe.Pointer) *QEventTransition {
	bcthis0 := NewQAbstractTransitionFromPointer(cthis)
	return &QEventTransition{bcthis0}
}

// /usr/include/qt/QtCore/qeventtransition.h:53
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QEventTransition) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QEventTransition10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qeventtransition.h:57
// index:0
// Public
// void QEventTransition(class QState *)
func NewQEventTransition(sourceState unsafe.Pointer) *QEventTransition {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QEventTransitionC2EP6QState", ffiqt.FFI_TYPE_VOID, cthis, sourceState)
	gopp.ErrPrint(err, rv)
	gothis := NewQEventTransitionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qeventtransition.h:58
// index:1
// Public
// void QEventTransition(class QObject *, class QEvent::Type, class QState *)
func NewQEventTransition_1(object unsafe.Pointer, type_ int, sourceState unsafe.Pointer) *QEventTransition {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QEventTransitionC2EP7QObjectN6QEvent4TypeEP6QState", ffiqt.FFI_TYPE_VOID, cthis, object, &type_, sourceState)
	gopp.ErrPrint(err, rv)
	gothis := NewQEventTransitionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qeventtransition.h:59
// index:0
// Public virtual
// void ~QEventTransition()
func DeleteQEventTransition(*QEventTransition) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QEventTransitionD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventtransition.h:61
// index:0
// Public
// QObject * eventSource()
func (this *QEventTransition) EventSource() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QEventTransition11eventSourceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qeventtransition.h:62
// index:0
// Public
// void setEventSource(class QObject *)
func (this *QEventTransition) SetEventSource(object unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QEventTransition14setEventSourceEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), object)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventtransition.h:64
// index:0
// Public
// QEvent::Type eventType()
func (this *QEventTransition) EventType() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QEventTransition9eventTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qeventtransition.h:65
// index:0
// Public
// void setEventType(class QEvent::Type)
func (this *QEventTransition) SetEventType(type_ int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QEventTransition12setEventTypeEN6QEvent4TypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventtransition.h:68
// index:0
// Protected virtual
// bool eventTest(class QEvent *)
func (this *QEventTransition) EventTest(event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QEventTransition9eventTestEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qeventtransition.h:69
// index:0
// Protected virtual
// void onTransition(class QEvent *)
func (this *QEventTransition) OnTransition(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QEventTransition12onTransitionEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventtransition.h:71
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QEventTransition) Event(e unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QEventTransition5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
