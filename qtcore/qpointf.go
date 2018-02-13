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
// extern C begin: 11
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QPointF struct {
	*qtrt.CObject
}
type QPointF_ITF interface {
	QPointF_PTR() *QPointF
}

func (ptr *QPointF) QPointF_PTR() *QPointF { return ptr }

func (this *QPointF) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPointF) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQPointFFromPointer(cthis unsafe.Pointer) *QPointF {
	return &QPointF{&qtrt.CObject{cthis}}
}
func (*QPointF) NewFromPointer(cthis unsafe.Pointer) *QPointF {
	return NewQPointFFromPointer(cthis)
}

// /usr/include/qt/QtCore/qpoint.h:222
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QPointF()
func NewQPointF() *QPointF {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPointFC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPointFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPointF)
	return gothis
}

// /usr/include/qt/QtCore/qpoint.h:223
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QPointF(const QPoint &)
func NewQPointF_1(p *QPoint) *QPointF {
	var convArg0 = p.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPointFC2ERK6QPoint", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPointFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPointF)
	return gothis
}

// /usr/include/qt/QtCore/qpoint.h:224
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QPointF(qreal, qreal)
func NewQPointF_2(xpos float64, ypos float64) *QPointF {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPointFC2Edd", qtrt.FFI_TYPE_POINTER, xpos, ypos)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPointFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPointF)
	return gothis
}

// /usr/include/qt/QtCore/qpoint.h:226
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal manhattanLength()
func (this *QPointF) ManhattanLength() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPointF15manhattanLengthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qpoint.h:228
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QPointF) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPointF6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qpoint.h:230
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal x()
func (this *QPointF) X() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPointF1xEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qpoint.h:231
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal y()
func (this *QPointF) Y() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPointF1yEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qpoint.h:232
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setX(qreal)
func (this *QPointF) SetX(x float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPointF4setXEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpoint.h:233
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setY(qreal)
func (this *QPointF) SetY(y float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPointF4setYEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpoint.h:235
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal & rx()
func (this *QPointF) Rx() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPointF2rxEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cpretval2go("float64", rv).(float64) // 3331
}

// /usr/include/qt/QtCore/qpoint.h:236
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal & ry()
func (this *QPointF) Ry() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPointF2ryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cpretval2go("float64", rv).(float64) // 3331
}

// /usr/include/qt/QtCore/qpoint.h:243
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] qreal dotProduct(const QPointF &, const QPointF &)
func (this *QPointF) DotProduct(p1 *QPointF, p2 *QPointF) float64 {
	var convArg0 = p1.GetCthis()
	var convArg1 = p2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPointF10dotProductERKS_S1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}
func QPointF_DotProduct(p1 *QPointF, p2 *QPointF) float64 {
	var nilthis *QPointF
	rv := nilthis.DotProduct(p1, p2)
	return rv
}

// /usr/include/qt/QtCore/qpoint.h:256
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint toPoint()
func (this *QPointF) ToPoint() *QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPointF7toPointEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPoint)
	return rv2
}

func DeleteQPointF(this *QPointF) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPointFD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
