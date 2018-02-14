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
// extern C begin: 19
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QPoint struct {
	*qtrt.CObject
}
type QPoint_ITF interface {
	QPoint_PTR() *QPoint
}

func (ptr *QPoint) QPoint_PTR() *QPoint { return ptr }

func (this *QPoint) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPoint) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQPointFromPointer(cthis unsafe.Pointer) *QPoint {
	return &QPoint{&qtrt.CObject{cthis}}
}
func (*QPoint) NewFromPointer(cthis unsafe.Pointer) *QPoint {
	return NewQPointFromPointer(cthis)
}

// /usr/include/qt/QtCore/qpoint.h:55
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QPoint()
func NewQPoint() *QPoint {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QPointC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPointFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPoint)
	return gothis
}

// /usr/include/qt/QtCore/qpoint.h:56
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QPoint(int, int)
func NewQPoint_1(xpos int, ypos int) *QPoint {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QPointC2Eii", qtrt.FFI_TYPE_POINTER, xpos, ypos)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPointFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPoint)
	return gothis
}

// /usr/include/qt/QtCore/qpoint.h:58
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QPoint) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QPoint6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qpoint.h:60
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int x()
func (this *QPoint) X() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QPoint1xEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qpoint.h:61
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int y()
func (this *QPoint) Y() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QPoint1yEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qpoint.h:62
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setX(int)
func (this *QPoint) SetX(x int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QPoint4setXEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpoint.h:63
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setY(int)
func (this *QPoint) SetY(y int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QPoint4setYEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpoint.h:65
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int manhattanLength()
func (this *QPoint) ManhattanLength() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QPoint15manhattanLengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qpoint.h:67
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int & rx()
func (this *QPoint) Rx() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QPoint2rxEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cpretval2go("int", rv).(int) // 3331
}

// /usr/include/qt/QtCore/qpoint.h:68
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int & ry()
func (this *QPoint) Ry() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QPoint2ryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cpretval2go("int", rv).(int) // 3331
}

// /usr/include/qt/QtCore/qpoint.h:79
// index:0
// Public static inline Visibility=Default Availability=Available
// [4] int dotProduct(const QPoint &, const QPoint &)
func (this *QPoint) DotProduct(p1 QPoint_ITF, p2 QPoint_ITF) int {
	var convArg0 = p1.QPoint_PTR().GetCthis()
	var convArg1 = p2.QPoint_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QPoint10dotProductERKS_S1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QPoint_DotProduct(p1 QPoint_ITF, p2 QPoint_ITF) int {
	var nilthis *QPoint
	rv := nilthis.DotProduct(p1, p2)
	return rv
}

func DeleteQPoint(this *QPoint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QPointD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
		qtrt.KeepMe()
	}
}

//  keep block end
