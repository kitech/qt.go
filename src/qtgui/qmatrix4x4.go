//  header block begin
// /usr/include/qt/QtGui/qmatrix4x4.h
// #include <qmatrix4x4.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 43
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
type QMatrix4x4 struct {
	*qtrt.CObject
}

func (this *QMatrix4x4) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtGui/qmatrix4x4.h:62
// index:0
// inline
// void QMatrix4x4()
func NewQMatrix4x4() *QMatrix4x4 {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x4C2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQMatrix4x4FromPointer(cthis)
	return gothis
}
func NewQMatrix4x4FromPointer(cthis unsafe.Pointer) *QMatrix4x4 {
	return &QMatrix4x4{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qmatrix4x4.h:63
// index:1
// inline
// void QMatrix4x4(Qt::Initialization)
func NewQMatrix4x4_1(arg0 int) *QMatrix4x4 {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x4C2EN2Qt14InitializationE", ffiqt.FFI_TYPE_VOID, cthis, &arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQMatrix4x4FromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qmatrix4x4.h:64
// index:2
// void QMatrix4x4(const float *)
func NewQMatrix4x4_2(values unsafe.Pointer) *QMatrix4x4 {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x4C2EPKf", ffiqt.FFI_TYPE_VOID, cthis, values)
	gopp.ErrPrint(err, rv)
	gothis := NewQMatrix4x4FromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qmatrix4x4.h:65
// index:3
// inline
// void QMatrix4x4(float, float, float, float, float, float, float, float, float, float, float, float, float, float, float, float)
func NewQMatrix4x4_3(m11 float32, m12 float32, m13 float32, m14 float32, m21 float32, m22 float32, m23 float32, m24 float32, m31 float32, m32 float32, m33 float32, m34 float32, m41 float32, m42 float32, m43 float32, m44 float32) *QMatrix4x4 {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x4C2Effffffffffffffff", ffiqt.FFI_TYPE_VOID, cthis, &m11, &m12, &m13, &m14, &m21, &m22, &m23, &m24, &m31, &m32, &m33, &m34, &m41, &m42, &m43, &m44)
	gopp.ErrPrint(err, rv)
	gothis := NewQMatrix4x4FromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qmatrix4x4.h:73
// index:4
// void QMatrix4x4(const float *, int, int)
func NewQMatrix4x4_4(values unsafe.Pointer, cols int, rows int) *QMatrix4x4 {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x4C2EPKfii", ffiqt.FFI_TYPE_VOID, cthis, values, &cols, &rows)
	gopp.ErrPrint(err, rv)
	gothis := NewQMatrix4x4FromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qmatrix4x4.h:74
// index:5
// void QMatrix4x4(const class QTransform &)
func NewQMatrix4x4_5(transform unsafe.Pointer) *QMatrix4x4 {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x4C2ERK10QTransform", ffiqt.FFI_TYPE_VOID, cthis, transform)
	gopp.ErrPrint(err, rv)
	gothis := NewQMatrix4x4FromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qmatrix4x4.h:75
// index:6
// void QMatrix4x4(const class QMatrix &)
func NewQMatrix4x4_6(matrix unsafe.Pointer) *QMatrix4x4 {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x4C2ERK7QMatrix", ffiqt.FFI_TYPE_VOID, cthis, matrix)
	gopp.ErrPrint(err, rv)
	gothis := NewQMatrix4x4FromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qmatrix4x4.h:81
// index:0
// inline
// QVector4D column(int)
func (this *QMatrix4x4) Column(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x46columnEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:82
// index:0
// inline
// void setColumn(int, const class QVector4D &)
func (this *QMatrix4x4) SetColumn(index int, value unsafe.Pointer) {
	// 0: (, index int, value const QVector4D &), (&index, value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x49setColumnEiRK9QVector4D", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:84
// index:0
// inline
// QVector4D row(int)
func (this *QMatrix4x4) Row(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x43rowEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:85
// index:0
// inline
// void setRow(int, const class QVector4D &)
func (this *QMatrix4x4) SetRow(index int, value unsafe.Pointer) {
	// 0: (, index int, value const QVector4D &), (&index, value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x46setRowEiRK9QVector4D", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:88
// index:0
// inline
// bool isAffine()
func (this *QMatrix4x4) IsAffine() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x48isAffineEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:90
// index:0
// inline
// bool isIdentity()
func (this *QMatrix4x4) IsIdentity() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x410isIdentityEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:91
// index:0
// inline
// void setToIdentity()
func (this *QMatrix4x4) SetToIdentity() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x413setToIdentityEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:93
// index:0
// inline
// void fill(float)
func (this *QMatrix4x4) Fill(value float32) {
	// 0: (, value float), (&value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x44fillEf", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:95
// index:0
// double determinant()
func (this *QMatrix4x4) Determinant() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x411determinantEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:96
// index:0
// QMatrix4x4 inverted(_Bool *)
func (this *QMatrix4x4) Inverted(invertible unsafe.Pointer) {
	// 0: (, invertible bool *), (invertible)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x48invertedEPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), invertible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:97
// index:0
// QMatrix4x4 transposed()
func (this *QMatrix4x4) Transposed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x410transposedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:98
// index:0
// QMatrix3x3 normalMatrix()
func (this *QMatrix4x4) NormalMatrix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x412normalMatrixEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:131
// index:0
// void scale(const class QVector3D &)
func (this *QMatrix4x4) Scale(vector unsafe.Pointer) {
	// 0: (, vector const QVector3D &), (vector)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x45scaleERK9QVector3D", ffiqt.FFI_TYPE_VOID, this.GetCthis(), vector)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:135
// index:1
// void scale(float, float)
func (this *QMatrix4x4) Scale_1(x float32, y float32) {
	// 1: (, x float, y float), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x45scaleEff", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:136
// index:2
// void scale(float, float, float)
func (this *QMatrix4x4) Scale_2(x float32, y float32, z float32) {
	// 2: (, x float, y float, z float), (&x, &y, &z)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x45scaleEfff", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &z)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:137
// index:3
// void scale(float)
func (this *QMatrix4x4) Scale_3(factor float32) {
	// 3: (, factor float), (&factor)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x45scaleEf", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &factor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:132
// index:0
// void translate(const class QVector3D &)
func (this *QMatrix4x4) Translate(vector unsafe.Pointer) {
	// 0: (, vector const QVector3D &), (vector)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x49translateERK9QVector3D", ffiqt.FFI_TYPE_VOID, this.GetCthis(), vector)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:138
// index:1
// void translate(float, float)
func (this *QMatrix4x4) Translate_1(x float32, y float32) {
	// 1: (, x float, y float), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x49translateEff", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:139
// index:2
// void translate(float, float, float)
func (this *QMatrix4x4) Translate_2(x float32, y float32, z float32) {
	// 2: (, x float, y float, z float), (&x, &y, &z)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x49translateEfff", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &z)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:133
// index:0
// void rotate(float, const class QVector3D &)
func (this *QMatrix4x4) Rotate(angle float32, vector unsafe.Pointer) {
	// 0: (, angle float, vector const QVector3D &), (&angle, vector)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x46rotateEfRK9QVector3D", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &angle, vector)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:140
// index:1
// void rotate(float, float, float, float)
func (this *QMatrix4x4) Rotate_1(angle float32, x float32, y float32, z float32) {
	// 1: (, angle float, x float, y float, z float), (&angle, &x, &y, &z)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x46rotateEffff", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &angle, &x, &y, &z)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:142
// index:2
// void rotate(const class QQuaternion &)
func (this *QMatrix4x4) Rotate_2(quaternion unsafe.Pointer) {
	// 2: (, quaternion const QQuaternion &), (quaternion)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x46rotateERK11QQuaternion", ffiqt.FFI_TYPE_VOID, this.GetCthis(), quaternion)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:145
// index:0
// void ortho(const class QRect &)
func (this *QMatrix4x4) Ortho(rect unsafe.Pointer) {
	// 0: (, rect const QRect &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x45orthoERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:146
// index:1
// void ortho(const class QRectF &)
func (this *QMatrix4x4) Ortho_1(rect unsafe.Pointer) {
	// 1: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x45orthoERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:147
// index:2
// void ortho(float, float, float, float, float, float)
func (this *QMatrix4x4) Ortho_2(left float32, right float32, bottom float32, top float32, nearPlane float32, farPlane float32) {
	// 2: (, left float, right float, bottom float, top float, nearPlane float, farPlane float), (&left, &right, &bottom, &top, &nearPlane, &farPlane)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x45orthoEffffff", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &left, &right, &bottom, &top, &nearPlane, &farPlane)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:148
// index:0
// void frustum(float, float, float, float, float, float)
func (this *QMatrix4x4) Frustum(left float32, right float32, bottom float32, top float32, nearPlane float32, farPlane float32) {
	// 0: (, left float, right float, bottom float, top float, nearPlane float, farPlane float), (&left, &right, &bottom, &top, &nearPlane, &farPlane)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x47frustumEffffff", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &left, &right, &bottom, &top, &nearPlane, &farPlane)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:149
// index:0
// void perspective(float, float, float, float)
func (this *QMatrix4x4) Perspective(verticalAngle float32, aspectRatio float32, nearPlane float32, farPlane float32) {
	// 0: (, verticalAngle float, aspectRatio float, nearPlane float, farPlane float), (&verticalAngle, &aspectRatio, &nearPlane, &farPlane)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x411perspectiveEffff", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &verticalAngle, &aspectRatio, &nearPlane, &farPlane)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:151
// index:0
// void lookAt(const class QVector3D &, const class QVector3D &, const class QVector3D &)
func (this *QMatrix4x4) LookAt(eye unsafe.Pointer, center unsafe.Pointer, up unsafe.Pointer) {
	// 0: (, eye const QVector3D &, center const QVector3D &, up const QVector3D &), (eye, center, up)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x46lookAtERK9QVector3DS2_S2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), eye, center, up)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:153
// index:0
// void viewport(const class QRectF &)
func (this *QMatrix4x4) Viewport(rect unsafe.Pointer) {
	// 0: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x48viewportERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:154
// index:1
// void viewport(float, float, float, float, float, float)
func (this *QMatrix4x4) Viewport_1(left float32, bottom float32, width float32, height float32, nearPlane float32, farPlane float32) {
	// 1: (, left float, bottom float, width float, height float, nearPlane float, farPlane float), (&left, &bottom, &width, &height, &nearPlane, &farPlane)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x48viewportEffffff", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &left, &bottom, &width, &height, &nearPlane, &farPlane)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:155
// index:0
// void flipCoordinates()
func (this *QMatrix4x4) FlipCoordinates() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x415flipCoordinatesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:157
// index:0
// void copyDataTo(float *)
func (this *QMatrix4x4) CopyDataTo(values unsafe.Pointer) {
	// 0: (, values float *), (values)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x410copyDataToEPf", ffiqt.FFI_TYPE_VOID, this.GetCthis(), values)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:159
// index:0
// QMatrix toAffine()
func (this *QMatrix4x4) ToAffine() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x48toAffineEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:160
// index:0
// QTransform toTransform()
func (this *QMatrix4x4) ToTransform() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x411toTransformEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:161
// index:1
// QTransform toTransform(float)
func (this *QMatrix4x4) ToTransform_1(distanceToPlane float32) {
	// 1: (, distanceToPlane float), (&distanceToPlane)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x411toTransformEf", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &distanceToPlane)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:163
// index:0
// QPoint map(const class QPoint &)
func (this *QMatrix4x4) Map(point unsafe.Pointer) {
	// 0: (, point const QPoint &), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x43mapERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:164
// index:1
// QPointF map(const class QPointF &)
func (this *QMatrix4x4) Map_1(point unsafe.Pointer) {
	// 1: (, point const QPointF &), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x43mapERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:166
// index:2
// QVector3D map(const class QVector3D &)
func (this *QMatrix4x4) Map_2(point unsafe.Pointer) {
	// 2: (, point const QVector3D &), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x43mapERK9QVector3D", ffiqt.FFI_TYPE_VOID, this.GetCthis(), point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:170
// index:3
// QVector4D map(const class QVector4D &)
func (this *QMatrix4x4) Map_3(point unsafe.Pointer) {
	// 3: (, point const QVector4D &), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x43mapERK9QVector4D", ffiqt.FFI_TYPE_VOID, this.GetCthis(), point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:167
// index:0
// QVector3D mapVector(const class QVector3D &)
func (this *QMatrix4x4) MapVector(vector unsafe.Pointer) {
	// 0: (, vector const QVector3D &), (vector)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x49mapVectorERK9QVector3D", ffiqt.FFI_TYPE_VOID, this.GetCthis(), vector)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:172
// index:0
// QRect mapRect(const class QRect &)
func (this *QMatrix4x4) MapRect(rect unsafe.Pointer) {
	// 0: (, rect const QRect &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x47mapRectERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:173
// index:1
// QRectF mapRect(const class QRectF &)
func (this *QMatrix4x4) MapRect_1(rect unsafe.Pointer) {
	// 1: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x47mapRectERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:178
// index:0
// inline
// float * data()
func (this *QMatrix4x4) Data() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x44dataEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:179
// index:1
// inline
// const float * data()
func (this *QMatrix4x4) Data_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x44dataEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:180
// index:0
// inline
// const float * constData()
func (this *QMatrix4x4) ConstData() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x49constDataEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:182
// index:0
// void optimize()
func (this *QMatrix4x4) Optimize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x48optimizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
