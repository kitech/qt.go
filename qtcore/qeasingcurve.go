package qtcore

// /usr/include/qt/QtCore/qeasingcurve.h
// #include <qeasingcurve.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QEasingCurve struct {
	*qtrt.CObject
}
type QEasingCurve_ITF interface {
	QEasingCurve_PTR() *QEasingCurve
}

func (ptr *QEasingCurve) QEasingCurve_PTR() *QEasingCurve { return ptr }

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
// [-2] void QEasingCurve(QEasingCurve::Type)

/*
Constructs an easing curve of the given type.
*/
func NewQEasingCurve(type_ int) *QEasingCurve {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QEasingCurveC2ENS_4TypeE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQEasingCurveFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQEasingCurve)
	return gothis
}

// /usr/include/qt/QtCore/qeasingcurve.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QEasingCurve(QEasingCurve::Type)

/*
Constructs an easing curve of the given type.
*/
func NewQEasingCurve__() *QEasingCurve {
	// arg: 0, QEasingCurve::Type=Enum, QEasingCurve::Type=Enum,
	type_ := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QEasingCurveC2ENS_4TypeE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQEasingCurveFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQEasingCurve)
	return gothis
}

// /usr/include/qt/QtCore/qeasingcurve.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QEasingCurve()

/*

 */
func DeleteQEasingCurve(this *QEasingCurve) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QEasingCurveD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qeasingcurve.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QEasingCurve & operator=(const QEasingCurve &)

/*

 */
func (this *QEasingCurve) Operator_equal(other QEasingCurve_ITF) *QEasingCurve {
	var convArg0 unsafe.Pointer
	if other != nil && other.QEasingCurve_PTR() != nil {
		convArg0 = other.QEasingCurve_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QEasingCurveaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQEasingCurveFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQEasingCurve)
	return rv2
}

// /usr/include/qt/QtCore/qeasingcurve.h:85
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QEasingCurve & operator=(QEasingCurve &&)

/*

 */
func (this *QEasingCurve) Operator_equal_1(other unsafe.Pointer /*333*/) *QEasingCurve {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QEasingCurveaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQEasingCurveFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQEasingCurve)
	return rv2
}

// /usr/include/qt/QtCore/qeasingcurve.h:89
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QEasingCurve &)

/*
Swaps curve other with this curve. This operation is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func (this *QEasingCurve) Swap(other QEasingCurve_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QEasingCurve_PTR() != nil {
		convArg0 = other.QEasingCurve_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QEasingCurve4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:91
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QEasingCurve &) const

/*

 */
func (this *QEasingCurve) Operator_equal_equal(other QEasingCurve_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QEasingCurve_PTR() != nil {
		convArg0 = other.QEasingCurve_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QEasingCurveeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qeasingcurve.h:92
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QEasingCurve &) const

/*

 */
func (this *QEasingCurve) Operator_not_equal(other QEasingCurve_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QEasingCurve_PTR() != nil {
		convArg0 = other.QEasingCurve_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QEasingCurveneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qeasingcurve.h:95
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal amplitude() const

/*
Returns the amplitude. This is not applicable for all curve types. It is only applicable for bounce and elastic curves (curves of type() QEasingCurve::InBounce, QEasingCurve::OutBounce, QEasingCurve::InOutBounce, QEasingCurve::OutInBounce, QEasingCurve::InElastic, QEasingCurve::OutElastic, QEasingCurve::InOutElastic or QEasingCurve::OutInElastic).

See also setAmplitude().
*/
func (this *QEasingCurve) Amplitude() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QEasingCurve9amplitudeEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qeasingcurve.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAmplitude(qreal)

/*
Sets the amplitude to amplitude.

This will set the amplitude of the bounce or the amplitude of the elastic "spring" effect. The higher the number, the higher the amplitude.

See also amplitude().
*/
func (this *QEasingCurve) SetAmplitude(amplitude float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QEasingCurve12setAmplitudeEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), amplitude)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal period() const

/*
Returns the period. This is not applicable for all curve types. It is only applicable if type() is QEasingCurve::InElastic, QEasingCurve::OutElastic, QEasingCurve::InOutElastic or QEasingCurve::OutInElastic.

See also setPeriod().
*/
func (this *QEasingCurve) Period() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QEasingCurve6periodEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qeasingcurve.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPeriod(qreal)

/*
Sets the period to period. Setting a small period value will give a high frequency of the curve. A large period will give it a small frequency.

See also period().
*/
func (this *QEasingCurve) SetPeriod(period float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QEasingCurve9setPeriodEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), period)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal overshoot() const

/*
Returns the overshoot. This is not applicable for all curve types. It is only applicable if type() is QEasingCurve::InBack, QEasingCurve::OutBack, QEasingCurve::InOutBack or QEasingCurve::OutInBack.

See also setOvershoot().
*/
func (this *QEasingCurve) Overshoot() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QEasingCurve9overshootEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qeasingcurve.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOvershoot(qreal)

/*
Sets the overshoot to overshoot.

0 produces no overshoot, and the default value of 1.70158 produces an overshoot of 10 percent.

See also overshoot().
*/
func (this *QEasingCurve) SetOvershoot(overshoot float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QEasingCurve12setOvershootEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), overshoot)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addCubicBezierSegment(const QPointF &, const QPointF &, const QPointF &)

