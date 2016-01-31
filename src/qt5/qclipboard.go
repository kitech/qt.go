package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  bool QClipboard::ownsClipboard();
extern void C_ZNK10QClipboard13ownsClipboardEv(void* qthis); // 4
  // proto:  bool QClipboard::supportsFindBuffer();
extern void C_ZNK10QClipboard18supportsFindBufferEv(void* qthis); // 4
  // proto:  bool QClipboard::ownsFindBuffer();
extern void C_ZNK10QClipboard14ownsFindBufferEv(void* qthis); // 4
  // proto:  bool QClipboard::supportsSelection();
extern void C_ZNK10QClipboard17supportsSelectionEv(void* qthis); // 4
  // proto:  const QMetaObject * QClipboard::metaObject();
extern void C_ZNK10QClipboard10metaObjectEv(void* qthis); // 4
  // proto:  bool QClipboard::ownsSelection();
extern void C_ZNK10QClipboard13ownsSelectionEv(void* qthis); // 4
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

// ownsClipboard()
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
    C.C_ZNK10QClipboard13ownsClipboardEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QClipboard", "ownsClipboard", args)
  }

}

// supportsFindBuffer()
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
    C.C_ZNK10QClipboard18supportsFindBufferEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QClipboard", "supportsFindBuffer", args)
  }

}

// ownsFindBuffer()
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
    C.C_ZNK10QClipboard14ownsFindBufferEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QClipboard", "ownsFindBuffer", args)
  }

}

// supportsSelection()
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
    C.C_ZNK10QClipboard17supportsSelectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QClipboard", "supportsSelection", args)
  }

}

// metaObject()
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
    C.C_ZNK10QClipboard10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QClipboard", "metaObject", args)
  }

}

// ownsSelection()
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
    C.C_ZNK10QClipboard13ownsSelectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QClipboard", "ownsSelection", args)
  }

}

// <= body block end

