package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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
extern double C_ZNK15QRadialGradient12centerRadiusEv(void* qthis); // 4
  // proto:  void QRadialGradient::setFocalPoint(const QPointF & focalPoint);
extern void C_ZN15QRadialGradient13setFocalPointERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QRadialGradient::setFocalPoint(qreal x, qreal y);
extern void C_ZN15QRadialGradient13setFocalPointEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QRadialGradient::setRadius(qreal radius);
extern void C_ZN15QRadialGradient9setRadiusEd(void* qthis, double arg0); // 4
  // proto:  QPointF QRadialGradient::focalPoint();
extern void* C_ZNK15QRadialGradient10focalPointEv(void* qthis); // 4
  // proto:  void QRadialGradient::QRadialGradient(qreal cx, qreal cy, qreal radius, qreal fx, qreal fy);
extern void* C_ZN15QRadialGradientC2Eddddd(double arg0, double arg1, double arg2, double arg3, double arg4); // 3
  // proto:  void QRadialGradient::QRadialGradient();
extern void* C_ZN15QRadialGradientC2Ev(); // 3
  // proto:  void QRadialGradient::QRadialGradient(const QPointF & center, qreal radius, const QPointF & focalPoint);
extern void* C_ZN15QRadialGradientC2ERK7QPointFdS2_(void* arg0, double arg1, void* arg2); // 3
  // proto:  void QRadialGradient::QRadialGradient(const QPointF & center, qreal centerRadius, const QPointF & focalPoint, qreal focalRadius);
extern void* C_ZN15QRadialGradientC2ERK7QPointFdS2_d(void* arg0, double arg1, void* arg2, double arg3); // 3
  // proto:  void QRadialGradient::QRadialGradient(qreal cx, qreal cy, qreal centerRadius, qreal fx, qreal fy, qreal focalRadius);
extern void* C_ZN15QRadialGradientC2Edddddd(double arg0, double arg1, double arg2, double arg3, double arg4, double arg5); // 3
  // proto:  void QRadialGradient::QRadialGradient(const QPointF & center, qreal radius);
extern void* C_ZN15QRadialGradientC2ERK7QPointFd(void* arg0, double arg1); // 3
  // proto:  void QRadialGradient::QRadialGradient(qreal cx, qreal cy, qreal radius);
extern void* C_ZN15QRadialGradientC2Eddd(double arg0, double arg1, double arg2); // 3
  // proto:  void QRadialGradient::setCenter(const QPointF & center);
extern void C_ZN15QRadialGradient9setCenterERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QRadialGradient::setCenter(qreal x, qreal y);
extern void C_ZN15QRadialGradient9setCenterEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QRadialGradient::setCenterRadius(qreal radius);
extern void C_ZN15QRadialGradient15setCenterRadiusEd(void* qthis, double arg0); // 4
  // proto:  qreal QRadialGradient::radius();
extern double C_ZNK15QRadialGradient6radiusEv(void* qthis); // 4
  // proto:  void QRadialGradient::setFocalRadius(qreal radius);
extern void C_ZN15QRadialGradient14setFocalRadiusEd(void* qthis, double arg0); // 4
  // proto:  qreal QRadialGradient::focalRadius();
extern double C_ZNK15QRadialGradient11focalRadiusEv(void* qthis); // 4
  // proto:  QPointF QRadialGradient::center();
extern void* C_ZNK15QRadialGradient6centerEv(void* qthis); // 4
  // proto:  void QConicalGradient::QConicalGradient(qreal cx, qreal cy, qreal startAngle);
extern void* C_ZN16QConicalGradientC2Eddd(double arg0, double arg1, double arg2); // 3
  // proto:  void QConicalGradient::QConicalGradient(const QPointF & center, qreal startAngle);
extern void* C_ZN16QConicalGradientC2ERK7QPointFd(void* arg0, double arg1); // 3
  // proto:  void QConicalGradient::QConicalGradient();
extern void* C_ZN16QConicalGradientC2Ev(); // 3
  // proto:  void QConicalGradient::setCenter(const QPointF & center);
