package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtWidgets/qtreeview.h
// dst-file: /src/widgets/qtreeview.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QTreeView::setHeader(QHeaderView * header);
extern void _ZN9QTreeView9setHeaderEP11QHeaderView(void* qthis, void* arg0);
  // proto:  bool QTreeView::isAnimated();
extern void _ZNK9QTreeView10isAnimatedEv(void* qthis);
  // proto:  bool QTreeView::isExpanded(const QModelIndex & index);
extern void _ZNK9QTreeView10isExpandedERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QTreeView::setColumnHidden(int column, bool hide);
extern void _ZN9QTreeView15setColumnHiddenEib(void* qthis, int32_t arg0, bool arg1);
  // proto:  void QTreeView::setIndentation(int i);
extern void _ZN9QTreeView14setIndentationEi(void* qthis, int32_t arg0);
  // proto:  const QMetaObject * QTreeView::metaObject();
extern void _ZNK9QTreeView10metaObjectEv(void* qthis);
  // proto:  void QTreeView::reset();
extern void _ZN9QTreeView5resetEv(void* qthis);
  // proto:  void QTreeView::setExpandsOnDoubleClick(bool enable);
extern void _ZN9QTreeView23setExpandsOnDoubleClickEb(void* qthis, bool arg0);
  // proto:  void QTreeView::setFirstColumnSpanned(int row, const QModelIndex & parent, bool span);
extern void _ZN9QTreeView21setFirstColumnSpannedEiRK11QModelIndexb(void* qthis, int32_t arg0, void* arg1, bool arg2);
  // proto:  void QTreeView::sortByColumn(int column);
extern void _ZN9QTreeView12sortByColumnEi(void* qthis, int32_t arg0);
  // proto:  void QTreeView::setRowHidden(int row, const QModelIndex & parent, bool hide);
extern void _ZN9QTreeView12setRowHiddenEiRK11QModelIndexb(void* qthis, int32_t arg0, void* arg1, bool arg2);
  // proto:  void QTreeView::expand(const QModelIndex & index);
extern void _ZN9QTreeView6expandERK11QModelIndex(void* qthis, void* arg0);
  // proto:  int QTreeView::autoExpandDelay();
extern void _ZNK9QTreeView15autoExpandDelayEv(void* qthis);
  // proto:  void QTreeView::QTreeView(QWidget * parent);
extern void* dector_ZN9QTreeViewC1EP7QWidget(void* arg0);
extern void _ZN9QTreeViewC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QTreeView::~QTreeView();
extern void _ZN9QTreeViewD0Ev(void* qthis);
  // proto:  int QTreeView::indentation();
extern void _ZNK9QTreeView11indentationEv(void* qthis);
  // proto:  int QTreeView::columnViewportPosition(int column);
extern void _ZNK9QTreeView22columnViewportPositionEi(void* qthis, int32_t arg0);
  // proto:  bool QTreeView::expandsOnDoubleClick();
extern void _ZNK9QTreeView20expandsOnDoubleClickEv(void* qthis);
  // proto:  bool QTreeView::isSortingEnabled();
extern void _ZNK9QTreeView16isSortingEnabledEv(void* qthis);
  // proto:  void QTreeView::showColumn(int column);
extern void _ZN9QTreeView10showColumnEi(void* qthis, int32_t arg0);
  // proto:  QRect QTreeView::visualRect(const QModelIndex & index);
extern void _ZNK9QTreeView10visualRectERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QTreeView::collapse(const QModelIndex & index);
extern void _ZN9QTreeView8collapseERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QTreeView::setWordWrap(bool on);
extern void _ZN9QTreeView11setWordWrapEb(void* qthis, bool arg0);
  // proto:  QModelIndex QTreeView::indexAbove(const QModelIndex & index);
extern void _ZNK9QTreeView10indexAboveERK11QModelIndex(void* qthis, void* arg0);
  // proto:  bool QTreeView::rootIsDecorated();
