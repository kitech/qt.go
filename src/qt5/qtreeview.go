package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QTreeView::keyboardSearch(const QString & search);
extern void C_ZN9QTreeView14keyboardSearchERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QTreeView::setColumnWidth(int column, int width);
extern void C_ZN9QTreeView14setColumnWidthEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QTreeView::QTreeView(QWidget * parent);
extern void C_ZN9QTreeViewC2EP7QWidget(void* qthis, void* arg0); // 3
  // proto:  void QTreeView::setHeaderHidden(bool hide);
extern void C_ZN9QTreeView15setHeaderHiddenEb(void* qthis, bool arg0); // 4
  // proto:  bool QTreeView::expandsOnDoubleClick();
extern void C_ZNK9QTreeView20expandsOnDoubleClickEv(void* qthis); // 4
  // proto:  bool QTreeView::isExpanded(const QModelIndex & index);
extern void C_ZNK9QTreeView10isExpandedERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QTreeView::setFirstColumnSpanned(int row, const QModelIndex & parent, bool span);
extern void C_ZN9QTreeView21setFirstColumnSpannedEiRK11QModelIndexb(void* qthis, int32_t arg0, void* arg1, bool arg2); // 4
  // proto:  void QTreeView::resetIndentation();
extern void C_ZN9QTreeView16resetIndentationEv(void* qthis); // 4
  // proto:  QModelIndex QTreeView::indexAbove(const QModelIndex & index);
extern void C_ZNK9QTreeView10indexAboveERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QTreeView::uniformRowHeights();
extern void C_ZNK9QTreeView17uniformRowHeightsEv(void* qthis); // 4
  // proto:  void QTreeView::setRootIsDecorated(bool show);
extern void C_ZN9QTreeView18setRootIsDecoratedEb(void* qthis, bool arg0); // 4
  // proto:  void QTreeView::setSelectionModel(QItemSelectionModel * selectionModel);
extern void C_ZN9QTreeView17setSelectionModelEP19QItemSelectionModel(void* qthis, void* arg0); // 4
  // proto:  void QTreeView::setExpandsOnDoubleClick(bool enable);
extern void C_ZN9QTreeView23setExpandsOnDoubleClickEb(void* qthis, bool arg0); // 4
  // proto:  bool QTreeView::isFirstColumnSpanned(int row, const QModelIndex & parent);
