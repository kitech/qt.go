//  header block begin
// /usr/include/qt/QtGui/qbrush.h
// #include <qbrush.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 60
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
type QBrush struct {
	*qtrt.CObject
}

func (this *QBrush) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtGui/qbrush.h:66
// index:0
// void QBrush()
func NewQBrush() *QBrush {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrushC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(cthis)
	return gothis
}
func NewQBrushFromPointer(cthis unsafe.Pointer) *QBrush {
	return &QBrush{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qbrush.h:67
// index:1
// void QBrush(Qt::BrushStyle)
func NewQBrush_1(bs int) *QBrush {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrushC2EN2Qt10BrushStyleE", ffiqt.FFI_TYPE_VOID, cthis, &bs)
	gopp.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:68
// index:2
// void QBrush(const class QColor &, Qt::BrushStyle)
func NewQBrush_2(color unsafe.Pointer, bs int) *QBrush {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrushC2ERK6QColorN2Qt10BrushStyleE", ffiqt.FFI_TYPE_VOID, cthis, color, &bs)
	gopp.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:69
// index:3
// void QBrush(Qt::GlobalColor, Qt::BrushStyle)
func NewQBrush_3(color int, bs int) *QBrush {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrushC2EN2Qt11GlobalColorENS0_10BrushStyleE", ffiqt.FFI_TYPE_VOID, cthis, &color, &bs)
	gopp.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:71
// index:4
// void QBrush(const class QColor &, const class QPixmap &)
func NewQBrush_4(color unsafe.Pointer, pixmap unsafe.Pointer) *QBrush {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrushC2ERK6QColorRK7QPixmap", ffiqt.FFI_TYPE_VOID, cthis, color, pixmap)
	gopp.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:72
// index:5
// void QBrush(Qt::GlobalColor, const class QPixmap &)
func NewQBrush_5(color int, pixmap unsafe.Pointer) *QBrush {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrushC2EN2Qt11GlobalColorERK7QPixmap", ffiqt.FFI_TYPE_VOID, cthis, &color, pixmap)
	gopp.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:73
// index:6
// void QBrush(const class QPixmap &)
func NewQBrush_6(pixmap unsafe.Pointer) *QBrush {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrushC2ERK7QPixmap", ffiqt.FFI_TYPE_VOID, cthis, pixmap)
	gopp.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:74
// index:7
// void QBrush(const class QImage &)
func NewQBrush_7(image unsafe.Pointer) *QBrush {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrushC2ERK6QImage", ffiqt.FFI_TYPE_VOID, cthis, image)
	gopp.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:78
// index:8
// void QBrush(const class QGradient &)
func NewQBrush_8(gradient unsafe.Pointer) *QBrush {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrushC2ERK9QGradient", ffiqt.FFI_TYPE_VOID, cthis, gradient)
	gopp.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:80
// index:0
// void ~QBrush()
func DeleteQBrush(*QBrush) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrushD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:86
// index:0
// inline
// void swap(class QBrush &)
func (this *QBrush) Swap(other unsafe.Pointer) {
	// 0: (, other QBrush &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrush4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:91
// index:0
// inline
// Qt::BrushStyle style()
func (this *QBrush) Style() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QBrush5styleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:92
// index:0
// void setStyle(Qt::BrushStyle)
func (this *QBrush) SetStyle(arg0 int) {
	// 0: (, Qt::BrushStyle arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrush8setStyleEN2Qt10BrushStyleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:94
// index:0
// inline
// const QMatrix & matrix()
func (this *QBrush) Matrix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QBrush6matrixEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:95
// index:0
// void setMatrix(const class QMatrix &)
func (this *QBrush) SetMatrix(mat unsafe.Pointer) {
	// 0: (, mat const QMatrix &), (mat)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrush9setMatrixERK7QMatrix", ffiqt.FFI_TYPE_VOID, this.GetCthis(), mat)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:97
// index:0
// inline
// QTransform transform()
func (this *QBrush) Transform() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QBrush9transformEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:98
// index:0
// void setTransform(const class QTransform &)
func (this *QBrush) SetTransform(arg0 unsafe.Pointer) {
	// 0: (, const QTransform & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrush12setTransformERK10QTransform", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:100
// index:0
// QPixmap texture()
func (this *QBrush) Texture() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QBrush7textureEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:101
// index:0
// void setTexture(const class QPixmap &)
func (this *QBrush) SetTexture(pixmap unsafe.Pointer) {
	// 0: (, pixmap const QPixmap &), (pixmap)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrush10setTextureERK7QPixmap", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pixmap)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:103
// index:0
// QImage textureImage()
func (this *QBrush) TextureImage() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QBrush12textureImageEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:104
// index:0
// void setTextureImage(const class QImage &)
func (this *QBrush) SetTextureImage(image unsafe.Pointer) {
	// 0: (, image const QImage &), (image)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrush15setTextureImageERK6QImage", ffiqt.FFI_TYPE_VOID, this.GetCthis(), image)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:106
// index:0
// inline
// const QColor & color()
func (this *QBrush) Color() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QBrush5colorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:107
// index:0
// void setColor(const class QColor &)
func (this *QBrush) SetColor(color unsafe.Pointer) {
	// 0: (, color const QColor &), (color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrush8setColorERK6QColor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:108
// index:1
// inline
// void setColor(Qt::GlobalColor)
func (this *QBrush) SetColor_1(color int) {
	// 1: (, color Qt::GlobalColor), (&color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrush8setColorEN2Qt11GlobalColorE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:110
// index:0
// const QGradient * gradient()
func (this *QBrush) Gradient() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QBrush8gradientEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:112
// index:0
// bool isOpaque()
func (this *QBrush) IsOpaque() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QBrush8isOpaqueEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:129
// index:0
// inline
// bool isDetached()
func (this *QBrush) IsDetached() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QBrush10isDetachedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
