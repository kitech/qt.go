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

/*

 */
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

/*
Constructs an empty QImageReader object. Before reading an image, call setDevice() or setFileName().
*/
func (*QImageReader) NewForInherit() *QImageReader {
	return NewQImageReader()
}
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

/*
Constructs an empty QImageReader object. Before reading an image, call setDevice() or setFileName().
*/
func (*QImageReader) NewForInherit_1(device qtcore.QIODevice_ITF /*777 QIODevice **/, format qtcore.QByteArray_ITF) *QImageReader {
	return NewQImageReader_1(device, format)
}
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

// /usr/include/qt/QtGui/qimagereader.h:72
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QImageReader(QIODevice *, const QByteArray &)

/*
Constructs an empty QImageReader object. Before reading an image, call setDevice() or setFileName().
*/
func (*QImageReader) NewForInherit_1_(device qtcore.QIODevice_ITF /*777 QIODevice **/) *QImageReader {
	return NewQImageReader_1_(device)
}
func NewQImageReader_1_(device qtcore.QIODevice_ITF /*777 QIODevice **/) *QImageReader {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	// arg: 1, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
	var convArg1 = qtcore.NewQByteArray()
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

/*
Constructs an empty QImageReader object. Before reading an image, call setDevice() or setFileName().
*/
func (*QImageReader) NewForInherit_2(fileName string, format qtcore.QByteArray_ITF) *QImageReader {
	return NewQImageReader_2(fileName, format)
}
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

// /usr/include/qt/QtGui/qimagereader.h:73
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QImageReader(const QString &, const QByteArray &)

/*
Constructs an empty QImageReader object. Before reading an image, call setDevice() or setFileName().
*/
func (*QImageReader) NewForInherit_2_(fileName string) *QImageReader {
	return NewQImageReader_2_(fileName)
}
func NewQImageReader_2_(fileName string) *QImageReader {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
	var convArg1 = qtcore.NewQByteArray()
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

/*

 */
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

/*
Sets the format QImageReader will use when reading images, to format. format is a case insensitive text string. Example:


  QImageReader reader;
  reader.setFormat("png"); // same as reader.setFormat("PNG");



You can call supportedImageFormats() for the full list of formats QImageReader supports.

See also format().
*/
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
// [8] QByteArray format() const

/*
Returns the format QImageReader uses for reading images.

You can call this function after assigning a device to the reader to determine the format of the device. For example:


  QImageReader reader("image.png");
  // reader.format() == "png"



If the reader cannot read any image from the device (e.g., there is no image there, or the image has already been read), or if the format is unsupported, this function returns an empty QByteArray().

See also setFormat() and supportedImageFormats().
*/
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
// [-2] void setAutoDetectImageFormat(bool)

/*
If enabled is true, image format autodetection is enabled; otherwise, it is disabled. By default, autodetection is enabled.

QImageReader uses an extensive approach to detecting the image format; firstly, if you pass a file name to QImageReader, it will attempt to detect the file extension if the given file name does not point to an existing file, by appending supported default extensions to the given file name, one at a time. It then uses the following approach to detect the image format:


Image plugins are queried first, based on either the optional format string, or the file name suffix (if the source device is a file). No content detection is done at this stage. QImageReader will choose the first plugin that supports reading for this format.
If no plugin supports the image format, Qt's built-in handlers are checked based on either the optional format string, or the file name suffix.
If no capable plugins or built-in handlers are found, each plugin is tested by inspecting the content of the data stream.
If no plugins could detect the image format based on data contents, each built-in image handler is tested by inspecting the contents.
Finally, if all above approaches fail, QImageReader will report failure when trying to read the image.


By disabling image format autodetection, QImageReader will only query the plugins and built-in handlers based on the format string (i.e., no file name extensions are tested).

See also autoDetectImageFormat(), QImageIOHandler::canRead(), and QImageIOPlugin::capabilities().
*/
func (this *QImageReader) SetAutoDetectImageFormat(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReader24setAutoDetectImageFormatEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoDetectImageFormat() const

/*
Returns true if image format autodetection is enabled on this image reader; otherwise returns false. By default, autodetection is enabled.

See also setAutoDetectImageFormat().
*/
func (this *QImageReader) AutoDetectImageFormat() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader21autoDetectImageFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimagereader.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDecideFormatFromContent(bool)

/*
If ignored is set to true, then the image reader will ignore specified formats or file extensions and decide which plugin to use only based on the contents in the datastream.

Setting this flag means that all image plugins gets loaded. Each plugin will read the first bytes in the image data and decide if the plugin is compatible or not.

This also disables auto detecting the image format.

See also decideFormatFromContent().
*/
func (this *QImageReader) SetDecideFormatFromContent(ignored bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReader26setDecideFormatFromContentEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ignored)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:83
// index:0
// Public Visibility=Default Availability=Available
// [1] bool decideFormatFromContent() const

/*
Returns whether the image reader should decide which plugin to use only based on the contents of the datastream rather than on the file extension.

See also setDecideFormatFromContent().
*/
func (this *QImageReader) DecideFormatFromContent() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader23decideFormatFromContentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimagereader.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDevice(QIODevice *)

/*
Sets QImageReader's device to device. If a device has already been set, the old device is removed from QImageReader and is otherwise left unchanged.

If the device is not already open, QImageReader will attempt to open the device in QIODevice::ReadOnly mode by calling open(). Note that this does not work for certain devices, such as QProcess, QTcpSocket and QUdpSocket, where more logic is required to open the device.

See also device() and setFileName().
*/
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
// [8] QIODevice * device() const

/*
Returns the device currently assigned to QImageReader, or 0 if no device has been assigned.

See also setDevice().
*/
func (this *QImageReader) Device() *qtcore.QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader6deviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qimagereader.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFileName(const QString &)

/*
Sets the file name of QImageReader to fileName. Internally, QImageReader will create a QFile object and open it in QIODevice::ReadOnly mode, and use this when reading images.

If fileName does not include a file extension (e.g., .png or .bmp), QImageReader will cycle through all supported extensions until it finds a matching file.

See also fileName(), setDevice(), and supportedImageFormats().
*/
func (this *QImageReader) SetFileName(fileName string) {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReader11setFileNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:89
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fileName() const

/*
If the currently assigned device is a QFile, or if setFileName() has been called, this function returns the name of the file QImageReader reads from. Otherwise (i.e., if no device has been assigned or the device is not a QFile), an empty QString is returned.

See also setFileName() and setDevice().
*/
func (this *QImageReader) FileName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader8fileNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qimagereader.h:91
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize size() const

/*
Returns the size of the image, without actually reading the image contents.

If the image format does not support this feature, this function returns an invalid size. Qt's built-in image handlers all support this feature, but custom image format plugins are not required to do so.

See also QImageIOHandler::ImageOption, QImageIOHandler::option(), and QImageIOHandler::supportsOption().
*/
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
// [4] QImage::Format imageFormat() const

/*
Returns the format of the image, without actually reading the image contents. The format describes the image format QImageReader::read() returns, not the format of the actual image.

If the image format does not support this feature, this function returns an invalid format.

This function was introduced in  Qt 4.5.

See also QImageIOHandler::ImageOption, QImageIOHandler::option(), and QImageIOHandler::supportsOption().
*/
func (this *QImageReader) ImageFormat() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader11imageFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qimagereader.h:143
// index:1
// Public static Visibility=Default Availability=Available
// [8] QByteArray imageFormat(const QString &)

/*
Returns the format of the image, without actually reading the image contents. The format describes the image format QImageReader::read() returns, not the format of the actual image.

If the image format does not support this feature, this function returns an invalid format.

This function was introduced in  Qt 4.5.

See also QImageIOHandler::ImageOption, QImageIOHandler::option(), and QImageIOHandler::supportsOption().
*/
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

/*
Returns the format of the image, without actually reading the image contents. The format describes the image format QImageReader::read() returns, not the format of the actual image.

If the image format does not support this feature, this function returns an invalid format.

This function was introduced in  Qt 4.5.

See also QImageIOHandler::ImageOption, QImageIOHandler::option(), and QImageIOHandler::supportsOption().
*/
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
// [8] QStringList textKeys() const

/*
Returns the text keys for this image. You can use these keys with text() to list the image text for a certain key.

Support for this option is implemented through QImageIOHandler::Description.

This function was introduced in  Qt 4.1.

See also text(), QImageWriter::setText(), and QImage::textKeys().
*/
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
// [8] QString text(const QString &) const

/*
Returns the image text associated with key.

Support for this option is implemented through QImageIOHandler::Description.

This function was introduced in  Qt 4.1.

See also textKeys() and QImageWriter::setText().
*/
func (this *QImageReader) Text(key string) string {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader4textERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qimagereader.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClipRect(const QRect &)

/*
Sets the image clip rect (also known as the ROI, or Region Of Interest) to rect. The coordinates of rect are relative to the untransformed image size, as returned by size().

See also clipRect(), setScaledSize(), and setScaledClipRect().
*/
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
// [16] QRect clipRect() const

/*
Returns the clip rect (also known as the ROI, or Region Of Interest) of the image. If no clip rect has been set, an invalid QRect is returned.

See also setClipRect().
*/
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

/*
Sets the scaled size of the image to size. The scaling is performed after the initial clip rect, but before the scaled clip rect is applied. The algorithm used for scaling depends on the image format. By default (i.e., if the image format does not support scaling), QImageReader will use QImage::scale() with Qt::SmoothScaling.

See also scaledSize(), setClipRect(), and setScaledClipRect().
*/
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
// [8] QSize scaledSize() const

/*
Returns the scaled size of the image.

See also setScaledSize().
*/
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

/*
Sets the quality setting of the image format to quality.

Some image formats, in particular lossy ones, entail a tradeoff between a) visual quality of the resulting image, and b) decoding execution time. This function sets the level of that tradeoff for image formats that support it.

In case of scaled image reading, the quality setting may also influence the tradeoff level between visual quality and execution speed of the scaling algorithm.

The value range of quality depends on the image format. For example, the "jpeg" format supports a quality range from 0 (low visual quality) to 100 (high visual quality).

This function was introduced in  Qt 4.2.

See also quality() and setScaledSize().
*/
func (this *QImageReader) SetQuality(quality int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReader10setQualityEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), quality)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:105
// index:0
// Public Visibility=Default Availability=Available
// [4] int quality() const

/*
Returns the quality setting of the image format.

This function was introduced in  Qt 4.2.

See also setQuality().
*/
func (this *QImageReader) Quality() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader7qualityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimagereader.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScaledClipRect(const QRect &)

/*
Sets the scaled clip rect to rect. The scaled clip rect is the clip rect (also known as ROI, or Region Of Interest) that is applied after the image has been scaled.

See also scaledClipRect() and setScaledSize().
*/
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
// [16] QRect scaledClipRect() const

/*
Returns the scaled clip rect of the image.

See also setScaledClipRect().
*/
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

/*
Sets the background color to color. Image formats that support this operation are expected to initialize the background to color before reading an image.

This function was introduced in  Qt 4.1.

See also backgroundColor() and read().
*/
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
// [16] QColor backgroundColor() const

/*
Returns the background color that's used when reading an image. If the image format does not support setting the background color an invalid color is returned.

This function was introduced in  Qt 4.1.

See also setBackgroundColor() and read().
*/
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
// [1] bool supportsAnimation() const

/*
Returns true if the image format supports animation; otherwise, false is returned.

This function was introduced in  Qt 4.1.

See also QMovie::supportedFormats().
*/
func (this *QImageReader) SupportsAnimation() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader17supportsAnimationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimagereader.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] QImageIOHandler::Transformations transformation() const

/*
Returns the transformation metadata of the image, including image orientation. If the format does not support transformation metadata QImageIOHandler::Transformation_None is returned.

This function was introduced in  Qt 5.5.

See also setAutoTransform() and autoTransform().
*/
func (this *QImageReader) Transformation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader14transformationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qimagereader.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoTransform(bool)

/*
Determines that images returned by read() should have transformation metadata automatically applied if enabled is true.

This function was introduced in  Qt 5.5.

See also autoTransform(), transformation(), and read().
*/
func (this *QImageReader) SetAutoTransform(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReader16setAutoTransformEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:118
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoTransform() const

/*
Returns true if the image handler will apply transformation metadata on read().

This function was introduced in  Qt 5.5.

See also setAutoTransform(), transformation(), and read().
*/
func (this *QImageReader) AutoTransform() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader13autoTransformEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimagereader.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGamma(float)

/*
This is an image format specific function that forces images with gamma information to be gamma corrected to gamma. For image formats that do not support gamma correction, this value is ignored.

To gamma correct to a standard PC color-space, set gamma to 1/2.2.

This function was introduced in  Qt 5.6.

See also gamma().
*/
func (this *QImageReader) SetGamma(gamma float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReader8setGammaEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gamma)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:121
// index:0
// Public Visibility=Default Availability=Available
// [4] float gamma() const

/*
Returns the gamma level of the decoded image. If setGamma() has been called and gamma correction is supported it will return the gamma set. If gamma level is not supported by the image format, 0.0 is returned.

This function was introduced in  Qt 5.6.

See also setGamma().
*/
func (this *QImageReader) Gamma() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader5gammaEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qimagereader.h:123
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray subType() const

/*
Returns the subtype of the image.

This function was introduced in  Qt 5.4.
*/
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
// [1] bool canRead() const

/*
Returns true if an image can be read for the device (i.e., the image format is supported, and the device seems to contain valid data); otherwise returns false.

canRead() is a lightweight function that only does a quick test to see if the image data is valid. read() may still return false after canRead() returns true, if the image data is corrupt.

Note: A QMimeDatabase lookup is normally a better approach than this function for identifying potentially non-image files or data.

For images that support animation, canRead() returns false when all frames have been read.

See also read(), supportedImageFormats(), and QMimeDatabase.
*/
func (this *QImageReader) CanRead() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader7canReadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimagereader.h:127
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage read()

/*
Reads an image from the device. On success, the image that was read is returned; otherwise, a null QImage is returned. You can then call error() to find the type of error that occurred, or errorString() to get a human readable description of the error.

For image formats that support animation, calling read() repeatedly will return the next frame. When all frames have been read, a null image will be returned.

See also canRead(), supportedImageFormats(), supportsAnimation(), and QMovie.
*/
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

/*
Reads an image from the device. On success, the image that was read is returned; otherwise, a null QImage is returned. You can then call error() to find the type of error that occurred, or errorString() to get a human readable description of the error.

For image formats that support animation, calling read() repeatedly will return the next frame. When all frames have been read, a null image will be returned.

See also canRead(), supportedImageFormats(), supportsAnimation(), and QMovie.
*/
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

/*
For image formats that support animation, this function steps over the current image, returning true if successful or false if there is no following image in the animation.

The default implementation calls read(), then discards the resulting image, but the image handler may have a more efficient way of implementing this operation.

See also jumpToImage() and QImageIOHandler::jumpToNextImage().
*/
func (this *QImageReader) JumpToNextImage() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReader15jumpToNextImageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimagereader.h:131
// index:0
// Public Visibility=Default Availability=Available
// [1] bool jumpToImage(int)

/*
For image formats that support animation, this function skips to the image whose sequence number is imageNumber, returning true if successful or false if the corresponding image cannot be found.

The next call to read() will attempt to read this image.

See also jumpToNextImage() and QImageIOHandler::jumpToImage().
*/
func (this *QImageReader) JumpToImage(imageNumber int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageReader11jumpToImageEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), imageNumber)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimagereader.h:132
// index:0
// Public Visibility=Default Availability=Available
// [4] int loopCount() const

/*
For image formats that support animation, this function returns the number of times the animation should loop. If this function returns -1, it can either mean the animation should loop forever, or that an error occurred. If an error occurred, canRead() will return false.

See also supportsAnimation(), QImageIOHandler::loopCount(), and canRead().
*/
func (this *QImageReader) LoopCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader9loopCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimagereader.h:133
// index:0
// Public Visibility=Default Availability=Available
// [4] int imageCount() const

/*
For image formats that support animation, this function returns the total number of images in the animation. If the format does not support animation, 0 is returned.

This function returns -1 if an error occurred.

See also supportsAnimation(), QImageIOHandler::imageCount(), and canRead().
*/
func (this *QImageReader) ImageCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader10imageCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimagereader.h:134
// index:0
// Public Visibility=Default Availability=Available
// [4] int nextImageDelay() const

/*
For image formats that support animation, this function returns the number of milliseconds to wait until displaying the next frame in the animation. If the image format doesn't support animation, 0 is returned.

This function returns -1 if an error occurred.

See also supportsAnimation(), QImageIOHandler::nextImageDelay(), and canRead().
*/
func (this *QImageReader) NextImageDelay() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader14nextImageDelayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimagereader.h:135
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentImageNumber() const

/*
For image formats that support animation, this function returns the sequence number of the current frame. If the image format doesn't support animation, 0 is returned.

This function returns -1 if an error occurred.

See also supportsAnimation(), QImageIOHandler::currentImageNumber(), and canRead().
*/
func (this *QImageReader) CurrentImageNumber() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader18currentImageNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimagereader.h:136
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect currentImageRect() const

/*
For image formats that support animation, this function returns the rect for the current frame. Otherwise, a null rect is returned.

See also supportsAnimation() and QImageIOHandler::currentImageRect().
*/
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
// [4] QImageReader::ImageReaderError error() const

/*
Returns the type of error that occurred last.

See also ImageReaderError and errorString().
*/
func (this *QImageReader) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qimagereader.h:139
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const

/*
Returns a human readable description of the last error that occurred.

See also error().
*/
func (this *QImageReader) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qimagereader.h:141
// index:0
// Public Visibility=Default Availability=Available
// [1] bool supportsOption(QImageIOHandler::ImageOption) const

/*
Returns true if the reader supports option; otherwise returns false.

Different image formats support different options. Call this function to determine whether a certain option is supported by the current format. For example, the PNG format allows you to embed text into the image's metadata (see text()), and the BMP format allows you to determine the image's size without loading the whole image into memory (see size()).


  QImageReader reader(":/image.png");
  if (reader.supportsOption(QImageIOHandler::Size))
      qDebug() << "Size:" << reader.size();



This function was introduced in  Qt 4.2.

See also QImageWriter::supportsOption().
*/
func (this *QImageReader) SupportsOption(option int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageReader14supportsOptionEN15QImageIOHandler11ImageOptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*
This enum describes the different types of errors that can occur when reading images with QImageReader.


*/
type QImageReader__ImageReaderError = int

// An unknown error occurred. If you get this value after calling read(), it is most likely caused by a bug in QImageReader.
const QImageReader__UnknownError QImageReader__ImageReaderError = 0

// QImageReader was used with a file name, but not file was found with that name. This can also happen if the file name contained no extension, and the file with the correct extension is not supported by Qt.
const QImageReader__FileNotFoundError QImageReader__ImageReaderError = 1

// QImageReader encountered a device error when reading the image. You can consult your particular device for more details on what went wrong.
const QImageReader__DeviceError QImageReader__ImageReaderError = 2

// Qt does not support the requested image format.
const QImageReader__UnsupportedFormatError QImageReader__ImageReaderError = 3

// The image data was invalid, and QImageReader was unable to read an image from it. The can happen if the image file is damaged.
const QImageReader__InvalidDataError QImageReader__ImageReaderError = 4

func (this *QImageReader) ImageReaderErrorItemName(val int) string {
	switch val {
	case QImageReader__UnknownError: // 0
		return "UnknownError"
	case QImageReader__FileNotFoundError: // 1
		return "FileNotFoundError"
	case QImageReader__DeviceError: // 2
		return "DeviceError"
	case QImageReader__UnsupportedFormatError: // 3
		return "UnsupportedFormatError"
	case QImageReader__InvalidDataError: // 4
		return "InvalidDataError"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QImageReader_ImageReaderErrorItemName(val int) string {
	var nilthis *QImageReader
	return nilthis.ImageReaderErrorItemName(val)
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
