package qtgui

// /usr/include/qt/QtGui/qmatrix4x4.h
// #include <qmatrix4x4.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 47
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
type QMatrix4x4 struct {
	*qtrt.CObject
}
type QMatrix4x4_ITF interface {
	QMatrix4x4_PTR() *QMatrix4x4
}

func (ptr *QMatrix4x4) QMatrix4x4_PTR() *QMatrix4x4 { return ptr }

func (this *QMatrix4x4) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMatrix4x4) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMatrix4x4FromPointer(cthis unsafe.Pointer) *QMatrix4x4 {
	return &QMatrix4x4{&qtrt.CObject{cthis}}
}
func (*QMatrix4x4) NewFromPointer(cthis unsafe.Pointer) *QMatrix4x4 {
	return NewQMatrix4x4FromPointer(cthis)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:62
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QMatrix4x4()

/*
Constructs an identity matrix.
*/
func NewQMatrix4x4() *QMatrix4x4 {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x4C2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMatrix4x4)
	return gothis
}

// /usr/include/qt/QtGui/qmatrix4x4.h:63
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QMatrix4x4(Qt::Initialization)

/*
Constructs an identity matrix.
*/
func NewQMatrix4x4_1(arg0 int) *QMatrix4x4 {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x4C2EN2Qt14InitializationE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMatrix4x4)
	return gothis
}

// /usr/include/qt/QtGui/qmatrix4x4.h:64
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QMatrix4x4(const float *)

/*
Constructs an identity matrix.
*/
func NewQMatrix4x4_2(values unsafe.Pointer /*666*/) *QMatrix4x4 {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x4C2EPKf", qtrt.FFI_TYPE_POINTER, values)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMatrix4x4)
	return gothis
}

// /usr/include/qt/QtGui/qmatrix4x4.h:65
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void QMatrix4x4(float, float, float, float, float, float, float, float, float, float, float, float, float, float, float, float)

/*
Constructs an identity matrix.
*/
func NewQMatrix4x4_3(m11 float32, m12 float32, m13 float32, m14 float32, m21 float32, m22 float32, m23 float32, m24 float32, m31 float32, m32 float32, m33 float32, m34 float32, m41 float32, m42 float32, m43 float32, m44 float32) *QMatrix4x4 {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x4C2Effffffffffffffff", qtrt.FFI_TYPE_POINTER, m11, m12, m13, m14, m21, m22, m23, m24, m31, m32, m33, m34, m41, m42, m43, m44)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMatrix4x4)
	return gothis
}

// /usr/include/qt/QtGui/qmatrix4x4.h:73
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QMatrix4x4(const float *, int, int)

/*
Constructs an identity matrix.
*/
func NewQMatrix4x4_4(values unsafe.Pointer /*666*/, cols int, rows int) *QMatrix4x4 {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x4C2EPKfii", qtrt.FFI_TYPE_POINTER, values, cols, rows)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMatrix4x4)
	return gothis
}

// /usr/include/qt/QtGui/qmatrix4x4.h:74
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QMatrix4x4(const QTransform &)

/*
Constructs an identity matrix.
*/
func NewQMatrix4x4_5(transform QTransform_ITF) *QMatrix4x4 {
	var convArg0 unsafe.Pointer
	if transform != nil && transform.QTransform_PTR() != nil {
		convArg0 = transform.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x4C2ERK10QTransform", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMatrix4x4)
	return gothis
}

// /usr/include/qt/QtGui/qmatrix4x4.h:75
// index:6
// Public Visibility=Default Availability=Available
// [-2] void QMatrix4x4(const QMatrix &)

/*
Constructs an identity matrix.
*/
func NewQMatrix4x4_6(matrix QMatrix_ITF) *QMatrix4x4 {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QMatrix_PTR() != nil {
		convArg0 = matrix.QMatrix_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x4C2ERK7QMatrix", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMatrix4x4)
	return gothis
}

// /usr/include/qt/QtGui/qmatrix4x4.h:77
// index:0
// Public inline Visibility=Default Availability=Available
// [4] const float & operator()(int, int) const

/*

 */
func (this *QMatrix4x4) Operator_fncall(row int, column int) float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x4clEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cpretval2go("float32", rv).(float32) // 3331
}

// /usr/include/qt/QtGui/qmatrix4x4.h:78
// index:1
// Public inline Visibility=Default Availability=Available
// [4] float & operator()(int, int)

/*

 */
func (this *QMatrix4x4) Operator_fncall_1(row int, column int) float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x4clEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cpretval2go("float32", rv).(float32) // 3331
}

