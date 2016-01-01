package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qlistwidget.h
// dst-file: /src/widgets/qlistwidget.go
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
// class sizeof(QListWidgetItem)=1
type QListWidgetItem struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QListWidget)=1
type QListWidget struct {
  /*qbase*/ QListView;
  qclsinst uint64 /* *mut c_void*/;
//  _itemDoubleClicked QListWidget_itemDoubleClicked_signal;
//  _itemClicked QListWidget_itemClicked_signal;
//  _currentItemChanged QListWidget_currentItemChanged_signal;
//  _itemEntered QListWidget_itemEntered_signal;
//  _itemPressed QListWidget_itemPressed_signal;
//  _itemSelectionChanged QListWidget_itemSelectionChanged_signal;
//  _itemActivated QListWidget_itemActivated_signal;
//  _itemChanged QListWidget_itemChanged_signal;
//  _currentRowChanged QListWidget_currentRowChanged_signal;
//  _currentTextChanged QListWidget_currentTextChanged_signal;
}


func (this *QListWidgetItem) FreeQListWidgetItem(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QListWidgetItem", "~QListWidgetItem", args)
  }

}


func (this *QListWidgetItem) isHidden(args ...interface{}) () {
  // isHidden()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem8isHiddenEv
  default:
    qtrt.ErrorResolve("QListWidgetItem", "isHidden", args)
  }

}


func (this *QListWidgetItem) setData(args ...interface{}) () {
  // setData(int, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem7setDataEiRK8QVariant
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setData", args)
  }

}


func (this *QListWidgetItem) setBackground(args ...interface{}) () {
  // setBackground(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem13setBackgroundERK6QBrush
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setBackground", args)
  }

}


func (this *QListWidgetItem) setSelected(args ...interface{}) () {
  // setSelected(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem11setSelectedEb
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setSelected", args)
  }

}


func (this *QListWidgetItem) font(args ...interface{}) () {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem4fontEv
  default:
    qtrt.ErrorResolve("QListWidgetItem", "font", args)
  }

}


func (this *QListWidgetItem) setTextAlignment(args ...interface{}) () {
  // setTextAlignment(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem16setTextAlignmentEi
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setTextAlignment", args)
  }

}


func NewQListWidgetItem(args ...interface{}) QListWidgetItem {
  return QListWidgetItem{}
}


func (this *QListWidgetItem) write(args ...interface{}) () {
  // write(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDataStream{}) // "QDataStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem5writeER11QDataStream
  default:
    qtrt.ErrorResolve("QListWidgetItem", "write", args)
  }

}


func (this *QListWidgetItem) whatsThis(args ...interface{}) () {
  // whatsThis()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem9whatsThisEv
  default:
    qtrt.ErrorResolve("QListWidgetItem", "whatsThis", args)
  }

}


func (this *QListWidgetItem) type_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QListWidgetItem", "type", args)
  }

}


func (this *QListWidgetItem) icon(args ...interface{}) () {
  // icon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem4iconEv
  default:
    qtrt.ErrorResolve("QListWidgetItem", "icon", args)
  }

}


func (this *QListWidgetItem) textColor(args ...interface{}) () {
  // textColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem9textColorEv
  default:
    qtrt.ErrorResolve("QListWidgetItem", "textColor", args)
  }

}


func (this *QListWidgetItem) foreground(args ...interface{}) () {
  // foreground()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem10foregroundEv
  default:
    qtrt.ErrorResolve("QListWidgetItem", "foreground", args)
  }

}


func (this *QListWidgetItem) background(args ...interface{}) () {
  // background()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem10backgroundEv
  default:
    qtrt.ErrorResolve("QListWidgetItem", "background", args)
  }

}


func (this *QListWidgetItem) setStatusTip(args ...interface{}) () {
  // setStatusTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem12setStatusTipERK7QString
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setStatusTip", args)
  }

}


func (this *QListWidgetItem) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem4textEv
  default:
    qtrt.ErrorResolve("QListWidgetItem", "text", args)
  }

}


