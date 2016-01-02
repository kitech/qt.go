package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtGui/qbrush.h
// dst-file: /src/gui/qbrush.go
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
  // proto:  void QRadialGradient::QRadialGradient(qreal cx, qreal cy, qreal radius, qreal fx, qreal fy);
extern void* dector_ZN15QRadialGradientC1Eddddd(double arg0, double arg1, double arg2, double arg3, double arg4);
extern void _ZN15QRadialGradientC1Eddddd(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4);
  // proto:  void QRadialGradient::setFocalPoint(qreal x, qreal y);
extern void demth_ZN15QRadialGradient13setFocalPointEdd(void* qthis, double arg0, double arg1);
  // proto:  void QRadialGradient::QRadialGradient();
extern void* dector_ZN15QRadialGradientC1Ev();
extern void _ZN15QRadialGradientC1Ev(void* qthis);
  // proto:  void QRadialGradient::QRadialGradient(const QPointF & center, qreal radius, const QPointF & focalPoint);
extern void* dector_ZN15QRadialGradientC1ERK7QPointFdS2_(void* arg0, double arg1, void* arg2);
extern void _ZN15QRadialGradientC1ERK7QPointFdS2_(void* qthis, void* arg0, double arg1, void* arg2);
  // proto:  qreal QRadialGradient::radius();
extern void _ZNK15QRadialGradient6radiusEv(void* qthis);
  // proto:  void QRadialGradient::setFocalPoint(const QPointF & focalPoint);
extern void _ZN15QRadialGradient13setFocalPointERK7QPointF(void* qthis, void* arg0);
  // proto:  void QRadialGradient::QRadialGradient(const QPointF & center, qreal centerRadius, const QPointF & focalPoint, qreal focalRadius);
extern void* dector_ZN15QRadialGradientC1ERK7QPointFdS2_d(void* arg0, double arg1, void* arg2, double arg3);
extern void _ZN15QRadialGradientC1ERK7QPointFdS2_d(void* qthis, void* arg0, double arg1, void* arg2, double arg3);
  // proto:  void QRadialGradient::QRadialGradient(qreal cx, qreal cy, qreal centerRadius, qreal fx, qreal fy, qreal focalRadius);
extern void* dector_ZN15QRadialGradientC1Edddddd(double arg0, double arg1, double arg2, double arg3, double arg4, double arg5);
extern void _ZN15QRadialGradientC1Edddddd(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4, double arg5);
  // proto:  qreal QRadialGradient::centerRadius();
extern void _ZNK15QRadialGradient12centerRadiusEv(void* qthis);
  // proto:  QPointF QRadialGradient::focalPoint();
extern void _ZNK15QRadialGradient10focalPointEv(void* qthis);
  // proto:  qreal QRadialGradient::focalRadius();
extern void _ZNK15QRadialGradient11focalRadiusEv(void* qthis);
  // proto:  QPointF QRadialGradient::center();
extern void _ZNK15QRadialGradient6centerEv(void* qthis);
  // proto:  void QRadialGradient::setCenter(const QPointF & center);
extern void _ZN15QRadialGradient9setCenterERK7QPointF(void* qthis, void* arg0);
  // proto:  void QRadialGradient::QRadialGradient(const QPointF & center, qreal radius);
extern void* dector_ZN15QRadialGradientC1ERK7QPointFd(void* arg0, double arg1);
extern void _ZN15QRadialGradientC1ERK7QPointFd(void* qthis, void* arg0, double arg1);
  // proto:  void QRadialGradient::setCenterRadius(qreal radius);
extern void _ZN15QRadialGradient15setCenterRadiusEd(void* qthis, double arg0);
  // proto:  void QRadialGradient::setFocalRadius(qreal radius);
extern void _ZN15QRadialGradient14setFocalRadiusEd(void* qthis, double arg0);
  // proto:  void QRadialGradient::setRadius(qreal radius);
