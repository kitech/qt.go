//  header block begin
// /usr/include/qt/QtGui/qbrush.h
// #include <qbrush.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 51
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
func NewQBrushFromPointer(cthis unsafe.Pointer) *QBrush {
	return &QBrush{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qbrush.h:66
// index:0
// Public
// void QBrush()
func NewQBrush() *QBrush {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrushC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:67
// index:1
// Public
// void QBrush(Qt::BrushStyle)
func NewQBrush_1(bs int) *QBrush {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrushC2EN2Qt10BrushStyleE", ffiqt.FFI_TYPE_VOID, cthis, &bs)
	gopp.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:68
// index:2
// Public
// void QBrush(const class QColor &, Qt::BrushStyle)
func NewQBrush_2(color *QColor, bs int) *QBrush {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrushC2ERK6QColorN2Qt10BrushStyleE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &bs)
	gopp.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:69
// index:3
// Public
// void QBrush(Qt::GlobalColor, Qt::BrushStyle)
func NewQBrush_3(color int, bs int) *QBrush {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrushC2EN2Qt11GlobalColorENS0_10BrushStyleE", ffiqt.FFI_TYPE_VOID, cthis, &color, &bs)
	gopp.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:71
// index:4
// Public
// void QBrush(const class QColor &, const class QPixmap &)
func NewQBrush_4(color *QColor, pixmap *QPixmap) *QBrush {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = color.GetCthis()
	var convArg1 = pixmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrushC2ERK6QColorRK7QPixmap", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:72
// index:5
// Public
// void QBrush(Qt::GlobalColor, const class QPixmap &)
func NewQBrush_5(color int, pixmap *QPixmap) *QBrush {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg1 = pixmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrushC2EN2Qt11GlobalColorERK7QPixmap", ffiqt.FFI_TYPE_VOID, cthis, &color, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:73
// index:6
// Public
// void QBrush(const class QPixmap &)
func NewQBrush_6(pixmap *QPixmap) *QBrush {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = pixmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrushC2ERK7QPixmap", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:74
// index:7
// Public
// void QBrush(const class QImage &)
func NewQBrush_7(image *QImage) *QBrush {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = image.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrushC2ERK6QImage", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:78
// index:8
// Public
// void QBrush(const class QGradient &)
func NewQBrush_8(gradient *QGradient) *QBrush {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = gradient.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrushC2ERK9QGradient", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:80
// index:0
// Public
// void ~QBrush()
func DeleteQBrush(*QBrush) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrushD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:86
// index:0
// Public inline
// void swap(class QBrush &)
func (this *QBrush) Swap(other *QBrush) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrush4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:91
// index:0
// Public inline
// Qt::BrushStyle style()
func (this *QBrush) Style() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QBrush5styleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qbrush.h:92
// index:0
// Public
// void setStyle(Qt::BrushStyle)
func (this *QBrush) SetStyle(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrush8setStyleEN2Qt10BrushStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:94
// index:0
// Public inline
// const QMatrix & matrix()
func (this *QBrush) Matrix() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QBrush6matrixEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qbrush.h:95
// index:0
// Public
// void setMatrix(const class QMatrix &)
func (this *QBrush) SetMatrix(mat *QMatrix) {
	var convArg0 = mat.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrush9setMatrixERK7QMatrix", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:97
// index:0
// Public inline
// QTransform transform()
func (this *QBrush) Transform() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QBrush9transformEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qbrush.h:98
// index:0
// Public
// void setTransform(const class QTransform &)
func (this *QBrush) SetTransform(arg0 *QTransform) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrush12setTransformERK10QTransform", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:100
// index:0
// Public
// QPixmap texture()
func (this *QBrush) Texture() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QBrush7textureEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qbrush.h:101
// index:0
// Public
// void setTexture(const class QPixmap &)
func (this *QBrush) SetTexture(pixmap *QPixmap) {
	var convArg0 = pixmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrush10setTextureERK7QPixmap", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:103
// index:0
// Public
// QImage textureImage()
func (this *QBrush) TextureImage() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QBrush12textureImageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qbrush.h:104
// index:0
// Public
// void setTextureImage(const class QImage &)
func (this *QBrush) SetTextureImage(image *QImage) {
	var convArg0 = image.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrush15setTextureImageERK6QImage", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:106
// index:0
// Public inline
// const QColor & color()
func (this *QBrush) Color() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QBrush5colorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qbrush.h:107
// index:0
// Public
// void setColor(const class QColor &)
func (this *QBrush) SetColor(color *QColor) {
	var convArg0 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrush8setColorERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:108
// index:1
// Public inline
// void setColor(Qt::GlobalColor)
func (this *QBrush) SetColor_1(color int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QBrush8setColorEN2Qt11GlobalColorE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:110
// index:0
// Public
// const QGradient * gradient()
func (this *QBrush) Gradient() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QBrush8gradientEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qbrush.h:112
// index:0
// Public
// bool isOpaque()
func (this *QBrush) IsOpaque() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QBrush8isOpaqueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qbrush.h:129
// index:0
// Public inline
// bool isDetached()
func (this *QBrush) IsDetached() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QBrush10isDetachedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
