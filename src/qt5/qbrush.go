package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:13 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  qreal QRadialGradient::centerRadius();
extern void _ZNK15QRadialGradient12centerRadiusEv(void* qthis); // 4
  // proto:  void QRadialGradient::setFocalPoint(const QPointF & focalPoint);
extern void _ZN15QRadialGradient13setFocalPointERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QRadialGradient::setFocalPoint(qreal x, qreal y);
extern void _ZN15QRadialGradient13setFocalPointEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QRadialGradient::setRadius(qreal radius);
extern void _ZN15QRadialGradient9setRadiusEd(void* qthis, double arg0); // 4
  // proto:  QPointF QRadialGradient::focalPoint();
extern void _ZNK15QRadialGradient10focalPointEv(void* qthis); // 4
  // proto:  void QRadialGradient::QRadialGradient(qreal cx, qreal cy, qreal radius, qreal fx, qreal fy);
extern void _ZN15QRadialGradientC2Eddddd(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4); // 3
  // proto:  void QRadialGradient::QRadialGradient();
extern void _ZN15QRadialGradientC2Ev(void* qthis); // 3
  // proto:  void QRadialGradient::QRadialGradient(const QPointF & center, qreal radius, const QPointF & focalPoint);
extern void _ZN15QRadialGradientC2ERK7QPointFdS2_(void* qthis, void* arg0, double arg1, void* arg2); // 3
  // proto:  void QRadialGradient::QRadialGradient(const QPointF & center, qreal centerRadius, const QPointF & focalPoint, qreal focalRadius);
extern void _ZN15QRadialGradientC2ERK7QPointFdS2_d(void* qthis, void* arg0, double arg1, void* arg2, double arg3); // 3
  // proto:  void QRadialGradient::QRadialGradient(qreal cx, qreal cy, qreal centerRadius, qreal fx, qreal fy, qreal focalRadius);
extern void _ZN15QRadialGradientC2Edddddd(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4, double arg5); // 3
  // proto:  void QRadialGradient::QRadialGradient(const QPointF & center, qreal radius);
extern void _ZN15QRadialGradientC2ERK7QPointFd(void* qthis, void* arg0, double arg1); // 3
  // proto:  void QRadialGradient::QRadialGradient(qreal cx, qreal cy, qreal radius);
extern void _ZN15QRadialGradientC2Eddd(void* qthis, double arg0, double arg1, double arg2); // 3
  // proto:  void QRadialGradient::setCenter(const QPointF & center);
extern void _ZN15QRadialGradient9setCenterERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QRadialGradient::setCenter(qreal x, qreal y);
extern void _ZN15QRadialGradient9setCenterEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QRadialGradient::setCenterRadius(qreal radius);
extern void _ZN15QRadialGradient15setCenterRadiusEd(void* qthis, double arg0); // 4
  // proto:  qreal QRadialGradient::radius();
extern void _ZNK15QRadialGradient6radiusEv(void* qthis); // 4
  // proto:  void QRadialGradient::setFocalRadius(qreal radius);
extern void _ZN15QRadialGradient14setFocalRadiusEd(void* qthis, double arg0); // 4
  // proto:  qreal QRadialGradient::focalRadius();
extern void _ZNK15QRadialGradient11focalRadiusEv(void* qthis); // 4
  // proto:  QPointF QRadialGradient::center();
extern void _ZNK15QRadialGradient6centerEv(void* qthis); // 4
  // proto:  void QConicalGradient::QConicalGradient(qreal cx, qreal cy, qreal startAngle);
extern void _ZN16QConicalGradientC2Eddd(void* qthis, double arg0, double arg1, double arg2); // 3
  // proto:  void QConicalGradient::QConicalGradient(const QPointF & center, qreal startAngle);
extern void _ZN16QConicalGradientC2ERK7QPointFd(void* qthis, void* arg0, double arg1); // 3
  // proto:  void QConicalGradient::QConicalGradient();
extern void _ZN16QConicalGradientC2Ev(void* qthis); // 3
  // proto:  void QConicalGradient::setCenter(const QPointF & center);
extern void _ZN16QConicalGradient9setCenterERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QConicalGradient::setCenter(qreal x, qreal y);
extern void _ZN16QConicalGradient9setCenterEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  qreal QConicalGradient::angle();
extern void _ZNK16QConicalGradient5angleEv(void* qthis); // 4
  // proto:  QPointF QConicalGradient::center();
