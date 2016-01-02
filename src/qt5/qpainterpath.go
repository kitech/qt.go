package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
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
  // proto:  void QPainterPath::setElementPositionAt(int i, qreal x, qreal y);
extern void _ZN12QPainterPath20setElementPositionAtEidd(void* qthis, int arg0, double arg1, double arg2);
  // proto:  QPolygonF QPainterPath::toFillPolygon(const QMatrix & matrix);
extern void _ZNK12QPainterPath13toFillPolygonERK7QMatrix(void* qthis, void* arg0);
  // proto:  QPainterPath QPainterPath::translated(qreal dx, qreal dy);
extern void _ZNK12QPainterPath10translatedEdd(void* qthis, double arg0, double arg1);
  // proto:  void QPainterPath::~QPainterPath();
extern void _ZN12QPainterPathD0Ev(void* qthis);
  // proto:  QList<QPolygonF> QPainterPath::toSubpathPolygons(const QTransform & matrix);
extern void _ZNK12QPainterPath17toSubpathPolygonsERK10QTransform(void* qthis, void* arg0);
  // proto:  QRectF QPainterPath::controlPointRect();
extern void _ZNK12QPainterPath16controlPointRectEv(void* qthis);
  // proto:  QList<QPolygonF> QPainterPath::toFillPolygons(const QMatrix & matrix);
extern void _ZNK12QPainterPath14toFillPolygonsERK7QMatrix(void* qthis, void* arg0);
  // proto:  QPainterPath QPainterPath::translated(const QPointF & offset);
extern void demth_ZNK12QPainterPath10translatedERK7QPointF(void* qthis, void* arg0);
  // proto:  void QPainterPath::quadTo(const QPointF & ctrlPt, const QPointF & endPt);
extern void _ZN12QPainterPath6quadToERK7QPointFS2_(void* qthis, void* arg0, void* arg1);
  // proto:  QList<QPolygonF> QPainterPath::toFillPolygons(const QTransform & matrix);
extern void _ZNK12QPainterPath14toFillPolygonsERK10QTransform(void* qthis, void* arg0);
  // proto:  void QPainterPath::arcTo(qreal x, qreal y, qreal w, qreal h, qreal startAngle, qreal arcLength);
extern void demth_ZN12QPainterPath5arcToEdddddd(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4, double arg5);
  // proto:  void QPainterPath::addRect(const QRectF & rect);
extern void _ZN12QPainterPath7addRectERK6QRectF(void* qthis, void* arg0);
  // proto:  void QPainterPath::addRoundRect(const QRectF & rect, int xRnd, int yRnd);
extern void _ZN12QPainterPath12addRoundRectERK6QRectFii(void* qthis, void* arg0, int arg1, int arg2);
  // proto:  void QPainterPath::addText(qreal x, qreal y, const QFont & f, const QString & text);
extern void demth_ZN12QPainterPath7addTextEddRK5QFontRK7QString(void* qthis, double arg0, double arg1, void* arg2, void* arg3);
  // proto:  bool QPainterPath::intersects(const QRectF & rect);
extern void _ZNK12QPainterPath10intersectsERK6QRectF(void* qthis, void* arg0);
  // proto:  bool QPainterPath::contains(const QPointF & pt);
extern void _ZNK12QPainterPath8containsERK7QPointF(void* qthis, void* arg0);
  // proto:  void QPainterPath::arcTo(const QRectF & rect, qreal startAngle, qreal arcLength);
extern void _ZN12QPainterPath5arcToERK6QRectFdd(void* qthis, void* arg0, double arg1, double arg2);
  // proto:  void QPainterPath::addRoundRect(const QRectF & rect, int roundness);
extern void demth_ZN12QPainterPath12addRoundRectERK6QRectFi(void* qthis, void* arg0, int arg1);
  // proto:  void QPainterPath::addEllipse(const QPointF & center, qreal rx, qreal ry);
