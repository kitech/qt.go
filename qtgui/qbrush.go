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
// extern C begin: 61
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QBrush struct {
	*qtrt.CObject
}
type QBrush_ITF interface {
	QBrush_PTR() *QBrush
}

func (ptr *QBrush) QBrush_PTR() *QBrush { return ptr }

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

/*
Constructs a default black brush with the style Qt::NoBrush (i.e. this brush will not fill shapes).
*/
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

/*
Constructs a default black brush with the style Qt::NoBrush (i.e. this brush will not fill shapes).
*/
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

/*
Constructs a default black brush with the style Qt::NoBrush (i.e. this brush will not fill shapes).
*/
func NewQBrush_2(color QColor_ITF, bs int) *QBrush {
	var convArg0 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg0 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrushC2ERK6QColorN2Qt10BrushStyleE", qtrt.FFI_TYPE_POINTER, convArg0, bs)
	qtrt.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQBrush)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:68
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QBrush(const QColor &, Qt::BrushStyle)

/*
Constructs a default black brush with the style Qt::NoBrush (i.e. this brush will not fill shapes).
*/
func NewQBrush_2_(color QColor_ITF) *QBrush {
	var convArg0 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg0 = color.QColor_PTR().GetCthis()
	}
	// arg: 1, Qt::BrushStyle=Elaborated, Qt::BrushStyle=Enum,
	bs := 0
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

/*
Constructs a default black brush with the style Qt::NoBrush (i.e. this brush will not fill shapes).
*/
func NewQBrush_3(color int, bs int) *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrushC2EN2Qt11GlobalColorENS0_10BrushStyleE", qtrt.FFI_TYPE_POINTER, color, bs)
	qtrt.ErrPrint(err, rv)
	gothis := NewQBrushFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQBrush)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:69
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QBrush(Qt::GlobalColor, Qt::BrushStyle)

