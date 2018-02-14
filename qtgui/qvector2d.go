package qtgui

// /usr/include/qt/QtGui/qvector2d.h
// #include <qvector2d.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QVector2D struct {
	*qtrt.CObject
}
type QVector2D_ITF interface {
	QVector2D_PTR() *QVector2D
}

func (ptr *QVector2D) QVector2D_PTR() *QVector2D { return ptr }

func (this *QVector2D) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QVector2D) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQVector2DFromPointer(cthis unsafe.Pointer) *QVector2D {
	return &QVector2D{&qtrt.CObject{cthis}}
}
func (*QVector2D) NewFromPointer(cthis unsafe.Pointer) *QVector2D {
	return NewQVector2DFromPointer(cthis)
}

// /usr/include/qt/QtGui/qvector2d.h:59
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QVector2D()
func NewQVector2D() *QVector2D {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2DC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector2D)
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:60
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QVector2D(Qt::Initialization)
func NewQVector2D_1(arg0 int) *QVector2D {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2DC2EN2Qt14InitializationE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector2D)
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:61
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QVector2D(float, float)
func NewQVector2D_2(xpos float32, ypos float32) *QVector2D {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2DC2Eff", qtrt.FFI_TYPE_POINTER, xpos, ypos)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector2D)
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:62
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void QVector2D(const QPoint &)
func NewQVector2D_3(point qtcore.QPoint_ITF) *QVector2D {
	var convArg0 = point.QPoint_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2DC2ERK6QPoint", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector2D)
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:63
// index:4
// Public inline Visibility=Default Availability=Available
// [-2] void QVector2D(const QPointF &)
func NewQVector2D_4(point qtcore.QPointF_ITF) *QVector2D {
	var convArg0 = point.QPointF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2DC2ERK7QPointF", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector2D)
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:65
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QVector2D(const QVector3D &)
func NewQVector2D_5(vector QVector3D_ITF) *QVector2D {
	var convArg0 = vector.QVector3D_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2DC2ERK9QVector3D", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector2D)
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:68
// index:6
// Public Visibility=Default Availability=Available
// [-2] void QVector2D(const QVector4D &)
func NewQVector2D_6(vector QVector4D_ITF) *QVector2D {
	var convArg0 = vector.QVector4D_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2DC2ERK9QVector4D", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector2D)
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:71
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QVector2D) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector2D6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qvector2d.h:73
// index:0
// Public inline Visibility=Default Availability=Available
// [4] float x()
func (this *QVector2D) X() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector2D1xEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector2d.h:74
// index:0
// Public inline Visibility=Default Availability=Available
// [4] float y()
func (this *QVector2D) Y() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector2D1yEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector2d.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setX(float)
func (this *QVector2D) SetX(x float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2D4setXEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector2d.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setY(float)
func (this *QVector2D) SetY(y float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2D4setYEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector2d.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] float length()
func (this *QVector2D) Length() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector2D6lengthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector2d.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] float lengthSquared()
func (this *QVector2D) LengthSquared() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector2D13lengthSquaredEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector2d.h:85
// index:0
// Public Visibility=Default Availability=Available
// [8] QVector2D normalized()
func (this *QVector2D) Normalized() *QVector2D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector2D10normalizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector2D)
	return rv2
}

// /usr/include/qt/QtGui/qvector2d.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void normalize()
func (this *QVector2D) Normalize() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2D9normalizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector2d.h:88
// index:0
// Public Visibility=Default Availability=Available
// [4] float distanceToPoint(const QVector2D &)
func (this *QVector2D) DistanceToPoint(point QVector2D_ITF) float32 {
	var convArg0 = point.QVector2D_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector2D15distanceToPointERKS_", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector2d.h:89
// index:0
// Public Visibility=Default Availability=Available
// [4] float distanceToLine(const QVector2D &, const QVector2D &)
func (this *QVector2D) DistanceToLine(point QVector2D_ITF, direction QVector2D_ITF) float32 {
	var convArg0 = point.QVector2D_PTR().GetCthis()
	var convArg1 = direction.QVector2D_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector2D14distanceToLineERKS_S1_", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector2d.h:98
// index:0
// Public static Visibility=Default Availability=Available
// [4] float dotProduct(const QVector2D &, const QVector2D &)
func (this *QVector2D) DotProduct(v1 QVector2D_ITF, v2 QVector2D_ITF) float32 {
	var convArg0 = v1.QVector2D_PTR().GetCthis()
	var convArg1 = v2.QVector2D_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2D10dotProductERKS_S1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}
func QVector2D_DotProduct(v1 QVector2D_ITF, v2 QVector2D_ITF) float32 {
	var nilthis *QVector2D
	rv := nilthis.DotProduct(v1, v2)
	return rv
}

// /usr/include/qt/QtGui/qvector2d.h:114
// index:0
// Public Visibility=Default Availability=Available
// [12] QVector3D toVector3D()
func (this *QVector2D) ToVector3D() *QVector3D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector2D10toVector3DEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}

// /usr/include/qt/QtGui/qvector2d.h:117
// index:0
// Public Visibility=Default Availability=Available
// [16] QVector4D toVector4D()
func (this *QVector2D) ToVector4D() *QVector4D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector2D10toVector4DEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector4DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector4D)
	return rv2
}

// /usr/include/qt/QtGui/qvector2d.h:120
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint toPoint()
func (this *QVector2D) ToPoint() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector2D7toPointEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qvector2d.h:121
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QPointF toPointF()
func (this *QVector2D) ToPointF() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector2D8toPointFEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

func DeleteQVector2D(this *QVector2D) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2DD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
