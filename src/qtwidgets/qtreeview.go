package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
import "runtime"
import "qtrt"
import "qtcore"
import "qtgui"
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
extern void* C_ZN9QTreeViewC2EP7QWidget(void* arg0); // 3
  // proto:  void QTreeView::setHeaderHidden(bool hide);
extern void C_ZN9QTreeView15setHeaderHiddenEb(void* qthis, bool arg0); // 4
  // proto:  bool QTreeView::expandsOnDoubleClick();
extern bool C_ZNK9QTreeView20expandsOnDoubleClickEv(void* qthis); // 4
  // proto:  bool QTreeView::isExpanded(const QModelIndex & index);
extern bool C_ZNK9QTreeView10isExpandedERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QTreeView::setFirstColumnSpanned(int row, const QModelIndex & parent, bool span);
extern void C_ZN9QTreeView21setFirstColumnSpannedEiRK11QModelIndexb(void* qthis, int32_t arg0, void* arg1, bool arg2); // 4
  // proto:  void QTreeView::resetIndentation();
extern void C_ZN9QTreeView16resetIndentationEv(void* qthis); // 4
  // proto:  QModelIndex QTreeView::indexAbove(const QModelIndex & index);
extern void* C_ZNK9QTreeView10indexAboveERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QTreeView::uniformRowHeights();
extern bool C_ZNK9QTreeView17uniformRowHeightsEv(void* qthis); // 4
  // proto:  void QTreeView::setRootIsDecorated(bool show);
extern void C_ZN9QTreeView18setRootIsDecoratedEb(void* qthis, bool arg0); // 4
  // proto:  void QTreeView::setSelectionModel(QItemSelectionModel * selectionModel);
extern void C_ZN9QTreeView17setSelectionModelEP19QItemSelectionModel(void* qthis, void* arg0); // 4
  // proto:  void QTreeView::setExpandsOnDoubleClick(bool enable);
extern void C_ZN9QTreeView23setExpandsOnDoubleClickEb(void* qthis, bool arg0); // 4
  // proto:  bool QTreeView::isFirstColumnSpanned(int row, const QModelIndex & parent);
