package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  qreal QTextItem::descent();
extern void C_ZNK9QTextItem7descentEv(void* qthis); // 4
  // proto:  QString QTextItem::text();
extern void C_ZNK9QTextItem4textEv(void* qthis); // 4
  // proto:  qreal QTextItem::width();
extern void C_ZNK9QTextItem5widthEv(void* qthis); // 4
  // proto:  RenderFlags QTextItem::renderFlags();
extern void C_ZNK9QTextItem11renderFlagsEv(void* qthis); // 4
  // proto:  QFont QTextItem::font();
extern void C_ZNK9QTextItem4fontEv(void* qthis); // 4
  // proto:  qreal QTextItem::ascent();
extern void C_ZNK9QTextItem6ascentEv(void* qthis); // 4
  // proto:  qreal QPaintEngineState::opacity();
extern void C_ZNK17QPaintEngineState7opacityEv(void* qthis); // 4
  // proto:  bool QPaintEngineState::penNeedsResolving();
extern void C_ZNK17QPaintEngineState17penNeedsResolvingEv(void* qthis); // 4
  // proto:  bool QPaintEngineState::isClipEnabled();
extern void C_ZNK17QPaintEngineState13isClipEnabledEv(void* qthis); // 4
  // proto:  QMatrix QPaintEngineState::matrix();
extern void C_ZNK17QPaintEngineState6matrixEv(void* qthis); // 4
  // proto:  QPen QPaintEngineState::pen();
extern void C_ZNK17QPaintEngineState3penEv(void* qthis); // 4
  // proto:  QBrush QPaintEngineState::backgroundBrush();
extern void C_ZNK17QPaintEngineState15backgroundBrushEv(void* qthis); // 4
  // proto:  QPainter::CompositionMode QPaintEngineState::compositionMode();
extern void C_ZNK17QPaintEngineState15compositionModeEv(void* qthis); // 4
  // proto:  QTransform QPaintEngineState::transform();
extern void C_ZNK17QPaintEngineState9transformEv(void* qthis); // 4
  // proto:  QPainter::RenderHints QPaintEngineState::renderHints();
extern void C_ZNK17QPaintEngineState11renderHintsEv(void* qthis); // 4
  // proto:  QPaintEngine::DirtyFlags QPaintEngineState::state();
extern void C_ZNK17QPaintEngineState5stateEv(void* qthis); // 2
  // proto:  Qt::BGMode QPaintEngineState::backgroundMode();
extern void C_ZNK17QPaintEngineState14backgroundModeEv(void* qthis); // 4
  // proto:  Qt::ClipOperation QPaintEngineState::clipOperation();
extern void C_ZNK17QPaintEngineState13clipOperationEv(void* qthis); // 4
  // proto:  QBrush QPaintEngineState::brush();
extern void C_ZNK17QPaintEngineState5brushEv(void* qthis); // 4
  // proto:  QPainterPath QPaintEngineState::clipPath();
extern void C_ZNK17QPaintEngineState8clipPathEv(void* qthis); // 4
  // proto:  QFont QPaintEngineState::font();
extern void C_ZNK17QPaintEngineState4fontEv(void* qthis); // 4
  // proto:  QRegion QPaintEngineState::clipRegion();
extern void C_ZNK17QPaintEngineState10clipRegionEv(void* qthis); // 4
  // proto:  QPointF QPaintEngineState::brushOrigin();
extern void C_ZNK17QPaintEngineState11brushOriginEv(void* qthis); // 4
  // proto:  QPainter * QPaintEngineState::painter();
extern void C_ZNK17QPaintEngineState7painterEv(void* qthis); // 4
  // proto:  bool QPaintEngineState::brushNeedsResolving();
extern void C_ZNK17QPaintEngineState19brushNeedsResolvingEv(void* qthis); // 4
  // proto:  QPaintDevice * QPaintEngine::paintDevice();
extern void C_ZNK12QPaintEngine11paintDeviceEv(void* qthis); // 4
  // proto:  void QPaintEngine::drawRects(const QRectF * rects, int rectCount);
