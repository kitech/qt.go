package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtCore/qpoint.h
// dst-file: /src/core/qpoint.go
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
  // proto:  bool QPoint::isNull();
extern bool C_ZNK6QPoint6isNullEv(void* qthis); // 2
  // proto:  int & QPoint::rx();
extern void C_ZN6QPoint2rxEv(void* qthis); // 2
  // proto:  int & QPoint::ry();
extern void C_ZN6QPoint2ryEv(void* qthis); // 2
  // proto:  void QPoint::QPoint();
extern void* C_ZN6QPointC2Ev(); // 1
  // proto:  void QPoint::QPoint(int xpos, int ypos);
extern void* C_ZN6QPointC2Eii(int32_t arg0, int32_t arg1); // 1
  // proto:  int QPoint::y();
extern void C_ZNK6QPoint1yEv(void* qthis); // 2
  // proto: static int QPoint::dotProduct(const QPoint & p1, const QPoint & p2);
extern int32_t C_ZN6QPoint10dotProductERKS_S1_(void* arg0, void* arg1); // 2
  // proto:  int QPoint::x();
extern void C_ZNK6QPoint1xEv(void* qthis); // 2
  // proto:  void QPoint::setX(int x);
extern void C_ZN6QPoint4setXEi(void* qthis, int32_t arg0); // 2
  // proto:  void QPoint::setY(int y);
extern void C_ZN6QPoint4setYEi(void* qthis, int32_t arg0); // 2
  // proto:  int QPoint::manhattanLength();
extern int32_t C_ZNK6QPoint15manhattanLengthEv(void* qthis); // 2
  // proto:  QPoint QPointF::toPoint();
extern void* C_ZNK7QPointF7toPointEv(void* qthis); // 2
  // proto: static qreal QPointF::dotProduct(const QPointF & p1, const QPointF & p2);
extern double C_ZN7QPointF10dotProductERKS_S1_(void* arg0, void* arg1); // 2
  // proto:  qreal & QPointF::rx();
extern void C_ZN7QPointF2rxEv(void* qthis); // 2
  // proto:  qreal & QPointF::ry();
extern void C_ZN7QPointF2ryEv(void* qthis); // 2
  // proto:  bool QPointF::isNull();
extern bool C_ZNK7QPointF6isNullEv(void* qthis); // 2
  // proto:  qreal QPointF::x();
extern void C_ZNK7QPointF1xEv(void* qthis); // 2
  // proto:  qreal QPointF::y();
extern void C_ZNK7QPointF1yEv(void* qthis); // 2
  // proto:  void QPointF::QPointF(qreal xpos, qreal ypos);
extern void* C_ZN7QPointFC2Edd(double arg0, double arg1); // 1
  // proto:  void QPointF::QPointF();
extern void* C_ZN7QPointFC2Ev(); // 1
  // proto:  void QPointF::QPointF(const QPoint & p);
extern void* C_ZN7QPointFC2ERK6QPoint(void* arg0); // 1
  // proto:  void QPointF::setX(qreal x);
extern void C_ZN7QPointF4setXEd(void* qthis, double arg0); // 2
  // proto:  void QPointF::setY(qreal y);
extern void C_ZN7QPointF4setYEd(void* qthis, double arg0); // 2
  // proto:  qreal QPointF::manhattanLength();
extern double C_ZNK7QPointF15manhattanLengthEv(void* qthis); // 2
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

// class sizeof(QPoint)=8
type QPoint struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QPointF)=16
type QPointF struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// isNull()
func (this *QPoint) Isnull(args ...interface{}) (ret interface{}) {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QPoint6isNullEv
    // invoke: bool isNull()
    var ret0 = C.C_ZNK6QPoint6isNullEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPoint", "isNull", args)
  }

  return
}

// rx()
func (this *QPoint) Rx(args ...interface{}) () {
  // rx()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QPoint2rxEv
    // invoke: int & rx()
    C.C_ZN6QPoint2rxEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPoint", "rx", args)
  }

  return
}

// ry()
func (this *QPoint) Ry(args ...interface{}) () {
  // ry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QPoint2ryEv
    // invoke: int & ry()
    C.C_ZN6QPoint2ryEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPoint", "ry", args)
  }

  return
}

// QPoint()
func NewQPoint(args ...interface{}) *QPoint {
  // QPoint()
  // QPoint(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QPointC1Ev
    // invoke: void QPoint()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QPointC2Ev()
    return &QPoint{Qclsinst:qthis}
  case 1:
    // invoke: _ZN6QPointC1Eii
    // invoke: void QPoint(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QPointC2Eii(arg0, arg1)
    return &QPoint{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QPoint", "QPoint", args)
  }

  return nil // QPoint{Qclsinst:qthis}
}

// y()
func (this *QPoint) Y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QPoint1yEv
    // invoke: int y()
    C.C_ZNK6QPoint1yEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPoint", "y", args)
  }

  return
}

