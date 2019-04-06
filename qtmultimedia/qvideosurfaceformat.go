package qtmultimedia

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h
// #include <qvideosurfaceformat.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 39
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

/*

 */
type QVideoSurfaceFormat struct {
	*qtrt.CObject
}
type QVideoSurfaceFormat_ITF interface {
	QVideoSurfaceFormat_PTR() *QVideoSurfaceFormat
}

func (ptr *QVideoSurfaceFormat) QVideoSurfaceFormat_PTR() *QVideoSurfaceFormat { return ptr }

func (this *QVideoSurfaceFormat) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QVideoSurfaceFormat) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQVideoSurfaceFormatFromPointer(cthis unsafe.Pointer) *QVideoSurfaceFormat {
	return &QVideoSurfaceFormat{&qtrt.CObject{cthis}}
}
func (*QVideoSurfaceFormat) NewFromPointer(cthis unsafe.Pointer) *QVideoSurfaceFormat {
	return NewQVideoSurfaceFormatFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QVideoSurfaceFormat()

/*
Constructs a null video stream format.
*/
func (*QVideoSurfaceFormat) NewForInherit() *QVideoSurfaceFormat {
	return NewQVideoSurfaceFormat()
}
func NewQVideoSurfaceFormat() *QVideoSurfaceFormat {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoSurfaceFormatC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVideoSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVideoSurfaceFormat)
	return gothis
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:80
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QVideoSurfaceFormat(const QSize &, QVideoFrame::PixelFormat, QAbstractVideoBuffer::HandleType)

/*
Constructs a null video stream format.
*/
func (*QVideoSurfaceFormat) NewForInherit1(size qtcore.QSize_ITF, pixelFormat int, handleType int) *QVideoSurfaceFormat {
	return NewQVideoSurfaceFormat1(size, pixelFormat, handleType)
}
func NewQVideoSurfaceFormat1(size qtcore.QSize_ITF, pixelFormat int, handleType int) *QVideoSurfaceFormat {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoSurfaceFormatC2ERK5QSizeN11QVideoFrame11PixelFormatEN20QAbstractVideoBuffer10HandleTypeE", qtrt.FFI_TYPE_POINTER, convArg0, pixelFormat, handleType)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVideoSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVideoSurfaceFormat)
	return gothis
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:80
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QVideoSurfaceFormat(const QSize &, QVideoFrame::PixelFormat, QAbstractVideoBuffer::HandleType)

/*
Constructs a null video stream format.
*/
func (*QVideoSurfaceFormat) NewForInherit1p(size qtcore.QSize_ITF, pixelFormat int) *QVideoSurfaceFormat {
	return NewQVideoSurfaceFormat1p(size, pixelFormat)
}
func NewQVideoSurfaceFormat1p(size qtcore.QSize_ITF, pixelFormat int) *QVideoSurfaceFormat {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	// arg: 2, QAbstractVideoBuffer::HandleType=Elaborated, QAbstractVideoBuffer::HandleType=Enum, , Invalid
	handleType := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoSurfaceFormatC2ERK5QSizeN11QVideoFrame11PixelFormatEN20QAbstractVideoBuffer10HandleTypeE", qtrt.FFI_TYPE_POINTER, convArg0, pixelFormat, handleType)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVideoSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVideoSurfaceFormat)
	return gothis
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QVideoSurfaceFormat()

/*

 */
func DeleteQVideoSurfaceFormat(this *QVideoSurfaceFormat) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoSurfaceFormatD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:87
// index:0
// Public Visibility=Default Availability=Available
// [8] QVideoSurfaceFormat & operator=(const QVideoSurfaceFormat &)

/*

 */
func (this *QVideoSurfaceFormat) Operator_equal(format QVideoSurfaceFormat_ITF) *QVideoSurfaceFormat {
	var convArg0 unsafe.Pointer
	if format != nil && format.QVideoSurfaceFormat_PTR() != nil {
		convArg0 = format.QVideoSurfaceFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoSurfaceFormataSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVideoSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVideoSurfaceFormat)
	return rv2
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:89
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QVideoSurfaceFormat &) const

/*

 */
func (this *QVideoSurfaceFormat) Operator_equal_equal(format QVideoSurfaceFormat_ITF) bool {
	var convArg0 unsafe.Pointer
	if format != nil && format.QVideoSurfaceFormat_PTR() != nil {
		convArg0 = format.QVideoSurfaceFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QVideoSurfaceFormateqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:90
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QVideoSurfaceFormat &) const

/*

 */
func (this *QVideoSurfaceFormat) Operator_not_equal(format QVideoSurfaceFormat_ITF) bool {
	var convArg0 unsafe.Pointer
	if format != nil && format.QVideoSurfaceFormat_PTR() != nil {
		convArg0 = format.QVideoSurfaceFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QVideoSurfaceFormatneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:92
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Identifies if a video surface format has a valid pixel format and frame size.

Returns true if the format is valid, and false otherwise.
*/
func (this *QVideoSurfaceFormat) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QVideoSurfaceFormat7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:94
// index:0
// Public Visibility=Default Availability=Available
// [4] QVideoFrame::PixelFormat pixelFormat() const

/*
Returns the pixel format of frames in a video stream.
*/
func (this *QVideoSurfaceFormat) PixelFormat() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QVideoSurfaceFormat11pixelFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:95
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractVideoBuffer::HandleType handleType() const

/*
Returns the type of handle the surface uses to present the frame data.

If the handle type is QAbstractVideoBuffer::NoHandle, buffers with any handle type are valid provided they can be mapped with the QAbstractVideoBuffer::ReadOnly flag. If the handleType() is not QAbstractVideoBuffer::NoHandle then the handle type of the buffer must be the same as that of the surface format.
*/
func (this *QVideoSurfaceFormat) HandleType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QVideoSurfaceFormat10handleTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize frameSize() const

/*
Returns the dimensions of frames in a video stream.

See also setFrameSize(), frameWidth(), and frameHeight().
*/
func (this *QVideoSurfaceFormat) FrameSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QVideoSurfaceFormat9frameSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFrameSize(const QSize &)

/*
Sets the size of frames in a video stream to size.

This will reset the viewport() to fill the entire frame.

See also frameSize().
*/
func (this *QVideoSurfaceFormat) SetFrameSize(size qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoSurfaceFormat12setFrameSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:99
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setFrameSize(int, int)

/*
Sets the size of frames in a video stream to size.

This will reset the viewport() to fill the entire frame.

See also frameSize().
*/
func (this *QVideoSurfaceFormat) SetFrameSize1(width int, height int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoSurfaceFormat12setFrameSizeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width, height)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] int frameWidth() const

/*
Returns the width of frames in a video stream.

See also frameSize() and frameHeight().
*/
func (this *QVideoSurfaceFormat) FrameWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QVideoSurfaceFormat10frameWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:102
// index:0
// Public Visibility=Default Availability=Available
// [4] int frameHeight() const

/*
Returns the height of frame in a video stream.
*/
func (this *QVideoSurfaceFormat) FrameHeight() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QVideoSurfaceFormat11frameHeightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:104
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect viewport() const

/*
Returns the viewport of a video stream.

The viewport is the region of a video frame that is actually displayed.

By default the viewport covers an entire frame.

See also setViewport().
*/
func (this *QVideoSurfaceFormat) Viewport() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QVideoSurfaceFormat8viewportEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setViewport(const QRect &)

/*
Sets the viewport of a video stream to viewport.

See also viewport().
*/
func (this *QVideoSurfaceFormat) SetViewport(viewport qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if viewport != nil && viewport.QRect_PTR() != nil {
		convArg0 = viewport.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoSurfaceFormat11setViewportERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:107
// index:0
// Public Visibility=Default Availability=Available
// [4] QVideoSurfaceFormat::Direction scanLineDirection() const

/*
Returns the direction of scan lines.

See also setScanLineDirection().
*/
func (this *QVideoSurfaceFormat) ScanLineDirection() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QVideoSurfaceFormat17scanLineDirectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScanLineDirection(QVideoSurfaceFormat::Direction)

/*
Sets the direction of scan lines.

See also scanLineDirection().
*/
func (this *QVideoSurfaceFormat) SetScanLineDirection(direction int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoSurfaceFormat20setScanLineDirectionENS_9DirectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), direction)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:110
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal frameRate() const

/*
Returns the frame rate of a video stream in frames per second.

See also setFrameRate().
*/
func (this *QVideoSurfaceFormat) FrameRate() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QVideoSurfaceFormat9frameRateEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFrameRate(qreal)

/*
Sets the frame rate of a video stream in frames per second.

See also frameRate().
*/
func (this *QVideoSurfaceFormat) SetFrameRate(rate float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoSurfaceFormat12setFrameRateEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rate)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:113
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize pixelAspectRatio() const

/*
Returns a video stream's pixel aspect ratio.

See also setPixelAspectRatio().
*/
func (this *QVideoSurfaceFormat) PixelAspectRatio() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QVideoSurfaceFormat16pixelAspectRatioEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPixelAspectRatio(const QSize &)

/*
Sets a video stream's pixel aspect ratio.

See also pixelAspectRatio().
*/
func (this *QVideoSurfaceFormat) SetPixelAspectRatio(ratio qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if ratio != nil && ratio.QSize_PTR() != nil {
		convArg0 = ratio.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoSurfaceFormat19setPixelAspectRatioERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:115
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setPixelAspectRatio(int, int)

/*
Sets a video stream's pixel aspect ratio.

See also pixelAspectRatio().
*/
func (this *QVideoSurfaceFormat) SetPixelAspectRatio1(width int, height int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoSurfaceFormat19setPixelAspectRatioEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width, height)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:117
// index:0
// Public Visibility=Default Availability=Available
// [4] QVideoSurfaceFormat::YCbCrColorSpace yCbCrColorSpace() const

/*
Returns the Y'CbCr color space of a video stream.

See also setYCbCrColorSpace().
*/
func (this *QVideoSurfaceFormat) YCbCrColorSpace() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QVideoSurfaceFormat15yCbCrColorSpaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setYCbCrColorSpace(QVideoSurfaceFormat::YCbCrColorSpace)

/*
Sets the Y'CbCr color space of a video stream. It is only used with raw YUV frame types.

See also yCbCrColorSpace().
*/
func (this *QVideoSurfaceFormat) SetYCbCrColorSpace(colorSpace int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoSurfaceFormat18setYCbCrColorSpaceENS_15YCbCrColorSpaceE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), colorSpace)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:120
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMirrored() const

/*
Returns true if the surface is mirrored around its vertical axis. This is typically needed for video frames coming from a front camera of a mobile device.

Note: The mirroring here differs from QImage::mirrored, as a vertically mirrored QImage will be mirrored around its x-axis.

This function was introduced in  Qt 5.11.
*/
func (this *QVideoSurfaceFormat) IsMirrored() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QVideoSurfaceFormat10isMirroredEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMirrored(bool)

/*
Sets if the surface is mirrored around its vertical axis. This is typically needed for video frames coming from a front camera of a mobile device. Default value is false.

Note: The mirroring here differs from QImage::mirrored, as a vertically mirrored QImage will be mirrored around its x-axis.

This function was introduced in  Qt 5.11.

See also isMirrored().
*/
func (this *QVideoSurfaceFormat) SetMirrored(mirrored bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoSurfaceFormat11setMirroredEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mirrored)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:123
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Returns a suggested size in pixels for the video stream.

This is the size of the viewport scaled according to the pixel aspect ratio.
*/
func (this *QVideoSurfaceFormat) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QVideoSurfaceFormat8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:126
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant property(const char *) const

/*
Returns the value of the video format's name property.

See also setProperty().
*/
func (this *QVideoSurfaceFormat) Property(name string) *qtcore.QVariant /*123*/ {
	var convArg0 = qtrt.CString(name)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QVideoSurfaceFormat8propertyEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtMultimedia/qvideosurfaceformat.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProperty(const char *, const QVariant &)

/*
Sets the video format's name property to value.

Trying to set a read only property will be ignored.

See also property().
*/
func (this *QVideoSurfaceFormat) SetProperty(name string, value qtcore.QVariant_ITF) {
	var convArg0 = qtrt.CString(name)
	defer qtrt.FreeMem(convArg0)
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoSurfaceFormat11setPropertyEPKcRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

/*
Enumerates the layout direction of video scan lines.


*/
type QVideoSurfaceFormat__Direction = int

// Scan lines are arranged from the top of the frame to the bottom.
const QVideoSurfaceFormat__TopToBottom QVideoSurfaceFormat__Direction = 0

// Scan lines are arranged from the bottom of the frame to the top.
const QVideoSurfaceFormat__BottomToTop QVideoSurfaceFormat__Direction = 1

func (this *QVideoSurfaceFormat) DirectionItemName(val int) string {
	switch val {
	case QVideoSurfaceFormat__TopToBottom: // 0
		return "TopToBottom"
	case QVideoSurfaceFormat__BottomToTop: // 1
		return "BottomToTop"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QVideoSurfaceFormat_DirectionItemName(val int) string {
	var nilthis *QVideoSurfaceFormat
	return nilthis.DirectionItemName(val)
}

/*
Enumerates the Y'CbCr color space of video frames.


*/
type QVideoSurfaceFormat__YCbCrColorSpace = int

// No color space is specified.
const QVideoSurfaceFormat__YCbCr_Undefined QVideoSurfaceFormat__YCbCrColorSpace = 0

//
const QVideoSurfaceFormat__YCbCr_BT601 QVideoSurfaceFormat__YCbCrColorSpace = 1

//
const QVideoSurfaceFormat__YCbCr_BT709 QVideoSurfaceFormat__YCbCrColorSpace = 2

//
const QVideoSurfaceFormat__YCbCr_xvYCC601 QVideoSurfaceFormat__YCbCrColorSpace = 3

//
const QVideoSurfaceFormat__YCbCr_xvYCC709 QVideoSurfaceFormat__YCbCrColorSpace = 4

// The full range Y'CbCr color space used in JPEG files.
const QVideoSurfaceFormat__YCbCr_JPEG QVideoSurfaceFormat__YCbCrColorSpace = 5

//
const QVideoSurfaceFormat__YCbCr_CustomMatrix QVideoSurfaceFormat__YCbCrColorSpace = 6

func (this *QVideoSurfaceFormat) YCbCrColorSpaceItemName(val int) string {
	switch val {
	case QVideoSurfaceFormat__YCbCr_Undefined: // 0
		return "YCbCr_Undefined"
	case QVideoSurfaceFormat__YCbCr_BT601: // 1
		return "YCbCr_BT601"
	case QVideoSurfaceFormat__YCbCr_BT709: // 2
		return "YCbCr_BT709"
	case QVideoSurfaceFormat__YCbCr_xvYCC601: // 3
		return "YCbCr_xvYCC601"
	case QVideoSurfaceFormat__YCbCr_xvYCC709: // 4
		return "YCbCr_xvYCC709"
	case QVideoSurfaceFormat__YCbCr_JPEG: // 5
		return "YCbCr_JPEG"
	case QVideoSurfaceFormat__YCbCr_CustomMatrix: // 6
		return "YCbCr_CustomMatrix"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QVideoSurfaceFormat_YCbCrColorSpaceItemName(val int) string {
	var nilthis *QVideoSurfaceFormat
	return nilthis.YCbCrColorSpaceItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11739() {
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
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
