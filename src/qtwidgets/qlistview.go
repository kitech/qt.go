package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
  // proto:  bool QListView::isWrapping();
extern bool C_ZNK9QListView10isWrappingEv(void* qthis); // 4
  // proto:  int QListView::batchSize();
extern int32_t C_ZNK9QListView9batchSizeEv(void* qthis); // 4
  // proto:  void QListView::setModelColumn(int column);
extern void C_ZN9QListView14setModelColumnEi(void* qthis, int32_t arg0); // 4
  // proto:  void QListView::doItemsLayout();
extern void C_ZN9QListView13doItemsLayoutEv(void* qthis); // 4
  // proto:  bool QListView::wordWrap();
extern bool C_ZNK9QListView8wordWrapEv(void* qthis); // 4
  // proto:  QListView::ViewMode QListView::viewMode();
extern void C_ZNK9QListView8viewModeEv(void* qthis); // 4
  // proto:  void QListView::setSpacing(int space);
extern void C_ZN9QListView10setSpacingEi(void* qthis, int32_t arg0); // 4
  // proto:  void QListView::QListView(QWidget * parent);
extern void* C_ZN9QListViewC2EP7QWidget(void* arg0); // 3
  // proto:  QListView::LayoutMode QListView::layoutMode();
extern void C_ZNK9QListView10layoutModeEv(void* qthis); // 4
  // proto:  void QListView::setWrapping(bool enable);
extern void C_ZN9QListView11setWrappingEb(void* qthis, bool arg0); // 4
  // proto:  QModelIndex QListView::indexAt(const QPoint & p);
extern void* C_ZNK9QListView7indexAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QListView::setWordWrap(bool on);
extern void C_ZN9QListView11setWordWrapEb(void* qthis, bool arg0); // 4
  // proto:  QRect QListView::visualRect(const QModelIndex & index);
extern void* C_ZNK9QListView10visualRectERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QListView::Movement QListView::movement();
extern void C_ZNK9QListView8movementEv(void* qthis); // 4
  // proto:  void QListView::clearPropertyFlags();
extern void C_ZN9QListView18clearPropertyFlagsEv(void* qthis); // 4
  // proto:  void QListView::reset();
extern void C_ZN9QListView5resetEv(void* qthis); // 4
  // proto:  QSize QListView::gridSize();
extern void* C_ZNK9QListView8gridSizeEv(void* qthis); // 4
  // proto:  int QListView::spacing();
extern int32_t C_ZNK9QListView7spacingEv(void* qthis); // 4
  // proto:  void QListView::setRootIndex(const QModelIndex & index);
extern void C_ZN9QListView12setRootIndexERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QListView::ResizeMode QListView::resizeMode();
extern void C_ZNK9QListView10resizeModeEv(void* qthis); // 4
  // proto:  void QListView::setBatchSize(int batchSize);
extern void C_ZN9QListView12setBatchSizeEi(void* qthis, int32_t arg0); // 4
  // proto:  void QListView::setGridSize(const QSize & size);
extern void C_ZN9QListView11setGridSizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  bool QListView::isRowHidden(int row);
extern bool C_ZNK9QListView11isRowHiddenEi(void* qthis, int32_t arg0); // 4
  // proto:  const QMetaObject * QListView::metaObject();
extern void C_ZNK9QListView10metaObjectEv(void* qthis); // 4
  // proto:  void QListView::~QListView();
extern void C_ZN9QListViewD2Ev(void* qthis); // 4
  // proto:  QListView::Flow QListView::flow();
extern void C_ZNK9QListView4flowEv(void* qthis); // 4
  // proto:  bool QListView::uniformItemSizes();
extern bool C_ZNK9QListView16uniformItemSizesEv(void* qthis); // 4
  // proto:  void QListView::setUniformItemSizes(bool enable);
extern void C_ZN9QListView19setUniformItemSizesEb(void* qthis, bool arg0); // 4
  // proto:  void QListView::setRowHidden(int row, bool hide);
extern void C_ZN9QListView12setRowHiddenEib(void* qthis, int32_t arg0, bool arg1); // 4
  // proto:  bool QListView::isSelectionRectVisible();
extern bool C_ZNK9QListView22isSelectionRectVisibleEv(void* qthis); // 4
  // proto:  void QListView::setSelectionRectVisible(bool show);
extern void C_ZN9QListView23setSelectionRectVisibleEb(void* qthis, bool arg0); // 4
  // proto:  int QListView::modelColumn();
extern int32_t C_ZNK9QListView11modelColumnEv(void* qthis); // 4
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
}

// class sizeof(QListView)=1
type QListView struct {
  /*qbase*/ QAbstractItemView;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _indexesMoved QListView_indexesMoved_signal;
}

// isWrapping()
func (this *QListView) Iswrapping(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QListView10isWrappingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListView", "isWrapping", args)
  }

  return
}

// batchSize()
func (this *QListView) Batchsize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QListView9batchSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListView", "batchSize", args)
  }

  return
}

// setModelColumn(int)
func (this *QListView) Setmodelcolumn(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QListView14setModelColumnEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListView", "setModelColumn", args)
  }

  return
}

// doItemsLayout()
func (this *QListView) Doitemslayout(args ...interface{}) () {
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
    C.C_ZN9QListView13doItemsLayoutEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "doItemsLayout", args)
  }

  return
}

// wordWrap()
func (this *QListView) Wordwrap(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QListView8wordWrapEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListView", "wordWrap", args)
  }

  return
}

