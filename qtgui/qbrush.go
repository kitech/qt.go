package qtgui

// /usr/include/qt/QtGui/qbrush.h
// #include <qbrush.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 58
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QBrush struct {
	*qtrt.CObject
}

func (this *QBrush) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QBrush) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQBrushFromPointer(cthis unsafe.Pointer) *QBrush {
	return &QBrush{&qtrt.CObject{cthis}}
}
func (*QBrush) NewFromPointer(cthis unsafe.Pointer) *QBrush {
	return NewQBrushFromPointer(cthis)
}

// /usr/include/qt/QtGui/qbrush.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QBrush()
func NewQBrush() *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrushC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQBrush)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:67
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QBrush(Qt::BrushStyle)
func NewQBrush_1(bs int) *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrushC2EN2Qt10BrushStyleE", qtrt.FFI_TYPE_POINTER, bs)
	qtrt.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQBrush)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:68
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QBrush(const QColor &, Qt::BrushStyle)
func NewQBrush_2(color *QColor, bs int) *QBrush {
	var convArg0 = color.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrushC2ERK6QColorN2Qt10BrushStyleE", qtrt.FFI_TYPE_POINTER, convArg0, bs)
	qtrt.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQBrush)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:69
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QBrush(Qt::GlobalColor, Qt::BrushStyle)
func NewQBrush_3(color int, bs int) *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrushC2EN2Qt11GlobalColorENS0_10BrushStyleE", qtrt.FFI_TYPE_POINTER, color, bs)
	qtrt.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQBrush)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:71
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QBrush(const QColor &, const QPixmap &)
func NewQBrush_4(color *QColor, pixmap *QPixmap) *QBrush {
	var convArg0 = color.GetCthis()
	var convArg1 = pixmap.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrushC2ERK6QColorRK7QPixmap", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQBrush)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:72
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QBrush(Qt::GlobalColor, const QPixmap &)
func NewQBrush_5(color int, pixmap *QPixmap) *QBrush {
	var convArg1 = pixmap.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrushC2EN2Qt11GlobalColorERK7QPixmap", qtrt.FFI_TYPE_POINTER, color, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQBrush)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:73
// index:6
// Public Visibility=Default Availability=Available
// [-2] void QBrush(const QPixmap &)
func NewQBrush_6(pixmap *QPixmap) *QBrush {
	var convArg0 = pixmap.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrushC2ERK7QPixmap", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQBrush)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:74
// index:7
// Public Visibility=Default Availability=Available
// [-2] void QBrush(const QImage &)
func NewQBrush_7(image *QImage) *QBrush {
	var convArg0 = image.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrushC2ERK6QImage", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQBrush)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:78
// index:8
// Public Visibility=Default Availability=Available
// [-2] void QBrush(const QGradient &)
func NewQBrush_8(gradient *QGradient) *QBrush {
	var convArg0 = gradient.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrushC2ERK9QGradient", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQBrush)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QBrush()
func DeleteQBrush(this *QBrush) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrushD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qbrush.h:86
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QBrush &)
func (this *QBrush) Swap(other *QBrush) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrush4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:91
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::BrushStyle style()
func (this *QBrush) Style() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QBrush5styleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qbrush.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStyle(Qt::BrushStyle)
func (this *QBrush) SetStyle(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrush8setStyleEN2Qt10BrushStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:94
// index:0
// Public inline Visibility=Default Availability=Available
// [48] const QMatrix & matrix()
func (this *QBrush) Matrix() *QMatrix {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QBrush6matrixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix)
	return rv2
}

// /usr/include/qt/QtGui/qbrush.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMatrix(const QMatrix &)
func (this *QBrush) SetMatrix(mat *QMatrix) {
	var convArg0 = mat.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrush9setMatrixERK7QMatrix", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:97
// index:0
// Public inline Visibility=Default Availability=Available
// [88] QTransform transform()
func (this *QBrush) Transform() *QTransform /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QBrush9transformEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qbrush.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTransform(const QTransform &)
func (this *QBrush) SetTransform(arg0 *QTransform) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrush12setTransformERK10QTransform", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:100
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap texture()
func (this *QBrush) Texture() *QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QBrush7textureEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qbrush.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTexture(const QPixmap &)
func (this *QBrush) SetTexture(pixmap *QPixmap) {
	var convArg0 = pixmap.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrush10setTextureERK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:103
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage textureImage()
func (this *QBrush) TextureImage() *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QBrush12textureImageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qbrush.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextureImage(const QImage &)
func (this *QBrush) SetTextureImage(image *QImage) {
	var convArg0 = image.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrush15setTextureImageERK6QImage", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:106
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QColor & color()
func (this *QBrush) Color() *QColor {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QBrush5colorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}

// /usr/include/qt/QtGui/qbrush.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColor(const QColor &)
func (this *QBrush) SetColor(color *QColor) {
	var convArg0 = color.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrush8setColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:108
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setColor(Qt::GlobalColor)
func (this *QBrush) SetColor_1(color int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrush8setColorEN2Qt11GlobalColorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), color)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:110
// index:0
// Public Visibility=Default Availability=Available
// [8] const QGradient * gradient()
func (this *QBrush) Gradient() *QGradient /*777 const QGradient **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QBrush8gradientEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGradientFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qbrush.h:112
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isOpaque()
func (this *QBrush) IsOpaque() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QBrush8isOpaqueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qbrush.h:129
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDetached()
func (this *QBrush) IsDetached() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QBrush10isDetachedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
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
