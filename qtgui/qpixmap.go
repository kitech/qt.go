package qtgui

// /usr/include/qt/QtGui/qpixmap.h
// #include <qpixmap.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 112
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

// int metric(QPaintDevice::PaintDeviceMetric)
func (this *QPixmap) InheritMetric(f func(arg0 int) int) {
	qtrt.SetAllInheritCallback(this, "metric", f)
}

// QPixmap fromImageInPlace(QImage &, Qt::ImageConversionFlags)
func (this *QPixmap) InheritFromImageInPlace(f func(image *QImage, flags int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "fromImageInPlace", f)
}

/*

 */
type QPixmap struct {
	*QPaintDevice
}
type QPixmap_ITF interface {
	QPaintDevice_ITF
	QPixmap_PTR() *QPixmap
}

func (ptr *QPixmap) QPixmap_PTR() *QPixmap { return ptr }

func (this *QPixmap) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QPaintDevice.GetCthis()
	}
}
func (this *QPixmap) SetCthis(cthis unsafe.Pointer) {
	this.QPaintDevice = NewQPaintDeviceFromPointer(cthis)
}
func NewQPixmapFromPointer(cthis unsafe.Pointer) *QPixmap {
	bcthis0 := NewQPaintDeviceFromPointer(cthis)
	return &QPixmap{bcthis0}
}
func (*QPixmap) NewFromPointer(cthis unsafe.Pointer) *QPixmap {
	return NewQPixmapFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpixmap.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPixmap()

/*
Constructs a null pixmap.

See also isNull().
*/
func (*QPixmap) NewForInherit() *QPixmap {
	return NewQPixmap()
}
func NewQPixmap() *QPixmap {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmapC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPixmap)
	return gothis
}

// /usr/include/qt/QtGui/qpixmap.h:66
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPixmap(int, int)

/*
Constructs a null pixmap.

See also isNull().
*/
func (*QPixmap) NewForInherit1(w int, h int) *QPixmap {
	return NewQPixmap1(w, h)
}
func NewQPixmap1(w int, h int) *QPixmap {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmapC2Eii", qtrt.FFI_TYPE_POINTER, w, h)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPixmap)
	return gothis
}

// /usr/include/qt/QtGui/qpixmap.h:67
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QPixmap(const QSize &)

/*
Constructs a null pixmap.

See also isNull().
*/
func (*QPixmap) NewForInherit2(arg0 qtcore.QSize_ITF) *QPixmap {
	return NewQPixmap2(arg0)
}
func NewQPixmap2(arg0 qtcore.QSize_ITF) *QPixmap {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSize_PTR() != nil {
		convArg0 = arg0.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmapC2ERK5QSize", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPixmap)
	return gothis
}

// /usr/include/qt/QtGui/qpixmap.h:68
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QPixmap(const QString &, const char *, Qt::ImageConversionFlags)

/*
Constructs a null pixmap.

See also isNull().
*/
func (*QPixmap) NewForInherit3(fileName string, format string, flags int) *QPixmap {
	return NewQPixmap3(fileName, format, flags)
}
func NewQPixmap3(fileName string, format string, flags int) *QPixmap {
	var tmpArg0 = qtcore.NewQString5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmapC2ERK7QStringPKc6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPixmap)
	return gothis
}

// /usr/include/qt/QtGui/qpixmap.h:68
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QPixmap(const QString &, const char *, Qt::ImageConversionFlags)

/*
Constructs a null pixmap.

See also isNull().
*/
func (*QPixmap) NewForInherit3p(fileName string) *QPixmap {
	return NewQPixmap3p(fileName)
}
func NewQPixmap3p(fileName string) *QPixmap {
	var tmpArg0 = qtcore.NewQString5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const char *=Pointer, =Invalid, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmapC2ERK7QStringPKc6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPixmap)
	return gothis
}

// /usr/include/qt/QtGui/qpixmap.h:68
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QPixmap(const QString &, const char *, Qt::ImageConversionFlags)

/*
Constructs a null pixmap.

See also isNull().
*/
func (*QPixmap) NewForInherit3p1(fileName string, format string) *QPixmap {
	return NewQPixmap3p1(fileName, format)
}
func NewQPixmap3p1(fileName string, format string) *QPixmap {
	var tmpArg0 = qtcore.NewQString5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	// arg: 2, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmapC2ERK7QStringPKc6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPixmap)
	return gothis
}

// /usr/include/qt/QtGui/qpixmap.h:70
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QPixmap(const char *const *)

/*
Constructs a null pixmap.

See also isNull().
*/
func (*QPixmap) NewForInherit4(xpm []string) *QPixmap {
	return NewQPixmap4(xpm)
}
func NewQPixmap4(xpm []string) *QPixmap {
	var convArg0 = qtrt.StringSliceToCCharPP(xpm)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmapC2EPKPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPixmap)
	return gothis
}

// /usr/include/qt/QtGui/qpixmap.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QPixmap()

/*

 */
func DeleteQPixmap(this *QPixmap) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmapD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpixmap.h:75
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap & operator=(const QPixmap &)

/*

 */
