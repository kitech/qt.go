package qtgui

// /usr/include/qt/QtGui/qimageiohandler.h
// #include <qimageiohandler.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
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
type QImageIOHandler struct {
	*qtrt.CObject
}
type QImageIOHandler_ITF interface {
	QImageIOHandler_PTR() *QImageIOHandler
}

func (ptr *QImageIOHandler) QImageIOHandler_PTR() *QImageIOHandler { return ptr }

func (this *QImageIOHandler) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QImageIOHandler) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQImageIOHandlerFromPointer(cthis unsafe.Pointer) *QImageIOHandler {
	return &QImageIOHandler{&qtrt.CObject{cthis}}
}
func (*QImageIOHandler) NewFromPointer(cthis unsafe.Pointer) *QImageIOHandler {
	return NewQImageIOHandlerFromPointer(cthis)
}

// /usr/include/qt/QtGui/qimageiohandler.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QImageIOHandler()

/*
Constructs a QImageIOHandler object.
*/
func NewQImageIOHandler() *QImageIOHandler {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QImageIOHandlerC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageIOHandlerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImageIOHandler)
	return gothis
}

// /usr/include/qt/QtGui/qimageiohandler.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QImageIOHandler()

/*

 */
func DeleteQImageIOHandler(this *QImageIOHandler) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QImageIOHandlerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qimageiohandler.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDevice(QIODevice *)

/*
Sets the device of the QImageIOHandler to device. The image handler will use this device when reading and writing images.

The device can only be set once and must be set before calling canRead(), read(), write(), etc. If you need to read multiple files, construct multiple instances of the appropriate QImageIOHandler subclass.

See also device().
*/
func (this *QImageIOHandler) SetDevice(device qtcore.QIODevice_ITF /*777 QIODevice **/) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QImageIOHandler9setDeviceEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QIODevice * device() const

/*
Returns the device currently assigned to the QImageIOHandler. If not device has been assigned, 0 is returned.

See also setDevice().
*/
func (this *QImageIOHandler) Device() *qtcore.QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QImageIOHandler6deviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qimageiohandler.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFormat(const QByteArray &)