extern void C_ZNK9QTreeView20isFirstColumnSpannedEiRK11QModelIndex(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  const QMetaObject * QTreeView::metaObject();
extern void C_ZNK9QTreeView10metaObjectEv(void* qthis); // 4
  // proto:  QHeaderView * QTreeView::header();
extern void C_ZNK9QTreeView6headerEv(void* qthis); // 4
  // proto:  void QTreeView::doItemsLayout();
extern void C_ZN9QTreeView13doItemsLayoutEv(void* qthis); // 4
  // proto:  void QTreeView::resizeColumnToContents(int column);
extern void C_ZN9QTreeView22resizeColumnToContentsEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTreeView::columnAt(int x);
extern void C_ZNK9QTreeView8columnAtEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTreeView::setAllColumnsShowFocus(bool enable);
extern void C_ZN9QTreeView22setAllColumnsShowFocusEb(void* qthis, bool arg0); // 4
  // proto:  void QTreeView::showColumn(int column);
extern void C_ZN9QTreeView10showColumnEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTreeView::setExpanded(const QModelIndex & index, bool expand);
extern void C_ZN9QTreeView11setExpandedERK11QModelIndexb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  void QTreeView::setRowHidden(int row, const QModelIndex & parent, bool hide);
extern void C_ZN9QTreeView12setRowHiddenEiRK11QModelIndexb(void* qthis, int32_t arg0, void* arg1, bool arg2); // 4
  // proto:  bool QTreeView::itemsExpandable();
extern void C_ZNK9QTreeView15itemsExpandableEv(void* qthis); // 4
  // proto:  bool QTreeView::isHeaderHidden();
extern void C_ZNK9QTreeView14isHeaderHiddenEv(void* qthis); // 4
  // proto:  void QTreeView::setUniformRowHeights(bool uniform);
extern void C_ZN9QTreeView20setUniformRowHeightsEb(void* qthis, bool arg0); // 4
  // proto:  void QTreeView::sortByColumn(int column);
extern void C_ZN9QTreeView12sortByColumnEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTreeView::indentation();
extern void C_ZNK9QTreeView11indentationEv(void* qthis); // 4
  // proto:  void QTreeView::selectAll();
extern void C_ZN9QTreeView9selectAllEv(void* qthis); // 4
  // proto:  void QTreeView::setTreePosition(int logicalIndex);
extern void C_ZN9QTreeView15setTreePositionEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTreeView::autoExpandDelay();
extern void C_ZNK9QTreeView15autoExpandDelayEv(void* qthis); // 4
  // proto:  bool QTreeView::allColumnsShowFocus();
extern void C_ZNK9QTreeView19allColumnsShowFocusEv(void* qthis); // 4
  // proto:  bool QTreeView::isAnimated();
extern void C_ZNK9QTreeView10isAnimatedEv(void* qthis); // 4
  // proto:  void QTreeView::setColumnHidden(int column, bool hide);
extern void C_ZN9QTreeView15setColumnHiddenEib(void* qthis, int32_t arg0, bool arg1); // 4
  // proto:  bool QTreeView::isColumnHidden(int column);
extern void C_ZNK9QTreeView14isColumnHiddenEi(void* qthis, int32_t arg0); // 4
  // proto:  QRect QTreeView::visualRect(const QModelIndex & index);
extern void C_ZNK9QTreeView10visualRectERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QModelIndex QTreeView::indexAt(const QPoint & p);
extern void C_ZNK9QTreeView7indexAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QTreeView::collapse(const QModelIndex & index);
extern void C_ZN9QTreeView8collapseERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QTreeView::isSortingEnabled();
extern void C_ZNK9QTreeView16isSortingEnabledEv(void* qthis); // 4
  // proto:  void QTreeView::setRootIndex(const QModelIndex & index);
extern void C_ZN9QTreeView12setRootIndexERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QTreeView::reset();
extern void C_ZN9QTreeView5resetEv(void* qthis); // 4
  // proto:  int QTreeView::treePosition();
extern void C_ZNK9QTreeView12treePositionEv(void* qthis); // 4
  // proto:  bool QTreeView::rootIsDecorated();
extern void C_ZNK9QTreeView15rootIsDecoratedEv(void* qthis); // 4
  // proto:  void QTreeView::setSortingEnabled(bool enable);
extern void C_ZN9QTreeView17setSortingEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QTreeView::expandToDepth(int depth);
extern void C_ZN9QTreeView13expandToDepthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTreeView::setAutoExpandDelay(int delay);
extern void C_ZN9QTreeView18setAutoExpandDelayEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTreeView::collapseAll();
extern void C_ZN9QTreeView11collapseAllEv(void* qthis); // 4
  // proto:  bool QTreeView::wordWrap();
extern void C_ZNK9QTreeView8wordWrapEv(void* qthis); // 4
  // proto:  void QTreeView::hideColumn(int column);
extern void C_ZN9QTreeView10hideColumnEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTreeView::~QTreeView();
extern void C_ZN9QTreeViewD2Ev(void* qthis); // 4
  // proto:  void QTreeView::setWordWrap(bool on);
extern void C_ZN9QTreeView11setWordWrapEb(void* qthis, bool arg0); // 4
  // proto:  bool QTreeView::isRowHidden(int row, const QModelIndex & parent);
extern void C_ZNK9QTreeView11isRowHiddenEiRK11QModelIndex(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QTreeView::expandAll();
extern void C_ZN9QTreeView9expandAllEv(void* qthis); // 4
  // proto:  int QTreeView::columnViewportPosition(int column);
extern void C_ZNK9QTreeView22columnViewportPositionEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTreeView::setIndentation(int i);
extern void C_ZN9QTreeView14setIndentationEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTreeView::setModel(QAbstractItemModel * model);
extern void C_ZN9QTreeView8setModelEP18QAbstractItemModel(void* qthis, void* arg0); // 4
  // proto:  void QTreeView::setAnimated(bool enable);
extern void C_ZN9QTreeView11setAnimatedEb(void* qthis, bool arg0); // 4
  // proto:  QModelIndex QTreeView::indexBelow(const QModelIndex & index);
extern void C_ZNK9QTreeView10indexBelowERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QTreeView::expand(const QModelIndex & index);
extern void C_ZN9QTreeView6expandERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QTreeView::setItemsExpandable(bool enable);
extern void C_ZN9QTreeView18setItemsExpandableEb(void* qthis, bool arg0); // 4
  // proto:  int QTreeView::columnWidth(int column);
extern void C_ZNK9QTreeView11columnWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTreeView::setHeader(QHeaderView * header);
extern void C_ZN9QTreeView9setHeaderEP11QHeaderView(void* qthis, void* arg0); // 4
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

// keyboardSearch(const class QString &)
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
    C.C_ZN9QTreeView14keyboardSearchERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "keyboardSearch", args)
  }

}

