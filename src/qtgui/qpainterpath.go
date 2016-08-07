package qtgui
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
// src-file: /QtGui/qpainterpath.h
// dst-file: /src/gui/qpainterpath.go
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
  // proto:  void QPainterPath::addRect(const QRectF & rect);
extern void C_ZN12QPainterPath7addRectERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QPainterPath::addRect(qreal x, qreal y, qreal w, qreal h);
extern void C_ZN12QPainterPath7addRectEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  QRectF QPainterPath::controlPointRect();
extern void* C_ZNK12QPainterPath16controlPointRectEv(void* qthis); // 4
  // proto:  void QPainterPath::moveTo(qreal x, qreal y);
extern void C_ZN12QPainterPath6moveToEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QPainterPath::moveTo(const QPointF & p);
extern void C_ZN12QPainterPath6moveToERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QPainterPath::lineTo(const QPointF & p);
extern void C_ZN12QPainterPath6lineToERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QPainterPath::lineTo(qreal x, qreal y);
extern void C_ZN12QPainterPath6lineToEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  bool QPainterPath::intersects(const QPainterPath & p);
extern bool C_ZNK12QPainterPath10intersectsERKS_(void* qthis, void* arg0); // 4
  // proto:  bool QPainterPath::intersects(const QRectF & rect);
extern bool C_ZNK12QPainterPath10intersectsERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  int QPainterPath::elementCount();
extern int32_t C_ZNK12QPainterPath12elementCountEv(void* qthis); // 4
  // proto:  void QPainterPath::closeSubpath();
extern void C_ZN12QPainterPath12closeSubpathEv(void* qthis); // 4
  // proto:  QList<QPolygonF> QPainterPath::toSubpathPolygons(const QTransform & matrix);
extern void C_ZNK12QPainterPath17toSubpathPolygonsERK10QTransform(void* qthis, void* arg0); // 4
  // proto:  QList<QPolygonF> QPainterPath::toSubpathPolygons(const QMatrix & matrix);
extern void C_ZNK12QPainterPath17toSubpathPolygonsERK7QMatrix(void* qthis, void* arg0); // 4
  // proto:  void QPainterPath::QPainterPath(const QPointF & startPoint);
extern void* C_ZN12QPainterPathC2ERK7QPointF(void* arg0); // 3
  // proto:  void QPainterPath::QPainterPath();
extern void* C_ZN12QPainterPathC2Ev(); // 3
  // proto:  void QPainterPath::QPainterPath(const QPainterPath & other);
extern void* C_ZN12QPainterPathC2ERKS_(void* arg0); // 3
  // proto:  qreal QPainterPath::angleAtPercent(qreal t);
extern double C_ZNK12QPainterPath14angleAtPercentEd(void* qthis, double arg0); // 4
  // proto:  void QPainterPath::addEllipse(const QRectF & rect);
extern void C_ZN12QPainterPath10addEllipseERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QPainterPath::addEllipse(const QPointF & center, qreal rx, qreal ry);
extern void C_ZN12QPainterPath10addEllipseERK7QPointFdd(void* qthis, void* arg0, double arg1, double arg2); // 2
  // proto:  void QPainterPath::addEllipse(qreal x, qreal y, qreal w, qreal h);
extern void C_ZN12QPainterPath10addEllipseEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  void QPainterPath::quadTo(const QPointF & ctrlPt, const QPointF & endPt);
extern void C_ZN12QPainterPath6quadToERK7QPointFS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QPainterPath::quadTo(qreal ctrlPtx, qreal ctrlPty, qreal endPtx, qreal endPty);
extern void C_ZN12QPainterPath6quadToEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  Qt::FillRule QPainterPath::fillRule();
extern void C_ZNK12QPainterPath8fillRuleEv(void* qthis); // 4
  // proto:  QPainterPath QPainterPath::intersected(const QPainterPath & r);
extern void* C_ZNK12QPainterPath11intersectedERKS_(void* qthis, void* arg0); // 4
  // proto:  void QPainterPath::translate(qreal dx, qreal dy);
extern void C_ZN12QPainterPath9translateEdd(void* qthis, double arg0, double arg1); // 4
  // proto:  void QPainterPath::translate(const QPointF & offset);
extern void C_ZN12QPainterPath9translateERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  bool QPainterPath::contains(const QPainterPath & p);
extern bool C_ZNK12QPainterPath8containsERKS_(void* qthis, void* arg0); // 4
  // proto:  bool QPainterPath::contains(const QRectF & rect);
extern bool C_ZNK12QPainterPath8containsERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  bool QPainterPath::contains(const QPointF & pt);
extern bool C_ZNK12QPainterPath8containsERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QPainterPath::setElementPositionAt(int i, qreal x, qreal y);
extern void C_ZN12QPainterPath20setElementPositionAtEidd(void* qthis, int32_t arg0, double arg1, double arg2); // 4
  // proto:  void QPainterPath::~QPainterPath();
extern void C_ZN12QPainterPathD2Ev(void* qthis); // 4
  // proto:  void QPainterPath::swap(QPainterPath & other);
extern void C_ZN12QPainterPath4swapERS_(void* qthis, void* arg0); // 2
  // proto:  QList<QPolygonF> QPainterPath::toFillPolygons(const QMatrix & matrix);
extern void C_ZNK12QPainterPath14toFillPolygonsERK7QMatrix(void* qthis, void* arg0); // 4
  // proto:  QList<QPolygonF> QPainterPath::toFillPolygons(const QTransform & matrix);
