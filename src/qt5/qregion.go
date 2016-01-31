package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtGui/qregion.h
// dst-file: /src/gui/qregion.go
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
  // proto:  QRect QRegion::boundingRect();
extern void C_ZNK7QRegion12boundingRectEv(void* qthis); // 4
  // proto:  void QRegion::QRegion(const QBitmap & bitmap);
extern void* C_ZN7QRegionC2ERK7QBitmap(void* arg0); // 3
  // proto:  void QRegion::QRegion(const QRegion & region);
extern void* C_ZN7QRegionC2ERKS_(void* arg0); // 3
  // proto:  void QRegion::QRegion();
extern void* C_ZN7QRegionC2Ev(); // 3
  // proto:  QRegion QRegion::xored(const QRegion & r);
extern void C_ZNK7QRegion5xoredERKS_(void* qthis, void* arg0); // 4
  // proto:  bool QRegion::intersects(const QRect & r);
extern void C_ZNK7QRegion10intersectsERK5QRect(void* qthis, void* arg0); // 4
  // proto:  bool QRegion::intersects(const QRegion & r);
extern void C_ZNK7QRegion10intersectsERKS_(void* qthis, void* arg0); // 4
  // proto:  QRegion QRegion::united(const QRegion & r);
extern void C_ZNK7QRegion6unitedERKS_(void* qthis, void* arg0); // 4
  // proto:  QRegion QRegion::united(const QRect & r);
extern void C_ZNK7QRegion6unitedERK5QRect(void* qthis, void* arg0); // 4
  // proto:  QRegion QRegion::intersected(const QRegion & r);
extern void C_ZNK7QRegion11intersectedERKS_(void* qthis, void* arg0); // 4
  // proto:  QRegion QRegion::intersected(const QRect & r);
extern void C_ZNK7QRegion11intersectedERK5QRect(void* qthis, void* arg0); // 4
  // proto:  bool QRegion::contains(const QPoint & p);
extern void C_ZNK7QRegion8containsERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  bool QRegion::contains(const QRect & r);
extern void C_ZNK7QRegion8containsERK5QRect(void* qthis, void* arg0); // 4
  // proto:  bool QRegion::isEmpty();
extern void C_ZNK7QRegion7isEmptyEv(void* qthis); // 4
  // proto:  void QRegion::swap(QRegion & other);
extern void C_ZN7QRegion4swapERS_(void* qthis, void* arg0); // 2
  // proto:  void QRegion::translate(int dx, int dy);
extern void C_ZN7QRegion9translateEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QRegion::translate(const QPoint & p);
extern void C_ZN7QRegion9translateERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  void QRegion::~QRegion();
extern void C_ZN7QRegionD2Ev(void* qthis); // 4
  // proto:  int QRegion::rectCount();
extern void C_ZNK7QRegion9rectCountEv(void* qthis); // 4
  // proto:  QVector<QRect> QRegion::rects();
extern void C_ZNK7QRegion5rectsEv(void* qthis); // 4
  // proto:  void QRegion::setRects(const QRect * rect, int num);
extern void C_ZN7QRegion8setRectsEPK5QRecti(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QRegion QRegion::subtracted(const QRegion & r);
extern void C_ZNK7QRegion10subtractedERKS_(void* qthis, void* arg0); // 4
  // proto:  bool QRegion::isNull();
extern void C_ZNK7QRegion6isNullEv(void* qthis); // 4
  // proto:  QRegion QRegion::translated(int dx, int dy);
extern void C_ZNK7QRegion10translatedEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QRegion QRegion::translated(const QPoint & p);
extern void C_ZNK7QRegion10translatedERK6QPoint(void* qthis, void* arg0); // 2
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

// class sizeof(QRegion)=8
type QRegion struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// boundingRect()
func (this *QRegion) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion12boundingRectEv
    // invoke: QRect boundingRect()
    var ret = C.C_ZNK7QRegion12boundingRectEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegion", "boundingRect", args)
  }

}

// QRegion(const class QBitmap &)
func NewQRegion(args ...interface{}) *QRegion {
  // QRegion(const class QBitmap &)
  // QRegion(const class QRegion &)
  // QRegion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBitmap{}) // "const QBitmap &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"
  vtys[2] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegionC1ERK7QBitmap
    // invoke: void QRegion(const class QBitmap &)
    var arg0 = args[0].(QBitmap).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QRegionC2ERK7QBitmap(arg0)
    return &QRegion{qclsinst:qthis}
  case 1:
    // invoke: _ZN7QRegionC1ERKS_
    // invoke: void QRegion(const class QRegion &)
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QRegionC2ERKS_(arg0)
    return &QRegion{qclsinst:qthis}
  case 2:
    // invoke: _ZN7QRegionC1Ev
    // invoke: void QRegion()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QRegionC2Ev()
    return &QRegion{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QRegion", "QRegion", args)
  }

  return nil // QRegion{qclsinst:qthis}
}