// viewMode()
func (this *QListView) Viewmode(args ...interface{}) () {
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
    C.C_ZNK9QListView8viewModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "viewMode", args)
  }

  return
}

// setSpacing(int)
func (this *QListView) Setspacing(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QListView10setSpacingEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListView", "setSpacing", args)
  }

  return
}

// QListView(class QWidget *)
func NewQListView(args ...interface{}) *QListView {
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QListViewC2EP7QWidget(arg0)
    return &QListView{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QListView", "QListView", args)
  }

  return nil // QListView{Qclsinst:qthis}
}

// layoutMode()
func (this *QListView) Layoutmode(args ...interface{}) () {
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
    C.C_ZNK9QListView10layoutModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "layoutMode", args)
  }

  return
}

// setWrapping(_Bool)
func (this *QListView) Setwrapping(args ...interface{}) () {
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
    C.C_ZN9QListView11setWrappingEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListView", "setWrapping", args)
  }

  return
}

// indexAt(const class QPoint &)
func (this *QListView) Indexat(args ...interface{}) (ret interface{}) {
  // indexAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListView7indexAtERK6QPoint
    // invoke: QModelIndex indexAt(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QListView7indexAtERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListView", "indexAt", args)
  }

  return
}

// setWordWrap(_Bool)
func (this *QListView) Setwordwrap(args ...interface{}) () {
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
    C.C_ZN9QListView11setWordWrapEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListView", "setWordWrap", args)
  }

  return
}

// visualRect(const class QModelIndex &)
func (this *QListView) Visualrect(args ...interface{}) (ret interface{}) {
  // visualRect(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListView10visualRectERK11QModelIndex
    // invoke: QRect visualRect(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QListView10visualRectERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListView", "visualRect", args)
  }

  return
}

// movement()
func (this *QListView) Movement(args ...interface{}) () {
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
    C.C_ZNK9QListView8movementEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "movement", args)
  }

  return
}

// clearPropertyFlags()
func (this *QListView) Clearpropertyflags(args ...interface{}) () {
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
    C.C_ZN9QListView18clearPropertyFlagsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "clearPropertyFlags", args)
  }

  return
}

// reset()
func (this *QListView) Reset(args ...interface{}) () {
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
    C.C_ZN9QListView5resetEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "reset", args)
  }

  return
}

// gridSize()
func (this *QListView) Gridsize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QListView8gridSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListView", "gridSize", args)
  }

  return
}

// spacing()
func (this *QListView) Spacing(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QListView7spacingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListView", "spacing", args)
  }

  return
}

// setRootIndex(const class QModelIndex &)
func (this *QListView) Setrootindex(args ...interface{}) () {
  // setRootIndex(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListView12setRootIndexERK11QModelIndex
    // invoke: void setRootIndex(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QListView12setRootIndexERK11QModelIndex(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListView", "setRootIndex", args)
  }

  return
}

// resizeMode()
func (this *QListView) Resizemode(args ...interface{}) () {
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
    C.C_ZNK9QListView10resizeModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "resizeMode", args)
  }

  return
}

// setBatchSize(int)
func (this *QListView) Setbatchsize(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QListView12setBatchSizeEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListView", "setBatchSize", args)
  }

  return
}

// setGridSize(const class QSize &)
func (this *QListView) Setgridsize(args ...interface{}) () {
  // setGridSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListView11setGridSizeERK5QSize
    // invoke: void setGridSize(const class QSize &)
    var arg0 = args[0].(*qtcore.QSize).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QListView11setGridSizeERK5QSize(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListView", "setGridSize", args)
  }

  return
}

// isRowHidden(int)
func (this *QListView) Isrowhidden(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QListView11isRowHiddenEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListView", "isRowHidden", args)
  }

  return
}

// metaObject()
func (this *QListView) Metaobject(args ...interface{}) () {
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
    C.C_ZNK9QListView10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "metaObject", args)
  }

  return
}

// ~QListView()
func (this *QListView) Freeqlistview(args ...interface{}) () {
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
    C.C_ZN9QListViewD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "~QListView", args)
  }

  return
}

// flow()
func (this *QListView) Flow(args ...interface{}) () {
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
    C.C_ZNK9QListView4flowEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QListView", "flow", args)
  }

  return
}

// uniformItemSizes()
func (this *QListView) Uniformitemsizes(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QListView16uniformItemSizesEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListView", "uniformItemSizes", args)
  }

  return
}

// setUniformItemSizes(_Bool)
func (this *QListView) Setuniformitemsizes(args ...interface{}) () {
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
    C.C_ZN9QListView19setUniformItemSizesEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListView", "setUniformItemSizes", args)
  }

  return
}

// setRowHidden(int, _Bool)
func (this *QListView) Setrowhidden(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN9QListView12setRowHiddenEib(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QListView", "setRowHidden", args)
  }

  return
}

// isSelectionRectVisible()
func (this *QListView) Isselectionrectvisible(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QListView22isSelectionRectVisibleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListView", "isSelectionRectVisible", args)
  }

  return
}

// setSelectionRectVisible(_Bool)
func (this *QListView) Setselectionrectvisible(args ...interface{}) () {
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
    C.C_ZN9QListView23setSelectionRectVisibleEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListView", "setSelectionRectVisible", args)
  }

  return
}

// modelColumn()
func (this *QListView) Modelcolumn(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QListView11modelColumnEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListView", "modelColumn", args)
  }

  return
}

// <= body block end