extern void C_ZNK12QPainterPath14toFillPolygonsERK10QTransform(void* qthis, void* arg0); // 4
  // proto:  bool QPainterPath::isEmpty();
extern bool C_ZNK12QPainterPath7isEmptyEv(void* qthis); // 4
  // proto:  void QPainterPath::addRegion(const QRegion & region);
extern void C_ZN12QPainterPath9addRegionERK7QRegion(void* qthis, void* arg0); // 4
  // proto:  void QPainterPath::addPath(const QPainterPath & path);
extern void C_ZN12QPainterPath7addPathERKS_(void* qthis, void* arg0); // 4
  // proto:  void QPainterPath::addPolygon(const QPolygonF & polygon);
extern void C_ZN12QPainterPath10addPolygonERK9QPolygonF(void* qthis, void* arg0); // 4
  // proto:  void QPainterPath::arcTo(const QRectF & rect, qreal startAngle, qreal arcLength);
extern void C_ZN12QPainterPath5arcToERK6QRectFdd(void* qthis, void* arg0, double arg1, double arg2); // 4
  // proto:  void QPainterPath::arcTo(qreal x, qreal y, qreal w, qreal h, qreal startAngle, qreal arcLength);
extern void C_ZN12QPainterPath5arcToEdddddd(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4, double arg5); // 2
  // proto:  qreal QPainterPath::slopeAtPercent(qreal t);
extern double C_ZNK12QPainterPath14slopeAtPercentEd(void* qthis, double arg0); // 4
  // proto:  void QPainterPath::addText(const QPointF & point, const QFont & f, const QString & text);
