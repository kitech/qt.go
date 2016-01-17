package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QPainterPath::addRect(const QRectF & rect);
extern void _ZN12QPainterPath7addRectERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QPainterPath::addRect(qreal x, qreal y, qreal w, qreal h);
extern void _ZN12QPainterPath7addRectEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  QRectF QPainterPath::controlPointRect();
extern void _ZNK12QPainterPath16controlPointRectEv(void* qthis); // 4
  // proto:  void QPainterPath::moveTo(qreal x, qreal y);
extern void _ZN12QPainterPath6moveToEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QPainterPath::moveTo(const QPointF & p);
extern void _ZN12QPainterPath6moveToERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QPainterPath::lineTo(const QPointF & p);
extern void _ZN12QPainterPath6lineToERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QPainterPath::lineTo(qreal x, qreal y);
extern void _ZN12QPainterPath6lineToEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  bool QPainterPath::intersects(const QPainterPath & p);
extern void _ZNK12QPainterPath10intersectsERKS_(void* qthis, void* arg0); // 4
  // proto:  bool QPainterPath::intersects(const QRectF & rect);
extern void _ZNK12QPainterPath10intersectsERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  int QPainterPath::elementCount();
extern void _ZNK12QPainterPath12elementCountEv(void* qthis); // 4
  // proto:  void QPainterPath::closeSubpath();
extern void _ZN12QPainterPath12closeSubpathEv(void* qthis); // 4
  // proto:  QList<QPolygonF> QPainterPath::toSubpathPolygons(const QTransform & matrix);
extern void _ZNK12QPainterPath17toSubpathPolygonsERK10QTransform(void* qthis, void* arg0); // 4
  // proto:  QList<QPolygonF> QPainterPath::toSubpathPolygons(const QMatrix & matrix);
extern void _ZNK12QPainterPath17toSubpathPolygonsERK7QMatrix(void* qthis, void* arg0); // 4
  // proto:  void QPainterPath::QPainterPath(const QPointF & startPoint);
extern void _ZN12QPainterPathC2ERK7QPointF(void* qthis, void* arg0); // 3
  // proto:  void QPainterPath::QPainterPath();
extern void _ZN12QPainterPathC2Ev(void* qthis); // 3
  // proto:  void QPainterPath::QPainterPath(const QPainterPath & other);
extern void _ZN12QPainterPathC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  qreal QPainterPath::angleAtPercent(qreal t);
extern void _ZNK12QPainterPath14angleAtPercentEd(void* qthis, double arg0); // 4
  // proto:  void QPainterPath::addEllipse(const QRectF & rect);
extern void _ZN12QPainterPath10addEllipseERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QPainterPath::addEllipse(const QPointF & center, qreal rx, qreal ry);
extern void _ZN12QPainterPath10addEllipseERK7QPointFdd(void* qthis, void* arg0, double arg1, double arg2); // 2
  // proto:  void QPainterPath::addEllipse(qreal x, qreal y, qreal w, qreal h);
extern void _ZN12QPainterPath10addEllipseEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  void QPainterPath::quadTo(const QPointF & ctrlPt, const QPointF & endPt);
extern void _ZN12QPainterPath6quadToERK7QPointFS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QPainterPath::quadTo(qreal ctrlPtx, qreal ctrlPty, qreal endPtx, qreal endPty);
extern void _ZN12QPainterPath6quadToEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  Qt::FillRule QPainterPath::fillRule();
extern void _ZNK12QPainterPath8fillRuleEv(void* qthis); // 4
  // proto:  QPainterPath QPainterPath::intersected(const QPainterPath & r);
extern void _ZNK12QPainterPath11intersectedERKS_(void* qthis, void* arg0); // 4
  // proto:  void QPainterPath::translate(qreal dx, qreal dy);
extern void _ZN12QPainterPath9translateEdd(void* qthis, double arg0, double arg1); // 4
  // proto:  void QPainterPath::translate(const QPointF & offset);
extern void _ZN12QPainterPath9translateERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  bool QPainterPath::contains(const QPainterPath & p);
extern void _ZNK12QPainterPath8containsERKS_(void* qthis, void* arg0); // 4
  // proto:  bool QPainterPath::contains(const QRectF & rect);
extern void _ZNK12QPainterPath8containsERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  bool QPainterPath::contains(const QPointF & pt);
extern void _ZNK12QPainterPath8containsERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QPainterPath::setElementPositionAt(int i, qreal x, qreal y);
extern void _ZN12QPainterPath20setElementPositionAtEidd(void* qthis, int32_t arg0, double arg1, double arg2); // 4
  // proto:  void QPainterPath::~QPainterPath();
extern void _ZN12QPainterPathD2Ev(void* qthis); // 4
  // proto:  void QPainterPath::swap(QPainterPath & other);
extern void _ZN12QPainterPath4swapERS_(void* qthis, void* arg0); // 2
  // proto:  QList<QPolygonF> QPainterPath::toFillPolygons(const QMatrix & matrix);
