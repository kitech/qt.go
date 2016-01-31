package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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
extern void C_ZNK6QPoint6isNullEv(void* qthis); // 2
  // proto:  int & QPoint::rx();
extern void C_ZN6QPoint2rxEv(void* qthis); // 2
  // proto:  int & QPoint::ry();
extern void C_ZN6QPoint2ryEv(void* qthis); // 2
  // proto:  void QPoint::QPoint();
extern void C_ZN6QPointC2Ev(void* qthis); // 1
  // proto:  void QPoint::QPoint(int xpos, int ypos);
extern void C_ZN6QPointC2Eii(void* qthis, int32_t arg0, int32_t arg1); // 1
  // proto:  int QPoint::y();
extern void C_ZNK6QPoint1yEv(void* qthis); // 2
  // proto: static int QPoint::dotProduct(const QPoint & p1, const QPoint & p2);
extern void C_ZN6QPoint10dotProductERKS_S1_(void* arg0, void* arg1); // 2
  // proto:  int QPoint::x();
extern void C_ZNK6QPoint1xEv(void* qthis); // 2
  // proto:  void QPoint::setX(int x);
extern void C_ZN6QPoint4setXEi(void* qthis, int32_t arg0); // 2
  // proto:  void QPoint::setY(int y);
extern void C_ZN6QPoint4setYEi(void* qthis, int32_t arg0); // 2
  // proto:  int QPoint::manhattanLength();
extern void C_ZNK6QPoint15manhattanLengthEv(void* qthis); // 2
  // proto:  QPoint QPointF::toPoint();
extern void C_ZNK7QPointF7toPointEv(void* qthis); // 2
  // proto: static qreal QPointF::dotProduct(const QPointF & p1, const QPointF & p2);
extern void C_ZN7QPointF10dotProductERKS_S1_(void* arg0, void* arg1); // 2
  // proto:  qreal & QPointF::rx();
extern void C_ZN7QPointF2rxEv(void* qthis); // 2
  // proto:  qreal & QPointF::ry();
extern void C_ZN7QPointF2ryEv(void* qthis); // 2
  // proto:  bool QPointF::isNull();
extern void C_ZNK7QPointF6isNullEv(void* qthis); // 2
  // proto:  qreal QPointF::x();
extern void C_ZNK7QPointF1xEv(void* qthis); // 2
  // proto:  qreal QPointF::y();
extern void C_ZNK7QPointF1yEv(void* qthis); // 2
  // proto:  void QPointF::QPointF(qreal xpos, qreal ypos);
extern void C_ZN7QPointFC2Edd(void* qthis, double arg0, double arg1); // 1
  // proto:  void QPointF::QPointF();
extern void C_ZN7QPointFC2Ev(void* qthis); // 1
  // proto:  void QPointF::QPointF(const QPoint & p);
extern void C_ZN7QPointFC2ERK6QPoint(void* qthis, void* arg0); // 1
  // proto:  void QPointF::setX(qreal x);
extern void C_ZN7QPointF4setXEd(void* qthis, double arg0); // 2
  // proto:  void QPointF::setY(qreal y);
extern void C_ZN7QPointF4setYEd(void* qthis, double arg0); // 2
  // proto:  qreal QPointF::manhattanLength();
extern void C_ZNK7QPointF15manhattanLengthEv(void* qthis); // 2
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QPointF)=16
type QPointF struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// isNull()
func (this *QPoint) isNull(args ...interface{}) () {
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
    C.C_ZNK6QPoint6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPoint", "isNull", args)
  }

}

// rx()
func (this *QPoint) rx(args ...interface{}) () {
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
    C.C_ZN6QPoint2rxEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPoint", "rx", args)
  }

}

// ry()
func (this *QPoint) ry(args ...interface{}) () {
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
    C.C_ZN6QPoint2ryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPoint", "ry", args)
  }

}

// QPoint()
func NewQPoint(args ...interface{}) QPoint {
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
    C.C_ZN6QPointC2Ev(qthis)
  case 1:
    // invoke: _ZN6QPointC1Eii
    // invoke: void QPoint(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN6QPointC2Eii(qthis, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPoint", "QPoint", args)
  }

  return QPoint{}
}

// y()
func (this *QPoint) y(args ...interface{}) () {
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
    C.C_ZNK6QPoint1yEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPoint", "y", args)
  }

}

// dotProduct(const class QPoint &, const class QPoint &)
func (this *QPoint) dotProduct_s(args ...interface{}) () {
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
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPoint).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN6QPoint10dotProductERKS_S1_(arg0, arg1)
  default:
    qtrt.ErrorResolve("QPoint", "dotProduct", args)
  }

}

// x()
func (this *QPoint) x(args ...interface{}) () {
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
    C.C_ZNK6QPoint1xEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPoint", "x", args)
  }

}

// setX(int)
func (this *QPoint) setX(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QPoint4setXEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPoint", "setX", args)
  }

}

// setY(int)
func (this *QPoint) setY(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QPoint4setYEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPoint", "setY", args)
  }

}

// manhattanLength()
func (this *QPoint) manhattanLength(args ...interface{}) () {
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
    C.C_ZNK6QPoint15manhattanLengthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPoint", "manhattanLength", args)
  }

}

// toPoint()
func (this *QPointF) toPoint(args ...interface{}) () {
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
    C.C_ZNK7QPointF7toPointEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPointF", "toPoint", args)
  }

}

// dotProduct(const class QPointF &, const class QPointF &)
func (this *QPointF) dotProduct_s(args ...interface{}) () {
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
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPointF).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN7QPointF10dotProductERKS_S1_(arg0, arg1)
  default:
    qtrt.ErrorResolve("QPointF", "dotProduct", args)
  }

}

// rx()
func (this *QPointF) rx(args ...interface{}) () {
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
    C.C_ZN7QPointF2rxEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPointF", "rx", args)
  }

}

// ry()
func (this *QPointF) ry(args ...interface{}) () {
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
    C.C_ZN7QPointF2ryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPointF", "ry", args)
  }

}

// isNull()
func (this *QPointF) isNull(args ...interface{}) () {
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
    C.C_ZNK7QPointF6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPointF", "isNull", args)
  }

}

// x()
func (this *QPointF) x(args ...interface{}) () {
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
    C.C_ZNK7QPointF1xEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPointF", "x", args)
  }

}

// y()
func (this *QPointF) y(args ...interface{}) () {
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
    C.C_ZNK7QPointF1yEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPointF", "y", args)
  }

}

// QPointF(qreal, qreal)
func NewQPointF(args ...interface{}) QPointF {
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN7QPointFC2Edd(qthis, arg0, arg1)
  case 1:
    // invoke: _ZN7QPointFC1Ev
    // invoke: void QPointF()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN7QPointFC2Ev(qthis)
  case 2:
    // invoke: _ZN7QPointFC1ERK6QPoint
    // invoke: void QPointF(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN7QPointFC2ERK6QPoint(qthis, arg0)
  default:
    qtrt.ErrorResolve("QPointF", "QPointF", args)
  }

  return QPointF{}
}

// setX(qreal)
func (this *QPointF) setX(args ...interface{}) () {
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN7QPointF4setXEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPointF", "setX", args)
  }

}

// setY(qreal)
func (this *QPointF) setY(args ...interface{}) () {
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN7QPointF4setYEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPointF", "setY", args)
  }

}

// manhattanLength()
func (this *QPointF) manhattanLength(args ...interface{}) () {
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
    C.C_ZNK7QPointF15manhattanLengthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPointF", "manhattanLength", args)
  }

}

// <= body block end

