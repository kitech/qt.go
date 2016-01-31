package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
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
extern void* C_ZNK5QLine2p2Ev(void* qthis); // 2
  // proto:  int QLine::y2();
extern int32_t C_ZNK5QLine2y2Ev(void* qthis); // 2
  // proto:  int QLine::dx();
extern int32_t C_ZNK5QLine2dxEv(void* qthis); // 2
  // proto:  int QLine::y1();
extern int32_t C_ZNK5QLine2y1Ev(void* qthis); // 2
  // proto:  void QLine::setP2(const QPoint & p2);
extern void C_ZN5QLine5setP2ERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  int QLine::x1();
extern int32_t C_ZNK5QLine2x1Ev(void* qthis); // 2
  // proto:  void QLine::setLine(int x1, int y1, int x2, int y2);
extern void C_ZN5QLine7setLineEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 2
  // proto:  int QLine::x2();
extern int32_t C_ZNK5QLine2x2Ev(void* qthis); // 2
  // proto:  bool QLine::isNull();
extern bool C_ZNK5QLine6isNullEv(void* qthis); // 2
  // proto:  QPoint QLine::p1();
extern void* C_ZNK5QLine2p1Ev(void* qthis); // 2
  // proto:  void QLine::QLine(int x1, int y1, int x2, int y2);
extern void* C_ZN5QLineC2Eiiii(int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 1
  // proto:  void QLine::QLine(const QPoint & pt1, const QPoint & pt2);
extern void* C_ZN5QLineC2ERK6QPointS2_(void* arg0, void* arg1); // 1
  // proto:  void QLine::QLine();
extern void* C_ZN5QLineC2Ev(); // 1
  // proto:  int QLine::dy();
extern int32_t C_ZNK5QLine2dyEv(void* qthis); // 2
  // proto:  QLine QLine::translated(const QPoint & p);
extern void* C_ZNK5QLine10translatedERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  QLine QLine::translated(int dx, int dy);
extern void* C_ZNK5QLine10translatedEii(void* qthis, int32_t arg0, int32_t arg1); // 2
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
extern void* C_ZN6QLineF9fromPolarEdd(double arg0, double arg1); // 4
  // proto:  qreal QLineF::y1();
extern double C_ZNK6QLineF2y1Ev(void* qthis); // 2
  // proto:  void QLineF::setP1(const QPointF & p1);
extern void C_ZN6QLineF5setP1ERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  qreal QLineF::y2();
extern double C_ZNK6QLineF2y2Ev(void* qthis); // 2
  // proto:  QLineF QLineF::normalVector();
extern void* C_ZNK6QLineF12normalVectorEv(void* qthis); // 2
  // proto:  qreal QLineF::angle(const QLineF & l);
extern double C_ZNK6QLineF5angleERKS_(void* qthis, void* arg0); // 4
  // proto:  qreal QLineF::angle();
extern double C_ZNK6QLineF5angleEv(void* qthis); // 4
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
extern void* C_ZNK6QLineF10unitVectorEv(void* qthis); // 4
  // proto:  qreal QLineF::angleTo(const QLineF & l);
extern double C_ZNK6QLineF7angleToERKS_(void* qthis, void* arg0); // 4
  // proto:  qreal QLineF::dx();
extern double C_ZNK6QLineF2dxEv(void* qthis); // 2
  // proto:  qreal QLineF::dy();
extern double C_ZNK6QLineF2dyEv(void* qthis); // 2
  // proto:  qreal QLineF::x2();
extern double C_ZNK6QLineF2x2Ev(void* qthis); // 2
  // proto:  qreal QLineF::x1();
extern double C_ZNK6QLineF2x1Ev(void* qthis); // 2
  // proto:  void QLineF::setPoints(const QPointF & p1, const QPointF & p2);
extern void C_ZN6QLineF9setPointsERK7QPointFS2_(void* qthis, void* arg0, void* arg1); // 2
  // proto:  void QLineF::setLength(qreal len);
extern void C_ZN6QLineF9setLengthEd(void* qthis, double arg0); // 2
  // proto:  QPointF QLineF::p2();
extern void* C_ZNK6QLineF2p2Ev(void* qthis); // 2
  // proto:  QPointF QLineF::p1();
extern void* C_ZNK6QLineF2p1Ev(void* qthis); // 2
  // proto:  QPointF QLineF::pointAt(qreal t);
extern void* C_ZNK6QLineF7pointAtEd(void* qthis, double arg0); // 2
  // proto:  QLine QLineF::toLine();
extern void* C_ZNK6QLineF6toLineEv(void* qthis); // 2
  // proto:  bool QLineF::isNull();
extern bool C_ZNK6QLineF6isNullEv(void* qthis); // 2
  // proto:  qreal QLineF::length();
extern double C_ZNK6QLineF6lengthEv(void* qthis); // 4
  // proto:  QLineF QLineF::translated(const QPointF & p);
extern void* C_ZNK6QLineF10translatedERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  QLineF QLineF::translated(qreal dx, qreal dy);
extern void* C_ZNK6QLineF10translatedEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QLineF::QLineF(qreal x1, qreal y1, qreal x2, qreal y2);
extern void* C_ZN6QLineFC2Edddd(double arg0, double arg1, double arg2, double arg3); // 1
  // proto:  void QLineF::QLineF(const QLine & line);
extern void* C_ZN6QLineFC2ERK5QLine(void* arg0); // 1
  // proto:  void QLineF::QLineF(const QPointF & pt1, const QPointF & pt2);
extern void* C_ZN6QLineFC2ERK7QPointFS2_(void* arg0, void* arg1); // 1
  // proto:  void QLineF::QLineF();
extern void* C_ZN6QLineFC2Ev(); // 1
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
func (this *QLine) P2(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QLine2p2Ev(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPoint{}) // "QPoint"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLine", "p2", args)
  }

  return
}