extern void _ZNK12QPainterPath14toFillPolygonsERK7QMatrix(void* qthis, void* arg0); // 4
  // proto:  QList<QPolygonF> QPainterPath::toFillPolygons(const QTransform & matrix);
extern void _ZNK12QPainterPath14toFillPolygonsERK10QTransform(void* qthis, void* arg0); // 4
  // proto:  bool QPainterPath::isEmpty();
extern void _ZNK12QPainterPath7isEmptyEv(void* qthis); // 4
  // proto:  void QPainterPath::addRegion(const QRegion & region);
extern void _ZN12QPainterPath9addRegionERK7QRegion(void* qthis, void* arg0); // 4
  // proto:  void QPainterPath::addPath(const QPainterPath & path);
extern void _ZN12QPainterPath7addPathERKS_(void* qthis, void* arg0); // 4
  // proto:  void QPainterPath::addPolygon(const QPolygonF & polygon);
extern void _ZN12QPainterPath10addPolygonERK9QPolygonF(void* qthis, void* arg0); // 4
  // proto:  void QPainterPath::arcTo(const QRectF & rect, qreal startAngle, qreal arcLength);
extern void _ZN12QPainterPath5arcToERK6QRectFdd(void* qthis, void* arg0, double arg1, double arg2); // 4
  // proto:  void QPainterPath::arcTo(qreal x, qreal y, qreal w, qreal h, qreal startAngle, qreal arcLength);
extern void _ZN12QPainterPath5arcToEdddddd(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4, double arg5); // 2
  // proto:  qreal QPainterPath::slopeAtPercent(qreal t);
extern void _ZNK12QPainterPath14slopeAtPercentEd(void* qthis, double arg0); // 4
  // proto:  void QPainterPath::addText(const QPointF & point, const QFont & f, const QString & text);
