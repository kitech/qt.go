package qtgui
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QMatrix::QMatrix(const QMatrix & other);
extern void* C_ZN7QMatrixC2ERKS_(void* arg0); // 3
  // proto:  void QMatrix::QMatrix(qreal m11, qreal m12, qreal m21, qreal m22, qreal dx, qreal dy);
extern void* C_ZN7QMatrixC2Edddddd(double arg0, double arg1, double arg2, double arg3, double arg4, double arg5); // 3
  // proto:  void QMatrix::QMatrix();
extern void* C_ZN7QMatrixC2Ev(); // 3
  // proto:  void QMatrix::setMatrix(qreal m11, qreal m12, qreal m21, qreal m22, qreal dx, qreal dy);
extern void C_ZN7QMatrix9setMatrixEdddddd(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4, double arg5); // 4
  // proto:  QRect QMatrix::mapRect(const QRect & );
extern void* C_ZNK7QMatrix7mapRectERK5QRect(void* qthis, void* arg0); // 4
  // proto:  QRectF QMatrix::mapRect(const QRectF & );
extern void* C_ZNK7QMatrix7mapRectERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  QMatrix QMatrix::inverted(bool * invertible);
extern void* C_ZNK7QMatrix8invertedEPb(void* qthis, void* arg0); // 4
  // proto:  qreal QMatrix::m11();
extern double C_ZNK7QMatrix3m11Ev(void* qthis); // 2
  // proto:  bool QMatrix::isInvertible();
extern bool C_ZNK7QMatrix12isInvertibleEv(void* qthis); // 2
  // proto:  QMatrix & QMatrix::scale(qreal sx, qreal sy);
extern void* C_ZN7QMatrix5scaleEdd(void* qthis, double arg0, double arg1); // 4
  // proto:  QPolygon QMatrix::mapToPolygon(const QRect & r);
extern void* C_ZNK7QMatrix12mapToPolygonERK5QRect(void* qthis, void* arg0); // 4
  // proto:  QMatrix & QMatrix::translate(qreal dx, qreal dy);
extern void* C_ZN7QMatrix9translateEdd(void* qthis, double arg0, double arg1); // 4
  // proto:  QMatrix & QMatrix::shear(qreal sh, qreal sv);
extern void* C_ZN7QMatrix5shearEdd(void* qthis, double arg0, double arg1); // 4
  // proto:  QPainterPath QMatrix::map(const QPainterPath & p);
extern void* C_ZNK7QMatrix3mapERK12QPainterPath(void* qthis, void* arg0); // 4
  // proto:  QPolygon QMatrix::map(const QPolygon & a);
extern void* C_ZNK7QMatrix3mapERK8QPolygon(void* qthis, void* arg0); // 4
  // proto:  QLineF QMatrix::map(const QLineF & l);
extern void* C_ZNK7QMatrix3mapERK6QLineF(void* qthis, void* arg0); // 4
  // proto:  void QMatrix::map(qreal x, qreal y, qreal * tx, qreal * ty);
extern void C_ZNK7QMatrix3mapEddPdS0_(void* qthis, double arg0, double arg1, void* arg2, void* arg3); // 4
  // proto:  QPolygonF QMatrix::map(const QPolygonF & a);
extern void* C_ZNK7QMatrix3mapERK9QPolygonF(void* qthis, void* arg0); // 4
  // proto:  QPoint QMatrix::map(const QPoint & p);
extern void* C_ZNK7QMatrix3mapERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QLine QMatrix::map(const QLine & l);
extern void* C_ZNK7QMatrix3mapERK5QLine(void* qthis, void* arg0); // 4
  // proto:  void QMatrix::map(int x, int y, int * tx, int * ty);
extern void C_ZNK7QMatrix3mapEiiPiS0_(void* qthis, int32_t arg0, int32_t arg1, void* arg2, void* arg3); // 4
  // proto:  QRegion QMatrix::map(const QRegion & r);