extern void _ZNK16QConicalGradient6centerEv(void* qthis); // 4
  // proto:  void QConicalGradient::setAngle(qreal angle);
extern void _ZN16QConicalGradient8setAngleEd(void* qthis, double arg0); // 4
  // proto:  void QBrush::setTexture(const QPixmap & pixmap);
extern void _ZN6QBrush10setTextureERK7QPixmap(void* qthis, void* arg0); // 4
  // proto:  void QBrush::setTransform(const QTransform & );
extern void _ZN6QBrush12setTransformERK10QTransform(void* qthis, void* arg0); // 4
  // proto:  void QBrush::setMatrix(const QMatrix & mat);
extern void _ZN6QBrush9setMatrixERK7QMatrix(void* qthis, void* arg0); // 4
  // proto:  bool QBrush::isOpaque();
extern void _ZNK6QBrush8isOpaqueEv(void* qthis); // 4
  // proto:  Qt::BrushStyle QBrush::style();
extern void _ZNK6QBrush5styleEv(void* qthis); // 2
  // proto:  const QMatrix & QBrush::matrix();
extern void _ZNK6QBrush6matrixEv(void* qthis); // 2
  // proto:  const QGradient * QBrush::gradient();
extern void _ZNK6QBrush8gradientEv(void* qthis); // 4
  // proto:  QImage QBrush::textureImage();
extern void _ZNK6QBrush12textureImageEv(void* qthis); // 4
  // proto:  QTransform QBrush::transform();
extern void _ZNK6QBrush9transformEv(void* qthis); // 2
  // proto:  QPixmap QBrush::texture();
extern void _ZNK6QBrush7textureEv(void* qthis); // 4
  // proto:  const QColor & QBrush::color();
extern void _ZNK6QBrush5colorEv(void* qthis); // 2
  // proto:  void QBrush::swap(QBrush & other);
extern void _ZN6QBrush4swapERS_(void* qthis, void* arg0); // 2
  // proto:  void QBrush::QBrush();
extern void _ZN6QBrushC2Ev(void* qthis); // 3
  // proto:  void QBrush::QBrush(const QPixmap & pixmap);
extern void _ZN6QBrushC2ERK7QPixmap(void* qthis, void* arg0); // 3
  // proto:  void QBrush::QBrush(const QBrush & brush);
extern void _ZN6QBrushC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  void QBrush::QBrush(const QImage & image);
extern void _ZN6QBrushC2ERK6QImage(void* qthis, void* arg0); // 3
  // proto:  void QBrush::QBrush(const QColor & color, const QPixmap & pixmap);
extern void _ZN6QBrushC2ERK6QColorRK7QPixmap(void* qthis, void* arg0, void* arg1); // 3
  // proto:  void QBrush::QBrush(const QGradient & gradient);
extern void _ZN6QBrushC2ERK9QGradient(void* qthis, void* arg0); // 3
  // proto:  bool QBrush::isDetached();
extern void _ZNK6QBrush10isDetachedEv(void* qthis); // 2
  // proto:  void QBrush::setTextureImage(const QImage & image);
extern void _ZN6QBrush15setTextureImageERK6QImage(void* qthis, void* arg0); // 4
  // proto:  void QBrush::~QBrush();
extern void _ZN6QBrushD2Ev(void* qthis); // 4
  // proto:  void QBrush::setColor(const QColor & color);
extern void _ZN6QBrush8setColorERK6QColor(void* qthis, void* arg0); // 4
  // proto:  QGradient::InterpolationMode QGradient::interpolationMode();
extern void _ZNK9QGradient17interpolationModeEv(void* qthis); // 4
  // proto:  void QGradient::QGradient();
extern void _ZN9QGradientC2Ev(void* qthis); // 3
  // proto:  void QGradient::setColorAt(qreal pos, const QColor & color);
extern void _ZN9QGradient10setColorAtEdRK6QColor(void* qthis, double arg0, void* arg1); // 4
  // proto:  QGradientStops QGradient::stops();
extern void _ZNK9QGradient5stopsEv(void* qthis); // 4
  // proto:  QGradient::CoordinateMode QGradient::coordinateMode();
extern void _ZNK9QGradient14coordinateModeEv(void* qthis); // 4
  // proto:  QGradient::Spread QGradient::spread();
