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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qeventtransition.h:53
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QEventTransition) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QEventTransition10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventtransition.h:57
// index:0
// void QEventTransition(class QState *)
func NewQEventTransition(sourceState unsafe.Pointer) *QEventTransition {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QEventTransitionC2EP6QState", ffiqt.FFI_TYPE_VOID, cthis, sourceState)
	gopp.ErrPrint(err, rv)
	return &QEventTransition{cthis}
}

// /usr/include/qt/QtCore/qeventtransition.h:58
// index:1
// void QEventTransition(class QObject *, class QEvent::Type, class QState *)
func NewQEventTransition_1(object unsafe.Pointer, type_ int, sourceState unsafe.Pointer) *QEventTransition {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QEventTransitionC2EP7QObjectN6QEvent4TypeEP6QState", ffiqt.FFI_TYPE_VOID, cthis, object, &type_, sourceState)
	gopp.ErrPrint(err, rv)
	return &QEventTransition{cthis}
}

// /usr/include/qt/QtCore/qeventtransition.h:59
// index:0
// virtual
// void ~QEventTransition()
func DeleteQEventTransition(*QEventTransition) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QEventTransitionD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventtransition.h:61
// index:0
// QObject * eventSource()
func (this *QEventTransition) EventSource() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QEventTransition11eventSourceEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventtransition.h:62
// index:0
// void setEventSource(class QObject *)
func (this *QEventTransition) SetEventSource(object unsafe.Pointer) {
	// 0: (, QObject * object), (object)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QEventTransition14setEventSourceEP7QObject", ffiqt.FFI_TYPE_VOID, this.cthis, object)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventtransition.h:64
// index:0
// QEvent::Type eventType()
func (this *QEventTransition) EventType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QEventTransition9eventTypeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventtransition.h:65
// index:0
// void setEventType(class QEvent::Type)
func (this *QEventTransition) SetEventType(type_ int) {
	// 0: (, QEvent::Type type), (&type_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QEventTransition12setEventTypeEN6QEvent4TypeE", ffiqt.FFI_TYPE_VOID, this.cthis, &type_)
	gopp.ErrPrint(err, rv)
}

//  body block end
