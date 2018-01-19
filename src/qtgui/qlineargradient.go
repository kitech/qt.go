//  header block begin
// /usr/include/qt/QtGui/qbrush.h
// #include <qbrush.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qbrush.h:257
// index:0
// void QLinearGradient()
func NewQLinearGradient() *QLinearGradient {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QLinearGradientC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QLinearGradient{cthis}
}

// /usr/include/qt/QtGui/qbrush.h:258
// index:1
// void QLinearGradient(const class QPointF &, const class QPointF &)
func NewQLinearGradient_1(start unsafe.Pointer, finalStop unsafe.Pointer) *QLinearGradient {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QLinearGradientC2ERK7QPointFS2_", ffiqt.FFI_TYPE_VOID, cthis, start, finalStop)
	gopp.ErrPrint(err, rv)
	return &QLinearGradient{cthis}
}

// /usr/include/qt/QtGui/qbrush.h:259
// index:2
// void QLinearGradient(qreal, qreal, qreal, qreal)
func NewQLinearGradient_2(xStart float64, yStart float64, xFinalStop float64, yFinalStop float64) *QLinearGradient {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QLinearGradientC2Edddd", ffiqt.FFI_TYPE_VOID, cthis, &xStart, &yStart, &xFinalStop, &yFinalStop)
	gopp.ErrPrint(err, rv)
	return &QLinearGradient{cthis}
}

// /usr/include/qt/QtGui/qbrush.h:261
// index:0
// QPointF start()
func (this *QLinearGradient) Start() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QLinearGradient5startEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:262
// index:0
// void setStart(const class QPointF &)
func (this *QLinearGradient) SetStart(start unsafe.Pointer) {
	// 0: (, const QPointF & start), (start)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QLinearGradient8setStartERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, start)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:263
// index:1
// inline
// void setStart(qreal, qreal)
func (this *QLinearGradient) SetStart_1(x float64, y float64) {
	// 1: (, qreal x, qreal y), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QLinearGradient8setStartEdd", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:265
// index:0
// QPointF finalStop()
func (this *QLinearGradient) FinalStop() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QLinearGradient9finalStopEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:266
// index:0
// void setFinalStop(const class QPointF &)
func (this *QLinearGradient) SetFinalStop(stop unsafe.Pointer) {
	// 0: (, const QPointF & stop), (stop)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QLinearGradient12setFinalStopERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, stop)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:267
// index:1
// inline
// void setFinalStop(qreal, qreal)
func (this *QLinearGradient) SetFinalStop_1(x float64, y float64) {
	// 1: (, qreal x, qreal y), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QLinearGradient12setFinalStopEdd", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y)
	gopp.ErrPrint(err, rv)
}

//  body block end
