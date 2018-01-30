package qtgui

// /usr/include/qt/QtGui/qvector2d.h
// #include <qvector2d.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
type QVector2D struct {
	*qtrt.CObject
}

func (this *QVector2D) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QVector2D) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector2DC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:60
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QVector2D(Qt::Initialization)
func NewQVector2D_1(arg0 int) *QVector2D {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector2DC2EN2Qt14InitializationE", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:61
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QVector2D(float, float)
func NewQVector2D_2(xpos float32, ypos float32) *QVector2D {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector2DC2Eff", ffiqt.FFI_TYPE_POINTER, xpos, ypos)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:62
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void QVector2D(const QPoint &)
func NewQVector2D_3(point *qtcore.QPoint) *QVector2D {
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector2DC2ERK6QPoint", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:63
// index:4
// Public inline Visibility=Default Availability=Available
// [-2] void QVector2D(const QPointF &)
func NewQVector2D_4(point *qtcore.QPointF) *QVector2D {
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector2DC2ERK7QPointF", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:65
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QVector2D(const QVector3D &)
func NewQVector2D_5(vector *QVector3D) *QVector2D {
	var convArg0 = vector.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector2DC2ERK9QVector3D", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:68
// index:6
// Public Visibility=Default Availability=Available
// [-2] void QVector2D(const QVector4D &)
func NewQVector2D_6(vector *QVector4D) *QVector2D {
	var convArg0 = vector.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector2DC2ERK9QVector4D", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:71
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QVector2D) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector2D6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qvector2d.h:73
// index:0
// Public inline Visibility=Default Availability=Available
// [4] float x()
func (this *QVector2D) X() float32 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector2D1xEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector2d.h:74
// index:0
// Public inline Visibility=Default Availability=Available
// [4] float y()
func (this *QVector2D) Y() float32 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector2D1yEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector2d.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setX(float)
func (this *QVector2D) SetX(x float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector2D4setXEf", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector2d.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setY(float)
func (this *QVector2D) SetY(y float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector2D4setYEf", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector2d.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] float length()
func (this *QVector2D) Length() float32 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector2D6lengthEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector2d.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] float lengthSquared()
func (this *QVector2D) LengthSquared() float32 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector2D13lengthSquaredEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector2d.h:85
// index:0
// Public Visibility=Default Availability=Available
// [8] QVector2D normalized()
func (this *QVector2D) Normalized() *QVector2D /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector2D10normalizedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qvector2d.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void normalize()
func (this *QVector2D) Normalize() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector2D9normalizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector2d.h:88
// index:0
// Public Visibility=Default Availability=Available
// [4] float distanceToPoint(const QVector2D &)
func (this *QVector2D) DistanceToPoint(point *QVector2D) float32 {
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector2D15distanceToPointERKS_", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector2d.h:89
// index:0
// Public Visibility=Default Availability=Available
// [4] float distanceToLine(const QVector2D &, const QVector2D &)
func (this *QVector2D) DistanceToLine(point *QVector2D, direction *QVector2D) float32 {
	var convArg0 = point.GetCthis()
	var convArg1 = direction.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector2D14distanceToLineERKS_S1_", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector2d.h:98
// index:0
// Public static Visibility=Default Availability=Available
// [4] float dotProduct(const QVector2D &, const QVector2D &)
func (this *QVector2D) DotProduct(v1 *QVector2D, v2 *QVector2D) float32 {
	var convArg0 = v1.GetCthis()
	var convArg1 = v2.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector2D10dotProductERKS_S1_", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}
func QVector2D_DotProduct(v1 *QVector2D, v2 *QVector2D) float32 {
	var nilthis *QVector2D
	rv := nilthis.DotProduct(v1, v2)
	return rv
}

// /usr/include/qt/QtGui/qvector2d.h:114
// index:0
// Public Visibility=Default Availability=Available
// [12] QVector3D toVector3D()
func (this *QVector2D) ToVector3D() *QVector3D /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector2D10toVector3DEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qvector2d.h:117
// index:0
// Public Visibility=Default Availability=Available
// [16] QVector4D toVector4D()
func (this *QVector2D) ToVector4D() *QVector4D /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector2D10toVector4DEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQVector4DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qvector2d.h:120
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint toPoint()
func (this *QVector2D) ToPoint() *qtcore.QPoint /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector2D7toPointEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qvector2d.h:121
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QPointF toPointF()
func (this *QVector2D) ToPointF() *qtcore.QPointF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector2D8toPointFEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
