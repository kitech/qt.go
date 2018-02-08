package qtcore

// /usr/include/qt/QtCore/qeasingcurve.h
// #include <qeasingcurve.h>
// #include <QtCore>

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
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin

type QEasingCurve struct {
	*qtrt.CObject
}

func (this *QEasingCurve) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QEasingCurve) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQEasingCurveFromPointer(cthis unsafe.Pointer) *QEasingCurve {
	return &QEasingCurve{&qtrt.CObject{cthis}}
}
func (*QEasingCurve) NewFromPointer(cthis unsafe.Pointer) *QEasingCurve {
	return NewQEasingCurveFromPointer(cthis)
}

// /usr/include/qt/QtCore/qeasingcurve.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QEasingCurve(enum QEasingCurve::Type)
func NewQEasingCurve(type_ int) *QEasingCurve {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QEasingCurveC2ENS_4TypeE", qtrt.FFI_TYPE_POINTER, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQEasingCurveFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQEasingCurve)
	return gothis
}

// /usr/include/qt/QtCore/qeasingcurve.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QEasingCurve()
func DeleteQEasingCurve(this *QEasingCurve) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QEasingCurveD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qeasingcurve.h:89
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QEasingCurve &)
func (this *QEasingCurve) Swap(other *QEasingCurve) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QEasingCurve4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:95
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal amplitude()
func (this *QEasingCurve) Amplitude() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QEasingCurve9amplitudeEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qeasingcurve.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAmplitude(qreal)
func (this *QEasingCurve) SetAmplitude(amplitude float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QEasingCurve12setAmplitudeEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), amplitude)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal period()
func (this *QEasingCurve) Period() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QEasingCurve6periodEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qeasingcurve.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPeriod(qreal)
func (this *QEasingCurve) SetPeriod(period float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QEasingCurve9setPeriodEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), period)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal overshoot()
func (this *QEasingCurve) Overshoot() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QEasingCurve9overshootEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qeasingcurve.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOvershoot(qreal)
func (this *QEasingCurve) SetOvershoot(overshoot float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QEasingCurve12setOvershootEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), overshoot)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addCubicBezierSegment(const QPointF &, const QPointF &, const QPointF &)
func (this *QEasingCurve) AddCubicBezierSegment(c1 *QPointF, c2 *QPointF, endPoint *QPointF) {
	var convArg0 = c1.GetCthis()
	var convArg1 = c2.GetCthis()
	var convArg2 = endPoint.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QEasingCurve21addCubicBezierSegmentERK7QPointFS2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addTCBSegment(const QPointF &, qreal, qreal, qreal)
func (this *QEasingCurve) AddTCBSegment(nextPoint *QPointF, t float64, c float64, b float64) {
	var convArg0 = nextPoint.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QEasingCurve13addTCBSegmentERK7QPointFddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, t, c, b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:111
// index:0
// Public Visibility=Default Availability=Available
// [4] QEasingCurve::Type type()
func (this *QEasingCurve) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QEasingCurve4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setType(enum QEasingCurve::Type)
func (this *QEasingCurve) SetType(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QEasingCurve7setTypeENS_4TypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:115
// index:0
// Public Visibility=Default Availability=Available
// [8] QEasingCurve::EasingFunction customType()
func (this *QEasingCurve) CustomType() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QEasingCurve10customTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qeasingcurve.h:117
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal valueForProgress(qreal)
func (this *QEasingCurve) ValueForProgress(progress float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QEasingCurve16valueForProgressEd", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), progress)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

type QEasingCurve__Type = int

const QEasingCurve__Linear QEasingCurve__Type = 0
const QEasingCurve__InQuad QEasingCurve__Type = 1
const QEasingCurve__OutQuad QEasingCurve__Type = 2
const QEasingCurve__InOutQuad QEasingCurve__Type = 3
const QEasingCurve__OutInQuad QEasingCurve__Type = 4
const QEasingCurve__InCubic QEasingCurve__Type = 5
const QEasingCurve__OutCubic QEasingCurve__Type = 6
const QEasingCurve__InOutCubic QEasingCurve__Type = 7
const QEasingCurve__OutInCubic QEasingCurve__Type = 8
const QEasingCurve__InQuart QEasingCurve__Type = 9
const QEasingCurve__OutQuart QEasingCurve__Type = 10
const QEasingCurve__InOutQuart QEasingCurve__Type = 11
const QEasingCurve__OutInQuart QEasingCurve__Type = 12
const QEasingCurve__InQuint QEasingCurve__Type = 13
const QEasingCurve__OutQuint QEasingCurve__Type = 14
const QEasingCurve__InOutQuint QEasingCurve__Type = 15
const QEasingCurve__OutInQuint QEasingCurve__Type = 16
const QEasingCurve__InSine QEasingCurve__Type = 17
const QEasingCurve__OutSine QEasingCurve__Type = 18
const QEasingCurve__InOutSine QEasingCurve__Type = 19
const QEasingCurve__OutInSine QEasingCurve__Type = 20
const QEasingCurve__InExpo QEasingCurve__Type = 21
const QEasingCurve__OutExpo QEasingCurve__Type = 22
const QEasingCurve__InOutExpo QEasingCurve__Type = 23
const QEasingCurve__OutInExpo QEasingCurve__Type = 24
const QEasingCurve__InCirc QEasingCurve__Type = 25
const QEasingCurve__OutCirc QEasingCurve__Type = 26
const QEasingCurve__InOutCirc QEasingCurve__Type = 27
const QEasingCurve__OutInCirc QEasingCurve__Type = 28
const QEasingCurve__InElastic QEasingCurve__Type = 29
const QEasingCurve__OutElastic QEasingCurve__Type = 30
const QEasingCurve__InOutElastic QEasingCurve__Type = 31
const QEasingCurve__OutInElastic QEasingCurve__Type = 32
const QEasingCurve__InBack QEasingCurve__Type = 33
const QEasingCurve__OutBack QEasingCurve__Type = 34
const QEasingCurve__InOutBack QEasingCurve__Type = 35
const QEasingCurve__OutInBack QEasingCurve__Type = 36
const QEasingCurve__InBounce QEasingCurve__Type = 37
const QEasingCurve__OutBounce QEasingCurve__Type = 38
const QEasingCurve__InOutBounce QEasingCurve__Type = 39
const QEasingCurve__OutInBounce QEasingCurve__Type = 40
const QEasingCurve__InCurve QEasingCurve__Type = 41
const QEasingCurve__OutCurve QEasingCurve__Type = 42
const QEasingCurve__SineCurve QEasingCurve__Type = 43
const QEasingCurve__CosineCurve QEasingCurve__Type = 44
const QEasingCurve__BezierSpline QEasingCurve__Type = 45
const QEasingCurve__TCBSpline QEasingCurve__Type = 46
const QEasingCurve__Custom QEasingCurve__Type = 47
const QEasingCurve__NCurveTypes QEasingCurve__Type = 48

//  body block end
