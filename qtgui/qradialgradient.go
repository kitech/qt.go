package qtgui

// /usr/include/qt/QtGui/qbrush.h
// #include <qbrush.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QRadialGradient struct {
	*QGradient
}
type QRadialGradient_ITF interface {
	QGradient_ITF
	QRadialGradient_PTR() *QRadialGradient
}

func (ptr *QRadialGradient) QRadialGradient_PTR() *QRadialGradient { return ptr }

func (this *QRadialGradient) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGradient.GetCthis()
	}
}
func (this *QRadialGradient) SetCthis(cthis unsafe.Pointer) {
	this.QGradient = NewQGradientFromPointer(cthis)
}
func NewQRadialGradientFromPointer(cthis unsafe.Pointer) *QRadialGradient {
	bcthis0 := NewQGradientFromPointer(cthis)
	return &QRadialGradient{bcthis0}
}
func (*QRadialGradient) NewFromPointer(cthis unsafe.Pointer) *QRadialGradient {
	return NewQRadialGradientFromPointer(cthis)
}

// /usr/include/qt/QtGui/qbrush.h:274
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRadialGradient()

/*

 */
func (*QRadialGradient) NewForInherit() *QRadialGradient {
	return NewQRadialGradient()
}
func NewQRadialGradient() *QRadialGradient {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QRadialGradientC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRadialGradientFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRadialGradient)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:275
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QRadialGradient(const QPointF &, qreal, const QPointF &)

/*

 */