extern void C_ZN12QPainterPath7addTextERK7QPointFRK5QFontRK7QString(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QPainterPath::addText(qreal x, qreal y, const QFont & f, const QString & text);
extern void C_ZN12QPainterPath7addTextEddRK5QFontRK7QString(void* qthis, double arg0, double arg1, void* arg2, void* arg3); // 2
  // proto:  void QPainterPath::addRoundRect(qreal x, qreal y, qreal w, qreal h, int xRnd, int yRnd);
extern void C_ZN12QPainterPath12addRoundRectEddddii(void* qthis, double arg0, double arg1, double arg2, double arg3, int32_t arg4, int32_t arg5); // 2
  // proto:  void QPainterPath::addRoundRect(const QRectF & rect, int xRnd, int yRnd);
extern void C_ZN12QPainterPath12addRoundRectERK6QRectFii(void* qthis, void* arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  void QPainterPath::addRoundRect(qreal x, qreal y, qreal w, qreal h, int roundness);
extern void C_ZN12QPainterPath12addRoundRectEddddi(void* qthis, double arg0, double arg1, double arg2, double arg3, int32_t arg4); // 2
  // proto:  void QPainterPath::addRoundRect(const QRectF & rect, int roundness);
extern void C_ZN12QPainterPath12addRoundRectERK6QRectFi(void* qthis, void* arg0, int32_t arg1); // 2
  // proto:  QPainterPath::Element QPainterPath::elementAt(int i);
extern void C_ZNK12QPainterPath9elementAtEi(void* qthis, int32_t arg0); // 4
  // proto:  void QPainterPath::connectPath(const QPainterPath & path);
extern void C_ZN12QPainterPath11connectPathERKS_(void* qthis, void* arg0); // 4
  // proto:  QRectF QPainterPath::boundingRect();
extern void* C_ZNK12QPainterPath12boundingRectEv(void* qthis); // 4
  // proto:  void QPainterPath::cubicTo(const QPointF & ctrlPt1, const QPointF & ctrlPt2, const QPointF & endPt);
extern void C_ZN12QPainterPath7cubicToERK7QPointFS2_S2_(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QPainterPath::cubicTo(qreal ctrlPt1x, qreal ctrlPt1y, qreal ctrlPt2x, qreal ctrlPt2y, qreal endPtx, qreal endPty);
extern void C_ZN12QPainterPath7cubicToEdddddd(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4, double arg5); // 2
  // proto:  QPainterPath QPainterPath::toReversed();
extern void* C_ZNK12QPainterPath10toReversedEv(void* qthis); // 4
  // proto:  QPainterPath QPainterPath::subtractedInverted(const QPainterPath & r);
extern void* C_ZNK12QPainterPath18subtractedInvertedERKS_(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QPainterPath::toFillPolygon(const QMatrix & matrix);
extern void* C_ZNK12QPainterPath13toFillPolygonERK7QMatrix(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QPainterPath::toFillPolygon(const QTransform & matrix);
extern void* C_ZNK12QPainterPath13toFillPolygonERK10QTransform(void* qthis, void* arg0); // 4
  // proto:  QPainterPath QPainterPath::subtracted(const QPainterPath & r);
extern void* C_ZNK12QPainterPath10subtractedERKS_(void* qthis, void* arg0); // 4
  // proto:  QPointF QPainterPath::pointAtPercent(qreal t);
extern void* C_ZNK12QPainterPath14pointAtPercentEd(void* qthis, double arg0); // 4
  // proto:  QPainterPath QPainterPath::united(const QPainterPath & r);
extern void* C_ZNK12QPainterPath6unitedERKS_(void* qthis, void* arg0); // 4
  // proto:  QPointF QPainterPath::currentPosition();
extern void* C_ZNK12QPainterPath15currentPositionEv(void* qthis); // 4
  // proto:  qreal QPainterPath::length();
extern double C_ZNK12QPainterPath6lengthEv(void* qthis); // 4
  // proto:  QPainterPath QPainterPath::translated(qreal dx, qreal dy);
extern void* C_ZNK12QPainterPath10translatedEdd(void* qthis, double arg0, double arg1); // 4
  // proto:  QPainterPath QPainterPath::translated(const QPointF & offset);
extern void* C_ZNK12QPainterPath10translatedERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  QPainterPath QPainterPath::simplified();
extern void* C_ZNK12QPainterPath10simplifiedEv(void* qthis); // 4
  // proto:  void QPainterPath::arcMoveTo(const QRectF & rect, qreal angle);
extern void C_ZN12QPainterPath9arcMoveToERK6QRectFd(void* qthis, void* arg0, double arg1); // 4
  // proto:  void QPainterPath::arcMoveTo(qreal x, qreal y, qreal w, qreal h, qreal angle);
extern void C_ZN12QPainterPath9arcMoveToEddddd(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4); // 2
  // proto:  qreal QPainterPath::percentAtLength(qreal t);
extern double C_ZNK12QPainterPath15percentAtLengthEd(void* qthis, double arg0); // 4
  // proto:  QPainterPath QPainterPathStroker::createStroke(const QPainterPath & path);
extern void* C_ZNK19QPainterPathStroker12createStrokeERK12QPainterPath(void* qthis, void* arg0); // 4
  // proto:  QVector<qreal> QPainterPathStroker::dashPattern();
extern void C_ZNK19QPainterPathStroker11dashPatternEv(void* qthis); // 4
  // proto:  qreal QPainterPathStroker::miterLimit();
extern double C_ZNK19QPainterPathStroker10miterLimitEv(void* qthis); // 4
  // proto:  void QPainterPathStroker::QPainterPathStroker(const QPen & pen);
extern void* C_ZN19QPainterPathStrokerC2ERK4QPen(void* arg0); // 3
  // proto:  void QPainterPathStroker::QPainterPathStroker();
extern void* C_ZN19QPainterPathStrokerC2Ev(); // 3
  // proto:  void QPainterPathStroker::setMiterLimit(qreal length);
extern void C_ZN19QPainterPathStroker13setMiterLimitEd(void* qthis, double arg0); // 4
  // proto:  qreal QPainterPathStroker::dashOffset();
extern double C_ZNK19QPainterPathStroker10dashOffsetEv(void* qthis); // 4
  // proto:  qreal QPainterPathStroker::curveThreshold();
extern double C_ZNK19QPainterPathStroker14curveThresholdEv(void* qthis); // 4
  // proto:  Qt::PenJoinStyle QPainterPathStroker::joinStyle();
extern void C_ZNK19QPainterPathStroker9joinStyleEv(void* qthis); // 4
  // proto:  void QPainterPathStroker::setDashOffset(qreal offset);
extern void C_ZN19QPainterPathStroker13setDashOffsetEd(void* qthis, double arg0); // 4
  // proto:  qreal QPainterPathStroker::width();
extern double C_ZNK19QPainterPathStroker5widthEv(void* qthis); // 4
  // proto:  void QPainterPathStroker::setCurveThreshold(qreal threshold);
extern void C_ZN19QPainterPathStroker17setCurveThresholdEd(void* qthis, double arg0); // 4
  // proto:  void QPainterPathStroker::~QPainterPathStroker();
extern void C_ZN19QPainterPathStrokerD2Ev(void* qthis); // 4
  // proto:  Qt::PenCapStyle QPainterPathStroker::capStyle();
extern void C_ZNK19QPainterPathStroker8capStyleEv(void* qthis); // 4
  // proto:  void QPainterPathStroker::setWidth(qreal width);
extern void C_ZN19QPainterPathStroker8setWidthEd(void* qthis, double arg0); // 4
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

// class sizeof(QPainterPath)=1
type QPainterPath struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QPainterPathStroker)=1
type QPainterPathStroker struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// addRect(const class QRectF &)
func (this *QPainterPath) Addrect(args ...interface{}) () {
  // addRect(const class QRectF &)
  // addRect(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath7addRectERK6QRectF
    // invoke: void addRect(const class QRectF &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QPainterPath7addRectERK6QRectF(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN12QPainterPath7addRectEdddd
    // invoke: void addRect(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN12QPainterPath7addRectEdddd(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QPainterPath", "addRect", args)
  }

  return
}

// controlPointRect()
func (this *QPainterPath) Controlpointrect(args ...interface{}) (ret interface{}) {
  // controlPointRect()
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
    // invoke: _ZNK12QPainterPath16controlPointRectEv
    // invoke: QRectF controlPointRect()
    var ret0 = C.C_ZNK12QPainterPath16controlPointRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainterPath", "controlPointRect", args)
  }

  return
}

// moveTo(qreal, qreal)
func (this *QPainterPath) Moveto(args ...interface{}) () {
  // moveTo(qreal, qreal)
  // moveTo(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath6moveToEdd
    // invoke: void moveTo(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN12QPainterPath6moveToEdd(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN12QPainterPath6moveToERK7QPointF
    // invoke: void moveTo(const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QPainterPath6moveToERK7QPointF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "moveTo", args)
  }

  return
}

// lineTo(const class QPointF &)
func (this *QPainterPath) Lineto(args ...interface{}) () {
  // lineTo(const class QPointF &)
  // lineTo(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath6lineToERK7QPointF
    // invoke: void lineTo(const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QPainterPath6lineToERK7QPointF(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN12QPainterPath6lineToEdd
    // invoke: void lineTo(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN12QPainterPath6lineToEdd(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainterPath", "lineTo", args)
  }

  return
}

