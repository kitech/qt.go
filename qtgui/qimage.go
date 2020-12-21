package qtgui

// /usr/include/qt/QtGui/qimage.h
// #include <qimage.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
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
// size 32
type QImage struct {
	*QPaintDevice
}
type QImage_ITF interface {
	QPaintDevice_ITF
	QImage_PTR() *QImage
}

func (ptr *QImage) QImage_PTR() *QImage { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QImageFromptr(cthis Voidptr) *QImage {
	bcthis0 := QPaintDeviceFromptr(cthis)
	return &QImage{bcthis0}
}
func (*QImage) Fromptr(cthis Voidptr) *QImage {
	return QImageFromptr(cthis)
}

// /usr/include/qt/QtGui/qimage.h:141
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QImage()

/*
 */
func (*QImage) NewForInherit() *QImage {
	return NewQImage()
}
func NewQImage() *QImage {
	cthis := qtrt.Malloc(32)
	rv, err := qtrt.Qtcc3(380729056, "_ZN6QImageC2Ev", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, Voidptr(&cthis))
	qtrt.ErrPrint2(err, rv)
	gothis := QImageFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQImage)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:142
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QImage(const QSize &, QImage::Format)

/*
 */
func (*QImage) NewForInherit1(size qtcore.QSize_ITF, format int) *QImage {
	return NewQImage1(size, format)
}
func NewQImage1(size qtcore.QSize_ITF, format int) *QImage {
	var convArg0 Voidptr
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(32)
	rv, err := qtrt.Qtcc3(2386780666, "_ZN6QImageC2ERK5QSizeNS_6FormatE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&format))
	qtrt.ErrPrint2(err, rv)
	gothis := QImageFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQImage)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:143
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QImage(int, int, QImage::Format)

/*
 */
func (*QImage) NewForInherit2(width int, height int, format int) *QImage {
	return NewQImage2(width, height, format)
}
func NewQImage2(width int, height int, format int) *QImage {
	cthis := qtrt.Malloc(32)
	rv, err := qtrt.Qtcc3(355905797, "_ZN6QImageC2EiiNS_6FormatE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&width), Voidptr(&height), Voidptr(&format))
	qtrt.ErrPrint2(err, rv)
	gothis := QImageFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQImage)
	return gothis
}

func DeleteQImage(this *QImage) {
	rv, err := qtrt.Qtcc1(0, "_ZN6QImageD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QImage__InvertMode = int

//
const QImage__InvertRgb QImage__InvertMode = 0

//
const QImage__InvertRgba QImage__InvertMode = 1

func (this *QImage) InvertModeItemName(val int) string {
	switch val {
	case QImage__InvertRgb: // 0
		return "InvertRgb"
	case QImage__InvertRgba: // 1
		return "InvertRgba"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QImage_InvertModeItemName(val int) string {
	var nilthis *QImage
	return nilthis.InvertModeItemName(val)
}

/*


 */
type QImage__Format = int

//
const QImage__Format_Invalid QImage__Format = 0

//
const QImage__Format_Mono QImage__Format = 1

//
const QImage__Format_MonoLSB QImage__Format = 2

//
const QImage__Format_Indexed8 QImage__Format = 3

//
const QImage__Format_RGB32 QImage__Format = 4

//
const QImage__Format_ARGB32 QImage__Format = 5

//
const QImage__Format_ARGB32_Premultiplied QImage__Format = 6

//
const QImage__Format_RGB16 QImage__Format = 7

//
const QImage__Format_ARGB8565_Premultiplied QImage__Format = 8

//
const QImage__Format_RGB666 QImage__Format = 9

//
const QImage__Format_ARGB6666_Premultiplied QImage__Format = 10

//
const QImage__Format_RGB555 QImage__Format = 11

//
const QImage__Format_ARGB8555_Premultiplied QImage__Format = 12

//
const QImage__Format_RGB888 QImage__Format = 13

//
const QImage__Format_RGB444 QImage__Format = 14

//
const QImage__Format_ARGB4444_Premultiplied QImage__Format = 15

//
const QImage__Format_RGBX8888 QImage__Format = 16

//
const QImage__Format_RGBA8888 QImage__Format = 17

//
const QImage__Format_RGBA8888_Premultiplied QImage__Format = 18

//
const QImage__Format_BGR30 QImage__Format = 19

//
const QImage__Format_A2BGR30_Premultiplied QImage__Format = 20

//
const QImage__Format_RGB30 QImage__Format = 21

//
const QImage__Format_A2RGB30_Premultiplied QImage__Format = 22

//
const QImage__Format_Alpha8 QImage__Format = 23

//
const QImage__Format_Grayscale8 QImage__Format = 24

//
const QImage__Format_RGBX64 QImage__Format = 25

//
const QImage__Format_RGBA64 QImage__Format = 26

//
const QImage__Format_RGBA64_Premultiplied QImage__Format = 27

//
const QImage__Format_Grayscale16 QImage__Format = 28

//
const QImage__Format_BGR888 QImage__Format = 29

//
const QImage__NImageFormats QImage__Format = 30

func (this *QImage) FormatItemName(val int) string {
	switch val {
	case QImage__Format_Invalid: // 0
		return "Format_Invalid"
	case QImage__Format_Mono: // 1
		return "Format_Mono"
	case QImage__Format_MonoLSB: // 2
		return "Format_MonoLSB"
	case QImage__Format_Indexed8: // 3
		return "Format_Indexed8"
	case QImage__Format_RGB32: // 4
		return "Format_RGB32"
	case QImage__Format_ARGB32: // 5
		return "Format_ARGB32"
	case QImage__Format_ARGB32_Premultiplied: // 6
		return "Format_ARGB32_Premultiplied"
	case QImage__Format_RGB16: // 7
		return "Format_RGB16"
	case QImage__Format_ARGB8565_Premultiplied: // 8
		return "Format_ARGB8565_Premultiplied"
	case QImage__Format_RGB666: // 9
		return "Format_RGB666"
	case QImage__Format_ARGB6666_Premultiplied: // 10
		return "Format_ARGB6666_Premultiplied"
	case QImage__Format_RGB555: // 11
		return "Format_RGB555"
	case QImage__Format_ARGB8555_Premultiplied: // 12
		return "Format_ARGB8555_Premultiplied"
	case QImage__Format_RGB888: // 13
		return "Format_RGB888"
	case QImage__Format_RGB444: // 14
		return "Format_RGB444"
	case QImage__Format_ARGB4444_Premultiplied: // 15
		return "Format_ARGB4444_Premultiplied"
	case QImage__Format_RGBX8888: // 16
		return "Format_RGBX8888"
	case QImage__Format_RGBA8888: // 17
		return "Format_RGBA8888"
	case QImage__Format_RGBA8888_Premultiplied: // 18
		return "Format_RGBA8888_Premultiplied"
	case QImage__Format_BGR30: // 19
		return "Format_BGR30"
	case QImage__Format_A2BGR30_Premultiplied: // 20
		return "Format_A2BGR30_Premultiplied"
	case QImage__Format_RGB30: // 21
		return "Format_RGB30"
	case QImage__Format_A2RGB30_Premultiplied: // 22
		return "Format_A2RGB30_Premultiplied"
	case QImage__Format_Alpha8: // 23
		return "Format_Alpha8"
	case QImage__Format_Grayscale8: // 24
		return "Format_Grayscale8"
	case QImage__Format_RGBX64: // 25
		return "Format_RGBX64"
	case QImage__Format_RGBA64: // 26
		return "Format_RGBA64"
	case QImage__Format_RGBA64_Premultiplied: // 27
		return "Format_RGBA64_Premultiplied"
	case QImage__Format_Grayscale16: // 28
		return "Format_Grayscale16"
	case QImage__Format_BGR888: // 29
		return "Format_BGR888"
	case QImage__NImageFormats: // 30
		return "NImageFormats"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QImage_FormatItemName(val int) string {
	var nilthis *QImage
	return nilthis.FormatItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10053() {
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
