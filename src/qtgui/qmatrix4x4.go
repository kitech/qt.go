package qtgui

// /usr/include/qt/QtGui/qmatrix4x4.h
// #include <qmatrix4x4.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 42
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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMatrix4x4) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQMatrix4x4FromPointer(cthis unsafe.Pointer) *QMatrix4x4 {
	return &QMatrix4x4{&qtrt.CObject{cthis}}
}
func (*QMatrix4x4) NewFromPointer(cthis unsafe.Pointer) *QMatrix4x4 {
	return NewQMatrix4x4FromPointer(cthis)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:62
// index:0
// Public inline
// void QMatrix4x4()
func NewQMatrix4x4() *QMatrix4x4 {
	cthis := qtrt.Calloc(1, 256) // 68
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x4C2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQMatrix4x4FromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qmatrix4x4.h:63
// index:1
// Public inline
// void QMatrix4x4(Qt::Initialization)
func NewQMatrix4x4_1(arg0 int) *QMatrix4x4 {
	cthis := qtrt.Calloc(1, 256) // 68
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x4C2EN2Qt14InitializationE", ffiqt.FFI_TYPE_VOID, cthis, arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQMatrix4x4FromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qmatrix4x4.h:64
// index:2
// Public
// void QMatrix4x4(const float *)
func NewQMatrix4x4_2(values unsafe.Pointer /*666*/) *QMatrix4x4 {
	cthis := qtrt.Calloc(1, 256) // 68
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x4C2EPKf", ffiqt.FFI_TYPE_VOID, cthis, &values)
	gopp.ErrPrint(err, rv)
	gothis := NewQMatrix4x4FromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qmatrix4x4.h:65
// index:3
// Public inline
// void QMatrix4x4(float, float, float, float, float, float, float, float, float, float, float, float, float, float, float, float)
func NewQMatrix4x4_3(m11 float32, m12 float32, m13 float32, m14 float32, m21 float32, m22 float32, m23 float32, m24 float32, m31 float32, m32 float32, m33 float32, m34 float32, m41 float32, m42 float32, m43 float32, m44 float32) *QMatrix4x4 {
	cthis := qtrt.Calloc(1, 256) // 68
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x4C2Effffffffffffffff", ffiqt.FFI_TYPE_VOID, cthis, m11, m12, m13, m14, m21, m22, m23, m24, m31, m32, m33, m34, m41, m42, m43, m44)
	gopp.ErrPrint(err, rv)
	gothis := NewQMatrix4x4FromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qmatrix4x4.h:73
// index:4
// Public
// void QMatrix4x4(const float *, int, int)
func NewQMatrix4x4_4(values unsafe.Pointer /*666*/, cols int, rows int) *QMatrix4x4 {
	cthis := qtrt.Calloc(1, 256) // 68
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x4C2EPKfii", ffiqt.FFI_TYPE_VOID, cthis, &values, cols, rows)
	gopp.ErrPrint(err, rv)
	gothis := NewQMatrix4x4FromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qmatrix4x4.h:74
// index:5
// Public
// void QMatrix4x4(const class QTransform &)
func NewQMatrix4x4_5(transform *QTransform) *QMatrix4x4 {
	cthis := qtrt.Calloc(1, 256) // 68
	var convArg0 = transform.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x4C2ERK10QTransform", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQMatrix4x4FromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qmatrix4x4.h:75
// index:6
// Public
// void QMatrix4x4(const class QMatrix &)
func NewQMatrix4x4_6(matrix *QMatrix) *QMatrix4x4 {
	cthis := qtrt.Calloc(1, 256) // 68
	var convArg0 = matrix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x4C2ERK7QMatrix", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQMatrix4x4FromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qmatrix4x4.h:81
// index:0
// Public inline
// QVector4D column(int)
func (this *QMatrix4x4) Column(index int) *QVector4D /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x46columnEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQVector4DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:82
// index:0
// Public inline
// void setColumn(int, const class QVector4D &)
func (this *QMatrix4x4) SetColumn(index int, value *QVector4D) {
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x49setColumnEiRK9QVector4D", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:84
// index:0
// Public inline
// QVector4D row(int)
func (this *QMatrix4x4) Row(index int) *QVector4D /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x43rowEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQVector4DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:85
// index:0
// Public inline
// void setRow(int, const class QVector4D &)
func (this *QMatrix4x4) SetRow(index int, value *QVector4D) {
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x46setRowEiRK9QVector4D", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:88
// index:0
// Public inline
// bool isAffine()
func (this *QMatrix4x4) IsAffine() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x48isAffineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qmatrix4x4.h:90
// index:0
// Public inline
// bool isIdentity()
func (this *QMatrix4x4) IsIdentity() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x410isIdentityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qmatrix4x4.h:91
// index:0
// Public inline
// void setToIdentity()
func (this *QMatrix4x4) SetToIdentity() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x413setToIdentityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:93
// index:0
// Public inline
// void fill(float)
func (this *QMatrix4x4) Fill(value float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x44fillEf", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:95
// index:0
// Public
// double determinant()
func (this *QMatrix4x4) Determinant() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x411determinantEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 111
}

// /usr/include/qt/QtGui/qmatrix4x4.h:96
// index:0
// Public
// QMatrix4x4 inverted(_Bool *)
func (this *QMatrix4x4) Inverted(invertible unsafe.Pointer /*666*/) *QMatrix4x4 /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x48invertedEPb", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), &invertible)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:97
// index:0
// Public
// QMatrix4x4 transposed()
func (this *QMatrix4x4) Transposed() *QMatrix4x4 /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x410transposedEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:131
// index:0
// Public
// void scale(const class QVector3D &)
func (this *QMatrix4x4) Scale(vector *QVector3D) {
	var convArg0 = vector.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x45scaleERK9QVector3D", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:135
// index:1
// Public
// void scale(float, float)
func (this *QMatrix4x4) Scale_1(x float32, y float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x45scaleEff", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:136
// index:2
// Public
// void scale(float, float, float)
func (this *QMatrix4x4) Scale_2(x float32, y float32, z float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x45scaleEfff", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, z)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:137
// index:3
// Public
// void scale(float)
func (this *QMatrix4x4) Scale_3(factor float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x45scaleEf", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), factor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:132
// index:0
// Public
// void translate(const class QVector3D &)
func (this *QMatrix4x4) Translate(vector *QVector3D) {
	var convArg0 = vector.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x49translateERK9QVector3D", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:138
// index:1
// Public
// void translate(float, float)
func (this *QMatrix4x4) Translate_1(x float32, y float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x49translateEff", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:139
// index:2
// Public
// void translate(float, float, float)
func (this *QMatrix4x4) Translate_2(x float32, y float32, z float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x49translateEfff", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, z)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:133
// index:0
// Public
// void rotate(float, const class QVector3D &)
func (this *QMatrix4x4) Rotate(angle float32, vector *QVector3D) {
	var convArg1 = vector.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x46rotateEfRK9QVector3D", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), angle, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:140
// index:1
// Public
// void rotate(float, float, float, float)
func (this *QMatrix4x4) Rotate_1(angle float32, x float32, y float32, z float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x46rotateEffff", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), angle, x, y, z)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:142
// index:2
// Public
// void rotate(const class QQuaternion &)
func (this *QMatrix4x4) Rotate_2(quaternion *QQuaternion) {
	var convArg0 = quaternion.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x46rotateERK11QQuaternion", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:145
// index:0
// Public
// void ortho(const class QRect &)
func (this *QMatrix4x4) Ortho(rect *qtcore.QRect) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x45orthoERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:146
// index:1
// Public
// void ortho(const class QRectF &)
func (this *QMatrix4x4) Ortho_1(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x45orthoERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:147
// index:2
// Public
// void ortho(float, float, float, float, float, float)
func (this *QMatrix4x4) Ortho_2(left float32, right float32, bottom float32, top float32, nearPlane float32, farPlane float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x45orthoEffffff", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), left, right, bottom, top, nearPlane, farPlane)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:148
// index:0
// Public
// void frustum(float, float, float, float, float, float)
func (this *QMatrix4x4) Frustum(left float32, right float32, bottom float32, top float32, nearPlane float32, farPlane float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x47frustumEffffff", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), left, right, bottom, top, nearPlane, farPlane)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:149
// index:0
// Public
// void perspective(float, float, float, float)
func (this *QMatrix4x4) Perspective(verticalAngle float32, aspectRatio float32, nearPlane float32, farPlane float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x411perspectiveEffff", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), verticalAngle, aspectRatio, nearPlane, farPlane)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:151
// index:0
// Public
// void lookAt(const class QVector3D &, const class QVector3D &, const class QVector3D &)
func (this *QMatrix4x4) LookAt(eye *QVector3D, center *QVector3D, up *QVector3D) {
	var convArg0 = eye.GetCthis()
	var convArg1 = center.GetCthis()
	var convArg2 = up.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x46lookAtERK9QVector3DS2_S2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:153
// index:0
// Public
// void viewport(const class QRectF &)
func (this *QMatrix4x4) Viewport(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x48viewportERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:154
// index:1
// Public
// void viewport(float, float, float, float, float, float)
func (this *QMatrix4x4) Viewport_1(left float32, bottom float32, width float32, height float32, nearPlane float32, farPlane float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x48viewportEffffff", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), left, bottom, width, height, nearPlane, farPlane)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:155
// index:0
// Public
// void flipCoordinates()
func (this *QMatrix4x4) FlipCoordinates() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x415flipCoordinatesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:157
// index:0
// Public
// void copyDataTo(float *)
func (this *QMatrix4x4) CopyDataTo(values unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x410copyDataToEPf", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &values)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmatrix4x4.h:159
// index:0
// Public
// QMatrix toAffine()
func (this *QMatrix4x4) ToAffine() *QMatrix /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x48toAffineEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:160
// index:0
// Public
// QTransform toTransform()
func (this *QMatrix4x4) ToTransform() *QTransform /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x411toTransformEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:161
// index:1
// Public
// QTransform toTransform(float)
func (this *QMatrix4x4) ToTransform_1(distanceToPlane float32) *QTransform /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x411toTransformEf", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), distanceToPlane)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:163
// index:0
// Public
// QPoint map(const class QPoint &)
func (this *QMatrix4x4) Map(point *qtcore.QPoint) *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x43mapERK6QPoint", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:164
// index:1
// Public
// QPointF map(const class QPointF &)
func (this *QMatrix4x4) Map_1(point *qtcore.QPointF) *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x43mapERK7QPointF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:166
// index:2
// Public
// QVector3D map(const class QVector3D &)
func (this *QMatrix4x4) Map_2(point *QVector3D) *QVector3D /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x43mapERK9QVector3D", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:170
// index:3
// Public
// QVector4D map(const class QVector4D &)
func (this *QMatrix4x4) Map_3(point *QVector4D) *QVector4D /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x43mapERK9QVector4D", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQVector4DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:167
// index:0
// Public
// QVector3D mapVector(const class QVector3D &)
func (this *QMatrix4x4) MapVector(vector *QVector3D) *QVector3D /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = vector.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x49mapVectorERK9QVector3D", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:172
// index:0
// Public
// QRect mapRect(const class QRect &)
func (this *QMatrix4x4) MapRect(rect *qtcore.QRect) *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x47mapRectERK5QRect", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:173
// index:1
// Public
// QRectF mapRect(const class QRectF &)
func (this *QMatrix4x4) MapRect_1(rect *qtcore.QRectF) *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x47mapRectERK6QRectF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qmatrix4x4.h:178
// index:0
// Public inline
// float * data()
func (this *QMatrix4x4) Data() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x44dataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qmatrix4x4.h:179
// index:1
// Public inline
// const float * data()
func (this *QMatrix4x4) Data_1() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x44dataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qmatrix4x4.h:180
// index:0
// Public inline
// const float * constData()
func (this *QMatrix4x4) ConstData() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMatrix4x49constDataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qmatrix4x4.h:182
// index:0
// Public
// void optimize()
func (this *QMatrix4x4) Optimize() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMatrix4x48optimizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
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