// setColumnWidth(int, int)
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
    C.C_ZN9QTreeView14setColumnWidthEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeView", "setColumnWidth", args)
  }

}

// QTreeView(class QWidget *)
func NewQTreeView(args ...interface{}) QTreeView {
  // QTreeView(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTreeViewC1EP7QWidget
    // invoke: void QTreeView(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN9QTreeViewC2EP7QWidget(qthis, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "QTreeView", args)
  }

  return QTreeView{}
}

// setHeaderHidden(_Bool)
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
    C.C_ZN9QTreeView15setHeaderHiddenEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setHeaderHidden", args)
  }

}

// expandsOnDoubleClick()
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
    C.C_ZNK9QTreeView20expandsOnDoubleClickEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "expandsOnDoubleClick", args)
  }

}

// isExpanded(const class QModelIndex &)
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
    C.C_ZNK9QTreeView10isExpandedERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "isExpanded", args)
  }

}

// setFirstColumnSpanned(int, const class QModelIndex &, _Bool)
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
    C.C_ZN9QTreeView21setFirstColumnSpannedEiRK11QModelIndexb(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QTreeView", "setFirstColumnSpanned", args)
  }

}

// resetIndentation()
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
    C.C_ZN9QTreeView16resetIndentationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "resetIndentation", args)
  }

}

// indexAbove(const class QModelIndex &)
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
    C.C_ZNK9QTreeView10indexAboveERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "indexAbove", args)
  }

}

// uniformRowHeights()
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
    C.C_ZNK9QTreeView17uniformRowHeightsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "uniformRowHeights", args)
  }

}

// setRootIsDecorated(_Bool)
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
    C.C_ZN9QTreeView18setRootIsDecoratedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setRootIsDecorated", args)
  }

}

// setSelectionModel(class QItemSelectionModel *)
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
    C.C_ZN9QTreeView17setSelectionModelEP19QItemSelectionModel(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setSelectionModel", args)
  }

}

// setExpandsOnDoubleClick(_Bool)
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
    C.C_ZN9QTreeView23setExpandsOnDoubleClickEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setExpandsOnDoubleClick", args)
  }

}

// isFirstColumnSpanned(int, const class QModelIndex &)
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
    C.C_ZNK9QTreeView20isFirstColumnSpannedEiRK11QModelIndex(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeView", "isFirstColumnSpanned", args)
  }

}

// metaObject()
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
    C.C_ZNK9QTreeView10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "metaObject", args)
  }

}

// header()
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
    C.C_ZNK9QTreeView6headerEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "header", args)
  }

}

// doItemsLayout()
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
    C.C_ZN9QTreeView13doItemsLayoutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "doItemsLayout", args)
  }

}

// resizeColumnToContents(int)
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
    C.C_ZN9QTreeView22resizeColumnToContentsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "resizeColumnToContents", args)
  }

}

// columnAt(int)
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
    C.C_ZNK9QTreeView8columnAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "columnAt", args)
  }

}

// setAllColumnsShowFocus(_Bool)
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
    C.C_ZN9QTreeView22setAllColumnsShowFocusEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setAllColumnsShowFocus", args)
  }

}

// showColumn(int)
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
    C.C_ZN9QTreeView10showColumnEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "showColumn", args)
  }

}

// setExpanded(const class QModelIndex &, _Bool)
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
    C.C_ZN9QTreeView11setExpandedERK11QModelIndexb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeView", "setExpanded", args)
  }

}

// setRowHidden(int, const class QModelIndex &, _Bool)
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
    C.C_ZN9QTreeView12setRowHiddenEiRK11QModelIndexb(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QTreeView", "setRowHidden", args)
  }

}

// itemsExpandable()
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
    C.C_ZNK9QTreeView15itemsExpandableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "itemsExpandable", args)
  }

}

// isHeaderHidden()
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
    C.C_ZNK9QTreeView14isHeaderHiddenEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "isHeaderHidden", args)
  }

}

// setUniformRowHeights(_Bool)
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
    C.C_ZN9QTreeView20setUniformRowHeightsEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setUniformRowHeights", args)
  }

}

// sortByColumn(int)
func (this *QTreeView) sortByColumn(args ...interface{}) () {
  // sortByColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView12sortByColumnEi
    // invoke: void sortByColumn(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTreeView12sortByColumnEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "sortByColumn", args)
  }

}

// indentation()
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
    C.C_ZNK9QTreeView11indentationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "indentation", args)
  }

}

// selectAll()
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
    C.C_ZN9QTreeView9selectAllEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "selectAll", args)
  }

}

