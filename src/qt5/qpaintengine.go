package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtGui/qpaintengine.h
// dst-file: /src/gui/qpaintengine.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  qreal QTextItem::descent();
extern void _ZNK9QTextItem7descentEv(void* qthis);
  // proto:  qreal QTextItem::width();
extern void _ZNK9QTextItem5widthEv(void* qthis);
  // proto:  QFont QTextItem::font();
extern void _ZNK9QTextItem4fontEv(void* qthis);
  // proto:  qreal QTextItem::ascent();
extern void _ZNK9QTextItem6ascentEv(void* qthis);
  // proto:  QString QTextItem::text();
extern void _ZNK9QTextItem4textEv(void* qthis);
  // proto:  qreal QPaintEngineState::opacity();
extern void _ZNK17QPaintEngineState7opacityEv(void* qthis);
  // proto:  QMatrix QPaintEngineState::matrix();
extern void _ZNK17QPaintEngineState6matrixEv(void* qthis);
  // proto:  QPainter * QPaintEngineState::painter();
extern void _ZNK17QPaintEngineState7painterEv(void* qthis);
  // proto:  QTransform QPaintEngineState::transform();
extern void _ZNK17QPaintEngineState9transformEv(void* qthis);
  // proto:  QPointF QPaintEngineState::brushOrigin();
extern void _ZNK17QPaintEngineState11brushOriginEv(void* qthis);
  // proto:  bool QPaintEngineState::penNeedsResolving();
extern void _ZNK17QPaintEngineState17penNeedsResolvingEv(void* qthis);
  // proto:  bool QPaintEngineState::isClipEnabled();
extern void _ZNK17QPaintEngineState13isClipEnabledEv(void* qthis);
  // proto:  QFont QPaintEngineState::font();
extern void _ZNK17QPaintEngineState4fontEv(void* qthis);
  // proto:  bool QPaintEngineState::brushNeedsResolving();
extern void _ZNK17QPaintEngineState19brushNeedsResolvingEv(void* qthis);
  // proto:  QRegion QPaintEngineState::clipRegion();
extern void _ZNK17QPaintEngineState10clipRegionEv(void* qthis);
  // proto:  QPainterPath QPaintEngineState::clipPath();
extern void _ZNK17QPaintEngineState8clipPathEv(void* qthis);
  // proto:  QBrush QPaintEngineState::brush();
extern void _ZNK17QPaintEngineState5brushEv(void* qthis);
  // proto:  QPen QPaintEngineState::pen();
extern void _ZNK17QPaintEngineState3penEv(void* qthis);
  // proto:  QBrush QPaintEngineState::backgroundBrush();
extern void _ZNK17QPaintEngineState15backgroundBrushEv(void* qthis);
  // proto:  void QPaintEngine::drawEllipse(const QRect & r);
extern void _ZN12QPaintEngine11drawEllipseERK5QRect(void* qthis, void* arg0);
  // proto:  void QPaintEngine::QPaintEngine(const QPaintEngine & );
extern void* dector_ZN12QPaintEngineC1ERKS_(void* arg0);
extern void _ZN12QPaintEngineC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QPaintEngine::isActive();
extern void _ZNK12QPaintEngine8isActiveEv(void* qthis);
  // proto:  void QPaintEngine::drawPoints(const QPointF * points, int pointCount);
extern void _ZN12QPaintEngine10drawPointsEPK7QPointFi(void* qthis, void* arg0, int arg1);
  // proto:  QPoint QPaintEngine::coordinateOffset();
extern void _ZNK12QPaintEngine16coordinateOffsetEv(void* qthis);
  // proto:  void QPaintEngine::setPaintDevice(QPaintDevice * device);
extern void _ZN12QPaintEngine14setPaintDeviceEP12QPaintDevice(void* qthis, void* arg0);
  // proto:  void QPaintEngine::setSystemRect(const QRect & rect);
extern void _ZN12QPaintEngine13setSystemRectERK5QRect(void* qthis, void* arg0);
  // proto:  void QPaintEngine::~QPaintEngine();
extern void _ZN12QPaintEngineD0Ev(void* qthis);
  // proto:  bool QPaintEngine::end();