extern void C_ZN16QConicalGradient9setCenterERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QConicalGradient::setCenter(qreal x, qreal y);
extern void C_ZN16QConicalGradient9setCenterEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  qreal QConicalGradient::angle();
extern double C_ZNK16QConicalGradient5angleEv(void* qthis); // 4
  // proto:  QPointF QConicalGradient::center();
extern void* C_ZNK16QConicalGradient6centerEv(void* qthis); // 4
  // proto:  void QConicalGradient::setAngle(qreal angle);
extern void C_ZN16QConicalGradient8setAngleEd(void* qthis, double arg0); // 4
  // proto:  void QBrush::setTexture(const QPixmap & pixmap);
extern void C_ZN6QBrush10setTextureERK7QPixmap(void* qthis, void* arg0); // 4
  // proto:  void QBrush::setTransform(const QTransform & );
extern void C_ZN6QBrush12setTransformERK10QTransform(void* qthis, void* arg0); // 4
  // proto:  void QBrush::setMatrix(const QMatrix & mat);
extern void C_ZN6QBrush9setMatrixERK7QMatrix(void* qthis, void* arg0); // 4
  // proto:  bool QBrush::isOpaque();
extern bool C_ZNK6QBrush8isOpaqueEv(void* qthis); // 4
  // proto:  Qt::BrushStyle QBrush::style();
extern void C_ZNK6QBrush5styleEv(void* qthis); // 2
  // proto:  const QMatrix & QBrush::matrix();
extern void* C_ZNK6QBrush6matrixEv(void* qthis); // 2
  // proto:  const QGradient * QBrush::gradient();
extern void* C_ZNK6QBrush8gradientEv(void* qthis); // 4
  // proto:  QImage QBrush::textureImage();
extern void* C_ZNK6QBrush12textureImageEv(void* qthis); // 4
  // proto:  QTransform QBrush::transform();
extern void* C_ZNK6QBrush9transformEv(void* qthis); // 2
  // proto:  QPixmap QBrush::texture();
extern void* C_ZNK6QBrush7textureEv(void* qthis); // 4
  // proto:  const QColor & QBrush::color();
extern void* C_ZNK6QBrush5colorEv(void* qthis); // 2
  // proto:  void QBrush::swap(QBrush & other);
extern void C_ZN6QBrush4swapERS_(void* qthis, void* arg0); // 2
  // proto:  void QBrush::QBrush();
extern void* C_ZN6QBrushC2Ev(); // 3
  // proto:  void QBrush::QBrush(const QPixmap & pixmap);
extern void* C_ZN6QBrushC2ERK7QPixmap(void* arg0); // 3
  // proto:  void QBrush::QBrush(const QBrush & brush);
extern void* C_ZN6QBrushC2ERKS_(void* arg0); // 3
  // proto:  void QBrush::QBrush(const QImage & image);
extern void* C_ZN6QBrushC2ERK6QImage(void* arg0); // 3
  // proto:  void QBrush::QBrush(const QColor & color, const QPixmap & pixmap);
extern void* C_ZN6QBrushC2ERK6QColorRK7QPixmap(void* arg0, void* arg1); // 3
  // proto:  void QBrush::QBrush(const QGradient & gradient);
extern void* C_ZN6QBrushC2ERK9QGradient(void* arg0); // 3
  // proto:  bool QBrush::isDetached();
extern bool C_ZNK6QBrush10isDetachedEv(void* qthis); // 2
  // proto:  void QBrush::setTextureImage(const QImage & image);
extern void C_ZN6QBrush15setTextureImageERK6QImage(void* qthis, void* arg0); // 4
  // proto:  void QBrush::~QBrush();
extern void C_ZN6QBrushD2Ev(void* qthis); // 4
  // proto:  void QBrush::setColor(const QColor & color);
extern void C_ZN6QBrush8setColorERK6QColor(void* qthis, void* arg0); // 4
  // proto:  QGradient::InterpolationMode QGradient::interpolationMode();
extern void C_ZNK9QGradient17interpolationModeEv(void* qthis); // 4
  // proto:  void QGradient::QGradient();
extern void* C_ZN9QGradientC2Ev(); // 3
  // proto:  void QGradient::setColorAt(qreal pos, const QColor & color);
