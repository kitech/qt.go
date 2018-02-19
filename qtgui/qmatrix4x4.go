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
func (this *QMatrix4x4) Operator_fncall(row int, column int) float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x4clEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cpretval2go("float32", rv).(float32) // 3331
}

// /usr/include/qt/QtGui/qmatrix4x4.h:78
// index:1
// Public inline Visibility=Default Availability=Available
// [4] float & operator()(int, int)
func (this *QMatrix4x4) Operator_fncall_1(row int, column int) float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x4clEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cpretval2go("float32", rv).(float32) // 3331
}

// /usr/include/qt/QtGui/qmatrix4x4.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QVector4D column(int) const
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
func (this *QMatrix4x4) IsAffine() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x48isAffineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qmatrix4x4.h:90
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isIdentity() const
func (this *QMatrix4x4) IsIdentity() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x410isIdentityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qmatrix4x4.h:91
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setToIdentity()
func (this *QMatrix4x4) SetToIdentity() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x413setToIdentityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:93
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void fill(float)
func (this *QMatrix4x4) Fill(value float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x44fillEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:95
// index:0
// Public Visibility=Default Availability=Available
// [8] double determinant() const
func (this *QMatrix4x4) Determinant() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x411determinantEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qmatrix4x4.h:96
// index:0
// Public Visibility=Default Availability=Available
// [68] QMatrix4x4 inverted(_Bool *) const
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
// [68] QMatrix4x4 inverted(_Bool *) const
func (this *QMatrix4x4) Inverted__() *QMatrix4x4 /*123*/ {
	// arg: 0, bool *=Pointer, =Invalid,
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
func (this *QMatrix4x4) Scale_1(x float32, y float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x45scaleEff", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:136
// index:2
// Public Visibility=Default Availability=Available
// [-2] void scale(float, float, float)
func (this *QMatrix4x4) Scale_2(x float32, y float32, z float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x45scaleEfff", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, z)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:137
// index:3
// Public Visibility=Default Availability=Available
// [-2] void scale(float)
func (this *QMatrix4x4) Scale_3(factor float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x45scaleEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), factor)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void translate(const QVector3D &)
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
func (this *QMatrix4x4) Translate_1(x float32, y float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x49translateEff", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:139
// index:2
// Public Visibility=Default Availability=Available
// [-2] void translate(float, float, float)
func (this *QMatrix4x4) Translate_2(x float32, y float32, z float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x49translateEfff", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, z)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void rotate(float, const QVector3D &)
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
func (this *QMatrix4x4) Rotate_1(angle float32, x float32, y float32, z float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x46rotateEffff", qtrt.FFI_TYPE_POINTER, this.GetCthis(), angle, x, y, z)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:140
// index:1
// Public Visibility=Default Availability=Available
// [-2] void rotate(float, float, float, float)
func (this *QMatrix4x4) Rotate_1_(angle float32, x float32, y float32) {
	// arg: 3, float=Float, =Invalid,
	z := float32(0.0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x46rotateEffff", qtrt.FFI_TYPE_POINTER, this.GetCthis(), angle, x, y, z)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:142
// index:2
// Public Visibility=Default Availability=Available
// [-2] void rotate(const QQuaternion &)
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
func (this *QMatrix4x4) Ortho_2(left float32, right float32, bottom float32, top float32, nearPlane float32, farPlane float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x45orthoEffffff", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, right, bottom, top, nearPlane, farPlane)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void frustum(float, float, float, float, float, float)
func (this *QMatrix4x4) Frustum(left float32, right float32, bottom float32, top float32, nearPlane float32, farPlane float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x47frustumEffffff", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, right, bottom, top, nearPlane, farPlane)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void perspective(float, float, float, float)
func (this *QMatrix4x4) Perspective(verticalAngle float32, aspectRatio float32, nearPlane float32, farPlane float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x411perspectiveEffff", qtrt.FFI_TYPE_POINTER, this.GetCthis(), verticalAngle, aspectRatio, nearPlane, farPlane)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:151
// index:0
// Public Visibility=Default Availability=Available
// [-2] void lookAt(const QVector3D &, const QVector3D &, const QVector3D &)
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
func (this *QMatrix4x4) Viewport_1(left float32, bottom float32, width float32, height float32, nearPlane float32, farPlane float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x48viewportEffffff", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, bottom, width, height, nearPlane, farPlane)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:154
// index:1
// Public Visibility=Default Availability=Available
// [-2] void viewport(float, float, float, float, float, float)
func (this *QMatrix4x4) Viewport_1_(left float32, bottom float32, width float32, height float32) {
	// arg: 4, float=Float, =Invalid,
	nearPlane := float32(0.0)
	// arg: 5, float=Float, =Invalid,
	farPlane := float32(1.0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x48viewportEffffff", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, bottom, width, height, nearPlane, farPlane)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:154
// index:1
// Public Visibility=Default Availability=Available
// [-2] void viewport(float, float, float, float, float, float)
func (this *QMatrix4x4) Viewport_1_1(left float32, bottom float32, width float32, height float32, nearPlane float32) {
	// arg: 5, float=Float, =Invalid,
	farPlane := float32(1.0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x48viewportEffffff", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, bottom, width, height, nearPlane, farPlane)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void flipCoordinates()
func (this *QMatrix4x4) FlipCoordinates() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x415flipCoordinatesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void copyDataTo(float *) const
func (this *QMatrix4x4) CopyDataTo(values unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x410copyDataToEPf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), values)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:159
// index:0
// Public Visibility=Default Availability=Available
// [48] QMatrix toAffine() const
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
func (this *QMatrix4x4) Data() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x44dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qmatrix4x4.h:179
// index:1
// Public inline Visibility=Default Availability=Available
// [8] const float * data() const
func (this *QMatrix4x4) Data_1() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x44dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qmatrix4x4.h:180
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const float * constData() const
func (this *QMatrix4x4) ConstData() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMatrix4x49constDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qmatrix4x4.h:182
// index:0
// Public Visibility=Default Availability=Available
// [-2] void optimize()
func (this *QMatrix4x4) Optimize() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x48optimizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

func DeleteQMatrix4x4(this *QMatrix4x4) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMatrix4x4D2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QMatrix4x4__ = int

const QMatrix4x4__Identity QMatrix4x4__ = 0
const QMatrix4x4__Translation QMatrix4x4__ = 1
const QMatrix4x4__Scale QMatrix4x4__ = 2
const QMatrix4x4__Rotation2D QMatrix4x4__ = 4
const QMatrix4x4__Rotation QMatrix4x4__ = 8
const QMatrix4x4__Perspective QMatrix4x4__ = 16
const QMatrix4x4__General QMatrix4x4__ = 31

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
