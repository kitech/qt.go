package qtgui

// /usr/include/qt/QtGui/qquaternion.h
// #include <qquaternion.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 37
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
type QQuaternion struct {
	*qtrt.CObject
}
type QQuaternion_ITF interface {
	QQuaternion_PTR() *QQuaternion
}

func (ptr *QQuaternion) QQuaternion_PTR() *QQuaternion { return ptr }

func (this *QQuaternion) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQuaternion) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQQuaternionFromPointer(cthis unsafe.Pointer) *QQuaternion {
	return &QQuaternion{&qtrt.CObject{cthis}}
}
func (*QQuaternion) NewFromPointer(cthis unsafe.Pointer) *QQuaternion {
	return NewQQuaternionFromPointer(cthis)
}

// /usr/include/qt/QtGui/qquaternion.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuaternion()

/*
Constructs an identity quaternion (1, 0, 0, 0), i.e. with the vector (0, 0, 0) and scalar 1.
*/
func NewQQuaternion() *QQuaternion {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternionC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQuaternion)
	return gothis
}

// /usr/include/qt/QtGui/qquaternion.h:60
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QQuaternion(Qt::Initialization)

/*
Constructs an identity quaternion (1, 0, 0, 0), i.e. with the vector (0, 0, 0) and scalar 1.
*/
func NewQQuaternion_1(arg0 int) *QQuaternion {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternionC2EN2Qt14InitializationE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQuaternion)
	return gothis
}

// /usr/include/qt/QtGui/qquaternion.h:61
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QQuaternion(float, float, float, float)

/*
Constructs an identity quaternion (1, 0, 0, 0), i.e. with the vector (0, 0, 0) and scalar 1.
*/
func NewQQuaternion_2(scalar float32, xpos float32, ypos float32, zpos float32) *QQuaternion {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternionC2Effff", qtrt.FFI_TYPE_POINTER, scalar, xpos, ypos, zpos)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQuaternion)
	return gothis
}

// /usr/include/qt/QtGui/qquaternion.h:63
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QQuaternion(float, const QVector3D &)