extern void _ZN15QRadialGradient9setRadiusEd(void* qthis, double arg0);
  // proto:  void QRadialGradient::QRadialGradient(qreal cx, qreal cy, qreal radius);
extern void* dector_ZN15QRadialGradientC1Eddd(double arg0, double arg1, double arg2);
extern void _ZN15QRadialGradientC1Eddd(void* qthis, double arg0, double arg1, double arg2);
  // proto:  void QRadialGradient::setCenter(qreal x, qreal y);
extern void demth_ZN15QRadialGradient9setCenterEdd(void* qthis, double arg0, double arg1);
  // proto:  qreal QConicalGradient::angle();
extern void _ZNK16QConicalGradient5angleEv(void* qthis);
  // proto:  QPointF QConicalGradient::center();
extern void _ZNK16QConicalGradient6centerEv(void* qthis);
  // proto:  void QConicalGradient::QConicalGradient(const QPointF & center, qreal startAngle);
extern void* dector_ZN16QConicalGradientC1ERK7QPointFd(void* arg0, double arg1);
extern void _ZN16QConicalGradientC1ERK7QPointFd(void* qthis, void* arg0, double arg1);
  // proto:  void QConicalGradient::QConicalGradient();
extern void* dector_ZN16QConicalGradientC1Ev();
extern void _ZN16QConicalGradientC1Ev(void* qthis);
  // proto:  void QConicalGradient::setAngle(qreal angle);
extern void _ZN16QConicalGradient8setAngleEd(void* qthis, double arg0);
  // proto:  void QConicalGradient::setCenter(qreal x, qreal y);
extern void demth_ZN16QConicalGradient9setCenterEdd(void* qthis, double arg0, double arg1);
  // proto:  void QConicalGradient::setCenter(const QPointF & center);
extern void _ZN16QConicalGradient9setCenterERK7QPointF(void* qthis, void* arg0);
  // proto:  void QConicalGradient::QConicalGradient(qreal cx, qreal cy, qreal startAngle);
extern void* dector_ZN16QConicalGradientC1Eddd(double arg0, double arg1, double arg2);
extern void _ZN16QConicalGradientC1Eddd(void* qthis, double arg0, double arg1, double arg2);
  // proto:  void QBrush::QBrush();
extern void* dector_ZN6QBrushC1Ev();
extern void _ZN6QBrushC1Ev(void* qthis);
  // proto:  void QBrush::QBrush(const QPixmap & pixmap);
extern void* dector_ZN6QBrushC1ERK7QPixmap(void* arg0);
extern void _ZN6QBrushC1ERK7QPixmap(void* qthis, void* arg0);
  // proto:  void QBrush::setTexture(const QPixmap & pixmap);
extern void _ZN6QBrush10setTextureERK7QPixmap(void* qthis, void* arg0);
  // proto:  void QBrush::setTextureImage(const QImage & image);
extern void _ZN6QBrush15setTextureImageERK6QImage(void* qthis, void* arg0);
  // proto:  void QBrush::QBrush(const QColor & color, const QPixmap & pixmap);
extern void* dector_ZN6QBrushC1ERK6QColorRK7QPixmap(void* arg0, void* arg1);
extern void _ZN6QBrushC1ERK6QColorRK7QPixmap(void* qthis, void* arg0, void* arg1);
  // proto:  QPixmap QBrush::texture();
extern void _ZNK6QBrush7textureEv(void* qthis);
  // proto:  void QBrush::QBrush(const QGradient & gradient);
extern void* dector_ZN6QBrushC1ERK9QGradient(void* arg0);
extern void _ZN6QBrushC1ERK9QGradient(void* qthis, void* arg0);
  // proto:  QTransform QBrush::transform();
extern void demth_ZNK6QBrush9transformEv(void* qthis);
  // proto:  void QBrush::setTransform(const QTransform & );
extern void _ZN6QBrush12setTransformERK10QTransform(void* qthis, void* arg0);
  // proto:  bool QBrush::isOpaque();