// y2()
func (this *QLine) Y2(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QLine2y2Ev(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLine", "y2", args)
  }

  return
}

// dx()
func (this *QLine) Dx(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QLine2dxEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLine", "dx", args)
  }

  return
}

// y1()
func (this *QLine) Y1(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QLine2y1Ev(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLine", "y1", args)
  }

  return
}

// setP2(const class QPoint &)
func (this *QLine) Setp2(args ...interface{}) () {
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

  return
}

// x1()
func (this *QLine) X1(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QLine2x1Ev(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLine", "x1", args)
  }

  return
}

// setLine(int, int, int, int)
func (this *QLine) Setline(args ...interface{}) () {
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

  return
}

// x2()
func (this *QLine) X2(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QLine2x2Ev(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLine", "x2", args)
  }

  return
}

// isNull()
func (this *QLine) Isnull(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QLine6isNullEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLine", "isNull", args)
  }

  return
}

// p1()
func (this *QLine) P1(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QLine2p1Ev(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPoint{}) // "QPoint"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLine", "p1", args)
  }

  return
}

// QLine(int, int, int, int)
func NewQLine(args ...interface{}) *QLine {
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
    qthis = C.C_ZN5QLineC2Eiiii(arg0, arg1, arg2, arg3)
    return &QLine{qclsinst:qthis}
  case 1:
    // invoke: _ZN5QLineC1ERK6QPointS2_
    // invoke: void QLine(const class QPoint &, const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPoint).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QLineC2ERK6QPointS2_(arg0, arg1)
    return &QLine{qclsinst:qthis}
  case 2:
    // invoke: _ZN5QLineC1Ev
    // invoke: void QLine()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QLineC2Ev()
    return &QLine{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QLine", "QLine", args)
  }

  return nil // QLine{qclsinst:qthis}
}

// dy()
func (this *QLine) Dy(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QLine2dyEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLine", "dy", args)
  }

  return
}

// translated(const class QPoint &)
func (this *QLine) Translated(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QLine10translatedERK6QPoint(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLine{}) // "QLine"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZNK5QLine10translatedEii
    // invoke: QLine translated(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK5QLine10translatedEii(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLine{}) // "QLine"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLine", "translated", args)
  }

  return
}

// setP1(const class QPoint &)
func (this *QLine) Setp1(args ...interface{}) () {
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

  return
}

// translate(const class QPoint &)
func (this *QLine) Translate(args ...interface{}) () {
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

  return
}

// setPoints(const class QPoint &, const class QPoint &)
func (this *QLine) Setpoints(args ...interface{}) () {
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

  return
}

// setAngle(qreal)
func (this *QLineF) Setangle(args ...interface{}) () {
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

  return
}

// fromPolar(qreal, qreal)
func (this *QLineF) Frompolar_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN6QLineF9fromPolarEdd(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLineF{}) // "QLineF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLineF", "fromPolar", args)
  }

  return
}

// y1()
func (this *QLineF) Y1(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QLineF2y1Ev(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLineF", "y1", args)
  }

  return
}