// intersects(const class QPainterPath &)
func (this *QPainterPath) Intersects(args ...interface{}) (ret interface{}) {
  // intersects(const class QPainterPath &)
  // intersects(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath10intersectsERKS_
    // invoke: bool intersects(const class QPainterPath &)
    var arg0 = args[0].(*QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QPainterPath10intersectsERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK12QPainterPath10intersectsERK6QRectF
    // invoke: bool intersects(const class QRectF &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QPainterPath10intersectsERK6QRectF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainterPath", "intersects", args)
  }

  return
}

// elementCount()
func (this *QPainterPath) Elementcount(args ...interface{}) (ret interface{}) {
  // elementCount()
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
    // invoke: _ZNK12QPainterPath12elementCountEv
    // invoke: int elementCount()
    var ret0 = C.C_ZNK12QPainterPath12elementCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainterPath", "elementCount", args)
  }

  return
}

// closeSubpath()
func (this *QPainterPath) Closesubpath(args ...interface{}) () {
  // closeSubpath()
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
    // invoke: _ZN12QPainterPath12closeSubpathEv
    // invoke: void closeSubpath()
    C.C_ZN12QPainterPath12closeSubpathEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPainterPath", "closeSubpath", args)
  }

  return
}

// toSubpathPolygons(const class QTransform &)
func (this *QPainterPath) Tosubpathpolygons(args ...interface{}) () {
  // toSubpathPolygons(const class QTransform &)
  // toSubpathPolygons(const class QMatrix &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QMatrix{}) // "const QMatrix &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath17toSubpathPolygonsERK10QTransform
    // invoke: QList<QPolygonF> toSubpathPolygons(const class QTransform &)
    var arg0 = args[0].(*QTransform).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK12QPainterPath17toSubpathPolygonsERK10QTransform(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZNK12QPainterPath17toSubpathPolygonsERK7QMatrix
    // invoke: QList<QPolygonF> toSubpathPolygons(const class QMatrix &)
    var arg0 = args[0].(*QMatrix).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK12QPainterPath17toSubpathPolygonsERK7QMatrix(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "toSubpathPolygons", args)
  }

  return
}

// QPainterPath(const class QPointF &)
func NewQPainterPath(args ...interface{}) *QPainterPath {
  // QPainterPath(const class QPointF &)
  // QPainterPath()
  // QPainterPath(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPathC1ERK7QPointF
    // invoke: void QPainterPath(const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QPainterPathC2ERK7QPointF(arg0)
    return &QPainterPath{Qclsinst:qthis}
  case 1:
    // invoke: _ZN12QPainterPathC1Ev
    // invoke: void QPainterPath()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QPainterPathC2Ev()
    return &QPainterPath{Qclsinst:qthis}
  case 2:
    // invoke: _ZN12QPainterPathC1ERKS_
    // invoke: void QPainterPath(const class QPainterPath &)
    var arg0 = args[0].(*QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QPainterPathC2ERKS_(arg0)
    return &QPainterPath{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QPainterPath", "QPainterPath", args)
  }

  return nil // QPainterPath{Qclsinst:qthis}
}

// angleAtPercent(qreal)
func (this *QPainterPath) Angleatpercent(args ...interface{}) (ret interface{}) {
  // angleAtPercent(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath14angleAtPercentEd
    // invoke: qreal angleAtPercent(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QPainterPath14angleAtPercentEd(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainterPath", "angleAtPercent", args)
  }

  return
}

// addEllipse(const class QRectF &)
func (this *QPainterPath) Addellipse(args ...interface{}) () {
  // addEllipse(const class QRectF &)
  // addEllipse(const class QPointF &, qreal, qreal)
  // addEllipse(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][3] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath10addEllipseERK6QRectF
    // invoke: void addEllipse(const class QRectF &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QPainterPath10addEllipseERK6QRectF(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN12QPainterPath10addEllipseERK7QPointFdd
    // invoke: void addEllipse(const class QPointF &, qreal, qreal)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    C.C_ZN12QPainterPath10addEllipseERK7QPointFdd(this.Qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN12QPainterPath10addEllipseEdddd
    // invoke: void addEllipse(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN12QPainterPath10addEllipseEdddd(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QPainterPath", "addEllipse", args)
  }

  return
}

// quadTo(const class QPointF &, const class QPointF &)
func (this *QPainterPath) Quadto(args ...interface{}) () {
  // quadTo(const class QPointF &, const class QPointF &)
  // quadTo(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[0][1] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath6quadToERK7QPointFS2_
    // invoke: void quadTo(const class QPointF &, const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN12QPainterPath6quadToERK7QPointFS2_(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN12QPainterPath6quadToEdddd
    // invoke: void quadTo(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN12QPainterPath6quadToEdddd(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QPainterPath", "quadTo", args)
  }

  return
}

// fillRule()
func (this *QPainterPath) Fillrule(args ...interface{}) () {
  // fillRule()
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
    // invoke: _ZNK12QPainterPath8fillRuleEv
    // invoke: Qt::FillRule fillRule()
    C.C_ZNK12QPainterPath8fillRuleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPainterPath", "fillRule", args)
  }

  return
}

