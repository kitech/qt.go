package qtgui

// /usr/include/qt/QtGui/qvector4d.h
// #include <qvector4d.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 33
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QVector4D struct {
	*qtrt.CObject
}
type QVector4D_ITF interface {
	QVector4D_PTR() *QVector4D
}

func (ptr *QVector4D) QVector4D_PTR() *QVector4D { return ptr }

func (this *QVector4D) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QVector4D) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQVector4DFromPointer(cthis unsafe.Pointer) *QVector4D {
	return &QVector4D{&qtrt.CObject{cthis}}
}
func (*QVector4D) NewFromPointer(cthis unsafe.Pointer) *QVector4D {
	return NewQVector4DFromPointer(cthis)
}

// /usr/include/qt/QtGui/qvector4d.h:59
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QVector4D()
func NewQVector4D() *QVector4D {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector4DC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector4DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector4D)
	return gothis
}

// /usr/include/qt/QtGui/qvector4d.h:60
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QVector4D(Qt::Initialization)
func NewQVector4D_1(arg0 int) *QVector4D {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector4DC2EN2Qt14InitializationE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector4DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector4D)
	return gothis
}

// /usr/include/qt/QtGui/qvector4d.h:61
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QVector4D(float, float, float, float)
func NewQVector4D_2(xpos float32, ypos float32, zpos float32, wpos float32) *QVector4D {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector4DC2Effff", qtrt.FFI_TYPE_POINTER, xpos, ypos, zpos, wpos)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector4DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector4D)
	return gothis
}

// /usr/include/qt/QtGui/qvector4d.h:62
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void QVector4D(const QPoint &)
func NewQVector4D_3(point *qtcore.QPoint) *QVector4D {
	var convArg0 = point.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector4DC2ERK6QPoint", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector4DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector4D)
	return gothis
}

// /usr/include/qt/QtGui/qvector4d.h:63
// index:4
// Public inline Visibility=Default Availability=Available
// [-2] void QVector4D(const QPointF &)
func NewQVector4D_4(point *qtcore.QPointF) *QVector4D {
	var convArg0 = point.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector4DC2ERK7QPointF", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector4DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector4D)
	return gothis
}

// /usr/include/qt/QtGui/qvector4d.h:65
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QVector4D(const QVector2D &)
func NewQVector4D_5(vector *QVector2D) *QVector4D {
	var convArg0 = vector.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector4DC2ERK9QVector2D", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector4DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector4D)
	return gothis
}

// /usr/include/qt/QtGui/qvector4d.h:66
// index:6
// Public Visibility=Default Availability=Available
// [-2] void QVector4D(const QVector2D &, float, float)
func NewQVector4D_6(vector *QVector2D, zpos float32, wpos float32) *QVector4D {
	var convArg0 = vector.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector4DC2ERK9QVector2Dff", qtrt.FFI_TYPE_POINTER, convArg0, zpos, wpos)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector4DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector4D)
	return gothis
}

// /usr/include/qt/QtGui/qvector4d.h:69
// index:7
// Public Visibility=Default Availability=Available
// [-2] void QVector4D(const QVector3D &)
func NewQVector4D_7(vector *QVector3D) *QVector4D {
	var convArg0 = vector.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector4DC2ERK9QVector3D", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector4DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector4D)
	return gothis
}

// /usr/include/qt/QtGui/qvector4d.h:70
// index:8
// Public Visibility=Default Availability=Available
// [-2] void QVector4D(const QVector3D &, float)
func NewQVector4D_8(vector *QVector3D, wpos float32) *QVector4D {
	var convArg0 = vector.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector4DC2ERK9QVector3Df", qtrt.FFI_TYPE_POINTER, convArg0, wpos)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector4DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector4D)
	return gothis
}

// /usr/include/qt/QtGui/qvector4d.h:73
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QVector4D) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector4D6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qvector4d.h:75
// index:0
// Public inline Visibility=Default Availability=Available
// [4] float x()
func (this *QVector4D) X() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector4D1xEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector4d.h:76
// index:0
// Public inline Visibility=Default Availability=Available
// [4] float y()
func (this *QVector4D) Y() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector4D1yEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector4d.h:77
// index:0
// Public inline Visibility=Default Availability=Available
// [4] float z()
func (this *QVector4D) Z() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector4D1zEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector4d.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [4] float w()
func (this *QVector4D) W() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector4D1wEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector4d.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setX(float)
func (this *QVector4D) SetX(x float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector4D4setXEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setY(float)
func (this *QVector4D) SetY(y float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector4D4setYEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setZ(float)
func (this *QVector4D) SetZ(z float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector4D4setZEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), z)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setW(float)
func (this *QVector4D) SetW(w float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector4D4setWEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:88
// index:0
// Public Visibility=Default Availability=Available
// [4] float length()
func (this *QVector4D) Length() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector4D6lengthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector4d.h:89
// index:0
// Public Visibility=Default Availability=Available
// [4] float lengthSquared()
func (this *QVector4D) LengthSquared() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector4D13lengthSquaredEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector4d.h:91
// index:0
// Public Visibility=Default Availability=Available
// [16] QVector4D normalized()
func (this *QVector4D) Normalized() *QVector4D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector4D10normalizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector4DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector4D)
	return rv2
}

// /usr/include/qt/QtGui/qvector4d.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void normalize()
func (this *QVector4D) Normalize() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector4D9normalizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:101
// index:0
// Public static Visibility=Default Availability=Available
// [4] float dotProduct(const QVector4D &, const QVector4D &)
func (this *QVector4D) DotProduct(v1 *QVector4D, v2 *QVector4D) float32 {
	var convArg0 = v1.GetCthis()
	var convArg1 = v2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector4D10dotProductERKS_S1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}
func QVector4D_DotProduct(v1 *QVector4D, v2 *QVector4D) float32 {
	var nilthis *QVector4D
	rv := nilthis.DotProduct(v1, v2)
	return rv
}

// /usr/include/qt/QtGui/qvector4d.h:117
// index:0
// Public Visibility=Default Availability=Available
// [8] QVector2D toVector2D()
func (this *QVector4D) ToVector2D() *QVector2D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector4D10toVector2DEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector2D)
	return rv2
}

// /usr/include/qt/QtGui/qvector4d.h:118
// index:0
// Public Visibility=Default Availability=Available
// [8] QVector2D toVector2DAffine()
func (this *QVector4D) ToVector2DAffine() *QVector2D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector4D16toVector2DAffineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector2D)
	return rv2
}

// /usr/include/qt/QtGui/qvector4d.h:121
// index:0
// Public Visibility=Default Availability=Available
// [12] QVector3D toVector3D()
func (this *QVector4D) ToVector3D() *QVector3D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector4D10toVector3DEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}

// /usr/include/qt/QtGui/qvector4d.h:122
// index:0
// Public Visibility=Default Availability=Available
// [12] QVector3D toVector3DAffine()
func (this *QVector4D) ToVector3DAffine() *QVector3D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector4D16toVector3DAffineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}

// /usr/include/qt/QtGui/qvector4d.h:125
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint toPoint()
func (this *QVector4D) ToPoint() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector4D7toPointEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qvector4d.h:126
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QPointF toPointF()
func (this *QVector4D) ToPointF() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector4D8toPointFEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

func DeleteQVector4D(this *QVector4D) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector4DD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
