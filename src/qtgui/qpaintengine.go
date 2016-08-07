package qtgui
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  qreal QTextItem::descent();
extern double C_ZNK9QTextItem7descentEv(void* qthis); // 4
  // proto:  QString QTextItem::text();
extern void* C_ZNK9QTextItem4textEv(void* qthis); // 4
  // proto:  qreal QTextItem::width();
extern double C_ZNK9QTextItem5widthEv(void* qthis); // 4
  // proto:  RenderFlags QTextItem::renderFlags();
extern void C_ZNK9QTextItem11renderFlagsEv(void* qthis); // 4
  // proto:  QFont QTextItem::font();
extern void* C_ZNK9QTextItem4fontEv(void* qthis); // 4
  // proto:  qreal QTextItem::ascent();
extern double C_ZNK9QTextItem6ascentEv(void* qthis); // 4
  // proto:  qreal QPaintEngineState::opacity();
extern double C_ZNK17QPaintEngineState7opacityEv(void* qthis); // 4
  // proto:  bool QPaintEngineState::penNeedsResolving();
extern bool C_ZNK17QPaintEngineState17penNeedsResolvingEv(void* qthis); // 4
  // proto:  bool QPaintEngineState::isClipEnabled();
extern bool C_ZNK17QPaintEngineState13isClipEnabledEv(void* qthis); // 4
  // proto:  QMatrix QPaintEngineState::matrix();
extern void* C_ZNK17QPaintEngineState6matrixEv(void* qthis); // 4
  // proto:  QPen QPaintEngineState::pen();
extern void* C_ZNK17QPaintEngineState3penEv(void* qthis); // 4
  // proto:  QBrush QPaintEngineState::backgroundBrush();
extern void* C_ZNK17QPaintEngineState15backgroundBrushEv(void* qthis); // 4
  // proto:  QPainter::CompositionMode QPaintEngineState::compositionMode();
extern void C_ZNK17QPaintEngineState15compositionModeEv(void* qthis); // 4
  // proto:  QTransform QPaintEngineState::transform();
extern void* C_ZNK17QPaintEngineState9transformEv(void* qthis); // 4
  // proto:  QPainter::RenderHints QPaintEngineState::renderHints();
extern void C_ZNK17QPaintEngineState11renderHintsEv(void* qthis); // 4
  // proto:  QPaintEngine::DirtyFlags QPaintEngineState::state();
extern void C_ZNK17QPaintEngineState5stateEv(void* qthis); // 2
  // proto:  Qt::BGMode QPaintEngineState::backgroundMode();
extern void C_ZNK17QPaintEngineState14backgroundModeEv(void* qthis); // 4
  // proto:  Qt::ClipOperation QPaintEngineState::clipOperation();
extern void C_ZNK17QPaintEngineState13clipOperationEv(void* qthis); // 4
  // proto:  QBrush QPaintEngineState::brush();
extern void* C_ZNK17QPaintEngineState5brushEv(void* qthis); // 4
  // proto:  QPainterPath QPaintEngineState::clipPath();
extern void* C_ZNK17QPaintEngineState8clipPathEv(void* qthis); // 4
  // proto:  QFont QPaintEngineState::font();
extern void* C_ZNK17QPaintEngineState4fontEv(void* qthis); // 4
  // proto:  QRegion QPaintEngineState::clipRegion();
extern void* C_ZNK17QPaintEngineState10clipRegionEv(void* qthis); // 4
  // proto:  QPointF QPaintEngineState::brushOrigin();
extern void* C_ZNK17QPaintEngineState11brushOriginEv(void* qthis); // 4
  // proto:  QPainter * QPaintEngineState::painter();
extern void* C_ZNK17QPaintEngineState7painterEv(void* qthis); // 4
  // proto:  bool QPaintEngineState::brushNeedsResolving();
extern bool C_ZNK17QPaintEngineState19brushNeedsResolvingEv(void* qthis); // 4
  // proto:  QPaintDevice * QPaintEngine::paintDevice();