// setP1(const class QPointF &)
func (this *QLineF) Setp1(args ...interface{}) () {
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

  return
}

// y2()
func (this *QLineF) Y2(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QLineF2y2Ev(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLineF", "y2", args)
  }

  return
}

// normalVector()
func (this *QLineF) Normalvector(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QLineF12normalVectorEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLineF{}) // "QLineF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLineF", "normalVector", args)
  }

  return
}

// angle(const class QLineF &)
func (this *QLineF) Angle(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QLineF5angleERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZNK6QLineF5angleEv
    // invoke: qreal angle()
    var ret0 = C.C_ZNK6QLineF5angleEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLineF", "angle", args)
  }

  return
}

// intersect(const class QLineF &, class QPointF *)
func (this *QLineF) Intersect(args ...interface{}) () {
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

  return
}

// setLine(qreal, qreal, qreal, qreal)
func (this *QLineF) Setline(args ...interface{}) () {
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

  return
}

// translate(qreal, qreal)
func (this *QLineF) Translate(args ...interface{}) () {
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

  return
}

// setP2(const class QPointF &)
func (this *QLineF) Setp2(args ...interface{}) () {
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

  return
}

// unitVector()
func (this *QLineF) Unitvector(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QLineF10unitVectorEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLineF{}) // "QLineF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLineF", "unitVector", args)
  }

  return
}

// angleTo(const class QLineF &)
func (this *QLineF) Angleto(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QLineF7angleToERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLineF", "angleTo", args)
  }

  return
}

// dx()
func (this *QLineF) Dx(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QLineF2dxEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLineF", "dx", args)
  }

  return
}

// dy()
func (this *QLineF) Dy(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QLineF2dyEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLineF", "dy", args)
  }

  return
}

// x2()
func (this *QLineF) X2(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QLineF2x2Ev(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLineF", "x2", args)
  }

  return
}

// x1()
func (this *QLineF) X1(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QLineF2x1Ev(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLineF", "x1", args)
  }

  return
}

// setPoints(const class QPointF &, const class QPointF &)
func (this *QLineF) Setpoints(args ...interface{}) () {
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

  return
}

// setLength(qreal)
func (this *QLineF) Setlength(args ...interface{}) () {
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

  return
}

// p2()
func (this *QLineF) P2(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QLineF2p2Ev(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLineF", "p2", args)
  }

  return
}

// p1()
func (this *QLineF) P1(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QLineF2p1Ev(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLineF", "p1", args)
  }

  return
}

// pointAt(qreal)
func (this *QLineF) Pointat(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QLineF7pointAtEd(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLineF", "pointAt", args)
  }

  return
}

// toLine()
func (this *QLineF) Toline(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QLineF6toLineEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLine{}) // "QLine"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLineF", "toLine", args)
  }

  return
}

// isNull()
func (this *QLineF) Isnull(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QLineF6isNullEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLineF", "isNull", args)
  }

  return
}

// length()
func (this *QLineF) Length(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QLineF6lengthEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLineF", "length", args)
  }

  return
}

// translated(const class QPointF &)
func (this *QLineF) Translated(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QLineF10translatedERK7QPointF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLineF{}) // "QLineF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZNK6QLineF10translatedEdd
    // invoke: QLineF translated(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK6QLineF10translatedEdd(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLineF{}) // "QLineF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLineF", "translated", args)
  }

  return
}

// QLineF(qreal, qreal, qreal, qreal)
func NewQLineF(args ...interface{}) *QLineF {
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
    qthis = C.C_ZN6QLineFC2Edddd(arg0, arg1, arg2, arg3)
    return &QLineF{qclsinst:qthis}
  case 1:
    // invoke: _ZN6QLineFC1ERK5QLine
    // invoke: void QLineF(const class QLine &)
    var arg0 = args[0].(QLine).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QLineFC2ERK5QLine(arg0)
    return &QLineF{qclsinst:qthis}
  case 2:
    // invoke: _ZN6QLineFC1ERK7QPointFS2_
    // invoke: void QLineF(const class QPointF &, const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPointF).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QLineFC2ERK7QPointFS2_(arg0, arg1)
    return &QLineF{qclsinst:qthis}
  case 3:
    // invoke: _ZN6QLineFC1Ev
    // invoke: void QLineF()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QLineFC2Ev()
    return &QLineF{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QLineF", "QLineF", args)
  }

  return nil // QLineF{qclsinst:qthis}
}

// <= body block end

