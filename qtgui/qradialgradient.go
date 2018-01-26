package qtgui

// /usr/include/qt/QtGui/qbrush.h
// #include <qbrush.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
}

//  ext block end

//  body block begin
type QRadialGradient struct {
	*QGradient
}

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
// Public
// void QRadialGradient()
func NewQRadialGradient() *QRadialGradient {
	cthis := qtrt.Calloc(1, 256) // 64
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradientC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQRadialGradientFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:275
// index:1
// Public
// void QRadialGradient(const class QPointF &, qreal, const class QPointF &)
func NewQRadialGradient_1(center *qtcore.QPointF, radius float64, focalPoint *qtcore.QPointF) *QRadialGradient {
	cthis := qtrt.Calloc(1, 256) // 64
	var convArg0 = center.GetCthis()
	var convArg2 = focalPoint.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradientC2ERK7QPointFdS2_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, radius, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQRadialGradientFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:276
// index:2
// Public
// void QRadialGradient(qreal, qreal, qreal, qreal, qreal)
func NewQRadialGradient_2(cx float64, cy float64, radius float64, fx float64, fy float64) *QRadialGradient {
	cthis := qtrt.Calloc(1, 256) // 64
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradientC2Eddddd", ffiqt.FFI_TYPE_VOID, cthis, cx, cy, radius, fx, fy)
	gopp.ErrPrint(err, rv)
	gothis := NewQRadialGradientFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:278
// index:3
// Public
// void QRadialGradient(const class QPointF &, qreal)
func NewQRadialGradient_3(center *qtcore.QPointF, radius float64) *QRadialGradient {
	cthis := qtrt.Calloc(1, 256) // 64
	var convArg0 = center.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradientC2ERK7QPointFd", ffiqt.FFI_TYPE_VOID, cthis, convArg0, radius)
	gopp.ErrPrint(err, rv)
	gothis := NewQRadialGradientFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:279
// index:4
// Public
// void QRadialGradient(qreal, qreal, qreal)
func NewQRadialGradient_4(cx float64, cy float64, radius float64) *QRadialGradient {
	cthis := qtrt.Calloc(1, 256) // 64
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradientC2Eddd", ffiqt.FFI_TYPE_VOID, cthis, cx, cy, radius)
	gopp.ErrPrint(err, rv)
	gothis := NewQRadialGradientFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:281
// index:5
// Public
// void QRadialGradient(const class QPointF &, qreal, const class QPointF &, qreal)
func NewQRadialGradient_5(center *qtcore.QPointF, centerRadius float64, focalPoint *qtcore.QPointF, focalRadius float64) *QRadialGradient {
	cthis := qtrt.Calloc(1, 256) // 64
	var convArg0 = center.GetCthis()
	var convArg2 = focalPoint.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradientC2ERK7QPointFdS2_d", ffiqt.FFI_TYPE_VOID, cthis, convArg0, centerRadius, convArg2, focalRadius)
	gopp.ErrPrint(err, rv)
	gothis := NewQRadialGradientFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:282
// index:6
// Public
// void QRadialGradient(qreal, qreal, qreal, qreal, qreal, qreal)
func NewQRadialGradient_6(cx float64, cy float64, centerRadius float64, fx float64, fy float64, focalRadius float64) *QRadialGradient {
	cthis := qtrt.Calloc(1, 256) // 64
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradientC2Edddddd", ffiqt.FFI_TYPE_VOID, cthis, cx, cy, centerRadius, fx, fy, focalRadius)
	gopp.ErrPrint(err, rv)
	gothis := NewQRadialGradientFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:284
// index:0
// Public
// QPointF center()
func (this *QRadialGradient) Center() *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QRadialGradient6centerEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qbrush.h:285
// index:0
// Public
// void setCenter(const class QPointF &)
func (this *QRadialGradient) SetCenter(center *qtcore.QPointF) {
	var convArg0 = center.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradient9setCenterERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:286
// index:1
// Public inline
// void setCenter(qreal, qreal)
func (this *QRadialGradient) SetCenter_1(x float64, y float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradient9setCenterEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:288
// index:0
// Public
// QPointF focalPoint()
func (this *QRadialGradient) FocalPoint() *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QRadialGradient10focalPointEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qbrush.h:289
// index:0
// Public
// void setFocalPoint(const class QPointF &)
func (this *QRadialGradient) SetFocalPoint(focalPoint *qtcore.QPointF) {
	var convArg0 = focalPoint.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradient13setFocalPointERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:290
// index:1
// Public inline
// void setFocalPoint(qreal, qreal)
func (this *QRadialGradient) SetFocalPoint_1(x float64, y float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradient13setFocalPointEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:292
// index:0
// Public
// qreal radius()
func (this *QRadialGradient) Radius() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QRadialGradient6radiusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qbrush.h:293
// index:0
// Public
// void setRadius(qreal)
func (this *QRadialGradient) SetRadius(radius float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradient9setRadiusEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), radius)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:295
// index:0
// Public
// qreal centerRadius()
func (this *QRadialGradient) CenterRadius() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QRadialGradient12centerRadiusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qbrush.h:296
// index:0
// Public
// void setCenterRadius(qreal)
func (this *QRadialGradient) SetCenterRadius(radius float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradient15setCenterRadiusEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), radius)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:298
// index:0
// Public
// qreal focalRadius()
func (this *QRadialGradient) FocalRadius() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QRadialGradient11focalRadiusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qbrush.h:299
// index:0
// Public
// void setFocalRadius(qreal)
func (this *QRadialGradient) SetFocalRadius(radius float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradient14setFocalRadiusEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), radius)
	gopp.ErrPrint(err, rv)
}

//  body block end
