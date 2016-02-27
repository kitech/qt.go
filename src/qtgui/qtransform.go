package qtgui
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
import "runtime"
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
  // proto:  qreal QTransform::dx();
extern double C_ZNK10QTransform2dxEv(void* qthis); // 2
  // proto:  void QTransform::setMatrix(qreal m11, qreal m12, qreal m13, qreal m21, qreal m22, qreal m23, qreal m31, qreal m32, qreal m33);
extern void C_ZN10QTransform9setMatrixEddddddddd(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4, double arg5, double arg6, double arg7, double arg8); // 4
  // proto:  QRect QTransform::mapRect(const QRect & );
extern void* C_ZNK10QTransform7mapRectERK5QRect(void* qthis, void* arg0); // 4
  // proto:  QRectF QTransform::mapRect(const QRectF & );
extern void* C_ZNK10QTransform7mapRectERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  qreal QTransform::m21();
extern double C_ZNK10QTransform3m21Ev(void* qthis); // 2
  // proto:  qreal QTransform::m11();
extern double C_ZNK10QTransform3m11Ev(void* qthis); // 2
  // proto:  qreal QTransform::m13();
extern double C_ZNK10QTransform3m13Ev(void* qthis); // 2
  // proto:  bool QTransform::isInvertible();
extern bool C_ZNK10QTransform12isInvertibleEv(void* qthis); // 2
  // proto:  qreal QTransform::m23();
extern double C_ZNK10QTransform3m23Ev(void* qthis); // 2
  // proto:  QTransform & QTransform::scale(qreal sx, qreal sy);
extern void* C_ZN10QTransform5scaleEdd(void* qthis, double arg0, double arg1); // 4
  // proto:  QPolygon QTransform::mapToPolygon(const QRect & r);
extern void* C_ZNK10QTransform12mapToPolygonERK5QRect(void* qthis, void* arg0); // 4
  // proto:  QTransform QTransform::adjoint();
extern void* C_ZNK10QTransform7adjointEv(void* qthis); // 4
  // proto:  qreal QTransform::m33();
extern double C_ZNK10QTransform3m33Ev(void* qthis); // 2
  // proto:  qreal QTransform::m32();
extern double C_ZNK10QTransform3m32Ev(void* qthis); // 2
  // proto:  qreal QTransform::m31();
extern double C_ZNK10QTransform3m31Ev(void* qthis); // 2
  // proto:  QTransform & QTransform::translate(qreal dx, qreal dy);
extern void* C_ZN10QTransform9translateEdd(void* qthis, double arg0, double arg1); // 4
  // proto:  QTransform::TransformationType QTransform::type();
extern void C_ZNK10QTransform4typeEv(void* qthis); // 4
  // proto:  QTransform & QTransform::shear(qreal sh, qreal sv);
extern void* C_ZN10QTransform5shearEdd(void* qthis, double arg0, double arg1); // 4
  // proto:  QPoint QTransform::map(const QPoint & p);
extern void* C_ZNK10QTransform3mapERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QPainterPath QTransform::map(const QPainterPath & p);
extern void* C_ZNK10QTransform3mapERK12QPainterPath(void* qthis, void* arg0); // 4
  // proto:  QLine QTransform::map(const QLine & l);
