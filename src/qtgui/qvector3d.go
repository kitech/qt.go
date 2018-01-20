//  header block begin
// /usr/include/qt/QtGui/qvector3d.h
// #include <qvector3d.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 35
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
type QVector3D struct {
	*qtrt.CObject
}

func (this *QVector3D) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtGui/qvector3d.h:60
// index:0
// inline
// void QVector3D()
func NewQVector3D() *QVector3D {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3DC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(cthis)
	return gothis
}
func NewQVector3DFromPointer(cthis unsafe.Pointer) *QVector3D {
	return &QVector3D{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qvector3d.h:61
// index:1
// inline
// void QVector3D(Qt::Initialization)
func NewQVector3D_1(arg0 int) *QVector3D {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3DC2EN2Qt14InitializationE", ffiqt.FFI_TYPE_VOID, cthis, &arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:62
// index:2
// inline
// void QVector3D(float, float, float)
func NewQVector3D_2(xpos float32, ypos float32, zpos float32) *QVector3D {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3DC2Efff", ffiqt.FFI_TYPE_VOID, cthis, &xpos, &ypos, &zpos)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:64
// index:3
// inline
// void QVector3D(const class QPoint &)
func NewQVector3D_3(point unsafe.Pointer) *QVector3D {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3DC2ERK6QPoint", ffiqt.FFI_TYPE_VOID, cthis, point)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:65
// index:4
// inline
// void QVector3D(const class QPointF &)
func NewQVector3D_4(point unsafe.Pointer) *QVector3D {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3DC2ERK7QPointF", ffiqt.FFI_TYPE_VOID, cthis, point)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:67
// index:5
// void QVector3D(const class QVector2D &)
func NewQVector3D_5(vector unsafe.Pointer) *QVector3D {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3DC2ERK9QVector2D", ffiqt.FFI_TYPE_VOID, cthis, vector)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:68
// index:6
// void QVector3D(const class QVector2D &, float)
func NewQVector3D_6(vector unsafe.Pointer, zpos float32) *QVector3D {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3DC2ERK9QVector2Df", ffiqt.FFI_TYPE_VOID, cthis, vector, &zpos)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:71
// index:7
// void QVector3D(const class QVector4D &)
func NewQVector3D_7(vector unsafe.Pointer) *QVector3D {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3DC2ERK9QVector4D", ffiqt.FFI_TYPE_VOID, cthis, vector)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:74
// index:0
// bool isNull()
func (this *QVector3D) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D6isNullEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:76
// index:0
// inline
// float x()
func (this *QVector3D) X() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D1xEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:77
// index:0
// inline
// float y()
func (this *QVector3D) Y() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D1yEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:78
// index:0
// inline
// float z()
func (this *QVector3D) Z() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D1zEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:80
// index:0
// void setX(float)
func (this *QVector3D) SetX(x float32) {
	// 0: (, x float), (&x)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3D4setXEf", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:81
// index:0
// void setY(float)
func (this *QVector3D) SetY(y float32) {
	// 0: (, y float), (&y)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3D4setYEf", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:82
// index:0
// void setZ(float)
func (this *QVector3D) SetZ(z float32) {
	// 0: (, z float), (&z)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3D4setZEf", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &z)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:87
// index:0
// float length()
func (this *QVector3D) Length() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D6lengthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:88
// index:0
// float lengthSquared()
func (this *QVector3D) LengthSquared() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D13lengthSquaredEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:90
// index:0
// QVector3D normalized()
func (this *QVector3D) Normalized() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D10normalizedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:91
// index:0
// void normalize()
func (this *QVector3D) Normalize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3D9normalizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:100
// index:0
// static
// float dotProduct(const class QVector3D &, const class QVector3D &)
func (this *QVector3D) DotProduct(v1 unsafe.Pointer, v2 unsafe.Pointer) {
	// 0: (v1 const QVector3D &, v2 const QVector3D &), (v1, v2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3D10dotProductERKS_S1_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QVector3D_DotProduct(v1 unsafe.Pointer, v2 unsafe.Pointer) {
	// 0: (v1 const QVector3D &, v2 const QVector3D &), (v1, v2)
	var nilthis *QVector3D
	nilthis.DotProduct(v1, v2)
}

// /usr/include/qt/QtGui/qvector3d.h:101
// index:0
// static
// QVector3D crossProduct(const class QVector3D &, const class QVector3D &)
func (this *QVector3D) CrossProduct(v1 unsafe.Pointer, v2 unsafe.Pointer) {
	// 0: (v1 const QVector3D &, v2 const QVector3D &), (v1, v2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3D12crossProductERKS_S1_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QVector3D_CrossProduct(v1 unsafe.Pointer, v2 unsafe.Pointer) {
	// 0: (v1 const QVector3D &, v2 const QVector3D &), (v1, v2)
	var nilthis *QVector3D
	nilthis.CrossProduct(v1, v2)
}

// /usr/include/qt/QtGui/qvector3d.h:103
// index:0
// static
// QVector3D normal(const class QVector3D &, const class QVector3D &)
func (this *QVector3D) Normal(v1 unsafe.Pointer, v2 unsafe.Pointer) {
	// 0: (v1 const QVector3D &, v2 const QVector3D &), (v1, v2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3D6normalERKS_S1_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QVector3D_Normal(v1 unsafe.Pointer, v2 unsafe.Pointer) {
	// 0: (v1 const QVector3D &, v2 const QVector3D &), (v1, v2)
	var nilthis *QVector3D
	nilthis.Normal(v1, v2)
}

// /usr/include/qt/QtGui/qvector3d.h:104
// index:1
// static
// QVector3D normal(const class QVector3D &, const class QVector3D &, const class QVector3D &)
func (this *QVector3D) Normal_1(v1 unsafe.Pointer, v2 unsafe.Pointer, v3 unsafe.Pointer) {
	// 1: (v1 const QVector3D &, v2 const QVector3D &, v3 const QVector3D &), (v1, v2, v3)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3D6normalERKS_S1_S1_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QVector3D_Normal_1(v1 unsafe.Pointer, v2 unsafe.Pointer, v3 unsafe.Pointer) {
	// 1: (v1 const QVector3D &, v2 const QVector3D &, v3 const QVector3D &), (v1, v2, v3)
	var nilthis *QVector3D
	nilthis.Normal_1(v1, v2, v3)
}

// /usr/include/qt/QtGui/qvector3d.h:107
// index:0
// QVector3D project(const class QMatrix4x4 &, const class QMatrix4x4 &, const class QRect &)
func (this *QVector3D) Project(modelView unsafe.Pointer, projection unsafe.Pointer, viewport unsafe.Pointer) {
	// 0: (, modelView const QMatrix4x4 &, projection const QMatrix4x4 &, viewport const QRect &), (modelView, projection, viewport)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D7projectERK10QMatrix4x4S2_RK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), modelView, projection, viewport)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:108
// index:0
// QVector3D unproject(const class QMatrix4x4 &, const class QMatrix4x4 &, const class QRect &)
func (this *QVector3D) Unproject(modelView unsafe.Pointer, projection unsafe.Pointer, viewport unsafe.Pointer) {
	// 0: (, modelView const QMatrix4x4 &, projection const QMatrix4x4 &, viewport const QRect &), (modelView, projection, viewport)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D9unprojectERK10QMatrix4x4S2_RK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), modelView, projection, viewport)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:110
// index:0
// float distanceToPoint(const class QVector3D &)
func (this *QVector3D) DistanceToPoint(point unsafe.Pointer) {
	// 0: (, point const QVector3D &), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D15distanceToPointERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:111
// index:0
// float distanceToPlane(const class QVector3D &, const class QVector3D &)
func (this *QVector3D) DistanceToPlane(plane unsafe.Pointer, normal unsafe.Pointer) {
	// 0: (, plane const QVector3D &, normal const QVector3D &), (plane, normal)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D15distanceToPlaneERKS_S1_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), plane, normal)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:112
// index:1
// float distanceToPlane(const class QVector3D &, const class QVector3D &, const class QVector3D &)
func (this *QVector3D) DistanceToPlane_1(plane1 unsafe.Pointer, plane2 unsafe.Pointer, plane3 unsafe.Pointer) {
	// 1: (, plane1 const QVector3D &, plane2 const QVector3D &, plane3 const QVector3D &), (plane1, plane2, plane3)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D15distanceToPlaneERKS_S1_S1_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), plane1, plane2, plane3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:113
// index:0
// float distanceToLine(const class QVector3D &, const class QVector3D &)
func (this *QVector3D) DistanceToLine(point unsafe.Pointer, direction unsafe.Pointer) {
	// 0: (, point const QVector3D &, direction const QVector3D &), (point, direction)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D14distanceToLineERKS_S1_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), point, direction)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:129
// index:0
// QVector2D toVector2D()
func (this *QVector3D) ToVector2D() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D10toVector2DEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:132
// index:0
// QVector4D toVector4D()
func (this *QVector3D) ToVector4D() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D10toVector4DEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:135
// index:0
// inline
// QPoint toPoint()
func (this *QVector3D) ToPoint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D7toPointEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:136
// index:0
// inline
// QPointF toPointF()
func (this *QVector3D) ToPointF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D8toPointFEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
