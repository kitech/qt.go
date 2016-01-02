package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtGui/qtransform.h
// dst-file: /src/gui/qtransform.go
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
  // proto:  QPoint QTransform::map(const QPoint & p);
extern void _ZNK10QTransform3mapERK6QPoint(void* qthis, void* arg0);
  // proto:  QPainterPath QTransform::map(const QPainterPath & p);
extern void _ZNK10QTransform3mapERK12QPainterPath(void* qthis, void* arg0);
  // proto:  qreal QTransform::det();
extern void _ZNK10QTransform3detEv(void* qthis);
  // proto:  void QTransform::map(qreal x, qreal y, qreal * tx, qreal * ty);
extern void _ZNK10QTransform3mapEddPdS0_(void* qthis, double arg0, double arg1, double* arg2, double* arg3);
  // proto:  void QTransform::setMatrix(qreal m11, qreal m12, qreal m13, qreal m21, qreal m22, qreal m23, qreal m31, qreal m32, qreal m33);
extern void _ZN10QTransform9setMatrixEddddddddd(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4, double arg5, double arg6, double arg7, double arg8);
  // proto:  const QMatrix & QTransform::toAffine();
extern void _ZNK10QTransform8toAffineEv(void* qthis);
  // proto:  void QTransform::reset();
extern void _ZN10QTransform5resetEv(void* qthis);
  // proto:  qreal QTransform::determinant();
extern void demth_ZNK10QTransform11determinantEv(void* qthis);
  // proto: static QTransform QTransform::fromScale(qreal dx, qreal dy);
extern void _ZN10QTransform9fromScaleEdd(double arg0, double arg1);
  // proto:  bool QTransform::isTranslating();
extern void _ZNK10QTransform13isTranslatingEv(void* qthis);
  // proto:  QPolygon QTransform::mapToPolygon(const QRect & r);
extern void _ZNK10QTransform12mapToPolygonERK5QRect(void* qthis, void* arg0);
  // proto:  qreal QTransform::m22();
extern void _ZNK10QTransform3m22Ev(void* qthis);
  // proto:  QRect QTransform::mapRect(const QRect & );
extern void _ZNK10QTransform7mapRectERK5QRect(void* qthis, void* arg0);
  // proto:  void QTransform::QTransform();
extern void* dector_ZN10QTransformC1Ev();
extern void _ZN10QTransformC1Ev(void* qthis);
  // proto:  qreal QTransform::m32();
extern void _ZNK10QTransform3m32Ev(void* qthis);
  // proto:  void QTransform::map(int x, int y, int * tx, int * ty);
extern void _ZNK10QTransform3mapEiiPiS0_(void* qthis, int arg0, int arg1, int* arg2, int* arg3);
  // proto:  QTransform & QTransform::shear(qreal sh, qreal sv);
extern void _ZN10QTransform5shearEdd(void* qthis, double arg0, double arg1);
  // proto:  void QTransform::QTransform(qreal h11, qreal h12, qreal h21, qreal h22, qreal dx, qreal dy);
extern void* dector_ZN10QTransformC1Edddddd(double arg0, double arg1, double arg2, double arg3, double arg4, double arg5);
extern void _ZN10QTransformC1Edddddd(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4, double arg5);
  // proto:  QTransform & QTransform::scale(qreal sx, qreal sy);
extern void _ZN10QTransform5scaleEdd(void* qthis, double arg0, double arg1);
  // proto:  QPolygon QTransform::map(const QPolygon & a);
extern void _ZNK10QTransform3mapERK8QPolygon(void* qthis, void* arg0);
  // proto:  QTransform QTransform::transposed();
extern void _ZNK10QTransform10transposedEv(void* qthis);
  // proto:  QLineF QTransform::map(const QLineF & l);
extern void _ZNK10QTransform3mapERK6QLineF(void* qthis, void* arg0);
  // proto:  void QTransform::QTransform(bool );
extern void* dector_ZN10QTransformC1Eb(bool arg0);
extern void demth_ZN10QTransformC1Eb(void* qthis, bool arg0);
  // proto:  QTransform & QTransform::translate(qreal dx, qreal dy);
extern void _ZN10QTransform9translateEdd(void* qthis, double arg0, double arg1);
  // proto:  QRectF QTransform::mapRect(const QRectF & );
extern void _ZNK10QTransform7mapRectERK6QRectF(void* qthis, void* arg0);
  // proto: static QTransform QTransform::fromTranslate(qreal dx, qreal dy);
