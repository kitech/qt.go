package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qclipboard.h
// dst-file: /src/gui/qclipboard.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
// class sizeof(QClipboard)=1
type QClipboard struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _changed QClipboard_changed_signal;
//  _findBufferChanged QClipboard_findBufferChanged_signal;
//  _selectionChanged QClipboard_selectionChanged_signal;
//  _dataChanged QClipboard_dataChanged_signal;
}


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


func NewQClipboard(args ...interface{})() {
}


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
  default:
    qtrt.ErrorResolve("QClipboard", "supportsFindBuffer", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QClipboard", "ownsFindBuffer", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QClipboard", "ownsClipboard", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QClipboard", "metaObject", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QClipboard", "supportsSelection", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QClipboard", "ownsSelection", args)
 }

}

// <= body block end

