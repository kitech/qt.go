package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qdialogbuttonbox.h
// dst-file: /src/widgets/qdialogbuttonbox.go
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
// class sizeof(QDialogButtonBox)=1
type QDialogButtonBox struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
//  _helpRequested QDialogButtonBox_helpRequested_signal;
//  _accepted QDialogButtonBox_accepted_signal;
//  _clicked QDialogButtonBox_clicked_signal;
//  _rejected QDialogButtonBox_rejected_signal;
}


func (this *QDialogButtonBox) buttons(args ...interface{}) () {
  // buttons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QDialogButtonBox7buttonsEv
  default:
    qtrt.ErrorResolve("QDialogButtonBox", "buttons", args)
  }

}


func (this *QDialogButtonBox) setCenterButtons(args ...interface{}) () {
  // setCenterButtons(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QDialogButtonBox16setCenterButtonsEb
  default:
    qtrt.ErrorResolve("QDialogButtonBox", "setCenterButtons", args)
  }

}


func (this *QDialogButtonBox) centerButtons(args ...interface{}) () {
  // centerButtons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QDialogButtonBox13centerButtonsEv
  default:
    qtrt.ErrorResolve("QDialogButtonBox", "centerButtons", args)
  }

}


func (this *QDialogButtonBox) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QDialogButtonBox10metaObjectEv
  default:
    qtrt.ErrorResolve("QDialogButtonBox", "metaObject", args)
  }

}


func (this *QDialogButtonBox) removeButton(args ...interface{}) () {
  // removeButton(class QAbstractButton *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractButton{}) // "QAbstractButton *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QDialogButtonBox12removeButtonEP15QAbstractButton
  default:
    qtrt.ErrorResolve("QDialogButtonBox", "removeButton", args)
  }

}


func (this *QDialogButtonBox) FreeQDialogButtonBox(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDialogButtonBox", "~QDialogButtonBox", args)
  }

}


func NewQDialogButtonBox(args ...interface{}) QDialogButtonBox {
  return QDialogButtonBox{}
}


func (this *QDialogButtonBox) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QDialogButtonBox5clearEv
  default:
    qtrt.ErrorResolve("QDialogButtonBox", "clear", args)
  }

}

// <= body block end