extern void* C_ZNK10QTransform3mapERK5QLine(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QTransform::map(const QPolygonF & a);
extern void* C_ZNK10QTransform3mapERK9QPolygonF(void* qthis, void* arg0); // 4
  // proto:  QPointF QTransform::map(const QPointF & p);
extern void* C_ZNK10QTransform3mapERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QTransform::map(qreal x, qreal y, qreal * tx, qreal * ty);
extern void C_ZNK10QTransform3mapEddPdS0_(void* qthis, double arg0, double arg1, void* arg2, void* arg3); // 4
  // proto:  QPolygon QTransform::map(const QPolygon & a);
extern void* C_ZNK10QTransform3mapERK8QPolygon(void* qthis, void* arg0); // 4
  // proto:  void QTransform::map(int x, int y, int * tx, int * ty);
extern void C_ZNK10QTransform3mapEiiPiS0_(void* qthis, int32_t arg0, int32_t arg1, void* arg2, void* arg3); // 4
  // proto:  QRegion QTransform::map(const QRegion & r);
extern void* C_ZNK10QTransform3mapERK7QRegion(void* qthis, void* arg0); // 4
  // proto:  QLineF QTransform::map(const QLineF & l);
extern void* C_ZNK10QTransform3mapERK6QLineF(void* qthis, void* arg0); // 4
  // proto:  qreal QTransform::determinant();
extern double C_ZNK10QTransform11determinantEv(void* qthis); // 2
  // proto:  qreal QTransform::dy();
extern double C_ZNK10QTransform2dyEv(void* qthis); // 2
  // proto:  const QMatrix & QTransform::toAffine();
extern void* C_ZNK10QTransform8toAffineEv(void* qthis); // 4
  // proto:  bool QTransform::isRotating();
extern bool C_ZNK10QTransform10isRotatingEv(void* qthis); // 2
  // proto:  QTransform QTransform::inverted(bool * invertible);
extern void* C_ZNK10QTransform8invertedEPb(void* qthis, void* arg0); // 4
  // proto: static bool QTransform::squareToQuad(const QPolygonF & square, QTransform & result);
extern bool C_ZN10QTransform12squareToQuadERK9QPolygonFRS_(void* arg0, void* arg1); // 4
  // proto:  void QTransform::reset();
extern void C_ZN10QTransform5resetEv(void* qthis); // 4
  // proto:  QTransform QTransform::transposed();
extern void* C_ZNK10QTransform10transposedEv(void* qthis); // 4
  // proto:  qreal QTransform::det();
extern double C_ZNK10QTransform3detEv(void* qthis); // 2
  // proto:  void QTransform::QTransform();
extern void* C_ZN10QTransformC2Ev(); // 3
  // proto:  void QTransform::QTransform(const QMatrix & mtx);
extern void* C_ZN10QTransformC2ERK7QMatrix(void* arg0); // 3
  // proto:  void QTransform::QTransform(qreal h11, qreal h12, qreal h13, qreal h21, qreal h22, qreal h23, qreal h31, qreal h32, qreal h33);
extern void* C_ZN10QTransformC2Eddddddddd(double arg0, double arg1, double arg2, double arg3, double arg4, double arg5, double arg6, double arg7, double arg8); // 3
  // proto:  void QTransform::QTransform(qreal h11, qreal h12, qreal h21, qreal h22, qreal dx, qreal dy);
extern void* C_ZN10QTransformC2Edddddd(double arg0, double arg1, double arg2, double arg3, double arg4, double arg5); // 3
  // proto:  qreal QTransform::m22();
extern double C_ZNK10QTransform3m22Ev(void* qthis); // 2
  // proto: static QTransform QTransform::fromScale(qreal dx, qreal dy);
extern void* C_ZN10QTransform9fromScaleEdd(double arg0, double arg1); // 4
  // proto: static bool QTransform::quadToQuad(const QPolygonF & one, const QPolygonF & two, QTransform & result);
extern bool C_ZN10QTransform10quadToQuadERK9QPolygonFS2_RS_(void* arg0, void* arg1, void* arg2); // 4
  // proto:  qreal QTransform::m12();
extern double C_ZNK10QTransform3m12Ev(void* qthis); // 2
  // proto:  bool QTransform::isIdentity();
extern bool C_ZNK10QTransform10isIdentityEv(void* qthis); // 2
  // proto:  bool QTransform::isScaling();
extern bool C_ZNK10QTransform9isScalingEv(void* qthis); // 2
  // proto: static bool QTransform::quadToSquare(const QPolygonF & quad, QTransform & result);
extern bool C_ZN10QTransform12quadToSquareERK9QPolygonFRS_(void* arg0, void* arg1); // 4
  // proto:  bool QTransform::isAffine();
extern bool C_ZNK10QTransform8isAffineEv(void* qthis); // 2
  // proto:  bool QTransform::isTranslating();
extern bool C_ZNK10QTransform13isTranslatingEv(void* qthis); // 2
  // proto: static QTransform QTransform::fromTranslate(qreal dx, qreal dy);
extern void* C_ZN10QTransform13fromTranslateEdd(double arg0, double arg1); // 4
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
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QTransform)=88
type QTransform struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// dx()
func (this *QTransform) Dx(args ...interface{}) (ret interface{}) {
  // dx()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform2dxEv
    // invoke: qreal dx()
    var ret0 = C.C_ZNK10QTransform2dxEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "dx", args)
  }

  return
}