extern void* C_ZNK12QPaintEngine11paintDeviceEv(void* qthis); // 4
  // proto:  void QPaintEngine::drawRects(const QRectF * rects, int rectCount);
extern void C_ZN12QPaintEngine9drawRectsEPK6QRectFi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QPaintEngine::drawRects(const QRect * rects, int rectCount);
extern void C_ZN12QPaintEngine9drawRectsEPK5QRecti(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QPaintEngine::drawPoints(const QPointF * points, int pointCount);
extern void C_ZN12QPaintEngine10drawPointsEPK7QPointFi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QPaintEngine::drawPoints(const QPoint * points, int pointCount);
extern void C_ZN12QPaintEngine10drawPointsEPK6QPointi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QRect QPaintEngine::systemRect();
extern void* C_ZNK12QPaintEngine10systemRectEv(void* qthis); // 4
  // proto:  void QPaintEngine::drawPath(const QPainterPath & path);
extern void C_ZN12QPaintEngine8drawPathERK12QPainterPath(void* qthis, void* arg0); // 4
  // proto:  void QPaintEngine::setPaintDevice(QPaintDevice * device);
extern void C_ZN12QPaintEngine14setPaintDeviceEP12QPaintDevice(void* qthis, void* arg0); // 4
  // proto:  QPainter * QPaintEngine::painter();
extern void* C_ZNK12QPaintEngine7painterEv(void* qthis); // 4
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
extern void* C_ZNK12QPaintEngine10systemClipEv(void* qthis); // 4
  // proto:  void QPaintEngine::fix_neg_rect(int * x, int * y, int * w, int * h);
extern void C_ZN12QPaintEngine12fix_neg_rectEPiS0_S0_S0_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3); // 2
  // proto:  void QPaintEngine::setSystemClip(const QRegion & baseClip);
extern void C_ZN12QPaintEngine13setSystemClipERK7QRegion(void* qthis, void* arg0); // 4
  // proto:  QPoint QPaintEngine::coordinateOffset();
extern void* C_ZNK12QPaintEngine16coordinateOffsetEv(void* qthis); // 4
  // proto:  bool QPaintEngine::isActive();
extern bool C_ZNK12QPaintEngine8isActiveEv(void* qthis); // 2
  // proto:  void QPaintEngine::drawEllipse(const QRectF & r);
extern void C_ZN12QPaintEngine11drawEllipseERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QPaintEngine::drawEllipse(const QRect & r);
extern void C_ZN12QPaintEngine11drawEllipseERK5QRect(void* qthis, void* arg0); // 4
  // proto:  void QPaintEngine::drawTextItem(const QPointF & p, const QTextItem & textItem);
extern void C_ZN12QPaintEngine12drawTextItemERK7QPointFRK9QTextItem(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QPaintEngine::isExtended();
extern bool C_ZNK12QPaintEngine10isExtendedEv(void* qthis); // 2
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
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QTextItem)=1
type QTextItem struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QPaintEngineState)=1
type QPaintEngineState struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QPaintEngine)=1
type QPaintEngine struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// descent()
func (this *QTextItem) Descent(args ...interface{}) (ret interface{}) {
  // descent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextItem7descentEv
    // invoke: qreal descent()
    var ret0 = C.C_ZNK9QTextItem7descentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextItem", "descent", args)
  }

  return
}

// text()
func (this *QTextItem) Text(args ...interface{}) (ret interface{}) {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextItem4textEv
    // invoke: QString text()
    var ret0 = C.C_ZNK9QTextItem4textEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextItem", "text", args)
  }

  return
}

// width()
func (this *QTextItem) Width(args ...interface{}) (ret interface{}) {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextItem5widthEv
    // invoke: qreal width()
    var ret0 = C.C_ZNK9QTextItem5widthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextItem", "width", args)
  }

  return
}

// renderFlags()
func (this *QTextItem) Renderflags(args ...interface{}) () {
  // renderFlags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextItem11renderFlagsEv
    // invoke: RenderFlags renderFlags()
    C.C_ZNK9QTextItem11renderFlagsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextItem", "renderFlags", args)
  }

  return
}