extern void _ZNK6QBrush8isOpaqueEv(void* qthis);
  // proto:  const QGradient * QBrush::gradient();
extern void _ZNK6QBrush8gradientEv(void* qthis);
  // proto:  void QBrush::~QBrush();
extern void _ZN6QBrushD0Ev(void* qthis);
  // proto:  void QBrush::setMatrix(const QMatrix & mat);
extern void _ZN6QBrush9setMatrixERK7QMatrix(void* qthis, void* arg0);
  // proto:  void QBrush::setColor(const QColor & color);
extern void _ZN6QBrush8setColorERK6QColor(void* qthis, void* arg0);
  // proto:  void QBrush::QBrush(const QBrush & brush);
extern void* dector_ZN6QBrushC1ERKS_(void* arg0);
extern void _ZN6QBrushC1ERKS_(void* qthis, void* arg0);
  // proto:  const QMatrix & QBrush::matrix();
extern void demth_ZNK6QBrush6matrixEv(void* qthis);
  // proto:  QImage QBrush::textureImage();
extern void _ZNK6QBrush12textureImageEv(void* qthis);
  // proto:  bool QBrush::isDetached();
extern void demth_ZNK6QBrush10isDetachedEv(void* qthis);
  // proto:  void QBrush::swap(QBrush & other);
extern void demth_ZN6QBrush4swapERS_(void* qthis, void* arg0);
  // proto:  const QColor & QBrush::color();
extern void demth_ZNK6QBrush5colorEv(void* qthis);
  // proto:  void QBrush::QBrush(const QImage & image);
extern void* dector_ZN6QBrushC1ERK6QImage(void* arg0);
extern void _ZN6QBrushC1ERK6QImage(void* qthis, void* arg0);
  // proto:  void QGradient::setColorAt(qreal pos, const QColor & color);
extern void _ZN9QGradient10setColorAtEdRK6QColor(void* qthis, double arg0, void* arg1);
  // proto:  QGradientStops QGradient::stops();
extern void _ZNK9QGradient5stopsEv(void* qthis);
  // proto:  void QGradient::QGradient();
extern void* dector_ZN9QGradientC1Ev();
extern void _ZN9QGradientC1Ev(void* qthis);
  // proto:  void QLinearGradient::setFinalStop(qreal x, qreal y);
extern void demth_ZN15QLinearGradient12setFinalStopEdd(void* qthis, double arg0, double arg1);
  // proto:  QPointF QLinearGradient::start();
extern void _ZNK15QLinearGradient5startEv(void* qthis);
  // proto:  void QLinearGradient::QLinearGradient(qreal xStart, qreal yStart, qreal xFinalStop, qreal yFinalStop);
extern void* dector_ZN15QLinearGradientC1Edddd(double arg0, double arg1, double arg2, double arg3);
extern void _ZN15QLinearGradientC1Edddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  void QLinearGradient::QLinearGradient(const QPointF & start, const QPointF & finalStop);
extern void* dector_ZN15QLinearGradientC1ERK7QPointFS2_(void* arg0, void* arg1);
extern void _ZN15QLinearGradientC1ERK7QPointFS2_(void* qthis, void* arg0, void* arg1);
  // proto:  void QLinearGradient::setStart(qreal x, qreal y);
extern void demth_ZN15QLinearGradient8setStartEdd(void* qthis, double arg0, double arg1);
  // proto:  void QLinearGradient::setStart(const QPointF & start);
extern void _ZN15QLinearGradient8setStartERK7QPointF(void* qthis, void* arg0);
  // proto:  void QLinearGradient::QLinearGradient();
extern void* dector_ZN15QLinearGradientC1Ev();
extern void _ZN15QLinearGradientC1Ev(void* qthis);
  // proto:  QPointF QLinearGradient::finalStop();
extern void _ZNK15QLinearGradient9finalStopEv(void* qthis);
  // proto:  void QLinearGradient::setFinalStop(const QPointF & stop);