extern void demth_ZN12QPainterPath10addEllipseERK7QPointFdd(void* qthis, void* arg0, double arg1, double arg2);
  // proto:  void QPainterPath::lineTo(qreal x, qreal y);
extern void demth_ZN12QPainterPath6lineToEdd(void* qthis, double arg0, double arg1);
  // proto:  void QPainterPath::cubicTo(const QPointF & ctrlPt1, const QPointF & ctrlPt2, const QPointF & endPt);
extern void _ZN12QPainterPath7cubicToERK7QPointFS2_S2_(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  qreal QPainterPath::slopeAtPercent(qreal t);
extern void _ZNK12QPainterPath14slopeAtPercentEd(void* qthis, double arg0);
  // proto:  void QPainterPath::addEllipse(qreal x, qreal y, qreal w, qreal h);
extern void demth_ZN12QPainterPath10addEllipseEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  bool QPainterPath::intersects(const QPainterPath & p);
extern void _ZNK12QPainterPath10intersectsERKS_(void* qthis, void* arg0);
  // proto:  void QPainterPath::addRoundRect(qreal x, qreal y, qreal w, qreal h, int roundness);
extern void demth_ZN12QPainterPath12addRoundRectEddddi(void* qthis, double arg0, double arg1, double arg2, double arg3, int arg4);
  // proto:  void QPainterPath::QPainterPath(const QPointF & startPoint);
extern void* dector_ZN12QPainterPathC1ERK7QPointF(void* arg0);
extern void _ZN12QPainterPathC1ERK7QPointF(void* qthis, void* arg0);
  // proto:  QPainterPath QPainterPath::intersected(const QPainterPath & r);
extern void _ZNK12QPainterPath11intersectedERKS_(void* qthis, void* arg0);
  // proto:  void QPainterPath::translate(qreal dx, qreal dy);
extern void _ZN12QPainterPath9translateEdd(void* qthis, double arg0, double arg1);
  // proto:  void QPainterPath::addPolygon(const QPolygonF & polygon);
extern void _ZN12QPainterPath10addPolygonERK9QPolygonF(void* qthis, void* arg0);
  // proto:  void QPainterPath::translate(const QPointF & offset);
extern void demth_ZN12QPainterPath9translateERK7QPointF(void* qthis, void* arg0);
  // proto:  QPolygonF QPainterPath::toFillPolygon(const QTransform & matrix);
extern void _ZNK12QPainterPath13toFillPolygonERK10QTransform(void* qthis, void* arg0);
  // proto:  void QPainterPath::addPath(const QPainterPath & path);
extern void _ZN12QPainterPath7addPathERKS_(void* qthis, void* arg0);
  // proto:  void QPainterPath::quadTo(qreal ctrlPtx, qreal ctrlPty, qreal endPtx, qreal endPty);
extern void demth_ZN12QPainterPath6quadToEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  int QPainterPath::elementCount();
extern void _ZNK12QPainterPath12elementCountEv(void* qthis);
  // proto:  QPainterPath QPainterPath::simplified();
extern void _ZNK12QPainterPath10simplifiedEv(void* qthis);
  // proto:  bool QPainterPath::contains(const QRectF & rect);
extern void _ZNK12QPainterPath8containsERK6QRectF(void* qthis, void* arg0);
  // proto:  qreal QPainterPath::length();
extern void _ZNK12QPainterPath6lengthEv(void* qthis);
  // proto:  void QPainterPath::connectPath(const QPainterPath & path);
extern void _ZN12QPainterPath11connectPathERKS_(void* qthis, void* arg0);
  // proto:  void QPainterPath::addRegion(const QRegion & region);
extern void _ZN12QPainterPath9addRegionERK7QRegion(void* qthis, void* arg0);
  // proto:  QPointF QPainterPath::currentPosition();
extern void _ZNK12QPainterPath15currentPositionEv(void* qthis);
  // proto:  QPainterPath QPainterPath::toReversed();
extern void _ZNK12QPainterPath10toReversedEv(void* qthis);
  // proto:  void QPainterPath::addRoundRect(qreal x, qreal y, qreal w, qreal h, int xRnd, int yRnd);
extern void demth_ZN12QPainterPath12addRoundRectEddddii(void* qthis, double arg0, double arg1, double arg2, double arg3, int arg4, int arg5);
  // proto:  QRectF QPainterPath::boundingRect();
extern void _ZNK12QPainterPath12boundingRectEv(void* qthis);
  // proto:  void QPainterPath::swap(QPainterPath & other);
extern void demth_ZN12QPainterPath4swapERS_(void* qthis, void* arg0);
  // proto:  bool QPainterPath::contains(const QPainterPath & p);
extern void _ZNK12QPainterPath8containsERKS_(void* qthis, void* arg0);
  // proto:  void QPainterPath::moveTo(qreal x, qreal y);
extern void demth_ZN12QPainterPath6moveToEdd(void* qthis, double arg0, double arg1);
  // proto:  QPainterPath QPainterPath::subtracted(const QPainterPath & r);
extern void _ZNK12QPainterPath10subtractedERKS_(void* qthis, void* arg0);
  // proto:  void QPainterPath::QPainterPath();
extern void* dector_ZN12QPainterPathC1Ev();
extern void _ZN12QPainterPathC1Ev(void* qthis);
  // proto:  void QPainterPath::addText(const QPointF & point, const QFont & f, const QString & text);
extern void _ZN12QPainterPath7addTextERK7QPointFRK5QFontRK7QString(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QPainterPath::QPainterPath(const QPainterPath & other);
extern void* dector_ZN12QPainterPathC1ERKS_(void* arg0);
extern void _ZN12QPainterPathC1ERKS_(void* qthis, void* arg0);
  // proto:  QPointF QPainterPath::pointAtPercent(qreal t);
extern void _ZNK12QPainterPath14pointAtPercentEd(void* qthis, double arg0);
  // proto:  qreal QPainterPath::percentAtLength(qreal t);
extern void _ZNK12QPainterPath15percentAtLengthEd(void* qthis, double arg0);
  // proto:  void QPainterPath::cubicTo(qreal ctrlPt1x, qreal ctrlPt1y, qreal ctrlPt2x, qreal ctrlPt2y, qreal endPtx, qreal endPty);
extern void demth_ZN12QPainterPath7cubicToEdddddd(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4, double arg5);
  // proto:  void QPainterPath::lineTo(const QPointF & p);
extern void _ZN12QPainterPath6lineToERK7QPointF(void* qthis, void* arg0);
  // proto:  QPainterPath QPainterPath::subtractedInverted(const QPainterPath & r);
extern void _ZNK12QPainterPath18subtractedInvertedERKS_(void* qthis, void* arg0);
  // proto:  void QPainterPath::arcMoveTo(qreal x, qreal y, qreal w, qreal h, qreal angle);
extern void demth_ZN12QPainterPath9arcMoveToEddddd(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4);
  // proto:  bool QPainterPath::isEmpty();
extern void _ZNK12QPainterPath7isEmptyEv(void* qthis);
  // proto:  void QPainterPath::addRect(qreal x, qreal y, qreal w, qreal h);
extern void demth_ZN12QPainterPath7addRectEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  void QPainterPath::arcMoveTo(const QRectF & rect, qreal angle);
extern void _ZN12QPainterPath9arcMoveToERK6QRectFd(void* qthis, void* arg0, double arg1);
  // proto:  QList<QPolygonF> QPainterPath::toSubpathPolygons(const QMatrix & matrix);
extern void _ZNK12QPainterPath17toSubpathPolygonsERK7QMatrix(void* qthis, void* arg0);
  // proto:  QPainterPath QPainterPath::united(const QPainterPath & r);
extern void _ZNK12QPainterPath6unitedERKS_(void* qthis, void* arg0);
  // proto:  void QPainterPath::addEllipse(const QRectF & rect);
extern void _ZN12QPainterPath10addEllipseERK6QRectF(void* qthis, void* arg0);
  // proto:  void QPainterPath::moveTo(const QPointF & p);
extern void _ZN12QPainterPath6moveToERK7QPointF(void* qthis, void* arg0);
  // proto:  qreal QPainterPath::angleAtPercent(qreal t);
extern void _ZNK12QPainterPath14angleAtPercentEd(void* qthis, double arg0);
  // proto:  void QPainterPath::closeSubpath();
extern void _ZN12QPainterPath12closeSubpathEv(void* qthis);
  // proto:  qreal QPainterPathStroker::curveThreshold();
extern void _ZNK19QPainterPathStroker14curveThresholdEv(void* qthis);
  // proto:  void QPainterPathStroker::setWidth(qreal width);
extern void _ZN19QPainterPathStroker8setWidthEd(void* qthis, double arg0);
  // proto:  void QPainterPathStroker::~QPainterPathStroker();
extern void _ZN19QPainterPathStrokerD0Ev(void* qthis);
  // proto:  void QPainterPathStroker::setMiterLimit(qreal length);
extern void _ZN19QPainterPathStroker13setMiterLimitEd(void* qthis, double arg0);
  // proto:  void QPainterPathStroker::QPainterPathStroker(const QPen & pen);
extern void* dector_ZN19QPainterPathStrokerC1ERK4QPen(void* arg0);
extern void _ZN19QPainterPathStrokerC1ERK4QPen(void* qthis, void* arg0);
  // proto:  void QPainterPathStroker::setCurveThreshold(qreal threshold);
extern void _ZN19QPainterPathStroker17setCurveThresholdEd(void* qthis, double arg0);
  // proto:  QVector<qreal> QPainterPathStroker::dashPattern();
extern void _ZNK19QPainterPathStroker11dashPatternEv(void* qthis);
  // proto:  qreal QPainterPathStroker::dashOffset();
extern void _ZNK19QPainterPathStroker10dashOffsetEv(void* qthis);
  // proto:  void QPainterPathStroker::QPainterPathStroker();
extern void* dector_ZN19QPainterPathStrokerC1Ev();
extern void _ZN19QPainterPathStrokerC1Ev(void* qthis);
  // proto:  QPainterPath QPainterPathStroker::createStroke(const QPainterPath & path);
extern void _ZNK19QPainterPathStroker12createStrokeERK12QPainterPath(void* qthis, void* arg0);
  // proto:  void QPainterPathStroker::setDashOffset(qreal offset);
extern void _ZN19QPainterPathStroker13setDashOffsetEd(void* qthis, double arg0);
  // proto:  qreal QPainterPathStroker::width();
extern void _ZNK19QPainterPathStroker5widthEv(void* qthis);
  // proto:  void QPainterPathStroker::QPainterPathStroker(const QPainterPathStroker & );
extern void* dector_ZN19QPainterPathStrokerC1ERKS_(void* arg0);
extern void _ZN19QPainterPathStrokerC1ERKS_(void* qthis, void* arg0);
  // proto:  qreal QPainterPathStroker::miterLimit();
extern void _ZNK19QPainterPathStroker10miterLimitEv(void* qthis);
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
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QPainterPathStroker)=1
type QPainterPathStroker struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QPainterPath::setElementPositionAt(int i, qreal x, qreal y);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QPainterPath", "setElementPositionAt", args)
  }

}

  // proto:  QPolygonF QPainterPath::toFillPolygon(const QMatrix & matrix);
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
    var arg0 = args[0].(QMatrix).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK12QPainterPath13toFillPolygonERK10QTransform
    var arg0 = args[0].(QTransform).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPath", "toFillPolygon", args)
  }

}

  // proto:  QPainterPath QPainterPath::translated(qreal dx, qreal dy);
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZNK12QPainterPath10translatedERK7QPointF
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPath", "translated", args)
  }

}

  // proto:  void QPainterPath::~QPainterPath();