// font()
func (this *QTextItem) Font(args ...interface{}) (ret interface{}) {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextItem4fontEv
    // invoke: QFont font()
    var ret0 = C.C_ZNK9QTextItem4fontEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QFont{}) // "QFont"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextItem", "font", args)
  }

  return
}

// ascent()
func (this *QTextItem) Ascent(args ...interface{}) (ret interface{}) {
  // ascent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextItem6ascentEv
    // invoke: qreal ascent()
    var ret0 = C.C_ZNK9QTextItem6ascentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextItem", "ascent", args)
  }

  return
}

// opacity()
func (this *QPaintEngineState) Opacity(args ...interface{}) (ret interface{}) {
  // opacity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState7opacityEv
    // invoke: qreal opacity()
    var ret0 = C.C_ZNK17QPaintEngineState7opacityEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintEngineState", "opacity", args)
  }

  return
}

// penNeedsResolving()
func (this *QPaintEngineState) Penneedsresolving(args ...interface{}) (ret interface{}) {
  // penNeedsResolving()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState17penNeedsResolvingEv
    // invoke: bool penNeedsResolving()
    var ret0 = C.C_ZNK17QPaintEngineState17penNeedsResolvingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintEngineState", "penNeedsResolving", args)
  }

  return
}

// isClipEnabled()
func (this *QPaintEngineState) Isclipenabled(args ...interface{}) (ret interface{}) {
  // isClipEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState13isClipEnabledEv
    // invoke: bool isClipEnabled()
    var ret0 = C.C_ZNK17QPaintEngineState13isClipEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintEngineState", "isClipEnabled", args)
  }

  return
}

// matrix()
func (this *QPaintEngineState) Matrix(args ...interface{}) (ret interface{}) {
  // matrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState6matrixEv
    // invoke: QMatrix matrix()
    var ret0 = C.C_ZNK17QPaintEngineState6matrixEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMatrix{}) // "QMatrix"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintEngineState", "matrix", args)
  }

  return
}

// pen()
func (this *QPaintEngineState) Pen(args ...interface{}) (ret interface{}) {
  // pen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState3penEv
    // invoke: QPen pen()
    var ret0 = C.C_ZNK17QPaintEngineState3penEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPen{}) // "QPen"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintEngineState", "pen", args)
  }

  return
}

// backgroundBrush()
func (this *QPaintEngineState) Backgroundbrush(args ...interface{}) (ret interface{}) {
  // backgroundBrush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState15backgroundBrushEv
    // invoke: QBrush backgroundBrush()
    var ret0 = C.C_ZNK17QPaintEngineState15backgroundBrushEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "QBrush"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintEngineState", "backgroundBrush", args)
  }

  return
}

// compositionMode()
func (this *QPaintEngineState) Compositionmode(args ...interface{}) () {
  // compositionMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState15compositionModeEv
    // invoke: QPainter::CompositionMode compositionMode()
    C.C_ZNK17QPaintEngineState15compositionModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngineState", "compositionMode", args)
  }

  return
}

// transform()
func (this *QPaintEngineState) Transform(args ...interface{}) (ret interface{}) {
  // transform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState9transformEv
    // invoke: QTransform transform()
    var ret0 = C.C_ZNK17QPaintEngineState9transformEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTransform{}) // "QTransform"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintEngineState", "transform", args)
  }

  return
}

// renderHints()
func (this *QPaintEngineState) Renderhints(args ...interface{}) () {
  // renderHints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState11renderHintsEv
    // invoke: QPainter::RenderHints renderHints()
    C.C_ZNK17QPaintEngineState11renderHintsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngineState", "renderHints", args)
  }

  return
}

// state()
func (this *QPaintEngineState) State(args ...interface{}) () {
  // state()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState5stateEv
    // invoke: QPaintEngine::DirtyFlags state()
    C.C_ZNK17QPaintEngineState5stateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngineState", "state", args)
  }

  return
}