extern void* C_ZNK7QMatrix3mapERK7QRegion(void* qthis, void* arg0); // 4
  // proto:  QPointF QMatrix::map(const QPointF & p);
extern void* C_ZNK7QMatrix3mapERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  qreal QMatrix::determinant();
extern double C_ZNK7QMatrix11determinantEv(void* qthis); // 2
  // proto:  qreal QMatrix::dx();
extern double C_ZNK7QMatrix2dxEv(void* qthis); // 2
  // proto:  qreal QMatrix::dy();
extern double C_ZNK7QMatrix2dyEv(void* qthis); // 2
  // proto:  void QMatrix::reset();
extern void C_ZN7QMatrix5resetEv(void* qthis); // 4
  // proto:  QMatrix & QMatrix::rotate(qreal a);
extern void* C_ZN7QMatrix6rotateEd(void* qthis, double arg0); // 4
  // proto:  qreal QMatrix::m21();
extern double C_ZNK7QMatrix3m21Ev(void* qthis); // 2
  // proto:  qreal QMatrix::m22();
extern double C_ZNK7QMatrix3m22Ev(void* qthis); // 2
  // proto:  qreal QMatrix::m12();
extern double C_ZNK7QMatrix3m12Ev(void* qthis); // 2
  // proto:  bool QMatrix::isIdentity();
extern bool C_ZNK7QMatrix10isIdentityEv(void* qthis); // 2
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

// class sizeof(QMatrix)=48
type QMatrix struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// QMatrix(const class QMatrix &)
func NewQMatrix(args ...interface{}) *QMatrix {
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
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QMatrixC1ERKS_
    // invoke: void QMatrix(const class QMatrix &)
    var arg0 = args[0].(*QMatrix).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QMatrixC2ERKS_(arg0)
    return &QMatrix{Qclsinst:qthis}
  case 1:
    // invoke: _ZN7QMatrixC1Edddddd
    // invoke: void QMatrix(qreal, qreal, qreal, qreal, qreal, qreal)
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
    qthis = C.C_ZN7QMatrixC2Edddddd(arg0, arg1, arg2, arg3, arg4, arg5)
    return &QMatrix{Qclsinst:qthis}
  case 2:
    // invoke: _ZN7QMatrixC1Ev
    // invoke: void QMatrix()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QMatrixC2Ev()
    return &QMatrix{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QMatrix", "QMatrix", args)
  }

  return nil // QMatrix{Qclsinst:qthis}
}

// setMatrix(qreal, qreal, qreal, qreal, qreal, qreal)
func (this *QMatrix) Setmatrix(args ...interface{}) () {
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
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QMatrix9setMatrixEdddddd
    // invoke: void setMatrix(qreal, qreal, qreal, qreal, qreal, qreal)
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
    C.C_ZN7QMatrix9setMatrixEdddddd(this.Qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  default:
    qtrt.ErrorResolve("QMatrix", "setMatrix", args)
  }

  return
}

// mapRect(const class QRect &)
func (this *QMatrix) Maprect(args ...interface{}) (ret interface{}) {
  // mapRect(const class QRect &)
  // mapRect(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
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
    // invoke: _ZNK7QMatrix7mapRectERK5QRect
    // invoke: QRect mapRect(const class QRect &)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QMatrix7mapRectERK5QRect(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK7QMatrix7mapRectERK6QRectF
    // invoke: QRectF mapRect(const class QRectF &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QMatrix7mapRectERK6QRectF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMatrix", "mapRect", args)
  }

  return
}

// inverted(_Bool *)
func (this *QMatrix) Inverted(args ...interface{}) (ret interface{}) {
  // inverted(_Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QMatrix8invertedEPb
    // invoke: QMatrix inverted(_Bool *)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QMatrix8invertedEPb(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMatrix{}) // "QMatrix"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMatrix", "inverted", args)
  }

  return
}

