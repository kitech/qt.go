package qtgui

// /usr/include/qt/QtGui/qquaternion.h
// #include <qquaternion.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 29
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QQuaternion struct {
	*qtrt.CObject
}

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
func NewQQuaternion() *QQuaternion {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternionC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQuaternion)
	return gothis
}

// /usr/include/qt/QtGui/qquaternion.h:60
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QQuaternion(Qt::Initialization)
func NewQQuaternion_1(arg0 int) *QQuaternion {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternionC2EN2Qt14InitializationE", qtrt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQuaternion)
	return gothis
}

// /usr/include/qt/QtGui/qquaternion.h:61
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QQuaternion(float, float, float, float)
func NewQQuaternion_2(scalar float32, xpos float32, ypos float32, zpos float32) *QQuaternion {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternionC2Effff", qtrt.FFI_TYPE_POINTER, scalar, xpos, ypos, zpos)
	gopp.ErrPrint(err, rv)
	gothis := NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQuaternion)
	return gothis
}

// /usr/include/qt/QtGui/qquaternion.h:63
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QQuaternion(float, const QVector3D &)
func NewQQuaternion_3(scalar float32, vector *QVector3D) *QQuaternion {
	var convArg1 = vector.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternionC2EfRK9QVector3D", qtrt.FFI_TYPE_POINTER, scalar, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQuaternion)
	return gothis
}

// /usr/include/qt/QtGui/qquaternion.h:66
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QQuaternion(const QVector4D &)
func NewQQuaternion_4(vector *QVector4D) *QQuaternion {
	var convArg0 = vector.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternionC2ERK9QVector4D", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQuaternion)
	return gothis
}

// /usr/include/qt/QtGui/qquaternion.h:69
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QQuaternion) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qquaternion.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isIdentity()
func (this *QQuaternion) IsIdentity() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion10isIdentityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qquaternion.h:73
// index:0
// Public Visibility=Default Availability=Available
// [12] QVector3D vector()
func (this *QQuaternion) Vector() *QVector3D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion6vectorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}

// /usr/include/qt/QtGui/qquaternion.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVector(const QVector3D &)
func (this *QQuaternion) SetVector(vector *QVector3D) {
	var convArg0 = vector.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion9setVectorERK9QVector3D", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:76
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setVector(float, float, float)
func (this *QQuaternion) SetVector_1(x float32, y float32, z float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion9setVectorEfff", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, z)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:78
// index:0
// Public Visibility=Default Availability=Available
// [4] float x()
func (this *QQuaternion) X() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion1xEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qquaternion.h:79
// index:0
// Public Visibility=Default Availability=Available
// [4] float y()
func (this *QQuaternion) Y() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion1yEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qquaternion.h:80
// index:0
// Public Visibility=Default Availability=Available
// [4] float z()
func (this *QQuaternion) Z() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion1zEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qquaternion.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] float scalar()
func (this *QQuaternion) Scalar() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion6scalarEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qquaternion.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setX(float)
func (this *QQuaternion) SetX(x float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion4setXEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setY(float)
func (this *QQuaternion) SetY(y float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion4setYEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setZ(float)
func (this *QQuaternion) SetZ(z float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion4setZEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), z)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScalar(float)
func (this *QQuaternion) SetScalar(scalar float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion9setScalarEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), scalar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:88
// index:0
// Public static inline Visibility=Default Availability=Available
// [4] float dotProduct(const QQuaternion &, const QQuaternion &)
func (this *QQuaternion) DotProduct(q1 *QQuaternion, q2 *QQuaternion) float32 {
	var convArg0 = q1.GetCthis()
	var convArg1 = q2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion10dotProductERKS_S1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}
func QQuaternion_DotProduct(q1 *QQuaternion, q2 *QQuaternion) float32 {
	var nilthis *QQuaternion
	rv := nilthis.DotProduct(q1, q2)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:90
// index:0
// Public Visibility=Default Availability=Available
// [4] float length()
func (this *QQuaternion) Length() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion6lengthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qquaternion.h:91
// index:0
// Public Visibility=Default Availability=Available
// [4] float lengthSquared()
func (this *QQuaternion) LengthSquared() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion13lengthSquaredEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qquaternion.h:93
// index:0
// Public Visibility=Default Availability=Available
// [16] QQuaternion normalized()
func (this *QQuaternion) Normalized() *QQuaternion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion10normalizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}

// /usr/include/qt/QtGui/qquaternion.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void normalize()
func (this *QQuaternion) Normalize() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion9normalizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:96
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QQuaternion inverted()
func (this *QQuaternion) Inverted() *QQuaternion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion8invertedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}

// /usr/include/qt/QtGui/qquaternion.h:98
// index:0
// Public Visibility=Default Availability=Available
// [16] QQuaternion conjugated()
func (this *QQuaternion) Conjugated() *QQuaternion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion10conjugatedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}