// /usr/include/qt/QtGui/qmatrix4x4.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QVector4D column(int) const

/*
Returns the elements of column index as a 4D vector.

See also setColumn() and row().
*/
func (this *QMatrix4x4) Column(index int) *QVector4D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x46columnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector4DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector4D)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:82
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setColumn(int, const QVector4D &)

/*
Sets the elements of column index to the components of value.

See also column() and setRow().
*/
func (this *QMatrix4x4) SetColumn(index int, value QVector4D_ITF) {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVector4D_PTR() != nil {
		convArg1 = value.QVector4D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x49setColumnEiRK9QVector4D", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:84
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QVector4D row(int) const

/*
Returns the elements of row index as a 4D vector.

See also setRow() and column().
*/
func (this *QMatrix4x4) Row(index int) *QVector4D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x43rowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector4DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector4D)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:85
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setRow(int, const QVector4D &)

/*
Sets the elements of row index to the components of value.

See also row() and setColumn().
*/
func (this *QMatrix4x4) SetRow(index int, value QVector4D_ITF) {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVector4D_PTR() != nil {
		convArg1 = value.QVector4D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x46setRowEiRK9QVector4D", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:88
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isAffine() const

/*
Returns true if this matrix is affine matrix; false otherwise.

An affine matrix is a 4x4 matrix with row 3 equal to (0, 0, 0, 1), e.g. no projective coefficients.

This function was introduced in  Qt 5.5.

See also isIdentity().
*/
func (this *QMatrix4x4) IsAffine() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x48isAffineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qmatrix4x4.h:90
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isIdentity() const

/*
Returns true if this matrix is the identity; false otherwise.

See also setToIdentity().
*/
func (this *QMatrix4x4) IsIdentity() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x410isIdentityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qmatrix4x4.h:91
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setToIdentity()

/*
Sets this matrix to the identity.

See also isIdentity().
*/
func (this *QMatrix4x4) SetToIdentity() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x413setToIdentityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:93
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void fill(float)

/*
Fills all elements of this matrx with value.
*/
func (this *QMatrix4x4) Fill(value float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x44fillEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:95
// index:0
// Public Visibility=Default Availability=Available
// [8] double determinant() const

/*
Returns the determinant of this matrix.
*/
func (this *QMatrix4x4) Determinant() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x411determinantEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qmatrix4x4.h:96
// index:0
// Public Visibility=Default Availability=Available
// [68] QMatrix4x4 inverted(bool *) const

/*
Returns the inverse of this matrix. Returns the identity if this matrix cannot be inverted; i.e. determinant() is zero. If invertible is not null, then true will be written to that location if the matrix can be inverted; false otherwise.

If the matrix is recognized as the identity or an orthonormal matrix, then this function will quickly invert the matrix using optimized routines.

See also determinant() and normalMatrix().
*/
func (this *QMatrix4x4) Inverted(invertible *bool) *QMatrix4x4 /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x48invertedEPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), invertible)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix4x4)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:96
// index:0
// Public Visibility=Default Availability=Available
// [68] QMatrix4x4 inverted(bool *) const

/*
Returns the inverse of this matrix. Returns the identity if this matrix cannot be inverted; i.e. determinant() is zero. If invertible is not null, then true will be written to that location if the matrix can be inverted; false otherwise.

If the matrix is recognized as the identity or an orthonormal matrix, then this function will quickly invert the matrix using optimized routines.

See also determinant() and normalMatrix().
*/
func (this *QMatrix4x4) Inverted__() *QMatrix4x4 /*123*/ {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var invertible unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x48invertedEPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), invertible)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix4x4)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:97
// index:0
// Public Visibility=Default Availability=Available
// [68] QMatrix4x4 transposed() const