func (this *QListWidgetItem) backgroundColor(args ...interface{}) () {
  // backgroundColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem15backgroundColorEv
  default:
    qtrt.ErrorResolve("QListWidgetItem", "backgroundColor", args)
  }

}


func (this *QListWidgetItem) isSelected(args ...interface{}) () {
  // isSelected()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem10isSelectedEv
  default:
    qtrt.ErrorResolve("QListWidgetItem", "isSelected", args)
  }

}


func (this *QListWidgetItem) setFont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem7setFontERK5QFont
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setFont", args)
  }

}


func (this *QListWidgetItem) setText(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem7setTextERK7QString
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setText", args)
  }

}


func (this *QListWidgetItem) data(args ...interface{}) () {
  // data(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem4dataEi
  default:
    qtrt.ErrorResolve("QListWidgetItem", "data", args)
  }

}


func (this *QListWidgetItem) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem8sizeHintEv
  default:
    qtrt.ErrorResolve("QListWidgetItem", "sizeHint", args)
  }

}


func (this *QListWidgetItem) setWhatsThis(args ...interface{}) () {
  // setWhatsThis(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem12setWhatsThisERK7QString
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setWhatsThis", args)
  }

}


func (this *QListWidgetItem) read(args ...interface{}) () {
  // read(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDataStream{}) // "QDataStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem4readER11QDataStream
  default:
    qtrt.ErrorResolve("QListWidgetItem", "read", args)
  }

}


func (this *QListWidgetItem) setTextColor(args ...interface{}) () {
  // setTextColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem12setTextColorERK6QColor
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setTextColor", args)
  }

}


func (this *QListWidgetItem) setSizeHint(args ...interface{}) () {
  // setSizeHint(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem11setSizeHintERK5QSize
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setSizeHint", args)
  }

}


func (this *QListWidgetItem) listWidget(args ...interface{}) () {
  // listWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem10listWidgetEv
  default:
    qtrt.ErrorResolve("QListWidgetItem", "listWidget", args)
  }

}


func (this *QListWidgetItem) setIcon(args ...interface{}) () {
  // setIcon(const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem7setIconERK5QIcon
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setIcon", args)
  }

}


func (this *QListWidgetItem) clone(args ...interface{}) () {
  // clone()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem5cloneEv
  default:
    qtrt.ErrorResolve("QListWidgetItem", "clone", args)
  }

}


func (this *QListWidgetItem) setBackgroundColor(args ...interface{}) () {
  // setBackgroundColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem18setBackgroundColorERK6QColor
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setBackgroundColor", args)
  }

}


func (this *QListWidgetItem) setForeground(args ...interface{}) () {
  // setForeground(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem13setForegroundERK6QBrush
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setForeground", args)
  }

}


func (this *QListWidgetItem) setHidden(args ...interface{}) () {
  // setHidden(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem9setHiddenEb
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setHidden", args)
  }

}


func (this *QListWidgetItem) toolTip(args ...interface{}) () {
  // toolTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem7toolTipEv
  default:
    qtrt.ErrorResolve("QListWidgetItem", "toolTip", args)
  }

}


func (this *QListWidgetItem) textAlignment(args ...interface{}) () {
  // textAlignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem13textAlignmentEv
  default:
    qtrt.ErrorResolve("QListWidgetItem", "textAlignment", args)
  }

}


func (this *QListWidgetItem) statusTip(args ...interface{}) () {
  // statusTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem9statusTipEv
  default:
    qtrt.ErrorResolve("QListWidgetItem", "statusTip", args)
  }

}


func (this *QListWidgetItem) setToolTip(args ...interface{}) () {
  // setToolTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem10setToolTipERK7QString
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setToolTip", args)
  }

}


func (this *QListWidget) dropEvent(args ...interface{}) () {
  // dropEvent(class QDropEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDropEvent{}) // "QDropEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget9dropEventEP10QDropEvent
  default:
    qtrt.ErrorResolve("QListWidget", "dropEvent", args)
  }

}


