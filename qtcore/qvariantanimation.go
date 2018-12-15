package qtcore

// /usr/include/qt/QtCore/qvariantanimation.h
// #include <qvariantanimation.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 66
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// bool event(QEvent *)
func (this *QVariantAnimation) InheritEvent(f func(event *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void updateCurrentTime(int)
func (this *QVariantAnimation) InheritUpdateCurrentTime(f func(arg0 int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateCurrentTime", f)
}

// void updateState(QAbstractAnimation::State, QAbstractAnimation::State)
func (this *QVariantAnimation) InheritUpdateState(f func(newState int, oldState int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateState", f)
}

// void updateCurrentValue(const QVariant &)
func (this *QVariantAnimation) InheritUpdateCurrentValue(f func(value *QVariant) /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateCurrentValue", f)
}

// QVariant interpolated(const QVariant &, const QVariant &, qreal)
func (this *QVariantAnimation) InheritInterpolated(f func(from *QVariant, to *QVariant, progress float64) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "interpolated", f)
}

/*

 */
type QVariantAnimation struct {
	*QAbstractAnimation
}
type QVariantAnimation_ITF interface {
	QAbstractAnimation_ITF
	QVariantAnimation_PTR() *QVariantAnimation
}

func (ptr *QVariantAnimation) QVariantAnimation_PTR() *QVariantAnimation { return ptr }

func (this *QVariantAnimation) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractAnimation.GetCthis()
	}
}
func (this *QVariantAnimation) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractAnimation = NewQAbstractAnimationFromPointer(cthis)
}
func NewQVariantAnimationFromPointer(cthis unsafe.Pointer) *QVariantAnimation {
	bcthis0 := NewQAbstractAnimationFromPointer(cthis)
	return &QVariantAnimation{bcthis0}
}
func (*QVariantAnimation) NewFromPointer(cthis unsafe.Pointer) *QVariantAnimation {
	return NewQVariantAnimationFromPointer(cthis)
}

// /usr/include/qt/QtCore/qvariantanimation.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QVariantAnimation) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QVariantAnimation10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qvariantanimation.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QVariantAnimation(QObject *)

/*
Construct a QVariantAnimation object. parent is passed to QAbstractAnimation's constructor.
*/
func (*QVariantAnimation) NewForInherit(parent QObject_ITF /*777 QObject **/) *QVariantAnimation {
	return NewQVariantAnimation(parent)
}
func NewQVariantAnimation(parent QObject_ITF /*777 QObject **/) *QVariantAnimation {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QVariantAnimationC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantAnimationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QVariantAnimation")
	return gothis
}

// /usr/include/qt/QtCore/qvariantanimation.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QVariantAnimation(QObject *)

/*
Construct a QVariantAnimation object. parent is passed to QAbstractAnimation's constructor.
*/
func (*QVariantAnimation) NewForInheritp() *QVariantAnimation {
	return NewQVariantAnimationp()
}
func NewQVariantAnimationp() *QVariantAnimation {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN17QVariantAnimationC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantAnimationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QVariantAnimation")
	return gothis
}

// /usr/include/qt/QtCore/qvariantanimation.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QVariantAnimation()

/*

 */
func DeleteQVariantAnimation(this *QVariantAnimation) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QVariantAnimationD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qvariantanimation.h:71
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant startValue() const

/*

 */
func (this *QVariantAnimation) StartValue() *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QVariantAnimation10startValueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qvariantanimation.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStartValue(const QVariant &)

/*

 */
func (this *QVariantAnimation) SetStartValue(value QVariant_ITF) {
	var convArg0 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg0 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QVariantAnimation13setStartValueERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:74
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant endValue() const

/*

 */
func (this *QVariantAnimation) EndValue() *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QVariantAnimation8endValueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qvariantanimation.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEndValue(const QVariant &)

/*

 */
func (this *QVariantAnimation) SetEndValue(value QVariant_ITF) {
	var convArg0 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg0 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QVariantAnimation11setEndValueERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:77
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant keyValueAt(qreal) const

/*
Returns the key frame value for the given step. The given step must be in the range 0 to 1. If there is no KeyValue for step, it returns an invalid QVariant.

See also keyValues() and setKeyValueAt().
*/
func (this *QVariantAnimation) KeyValueAt(step float64) *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QVariantAnimation10keyValueAtEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), step)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qvariantanimation.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKeyValueAt(qreal, const QVariant &)

