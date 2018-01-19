//  header block begin
// /usr/include/qt/QtGui/qbrush.h
// #include <qbrush.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
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
type QGradient struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qbrush.h:206
// index:0
// void QGradient()
func NewQGradient() *QGradient {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGradientC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QGradient{cthis}
}

// /usr/include/qt/QtGui/qbrush.h:208
// index:0
// inline
// QGradient::Type type()
func (this *QGradient) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGradient4typeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:210
// index:0
// inline
// void setSpread(enum QGradient::Spread)
func (this *QGradient) SetSpread(spread int) {
	// 0: (, QGradient::Spread spread), (&spread)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGradient9setSpreadENS_6SpreadE", ffiqt.FFI_TYPE_VOID, this.cthis, &spread)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:211
// index:0
// inline
// QGradient::Spread spread()
func (this *QGradient) Spread() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGradient6spreadEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:213
// index:0
// void setColorAt(qreal, const class QColor &)
func (this *QGradient) SetColorAt(pos float64, color unsafe.Pointer) {
	// 0: (, qreal pos, const QColor & color), (&pos, color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGradient10setColorAtEdRK6QColor", ffiqt.FFI_TYPE_VOID, this.cthis, &pos, color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:216
// index:0
// QGradientStops stops()
func (this *QGradient) Stops() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGradient5stopsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:218
// index:0
// QGradient::CoordinateMode coordinateMode()
func (this *QGradient) CoordinateMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGradient14coordinateModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:219
// index:0
// void setCoordinateMode(enum QGradient::CoordinateMode)
func (this *QGradient) SetCoordinateMode(mode int) {
	// 0: (, QGradient::CoordinateMode mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGradient17setCoordinateModeENS_14CoordinateModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:221
// index:0
// QGradient::InterpolationMode interpolationMode()
func (this *QGradient) InterpolationMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGradient17interpolationModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:222
// index:0
// void setInterpolationMode(enum QGradient::InterpolationMode)
func (this *QGradient) SetInterpolationMode(mode int) {
	// 0: (, QGradient::InterpolationMode mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGradient20setInterpolationModeENS_17InterpolationModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &mode)
	gopp.ErrPrint(err, rv)
}

//  body block end