extern void _ZN12QPaintEngine3endEv(void* qthis);
  // proto:  void QPaintEngine::drawTiledPixmap(const QRectF & r, const QPixmap & pixmap, const QPointF & s);
extern void _ZN12QPaintEngine15drawTiledPixmapERK6QRectFRK7QPixmapRK7QPointF(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QPaintEngine::setActive(bool newState);
extern void _ZN12QPaintEngine9setActiveEb(void* qthis, bool arg0);
  // proto:  void QPaintEngine::drawPixmap(const QRectF & r, const QPixmap & pm, const QRectF & sr);
extern void _ZN12QPaintEngine10drawPixmapERK6QRectFRK7QPixmapS2_(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QPaintEngine::drawLines(const QLine * lines, int lineCount);
extern void _ZN12QPaintEngine9drawLinesEPK5QLinei(void* qthis, void* arg0, int arg1);
  // proto:  void QPaintEngine::drawPath(const QPainterPath & path);
extern void _ZN12QPaintEngine8drawPathERK12QPainterPath(void* qthis, void* arg0);
  // proto:  void QPaintEngine::drawLines(const QLineF * lines, int lineCount);
extern void _ZN12QPaintEngine9drawLinesEPK6QLineFi(void* qthis, void* arg0, int arg1);
  // proto:  void QPaintEngine::updateState(const QPaintEngineState & state);
extern void _ZN12QPaintEngine11updateStateERK17QPaintEngineState(void* qthis, void* arg0);
  // proto:  bool QPaintEngine::begin(QPaintDevice * pdev);
extern void _ZN12QPaintEngine5beginEP12QPaintDevice(void* qthis, void* arg0);
  // proto:  QRect QPaintEngine::systemRect();
extern void _ZNK12QPaintEngine10systemRectEv(void* qthis);
  // proto:  void QPaintEngine::drawRects(const QRectF * rects, int rectCount);
extern void _ZN12QPaintEngine9drawRectsEPK6QRectFi(void* qthis, void* arg0, int arg1);
  // proto:  void QPaintEngine::setSystemClip(const QRegion & baseClip);
extern void _ZN12QPaintEngine13setSystemClipERK7QRegion(void* qthis, void* arg0);
  // proto:  QRegion QPaintEngine::systemClip();
extern void _ZNK12QPaintEngine10systemClipEv(void* qthis);
  // proto:  QPaintDevice * QPaintEngine::paintDevice();
extern void _ZNK12QPaintEngine11paintDeviceEv(void* qthis);
  // proto:  void QPaintEngine::syncState();
extern void _ZN12QPaintEngine9syncStateEv(void* qthis);
  // proto:  QPainter * QPaintEngine::painter();
extern void _ZNK12QPaintEngine7painterEv(void* qthis);
  // proto:  void QPaintEngine::drawEllipse(const QRectF & r);
extern void _ZN12QPaintEngine11drawEllipseERK6QRectF(void* qthis, void* arg0);
  // proto:  void QPaintEngine::drawTextItem(const QPointF & p, const QTextItem & textItem);
extern void _ZN12QPaintEngine12drawTextItemERK7QPointFRK9QTextItem(void* qthis, void* arg0, void* arg1);
  // proto:  void QPaintEngine::fix_neg_rect(int * x, int * y, int * w, int * h);
extern void demth_ZN12QPaintEngine12fix_neg_rectEPiS0_S0_S0_(void* qthis, int* arg0, int* arg1, int* arg2, int* arg3);
  // proto:  bool QPaintEngine::isExtended();
extern void demth_ZNK12QPaintEngine10isExtendedEv(void* qthis);
  // proto:  void QPaintEngine::drawRects(const QRect * rects, int rectCount);
extern void _ZN12QPaintEngine9drawRectsEPK5QRecti(void* qthis, void* arg0, int arg1);
  // proto:  void QPaintEngine::drawPoints(const QPoint * points, int pointCount);
extern void _ZN12QPaintEngine10drawPointsEPK6QPointi(void* qthis, void* arg0, int arg1);
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QTextItem)=1
type QTextItem struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QPaintEngineState)=1
type QPaintEngineState struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QPaintEngine)=1
type QPaintEngine struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  qreal QTextItem::descent();
func (this *QTextItem) descent(args ...interface{}) () {
  // descent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextItem7descentEv
  default:
    qtrt.ErrorResolve("QTextItem", "descent", args)
  }

}

  // proto:  qreal QTextItem::width();