extern void _ZNK9QTreeView15rootIsDecoratedEv(void* qthis);
  // proto:  void QTreeView::collapseAll();
extern void _ZN9QTreeView11collapseAllEv(void* qthis);
  // proto:  void QTreeView::setHeaderHidden(bool hide);
extern void _ZN9QTreeView15setHeaderHiddenEb(void* qthis, bool arg0);
  // proto:  bool QTreeView::allColumnsShowFocus();
extern void _ZNK9QTreeView19allColumnsShowFocusEv(void* qthis);
  // proto:  int QTreeView::columnWidth(int column);
extern void _ZNK9QTreeView11columnWidthEi(void* qthis, int32_t arg0);
  // proto:  void QTreeView::resizeColumnToContents(int column);
extern void _ZN9QTreeView22resizeColumnToContentsEi(void* qthis, int32_t arg0);
  // proto:  void QTreeView::setAutoExpandDelay(int delay);
extern void _ZN9QTreeView18setAutoExpandDelayEi(void* qthis, int32_t arg0);
  // proto:  void QTreeView::setAllColumnsShowFocus(bool enable);
extern void _ZN9QTreeView22setAllColumnsShowFocusEb(void* qthis, bool arg0);
  // proto:  bool QTreeView::isFirstColumnSpanned(int row, const QModelIndex & parent);
extern void _ZNK9QTreeView20isFirstColumnSpannedEiRK11QModelIndex(void* qthis, int32_t arg0, void* arg1);
  // proto:  void QTreeView::hideColumn(int column);
extern void _ZN9QTreeView10hideColumnEi(void* qthis, int32_t arg0);
  // proto:  int QTreeView::treePosition();
extern void _ZNK9QTreeView12treePositionEv(void* qthis);
  // proto:  void QTreeView::setExpanded(const QModelIndex & index, bool expand);
extern void _ZN9QTreeView11setExpandedERK11QModelIndexb(void* qthis, void* arg0, bool arg1);
  // proto:  void QTreeView::resetIndentation();
extern void _ZN9QTreeView16resetIndentationEv(void* qthis);
  // proto:  bool QTreeView::isRowHidden(int row, const QModelIndex & parent);
extern void _ZNK9QTreeView11isRowHiddenEiRK11QModelIndex(void* qthis, int32_t arg0, void* arg1);
  // proto:  void QTreeView::QTreeView(const QTreeView & );
extern void* dector_ZN9QTreeViewC1ERKS_(void* arg0);
extern void _ZN9QTreeViewC1ERKS_(void* qthis, void* arg0);
  // proto:  void QTreeView::selectAll();
extern void _ZN9QTreeView9selectAllEv(void* qthis);
  // proto:  bool QTreeView::wordWrap();
extern void _ZNK9QTreeView8wordWrapEv(void* qthis);
  // proto:  void QTreeView::doItemsLayout();
extern void _ZN9QTreeView13doItemsLayoutEv(void* qthis);
  // proto:  void QTreeView::setTreePosition(int logicalIndex);
extern void _ZN9QTreeView15setTreePositionEi(void* qthis, int32_t arg0);
  // proto:  void QTreeView::keyboardSearch(const QString & search);
extern void _ZN9QTreeView14keyboardSearchERK7QString(void* qthis, void* arg0);
  // proto:  void QTreeView::setRootIndex(const QModelIndex & index);
extern void _ZN9QTreeView12setRootIndexERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QTreeView::setItemsExpandable(bool enable);
extern void _ZN9QTreeView18setItemsExpandableEb(void* qthis, bool arg0);
  // proto:  void QTreeView::setSelectionModel(QItemSelectionModel * selectionModel);
extern void _ZN9QTreeView17setSelectionModelEP19QItemSelectionModel(void* qthis, void* arg0);
  // proto:  QHeaderView * QTreeView::header();
