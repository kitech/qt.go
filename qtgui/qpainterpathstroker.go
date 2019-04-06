package qtgui

// /usr/include/qt/QtGui/qpainterpath.h
// #include <qpainterpath.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 77
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QPainterPathStroker struct {
	*qtrt.CObject
}
type QPainterPathStroker_ITF interface {
	QPainterPathStroker_PTR() *QPainterPathStroker
}

func (ptr *QPainterPathStroker) QPainterPathStroker_PTR() *QPainterPathStroker { return ptr }

func (this *QPainterPathStroker) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPainterPathStroker) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQPainterPathStrokerFromPointer(cthis unsafe.Pointer) *QPainterPathStroker {
	return &QPainterPathStroker{&qtrt.CObject{cthis}}
}
func (*QPainterPathStroker) NewFromPointer(cthis unsafe.Pointer) *QPainterPathStroker {
	return NewQPainterPathStrokerFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpainterpath.h:246
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPainterPathStroker()

/*

 */
func (*QPainterPathStroker) NewForInherit() *QPainterPathStroker {
	return NewQPainterPathStroker()
}
func NewQPainterPathStroker() *QPainterPathStroker {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QPainterPathStrokerC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPainterPathStrokerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPainterPathStroker)
	return gothis
}

// /usr/include/qt/QtGui/qpainterpath.h:247
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPainterPathStroker(const QPen &)

/*

 */
func (*QPainterPathStroker) NewForInherit1(pen QPen_ITF) *QPainterPathStroker {
	return NewQPainterPathStroker1(pen)
}
func NewQPainterPathStroker1(pen QPen_ITF) *QPainterPathStroker {
	var convArg0 unsafe.Pointer
	if pen != nil && pen.QPen_PTR() != nil {
		convArg0 = pen.QPen_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QPainterPathStrokerC2ERK4QPen", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPainterPathStrokerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPainterPathStroker)
	return gothis
}

// /usr/include/qt/QtGui/qpainterpath.h:248
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QPainterPathStroker()

/*

 */
func DeleteQPainterPathStroker(this *QPainterPathStroker) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QPainterPathStrokerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpainterpath.h:250
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidth(qreal)

/*

 */
func (this *QPainterPathStroker) SetWidth(width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QPainterPathStroker8setWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:251
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal width() const

/*

 */
func (this *QPainterPathStroker) Width() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QPainterPathStroker5widthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpainterpath.h:253
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCapStyle(Qt::PenCapStyle)

/*

 */
func (this *QPainterPathStroker) SetCapStyle(style int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QPainterPathStroker11setCapStyleEN2Qt11PenCapStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), style)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:254
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::PenCapStyle capStyle() const

/*

 */
func (this *QPainterPathStroker) CapStyle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QPainterPathStroker8capStyleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:256
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setJoinStyle(Qt::PenJoinStyle)

/*

 */
func (this *QPainterPathStroker) SetJoinStyle(style int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QPainterPathStroker12setJoinStyleEN2Qt12PenJoinStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), style)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:257
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::PenJoinStyle joinStyle() const

/*

 */
func (this *QPainterPathStroker) JoinStyle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QPainterPathStroker9joinStyleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:259
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMiterLimit(qreal)

/*

 */
func (this *QPainterPathStroker) SetMiterLimit(length float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QPainterPathStroker13setMiterLimitEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), length)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:260
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal miterLimit() const

/*

 */
func (this *QPainterPathStroker) MiterLimit() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QPainterPathStroker10miterLimitEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpainterpath.h:262
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurveThreshold(qreal)

/*

 */
func (this *QPainterPathStroker) SetCurveThreshold(threshold float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QPainterPathStroker17setCurveThresholdEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), threshold)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:263
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal curveThreshold() const

/*

 */
func (this *QPainterPathStroker) CurveThreshold() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QPainterPathStroker14curveThresholdEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpainterpath.h:265
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDashPattern(Qt::PenStyle)

/*

 */
func (this *QPainterPathStroker) SetDashPattern(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QPainterPathStroker14setDashPatternEN2Qt8PenStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:269
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDashOffset(qreal)

/*

 */
func (this *QPainterPathStroker) SetDashOffset(offset float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QPainterPathStroker13setDashOffsetEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), offset)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:270
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal dashOffset() const

/*

 */
func (this *QPainterPathStroker) DashOffset() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QPainterPathStroker10dashOffsetEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpainterpath.h:272
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath createStroke(const QPainterPath &) const

/*

 */
func (this *QPainterPathStroker) CreateStroke(path QPainterPath_ITF) *QPainterPath /*123*/ {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QPainterPathStroker12createStrokeERK12QPainterPath", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

//  body block end

//  keep block begin

func init_unused_10731() {
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