extern void _ZN12QPainterPath7addTextERK7QPointFRK5QFontRK7QString(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QPainterPath::addText(qreal x, qreal y, const QFont & f, const QString & text);
extern void _ZN12QPainterPath7addTextEddRK5QFontRK7QString(void* qthis, double arg0, double arg1, void* arg2, void* arg3); // 2
  // proto:  void QPainterPath::addRoundRect(qreal x, qreal y, qreal w, qreal h, int xRnd, int yRnd);
extern void _ZN12QPainterPath12addRoundRectEddddii(void* qthis, double arg0, double arg1, double arg2, double arg3, int32_t arg4, int32_t arg5); // 2
  // proto:  void QPainterPath::addRoundRect(const QRectF & rect, int xRnd, int yRnd);
extern void _ZN12QPainterPath12addRoundRectERK6QRectFii(void* qthis, void* arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  void QPainterPath::addRoundRect(qreal x, qreal y, qreal w, qreal h, int roundness);
extern void _ZN12QPainterPath12addRoundRectEddddi(void* qthis, double arg0, double arg1, double arg2, double arg3, int32_t arg4); // 2
  // proto:  void QPainterPath::addRoundRect(const QRectF & rect, int roundness);
extern void _ZN12QPainterPath12addRoundRectERK6QRectFi(void* qthis, void* arg0, int32_t arg1); // 2
  // proto:  QPainterPath::Element QPainterPath::elementAt(int i);
extern void _ZNK12QPainterPath9elementAtEi(void* qthis, int32_t arg0); // 4
  // proto:  void QPainterPath::connectPath(const QPainterPath & path);
extern void _ZN12QPainterPath11connectPathERKS_(void* qthis, void* arg0); // 4
  // proto:  QRectF QPainterPath::boundingRect();
extern void _ZNK12QPainterPath12boundingRectEv(void* qthis); // 4
  // proto:  void QPainterPath::cubicTo(const QPointF & ctrlPt1, const QPointF & ctrlPt2, const QPointF & endPt);
extern void _ZN12QPainterPath7cubicToERK7QPointFS2_S2_(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QPainterPath::cubicTo(qreal ctrlPt1x, qreal ctrlPt1y, qreal ctrlPt2x, qreal ctrlPt2y, qreal endPtx, qreal endPty);
extern void _ZN12QPainterPath7cubicToEdddddd(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4, double arg5); // 2
  // proto:  QPainterPath QPainterPath::toReversed();
extern void _ZNK12QPainterPath10toReversedEv(void* qthis); // 4
  // proto:  QPainterPath QPainterPath::subtractedInverted(const QPainterPath & r);
extern void _ZNK12QPainterPath18subtractedInvertedERKS_(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QPainterPath::toFillPolygon(const QMatrix & matrix);
extern void _ZNK12QPainterPath13toFillPolygonERK7QMatrix(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QPainterPath::toFillPolygon(const QTransform & matrix);
extern void _ZNK12QPainterPath13toFillPolygonERK10QTransform(void* qthis, void* arg0); // 4
  // proto:  QPainterPath QPainterPath::subtracted(const QPainterPath & r);
extern void _ZNK12QPainterPath10subtractedERKS_(void* qthis, void* arg0); // 4
  // proto:  QPointF QPainterPath::pointAtPercent(qreal t);
extern void _ZNK12QPainterPath14pointAtPercentEd(void* qthis, double arg0); // 4
  // proto:  QPainterPath QPainterPath::united(const QPainterPath & r);
extern void _ZNK12QPainterPath6unitedERKS_(void* qthis, void* arg0); // 4
  // proto:  QPointF QPainterPath::currentPosition();
extern void _ZNK12QPainterPath15currentPositionEv(void* qthis); // 4
  // proto:  qreal QPainterPath::length();
extern void _ZNK12QPainterPath6lengthEv(void* qthis); // 4
  // proto:  QPainterPath QPainterPath::translated(qreal dx, qreal dy);
extern void _ZNK12QPainterPath10translatedEdd(void* qthis, double arg0, double arg1); // 4
  // proto:  QPainterPath QPainterPath::translated(const QPointF & offset);
extern void _ZNK12QPainterPath10translatedERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  QPainterPath QPainterPath::simplified();
extern void _ZNK12QPainterPath10simplifiedEv(void* qthis); // 4
  // proto:  void QPainterPath::arcMoveTo(const QRectF & rect, qreal angle);
extern void _ZN12QPainterPath9arcMoveToERK6QRectFd(void* qthis, void* arg0, double arg1); // 4
  // proto:  void QPainterPath::arcMoveTo(qreal x, qreal y, qreal w, qreal h, qreal angle);
extern void _ZN12QPainterPath9arcMoveToEddddd(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4); // 2
  // proto:  qreal QPainterPath::percentAtLength(qreal t);
extern void _ZNK12QPainterPath15percentAtLengthEd(void* qthis, double arg0); // 4
  // proto:  QPainterPath QPainterPathStroker::createStroke(const QPainterPath & path);
extern void _ZNK19QPainterPathStroker12createStrokeERK12QPainterPath(void* qthis, void* arg0); // 4
  // proto:  QVector<qreal> QPainterPathStroker::dashPattern();
extern void _ZNK19QPainterPathStroker11dashPatternEv(void* qthis); // 4
  // proto:  qreal QPainterPathStroker::miterLimit();
extern void _ZNK19QPainterPathStroker10miterLimitEv(void* qthis); // 4
  // proto:  void QPainterPathStroker::QPainterPathStroker(const QPen & pen);
extern void _ZN19QPainterPathStrokerC2ERK4QPen(void* qthis, void* arg0); // 3
  // proto:  void QPainterPathStroker::QPainterPathStroker();
extern void _ZN19QPainterPathStrokerC2Ev(void* qthis); // 3
  // proto:  void QPainterPathStroker::setMiterLimit(qreal length);
extern void _ZN19QPainterPathStroker13setMiterLimitEd(void* qthis, double arg0); // 4
  // proto:  qreal QPainterPathStroker::dashOffset();
extern void _ZNK19QPainterPathStroker10dashOffsetEv(void* qthis); // 4
  // proto:  qreal QPainterPathStroker::curveThreshold();
extern void _ZNK19QPainterPathStroker14curveThresholdEv(void* qthis); // 4
  // proto:  Qt::PenJoinStyle QPainterPathStroker::joinStyle();
extern void _ZNK19QPainterPathStroker9joinStyleEv(void* qthis); // 4
  // proto:  void QPainterPathStroker::setDashOffset(qreal offset);
extern void _ZN19QPainterPathStroker13setDashOffsetEd(void* qthis, double arg0); // 4
  // proto:  qreal QPainterPathStroker::width();
extern void _ZNK19QPainterPathStroker5widthEv(void* qthis); // 4
  // proto:  void QPainterPathStroker::setCurveThreshold(qreal threshold);
extern void _ZN19QPainterPathStroker17setCurveThresholdEd(void* qthis, double arg0); // 4
  // proto:  void QPainterPathStroker::~QPainterPathStroker();
extern void _ZN19QPainterPathStrokerD2Ev(void* qthis); // 4
  // proto:  Qt::PenCapStyle QPainterPathStroker::capStyle();
extern void _ZNK19QPainterPathStroker8capStyleEv(void* qthis); // 4
  // proto:  void QPainterPathStroker::setWidth(qreal width);
extern void _ZN19QPainterPathStroker8setWidthEd(void* qthis, double arg0); // 4
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

// class sizeof(QPainterPath)=1
type QPainterPath struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QPainterPathStroker)=1
type QPainterPathStroker struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// addRect(const class QRectF &)
func (this *QPainterPath) addRect(args ...interface{}) () {
  // addRect(const class QRectF &)
  // addRect(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath7addRectERK6QRectF
    // invoke: void addRect(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN12QPainterPath7addRectERK6QRectF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN12QPainterPath7addRectEdddd
    // invoke: void addRect(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    C._ZN12QPainterPath7addRectEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QPainterPath", "addRect", args)
  }

}

// controlPointRect()
func (this *QPainterPath) controlPointRect(args ...interface{}) () {
  // controlPointRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath16controlPointRectEv
    // invoke: QRectF controlPointRect()
    C._ZNK12QPainterPath16controlPointRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainterPath", "controlPointRect", args)
  }

}

// moveTo(qreal, qreal)
func (this *QPainterPath) moveTo(args ...interface{}) () {
  // moveTo(qreal, qreal)
  // moveTo(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath6moveToEdd
    // invoke: void moveTo(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C._ZN12QPainterPath6moveToEdd(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN12QPainterPath6moveToERK7QPointF
    // invoke: void moveTo(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN12QPainterPath6moveToERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "moveTo", args)
  }

}

// lineTo(const class QPointF &)
func (this *QPainterPath) lineTo(args ...interface{}) () {
  // lineTo(const class QPointF &)
  // lineTo(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath6lineToERK7QPointF
    // invoke: void lineTo(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN12QPainterPath6lineToERK7QPointF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN12QPainterPath6lineToEdd
    // invoke: void lineTo(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C._ZN12QPainterPath6lineToEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainterPath", "lineTo", args)
  }

}

