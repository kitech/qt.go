package qtwidgets

// /usr/include/qt/QtWidgets/qabstractslider.h
// #include <qabstractslider.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 58
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
// bool event(class QEvent *)
func (this *QAbstractSlider) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void setRepeatAction(enum QAbstractSlider::SliderAction, int, int)
func (this *QAbstractSlider) InheritSetRepeatAction(f func(action int, thresholdTime int, repeatTime int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setRepeatAction", f)
}

// QAbstractSlider::SliderAction repeatAction()
func (this *QAbstractSlider) InheritRepeatAction(f func() int) {
	qtrt.SetAllInheritCallback(this, "repeatAction", f)
}

// void sliderChange(enum QAbstractSlider::SliderChange)
func (this *QAbstractSlider) InheritSliderChange(f func(change int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "sliderChange", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QAbstractSlider) InheritKeyPressEvent(f func(ev *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void timerEvent(class QTimerEvent *)
func (this *QAbstractSlider) InheritTimerEvent(f func(arg0 *qtcore.QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

// void wheelEvent(class QWheelEvent *)
func (this *QAbstractSlider) InheritWheelEvent(f func(e *qtgui.QWheelEvent /*777 QWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void changeEvent(class QEvent *)
func (this *QAbstractSlider) InheritChangeEvent(f func(e *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

type QAbstractSlider struct {
	*QWidget
}

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
// [8] const QMetaObject * metaObject()
func (this *QAbstractSlider) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSlider10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractslider.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractSlider(QWidget *)
func NewQAbstractSlider(parent *QWidget /*777 QWidget **/) *QAbstractSlider {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSliderC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractSliderFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractslider.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractSlider()
func DeleteQAbstractSlider(this *QAbstractSlider) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSliderD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:73
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Orientation orientation()
func (this *QAbstractSlider) Orientation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSlider11orientationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimum(int)
func (this *QAbstractSlider) SetMinimum(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider10setMinimumEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] int minimum()
func (this *QAbstractSlider) Minimum() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSlider7minimumEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractslider.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximum(int)
func (this *QAbstractSlider) SetMaximum(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider10setMaximumEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:79
// index:0
// Public Visibility=Default Availability=Available
// [4] int maximum()
func (this *QAbstractSlider) Maximum() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSlider7maximumEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractslider.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSingleStep(int)
func (this *QAbstractSlider) SetSingleStep(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider13setSingleStepEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] int singleStep()
func (this *QAbstractSlider) SingleStep() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSlider10singleStepEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractslider.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPageStep(int)
func (this *QAbstractSlider) SetPageStep(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider11setPageStepEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:85
// index:0
// Public Visibility=Default Availability=Available
// [4] int pageStep()
func (this *QAbstractSlider) PageStep() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSlider8pageStepEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractslider.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTracking(_Bool)
func (this *QAbstractSlider) SetTracking(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider11setTrackingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:88
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasTracking()
func (this *QAbstractSlider) HasTracking() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSlider11hasTrackingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractslider.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSliderDown(_Bool)
func (this *QAbstractSlider) SetSliderDown(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider13setSliderDownEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:91
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSliderDown()
func (this *QAbstractSlider) IsSliderDown() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSlider12isSliderDownEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractslider.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSliderPosition(int)
func (this *QAbstractSlider) SetSliderPosition(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider17setSliderPositionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:94
// index:0
// Public Visibility=Default Availability=Available
// [4] int sliderPosition()
func (this *QAbstractSlider) SliderPosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSlider14sliderPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractslider.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInvertedAppearance(_Bool)
func (this *QAbstractSlider) SetInvertedAppearance(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider21setInvertedAppearanceEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:97
// index:0
// Public Visibility=Default Availability=Available
// [1] bool invertedAppearance()
func (this *QAbstractSlider) InvertedAppearance() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSlider18invertedAppearanceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractslider.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInvertedControls(_Bool)
func (this *QAbstractSlider) SetInvertedControls(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider19setInvertedControlsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:100
// index:0
// Public Visibility=Default Availability=Available
// [1] bool invertedControls()
func (this *QAbstractSlider) InvertedControls() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSlider16invertedControlsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractslider.h:113
// index:0
// Public Visibility=Default Availability=Available
// [4] int value()
func (this *QAbstractSlider) Value() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSlider5valueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractslider.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void triggerAction(enum QAbstractSlider::SliderAction)
func (this *QAbstractSlider) TriggerAction(action int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider13triggerActionENS_12SliderActionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setValue(int)
func (this *QAbstractSlider) SetValue(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider8setValueEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOrientation(Qt::Orientation)
func (this *QAbstractSlider) SetOrientation(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider14setOrientationEN2Qt11OrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRange(int, int)
func (this *QAbstractSlider) SetRange(min int, max int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider8setRangeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), min, max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void valueChanged(int)
func (this *QAbstractSlider) ValueChanged(value int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider12valueChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sliderPressed()
func (this *QAbstractSlider) SliderPressed() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider13sliderPressedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sliderMoved(int)
func (this *QAbstractSlider) SliderMoved(position int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider11sliderMovedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sliderReleased()
func (this *QAbstractSlider) SliderReleased() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider14sliderReleasedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void rangeChanged(int, int)
func (this *QAbstractSlider) RangeChanged(min int, max int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider12rangeChangedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), min, max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void actionTriggered(int)
func (this *QAbstractSlider) ActionTriggered(action int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider15actionTriggeredEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:134
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QAbstractSlider) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractslider.h:136
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setRepeatAction(enum QAbstractSlider::SliderAction, int, int)
func (this *QAbstractSlider) SetRepeatAction(action int, thresholdTime int, repeatTime int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider15setRepeatActionENS_12SliderActionEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action, thresholdTime, repeatTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:137
// index:0
// Protected Visibility=Default Availability=Available
// [4] QAbstractSlider::SliderAction repeatAction()
func (this *QAbstractSlider) RepeatAction() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSlider12repeatActionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:145
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void sliderChange(enum QAbstractSlider::SliderChange)
func (this *QAbstractSlider) SliderChange(change int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider12sliderChangeENS_12SliderChangeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), change)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:147
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QAbstractSlider) KeyPressEvent(ev *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = ev.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:148
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)
func (this *QAbstractSlider) TimerEvent(arg0 *qtcore.QTimerEvent /*777 QTimerEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:150
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)
func (this *QAbstractSlider) WheelEvent(e *qtgui.QWheelEvent /*777 QWheelEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:152
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)
func (this *QAbstractSlider) ChangeEvent(e *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSlider11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QAbstractSlider__SliderAction = int

const QAbstractSlider__SliderNoAction QAbstractSlider__SliderAction = 0
const QAbstractSlider__SliderSingleStepAdd QAbstractSlider__SliderAction = 1
const QAbstractSlider__SliderSingleStepSub QAbstractSlider__SliderAction = 2
const QAbstractSlider__SliderPageStepAdd QAbstractSlider__SliderAction = 3
const QAbstractSlider__SliderPageStepSub QAbstractSlider__SliderAction = 4
const QAbstractSlider__SliderToMinimum QAbstractSlider__SliderAction = 5
const QAbstractSlider__SliderToMaximum QAbstractSlider__SliderAction = 6
const QAbstractSlider__SliderMove QAbstractSlider__SliderAction = 7

type QAbstractSlider__SliderChange = int

const QAbstractSlider__SliderRangeChange QAbstractSlider__SliderChange = 0
const QAbstractSlider__SliderOrientationChange QAbstractSlider__SliderChange = 1
const QAbstractSlider__SliderStepsChange QAbstractSlider__SliderChange = 2
const QAbstractSlider__SliderValueChange QAbstractSlider__SliderChange = 3

//  body block end
