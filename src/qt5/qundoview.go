package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qundoview.h
// dst-file: /src/widgets/qundoview.go
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
// class sizeof(QUndoView)=1
type QUndoView struct {
  /*qbase*/ QListView;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQUndoView(args ...interface{}) QUndoView {
  return QUndoView{}
}


func (this *QUndoView) setStack(args ...interface{}) () {
  // setStack(class QUndoStack *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QUndoStack{}) // "QUndoStack *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QUndoView8setStackEP10QUndoStack
  default:
    qtrt.ErrorResolve("QUndoView", "setStack", args)
  }

}


func (this *QUndoView) setEmptyLabel(args ...interface{}) () {
  // setEmptyLabel(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QUndoView13setEmptyLabelERK7QString
  default:
    qtrt.ErrorResolve("QUndoView", "setEmptyLabel", args)
  }

}


func (this *QUndoView) setCleanIcon(args ...interface{}) () {
  // setCleanIcon(const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QUndoView12setCleanIconERK5QIcon
  default:
    qtrt.ErrorResolve("QUndoView", "setCleanIcon", args)
  }

}


func (this *QUndoView) setGroup(args ...interface{}) () {
  // setGroup(class QUndoGroup *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QUndoGroup{}) // "QUndoGroup *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QUndoView8setGroupEP10QUndoGroup
  default:
    qtrt.ErrorResolve("QUndoView", "setGroup", args)
  }

}


func (this *QUndoView) group(args ...interface{}) () {
  // group()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QUndoView5groupEv
  default:
    qtrt.ErrorResolve("QUndoView", "group", args)
  }

}


func (this *QUndoView) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QUndoView10metaObjectEv
  default:
    qtrt.ErrorResolve("QUndoView", "metaObject", args)
  }

}


func (this *QUndoView) stack(args ...interface{}) () {
  // stack()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QUndoView5stackEv
  default:
    qtrt.ErrorResolve("QUndoView", "stack", args)
  }

}


func (this *QUndoView) cleanIcon(args ...interface{}) () {
  // cleanIcon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QUndoView9cleanIconEv
  default:
    qtrt.ErrorResolve("QUndoView", "cleanIcon", args)
  }

}


func (this *QUndoView) emptyLabel(args ...interface{}) () {
  // emptyLabel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QUndoView10emptyLabelEv
  default:
    qtrt.ErrorResolve("QUndoView", "emptyLabel", args)
  }

}


func (this *QUndoView) FreeQUndoView(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QUndoView", "~QUndoView", args)
  }

}

// <= body block end

