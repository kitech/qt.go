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
func NewQImageReaderFromPointer(cthis unsafe.Pointer) *QImageReader {
	return &QImageReader{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qimagereader.h:71
// index:0
// Public
// void QImageReader()
func NewQImageReader() *QImageReader {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReaderC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQImageReaderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qimagereader.h:72
// index:1
// Public
// void QImageReader(class QIODevice *, const class QByteArray &)
func NewQImageReader_1(device unsafe.Pointer, format *qtcore.QByteArray) *QImageReader {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg1 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReaderC2EP9QIODeviceRK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, device, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQImageReaderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qimagereader.h:73
// index:2
// Public
// void QImageReader(const class QString &, const class QByteArray &)
func NewQImageReader_2(fileName *qtcore.QString, format *qtcore.QByteArray) *QImageReader {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = fileName.GetCthis()
	var convArg1 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReaderC2ERK7QStringRK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQImageReaderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qimagereader.h:74
// index:0
// Public
// void ~QImageReader()
func DeleteQImageReader(*QImageReader) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReaderD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:76
// index:0
// Public
// void setFormat(const class QByteArray &)
func (this *QImageReader) SetFormat(format *qtcore.QByteArray) {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader9setFormatERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:77
// index:0
// Public
// QByteArray format()
func (this *QImageReader) Format() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader6formatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:79
// index:0
// Public
// void setAutoDetectImageFormat(_Bool)
func (this *QImageReader) SetAutoDetectImageFormat(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader24setAutoDetectImageFormatEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:80
// index:0
// Public
// bool autoDetectImageFormat()
func (this *QImageReader) AutoDetectImageFormat() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader21autoDetectImageFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:82
// index:0
// Public
// void setDecideFormatFromContent(_Bool)
func (this *QImageReader) SetDecideFormatFromContent(ignored bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader26setDecideFormatFromContentEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &ignored)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:83
// index:0
// Public
// bool decideFormatFromContent()
func (this *QImageReader) DecideFormatFromContent() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader23decideFormatFromContentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:85
// index:0
// Public
// void setDevice(class QIODevice *)
func (this *QImageReader) SetDevice(device unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader9setDeviceEP9QIODevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), device)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:86
// index:0
// Public
// QIODevice * device()
func (this *QImageReader) Device() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader6deviceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:88
// index:0
// Public
// void setFileName(const class QString &)
func (this *QImageReader) SetFileName(fileName *qtcore.QString) {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader11setFileNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:89
// index:0
// Public
// QString fileName()
func (this *QImageReader) FileName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader8fileNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:91
// index:0
// Public
// QSize size()
func (this *QImageReader) Size() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:93
// index:0
// Public
// QImage::Format imageFormat()
func (this *QImageReader) ImageFormat() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader11imageFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:143
// index:1
// Public static
// QByteArray imageFormat(const class QString &)
func (this *QImageReader) ImageFormat_1(fileName *qtcore.QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader11imageFormatERK7QString", ffiqt.FFI_TYPE_POINTER, fileName)
	gopp.ErrPrint(err, rv)
	return rv
}
func QImageReader_ImageFormat_1(fileName *qtcore.QString) {
	var nilthis *QImageReader
	nilthis.ImageFormat_1(fileName)
}

// /usr/include/qt/QtGui/qimagereader.h:144
// index:2
// Public static
// QByteArray imageFormat(class QIODevice *)
func (this *QImageReader) ImageFormat_2(device unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader11imageFormatEP9QIODevice", ffiqt.FFI_TYPE_POINTER, device)
	gopp.ErrPrint(err, rv)
	return rv
}
func QImageReader_ImageFormat_2(device unsafe.Pointer) {
	var nilthis *QImageReader
	nilthis.ImageFormat_2(device)
}

// /usr/include/qt/QtGui/qimagereader.h:95
// index:0
// Public
// QStringList textKeys()
func (this *QImageReader) TextKeys() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader8textKeysEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:96
// index:0
// Public
// QString text(const class QString &)
func (this *QImageReader) Text(key *qtcore.QString) interface{} {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader4textERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:98
// index:0
// Public
// void setClipRect(const class QRect &)
func (this *QImageReader) SetClipRect(rect *qtcore.QRect) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader11setClipRectERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:99
// index:0
// Public
// QRect clipRect()
func (this *QImageReader) ClipRect() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader8clipRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:101
// index:0
// Public
// void setScaledSize(const class QSize &)
func (this *QImageReader) SetScaledSize(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader13setScaledSizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:102
// index:0
// Public
// QSize scaledSize()
func (this *QImageReader) ScaledSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader10scaledSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:104
// index:0
// Public
// void setQuality(int)
func (this *QImageReader) SetQuality(quality int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader10setQualityEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &quality)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:105
// index:0
// Public
// int quality()
func (this *QImageReader) Quality() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader7qualityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:107
// index:0
// Public
// void setScaledClipRect(const class QRect &)
func (this *QImageReader) SetScaledClipRect(rect *qtcore.QRect) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader17setScaledClipRectERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:108
// index:0
// Public
// QRect scaledClipRect()
func (this *QImageReader) ScaledClipRect() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader14scaledClipRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:110
// index:0
// Public
// void setBackgroundColor(const class QColor &)
func (this *QImageReader) SetBackgroundColor(color *QColor) {
	var convArg0 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader18setBackgroundColorERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:111
// index:0
// Public
// QColor backgroundColor()
func (this *QImageReader) BackgroundColor() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader15backgroundColorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:113
// index:0
// Public
// bool supportsAnimation()
func (this *QImageReader) SupportsAnimation() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader17supportsAnimationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:115
// index:0
// Public
// QImageIOHandler::Transformations transformation()
func (this *QImageReader) Transformation() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader14transformationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:117
// index:0
// Public
// void setAutoTransform(_Bool)
func (this *QImageReader) SetAutoTransform(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader16setAutoTransformEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:118
// index:0
// Public
// bool autoTransform()
func (this *QImageReader) AutoTransform() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader13autoTransformEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:120
// index:0
// Public
// void setGamma(float)
func (this *QImageReader) SetGamma(gamma float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader8setGammaEf", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &gamma)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:121
// index:0
// Public
// float gamma()
func (this *QImageReader) Gamma() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader5gammaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:123
// index:0
// Public
// QByteArray subType()
func (this *QImageReader) SubType() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader7subTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:124
// index:0
// Public
// QList<QByteArray> supportedSubTypes()
func (this *QImageReader) SupportedSubTypes() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader17supportedSubTypesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:126
// index:0
// Public
// bool canRead()
func (this *QImageReader) CanRead() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader7canReadEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:127
// index:0
// Public
// QImage read()
func (this *QImageReader) Read() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader4readEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:128
// index:1
// Public
// bool read(class QImage *)
func (this *QImageReader) Read_1(image unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader4readEP6QImage", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), image)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:130
// index:0
// Public
// bool jumpToNextImage()
func (this *QImageReader) JumpToNextImage() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader15jumpToNextImageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:131
// index:0
// Public
// bool jumpToImage(int)
func (this *QImageReader) JumpToImage(imageNumber int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader11jumpToImageEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &imageNumber)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:132
// index:0
// Public
// int loopCount()
func (this *QImageReader) LoopCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader9loopCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:133
// index:0
// Public
// int imageCount()
func (this *QImageReader) ImageCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader10imageCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:134
// index:0
// Public
// int nextImageDelay()
func (this *QImageReader) NextImageDelay() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader14nextImageDelayEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:135
// index:0
// Public
// int currentImageNumber()
func (this *QImageReader) CurrentImageNumber() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader18currentImageNumberEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:136
// index:0
// Public
// QRect currentImageRect()
func (this *QImageReader) CurrentImageRect() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader16currentImageRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:138
// index:0
// Public
// QImageReader::ImageReaderError error()
func (this *QImageReader) Error() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader5errorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:139
// index:0
// Public
// QString errorString()
func (this *QImageReader) ErrorString() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader11errorStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:141
// index:0
// Public
// bool supportsOption(class QImageIOHandler::ImageOption)
func (this *QImageReader) SupportsOption(option int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader14supportsOptionEN15QImageIOHandler11ImageOptionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &option)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:145
// index:0
// Public static
// QList<QByteArray> supportedImageFormats()
func (this *QImageReader) SupportedImageFormats() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader21supportedImageFormatsEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QImageReader_SupportedImageFormats() {
	var nilthis *QImageReader
	nilthis.SupportedImageFormats()
}

// /usr/include/qt/QtGui/qimagereader.h:146
// index:0
// Public static
// QList<QByteArray> supportedMimeTypes()
func (this *QImageReader) SupportedMimeTypes() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader18supportedMimeTypesEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QImageReader_SupportedMimeTypes() {
	var nilthis *QImageReader
	nilthis.SupportedMimeTypes()
}

//  body block end