// setMatrix(qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal)
func (this *QTransform) SetMatrix(args ...interface{}) () {
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
    // invoke: void setMatrix(qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal)
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
    var arg6 = C.double(qtrt.PrimConv(args[6], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg6)}
    var arg7 = C.double(qtrt.PrimConv(args[7], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg7)}
    var arg8 = C.double(qtrt.PrimConv(args[8], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg8)}
    C.C_ZN10QTransform9setMatrixEddddddddd(this.Qclsinst, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
  default:
    qtrt.ErrorResolve("QTransform", "setMatrix", args)
  }

  return
}

// mapRect(const class QRect &)
func (this *QTransform) MapRect(args ...interface{}) (ret interface{}) {
  // mapRect(const class QRect &)
  // mapRect(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform7mapRectERK5QRect
    // invoke: QRect mapRect(const class QRect &)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTransform7mapRectERK5QRect(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK10QTransform7mapRectERK6QRectF
    // invoke: QRectF mapRect(const class QRectF &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTransform7mapRectERK6QRectF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "mapRect", args)
  }

  return
}

// m21()
func (this *QTransform) M21(args ...interface{}) (ret interface{}) {
  // m21()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3m21Ev
    // invoke: qreal m21()
    var ret0 = C.C_ZNK10QTransform3m21Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "m21", args)
  }

  return
}

// m11()
func (this *QTransform) M11(args ...interface{}) (ret interface{}) {
  // m11()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3m11Ev
    // invoke: qreal m11()
    var ret0 = C.C_ZNK10QTransform3m11Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "m11", args)
  }

  return
}

// m13()
func (this *QTransform) M13(args ...interface{}) (ret interface{}) {
  // m13()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3m13Ev
    // invoke: qreal m13()
    var ret0 = C.C_ZNK10QTransform3m13Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "m13", args)
  }

  return
}

// isInvertible()
func (this *QTransform) IsInvertible(args ...interface{}) (ret interface{}) {
  // isInvertible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform12isInvertibleEv
    // invoke: bool isInvertible()
    var ret0 = C.C_ZNK10QTransform12isInvertibleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "isInvertible", args)
  }

  return
}

// m23()
func (this *QTransform) M23(args ...interface{}) (ret interface{}) {
  // m23()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3m23Ev
    // invoke: qreal m23()
    var ret0 = C.C_ZNK10QTransform3m23Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "m23", args)
  }

  return
}

// scale(qreal, qreal)
func (this *QTransform) Scale(args ...interface{}) (ret interface{}) {
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
    // invoke: QTransform & scale(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QTransform5scaleEdd(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTransform{}) // "QTransform &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "scale", args)
  }

  return
}

// mapToPolygon(const class QRect &)
func (this *QTransform) MapToPolygon(args ...interface{}) (ret interface{}) {
  // mapToPolygon(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform12mapToPolygonERK5QRect
    // invoke: QPolygon mapToPolygon(const class QRect &)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTransform12mapToPolygonERK5QRect(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygon{}) // "QPolygon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "mapToPolygon", args)
  }

  return
}

// adjoint()
func (this *QTransform) Adjoint(args ...interface{}) (ret interface{}) {
  // adjoint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform7adjointEv
    // invoke: QTransform adjoint()
    var ret0 = C.C_ZNK10QTransform7adjointEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTransform{}) // "QTransform"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "adjoint", args)
  }

  return
}

// m33()
func (this *QTransform) M33(args ...interface{}) (ret interface{}) {
  // m33()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3m33Ev
    // invoke: qreal m33()
    var ret0 = C.C_ZNK10QTransform3m33Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "m33", args)
  }

  return
}

// m32()
func (this *QTransform) M32(args ...interface{}) (ret interface{}) {
  // m32()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3m32Ev
    // invoke: qreal m32()
    var ret0 = C.C_ZNK10QTransform3m32Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "m32", args)
  }

  return
}

