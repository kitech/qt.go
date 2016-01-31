package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QPolygon QPolygon::united(const QPolygon & r);
extern void C_ZNK8QPolygon6unitedERKS_(void* qthis, void* arg0); // 4
  // proto:  QPolygon QPolygon::subtracted(const QPolygon & r);
extern void C_ZNK8QPolygon10subtractedERKS_(void* qthis, void* arg0); // 4
  // proto:  QRect QPolygon::boundingRect();
extern void C_ZNK8QPolygon12boundingRectEv(void* qthis); // 4
  // proto:  QPolygon QPolygon::intersected(const QPolygon & r);
extern void C_ZNK8QPolygon11intersectedERKS_(void* qthis, void* arg0); // 4
  // proto:  void QPolygon::QPolygon(const QRect & r, bool closed);
extern void C_ZN8QPolygonC2ERK5QRectb(void* qthis, void* arg0, bool arg1); // 3
  // proto:  void QPolygon::QPolygon(const QPolygon & a);
extern void C_ZN8QPolygonC2ERKS_(void* qthis, void* arg0); // 1
  // proto:  void QPolygon::QPolygon(int nPoints, const int * points);
extern void C_ZN8QPolygonC2EiPKi(void* qthis, int32_t arg0, int32_t* arg1); // 3
  // proto:  void QPolygon::QPolygon(int size);
extern void C_ZN8QPolygonC2Ei(void* qthis, int32_t arg0); // 1
  // proto:  void QPolygon::QPolygon();
extern void C_ZN8QPolygonC2Ev(void* qthis); // 1
  // proto:  void QPolygon::setPoint(int index, const QPoint & p);
extern void C_ZN8QPolygon8setPointEiRK6QPoint(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  void QPolygon::setPoint(int index, int x, int y);
extern void C_ZN8QPolygon8setPointEiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2); // 2
  // proto:  void QPolygon::point(int i, int * x, int * y);
extern void C_ZNK8QPolygon5pointEiPiS0_(void* qthis, int32_t arg0, int32_t* arg1, int32_t* arg2); // 4
  // proto:  QPoint QPolygon::point(int i);
extern void C_ZNK8QPolygon5pointEi(void* qthis, int32_t arg0); // 2
  // proto:  void QPolygon::swap(QPolygon & other);
extern void C_ZN8QPolygon4swapERS_(void* qthis, void* arg0); // 2
  // proto:  void QPolygon::putPoints(int index, int nPoints, const int * points);
extern void C_ZN8QPolygon9putPointsEiiPKi(void* qthis, int32_t arg0, int32_t arg1, int32_t* arg2); // 4
  // proto:  void QPolygon::putPoints(int index, int nPoints, const QPolygon & from, int fromIndex);
extern void C_ZN8QPolygon9putPointsEiiRKS_i(void* qthis, int32_t arg0, int32_t arg1, void* arg2, int32_t arg3); // 4
  // proto:  void QPolygon::putPoints(int index, int nPoints, int firstx, int firsty);
extern void C_ZN8QPolygon9putPointsEiiiiz(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  QPolygon QPolygon::translated(int dx, int dy);
extern void C_ZNK8QPolygon10translatedEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QPolygon QPolygon::translated(const QPoint & offset);
extern void C_ZNK8QPolygon10translatedERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  void QPolygon::translate(int dx, int dy);
extern void C_ZN8QPolygon9translateEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QPolygon::translate(const QPoint & offset);
extern void C_ZN8QPolygon9translateERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  void QPolygon::setPoints(int nPoints, int firstx, int firsty);
extern void C_ZN8QPolygon9setPointsEiiiz(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  void QPolygon::setPoints(int nPoints, const int * points);
extern void C_ZN8QPolygon9setPointsEiPKi(void* qthis, int32_t arg0, int32_t* arg1); // 4
  // proto:  void QPolygon::~QPolygon();
extern void C_ZN8QPolygonD2Ev(void* qthis); // 2
  // proto:  QPolygonF QPolygonF::united(const QPolygonF & r);
extern void C_ZNK9QPolygonF6unitedERKS_(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QPolygonF::subtracted(const QPolygonF & r);
extern void C_ZNK9QPolygonF10subtractedERKS_(void* qthis, void* arg0); // 4
  // proto:  QRectF QPolygonF::boundingRect();
extern void C_ZNK9QPolygonF12boundingRectEv(void* qthis); // 4
  // proto:  QPolygonF QPolygonF::intersected(const QPolygonF & r);
extern void C_ZNK9QPolygonF11intersectedERKS_(void* qthis, void* arg0); // 4
  // proto:  void QPolygonF::~QPolygonF();
extern void C_ZN9QPolygonFD2Ev(void* qthis); // 2
  // proto:  void QPolygonF::QPolygonF(const QRectF & r);
extern void C_ZN9QPolygonFC2ERK6QRectF(void* qthis, void* arg0); // 3
  // proto:  void QPolygonF::QPolygonF(const QPolygonF & a);
extern void C_ZN9QPolygonFC2ERKS_(void* qthis, void* arg0); // 1
  // proto:  void QPolygonF::QPolygonF();
extern void C_ZN9QPolygonFC2Ev(void* qthis); // 1
  // proto:  void QPolygonF::QPolygonF(int size);
extern void C_ZN9QPolygonFC2Ei(void* qthis, int32_t arg0); // 1
  // proto:  void QPolygonF::QPolygonF(const QPolygon & a);
extern void C_ZN9QPolygonFC2ERK8QPolygon(void* qthis, void* arg0); // 3
  // proto:  bool QPolygonF::isClosed();
extern void C_ZNK9QPolygonF8isClosedEv(void* qthis); // 2
  // proto:  void QPolygonF::swap(QPolygonF & other);
extern void C_ZN9QPolygonF4swapERS_(void* qthis, void* arg0); // 2
  // proto:  QPolygon QPolygonF::toPolygon();
extern void C_ZNK9QPolygonF9toPolygonEv(void* qthis); // 4
  // proto:  QPolygonF QPolygonF::translated(qreal dx, qreal dy);
extern void C_ZNK9QPolygonF10translatedEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  QPolygonF QPolygonF::translated(const QPointF & offset);
extern void C_ZNK9QPolygonF10translatedERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QPolygonF::translate(const QPointF & offset);
extern void C_ZN9QPolygonF9translateERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QPolygonF::translate(qreal dx, qreal dy);
extern void C_ZN9QPolygonF9translateEdd(void* qthis, double arg0, double arg1); // 2
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QPolygonF)=1
type QPolygonF struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// united(const class QPolygon &)
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
    C.C_ZNK8QPolygon6unitedERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPolygon", "united", args)
  }

}

