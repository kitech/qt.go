package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
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
  // proto:  bool QLine::isNull();
extern void _ZNK5QLine6isNullEv(void* qthis);
  // proto:  QLine QLine::translated(const QPoint & p);
extern void _ZNK5QLine10translatedERK6QPoint(void* qthis, void* arg0);
  // proto:  void QLine::setP2(const QPoint & p2);
extern void demth_ZN5QLine5setP2ERK6QPoint(void* qthis, void* arg0);
  // proto:  int QLine::x2();
extern void _ZNK5QLine2x2Ev(void* qthis);
  // proto:  void QLine::QLine(const QPoint & pt1, const QPoint & pt2);
extern void* dector_ZN5QLineC1ERK6QPointS2_(void* arg0, void* arg1);
extern void _ZN5QLineC1ERK6QPointS2_(void* qthis, void* arg0, void* arg1);
  // proto:  void QLine::setP1(const QPoint & p1);
extern void demth_ZN5QLine5setP1ERK6QPoint(void* qthis, void* arg0);
  // proto:  void QLine::translate(const QPoint & p);
extern void demth_ZN5QLine9translateERK6QPoint(void* qthis, void* arg0);
  // proto:  int QLine::dx();
extern void _ZNK5QLine2dxEv(void* qthis);
  // proto:  int QLine::y2();
extern void _ZNK5QLine2y2Ev(void* qthis);
  // proto:  int QLine::dy();
extern void _ZNK5QLine2dyEv(void* qthis);
  // proto:  int QLine::y1();
extern void _ZNK5QLine2y1Ev(void* qthis);
  // proto:  QPoint QLine::p1();
extern void _ZNK5QLine2p1Ev(void* qthis);
  // proto:  QPoint QLine::p2();
extern void _ZNK5QLine2p2Ev(void* qthis);
  // proto:  void QLine::QLine(int x1, int y1, int x2, int y2);
extern void* dector_ZN5QLineC1Eiiii(int arg0, int arg1, int arg2, int arg3);
extern void _ZN5QLineC1Eiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  void QLine::translate(int dx, int dy);
extern void demth_ZN5QLine9translateEii(void* qthis, int arg0, int arg1);
  // proto:  QLine QLine::translated(int dx, int dy);
extern void _ZNK5QLine10translatedEii(void* qthis, int arg0, int arg1);
  // proto:  void QLine::setPoints(const QPoint & p1, const QPoint & p2);
extern void demth_ZN5QLine9setPointsERK6QPointS2_(void* qthis, void* arg0, void* arg1);
  // proto:  void QLine::setLine(int x1, int y1, int x2, int y2);
extern void demth_ZN5QLine7setLineEiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  int QLine::x1();
extern void _ZNK5QLine2x1Ev(void* qthis);
  // proto:  void QLine::QLine();
extern void* dector_ZN5QLineC1Ev();
extern void _ZN5QLineC1Ev(void* qthis);
  // proto:  void QLineF::translate(qreal dx, qreal dy);
extern void demth_ZN6QLineF9translateEdd(void* qthis, double arg0, double arg1);
  // proto:  void QLineF::setPoints(const QPointF & p1, const QPointF & p2);
extern void demth_ZN6QLineF9setPointsERK7QPointFS2_(void* qthis, void* arg0, void* arg1);
  // proto:  void QLineF::setP2(const QPointF & p2);
extern void demth_ZN6QLineF5setP2ERK7QPointF(void* qthis, void* arg0);
  // proto:  QLineF QLineF::translated(qreal dx, qreal dy);
extern void _ZNK6QLineF10translatedEdd(void* qthis, double arg0, double arg1);
  // proto:  void QLineF::setLength(qreal len);
extern void _ZN6QLineF9setLengthEd(void* qthis, double arg0);
  // proto:  qreal QLineF::x1();
extern void _ZNK6QLineF2x1Ev(void* qthis);
  // proto:  qreal QLineF::angle();
extern void _ZNK6QLineF5angleEv(void* qthis);
  // proto:  void QLineF::QLineF(const QPointF & pt1, const QPointF & pt2);
extern void* dector_ZN6QLineFC1ERK7QPointFS2_(void* arg0, void* arg1);
extern void _ZN6QLineFC1ERK7QPointFS2_(void* qthis, void* arg0, void* arg1);
  // proto:  qreal QLineF::length();
extern void _ZNK6QLineF6lengthEv(void* qthis);
  // proto:  void QLineF::QLineF(const QLine & line);
extern void* dector_ZN6QLineFC1ERK5QLine(void* arg0);
extern void _ZN6QLineFC1ERK5QLine(void* qthis, void* arg0);
  // proto:  void QLineF::setAngle(qreal angle);