func (*QRadialGradient) NewForInherit1(center qtcore.QPointF_ITF, radius float64, focalPoint qtcore.QPointF_ITF) *QRadialGradient {
	return NewQRadialGradient1(center, radius, focalPoint)
}
func NewQRadialGradient1(center qtcore.QPointF_ITF, radius float64, focalPoint qtcore.QPointF_ITF) *QRadialGradient {
	var convArg0 unsafe.Pointer
	if center != nil && center.QPointF_PTR() != nil {
		convArg0 = center.QPointF_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if focalPoint != nil && focalPoint.QPointF_PTR() != nil {
		convArg2 = focalPoint.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QRadialGradientC2ERK7QPointFdS2_", qtrt.FFI_TYPE_POINTER, convArg0, radius, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRadialGradientFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRadialGradient)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:276
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QRadialGradient(qreal, qreal, qreal, qreal, qreal)

/*

 */
func (*QRadialGradient) NewForInherit2(cx float64, cy float64, radius float64, fx float64, fy float64) *QRadialGradient {
	return NewQRadialGradient2(cx, cy, radius, fx, fy)
}
func NewQRadialGradient2(cx float64, cy float64, radius float64, fx float64, fy float64) *QRadialGradient {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QRadialGradientC2Eddddd", qtrt.FFI_TYPE_POINTER, cx, cy, radius, fx, fy)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRadialGradientFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRadialGradient)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:278
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QRadialGradient(const QPointF &, qreal)

/*

 */
func (*QRadialGradient) NewForInherit3(center qtcore.QPointF_ITF, radius float64) *QRadialGradient {
	return NewQRadialGradient3(center, radius)
}
func NewQRadialGradient3(center qtcore.QPointF_ITF, radius float64) *QRadialGradient {
	var convArg0 unsafe.Pointer
	if center != nil && center.QPointF_PTR() != nil {
		convArg0 = center.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QRadialGradientC2ERK7QPointFd", qtrt.FFI_TYPE_POINTER, convArg0, radius)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRadialGradientFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRadialGradient)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:279
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QRadialGradient(qreal, qreal, qreal)

/*

 */
func (*QRadialGradient) NewForInherit4(cx float64, cy float64, radius float64) *QRadialGradient {
	return NewQRadialGradient4(cx, cy, radius)
}
func NewQRadialGradient4(cx float64, cy float64, radius float64) *QRadialGradient {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QRadialGradientC2Eddd", qtrt.FFI_TYPE_POINTER, cx, cy, radius)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRadialGradientFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRadialGradient)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:281
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QRadialGradient(const QPointF &, qreal, const QPointF &, qreal)

/*

 */
func (*QRadialGradient) NewForInherit5(center qtcore.QPointF_ITF, centerRadius float64, focalPoint qtcore.QPointF_ITF, focalRadius float64) *QRadialGradient {
	return NewQRadialGradient5(center, centerRadius, focalPoint, focalRadius)
}
func NewQRadialGradient5(center qtcore.QPointF_ITF, centerRadius float64, focalPoint qtcore.QPointF_ITF, focalRadius float64) *QRadialGradient {
	var convArg0 unsafe.Pointer
	if center != nil && center.QPointF_PTR() != nil {
		convArg0 = center.QPointF_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if focalPoint != nil && focalPoint.QPointF_PTR() != nil {
		convArg2 = focalPoint.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QRadialGradientC2ERK7QPointFdS2_d", qtrt.FFI_TYPE_POINTER, convArg0, centerRadius, convArg2, focalRadius)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRadialGradientFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRadialGradient)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:282
// index:6
// Public Visibility=Default Availability=Available
// [-2] void QRadialGradient(qreal, qreal, qreal, qreal, qreal, qreal)

/*

 */
func (*QRadialGradient) NewForInherit6(cx float64, cy float64, centerRadius float64, fx float64, fy float64, focalRadius float64) *QRadialGradient {
	return NewQRadialGradient6(cx, cy, centerRadius, fx, fy, focalRadius)
}
func NewQRadialGradient6(cx float64, cy float64, centerRadius float64, fx float64, fy float64, focalRadius float64) *QRadialGradient {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QRadialGradientC2Edddddd", qtrt.FFI_TYPE_POINTER, cx, cy, centerRadius, fx, fy, focalRadius)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRadialGradientFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRadialGradient)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:284
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF center() const

/*

 */
func (this *QRadialGradient) Center() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QRadialGradient6centerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qbrush.h:285
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCenter(const QPointF &)

/*

 */
func (this *QRadialGradient) SetCenter(center qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if center != nil && center.QPointF_PTR() != nil {
		convArg0 = center.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QRadialGradient9setCenterERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:286
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setCenter(qreal, qreal)

/*

 */
func (this *QRadialGradient) SetCenter1(x float64, y float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QRadialGradient9setCenterEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:288
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF focalPoint() const

/*

 */
func (this *QRadialGradient) FocalPoint() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QRadialGradient10focalPointEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qbrush.h:289
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFocalPoint(const QPointF &)

/*

 */
func (this *QRadialGradient) SetFocalPoint(focalPoint qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if focalPoint != nil && focalPoint.QPointF_PTR() != nil {
		convArg0 = focalPoint.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QRadialGradient13setFocalPointERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:290
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setFocalPoint(qreal, qreal)

/*

 */
func (this *QRadialGradient) SetFocalPoint1(x float64, y float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QRadialGradient13setFocalPointEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:292
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal radius() const

/*

 */
func (this *QRadialGradient) Radius() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QRadialGradient6radiusEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qbrush.h:293
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRadius(qreal)

/*

 */
func (this *QRadialGradient) SetRadius(radius float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QRadialGradient9setRadiusEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), radius)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:295
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal centerRadius() const

/*

 */
func (this *QRadialGradient) CenterRadius() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QRadialGradient12centerRadiusEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qbrush.h:296
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCenterRadius(qreal)

/*

 */
func (this *QRadialGradient) SetCenterRadius(radius float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QRadialGradient15setCenterRadiusEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), radius)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:298
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal focalRadius() const

/*

 */
func (this *QRadialGradient) FocalRadius() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QRadialGradient11focalRadiusEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qbrush.h:299
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFocalRadius(qreal)

/*

 */
func (this *QRadialGradient) SetFocalRadius(radius float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QRadialGradient14setFocalRadiusEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), radius)
	qtrt.ErrPrint(err, rv)
}

func DeleteQRadialGradient(this *QRadialGradient) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QRadialGradientD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
}

//  keep block end