// subtracted(const class QPolygon &)
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
    C.C_ZNK8QPolygon10subtractedERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPolygon", "subtracted", args)
  }

}

// boundingRect()
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
    C.C_ZNK8QPolygon12boundingRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPolygon", "boundingRect", args)
  }

}

// intersected(const class QPolygon &)
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
    C.C_ZNK8QPolygon11intersectedERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPolygon", "intersected", args)
  }

}

// QPolygon(const class QRect &, _Bool)
func NewQPolygon(args ...interface{}) QPolygon {
  // QPolygon(const class QRect &, _Bool)
  // QPolygon(const class QPolygon &)
  // QPolygon(int, const int *)
  // QPolygon(int)
  // QPolygon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPolygon{}) // "const QPolygon &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(true) // "const int *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPolygonC1ERK5QRectb
    // invoke: void QPolygon(const class QRect &, _Bool)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN8QPolygonC2ERK5QRectb(qthis, arg0, arg1)
  case 1:
    // invoke: _ZN8QPolygonC1ERKS_
    // invoke: void QPolygon(const class QPolygon &)
    var arg0 = args[0].(QPolygon).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN8QPolygonC2ERKS_(qthis, arg0)
  case 2:
    // invoke: _ZN8QPolygonC1EiPKi
    // invoke: void QPolygon(int, const int *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN8QPolygonC2EiPKi(qthis, arg0, arg1)
  case 3:
    // invoke: _ZN8QPolygonC1Ei
    // invoke: void QPolygon(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN8QPolygonC2Ei(qthis, arg0)
  case 4:
    // invoke: _ZN8QPolygonC1Ev
    // invoke: void QPolygon()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN8QPolygonC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QPolygon", "QPolygon", args)
  }

  return QPolygon{}
}

// setPoint(int, const class QPoint &)
func (this *QPolygon) setPoint(args ...interface{}) () {
  // setPoint(int, const class QPoint &)
  // setPoint(int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPolygon8setPointEiRK6QPoint
    // invoke: void setPoint(int, const class QPoint &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPoint).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QPolygon8setPointEiRK6QPoint(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QPolygon8setPointEiii
    // invoke: void setPoint(int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN8QPolygon8setPointEiii(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QPolygon", "setPoint", args)
  }

}

// point(int, int *, int *)
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
    C.C_ZNK8QPolygon5pointEiPiS0_(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZNK8QPolygon5pointEi
    // invoke: QPoint point(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK8QPolygon5pointEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPolygon", "point", args)
  }

}

// swap(class QPolygon &)
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
    C.C_ZN8QPolygon4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPolygon", "swap", args)
  }

}

// putPoints(int, int, const int *)
func (this *QPolygon) putPoints(args ...interface{}) () {
  // putPoints(int, int, const int *)
  // putPoints(int, int, const class QPolygon &, int)
  // putPoints(int, int, int, int, ...)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(true) // "const int *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QPolygon{}) // "const QPolygon &"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[2][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPolygon9putPointsEiiPKi
    // invoke: void putPoints(int, int, const int *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    C.C_ZN8QPolygon9putPointsEiiPKi(this.qclsinst, arg0, arg1, arg2)
  case 1:
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
    C.C_ZN8QPolygon9putPointsEiiRKS_i(this.qclsinst, arg0, arg1, arg2, arg3)
  case 2:
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
    C.C_ZN8QPolygon9putPointsEiiiiz(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QPolygon", "putPoints", args)
  }

}

