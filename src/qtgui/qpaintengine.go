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
func NewQPaintEngineFromPointer(cthis unsafe.Pointer) *QPaintEngine {
	return &QPaintEngine{&qtrt.CObject{cthis}}
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine9setActiveEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &newState)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine9drawRectsEPK5QRecti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &rectCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:159
// index:1
// Public virtual
// void drawRects(const class QRectF *, int)
func (this *QPaintEngine) DrawRects_1(rects *qtcore.QRectF /*444 const QRectF **/, rectCount int) {
	var convArg0 = rects.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine9drawRectsEPK6QRectFi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &rectCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:161
// index:0
// Public virtual
// void drawLines(const class QLine *, int)
func (this *QPaintEngine) DrawLines(lines *qtcore.QLine /*444 const QLine **/, lineCount int) {
	var convArg0 = lines.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine9drawLinesEPK5QLinei", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &lineCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:162
// index:1
// Public virtual
// void drawLines(const class QLineF *, int)
func (this *QPaintEngine) DrawLines_1(lines *qtcore.QLineF /*444 const QLineF **/, lineCount int) {
	var convArg0 = lines.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine9drawLinesEPK6QLineFi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &lineCount)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine10drawPointsEPK7QPointFi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &pointCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:170
// index:1
// Public virtual
// void drawPoints(const class QPoint *, int)
func (this *QPaintEngine) DrawPoints_1(points *qtcore.QPoint /*444 const QPoint **/, pointCount int) {
	var convArg0 = points.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine10drawPointsEPK6QPointi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &pointCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:172
// index:0
// Public virtual
// void drawPolygon(const class QPointF *, int, enum QPaintEngine::PolygonDrawMode)
func (this *QPaintEngine) DrawPolygon(points *qtcore.QPointF /*444 const QPointF **/, pointCount int, mode int) {
	var convArg0 = points.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine11drawPolygonEPK7QPointFiNS_15PolygonDrawModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &pointCount, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:173
// index:1
// Public virtual
// void drawPolygon(const class QPoint *, int, enum QPaintEngine::PolygonDrawMode)
func (this *QPaintEngine) DrawPolygon_1(points *qtcore.QPoint /*444 const QPoint **/, pointCount int, mode int) {
	var convArg0 = points.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine11drawPolygonEPK6QPointiNS_15PolygonDrawModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &pointCount, &mode)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintEngine10systemClipEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintEngine10systemRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:191
// index:0
// Public virtual
// QPoint coordinateOffset()
func (this *QPaintEngine) CoordinateOffset() *qtcore.QPoint /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintEngine16coordinateOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine12fix_neg_rectEPiS0_S0_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
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

//  body block end
