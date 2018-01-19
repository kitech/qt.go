//  header block begin
// /usr/include/qt/QtGui/qpainterpath.h
// #include <qpainterpath.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 69
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qpainterpath.h:246
// index:0
// void QPainterPathStroker()
func NewQPainterPathStroker() *QPainterPathStroker {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QPainterPathStrokerC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QPainterPathStroker{cthis}
}

// /usr/include/qt/QtGui/qpainterpath.h:247
// index:1
// void QPainterPathStroker(const class QPen &)
func NewQPainterPathStroker_1(pen unsafe.Pointer) *QPainterPathStroker {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QPainterPathStrokerC2ERK4QPen", ffiqt.FFI_TYPE_VOID, cthis, pen)
	gopp.ErrPrint(err, rv)
	return &QPainterPathStroker{cthis}
}

// /usr/include/qt/QtGui/qpainterpath.h:248
// index:0
// void ~QPainterPathStroker()
func DeleteQPainterPathStroker(*QPainterPathStroker) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QPainterPathStrokerD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:250
// index:0
// void setWidth(qreal)
func (this *QPainterPathStroker) SetWidth(width float64) {
	// 0: (, qreal width), (&width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QPainterPathStroker8setWidthEd", ffiqt.FFI_TYPE_VOID, this.cthis, &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:251
// index:0
// qreal width()
func (this *QPainterPathStroker) Width() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QPainterPathStroker5widthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:253
// index:0
// void setCapStyle(Qt::PenCapStyle)
func (this *QPainterPathStroker) SetCapStyle(style int) {
	// 0: (, Qt::PenCapStyle style), (&style)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QPainterPathStroker11setCapStyleEN2Qt11PenCapStyleE", ffiqt.FFI_TYPE_VOID, this.cthis, &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:254
// index:0
// Qt::PenCapStyle capStyle()
func (this *QPainterPathStroker) CapStyle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QPainterPathStroker8capStyleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:256
// index:0
// void setJoinStyle(Qt::PenJoinStyle)
func (this *QPainterPathStroker) SetJoinStyle(style int) {
	// 0: (, Qt::PenJoinStyle style), (&style)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QPainterPathStroker12setJoinStyleEN2Qt12PenJoinStyleE", ffiqt.FFI_TYPE_VOID, this.cthis, &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:257
// index:0
// Qt::PenJoinStyle joinStyle()
func (this *QPainterPathStroker) JoinStyle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QPainterPathStroker9joinStyleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:259
// index:0
// void setMiterLimit(qreal)
func (this *QPainterPathStroker) SetMiterLimit(length float64) {
	// 0: (, qreal length), (&length)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QPainterPathStroker13setMiterLimitEd", ffiqt.FFI_TYPE_VOID, this.cthis, &length)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:260
// index:0
// qreal miterLimit()
func (this *QPainterPathStroker) MiterLimit() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QPainterPathStroker10miterLimitEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:262
// index:0
// void setCurveThreshold(qreal)
func (this *QPainterPathStroker) SetCurveThreshold(threshold float64) {
	// 0: (, qreal threshold), (&threshold)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QPainterPathStroker17setCurveThresholdEd", ffiqt.FFI_TYPE_VOID, this.cthis, &threshold)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:263
// index:0
// qreal curveThreshold()
func (this *QPainterPathStroker) CurveThreshold() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QPainterPathStroker14curveThresholdEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:265
// index:0
// void setDashPattern(Qt::PenStyle)
func (this *QPainterPathStroker) SetDashPattern(arg0 int) {
	// 0: (, Qt::PenStyle arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QPainterPathStroker14setDashPatternEN2Qt8PenStyleE", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:267
// index:0
// QVector<qreal> dashPattern()
func (this *QPainterPathStroker) DashPattern() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QPainterPathStroker11dashPatternEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:269
// index:0
// void setDashOffset(qreal)
func (this *QPainterPathStroker) SetDashOffset(offset float64) {
	// 0: (, qreal offset), (&offset)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QPainterPathStroker13setDashOffsetEd", ffiqt.FFI_TYPE_VOID, this.cthis, &offset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:270
// index:0
// qreal dashOffset()
func (this *QPainterPathStroker) DashOffset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QPainterPathStroker10dashOffsetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:272
// index:0
// QPainterPath createStroke(const class QPainterPath &)
func (this *QPainterPathStroker) CreateStroke(path unsafe.Pointer) {
	// 0: (, const QPainterPath & path), (path)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QPainterPathStroker12createStrokeERK12QPainterPath", ffiqt.FFI_TYPE_VOID, this.cthis, path)
	gopp.ErrPrint(err, rv)
}

//  body block end
