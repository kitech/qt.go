package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtGui/qmatrix.h
// dst-file: /src/gui/qmatrix.go
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
  // proto:  qreal QMatrix::dx();
extern void _ZNK7QMatrix2dxEv(void* qthis);
  // proto:  void QMatrix::QMatrix(bool );
extern void* dector_ZN7QMatrixC1Eb(bool arg0);
extern void demth_ZN7QMatrixC1Eb(void* qthis, bool arg0);
  // proto:  qreal QMatrix::dy();
extern void _ZNK7QMatrix2dyEv(void* qthis);
  // proto:  QMatrix & QMatrix::scale(qreal sx, qreal sy);
extern void _ZN7QMatrix5scaleEdd(void* qthis, double arg0, double arg1);
  // proto:  QMatrix & QMatrix::translate(qreal dx, qreal dy);
extern void _ZN7QMatrix9translateEdd(void* qthis, double arg0, double arg1);
  // proto:  qreal QMatrix::determinant();
extern void _ZNK7QMatrix11determinantEv(void* qthis);
  // proto:  QMatrix & QMatrix::shear(qreal sh, qreal sv);
extern void _ZN7QMatrix5shearEdd(void* qthis, double arg0, double arg1);
  // proto:  void QMatrix::QMatrix();
extern void* dector_ZN7QMatrixC1Ev();
extern void _ZN7QMatrixC1Ev(void* qthis);
  // proto:  qreal QMatrix::m21();
extern void _ZNK7QMatrix3m21Ev(void* qthis);
  // proto:  QPointF QMatrix::map(const QPointF & p);
extern void _ZNK7QMatrix3mapERK7QPointF(void* qthis, void* arg0);
  // proto:  QPolygonF QMatrix::map(const QPolygonF & a);
extern void _ZNK7QMatrix3mapERK9QPolygonF(void* qthis, void* arg0);
  // proto:  void QMatrix::map(qreal x, qreal y, qreal * tx, qreal * ty);
extern void _ZNK7QMatrix3mapEddPdS0_(void* qthis, double arg0, double arg1, double* arg2, double* arg3);
  // proto:  QMatrix & QMatrix::rotate(qreal a);
extern void _ZN7QMatrix6rotateEd(void* qthis, double arg0);
  // proto:  QRegion QMatrix::map(const QRegion & r);
extern void _ZNK7QMatrix3mapERK7QRegion(void* qthis, void* arg0);
  // proto:  void QMatrix::setMatrix(qreal m11, qreal m12, qreal m21, qreal m22, qreal dx, qreal dy);
extern void _ZN7QMatrix9setMatrixEdddddd(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4, double arg5);
  // proto:  void QMatrix::QMatrix(const QMatrix & matrix);
extern void* dector_ZN7QMatrixC1ERKS_(void* arg0);
extern void _ZN7QMatrixC1ERKS_(void* qthis, void* arg0);
  // proto:  void QMatrix::reset();
extern void _ZN7QMatrix5resetEv(void* qthis);
  // proto:  QLineF QMatrix::map(const QLineF & l);
extern void _ZNK7QMatrix3mapERK6QLineF(void* qthis, void* arg0);
  // proto:  QPainterPath QMatrix::map(const QPainterPath & p);
extern void _ZNK7QMatrix3mapERK12QPainterPath(void* qthis, void* arg0);
  // proto:  qreal QMatrix::m11();
extern void _ZNK7QMatrix3m11Ev(void* qthis);
  // proto:  QPolygon QMatrix::mapToPolygon(const QRect & r);
extern void _ZNK7QMatrix12mapToPolygonERK5QRect(void* qthis, void* arg0);
  // proto:  QMatrix QMatrix::inverted(bool * invertible);
extern void _ZNK7QMatrix8invertedEPb(void* qthis, bool* arg0);
  // proto:  QPoint QMatrix::map(const QPoint & p);
extern void _ZNK7QMatrix3mapERK6QPoint(void* qthis, void* arg0);
  // proto:  void QMatrix::map(int x, int y, int * tx, int * ty);
extern void _ZNK7QMatrix3mapEiiPiS0_(void* qthis, int arg0, int arg1, int* arg2, int* arg3);
  // proto:  QLine QMatrix::map(const QLine & l);
extern void _ZNK7QMatrix3mapERK5QLine(void* qthis, void* arg0);
  // proto:  QRectF QMatrix::mapRect(const QRectF & );
extern void _ZNK7QMatrix7mapRectERK6QRectF(void* qthis, void* arg0);
  // proto:  bool QMatrix::isIdentity();
extern void demth_ZNK7QMatrix10isIdentityEv(void* qthis);
  // proto:  void QMatrix::QMatrix(qreal am11, qreal am12, qreal am21, qreal am22, qreal adx, qreal ady, bool );
extern void* dector_ZN7QMatrixC1Eddddddb(double arg0, double arg1, double arg2, double arg3, double arg4, double arg5, bool arg6);
extern void demth_ZN7QMatrixC1Eddddddb(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4, double arg5, bool arg6);
  // proto:  qreal QMatrix::m12();
extern void _ZNK7QMatrix3m12Ev(void* qthis);
  // proto:  bool QMatrix::isInvertible();
extern void _ZNK7QMatrix12isInvertibleEv(void* qthis);
  // proto:  QRect QMatrix::mapRect(const QRect & );
extern void _ZNK7QMatrix7mapRectERK5QRect(void* qthis, void* arg0);
  // proto:  void QMatrix::QMatrix(qreal m11, qreal m12, qreal m21, qreal m22, qreal dx, qreal dy);
