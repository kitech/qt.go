//  header block begin
// /usr/include/qt/QtGui/qtransform.h
// #include <qtransform.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
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
type QTransform struct {
	*qtrt.CObject
}

func (this *QTransform) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtGui/qtransform.h:69
// index:0
// inline
// void QTransform(Qt::Initialization)
func NewQTransform(arg0 int) *QTransform {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTransformC2EN2Qt14InitializationE", ffiqt.FFI_TYPE_VOID, cthis, &arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTransformFromPointer(cthis)
	return gothis
}
func NewQTransformFromPointer(cthis unsafe.Pointer) *QTransform {
	return &QTransform{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qtransform.h:70
// index:1
// void QTransform()
func NewQTransform_1() *QTransform {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTransformC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTransformFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtransform.h:71
// index:2
// void QTransform(qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal)
func NewQTransform_2(h11 float64, h12 float64, h13 float64, h21 float64, h22 float64, h23 float64, h31 float64, h32 float64, h33 float64) *QTransform {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTransformC2Eddddddddd", ffiqt.FFI_TYPE_VOID, cthis, &h11, &h12, &h13, &h21, &h22, &h23, &h31, &h32, &h33)
	gopp.ErrPrint(err, rv)
	gothis := NewQTransformFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtransform.h:74
// index:3
// void QTransform(qreal, qreal, qreal, qreal, qreal, qreal)
func NewQTransform_3(h11 float64, h12 float64, h21 float64, h22 float64, dx float64, dy float64) *QTransform {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTransformC2Edddddd", ffiqt.FFI_TYPE_VOID, cthis, &h11, &h12, &h21, &h22, &dx, &dy)
	gopp.ErrPrint(err, rv)
	gothis := NewQTransformFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtransform.h:76
// index:4
// void QTransform(const class QMatrix &)
func NewQTransform_4(mtx unsafe.Pointer) *QTransform {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTransformC2ERK7QMatrix", ffiqt.FFI_TYPE_VOID, cthis, mtx)
	gopp.ErrPrint(err, rv)
	gothis := NewQTransformFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtransform.h:91
// index:0
// bool isAffine()
func (this *QTransform) IsAffine() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform8isAffineEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:92
// index:0
// bool isIdentity()
func (this *QTransform) IsIdentity() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform10isIdentityEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:93
// index:0
// bool isInvertible()
func (this *QTransform) IsInvertible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform12isInvertibleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:94
// index:0
// bool isScaling()
func (this *QTransform) IsScaling() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform9isScalingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:95
// index:0
// bool isRotating()
func (this *QTransform) IsRotating() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform10isRotatingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:96
// index:0
// bool isTranslating()
func (this *QTransform) IsTranslating() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform13isTranslatingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:98
// index:0
// QTransform::TransformationType type()
func (this *QTransform) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform4typeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:100
// index:0
// inline
// qreal determinant()
func (this *QTransform) Determinant() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform11determinantEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:101
// index:0
// qreal det()
func (this *QTransform) Det() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform3detEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:103
// index:0
// qreal m11()
func (this *QTransform) M11() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform3m11Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:104
// index:0
// qreal m12()
func (this *QTransform) M12() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform3m12Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:105
// index:0
// qreal m13()
func (this *QTransform) M13() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform3m13Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:106
// index:0
// qreal m21()
func (this *QTransform) M21() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform3m21Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:107
// index:0
// qreal m22()
func (this *QTransform) M22() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform3m22Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:108
// index:0
// qreal m23()
func (this *QTransform) M23() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform3m23Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:109
// index:0
// qreal m31()
func (this *QTransform) M31() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform3m31Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:110
// index:0
// qreal m32()
func (this *QTransform) M32() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform3m32Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:111
// index:0
// qreal m33()
func (this *QTransform) M33() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform3m33Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:112
// index:0
// qreal dx()
func (this *QTransform) Dx() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform2dxEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:113
// index:0
// qreal dy()
func (this *QTransform) Dy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform2dyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:115
// index:0
// void setMatrix(qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal)
func (this *QTransform) SetMatrix(m11 float64, m12 float64, m13 float64, m21 float64, m22 float64, m23 float64, m31 float64, m32 float64, m33 float64) {
	// 0: (, m11 qreal, m12 qreal, m13 qreal, m21 qreal, m22 qreal, m23 qreal, m31 qreal, m32 qreal, m33 qreal), (&m11, &m12, &m13, &m21, &m22, &m23, &m31, &m32, &m33)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTransform9setMatrixEddddddddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &m11, &m12, &m13, &m21, &m22, &m23, &m31, &m32, &m33)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:119
// index:0
// QTransform inverted(_Bool *)
func (this *QTransform) Inverted(invertible unsafe.Pointer) {
	// 0: (, invertible bool *), (invertible)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform8invertedEPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), invertible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:120
// index:0
// QTransform adjoint()
func (this *QTransform) Adjoint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform7adjointEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:121
// index:0
// QTransform transposed()
func (this *QTransform) Transposed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform10transposedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:123
// index:0
// QTransform & translate(qreal, qreal)
func (this *QTransform) Translate(dx float64, dy float64) {
	// 0: (, dx qreal, dy qreal), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTransform9translateEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:124
// index:0
// QTransform & scale(qreal, qreal)
func (this *QTransform) Scale(sx float64, sy float64) {
	// 0: (, sx qreal, sy qreal), (&sx, &sy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTransform5scaleEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &sx, &sy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:125
// index:0
// QTransform & shear(qreal, qreal)
func (this *QTransform) Shear(sh float64, sv float64) {
	// 0: (, sh qreal, sv qreal), (&sh, &sv)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTransform5shearEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &sh, &sv)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:126
// index:0
// QTransform & rotate(qreal, Qt::Axis)
func (this *QTransform) Rotate(a float64, axis int) {
	// 0: (, a qreal, axis Qt::Axis), (&a, &axis)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTransform6rotateEdN2Qt4AxisE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &a, &axis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:127
// index:0
// QTransform & rotateRadians(qreal, Qt::Axis)
func (this *QTransform) RotateRadians(a float64, axis int) {
	// 0: (, a qreal, axis Qt::Axis), (&a, &axis)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTransform13rotateRadiansEdN2Qt4AxisE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &a, &axis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:129
// index:0
// static
// bool squareToQuad(const class QPolygonF &, class QTransform &)
func (this *QTransform) SquareToQuad(square unsafe.Pointer, result unsafe.Pointer) {
	// 0: (square const QPolygonF &, result QTransform &), (square, result)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTransform12squareToQuadERK9QPolygonFRS_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTransform_SquareToQuad(square unsafe.Pointer, result unsafe.Pointer) {
	// 0: (square const QPolygonF &, result QTransform &), (square, result)
	var nilthis *QTransform
	nilthis.SquareToQuad(square, result)
}

// /usr/include/qt/QtGui/qtransform.h:130
// index:0
// static
// bool quadToSquare(const class QPolygonF &, class QTransform &)
func (this *QTransform) QuadToSquare(quad unsafe.Pointer, result unsafe.Pointer) {
	// 0: (quad const QPolygonF &, result QTransform &), (quad, result)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTransform12quadToSquareERK9QPolygonFRS_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTransform_QuadToSquare(quad unsafe.Pointer, result unsafe.Pointer) {
	// 0: (quad const QPolygonF &, result QTransform &), (quad, result)
	var nilthis *QTransform
	nilthis.QuadToSquare(quad, result)
}

// /usr/include/qt/QtGui/qtransform.h:131
// index:0
// static
// bool quadToQuad(const class QPolygonF &, const class QPolygonF &, class QTransform &)
func (this *QTransform) QuadToQuad(one unsafe.Pointer, two unsafe.Pointer, result unsafe.Pointer) {
	// 0: (one const QPolygonF &, two const QPolygonF &, result QTransform &), (one, two, result)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTransform10quadToQuadERK9QPolygonFS2_RS_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTransform_QuadToQuad(one unsafe.Pointer, two unsafe.Pointer, result unsafe.Pointer) {
	// 0: (one const QPolygonF &, two const QPolygonF &, result QTransform &), (one, two, result)
	var nilthis *QTransform
	nilthis.QuadToQuad(one, two, result)
}

// /usr/include/qt/QtGui/qtransform.h:143
// index:0
// void reset()
func (this *QTransform) Reset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTransform5resetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:144
// index:0
// QPoint map(const class QPoint &)
func (this *QTransform) Map(p unsafe.Pointer) {
	// 0: (, p const QPoint &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform3mapERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:145
// index:1
// QPointF map(const class QPointF &)
func (this *QTransform) Map_1(p unsafe.Pointer) {
	// 1: (, p const QPointF &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform3mapERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:146
// index:2
// QLine map(const class QLine &)
func (this *QTransform) Map_2(l unsafe.Pointer) {
	// 2: (, l const QLine &), (l)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform3mapERK5QLine", ffiqt.FFI_TYPE_VOID, this.GetCthis(), l)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:147
// index:3
// QLineF map(const class QLineF &)
func (this *QTransform) Map_3(l unsafe.Pointer) {
	// 3: (, l const QLineF &), (l)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform3mapERK6QLineF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), l)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:148
// index:4
// QPolygonF map(const class QPolygonF &)
func (this *QTransform) Map_4(a unsafe.Pointer) {
	// 4: (, a const QPolygonF &), (a)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform3mapERK9QPolygonF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:149
// index:5
// QPolygon map(const class QPolygon &)
func (this *QTransform) Map_5(a unsafe.Pointer) {
	// 5: (, a const QPolygon &), (a)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform3mapERK8QPolygon", ffiqt.FFI_TYPE_VOID, this.GetCthis(), a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:150
// index:6
// QRegion map(const class QRegion &)
func (this *QTransform) Map_6(r unsafe.Pointer) {
	// 6: (, r const QRegion &), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform3mapERK7QRegion", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:151
// index:7
// QPainterPath map(const class QPainterPath &)
func (this *QTransform) Map_7(p unsafe.Pointer) {
	// 7: (, p const QPainterPath &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform3mapERK12QPainterPath", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:155
// index:8
// void map(int, int, int *, int *)
func (this *QTransform) Map_8(x int, y int, tx unsafe.Pointer, ty unsafe.Pointer) {
	// 8: (, x int, y int, tx int *, ty int *), (&x, &y, tx, ty)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform3mapEiiPiS0_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, tx, ty)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:156
// index:9
// void map(qreal, qreal, qreal *, qreal *)
func (this *QTransform) Map_9(x float64, y float64, tx unsafe.Pointer, ty unsafe.Pointer) {
	// 9: (, x qreal, y qreal, tx qreal *, ty qreal *), (&x, &y, tx, ty)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform3mapEddPdS0_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, tx, ty)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:152
// index:0
// QPolygon mapToPolygon(const class QRect &)
func (this *QTransform) MapToPolygon(r unsafe.Pointer) {
	// 0: (, r const QRect &), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform12mapToPolygonERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:153
// index:0
// QRect mapRect(const class QRect &)
func (this *QTransform) MapRect(arg0 unsafe.Pointer) {
	// 0: (, const QRect & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform7mapRectERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:154
// index:1
// QRectF mapRect(const class QRectF &)
func (this *QTransform) MapRect_1(arg0 unsafe.Pointer) {
	// 1: (, const QRectF & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform7mapRectERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:158
// index:0
// const QMatrix & toAffine()
func (this *QTransform) ToAffine() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTransform8toAffineEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:165
// index:0
// static
// QTransform fromTranslate(qreal, qreal)
func (this *QTransform) FromTranslate(dx float64, dy float64) {
	// 0: (dx qreal, dy qreal), (dx, dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTransform13fromTranslateEdd", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTransform_FromTranslate(dx float64, dy float64) {
	// 0: (dx qreal, dy qreal), (dx, dy)
	var nilthis *QTransform
	nilthis.FromTranslate(dx, dy)
}

// /usr/include/qt/QtGui/qtransform.h:166
// index:0
// static
// QTransform fromScale(qreal, qreal)
func (this *QTransform) FromScale(dx float64, dy float64) {
	// 0: (dx qreal, dy qreal), (dx, dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTransform9fromScaleEdd", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTransform_FromScale(dx float64, dy float64) {
	// 0: (dx qreal, dy qreal), (dx, dy)
	var nilthis *QTransform
	nilthis.FromScale(dx, dy)
}

//  body block end
