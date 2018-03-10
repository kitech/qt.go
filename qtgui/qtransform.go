package qtgui

// /usr/include/qt/QtGui/qtransform.h
// #include <qtransform.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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
type QTransform struct {
	*qtrt.CObject
}
type QTransform_ITF interface {
	QTransform_PTR() *QTransform
}

func (ptr *QTransform) QTransform_PTR() *QTransform { return ptr }

func (this *QTransform) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTransform) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTransformFromPointer(cthis unsafe.Pointer) *QTransform {
	return &QTransform{&qtrt.CObject{cthis}}
}
func (*QTransform) NewFromPointer(cthis unsafe.Pointer) *QTransform {
	return NewQTransformFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtransform.h:69
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QTransform(Qt::Initialization)

/*
Constructs an identity matrix.

All elements are set to zero except m11 and m22 (specifying the scale) and m33 which are set to 1.

See also reset().
*/
func NewQTransform(arg0 int) *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransformC2EN2Qt14InitializationE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTransformFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTransform)
	return gothis
}

// /usr/include/qt/QtGui/qtransform.h:70
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTransform()

/*
Constructs an identity matrix.

All elements are set to zero except m11 and m22 (specifying the scale) and m33 which are set to 1.

See also reset().
*/
func NewQTransform_1() *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransformC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTransformFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTransform)
	return gothis
}

// /usr/include/qt/QtGui/qtransform.h:71
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QTransform(qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal)

/*
Constructs an identity matrix.

All elements are set to zero except m11 and m22 (specifying the scale) and m33 which are set to 1.

See also reset().
*/
func NewQTransform_2(h11 float64, h12 float64, h13 float64, h21 float64, h22 float64, h23 float64, h31 float64, h32 float64, h33 float64) *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransformC2Eddddddddd", qtrt.FFI_TYPE_POINTER, h11, h12, h13, h21, h22, h23, h31, h32, h33)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTransformFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTransform)
	return gothis
}

// /usr/include/qt/QtGui/qtransform.h:71
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QTransform(qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal)

/*
Constructs an identity matrix.

All elements are set to zero except m11 and m22 (specifying the scale) and m33 which are set to 1.

See also reset().
*/
func NewQTransform_2_(h11 float64, h12 float64, h13 float64, h21 float64, h22 float64, h23 float64, h31 float64, h32 float64) *QTransform {
	// arg: 8, qreal=Typedef, qreal=Typedef, double
	h33 := float64(1.0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransformC2Eddddddddd", qtrt.FFI_TYPE_POINTER, h11, h12, h13, h21, h22, h23, h31, h32, h33)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTransformFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTransform)
	return gothis
}

// /usr/include/qt/QtGui/qtransform.h:74
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QTransform(qreal, qreal, qreal, qreal, qreal, qreal)

/*
Constructs an identity matrix.

All elements are set to zero except m11 and m22 (specifying the scale) and m33 which are set to 1.

See also reset().
*/
func NewQTransform_3(h11 float64, h12 float64, h21 float64, h22 float64, dx float64, dy float64) *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransformC2Edddddd", qtrt.FFI_TYPE_POINTER, h11, h12, h21, h22, dx, dy)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTransformFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTransform)
	return gothis
}

// /usr/include/qt/QtGui/qtransform.h:76
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QTransform(const QMatrix &)

