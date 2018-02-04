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
// Public Visibility=Default Availability=Available
// [-2] void QConicalGradient()
func NewQConicalGradient() *QConicalGradient {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QConicalGradientC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQConicalGradientFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQConicalGradient)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:307
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QConicalGradient(const QPointF &, qreal)
func NewQConicalGradient_1(center *qtcore.QPointF, startAngle float64) *QConicalGradient {
	var convArg0 = center.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QConicalGradientC2ERK7QPointFd", qtrt.FFI_TYPE_POINTER, convArg0, startAngle)
	gopp.ErrPrint(err, rv)
	gothis := NewQConicalGradientFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQConicalGradient)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:308
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QConicalGradient(qreal, qreal, qreal)
func NewQConicalGradient_2(cx float64, cy float64, startAngle float64) *QConicalGradient {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QConicalGradientC2Eddd", qtrt.FFI_TYPE_POINTER, cx, cy, startAngle)
	gopp.ErrPrint(err, rv)
	gothis := NewQConicalGradientFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQConicalGradient)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:310
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF center()
func (this *QConicalGradient) Center() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QConicalGradient6centerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qbrush.h:311
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCenter(const QPointF &)
func (this *QConicalGradient) SetCenter(center *qtcore.QPointF) {
	var convArg0 = center.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QConicalGradient9setCenterERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:312
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setCenter(qreal, qreal)
func (this *QConicalGradient) SetCenter_1(x float64, y float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QConicalGradient9setCenterEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:314
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal angle()
func (this *QConicalGradient) Angle() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QConicalGradient5angleEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qbrush.h:315
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAngle(qreal)
func (this *QConicalGradient) SetAngle(angle float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QConicalGradient8setAngleEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), angle)
	gopp.ErrPrint(err, rv)
}

func DeleteQConicalGradient(this *QConicalGradient) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QConicalGradientD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
