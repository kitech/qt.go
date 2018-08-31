package qtgui

// /usr/include/qt/QtGui/qpixelformat.h
// #include <qpixelformat.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
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
type QPixelFormat struct {
	*qtrt.CObject
}
type QPixelFormat_ITF interface {
	QPixelFormat_PTR() *QPixelFormat
}

func (ptr *QPixelFormat) QPixelFormat_PTR() *QPixelFormat { return ptr }

func (this *QPixelFormat) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPixelFormat) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQPixelFormatFromPointer(cthis unsafe.Pointer) *QPixelFormat {
	return &QPixelFormat{&qtrt.CObject{cthis}}
}
func (*QPixelFormat) NewFromPointer(cthis unsafe.Pointer) *QPixelFormat {
	return NewQPixelFormatFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpixelformat.h:163
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QPixelFormat()

/*
Creates a null pixelformat. This format maps to QImage::Format_Invalid.
*/
func (*QPixelFormat) NewForInherit() *QPixelFormat {
	return NewQPixelFormat()
}
func NewQPixelFormat() *QPixelFormat {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPixelFormatC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPixelFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPixelFormat)
	return gothis
}

// /usr/include/qt/QtGui/qpixelformat.h:164
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QPixelFormat(QPixelFormat::ColorModel, uchar, uchar, uchar, uchar, uchar, uchar, QPixelFormat::AlphaUsage, QPixelFormat::AlphaPosition, QPixelFormat::AlphaPremultiplied, QPixelFormat::TypeInterpretation, QPixelFormat::ByteOrder, uchar)

/*
Creates a null pixelformat. This format maps to QImage::Format_Invalid.
*/
func (*QPixelFormat) NewForInherit_1(colorModel int, firstSize byte, secondSize byte, thirdSize byte, fourthSize byte, fifthSize byte, alphaSize byte, alphaUsage int, alphaPosition int, premultiplied int, typeInterpretation int, byteOrder int, subEnum byte) *QPixelFormat {
	return NewQPixelFormat_1(colorModel, firstSize, secondSize, thirdSize, fourthSize, fifthSize, alphaSize, alphaUsage, alphaPosition, premultiplied, typeInterpretation, byteOrder, subEnum)
}
func NewQPixelFormat_1(colorModel int, firstSize byte, secondSize byte, thirdSize byte, fourthSize byte, fifthSize byte, alphaSize byte, alphaUsage int, alphaPosition int, premultiplied int, typeInterpretation int, byteOrder int, subEnum byte) *QPixelFormat {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPixelFormatC2ENS_10ColorModelEhhhhhhNS_10AlphaUsageENS_13AlphaPositionENS_18AlphaPremultipliedENS_18TypeInterpretationENS_9ByteOrderEh", qtrt.FFI_TYPE_POINTER, colorModel, firstSize, secondSize, thirdSize, fourthSize, fifthSize, alphaSize, alphaUsage, alphaPosition, premultiplied, typeInterpretation, byteOrder, subEnum)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPixelFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPixelFormat)
	return gothis
}

// /usr/include/qt/QtGui/qpixelformat.h:164
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QPixelFormat(QPixelFormat::ColorModel, uchar, uchar, uchar, uchar, uchar, uchar, QPixelFormat::AlphaUsage, QPixelFormat::AlphaPosition, QPixelFormat::AlphaPremultiplied, QPixelFormat::TypeInterpretation, QPixelFormat::ByteOrder, uchar)