extern void _ZNK9QTreeView6headerEv(void* qthis);
  // proto:  void QTreeView::setAnimated(bool enable);
extern void _ZN9QTreeView11setAnimatedEb(void* qthis, bool arg0);
  // proto:  void QTreeView::setSortingEnabled(bool enable);
extern void _ZN9QTreeView17setSortingEnabledEb(void* qthis, bool arg0);
  // proto:  bool QTreeView::itemsExpandable();
extern void _ZNK9QTreeView15itemsExpandableEv(void* qthis);
  // proto:  void QTreeView::setRootIsDecorated(bool show);
extern void _ZN9QTreeView18setRootIsDecoratedEb(void* qthis, bool arg0);
  // proto:  bool QTreeView::isHeaderHidden();
extern void _ZNK9QTreeView14isHeaderHiddenEv(void* qthis);
  // proto:  int QTreeView::columnAt(int x);
extern void _ZNK9QTreeView8columnAtEi(void* qthis, int32_t arg0);
  // proto:  void QTreeView::setModel(QAbstractItemModel * model);
extern void _ZN9QTreeView8setModelEP18QAbstractItemModel(void* qthis, void* arg0);
  // proto:  bool QTreeView::isColumnHidden(int column);
extern void _ZNK9QTreeView14isColumnHiddenEi(void* qthis, int32_t arg0);
  // proto:  bool QTreeView::uniformRowHeights();
extern void _ZNK9QTreeView17uniformRowHeightsEv(void* qthis);
  // proto:  void QTreeView::setUniformRowHeights(bool uniform);
extern void _ZN9QTreeView20setUniformRowHeightsEb(void* qthis, bool arg0);
  // proto:  void QTreeView::expandToDepth(int depth);
extern void _ZN9QTreeView13expandToDepthEi(void* qthis, int32_t arg0);
  // proto:  QModelIndex QTreeView::indexBelow(const QModelIndex & index);
extern void _ZNK9QTreeView10indexBelowERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QTreeView::expandAll();
extern void _ZN9QTreeView9expandAllEv(void* qthis);
  // proto:  QModelIndex QTreeView::indexAt(const QPoint & p);
extern void _ZNK9QTreeView7indexAtERK6QPoint(void* qthis, void* arg0);
  // proto:  void QTreeView::setColumnWidth(int column, int width);
