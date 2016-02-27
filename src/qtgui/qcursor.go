package qtgui
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtGui/qcursor.h
// dst-file: /src/gui/qcursor.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "runtime"
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
  // proto:  QPixmap QCursor::pixmap();
extern void* C_ZNK7QCursor6pixmapEv(void* qthis); // 4
  // proto: static void QCursor::setPos(QScreen * screen, int x, int y);
extern void C_ZN7QCursor6setPosEP7QScreenii(void* arg0, int32_t arg1, int32_t arg2); // 4
  // proto: static void QCursor::setPos(QScreen * screen, const QPoint & p);
extern void C_ZN7QCursor6setPosEP7QScreenRK6QPoint(void* arg0, void* arg1); // 2
  // proto: static void QCursor::setPos(int x, int y);
extern void C_ZN7QCursor6setPosEii(int32_t arg0, int32_t arg1); // 4
  // proto: static void QCursor::setPos(const QPoint & p);
extern void C_ZN7QCursor6setPosERK6QPoint(void* arg0); // 2
  // proto:  const QBitmap * QCursor::mask();
extern void* C_ZNK7QCursor4maskEv(void* qthis); // 4
  // proto: static QPoint QCursor::pos();
extern void* C_ZN7QCursor3posEv(); // 4
  // proto: static QPoint QCursor::pos(const QScreen * screen);
extern void* C_ZN7QCursor3posEPK7QScreen(void* arg0); // 4
  // proto:  QPoint QCursor::hotSpot();
extern void* C_ZNK7QCursor7hotSpotEv(void* qthis); // 4
  // proto:  const QBitmap * QCursor::bitmap();
extern void* C_ZNK7QCursor6bitmapEv(void* qthis); // 4
  // proto:  void QCursor::~QCursor();
extern void C_ZN7QCursorD2Ev(void* qthis); // 4
  // proto:  Qt::CursorShape QCursor::shape();
extern void C_ZNK7QCursor5shapeEv(void* qthis); // 4
  // proto:  void QCursor::QCursor(const QBitmap & bitmap, const QBitmap & mask, int hotX, int hotY);
extern void* C_ZN7QCursorC2ERK7QBitmapS2_ii(void* arg0, void* arg1, int32_t arg2, int32_t arg3); // 3
  // proto:  void QCursor::QCursor(const QPixmap & pixmap, int hotX, int hotY);
extern void* C_ZN7QCursorC2ERK7QPixmapii(void* arg0, int32_t arg1, int32_t arg2); // 3
  // proto:  void QCursor::QCursor(const QCursor & cursor);
extern void* C_ZN7QCursorC2ERKS_(void* arg0); // 3
  // proto:  void QCursor::QCursor();
extern void* C_ZN7QCursorC2Ev(); // 3
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
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QCursor)=8
type QCursor struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// pixmap()
func (this *QCursor) Pixmap(args ...interface{}) (ret interface{}) {
  // pixmap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QCursor6pixmapEv
    // invoke: QPixmap pixmap()
    var ret0 = C.C_ZNK7QCursor6pixmapEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPixmap{}) // "QPixmap"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCursor", "pixmap", args)
  }

  return
}

// setPos(class QScreen *, int, int)
func (this *QCursor) SetPos_s(args ...interface{}) () {
  // setPos(class QScreen *, int, int)
  // setPos(class QScreen *, const class QPoint &)
  // setPos(int, int)
  // setPos(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QScreen{}) // "QScreen *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QScreen{}) // "QScreen *"
  vtys[1][1] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QCursor6setPosEP7QScreenii
    // invoke: void setPos(class QScreen *, int, int)
    var arg0 = args[0].(*QScreen).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN7QCursor6setPosEP7QScreenii(arg0, arg1, arg2)
  case 1:
    // invoke: _ZN7QCursor6setPosEP7QScreenRK6QPoint
    // invoke: void setPos(class QScreen *, const class QPoint &)
    var arg0 = args[0].(*QScreen).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN7QCursor6setPosEP7QScreenRK6QPoint(arg0, arg1)
  case 2:
    // invoke: _ZN7QCursor6setPosEii
    // invoke: void setPos(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN7QCursor6setPosEii(arg0, arg1)
  case 3:
    // invoke: _ZN7QCursor6setPosERK6QPoint
    // invoke: void setPos(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QCursor6setPosERK6QPoint(arg0)
  default:
    qtrt.ErrorResolve("QCursor", "setPos", args)
  }

  return
}

