//  header block begin
// /usr/include/qt/QtGui/qmatrix.h
// #include <qmatrix.h>
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
type QMatrix struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qmatrix.h:60
// index:0
// inline
// void QMatrix(Qt::Initialization)
func NewQMatrix(arg0 int) *QMatrix {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QMatrixC2EN2Qt14InitializationE", ffiqt.FFI_TYPE_VOID, cthis, &arg0)
	gopp.ErrPrint(err, rv)
	return &QMatrix{cthis}
}

// /usr/include/qt/QtGui/qmatrix.h:61
// index:1
// void QMatrix()
func NewQMatrix_1() *QMatrix {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QMatrixC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QMatrix{cthis}
}

// /usr/include/qt/QtGui/qmatrix.h:62
// index:2
// void QMatrix(qreal, qreal, qreal, qreal, qreal, qreal)
func NewQMatrix_2(m11 float64, m12 float64, m21 float64, m22 float64, dx float64, dy float64) *QMatrix {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QMatrixC2Edddddd", ffiqt.FFI_TYPE_VOID, cthis, &m11, &m12, &m21, &m22, &dx, &dy)
	gopp.ErrPrint(err, rv)
	return &QMatrix{cthis}
}

// /usr/include/qt/QtGui/qmatrix.h:75
// index:0
// void setMatrix(qreal, qreal, qreal, qreal, qreal, qreal)
func (this *QMatrix) SetMatrix(m11 float64, m12 float64, m21 float64, m22 float64, dx float64, dy float64) {
	// 0: (, qreal m11, qreal m12, qreal m21, qreal m22, qreal dx, qreal dy), (&m11, &m12, &m21, &m22, &dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QMatrix9setMatrixEdddddd", ffiqt.FFI_TYPE_VOID, this.cthis, &m11, &m12, &m21, &m22, &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:78
// index:0
// inline
// qreal m11()
func (this *QMatrix) M11() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3m11Ev", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:79
// index:0
// inline
// qreal m12()
func (this *QMatrix) M12() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3m12Ev", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:80
// index:0
// inline
// qreal m21()
func (this *QMatrix) M21() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3m21Ev", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:81
// index:0
// inline
// qreal m22()
func (this *QMatrix) M22() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3m22Ev", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:82
// index:0
// inline
// qreal dx()
func (this *QMatrix) Dx() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix2dxEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:83
// index:0
// inline
// qreal dy()
func (this *QMatrix) Dy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix2dyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:85
// index:0
// void map(int, int, int *, int *)
func (this *QMatrix) Map(x int, y int, tx unsafe.Pointer, ty unsafe.Pointer) {
	// 0: (, int x, int y, int * tx, int * ty), (&x, &y, tx, ty)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3mapEiiPiS0_", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y, tx, ty)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:86
// index:1
// void map(qreal, qreal, qreal *, qreal *)
func (this *QMatrix) Map_1(x float64, y float64, tx unsafe.Pointer, ty unsafe.Pointer) {
	// 1: (, qreal x, qreal y, qreal * tx, qreal * ty), (&x, &y, tx, ty)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3mapEddPdS0_", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y, tx, ty)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:90
// index:2
// QPoint map(const class QPoint &)
func (this *QMatrix) Map_2(p unsafe.Pointer) {
	// 2: (, const QPoint & p), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3mapERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:91
// index:3
// QPointF map(const class QPointF &)
func (this *QMatrix) Map_3(p unsafe.Pointer) {
	// 3: (, const QPointF & p), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3mapERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:92
// index:4
// QLine map(const class QLine &)
func (this *QMatrix) Map_4(l unsafe.Pointer) {
	// 4: (, const QLine & l), (l)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3mapERK5QLine", ffiqt.FFI_TYPE_VOID, this.cthis, l)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:93
// index:5
// QLineF map(const class QLineF &)
func (this *QMatrix) Map_5(l unsafe.Pointer) {
	// 5: (, const QLineF & l), (l)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3mapERK6QLineF", ffiqt.FFI_TYPE_VOID, this.cthis, l)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:94
// index:6
// QPolygonF map(const class QPolygonF &)
func (this *QMatrix) Map_6(a unsafe.Pointer) {
	// 6: (, const QPolygonF & a), (a)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3mapERK9QPolygonF", ffiqt.FFI_TYPE_VOID, this.cthis, a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:95
// index:7
// QPolygon map(const class QPolygon &)
func (this *QMatrix) Map_7(a unsafe.Pointer) {
	// 7: (, const QPolygon & a), (a)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3mapERK8QPolygon", ffiqt.FFI_TYPE_VOID, this.cthis, a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:96
// index:8
// QRegion map(const class QRegion &)
func (this *QMatrix) Map_8(r unsafe.Pointer) {
	// 8: (, const QRegion & r), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3mapERK7QRegion", ffiqt.FFI_TYPE_VOID, this.cthis, r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:97
// index:9
// QPainterPath map(const class QPainterPath &)
func (this *QMatrix) Map_9(p unsafe.Pointer) {
	// 9: (, const QPainterPath & p), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3mapERK12QPainterPath", ffiqt.FFI_TYPE_VOID, this.cthis, p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:87
// index:0
// QRect mapRect(const class QRect &)
func (this *QMatrix) MapRect(arg0 unsafe.Pointer) {
	// 0: (, const QRect & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix7mapRectERK5QRect", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:88
// index:1
// QRectF mapRect(const class QRectF &)
func (this *QMatrix) MapRect_1(arg0 unsafe.Pointer) {
	// 1: (, const QRectF & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix7mapRectERK6QRectF", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:98
// index:0
// QPolygon mapToPolygon(const class QRect &)
func (this *QMatrix) MapToPolygon(r unsafe.Pointer) {
	// 0: (, const QRect & r), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix12mapToPolygonERK5QRect", ffiqt.FFI_TYPE_VOID, this.cthis, r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:100
// index:0
// void reset()
func (this *QMatrix) Reset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QMatrix5resetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:101
// index:0
// inline
// bool isIdentity()
func (this *QMatrix) IsIdentity() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix10isIdentityEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:103
// index:0
// QMatrix & translate(qreal, qreal)
func (this *QMatrix) Translate(dx float64, dy float64) {
	// 0: (, qreal dx, qreal dy), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QMatrix9translateEdd", ffiqt.FFI_TYPE_VOID, this.cthis, &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:104
// index:0
// QMatrix & scale(qreal, qreal)
func (this *QMatrix) Scale(sx float64, sy float64) {
	// 0: (, qreal sx, qreal sy), (&sx, &sy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QMatrix5scaleEdd", ffiqt.FFI_TYPE_VOID, this.cthis, &sx, &sy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:105
// index:0
// QMatrix & shear(qreal, qreal)
func (this *QMatrix) Shear(sh float64, sv float64) {
	// 0: (, qreal sh, qreal sv), (&sh, &sv)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QMatrix5shearEdd", ffiqt.FFI_TYPE_VOID, this.cthis, &sh, &sv)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:106
// index:0
// QMatrix & rotate(qreal)
func (this *QMatrix) Rotate(a float64) {
	// 0: (, qreal a), (&a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QMatrix6rotateEd", ffiqt.FFI_TYPE_VOID, this.cthis, &a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:108
// index:0
// inline
// bool isInvertible()
func (this *QMatrix) IsInvertible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix12isInvertibleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:109
// index:0
// inline
// qreal determinant()
func (this *QMatrix) Determinant() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix11determinantEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:111
// index:0
// QMatrix inverted(_Bool *)
func (this *QMatrix) Inverted(invertible unsafe.Pointer) {
	// 0: (, bool * invertible), (invertible)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix8invertedEPb", ffiqt.FFI_TYPE_VOID, this.cthis, invertible)
	gopp.ErrPrint(err, rv)
}

//  body block end
