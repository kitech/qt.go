package qtgui

// /usr/include/qt/QtGui/qimagewriter.h
// #include <qimagewriter.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 50
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
type QImageWriter struct {
	*qtrt.CObject
}
type QImageWriter_ITF interface {
	QImageWriter_PTR() *QImageWriter
}

func (ptr *QImageWriter) QImageWriter_PTR() *QImageWriter { return ptr }

func (this *QImageWriter) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QImageWriter) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQImageWriterFromPointer(cthis unsafe.Pointer) *QImageWriter {
	return &QImageWriter{&qtrt.CObject{cthis}}
}
func (*QImageWriter) NewFromPointer(cthis unsafe.Pointer) *QImageWriter {
	return NewQImageWriterFromPointer(cthis)
}

// /usr/include/qt/QtGui/qimagewriter.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QImageWriter()

/*
Constructs an empty QImageWriter object. Before writing, you must call setFormat() to set an image format, then setDevice() or setFileName().
*/
func NewQImageWriter() *QImageWriter {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriterC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageWriterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImageWriter)
	return gothis
}

// /usr/include/qt/QtGui/qimagewriter.h:68
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QImageWriter(QIODevice *, const QByteArray &)

/*
Constructs an empty QImageWriter object. Before writing, you must call setFormat() to set an image format, then setDevice() or setFileName().
*/
func NewQImageWriter_1(device qtcore.QIODevice_ITF /*777 QIODevice **/, format qtcore.QByteArray_ITF) *QImageWriter {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if format != nil && format.QByteArray_PTR() != nil {
		convArg1 = format.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriterC2EP9QIODeviceRK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageWriterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImageWriter)
	return gothis
}

// /usr/include/qt/QtGui/qimagewriter.h:69
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QImageWriter(const QString &, const QByteArray &)

/*
Constructs an empty QImageWriter object. Before writing, you must call setFormat() to set an image format, then setDevice() or setFileName().
*/
func NewQImageWriter_2(fileName string, format qtcore.QByteArray_ITF) *QImageWriter {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if format != nil && format.QByteArray_PTR() != nil {
		convArg1 = format.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriterC2ERK7QStringRK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageWriterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImageWriter)
	return gothis
}

// /usr/include/qt/QtGui/qimagewriter.h:69
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QImageWriter(const QString &, const QByteArray &)

/*
Constructs an empty QImageWriter object. Before writing, you must call setFormat() to set an image format, then setDevice() or setFileName().
*/
func NewQImageWriter_2_(fileName string) *QImageWriter {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QByteArray &=LValueReference, QByteArray=Record,
	var convArg1 = qtcore.NewQByteArray()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriterC2ERK7QStringRK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageWriterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImageWriter)
	return gothis
}

// /usr/include/qt/QtGui/qimagewriter.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QImageWriter()

/*

 */
func DeleteQImageWriter(this *QImageWriter) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriterD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qimagewriter.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFormat(const QByteArray &)

