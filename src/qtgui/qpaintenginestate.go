package qtgui

// /usr/include/qt/QtGui/qpaintengine.h
// #include <qpaintengine.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 36
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
type QPaintEngineState struct {
	*qtrt.CObject
}

func (this *QPaintEngineState) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPaintEngineState) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQPaintEngineStateFromPointer(cthis unsafe.Pointer) *QPaintEngineState {
	return &QPaintEngineState{&qtrt.CObject{cthis}}
}
func (*QPaintEngineState) NewFromPointer(cthis unsafe.Pointer) *QPaintEngineState {
	return NewQPaintEngineStateFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpaintengine.h:268
// index:0
// Public inline
// QPaintEngine::DirtyFlags state()
func (this *QPaintEngineState) State() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPaintEngineState5stateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:270
// index:0
// Public
// QPen pen()
func (this *QPaintEngineState) Pen() *QPen /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPaintEngineState3penEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPenFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:271
// index:0
// Public
// QBrush brush()
func (this *QPaintEngineState) Brush() *QBrush /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPaintEngineState5brushEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:272
// index:0
// Public
// QPointF brushOrigin()
func (this *QPaintEngineState) BrushOrigin() *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPaintEngineState11brushOriginEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:273
// index:0
// Public
// QBrush backgroundBrush()
func (this *QPaintEngineState) BackgroundBrush() *QBrush /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPaintEngineState15backgroundBrushEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:274
// index:0
// Public
// Qt::BGMode backgroundMode()
func (this *QPaintEngineState) BackgroundMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPaintEngineState14backgroundModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:275
// index:0
// Public
// QFont font()
func (this *QPaintEngineState) Font() *QFont /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPaintEngineState4fontEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:276
// index:0
// Public
// QMatrix matrix()
func (this *QPaintEngineState) Matrix() *QMatrix /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPaintEngineState6matrixEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:277
// index:0
// Public
// QTransform transform()
func (this *QPaintEngineState) Transform() *QTransform /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPaintEngineState9transformEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:279
// index:0
// Public
// Qt::ClipOperation clipOperation()
func (this *QPaintEngineState) ClipOperation() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPaintEngineState13clipOperationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:280
// index:0
// Public
// QRegion clipRegion()
func (this *QPaintEngineState) ClipRegion() *QRegion /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPaintEngineState10clipRegionEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:281
// index:0
// Public
// QPainterPath clipPath()
func (this *QPaintEngineState) ClipPath() *QPainterPath /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPaintEngineState8clipPathEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:282
// index:0
// Public
// bool isClipEnabled()
func (this *QPaintEngineState) IsClipEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPaintEngineState13isClipEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpaintengine.h:284
// index:0
// Public
// QPainter::RenderHints renderHints()
func (this *QPaintEngineState) RenderHints() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPaintEngineState11renderHintsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:285
// index:0
// Public
// QPainter::CompositionMode compositionMode()
func (this *QPaintEngineState) CompositionMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPaintEngineState15compositionModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:286
// index:0
// Public
// qreal opacity()
func (this *QPaintEngineState) Opacity() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPaintEngineState7opacityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qpaintengine.h:288
// index:0
// Public
// QPainter * painter()
func (this *QPaintEngineState) Painter() *QPainter /*777 QPainter **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPaintEngineState7painterEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPainterFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:290
// index:0
// Public
// bool brushNeedsResolving()
func (this *QPaintEngineState) BrushNeedsResolving() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPaintEngineState19brushNeedsResolvingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpaintengine.h:291
// index:0
// Public
// bool penNeedsResolving()
func (this *QPaintEngineState) PenNeedsResolving() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPaintEngineState17penNeedsResolvingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
