package qtcore

// /usr/include/qt/QtCore/qpoint.h
// #include <qpoint.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
type QPointF struct {
	*qtrt.CObject
}

func (this *QPointF) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPointF) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQPointFFromPointer(cthis unsafe.Pointer) *QPointF {
	return &QPointF{&qtrt.CObject{cthis}}
}
func (*QPointF) NewFromPointer(cthis unsafe.Pointer) *QPointF {
	return NewQPointFFromPointer(cthis)
}

// /usr/include/qt/QtCore/qpoint.h:222
// index:0
// Public inline
// void QPointF()
func NewQPointF() *QPointF {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPointFC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQPointFFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qpoint.h:223
// index:1
// Public inline
// void QPointF(const QPoint &)
func NewQPointF_1(p *QPoint) *QPointF {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPointFC2ERK6QPoint", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPointFFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qpoint.h:224
// index:2
// Public inline
// void QPointF(qreal, qreal)
func NewQPointF_2(xpos float64, ypos float64) *QPointF {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPointFC2Edd", ffiqt.FFI_TYPE_VOID, cthis, xpos, ypos)
	gopp.ErrPrint(err, rv)
	gothis := NewQPointFFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qpoint.h:226
// index:0
// Public inline
// qreal manhattanLength()
func (this *QPointF) ManhattanLength() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPointF15manhattanLengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtCore/qpoint.h:228
// index:0
// Public inline
// bool isNull()
func (this *QPointF) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPointF6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qpoint.h:230
// index:0
// Public inline
// qreal x()
func (this *QPointF) X() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPointF1xEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtCore/qpoint.h:231
// index:0
// Public inline
// qreal y()
func (this *QPointF) Y() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPointF1yEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtCore/qpoint.h:232
// index:0
// Public inline
// void setX(qreal)
func (this *QPointF) SetX(x float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPointF4setXEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpoint.h:233
// index:0
// Public inline
// void setY(qreal)
func (this *QPointF) SetY(y float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPointF4setYEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpoint.h:235
// index:0
// Public inline
// qreal & rx()
func (this *QPointF) Rx() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPointF2rxEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 3331
}

// /usr/include/qt/QtCore/qpoint.h:236
// index:0
// Public inline
// qreal & ry()
func (this *QPointF) Ry() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPointF2ryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 3331
}

// /usr/include/qt/QtCore/qpoint.h:243
// index:0
// Public static inline
// qreal dotProduct(const QPointF &, const QPointF &)
func (this *QPointF) DotProduct(p1 *QPointF, p2 *QPointF) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPointF10dotProductERKS_S1_", ffiqt.FFI_TYPE_POINTER, p1, p2)
	gopp.ErrPrint(err, rv)
	// return rv
	return float64(rv) // 222
}
func QPointF_DotProduct(p1 *QPointF, p2 *QPointF) float64 {
	var nilthis *QPointF
	rv := nilthis.DotProduct(p1, p2)
	return rv
}

// /usr/include/qt/QtCore/qpoint.h:256
// index:0
// Public inline
// QPoint toPoint()
func (this *QPointF) ToPoint() *QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPointF7toPointEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
