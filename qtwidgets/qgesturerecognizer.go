package qtwidgets

// /usr/include/qt/QtWidgets/qgesturerecognizer.h
// #include <qgesturerecognizer.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

/*

 */
type QGestureRecognizer struct {
	*qtrt.CObject
}
type QGestureRecognizer_ITF interface {
	QGestureRecognizer_PTR() *QGestureRecognizer
}

func (ptr *QGestureRecognizer) QGestureRecognizer_PTR() *QGestureRecognizer { return ptr }

func (this *QGestureRecognizer) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QGestureRecognizer) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQGestureRecognizerFromPointer(cthis unsafe.Pointer) *QGestureRecognizer {
	return &QGestureRecognizer{&qtrt.CObject{cthis}}
}
func (*QGestureRecognizer) NewFromPointer(cthis unsafe.Pointer) *QGestureRecognizer {
	return NewQGestureRecognizerFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgesturerecognizer.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGestureRecognizer()

/*
Constructs a new gesture recognizer object.
*/
func NewQGestureRecognizer() *QGestureRecognizer {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QGestureRecognizerC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGestureRecognizerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGestureRecognizer)
	return gothis
}

// /usr/include/qt/QtWidgets/qgesturerecognizer.h:78
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGestureRecognizer()

/*

 */
func DeleteQGestureRecognizer(this *QGestureRecognizer) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QGestureRecognizerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgesturerecognizer.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QGesture * create(QObject *)

/*
This function is called by Qt to create a new QGesture object for the given target (QWidget or QGraphicsObject).

Reimplement this function to create a custom QGesture-derived gesture object if necessary.

The application takes ownership of the created gesture object.
*/
func (this *QGestureRecognizer) Create(target qtcore.QObject_ITF /*777 QObject **/) *QGesture /*777 QGesture **/ {
	var convArg0 unsafe.Pointer
	if target != nil && target.QObject_PTR() != nil {
		convArg0 = target.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QGestureRecognizer6createEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGestureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgesturerecognizer.h:81
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QGestureRecognizer::Result recognize(QGesture *, QObject *, QEvent *)

/*
Handles the given event for the watched object, updating the state of the gesture object as required, and returns a suitable result for the current recognition step.

This function is called by the framework to allow the recognizer to filter input events dispatched to QWidget or QGraphicsObject instances that it is monitoring.

The result reflects how much of the gesture has been recognized. The state of the gesture object is set depending on the result.

See also QGestureRecognizer::Result.
*/
func (this *QGestureRecognizer) Recognize(state QGesture_ITF /*777 QGesture **/, watched qtcore.QObject_ITF /*777 QObject **/, event qtcore.QEvent_ITF /*777 QEvent **/) int {
	var convArg0 unsafe.Pointer
	if state != nil && state.QGesture_PTR() != nil {
		convArg0 = state.QGesture_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if watched != nil && watched.QObject_PTR() != nil {
		convArg1 = watched.QObject_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg2 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QGestureRecognizer9recognizeEP8QGestureP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgesturerecognizer.h:83
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void reset(QGesture *)

/*
This function is called by the framework to reset a given gesture.

Reimplement this function to implement additional requirements for custom QGesture objects. This may be necessary if you implement a custom QGesture whose properties need special handling when the gesture is reset.
*/
func (this *QGestureRecognizer) Reset(state QGesture_ITF /*777 QGesture **/) {
	var convArg0 unsafe.Pointer
	if state != nil && state.QGesture_PTR() != nil {
		convArg0 = state.QGesture_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QGestureRecognizer5resetEP8QGesture", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesturerecognizer.h:85
// index:0
// Public static Visibility=Default Availability=Available
// [4] Qt::GestureType registerRecognizer(QGestureRecognizer *)

/*
Registers the given recognizer in the gesture framework and returns a gesture ID for it.

The application takes ownership of the recognizer and returns the gesture type ID associated with it. For gesture recognizers which handle custom QGesture objects (i.e., those which return Qt::CustomGesture in a QGesture::gestureType() function) the return value is a generated gesture ID with the Qt::CustomGesture flag set.

See also unregisterRecognizer(), QGestureRecognizer::create(), and QGesture.
*/
func (this *QGestureRecognizer) RegisterRecognizer(recognizer QGestureRecognizer_ITF /*777 QGestureRecognizer **/) int {
	var convArg0 unsafe.Pointer
	if recognizer != nil && recognizer.QGestureRecognizer_PTR() != nil {
		convArg0 = recognizer.QGestureRecognizer_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QGestureRecognizer18registerRecognizerEPS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QGestureRecognizer_RegisterRecognizer(recognizer QGestureRecognizer_ITF /*777 QGestureRecognizer **/) int {
	var nilthis *QGestureRecognizer
	rv := nilthis.RegisterRecognizer(recognizer)
	return rv
}

// /usr/include/qt/QtWidgets/qgesturerecognizer.h:86
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void unregisterRecognizer(Qt::GestureType)

/*
Unregisters all gesture recognizers of the specified type.

See also registerRecognizer().
*/
func (this *QGestureRecognizer) UnregisterRecognizer(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QGestureRecognizer20unregisterRecognizerEN2Qt11GestureTypeE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
}
func QGestureRecognizer_UnregisterRecognizer(type_ int) {
	var nilthis *QGestureRecognizer
	nilthis.UnregisterRecognizer(type_)
}

/*


 */
type QGestureRecognizer__ResultFlag = int

//
const QGestureRecognizer__Ignore QGestureRecognizer__ResultFlag = 1

//
const QGestureRecognizer__MayBeGesture QGestureRecognizer__ResultFlag = 2

//
const QGestureRecognizer__TriggerGesture QGestureRecognizer__ResultFlag = 4

//
const QGestureRecognizer__FinishGesture QGestureRecognizer__ResultFlag = 8

//
const QGestureRecognizer__CancelGesture QGestureRecognizer__ResultFlag = 16

//
const QGestureRecognizer__ResultState_Mask QGestureRecognizer__ResultFlag = 255

//
const QGestureRecognizer__ConsumeEventHint QGestureRecognizer__ResultFlag = 256

//
const QGestureRecognizer__ResultHint_Mask QGestureRecognizer__ResultFlag = 65280

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