extern void _ZN6QLineF8setAngleEd(void* qthis, double arg0);
  // proto:  qreal QLineF::x2();
extern void _ZNK6QLineF2x2Ev(void* qthis);
  // proto:  void QLineF::translate(const QPointF & p);
extern void demth_ZN6QLineF9translateERK7QPointF(void* qthis, void* arg0);
  // proto:  qreal QLineF::dx();
extern void _ZNK6QLineF2dxEv(void* qthis);
  // proto:  void QLineF::QLineF();
extern void* dector_ZN6QLineFC1Ev();
extern void _ZN6QLineFC1Ev(void* qthis);
  // proto:  QPointF QLineF::p1();
extern void _ZNK6QLineF2p1Ev(void* qthis);
  // proto:  QLineF QLineF::normalVector();
extern void _ZNK6QLineF12normalVectorEv(void* qthis);
  // proto:  QLine QLineF::toLine();
extern void _ZNK6QLineF6toLineEv(void* qthis);
  // proto:  QPointF QLineF::pointAt(qreal t);
extern void _ZNK6QLineF7pointAtEd(void* qthis, double arg0);
  // proto:  QPointF QLineF::p2();
extern void _ZNK6QLineF2p2Ev(void* qthis);
  // proto:  qreal QLineF::y2();
extern void _ZNK6QLineF2y2Ev(void* qthis);
  // proto:  void QLineF::QLineF(qreal x1, qreal y1, qreal x2, qreal y2);
extern void* dector_ZN6QLineFC1Edddd(double arg0, double arg1, double arg2, double arg3);
extern void _ZN6QLineFC1Edddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  qreal QLineF::dy();
extern void _ZNK6QLineF2dyEv(void* qthis);
  // proto:  QLineF QLineF::unitVector();
extern void _ZNK6QLineF10unitVectorEv(void* qthis);
  // proto:  bool QLineF::isNull();
extern void _ZNK6QLineF6isNullEv(void* qthis);
  // proto:  qreal QLineF::y1();
extern void _ZNK6QLineF2y1Ev(void* qthis);
  // proto:  qreal QLineF::angleTo(const QLineF & l);
extern void _ZNK6QLineF7angleToERKS_(void* qthis, void* arg0);
  // proto:  QLineF QLineF::translated(const QPointF & p);
extern void _ZNK6QLineF10translatedERK7QPointF(void* qthis, void* arg0);
  // proto:  void QLineF::setLine(qreal x1, qreal y1, qreal x2, qreal y2);
extern void demth_ZN6QLineF7setLineEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto: static QLineF QLineF::fromPolar(qreal length, qreal angle);
extern void _ZN6QLineF9fromPolarEdd(double arg0, double arg1);
  // proto:  void QLineF::setP1(const QPointF & p1);
extern void demth_ZN6QLineF5setP1ERK7QPointF(void* qthis, void* arg0);
  // proto:  qreal QLineF::angle(const QLineF & l);
extern void _ZNK6QLineF5angleERKS_(void* qthis, void* arg0);
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
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QLineF)=32
type QLineF struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  bool QLine::isNull();
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
  default:
    qtrt.ErrorResolve("QLine", "isNull", args)
  }

}

  // proto:  QLine QLine::translated(const QPoint & p);
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
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK5QLine10translatedEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QLine", "translated", args)
  }

}

  // proto:  void QLine::setP2(const QPoint & p2);
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
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QLine", "setP2", args)
  }

}

  // proto:  int QLine::x2();
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
  default:
    qtrt.ErrorResolve("QLine", "x2", args)
  }

}

  // proto:  void QLine::QLine(const QPoint & pt1, const QPoint & pt2);
