package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qcombobox.h
// dst-file: /src/widgets/qcombobox.go
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
// class sizeof(QComboBox)=1
type QComboBox struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
//  _currentIndexChanged QComboBox_currentIndexChanged_signal;
//  _currentTextChanged QComboBox_currentTextChanged_signal;
//  _highlighted QComboBox_highlighted_signal;
//  _activated QComboBox_activated_signal;
//  _editTextChanged QComboBox_editTextChanged_signal;
}


func (this *QComboBox) setModel(args ...interface{}) () {
  // setModel(class QAbstractItemModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemModel{}) // "QAbstractItemModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox8setModelEP18QAbstractItemModel
  default:
    qtrt.ErrorResolve("QComboBox", "setModel", args)
  }

}


func (this *QComboBox) clearEditText(args ...interface{}) () {
  // clearEditText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox13clearEditTextEv
  default:
    qtrt.ErrorResolve("QComboBox", "clearEditText", args)
  }

}


func (this *QComboBox) setAutoCompletion(args ...interface{}) () {
  // setAutoCompletion(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox17setAutoCompletionEb
  default:
    qtrt.ErrorResolve("QComboBox", "setAutoCompletion", args)
  }

}


func (this *QComboBox) setFrame(args ...interface{}) () {
  // setFrame(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox8setFrameEb
  default:
    qtrt.ErrorResolve("QComboBox", "setFrame", args)
  }

}


func (this *QComboBox) setIconSize(args ...interface{}) () {
  // setIconSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox11setIconSizeERK5QSize
  default:
    qtrt.ErrorResolve("QComboBox", "setIconSize", args)
  }

}


func (this *QComboBox) setView(args ...interface{}) () {
  // setView(class QAbstractItemView *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemView{}) // "QAbstractItemView *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox7setViewEP17QAbstractItemView
  default:
    qtrt.ErrorResolve("QComboBox", "setView", args)
  }

}


func (this *QComboBox) view(args ...interface{}) () {
  // view()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox4viewEv
  default:
    qtrt.ErrorResolve("QComboBox", "view", args)
  }

}


func (this *QComboBox) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox15minimumSizeHintEv
  default:
    qtrt.ErrorResolve("QComboBox", "minimumSizeHint", args)
  }

}


func (this *QComboBox) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox5clearEv
  default:
    qtrt.ErrorResolve("QComboBox", "clear", args)
  }

}


func (this *QComboBox) maxCount(args ...interface{}) () {
  // maxCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox8maxCountEv
  default:
    qtrt.ErrorResolve("QComboBox", "maxCount", args)
  }

}


func (this *QComboBox) addItem(args ...interface{}) () {
  // addItem(const class QIcon &, const class QString &, const class QVariant &)
  // addItem(const class QString &, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QVariant{}) // "const QVariant &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox7addItemERK5QIconRK7QStringRK8QVariant
  case 1:
    // invoke: _ZN9QComboBox7addItemERK7QStringRK8QVariant
  default:
    qtrt.ErrorResolve("QComboBox", "addItem", args)
  }

}


func (this *QComboBox) insertItems(args ...interface{}) () {
  // insertItems(int, const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox11insertItemsEiRK11QStringList
  default:
    qtrt.ErrorResolve("QComboBox", "insertItems", args)
  }

}


func (this *QComboBox) iconSize(args ...interface{}) () {
  // iconSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox8iconSizeEv
  default:
    qtrt.ErrorResolve("QComboBox", "iconSize", args)
  }

}


func (this *QComboBox) rootModelIndex(args ...interface{}) () {
  // rootModelIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox14rootModelIndexEv
  default:
    qtrt.ErrorResolve("QComboBox", "rootModelIndex", args)
  }

}


func (this *QComboBox) setEditable(args ...interface{}) () {
  // setEditable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox11setEditableEb
  default:
    qtrt.ErrorResolve("QComboBox", "setEditable", args)
  }

}


func (this *QComboBox) setItemIcon(args ...interface{}) () {
  // setItemIcon(int, const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox11setItemIconEiRK5QIcon
  default:
    qtrt.ErrorResolve("QComboBox", "setItemIcon", args)
  }

}


func (this *QComboBox) autoCompletion(args ...interface{}) () {
  // autoCompletion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox14autoCompletionEv
  default:
    qtrt.ErrorResolve("QComboBox", "autoCompletion", args)
  }

}


func (this *QComboBox) currentData(args ...interface{}) () {
  // currentData(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox11currentDataEi
  default:
    qtrt.ErrorResolve("QComboBox", "currentData", args)
  }

}


func (this *QComboBox) hasFrame(args ...interface{}) () {
  // hasFrame()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox8hasFrameEv
  default:
    qtrt.ErrorResolve("QComboBox", "hasFrame", args)
  }

}