extern void C_ZN9QGradient10setColorAtEdRK6QColor(void* qthis, double arg0, void* arg1); // 4
  // proto:  QGradientStops QGradient::stops();
extern void C_ZNK9QGradient5stopsEv(void* qthis); // 4
  // proto:  QGradient::CoordinateMode QGradient::coordinateMode();
extern void C_ZNK9QGradient14coordinateModeEv(void* qthis); // 4
  // proto:  QGradient::Spread QGradient::spread();
extern void C_ZNK9QGradient6spreadEv(void* qthis); // 2
  // proto:  QGradient::Type QGradient::type();
extern void C_ZNK9QGradient4typeEv(void* qthis); // 2
  // proto:  QPointF QLinearGradient::start();
extern void* C_ZNK15QLinearGradient5startEv(void* qthis); // 4
  // proto:  QPointF QLinearGradient::finalStop();
extern void* C_ZNK15QLinearGradient9finalStopEv(void* qthis); // 4
  // proto:  void QLinearGradient::setFinalStop(qreal x, qreal y);
extern void C_ZN15QLinearGradient12setFinalStopEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QLinearGradient::setFinalStop(const QPointF & stop);
extern void C_ZN15QLinearGradient12setFinalStopERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QLinearGradient::setStart(const QPointF & start);
extern void C_ZN15QLinearGradient8setStartERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QLinearGradient::setStart(qreal x, qreal y);
extern void C_ZN15QLinearGradient8setStartEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QLinearGradient::QLinearGradient();
extern void* C_ZN15QLinearGradientC2Ev(); // 3
  // proto:  void QLinearGradient::QLinearGradient(qreal xStart, qreal yStart, qreal xFinalStop, qreal yFinalStop);
extern void* C_ZN15QLinearGradientC2Edddd(double arg0, double arg1, double arg2, double arg3); // 3
  // proto:  void QLinearGradient::QLinearGradient(const QPointF & start, const QPointF & finalStop);
extern void* C_ZN15QLinearGradientC2ERK7QPointFS2_(void* arg0, void* arg1); // 3
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QConicalGradient)=1
type QConicalGradient struct {
  /*qbase*/ QGradient;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QBrush)=1
type QBrush struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGradient)=1
type QGradient struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QBrushData)=1
type QBrushData struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QLinearGradient)=1
type QLinearGradient struct {
  /*qbase*/ QGradient;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// centerRadius()
func (this *QRadialGradient) Centerradius(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QRadialGradient12centerRadiusEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRadialGradient", "centerRadius", args)
  }

  return
}

// setFocalPoint(const class QPointF &)
func (this *QRadialGradient) Setfocalpoint(args ...interface{}) () {
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
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QRadialGradient13setFocalPointERK7QPointF(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN15QRadialGradient13setFocalPointEdd
    // invoke: void setFocalPoint(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN15QRadialGradient13setFocalPointEdd(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QRadialGradient", "setFocalPoint", args)
  }

  return
}

// setRadius(qreal)
func (this *QRadialGradient) Setradius(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN15QRadialGradient9setRadiusEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRadialGradient", "setRadius", args)
  }

  return
}

// focalPoint()
func (this *QRadialGradient) Focalpoint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QRadialGradient10focalPointEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRadialGradient", "focalPoint", args)
  }

  return
}