// backgroundMode()
func (this *QPaintEngineState) Backgroundmode(args ...interface{}) () {
  // backgroundMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState14backgroundModeEv
    // invoke: Qt::BGMode backgroundMode()
    C.C_ZNK17QPaintEngineState14backgroundModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngineState", "backgroundMode", args)
  }

  return
}

// clipOperation()
func (this *QPaintEngineState) Clipoperation(args ...interface{}) () {
  // clipOperation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState13clipOperationEv
    // invoke: Qt::ClipOperation clipOperation()
    C.C_ZNK17QPaintEngineState13clipOperationEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngineState", "clipOperation", args)
  }

  return
}

// brush()
func (this *QPaintEngineState) Brush(args ...interface{}) (ret interface{}) {
  // brush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState5brushEv
    // invoke: QBrush brush()
    var ret0 = C.C_ZNK17QPaintEngineState5brushEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "QBrush"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintEngineState", "brush", args)
  }

  return
}

// clipPath()
func (this *QPaintEngineState) Clippath(args ...interface{}) (ret interface{}) {
  // clipPath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState8clipPathEv
    // invoke: QPainterPath clipPath()
    var ret0 = C.C_ZNK17QPaintEngineState8clipPathEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintEngineState", "clipPath", args)
  }

  return
}

// font()
func (this *QPaintEngineState) Font(args ...interface{}) (ret interface{}) {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState4fontEv
    // invoke: QFont font()
    var ret0 = C.C_ZNK17QPaintEngineState4fontEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QFont{}) // "QFont"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintEngineState", "font", args)
  }

  return
}

// clipRegion()
func (this *QPaintEngineState) Clipregion(args ...interface{}) (ret interface{}) {
  // clipRegion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState10clipRegionEv
    // invoke: QRegion clipRegion()
    var ret0 = C.C_ZNK17QPaintEngineState10clipRegionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRegion{}) // "QRegion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintEngineState", "clipRegion", args)
  }

  return
}

// brushOrigin()
func (this *QPaintEngineState) Brushorigin(args ...interface{}) (ret interface{}) {
  // brushOrigin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState11brushOriginEv
    // invoke: QPointF brushOrigin()
    var ret0 = C.C_ZNK17QPaintEngineState11brushOriginEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintEngineState", "brushOrigin", args)
  }

  return
}

// painter()
func (this *QPaintEngineState) Painter(args ...interface{}) (ret interface{}) {
  // painter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState7painterEv
    // invoke: QPainter * painter()
    var ret0 = C.C_ZNK17QPaintEngineState7painterEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainter{}) // "QPainter *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintEngineState", "painter", args)
  }

  return
}

// brushNeedsResolving()
func (this *QPaintEngineState) Brushneedsresolving(args ...interface{}) (ret interface{}) {
  // brushNeedsResolving()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState19brushNeedsResolvingEv
    // invoke: bool brushNeedsResolving()
    var ret0 = C.C_ZNK17QPaintEngineState19brushNeedsResolvingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintEngineState", "brushNeedsResolving", args)
  }

  return
}

// paintDevice()
func (this *QPaintEngine) Paintdevice(args ...interface{}) (ret interface{}) {
  // paintDevice()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintEngine11paintDeviceEv
    // invoke: QPaintDevice * paintDevice()
    var ret0 = C.C_ZNK12QPaintEngine11paintDeviceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPaintDevice{}) // "QPaintDevice *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintEngine", "paintDevice", args)
  }

  return
}

// drawRects(const class QRectF *, int)
func (this *QPaintEngine) Drawrects(args ...interface{}) () {
  // drawRects(const class QRectF *, int)
  // drawRects(const class QRect *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine9drawRectsEPK6QRectFi
    // invoke: void drawRects(const class QRectF *, int)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN12QPaintEngine9drawRectsEPK6QRectFi(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN12QPaintEngine9drawRectsEPK5QRecti
    // invoke: void drawRects(const class QRect *, int)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN12QPaintEngine9drawRectsEPK5QRecti(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawRects", args)
  }

  return
}