func (this *QPainterPath) FreeQPainterPath(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPainterPath", "~QPainterPath", args)
  }

}

  // proto:  QList<QPolygonF> QPainterPath::toSubpathPolygons(const QTransform & matrix);
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
    var arg0 = args[0].(QTransform).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK12QPainterPath17toSubpathPolygonsERK7QMatrix
    var arg0 = args[0].(QMatrix).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPath", "toSubpathPolygons", args)
  }

}

  // proto:  QRectF QPainterPath::controlPointRect();
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
  default:
    qtrt.ErrorResolve("QPainterPath", "controlPointRect", args)
  }

}

  // proto:  QList<QPolygonF> QPainterPath::toFillPolygons(const QMatrix & matrix);
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
    var arg0 = args[0].(QMatrix).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK12QPainterPath14toFillPolygonsERK10QTransform
    var arg0 = args[0].(QTransform).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPath", "toFillPolygons", args)
  }

}

  // proto:  void QPainterPath::quadTo(const QPointF & ctrlPt, const QPointF & endPt);
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
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPointF).qclsinst
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN12QPainterPath6quadToEdddd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QPainterPath", "quadTo", args)
  }

}

  // proto:  void QPainterPath::arcTo(qreal x, qreal y, qreal w, qreal h, qreal startAngle, qreal arcLength);
