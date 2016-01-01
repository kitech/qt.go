package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qdialog.h
// dst-file: /src/widgets/qdialog.go
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
// class sizeof(QDialog)=1
type QDialog struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
//  _finished QDialog_finished_signal;
//  _accepted QDialog_accepted_signal;
//  _rejected QDialog_rejected_signal;
}


func (this *QDialog) setExtension(args ...interface{}) () {
  // setExtension(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QDialog12setExtensionEP7QWidget
  default:
    qtrt.ErrorResolve("QDialog", "setExtension", args)
  }

}


func (this *QDialog) result(args ...interface{}) () {
  // result()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QDialog6resultEv
  default:
    qtrt.ErrorResolve("QDialog", "result", args)
  }

}


func (this *QDialog) done(args ...interface{}) () {
  // done(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QDialog4doneEi
  default:
    qtrt.ErrorResolve("QDialog", "done", args)
  }

}


func (this *QDialog) open(args ...interface{}) () {
  // open()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QDialog4openEv
  default:
    qtrt.ErrorResolve("QDialog", "open", args)
  }

}


func (this *QDialog) FreeQDialog(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDialog", "~QDialog", args)
  }

}


func (this *QDialog) setResult(args ...interface{}) () {
  // setResult(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QDialog9setResultEi
  default:
    qtrt.ErrorResolve("QDialog", "setResult", args)
  }

}


func (this *QDialog) setSizeGripEnabled(args ...interface{}) () {
  // setSizeGripEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QDialog18setSizeGripEnabledEb
  default:
    qtrt.ErrorResolve("QDialog", "setSizeGripEnabled", args)
  }

}


func (this *QDialog) showExtension(args ...interface{}) () {
  // showExtension(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QDialog13showExtensionEb
  default:
    qtrt.ErrorResolve("QDialog", "showExtension", args)
  }

}


func (this *QDialog) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QDialog10metaObjectEv
  default:
    qtrt.ErrorResolve("QDialog", "metaObject", args)
  }

}


func (this *QDialog) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QDialog15minimumSizeHintEv
  default:
    qtrt.ErrorResolve("QDialog", "minimumSizeHint", args)
  }

}


func (this *QDialog) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QDialog8sizeHintEv
  default:
    qtrt.ErrorResolve("QDialog", "sizeHint", args)
  }

}


func (this *QDialog) accept(args ...interface{}) () {
  // accept()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QDialog6acceptEv
  default:
    qtrt.ErrorResolve("QDialog", "accept", args)
  }

}


func (this *QDialog) setVisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QDialog10setVisibleEb
  default:
    qtrt.ErrorResolve("QDialog", "setVisible", args)
  }

}


func (this *QDialog) extension(args ...interface{}) () {
  // extension()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QDialog9extensionEv
  default:
    qtrt.ErrorResolve("QDialog", "extension", args)
  }

}


func (this *QDialog) exec(args ...interface{}) () {
  // exec()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QDialog4execEv
  default:
    qtrt.ErrorResolve("QDialog", "exec", args)
  }

}


func (this *QDialog) reject(args ...interface{}) () {
  // reject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QDialog6rejectEv
  default:
    qtrt.ErrorResolve("QDialog", "reject", args)
  }

}


func NewQDialog(args ...interface{}) QDialog {
  return QDialog{}
}


func (this *QDialog) isSizeGripEnabled(args ...interface{}) () {
  // isSizeGripEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QDialog17isSizeGripEnabledEv
  default:
    qtrt.ErrorResolve("QDialog", "isSizeGripEnabled", args)
  }

}


func (this *QDialog) setModal(args ...interface{}) () {
  // setModal(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QDialog8setModalEb
  default:
    qtrt.ErrorResolve("QDialog", "setModal", args)
  }

}

// <= body block end

