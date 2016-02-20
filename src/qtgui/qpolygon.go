package qtgui
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QPolygon QPolygon::united(const QPolygon & r);
extern void* C_ZNK8QPolygon6unitedERKS_(void* qthis, void* arg0); // 4
  // proto:  QPolygon QPolygon::subtracted(const QPolygon & r);
extern void* C_ZNK8QPolygon10subtractedERKS_(void* qthis, void* arg0); // 4
  // proto:  QRect QPolygon::boundingRect();
extern void* C_ZNK8QPolygon12boundingRectEv(void* qthis); // 4
  // proto:  QPolygon QPolygon::intersected(const QPolygon & r);
extern void* C_ZNK8QPolygon11intersectedERKS_(void* qthis, void* arg0); // 4
  // proto:  void QPolygon::QPolygon(const QRect & r, bool closed);
extern void* C_ZN8QPolygonC2ERK5QRectb(void* arg0, bool arg1); // 3
  // proto:  void QPolygon::QPolygon(const QPolygon & a);
extern void* C_ZN8QPolygonC2ERKS_(void* arg0); // 1
  // proto:  void QPolygon::QPolygon(int nPoints, const int * points);
extern void* C_ZN8QPolygonC2EiPKi(int32_t arg0, void* arg1); // 3
  // proto:  void QPolygon::QPolygon(int size);
extern void* C_ZN8QPolygonC2Ei(int32_t arg0); // 1
  // proto:  void QPolygon::QPolygon();
extern void* C_ZN8QPolygonC2Ev(); // 1
  // proto:  void QPolygon::setPoint(int index, const QPoint & p);
extern void C_ZN8QPolygon8setPointEiRK6QPoint(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  void QPolygon::setPoint(int index, int x, int y);
extern void C_ZN8QPolygon8setPointEiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2); // 2
  // proto:  void QPolygon::point(int i, int * x, int * y);
extern void C_ZNK8QPolygon5pointEiPiS0_(void* qthis, int32_t arg0, void* arg1, void* arg2); // 4
  // proto:  QPoint QPolygon::point(int i);
extern void* C_ZNK8QPolygon5pointEi(void* qthis, int32_t arg0); // 2
  // proto:  void QPolygon::swap(QPolygon & other);
extern void C_ZN8QPolygon4swapERS_(void* qthis, void* arg0); // 2
  // proto:  void QPolygon::putPoints(int index, int nPoints, const int * points);
extern void C_ZN8QPolygon9putPointsEiiPKi(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QPolygon::putPoints(int index, int nPoints, const QPolygon & from, int fromIndex);
extern void C_ZN8QPolygon9putPointsEiiRKS_i(void* qthis, int32_t arg0, int32_t arg1, void* arg2, int32_t arg3); // 4
  // proto:  void QPolygon::putPoints(int index, int nPoints, int firstx, int firsty);
extern void C_ZN8QPolygon9putPointsEiiiiz(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  QPolygon QPolygon::translated(int dx, int dy);
extern void* C_ZNK8QPolygon10translatedEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QPolygon QPolygon::translated(const QPoint & offset);
extern void* C_ZNK8QPolygon10translatedERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  void QPolygon::translate(int dx, int dy);
extern void C_ZN8QPolygon9translateEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QPolygon::translate(const QPoint & offset);
extern void C_ZN8QPolygon9translateERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  void QPolygon::setPoints(int nPoints, int firstx, int firsty);
extern void C_ZN8QPolygon9setPointsEiiiz(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  void QPolygon::setPoints(int nPoints, const int * points);
extern void C_ZN8QPolygon9setPointsEiPKi(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QPolygon::~QPolygon();
extern void C_ZN8QPolygonD2Ev(void* qthis); // 2
  // proto:  QPolygonF QPolygonF::united(const QPolygonF & r);
extern void* C_ZNK9QPolygonF6unitedERKS_(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QPolygonF::subtracted(const QPolygonF & r);
extern void* C_ZNK9QPolygonF10subtractedERKS_(void* qthis, void* arg0); // 4
  // proto:  QRectF QPolygonF::boundingRect();
extern void* C_ZNK9QPolygonF12boundingRectEv(void* qthis); // 4
  // proto:  QPolygonF QPolygonF::intersected(const QPolygonF & r);
extern void* C_ZNK9QPolygonF11intersectedERKS_(void* qthis, void* arg0); // 4
  // proto:  void QPolygonF::~QPolygonF();
extern void C_ZN9QPolygonFD2Ev(void* qthis); // 2
  // proto:  void QPolygonF::QPolygonF(const QRectF & r);
extern void* C_ZN9QPolygonFC2ERK6QRectF(void* arg0); // 3
  // proto:  void QPolygonF::QPolygonF(const QPolygonF & a);
extern void* C_ZN9QPolygonFC2ERKS_(void* arg0); // 1
  // proto:  void QPolygonF::QPolygonF();
extern void* C_ZN9QPolygonFC2Ev(); // 1
  // proto:  void QPolygonF::QPolygonF(int size);
extern void* C_ZN9QPolygonFC2Ei(int32_t arg0); // 1
  // proto:  void QPolygonF::QPolygonF(const QPolygon & a);
extern void* C_ZN9QPolygonFC2ERK8QPolygon(void* arg0); // 3
  // proto:  bool QPolygonF::isClosed();
extern bool C_ZNK9QPolygonF8isClosedEv(void* qthis); // 2
  // proto:  void QPolygonF::swap(QPolygonF & other);
extern void C_ZN9QPolygonF4swapERS_(void* qthis, void* arg0); // 2
  // proto:  QPolygon QPolygonF::toPolygon();
extern void* C_ZNK9QPolygonF9toPolygonEv(void* qthis); // 4
  // proto:  QPolygonF QPolygonF::translated(qreal dx, qreal dy);
extern void* C_ZNK9QPolygonF10translatedEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  QPolygonF QPolygonF::translated(const QPointF & offset);
extern void* C_ZNK9QPolygonF10translatedERK7QPointF(void* qthis, void* arg0); // 4
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
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QPolygon)=1
type QPolygon struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QPolygonF)=1
type QPolygonF struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// united(const class QPolygon &)
func (this *QPolygon) United(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QPolygon).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK8QPolygon6unitedERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygon{}) // "QPolygon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPolygon", "united", args)
  }

  return
}

