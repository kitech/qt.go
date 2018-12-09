// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qscroller.h
// #include <qscroller.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
type QScroller struct {
	*qtcore.QObject
}
type QScroller_ITF interface {
	qtcore.QObject_ITF
	QScroller_PTR() *QScroller
}

func (ptr *QScroller) QScroller_PTR() *QScroller { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QScroller) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QScroller10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qscroller.h:92
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool hasScroller(QObject *)

/*
Returns true if a QScroller object was already created for target; false otherwise.

See also scroller().
*/
func (this *QScroller) HasScroller(target qtcore.QObject_ITF /*777 QObject **/) bool {
	var convArg0 unsafe.Pointer
	if target != nil && target.QObject_PTR() != nil {
		convArg0 = target.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller11hasScrollerEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QScroller_HasScroller(target qtcore.QObject_ITF /*777 QObject **/) bool {
	var nilthis *QScroller
	rv := nilthis.HasScroller(target)
	return rv
}

// /usr/include/qt/QtWidgets/qscroller.h:94
// index:0
// Public static Visibility=Default Availability=Available
// [8] QScroller * scroller(QObject *)

/*
Returns the scroller for the given target. As long as the object exists this function will always return the same QScroller instance. If no QScroller exists for the target, one will implicitly be created. At no point more than one QScroller will be active on an object.

See also hasScroller() and target().
*/
func (this *QScroller) Scroller(target qtcore.QObject_ITF /*777 QObject **/) *QScroller /*777 QScroller **/ {
	var convArg0 unsafe.Pointer
	if target != nil && target.QObject_PTR() != nil {
		convArg0 = target.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller8scrollerEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQScrollerFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QScroller_Scroller(target qtcore.QObject_ITF /*777 QObject **/) *QScroller /*777 QScroller **/ {
	var nilthis *QScroller
	rv := nilthis.Scroller(target)
	return rv
}

// /usr/include/qt/QtWidgets/qscroller.h:95
// index:1
// Public static Visibility=Default Availability=Available
// [8] const QScroller * scroller(const QObject *)

/*
Returns the scroller for the given target. As long as the object exists this function will always return the same QScroller instance. If no QScroller exists for the target, one will implicitly be created. At no point more than one QScroller will be active on an object.

See also hasScroller() and target().
*/
func (this *QScroller) Scroller1(target qtcore.QObject_ITF /*777 const QObject **/) *QScroller /*777 const QScroller **/ {
	var convArg0 unsafe.Pointer
	if target != nil && target.QObject_PTR() != nil {
		convArg0 = target.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller8scrollerEPK7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQScrollerFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QScroller_Scroller1(target qtcore.QObject_ITF /*777 const QObject **/) *QScroller /*777 const QScroller **/ {
	var nilthis *QScroller
	rv := nilthis.Scroller1(target)
	return rv
}

// /usr/include/qt/QtWidgets/qscroller.h:98
// index:0
// Public static Visibility=Default Availability=Available
// [4] Qt::GestureType grabGesture(QObject *, QScroller::ScrollerGestureType)

/*
Registers a custom scroll gesture recognizer, grabs it for the target and returns the resulting gesture type. If scrollGestureType is set to TouchGesture the gesture triggers on touch events. If it is set to one of LeftMouseButtonGesture, RightMouseButtonGesture or MiddleMouseButtonGesture it triggers on mouse events of the corresponding button.

Only one scroll gesture can be active on a single object at the same time. If you call this function twice on the same object, it will ungrab the existing gesture before grabbing the new one.

Note: To avoid unwanted side-effects, mouse events are consumed while the gesture is triggered. Since the initial mouse press event is not consumed, the gesture sends a fake mouse release event at the global position (INT_MIN, INT_MIN). This ensures that internal states of the widget that received the original mouse press are consistent.

See also ungrabGesture() and grabbedGesture().
*/
func (this *QScroller) GrabGesture(target qtcore.QObject_ITF /*777 QObject **/, gestureType int) int {
	var convArg0 unsafe.Pointer
	if target != nil && target.QObject_PTR() != nil {
		convArg0 = target.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller11grabGestureEP7QObjectNS_19ScrollerGestureTypeE", qtrt.FFI_TYPE_POINTER, convArg0, gestureType)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QScroller_GrabGesture(target qtcore.QObject_ITF /*777 QObject **/, gestureType int) int {
	var nilthis *QScroller
	rv := nilthis.GrabGesture(target, gestureType)
	return rv
}

// /usr/include/qt/QtWidgets/qscroller.h:98
// index:0
// Public static Visibility=Default Availability=Available
// [4] Qt::GestureType grabGesture(QObject *, QScroller::ScrollerGestureType)

/*
Registers a custom scroll gesture recognizer, grabs it for the target and returns the resulting gesture type. If scrollGestureType is set to TouchGesture the gesture triggers on touch events. If it is set to one of LeftMouseButtonGesture, RightMouseButtonGesture or MiddleMouseButtonGesture it triggers on mouse events of the corresponding button.

Only one scroll gesture can be active on a single object at the same time. If you call this function twice on the same object, it will ungrab the existing gesture before grabbing the new one.

Note: To avoid unwanted side-effects, mouse events are consumed while the gesture is triggered. Since the initial mouse press event is not consumed, the gesture sends a fake mouse release event at the global position (INT_MIN, INT_MIN). This ensures that internal states of the widget that received the original mouse press are consistent.

See also ungrabGesture() and grabbedGesture().
*/
func (this *QScroller) GrabGesturep(target qtcore.QObject_ITF /*777 QObject **/) int {
	var convArg0 unsafe.Pointer
	if target != nil && target.QObject_PTR() != nil {
		convArg0 = target.QObject_PTR().GetCthis()
	}
	// arg: 1, QScroller::ScrollerGestureType=Enum, QScroller::ScrollerGestureType=Enum, , Invalid
	gestureType := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller11grabGestureEP7QObjectNS_19ScrollerGestureTypeE", qtrt.FFI_TYPE_POINTER, convArg0, gestureType)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:99
// index:0
// Public static Visibility=Default Availability=Available
// [4] Qt::GestureType grabbedGesture(QObject *)

/*
Returns the gesture type currently grabbed for the target or 0 if no gesture is grabbed.

See also grabGesture() and ungrabGesture().
*/
func (this *QScroller) GrabbedGesture(target qtcore.QObject_ITF /*777 QObject **/) int {
	var convArg0 unsafe.Pointer
	if target != nil && target.QObject_PTR() != nil {
		convArg0 = target.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller14grabbedGestureEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QScroller_GrabbedGesture(target qtcore.QObject_ITF /*777 QObject **/) int {
	var nilthis *QScroller
	rv := nilthis.GrabbedGesture(target)
	return rv
}

// /usr/include/qt/QtWidgets/qscroller.h:100
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void ungrabGesture(QObject *)

/*
Ungrabs the gesture for the target. Does nothing if no gesture is grabbed.

See also grabGesture() and grabbedGesture().
*/
func (this *QScroller) UngrabGesture(target qtcore.QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if target != nil && target.QObject_PTR() != nil {
		convArg0 = target.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller13ungrabGestureEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QScroller_UngrabGesture(target qtcore.QObject_ITF /*777 QObject **/) {
	var nilthis *QScroller
	nilthis.UngrabGesture(target)
}

// /usr/include/qt/QtWidgets/qscroller.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * target() const

/*
Returns the target object of this scroller.

See also hasScroller() and scroller().
*/
func (this *QScroller) Target() *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QScroller6targetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qscroller.h:107
// index:0
// Public Visibility=Default Availability=Available
// [4] QScroller::State state() const

/*

 */
func (this *QScroller) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QScroller5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:109
// index:0
// Public Visibility=Default Availability=Available
// [1] bool handleInput(QScroller::Input, const QPointF &, qint64)

/*
This function is used by gesture recognizers to inform the scroller about a new input event. The scroller changes its internal state() according to the input event and its attached scroller properties. The scroller doesn't distinguish between the kind of input device the event came from. Therefore the event needs to be split into the input type, a position and a milli-second timestamp. The position needs to be in the target's coordinate system.

The return value is true if the event should be consumed by the calling filter or false if the event should be forwarded to the control.

Note: Using grabGesture() should be sufficient for most use cases.
*/
func (this *QScroller) HandleInput(input int, position qtcore.QPointF_ITF, timestamp int64) bool {
	var convArg1 unsafe.Pointer
	if position != nil && position.QPointF_PTR() != nil {
		convArg1 = position.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller11handleInputENS_5InputERK7QPointFx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), input, convArg1, timestamp)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qscroller.h:109
// index:0
// Public Visibility=Default Availability=Available
// [1] bool handleInput(QScroller::Input, const QPointF &, qint64)

/*
This function is used by gesture recognizers to inform the scroller about a new input event. The scroller changes its internal state() according to the input event and its attached scroller properties. The scroller doesn't distinguish between the kind of input device the event came from. Therefore the event needs to be split into the input type, a position and a milli-second timestamp. The position needs to be in the target's coordinate system.

The return value is true if the event should be consumed by the calling filter or false if the event should be forwarded to the control.

Note: Using grabGesture() should be sufficient for most use cases.
*/
func (this *QScroller) HandleInputp(input int, position qtcore.QPointF_ITF) bool {
	var convArg1 unsafe.Pointer
	if position != nil && position.QPointF_PTR() != nil {
		convArg1 = position.QPointF_PTR().GetCthis()
	}
	// arg: 2, qint64=Typedef, qint64=Typedef, long long, LongLong
	timestamp := int64(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller11handleInputENS_5InputERK7QPointFx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), input, convArg1, timestamp)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qscroller.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stop()

/*
Stops the scroller and resets its state back to Inactive.
*/
func (this *QScroller) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:112
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF velocity() const

/*
Returns the current scrolling velocity in meter per second when the state is Scrolling or Dragging. Returns a zero velocity otherwise.

The velocity is reported for both the x and y axis separately by using a QPointF.

See also pixelPerMeter().
*/
func (this *QScroller) Velocity() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QScroller8velocityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qscroller.h:113
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF finalPosition() const

/*
Returns the estimated final position for the current scroll movement. Returns the current position if the scroller state is not Scrolling. The result is undefined when the scroller state is Inactive.

The target position is in pixel.

See also pixelPerMeter() and scrollTo().
*/
func (this *QScroller) FinalPosition() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QScroller13finalPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qscroller.h:114
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF pixelPerMeter() const

/*
Returns the pixel per meter metric for the scrolled widget.

The value is reported for both the x and y axis separately by using a QPointF.

Note: Please note that this value should be physically correct. The actual DPI settings that Qt returns for the display may be reported wrongly on purpose by the underlying windowing system, for example on macOS.
*/
func (this *QScroller) PixelPerMeter() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QScroller13pixelPerMeterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qscroller.h:116
// index:0
// Public Visibility=Default Availability=Available
// [16] QScrollerProperties scrollerProperties() const

/*

 */
func (this *QScroller) ScrollerProperties() *QScrollerProperties /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QScroller18scrollerPropertiesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQScrollerPropertiesFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQScrollerProperties)
	return rv2
}

// /usr/include/qt/QtWidgets/qscroller.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSnapPositionsX(qreal, qreal)

/*
Set the snap positions for the horizontal axis to a list of positions. This overwrites all previously set snap positions and also a previously set snapping interval. Snapping can be deactivated by setting an empty list of positions.
*/
func (this *QScroller) SetSnapPositionsX(first float64, interval float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller17setSnapPositionsXEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), first, interval)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSnapPositionsY(qreal, qreal)

/*
Set the snap positions for the vertical axis to a list of positions. This overwrites all previously set snap positions and also a previously set snapping interval. Snapping can be deactivated by setting an empty list of positions.
*/
func (this *QScroller) SetSnapPositionsY(first float64, interval float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller17setSnapPositionsYEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), first, interval)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScrollerProperties(const QScrollerProperties &)

/*

 */
func (this *QScroller) SetScrollerProperties(prop QScrollerProperties_ITF) {
	var convArg0 unsafe.Pointer
	if prop != nil && prop.QScrollerProperties_PTR() != nil {
		convArg0 = prop.QScrollerProperties_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller21setScrollerPropertiesERK19QScrollerProperties", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scrollTo(const QPointF &)

/*
Starts scrolling the widget so that point pos is at the top-left position in the viewport.

The behaviour when scrolling outside the valid scroll area is undefined. In this case the scroller might or might not overshoot.

The scrolling speed will be calculated so that the given position will be reached after a platform-defined time span.

pos is given in viewport coordinates.

See also ensureVisible().
*/
func (this *QScroller) ScrollTo(pos qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller8scrollToERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:126
// index:1
// Public Visibility=Default Availability=Available
// [-2] void scrollTo(const QPointF &, int)

/*
Starts scrolling the widget so that point pos is at the top-left position in the viewport.

The behaviour when scrolling outside the valid scroll area is undefined. In this case the scroller might or might not overshoot.

The scrolling speed will be calculated so that the given position will be reached after a platform-defined time span.

pos is given in viewport coordinates.

See also ensureVisible().
*/
func (this *QScroller) ScrollTo1(pos qtcore.QPointF_ITF, scrollTime int) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller8scrollToERK7QPointFi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, scrollTime)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ensureVisible(const QRectF &, qreal, qreal)

/*
Starts scrolling so that the rectangle rect is visible inside the viewport with additional margins specified in pixels by xmargin and ymargin around the rect.

In cases where it is not possible to fit the rect plus margins inside the viewport the contents are scrolled so that as much as possible is visible from rect.

The scrolling speed is calculated so that the given position is reached after a platform-defined time span.

This function performs the actual scrolling by calling scrollTo().

See also scrollTo().
*/
func (this *QScroller) EnsureVisible(rect qtcore.QRectF_ITF, xmargin float64, ymargin float64) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller13ensureVisibleERK6QRectFdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:128
// index:1
// Public Visibility=Default Availability=Available
// [-2] void ensureVisible(const QRectF &, qreal, qreal, int)

/*
Starts scrolling so that the rectangle rect is visible inside the viewport with additional margins specified in pixels by xmargin and ymargin around the rect.

In cases where it is not possible to fit the rect plus margins inside the viewport the contents are scrolled so that as much as possible is visible from rect.

The scrolling speed is calculated so that the given position is reached after a platform-defined time span.

This function performs the actual scrolling by calling scrollTo().

See also scrollTo().
*/
func (this *QScroller) EnsureVisible1(rect qtcore.QRectF_ITF, xmargin float64, ymargin float64, scrollTime int) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller13ensureVisibleERK6QRectFddi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xmargin, ymargin, scrollTime)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resendPrepareEvent()

/*
This function resends the QScrollPrepareEvent. Calling resendPrepareEvent triggers a QScrollPrepareEvent from the scroller. This allows the receiver to re-set content position and content size while scrolling. Calling this function while in the Inactive state is useless as the prepare event is sent again before scrolling starts.
*/
func (this *QScroller) ResendPrepareEvent() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller18resendPrepareEventEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(QScroller::State)

/*
QScroller emits this signal whenever the state changes. newState is the new State.

Note: Notifier signal for property state.

See also state.
*/
func (this *QScroller) StateChanged(newstate int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller12stateChangedENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newstate)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscroller.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scrollerPropertiesChanged(const QScrollerProperties &)

/*
QScroller emits this signal whenever its scroller properties change. newProperties are the new scroller properties.

Note: Notifier signal for property scrollerProperties.

See also scrollerProperties.
*/
func (this *QScroller) ScrollerPropertiesChanged(arg0 QScrollerProperties_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QScrollerProperties_PTR() != nil {
		convArg0 = arg0.QScrollerProperties_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScroller25scrollerPropertiesChangedERK19QScrollerProperties", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

func DeleteQScroller(this *QScroller) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QScrollerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
This enum contains the different QScroller states.


*/
type QScroller__State = int

// The scroller is not scrolling and nothing is pressed.
const QScroller__Inactive QScroller__State = 0

// A touch event was received or the mouse button was pressed but the scroll area is currently not dragged.
const QScroller__Pressed QScroller__State = 1

// The scroll area is currently following the touch point or mouse.
const QScroller__Dragging QScroller__State = 2

// The scroll area is moving on it's own.
const QScroller__Scrolling QScroller__State = 3

func (this *QScroller) StateItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QScroller_StateItemName(val int) string {
	var nilthis *QScroller
	return nilthis.StateItemName(val)
}

/*
This enum contains the different gesture types that are supported by the QScroller gesture recognizer.


*/
type QScroller__ScrollerGestureType = int

// The gesture recognizer will only trigger on touch events. Specifically it will react on single touch points when using a touch screen and dual touch points when using a touchpad.
const QScroller__TouchGesture QScroller__ScrollerGestureType = 0

// The gesture recognizer will only trigger on left mouse button events.
const QScroller__LeftMouseButtonGesture QScroller__ScrollerGestureType = 1

// The gesture recognizer will only trigger on right mouse button events.
const QScroller__RightMouseButtonGesture QScroller__ScrollerGestureType = 2

// The gesture recognizer will only trigger on middle mouse button events.
const QScroller__MiddleMouseButtonGesture QScroller__ScrollerGestureType = 3

func (this *QScroller) ScrollerGestureTypeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QScroller_ScrollerGestureTypeItemName(val int) string {
	var nilthis *QScroller
	return nilthis.ScrollerGestureTypeItemName(val)
}

/*
This enum contains an input device agnostic view of input events that are relevant for QScroller.


*/
type QScroller__Input = int

// The user pressed the input device (e.g. QEvent::MouseButtonPress, QEvent::GraphicsSceneMousePress, QEvent::TouchBegin)
const QScroller__InputPress QScroller__Input = 1

// The user moved the input device (e.g. QEvent::MouseMove, QEvent::GraphicsSceneMouseMove, QEvent::TouchUpdate)
const QScroller__InputMove QScroller__Input = 2

// The user released the input device (e.g. QEvent::MouseButtonRelease, QEvent::GraphicsSceneMouseRelease, QEvent::TouchEnd)
const QScroller__InputRelease QScroller__Input = 3

func (this *QScroller) InputItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QScroller_InputItemName(val int) string {
	var nilthis *QScroller
	return nilthis.InputItemName(val)
}

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