extern void _ZN9QTreeView14setColumnWidthEii(void* qthis, int32_t arg0, int32_t arg1);
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QTreeView)=1
type QTreeView struct {
  /*qbase*/ QAbstractItemView;
  qclsinst unsafe.Pointer /* *C.void */;
//  _collapsed QTreeView_collapsed_signal;
//  _expanded QTreeView_expanded_signal;
}

  // proto:  void QTreeView::setHeader(QHeaderView * header);
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
    // invoke: void setHeader(class QHeaderView *)
    var arg0 = args[0].(QHeaderView).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTreeView9setHeaderEP11QHeaderView(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setHeader", args)
  }

}

  // proto:  bool QTreeView::isAnimated();
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
    // invoke: bool isAnimated()
    C._ZNK9QTreeView10isAnimatedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "isAnimated", args)
  }

}

  // proto:  bool QTreeView::isExpanded(const QModelIndex & index);
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
    // invoke: bool isExpanded(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK9QTreeView10isExpandedERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "isExpanded", args)
  }

}

  // proto:  void QTreeView::setColumnHidden(int column, bool hide);
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
    // invoke: void setColumnHidden(int, _Bool)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C._ZN9QTreeView15setColumnHiddenEib(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeView", "setColumnHidden", args)
  }

}

  // proto:  void QTreeView::setIndentation(int i);
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
    // invoke: void setIndentation(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QTreeView14setIndentationEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setIndentation", args)
  }

}

  // proto:  const QMetaObject * QTreeView::metaObject();
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
    // invoke: const QMetaObject * metaObject()
    C._ZNK9QTreeView10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "metaObject", args)
  }

}

  // proto:  void QTreeView::reset();
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
    // invoke: void reset()
    C._ZN9QTreeView5resetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "reset", args)
  }

}

  // proto:  void QTreeView::setExpandsOnDoubleClick(bool enable);
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
    // invoke: void setExpandsOnDoubleClick(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QTreeView23setExpandsOnDoubleClickEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setExpandsOnDoubleClick", args)
  }

}

  // proto:  void QTreeView::setFirstColumnSpanned(int row, const QModelIndex & parent, bool span);
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
    // invoke: void setFirstColumnSpanned(int, const class QModelIndex &, _Bool)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QModelIndex).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.bool(args[2].(bool))
    if false {fmt.Println(arg2)}
    C._ZN9QTreeView21setFirstColumnSpannedEiRK11QModelIndexb(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QTreeView", "setFirstColumnSpanned", args)
  }

}

  // proto:  void QTreeView::sortByColumn(int column);
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
    // invoke: void sortByColumn(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QTreeView12sortByColumnEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "sortByColumn", args)
  }

}

  // proto:  void QTreeView::setRowHidden(int row, const QModelIndex & parent, bool hide);
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
    // invoke: void setRowHidden(int, const class QModelIndex &, _Bool)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QModelIndex).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.bool(args[2].(bool))
    if false {fmt.Println(arg2)}
    C._ZN9QTreeView12setRowHiddenEiRK11QModelIndexb(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QTreeView", "setRowHidden", args)
  }

}

  // proto:  void QTreeView::expand(const QModelIndex & index);
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
    // invoke: void expand(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTreeView6expandERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "expand", args)
  }

}

  // proto:  int QTreeView::autoExpandDelay();
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
    // invoke: int autoExpandDelay()
    C._ZNK9QTreeView15autoExpandDelayEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "autoExpandDelay", args)
  }

}

  // proto:  void QTreeView::QTreeView(QWidget * parent);
