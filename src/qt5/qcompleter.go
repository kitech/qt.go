package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qcompleter.h
// dst-file: /src/widgets/qcompleter.go
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
// class sizeof(QCompleter)=1
type QCompleter struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _highlighted QCompleter_highlighted_signal;
//  _activated QCompleter_activated_signal;
}


func NewQCompleter(args ...interface{}) QCompleter {
  return QCompleter{}
}


func (this *QCompleter) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter10metaObjectEv
  default:
    qtrt.ErrorResolve("QCompleter", "metaObject", args)
  }

}


func (this *QCompleter) popup(args ...interface{}) () {
  // popup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter5popupEv
  default:
    qtrt.ErrorResolve("QCompleter", "popup", args)
  }

}


func (this *QCompleter) complete(args ...interface{}) () {
  // complete(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QCompleter8completeERK5QRect
  default:
    qtrt.ErrorResolve("QCompleter", "complete", args)
  }

}


func (this *QCompleter) setCompletionRole(args ...interface{}) () {
  // setCompletionRole(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QCompleter17setCompletionRoleEi
  default:
    qtrt.ErrorResolve("QCompleter", "setCompletionRole", args)
  }

}


func (this *QCompleter) completionCount(args ...interface{}) () {
  // completionCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter15completionCountEv
  default:
    qtrt.ErrorResolve("QCompleter", "completionCount", args)
  }

}


func (this *QCompleter) currentIndex(args ...interface{}) () {
  // currentIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter12currentIndexEv
  default:
    qtrt.ErrorResolve("QCompleter", "currentIndex", args)
  }

}


func (this *QCompleter) pathFromIndex(args ...interface{}) () {
  // pathFromIndex(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter13pathFromIndexERK11QModelIndex
  default:
    qtrt.ErrorResolve("QCompleter", "pathFromIndex", args)
  }

}


func (this *QCompleter) setMaxVisibleItems(args ...interface{}) () {
  // setMaxVisibleItems(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QCompleter18setMaxVisibleItemsEi
  default:
    qtrt.ErrorResolve("QCompleter", "setMaxVisibleItems", args)
  }

}


func (this *QCompleter) completionColumn(args ...interface{}) () {
  // completionColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter16completionColumnEv
  default:
    qtrt.ErrorResolve("QCompleter", "completionColumn", args)
  }

}


func (this *QCompleter) maxVisibleItems(args ...interface{}) () {
  // maxVisibleItems()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter15maxVisibleItemsEv
  default:
    qtrt.ErrorResolve("QCompleter", "maxVisibleItems", args)
  }

}


func (this *QCompleter) FreeQCompleter(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCompleter", "~QCompleter", args)
  }

}


func (this *QCompleter) setWrapAround(args ...interface{}) () {
  // setWrapAround(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QCompleter13setWrapAroundEb
  default:
    qtrt.ErrorResolve("QCompleter", "setWrapAround", args)
  }

}


func (this *QCompleter) splitPath(args ...interface{}) () {
  // splitPath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter9splitPathERK7QString
  default:
    qtrt.ErrorResolve("QCompleter", "splitPath", args)
  }

}


func (this *QCompleter) model(args ...interface{}) () {
  // model()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter5modelEv
  default:
    qtrt.ErrorResolve("QCompleter", "model", args)
  }

}


func (this *QCompleter) currentCompletion(args ...interface{}) () {
  // currentCompletion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter17currentCompletionEv
  default:
    qtrt.ErrorResolve("QCompleter", "currentCompletion", args)
  }

}


func (this *QCompleter) setCompletionColumn(args ...interface{}) () {
  // setCompletionColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QCompleter19setCompletionColumnEi
  default:
    qtrt.ErrorResolve("QCompleter", "setCompletionColumn", args)
  }

}


func (this *QCompleter) setCompletionPrefix(args ...interface{}) () {
  // setCompletionPrefix(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QCompleter19setCompletionPrefixERK7QString
  default:
    qtrt.ErrorResolve("QCompleter", "setCompletionPrefix", args)
  }

}


func (this *QCompleter) completionModel(args ...interface{}) () {
  // completionModel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter15completionModelEv
  default:
    qtrt.ErrorResolve("QCompleter", "completionModel", args)
  }

}


func (this *QCompleter) setCurrentRow(args ...interface{}) () {
  // setCurrentRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QCompleter13setCurrentRowEi
  default:
    qtrt.ErrorResolve("QCompleter", "setCurrentRow", args)
  }

}


func (this *QCompleter) currentRow(args ...interface{}) () {
  // currentRow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter10currentRowEv
  default:
    qtrt.ErrorResolve("QCompleter", "currentRow", args)
  }

}


func (this *QCompleter) setModel(args ...interface{}) () {
  // setModel(class QAbstractItemModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemModel{}) // "QAbstractItemModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QCompleter8setModelEP18QAbstractItemModel
  default:
    qtrt.ErrorResolve("QCompleter", "setModel", args)
  }

}


func (this *QCompleter) wrapAround(args ...interface{}) () {
  // wrapAround()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter10wrapAroundEv
  default:
    qtrt.ErrorResolve("QCompleter", "wrapAround", args)
  }

}


func (this *QCompleter) setPopup(args ...interface{}) () {
  // setPopup(class QAbstractItemView *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemView{}) // "QAbstractItemView *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QCompleter8setPopupEP17QAbstractItemView
  default:
    qtrt.ErrorResolve("QCompleter", "setPopup", args)
  }

}


func (this *QCompleter) widget(args ...interface{}) () {
  // widget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter6widgetEv
  default:
    qtrt.ErrorResolve("QCompleter", "widget", args)
  }

}


func (this *QCompleter) completionRole(args ...interface{}) () {
  // completionRole()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter14completionRoleEv
  default:
    qtrt.ErrorResolve("QCompleter", "completionRole", args)
  }

}


func (this *QCompleter) completionPrefix(args ...interface{}) () {
  // completionPrefix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter16completionPrefixEv
  default:
    qtrt.ErrorResolve("QCompleter", "completionPrefix", args)
  }

}


func (this *QCompleter) setWidget(args ...interface{}) () {
  // setWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QCompleter9setWidgetEP7QWidget
  default:
    qtrt.ErrorResolve("QCompleter", "setWidget", args)
  }

}

// <= body block end

