package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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
extern void* C_ZNK7QRegion12boundingRectEv(void* qthis); // 4
  // proto:  void QRegion::QRegion(const QBitmap & bitmap);
extern void* C_ZN7QRegionC2ERK7QBitmap(void* arg0); // 3
  // proto:  void QRegion::QRegion(const QRegion & region);
extern void* C_ZN7QRegionC2ERKS_(void* arg0); // 3
  // proto:  void QRegion::QRegion();
extern void* C_ZN7QRegionC2Ev(); // 3
  // proto:  QRegion QRegion::xored(const QRegion & r);
extern void* C_ZNK7QRegion5xoredERKS_(void* qthis, void* arg0); // 4
  // proto:  bool QRegion::intersects(const QRect & r);
extern bool C_ZNK7QRegion10intersectsERK5QRect(void* qthis, void* arg0); // 4
  // proto:  bool QRegion::intersects(const QRegion & r);
extern bool C_ZNK7QRegion10intersectsERKS_(void* qthis, void* arg0); // 4
  // proto:  QRegion QRegion::united(const QRegion & r);
extern void* C_ZNK7QRegion6unitedERKS_(void* qthis, void* arg0); // 4
  // proto:  QRegion QRegion::united(const QRect & r);
extern void* C_ZNK7QRegion6unitedERK5QRect(void* qthis, void* arg0); // 4
  // proto:  QRegion QRegion::intersected(const QRegion & r);
extern void* C_ZNK7QRegion11intersectedERKS_(void* qthis, void* arg0); // 4
  // proto:  QRegion QRegion::intersected(const QRect & r);
extern void* C_ZNK7QRegion11intersectedERK5QRect(void* qthis, void* arg0); // 4
  // proto:  bool QRegion::contains(const QPoint & p);
extern bool C_ZNK7QRegion8containsERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  bool QRegion::contains(const QRect & r);
extern bool C_ZNK7QRegion8containsERK5QRect(void* qthis, void* arg0); // 4
  // proto:  bool QRegion::isEmpty();
extern bool C_ZNK7QRegion7isEmptyEv(void* qthis); // 4
  // proto:  void QRegion::swap(QRegion & other);
extern void C_ZN7QRegion4swapERS_(void* qthis, void* arg0); // 2
  // proto:  void QRegion::translate(int dx, int dy);
extern void C_ZN7QRegion9translateEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QRegion::translate(const QPoint & p);
extern void C_ZN7QRegion9translateERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  void QRegion::~QRegion();
extern void C_ZN7QRegionD2Ev(void* qthis); // 4
  // proto:  int QRegion::rectCount();
extern int32_t C_ZNK7QRegion9rectCountEv(void* qthis); // 4
  // proto:  QVector<QRect> QRegion::rects();
extern void C_ZNK7QRegion5rectsEv(void* qthis); // 4
  // proto:  void QRegion::setRects(const QRect * rect, int num);
extern void C_ZN7QRegion8setRectsEPK5QRecti(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QRegion QRegion::subtracted(const QRegion & r);
extern void* C_ZNK7QRegion10subtractedERKS_(void* qthis, void* arg0); // 4
  // proto:  bool QRegion::isNull();
extern bool C_ZNK7QRegion6isNullEv(void* qthis); // 4
  // proto:  QRegion QRegion::translated(int dx, int dy);
extern void* C_ZNK7QRegion10translatedEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QRegion QRegion::translated(const QPoint & p);
extern void* C_ZNK7QRegion10translatedERK6QPoint(void* qthis, void* arg0); // 2
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// boundingRect()
func (this *QRegion) Boundingrect(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QRegion12boundingRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRegion", "boundingRect", args)
  }

  return
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
    var arg0 = args[0].(*QBitmap).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QRegionC2ERK7QBitmap(arg0)
    return &QRegion{Qclsinst:qthis}
  case 1:
    // invoke: _ZN7QRegionC1ERKS_
    // invoke: void QRegion(const class QRegion &)
    var arg0 = args[0].(*QRegion).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QRegionC2ERKS_(arg0)
    return &QRegion{Qclsinst:qthis}
  case 2:
    // invoke: _ZN7QRegionC1Ev
    // invoke: void QRegion()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QRegionC2Ev()
    return &QRegion{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QRegion", "QRegion", args)
  }

  return nil // QRegion{Qclsinst:qthis}
}