func NewQTreeView(args ...interface{}) QTreeView {
  return QTreeView{}
}

  // proto:  void QTreeView::~QTreeView();
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

  // proto:  int QTreeView::indentation();
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
    // invoke: int indentation()
    C._ZNK9QTreeView11indentationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "indentation", args)
  }

}

  // proto:  int QTreeView::columnViewportPosition(int column);
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
    // invoke: int columnViewportPosition(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK9QTreeView22columnViewportPositionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "columnViewportPosition", args)
  }

}

  // proto:  bool QTreeView::expandsOnDoubleClick();
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
    // invoke: bool expandsOnDoubleClick()
    C._ZNK9QTreeView20expandsOnDoubleClickEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "expandsOnDoubleClick", args)
  }

}

  // proto:  bool QTreeView::isSortingEnabled();
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
    // invoke: bool isSortingEnabled()
    C._ZNK9QTreeView16isSortingEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "isSortingEnabled", args)
  }

}

  // proto:  void QTreeView::showColumn(int column);
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
    // invoke: void showColumn(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QTreeView10showColumnEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "showColumn", args)
  }

}

  // proto:  QRect QTreeView::visualRect(const QModelIndex & index);
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
    // invoke: QRect visualRect(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK9QTreeView10visualRectERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "visualRect", args)
  }

}

  // proto:  void QTreeView::collapse(const QModelIndex & index);
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
    // invoke: void collapse(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTreeView8collapseERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "collapse", args)
  }

}

  // proto:  void QTreeView::setWordWrap(bool on);
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
    // invoke: void setWordWrap(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QTreeView11setWordWrapEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setWordWrap", args)
  }

}

  // proto:  QModelIndex QTreeView::indexAbove(const QModelIndex & index);
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
    // invoke: QModelIndex indexAbove(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK9QTreeView10indexAboveERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "indexAbove", args)
  }

}

  // proto:  bool QTreeView::rootIsDecorated();
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
    // invoke: bool rootIsDecorated()
    C._ZNK9QTreeView15rootIsDecoratedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "rootIsDecorated", args)
  }

}

  // proto:  void QTreeView::collapseAll();
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
    // invoke: void collapseAll()
    C._ZN9QTreeView11collapseAllEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "collapseAll", args)
  }

}

  // proto:  void QTreeView::setHeaderHidden(bool hide);
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
    // invoke: void setHeaderHidden(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QTreeView15setHeaderHiddenEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setHeaderHidden", args)
  }

}

  // proto:  bool QTreeView::allColumnsShowFocus();
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
    // invoke: bool allColumnsShowFocus()
    C._ZNK9QTreeView19allColumnsShowFocusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "allColumnsShowFocus", args)
  }

}

  // proto:  int QTreeView::columnWidth(int column);
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
    // invoke: int columnWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK9QTreeView11columnWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "columnWidth", args)
  }

}

  // proto:  void QTreeView::resizeColumnToContents(int column);
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
    // invoke: void resizeColumnToContents(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QTreeView22resizeColumnToContentsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "resizeColumnToContents", args)
  }

}

  // proto:  void QTreeView::setAutoExpandDelay(int delay);
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
    // invoke: void setAutoExpandDelay(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QTreeView18setAutoExpandDelayEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setAutoExpandDelay", args)
  }

}

  // proto:  void QTreeView::setAllColumnsShowFocus(bool enable);
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
    // invoke: void setAllColumnsShowFocus(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QTreeView22setAllColumnsShowFocusEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setAllColumnsShowFocus", args)
  }

}

  // proto:  bool QTreeView::isFirstColumnSpanned(int row, const QModelIndex & parent);
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
    // invoke: bool isFirstColumnSpanned(int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QModelIndex).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK9QTreeView20isFirstColumnSpannedEiRK11QModelIndex(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeView", "isFirstColumnSpanned", args)
  }

}

  // proto:  void QTreeView::hideColumn(int column);
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
    // invoke: void hideColumn(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QTreeView10hideColumnEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "hideColumn", args)
  }

}

  // proto:  int QTreeView::treePosition();
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
    // invoke: int treePosition()
    C._ZNK9QTreeView12treePositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "treePosition", args)
  }

}

  // proto:  void QTreeView::setExpanded(const QModelIndex & index, bool expand);
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
    // invoke: void setExpanded(const class QModelIndex &, _Bool)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C._ZN9QTreeView11setExpandedERK11QModelIndexb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeView", "setExpanded", args)
  }

}

  // proto:  void QTreeView::resetIndentation();
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
    // invoke: void resetIndentation()
    C._ZN9QTreeView16resetIndentationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "resetIndentation", args)
  }

}

  // proto:  bool QTreeView::isRowHidden(int row, const QModelIndex & parent);
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
    // invoke: bool isRowHidden(int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QModelIndex).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK9QTreeView11isRowHiddenEiRK11QModelIndex(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeView", "isRowHidden", args)
  }

}

  // proto:  void QTreeView::selectAll();
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
    // invoke: void selectAll()
    C._ZN9QTreeView9selectAllEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "selectAll", args)
  }

}

  // proto:  bool QTreeView::wordWrap();
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
    // invoke: bool wordWrap()
    C._ZNK9QTreeView8wordWrapEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "wordWrap", args)
  }

}

  // proto:  void QTreeView::doItemsLayout();
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
    // invoke: void doItemsLayout()
    C._ZN9QTreeView13doItemsLayoutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "doItemsLayout", args)
  }

}

  // proto:  void QTreeView::setTreePosition(int logicalIndex);
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
    // invoke: void setTreePosition(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QTreeView15setTreePositionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setTreePosition", args)
  }

}

  // proto:  void QTreeView::keyboardSearch(const QString & search);
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
    // invoke: void keyboardSearch(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTreeView14keyboardSearchERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "keyboardSearch", args)
  }

}

  // proto:  void QTreeView::setRootIndex(const QModelIndex & index);
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
    // invoke: void setRootIndex(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTreeView12setRootIndexERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setRootIndex", args)
  }

}

  // proto:  void QTreeView::setItemsExpandable(bool enable);
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
    // invoke: void setItemsExpandable(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QTreeView18setItemsExpandableEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setItemsExpandable", args)
  }

}

  // proto:  void QTreeView::setSelectionModel(QItemSelectionModel * selectionModel);
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
    // invoke: void setSelectionModel(class QItemSelectionModel *)
    var arg0 = args[0].(QItemSelectionModel).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTreeView17setSelectionModelEP19QItemSelectionModel(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setSelectionModel", args)
  }

}

  // proto:  QHeaderView * QTreeView::header();
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
    // invoke: QHeaderView * header()
    C._ZNK9QTreeView6headerEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "header", args)
  }

}

  // proto:  void QTreeView::setAnimated(bool enable);
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
    // invoke: void setAnimated(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QTreeView11setAnimatedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setAnimated", args)
  }

}

  // proto:  void QTreeView::setSortingEnabled(bool enable);
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
    // invoke: void setSortingEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QTreeView17setSortingEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setSortingEnabled", args)
  }

}

  // proto:  bool QTreeView::itemsExpandable();
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
    // invoke: bool itemsExpandable()
    C._ZNK9QTreeView15itemsExpandableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "itemsExpandable", args)
  }

}

  // proto:  void QTreeView::setRootIsDecorated(bool show);
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
    // invoke: void setRootIsDecorated(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QTreeView18setRootIsDecoratedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setRootIsDecorated", args)
  }

}

  // proto:  bool QTreeView::isHeaderHidden();
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
    // invoke: bool isHeaderHidden()
    C._ZNK9QTreeView14isHeaderHiddenEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "isHeaderHidden", args)
  }

}

  // proto:  int QTreeView::columnAt(int x);
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
    // invoke: int columnAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK9QTreeView8columnAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "columnAt", args)
  }

}

  // proto:  void QTreeView::setModel(QAbstractItemModel * model);
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
    // invoke: void setModel(class QAbstractItemModel *)
    var arg0 = args[0].(QAbstractItemModel).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTreeView8setModelEP18QAbstractItemModel(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setModel", args)
  }

}

  // proto:  bool QTreeView::isColumnHidden(int column);
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
    // invoke: bool isColumnHidden(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK9QTreeView14isColumnHiddenEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "isColumnHidden", args)
  }

}

  // proto:  bool QTreeView::uniformRowHeights();
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
    // invoke: bool uniformRowHeights()
    C._ZNK9QTreeView17uniformRowHeightsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "uniformRowHeights", args)
  }

}

  // proto:  void QTreeView::setUniformRowHeights(bool uniform);
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
    // invoke: void setUniformRowHeights(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QTreeView20setUniformRowHeightsEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setUniformRowHeights", args)
  }

}

  // proto:  void QTreeView::expandToDepth(int depth);
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
    // invoke: void expandToDepth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QTreeView13expandToDepthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "expandToDepth", args)
  }

}

  // proto:  QModelIndex QTreeView::indexBelow(const QModelIndex & index);
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
    // invoke: QModelIndex indexBelow(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK9QTreeView10indexBelowERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "indexBelow", args)
  }

}

  // proto:  void QTreeView::expandAll();
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
    // invoke: void expandAll()
    C._ZN9QTreeView9expandAllEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "expandAll", args)
  }

}

  // proto:  QModelIndex QTreeView::indexAt(const QPoint & p);
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
    // invoke: QModelIndex indexAt(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK9QTreeView7indexAtERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "indexAt", args)
  }

}

  // proto:  void QTreeView::setColumnWidth(int column, int width);
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
    // invoke: void setColumnWidth(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN9QTreeView14setColumnWidthEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeView", "setColumnWidth", args)
  }

}

// <= body block end

