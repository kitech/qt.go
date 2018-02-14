package qtgui

// /usr/include/qt/QtGui/qimagereader.h
// #include <qimagereader.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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

type QImageReader struct {
	*qtrt.CObject
}
type QImageReader_ITF interface {
	QImageReader_PTR() *QImageReader
}

func (ptr *QImageReader) QImageReader_PTR() *QImageReader { return ptr }

func (this *QImageReader) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QImageReader) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQImageReaderFromPointer(cthis unsafe.Pointer) *QImageReader {
	return &QImageReader{&qtrt.CObject{cthis}}
}
func (*QImageReader) NewFromPointer(cthis unsafe.Pointer) *QImageReader {
	return NewQImageReaderFromPointer(cthis)
}

// /usr/include/qt/QtGui/qimagereader.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QImageReader()
func NewQImageReader() *QImageReader {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReaderC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageReaderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImageReader)
	return gothis
}

// /usr/include/qt/QtGui/qimagereader.h:72
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QImageReader(QIODevice *, const QByteArray &)
func NewQImageReader_1(device qtcore.QIODevice_ITF /*777 QIODevice **/, format qtcore.QByteArray_ITF) *QImageReader {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if format != nil && format.QByteArray_PTR() != nil {
		convArg1 = format.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReaderC2EP9QIODeviceRK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageReaderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImageReader)
	return gothis
}

// /usr/include/qt/QtGui/qimagereader.h:73
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QImageReader(const QString &, const QByteArray &)
func NewQImageReader_2(fileName string, format qtcore.QByteArray_ITF) *QImageReader {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if format != nil && format.QByteArray_PTR() != nil {
		convArg1 = format.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReaderC2ERK7QStringRK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageReaderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImageReader)
	return gothis
}