// intersected(const class QPainterPath &)
func (this *QPainterPath) Intersected(args ...interface{}) (ret interface{}) {
  // intersected(const class QPainterPath &)
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
    // invoke: _ZNK12QPainterPath11intersectedERKS_
    // invoke: QPainterPath intersected(const class QPainterPath &)
    var arg0 = args[0].(*QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QPainterPath11intersectedERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainterPath", "intersected", args)
  }

  return
}

// translate(qreal, qreal)
func (this *QPainterPath) Translate(args ...interface{}) () {
  // translate(qreal, qreal)
  // translate(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath9translateEdd
    // invoke: void translate(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN12QPainterPath9translateEdd(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN12QPainterPath9translateERK7QPointF
    // invoke: void translate(const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QPainterPath9translateERK7QPointF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "translate", args)
  }

  return
}

// contains(const class QPainterPath &)
func (this *QPainterPath) Contains(args ...interface{}) (ret interface{}) {
  // contains(const class QPainterPath &)
  // contains(const class QRectF &)
  // contains(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath8containsERKS_
    // invoke: bool contains(const class QPainterPath &)
    var arg0 = args[0].(*QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QPainterPath8containsERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK12QPainterPath8containsERK6QRectF
    // invoke: bool contains(const class QRectF &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QPainterPath8containsERK6QRectF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK12QPainterPath8containsERK7QPointF
    // invoke: bool contains(const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QPainterPath8containsERK7QPointF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainterPath", "contains", args)
  }

  return
}

// setElementPositionAt(int, qreal, qreal)
func (this *QPainterPath) Setelementpositionat(args ...interface{}) () {
  // setElementPositionAt(int, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath20setElementPositionAtEidd
    // invoke: void setElementPositionAt(int, qreal, qreal)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    C.C_ZN12QPainterPath20setElementPositionAtEidd(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QPainterPath", "setElementPositionAt", args)
  }

  return
}

// ~QPainterPath()
func (this *QPainterPath) Freeqpainterpath(args ...interface{}) () {
  // ~QPainterPath()
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
    // invoke: _ZN12QPainterPathD0Ev
    // invoke: void ~QPainterPath()
    C.C_ZN12QPainterPathD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPainterPath", "~QPainterPath", args)
  }

  return
}

// swap(class QPainterPath &)
func (this *QPainterPath) Swap(args ...interface{}) () {
  // swap(class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "QPainterPath &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath4swapERS_
    // invoke: void swap(class QPainterPath &)
    var arg0 = args[0].(*QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QPainterPath4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "swap", args)
  }

  return
}

// toFillPolygons(const class QMatrix &)
func (this *QPainterPath) Tofillpolygons(args ...interface{}) () {
  // toFillPolygons(const class QMatrix &)
  // toFillPolygons(const class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix{}) // "const QMatrix &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath14toFillPolygonsERK7QMatrix
    // invoke: QList<QPolygonF> toFillPolygons(const class QMatrix &)
    var arg0 = args[0].(*QMatrix).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK12QPainterPath14toFillPolygonsERK7QMatrix(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZNK12QPainterPath14toFillPolygonsERK10QTransform
    // invoke: QList<QPolygonF> toFillPolygons(const class QTransform &)
    var arg0 = args[0].(*QTransform).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK12QPainterPath14toFillPolygonsERK10QTransform(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "toFillPolygons", args)
  }

  return
}

// isEmpty()
func (this *QPainterPath) Isempty(args ...interface{}) (ret interface{}) {
  // isEmpty()
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
    // invoke: _ZNK12QPainterPath7isEmptyEv
    // invoke: bool isEmpty()
    var ret0 = C.C_ZNK12QPainterPath7isEmptyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainterPath", "isEmpty", args)
  }

  return
}

// addRegion(const class QRegion &)
func (this *QPainterPath) Addregion(args ...interface{}) () {
  // addRegion(const class QRegion &)
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
    // invoke: _ZN12QPainterPath9addRegionERK7QRegion
    // invoke: void addRegion(const class QRegion &)
    var arg0 = args[0].(*QRegion).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QPainterPath9addRegionERK7QRegion(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "addRegion", args)
  }

  return
}

// addPath(const class QPainterPath &)
func (this *QPainterPath) Addpath(args ...interface{}) () {
  // addPath(const class QPainterPath &)
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
    // invoke: _ZN12QPainterPath7addPathERKS_
    // invoke: void addPath(const class QPainterPath &)
    var arg0 = args[0].(*QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QPainterPath7addPathERKS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "addPath", args)
  }

  return
}

// addPolygon(const class QPolygonF &)
func (this *QPainterPath) Addpolygon(args ...interface{}) () {
  // addPolygon(const class QPolygonF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath10addPolygonERK9QPolygonF
    // invoke: void addPolygon(const class QPolygonF &)
    var arg0 = args[0].(*QPolygonF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QPainterPath10addPolygonERK9QPolygonF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "addPolygon", args)
  }

  return
}

// arcTo(const class QRectF &, qreal, qreal)
func (this *QPainterPath) Arcto(args ...interface{}) () {
  // arcTo(const class QRectF &, qreal, qreal)
  // arcTo(qreal, qreal, qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][5] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath5arcToERK6QRectFdd
    // invoke: void arcTo(const class QRectF &, qreal, qreal)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    C.C_ZN12QPainterPath5arcToERK6QRectFdd(this.Qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN12QPainterPath5arcToEdddddd
    // invoke: void arcTo(qreal, qreal, qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.double(qtrt.PrimConv(args[4], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg4)}
    var arg5 = C.double(qtrt.PrimConv(args[5], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg5)}
    C.C_ZN12QPainterPath5arcToEdddddd(this.Qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  default:
    qtrt.ErrorResolve("QPainterPath", "arcTo", args)
  }

  return
}

