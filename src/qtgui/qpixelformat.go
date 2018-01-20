//  header block begin
// /usr/include/qt/QtGui/qpixelformat.h
// #include <qpixelformat.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
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
type QPixelFormat struct {
	*qtrt.CObject
}

func (this *QPixelFormat) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQPixelFormatFromPointer(cthis unsafe.Pointer) *QPixelFormat {
	return &QPixelFormat{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qpixelformat.h:163
// index:0
// Public inline
// void QPixelFormat()
func NewQPixelFormat() *QPixelFormat {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPixelFormatC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQPixelFormatFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpixelformat.h:164
// index:1
// Public inline
// void QPixelFormat(enum QPixelFormat::ColorModel, uchar, uchar, uchar, uchar, uchar, uchar, enum QPixelFormat::AlphaUsage, enum QPixelFormat::AlphaPosition, enum QPixelFormat::AlphaPremultiplied, enum QPixelFormat::TypeInterpretation, enum QPixelFormat::ByteOrder, uchar)
func NewQPixelFormat_1(colorModel int, firstSize byte, secondSize byte, thirdSize byte, fourthSize byte, fifthSize byte, alphaSize byte, alphaUsage int, alphaPosition int, premultiplied int, typeInterpretation int, byteOrder int, subEnum byte) *QPixelFormat {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPixelFormatC2ENS_10ColorModelEhhhhhhNS_10AlphaUsageENS_13AlphaPositionENS_18AlphaPremultipliedENS_18TypeInterpretationENS_9ByteOrderEh", ffiqt.FFI_TYPE_VOID, cthis, &colorModel, &firstSize, &secondSize, &thirdSize, &fourthSize, &fifthSize, &alphaSize, &alphaUsage, &alphaPosition, &premultiplied, &typeInterpretation, &byteOrder, &subEnum)
	gopp.ErrPrint(err, rv)
	gothis := NewQPixelFormatFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpixelformat.h:178
// index:0
// Public inline
// QPixelFormat::ColorModel colorModel()
func (this *QPixelFormat) ColorModel() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat10colorModelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpixelformat.h:179
// index:0
// Public inline
// uchar channelCount()
func (this *QPixelFormat) ChannelCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat12channelCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpixelformat.h:186
// index:0
// Public inline
// uchar redSize()
func (this *QPixelFormat) RedSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat7redSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpixelformat.h:187
// index:0
// Public inline
// uchar greenSize()
func (this *QPixelFormat) GreenSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat9greenSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpixelformat.h:188
// index:0
// Public inline
// uchar blueSize()
func (this *QPixelFormat) BlueSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat8blueSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpixelformat.h:190
// index:0
// Public inline
// uchar cyanSize()
func (this *QPixelFormat) CyanSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat8cyanSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpixelformat.h:191
// index:0
// Public inline
// uchar magentaSize()
func (this *QPixelFormat) MagentaSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat11magentaSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpixelformat.h:192
// index:0
// Public inline
// uchar yellowSize()
func (this *QPixelFormat) YellowSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat10yellowSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpixelformat.h:193
// index:0
// Public inline
// uchar blackSize()
func (this *QPixelFormat) BlackSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat9blackSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpixelformat.h:195
// index:0
// Public inline
// uchar hueSize()
func (this *QPixelFormat) HueSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat7hueSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpixelformat.h:196
// index:0
// Public inline
// uchar saturationSize()
func (this *QPixelFormat) SaturationSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat14saturationSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpixelformat.h:197
// index:0
// Public inline
// uchar lightnessSize()
func (this *QPixelFormat) LightnessSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat13lightnessSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpixelformat.h:198
// index:0
// Public inline
// uchar brightnessSize()
func (this *QPixelFormat) BrightnessSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat14brightnessSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpixelformat.h:200
// index:0
// Public inline
// uchar alphaSize()
func (this *QPixelFormat) AlphaSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat9alphaSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpixelformat.h:202
// index:0
// Public inline
// uchar bitsPerPixel()
func (this *QPixelFormat) BitsPerPixel() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat12bitsPerPixelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpixelformat.h:209
// index:0
// Public inline
// QPixelFormat::AlphaUsage alphaUsage()
func (this *QPixelFormat) AlphaUsage() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat10alphaUsageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpixelformat.h:210
// index:0
// Public inline
// QPixelFormat::AlphaPosition alphaPosition()
func (this *QPixelFormat) AlphaPosition() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat13alphaPositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpixelformat.h:211
// index:0
// Public inline
// QPixelFormat::AlphaPremultiplied premultiplied()
func (this *QPixelFormat) Premultiplied() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat13premultipliedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpixelformat.h:212
// index:0
// Public inline
// QPixelFormat::TypeInterpretation typeInterpretation()
func (this *QPixelFormat) TypeInterpretation() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat18typeInterpretationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpixelformat.h:213
// index:0
// Public inline
// QPixelFormat::ByteOrder byteOrder()
func (this *QPixelFormat) ByteOrder() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat9byteOrderEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpixelformat.h:215
// index:0
// Public inline
// QPixelFormat::YUVLayout yuvLayout()
func (this *QPixelFormat) YuvLayout() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat9yuvLayoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpixelformat.h:216
// index:0
// Public inline
// uchar subEnum()
func (this *QPixelFormat) SubEnum() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat7subEnumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