func (this *QPainterPath) arcTo(args ...interface{}) () {
  // arcTo(qreal, qreal, qreal, qreal, qreal, qreal)
  // arcTo(const class QRectF &, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][5] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath5arcToEdddddd
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
  case 1:
    // invoke: _ZN12QPainterPath5arcToERK6QRectFdd
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QPainterPath", "arcTo", args)
  }

}

  // proto:  void QPainterPath::addRect(const QRectF & rect);
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
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN12QPainterPath7addRectEdddd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QPainterPath", "addRect", args)
  }

}

  // proto:  void QPainterPath::addRoundRect(const QRectF & rect, int xRnd, int yRnd);
func (this *QPainterPath) addRoundRect(args ...interface{}) () {
  // addRoundRect(const class QRectF &, int, int)
  // addRoundRect(const class QRectF &, int)
  // addRoundRect(qreal, qreal, qreal, qreal, int)
  // addRoundRect(qreal, qreal, qreal, qreal, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][4] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][4] = qtrt.Int32Ty(false) // "int"
  vtys[3][5] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath12addRoundRectERK6QRectFii
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 1:
    // invoke: _ZN12QPainterPath12addRoundRectERK6QRectFi
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  case 2:
    // invoke: _ZN12QPainterPath12addRoundRectEddddi
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
  case 3:
    // invoke: _ZN12QPainterPath12addRoundRectEddddii
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
  default:
    qtrt.ErrorResolve("QPainterPath", "addRoundRect", args)
  }

}

  // proto:  void QPainterPath::addText(qreal x, qreal y, const QFont & f, const QString & text);
