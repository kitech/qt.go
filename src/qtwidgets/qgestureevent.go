//  header block begin
// /usr/include/qt/QtWidgets/qgesture.h
// #include <qgesture.h>
// #include <QtWidgets>
package qtwidgets

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
type QGestureEvent struct {
	*qtcore.QEvent
}

func (this *QGestureEvent) GetCthis() unsafe.Pointer {
	return this.QEvent.GetCthis()
}

// /usr/include/qt/QtWidgets/qgesture.h:278
// index:0
// virtual
// void ~QGestureEvent()
func DeleteQGestureEvent(*QGestureEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGestureEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:280
// index:0
// QList<QGesture *> gestures()
func (this *QGestureEvent) Gestures() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGestureEvent8gesturesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:281
// index:0
// QGesture * gesture(Qt::GestureType)
func (this *QGestureEvent) Gesture(type_ int) {
	// 0: (, type Qt::GestureType), (&type_)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGestureEvent7gestureEN2Qt11GestureTypeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:283
// index:0
// QList<QGesture *> activeGestures()
func (this *QGestureEvent) ActiveGestures() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGestureEvent14activeGesturesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:284
// index:0
// QList<QGesture *> canceledGestures()
func (this *QGestureEvent) CanceledGestures() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGestureEvent16canceledGesturesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:291
// index:0
// void setAccepted(class QGesture *, _Bool)
func (this *QGestureEvent) SetAccepted(arg0 unsafe.Pointer, arg1 bool) {
	// 0: (, QGesture * arg0, bool arg1), (arg0, &arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGestureEvent11setAcceptedEP8QGestureb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, &arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:296
// index:1
// void setAccepted(Qt::GestureType, _Bool)
func (this *QGestureEvent) SetAccepted_1(arg0 int, arg1 bool) {
	// 1: (, Qt::GestureType arg0, bool arg1), (&arg0, &arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGestureEvent11setAcceptedEN2Qt11GestureTypeEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0, &arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:292
// index:0
// void accept(class QGesture *)
func (this *QGestureEvent) Accept(arg0 unsafe.Pointer) {
	// 0: (, QGesture * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGestureEvent6acceptEP8QGesture", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:297
// index:1
// void accept(Qt::GestureType)
func (this *QGestureEvent) Accept_1(arg0 int) {
	// 1: (, Qt::GestureType arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGestureEvent6acceptEN2Qt11GestureTypeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:293
// index:0
// void ignore(class QGesture *)
func (this *QGestureEvent) Ignore(arg0 unsafe.Pointer) {
	// 0: (, QGesture * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGestureEvent6ignoreEP8QGesture", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:298
// index:1
// void ignore(Qt::GestureType)
func (this *QGestureEvent) Ignore_1(arg0 int) {
	// 1: (, Qt::GestureType arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGestureEvent6ignoreEN2Qt11GestureTypeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:294
// index:0
// bool isAccepted(class QGesture *)
func (this *QGestureEvent) IsAccepted(arg0 unsafe.Pointer) {
	// 0: (, QGesture * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGestureEvent10isAcceptedEP8QGesture", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:299
// index:1
// bool isAccepted(Qt::GestureType)
func (this *QGestureEvent) IsAccepted_1(arg0 int) {
	// 1: (, Qt::GestureType arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGestureEvent10isAcceptedEN2Qt11GestureTypeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:301
// index:0
// void setWidget(class QWidget *)
func (this *QGestureEvent) SetWidget(widget unsafe.Pointer) {
	// 0: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGestureEvent9setWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:302
// index:0
// QWidget * widget()
func (this *QGestureEvent) Widget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGestureEvent6widgetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:305
// index:0
// QPointF mapToGraphicsScene(const class QPointF &)
func (this *QGestureEvent) MapToGraphicsScene(gesturePoint unsafe.Pointer) {
	// 0: (, gesturePoint const QPointF &), (gesturePoint)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGestureEvent18mapToGraphicsSceneERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), gesturePoint)
	gopp.ErrPrint(err, rv)
}

//  body block end
