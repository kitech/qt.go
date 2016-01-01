package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtWidgets/qtreeview.h
// dst-file: /src/widgets/qtreeview.go
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
// class sizeof(QTreeView)=1
type QTreeView struct {
  /*qbase*/ QAbstractItemView;
  qclsinst uint64 /* *mut c_void*/;
//  _collapsed QTreeView_collapsed_signal;
//  _expanded QTreeView_expanded_signal;
}


func (this *QTreeView) setHeader(args ...interface{}) () {
  // setHeader(class QHeaderView *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QHeaderView{}) // "QHeaderView *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView9setHeaderEP11QHeaderView
  default:
    qtrt.ErrorResolve("QTreeView", "setHeader", args)
 }

}


func (this *QTreeView) isAnimated(args ...interface{}) () {
  // isAnimated()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView10isAnimatedEv
  default:
    qtrt.ErrorResolve("QTreeView", "isAnimated", args)
 }

}


func (this *QTreeView) isExpanded(args ...interface{}) () {
  // isExpanded(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView10isExpandedERK11QModelIndex
  default:
    qtrt.ErrorResolve("QTreeView", "isExpanded", args)
 }

}


func (this *QTreeView) setColumnHidden(args ...interface{}) () {
  // setColumnHidden(int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView15setColumnHiddenEib
  default:
    qtrt.ErrorResolve("QTreeView", "setColumnHidden", args)
 }

}


func (this *QTreeView) setIndentation(args ...interface{}) () {
  // setIndentation(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView14setIndentationEi
  default:
    qtrt.ErrorResolve("QTreeView", "setIndentation", args)
 }

}


func (this *QTreeView) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView10metaObjectEv
  default:
    qtrt.ErrorResolve("QTreeView", "metaObject", args)
 }

}


func (this *QTreeView) reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView5resetEv
  default:
    qtrt.ErrorResolve("QTreeView", "reset", args)
 }

}


func (this *QTreeView) setExpandsOnDoubleClick(args ...interface{}) () {
  // setExpandsOnDoubleClick(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView23setExpandsOnDoubleClickEb
  default:
    qtrt.ErrorResolve("QTreeView", "setExpandsOnDoubleClick", args)
 }

}


func (this *QTreeView) setFirstColumnSpanned(args ...interface{}) () {
  // setFirstColumnSpanned(int, const class QModelIndex &, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][2] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView21setFirstColumnSpannedEiRK11QModelIndexb
  default:
    qtrt.ErrorResolve("QTreeView", "setFirstColumnSpanned", args)
 }

}


func (this *QTreeView) sortByColumn(args ...interface{}) () {
  // sortByColumn(int)
  // sortByColumn(int, Qt::SortOrder)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "Qt::SortOrder"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView12sortByColumnEi
  case 1:
    // invoke: _ZN9QTreeView12sortByColumnEiN2Qt9SortOrderE
  default:
    qtrt.ErrorResolve("QTreeView", "sortByColumn", args)
 }

}


func (this *QTreeView) setRowHidden(args ...interface{}) () {
  // setRowHidden(int, const class QModelIndex &, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][2] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView12setRowHiddenEiRK11QModelIndexb
  default:
    qtrt.ErrorResolve("QTreeView", "setRowHidden", args)
 }

}


func (this *QTreeView) expand(args ...interface{}) () {
  // expand(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView6expandERK11QModelIndex
  default:
    qtrt.ErrorResolve("QTreeView", "expand", args)
 }

}


func (this *QTreeView) autoExpandDelay(args ...interface{}) () {
  // autoExpandDelay()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView15autoExpandDelayEv
  default:
    qtrt.ErrorResolve("QTreeView", "autoExpandDelay", args)
 }

}


func NewQTreeView(args ...interface{})() {
}


func (this *QTreeView) FreeQTreeView(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTreeView", "~QTreeView", args)
 }

}


func (this *QTreeView) indentation(args ...interface{}) () {
  // indentation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView11indentationEv
  default:
    qtrt.ErrorResolve("QTreeView", "indentation", args)
 }

}


func (this *QTreeView) columnViewportPosition(args ...interface{}) () {
  // columnViewportPosition(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView22columnViewportPositionEi
  default:
    qtrt.ErrorResolve("QTreeView", "columnViewportPosition", args)
 }

}


func (this *QTreeView) expandsOnDoubleClick(args ...interface{}) () {
  // expandsOnDoubleClick()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView20expandsOnDoubleClickEv
  default:
    qtrt.ErrorResolve("QTreeView", "expandsOnDoubleClick", args)
 }

}


func (this *QTreeView) isSortingEnabled(args ...interface{}) () {
  // isSortingEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView16isSortingEnabledEv
  default:
    qtrt.ErrorResolve("QTreeView", "isSortingEnabled", args)
 }

}


func (this *QTreeView) showColumn(args ...interface{}) () {
  // showColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView10showColumnEi
  default:
    qtrt.ErrorResolve("QTreeView", "showColumn", args)
 }

}


