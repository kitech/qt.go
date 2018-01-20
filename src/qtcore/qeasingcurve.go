//  header block begin
// /usr/include/qt/QtCore/qeasingcurve.h
// #include <qeasingcurve.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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

// /usr/include/qt/QtCore/qeasingcurve.h:77
// index:0
// void QEasingCurve(enum QEasingCurve::Type)
func NewQEasingCurve(type_ int) *QEasingCurve {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QEasingCurveC2ENS_4TypeE", ffiqt.FFI_TYPE_VOID, cthis, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQEasingCurveFromPointer(cthis)
	return gothis
}
func NewQEasingCurveFromPointer(cthis unsafe.Pointer) *QEasingCurve {
	return &QEasingCurve{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qeasingcurve.h:79
// index:0
// void ~QEasingCurve()
func DeleteQEasingCurve(*QEasingCurve) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QEasingCurveD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:89
// index:0
// inline
// void swap(class QEasingCurve &)
func (this *QEasingCurve) Swap(other unsafe.Pointer) {
	// 0: (, other QEasingCurve &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QEasingCurve4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:95
// index:0
// qreal amplitude()
func (this *QEasingCurve) Amplitude() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QEasingCurve9amplitudeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:96
// index:0
// void setAmplitude(qreal)
func (this *QEasingCurve) SetAmplitude(amplitude float64) {
	// 0: (, amplitude qreal), (&amplitude)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QEasingCurve12setAmplitudeEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &amplitude)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:98
// index:0
// qreal period()
func (this *QEasingCurve) Period() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QEasingCurve6periodEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:99
// index:0
// void setPeriod(qreal)
func (this *QEasingCurve) SetPeriod(period float64) {
	// 0: (, period qreal), (&period)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QEasingCurve9setPeriodEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &period)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:101
// index:0
// qreal overshoot()
func (this *QEasingCurve) Overshoot() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QEasingCurve9overshootEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:102
// index:0
// void setOvershoot(qreal)
func (this *QEasingCurve) SetOvershoot(overshoot float64) {
	// 0: (, overshoot qreal), (&overshoot)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QEasingCurve12setOvershootEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &overshoot)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:104
// index:0
// void addCubicBezierSegment(const class QPointF &, const class QPointF &, const class QPointF &)
func (this *QEasingCurve) AddCubicBezierSegment(c1 unsafe.Pointer, c2 unsafe.Pointer, endPoint unsafe.Pointer) {
	// 0: (, c1 const QPointF &, c2 const QPointF &, endPoint const QPointF &), (c1, c2, endPoint)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QEasingCurve21addCubicBezierSegmentERK7QPointFS2_S2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), c1, c2, endPoint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:105
// index:0
// void addTCBSegment(const class QPointF &, qreal, qreal, qreal)
func (this *QEasingCurve) AddTCBSegment(nextPoint unsafe.Pointer, t float64, c float64, b float64) {
	// 0: (, nextPoint const QPointF &, t qreal, c qreal, b qreal), (nextPoint, &t, &c, &b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QEasingCurve13addTCBSegmentERK7QPointFddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), nextPoint, &t, &c, &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:106
// index:0
// QVector<QPointF> toCubicSpline()
func (this *QEasingCurve) ToCubicSpline() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QEasingCurve13toCubicSplineEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:111
// index:0
// QEasingCurve::Type type()
func (this *QEasingCurve) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QEasingCurve4typeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:112
// index:0
// void setType(enum QEasingCurve::Type)
func (this *QEasingCurve) SetType(type_ int) {
	// 0: (, type QEasingCurve::Type), (&type_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QEasingCurve7setTypeENS_4TypeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:115
// index:0
// QEasingCurve::EasingFunction customType()
func (this *QEasingCurve) CustomType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QEasingCurve10customTypeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:117
// index:0
// qreal valueForProgress(qreal)
func (this *QEasingCurve) ValueForProgress(progress float64) {
	// 0: (, progress qreal), (&progress)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QEasingCurve16valueForProgressEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &progress)
	gopp.ErrPrint(err, rv)
}

//  body block end