// drawPoints(const class QPointF *, int)
func (this *QPaintEngine) Drawpoints(args ...interface{}) () {
  // drawPoints(const class QPointF *, int)
  // drawPoints(const class QPoint *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine10drawPointsEPK7QPointFi
    // invoke: void drawPoints(const class QPointF *, int)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN12QPaintEngine10drawPointsEPK7QPointFi(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN12QPaintEngine10drawPointsEPK6QPointi
    // invoke: void drawPoints(const class QPoint *, int)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN12QPaintEngine10drawPointsEPK6QPointi(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawPoints", args)
  }

  return
}

// systemRect()
func (this *QPaintEngine) Systemrect(args ...interface{}) (ret interface{}) {
  // systemRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintEngine10systemRectEv
    // invoke: QRect systemRect()
    var ret0 = C.C_ZNK12QPaintEngine10systemRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintEngine", "systemRect", args)
  }

  return
}

// drawPath(const class QPainterPath &)
func (this *QPaintEngine) Drawpath(args ...interface{}) () {
  // drawPath(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine8drawPathERK12QPainterPath
    // invoke: void drawPath(const class QPainterPath &)
    var arg0 = args[0].(*QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QPaintEngine8drawPathERK12QPainterPath(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawPath", args)
  }

  return
}

// setPaintDevice(class QPaintDevice *)
func (this *QPaintEngine) Setpaintdevice(args ...interface{}) () {
  // setPaintDevice(class QPaintDevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPaintDevice{}) // "QPaintDevice *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine14setPaintDeviceEP12QPaintDevice
    // invoke: void setPaintDevice(class QPaintDevice *)
    var arg0 = args[0].(*QPaintDevice).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QPaintEngine14setPaintDeviceEP12QPaintDevice(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPaintEngine", "setPaintDevice", args)
  }

  return
}

// painter()
func (this *QPaintEngine) Painter(args ...interface{}) (ret interface{}) {
  // painter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintEngine7painterEv
    // invoke: QPainter * painter()
    var ret0 = C.C_ZNK12QPaintEngine7painterEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainter{}) // "QPainter *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintEngine", "painter", args)
  }

  return
}

// drawTiledPixmap(const class QRectF &, const class QPixmap &, const class QPointF &)
func (this *QPaintEngine) Drawtiledpixmap(args ...interface{}) () {
  // drawTiledPixmap(const class QRectF &, const class QPixmap &, const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[0][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[0][2] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine15drawTiledPixmapERK6QRectFRK7QPixmapRK7QPointF
    // invoke: void drawTiledPixmap(const class QRectF &, const class QPixmap &, const class QPointF &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPixmap).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN12QPaintEngine15drawTiledPixmapERK6QRectFRK7QPixmapRK7QPointF(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawTiledPixmap", args)
  }

  return
}

// setActive(_Bool)
func (this *QPaintEngine) Setactive(args ...interface{}) () {
  // setActive(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine9setActiveEb
    // invoke: void setActive(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN12QPaintEngine9setActiveEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPaintEngine", "setActive", args)
  }

  return
}

// setSystemRect(const class QRect &)
func (this *QPaintEngine) Setsystemrect(args ...interface{}) () {
  // setSystemRect(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine13setSystemRectERK5QRect
    // invoke: void setSystemRect(const class QRect &)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QPaintEngine13setSystemRectERK5QRect(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPaintEngine", "setSystemRect", args)
  }

  return
}

// syncState()
func (this *QPaintEngine) Syncstate(args ...interface{}) () {
  // syncState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine9syncStateEv
    // invoke: void syncState()
    C.C_ZN12QPaintEngine9syncStateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngine", "syncState", args)
  }

  return
}

// ~QPaintEngine()
func (this *QPaintEngine) Freeqpaintengine(args ...interface{}) () {
  // ~QPaintEngine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngineD0Ev
    // invoke: void ~QPaintEngine()
    C.C_ZN12QPaintEngineD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEngine", "~QPaintEngine", args)
  }

  return
}

// systemClip()
func (this *QPaintEngine) Systemclip(args ...interface{}) (ret interface{}) {
  // systemClip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintEngine10systemClipEv
    // invoke: QRegion systemClip()
    var ret0 = C.C_ZNK12QPaintEngine10systemClipEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRegion{}) // "QRegion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintEngine", "systemClip", args)
  }

  return
}

