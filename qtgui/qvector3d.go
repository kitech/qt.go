package qtgui

// /usr/include/qt/QtGui/qvector3d.h
// #include <qvector3d.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 32
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QVector3D struct {
	*qtrt.CObject
}
type QVector3D_ITF interface {
	QVector3D_PTR() *QVector3D
}

func (ptr *QVector3D) QVector3D_PTR() *QVector3D { return ptr }

func (this *QVector3D) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QVector3D) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQVector3DFromPointer(cthis unsafe.Pointer) *QVector3D {
	return &QVector3D{&qtrt.CObject{cthis}}
}
func (*QVector3D) NewFromPointer(cthis unsafe.Pointer) *QVector3D {
	return NewQVector3DFromPointer(cthis)
}

// /usr/include/qt/QtGui/qvector3d.h:60
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QVector3D()
func NewQVector3D() *QVector3D {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3DC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector3D)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:61
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QVector3D(Qt::Initialization)
func NewQVector3D_1(arg0 int) *QVector3D {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3DC2EN2Qt14InitializationE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector3D)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:62
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QVector3D(float, float, float)
func NewQVector3D_2(xpos float32, ypos float32, zpos float32) *QVector3D {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3DC2Efff", qtrt.FFI_TYPE_POINTER, xpos, ypos, zpos)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector3D)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:64
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void QVector3D(const QPoint &)
func NewQVector3D_3(point *qtcore.QPoint) *QVector3D {
	var convArg0 = point.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3DC2ERK6QPoint", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector3D)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:65
// index:4
// Public inline Visibility=Default Availability=Available
// [-2] void QVector3D(const QPointF &)
func NewQVector3D_4(point *qtcore.QPointF) *QVector3D {
	var convArg0 = point.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3DC2ERK7QPointF", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector3D)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:67
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QVector3D(const QVector2D &)
func NewQVector3D_5(vector *QVector2D) *QVector3D {
	var convArg0 = vector.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3DC2ERK9QVector2D", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector3D)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:68
// index:6
// Public Visibility=Default Availability=Available
// [-2] void QVector3D(const QVector2D &, float)
func NewQVector3D_6(vector *QVector2D, zpos float32) *QVector3D {
	var convArg0 = vector.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3DC2ERK9QVector2Df", qtrt.FFI_TYPE_POINTER, convArg0, zpos)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector3D)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:71
// index:7
// Public Visibility=Default Availability=Available
// [-2] void QVector3D(const QVector4D &)
func NewQVector3D_7(vector *QVector4D) *QVector3D {
	var convArg0 = vector.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3DC2ERK9QVector4D", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector3D)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QVector3D) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qvector3d.h:76
// index:0
// Public inline Visibility=Default Availability=Available
// [4] float x()
func (this *QVector3D) X() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D1xEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector3d.h:77
// index:0
// Public inline Visibility=Default Availability=Available
// [4] float y()
func (this *QVector3D) Y() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D1yEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector3d.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [4] float z()
func (this *QVector3D) Z() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D1zEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector3d.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setX(float)
func (this *QVector3D) SetX(x float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3D4setXEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setY(float)
func (this *QVector3D) SetY(y float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3D4setYEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setZ(float)
func (this *QVector3D) SetZ(z float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3D4setZEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), z)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:87
// index:0
// Public Visibility=Default Availability=Available
// [4] float length()
func (this *QVector3D) Length() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D6lengthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector3d.h:88
// index:0
// Public Visibility=Default Availability=Available
// [4] float lengthSquared()
func (this *QVector3D) LengthSquared() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D13lengthSquaredEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector3d.h:90
// index:0
// Public Visibility=Default Availability=Available
// [12] QVector3D normalized()
func (this *QVector3D) Normalized() *QVector3D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D10normalizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}

// /usr/include/qt/QtGui/qvector3d.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void normalize()
func (this *QVector3D) Normalize() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3D9normalizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:100
// index:0
// Public static Visibility=Default Availability=Available
// [4] float dotProduct(const QVector3D &, const QVector3D &)
func (this *QVector3D) DotProduct(v1 *QVector3D, v2 *QVector3D) float32 {
	var convArg0 = v1.GetCthis()
	var convArg1 = v2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3D10dotProductERKS_S1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}
func QVector3D_DotProduct(v1 *QVector3D, v2 *QVector3D) float32 {
	var nilthis *QVector3D
	rv := nilthis.DotProduct(v1, v2)
	return rv
}

// /usr/include/qt/QtGui/qvector3d.h:101
// index:0
// Public static Visibility=Default Availability=Available
// [12] QVector3D crossProduct(const QVector3D &, const QVector3D &)
func (this *QVector3D) CrossProduct(v1 *QVector3D, v2 *QVector3D) *QVector3D /*123*/ {
	var convArg0 = v1.GetCthis()
	var convArg1 = v2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3D12crossProductERKS_S1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}
func QVector3D_CrossProduct(v1 *QVector3D, v2 *QVector3D) *QVector3D /*123*/ {
	var nilthis *QVector3D
	rv := nilthis.CrossProduct(v1, v2)
	return rv
}

// /usr/include/qt/QtGui/qvector3d.h:103
// index:0
// Public static Visibility=Default Availability=Available
// [12] QVector3D normal(const QVector3D &, const QVector3D &)
func (this *QVector3D) Normal(v1 *QVector3D, v2 *QVector3D) *QVector3D /*123*/ {
	var convArg0 = v1.GetCthis()
	var convArg1 = v2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3D6normalERKS_S1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}
func QVector3D_Normal(v1 *QVector3D, v2 *QVector3D) *QVector3D /*123*/ {
	var nilthis *QVector3D
	rv := nilthis.Normal(v1, v2)
	return rv
}

// /usr/include/qt/QtGui/qvector3d.h:104
// index:1
// Public static Visibility=Default Availability=Available
// [12] QVector3D normal(const QVector3D &, const QVector3D &, const QVector3D &)
func (this *QVector3D) Normal_1(v1 *QVector3D, v2 *QVector3D, v3 *QVector3D) *QVector3D /*123*/ {
	var convArg0 = v1.GetCthis()
	var convArg1 = v2.GetCthis()
	var convArg2 = v3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3D6normalERKS_S1_S1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}
func QVector3D_Normal_1(v1 *QVector3D, v2 *QVector3D, v3 *QVector3D) *QVector3D /*123*/ {
	var nilthis *QVector3D
	rv := nilthis.Normal_1(v1, v2, v3)
	return rv
}

// /usr/include/qt/QtGui/qvector3d.h:107
// index:0
// Public Visibility=Default Availability=Available
// [12] QVector3D project(const QMatrix4x4 &, const QMatrix4x4 &, const QRect &)
func (this *QVector3D) Project(modelView *QMatrix4x4, projection *QMatrix4x4, viewport *qtcore.QRect) *QVector3D /*123*/ {
	var convArg0 = modelView.GetCthis()
	var convArg1 = projection.GetCthis()
	var convArg2 = viewport.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D7projectERK10QMatrix4x4S2_RK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}

// /usr/include/qt/QtGui/qvector3d.h:108
// index:0
// Public Visibility=Default Availability=Available
// [12] QVector3D unproject(const QMatrix4x4 &, const QMatrix4x4 &, const QRect &)
func (this *QVector3D) Unproject(modelView *QMatrix4x4, projection *QMatrix4x4, viewport *qtcore.QRect) *QVector3D /*123*/ {
	var convArg0 = modelView.GetCthis()
	var convArg1 = projection.GetCthis()
	var convArg2 = viewport.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D9unprojectERK10QMatrix4x4S2_RK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}

// /usr/include/qt/QtGui/qvector3d.h:110
// index:0
// Public Visibility=Default Availability=Available
// [4] float distanceToPoint(const QVector3D &)
func (this *QVector3D) DistanceToPoint(point *QVector3D) float32 {
	var convArg0 = point.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D15distanceToPointERKS_", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector3d.h:111
// index:0
// Public Visibility=Default Availability=Available
// [4] float distanceToPlane(const QVector3D &, const QVector3D &)
func (this *QVector3D) DistanceToPlane(plane *QVector3D, normal *QVector3D) float32 {
	var convArg0 = plane.GetCthis()
	var convArg1 = normal.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D15distanceToPlaneERKS_S1_", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector3d.h:112
// index:1
// Public Visibility=Default Availability=Available
// [4] float distanceToPlane(const QVector3D &, const QVector3D &, const QVector3D &)
func (this *QVector3D) DistanceToPlane_1(plane1 *QVector3D, plane2 *QVector3D, plane3 *QVector3D) float32 {
	var convArg0 = plane1.GetCthis()
	var convArg1 = plane2.GetCthis()
	var convArg2 = plane3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D15distanceToPlaneERKS_S1_S1_", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector3d.h:113
// index:0
// Public Visibility=Default Availability=Available
// [4] float distanceToLine(const QVector3D &, const QVector3D &)
func (this *QVector3D) DistanceToLine(point *QVector3D, direction *QVector3D) float32 {
	var convArg0 = point.GetCthis()
	var convArg1 = direction.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D14distanceToLineERKS_S1_", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector3d.h:129
// index:0
// Public Visibility=Default Availability=Available
// [8] QVector2D toVector2D()
func (this *QVector3D) ToVector2D() *QVector2D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D10toVector2DEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector2D)
	return rv2
}

// /usr/include/qt/QtGui/qvector3d.h:132
// index:0
// Public Visibility=Default Availability=Available
// [16] QVector4D toVector4D()
func (this *QVector3D) ToVector4D() *QVector4D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D10toVector4DEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector4DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector4D)
	return rv2
}

// /usr/include/qt/QtGui/qvector3d.h:135
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint toPoint()
func (this *QVector3D) ToPoint() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D7toPointEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qvector3d.h:136
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QPointF toPointF()
func (this *QVector3D) ToPointF() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D8toPointFEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

func DeleteQVector3D(this *QVector3D) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3DD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
