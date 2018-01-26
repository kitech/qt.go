package qtgui

// /usr/include/qt/QtGui/qpainterpath.h
// #include <qpainterpath.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 65
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
type QPainterPathStroker struct {
	*qtrt.CObject
}

func (this *QPainterPathStroker) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPainterPathStroker) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQPainterPathStrokerFromPointer(cthis unsafe.Pointer) *QPainterPathStroker {
	return &QPainterPathStroker{&qtrt.CObject{cthis}}
}
func (*QPainterPathStroker) NewFromPointer(cthis unsafe.Pointer) *QPainterPathStroker {
	return NewQPainterPathStrokerFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpainterpath.h:246
// index:0
// Public
// void QPainterPathStroker()
func NewQPainterPathStroker() *QPainterPathStroker {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QPainterPathStrokerC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQPainterPathStrokerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpainterpath.h:247
// index:1
// Public
// void QPainterPathStroker(const class QPen &)
func NewQPainterPathStroker_1(pen *QPen) *QPainterPathStroker {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = pen.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QPainterPathStrokerC2ERK4QPen", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPainterPathStrokerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpainterpath.h:248
// index:0
// Public
// void ~QPainterPathStroker()
func DeleteQPainterPathStroker(*QPainterPathStroker) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QPainterPathStrokerD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:250
// index:0
// Public
// void setWidth(qreal)
func (this *QPainterPathStroker) SetWidth(width float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QPainterPathStroker8setWidthEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:251
// index:0
// Public
// qreal width()
func (this *QPainterPathStroker) Width() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QPainterPathStroker5widthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qpainterpath.h:253
// index:0
// Public
// void setCapStyle(Qt::PenCapStyle)
func (this *QPainterPathStroker) SetCapStyle(style int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QPainterPathStroker11setCapStyleEN2Qt11PenCapStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:254
// index:0
// Public
// Qt::PenCapStyle capStyle()
func (this *QPainterPathStroker) CapStyle() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QPainterPathStroker8capStyleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:256
// index:0
// Public
// void setJoinStyle(Qt::PenJoinStyle)
func (this *QPainterPathStroker) SetJoinStyle(style int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QPainterPathStroker12setJoinStyleEN2Qt12PenJoinStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:257
// index:0
// Public
// Qt::PenJoinStyle joinStyle()
func (this *QPainterPathStroker) JoinStyle() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QPainterPathStroker9joinStyleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:259
// index:0
// Public
// void setMiterLimit(qreal)
func (this *QPainterPathStroker) SetMiterLimit(length float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QPainterPathStroker13setMiterLimitEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), length)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:260
// index:0
// Public
// qreal miterLimit()
func (this *QPainterPathStroker) MiterLimit() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QPainterPathStroker10miterLimitEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qpainterpath.h:262
// index:0
// Public
// void setCurveThreshold(qreal)
func (this *QPainterPathStroker) SetCurveThreshold(threshold float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QPainterPathStroker17setCurveThresholdEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), threshold)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:263
// index:0
// Public
// qreal curveThreshold()
func (this *QPainterPathStroker) CurveThreshold() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QPainterPathStroker14curveThresholdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qpainterpath.h:265
// index:0
// Public
// void setDashPattern(Qt::PenStyle)
func (this *QPainterPathStroker) SetDashPattern(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QPainterPathStroker14setDashPatternEN2Qt8PenStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:269
// index:0
// Public
// void setDashOffset(qreal)
func (this *QPainterPathStroker) SetDashOffset(offset float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QPainterPathStroker13setDashOffsetEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), offset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:270
// index:0
// Public
// qreal dashOffset()
func (this *QPainterPathStroker) DashOffset() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QPainterPathStroker10dashOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qpainterpath.h:272
// index:0
// Public
// QPainterPath createStroke(const class QPainterPath &)
func (this *QPainterPathStroker) CreateStroke(path *QPainterPath) *QPainterPath /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QPainterPathStroker12createStrokeERK12QPainterPath", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
