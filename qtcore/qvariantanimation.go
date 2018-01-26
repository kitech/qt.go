package qtcore

// /usr/include/qt/QtCore/qvariantanimation.h
// #include <qvariantanimation.h>
// #include <QtCore>

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
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"

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
// Public virtual
// const QMetaObject * metaObject()
func (this *QVariantAnimation) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QVariantAnimation10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qvariantanimation.h:68
// index:0
// Public
// void QVariantAnimation(class QObject *)
func NewQVariantAnimation(parent *QObject /*777 QObject **/) *QVariantAnimation {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QVariantAnimationC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantAnimationFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariantanimation.h:69
// index:0
// Public virtual
// void ~QVariantAnimation()
func DeleteQVariantAnimation(*QVariantAnimation) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QVariantAnimationD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:71
// index:0
// Public
// QVariant startValue()
func (this *QVariantAnimation) StartValue() *QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QVariantAnimation10startValueEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qvariantanimation.h:72
// index:0
// Public
// void setStartValue(const class QVariant &)
func (this *QVariantAnimation) SetStartValue(value *QVariant) {
	var convArg0 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QVariantAnimation13setStartValueERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:74
// index:0
// Public
// QVariant endValue()
func (this *QVariantAnimation) EndValue() *QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QVariantAnimation8endValueEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qvariantanimation.h:75
// index:0
// Public
// void setEndValue(const class QVariant &)
func (this *QVariantAnimation) SetEndValue(value *QVariant) {
	var convArg0 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QVariantAnimation11setEndValueERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:77
// index:0
// Public
// QVariant keyValueAt(qreal)
func (this *QVariantAnimation) KeyValueAt(step float64) *QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QVariantAnimation10keyValueAtEd", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), step)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qvariantanimation.h:78
// index:0
// Public
// void setKeyValueAt(qreal, const class QVariant &)
func (this *QVariantAnimation) SetKeyValueAt(step float64, value *QVariant) {
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QVariantAnimation13setKeyValueAtEdRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), step, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:83
// index:0
// Public
// QVariant currentValue()
func (this *QVariantAnimation) CurrentValue() *QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QVariantAnimation12currentValueEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qvariantanimation.h:85
// index:0
// Public virtual
// int duration()
func (this *QVariantAnimation) Duration() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QVariantAnimation8durationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qvariantanimation.h:86
// index:0
// Public
// void setDuration(int)
func (this *QVariantAnimation) SetDuration(msecs int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QVariantAnimation11setDurationEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:88
// index:0
// Public
// QEasingCurve easingCurve()
func (this *QVariantAnimation) EasingCurve() *QEasingCurve /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QVariantAnimation11easingCurveEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQEasingCurveFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qvariantanimation.h:89
// index:0
// Public
// void setEasingCurve(const class QEasingCurve &)
func (this *QVariantAnimation) SetEasingCurve(easing *QEasingCurve) {
	var convArg0 = easing.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QVariantAnimation14setEasingCurveERK12QEasingCurve", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:94
// index:0
// Public
// void valueChanged(const class QVariant &)
func (this *QVariantAnimation) ValueChanged(value *QVariant) {
	var convArg0 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QVariantAnimation12valueChangedERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:98
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QVariantAnimation) Event(event *QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QVariantAnimation5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qvariantanimation.h:100
// index:0
// Protected virtual
// void updateCurrentTime(int)
func (this *QVariantAnimation) UpdateCurrentTime(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QVariantAnimation17updateCurrentTimeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:101
// index:0
// Protected virtual
// void updateState(class QAbstractAnimation::State, class QAbstractAnimation::State)
func (this *QVariantAnimation) UpdateState(newState int, oldState int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QVariantAnimation11updateStateEN18QAbstractAnimation5StateES1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), newState, oldState)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:103
// index:0
// Protected virtual
// void updateCurrentValue(const class QVariant &)
func (this *QVariantAnimation) UpdateCurrentValue(value *QVariant) {
	var convArg0 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QVariantAnimation18updateCurrentValueERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:104
// index:0
// Protected virtual
// QVariant interpolated(const class QVariant &, const class QVariant &, qreal)
func (this *QVariantAnimation) Interpolated(from *QVariant, to *QVariant, progress float64) *QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = from.GetCthis()
	var convArg1 = to.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QVariantAnimation12interpolatedERK8QVariantS2_d", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1, progress)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
