package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
  // proto: static void QCursor::setPos(QScreen * screen, int x, int y);
extern void _ZN7QCursor6setPosEP7QScreenii(void* arg0, int arg1, int arg2);
  // proto:  QPixmap QCursor::pixmap();
extern void _ZNK7QCursor6pixmapEv(void* qthis);
  // proto:  void QCursor::QCursor(const QBitmap & bitmap, const QBitmap & mask, int hotX, int hotY);
extern void* dector_ZN7QCursorC1ERK7QBitmapS2_ii(void* arg0, void* arg1, int arg2, int arg3);
extern void _ZN7QCursorC1ERK7QBitmapS2_ii(void* qthis, void* arg0, void* arg1, int arg2, int arg3);
  // proto:  void QCursor::QCursor(const QPixmap & pixmap, int hotX, int hotY);
extern void* dector_ZN7QCursorC1ERK7QPixmapii(void* arg0, int arg1, int arg2);
extern void _ZN7QCursorC1ERK7QPixmapii(void* qthis, void* arg0, int arg1, int arg2);
  // proto:  void QCursor::~QCursor();
extern void _ZN7QCursorD0Ev(void* qthis);
  // proto:  const QBitmap * QCursor::mask();
extern void _ZNK7QCursor4maskEv(void* qthis);
  // proto:  void QCursor::QCursor(const QCursor & cursor);
extern void* dector_ZN7QCursorC1ERKS_(void* arg0);
extern void _ZN7QCursorC1ERKS_(void* qthis, void* arg0);
  // proto: static void QCursor::setPos(int x, int y);
extern void _ZN7QCursor6setPosEii(int arg0, int arg1);
  // proto: static void QCursor::setPos(QScreen * screen, const QPoint & p);
extern void demth_ZN7QCursor6setPosEP7QScreenRK6QPoint(void* arg0, void* arg1);
  // proto: static void QCursor::setPos(const QPoint & p);
extern void demth_ZN7QCursor6setPosERK6QPoint(void* arg0);
  // proto:  const QBitmap * QCursor::bitmap();
extern void _ZNK7QCursor6bitmapEv(void* qthis);
  // proto: static QPoint QCursor::pos(const QScreen * screen);
extern void _ZN7QCursor3posEPK7QScreen(void* arg0);
  // proto: static QPoint QCursor::pos();
extern void _ZN7QCursor3posEv();
  // proto:  void QCursor::QCursor();
extern void* dector_ZN7QCursorC1Ev();
extern void _ZN7QCursorC1Ev(void* qthis);
  // proto:  QPoint QCursor::hotSpot();
extern void _ZNK7QCursor7hotSpotEv(void* qthis);
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

// class sizeof(QCursor)=8
type QCursor struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto: static void QCursor::setPos(QScreen * screen, int x, int y);
func (this *QCursor) setPos_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCursor", "setPos", args)
  }

}

  // proto:  QPixmap QCursor::pixmap();
func (this *QCursor) pixmap(args ...interface{}) () {
  // pixmap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QCursor6pixmapEv
  default:
    qtrt.ErrorResolve("QCursor", "pixmap", args)
  }

}

  // proto:  void QCursor::QCursor(const QBitmap & bitmap, const QBitmap & mask, int hotX, int hotY);
func NewQCursor(args ...interface{}) QCursor {
  return QCursor{}
}

  // proto:  void QCursor::~QCursor();
func (this *QCursor) FreeQCursor(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCursor", "~QCursor", args)
  }

}

  // proto:  const QBitmap * QCursor::mask();
func (this *QCursor) mask(args ...interface{}) () {
  // mask()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QCursor4maskEv
  default:
    qtrt.ErrorResolve("QCursor", "mask", args)
  }

}

  // proto:  const QBitmap * QCursor::bitmap();
func (this *QCursor) bitmap(args ...interface{}) () {
  // bitmap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QCursor6bitmapEv
  default:
    qtrt.ErrorResolve("QCursor", "bitmap", args)
  }

}

  // proto: static QPoint QCursor::pos(const QScreen * screen);
func (this *QCursor) pos_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCursor", "pos", args)
  }

}

  // proto:  QPoint QCursor::hotSpot();
func (this *QCursor) hotSpot(args ...interface{}) () {
  // hotSpot()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QCursor7hotSpotEv
  default:
    qtrt.ErrorResolve("QCursor", "hotSpot", args)
  }

}

// <= body block end