/*
Constructs an identity matrix.

All elements are set to zero except m11 and m22 (specifying the scale) and m33 which are set to 1.

See also reset().
*/
func NewQTransform_4(mtx QMatrix_ITF) *QTransform {
	var convArg0 unsafe.Pointer
	if mtx != nil && mtx.QMatrix_PTR() != nil {
		convArg0 = mtx.QMatrix_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransformC2ERK7QMatrix", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTransformFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTransform)
	return gothis
}

// /usr/include/qt/QtGui/qtransform.h:80
// index:0
// Public inline Visibility=Default Availability=Available
// [88] QTransform & operator=(QTransform &&)

/*

 */
func (this *QTransform) Operator_equal(other unsafe.Pointer /*333*/) *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransformaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:82
// index:1
// Public Visibility=Default Availability=Available
// [88] QTransform & operator=(const QTransform &)

/*

 */
func (this *QTransform) Operator_equal_1(arg0 QTransform_ITF) *QTransform {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTransform_PTR() != nil {
		convArg0 = arg0.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransformaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:91
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAffine() const

/*
Returns true if the matrix represent an affine transformation, otherwise returns false.
*/
func (this *QTransform) IsAffine() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform8isAffineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtransform.h:92
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isIdentity() const

/*
Returns true if the matrix is the identity matrix, otherwise returns false.

See also reset().
*/
func (this *QTransform) IsIdentity() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform10isIdentityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtransform.h:93
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isInvertible() const

/*
Returns true if the matrix is invertible, otherwise returns false.

See also inverted().
*/
func (this *QTransform) IsInvertible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform12isInvertibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtransform.h:94
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isScaling() const

/*
Returns true if the matrix represents a scaling transformation, otherwise returns false.

See also reset().
*/
func (this *QTransform) IsScaling() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform9isScalingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtransform.h:95
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRotating() const

/*
Returns true if the matrix represents some kind of a rotating transformation, otherwise returns false.

Note: A rotation transformation of 180 degrees and/or 360 degrees is treated as a scaling transformation.

See also reset().
*/
func (this *QTransform) IsRotating() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform10isRotatingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtransform.h:96
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isTranslating() const

/*
Returns true if the matrix represents a translating transformation, otherwise returns false.

See also reset().
*/
func (this *QTransform) IsTranslating() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform13isTranslatingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtransform.h:98
// index:0
// Public Visibility=Default Availability=Available
// [4] QTransform::TransformationType type() const

/*
Returns the transformation type of this matrix.

The transformation type is the highest enumeration value capturing all of the matrix's transformations. For example, if the matrix both scales and shears, the type would be TxShear, because TxShear has a higher enumeration value than TxScale.

Knowing the transformation type of a matrix is useful for optimization: you can often handle specific types more optimally than handling the generic case.
*/
func (this *QTransform) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtransform.h:100
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal determinant() const

/*
Returns the matrix's determinant.
*/
func (this *QTransform) Determinant() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform11determinantEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtransform.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal det() const

/*

 */
func (this *QTransform) Det() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3detEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtransform.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal m11() const

/*
Returns the horizontal scaling factor.

See also scale() and Basic Matrix Operations.
*/
func (this *QTransform) M11() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3m11Ev", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtransform.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal m12() const

/*
Returns the vertical shearing factor.

See also shear() and Basic Matrix Operations.
*/
func (this *QTransform) M12() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3m12Ev", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtransform.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal m13() const

/*
Returns the horizontal projection factor.

See also translate() and Basic Matrix Operations.
*/
func (this *QTransform) M13() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3m13Ev", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtransform.h:106
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal m21() const

/*
Returns the horizontal shearing factor.

See also shear() and Basic Matrix Operations.
*/
func (this *QTransform) M21() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3m21Ev", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtransform.h:107
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal m22() const

/*
Returns the vertical scaling factor.

See also scale() and Basic Matrix Operations.
*/
func (this *QTransform) M22() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3m22Ev", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtransform.h:108
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal m23() const

/*
Returns the vertical projection factor.

See also translate() and Basic Matrix Operations.
*/
func (this *QTransform) M23() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3m23Ev", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtransform.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal m31() const

/*
Returns the horizontal translation factor.

See also dx(), translate(), and Basic Matrix Operations.
*/
func (this *QTransform) M31() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3m31Ev", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtransform.h:110
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal m32() const

/*
Returns the vertical translation factor.

See also dy(), translate(), and Basic Matrix Operations.
*/
func (this *QTransform) M32() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3m32Ev", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtransform.h:111
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal m33() const

/*
Returns the division factor.

See also translate() and Basic Matrix Operations.
*/
func (this *QTransform) M33() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3m33Ev", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtransform.h:112
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal dx() const

/*
Returns the horizontal translation factor.

See also m31(), translate(), and Basic Matrix Operations.
*/
func (this *QTransform) Dx() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform2dxEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtransform.h:113
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal dy() const

/*
Returns the vertical translation factor.

See also translate() and Basic Matrix Operations.
*/
func (this *QTransform) Dy() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform2dyEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtransform.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMatrix(qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal)

/*
Sets the matrix elements to the specified values, m11, m12, m13 m21, m22, m23 m31, m32 and m33. Note that this function replaces the previous values. QTransform provides the translate(), rotate(), scale() and shear() convenience functions to manipulate the various matrix elements based on the currently defined coordinate system.

See also QTransform().
*/
func (this *QTransform) SetMatrix(m11 float64, m12 float64, m13 float64, m21 float64, m22 float64, m23 float64, m31 float64, m32 float64, m33 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransform9setMatrixEddddddddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), m11, m12, m13, m21, m22, m23, m31, m32, m33)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:119
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform inverted(_Bool *) const

/*
Returns an inverted copy of this matrix.

If the matrix is singular (not invertible), the returned matrix is the identity matrix. If invertible is valid (i.e. not 0), its value is set to true if the matrix is invertible, otherwise it is set to false.

See also isInvertible().
*/
func (this *QTransform) Inverted(invertible *bool) *QTransform /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform8invertedEPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), invertible)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:119
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform inverted(_Bool *) const

