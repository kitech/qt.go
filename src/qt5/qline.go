package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtCore/qline.h
// dst-file: /src/core/qline.go
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
  // proto:  QPoint QLine::p2();
extern void C_ZNK5QLine2p2Ev(void* qthis); // 2
  // proto:  int QLine::y2();
extern void C_ZNK5QLine2y2Ev(void* qthis); // 2
  // proto:  int QLine::dx();
extern void C_ZNK5QLine2dxEv(void* qthis); // 2
  // proto:  int QLine::y1();
extern void C_ZNK5QLine2y1Ev(void* qthis); // 2
  // proto:  void QLine::setP2(const QPoint & p2);
extern void C_ZN5QLine5setP2ERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  int QLine::x1();
extern void C_ZNK5QLine2x1Ev(void* qthis); // 2
  // proto:  void QLine::setLine(int x1, int y1, int x2, int y2);
extern void C_ZN5QLine7setLineEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 2
  // proto:  int QLine::x2();
extern void C_ZNK5QLine2x2Ev(void* qthis); // 2
  // proto:  bool QLine::isNull();
extern void C_ZNK5QLine6isNullEv(void* qthis); // 2
  // proto:  QPoint QLine::p1();
extern void C_ZNK5QLine2p1Ev(void* qthis); // 2
  // proto:  void QLine::QLine(int x1, int y1, int x2, int y2);
extern void C_ZN5QLineC2Eiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 1
  // proto:  void QLine::QLine(const QPoint & pt1, const QPoint & pt2);
extern void C_ZN5QLineC2ERK6QPointS2_(void* qthis, void* arg0, void* arg1); // 1
  // proto:  void QLine::QLine();
extern void C_ZN5QLineC2Ev(void* qthis); // 1
  // proto:  int QLine::dy();
extern void C_ZNK5QLine2dyEv(void* qthis); // 2
  // proto:  QLine QLine::translated(const QPoint & p);
extern void C_ZNK5QLine10translatedERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  QLine QLine::translated(int dx, int dy);
extern void C_ZNK5QLine10translatedEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  void QLine::setP1(const QPoint & p1);
extern void C_ZN5QLine5setP1ERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  void QLine::translate(const QPoint & p);
extern void C_ZN5QLine9translateERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  void QLine::translate(int dx, int dy);
extern void C_ZN5QLine9translateEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  void QLine::setPoints(const QPoint & p1, const QPoint & p2);
extern void C_ZN5QLine9setPointsERK6QPointS2_(void* qthis, void* arg0, void* arg1); // 2
  // proto:  void QLineF::setAngle(qreal angle);
extern void C_ZN6QLineF8setAngleEd(void* qthis, double arg0); // 4
  // proto: static QLineF QLineF::fromPolar(qreal length, qreal angle);
extern void C_ZN6QLineF9fromPolarEdd(double arg0, double arg1); // 4
  // proto:  qreal QLineF::y1();
extern void C_ZNK6QLineF2y1Ev(void* qthis); // 2
  // proto:  void QLineF::setP1(const QPointF & p1);
extern void C_ZN6QLineF5setP1ERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  qreal QLineF::y2();
extern void C_ZNK6QLineF2y2Ev(void* qthis); // 2
  // proto:  QLineF QLineF::normalVector();
extern void C_ZNK6QLineF12normalVectorEv(void* qthis); // 2
  // proto:  qreal QLineF::angle(const QLineF & l);
extern void C_ZNK6QLineF5angleERKS_(void* qthis, void* arg0); // 4
  // proto:  qreal QLineF::angle();
extern void C_ZNK6QLineF5angleEv(void* qthis); // 4
  // proto:  QLineF::IntersectType QLineF::intersect(const QLineF & l, QPointF * intersectionPoint);