// intersects(const class QPainterPath &)
func (this *QPainterPath) intersects(args ...interface{}) () {
  // intersects(const class QPainterPath &)
  // intersects(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath10intersectsERKS_
    // invoke: bool intersects(const class QPainterPath &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK12QPainterPath10intersectsERKS_(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK12QPainterPath10intersectsERK6QRectF
    // invoke: bool intersects(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK12QPainterPath10intersectsERK6QRectF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "intersects", args)
  }

}

// elementCount()
func (this *QPainterPath) elementCount(args ...interface{}) () {
  // elementCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath12elementCountEv
    // invoke: int elementCount()
    C._ZNK12QPainterPath12elementCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainterPath", "elementCount", args)
  }

}

// closeSubpath()
func (this *QPainterPath) closeSubpath(args ...interface{}) () {
  // closeSubpath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath12closeSubpathEv
    // invoke: void closeSubpath()
    C._ZN12QPainterPath12closeSubpathEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainterPath", "closeSubpath", args)
  }

}

// toSubpathPolygons(const class QTransform &)
func (this *QPainterPath) toSubpathPolygons(args ...interface{}) () {
  // toSubpathPolygons(const class QTransform &)
  // toSubpathPolygons(const class QMatrix &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QMatrix{}) // "const QMatrix &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath17toSubpathPolygonsERK10QTransform
    // invoke: QList<QPolygonF> toSubpathPolygons(const class QTransform &)
    var arg0 = args[0].(QTransform).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK12QPainterPath17toSubpathPolygonsERK10QTransform(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK12QPainterPath17toSubpathPolygonsERK7QMatrix
    // invoke: QList<QPolygonF> toSubpathPolygons(const class QMatrix &)
    var arg0 = args[0].(QMatrix).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK12QPainterPath17toSubpathPolygonsERK7QMatrix(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "toSubpathPolygons", args)
  }

}

// QPainterPath(const class QPointF &)
func NewQPainterPath(args ...interface{}) QPainterPath {
  // QPainterPath(const class QPointF &)
  // QPainterPath()
  // QPainterPath(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPathC1ERK7QPointF
    // invoke: void QPainterPath(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN12QPainterPathC2ERK7QPointF(qthis, arg0)
  case 1:
    // invoke: _ZN12QPainterPathC1Ev
    // invoke: void QPainterPath()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN12QPainterPathC2Ev(qthis)
  case 2:
    // invoke: _ZN12QPainterPathC1ERKS_
    // invoke: void QPainterPath(const class QPainterPath &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN12QPainterPathC2ERKS_(qthis, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "QPainterPath", args)
  }

  return QPainterPath{}
}

// angleAtPercent(qreal)
func (this *QPainterPath) angleAtPercent(args ...interface{}) () {
  // angleAtPercent(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath14angleAtPercentEd
    // invoke: qreal angleAtPercent(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZNK12QPainterPath14angleAtPercentEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "angleAtPercent", args)
  }

}

// addEllipse(const class QRectF &)
func (this *QPainterPath) addEllipse(args ...interface{}) () {
  // addEllipse(const class QRectF &)
  // addEllipse(const class QPointF &, qreal, qreal)
  // addEllipse(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath10addEllipseERK6QRectF
    // invoke: void addEllipse(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN12QPainterPath10addEllipseERK6QRectF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN12QPainterPath10addEllipseERK7QPointFdd
    // invoke: void addEllipse(const class QPointF &, qreal, qreal)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    C._ZN12QPainterPath10addEllipseERK7QPointFdd(this.qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN12QPainterPath10addEllipseEdddd
    // invoke: void addEllipse(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    C._ZN12QPainterPath10addEllipseEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QPainterPath", "addEllipse", args)
  }

}

