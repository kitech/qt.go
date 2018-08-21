package qtwidgets

// /usr/include/qt/QtWidgets/qabstractslider.h
// #include <qabstractslider.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 58
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

// bool event(QEvent *)
func (this *QAbstractSlider) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void setRepeatAction(QAbstractSlider::SliderAction, int, int)
func (this *QAbstractSlider) InheritSetRepeatAction(f func(action int, thresholdTime int, repeatTime int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setRepeatAction", f)
}

// QAbstractSlider::SliderAction repeatAction()
func (this *QAbstractSlider) InheritRepeatAction(f func() int) {
	qtrt.SetAllInheritCallback(this, "repeatAction", f)
}

// void sliderChange(QAbstractSlider::SliderChange)
func (this *QAbstractSlider) InheritSliderChange(f func(change int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "sliderChange", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QAbstractSlider) InheritKeyPressEvent(f func(ev *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void timerEvent(QTimerEvent *)
func (this *QAbstractSlider) InheritTimerEvent(f func(arg0 *qtcore.QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

// void wheelEvent(QWheelEvent *)
func (this *QAbstractSlider) InheritWheelEvent(f func(e *qtgui.QWheelEvent /*777 QWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void changeEvent(QEvent *)
func (this *QAbstractSlider) InheritChangeEvent(f func(e *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

/*

 */
type QAbstractSlider struct {
	*QWidget
}
type QAbstractSlider_ITF interface {
	QWidget_ITF
	QAbstractSlider_PTR() *QAbstractSlider
}

func (ptr *QAbstractSlider) QAbstractSlider_PTR() *QAbstractSlider { return ptr }

func (this *QAbstractSlider) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QAbstractSlider) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQAbstractSliderFromPointer(cthis unsafe.Pointer) *QAbstractSlider {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QAbstractSlider{bcthis0}
}
func (*QAbstractSlider) NewFromPointer(cthis unsafe.Pointer) *QAbstractSlider {
	return NewQAbstractSliderFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAbstractSlider) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSlider10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractslider.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractSlider(QWidget *)

/*
Constructs an abstract slider.

The parent argument is sent to the QWidget constructor.

The minimum defaults to 0, the maximum to 99, with a singleStep size of 1 and a pageStep size of 10, and an initial value of 0.
*/
func NewQAbstractSlider(parent QWidget_ITF /*777 QWidget **/) *QAbstractSlider {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSliderC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractSliderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractSlider")
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractslider.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractSlider(QWidget *)

/*
Constructs an abstract slider.

The parent argument is sent to the QWidget constructor.

The minimum defaults to 0, the maximum to 99, with a singleStep size of 1 and a pageStep size of 10, and an initial value of 0.
*/
func NewQAbstractSlider__() *QAbstractSlider {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSliderC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractSliderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractSlider")
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractslider.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractSlider()

/*

 */
func DeleteQAbstractSlider(this *QAbstractSlider) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSliderD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:73
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Orientation orientation() const

/*

 */
func (this *QAbstractSlider) Orientation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSlider11orientationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimum(int)

/*

 */
func (this *QAbstractSlider) SetMinimum(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider10setMinimumEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] int minimum() const

/*

 */
func (this *QAbstractSlider) Minimum() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSlider7minimumEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractslider.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximum(int)

/*

 */
func (this *QAbstractSlider) SetMaximum(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider10setMaximumEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:79
// index:0
// Public Visibility=Default Availability=Available
// [4] int maximum() const

/*

 */
func (this *QAbstractSlider) Maximum() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSlider7maximumEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractslider.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSingleStep(int)

/*

 */
func (this *QAbstractSlider) SetSingleStep(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider13setSingleStepEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] int singleStep() const

/*

 */
func (this *QAbstractSlider) SingleStep() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSlider10singleStepEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractslider.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPageStep(int)

/*

 */
func (this *QAbstractSlider) SetPageStep(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider11setPageStepEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:85
// index:0
// Public Visibility=Default Availability=Available
// [4] int pageStep() const

/*

 */
func (this *QAbstractSlider) PageStep() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSlider8pageStepEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractslider.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTracking(bool)

/*

 */
func (this *QAbstractSlider) SetTracking(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider11setTrackingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:88
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasTracking() const

/*

 */
func (this *QAbstractSlider) HasTracking() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSlider11hasTrackingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractslider.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSliderDown(bool)

/*

 */
func (this *QAbstractSlider) SetSliderDown(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider13setSliderDownEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:91
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSliderDown() const

/*

 */
func (this *QAbstractSlider) IsSliderDown() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSlider12isSliderDownEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractslider.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSliderPosition(int)

/*

 */
func (this *QAbstractSlider) SetSliderPosition(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider17setSliderPositionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:94
// index:0
// Public Visibility=Default Availability=Available
// [4] int sliderPosition() const

/*

 */
func (this *QAbstractSlider) SliderPosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSlider14sliderPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractslider.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInvertedAppearance(bool)

/*

 */
func (this *QAbstractSlider) SetInvertedAppearance(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider21setInvertedAppearanceEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:97
// index:0
// Public Visibility=Default Availability=Available
// [1] bool invertedAppearance() const

/*

 */
func (this *QAbstractSlider) InvertedAppearance() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSlider18invertedAppearanceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractslider.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInvertedControls(bool)

/*

 */
func (this *QAbstractSlider) SetInvertedControls(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider19setInvertedControlsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:100
// index:0
// Public Visibility=Default Availability=Available
// [1] bool invertedControls() const

/*

 */
func (this *QAbstractSlider) InvertedControls() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSlider16invertedControlsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractslider.h:113
// index:0
// Public Visibility=Default Availability=Available
// [4] int value() const

/*

 */
func (this *QAbstractSlider) Value() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSlider5valueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractslider.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void triggerAction(QAbstractSlider::SliderAction)

/*
Triggers a slider action. Possible actions are SliderSingleStepAdd, SliderSingleStepSub, SliderPageStepAdd, SliderPageStepSub, SliderToMinimum, SliderToMaximum, and SliderMove.

See also actionTriggered().
*/
func (this *QAbstractSlider) TriggerAction(action int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider13triggerActionENS_12SliderActionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setValue(int)

/*

 */
func (this *QAbstractSlider) SetValue(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider8setValueEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOrientation(Qt::Orientation)

/*

 */
func (this *QAbstractSlider) SetOrientation(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider14setOrientationEN2Qt11OrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRange(int, int)

/*
Sets the slider's minimum to min and its maximum to max.

If max is smaller than min, min becomes the only legal value.

See also minimum and maximum.
*/
func (this *QAbstractSlider) SetRange(min int, max int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider8setRangeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), min, max)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void valueChanged(int)

/*
This signal is emitted when the slider value has changed, with the new slider value as argument.

Note: Notifier signal for property value.
*/
func (this *QAbstractSlider) ValueChanged(value int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider12valueChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sliderPressed()

/*
This signal is emitted when the user presses the slider with the mouse, or programmatically when setSliderDown(true) is called.

See also sliderReleased(), sliderMoved(), and isSliderDown().
*/
func (this *QAbstractSlider) SliderPressed() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider13sliderPressedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sliderMoved(int)

/*
This signal is emitted when sliderDown is true and the slider moves. This usually happens when the user is dragging the slider. The value is the new slider position.

This signal is emitted even when tracking is turned off.

Note: Notifier signal for property sliderPosition.

See also setTracking(), valueChanged(), isSliderDown(), sliderPressed(), and sliderReleased().
*/
func (this *QAbstractSlider) SliderMoved(position int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider11sliderMovedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sliderReleased()

/*
This signal is emitted when the user releases the slider with the mouse, or programmatically when setSliderDown(false) is called.

See also sliderPressed(), sliderMoved(), and sliderDown.
*/
func (this *QAbstractSlider) SliderReleased() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider14sliderReleasedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void rangeChanged(int, int)

/*
This signal is emitted when the slider range has changed, with min being the new minimum, and max being the new maximum.

See also minimum and maximum.
*/
func (this *QAbstractSlider) RangeChanged(min int, max int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider12rangeChangedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), min, max)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void actionTriggered(int)

/*
This signal is emitted when the slider action action is triggered. Actions are SliderSingleStepAdd, SliderSingleStepSub, SliderPageStepAdd, SliderPageStepSub, SliderToMinimum, SliderToMaximum, and SliderMove.

When the signal is emitted, the sliderPosition has been adjusted according to the action, but the value has not yet been propagated (meaning the valueChanged() signal was not yet emitted), and the visual display has not been updated. In slots connected to this signal you can thus safely adjust any action by calling setSliderPosition() yourself, based on both the action and the slider's value.

See also triggerAction().
*/
func (this *QAbstractSlider) ActionTriggered(action int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider15actionTriggeredEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:134
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QAbstractSlider) Event(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractslider.h:136
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setRepeatAction(QAbstractSlider::SliderAction, int, int)

/*
Sets action action to be triggered repetitively in intervals of repeatTime, after an initial delay of thresholdTime.

See also triggerAction() and repeatAction().
*/
func (this *QAbstractSlider) SetRepeatAction(action int, thresholdTime int, repeatTime int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider15setRepeatActionENS_12SliderActionEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action, thresholdTime, repeatTime)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:136
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setRepeatAction(QAbstractSlider::SliderAction, int, int)

/*
Sets action action to be triggered repetitively in intervals of repeatTime, after an initial delay of thresholdTime.

See also triggerAction() and repeatAction().
*/
func (this *QAbstractSlider) SetRepeatAction__(action int) {
	// arg: 1, int=Int, =Invalid, , Invalid
	thresholdTime := int(500)
	// arg: 2, int=Int, =Invalid, , Invalid
	repeatTime := int(50)
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider15setRepeatActionENS_12SliderActionEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action, thresholdTime, repeatTime)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:136
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setRepeatAction(QAbstractSlider::SliderAction, int, int)

/*
Sets action action to be triggered repetitively in intervals of repeatTime, after an initial delay of thresholdTime.

See also triggerAction() and repeatAction().
*/
func (this *QAbstractSlider) SetRepeatAction__1(action int, thresholdTime int) {
	// arg: 2, int=Int, =Invalid, , Invalid
	repeatTime := int(50)
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider15setRepeatActionENS_12SliderActionEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action, thresholdTime, repeatTime)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:137
// index:0
// Protected Visibility=Default Availability=Available
// [4] QAbstractSlider::SliderAction repeatAction() const

/*
Returns the current repeat action.

See also setRepeatAction().
*/
func (this *QAbstractSlider) RepeatAction() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSlider12repeatActionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:145
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void sliderChange(QAbstractSlider::SliderChange)

/*
Reimplement this virtual function to track slider changes such as SliderRangeChange, SliderOrientationChange, SliderStepsChange, or SliderValueChange. The default implementation only updates the display and ignores the change parameter.
*/
func (this *QAbstractSlider) SliderChange(change int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider12sliderChangeENS_12SliderChangeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), change)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:147
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
Reimplemented from QWidget::keyPressEvent().
*/
func (this *QAbstractSlider) KeyPressEvent(ev qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if ev != nil && ev.QKeyEvent_PTR() != nil {
		convArg0 = ev.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:148
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)

/*
Reimplemented from QObject::timerEvent().
*/
func (this *QAbstractSlider) TimerEvent(arg0 qtcore.QTimerEvent_ITF /*777 QTimerEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTimerEvent_PTR() != nil {
		convArg0 = arg0.QTimerEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:150
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)

/*
Reimplemented from QWidget::wheelEvent().
*/
func (this *QAbstractSlider) WheelEvent(e qtgui.QWheelEvent_ITF /*777 QWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QWheelEvent_PTR() != nil {
		convArg0 = e.QWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:152
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)

/*
Reimplemented from QWidget::changeEvent().
*/
func (this *QAbstractSlider) ChangeEvent(e qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*
ConstantValue
QAbstractSlider::SliderNoAction0
QAbstractSlider::SliderSingleStepAdd1
QAbstractSlider::SliderSingleStepSub2
QAbstractSlider::SliderPageStepAdd3
QAbstractSlider::SliderPageStepSub4
QAbstractSlider::SliderToMinimum5
QAbstractSlider::SliderToMaximum6
QAbstractSlider::SliderMove7

*/
type QAbstractSlider__SliderAction = int

//
const QAbstractSlider__SliderNoAction QAbstractSlider__SliderAction = 0

//
const QAbstractSlider__SliderSingleStepAdd QAbstractSlider__SliderAction = 1

//
const QAbstractSlider__SliderSingleStepSub QAbstractSlider__SliderAction = 2

//
const QAbstractSlider__SliderPageStepAdd QAbstractSlider__SliderAction = 3

//
const QAbstractSlider__SliderPageStepSub QAbstractSlider__SliderAction = 4

//
const QAbstractSlider__SliderToMinimum QAbstractSlider__SliderAction = 5

//
const QAbstractSlider__SliderToMaximum QAbstractSlider__SliderAction = 6

//
const QAbstractSlider__SliderMove QAbstractSlider__SliderAction = 7

func (this *QAbstractSlider) SliderActionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractSlider_SliderActionItemName(val int) string {
	var nilthis *QAbstractSlider
	return nilthis.SliderActionItemName(val)
}

/*
ConstantValue
QAbstractSlider::SliderRangeChange0
QAbstractSlider::SliderOrientationChange1
QAbstractSlider::SliderStepsChange2
QAbstractSlider::SliderValueChange3

*/
type QAbstractSlider__SliderChange = int

//
const QAbstractSlider__SliderRangeChange QAbstractSlider__SliderChange = 0

//
const QAbstractSlider__SliderOrientationChange QAbstractSlider__SliderChange = 1

//
const QAbstractSlider__SliderStepsChange QAbstractSlider__SliderChange = 2

//
const QAbstractSlider__SliderValueChange QAbstractSlider__SliderChange = 3

func (this *QAbstractSlider) SliderChangeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractSlider_SliderChangeItemName(val int) string {
	var nilthis *QAbstractSlider
	return nilthis.SliderChangeItemName(val)
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
