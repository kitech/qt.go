package qtgui

// /usr/include/qt/QtGui/qpixmap.h
// #include <qpixmap.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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
type QPixmap struct {
	*QPaintDevice
}
type QPixmap_ITF interface {
	QPaintDevice_ITF
	QPixmap_PTR() *QPixmap
}

func (ptr *QPixmap) QPixmap_PTR() *QPixmap { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QPixmapFromptr(cthis Voidptr) *QPixmap {
	bcthis0 := QPaintDeviceFromptr(cthis)
	return &QPixmap{bcthis0}
}
func (*QPixmap) Fromptr(cthis Voidptr) *QPixmap {
	return QPixmapFromptr(cthis)
}

// /usr/include/qt/QtGui/qpixmap.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPixmap()

/*
 */
func (*QPixmap) NewForInherit() *QPixmap {
	return NewQPixmap()
}
func NewQPixmap() *QPixmap {
	cthis := qtrt.Malloc(32)
	rv, err := qtrt.Qtcc3(440989401, "_ZN7QPixmapC2Ev", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, Voidptr(&cthis))
	qtrt.ErrPrint3(err, rv)
	gothis := QPixmapFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQPixmap)
	return gothis
}

// /usr/include/qt/QtGui/qpixmap.h:66
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPixmap(int, int)

/*
 */
func (*QPixmap) NewForInherit1(w int, h int) *QPixmap {
	return NewQPixmap1(w, h)
}
func NewQPixmap1(w int, h int) *QPixmap {
	cthis := qtrt.Malloc(32)
	rv, err := qtrt.Qtcc3(3559070313, "_ZN7QPixmapC2Eii", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, Voidptr(&cthis), Voidptr(&w), Voidptr(&h))
	qtrt.ErrPrint3(err, rv)
	gothis := QPixmapFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQPixmap)
	return gothis
}

// /usr/include/qt/QtGui/qpixmap.h:86
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int width() const

/*
 */
func (this *QPixmap) Width() int {
	rv, err := qtrt.Qtcc3(494131390, "_ZNK7QPixmap5widthEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtGui/qpixmap.h:87
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int height() const

/*
 */
func (this *QPixmap) Height() int {
	rv, err := qtrt.Qtcc3(1443543034, "_ZNK7QPixmap6heightEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtGui/qpixmap.h:89
// index:0
// Public Direct Visibility=Default Availability=Available
// [16] QRect rect() const

/*
 */
func (this *QPixmap) Rect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.Qtcc3(4207271430, "_ZNK7QPixmap4rectEv", qtrt.FFITO_UINT128,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	cthis := qtrt.Malloc(16)
	qtrt.Cmemcpy(cthis, rv.Addr(), 16)
	rv2 := qtcore.QRectFromptr(cthis) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:94
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void fill(const QColor &)

/*
 */
func (this *QPixmap) Fill(fillColor QColor_ITF) {
	var convArg0 Voidptr
	if fillColor != nil && fillColor.QColor_PTR() != nil {
		convArg0 = fillColor.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(2168313289, "_ZN7QPixmap4fillERK6QColor", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:94
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void fill(const QColor &)

/*
 */
func (this *QPixmap) Fillp() {
	// arg: 0, const QColor &=LValueReference, QColor=Record, , Invalid
	var convArg0 Voidptr
	rv, err := qtrt.Qtcc3(2168313289, "_ZN7QPixmap4fillERK6QColor", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
}

func DeleteQPixmap(this *QPixmap) {
	rv, err := qtrt.Qtcc3(2275397216, "_ZN7QPixmapD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10149() {
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
