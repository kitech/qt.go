//  header block begin
// /usr/include/qt/QtGui/qquaternion.h
// #include <qquaternion.h>
// #include <QtGui>
package qtgui

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
type QQuaternion struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qquaternion.h:59
// index:0
// void QQuaternion()
func NewQQuaternion() *QQuaternion {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternionC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QQuaternion{cthis}
}

// /usr/include/qt/QtGui/qquaternion.h:60
// index:1
// inline
// void QQuaternion(Qt::Initialization)
func NewQQuaternion_1(arg0 int) *QQuaternion {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternionC2EN2Qt14InitializationE", ffiqt.FFI_TYPE_VOID, cthis, &arg0)
	gopp.ErrPrint(err, rv)
	return &QQuaternion{cthis}
}

// /usr/include/qt/QtGui/qquaternion.h:61
// index:2
// void QQuaternion(float, float, float, float)
func NewQQuaternion_2(scalar float32, xpos float32, ypos float32, zpos float32) *QQuaternion {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternionC2Effff", ffiqt.FFI_TYPE_VOID, cthis, &scalar, &xpos, &ypos, &zpos)
	gopp.ErrPrint(err, rv)
	return &QQuaternion{cthis}
}

// /usr/include/qt/QtGui/qquaternion.h:63
// index:3
// void QQuaternion(float, const class QVector3D &)
func NewQQuaternion_3(scalar float32, vector unsafe.Pointer) *QQuaternion {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternionC2EfRK9QVector3D", ffiqt.FFI_TYPE_VOID, cthis, &scalar, vector)
	gopp.ErrPrint(err, rv)
	return &QQuaternion{cthis}
}

// /usr/include/qt/QtGui/qquaternion.h:66
// index:4
// void QQuaternion(const class QVector4D &)
func NewQQuaternion_4(vector unsafe.Pointer) *QQuaternion {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternionC2ERK9QVector4D", ffiqt.FFI_TYPE_VOID, cthis, vector)
	gopp.ErrPrint(err, rv)
	return &QQuaternion{cthis}
}

