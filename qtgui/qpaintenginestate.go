package qtgui

// /usr/include/qt/QtGui/qpaintengine.h
// #include <qpaintengine.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 38
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QPaintEngineState struct {
	*qtrt.CObject
}
type QPaintEngineState_ITF interface {
	QPaintEngineState_PTR() *QPaintEngineState
}

func (ptr *QPaintEngineState) QPaintEngineState_PTR() *QPaintEngineState { return ptr }

func (this *QPaintEngineState) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPaintEngineState) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQPaintEngineStateFromPointer(cthis unsafe.Pointer) *QPaintEngineState {
	return &QPaintEngineState{&qtrt.CObject{cthis}}
}
func (*QPaintEngineState) NewFromPointer(cthis unsafe.Pointer) *QPaintEngineState {
	return NewQPaintEngineStateFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpaintengine.h:268
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QPaintEngine::DirtyFlags state()
func (this *QPaintEngineState) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPaintEngineState5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:270
// index:0
// Public Visibility=Default Availability=Available
// [8] QPen pen()
func (this *QPaintEngineState) Pen() *QPen /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPaintEngineState3penEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPenFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPen)
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:271
// index:0
// Public Visibility=Default Availability=Available
// [8] QBrush brush()
func (this *QPaintEngineState) Brush() *QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPaintEngineState5brushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:272
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF brushOrigin()
func (this *QPaintEngineState) BrushOrigin() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPaintEngineState11brushOriginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:273
// index:0
// Public Visibility=Default Availability=Available
// [8] QBrush backgroundBrush()
func (this *QPaintEngineState) BackgroundBrush() *QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPaintEngineState15backgroundBrushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:274
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::BGMode backgroundMode()
func (this *QPaintEngineState) BackgroundMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPaintEngineState14backgroundModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:275
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont font()
func (this *QPaintEngineState) Font() *QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPaintEngineState4fontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFont)
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:276
// index:0
// Public Visibility=Default Availability=Available
// [48] QMatrix matrix()
func (this *QPaintEngineState) Matrix() *QMatrix /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPaintEngineState6matrixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix)
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:277
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform transform()
func (this *QPaintEngineState) Transform() *QTransform /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPaintEngineState9transformEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:279
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ClipOperation clipOperation()
func (this *QPaintEngineState) ClipOperation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPaintEngineState13clipOperationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:280
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion clipRegion()
func (this *QPaintEngineState) ClipRegion() *QRegion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPaintEngineState10clipRegionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:281
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath clipPath()
func (this *QPaintEngineState) ClipPath() *QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPaintEngineState8clipPathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:282
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isClipEnabled()
func (this *QPaintEngineState) IsClipEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPaintEngineState13isClipEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpaintengine.h:284
// index:0
// Public Visibility=Default Availability=Available
// [4] QPainter::RenderHints renderHints()
func (this *QPaintEngineState) RenderHints() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPaintEngineState11renderHintsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:285
// index:0
// Public Visibility=Default Availability=Available
// [4] QPainter::CompositionMode compositionMode()
func (this *QPaintEngineState) CompositionMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPaintEngineState15compositionModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:286
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal opacity()
func (this *QPaintEngineState) Opacity() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPaintEngineState7opacityEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpaintengine.h:288
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainter * painter()
func (this *QPaintEngineState) Painter() *QPainter /*777 QPainter **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPaintEngineState7painterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPainterFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qpaintengine.h:290
// index:0
// Public Visibility=Default Availability=Available
// [1] bool brushNeedsResolving()
func (this *QPaintEngineState) BrushNeedsResolving() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPaintEngineState19brushNeedsResolvingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpaintengine.h:291
// index:0
// Public Visibility=Default Availability=Available
// [1] bool penNeedsResolving()
func (this *QPaintEngineState) PenNeedsResolving() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPaintEngineState17penNeedsResolvingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

func DeleteQPaintEngineState(this *QPaintEngineState) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QPaintEngineStateD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