extern bool C_ZNK9QTreeView20isFirstColumnSpannedEiRK11QModelIndex(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  const QMetaObject * QTreeView::metaObject();
extern void C_ZNK9QTreeView10metaObjectEv(void* qthis); // 4
  // proto:  QHeaderView * QTreeView::header();
extern void* C_ZNK9QTreeView6headerEv(void* qthis); // 4
  // proto:  void QTreeView::doItemsLayout();
extern void C_ZN9QTreeView13doItemsLayoutEv(void* qthis); // 4
  // proto:  void QTreeView::resizeColumnToContents(int column);
extern void C_ZN9QTreeView22resizeColumnToContentsEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTreeView::columnAt(int x);
extern int32_t C_ZNK9QTreeView8columnAtEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTreeView::setAllColumnsShowFocus(bool enable);
extern void C_ZN9QTreeView22setAllColumnsShowFocusEb(void* qthis, bool arg0); // 4
  // proto:  void QTreeView::showColumn(int column);
extern void C_ZN9QTreeView10showColumnEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTreeView::setExpanded(const QModelIndex & index, bool expand);
extern void C_ZN9QTreeView11setExpandedERK11QModelIndexb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  void QTreeView::setRowHidden(int row, const QModelIndex & parent, bool hide);
extern void C_ZN9QTreeView12setRowHiddenEiRK11QModelIndexb(void* qthis, int32_t arg0, void* arg1, bool arg2); // 4
  // proto:  bool QTreeView::itemsExpandable();
extern bool C_ZNK9QTreeView15itemsExpandableEv(void* qthis); // 4
  // proto:  bool QTreeView::isHeaderHidden();
extern bool C_ZNK9QTreeView14isHeaderHiddenEv(void* qthis); // 4
  // proto:  void QTreeView::setUniformRowHeights(bool uniform);
extern void C_ZN9QTreeView20setUniformRowHeightsEb(void* qthis, bool arg0); // 4
  // proto:  void QTreeView::sortByColumn(int column);
extern void C_ZN9QTreeView12sortByColumnEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTreeView::indentation();
extern int32_t C_ZNK9QTreeView11indentationEv(void* qthis); // 4
  // proto:  void QTreeView::selectAll();
extern void C_ZN9QTreeView9selectAllEv(void* qthis); // 4
  // proto:  void QTreeView::setTreePosition(int logicalIndex);
extern void C_ZN9QTreeView15setTreePositionEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTreeView::autoExpandDelay();
extern int32_t C_ZNK9QTreeView15autoExpandDelayEv(void* qthis); // 4
  // proto:  bool QTreeView::allColumnsShowFocus();
extern bool C_ZNK9QTreeView19allColumnsShowFocusEv(void* qthis); // 4
  // proto:  bool QTreeView::isAnimated();
extern bool C_ZNK9QTreeView10isAnimatedEv(void* qthis); // 4
  // proto:  void QTreeView::setColumnHidden(int column, bool hide);
extern void C_ZN9QTreeView15setColumnHiddenEib(void* qthis, int32_t arg0, bool arg1); // 4
  // proto:  bool QTreeView::isColumnHidden(int column);
extern bool C_ZNK9QTreeView14isColumnHiddenEi(void* qthis, int32_t arg0); // 4
  // proto:  QRect QTreeView::visualRect(const QModelIndex & index);
extern void* C_ZNK9QTreeView10visualRectERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QModelIndex QTreeView::indexAt(const QPoint & p);
extern void* C_ZNK9QTreeView7indexAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QTreeView::collapse(const QModelIndex & index);
extern void C_ZN9QTreeView8collapseERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QTreeView::isSortingEnabled();
extern bool C_ZNK9QTreeView16isSortingEnabledEv(void* qthis); // 4
  // proto:  void QTreeView::setRootIndex(const QModelIndex & index);
extern void C_ZN9QTreeView12setRootIndexERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QTreeView::reset();
extern void C_ZN9QTreeView5resetEv(void* qthis); // 4
  // proto:  int QTreeView::treePosition();
extern int32_t C_ZNK9QTreeView12treePositionEv(void* qthis); // 4
  // proto:  bool QTreeView::rootIsDecorated();
extern bool C_ZNK9QTreeView15rootIsDecoratedEv(void* qthis); // 4
  // proto:  void QTreeView::setSortingEnabled(bool enable);
extern void C_ZN9QTreeView17setSortingEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QTreeView::expandToDepth(int depth);
extern void C_ZN9QTreeView13expandToDepthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTreeView::setAutoExpandDelay(int delay);
extern void C_ZN9QTreeView18setAutoExpandDelayEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTreeView::collapseAll();
extern void C_ZN9QTreeView11collapseAllEv(void* qthis); // 4
  // proto:  bool QTreeView::wordWrap();
extern bool C_ZNK9QTreeView8wordWrapEv(void* qthis); // 4
  // proto:  void QTreeView::hideColumn(int column);
extern void C_ZN9QTreeView10hideColumnEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTreeView::~QTreeView();
extern void C_ZN9QTreeViewD2Ev(void* qthis); // 4
  // proto:  void QTreeView::setWordWrap(bool on);
extern void C_ZN9QTreeView11setWordWrapEb(void* qthis, bool arg0); // 4
  // proto:  bool QTreeView::isRowHidden(int row, const QModelIndex & parent);
extern bool C_ZNK9QTreeView11isRowHiddenEiRK11QModelIndex(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QTreeView::expandAll();
extern void C_ZN9QTreeView9expandAllEv(void* qthis); // 4
  // proto:  int QTreeView::columnViewportPosition(int column);
extern int32_t C_ZNK9QTreeView22columnViewportPositionEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTreeView::setIndentation(int i);
extern void C_ZN9QTreeView14setIndentationEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTreeView::setModel(QAbstractItemModel * model);
extern void C_ZN9QTreeView8setModelEP18QAbstractItemModel(void* qthis, void* arg0); // 4
  // proto:  void QTreeView::setAnimated(bool enable);
extern void C_ZN9QTreeView11setAnimatedEb(void* qthis, bool arg0); // 4
  // proto:  QModelIndex QTreeView::indexBelow(const QModelIndex & index);
extern void* C_ZNK9QTreeView10indexBelowERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QTreeView::expand(const QModelIndex & index);
extern void C_ZN9QTreeView6expandERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QTreeView::setItemsExpandable(bool enable);
extern void C_ZN9QTreeView18setItemsExpandableEb(void* qthis, bool arg0); // 4
  // proto:  int QTreeView::columnWidth(int column);
extern int32_t C_ZNK9QTreeView11columnWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTreeView::setHeader(QHeaderView * header);
extern void C_ZN9QTreeView9setHeaderEP11QHeaderView(void* qthis, void* arg0); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QTreeView)=1
type QTreeView struct {
  /*qbase*/ QAbstractItemView;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _collapsed QTreeView_collapsed_signal;
//  _expanded QTreeView_expanded_signal;
}

// keyboardSearch(const class QString &)
func (this *QTreeView) KeyboardSearch(args ...interface{}) () {
  // keyboardSearch(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView14keyboardSearchERK7QString
    // invoke: void keyboardSearch(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTreeView14keyboardSearchERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "keyboardSearch", args)
  }

  return
}