/*
Sets the format QImageWriter will use when writing images, to format. format is a case insensitive text string. Example:


  QImageWriter writer;
  writer.setFormat("png"); // same as writer.setFormat("PNG");



You can call supportedImageFormats() for the full list of formats QImageWriter supports.

See also format().
*/
func (this *QImageWriter) SetFormat(format qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QByteArray_PTR() != nil {
		convArg0 = format.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriter9setFormatERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray format() const

/*
Returns the format QImageWriter uses for writing images.

See also setFormat().
*/
func (this *QImageWriter) Format() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtGui/qimagewriter.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDevice(QIODevice *)

/*
Sets QImageWriter's device to device. If a device has already been set, the old device is removed from QImageWriter and is otherwise left unchanged.

If the device is not already open, QImageWriter will attempt to open the device in QIODevice::WriteOnly mode by calling open(). Note that this does not work for certain devices, such as QProcess, QTcpSocket and QUdpSocket, where more logic is required to open the device.

See also device() and setFileName().
*/
func (this *QImageWriter) SetDevice(device qtcore.QIODevice_ITF /*777 QIODevice **/) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriter9setDeviceEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QIODevice * device() const

/*
Returns the device currently assigned to QImageWriter, or 0 if no device has been assigned.

See also setDevice().
*/
func (this *QImageWriter) Device() *qtcore.QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter6deviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qimagewriter.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFileName(const QString &)

/*
Sets the file name of QImageWriter to fileName. Internally, QImageWriter will create a QFile and open it in QIODevice::WriteOnly mode, and use this file when writing images.

See also fileName() and setDevice().
*/
func (this *QImageWriter) SetFileName(fileName string) {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriter11setFileNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fileName() const

/*
If the currently assigned device is a QFile, or if setFileName() has been called, this function returns the name of the file QImageWriter writes to. Otherwise (i.e., if no device has been assigned or the device is not a QFile), an empty QString is returned.

See also setFileName() and setDevice().
*/
func (this *QImageWriter) FileName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter8fileNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qimagewriter.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setQuality(int)

/*
Sets the quality setting of the image format to quality.

Some image formats, in particular lossy ones, entail a tradeoff between a) visual quality of the resulting image, and b) encoding execution time and compression level. This function sets the level of that tradeoff for image formats that support it. For other formats, this value is ignored.

The value range of quality depends on the image format. For example, the "jpeg" format supports a quality range from 0 (low visual quality, high compression) to 100 (high visual quality, low compression).

See also quality().
*/
func (this *QImageWriter) SetQuality(quality int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriter10setQualityEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), quality)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] int quality() const

/*
Returns the quality setting of the image format.

See also setQuality().
*/
func (this *QImageWriter) Quality() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter7qualityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimagewriter.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCompression(int)

/*
This is an image format specific function that set the compression of an image. For image formats that do not support setting the compression, this value is ignored.

The value range of compression depends on the image format. For example, the "tiff" format supports two values, 0(no compression) and 1(LZW-compression).

See also compression().
*/
func (this *QImageWriter) SetCompression(compression int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriter14setCompressionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), compression)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:85
// index:0
// Public Visibility=Default Availability=Available
// [4] int compression() const

/*
Returns the compression of the image.

See also setCompression().
*/
func (this *QImageWriter) Compression() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter11compressionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimagewriter.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGamma(float)

/*
This is an image format specific function that sets the gamma level of the image to gamma. For image formats that do not support setting the gamma level, this value is ignored.

The value range of gamma depends on the image format. For example, the "png" format supports a gamma range from 0.0 to 1.0.

See also gamma() and quality().
*/
func (this *QImageWriter) SetGamma(gamma float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriter8setGammaEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gamma)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:88
// index:0
// Public Visibility=Default Availability=Available
// [4] float gamma() const

/*
Returns the gamma level of the image.

See also setGamma().
*/
func (this *QImageWriter) Gamma() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter5gammaEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qimagewriter.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSubType(const QByteArray &)

/*
This is an image format specific function that sets the subtype of the image to type. Subtype can be used by a handler to determine which format it should use while saving the image.

For example, saving an image in DDS format with A8R8G8R8 subtype:


  QImageWriter writer("some/image.dds");
  if (writer.supportsOption(QImageIOHandler::SubType))
      writer.setSubType("A8R8G8B8");
  writer.write(image);



This function was introduced in  Qt 5.4.

See also subType().
*/
func (this *QImageWriter) SetSubType(type_ qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if type_ != nil && type_.QByteArray_PTR() != nil {
		convArg0 = type_.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriter10setSubTypeERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:91
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray subType() const

/*
Returns the subtype of the image.

This function was introduced in  Qt 5.4.

See also setSubType().
*/
func (this *QImageWriter) SubType() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter7subTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtGui/qimagewriter.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOptimizedWrite(bool)

/*
This is an image format-specific function which sets the optimize flags when writing images. For image formats that do not support setting an optimize flag, this value is ignored.

The default is false.

This function was introduced in  Qt 5.5.

See also optimizedWrite().
*/
func (this *QImageWriter) SetOptimizedWrite(optimize bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriter17setOptimizedWriteEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), optimize)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:95
// index:0
// Public Visibility=Default Availability=Available
// [1] bool optimizedWrite() const

/*
Returns whether optimization has been turned on for writing the image.

This function was introduced in  Qt 5.5.

See also setOptimizedWrite().
*/
func (this *QImageWriter) OptimizedWrite() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter14optimizedWriteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimagewriter.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProgressiveScanWrite(bool)

/*
This is an image format-specific function which turns on progressive scanning when writing images. For image formats that do not support setting a progressive scan flag, this value is ignored.

The default is false.

This function was introduced in  Qt 5.5.

See also progressiveScanWrite().
*/
func (this *QImageWriter) SetProgressiveScanWrite(progressive bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriter23setProgressiveScanWriteEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), progressive)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:98
// index:0
// Public Visibility=Default Availability=Available
// [1] bool progressiveScanWrite() const