// quadTo(const class QPointF &, const class QPointF &)
func (this *QPainterPath) quadTo(args ...interface{}) () {
  // quadTo(const class QPointF &, const class QPointF &)
  // quadTo(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath6quadToERK7QPointFS2_
    // invoke: void quadTo(const class QPointF &, const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPointF).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN12QPainterPath6quadToERK7QPointFS2_(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN12QPainterPath6quadToEdddd
    // invoke: void quadTo(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    C._ZN12QPainterPath6quadToEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QPainterPath", "quadTo", args)
  }

}

// fillRule()
func (this *QPainterPath) fillRule(args ...interface{}) () {
  // fillRule()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath8fillRuleEv
    // invoke: Qt::FillRule fillRule()
    C._ZNK12QPainterPath8fillRuleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainterPath", "fillRule", args)
  }

}

// intersected(const class QPainterPath &)
func (this *QPainterPath) intersected(args ...interface{}) () {
  // intersected(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath11intersectedERKS_
    // invoke: QPainterPath intersected(const class QPainterPath &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK12QPainterPath11intersectedERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "intersected", args)
  }

}

// translate(qreal, qreal)
func (this *QPainterPath) translate(args ...interface{}) () {
  // translate(qreal, qreal)
  // translate(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath9translateEdd
    // invoke: void translate(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C._ZN12QPainterPath9translateEdd(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN12QPainterPath9translateERK7QPointF
    // invoke: void translate(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN12QPainterPath9translateERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "translate", args)
  }

}

// contains(const class QPainterPath &)
func (this *QPainterPath) contains(args ...interface{}) () {
  // contains(const class QPainterPath &)
  // contains(const class QRectF &)
  // contains(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath8containsERKS_
    // invoke: bool contains(const class QPainterPath &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK12QPainterPath8containsERKS_(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK12QPainterPath8containsERK6QRectF
    // invoke: bool contains(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK12QPainterPath8containsERK6QRectF(this.qclsinst, arg0)
  case 2:
    // invoke: _ZNK12QPainterPath8containsERK7QPointF
    // invoke: bool contains(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK12QPainterPath8containsERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "contains", args)
  }

}

// setElementPositionAt(int, qreal, qreal)
func (this *QPainterPath) setElementPositionAt(args ...interface{}) () {
  // setElementPositionAt(int, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath20setElementPositionAtEidd
    // invoke: void setElementPositionAt(int, qreal, qreal)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    C._ZN12QPainterPath20setElementPositionAtEidd(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QPainterPath", "setElementPositionAt", args)
  }

}

// ~QPainterPath()
func (this *QPainterPath) FreeQPainterPath(args ...interface{}) () {
  // ~QPainterPath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPathD0Ev
    // invoke: void ~QPainterPath()
    C._ZN12QPainterPathD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainterPath", "~QPainterPath", args)
  }

}

// swap(class QPainterPath &)
func (this *QPainterPath) swap(args ...interface{}) () {
  // swap(class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath4swapERS_
    // invoke: void swap(class QPainterPath &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN12QPainterPath4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "swap", args)
  }

}

// toFillPolygons(const class QMatrix &)
func (this *QPainterPath) toFillPolygons(args ...interface{}) () {
  // toFillPolygons(const class QMatrix &)
  // toFillPolygons(const class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix{}) // "const QMatrix &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath14toFillPolygonsERK7QMatrix
    // invoke: QList<QPolygonF> toFillPolygons(const class QMatrix &)
    var arg0 = args[0].(QMatrix).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK12QPainterPath14toFillPolygonsERK7QMatrix(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK12QPainterPath14toFillPolygonsERK10QTransform
    // invoke: QList<QPolygonF> toFillPolygons(const class QTransform &)
    var arg0 = args[0].(QTransform).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK12QPainterPath14toFillPolygonsERK10QTransform(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "toFillPolygons", args)
  }

}

// isEmpty()
func (this *QPainterPath) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath7isEmptyEv
    // invoke: bool isEmpty()
    C._ZNK12QPainterPath7isEmptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainterPath", "isEmpty", args)
  }

}

// addRegion(const class QRegion &)
func (this *QPainterPath) addRegion(args ...interface{}) () {
  // addRegion(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath9addRegionERK7QRegion
    // invoke: void addRegion(const class QRegion &)
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN12QPainterPath9addRegionERK7QRegion(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "addRegion", args)
  }

}

// addPath(const class QPainterPath &)
func (this *QPainterPath) addPath(args ...interface{}) () {
  // addPath(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath7addPathERKS_
    // invoke: void addPath(const class QPainterPath &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN12QPainterPath7addPathERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "addPath", args)
  }

}

