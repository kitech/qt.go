package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtWidgets/qtablewidget.h
// dst-file: /src/widgets/qtablewidget.go
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
// class sizeof(QTableWidgetSelectionRange)=16
type QTableWidgetSelectionRange struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTableWidget)=1
type QTableWidget struct {
  /*qbase*/ QTableView;
  qclsinst uint64 /* *mut c_void*/;
//  _itemDoubleClicked QTableWidget_itemDoubleClicked_signal;
//  _cellEntered QTableWidget_cellEntered_signal;
//  _itemClicked QTableWidget_itemClicked_signal;
//  _currentItemChanged QTableWidget_currentItemChanged_signal;
//  _itemEntered QTableWidget_itemEntered_signal;
//  _itemPressed QTableWidget_itemPressed_signal;
//  _cellClicked QTableWidget_cellClicked_signal;
//  _itemSelectionChanged QTableWidget_itemSelectionChanged_signal;
//  _cellChanged QTableWidget_cellChanged_signal;
//  _itemActivated QTableWidget_itemActivated_signal;
//  _cellActivated QTableWidget_cellActivated_signal;
//  _itemChanged QTableWidget_itemChanged_signal;
//  _currentCellChanged QTableWidget_currentCellChanged_signal;
//  _cellDoubleClicked QTableWidget_cellDoubleClicked_signal;
//  _cellPressed QTableWidget_cellPressed_signal;
}

// class sizeof(QTableWidgetItem)=1
type QTableWidgetItem struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQTableWidgetSelectionRange(args ...interface{})() {
}


func (this *QTableWidgetSelectionRange) columnCount(args ...interface{}) () {
  // columnCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK26QTableWidgetSelectionRange11columnCountEv
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "columnCount", args)
 }

}


func (this *QTableWidgetSelectionRange) rowCount(args ...interface{}) () {
  // rowCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK26QTableWidgetSelectionRange8rowCountEv
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "rowCount", args)
 }

}


func (this *QTableWidgetSelectionRange) leftColumn(args ...interface{}) () {
  // leftColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK26QTableWidgetSelectionRange10leftColumnEv
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "leftColumn", args)
 }

}


func (this *QTableWidgetSelectionRange) FreeQTableWidgetSelectionRange(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "~QTableWidgetSelectionRange", args)
 }

}


func (this *QTableWidgetSelectionRange) topRow(args ...interface{}) () {
  // topRow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK26QTableWidgetSelectionRange6topRowEv
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "topRow", args)
 }

}


func (this *QTableWidgetSelectionRange) rightColumn(args ...interface{}) () {
  // rightColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK26QTableWidgetSelectionRange11rightColumnEv
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "rightColumn", args)
 }

}


func (this *QTableWidgetSelectionRange) bottomRow(args ...interface{}) () {
  // bottomRow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK26QTableWidgetSelectionRange9bottomRowEv
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "bottomRow", args)
 }

}


func (this *QTableWidget) setColumnCount(args ...interface{}) () {
  // setColumnCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget14setColumnCountEi
  default:
    qtrt.ErrorResolve("QTableWidget", "setColumnCount", args)
 }

}


func (this *QTableWidget) FreeQTableWidget(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTableWidget", "~QTableWidget", args)
 }

}


func (this *QTableWidget) selectedItems(args ...interface{}) () {
  // selectedItems()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget13selectedItemsEv
  default:
    qtrt.ErrorResolve("QTableWidget", "selectedItems", args)
 }

}


func (this *QTableWidget) isSortingEnabled(args ...interface{}) () {
  // isSortingEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget16isSortingEnabledEv
  default:
    qtrt.ErrorResolve("QTableWidget", "isSortingEnabled", args)
 }

}


func (this *QTableWidget) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget10metaObjectEv
  default:
    qtrt.ErrorResolve("QTableWidget", "metaObject", args)
 }

}


func NewQTableWidget(args ...interface{})() {
}


func (this *QTableWidget) closePersistentEditor(args ...interface{}) () {
  // closePersistentEditor(class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget21closePersistentEditorEP16QTableWidgetItem
  default:
    qtrt.ErrorResolve("QTableWidget", "closePersistentEditor", args)
 }

}


