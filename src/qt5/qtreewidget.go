package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qtreewidget.h
// dst-file: /src/widgets/qtreewidget.go
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
// class sizeof(QTreeWidget)=1
type QTreeWidget struct {
  /*qbase*/ QTreeView;
  qclsinst uint64 /* *mut c_void*/;
//  _itemDoubleClicked QTreeWidget_itemDoubleClicked_signal;
//  _itemClicked QTreeWidget_itemClicked_signal;
//  _currentItemChanged QTreeWidget_currentItemChanged_signal;
//  _itemEntered QTreeWidget_itemEntered_signal;
//  _itemPressed QTreeWidget_itemPressed_signal;
//  _itemSelectionChanged QTreeWidget_itemSelectionChanged_signal;
//  _itemActivated QTreeWidget_itemActivated_signal;
//  _itemExpanded QTreeWidget_itemExpanded_signal;
//  _itemChanged QTreeWidget_itemChanged_signal;
//  _itemCollapsed QTreeWidget_itemCollapsed_signal;
}

// class sizeof(QTreeWidgetItem)=1
type QTreeWidgetItem struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QTreeWidget) setColumnCount(args ...interface{}) () {
  // setColumnCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget14setColumnCountEi
  default:
    qtrt.ErrorResolve("QTreeWidget", "setColumnCount", args)
  }

}


func (this *QTreeWidget) FreeQTreeWidget(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTreeWidget", "~QTreeWidget", args)
  }

}


func (this *QTreeWidget) selectedItems(args ...interface{}) () {
  // selectedItems()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget13selectedItemsEv
  default:
    qtrt.ErrorResolve("QTreeWidget", "selectedItems", args)
  }

}


func (this *QTreeWidget) isItemExpanded(args ...interface{}) () {
  // isItemExpanded(const class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget14isItemExpandedEPK15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "isItemExpanded", args)
  }

}


func NewQTreeWidget(args ...interface{}) QTreeWidget {
  return QTreeWidget{}
}


func (this *QTreeWidget) setItemHidden(args ...interface{}) () {
  // setItemHidden(const class QTreeWidgetItem *, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget13setItemHiddenEPK15QTreeWidgetItemb
  default:
    qtrt.ErrorResolve("QTreeWidget", "setItemHidden", args)
  }

}


func (this *QTreeWidget) indexOfTopLevelItem(args ...interface{}) () {
  // indexOfTopLevelItem(class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget19indexOfTopLevelItemEP15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "indexOfTopLevelItem", args)
  }

}


func (this *QTreeWidget) insertTopLevelItem(args ...interface{}) () {
  // insertTopLevelItem(int, class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget18insertTopLevelItemEiP15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "insertTopLevelItem", args)
  }

}


func (this *QTreeWidget) setItemWidget(args ...interface{}) () {
  // setItemWidget(class QTreeWidgetItem *, int, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget13setItemWidgetEP15QTreeWidgetItemiP7QWidget
  default:
    qtrt.ErrorResolve("QTreeWidget", "setItemWidget", args)
  }

}


func (this *QTreeWidget) isItemSelected(args ...interface{}) () {
  // isItemSelected(const class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget14isItemSelectedEPK15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "isItemSelected", args)
  }

}


func (this *QTreeWidget) currentColumn(args ...interface{}) () {
  // currentColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget13currentColumnEv
  default:
    qtrt.ErrorResolve("QTreeWidget", "currentColumn", args)
  }

}


func (this *QTreeWidget) isFirstItemColumnSpanned(args ...interface{}) () {
  // isFirstItemColumnSpanned(const class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget24isFirstItemColumnSpannedEPK15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "isFirstItemColumnSpanned", args)
  }

}


func (this *QTreeWidget) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget5clearEv
  default:
    qtrt.ErrorResolve("QTreeWidget", "clear", args)
  }

}


func (this *QTreeWidget) setHeaderLabels(args ...interface{}) () {
  // setHeaderLabels(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget15setHeaderLabelsERK11QStringList
  default:
    qtrt.ErrorResolve("QTreeWidget", "setHeaderLabels", args)
  }

}