/*
Creates a null pixelformat. This format maps to QImage::Format_Invalid.
*/
func (*QPixelFormat) NewForInherit_1_(colorModel int, firstSize byte, secondSize byte, thirdSize byte, fourthSize byte, fifthSize byte, alphaSize byte, alphaUsage int, alphaPosition int, premultiplied int, typeInterpretation int) *QPixelFormat {
	return NewQPixelFormat_1_(colorModel, firstSize, secondSize, thirdSize, fourthSize, fifthSize, alphaSize, alphaUsage, alphaPosition, premultiplied, typeInterpretation)
}
func NewQPixelFormat_1_(colorModel int, firstSize byte, secondSize byte, thirdSize byte, fourthSize byte, fifthSize byte, alphaSize byte, alphaUsage int, alphaPosition int, premultiplied int, typeInterpretation int) *QPixelFormat {
	// arg: 11, QPixelFormat::ByteOrder=Enum, QPixelFormat::ByteOrder=Enum, , Invalid
	byteOrder := 0
	// arg: 12, uchar=Typedef, uchar=Typedef, unsigned char, UChar
	subEnum := byte(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPixelFormatC2ENS_10ColorModelEhhhhhhNS_10AlphaUsageENS_13AlphaPositionENS_18AlphaPremultipliedENS_18TypeInterpretationENS_9ByteOrderEh", qtrt.FFI_TYPE_POINTER, colorModel, firstSize, secondSize, thirdSize, fourthSize, fifthSize, alphaSize, alphaUsage, alphaPosition, premultiplied, typeInterpretation, byteOrder, subEnum)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPixelFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPixelFormat)
	return gothis
}

// /usr/include/qt/QtGui/qpixelformat.h:164
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QPixelFormat(QPixelFormat::ColorModel, uchar, uchar, uchar, uchar, uchar, uchar, QPixelFormat::AlphaUsage, QPixelFormat::AlphaPosition, QPixelFormat::AlphaPremultiplied, QPixelFormat::TypeInterpretation, QPixelFormat::ByteOrder, uchar)

/*
Creates a null pixelformat. This format maps to QImage::Format_Invalid.
*/
func (*QPixelFormat) NewForInherit_1_1(colorModel int, firstSize byte, secondSize byte, thirdSize byte, fourthSize byte, fifthSize byte, alphaSize byte, alphaUsage int, alphaPosition int, premultiplied int, typeInterpretation int, byteOrder int) *QPixelFormat {
	return NewQPixelFormat_1_1(colorModel, firstSize, secondSize, thirdSize, fourthSize, fifthSize, alphaSize, alphaUsage, alphaPosition, premultiplied, typeInterpretation, byteOrder)
}
func NewQPixelFormat_1_1(colorModel int, firstSize byte, secondSize byte, thirdSize byte, fourthSize byte, fifthSize byte, alphaSize byte, alphaUsage int, alphaPosition int, premultiplied int, typeInterpretation int, byteOrder int) *QPixelFormat {
	// arg: 12, uchar=Typedef, uchar=Typedef, unsigned char, UChar
	subEnum := byte(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPixelFormatC2ENS_10ColorModelEhhhhhhNS_10AlphaUsageENS_13AlphaPositionENS_18AlphaPremultipliedENS_18TypeInterpretationENS_9ByteOrderEh", qtrt.FFI_TYPE_POINTER, colorModel, firstSize, secondSize, thirdSize, fourthSize, fifthSize, alphaSize, alphaUsage, alphaPosition, premultiplied, typeInterpretation, byteOrder, subEnum)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPixelFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPixelFormat)
	return gothis
}

// /usr/include/qt/QtGui/qpixelformat.h:178
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QPixelFormat::ColorModel colorModel() const

/*
Accessor function for getting the colorModel.
*/
func (this *QPixelFormat) ColorModel() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPixelFormat10colorModelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpixelformat.h:179
// index:0
// Public inline Visibility=Default Availability=Available
// [1] uchar channelCount() const

/*
Accessor function for getting the channelCount. Channel Count is deduced by color channels with a size > 0 and if the size of the alpha channel is > 0.
*/
func (this *QPixelFormat) ChannelCount() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPixelFormat12channelCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:186
// index:0
// Public inline Visibility=Default Availability=Available
// [1] uchar redSize() const

/*
Accessor function for the size of the red color channel.
*/
func (this *QPixelFormat) RedSize() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPixelFormat7redSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:187
// index:0
// Public inline Visibility=Default Availability=Available
// [1] uchar greenSize() const

/*
Accessor function for the size of the green color channel.
*/
func (this *QPixelFormat) GreenSize() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPixelFormat9greenSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:188
// index:0
// Public inline Visibility=Default Availability=Available
// [1] uchar blueSize() const

/*
Accessor function for the size of the blue color channel.
*/
func (this *QPixelFormat) BlueSize() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPixelFormat8blueSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:190
// index:0
// Public inline Visibility=Default Availability=Available
// [1] uchar cyanSize() const

/*
Accessor function for the cyan color channel.
*/
func (this *QPixelFormat) CyanSize() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPixelFormat8cyanSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:191
// index:0
// Public inline Visibility=Default Availability=Available
// [1] uchar magentaSize() const

/*
Accessor function for the megenta color channel.
*/
func (this *QPixelFormat) MagentaSize() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPixelFormat11magentaSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:192
// index:0
// Public inline Visibility=Default Availability=Available
// [1] uchar yellowSize() const

/*
Accessor function for the yellow color channel.
*/
func (this *QPixelFormat) YellowSize() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPixelFormat10yellowSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:193
// index:0
// Public inline Visibility=Default Availability=Available
// [1] uchar blackSize() const

/*
Accessor function for the black/key color channel.
*/
func (this *QPixelFormat) BlackSize() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPixelFormat9blackSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:195
// index:0
// Public inline Visibility=Default Availability=Available
// [1] uchar hueSize() const

/*
Accessor function for the hue channel size.
*/
func (this *QPixelFormat) HueSize() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPixelFormat7hueSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:196
// index:0
// Public inline Visibility=Default Availability=Available
// [1] uchar saturationSize() const

/*
Accessor function for the saturation channel size.
*/
func (this *QPixelFormat) SaturationSize() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPixelFormat14saturationSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:197
// index:0
// Public inline Visibility=Default Availability=Available
// [1] uchar lightnessSize() const

/*
Accessor function for the lightness channel size.
*/
func (this *QPixelFormat) LightnessSize() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPixelFormat13lightnessSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:198
// index:0
// Public inline Visibility=Default Availability=Available
// [1] uchar brightnessSize() const

/*
Accessor function for the brightness channel size.
*/
func (this *QPixelFormat) BrightnessSize() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPixelFormat14brightnessSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:200
// index:0
// Public inline Visibility=Default Availability=Available
// [1] uchar alphaSize() const

/*
Accessor function for the alpha channel size.
*/
func (this *QPixelFormat) AlphaSize() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPixelFormat9alphaSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:202
// index:0
// Public inline Visibility=Default Availability=Available
// [1] uchar bitsPerPixel() const

/*
Accessor function for the bits used per pixel. This function returns the sum of the color channels + the size of the alpha channel.
*/
func (this *QPixelFormat) BitsPerPixel() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPixelFormat12bitsPerPixelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:209
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QPixelFormat::AlphaUsage alphaUsage() const

/*
Accessor function for alphaUsage.
*/
func (this *QPixelFormat) AlphaUsage() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPixelFormat10alphaUsageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpixelformat.h:210
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QPixelFormat::AlphaPosition alphaPosition() const

/*
Accessor function for alphaPosition.
*/
func (this *QPixelFormat) AlphaPosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPixelFormat13alphaPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpixelformat.h:211
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QPixelFormat::AlphaPremultiplied premultiplied() const

/*
Accessor function for the AlphaPremultiplied enum. This indicates if the alpha channel is multiplied in to the color channels.
*/
func (this *QPixelFormat) Premultiplied() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPixelFormat13premultipliedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpixelformat.h:212
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QPixelFormat::TypeInterpretation typeInterpretation() const

/*
Accessor function for the type representation of a color channel or a pixel.

See also TypeInterpretation.
*/
func (this *QPixelFormat) TypeInterpretation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPixelFormat18typeInterpretationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpixelformat.h:213
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QPixelFormat::ByteOrder byteOrder() const

/*
The byte order is almost always set the the byte order of the current system. However, it can be useful to describe some YUV formats. This function should never return QPixelFormat::CurrentSystemEndian as this value is translated to a endian value in the constructor.
*/
func (this *QPixelFormat) ByteOrder() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPixelFormat9byteOrderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpixelformat.h:215
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QPixelFormat::YUVLayout yuvLayout() const

/*
Accessor function for the YUVLayout. It is difficult to describe the color channels of a YUV pixel format since YUV color model uses macro pixels. Instead the layout of the pixels are stored as an enum.
*/
func (this *QPixelFormat) YuvLayout() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPixelFormat9yuvLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpixelformat.h:216
// index:0
// Public inline Visibility=Default Availability=Available
// [1] uchar subEnum() const

/*

 */
func (this *QPixelFormat) SubEnum() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPixelFormat7subEnumEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return byte(rv) // 222
}

func DeleteQPixelFormat(this *QPixelFormat) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPixelFormatD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QPixelFormat__FieldWidth = int

//
const QPixelFormat__ModelFieldWidth QPixelFormat__FieldWidth = 4

//
const QPixelFormat__FirstFieldWidth QPixelFormat__FieldWidth = 6

//
const QPixelFormat__SecondFieldWidth QPixelFormat__FieldWidth = 6

//
const QPixelFormat__ThirdFieldWidth QPixelFormat__FieldWidth = 6

//
const QPixelFormat__FourthFieldWidth QPixelFormat__FieldWidth = 6

//
const QPixelFormat__FifthFieldWidth QPixelFormat__FieldWidth = 6

//
const QPixelFormat__AlphaFieldWidth QPixelFormat__FieldWidth = 6

//
const QPixelFormat__AlphaUsageFieldWidth QPixelFormat__FieldWidth = 1

//
const QPixelFormat__AlphaPositionFieldWidth QPixelFormat__FieldWidth = 1

//
const QPixelFormat__PremulFieldWidth QPixelFormat__FieldWidth = 1

//
const QPixelFormat__TypeInterpretationFieldWidth QPixelFormat__FieldWidth = 4

//
const QPixelFormat__ByteOrderFieldWidth QPixelFormat__FieldWidth = 2

//
const QPixelFormat__SubEnumFieldWidth QPixelFormat__FieldWidth = 6

//
const QPixelFormat__UnusedFieldWidth QPixelFormat__FieldWidth = 9

//
const QPixelFormat__TotalFieldWidthByWidths QPixelFormat__FieldWidth = 64

func (this *QPixelFormat) FieldWidthItemName(val int) string {
	switch val {
	case QPixelFormat__ModelFieldWidth: // 4
		return "ModelFieldWidth,TypeInterpretationFieldWidth"
	case QPixelFormat__FirstFieldWidth: // 6
		return "FirstFieldWidth,SecondFieldWidth,ThirdFieldWidth,FourthFieldWidth,FifthFieldWidth,AlphaFieldWidth,SubEnumFieldWidth"
		// case QPixelFormat__SecondFieldWidth: // 6
		// return ""
		// case QPixelFormat__ThirdFieldWidth: // 6
		// return ""
		// case QPixelFormat__FourthFieldWidth: // 6
		// return ""
		// case QPixelFormat__FifthFieldWidth: // 6
		// return ""
		// case QPixelFormat__AlphaFieldWidth: // 6
		// return ""
	case QPixelFormat__AlphaUsageFieldWidth: // 1
		return "AlphaUsageFieldWidth,AlphaPositionFieldWidth,PremulFieldWidth"
		// case QPixelFormat__AlphaPositionFieldWidth: // 1
		// return ""
		// case QPixelFormat__PremulFieldWidth: // 1
		// return ""
		// case QPixelFormat__TypeInterpretationFieldWidth: // 4
		// return ""
	case QPixelFormat__ByteOrderFieldWidth: // 2
		return "ByteOrderFieldWidth"
		// case QPixelFormat__SubEnumFieldWidth: // 6
		// return ""
	case QPixelFormat__UnusedFieldWidth: // 9
		return "UnusedFieldWidth"
	case QPixelFormat__TotalFieldWidthByWidths: // 64
		return "TotalFieldWidthByWidths"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QPixelFormat_FieldWidthItemName(val int) string {
	var nilthis *QPixelFormat
	return nilthis.FieldWidthItemName(val)
}

/*


 */
type QPixelFormat__Field = int

//
const QPixelFormat__ModelField QPixelFormat__Field = 0

//
const QPixelFormat__FirstField QPixelFormat__Field = 4

//
const QPixelFormat__SecondField QPixelFormat__Field = 10

//
const QPixelFormat__ThirdField QPixelFormat__Field = 16

//
const QPixelFormat__FourthField QPixelFormat__Field = 22

//
const QPixelFormat__FifthField QPixelFormat__Field = 28

//
const QPixelFormat__AlphaField QPixelFormat__Field = 34

//
const QPixelFormat__AlphaUsageField QPixelFormat__Field = 40

//
const QPixelFormat__AlphaPositionField QPixelFormat__Field = 41

//
const QPixelFormat__PremulField QPixelFormat__Field = 42

//
const QPixelFormat__TypeInterpretationField QPixelFormat__Field = 43

//
const QPixelFormat__ByteOrderField QPixelFormat__Field = 47

//
const QPixelFormat__SubEnumField QPixelFormat__Field = 49

//
const QPixelFormat__UnusedField QPixelFormat__Field = 55

//
const QPixelFormat__TotalFieldWidthByOffsets QPixelFormat__Field = 64

func (this *QPixelFormat) FieldItemName(val int) string {
	switch val {
	case QPixelFormat__ModelField: // 0
		return "ModelField"
	case QPixelFormat__FirstField: // 4
		return "FirstField"
	case QPixelFormat__SecondField: // 10
		return "SecondField"
	case QPixelFormat__ThirdField: // 16
		return "ThirdField"
	case QPixelFormat__FourthField: // 22
		return "FourthField"
	case QPixelFormat__FifthField: // 28
		return "FifthField"
	case QPixelFormat__AlphaField: // 34
		return "AlphaField"
	case QPixelFormat__AlphaUsageField: // 40
		return "AlphaUsageField"
	case QPixelFormat__AlphaPositionField: // 41
		return "AlphaPositionField"
	case QPixelFormat__PremulField: // 42
		return "PremulField"
	case QPixelFormat__TypeInterpretationField: // 43
		return "TypeInterpretationField"
	case QPixelFormat__ByteOrderField: // 47
		return "ByteOrderField"
	case QPixelFormat__SubEnumField: // 49
		return "SubEnumField"
	case QPixelFormat__UnusedField: // 55
		return "UnusedField"
	case QPixelFormat__TotalFieldWidthByOffsets: // 64
		return "TotalFieldWidthByOffsets"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QPixelFormat_FieldItemName(val int) string {
	var nilthis *QPixelFormat
	return nilthis.FieldItemName(val)
}

/*
This enum type is used to describe the color model of the pixelformat. Alpha was added in 5.5.


*/
type QPixelFormat__ColorModel = int

// The color model is RGB.
const QPixelFormat__RGB QPixelFormat__ColorModel = 0

// This is logically the opposite endian version of RGB. However, for ease of use it has its own model.
const QPixelFormat__BGR QPixelFormat__ColorModel = 1

// The color model uses a color palette.
const QPixelFormat__Indexed QPixelFormat__ColorModel = 2

// The color model is Grayscale.
const QPixelFormat__Grayscale QPixelFormat__ColorModel = 3

// The color model is CMYK.
const QPixelFormat__CMYK QPixelFormat__ColorModel = 4

// The color model is HSL.
const QPixelFormat__HSL QPixelFormat__ColorModel = 5

// The color model is HSV.
const QPixelFormat__HSV QPixelFormat__ColorModel = 6

// The color model is YUV.
const QPixelFormat__YUV QPixelFormat__ColorModel = 7

// There is no color model, only alpha is used.
const QPixelFormat__Alpha QPixelFormat__ColorModel = 8

func (this *QPixelFormat) ColorModelItemName(val int) string {
	switch val {
	case QPixelFormat__RGB: // 0
		return "RGB"
	case QPixelFormat__BGR: // 1
		return "BGR"
	case QPixelFormat__Indexed: // 2
		return "Indexed"
	case QPixelFormat__Grayscale: // 3
		return "Grayscale"
	case QPixelFormat__CMYK: // 4
		return "CMYK"
	case QPixelFormat__HSL: // 5
		return "HSL"
	case QPixelFormat__HSV: // 6
		return "HSV"
	case QPixelFormat__YUV: // 7
		return "YUV"
	case QPixelFormat__Alpha: // 8
		return "Alpha"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QPixelFormat_ColorModelItemName(val int) string {
	var nilthis *QPixelFormat
	return nilthis.ColorModelItemName(val)
}

/*
This enum describes if the alpha channel is used or not. Sometimes the pixelformat will have a size for the alpha channel, but the pixel format does actually not use the alpha channel. For example RGB32 is such a format. The RGB channels are 8 bits each, and there is no alpha channel. But the complete size for each pixel is 32. Therefore the alpha channel size is 8, but the alpha channel is ignored. Its important to note that in such situations the position of the alpha channel is significant.


*/
type QPixelFormat__AlphaUsage = int

// The alpha channel is used.
const QPixelFormat__UsesAlpha QPixelFormat__AlphaUsage = 0

// The alpha channel is not used.
const QPixelFormat__IgnoresAlpha QPixelFormat__AlphaUsage = 1

func (this *QPixelFormat) AlphaUsageItemName(val int) string {
	switch val {
	case QPixelFormat__UsesAlpha: // 0
		return "UsesAlpha"
	case QPixelFormat__IgnoresAlpha: // 1
		return "IgnoresAlpha"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QPixelFormat_AlphaUsageItemName(val int) string {
	var nilthis *QPixelFormat
	return nilthis.AlphaUsageItemName(val)
}

/*
This enum type is used to describe the alpha channels position relative to the color channels.


*/
type QPixelFormat__AlphaPosition = int

// The alpha channel will be put in front of the color channels . E.g. ARGB.
const QPixelFormat__AtBeginning QPixelFormat__AlphaPosition = 0

// The alpha channel will be put in the back of the color channels. E.g. RGBA.
const QPixelFormat__AtEnd QPixelFormat__AlphaPosition = 1

func (this *QPixelFormat) AlphaPositionItemName(val int) string {
	switch val {
	case QPixelFormat__AtBeginning: // 0
		return "AtBeginning"
	case QPixelFormat__AtEnd: // 1
		return "AtEnd"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QPixelFormat_AlphaPositionItemName(val int) string {
	var nilthis *QPixelFormat
	return nilthis.AlphaPositionItemName(val)
}

/*
This enum type describes the boolean state if the alpha channel is multiplied into the color channels or not.


*/
type QPixelFormat__AlphaPremultiplied = int

// The alpha channel is not multiplied into the color channels.
const QPixelFormat__NotPremultiplied QPixelFormat__AlphaPremultiplied = 0

// The alpha channel is multiplied into the color channels.
const QPixelFormat__Premultiplied QPixelFormat__AlphaPremultiplied = 1

func (this *QPixelFormat) AlphaPremultipliedItemName(val int) string {
	switch val {
	case QPixelFormat__NotPremultiplied: // 0
		return "NotPremultiplied"
	case QPixelFormat__Premultiplied: // 1
		return "Premultiplied"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QPixelFormat_AlphaPremultipliedItemName(val int) string {
	var nilthis *QPixelFormat
	return nilthis.AlphaPremultipliedItemName(val)
}

/*
This enum describes how each pixel is interpreted. If a pixel is read as a full 32 bit unsigned integer and then each channel is masked out, or if each byte is read as unsigned char values. Typically QImage formats interpret one pixel as an unsigned integer and then the color channels are masked out. OpenGL on the other hand typically interpreted pixels "one byte after the other", Ie. unsigned byte.

QImage also have the format Format_RGBA8888 (and its derivatives), where the pixels are interpreted as unsigned bytes. OpenGL has extensions that makes it possible to upload pixel buffers in an unsigned integer format.



The image above shows a ARGB pixel in memory read as an unsigned integer. However, if this pixel was read byte for byte on a little endian system the first byte would be the byte containing the B-channel. The next byte would be the G-channel, then the R-channel and finally the A-channel. This shows that on little endian systems, how each pixel is interpreted is significant for integer formats. This is not the case on big endian systems.

ConstantValue
QPixelFormat::UnsignedInteger0
QPixelFormat::UnsignedShort1
QPixelFormat::UnsignedByte2
QPixelFormat::FloatingPoint3

*/
type QPixelFormat__TypeInterpretation = int

//
const QPixelFormat__UnsignedInteger QPixelFormat__TypeInterpretation = 0

//
const QPixelFormat__UnsignedShort QPixelFormat__TypeInterpretation = 1

//
const QPixelFormat__UnsignedByte QPixelFormat__TypeInterpretation = 2

//
const QPixelFormat__FloatingPoint QPixelFormat__TypeInterpretation = 3

func (this *QPixelFormat) TypeInterpretationItemName(val int) string {
	switch val {
	case QPixelFormat__UnsignedInteger: // 0
		return "UnsignedInteger"
	case QPixelFormat__UnsignedShort: // 1
		return "UnsignedShort"
	case QPixelFormat__UnsignedByte: // 2
		return "UnsignedByte"
	case QPixelFormat__FloatingPoint: // 3
		return "FloatingPoint"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QPixelFormat_TypeInterpretationItemName(val int) string {
	var nilthis *QPixelFormat
	return nilthis.TypeInterpretationItemName(val)
}

/*
YUV is not represented by describing the size of the color channels. This is because YUV often use macro pixels, making the concept of sperate color channels invalid. Instead the different YUV layouts are described with this enum.

ConstantValue
QPixelFormat::UYVY6
QPixelFormat::YUYV7

*/
type QPixelFormat__YUVLayout = int

//
const QPixelFormat__YUV444 QPixelFormat__YUVLayout = 0

//
const QPixelFormat__YUV422 QPixelFormat__YUVLayout = 1

//
const QPixelFormat__YUV411 QPixelFormat__YUVLayout = 2

//
const QPixelFormat__YUV420P QPixelFormat__YUVLayout = 3

//
const QPixelFormat__YUV420SP QPixelFormat__YUVLayout = 4

//
const QPixelFormat__YV12 QPixelFormat__YUVLayout = 5

//
const QPixelFormat__UYVY QPixelFormat__YUVLayout = 6

//
const QPixelFormat__YUYV QPixelFormat__YUVLayout = 7

//
const QPixelFormat__NV12 QPixelFormat__YUVLayout = 8

//
const QPixelFormat__NV21 QPixelFormat__YUVLayout = 9

// 0
const QPixelFormat__IMC1 QPixelFormat__YUVLayout = 10

// 1
const QPixelFormat__IMC2 QPixelFormat__YUVLayout = 11

// 2
const QPixelFormat__IMC3 QPixelFormat__YUVLayout = 12

// 3
const QPixelFormat__IMC4 QPixelFormat__YUVLayout = 13

// 4
const QPixelFormat__Y8 QPixelFormat__YUVLayout = 14

// 5
const QPixelFormat__Y16 QPixelFormat__YUVLayout = 15

func (this *QPixelFormat) YUVLayoutItemName(val int) string {
	switch val {
	case QPixelFormat__YUV444: // 0
		return "YUV444"
	case QPixelFormat__YUV422: // 1
		return "YUV422"
	case QPixelFormat__YUV411: // 2
		return "YUV411"
	case QPixelFormat__YUV420P: // 3
		return "YUV420P"
	case QPixelFormat__YUV420SP: // 4
		return "YUV420SP"
	case QPixelFormat__YV12: // 5
		return "YV12"
	case QPixelFormat__UYVY: // 6
		return "UYVY"
	case QPixelFormat__YUYV: // 7
		return "YUYV"
	case QPixelFormat__NV12: // 8
		return "NV12"
	case QPixelFormat__NV21: // 9
		return "NV21"
	case QPixelFormat__IMC1: // 10
		return "IMC1"
	case QPixelFormat__IMC2: // 11
		return "IMC2"
	case QPixelFormat__IMC3: // 12
		return "IMC3"
	case QPixelFormat__IMC4: // 13
		return "IMC4"
	case QPixelFormat__Y8: // 14
		return "Y8"
	case QPixelFormat__Y16: // 15
		return "Y16"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QPixelFormat_YUVLayoutItemName(val int) string {
	var nilthis *QPixelFormat
	return nilthis.YUVLayoutItemName(val)
}

/*
This enum describes the ByteOrder of the pixel format. This enum is mostly ignored but have some use cases for YUV formats. BGR formats have their own color model, and should not be described by using the opposite endianness on an RGB format.


*/
type QPixelFormat__ByteOrder = int

// The byte order is little endian.
const QPixelFormat__LittleEndian QPixelFormat__ByteOrder = 0

// The byte order is big endian.
const QPixelFormat__BigEndian QPixelFormat__ByteOrder = 1

// This enum will not be stored, but is converted in the constructor to the endian enum that matches the enum of the current system.
const QPixelFormat__CurrentSystemEndian QPixelFormat__ByteOrder = 2

func (this *QPixelFormat) ByteOrderItemName(val int) string {
	switch val {
	case QPixelFormat__LittleEndian: // 0
		return "LittleEndian"
	case QPixelFormat__BigEndian: // 1
		return "BigEndian"
	case QPixelFormat__CurrentSystemEndian: // 2
		return "CurrentSystemEndian"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QPixelFormat_ByteOrderItemName(val int) string {
	var nilthis *QPixelFormat
	return nilthis.ByteOrderItemName(val)
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