// /usr/include/qt/QtGui/qquaternion.h:69
// index:0
// bool isNull()
func (this *QQuaternion) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion6isNullEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:70
// index:0
// bool isIdentity()
func (this *QQuaternion) IsIdentity() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion10isIdentityEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:73
// index:0
// QVector3D vector()
func (this *QQuaternion) Vector() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion6vectorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:74
// index:0
// void setVector(const class QVector3D &)
func (this *QQuaternion) SetVector(vector unsafe.Pointer) {
	// 0: (, const QVector3D & vector), (vector)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion9setVectorERK9QVector3D", ffiqt.FFI_TYPE_VOID, this.cthis, vector)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:76
// index:1
// void setVector(float, float, float)
func (this *QQuaternion) SetVector_1(x float32, y float32, z float32) {
	// 1: (, float x, float y, float z), (&x, &y, &z)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion9setVectorEfff", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y, &z)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:78
// index:0
// float x()
func (this *QQuaternion) X() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion1xEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:79
// index:0
// float y()
func (this *QQuaternion) Y() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion1yEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:80
// index:0
// float z()
func (this *QQuaternion) Z() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion1zEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:81
// index:0
// float scalar()
func (this *QQuaternion) Scalar() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion6scalarEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:83
// index:0
// void setX(float)
func (this *QQuaternion) SetX(x float32) {
	// 0: (, float x), (&x)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion4setXEf", ffiqt.FFI_TYPE_VOID, this.cthis, &x)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:84
// index:0
// void setY(float)
func (this *QQuaternion) SetY(y float32) {
	// 0: (, float y), (&y)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion4setYEf", ffiqt.FFI_TYPE_VOID, this.cthis, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:85
// index:0
// void setZ(float)
func (this *QQuaternion) SetZ(z float32) {
	// 0: (, float z), (&z)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion4setZEf", ffiqt.FFI_TYPE_VOID, this.cthis, &z)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:86
// index:0
// void setScalar(float)
func (this *QQuaternion) SetScalar(scalar float32) {
	// 0: (, float scalar), (&scalar)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion9setScalarEf", ffiqt.FFI_TYPE_VOID, this.cthis, &scalar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:88
// index:0
// static inline
// float dotProduct(const class QQuaternion &, const class QQuaternion &)
func (this *QQuaternion) DotProduct(q1 unsafe.Pointer, q2 unsafe.Pointer) {
	// 0: (const QQuaternion & q1, const QQuaternion & q2), (q1, q2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion10dotProductERKS_S1_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QQuaternion_DotProduct(q1 unsafe.Pointer, q2 unsafe.Pointer) {
	// 0: (const QQuaternion & q1, const QQuaternion & q2), (q1, q2)
	var nilthis *QQuaternion
	nilthis.DotProduct(q1, q2)
}

// /usr/include/qt/QtGui/qquaternion.h:90
// index:0
// float length()
func (this *QQuaternion) Length() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion6lengthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:91
// index:0
// float lengthSquared()
func (this *QQuaternion) LengthSquared() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion13lengthSquaredEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:93
// index:0
// QQuaternion normalized()
func (this *QQuaternion) Normalized() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion10normalizedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:94
// index:0
// void normalize()
func (this *QQuaternion) Normalize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion9normalizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:96
// index:0
// inline
// QQuaternion inverted()
func (this *QQuaternion) Inverted() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion8invertedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:98
// index:0
// QQuaternion conjugated()
func (this *QQuaternion) Conjugated() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion10conjugatedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:100
// index:0
// QQuaternion conjugate()
func (this *QQuaternion) Conjugate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion9conjugateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:103
// index:0
// QVector3D rotatedVector(const class QVector3D &)
func (this *QQuaternion) RotatedVector(vector unsafe.Pointer) {
	// 0: (, const QVector3D & vector), (vector)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion13rotatedVectorERK9QVector3D", ffiqt.FFI_TYPE_VOID, this.cthis, vector)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:124
// index:0
// QVector4D toVector4D()
func (this *QQuaternion) ToVector4D() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion10toVector4DEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:130
// index:0
// inline
// void getAxisAndAngle(class QVector3D *, float *)
func (this *QQuaternion) GetAxisAndAngle(axis unsafe.Pointer, angle unsafe.Pointer) {
	// 0: (, QVector3D * axis, float * angle), (axis, angle)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion15getAxisAndAngleEP9QVector3DPf", ffiqt.FFI_TYPE_VOID, this.cthis, axis, angle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:133
// index:1
// void getAxisAndAngle(float *, float *, float *, float *)
func (this *QQuaternion) GetAxisAndAngle_1(x unsafe.Pointer, y unsafe.Pointer, z unsafe.Pointer, angle unsafe.Pointer) {
	// 1: (, float * x, float * y, float * z, float * angle), (x, y, z, angle)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion15getAxisAndAngleEPfS0_S0_S0_", ffiqt.FFI_TYPE_VOID, this.cthis, x, y, z, angle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:131
// index:0
// static
// QQuaternion fromAxisAndAngle(const class QVector3D &, float)
func (this *QQuaternion) FromAxisAndAngle(axis unsafe.Pointer, angle float32) {
	// 0: (const QVector3D & axis, float angle), (axis, angle)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion16fromAxisAndAngleERK9QVector3Df", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QQuaternion_FromAxisAndAngle(axis unsafe.Pointer, angle float32) {
	// 0: (const QVector3D & axis, float angle), (axis, angle)
	var nilthis *QQuaternion
	nilthis.FromAxisAndAngle(axis, angle)
}

// /usr/include/qt/QtGui/qquaternion.h:134
// index:1
// static
// QQuaternion fromAxisAndAngle(float, float, float, float)
func (this *QQuaternion) FromAxisAndAngle_1(x float32, y float32, z float32, angle float32) {
	// 1: (float x, float y, float z, float angle), (x, y, z, angle)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion16fromAxisAndAngleEffff", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QQuaternion_FromAxisAndAngle_1(x float32, y float32, z float32, angle float32) {
	// 1: (float x, float y, float z, float angle), (x, y, z, angle)
	var nilthis *QQuaternion
	nilthis.FromAxisAndAngle_1(x, y, z, angle)
}

// /usr/include/qt/QtGui/qquaternion.h:138
// index:0
// inline
// QVector3D toEulerAngles()
func (this *QQuaternion) ToEulerAngles() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion13toEulerAnglesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:139
// index:0
// static inline
// QQuaternion fromEulerAngles(const class QVector3D &)
func (this *QQuaternion) FromEulerAngles(eulerAngles unsafe.Pointer) {
	// 0: (const QVector3D & eulerAngles), (eulerAngles)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion15fromEulerAnglesERK9QVector3D", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QQuaternion_FromEulerAngles(eulerAngles unsafe.Pointer) {
	// 0: (const QVector3D & eulerAngles), (eulerAngles)
	var nilthis *QQuaternion
	nilthis.FromEulerAngles(eulerAngles)
}

// /usr/include/qt/QtGui/qquaternion.h:142
// index:1
// static
// QQuaternion fromEulerAngles(float, float, float)
func (this *QQuaternion) FromEulerAngles_1(pitch float32, yaw float32, roll float32) {
	// 1: (float pitch, float yaw, float roll), (pitch, yaw, roll)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion15fromEulerAnglesEfff", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QQuaternion_FromEulerAngles_1(pitch float32, yaw float32, roll float32) {
	// 1: (float pitch, float yaw, float roll), (pitch, yaw, roll)
	var nilthis *QQuaternion
	nilthis.FromEulerAngles_1(pitch, yaw, roll)
}

// /usr/include/qt/QtGui/qquaternion.h:141
// index:0
// void getEulerAngles(float *, float *, float *)
func (this *QQuaternion) GetEulerAngles(pitch unsafe.Pointer, yaw unsafe.Pointer, roll unsafe.Pointer) {
	// 0: (, float * pitch, float * yaw, float * roll), (pitch, yaw, roll)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion14getEulerAnglesEPfS0_S0_", ffiqt.FFI_TYPE_VOID, this.cthis, pitch, yaw, roll)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:144
// index:0
// QMatrix3x3 toRotationMatrix()
func (this *QQuaternion) ToRotationMatrix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion16toRotationMatrixEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:148
// index:0
// void getAxes(class QVector3D *, class QVector3D *, class QVector3D *)
func (this *QQuaternion) GetAxes(xAxis unsafe.Pointer, yAxis unsafe.Pointer, zAxis unsafe.Pointer) {
	// 0: (, QVector3D * xAxis, QVector3D * yAxis, QVector3D * zAxis), (xAxis, yAxis, zAxis)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion7getAxesEP9QVector3DS1_S1_", ffiqt.FFI_TYPE_VOID, this.cthis, xAxis, yAxis, zAxis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:149
// index:0
// static
// QQuaternion fromAxes(const class QVector3D &, const class QVector3D &, const class QVector3D &)
func (this *QQuaternion) FromAxes(xAxis unsafe.Pointer, yAxis unsafe.Pointer, zAxis unsafe.Pointer) {
	// 0: (const QVector3D & xAxis, const QVector3D & yAxis, const QVector3D & zAxis), (xAxis, yAxis, zAxis)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion8fromAxesERK9QVector3DS2_S2_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QQuaternion_FromAxes(xAxis unsafe.Pointer, yAxis unsafe.Pointer, zAxis unsafe.Pointer) {
	// 0: (const QVector3D & xAxis, const QVector3D & yAxis, const QVector3D & zAxis), (xAxis, yAxis, zAxis)
	var nilthis *QQuaternion
	nilthis.FromAxes(xAxis, yAxis, zAxis)
}

// /usr/include/qt/QtGui/qquaternion.h:151
// index:0
// static
// QQuaternion fromDirection(const class QVector3D &, const class QVector3D &)
func (this *QQuaternion) FromDirection(direction unsafe.Pointer, up unsafe.Pointer) {
	// 0: (const QVector3D & direction, const QVector3D & up), (direction, up)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion13fromDirectionERK9QVector3DS2_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QQuaternion_FromDirection(direction unsafe.Pointer, up unsafe.Pointer) {
	// 0: (const QVector3D & direction, const QVector3D & up), (direction, up)
	var nilthis *QQuaternion
	nilthis.FromDirection(direction, up)
}

// /usr/include/qt/QtGui/qquaternion.h:153
// index:0
// static
// QQuaternion rotationTo(const class QVector3D &, const class QVector3D &)
func (this *QQuaternion) RotationTo(from unsafe.Pointer, to unsafe.Pointer) {
	// 0: (const QVector3D & from, const QVector3D & to), (from, to)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion10rotationToERK9QVector3DS2_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QQuaternion_RotationTo(from unsafe.Pointer, to unsafe.Pointer) {
	// 0: (const QVector3D & from, const QVector3D & to), (from, to)
	var nilthis *QQuaternion
	nilthis.RotationTo(from, to)
}

// /usr/include/qt/QtGui/qquaternion.h:156
// index:0
// static
// QQuaternion slerp(const class QQuaternion &, const class QQuaternion &, float)
func (this *QQuaternion) Slerp(q1 unsafe.Pointer, q2 unsafe.Pointer, t float32) {
	// 0: (const QQuaternion & q1, const QQuaternion & q2, float t), (q1, q2, t)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion5slerpERKS_S1_f", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QQuaternion_Slerp(q1 unsafe.Pointer, q2 unsafe.Pointer, t float32) {
	// 0: (const QQuaternion & q1, const QQuaternion & q2, float t), (q1, q2, t)
	var nilthis *QQuaternion
	nilthis.Slerp(q1, q2, t)
}

// /usr/include/qt/QtGui/qquaternion.h:158
// index:0
// static
// QQuaternion nlerp(const class QQuaternion &, const class QQuaternion &, float)
func (this *QQuaternion) Nlerp(q1 unsafe.Pointer, q2 unsafe.Pointer, t float32) {
	// 0: (const QQuaternion & q1, const QQuaternion & q2, float t), (q1, q2, t)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion5nlerpERKS_S1_f", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QQuaternion_Nlerp(q1 unsafe.Pointer, q2 unsafe.Pointer, t float32) {
	// 0: (const QQuaternion & q1, const QQuaternion & q2, float t), (q1, q2, t)
	var nilthis *QQuaternion
	nilthis.Nlerp(q1, q2, t)
}

//  body block end
