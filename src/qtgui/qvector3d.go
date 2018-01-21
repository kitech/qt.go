package qtgui

// /usr/include/qt/QtGui/qvector3d.h
// #include <qvector3d.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 31
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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQVector3DFromPointer(cthis unsafe.Pointer) *QVector3D {
	return &QVector3D{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qvector3d.h:60
// index:0
// Public inline
// void QVector3D()
func NewQVector3D() *QVector3D {
	cthis := qtrt.Calloc(1, 256) // 12
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3DC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:61
// index:1
// Public inline
// void QVector3D(Qt::Initialization)
func NewQVector3D_1(arg0 int) *QVector3D {
	cthis := qtrt.Calloc(1, 256) // 12
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3DC2EN2Qt14InitializationE", ffiqt.FFI_TYPE_VOID, cthis, &arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:62
// index:2
// Public inline
// void QVector3D(float, float, float)
func NewQVector3D_2(xpos float32, ypos float32, zpos float32) *QVector3D {
	cthis := qtrt.Calloc(1, 256) // 12
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3DC2Efff", ffiqt.FFI_TYPE_VOID, cthis, &xpos, &ypos, &zpos)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:64
// index:3
// Public inline
// void QVector3D(const class QPoint &)
func NewQVector3D_3(point *qtcore.QPoint) *QVector3D {
	cthis := qtrt.Calloc(1, 256) // 12
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3DC2ERK6QPoint", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:65
// index:4
// Public inline
// void QVector3D(const class QPointF &)
func NewQVector3D_4(point *qtcore.QPointF) *QVector3D {
	cthis := qtrt.Calloc(1, 256) // 12
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3DC2ERK7QPointF", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:67
// index:5
// Public
// void QVector3D(const class QVector2D &)
func NewQVector3D_5(vector *QVector2D) *QVector3D {
	cthis := qtrt.Calloc(1, 256) // 12
	var convArg0 = vector.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3DC2ERK9QVector2D", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:68
// index:6
// Public
// void QVector3D(const class QVector2D &, float)
func NewQVector3D_6(vector *QVector2D, zpos float32) *QVector3D {
	cthis := qtrt.Calloc(1, 256) // 12
	var convArg0 = vector.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3DC2ERK9QVector2Df", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &zpos)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:71
// index:7
// Public
// void QVector3D(const class QVector4D &)
func NewQVector3D_7(vector *QVector4D) *QVector3D {
	cthis := qtrt.Calloc(1, 256) // 12
	var convArg0 = vector.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3DC2ERK9QVector4D", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:74
// index:0
// Public
// bool isNull()
func (this *QVector3D) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qvector3d.h:76
// index:0
// Public inline
// float x()
func (this *QVector3D) X() float32 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D1xEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float32(rv) // 111
}

// /usr/include/qt/QtGui/qvector3d.h:77
// index:0
// Public inline
// float y()
func (this *QVector3D) Y() float32 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D1yEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float32(rv) // 111
}

// /usr/include/qt/QtGui/qvector3d.h:78
// index:0
// Public inline
// float z()
func (this *QVector3D) Z() float32 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D1zEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float32(rv) // 111
}

// /usr/include/qt/QtGui/qvector3d.h:80
// index:0
// Public
// void setX(float)
func (this *QVector3D) SetX(x float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3D4setXEf", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:81
// index:0
// Public
// void setY(float)
func (this *QVector3D) SetY(y float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3D4setYEf", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:82
// index:0
// Public
// void setZ(float)
func (this *QVector3D) SetZ(z float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3D4setZEf", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &z)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:87
// index:0
// Public
// float length()
func (this *QVector3D) Length() float32 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D6lengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float32(rv) // 111
}

// /usr/include/qt/QtGui/qvector3d.h:88
// index:0
// Public
// float lengthSquared()
func (this *QVector3D) LengthSquared() float32 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D13lengthSquaredEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float32(rv) // 111
}

// /usr/include/qt/QtGui/qvector3d.h:90
// index:0
// Public
// QVector3D normalized()
func (this *QVector3D) Normalized() *QVector3D /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D10normalizedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qvector3d.h:91
// index:0
// Public
// void normalize()
func (this *QVector3D) Normalize() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3D9normalizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:100
// index:0
// Public static
// float dotProduct(const class QVector3D &, const class QVector3D &)
func (this *QVector3D) DotProduct(v1 *QVector3D, v2 *QVector3D) float32 {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3D10dotProductERKS_S1_", ffiqt.FFI_TYPE_POINTER, v1, v2)
	gopp.ErrPrint(err, rv)
	// return rv
	return float32(rv) // 111
}
func QVector3D_DotProduct(v1 *QVector3D, v2 *QVector3D) float32 {
	var nilthis *QVector3D
	rv := nilthis.DotProduct(v1, v2)
	return rv
}

// /usr/include/qt/QtGui/qvector3d.h:101
// index:0
// Public static
// QVector3D crossProduct(const class QVector3D &, const class QVector3D &)
func (this *QVector3D) CrossProduct(v1 *QVector3D, v2 *QVector3D) *QVector3D /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3D12crossProductERKS_S1_", ffiqt.FFI_TYPE_POINTER, v1, v2)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QVector3D_CrossProduct(v1 *QVector3D, v2 *QVector3D) *QVector3D /*123*/ {
	var nilthis *QVector3D
	rv := nilthis.CrossProduct(v1, v2)
	return rv
}

// /usr/include/qt/QtGui/qvector3d.h:103
// index:0
// Public static
// QVector3D normal(const class QVector3D &, const class QVector3D &)
func (this *QVector3D) Normal(v1 *QVector3D, v2 *QVector3D) *QVector3D /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3D6normalERKS_S1_", ffiqt.FFI_TYPE_POINTER, v1, v2)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QVector3D_Normal(v1 *QVector3D, v2 *QVector3D) *QVector3D /*123*/ {
	var nilthis *QVector3D
	rv := nilthis.Normal(v1, v2)
	return rv
}

// /usr/include/qt/QtGui/qvector3d.h:104
// index:1
// Public static
// QVector3D normal(const class QVector3D &, const class QVector3D &, const class QVector3D &)
func (this *QVector3D) Normal_1(v1 *QVector3D, v2 *QVector3D, v3 *QVector3D) *QVector3D /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector3D6normalERKS_S1_S1_", ffiqt.FFI_TYPE_POINTER, v1, v2, v3)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QVector3D_Normal_1(v1 *QVector3D, v2 *QVector3D, v3 *QVector3D) *QVector3D /*123*/ {
	var nilthis *QVector3D
	rv := nilthis.Normal_1(v1, v2, v3)
	return rv
}

// /usr/include/qt/QtGui/qvector3d.h:107
// index:0
// Public
// QVector3D project(const class QMatrix4x4 &, const class QMatrix4x4 &, const class QRect &)
func (this *QVector3D) Project(modelView *QMatrix4x4, projection *QMatrix4x4, viewport *qtcore.QRect) *QVector3D /*123*/ {
	var convArg0 = modelView.GetCthis()
	var convArg1 = projection.GetCthis()
	var convArg2 = viewport.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D7projectERK10QMatrix4x4S2_RK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qvector3d.h:108
// index:0
// Public
// QVector3D unproject(const class QMatrix4x4 &, const class QMatrix4x4 &, const class QRect &)
func (this *QVector3D) Unproject(modelView *QMatrix4x4, projection *QMatrix4x4, viewport *qtcore.QRect) *QVector3D /*123*/ {
	var convArg0 = modelView.GetCthis()
	var convArg1 = projection.GetCthis()
	var convArg2 = viewport.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D9unprojectERK10QMatrix4x4S2_RK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qvector3d.h:110
// index:0
// Public
// float distanceToPoint(const class QVector3D &)
func (this *QVector3D) DistanceToPoint(point *QVector3D) float32 {
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D15distanceToPointERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float32(rv) // 111
}

// /usr/include/qt/QtGui/qvector3d.h:111
// index:0
// Public
// float distanceToPlane(const class QVector3D &, const class QVector3D &)
func (this *QVector3D) DistanceToPlane(plane *QVector3D, normal *QVector3D) float32 {
	var convArg0 = plane.GetCthis()
	var convArg1 = normal.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D15distanceToPlaneERKS_S1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float32(rv) // 111
}

// /usr/include/qt/QtGui/qvector3d.h:112
// index:1
// Public
// float distanceToPlane(const class QVector3D &, const class QVector3D &, const class QVector3D &)
func (this *QVector3D) DistanceToPlane_1(plane1 *QVector3D, plane2 *QVector3D, plane3 *QVector3D) float32 {
	var convArg0 = plane1.GetCthis()
	var convArg1 = plane2.GetCthis()
	var convArg2 = plane3.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D15distanceToPlaneERKS_S1_S1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float32(rv) // 111
}

// /usr/include/qt/QtGui/qvector3d.h:113
// index:0
// Public
// float distanceToLine(const class QVector3D &, const class QVector3D &)
func (this *QVector3D) DistanceToLine(point *QVector3D, direction *QVector3D) float32 {
	var convArg0 = point.GetCthis()
	var convArg1 = direction.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D14distanceToLineERKS_S1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float32(rv) // 111
}

// /usr/include/qt/QtGui/qvector3d.h:129
// index:0
// Public
// QVector2D toVector2D()
func (this *QVector3D) ToVector2D() *QVector2D /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D10toVector2DEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qvector3d.h:132
// index:0
// Public
// QVector4D toVector4D()
func (this *QVector3D) ToVector4D() *QVector4D /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D10toVector4DEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQVector4DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qvector3d.h:135
// index:0
// Public inline
// QPoint toPoint()
func (this *QVector3D) ToPoint() *qtcore.QPoint /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D7toPointEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qvector3d.h:136
// index:0
// Public inline
// QPointF toPointF()
func (this *QVector3D) ToPointF() *qtcore.QPointF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector3D8toPointFEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
