package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qkeysequenceedit.h
// dst-file: /src/widgets/qkeysequenceedit.go
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
// class sizeof(QKeySequenceEdit)=1
type QKeySequenceEdit struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
//  _editingFinished QKeySequenceEdit_editingFinished_signal;
//  _keySequenceChanged QKeySequenceEdit_keySequenceChanged_signal;
}


func NewQKeySequenceEdit(args ...interface{}) QKeySequenceEdit {
  return QKeySequenceEdit{}
}


func (this *QKeySequenceEdit) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QKeySequenceEdit5clearEv
  default:
    qtrt.ErrorResolve("QKeySequenceEdit", "clear", args)
  }

}


func (this *QKeySequenceEdit) setKeySequence(args ...interface{}) () {
  // setKeySequence(const class QKeySequence &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QKeySequence{}) // "const QKeySequence &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QKeySequenceEdit14setKeySequenceERK12QKeySequence
  default:
    qtrt.ErrorResolve("QKeySequenceEdit", "setKeySequence", args)
  }

}


func (this *QKeySequenceEdit) keySequence(args ...interface{}) () {
  // keySequence()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QKeySequenceEdit11keySequenceEv
  default:
    qtrt.ErrorResolve("QKeySequenceEdit", "keySequence", args)
  }

}


func (this *QKeySequenceEdit) FreeQKeySequenceEdit(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QKeySequenceEdit", "~QKeySequenceEdit", args)
  }

}


func (this *QKeySequenceEdit) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QKeySequenceEdit10metaObjectEv
  default:
    qtrt.ErrorResolve("QKeySequenceEdit", "metaObject", args)
  }

}

// <= body block end

