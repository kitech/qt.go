package qtgui

// /usr/include/qt/QtGui/qbrush.h
// #include <qbrush.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QLinearGradientC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QLinearGradientC2ERK7QPointFS2_", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQLinearGradientFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLinearGradient)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:259
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QLinearGradient(qreal, qreal, qreal, qreal)
func NewQLinearGradient_2(xStart float64, yStart float64, xFinalStop float64, yFinalStop float64) *QLinearGradient {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QLinearGradientC2Edddd", ffiqt.FFI_TYPE_POINTER, xStart, yStart, xFinalStop, yFinalStop)
	gopp.ErrPrint(err, rv)
	gothis := NewQLinearGradientFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLinearGradient)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:261
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF start()
func (this *QLinearGradient) Start() *qtcore.QPointF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QLinearGradient5startEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QLinearGradient8setStartERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:263
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setStart(qreal, qreal)
func (this *QLinearGradient) SetStart_1(x float64, y float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QLinearGradient8setStartEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:265
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF finalStop()
func (this *QLinearGradient) FinalStop() *qtcore.QPointF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QLinearGradient9finalStopEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QLinearGradient12setFinalStopERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:267
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setFinalStop(qreal, qreal)
func (this *QLinearGradient) SetFinalStop_1(x float64, y float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QLinearGradient12setFinalStopEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
}

func DeleteQLinearGradient(this *QLinearGradient) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QLinearGradientD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