func (this *QTextItem) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextItem5widthEv
  default:
    qtrt.ErrorResolve("QTextItem", "width", args)
  }

}

  // proto:  QFont QTextItem::font();
func (this *QTextItem) font(args ...interface{}) () {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextItem4fontEv
  default:
    qtrt.ErrorResolve("QTextItem", "font", args)
  }

}

  // proto:  qreal QTextItem::ascent();
func (this *QTextItem) ascent(args ...interface{}) () {
  // ascent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextItem6ascentEv
  default:
    qtrt.ErrorResolve("QTextItem", "ascent", args)
  }

}

  // proto:  QString QTextItem::text();
func (this *QTextItem) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextItem4textEv
  default:
    qtrt.ErrorResolve("QTextItem", "text", args)
  }

}

  // proto:  qreal QPaintEngineState::opacity();
func (this *QPaintEngineState) opacity(args ...interface{}) () {
  // opacity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState7opacityEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "opacity", args)
  }

}

  // proto:  QMatrix QPaintEngineState::matrix();
func (this *QPaintEngineState) matrix(args ...interface{}) () {
  // matrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState6matrixEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "matrix", args)
  }

}

  // proto:  QPainter * QPaintEngineState::painter();
func (this *QPaintEngineState) painter(args ...interface{}) () {
  // painter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState7painterEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "painter", args)
  }

}

  // proto:  QTransform QPaintEngineState::transform();
func (this *QPaintEngineState) transform(args ...interface{}) () {
  // transform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState9transformEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "transform", args)
  }

}

  // proto:  QPointF QPaintEngineState::brushOrigin();
func (this *QPaintEngineState) brushOrigin(args ...interface{}) () {
  // brushOrigin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState11brushOriginEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "brushOrigin", args)
  }

}

  // proto:  bool QPaintEngineState::penNeedsResolving();
func (this *QPaintEngineState) penNeedsResolving(args ...interface{}) () {
  // penNeedsResolving()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState17penNeedsResolvingEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "penNeedsResolving", args)
  }

}

  // proto:  bool QPaintEngineState::isClipEnabled();
func (this *QPaintEngineState) isClipEnabled(args ...interface{}) () {
  // isClipEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState13isClipEnabledEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "isClipEnabled", args)
  }

}

  // proto:  QFont QPaintEngineState::font();
func (this *QPaintEngineState) font(args ...interface{}) () {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState4fontEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "font", args)
  }

}

  // proto:  bool QPaintEngineState::brushNeedsResolving();
func (this *QPaintEngineState) brushNeedsResolving(args ...interface{}) () {
  // brushNeedsResolving()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState19brushNeedsResolvingEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "brushNeedsResolving", args)
  }

}

  // proto:  QRegion QPaintEngineState::clipRegion();
func (this *QPaintEngineState) clipRegion(args ...interface{}) () {
  // clipRegion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState10clipRegionEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "clipRegion", args)
  }

}

  // proto:  QPainterPath QPaintEngineState::clipPath();
func (this *QPaintEngineState) clipPath(args ...interface{}) () {
  // clipPath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState8clipPathEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "clipPath", args)
  }

}

  // proto:  QBrush QPaintEngineState::brush();
func (this *QPaintEngineState) brush(args ...interface{}) () {
  // brush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState5brushEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "brush", args)
  }

}

  // proto:  QPen QPaintEngineState::pen();
func (this *QPaintEngineState) pen(args ...interface{}) () {
  // pen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState3penEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "pen", args)
  }

}

  // proto:  QBrush QPaintEngineState::backgroundBrush();
func (this *QPaintEngineState) backgroundBrush(args ...interface{}) () {
  // backgroundBrush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState15backgroundBrushEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "backgroundBrush", args)
  }

}

  // proto:  void QPaintEngine::drawEllipse(const QRect & r);
