package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
// src-file: /QtGui/qclipboard.h
// dst-file: /src/gui/qclipboard.go
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
  // proto:  void QClipboard::~QClipboard();
extern void _ZN10QClipboardD0Ev(void* qthis);
  // proto:  void QClipboard::QClipboard(QObject * parent);
extern void* dector_ZN10QClipboardC1EP7QObject(void* arg0);
extern void _ZN10QClipboardC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QClipboard::QClipboard(const QClipboard & );
extern void* dector_ZN10QClipboardC1ERKS_(void* arg0);
extern void _ZN10QClipboardC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QClipboard::supportsFindBuffer();
extern void _ZNK10QClipboard18supportsFindBufferEv(void* qthis);
  // proto:  bool QClipboard::ownsFindBuffer();
extern void _ZNK10QClipboard14ownsFindBufferEv(void* qthis);
  // proto:  bool QClipboard::ownsClipboard();
extern void _ZNK10QClipboard13ownsClipboardEv(void* qthis);
  // proto:  const QMetaObject * QClipboard::metaObject();
extern void _ZNK10QClipboard10metaObjectEv(void* qthis);
  // proto:  bool QClipboard::supportsSelection();
extern void _ZNK10QClipboard17supportsSelectionEv(void* qthis);
  // proto:  bool QClipboard::ownsSelection();
extern void _ZNK10QClipboard13ownsSelectionEv(void* qthis);
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

// class sizeof(QClipboard)=1
type QClipboard struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _changed QClipboard_changed_signal;
//  _findBufferChanged QClipboard_findBufferChanged_signal;
//  _selectionChanged QClipboard_selectionChanged_signal;
//  _dataChanged QClipboard_dataChanged_signal;
}

  // proto:  void QClipboard::~QClipboard();
func (this *QClipboard) FreeQClipboard(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QClipboard", "~QClipboard", args)
  }

}

  // proto:  void QClipboard::QClipboard(QObject * parent);
func NewQClipboard(args ...interface{}) QClipboard {
  return QClipboard{}
}

  // proto:  bool QClipboard::supportsFindBuffer();
func (this *QClipboard) supportsFindBuffer(args ...interface{}) () {
  // supportsFindBuffer()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QClipboard18supportsFindBufferEv
    // invoke: bool supportsFindBuffer()
    C._ZNK10QClipboard18supportsFindBufferEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QClipboard", "supportsFindBuffer", args)
  }

}

  // proto:  bool QClipboard::ownsFindBuffer();
func (this *QClipboard) ownsFindBuffer(args ...interface{}) () {
  // ownsFindBuffer()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QClipboard14ownsFindBufferEv
    // invoke: bool ownsFindBuffer()
    C._ZNK10QClipboard14ownsFindBufferEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QClipboard", "ownsFindBuffer", args)
  }

}

  // proto:  bool QClipboard::ownsClipboard();
func (this *QClipboard) ownsClipboard(args ...interface{}) () {
  // ownsClipboard()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QClipboard13ownsClipboardEv
    // invoke: bool ownsClipboard()
    C._ZNK10QClipboard13ownsClipboardEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QClipboard", "ownsClipboard", args)
  }

}

  // proto:  const QMetaObject * QClipboard::metaObject();
func (this *QClipboard) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QClipboard10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK10QClipboard10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QClipboard", "metaObject", args)
  }

}

  // proto:  bool QClipboard::supportsSelection();
func (this *QClipboard) supportsSelection(args ...interface{}) () {
  // supportsSelection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QClipboard17supportsSelectionEv
    // invoke: bool supportsSelection()
    C._ZNK10QClipboard17supportsSelectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QClipboard", "supportsSelection", args)
  }

}

  // proto:  bool QClipboard::ownsSelection();
func (this *QClipboard) ownsSelection(args ...interface{}) () {
  // ownsSelection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QClipboard13ownsSelectionEv
    // invoke: bool ownsSelection()
    C._ZNK10QClipboard13ownsSelectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QClipboard", "ownsSelection", args)
  }

}

// <= body block end