extern void C_ZNK6QLineF9intersectERKS_P7QPointF(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QLineF::setLine(qreal x1, qreal y1, qreal x2, qreal y2);
extern void C_ZN6QLineF7setLineEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  void QLineF::translate(qreal dx, qreal dy);
extern void C_ZN6QLineF9translateEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QLineF::translate(const QPointF & p);
extern void C_ZN6QLineF9translateERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  void QLineF::setP2(const QPointF & p2);
extern void C_ZN6QLineF5setP2ERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  QLineF QLineF::unitVector();
extern void C_ZNK6QLineF10unitVectorEv(void* qthis); // 4
  // proto:  qreal QLineF::angleTo(const QLineF & l);
extern void C_ZNK6QLineF7angleToERKS_(void* qthis, void* arg0); // 4
  // proto:  qreal QLineF::dx();
extern void C_ZNK6QLineF2dxEv(void* qthis); // 2
  // proto:  qreal QLineF::dy();
extern void C_ZNK6QLineF2dyEv(void* qthis); // 2
  // proto:  qreal QLineF::x2();
extern void C_ZNK6QLineF2x2Ev(void* qthis); // 2
  // proto:  qreal QLineF::x1();
extern void C_ZNK6QLineF2x1Ev(void* qthis); // 2
  // proto:  void QLineF::setPoints(const QPointF & p1, const QPointF & p2);
extern void C_ZN6QLineF9setPointsERK7QPointFS2_(void* qthis, void* arg0, void* arg1); // 2
  // proto:  void QLineF::setLength(qreal len);
extern void C_ZN6QLineF9setLengthEd(void* qthis, double arg0); // 2
  // proto:  QPointF QLineF::p2();
extern void C_ZNK6QLineF2p2Ev(void* qthis); // 2
  // proto:  QPointF QLineF::p1();
extern void C_ZNK6QLineF2p1Ev(void* qthis); // 2
  // proto:  QPointF QLineF::pointAt(qreal t);
extern void C_ZNK6QLineF7pointAtEd(void* qthis, double arg0); // 2
  // proto:  QLine QLineF::toLine();
extern void C_ZNK6QLineF6toLineEv(void* qthis); // 2
  // proto:  bool QLineF::isNull();
extern void C_ZNK6QLineF6isNullEv(void* qthis); // 2
  // proto:  qreal QLineF::length();
extern void C_ZNK6QLineF6lengthEv(void* qthis); // 4
  // proto:  QLineF QLineF::translated(const QPointF & p);
extern void C_ZNK6QLineF10translatedERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  QLineF QLineF::translated(qreal dx, qreal dy);
extern void C_ZNK6QLineF10translatedEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QLineF::QLineF(qreal x1, qreal y1, qreal x2, qreal y2);
extern void C_ZN6QLineFC2Edddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 1
  // proto:  void QLineF::QLineF(const QLine & line);
extern void C_ZN6QLineFC2ERK5QLine(void* qthis, void* arg0); // 1
  // proto:  void QLineF::QLineF(const QPointF & pt1, const QPointF & pt2);
extern void C_ZN6QLineFC2ERK7QPointFS2_(void* qthis, void* arg0, void* arg1); // 1
  // proto:  void QLineF::QLineF();
extern void C_ZN6QLineFC2Ev(void* qthis); // 1
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

// class sizeof(QLine)=16
type QLine struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QLineF)=32
type QLineF struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// p2()
func (this *QLine) p2(args ...interface{}) () {
  // p2()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QLine2p2Ev
    // invoke: QPoint p2()
    C.C_ZNK5QLine2p2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLine", "p2", args)
  }

}

// y2()
func (this *QLine) y2(args ...interface{}) () {
  // y2()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QLine2y2Ev
    // invoke: int y2()
    C.C_ZNK5QLine2y2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLine", "y2", args)
  }

}

// dx()
func (this *QLine) dx(args ...interface{}) () {
  // dx()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QLine2dxEv
    // invoke: int dx()
    C.C_ZNK5QLine2dxEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLine", "dx", args)
  }

}

// y1()
func (this *QLine) y1(args ...interface{}) () {
  // y1()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QLine2y1Ev
    // invoke: int y1()
    C.C_ZNK5QLine2y1Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLine", "y1", args)
  }

}