/*
Constructs an identity quaternion (1, 0, 0, 0), i.e. with the vector (0, 0, 0) and scalar 1.
*/
func NewQQuaternion_3(scalar float32, vector QVector3D_ITF) *QQuaternion {
	var convArg1 unsafe.Pointer
	if vector != nil && vector.QVector3D_PTR() != nil {
		convArg1 = vector.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternionC2EfRK9QVector3D", qtrt.FFI_TYPE_POINTER, scalar, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQuaternion)
	return gothis
}

// /usr/include/qt/QtGui/qquaternion.h:66
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QQuaternion(const QVector4D &)

/*
Constructs an identity quaternion (1, 0, 0, 0), i.e. with the vector (0, 0, 0) and scalar 1.
*/
func NewQQuaternion_4(vector QVector4D_ITF) *QQuaternion {
	var convArg0 unsafe.Pointer
	if vector != nil && vector.QVector4D_PTR() != nil {
		convArg0 = vector.QVector4D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternionC2ERK9QVector4D", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQuaternion)
	return gothis
}

// /usr/include/qt/QtGui/qquaternion.h:69
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if the x, y, z, and scalar components of this quaternion are set to 0.0; otherwise returns false.
*/
func (this *QQuaternion) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qquaternion.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isIdentity() const

/*
Returns true if the x, y, and z components of this quaternion are set to 0.0, and the scalar component is set to 1.0; otherwise returns false.
*/
func (this *QQuaternion) IsIdentity() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion10isIdentityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qquaternion.h:73
// index:0
// Public Visibility=Default Availability=Available
// [12] QVector3D vector() const

/*
Returns the vector component of this quaternion.

See also setVector() and scalar().
*/
func (this *QQuaternion) Vector() *QVector3D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion6vectorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}

// /usr/include/qt/QtGui/qquaternion.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVector(const QVector3D &)

/*
Sets the vector component of this quaternion to vector.

See also vector() and setScalar().
*/
func (this *QQuaternion) SetVector(vector QVector3D_ITF) {
	var convArg0 unsafe.Pointer
	if vector != nil && vector.QVector3D_PTR() != nil {
		convArg0 = vector.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion9setVectorERK9QVector3D", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:76
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setVector(float, float, float)

/*
Sets the vector component of this quaternion to vector.

See also vector() and setScalar().
*/
func (this *QQuaternion) SetVector_1(x float32, y float32, z float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion9setVectorEfff", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, z)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:78
// index:0
// Public Visibility=Default Availability=Available
// [4] float x() const

/*
Returns the x coordinate of this quaternion's vector.

See also setX(), y(), z(), and scalar().
*/
func (this *QQuaternion) X() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion1xEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qquaternion.h:79
// index:0
// Public Visibility=Default Availability=Available
// [4] float y() const

/*
Returns the y coordinate of this quaternion's vector.

See also setY(), x(), z(), and scalar().
*/
func (this *QQuaternion) Y() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion1yEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qquaternion.h:80
// index:0
// Public Visibility=Default Availability=Available
// [4] float z() const

/*
Returns the z coordinate of this quaternion's vector.

See also setZ(), x(), y(), and scalar().
*/
func (this *QQuaternion) Z() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion1zEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qquaternion.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] float scalar() const

/*
Returns the scalar component of this quaternion.

See also setScalar(), x(), y(), and z().
*/
func (this *QQuaternion) Scalar() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion6scalarEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qquaternion.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setX(float)

/*
Sets the x coordinate of this quaternion's vector to the given x coordinate.

See also x(), setY(), setZ(), and setScalar().
*/
func (this *QQuaternion) SetX(x float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion4setXEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setY(float)

/*
Sets the y coordinate of this quaternion's vector to the given y coordinate.

See also y(), setX(), setZ(), and setScalar().
*/
func (this *QQuaternion) SetY(y float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion4setYEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setZ(float)

/*
Sets the z coordinate of this quaternion's vector to the given z coordinate.

See also z(), setX(), setY(), and setScalar().
*/
func (this *QQuaternion) SetZ(z float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion4setZEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), z)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScalar(float)

/*
Sets the scalar component of this quaternion to scalar.

See also scalar(), setX(), setY(), and setZ().
*/
func (this *QQuaternion) SetScalar(scalar float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion9setScalarEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), scalar)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:88
// index:0
// Public static inline Visibility=Default Availability=Available
// [4] float dotProduct(const QQuaternion &, const QQuaternion &)

/*
Returns the dot product of q1 and q2.

This function was introduced in  Qt 5.5.

See also length().
*/
func (this *QQuaternion) DotProduct(q1 QQuaternion_ITF, q2 QQuaternion_ITF) float32 {
	var convArg0 unsafe.Pointer
	if q1 != nil && q1.QQuaternion_PTR() != nil {
		convArg0 = q1.QQuaternion_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if q2 != nil && q2.QQuaternion_PTR() != nil {
		convArg1 = q2.QQuaternion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion10dotProductERKS_S1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}
func QQuaternion_DotProduct(q1 QQuaternion_ITF, q2 QQuaternion_ITF) float32 {
	var nilthis *QQuaternion
	rv := nilthis.DotProduct(q1, q2)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:90
// index:0
// Public Visibility=Default Availability=Available
// [4] float length() const

/*
Returns the length of the quaternion. This is also called the "norm".

See also lengthSquared(), normalized(), and dotProduct().
*/
func (this *QQuaternion) Length() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion6lengthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qquaternion.h:91
// index:0
// Public Visibility=Default Availability=Available
// [4] float lengthSquared() const

/*
Returns the squared length of the quaternion.

See also length() and dotProduct().
*/
func (this *QQuaternion) LengthSquared() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion13lengthSquaredEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qquaternion.h:93
// index:0
// Public Visibility=Default Availability=Available
// [16] QQuaternion normalized() const

/*
Returns the normalized unit form of this quaternion.

If this quaternion is null, then a null quaternion is returned. If the length of the quaternion is very close to 1, then the quaternion will be returned as-is. Otherwise the normalized form of the quaternion of length 1 will be returned.

See also normalize(), length(), and dotProduct().
*/
func (this *QQuaternion) Normalized() *QQuaternion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion10normalizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}

// /usr/include/qt/QtGui/qquaternion.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void normalize()

/*
Normalizes the current quaternion in place. Nothing happens if this is a null quaternion or the length of the quaternion is very close to 1.

See also length() and normalized().
*/
func (this *QQuaternion) Normalize() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion9normalizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:96
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QQuaternion inverted() const

/*
Returns the inverse of this quaternion. If this quaternion is null, then a null quaternion is returned.

This function was introduced in  Qt 5.5.

See also isNull() and length().
*/
func (this *QQuaternion) Inverted() *QQuaternion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion8invertedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}

// /usr/include/qt/QtGui/qquaternion.h:98
// index:0
// Public Visibility=Default Availability=Available
// [16] QQuaternion conjugated() const

/*
Returns the conjugate of this quaternion, which is (-x, -y, -z, scalar).

This function was introduced in  Qt 5.5.
*/
func (this *QQuaternion) Conjugated() *QQuaternion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion10conjugatedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}

// /usr/include/qt/QtGui/qquaternion.h:100
// index:0
// Public Visibility=Default Availability=Available
// [16] QQuaternion conjugate() const

/*

 */
func (this *QQuaternion) Conjugate() *QQuaternion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion9conjugateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}

// /usr/include/qt/QtGui/qquaternion.h:103
// index:0
// Public Visibility=Default Availability=Available
// [12] QVector3D rotatedVector(const QVector3D &) const

/*
Rotates vector with this quaternion to produce a new vector in 3D space. The following code:


  QVector3D result = q.rotatedVector(vector);



is equivalent to the following:


  QVector3D result = (q * QQuaternion(0, vector) * q.conjugated()).vector();
*/
func (this *QQuaternion) RotatedVector(vector QVector3D_ITF) *QVector3D /*123*/ {
	var convArg0 unsafe.Pointer
	if vector != nil && vector.QVector3D_PTR() != nil {
		convArg0 = vector.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion13rotatedVectorERK9QVector3D", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}

// /usr/include/qt/QtGui/qquaternion.h:105
// index:0
// Public Visibility=Default Availability=Available
// [16] QQuaternion & operator+=(const QQuaternion &)

/*

 */
func (this *QQuaternion) Operator_add_equal(quaternion QQuaternion_ITF) *QQuaternion {
	var convArg0 unsafe.Pointer
	if quaternion != nil && quaternion.QQuaternion_PTR() != nil {
		convArg0 = quaternion.QQuaternion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternionpLERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}

// /usr/include/qt/QtGui/qquaternion.h:106
// index:0
// Public Visibility=Default Availability=Available
// [16] QQuaternion & operator-=(const QQuaternion &)

/*

 */
func (this *QQuaternion) Operator_minus_equal(quaternion QQuaternion_ITF) *QQuaternion {
	var convArg0 unsafe.Pointer
	if quaternion != nil && quaternion.QQuaternion_PTR() != nil {
		convArg0 = quaternion.QQuaternion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternionmIERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}

// /usr/include/qt/QtGui/qquaternion.h:107
// index:0
// Public Visibility=Default Availability=Available
// [16] QQuaternion & operator*=(float)

/*

 */
func (this *QQuaternion) Operator_mul_equal(factor float32) *QQuaternion {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternionmLEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), factor)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}

// /usr/include/qt/QtGui/qquaternion.h:108
// index:1
// Public Visibility=Default Availability=Available
// [16] QQuaternion & operator*=(const QQuaternion &)

/*

 */
func (this *QQuaternion) Operator_mul_equal_1(quaternion QQuaternion_ITF) *QQuaternion {
	var convArg0 unsafe.Pointer
	if quaternion != nil && quaternion.QQuaternion_PTR() != nil {
		convArg0 = quaternion.QQuaternion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternionmLERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}

// /usr/include/qt/QtGui/qquaternion.h:109
// index:0
// Public Visibility=Default Availability=Available
// [16] QQuaternion & operator/=(float)

/*

 */
func (this *QQuaternion) Operator_div_equal(divisor float32) *QQuaternion {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaterniondVEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), divisor)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}

// /usr/include/qt/QtGui/qquaternion.h:124
// index:0
// Public Visibility=Default Availability=Available
// [16] QVector4D toVector4D() const

/*
Returns this quaternion as a 4D vector.
*/
func (this *QQuaternion) ToVector4D() *QVector4D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion10toVector4DEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector4DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector4D)
	return rv2
}

// /usr/include/qt/QtGui/qquaternion.h:130
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void getAxisAndAngle(QVector3D *, float *) const

/*
Extracts a 3D axis (x, y, z) and a rotating angle angle (in degrees) that corresponds to this quaternion.

This function was introduced in  Qt 5.5.

See also fromAxisAndAngle().
*/
func (this *QQuaternion) GetAxisAndAngle(axis QVector3D_ITF /*777 QVector3D **/, angle unsafe.Pointer /*666*/) {
	var convArg0 unsafe.Pointer
	if axis != nil && axis.QVector3D_PTR() != nil {
		convArg0 = axis.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion15getAxisAndAngleEP9QVector3DPf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, angle)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:133
// index:1
// Public Visibility=Default Availability=Available
// [-2] void getAxisAndAngle(float *, float *, float *, float *) const

/*
Extracts a 3D axis (x, y, z) and a rotating angle angle (in degrees) that corresponds to this quaternion.

This function was introduced in  Qt 5.5.

See also fromAxisAndAngle().
*/
func (this *QQuaternion) GetAxisAndAngle_1(x unsafe.Pointer /*666*/, y unsafe.Pointer /*666*/, z unsafe.Pointer /*666*/, angle unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion15getAxisAndAngleEPfS0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, z, angle)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:131
// index:0
// Public static Visibility=Default Availability=Available
// [16] QQuaternion fromAxisAndAngle(const QVector3D &, float)

/*
Creates a normalized quaternion that corresponds to rotating through angle degrees about the specified 3D axis.

See also getAxisAndAngle().
*/
func (this *QQuaternion) FromAxisAndAngle(axis QVector3D_ITF, angle float32) *QQuaternion /*123*/ {
	var convArg0 unsafe.Pointer
	if axis != nil && axis.QVector3D_PTR() != nil {
		convArg0 = axis.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion16fromAxisAndAngleERK9QVector3Df", qtrt.FFI_TYPE_POINTER, convArg0, angle)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}
func QQuaternion_FromAxisAndAngle(axis QVector3D_ITF, angle float32) *QQuaternion /*123*/ {
	var nilthis *QQuaternion
	rv := nilthis.FromAxisAndAngle(axis, angle)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:134
// index:1
// Public static Visibility=Default Availability=Available
// [16] QQuaternion fromAxisAndAngle(float, float, float, float)

/*
Creates a normalized quaternion that corresponds to rotating through angle degrees about the specified 3D axis.

See also getAxisAndAngle().
*/
func (this *QQuaternion) FromAxisAndAngle_1(x float32, y float32, z float32, angle float32) *QQuaternion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion16fromAxisAndAngleEffff", qtrt.FFI_TYPE_POINTER, x, y, z, angle)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}
func QQuaternion_FromAxisAndAngle_1(x float32, y float32, z float32, angle float32) *QQuaternion /*123*/ {
	var nilthis *QQuaternion
	rv := nilthis.FromAxisAndAngle_1(x, y, z, angle)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:138
// index:0
// Public inline Visibility=Default Availability=Available
// [12] QVector3D toEulerAngles() const

/*
This is an overloaded function.

Calculates roll, pitch, and yaw Euler angles (in degrees) that corresponds to this quaternion.

This function was introduced in  Qt 5.5.

See also fromEulerAngles().
*/
func (this *QQuaternion) ToEulerAngles() *QVector3D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion13toEulerAnglesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}

// /usr/include/qt/QtGui/qquaternion.h:139
// index:0
// Public static inline Visibility=Default Availability=Available
// [16] QQuaternion fromEulerAngles(const QVector3D &)

/*
Creates a quaternion that corresponds to a rotation of roll degrees around the z axis, pitch degrees around the x axis, and yaw degrees around the y axis (in that order).

This function was introduced in  Qt 5.5.

See also getEulerAngles().
*/
func (this *QQuaternion) FromEulerAngles(eulerAngles QVector3D_ITF) *QQuaternion /*123*/ {
	var convArg0 unsafe.Pointer
	if eulerAngles != nil && eulerAngles.QVector3D_PTR() != nil {
		convArg0 = eulerAngles.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion15fromEulerAnglesERK9QVector3D", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}
func QQuaternion_FromEulerAngles(eulerAngles QVector3D_ITF) *QQuaternion /*123*/ {
	var nilthis *QQuaternion
	rv := nilthis.FromEulerAngles(eulerAngles)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:142
// index:1
// Public static Visibility=Default Availability=Available
// [16] QQuaternion fromEulerAngles(float, float, float)

/*
Creates a quaternion that corresponds to a rotation of roll degrees around the z axis, pitch degrees around the x axis, and yaw degrees around the y axis (in that order).

This function was introduced in  Qt 5.5.

See also getEulerAngles().
*/
func (this *QQuaternion) FromEulerAngles_1(pitch float32, yaw float32, roll float32) *QQuaternion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion15fromEulerAnglesEfff", qtrt.FFI_TYPE_POINTER, pitch, yaw, roll)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}
func QQuaternion_FromEulerAngles_1(pitch float32, yaw float32, roll float32) *QQuaternion /*123*/ {
	var nilthis *QQuaternion
	rv := nilthis.FromEulerAngles_1(pitch, yaw, roll)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:141
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getEulerAngles(float *, float *, float *) const

/*
Calculates roll, pitch, and yaw Euler angles (in degrees) that corresponds to this quaternion.

This function was introduced in  Qt 5.5.

See also fromEulerAngles().
*/
func (this *QQuaternion) GetEulerAngles(pitch unsafe.Pointer /*666*/, yaw unsafe.Pointer /*666*/, roll unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion14getEulerAnglesEPfS0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pitch, yaw, roll)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getAxes(QVector3D *, QVector3D *, QVector3D *) const

/*
Returns the 3 orthonormal axes (xAxis, yAxis, zAxis) defining the quaternion.

This function was introduced in  Qt 5.5.

See also fromAxes() and toRotationMatrix().
*/
func (this *QQuaternion) GetAxes(xAxis QVector3D_ITF /*777 QVector3D **/, yAxis QVector3D_ITF /*777 QVector3D **/, zAxis QVector3D_ITF /*777 QVector3D **/) {
	var convArg0 unsafe.Pointer
	if xAxis != nil && xAxis.QVector3D_PTR() != nil {
		convArg0 = xAxis.QVector3D_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if yAxis != nil && yAxis.QVector3D_PTR() != nil {
		convArg1 = yAxis.QVector3D_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if zAxis != nil && zAxis.QVector3D_PTR() != nil {
		convArg2 = zAxis.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion7getAxesEP9QVector3DS1_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:149
// index:0
// Public static Visibility=Default Availability=Available
// [16] QQuaternion fromAxes(const QVector3D &, const QVector3D &, const QVector3D &)

/*
Constructs the quaternion using 3 axes (xAxis, yAxis, zAxis).

Note: The axes are assumed to be orthonormal.

This function was introduced in  Qt 5.5.

See also getAxes() and fromRotationMatrix().
*/
func (this *QQuaternion) FromAxes(xAxis QVector3D_ITF, yAxis QVector3D_ITF, zAxis QVector3D_ITF) *QQuaternion /*123*/ {
	var convArg0 unsafe.Pointer
	if xAxis != nil && xAxis.QVector3D_PTR() != nil {
		convArg0 = xAxis.QVector3D_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if yAxis != nil && yAxis.QVector3D_PTR() != nil {
		convArg1 = yAxis.QVector3D_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if zAxis != nil && zAxis.QVector3D_PTR() != nil {
		convArg2 = zAxis.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion8fromAxesERK9QVector3DS2_S2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}
func QQuaternion_FromAxes(xAxis QVector3D_ITF, yAxis QVector3D_ITF, zAxis QVector3D_ITF) *QQuaternion /*123*/ {
	var nilthis *QQuaternion
	rv := nilthis.FromAxes(xAxis, yAxis, zAxis)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:151
// index:0
// Public static Visibility=Default Availability=Available
// [16] QQuaternion fromDirection(const QVector3D &, const QVector3D &)

/*
Constructs the quaternion using specified forward direction direction and upward direction up. If the upward direction was not specified or the forward and upward vectors are collinear, a new orthonormal upward direction will be generated.

This function was introduced in  Qt 5.5.

See also fromAxes() and rotationTo().
*/
func (this *QQuaternion) FromDirection(direction QVector3D_ITF, up QVector3D_ITF) *QQuaternion /*123*/ {
	var convArg0 unsafe.Pointer
	if direction != nil && direction.QVector3D_PTR() != nil {
		convArg0 = direction.QVector3D_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if up != nil && up.QVector3D_PTR() != nil {
		convArg1 = up.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion13fromDirectionERK9QVector3DS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}
func QQuaternion_FromDirection(direction QVector3D_ITF, up QVector3D_ITF) *QQuaternion /*123*/ {
	var nilthis *QQuaternion
	rv := nilthis.FromDirection(direction, up)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:153
// index:0
// Public static Visibility=Default Availability=Available
// [16] QQuaternion rotationTo(const QVector3D &, const QVector3D &)

/*
Returns the shortest arc quaternion to rotate from the direction described by the vector from to the direction described by the vector to.

This function was introduced in  Qt 5.5.

See also fromDirection().
*/
func (this *QQuaternion) RotationTo(from QVector3D_ITF, to QVector3D_ITF) *QQuaternion /*123*/ {
	var convArg0 unsafe.Pointer
	if from != nil && from.QVector3D_PTR() != nil {
		convArg0 = from.QVector3D_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if to != nil && to.QVector3D_PTR() != nil {
		convArg1 = to.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion10rotationToERK9QVector3DS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}
func QQuaternion_RotationTo(from QVector3D_ITF, to QVector3D_ITF) *QQuaternion /*123*/ {
	var nilthis *QQuaternion
	rv := nilthis.RotationTo(from, to)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:156
// index:0
// Public static Visibility=Default Availability=Available
// [16] QQuaternion slerp(const QQuaternion &, const QQuaternion &, float)

/*
Interpolates along the shortest spherical path between the rotational positions q1 and q2. The value t should be between 0 and 1, indicating the spherical distance to travel between q1 and q2.

If t is less than or equal to 0, then q1 will be returned. If t is greater than or equal to 1, then q2 will be returned.

See also nlerp().
*/
func (this *QQuaternion) Slerp(q1 QQuaternion_ITF, q2 QQuaternion_ITF, t float32) *QQuaternion /*123*/ {
	var convArg0 unsafe.Pointer
	if q1 != nil && q1.QQuaternion_PTR() != nil {
		convArg0 = q1.QQuaternion_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if q2 != nil && q2.QQuaternion_PTR() != nil {
		convArg1 = q2.QQuaternion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion5slerpERKS_S1_f", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, t)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}
func QQuaternion_Slerp(q1 QQuaternion_ITF, q2 QQuaternion_ITF, t float32) *QQuaternion /*123*/ {
	var nilthis *QQuaternion
	rv := nilthis.Slerp(q1, q2, t)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:158
// index:0
// Public static Visibility=Default Availability=Available
// [16] QQuaternion nlerp(const QQuaternion &, const QQuaternion &, float)

/*
Interpolates along the shortest linear path between the rotational positions q1 and q2. The value t should be between 0 and 1, indicating the distance to travel between q1 and q2. The result will be normalized().

If t is less than or equal to 0, then q1 will be returned. If t is greater than or equal to 1, then q2 will be returned.

The nlerp() function is typically faster than slerp() and will give approximate results to spherical interpolation that are good enough for some applications.

See also slerp().
*/
func (this *QQuaternion) Nlerp(q1 QQuaternion_ITF, q2 QQuaternion_ITF, t float32) *QQuaternion /*123*/ {
	var convArg0 unsafe.Pointer
	if q1 != nil && q1.QQuaternion_PTR() != nil {
		convArg0 = q1.QQuaternion_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if q2 != nil && q2.QQuaternion_PTR() != nil {
		convArg1 = q2.QQuaternion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion5nlerpERKS_S1_f", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, t)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}
func QQuaternion_Nlerp(q1 QQuaternion_ITF, q2 QQuaternion_ITF, t float32) *QQuaternion /*123*/ {
	var nilthis *QQuaternion
	rv := nilthis.Nlerp(q1, q2, t)
	return rv
}

func DeleteQQuaternion(this *QQuaternion) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
