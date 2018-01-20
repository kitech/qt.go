//  header block begin
// /usr/include/qt/QtCore/qeasingcurve.h
// #include <qeasingcurve.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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
type QEasingCurve struct {
	*qtrt.CObject
}

func (this *QEasingCurve) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQEasingCurveFromPointer(cthis unsafe.Pointer) *QEasingCurve {
	return &QEasingCurve{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qeasingcurve.h:77
// index:0
// Public
// void QEasingCurve(enum QEasingCurve::Type)
func NewQEasingCurve(type_ int) *QEasingCurve {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QEasingCurveC2ENS_4TypeE", ffiqt.FFI_TYPE_VOID, cthis, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQEasingCurveFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qeasingcurve.h:79
// index:0
// Public
// void ~QEasingCurve()
func DeleteQEasingCurve(*QEasingCurve) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QEasingCurveD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:89
// index:0
// Public inline
// void swap(class QEasingCurve &)
func (this *QEasingCurve) Swap(other *QEasingCurve) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QEasingCurve4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:95
// index:0
// Public
// qreal amplitude()
func (this *QEasingCurve) Amplitude() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QEasingCurve9amplitudeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qeasingcurve.h:96
// index:0
// Public
// void setAmplitude(qreal)
func (this *QEasingCurve) SetAmplitude(amplitude float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QEasingCurve12setAmplitudeEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &amplitude)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:98
// index:0
// Public
// qreal period()
func (this *QEasingCurve) Period() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QEasingCurve6periodEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qeasingcurve.h:99
// index:0
// Public
// void setPeriod(qreal)
func (this *QEasingCurve) SetPeriod(period float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QEasingCurve9setPeriodEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &period)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:101
// index:0
// Public
// qreal overshoot()
func (this *QEasingCurve) Overshoot() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QEasingCurve9overshootEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qeasingcurve.h:102
// index:0
// Public
// void setOvershoot(qreal)
func (this *QEasingCurve) SetOvershoot(overshoot float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QEasingCurve12setOvershootEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &overshoot)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:104
// index:0
// Public
// void addCubicBezierSegment(const class QPointF &, const class QPointF &, const class QPointF &)
func (this *QEasingCurve) AddCubicBezierSegment(c1 *QPointF, c2 *QPointF, endPoint *QPointF) {
	var convArg0 = c1.GetCthis()
	var convArg1 = c2.GetCthis()
	var convArg2 = endPoint.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QEasingCurve21addCubicBezierSegmentERK7QPointFS2_S2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:105
// index:0
// Public
// void addTCBSegment(const class QPointF &, qreal, qreal, qreal)
func (this *QEasingCurve) AddTCBSegment(nextPoint *QPointF, t float64, c float64, b float64) {
	var convArg0 = nextPoint.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QEasingCurve13addTCBSegmentERK7QPointFddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &t, &c, &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:106
// index:0
// Public
// QVector<QPointF> toCubicSpline()
func (this *QEasingCurve) ToCubicSpline() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QEasingCurve13toCubicSplineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qeasingcurve.h:111
// index:0
// Public
// QEasingCurve::Type type()
func (this *QEasingCurve) Type() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QEasingCurve4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qeasingcurve.h:112
// index:0
// Public
// void setType(enum QEasingCurve::Type)
func (this *QEasingCurve) SetType(type_ int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QEasingCurve7setTypeENS_4TypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:115
// index:0
// Public
// QEasingCurve::EasingFunction customType()
func (this *QEasingCurve) CustomType() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QEasingCurve10customTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qeasingcurve.h:117
// index:0
// Public
// qreal valueForProgress(qreal)
func (this *QEasingCurve) ValueForProgress(progress float64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QEasingCurve16valueForProgressEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &progress)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
