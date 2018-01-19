//  header block begin
// /usr/include/qt/QtCore/qpoint.h
// #include <qpoint.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
}

//  ext block end

//  body block begin
type QPoint struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qpoint.h:55
// index:0
// inline
// void QPoint()
func NewQPoint() *QPoint {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QPointC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QPoint{cthis}
}

// /usr/include/qt/QtCore/qpoint.h:56
// index:1
// inline
// void QPoint(int, int)
func NewQPoint_1(xpos int, ypos int) *QPoint {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QPointC2Eii", ffiqt.FFI_TYPE_VOID, cthis, &xpos, &ypos)
	gopp.ErrPrint(err, rv)
	return &QPoint{cthis}
}

// /usr/include/qt/QtCore/qpoint.h:58
// index:0
// inline
// bool isNull()
func (this *QPoint) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QPoint6isNullEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpoint.h:60
// index:0
// inline
// int x()
func (this *QPoint) X() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QPoint1xEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpoint.h:61
// index:0
// inline
// int y()
func (this *QPoint) Y() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QPoint1yEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpoint.h:62
// index:0
// inline
// void setX(int)
func (this *QPoint) SetX(x int) {
	// 0: (, int x), (&x)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QPoint4setXEi", ffiqt.FFI_TYPE_VOID, this.cthis, &x)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpoint.h:63
// index:0
// inline
// void setY(int)
func (this *QPoint) SetY(y int) {
	// 0: (, int y), (&y)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QPoint4setYEi", ffiqt.FFI_TYPE_VOID, this.cthis, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpoint.h:65
// index:0
// inline
// int manhattanLength()
func (this *QPoint) ManhattanLength() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QPoint15manhattanLengthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpoint.h:67
// index:0
// inline
// int & rx()
func (this *QPoint) Rx() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QPoint2rxEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpoint.h:68
// index:0
// inline
// int & ry()
func (this *QPoint) Ry() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QPoint2ryEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpoint.h:79
// index:0
// static inline
// int dotProduct(const class QPoint &, const class QPoint &)
func (this *QPoint) DotProduct(p1 unsafe.Pointer, p2 unsafe.Pointer) {
	// 0: (const QPoint & p1, const QPoint & p2), (p1, p2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QPoint10dotProductERKS_S1_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPoint_DotProduct(p1 unsafe.Pointer, p2 unsafe.Pointer) {
	// 0: (const QPoint & p1, const QPoint & p2), (p1, p2)
	var nilthis *QPoint
	nilthis.DotProduct(p1, p2)
}

//  body block end