// QRadialGradient(qreal, qreal, qreal, qreal, qreal)
func NewQRadialGradient(args ...interface{}) *QRadialGradient {
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
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QRadialGradientC2Eddddd(arg0, arg1, arg2, arg3, arg4)
    return &QRadialGradient{Qclsinst:qthis}
  case 1:
    // invoke: _ZN15QRadialGradientC1Ev
    // invoke: void QRadialGradient()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QRadialGradientC2Ev()
    return &QRadialGradient{Qclsinst:qthis}
  case 2:
    // invoke: _ZN15QRadialGradientC1ERK7QPointFdS2_
    // invoke: void QRadialGradient(const class QPointF &, qreal, const class QPointF &)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QPointF).Qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QRadialGradientC2ERK7QPointFdS2_(arg0, arg1, arg2)
    return &QRadialGradient{Qclsinst:qthis}
  case 3:
    // invoke: _ZN15QRadialGradientC1ERK7QPointFdS2_d
    // invoke: void QRadialGradient(const class QPointF &, qreal, const class QPointF &, qreal)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QPointF).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QRadialGradientC2ERK7QPointFdS2_d(arg0, arg1, arg2, arg3)
    return &QRadialGradient{Qclsinst:qthis}
  case 4:
    // invoke: _ZN15QRadialGradientC1Edddddd
    // invoke: void QRadialGradient(qreal, qreal, qreal, qreal, qreal, qreal)
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
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QRadialGradientC2Edddddd(arg0, arg1, arg2, arg3, arg4, arg5)
    return &QRadialGradient{Qclsinst:qthis}
  case 5:
    // invoke: _ZN15QRadialGradientC1ERK7QPointFd
    // invoke: void QRadialGradient(const class QPointF &, qreal)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QRadialGradientC2ERK7QPointFd(arg0, arg1)
    return &QRadialGradient{Qclsinst:qthis}
  case 6:
    // invoke: _ZN15QRadialGradientC1Eddd
    // invoke: void QRadialGradient(qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QRadialGradientC2Eddd(arg0, arg1, arg2)
    return &QRadialGradient{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QRadialGradient", "QRadialGradient", args)
  }

  return nil // QRadialGradient{Qclsinst:qthis}
}

// setCenter(const class QPointF &)
func (this *QRadialGradient) Setcenter(args ...interface{}) () {
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
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QRadialGradient9setCenterERK7QPointF(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN15QRadialGradient9setCenterEdd
    // invoke: void setCenter(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN15QRadialGradient9setCenterEdd(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QRadialGradient", "setCenter", args)
  }

  return
}

// setCenterRadius(qreal)
func (this *QRadialGradient) Setcenterradius(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN15QRadialGradient15setCenterRadiusEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRadialGradient", "setCenterRadius", args)
  }

  return
}

// radius()
func (this *QRadialGradient) Radius(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QRadialGradient6radiusEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRadialGradient", "radius", args)
  }

  return
}

// setFocalRadius(qreal)
func (this *QRadialGradient) Setfocalradius(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN15QRadialGradient14setFocalRadiusEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRadialGradient", "setFocalRadius", args)
  }

  return
}

// focalRadius()
func (this *QRadialGradient) Focalradius(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QRadialGradient11focalRadiusEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRadialGradient", "focalRadius", args)
  }

  return
}

// center()
func (this *QRadialGradient) Center(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QRadialGradient6centerEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRadialGradient", "center", args)
  }

  return
}

// QConicalGradient(qreal, qreal, qreal)
func NewQConicalGradient(args ...interface{}) *QConicalGradient {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QConicalGradientC2Eddd(arg0, arg1, arg2)
    return &QConicalGradient{Qclsinst:qthis}
  case 1:
    // invoke: _ZN16QConicalGradientC1ERK7QPointFd
    // invoke: void QConicalGradient(const class QPointF &, qreal)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QConicalGradientC2ERK7QPointFd(arg0, arg1)
    return &QConicalGradient{Qclsinst:qthis}
  case 2:
    // invoke: _ZN16QConicalGradientC1Ev
    // invoke: void QConicalGradient()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QConicalGradientC2Ev()
    return &QConicalGradient{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QConicalGradient", "QConicalGradient", args)
  }

  return nil // QConicalGradient{Qclsinst:qthis}
}

// setCenter(const class QPointF &)
func (this *QConicalGradient) Setcenter(args ...interface{}) () {
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
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QConicalGradient9setCenterERK7QPointF(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN16QConicalGradient9setCenterEdd
    // invoke: void setCenter(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN16QConicalGradient9setCenterEdd(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QConicalGradient", "setCenter", args)
  }

  return
}

// angle()
func (this *QConicalGradient) Angle(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK16QConicalGradient5angleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QConicalGradient", "angle", args)
  }

  return
}

// center()
func (this *QConicalGradient) Center(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK16QConicalGradient6centerEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QConicalGradient", "center", args)
  }

  return
}