func (this *QListWidget) itemWidget(args ...interface{}) () {
  // itemWidget(class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget10itemWidgetEP15QListWidgetItem
  default:
    qtrt.ErrorResolve("QListWidget", "itemWidget", args)
  }

}


func NewQListWidget(args ...interface{}) QListWidget {
  return QListWidget{}
}


func (this *QListWidget) currentRow(args ...interface{}) () {
  // currentRow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget10currentRowEv
  default:
    qtrt.ErrorResolve("QListWidget", "currentRow", args)
  }

}


func (this *QListWidget) item(args ...interface{}) () {
  // item(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget4itemEi
  default:
    qtrt.ErrorResolve("QListWidget", "item", args)
  }

}


func (this *QListWidget) itemAt(args ...interface{}) () {
  // itemAt(const class QPoint &)
  // itemAt(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget6itemAtERK6QPoint
  case 1:
    // invoke: _ZNK11QListWidget6itemAtEii
  default:
    qtrt.ErrorResolve("QListWidget", "itemAt", args)
  }

}


func (this *QListWidget) insertItem(args ...interface{}) () {
  // insertItem(int, const class QString &)
  // insertItem(int, class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget10insertItemEiRK7QString
  case 1:
    // invoke: _ZN11QListWidget10insertItemEiP15QListWidgetItem
  default:
    qtrt.ErrorResolve("QListWidget", "insertItem", args)
  }

}


func (this *QListWidget) row(args ...interface{}) () {
  // row(const class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "const QListWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget3rowEPK15QListWidgetItem
  default:
    qtrt.ErrorResolve("QListWidget", "row", args)
  }

}


func (this *QListWidget) openPersistentEditor(args ...interface{}) () {
  // openPersistentEditor(class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget20openPersistentEditorEP15QListWidgetItem
  default:
    qtrt.ErrorResolve("QListWidget", "openPersistentEditor", args)
  }

}


func (this *QListWidget) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget5clearEv
  default:
    qtrt.ErrorResolve("QListWidget", "clear", args)
  }

}


func (this *QListWidget) editItem(args ...interface{}) () {
  // editItem(class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget8editItemEP15QListWidgetItem
  default:
    qtrt.ErrorResolve("QListWidget", "editItem", args)
  }

}


func (this *QListWidget) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget5countEv
  default:
    qtrt.ErrorResolve("QListWidget", "count", args)
  }

}


func (this *QListWidget) setItemHidden(args ...interface{}) () {
  // setItemHidden(const class QListWidgetItem *, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "const QListWidgetItem *"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget13setItemHiddenEPK15QListWidgetItemb
  default:
    qtrt.ErrorResolve("QListWidget", "setItemHidden", args)
  }

}


func (this *QListWidget) FreeQListWidget(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QListWidget", "~QListWidget", args)
  }

}


func (this *QListWidget) addItem(args ...interface{}) () {
  // addItem(class QListWidgetItem *)
  // addItem(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget7addItemEP15QListWidgetItem
  case 1:
    // invoke: _ZN11QListWidget7addItemERK7QString
  default:
    qtrt.ErrorResolve("QListWidget", "addItem", args)
  }

}


func (this *QListWidget) takeItem(args ...interface{}) () {
  // takeItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget8takeItemEi
  default:
    qtrt.ErrorResolve("QListWidget", "takeItem", args)
  }

}


func (this *QListWidget) isSortingEnabled(args ...interface{}) () {
  // isSortingEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget16isSortingEnabledEv
  default:
    qtrt.ErrorResolve("QListWidget", "isSortingEnabled", args)
  }

}


func (this *QListWidget) addItems(args ...interface{}) () {
  // addItems(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget8addItemsERK11QStringList
  default:
    qtrt.ErrorResolve("QListWidget", "addItems", args)
  }

}


func (this *QListWidget) selectedItems(args ...interface{}) () {
  // selectedItems()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget13selectedItemsEv
  default:
    qtrt.ErrorResolve("QListWidget", "selectedItems", args)
  }

}