// mask()
func (this *QCursor) Mask(args ...interface{}) (ret interface{}) {
  // mask()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QCursor4maskEv
    // invoke: const QBitmap * mask()
    var ret0 = C.C_ZNK7QCursor4maskEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBitmap{}) // "const QBitmap *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCursor", "mask", args)
  }

  return
}

// pos()
func (this *QCursor) Pos_s(args ...interface{}) (ret interface{}) {
  // pos()
  // pos(const class QScreen *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QScreen{}) // "const QScreen *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QCursor3posEv
    // invoke: QPoint pos()
    var ret0 = C.C_ZN7QCursor3posEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN7QCursor3posEPK7QScreen
    // invoke: QPoint pos(const class QScreen *)
    var arg0 = args[0].(*QScreen).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN7QCursor3posEPK7QScreen(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCursor", "pos", args)
  }

  return
}

// hotSpot()
func (this *QCursor) HotSpot(args ...interface{}) (ret interface{}) {
  // hotSpot()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QCursor7hotSpotEv
    // invoke: QPoint hotSpot()
    var ret0 = C.C_ZNK7QCursor7hotSpotEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCursor", "hotSpot", args)
  }

  return
}

// bitmap()
func (this *QCursor) Bitmap(args ...interface{}) (ret interface{}) {
  // bitmap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QCursor6bitmapEv
    // invoke: const QBitmap * bitmap()
    var ret0 = C.C_ZNK7QCursor6bitmapEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBitmap{}) // "const QBitmap *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCursor", "bitmap", args)
  }

  return
}

// ~QCursor()
func (this *QCursor) Free(args ...interface{}) () {
  // ~QCursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QCursorD0Ev
    // invoke: void ~QCursor()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN7QCursorD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QCursor", "~QCursor", args)
  }

  return
}

// shape()
func (this *QCursor) Shape(args ...interface{}) () {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QCursor5shapeEv
    // invoke: Qt::CursorShape shape()
    C.C_ZNK7QCursor5shapeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCursor", "shape", args)
  }

  return
}

// QCursor(const class QBitmap &, const class QBitmap &, int, int)
func GcfreeQCursor(this *QCursor) {
  qtrt.UniverseFree(this)
}
func NewQCursor(args ...interface{}) *QCursor {
  // QCursor(const class QBitmap &, const class QBitmap &, int, int)
  // QCursor(const class QPixmap &, int, int)
  // QCursor(const class QCursor &)
  // QCursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBitmap{}) // "const QBitmap &"
  vtys[0][1] = reflect.TypeOf(QBitmap{}) // "const QBitmap &"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QCursor{}) // "const QCursor &"
  vtys[3] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QCursorC1ERK7QBitmapS2_ii
    // invoke: void QCursor(const class QBitmap &, const class QBitmap &, int, int)
    var arg0 = args[0].(*QBitmap).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QBitmap).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QCursorC2ERK7QBitmapS2_ii(arg0, arg1, arg2, arg3)
    this := &QCursor{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQCursor)
    return this
  case 1:
    // invoke: _ZN7QCursorC1ERK7QPixmapii
    // invoke: void QCursor(const class QPixmap &, int, int)
    var arg0 = args[0].(*QPixmap).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QCursorC2ERK7QPixmapii(arg0, arg1, arg2)
    this := &QCursor{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQCursor)
    return this
  case 2:
    // invoke: _ZN7QCursorC1ERKS_
    // invoke: void QCursor(const class QCursor &)
    var arg0 = args[0].(*QCursor).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QCursorC2ERKS_(arg0)
    this := &QCursor{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQCursor)
    return this
  case 3:
    // invoke: _ZN7QCursorC1Ev
    // invoke: void QCursor()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QCursorC2Ev()
    this := &QCursor{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQCursor)
    return this
  default:
    qtrt.ErrorResolve("QCursor", "QCursor", args)
  }

  return nil // QCursor{Qclsinst:qthis}
}

// <= body block end

