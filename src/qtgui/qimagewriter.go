//  header block begin
// /usr/include/qt/QtGui/qimagewriter.h
// #include <qimagewriter.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 53
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
type QImageWriter struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qimagewriter.h:67
// index:0
// void QImageWriter()
func NewQImageWriter() *QImageWriter {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriterC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QImageWriter{cthis}
}

// /usr/include/qt/QtGui/qimagewriter.h:68
// index:1
// void QImageWriter(class QIODevice *, const class QByteArray &)
func NewQImageWriter_1(device unsafe.Pointer, format unsafe.Pointer) *QImageWriter {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriterC2EP9QIODeviceRK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, device, format)
	gopp.ErrPrint(err, rv)
	return &QImageWriter{cthis}
}

// /usr/include/qt/QtGui/qimagewriter.h:69
// index:2
// void QImageWriter(const class QString &, const class QByteArray &)
func NewQImageWriter_2(fileName unsafe.Pointer, format unsafe.Pointer) *QImageWriter {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriterC2ERK7QStringRK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, fileName, format)
	gopp.ErrPrint(err, rv)
	return &QImageWriter{cthis}
}

// /usr/include/qt/QtGui/qimagewriter.h:70
// index:0
// void ~QImageWriter()
func DeleteQImageWriter(*QImageWriter) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriterD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:72
// index:0
// void setFormat(const class QByteArray &)
func (this *QImageWriter) SetFormat(format unsafe.Pointer) {
	// 0: (, const QByteArray & format), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriter9setFormatERK10QByteArray", ffiqt.FFI_TYPE_VOID, this.cthis, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:73
// index:0
// QByteArray format()
func (this *QImageWriter) Format() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter6formatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:75
// index:0
// void setDevice(class QIODevice *)
func (this *QImageWriter) SetDevice(device unsafe.Pointer) {
	// 0: (, QIODevice * device), (device)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriter9setDeviceEP9QIODevice", ffiqt.FFI_TYPE_VOID, this.cthis, device)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:76
// index:0
// QIODevice * device()
func (this *QImageWriter) Device() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter6deviceEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:78
// index:0
// void setFileName(const class QString &)
func (this *QImageWriter) SetFileName(fileName unsafe.Pointer) {
	// 0: (, const QString & fileName), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriter11setFileNameERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, fileName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:79
// index:0
// QString fileName()
func (this *QImageWriter) FileName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter8fileNameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:81
// index:0
// void setQuality(int)
func (this *QImageWriter) SetQuality(quality int) {
	// 0: (, int quality), (&quality)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriter10setQualityEi", ffiqt.FFI_TYPE_VOID, this.cthis, &quality)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:82
// index:0
// int quality()
func (this *QImageWriter) Quality() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter7qualityEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:84
// index:0
// void setCompression(int)
func (this *QImageWriter) SetCompression(compression int) {
	// 0: (, int compression), (&compression)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriter14setCompressionEi", ffiqt.FFI_TYPE_VOID, this.cthis, &compression)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:85
// index:0
// int compression()
func (this *QImageWriter) Compression() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter11compressionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:87
// index:0
// void setGamma(float)
func (this *QImageWriter) SetGamma(gamma float32) {
	// 0: (, float gamma), (&gamma)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriter8setGammaEf", ffiqt.FFI_TYPE_VOID, this.cthis, &gamma)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:88
// index:0
// float gamma()
func (this *QImageWriter) Gamma() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter5gammaEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:90
// index:0
// void setSubType(const class QByteArray &)
func (this *QImageWriter) SetSubType(type_ unsafe.Pointer) {
	// 0: (, const QByteArray & type), (type_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriter10setSubTypeERK10QByteArray", ffiqt.FFI_TYPE_VOID, this.cthis, type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:91
// index:0
// QByteArray subType()
func (this *QImageWriter) SubType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter7subTypeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:92
// index:0
// QList<QByteArray> supportedSubTypes()
func (this *QImageWriter) SupportedSubTypes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter17supportedSubTypesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:94
// index:0
// void setOptimizedWrite(_Bool)
func (this *QImageWriter) SetOptimizedWrite(optimize bool) {
	// 0: (, bool optimize), (&optimize)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriter17setOptimizedWriteEb", ffiqt.FFI_TYPE_VOID, this.cthis, &optimize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:95
// index:0
// bool optimizedWrite()
func (this *QImageWriter) OptimizedWrite() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter14optimizedWriteEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:97
// index:0
// void setProgressiveScanWrite(_Bool)
func (this *QImageWriter) SetProgressiveScanWrite(progressive bool) {
	// 0: (, bool progressive), (&progressive)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriter23setProgressiveScanWriteEb", ffiqt.FFI_TYPE_VOID, this.cthis, &progressive)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:98
// index:0
// bool progressiveScanWrite()
func (this *QImageWriter) ProgressiveScanWrite() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter20progressiveScanWriteEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:100
// index:0
// QImageIOHandler::Transformations transformation()
func (this *QImageWriter) Transformation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter14transformationEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:104
// index:0
// void setDescription(const class QString &)
func (this *QImageWriter) SetDescription(description unsafe.Pointer) {
	// 0: (, const QString & description), (description)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriter14setDescriptionERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, description)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:105
// index:0
// QString description()
func (this *QImageWriter) Description() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter11descriptionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:107
// index:0
// void setText(const class QString &, const class QString &)
func (this *QImageWriter) SetText(key unsafe.Pointer, text unsafe.Pointer) {
	// 0: (, const QString & key, const QString & text), (key, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriter7setTextERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.cthis, key, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:109
// index:0
// bool canWrite()
func (this *QImageWriter) CanWrite() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter8canWriteEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:110
// index:0
// bool write(const class QImage &)
func (this *QImageWriter) Write(image unsafe.Pointer) {
	// 0: (, const QImage & image), (image)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriter5writeERK6QImage", ffiqt.FFI_TYPE_VOID, this.cthis, image)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:112
// index:0
// QImageWriter::ImageWriterError error()
func (this *QImageWriter) Error() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter5errorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:113
// index:0
// QString errorString()
func (this *QImageWriter) ErrorString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter11errorStringEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:115
// index:0
// bool supportsOption(class QImageIOHandler::ImageOption)
func (this *QImageWriter) SupportsOption(option int) {
	// 0: (, QImageIOHandler::ImageOption option), (&option)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter14supportsOptionEN15QImageIOHandler11ImageOptionE", ffiqt.FFI_TYPE_VOID, this.cthis, &option)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:117
// index:0
// static
// QList<QByteArray> supportedImageFormats()
func (this *QImageWriter) SupportedImageFormats() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriter21supportedImageFormatsEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QImageWriter_SupportedImageFormats() {
	// 0: (), ()
	var nilthis *QImageWriter
	nilthis.SupportedImageFormats()
}

// /usr/include/qt/QtGui/qimagewriter.h:118
// index:0
// static
// QList<QByteArray> supportedMimeTypes()
func (this *QImageWriter) SupportedMimeTypes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriter18supportedMimeTypesEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QImageWriter_SupportedMimeTypes() {
	// 0: (), ()
	var nilthis *QImageWriter
	nilthis.SupportedMimeTypes()
}

//  body block end
