package qtcore

// /usr/include/qt/QtCore/qrect.h
// #include <qrect.h>
// #include <QtCore>

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

//  ext block end

//  body block begin

/*
 */
// size 16
type QRect struct {
	*qtrt.CObject
}
type QRect_ITF interface {
	QRect_PTR() *QRect
}

func (ptr *QRect) QRect_PTR() *QRect { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QRectFromptr(cthis Voidptr) *QRect {
	return &QRect{&qtrt.CObject{cthis}}
}
func (*QRect) Fromptr(cthis Voidptr) *QRect {
	return QRectFromptr(cthis)
}

// /usr/include/qt/QtCore/qrect.h:65
// index:0
// Public inline Extend Visibility=Default Availability=Available
// [1] bool isNull() const

/*
 */
func (this *QRect) IsNull() bool {
	rv, err := qtrt.Qtcc3(3424420210, "_ZNK5QRect6isNullEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qrect.h:66
// index:0
// Public inline Extend Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
 */
func (this *QRect) IsEmpty() bool {
	rv, err := qtrt.Qtcc3(1044706320, "_ZNK5QRect7isEmptyEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qrect.h:67
// index:0
// Public inline Extend Visibility=Default Availability=Available
// [1] bool isValid() const

/*
 */
func (this *QRect) IsValid() bool {
	rv, err := qtrt.Qtcc3(204306394, "_ZNK5QRect7isValidEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qrect.h:69
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int left() const

/*
 */
func (this *QRect) Left() int {
	rv, err := qtrt.Qtcc3(3169620863, "_ZNK5QRect4leftEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtCore/qrect.h:70
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int top() const

/*
 */
func (this *QRect) Top() int {
	rv, err := qtrt.Qtcc3(3329531928, "_ZNK5QRect3topEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtCore/qrect.h:71
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int right() const

/*
 */
func (this *QRect) Right() int {
	rv, err := qtrt.Qtcc3(3515267020, "_ZNK5QRect5rightEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtCore/qrect.h:72
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int bottom() const

/*
 */
func (this *QRect) Bottom() int {
	rv, err := qtrt.Qtcc3(3137250108, "_ZNK5QRect6bottomEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtCore/qrect.h:75
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int x() const

/*
 */
func (this *QRect) X() int {
	rv, err := qtrt.Qtcc3(1050399888, "_ZNK5QRect1xEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtCore/qrect.h:76
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int y() const

/*
 */
func (this *QRect) Y() int {
	rv, err := qtrt.Qtcc3(1062846119, "_ZNK5QRect1yEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtCore/qrect.h:81
// index:0
// Public inline Ignore Visibility=Default Availability=Available
// [-2] void setX(int)

/*
 */
func (this *QRect) SetX(x int) {
	rv, err := qtrt.Qtcc3(2433102519, "_ZN5QRect4setXEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&x))
	qtrt.ErrPrint3(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:82
// index:0
// Public inline Ignore Visibility=Default Availability=Available
// [-2] void setY(int)

/*
 */
func (this *QRect) SetY(y int) {
	rv, err := qtrt.Qtcc3(2428786816, "_ZN5QRect4setYEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&y))
	qtrt.ErrPrint3(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:126
// index:0
// Public inline Ignore Visibility=Default Availability=Available
// [-2] void setWidth(int)

/*
 */
func (this *QRect) SetWidth(w int) {
	rv, err := qtrt.Qtcc3(940159144, "_ZN5QRect8setWidthEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&w))
	qtrt.ErrPrint3(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:127
// index:0
// Public inline Ignore Visibility=Default Availability=Available
// [-2] void setHeight(int)

/*
 */
func (this *QRect) SetHeight(h int) {
	rv, err := qtrt.Qtcc3(1082198786, "_ZN5QRect9setHeightEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&h))
	qtrt.ErrPrint3(err, rv)
}

func DeleteQRect(this *QRect) {
	rv, err := qtrt.Qtcc3(3544253908, "_ZdlPv", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10059() {
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
}

//  keep block end