// setColumnWidth(int, int)
func (this *QTreeView) SetColumnWidth(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN9QTreeView14setColumnWidthEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeView", "setColumnWidth", args)
  }

  return
}

// QTreeView(class QWidget *)
func GcfreeQTreeView(this *QTreeView) {
  qtrt.UniverseFree(this)
}
func NewQTreeView(args ...interface{}) *QTreeView {
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QTreeViewC2EP7QWidget(arg0)
    this := &QTreeView{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQTreeView)
    return this
  default:
    qtrt.ErrorResolve("QTreeView", "QTreeView", args)
  }

  return nil // QTreeView{Qclsinst:qthis}
}

// setHeaderHidden(_Bool)
func (this *QTreeView) SetHeaderHidden(args ...interface{}) () {
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
    C.C_ZN9QTreeView15setHeaderHiddenEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setHeaderHidden", args)
  }

  return
}

// expandsOnDoubleClick()
func (this *QTreeView) ExpandsOnDoubleClick(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTreeView20expandsOnDoubleClickEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTreeView", "expandsOnDoubleClick", args)
  }

  return
}

// isExpanded(const class QModelIndex &)
func (this *QTreeView) IsExpanded(args ...interface{}) (ret interface{}) {
  // isExpanded(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView10isExpandedERK11QModelIndex
    // invoke: bool isExpanded(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QTreeView10isExpandedERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTreeView", "isExpanded", args)
  }

  return
}

// setFirstColumnSpanned(int, const class QModelIndex &, _Bool)
func (this *QTreeView) SetFirstColumnSpanned(args ...interface{}) () {
  // setFirstColumnSpanned(int, const class QModelIndex &, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  vtys[0][2] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView21setFirstColumnSpannedEiRK11QModelIndexb
    // invoke: void setFirstColumnSpanned(int, const class QModelIndex &, _Bool)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.bool(args[2].(bool))
    if false {fmt.Println(arg2)}
    C.C_ZN9QTreeView21setFirstColumnSpannedEiRK11QModelIndexb(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QTreeView", "setFirstColumnSpanned", args)
  }

  return
}

// resetIndentation()
func (this *QTreeView) ResetIndentation(args ...interface{}) () {
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
    C.C_ZN9QTreeView16resetIndentationEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "resetIndentation", args)
  }

  return
}

// indexAbove(const class QModelIndex &)
func (this *QTreeView) IndexAbove(args ...interface{}) (ret interface{}) {
  // indexAbove(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView10indexAboveERK11QModelIndex
    // invoke: QModelIndex indexAbove(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QTreeView10indexAboveERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTreeView", "indexAbove", args)
  }

  return
}

// uniformRowHeights()
func (this *QTreeView) UniformRowHeights(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTreeView17uniformRowHeightsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTreeView", "uniformRowHeights", args)
  }

  return
}

