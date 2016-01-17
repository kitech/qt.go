package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  bool QListView::isWrapping();
extern void _ZNK9QListView10isWrappingEv(void* qthis); // 4
  // proto:  int QListView::batchSize();
extern void _ZNK9QListView9batchSizeEv(void* qthis); // 4
  // proto:  void QListView::setModelColumn(int column);
extern void _ZN9QListView14setModelColumnEi(void* qthis, int32_t arg0); // 4
  // proto:  void QListView::doItemsLayout();
extern void _ZN9QListView13doItemsLayoutEv(void* qthis); // 4
  // proto:  bool QListView::wordWrap();
extern void _ZNK9QListView8wordWrapEv(void* qthis); // 4
  // proto:  QListView::ViewMode QListView::viewMode();
extern void _ZNK9QListView8viewModeEv(void* qthis); // 4
  // proto:  void QListView::setSpacing(int space);
extern void _ZN9QListView10setSpacingEi(void* qthis, int32_t arg0); // 4
  // proto:  void QListView::QListView(QWidget * parent);
extern void _ZN9QListViewC2EP7QWidget(void* qthis, void* arg0); // 3
  // proto:  QListView::LayoutMode QListView::layoutMode();
extern void _ZNK9QListView10layoutModeEv(void* qthis); // 4
  // proto:  void QListView::setWrapping(bool enable);
extern void _ZN9QListView11setWrappingEb(void* qthis, bool arg0); // 4
  // proto:  QModelIndex QListView::indexAt(const QPoint & p);
extern void _ZNK9QListView7indexAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QListView::setWordWrap(bool on);
extern void _ZN9QListView11setWordWrapEb(void* qthis, bool arg0); // 4
  // proto:  QRect QListView::visualRect(const QModelIndex & index);
extern void _ZNK9QListView10visualRectERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QListView::Movement QListView::movement();
extern void _ZNK9QListView8movementEv(void* qthis); // 4
  // proto:  void QListView::clearPropertyFlags();
extern void _ZN9QListView18clearPropertyFlagsEv(void* qthis); // 4
  // proto:  void QListView::reset();
extern void _ZN9QListView5resetEv(void* qthis); // 4
  // proto:  QSize QListView::gridSize();
extern void _ZNK9QListView8gridSizeEv(void* qthis); // 4
  // proto:  int QListView::spacing();
extern void _ZNK9QListView7spacingEv(void* qthis); // 4
  // proto:  void QListView::setRootIndex(const QModelIndex & index);
extern void _ZN9QListView12setRootIndexERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QListView::ResizeMode QListView::resizeMode();
extern void _ZNK9QListView10resizeModeEv(void* qthis); // 4
  // proto:  void QListView::setBatchSize(int batchSize);
extern void _ZN9QListView12setBatchSizeEi(void* qthis, int32_t arg0); // 4
  // proto:  void QListView::setGridSize(const QSize & size);
extern void _ZN9QListView11setGridSizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  bool QListView::isRowHidden(int row);
extern void _ZNK9QListView11isRowHiddenEi(void* qthis, int32_t arg0); // 4
  // proto:  const QMetaObject * QListView::metaObject();
extern void _ZNK9QListView10metaObjectEv(void* qthis); // 4
  // proto:  void QListView::~QListView();
extern void _ZN9QListViewD2Ev(void* qthis); // 4
  // proto:  QListView::Flow QListView::flow();
extern void _ZNK9QListView4flowEv(void* qthis); // 4
  // proto:  bool QListView::uniformItemSizes();
extern void _ZNK9QListView16uniformItemSizesEv(void* qthis); // 4
  // proto:  void QListView::setUniformItemSizes(bool enable);
extern void _ZN9QListView19setUniformItemSizesEb(void* qthis, bool arg0); // 4
  // proto:  void QListView::setRowHidden(int row, bool hide);
extern void _ZN9QListView12setRowHiddenEib(void* qthis, int32_t arg0, bool arg1); // 4
  // proto:  bool QListView::isSelectionRectVisible();
extern void _ZNK9QListView22isSelectionRectVisibleEv(void* qthis); // 4
  // proto:  void QListView::setSelectionRectVisible(bool show);
extern void _ZN9QListView23setSelectionRectVisibleEb(void* qthis, bool arg0); // 4
  // proto:  int QListView::modelColumn();
extern void _ZNK9QListView11modelColumnEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
//  _indexesMoved QListView_indexesMoved_signal;
}

// isWrapping()
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

// batchSize()
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

// setModelColumn(int)
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

// doItemsLayout()
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

// wordWrap()
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

// viewMode()
func (this *QListView) viewMode(args ...interface{}) () {
  // viewMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListView8viewModeEv
    // invoke: QListView::ViewMode viewMode()
    C._ZNK9QListView8viewModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "viewMode", args)
  }

}

// setSpacing(int)
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

// QListView(class QWidget *)
func NewQListView(args ...interface{}) QListView {
  // QListView(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListViewC1EP7QWidget
    // invoke: void QListView(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN9QListViewC2EP7QWidget(qthis, arg0)
  default:
    qtrt.ErrorResolve("QListView", "QListView", args)
  }

  return QListView{}
}

// layoutMode()
func (this *QListView) layoutMode(args ...interface{}) () {
  // layoutMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListView10layoutModeEv
    // invoke: QListView::LayoutMode layoutMode()
    C._ZNK9QListView10layoutModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "layoutMode", args)
  }

}

// setWrapping(_Bool)
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QListView11setWrappingEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListView", "setWrapping", args)
  }

}

// indexAt(const class QPoint &)
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

// setWordWrap(_Bool)
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QListView11setWordWrapEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListView", "setWordWrap", args)
  }

}

// visualRect(const class QModelIndex &)
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

// movement()
func (this *QListView) movement(args ...interface{}) () {
  // movement()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListView8movementEv
    // invoke: QListView::Movement movement()
    C._ZNK9QListView8movementEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "movement", args)
  }

}

// clearPropertyFlags()
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

// reset()
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

// gridSize()
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

// spacing()
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

// setRootIndex(const class QModelIndex &)
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

// resizeMode()
func (this *QListView) resizeMode(args ...interface{}) () {
  // resizeMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListView10resizeModeEv
    // invoke: QListView::ResizeMode resizeMode()
    C._ZNK9QListView10resizeModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "resizeMode", args)
  }

}

// setBatchSize(int)
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

// setGridSize(const class QSize &)
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

// isRowHidden(int)
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

// metaObject()
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

// ~QListView()
func (this *QListView) FreeQListView(args ...interface{}) () {
  // ~QListView()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListViewD0Ev
    // invoke: void ~QListView()
    C._ZN9QListViewD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "~QListView", args)
  }

}

// flow()
func (this *QListView) flow(args ...interface{}) () {
  // flow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListView4flowEv
    // invoke: QListView::Flow flow()
    C._ZNK9QListView4flowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "flow", args)
  }

}

// uniformItemSizes()
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

// setUniformItemSizes(_Bool)
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QListView19setUniformItemSizesEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListView", "setUniformItemSizes", args)
  }

}

// setRowHidden(int, _Bool)
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
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C._ZN9QListView12setRowHiddenEib(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QListView", "setRowHidden", args)
  }

}

// isSelectionRectVisible()
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

// setSelectionRectVisible(_Bool)
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QListView23setSelectionRectVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListView", "setSelectionRectVisible", args)
  }

}

// modelColumn()
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

// <= body block end

