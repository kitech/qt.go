package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QMatrix::QMatrix(const QMatrix & matrix);
extern void _ZN7QMatrixC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  void QMatrix::QMatrix(qreal m11, qreal m12, qreal m21, qreal m22, qreal dx, qreal dy);
extern void _ZN7QMatrixC2Edddddd(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4, double arg5); // 3
  // proto:  void QMatrix::QMatrix();
extern void _ZN7QMatrixC2Ev(void* qthis); // 3
  // proto:  void QMatrix::setMatrix(qreal m11, qreal m12, qreal m21, qreal m22, qreal dx, qreal dy);
extern void _ZN7QMatrix9setMatrixEdddddd(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4, double arg5); // 4
  // proto:  QRect QMatrix::mapRect(const QRect & );
extern void _ZNK7QMatrix7mapRectERK5QRect(void* qthis, void* arg0); // 4
  // proto:  QRectF QMatrix::mapRect(const QRectF & );
extern void _ZNK7QMatrix7mapRectERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  QMatrix QMatrix::inverted(bool * invertible);
extern void _ZNK7QMatrix8invertedEPb(void* qthis, bool* arg0); // 4
  // proto:  qreal QMatrix::m11();
extern void _ZNK7QMatrix3m11Ev(void* qthis); // 2
  // proto:  bool QMatrix::isInvertible();
extern void _ZNK7QMatrix12isInvertibleEv(void* qthis); // 2
  // proto:  QMatrix & QMatrix::scale(qreal sx, qreal sy);
extern void _ZN7QMatrix5scaleEdd(void* qthis, double arg0, double arg1); // 4
  // proto:  QPolygon QMatrix::mapToPolygon(const QRect & r);
extern void _ZNK7QMatrix12mapToPolygonERK5QRect(void* qthis, void* arg0); // 4
  // proto:  QMatrix & QMatrix::translate(qreal dx, qreal dy);
extern void _ZN7QMatrix9translateEdd(void* qthis, double arg0, double arg1); // 4
  // proto:  QMatrix & QMatrix::shear(qreal sh, qreal sv);
extern void _ZN7QMatrix5shearEdd(void* qthis, double arg0, double arg1); // 4
  // proto:  QPainterPath QMatrix::map(const QPainterPath & p);
extern void _ZNK7QMatrix3mapERK12QPainterPath(void* qthis, void* arg0); // 4
  // proto:  QPolygon QMatrix::map(const QPolygon & a);
extern void _ZNK7QMatrix3mapERK8QPolygon(void* qthis, void* arg0); // 4
  // proto:  QLineF QMatrix::map(const QLineF & l);
extern void _ZNK7QMatrix3mapERK6QLineF(void* qthis, void* arg0); // 4
  // proto:  void QMatrix::map(qreal x, qreal y, qreal * tx, qreal * ty);
extern void _ZNK7QMatrix3mapEddPdS0_(void* qthis, double arg0, double arg1, double* arg2, double* arg3); // 4
  // proto:  QPolygonF QMatrix::map(const QPolygonF & a);
extern void _ZNK7QMatrix3mapERK9QPolygonF(void* qthis, void* arg0); // 4
  // proto:  QPoint QMatrix::map(const QPoint & p);
extern void _ZNK7QMatrix3mapERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QLine QMatrix::map(const QLine & l);
extern void _ZNK7QMatrix3mapERK5QLine(void* qthis, void* arg0); // 4
  // proto:  void QMatrix::map(int x, int y, int * tx, int * ty);
extern void _ZNK7QMatrix3mapEiiPiS0_(void* qthis, int32_t arg0, int32_t arg1, int32_t* arg2, int32_t* arg3); // 4
  // proto:  QRegion QMatrix::map(const QRegion & r);
extern void _ZNK7QMatrix3mapERK7QRegion(void* qthis, void* arg0); // 4
  // proto:  QPointF QMatrix::map(const QPointF & p);
extern void _ZNK7QMatrix3mapERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  qreal QMatrix::determinant();
extern void _ZNK7QMatrix11determinantEv(void* qthis); // 2
  // proto:  qreal QMatrix::dx();
extern void _ZNK7QMatrix2dxEv(void* qthis); // 2
  // proto:  qreal QMatrix::dy();
