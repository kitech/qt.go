package qtgui

// /usr/include/qt/QtGui/qmatrix.h
// #include <qmatrix.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
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
type QMatrix struct {
	*qtrt.CObject
}
type QMatrix_ITF interface {
	QMatrix_PTR() *QMatrix
}

func (ptr *QMatrix) QMatrix_PTR() *QMatrix { return ptr }

func (this *QMatrix) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMatrix) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
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

/*
Constructs an identity matrix.

All elements are set to zero except m11 and m22 (specifying the scale), which are set to 1.

See also reset().
*/
func (*QMatrix) NewForInherit(arg0 int) *QMatrix {
	return NewQMatrix(arg0)
}
func NewQMatrix(arg0 int) *QMatrix {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QMatrixC2EN2Qt14InitializationE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMatrix)
	return gothis
}

// /usr/include/qt/QtGui/qmatrix.h:61
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QMatrix()

/*
Constructs an identity matrix.

All elements are set to zero except m11 and m22 (specifying the scale), which are set to 1.

See also reset().
*/
func (*QMatrix) NewForInherit1() *QMatrix {
	return NewQMatrix1()
}
func NewQMatrix1() *QMatrix {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QMatrixC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMatrix)
	return gothis
}

// /usr/include/qt/QtGui/qmatrix.h:62
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QMatrix(qreal, qreal, qreal, qreal, qreal, qreal)

/*
Constructs an identity matrix.

All elements are set to zero except m11 and m22 (specifying the scale), which are set to 1.

See also reset().
*/
func (*QMatrix) NewForInherit2(m11 float64, m12 float64, m21 float64, m22 float64, dx float64, dy float64) *QMatrix {
	return NewQMatrix2(m11, m12, m21, m22, dx, dy)
}
func NewQMatrix2(m11 float64, m12 float64, m21 float64, m22 float64, dx float64, dy float64) *QMatrix {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QMatrixC2Edddddd", qtrt.FFI_TYPE_POINTER, m11, m12, m21, m22, dx, dy)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMatrix)
	return gothis
}

// /usr/include/qt/QtGui/qmatrix.h:67
// index:0
// Public inline Visibility=Default Availability=Available
// [48] QMatrix & operator=(QMatrix &&)

/*

 */
