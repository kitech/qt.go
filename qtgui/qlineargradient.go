package qtgui

// /usr/include/qt/QtGui/qbrush.h
// #include <qbrush.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QLinearGradient struct {
	*QGradient
}

func (this *QLinearGradient) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGradient.GetCthis()
	}
}
func (this *QLinearGradient) SetCthis(cthis unsafe.Pointer) {
	this.QGradient = NewQGradientFromPointer(cthis)
}
func NewQLinearGradientFromPointer(cthis unsafe.Pointer) *QLinearGradient {
	bcthis0 := NewQGradientFromPointer(cthis)
	return &QLinearGradient{bcthis0}
}
func (*QLinearGradient) NewFromPointer(cthis unsafe.Pointer) *QLinearGradient {
	return NewQLinearGradientFromPointer(cthis)
}

// /usr/include/qt/QtGui/qbrush.h:257
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLinearGradient()
func NewQLinearGradient() *QLinearGradient {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QLinearGradientC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLinearGradientFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLinearGradient)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:258
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QLinearGradient(const QPointF &, const QPointF &)
func NewQLinearGradient_1(start *qtcore.QPointF, finalStop *qtcore.QPointF) *QLinearGradient {
	var convArg0 = start.GetCthis()
	var convArg1 = finalStop.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QLinearGradientC2ERK7QPointFS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLinearGradientFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLinearGradient)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:259
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QLinearGradient(qreal, qreal, qreal, qreal)
func NewQLinearGradient_2(xStart float64, yStart float64, xFinalStop float64, yFinalStop float64) *QLinearGradient {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QLinearGradientC2Edddd", qtrt.FFI_TYPE_POINTER, xStart, yStart, xFinalStop, yFinalStop)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLinearGradientFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLinearGradient)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:261
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF start()
func (this *QLinearGradient) Start() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QLinearGradient5startEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qbrush.h:262
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStart(const QPointF &)
func (this *QLinearGradient) SetStart(start *qtcore.QPointF) {
	var convArg0 = start.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QLinearGradient8setStartERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:263
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setStart(qreal, qreal)
func (this *QLinearGradient) SetStart_1(x float64, y float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QLinearGradient8setStartEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:265
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF finalStop()
func (this *QLinearGradient) FinalStop() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QLinearGradient9finalStopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qbrush.h:266
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFinalStop(const QPointF &)
func (this *QLinearGradient) SetFinalStop(stop *qtcore.QPointF) {
	var convArg0 = stop.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QLinearGradient12setFinalStopERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:267
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setFinalStop(qreal, qreal)
func (this *QLinearGradient) SetFinalStop_1(x float64, y float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QLinearGradient12setFinalStopEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
}

func DeleteQLinearGradient(this *QLinearGradient) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QLinearGradientD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