extern void _ZN10QTransform13fromTranslateEdd(double arg0, double arg1);
  // proto:  QLine QTransform::map(const QLine & l);
extern void _ZNK10QTransform3mapERK5QLine(void* qthis, void* arg0);
  // proto:  bool QTransform::isInvertible();
extern void _ZNK10QTransform12isInvertibleEv(void* qthis);
  // proto: static bool QTransform::quadToQuad(const QPolygonF & one, const QPolygonF & two, QTransform & result);
extern void _ZN10QTransform10quadToQuadERK9QPolygonFS2_RS_(void* arg0, void* arg1, void* arg2);
  // proto: static bool QTransform::squareToQuad(const QPolygonF & square, QTransform & result);
extern void _ZN10QTransform12squareToQuadERK9QPolygonFRS_(void* arg0, void* arg1);
  // proto:  QPointF QTransform::map(const QPointF & p);
extern void _ZNK10QTransform3mapERK7QPointF(void* qthis, void* arg0);
  // proto:  QPolygonF QTransform::map(const QPolygonF & a);
extern void _ZNK10QTransform3mapERK9QPolygonF(void* qthis, void* arg0);
  // proto:  qreal QTransform::m31();
extern void _ZNK10QTransform3m31Ev(void* qthis);
  // proto:  void QTransform::QTransform(const QMatrix & mtx);
extern void* dector_ZN10QTransformC1ERK7QMatrix(void* arg0);
extern void _ZN10QTransformC1ERK7QMatrix(void* qthis, void* arg0);
  // proto:  void QTransform::QTransform(qreal h11, qreal h12, qreal h13, qreal h21, qreal h22, qreal h23, qreal h31, qreal h32, qreal h33);
extern void* dector_ZN10QTransformC1Eddddddddd(double arg0, double arg1, double arg2, double arg3, double arg4, double arg5, double arg6, double arg7, double arg8);
extern void _ZN10QTransformC1Eddddddddd(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4, double arg5, double arg6, double arg7, double arg8);
  // proto:  QRegion QTransform::map(const QRegion & r);
extern void _ZNK10QTransform3mapERK7QRegion(void* qthis, void* arg0);
  // proto:  bool QTransform::isRotating();
extern void _ZNK10QTransform10isRotatingEv(void* qthis);
  // proto:  qreal QTransform::m33();
extern void _ZNK10QTransform3m33Ev(void* qthis);
  // proto:  qreal QTransform::m13();
extern void _ZNK10QTransform3m13Ev(void* qthis);
  // proto:  qreal QTransform::m21();
extern void _ZNK10QTransform3m21Ev(void* qthis);
  // proto:  bool QTransform::isScaling();
extern void _ZNK10QTransform9isScalingEv(void* qthis);
  // proto:  QTransform QTransform::inverted(bool * invertible);
extern void _ZNK10QTransform8invertedEPb(void* qthis, bool* arg0);
  // proto:  bool QTransform::isAffine();
extern void _ZNK10QTransform8isAffineEv(void* qthis);
  // proto:  qreal QTransform::m11();
extern void _ZNK10QTransform3m11Ev(void* qthis);
  // proto:  bool QTransform::isIdentity();
extern void _ZNK10QTransform10isIdentityEv(void* qthis);
  // proto: static bool QTransform::quadToSquare(const QPolygonF & quad, QTransform & result);
extern void _ZN10QTransform12quadToSquareERK9QPolygonFRS_(void* arg0, void* arg1);
  // proto:  QTransform QTransform::adjoint();
extern void _ZNK10QTransform7adjointEv(void* qthis);
  // proto:  qreal QTransform::dx();
extern void _ZNK10QTransform2dxEv(void* qthis);
  // proto:  qreal QTransform::m23();
extern void _ZNK10QTransform3m23Ev(void* qthis);
  // proto:  qreal QTransform::dy();
extern void _ZNK10QTransform2dyEv(void* qthis);
  // proto:  void QTransform::QTransform(qreal h11, qreal h12, qreal h13, qreal h21, qreal h22, qreal h23, qreal h31, qreal h32, qreal h33, bool );
extern void* dector_ZN10QTransformC1Edddddddddb(double arg0, double arg1, double arg2, double arg3, double arg4, double arg5, double arg6, double arg7, double arg8, bool arg9);
extern void demth_ZN10QTransformC1Edddddddddb(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4, double arg5, double arg6, double arg7, double arg8, bool arg9);
  // proto:  qreal QTransform::m12();