/*
Constructs a default black brush with the style Qt::NoBrush (i.e. this brush will not fill shapes).
*/
func NewQBrush_3_(color int) *QBrush {
	// arg: 1, Qt::BrushStyle=Elaborated, Qt::BrushStyle=Enum,
	bs := 0
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

/*
Constructs a default black brush with the style Qt::NoBrush (i.e. this brush will not fill shapes).
*/
func NewQBrush_4(color QColor_ITF, pixmap QPixmap_ITF) *QBrush {
	var convArg0 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg0 = color.QColor_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg1 = pixmap.QPixmap_PTR().GetCthis()
	}
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

/*
Constructs a default black brush with the style Qt::NoBrush (i.e. this brush will not fill shapes).
*/
func NewQBrush_5(color int, pixmap QPixmap_ITF) *QBrush {
	var convArg1 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg1 = pixmap.QPixmap_PTR().GetCthis()
	}
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

/*
Constructs a default black brush with the style Qt::NoBrush (i.e. this brush will not fill shapes).
*/
func NewQBrush_6(pixmap QPixmap_ITF) *QBrush {
	var convArg0 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg0 = pixmap.QPixmap_PTR().GetCthis()
	}
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

/*
Constructs a default black brush with the style Qt::NoBrush (i.e. this brush will not fill shapes).
*/
func NewQBrush_7(image QImage_ITF) *QBrush {
	var convArg0 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg0 = image.QImage_PTR().GetCthis()
	}
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

/*
Constructs a default black brush with the style Qt::NoBrush (i.e. this brush will not fill shapes).
*/
func NewQBrush_8(gradient QGradient_ITF) *QBrush {
	var convArg0 unsafe.Pointer
	if gradient != nil && gradient.QGradient_PTR() != nil {
		convArg0 = gradient.QGradient_PTR().GetCthis()
	}
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

/*

 */
func DeleteQBrush(this *QBrush) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrushD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qbrush.h:81
// index:0
// Public Visibility=Default Availability=Available
// [8] QBrush & operator=(const QBrush &)

/*

 */
func (this *QBrush) Operator_equal(brush QBrush_ITF) *QBrush {
	var convArg0 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg0 = brush.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrushaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qbrush.h:83
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QBrush & operator=(QBrush &&)

/*

 */
func (this *QBrush) Operator_equal_1(other unsafe.Pointer /*333*/) *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrushaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qbrush.h:86
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QBrush &)

/*
Swaps brush other with this brush. This operation is very fast and never fails.

This function was introduced in  Qt 4.8.
*/
func (this *QBrush) Swap(other QBrush_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QBrush_PTR() != nil {
		convArg0 = other.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrush4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:91
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::BrushStyle style() const

/*
Returns the brush style.

See also setStyle().
*/
func (this *QBrush) Style() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QBrush5styleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qbrush.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStyle(Qt::BrushStyle)

/*
Sets the brush style to style.

See also style().
*/
func (this *QBrush) SetStyle(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrush8setStyleEN2Qt10BrushStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:94
// index:0
// Public inline Visibility=Default Availability=Available
// [48] const QMatrix & matrix() const

/*
Returns the current transformation matrix for the brush.

This function was introduced in  Qt 4.2.

See also setMatrix().
*/
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

/*
Sets matrix as an explicit transformation matrix on the current brush. The brush transformation matrix is merged with QPainter transformation matrix to produce the final result.

This function was introduced in  Qt 4.2.

See also matrix().
*/
func (this *QBrush) SetMatrix(mat QMatrix_ITF) {
	var convArg0 unsafe.Pointer
	if mat != nil && mat.QMatrix_PTR() != nil {
		convArg0 = mat.QMatrix_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrush9setMatrixERK7QMatrix", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:97
// index:0
// Public inline Visibility=Default Availability=Available
// [88] QTransform transform() const

/*
Returns the current transformation matrix for the brush.

This function was introduced in  Qt 4.3.

See also setTransform().
*/
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

/*
Sets matrix as an explicit transformation matrix on the current brush. The brush transformation matrix is merged with QPainter transformation matrix to produce the final result.

This function was introduced in  Qt 4.3.

See also transform().
*/
func (this *QBrush) SetTransform(arg0 QTransform_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTransform_PTR() != nil {
		convArg0 = arg0.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrush12setTransformERK10QTransform", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:100
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap texture() const

/*
Returns the custom brush pattern, or a null pixmap if no custom brush pattern has been set.

See also setTexture().
*/
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

/*
Sets the brush pixmap to pixmap. The style is set to Qt::TexturePattern.

The current brush color will only have an effect for monochrome pixmaps, i.e. for QPixmap::depth() == 1 (QBitmaps).

See also texture().
*/
func (this *QBrush) SetTexture(pixmap QPixmap_ITF) {
	var convArg0 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg0 = pixmap.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrush10setTextureERK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:103
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage textureImage() const

/*
Returns the custom brush pattern, or a null image if no custom brush pattern has been set.

If the texture was set as a QPixmap it will be converted to a QImage.

This function was introduced in  Qt 4.2.

See also setTextureImage().
*/
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

/*
Sets the brush image to image. The style is set to Qt::TexturePattern.

Note the current brush color will not have any affect on monochrome images, as opposed to calling setTexture() with a QBitmap. If you want to change the color of monochrome image brushes, either convert the image to QBitmap with QBitmap::fromImage() and set the resulting QBitmap as a texture, or change the entries in the color table for the image.

This function was introduced in  Qt 4.2.

See also textureImage() and setTexture().
*/
func (this *QBrush) SetTextureImage(image QImage_ITF) {
	var convArg0 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg0 = image.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrush15setTextureImageERK6QImage", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:106
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QColor & color() const

/*
Returns the brush color.

See also setColor().
*/
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

/*
Sets the brush color to the given color.

Note that calling setColor() will not make a difference if the style is a gradient. The same is the case if the style is Qt::TexturePattern style unless the current texture is a QBitmap.

See also color().
*/
func (this *QBrush) SetColor(color QColor_ITF) {
	var convArg0 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg0 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrush8setColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:108
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setColor(Qt::GlobalColor)

/*
Sets the brush color to the given color.

Note that calling setColor() will not make a difference if the style is a gradient. The same is the case if the style is Qt::TexturePattern style unless the current texture is a QBitmap.

See also color().
*/
func (this *QBrush) SetColor_1(color int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QBrush8setColorEN2Qt11GlobalColorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), color)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:110
// index:0
// Public Visibility=Default Availability=Available
// [8] const QGradient * gradient() const

/*
Returns the gradient describing this brush.
*/
func (this *QBrush) Gradient() *QGradient /*777 const QGradient **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QBrush8gradientEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGradientFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qbrush.h:112
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isOpaque() const

/*
Returns true if the brush is fully opaque otherwise false. A brush is considered opaque if:


The alpha component of the color() is 255.
Its texture() does not have an alpha channel and is not a QBitmap.
The colors in the gradient() all have an alpha component that is 255.
It is an extended radial gradient.
*/
func (this *QBrush) IsOpaque() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QBrush8isOpaqueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qbrush.h:114
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QBrush &) const

/*

 */
func (this *QBrush) Operator_equal_equal(b QBrush_ITF) bool {
	var convArg0 unsafe.Pointer
	if b != nil && b.QBrush_PTR() != nil {
		convArg0 = b.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QBrusheqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qbrush.h:115
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QBrush &) const

/*

 */
func (this *QBrush) Operator_not_equal(b QBrush_ITF) bool {
	var convArg0 unsafe.Pointer
	if b != nil && b.QBrush_PTR() != nil {
		convArg0 = b.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QBrushneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qbrush.h:129
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDetached() const

/*

 */
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