// setP2(const class QPoint &)
func (this *QLine) setP2(args ...interface{}) () {
  // setP2(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QLine5setP2ERK6QPoint
    // invoke: void setP2(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QLine5setP2ERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLine", "setP2", args)
  }

}

// x1()
func (this *QLine) x1(args ...interface{}) () {
  // x1()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QLine2x1Ev
    // invoke: int x1()
    C.C_ZNK5QLine2x1Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLine", "x1", args)
  }

}

// setLine(int, int, int, int)
func (this *QLine) setLine(args ...interface{}) () {
  // setLine(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QLine7setLineEiiii
    // invoke: void setLine(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN5QLine7setLineEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QLine", "setLine", args)
  }

}

// x2()
func (this *QLine) x2(args ...interface{}) () {
  // x2()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QLine2x2Ev
    // invoke: int x2()
    C.C_ZNK5QLine2x2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLine", "x2", args)
  }

}

// isNull()
func (this *QLine) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QLine6isNullEv
    // invoke: bool isNull()
    C.C_ZNK5QLine6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLine", "isNull", args)
  }

}

// p1()
func (this *QLine) p1(args ...interface{}) () {
  // p1()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QLine2p1Ev
    // invoke: QPoint p1()
    C.C_ZNK5QLine2p1Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLine", "p1", args)
  }

}

// QLine(int, int, int, int)
func NewQLine(args ...interface{}) QLine {
  // QLine(int, int, int, int)
  // QLine(const class QPoint &, const class QPoint &)
  // QLine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1][1] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[2] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QLineC1Eiiii
    // invoke: void QLine(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN5QLineC2Eiiii(qthis, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN5QLineC1ERK6QPointS2_
    // invoke: void QLine(const class QPoint &, const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPoint).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN5QLineC2ERK6QPointS2_(qthis, arg0, arg1)
  case 2:
    // invoke: _ZN5QLineC1Ev
    // invoke: void QLine()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN5QLineC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QLine", "QLine", args)
  }

  return QLine{}
}

// dy()
func (this *QLine) dy(args ...interface{}) () {
  // dy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QLine2dyEv
    // invoke: int dy()
    C.C_ZNK5QLine2dyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLine", "dy", args)
  }

}

// translated(const class QPoint &)
func (this *QLine) translated(args ...interface{}) () {
  // translated(const class QPoint &)
  // translated(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QLine10translatedERK6QPoint
    // invoke: QLine translated(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK5QLine10translatedERK6QPoint(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK5QLine10translatedEii
    // invoke: QLine translated(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK5QLine10translatedEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLine", "translated", args)
  }

}

// setP1(const class QPoint &)
func (this *QLine) setP1(args ...interface{}) () {
  // setP1(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QLine5setP1ERK6QPoint
    // invoke: void setP1(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QLine5setP1ERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLine", "setP1", args)
  }

}

// translate(const class QPoint &)
func (this *QLine) translate(args ...interface{}) () {
  // translate(const class QPoint &)
  // translate(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QLine9translateERK6QPoint
    // invoke: void translate(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QLine9translateERK6QPoint(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN5QLine9translateEii
    // invoke: void translate(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN5QLine9translateEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLine", "translate", args)
  }

}

// setPoints(const class QPoint &, const class QPoint &)
func (this *QLine) setPoints(args ...interface{}) () {
  // setPoints(const class QPoint &, const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[0][1] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QLine9setPointsERK6QPointS2_
    // invoke: void setPoints(const class QPoint &, const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPoint).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN5QLine9setPointsERK6QPointS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLine", "setPoints", args)
  }

}

// setAngle(qreal)
func (this *QLineF) setAngle(args ...interface{}) () {
  // setAngle(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLineF8setAngleEd
    // invoke: void setAngle(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN6QLineF8setAngleEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineF", "setAngle", args)
  }

}

// fromPolar(qreal, qreal)
func (this *QLineF) fromPolar_s(args ...interface{}) () {
  // fromPolar(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLineF9fromPolarEdd
    // invoke: QLineF fromPolar(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN6QLineF9fromPolarEdd(arg0, arg1)
  default:
    qtrt.ErrorResolve("QLineF", "fromPolar", args)
  }

}

