package qtwidgets

// /usr/include/qt/QtWidgets/qgesturerecognizer.h
// #include <qgesturerecognizer.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QGestureRecognizer) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQGestureRecognizerFromPointer(cthis unsafe.Pointer) *QGestureRecognizer {
	return &QGestureRecognizer{&qtrt.CObject{cthis}}
}
func (*QGestureRecognizer) NewFromPointer(cthis unsafe.Pointer) *QGestureRecognizer {
	return NewQGestureRecognizerFromPointer(cthis)
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
// QGesture * create(QObject *)
func (this *QGestureRecognizer) Create(target *qtcore.QObject /*777 QObject **/) *QGesture /*777 QGesture **/ {
	var convArg0 = target.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGestureRecognizer6createEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGestureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgesturerecognizer.h:81
// index:0
// Public pure virtual
// QGestureRecognizer::Result recognize(QGesture *, QObject *, QEvent *)
func (this *QGestureRecognizer) Recognize(state *QGesture /*777 QGesture **/, watched *qtcore.QObject /*777 QObject **/, event *qtcore.QEvent /*777 QEvent **/) int {
	var convArg0 = state.GetCthis()
	var convArg1 = watched.GetCthis()
	var convArg2 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGestureRecognizer9recognizeEP8QGestureP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgesturerecognizer.h:83
// index:0
// Public virtual
// void reset(QGesture *)
func (this *QGestureRecognizer) Reset(state *QGesture /*777 QGesture **/) {
	var convArg0 = state.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGestureRecognizer5resetEP8QGesture", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesturerecognizer.h:85
// index:0
// Public static
// Qt::GestureType registerRecognizer(QGestureRecognizer *)
func (this *QGestureRecognizer) RegisterRecognizer(recognizer *QGestureRecognizer /*777 QGestureRecognizer **/) int {
	var convArg0 = recognizer.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGestureRecognizer18registerRecognizerEPS_", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QGestureRecognizer_RegisterRecognizer(recognizer *QGestureRecognizer /*777 QGestureRecognizer **/) int {
	var nilthis *QGestureRecognizer
	rv := nilthis.RegisterRecognizer(recognizer)
	return rv
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

type QGestureRecognizer__ResultFlag = int

const QGestureRecognizer__Ignore QGestureRecognizer__ResultFlag = 1
const QGestureRecognizer__MayBeGesture QGestureRecognizer__ResultFlag = 2
const QGestureRecognizer__TriggerGesture QGestureRecognizer__ResultFlag = 4
const QGestureRecognizer__FinishGesture QGestureRecognizer__ResultFlag = 8
const QGestureRecognizer__CancelGesture QGestureRecognizer__ResultFlag = 16
const QGestureRecognizer__ResultState_Mask QGestureRecognizer__ResultFlag = 255
const QGestureRecognizer__ConsumeEventHint QGestureRecognizer__ResultFlag = 256
const QGestureRecognizer__ResultHint_Mask QGestureRecognizer__ResultFlag = 65280

//  body block end