func (this *QTreeWidget) invisibleRootItem(args ...interface{}) () {
  // invisibleRootItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget17invisibleRootItemEv
  default:
    qtrt.ErrorResolve("QTreeWidget", "invisibleRootItem", args)
  }

}


func (this *QTreeWidget) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget10metaObjectEv
  default:
    qtrt.ErrorResolve("QTreeWidget", "metaObject", args)
  }

}


func (this *QTreeWidget) itemBelow(args ...interface{}) () {
  // itemBelow(const class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget9itemBelowEPK15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "itemBelow", args)
  }

}


func (this *QTreeWidget) sortColumn(args ...interface{}) () {
  // sortColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget10sortColumnEv
  default:
    qtrt.ErrorResolve("QTreeWidget", "sortColumn", args)
  }

}


func (this *QTreeWidget) itemAt(args ...interface{}) () {
  // itemAt(int, int)
  // itemAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget6itemAtEii
  case 1:
    // invoke: _ZNK11QTreeWidget6itemAtERK6QPoint
  default:
    qtrt.ErrorResolve("QTreeWidget", "itemAt", args)
  }

}


func (this *QTreeWidget) currentItem(args ...interface{}) () {
  // currentItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget11currentItemEv
  default:
    qtrt.ErrorResolve("QTreeWidget", "currentItem", args)
  }

}


func (this *QTreeWidget) setCurrentItem(args ...interface{}) () {
  // setCurrentItem(class QTreeWidgetItem *, int, class QItemSelectionModel::SelectionFlags)
  // setCurrentItem(class QTreeWidgetItem *, int)
  // setCurrentItem(class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int64Ty(false) // "QItemSelectionModel::SelectionFlags"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItemi6QFlagsIN19QItemSelectionModel13SelectionFlagEE
  case 1:
    // invoke: _ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItemi
  case 2:
    // invoke: _ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "setCurrentItem", args)
  }

}


func (this *QTreeWidget) topLevelItem(args ...interface{}) () {
  // topLevelItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget12topLevelItemEi
  default:
    qtrt.ErrorResolve("QTreeWidget", "topLevelItem", args)
  }

}


func (this *QTreeWidget) topLevelItemCount(args ...interface{}) () {
  // topLevelItemCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget17topLevelItemCountEv
  default:
    qtrt.ErrorResolve("QTreeWidget", "topLevelItemCount", args)
  }

}


func (this *QTreeWidget) headerItem(args ...interface{}) () {
  // headerItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget10headerItemEv
  default:
    qtrt.ErrorResolve("QTreeWidget", "headerItem", args)
  }

}


func (this *QTreeWidget) setFirstItemColumnSpanned(args ...interface{}) () {
  // setFirstItemColumnSpanned(const class QTreeWidgetItem *, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget25setFirstItemColumnSpannedEPK15QTreeWidgetItemb
  default:
    qtrt.ErrorResolve("QTreeWidget", "setFirstItemColumnSpanned", args)
  }

}


func (this *QTreeWidget) removeItemWidget(args ...interface{}) () {
  // removeItemWidget(class QTreeWidgetItem *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget16removeItemWidgetEP15QTreeWidgetItemi
  default:
    qtrt.ErrorResolve("QTreeWidget", "removeItemWidget", args)
  }

}


func (this *QTreeWidget) itemAbove(args ...interface{}) () {
  // itemAbove(const class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget9itemAboveEPK15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "itemAbove", args)
  }

}


func (this *QTreeWidget) expandItem(args ...interface{}) () {
  // expandItem(const class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget10expandItemEPK15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "expandItem", args)
  }

}


func (this *QTreeWidget) setHeaderItem(args ...interface{}) () {
  // setHeaderItem(class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget13setHeaderItemEP15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "setHeaderItem", args)
  }

}


func (this *QTreeWidget) collapseItem(args ...interface{}) () {
  // collapseItem(const class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget12collapseItemEPK15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "collapseItem", args)
  }

}


