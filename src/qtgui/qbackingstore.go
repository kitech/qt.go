package qtgui
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QBackingStore::QBackingStore(QWindow * window);
extern void* C_ZN13QBackingStoreC2EP7QWindow(void* arg0); // 3
  // proto:  QPlatformBackingStore * QBackingStore::handle();
extern void C_ZNK13QBackingStore6handleEv(void* qthis); // 4
  // proto:  QPaintDevice * QBackingStore::paintDevice();
extern void* C_ZN13QBackingStore11paintDeviceEv(void* qthis); // 4
  // proto:  QRegion QBackingStore::staticContents();
extern void* C_ZNK13QBackingStore14staticContentsEv(void* qthis); // 4
  // proto:  void QBackingStore::beginPaint(const QRegion & );
extern void C_ZN13QBackingStore10beginPaintERK7QRegion(void* qthis, void* arg0); // 4
  // proto:  void QBackingStore::endPaint();
extern void C_ZN13QBackingStore8endPaintEv(void* qthis); // 4
  // proto:  bool QBackingStore::hasStaticContents();
extern bool C_ZNK13QBackingStore17hasStaticContentsEv(void* qthis); // 4
  // proto:  QWindow * QBackingStore::window();
extern void* C_ZNK13QBackingStore6windowEv(void* qthis); // 4
  // proto:  void QBackingStore::setStaticContents(const QRegion & region);
extern void C_ZN13QBackingStore17setStaticContentsERK7QRegion(void* qthis, void* arg0); // 4
  // proto:  void QBackingStore::flush(const QRegion & region, QWindow * window, const QPoint & offset);
extern void C_ZN13QBackingStore5flushERK7QRegionP7QWindowRK6QPoint(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  QSize QBackingStore::size();
extern void* C_ZNK13QBackingStore4sizeEv(void* qthis); // 4
  // proto:  bool QBackingStore::scroll(const QRegion & area, int dx, int dy);
extern bool C_ZN13QBackingStore6scrollERK7QRegionii(void* qthis, void* arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  void QBackingStore::resize(const QSize & size);
extern void C_ZN13QBackingStore6resizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  void QBackingStore::~QBackingStore();
extern void C_ZN13QBackingStoreD2Ev(void* qthis); // 4
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

// class sizeof(QBackingStore)=1
type QBackingStore struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// QBackingStore(class QWindow *)
func NewQBackingStore(args ...interface{}) *QBackingStore {
  // QBackingStore(class QWindow *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWindow{}) // "QWindow *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QBackingStoreC1EP7QWindow
    // invoke: void QBackingStore(class QWindow *)
    var arg0 = args[0].(*QWindow).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QBackingStoreC2EP7QWindow(arg0)
    return &QBackingStore{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QBackingStore", "QBackingStore", args)
  }

  return nil // QBackingStore{Qclsinst:qthis}
}

// handle()
func (this *QBackingStore) Handle(args ...interface{}) () {
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
    C.C_ZNK13QBackingStore6handleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QBackingStore", "handle", args)
  }

  return
}

// paintDevice()
func (this *QBackingStore) Paintdevice(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN13QBackingStore11paintDeviceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPaintDevice{}) // "QPaintDevice *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBackingStore", "paintDevice", args)
  }

  return
}

// staticContents()
func (this *QBackingStore) Staticcontents(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QBackingStore14staticContentsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRegion{}) // "QRegion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBackingStore", "staticContents", args)
  }

  return
}

// beginPaint(const class QRegion &)
func (this *QBackingStore) Beginpaint(args ...interface{}) () {
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
    var arg0 = args[0].(*QRegion).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QBackingStore10beginPaintERK7QRegion(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBackingStore", "beginPaint", args)
  }

  return
}

// endPaint()
func (this *QBackingStore) Endpaint(args ...interface{}) () {
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
    C.C_ZN13QBackingStore8endPaintEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QBackingStore", "endPaint", args)
  }

  return
}

// hasStaticContents()
func (this *QBackingStore) Hasstaticcontents(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QBackingStore17hasStaticContentsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBackingStore", "hasStaticContents", args)
  }

  return
}

// window()
func (this *QBackingStore) Window(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QBackingStore6windowEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWindow{}) // "QWindow *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBackingStore", "window", args)
  }

  return
}

// setStaticContents(const class QRegion &)
func (this *QBackingStore) Setstaticcontents(args ...interface{}) () {
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
    var arg0 = args[0].(*QRegion).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QBackingStore17setStaticContentsERK7QRegion(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBackingStore", "setStaticContents", args)
  }

  return
}

// flush(const class QRegion &, class QWindow *, const class QPoint &)
func (this *QBackingStore) Flush(args ...interface{}) () {
  // flush(const class QRegion &, class QWindow *, const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"
  vtys[0][1] = reflect.TypeOf(QWindow{}) // "QWindow *"
  vtys[0][2] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QBackingStore5flushERK7QRegionP7QWindowRK6QPoint
    // invoke: void flush(const class QRegion &, class QWindow *, const class QPoint &)
    var arg0 = args[0].(*QRegion).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWindow).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN13QBackingStore5flushERK7QRegionP7QWindowRK6QPoint(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QBackingStore", "flush", args)
  }

  return
}

// size()
func (this *QBackingStore) Size(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QBackingStore4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBackingStore", "size", args)
  }

  return
}

// scroll(const class QRegion &, int, int)
func (this *QBackingStore) Scroll(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QRegion).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN13QBackingStore6scrollERK7QRegionii(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBackingStore", "scroll", args)
  }

  return
}

// resize(const class QSize &)
func (this *QBackingStore) Resize(args ...interface{}) () {
  // resize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QBackingStore6resizeERK5QSize
    // invoke: void resize(const class QSize &)
    var arg0 = args[0].(*qtcore.QSize).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QBackingStore6resizeERK5QSize(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBackingStore", "resize", args)
  }

  return
}

// ~QBackingStore()
func (this *QBackingStore) Freeqbackingstore(args ...interface{}) () {
  // ~QBackingStore()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QBackingStoreD0Ev
    // invoke: void ~QBackingStore()
    C.C_ZN13QBackingStoreD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QBackingStore", "~QBackingStore", args)
  }

  return
}

// <= body block end