extern void _ZNK9QGradient6spreadEv(void* qthis); // 2
  // proto:  QGradient::Type QGradient::type();
extern void _ZNK9QGradient4typeEv(void* qthis); // 2
  // proto:  QPointF QLinearGradient::start();
extern void _ZNK15QLinearGradient5startEv(void* qthis); // 4
  // proto:  QPointF QLinearGradient::finalStop();
extern void _ZNK15QLinearGradient9finalStopEv(void* qthis); // 4
  // proto:  void QLinearGradient::setFinalStop(qreal x, qreal y);
extern void _ZN15QLinearGradient12setFinalStopEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QLinearGradient::setFinalStop(const QPointF & stop);
extern void _ZN15QLinearGradient12setFinalStopERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QLinearGradient::setStart(const QPointF & start);
extern void _ZN15QLinearGradient8setStartERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QLinearGradient::setStart(qreal x, qreal y);
extern void _ZN15QLinearGradient8setStartEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QLinearGradient::QLinearGradient();
extern void _ZN15QLinearGradientC2Ev(void* qthis); // 3
  // proto:  void QLinearGradient::QLinearGradient(qreal xStart, qreal yStart, qreal xFinalStop, qreal yFinalStop);
extern void _ZN15QLinearGradientC2Edddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 3
  // proto:  void QLinearGradient::QLinearGradient(const QPointF & start, const QPointF & finalStop);
extern void _ZN15QLinearGradientC2ERK7QPointFS2_(void* qthis, void* arg0, void* arg1); // 3
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QConicalGradient)=1
type QConicalGradient struct {
  /*qbase*/ QGradient;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QBrush)=1
type QBrush struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGradient)=1
type QGradient struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QBrushData)=1
type QBrushData struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QLinearGradient)=1
type QLinearGradient struct {
  /*qbase*/ QGradient;
  qclsinst unsafe.Pointer /* *C.void */;
}

// centerRadius()
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
    // invoke: qreal centerRadius()
    C._ZNK15QRadialGradient12centerRadiusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRadialGradient", "centerRadius", args)
  }

}

// setFocalPoint(const class QPointF &)
func (this *QRadialGradient) setFocalPoint(args ...interface{}) () {
  // setFocalPoint(const class QPointF &)
  // setFocalPoint(qreal, qreal)
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
    // invoke: _ZN15QRadialGradient13setFocalPointERK7QPointF
    // invoke: void setFocalPoint(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QRadialGradient13setFocalPointERK7QPointF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN15QRadialGradient13setFocalPointEdd
    // invoke: void setFocalPoint(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C._ZN15QRadialGradient13setFocalPointEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QRadialGradient", "setFocalPoint", args)
  }

}

// setRadius(qreal)
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
    // invoke: void setRadius(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN15QRadialGradient9setRadiusEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRadialGradient", "setRadius", args)
  }

}

// focalPoint()
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
    // invoke: QPointF focalPoint()
    C._ZNK15QRadialGradient10focalPointEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRadialGradient", "focalPoint", args)
  }

}