extern void C_ZN12QPaintEngine9drawRectsEPK6QRectFi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QPaintEngine::drawRects(const QRect * rects, int rectCount);
extern void C_ZN12QPaintEngine9drawRectsEPK5QRecti(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QPaintEngine::drawPoints(const QPointF * points, int pointCount);
extern void C_ZN12QPaintEngine10drawPointsEPK7QPointFi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QPaintEngine::drawPoints(const QPoint * points, int pointCount);
extern void C_ZN12QPaintEngine10drawPointsEPK6QPointi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QRect QPaintEngine::systemRect();
extern void C_ZNK12QPaintEngine10systemRectEv(void* qthis); // 4
  // proto:  void QPaintEngine::drawPath(const QPainterPath & path);
extern void C_ZN12QPaintEngine8drawPathERK12QPainterPath(void* qthis, void* arg0); // 4
  // proto:  void QPaintEngine::setPaintDevice(QPaintDevice * device);
extern void C_ZN12QPaintEngine14setPaintDeviceEP12QPaintDevice(void* qthis, void* arg0); // 4
  // proto:  QPainter * QPaintEngine::painter();
extern void C_ZNK12QPaintEngine7painterEv(void* qthis); // 4
  // proto:  void QPaintEngine::drawTiledPixmap(const QRectF & r, const QPixmap & pixmap, const QPointF & s);
extern void C_ZN12QPaintEngine15drawTiledPixmapERK6QRectFRK7QPixmapRK7QPointF(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QPaintEngine::setActive(bool newState);
extern void C_ZN12QPaintEngine9setActiveEb(void* qthis, bool arg0); // 2
  // proto:  void QPaintEngine::setSystemRect(const QRect & rect);
extern void C_ZN12QPaintEngine13setSystemRectERK5QRect(void* qthis, void* arg0); // 4
  // proto:  void QPaintEngine::syncState();
extern void C_ZN12QPaintEngine9syncStateEv(void* qthis); // 4
  // proto:  void QPaintEngine::~QPaintEngine();
extern void C_ZN12QPaintEngineD2Ev(void* qthis); // 4
  // proto:  QRegion QPaintEngine::systemClip();
extern void C_ZNK12QPaintEngine10systemClipEv(void* qthis); // 4
  // proto:  void QPaintEngine::fix_neg_rect(int * x, int * y, int * w, int * h);
extern void C_ZN12QPaintEngine12fix_neg_rectEPiS0_S0_S0_(void* qthis, int32_t* arg0, int32_t* arg1, int32_t* arg2, int32_t* arg3); // 2
  // proto:  void QPaintEngine::setSystemClip(const QRegion & baseClip);
extern void C_ZN12QPaintEngine13setSystemClipERK7QRegion(void* qthis, void* arg0); // 4
  // proto:  QPoint QPaintEngine::coordinateOffset();
extern void C_ZNK12QPaintEngine16coordinateOffsetEv(void* qthis); // 4
  // proto:  bool QPaintEngine::isActive();
extern void C_ZNK12QPaintEngine8isActiveEv(void* qthis); // 2
  // proto:  void QPaintEngine::drawEllipse(const QRectF & r);
extern void C_ZN12QPaintEngine11drawEllipseERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QPaintEngine::drawEllipse(const QRect & r);
extern void C_ZN12QPaintEngine11drawEllipseERK5QRect(void* qthis, void* arg0); // 4
  // proto:  void QPaintEngine::drawTextItem(const QPointF & p, const QTextItem & textItem);
extern void C_ZN12QPaintEngine12drawTextItemERK7QPointFRK9QTextItem(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QPaintEngine::isExtended();
extern void C_ZNK12QPaintEngine10isExtendedEv(void* qthis); // 2
  // proto:  void QPaintEngine::drawLines(const QLine * lines, int lineCount);
extern void C_ZN12QPaintEngine9drawLinesEPK5QLinei(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QPaintEngine::drawLines(const QLineF * lines, int lineCount);
extern void C_ZN12QPaintEngine9drawLinesEPK6QLineFi(void* qthis, void* arg0, int32_t arg1); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QPaintEngineState)=1
type QPaintEngineState struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QPaintEngine)=1
type QPaintEngine struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// descent()
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
    // invoke: qreal descent()
    C.C_ZNK9QTextItem7descentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextItem", "descent", args)
  }

}

// text()
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
    // invoke: QString text()
    C.C_ZNK9QTextItem4textEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextItem", "text", args)
  }

}

// width()
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
    // invoke: qreal width()
    C.C_ZNK9QTextItem5widthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextItem", "width", args)
  }

}

// renderFlags()
func (this *QTextItem) renderFlags(args ...interface{}) () {
  // renderFlags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextItem11renderFlagsEv
    // invoke: RenderFlags renderFlags()
    C.C_ZNK9QTextItem11renderFlagsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextItem", "renderFlags", args)
  }

}

// font()
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
    // invoke: QFont font()
    C.C_ZNK9QTextItem4fontEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextItem", "font", args)
  }

}

// ascent()
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
    // invoke: qreal ascent()
    C.C_ZNK9QTextItem6ascentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextItem", "ascent", args)
  }

}

