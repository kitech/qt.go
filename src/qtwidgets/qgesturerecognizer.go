//  header block begin
// /usr/include/qt/QtWidgets/qgesturerecognizer.h
// #include <qgesturerecognizer.h>
// #include <QtWidgets>
package qtwidgets

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
import "qtcore"
import "qtgui"

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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QGestureRecognizer struct {
	*qtrt.CObject
}

func (this *QGestureRecognizer) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQGestureRecognizerFromPointer(cthis unsafe.Pointer) *QGestureRecognizer {
	return &QGestureRecognizer{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtWidgets/qgesturerecognizer.h:77
// index:0
// Public
// void QGestureRecognizer()
func NewQGestureRecognizer() *QGestureRecognizer {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGestureRecognizerC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQGestureRecognizerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgesturerecognizer.h:78
// index:0
// Public virtual
// void ~QGestureRecognizer()
func DeleteQGestureRecognizer(*QGestureRecognizer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGestureRecognizerD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesturerecognizer.h:80
// index:0
// Public virtual
// QGesture * create(class QObject *)
func (this *QGestureRecognizer) Create(target unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGestureRecognizer6createEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), target)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgesturerecognizer.h:81
// index:0
// Public pure virtual
// QGestureRecognizer::Result recognize(class QGesture *, class QObject *, class QEvent *)
func (this *QGestureRecognizer) Recognize(state unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGestureRecognizer9recognizeEP8QGestureP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), state, watched, event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgesturerecognizer.h:83
// index:0
// Public virtual
// void reset(class QGesture *)
func (this *QGestureRecognizer) Reset(state unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGestureRecognizer5resetEP8QGesture", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesturerecognizer.h:85
// index:0
// Public static
// Qt::GestureType registerRecognizer(class QGestureRecognizer *)
func (this *QGestureRecognizer) RegisterRecognizer(recognizer unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGestureRecognizer18registerRecognizerEPS_", ffiqt.FFI_TYPE_POINTER, recognizer)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGestureRecognizer_RegisterRecognizer(recognizer unsafe.Pointer) {
	var nilthis *QGestureRecognizer
	nilthis.RegisterRecognizer(recognizer)
}

// /usr/include/qt/QtWidgets/qgesturerecognizer.h:86
// index:0
// Public static
// void unregisterRecognizer(Qt::GestureType)
func (this *QGestureRecognizer) UnregisterRecognizer(type_ int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGestureRecognizer20unregisterRecognizerEN2Qt11GestureTypeE", ffiqt.FFI_TYPE_POINTER, type_)
	gopp.ErrPrint(err, rv)
}
func QGestureRecognizer_UnregisterRecognizer(type_ int) {
	var nilthis *QGestureRecognizer
	nilthis.UnregisterRecognizer(type_)
}

//  body block end
