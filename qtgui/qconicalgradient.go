package qtgui

// /usr/include/qt/QtGui/qbrush.h
// #include <qbrush.h>
// #include <QtGui>

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
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QConicalGradient struct {
	*QGradient
}

func (this *QConicalGradient) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGradient.GetCthis()
	}
}
func (this *QConicalGradient) SetCthis(cthis unsafe.Pointer) {
	this.QGradient = NewQGradientFromPointer(cthis)
}
func NewQConicalGradientFromPointer(cthis unsafe.Pointer) *QConicalGradient {
	bcthis0 := NewQGradientFromPointer(cthis)
	return &QConicalGradient{bcthis0}
}
func (*QConicalGradient) NewFromPointer(cthis unsafe.Pointer) *QConicalGradient {
	return NewQConicalGradientFromPointer(cthis)
}

// /usr/include/qt/QtGui/qbrush.h:306
// index:0
// Public
// void QConicalGradient()
func NewQConicalGradient() *QConicalGradient {
	cthis := qtrt.Calloc(1, 256) // 64
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QConicalGradientC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQConicalGradientFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:307
// index:1
// Public
// void QConicalGradient(const QPointF &, qreal)
func NewQConicalGradient_1(center *qtcore.QPointF, startAngle float64) *QConicalGradient {
	cthis := qtrt.Calloc(1, 256) // 64
	var convArg0 = center.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QConicalGradientC2ERK7QPointFd", ffiqt.FFI_TYPE_VOID, cthis, convArg0, startAngle)
	gopp.ErrPrint(err, rv)
	gothis := NewQConicalGradientFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:308
// index:2
// Public
// void QConicalGradient(qreal, qreal, qreal)
func NewQConicalGradient_2(cx float64, cy float64, startAngle float64) *QConicalGradient {
	cthis := qtrt.Calloc(1, 256) // 64
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QConicalGradientC2Eddd", ffiqt.FFI_TYPE_VOID, cthis, cx, cy, startAngle)
	gopp.ErrPrint(err, rv)
	gothis := NewQConicalGradientFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:310
// index:0
// Public
// QPointF center()
func (this *QConicalGradient) Center() *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QConicalGradient6centerEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qbrush.h:311
// index:0
// Public
// void setCenter(const QPointF &)
func (this *QConicalGradient) SetCenter(center *qtcore.QPointF) {
	var convArg0 = center.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QConicalGradient9setCenterERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:312
// index:1
// Public inline
// void setCenter(qreal, qreal)
func (this *QConicalGradient) SetCenter_1(x float64, y float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QConicalGradient9setCenterEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:314
// index:0
// Public
// qreal angle()
func (this *QConicalGradient) Angle() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QConicalGradient5angleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qbrush.h:315
// index:0
// Public
// void setAngle(qreal)
func (this *QConicalGradient) SetAngle(angle float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QConicalGradient8setAngleEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), angle)
	gopp.ErrPrint(err, rv)
}

//  body block end