func (this *QPixmap) Operator_equal(arg0 QPixmap_ITF) *QPixmap {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPixmap_PTR() != nil {
		convArg0 = arg0.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmapaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:77
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QPixmap & operator=(QPixmap &&)

/*

 */
func (this *QPixmap) Operator_equal1(other unsafe.Pointer /*333*/) *QPixmap {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmapaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:80
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QPixmap &)

/*
Swaps pixmap other with this pixmap. This operation is very fast and never fails.

This function was introduced in  Qt 4.8.
*/
func (this *QPixmap) Swap(other QPixmap_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPixmap_PTR() != nil {
		convArg0 = other.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:85
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if this is a null pixmap; otherwise returns false.

A null pixmap has zero width, zero height and no contents. You cannot draw in a null pixmap.
*/
func (this *QPixmap) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int devType() const

/*

 */
func (this *QPixmap) DevType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap7devTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpixmap.h:88
// index:0
// Public Visibility=Default Availability=Available
// [4] int width() const

/*
Returns the width of the pixmap.

See also size() and Pixmap Information.
*/
func (this *QPixmap) Width() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap5widthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpixmap.h:89
// index:0
// Public Visibility=Default Availability=Available
// [4] int height() const

/*
Returns the height of the pixmap.

See also size() and Pixmap Information.
*/
func (this *QPixmap) Height() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap6heightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpixmap.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize size() const

/*
Returns the size of the pixmap.

See also width(), height(), and Pixmap Information.
*/
func (this *QPixmap) Size() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:91
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect rect() const

/*
Returns the pixmap's enclosing rectangle.

See also Pixmap Information.
*/
func (this *QPixmap) Rect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap4rectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:92
// index:0
// Public Visibility=Default Availability=Available
// [4] int depth() const

/*
Returns the depth of the pixmap.

The pixmap depth is also called bits per pixel (bpp) or bit planes of a pixmap. A null pixmap has depth 0.

See also defaultDepth() and Pixmap Information.
*/
func (this *QPixmap) Depth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap5depthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpixmap.h:94
// index:0
// Public static Visibility=Default Availability=Available
// [4] int defaultDepth()

/*
Returns the default pixmap depth used by the application.

On all platforms the depth of the primary screen will be returned.

Note: QGuiApplication must be created before calling this function.

See also depth(), QColormap::depth(), and Pixmap Information.
*/
func (this *QPixmap) DefaultDepth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap12defaultDepthEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QPixmap_DefaultDepth() int {
	var nilthis *QPixmap
	rv := nilthis.DefaultDepth()
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void fill(const QColor &)

/*
Fills the pixmap with the given color.

The effect of this function is undefined when the pixmap is being painted on.

See also Pixmap Transformations.
*/
func (this *QPixmap) Fill(fillColor QColor_ITF) {
	var convArg0 unsafe.Pointer
	if fillColor != nil && fillColor.QColor_PTR() != nil {
		convArg0 = fillColor.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap4fillERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void fill(const QColor &)

/*
Fills the pixmap with the given color.

The effect of this function is undefined when the pixmap is being painted on.

See also Pixmap Transformations.
*/
func (this *QPixmap) Fillp() {
	// arg: 0, const QColor &=LValueReference, QColor=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap4fillERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:97
// index:1
// Public Visibility=Default Availability=Available
// [-2] void fill(const QPaintDevice *, const QPoint &)

/*
Fills the pixmap with the given color.

The effect of this function is undefined when the pixmap is being painted on.

See also Pixmap Transformations.
*/
func (this *QPixmap) Fill1(device QPaintDevice_ITF /*777 const QPaintDevice **/, ofs qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QPaintDevice_PTR() != nil {
		convArg0 = device.QPaintDevice_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if ofs != nil && ofs.QPoint_PTR() != nil {
		convArg1 = ofs.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap4fillEPK12QPaintDeviceRK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:98
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void fill(const QPaintDevice *, int, int)

/*
Fills the pixmap with the given color.

The effect of this function is undefined when the pixmap is being painted on.

See also Pixmap Transformations.
*/
func (this *QPixmap) Fill2(device QPaintDevice_ITF /*777 const QPaintDevice **/, xofs int, yofs int) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QPaintDevice_PTR() != nil {
		convArg0 = device.QPaintDevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap4fillEPK12QPaintDeviceii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xofs, yofs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:100
// index:0
// Public Visibility=Default Availability=Available
// [32] QBitmap mask() const

/*
Extracts a bitmap mask from the pixmap's alpha channel.

Warning: This is potentially an expensive operation. The mask of the pixmap is extracted dynamically from the pixeldata.

See also setMask() and Pixmap Information.
*/
func (this *QPixmap) Mask() *QBitmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap4maskEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBitmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBitmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMask(const QBitmap &)

/*
Sets a mask bitmap.

This function merges the mask with the pixmap's alpha channel. A pixel value of 1 on the mask means the pixmap's pixel is unchanged; a value of 0 means the pixel is transparent. The mask must have the same size as this pixmap.

Setting a null mask resets the mask, leaving the previously transparent pixels black. The effect of this function is undefined when the pixmap is being painted on.

Warning: This is potentially an expensive operation.

See also mask(), Pixmap Transformations, and QBitmap.
*/
func (this *QPixmap) SetMask(arg0 QBitmap_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QBitmap_PTR() != nil {
		convArg0 = arg0.QBitmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap7setMaskERK7QBitmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal devicePixelRatio() const

/*
Returns the device pixel ratio for the pixmap. This is the ratio between device pixels and device independent pixels.

Use this function when calculating layout geometry based on the pixmap size: QSize layoutSize = image.size() / image.devicePixelRatio()

The default value is 1.0.

See also setDevicePixelRatio() and QImageReader.
*/
func (this *QPixmap) DevicePixelRatio() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap16devicePixelRatioEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpixmap.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDevicePixelRatio(qreal)

/*
Sets the device pixel ratio for the pixmap. This is the ratio between image pixels and device-independent pixels.

The default scaleFactor is 1.0. Setting it to something else has two effects:

QPainters that are opened on the pixmap will be scaled. For example, painting on a 200x200 image if with a ratio of 2.0 will result in effective (device-independent) painting bounds of 100x100.

Code paths in Qt that calculate layout geometry based on the pixmap size will take the ratio into account: QSize layoutSize = pixmap.size() / pixmap.devicePixelRatio() The net effect of this is that the pixmap is displayed as high-DPI pixmap rather than a large pixmap (see Drawing High Resolution Versions of Pixmaps and Images).

See also devicePixelRatio().
*/
func (this *QPixmap) SetDevicePixelRatio(scaleFactor float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap19setDevicePixelRatioEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), scaleFactor)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:106
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasAlpha() const

/*
Returns true if this pixmap has an alpha channel, or has a mask, otherwise returns false.

See also hasAlphaChannel() and mask().
*/
func (this *QPixmap) HasAlpha() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap8hasAlphaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:107
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasAlphaChannel() const

/*
Returns true if the pixmap has a format that respects the alpha channel, otherwise returns false.

See also hasAlpha().
*/
func (this *QPixmap) HasAlphaChannel() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap15hasAlphaChannelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:110
// index:0
// Public Visibility=Default Availability=Available
// [32] QBitmap createHeuristicMask(bool) const

/*
Creates and returns a heuristic mask for this pixmap.

The function works by selecting a color from one of the corners and then chipping away pixels of that color, starting at all the edges. If clipTight is true (the default) the mask is just large enough to cover the pixels; otherwise, the mask is larger than the data pixels.

The mask may not be perfect but it should be reasonable, so you can do things such as the following:


  QPixmap myPixmap;
  myPixmap.setMask(myPixmap.createHeuristicMask());



This function is slow because it involves converting to/from a QImage, and non-trivial computations.

See also QImage::createHeuristicMask() and createMaskFromColor().
*/
func (this *QPixmap) CreateHeuristicMask(clipTight bool) *QBitmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap19createHeuristicMaskEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), clipTight)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBitmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBitmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:110
// index:0
// Public Visibility=Default Availability=Available
// [32] QBitmap createHeuristicMask(bool) const

/*
Creates and returns a heuristic mask for this pixmap.

The function works by selecting a color from one of the corners and then chipping away pixels of that color, starting at all the edges. If clipTight is true (the default) the mask is just large enough to cover the pixels; otherwise, the mask is larger than the data pixels.

The mask may not be perfect but it should be reasonable, so you can do things such as the following:


  QPixmap myPixmap;
  myPixmap.setMask(myPixmap.createHeuristicMask());



This function is slow because it involves converting to/from a QImage, and non-trivial computations.

See also QImage::createHeuristicMask() and createMaskFromColor().
*/
func (this *QPixmap) CreateHeuristicMaskp() *QBitmap /*123*/ {
	// arg: 0, bool=Bool, =Invalid, , Invalid
	clipTight := true
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap19createHeuristicMaskEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), clipTight)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBitmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBitmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:112
// index:0
// Public Visibility=Default Availability=Available
// [32] QBitmap createMaskFromColor(const QColor &, Qt::MaskMode) const

/*
Creates and returns a mask for this pixmap based on the given maskColor. If the mode is Qt::MaskInColor, all pixels matching the maskColor will be transparent. If mode is Qt::MaskOutColor, all pixels matching the maskColor will be opaque.

This function is slow because it involves converting to/from a QImage.

See also createHeuristicMask() and QImage::createMaskFromColor().
*/
func (this *QPixmap) CreateMaskFromColor(maskColor QColor_ITF, mode int) *QBitmap /*123*/ {
	var convArg0 unsafe.Pointer
	if maskColor != nil && maskColor.QColor_PTR() != nil {
		convArg0 = maskColor.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap19createMaskFromColorERK6QColorN2Qt8MaskModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBitmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBitmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:112
// index:0
// Public Visibility=Default Availability=Available
// [32] QBitmap createMaskFromColor(const QColor &, Qt::MaskMode) const

/*
Creates and returns a mask for this pixmap based on the given maskColor. If the mode is Qt::MaskInColor, all pixels matching the maskColor will be transparent. If mode is Qt::MaskOutColor, all pixels matching the maskColor will be opaque.

This function is slow because it involves converting to/from a QImage.

See also createHeuristicMask() and QImage::createMaskFromColor().
*/
func (this *QPixmap) CreateMaskFromColorp(maskColor QColor_ITF) *QBitmap /*123*/ {
	var convArg0 unsafe.Pointer
	if maskColor != nil && maskColor.QColor_PTR() != nil {
		convArg0 = maskColor.QColor_PTR().GetCthis()
	}
	// arg: 1, Qt::MaskMode=Elaborated, Qt::MaskMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap19createMaskFromColorERK6QColorN2Qt8MaskModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBitmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBitmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:114
// index:0
// Public static Visibility=Default Availability=Available
// [32] QPixmap grabWindow(WId, int, int, int, int)

/*

 */
func (this *QPixmap) GrabWindow(arg0 uint64, x int, y int, w int, h int) *QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap10grabWindowEyiiii", qtrt.FFI_TYPE_POINTER, arg0, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}
func QPixmap_GrabWindow(arg0 uint64, x int, y int, w int, h int) *QPixmap /*123*/ {
	var nilthis *QPixmap
	rv := nilthis.GrabWindow(arg0, x, y, w, h)
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:114
// index:0
// Public static Visibility=Default Availability=Available
// [32] QPixmap grabWindow(WId, int, int, int, int)

/*

 */
func (this *QPixmap) GrabWindowp(arg0 uint64) *QPixmap /*123*/ {
	// arg: 1, int=Int, =Invalid, , Invalid
	x := int(0)
	// arg: 2, int=Int, =Invalid, , Invalid
	y := int(0)
	// arg: 3, int=Int, =Invalid, , Invalid
	w := int(-1)
	// arg: 4, int=Int, =Invalid, , Invalid
	h := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap10grabWindowEyiiii", qtrt.FFI_TYPE_POINTER, arg0, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:114
// index:0
// Public static Visibility=Default Availability=Available
// [32] QPixmap grabWindow(WId, int, int, int, int)

/*

 */
func (this *QPixmap) GrabWindowp1(arg0 uint64, x int) *QPixmap /*123*/ {
	// arg: 2, int=Int, =Invalid, , Invalid
	y := int(0)
	// arg: 3, int=Int, =Invalid, , Invalid
	w := int(-1)
	// arg: 4, int=Int, =Invalid, , Invalid
	h := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap10grabWindowEyiiii", qtrt.FFI_TYPE_POINTER, arg0, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:114
// index:0
// Public static Visibility=Default Availability=Available
// [32] QPixmap grabWindow(WId, int, int, int, int)

/*

 */
func (this *QPixmap) GrabWindowp2(arg0 uint64, x int, y int) *QPixmap /*123*/ {
	// arg: 3, int=Int, =Invalid, , Invalid
	w := int(-1)
	// arg: 4, int=Int, =Invalid, , Invalid
	h := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap10grabWindowEyiiii", qtrt.FFI_TYPE_POINTER, arg0, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:114
// index:0
// Public static Visibility=Default Availability=Available
// [32] QPixmap grabWindow(WId, int, int, int, int)

/*

 */
func (this *QPixmap) GrabWindowp3(arg0 uint64, x int, y int, w int) *QPixmap /*123*/ {
	// arg: 4, int=Int, =Invalid, , Invalid
	h := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap10grabWindowEyiiii", qtrt.FFI_TYPE_POINTER, arg0, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:115
// index:0
// Public static Visibility=Default Availability=Available
// [32] QPixmap grabWidget(QObject *, const QRect &)

/*

 */
func (this *QPixmap) GrabWidget(widget qtcore.QObject_ITF /*777 QObject **/, rect qtcore.QRect_ITF) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QObject_PTR() != nil {
		convArg0 = widget.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg1 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap10grabWidgetEP7QObjectRK5QRect", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}
func QPixmap_GrabWidget(widget qtcore.QObject_ITF /*777 QObject **/, rect qtcore.QRect_ITF) *QPixmap /*123*/ {
	var nilthis *QPixmap
	rv := nilthis.GrabWidget(widget, rect)
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:116
// index:1
// Public static inline Visibility=Default Availability=Available
// [32] QPixmap grabWidget(QObject *, int, int, int, int)

/*

 */
func (this *QPixmap) GrabWidget1(widget qtcore.QObject_ITF /*777 QObject **/, x int, y int, w int, h int) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QObject_PTR() != nil {
		convArg0 = widget.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap10grabWidgetEP7QObjectiiii", qtrt.FFI_TYPE_POINTER, convArg0, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}
func QPixmap_GrabWidget1(widget qtcore.QObject_ITF /*777 QObject **/, x int, y int, w int, h int) *QPixmap /*123*/ {
	var nilthis *QPixmap
	rv := nilthis.GrabWidget1(widget, x, y, w, h)
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:116
// index:1
// Public static inline Visibility=Default Availability=Available
// [32] QPixmap grabWidget(QObject *, int, int, int, int)

/*

 */
func (this *QPixmap) GrabWidget1p(widget qtcore.QObject_ITF /*777 QObject **/) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QObject_PTR() != nil {
		convArg0 = widget.QObject_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	x := int(0)
	// arg: 2, int=Int, =Invalid, , Invalid
	y := int(0)
	// arg: 3, int=Int, =Invalid, , Invalid
	w := int(-1)
	// arg: 4, int=Int, =Invalid, , Invalid
	h := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap10grabWidgetEP7QObjectiiii", qtrt.FFI_TYPE_POINTER, convArg0, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:116
// index:1
// Public static inline Visibility=Default Availability=Available
// [32] QPixmap grabWidget(QObject *, int, int, int, int)

/*

 */
func (this *QPixmap) GrabWidget1p1(widget qtcore.QObject_ITF /*777 QObject **/, x int) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QObject_PTR() != nil {
		convArg0 = widget.QObject_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid, , Invalid
	y := int(0)
	// arg: 3, int=Int, =Invalid, , Invalid
	w := int(-1)
	// arg: 4, int=Int, =Invalid, , Invalid
	h := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap10grabWidgetEP7QObjectiiii", qtrt.FFI_TYPE_POINTER, convArg0, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:116
// index:1
// Public static inline Visibility=Default Availability=Available
// [32] QPixmap grabWidget(QObject *, int, int, int, int)

/*

 */
func (this *QPixmap) GrabWidget1p2(widget qtcore.QObject_ITF /*777 QObject **/, x int, y int) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QObject_PTR() != nil {
		convArg0 = widget.QObject_PTR().GetCthis()
	}
	// arg: 3, int=Int, =Invalid, , Invalid
	w := int(-1)
	// arg: 4, int=Int, =Invalid, , Invalid
	h := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap10grabWidgetEP7QObjectiiii", qtrt.FFI_TYPE_POINTER, convArg0, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:116
// index:1
// Public static inline Visibility=Default Availability=Available
// [32] QPixmap grabWidget(QObject *, int, int, int, int)

/*

 */
func (this *QPixmap) GrabWidget1p3(widget qtcore.QObject_ITF /*777 QObject **/, x int, y int, w int) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QObject_PTR() != nil {
		convArg0 = widget.QObject_PTR().GetCthis()
	}
	// arg: 4, int=Int, =Invalid, , Invalid
	h := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap10grabWidgetEP7QObjectiiii", qtrt.FFI_TYPE_POINTER, convArg0, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:119
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QPixmap scaled(int, int, Qt::AspectRatioMode, Qt::TransformationMode) const

/*
Scales the pixmap to the given size, using the aspect ratio and transformation modes specified by aspectRatioMode and transformMode.




If aspectRatioMode is Qt::IgnoreAspectRatio, the pixmap is scaled to size.
If aspectRatioMode is Qt::KeepAspectRatio, the pixmap is scaled to a rectangle as large as possible inside size, preserving the aspect ratio.
If aspectRatioMode is Qt::KeepAspectRatioByExpanding, the pixmap is scaled to a rectangle as small as possible outside size, preserving the aspect ratio.


If the given size is empty, this function returns a null pixmap.

In some cases it can be more beneficial to draw the pixmap to a painter with a scale set rather than scaling the pixmap. This is the case when the painter is for instance based on OpenGL or when the scale factor changes rapidly.

See also isNull() and Pixmap Transformations.
*/
func (this *QPixmap) Scaled(w int, h int, aspectMode int, mode int) *QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap6scaledEiiN2Qt15AspectRatioModeENS0_18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h, aspectMode, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:119
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QPixmap scaled(int, int, Qt::AspectRatioMode, Qt::TransformationMode) const

/*
Scales the pixmap to the given size, using the aspect ratio and transformation modes specified by aspectRatioMode and transformMode.




If aspectRatioMode is Qt::IgnoreAspectRatio, the pixmap is scaled to size.
If aspectRatioMode is Qt::KeepAspectRatio, the pixmap is scaled to a rectangle as large as possible inside size, preserving the aspect ratio.
If aspectRatioMode is Qt::KeepAspectRatioByExpanding, the pixmap is scaled to a rectangle as small as possible outside size, preserving the aspect ratio.


If the given size is empty, this function returns a null pixmap.

In some cases it can be more beneficial to draw the pixmap to a painter with a scale set rather than scaling the pixmap. This is the case when the painter is for instance based on OpenGL or when the scale factor changes rapidly.

See also isNull() and Pixmap Transformations.
*/
func (this *QPixmap) Scaledp(w int, h int) *QPixmap /*123*/ {
	// arg: 2, Qt::AspectRatioMode=Elaborated, Qt::AspectRatioMode=Enum, , Invalid
	aspectMode := 0
	// arg: 3, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap6scaledEiiN2Qt15AspectRatioModeENS0_18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h, aspectMode, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:119
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QPixmap scaled(int, int, Qt::AspectRatioMode, Qt::TransformationMode) const

/*
Scales the pixmap to the given size, using the aspect ratio and transformation modes specified by aspectRatioMode and transformMode.




If aspectRatioMode is Qt::IgnoreAspectRatio, the pixmap is scaled to size.
If aspectRatioMode is Qt::KeepAspectRatio, the pixmap is scaled to a rectangle as large as possible inside size, preserving the aspect ratio.
If aspectRatioMode is Qt::KeepAspectRatioByExpanding, the pixmap is scaled to a rectangle as small as possible outside size, preserving the aspect ratio.


If the given size is empty, this function returns a null pixmap.

In some cases it can be more beneficial to draw the pixmap to a painter with a scale set rather than scaling the pixmap. This is the case when the painter is for instance based on OpenGL or when the scale factor changes rapidly.

See also isNull() and Pixmap Transformations.
*/
func (this *QPixmap) Scaledp1(w int, h int, aspectMode int) *QPixmap /*123*/ {
	// arg: 3, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap6scaledEiiN2Qt15AspectRatioModeENS0_18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h, aspectMode, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:122
// index:1
// Public Visibility=Default Availability=Available
// [32] QPixmap scaled(const QSize &, Qt::AspectRatioMode, Qt::TransformationMode) const

/*
Scales the pixmap to the given size, using the aspect ratio and transformation modes specified by aspectRatioMode and transformMode.




If aspectRatioMode is Qt::IgnoreAspectRatio, the pixmap is scaled to size.
If aspectRatioMode is Qt::KeepAspectRatio, the pixmap is scaled to a rectangle as large as possible inside size, preserving the aspect ratio.
If aspectRatioMode is Qt::KeepAspectRatioByExpanding, the pixmap is scaled to a rectangle as small as possible outside size, preserving the aspect ratio.


If the given size is empty, this function returns a null pixmap.

In some cases it can be more beneficial to draw the pixmap to a painter with a scale set rather than scaling the pixmap. This is the case when the painter is for instance based on OpenGL or when the scale factor changes rapidly.

See also isNull() and Pixmap Transformations.
*/
func (this *QPixmap) Scaled1(s qtcore.QSize_ITF, aspectMode int, mode int) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if s != nil && s.QSize_PTR() != nil {
		convArg0 = s.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap6scaledERK5QSizeN2Qt15AspectRatioModeENS3_18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, aspectMode, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:122
// index:1
// Public Visibility=Default Availability=Available
// [32] QPixmap scaled(const QSize &, Qt::AspectRatioMode, Qt::TransformationMode) const

/*
Scales the pixmap to the given size, using the aspect ratio and transformation modes specified by aspectRatioMode and transformMode.




If aspectRatioMode is Qt::IgnoreAspectRatio, the pixmap is scaled to size.
If aspectRatioMode is Qt::KeepAspectRatio, the pixmap is scaled to a rectangle as large as possible inside size, preserving the aspect ratio.
If aspectRatioMode is Qt::KeepAspectRatioByExpanding, the pixmap is scaled to a rectangle as small as possible outside size, preserving the aspect ratio.


If the given size is empty, this function returns a null pixmap.

In some cases it can be more beneficial to draw the pixmap to a painter with a scale set rather than scaling the pixmap. This is the case when the painter is for instance based on OpenGL or when the scale factor changes rapidly.

See also isNull() and Pixmap Transformations.
*/
func (this *QPixmap) Scaled1p(s qtcore.QSize_ITF) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if s != nil && s.QSize_PTR() != nil {
		convArg0 = s.QSize_PTR().GetCthis()
	}
	// arg: 1, Qt::AspectRatioMode=Elaborated, Qt::AspectRatioMode=Enum, , Invalid
	aspectMode := 0
	// arg: 2, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap6scaledERK5QSizeN2Qt15AspectRatioModeENS3_18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, aspectMode, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:122
// index:1
// Public Visibility=Default Availability=Available
// [32] QPixmap scaled(const QSize &, Qt::AspectRatioMode, Qt::TransformationMode) const

/*
Scales the pixmap to the given size, using the aspect ratio and transformation modes specified by aspectRatioMode and transformMode.




If aspectRatioMode is Qt::IgnoreAspectRatio, the pixmap is scaled to size.
If aspectRatioMode is Qt::KeepAspectRatio, the pixmap is scaled to a rectangle as large as possible inside size, preserving the aspect ratio.
If aspectRatioMode is Qt::KeepAspectRatioByExpanding, the pixmap is scaled to a rectangle as small as possible outside size, preserving the aspect ratio.


If the given size is empty, this function returns a null pixmap.

In some cases it can be more beneficial to draw the pixmap to a painter with a scale set rather than scaling the pixmap. This is the case when the painter is for instance based on OpenGL or when the scale factor changes rapidly.

See also isNull() and Pixmap Transformations.
*/
func (this *QPixmap) Scaled1p1(s qtcore.QSize_ITF, aspectMode int) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if s != nil && s.QSize_PTR() != nil {
		convArg0 = s.QSize_PTR().GetCthis()
	}
	// arg: 2, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap6scaledERK5QSizeN2Qt15AspectRatioModeENS3_18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, aspectMode, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:124
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap scaledToWidth(int, Qt::TransformationMode) const

/*
Returns a scaled copy of the image. The returned image is scaled to the given width using the specified transformation mode. The height of the pixmap is automatically calculated so that the aspect ratio of the pixmap is preserved.

If width is 0 or negative, a null pixmap is returned.

See also isNull() and Pixmap Transformations.
*/
func (this *QPixmap) ScaledToWidth(w int, mode int) *QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap13scaledToWidthEiN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:124
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap scaledToWidth(int, Qt::TransformationMode) const

/*
Returns a scaled copy of the image. The returned image is scaled to the given width using the specified transformation mode. The height of the pixmap is automatically calculated so that the aspect ratio of the pixmap is preserved.

If width is 0 or negative, a null pixmap is returned.

See also isNull() and Pixmap Transformations.
*/
func (this *QPixmap) ScaledToWidthp(w int) *QPixmap /*123*/ {
	// arg: 1, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap13scaledToWidthEiN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:125
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap scaledToHeight(int, Qt::TransformationMode) const

/*
Returns a scaled copy of the image. The returned image is scaled to the given height using the specified transformation mode. The width of the pixmap is automatically calculated so that the aspect ratio of the pixmap is preserved.

If height is 0 or negative, a null pixmap is returned.

See also isNull() and Pixmap Transformations.
*/
func (this *QPixmap) ScaledToHeight(h int, mode int) *QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap14scaledToHeightEiN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), h, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:125
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap scaledToHeight(int, Qt::TransformationMode) const

/*
Returns a scaled copy of the image. The returned image is scaled to the given height using the specified transformation mode. The width of the pixmap is automatically calculated so that the aspect ratio of the pixmap is preserved.

If height is 0 or negative, a null pixmap is returned.

See also isNull() and Pixmap Transformations.
*/
func (this *QPixmap) ScaledToHeightp(h int) *QPixmap /*123*/ {
	// arg: 1, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap14scaledToHeightEiN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), h, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:126
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap transformed(const QMatrix &, Qt::TransformationMode) const

/*
Returns a copy of the pixmap that is transformed using the given transformation transform and transformation mode. The original pixmap is not changed.

The transformation transform is internally adjusted to compensate for unwanted translation; i.e. the pixmap produced is the smallest pixmap that contains all the transformed points of the original pixmap. Use the trueMatrix() function to retrieve the actual matrix used for transforming the pixmap.

This function is slow because it involves transformation to a QImage, non-trivial computations and a transformation back to a QPixmap.

See also trueMatrix() and Pixmap Transformations.
*/
func (this *QPixmap) Transformed(arg0 QMatrix_ITF, mode int) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMatrix_PTR() != nil {
		convArg0 = arg0.QMatrix_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap11transformedERK7QMatrixN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:126
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap transformed(const QMatrix &, Qt::TransformationMode) const

/*
Returns a copy of the pixmap that is transformed using the given transformation transform and transformation mode. The original pixmap is not changed.

The transformation transform is internally adjusted to compensate for unwanted translation; i.e. the pixmap produced is the smallest pixmap that contains all the transformed points of the original pixmap. Use the trueMatrix() function to retrieve the actual matrix used for transforming the pixmap.

This function is slow because it involves transformation to a QImage, non-trivial computations and a transformation back to a QPixmap.

See also trueMatrix() and Pixmap Transformations.
*/
func (this *QPixmap) Transformedp(arg0 QMatrix_ITF) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMatrix_PTR() != nil {
		convArg0 = arg0.QMatrix_PTR().GetCthis()
	}
	// arg: 1, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap11transformedERK7QMatrixN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:128
// index:1
// Public Visibility=Default Availability=Available
// [32] QPixmap transformed(const QTransform &, Qt::TransformationMode) const

/*
Returns a copy of the pixmap that is transformed using the given transformation transform and transformation mode. The original pixmap is not changed.

The transformation transform is internally adjusted to compensate for unwanted translation; i.e. the pixmap produced is the smallest pixmap that contains all the transformed points of the original pixmap. Use the trueMatrix() function to retrieve the actual matrix used for transforming the pixmap.

This function is slow because it involves transformation to a QImage, non-trivial computations and a transformation back to a QPixmap.

See also trueMatrix() and Pixmap Transformations.
*/
func (this *QPixmap) Transformed1(arg0 QTransform_ITF, mode int) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTransform_PTR() != nil {
		convArg0 = arg0.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap11transformedERK10QTransformN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:128
// index:1
// Public Visibility=Default Availability=Available
// [32] QPixmap transformed(const QTransform &, Qt::TransformationMode) const

/*
Returns a copy of the pixmap that is transformed using the given transformation transform and transformation mode. The original pixmap is not changed.

The transformation transform is internally adjusted to compensate for unwanted translation; i.e. the pixmap produced is the smallest pixmap that contains all the transformed points of the original pixmap. Use the trueMatrix() function to retrieve the actual matrix used for transforming the pixmap.

This function is slow because it involves transformation to a QImage, non-trivial computations and a transformation back to a QPixmap.

See also trueMatrix() and Pixmap Transformations.
*/
func (this *QPixmap) Transformed1p(arg0 QTransform_ITF) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTransform_PTR() != nil {
		convArg0 = arg0.QTransform_PTR().GetCthis()
	}
	// arg: 1, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap11transformedERK10QTransformN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:127
// index:0
// Public static Visibility=Default Availability=Available
// [48] QMatrix trueMatrix(const QMatrix &, int, int)

/*
Returns the actual matrix used for transforming a pixmap with the given width, height and matrix.

When transforming a pixmap using the transformed() function, the transformation matrix is internally adjusted to compensate for unwanted translation, i.e. transformed() returns the smallest pixmap containing all transformed points of the original pixmap. This function returns the modified matrix, which maps points correctly from the original pixmap into the new pixmap.

See also transformed() and Pixmap Transformations.
*/
func (this *QPixmap) TrueMatrix(m QMatrix_ITF, w int, h int) *QMatrix /*123*/ {
	var convArg0 unsafe.Pointer
	if m != nil && m.QMatrix_PTR() != nil {
		convArg0 = m.QMatrix_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap10trueMatrixERK7QMatrixii", qtrt.FFI_TYPE_POINTER, convArg0, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix)
	return rv2
}
func QPixmap_TrueMatrix(m QMatrix_ITF, w int, h int) *QMatrix /*123*/ {
	var nilthis *QPixmap
	rv := nilthis.TrueMatrix(m, w, h)
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:129
// index:1
// Public static Visibility=Default Availability=Available
// [88] QTransform trueMatrix(const QTransform &, int, int)

/*
Returns the actual matrix used for transforming a pixmap with the given width, height and matrix.

When transforming a pixmap using the transformed() function, the transformation matrix is internally adjusted to compensate for unwanted translation, i.e. transformed() returns the smallest pixmap containing all transformed points of the original pixmap. This function returns the modified matrix, which maps points correctly from the original pixmap into the new pixmap.

See also transformed() and Pixmap Transformations.
*/
func (this *QPixmap) TrueMatrix1(m QTransform_ITF, w int, h int) *QTransform /*123*/ {
	var convArg0 unsafe.Pointer
	if m != nil && m.QTransform_PTR() != nil {
		convArg0 = m.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap10trueMatrixERK10QTransformii", qtrt.FFI_TYPE_POINTER, convArg0, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}
func QPixmap_TrueMatrix1(m QTransform_ITF, w int, h int) *QTransform /*123*/ {
	var nilthis *QPixmap
	rv := nilthis.TrueMatrix1(m, w, h)
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:131
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage toImage() const

/*
Converts the pixmap to a QImage. Returns a null image if the conversion fails.

If the pixmap has 1-bit depth, the returned image will also be 1 bit deep. Images with more bits will be returned in a format closely represents the underlying system. Usually this will be QImage::Format_ARGB32_Premultiplied for pixmaps with an alpha and QImage::Format_RGB32 or QImage::Format_RGB16 for pixmaps without alpha.

Note that for the moment, alpha masks on monochrome images are ignored.

See also fromImage() and Image Formats.
*/
func (this *QPixmap) ToImage() *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap7toImageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:132
// index:0
// Public static Visibility=Default Availability=Available
// [32] QPixmap fromImage(const QImage &, Qt::ImageConversionFlags)

/*
Converts the given image to a pixmap using the specified flags to control the conversion. The flags argument is a bitwise-OR of the Qt::ImageConversionFlags. Passing 0 for flags sets all the default options.

In case of monochrome and 8-bit images, the image is first converted to a 32-bit pixmap and then filled with the colors in the color table. If this is too expensive an operation, you can use QBitmap::fromImage() instead.

See also fromImageReader(), toImage(), and Pixmap Conversion.
*/
func (this *QPixmap) FromImage(image QImage_ITF, flags int) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg0 = image.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap9fromImageERK6QImage6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}
func QPixmap_FromImage(image QImage_ITF, flags int) *QPixmap /*123*/ {
	var nilthis *QPixmap
	rv := nilthis.FromImage(image, flags)
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:132
// index:0
// Public static Visibility=Default Availability=Available
// [32] QPixmap fromImage(const QImage &, Qt::ImageConversionFlags)

/*
Converts the given image to a pixmap using the specified flags to control the conversion. The flags argument is a bitwise-OR of the Qt::ImageConversionFlags. Passing 0 for flags sets all the default options.

In case of monochrome and 8-bit images, the image is first converted to a 32-bit pixmap and then filled with the colors in the color table. If this is too expensive an operation, you can use QBitmap::fromImage() instead.

See also fromImageReader(), toImage(), and Pixmap Conversion.
*/
func (this *QPixmap) FromImagep(image QImage_ITF) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg0 = image.QImage_PTR().GetCthis()
	}
	// arg: 1, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap9fromImageERK6QImage6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:135
// index:1
// Public static inline Visibility=Default Availability=Available
// [32] QPixmap fromImage(QImage &&, Qt::ImageConversionFlags)

/*
Converts the given image to a pixmap using the specified flags to control the conversion. The flags argument is a bitwise-OR of the Qt::ImageConversionFlags. Passing 0 for flags sets all the default options.

In case of monochrome and 8-bit images, the image is first converted to a 32-bit pixmap and then filled with the colors in the color table. If this is too expensive an operation, you can use QBitmap::fromImage() instead.

See also fromImageReader(), toImage(), and Pixmap Conversion.
*/
func (this *QPixmap) FromImage1(image unsafe.Pointer /*333*/, flags int) *QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap9fromImageEO6QImage6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, image, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}
func QPixmap_FromImage1(image unsafe.Pointer /*333*/, flags int) *QPixmap /*123*/ {
	var nilthis *QPixmap
	rv := nilthis.FromImage1(image, flags)
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:135
// index:1
// Public static inline Visibility=Default Availability=Available
// [32] QPixmap fromImage(QImage &&, Qt::ImageConversionFlags)

/*
Converts the given image to a pixmap using the specified flags to control the conversion. The flags argument is a bitwise-OR of the Qt::ImageConversionFlags. Passing 0 for flags sets all the default options.

In case of monochrome and 8-bit images, the image is first converted to a 32-bit pixmap and then filled with the colors in the color table. If this is too expensive an operation, you can use QBitmap::fromImage() instead.

See also fromImageReader(), toImage(), and Pixmap Conversion.
*/
func (this *QPixmap) FromImage1p(image unsafe.Pointer /*333*/) *QPixmap /*123*/ {
	// arg: 1, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap9fromImageEO6QImage6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, image, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:133
// index:0
// Public static Visibility=Default Availability=Available
// [32] QPixmap fromImageReader(QImageReader *, Qt::ImageConversionFlags)

/*
Create a QPixmap from an image read directly from an imageReader. The flags argument is a bitwise-OR of the Qt::ImageConversionFlags. Passing 0 for flags sets all the default options.

On some systems, reading an image directly to QPixmap can use less memory than reading a QImage to convert it to QPixmap.

See also fromImage(), toImage(), and Pixmap Conversion.
*/
func (this *QPixmap) FromImageReader(imageReader QImageReader_ITF /*777 QImageReader **/, flags int) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if imageReader != nil && imageReader.QImageReader_PTR() != nil {
		convArg0 = imageReader.QImageReader_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap15fromImageReaderEP12QImageReader6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}
func QPixmap_FromImageReader(imageReader QImageReader_ITF /*777 QImageReader **/, flags int) *QPixmap /*123*/ {
	var nilthis *QPixmap
	rv := nilthis.FromImageReader(imageReader, flags)
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:133
// index:0
// Public static Visibility=Default Availability=Available
// [32] QPixmap fromImageReader(QImageReader *, Qt::ImageConversionFlags)

/*
Create a QPixmap from an image read directly from an imageReader. The flags argument is a bitwise-OR of the Qt::ImageConversionFlags. Passing 0 for flags sets all the default options.

On some systems, reading an image directly to QPixmap can use less memory than reading a QImage to convert it to QPixmap.

See also fromImage(), toImage(), and Pixmap Conversion.
*/
func (this *QPixmap) FromImageReaderp(imageReader QImageReader_ITF /*777 QImageReader **/) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if imageReader != nil && imageReader.QImageReader_PTR() != nil {
		convArg0 = imageReader.QImageReader_PTR().GetCthis()
	}
	// arg: 1, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap15fromImageReaderEP12QImageReader6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:141
// index:0
// Public Visibility=Default Availability=Available
// [1] bool load(const QString &, const char *, Qt::ImageConversionFlags)

/*
Loads a pixmap from the file with the given fileName. Returns true if the pixmap was successfully loaded; otherwise invalidates the pixmap and returns false.

The loader attempts to read the pixmap using the specified format. If the format is not specified (which is the default), the loader probes the file for a header to guess the file format.

The file name can either refer to an actual file on disk or to one of the application's embedded resources. See the Resource System overview for details on how to embed pixmaps and other resource files in the application's executable.

If the data needs to be modified to fit in a lower-resolution result (e.g. converting from 32-bit to 8-bit), use the flags to control the conversion.

Note that QPixmaps are automatically added to the QPixmapCache when loaded from a file; the key used is internal and can not be acquired.

See also loadFromData() and Reading and Writing Image Files.
*/
func (this *QPixmap) Load(fileName string, format string, flags int) bool {
	var tmpArg0 = qtcore.NewQString5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap4loadERK7QStringPKc6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:141
// index:0
// Public Visibility=Default Availability=Available
// [1] bool load(const QString &, const char *, Qt::ImageConversionFlags)

/*
Loads a pixmap from the file with the given fileName. Returns true if the pixmap was successfully loaded; otherwise invalidates the pixmap and returns false.

The loader attempts to read the pixmap using the specified format. If the format is not specified (which is the default), the loader probes the file for a header to guess the file format.

The file name can either refer to an actual file on disk or to one of the application's embedded resources. See the Resource System overview for details on how to embed pixmaps and other resource files in the application's executable.

If the data needs to be modified to fit in a lower-resolution result (e.g. converting from 32-bit to 8-bit), use the flags to control the conversion.

Note that QPixmaps are automatically added to the QPixmapCache when loaded from a file; the key used is internal and can not be acquired.

See also loadFromData() and Reading and Writing Image Files.
*/
func (this *QPixmap) Loadp(fileName string) bool {
	var tmpArg0 = qtcore.NewQString5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const char *=Pointer, =Invalid, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap4loadERK7QStringPKc6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:141
// index:0
// Public Visibility=Default Availability=Available
// [1] bool load(const QString &, const char *, Qt::ImageConversionFlags)

/*
Loads a pixmap from the file with the given fileName. Returns true if the pixmap was successfully loaded; otherwise invalidates the pixmap and returns false.

The loader attempts to read the pixmap using the specified format. If the format is not specified (which is the default), the loader probes the file for a header to guess the file format.

The file name can either refer to an actual file on disk or to one of the application's embedded resources. See the Resource System overview for details on how to embed pixmaps and other resource files in the application's executable.

If the data needs to be modified to fit in a lower-resolution result (e.g. converting from 32-bit to 8-bit), use the flags to control the conversion.

Note that QPixmaps are automatically added to the QPixmapCache when loaded from a file; the key used is internal and can not be acquired.

See also loadFromData() and Reading and Writing Image Files.
*/
func (this *QPixmap) Loadp1(fileName string, format string) bool {
	var tmpArg0 = qtcore.NewQString5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	// arg: 2, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap4loadERK7QStringPKc6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:142
// index:0
// Public Visibility=Default Availability=Available
// [1] bool loadFromData(const uchar *, uint, const char *, Qt::ImageConversionFlags)

/*
Loads a pixmap from the len first bytes of the given binary data. Returns true if the pixmap was loaded successfully; otherwise invalidates the pixmap and returns false.

The loader attempts to read the pixmap using the specified format. If the format is not specified (which is the default), the loader probes the file for a header to guess the file format.

If the data needs to be modified to fit in a lower-resolution result (e.g. converting from 32-bit to 8-bit), use the flags to control the conversion.

See also load() and Reading and Writing Image Files.
*/
func (this *QPixmap) LoadFromData(buf unsafe.Pointer /*666*/, len_ uint, format string, flags int) bool {
	var convArg2 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap12loadFromDataEPKhjPKc6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), buf, len_, convArg2, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:142
// index:0
// Public Visibility=Default Availability=Available
// [1] bool loadFromData(const uchar *, uint, const char *, Qt::ImageConversionFlags)

/*
Loads a pixmap from the len first bytes of the given binary data. Returns true if the pixmap was loaded successfully; otherwise invalidates the pixmap and returns false.

The loader attempts to read the pixmap using the specified format. If the format is not specified (which is the default), the loader probes the file for a header to guess the file format.

If the data needs to be modified to fit in a lower-resolution result (e.g. converting from 32-bit to 8-bit), use the flags to control the conversion.

See also load() and Reading and Writing Image Files.
*/
func (this *QPixmap) LoadFromDatap(buf unsafe.Pointer /*666*/, len_ uint) bool {
	// arg: 2, const char *=Pointer, =Invalid, , Invalid
	var convArg2 unsafe.Pointer
	// arg: 3, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap12loadFromDataEPKhjPKc6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), buf, len_, convArg2, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:142
// index:0
// Public Visibility=Default Availability=Available
// [1] bool loadFromData(const uchar *, uint, const char *, Qt::ImageConversionFlags)

/*
Loads a pixmap from the len first bytes of the given binary data. Returns true if the pixmap was loaded successfully; otherwise invalidates the pixmap and returns false.

The loader attempts to read the pixmap using the specified format. If the format is not specified (which is the default), the loader probes the file for a header to guess the file format.

If the data needs to be modified to fit in a lower-resolution result (e.g. converting from 32-bit to 8-bit), use the flags to control the conversion.

See also load() and Reading and Writing Image Files.
*/
func (this *QPixmap) LoadFromDatap1(buf unsafe.Pointer /*666*/, len_ uint, format string) bool {
	var convArg2 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg2)
	// arg: 3, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap12loadFromDataEPKhjPKc6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), buf, len_, convArg2, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:143
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool loadFromData(const QByteArray &, const char *, Qt::ImageConversionFlags)

/*
Loads a pixmap from the len first bytes of the given binary data. Returns true if the pixmap was loaded successfully; otherwise invalidates the pixmap and returns false.

The loader attempts to read the pixmap using the specified format. If the format is not specified (which is the default), the loader probes the file for a header to guess the file format.

If the data needs to be modified to fit in a lower-resolution result (e.g. converting from 32-bit to 8-bit), use the flags to control the conversion.

See also load() and Reading and Writing Image Files.
*/
func (this *QPixmap) LoadFromData1(data qtcore.QByteArray_ITF, format string, flags int) bool {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap12loadFromDataERK10QByteArrayPKc6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:143
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool loadFromData(const QByteArray &, const char *, Qt::ImageConversionFlags)

/*
Loads a pixmap from the len first bytes of the given binary data. Returns true if the pixmap was loaded successfully; otherwise invalidates the pixmap and returns false.

The loader attempts to read the pixmap using the specified format. If the format is not specified (which is the default), the loader probes the file for a header to guess the file format.

If the data needs to be modified to fit in a lower-resolution result (e.g. converting from 32-bit to 8-bit), use the flags to control the conversion.

See also load() and Reading and Writing Image Files.
*/
func (this *QPixmap) LoadFromData1p(data qtcore.QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	// arg: 1, const char *=Pointer, =Invalid, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap12loadFromDataERK10QByteArrayPKc6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:143
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool loadFromData(const QByteArray &, const char *, Qt::ImageConversionFlags)

/*
Loads a pixmap from the len first bytes of the given binary data. Returns true if the pixmap was loaded successfully; otherwise invalidates the pixmap and returns false.

The loader attempts to read the pixmap using the specified format. If the format is not specified (which is the default), the loader probes the file for a header to guess the file format.

If the data needs to be modified to fit in a lower-resolution result (e.g. converting from 32-bit to 8-bit), use the flags to control the conversion.

See also load() and Reading and Writing Image Files.
*/
func (this *QPixmap) LoadFromData1p1(data qtcore.QByteArray_ITF, format string) bool {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	// arg: 2, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap12loadFromDataERK10QByteArrayPKc6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:144
// index:0
// Public Visibility=Default Availability=Available
// [1] bool save(const QString &, const char *, int) const

/*
Saves the pixmap to the file with the given fileName using the specified image file format and quality factor. Returns true if successful; otherwise returns false.

The quality factor must be in the range [0,100] or -1. Specify 0 to obtain small compressed files, 100 for large uncompressed files, and -1 to use the default settings.

If format is 0, an image format will be chosen from fileName's suffix.

See also Reading and Writing Image Files.
*/
func (this *QPixmap) Save(fileName string, format string, quality int) bool {
	var tmpArg0 = qtcore.NewQString5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap4saveERK7QStringPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, quality)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:144
// index:0
// Public Visibility=Default Availability=Available
// [1] bool save(const QString &, const char *, int) const

/*
Saves the pixmap to the file with the given fileName using the specified image file format and quality factor. Returns true if successful; otherwise returns false.

The quality factor must be in the range [0,100] or -1. Specify 0 to obtain small compressed files, 100 for large uncompressed files, and -1 to use the default settings.

If format is 0, an image format will be chosen from fileName's suffix.

See also Reading and Writing Image Files.
*/
func (this *QPixmap) Savep(fileName string) bool {
	var tmpArg0 = qtcore.NewQString5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const char *=Pointer, =Invalid, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, int=Int, =Invalid, , Invalid
	quality := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap4saveERK7QStringPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, quality)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:144
// index:0
// Public Visibility=Default Availability=Available
// [1] bool save(const QString &, const char *, int) const

/*
Saves the pixmap to the file with the given fileName using the specified image file format and quality factor. Returns true if successful; otherwise returns false.

The quality factor must be in the range [0,100] or -1. Specify 0 to obtain small compressed files, 100 for large uncompressed files, and -1 to use the default settings.

If format is 0, an image format will be chosen from fileName's suffix.

See also Reading and Writing Image Files.
*/
func (this *QPixmap) Savep1(fileName string, format string) bool {
	var tmpArg0 = qtcore.NewQString5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	// arg: 2, int=Int, =Invalid, , Invalid
	quality := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap4saveERK7QStringPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, quality)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:145
// index:1
// Public Visibility=Default Availability=Available
// [1] bool save(QIODevice *, const char *, int) const

/*
Saves the pixmap to the file with the given fileName using the specified image file format and quality factor. Returns true if successful; otherwise returns false.

The quality factor must be in the range [0,100] or -1. Specify 0 to obtain small compressed files, 100 for large uncompressed files, and -1 to use the default settings.

If format is 0, an image format will be chosen from fileName's suffix.

See also Reading and Writing Image Files.
*/
func (this *QPixmap) Save1(device qtcore.QIODevice_ITF /*777 QIODevice **/, format string, quality int) bool {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap4saveEP9QIODevicePKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, quality)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:145
// index:1
// Public Visibility=Default Availability=Available
// [1] bool save(QIODevice *, const char *, int) const

/*
Saves the pixmap to the file with the given fileName using the specified image file format and quality factor. Returns true if successful; otherwise returns false.

The quality factor must be in the range [0,100] or -1. Specify 0 to obtain small compressed files, 100 for large uncompressed files, and -1 to use the default settings.

If format is 0, an image format will be chosen from fileName's suffix.

See also Reading and Writing Image Files.
*/
func (this *QPixmap) Save1p(device qtcore.QIODevice_ITF /*777 QIODevice **/) bool {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	// arg: 1, const char *=Pointer, =Invalid, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, int=Int, =Invalid, , Invalid
	quality := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap4saveEP9QIODevicePKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, quality)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:145
// index:1
// Public Visibility=Default Availability=Available
// [1] bool save(QIODevice *, const char *, int) const

/*
Saves the pixmap to the file with the given fileName using the specified image file format and quality factor. Returns true if successful; otherwise returns false.

The quality factor must be in the range [0,100] or -1. Specify 0 to obtain small compressed files, 100 for large uncompressed files, and -1 to use the default settings.

If format is 0, an image format will be chosen from fileName's suffix.

See also Reading and Writing Image Files.
*/
func (this *QPixmap) Save1p1(device qtcore.QIODevice_ITF /*777 QIODevice **/, format string) bool {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	// arg: 2, int=Int, =Invalid, , Invalid
	quality := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap4saveEP9QIODevicePKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, quality)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:147
// index:0
// Public Visibility=Default Availability=Available
// [1] bool convertFromImage(const QImage &, Qt::ImageConversionFlags)

/*
Replaces this pixmap's data with the given image using the specified flags to control the conversion. The flags argument is a bitwise-OR of the Qt::ImageConversionFlags. Passing 0 for flags sets all the default options. Returns true if the result is that this pixmap is not null.

Note: this function was part of Qt 3 support in Qt 4.6 and earlier. It has been promoted to official API status in 4.7 to support updating the pixmap's image without creating a new QPixmap as fromImage() would.

This function was introduced in  Qt 4.7.

See also fromImage().
*/
func (this *QPixmap) ConvertFromImage(img QImage_ITF, flags int) bool {
	var convArg0 unsafe.Pointer
	if img != nil && img.QImage_PTR() != nil {
		convArg0 = img.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap16convertFromImageERK6QImage6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:147
// index:0
// Public Visibility=Default Availability=Available
// [1] bool convertFromImage(const QImage &, Qt::ImageConversionFlags)

/*
Replaces this pixmap's data with the given image using the specified flags to control the conversion. The flags argument is a bitwise-OR of the Qt::ImageConversionFlags. Passing 0 for flags sets all the default options. Returns true if the result is that this pixmap is not null.

Note: this function was part of Qt 3 support in Qt 4.6 and earlier. It has been promoted to official API status in 4.7 to support updating the pixmap's image without creating a new QPixmap as fromImage() would.

This function was introduced in  Qt 4.7.

See also fromImage().
*/
func (this *QPixmap) ConvertFromImagep(img QImage_ITF) bool {
	var convArg0 unsafe.Pointer
	if img != nil && img.QImage_PTR() != nil {
		convArg0 = img.QImage_PTR().GetCthis()
	}
	// arg: 1, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap16convertFromImageERK6QImage6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:149
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QPixmap copy(int, int, int, int) const

/*
Returns a deep copy of the subset of the pixmap that is specified by the given rectangle. For more information on deep copies, see the Implicit Data Sharing documentation.

If the given rectangle is empty, the whole image is copied.

See also operator=(), QPixmap(), and Pixmap Transformations.
*/
func (this *QPixmap) Copy(x int, y int, width int, height int) *QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap4copyEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, width, height)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:150
// index:1
// Public Visibility=Default Availability=Available
// [32] QPixmap copy(const QRect &) const

/*
Returns a deep copy of the subset of the pixmap that is specified by the given rectangle. For more information on deep copies, see the Implicit Data Sharing documentation.

If the given rectangle is empty, the whole image is copied.

See also operator=(), QPixmap(), and Pixmap Transformations.
*/
func (this *QPixmap) Copy1(rect qtcore.QRect_ITF) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap4copyERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:150
// index:1
// Public Visibility=Default Availability=Available
// [32] QPixmap copy(const QRect &) const

/*
Returns a deep copy of the subset of the pixmap that is specified by the given rectangle. For more information on deep copies, see the Implicit Data Sharing documentation.

If the given rectangle is empty, the whole image is copied.

See also operator=(), QPixmap(), and Pixmap Transformations.
*/
func (this *QPixmap) Copy1p() *QPixmap /*123*/ {
	// arg: 0, const QRect &=LValueReference, QRect=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap4copyERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:152
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void scroll(int, int, int, int, int, int, QRegion *)

/*
This convenience function is equivalent to calling QPixmap::scroll(dx, dy, QRect(x, y, width, height), exposed).

This function was introduced in  Qt 4.6.

See also QWidget::scroll() and QGraphicsItem::scroll().
*/
func (this *QPixmap) Scroll(dx int, dy int, x int, y int, width int, height int, exposed QRegion_ITF /*777 QRegion **/) {
	var convArg6 unsafe.Pointer
	if exposed != nil && exposed.QRegion_PTR() != nil {
		convArg6 = exposed.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap6scrollEiiiiiiP7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy, x, y, width, height, convArg6)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:152
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void scroll(int, int, int, int, int, int, QRegion *)

/*
This convenience function is equivalent to calling QPixmap::scroll(dx, dy, QRect(x, y, width, height), exposed).

This function was introduced in  Qt 4.6.

See also QWidget::scroll() and QGraphicsItem::scroll().
*/
func (this *QPixmap) Scrollp(dx int, dy int, x int, y int, width int, height int) {
	// arg: 6, QRegion *=Pointer, QRegion=Record, , Invalid
	var convArg6 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap6scrollEiiiiiiP7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy, x, y, width, height, convArg6)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:153
// index:1
// Public Visibility=Default Availability=Available
// [-2] void scroll(int, int, const QRect &, QRegion *)

/*
This convenience function is equivalent to calling QPixmap::scroll(dx, dy, QRect(x, y, width, height), exposed).

This function was introduced in  Qt 4.6.

See also QWidget::scroll() and QGraphicsItem::scroll().
*/
func (this *QPixmap) Scroll1(dx int, dy int, rect qtcore.QRect_ITF, exposed QRegion_ITF /*777 QRegion **/) {
	var convArg2 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg2 = rect.QRect_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if exposed != nil && exposed.QRegion_PTR() != nil {
		convArg3 = exposed.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap6scrollEiiRK5QRectP7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:153
// index:1
// Public Visibility=Default Availability=Available
// [-2] void scroll(int, int, const QRect &, QRegion *)

/*
This convenience function is equivalent to calling QPixmap::scroll(dx, dy, QRect(x, y, width, height), exposed).

This function was introduced in  Qt 4.6.

See also QWidget::scroll() and QGraphicsItem::scroll().
*/
func (this *QPixmap) Scroll1p(dx int, dy int, rect qtcore.QRect_ITF) {
	var convArg2 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg2 = rect.QRect_PTR().GetCthis()
	}
	// arg: 3, QRegion *=Pointer, QRegion=Record, , Invalid
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap6scrollEiiRK5QRectP7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:158
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 cacheKey() const

/*
Returns a number that identifies this QPixmap. Distinct QPixmap objects can only have the same cache key if they refer to the same contents.

The cacheKey() will change when the pixmap is altered.
*/
func (this *QPixmap) CacheKey() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap8cacheKeyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtGui/qpixmap.h:160
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDetached() const

/*

 */
func (this *QPixmap) IsDetached() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap10isDetachedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:161
// index:0
// Public Visibility=Default Availability=Available
// [-2] void detach()

/*
Detaches the pixmap from shared pixmap data.

A pixmap is automatically detached by Qt whenever its contents are about to change. This is done in almost all QPixmap member functions that modify the pixmap (fill(), fromImage(), load(), etc.), and in QPainter::begin() on a pixmap.

There are two exceptions in which detach() must be called explicitly, that is when calling the handle() or the x11PictureHandle() function (only available on X11). Otherwise, any modifications done using system calls, will be performed on the shared data.

The detach() function returns immediately if there is just a single reference or if the pixmap has not been initialized yet.
*/
func (this *QPixmap) Detach() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap6detachEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:163
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isQBitmap() const

/*
Returns true if this is a QBitmap; otherwise returns false.
*/
func (this *QPixmap) IsQBitmap() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap9isQBitmapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:165
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPaintEngine * paintEngine() const

/*

 */
func (this *QPixmap) PaintEngine() *QPaintEngine /*777 QPaintEngine **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap11paintEngineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPaintEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qpixmap.h:167
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!() const

/*

 */
func (this *QPixmap) Operator_not() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmapntEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:175
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int metric(QPaintDevice::PaintDeviceMetric) const

/*

 */
func (this *QPixmap) Metric(arg0 int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap6metricEN12QPaintDevice17PaintDeviceMetricE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpixmap.h:176
// index:0
// Protected static Visibility=Default Availability=Available
// [32] QPixmap fromImageInPlace(QImage &, Qt::ImageConversionFlags)

/*

 */
func (this *QPixmap) FromImageInPlace(image QImage_ITF, flags int) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg0 = image.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap16fromImageInPlaceER6QImage6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}
func QPixmap_FromImageInPlace(image QImage_ITF, flags int) *QPixmap /*123*/ {
	var nilthis *QPixmap
	rv := nilthis.FromImageInPlace(image, flags)
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:176
// index:0
// Protected static Visibility=Default Availability=Available
// [32] QPixmap fromImageInPlace(QImage &, Qt::ImageConversionFlags)

/*

 */
func (this *QPixmap) FromImageInPlacep(image QImage_ITF) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg0 = image.QImage_PTR().GetCthis()
	}
	// arg: 1, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap16fromImageInPlaceER6QImage6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
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