func (this *QPaintEngine) drawEllipse(args ...interface{}) () {
  // drawEllipse(const class QRect &)
  // drawEllipse(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine11drawEllipseERK5QRect
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN12QPaintEngine11drawEllipseERK6QRectF
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawEllipse", args)
  }

}

  // proto:  void QPaintEngine::QPaintEngine(const QPaintEngine & );
func NewQPaintEngine(args ...interface{}) QPaintEngine {
  return QPaintEngine{}
}

  // proto:  bool QPaintEngine::isActive();
func (this *QPaintEngine) isActive(args ...interface{}) () {
  // isActive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintEngine8isActiveEv
  default:
    qtrt.ErrorResolve("QPaintEngine", "isActive", args)
  }

}

  // proto:  void QPaintEngine::drawPoints(const QPointF * points, int pointCount);
func (this *QPaintEngine) drawPoints(args ...interface{}) () {
  // drawPoints(const class QPointF *, int)
  // drawPoints(const class QPoint *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine10drawPointsEPK7QPointFi
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN12QPaintEngine10drawPointsEPK6QPointi
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawPoints", args)
  }

}

  // proto:  QPoint QPaintEngine::coordinateOffset();
func (this *QPaintEngine) coordinateOffset(args ...interface{}) () {
  // coordinateOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintEngine16coordinateOffsetEv
  default:
    qtrt.ErrorResolve("QPaintEngine", "coordinateOffset", args)
  }

}

  // proto:  void QPaintEngine::setPaintDevice(QPaintDevice * device);
func (this *QPaintEngine) setPaintDevice(args ...interface{}) () {
  // setPaintDevice(class QPaintDevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPaintDevice{}) // "QPaintDevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine14setPaintDeviceEP12QPaintDevice
    var arg0 = args[0].(QPaintDevice).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPaintEngine", "setPaintDevice", args)
  }

}

  // proto:  void QPaintEngine::setSystemRect(const QRect & rect);
func (this *QPaintEngine) setSystemRect(args ...interface{}) () {
  // setSystemRect(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine13setSystemRectERK5QRect
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPaintEngine", "setSystemRect", args)
  }

}

  // proto:  void QPaintEngine::~QPaintEngine();
func (this *QPaintEngine) FreeQPaintEngine(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPaintEngine", "~QPaintEngine", args)
  }

}

  // proto:  bool QPaintEngine::end();
func (this *QPaintEngine) end(args ...interface{}) () {
  // end()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine3endEv
  default:
    qtrt.ErrorResolve("QPaintEngine", "end", args)
  }

}

  // proto:  void QPaintEngine::drawTiledPixmap(const QRectF & r, const QPixmap & pixmap, const QPointF & s);
func (this *QPaintEngine) drawTiledPixmap(args ...interface{}) () {
  // drawTiledPixmap(const class QRectF &, const class QPixmap &, const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[0][2] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine15drawTiledPixmapERK6QRectFRK7QPixmapRK7QPointF
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPixmap).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QPointF).qclsinst
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawTiledPixmap", args)
  }

}

  // proto:  void QPaintEngine::setActive(bool newState);
func (this *QPaintEngine) setActive(args ...interface{}) () {
  // setActive(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine9setActiveEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPaintEngine", "setActive", args)
  }

}

  // proto:  void QPaintEngine::drawPixmap(const QRectF & r, const QPixmap & pm, const QRectF & sr);
func (this *QPaintEngine) drawPixmap(args ...interface{}) () {
  // drawPixmap(const class QRectF &, const class QPixmap &, const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[0][2] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine10drawPixmapERK6QRectFRK7QPixmapS2_
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPixmap).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QRectF).qclsinst
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawPixmap", args)
  }

}

  // proto:  void QPaintEngine::drawLines(const QLine * lines, int lineCount);
func (this *QPaintEngine) drawLines(args ...interface{}) () {
  // drawLines(const class QLine *, int)
  // drawLines(const class QLineF *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLine{}) // "const QLine *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QLineF{}) // "const QLineF *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine9drawLinesEPK5QLinei
    var arg0 = args[0].(QLine).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN12QPaintEngine9drawLinesEPK6QLineFi
    var arg0 = args[0].(QLineF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawLines", args)
  }

}

  // proto:  void QPaintEngine::drawPath(const QPainterPath & path);
