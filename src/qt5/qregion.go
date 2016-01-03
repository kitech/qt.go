package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
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
  // proto:  bool QRegion::isNull();
extern void _ZNK7QRegion6isNullEv(void* qthis);
  // proto:  QRect QRegion::boundingRect();
extern void _ZNK7QRegion12boundingRectEv(void* qthis);
  // proto:  void QRegion::QRegion(const QRegion & region);
extern void* dector_ZN7QRegionC1ERKS_(void* arg0);
extern void _ZN7QRegionC1ERKS_(void* qthis, void* arg0);
  // proto:  int QRegion::rectCount();
extern void _ZNK7QRegion9rectCountEv(void* qthis);
  // proto:  void QRegion::translate(int dx, int dy);
extern void _ZN7QRegion9translateEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  QRegion QRegion::united(const QRegion & r);
extern void _ZNK7QRegion6unitedERKS_(void* qthis, void* arg0);
  // proto:  QRegion QRegion::translated(const QPoint & p);
extern void demth_ZNK7QRegion10translatedERK6QPoint(void* qthis, void* arg0);
  // proto:  void QRegion::swap(QRegion & other);
extern void demth_ZN7QRegion4swapERS_(void* qthis, void* arg0);
  // proto:  void QRegion::QRegion(const QBitmap & bitmap);
extern void* dector_ZN7QRegionC1ERK7QBitmap(void* arg0);
extern void _ZN7QRegionC1ERK7QBitmap(void* qthis, void* arg0);
  // proto:  void QRegion::~QRegion();
extern void _ZN7QRegionD0Ev(void* qthis);
  // proto:  void QRegion::translate(const QPoint & p);
extern void demth_ZN7QRegion9translateERK6QPoint(void* qthis, void* arg0);
  // proto:  void QRegion::QRegion();
extern void* dector_ZN7QRegionC1Ev();
extern void _ZN7QRegionC1Ev(void* qthis);
  // proto:  bool QRegion::contains(const QRect & r);
extern void _ZNK7QRegion8containsERK5QRect(void* qthis, void* arg0);
  // proto:  bool QRegion::isEmpty();
extern void _ZNK7QRegion7isEmptyEv(void* qthis);
  // proto:  QRegion QRegion::intersected(const QRect & r);
extern void _ZNK7QRegion11intersectedERK5QRect(void* qthis, void* arg0);
  // proto:  void QRegion::setRects(const QRect * rect, int num);
extern void _ZN7QRegion8setRectsEPK5QRecti(void* qthis, void* arg0, int32_t arg1);
  // proto:  QVector<QRect> QRegion::rects();
extern void _ZNK7QRegion5rectsEv(void* qthis);
  // proto:  QRegion QRegion::subtracted(const QRegion & r);
extern void _ZNK7QRegion10subtractedERKS_(void* qthis, void* arg0);
  // proto:  bool QRegion::intersects(const QRect & r);
extern void _ZNK7QRegion10intersectsERK5QRect(void* qthis, void* arg0);
  // proto:  QRegion QRegion::translated(int dx, int dy);
extern void _ZNK7QRegion10translatedEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  QRegion QRegion::intersected(const QRegion & r);
extern void _ZNK7QRegion11intersectedERKS_(void* qthis, void* arg0);
  // proto:  QRegion QRegion::united(const QRect & r);
extern void _ZNK7QRegion6unitedERK5QRect(void* qthis, void* arg0);
  // proto:  QRegion QRegion::xored(const QRegion & r);
extern void _ZNK7QRegion5xoredERKS_(void* qthis, void* arg0);
  // proto:  bool QRegion::intersects(const QRegion & r);
extern void _ZNK7QRegion10intersectsERKS_(void* qthis, void* arg0);
  // proto:  bool QRegion::contains(const QPoint & p);
extern void _ZNK7QRegion8containsERK6QPoint(void* qthis, void* arg0);
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

  // proto:  bool QRegion::isNull();
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
    C._ZNK7QRegion6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegion", "isNull", args)
  }

}

  // proto:  QRect QRegion::boundingRect();
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
    C._ZNK7QRegion12boundingRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegion", "boundingRect", args)
  }

}

  // proto:  void QRegion::QRegion(const QRegion & region);
func NewQRegion(args ...interface{}) QRegion {
  return QRegion{}
}

  // proto:  int QRegion::rectCount();
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
    C._ZNK7QRegion9rectCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegion", "rectCount", args)
  }

}

  // proto:  void QRegion::translate(int dx, int dy);
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
    C._ZN7QRegion9translateEii(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN7QRegion9translateERK6QPoint
    // invoke: void translate(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN7QRegion9translateERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegion", "translate", args)
  }

}

  // proto:  QRegion QRegion::united(const QRegion & r);
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
    C._ZNK7QRegion6unitedERKS_(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK7QRegion6unitedERK5QRect
    // invoke: QRegion united(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QRegion6unitedERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegion", "united", args)
  }

}

  // proto:  QRegion QRegion::translated(const QPoint & p);
func (this *QRegion) translated(args ...interface{}) () {
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
    // invoke: _ZNK7QRegion10translatedERK6QPoint
    // invoke: QRegion translated(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZNK7QRegion10translatedERK6QPoint(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK7QRegion10translatedEii
    // invoke: QRegion translated(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK7QRegion10translatedEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QRegion", "translated", args)
  }

}

  // proto:  void QRegion::swap(QRegion & other);
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
    C.demth_ZN7QRegion4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegion", "swap", args)
  }

}

  // proto:  void QRegion::~QRegion();
func (this *QRegion) FreeQRegion(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QRegion", "~QRegion", args)
  }

}

  // proto:  bool QRegion::contains(const QRect & r);
func (this *QRegion) contains(args ...interface{}) () {
  // contains(const class QRect &)
  // contains(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion8containsERK5QRect
    // invoke: bool contains(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QRegion8containsERK5QRect(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK7QRegion8containsERK6QPoint
    // invoke: bool contains(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QRegion8containsERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegion", "contains", args)
  }

}

  // proto:  bool QRegion::isEmpty();
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
    C._ZNK7QRegion7isEmptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegion", "isEmpty", args)
  }

}

  // proto:  QRegion QRegion::intersected(const QRect & r);
func (this *QRegion) intersected(args ...interface{}) () {
  // intersected(const class QRect &)
  // intersected(const class QRegion &)
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
    // invoke: _ZNK7QRegion11intersectedERK5QRect
    // invoke: QRegion intersected(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QRegion11intersectedERK5QRect(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK7QRegion11intersectedERKS_
    // invoke: QRegion intersected(const class QRegion &)
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QRegion11intersectedERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegion", "intersected", args)
  }

}

  // proto:  void QRegion::setRects(const QRect * rect, int num);
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
    C._ZN7QRegion8setRectsEPK5QRecti(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QRegion", "setRects", args)
  }

}

  // proto:  QVector<QRect> QRegion::rects();
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
    C._ZNK7QRegion5rectsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegion", "rects", args)
  }

}

  // proto:  QRegion QRegion::subtracted(const QRegion & r);
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
    C._ZNK7QRegion10subtractedERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegion", "subtracted", args)
  }

}

  // proto:  bool QRegion::intersects(const QRect & r);
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
    C._ZNK7QRegion10intersectsERK5QRect(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK7QRegion10intersectsERKS_
    // invoke: bool intersects(const class QRegion &)
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QRegion10intersectsERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegion", "intersects", args)
  }

}

  // proto:  QRegion QRegion::xored(const QRegion & r);
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
    C._ZNK7QRegion5xoredERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegion", "xored", args)
  }

}

// <= body block end