/*
Adds a segment of a cubic bezier spline to define a custom easing curve. It is only applicable if type() is QEasingCurve::BezierSpline. Note that the spline implicitly starts at (0.0, 0.0) and has to end at (1.0, 1.0) to be a valid easing curve. c1 and c2 are the control points used for drawing the curve. endPoint is the endpoint of the curve.
*/
func (this *QEasingCurve) AddCubicBezierSegment(c1 QPointF_ITF, c2 QPointF_ITF, endPoint QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if c1 != nil && c1.QPointF_PTR() != nil {
		convArg0 = c1.QPointF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if c2 != nil && c2.QPointF_PTR() != nil {
		convArg1 = c2.QPointF_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if endPoint != nil && endPoint.QPointF_PTR() != nil {
		convArg2 = endPoint.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QEasingCurve21addCubicBezierSegmentERK7QPointFS2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addTCBSegment(const QPointF &, qreal, qreal, qreal)

/*
Adds a segment of a TCB bezier spline to define a custom easing curve. It is only applicable if type() is QEasingCurve::TCBSpline. The spline has to start explitly at (0.0, 0.0) and has to end at (1.0, 1.0) to be a valid easing curve. The tension t changes the length of the tangent vector. The continuity c changes the sharpness in change between the tangents. The bias b changes the direction of the tangent vector. nextPoint is the sample position. All three parameters are valid between -1 and 1 and define the tangent of the control point. If all three parameters are 0 the resulting spline is a Catmull-Rom spline. The begin and endpoint always have a bias of -1 and 1, since the outer tangent is not defined.
*/
func (this *QEasingCurve) AddTCBSegment(nextPoint QPointF_ITF, t float64, c float64, b float64) {
	var convArg0 unsafe.Pointer
	if nextPoint != nil && nextPoint.QPointF_PTR() != nil {
		convArg0 = nextPoint.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QEasingCurve13addTCBSegmentERK7QPointFddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, t, c, b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:111
// index:0
// Public Visibility=Default Availability=Available
// [4] QEasingCurve::Type type() const

/*
Returns the type of the easing curve.

See also setType().
*/
func (this *QEasingCurve) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QEasingCurve4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setType(QEasingCurve::Type)

/*
Sets the type of the easing curve to type.

See also type().
*/
func (this *QEasingCurve) SetType(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QEasingCurve7setTypeENS_4TypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeasingcurve.h:115
// index:0
// Public Visibility=Default Availability=Available
// [8] QEasingCurve::EasingFunction customType() const

/*
Returns the function pointer to the custom easing curve. If type() does not return QEasingCurve::Custom, this function will return 0.

See also setCustomType().
*/
func (this *QEasingCurve) CustomType() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QEasingCurve10customTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qeasingcurve.h:117
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal valueForProgress(qreal) const

/*
Return the effective progress for the easing curve at progress. Whereas progress must be between 0 and 1, the returned effective progress can be outside those bounds. For example, QEasingCurve::InBack will return negative values in the beginning of the function.
*/
func (this *QEasingCurve) ValueForProgress(progress float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QEasingCurve16valueForProgressEd", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), progress)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

/*
The type of easing curve.

ConstantValue
QEasingCurve::Linear0




Easing curve for a linear (t) function: velocity is constant.

ConstantValue
QEasingCurve::InQuad1




Easing curve for a quadratic (t^2) function: accelerating from zero velocity.

ConstantValue
QEasingCurve::OutQuad2




Easing curve for a quadratic (t^2) function: decelerating to zero velocity.

ConstantValue
QEasingCurve::InOutQuad3




Easing curve for a quadratic (t^2) function: acceleration until halfway, then deceleration.

ConstantValue
QEasingCurve::OutInQuad4




Easing curve for a quadratic (t^2) function: deceleration until halfway, then acceleration.

ConstantValue
QEasingCurve::InCubic5




Easing curve for a cubic (t^3) function: accelerating from zero velocity.

ConstantValue
QEasingCurve::OutCubic6




Easing curve for a cubic (t^3) function: decelerating to zero velocity.

ConstantValue
QEasingCurve::InOutCubic7




Easing curve for a cubic (t^3) function: acceleration until halfway, then deceleration.

ConstantValue
QEasingCurve::OutInCubic8




Easing curve for a cubic (t^3) function: deceleration until halfway, then acceleration.

ConstantValue
QEasingCurve::InQuart9




Easing curve for a quartic (t^4) function: accelerating from zero velocity.

ConstantValue




Easing curve for a quartic (t^4) function: decelerating to zero velocity.

ConstantValue




Easing curve for a quartic (t^4) function: acceleration until halfway, then deceleration.

ConstantValue




Easing curve for a quartic (t^4) function: deceleration until halfway, then acceleration.

ConstantValue




Easing curve for a quintic (t^5) easing in: accelerating from zero velocity.

ConstantValue




Easing curve for a quintic (t^5) function: decelerating to zero velocity.

ConstantValue




Easing curve for a quintic (t^5) function: acceleration until halfway, then deceleration.

ConstantValue




Easing curve for a quintic (t^5) function: deceleration until halfway, then acceleration.

ConstantValue




Easing curve for a sinusoidal (sin(t)) function: accelerating from zero velocity.

ConstantValue




Easing curve for a sinusoidal (sin(t)) function: decelerating to zero velocity.

ConstantValue




Easing curve for a sinusoidal (sin(t)) function: acceleration until halfway, then deceleration.

ConstantValue




Easing curve for a sinusoidal (sin(t)) function: deceleration until halfway, then acceleration.

ConstantValue




Easing curve for an exponential (2^t) function: accelerating from zero velocity.

ConstantValue




Easing curve for an exponential (2^t) function: decelerating to zero velocity.

ConstantValue




Easing curve for an exponential (2^t) function: acceleration until halfway, then deceleration.

ConstantValue




Easing curve for an exponential (2^t) function: deceleration until halfway, then acceleration.

ConstantValue




Easing curve for a circular (sqrt(1-t^2)) function: accelerating from zero velocity.

ConstantValue




Easing curve for a circular (sqrt(1-t^2)) function: decelerating to zero velocity.

ConstantValue




Easing curve for a circular (sqrt(1-t^2)) function: acceleration until halfway, then deceleration.

ConstantValue




Easing curve for a circular (sqrt(1-t^2)) function: deceleration until halfway, then acceleration.

ConstantValue




Easing curve for an elastic (exponentially decaying sine wave) function: accelerating from zero velocity. The peak amplitude can be set with the amplitude parameter, and the period of decay by the period parameter.

ConstantValue




Easing curve for an elastic (exponentially decaying sine wave) function: decelerating to zero velocity. The peak amplitude can be set with the amplitude parameter, and the period of decay by the period parameter.

ConstantValue




Easing curve for an elastic (exponentially decaying sine wave) function: acceleration until halfway, then deceleration.

ConstantValue




Easing curve for an elastic (exponentially decaying sine wave) function: deceleration until halfway, then acceleration.

ConstantValue




Easing curve for a back (overshooting cubic function: (s+1)*t^3 - s*t^2) easing in: accelerating from zero velocity.

ConstantValue




Easing curve for a back (overshooting cubic function: (s+1)*t^3 - s*t^2) easing out: decelerating to zero velocity.

ConstantValue




Easing curve for a back (overshooting cubic function: (s+1)*t^3 - s*t^2) easing in/out: acceleration until halfway, then deceleration.

ConstantValue




Easing curve for a back (overshooting cubic easing: (s+1)*t^3 - s*t^2) easing out/in: deceleration until halfway, then acceleration.

ConstantValue




Easing curve for a bounce (exponentially decaying parabolic bounce) function: accelerating from zero velocity.

ConstantValue




Easing curve for a bounce (exponentially decaying parabolic bounce) function: decelerating from zero velocity.

ConstantValue




Easing curve for a bounce (exponentially decaying parabolic bounce) function easing in/out: acceleration until halfway, then deceleration.

ConstantValue




Easing curve for a bounce (exponentially decaying parabolic bounce) function easing out/in: deceleration until halfway, then acceleration.



See also addCubicBezierSegment() and addTCBSegment().

*/
type QEasingCurve__Type = int

//
const QEasingCurve__Linear QEasingCurve__Type = 0

//
const QEasingCurve__InQuad QEasingCurve__Type = 1

//
const QEasingCurve__OutQuad QEasingCurve__Type = 2

//
const QEasingCurve__InOutQuad QEasingCurve__Type = 3

//
const QEasingCurve__OutInQuad QEasingCurve__Type = 4

//
const QEasingCurve__InCubic QEasingCurve__Type = 5

//
const QEasingCurve__OutCubic QEasingCurve__Type = 6

//
const QEasingCurve__InOutCubic QEasingCurve__Type = 7

//
const QEasingCurve__OutInCubic QEasingCurve__Type = 8

//
const QEasingCurve__InQuart QEasingCurve__Type = 9

// 0
const QEasingCurve__OutQuart QEasingCurve__Type = 10

// 1
const QEasingCurve__InOutQuart QEasingCurve__Type = 11

// 2
const QEasingCurve__OutInQuart QEasingCurve__Type = 12

// 3
const QEasingCurve__InQuint QEasingCurve__Type = 13

// 4
const QEasingCurve__OutQuint QEasingCurve__Type = 14

// 5
const QEasingCurve__InOutQuint QEasingCurve__Type = 15

// 6
const QEasingCurve__OutInQuint QEasingCurve__Type = 16

// 7
const QEasingCurve__InSine QEasingCurve__Type = 17

// 8
const QEasingCurve__OutSine QEasingCurve__Type = 18

// 9
const QEasingCurve__InOutSine QEasingCurve__Type = 19

// 0
const QEasingCurve__OutInSine QEasingCurve__Type = 20

// 1
const QEasingCurve__InExpo QEasingCurve__Type = 21

// 2
const QEasingCurve__OutExpo QEasingCurve__Type = 22

// 3
const QEasingCurve__InOutExpo QEasingCurve__Type = 23

// 4
const QEasingCurve__OutInExpo QEasingCurve__Type = 24

// 5
const QEasingCurve__InCirc QEasingCurve__Type = 25

// 6
const QEasingCurve__OutCirc QEasingCurve__Type = 26

// 7
const QEasingCurve__InOutCirc QEasingCurve__Type = 27

// 8
const QEasingCurve__OutInCirc QEasingCurve__Type = 28

// 9
const QEasingCurve__InElastic QEasingCurve__Type = 29

// 0
const QEasingCurve__OutElastic QEasingCurve__Type = 30

// 1
const QEasingCurve__InOutElastic QEasingCurve__Type = 31

// 2
const QEasingCurve__OutInElastic QEasingCurve__Type = 32

// 3
const QEasingCurve__InBack QEasingCurve__Type = 33

// 4
const QEasingCurve__OutBack QEasingCurve__Type = 34

// 5
const QEasingCurve__InOutBack QEasingCurve__Type = 35

// 6
const QEasingCurve__OutInBack QEasingCurve__Type = 36

// 7
const QEasingCurve__InBounce QEasingCurve__Type = 37

// 8
const QEasingCurve__OutBounce QEasingCurve__Type = 38

// 9
const QEasingCurve__InOutBounce QEasingCurve__Type = 39

// 0
const QEasingCurve__OutInBounce QEasingCurve__Type = 40

//
const QEasingCurve__InCurve QEasingCurve__Type = 41

//
const QEasingCurve__OutCurve QEasingCurve__Type = 42

//
const QEasingCurve__SineCurve QEasingCurve__Type = 43

//
const QEasingCurve__CosineCurve QEasingCurve__Type = 44

//
const QEasingCurve__BezierSpline QEasingCurve__Type = 45

//
const QEasingCurve__TCBSpline QEasingCurve__Type = 46

//
const QEasingCurve__Custom QEasingCurve__Type = 47

//
const QEasingCurve__NCurveTypes QEasingCurve__Type = 48

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