func (this *QListWidget) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget10metaObjectEv
  default:
    qtrt.ErrorResolve("QListWidget", "metaObject", args)
  }

}


func (this *QListWidget) setItemSelected(args ...interface{}) () {
  // setItemSelected(const class QListWidgetItem *, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "const QListWidgetItem *"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget15setItemSelectedEPK15QListWidgetItemb
  default:
    qtrt.ErrorResolve("QListWidget", "setItemSelected", args)
  }

}


func (this *QListWidget) setCurrentRow(args ...interface{}) () {
  // setCurrentRow(int, class QItemSelectionModel::SelectionFlags)
  // setCurrentRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int64Ty(false) // "QItemSelectionModel::SelectionFlags"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget13setCurrentRowEi6QFlagsIN19QItemSelectionModel13SelectionFlagEE
  case 1:
    // invoke: _ZN11QListWidget13setCurrentRowEi
  default:
    qtrt.ErrorResolve("QListWidget", "setCurrentRow", args)
  }

}


func (this *QListWidget) setSortingEnabled(args ...interface{}) () {
  // setSortingEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget17setSortingEnabledEb
  default:
    qtrt.ErrorResolve("QListWidget", "setSortingEnabled", args)
  }

}


func (this *QListWidget) visualItemRect(args ...interface{}) () {
  // visualItemRect(const class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "const QListWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget14visualItemRectEPK15QListWidgetItem
  default:
    qtrt.ErrorResolve("QListWidget", "visualItemRect", args)
  }

}


func (this *QListWidget) removeItemWidget(args ...interface{}) () {
  // removeItemWidget(class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget16removeItemWidgetEP15QListWidgetItem
  default:
    qtrt.ErrorResolve("QListWidget", "removeItemWidget", args)
  }

}


func (this *QListWidget) closePersistentEditor(args ...interface{}) () {
  // closePersistentEditor(class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget21closePersistentEditorEP15QListWidgetItem
  default:
    qtrt.ErrorResolve("QListWidget", "closePersistentEditor", args)
  }

}


func (this *QListWidget) isItemHidden(args ...interface{}) () {
  // isItemHidden(const class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "const QListWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget12isItemHiddenEPK15QListWidgetItem
  default:
    qtrt.ErrorResolve("QListWidget", "isItemHidden", args)
  }

}


func (this *QListWidget) insertItems(args ...interface{}) () {
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
    // invoke: _ZN11QListWidget11insertItemsEiRK11QStringList
  default:
    qtrt.ErrorResolve("QListWidget", "insertItems", args)
  }

}


func (this *QListWidget) currentItem(args ...interface{}) () {
  // currentItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget11currentItemEv
  default:
    qtrt.ErrorResolve("QListWidget", "currentItem", args)
  }

}


func (this *QListWidget) setCurrentItem(args ...interface{}) () {
  // setCurrentItem(class QListWidgetItem *, class QItemSelectionModel::SelectionFlags)
  // setCurrentItem(class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"
  vtys[0][1] = qtrt.Int64Ty(false) // "QItemSelectionModel::SelectionFlags"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget14setCurrentItemEP15QListWidgetItem6QFlagsIN19QItemSelectionModel13SelectionFlagEE
  case 1:
    // invoke: _ZN11QListWidget14setCurrentItemEP15QListWidgetItem
  default:
    qtrt.ErrorResolve("QListWidget", "setCurrentItem", args)
  }

}


func (this *QListWidget) setItemWidget(args ...interface{}) () {
  // setItemWidget(class QListWidgetItem *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget13setItemWidgetEP15QListWidgetItemP7QWidget
  default:
    qtrt.ErrorResolve("QListWidget", "setItemWidget", args)
  }

}


func (this *QListWidget) isItemSelected(args ...interface{}) () {
  // isItemSelected(const class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "const QListWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget14isItemSelectedEPK15QListWidgetItem
  default:
    qtrt.ErrorResolve("QListWidget", "isItemSelected", args)
  }

}

// <= body block end

