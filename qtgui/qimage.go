package qtgui

// /usr/include/qt/QtGui/qimage.h
// #include <qimage.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 24
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

// int metric(enum QPaintDevice::PaintDeviceMetric)
func (this *QImage) InheritMetric(f func(metric int) int) {
	qtrt.SetAllInheritCallback(this, "metric", f)
}

// QImage mirrored_helper(_Bool, _Bool)
func (this *QImage) InheritMirrored_helper(f func(horizontal bool, vertical bool) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "mirrored_helper", f)
}

// QImage rgbSwapped_helper()
func (this *QImage) InheritRgbSwapped_helper(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "rgbSwapped_helper", f)
}

// void mirrored_inplace(_Bool, _Bool)
func (this *QImage) InheritMirrored_inplace(f func(horizontal bool, vertical bool) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mirrored_inplace", f)
}

// void rgbSwapped_inplace()
func (this *QImage) InheritRgbSwapped_inplace(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "rgbSwapped_inplace", f)
}

// QImage convertToFormat_helper(enum QImage::Format, Qt::ImageConversionFlags)
func (this *QImage) InheritConvertToFormat_helper(f func(format int, flags int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "convertToFormat_helper", f)
}

// bool convertToFormat_inplace(enum QImage::Format, Qt::ImageConversionFlags)
func (this *QImage) InheritConvertToFormat_inplace(f func(format int, flags int) bool) {
	qtrt.SetAllInheritCallback(this, "convertToFormat_inplace", f)
}

// QImage smoothScaled(int, int)
func (this *QImage) InheritSmoothScaled(f func(w int, h int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "smoothScaled", f)
}

type QImage struct {
	*QPaintDevice
}
type QImage_ITF interface {
	QPaintDevice_ITF
	QImage_PTR() *QImage
}

func (ptr *QImage) QImage_PTR() *QImage { return ptr }

func (this *QImage) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QPaintDevice.GetCthis()
	}
}
func (this *QImage) SetCthis(cthis unsafe.Pointer) {
	this.QPaintDevice = NewQPaintDeviceFromPointer(cthis)
}
func NewQImageFromPointer(cthis unsafe.Pointer) *QImage {
	bcthis0 := NewQPaintDeviceFromPointer(cthis)
	return &QImage{bcthis0}
}
func (*QImage) NewFromPointer(cthis unsafe.Pointer) *QImage {
	return NewQImageFromPointer(cthis)
}

