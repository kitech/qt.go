package qtgui

// /usr/include/qt/QtGui/qimagewriter.h
// #include <qimagewriter.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 49
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
	*qtrt.CObject
}

func (this *QImageWriter) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QImageWriter) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQImageWriterFromPointer(cthis unsafe.Pointer) *QImageWriter {
	return &QImageWriter{&qtrt.CObject{cthis}}
}
func (*QImageWriter) NewFromPointer(cthis unsafe.Pointer) *QImageWriter {
	return NewQImageWriterFromPointer(cthis)
}

// /usr/include/qt/QtGui/qimagewriter.h:67
// index:0
// Public
// void QImageWriter()
func NewQImageWriter() *QImageWriter {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriterC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQImageWriterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qimagewriter.h:68
// index:1
// Public
// void QImageWriter(QIODevice *, const QByteArray &)
func NewQImageWriter_1(device *qtcore.QIODevice /*777 QIODevice **/, format *qtcore.QByteArray) *QImageWriter {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = device.GetCthis()
	var convArg1 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriterC2EP9QIODeviceRK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQImageWriterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qimagewriter.h:69
// index:2
// Public
// void QImageWriter(const QString &, const QByteArray &)
func NewQImageWriter_2(fileName *qtcore.QString, format *qtcore.QByteArray) *QImageWriter {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = fileName.GetCthis()
	var convArg1 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriterC2ERK7QStringRK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQImageWriterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qimagewriter.h:70
// index:0
// Public
// void ~QImageWriter()
func DeleteQImageWriter(*QImageWriter) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriterD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:72
// index:0
// Public
// void setFormat(const QByteArray &)
func (this *QImageWriter) SetFormat(format *qtcore.QByteArray) {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriter9setFormatERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:73
// index:0
// Public
// QByteArray format()
func (this *QImageWriter) Format() *qtcore.QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter6formatEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qimagewriter.h:75
// index:0
// Public
// void setDevice(QIODevice *)
func (this *QImageWriter) SetDevice(device *qtcore.QIODevice /*777 QIODevice **/) {
	var convArg0 = device.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriter9setDeviceEP9QIODevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:76
// index:0
// Public
// QIODevice * device()
func (this *QImageWriter) Device() *qtcore.QIODevice /*777 QIODevice **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter6deviceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qimagewriter.h:78
// index:0
// Public
// void setFileName(const QString &)
func (this *QImageWriter) SetFileName(fileName *qtcore.QString) {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriter11setFileNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:79
// index:0
// Public
// QString fileName()
func (this *QImageWriter) FileName() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter8fileNameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qimagewriter.h:81
// index:0
// Public
// void setQuality(int)
func (this *QImageWriter) SetQuality(quality int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriter10setQualityEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), quality)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:82
// index:0
// Public
// int quality()
func (this *QImageWriter) Quality() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter7qualityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qimagewriter.h:84
// index:0
// Public
// void setCompression(int)
func (this *QImageWriter) SetCompression(compression int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriter14setCompressionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), compression)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:85
// index:0
// Public
// int compression()
func (this *QImageWriter) Compression() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter11compressionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qimagewriter.h:87
// index:0
// Public
// void setGamma(float)
func (this *QImageWriter) SetGamma(gamma float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriter8setGammaEf", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), gamma)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:88
// index:0
// Public
// float gamma()
func (this *QImageWriter) Gamma() float32 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter5gammaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float32(rv) // 111
}

// /usr/include/qt/QtGui/qimagewriter.h:90
// index:0
// Public
// void setSubType(const QByteArray &)
func (this *QImageWriter) SetSubType(type_ *qtcore.QByteArray) {
	var convArg0 = type_.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriter10setSubTypeERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:91
// index:0
// Public
// QByteArray subType()
func (this *QImageWriter) SubType() *qtcore.QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter7subTypeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qimagewriter.h:94
// index:0
// Public
// void setOptimizedWrite(bool)
func (this *QImageWriter) SetOptimizedWrite(optimize bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriter17setOptimizedWriteEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), optimize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:95
// index:0
// Public
// bool optimizedWrite()
func (this *QImageWriter) OptimizedWrite() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter14optimizedWriteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qimagewriter.h:97
// index:0
// Public
// void setProgressiveScanWrite(bool)
func (this *QImageWriter) SetProgressiveScanWrite(progressive bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriter23setProgressiveScanWriteEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), progressive)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:98
// index:0
// Public
// bool progressiveScanWrite()
func (this *QImageWriter) ProgressiveScanWrite() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter20progressiveScanWriteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qimagewriter.h:100
// index:0
// Public
// QImageIOHandler::Transformations transformation()
func (this *QImageWriter) Transformation() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter14transformationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:104
// index:0
// Public
// void setDescription(const QString &)
func (this *QImageWriter) SetDescription(description *qtcore.QString) {
	var convArg0 = description.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriter14setDescriptionERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:105
// index:0
// Public
// QString description()
func (this *QImageWriter) Description() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter11descriptionEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qimagewriter.h:107
// index:0
// Public
// void setText(const QString &, const QString &)
func (this *QImageWriter) SetText(key *qtcore.QString, text *qtcore.QString) {
	var convArg0 = key.GetCthis()
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriter7setTextERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:109
// index:0
// Public
// bool canWrite()
func (this *QImageWriter) CanWrite() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter8canWriteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qimagewriter.h:110
// index:0
// Public
// bool write(const QImage &)
func (this *QImageWriter) Write(image *QImage) bool {
	var convArg0 = image.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageWriter5writeERK6QImage", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qimagewriter.h:112
// index:0
// Public
// QImageWriter::ImageWriterError error()
func (this *QImageWriter) Error() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter5errorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:113
// index:0
// Public
// QString errorString()
func (this *QImageWriter) ErrorString() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter11errorStringEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qimagewriter.h:115
// index:0
// Public
// bool supportsOption(QImageIOHandler::ImageOption)
func (this *QImageWriter) SupportsOption(option int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageWriter14supportsOptionEN15QImageIOHandler11ImageOptionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

type QImageWriter__ImageWriterError = int

const QImageWriter__UnknownError QImageWriter__ImageWriterError = 0
const QImageWriter__DeviceError QImageWriter__ImageWriterError = 1
const QImageWriter__UnsupportedFormatError QImageWriter__ImageWriterError = 2
const QImageWriter__InvalidImageError QImageWriter__ImageWriterError = 3

//  body block end