func (this *QTableWidget) setHorizontalHeaderLabels(args ...interface{}) () {
  // setHorizontalHeaderLabels(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget25setHorizontalHeaderLabelsERK11QStringList
  default:
    qtrt.ErrorResolve("QTableWidget", "setHorizontalHeaderLabels", args)
 }

}


func (this *QTableWidget) setItemSelected(args ...interface{}) () {
  // setItemSelected(const class QTableWidgetItem *, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "const QTableWidgetItem *"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget15setItemSelectedEPK16QTableWidgetItemb
  default:
    qtrt.ErrorResolve("QTableWidget", "setItemSelected", args)
 }

}


func (this *QTableWidget) takeItem(args ...interface{}) () {
  // takeItem(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget8takeItemEii
  default:
    qtrt.ErrorResolve("QTableWidget", "takeItem", args)
 }

}


func (this *QTableWidget) removeCellWidget(args ...interface{}) () {
  // removeCellWidget(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget16removeCellWidgetEii
  default:
    qtrt.ErrorResolve("QTableWidget", "removeCellWidget", args)
 }

}


func (this *QTableWidget) setVerticalHeaderItem(args ...interface{}) () {
  // setVerticalHeaderItem(int, class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget21setVerticalHeaderItemEiP16QTableWidgetItem
  default:
    qtrt.ErrorResolve("QTableWidget", "setVerticalHeaderItem", args)
 }

}


func (this *QTableWidget) visualItemRect(args ...interface{}) () {
  // visualItemRect(const class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "const QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget14visualItemRectEPK16QTableWidgetItem
  default:
    qtrt.ErrorResolve("QTableWidget", "visualItemRect", args)
 }

}


func (this *QTableWidget) currentItem(args ...interface{}) () {
  // currentItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget11currentItemEv
  default:
    qtrt.ErrorResolve("QTableWidget", "currentItem", args)
 }

}


func (this *QTableWidget) row(args ...interface{}) () {
  // row(const class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "const QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget3rowEPK16QTableWidgetItem
  default:
    qtrt.ErrorResolve("QTableWidget", "row", args)
 }

}


func (this *QTableWidget) removeRow(args ...interface{}) () {
  // removeRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget9removeRowEi
  default:
    qtrt.ErrorResolve("QTableWidget", "removeRow", args)
 }

}


func (this *QTableWidget) setItemPrototype(args ...interface{}) () {
  // setItemPrototype(const class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "const QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget16setItemPrototypeEPK16QTableWidgetItem
  default:
    qtrt.ErrorResolve("QTableWidget", "setItemPrototype", args)
 }

}


func (this *QTableWidget) visualRow(args ...interface{}) () {
  // visualRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget9visualRowEi
  default:
    qtrt.ErrorResolve("QTableWidget", "visualRow", args)
 }

}


func (this *QTableWidget) setCellWidget(args ...interface{}) () {
  // setCellWidget(int, int, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget13setCellWidgetEiiP7QWidget
  default:
    qtrt.ErrorResolve("QTableWidget", "setCellWidget", args)
 }

}


func (this *QTableWidget) openPersistentEditor(args ...interface{}) () {
  // openPersistentEditor(class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget20openPersistentEditorEP16QTableWidgetItem
  default:
    qtrt.ErrorResolve("QTableWidget", "openPersistentEditor", args)
 }

}


func (this *QTableWidget) columnCount(args ...interface{}) () {
  // columnCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget11columnCountEv
  default:
    qtrt.ErrorResolve("QTableWidget", "columnCount", args)
 }

}


func (this *QTableWidget) currentRow(args ...interface{}) () {
  // currentRow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget10currentRowEv
  default:
    qtrt.ErrorResolve("QTableWidget", "currentRow", args)
 }

}


func (this *QTableWidget) setCurrentItem(args ...interface{}) () {
  // setCurrentItem(class QTableWidgetItem *, class QItemSelectionModel::SelectionFlags)
  // setCurrentItem(class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"
  vtys[0][1] = qtrt.Int64Ty(false) // "QItemSelectionModel::SelectionFlags"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget14setCurrentItemEP16QTableWidgetItem6QFlagsIN19QItemSelectionModel13SelectionFlagEE
  case 1:
    // invoke: _ZN12QTableWidget14setCurrentItemEP16QTableWidgetItem
  default:
    qtrt.ErrorResolve("QTableWidget", "setCurrentItem", args)
 }

}