// m31()
func (this *QTransform) M31(args ...interface{}) (ret interface{}) {
  // m31()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3m31Ev
    // invoke: qreal m31()
    var ret0 = C.C_ZNK10QTransform3m31Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "m31", args)
  }

  return
}

// translate(qreal, qreal)
func (this *QTransform) Translate(args ...interface{}) (ret interface{}) {
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
    // invoke: QTransform & translate(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QTransform9translateEdd(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTransform{}) // "QTransform &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "translate", args)
  }

  return
}

// type()
func (this *QTransform) Type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform4typeEv
    // invoke: QTransform::TransformationType type()
    C.C_ZNK10QTransform4typeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTransform", "type", args)
  }

  return
}

// shear(qreal, qreal)
func (this *QTransform) Shear(args ...interface{}) (ret interface{}) {
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
    // invoke: QTransform & shear(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QTransform5shearEdd(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTransform{}) // "QTransform &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "shear", args)
  }

  return
}

// map(const class QPoint &)
func (this *QTransform) Map_(args ...interface{}) (ret interface{}) {
  // map(const class QPoint &)
  // map(const class QPainterPath &)
  // map(const class QLine &)
  // map(const class QPolygonF &)
  // map(const class QPointF &)
  // map(qreal, qreal, qreal *, qreal *)
  // map(const class QPolygon &)
  // map(int, int, int *, int *)
  // map(const class QRegion &)
  // map(const class QLineF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QLine{}) // "const QLine &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[5][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[5][2] = qtrt.DoubleTy(true) // "qreal *"
  vtys[5][3] = qtrt.DoubleTy(true) // "qreal *"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QPolygon{}) // "const QPolygon &"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "int"
  vtys[7][1] = qtrt.Int32Ty(false) // "int"
  vtys[7][2] = qtrt.Int32Ty(true) // "int *"
  vtys[7][3] = qtrt.Int32Ty(true) // "int *"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = reflect.TypeOf(qtcore.QLineF{}) // "const QLineF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3mapERK6QPoint
    // invoke: QPoint map(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTransform3mapERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK10QTransform3mapERK12QPainterPath
    // invoke: QPainterPath map(const class QPainterPath &)
    var arg0 = args[0].(*QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTransform3mapERK12QPainterPath(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK10QTransform3mapERK5QLine
    // invoke: QLine map(const class QLine &)
    var arg0 = args[0].(*qtcore.QLine).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTransform3mapERK5QLine(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QLine{}) // "QLine"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZNK10QTransform3mapERK9QPolygonF
    // invoke: QPolygonF map(const class QPolygonF &)
    var arg0 = args[0].(*QPolygonF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTransform3mapERK9QPolygonF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 4:
    // invoke: _ZNK10QTransform3mapERK7QPointF
    // invoke: QPointF map(const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTransform3mapERK7QPointF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 5:
    // invoke: _ZNK10QTransform3mapEddPdS0_
    // invoke: void map(qreal, qreal, qreal *, qreal *)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*float64))
    if false {fmt.Println(arg3)}
    C.C_ZNK10QTransform3mapEddPdS0_(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 6:
    // invoke: _ZNK10QTransform3mapERK8QPolygon
    // invoke: QPolygon map(const class QPolygon &)
    var arg0 = args[0].(*QPolygon).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTransform3mapERK8QPolygon(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygon{}) // "QPolygon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 7:
    // invoke: _ZNK10QTransform3mapEiiPiS0_
    // invoke: void map(int, int, int *, int *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    C.C_ZNK10QTransform3mapEiiPiS0_(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 8:
    // invoke: _ZNK10QTransform3mapERK7QRegion
    // invoke: QRegion map(const class QRegion &)
    var arg0 = args[0].(*QRegion).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTransform3mapERK7QRegion(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRegion{}) // "QRegion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 9:
    // invoke: _ZNK10QTransform3mapERK6QLineF
    // invoke: QLineF map(const class QLineF &)
    var arg0 = args[0].(*qtcore.QLineF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTransform3mapERK6QLineF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QLineF{}) // "QLineF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "map", args)
  }

  return
}

// determinant()
func (this *QTransform) Determinant(args ...interface{}) (ret interface{}) {
  // determinant()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform11determinantEv
    // invoke: qreal determinant()
    var ret0 = C.C_ZNK10QTransform11determinantEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "determinant", args)
  }

  return
}

// dy()
func (this *QTransform) Dy(args ...interface{}) (ret interface{}) {
  // dy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform2dyEv
    // invoke: qreal dy()
    var ret0 = C.C_ZNK10QTransform2dyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "dy", args)
  }

  return
}