/*
Returns this matrix, transposed about its diagonal.
*/
func (this *QMatrix4x4) Transposed() *QMatrix4x4 /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x410transposedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix4x4)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:100
// index:0
// Public inline Visibility=Default Availability=Available
// [68] QMatrix4x4 & operator+=(const QMatrix4x4 &)

/*

 */
func (this *QMatrix4x4) Operator_add_equal(other QMatrix4x4_ITF) *QMatrix4x4 {
	var convArg0 unsafe.Pointer
	if other != nil && other.QMatrix4x4_PTR() != nil {
		convArg0 = other.QMatrix4x4_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x4pLERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix4x4)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:101
// index:0
// Public inline Visibility=Default Availability=Available
// [68] QMatrix4x4 & operator-=(const QMatrix4x4 &)

/*

 */
func (this *QMatrix4x4) Operator_minus_equal(other QMatrix4x4_ITF) *QMatrix4x4 {
	var convArg0 unsafe.Pointer
	if other != nil && other.QMatrix4x4_PTR() != nil {
		convArg0 = other.QMatrix4x4_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x4mIERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix4x4)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:102
// index:0
// Public inline Visibility=Default Availability=Available
// [68] QMatrix4x4 & operator*=(const QMatrix4x4 &)

/*

 */
func (this *QMatrix4x4) Operator_mul_equal(other QMatrix4x4_ITF) *QMatrix4x4 {
	var convArg0 unsafe.Pointer
	if other != nil && other.QMatrix4x4_PTR() != nil {
		convArg0 = other.QMatrix4x4_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x4mLERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix4x4)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:103
// index:1
// Public inline Visibility=Default Availability=Available
// [68] QMatrix4x4 & operator*=(float)

/*

 */
func (this *QMatrix4x4) Operator_mul_equal_1(factor float32) *QMatrix4x4 {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x4mLEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), factor)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix4x4)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:104
// index:0
// Public Visibility=Default Availability=Available
// [68] QMatrix4x4 & operator/=(float)

/*

 */
func (this *QMatrix4x4) Operator_div_equal(divisor float32) *QMatrix4x4 {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x4dVEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), divisor)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix4x4)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:105
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QMatrix4x4 &) const

/*

 */