func (this *QTableWidget) cellWidget(args ...interface{}) () {
  // cellWidget(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget10cellWidgetEii
  default:
    qtrt.ErrorResolve("QTableWidget", "cellWidget", args)
 }

}


func (this *QTableWidget) setSortingEnabled(args ...interface{}) () {
  // setSortingEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget17setSortingEnabledEb
  default:
    qtrt.ErrorResolve("QTableWidget", "setSortingEnabled", args)
 }

}


func (this *QTableWidget) setItem(args ...interface{}) () {
  // setItem(int, int, class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget7setItemEiiP16QTableWidgetItem
  default:
    qtrt.ErrorResolve("QTableWidget", "setItem", args)
 }

}


func (this *QTableWidget) horizontalHeaderItem(args ...interface{}) () {
  // horizontalHeaderItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget20horizontalHeaderItemEi
  default:
    qtrt.ErrorResolve("QTableWidget", "horizontalHeaderItem", args)
 }

}


func (this *QTableWidget) editItem(args ...interface{}) () {
  // editItem(class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget8editItemEP16QTableWidgetItem
  default:
    qtrt.ErrorResolve("QTableWidget", "editItem", args)
 }

}


func (this *QTableWidget) selectedRanges(args ...interface{}) () {
  // selectedRanges()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget14selectedRangesEv
  default:
    qtrt.ErrorResolve("QTableWidget", "selectedRanges", args)
 }

}


func (this *QTableWidget) currentColumn(args ...interface{}) () {
  // currentColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget13currentColumnEv
  default:
    qtrt.ErrorResolve("QTableWidget", "currentColumn", args)
 }

}


func (this *QTableWidget) removeColumn(args ...interface{}) () {
  // removeColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget12removeColumnEi
  default:
    qtrt.ErrorResolve("QTableWidget", "removeColumn", args)
 }

}


func (this *QTableWidget) setRangeSelected(args ...interface{}) () {
  // setRangeSelected(const class QTableWidgetSelectionRange &, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetSelectionRange{}) // "const QTableWidgetSelectionRange &"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget16setRangeSelectedERK26QTableWidgetSelectionRangeb
  default:
    qtrt.ErrorResolve("QTableWidget", "setRangeSelected", args)
 }

}


func (this *QTableWidget) column(args ...interface{}) () {
  // column(const class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "const QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget6columnEPK16QTableWidgetItem
  default:
    qtrt.ErrorResolve("QTableWidget", "column", args)
 }

}


func (this *QTableWidget) isItemSelected(args ...interface{}) () {
  // isItemSelected(const class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "const QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget14isItemSelectedEPK16QTableWidgetItem
  default:
    qtrt.ErrorResolve("QTableWidget", "isItemSelected", args)
 }

}


func (this *QTableWidget) takeVerticalHeaderItem(args ...interface{}) () {
  // takeVerticalHeaderItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget22takeVerticalHeaderItemEi
  default:
    qtrt.ErrorResolve("QTableWidget", "takeVerticalHeaderItem", args)
 }

}


func (this *QTableWidget) insertRow(args ...interface{}) () {
  // insertRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget9insertRowEi
  default:
    qtrt.ErrorResolve("QTableWidget", "insertRow", args)
 }

}


func (this *QTableWidget) rowCount(args ...interface{}) () {
  // rowCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget8rowCountEv
  default:
    qtrt.ErrorResolve("QTableWidget", "rowCount", args)
 }

}


func (this *QTableWidget) item(args ...interface{}) () {
  // item(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget4itemEii
  default:
    qtrt.ErrorResolve("QTableWidget", "item", args)
 }

}


func (this *QTableWidget) setVerticalHeaderLabels(args ...interface{}) () {
  // setVerticalHeaderLabels(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget23setVerticalHeaderLabelsERK11QStringList
  default:
    qtrt.ErrorResolve("QTableWidget", "setVerticalHeaderLabels", args)
 }

}