// dotProduct(const class QPoint &, const class QPoint &)
func (this *QPoint) Dotproduct_S(args ...interface{}) (ret interface{}) {
  // dotProduct(const class QPoint &, const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[0][1] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QPoint10dotProductERKS_S1_
    // invoke: int dotProduct(const class QPoint &, const class QPoint &)
    var arg0 = args[0].(*QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPoint).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN6QPoint10dotProductERKS_S1_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPoint", "dotProduct", args)
  }

  return
}

// x()
func (this *QPoint) X(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QPoint1xEv
    // invoke: int x()
    C.C_ZNK6QPoint1xEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPoint", "x", args)
  }

  return
}

// setX(int)
func (this *QPoint) Setx(args ...interface{}) () {
  // setX(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QPoint4setXEi
    // invoke: void setX(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QPoint4setXEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPoint", "setX", args)
  }

  return
}

// setY(int)
func (this *QPoint) Sety(args ...interface{}) () {
  // setY(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QPoint4setYEi
    // invoke: void setY(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QPoint4setYEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPoint", "setY", args)
  }

  return
}

// manhattanLength()
func (this *QPoint) Manhattanlength(args ...interface{}) (ret interface{}) {
  // manhattanLength()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QPoint15manhattanLengthEv
    // invoke: int manhattanLength()
    var ret0 = C.C_ZNK6QPoint15manhattanLengthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPoint", "manhattanLength", args)
  }

  return
}

// toPoint()
func (this *QPointF) Topoint(args ...interface{}) (ret interface{}) {
  // toPoint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPointF7toPointEv
    // invoke: QPoint toPoint()
    var ret0 = C.C_ZNK7QPointF7toPointEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPointF", "toPoint", args)
  }

  return
}

// dotProduct(const class QPointF &, const class QPointF &)
func (this *QPointF) Dotproduct_S(args ...interface{}) (ret interface{}) {
  // dotProduct(const class QPointF &, const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPointF10dotProductERKS_S1_
    // invoke: qreal dotProduct(const class QPointF &, const class QPointF &)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPointF).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QPointF10dotProductERKS_S1_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPointF", "dotProduct", args)
  }

  return
}

// rx()
func (this *QPointF) Rx(args ...interface{}) () {
  // rx()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPointF2rxEv
    // invoke: qreal & rx()
    C.C_ZN7QPointF2rxEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPointF", "rx", args)
  }

  return
}

// ry()
func (this *QPointF) Ry(args ...interface{}) () {
  // ry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPointF2ryEv
    // invoke: qreal & ry()
    C.C_ZN7QPointF2ryEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPointF", "ry", args)
  }

  return
}

// isNull()
func (this *QPointF) Isnull(args ...interface{}) (ret interface{}) {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPointF6isNullEv
    // invoke: bool isNull()
    var ret0 = C.C_ZNK7QPointF6isNullEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPointF", "isNull", args)
  }

  return
}

// x()
func (this *QPointF) X(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPointF1xEv
    // invoke: qreal x()
    C.C_ZNK7QPointF1xEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPointF", "x", args)
  }

  return
}

// y()
func (this *QPointF) Y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPointF1yEv
    // invoke: qreal y()
    C.C_ZNK7QPointF1yEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPointF", "y", args)
  }

  return
}

// QPointF(qreal, qreal)
func NewQPointF(args ...interface{}) *QPointF {
  // QPointF(qreal, qreal)
  // QPointF()
  // QPointF(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPointFC1Edd
    // invoke: void QPointF(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QPointFC2Edd(arg0, arg1)
    return &QPointF{Qclsinst:qthis}
  case 1:
    // invoke: _ZN7QPointFC1Ev
    // invoke: void QPointF()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QPointFC2Ev()
    return &QPointF{Qclsinst:qthis}
  case 2:
    // invoke: _ZN7QPointFC1ERK6QPoint
    // invoke: void QPointF(const class QPoint &)
    var arg0 = args[0].(*QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QPointFC2ERK6QPoint(arg0)
    return &QPointF{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QPointF", "QPointF", args)
  }

  return nil // QPointF{Qclsinst:qthis}
}

// setX(qreal)
func (this *QPointF) Setx(args ...interface{}) () {
  // setX(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPointF4setXEd
    // invoke: void setX(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN7QPointF4setXEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPointF", "setX", args)
  }

  return
}

// setY(qreal)
func (this *QPointF) Sety(args ...interface{}) () {
  // setY(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPointF4setYEd
    // invoke: void setY(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN7QPointF4setYEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPointF", "setY", args)
  }

  return
}

// manhattanLength()
func (this *QPointF) Manhattanlength(args ...interface{}) (ret interface{}) {
  // manhattanLength()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPointF15manhattanLengthEv
    // invoke: qreal manhattanLength()
    var ret0 = C.C_ZNK7QPointF15manhattanLengthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPointF", "manhattanLength", args)
  }

  return
}

// <= body block end