// /usr/include/qt/QtGui/qquaternion.h:100
// index:0
// Public Visibility=Default Availability=Available
// [16] QQuaternion conjugate()
func (this *QQuaternion) Conjugate() *QQuaternion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion9conjugateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}

// /usr/include/qt/QtGui/qquaternion.h:103
// index:0
// Public Visibility=Default Availability=Available
// [12] QVector3D rotatedVector(const QVector3D &)
func (this *QQuaternion) RotatedVector(vector *QVector3D) *QVector3D /*123*/ {
	var convArg0 = vector.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion13rotatedVectorERK9QVector3D", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}

// /usr/include/qt/QtGui/qquaternion.h:124
// index:0
// Public Visibility=Default Availability=Available
// [16] QVector4D toVector4D()
func (this *QQuaternion) ToVector4D() *QVector4D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion10toVector4DEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQVector4DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector4D)
	return rv2
}

// /usr/include/qt/QtGui/qquaternion.h:130
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void getAxisAndAngle(QVector3D *, float *)
func (this *QQuaternion) GetAxisAndAngle(axis *QVector3D /*777 QVector3D **/, angle unsafe.Pointer /*666*/) {
	var convArg0 = axis.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion15getAxisAndAngleEP9QVector3DPf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &angle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:133
// index:1
// Public Visibility=Default Availability=Available
// [-2] void getAxisAndAngle(float *, float *, float *, float *)
func (this *QQuaternion) GetAxisAndAngle_1(x unsafe.Pointer /*666*/, y unsafe.Pointer /*666*/, z unsafe.Pointer /*666*/, angle unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion15getAxisAndAngleEPfS0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &x, &y, &z, &angle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:131
// index:0
// Public static Visibility=Default Availability=Available
// [16] QQuaternion fromAxisAndAngle(const QVector3D &, float)
func (this *QQuaternion) FromAxisAndAngle(axis *QVector3D, angle float32) *QQuaternion /*123*/ {
	var convArg0 = axis.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion16fromAxisAndAngleERK9QVector3Df", qtrt.FFI_TYPE_POINTER, convArg0, angle)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}
func QQuaternion_FromAxisAndAngle(axis *QVector3D, angle float32) *QQuaternion /*123*/ {
	var nilthis *QQuaternion
	rv := nilthis.FromAxisAndAngle(axis, angle)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:134
// index:1
// Public static Visibility=Default Availability=Available
// [16] QQuaternion fromAxisAndAngle(float, float, float, float)
func (this *QQuaternion) FromAxisAndAngle_1(x float32, y float32, z float32, angle float32) *QQuaternion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion16fromAxisAndAngleEffff", qtrt.FFI_TYPE_POINTER, x, y, z, angle)
	gopp.ErrPrint(err, rv)
	// return rv
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
// [12] QVector3D toEulerAngles()
func (this *QQuaternion) ToEulerAngles() *QVector3D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion13toEulerAnglesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}

// /usr/include/qt/QtGui/qquaternion.h:139
// index:0
// Public static inline Visibility=Default Availability=Available
// [16] QQuaternion fromEulerAngles(const QVector3D &)
func (this *QQuaternion) FromEulerAngles(eulerAngles *QVector3D) *QQuaternion /*123*/ {
	var convArg0 = eulerAngles.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion15fromEulerAnglesERK9QVector3D", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}
func QQuaternion_FromEulerAngles(eulerAngles *QVector3D) *QQuaternion /*123*/ {
	var nilthis *QQuaternion
	rv := nilthis.FromEulerAngles(eulerAngles)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:142
// index:1
// Public static Visibility=Default Availability=Available
// [16] QQuaternion fromEulerAngles(float, float, float)
func (this *QQuaternion) FromEulerAngles_1(pitch float32, yaw float32, roll float32) *QQuaternion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion15fromEulerAnglesEfff", qtrt.FFI_TYPE_POINTER, pitch, yaw, roll)
	gopp.ErrPrint(err, rv)
	// return rv
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
// [-2] void getEulerAngles(float *, float *, float *)
func (this *QQuaternion) GetEulerAngles(pitch unsafe.Pointer /*666*/, yaw unsafe.Pointer /*666*/, roll unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion14getEulerAnglesEPfS0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &pitch, &yaw, &roll)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getAxes(QVector3D *, QVector3D *, QVector3D *)
func (this *QQuaternion) GetAxes(xAxis *QVector3D /*777 QVector3D **/, yAxis *QVector3D /*777 QVector3D **/, zAxis *QVector3D /*777 QVector3D **/) {
	var convArg0 = xAxis.GetCthis()
	var convArg1 = yAxis.GetCthis()
	var convArg2 = zAxis.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQuaternion7getAxesEP9QVector3DS1_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:149
// index:0
// Public static Visibility=Default Availability=Available
// [16] QQuaternion fromAxes(const QVector3D &, const QVector3D &, const QVector3D &)
func (this *QQuaternion) FromAxes(xAxis *QVector3D, yAxis *QVector3D, zAxis *QVector3D) *QQuaternion /*123*/ {
	var convArg0 = xAxis.GetCthis()
	var convArg1 = yAxis.GetCthis()
	var convArg2 = zAxis.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion8fromAxesERK9QVector3DS2_S2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}
func QQuaternion_FromAxes(xAxis *QVector3D, yAxis *QVector3D, zAxis *QVector3D) *QQuaternion /*123*/ {
	var nilthis *QQuaternion
	rv := nilthis.FromAxes(xAxis, yAxis, zAxis)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:151
// index:0
// Public static Visibility=Default Availability=Available
// [16] QQuaternion fromDirection(const QVector3D &, const QVector3D &)
func (this *QQuaternion) FromDirection(direction *QVector3D, up *QVector3D) *QQuaternion /*123*/ {
	var convArg0 = direction.GetCthis()
	var convArg1 = up.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion13fromDirectionERK9QVector3DS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}
func QQuaternion_FromDirection(direction *QVector3D, up *QVector3D) *QQuaternion /*123*/ {
	var nilthis *QQuaternion
	rv := nilthis.FromDirection(direction, up)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:153
// index:0
// Public static Visibility=Default Availability=Available
// [16] QQuaternion rotationTo(const QVector3D &, const QVector3D &)
func (this *QQuaternion) RotationTo(from *QVector3D, to *QVector3D) *QQuaternion /*123*/ {
	var convArg0 = from.GetCthis()
	var convArg1 = to.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion10rotationToERK9QVector3DS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}
func QQuaternion_RotationTo(from *QVector3D, to *QVector3D) *QQuaternion /*123*/ {
	var nilthis *QQuaternion
	rv := nilthis.RotationTo(from, to)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:156
// index:0
// Public static Visibility=Default Availability=Available
// [16] QQuaternion slerp(const QQuaternion &, const QQuaternion &, float)
func (this *QQuaternion) Slerp(q1 *QQuaternion, q2 *QQuaternion, t float32) *QQuaternion /*123*/ {
	var convArg0 = q1.GetCthis()
	var convArg1 = q2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion5slerpERKS_S1_f", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, t)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}
func QQuaternion_Slerp(q1 *QQuaternion, q2 *QQuaternion, t float32) *QQuaternion /*123*/ {
	var nilthis *QQuaternion
	rv := nilthis.Slerp(q1, q2, t)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:158
// index:0
// Public static Visibility=Default Availability=Available
// [16] QQuaternion nlerp(const QQuaternion &, const QQuaternion &, float)
func (this *QQuaternion) Nlerp(q1 *QQuaternion, q2 *QQuaternion, t float32) *QQuaternion /*123*/ {
	var convArg0 = q1.GetCthis()
	var convArg1 = q2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternion5nlerpERKS_S1_f", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, t)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}
func QQuaternion_Nlerp(q1 *QQuaternion, q2 *QQuaternion, t float32) *QQuaternion /*123*/ {
	var nilthis *QQuaternion
	rv := nilthis.Nlerp(q1, q2, t)
	return rv
}

func DeleteQQuaternion(this *QQuaternion) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuaternionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