func (this *QTreeWidget) takeTopLevelItem(args ...interface{}) () {
  // takeTopLevelItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget16takeTopLevelItemEi
  default:
    qtrt.ErrorResolve("QTreeWidget", "takeTopLevelItem", args)
  }

}


func (this *QTreeWidget) itemWidget(args ...interface{}) () {
  // itemWidget(class QTreeWidgetItem *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget10itemWidgetEP15QTreeWidgetItemi
  default:
    qtrt.ErrorResolve("QTreeWidget", "itemWidget", args)
  }

}


func (this *QTreeWidget) editItem(args ...interface{}) () {
  // editItem(class QTreeWidgetItem *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget8editItemEP15QTreeWidgetItemi
  default:
    qtrt.ErrorResolve("QTreeWidget", "editItem", args)
  }

}


func (this *QTreeWidget) setItemExpanded(args ...interface{}) () {
  // setItemExpanded(const class QTreeWidgetItem *, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget15setItemExpandedEPK15QTreeWidgetItemb
  default:
    qtrt.ErrorResolve("QTreeWidget", "setItemExpanded", args)
  }

}


func (this *QTreeWidget) addTopLevelItem(args ...interface{}) () {
  // addTopLevelItem(class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget15addTopLevelItemEP15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "addTopLevelItem", args)
  }

}


func (this *QTreeWidget) closePersistentEditor(args ...interface{}) () {
  // closePersistentEditor(class QTreeWidgetItem *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget21closePersistentEditorEP15QTreeWidgetItemi
  default:
    qtrt.ErrorResolve("QTreeWidget", "closePersistentEditor", args)
  }

}


func (this *QTreeWidget) setSelectionModel(args ...interface{}) () {
  // setSelectionModel(class QItemSelectionModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QItemSelectionModel{}) // "QItemSelectionModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget17setSelectionModelEP19QItemSelectionModel
  default:
    qtrt.ErrorResolve("QTreeWidget", "setSelectionModel", args)
  }

}


func (this *QTreeWidget) visualItemRect(args ...interface{}) () {
  // visualItemRect(const class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget14visualItemRectEPK15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "visualItemRect", args)
  }

}


func (this *QTreeWidget) setHeaderLabel(args ...interface{}) () {
  // setHeaderLabel(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget14setHeaderLabelERK7QString
  default:
    qtrt.ErrorResolve("QTreeWidget", "setHeaderLabel", args)
  }

}


func (this *QTreeWidget) isItemHidden(args ...interface{}) () {
  // isItemHidden(const class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget12isItemHiddenEPK15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "isItemHidden", args)
  }

}


func (this *QTreeWidget) openPersistentEditor(args ...interface{}) () {
  // openPersistentEditor(class QTreeWidgetItem *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget20openPersistentEditorEP15QTreeWidgetItemi
  default:
    qtrt.ErrorResolve("QTreeWidget", "openPersistentEditor", args)
  }

}


func (this *QTreeWidget) columnCount(args ...interface{}) () {
  // columnCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget11columnCountEv
  default:
    qtrt.ErrorResolve("QTreeWidget", "columnCount", args)
  }

}


func (this *QTreeWidget) setItemSelected(args ...interface{}) () {
  // setItemSelected(const class QTreeWidgetItem *, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget15setItemSelectedEPK15QTreeWidgetItemb
  default:
    qtrt.ErrorResolve("QTreeWidget", "setItemSelected", args)
  }

}


func (this *QTreeWidgetItem) setFirstColumnSpanned(args ...interface{}) () {
  // setFirstColumnSpanned(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem21setFirstColumnSpannedEb
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setFirstColumnSpanned", args)
  }

}


func (this *QTreeWidgetItem) indexOfChild(args ...interface{}) () {
  // indexOfChild(class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem12indexOfChildEPS_
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "indexOfChild", args)
  }

}


func (this *QTreeWidgetItem) data(args ...interface{}) () {
  // data(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem4dataEii
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "data", args)
  }

}


