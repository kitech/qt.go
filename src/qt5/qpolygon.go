package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtGui/qpolygon.h
// dst-file: /src/gui/qpolygon.go
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
  // proto:  QRect QPolygon::boundingRect();
extern void _ZNK8QPolygon12boundingRectEv(void* qthis);
  // proto:  void QPolygon::setPoint(int index, int x, int y);
extern void demth_ZN8QPolygon8setPointEiii(void* qthis, int arg0, int arg1, int arg2);
  // proto:  void QPolygon::~QPolygon();
extern void demth_ZN8QPolygonD0Ev(void* qthis);
  // proto:  void QPolygon::putPoints(int index, int nPoints, const QPolygon & from, int fromIndex);
extern void _ZN8QPolygon9putPointsEiiRKS_i(void* qthis, int arg0, int arg1, void* arg2, int arg3);
  // proto:  QPolygon QPolygon::translated(const QPoint & offset);
extern void demth_ZNK8QPolygon10translatedERK6QPoint(void* qthis, void* arg0);
  // proto:  QPolygon QPolygon::subtracted(const QPolygon & r);
extern void _ZNK8QPolygon10subtractedERKS_(void* qthis, void* arg0);
  // proto:  QPolygon QPolygon::intersected(const QPolygon & r);
extern void _ZNK8QPolygon11intersectedERKS_(void* qthis, void* arg0);
  // proto:  void QPolygon::setPoint(int index, const QPoint & p);
extern void demth_ZN8QPolygon8setPointEiRK6QPoint(void* qthis, int arg0, void* arg1);
  // proto:  void QPolygon::point(int i, int * x, int * y);
extern void _ZNK8QPolygon5pointEiPiS0_(void* qthis, int arg0, int* arg1, int* arg2);
  // proto:  void QPolygon::translate(int dx, int dy);
extern void _ZN8QPolygon9translateEii(void* qthis, int arg0, int arg1);
  // proto:  void QPolygon::putPoints(int index, int nPoints, int firstx, int firsty);