func (this *QPainterPath) addText(args ...interface{}) () {
  // addText(qreal, qreal, const class QFont &, const class QString &)
  // addText(const class QPointF &, const class QFont &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = reflect.TypeOf(QFont{}) // "const QFont &"
  vtys[0][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1][1] = reflect.TypeOf(QFont{}) // "const QFont &"
  vtys[1][2] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath7addTextEddRK5QFontRK7QString
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QFont).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QString).qclsinst
    if false {fmt.Println(arg3)}
  case 1:
    // invoke: _ZN12QPainterPath7addTextERK7QPointFRK5QFontRK7QString
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QFont).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QPainterPath", "addText", args)
  }

}

  // proto:  bool QPainterPath::intersects(const QRectF & rect);
func (this *QPainterPath) intersects(args ...interface{}) () {
  // intersects(const class QRectF &)
  // intersects(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath10intersectsERK6QRectF
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK12QPainterPath10intersectsERKS_
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPath", "intersects", args)
  }

}

  // proto:  bool QPainterPath::contains(const QPointF & pt);
func (this *QPainterPath) contains(args ...interface{}) () {
  // contains(const class QPointF &)
  // contains(const class QRectF &)
  // contains(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath8containsERK7QPointF
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK12QPainterPath8containsERK6QRectF
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
  case 2:
    // invoke: _ZNK12QPainterPath8containsERKS_
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPath", "contains", args)
  }

}

  // proto:  void QPainterPath::addEllipse(const QPointF & center, qreal rx, qreal ry);