func (this *QTreeWidgetItem) parent(args ...interface{}) () {
  // parent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem6parentEv
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "parent", args)
  }

}


func (this *QTreeWidgetItem) setFont(args ...interface{}) () {
  // setFont(int, const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem7setFontEiRK5QFont
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setFont", args)
  }

}


func (this *QTreeWidgetItem) setData(args ...interface{}) () {
  // setData(int, int, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem7setDataEiiRK8QVariant
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setData", args)
  }

}


func (this *QTreeWidgetItem) font(args ...interface{}) () {
  // font(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem4fontEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "font", args)
  }

}


func (this *QTreeWidgetItem) setStatusTip(args ...interface{}) () {
  // setStatusTip(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem12setStatusTipEiRK7QString
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setStatusTip", args)
  }

}


func (this *QTreeWidgetItem) setExpanded(args ...interface{}) () {
  // setExpanded(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem11setExpandedEb
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setExpanded", args)
  }

}


func (this *QTreeWidgetItem) write(args ...interface{}) () {
  // write(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDataStream{}) // "QDataStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem5writeER11QDataStream
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "write", args)
  }

}


func (this *QTreeWidgetItem) isExpanded(args ...interface{}) () {
  // isExpanded()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem10isExpandedEv
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "isExpanded", args)
  }

}


func (this *QTreeWidgetItem) takeChildren(args ...interface{}) () {
  // takeChildren()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem12takeChildrenEv
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "takeChildren", args)
  }

}


func NewQTreeWidgetItem(args ...interface{}) QTreeWidgetItem {
  return QTreeWidgetItem{}
}


func (this *QTreeWidgetItem) setIcon(args ...interface{}) () {
  // setIcon(int, const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem7setIconEiRK5QIcon
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setIcon", args)
  }

}


func (this *QTreeWidgetItem) toolTip(args ...interface{}) () {
  // toolTip(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem7toolTipEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "toolTip", args)
  }

}


func (this *QTreeWidgetItem) backgroundColor(args ...interface{}) () {
  // backgroundColor(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem15backgroundColorEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "backgroundColor", args)
  }

}


func (this *QTreeWidgetItem) text(args ...interface{}) () {
  // text(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem4textEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "text", args)
  }

}


func (this *QTreeWidgetItem) isHidden(args ...interface{}) () {
  // isHidden()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem8isHiddenEv
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "isHidden", args)
  }

}


func (this *QTreeWidgetItem) setTextAlignment(args ...interface{}) () {
  // setTextAlignment(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem16setTextAlignmentEii
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setTextAlignment", args)
  }

}


func (this *QTreeWidgetItem) insertChild(args ...interface{}) () {
  // insertChild(int, class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem11insertChildEiPS_
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "insertChild", args)
  }

}


func (this *QTreeWidgetItem) isDisabled(args ...interface{}) () {
  // isDisabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem10isDisabledEv
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "isDisabled", args)
  }

}


func (this *QTreeWidgetItem) setText(args ...interface{}) () {
  // setText(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem7setTextEiRK7QString
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setText", args)
  }

}


func (this *QTreeWidgetItem) setTextColor(args ...interface{}) () {
  // setTextColor(int, const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem12setTextColorEiRK6QColor
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setTextColor", args)
  }

}


func (this *QTreeWidgetItem) sizeHint(args ...interface{}) () {
  // sizeHint(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem8sizeHintEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "sizeHint", args)
  }

}


func (this *QTreeWidgetItem) whatsThis(args ...interface{}) () {
  // whatsThis(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem9whatsThisEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "whatsThis", args)
  }

}


func (this *QTreeWidgetItem) setWhatsThis(args ...interface{}) () {
  // setWhatsThis(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem12setWhatsThisEiRK7QString
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setWhatsThis", args)
  }

}


func (this *QTreeWidgetItem) textColor(args ...interface{}) () {
  // textColor(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem9textColorEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "textColor", args)
  }

}


func (this *QTreeWidgetItem) icon(args ...interface{}) () {
  // icon(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem4iconEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "icon", args)
  }

}


