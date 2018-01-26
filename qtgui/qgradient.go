package qtgui

// /usr/include/qt/QtGui/qbrush.h
// #include <qbrush.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
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
	*qtrt.CObject
}

func (this *QGradient) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QGradient) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQGradientFromPointer(cthis unsafe.Pointer) *QGradient {
	return &QGradient{&qtrt.CObject{cthis}}
}
func (*QGradient) NewFromPointer(cthis unsafe.Pointer) *QGradient {
	return NewQGradientFromPointer(cthis)
}

// /usr/include/qt/QtGui/qbrush.h:206
// index:0
// Public
// void QGradient()
func NewQGradient() *QGradient {
	cthis := qtrt.Calloc(1, 256) // 64
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGradientC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQGradientFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:208
// index:0
// Public inline
// QGradient::Type type()
func (this *QGradient) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGradient4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qbrush.h:210
// index:0
// Public inline
// void setSpread(enum QGradient::Spread)
func (this *QGradient) SetSpread(spread int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGradient9setSpreadENS_6SpreadE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), spread)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:211
// index:0
// Public inline
// QGradient::Spread spread()
func (this *QGradient) Spread() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGradient6spreadEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qbrush.h:213
// index:0
// Public
// void setColorAt(qreal, const class QColor &)
func (this *QGradient) SetColorAt(pos float64, color *QColor) {
	var convArg1 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGradient10setColorAtEdRK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pos, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:218
// index:0
// Public
// QGradient::CoordinateMode coordinateMode()
func (this *QGradient) CoordinateMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGradient14coordinateModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qbrush.h:219
// index:0
// Public
// void setCoordinateMode(enum QGradient::CoordinateMode)
func (this *QGradient) SetCoordinateMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGradient17setCoordinateModeENS_14CoordinateModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:221
// index:0
// Public
// QGradient::InterpolationMode interpolationMode()
func (this *QGradient) InterpolationMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGradient17interpolationModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qbrush.h:222
// index:0
// Public
// void setInterpolationMode(enum QGradient::InterpolationMode)
func (this *QGradient) SetInterpolationMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGradient20setInterpolationModeENS_17InterpolationModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

type QGradient__Type = int

const QGradient__LinearGradient QGradient__Type = 0
const QGradient__RadialGradient QGradient__Type = 1
const QGradient__ConicalGradient QGradient__Type = 2
const QGradient__NoGradient QGradient__Type = 3

type QGradient__Spread = int

const QGradient__PadSpread QGradient__Spread = 0
const QGradient__ReflectSpread QGradient__Spread = 1
const QGradient__RepeatSpread QGradient__Spread = 2

type QGradient__CoordinateMode = int

const QGradient__LogicalMode QGradient__CoordinateMode = 0
const QGradient__StretchToDeviceMode QGradient__CoordinateMode = 1
const QGradient__ObjectBoundingMode QGradient__CoordinateMode = 2

type QGradient__InterpolationMode = int

const QGradient__ColorInterpolation QGradient__InterpolationMode = 0
const QGradient__ComponentInterpolation QGradient__InterpolationMode = 1

//  body block end