// subtracted(const class QPolygon &)
func (this *QPolygon) Subtracted(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QPolygon).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK8QPolygon10subtractedERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygon{}) // "QPolygon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPolygon", "subtracted", args)
  }

  return
}

// boundingRect()
func (this *QPolygon) Boundingrect(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QPolygon12boundingRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPolygon", "boundingRect", args)
  }

  return
}

// intersected(const class QPolygon &)
func (this *QPolygon) Intersected(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QPolygon).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK8QPolygon11intersectedERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygon{}) // "QPolygon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPolygon", "intersected", args)
  }

  return
}

// QPolygon(const class QRect &, _Bool)
func NewQPolygon(args ...interface{}) *QPolygon {
  // QPolygon(const class QRect &, _Bool)
  // QPolygon(const class QPolygon &)
  // QPolygon(int, const int *)
  // QPolygon(int)
  // QPolygon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
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
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QPolygonC2ERK5QRectb(arg0, arg1)
    return &QPolygon{Qclsinst:qthis}
  case 1:
    // invoke: _ZN8QPolygonC1ERKS_
    // invoke: void QPolygon(const class QPolygon &)
    var arg0 = args[0].(*QPolygon).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QPolygonC2ERKS_(arg0)
    return &QPolygon{Qclsinst:qthis}
  case 2:
    // invoke: _ZN8QPolygonC1EiPKi
    // invoke: void QPolygon(int, const int *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QPolygonC2EiPKi(arg0, arg1)
    return &QPolygon{Qclsinst:qthis}
  case 3:
    // invoke: _ZN8QPolygonC1Ei
    // invoke: void QPolygon(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QPolygonC2Ei(arg0)
    return &QPolygon{Qclsinst:qthis}
  case 4:
    // invoke: _ZN8QPolygonC1Ev
    // invoke: void QPolygon()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QPolygonC2Ev()
    return &QPolygon{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QPolygon", "QPolygon", args)
  }

  return nil // QPolygon{Qclsinst:qthis}
}

// setPoint(int, const class QPoint &)
func (this *QPolygon) Setpoint(args ...interface{}) () {
  // setPoint(int, const class QPoint &)
  // setPoint(int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QPolygon8setPointEiRK6QPoint(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QPolygon8setPointEiii
    // invoke: void setPoint(int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN8QPolygon8setPointEiii(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QPolygon", "setPoint", args)
  }

  return
}

// point(int, int *, int *)
func (this *QPolygon) Point(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    C.C_ZNK8QPolygon5pointEiPiS0_(this.Qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZNK8QPolygon5pointEi
    // invoke: QPoint point(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK8QPolygon5pointEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
  default:
    qtrt.ErrorResolve("QPolygon", "point", args)
  }

  return
}

// swap(class QPolygon &)
func (this *QPolygon) Swap(args ...interface{}) () {
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
    var arg0 = args[0].(*QPolygon).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPolygon4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPolygon", "swap", args)
  }

  return
}

// putPoints(int, int, const int *)
func (this *QPolygon) Putpoints(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    C.C_ZN8QPolygon9putPointsEiiPKi(this.Qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN8QPolygon9putPointsEiiRKS_i
    // invoke: void putPoints(int, int, const class QPolygon &, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QPolygon).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN8QPolygon9putPointsEiiRKS_i(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 2:
    // invoke: _ZN8QPolygon9putPointsEiiiiz
    // invoke: void putPoints(int, int, int, int, ...)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN8QPolygon9putPointsEiiiiz(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QPolygon", "putPoints", args)
  }

  return
}

// translated(int, int)
func (this *QPolygon) Translated(args ...interface{}) (ret interface{}) {
  // translated(int, int)
  // translated(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPolygon10translatedEii
    // invoke: QPolygon translated(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK8QPolygon10translatedEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygon{}) // "QPolygon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK8QPolygon10translatedERK6QPoint
    // invoke: QPolygon translated(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK8QPolygon10translatedERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygon{}) // "QPolygon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPolygon", "translated", args)
  }

  return
}

