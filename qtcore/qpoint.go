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
// extern C begin: 1
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
// size 8
type QPoint struct {
	*qtrt.CObject
}
type QPoint_ITF interface {
	QPoint_PTR() *QPoint
}

func (ptr *QPoint) QPoint_PTR() *QPoint { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QPointFromptr(cthis Voidptr) *QPoint {
	return &QPoint{&qtrt.CObject{cthis}}
}
func (*QPoint) Fromptr(cthis Voidptr) *QPoint {
	return QPointFromptr(cthis)
}

// /usr/include/qt/QtCore/qpoint.h:58
// index:0
// Public inline Extend Visibility=Default Availability=Available
// [1] bool isNull() const

/*
 */
func (this *QPoint) IsNull() bool {
	rv, err := qtrt.Qtcc3(1390804939, "_ZNK6QPoint6isNullEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qpoint.h:60
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int x() const

/*
 */
func (this *QPoint) X() int {
	rv, err := qtrt.Qtcc3(359386261, "_ZNK6QPoint1xEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtCore/qpoint.h:61
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int y() const

/*
 */
func (this *QPoint) Y() int {
	rv, err := qtrt.Qtcc3(346662562, "_ZNK6QPoint1yEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtCore/qpoint.h:62
// index:0
// Public inline Ignore Visibility=Default Availability=Available
// [-2] void setX(int)

/*
 */
func (this *QPoint) SetX(x int) {
	rv, err := qtrt.Qtcc3(148123695, "_ZN6QPoint4setXEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&x))
	qtrt.ErrPrint3(err, rv)
}

// /usr/include/qt/QtCore/qpoint.h:63
// index:0
// Public inline Ignore Visibility=Default Availability=Available
// [-2] void setY(int)

/*
 */
func (this *QPoint) SetY(y int) {
	rv, err := qtrt.Qtcc3(152459800, "_ZN6QPoint4setYEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&y))
	qtrt.ErrPrint3(err, rv)
}

// /usr/include/qt/QtCore/qpoint.h:65
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int manhattanLength() const

/*
 */
func (this *QPoint) ManhattanLength() int {
	rv, err := qtrt.Qtcc3(3380245230, "_ZNK6QPoint15manhattanLengthEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

func DeleteQPoint(this *QPoint) {
	rv, err := qtrt.Qtcc3(2151685267, "_ZN6QPointD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10039() {
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