/*
Returns an inverted copy of this matrix.

If the matrix is singular (not invertible), the returned matrix is the identity matrix. If invertible is valid (i.e. not 0), its value is set to true if the matrix is invertible, otherwise it is set to false.

See also isInvertible().
*/
func (this *QTransform) Inverted__() *QTransform /*123*/ {
	// arg: 0, bool *=Pointer, =Invalid,
	var invertible unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform8invertedEPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), invertible)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:120
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform adjoint() const

/*
Returns the adjoint of this matrix.
*/
func (this *QTransform) Adjoint() *QTransform /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform7adjointEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:121
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform transposed() const

/*
Returns the transpose of this matrix.
*/
func (this *QTransform) Transposed() *QTransform /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform10transposedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:123
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform & translate(qreal, qreal)

/*
Moves the coordinate system dx along the x axis and dy along the y axis, and returns a reference to the matrix.

See also setMatrix().
*/
func (this *QTransform) Translate(dx float64, dy float64) *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransform9translateEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:124
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform & scale(qreal, qreal)

/*
Scales the coordinate system by sx horizontally and sy vertically, and returns a reference to the matrix.

See also setMatrix().
*/
func (this *QTransform) Scale(sx float64, sy float64) *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransform5scaleEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sx, sy)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:125
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform & shear(qreal, qreal)

/*
Shears the coordinate system by sh horizontally and sv vertically, and returns a reference to the matrix.

See also setMatrix().
*/
func (this *QTransform) Shear(sh float64, sv float64) *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransform5shearEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sh, sv)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:126
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform & rotate(qreal, Qt::Axis)

/*
Rotates the coordinate system counterclockwise by the given angle about the specified axis and returns a reference to the matrix.

Note that if you apply a QTransform to a point defined in widget coordinates, the direction of the rotation will be clockwise because the y-axis points downwards.

The angle is specified in degrees.

See also setMatrix().
*/
func (this *QTransform) Rotate(a float64, axis int) *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransform6rotateEdN2Qt4AxisE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), a, axis)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:126
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform & rotate(qreal, Qt::Axis)

