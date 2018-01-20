//  header block begin
// /usr/include/qt/QtWidgets/qscroller.h
// #include <qscroller.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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
type QScroller struct {
	*qtcore.QObject
}

func (this *QScroller) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQScrollerFromPointer(cthis unsafe.Pointer) *QScroller {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QScroller{bcthis0}
}

// /usr/include/qt/QtWidgets/qscroller.h:63
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QScroller) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QScroller10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qscroller.h:92
// index:0
// Public static
// bool hasScroller(class QObject *)
func (this *QScroller) HasScroller(target unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller11hasScrollerEP7QObject", ffiqt.FFI_TYPE_POINTER, target)
	gopp.ErrPrint(err, rv)
	return rv
}
func QScroller_HasScroller(target unsafe.Pointer) {
	var nilthis *QScroller
	nilthis.HasScroller(target)
}

// /usr/include/qt/QtWidgets/qscroller.h:94
// index:0
// Public static
// QScroller * scroller(class QObject *)
func (this *QScroller) Scroller(target unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller8scrollerEP7QObject", ffiqt.FFI_TYPE_POINTER, target)
	gopp.ErrPrint(err, rv)
	return rv
}
func QScroller_Scroller(target unsafe.Pointer) {
	var nilthis *QScroller
	nilthis.Scroller(target)
}

// /usr/include/qt/QtWidgets/qscroller.h:95
// index:1
// Public static
// const QScroller * scroller(const class QObject *)
func (this *QScroller) Scroller_1(target unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller8scrollerEPK7QObject", ffiqt.FFI_TYPE_POINTER, target)
	gopp.ErrPrint(err, rv)
	return rv
}
func QScroller_Scroller_1(target unsafe.Pointer) {
	var nilthis *QScroller
	nilthis.Scroller_1(target)
}

// /usr/include/qt/QtWidgets/qscroller.h:98
// index:0
// Public static
// Qt::GestureType grabGesture(class QObject *, enum QScroller::ScrollerGestureType)
func (this *QScroller) GrabGesture(target unsafe.Pointer, gestureType int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller11grabGestureEP7QObjectNS_19ScrollerGestureTypeE", ffiqt.FFI_TYPE_POINTER, target, gestureType)
	gopp.ErrPrint(err, rv)
	return rv
}
func QScroller_GrabGesture(target unsafe.Pointer, gestureType int) {
	var nilthis *QScroller
	nilthis.GrabGesture(target, gestureType)
}

// /usr/include/qt/QtWidgets/qscroller.h:99
// index:0
// Public static
// Qt::GestureType grabbedGesture(class QObject *)
func (this *QScroller) GrabbedGesture(target unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller14grabbedGestureEP7QObject", ffiqt.FFI_TYPE_POINTER, target)
	gopp.ErrPrint(err, rv)
	return rv
}
func QScroller_GrabbedGesture(target unsafe.Pointer) {
	var nilthis *QScroller
	nilthis.GrabbedGesture(target)
}

// /usr/include/qt/QtWidgets/qscroller.h:100
// index:0
// Public static
// void ungrabGesture(class QObject *)
func (this *QScroller) UngrabGesture(target unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller13ungrabGestureEP7QObject", ffiqt.FFI_TYPE_POINTER, target)
	gopp.ErrPrint(err, rv)
}
func QScroller_UngrabGesture(target unsafe.Pointer) {
	var nilthis *QScroller
	nilthis.UngrabGesture(target)
}

// /usr/include/qt/QtWidgets/qscroller.h:103
// index:0
// Public static
// QList<QScroller *> activeScrollers()
func (this *QScroller) ActiveScrollers() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller15activeScrollersEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QScroller_ActiveScrollers() {
	var nilthis *QScroller
	nilthis.ActiveScrollers()
}

// /usr/include/qt/QtWidgets/qscroller.h:105
// index:0
// Public
// QObject * target()
func (this *QScroller) Target() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QScroller6targetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qscroller.h:107
// index:0
// Public
// QScroller::State state()
func (this *QScroller) State() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QScroller5stateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qscroller.h:109
// index:0
// Public
// bool handleInput(enum QScroller::Input, const class QPointF &, qint64)
func (this *QScroller) HandleInput(input int, position *qtcore.QPointF, timestamp int64) interface{} {
	var convArg1 = position.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller11handleInputENS_5InputERK7QPointFx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &input, convArg1, &timestamp)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qscroller.h:111
// index:0
// Public
// void stop()
func (this *QScroller) Stop() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller4stopEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:112
// index:0
// Public
// QPointF velocity()
func (this *QScroller) Velocity() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QScroller8velocityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qscroller.h:113
// index:0
// Public
// QPointF finalPosition()
func (this *QScroller) FinalPosition() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QScroller13finalPositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qscroller.h:114
// index:0
// Public
// QPointF pixelPerMeter()
func (this *QScroller) PixelPerMeter() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QScroller13pixelPerMeterEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qscroller.h:116
// index:0
// Public
// QScrollerProperties scrollerProperties()
func (this *QScroller) ScrollerProperties() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QScroller18scrollerPropertiesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qscroller.h:119
// index:0
// Public
// void setSnapPositionsX(qreal, qreal)
func (this *QScroller) SetSnapPositionsX(first float64, interval float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller17setSnapPositionsXEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &first, &interval)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:121
// index:0
// Public
// void setSnapPositionsY(qreal, qreal)
func (this *QScroller) SetSnapPositionsY(first float64, interval float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller17setSnapPositionsYEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &first, &interval)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:124
// index:0
// Public
// void setScrollerProperties(const class QScrollerProperties &)
func (this *QScroller) SetScrollerProperties(prop *QScrollerProperties) {
	var convArg0 = prop.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller21setScrollerPropertiesERK19QScrollerProperties", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:125
// index:0
// Public
// void scrollTo(const class QPointF &)
func (this *QScroller) ScrollTo(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller8scrollToERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:126
// index:1
// Public
// void scrollTo(const class QPointF &, int)
func (this *QScroller) ScrollTo_1(pos *qtcore.QPointF, scrollTime int) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller8scrollToERK7QPointFi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &scrollTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:127
// index:0
// Public
// void ensureVisible(const class QRectF &, qreal, qreal)
func (this *QScroller) EnsureVisible(rect *qtcore.QRectF, xmargin float64, ymargin float64) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller13ensureVisibleERK6QRectFdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &xmargin, &ymargin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:128
// index:1
// Public
// void ensureVisible(const class QRectF &, qreal, qreal, int)
func (this *QScroller) EnsureVisible_1(rect *qtcore.QRectF, xmargin float64, ymargin float64, scrollTime int) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller13ensureVisibleERK6QRectFddi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &xmargin, &ymargin, &scrollTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:129
// index:0
// Public
// void resendPrepareEvent()
func (this *QScroller) ResendPrepareEvent() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller18resendPrepareEventEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:132
// index:0
// Public
// void stateChanged(class QScroller::State)
func (this *QScroller) StateChanged(newstate int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller12stateChangedENS_5StateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &newstate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:133
// index:0
// Public
// void scrollerPropertiesChanged(const class QScrollerProperties &)
func (this *QScroller) ScrollerPropertiesChanged(arg0 *QScrollerProperties) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller25scrollerPropertiesChangedERK19QScrollerProperties", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