func (this *QComboBox) validator(args ...interface{}) () {
  // validator()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox9validatorEv
  default:
    qtrt.ErrorResolve("QComboBox", "validator", args)
  }

}


func (this *QComboBox) itemText(args ...interface{}) () {
  // itemText(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox8itemTextEi
  default:
    qtrt.ErrorResolve("QComboBox", "itemText", args)
  }

}


func (this *QComboBox) setItemData(args ...interface{}) () {
  // setItemData(int, const class QVariant &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox11setItemDataEiRK8QVarianti
  default:
    qtrt.ErrorResolve("QComboBox", "setItemData", args)
  }

}


func (this *QComboBox) hidePopup(args ...interface{}) () {
  // hidePopup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox9hidePopupEv
  default:
    qtrt.ErrorResolve("QComboBox", "hidePopup", args)
  }

}


func (this *QComboBox) insertItem(args ...interface{}) () {
  // insertItem(int, const class QIcon &, const class QString &, const class QVariant &)
  // insertItem(int, const class QString &, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][3] = reflect.TypeOf(QVariant{}) // "const QVariant &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][2] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox10insertItemEiRK5QIconRK7QStringRK8QVariant
  case 1:
    // invoke: _ZN9QComboBox10insertItemEiRK7QStringRK8QVariant
  default:
    qtrt.ErrorResolve("QComboBox", "insertItem", args)
  }

}


func (this *QComboBox) setCurrentText(args ...interface{}) () {
  // setCurrentText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox14setCurrentTextERK7QString
  default:
    qtrt.ErrorResolve("QComboBox", "setCurrentText", args)
  }

}


func (this *QComboBox) modelColumn(args ...interface{}) () {
  // modelColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox11modelColumnEv
  default:
    qtrt.ErrorResolve("QComboBox", "modelColumn", args)
  }

}


func (this *QComboBox) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox8sizeHintEv
  default:
    qtrt.ErrorResolve("QComboBox", "sizeHint", args)
  }

}


func (this *QComboBox) itemData(args ...interface{}) () {
  // itemData(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox8itemDataEii
  default:
    qtrt.ErrorResolve("QComboBox", "itemData", args)
  }

}


func (this *QComboBox) setCompleter(args ...interface{}) () {
  // setCompleter(class QCompleter *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCompleter{}) // "QCompleter *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox12setCompleterEP10QCompleter
  default:
    qtrt.ErrorResolve("QComboBox", "setCompleter", args)
  }

}


func (this *QComboBox) maxVisibleItems(args ...interface{}) () {
  // maxVisibleItems()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox15maxVisibleItemsEv
  default:
    qtrt.ErrorResolve("QComboBox", "maxVisibleItems", args)
  }

}


func NewQComboBox(args ...interface{}) QComboBox {
  return QComboBox{}
}


func (this *QComboBox) setCurrentIndex(args ...interface{}) () {
  // setCurrentIndex(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox15setCurrentIndexEi
  default:
    qtrt.ErrorResolve("QComboBox", "setCurrentIndex", args)
  }

}


func (this *QComboBox) setRootModelIndex(args ...interface{}) () {
  // setRootModelIndex(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox17setRootModelIndexERK11QModelIndex
  default:
    qtrt.ErrorResolve("QComboBox", "setRootModelIndex", args)
  }

}


func (this *QComboBox) setEditText(args ...interface{}) () {
  // setEditText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox11setEditTextERK7QString
  default:
    qtrt.ErrorResolve("QComboBox", "setEditText", args)
  }

}


func (this *QComboBox) completer(args ...interface{}) () {
  // completer()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox9completerEv
  default:
    qtrt.ErrorResolve("QComboBox", "completer", args)
  }

}


func (this *QComboBox) removeItem(args ...interface{}) () {
  // removeItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox10removeItemEi
  default:
    qtrt.ErrorResolve("QComboBox", "removeItem", args)
  }

}


func (this *QComboBox) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox5countEv
  default:
    qtrt.ErrorResolve("QComboBox", "count", args)
  }

}


func (this *QComboBox) setItemDelegate(args ...interface{}) () {
  // setItemDelegate(class QAbstractItemDelegate *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemDelegate{}) // "QAbstractItemDelegate *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox15setItemDelegateEP21QAbstractItemDelegate
  default:
    qtrt.ErrorResolve("QComboBox", "setItemDelegate", args)
  }

}


func (this *QComboBox) addItems(args ...interface{}) () {
  // addItems(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox8addItemsERK11QStringList
  default:
    qtrt.ErrorResolve("QComboBox", "addItems", args)
  }

}