/*
Rotates the coordinate system counterclockwise by the given angle about the specified axis and returns a reference to the matrix.

Note that if you apply a QTransform to a point defined in widget coordinates, the direction of the rotation will be clockwise because the y-axis points downwards.

The angle is specified in degrees.

See also setMatrix().
*/
func (this *QTransform) Rotate__(a float64) *QTransform {
	// arg: 1, Qt::Axis=Elaborated, Qt::Axis=Enum,
	axis := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransform6rotateEdN2Qt4AxisE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), a, axis)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:127
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform & rotateRadians(qreal, Qt::Axis)

/*
Rotates the coordinate system counterclockwise by the given angle about the specified axis and returns a reference to the matrix.

Note that if you apply a QTransform to a point defined in widget coordinates, the direction of the rotation will be clockwise because the y-axis points downwards.

The angle is specified in radians.

See also setMatrix().
*/
func (this *QTransform) RotateRadians(a float64, axis int) *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransform13rotateRadiansEdN2Qt4AxisE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), a, axis)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:127
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform & rotateRadians(qreal, Qt::Axis)

/*
Rotates the coordinate system counterclockwise by the given angle about the specified axis and returns a reference to the matrix.

Note that if you apply a QTransform to a point defined in widget coordinates, the direction of the rotation will be clockwise because the y-axis points downwards.

The angle is specified in radians.

See also setMatrix().
*/
func (this *QTransform) RotateRadians__(a float64) *QTransform {
	// arg: 1, Qt::Axis=Elaborated, Qt::Axis=Enum,
	axis := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransform13rotateRadiansEdN2Qt4AxisE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), a, axis)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:129
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool squareToQuad(const QPolygonF &, QTransform &)