extern void _ZN15QLinearGradient12setFinalStopERK7QPointF(void* qthis, void* arg0);
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

// class sizeof(QRadialGradient)=1
type QRadialGradient struct {
  /*qbase*/ QGradient;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QConicalGradient)=1
type QConicalGradient struct {
  /*qbase*/ QGradient;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QBrush)=1
type QBrush struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QGradient)=1
type QGradient struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QBrushData)=1
type QBrushData struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QLinearGradient)=1
type QLinearGradient struct {
  /*qbase*/ QGradient;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QRadialGradient::QRadialGradient(qreal cx, qreal cy, qreal radius, qreal fx, qreal fy);
func NewQRadialGradient(args ...interface{}) QRadialGradient {
  return QRadialGradient{}
}

  // proto:  void QRadialGradient::setFocalPoint(qreal x, qreal y);
func (this *QRadialGradient) setFocalPoint(args ...interface{}) () {
  // setFocalPoint(qreal, qreal)
  // setFocalPoint(const class QPointF &)
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
    // invoke: _ZN15QRadialGradient13setFocalPointEdd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN15QRadialGradient13setFocalPointERK7QPointF
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRadialGradient", "setFocalPoint", args)
  }

}

  // proto:  qreal QRadialGradient::radius();
func (this *QRadialGradient) radius(args ...interface{}) () {
  // radius()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QRadialGradient6radiusEv
  default:
    qtrt.ErrorResolve("QRadialGradient", "radius", args)
  }

}

  // proto:  qreal QRadialGradient::centerRadius();
func (this *QRadialGradient) centerRadius(args ...interface{}) () {
  // centerRadius()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QRadialGradient12centerRadiusEv
  default:
    qtrt.ErrorResolve("QRadialGradient", "centerRadius", args)
  }

}

  // proto:  QPointF QRadialGradient::focalPoint();
func (this *QRadialGradient) focalPoint(args ...interface{}) () {
  // focalPoint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QRadialGradient10focalPointEv
  default:
    qtrt.ErrorResolve("QRadialGradient", "focalPoint", args)
  }

}

  // proto:  qreal QRadialGradient::focalRadius();
func (this *QRadialGradient) focalRadius(args ...interface{}) () {
  // focalRadius()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QRadialGradient11focalRadiusEv
  default:
    qtrt.ErrorResolve("QRadialGradient", "focalRadius", args)
  }

}

  // proto:  QPointF QRadialGradient::center();
func (this *QRadialGradient) center(args ...interface{}) () {
  // center()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QRadialGradient6centerEv
  default:
    qtrt.ErrorResolve("QRadialGradient", "center", args)
  }

}

  // proto:  void QRadialGradient::setCenter(const QPointF & center);
func (this *QRadialGradient) setCenter(args ...interface{}) () {
  // setCenter(const class QPointF &)
  // setCenter(qreal, qreal)
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
    // invoke: _ZN15QRadialGradient9setCenterERK7QPointF
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN15QRadialGradient9setCenterEdd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QRadialGradient", "setCenter", args)
  }

}

  // proto:  void QRadialGradient::setCenterRadius(qreal radius);
func (this *QRadialGradient) setCenterRadius(args ...interface{}) () {
  // setCenterRadius(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QRadialGradient15setCenterRadiusEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRadialGradient", "setCenterRadius", args)
  }

}

  // proto:  void QRadialGradient::setFocalRadius(qreal radius);
func (this *QRadialGradient) setFocalRadius(args ...interface{}) () {
  // setFocalRadius(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QRadialGradient14setFocalRadiusEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRadialGradient", "setFocalRadius", args)
  }

}

  // proto:  void QRadialGradient::setRadius(qreal radius);
func (this *QRadialGradient) setRadius(args ...interface{}) () {
  // setRadius(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QRadialGradient9setRadiusEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRadialGradient", "setRadius", args)
  }

}

  // proto:  qreal QConicalGradient::angle();