// opacity()
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
    // invoke: qreal opacity()
    C.C_ZNK17QPaintEngineState7opacityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngineState", "opacity", args)
  }

}

// penNeedsResolving()
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
    // invoke: bool penNeedsResolving()
    C.C_ZNK17QPaintEngineState17penNeedsResolvingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngineState", "penNeedsResolving", args)
  }

}

// isClipEnabled()
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
    // invoke: bool isClipEnabled()
    C.C_ZNK17QPaintEngineState13isClipEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngineState", "isClipEnabled", args)
  }

}

// matrix()
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
    // invoke: QMatrix matrix()
    C.C_ZNK17QPaintEngineState6matrixEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngineState", "matrix", args)
  }

}

// pen()
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
    // invoke: QPen pen()
    C.C_ZNK17QPaintEngineState3penEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngineState", "pen", args)
  }

}

// backgroundBrush()
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
    // invoke: QBrush backgroundBrush()
    C.C_ZNK17QPaintEngineState15backgroundBrushEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngineState", "backgroundBrush", args)
  }

}

// compositionMode()
func (this *QPaintEngineState) compositionMode(args ...interface{}) () {
  // compositionMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState15compositionModeEv
    // invoke: QPainter::CompositionMode compositionMode()
    C.C_ZNK17QPaintEngineState15compositionModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngineState", "compositionMode", args)
  }

}

// transform()
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
    // invoke: QTransform transform()
    C.C_ZNK17QPaintEngineState9transformEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngineState", "transform", args)
  }

}

// renderHints()
func (this *QPaintEngineState) renderHints(args ...interface{}) () {
  // renderHints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState11renderHintsEv
    // invoke: QPainter::RenderHints renderHints()
    C.C_ZNK17QPaintEngineState11renderHintsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngineState", "renderHints", args)
  }

}

// state()
func (this *QPaintEngineState) state(args ...interface{}) () {
  // state()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState5stateEv
    // invoke: QPaintEngine::DirtyFlags state()
    C.C_ZNK17QPaintEngineState5stateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngineState", "state", args)
  }

}

// backgroundMode()
func (this *QPaintEngineState) backgroundMode(args ...interface{}) () {
  // backgroundMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState14backgroundModeEv
    // invoke: Qt::BGMode backgroundMode()
    C.C_ZNK17QPaintEngineState14backgroundModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngineState", "backgroundMode", args)
  }

}

// clipOperation()
func (this *QPaintEngineState) clipOperation(args ...interface{}) () {
  // clipOperation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState13clipOperationEv
    // invoke: Qt::ClipOperation clipOperation()
    C.C_ZNK17QPaintEngineState13clipOperationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngineState", "clipOperation", args)
  }

}

// brush()
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
    // invoke: QBrush brush()
    C.C_ZNK17QPaintEngineState5brushEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngineState", "brush", args)
  }

}

// clipPath()
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
    // invoke: QPainterPath clipPath()
    C.C_ZNK17QPaintEngineState8clipPathEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngineState", "clipPath", args)
  }

}

// font()
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
    // invoke: QFont font()
    C.C_ZNK17QPaintEngineState4fontEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngineState", "font", args)
  }

}

// clipRegion()
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
    // invoke: QRegion clipRegion()
    C.C_ZNK17QPaintEngineState10clipRegionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngineState", "clipRegion", args)
  }

}

// brushOrigin()
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
    // invoke: QPointF brushOrigin()
    C.C_ZNK17QPaintEngineState11brushOriginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngineState", "brushOrigin", args)
  }

}

// painter()
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
    // invoke: QPainter * painter()
    C.C_ZNK17QPaintEngineState7painterEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngineState", "painter", args)
  }

}

// brushNeedsResolving()
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
    // invoke: bool brushNeedsResolving()
    C.C_ZNK17QPaintEngineState19brushNeedsResolvingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngineState", "brushNeedsResolving", args)
  }

}

// paintDevice()
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
    // invoke: QPaintDevice * paintDevice()
    C.C_ZNK12QPaintEngine11paintDeviceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngine", "paintDevice", args)
  }

}

// drawRects(const class QRectF *, int)
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
    // invoke: void drawRects(const class QRectF *, int)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN12QPaintEngine9drawRectsEPK6QRectFi(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN12QPaintEngine9drawRectsEPK5QRecti
    // invoke: void drawRects(const class QRect *, int)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN12QPaintEngine9drawRectsEPK5QRecti(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawRects", args)
  }

}