// setAngle(qreal)
func (this *QConicalGradient) Setangle(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QConicalGradient8setAngleEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QConicalGradient", "setAngle", args)
  }

  return
}

// setTexture(const class QPixmap &)
func (this *QBrush) Settexture(args ...interface{}) () {
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
    var arg0 = args[0].(*QPixmap).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QBrush10setTextureERK7QPixmap(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBrush", "setTexture", args)
  }

  return
}

// setTransform(const class QTransform &)
func (this *QBrush) Settransform(args ...interface{}) () {
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
    var arg0 = args[0].(*QTransform).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QBrush12setTransformERK10QTransform(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBrush", "setTransform", args)
  }

  return
}

// setMatrix(const class QMatrix &)
func (this *QBrush) Setmatrix(args ...interface{}) () {
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
    var arg0 = args[0].(*QMatrix).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QBrush9setMatrixERK7QMatrix(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBrush", "setMatrix", args)
  }

  return
}

// isOpaque()
func (this *QBrush) Isopaque(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QBrush8isOpaqueEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBrush", "isOpaque", args)
  }

  return
}

// style()
func (this *QBrush) Style(args ...interface{}) () {
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
    C.C_ZNK6QBrush5styleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QBrush", "style", args)
  }

  return
}

// matrix()
func (this *QBrush) Matrix(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QBrush6matrixEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMatrix{}) // "const QMatrix &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBrush", "matrix", args)
  }

  return
}

// gradient()
func (this *QBrush) Gradient(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QBrush8gradientEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QGradient{}) // "const QGradient *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBrush", "gradient", args)
  }

  return
}

// textureImage()
func (this *QBrush) Textureimage(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QBrush12textureImageEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QImage{}) // "QImage"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBrush", "textureImage", args)
  }

  return
}

// transform()
func (this *QBrush) Transform(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QBrush9transformEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTransform{}) // "QTransform"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBrush", "transform", args)
  }

  return
}

// texture()
func (this *QBrush) Texture(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QBrush7textureEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPixmap{}) // "QPixmap"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBrush", "texture", args)
  }

  return
}

// color()
func (this *QBrush) Color(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QBrush5colorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "const QColor &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBrush", "color", args)
  }

  return
}

// swap(class QBrush &)
func (this *QBrush) Swap(args ...interface{}) () {
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
    var arg0 = args[0].(*QBrush).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QBrush4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBrush", "swap", args)
  }

  return
}

// QBrush()
func NewQBrush(args ...interface{}) *QBrush {
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
    qthis = C.C_ZN6QBrushC2Ev()
    return &QBrush{Qclsinst:qthis}
  case 1:
    // invoke: _ZN6QBrushC1ERK7QPixmap
    // invoke: void QBrush(const class QPixmap &)
    var arg0 = args[0].(*QPixmap).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QBrushC2ERK7QPixmap(arg0)
    return &QBrush{Qclsinst:qthis}
  case 2:
    // invoke: _ZN6QBrushC1ERKS_
    // invoke: void QBrush(const class QBrush &)
    var arg0 = args[0].(*QBrush).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QBrushC2ERKS_(arg0)
    return &QBrush{Qclsinst:qthis}
  case 3:
    // invoke: _ZN6QBrushC1ERK6QImage
    // invoke: void QBrush(const class QImage &)
    var arg0 = args[0].(*QImage).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QBrushC2ERK6QImage(arg0)
    return &QBrush{Qclsinst:qthis}
  case 4:
    // invoke: _ZN6QBrushC1ERK6QColorRK7QPixmap
    // invoke: void QBrush(const class QColor &, const class QPixmap &)
    var arg0 = args[0].(*QColor).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPixmap).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QBrushC2ERK6QColorRK7QPixmap(arg0, arg1)
    return &QBrush{Qclsinst:qthis}
  case 5:
    // invoke: _ZN6QBrushC1ERK9QGradient
    // invoke: void QBrush(const class QGradient &)
    var arg0 = args[0].(*QGradient).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QBrushC2ERK9QGradient(arg0)
    return &QBrush{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QBrush", "QBrush", args)
  }

  return nil // QBrush{Qclsinst:qthis}
}