// slopeAtPercent(qreal)
func (this *QPainterPath) Slopeatpercent(args ...interface{}) (ret interface{}) {
  // slopeAtPercent(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath14slopeAtPercentEd
    // invoke: qreal slopeAtPercent(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QPainterPath14slopeAtPercentEd(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainterPath", "slopeAtPercent", args)
  }

  return
}

// addText(const class QPointF &, const class QFont &, const class QString &)
func (this *QPainterPath) Addtext(args ...interface{}) () {
  // addText(const class QPointF &, const class QFont &, const class QString &)
  // addText(qreal, qreal, const class QFont &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[0][1] = reflect.TypeOf(QFont{}) // "const QFont &"
  vtys[0][2] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = reflect.TypeOf(QFont{}) // "const QFont &"
  vtys[1][3] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath7addTextERK7QPointFRK5QFontRK7QString
    // invoke: void addText(const class QPointF &, const class QFont &, const class QString &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QFont).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN12QPainterPath7addTextERK7QPointFRK5QFontRK7QString(this.Qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN12QPainterPath7addTextEddRK5QFontRK7QString
    // invoke: void addText(qreal, qreal, const class QFont &, const class QString &)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QFont).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg3)}
    C.C_ZN12QPainterPath7addTextEddRK5QFontRK7QString(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QPainterPath", "addText", args)
  }

  return
}

// addRoundRect(qreal, qreal, qreal, qreal, int, int)
func (this *QPainterPath) Addroundrect(args ...interface{}) () {
  // addRoundRect(qreal, qreal, qreal, qreal, int, int)
  // addRoundRect(const class QRectF &, int, int)
  // addRoundRect(qreal, qreal, qreal, qreal, int)
  // addRoundRect(const class QRectF &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"
  vtys[0][5] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][4] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath12addRoundRectEddddii
    // invoke: void addRoundRect(qreal, qreal, qreal, qreal, int, int)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(qtrt.PrimConv(args[4], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(qtrt.PrimConv(args[5], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg5)}
    C.C_ZN12QPainterPath12addRoundRectEddddii(this.Qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  case 1:
    // invoke: _ZN12QPainterPath12addRoundRectERK6QRectFii
    // invoke: void addRoundRect(const class QRectF &, int, int)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN12QPainterPath12addRoundRectERK6QRectFii(this.Qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN12QPainterPath12addRoundRectEddddi
    // invoke: void addRoundRect(qreal, qreal, qreal, qreal, int)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(qtrt.PrimConv(args[4], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg4)}
    C.C_ZN12QPainterPath12addRoundRectEddddi(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
  case 3:
    // invoke: _ZN12QPainterPath12addRoundRectERK6QRectFi
    // invoke: void addRoundRect(const class QRectF &, int)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN12QPainterPath12addRoundRectERK6QRectFi(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainterPath", "addRoundRect", args)
  }

  return
}

// elementAt(int)
func (this *QPainterPath) Elementat(args ...interface{}) () {
  // elementAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath9elementAtEi
    // invoke: QPainterPath::Element elementAt(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK12QPainterPath9elementAtEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "elementAt", args)
  }

  return
}

// connectPath(const class QPainterPath &)
func (this *QPainterPath) Connectpath(args ...interface{}) () {
  // connectPath(const class QPainterPath &)
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
    // invoke: _ZN12QPainterPath11connectPathERKS_
    // invoke: void connectPath(const class QPainterPath &)
    var arg0 = args[0].(*QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QPainterPath11connectPathERKS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "connectPath", args)
  }

  return
}

// boundingRect()
func (this *QPainterPath) Boundingrect(args ...interface{}) (ret interface{}) {
  // boundingRect()
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
    // invoke: _ZNK12QPainterPath12boundingRectEv
    // invoke: QRectF boundingRect()
    var ret0 = C.C_ZNK12QPainterPath12boundingRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainterPath", "boundingRect", args)
  }

  return
}

// cubicTo(const class QPointF &, const class QPointF &, const class QPointF &)
func (this *QPainterPath) Cubicto(args ...interface{}) () {
  // cubicTo(const class QPointF &, const class QPointF &, const class QPointF &)
  // cubicTo(qreal, qreal, qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[0][1] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[0][2] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][5] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath7cubicToERK7QPointFS2_S2_
    // invoke: void cubicTo(const class QPointF &, const class QPointF &, const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN12QPainterPath7cubicToERK7QPointFS2_S2_(this.Qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN12QPainterPath7cubicToEdddddd
    // invoke: void cubicTo(qreal, qreal, qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.double(qtrt.PrimConv(args[4], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg4)}
    var arg5 = C.double(qtrt.PrimConv(args[5], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg5)}
    C.C_ZN12QPainterPath7cubicToEdddddd(this.Qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  default:
    qtrt.ErrorResolve("QPainterPath", "cubicTo", args)
  }

  return
}

// toReversed()
func (this *QPainterPath) Toreversed(args ...interface{}) (ret interface{}) {
  // toReversed()
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
    // invoke: _ZNK12QPainterPath10toReversedEv
    // invoke: QPainterPath toReversed()
    var ret0 = C.C_ZNK12QPainterPath10toReversedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainterPath", "toReversed", args)
  }

  return
}

