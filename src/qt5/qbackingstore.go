package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtGui/qbackingstore.h
// dst-file: /src/gui/qbackingstore.go
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
  // proto:  QPaintDevice * QBackingStore::paintDevice();
extern void _ZN13QBackingStore11paintDeviceEv(void* qthis);
  // proto:  QWindow * QBackingStore::window();
extern void _ZNK13QBackingStore6windowEv(void* qthis);
  // proto:  QSize QBackingStore::size();
extern void _ZNK13QBackingStore4sizeEv(void* qthis);
  // proto:  QRegion QBackingStore::staticContents();
extern void _ZNK13QBackingStore14staticContentsEv(void* qthis);
  // proto:  void QBackingStore::~QBackingStore();
extern void _ZN13QBackingStoreD0Ev(void* qthis);
  // proto:  void QBackingStore::setStaticContents(const QRegion & region);
extern void _ZN13QBackingStore17setStaticContentsERK7QRegion(void* qthis, void* arg0);
  // proto:  void QBackingStore::resize(const QSize & size);
extern void _ZN13QBackingStore6resizeERK5QSize(void* qthis, void* arg0);
  // proto:  void QBackingStore::flush(const QRegion & region, QWindow * window, const QPoint & offset);
extern void _ZN13QBackingStore5flushERK7QRegionP7QWindowRK6QPoint(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QBackingStore::beginPaint(const QRegion & );
extern void _ZN13QBackingStore10beginPaintERK7QRegion(void* qthis, void* arg0);
  // proto:  bool QBackingStore::hasStaticContents();
extern void _ZNK13QBackingStore17hasStaticContentsEv(void* qthis);
  // proto:  void QBackingStore::endPaint();
extern void _ZN13QBackingStore8endPaintEv(void* qthis);
  // proto:  bool QBackingStore::scroll(const QRegion & area, int dx, int dy);
extern void _ZN13QBackingStore6scrollERK7QRegionii(void* qthis, void* arg0, int arg1, int arg2);
  // proto:  QPlatformBackingStore * QBackingStore::handle();
extern void _ZNK13QBackingStore6handleEv(void* qthis);
  // proto:  void QBackingStore::QBackingStore(QWindow * window);
extern void* dector_ZN13QBackingStoreC1EP7QWindow(void* arg0);
extern void _ZN13QBackingStoreC1EP7QWindow(void* qthis, void* arg0);
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

// class sizeof(QBackingStore)=1
type QBackingStore struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QPaintDevice * QBackingStore::paintDevice();
func (this *QBackingStore) paintDevice(args ...interface{}) () {
  // paintDevice()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QBackingStore11paintDeviceEv
    // invoke: QPaintDevice * paintDevice()
    C._ZN13QBackingStore11paintDeviceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBackingStore", "paintDevice", args)
  }

}

  // proto:  QWindow * QBackingStore::window();
func (this *QBackingStore) window(args ...interface{}) () {
  // window()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QBackingStore6windowEv
    // invoke: QWindow * window()
    C._ZNK13QBackingStore6windowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBackingStore", "window", args)
  }

}

  // proto:  QSize QBackingStore::size();
func (this *QBackingStore) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QBackingStore4sizeEv
    // invoke: QSize size()
    C._ZNK13QBackingStore4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBackingStore", "size", args)
  }

}

  // proto:  QRegion QBackingStore::staticContents();
func (this *QBackingStore) staticContents(args ...interface{}) () {
  // staticContents()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QBackingStore14staticContentsEv
    // invoke: QRegion staticContents()
    C._ZNK13QBackingStore14staticContentsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBackingStore", "staticContents", args)
  }

}

  // proto:  void QBackingStore::~QBackingStore();
func (this *QBackingStore) FreeQBackingStore(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QBackingStore", "~QBackingStore", args)
  }

}

  // proto:  void QBackingStore::setStaticContents(const QRegion & region);
func (this *QBackingStore) setStaticContents(args ...interface{}) () {
  // setStaticContents(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QBackingStore17setStaticContentsERK7QRegion
    // invoke: void setStaticContents(const class QRegion &)
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QBackingStore17setStaticContentsERK7QRegion(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBackingStore", "setStaticContents", args)
  }

}

  // proto:  void QBackingStore::resize(const QSize & size);
func (this *QBackingStore) resize(args ...interface{}) () {
  // resize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QBackingStore6resizeERK5QSize
    // invoke: void resize(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QBackingStore6resizeERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBackingStore", "resize", args)
  }

}

  // proto:  void QBackingStore::flush(const QRegion & region, QWindow * window, const QPoint & offset);
func (this *QBackingStore) flush(args ...interface{}) () {
  // flush(const class QRegion &, class QWindow *, const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"
  vtys[0][1] = reflect.TypeOf(QWindow{}) // "QWindow *"
  vtys[0][2] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QBackingStore5flushERK7QRegionP7QWindowRK6QPoint
    // invoke: void flush(const class QRegion &, class QWindow *, const class QPoint &)
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWindow).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QPoint).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN13QBackingStore5flushERK7QRegionP7QWindowRK6QPoint(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QBackingStore", "flush", args)
  }

}

  // proto:  void QBackingStore::beginPaint(const QRegion & );
func (this *QBackingStore) beginPaint(args ...interface{}) () {
  // beginPaint(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QBackingStore10beginPaintERK7QRegion
    // invoke: void beginPaint(const class QRegion &)
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QBackingStore10beginPaintERK7QRegion(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBackingStore", "beginPaint", args)
  }

}

  // proto:  bool QBackingStore::hasStaticContents();
func (this *QBackingStore) hasStaticContents(args ...interface{}) () {
  // hasStaticContents()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QBackingStore17hasStaticContentsEv
    // invoke: bool hasStaticContents()
    C._ZNK13QBackingStore17hasStaticContentsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBackingStore", "hasStaticContents", args)
  }

}

  // proto:  void QBackingStore::endPaint();
func (this *QBackingStore) endPaint(args ...interface{}) () {
  // endPaint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QBackingStore8endPaintEv
    // invoke: void endPaint()
    C._ZN13QBackingStore8endPaintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBackingStore", "endPaint", args)
  }

}

  // proto:  bool QBackingStore::scroll(const QRegion & area, int dx, int dy);
func (this *QBackingStore) scroll(args ...interface{}) () {
  // scroll(const class QRegion &, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QBackingStore6scrollERK7QRegionii
    // invoke: bool scroll(const class QRegion &, int, int)
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN13QBackingStore6scrollERK7QRegionii(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QBackingStore", "scroll", args)
  }

}

  // proto:  QPlatformBackingStore * QBackingStore::handle();
func (this *QBackingStore) handle(args ...interface{}) () {
  // handle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QBackingStore6handleEv
    // invoke: QPlatformBackingStore * handle()
    C._ZNK13QBackingStore6handleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBackingStore", "handle", args)
  }

}

  // proto:  void QBackingStore::QBackingStore(QWindow * window);
func NewQBackingStore(args ...interface{}) QBackingStore {
  return QBackingStore{}
}

// <= body block end