// addPolygon(const class QPolygonF &)
func (this *QPainterPath) addPolygon(args ...interface{}) () {
  // addPolygon(const class QPolygonF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath10addPolygonERK9QPolygonF
    // invoke: void addPolygon(const class QPolygonF &)
    var arg0 = args[0].(QPolygonF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN12QPainterPath10addPolygonERK9QPolygonF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "addPolygon", args)
  }

}

// arcTo(const class QRectF &, qreal, qreal)
func (this *QPainterPath) arcTo(args ...interface{}) () {
  // arcTo(const class QRectF &, qreal, qreal)
  // arcTo(qreal, qreal, qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][5] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath5arcToERK6QRectFdd
    // invoke: void arcTo(const class QRectF &, qreal, qreal)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    C._ZN12QPainterPath5arcToERK6QRectFdd(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN12QPainterPath5arcToEdddddd
    // invoke: void arcTo(qreal, qreal, qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.double(args[4].(float64))
    if false {fmt.Println(arg4)}
    var arg5 = C.double(args[5].(float64))
    if false {fmt.Println(arg5)}
    C._ZN12QPainterPath5arcToEdddddd(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  default:
    qtrt.ErrorResolve("QPainterPath", "arcTo", args)
  }

}

// slopeAtPercent(qreal)
func (this *QPainterPath) slopeAtPercent(args ...interface{}) () {
  // slopeAtPercent(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath14slopeAtPercentEd
    // invoke: qreal slopeAtPercent(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZNK12QPainterPath14slopeAtPercentEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "slopeAtPercent", args)
  }

}

// addText(const class QPointF &, const class QFont &, const class QString &)
func (this *QPainterPath) addText(args ...interface{}) () {
  // addText(const class QPointF &, const class QFont &, const class QString &)
  // addText(qreal, qreal, const class QFont &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][1] = reflect.TypeOf(QFont{}) // "const QFont &"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = reflect.TypeOf(QFont{}) // "const QFont &"
  vtys[1][3] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath7addTextERK7QPointFRK5QFontRK7QString
    // invoke: void addText(const class QPointF &, const class QFont &, const class QString &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QFont).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN12QPainterPath7addTextERK7QPointFRK5QFontRK7QString(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN12QPainterPath7addTextEddRK5QFontRK7QString
    // invoke: void addText(qreal, qreal, const class QFont &, const class QString &)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QFont).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QString).qclsinst
    if false {fmt.Println(arg3)}
    C._ZN12QPainterPath7addTextEddRK5QFontRK7QString(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QPainterPath", "addText", args)
  }

}

// addRoundRect(qreal, qreal, qreal, qreal, int, int)
func (this *QPainterPath) addRoundRect(args ...interface{}) () {
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
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][4] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath12addRoundRectEddddii
    // invoke: void addRoundRect(qreal, qreal, qreal, qreal, int, int)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(args[5].(int32))
    if false {fmt.Println(arg5)}
    C._ZN12QPainterPath12addRoundRectEddddii(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  case 1:
    // invoke: _ZN12QPainterPath12addRoundRectERK6QRectFii
    // invoke: void addRoundRect(const class QRectF &, int, int)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN12QPainterPath12addRoundRectERK6QRectFii(this.qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN12QPainterPath12addRoundRectEddddi
    // invoke: void addRoundRect(qreal, qreal, qreal, qreal, int)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    C._ZN12QPainterPath12addRoundRectEddddi(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
  case 3:
    // invoke: _ZN12QPainterPath12addRoundRectERK6QRectFi
    // invoke: void addRoundRect(const class QRectF &, int)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN12QPainterPath12addRoundRectERK6QRectFi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainterPath", "addRoundRect", args)
  }

}

// elementAt(int)
func (this *QPainterPath) elementAt(args ...interface{}) () {
  // elementAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath9elementAtEi
    // invoke: QPainterPath::Element elementAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK12QPainterPath9elementAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "elementAt", args)
  }

}

// connectPath(const class QPainterPath &)
func (this *QPainterPath) connectPath(args ...interface{}) () {
  // connectPath(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath11connectPathERKS_
    // invoke: void connectPath(const class QPainterPath &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN12QPainterPath11connectPathERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "connectPath", args)
  }

}

// boundingRect()
func (this *QPainterPath) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath12boundingRectEv
    // invoke: QRectF boundingRect()
    C._ZNK12QPainterPath12boundingRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainterPath", "boundingRect", args)
  }

}

// cubicTo(const class QPointF &, const class QPointF &, const class QPointF &)
func (this *QPainterPath) cubicTo(args ...interface{}) () {
  // cubicTo(const class QPointF &, const class QPointF &, const class QPointF &)
  // cubicTo(qreal, qreal, qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][2] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][5] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath7cubicToERK7QPointFS2_S2_
    // invoke: void cubicTo(const class QPointF &, const class QPointF &, const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPointF).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QPointF).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN12QPainterPath7cubicToERK7QPointFS2_S2_(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN12QPainterPath7cubicToEdddddd
    // invoke: void cubicTo(qreal, qreal, qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.double(args[4].(float64))
    if false {fmt.Println(arg4)}
    var arg5 = C.double(args[5].(float64))
    if false {fmt.Println(arg5)}
    C._ZN12QPainterPath7cubicToEdddddd(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  default:
    qtrt.ErrorResolve("QPainterPath", "cubicTo", args)
  }

}

