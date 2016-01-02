package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtWidgets/qlistview.h
// dst-file: /src/widgets/qlistview.go
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
  // proto:  void QListView::QListView(QWidget * parent);
extern void* dector_ZN9QListViewC1EP7QWidget(void* arg0);
extern void _ZN9QListViewC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QListView::setWordWrap(bool on);
extern void _ZN9QListView11setWordWrapEb(void* qthis, bool arg0);
  // proto:  void QListView::doItemsLayout();
extern void _ZN9QListView13doItemsLayoutEv(void* qthis);
  // proto:  int QListView::spacing();
extern void _ZNK9QListView7spacingEv(void* qthis);
  // proto:  void QListView::setGridSize(const QSize & size);
extern void _ZN9QListView11setGridSizeERK5QSize(void* qthis, void* arg0);
  // proto:  QModelIndex QListView::indexAt(const QPoint & p);
extern void _ZNK9QListView7indexAtERK6QPoint(void* qthis, void* arg0);
  // proto:  void QListView::setWrapping(bool enable);
extern void _ZN9QListView11setWrappingEb(void* qthis, bool arg0);
  // proto:  void QListView::setSelectionRectVisible(bool show);
extern void _ZN9QListView23setSelectionRectVisibleEb(void* qthis, bool arg0);
  // proto:  void QListView::setBatchSize(int batchSize);
extern void _ZN9QListView12setBatchSizeEi(void* qthis, int arg0);
  // proto:  bool QListView::uniformItemSizes();
extern void _ZNK9QListView16uniformItemSizesEv(void* qthis);
  // proto:  void QListView::setRootIndex(const QModelIndex & index);
extern void _ZN9QListView12setRootIndexERK11QModelIndex(void* qthis, void* arg0);
  // proto:  bool QListView::isWrapping();
extern void _ZNK9QListView10isWrappingEv(void* qthis);
  // proto:  void QListView::reset();
extern void _ZN9QListView5resetEv(void* qthis);
  // proto:  QSize QListView::gridSize();
extern void _ZNK9QListView8gridSizeEv(void* qthis);
  // proto:  void QListView::setModelColumn(int column);
extern void _ZN9QListView14setModelColumnEi(void* qthis, int arg0);
  // proto:  void QListView::QListView(const QListView & );
extern void* dector_ZN9QListViewC1ERKS_(void* arg0);
extern void _ZN9QListViewC1ERKS_(void* qthis, void* arg0);
  // proto:  void QListView::setSpacing(int space);
extern void _ZN9QListView10setSpacingEi(void* qthis, int arg0);
  // proto:  QRect QListView::visualRect(const QModelIndex & index);
extern void _ZNK9QListView10visualRectERK11QModelIndex(void* qthis, void* arg0);
  // proto:  bool QListView::isRowHidden(int row);
extern void _ZNK9QListView11isRowHiddenEi(void* qthis, int arg0);
  // proto:  void QListView::~QListView();
extern void _ZN9QListViewD0Ev(void* qthis);
  // proto:  const QMetaObject * QListView::metaObject();
extern void _ZNK9QListView10metaObjectEv(void* qthis);
  // proto:  int QListView::batchSize();
extern void _ZNK9QListView9batchSizeEv(void* qthis);
  // proto:  bool QListView::isSelectionRectVisible();
extern void _ZNK9QListView22isSelectionRectVisibleEv(void* qthis);
  // proto:  bool QListView::wordWrap();
extern void _ZNK9QListView8wordWrapEv(void* qthis);
  // proto:  void QListView::setRowHidden(int row, bool hide);
extern void _ZN9QListView12setRowHiddenEib(void* qthis, int arg0, bool arg1);
  // proto:  void QListView::clearPropertyFlags();
extern void _ZN9QListView18clearPropertyFlagsEv(void* qthis);
  // proto:  int QListView::modelColumn();
extern void _ZNK9QListView11modelColumnEv(void* qthis);
  // proto:  void QListView::setUniformItemSizes(bool enable);