extern void _ZNK10QTransform3m12Ev(void* qthis);
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

// class sizeof(QTransform)=88
type QTransform struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QPoint QTransform::map(const QPoint & p);
func (this *QTransform) map_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTransform", "map", args)
  }

}

  // proto:  qreal QTransform::det();
func (this *QTransform) det(args ...interface{}) () {
  // det()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3detEv
  default:
    qtrt.ErrorResolve("QTransform", "det", args)
  }

}

  // proto:  void QTransform::setMatrix(qreal m11, qreal m12, qreal m13, qreal m21, qreal m22, qreal m23, qreal m31, qreal m32, qreal m33);
func (this *QTransform) setMatrix(args ...interface{}) () {
  // setMatrix(qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][5] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][6] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][7] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][8] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTransform9setMatrixEddddddddd
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
    var arg6 = C.double(args[6].(float64))
    if false {fmt.Println(arg6)}
    var arg7 = C.double(args[7].(float64))
    if false {fmt.Println(arg7)}
    var arg8 = C.double(args[8].(float64))
    if false {fmt.Println(arg8)}
  default:
    qtrt.ErrorResolve("QTransform", "setMatrix", args)
  }

}

  // proto:  const QMatrix & QTransform::toAffine();
func (this *QTransform) toAffine(args ...interface{}) () {
  // toAffine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform8toAffineEv
  default:
    qtrt.ErrorResolve("QTransform", "toAffine", args)
  }

}

  // proto:  void QTransform::reset();
func (this *QTransform) reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTransform5resetEv
  default:
    qtrt.ErrorResolve("QTransform", "reset", args)
  }

}

  // proto:  qreal QTransform::determinant();
func (this *QTransform) determinant(args ...interface{}) () {
  // determinant()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform11determinantEv
  default:
    qtrt.ErrorResolve("QTransform", "determinant", args)
  }

}

  // proto: static QTransform QTransform::fromScale(qreal dx, qreal dy);
func (this *QTransform) fromScale_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTransform", "fromScale", args)
  }

}

  // proto:  bool QTransform::isTranslating();
func (this *QTransform) isTranslating(args ...interface{}) () {
  // isTranslating()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform13isTranslatingEv
  default:
    qtrt.ErrorResolve("QTransform", "isTranslating", args)
  }

}

  // proto:  QPolygon QTransform::mapToPolygon(const QRect & r);
func (this *QTransform) mapToPolygon(args ...interface{}) () {
  // mapToPolygon(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform12mapToPolygonERK5QRect
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTransform", "mapToPolygon", args)
  }

}

  // proto:  qreal QTransform::m22();
func (this *QTransform) m22(args ...interface{}) () {
  // m22()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3m22Ev
  default:
    qtrt.ErrorResolve("QTransform", "m22", args)
  }

}

  // proto:  QRect QTransform::mapRect(const QRect & );
func (this *QTransform) mapRect(args ...interface{}) () {
  // mapRect(const class QRect &)
  // mapRect(const class QRectF &)
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
    // invoke: _ZNK10QTransform7mapRectERK5QRect
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK10QTransform7mapRectERK6QRectF
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTransform", "mapRect", args)
  }

}

  // proto:  void QTransform::QTransform();
func NewQTransform(args ...interface{}) QTransform {
  return QTransform{}
}

  // proto:  qreal QTransform::m32();
func (this *QTransform) m32(args ...interface{}) () {
  // m32()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3m32Ev
  default:
    qtrt.ErrorResolve("QTransform", "m32", args)
  }

}

  // proto:  QTransform & QTransform::shear(qreal sh, qreal sv);
func (this *QTransform) shear(args ...interface{}) () {
  // shear(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTransform5shearEdd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTransform", "shear", args)
  }

}

  // proto:  QTransform & QTransform::scale(qreal sx, qreal sy);
func (this *QTransform) scale(args ...interface{}) () {
  // scale(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTransform5scaleEdd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTransform", "scale", args)
  }

}

  // proto:  QTransform QTransform::transposed();
func (this *QTransform) transposed(args ...interface{}) () {
  // transposed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform10transposedEv
  default:
    qtrt.ErrorResolve("QTransform", "transposed", args)
  }

}

  // proto:  QTransform & QTransform::translate(qreal dx, qreal dy);
func (this *QTransform) translate(args ...interface{}) () {
  // translate(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTransform9translateEdd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTransform", "translate", args)
  }

}

  // proto: static QTransform QTransform::fromTranslate(qreal dx, qreal dy);
