//  header block begin
// /usr/include/qt/QtGui/qbrush.h
// #include <qbrush.h>
// #include <QtGui>
package qtgui

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
type QConicalGradient struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qbrush.h:306
// index:0
// void QConicalGradient()
func NewQConicalGradient() *QConicalGradient {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QConicalGradientC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QConicalGradient{cthis}
}

// /usr/include/qt/QtGui/qbrush.h:307
// index:1
// void QConicalGradient(const class QPointF &, qreal)
func NewQConicalGradient_1(center unsafe.Pointer, startAngle float64) *QConicalGradient {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QConicalGradientC2ERK7QPointFd", ffiqt.FFI_TYPE_VOID, cthis, center, &startAngle)
	gopp.ErrPrint(err, rv)
	return &QConicalGradient{cthis}
}

// /usr/include/qt/QtGui/qbrush.h:308
// index:2
// void QConicalGradient(qreal, qreal, qreal)
func NewQConicalGradient_2(cx float64, cy float64, startAngle float64) *QConicalGradient {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QConicalGradientC2Eddd", ffiqt.FFI_TYPE_VOID, cthis, &cx, &cy, &startAngle)
	gopp.ErrPrint(err, rv)
	return &QConicalGradient{cthis}
}

// /usr/include/qt/QtGui/qbrush.h:310
// index:0
// QPointF center()
func (this *QConicalGradient) Center() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QConicalGradient6centerEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:311
// index:0
// void setCenter(const class QPointF &)
func (this *QConicalGradient) SetCenter(center unsafe.Pointer) {
	// 0: (, const QPointF & center), (center)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QConicalGradient9setCenterERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, center)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:312
// index:1
// inline
// void setCenter(qreal, qreal)
func (this *QConicalGradient) SetCenter_1(x float64, y float64) {
	// 1: (, qreal x, qreal y), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QConicalGradient9setCenterEdd", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:314
// index:0
// qreal angle()
func (this *QConicalGradient) Angle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QConicalGradient5angleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:315
// index:0
// void setAngle(qreal)
func (this *QConicalGradient) SetAngle(angle float64) {
	// 0: (, qreal angle), (&angle)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QConicalGradient8setAngleEd", ffiqt.FFI_TYPE_VOID, this.cthis, &angle)
	gopp.ErrPrint(err, rv)
}

//  body block end