/*
Creates a key frame at the given step with the given value. The given step must be in the range 0 to 1.

See also setKeyValues() and keyValueAt().
*/
func (this *QVariantAnimation) SetKeyValueAt(step float64, value QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QVariantAnimation13setKeyValueAtEdRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), step, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:83
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant currentValue() const

/*

 */
func (this *QVariantAnimation) CurrentValue() *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QVariantAnimation12currentValueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qvariantanimation.h:85
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int duration() const

/*

 */
func (this *QVariantAnimation) Duration() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QVariantAnimation8durationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qvariantanimation.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDuration(int)

/*

 */
func (this *QVariantAnimation) SetDuration(msecs int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QVariantAnimation11setDurationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QEasingCurve easingCurve() const

/*

 */
func (this *QVariantAnimation) EasingCurve() *QEasingCurve /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QVariantAnimation11easingCurveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQEasingCurveFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQEasingCurve)
	return rv2
}

// /usr/include/qt/QtCore/qvariantanimation.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEasingCurve(const QEasingCurve &)

/*

 */
func (this *QVariantAnimation) SetEasingCurve(easing QEasingCurve_ITF) {
	var convArg0 unsafe.Pointer
	if easing != nil && easing.QEasingCurve_PTR() != nil {
		convArg0 = easing.QEasingCurve_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QVariantAnimation14setEasingCurveERK12QEasingCurve", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void valueChanged(const QVariant &)

/*
QVariantAnimation emits this signal whenever the current value changes.

Note: Notifier signal for property currentValue.

See also currentValue, startValue, and endValue.
*/
func (this *QVariantAnimation) ValueChanged(value QVariant_ITF) {
	var convArg0 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg0 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QVariantAnimation12valueChangedERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:98
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QVariantAnimation) Event(event QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QVariantAnimation5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qvariantanimation.h:100
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateCurrentTime(int)

/*
Reimplemented from QAbstractAnimation::updateCurrentTime().
*/
func (this *QVariantAnimation) UpdateCurrentTime(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QVariantAnimation17updateCurrentTimeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:101
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateState(QAbstractAnimation::State, QAbstractAnimation::State)

/*
Reimplemented from QAbstractAnimation::updateState().
*/
func (this *QVariantAnimation) UpdateState(newState int, oldState int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QVariantAnimation11updateStateEN18QAbstractAnimation5StateES1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newState, oldState)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:103
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateCurrentValue(const QVariant &)

/*
This virtual function is called every time the animation's current value changes. The value argument is the new current value.

The base class implementation does nothing.

See also currentValue.
*/
func (this *QVariantAnimation) UpdateCurrentValue(value QVariant_ITF) {
	var convArg0 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg0 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QVariantAnimation18updateCurrentValueERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:104
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QVariant interpolated(const QVariant &, const QVariant &, qreal) const

/*
This virtual function returns the linear interpolation between variants from and to, at progress, usually a value between 0 and 1. You can reimplement this function in a subclass of QVariantAnimation to provide your own interpolation algorithm.

Note that in order for the interpolation to work with a QEasingCurve that return a value smaller than 0 or larger than 1 (such as QEasingCurve::InBack) you should make sure that it can extrapolate. If the semantic of the datatype does not allow extrapolation this function should handle that gracefully.

You should call the QVariantAnimation implementation of this function if you want your class to handle the types already supported by Qt (see class QVariantAnimation description for a list of supported types).

See also QEasingCurve.
*/
func (this *QVariantAnimation) Interpolated(from QVariant_ITF, to QVariant_ITF, progress float64) *QVariant /*123*/ {
	var convArg0 unsafe.Pointer
	if from != nil && from.QVariant_PTR() != nil {
		convArg0 = from.QVariant_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if to != nil && to.QVariant_PTR() != nil {
		convArg1 = to.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QVariantAnimation12interpolatedERK8QVariantS2_d", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, progress)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
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
}

//  keep block end