func (this *QPainterPath) addEllipse(args ...interface{}) () {
  // addEllipse(const class QPointF &, qreal, qreal)
  // addEllipse(qreal, qreal, qreal, qreal)
  // addEllipse(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath10addEllipseERK7QPointFdd
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
  case 1:
    // invoke: _ZN12QPainterPath10addEllipseEdddd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
  case 2:
    // invoke: _ZN12QPainterPath10addEllipseERK6QRectF
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPath", "addEllipse", args)
  }

}

  // proto:  void QPainterPath::lineTo(qreal x, qreal y);
func (this *QPainterPath) lineTo(args ...interface{}) () {
  // lineTo(qreal, qreal)
  // lineTo(const class QPointF &)
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
    // invoke: _ZN12QPainterPath6lineToEdd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN12QPainterPath6lineToERK7QPointF
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPath", "lineTo", args)
  }

}

  // proto:  void QPainterPath::cubicTo(const QPointF & ctrlPt1, const QPointF & ctrlPt2, const QPointF & endPt);
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
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPointF).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QPointF).qclsinst
    if false {fmt.Println(arg2)}
  case 1:
    // invoke: _ZN12QPainterPath7cubicToEdddddd
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
  default:
    qtrt.ErrorResolve("QPainterPath", "cubicTo", args)
  }

}

  // proto:  qreal QPainterPath::slopeAtPercent(qreal t);
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPath", "slopeAtPercent", args)
  }

}

  // proto:  void QPainterPath::QPainterPath(const QPointF & startPoint);
func NewQPainterPath(args ...interface{}) QPainterPath {
  return QPainterPath{}
}

  // proto:  QPainterPath QPainterPath::intersected(const QPainterPath & r);
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
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPath", "intersected", args)
  }

}

  // proto:  void QPainterPath::translate(qreal dx, qreal dy);
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN12QPainterPath9translateERK7QPointF
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPath", "translate", args)
  }

}

  // proto:  void QPainterPath::addPolygon(const QPolygonF & polygon);
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
    var arg0 = args[0].(QPolygonF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPath", "addPolygon", args)
  }

}

  // proto:  void QPainterPath::addPath(const QPainterPath & path);
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
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPath", "addPath", args)
  }

}

  // proto:  int QPainterPath::elementCount();
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
  default:
    qtrt.ErrorResolve("QPainterPath", "elementCount", args)
  }

}

  // proto:  QPainterPath QPainterPath::simplified();
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
  default:
    qtrt.ErrorResolve("QPainterPath", "simplified", args)
  }

}

  // proto:  qreal QPainterPath::length();
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
  default:
    qtrt.ErrorResolve("QPainterPath", "length", args)
  }

}

  // proto:  void QPainterPath::connectPath(const QPainterPath & path);
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
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPath", "connectPath", args)
  }

}

  // proto:  void QPainterPath::addRegion(const QRegion & region);
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
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPath", "addRegion", args)
  }

}

  // proto:  QPointF QPainterPath::currentPosition();
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
  default:
    qtrt.ErrorResolve("QPainterPath", "currentPosition", args)
  }

}

  // proto:  QPainterPath QPainterPath::toReversed();
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
  default:
    qtrt.ErrorResolve("QPainterPath", "toReversed", args)
  }

}

  // proto:  QRectF QPainterPath::boundingRect();
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
  default:
    qtrt.ErrorResolve("QPainterPath", "boundingRect", args)
  }

}

  // proto:  void QPainterPath::swap(QPainterPath & other);
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
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPath", "swap", args)
  }

}

  // proto:  void QPainterPath::moveTo(qreal x, qreal y);
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN12QPainterPath6moveToERK7QPointF
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPath", "moveTo", args)
  }

}

  // proto:  QPainterPath QPainterPath::subtracted(const QPainterPath & r);
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
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPath", "subtracted", args)
  }

}

  // proto:  QPointF QPainterPath::pointAtPercent(qreal t);
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPath", "pointAtPercent", args)
  }

}

  // proto:  qreal QPainterPath::percentAtLength(qreal t);
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPath", "percentAtLength", args)
  }

}

  // proto:  QPainterPath QPainterPath::subtractedInverted(const QPainterPath & r);
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
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPath", "subtractedInverted", args)
  }

}

  // proto:  void QPainterPath::arcMoveTo(qreal x, qreal y, qreal w, qreal h, qreal angle);
