//  header block begin
// /usr/include/qt/QtGui/qimage.h
// #include <qimage.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 24
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
type QImage struct {
	*QPaintDevice
}

func (this *QImage) GetCthis() unsafe.Pointer {
	return this.QPaintDevice.GetCthis()
}

// /usr/include/qt/QtGui/qimage.h:136
// index:0
// void QImage()
func NewQImage() *QImage {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImageC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(cthis)
	return gothis
}
func NewQImageFromPointer(cthis unsafe.Pointer) *QImage {
	bcthis0 := NewQPaintDeviceFromPointer(cthis)
	return &QImage{bcthis0}
}

// /usr/include/qt/QtGui/qimage.h:137
// index:1
// void QImage(const class QSize &, enum QImage::Format)
func NewQImage_1(size unsafe.Pointer, format int) *QImage {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImageC2ERK5QSizeNS_6FormatE", ffiqt.FFI_TYPE_VOID, cthis, size, &format)
	gopp.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:138
// index:2
// void QImage(int, int, enum QImage::Format)
func NewQImage_2(width int, height int, format int) *QImage {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImageC2EiiNS_6FormatE", ffiqt.FFI_TYPE_VOID, cthis, &width, &height, &format)
	gopp.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:145
// index:3
// void QImage(const char *const *)
func NewQImage_3(xpm []interface{}) *QImage {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImageC2EPKPKc", ffiqt.FFI_TYPE_VOID, cthis, xpm)
	gopp.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:147
// index:4
// void QImage(const class QString &, const char *)
func NewQImage_4(fileName unsafe.Pointer, format unsafe.Pointer) *QImage {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImageC2ERK7QStringPKc", ffiqt.FFI_TYPE_VOID, cthis, fileName, format)
	gopp.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:155
// index:0
// virtual
// void ~QImage()
func DeleteQImage(*QImage) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImageD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:162
// index:0
// inline
// void swap(class QImage &)
func (this *QImage) Swap(other unsafe.Pointer) {
	// 0: (, other QImage &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:165
// index:0
// bool isNull()
func (this *QImage) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage6isNullEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:167
// index:0
// virtual
// int devType()
func (this *QImage) DevType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage7devTypeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:172
// index:0
// void detach()
func (this *QImage) Detach() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage6detachEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:173
// index:0
// bool isDetached()
func (this *QImage) IsDetached() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage10isDetachedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:175
// index:0
// QImage copy(const class QRect &)
func (this *QImage) Copy(rect unsafe.Pointer) {
	// 0: (, rect const QRect &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage4copyERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:176
// index:1
// inline
// QImage copy(int, int, int, int)
func (this *QImage) Copy_1(x int, y int, w int, h int) {
	// 1: (, x int, y int, w int, h int), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage4copyEiiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:179
// index:0
// QImage::Format format()
func (this *QImage) Format() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage6formatEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:182
// index:0
// inline
// QImage convertToFormat(enum QImage::Format, Qt::ImageConversionFlags)
func (this *QImage) ConvertToFormat(f int, flags int) {
	// 0: (, f QImage::Format, QFlags<Qt::ImageConversionFlag> flags), (&f, &flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZNKR6QImage15convertToFormatENS_6FormatE6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &f, &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:184
// index:1
// inline
// QImage convertToFormat(enum QImage::Format, Qt::ImageConversionFlags)
func (this *QImage) ConvertToFormat_1(f int, flags int) {
	// 1: (, f QImage::Format, QFlags<Qt::ImageConversionFlag> flags), (&f, &flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZNO6QImage15convertToFormatENS_6FormatE6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &f, &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:195
// index:0
// bool reinterpretAsFormat(enum QImage::Format)
func (this *QImage) ReinterpretAsFormat(f int) {
	// 0: (, f QImage::Format), (&f)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage19reinterpretAsFormatENS_6FormatE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &f)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:197
// index:0
// int width()
func (this *QImage) Width() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage5widthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:198
// index:0
// int height()
func (this *QImage) Height() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage6heightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:199
// index:0
// QSize size()
func (this *QImage) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage4sizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:200
// index:0
// QRect rect()
func (this *QImage) Rect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage4rectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:202
// index:0
// int depth()
func (this *QImage) Depth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage5depthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:203
// index:0
// int colorCount()
func (this *QImage) ColorCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage10colorCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:204
// index:0
// int bitPlaneCount()
func (this *QImage) BitPlaneCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage13bitPlaneCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:206
// index:0
// QRgb color(int)
func (this *QImage) Color(i int) {
	// 0: (, i int), (&i)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage5colorEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:207
// index:0
// void setColor(int, QRgb)
func (this *QImage) SetColor(i int, c uint) {
	// 0: (, i int, c QRgb), (&i, &c)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage8setColorEij", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &i, &c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:208
// index:0
// void setColorCount(int)
func (this *QImage) SetColorCount(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage13setColorCountEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:210
// index:0
// bool allGray()
func (this *QImage) AllGray() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage7allGrayEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:211
// index:0
// bool isGrayscale()
func (this *QImage) IsGrayscale() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage11isGrayscaleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:213
// index:0
// uchar * bits()
func (this *QImage) Bits() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage4bitsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:214
// index:1
// const uchar * bits()
func (this *QImage) Bits_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage4bitsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:215
// index:0
// const uchar * constBits()
func (this *QImage) ConstBits() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage9constBitsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:218
// index:0
// int byteCount()
func (this *QImage) ByteCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage9byteCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:220
// index:0
// qsizetype sizeInBytes()
func (this *QImage) SizeInBytes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage11sizeInBytesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:222
// index:0
// uchar * scanLine(int)
func (this *QImage) ScanLine(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage8scanLineEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:223
// index:1
// const uchar * scanLine(int)
func (this *QImage) ScanLine_1(arg0 int) {
	// 1: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage8scanLineEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:224
// index:0
// const uchar * constScanLine(int)
func (this *QImage) ConstScanLine(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage13constScanLineEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:225
// index:0
// int bytesPerLine()
func (this *QImage) BytesPerLine() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage12bytesPerLineEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:227
// index:0
// bool valid(int, int)
func (this *QImage) Valid(x int, y int) {
	// 0: (, x int, y int), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage5validEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:228
// index:1
// bool valid(const class QPoint &)
func (this *QImage) Valid_1(pt unsafe.Pointer) {
	// 1: (, pt const QPoint &), (pt)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage5validERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pt)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:230
// index:0
// int pixelIndex(int, int)
func (this *QImage) PixelIndex(x int, y int) {
	// 0: (, x int, y int), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage10pixelIndexEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:231
// index:1
// int pixelIndex(const class QPoint &)
func (this *QImage) PixelIndex_1(pt unsafe.Pointer) {
	// 1: (, pt const QPoint &), (pt)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage10pixelIndexERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pt)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:233
// index:0
// QRgb pixel(int, int)
func (this *QImage) Pixel(x int, y int) {
	// 0: (, x int, y int), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage5pixelEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:234
// index:1
// QRgb pixel(const class QPoint &)
func (this *QImage) Pixel_1(pt unsafe.Pointer) {
	// 1: (, pt const QPoint &), (pt)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage5pixelERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pt)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:236
// index:0
// void setPixel(int, int, uint)
func (this *QImage) SetPixel(x int, y int, index_or_rgb uint) {
	// 0: (, x int, y int, index_or_rgb uint), (&x, &y, &index_or_rgb)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage8setPixelEiij", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &index_or_rgb)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:237
// index:1
// void setPixel(const class QPoint &, uint)
func (this *QImage) SetPixel_1(pt unsafe.Pointer, index_or_rgb uint) {
	// 1: (, pt const QPoint &, index_or_rgb uint), (pt, &index_or_rgb)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage8setPixelERK6QPointj", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pt, &index_or_rgb)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:239
// index:0
// QColor pixelColor(int, int)
func (this *QImage) PixelColor(x int, y int) {
	// 0: (, x int, y int), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage10pixelColorEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:240
// index:1
// QColor pixelColor(const class QPoint &)
func (this *QImage) PixelColor_1(pt unsafe.Pointer) {
	// 1: (, pt const QPoint &), (pt)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage10pixelColorERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pt)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:242
// index:0
// void setPixelColor(int, int, const class QColor &)
func (this *QImage) SetPixelColor(x int, y int, c unsafe.Pointer) {
	// 0: (, x int, y int, c const QColor &), (&x, &y, c)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage13setPixelColorEiiRK6QColor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:243
// index:1
// void setPixelColor(const class QPoint &, const class QColor &)
func (this *QImage) SetPixelColor_1(pt unsafe.Pointer, c unsafe.Pointer) {
	// 1: (, pt const QPoint &, c const QColor &), (pt, c)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage13setPixelColorERK6QPointRK6QColor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pt, c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:245
// index:0
// QVector<QRgb> colorTable()
func (this *QImage) ColorTable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage10colorTableEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:252
// index:0
// qreal devicePixelRatio()
func (this *QImage) DevicePixelRatio() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage16devicePixelRatioEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:253
// index:0
// void setDevicePixelRatio(qreal)
func (this *QImage) SetDevicePixelRatio(scaleFactor float64) {
	// 0: (, scaleFactor qreal), (&scaleFactor)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage19setDevicePixelRatioEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &scaleFactor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:255
// index:0
// void fill(uint)
func (this *QImage) Fill(pixel uint) {
	// 0: (, pixel uint), (&pixel)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage4fillEj", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &pixel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:256
// index:1
// void fill(const class QColor &)
func (this *QImage) Fill_1(color unsafe.Pointer) {
	// 1: (, color const QColor &), (color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage4fillERK6QColor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:257
// index:2
// void fill(Qt::GlobalColor)
func (this *QImage) Fill_2(color int) {
	// 2: (, color Qt::GlobalColor), (&color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage4fillEN2Qt11GlobalColorE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:260
// index:0
// bool hasAlphaChannel()
func (this *QImage) HasAlphaChannel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage15hasAlphaChannelEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:261
// index:0
// void setAlphaChannel(const class QImage &)
func (this *QImage) SetAlphaChannel(alphaChannel unsafe.Pointer) {
	// 0: (, alphaChannel const QImage &), (alphaChannel)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage15setAlphaChannelERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), alphaChannel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:262
// index:0
// QImage alphaChannel()
func (this *QImage) AlphaChannel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage12alphaChannelEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:263
// index:0
// QImage createAlphaMask(Qt::ImageConversionFlags)
func (this *QImage) CreateAlphaMask(flags int) {
	// 0: (, QFlags<Qt::ImageConversionFlag> flags), (&flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage15createAlphaMaskE6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:265
// index:0
// QImage createHeuristicMask(_Bool)
func (this *QImage) CreateHeuristicMask(clipTight bool) {
	// 0: (, clipTight bool), (&clipTight)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage19createHeuristicMaskEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &clipTight)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:267
// index:0
// QImage createMaskFromColor(QRgb, Qt::MaskMode)
func (this *QImage) CreateMaskFromColor(color uint, mode int) {
	// 0: (, color QRgb, mode Qt::MaskMode), (&color, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage19createMaskFromColorEjN2Qt8MaskModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &color, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:269
// index:0
// inline
// QImage scaled(int, int, Qt::AspectRatioMode, Qt::TransformationMode)
func (this *QImage) Scaled(w int, h int, aspectMode int, mode int) {
	// 0: (, w int, h int, aspectMode Qt::AspectRatioMode, mode Qt::TransformationMode), (&w, &h, &aspectMode, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage6scaledEiiN2Qt15AspectRatioModeENS0_18TransformationModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &w, &h, &aspectMode, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:272
// index:1
// QImage scaled(const class QSize &, Qt::AspectRatioMode, Qt::TransformationMode)
func (this *QImage) Scaled_1(s unsafe.Pointer, aspectMode int, mode int) {
	// 1: (, s const QSize &, aspectMode Qt::AspectRatioMode, mode Qt::TransformationMode), (s, &aspectMode, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage6scaledERK5QSizeN2Qt15AspectRatioModeENS3_18TransformationModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, &aspectMode, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:274
// index:0
// QImage scaledToWidth(int, Qt::TransformationMode)
func (this *QImage) ScaledToWidth(w int, mode int) {
	// 0: (, w int, mode Qt::TransformationMode), (&w, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage13scaledToWidthEiN2Qt18TransformationModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &w, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:275
// index:0
// QImage scaledToHeight(int, Qt::TransformationMode)
func (this *QImage) ScaledToHeight(h int, mode int) {
	// 0: (, h int, mode Qt::TransformationMode), (&h, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage14scaledToHeightEiN2Qt18TransformationModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &h, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:276
// index:0
// QImage transformed(const class QMatrix &, Qt::TransformationMode)
func (this *QImage) Transformed(matrix unsafe.Pointer, mode int) {
	// 0: (, matrix const QMatrix &, mode Qt::TransformationMode), (matrix, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage11transformedERK7QMatrixN2Qt18TransformationModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), matrix, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:278
// index:1
// QImage transformed(const class QTransform &, Qt::TransformationMode)
func (this *QImage) Transformed_1(matrix unsafe.Pointer, mode int) {
	// 1: (, matrix const QTransform &, mode Qt::TransformationMode), (matrix, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage11transformedERK10QTransformN2Qt18TransformationModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), matrix, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:277
// index:0
// static
// QMatrix trueMatrix(const class QMatrix &, int, int)
func (this *QImage) TrueMatrix(arg0 unsafe.Pointer, w int, h int) {
	// 0: (const QMatrix & arg0, w int, h int), (arg0, w, h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage10trueMatrixERK7QMatrixii", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QImage_TrueMatrix(arg0 unsafe.Pointer, w int, h int) {
	// 0: (const QMatrix & arg0, w int, h int), (arg0, w, h)
	var nilthis *QImage
	nilthis.TrueMatrix(arg0, w, h)
}

// /usr/include/qt/QtGui/qimage.h:279
// index:1
// static
// QTransform trueMatrix(const class QTransform &, int, int)
func (this *QImage) TrueMatrix_1(arg0 unsafe.Pointer, w int, h int) {
	// 1: (const QTransform & arg0, w int, h int), (arg0, w, h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage10trueMatrixERK10QTransformii", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QImage_TrueMatrix_1(arg0 unsafe.Pointer, w int, h int) {
	// 1: (const QTransform & arg0, w int, h int), (arg0, w, h)
	var nilthis *QImage
	nilthis.TrueMatrix_1(arg0, w, h)
}

// /usr/include/qt/QtGui/qimage.h:281
// index:0
// inline
// QImage mirrored(_Bool, _Bool)
func (this *QImage) Mirrored(horizontally bool, vertically bool) {
	// 0: (, horizontally bool, vertically bool), (&horizontally, &vertically)
	rv, err := ffiqt.InvokeQtFunc6("_ZNKR6QImage8mirroredEbb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &horizontally, &vertically)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:283
// index:1
// inline
// QImage && mirrored(_Bool, _Bool)
func (this *QImage) Mirrored_1(horizontally bool, vertically bool) {
	// 1: (, horizontally bool, vertically bool), (&horizontally, &vertically)
	rv, err := ffiqt.InvokeQtFunc6("_ZNO6QImage8mirroredEbb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &horizontally, &vertically)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:285
// index:0
// inline
// QImage rgbSwapped()
func (this *QImage) RgbSwapped() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNKR6QImage10rgbSwappedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:287
// index:1
// inline
// QImage && rgbSwapped()
func (this *QImage) RgbSwapped_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNO6QImage10rgbSwappedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:293
// index:0
// void invertPixels(enum QImage::InvertMode)
func (this *QImage) InvertPixels(arg0 int) {
	// 0: (, QImage::InvertMode arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage12invertPixelsENS_10InvertModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:296
// index:0
// bool load(class QIODevice *, const char *)
func (this *QImage) Load(device unsafe.Pointer, format unsafe.Pointer) {
	// 0: (, device QIODevice *, format const char *), (device, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage4loadEP9QIODevicePKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), device, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:297
// index:1
// bool load(const class QString &, const char *)
func (this *QImage) Load_1(fileName unsafe.Pointer, format unsafe.Pointer) {
	// 1: (, fileName const QString &, format const char *), (fileName, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage4loadERK7QStringPKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), fileName, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:298
// index:0
// bool loadFromData(const uchar *, int, const char *)
func (this *QImage) LoadFromData(buf unsafe.Pointer, len int, format unsafe.Pointer) {
	// 0: (, buf const uchar *, len int, format const char *), (buf, &len, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage12loadFromDataEPKhiPKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), buf, &len, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:299
// index:1
// inline
// bool loadFromData(const class QByteArray &, const char *)
func (this *QImage) LoadFromData_1(data unsafe.Pointer, aformat unsafe.Pointer) {
	// 1: (, data const QByteArray &, aformat const char *), (data, aformat)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage12loadFromDataERK10QByteArrayPKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data, aformat)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:302
// index:0
// bool save(const class QString &, const char *, int)
func (this *QImage) Save(fileName unsafe.Pointer, format unsafe.Pointer, quality int) {
	// 0: (, fileName const QString &, format const char *, quality int), (fileName, format, &quality)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage4saveERK7QStringPKci", ffiqt.FFI_TYPE_VOID, this.GetCthis(), fileName, format, &quality)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:303
// index:1
// bool save(class QIODevice *, const char *, int)
func (this *QImage) Save_1(device unsafe.Pointer, format unsafe.Pointer, quality int) {
	// 1: (, device QIODevice *, format const char *, quality int), (device, format, &quality)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage4saveEP9QIODevicePKci", ffiqt.FFI_TYPE_VOID, this.GetCthis(), device, format, &quality)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:305
// index:0
// static
// QImage fromData(const uchar *, int, const char *)
func (this *QImage) FromData(data unsafe.Pointer, size int, format unsafe.Pointer) {
	// 0: (data const uchar *, size int, format const char *), (data, size, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage8fromDataEPKhiPKc", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QImage_FromData(data unsafe.Pointer, size int, format unsafe.Pointer) {
	// 0: (data const uchar *, size int, format const char *), (data, size, format)
	var nilthis *QImage
	nilthis.FromData(data, size, format)
}

// /usr/include/qt/QtGui/qimage.h:306
// index:1
// static inline
// QImage fromData(const class QByteArray &, const char *)
func (this *QImage) FromData_1(data unsafe.Pointer, format unsafe.Pointer) {
	// 1: (data const QByteArray &, format const char *), (data, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage8fromDataERK10QByteArrayPKc", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QImage_FromData_1(data unsafe.Pointer, format unsafe.Pointer) {
	// 1: (data const QByteArray &, format const char *), (data, format)
	var nilthis *QImage
	nilthis.FromData_1(data, format)
}

// /usr/include/qt/QtGui/qimage.h:312
// index:0
// qint64 cacheKey()
func (this *QImage) CacheKey() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage8cacheKeyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:314
// index:0
// virtual
// QPaintEngine * paintEngine()
func (this *QImage) PaintEngine() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage11paintEngineEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:317
// index:0
// int dotsPerMeterX()
func (this *QImage) DotsPerMeterX() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage13dotsPerMeterXEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:318
// index:0
// int dotsPerMeterY()
func (this *QImage) DotsPerMeterY() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage13dotsPerMeterYEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:319
// index:0
// void setDotsPerMeterX(int)
func (this *QImage) SetDotsPerMeterX(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage16setDotsPerMeterXEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:320
// index:0
// void setDotsPerMeterY(int)
func (this *QImage) SetDotsPerMeterY(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage16setDotsPerMeterYEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:321
// index:0
// QPoint offset()
func (this *QImage) Offset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage6offsetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:322
// index:0
// void setOffset(const class QPoint &)
func (this *QImage) SetOffset(arg0 unsafe.Pointer) {
	// 0: (, const QPoint & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage9setOffsetERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:324
// index:0
// QStringList textKeys()
func (this *QImage) TextKeys() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage8textKeysEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:325
// index:0
// QString text(const class QString &)
func (this *QImage) Text(key unsafe.Pointer) {
	// 0: (, key const QString &), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage4textERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:326
// index:0
// void setText(const class QString &, const class QString &)
func (this *QImage) SetText(key unsafe.Pointer, value unsafe.Pointer) {
	// 0: (, key const QString &, value const QString &), (key, value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage7setTextERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), key, value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:328
// index:0
// QPixelFormat pixelFormat()
func (this *QImage) PixelFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage11pixelFormatEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:329
// index:0
// static
// QPixelFormat toPixelFormat(class QImage::Format)
func (this *QImage) ToPixelFormat(format int) {
	// 0: (format QImage::Format), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage13toPixelFormatENS_6FormatE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QImage_ToPixelFormat(format int) {
	// 0: (format QImage::Format), (format)
	var nilthis *QImage
	nilthis.ToPixelFormat(format)
}

// /usr/include/qt/QtGui/qimage.h:330
// index:0
// static
// QImage::Format toImageFormat(class QPixelFormat)
func (this *QImage) ToImageFormat(format unsafe.Pointer) {
	// 0: (format QPixelFormat), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage13toImageFormatE12QPixelFormat", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QImage_ToImageFormat(format unsafe.Pointer) {
	// 0: (format QPixelFormat), (format)
	var nilthis *QImage
	nilthis.ToImageFormat(format)
}

// /usr/include/qt/QtGui/qimage.h:352
// index:0
// virtual
// int metric(enum QPaintDevice::PaintDeviceMetric)
func (this *QImage) Metric(metric int) {
	// 0: (, metric QPaintDevice::PaintDeviceMetric), (&metric)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage6metricEN12QPaintDevice17PaintDeviceMetricE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &metric)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:353
// index:0
// QImage mirrored_helper(_Bool, _Bool)
func (this *QImage) Mirrored_helper(horizontal bool, vertical bool) {
	// 0: (, horizontal bool, vertical bool), (&horizontal, &vertical)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage15mirrored_helperEbb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &horizontal, &vertical)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:354
// index:0
// QImage rgbSwapped_helper()
func (this *QImage) RgbSwapped_helper() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage17rgbSwapped_helperEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:355
// index:0
// void mirrored_inplace(_Bool, _Bool)
func (this *QImage) Mirrored_inplace(horizontal bool, vertical bool) {
	// 0: (, horizontal bool, vertical bool), (&horizontal, &vertical)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage16mirrored_inplaceEbb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &horizontal, &vertical)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:356
// index:0
// void rgbSwapped_inplace()
func (this *QImage) RgbSwapped_inplace() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage18rgbSwapped_inplaceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:357
// index:0
// QImage convertToFormat_helper(enum QImage::Format, Qt::ImageConversionFlags)
func (this *QImage) ConvertToFormat_helper(format int, flags int) {
	// 0: (, format QImage::Format, QFlags<Qt::ImageConversionFlag> flags), (&format, &flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage22convertToFormat_helperENS_6FormatE6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &format, &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:358
// index:0
// bool convertToFormat_inplace(enum QImage::Format, Qt::ImageConversionFlags)
func (this *QImage) ConvertToFormat_inplace(format int, flags int) {
	// 0: (, format QImage::Format, QFlags<Qt::ImageConversionFlag> flags), (&format, &flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QImage23convertToFormat_inplaceENS_6FormatE6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &format, &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:359
// index:0
// QImage smoothScaled(int, int)
func (this *QImage) SmoothScaled(w int, h int) {
	// 0: (, w int, h int), (&w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QImage12smoothScaledEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &w, &h)
	gopp.ErrPrint(err, rv)
}

//  body block end
