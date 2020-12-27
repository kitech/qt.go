package qtgui

// /usr/include/qt/QtGui/qcolor.h
// #include <qcolor.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
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
// size 16
type QColor struct {
	*qtrt.CObject
}
type QColor_ITF interface {
	QColor_PTR() *QColor
}

func (ptr *QColor) QColor_PTR() *QColor { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QColorFromptr(cthis Voidptr) *QColor {
	return &QColor{&qtrt.CObject{cthis}}
}
func (*QColor) Fromptr(cthis Voidptr) *QColor {
	return QColorFromptr(cthis)
}

// /usr/include/qt/QtGui/qcolor.h:70
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QColor()

/*
 */
func (*QColor) NewForInherit() *QColor {
	return NewQColor()
}
func NewQColor() *QColor {
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(2667467580, "_ZN6QColorC2Ev", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, Voidptr(&cthis))
	qtrt.ErrPrint3(err, rv)
	gothis := QColorFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQColor)
	return gothis
}

// /usr/include/qt/QtGui/qcolor.h:73
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QColor(int, int, int, int)

/*
 */
func (*QColor) NewForInherit1(r int, g int, b int, a int) *QColor {
	return NewQColor1(r, g, b, a)
}
func NewQColor1(r int, g int, b int, a int) *QColor {
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(3599673984, "_ZN6QColorC2Eiiii", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, qtrt.FFITO_INT, qtrt.FFITO_INT, Voidptr(&cthis), Voidptr(&r), Voidptr(&g), Voidptr(&b), Voidptr(&a))
	qtrt.ErrPrint3(err, rv)
	gothis := QColorFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQColor)
	return gothis
}

// /usr/include/qt/QtGui/qcolor.h:73
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QColor(int, int, int, int)

/*
 */
func (*QColor) NewForInherit1p(r int, g int, b int) *QColor {
	return NewQColor1p(r, g, b)
}
func NewQColor1p(r int, g int, b int) *QColor {
	// arg: 3, int=Int, =Invalid, , Invalid
	a := int(255)
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(3599673984, "_ZN6QColorC2Eiiii", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, qtrt.FFITO_INT, qtrt.FFITO_INT, Voidptr(&cthis), Voidptr(&r), Voidptr(&g), Voidptr(&b), Voidptr(&a))
	qtrt.ErrPrint3(err, rv)
	gothis := QColorFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQColor)
	return gothis
}

func DeleteQColor(this *QColor) {
	rv, err := qtrt.Qtcc3(53046149, "_ZN6QColorD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

/*


 */
type QColor__Spec = int

//
const QColor__Invalid QColor__Spec = 0

//
const QColor__Rgb QColor__Spec = 1

//
const QColor__Hsv QColor__Spec = 2

//
const QColor__Cmyk QColor__Spec = 3

//
const QColor__Hsl QColor__Spec = 4

//
const QColor__ExtendedRgb QColor__Spec = 5

func (this *QColor) SpecItemName(val int) string {
	switch val {
	case QColor__Invalid: // 0
		return "Invalid"
	case QColor__Rgb: // 1
		return "Rgb"
	case QColor__Hsv: // 2
		return "Hsv"
	case QColor__Cmyk: // 3
		return "Cmyk"
	case QColor__Hsl: // 4
		return "Hsl"
	case QColor__ExtendedRgb: // 5
		return "ExtendedRgb"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QColor_SpecItemName(val int) string {
	var nilthis *QColor
	return nilthis.SpecItemName(val)
}

/*


 */
type QColor__NameFormat = int

//
const QColor__HexRgb QColor__NameFormat = 0

//
const QColor__HexArgb QColor__NameFormat = 1

func (this *QColor) NameFormatItemName(val int) string {
	switch val {
	case QColor__HexRgb: // 0
		return "HexRgb"
	case QColor__HexArgb: // 1
		return "HexArgb"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QColor_NameFormatItemName(val int) string {
	var nilthis *QColor
	return nilthis.NameFormatItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10065() {
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