func (this *QTableWidget) itemPrototype(args ...interface{}) () {
  // itemPrototype()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget13itemPrototypeEv
  default:
    qtrt.ErrorResolve("QTableWidget", "itemPrototype", args)
 }

}


func (this *QTableWidget) itemAt(args ...interface{}) () {
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
    // invoke: _ZNK12QTableWidget6itemAtERK6QPoint
  case 1:
    // invoke: _ZNK12QTableWidget6itemAtEii
  default:
    qtrt.ErrorResolve("QTableWidget", "itemAt", args)
 }

}


func (this *QTableWidget) clearContents(args ...interface{}) () {
  // clearContents()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget13clearContentsEv
  default:
    qtrt.ErrorResolve("QTableWidget", "clearContents", args)
 }

}


func (this *QTableWidget) setCurrentCell(args ...interface{}) () {
  // setCurrentCell(int, int, class QItemSelectionModel::SelectionFlags)
  // setCurrentCell(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int64Ty(false) // "QItemSelectionModel::SelectionFlags"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget14setCurrentCellEii6QFlagsIN19QItemSelectionModel13SelectionFlagEE
  case 1:
    // invoke: _ZN12QTableWidget14setCurrentCellEii
  default:
    qtrt.ErrorResolve("QTableWidget", "setCurrentCell", args)
 }

}


func (this *QTableWidget) setRowCount(args ...interface{}) () {
  // setRowCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget11setRowCountEi
  default:
    qtrt.ErrorResolve("QTableWidget", "setRowCount", args)
 }

}


func (this *QTableWidget) setHorizontalHeaderItem(args ...interface{}) () {
  // setHorizontalHeaderItem(int, class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget23setHorizontalHeaderItemEiP16QTableWidgetItem
  default:
    qtrt.ErrorResolve("QTableWidget", "setHorizontalHeaderItem", args)
 }

}


func (this *QTableWidget) visualColumn(args ...interface{}) () {
  // visualColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget12visualColumnEi
  default:
    qtrt.ErrorResolve("QTableWidget", "visualColumn", args)
 }

}


func (this *QTableWidget) takeHorizontalHeaderItem(args ...interface{}) () {
  // takeHorizontalHeaderItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget24takeHorizontalHeaderItemEi
  default:
    qtrt.ErrorResolve("QTableWidget", "takeHorizontalHeaderItem", args)
 }

}


func (this *QTableWidget) verticalHeaderItem(args ...interface{}) () {
  // verticalHeaderItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget18verticalHeaderItemEi
  default:
    qtrt.ErrorResolve("QTableWidget", "verticalHeaderItem", args)
 }

}


func (this *QTableWidget) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget5clearEv
  default:
    qtrt.ErrorResolve("QTableWidget", "clear", args)
 }

}


func (this *QTableWidget) insertColumn(args ...interface{}) () {
  // insertColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget12insertColumnEi
  default:
    qtrt.ErrorResolve("QTableWidget", "insertColumn", args)
 }

}


func (this *QTableWidgetItem) backgroundColor(args ...interface{}) () {
  // backgroundColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem15backgroundColorEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "backgroundColor", args)
 }

}


func (this *QTableWidgetItem) data(args ...interface{}) () {
  // data(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem4dataEi
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "data", args)
 }

}


func (this *QTableWidgetItem) setSelected(args ...interface{}) () {
  // setSelected(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem11setSelectedEb
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setSelected", args)
 }

}


func (this *QTableWidgetItem) setStatusTip(args ...interface{}) () {
  // setStatusTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem12setStatusTipERK7QString
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setStatusTip", args)
 }

}


func (this *QTableWidgetItem) textColor(args ...interface{}) () {
  // textColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem9textColorEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "textColor", args)
 }

}


func (this *QTableWidgetItem) FreeQTableWidgetItem(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "~QTableWidgetItem", args)
 }

}


func (this *QTableWidgetItem) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem4textEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "text", args)
 }

}


func (this *QTableWidgetItem) setSizeHint(args ...interface{}) () {
  // setSizeHint(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem11setSizeHintERK5QSize
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setSizeHint", args)
 }

}


