//  header block begin
// /usr/include/qt/QtGui/qpaintengine.h
// #include <qpaintengine.h>
// #include <QtGui>
package qtgui

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
	return this.Cthis
}

// /usr/include/qt/QtGui/qpaintengine.h:147
// index:0
// void QPaintEngine(QPaintEngine::PaintEngineFeatures)
func NewQPaintEngine(features int) *QPaintEngine {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngineC1E6QFlagsINS_18PaintEngineFeatureEE", ffiqt.FFI_TYPE_VOID, cthis, features)
	gopp.ErrPrint(err, rv)
	gothis := NewQPaintEngineFromPointer(cthis)
	return gothis
}
func NewQPaintEngineFromPointer(cthis unsafe.Pointer) *QPaintEngine {
	return &QPaintEngine{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qpaintengine.h:230
// index:1
// void QPaintEngine(class QPaintEnginePrivate &, QPaintEngine::PaintEngineFeatures)
func NewQPaintEngine_1(data unsafe.Pointer, devcaps int) *QPaintEngine {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngineC1ER19QPaintEnginePrivate6QFlagsINS_18PaintEngineFeatureEE", ffiqt.FFI_TYPE_VOID, cthis, data, devcaps)
	gopp.ErrPrint(err, rv)
	gothis := NewQPaintEngineFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpaintengine.h:148
// index:0
// virtual
// void ~QPaintEngine()
func DeleteQPaintEngine(*QPaintEngine) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngineD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:150
// index:0
// inline
// bool isActive()
func (this *QPaintEngine) IsActive() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintEngine8isActiveEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:151
// index:0
// inline
// void setActive(_Bool)
func (this *QPaintEngine) SetActive(newState bool) {
	// 0: (, newState bool), (&newState)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine9setActiveEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &newState)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:153
// index:0
// pure virtual
// bool begin(class QPaintDevice *)
func (this *QPaintEngine) Begin(pdev unsafe.Pointer) {
	// 0: (, pdev QPaintDevice *), (pdev)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine5beginEP12QPaintDevice", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pdev)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:154
// index:0
// pure virtual
// bool end()
func (this *QPaintEngine) End() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine3endEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:156
// index:0
// pure virtual
// void updateState(const class QPaintEngineState &)
func (this *QPaintEngine) UpdateState(state unsafe.Pointer) {
	// 0: (, state const QPaintEngineState &), (state)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine11updateStateERK17QPaintEngineState", ffiqt.FFI_TYPE_VOID, this.GetCthis(), state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:158
// index:0
// virtual
// void drawRects(const class QRect *, int)
func (this *QPaintEngine) DrawRects(rects unsafe.Pointer, rectCount int) {
	// 0: (, rects const QRect *, rectCount int), (rects, &rectCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine9drawRectsEPK5QRecti", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rects, &rectCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:159
// index:1
// virtual
// void drawRects(const class QRectF *, int)
func (this *QPaintEngine) DrawRects_1(rects unsafe.Pointer, rectCount int) {
	// 1: (, rects const QRectF *, rectCount int), (rects, &rectCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine9drawRectsEPK6QRectFi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rects, &rectCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:161
// index:0
// virtual
// void drawLines(const class QLine *, int)
func (this *QPaintEngine) DrawLines(lines unsafe.Pointer, lineCount int) {
	// 0: (, lines const QLine *, lineCount int), (lines, &lineCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine9drawLinesEPK5QLinei", ffiqt.FFI_TYPE_VOID, this.GetCthis(), lines, &lineCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:162
// index:1
// virtual
// void drawLines(const class QLineF *, int)
func (this *QPaintEngine) DrawLines_1(lines unsafe.Pointer, lineCount int) {
	// 1: (, lines const QLineF *, lineCount int), (lines, &lineCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine9drawLinesEPK6QLineFi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), lines, &lineCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:164
// index:0
// virtual
// void drawEllipse(const class QRectF &)
func (this *QPaintEngine) DrawEllipse(r unsafe.Pointer) {
	// 0: (, r const QRectF &), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine11drawEllipseERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:165
// index:1
// virtual
// void drawEllipse(const class QRect &)
func (this *QPaintEngine) DrawEllipse_1(r unsafe.Pointer) {
	// 1: (, r const QRect &), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine11drawEllipseERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:167
// index:0
// virtual
// void drawPath(const class QPainterPath &)
func (this *QPaintEngine) DrawPath(path unsafe.Pointer) {
	// 0: (, path const QPainterPath &), (path)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine8drawPathERK12QPainterPath", ffiqt.FFI_TYPE_VOID, this.GetCthis(), path)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:169
// index:0
// virtual
// void drawPoints(const class QPointF *, int)
func (this *QPaintEngine) DrawPoints(points unsafe.Pointer, pointCount int) {
	// 0: (, points const QPointF *, pointCount int), (points, &pointCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine10drawPointsEPK7QPointFi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), points, &pointCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:170
// index:1
// virtual
// void drawPoints(const class QPoint *, int)
func (this *QPaintEngine) DrawPoints_1(points unsafe.Pointer, pointCount int) {
	// 1: (, points const QPoint *, pointCount int), (points, &pointCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine10drawPointsEPK6QPointi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), points, &pointCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:172
// index:0
// virtual
// void drawPolygon(const class QPointF *, int, enum QPaintEngine::PolygonDrawMode)
func (this *QPaintEngine) DrawPolygon(points unsafe.Pointer, pointCount int, mode int) {
	// 0: (, points const QPointF *, pointCount int, mode QPaintEngine::PolygonDrawMode), (points, &pointCount, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine11drawPolygonEPK7QPointFiNS_15PolygonDrawModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), points, &pointCount, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:173
// index:1
// virtual
// void drawPolygon(const class QPoint *, int, enum QPaintEngine::PolygonDrawMode)
func (this *QPaintEngine) DrawPolygon_1(points unsafe.Pointer, pointCount int, mode int) {
	// 1: (, points const QPoint *, pointCount int, mode QPaintEngine::PolygonDrawMode), (points, &pointCount, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine11drawPolygonEPK6QPointiNS_15PolygonDrawModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), points, &pointCount, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:175
// index:0
// pure virtual
// void drawPixmap(const class QRectF &, const class QPixmap &, const class QRectF &)
func (this *QPaintEngine) DrawPixmap(r unsafe.Pointer, pm unsafe.Pointer, sr unsafe.Pointer) {
	// 0: (, r const QRectF &, pm const QPixmap &, sr const QRectF &), (r, pm, sr)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine10drawPixmapERK6QRectFRK7QPixmapS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r, pm, sr)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:176
// index:0
// virtual
// void drawTextItem(const class QPointF &, const class QTextItem &)
func (this *QPaintEngine) DrawTextItem(p unsafe.Pointer, textItem unsafe.Pointer) {
	// 0: (, p const QPointF &, textItem const QTextItem &), (p, textItem)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine12drawTextItemERK7QPointFRK9QTextItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p, textItem)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:177
// index:0
// virtual
// void drawTiledPixmap(const class QRectF &, const class QPixmap &, const class QPointF &)
func (this *QPaintEngine) DrawTiledPixmap(r unsafe.Pointer, pixmap unsafe.Pointer, s unsafe.Pointer) {
	// 0: (, r const QRectF &, pixmap const QPixmap &, s const QPointF &), (r, pixmap, s)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine15drawTiledPixmapERK6QRectFRK7QPixmapRK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r, pixmap, s)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:178
// index:0
// virtual
// void drawImage(const class QRectF &, const class QImage &, const class QRectF &, Qt::ImageConversionFlags)
func (this *QPaintEngine) DrawImage(r unsafe.Pointer, pm unsafe.Pointer, sr unsafe.Pointer, flags int) {
	// 0: (, r const QRectF &, pm const QImage &, sr const QRectF &, QFlags<Qt::ImageConversionFlag> flags), (r, pm, sr, &flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine9drawImageERK6QRectFRK6QImageS2_6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r, pm, sr, &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:181
// index:0
// void setPaintDevice(class QPaintDevice *)
func (this *QPaintEngine) SetPaintDevice(device unsafe.Pointer) {
	// 0: (, device QPaintDevice *), (device)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine14setPaintDeviceEP12QPaintDevice", ffiqt.FFI_TYPE_VOID, this.GetCthis(), device)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:182
// index:0
// QPaintDevice * paintDevice()
func (this *QPaintEngine) PaintDevice() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintEngine11paintDeviceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:184
// index:0
// void setSystemClip(const class QRegion &)
func (this *QPaintEngine) SetSystemClip(baseClip unsafe.Pointer) {
	// 0: (, baseClip const QRegion &), (baseClip)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine13setSystemClipERK7QRegion", ffiqt.FFI_TYPE_VOID, this.GetCthis(), baseClip)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:185
// index:0
// QRegion systemClip()
func (this *QPaintEngine) SystemClip() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintEngine10systemClipEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:187
// index:0
// void setSystemRect(const class QRect &)
func (this *QPaintEngine) SetSystemRect(rect unsafe.Pointer) {
	// 0: (, rect const QRect &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine13setSystemRectERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:188
// index:0
// QRect systemRect()
func (this *QPaintEngine) SystemRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintEngine10systemRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:191
// index:0
// virtual
// QPoint coordinateOffset()
func (this *QPaintEngine) CoordinateOffset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintEngine16coordinateOffsetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:214
// index:0
// pure virtual
// QPaintEngine::Type type()
func (this *QPaintEngine) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintEngine4typeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:216
// index:0
// inline
// void fix_neg_rect(int *, int *, int *, int *)
func (this *QPaintEngine) Fix_neg_rect(x unsafe.Pointer, y unsafe.Pointer, w unsafe.Pointer, h unsafe.Pointer) {
	// 0: (, x int *, y int *, w int *, h int *), (x, y, w, h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine12fix_neg_rectEPiS0_S0_S0_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:218
// index:0
// inline
// bool testDirty(QPaintEngine::DirtyFlags)
func (this *QPaintEngine) TestDirty(df int) {
	// 0: (, QFlags<QPaintEngine::DirtyFlag> df), (df)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine9testDirtyE6QFlagsINS_9DirtyFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), df)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:219
// index:0
// inline
// void setDirty(QPaintEngine::DirtyFlags)
func (this *QPaintEngine) SetDirty(df int) {
	// 0: (, QFlags<QPaintEngine::DirtyFlag> df), (df)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine8setDirtyE6QFlagsINS_9DirtyFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), df)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:220
// index:0
// inline
// void clearDirty(QPaintEngine::DirtyFlags)
func (this *QPaintEngine) ClearDirty(df int) {
	// 0: (, QFlags<QPaintEngine::DirtyFlag> df), (df)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine10clearDirtyE6QFlagsINS_9DirtyFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), df)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:222
// index:0
// inline
// bool hasFeature(QPaintEngine::PaintEngineFeatures)
func (this *QPaintEngine) HasFeature(feature int) {
	// 0: (, QFlags<QPaintEngine::PaintEngineFeature> feature), (feature)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintEngine10hasFeatureE6QFlagsINS_18PaintEngineFeatureEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), feature)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:224
// index:0
// QPainter * painter()
func (this *QPaintEngine) Painter() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintEngine7painterEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:226
// index:0
// void syncState()
func (this *QPaintEngine) SyncState() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintEngine9syncStateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:227
// index:0
// inline
// bool isExtended()
func (this *QPaintEngine) IsExtended() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintEngine10isExtendedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