// translated(int, int)
func (this *QPolygon) translated(args ...interface{}) () {
  // translated(int, int)
  // translated(const class QPoint &)
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
    // invoke: _ZNK8QPolygon10translatedEii
    // invoke: QPolygon translated(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK8QPolygon10translatedEii(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK8QPolygon10translatedERK6QPoint
    // invoke: QPolygon translated(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK8QPolygon10translatedERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPolygon", "translated", args)
  }

}

// translate(int, int)
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
    C.C_ZN8QPolygon9translateEii(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QPolygon9translateERK6QPoint
    // invoke: void translate(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPolygon9translateERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPolygon", "translate", args)
  }

}

// setPoints(int, int, int, ...)
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
    C.C_ZN8QPolygon9setPointsEiiiz(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN8QPolygon9setPointsEiPKi
    // invoke: void setPoints(int, const int *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPolygon9setPointsEiPKi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPolygon", "setPoints", args)
  }

}

// ~QPolygon()
func (this *QPolygon) FreeQPolygon(args ...interface{}) () {
  // ~QPolygon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPolygonD0Ev
    // invoke: void ~QPolygon()
    C.C_ZN8QPolygonD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPolygon", "~QPolygon", args)
  }

}

// united(const class QPolygonF &)
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
    C.C_ZNK9QPolygonF6unitedERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPolygonF", "united", args)
  }

}

// subtracted(const class QPolygonF &)
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
    C.C_ZNK9QPolygonF10subtractedERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPolygonF", "subtracted", args)
  }

}

// boundingRect()
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
    C.C_ZNK9QPolygonF12boundingRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPolygonF", "boundingRect", args)
  }

}

// intersected(const class QPolygonF &)
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
    C.C_ZNK9QPolygonF11intersectedERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPolygonF", "intersected", args)
  }

}

// ~QPolygonF()
func (this *QPolygonF) FreeQPolygonF(args ...interface{}) () {
  // ~QPolygonF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QPolygonFD0Ev
    // invoke: void ~QPolygonF()
    C.C_ZN9QPolygonFD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPolygonF", "~QPolygonF", args)
  }

}

// QPolygonF(const class QRectF &)
func NewQPolygonF(args ...interface{}) QPolygonF {
  // QPolygonF(const class QRectF &)
  // QPolygonF(const class QPolygonF &)
  // QPolygonF()
  // QPolygonF(int)
  // QPolygonF(const class QPolygon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QPolygon{}) // "const QPolygon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QPolygonFC1ERK6QRectF
    // invoke: void QPolygonF(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN9QPolygonFC2ERK6QRectF(qthis, arg0)
  case 1:
    // invoke: _ZN9QPolygonFC1ERKS_
    // invoke: void QPolygonF(const class QPolygonF &)
    var arg0 = args[0].(QPolygonF).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN9QPolygonFC2ERKS_(qthis, arg0)
  case 2:
    // invoke: _ZN9QPolygonFC1Ev
    // invoke: void QPolygonF()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN9QPolygonFC2Ev(qthis)
  case 3:
    // invoke: _ZN9QPolygonFC1Ei
    // invoke: void QPolygonF(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN9QPolygonFC2Ei(qthis, arg0)
  case 4:
    // invoke: _ZN9QPolygonFC1ERK8QPolygon
    // invoke: void QPolygonF(const class QPolygon &)
    var arg0 = args[0].(QPolygon).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN9QPolygonFC2ERK8QPolygon(qthis, arg0)
  default:
    qtrt.ErrorResolve("QPolygonF", "QPolygonF", args)
  }

  return QPolygonF{}
}

// isClosed()
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
    C.C_ZNK9QPolygonF8isClosedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPolygonF", "isClosed", args)
  }

}

// swap(class QPolygonF &)
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
    C.C_ZN9QPolygonF4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPolygonF", "swap", args)
  }

}

// toPolygon()
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
    C.C_ZNK9QPolygonF9toPolygonEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPolygonF", "toPolygon", args)
  }

}

// translated(qreal, qreal)
func (this *QPolygonF) translated(args ...interface{}) () {
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
    // invoke: _ZNK9QPolygonF10translatedEdd
    // invoke: QPolygonF translated(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZNK9QPolygonF10translatedEdd(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK9QPolygonF10translatedERK7QPointF
    // invoke: QPolygonF translated(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK9QPolygonF10translatedERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPolygonF", "translated", args)
  }

}

// translate(const class QPointF &)
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
    C.C_ZN9QPolygonF9translateERK7QPointF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN9QPolygonF9translateEdd
    // invoke: void translate(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN9QPolygonF9translateEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPolygonF", "translate", args)
  }

}

// <= body block end