func (this *QMatrix4x4) Operator_equal_equal(other QMatrix4x4_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QMatrix4x4_PTR() != nil {
		convArg0 = other.QMatrix4x4_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x4eqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qmatrix4x4.h:106
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QMatrix4x4 &) const

/*

 */
func (this *QMatrix4x4) Operator_not_equal(other QMatrix4x4_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QMatrix4x4_PTR() != nil {
		convArg0 = other.QMatrix4x4_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x4neERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qmatrix4x4.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scale(const QVector3D &)

/*
Multiplies this matrix by another that scales coordinates by the components of vector.

See also translate() and rotate().
*/
func (this *QMatrix4x4) Scale(vector QVector3D_ITF) {
	var convArg0 unsafe.Pointer
	if vector != nil && vector.QVector3D_PTR() != nil {
		convArg0 = vector.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x45scaleERK9QVector3D", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:135
// index:1
// Public Visibility=Default Availability=Available
// [-2] void scale(float, float)

/*
Multiplies this matrix by another that scales coordinates by the components of vector.

See also translate() and rotate().
*/
func (this *QMatrix4x4) Scale_1(x float32, y float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x45scaleEff", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:136
// index:2
// Public Visibility=Default Availability=Available
// [-2] void scale(float, float, float)

/*
Multiplies this matrix by another that scales coordinates by the components of vector.

See also translate() and rotate().
*/
func (this *QMatrix4x4) Scale_2(x float32, y float32, z float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x45scaleEfff", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, z)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:137
// index:3
// Public Visibility=Default Availability=Available
// [-2] void scale(float)

/*
Multiplies this matrix by another that scales coordinates by the components of vector.

See also translate() and rotate().
*/
func (this *QMatrix4x4) Scale_3(factor float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x45scaleEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), factor)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void translate(const QVector3D &)

/*
Multiplies this matrix by another that translates coordinates by the components of vector.

See also scale() and rotate().
*/
func (this *QMatrix4x4) Translate(vector QVector3D_ITF) {
	var convArg0 unsafe.Pointer
	if vector != nil && vector.QVector3D_PTR() != nil {
		convArg0 = vector.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x49translateERK9QVector3D", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:138
// index:1
// Public Visibility=Default Availability=Available
// [-2] void translate(float, float)

/*
Multiplies this matrix by another that translates coordinates by the components of vector.

See also scale() and rotate().
*/
func (this *QMatrix4x4) Translate_1(x float32, y float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x49translateEff", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:139
// index:2
// Public Visibility=Default Availability=Available
// [-2] void translate(float, float, float)

/*
Multiplies this matrix by another that translates coordinates by the components of vector.

See also scale() and rotate().
*/
func (this *QMatrix4x4) Translate_2(x float32, y float32, z float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x49translateEfff", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, z)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void rotate(float, const QVector3D &)

/*
Multiples this matrix by another that rotates coordinates through angle degrees about vector.

See also scale() and translate().
*/
func (this *QMatrix4x4) Rotate(angle float32, vector QVector3D_ITF) {
	var convArg1 unsafe.Pointer
	if vector != nil && vector.QVector3D_PTR() != nil {
		convArg1 = vector.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x46rotateEfRK9QVector3D", qtrt.FFI_TYPE_POINTER, this.GetCthis(), angle, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:140
// index:1
// Public Visibility=Default Availability=Available
// [-2] void rotate(float, float, float, float)

/*
Multiples this matrix by another that rotates coordinates through angle degrees about vector.

See also scale() and translate().
*/
func (this *QMatrix4x4) Rotate_1(angle float32, x float32, y float32, z float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x46rotateEffff", qtrt.FFI_TYPE_POINTER, this.GetCthis(), angle, x, y, z)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:140
// index:1
// Public Visibility=Default Availability=Available
// [-2] void rotate(float, float, float, float)

/*
Multiples this matrix by another that rotates coordinates through angle degrees about vector.

See also scale() and translate().
*/
func (this *QMatrix4x4) Rotate_1_(angle float32, x float32, y float32) {
	// arg: 3, float=Float, =Invalid, , Invalid
	z := float32(0.0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x46rotateEffff", qtrt.FFI_TYPE_POINTER, this.GetCthis(), angle, x, y, z)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:142
// index:2
// Public Visibility=Default Availability=Available
// [-2] void rotate(const QQuaternion &)

/*
Multiples this matrix by another that rotates coordinates through angle degrees about vector.

See also scale() and translate().
*/
func (this *QMatrix4x4) Rotate_2(quaternion QQuaternion_ITF) {
	var convArg0 unsafe.Pointer
	if quaternion != nil && quaternion.QQuaternion_PTR() != nil {
		convArg0 = quaternion.QQuaternion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x46rotateERK11QQuaternion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:145
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ortho(const QRect &)

/*
Multiplies this matrix by another that applies an orthographic projection for a window with lower-left corner (left, bottom), upper-right corner (right, top), and the specified nearPlane and farPlane clipping planes.

See also frustum() and perspective().
*/
func (this *QMatrix4x4) Ortho(rect qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x45orthoERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:146
// index:1
// Public Visibility=Default Availability=Available
// [-2] void ortho(const QRectF &)

/*
Multiplies this matrix by another that applies an orthographic projection for a window with lower-left corner (left, bottom), upper-right corner (right, top), and the specified nearPlane and farPlane clipping planes.

See also frustum() and perspective().
*/
func (this *QMatrix4x4) Ortho_1(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x45orthoERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:147
// index:2
// Public Visibility=Default Availability=Available
// [-2] void ortho(float, float, float, float, float, float)

/*
Multiplies this matrix by another that applies an orthographic projection for a window with lower-left corner (left, bottom), upper-right corner (right, top), and the specified nearPlane and farPlane clipping planes.

See also frustum() and perspective().
*/
func (this *QMatrix4x4) Ortho_2(left float32, right float32, bottom float32, top float32, nearPlane float32, farPlane float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x45orthoEffffff", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, right, bottom, top, nearPlane, farPlane)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void frustum(float, float, float, float, float, float)

/*
Multiplies this matrix by another that applies a perspective frustum projection for a window with lower-left corner (left, bottom), upper-right corner (right, top), and the specified nearPlane and farPlane clipping planes.

See also ortho() and perspective().
*/
func (this *QMatrix4x4) Frustum(left float32, right float32, bottom float32, top float32, nearPlane float32, farPlane float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x47frustumEffffff", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, right, bottom, top, nearPlane, farPlane)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void perspective(float, float, float, float)

/*
Multiplies this matrix by another that applies a perspective projection. The vertical field of view will be verticalAngle degrees within a window with a given aspectRatio that determines the horizontal field of view. The projection will have the specified nearPlane and farPlane clipping planes which are the distances from the viewer to the corresponding planes.

See also ortho() and frustum().
*/
func (this *QMatrix4x4) Perspective(verticalAngle float32, aspectRatio float32, nearPlane float32, farPlane float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x411perspectiveEffff", qtrt.FFI_TYPE_POINTER, this.GetCthis(), verticalAngle, aspectRatio, nearPlane, farPlane)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:151
// index:0
// Public Visibility=Default Availability=Available
// [-2] void lookAt(const QVector3D &, const QVector3D &, const QVector3D &)

/*
Multiplies this matrix by a viewing matrix derived from an eye point. The center value indicates the center of the view that the eye is looking at. The up value indicates which direction should be considered up with respect to the eye.

Note: The up vector must not be parallel to the line of sight from eye to center.
*/
func (this *QMatrix4x4) LookAt(eye QVector3D_ITF, center QVector3D_ITF, up QVector3D_ITF) {
	var convArg0 unsafe.Pointer
	if eye != nil && eye.QVector3D_PTR() != nil {
		convArg0 = eye.QVector3D_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if center != nil && center.QVector3D_PTR() != nil {
		convArg1 = center.QVector3D_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if up != nil && up.QVector3D_PTR() != nil {
		convArg2 = up.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x46lookAtERK9QVector3DS2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void viewport(const QRectF &)

/*
Multiplies this matrix by another that performs the scale and bias transformation used by OpenGL to transform from normalized device coordinates (NDC) to viewport (window) coordinates. That is it maps points from the cube ranging over [-1, 1] in each dimension to the viewport with it's near-lower-left corner at (left, bottom, nearPlane) and with size (width, height, farPlane - nearPlane).

This matches the transform used by the fixed function OpenGL viewport transform controlled by the functions glViewport() and glDepthRange().
*/
func (this *QMatrix4x4) Viewport(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x48viewportERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:154
// index:1
// Public Visibility=Default Availability=Available
// [-2] void viewport(float, float, float, float, float, float)

/*
Multiplies this matrix by another that performs the scale and bias transformation used by OpenGL to transform from normalized device coordinates (NDC) to viewport (window) coordinates. That is it maps points from the cube ranging over [-1, 1] in each dimension to the viewport with it's near-lower-left corner at (left, bottom, nearPlane) and with size (width, height, farPlane - nearPlane).

This matches the transform used by the fixed function OpenGL viewport transform controlled by the functions glViewport() and glDepthRange().
*/
func (this *QMatrix4x4) Viewport_1(left float32, bottom float32, width float32, height float32, nearPlane float32, farPlane float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x48viewportEffffff", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, bottom, width, height, nearPlane, farPlane)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:154
// index:1
// Public Visibility=Default Availability=Available
// [-2] void viewport(float, float, float, float, float, float)

/*
Multiplies this matrix by another that performs the scale and bias transformation used by OpenGL to transform from normalized device coordinates (NDC) to viewport (window) coordinates. That is it maps points from the cube ranging over [-1, 1] in each dimension to the viewport with it's near-lower-left corner at (left, bottom, nearPlane) and with size (width, height, farPlane - nearPlane).

This matches the transform used by the fixed function OpenGL viewport transform controlled by the functions glViewport() and glDepthRange().
*/
func (this *QMatrix4x4) Viewport_1_(left float32, bottom float32, width float32, height float32) {
	// arg: 4, float=Float, =Invalid, , Invalid
	nearPlane := float32(0.0)
	// arg: 5, float=Float, =Invalid, , Invalid
	farPlane := float32(1.0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x48viewportEffffff", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, bottom, width, height, nearPlane, farPlane)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:154
// index:1
// Public Visibility=Default Availability=Available
// [-2] void viewport(float, float, float, float, float, float)

/*
Multiplies this matrix by another that performs the scale and bias transformation used by OpenGL to transform from normalized device coordinates (NDC) to viewport (window) coordinates. That is it maps points from the cube ranging over [-1, 1] in each dimension to the viewport with it's near-lower-left corner at (left, bottom, nearPlane) and with size (width, height, farPlane - nearPlane).

This matches the transform used by the fixed function OpenGL viewport transform controlled by the functions glViewport() and glDepthRange().
*/
func (this *QMatrix4x4) Viewport_1_1(left float32, bottom float32, width float32, height float32, nearPlane float32) {
	// arg: 5, float=Float, =Invalid, , Invalid
	farPlane := float32(1.0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x48viewportEffffff", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, bottom, width, height, nearPlane, farPlane)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void flipCoordinates()

/*

 */
func (this *QMatrix4x4) FlipCoordinates() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x415flipCoordinatesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void copyDataTo(float *) const

/*
Retrieves the 16 items in this matrix and copies them to values in row-major order.
*/
func (this *QMatrix4x4) CopyDataTo(values unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x410copyDataToEPf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), values)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:159
// index:0
// Public Visibility=Default Availability=Available
// [48] QMatrix toAffine() const

/*
Returns the conventional Qt 2D affine transformation matrix that corresponds to this matrix. It is assumed that this matrix only contains 2D affine transformation elements.

See also toTransform().
*/
func (this *QMatrix4x4) ToAffine() *QMatrix /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x48toAffineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:160
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform toTransform() const

/*
Returns the conventional Qt 2D transformation matrix that corresponds to this matrix.

The returned QTransform is formed by simply dropping the third row and third column of the QMatrix4x4. This is suitable for implementing orthographic projections where the z co-ordinate should be dropped rather than projected.

See also toAffine().
*/
func (this *QMatrix4x4) ToTransform() *QTransform /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x411toTransformEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:161
// index:1
// Public Visibility=Default Availability=Available
// [88] QTransform toTransform(float) const

/*
Returns the conventional Qt 2D transformation matrix that corresponds to this matrix.

The returned QTransform is formed by simply dropping the third row and third column of the QMatrix4x4. This is suitable for implementing orthographic projections where the z co-ordinate should be dropped rather than projected.

See also toAffine().
*/
func (this *QMatrix4x4) ToTransform_1(distanceToPlane float32) *QTransform /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x411toTransformEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), distanceToPlane)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:163
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint map(const QPoint &) const

/*
Maps point by multiplying this matrix by point.

See also mapRect().
*/
func (this *QMatrix4x4) Map(point qtcore.QPoint_ITF) *qtcore.QPoint /*123*/ {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPoint_PTR() != nil {
		convArg0 = point.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x43mapERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:164
// index:1
// Public Visibility=Default Availability=Available
// [16] QPointF map(const QPointF &) const

/*
Maps point by multiplying this matrix by point.

See also mapRect().
*/
func (this *QMatrix4x4) Map_1(point qtcore.QPointF_ITF) *qtcore.QPointF /*123*/ {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPointF_PTR() != nil {
		convArg0 = point.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x43mapERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:166
// index:2
// Public Visibility=Default Availability=Available
// [12] QVector3D map(const QVector3D &) const

/*
Maps point by multiplying this matrix by point.

See also mapRect().
*/
func (this *QMatrix4x4) Map_2(point QVector3D_ITF) *QVector3D /*123*/ {
	var convArg0 unsafe.Pointer
	if point != nil && point.QVector3D_PTR() != nil {
		convArg0 = point.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x43mapERK9QVector3D", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:170
// index:3
// Public Visibility=Default Availability=Available
// [16] QVector4D map(const QVector4D &) const

/*
Maps point by multiplying this matrix by point.

See also mapRect().
*/
func (this *QMatrix4x4) Map_3(point QVector4D_ITF) *QVector4D /*123*/ {
	var convArg0 unsafe.Pointer
	if point != nil && point.QVector4D_PTR() != nil {
		convArg0 = point.QVector4D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x43mapERK9QVector4D", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector4DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector4D)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:167
// index:0
// Public Visibility=Default Availability=Available
// [12] QVector3D mapVector(const QVector3D &) const

/*
Maps vector by multiplying the top 3x3 portion of this matrix by vector. The translation and projection components of this matrix are ignored.

See also map().
*/
func (this *QMatrix4x4) MapVector(vector QVector3D_ITF) *QVector3D /*123*/ {
	var convArg0 unsafe.Pointer
	if vector != nil && vector.QVector3D_PTR() != nil {
		convArg0 = vector.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x49mapVectorERK9QVector3D", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:172
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect mapRect(const QRect &) const

/*
Maps rect by multiplying this matrix by the corners of rect and then forming a new rectangle from the results. The returned rectangle will be an ordinary 2D rectangle with sides parallel to the horizontal and vertical axes.

See also map().
*/
func (this *QMatrix4x4) MapRect(rect qtcore.QRect_ITF) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x47mapRectERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:173
// index:1
// Public Visibility=Default Availability=Available
// [32] QRectF mapRect(const QRectF &) const

/*
Maps rect by multiplying this matrix by the corners of rect and then forming a new rectangle from the results. The returned rectangle will be an ordinary 2D rectangle with sides parallel to the horizontal and vertical axes.

See also map().
*/
func (this *QMatrix4x4) MapRect_1(rect qtcore.QRectF_ITF) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x47mapRectERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:178
// index:0
// Public inline Visibility=Default Availability=Available
// [8] float * data()

/*
Returns a pointer to the raw data of this matrix.

See also constData() and optimize().
*/
func (this *QMatrix4x4) Data() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x44dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qmatrix4x4.h:179
// index:1
// Public inline Visibility=Default Availability=Available
// [8] const float * data() const

/*
Returns a pointer to the raw data of this matrix.

See also constData() and optimize().
*/
func (this *QMatrix4x4) Data_1() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x44dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qmatrix4x4.h:180
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const float * constData() const

/*
Returns a constant pointer to the raw data of this matrix. This raw data is stored in column-major format.

See also data().
*/
func (this *QMatrix4x4) ConstData() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x49constDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qmatrix4x4.h:182
// index:0
// Public Visibility=Default Availability=Available
// [-2] void optimize()

/*
Optimize the usage of this matrix from its current elements.

Some operations such as translate(), scale(), and rotate() can be performed more efficiently if the matrix being modified is already known to be the identity, a previous translate(), a previous scale(), etc.

Normally the QMatrix4x4 class keeps track of this special type internally as operations are performed. However, if the matrix is modified directly with {QLoggingCategory::operator()}{operator()()} or data(), then QMatrix4x4 will lose track of the special type and will revert to the safest but least efficient operations thereafter.

By calling optimize() after directly modifying the matrix, the programmer can force QMatrix4x4 to recover the special type if the elements appear to conform to one of the known optimized types.

See also operator()(), data(), and translate().
*/
func (this *QMatrix4x4) Optimize() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x48optimizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

func DeleteQMatrix4x4(this *QMatrix4x4) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x4D2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QMatrix4x4__ = int

//
const QMatrix4x4__Identity QMatrix4x4__ = 0

//
const QMatrix4x4__Translation QMatrix4x4__ = 1

//
const QMatrix4x4__Scale QMatrix4x4__ = 2

//
const QMatrix4x4__Rotation2D QMatrix4x4__ = 4

//
const QMatrix4x4__Rotation QMatrix4x4__ = 8

//
const QMatrix4x4__Perspective QMatrix4x4__ = 16

//
const QMatrix4x4__General QMatrix4x4__ = 31

func (this *QMatrix4x4) ItemName(val int) string {
	switch val {
	case QMatrix4x4__Identity: // 0
		return "Identity"
	case QMatrix4x4__Translation: // 1
		return "Translation"
	case QMatrix4x4__Scale: // 2
		return "Scale"
	case QMatrix4x4__Rotation2D: // 4
		return "Rotation2D"
	case QMatrix4x4__Rotation: // 8
		return "Rotation"
	case QMatrix4x4__Perspective: // 16
		return "Perspective"
	case QMatrix4x4__General: // 31
		return "General"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QMatrix4x4_ItemName(val int) string {
	var nilthis *QMatrix4x4
	return nilthis.ItemName(val)
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