// subtractedInverted(const class QPainterPath &)
func (this *QPainterPath) Subtractedinverted(args ...interface{}) (ret interface{}) {
  // subtractedInverted(const class QPainterPath &)
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
    // invoke: _ZNK12QPainterPath18subtractedInvertedERKS_
    // invoke: QPainterPath subtractedInverted(const class QPainterPath &)
    var arg0 = args[0].(*QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QPainterPath18subtractedInvertedERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainterPath", "subtractedInverted", args)
  }

  return
}

// toFillPolygon(const class QMatrix &)
func (this *QPainterPath) Tofillpolygon(args ...interface{}) (ret interface{}) {
  // toFillPolygon(const class QMatrix &)
  // toFillPolygon(const class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix{}) // "const QMatrix &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath13toFillPolygonERK7QMatrix
    // invoke: QPolygonF toFillPolygon(const class QMatrix &)
    var arg0 = args[0].(*QMatrix).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QPainterPath13toFillPolygonERK7QMatrix(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK12QPainterPath13toFillPolygonERK10QTransform
    // invoke: QPolygonF toFillPolygon(const class QTransform &)
    var arg0 = args[0].(*QTransform).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QPainterPath13toFillPolygonERK10QTransform(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainterPath", "toFillPolygon", args)
  }

  return
}

// subtracted(const class QPainterPath &)
func (this *QPainterPath) Subtracted(args ...interface{}) (ret interface{}) {
  // subtracted(const class QPainterPath &)
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
    // invoke: _ZNK12QPainterPath10subtractedERKS_
    // invoke: QPainterPath subtracted(const class QPainterPath &)
    var arg0 = args[0].(*QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QPainterPath10subtractedERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainterPath", "subtracted", args)
  }

  return
}

// pointAtPercent(qreal)
func (this *QPainterPath) Pointatpercent(args ...interface{}) (ret interface{}) {
  // pointAtPercent(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath14pointAtPercentEd
    // invoke: QPointF pointAtPercent(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QPainterPath14pointAtPercentEd(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainterPath", "pointAtPercent", args)
  }

  return
}

// united(const class QPainterPath &)
func (this *QPainterPath) United(args ...interface{}) (ret interface{}) {
  // united(const class QPainterPath &)
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
    // invoke: _ZNK12QPainterPath6unitedERKS_
    // invoke: QPainterPath united(const class QPainterPath &)
    var arg0 = args[0].(*QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QPainterPath6unitedERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainterPath", "united", args)
  }

  return
}

// currentPosition()
func (this *QPainterPath) Currentposition(args ...interface{}) (ret interface{}) {
  // currentPosition()
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
    // invoke: _ZNK12QPainterPath15currentPositionEv
    // invoke: QPointF currentPosition()
    var ret0 = C.C_ZNK12QPainterPath15currentPositionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainterPath", "currentPosition", args)
  }

  return
}

// length()
func (this *QPainterPath) Length(args ...interface{}) (ret interface{}) {
  // length()
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
    // invoke: _ZNK12QPainterPath6lengthEv
    // invoke: qreal length()
    var ret0 = C.C_ZNK12QPainterPath6lengthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainterPath", "length", args)
  }

  return
}

// translated(qreal, qreal)
func (this *QPainterPath) Translated(args ...interface{}) (ret interface{}) {
  // translated(qreal, qreal)
  // translated(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath10translatedEdd
    // invoke: QPainterPath translated(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK12QPainterPath10translatedEdd(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK12QPainterPath10translatedERK7QPointF
    // invoke: QPainterPath translated(const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QPainterPath10translatedERK7QPointF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainterPath", "translated", args)
  }

  return
}

// simplified()
func (this *QPainterPath) Simplified(args ...interface{}) (ret interface{}) {
  // simplified()
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
    // invoke: _ZNK12QPainterPath10simplifiedEv
    // invoke: QPainterPath simplified()
    var ret0 = C.C_ZNK12QPainterPath10simplifiedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainterPath", "simplified", args)
  }

  return
}

// arcMoveTo(const class QRectF &, qreal)
func (this *QPainterPath) Arcmoveto(args ...interface{}) () {
  // arcMoveTo(const class QRectF &, qreal)
  // arcMoveTo(qreal, qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][4] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath9arcMoveToERK6QRectFd
    // invoke: void arcMoveTo(const class QRectF &, qreal)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN12QPainterPath9arcMoveToERK6QRectFd(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN12QPainterPath9arcMoveToEddddd
    // invoke: void arcMoveTo(qreal, qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.double(qtrt.PrimConv(args[4], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg4)}
    C.C_ZN12QPainterPath9arcMoveToEddddd(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
  default:
    qtrt.ErrorResolve("QPainterPath", "arcMoveTo", args)
  }

  return
}

// percentAtLength(qreal)
func (this *QPainterPath) Percentatlength(args ...interface{}) (ret interface{}) {
  // percentAtLength(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath15percentAtLengthEd
    // invoke: qreal percentAtLength(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QPainterPath15percentAtLengthEd(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainterPath", "percentAtLength", args)
  }

  return
}

