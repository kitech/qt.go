package qtcore

// /usr/include/qt/QtCore/qpoint.h
// #include <qpoint.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"

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
	*qtrt.CObject
}

func (this *QPoint) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPoint) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQPointFromPointer(cthis unsafe.Pointer) *QPoint {
	return &QPoint{&qtrt.CObject{cthis}}
}
func (*QPoint) NewFromPointer(cthis unsafe.Pointer) *QPoint {
	return NewQPointFromPointer(cthis)
}

// /usr/include/qt/QtCore/qpoint.h:55
// index:0
// Public inline
// void QPoint()
func NewQPoint() *QPoint {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QPointC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQPointFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qpoint.h:56
// index:1
// Public inline
// void QPoint(int, int)
func NewQPoint_1(xpos int, ypos int) *QPoint {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QPointC2Eii", ffiqt.FFI_TYPE_VOID, cthis, xpos, ypos)
	gopp.ErrPrint(err, rv)
	gothis := NewQPointFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qpoint.h:58
// index:0
// Public inline
// bool isNull()
func (this *QPoint) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QPoint6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qpoint.h:60
// index:0
// Public inline
// int x()
func (this *QPoint) X() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QPoint1xEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qpoint.h:61
// index:0
// Public inline
// int y()
func (this *QPoint) Y() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QPoint1yEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qpoint.h:62
// index:0
// Public inline
// void setX(int)
func (this *QPoint) SetX(x int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QPoint4setXEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpoint.h:63
// index:0
// Public inline
// void setY(int)
func (this *QPoint) SetY(y int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QPoint4setYEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpoint.h:65
// index:0
// Public inline
// int manhattanLength()
func (this *QPoint) ManhattanLength() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QPoint15manhattanLengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qpoint.h:67
// index:0
// Public inline
// int & rx()
func (this *QPoint) Rx() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QPoint2rxEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 3331
}

// /usr/include/qt/QtCore/qpoint.h:68
// index:0
// Public inline
// int & ry()
func (this *QPoint) Ry() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QPoint2ryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 3331
}

// /usr/include/qt/QtCore/qpoint.h:79
// index:0
// Public static inline
// int dotProduct(const class QPoint &, const class QPoint &)
func (this *QPoint) DotProduct(p1 *QPoint, p2 *QPoint) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QPoint10dotProductERKS_S1_", ffiqt.FFI_TYPE_POINTER, p1, p2)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QPoint_DotProduct(p1 *QPoint, p2 *QPoint) int {
	var nilthis *QPoint
	rv := nilthis.DotProduct(p1, p2)
	return rv
}

//  body block end