extern void _ZNK7QMatrix2dyEv(void* qthis); // 2
  // proto:  void QMatrix::reset();
extern void _ZN7QMatrix5resetEv(void* qthis); // 4
  // proto:  QMatrix & QMatrix::rotate(qreal a);
extern void _ZN7QMatrix6rotateEd(void* qthis, double arg0); // 4
  // proto:  qreal QMatrix::m21();
extern void _ZNK7QMatrix3m21Ev(void* qthis); // 2
  // proto:  qreal QMatrix::m22();
extern void _ZNK7QMatrix3m22Ev(void* qthis); // 2
  // proto:  qreal QMatrix::m12();
extern void _ZNK7QMatrix3m12Ev(void* qthis); // 2
  // proto:  bool QMatrix::isIdentity();
extern void _ZNK7QMatrix10isIdentityEv(void* qthis); // 2
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// QMatrix(const class QMatrix &)
func NewQMatrix(args ...interface{}) QMatrix {
  // QMatrix(const class QMatrix &)
  // QMatrix(qreal, qreal, qreal, qreal, qreal, qreal)
  // QMatrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix{}) // "const QMatrix &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][5] = qtrt.DoubleTy(false) // "qreal"
  vtys[2] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QMatrixC1ERKS_
    // invoke: void QMatrix(const class QMatrix &)
    var arg0 = args[0].(QMatrix).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN7QMatrixC2ERKS_(qthis, arg0)
  case 1:
    // invoke: _ZN7QMatrixC1Edddddd
    // invoke: void QMatrix(qreal, qreal, qreal, qreal, qreal, qreal)
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
    C._ZN7QMatrixC2Edddddd(qthis, arg0, arg1, arg2, arg3, arg4, arg5)
  case 2:
    // invoke: _ZN7QMatrixC1Ev
    // invoke: void QMatrix()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN7QMatrixC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QMatrix", "QMatrix", args)
  }

  return QMatrix{}
}

// setMatrix(qreal, qreal, qreal, qreal, qreal, qreal)
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
    // invoke: void setMatrix(qreal, qreal, qreal, qreal, qreal, qreal)
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
    C._ZN7QMatrix9setMatrixEdddddd(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  default:
    qtrt.ErrorResolve("QMatrix", "setMatrix", args)
  }

}

// mapRect(const class QRect &)
func (this *QMatrix) mapRect(args ...interface{}) () {
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
    // invoke: _ZNK7QMatrix7mapRectERK5QRect
    // invoke: QRect mapRect(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QMatrix7mapRectERK5QRect(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK7QMatrix7mapRectERK6QRectF
    // invoke: QRectF mapRect(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QMatrix7mapRectERK6QRectF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMatrix", "mapRect", args)
  }

}

// inverted(_Bool *)
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
    // invoke: QMatrix inverted(_Bool *)
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    C._ZNK7QMatrix8invertedEPb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMatrix", "inverted", args)
  }

}

// m11()
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
    // invoke: qreal m11()
    C._ZNK7QMatrix3m11Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix", "m11", args)
  }

}

// isInvertible()
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
    // invoke: bool isInvertible()
    C._ZNK7QMatrix12isInvertibleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix", "isInvertible", args)
  }

}

// scale(qreal, qreal)
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
    // invoke: QMatrix & scale(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C._ZN7QMatrix5scaleEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMatrix", "scale", args)
  }

}

// mapToPolygon(const class QRect &)
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
    // invoke: QPolygon mapToPolygon(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QMatrix12mapToPolygonERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMatrix", "mapToPolygon", args)
  }

}

// translate(qreal, qreal)
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
    // invoke: QMatrix & translate(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C._ZN7QMatrix9translateEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMatrix", "translate", args)
  }

}

// shear(qreal, qreal)
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
    // invoke: QMatrix & shear(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C._ZN7QMatrix5shearEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMatrix", "shear", args)
  }

}