func (this *QMatrix) Operator_equal(other unsafe.Pointer /*333*/) *QMatrix {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QMatrixaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:69
// index:1
// Public Visibility=Default Availability=Available
// [48] QMatrix & operator=(const QMatrix &)

/*

 */
func (this *QMatrix) Operator_equal1(arg0 QMatrix_ITF) *QMatrix {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMatrix_PTR() != nil {
		convArg0 = arg0.QMatrix_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QMatrixaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMatrix(qreal, qreal, qreal, qreal, qreal, qreal)

/*
Sets the matrix elements to the specified values, m11, m12, m21, m22, dx and dy.

Note that this function replaces the previous values. QMatrix provide the translate(), rotate(), scale() and shear() convenience functions to manipulate the various matrix elements based on the currently defined coordinate system.

See also QMatrix().
*/
func (this *QMatrix) SetMatrix(m11 float64, m12 float64, m21 float64, m22 float64, dx float64, dy float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QMatrix9setMatrixEdddddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), m11, m12, m21, m22, dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal m11() const

/*
Returns the horizontal scaling factor.

See also scale() and Basic Matrix Operations.
*/
func (this *QMatrix) M11() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QMatrix3m11Ev", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qmatrix.h:79
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal m12() const

/*
Returns the vertical shearing factor.

See also shear() and Basic Matrix Operations.
*/
func (this *QMatrix) M12() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QMatrix3m12Ev", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qmatrix.h:80
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal m21() const

/*
Returns the horizontal shearing factor.

See also shear() and Basic Matrix Operations.
*/
func (this *QMatrix) M21() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QMatrix3m21Ev", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qmatrix.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal m22() const

/*
Returns the vertical scaling factor.

See also scale() and Basic Matrix Operations.
*/
func (this *QMatrix) M22() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QMatrix3m22Ev", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qmatrix.h:82
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal dx() const

/*
Returns the horizontal translation factor.

See also translate() and Basic Matrix Operations.
*/
func (this *QMatrix) Dx() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QMatrix2dxEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qmatrix.h:83
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal dy() const

/*
Returns the vertical translation factor.

See also translate() and Basic Matrix Operations.
*/
func (this *QMatrix) Dy() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QMatrix2dyEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qmatrix.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void map(int, int, int *, int *) const

/*
Maps the given coordinates x and y into the coordinate system defined by this matrix. The resulting values are put in *tx and *ty, respectively.

The coordinates are transformed using the following formulas:


  x' = m11*x + m21*y + dx
  y' = m22*y + m12*x + dy



The point (x, y) is the original point, and (x', y') is the transformed point.

See also Basic Matrix Operations.
*/
func (this *QMatrix) Map(x int, y int, tx unsafe.Pointer /*666*/, ty unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QMatrix3mapEiiPiS0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, tx, ty)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:86
// index:1
// Public Visibility=Default Availability=Available
// [-2] void map(qreal, qreal, qreal *, qreal *) const

/*
Maps the given coordinates x and y into the coordinate system defined by this matrix. The resulting values are put in *tx and *ty, respectively.

The coordinates are transformed using the following formulas:


  x' = m11*x + m21*y + dx
  y' = m22*y + m12*x + dy



The point (x, y) is the original point, and (x', y') is the transformed point.

See also Basic Matrix Operations.
*/
func (this *QMatrix) Map1(x float64, y float64, tx unsafe.Pointer /*666*/, ty unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QMatrix3mapEddPdS0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, tx, ty)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:90
// index:2
// Public Visibility=Default Availability=Available
// [8] QPoint map(const QPoint &) const

/*
Maps the given coordinates x and y into the coordinate system defined by this matrix. The resulting values are put in *tx and *ty, respectively.

The coordinates are transformed using the following formulas:


  x' = m11*x + m21*y + dx
  y' = m22*y + m12*x + dy



The point (x, y) is the original point, and (x', y') is the transformed point.

See also Basic Matrix Operations.
*/
func (this *QMatrix) Map2(p qtcore.QPoint_ITF) *qtcore.QPoint /*123*/ {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QMatrix3mapERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:91
// index:3
// Public Visibility=Default Availability=Available
// [16] QPointF map(const QPointF &) const

/*
Maps the given coordinates x and y into the coordinate system defined by this matrix. The resulting values are put in *tx and *ty, respectively.

The coordinates are transformed using the following formulas:


  x' = m11*x + m21*y + dx
  y' = m22*y + m12*x + dy



The point (x, y) is the original point, and (x', y') is the transformed point.

See also Basic Matrix Operations.
*/
func (this *QMatrix) Map3(p qtcore.QPointF_ITF) *qtcore.QPointF /*123*/ {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QMatrix3mapERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:92
// index:4
// Public Visibility=Default Availability=Available
// [16] QLine map(const QLine &) const

/*
Maps the given coordinates x and y into the coordinate system defined by this matrix. The resulting values are put in *tx and *ty, respectively.

The coordinates are transformed using the following formulas:


  x' = m11*x + m21*y + dx
  y' = m22*y + m12*x + dy



The point (x, y) is the original point, and (x', y') is the transformed point.

See also Basic Matrix Operations.
*/
func (this *QMatrix) Map4(l qtcore.QLine_ITF) *qtcore.QLine /*123*/ {
	var convArg0 unsafe.Pointer
	if l != nil && l.QLine_PTR() != nil {
		convArg0 = l.QLine_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QMatrix3mapERK5QLine", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQLineFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQLine)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:93
// index:5
// Public Visibility=Default Availability=Available
// [32] QLineF map(const QLineF &) const

/*
Maps the given coordinates x and y into the coordinate system defined by this matrix. The resulting values are put in *tx and *ty, respectively.

The coordinates are transformed using the following formulas:


  x' = m11*x + m21*y + dx
  y' = m22*y + m12*x + dy



The point (x, y) is the original point, and (x', y') is the transformed point.

See also Basic Matrix Operations.
*/
func (this *QMatrix) Map5(l qtcore.QLineF_ITF) *qtcore.QLineF /*123*/ {
	var convArg0 unsafe.Pointer
	if l != nil && l.QLineF_PTR() != nil {
		convArg0 = l.QLineF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QMatrix3mapERK6QLineF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQLineFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQLineF)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:94
// index:6
// Public Visibility=Default Availability=Available
// [8] QPolygonF map(const QPolygonF &) const

/*
Maps the given coordinates x and y into the coordinate system defined by this matrix. The resulting values are put in *tx and *ty, respectively.

The coordinates are transformed using the following formulas:


  x' = m11*x + m21*y + dx
  y' = m22*y + m12*x + dy



The point (x, y) is the original point, and (x', y') is the transformed point.

See also Basic Matrix Operations.
*/
func (this *QMatrix) Map6(a QPolygonF_ITF) *QPolygonF /*123*/ {
	var convArg0 unsafe.Pointer
	if a != nil && a.QPolygonF_PTR() != nil {
		convArg0 = a.QPolygonF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QMatrix3mapERK9QPolygonF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:95
// index:7
// Public Visibility=Default Availability=Available
// [8] QPolygon map(const QPolygon &) const

/*
Maps the given coordinates x and y into the coordinate system defined by this matrix. The resulting values are put in *tx and *ty, respectively.

The coordinates are transformed using the following formulas:


  x' = m11*x + m21*y + dx
  y' = m22*y + m12*x + dy



The point (x, y) is the original point, and (x', y') is the transformed point.

See also Basic Matrix Operations.
*/
func (this *QMatrix) Map7(a QPolygon_ITF) *QPolygon /*123*/ {
	var convArg0 unsafe.Pointer
	if a != nil && a.QPolygon_PTR() != nil {
		convArg0 = a.QPolygon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QMatrix3mapERK8QPolygon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygon)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:96
// index:8
// Public Visibility=Default Availability=Available
// [8] QRegion map(const QRegion &) const

/*
Maps the given coordinates x and y into the coordinate system defined by this matrix. The resulting values are put in *tx and *ty, respectively.

The coordinates are transformed using the following formulas:


  x' = m11*x + m21*y + dx
  y' = m22*y + m12*x + dy



The point (x, y) is the original point, and (x', y') is the transformed point.

See also Basic Matrix Operations.
*/
func (this *QMatrix) Map8(r QRegion_ITF) *QRegion /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRegion_PTR() != nil {
		convArg0 = r.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QMatrix3mapERK7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:97
// index:9
// Public Visibility=Default Availability=Available
// [8] QPainterPath map(const QPainterPath &) const

/*
Maps the given coordinates x and y into the coordinate system defined by this matrix. The resulting values are put in *tx and *ty, respectively.

The coordinates are transformed using the following formulas:


  x' = m11*x + m21*y + dx
  y' = m22*y + m12*x + dy



The point (x, y) is the original point, and (x', y') is the transformed point.

See also Basic Matrix Operations.
*/
func (this *QMatrix) Map9(p QPainterPath_ITF) *QPainterPath /*123*/ {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPainterPath_PTR() != nil {
		convArg0 = p.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QMatrix3mapERK12QPainterPath", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:87
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect mapRect(const QRect &) const

/*
Creates and returns a QRectF object that is a copy of the given rectangle, mapped into the coordinate system defined by this matrix.

The rectangle's coordinates are transformed using the following formulas:


  x' = m11*x + m21*y + dx
  y' = m22*y + m12*x + dy



If rotation or shearing has been specified, this function returns the bounding rectangle. To retrieve the exact region the given rectangle maps to, use the mapToPolygon() function instead.

See also mapToPolygon() and Basic Matrix Operations.
*/
func (this *QMatrix) MapRect(arg0 qtcore.QRect_ITF) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRect_PTR() != nil {
		convArg0 = arg0.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QMatrix7mapRectERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:88
// index:1
// Public Visibility=Default Availability=Available
// [32] QRectF mapRect(const QRectF &) const

/*
Creates and returns a QRectF object that is a copy of the given rectangle, mapped into the coordinate system defined by this matrix.

The rectangle's coordinates are transformed using the following formulas:


  x' = m11*x + m21*y + dx
  y' = m22*y + m12*x + dy



If rotation or shearing has been specified, this function returns the bounding rectangle. To retrieve the exact region the given rectangle maps to, use the mapToPolygon() function instead.

See also mapToPolygon() and Basic Matrix Operations.
*/
func (this *QMatrix) MapRect1(arg0 qtcore.QRectF_ITF) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRectF_PTR() != nil {
		convArg0 = arg0.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QMatrix7mapRectERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] QPolygon mapToPolygon(const QRect &) const

/*
Creates and returns a QPolygon representation of the given rectangle, mapped into the coordinate system defined by this matrix.

The rectangle's coordinates are transformed using the following formulas:


  x' = m11*x + m21*y + dx
  y' = m22*y + m12*x + dy



Polygons and rectangles behave slightly differently when transformed (due to integer rounding), so matrix.map(QPolygon(rectangle)) is not always the same as matrix.mapToPolygon(rectangle).

See also mapRect() and Basic Matrix Operations.
*/
func (this *QMatrix) MapToPolygon(r qtcore.QRect_ITF) *QPolygon /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QMatrix12mapToPolygonERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygon)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reset()

/*
Resets the matrix to an identity matrix, i.e. all elements are set to zero, except m11 and m22 (specifying the scale) which are set to 1.

See also QMatrix(), isIdentity(), and Basic Matrix Operations.
*/
func (this *QMatrix) Reset() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QMatrix5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix.h:101
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isIdentity() const

/*
Returns true if the matrix is the identity matrix, otherwise returns false.

See also reset().
*/
func (this *QMatrix) IsIdentity() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QMatrix10isIdentityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qmatrix.h:103
// index:0
// Public Visibility=Default Availability=Available
// [48] QMatrix & translate(qreal, qreal)

/*
Moves the coordinate system dx along the x axis and dy along the y axis, and returns a reference to the matrix.

See also setMatrix().
*/
func (this *QMatrix) Translate(dx float64, dy float64) *QMatrix {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QMatrix9translateEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:104
// index:0
// Public Visibility=Default Availability=Available
// [48] QMatrix & scale(qreal, qreal)

/*
Scales the coordinate system by sx horizontally and sy vertically, and returns a reference to the matrix.

See also setMatrix().
*/
func (this *QMatrix) Scale(sx float64, sy float64) *QMatrix {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QMatrix5scaleEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sx, sy)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:105
// index:0
// Public Visibility=Default Availability=Available
// [48] QMatrix & shear(qreal, qreal)

/*
Shears the coordinate system by sh horizontally and sv vertically, and returns a reference to the matrix.

See also setMatrix().
*/
func (this *QMatrix) Shear(sh float64, sv float64) *QMatrix {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QMatrix5shearEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sh, sv)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:106
// index:0
// Public Visibility=Default Availability=Available
// [48] QMatrix & rotate(qreal)

/*
Rotates the coordinate system the given degrees counterclockwise.

Note that if you apply a QMatrix to a point defined in widget coordinates, the direction of the rotation will be clockwise because the y-axis points downwards.

Returns a reference to the matrix.

See also setMatrix().
*/
func (this *QMatrix) Rotate(a float64) *QMatrix {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QMatrix6rotateEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), a)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:108
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isInvertible() const

/*
Returns true if the matrix is invertible, otherwise returns false.

See also inverted().
*/
func (this *QMatrix) IsInvertible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QMatrix12isInvertibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qmatrix.h:109
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal determinant() const

/*
Returns the matrix's determinant.

This function was introduced in  Qt 4.6.
*/
func (this *QMatrix) Determinant() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QMatrix11determinantEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qmatrix.h:111
// index:0
// Public Visibility=Default Availability=Available
// [48] QMatrix inverted(bool *) const

/*
Returns an inverted copy of this matrix.

If the matrix is singular (not invertible), the returned matrix is the identity matrix. If invertible is valid (i.e. not 0), its value is set to true if the matrix is invertible, otherwise it is set to false.

See also isInvertible().
*/
func (this *QMatrix) Inverted(invertible *bool) *QMatrix /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QMatrix8invertedEPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), invertible)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:111
// index:0
// Public Visibility=Default Availability=Available
// [48] QMatrix inverted(bool *) const

/*
Returns an inverted copy of this matrix.

If the matrix is singular (not invertible), the returned matrix is the identity matrix. If invertible is valid (i.e. not 0), its value is set to true if the matrix is invertible, otherwise it is set to false.

See also isInvertible().
*/
func (this *QMatrix) Invertedp() *QMatrix /*123*/ {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var invertible unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QMatrix8invertedEPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), invertible)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QMatrix &) const

/*

 */
func (this *QMatrix) Operator_equal_equal(arg0 QMatrix_ITF) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMatrix_PTR() != nil {
		convArg0 = arg0.QMatrix_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QMatrixeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qmatrix.h:114
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QMatrix &) const

/*

 */
func (this *QMatrix) Operator_not_equal(arg0 QMatrix_ITF) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMatrix_PTR() != nil {
		convArg0 = arg0.QMatrix_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QMatrixneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qmatrix.h:116
// index:0
// Public Visibility=Default Availability=Available
// [48] QMatrix & operator*=(const QMatrix &)

/*

 */
func (this *QMatrix) Operator_mul_equal(arg0 QMatrix_ITF) *QMatrix {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMatrix_PTR() != nil {
		convArg0 = arg0.QMatrix_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QMatrixmLERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix.h:117
// index:0
// Public Visibility=Default Availability=Available
// [48] QMatrix operator*(const QMatrix &) const

/*

 */
func (this *QMatrix) Operator_mul(o QMatrix_ITF) *QMatrix /*123*/ {
	var convArg0 unsafe.Pointer
	if o != nil && o.QMatrix_PTR() != nil {
		convArg0 = o.QMatrix_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QMatrixmlERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix)
	return rv2
}

func DeleteQMatrix(this *QMatrix) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QMatrixD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