// QRadialGradient(qreal, qreal, qreal, qreal, qreal)
func NewQRadialGradient(args ...interface{}) QRadialGradient {
  // QRadialGradient(qreal, qreal, qreal, qreal, qreal)
  // QRadialGradient()
  // QRadialGradient(const class QPointF &, qreal, const class QPointF &)
  // QRadialGradient(const class QPointF &, qreal, const class QPointF &, qreal)
  // QRadialGradient(qreal, qreal, qreal, qreal, qreal, qreal)
  // QRadialGradient(const class QPointF &, qreal)
  // QRadialGradient(qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][2] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[3][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][2] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[3][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[4][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[4][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[4][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[4][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[4][5] = qtrt.DoubleTy(false) // "qreal"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[5][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[6][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[6][2] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QRadialGradientC1Eddddd
    // invoke: void QRadialGradient(qreal, qreal, qreal, qreal, qreal)
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
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN15QRadialGradientC2Eddddd(qthis, arg0, arg1, arg2, arg3, arg4)
  case 1:
    // invoke: _ZN15QRadialGradientC1Ev
    // invoke: void QRadialGradient()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN15QRadialGradientC2Ev(qthis)
  case 2:
    // invoke: _ZN15QRadialGradientC1ERK7QPointFdS2_
    // invoke: void QRadialGradient(const class QPointF &, qreal, const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QPointF).qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN15QRadialGradientC2ERK7QPointFdS2_(qthis, arg0, arg1, arg2)
  case 3:
    // invoke: _ZN15QRadialGradientC1ERK7QPointFdS2_d
    // invoke: void QRadialGradient(const class QPointF &, qreal, const class QPointF &, qreal)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QPointF).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN15QRadialGradientC2ERK7QPointFdS2_d(qthis, arg0, arg1, arg2, arg3)
  case 4:
    // invoke: _ZN15QRadialGradientC1Edddddd
    // invoke: void QRadialGradient(qreal, qreal, qreal, qreal, qreal, qreal)
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
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN15QRadialGradientC2Edddddd(qthis, arg0, arg1, arg2, arg3, arg4, arg5)
  case 5:
    // invoke: _ZN15QRadialGradientC1ERK7QPointFd
    // invoke: void QRadialGradient(const class QPointF &, qreal)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN15QRadialGradientC2ERK7QPointFd(qthis, arg0, arg1)
  case 6:
    // invoke: _ZN15QRadialGradientC1Eddd
    // invoke: void QRadialGradient(qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN15QRadialGradientC2Eddd(qthis, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QRadialGradient", "QRadialGradient", args)
  }

  return QRadialGradient{}
}

// setCenter(const class QPointF &)
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
    // invoke: void setCenter(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QRadialGradient9setCenterERK7QPointF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN15QRadialGradient9setCenterEdd
    // invoke: void setCenter(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C._ZN15QRadialGradient9setCenterEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QRadialGradient", "setCenter", args)
  }

}

// setCenterRadius(qreal)
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
    // invoke: void setCenterRadius(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN15QRadialGradient15setCenterRadiusEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRadialGradient", "setCenterRadius", args)
  }

}

// radius()
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
    // invoke: qreal radius()
    C._ZNK15QRadialGradient6radiusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRadialGradient", "radius", args)
  }

}

// setFocalRadius(qreal)
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
    // invoke: void setFocalRadius(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN15QRadialGradient14setFocalRadiusEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRadialGradient", "setFocalRadius", args)
  }

}

// focalRadius()
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
    // invoke: qreal focalRadius()
    C._ZNK15QRadialGradient11focalRadiusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRadialGradient", "focalRadius", args)
  }

}

// center()
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
    // invoke: QPointF center()
    C._ZNK15QRadialGradient6centerEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRadialGradient", "center", args)
  }

}

// QConicalGradient(qreal, qreal, qreal)
func NewQConicalGradient(args ...interface{}) QConicalGradient {
  // QConicalGradient(qreal, qreal, qreal)
  // QConicalGradient(const class QPointF &, qreal)
  // QConicalGradient()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QConicalGradientC1Eddd
    // invoke: void QConicalGradient(qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN16QConicalGradientC2Eddd(qthis, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN16QConicalGradientC1ERK7QPointFd
    // invoke: void QConicalGradient(const class QPointF &, qreal)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN16QConicalGradientC2ERK7QPointFd(qthis, arg0, arg1)
  case 2:
    // invoke: _ZN16QConicalGradientC1Ev
    // invoke: void QConicalGradient()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN16QConicalGradientC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QConicalGradient", "QConicalGradient", args)
  }

  return QConicalGradient{}
}

// setCenter(const class QPointF &)
func (this *QConicalGradient) setCenter(args ...interface{}) () {
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
    // invoke: _ZN16QConicalGradient9setCenterERK7QPointF
    // invoke: void setCenter(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN16QConicalGradient9setCenterERK7QPointF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN16QConicalGradient9setCenterEdd
    // invoke: void setCenter(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C._ZN16QConicalGradient9setCenterEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QConicalGradient", "setCenter", args)
  }

}

// angle()
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
    // invoke: qreal angle()
    C._ZNK16QConicalGradient5angleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QConicalGradient", "angle", args)
  }

}

// center()
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
    // invoke: QPointF center()
    C._ZNK16QConicalGradient6centerEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QConicalGradient", "center", args)
  }

}

// setAngle(qreal)
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
    // invoke: void setAngle(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN16QConicalGradient8setAngleEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QConicalGradient", "setAngle", args)
  }

}