func (this *QTransform) fromTranslate_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTransform", "fromTranslate", args)
  }

}

  // proto:  bool QTransform::isInvertible();
func (this *QTransform) isInvertible(args ...interface{}) () {
  // isInvertible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform12isInvertibleEv
  default:
    qtrt.ErrorResolve("QTransform", "isInvertible", args)
  }

}

  // proto: static bool QTransform::quadToQuad(const QPolygonF & one, const QPolygonF & two, QTransform & result);
func (this *QTransform) quadToQuad_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTransform", "quadToQuad", args)
  }

}

  // proto: static bool QTransform::squareToQuad(const QPolygonF & square, QTransform & result);
func (this *QTransform) squareToQuad_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTransform", "squareToQuad", args)
  }

}

  // proto:  qreal QTransform::m31();
func (this *QTransform) m31(args ...interface{}) () {
  // m31()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3m31Ev
  default:
    qtrt.ErrorResolve("QTransform", "m31", args)
  }

}

  // proto:  bool QTransform::isRotating();
func (this *QTransform) isRotating(args ...interface{}) () {
  // isRotating()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform10isRotatingEv
  default:
    qtrt.ErrorResolve("QTransform", "isRotating", args)
  }

}

  // proto:  qreal QTransform::m33();
func (this *QTransform) m33(args ...interface{}) () {
  // m33()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3m33Ev
  default:
    qtrt.ErrorResolve("QTransform", "m33", args)
  }

}

  // proto:  qreal QTransform::m13();
func (this *QTransform) m13(args ...interface{}) () {
  // m13()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3m13Ev
  default:
    qtrt.ErrorResolve("QTransform", "m13", args)
  }

}

  // proto:  qreal QTransform::m21();
func (this *QTransform) m21(args ...interface{}) () {
  // m21()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3m21Ev
  default:
    qtrt.ErrorResolve("QTransform", "m21", args)
  }

}

  // proto:  bool QTransform::isScaling();
func (this *QTransform) isScaling(args ...interface{}) () {
  // isScaling()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform9isScalingEv
  default:
    qtrt.ErrorResolve("QTransform", "isScaling", args)
  }

}

  // proto:  QTransform QTransform::inverted(bool * invertible);
func (this *QTransform) inverted(args ...interface{}) () {
  // inverted(_Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform8invertedEPb
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTransform", "inverted", args)
  }

}

  // proto:  bool QTransform::isAffine();
func (this *QTransform) isAffine(args ...interface{}) () {
  // isAffine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform8isAffineEv
  default:
    qtrt.ErrorResolve("QTransform", "isAffine", args)
  }

}

  // proto:  qreal QTransform::m11();
func (this *QTransform) m11(args ...interface{}) () {
  // m11()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3m11Ev
  default:
    qtrt.ErrorResolve("QTransform", "m11", args)
  }

}

  // proto:  bool QTransform::isIdentity();
func (this *QTransform) isIdentity(args ...interface{}) () {
  // isIdentity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform10isIdentityEv
  default:
    qtrt.ErrorResolve("QTransform", "isIdentity", args)
  }

}

  // proto: static bool QTransform::quadToSquare(const QPolygonF & quad, QTransform & result);
func (this *QTransform) quadToSquare_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTransform", "quadToSquare", args)
  }

}

  // proto:  QTransform QTransform::adjoint();
func (this *QTransform) adjoint(args ...interface{}) () {
  // adjoint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform7adjointEv
  default:
    qtrt.ErrorResolve("QTransform", "adjoint", args)
  }

}

  // proto:  qreal QTransform::dx();
func (this *QTransform) dx(args ...interface{}) () {
  // dx()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform2dxEv
  default:
    qtrt.ErrorResolve("QTransform", "dx", args)
  }

}

  // proto:  qreal QTransform::m23();
func (this *QTransform) m23(args ...interface{}) () {
  // m23()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3m23Ev
  default:
    qtrt.ErrorResolve("QTransform", "m23", args)
  }

}

  // proto:  qreal QTransform::dy();
func (this *QTransform) dy(args ...interface{}) () {
  // dy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform2dyEv
  default:
    qtrt.ErrorResolve("QTransform", "dy", args)
  }

}

  // proto:  qreal QTransform::m12();
func (this *QTransform) m12(args ...interface{}) () {
  // m12()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3m12Ev
  default:
    qtrt.ErrorResolve("QTransform", "m12", args)
  }

}

// <= body block end