// drawPoints(const class QPointF *, int)
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
    // invoke: void drawPoints(const class QPointF *, int)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN12QPaintEngine10drawPointsEPK7QPointFi(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN12QPaintEngine10drawPointsEPK6QPointi
    // invoke: void drawPoints(const class QPoint *, int)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN12QPaintEngine10drawPointsEPK6QPointi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawPoints", args)
  }

}

// systemRect()
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
    // invoke: QRect systemRect()
    C.C_ZNK12QPaintEngine10systemRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngine", "systemRect", args)
  }

}

// drawPath(const class QPainterPath &)
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
    // invoke: void drawPath(const class QPainterPath &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QPaintEngine8drawPathERK12QPainterPath(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawPath", args)
  }

}

// setPaintDevice(class QPaintDevice *)
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
    // invoke: void setPaintDevice(class QPaintDevice *)
    var arg0 = args[0].(QPaintDevice).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QPaintEngine14setPaintDeviceEP12QPaintDevice(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPaintEngine", "setPaintDevice", args)
  }

}

// painter()
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
    // invoke: QPainter * painter()
    C.C_ZNK12QPaintEngine7painterEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngine", "painter", args)
  }

}

// drawTiledPixmap(const class QRectF &, const class QPixmap &, const class QPointF &)
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
    // invoke: void drawTiledPixmap(const class QRectF &, const class QPixmap &, const class QPointF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPixmap).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QPointF).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN12QPaintEngine15drawTiledPixmapERK6QRectFRK7QPixmapRK7QPointF(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawTiledPixmap", args)
  }

}

// setActive(_Bool)
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
    // invoke: void setActive(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN12QPaintEngine9setActiveEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPaintEngine", "setActive", args)
  }

}

// setSystemRect(const class QRect &)
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
    // invoke: void setSystemRect(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QPaintEngine13setSystemRectERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPaintEngine", "setSystemRect", args)
  }

}

// syncState()
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
    // invoke: void syncState()
    C.C_ZN12QPaintEngine9syncStateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngine", "syncState", args)
  }

}

// ~QPaintEngine()
func (this *QPaintEngine) FreeQPaintEngine(args ...interface{}) () {
  // ~QPaintEngine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngineD0Ev
    // invoke: void ~QPaintEngine()
    C.C_ZN12QPaintEngineD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngine", "~QPaintEngine", args)
  }

}

// systemClip()
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
    // invoke: QRegion systemClip()
    C.C_ZNK12QPaintEngine10systemClipEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngine", "systemClip", args)
  }

}

// fix_neg_rect(int *, int *, int *, int *)
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
    // invoke: void fix_neg_rect(int *, int *, int *, int *)
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int32_t)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    C.C_ZN12QPaintEngine12fix_neg_rectEPiS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QPaintEngine", "fix_neg_rect", args)
  }

}

// setSystemClip(const class QRegion &)
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
    // invoke: void setSystemClip(const class QRegion &)
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QPaintEngine13setSystemClipERK7QRegion(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPaintEngine", "setSystemClip", args)
  }

}

// coordinateOffset()
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
    // invoke: QPoint coordinateOffset()
    C.C_ZNK12QPaintEngine16coordinateOffsetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngine", "coordinateOffset", args)
  }

}

// isActive()
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
    // invoke: bool isActive()
    C.C_ZNK12QPaintEngine8isActiveEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngine", "isActive", args)
  }

}

// drawEllipse(const class QRectF &)
func (this *QPaintEngine) drawEllipse(args ...interface{}) () {
  // drawEllipse(const class QRectF &)
  // drawEllipse(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine11drawEllipseERK6QRectF
    // invoke: void drawEllipse(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QPaintEngine11drawEllipseERK6QRectF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN12QPaintEngine11drawEllipseERK5QRect
    // invoke: void drawEllipse(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QPaintEngine11drawEllipseERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawEllipse", args)
  }

}

// drawTextItem(const class QPointF &, const class QTextItem &)
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
    // invoke: void drawTextItem(const class QPointF &, const class QTextItem &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTextItem).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN12QPaintEngine12drawTextItemERK7QPointFRK9QTextItem(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawTextItem", args)
  }

}

// isExtended()
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
    // invoke: bool isExtended()
    C.C_ZNK12QPaintEngine10isExtendedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngine", "isExtended", args)
  }

}

// drawLines(const class QLine *, int)
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
    // invoke: void drawLines(const class QLine *, int)
    var arg0 = args[0].(QLine).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN12QPaintEngine9drawLinesEPK5QLinei(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN12QPaintEngine9drawLinesEPK6QLineFi
    // invoke: void drawLines(const class QLineF *, int)
    var arg0 = args[0].(QLineF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN12QPaintEngine9drawLinesEPK6QLineFi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawLines", args)
  }

}

// <= body block end