// fix_neg_rect(int *, int *, int *, int *)
func (this *QPaintEngine) Fix_Neg_Rect(args ...interface{}) () {
  // fix_neg_rect(int *, int *, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "int *"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine12fix_neg_rectEPiS0_S0_S0_
    // invoke: void fix_neg_rect(int *, int *, int *, int *)
    var arg0 = (unsafe.Pointer)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    C.C_ZN12QPaintEngine12fix_neg_rectEPiS0_S0_S0_(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QPaintEngine", "fix_neg_rect", args)
  }

  return
}

// setSystemClip(const class QRegion &)
func (this *QPaintEngine) Setsystemclip(args ...interface{}) () {
  // setSystemClip(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine13setSystemClipERK7QRegion
    // invoke: void setSystemClip(const class QRegion &)
    var arg0 = args[0].(*QRegion).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QPaintEngine13setSystemClipERK7QRegion(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPaintEngine", "setSystemClip", args)
  }

  return
}

// coordinateOffset()
func (this *QPaintEngine) Coordinateoffset(args ...interface{}) (ret interface{}) {
  // coordinateOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintEngine16coordinateOffsetEv
    // invoke: QPoint coordinateOffset()
    var ret0 = C.C_ZNK12QPaintEngine16coordinateOffsetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintEngine", "coordinateOffset", args)
  }

  return
}

// isActive()
func (this *QPaintEngine) Isactive(args ...interface{}) (ret interface{}) {
  // isActive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintEngine8isActiveEv
    // invoke: bool isActive()
    var ret0 = C.C_ZNK12QPaintEngine8isActiveEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintEngine", "isActive", args)
  }

  return
}

// drawEllipse(const class QRectF &)
func (this *QPaintEngine) Drawellipse(args ...interface{}) () {
  // drawEllipse(const class QRectF &)
  // drawEllipse(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine11drawEllipseERK6QRectF
    // invoke: void drawEllipse(const class QRectF &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QPaintEngine11drawEllipseERK6QRectF(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN12QPaintEngine11drawEllipseERK5QRect
    // invoke: void drawEllipse(const class QRect &)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QPaintEngine11drawEllipseERK5QRect(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawEllipse", args)
  }

  return
}

// drawTextItem(const class QPointF &, const class QTextItem &)
func (this *QPaintEngine) Drawtextitem(args ...interface{}) () {
  // drawTextItem(const class QPointF &, const class QTextItem &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[0][1] = reflect.TypeOf(QTextItem{}) // "const QTextItem &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine12drawTextItemERK7QPointFRK9QTextItem
    // invoke: void drawTextItem(const class QPointF &, const class QTextItem &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QTextItem).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN12QPaintEngine12drawTextItemERK7QPointFRK9QTextItem(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawTextItem", args)
  }

  return
}

// isExtended()
func (this *QPaintEngine) Isextended(args ...interface{}) (ret interface{}) {
  // isExtended()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintEngine10isExtendedEv
    // invoke: bool isExtended()
    var ret0 = C.C_ZNK12QPaintEngine10isExtendedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintEngine", "isExtended", args)
  }

  return
}

// drawLines(const class QLine *, int)
func (this *QPaintEngine) Drawlines(args ...interface{}) () {
  // drawLines(const class QLine *, int)
  // drawLines(const class QLineF *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QLine{}) // "const QLine *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QLineF{}) // "const QLineF *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine9drawLinesEPK5QLinei
    // invoke: void drawLines(const class QLine *, int)
    var arg0 = args[0].(*qtcore.QLine).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN12QPaintEngine9drawLinesEPK5QLinei(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN12QPaintEngine9drawLinesEPK6QLineFi
    // invoke: void drawLines(const class QLineF *, int)
    var arg0 = args[0].(*qtcore.QLineF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN12QPaintEngine9drawLinesEPK6QLineFi(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawLines", args)
  }

  return
}

// <= body block end