extern void _ZN8QPolygon9putPointsEiiiiz(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  void QPolygon::setPoints(int nPoints, int firstx, int firsty);
extern void _ZN8QPolygon9setPointsEiiiz(void* qthis, int arg0, int arg1, int arg2);
  // proto:  void QPolygon::translate(const QPoint & offset);
extern void demth_ZN8QPolygon9translateERK6QPoint(void* qthis, void* arg0);
  // proto:  void QPolygon::swap(QPolygon & other);
extern void demth_ZN8QPolygon4swapERS_(void* qthis, void* arg0);
  // proto:  QPoint QPolygon::point(int i);
extern void demth_ZNK8QPolygon5pointEi(void* qthis, int arg0);
  // proto:  void QPolygon::QPolygon(const QPolygon & a);
extern void* dector_ZN8QPolygonC1ERKS_(void* arg0);
extern void demth_ZN8QPolygonC1ERKS_(void* qthis, void* arg0);
  // proto:  void QPolygon::QPolygon(int nPoints, const int * points);
extern void* dector_ZN8QPolygonC1EiPKi(int arg0, int* arg1);
extern void _ZN8QPolygonC1EiPKi(void* qthis, int arg0, int* arg1);
  // proto:  QPolygon QPolygon::united(const QPolygon & r);
extern void _ZNK8QPolygon6unitedERKS_(void* qthis, void* arg0);
  // proto:  QPolygon QPolygon::translated(int dx, int dy);
extern void _ZNK8QPolygon10translatedEii(void* qthis, int arg0, int arg1);
  // proto:  void QPolygon::putPoints(int index, int nPoints, const int * points);
extern void _ZN8QPolygon9putPointsEiiPKi(void* qthis, int arg0, int arg1, int* arg2);
  // proto:  void QPolygon::setPoints(int nPoints, const int * points);
extern void _ZN8QPolygon9setPointsEiPKi(void* qthis, int arg0, int* arg1);
  // proto:  void QPolygon::QPolygon(int size);
extern void* dector_ZN8QPolygonC1Ei(int arg0);
extern void demth_ZN8QPolygonC1Ei(void* qthis, int arg0);
  // proto:  void QPolygon::QPolygon();
extern void* dector_ZN8QPolygonC1Ev();
extern void demth_ZN8QPolygonC1Ev(void* qthis);
  // proto:  void QPolygon::QPolygon(const QRect & r, bool closed);
extern void* dector_ZN8QPolygonC1ERK5QRectb(void* arg0, bool arg1);
extern void _ZN8QPolygonC1ERK5QRectb(void* qthis, void* arg0, bool arg1);
  // proto:  QRectF QPolygonF::boundingRect();
extern void _ZNK9QPolygonF12boundingRectEv(void* qthis);
  // proto:  QPolygonF QPolygonF::intersected(const QPolygonF & r);
extern void _ZNK9QPolygonF11intersectedERKS_(void* qthis, void* arg0);
  // proto:  void QPolygonF::QPolygonF(const QPolygon & a);
extern void* dector_ZN9QPolygonFC1ERK8QPolygon(void* arg0);
extern void _ZN9QPolygonFC1ERK8QPolygon(void* qthis, void* arg0);
  // proto:  void QPolygonF::QPolygonF(const QRectF & r);
extern void* dector_ZN9QPolygonFC1ERK6QRectF(void* arg0);
extern void _ZN9QPolygonFC1ERK6QRectF(void* qthis, void* arg0);
  // proto:  QPolygon QPolygonF::toPolygon();
extern void _ZNK9QPolygonF9toPolygonEv(void* qthis);
  // proto:  void QPolygonF::~QPolygonF();
extern void demth_ZN9QPolygonFD0Ev(void* qthis);
  // proto:  void QPolygonF::QPolygonF(int size);
extern void* dector_ZN9QPolygonFC1Ei(int arg0);
extern void demth_ZN9QPolygonFC1Ei(void* qthis, int arg0);
  // proto:  QPolygonF QPolygonF::subtracted(const QPolygonF & r);
extern void _ZNK9QPolygonF10subtractedERKS_(void* qthis, void* arg0);
  // proto:  void QPolygonF::QPolygonF();
extern void* dector_ZN9QPolygonFC1Ev();
extern void demth_ZN9QPolygonFC1Ev(void* qthis);
  // proto:  void QPolygonF::translate(const QPointF & offset);
extern void _ZN9QPolygonF9translateERK7QPointF(void* qthis, void* arg0);
  // proto:  void QPolygonF::swap(QPolygonF & other);
extern void demth_ZN9QPolygonF4swapERS_(void* qthis, void* arg0);
  // proto:  QPolygonF QPolygonF::translated(const QPointF & offset);
extern void _ZNK9QPolygonF10translatedERK7QPointF(void* qthis, void* arg0);
  // proto:  void QPolygonF::translate(qreal dx, qreal dy);
extern void demth_ZN9QPolygonF9translateEdd(void* qthis, double arg0, double arg1);
  // proto:  void QPolygonF::QPolygonF(const QPolygonF & a);
extern void* dector_ZN9QPolygonFC1ERKS_(void* arg0);
extern void demth_ZN9QPolygonFC1ERKS_(void* qthis, void* arg0);
  // proto:  QPolygonF QPolygonF::translated(qreal dx, qreal dy);
extern void demth_ZNK9QPolygonF10translatedEdd(void* qthis, double arg0, double arg1);
  // proto:  bool QPolygonF::isClosed();
extern void _ZNK9QPolygonF8isClosedEv(void* qthis);
  // proto:  QPolygonF QPolygonF::united(const QPolygonF & r);
extern void _ZNK9QPolygonF6unitedERKS_(void* qthis, void* arg0);
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

// class sizeof(QPolygon)=1
type QPolygon struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QPolygonF)=1
type QPolygonF struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QRect QPolygon::boundingRect();
func (this *QPolygon) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPolygon12boundingRectEv
    // invoke: QRect boundingRect()
    C._ZNK8QPolygon12boundingRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPolygon", "boundingRect", args)
  }

}

  // proto:  void QPolygon::setPoint(int index, int x, int y);
func (this *QPolygon) setPoint(args ...interface{}) () {
  // setPoint(int, int, int)
  // setPoint(int, const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPolygon8setPointEiii
    // invoke: void setPoint(int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C.demth_ZN8QPolygon8setPointEiii(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN8QPolygon8setPointEiRK6QPoint
    // invoke: void setPoint(int, const class QPoint &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPoint).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZN8QPolygon8setPointEiRK6QPoint(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPolygon", "setPoint", args)
  }

}

  // proto:  void QPolygon::~QPolygon();