func (this *QConicalGradient) angle(args ...interface{}) () {
  // angle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QConicalGradient5angleEv
  default:
    qtrt.ErrorResolve("QConicalGradient", "angle", args)
  }

}

  // proto:  QPointF QConicalGradient::center();
func (this *QConicalGradient) center(args ...interface{}) () {
  // center()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QConicalGradient6centerEv
  default:
    qtrt.ErrorResolve("QConicalGradient", "center", args)
  }

}

  // proto:  void QConicalGradient::QConicalGradient(const QPointF & center, qreal startAngle);
func NewQConicalGradient(args ...interface{}) QConicalGradient {
  return QConicalGradient{}
}

  // proto:  void QConicalGradient::setAngle(qreal angle);
func (this *QConicalGradient) setAngle(args ...interface{}) () {
  // setAngle(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QConicalGradient8setAngleEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QConicalGradient", "setAngle", args)
  }

}

  // proto:  void QConicalGradient::setCenter(qreal x, qreal y);
func (this *QConicalGradient) setCenter(args ...interface{}) () {
  // setCenter(qreal, qreal)
  // setCenter(const class QPointF &)
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
    // invoke: _ZN16QConicalGradient9setCenterEdd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN16QConicalGradient9setCenterERK7QPointF
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QConicalGradient", "setCenter", args)
  }

}

  // proto:  void QBrush::QBrush();
func NewQBrush(args ...interface{}) QBrush {
  return QBrush{}
}

  // proto:  void QBrush::setTexture(const QPixmap & pixmap);
func (this *QBrush) setTexture(args ...interface{}) () {
  // setTexture(const class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QBrush10setTextureERK7QPixmap
    var arg0 = args[0].(QPixmap).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QBrush", "setTexture", args)
  }

}

  // proto:  void QBrush::setTextureImage(const QImage & image);
func (this *QBrush) setTextureImage(args ...interface{}) () {
  // setTextureImage(const class QImage &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QImage{}) // "const QImage &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QBrush15setTextureImageERK6QImage
    var arg0 = args[0].(QImage).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QBrush", "setTextureImage", args)
  }

}

  // proto:  QPixmap QBrush::texture();
func (this *QBrush) texture(args ...interface{}) () {
  // texture()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QBrush7textureEv
  default:
    qtrt.ErrorResolve("QBrush", "texture", args)
  }

}

  // proto:  QTransform QBrush::transform();
func (this *QBrush) transform(args ...interface{}) () {
  // transform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QBrush9transformEv
  default:
    qtrt.ErrorResolve("QBrush", "transform", args)
  }

}

  // proto:  void QBrush::setTransform(const QTransform & );
func (this *QBrush) setTransform(args ...interface{}) () {
  // setTransform(const class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QBrush12setTransformERK10QTransform
    var arg0 = args[0].(QTransform).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QBrush", "setTransform", args)
  }

}

  // proto:  bool QBrush::isOpaque();
func (this *QBrush) isOpaque(args ...interface{}) () {
  // isOpaque()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QBrush8isOpaqueEv
  default:
    qtrt.ErrorResolve("QBrush", "isOpaque", args)
  }

}

  // proto:  const QGradient * QBrush::gradient();
func (this *QBrush) gradient(args ...interface{}) () {
  // gradient()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QBrush8gradientEv
  default:
    qtrt.ErrorResolve("QBrush", "gradient", args)
  }

}

  // proto:  void QBrush::~QBrush();
func (this *QBrush) FreeQBrush(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QBrush", "~QBrush", args)
  }

}

  // proto:  void QBrush::setMatrix(const QMatrix & mat);
func (this *QBrush) setMatrix(args ...interface{}) () {
  // setMatrix(const class QMatrix &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix{}) // "const QMatrix &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QBrush9setMatrixERK7QMatrix
    var arg0 = args[0].(QMatrix).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QBrush", "setMatrix", args)
  }

}

  // proto:  void QBrush::setColor(const QColor & color);
