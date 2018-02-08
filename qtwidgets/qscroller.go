package qtwidgets

// /usr/include/qt/QtWidgets/qscroller.h
// #include <qscroller.h>
// #include <QtWidgets>

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

type QScroller struct {
	*qtcore.QObject
}

func (this *QScroller) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QScroller) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQScrollerFromPointer(cthis unsafe.Pointer) *QScroller {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QScroller{bcthis0}
}
func (*QScroller) NewFromPointer(cthis unsafe.Pointer) *QScroller {
	return NewQScrollerFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qscroller.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QScroller) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QScroller10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qscroller.h:92
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool hasScroller(QObject *)
func (this *QScroller) HasScroller(target *qtcore.QObject /*777 QObject **/) bool {
	var convArg0 = target.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller11hasScrollerEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}
func QScroller_HasScroller(target *qtcore.QObject /*777 QObject **/) bool {
	var nilthis *QScroller
	rv := nilthis.HasScroller(target)
	return rv
}

// /usr/include/qt/QtWidgets/qscroller.h:94
// index:0
// Public static Visibility=Default Availability=Available
// [8] QScroller * scroller(QObject *)
func (this *QScroller) Scroller(target *qtcore.QObject /*777 QObject **/) *QScroller /*777 QScroller **/ {
	var convArg0 = target.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller8scrollerEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQScrollerFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QScroller_Scroller(target *qtcore.QObject /*777 QObject **/) *QScroller /*777 QScroller **/ {
	var nilthis *QScroller
	rv := nilthis.Scroller(target)
	return rv
}

// /usr/include/qt/QtWidgets/qscroller.h:95
// index:1
// Public static Visibility=Default Availability=Available
// [8] const QScroller * scroller(const QObject *)
func (this *QScroller) Scroller_1(target *qtcore.QObject /*777 const QObject **/) *QScroller /*777 const QScroller **/ {
	var convArg0 = target.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller8scrollerEPK7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQScrollerFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QScroller_Scroller_1(target *qtcore.QObject /*777 const QObject **/) *QScroller /*777 const QScroller **/ {
	var nilthis *QScroller
	rv := nilthis.Scroller_1(target)
	return rv
}

// /usr/include/qt/QtWidgets/qscroller.h:98
// index:0
// Public static Visibility=Default Availability=Available
// [4] Qt::GestureType grabGesture(QObject *, enum QScroller::ScrollerGestureType)
func (this *QScroller) GrabGesture(target *qtcore.QObject /*777 QObject **/, gestureType int) int {
	var convArg0 = target.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller11grabGestureEP7QObjectNS_19ScrollerGestureTypeE", qtrt.FFI_TYPE_POINTER, convArg0, gestureType)
	gopp.ErrPrint(err, rv)
	return int(rv)
}
func QScroller_GrabGesture(target *qtcore.QObject /*777 QObject **/, gestureType int) int {
	var nilthis *QScroller
	rv := nilthis.GrabGesture(target, gestureType)
	return rv
}

// /usr/include/qt/QtWidgets/qscroller.h:99
// index:0
// Public static Visibility=Default Availability=Available
// [4] Qt::GestureType grabbedGesture(QObject *)
func (this *QScroller) GrabbedGesture(target *qtcore.QObject /*777 QObject **/) int {
	var convArg0 = target.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller14grabbedGestureEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	return int(rv)
}
func QScroller_GrabbedGesture(target *qtcore.QObject /*777 QObject **/) int {
	var nilthis *QScroller
	rv := nilthis.GrabbedGesture(target)
	return rv
}

// /usr/include/qt/QtWidgets/qscroller.h:100
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void ungrabGesture(QObject *)
func (this *QScroller) UngrabGesture(target *qtcore.QObject /*777 QObject **/) {
	var convArg0 = target.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller13ungrabGestureEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QScroller_UngrabGesture(target *qtcore.QObject /*777 QObject **/) {
	var nilthis *QScroller
	nilthis.UngrabGesture(target)
}

// /usr/include/qt/QtWidgets/qscroller.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * target()
func (this *QScroller) Target() *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QScroller6targetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qscroller.h:107
// index:0
// Public Visibility=Default Availability=Available
// [4] QScroller::State state()
func (this *QScroller) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QScroller5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:109
// index:0
// Public Visibility=Default Availability=Available
// [1] bool handleInput(enum QScroller::Input, const QPointF &, qint64)
func (this *QScroller) HandleInput(input int, position *qtcore.QPointF, timestamp int64) bool {
	var convArg1 = position.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller11handleInputENS_5InputERK7QPointFx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), input, convArg1, timestamp)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qscroller.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stop()
func (this *QScroller) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:112
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF velocity()
func (this *QScroller) Velocity() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QScroller8velocityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qscroller.h:113
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF finalPosition()
func (this *QScroller) FinalPosition() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QScroller13finalPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qscroller.h:114
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF pixelPerMeter()
func (this *QScroller) PixelPerMeter() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QScroller13pixelPerMeterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qscroller.h:116
// index:0
// Public Visibility=Default Availability=Available
// [16] QScrollerProperties scrollerProperties()
func (this *QScroller) ScrollerProperties() *QScrollerProperties /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QScroller18scrollerPropertiesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQScrollerPropertiesFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQScrollerProperties)
	return rv2
}