// setRootIsDecorated(_Bool)
func (this *QTreeView) SetRootIsDecorated(args ...interface{}) () {
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
    C.C_ZN9QTreeView18setRootIsDecoratedEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setRootIsDecorated", args)
  }

  return
}

// setSelectionModel(class QItemSelectionModel *)
func (this *QTreeView) SetSelectionModel(args ...interface{}) () {
  // setSelectionModel(class QItemSelectionModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QItemSelectionModel{}) // "QItemSelectionModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView17setSelectionModelEP19QItemSelectionModel
    // invoke: void setSelectionModel(class QItemSelectionModel *)
    var arg0 = args[0].(*qtcore.QItemSelectionModel).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTreeView17setSelectionModelEP19QItemSelectionModel(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setSelectionModel", args)
  }

  return
}

// setExpandsOnDoubleClick(_Bool)
func (this *QTreeView) SetExpandsOnDoubleClick(args ...interface{}) () {
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
    C.C_ZN9QTreeView23setExpandsOnDoubleClickEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setExpandsOnDoubleClick", args)
  }

  return
}

// isFirstColumnSpanned(int, const class QModelIndex &)
func (this *QTreeView) IsFirstColumnSpanned(args ...interface{}) (ret interface{}) {
  // isFirstColumnSpanned(int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView20isFirstColumnSpannedEiRK11QModelIndex
    // invoke: bool isFirstColumnSpanned(int, const class QModelIndex &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK9QTreeView20isFirstColumnSpannedEiRK11QModelIndex(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTreeView", "isFirstColumnSpanned", args)
  }

  return
}

// metaObject()
func (this *QTreeView) MetaObject(args ...interface{}) () {
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
    C.C_ZNK9QTreeView10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "metaObject", args)
  }

  return
}

// header()
func (this *QTreeView) Header(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTreeView6headerEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QHeaderView{}) // "QHeaderView *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTreeView", "header", args)
  }

  return
}

// doItemsLayout()
func (this *QTreeView) DoItemsLayout(args ...interface{}) () {
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
    C.C_ZN9QTreeView13doItemsLayoutEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "doItemsLayout", args)
  }

  return
}

// resizeColumnToContents(int)
func (this *QTreeView) ResizeColumnToContents(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTreeView22resizeColumnToContentsEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "resizeColumnToContents", args)
  }

  return
}

// columnAt(int)
func (this *QTreeView) ColumnAt(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QTreeView8columnAtEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTreeView", "columnAt", args)
  }

  return
}

// setAllColumnsShowFocus(_Bool)
func (this *QTreeView) SetAllColumnsShowFocus(args ...interface{}) () {
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
    C.C_ZN9QTreeView22setAllColumnsShowFocusEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setAllColumnsShowFocus", args)
  }

  return
}

// showColumn(int)
func (this *QTreeView) ShowColumn(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTreeView10showColumnEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "showColumn", args)
  }

  return
}

// setExpanded(const class QModelIndex &, _Bool)
func (this *QTreeView) SetExpanded(args ...interface{}) () {
  // setExpanded(const class QModelIndex &, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView11setExpandedERK11QModelIndexb
    // invoke: void setExpanded(const class QModelIndex &, _Bool)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN9QTreeView11setExpandedERK11QModelIndexb(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeView", "setExpanded", args)
  }

  return
}

// setRowHidden(int, const class QModelIndex &, _Bool)
func (this *QTreeView) SetRowHidden(args ...interface{}) () {
  // setRowHidden(int, const class QModelIndex &, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  vtys[0][2] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView12setRowHiddenEiRK11QModelIndexb
    // invoke: void setRowHidden(int, const class QModelIndex &, _Bool)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.bool(args[2].(bool))
    if false {fmt.Println(arg2)}
    C.C_ZN9QTreeView12setRowHiddenEiRK11QModelIndexb(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QTreeView", "setRowHidden", args)
  }

  return
}

// itemsExpandable()
func (this *QTreeView) ItemsExpandable(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTreeView15itemsExpandableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTreeView", "itemsExpandable", args)
  }

  return
}