func (this *QBrush) setColor(args ...interface{}) () {
  // setColor(const class QColor &)
  // setColor(Qt::GlobalColor)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "Qt::GlobalColor"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QBrush8setColorERK6QColor
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN6QBrush8setColorEN2Qt11GlobalColorE
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QBrush", "setColor", args)
  }

}

  // proto:  const QMatrix & QBrush::matrix();
func (this *QBrush) matrix(args ...interface{}) () {
  // matrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QBrush6matrixEv
  default:
    qtrt.ErrorResolve("QBrush", "matrix", args)
  }

}

  // proto:  QImage QBrush::textureImage();
func (this *QBrush) textureImage(args ...interface{}) () {
  // textureImage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QBrush12textureImageEv
  default:
    qtrt.ErrorResolve("QBrush", "textureImage", args)
  }

}

  // proto:  bool QBrush::isDetached();
func (this *QBrush) isDetached(args ...interface{}) () {
  // isDetached()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QBrush10isDetachedEv
  default:
    qtrt.ErrorResolve("QBrush", "isDetached", args)
  }

}

  // proto:  void QBrush::swap(QBrush & other);
func (this *QBrush) swap(args ...interface{}) () {
  // swap(class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QBrush4swapERS_
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QBrush", "swap", args)
  }

}

  // proto:  const QColor & QBrush::color();
func (this *QBrush) color(args ...interface{}) () {
  // color()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QBrush5colorEv
  default:
    qtrt.ErrorResolve("QBrush", "color", args)
  }

}

  // proto:  void QGradient::setColorAt(qreal pos, const QColor & color);
func (this *QGradient) setColorAt(args ...interface{}) () {
  // setColorAt(qreal, const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGradient10setColorAtEdRK6QColor
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QColor).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QGradient", "setColorAt", args)
  }

}

  // proto:  QGradientStops QGradient::stops();
func (this *QGradient) stops(args ...interface{}) () {
  // stops()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGradient5stopsEv
  default:
    qtrt.ErrorResolve("QGradient", "stops", args)
  }

}

  // proto:  void QGradient::QGradient();
func NewQGradient(args ...interface{}) QGradient {
  return QGradient{}
}

  // proto:  void QLinearGradient::setFinalStop(qreal x, qreal y);
func (this *QLinearGradient) setFinalStop(args ...interface{}) () {
  // setFinalStop(qreal, qreal)
  // setFinalStop(const class QPointF &)
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
    // invoke: _ZN15QLinearGradient12setFinalStopEdd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN15QLinearGradient12setFinalStopERK7QPointF
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QLinearGradient", "setFinalStop", args)
  }

}

  // proto:  QPointF QLinearGradient::start();
func (this *QLinearGradient) start(args ...interface{}) () {
  // start()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QLinearGradient5startEv
  default:
    qtrt.ErrorResolve("QLinearGradient", "start", args)
  }

}

  // proto:  void QLinearGradient::QLinearGradient(qreal xStart, qreal yStart, qreal xFinalStop, qreal yFinalStop);
func NewQLinearGradient(args ...interface{}) QLinearGradient {
  return QLinearGradient{}
}

  // proto:  void QLinearGradient::setStart(qreal x, qreal y);
func (this *QLinearGradient) setStart(args ...interface{}) () {
  // setStart(qreal, qreal)
  // setStart(const class QPointF &)
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
    // invoke: _ZN15QLinearGradient8setStartEdd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN15QLinearGradient8setStartERK7QPointF
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QLinearGradient", "setStart", args)
  }

}

  // proto:  QPointF QLinearGradient::finalStop();
func (this *QLinearGradient) finalStop(args ...interface{}) () {
  // finalStop()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QLinearGradient9finalStopEv
  default:
    qtrt.ErrorResolve("QLinearGradient", "finalStop", args)
  }

}

// <= body block end

