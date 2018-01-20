//  header block begin
// /usr/include/qt/QtWidgets/qabstractslider.h
// #include <qabstractslider.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 59
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
type QAbstractSlider struct {
	*QWidget
}

func (this *QAbstractSlider) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}

// /usr/include/qt/QtWidgets/qabstractslider.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QAbstractSlider) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSlider10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:70
// index:0
// void QAbstractSlider(class QWidget *)
func NewQAbstractSlider(parent unsafe.Pointer) *QAbstractSlider {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSliderC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractSliderFromPointer(cthis)
	return gothis
}
func NewQAbstractSliderFromPointer(cthis unsafe.Pointer) *QAbstractSlider {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QAbstractSlider{bcthis0}
}

// /usr/include/qt/QtWidgets/qabstractslider.h:156
// index:1
// void QAbstractSlider(class QAbstractSliderPrivate &, class QWidget *)
func NewQAbstractSlider_1(dd unsafe.Pointer, parent unsafe.Pointer) *QAbstractSlider {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSliderC2ER22QAbstractSliderPrivateP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractSliderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractslider.h:71
// index:0
// virtual
// void ~QAbstractSlider()
func DeleteQAbstractSlider(*QAbstractSlider) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSliderD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:73
// index:0
// Qt::Orientation orientation()
func (this *QAbstractSlider) Orientation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSlider11orientationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:75
// index:0
// void setMinimum(int)
func (this *QAbstractSlider) SetMinimum(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSlider10setMinimumEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:76
// index:0
// int minimum()
func (this *QAbstractSlider) Minimum() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSlider7minimumEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:78
// index:0
// void setMaximum(int)
func (this *QAbstractSlider) SetMaximum(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSlider10setMaximumEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:79
// index:0
// int maximum()
func (this *QAbstractSlider) Maximum() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSlider7maximumEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:81
// index:0
// void setSingleStep(int)
func (this *QAbstractSlider) SetSingleStep(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSlider13setSingleStepEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:82
// index:0
// int singleStep()
func (this *QAbstractSlider) SingleStep() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSlider10singleStepEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:84
// index:0
// void setPageStep(int)
func (this *QAbstractSlider) SetPageStep(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSlider11setPageStepEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:85
// index:0
// int pageStep()
func (this *QAbstractSlider) PageStep() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSlider8pageStepEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:87
// index:0
// void setTracking(_Bool)
func (this *QAbstractSlider) SetTracking(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSlider11setTrackingEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:88
// index:0
// bool hasTracking()
func (this *QAbstractSlider) HasTracking() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSlider11hasTrackingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:90
// index:0
// void setSliderDown(_Bool)
func (this *QAbstractSlider) SetSliderDown(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSlider13setSliderDownEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:91
// index:0
// bool isSliderDown()
func (this *QAbstractSlider) IsSliderDown() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSlider12isSliderDownEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:93
// index:0
// void setSliderPosition(int)
func (this *QAbstractSlider) SetSliderPosition(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSlider17setSliderPositionEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:94
// index:0
// int sliderPosition()
func (this *QAbstractSlider) SliderPosition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSlider14sliderPositionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:96
// index:0
// void setInvertedAppearance(_Bool)
func (this *QAbstractSlider) SetInvertedAppearance(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSlider21setInvertedAppearanceEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:97
// index:0
// bool invertedAppearance()
func (this *QAbstractSlider) InvertedAppearance() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSlider18invertedAppearanceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:99
// index:0
// void setInvertedControls(_Bool)
func (this *QAbstractSlider) SetInvertedControls(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSlider19setInvertedControlsEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:100
// index:0
// bool invertedControls()
func (this *QAbstractSlider) InvertedControls() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSlider16invertedControlsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:113
// index:0
// int value()
func (this *QAbstractSlider) Value() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSlider5valueEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:115
// index:0
// void triggerAction(enum QAbstractSlider::SliderAction)
func (this *QAbstractSlider) TriggerAction(action int) {
	// 0: (, action QAbstractSlider::SliderAction), (&action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSlider13triggerActionENS_12SliderActionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:118
// index:0
// void setValue(int)
func (this *QAbstractSlider) SetValue(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSlider8setValueEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:119
// index:0
// void setOrientation(Qt::Orientation)
func (this *QAbstractSlider) SetOrientation(arg0 int) {
	// 0: (, Qt::Orientation arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSlider14setOrientationEN2Qt11OrientationE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:120
// index:0
// void setRange(int, int)
func (this *QAbstractSlider) SetRange(min int, max int) {
	// 0: (, min int, max int), (&min, &max)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSlider8setRangeEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &min, &max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:123
// index:0
// void valueChanged(int)
func (this *QAbstractSlider) ValueChanged(value int) {
	// 0: (, value int), (&value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSlider12valueChangedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:125
// index:0
// void sliderPressed()
func (this *QAbstractSlider) SliderPressed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSlider13sliderPressedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:126
// index:0
// void sliderMoved(int)
func (this *QAbstractSlider) SliderMoved(position int) {
	// 0: (, position int), (&position)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSlider11sliderMovedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:127
// index:0
// void sliderReleased()
func (this *QAbstractSlider) SliderReleased() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSlider14sliderReleasedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:129
// index:0
// void rangeChanged(int, int)
func (this *QAbstractSlider) RangeChanged(min int, max int) {
	// 0: (, min int, max int), (&min, &max)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSlider12rangeChangedEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &min, &max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:131
// index:0
// void actionTriggered(int)
func (this *QAbstractSlider) ActionTriggered(action int) {
	// 0: (, action int), (&action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSlider15actionTriggeredEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:134
// index:0
// virtual
// bool event(class QEvent *)
func (this *QAbstractSlider) Event(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSlider5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:136
// index:0
// void setRepeatAction(enum QAbstractSlider::SliderAction, int, int)
func (this *QAbstractSlider) SetRepeatAction(action int, thresholdTime int, repeatTime int) {
	// 0: (, action QAbstractSlider::SliderAction, thresholdTime int, repeatTime int), (&action, &thresholdTime, &repeatTime)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSlider15setRepeatActionENS_12SliderActionEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &action, &thresholdTime, &repeatTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:137
// index:0
// QAbstractSlider::SliderAction repeatAction()
func (this *QAbstractSlider) RepeatAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSlider12repeatActionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:145
// index:0
// virtual
// void sliderChange(enum QAbstractSlider::SliderChange)
func (this *QAbstractSlider) SliderChange(change int) {
	// 0: (, change QAbstractSlider::SliderChange), (&change)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSlider12sliderChangeENS_12SliderChangeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &change)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:147
// index:0
// virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QAbstractSlider) KeyPressEvent(ev unsafe.Pointer) {
	// 0: (, ev QKeyEvent *), (ev)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSlider13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ev)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:148
// index:0
// virtual
// void timerEvent(class QTimerEvent *)
func (this *QAbstractSlider) TimerEvent(arg0 unsafe.Pointer) {
	// 0: (, QTimerEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSlider10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:150
// index:0
// virtual
// void wheelEvent(class QWheelEvent *)
func (this *QAbstractSlider) WheelEvent(e unsafe.Pointer) {
	// 0: (, e QWheelEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSlider10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractslider.h:152
// index:0
// virtual
// void changeEvent(class QEvent *)
func (this *QAbstractSlider) ChangeEvent(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSlider11changeEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

//  body block end