extern void* dector_ZN7QMatrixC1Edddddd(double arg0, double arg1, double arg2, double arg3, double arg4, double arg5);
extern void _ZN7QMatrixC1Edddddd(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4, double arg5);
  // proto:  qreal QMatrix::m22();
extern void _ZNK7QMatrix3m22Ev(void* qthis);
  // proto:  QPolygon QMatrix::map(const QPolygon & a);
extern void _ZNK7QMatrix3mapERK8QPolygon(void* qthis, void* arg0);
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

// class sizeof(QMatrix)=48
type QMatrix struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  qreal QMatrix::dx();
func (this *QMatrix) dx(args ...interface{}) () {
  // dx()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QMatrix2dxEv
  default:
    qtrt.ErrorResolve("QMatrix", "dx", args)
  }

}

  // proto:  void QMatrix::QMatrix(bool );
func NewQMatrix(args ...interface{}) QMatrix {
  return QMatrix{}
}

  // proto:  qreal QMatrix::dy();
func (this *QMatrix) dy(args ...interface{}) () {
  // dy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QMatrix2dyEv
  default:
    qtrt.ErrorResolve("QMatrix", "dy", args)
  }

}

  // proto:  QMatrix & QMatrix::scale(qreal sx, qreal sy);
func (this *QMatrix) scale(args ...interface{}) () {
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
    // invoke: _ZN7QMatrix5scaleEdd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QMatrix", "scale", args)
  }

}

  // proto:  QMatrix & QMatrix::translate(qreal dx, qreal dy);
func (this *QMatrix) translate(args ...interface{}) () {
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
    // invoke: _ZN7QMatrix9translateEdd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QMatrix", "translate", args)
  }

}

  // proto:  qreal QMatrix::determinant();
func (this *QMatrix) determinant(args ...interface{}) () {
  // determinant()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QMatrix11determinantEv
  default:
    qtrt.ErrorResolve("QMatrix", "determinant", args)
  }

}

  // proto:  QMatrix & QMatrix::shear(qreal sh, qreal sv);
func (this *QMatrix) shear(args ...interface{}) () {
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
    // invoke: _ZN7QMatrix5shearEdd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QMatrix", "shear", args)
  }

}

  // proto:  qreal QMatrix::m21();
func (this *QMatrix) m21(args ...interface{}) () {
  // m21()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QMatrix3m21Ev
  default:
    qtrt.ErrorResolve("QMatrix", "m21", args)
  }

}

  // proto:  QPointF QMatrix::map(const QPointF & p);
func (this *QMatrix) map_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMatrix", "map", args)
  }

}

  // proto:  QMatrix & QMatrix::rotate(qreal a);
func (this *QMatrix) rotate(args ...interface{}) () {
  // rotate(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QMatrix6rotateEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMatrix", "rotate", args)
  }

}

  // proto:  void QMatrix::setMatrix(qreal m11, qreal m12, qreal m21, qreal m22, qreal dx, qreal dy);
func (this *QMatrix) setMatrix(args ...interface{}) () {
  // setMatrix(qreal, qreal, qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][5] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QMatrix9setMatrixEdddddd
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
    qtrt.ErrorResolve("QMatrix", "setMatrix", args)
  }

}

  // proto:  void QMatrix::reset();
func (this *QMatrix) reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QMatrix5resetEv
  default:
    qtrt.ErrorResolve("QMatrix", "reset", args)
  }

}

  // proto:  qreal QMatrix::m11();
func (this *QMatrix) m11(args ...interface{}) () {
  // m11()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QMatrix3m11Ev
  default:
    qtrt.ErrorResolve("QMatrix", "m11", args)
  }

}

  // proto:  QPolygon QMatrix::mapToPolygon(const QRect & r);
func (this *QMatrix) mapToPolygon(args ...interface{}) () {
  // mapToPolygon(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QMatrix12mapToPolygonERK5QRect
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMatrix", "mapToPolygon", args)
  }

}

  // proto:  QMatrix QMatrix::inverted(bool * invertible);
func (this *QMatrix) inverted(args ...interface{}) () {
  // inverted(_Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QMatrix8invertedEPb
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMatrix", "inverted", args)
  }

}

  // proto:  QRectF QMatrix::mapRect(const QRectF & );
func (this *QMatrix) mapRect(args ...interface{}) () {
  // mapRect(const class QRectF &)
  // mapRect(const class QRect &)
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
    // invoke: _ZNK7QMatrix7mapRectERK6QRectF
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK7QMatrix7mapRectERK5QRect
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMatrix", "mapRect", args)
  }

}

  // proto:  bool QMatrix::isIdentity();
func (this *QMatrix) isIdentity(args ...interface{}) () {
  // isIdentity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QMatrix10isIdentityEv
  default:
    qtrt.ErrorResolve("QMatrix", "isIdentity", args)
  }

}

  // proto:  qreal QMatrix::m12();
func (this *QMatrix) m12(args ...interface{}) () {
  // m12()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QMatrix3m12Ev
  default:
    qtrt.ErrorResolve("QMatrix", "m12", args)
  }

}

  // proto:  bool QMatrix::isInvertible();
func (this *QMatrix) isInvertible(args ...interface{}) () {
  // isInvertible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QMatrix12isInvertibleEv
  default:
    qtrt.ErrorResolve("QMatrix", "isInvertible", args)
  }

}

  // proto:  qreal QMatrix::m22();
func (this *QMatrix) m22(args ...interface{}) () {
  // m22()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QMatrix3m22Ev
  default:
    qtrt.ErrorResolve("QMatrix", "m22", args)
  }

}

// <= body block end