// m11()
func (this *QMatrix) M11(args ...interface{}) (ret interface{}) {
  // m11()
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
    // invoke: _ZNK7QMatrix3m11Ev
    // invoke: qreal m11()
    var ret0 = C.C_ZNK7QMatrix3m11Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMatrix", "m11", args)
  }

  return
}

// isInvertible()
func (this *QMatrix) Isinvertible(args ...interface{}) (ret interface{}) {
  // isInvertible()
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
    // invoke: _ZNK7QMatrix12isInvertibleEv
    // invoke: bool isInvertible()
    var ret0 = C.C_ZNK7QMatrix12isInvertibleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMatrix", "isInvertible", args)
  }

  return
}

// scale(qreal, qreal)
func (this *QMatrix) Scale(args ...interface{}) (ret interface{}) {
  // scale(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QMatrix5scaleEdd
    // invoke: QMatrix & scale(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QMatrix5scaleEdd(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMatrix{}) // "QMatrix &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMatrix", "scale", args)
  }

  return
}

// mapToPolygon(const class QRect &)
func (this *QMatrix) Maptopolygon(args ...interface{}) (ret interface{}) {
  // mapToPolygon(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QMatrix12mapToPolygonERK5QRect
    // invoke: QPolygon mapToPolygon(const class QRect &)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QMatrix12mapToPolygonERK5QRect(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygon{}) // "QPolygon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMatrix", "mapToPolygon", args)
  }

  return
}

// translate(qreal, qreal)
func (this *QMatrix) Translate(args ...interface{}) (ret interface{}) {
  // translate(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QMatrix9translateEdd
    // invoke: QMatrix & translate(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QMatrix9translateEdd(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMatrix{}) // "QMatrix &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMatrix", "translate", args)
  }

  return
}

// shear(qreal, qreal)
func (this *QMatrix) Shear(args ...interface{}) (ret interface{}) {
  // shear(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QMatrix5shearEdd
    // invoke: QMatrix & shear(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QMatrix5shearEdd(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMatrix{}) // "QMatrix &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMatrix", "shear", args)
  }

  return
}