// xored(const class QRegion &)
func (this *QRegion) Xored(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QRegion).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QRegion5xoredERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRegion{}) // "QRegion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRegion", "xored", args)
  }

  return
}

// intersects(const class QRect &)
func (this *QRegion) Intersects(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QRegion10intersectsERK5QRect(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK7QRegion10intersectsERKS_
    // invoke: bool intersects(const class QRegion &)
    var arg0 = args[0].(*QRegion).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QRegion10intersectsERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRegion", "intersects", args)
  }

  return
}

// united(const class QRegion &)
func (this *QRegion) United(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QRegion).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QRegion6unitedERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRegion{}) // "QRegion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK7QRegion6unitedERK5QRect
    // invoke: QRegion united(const class QRect &)
    var arg0 = args[0].(*QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QRegion6unitedERK5QRect(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRegion{}) // "QRegion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRegion", "united", args)
  }

  return
}

// intersected(const class QRegion &)
func (this *QRegion) Intersected(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QRegion).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QRegion11intersectedERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRegion{}) // "QRegion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK7QRegion11intersectedERK5QRect
    // invoke: QRegion intersected(const class QRect &)
    var arg0 = args[0].(*QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QRegion11intersectedERK5QRect(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRegion{}) // "QRegion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRegion", "intersected", args)
  }

  return
}

// contains(const class QPoint &)
func (this *QRegion) Contains(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QRegion8containsERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK7QRegion8containsERK5QRect
    // invoke: bool contains(const class QRect &)
    var arg0 = args[0].(*QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QRegion8containsERK5QRect(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRegion", "contains", args)
  }

  return
}

// isEmpty()
func (this *QRegion) Isempty(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QRegion7isEmptyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRegion", "isEmpty", args)
  }

  return
}

// swap(class QRegion &)
func (this *QRegion) Swap(args ...interface{}) () {
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
    var arg0 = args[0].(*QRegion).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QRegion4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegion", "swap", args)
  }

  return
}

// translate(int, int)
func (this *QRegion) Translate(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN7QRegion9translateEii(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN7QRegion9translateERK6QPoint
    // invoke: void translate(const class QPoint &)
    var arg0 = args[0].(*QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QRegion9translateERK6QPoint(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegion", "translate", args)
  }

  return
}

// ~QRegion()
func (this *QRegion) Freeqregion(args ...interface{}) () {
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
    C.C_ZN7QRegionD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QRegion", "~QRegion", args)
  }

  return
}

// rectCount()
func (this *QRegion) Rectcount(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QRegion9rectCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRegion", "rectCount", args)
  }

  return
}

// rects()
func (this *QRegion) Rects(args ...interface{}) () {
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
    C.C_ZNK7QRegion5rectsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QRegion", "rects", args)
  }

  return
}

// setRects(const class QRect *, int)
func (this *QRegion) Setrects(args ...interface{}) () {
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
    var arg0 = args[0].(*QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN7QRegion8setRectsEPK5QRecti(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QRegion", "setRects", args)
  }

  return
}

// subtracted(const class QRegion &)
func (this *QRegion) Subtracted(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QRegion).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QRegion10subtractedERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRegion{}) // "QRegion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRegion", "subtracted", args)
  }

  return
}

// isNull()
func (this *QRegion) Isnull(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QRegion6isNullEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRegion", "isNull", args)
  }

  return
}

// translated(int, int)
func (this *QRegion) Translated(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QRegion10translatedEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRegion{}) // "QRegion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK7QRegion10translatedERK6QPoint
    // invoke: QRegion translated(const class QPoint &)
    var arg0 = args[0].(*QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QRegion10translatedERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRegion{}) // "QRegion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRegion", "translated", args)
  }

  return
}

// <= body block end