/*
Creates a transformation matrix, trans, that maps a unit square to a four-sided polygon, quad. Returns true if the transformation is constructed or false if such a transformation does not exist.

See also quadToSquare() and quadToQuad().
*/
func (this *QTransform) SquareToQuad(square QPolygonF_ITF, result QTransform_ITF) bool {
	var convArg0 unsafe.Pointer
	if square != nil && square.QPolygonF_PTR() != nil {
		convArg0 = square.QPolygonF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if result != nil && result.QTransform_PTR() != nil {
		convArg1 = result.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransform12squareToQuadERK9QPolygonFRS_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QTransform_SquareToQuad(square QPolygonF_ITF, result QTransform_ITF) bool {
	var nilthis *QTransform
	rv := nilthis.SquareToQuad(square, result)
	return rv
}

// /usr/include/qt/QtGui/qtransform.h:130
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool quadToSquare(const QPolygonF &, QTransform &)

/*
Creates a transformation matrix, trans, that maps a four-sided polygon, quad, to a unit square. Returns true if the transformation is constructed or false if such a transformation does not exist.

See also squareToQuad() and quadToQuad().
*/
func (this *QTransform) QuadToSquare(quad QPolygonF_ITF, result QTransform_ITF) bool {
	var convArg0 unsafe.Pointer
	if quad != nil && quad.QPolygonF_PTR() != nil {
		convArg0 = quad.QPolygonF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if result != nil && result.QTransform_PTR() != nil {
		convArg1 = result.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransform12quadToSquareERK9QPolygonFRS_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QTransform_QuadToSquare(quad QPolygonF_ITF, result QTransform_ITF) bool {
	var nilthis *QTransform
	rv := nilthis.QuadToSquare(quad, result)
	return rv
}

// /usr/include/qt/QtGui/qtransform.h:131
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool quadToQuad(const QPolygonF &, const QPolygonF &, QTransform &)

/*
Creates a transformation matrix, trans, that maps a four-sided polygon, one, to another four-sided polygon, two. Returns true if the transformation is possible; otherwise returns false.

This is a convenience method combining quadToSquare() and squareToQuad() methods. It allows the input quad to be transformed into any other quad.

See also squareToQuad() and quadToSquare().
*/
func (this *QTransform) QuadToQuad(one QPolygonF_ITF, two QPolygonF_ITF, result QTransform_ITF) bool {
	var convArg0 unsafe.Pointer
	if one != nil && one.QPolygonF_PTR() != nil {
		convArg0 = one.QPolygonF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if two != nil && two.QPolygonF_PTR() != nil {
		convArg1 = two.QPolygonF_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if result != nil && result.QTransform_PTR() != nil {
		convArg2 = result.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransform10quadToQuadERK9QPolygonFS2_RS_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QTransform_QuadToQuad(one QPolygonF_ITF, two QPolygonF_ITF, result QTransform_ITF) bool {
	var nilthis *QTransform
	rv := nilthis.QuadToQuad(one, two, result)
	return rv
}

// /usr/include/qt/QtGui/qtransform.h:135
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QTransform &) const

/*

 */
func (this *QTransform) Operator_equal_equal(arg0 QTransform_ITF) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTransform_PTR() != nil {
		convArg0 = arg0.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransformeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtransform.h:136
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QTransform &) const

/*

 */
func (this *QTransform) Operator_not_equal(arg0 QTransform_ITF) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTransform_PTR() != nil {
		convArg0 = arg0.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransformneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtransform.h:138
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform & operator*=(const QTransform &)

/*

 */
func (this *QTransform) Operator_mul_equal(arg0 QTransform_ITF) *QTransform {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTransform_PTR() != nil {
		convArg0 = arg0.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransformmLERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:160
// index:1
// Public Visibility=Default Availability=Available
// [88] QTransform & operator*=(qreal)

/*

 */
func (this *QTransform) Operator_mul_equal_1(div float64) *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransformmLEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), div)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:139
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform operator*(const QTransform &) const

/*

 */
func (this *QTransform) Operator_mul(o QTransform_ITF) *QTransform /*123*/ {
	var convArg0 unsafe.Pointer
	if o != nil && o.QTransform_PTR() != nil {
		convArg0 = o.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransformmlERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reset()

/*
Resets the matrix to an identity matrix, i.e. all elements are set to zero, except m11 and m22 (specifying the scale) and m33 which are set to 1.

See also QTransform(), isIdentity(), and Basic Matrix Operations.
*/
func (this *QTransform) Reset() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransform5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:144
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint map(const QPoint &) const

/*
Maps the given coordinates x and y into the coordinate system defined by this matrix. The resulting values are put in *tx and *ty, respectively.

The coordinates are transformed using the following formulas:


  x' = m11*x + m21*y + dx
  y' = m22*y + m12*x + dy
  if (is not affine) {
      w' = m13*x + m23*y + m33
      x' /= w'
      y' /= w'
  }



The point (x, y) is the original point, and (x', y') is the transformed point.

See also Basic Matrix Operations.
*/
func (this *QTransform) Map(p qtcore.QPoint_ITF) *qtcore.QPoint /*123*/ {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3mapERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:145
// index:1
// Public Visibility=Default Availability=Available
// [16] QPointF map(const QPointF &) const

/*
Maps the given coordinates x and y into the coordinate system defined by this matrix. The resulting values are put in *tx and *ty, respectively.

The coordinates are transformed using the following formulas:


  x' = m11*x + m21*y + dx
  y' = m22*y + m12*x + dy
  if (is not affine) {
      w' = m13*x + m23*y + m33
      x' /= w'
      y' /= w'
  }



The point (x, y) is the original point, and (x', y') is the transformed point.

See also Basic Matrix Operations.
*/
func (this *QTransform) Map_1(p qtcore.QPointF_ITF) *qtcore.QPointF /*123*/ {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3mapERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:146
// index:2
// Public Visibility=Default Availability=Available
// [16] QLine map(const QLine &) const

/*
Maps the given coordinates x and y into the coordinate system defined by this matrix. The resulting values are put in *tx and *ty, respectively.

The coordinates are transformed using the following formulas:


  x' = m11*x + m21*y + dx
  y' = m22*y + m12*x + dy
  if (is not affine) {
      w' = m13*x + m23*y + m33
      x' /= w'
      y' /= w'
  }



The point (x, y) is the original point, and (x', y') is the transformed point.

See also Basic Matrix Operations.
*/
func (this *QTransform) Map_2(l qtcore.QLine_ITF) *qtcore.QLine /*123*/ {
	var convArg0 unsafe.Pointer
	if l != nil && l.QLine_PTR() != nil {
		convArg0 = l.QLine_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3mapERK5QLine", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQLineFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQLine)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:147
// index:3
// Public Visibility=Default Availability=Available
// [32] QLineF map(const QLineF &) const

/*
Maps the given coordinates x and y into the coordinate system defined by this matrix. The resulting values are put in *tx and *ty, respectively.

The coordinates are transformed using the following formulas:


  x' = m11*x + m21*y + dx
  y' = m22*y + m12*x + dy
  if (is not affine) {
      w' = m13*x + m23*y + m33
      x' /= w'
      y' /= w'
  }



The point (x, y) is the original point, and (x', y') is the transformed point.

See also Basic Matrix Operations.
*/
func (this *QTransform) Map_3(l qtcore.QLineF_ITF) *qtcore.QLineF /*123*/ {
	var convArg0 unsafe.Pointer
	if l != nil && l.QLineF_PTR() != nil {
		convArg0 = l.QLineF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3mapERK6QLineF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQLineFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQLineF)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:148
// index:4
// Public Visibility=Default Availability=Available
// [8] QPolygonF map(const QPolygonF &) const

/*
Maps the given coordinates x and y into the coordinate system defined by this matrix. The resulting values are put in *tx and *ty, respectively.

The coordinates are transformed using the following formulas:


  x' = m11*x + m21*y + dx
  y' = m22*y + m12*x + dy
  if (is not affine) {
      w' = m13*x + m23*y + m33
      x' /= w'
      y' /= w'
  }



The point (x, y) is the original point, and (x', y') is the transformed point.

See also Basic Matrix Operations.
*/
func (this *QTransform) Map_4(a QPolygonF_ITF) *QPolygonF /*123*/ {
	var convArg0 unsafe.Pointer
	if a != nil && a.QPolygonF_PTR() != nil {
		convArg0 = a.QPolygonF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3mapERK9QPolygonF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:149
// index:5
// Public Visibility=Default Availability=Available
// [8] QPolygon map(const QPolygon &) const

/*
Maps the given coordinates x and y into the coordinate system defined by this matrix. The resulting values are put in *tx and *ty, respectively.

The coordinates are transformed using the following formulas:


  x' = m11*x + m21*y + dx
  y' = m22*y + m12*x + dy
  if (is not affine) {
      w' = m13*x + m23*y + m33
      x' /= w'
      y' /= w'
  }



The point (x, y) is the original point, and (x', y') is the transformed point.

See also Basic Matrix Operations.
*/
func (this *QTransform) Map_5(a QPolygon_ITF) *QPolygon /*123*/ {
	var convArg0 unsafe.Pointer
	if a != nil && a.QPolygon_PTR() != nil {
		convArg0 = a.QPolygon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3mapERK8QPolygon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygon)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:150
// index:6
// Public Visibility=Default Availability=Available
// [8] QRegion map(const QRegion &) const

/*
Maps the given coordinates x and y into the coordinate system defined by this matrix. The resulting values are put in *tx and *ty, respectively.

The coordinates are transformed using the following formulas:


  x' = m11*x + m21*y + dx
  y' = m22*y + m12*x + dy
  if (is not affine) {
      w' = m13*x + m23*y + m33
      x' /= w'
      y' /= w'
  }



The point (x, y) is the original point, and (x', y') is the transformed point.

See also Basic Matrix Operations.
*/
func (this *QTransform) Map_6(r QRegion_ITF) *QRegion /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRegion_PTR() != nil {
		convArg0 = r.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3mapERK7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:151
// index:7
// Public Visibility=Default Availability=Available
// [8] QPainterPath map(const QPainterPath &) const

/*
Maps the given coordinates x and y into the coordinate system defined by this matrix. The resulting values are put in *tx and *ty, respectively.

The coordinates are transformed using the following formulas:


  x' = m11*x + m21*y + dx
  y' = m22*y + m12*x + dy
  if (is not affine) {
      w' = m13*x + m23*y + m33
      x' /= w'
      y' /= w'
  }



The point (x, y) is the original point, and (x', y') is the transformed point.

See also Basic Matrix Operations.
*/
func (this *QTransform) Map_7(p QPainterPath_ITF) *QPainterPath /*123*/ {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPainterPath_PTR() != nil {
		convArg0 = p.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3mapERK12QPainterPath", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:155
// index:8
// Public Visibility=Default Availability=Available
// [-2] void map(int, int, int *, int *) const

/*
Maps the given coordinates x and y into the coordinate system defined by this matrix. The resulting values are put in *tx and *ty, respectively.

The coordinates are transformed using the following formulas:


  x' = m11*x + m21*y + dx
  y' = m22*y + m12*x + dy
  if (is not affine) {
      w' = m13*x + m23*y + m33
      x' /= w'
      y' /= w'
  }



The point (x, y) is the original point, and (x', y') is the transformed point.

See also Basic Matrix Operations.
*/
func (this *QTransform) Map_8(x int, y int, tx unsafe.Pointer /*666*/, ty unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3mapEiiPiS0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, tx, ty)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:156
// index:9
// Public Visibility=Default Availability=Available
// [-2] void map(qreal, qreal, qreal *, qreal *) const

/*
Maps the given coordinates x and y into the coordinate system defined by this matrix. The resulting values are put in *tx and *ty, respectively.

The coordinates are transformed using the following formulas:


  x' = m11*x + m21*y + dx
  y' = m22*y + m12*x + dy
  if (is not affine) {
      w' = m13*x + m23*y + m33
      x' /= w'
      y' /= w'
  }



The point (x, y) is the original point, and (x', y') is the transformed point.

See also Basic Matrix Operations.
*/
func (this *QTransform) Map_9(x float64, y float64, tx unsafe.Pointer /*666*/, ty unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3mapEddPdS0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, tx, ty)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:152
// index:0
// Public Visibility=Default Availability=Available
// [8] QPolygon mapToPolygon(const QRect &) const

/*
Creates and returns a QPolygon representation of the given rectangle, mapped into the coordinate system defined by this matrix.

The rectangle's coordinates are transformed using the following formulas:


  x' = m11*x + m21*y + dx
  y' = m22*y + m12*x + dy
  if (is not affine) {
      w' = m13*x + m23*y + m33
      x' /= w'
      y' /= w'
  }



Polygons and rectangles behave slightly differently when transformed (due to integer rounding), so matrix.map(QPolygon(rectangle)) is not always the same as matrix.mapToPolygon(rectangle).

See also mapRect() and Basic Matrix Operations.
*/
func (this *QTransform) MapToPolygon(r qtcore.QRect_ITF) *QPolygon /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform12mapToPolygonERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygon)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:153
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect mapRect(const QRect &) const

/*
Creates and returns a QRectF object that is a copy of the given rectangle, mapped into the coordinate system defined by this matrix.

The rectangle's coordinates are transformed using the following formulas:


  x' = m11*x + m21*y + dx
  y' = m22*y + m12*x + dy
  if (is not affine) {
      w' = m13*x + m23*y + m33
      x' /= w'
      y' /= w'
  }



If rotation or shearing has been specified, this function returns the bounding rectangle. To retrieve the exact region the given rectangle maps to, use the mapToPolygon() function instead.

See also mapToPolygon() and Basic Matrix Operations.
*/
func (this *QTransform) MapRect(arg0 qtcore.QRect_ITF) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRect_PTR() != nil {
		convArg0 = arg0.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform7mapRectERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:154
// index:1
// Public Visibility=Default Availability=Available
// [32] QRectF mapRect(const QRectF &) const

/*
Creates and returns a QRectF object that is a copy of the given rectangle, mapped into the coordinate system defined by this matrix.

The rectangle's coordinates are transformed using the following formulas:


  x' = m11*x + m21*y + dx
  y' = m22*y + m12*x + dy
  if (is not affine) {
      w' = m13*x + m23*y + m33
      x' /= w'
      y' /= w'
  }



If rotation or shearing has been specified, this function returns the bounding rectangle. To retrieve the exact region the given rectangle maps to, use the mapToPolygon() function instead.

See also mapToPolygon() and Basic Matrix Operations.
*/
func (this *QTransform) MapRect_1(arg0 qtcore.QRectF_ITF) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRectF_PTR() != nil {
		convArg0 = arg0.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform7mapRectERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:158
// index:0
// Public Visibility=Default Availability=Available
// [48] const QMatrix & toAffine() const

/*
Returns the QTransform as an affine matrix.

Warning: If a perspective transformation has been specified, then the conversion will cause loss of data.
*/
func (this *QTransform) ToAffine() *QMatrix {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform8toAffineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:161
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform & operator/=(qreal)

/*

 */
func (this *QTransform) Operator_div_equal(div float64) *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransformdVEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), div)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:162
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform & operator+=(qreal)

/*

 */
func (this *QTransform) Operator_add_equal(div float64) *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransformpLEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), div)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:163
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform & operator-=(qreal)

/*

 */
func (this *QTransform) Operator_minus_equal(div float64) *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransformmIEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), div)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:165
// index:0
// Public static Visibility=Default Availability=Available
// [88] QTransform fromTranslate(qreal, qreal)

/*
Creates a matrix which corresponds to a translation of dx along the x axis and dy along the y axis. This is the same as QTransform().translate(dx, dy) but slightly faster.

This function was introduced in  Qt 4.5.
*/
func (this *QTransform) FromTranslate(dx float64, dy float64) *QTransform /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransform13fromTranslateEdd", qtrt.FFI_TYPE_POINTER, dx, dy)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}
func QTransform_FromTranslate(dx float64, dy float64) *QTransform /*123*/ {
	var nilthis *QTransform
	rv := nilthis.FromTranslate(dx, dy)
	return rv
}

// /usr/include/qt/QtGui/qtransform.h:166
// index:0
// Public static Visibility=Default Availability=Available
// [88] QTransform fromScale(qreal, qreal)

/*
Creates a matrix which corresponds to a scaling of sx horizontally and sy vertically. This is the same as QTransform().scale(sx, sy) but slightly faster.

This function was introduced in  Qt 4.5.
*/
func (this *QTransform) FromScale(dx float64, dy float64) *QTransform /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransform9fromScaleEdd", qtrt.FFI_TYPE_POINTER, dx, dy)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}
func QTransform_FromScale(dx float64, dy float64) *QTransform /*123*/ {
	var nilthis *QTransform
	rv := nilthis.FromScale(dx, dy)
	return rv
}

func DeleteQTransform(this *QTransform) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransformD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
ConstantValue

*/
type QTransform__TransformationType = int

//
const QTransform__TxNone QTransform__TransformationType = 0

//
const QTransform__TxTranslate QTransform__TransformationType = 1

//
const QTransform__TxScale QTransform__TransformationType = 2

//
const QTransform__TxRotate QTransform__TransformationType = 4

//
const QTransform__TxShear QTransform__TransformationType = 8

//
const QTransform__TxProject QTransform__TransformationType = 16

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
