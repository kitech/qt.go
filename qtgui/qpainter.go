package qtgui

// /usr/include/qt/QtGui/qpainter.h
// #include <qpainter.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QPainter struct {
	*qtrt.CObject
}
type QPainter_ITF interface {
	QPainter_PTR() *QPainter
}

func (ptr *QPainter) QPainter_PTR() *QPainter { return ptr }

func (this *QPainter) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPainter) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQPainterFromPointer(cthis unsafe.Pointer) *QPainter {
	return &QPainter{&qtrt.CObject{cthis}}
}
func (*QPainter) NewFromPointer(cthis unsafe.Pointer) *QPainter {
	return NewQPainterFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpainter.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPainter()
func NewQPainter() *QPainter {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainterC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPainterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPainter)
	return gothis
}

// /usr/include/qt/QtGui/qpainter.h:125
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPainter(QPaintDevice *)
func NewQPainter_1(arg0 QPaintDevice_ITF /*777 QPaintDevice **/) *QPainter {
	var convArg0 = arg0.QPaintDevice_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainterC2EP12QPaintDevice", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPainterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPainter)
	return gothis
}

// /usr/include/qt/QtGui/qpainter.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QPainter()
func DeleteQPainter(this *QPainter) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainterD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpainter.h:128
// index:0
// Public Visibility=Default Availability=Available
// [8] QPaintDevice * device()
func (this *QPainter) Device() *QPaintDevice /*777 QPaintDevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter6deviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPaintDeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qpainter.h:130
// index:0
// Public Visibility=Default Availability=Available
// [1] bool begin(QPaintDevice *)
func (this *QPainter) Begin(arg0 QPaintDevice_ITF /*777 QPaintDevice **/) bool {
	var convArg0 = arg0.QPaintDevice_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter5beginEP12QPaintDevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainter.h:131
// index:0
// Public Visibility=Default Availability=Available
// [1] bool end()
func (this *QPainter) End() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainter.h:132
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isActive()
func (this *QPainter) IsActive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter8isActiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainter.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void initFrom(const QPaintDevice *)
func (this *QPainter) InitFrom(device QPaintDevice_ITF /*777 const QPaintDevice **/) {
	var convArg0 = device.QPaintDevice_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8initFromEPK12QPaintDevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:180
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCompositionMode(enum QPainter::CompositionMode)
func (this *QPainter) SetCompositionMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter18setCompositionModeENS_15CompositionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:181
// index:0
// Public Visibility=Default Availability=Available
// [4] QPainter::CompositionMode compositionMode()
func (this *QPainter) CompositionMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter15compositionModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpainter.h:183
// index:0
// Public Visibility=Default Availability=Available
// [16] const QFont & font()
func (this *QPainter) Font() *QFont {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter4fontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFont)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:184
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFont(const QFont &)
func (this *QPainter) SetFont(f QFont_ITF) {
	var convArg0 = f.QFont_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter7setFontERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:186
// index:0
// Public Visibility=Default Availability=Available
// [8] QFontMetrics fontMetrics()
func (this *QPainter) FontMetrics() *QFontMetrics /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter11fontMetricsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontMetricsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFontMetrics)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:187
// index:0
// Public Visibility=Default Availability=Available
// [8] QFontInfo fontInfo()
func (this *QPainter) FontInfo() *QFontInfo /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter8fontInfoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFontInfo)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPen(const QColor &)
func (this *QPainter) SetPen(color QColor_ITF) {
	var convArg0 = color.QColor_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter6setPenERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:190
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setPen(const QPen &)
func (this *QPainter) SetPen_1(pen QPen_ITF) {
	var convArg0 = pen.QPen_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter6setPenERK4QPen", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:191
// index:2
// Public Visibility=Default Availability=Available
// [-2] void setPen(Qt::PenStyle)
func (this *QPainter) SetPen_2(style int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter6setPenEN2Qt8PenStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), style)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:192
// index:0
// Public Visibility=Default Availability=Available
// [8] const QPen & pen()
func (this *QPainter) Pen() *QPen {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter3penEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPenFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPen)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:194
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBrush(const QBrush &)
func (this *QPainter) SetBrush(brush QBrush_ITF) {
	var convArg0 = brush.QBrush_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8setBrushERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:195
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setBrush(Qt::BrushStyle)
func (this *QPainter) SetBrush_1(style int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8setBrushEN2Qt10BrushStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), style)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:196
// index:0
// Public Visibility=Default Availability=Available
// [8] const QBrush & brush()
func (this *QPainter) Brush() *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter5brushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:199
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBackgroundMode(Qt::BGMode)
func (this *QPainter) SetBackgroundMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter17setBackgroundModeEN2Qt6BGModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:200
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::BGMode backgroundMode()
func (this *QPainter) BackgroundMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter14backgroundModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpainter.h:202
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint brushOrigin()
func (this *QPainter) BrushOrigin() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter11brushOriginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:203
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBrushOrigin(int, int)
func (this *QPainter) SetBrushOrigin(x int, y int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter14setBrushOriginEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:204
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setBrushOrigin(const QPoint &)
func (this *QPainter) SetBrushOrigin_1(arg0 qtcore.QPoint_ITF) {
	var convArg0 = arg0.QPoint_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter14setBrushOriginERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:205
// index:2
// Public Visibility=Default Availability=Available
// [-2] void setBrushOrigin(const QPointF &)
func (this *QPainter) SetBrushOrigin_2(arg0 qtcore.QPointF_ITF) {
	var convArg0 = arg0.QPointF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter14setBrushOriginERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:207
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBackground(const QBrush &)
func (this *QPainter) SetBackground(bg QBrush_ITF) {
	var convArg0 = bg.QBrush_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter13setBackgroundERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:208
// index:0
// Public Visibility=Default Availability=Available
// [8] const QBrush & background()
func (this *QPainter) Background() *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter10backgroundEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:210
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal opacity()
func (this *QPainter) Opacity() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter7opacityEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpainter.h:211
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpacity(qreal)
func (this *QPainter) SetOpacity(opacity float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10setOpacityEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opacity)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:214
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion clipRegion()
func (this *QPainter) ClipRegion() *QRegion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter10clipRegionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:215
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath clipPath()
func (this *QPainter) ClipPath() *QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter8clipPathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:217
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClipRect(const QRectF &, Qt::ClipOperation)
func (this *QPainter) SetClipRect(arg0 qtcore.QRectF_ITF, op int) {
	var convArg0 = arg0.QRectF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11setClipRectERK6QRectFN2Qt13ClipOperationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, op)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:218
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setClipRect(const QRect &, Qt::ClipOperation)
func (this *QPainter) SetClipRect_1(arg0 qtcore.QRect_ITF, op int) {
	var convArg0 = arg0.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11setClipRectERK5QRectN2Qt13ClipOperationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, op)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:219
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void setClipRect(int, int, int, int, Qt::ClipOperation)
func (this *QPainter) SetClipRect_2(x int, y int, w int, h int, op int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11setClipRectEiiiiN2Qt13ClipOperationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, op)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:221
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClipRegion(const QRegion &, Qt::ClipOperation)
func (this *QPainter) SetClipRegion(arg0 QRegion_ITF, op int) {
	var convArg0 = arg0.QRegion_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter13setClipRegionERK7QRegionN2Qt13ClipOperationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, op)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:223
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClipPath(const QPainterPath &, Qt::ClipOperation)
func (this *QPainter) SetClipPath(path QPainterPath_ITF, op int) {
	var convArg0 = path.QPainterPath_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11setClipPathERK12QPainterPathN2Qt13ClipOperationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, op)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:225
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClipping(_Bool)
func (this *QPainter) SetClipping(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11setClippingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:226
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasClipping()
func (this *QPainter) HasClipping() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter11hasClippingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainter.h:228
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF clipBoundingRect()
func (this *QPainter) ClipBoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter16clipBoundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:230
// index:0
// Public Visibility=Default Availability=Available
// [-2] void save()
func (this *QPainter) Save() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter4saveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:231
// index:0
// Public Visibility=Default Availability=Available
// [-2] void restore()
func (this *QPainter) Restore() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter7restoreEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:234
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMatrix(const QMatrix &, _Bool)
func (this *QPainter) SetMatrix(matrix QMatrix_ITF, combine bool) {
	var convArg0 = matrix.QMatrix_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9setMatrixERK7QMatrixb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:235
// index:0
// Public Visibility=Default Availability=Available
// [48] const QMatrix & matrix()
func (this *QPainter) Matrix() *QMatrix {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter6matrixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:236
// index:0
// Public Visibility=Default Availability=Available
// [48] const QMatrix & deviceMatrix()
func (this *QPainter) DeviceMatrix() *QMatrix {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter12deviceMatrixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:237
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetMatrix()
func (this *QPainter) ResetMatrix() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11resetMatrixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:239
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTransform(const QTransform &, _Bool)
func (this *QPainter) SetTransform(transform QTransform_ITF, combine bool) {
	var convArg0 = transform.QTransform_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12setTransformERK10QTransformb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:240
// index:0
// Public Visibility=Default Availability=Available
// [88] const QTransform & transform()
func (this *QPainter) Transform() *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter9transformEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:241
// index:0
// Public Visibility=Default Availability=Available
// [88] const QTransform & deviceTransform()
func (this *QPainter) DeviceTransform() *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter15deviceTransformEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:242
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetTransform()
func (this *QPainter) ResetTransform() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter14resetTransformEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:244
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWorldMatrix(const QMatrix &, _Bool)
func (this *QPainter) SetWorldMatrix(matrix QMatrix_ITF, combine bool) {
	var convArg0 = matrix.QMatrix_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter14setWorldMatrixERK7QMatrixb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:245
// index:0
// Public Visibility=Default Availability=Available
// [48] const QMatrix & worldMatrix()
func (this *QPainter) WorldMatrix() *QMatrix {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter11worldMatrixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:247
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWorldTransform(const QTransform &, _Bool)
func (this *QPainter) SetWorldTransform(matrix QTransform_ITF, combine bool) {
	var convArg0 = matrix.QTransform_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter17setWorldTransformERK10QTransformb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:248
// index:0
// Public Visibility=Default Availability=Available
// [88] const QTransform & worldTransform()
func (this *QPainter) WorldTransform() *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter14worldTransformEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:250
// index:0
// Public Visibility=Default Availability=Available
// [48] QMatrix combinedMatrix()
func (this *QPainter) CombinedMatrix() *QMatrix /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter14combinedMatrixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:251
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform combinedTransform()
func (this *QPainter) CombinedTransform() *QTransform /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter17combinedTransformEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:253
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMatrixEnabled(_Bool)
func (this *QPainter) SetMatrixEnabled(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter16setMatrixEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:254
// index:0
// Public Visibility=Default Availability=Available
// [1] bool matrixEnabled()
func (this *QPainter) MatrixEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter13matrixEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainter.h:256
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWorldMatrixEnabled(_Bool)
func (this *QPainter) SetWorldMatrixEnabled(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter21setWorldMatrixEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:257
// index:0
// Public Visibility=Default Availability=Available
// [1] bool worldMatrixEnabled()
func (this *QPainter) WorldMatrixEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter18worldMatrixEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainter.h:259
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scale(qreal, qreal)
func (this *QPainter) Scale(sx float64, sy float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter5scaleEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sx, sy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:260
// index:0
// Public Visibility=Default Availability=Available
// [-2] void shear(qreal, qreal)
func (this *QPainter) Shear(sh float64, sv float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter5shearEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sh, sv)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:261
// index:0
// Public Visibility=Default Availability=Available
// [-2] void rotate(qreal)
func (this *QPainter) Rotate(a float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter6rotateEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), a)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:263
// index:0
// Public Visibility=Default Availability=Available
// [-2] void translate(const QPointF &)
func (this *QPainter) Translate(offset qtcore.QPointF_ITF) {
	var convArg0 = offset.QPointF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9translateERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:264
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void translate(const QPoint &)
func (this *QPainter) Translate_1(offset qtcore.QPoint_ITF) {
	var convArg0 = offset.QPoint_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9translateERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:265
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void translate(qreal, qreal)
func (this *QPainter) Translate_2(dx float64, dy float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9translateEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:267
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect window()
func (this *QPainter) Window() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter6windowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:268
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindow(const QRect &)
func (this *QPainter) SetWindow(window qtcore.QRect_ITF) {
	var convArg0 = window.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9setWindowERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:269
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setWindow(int, int, int, int)
func (this *QPainter) SetWindow_1(x int, y int, w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9setWindowEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:271
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect viewport()
func (this *QPainter) Viewport() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter8viewportEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:272
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setViewport(const QRect &)
func (this *QPainter) SetViewport(viewport qtcore.QRect_ITF) {
	var convArg0 = viewport.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11setViewportERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:273
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setViewport(int, int, int, int)
func (this *QPainter) SetViewport_1(x int, y int, w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11setViewportEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:275
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setViewTransformEnabled(_Bool)
func (this *QPainter) SetViewTransformEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter23setViewTransformEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:276
// index:0
// Public Visibility=Default Availability=Available
// [1] bool viewTransformEnabled()
func (this *QPainter) ViewTransformEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter20viewTransformEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainter.h:279
// index:0
// Public Visibility=Default Availability=Available
// [-2] void strokePath(const QPainterPath &, const QPen &)
func (this *QPainter) StrokePath(path QPainterPath_ITF, pen QPen_ITF) {
	var convArg0 = path.QPainterPath_PTR().GetCthis()
	var convArg1 = pen.QPen_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10strokePathERK12QPainterPathRK4QPen", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:280
// index:0
// Public Visibility=Default Availability=Available
// [-2] void fillPath(const QPainterPath &, const QBrush &)
func (this *QPainter) FillPath(path QPainterPath_ITF, brush QBrush_ITF) {
	var convArg0 = path.QPainterPath_PTR().GetCthis()
	var convArg1 = brush.QBrush_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8fillPathERK12QPainterPathRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:281
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawPath(const QPainterPath &)
func (this *QPainter) DrawPath(path QPainterPath_ITF) {
	var convArg0 = path.QPainterPath_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawPathERK12QPainterPath", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:283
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void drawPoint(const QPointF &)
func (this *QPainter) DrawPoint(pt qtcore.QPointF_ITF) {
	var convArg0 = pt.QPointF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawPointERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:284
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawPoint(const QPoint &)
func (this *QPainter) DrawPoint_1(p qtcore.QPoint_ITF) {
	var convArg0 = p.QPoint_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawPointERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:285
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawPoint(int, int)
func (this *QPainter) DrawPoint_2(x int, y int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawPointEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:287
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawPoints(const QPointF *, int)
func (this *QPainter) DrawPoints(points qtcore.QPointF_ITF /*777 const QPointF **/, pointCount int) {
	var convArg0 = points.QPointF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPointsEPK7QPointFi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:288
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawPoints(const QPolygonF &)
func (this *QPainter) DrawPoints_1(points QPolygonF_ITF) {
	var convArg0 = points.QPolygonF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPointsERK9QPolygonF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:289
// index:2
// Public Visibility=Default Availability=Available
// [-2] void drawPoints(const QPoint *, int)
func (this *QPainter) DrawPoints_2(points qtcore.QPoint_ITF /*777 const QPoint **/, pointCount int) {
	var convArg0 = points.QPoint_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPointsEPK6QPointi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:290
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void drawPoints(const QPolygon &)
func (this *QPainter) DrawPoints_3(points QPolygon_ITF) {
	var convArg0 = points.QPolygon_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPointsERK8QPolygon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:292
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void drawLine(const QLineF &)
func (this *QPainter) DrawLine(line qtcore.QLineF_ITF) {
	var convArg0 = line.QLineF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawLineERK6QLineF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:293
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawLine(const QLine &)
func (this *QPainter) DrawLine_1(line qtcore.QLine_ITF) {
	var convArg0 = line.QLine_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawLineERK5QLine", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:294
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawLine(int, int, int, int)
func (this *QPainter) DrawLine_2(x1 int, y1 int, x2 int, y2 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawLineEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, x2, y2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:295
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void drawLine(const QPoint &, const QPoint &)
func (this *QPainter) DrawLine_3(p1 qtcore.QPoint_ITF, p2 qtcore.QPoint_ITF) {
	var convArg0 = p1.QPoint_PTR().GetCthis()
	var convArg1 = p2.QPoint_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawLineERK6QPointS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:296
// index:4
// Public inline Visibility=Default Availability=Available
// [-2] void drawLine(const QPointF &, const QPointF &)
func (this *QPainter) DrawLine_4(p1 qtcore.QPointF_ITF, p2 qtcore.QPointF_ITF) {
	var convArg0 = p1.QPointF_PTR().GetCthis()
	var convArg1 = p2.QPointF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawLineERK7QPointFS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:298
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawLines(const QLineF *, int)
func (this *QPainter) DrawLines(lines qtcore.QLineF_ITF /*777 const QLineF **/, lineCount int) {
	var convArg0 = lines.QLineF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawLinesEPK6QLineFi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, lineCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:300
// index:1
// Public Visibility=Default Availability=Available
// [-2] void drawLines(const QPointF *, int)
func (this *QPainter) DrawLines_1(pointPairs qtcore.QPointF_ITF /*777 const QPointF **/, lineCount int) {
	var convArg0 = pointPairs.QPointF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawLinesEPK7QPointFi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, lineCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:302
// index:2
// Public Visibility=Default Availability=Available
// [-2] void drawLines(const QLine *, int)
func (this *QPainter) DrawLines_2(lines qtcore.QLine_ITF /*777 const QLine **/, lineCount int) {
	var convArg0 = lines.QLine_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawLinesEPK5QLinei", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, lineCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:304
// index:3
// Public Visibility=Default Availability=Available
// [-2] void drawLines(const QPoint *, int)
func (this *QPainter) DrawLines_3(pointPairs qtcore.QPoint_ITF /*777 const QPoint **/, lineCount int) {
	var convArg0 = pointPairs.QPoint_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawLinesEPK6QPointi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, lineCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:307
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void drawRect(const QRectF &)
func (this *QPainter) DrawRect(rect qtcore.QRectF_ITF) {
	var convArg0 = rect.QRectF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawRectERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:308
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawRect(int, int, int, int)
func (this *QPainter) DrawRect_1(x1 int, y1 int, w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawRectEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:309
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawRect(const QRect &)
func (this *QPainter) DrawRect_2(rect qtcore.QRect_ITF) {
	var convArg0 = rect.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawRectERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:311
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawRects(const QRectF *, int)
func (this *QPainter) DrawRects(rects qtcore.QRectF_ITF /*777 const QRectF **/, rectCount int) {
	var convArg0 = rects.QRectF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawRectsEPK6QRectFi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, rectCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:313
// index:1
// Public Visibility=Default Availability=Available
// [-2] void drawRects(const QRect *, int)
func (this *QPainter) DrawRects_1(rects qtcore.QRect_ITF /*777 const QRect **/, rectCount int) {
	var convArg0 = rects.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawRectsEPK5QRecti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, rectCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:316
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawEllipse(const QRectF &)
func (this *QPainter) DrawEllipse(r qtcore.QRectF_ITF) {
	var convArg0 = r.QRectF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawEllipseERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:317
// index:1
// Public Visibility=Default Availability=Available
// [-2] void drawEllipse(const QRect &)
func (this *QPainter) DrawEllipse_1(r qtcore.QRect_ITF) {
	var convArg0 = r.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawEllipseERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:318
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawEllipse(int, int, int, int)
func (this *QPainter) DrawEllipse_2(x int, y int, w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawEllipseEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:320
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void drawEllipse(const QPointF &, qreal, qreal)
func (this *QPainter) DrawEllipse_3(center qtcore.QPointF_ITF, rx float64, ry float64) {
	var convArg0 = center.QPointF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawEllipseERK7QPointFdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, rx, ry)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:321
// index:4
// Public inline Visibility=Default Availability=Available
// [-2] void drawEllipse(const QPoint &, int, int)
func (this *QPainter) DrawEllipse_4(center qtcore.QPoint_ITF, rx int, ry int) {
	var convArg0 = center.QPoint_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawEllipseERK6QPointii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, rx, ry)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:323
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawPolyline(const QPointF *, int)
func (this *QPainter) DrawPolyline(points qtcore.QPointF_ITF /*777 const QPointF **/, pointCount int) {
	var convArg0 = points.QPointF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12drawPolylineEPK7QPointFi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:324
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawPolyline(const QPolygonF &)
func (this *QPainter) DrawPolyline_1(polyline QPolygonF_ITF) {
	var convArg0 = polyline.QPolygonF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12drawPolylineERK9QPolygonF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:325
// index:2
// Public Visibility=Default Availability=Available
// [-2] void drawPolyline(const QPoint *, int)
func (this *QPainter) DrawPolyline_2(points qtcore.QPoint_ITF /*777 const QPoint **/, pointCount int) {
	var convArg0 = points.QPoint_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12drawPolylineEPK6QPointi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:326
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void drawPolyline(const QPolygon &)
func (this *QPainter) DrawPolyline_3(polygon QPolygon_ITF) {
	var convArg0 = polygon.QPolygon_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12drawPolylineERK8QPolygon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:328
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawPolygon(const QPointF *, int, Qt::FillRule)
func (this *QPainter) DrawPolygon(points qtcore.QPointF_ITF /*777 const QPointF **/, pointCount int, fillRule int) {
	var convArg0 = points.QPointF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawPolygonEPK7QPointFiN2Qt8FillRuleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount, fillRule)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:329
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawPolygon(const QPolygonF &, Qt::FillRule)
func (this *QPainter) DrawPolygon_1(polygon QPolygonF_ITF, fillRule int) {
	var convArg0 = polygon.QPolygonF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawPolygonERK9QPolygonFN2Qt8FillRuleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, fillRule)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:330
// index:2
// Public Visibility=Default Availability=Available
// [-2] void drawPolygon(const QPoint *, int, Qt::FillRule)
func (this *QPainter) DrawPolygon_2(points qtcore.QPoint_ITF /*777 const QPoint **/, pointCount int, fillRule int) {
	var convArg0 = points.QPoint_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawPolygonEPK6QPointiN2Qt8FillRuleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount, fillRule)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:331
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void drawPolygon(const QPolygon &, Qt::FillRule)
func (this *QPainter) DrawPolygon_3(polygon QPolygon_ITF, fillRule int) {
	var convArg0 = polygon.QPolygon_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawPolygonERK8QPolygonN2Qt8FillRuleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, fillRule)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:333
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawConvexPolygon(const QPointF *, int)
func (this *QPainter) DrawConvexPolygon(points qtcore.QPointF_ITF /*777 const QPointF **/, pointCount int) {
	var convArg0 = points.QPointF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter17drawConvexPolygonEPK7QPointFi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:334
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawConvexPolygon(const QPolygonF &)
func (this *QPainter) DrawConvexPolygon_1(polygon QPolygonF_ITF) {
	var convArg0 = polygon.QPolygonF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter17drawConvexPolygonERK9QPolygonF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:335
// index:2
// Public Visibility=Default Availability=Available
// [-2] void drawConvexPolygon(const QPoint *, int)
func (this *QPainter) DrawConvexPolygon_2(points qtcore.QPoint_ITF /*777 const QPoint **/, pointCount int) {
	var convArg0 = points.QPoint_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter17drawConvexPolygonEPK6QPointi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:336
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void drawConvexPolygon(const QPolygon &)
func (this *QPainter) DrawConvexPolygon_3(polygon QPolygon_ITF) {
	var convArg0 = polygon.QPolygon_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter17drawConvexPolygonERK8QPolygon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:338
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawArc(const QRectF &, int, int)
func (this *QPainter) DrawArc(rect qtcore.QRectF_ITF, a int, alen int) {
	var convArg0 = rect.QRectF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter7drawArcERK6QRectFii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, a, alen)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:339
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawArc(const QRect &, int, int)
func (this *QPainter) DrawArc_1(arg0 qtcore.QRect_ITF, a int, alen int) {
	var convArg0 = arg0.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter7drawArcERK5QRectii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, a, alen)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:340
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawArc(int, int, int, int, int, int)
func (this *QPainter) DrawArc_2(x int, y int, w int, h int, a int, alen int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter7drawArcEiiiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, a, alen)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:342
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawPie(const QRectF &, int, int)
func (this *QPainter) DrawPie(rect qtcore.QRectF_ITF, a int, alen int) {
	var convArg0 = rect.QRectF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter7drawPieERK6QRectFii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, a, alen)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:343
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawPie(int, int, int, int, int, int)
func (this *QPainter) DrawPie_1(x int, y int, w int, h int, a int, alen int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter7drawPieEiiiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, a, alen)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:344
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawPie(const QRect &, int, int)
func (this *QPainter) DrawPie_2(arg0 qtcore.QRect_ITF, a int, alen int) {
	var convArg0 = arg0.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter7drawPieERK5QRectii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, a, alen)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:346
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawChord(const QRectF &, int, int)
func (this *QPainter) DrawChord(rect qtcore.QRectF_ITF, a int, alen int) {
	var convArg0 = rect.QRectF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawChordERK6QRectFii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, a, alen)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:347
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawChord(int, int, int, int, int, int)
func (this *QPainter) DrawChord_1(x int, y int, w int, h int, a int, alen int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawChordEiiiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, a, alen)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:348
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawChord(const QRect &, int, int)
func (this *QPainter) DrawChord_2(arg0 qtcore.QRect_ITF, a int, alen int) {
	var convArg0 = arg0.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawChordERK5QRectii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, a, alen)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:350
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawRoundedRect(const QRectF &, qreal, qreal, Qt::SizeMode)
func (this *QPainter) DrawRoundedRect(rect qtcore.QRectF_ITF, xRadius float64, yRadius float64, mode int) {
	var convArg0 = rect.QRectF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter15drawRoundedRectERK6QRectFddN2Qt8SizeModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xRadius, yRadius, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:352
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawRoundedRect(int, int, int, int, qreal, qreal, Qt::SizeMode)
func (this *QPainter) DrawRoundedRect_1(x int, y int, w int, h int, xRadius float64, yRadius float64, mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter15drawRoundedRectEiiiiddN2Qt8SizeModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, xRadius, yRadius, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:354
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawRoundedRect(const QRect &, qreal, qreal, Qt::SizeMode)
func (this *QPainter) DrawRoundedRect_2(rect qtcore.QRect_ITF, xRadius float64, yRadius float64, mode int) {
	var convArg0 = rect.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter15drawRoundedRectERK5QRectddN2Qt8SizeModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xRadius, yRadius, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:357
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawRoundRect(const QRectF &, int, int)
func (this *QPainter) DrawRoundRect(r qtcore.QRectF_ITF, xround int, yround int) {
	var convArg0 = r.QRectF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter13drawRoundRectERK6QRectFii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xround, yround)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:358
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawRoundRect(int, int, int, int, int, int)
func (this *QPainter) DrawRoundRect_1(x int, y int, w int, h int, arg4 int, arg5 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter13drawRoundRectEiiiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, arg4, arg5)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:359
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawRoundRect(const QRect &, int, int)
func (this *QPainter) DrawRoundRect_2(r qtcore.QRect_ITF, xround int, yround int) {
	var convArg0 = r.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter13drawRoundRectERK5QRectii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xround, yround)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:361
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawTiledPixmap(const QRectF &, const QPixmap &, const QPointF &)
func (this *QPainter) DrawTiledPixmap(rect qtcore.QRectF_ITF, pm QPixmap_ITF, offset qtcore.QPointF_ITF) {
	var convArg0 = rect.QRectF_PTR().GetCthis()
	var convArg1 = pm.QPixmap_PTR().GetCthis()
	var convArg2 = offset.QPointF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter15drawTiledPixmapERK6QRectFRK7QPixmapRK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:362
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawTiledPixmap(int, int, int, int, const QPixmap &, int, int)
func (this *QPainter) DrawTiledPixmap_1(x int, y int, w int, h int, arg4 QPixmap_ITF, sx int, sy int) {
	var convArg4 = arg4.QPixmap_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter15drawTiledPixmapEiiiiRK7QPixmapii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, convArg4, sx, sy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:363
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawTiledPixmap(const QRect &, const QPixmap &, const QPoint &)
func (this *QPainter) DrawTiledPixmap_2(arg0 qtcore.QRect_ITF, arg1 QPixmap_ITF, arg2 qtcore.QPoint_ITF) {
	var convArg0 = arg0.QRect_PTR().GetCthis()
	var convArg1 = arg1.QPixmap_PTR().GetCthis()
	var convArg2 = arg2.QPoint_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter15drawTiledPixmapERK5QRectRK7QPixmapRK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:365
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawPicture(const QPointF &, const QPicture &)
func (this *QPainter) DrawPicture(p qtcore.QPointF_ITF, picture QPicture_ITF) {
	var convArg0 = p.QPointF_PTR().GetCthis()
	var convArg1 = picture.QPicture_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawPictureERK7QPointFRK8QPicture", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:366
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawPicture(int, int, const QPicture &)
func (this *QPainter) DrawPicture_1(x int, y int, picture QPicture_ITF) {
	var convArg2 = picture.QPicture_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawPictureEiiRK8QPicture", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:367
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawPicture(const QPoint &, const QPicture &)
func (this *QPainter) DrawPicture_2(p qtcore.QPoint_ITF, picture QPicture_ITF) {
	var convArg0 = p.QPoint_PTR().GetCthis()
	var convArg1 = picture.QPicture_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawPictureERK6QPointRK8QPicture", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:370
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawPixmap(const QRectF &, const QPixmap &, const QRectF &)
func (this *QPainter) DrawPixmap(targetRect qtcore.QRectF_ITF, pixmap QPixmap_ITF, sourceRect qtcore.QRectF_ITF) {
	var convArg0 = targetRect.QRectF_PTR().GetCthis()
	var convArg1 = pixmap.QPixmap_PTR().GetCthis()
	var convArg2 = sourceRect.QRectF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK6QRectFRK7QPixmapS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:371
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawPixmap(const QRect &, const QPixmap &, const QRect &)
func (this *QPainter) DrawPixmap_1(targetRect qtcore.QRect_ITF, pixmap QPixmap_ITF, sourceRect qtcore.QRect_ITF) {
	var convArg0 = targetRect.QRect_PTR().GetCthis()
	var convArg1 = pixmap.QPixmap_PTR().GetCthis()
	var convArg2 = sourceRect.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK5QRectRK7QPixmapS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:372
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawPixmap(int, int, int, int, const QPixmap &, int, int, int, int)
func (this *QPainter) DrawPixmap_2(x int, y int, w int, h int, pm QPixmap_ITF, sx int, sy int, sw int, sh int) {
	var convArg4 = pm.QPixmap_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPixmapEiiiiRK7QPixmapiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, convArg4, sx, sy, sw, sh)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:374
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void drawPixmap(int, int, const QPixmap &, int, int, int, int)
func (this *QPainter) DrawPixmap_3(x int, y int, pm QPixmap_ITF, sx int, sy int, sw int, sh int) {
	var convArg2 = pm.QPixmap_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPixmapEiiRK7QPixmapiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2, sx, sy, sw, sh)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:376
// index:4
// Public inline Visibility=Default Availability=Available
// [-2] void drawPixmap(const QPointF &, const QPixmap &, const QRectF &)
func (this *QPainter) DrawPixmap_4(p qtcore.QPointF_ITF, pm QPixmap_ITF, sr qtcore.QRectF_ITF) {
	var convArg0 = p.QPointF_PTR().GetCthis()
	var convArg1 = pm.QPixmap_PTR().GetCthis()
	var convArg2 = sr.QRectF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK7QPointFRK7QPixmapRK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:377
// index:5
// Public inline Visibility=Default Availability=Available
// [-2] void drawPixmap(const QPoint &, const QPixmap &, const QRect &)
func (this *QPainter) DrawPixmap_5(p qtcore.QPoint_ITF, pm QPixmap_ITF, sr qtcore.QRect_ITF) {
	var convArg0 = p.QPoint_PTR().GetCthis()
	var convArg1 = pm.QPixmap_PTR().GetCthis()
	var convArg2 = sr.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK6QPointRK7QPixmapRK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:378
// index:6
// Public Visibility=Default Availability=Available
// [-2] void drawPixmap(const QPointF &, const QPixmap &)
func (this *QPainter) DrawPixmap_6(p qtcore.QPointF_ITF, pm QPixmap_ITF) {
	var convArg0 = p.QPointF_PTR().GetCthis()
	var convArg1 = pm.QPixmap_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK7QPointFRK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:379
// index:7
// Public inline Visibility=Default Availability=Available
// [-2] void drawPixmap(const QPoint &, const QPixmap &)
func (this *QPainter) DrawPixmap_7(p qtcore.QPoint_ITF, pm QPixmap_ITF) {
	var convArg0 = p.QPoint_PTR().GetCthis()
	var convArg1 = pm.QPixmap_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK6QPointRK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:380
// index:8
// Public inline Visibility=Default Availability=Available
// [-2] void drawPixmap(int, int, const QPixmap &)
func (this *QPainter) DrawPixmap_8(x int, y int, pm QPixmap_ITF) {
	var convArg2 = pm.QPixmap_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPixmapEiiRK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:381
// index:9
// Public inline Visibility=Default Availability=Available
// [-2] void drawPixmap(const QRect &, const QPixmap &)
func (this *QPainter) DrawPixmap_9(r qtcore.QRect_ITF, pm QPixmap_ITF) {
	var convArg0 = r.QRect_PTR().GetCthis()
	var convArg1 = pm.QPixmap_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK5QRectRK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:382
// index:10
// Public inline Visibility=Default Availability=Available
// [-2] void drawPixmap(int, int, int, int, const QPixmap &)
func (this *QPainter) DrawPixmap_10(x int, y int, w int, h int, pm QPixmap_ITF) {
	var convArg4 = pm.QPixmap_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPixmapEiiiiRK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, convArg4)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:387
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawImage(const QRectF &, const QImage &, const QRectF &, Qt::ImageConversionFlags)
func (this *QPainter) DrawImage(targetRect qtcore.QRectF_ITF, image QImage_ITF, sourceRect qtcore.QRectF_ITF, flags int) {
	var convArg0 = targetRect.QRectF_PTR().GetCthis()
	var convArg1 = image.QImage_PTR().GetCthis()
	var convArg2 = sourceRect.QRectF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawImageERK6QRectFRK6QImageS2_6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:389
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawImage(const QRect &, const QImage &, const QRect &, Qt::ImageConversionFlags)
func (this *QPainter) DrawImage_1(targetRect qtcore.QRect_ITF, image QImage_ITF, sourceRect qtcore.QRect_ITF, flags int) {
	var convArg0 = targetRect.QRect_PTR().GetCthis()
	var convArg1 = image.QImage_PTR().GetCthis()
	var convArg2 = sourceRect.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawImageERK5QRectRK6QImageS2_6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:391
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawImage(const QPointF &, const QImage &, const QRectF &, Qt::ImageConversionFlags)
func (this *QPainter) DrawImage_2(p qtcore.QPointF_ITF, image QImage_ITF, sr qtcore.QRectF_ITF, flags int) {
	var convArg0 = p.QPointF_PTR().GetCthis()
	var convArg1 = image.QImage_PTR().GetCthis()
	var convArg2 = sr.QRectF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawImageERK7QPointFRK6QImageRK6QRectF6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:393
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void drawImage(const QPoint &, const QImage &, const QRect &, Qt::ImageConversionFlags)
func (this *QPainter) DrawImage_3(p qtcore.QPoint_ITF, image QImage_ITF, sr qtcore.QRect_ITF, flags int) {
	var convArg0 = p.QPoint_PTR().GetCthis()
	var convArg1 = image.QImage_PTR().GetCthis()
	var convArg2 = sr.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawImageERK6QPointRK6QImageRK5QRect6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:395
// index:4
// Public inline Visibility=Default Availability=Available
// [-2] void drawImage(const QRectF &, const QImage &)
func (this *QPainter) DrawImage_4(r qtcore.QRectF_ITF, image QImage_ITF) {
	var convArg0 = r.QRectF_PTR().GetCthis()
	var convArg1 = image.QImage_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawImageERK6QRectFRK6QImage", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:396
// index:5
// Public inline Visibility=Default Availability=Available
// [-2] void drawImage(const QRect &, const QImage &)
func (this *QPainter) DrawImage_5(r qtcore.QRect_ITF, image QImage_ITF) {
	var convArg0 = r.QRect_PTR().GetCthis()
	var convArg1 = image.QImage_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawImageERK5QRectRK6QImage", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:397
// index:6
// Public Visibility=Default Availability=Available
// [-2] void drawImage(const QPointF &, const QImage &)
func (this *QPainter) DrawImage_6(p qtcore.QPointF_ITF, image QImage_ITF) {
	var convArg0 = p.QPointF_PTR().GetCthis()
	var convArg1 = image.QImage_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawImageERK7QPointFRK6QImage", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:398
// index:7
// Public inline Visibility=Default Availability=Available
// [-2] void drawImage(const QPoint &, const QImage &)
func (this *QPainter) DrawImage_7(p qtcore.QPoint_ITF, image QImage_ITF) {
	var convArg0 = p.QPoint_PTR().GetCthis()
	var convArg1 = image.QImage_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawImageERK6QPointRK6QImage", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:399
// index:8
// Public inline Visibility=Default Availability=Available
// [-2] void drawImage(int, int, const QImage &, int, int, int, int, Qt::ImageConversionFlags)
func (this *QPainter) DrawImage_8(x int, y int, image QImage_ITF, sx int, sy int, sw int, sh int, flags int) {
	var convArg2 = image.QImage_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawImageEiiRK6QImageiiii6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2, sx, sy, sw, sh, flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:402
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLayoutDirection(Qt::LayoutDirection)
func (this *QPainter) SetLayoutDirection(direction int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter18setLayoutDirectionEN2Qt15LayoutDirectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), direction)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:403
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::LayoutDirection layoutDirection()
func (this *QPainter) LayoutDirection() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter15layoutDirectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpainter.h:406
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawGlyphRun(const QPointF &, const QGlyphRun &)
func (this *QPainter) DrawGlyphRun(position qtcore.QPointF_ITF, glyphRun QGlyphRun_ITF) {
	var convArg0 = position.QPointF_PTR().GetCthis()
	var convArg1 = glyphRun.QGlyphRun_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12drawGlyphRunERK7QPointFRK9QGlyphRun", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:409
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawStaticText(const QPointF &, const QStaticText &)
func (this *QPainter) DrawStaticText(topLeftPosition qtcore.QPointF_ITF, staticText QStaticText_ITF) {
	var convArg0 = topLeftPosition.QPointF_PTR().GetCthis()
	var convArg1 = staticText.QStaticText_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter14drawStaticTextERK7QPointFRK11QStaticText", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:410
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawStaticText(const QPoint &, const QStaticText &)
func (this *QPainter) DrawStaticText_1(topLeftPosition qtcore.QPoint_ITF, staticText QStaticText_ITF) {
	var convArg0 = topLeftPosition.QPoint_PTR().GetCthis()
	var convArg1 = staticText.QStaticText_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter14drawStaticTextERK6QPointRK11QStaticText", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:411
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawStaticText(int, int, const QStaticText &)
func (this *QPainter) DrawStaticText_2(left int, top int, staticText QStaticText_ITF) {
	var convArg2 = staticText.QStaticText_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter14drawStaticTextEiiRK11QStaticText", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, top, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:413
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawText(const QPointF &, const QString &)
func (this *QPainter) DrawText(p qtcore.QPointF_ITF, s string) {
	var convArg0 = p.QPointF_PTR().GetCthis()
	var tmpArg1 = qtcore.NewQString_5(s)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawTextERK7QPointFRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:414
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawText(const QPoint &, const QString &)
func (this *QPainter) DrawText_1(p qtcore.QPoint_ITF, s string) {
	var convArg0 = p.QPoint_PTR().GetCthis()
	var tmpArg1 = qtcore.NewQString_5(s)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawTextERK6QPointRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:415
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawText(int, int, const QString &)
func (this *QPainter) DrawText_2(x int, y int, s string) {
	var tmpArg2 = qtcore.NewQString_5(s)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawTextEiiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:417
// index:3
// Public Visibility=Default Availability=Available
// [-2] void drawText(const QPointF &, const QString &, int, int)
func (this *QPainter) DrawText_3(p qtcore.QPointF_ITF, str string, tf int, justificationPadding int) {
	var convArg0 = p.QPointF_PTR().GetCthis()
	var tmpArg1 = qtcore.NewQString_5(str)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawTextERK7QPointFRK7QStringii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, tf, justificationPadding)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:419
// index:4
// Public Visibility=Default Availability=Available
// [-2] void drawText(const QRectF &, int, const QString &, QRectF *)
func (this *QPainter) DrawText_4(r qtcore.QRectF_ITF, flags int, text string, br qtcore.QRectF_ITF /*777 QRectF **/) {
	var convArg0 = r.QRectF_PTR().GetCthis()
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	var convArg3 = br.QRectF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawTextERK6QRectFiRK7QStringPS0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:420
// index:5
// Public Visibility=Default Availability=Available
// [-2] void drawText(const QRect &, int, const QString &, QRect *)
func (this *QPainter) DrawText_5(r qtcore.QRect_ITF, flags int, text string, br qtcore.QRect_ITF /*777 QRect **/) {
	var convArg0 = r.QRect_PTR().GetCthis()
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	var convArg3 = br.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawTextERK5QRectiRK7QStringPS0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:421
// index:6
// Public inline Visibility=Default Availability=Available
// [-2] void drawText(int, int, int, int, int, const QString &, QRect *)
func (this *QPainter) DrawText_6(x int, y int, w int, h int, flags int, text string, br qtcore.QRect_ITF /*777 QRect **/) {
	var tmpArg5 = qtcore.NewQString_5(text)
	var convArg5 = tmpArg5.GetCthis()
	var convArg6 = br.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawTextEiiiiiRK7QStringP5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, flags, convArg5, convArg6)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:423
// index:7
// Public Visibility=Default Availability=Available
// [-2] void drawText(const QRectF &, const QString &, const QTextOption &)
func (this *QPainter) DrawText_7(r qtcore.QRectF_ITF, text string, o QTextOption_ITF) {
	var convArg0 = r.QRectF_PTR().GetCthis()
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 = o.QTextOption_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawTextERK6QRectFRK7QStringRK11QTextOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:425
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect(const QRectF &, int, const QString &)
func (this *QPainter) BoundingRect(rect qtcore.QRectF_ITF, flags int, text string) *qtcore.QRectF /*123*/ {
	var convArg0 = rect.QRectF_PTR().GetCthis()
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12boundingRectERK6QRectFiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:426
// index:1
// Public Visibility=Default Availability=Available
// [16] QRect boundingRect(const QRect &, int, const QString &)
func (this *QPainter) BoundingRect_1(rect qtcore.QRect_ITF, flags int, text string) *qtcore.QRect /*123*/ {
	var convArg0 = rect.QRect_PTR().GetCthis()
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12boundingRectERK5QRectiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:427
// index:2
// Public inline Visibility=Default Availability=Available
// [16] QRect boundingRect(int, int, int, int, int, const QString &)
func (this *QPainter) BoundingRect_2(x int, y int, w int, h int, flags int, text string) *qtcore.QRect /*123*/ {
	var tmpArg5 = qtcore.NewQString_5(text)
	var convArg5 = tmpArg5.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12boundingRectEiiiiiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, flags, convArg5)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:429
// index:3
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect(const QRectF &, const QString &, const QTextOption &)
func (this *QPainter) BoundingRect_3(rect qtcore.QRectF_ITF, text string, o QTextOption_ITF) *qtcore.QRectF /*123*/ {
	var convArg0 = rect.QRectF_PTR().GetCthis()
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 = o.QTextOption_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12boundingRectERK6QRectFRK7QStringRK11QTextOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:431
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawTextItem(const QPointF &, const QTextItem &)
func (this *QPainter) DrawTextItem(p qtcore.QPointF_ITF, ti QTextItem_ITF) {
	var convArg0 = p.QPointF_PTR().GetCthis()
	var convArg1 = ti.QTextItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12drawTextItemERK7QPointFRK9QTextItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:432
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawTextItem(int, int, const QTextItem &)
func (this *QPainter) DrawTextItem_1(x int, y int, ti QTextItem_ITF) {
	var convArg2 = ti.QTextItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12drawTextItemEiiRK9QTextItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:433
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawTextItem(const QPoint &, const QTextItem &)
func (this *QPainter) DrawTextItem_2(p qtcore.QPoint_ITF, ti QTextItem_ITF) {
	var convArg0 = p.QPoint_PTR().GetCthis()
	var convArg1 = ti.QTextItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12drawTextItemERK6QPointRK9QTextItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:435
// index:0
// Public Visibility=Default Availability=Available
// [-2] void fillRect(const QRectF &, const QBrush &)
func (this *QPainter) FillRect(arg0 qtcore.QRectF_ITF, arg1 QBrush_ITF) {
	var convArg0 = arg0.QRectF_PTR().GetCthis()
	var convArg1 = arg1.QBrush_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8fillRectERK6QRectFRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:436
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void fillRect(int, int, int, int, const QBrush &)
func (this *QPainter) FillRect_1(x int, y int, w int, h int, arg4 QBrush_ITF) {
	var convArg4 = arg4.QBrush_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8fillRectEiiiiRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, convArg4)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:437
// index:2
// Public Visibility=Default Availability=Available
// [-2] void fillRect(const QRect &, const QBrush &)
func (this *QPainter) FillRect_2(arg0 qtcore.QRect_ITF, arg1 QBrush_ITF) {
	var convArg0 = arg0.QRect_PTR().GetCthis()
	var convArg1 = arg1.QBrush_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8fillRectERK5QRectRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:439
// index:3
// Public Visibility=Default Availability=Available
// [-2] void fillRect(const QRectF &, const QColor &)
func (this *QPainter) FillRect_3(arg0 qtcore.QRectF_ITF, color QColor_ITF) {
	var convArg0 = arg0.QRectF_PTR().GetCthis()
	var convArg1 = color.QColor_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8fillRectERK6QRectFRK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:440
// index:4
// Public inline Visibility=Default Availability=Available
// [-2] void fillRect(int, int, int, int, const QColor &)
func (this *QPainter) FillRect_4(x int, y int, w int, h int, color QColor_ITF) {
	var convArg4 = color.QColor_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8fillRectEiiiiRK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, convArg4)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:441
// index:5
// Public Visibility=Default Availability=Available
// [-2] void fillRect(const QRect &, const QColor &)
func (this *QPainter) FillRect_5(arg0 qtcore.QRect_ITF, color QColor_ITF) {
	var convArg0 = arg0.QRect_PTR().GetCthis()
	var convArg1 = color.QColor_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8fillRectERK5QRectRK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:443
// index:6
// Public inline Visibility=Default Availability=Available
// [-2] void fillRect(int, int, int, int, Qt::GlobalColor)
func (this *QPainter) FillRect_6(x int, y int, w int, h int, c int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8fillRectEiiiiN2Qt11GlobalColorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, c)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:444
// index:7
// Public inline Visibility=Default Availability=Available
// [-2] void fillRect(const QRect &, Qt::GlobalColor)
func (this *QPainter) FillRect_7(r qtcore.QRect_ITF, c int) {
	var convArg0 = r.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8fillRectERK5QRectN2Qt11GlobalColorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, c)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:445
// index:8
// Public inline Visibility=Default Availability=Available
// [-2] void fillRect(const QRectF &, Qt::GlobalColor)
func (this *QPainter) FillRect_8(r qtcore.QRectF_ITF, c int) {
	var convArg0 = r.QRectF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8fillRectERK6QRectFN2Qt11GlobalColorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, c)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:447
// index:9
// Public inline Visibility=Default Availability=Available
// [-2] void fillRect(int, int, int, int, Qt::BrushStyle)
func (this *QPainter) FillRect_9(x int, y int, w int, h int, style int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8fillRectEiiiiN2Qt10BrushStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, style)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:448
// index:10
// Public inline Visibility=Default Availability=Available
// [-2] void fillRect(const QRect &, Qt::BrushStyle)
func (this *QPainter) FillRect_10(r qtcore.QRect_ITF, style int) {
	var convArg0 = r.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8fillRectERK5QRectN2Qt10BrushStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, style)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:449
// index:11
// Public inline Visibility=Default Availability=Available
// [-2] void fillRect(const QRectF &, Qt::BrushStyle)
func (this *QPainter) FillRect_11(r qtcore.QRectF_ITF, style int) {
	var convArg0 = r.QRectF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8fillRectERK6QRectFN2Qt10BrushStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, style)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:451
// index:0
// Public Visibility=Default Availability=Available
// [-2] void eraseRect(const QRectF &)
func (this *QPainter) EraseRect(arg0 qtcore.QRectF_ITF) {
	var convArg0 = arg0.QRectF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9eraseRectERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:452
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void eraseRect(int, int, int, int)
func (this *QPainter) EraseRect_1(x int, y int, w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9eraseRectEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:453
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void eraseRect(const QRect &)
func (this *QPainter) EraseRect_2(arg0 qtcore.QRect_ITF) {
	var convArg0 = arg0.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9eraseRectERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:455
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRenderHint(enum QPainter::RenderHint, _Bool)
func (this *QPainter) SetRenderHint(hint int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter13setRenderHintENS_10RenderHintEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hint, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:456
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRenderHints(QPainter::RenderHints, _Bool)
func (this *QPainter) SetRenderHints(hints int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter14setRenderHintsE6QFlagsINS_10RenderHintEEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hints, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:457
// index:0
// Public Visibility=Default Availability=Available
// [4] QPainter::RenderHints renderHints()
func (this *QPainter) RenderHints() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter11renderHintsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpainter.h:458
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool testRenderHint(enum QPainter::RenderHint)
func (this *QPainter) TestRenderHint(hint int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter14testRenderHintENS_10RenderHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hint)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainter.h:460
// index:0
// Public Visibility=Default Availability=Available
// [8] QPaintEngine * paintEngine()
func (this *QPainter) PaintEngine() *QPaintEngine /*777 QPaintEngine **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter11paintEngineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPaintEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qpainter.h:462
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setRedirected(const QPaintDevice *, QPaintDevice *, const QPoint &)
func (this *QPainter) SetRedirected(device QPaintDevice_ITF /*777 const QPaintDevice **/, replacement QPaintDevice_ITF /*777 QPaintDevice **/, offset qtcore.QPoint_ITF) {
	var convArg0 = device.QPaintDevice_PTR().GetCthis()
	var convArg1 = replacement.QPaintDevice_PTR().GetCthis()
	var convArg2 = offset.QPoint_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter13setRedirectedEPK12QPaintDevicePS0_RK6QPoint", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}
func QPainter_SetRedirected(device QPaintDevice_ITF /*777 const QPaintDevice **/, replacement QPaintDevice_ITF /*777 QPaintDevice **/, offset qtcore.QPoint_ITF) {
	var nilthis *QPainter
	nilthis.SetRedirected(device, replacement, offset)
}

// /usr/include/qt/QtGui/qpainter.h:464
// index:0
// Public static Visibility=Default Availability=Available
// [8] QPaintDevice * redirected(const QPaintDevice *, QPoint *)
func (this *QPainter) Redirected(device QPaintDevice_ITF /*777 const QPaintDevice **/, offset qtcore.QPoint_ITF /*777 QPoint **/) *QPaintDevice /*777 QPaintDevice **/ {
	var convArg0 = device.QPaintDevice_PTR().GetCthis()
	var convArg1 = offset.QPoint_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10redirectedEPK12QPaintDeviceP6QPoint", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPaintDeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QPainter_Redirected(device QPaintDevice_ITF /*777 const QPaintDevice **/, offset qtcore.QPoint_ITF /*777 QPoint **/) *QPaintDevice /*777 QPaintDevice **/ {
	var nilthis *QPainter
	rv := nilthis.Redirected(device, offset)
	return rv
}

// /usr/include/qt/QtGui/qpainter.h:465
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void restoreRedirected(const QPaintDevice *)
func (this *QPainter) RestoreRedirected(device QPaintDevice_ITF /*777 const QPaintDevice **/) {
	var convArg0 = device.QPaintDevice_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter17restoreRedirectedEPK12QPaintDevice", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QPainter_RestoreRedirected(device QPaintDevice_ITF /*777 const QPaintDevice **/) {
	var nilthis *QPainter
	nilthis.RestoreRedirected(device)
}

// /usr/include/qt/QtGui/qpainter.h:467
// index:0
// Public Visibility=Default Availability=Available
// [-2] void beginNativePainting()
func (this *QPainter) BeginNativePainting() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter19beginNativePaintingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:468
// index:0
// Public Visibility=Default Availability=Available
// [-2] void endNativePainting()
func (this *QPainter) EndNativePainting() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter17endNativePaintingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

type QPainter__RenderHint = int

const QPainter__Antialiasing QPainter__RenderHint = 1
const QPainter__TextAntialiasing QPainter__RenderHint = 2
const QPainter__SmoothPixmapTransform QPainter__RenderHint = 4
const QPainter__HighQualityAntialiasing QPainter__RenderHint = 8
const QPainter__NonCosmeticDefaultPen QPainter__RenderHint = 16
const QPainter__Qt4CompatiblePainting QPainter__RenderHint = 32

type QPainter__PixmapFragmentHint = int

const QPainter__OpaqueHint QPainter__PixmapFragmentHint = 1

type QPainter__CompositionMode = int

const QPainter__CompositionMode_SourceOver QPainter__CompositionMode = 0
const QPainter__CompositionMode_DestinationOver QPainter__CompositionMode = 1
const QPainter__CompositionMode_Clear QPainter__CompositionMode = 2
const QPainter__CompositionMode_Source QPainter__CompositionMode = 3
const QPainter__CompositionMode_Destination QPainter__CompositionMode = 4
const QPainter__CompositionMode_SourceIn QPainter__CompositionMode = 5
const QPainter__CompositionMode_DestinationIn QPainter__CompositionMode = 6
const QPainter__CompositionMode_SourceOut QPainter__CompositionMode = 7
const QPainter__CompositionMode_DestinationOut QPainter__CompositionMode = 8
const QPainter__CompositionMode_SourceAtop QPainter__CompositionMode = 9
const QPainter__CompositionMode_DestinationAtop QPainter__CompositionMode = 10
const QPainter__CompositionMode_Xor QPainter__CompositionMode = 11
const QPainter__CompositionMode_Plus QPainter__CompositionMode = 12
const QPainter__CompositionMode_Multiply QPainter__CompositionMode = 13
const QPainter__CompositionMode_Screen QPainter__CompositionMode = 14
const QPainter__CompositionMode_Overlay QPainter__CompositionMode = 15
const QPainter__CompositionMode_Darken QPainter__CompositionMode = 16
const QPainter__CompositionMode_Lighten QPainter__CompositionMode = 17
const QPainter__CompositionMode_ColorDodge QPainter__CompositionMode = 18
const QPainter__CompositionMode_ColorBurn QPainter__CompositionMode = 19
const QPainter__CompositionMode_HardLight QPainter__CompositionMode = 20
const QPainter__CompositionMode_SoftLight QPainter__CompositionMode = 21
const QPainter__CompositionMode_Difference QPainter__CompositionMode = 22
const QPainter__CompositionMode_Exclusion QPainter__CompositionMode = 23
const QPainter__RasterOp_SourceOrDestination QPainter__CompositionMode = 24
const QPainter__RasterOp_SourceAndDestination QPainter__CompositionMode = 25
const QPainter__RasterOp_SourceXorDestination QPainter__CompositionMode = 26
const QPainter__RasterOp_NotSourceAndNotDestination QPainter__CompositionMode = 27
const QPainter__RasterOp_NotSourceOrNotDestination QPainter__CompositionMode = 28
const QPainter__RasterOp_NotSourceXorDestination QPainter__CompositionMode = 29
const QPainter__RasterOp_NotSource QPainter__CompositionMode = 30
const QPainter__RasterOp_NotSourceAndDestination QPainter__CompositionMode = 31
const QPainter__RasterOp_SourceAndNotDestination QPainter__CompositionMode = 32
const QPainter__RasterOp_NotSourceOrDestination QPainter__CompositionMode = 33
const QPainter__RasterOp_SourceOrNotDestination QPainter__CompositionMode = 34
const QPainter__RasterOp_ClearDestination QPainter__CompositionMode = 35
const QPainter__RasterOp_SetDestination QPainter__CompositionMode = 36
const QPainter__RasterOp_NotDestination QPainter__CompositionMode = 37

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
