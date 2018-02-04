package qtwidgets

// /usr/include/qt/QtWidgets/qgesture.h
// #include <qgesture.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
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

type QGestureEvent struct {
	*qtcore.QEvent
}

func (this *QGestureEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QGestureEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQGestureEventFromPointer(cthis unsafe.Pointer) *QGestureEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QGestureEvent{bcthis0}
}
func (*QGestureEvent) NewFromPointer(cthis unsafe.Pointer) *QGestureEvent {
	return NewQGestureEventFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgesture.h:278
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGestureEvent()
func DeleteQGestureEvent(this *QGestureEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGestureEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 56)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgesture.h:281
// index:0
// Public Visibility=Default Availability=Available
// [8] QGesture * gesture(Qt::GestureType)
func (this *QGestureEvent) Gesture(type_ int) *QGesture /*777 QGesture **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGestureEvent7gestureEN2Qt11GestureTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGestureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgesture.h:291
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAccepted(QGesture *, _Bool)
func (this *QGestureEvent) SetAccepted(arg0 *QGesture /*777 QGesture **/, arg1 bool) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGestureEvent11setAcceptedEP8QGestureb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:296
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setAccepted(Qt::GestureType, _Bool)
func (this *QGestureEvent) SetAccepted_1(arg0 int, arg1 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGestureEvent11setAcceptedEN2Qt11GestureTypeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:292
// index:0
// Public Visibility=Default Availability=Available
// [-2] void accept(QGesture *)
func (this *QGestureEvent) Accept(arg0 *QGesture /*777 QGesture **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGestureEvent6acceptEP8QGesture", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:297
// index:1
// Public Visibility=Default Availability=Available
// [-2] void accept(Qt::GestureType)
func (this *QGestureEvent) Accept_1(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGestureEvent6acceptEN2Qt11GestureTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:293
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ignore(QGesture *)
func (this *QGestureEvent) Ignore(arg0 *QGesture /*777 QGesture **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGestureEvent6ignoreEP8QGesture", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:298
// index:1
// Public Visibility=Default Availability=Available
// [-2] void ignore(Qt::GestureType)
func (this *QGestureEvent) Ignore_1(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGestureEvent6ignoreEN2Qt11GestureTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:294
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAccepted(QGesture *)
func (this *QGestureEvent) IsAccepted(arg0 *QGesture /*777 QGesture **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGestureEvent10isAcceptedEP8QGesture", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgesture.h:299
// index:1
// Public Visibility=Default Availability=Available
// [1] bool isAccepted(Qt::GestureType)
func (this *QGestureEvent) IsAccepted_1(arg0 int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGestureEvent10isAcceptedEN2Qt11GestureTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgesture.h:301
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidget(QWidget *)
func (this *QGestureEvent) SetWidget(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGestureEvent9setWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:302
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * widget()
func (this *QGestureEvent) Widget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGestureEvent6widgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgesture.h:305
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF mapToGraphicsScene(const QPointF &)
func (this *QGestureEvent) MapToGraphicsScene(gesturePoint *qtcore.QPointF) *qtcore.QPointF /*123*/ {
	var convArg0 = gesturePoint.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGestureEvent18mapToGraphicsSceneERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

//  body block end