// setTexture(const class QPixmap &)
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
    // invoke: void setTexture(const class QPixmap &)
    var arg0 = args[0].(QPixmap).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN6QBrush10setTextureERK7QPixmap(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBrush", "setTexture", args)
  }

}

// setTransform(const class QTransform &)
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
    // invoke: void setTransform(const class QTransform &)
    var arg0 = args[0].(QTransform).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN6QBrush12setTransformERK10QTransform(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBrush", "setTransform", args)
  }

}

// setMatrix(const class QMatrix &)
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
    // invoke: void setMatrix(const class QMatrix &)
    var arg0 = args[0].(QMatrix).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN6QBrush9setMatrixERK7QMatrix(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBrush", "setMatrix", args)
  }

}

// isOpaque()
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
    // invoke: bool isOpaque()
    C._ZNK6QBrush8isOpaqueEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBrush", "isOpaque", args)
  }

}

// style()
func (this *QBrush) style(args ...interface{}) () {
  // style()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QBrush5styleEv
    // invoke: Qt::BrushStyle style()
    C._ZNK6QBrush5styleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBrush", "style", args)
  }

}

// matrix()
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
    // invoke: const QMatrix & matrix()
    C._ZNK6QBrush6matrixEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBrush", "matrix", args)
  }

}

// gradient()
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
    // invoke: const QGradient * gradient()
    C._ZNK6QBrush8gradientEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBrush", "gradient", args)
  }

}

// textureImage()
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
    // invoke: QImage textureImage()
    C._ZNK6QBrush12textureImageEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBrush", "textureImage", args)
  }

}

// transform()
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
    // invoke: QTransform transform()
    C._ZNK6QBrush9transformEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBrush", "transform", args)
  }

}

// texture()
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
    // invoke: QPixmap texture()
    C._ZNK6QBrush7textureEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBrush", "texture", args)
  }

}

// color()
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
    // invoke: const QColor & color()
    C._ZNK6QBrush5colorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBrush", "color", args)
  }

}

// swap(class QBrush &)
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
    // invoke: void swap(class QBrush &)
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN6QBrush4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBrush", "swap", args)
  }

}

// QBrush()
func NewQBrush(args ...interface{}) QBrush {
  // QBrush()
  // QBrush(const class QPixmap &)
  // QBrush(const class QBrush &)
  // QBrush(const class QImage &)
  // QBrush(const class QColor &, const class QPixmap &)
  // QBrush(const class QGradient &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QImage{}) // "const QImage &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[4][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QGradient{}) // "const QGradient &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QBrushC1Ev
    // invoke: void QBrush()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN6QBrushC2Ev(qthis)
  case 1:
    // invoke: _ZN6QBrushC1ERK7QPixmap
    // invoke: void QBrush(const class QPixmap &)
    var arg0 = args[0].(QPixmap).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN6QBrushC2ERK7QPixmap(qthis, arg0)
  case 2:
    // invoke: _ZN6QBrushC1ERKS_
    // invoke: void QBrush(const class QBrush &)
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN6QBrushC2ERKS_(qthis, arg0)
  case 3:
    // invoke: _ZN6QBrushC1ERK6QImage
    // invoke: void QBrush(const class QImage &)
    var arg0 = args[0].(QImage).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN6QBrushC2ERK6QImage(qthis, arg0)
  case 4:
    // invoke: _ZN6QBrushC1ERK6QColorRK7QPixmap
    // invoke: void QBrush(const class QColor &, const class QPixmap &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPixmap).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN6QBrushC2ERK6QColorRK7QPixmap(qthis, arg0, arg1)
  case 5:
    // invoke: _ZN6QBrushC1ERK9QGradient
    // invoke: void QBrush(const class QGradient &)
    var arg0 = args[0].(QGradient).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN6QBrushC2ERK9QGradient(qthis, arg0)
  default:
    qtrt.ErrorResolve("QBrush", "QBrush", args)
  }

  return QBrush{}
}

// isDetached()
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
    // invoke: bool isDetached()
    C._ZNK6QBrush10isDetachedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBrush", "isDetached", args)
  }

}

// setTextureImage(const class QImage &)
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
    // invoke: void setTextureImage(const class QImage &)
    var arg0 = args[0].(QImage).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN6QBrush15setTextureImageERK6QImage(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBrush", "setTextureImage", args)
  }

}