// map(const class QPainterPath &)
func (this *QMatrix) map_(args ...interface{}) () {
  // map(const class QPainterPath &)
  // map(const class QPolygon &)
  // map(const class QLineF &)
  // map(qreal, qreal, qreal *, qreal *)
  // map(const class QPolygonF &)
  // map(const class QPoint &)
  // map(const class QLine &)
  // map(int, int, int *, int *)
  // map(const class QRegion &)
  // map(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPolygon{}) // "const QPolygon &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QLineF{}) // "const QLineF &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][2] = qtrt.DoubleTy(true) // "qreal *"
  vtys[3][3] = qtrt.DoubleTy(true) // "qreal *"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QLine{}) // "const QLine &"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "int"
  vtys[7][1] = qtrt.Int32Ty(false) // "int"
  vtys[7][2] = qtrt.Int32Ty(true) // "int *"
  vtys[7][3] = qtrt.Int32Ty(true) // "int *"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QMatrix3mapERK12QPainterPath
    // invoke: QPainterPath map(const class QPainterPath &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QMatrix3mapERK12QPainterPath(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK7QMatrix3mapERK8QPolygon
    // invoke: QPolygon map(const class QPolygon &)
    var arg0 = args[0].(QPolygon).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QMatrix3mapERK8QPolygon(this.qclsinst, arg0)
  case 2:
    // invoke: _ZNK7QMatrix3mapERK6QLineF
    // invoke: QLineF map(const class QLineF &)
    var arg0 = args[0].(QLineF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QMatrix3mapERK6QLineF(this.qclsinst, arg0)
  case 3:
    // invoke: _ZNK7QMatrix3mapEddPdS0_
    // invoke: void map(qreal, qreal, qreal *, qreal *)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.double)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.double)(args[3].(*float64))
    if false {fmt.Println(arg3)}
    C._ZNK7QMatrix3mapEddPdS0_(this.qclsinst, arg0, arg1, arg2, arg3)
  case 4:
    // invoke: _ZNK7QMatrix3mapERK9QPolygonF
    // invoke: QPolygonF map(const class QPolygonF &)
    var arg0 = args[0].(QPolygonF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QMatrix3mapERK9QPolygonF(this.qclsinst, arg0)
  case 5:
    // invoke: _ZNK7QMatrix3mapERK6QPoint
    // invoke: QPoint map(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QMatrix3mapERK6QPoint(this.qclsinst, arg0)
  case 6:
    // invoke: _ZNK7QMatrix3mapERK5QLine
    // invoke: QLine map(const class QLine &)
    var arg0 = args[0].(QLine).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QMatrix3mapERK5QLine(this.qclsinst, arg0)
  case 7:
    // invoke: _ZNK7QMatrix3mapEiiPiS0_
    // invoke: void map(int, int, int *, int *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int32_t)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    C._ZNK7QMatrix3mapEiiPiS0_(this.qclsinst, arg0, arg1, arg2, arg3)
  case 8:
    // invoke: _ZNK7QMatrix3mapERK7QRegion
    // invoke: QRegion map(const class QRegion &)
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QMatrix3mapERK7QRegion(this.qclsinst, arg0)
  case 9:
    // invoke: _ZNK7QMatrix3mapERK7QPointF
    // invoke: QPointF map(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QMatrix3mapERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMatrix", "map", args)
  }

}

// determinant()
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
    // invoke: qreal determinant()
    C._ZNK7QMatrix11determinantEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix", "determinant", args)
  }

}

// dx()
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
    // invoke: qreal dx()
    C._ZNK7QMatrix2dxEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix", "dx", args)
  }

}

// dy()
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
    // invoke: qreal dy()
    C._ZNK7QMatrix2dyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix", "dy", args)
  }

}

// reset()
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
    // invoke: void reset()
    C._ZN7QMatrix5resetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix", "reset", args)
  }

}

// rotate(qreal)
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
    // invoke: QMatrix & rotate(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN7QMatrix6rotateEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMatrix", "rotate", args)
  }

}

// m21()
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
    // invoke: qreal m21()
    C._ZNK7QMatrix3m21Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix", "m21", args)
  }

}

// m22()
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
    // invoke: qreal m22()
    C._ZNK7QMatrix3m22Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix", "m22", args)
  }

}

// m12()
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
    // invoke: qreal m12()
    C._ZNK7QMatrix3m12Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix", "m12", args)
  }

}

// isIdentity()
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
    // invoke: bool isIdentity()
    C._ZNK7QMatrix10isIdentityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix", "isIdentity", args)
  }

}

// <= body block end