// setTreePosition(int)
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
    C.C_ZN9QTreeView15setTreePositionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setTreePosition", args)
  }

}

// autoExpandDelay()
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
    C.C_ZNK9QTreeView15autoExpandDelayEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "autoExpandDelay", args)
  }

}

// allColumnsShowFocus()
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
    C.C_ZNK9QTreeView19allColumnsShowFocusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "allColumnsShowFocus", args)
  }

}

// isAnimated()
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
    C.C_ZNK9QTreeView10isAnimatedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "isAnimated", args)
  }

}

// setColumnHidden(int, _Bool)
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
    C.C_ZN9QTreeView15setColumnHiddenEib(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeView", "setColumnHidden", args)
  }

}

// isColumnHidden(int)
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
    C.C_ZNK9QTreeView14isColumnHiddenEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "isColumnHidden", args)
  }

}

// visualRect(const class QModelIndex &)
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
    C.C_ZNK9QTreeView10visualRectERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "visualRect", args)
  }

}

// indexAt(const class QPoint &)
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
    C.C_ZNK9QTreeView7indexAtERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "indexAt", args)
  }

}

// collapse(const class QModelIndex &)
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
    C.C_ZN9QTreeView8collapseERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "collapse", args)
  }

}

// isSortingEnabled()
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
    C.C_ZNK9QTreeView16isSortingEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "isSortingEnabled", args)
  }

}

// setRootIndex(const class QModelIndex &)
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
    C.C_ZN9QTreeView12setRootIndexERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setRootIndex", args)
  }

}

// reset()
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
    C.C_ZN9QTreeView5resetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "reset", args)
  }

}

// treePosition()
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
    C.C_ZNK9QTreeView12treePositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "treePosition", args)
  }

}

// rootIsDecorated()
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
    C.C_ZNK9QTreeView15rootIsDecoratedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "rootIsDecorated", args)
  }

}

// setSortingEnabled(_Bool)
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
    C.C_ZN9QTreeView17setSortingEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setSortingEnabled", args)
  }

}

// expandToDepth(int)
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
    C.C_ZN9QTreeView13expandToDepthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "expandToDepth", args)
  }

}

// setAutoExpandDelay(int)
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
    C.C_ZN9QTreeView18setAutoExpandDelayEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setAutoExpandDelay", args)
  }

}

// collapseAll()
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
    C.C_ZN9QTreeView11collapseAllEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "collapseAll", args)
  }

}

// wordWrap()
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
    C.C_ZNK9QTreeView8wordWrapEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "wordWrap", args)
  }

}

// hideColumn(int)
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
    C.C_ZN9QTreeView10hideColumnEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "hideColumn", args)
  }

}

// ~QTreeView()
func (this *QTreeView) FreeQTreeView(args ...interface{}) () {
  // ~QTreeView()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTreeViewD0Ev
    // invoke: void ~QTreeView()
    C.C_ZN9QTreeViewD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "~QTreeView", args)
  }

}

// setWordWrap(_Bool)
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
    C.C_ZN9QTreeView11setWordWrapEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setWordWrap", args)
  }

}

// isRowHidden(int, const class QModelIndex &)
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
    C.C_ZNK9QTreeView11isRowHiddenEiRK11QModelIndex(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeView", "isRowHidden", args)
  }

}

// expandAll()
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
    C.C_ZN9QTreeView9expandAllEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "expandAll", args)
  }

}

// columnViewportPosition(int)
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
    C.C_ZNK9QTreeView22columnViewportPositionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "columnViewportPosition", args)
  }

}

// setIndentation(int)
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
    C.C_ZN9QTreeView14setIndentationEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setIndentation", args)
  }

}

// setModel(class QAbstractItemModel *)
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
    C.C_ZN9QTreeView8setModelEP18QAbstractItemModel(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setModel", args)
  }

}

// setAnimated(_Bool)
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
    C.C_ZN9QTreeView11setAnimatedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setAnimated", args)
  }

}

// indexBelow(const class QModelIndex &)
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
    C.C_ZNK9QTreeView10indexBelowERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "indexBelow", args)
  }

}

// expand(const class QModelIndex &)
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
    C.C_ZN9QTreeView6expandERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "expand", args)
  }

}

// setItemsExpandable(_Bool)
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
    C.C_ZN9QTreeView18setItemsExpandableEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setItemsExpandable", args)
  }

}

// columnWidth(int)
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
    C.C_ZNK9QTreeView11columnWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "columnWidth", args)
  }

}

// setHeader(class QHeaderView *)
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
    C.C_ZN9QTreeView9setHeaderEP11QHeaderView(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setHeader", args)
  }

}

// <= body block end