func (this *QPolygon) FreeQPolygon(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPolygon", "~QPolygon", args)
  }

}

  // proto:  void QPolygon::putPoints(int index, int nPoints, const QPolygon & from, int fromIndex);
func (this *QPolygon) putPoints(args ...interface{}) () {
  // putPoints(int, int, const class QPolygon &, int)
  // putPoints(int, int, int, int, ...)
  // putPoints(int, int, const int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QPolygon{}) // "const QPolygon &"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(true) // "const int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPolygon9putPointsEiiRKS_i
    // invoke: void putPoints(int, int, const class QPolygon &, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QPolygon).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C._ZN8QPolygon9putPointsEiiRKS_i(this.qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN8QPolygon9putPointsEiiiiz
    // invoke: void putPoints(int, int, int, int, ...)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C._ZN8QPolygon9putPointsEiiiiz(this.qclsinst, arg0, arg1, arg2, arg3)
  case 2:
    // invoke: _ZN8QPolygon9putPointsEiiPKi
    // invoke: void putPoints(int, int, const int *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    C._ZN8QPolygon9putPointsEiiPKi(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QPolygon", "putPoints", args)
  }

}

  // proto:  QPolygon QPolygon::translated(const QPoint & offset);
func (this *QPolygon) translated(args ...interface{}) () {
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
    // invoke: _ZNK8QPolygon10translatedERK6QPoint
    // invoke: QPolygon translated(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZNK8QPolygon10translatedERK6QPoint(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK8QPolygon10translatedEii
    // invoke: QPolygon translated(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK8QPolygon10translatedEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPolygon", "translated", args)
  }

}

  // proto:  QPolygon QPolygon::subtracted(const QPolygon & r);
func (this *QPolygon) subtracted(args ...interface{}) () {
  // subtracted(const class QPolygon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygon{}) // "const QPolygon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPolygon10subtractedERKS_
    // invoke: QPolygon subtracted(const class QPolygon &)
    var arg0 = args[0].(QPolygon).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK8QPolygon10subtractedERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPolygon", "subtracted", args)
  }

}

  // proto:  QPolygon QPolygon::intersected(const QPolygon & r);
func (this *QPolygon) intersected(args ...interface{}) () {
  // intersected(const class QPolygon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygon{}) // "const QPolygon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPolygon11intersectedERKS_
    // invoke: QPolygon intersected(const class QPolygon &)
    var arg0 = args[0].(QPolygon).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK8QPolygon11intersectedERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPolygon", "intersected", args)
  }

}

  // proto:  void QPolygon::point(int i, int * x, int * y);
func (this *QPolygon) point(args ...interface{}) () {
  // point(int, int *, int *)
  // point(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPolygon5pointEiPiS0_
    // invoke: void point(int, int *, int *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    C._ZNK8QPolygon5pointEiPiS0_(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZNK8QPolygon5pointEi
    // invoke: QPoint point(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.demth_ZNK8QPolygon5pointEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPolygon", "point", args)
  }

}

  // proto:  void QPolygon::translate(int dx, int dy);
func (this *QPolygon) translate(args ...interface{}) () {
  // translate(int, int)
  // translate(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPolygon9translateEii
    // invoke: void translate(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN8QPolygon9translateEii(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QPolygon9translateERK6QPoint
    // invoke: void translate(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN8QPolygon9translateERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPolygon", "translate", args)
  }

}

  // proto:  void QPolygon::setPoints(int nPoints, int firstx, int firsty);
func (this *QPolygon) setPoints(args ...interface{}) () {
  // setPoints(int, int, int, ...)
  // setPoints(int, const int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(true) // "const int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPolygon9setPointsEiiiz
    // invoke: void setPoints(int, int, int, ...)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN8QPolygon9setPointsEiiiz(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN8QPolygon9setPointsEiPKi
    // invoke: void setPoints(int, const int *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    C._ZN8QPolygon9setPointsEiPKi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPolygon", "setPoints", args)
  }

}

  // proto:  void QPolygon::swap(QPolygon & other);
func (this *QPolygon) swap(args ...interface{}) () {
  // swap(class QPolygon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygon{}) // "QPolygon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPolygon4swapERS_
    // invoke: void swap(class QPolygon &)
    var arg0 = args[0].(QPolygon).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN8QPolygon4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPolygon", "swap", args)
  }

}

  // proto:  void QPolygon::QPolygon(const QPolygon & a);