func (this *QPainterPath) arcMoveTo(args ...interface{}) () {
  // arcMoveTo(qreal, qreal, qreal, qreal, qreal)
  // arcMoveTo(const class QRectF &, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath9arcMoveToEddddd
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
  case 1:
    // invoke: _ZN12QPainterPath9arcMoveToERK6QRectFd
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QPainterPath", "arcMoveTo", args)
  }

}

  // proto:  bool QPainterPath::isEmpty();
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
  default:
    qtrt.ErrorResolve("QPainterPath", "isEmpty", args)
  }

}

  // proto:  QPainterPath QPainterPath::united(const QPainterPath & r);
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
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPath", "united", args)
  }

}

  // proto:  qreal QPainterPath::angleAtPercent(qreal t);
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPath", "angleAtPercent", args)
  }

}

  // proto:  void QPainterPath::closeSubpath();
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
  default:
    qtrt.ErrorResolve("QPainterPath", "closeSubpath", args)
  }

}

  // proto:  qreal QPainterPathStroker::curveThreshold();
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
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "curveThreshold", args)
  }

}

  // proto:  void QPainterPathStroker::setWidth(qreal width);
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "setWidth", args)
  }

}

  // proto:  void QPainterPathStroker::~QPainterPathStroker();
func (this *QPainterPathStroker) FreeQPainterPathStroker(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "~QPainterPathStroker", args)
  }

}

  // proto:  void QPainterPathStroker::setMiterLimit(qreal length);
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "setMiterLimit", args)
  }

}

  // proto:  void QPainterPathStroker::QPainterPathStroker(const QPen & pen);
func NewQPainterPathStroker(args ...interface{}) QPainterPathStroker {
  return QPainterPathStroker{}
}

  // proto:  void QPainterPathStroker::setCurveThreshold(qreal threshold);
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "setCurveThreshold", args)
  }

}

  // proto:  QVector<qreal> QPainterPathStroker::dashPattern();
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
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "dashPattern", args)
  }

}

  // proto:  qreal QPainterPathStroker::dashOffset();
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
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "dashOffset", args)
  }

}

  // proto:  QPainterPath QPainterPathStroker::createStroke(const QPainterPath & path);
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
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "createStroke", args)
  }

}

  // proto:  void QPainterPathStroker::setDashOffset(qreal offset);
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "setDashOffset", args)
  }

}

  // proto:  qreal QPainterPathStroker::width();
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
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "width", args)
  }

}

  // proto:  qreal QPainterPathStroker::miterLimit();
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
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "miterLimit", args)
  }

}

// <= body block end