func (this *QTreeView) visualRect(args ...interface{}) () {
  // visualRect(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView10visualRectERK11QModelIndex
  default:
    qtrt.ErrorResolve("QTreeView", "visualRect", args)
 }

}


func (this *QTreeView) collapse(args ...interface{}) () {
  // collapse(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView8collapseERK11QModelIndex
  default:
    qtrt.ErrorResolve("QTreeView", "collapse", args)
 }

}


func (this *QTreeView) setWordWrap(args ...interface{}) () {
  // setWordWrap(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView11setWordWrapEb
  default:
    qtrt.ErrorResolve("QTreeView", "setWordWrap", args)
 }

}


func (this *QTreeView) indexAbove(args ...interface{}) () {
  // indexAbove(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView10indexAboveERK11QModelIndex
  default:
    qtrt.ErrorResolve("QTreeView", "indexAbove", args)
 }

}


func (this *QTreeView) rootIsDecorated(args ...interface{}) () {
  // rootIsDecorated()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView15rootIsDecoratedEv
  default:
    qtrt.ErrorResolve("QTreeView", "rootIsDecorated", args)
 }

}


func (this *QTreeView) collapseAll(args ...interface{}) () {
  // collapseAll()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView11collapseAllEv
  default:
    qtrt.ErrorResolve("QTreeView", "collapseAll", args)
 }

}


func (this *QTreeView) setHeaderHidden(args ...interface{}) () {
  // setHeaderHidden(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView15setHeaderHiddenEb
  default:
    qtrt.ErrorResolve("QTreeView", "setHeaderHidden", args)
 }

}


func (this *QTreeView) allColumnsShowFocus(args ...interface{}) () {
  // allColumnsShowFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView19allColumnsShowFocusEv
  default:
    qtrt.ErrorResolve("QTreeView", "allColumnsShowFocus", args)
 }

}


func (this *QTreeView) columnWidth(args ...interface{}) () {
  // columnWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView11columnWidthEi
  default:
    qtrt.ErrorResolve("QTreeView", "columnWidth", args)
 }

}


func (this *QTreeView) resizeColumnToContents(args ...interface{}) () {
  // resizeColumnToContents(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView22resizeColumnToContentsEi
  default:
    qtrt.ErrorResolve("QTreeView", "resizeColumnToContents", args)
 }

}


func (this *QTreeView) setAutoExpandDelay(args ...interface{}) () {
  // setAutoExpandDelay(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView18setAutoExpandDelayEi
  default:
    qtrt.ErrorResolve("QTreeView", "setAutoExpandDelay", args)
 }

}


func (this *QTreeView) setAllColumnsShowFocus(args ...interface{}) () {
  // setAllColumnsShowFocus(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView22setAllColumnsShowFocusEb
  default:
    qtrt.ErrorResolve("QTreeView", "setAllColumnsShowFocus", args)
 }

}


func (this *QTreeView) isFirstColumnSpanned(args ...interface{}) () {
  // isFirstColumnSpanned(int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView20isFirstColumnSpannedEiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QTreeView", "isFirstColumnSpanned", args)
 }

}


func (this *QTreeView) hideColumn(args ...interface{}) () {
  // hideColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView10hideColumnEi
  default:
    qtrt.ErrorResolve("QTreeView", "hideColumn", args)
 }

}


func (this *QTreeView) treePosition(args ...interface{}) () {
  // treePosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView12treePositionEv
  default:
    qtrt.ErrorResolve("QTreeView", "treePosition", args)
 }

}


func (this *QTreeView) setExpanded(args ...interface{}) () {
  // setExpanded(const class QModelIndex &, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView11setExpandedERK11QModelIndexb
  default:
    qtrt.ErrorResolve("QTreeView", "setExpanded", args)
 }

}


func (this *QTreeView) resetIndentation(args ...interface{}) () {
  // resetIndentation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView16resetIndentationEv
  default:
    qtrt.ErrorResolve("QTreeView", "resetIndentation", args)
 }

}


func (this *QTreeView) isRowHidden(args ...interface{}) () {
  // isRowHidden(int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView11isRowHiddenEiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QTreeView", "isRowHidden", args)
 }

}


func (this *QTreeView) selectAll(args ...interface{}) () {
  // selectAll()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView9selectAllEv
  default:
    qtrt.ErrorResolve("QTreeView", "selectAll", args)
 }

}


func (this *QTreeView) wordWrap(args ...interface{}) () {
  // wordWrap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView8wordWrapEv
  default:
    qtrt.ErrorResolve("QTreeView", "wordWrap", args)
 }

}


func (this *QTreeView) doItemsLayout(args ...interface{}) () {
  // doItemsLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView13doItemsLayoutEv
  default:
    qtrt.ErrorResolve("QTreeView", "doItemsLayout", args)
 }

}