func (this *QTreeWidgetItem) setToolTip(args ...interface{}) () {
  // setToolTip(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem10setToolTipEiRK7QString
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setToolTip", args)
  }

}


func (this *QTreeWidgetItem) isFirstColumnSpanned(args ...interface{}) () {
  // isFirstColumnSpanned()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem20isFirstColumnSpannedEv
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "isFirstColumnSpanned", args)
  }

}


func (this *QTreeWidgetItem) textAlignment(args ...interface{}) () {
  // textAlignment(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem13textAlignmentEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "textAlignment", args)
  }

}


func (this *QTreeWidgetItem) child(args ...interface{}) () {
  // child(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem5childEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "child", args)
  }

}


func (this *QTreeWidgetItem) setSelected(args ...interface{}) () {
  // setSelected(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem11setSelectedEb
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setSelected", args)
  }

}


func (this *QTreeWidgetItem) FreeQTreeWidgetItem(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "~QTreeWidgetItem", args)
  }

}


func (this *QTreeWidgetItem) setHidden(args ...interface{}) () {
  // setHidden(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem9setHiddenEb
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setHidden", args)
  }

}


func (this *QTreeWidgetItem) columnCount(args ...interface{}) () {
  // columnCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem11columnCountEv
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "columnCount", args)
  }

}


func (this *QTreeWidgetItem) takeChild(args ...interface{}) () {
  // takeChild(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem9takeChildEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "takeChild", args)
  }

}


func (this *QTreeWidgetItem) setDisabled(args ...interface{}) () {
  // setDisabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem11setDisabledEb
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setDisabled", args)
  }

}


func (this *QTreeWidgetItem) setBackground(args ...interface{}) () {
  // setBackground(int, const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem13setBackgroundEiRK6QBrush
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setBackground", args)
  }

}


func (this *QTreeWidgetItem) addChild(args ...interface{}) () {
  // addChild(class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem8addChildEPS_
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "addChild", args)
  }

}


func (this *QTreeWidgetItem) removeChild(args ...interface{}) () {
  // removeChild(class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem11removeChildEPS_
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "removeChild", args)
  }

}


func (this *QTreeWidgetItem) clone(args ...interface{}) () {
  // clone()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem5cloneEv
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "clone", args)
  }

}


func (this *QTreeWidgetItem) setSizeHint(args ...interface{}) () {
  // setSizeHint(int, const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem11setSizeHintEiRK5QSize
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setSizeHint", args)
  }

}


func (this *QTreeWidgetItem) foreground(args ...interface{}) () {
  // foreground(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem10foregroundEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "foreground", args)
  }

}


func (this *QTreeWidgetItem) childCount(args ...interface{}) () {
  // childCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem10childCountEv
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "childCount", args)
  }

}


func (this *QTreeWidgetItem) setBackgroundColor(args ...interface{}) () {
  // setBackgroundColor(int, const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem18setBackgroundColorEiRK6QColor
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setBackgroundColor", args)
  }

}


func (this *QTreeWidgetItem) statusTip(args ...interface{}) () {
  // statusTip(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem9statusTipEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "statusTip", args)
  }

}


func (this *QTreeWidgetItem) background(args ...interface{}) () {
  // background(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem10backgroundEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "background", args)
  }

}


func (this *QTreeWidgetItem) type_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "type", args)
  }

}


func (this *QTreeWidgetItem) treeWidget(args ...interface{}) () {
  // treeWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem10treeWidgetEv
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "treeWidget", args)
  }

}


func (this *QTreeWidgetItem) read(args ...interface{}) () {
  // read(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDataStream{}) // "QDataStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem4readER11QDataStream
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "read", args)
  }

}


func (this *QTreeWidgetItem) setForeground(args ...interface{}) () {
  // setForeground(int, const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem13setForegroundEiRK6QBrush
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setForeground", args)
  }

}


func (this *QTreeWidgetItem) isSelected(args ...interface{}) () {
  // isSelected()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem10isSelectedEv
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "isSelected", args)
  }

}

// <= body block end