func (this *QPaintEngine) drawPath(args ...interface{}) () {
  // drawPath(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine8drawPathERK12QPainterPath
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawPath", args)
  }

}

  // proto:  void QPaintEngine::updateState(const QPaintEngineState & state);
func (this *QPaintEngine) updateState(args ...interface{}) () {
  // updateState(const class QPaintEngineState &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPaintEngineState{}) // "const QPaintEngineState &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine11updateStateERK17QPaintEngineState
    var arg0 = args[0].(QPaintEngineState).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPaintEngine", "updateState", args)
  }

}

  // proto:  bool QPaintEngine::begin(QPaintDevice * pdev);
func (this *QPaintEngine) begin(args ...interface{}) () {
  // begin(class QPaintDevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPaintDevice{}) // "QPaintDevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine5beginEP12QPaintDevice
    var arg0 = args[0].(QPaintDevice).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPaintEngine", "begin", args)
  }

}

  // proto:  QRect QPaintEngine::systemRect();
func (this *QPaintEngine) systemRect(args ...interface{}) () {
  // systemRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintEngine10systemRectEv
  default:
    qtrt.ErrorResolve("QPaintEngine", "systemRect", args)
  }

}

  // proto:  void QPaintEngine::drawRects(const QRectF * rects, int rectCount);
func (this *QPaintEngine) drawRects(args ...interface{}) () {
  // drawRects(const class QRectF *, int)
  // drawRects(const class QRect *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine9drawRectsEPK6QRectFi
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN12QPaintEngine9drawRectsEPK5QRecti
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawRects", args)
  }

}

  // proto:  void QPaintEngine::setSystemClip(const QRegion & baseClip);
func (this *QPaintEngine) setSystemClip(args ...interface{}) () {
  // setSystemClip(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine13setSystemClipERK7QRegion
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPaintEngine", "setSystemClip", args)
  }

}

  // proto:  QRegion QPaintEngine::systemClip();
func (this *QPaintEngine) systemClip(args ...interface{}) () {
  // systemClip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintEngine10systemClipEv
  default:
    qtrt.ErrorResolve("QPaintEngine", "systemClip", args)
  }

}

  // proto:  QPaintDevice * QPaintEngine::paintDevice();
func (this *QPaintEngine) paintDevice(args ...interface{}) () {
  // paintDevice()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintEngine11paintDeviceEv
  default:
    qtrt.ErrorResolve("QPaintEngine", "paintDevice", args)
  }

}

  // proto:  void QPaintEngine::syncState();
func (this *QPaintEngine) syncState(args ...interface{}) () {
  // syncState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine9syncStateEv
  default:
    qtrt.ErrorResolve("QPaintEngine", "syncState", args)
  }

}

  // proto:  QPainter * QPaintEngine::painter();
func (this *QPaintEngine) painter(args ...interface{}) () {
  // painter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintEngine7painterEv
  default:
    qtrt.ErrorResolve("QPaintEngine", "painter", args)
  }

}

  // proto:  void QPaintEngine::drawTextItem(const QPointF & p, const QTextItem & textItem);
func (this *QPaintEngine) drawTextItem(args ...interface{}) () {
  // drawTextItem(const class QPointF &, const class QTextItem &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][1] = reflect.TypeOf(QTextItem{}) // "const QTextItem &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine12drawTextItemERK7QPointFRK9QTextItem
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTextItem).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawTextItem", args)
  }

}

  // proto:  void QPaintEngine::fix_neg_rect(int * x, int * y, int * w, int * h);
func (this *QPaintEngine) fix_neg_rect(args ...interface{}) () {
  // fix_neg_rect(int *, int *, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "int *"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine12fix_neg_rectEPiS0_S0_S0_
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int32_t)(args[3].(*int32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QPaintEngine", "fix_neg_rect", args)
  }

}

  // proto:  bool QPaintEngine::isExtended();
func (this *QPaintEngine) isExtended(args ...interface{}) () {
  // isExtended()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintEngine10isExtendedEv
  default:
    qtrt.ErrorResolve("QPaintEngine", "isExtended", args)
  }

}

// <= body block end