// /usr/include/qt/QtWidgets/qscroller.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSnapPositionsX(qreal, qreal)
func (this *QScroller) SetSnapPositionsX(first float64, interval float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller17setSnapPositionsXEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), first, interval)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSnapPositionsY(qreal, qreal)
func (this *QScroller) SetSnapPositionsY(first float64, interval float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller17setSnapPositionsYEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), first, interval)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScrollerProperties(const QScrollerProperties &)
func (this *QScroller) SetScrollerProperties(prop *QScrollerProperties) {
	var convArg0 = prop.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller21setScrollerPropertiesERK19QScrollerProperties", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scrollTo(const QPointF &)
func (this *QScroller) ScrollTo(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller8scrollToERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:126
// index:1
// Public Visibility=Default Availability=Available
// [-2] void scrollTo(const QPointF &, int)
func (this *QScroller) ScrollTo_1(pos *qtcore.QPointF, scrollTime int) {
	var convArg0 = pos.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller8scrollToERK7QPointFi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, scrollTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ensureVisible(const QRectF &, qreal, qreal)
func (this *QScroller) EnsureVisible(rect *qtcore.QRectF, xmargin float64, ymargin float64) {
	var convArg0 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller13ensureVisibleERK6QRectFdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xmargin, ymargin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:128
// index:1
// Public Visibility=Default Availability=Available
// [-2] void ensureVisible(const QRectF &, qreal, qreal, int)
func (this *QScroller) EnsureVisible_1(rect *qtcore.QRectF, xmargin float64, ymargin float64, scrollTime int) {
	var convArg0 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller13ensureVisibleERK6QRectFddi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xmargin, ymargin, scrollTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resendPrepareEvent()
func (this *QScroller) ResendPrepareEvent() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller18resendPrepareEventEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(QScroller::State)
func (this *QScroller) StateChanged(newstate int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller12stateChangedENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newstate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scrollerPropertiesChanged(const QScrollerProperties &)
func (this *QScroller) ScrollerPropertiesChanged(arg0 *QScrollerProperties) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller25scrollerPropertiesChangedERK19QScrollerProperties", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

func DeleteQScroller(this *QScroller) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScrollerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QScroller__State = int

const QScroller__Inactive QScroller__State = 0
const QScroller__Pressed QScroller__State = 1
const QScroller__Dragging QScroller__State = 2
const QScroller__Scrolling QScroller__State = 3

type QScroller__ScrollerGestureType = int

const QScroller__TouchGesture QScroller__ScrollerGestureType = 0
const QScroller__LeftMouseButtonGesture QScroller__ScrollerGestureType = 1
const QScroller__RightMouseButtonGesture QScroller__ScrollerGestureType = 2
const QScroller__MiddleMouseButtonGesture QScroller__ScrollerGestureType = 3

type QScroller__Input = int

const QScroller__InputPress QScroller__Input = 1
const QScroller__InputMove QScroller__Input = 2
const QScroller__InputRelease QScroller__Input = 3

//  body block end
