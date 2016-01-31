package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
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
extern bool C_ZNK10QClipboard13ownsClipboardEv(void* qthis); // 4
  // proto:  bool QClipboard::supportsFindBuffer();
extern bool C_ZNK10QClipboard18supportsFindBufferEv(void* qthis); // 4
  // proto:  bool QClipboard::ownsFindBuffer();
extern bool C_ZNK10QClipboard14ownsFindBufferEv(void* qthis); // 4
  // proto:  bool QClipboard::supportsSelection();
extern bool C_ZNK10QClipboard17supportsSelectionEv(void* qthis); // 4
  // proto:  const QMetaObject * QClipboard::metaObject();
extern void C_ZNK10QClipboard10metaObjectEv(void* qthis); // 4
  // proto:  bool QClipboard::ownsSelection();
extern bool C_ZNK10QClipboard13ownsSelectionEv(void* qthis); // 4
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
func (this *QClipboard) Ownsclipboard(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QClipboard13ownsClipboardEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QClipboard", "ownsClipboard", args)
  }

  return
}

// supportsFindBuffer()
func (this *QClipboard) Supportsfindbuffer(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QClipboard18supportsFindBufferEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QClipboard", "supportsFindBuffer", args)
  }

  return
}

// ownsFindBuffer()
func (this *QClipboard) Ownsfindbuffer(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QClipboard14ownsFindBufferEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QClipboard", "ownsFindBuffer", args)
  }

  return
}

// supportsSelection()
func (this *QClipboard) Supportsselection(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QClipboard17supportsSelectionEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QClipboard", "supportsSelection", args)
  }

  return
}

// metaObject()
func (this *QClipboard) Metaobject(args ...interface{}) () {
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

  return
}

// ownsSelection()
func (this *QClipboard) Ownsselection(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QClipboard13ownsSelectionEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QClipboard", "ownsSelection", args)
  }

  return
}

// <= body block end