// y1()
func (this *QLineF) y1(args ...interface{}) () {
  // y1()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF2y1Ev
    // invoke: qreal y1()
    C.C_ZNK6QLineF2y1Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineF", "y1", args)
  }

}

// setP1(const class QPointF &)
func (this *QLineF) setP1(args ...interface{}) () {
  // setP1(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLineF5setP1ERK7QPointF
    // invoke: void setP1(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QLineF5setP1ERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineF", "setP1", args)
  }

}

// y2()
func (this *QLineF) y2(args ...interface{}) () {
  // y2()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF2y2Ev
    // invoke: qreal y2()
    C.C_ZNK6QLineF2y2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineF", "y2", args)
  }

}

// normalVector()
func (this *QLineF) normalVector(args ...interface{}) () {
  // normalVector()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF12normalVectorEv
    // invoke: QLineF normalVector()
    C.C_ZNK6QLineF12normalVectorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineF", "normalVector", args)
  }

}

// angle(const class QLineF &)
func (this *QLineF) angle(args ...interface{}) () {
  // angle(const class QLineF &)
  // angle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLineF{}) // "const QLineF &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF5angleERKS_
    // invoke: qreal angle(const class QLineF &)
    var arg0 = args[0].(QLineF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK6QLineF5angleERKS_(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK6QLineF5angleEv
    // invoke: qreal angle()
    C.C_ZNK6QLineF5angleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineF", "angle", args)
  }

}

// intersect(const class QLineF &, class QPointF *)
func (this *QLineF) intersect(args ...interface{}) () {
  // intersect(const class QLineF &, class QPointF *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLineF{}) // "const QLineF &"
  vtys[0][1] = reflect.TypeOf(QPointF{}) // "QPointF *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF9intersectERKS_P7QPointF
    // invoke: QLineF::IntersectType intersect(const class QLineF &, class QPointF *)
    var arg0 = args[0].(QLineF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPointF).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZNK6QLineF9intersectERKS_P7QPointF(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLineF", "intersect", args)
  }

}

// setLine(qreal, qreal, qreal, qreal)
func (this *QLineF) setLine(args ...interface{}) () {
  // setLine(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLineF7setLineEdddd
    // invoke: void setLine(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN6QLineF7setLineEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QLineF", "setLine", args)
  }

}

// translate(qreal, qreal)
func (this *QLineF) translate(args ...interface{}) () {
  // translate(qreal, qreal)
  // translate(const class QPointF &)
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
    // invoke: _ZN6QLineF9translateEdd
    // invoke: void translate(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN6QLineF9translateEdd(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN6QLineF9translateERK7QPointF
    // invoke: void translate(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QLineF9translateERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineF", "translate", args)
  }

}

// setP2(const class QPointF &)
func (this *QLineF) setP2(args ...interface{}) () {
  // setP2(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLineF5setP2ERK7QPointF
    // invoke: void setP2(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QLineF5setP2ERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineF", "setP2", args)
  }

}

// unitVector()
func (this *QLineF) unitVector(args ...interface{}) () {
  // unitVector()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF10unitVectorEv
    // invoke: QLineF unitVector()
    C.C_ZNK6QLineF10unitVectorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineF", "unitVector", args)
  }

}

// angleTo(const class QLineF &)
func (this *QLineF) angleTo(args ...interface{}) () {
  // angleTo(const class QLineF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLineF{}) // "const QLineF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF7angleToERKS_
    // invoke: qreal angleTo(const class QLineF &)
    var arg0 = args[0].(QLineF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK6QLineF7angleToERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineF", "angleTo", args)
  }

}

// dx()
func (this *QLineF) dx(args ...interface{}) () {
  // dx()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF2dxEv
    // invoke: qreal dx()
    C.C_ZNK6QLineF2dxEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineF", "dx", args)
  }

}

// dy()
func (this *QLineF) dy(args ...interface{}) () {
  // dy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF2dyEv
    // invoke: qreal dy()
    C.C_ZNK6QLineF2dyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineF", "dy", args)
  }

}