// toAffine()
func (this *QTransform) ToAffine(args ...interface{}) (ret interface{}) {
  // toAffine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform8toAffineEv
    // invoke: const QMatrix & toAffine()
    var ret0 = C.C_ZNK10QTransform8toAffineEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMatrix{}) // "const QMatrix &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "toAffine", args)
  }

  return
}

// isRotating()
func (this *QTransform) IsRotating(args ...interface{}) (ret interface{}) {
  // isRotating()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform10isRotatingEv
    // invoke: bool isRotating()
    var ret0 = C.C_ZNK10QTransform10isRotatingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "isRotating", args)
  }

  return
}

// inverted(_Bool *)
func (this *QTransform) Inverted(args ...interface{}) (ret interface{}) {
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
    // invoke: QTransform inverted(_Bool *)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTransform8invertedEPb(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTransform{}) // "QTransform"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "inverted", args)
  }

  return
}

// squareToQuad(const class QPolygonF &, class QTransform &)
func (this *QTransform) SquareToQuad_s(args ...interface{}) (ret interface{}) {
  // squareToQuad(const class QPolygonF &, class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[0][1] = reflect.TypeOf(QTransform{}) // "QTransform &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTransform12squareToQuadERK9QPolygonFRS_
    // invoke: bool squareToQuad(const class QPolygonF &, class QTransform &)
    var arg0 = args[0].(*QPolygonF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QTransform).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QTransform12squareToQuadERK9QPolygonFRS_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "squareToQuad", args)
  }

  return
}

// reset()
func (this *QTransform) Reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTransform5resetEv
    // invoke: void reset()
    C.C_ZN10QTransform5resetEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTransform", "reset", args)
  }

  return
}

// transposed()
func (this *QTransform) Transposed(args ...interface{}) (ret interface{}) {
  // transposed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform10transposedEv
    // invoke: QTransform transposed()
    var ret0 = C.C_ZNK10QTransform10transposedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTransform{}) // "QTransform"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "transposed", args)
  }

  return
}

// det()
func (this *QTransform) Det(args ...interface{}) (ret interface{}) {
  // det()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3detEv
    // invoke: qreal det()
    var ret0 = C.C_ZNK10QTransform3detEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "det", args)
  }

  return
}

// QTransform()
func GcfreeQTransform(this *QTransform) {
  qtrt.UniverseFree(this)
}
func NewQTransform(args ...interface{}) *QTransform {
  // QTransform()
  // QTransform(const class QMatrix &)
  // QTransform(qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal)
  // QTransform(qreal, qreal, qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QMatrix{}) // "const QMatrix &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][5] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][6] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][7] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][8] = qtrt.DoubleTy(false) // "qreal"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][5] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTransformC1Ev
    // invoke: void QTransform()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QTransformC2Ev()
    this := &QTransform{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQTransform)
    return this
  case 1:
    // invoke: _ZN10QTransformC1ERK7QMatrix
    // invoke: void QTransform(const class QMatrix &)
    var arg0 = args[0].(*QMatrix).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QTransformC2ERK7QMatrix(arg0)
    this := &QTransform{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQTransform)
    return this
  case 2:
    // invoke: _ZN10QTransformC1Eddddddddd
    // invoke: void QTransform(qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal)
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
    var arg6 = C.double(qtrt.PrimConv(args[6], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg6)}
    var arg7 = C.double(qtrt.PrimConv(args[7], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg7)}
    var arg8 = C.double(qtrt.PrimConv(args[8], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg8)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QTransformC2Eddddddddd(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    this := &QTransform{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQTransform)
    return this
  case 3:
    // invoke: _ZN10QTransformC1Edddddd
    // invoke: void QTransform(qreal, qreal, qreal, qreal, qreal, qreal)
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
    qthis = C.C_ZN10QTransformC2Edddddd(arg0, arg1, arg2, arg3, arg4, arg5)
    this := &QTransform{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQTransform)
    return this
  default:
    qtrt.ErrorResolve("QTransform", "QTransform", args)
  }

  return nil // QTransform{Qclsinst:qthis}
}

// m22()
func (this *QTransform) M22(args ...interface{}) (ret interface{}) {
  // m22()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3m22Ev
    // invoke: qreal m22()
    var ret0 = C.C_ZNK10QTransform3m22Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "m22", args)
  }

  return
}

