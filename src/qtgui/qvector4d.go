//  header block begin
// /usr/include/qt/QtGui/qvector4d.h
// #include <qvector4d.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 33
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
type QVector4D struct {
	*qtrt.CObject
}

func (this *QVector4D) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQVector4DFromPointer(cthis unsafe.Pointer) *QVector4D {
	return &QVector4D{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qvector4d.h:59
// index:0
// Public inline
// void QVector4D()
func NewQVector4D() *QVector4D {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4DC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector4DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector4d.h:60
// index:1
// Public inline
// void QVector4D(Qt::Initialization)
func NewQVector4D_1(arg0 int) *QVector4D {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4DC2EN2Qt14InitializationE", ffiqt.FFI_TYPE_VOID, cthis, &arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector4DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector4d.h:61
// index:2
// Public inline
// void QVector4D(float, float, float, float)
func NewQVector4D_2(xpos float32, ypos float32, zpos float32, wpos float32) *QVector4D {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4DC2Effff", ffiqt.FFI_TYPE_VOID, cthis, &xpos, &ypos, &zpos, &wpos)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector4DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector4d.h:62
// index:3
// Public inline
// void QVector4D(const class QPoint &)
func NewQVector4D_3(point *qtcore.QPoint) *QVector4D {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4DC2ERK6QPoint", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector4DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector4d.h:63
// index:4
// Public inline
// void QVector4D(const class QPointF &)
func NewQVector4D_4(point *qtcore.QPointF) *QVector4D {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4DC2ERK7QPointF", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector4DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector4d.h:65
// index:5
// Public
// void QVector4D(const class QVector2D &)
func NewQVector4D_5(vector *QVector2D) *QVector4D {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = vector.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4DC2ERK9QVector2D", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector4DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector4d.h:66
// index:6
// Public
// void QVector4D(const class QVector2D &, float, float)
func NewQVector4D_6(vector *QVector2D, zpos float32, wpos float32) *QVector4D {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = vector.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4DC2ERK9QVector2Dff", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &zpos, &wpos)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector4DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector4d.h:69
// index:7
// Public
// void QVector4D(const class QVector3D &)
func NewQVector4D_7(vector *QVector3D) *QVector4D {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = vector.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4DC2ERK9QVector3D", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector4DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector4d.h:70
// index:8
// Public
// void QVector4D(const class QVector3D &, float)
func NewQVector4D_8(vector *QVector3D, wpos float32) *QVector4D {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = vector.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4DC2ERK9QVector3Df", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &wpos)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector4DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector4d.h:73
// index:0
// Public
// bool isNull()
func (this *QVector4D) IsNull() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qvector4d.h:75
// index:0
// Public inline
// float x()
func (this *QVector4D) X() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D1xEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qvector4d.h:76
// index:0
// Public inline
// float y()
func (this *QVector4D) Y() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D1yEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qvector4d.h:77
// index:0
// Public inline
// float z()
func (this *QVector4D) Z() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D1zEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qvector4d.h:78
// index:0
// Public inline
// float w()
func (this *QVector4D) W() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D1wEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qvector4d.h:80
// index:0
// Public
// void setX(float)
func (this *QVector4D) SetX(x float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4D4setXEf", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:81
// index:0
// Public
// void setY(float)
func (this *QVector4D) SetY(y float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4D4setYEf", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:82
// index:0
// Public
// void setZ(float)
func (this *QVector4D) SetZ(z float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4D4setZEf", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &z)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:83
// index:0
// Public
// void setW(float)
func (this *QVector4D) SetW(w float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4D4setWEf", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:88
// index:0
// Public
// float length()
func (this *QVector4D) Length() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D6lengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qvector4d.h:89
// index:0
// Public
// float lengthSquared()
func (this *QVector4D) LengthSquared() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D13lengthSquaredEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qvector4d.h:91
// index:0
// Public
// QVector4D normalized()
func (this *QVector4D) Normalized() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D10normalizedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qvector4d.h:92
// index:0
// Public
// void normalize()
func (this *QVector4D) Normalize() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4D9normalizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:101
// index:0
// Public static
// float dotProduct(const class QVector4D &, const class QVector4D &)
func (this *QVector4D) DotProduct(v1 *QVector4D, v2 *QVector4D) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4D10dotProductERKS_S1_", ffiqt.FFI_TYPE_POINTER, v1, v2)
	gopp.ErrPrint(err, rv)
	return rv
}
func QVector4D_DotProduct(v1 *QVector4D, v2 *QVector4D) {
	var nilthis *QVector4D
	nilthis.DotProduct(v1, v2)
}

// /usr/include/qt/QtGui/qvector4d.h:117
// index:0
// Public
// QVector2D toVector2D()
func (this *QVector4D) ToVector2D() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D10toVector2DEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qvector4d.h:118
// index:0
// Public
// QVector2D toVector2DAffine()
func (this *QVector4D) ToVector2DAffine() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D16toVector2DAffineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qvector4d.h:121
// index:0
// Public
// QVector3D toVector3D()
func (this *QVector4D) ToVector3D() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D10toVector3DEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qvector4d.h:122
// index:0
// Public
// QVector3D toVector3DAffine()
func (this *QVector4D) ToVector3DAffine() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D16toVector3DAffineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qvector4d.h:125
// index:0
// Public inline
// QPoint toPoint()
func (this *QVector4D) ToPoint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D7toPointEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qvector4d.h:126
// index:0
// Public inline
// QPointF toPointF()
func (this *QVector4D) ToPointF() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D8toPointFEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
