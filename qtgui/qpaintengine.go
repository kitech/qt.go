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
// extern C begin: 6
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QPaintEngine struct {
	*qtrt.CObject
}
type QPaintEngine_ITF interface {
	QPaintEngine_PTR() *QPaintEngine
}

func (ptr *QPaintEngine) QPaintEngine_PTR() *QPaintEngine { return ptr }

func (this *QPaintEngine) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPaintEngine) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQPaintEngineFromPointer(cthis unsafe.Pointer) *QPaintEngine {
	return &QPaintEngine{&qtrt.CObject{cthis}}
}
func (*QPaintEngine) NewFromPointer(cthis unsafe.Pointer) *QPaintEngine {
	return NewQPaintEngineFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpaintengine.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPaintEngine(QPaintEngine::PaintEngineFeatures)
func NewQPaintEngine(features int) *QPaintEngine {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngineC2E6QFlagsINS_18PaintEngineFeatureEE", qtrt.FFI_TYPE_POINTER, features)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPaintEngineFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPaintEngine)
	return gothis
}

// /usr/include/qt/QtGui/qpaintengine.h:148
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QPaintEngine()
func DeleteQPaintEngine(this *QPaintEngine) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngineD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpaintengine.h:150
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isActive()
func (this *QPaintEngine) IsActive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintEngine8isActiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpaintengine.h:151
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setActive(_Bool)
func (this *QPaintEngine) SetActive(newState bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngine9setActiveEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newState)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:153
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool begin(QPaintDevice *)
func (this *QPaintEngine) Begin(pdev QPaintDevice_ITF /*777 QPaintDevice **/) bool {
	var convArg0 = pdev.QPaintDevice_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngine5beginEP12QPaintDevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpaintengine.h:154
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool end()
func (this *QPaintEngine) End() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngine3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpaintengine.h:156
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void updateState(const QPaintEngineState &)
func (this *QPaintEngine) UpdateState(state QPaintEngineState_ITF) {
	var convArg0 = state.QPaintEngineState_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngine11updateStateERK17QPaintEngineState", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:158
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawRects(const QRect *, int)
func (this *QPaintEngine) DrawRects(rects qtcore.QRect_ITF /*777 const QRect **/, rectCount int) {
	var convArg0 = rects.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngine9drawRectsEPK5QRecti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, rectCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:159
// index:1
// Public virtual Visibility=Default Availability=Available
// [-2] void drawRects(const QRectF *, int)
func (this *QPaintEngine) DrawRects_1(rects qtcore.QRectF_ITF /*777 const QRectF **/, rectCount int) {
	var convArg0 = rects.QRectF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngine9drawRectsEPK6QRectFi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, rectCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:161
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawLines(const QLine *, int)
func (this *QPaintEngine) DrawLines(lines qtcore.QLine_ITF /*777 const QLine **/, lineCount int) {
	var convArg0 = lines.QLine_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngine9drawLinesEPK5QLinei", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, lineCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:162
// index:1
// Public virtual Visibility=Default Availability=Available
// [-2] void drawLines(const QLineF *, int)
func (this *QPaintEngine) DrawLines_1(lines qtcore.QLineF_ITF /*777 const QLineF **/, lineCount int) {
	var convArg0 = lines.QLineF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngine9drawLinesEPK6QLineFi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, lineCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:164
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawEllipse(const QRectF &)
func (this *QPaintEngine) DrawEllipse(r qtcore.QRectF_ITF) {
	var convArg0 = r.QRectF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngine11drawEllipseERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:165
// index:1
// Public virtual Visibility=Default Availability=Available
// [-2] void drawEllipse(const QRect &)
func (this *QPaintEngine) DrawEllipse_1(r qtcore.QRect_ITF) {
	var convArg0 = r.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngine11drawEllipseERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:167
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawPath(const QPainterPath &)
func (this *QPaintEngine) DrawPath(path QPainterPath_ITF) {
	var convArg0 = path.QPainterPath_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngine8drawPathERK12QPainterPath", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:169
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawPoints(const QPointF *, int)
func (this *QPaintEngine) DrawPoints(points qtcore.QPointF_ITF /*777 const QPointF **/, pointCount int) {
	var convArg0 = points.QPointF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngine10drawPointsEPK7QPointFi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:170
// index:1
// Public virtual Visibility=Default Availability=Available
// [-2] void drawPoints(const QPoint *, int)
func (this *QPaintEngine) DrawPoints_1(points qtcore.QPoint_ITF /*777 const QPoint **/, pointCount int) {
	var convArg0 = points.QPoint_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngine10drawPointsEPK6QPointi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:172
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawPolygon(const QPointF *, int, enum QPaintEngine::PolygonDrawMode)
func (this *QPaintEngine) DrawPolygon(points qtcore.QPointF_ITF /*777 const QPointF **/, pointCount int, mode int) {
	var convArg0 = points.QPointF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngine11drawPolygonEPK7QPointFiNS_15PolygonDrawModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:173
// index:1
// Public virtual Visibility=Default Availability=Available
// [-2] void drawPolygon(const QPoint *, int, enum QPaintEngine::PolygonDrawMode)
func (this *QPaintEngine) DrawPolygon_1(points qtcore.QPoint_ITF /*777 const QPoint **/, pointCount int, mode int) {
	var convArg0 = points.QPoint_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngine11drawPolygonEPK6QPointiNS_15PolygonDrawModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:175
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void drawPixmap(const QRectF &, const QPixmap &, const QRectF &)
func (this *QPaintEngine) DrawPixmap(r qtcore.QRectF_ITF, pm QPixmap_ITF, sr qtcore.QRectF_ITF) {
	var convArg0 = r.QRectF_PTR().GetCthis()
	var convArg1 = pm.QPixmap_PTR().GetCthis()
	var convArg2 = sr.QRectF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngine10drawPixmapERK6QRectFRK7QPixmapS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:176
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawTextItem(const QPointF &, const QTextItem &)
func (this *QPaintEngine) DrawTextItem(p qtcore.QPointF_ITF, textItem QTextItem_ITF) {
	var convArg0 = p.QPointF_PTR().GetCthis()
	var convArg1 = textItem.QTextItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngine12drawTextItemERK7QPointFRK9QTextItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:177
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawTiledPixmap(const QRectF &, const QPixmap &, const QPointF &)
func (this *QPaintEngine) DrawTiledPixmap(r qtcore.QRectF_ITF, pixmap QPixmap_ITF, s qtcore.QPointF_ITF) {
	var convArg0 = r.QRectF_PTR().GetCthis()
	var convArg1 = pixmap.QPixmap_PTR().GetCthis()
	var convArg2 = s.QPointF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngine15drawTiledPixmapERK6QRectFRK7QPixmapRK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:178
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawImage(const QRectF &, const QImage &, const QRectF &, Qt::ImageConversionFlags)
func (this *QPaintEngine) DrawImage(r qtcore.QRectF_ITF, pm QImage_ITF, sr qtcore.QRectF_ITF, flags int) {
	var convArg0 = r.QRectF_PTR().GetCthis()
	var convArg1 = pm.QImage_PTR().GetCthis()
	var convArg2 = sr.QRectF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngine9drawImageERK6QRectFRK6QImageS2_6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:181
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPaintDevice(QPaintDevice *)
func (this *QPaintEngine) SetPaintDevice(device QPaintDevice_ITF /*777 QPaintDevice **/) {
	var convArg0 = device.QPaintDevice_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngine14setPaintDeviceEP12QPaintDevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:182
// index:0
// Public Visibility=Default Availability=Available
// [8] QPaintDevice * paintDevice()
func (this *QPaintEngine) PaintDevice() *QPaintDevice /*777 QPaintDevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintEngine11paintDeviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPaintDeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qpaintengine.h:184
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSystemClip(const QRegion &)
func (this *QPaintEngine) SetSystemClip(baseClip QRegion_ITF) {
	var convArg0 = baseClip.QRegion_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngine13setSystemClipERK7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:185
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion systemClip()
func (this *QPaintEngine) SystemClip() *QRegion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintEngine10systemClipEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSystemRect(const QRect &)
func (this *QPaintEngine) SetSystemRect(rect qtcore.QRect_ITF) {
	var convArg0 = rect.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngine13setSystemRectERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:188
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect systemRect()
func (this *QPaintEngine) SystemRect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintEngine10systemRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:191
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPoint coordinateOffset()
func (this *QPaintEngine) CoordinateOffset() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintEngine16coordinateOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:214
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QPaintEngine::Type type()
func (this *QPaintEngine) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintEngine4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:216
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void fix_neg_rect(int *, int *, int *, int *)
func (this *QPaintEngine) Fix_neg_rect(x unsafe.Pointer /*666*/, y unsafe.Pointer /*666*/, w unsafe.Pointer /*666*/, h unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngine12fix_neg_rectEPiS0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &x, &y, &w, &h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:218
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool testDirty(QPaintEngine::DirtyFlags)
func (this *QPaintEngine) TestDirty(df int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngine9testDirtyE6QFlagsINS_9DirtyFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), df)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpaintengine.h:219
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setDirty(QPaintEngine::DirtyFlags)
func (this *QPaintEngine) SetDirty(df int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngine8setDirtyE6QFlagsINS_9DirtyFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), df)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:220
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void clearDirty(QPaintEngine::DirtyFlags)
func (this *QPaintEngine) ClearDirty(df int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngine10clearDirtyE6QFlagsINS_9DirtyFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), df)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:222
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool hasFeature(QPaintEngine::PaintEngineFeatures)
func (this *QPaintEngine) HasFeature(feature int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintEngine10hasFeatureE6QFlagsINS_18PaintEngineFeatureEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), feature)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpaintengine.h:224
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainter * painter()
func (this *QPaintEngine) Painter() *QPainter /*777 QPainter **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintEngine7painterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPainterFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qpaintengine.h:226
// index:0
// Public Visibility=Default Availability=Available
// [-2] void syncState()
func (this *QPaintEngine) SyncState() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintEngine9syncStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:227
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isExtended()
func (this *QPaintEngine) IsExtended() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintEngine10isExtendedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

type QPaintEngine__PaintEngineFeature = int

const QPaintEngine__PrimitiveTransform QPaintEngine__PaintEngineFeature = 1
const QPaintEngine__PatternTransform QPaintEngine__PaintEngineFeature = 2
const QPaintEngine__PixmapTransform QPaintEngine__PaintEngineFeature = 4
const QPaintEngine__PatternBrush QPaintEngine__PaintEngineFeature = 8
const QPaintEngine__LinearGradientFill QPaintEngine__PaintEngineFeature = 16
const QPaintEngine__RadialGradientFill QPaintEngine__PaintEngineFeature = 32
const QPaintEngine__ConicalGradientFill QPaintEngine__PaintEngineFeature = 64
const QPaintEngine__AlphaBlend QPaintEngine__PaintEngineFeature = 128
const QPaintEngine__PorterDuff QPaintEngine__PaintEngineFeature = 256
const QPaintEngine__PainterPaths QPaintEngine__PaintEngineFeature = 512
const QPaintEngine__Antialiasing QPaintEngine__PaintEngineFeature = 1024
const QPaintEngine__BrushStroke QPaintEngine__PaintEngineFeature = 2048
const QPaintEngine__ConstantOpacity QPaintEngine__PaintEngineFeature = 4096
const QPaintEngine__MaskedBrush QPaintEngine__PaintEngineFeature = 8192
const QPaintEngine__PerspectiveTransform QPaintEngine__PaintEngineFeature = 16384
const QPaintEngine__BlendModes QPaintEngine__PaintEngineFeature = 32768
const QPaintEngine__ObjectBoundingModeGradients QPaintEngine__PaintEngineFeature = 65536
const QPaintEngine__RasterOpModes QPaintEngine__PaintEngineFeature = 131072
const QPaintEngine__PaintOutsidePaintEvent QPaintEngine__PaintEngineFeature = 536870912
const QPaintEngine__AllFeatures QPaintEngine__PaintEngineFeature = -1

type QPaintEngine__DirtyFlag = int

const QPaintEngine__DirtyPen QPaintEngine__DirtyFlag = 1
const QPaintEngine__DirtyBrush QPaintEngine__DirtyFlag = 2
const QPaintEngine__DirtyBrushOrigin QPaintEngine__DirtyFlag = 4
const QPaintEngine__DirtyFont QPaintEngine__DirtyFlag = 8
const QPaintEngine__DirtyBackground QPaintEngine__DirtyFlag = 16
const QPaintEngine__DirtyBackgroundMode QPaintEngine__DirtyFlag = 32
const QPaintEngine__DirtyTransform QPaintEngine__DirtyFlag = 64
const QPaintEngine__DirtyClipRegion QPaintEngine__DirtyFlag = 128
const QPaintEngine__DirtyClipPath QPaintEngine__DirtyFlag = 256
const QPaintEngine__DirtyHints QPaintEngine__DirtyFlag = 512
const QPaintEngine__DirtyCompositionMode QPaintEngine__DirtyFlag = 1024
const QPaintEngine__DirtyClipEnabled QPaintEngine__DirtyFlag = 2048
const QPaintEngine__DirtyOpacity QPaintEngine__DirtyFlag = 4096
const QPaintEngine__AllDirty QPaintEngine__DirtyFlag = 65535

type QPaintEngine__PolygonDrawMode = int

const QPaintEngine__OddEvenMode QPaintEngine__PolygonDrawMode = 0
const QPaintEngine__WindingMode QPaintEngine__PolygonDrawMode = 1
const QPaintEngine__ConvexMode QPaintEngine__PolygonDrawMode = 2
const QPaintEngine__PolylineMode QPaintEngine__PolygonDrawMode = 3

type QPaintEngine__Type = int

const QPaintEngine__X11 QPaintEngine__Type = 0
const QPaintEngine__Windows QPaintEngine__Type = 1
const QPaintEngine__QuickDraw QPaintEngine__Type = 2
const QPaintEngine__CoreGraphics QPaintEngine__Type = 3
const QPaintEngine__MacPrinter QPaintEngine__Type = 4
const QPaintEngine__QWindowSystem QPaintEngine__Type = 5
const QPaintEngine__PostScript QPaintEngine__Type = 6
const QPaintEngine__OpenGL QPaintEngine__Type = 7
const QPaintEngine__Picture QPaintEngine__Type = 8
const QPaintEngine__SVG QPaintEngine__Type = 9
const QPaintEngine__Raster QPaintEngine__Type = 10
const QPaintEngine__Direct3D QPaintEngine__Type = 11
const QPaintEngine__Pdf QPaintEngine__Type = 12
const QPaintEngine__OpenVG QPaintEngine__Type = 13
const QPaintEngine__OpenGL2 QPaintEngine__Type = 14
const QPaintEngine__PaintBuffer QPaintEngine__Type = 15
const QPaintEngine__Blitter QPaintEngine__Type = 16
const QPaintEngine__Direct2D QPaintEngine__Type = 17
const QPaintEngine__User QPaintEngine__Type = 50
const QPaintEngine__MaxUser QPaintEngine__Type = 100

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