// fromScale(qreal, qreal)
func (this *QTransform) FromScale_s(args ...interface{}) (ret interface{}) {
  // fromScale(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTransform9fromScaleEdd
    // invoke: QTransform fromScale(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QTransform9fromScaleEdd(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTransform{}) // "QTransform"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "fromScale", args)
  }

  return
}

// quadToQuad(const class QPolygonF &, const class QPolygonF &, class QTransform &)
func (this *QTransform) QuadToQuad_s(args ...interface{}) (ret interface{}) {
  // quadToQuad(const class QPolygonF &, const class QPolygonF &, class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[0][1] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[0][2] = reflect.TypeOf(QTransform{}) // "QTransform &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTransform10quadToQuadERK9QPolygonFS2_RS_
    // invoke: bool quadToQuad(const class QPolygonF &, const class QPolygonF &, class QTransform &)
    var arg0 = args[0].(*QPolygonF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPolygonF).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QTransform).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN10QTransform10quadToQuadERK9QPolygonFS2_RS_(arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "quadToQuad", args)
  }

  return
}

// m12()
func (this *QTransform) M12(args ...interface{}) (ret interface{}) {
  // m12()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3m12Ev
    // invoke: qreal m12()
    var ret0 = C.C_ZNK10QTransform3m12Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "m12", args)
  }

  return
}

// isIdentity()
func (this *QTransform) IsIdentity(args ...interface{}) (ret interface{}) {
  // isIdentity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform10isIdentityEv
    // invoke: bool isIdentity()
    var ret0 = C.C_ZNK10QTransform10isIdentityEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "isIdentity", args)
  }

  return
}

// isScaling()
func (this *QTransform) IsScaling(args ...interface{}) (ret interface{}) {
  // isScaling()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform9isScalingEv
    // invoke: bool isScaling()
    var ret0 = C.C_ZNK10QTransform9isScalingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "isScaling", args)
  }

  return
}

// quadToSquare(const class QPolygonF &, class QTransform &)
func (this *QTransform) QuadToSquare_s(args ...interface{}) (ret interface{}) {
  // quadToSquare(const class QPolygonF &, class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[0][1] = reflect.TypeOf(QTransform{}) // "QTransform &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTransform12quadToSquareERK9QPolygonFRS_
    // invoke: bool quadToSquare(const class QPolygonF &, class QTransform &)
    var arg0 = args[0].(*QPolygonF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QTransform).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QTransform12quadToSquareERK9QPolygonFRS_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "quadToSquare", args)
  }

  return
}

// isAffine()
func (this *QTransform) IsAffine(args ...interface{}) (ret interface{}) {
  // isAffine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform8isAffineEv
    // invoke: bool isAffine()
    var ret0 = C.C_ZNK10QTransform8isAffineEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "isAffine", args)
  }

  return
}

// isTranslating()
func (this *QTransform) IsTranslating(args ...interface{}) (ret interface{}) {
  // isTranslating()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform13isTranslatingEv
    // invoke: bool isTranslating()
    var ret0 = C.C_ZNK10QTransform13isTranslatingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "isTranslating", args)
  }

  return
}

// fromTranslate(qreal, qreal)
func (this *QTransform) FromTranslate_s(args ...interface{}) (ret interface{}) {
  // fromTranslate(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTransform13fromTranslateEdd
    // invoke: QTransform fromTranslate(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QTransform13fromTranslateEdd(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTransform{}) // "QTransform"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTransform", "fromTranslate", args)
  }

  return
}

// <= body block end