// ~QBrush()
func (this *QBrush) FreeQBrush(args ...interface{}) () {
  // ~QBrush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QBrushD0Ev
    // invoke: void ~QBrush()
    C._ZN6QBrushD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBrush", "~QBrush", args)
  }

}

// setColor(const class QColor &)
func (this *QBrush) setColor(args ...interface{}) () {
  // setColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QBrush8setColorERK6QColor
    // invoke: void setColor(const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN6QBrush8setColorERK6QColor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBrush", "setColor", args)
  }

}

// interpolationMode()
func (this *QGradient) interpolationMode(args ...interface{}) () {
  // interpolationMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGradient17interpolationModeEv
    // invoke: QGradient::InterpolationMode interpolationMode()
    C._ZNK9QGradient17interpolationModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGradient", "interpolationMode", args)
  }

}

// QGradient()
func NewQGradient(args ...interface{}) QGradient {
  // QGradient()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGradientC1Ev
    // invoke: void QGradient()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN9QGradientC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QGradient", "QGradient", args)
  }

  return QGradient{}
}

// setColorAt(qreal, const class QColor &)
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
    // invoke: void setColorAt(qreal, const class QColor &)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QColor).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN9QGradient10setColorAtEdRK6QColor(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGradient", "setColorAt", args)
  }

}

// stops()
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
    // invoke: QGradientStops stops()
    C._ZNK9QGradient5stopsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGradient", "stops", args)
  }

}

// coordinateMode()
func (this *QGradient) coordinateMode(args ...interface{}) () {
  // coordinateMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGradient14coordinateModeEv
    // invoke: QGradient::CoordinateMode coordinateMode()
    C._ZNK9QGradient14coordinateModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGradient", "coordinateMode", args)
  }

}

// spread()
func (this *QGradient) spread(args ...interface{}) () {
  // spread()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGradient6spreadEv
    // invoke: QGradient::Spread spread()
    C._ZNK9QGradient6spreadEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGradient", "spread", args)
  }

}

// type()
func (this *QGradient) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGradient4typeEv
    // invoke: QGradient::Type type()
    C._ZNK9QGradient4typeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGradient", "type", args)
  }

}

// start()
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
    // invoke: QPointF start()
    C._ZNK15QLinearGradient5startEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLinearGradient", "start", args)
  }

}

// finalStop()
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
    // invoke: QPointF finalStop()
    C._ZNK15QLinearGradient9finalStopEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLinearGradient", "finalStop", args)
  }

}

// setFinalStop(qreal, qreal)
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
    // invoke: void setFinalStop(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C._ZN15QLinearGradient12setFinalStopEdd(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN15QLinearGradient12setFinalStopERK7QPointF
    // invoke: void setFinalStop(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QLinearGradient12setFinalStopERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLinearGradient", "setFinalStop", args)
  }

}

// setStart(const class QPointF &)
func (this *QLinearGradient) setStart(args ...interface{}) () {
  // setStart(const class QPointF &)
  // setStart(qreal, qreal)
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
    // invoke: _ZN15QLinearGradient8setStartERK7QPointF
    // invoke: void setStart(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QLinearGradient8setStartERK7QPointF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN15QLinearGradient8setStartEdd
    // invoke: void setStart(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C._ZN15QLinearGradient8setStartEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLinearGradient", "setStart", args)
  }

}

// QLinearGradient()
func NewQLinearGradient(args ...interface{}) QLinearGradient {
  // QLinearGradient()
  // QLinearGradient(qreal, qreal, qreal, qreal)
  // QLinearGradient(const class QPointF &, const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[2][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QLinearGradientC1Ev
    // invoke: void QLinearGradient()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN15QLinearGradientC2Ev(qthis)
  case 1:
    // invoke: _ZN15QLinearGradientC1Edddd
    // invoke: void QLinearGradient(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN15QLinearGradientC2Edddd(qthis, arg0, arg1, arg2, arg3)
  case 2:
    // invoke: _ZN15QLinearGradientC1ERK7QPointFS2_
    // invoke: void QLinearGradient(const class QPointF &, const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPointF).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN15QLinearGradientC2ERK7QPointFS2_(qthis, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLinearGradient", "QLinearGradient", args)
  }

  return QLinearGradient{}
}

// <= body block end