func NewQPolygon(args ...interface{}) QPolygon {
  return QPolygon{}
}

  // proto:  QPolygon QPolygon::united(const QPolygon & r);
func (this *QPolygon) united(args ...interface{}) () {
  // united(const class QPolygon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygon{}) // "const QPolygon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPolygon6unitedERKS_
    // invoke: QPolygon united(const class QPolygon &)
    var arg0 = args[0].(QPolygon).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK8QPolygon6unitedERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPolygon", "united", args)
  }

}

  // proto:  QRectF QPolygonF::boundingRect();
func (this *QPolygonF) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QPolygonF12boundingRectEv
    // invoke: QRectF boundingRect()
    C._ZNK9QPolygonF12boundingRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPolygonF", "boundingRect", args)
  }

}

  // proto:  QPolygonF QPolygonF::intersected(const QPolygonF & r);
func (this *QPolygonF) intersected(args ...interface{}) () {
  // intersected(const class QPolygonF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QPolygonF11intersectedERKS_
    // invoke: QPolygonF intersected(const class QPolygonF &)
    var arg0 = args[0].(QPolygonF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK9QPolygonF11intersectedERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPolygonF", "intersected", args)
  }

}

  // proto:  void QPolygonF::QPolygonF(const QPolygon & a);
func NewQPolygonF(args ...interface{}) QPolygonF {
  return QPolygonF{}
}

  // proto:  QPolygon QPolygonF::toPolygon();
func (this *QPolygonF) toPolygon(args ...interface{}) () {
  // toPolygon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QPolygonF9toPolygonEv
    // invoke: QPolygon toPolygon()
    C._ZNK9QPolygonF9toPolygonEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPolygonF", "toPolygon", args)
  }

}

  // proto:  void QPolygonF::~QPolygonF();
func (this *QPolygonF) FreeQPolygonF(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPolygonF", "~QPolygonF", args)
  }

}

  // proto:  QPolygonF QPolygonF::subtracted(const QPolygonF & r);
func (this *QPolygonF) subtracted(args ...interface{}) () {
  // subtracted(const class QPolygonF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QPolygonF10subtractedERKS_
    // invoke: QPolygonF subtracted(const class QPolygonF &)
    var arg0 = args[0].(QPolygonF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK9QPolygonF10subtractedERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPolygonF", "subtracted", args)
  }

}

  // proto:  void QPolygonF::translate(const QPointF & offset);
func (this *QPolygonF) translate(args ...interface{}) () {
  // translate(const class QPointF &)
  // translate(qreal, qreal)
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
    // invoke: _ZN9QPolygonF9translateERK7QPointF
    // invoke: void translate(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QPolygonF9translateERK7QPointF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN9QPolygonF9translateEdd
    // invoke: void translate(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.demth_ZN9QPolygonF9translateEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPolygonF", "translate", args)
  }

}

  // proto:  void QPolygonF::swap(QPolygonF & other);
func (this *QPolygonF) swap(args ...interface{}) () {
  // swap(class QPolygonF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "QPolygonF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QPolygonF4swapERS_
    // invoke: void swap(class QPolygonF &)
    var arg0 = args[0].(QPolygonF).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN9QPolygonF4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPolygonF", "swap", args)
  }

}

  // proto:  QPolygonF QPolygonF::translated(const QPointF & offset);
func (this *QPolygonF) translated(args ...interface{}) () {
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
    // invoke: _ZNK9QPolygonF10translatedERK7QPointF
    // invoke: QPolygonF translated(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK9QPolygonF10translatedERK7QPointF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK9QPolygonF10translatedEdd
    // invoke: QPolygonF translated(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.demth_ZNK9QPolygonF10translatedEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPolygonF", "translated", args)
  }

}

  // proto:  bool QPolygonF::isClosed();
func (this *QPolygonF) isClosed(args ...interface{}) () {
  // isClosed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QPolygonF8isClosedEv
    // invoke: bool isClosed()
    C._ZNK9QPolygonF8isClosedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPolygonF", "isClosed", args)
  }

}

  // proto:  QPolygonF QPolygonF::united(const QPolygonF & r);
func (this *QPolygonF) united(args ...interface{}) () {
  // united(const class QPolygonF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QPolygonF6unitedERKS_
    // invoke: QPolygonF united(const class QPolygonF &)
    var arg0 = args[0].(QPolygonF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK9QPolygonF6unitedERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPolygonF", "united", args)
  }

}

// <= body block end