extern void _ZN9QListView19setUniformItemSizesEb(void* qthis, bool arg0);
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

// class sizeof(QListView)=1
type QListView struct {
  /*qbase*/ QAbstractItemView;
  qclsinst uint64 /* *mut c_void*/;
//  _indexesMoved QListView_indexesMoved_signal;
}

  // proto:  void QListView::QListView(QWidget * parent);
func NewQListView(args ...interface{}) QListView {
  return QListView{}
}

  // proto:  void QListView::setWordWrap(bool on);
func (this *QListView) setWordWrap(args ...interface{}) () {
  // setWordWrap(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListView11setWordWrapEb
    // invoke: void setWordWrap(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN9QListView11setWordWrapEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListView", "setWordWrap", args)
  }

}

  // proto:  void QListView::doItemsLayout();
func (this *QListView) doItemsLayout(args ...interface{}) () {
  // doItemsLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListView13doItemsLayoutEv
    // invoke: void doItemsLayout()
    C._ZN9QListView13doItemsLayoutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "doItemsLayout", args)
  }

}

  // proto:  int QListView::spacing();
func (this *QListView) spacing(args ...interface{}) () {
  // spacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListView7spacingEv
    // invoke: int spacing()
    C._ZNK9QListView7spacingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "spacing", args)
  }

}

  // proto:  void QListView::setGridSize(const QSize & size);
func (this *QListView) setGridSize(args ...interface{}) () {
  // setGridSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListView11setGridSizeERK5QSize
    // invoke: void setGridSize(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QListView11setGridSizeERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListView", "setGridSize", args)
  }

}

  // proto:  QModelIndex QListView::indexAt(const QPoint & p);
func (this *QListView) indexAt(args ...interface{}) () {
  // indexAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListView7indexAtERK6QPoint
    // invoke: QModelIndex indexAt(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK9QListView7indexAtERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListView", "indexAt", args)
  }

}

  // proto:  void QListView::setWrapping(bool enable);
func (this *QListView) setWrapping(args ...interface{}) () {
  // setWrapping(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListView11setWrappingEb
    // invoke: void setWrapping(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN9QListView11setWrappingEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListView", "setWrapping", args)
  }

}

  // proto:  void QListView::setSelectionRectVisible(bool show);
func (this *QListView) setSelectionRectVisible(args ...interface{}) () {
  // setSelectionRectVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListView23setSelectionRectVisibleEb
    // invoke: void setSelectionRectVisible(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN9QListView23setSelectionRectVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListView", "setSelectionRectVisible", args)
  }

}

  // proto:  void QListView::setBatchSize(int batchSize);
func (this *QListView) setBatchSize(args ...interface{}) () {
  // setBatchSize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListView12setBatchSizeEi
    // invoke: void setBatchSize(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QListView12setBatchSizeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListView", "setBatchSize", args)
  }

}

  // proto:  bool QListView::uniformItemSizes();
func (this *QListView) uniformItemSizes(args ...interface{}) () {
  // uniformItemSizes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListView16uniformItemSizesEv
    // invoke: bool uniformItemSizes()
    C._ZNK9QListView16uniformItemSizesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "uniformItemSizes", args)
  }

}

  // proto:  void QListView::setRootIndex(const QModelIndex & index);
func (this *QListView) setRootIndex(args ...interface{}) () {
  // setRootIndex(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListView12setRootIndexERK11QModelIndex
    // invoke: void setRootIndex(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QListView12setRootIndexERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListView", "setRootIndex", args)
  }

}

  // proto:  bool QListView::isWrapping();
func (this *QListView) isWrapping(args ...interface{}) () {
  // isWrapping()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListView10isWrappingEv
    // invoke: bool isWrapping()
    C._ZNK9QListView10isWrappingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "isWrapping", args)
  }

}

  // proto:  void QListView::reset();
func (this *QListView) reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListView5resetEv
    // invoke: void reset()
    C._ZN9QListView5resetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "reset", args)
  }

}

  // proto:  QSize QListView::gridSize();