/*
Returns whether the image should be written as a progressive image.

This function was introduced in  Qt 5.5.

See also setProgressiveScanWrite().
*/
func (this *QImageWriter) ProgressiveScanWrite() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter20progressiveScanWriteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimagewriter.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] QImageIOHandler::Transformations transformation() const

/*
Returns the transformation and orientation the image has been set to written with.

This function was introduced in  Qt 5.5.

See also setTransformation().
*/
func (this *QImageWriter) Transformation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter14transformationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTransformation(QImageIOHandler::Transformations)

/*
Sets the image transformations metadata including orientation to transform.

If transformation metadata is not supported by the image format, the transform is applied before writing.

This function was introduced in  Qt 5.5.

See also transformation() and write().
*/
func (this *QImageWriter) SetTransformation(orientation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriter17setTransformationE6QFlagsIN15QImageIOHandler14TransformationEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDescription(const QString &)

/*

 */
func (this *QImageWriter) SetDescription(description string) {
	var tmpArg0 = qtcore.NewQString_5(description)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriter14setDescriptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] QString description() const

/*

 */
func (this *QImageWriter) Description() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter11descriptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qimagewriter.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &, const QString &)

/*
Sets the image text associated with the key key to text. This is useful for storing copyright information or other information about the image. Example:


  QImage image("some/image.jpeg");
  QImageWriter writer("images/outimage.png", "png");
  writer.setText("Author", "John Smith");
  writer.write(image);



If you want to store a single block of data (e.g., a comment), you can pass an empty key, or use a generic key like "Description".

The key and text will be embedded into the image data after calling write().

Support for this option is implemented through QImageIOHandler::Description.

This function was introduced in  Qt 4.1.

See also QImage::setText() and QImageReader::text().
*/
func (this *QImageWriter) SetText(key string, text string) {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriter7setTextERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:109
// index:0
// Public Visibility=Default Availability=Available
// [1] bool canWrite() const

/*
Returns true if QImageWriter can write the image; i.e., the image format is supported and the assigned device is open for reading.

See also write(), setDevice(), and setFormat().
*/
func (this *QImageWriter) CanWrite() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter8canWriteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimagewriter.h:110
// index:0
// Public Visibility=Default Availability=Available
// [1] bool write(const QImage &)

/*
Writes the image image to the assigned device or file name. Returns true on success; otherwise returns false. If the operation fails, you can call error() to find the type of error that occurred, or errorString() to get a human readable description of the error.

See also canWrite(), error(), and errorString().
*/
func (this *QImageWriter) Write(image QImage_ITF) bool {
	var convArg0 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg0 = image.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriter5writeERK6QImage", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimagewriter.h:112
// index:0
// Public Visibility=Default Availability=Available
// [4] QImageWriter::ImageWriterError error() const

/*
Returns the type of error that last occurred.

See also ImageWriterError and errorString().
*/
func (this *QImageWriter) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:113
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const

/*
Returns a human readable description of the last error that occurred.

See also error().
*/
func (this *QImageWriter) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qimagewriter.h:115
// index:0
// Public Visibility=Default Availability=Available
// [1] bool supportsOption(QImageIOHandler::ImageOption) const

/*
Returns true if the writer supports option; otherwise returns false.

Different image formats support different options. Call this function to determine whether a certain option is supported by the current format. For example, the PNG format allows you to embed text into the image's metadata (see text()).


  QImageWriter writer(fileName);
  if (writer.supportsOption(QImageIOHandler::Description))
      writer.setText("Author", "John Smith");



Options can be tested after the writer has been associated with a format.

This function was introduced in  Qt 4.2.

See also QImageReader::supportsOption() and setFormat().
*/
func (this *QImageWriter) SupportsOption(option int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter14supportsOptionEN15QImageIOHandler11ImageOptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*
This enum describes errors that can occur when writing images with QImageWriter.


*/
type QImageWriter__ImageWriterError = int

// An unknown error occurred. If you get this value after calling write(), it is most likely caused by a bug in QImageWriter.
const QImageWriter__UnknownError QImageWriter__ImageWriterError = 0

// QImageWriter encountered a device error when writing the image data. Consult your device for more details on what went wrong.
const QImageWriter__DeviceError QImageWriter__ImageWriterError = 1

// Qt does not support the requested image format.
const QImageWriter__UnsupportedFormatError QImageWriter__ImageWriterError = 2

// An attempt was made to write an invalid QImage. An example of an invalid image would be a null QImage.
const QImageWriter__InvalidImageError QImageWriter__ImageWriterError = 3

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