// /usr/include/qt/QtGui/qimage.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QImage()
func NewQImage() *QImage {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImage)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:137
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QImage(const QSize &, enum QImage::Format)
func NewQImage_1(size qtcore.QSize_ITF, format int) *QImage {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageC2ERK5QSizeNS_6FormatE", qtrt.FFI_TYPE_POINTER, convArg0, format)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImage)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:138
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QImage(int, int, enum QImage::Format)
func NewQImage_2(width int, height int, format int) *QImage {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageC2EiiNS_6FormatE", qtrt.FFI_TYPE_POINTER, width, height, format)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImage)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:145
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QImage(const char *const *)
func NewQImage_3(xpm []string) *QImage {
	var convArg0 = qtrt.StringSliceToCCharPP(xpm)
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageC2EPKPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImage)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:147
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QImage(const QString &, const char *)
func NewQImage_4(fileName string, format string) *QImage {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageC2ERK7QStringPKc", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImage)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:147
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QImage(const QString &, const char *)
func NewQImage_4_(fileName string) *QImage {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const char *=Pointer, =Invalid,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageC2ERK7QStringPKc", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImage)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:155
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QImage()
func DeleteQImage(this *QImage) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qimage.h:157
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage & operator=(const QImage &)
func (this *QImage) Operator_equal(arg0 QImage_ITF) *QImage {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QImage_PTR() != nil {
		convArg0 = arg0.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:159
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QImage & operator=(QImage &&)
func (this *QImage) Operator_equal_1(other unsafe.Pointer /*333*/) *QImage {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:162
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QImage &)
func (this *QImage) Swap(other QImage_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QImage_PTR() != nil {
		convArg0 = other.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:165
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const
func (this *QImage) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:167
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int devType() const
func (this *QImage) DevType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage7devTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimage.h:169
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QImage &) const
func (this *QImage) Operator_equal_equal(arg0 QImage_ITF) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QImage_PTR() != nil {
		convArg0 = arg0.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImageeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:170
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QImage &) const
func (this *QImage) Operator_not_equal(arg0 QImage_ITF) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QImage_PTR() != nil {
		convArg0 = arg0.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImageneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:172
// index:0
// Public Visibility=Default Availability=Available
// [-2] void detach()
func (this *QImage) Detach() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage6detachEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:173
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDetached() const
func (this *QImage) IsDetached() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage10isDetachedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:175
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage copy(const QRect &) const
func (this *QImage) Copy(rect qtcore.QRect_ITF) *QImage /*123*/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4copyERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:175
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage copy(const QRect &) const
func (this *QImage) Copy__() *QImage /*123*/ {
	// arg: 0, const QRect &=LValueReference, QRect=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4copyERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:176
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QImage copy(int, int, int, int) const
func (this *QImage) Copy_1(x int, y int, w int, h int) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4copyEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:179
// index:0
// Public Visibility=Default Availability=Available
// [4] QImage::Format format() const
func (this *QImage) Format() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qimage.h:182
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QImage convertToFormat(enum QImage::Format, Qt::ImageConversionFlags) const
func (this *QImage) ConvertToFormat(f int, flags int) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNKR6QImage15convertToFormatENS_6FormatE6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:182
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QImage convertToFormat(enum QImage::Format, Qt::ImageConversionFlags) const
func (this *QImage) ConvertToFormat__(f int) *QImage /*123*/ {
	// arg: 1, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNKR6QImage15convertToFormatENS_6FormatE6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:184
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QImage convertToFormat(enum QImage::Format, Qt::ImageConversionFlags)
func (this *QImage) ConvertToFormat_1(f int, flags int) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNO6QImage15convertToFormatENS_6FormatE6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:184
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QImage convertToFormat(enum QImage::Format, Qt::ImageConversionFlags)
func (this *QImage) ConvertToFormat_1_(f int) *QImage /*123*/ {
	// arg: 1, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNO6QImage15convertToFormatENS_6FormatE6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:195
// index:0
// Public Visibility=Default Availability=Available
// [1] bool reinterpretAsFormat(enum QImage::Format)
func (this *QImage) ReinterpretAsFormat(f int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage19reinterpretAsFormatENS_6FormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:197
// index:0
// Public Visibility=Default Availability=Available
// [4] int width() const
func (this *QImage) Width() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage5widthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimage.h:198
// index:0
// Public Visibility=Default Availability=Available
// [4] int height() const
func (this *QImage) Height() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage6heightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimage.h:199
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize size() const
func (this *QImage) Size() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:200
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect rect() const
func (this *QImage) Rect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4rectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:202
// index:0
// Public Visibility=Default Availability=Available
// [4] int depth() const
func (this *QImage) Depth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage5depthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimage.h:203
// index:0
// Public Visibility=Default Availability=Available
// [4] int colorCount() const
func (this *QImage) ColorCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage10colorCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimage.h:204
// index:0
// Public Visibility=Default Availability=Available
// [4] int bitPlaneCount() const
func (this *QImage) BitPlaneCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage13bitPlaneCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimage.h:206
// index:0
// Public Visibility=Default Availability=Available
// [4] QRgb color(int) const
func (this *QImage) Color(i int) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage5colorEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qimage.h:207
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColor(int, QRgb)
func (this *QImage) SetColor(i int, c uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage8setColorEij", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i, c)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:208
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColorCount(int)
func (this *QImage) SetColorCount(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage13setColorCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:210
// index:0
// Public Visibility=Default Availability=Available
// [1] bool allGray() const
func (this *QImage) AllGray() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage7allGrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:211
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isGrayscale() const
func (this *QImage) IsGrayscale() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage11isGrayscaleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:213
// index:0
// Public Visibility=Default Availability=Available
// [8] uchar * bits()
func (this *QImage) Bits() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage4bitsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qimage.h:214
// index:1
// Public Visibility=Default Availability=Available
// [8] const uchar * bits() const
func (this *QImage) Bits_1() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4bitsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qimage.h:215
// index:0
// Public Visibility=Default Availability=Available
// [8] const uchar * constBits() const
func (this *QImage) ConstBits() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage9constBitsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qimage.h:218
// index:0
// Public Visibility=Default Availability=Available
// [4] int byteCount() const
func (this *QImage) ByteCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage9byteCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimage.h:220
// index:0
// Public Visibility=Default Availability=Available
// [8] qsizetype sizeInBytes() const
func (this *QImage) SizeInBytes() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage11sizeInBytesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtGui/qimage.h:222
// index:0
// Public Visibility=Default Availability=Available
// [8] uchar * scanLine(int)
func (this *QImage) ScanLine(arg0 int) unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage8scanLineEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qimage.h:223
// index:1
// Public Visibility=Default Availability=Available
// [8] const uchar * scanLine(int) const
func (this *QImage) ScanLine_1(arg0 int) unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage8scanLineEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qimage.h:224
// index:0
// Public Visibility=Default Availability=Available
// [8] const uchar * constScanLine(int) const
func (this *QImage) ConstScanLine(arg0 int) unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage13constScanLineEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qimage.h:225
// index:0
// Public Visibility=Default Availability=Available
// [4] int bytesPerLine() const
func (this *QImage) BytesPerLine() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage12bytesPerLineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimage.h:227
// index:0
// Public Visibility=Default Availability=Available
// [1] bool valid(int, int) const
func (this *QImage) Valid(x int, y int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage5validEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:228
// index:1
// Public Visibility=Default Availability=Available
// [1] bool valid(const QPoint &) const
func (this *QImage) Valid_1(pt qtcore.QPoint_ITF) bool {
	var convArg0 unsafe.Pointer
	if pt != nil && pt.QPoint_PTR() != nil {
		convArg0 = pt.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage5validERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:230
// index:0
// Public Visibility=Default Availability=Available
// [4] int pixelIndex(int, int) const
func (this *QImage) PixelIndex(x int, y int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage10pixelIndexEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimage.h:231
// index:1
// Public Visibility=Default Availability=Available
// [4] int pixelIndex(const QPoint &) const
func (this *QImage) PixelIndex_1(pt qtcore.QPoint_ITF) int {
	var convArg0 unsafe.Pointer
	if pt != nil && pt.QPoint_PTR() != nil {
		convArg0 = pt.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage10pixelIndexERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimage.h:233
// index:0
// Public Visibility=Default Availability=Available
// [4] QRgb pixel(int, int) const
func (this *QImage) Pixel(x int, y int) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage5pixelEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qimage.h:234
// index:1
// Public Visibility=Default Availability=Available
// [4] QRgb pixel(const QPoint &) const
func (this *QImage) Pixel_1(pt qtcore.QPoint_ITF) uint {
	var convArg0 unsafe.Pointer
	if pt != nil && pt.QPoint_PTR() != nil {
		convArg0 = pt.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage5pixelERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qimage.h:236
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPixel(int, int, uint)
func (this *QImage) SetPixel(x int, y int, index_or_rgb uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage8setPixelEiij", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, index_or_rgb)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:237
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setPixel(const QPoint &, uint)
func (this *QImage) SetPixel_1(pt qtcore.QPoint_ITF, index_or_rgb uint) {
	var convArg0 unsafe.Pointer
	if pt != nil && pt.QPoint_PTR() != nil {
		convArg0 = pt.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage8setPixelERK6QPointj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, index_or_rgb)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:239
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor pixelColor(int, int) const
func (this *QImage) PixelColor(x int, y int) *QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage10pixelColorEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:240
// index:1
// Public Visibility=Default Availability=Available
// [16] QColor pixelColor(const QPoint &) const
func (this *QImage) PixelColor_1(pt qtcore.QPoint_ITF) *QColor /*123*/ {
	var convArg0 unsafe.Pointer
	if pt != nil && pt.QPoint_PTR() != nil {
		convArg0 = pt.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage10pixelColorERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:242
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPixelColor(int, int, const QColor &)
func (this *QImage) SetPixelColor(x int, y int, c QColor_ITF) {
	var convArg2 unsafe.Pointer
	if c != nil && c.QColor_PTR() != nil {
		convArg2 = c.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage13setPixelColorEiiRK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:243
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setPixelColor(const QPoint &, const QColor &)
func (this *QImage) SetPixelColor_1(pt qtcore.QPoint_ITF, c QColor_ITF) {
	var convArg0 unsafe.Pointer
	if pt != nil && pt.QPoint_PTR() != nil {
		convArg0 = pt.QPoint_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if c != nil && c.QColor_PTR() != nil {
		convArg1 = c.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage13setPixelColorERK6QPointRK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:252
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal devicePixelRatio() const
func (this *QImage) DevicePixelRatio() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage16devicePixelRatioEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qimage.h:253
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDevicePixelRatio(qreal)
func (this *QImage) SetDevicePixelRatio(scaleFactor float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage19setDevicePixelRatioEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), scaleFactor)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:255
// index:0
// Public Visibility=Default Availability=Available
// [-2] void fill(uint)
func (this *QImage) Fill(pixel uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage4fillEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pixel)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:256
// index:1
// Public Visibility=Default Availability=Available
// [-2] void fill(const QColor &)
func (this *QImage) Fill_1(color QColor_ITF) {
	var convArg0 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg0 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage4fillERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:257
// index:2
// Public Visibility=Default Availability=Available
// [-2] void fill(Qt::GlobalColor)
func (this *QImage) Fill_2(color int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage4fillEN2Qt11GlobalColorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), color)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:260
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasAlphaChannel() const
func (this *QImage) HasAlphaChannel() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage15hasAlphaChannelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:261
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlphaChannel(const QImage &)
func (this *QImage) SetAlphaChannel(alphaChannel QImage_ITF) {
	var convArg0 unsafe.Pointer
	if alphaChannel != nil && alphaChannel.QImage_PTR() != nil {
		convArg0 = alphaChannel.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage15setAlphaChannelERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:262
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage alphaChannel() const
func (this *QImage) AlphaChannel() *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage12alphaChannelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:263
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage createAlphaMask(Qt::ImageConversionFlags) const
func (this *QImage) CreateAlphaMask(flags int) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage15createAlphaMaskE6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:263
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage createAlphaMask(Qt::ImageConversionFlags) const
func (this *QImage) CreateAlphaMask__() *QImage /*123*/ {
	// arg: 0, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage15createAlphaMaskE6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:265
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage createHeuristicMask(_Bool) const
func (this *QImage) CreateHeuristicMask(clipTight bool) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage19createHeuristicMaskEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), clipTight)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:265
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage createHeuristicMask(_Bool) const
func (this *QImage) CreateHeuristicMask__() *QImage /*123*/ {
	// arg: 0, bool=Bool, =Invalid,
	clipTight := true
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage19createHeuristicMaskEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), clipTight)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:267
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage createMaskFromColor(QRgb, Qt::MaskMode) const
func (this *QImage) CreateMaskFromColor(color uint, mode int) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage19createMaskFromColorEjN2Qt8MaskModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), color, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:267
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage createMaskFromColor(QRgb, Qt::MaskMode) const
func (this *QImage) CreateMaskFromColor__(color uint) *QImage /*123*/ {
	// arg: 1, Qt::MaskMode=Elaborated, Qt::MaskMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage19createMaskFromColorEjN2Qt8MaskModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), color, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:269
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QImage scaled(int, int, Qt::AspectRatioMode, Qt::TransformationMode) const
func (this *QImage) Scaled(w int, h int, aspectMode int, mode int) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage6scaledEiiN2Qt15AspectRatioModeENS0_18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h, aspectMode, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:269
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QImage scaled(int, int, Qt::AspectRatioMode, Qt::TransformationMode) const
func (this *QImage) Scaled__(w int, h int) *QImage /*123*/ {
	// arg: 2, Qt::AspectRatioMode=Elaborated, Qt::AspectRatioMode=Enum,
	aspectMode := 0
	// arg: 3, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage6scaledEiiN2Qt15AspectRatioModeENS0_18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h, aspectMode, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:269
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QImage scaled(int, int, Qt::AspectRatioMode, Qt::TransformationMode) const
func (this *QImage) Scaled__1(w int, h int, aspectMode int) *QImage /*123*/ {
	// arg: 3, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage6scaledEiiN2Qt15AspectRatioModeENS0_18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h, aspectMode, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:272
// index:1
// Public Visibility=Default Availability=Available
// [32] QImage scaled(const QSize &, Qt::AspectRatioMode, Qt::TransformationMode) const
func (this *QImage) Scaled_1(s qtcore.QSize_ITF, aspectMode int, mode int) *QImage /*123*/ {
	var convArg0 unsafe.Pointer
	if s != nil && s.QSize_PTR() != nil {
		convArg0 = s.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage6scaledERK5QSizeN2Qt15AspectRatioModeENS3_18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, aspectMode, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:272
// index:1
// Public Visibility=Default Availability=Available
// [32] QImage scaled(const QSize &, Qt::AspectRatioMode, Qt::TransformationMode) const
func (this *QImage) Scaled_1_(s qtcore.QSize_ITF) *QImage /*123*/ {
	var convArg0 unsafe.Pointer
	if s != nil && s.QSize_PTR() != nil {
		convArg0 = s.QSize_PTR().GetCthis()
	}
	// arg: 1, Qt::AspectRatioMode=Elaborated, Qt::AspectRatioMode=Enum,
	aspectMode := 0
	// arg: 2, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage6scaledERK5QSizeN2Qt15AspectRatioModeENS3_18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, aspectMode, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:272
// index:1
// Public Visibility=Default Availability=Available
// [32] QImage scaled(const QSize &, Qt::AspectRatioMode, Qt::TransformationMode) const
func (this *QImage) Scaled_1_1(s qtcore.QSize_ITF, aspectMode int) *QImage /*123*/ {
	var convArg0 unsafe.Pointer
	if s != nil && s.QSize_PTR() != nil {
		convArg0 = s.QSize_PTR().GetCthis()
	}
	// arg: 2, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage6scaledERK5QSizeN2Qt15AspectRatioModeENS3_18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, aspectMode, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:274
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage scaledToWidth(int, Qt::TransformationMode) const
func (this *QImage) ScaledToWidth(w int, mode int) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage13scaledToWidthEiN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:274
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage scaledToWidth(int, Qt::TransformationMode) const
func (this *QImage) ScaledToWidth__(w int) *QImage /*123*/ {
	// arg: 1, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage13scaledToWidthEiN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:275
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage scaledToHeight(int, Qt::TransformationMode) const
func (this *QImage) ScaledToHeight(h int, mode int) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage14scaledToHeightEiN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), h, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:275
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage scaledToHeight(int, Qt::TransformationMode) const
func (this *QImage) ScaledToHeight__(h int) *QImage /*123*/ {
	// arg: 1, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage14scaledToHeightEiN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), h, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:276
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage transformed(const QMatrix &, Qt::TransformationMode) const
func (this *QImage) Transformed(matrix QMatrix_ITF, mode int) *QImage /*123*/ {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QMatrix_PTR() != nil {
		convArg0 = matrix.QMatrix_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage11transformedERK7QMatrixN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:276
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage transformed(const QMatrix &, Qt::TransformationMode) const
func (this *QImage) Transformed__(matrix QMatrix_ITF) *QImage /*123*/ {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QMatrix_PTR() != nil {
		convArg0 = matrix.QMatrix_PTR().GetCthis()
	}
	// arg: 1, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage11transformedERK7QMatrixN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:278
// index:1
// Public Visibility=Default Availability=Available
// [32] QImage transformed(const QTransform &, Qt::TransformationMode) const
func (this *QImage) Transformed_1(matrix QTransform_ITF, mode int) *QImage /*123*/ {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QTransform_PTR() != nil {
		convArg0 = matrix.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage11transformedERK10QTransformN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:278
// index:1
// Public Visibility=Default Availability=Available
// [32] QImage transformed(const QTransform &, Qt::TransformationMode) const
func (this *QImage) Transformed_1_(matrix QTransform_ITF) *QImage /*123*/ {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QTransform_PTR() != nil {
		convArg0 = matrix.QTransform_PTR().GetCthis()
	}
	// arg: 1, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage11transformedERK10QTransformN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:277
// index:0
// Public static Visibility=Default Availability=Available
// [48] QMatrix trueMatrix(const QMatrix &, int, int)
func (this *QImage) TrueMatrix(arg0 QMatrix_ITF, w int, h int) *QMatrix /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMatrix_PTR() != nil {
		convArg0 = arg0.QMatrix_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage10trueMatrixERK7QMatrixii", qtrt.FFI_TYPE_POINTER, convArg0, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix)
	return rv2
}
func QImage_TrueMatrix(arg0 QMatrix_ITF, w int, h int) *QMatrix /*123*/ {
	var nilthis *QImage
	rv := nilthis.TrueMatrix(arg0, w, h)
	return rv
}

// /usr/include/qt/QtGui/qimage.h:279
// index:1
// Public static Visibility=Default Availability=Available
// [88] QTransform trueMatrix(const QTransform &, int, int)
func (this *QImage) TrueMatrix_1(arg0 QTransform_ITF, w int, h int) *QTransform /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTransform_PTR() != nil {
		convArg0 = arg0.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage10trueMatrixERK10QTransformii", qtrt.FFI_TYPE_POINTER, convArg0, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}
func QImage_TrueMatrix_1(arg0 QTransform_ITF, w int, h int) *QTransform /*123*/ {
	var nilthis *QImage
	rv := nilthis.TrueMatrix_1(arg0, w, h)
	return rv
}

// /usr/include/qt/QtGui/qimage.h:281
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QImage mirrored(_Bool, _Bool) const
func (this *QImage) Mirrored(horizontally bool, vertically bool) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNKR6QImage8mirroredEbb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), horizontally, vertically)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:281
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QImage mirrored(_Bool, _Bool) const
func (this *QImage) Mirrored__() *QImage /*123*/ {
	// arg: 0, bool=Bool, =Invalid,
	horizontally := false
	// arg: 1, bool=Bool, =Invalid,
	vertically := true
	rv, err := qtrt.InvokeQtFunc6("_ZNKR6QImage8mirroredEbb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), horizontally, vertically)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:281
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QImage mirrored(_Bool, _Bool) const
func (this *QImage) Mirrored__1(horizontally bool) *QImage /*123*/ {
	// arg: 1, bool=Bool, =Invalid,
	vertically := true
	rv, err := qtrt.InvokeQtFunc6("_ZNKR6QImage8mirroredEbb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), horizontally, vertically)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:283
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QImage && mirrored(_Bool, _Bool)
func (this *QImage) Mirrored_1(horizontally bool, vertically bool) unsafe.Pointer /*333*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNO6QImage8mirroredEbb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), horizontally, vertically)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv)) //777
}

// /usr/include/qt/QtGui/qimage.h:283
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QImage && mirrored(_Bool, _Bool)
func (this *QImage) Mirrored_1_() unsafe.Pointer /*333*/ {
	// arg: 0, bool=Bool, =Invalid,
	horizontally := false
	// arg: 1, bool=Bool, =Invalid,
	vertically := true
	rv, err := qtrt.InvokeQtFunc6("_ZNO6QImage8mirroredEbb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), horizontally, vertically)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv)) //777
}

// /usr/include/qt/QtGui/qimage.h:283
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QImage && mirrored(_Bool, _Bool)
func (this *QImage) Mirrored_1_1(horizontally bool) unsafe.Pointer /*333*/ {
	// arg: 1, bool=Bool, =Invalid,
	vertically := true
	rv, err := qtrt.InvokeQtFunc6("_ZNO6QImage8mirroredEbb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), horizontally, vertically)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv)) //777
}

// /usr/include/qt/QtGui/qimage.h:285
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QImage rgbSwapped() const
func (this *QImage) RgbSwapped() *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNKR6QImage10rgbSwappedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:287
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QImage && rgbSwapped()
func (this *QImage) RgbSwapped_1() unsafe.Pointer /*333*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNO6QImage10rgbSwappedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv)) //777
}

// /usr/include/qt/QtGui/qimage.h:293
// index:0
// Public Visibility=Default Availability=Available
// [-2] void invertPixels(enum QImage::InvertMode)
func (this *QImage) InvertPixels(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage12invertPixelsENS_10InvertModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:293
// index:0
// Public Visibility=Default Availability=Available
// [-2] void invertPixels(enum QImage::InvertMode)
func (this *QImage) InvertPixels__() {
	// arg: 0, QImage::InvertMode=Enum, QImage::InvertMode=Enum,
	arg0 := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage12invertPixelsENS_10InvertModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:296
// index:0
// Public Visibility=Default Availability=Available
// [1] bool load(QIODevice *, const char *)
func (this *QImage) Load(device qtcore.QIODevice_ITF /*777 QIODevice **/, format string) bool {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage4loadEP9QIODevicePKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:297
// index:1
// Public Visibility=Default Availability=Available
// [1] bool load(const QString &, const char *)
func (this *QImage) Load_1(fileName string, format string) bool {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage4loadERK7QStringPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:297
// index:1
// Public Visibility=Default Availability=Available
// [1] bool load(const QString &, const char *)
func (this *QImage) Load_1_(fileName string) bool {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const char *=Pointer, =Invalid,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage4loadERK7QStringPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:298
// index:0
// Public Visibility=Default Availability=Available
// [1] bool loadFromData(const uchar *, int, const char *)
func (this *QImage) LoadFromData(buf unsafe.Pointer /*666*/, len_ int, format string) bool {
	var convArg2 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage12loadFromDataEPKhiPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), buf, len_, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:298
// index:0
// Public Visibility=Default Availability=Available
// [1] bool loadFromData(const uchar *, int, const char *)
func (this *QImage) LoadFromData__(buf unsafe.Pointer /*666*/, len_ int) bool {
	// arg: 2, const char *=Pointer, =Invalid,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage12loadFromDataEPKhiPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), buf, len_, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:299
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool loadFromData(const QByteArray &, const char *)
func (this *QImage) LoadFromData_1(data qtcore.QByteArray_ITF, aformat string) bool {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(aformat)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage12loadFromDataERK10QByteArrayPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:299
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool loadFromData(const QByteArray &, const char *)
func (this *QImage) LoadFromData_1_(data qtcore.QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	// arg: 1, const char *=Pointer, =Invalid,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage12loadFromDataERK10QByteArrayPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:302
// index:0
// Public Visibility=Default Availability=Available
// [1] bool save(const QString &, const char *, int) const
func (this *QImage) Save(fileName string, format string, quality int) bool {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4saveERK7QStringPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, quality)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:302
// index:0
// Public Visibility=Default Availability=Available
// [1] bool save(const QString &, const char *, int) const
func (this *QImage) Save__(fileName string) bool {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const char *=Pointer, =Invalid,
	var convArg1 unsafe.Pointer
	// arg: 2, int=Int, =Invalid,
	quality := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4saveERK7QStringPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, quality)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:302
// index:0
// Public Visibility=Default Availability=Available
// [1] bool save(const QString &, const char *, int) const
func (this *QImage) Save__1(fileName string, format string) bool {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	// arg: 2, int=Int, =Invalid,
	quality := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4saveERK7QStringPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, quality)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:303
// index:1
// Public Visibility=Default Availability=Available
// [1] bool save(QIODevice *, const char *, int) const
func (this *QImage) Save_1(device qtcore.QIODevice_ITF /*777 QIODevice **/, format string, quality int) bool {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4saveEP9QIODevicePKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, quality)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:303
// index:1
// Public Visibility=Default Availability=Available
// [1] bool save(QIODevice *, const char *, int) const
func (this *QImage) Save_1_(device qtcore.QIODevice_ITF /*777 QIODevice **/) bool {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	// arg: 1, const char *=Pointer, =Invalid,
	var convArg1 unsafe.Pointer
	// arg: 2, int=Int, =Invalid,
	quality := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4saveEP9QIODevicePKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, quality)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:303
// index:1
// Public Visibility=Default Availability=Available
// [1] bool save(QIODevice *, const char *, int) const
func (this *QImage) Save_1_1(device qtcore.QIODevice_ITF /*777 QIODevice **/, format string) bool {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	// arg: 2, int=Int, =Invalid,
	quality := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4saveEP9QIODevicePKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, quality)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:305
// index:0
// Public static Visibility=Default Availability=Available
// [32] QImage fromData(const uchar *, int, const char *)
func (this *QImage) FromData(data unsafe.Pointer /*666*/, size int, format string) *QImage /*123*/ {
	var convArg2 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage8fromDataEPKhiPKc", qtrt.FFI_TYPE_POINTER, data, size, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}
func QImage_FromData(data unsafe.Pointer /*666*/, size int, format string) *QImage /*123*/ {
	var nilthis *QImage
	rv := nilthis.FromData(data, size, format)
	return rv
}

// /usr/include/qt/QtGui/qimage.h:305
// index:0
// Public static Visibility=Default Availability=Available
// [32] QImage fromData(const uchar *, int, const char *)
func (this *QImage) FromData__(data unsafe.Pointer /*666*/, size int) *QImage /*123*/ {
	// arg: 2, const char *=Pointer, =Invalid,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage8fromDataEPKhiPKc", qtrt.FFI_TYPE_POINTER, data, size, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:306
// index:1
// Public static inline Visibility=Default Availability=Available
// [32] QImage fromData(const QByteArray &, const char *)
func (this *QImage) FromData_1(data qtcore.QByteArray_ITF, format string) *QImage /*123*/ {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage8fromDataERK10QByteArrayPKc", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}
func QImage_FromData_1(data qtcore.QByteArray_ITF, format string) *QImage /*123*/ {
	var nilthis *QImage
	rv := nilthis.FromData_1(data, format)
	return rv
}

// /usr/include/qt/QtGui/qimage.h:306
// index:1
// Public static inline Visibility=Default Availability=Available
// [32] QImage fromData(const QByteArray &, const char *)
func (this *QImage) FromData_1_(data qtcore.QByteArray_ITF) *QImage /*123*/ {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	// arg: 1, const char *=Pointer, =Invalid,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage8fromDataERK10QByteArrayPKc", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:312
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 cacheKey() const
func (this *QImage) CacheKey() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage8cacheKeyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtGui/qimage.h:314
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPaintEngine * paintEngine() const
func (this *QImage) PaintEngine() *QPaintEngine /*777 QPaintEngine **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage11paintEngineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPaintEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qimage.h:317
// index:0
// Public Visibility=Default Availability=Available
// [4] int dotsPerMeterX() const
func (this *QImage) DotsPerMeterX() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage13dotsPerMeterXEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimage.h:318
// index:0
// Public Visibility=Default Availability=Available
// [4] int dotsPerMeterY() const
func (this *QImage) DotsPerMeterY() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage13dotsPerMeterYEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimage.h:319
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDotsPerMeterX(int)
func (this *QImage) SetDotsPerMeterX(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage16setDotsPerMeterXEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:320
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDotsPerMeterY(int)
func (this *QImage) SetDotsPerMeterY(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage16setDotsPerMeterYEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:321
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint offset() const
func (this *QImage) Offset() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage6offsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:322
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOffset(const QPoint &)
func (this *QImage) SetOffset(arg0 qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPoint_PTR() != nil {
		convArg0 = arg0.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage9setOffsetERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:324
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList textKeys() const
func (this *QImage) TextKeys() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage8textKeysEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:325
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text(const QString &) const
func (this *QImage) Text(key string) string {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4textERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qimage.h:325
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text(const QString &) const
func (this *QImage) Text__() string {
	// arg: 0, const QString &=LValueReference, QString=Record,
	var convArg0 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4textERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qimage.h:326
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &, const QString &)
func (this *QImage) SetText(key string, value string) {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(value)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage7setTextERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:328
// index:0
// Public Visibility=Default Availability=Available
// [8] QPixelFormat pixelFormat() const
func (this *QImage) PixelFormat() *QPixelFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage11pixelFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixelFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixelFormat)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:329
// index:0
// Public static Visibility=Default Availability=Available
// [8] QPixelFormat toPixelFormat(QImage::Format)
func (this *QImage) ToPixelFormat(format int) *QPixelFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage13toPixelFormatENS_6FormatE", qtrt.FFI_TYPE_POINTER, format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixelFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixelFormat)
	return rv2
}
func QImage_ToPixelFormat(format int) *QPixelFormat /*123*/ {
	var nilthis *QImage
	rv := nilthis.ToPixelFormat(format)
	return rv
}

// /usr/include/qt/QtGui/qimage.h:330
// index:0
// Public static Visibility=Default Availability=Available
// [4] QImage::Format toImageFormat(QPixelFormat)
func (this *QImage) ToImageFormat(format QPixelFormat_ITF /*123*/) int {
	var convArg0 unsafe.Pointer
	if format != nil && format.QPixelFormat_PTR() != nil {
		convArg0 = format.QPixelFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage13toImageFormatE12QPixelFormat", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QImage_ToImageFormat(format QPixelFormat_ITF /*123*/) int {
	var nilthis *QImage
	rv := nilthis.ToImageFormat(format)
	return rv
}

// /usr/include/qt/QtGui/qimage.h:352
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int metric(enum QPaintDevice::PaintDeviceMetric) const
func (this *QImage) Metric(metric int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage6metricEN12QPaintDevice17PaintDeviceMetricE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), metric)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimage.h:353
// index:0
// Protected Visibility=Default Availability=Available
// [32] QImage mirrored_helper(_Bool, _Bool) const
func (this *QImage) Mirrored_helper(horizontal bool, vertical bool) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage15mirrored_helperEbb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), horizontal, vertical)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:354
// index:0
// Protected Visibility=Default Availability=Available
// [32] QImage rgbSwapped_helper() const
func (this *QImage) RgbSwapped_helper() *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage17rgbSwapped_helperEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:355
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void mirrored_inplace(_Bool, _Bool)
func (this *QImage) Mirrored_inplace(horizontal bool, vertical bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage16mirrored_inplaceEbb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), horizontal, vertical)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:356
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void rgbSwapped_inplace()
func (this *QImage) RgbSwapped_inplace() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage18rgbSwapped_inplaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:357
// index:0
// Protected Visibility=Default Availability=Available
// [32] QImage convertToFormat_helper(enum QImage::Format, Qt::ImageConversionFlags) const
func (this *QImage) ConvertToFormat_helper(format int, flags int) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage22convertToFormat_helperENS_6FormatE6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:358
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool convertToFormat_inplace(enum QImage::Format, Qt::ImageConversionFlags)
func (this *QImage) ConvertToFormat_inplace(format int, flags int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage23convertToFormat_inplaceENS_6FormatE6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:359
// index:0
// Protected Visibility=Default Availability=Available
// [32] QImage smoothScaled(int, int) const
func (this *QImage) SmoothScaled(w int, h int) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage12smoothScaledEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

type QImage__InvertMode = int

const QImage__InvertRgb QImage__InvertMode = 0
const QImage__InvertRgba QImage__InvertMode = 1

type QImage__Format = int

const QImage__Format_Invalid QImage__Format = 0
const QImage__Format_Mono QImage__Format = 1
const QImage__Format_MonoLSB QImage__Format = 2
const QImage__Format_Indexed8 QImage__Format = 3
const QImage__Format_RGB32 QImage__Format = 4
const QImage__Format_ARGB32 QImage__Format = 5
const QImage__Format_ARGB32_Premultiplied QImage__Format = 6
const QImage__Format_RGB16 QImage__Format = 7
const QImage__Format_ARGB8565_Premultiplied QImage__Format = 8
const QImage__Format_RGB666 QImage__Format = 9
const QImage__Format_ARGB6666_Premultiplied QImage__Format = 10
const QImage__Format_RGB555 QImage__Format = 11
const QImage__Format_ARGB8555_Premultiplied QImage__Format = 12
const QImage__Format_RGB888 QImage__Format = 13
const QImage__Format_RGB444 QImage__Format = 14
const QImage__Format_ARGB4444_Premultiplied QImage__Format = 15
const QImage__Format_RGBX8888 QImage__Format = 16
const QImage__Format_RGBA8888 QImage__Format = 17
const QImage__Format_RGBA8888_Premultiplied QImage__Format = 18
const QImage__Format_BGR30 QImage__Format = 19
const QImage__Format_A2BGR30_Premultiplied QImage__Format = 20
const QImage__Format_RGB30 QImage__Format = 21
const QImage__Format_A2RGB30_Premultiplied QImage__Format = 22
const QImage__Format_Alpha8 QImage__Format = 23
const QImage__Format_Grayscale8 QImage__Format = 24
const QImage__NImageFormats QImage__Format = 25

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
