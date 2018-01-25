package qtgui

// /usr/include/qt/QtGui/qpaintengine.h
// #include <qpaintengine.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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
type QPaintEngine struct {
	*qtrt.CObject
}

func (this *QPaintEngine) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPaintEngine) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQPaintEngineFromPointer(cthis unsafe.Pointer) *QPaintEngine {
	return &QPaintEngine{&qtrt.CObject{cthis}}
}
func (*QPaintEngine) NewFromPointer(cthis unsafe.Pointer) *QPaintEngine {
	return NewQPaintEngineFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpaintengine.h:148
// index:0
// Public virtual
// void ~QPaintEngine()
func DeleteQPaintEngine(*QPaintEngine) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngineD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:150
// index:0
// Public inline
// bool isActive()
func (this *QPaintEngine) IsActive() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintEngine8isActiveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpaintengine.h:151
// index:0
// Public inline
// void setActive(_Bool)
func (this *QPaintEngine) SetActive(newState bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine9setActiveEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), newState)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:153
// index:0
// Public pure virtual
// bool begin(class QPaintDevice *)
func (this *QPaintEngine) Begin(pdev *QPaintDevice /*444 QPaintDevice **/) bool {
	var convArg0 = pdev.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine5beginEP12QPaintDevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpaintengine.h:154
// index:0
// Public pure virtual
// bool end()
func (this *QPaintEngine) End() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine3endEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpaintengine.h:156
// index:0
// Public pure virtual
// void updateState(const class QPaintEngineState &)
func (this *QPaintEngine) UpdateState(state *QPaintEngineState) {
	var convArg0 = state.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine11updateStateERK17QPaintEngineState", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:158
// index:0
// Public virtual
// void drawRects(const class QRect *, int)
func (this *QPaintEngine) DrawRects(rects *qtcore.QRect /*444 const QRect **/, rectCount int) {
	var convArg0 = rects.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine9drawRectsEPK5QRecti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, rectCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:159
// index:1
// Public virtual
// void drawRects(const class QRectF *, int)
func (this *QPaintEngine) DrawRects_1(rects *qtcore.QRectF /*444 const QRectF **/, rectCount int) {
	var convArg0 = rects.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine9drawRectsEPK6QRectFi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, rectCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:161
// index:0
// Public virtual
// void drawLines(const class QLine *, int)
func (this *QPaintEngine) DrawLines(lines *qtcore.QLine /*444 const QLine **/, lineCount int) {
	var convArg0 = lines.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine9drawLinesEPK5QLinei", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, lineCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:162
// index:1
// Public virtual
// void drawLines(const class QLineF *, int)
func (this *QPaintEngine) DrawLines_1(lines *qtcore.QLineF /*444 const QLineF **/, lineCount int) {
	var convArg0 = lines.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine9drawLinesEPK6QLineFi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, lineCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:164
// index:0
// Public virtual
// void drawEllipse(const class QRectF &)
func (this *QPaintEngine) DrawEllipse(r *qtcore.QRectF) {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine11drawEllipseERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:165
// index:1
// Public virtual
// void drawEllipse(const class QRect &)
func (this *QPaintEngine) DrawEllipse_1(r *qtcore.QRect) {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine11drawEllipseERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:167
// index:0
// Public virtual
// void drawPath(const class QPainterPath &)
func (this *QPaintEngine) DrawPath(path *QPainterPath) {
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine8drawPathERK12QPainterPath", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:169
// index:0
// Public virtual
// void drawPoints(const class QPointF *, int)
func (this *QPaintEngine) DrawPoints(points *qtcore.QPointF /*444 const QPointF **/, pointCount int) {
	var convArg0 = points.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine10drawPointsEPK7QPointFi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:170
// index:1
// Public virtual
// void drawPoints(const class QPoint *, int)
func (this *QPaintEngine) DrawPoints_1(points *qtcore.QPoint /*444 const QPoint **/, pointCount int) {
	var convArg0 = points.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine10drawPointsEPK6QPointi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:172
// index:0
// Public virtual
// void drawPolygon(const class QPointF *, int, enum QPaintEngine::PolygonDrawMode)
func (this *QPaintEngine) DrawPolygon(points *qtcore.QPointF /*444 const QPointF **/, pointCount int, mode int) {
	var convArg0 = points.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine11drawPolygonEPK7QPointFiNS_15PolygonDrawModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:173
// index:1
// Public virtual
// void drawPolygon(const class QPoint *, int, enum QPaintEngine::PolygonDrawMode)
func (this *QPaintEngine) DrawPolygon_1(points *qtcore.QPoint /*444 const QPoint **/, pointCount int, mode int) {
	var convArg0 = points.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine11drawPolygonEPK6QPointiNS_15PolygonDrawModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:175
// index:0
// Public pure virtual
// void drawPixmap(const class QRectF &, const class QPixmap &, const class QRectF &)
func (this *QPaintEngine) DrawPixmap(r *qtcore.QRectF, pm *QPixmap, sr *qtcore.QRectF) {
	var convArg0 = r.GetCthis()
	var convArg1 = pm.GetCthis()
	var convArg2 = sr.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine10drawPixmapERK6QRectFRK7QPixmapS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:176
// index:0
// Public virtual
// void drawTextItem(const class QPointF &, const class QTextItem &)
func (this *QPaintEngine) DrawTextItem(p *qtcore.QPointF, textItem *QTextItem) {
	var convArg0 = p.GetCthis()
	var convArg1 = textItem.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine12drawTextItemERK7QPointFRK9QTextItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:177
// index:0
// Public virtual
// void drawTiledPixmap(const class QRectF &, const class QPixmap &, const class QPointF &)
func (this *QPaintEngine) DrawTiledPixmap(r *qtcore.QRectF, pixmap *QPixmap, s *qtcore.QPointF) {
	var convArg0 = r.GetCthis()
	var convArg1 = pixmap.GetCthis()
	var convArg2 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine15drawTiledPixmapERK6QRectFRK7QPixmapRK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:178
// index:0
// Public virtual
// void drawImage(const class QRectF &, const class QImage &, const class QRectF &, Qt::ImageConversionFlags)
func (this *QPaintEngine) DrawImage(r *qtcore.QRectF, pm *QImage, sr *qtcore.QRectF, flags int) {
	var convArg0 = r.GetCthis()
	var convArg1 = pm.GetCthis()
	var convArg2 = sr.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine9drawImageERK6QRectFRK6QImageS2_6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:181
// index:0
// Public
// void setPaintDevice(class QPaintDevice *)
func (this *QPaintEngine) SetPaintDevice(device *QPaintDevice /*444 QPaintDevice **/) {
	var convArg0 = device.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine14setPaintDeviceEP12QPaintDevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:182
// index:0
// Public
// QPaintDevice * paintDevice()
func (this *QPaintEngine) PaintDevice() *QPaintDevice /*444 QPaintDevice **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintEngine11paintDeviceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPaintDeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:184
// index:0
// Public
// void setSystemClip(const class QRegion &)
func (this *QPaintEngine) SetSystemClip(baseClip *QRegion) {
	var convArg0 = baseClip.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine13setSystemClipERK7QRegion", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:185
// index:0
// Public
// QRegion systemClip()
func (this *QPaintEngine) SystemClip() *QRegion /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintEngine10systemClipEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:187
// index:0
// Public
// void setSystemRect(const class QRect &)
func (this *QPaintEngine) SetSystemRect(rect *qtcore.QRect) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine13setSystemRectERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:188
// index:0
// Public
// QRect systemRect()
func (this *QPaintEngine) SystemRect() *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintEngine10systemRectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:191
// index:0
// Public virtual
// QPoint coordinateOffset()
func (this *QPaintEngine) CoordinateOffset() *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintEngine16coordinateOffsetEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:214
// index:0
// Public pure virtual
// QPaintEngine::Type type()
func (this *QPaintEngine) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintEngine4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:216
// index:0
// Public inline
// void fix_neg_rect(int *, int *, int *, int *)
func (this *QPaintEngine) Fix_neg_rect(x unsafe.Pointer /*666*/, y unsafe.Pointer /*666*/, w unsafe.Pointer /*666*/, h unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine12fix_neg_rectEPiS0_S0_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:218
// index:0
// Public inline
// bool testDirty(QPaintEngine::DirtyFlags)
func (this *QPaintEngine) TestDirty(df int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine9testDirtyE6QFlagsINS_9DirtyFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), df)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpaintengine.h:219
// index:0
// Public inline
// void setDirty(QPaintEngine::DirtyFlags)
func (this *QPaintEngine) SetDirty(df int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine8setDirtyE6QFlagsINS_9DirtyFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), df)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:220
// index:0
// Public inline
// void clearDirty(QPaintEngine::DirtyFlags)
func (this *QPaintEngine) ClearDirty(df int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine10clearDirtyE6QFlagsINS_9DirtyFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), df)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:224
// index:0
// Public
// QPainter * painter()
func (this *QPaintEngine) Painter() *QPainter /*444 QPainter **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintEngine7painterEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPainterFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:226
// index:0
// Public
// void syncState()
func (this *QPaintEngine) SyncState() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine9syncStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:227
// index:0
// Public inline
// bool isExtended()
func (this *QPaintEngine) IsExtended() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintEngine10isExtendedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
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
const QPaintEngine__AllFeatures QPaintEngine__PaintEngineFeature = 4294967295

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