// isHeaderHidden()
func (this *QTreeView) IsHeaderHidden(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTreeView14isHeaderHiddenEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTreeView", "isHeaderHidden", args)
  }

  return
}

// setUniformRowHeights(_Bool)
func (this *QTreeView) SetUniformRowHeights(args ...interface{}) () {
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
    C.C_ZN9QTreeView20setUniformRowHeightsEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setUniformRowHeights", args)
  }

  return
}

// sortByColumn(int)
func (this *QTreeView) SortByColumn(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTreeView12sortByColumnEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "sortByColumn", args)
  }

  return
}

// indentation()
func (this *QTreeView) Indentation(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTreeView11indentationEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTreeView", "indentation", args)
  }

  return
}

// selectAll()
func (this *QTreeView) SelectAll(args ...interface{}) () {
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
    C.C_ZN9QTreeView9selectAllEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "selectAll", args)
  }

  return
}

// setTreePosition(int)
func (this *QTreeView) SetTreePosition(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTreeView15setTreePositionEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setTreePosition", args)
  }

  return
}

// autoExpandDelay()
func (this *QTreeView) AutoExpandDelay(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTreeView15autoExpandDelayEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTreeView", "autoExpandDelay", args)
  }

  return
}

// allColumnsShowFocus()
func (this *QTreeView) AllColumnsShowFocus(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTreeView19allColumnsShowFocusEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTreeView", "allColumnsShowFocus", args)
  }

  return
}

// isAnimated()
func (this *QTreeView) IsAnimated(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTreeView10isAnimatedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTreeView", "isAnimated", args)
  }

  return
}

// setColumnHidden(int, _Bool)
func (this *QTreeView) SetColumnHidden(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN9QTreeView15setColumnHiddenEib(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeView", "setColumnHidden", args)
  }

  return
}

// isColumnHidden(int)
func (this *QTreeView) IsColumnHidden(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QTreeView14isColumnHiddenEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTreeView", "isColumnHidden", args)
  }

  return
}

// visualRect(const class QModelIndex &)
func (this *QTreeView) VisualRect(args ...interface{}) (ret interface{}) {
  // visualRect(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView10visualRectERK11QModelIndex
    // invoke: QRect visualRect(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QTreeView10visualRectERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTreeView", "visualRect", args)
  }

  return
}

// indexAt(const class QPoint &)
func (this *QTreeView) IndexAt(args ...interface{}) (ret interface{}) {
  // indexAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView7indexAtERK6QPoint
    // invoke: QModelIndex indexAt(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QTreeView7indexAtERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTreeView", "indexAt", args)
  }

  return
}

// collapse(const class QModelIndex &)
func (this *QTreeView) Collapse(args ...interface{}) () {
  // collapse(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView8collapseERK11QModelIndex
    // invoke: void collapse(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTreeView8collapseERK11QModelIndex(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "collapse", args)
  }

  return
}

// isSortingEnabled()
func (this *QTreeView) IsSortingEnabled(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTreeView16isSortingEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTreeView", "isSortingEnabled", args)
  }

  return
}

// setRootIndex(const class QModelIndex &)
func (this *QTreeView) SetRootIndex(args ...interface{}) () {
  // setRootIndex(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView12setRootIndexERK11QModelIndex
    // invoke: void setRootIndex(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTreeView12setRootIndexERK11QModelIndex(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setRootIndex", args)
  }

  return
}

// reset()
func (this *QTreeView) Reset(args ...interface{}) () {
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
    C.C_ZN9QTreeView5resetEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "reset", args)
  }

  return
}

// treePosition()
func (this *QTreeView) TreePosition(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTreeView12treePositionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTreeView", "treePosition", args)
  }

  return
}

// rootIsDecorated()
func (this *QTreeView) RootIsDecorated(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTreeView15rootIsDecoratedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTreeView", "rootIsDecorated", args)
  }

  return
}

// setSortingEnabled(_Bool)
func (this *QTreeView) SetSortingEnabled(args ...interface{}) () {
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
    C.C_ZN9QTreeView17setSortingEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setSortingEnabled", args)
  }

  return
}

