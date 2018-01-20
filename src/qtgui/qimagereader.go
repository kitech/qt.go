//  header block begin
// /usr/include/qt/QtGui/qimagereader.h
// #include <qimagereader.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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
type QImageReader struct {
	*qtrt.CObject
}

func (this *QImageReader) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtGui/qimagereader.h:71
// index:0
// void QImageReader()
func NewQImageReader() *QImageReader {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReaderC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQImageReaderFromPointer(cthis)
	return gothis
}
func NewQImageReaderFromPointer(cthis unsafe.Pointer) *QImageReader {
	return &QImageReader{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qimagereader.h:72
// index:1
// void QImageReader(class QIODevice *, const class QByteArray &)
func NewQImageReader_1(device unsafe.Pointer, format unsafe.Pointer) *QImageReader {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReaderC2EP9QIODeviceRK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, device, format)
	gopp.ErrPrint(err, rv)
	gothis := NewQImageReaderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qimagereader.h:73
// index:2
// void QImageReader(const class QString &, const class QByteArray &)
func NewQImageReader_2(fileName unsafe.Pointer, format unsafe.Pointer) *QImageReader {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReaderC2ERK7QStringRK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, fileName, format)
	gopp.ErrPrint(err, rv)
	gothis := NewQImageReaderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qimagereader.h:74
// index:0
// void ~QImageReader()
func DeleteQImageReader(*QImageReader) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReaderD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:76
// index:0
// void setFormat(const class QByteArray &)
func (this *QImageReader) SetFormat(format unsafe.Pointer) {
	// 0: (, format const QByteArray &), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader9setFormatERK10QByteArray", ffiqt.FFI_TYPE_VOID, this.GetCthis(), format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:77
// index:0
// QByteArray format()
func (this *QImageReader) Format() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader6formatEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:79
// index:0
// void setAutoDetectImageFormat(_Bool)
func (this *QImageReader) SetAutoDetectImageFormat(enabled bool) {
	// 0: (, enabled bool), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader24setAutoDetectImageFormatEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:80
// index:0
// bool autoDetectImageFormat()
func (this *QImageReader) AutoDetectImageFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader21autoDetectImageFormatEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:82
// index:0
// void setDecideFormatFromContent(_Bool)
func (this *QImageReader) SetDecideFormatFromContent(ignored bool) {
	// 0: (, ignored bool), (&ignored)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader26setDecideFormatFromContentEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &ignored)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:83
// index:0
// bool decideFormatFromContent()
func (this *QImageReader) DecideFormatFromContent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader23decideFormatFromContentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:85
// index:0
// void setDevice(class QIODevice *)
func (this *QImageReader) SetDevice(device unsafe.Pointer) {
	// 0: (, device QIODevice *), (device)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader9setDeviceEP9QIODevice", ffiqt.FFI_TYPE_VOID, this.GetCthis(), device)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:86
// index:0
// QIODevice * device()
func (this *QImageReader) Device() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader6deviceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:88
// index:0
// void setFileName(const class QString &)
func (this *QImageReader) SetFileName(fileName unsafe.Pointer) {
	// 0: (, fileName const QString &), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader11setFileNameERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), fileName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:89
// index:0
// QString fileName()
func (this *QImageReader) FileName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader8fileNameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:91
// index:0
// QSize size()
func (this *QImageReader) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader4sizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:93
// index:0
// QImage::Format imageFormat()
func (this *QImageReader) ImageFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader11imageFormatEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:143
// index:1
// static
// QByteArray imageFormat(const class QString &)
func (this *QImageReader) ImageFormat_1(fileName unsafe.Pointer) {
	// 1: (fileName const QString &), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader11imageFormatERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QImageReader_ImageFormat_1(fileName unsafe.Pointer) {
	// 1: (fileName const QString &), (fileName)
	var nilthis *QImageReader
	nilthis.ImageFormat_1(fileName)
}

// /usr/include/qt/QtGui/qimagereader.h:144
// index:2
// static
// QByteArray imageFormat(class QIODevice *)
func (this *QImageReader) ImageFormat_2(device unsafe.Pointer) {
	// 2: (device QIODevice *), (device)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader11imageFormatEP9QIODevice", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QImageReader_ImageFormat_2(device unsafe.Pointer) {
	// 2: (device QIODevice *), (device)
	var nilthis *QImageReader
	nilthis.ImageFormat_2(device)
}

// /usr/include/qt/QtGui/qimagereader.h:95
// index:0
// QStringList textKeys()
func (this *QImageReader) TextKeys() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader8textKeysEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:96
// index:0
// QString text(const class QString &)
func (this *QImageReader) Text(key unsafe.Pointer) {
	// 0: (, key const QString &), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader4textERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:98
// index:0
// void setClipRect(const class QRect &)
func (this *QImageReader) SetClipRect(rect unsafe.Pointer) {
	// 0: (, rect const QRect &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader11setClipRectERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:99
// index:0
// QRect clipRect()
func (this *QImageReader) ClipRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader8clipRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:101
// index:0
// void setScaledSize(const class QSize &)
func (this *QImageReader) SetScaledSize(size unsafe.Pointer) {
	// 0: (, size const QSize &), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader13setScaledSizeERK5QSize", ffiqt.FFI_TYPE_VOID, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:102
// index:0
// QSize scaledSize()
func (this *QImageReader) ScaledSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader10scaledSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:104
// index:0
// void setQuality(int)
func (this *QImageReader) SetQuality(quality int) {
	// 0: (, quality int), (&quality)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader10setQualityEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &quality)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:105
// index:0
// int quality()
func (this *QImageReader) Quality() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader7qualityEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:107
// index:0
// void setScaledClipRect(const class QRect &)
func (this *QImageReader) SetScaledClipRect(rect unsafe.Pointer) {
	// 0: (, rect const QRect &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader17setScaledClipRectERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:108
// index:0
// QRect scaledClipRect()
func (this *QImageReader) ScaledClipRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader14scaledClipRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:110
// index:0
// void setBackgroundColor(const class QColor &)
func (this *QImageReader) SetBackgroundColor(color unsafe.Pointer) {
	// 0: (, color const QColor &), (color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader18setBackgroundColorERK6QColor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:111
// index:0
// QColor backgroundColor()
func (this *QImageReader) BackgroundColor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader15backgroundColorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:113
// index:0
// bool supportsAnimation()
func (this *QImageReader) SupportsAnimation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader17supportsAnimationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:115
// index:0
// QImageIOHandler::Transformations transformation()
func (this *QImageReader) Transformation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader14transformationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:117
// index:0
// void setAutoTransform(_Bool)
func (this *QImageReader) SetAutoTransform(enabled bool) {
	// 0: (, enabled bool), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader16setAutoTransformEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:118
// index:0
// bool autoTransform()
func (this *QImageReader) AutoTransform() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader13autoTransformEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:120
// index:0
// void setGamma(float)
func (this *QImageReader) SetGamma(gamma float32) {
	// 0: (, gamma float), (&gamma)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader8setGammaEf", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &gamma)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:121
// index:0
// float gamma()
func (this *QImageReader) Gamma() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader5gammaEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:123
// index:0
// QByteArray subType()
func (this *QImageReader) SubType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader7subTypeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:124
// index:0
// QList<QByteArray> supportedSubTypes()
func (this *QImageReader) SupportedSubTypes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader17supportedSubTypesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:126
// index:0
// bool canRead()
func (this *QImageReader) CanRead() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader7canReadEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:127
// index:0
// QImage read()
func (this *QImageReader) Read() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader4readEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:128
// index:1
// bool read(class QImage *)
func (this *QImageReader) Read_1(image unsafe.Pointer) {
	// 1: (, image QImage *), (image)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader4readEP6QImage", ffiqt.FFI_TYPE_VOID, this.GetCthis(), image)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:130
// index:0
// bool jumpToNextImage()
func (this *QImageReader) JumpToNextImage() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader15jumpToNextImageEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:131
// index:0
// bool jumpToImage(int)
func (this *QImageReader) JumpToImage(imageNumber int) {
	// 0: (, imageNumber int), (&imageNumber)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader11jumpToImageEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &imageNumber)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:132
// index:0
// int loopCount()
func (this *QImageReader) LoopCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader9loopCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:133
// index:0
// int imageCount()
func (this *QImageReader) ImageCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader10imageCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:134
// index:0
// int nextImageDelay()
func (this *QImageReader) NextImageDelay() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader14nextImageDelayEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:135
// index:0
// int currentImageNumber()
func (this *QImageReader) CurrentImageNumber() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader18currentImageNumberEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:136
// index:0
// QRect currentImageRect()
func (this *QImageReader) CurrentImageRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader16currentImageRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:138
// index:0
// QImageReader::ImageReaderError error()
func (this *QImageReader) Error() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader5errorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:139
// index:0
// QString errorString()
func (this *QImageReader) ErrorString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader11errorStringEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:141
// index:0
// bool supportsOption(class QImageIOHandler::ImageOption)
func (this *QImageReader) SupportsOption(option int) {
	// 0: (, option QImageIOHandler::ImageOption), (&option)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader14supportsOptionEN15QImageIOHandler11ImageOptionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &option)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:145
// index:0
// static
// QList<QByteArray> supportedImageFormats()
func (this *QImageReader) SupportedImageFormats() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader21supportedImageFormatsEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QImageReader_SupportedImageFormats() {
	// 0: (), ()
	var nilthis *QImageReader
	nilthis.SupportedImageFormats()
}

// /usr/include/qt/QtGui/qimagereader.h:146
// index:0
// static
// QList<QByteArray> supportedMimeTypes()
func (this *QImageReader) SupportedMimeTypes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader18supportedMimeTypesEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QImageReader_SupportedMimeTypes() {
	// 0: (), ()
	var nilthis *QImageReader
	nilthis.SupportedMimeTypes()
}

//  body block end