func (this *QTableWidgetItem) foreground(args ...interface{}) () {
  // foreground()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem10foregroundEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "foreground", args)
 }

}


func (this *QTableWidgetItem) type_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "type", args)
 }

}


func (this *QTableWidgetItem) column(args ...interface{}) () {
  // column()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem6columnEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "column", args)
 }

}


func (this *QTableWidgetItem) setTextAlignment(args ...interface{}) () {
  // setTextAlignment(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem16setTextAlignmentEi
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setTextAlignment", args)
 }

}


func (this *QTableWidgetItem) font(args ...interface{}) () {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem4fontEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "font", args)
 }

}


func (this *QTableWidgetItem) icon(args ...interface{}) () {
  // icon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem4iconEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "icon", args)
 }

}


func (this *QTableWidgetItem) write(args ...interface{}) () {
  // write(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDataStream{}) // "QDataStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem5writeER11QDataStream
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "write", args)
 }

}


func NewQTableWidgetItem(args ...interface{})() {
}


func (this *QTableWidgetItem) background(args ...interface{}) () {
  // background()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem10backgroundEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "background", args)
 }

}


func (this *QTableWidgetItem) setIcon(args ...interface{}) () {
  // setIcon(const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem7setIconERK5QIcon
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setIcon", args)
 }

}


func (this *QTableWidgetItem) statusTip(args ...interface{}) () {
  // statusTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem9statusTipEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "statusTip", args)
 }

}


func (this *QTableWidgetItem) clone(args ...interface{}) () {
  // clone()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem5cloneEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "clone", args)
 }

}


func (this *QTableWidgetItem) setWhatsThis(args ...interface{}) () {
  // setWhatsThis(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem12setWhatsThisERK7QString
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setWhatsThis", args)
 }

}


func (this *QTableWidgetItem) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem8sizeHintEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "sizeHint", args)
 }

}


func (this *QTableWidgetItem) setForeground(args ...interface{}) () {
  // setForeground(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem13setForegroundERK6QBrush
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setForeground", args)
 }

}


func (this *QTableWidgetItem) row(args ...interface{}) () {
  // row()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem3rowEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "row", args)
 }

}


func (this *QTableWidgetItem) setData(args ...interface{}) () {
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
    // invoke: _ZN16QTableWidgetItem7setDataEiRK8QVariant
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setData", args)
 }

}


func (this *QTableWidgetItem) tableWidget(args ...interface{}) () {
  // tableWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem11tableWidgetEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "tableWidget", args)
 }

}


func (this *QTableWidgetItem) textAlignment(args ...interface{}) () {
  // textAlignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem13textAlignmentEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "textAlignment", args)
 }

}


func (this *QTableWidgetItem) read(args ...interface{}) () {
  // read(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDataStream{}) // "QDataStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem4readER11QDataStream
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "read", args)
 }

}


func (this *QTableWidgetItem) toolTip(args ...interface{}) () {
  // toolTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem7toolTipEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "toolTip", args)
 }

}


func (this *QTableWidgetItem) isSelected(args ...interface{}) () {
  // isSelected()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem10isSelectedEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "isSelected", args)
 }

}


func (this *QTableWidgetItem) setBackgroundColor(args ...interface{}) () {
  // setBackgroundColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem18setBackgroundColorERK6QColor
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setBackgroundColor", args)
 }

}


func (this *QTableWidgetItem) setBackground(args ...interface{}) () {
  // setBackground(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem13setBackgroundERK6QBrush
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setBackground", args)
 }

}


func (this *QTableWidgetItem) setFont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem7setFontERK5QFont
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setFont", args)
 }

}


func (this *QTableWidgetItem) setTextColor(args ...interface{}) () {
  // setTextColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem12setTextColorERK6QColor
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setTextColor", args)
 }

}


func (this *QTableWidgetItem) setText(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem7setTextERK7QString
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setText", args)
 }

}


func (this *QTableWidgetItem) whatsThis(args ...interface{}) () {
  // whatsThis()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem9whatsThisEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "whatsThis", args)
 }

}


func (this *QTableWidgetItem) setToolTip(args ...interface{}) () {
  // setToolTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem10setToolTipERK7QString
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setToolTip", args)
 }

}

// <= body block end

