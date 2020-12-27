package qtcore

// /usr/include/qt/QtCore/qpoint.h
// #include <qpoint.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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
type QPointF struct {
	*qtrt.CObject
}
type QPointF_ITF interface {
	QPointF_PTR() *QPointF
}

func (ptr *QPointF) QPointF_PTR() *QPointF { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QPointFFromptr(cthis Voidptr) *QPointF {
	return &QPointF{&qtrt.CObject{cthis}}
}
func (*QPointF) Fromptr(cthis Voidptr) *QPointF {
	return QPointFFromptr(cthis)
}

// /usr/include/qt/QtCore/qpoint.h:228
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [8] qreal manhattanLength() const

/*
 */
func (this *QPointF) ManhattanLength() float64 {
	rv, err := qtrt.Qtcc3(2157545085, "_ZNK7QPointF15manhattanLengthEv", qtrt.FFITO_DOUBLE,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Float64() // 1111
}

// /usr/include/qt/QtCore/qpoint.h:230
// index:0
// Public inline Extend Visibility=Default Availability=Available
// [1] bool isNull() const

/*
 */
func (this *QPointF) IsNull() bool {
	rv, err := qtrt.Qtcc3(1609692812, "_ZNK7QPointF6isNullEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qpoint.h:232
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [8] qreal x() const

/*
 */
func (this *QPointF) X() float64 {
	rv, err := qtrt.Qtcc3(92750711, "_ZNK7QPointF1xEv", qtrt.FFITO_DOUBLE,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Float64() // 1111
}

// /usr/include/qt/QtCore/qpoint.h:233
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [8] qreal y() const

/*
 */
func (this *QPointF) Y() float64 {
	rv, err := qtrt.Qtcc3(71641408, "_ZNK7QPointF1yEv", qtrt.FFITO_DOUBLE,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Float64() // 1111
}

// /usr/include/qt/QtCore/qpoint.h:234
// index:0
// Public inline Ignore Visibility=Default Availability=Available
// [-2] void setX(qreal)

/*
 */
func (this *QPointF) SetX(x float64) {
	rv, err := qtrt.Qtcc3(2852303336, "_ZN7QPointF4setXEd", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&x))
	qtrt.ErrPrint3(err, rv)
}

// /usr/include/qt/QtCore/qpoint.h:235
// index:0
// Public inline Ignore Visibility=Default Availability=Available
// [-2] void setY(qreal)

/*
 */
func (this *QPointF) SetY(y float64) {
	rv, err := qtrt.Qtcc3(2881543135, "_ZN7QPointF4setYEd", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&y))
	qtrt.ErrPrint3(err, rv)
}

func DeleteQPointF(this *QPointF) {
	rv, err := qtrt.Qtcc3(134628044, "_ZN7QPointFD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10041() {
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
