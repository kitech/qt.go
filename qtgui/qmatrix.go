package qtgui

// /usr/include/qt/QtGui/qmatrix.h
// #include <qmatrix.h>
// #include <QtGui>

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
import "gopp"
import "qt.go/cffiqt"
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
type QMatrix struct {
	*qtrt.CObject
}

func (this *QMatrix) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMatrix) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQMatrixFromPointer(cthis unsafe.Pointer) *QMatrix {
	return &QMatrix{&qtrt.CObject{cthis}}
}
func (*QMatrix) NewFromPointer(cthis unsafe.Pointer) *QMatrix {
	return NewQMatrixFromPointer(cthis)
}

// /usr/include/qt/QtGui/qmatrix.h:60
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QMatrix(Qt::Initialization)
func NewQMatrix(arg0 int) *QMatrix {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QMatrixC2EN2Qt14InitializationE", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qmatrix.h:61
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QMatrix()
func NewQMatrix_1() *QMatrix {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QMatrixC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qmatrix.h:62
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QMatrix(qreal, qreal, qreal, qreal, qreal, qreal)
func NewQMatrix_2(m11 float64, m12 float64, m21 float64, m22 float64, dx float64, dy float64) *QMatrix {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QMatrixC2Edddddd", ffiqt.FFI_TYPE_POINTER, m11, m12, m21, m22, dx, dy)
	gopp.ErrPrint(err, rv)
	gothis := NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qmatrix.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMatrix(qreal, qreal, qreal, qreal, qreal, qreal)
func (this *QMatrix) SetMatrix(m11 float64, m12 float64, m21 float64, m22 float64, dx float64, dy float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QMatrix9setMatrixEdddddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), m11, m12, m21, m22, dx, dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal m11()
func (this *QMatrix) M11() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3m11Ev", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qmatrix.h:79
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal m12()
func (this *QMatrix) M12() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3m12Ev", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qmatrix.h:80
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal m21()
func (this *QMatrix) M21() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3m21Ev", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qmatrix.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal m22()
func (this *QMatrix) M22() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3m22Ev", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qmatrix.h:82
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal dx()
func (this *QMatrix) Dx() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix2dxEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qmatrix.h:83
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal dy()
func (this *QMatrix) Dy() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix2dyEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qmatrix.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void map(int, int, int *, int *)
func (this *QMatrix) Map(x int, y int, tx unsafe.Pointer /*666*/, ty unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3mapEiiPiS0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, &tx, &ty)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:86
// index:1
// Public Visibility=Default Availability=Available
// [-2] void map(qreal, qreal, qreal *, qreal *)
func (this *QMatrix) Map_1(x float64, y float64, tx unsafe.Pointer /*666*/, ty unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3mapEddPdS0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, &tx, &ty)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:90
// index:2
// Public Visibility=Default Availability=Available
// [8] QPoint map(const QPoint &)
func (this *QMatrix) Map_2(p *qtcore.QPoint) *qtcore.QPoint /*123*/ {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3mapERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:91
// index:3
// Public Visibility=Default Availability=Available
// [16] QPointF map(const QPointF &)
func (this *QMatrix) Map_3(p *qtcore.QPointF) *qtcore.QPointF /*123*/ {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3mapERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:92
// index:4
// Public Visibility=Default Availability=Available
// [16] QLine map(const QLine &)
func (this *QMatrix) Map_4(l *qtcore.QLine) *qtcore.QLine /*123*/ {
	var convArg0 = l.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3mapERK5QLine", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQLineFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:93
// index:5
// Public Visibility=Default Availability=Available
// [32] QLineF map(const QLineF &)
func (this *QMatrix) Map_5(l *qtcore.QLineF) *qtcore.QLineF /*123*/ {
	var convArg0 = l.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3mapERK6QLineF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQLineFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:94
// index:6
// Public Visibility=Default Availability=Available
// [8] QPolygonF map(const QPolygonF &)
func (this *QMatrix) Map_6(a *QPolygonF) *QPolygonF /*123*/ {
	var convArg0 = a.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3mapERK9QPolygonF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:95
// index:7
// Public Visibility=Default Availability=Available
// [8] QPolygon map(const QPolygon &)
func (this *QMatrix) Map_7(a *QPolygon) *QPolygon /*123*/ {
	var convArg0 = a.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3mapERK8QPolygon", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:96
// index:8
// Public Visibility=Default Availability=Available
// [8] QRegion map(const QRegion &)
func (this *QMatrix) Map_8(r *QRegion) *QRegion /*123*/ {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3mapERK7QRegion", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:97
// index:9
// Public Visibility=Default Availability=Available
// [8] QPainterPath map(const QPainterPath &)
func (this *QMatrix) Map_9(p *QPainterPath) *QPainterPath /*123*/ {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix3mapERK12QPainterPath", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:87
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect mapRect(const QRect &)
func (this *QMatrix) MapRect(arg0 *qtcore.QRect) *qtcore.QRect /*123*/ {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix7mapRectERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:88
// index:1
// Public Visibility=Default Availability=Available
// [32] QRectF mapRect(const QRectF &)
func (this *QMatrix) MapRect_1(arg0 *qtcore.QRectF) *qtcore.QRectF /*123*/ {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix7mapRectERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] QPolygon mapToPolygon(const QRect &)
func (this *QMatrix) MapToPolygon(r *qtcore.QRect) *QPolygon /*123*/ {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix12mapToPolygonERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reset()
func (this *QMatrix) Reset() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QMatrix5resetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:101
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isIdentity()
func (this *QMatrix) IsIdentity() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix10isIdentityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qmatrix.h:103
// index:0
// Public Visibility=Default Availability=Available
// [48] QMatrix & translate(qreal, qreal)
func (this *QMatrix) Translate(dx float64, dy float64) *QMatrix {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QMatrix9translateEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:104
// index:0
// Public Visibility=Default Availability=Available
// [48] QMatrix & scale(qreal, qreal)
func (this *QMatrix) Scale(sx float64, sy float64) *QMatrix {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QMatrix5scaleEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), sx, sy)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:105
// index:0
// Public Visibility=Default Availability=Available
// [48] QMatrix & shear(qreal, qreal)
func (this *QMatrix) Shear(sh float64, sv float64) *QMatrix {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QMatrix5shearEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), sh, sv)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:106
// index:0
// Public Visibility=Default Availability=Available
// [48] QMatrix & rotate(qreal)
func (this *QMatrix) Rotate(a float64) *QMatrix {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QMatrix6rotateEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), a)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:108
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isInvertible()
func (this *QMatrix) IsInvertible() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix12isInvertibleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qmatrix.h:109
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal determinant()
func (this *QMatrix) Determinant() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix11determinantEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qmatrix.h:111
// index:0
// Public Visibility=Default Availability=Available
// [48] QMatrix inverted(_Bool *)
func (this *QMatrix) Inverted(invertible unsafe.Pointer /*666*/) *QMatrix /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QMatrix8invertedEPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &invertible)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