// expandToDepth(int)
func (this *QTreeView) ExpandToDepth(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTreeView13expandToDepthEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "expandToDepth", args)
  }

  return
}

// setAutoExpandDelay(int)
func (this *QTreeView) SetAutoExpandDelay(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTreeView18setAutoExpandDelayEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setAutoExpandDelay", args)
  }

  return
}

// collapseAll()
func (this *QTreeView) CollapseAll(args ...interface{}) () {
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
    C.C_ZN9QTreeView11collapseAllEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "collapseAll", args)
  }

  return
}

// wordWrap()
func (this *QTreeView) WordWrap(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTreeView8wordWrapEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTreeView", "wordWrap", args)
  }

  return
}

// hideColumn(int)
func (this *QTreeView) HideColumn(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTreeView10hideColumnEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "hideColumn", args)
  }

  return
}

// ~QTreeView()
func (this *QTreeView) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN9QTreeViewD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QTreeView", "~QTreeView", args)
  }

  return
}

// setWordWrap(_Bool)
func (this *QTreeView) SetWordWrap(args ...interface{}) () {
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
    C.C_ZN9QTreeView11setWordWrapEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setWordWrap", args)
  }

  return
}

// isRowHidden(int, const class QModelIndex &)
func (this *QTreeView) IsRowHidden(args ...interface{}) (ret interface{}) {
  // isRowHidden(int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView11isRowHiddenEiRK11QModelIndex
    // invoke: bool isRowHidden(int, const class QModelIndex &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK9QTreeView11isRowHiddenEiRK11QModelIndex(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTreeView", "isRowHidden", args)
  }

  return
}

// expandAll()
func (this *QTreeView) ExpandAll(args ...interface{}) () {
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
    C.C_ZN9QTreeView9expandAllEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTreeView", "expandAll", args)
  }

  return
}

// columnViewportPosition(int)
func (this *QTreeView) ColumnViewportPosition(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QTreeView22columnViewportPositionEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTreeView", "columnViewportPosition", args)
  }

  return
}

// setIndentation(int)
func (this *QTreeView) SetIndentation(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTreeView14setIndentationEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setIndentation", args)
  }

  return
}

// setModel(class QAbstractItemModel *)
func (this *QTreeView) SetModel(args ...interface{}) () {
  // setModel(class QAbstractItemModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QAbstractItemModel{}) // "QAbstractItemModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView8setModelEP18QAbstractItemModel
    // invoke: void setModel(class QAbstractItemModel *)
    var arg0 = args[0].(*qtcore.QAbstractItemModel).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTreeView8setModelEP18QAbstractItemModel(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setModel", args)
  }

  return
}

// setAnimated(_Bool)
func (this *QTreeView) SetAnimated(args ...interface{}) () {
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
    C.C_ZN9QTreeView11setAnimatedEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setAnimated", args)
  }

  return
}

// indexBelow(const class QModelIndex &)
func (this *QTreeView) IndexBelow(args ...interface{}) (ret interface{}) {
  // indexBelow(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTreeView10indexBelowERK11QModelIndex
    // invoke: QModelIndex indexBelow(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QTreeView10indexBelowERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTreeView", "indexBelow", args)
  }

  return
}

// expand(const class QModelIndex &)
func (this *QTreeView) Expand(args ...interface{}) () {
  // expand(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTreeView6expandERK11QModelIndex
    // invoke: void expand(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTreeView6expandERK11QModelIndex(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "expand", args)
  }

  return
}

// setItemsExpandable(_Bool)
func (this *QTreeView) SetItemsExpandable(args ...interface{}) () {
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
    C.C_ZN9QTreeView18setItemsExpandableEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setItemsExpandable", args)
  }

  return
}

// columnWidth(int)
func (this *QTreeView) ColumnWidth(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QTreeView11columnWidthEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTreeView", "columnWidth", args)
  }

  return
}

// setHeader(class QHeaderView *)
func (this *QTreeView) SetHeader(args ...interface{}) () {
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
    var arg0 = args[0].(*QHeaderView).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTreeView9setHeaderEP11QHeaderView(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeView", "setHeader", args)
  }

  return
}

// <= body block end