func NewQLine(args ...interface{}) QLine {
  return QLine{}
}

  // proto:  void QLine::setP1(const QPoint & p1);
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
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QLine", "setP1", args)
  }

}

  // proto:  void QLine::translate(const QPoint & p);
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
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN5QLine9translateEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QLine", "translate", args)
  }

}

  // proto:  int QLine::dx();
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
  default:
    qtrt.ErrorResolve("QLine", "dx", args)
  }

}

  // proto:  int QLine::y2();
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
  default:
    qtrt.ErrorResolve("QLine", "y2", args)
  }

}

  // proto:  int QLine::dy();
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
  default:
    qtrt.ErrorResolve("QLine", "dy", args)
  }

}

  // proto:  int QLine::y1();
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
  default:
    qtrt.ErrorResolve("QLine", "y1", args)
  }

}

  // proto:  QPoint QLine::p1();
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
  default:
    qtrt.ErrorResolve("QLine", "p1", args)
  }

}

  // proto:  QPoint QLine::p2();
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
  default:
    qtrt.ErrorResolve("QLine", "p2", args)
  }

}

  // proto:  void QLine::setPoints(const QPoint & p1, const QPoint & p2);
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
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPoint).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QLine", "setPoints", args)
  }

}

  // proto:  void QLine::setLine(int x1, int y1, int x2, int y2);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QLine", "setLine", args)
  }

}

  // proto:  int QLine::x1();
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
  default:
    qtrt.ErrorResolve("QLine", "x1", args)
  }

}

  // proto:  void QLineF::translate(qreal dx, qreal dy);
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN6QLineF9translateERK7QPointF
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QLineF", "translate", args)
  }

}

  // proto:  void QLineF::setPoints(const QPointF & p1, const QPointF & p2);
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
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPointF).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QLineF", "setPoints", args)
  }

}

  // proto:  void QLineF::setP2(const QPointF & p2);
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
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QLineF", "setP2", args)
  }

}

  // proto:  QLineF QLineF::translated(qreal dx, qreal dy);
func (this *QLineF) translated(args ...interface{}) () {
  // translated(qreal, qreal)
  // translated(const class QPointF &)
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
    // invoke: _ZNK6QLineF10translatedEdd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZNK6QLineF10translatedERK7QPointF
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QLineF", "translated", args)
  }

}

  // proto:  void QLineF::setLength(qreal len);
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QLineF", "setLength", args)
  }

}

  // proto:  qreal QLineF::x1();
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
  default:
    qtrt.ErrorResolve("QLineF", "x1", args)
  }

}

  // proto:  qreal QLineF::angle();
func (this *QLineF) angle(args ...interface{}) () {
  // angle()
  // angle(const class QLineF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QLineF{}) // "const QLineF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF5angleEv
  case 1:
    // invoke: _ZNK6QLineF5angleERKS_
    var arg0 = args[0].(QLineF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QLineF", "angle", args)
  }

}

  // proto:  void QLineF::QLineF(const QPointF & pt1, const QPointF & pt2);
func NewQLineF(args ...interface{}) QLineF {
  return QLineF{}
}

  // proto:  qreal QLineF::length();
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
  default:
    qtrt.ErrorResolve("QLineF", "length", args)
  }

}

  // proto:  void QLineF::setAngle(qreal angle);
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QLineF", "setAngle", args)
  }

}

  // proto:  qreal QLineF::x2();
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
  default:
    qtrt.ErrorResolve("QLineF", "x2", args)
  }

}

  // proto:  qreal QLineF::dx();
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
  default:
    qtrt.ErrorResolve("QLineF", "dx", args)
  }

}

  // proto:  QPointF QLineF::p1();
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
  default:
    qtrt.ErrorResolve("QLineF", "p1", args)
  }

}

  // proto:  QLineF QLineF::normalVector();
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
  default:
    qtrt.ErrorResolve("QLineF", "normalVector", args)
  }

}

  // proto:  QLine QLineF::toLine();
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
  default:
    qtrt.ErrorResolve("QLineF", "toLine", args)
  }

}

  // proto:  QPointF QLineF::pointAt(qreal t);
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QLineF", "pointAt", args)
  }

}

  // proto:  QPointF QLineF::p2();
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
  default:
    qtrt.ErrorResolve("QLineF", "p2", args)
  }

}

  // proto:  qreal QLineF::y2();
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
  default:
    qtrt.ErrorResolve("QLineF", "y2", args)
  }

}

  // proto:  qreal QLineF::dy();
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
  default:
    qtrt.ErrorResolve("QLineF", "dy", args)
  }

}

  // proto:  QLineF QLineF::unitVector();
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
  default:
    qtrt.ErrorResolve("QLineF", "unitVector", args)
  }

}

  // proto:  bool QLineF::isNull();
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
  default:
    qtrt.ErrorResolve("QLineF", "isNull", args)
  }

}

  // proto:  qreal QLineF::y1();
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
  default:
    qtrt.ErrorResolve("QLineF", "y1", args)
  }

}

  // proto:  qreal QLineF::angleTo(const QLineF & l);
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
    var arg0 = args[0].(QLineF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QLineF", "angleTo", args)
  }

}

  // proto:  void QLineF::setLine(qreal x1, qreal y1, qreal x2, qreal y2);
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QLineF", "setLine", args)
  }

}

  // proto: static QLineF QLineF::fromPolar(qreal length, qreal angle);
func (this *QLineF) fromPolar_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QLineF", "fromPolar", args)
  }

}

  // proto:  void QLineF::setP1(const QPointF & p1);
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
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QLineF", "setP1", args)
  }

}

// <= body block end