/*
Sets the format of the QImageIOHandler to format. The format is most useful for handlers that support multiple image formats.

See also format().
*/
func (this *QImageIOHandler) SetFormat(format qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QByteArray_PTR() != nil {
		convArg0 = format.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QImageIOHandler9setFormatERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:69
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setFormat(const QByteArray &) const

/*
Sets the format of the QImageIOHandler to format. The format is most useful for handlers that support multiple image formats.

See also format().
*/
func (this *QImageIOHandler) SetFormat_1(format qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QByteArray_PTR() != nil {
		convArg0 = format.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QImageIOHandler9setFormatERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray format() const

/*
Returns the format that is currently assigned to QImageIOHandler. If no format has been assigned, an empty string is returned.

See also setFormat().
*/
func (this *QImageIOHandler) Format() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QImageIOHandler6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtGui/qimageiohandler.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QByteArray name() const

/*

 */
func (this *QImageIOHandler) Name() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QImageIOHandler4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtGui/qimageiohandler.h:74
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool canRead() const

/*
Returns true if an image can be read from the device (i.e., the image format is supported, the device can be read from and the initial header information suggests that the image can be read); otherwise returns false.

When reimplementing canRead(), make sure that the I/O device (device()) is left in its original state (e.g., by using peek() rather than read()).

See also read() and QIODevice::peek().
*/
func (this *QImageIOHandler) CanRead() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QImageIOHandler7canReadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimageiohandler.h:75
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool read(QImage *)

/*
Read an image from the device, and stores it in image. Returns true if the image is successfully read; otherwise returns false.

For image formats that support incremental loading, and for animation formats, the image handler can assume that image points to the previous frame.

See also canRead().
*/
func (this *QImageIOHandler) Read(image QImage_ITF /*777 QImage **/) bool {
	var convArg0 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg0 = image.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QImageIOHandler4readEP6QImage", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimageiohandler.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool write(const QImage &)

/*
Writes the image image to the assigned device. Returns true on success; otherwise returns false.

The default implementation does nothing, and simply returns false.
*/
func (this *QImageIOHandler) Write(image QImage_ITF) bool {
	var convArg0 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg0 = image.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QImageIOHandler5writeERK6QImage", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimageiohandler.h:115
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant option(QImageIOHandler::ImageOption) const

/*
Returns the value assigned to option as a QVariant. The type of the value depends on the option. For example, option(Size) returns a QSize variant.

See also setOption() and supportsOption().
*/
func (this *QImageIOHandler) Option(option int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QImageIOHandler6optionENS_11ImageOptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtGui/qimageiohandler.h:116
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setOption(QImageIOHandler::ImageOption, const QVariant &)

/*
Sets the option option with the value value.

See also option() and ImageOption.
*/
func (this *QImageIOHandler) SetOption(option int, value qtcore.QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QImageIOHandler9setOptionENS_11ImageOptionERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:117
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool supportsOption(QImageIOHandler::ImageOption) const

/*
Returns true if the QImageIOHandler supports the option option; otherwise returns false. For example, if the QImageIOHandler supports the Size option, supportsOption(Size) must return true.

See also setOption() and option().
*/
func (this *QImageIOHandler) SupportsOption(option int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QImageIOHandler14supportsOptionENS_11ImageOptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimageiohandler.h:120
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool jumpToNextImage()

/*
For image formats that support animation, this function jumps to the next image.

The default implementation does nothing, and returns false.
*/
func (this *QImageIOHandler) JumpToNextImage() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QImageIOHandler15jumpToNextImageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimageiohandler.h:121
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool jumpToImage(int)

/*
For image formats that support animation, this function jumps to the image whose sequence number is imageNumber. The next call to read() will attempt to read this image.

The default implementation does nothing, and returns false.
*/
func (this *QImageIOHandler) JumpToImage(imageNumber int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QImageIOHandler11jumpToImageEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), imageNumber)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimageiohandler.h:122
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int loopCount() const

/*
For image formats that support animation, this function returns the number of times the animation should loop. If the image format does not support animation, 0 is returned.
*/
func (this *QImageIOHandler) LoopCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QImageIOHandler9loopCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimageiohandler.h:123
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int imageCount() const

/*
For image formats that support animation, this function returns the number of images in the animation. If the image format does not support animation, or if it is unable to determine the number of images, 0 is returned.

The default implementation returns 1 if canRead() returns true; otherwise 0 is returned.
*/
func (this *QImageIOHandler) ImageCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QImageIOHandler10imageCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimageiohandler.h:124
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int nextImageDelay() const

/*
For image formats that support animation, this function returns the number of milliseconds to wait until reading the next image. If the image format does not support animation, 0 is returned.
*/
func (this *QImageIOHandler) NextImageDelay() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QImageIOHandler14nextImageDelayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimageiohandler.h:125
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int currentImageNumber() const

/*
For image formats that support animation, this function returns the sequence number of the current image in the animation. If this function is called before any image is read(), -1 is returned. The number of the first image in the sequence is 0.

If the image format does not support animation, 0 is returned.

See also read().
*/
func (this *QImageIOHandler) CurrentImageNumber() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QImageIOHandler18currentImageNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimageiohandler.h:126
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QRect currentImageRect() const

/*
Returns the rect of the current image. If no rect is defined for the image, and empty QRect() is returned.

This function is useful for animations, where only parts of the frame may be updated at a time.
*/
func (this *QImageIOHandler) CurrentImageRect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QImageIOHandler16currentImageRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

/*
This enum describes the different options supported by QImageIOHandler. Some options are used to query an image for properties, and others are used to toggle the way in which an image should be written.


*/
type QImageIOHandler__ImageOption = int

// The original size of an image. A handler that supports this option is expected to read the size of the image from the image metadata, and return this size from option() as a QSize.
const QImageIOHandler__Size QImageIOHandler__ImageOption = 0

// The clip rect, or ROI (Region Of Interest). A handler that supports this option is expected to only read the provided QRect area from the original image in read(), before any other transformation is applied.
const QImageIOHandler__ClipRect QImageIOHandler__ImageOption = 1

// The image description. Some image formats, such as GIF and PNG, allow embedding of text or comments into the image data (e.g., for storing copyright information). It's common that the text is stored in key-value pairs, but some formats store all text in one continuous block. QImageIOHandler returns the text as one QString, where keys and values are separated by a ':', and keys-value pairs are separated by two newlines (\n\n). For example, "Title: Sunset\n\nAuthor: Jim Smith\nSarah Jones\n\n". Formats that store text in a single block can use "Description" as the key.
const QImageIOHandler__Description QImageIOHandler__ImageOption = 2

// The scaled clip rect (or ROI, Region Of Interest) of the image. A handler that supports this option is expected to apply the provided clip rect (a QRect), after applying any scaling (ScaleSize) or regular clipping (ClipRect). If the handler does not support this option, QImageReader will apply the scaled clip rect after the image has been read.
const QImageIOHandler__ScaledClipRect QImageIOHandler__ImageOption = 3

// The scaled size of the image. A handler that supports this option is expected to scale the image to the provided size (a QSize), after applying any clip rect transformation (ClipRect). If the handler does not support this option, QImageReader will perform the scaling after the image has been read.
const QImageIOHandler__ScaledSize QImageIOHandler__ImageOption = 4

// The compression ratio of the image data. A handler that supports this option is expected to set its compression rate depending on the value of this option (an int) when writing.
const QImageIOHandler__CompressionRatio QImageIOHandler__ImageOption = 5

// The gamma level of the image. A handler that supports this option is expected to set the image gamma level depending on the value of this option (a float) when writing.
const QImageIOHandler__Gamma QImageIOHandler__ImageOption = 6

// The quality level of the image. A handler that supports this option is expected to set the image quality level depending on the value of this option (an int) when writing.
const QImageIOHandler__Quality QImageIOHandler__ImageOption = 7

// The name of the image. A handler that supports this option is expected to read the name from the image metadata and return this as a QString, or when writing an image it is expected to store the name in the image metadata.
const QImageIOHandler__Name QImageIOHandler__ImageOption = 8

// The subtype of the image. A handler that supports this option can use the subtype value to help when reading and writing images. For example, a PPM handler may have a subtype value of "ppm" or "ppmraw".
const QImageIOHandler__SubType QImageIOHandler__ImageOption = 9

//
const QImageIOHandler__IncrementalReading QImageIOHandler__ImageOption = 10

//
const QImageIOHandler__Endianness QImageIOHandler__ImageOption = 11

//
const QImageIOHandler__Animation QImageIOHandler__ImageOption = 12

//
const QImageIOHandler__BackgroundColor QImageIOHandler__ImageOption = 13

//
const QImageIOHandler__ImageFormat QImageIOHandler__ImageOption = 14

//
const QImageIOHandler__SupportedSubTypes QImageIOHandler__ImageOption = 15

//
const QImageIOHandler__OptimizedWrite QImageIOHandler__ImageOption = 16

//
const QImageIOHandler__ProgressiveScanWrite QImageIOHandler__ImageOption = 17

//
const QImageIOHandler__ImageTransformation QImageIOHandler__ImageOption = 18

//
const QImageIOHandler__TransformedByDefault QImageIOHandler__ImageOption = 19

func (this *QImageIOHandler) ImageOptionItemName(val int) string {
	switch val {
	case QImageIOHandler__Size: // 0
		return "Size"
	case QImageIOHandler__ClipRect: // 1
		return "ClipRect"
	case QImageIOHandler__Description: // 2
		return "Description"
	case QImageIOHandler__ScaledClipRect: // 3
		return "ScaledClipRect"
	case QImageIOHandler__ScaledSize: // 4
		return "ScaledSize"
	case QImageIOHandler__CompressionRatio: // 5
		return "CompressionRatio"
	case QImageIOHandler__Gamma: // 6
		return "Gamma"
	case QImageIOHandler__Quality: // 7
		return "Quality"
	case QImageIOHandler__Name: // 8
		return "Name"
	case QImageIOHandler__SubType: // 9
		return "SubType"
	case QImageIOHandler__IncrementalReading: // 10
		return "IncrementalReading"
	case QImageIOHandler__Endianness: // 11
		return "Endianness"
	case QImageIOHandler__Animation: // 12
		return "Animation"
	case QImageIOHandler__BackgroundColor: // 13
		return "BackgroundColor"
	case QImageIOHandler__ImageFormat: // 14
		return "ImageFormat"
	case QImageIOHandler__SupportedSubTypes: // 15
		return "SupportedSubTypes"
	case QImageIOHandler__OptimizedWrite: // 16
		return "OptimizedWrite"
	case QImageIOHandler__ProgressiveScanWrite: // 17
		return "ProgressiveScanWrite"
	case QImageIOHandler__ImageTransformation: // 18
		return "ImageTransformation"
	case QImageIOHandler__TransformedByDefault: // 19
		return "TransformedByDefault"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QImageIOHandler_ImageOptionItemName(val int) string {
	var nilthis *QImageIOHandler
	return nilthis.ImageOptionItemName(val)
}

/*


 */
type QImageIOHandler__Transformation = int

//
const QImageIOHandler__TransformationNone QImageIOHandler__Transformation = 0

//
const QImageIOHandler__TransformationMirror QImageIOHandler__Transformation = 1

//
const QImageIOHandler__TransformationFlip QImageIOHandler__Transformation = 2

//
const QImageIOHandler__TransformationRotate180 QImageIOHandler__Transformation = 3

//
const QImageIOHandler__TransformationRotate90 QImageIOHandler__Transformation = 4

//
const QImageIOHandler__TransformationMirrorAndRotate90 QImageIOHandler__Transformation = 5

//
const QImageIOHandler__TransformationFlipAndRotate90 QImageIOHandler__Transformation = 6

//
const QImageIOHandler__TransformationRotate270 QImageIOHandler__Transformation = 7

func (this *QImageIOHandler) TransformationItemName(val int) string {
	switch val {
	case QImageIOHandler__TransformationNone: // 0
		return "TransformationNone"
	case QImageIOHandler__TransformationMirror: // 1
		return "TransformationMirror"
	case QImageIOHandler__TransformationFlip: // 2
		return "TransformationFlip"
	case QImageIOHandler__TransformationRotate180: // 3
		return "TransformationRotate180"
	case QImageIOHandler__TransformationRotate90: // 4
		return "TransformationRotate90"
	case QImageIOHandler__TransformationMirrorAndRotate90: // 5
		return "TransformationMirrorAndRotate90"
	case QImageIOHandler__TransformationFlipAndRotate90: // 6
		return "TransformationFlipAndRotate90"
	case QImageIOHandler__TransformationRotate270: // 7
		return "TransformationRotate270"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QImageIOHandler_TransformationItemName(val int) string {
	var nilthis *QImageIOHandler
	return nilthis.TransformationItemName(val)
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