// map(const class QPainterPath &)
func (this *QMatrix) Map_(args ...interface{}) (ret interface{}) {
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
  vtys[2][0] = reflect.TypeOf(qtcore.QLineF{}) // "const QLineF &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][2] = qtrt.DoubleTy(true) // "qreal *"
  vtys[3][3] = qtrt.DoubleTy(true) // "qreal *"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(qtcore.QLine{}) // "const QLine &"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "int"
  vtys[7][1] = qtrt.Int32Ty(false) // "int"
  vtys[7][2] = qtrt.Int32Ty(true) // "int *"
  vtys[7][3] = qtrt.Int32Ty(true) // "int *"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QMatrix3mapERK12QPainterPath
    // invoke: QPainterPath map(const class QPainterPath &)
    var arg0 = args[0].(*QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QMatrix3mapERK12QPainterPath(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK7QMatrix3mapERK8QPolygon
    // invoke: QPolygon map(const class QPolygon &)
    var arg0 = args[0].(*QPolygon).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QMatrix3mapERK8QPolygon(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygon{}) // "QPolygon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK7QMatrix3mapERK6QLineF
    // invoke: QLineF map(const class QLineF &)
    var arg0 = args[0].(*qtcore.QLineF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QMatrix3mapERK6QLineF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QLineF{}) // "QLineF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZNK7QMatrix3mapEddPdS0_
    // invoke: void map(qreal, qreal, qreal *, qreal *)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*float64))
    if false {fmt.Println(arg3)}
    C.C_ZNK7QMatrix3mapEddPdS0_(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 4:
    // invoke: _ZNK7QMatrix3mapERK9QPolygonF
    // invoke: QPolygonF map(const class QPolygonF &)
    var arg0 = args[0].(*QPolygonF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QMatrix3mapERK9QPolygonF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 5:
    // invoke: _ZNK7QMatrix3mapERK6QPoint
    // invoke: QPoint map(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QMatrix3mapERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 6:
    // invoke: _ZNK7QMatrix3mapERK5QLine
    // invoke: QLine map(const class QLine &)
    var arg0 = args[0].(*qtcore.QLine).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QMatrix3mapERK5QLine(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QLine{}) // "QLine"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 7:
    // invoke: _ZNK7QMatrix3mapEiiPiS0_
    // invoke: void map(int, int, int *, int *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    C.C_ZNK7QMatrix3mapEiiPiS0_(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 8:
    // invoke: _ZNK7QMatrix3mapERK7QRegion
    // invoke: QRegion map(const class QRegion &)
    var arg0 = args[0].(*QRegion).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QMatrix3mapERK7QRegion(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRegion{}) // "QRegion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 9:
    // invoke: _ZNK7QMatrix3mapERK7QPointF
    // invoke: QPointF map(const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QMatrix3mapERK7QPointF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMatrix", "map", args)
  }

  return
}

// determinant()
func (this *QMatrix) Determinant(args ...interface{}) (ret interface{}) {
  // determinant()
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
    // invoke: _ZNK7QMatrix11determinantEv
    // invoke: qreal determinant()
    var ret0 = C.C_ZNK7QMatrix11determinantEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMatrix", "determinant", args)
  }

  return
}

// dx()
func (this *QMatrix) Dx(args ...interface{}) (ret interface{}) {
  // dx()
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
    // invoke: _ZNK7QMatrix2dxEv
    // invoke: qreal dx()
    var ret0 = C.C_ZNK7QMatrix2dxEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMatrix", "dx", args)
  }

  return
}

// dy()
func (this *QMatrix) Dy(args ...interface{}) (ret interface{}) {
  // dy()
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
    // invoke: _ZNK7QMatrix2dyEv
    // invoke: qreal dy()
    var ret0 = C.C_ZNK7QMatrix2dyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMatrix", "dy", args)
  }

  return
}

// reset()
func (this *QMatrix) Reset(args ...interface{}) () {
  // reset()
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
    // invoke: _ZN7QMatrix5resetEv
    // invoke: void reset()
    C.C_ZN7QMatrix5resetEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix", "reset", args)
  }

  return
}

// rotate(qreal)
func (this *QMatrix) Rotate(args ...interface{}) (ret interface{}) {
  // rotate(qreal)
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
    // invoke: _ZN7QMatrix6rotateEd
    // invoke: QMatrix & rotate(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN7QMatrix6rotateEd(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMatrix{}) // "QMatrix &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMatrix", "rotate", args)
  }

  return
}

// m21()
func (this *QMatrix) M21(args ...interface{}) (ret interface{}) {
  // m21()
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
    // invoke: _ZNK7QMatrix3m21Ev
    // invoke: qreal m21()
    var ret0 = C.C_ZNK7QMatrix3m21Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMatrix", "m21", args)
  }

  return
}

// m22()
func (this *QMatrix) M22(args ...interface{}) (ret interface{}) {
  // m22()
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
    // invoke: _ZNK7QMatrix3m22Ev
    // invoke: qreal m22()
    var ret0 = C.C_ZNK7QMatrix3m22Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMatrix", "m22", args)
  }

  return
}

// m12()
func (this *QMatrix) M12(args ...interface{}) (ret interface{}) {
  // m12()
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
    // invoke: _ZNK7QMatrix3m12Ev
    // invoke: qreal m12()
    var ret0 = C.C_ZNK7QMatrix3m12Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMatrix", "m12", args)
  }

  return
}

// isIdentity()
func (this *QMatrix) Isidentity(args ...interface{}) (ret interface{}) {
  // isIdentity()
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
    // invoke: _ZNK7QMatrix10isIdentityEv
    // invoke: bool isIdentity()
    var ret0 = C.C_ZNK7QMatrix10isIdentityEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMatrix", "isIdentity", args)
  }

  return
}

// <= body block end