// toReversed()
func (this *QPainterPath) toReversed(args ...interface{}) () {
  // toReversed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath10toReversedEv
    // invoke: QPainterPath toReversed()
    C._ZNK12QPainterPath10toReversedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainterPath", "toReversed", args)
  }

}

// subtractedInverted(const class QPainterPath &)
func (this *QPainterPath) subtractedInverted(args ...interface{}) () {
  // subtractedInverted(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath18subtractedInvertedERKS_
    // invoke: QPainterPath subtractedInverted(const class QPainterPath &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK12QPainterPath18subtractedInvertedERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "subtractedInverted", args)
  }

}

// toFillPolygon(const class QMatrix &)
func (this *QPainterPath) toFillPolygon(args ...interface{}) () {
  // toFillPolygon(const class QMatrix &)
  // toFillPolygon(const class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix{}) // "const QMatrix &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath13toFillPolygonERK7QMatrix
    // invoke: QPolygonF toFillPolygon(const class QMatrix &)
    var arg0 = args[0].(QMatrix).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK12QPainterPath13toFillPolygonERK7QMatrix(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK12QPainterPath13toFillPolygonERK10QTransform
    // invoke: QPolygonF toFillPolygon(const class QTransform &)
    var arg0 = args[0].(QTransform).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK12QPainterPath13toFillPolygonERK10QTransform(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "toFillPolygon", args)
  }

}

// subtracted(const class QPainterPath &)
func (this *QPainterPath) subtracted(args ...interface{}) () {
  // subtracted(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath10subtractedERKS_
    // invoke: QPainterPath subtracted(const class QPainterPath &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK12QPainterPath10subtractedERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "subtracted", args)
  }

}

// pointAtPercent(qreal)
func (this *QPainterPath) pointAtPercent(args ...interface{}) () {
  // pointAtPercent(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath14pointAtPercentEd
    // invoke: QPointF pointAtPercent(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZNK12QPainterPath14pointAtPercentEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "pointAtPercent", args)
  }

}

// united(const class QPainterPath &)
func (this *QPainterPath) united(args ...interface{}) () {
  // united(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath6unitedERKS_
    // invoke: QPainterPath united(const class QPainterPath &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK12QPainterPath6unitedERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "united", args)
  }

}

// currentPosition()
func (this *QPainterPath) currentPosition(args ...interface{}) () {
  // currentPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath15currentPositionEv
    // invoke: QPointF currentPosition()
    C._ZNK12QPainterPath15currentPositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainterPath", "currentPosition", args)
  }

}

// length()
func (this *QPainterPath) length(args ...interface{}) () {
  // length()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath6lengthEv
    // invoke: qreal length()
    C._ZNK12QPainterPath6lengthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainterPath", "length", args)
  }

}

// translated(qreal, qreal)
func (this *QPainterPath) translated(args ...interface{}) () {
  // translated(qreal, qreal)
  // translated(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath10translatedEdd
    // invoke: QPainterPath translated(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C._ZNK12QPainterPath10translatedEdd(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK12QPainterPath10translatedERK7QPointF
    // invoke: QPainterPath translated(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK12QPainterPath10translatedERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "translated", args)
  }

}

// simplified()
func (this *QPainterPath) simplified(args ...interface{}) () {
  // simplified()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath10simplifiedEv
    // invoke: QPainterPath simplified()
    C._ZNK12QPainterPath10simplifiedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainterPath", "simplified", args)
  }

}

// arcMoveTo(const class QRectF &, qreal)
func (this *QPainterPath) arcMoveTo(args ...interface{}) () {
  // arcMoveTo(const class QRectF &, qreal)
  // arcMoveTo(qreal, qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][4] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath9arcMoveToERK6QRectFd
    // invoke: void arcMoveTo(const class QRectF &, qreal)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C._ZN12QPainterPath9arcMoveToERK6QRectFd(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN12QPainterPath9arcMoveToEddddd
    // invoke: void arcMoveTo(qreal, qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.double(args[4].(float64))
    if false {fmt.Println(arg4)}
    C._ZN12QPainterPath9arcMoveToEddddd(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
  default:
    qtrt.ErrorResolve("QPainterPath", "arcMoveTo", args)
  }

}

// percentAtLength(qreal)
func (this *QPainterPath) percentAtLength(args ...interface{}) () {
  // percentAtLength(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath15percentAtLengthEd
    // invoke: qreal percentAtLength(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZNK12QPainterPath15percentAtLengthEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPath", "percentAtLength", args)
  }

}