func (this *QComboBox) setMinimumContentsLength(args ...interface{}) () {
  // setMinimumContentsLength(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox24setMinimumContentsLengthEi
  default:
    qtrt.ErrorResolve("QComboBox", "setMinimumContentsLength", args)
  }

}


func (this *QComboBox) duplicatesEnabled(args ...interface{}) () {
  // duplicatesEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox17duplicatesEnabledEv
  default:
    qtrt.ErrorResolve("QComboBox", "duplicatesEnabled", args)
  }

}


func (this *QComboBox) FreeQComboBox(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QComboBox", "~QComboBox", args)
  }

}


func (this *QComboBox) model(args ...interface{}) () {
  // model()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox5modelEv
  default:
    qtrt.ErrorResolve("QComboBox", "model", args)
  }

}


func (this *QComboBox) minimumContentsLength(args ...interface{}) () {
  // minimumContentsLength()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox21minimumContentsLengthEv
  default:
    qtrt.ErrorResolve("QComboBox", "minimumContentsLength", args)
  }

}


func (this *QComboBox) isEditable(args ...interface{}) () {
  // isEditable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox10isEditableEv
  default:
    qtrt.ErrorResolve("QComboBox", "isEditable", args)
  }

}


func (this *QComboBox) setMaxCount(args ...interface{}) () {
  // setMaxCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox11setMaxCountEi
  default:
    qtrt.ErrorResolve("QComboBox", "setMaxCount", args)
  }

}


func (this *QComboBox) currentIndex(args ...interface{}) () {
  // currentIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox12currentIndexEv
  default:
    qtrt.ErrorResolve("QComboBox", "currentIndex", args)
  }

}


func (this *QComboBox) setDuplicatesEnabled(args ...interface{}) () {
  // setDuplicatesEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox20setDuplicatesEnabledEb
  default:
    qtrt.ErrorResolve("QComboBox", "setDuplicatesEnabled", args)
  }

}


func (this *QComboBox) currentText(args ...interface{}) () {
  // currentText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox11currentTextEv
  default:
    qtrt.ErrorResolve("QComboBox", "currentText", args)
  }

}


func (this *QComboBox) showPopup(args ...interface{}) () {
  // showPopup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox9showPopupEv
  default:
    qtrt.ErrorResolve("QComboBox", "showPopup", args)
  }

}


func (this *QComboBox) lineEdit(args ...interface{}) () {
  // lineEdit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox8lineEditEv
  default:
    qtrt.ErrorResolve("QComboBox", "lineEdit", args)
  }

}


func (this *QComboBox) itemDelegate(args ...interface{}) () {
  // itemDelegate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox12itemDelegateEv
  default:
    qtrt.ErrorResolve("QComboBox", "itemDelegate", args)
  }

}


func (this *QComboBox) setMaxVisibleItems(args ...interface{}) () {
  // setMaxVisibleItems(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox18setMaxVisibleItemsEi
  default:
    qtrt.ErrorResolve("QComboBox", "setMaxVisibleItems", args)
  }

}


func (this *QComboBox) event(args ...interface{}) () {
  // event(class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox5eventEP6QEvent
  default:
    qtrt.ErrorResolve("QComboBox", "event", args)
  }

}


func (this *QComboBox) setModelColumn(args ...interface{}) () {
  // setModelColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox14setModelColumnEi
  default:
    qtrt.ErrorResolve("QComboBox", "setModelColumn", args)
  }

}


func (this *QComboBox) setItemText(args ...interface{}) () {
  // setItemText(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox11setItemTextEiRK7QString
  default:
    qtrt.ErrorResolve("QComboBox", "setItemText", args)
  }

}


func (this *QComboBox) setLineEdit(args ...interface{}) () {
  // setLineEdit(class QLineEdit *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLineEdit{}) // "QLineEdit *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox11setLineEditEP9QLineEdit
  default:
    qtrt.ErrorResolve("QComboBox", "setLineEdit", args)
  }

}


func (this *QComboBox) itemIcon(args ...interface{}) () {
  // itemIcon(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox8itemIconEi
  default:
    qtrt.ErrorResolve("QComboBox", "itemIcon", args)
  }

}


func (this *QComboBox) setValidator(args ...interface{}) () {
  // setValidator(const class QValidator *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QValidator{}) // "const QValidator *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox12setValidatorEPK10QValidator
  default:
    qtrt.ErrorResolve("QComboBox", "setValidator", args)
  }

}


func (this *QComboBox) insertSeparator(args ...interface{}) () {
  // insertSeparator(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox15insertSeparatorEi
  default:
    qtrt.ErrorResolve("QComboBox", "insertSeparator", args)
  }

}


func (this *QComboBox) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox10metaObjectEv
  default:
    qtrt.ErrorResolve("QComboBox", "metaObject", args)
  }

}

// <= body block end