// /usr/include/qt/QtGui/qimagereader.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QImageReader()
func DeleteQImageReader(this *QImageReader) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReaderD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qimagereader.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFormat(const QByteArray &)
func (this *QImageReader) SetFormat(format qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QByteArray_PTR() != nil {
		convArg0 = format.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReader9setFormatERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray format()
func (this *QImageReader) Format() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtGui/qimagereader.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoDetectImageFormat(_Bool)
func (this *QImageReader) SetAutoDetectImageFormat(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReader24setAutoDetectImageFormatEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoDetectImageFormat()
func (this *QImageReader) AutoDetectImageFormat() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader21autoDetectImageFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimagereader.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDecideFormatFromContent(_Bool)
func (this *QImageReader) SetDecideFormatFromContent(ignored bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReader26setDecideFormatFromContentEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ignored)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:83
// index:0
// Public Visibility=Default Availability=Available
// [1] bool decideFormatFromContent()
func (this *QImageReader) DecideFormatFromContent() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader23decideFormatFromContentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimagereader.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDevice(QIODevice *)
func (this *QImageReader) SetDevice(device qtcore.QIODevice_ITF /*777 QIODevice **/) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReader9setDeviceEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:86
// index:0
// Public Visibility=Default Availability=Available
// [8] QIODevice * device()
func (this *QImageReader) Device() *qtcore.QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader6deviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qimagereader.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFileName(const QString &)
func (this *QImageReader) SetFileName(fileName string) {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReader11setFileNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:89
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fileName()
func (this *QImageReader) FileName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader8fileNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qimagereader.h:91
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize size()
func (this *QImageReader) Size() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qimagereader.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] QImage::Format imageFormat()
func (this *QImageReader) ImageFormat() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader11imageFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qimagereader.h:143
// index:1
// Public static Visibility=Default Availability=Available
// [8] QByteArray imageFormat(const QString &)
func (this *QImageReader) ImageFormat_1(fileName string) *qtcore.QByteArray /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReader11imageFormatERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}
func QImageReader_ImageFormat_1(fileName string) *qtcore.QByteArray /*123*/ {
	var nilthis *QImageReader
	rv := nilthis.ImageFormat_1(fileName)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:144
// index:2
// Public static Visibility=Default Availability=Available
// [8] QByteArray imageFormat(QIODevice *)
func (this *QImageReader) ImageFormat_2(device qtcore.QIODevice_ITF /*777 QIODevice **/) *qtcore.QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReader11imageFormatEP9QIODevice", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}
func QImageReader_ImageFormat_2(device qtcore.QIODevice_ITF /*777 QIODevice **/) *qtcore.QByteArray /*123*/ {
	var nilthis *QImageReader
	rv := nilthis.ImageFormat_2(device)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:95
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList textKeys()
func (this *QImageReader) TextKeys() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader8textKeysEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtGui/qimagereader.h:96
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text(const QString &)
func (this *QImageReader) Text(key string) string {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader4textERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qimagereader.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClipRect(const QRect &)
func (this *QImageReader) SetClipRect(rect qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReader11setClipRectERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:99
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect clipRect()
func (this *QImageReader) ClipRect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader8clipRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qimagereader.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScaledSize(const QSize &)
func (this *QImageReader) SetScaledSize(size qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReader13setScaledSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:102
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize scaledSize()
func (this *QImageReader) ScaledSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader10scaledSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qimagereader.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setQuality(int)
func (this *QImageReader) SetQuality(quality int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReader10setQualityEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), quality)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:105
// index:0
// Public Visibility=Default Availability=Available
// [4] int quality()
func (this *QImageReader) Quality() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader7qualityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimagereader.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScaledClipRect(const QRect &)
func (this *QImageReader) SetScaledClipRect(rect qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReader17setScaledClipRectERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:108
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect scaledClipRect()
func (this *QImageReader) ScaledClipRect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader14scaledClipRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qimagereader.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBackgroundColor(const QColor &)
func (this *QImageReader) SetBackgroundColor(color QColor_ITF) {
	var convArg0 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg0 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReader18setBackgroundColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:111
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor backgroundColor()
func (this *QImageReader) BackgroundColor() *QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader15backgroundColorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}

// /usr/include/qt/QtGui/qimagereader.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool supportsAnimation()
func (this *QImageReader) SupportsAnimation() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader17supportsAnimationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimagereader.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] QImageIOHandler::Transformations transformation()
func (this *QImageReader) Transformation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader14transformationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qimagereader.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoTransform(_Bool)
func (this *QImageReader) SetAutoTransform(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReader16setAutoTransformEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:118
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoTransform()
func (this *QImageReader) AutoTransform() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader13autoTransformEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimagereader.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGamma(float)
func (this *QImageReader) SetGamma(gamma float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReader8setGammaEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gamma)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:121
// index:0
// Public Visibility=Default Availability=Available
// [4] float gamma()
func (this *QImageReader) Gamma() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader5gammaEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qimagereader.h:123
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray subType()
func (this *QImageReader) SubType() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader7subTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtGui/qimagereader.h:126
// index:0
// Public Visibility=Default Availability=Available
// [1] bool canRead()
func (this *QImageReader) CanRead() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader7canReadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimagereader.h:127
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage read()
func (this *QImageReader) Read() *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReader4readEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimagereader.h:128
// index:1
// Public Visibility=Default Availability=Available
// [1] bool read(QImage *)
func (this *QImageReader) Read_1(image QImage_ITF /*777 QImage **/) bool {
	var convArg0 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg0 = image.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReader4readEP6QImage", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimagereader.h:130
// index:0
// Public Visibility=Default Availability=Available
// [1] bool jumpToNextImage()
func (this *QImageReader) JumpToNextImage() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReader15jumpToNextImageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimagereader.h:131
// index:0
// Public Visibility=Default Availability=Available
// [1] bool jumpToImage(int)
func (this *QImageReader) JumpToImage(imageNumber int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReader11jumpToImageEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), imageNumber)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimagereader.h:132
// index:0
// Public Visibility=Default Availability=Available
// [4] int loopCount()
func (this *QImageReader) LoopCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader9loopCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimagereader.h:133
// index:0
// Public Visibility=Default Availability=Available
// [4] int imageCount()
func (this *QImageReader) ImageCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader10imageCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimagereader.h:134
// index:0
// Public Visibility=Default Availability=Available
// [4] int nextImageDelay()
func (this *QImageReader) NextImageDelay() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader14nextImageDelayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimagereader.h:135
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentImageNumber()
func (this *QImageReader) CurrentImageNumber() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader18currentImageNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimagereader.h:136
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect currentImageRect()
func (this *QImageReader) CurrentImageRect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader16currentImageRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qimagereader.h:138
// index:0
// Public Visibility=Default Availability=Available
// [4] QImageReader::ImageReaderError error()
func (this *QImageReader) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qimagereader.h:139
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString()
func (this *QImageReader) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qimagereader.h:141
// index:0
// Public Visibility=Default Availability=Available
// [1] bool supportsOption(QImageIOHandler::ImageOption)
func (this *QImageReader) SupportsOption(option int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader14supportsOptionEN15QImageIOHandler11ImageOptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

type QImageReader__ImageReaderError = int

const QImageReader__UnknownError QImageReader__ImageReaderError = 0
const QImageReader__FileNotFoundError QImageReader__ImageReaderError = 1
const QImageReader__DeviceError QImageReader__ImageReaderError = 2
const QImageReader__UnsupportedFormatError QImageReader__ImageReaderError = 3
const QImageReader__InvalidDataError QImageReader__ImageReaderError = 4

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
