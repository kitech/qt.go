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

// /usr/include/qt/QtWidgets/qscroller.h:63
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QScroller) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QScroller10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:92
// index:0
// static
// bool hasScroller(class QObject *)
func (this *QScroller) HasScroller(target unsafe.Pointer) {
	// 0: (target QObject *), (target)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller11hasScrollerEP7QObject", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QScroller_HasScroller(target unsafe.Pointer) {
	// 0: (target QObject *), (target)
	var nilthis *QScroller
	nilthis.HasScroller(target)
}

// /usr/include/qt/QtWidgets/qscroller.h:94
// index:0
// static
// QScroller * scroller(class QObject *)
func (this *QScroller) Scroller(target unsafe.Pointer) {
	// 0: (target QObject *), (target)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller8scrollerEP7QObject", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QScroller_Scroller(target unsafe.Pointer) {
	// 0: (target QObject *), (target)
	var nilthis *QScroller
	nilthis.Scroller(target)
}

// /usr/include/qt/QtWidgets/qscroller.h:95
// index:1
// static
// const QScroller * scroller(const class QObject *)
func (this *QScroller) Scroller_1(target unsafe.Pointer) {
	// 1: (target const QObject *), (target)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller8scrollerEPK7QObject", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QScroller_Scroller_1(target unsafe.Pointer) {
	// 1: (target const QObject *), (target)
	var nilthis *QScroller
	nilthis.Scroller_1(target)
}

// /usr/include/qt/QtWidgets/qscroller.h:98
// index:0
// static
// Qt::GestureType grabGesture(class QObject *, enum QScroller::ScrollerGestureType)
func (this *QScroller) GrabGesture(target unsafe.Pointer, gestureType int) {
	// 0: (target QObject *, gestureType QScroller::ScrollerGestureType), (target, gestureType)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller11grabGestureEP7QObjectNS_19ScrollerGestureTypeE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QScroller_GrabGesture(target unsafe.Pointer, gestureType int) {
	// 0: (target QObject *, gestureType QScroller::ScrollerGestureType), (target, gestureType)
	var nilthis *QScroller
	nilthis.GrabGesture(target, gestureType)
}

// /usr/include/qt/QtWidgets/qscroller.h:99
// index:0
// static
// Qt::GestureType grabbedGesture(class QObject *)
func (this *QScroller) GrabbedGesture(target unsafe.Pointer) {
	// 0: (target QObject *), (target)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller14grabbedGestureEP7QObject", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QScroller_GrabbedGesture(target unsafe.Pointer) {
	// 0: (target QObject *), (target)
	var nilthis *QScroller
	nilthis.GrabbedGesture(target)
}

// /usr/include/qt/QtWidgets/qscroller.h:100
// index:0
// static
// void ungrabGesture(class QObject *)
func (this *QScroller) UngrabGesture(target unsafe.Pointer) {
	// 0: (target QObject *), (target)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller13ungrabGestureEP7QObject", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QScroller_UngrabGesture(target unsafe.Pointer) {
	// 0: (target QObject *), (target)
	var nilthis *QScroller
	nilthis.UngrabGesture(target)
}

// /usr/include/qt/QtWidgets/qscroller.h:103
// index:0
// static
// QList<QScroller *> activeScrollers()
func (this *QScroller) ActiveScrollers() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller15activeScrollersEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QScroller_ActiveScrollers() {
	// 0: (), ()
	var nilthis *QScroller
	nilthis.ActiveScrollers()
}

// /usr/include/qt/QtWidgets/qscroller.h:105
// index:0
// QObject * target()
func (this *QScroller) Target() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QScroller6targetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:107
// index:0
// QScroller::State state()
func (this *QScroller) State() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QScroller5stateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:109
// index:0
// bool handleInput(enum QScroller::Input, const class QPointF &, qint64)
func (this *QScroller) HandleInput(input int, position unsafe.Pointer, timestamp int64) {
	// 0: (, input QScroller::Input, position const QPointF &, timestamp qint64), (&input, position, &timestamp)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller11handleInputENS_5InputERK7QPointFx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &input, position, &timestamp)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:111
// index:0
// void stop()
func (this *QScroller) Stop() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller4stopEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:112
// index:0
// QPointF velocity()
func (this *QScroller) Velocity() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QScroller8velocityEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:113
// index:0
// QPointF finalPosition()
func (this *QScroller) FinalPosition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QScroller13finalPositionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:114
// index:0
// QPointF pixelPerMeter()
func (this *QScroller) PixelPerMeter() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QScroller13pixelPerMeterEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:116
// index:0
// QScrollerProperties scrollerProperties()
func (this *QScroller) ScrollerProperties() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QScroller18scrollerPropertiesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:119
// index:0
// void setSnapPositionsX(qreal, qreal)
func (this *QScroller) SetSnapPositionsX(first float64, interval float64) {
	// 0: (, first qreal, interval qreal), (&first, &interval)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller17setSnapPositionsXEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &first, &interval)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:121
// index:0
// void setSnapPositionsY(qreal, qreal)
func (this *QScroller) SetSnapPositionsY(first float64, interval float64) {
	// 0: (, first qreal, interval qreal), (&first, &interval)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller17setSnapPositionsYEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &first, &interval)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:124
// index:0
// void setScrollerProperties(const class QScrollerProperties &)
func (this *QScroller) SetScrollerProperties(prop unsafe.Pointer) {
	// 0: (, prop const QScrollerProperties &), (prop)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller21setScrollerPropertiesERK19QScrollerProperties", ffiqt.FFI_TYPE_VOID, this.GetCthis(), prop)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:125
// index:0
// void scrollTo(const class QPointF &)
func (this *QScroller) ScrollTo(pos unsafe.Pointer) {
	// 0: (, pos const QPointF &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller8scrollToERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:126
// index:1
// void scrollTo(const class QPointF &, int)
func (this *QScroller) ScrollTo_1(pos unsafe.Pointer, scrollTime int) {
	// 1: (, pos const QPointF &, scrollTime int), (pos, &scrollTime)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller8scrollToERK7QPointFi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos, &scrollTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:127
// index:0
// void ensureVisible(const class QRectF &, qreal, qreal)
func (this *QScroller) EnsureVisible(rect unsafe.Pointer, xmargin float64, ymargin float64) {
	// 0: (, rect const QRectF &, xmargin qreal, ymargin qreal), (rect, &xmargin, &ymargin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller13ensureVisibleERK6QRectFdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &xmargin, &ymargin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:128
// index:1
// void ensureVisible(const class QRectF &, qreal, qreal, int)
func (this *QScroller) EnsureVisible_1(rect unsafe.Pointer, xmargin float64, ymargin float64, scrollTime int) {
	// 1: (, rect const QRectF &, xmargin qreal, ymargin qreal, scrollTime int), (rect, &xmargin, &ymargin, &scrollTime)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller13ensureVisibleERK6QRectFddi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &xmargin, &ymargin, &scrollTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:129
// index:0
// void resendPrepareEvent()
func (this *QScroller) ResendPrepareEvent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller18resendPrepareEventEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:132
// index:0
// void stateChanged(class QScroller::State)
func (this *QScroller) StateChanged(newstate int) {
	// 0: (, newstate QScroller::State), (&newstate)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller12stateChangedENS_5StateE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &newstate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:133
// index:0
// void scrollerPropertiesChanged(const class QScrollerProperties &)
func (this *QScroller) ScrollerPropertiesChanged(arg0 unsafe.Pointer) {
	// 0: (, const QScrollerProperties & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QScroller25scrollerPropertiesChangedERK19QScrollerProperties", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