// translate(int, int)
func (this *QPolygon) Translate(args ...interface{}) () {
  // translate(int, int)
  // translate(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPolygon9translateEii
    // invoke: void translate(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPolygon9translateEii(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QPolygon9translateERK6QPoint
    // invoke: void translate(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPolygon9translateERK6QPoint(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPolygon", "translate", args)
  }

  return
}

// setPoints(int, int, int, ...)
func (this *QPolygon) Setpoints(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN8QPolygon9setPointsEiiiz(this.Qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN8QPolygon9setPointsEiPKi
    // invoke: void setPoints(int, const int *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPolygon9setPointsEiPKi(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPolygon", "setPoints", args)
  }

  return
}

// ~QPolygon()
func (this *QPolygon) Freeqpolygon(args ...interface{}) () {
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
    C.C_ZN8QPolygonD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPolygon", "~QPolygon", args)
  }

  return
}

// united(const class QPolygonF &)
func (this *QPolygonF) United(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QPolygonF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QPolygonF6unitedERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPolygonF", "united", args)
  }

  return
}

// subtracted(const class QPolygonF &)
func (this *QPolygonF) Subtracted(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QPolygonF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QPolygonF10subtractedERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPolygonF", "subtracted", args)
  }

  return
}

// boundingRect()
func (this *QPolygonF) Boundingrect(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QPolygonF12boundingRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPolygonF", "boundingRect", args)
  }

  return
}

// intersected(const class QPolygonF &)
func (this *QPolygonF) Intersected(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QPolygonF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QPolygonF11intersectedERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPolygonF", "intersected", args)
  }

  return
}

// ~QPolygonF()
func (this *QPolygonF) Freeqpolygonf(args ...interface{}) () {
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
    C.C_ZN9QPolygonFD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPolygonF", "~QPolygonF", args)
  }

  return
}

// QPolygonF(const class QRectF &)
func NewQPolygonF(args ...interface{}) *QPolygonF {
  // QPolygonF(const class QRectF &)
  // QPolygonF(const class QPolygonF &)
  // QPolygonF()
  // QPolygonF(int)
  // QPolygonF(const class QPolygon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
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
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QPolygonFC2ERK6QRectF(arg0)
    return &QPolygonF{Qclsinst:qthis}
  case 1:
    // invoke: _ZN9QPolygonFC1ERKS_
    // invoke: void QPolygonF(const class QPolygonF &)
    var arg0 = args[0].(*QPolygonF).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QPolygonFC2ERKS_(arg0)
    return &QPolygonF{Qclsinst:qthis}
  case 2:
    // invoke: _ZN9QPolygonFC1Ev
    // invoke: void QPolygonF()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QPolygonFC2Ev()
    return &QPolygonF{Qclsinst:qthis}
  case 3:
    // invoke: _ZN9QPolygonFC1Ei
    // invoke: void QPolygonF(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QPolygonFC2Ei(arg0)
    return &QPolygonF{Qclsinst:qthis}
  case 4:
    // invoke: _ZN9QPolygonFC1ERK8QPolygon
    // invoke: void QPolygonF(const class QPolygon &)
    var arg0 = args[0].(*QPolygon).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QPolygonFC2ERK8QPolygon(arg0)
    return &QPolygonF{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QPolygonF", "QPolygonF", args)
  }

  return nil // QPolygonF{Qclsinst:qthis}
}

// isClosed()
func (this *QPolygonF) Isclosed(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QPolygonF8isClosedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPolygonF", "isClosed", args)
  }

  return
}

// swap(class QPolygonF &)
func (this *QPolygonF) Swap(args ...interface{}) () {
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
    var arg0 = args[0].(*QPolygonF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QPolygonF4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPolygonF", "swap", args)
  }

  return
}

// toPolygon()
func (this *QPolygonF) Topolygon(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QPolygonF9toPolygonEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygon{}) // "QPolygon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPolygonF", "toPolygon", args)
  }

  return
}

// translated(qreal, qreal)
func (this *QPolygonF) Translated(args ...interface{}) (ret interface{}) {
  // translated(qreal, qreal)
  // translated(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QPolygonF10translatedEdd
    // invoke: QPolygonF translated(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK9QPolygonF10translatedEdd(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK9QPolygonF10translatedERK7QPointF
    // invoke: QPolygonF translated(const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QPolygonF10translatedERK7QPointF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPolygonF", "translated", args)
  }

  return
}

// translate(const class QPointF &)
func (this *QPolygonF) Translate(args ...interface{}) () {
  // translate(const class QPointF &)
  // translate(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QPolygonF9translateERK7QPointF
    // invoke: void translate(const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QPolygonF9translateERK7QPointF(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN9QPolygonF9translateEdd
    // invoke: void translate(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN9QPolygonF9translateEdd(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPolygonF", "translate", args)
  }

  return
}

// <= body block end