// createStroke(const class QPainterPath &)
func (this *QPainterPathStroker) createStroke(args ...interface{}) () {
  // createStroke(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QPainterPathStroker12createStrokeERK12QPainterPath
    // invoke: QPainterPath createStroke(const class QPainterPath &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK19QPainterPathStroker12createStrokeERK12QPainterPath(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "createStroke", args)
  }

}

// dashPattern()
func (this *QPainterPathStroker) dashPattern(args ...interface{}) () {
  // dashPattern()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QPainterPathStroker11dashPatternEv
    // invoke: QVector<qreal> dashPattern()
    C._ZNK19QPainterPathStroker11dashPatternEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "dashPattern", args)
  }

}

// miterLimit()
func (this *QPainterPathStroker) miterLimit(args ...interface{}) () {
  // miterLimit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QPainterPathStroker10miterLimitEv
    // invoke: qreal miterLimit()
    C._ZNK19QPainterPathStroker10miterLimitEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "miterLimit", args)
  }

}

// QPainterPathStroker(const class QPen &)
func NewQPainterPathStroker(args ...interface{}) QPainterPathStroker {
  // QPainterPathStroker(const class QPen &)
  // QPainterPathStroker()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPen{}) // "const QPen &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QPainterPathStrokerC1ERK4QPen
    // invoke: void QPainterPathStroker(const class QPen &)
    var arg0 = args[0].(QPen).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN19QPainterPathStrokerC2ERK4QPen(qthis, arg0)
  case 1:
    // invoke: _ZN19QPainterPathStrokerC1Ev
    // invoke: void QPainterPathStroker()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN19QPainterPathStrokerC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "QPainterPathStroker", args)
  }

  return QPainterPathStroker{}
}

// setMiterLimit(qreal)
func (this *QPainterPathStroker) setMiterLimit(args ...interface{}) () {
  // setMiterLimit(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QPainterPathStroker13setMiterLimitEd
    // invoke: void setMiterLimit(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN19QPainterPathStroker13setMiterLimitEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "setMiterLimit", args)
  }

}

// dashOffset()
func (this *QPainterPathStroker) dashOffset(args ...interface{}) () {
  // dashOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QPainterPathStroker10dashOffsetEv
    // invoke: qreal dashOffset()
    C._ZNK19QPainterPathStroker10dashOffsetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "dashOffset", args)
  }

}

// curveThreshold()
func (this *QPainterPathStroker) curveThreshold(args ...interface{}) () {
  // curveThreshold()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QPainterPathStroker14curveThresholdEv
    // invoke: qreal curveThreshold()
    C._ZNK19QPainterPathStroker14curveThresholdEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "curveThreshold", args)
  }

}

// joinStyle()
func (this *QPainterPathStroker) joinStyle(args ...interface{}) () {
  // joinStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QPainterPathStroker9joinStyleEv
    // invoke: Qt::PenJoinStyle joinStyle()
    C._ZNK19QPainterPathStroker9joinStyleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "joinStyle", args)
  }

}

// setDashOffset(qreal)
func (this *QPainterPathStroker) setDashOffset(args ...interface{}) () {
  // setDashOffset(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QPainterPathStroker13setDashOffsetEd
    // invoke: void setDashOffset(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN19QPainterPathStroker13setDashOffsetEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "setDashOffset", args)
  }

}

// width()
func (this *QPainterPathStroker) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QPainterPathStroker5widthEv
    // invoke: qreal width()
    C._ZNK19QPainterPathStroker5widthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "width", args)
  }

}

// setCurveThreshold(qreal)
func (this *QPainterPathStroker) setCurveThreshold(args ...interface{}) () {
  // setCurveThreshold(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QPainterPathStroker17setCurveThresholdEd
    // invoke: void setCurveThreshold(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN19QPainterPathStroker17setCurveThresholdEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "setCurveThreshold", args)
  }

}

// ~QPainterPathStroker()
func (this *QPainterPathStroker) FreeQPainterPathStroker(args ...interface{}) () {
  // ~QPainterPathStroker()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QPainterPathStrokerD0Ev
    // invoke: void ~QPainterPathStroker()
    C._ZN19QPainterPathStrokerD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "~QPainterPathStroker", args)
  }

}

// capStyle()
func (this *QPainterPathStroker) capStyle(args ...interface{}) () {
  // capStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QPainterPathStroker8capStyleEv
    // invoke: Qt::PenCapStyle capStyle()
    C._ZNK19QPainterPathStroker8capStyleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "capStyle", args)
  }

}

// setWidth(qreal)
func (this *QPainterPathStroker) setWidth(args ...interface{}) () {
  // setWidth(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QPainterPathStroker8setWidthEd
    // invoke: void setWidth(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN19QPainterPathStroker8setWidthEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "setWidth", args)
  }

}

// <= body block end