// isDetached()
func (this *QBrush) Isdetached(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QBrush10isDetachedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBrush", "isDetached", args)
  }

  return
}

// setTextureImage(const class QImage &)
func (this *QBrush) Settextureimage(args ...interface{}) () {
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
    var arg0 = args[0].(*QImage).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QBrush15setTextureImageERK6QImage(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBrush", "setTextureImage", args)
  }

  return
}

// ~QBrush()
func (this *QBrush) Freeqbrush(args ...interface{}) () {
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
    C.C_ZN6QBrushD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QBrush", "~QBrush", args)
  }

  return
}

// setColor(const class QColor &)
func (this *QBrush) Setcolor(args ...interface{}) () {
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
    var arg0 = args[0].(*QColor).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QBrush8setColorERK6QColor(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBrush", "setColor", args)
  }

  return
}

// interpolationMode()
func (this *QGradient) Interpolationmode(args ...interface{}) () {
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
    C.C_ZNK9QGradient17interpolationModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGradient", "interpolationMode", args)
  }

  return
}

// QGradient()
func NewQGradient(args ...interface{}) *QGradient {
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
    qthis = C.C_ZN9QGradientC2Ev()
    return &QGradient{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGradient", "QGradient", args)
  }

  return nil // QGradient{Qclsinst:qthis}
}

// setColorAt(qreal, const class QColor &)
func (this *QGradient) Setcolorat(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QColor).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN9QGradient10setColorAtEdRK6QColor(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGradient", "setColorAt", args)
  }

  return
}

// stops()
func (this *QGradient) Stops(args ...interface{}) () {
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
    C.C_ZNK9QGradient5stopsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGradient", "stops", args)
  }

  return
}

// coordinateMode()
func (this *QGradient) Coordinatemode(args ...interface{}) () {
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
    C.C_ZNK9QGradient14coordinateModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGradient", "coordinateMode", args)
  }

  return
}

// spread()
func (this *QGradient) Spread(args ...interface{}) () {
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
    C.C_ZNK9QGradient6spreadEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGradient", "spread", args)
  }

  return
}

// type()
func (this *QGradient) Type_(args ...interface{}) () {
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
    C.C_ZNK9QGradient4typeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGradient", "type", args)
  }

  return
}

// start()
func (this *QLinearGradient) Start(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QLinearGradient5startEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLinearGradient", "start", args)
  }

  return
}

// finalStop()
func (this *QLinearGradient) Finalstop(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QLinearGradient9finalStopEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLinearGradient", "finalStop", args)
  }

  return
}

// setFinalStop(qreal, qreal)
func (this *QLinearGradient) Setfinalstop(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN15QLinearGradient12setFinalStopEdd(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN15QLinearGradient12setFinalStopERK7QPointF
    // invoke: void setFinalStop(const class QPointF &)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QLinearGradient12setFinalStopERK7QPointF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLinearGradient", "setFinalStop", args)
  }

  return
}

// setStart(const class QPointF &)
func (this *QLinearGradient) Setstart(args ...interface{}) () {
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
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QLinearGradient8setStartERK7QPointF(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN15QLinearGradient8setStartEdd
    // invoke: void setStart(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN15QLinearGradient8setStartEdd(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLinearGradient", "setStart", args)
  }

  return
}

// QLinearGradient()
func NewQLinearGradient(args ...interface{}) *QLinearGradient {
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
    qthis = C.C_ZN15QLinearGradientC2Ev()
    return &QLinearGradient{Qclsinst:qthis}
  case 1:
    // invoke: _ZN15QLinearGradientC1Edddd
    // invoke: void QLinearGradient(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QLinearGradientC2Edddd(arg0, arg1, arg2, arg3)
    return &QLinearGradient{Qclsinst:qthis}
  case 2:
    // invoke: _ZN15QLinearGradientC1ERK7QPointFS2_
    // invoke: void QLinearGradient(const class QPointF &, const class QPointF &)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPointF).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QLinearGradientC2ERK7QPointFS2_(arg0, arg1)
    return &QLinearGradient{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QLinearGradient", "QLinearGradient", args)
  }

  return nil // QLinearGradient{Qclsinst:qthis}
}

// <= body block end

