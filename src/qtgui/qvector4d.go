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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qvector4d.h:59
// index:0
// inline
// void QVector4D()
func NewQVector4D() *QVector4D {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4DC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QVector4D{cthis}
}

// /usr/include/qt/QtGui/qvector4d.h:60
// index:1
// inline
// void QVector4D(Qt::Initialization)
func NewQVector4D_1(arg0 int) *QVector4D {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4DC2EN2Qt14InitializationE", ffiqt.FFI_TYPE_VOID, cthis, &arg0)
	gopp.ErrPrint(err, rv)
	return &QVector4D{cthis}
}

// /usr/include/qt/QtGui/qvector4d.h:61
// index:2
// inline
// void QVector4D(float, float, float, float)
func NewQVector4D_2(xpos float32, ypos float32, zpos float32, wpos float32) *QVector4D {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4DC2Effff", ffiqt.FFI_TYPE_VOID, cthis, &xpos, &ypos, &zpos, &wpos)
	gopp.ErrPrint(err, rv)
	return &QVector4D{cthis}
}

// /usr/include/qt/QtGui/qvector4d.h:62
// index:3
// inline
// void QVector4D(const class QPoint &)
func NewQVector4D_3(point unsafe.Pointer) *QVector4D {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4DC2ERK6QPoint", ffiqt.FFI_TYPE_VOID, cthis, point)
	gopp.ErrPrint(err, rv)
	return &QVector4D{cthis}
}

// /usr/include/qt/QtGui/qvector4d.h:63
// index:4
// inline
// void QVector4D(const class QPointF &)
func NewQVector4D_4(point unsafe.Pointer) *QVector4D {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4DC2ERK7QPointF", ffiqt.FFI_TYPE_VOID, cthis, point)
	gopp.ErrPrint(err, rv)
	return &QVector4D{cthis}
}

// /usr/include/qt/QtGui/qvector4d.h:65
// index:5
// void QVector4D(const class QVector2D &)
func NewQVector4D_5(vector unsafe.Pointer) *QVector4D {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4DC2ERK9QVector2D", ffiqt.FFI_TYPE_VOID, cthis, vector)
	gopp.ErrPrint(err, rv)
	return &QVector4D{cthis}
}

// /usr/include/qt/QtGui/qvector4d.h:66
// index:6
// void QVector4D(const class QVector2D &, float, float)
func NewQVector4D_6(vector unsafe.Pointer, zpos float32, wpos float32) *QVector4D {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4DC2ERK9QVector2Dff", ffiqt.FFI_TYPE_VOID, cthis, vector, &zpos, &wpos)
	gopp.ErrPrint(err, rv)
	return &QVector4D{cthis}
}

// /usr/include/qt/QtGui/qvector4d.h:69
// index:7
// void QVector4D(const class QVector3D &)
func NewQVector4D_7(vector unsafe.Pointer) *QVector4D {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4DC2ERK9QVector3D", ffiqt.FFI_TYPE_VOID, cthis, vector)
	gopp.ErrPrint(err, rv)
	return &QVector4D{cthis}
}

// /usr/include/qt/QtGui/qvector4d.h:70
// index:8
// void QVector4D(const class QVector3D &, float)
func NewQVector4D_8(vector unsafe.Pointer, wpos float32) *QVector4D {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4DC2ERK9QVector3Df", ffiqt.FFI_TYPE_VOID, cthis, vector, &wpos)
	gopp.ErrPrint(err, rv)
	return &QVector4D{cthis}
}

// /usr/include/qt/QtGui/qvector4d.h:73
// index:0
// bool isNull()
func (this *QVector4D) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D6isNullEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:75
// index:0
// inline
// float x()
func (this *QVector4D) X() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D1xEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:76
// index:0
// inline
// float y()
func (this *QVector4D) Y() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D1yEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:77
// index:0
// inline
// float z()
func (this *QVector4D) Z() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D1zEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:78
// index:0
// inline
// float w()
func (this *QVector4D) W() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D1wEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:80
// index:0
// void setX(float)
func (this *QVector4D) SetX(x float32) {
	// 0: (, float x), (&x)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4D4setXEf", ffiqt.FFI_TYPE_VOID, this.cthis, &x)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:81
// index:0
// void setY(float)
func (this *QVector4D) SetY(y float32) {
	// 0: (, float y), (&y)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4D4setYEf", ffiqt.FFI_TYPE_VOID, this.cthis, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:82
// index:0
// void setZ(float)
func (this *QVector4D) SetZ(z float32) {
	// 0: (, float z), (&z)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4D4setZEf", ffiqt.FFI_TYPE_VOID, this.cthis, &z)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:83
// index:0
// void setW(float)
func (this *QVector4D) SetW(w float32) {
	// 0: (, float w), (&w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4D4setWEf", ffiqt.FFI_TYPE_VOID, this.cthis, &w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:88
// index:0
// float length()
func (this *QVector4D) Length() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D6lengthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:89
// index:0
// float lengthSquared()
func (this *QVector4D) LengthSquared() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D13lengthSquaredEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:91
// index:0
// QVector4D normalized()
func (this *QVector4D) Normalized() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D10normalizedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:92
// index:0
// void normalize()
func (this *QVector4D) Normalize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4D9normalizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:101
// index:0
// static
// float dotProduct(const class QVector4D &, const class QVector4D &)
func (this *QVector4D) DotProduct(v1 unsafe.Pointer, v2 unsafe.Pointer) {
	// 0: (const QVector4D & v1, const QVector4D & v2), (v1, v2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector4D10dotProductERKS_S1_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QVector4D_DotProduct(v1 unsafe.Pointer, v2 unsafe.Pointer) {
	// 0: (const QVector4D & v1, const QVector4D & v2), (v1, v2)
	var nilthis *QVector4D
	nilthis.DotProduct(v1, v2)
}

// /usr/include/qt/QtGui/qvector4d.h:117
// index:0
// QVector2D toVector2D()
func (this *QVector4D) ToVector2D() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D10toVector2DEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:118
// index:0
// QVector2D toVector2DAffine()
func (this *QVector4D) ToVector2DAffine() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D16toVector2DAffineEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:121
// index:0
// QVector3D toVector3D()
func (this *QVector4D) ToVector3D() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D10toVector3DEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:122
// index:0
// QVector3D toVector3DAffine()
func (this *QVector4D) ToVector3DAffine() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D16toVector3DAffineEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:125
// index:0
// inline
// QPoint toPoint()
func (this *QVector4D) ToPoint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D7toPointEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector4d.h:126
// index:0
// inline
// QPointF toPointF()
func (this *QVector4D) ToPointF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector4D8toPointFEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