// xored(const class QRegion &)
func (this *QRegion) xored(args ...interface{}) () {
  // xored(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion5xoredERKS_
    // invoke: QRegion xored(const class QRegion &)
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK7QRegion5xoredERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegion", "xored", args)
  }

}

// intersects(const class QRect &)
func (this *QRegion) intersects(args ...interface{}) () {
  // intersects(const class QRect &)
  // intersects(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion10intersectsERK5QRect
    // invoke: bool intersects(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK7QRegion10intersectsERK5QRect(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK7QRegion10intersectsERKS_
    // invoke: bool intersects(const class QRegion &)
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK7QRegion10intersectsERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegion", "intersects", args)
  }

}

// united(const class QRegion &)
func (this *QRegion) united(args ...interface{}) () {
  // united(const class QRegion &)
  // united(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion6unitedERKS_
    // invoke: QRegion united(const class QRegion &)
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK7QRegion6unitedERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK7QRegion6unitedERK5QRect
    // invoke: QRegion united(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK7QRegion6unitedERK5QRect(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegion", "united", args)
  }

}

// intersected(const class QRegion &)
func (this *QRegion) intersected(args ...interface{}) () {
  // intersected(const class QRegion &)
  // intersected(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion11intersectedERKS_
    // invoke: QRegion intersected(const class QRegion &)
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK7QRegion11intersectedERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK7QRegion11intersectedERK5QRect
    // invoke: QRegion intersected(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK7QRegion11intersectedERK5QRect(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegion", "intersected", args)
  }

}

// contains(const class QPoint &)
func (this *QRegion) contains(args ...interface{}) () {
  // contains(const class QPoint &)
  // contains(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion8containsERK6QPoint
    // invoke: bool contains(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK7QRegion8containsERK6QPoint(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK7QRegion8containsERK5QRect
    // invoke: bool contains(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK7QRegion8containsERK5QRect(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegion", "contains", args)
  }

}

// isEmpty()
func (this *QRegion) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion7isEmptyEv
    // invoke: bool isEmpty()
    var ret = C.C_ZNK7QRegion7isEmptyEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegion", "isEmpty", args)
  }

}

// swap(class QRegion &)
func (this *QRegion) swap(args ...interface{}) () {
  // swap(class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegion4swapERS_
    // invoke: void swap(class QRegion &)
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QRegion4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegion", "swap", args)
  }

}

// translate(int, int)
func (this *QRegion) translate(args ...interface{}) () {
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
    // invoke: _ZN7QRegion9translateEii
    // invoke: void translate(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN7QRegion9translateEii(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN7QRegion9translateERK6QPoint
    // invoke: void translate(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QRegion9translateERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegion", "translate", args)
  }

}

// ~QRegion()
func (this *QRegion) FreeQRegion(args ...interface{}) () {
  // ~QRegion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegionD0Ev
    // invoke: void ~QRegion()
    C.C_ZN7QRegionD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegion", "~QRegion", args)
  }

}

// rectCount()
func (this *QRegion) rectCount(args ...interface{}) () {
  // rectCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion9rectCountEv
    // invoke: int rectCount()
    var ret = C.C_ZNK7QRegion9rectCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegion", "rectCount", args)
  }

}

// rects()
func (this *QRegion) rects(args ...interface{}) () {
  // rects()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion5rectsEv
    // invoke: QVector<QRect> rects()
    C.C_ZNK7QRegion5rectsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegion", "rects", args)
  }

}

// setRects(const class QRect *, int)
func (this *QRegion) setRects(args ...interface{}) () {
  // setRects(const class QRect *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegion8setRectsEPK5QRecti
    // invoke: void setRects(const class QRect *, int)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN7QRegion8setRectsEPK5QRecti(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QRegion", "setRects", args)
  }

}

// subtracted(const class QRegion &)
func (this *QRegion) subtracted(args ...interface{}) () {
  // subtracted(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion10subtractedERKS_
    // invoke: QRegion subtracted(const class QRegion &)
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK7QRegion10subtractedERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegion", "subtracted", args)
  }

}

// isNull()
func (this *QRegion) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion6isNullEv
    // invoke: bool isNull()
    var ret = C.C_ZNK7QRegion6isNullEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegion", "isNull", args)
  }

}

// translated(int, int)
func (this *QRegion) translated(args ...interface{}) () {
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
    // invoke: _ZNK7QRegion10translatedEii
    // invoke: QRegion translated(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK7QRegion10translatedEii(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK7QRegion10translatedERK6QPoint
    // invoke: QRegion translated(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK7QRegion10translatedERK6QPoint(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegion", "translated", args)
  }

}

// <= body block end

