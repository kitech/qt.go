//  header block begin
// /usr/include/qt/QtCore/qvariantanimation.h
// #include <qvariantanimation.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 67
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
}

//  ext block end

//  body block begin
type QVariantAnimation struct {
	*QAbstractAnimation
}

func (this *QVariantAnimation) GetCthis() unsafe.Pointer {
	return this.QAbstractAnimation.GetCthis()
}

// /usr/include/qt/QtCore/qvariantanimation.h:57
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QVariantAnimation) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QVariantAnimation10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:68
// index:0
// void QVariantAnimation(class QObject *)
func NewQVariantAnimation(parent unsafe.Pointer) *QVariantAnimation {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QVariantAnimationC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantAnimationFromPointer(cthis)
	return gothis
}
func NewQVariantAnimationFromPointer(cthis unsafe.Pointer) *QVariantAnimation {
	bcthis0 := NewQAbstractAnimationFromPointer(cthis)
	return &QVariantAnimation{bcthis0}
}

// /usr/include/qt/QtCore/qvariantanimation.h:97
// index:1
// void QVariantAnimation(class QVariantAnimationPrivate &, class QObject *)
func NewQVariantAnimation_1(dd unsafe.Pointer, parent unsafe.Pointer) *QVariantAnimation {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QVariantAnimationC2ER24QVariantAnimationPrivateP7QObject", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantAnimationFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariantanimation.h:69
// index:0
// virtual
// void ~QVariantAnimation()
func DeleteQVariantAnimation(*QVariantAnimation) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QVariantAnimationD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:71
// index:0
// QVariant startValue()
func (this *QVariantAnimation) StartValue() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QVariantAnimation10startValueEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:72
// index:0
// void setStartValue(const class QVariant &)
func (this *QVariantAnimation) SetStartValue(value unsafe.Pointer) {
	// 0: (, value const QVariant &), (value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QVariantAnimation13setStartValueERK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:74
// index:0
// QVariant endValue()
func (this *QVariantAnimation) EndValue() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QVariantAnimation8endValueEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:75
// index:0
// void setEndValue(const class QVariant &)
func (this *QVariantAnimation) SetEndValue(value unsafe.Pointer) {
	// 0: (, value const QVariant &), (value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QVariantAnimation11setEndValueERK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:77
// index:0
// QVariant keyValueAt(qreal)
func (this *QVariantAnimation) KeyValueAt(step float64) {
	// 0: (, step qreal), (&step)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QVariantAnimation10keyValueAtEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &step)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:78
// index:0
// void setKeyValueAt(qreal, const class QVariant &)
func (this *QVariantAnimation) SetKeyValueAt(step float64, value unsafe.Pointer) {
	// 0: (, step qreal, value const QVariant &), (&step, value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QVariantAnimation13setKeyValueAtEdRK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &step, value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:80
// index:0
// QVariantAnimation::KeyValues keyValues()
func (this *QVariantAnimation) KeyValues() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QVariantAnimation9keyValuesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:83
// index:0
// QVariant currentValue()
func (this *QVariantAnimation) CurrentValue() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QVariantAnimation12currentValueEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:85
// index:0
// virtual
// int duration()
func (this *QVariantAnimation) Duration() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QVariantAnimation8durationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:86
// index:0
// void setDuration(int)
func (this *QVariantAnimation) SetDuration(msecs int) {
	// 0: (, msecs int), (&msecs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QVariantAnimation11setDurationEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:88
// index:0
// QEasingCurve easingCurve()
func (this *QVariantAnimation) EasingCurve() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QVariantAnimation11easingCurveEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:89
// index:0
// void setEasingCurve(const class QEasingCurve &)
func (this *QVariantAnimation) SetEasingCurve(easing unsafe.Pointer) {
	// 0: (, easing const QEasingCurve &), (easing)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QVariantAnimation14setEasingCurveERK12QEasingCurve", ffiqt.FFI_TYPE_VOID, this.GetCthis(), easing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:94
// index:0
// void valueChanged(const class QVariant &)
func (this *QVariantAnimation) ValueChanged(value unsafe.Pointer) {
	// 0: (, value const QVariant &), (value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QVariantAnimation12valueChangedERK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:98
// index:0
// virtual
// bool event(class QEvent *)
func (this *QVariantAnimation) Event(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QVariantAnimation5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:100
// index:0
// virtual
// void updateCurrentTime(int)
func (this *QVariantAnimation) UpdateCurrentTime(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QVariantAnimation17updateCurrentTimeEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:101
// index:0
// virtual
// void updateState(class QAbstractAnimation::State, class QAbstractAnimation::State)
func (this *QVariantAnimation) UpdateState(newState int, oldState int) {
	// 0: (, newState QAbstractAnimation::State, oldState QAbstractAnimation::State), (&newState, &oldState)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QVariantAnimation11updateStateEN18QAbstractAnimation5StateES1_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &newState, &oldState)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:103
// index:0
// virtual
// void updateCurrentValue(const class QVariant &)
func (this *QVariantAnimation) UpdateCurrentValue(value unsafe.Pointer) {
	// 0: (, value const QVariant &), (value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QVariantAnimation18updateCurrentValueERK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:104
// index:0
// virtual
// QVariant interpolated(const class QVariant &, const class QVariant &, qreal)
func (this *QVariantAnimation) Interpolated(from unsafe.Pointer, to unsafe.Pointer, progress float64) {
	// 0: (, from const QVariant &, to const QVariant &, progress qreal), (from, to, &progress)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QVariantAnimation12interpolatedERK8QVariantS2_d", ffiqt.FFI_TYPE_VOID, this.GetCthis(), from, to, &progress)
	gopp.ErrPrint(err, rv)
}

//  body block end