func (this *QTreeView) setTreePosition(args ...interface{}) () {
  // setTreePosition(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView15setTreePositionEi
  default:
    qtrt.ErrorResolve("QTreeView", "setTreePosition", args)
 }

}


func (this *QTreeView) keyboardSearch(args ...interface{}) () {
  // keyboardSearch(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView14keyboardSearchERK7QString
  default:
    qtrt.ErrorResolve("QTreeView", "keyboardSearch", args)
 }

}


func (this *QTreeView) setRootIndex(args ...interface{}) () {
  // setRootIndex(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView12setRootIndexERK11QModelIndex
  default:
    qtrt.ErrorResolve("QTreeView", "setRootIndex", args)
 }

}


func (this *QTreeView) setItemsExpandable(args ...interface{}) () {
  // setItemsExpandable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView18setItemsExpandableEb
  default:
    qtrt.ErrorResolve("QTreeView", "setItemsExpandable", args)
 }

}


func (this *QTreeView) setSelectionModel(args ...interface{}) () {
  // setSelectionModel(class QItemSelectionModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QItemSelectionModel{}) // "QItemSelectionModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView17setSelectionModelEP19QItemSelectionModel
  default:
    qtrt.ErrorResolve("QTreeView", "setSelectionModel", args)
 }

}


func (this *QTreeView) header(args ...interface{}) () {
  // header()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView6headerEv
  default:
    qtrt.ErrorResolve("QTreeView", "header", args)
 }

}


func (this *QTreeView) setAnimated(args ...interface{}) () {
  // setAnimated(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView11setAnimatedEb
  default:
    qtrt.ErrorResolve("QTreeView", "setAnimated", args)
 }

}


func (this *QTreeView) setSortingEnabled(args ...interface{}) () {
  // setSortingEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView17setSortingEnabledEb
  default:
    qtrt.ErrorResolve("QTreeView", "setSortingEnabled", args)
 }

}


func (this *QTreeView) itemsExpandable(args ...interface{}) () {
  // itemsExpandable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView15itemsExpandableEv
  default:
    qtrt.ErrorResolve("QTreeView", "itemsExpandable", args)
 }

}


func (this *QTreeView) setRootIsDecorated(args ...interface{}) () {
  // setRootIsDecorated(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView18setRootIsDecoratedEb
  default:
    qtrt.ErrorResolve("QTreeView", "setRootIsDecorated", args)
 }

}


func (this *QTreeView) isHeaderHidden(args ...interface{}) () {
  // isHeaderHidden()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView14isHeaderHiddenEv
  default:
    qtrt.ErrorResolve("QTreeView", "isHeaderHidden", args)
 }

}


func (this *QTreeView) columnAt(args ...interface{}) () {
  // columnAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView8columnAtEi
  default:
    qtrt.ErrorResolve("QTreeView", "columnAt", args)
 }

}


func (this *QTreeView) setModel(args ...interface{}) () {
  // setModel(class QAbstractItemModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemModel{}) // "QAbstractItemModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView8setModelEP18QAbstractItemModel
  default:
    qtrt.ErrorResolve("QTreeView", "setModel", args)
 }

}


func (this *QTreeView) isColumnHidden(args ...interface{}) () {
  // isColumnHidden(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView14isColumnHiddenEi
  default:
    qtrt.ErrorResolve("QTreeView", "isColumnHidden", args)
 }

}


func (this *QTreeView) uniformRowHeights(args ...interface{}) () {
  // uniformRowHeights()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView17uniformRowHeightsEv
  default:
    qtrt.ErrorResolve("QTreeView", "uniformRowHeights", args)
 }

}


func (this *QTreeView) setUniformRowHeights(args ...interface{}) () {
  // setUniformRowHeights(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView20setUniformRowHeightsEb
  default:
    qtrt.ErrorResolve("QTreeView", "setUniformRowHeights", args)
 }

}


func (this *QTreeView) expandToDepth(args ...interface{}) () {
  // expandToDepth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView13expandToDepthEi
  default:
    qtrt.ErrorResolve("QTreeView", "expandToDepth", args)
 }

}


func (this *QTreeView) indexBelow(args ...interface{}) () {
  // indexBelow(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView10indexBelowERK11QModelIndex
  default:
    qtrt.ErrorResolve("QTreeView", "indexBelow", args)
 }

}


func (this *QTreeView) expandAll(args ...interface{}) () {
  // expandAll()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView9expandAllEv
  default:
    qtrt.ErrorResolve("QTreeView", "expandAll", args)
 }

}


func (this *QTreeView) indexAt(args ...interface{}) () {
  // indexAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView7indexAtERK6QPoint
  default:
    qtrt.ErrorResolve("QTreeView", "indexAt", args)
 }

}


func (this *QTreeView) setColumnWidth(args ...interface{}) () {
  // setColumnWidth(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView14setColumnWidthEii
  default:
    qtrt.ErrorResolve("QTreeView", "setColumnWidth", args)
 }

}

// <= body block end