func (this *QListView) gridSize(args ...interface{}) () {
  // gridSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListView8gridSizeEv
    // invoke: QSize gridSize()
    C._ZNK9QListView8gridSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "gridSize", args)
  }

}

  // proto:  void QListView::setModelColumn(int column);
func (this *QListView) setModelColumn(args ...interface{}) () {
  // setModelColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListView14setModelColumnEi
    // invoke: void setModelColumn(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QListView14setModelColumnEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListView", "setModelColumn", args)
  }

}

  // proto:  void QListView::setSpacing(int space);
func (this *QListView) setSpacing(args ...interface{}) () {
  // setSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListView10setSpacingEi
    // invoke: void setSpacing(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QListView10setSpacingEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListView", "setSpacing", args)
  }

}

  // proto:  QRect QListView::visualRect(const QModelIndex & index);
func (this *QListView) visualRect(args ...interface{}) () {
  // visualRect(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListView10visualRectERK11QModelIndex
    // invoke: QRect visualRect(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK9QListView10visualRectERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListView", "visualRect", args)
  }

}

  // proto:  bool QListView::isRowHidden(int row);
func (this *QListView) isRowHidden(args ...interface{}) () {
  // isRowHidden(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListView11isRowHiddenEi
    // invoke: bool isRowHidden(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK9QListView11isRowHiddenEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListView", "isRowHidden", args)
  }

}

  // proto:  void QListView::~QListView();
func (this *QListView) FreeQListView(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QListView", "~QListView", args)
  }

}

  // proto:  const QMetaObject * QListView::metaObject();
func (this *QListView) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListView10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK9QListView10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "metaObject", args)
  }

}

  // proto:  int QListView::batchSize();
func (this *QListView) batchSize(args ...interface{}) () {
  // batchSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListView9batchSizeEv
    // invoke: int batchSize()
    C._ZNK9QListView9batchSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "batchSize", args)
  }

}

  // proto:  bool QListView::isSelectionRectVisible();
func (this *QListView) isSelectionRectVisible(args ...interface{}) () {
  // isSelectionRectVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListView22isSelectionRectVisibleEv
    // invoke: bool isSelectionRectVisible()
    C._ZNK9QListView22isSelectionRectVisibleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "isSelectionRectVisible", args)
  }

}

  // proto:  bool QListView::wordWrap();
func (this *QListView) wordWrap(args ...interface{}) () {
  // wordWrap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListView8wordWrapEv
    // invoke: bool wordWrap()
    C._ZNK9QListView8wordWrapEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "wordWrap", args)
  }

}

  // proto:  void QListView::setRowHidden(int row, bool hide);
func (this *QListView) setRowHidden(args ...interface{}) () {
  // setRowHidden(int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListView12setRowHiddenEib
    // invoke: void setRowHidden(int, _Bool)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int8_t(args[1].(int8))
    if false {fmt.Println(arg1)}
    C._ZN9QListView12setRowHiddenEib(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QListView", "setRowHidden", args)
  }

}

  // proto:  void QListView::clearPropertyFlags();
func (this *QListView) clearPropertyFlags(args ...interface{}) () {
  // clearPropertyFlags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListView18clearPropertyFlagsEv
    // invoke: void clearPropertyFlags()
    C._ZN9QListView18clearPropertyFlagsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "clearPropertyFlags", args)
  }

}

  // proto:  int QListView::modelColumn();
func (this *QListView) modelColumn(args ...interface{}) () {
  // modelColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListView11modelColumnEv
    // invoke: int modelColumn()
    C._ZNK9QListView11modelColumnEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "modelColumn", args)
  }

}

  // proto:  void QListView::setUniformItemSizes(bool enable);
func (this *QListView) setUniformItemSizes(args ...interface{}) () {
  // setUniformItemSizes(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListView19setUniformItemSizesEb
    // invoke: void setUniformItemSizes(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN9QListView19setUniformItemSizesEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListView", "setUniformItemSizes", args)
  }

}

// <= body block end

