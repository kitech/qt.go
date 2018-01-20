//  header block begin
// /usr/include/qt/QtGui/qbrush.h
// #include <qbrush.h>
// #include <QtGui>
package qtgui

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
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
	return this.QGradient.GetCthis()
}

// /usr/include/qt/QtGui/qbrush.h:274
// index:0
// void QRadialGradient()
func NewQRadialGradient() *QRadialGradient {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradientC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQRadialGradientFromPointer(cthis)
	return gothis
}
func NewQRadialGradientFromPointer(cthis unsafe.Pointer) *QRadialGradient {
	bcthis0 := NewQGradientFromPointer(cthis)
	return &QRadialGradient{bcthis0}
}

// /usr/include/qt/QtGui/qbrush.h:275
// index:1
// void QRadialGradient(const class QPointF &, qreal, const class QPointF &)
func NewQRadialGradient_1(center unsafe.Pointer, radius float64, focalPoint unsafe.Pointer) *QRadialGradient {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradientC2ERK7QPointFdS2_", ffiqt.FFI_TYPE_VOID, cthis, center, &radius, focalPoint)
	gopp.ErrPrint(err, rv)
	gothis := NewQRadialGradientFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:276
// index:2
// void QRadialGradient(qreal, qreal, qreal, qreal, qreal)
func NewQRadialGradient_2(cx float64, cy float64, radius float64, fx float64, fy float64) *QRadialGradient {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradientC2Eddddd", ffiqt.FFI_TYPE_VOID, cthis, &cx, &cy, &radius, &fx, &fy)
	gopp.ErrPrint(err, rv)
	gothis := NewQRadialGradientFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:278
// index:3
// void QRadialGradient(const class QPointF &, qreal)
func NewQRadialGradient_3(center unsafe.Pointer, radius float64) *QRadialGradient {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradientC2ERK7QPointFd", ffiqt.FFI_TYPE_VOID, cthis, center, &radius)
	gopp.ErrPrint(err, rv)
	gothis := NewQRadialGradientFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:279
// index:4
// void QRadialGradient(qreal, qreal, qreal)
func NewQRadialGradient_4(cx float64, cy float64, radius float64) *QRadialGradient {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradientC2Eddd", ffiqt.FFI_TYPE_VOID, cthis, &cx, &cy, &radius)
	gopp.ErrPrint(err, rv)
	gothis := NewQRadialGradientFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:281
// index:5
// void QRadialGradient(const class QPointF &, qreal, const class QPointF &, qreal)
func NewQRadialGradient_5(center unsafe.Pointer, centerRadius float64, focalPoint unsafe.Pointer, focalRadius float64) *QRadialGradient {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradientC2ERK7QPointFdS2_d", ffiqt.FFI_TYPE_VOID, cthis, center, &centerRadius, focalPoint, &focalRadius)
	gopp.ErrPrint(err, rv)
	gothis := NewQRadialGradientFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:282
// index:6
// void QRadialGradient(qreal, qreal, qreal, qreal, qreal, qreal)
func NewQRadialGradient_6(cx float64, cy float64, centerRadius float64, fx float64, fy float64, focalRadius float64) *QRadialGradient {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradientC2Edddddd", ffiqt.FFI_TYPE_VOID, cthis, &cx, &cy, &centerRadius, &fx, &fy, &focalRadius)
	gopp.ErrPrint(err, rv)
	gothis := NewQRadialGradientFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:284
// index:0
// QPointF center()
func (this *QRadialGradient) Center() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QRadialGradient6centerEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:285
// index:0
// void setCenter(const class QPointF &)
func (this *QRadialGradient) SetCenter(center unsafe.Pointer) {
	// 0: (, center const QPointF &), (center)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradient9setCenterERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), center)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:286
// index:1
// inline
// void setCenter(qreal, qreal)
func (this *QRadialGradient) SetCenter_1(x float64, y float64) {
	// 1: (, x qreal, y qreal), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradient9setCenterEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:288
// index:0
// QPointF focalPoint()
func (this *QRadialGradient) FocalPoint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QRadialGradient10focalPointEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:289
// index:0
// void setFocalPoint(const class QPointF &)
func (this *QRadialGradient) SetFocalPoint(focalPoint unsafe.Pointer) {
	// 0: (, focalPoint const QPointF &), (focalPoint)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradient13setFocalPointERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), focalPoint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:290
// index:1
// inline
// void setFocalPoint(qreal, qreal)
func (this *QRadialGradient) SetFocalPoint_1(x float64, y float64) {
	// 1: (, x qreal, y qreal), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradient13setFocalPointEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:292
// index:0
// qreal radius()
func (this *QRadialGradient) Radius() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QRadialGradient6radiusEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:293
// index:0
// void setRadius(qreal)
func (this *QRadialGradient) SetRadius(radius float64) {
	// 0: (, radius qreal), (&radius)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradient9setRadiusEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &radius)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:295
// index:0
// qreal centerRadius()
func (this *QRadialGradient) CenterRadius() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QRadialGradient12centerRadiusEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:296
// index:0
// void setCenterRadius(qreal)
func (this *QRadialGradient) SetCenterRadius(radius float64) {
	// 0: (, radius qreal), (&radius)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradient15setCenterRadiusEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &radius)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:298
// index:0
// qreal focalRadius()
func (this *QRadialGradient) FocalRadius() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QRadialGradient11focalRadiusEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:299
// index:0
// void setFocalRadius(qreal)
func (this *QRadialGradient) SetFocalRadius(radius float64) {
	// 0: (, radius qreal), (&radius)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QRadialGradient14setFocalRadiusEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &radius)
	gopp.ErrPrint(err, rv)
}

//  body block end