// x2()
func (this *QLineF) x2(args ...interface{}) () {
  // x2()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF2x2Ev
    // invoke: qreal x2()
    C.C_ZNK6QLineF2x2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineF", "x2", args)
  }

}

// x1()
func (this *QLineF) x1(args ...interface{}) () {
  // x1()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF2x1Ev
    // invoke: qreal x1()
    C.C_ZNK6QLineF2x1Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineF", "x1", args)
  }

}

// setPoints(const class QPointF &, const class QPointF &)
func (this *QLineF) setPoints(args ...interface{}) () {
  // setPoints(const class QPointF &, const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLineF9setPointsERK7QPointFS2_
    // invoke: void setPoints(const class QPointF &, const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPointF).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN6QLineF9setPointsERK7QPointFS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLineF", "setPoints", args)
  }

}

// setLength(qreal)
func (this *QLineF) setLength(args ...interface{}) () {
  // setLength(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLineF9setLengthEd
    // invoke: void setLength(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN6QLineF9setLengthEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineF", "setLength", args)
  }

}

// p2()
func (this *QLineF) p2(args ...interface{}) () {
  // p2()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF2p2Ev
    // invoke: QPointF p2()
    C.C_ZNK6QLineF2p2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineF", "p2", args)
  }

}

// p1()
func (this *QLineF) p1(args ...interface{}) () {
  // p1()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF2p1Ev
    // invoke: QPointF p1()
    C.C_ZNK6QLineF2p1Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineF", "p1", args)
  }

}

// pointAt(qreal)
func (this *QLineF) pointAt(args ...interface{}) () {
  // pointAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF7pointAtEd
    // invoke: QPointF pointAt(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZNK6QLineF7pointAtEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineF", "pointAt", args)
  }

}

// toLine()
func (this *QLineF) toLine(args ...interface{}) () {
  // toLine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF6toLineEv
    // invoke: QLine toLine()
    C.C_ZNK6QLineF6toLineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineF", "toLine", args)
  }

}

// isNull()
func (this *QLineF) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF6isNullEv
    // invoke: bool isNull()
    C.C_ZNK6QLineF6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineF", "isNull", args)
  }

}

// length()
func (this *QLineF) length(args ...interface{}) () {
  // length()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF6lengthEv
    // invoke: qreal length()
    C.C_ZNK6QLineF6lengthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineF", "length", args)
  }

}

// translated(const class QPointF &)
func (this *QLineF) translated(args ...interface{}) () {
  // translated(const class QPointF &)
  // translated(qreal, qreal)
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
    // invoke: _ZNK6QLineF10translatedERK7QPointF
    // invoke: QLineF translated(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK6QLineF10translatedERK7QPointF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK6QLineF10translatedEdd
    // invoke: QLineF translated(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZNK6QLineF10translatedEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLineF", "translated", args)
  }

}

// QLineF(qreal, qreal, qreal, qreal)
func NewQLineF(args ...interface{}) QLineF {
  // QLineF(qreal, qreal, qreal, qreal)
  // QLineF(const class QLine &)
  // QLineF(const class QPointF &, const class QPointF &)
  // QLineF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QLine{}) // "const QLine &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[2][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[3] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLineFC1Edddd
    // invoke: void QLineF(qreal, qreal, qreal, qreal)
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
    C.C_ZN6QLineFC2Edddd(qthis, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN6QLineFC1ERK5QLine
    // invoke: void QLineF(const class QLine &)
    var arg0 = args[0].(QLine).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN6QLineFC2ERK5QLine(qthis, arg0)
  case 2:
    // invoke: _ZN6QLineFC1ERK7QPointFS2_
    // invoke: void QLineF(const class QPointF &, const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPointF).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN6QLineFC2ERK7QPointFS2_(qthis, arg0, arg1)
  case 3:
    // invoke: _ZN6QLineFC1Ev
    // invoke: void QLineF()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN6QLineFC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QLineF", "QLineF", args)
  }

  return QLineF{}
}

// <= body block end