// createStroke(const class QPainterPath &)
func (this *QPainterPathStroker) Createstroke(args ...interface{}) (ret interface{}) {
  // createStroke(const class QPainterPath &)
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
    // invoke: _ZNK19QPainterPathStroker12createStrokeERK12QPainterPath
    // invoke: QPainterPath createStroke(const class QPainterPath &)
    var arg0 = args[0].(*QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QPainterPathStroker12createStrokeERK12QPainterPath(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "createStroke", args)
  }

  return
}

// dashPattern()
func (this *QPainterPathStroker) Dashpattern(args ...interface{}) () {
  // dashPattern()
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
    // invoke: _ZNK19QPainterPathStroker11dashPatternEv
    // invoke: QVector<qreal> dashPattern()
    C.C_ZNK19QPainterPathStroker11dashPatternEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "dashPattern", args)
  }

  return
}

// miterLimit()
func (this *QPainterPathStroker) Miterlimit(args ...interface{}) (ret interface{}) {
  // miterLimit()
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
    // invoke: _ZNK19QPainterPathStroker10miterLimitEv
    // invoke: qreal miterLimit()
    var ret0 = C.C_ZNK19QPainterPathStroker10miterLimitEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "miterLimit", args)
  }

  return
}

// QPainterPathStroker(const class QPen &)
func NewQPainterPathStroker(args ...interface{}) *QPainterPathStroker {
  // QPainterPathStroker(const class QPen &)
  // QPainterPathStroker()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPen{}) // "const QPen &"
  vtys[1] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QPainterPathStrokerC1ERK4QPen
    // invoke: void QPainterPathStroker(const class QPen &)
    var arg0 = args[0].(*QPen).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QPainterPathStrokerC2ERK4QPen(arg0)
    return &QPainterPathStroker{Qclsinst:qthis}
  case 1:
    // invoke: _ZN19QPainterPathStrokerC1Ev
    // invoke: void QPainterPathStroker()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QPainterPathStrokerC2Ev()
    return &QPainterPathStroker{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "QPainterPathStroker", args)
  }

  return nil // QPainterPathStroker{Qclsinst:qthis}
}

// setMiterLimit(qreal)
func (this *QPainterPathStroker) Setmiterlimit(args ...interface{}) () {
  // setMiterLimit(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QPainterPathStroker13setMiterLimitEd
    // invoke: void setMiterLimit(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN19QPainterPathStroker13setMiterLimitEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "setMiterLimit", args)
  }

  return
}

// dashOffset()
func (this *QPainterPathStroker) Dashoffset(args ...interface{}) (ret interface{}) {
  // dashOffset()
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
    // invoke: _ZNK19QPainterPathStroker10dashOffsetEv
    // invoke: qreal dashOffset()
    var ret0 = C.C_ZNK19QPainterPathStroker10dashOffsetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "dashOffset", args)
  }

  return
}

// curveThreshold()
func (this *QPainterPathStroker) Curvethreshold(args ...interface{}) (ret interface{}) {
  // curveThreshold()
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
    // invoke: _ZNK19QPainterPathStroker14curveThresholdEv
    // invoke: qreal curveThreshold()
    var ret0 = C.C_ZNK19QPainterPathStroker14curveThresholdEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "curveThreshold", args)
  }

  return
}

// joinStyle()
func (this *QPainterPathStroker) Joinstyle(args ...interface{}) () {
  // joinStyle()
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
    // invoke: _ZNK19QPainterPathStroker9joinStyleEv
    // invoke: Qt::PenJoinStyle joinStyle()
    C.C_ZNK19QPainterPathStroker9joinStyleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "joinStyle", args)
  }

  return
}

// setDashOffset(qreal)
func (this *QPainterPathStroker) Setdashoffset(args ...interface{}) () {
  // setDashOffset(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QPainterPathStroker13setDashOffsetEd
    // invoke: void setDashOffset(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN19QPainterPathStroker13setDashOffsetEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "setDashOffset", args)
  }

  return
}

// width()
func (this *QPainterPathStroker) Width(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK19QPainterPathStroker5widthEv
    // invoke: qreal width()
    var ret0 = C.C_ZNK19QPainterPathStroker5widthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "width", args)
  }

  return
}

// setCurveThreshold(qreal)
func (this *QPainterPathStroker) Setcurvethreshold(args ...interface{}) () {
  // setCurveThreshold(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QPainterPathStroker17setCurveThresholdEd
    // invoke: void setCurveThreshold(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN19QPainterPathStroker17setCurveThresholdEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "setCurveThreshold", args)
  }

  return
}

// ~QPainterPathStroker()
func (this *QPainterPathStroker) Freeqpainterpathstroker(args ...interface{}) () {
  // ~QPainterPathStroker()
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
    // invoke: _ZN19QPainterPathStrokerD0Ev
    // invoke: void ~QPainterPathStroker()
    C.C_ZN19QPainterPathStrokerD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "~QPainterPathStroker", args)
  }

  return
}

// capStyle()
func (this *QPainterPathStroker) Capstyle(args ...interface{}) () {
  // capStyle()
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
    // invoke: _ZNK19QPainterPathStroker8capStyleEv
    // invoke: Qt::PenCapStyle capStyle()
    C.C_ZNK19QPainterPathStroker8capStyleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "capStyle", args)
  }

  return
}

// setWidth(qreal)
func (this *QPainterPathStroker) Setwidth(args ...interface{}) () {
  // setWidth(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QPainterPathStroker8setWidthEd
    // invoke: void setWidth(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN19QPainterPathStroker8setWidthEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "setWidth", args)
  }

  return
}

// <= body block end

