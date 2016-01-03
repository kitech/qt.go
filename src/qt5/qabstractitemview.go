package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
// src-file: /QtWidgets/qabstractitemview.h
// dst-file: /src/widgets/qabstractitemview.go
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
  // proto:  QWidget * QAbstractItemView::indexWidget(const QModelIndex & index);
extern void _ZNK17QAbstractItemView11indexWidgetERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QAbstractItemView::scrollToBottom();
extern void _ZN17QAbstractItemView14scrollToBottomEv(void* qthis);
  // proto:  void QAbstractItemView::setDropIndicatorShown(bool enable);
extern void _ZN17QAbstractItemView21setDropIndicatorShownEb(void* qthis, bool arg0);
  // proto:  QSize QAbstractItemView::sizeHintForIndex(const QModelIndex & index);
extern void _ZNK17QAbstractItemView16sizeHintForIndexERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QAbstractItemView::setItemDelegateForColumn(int column, QAbstractItemDelegate * delegate);
extern void _ZN17QAbstractItemView24setItemDelegateForColumnEiP21QAbstractItemDelegate(void* qthis, int32_t arg0, void* arg1);
  // proto:  bool QAbstractItemView::dragEnabled();
extern void _ZNK17QAbstractItemView11dragEnabledEv(void* qthis);
  // proto:  QAbstractItemDelegate * QAbstractItemView::itemDelegate(const QModelIndex & index);
extern void _ZNK17QAbstractItemView12itemDelegateERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QAbstractItemView::keyboardSearch(const QString & search);
extern void _ZN17QAbstractItemView14keyboardSearchERK7QString(void* qthis, void* arg0);
  // proto:  QModelIndex QAbstractItemView::rootIndex();
extern void _ZNK17QAbstractItemView9rootIndexEv(void* qthis);
  // proto:  QAbstractItemDelegate * QAbstractItemView::itemDelegateForColumn(int column);
extern void _ZNK17QAbstractItemView21itemDelegateForColumnEi(void* qthis, int32_t arg0);
  // proto:  QItemSelectionModel * QAbstractItemView::selectionModel();
extern void _ZNK17QAbstractItemView14selectionModelEv(void* qthis);
  // proto:  void QAbstractItemView::reset();
extern void _ZN17QAbstractItemView5resetEv(void* qthis);
  // proto:  void QAbstractItemView::setItemDelegateForRow(int row, QAbstractItemDelegate * delegate);
extern void _ZN17QAbstractItemView21setItemDelegateForRowEiP21QAbstractItemDelegate(void* qthis, int32_t arg0, void* arg1);
  // proto:  void QAbstractItemView::setRootIndex(const QModelIndex & index);
extern void _ZN17QAbstractItemView12setRootIndexERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QAbstractItemView::setAutoScrollMargin(int margin);
extern void _ZN17QAbstractItemView19setAutoScrollMarginEi(void* qthis, int32_t arg0);
  // proto:  void QAbstractItemView::doItemsLayout();
extern void _ZN17QAbstractItemView13doItemsLayoutEv(void* qthis);
  // proto:  void QAbstractItemView::~QAbstractItemView();
extern void _ZN17QAbstractItemViewD0Ev(void* qthis);
  // proto:  QAbstractItemModel * QAbstractItemView::model();
extern void _ZNK17QAbstractItemView5modelEv(void* qthis);
  // proto:  QSize QAbstractItemView::iconSize();
extern void _ZNK17QAbstractItemView8iconSizeEv(void* qthis);
  // proto:  void QAbstractItemView::setItemDelegate(QAbstractItemDelegate * delegate);
extern void _ZN17QAbstractItemView15setItemDelegateEP21QAbstractItemDelegate(void* qthis, void* arg0);
  // proto:  void QAbstractItemView::setDragEnabled(bool enable);
extern void _ZN17QAbstractItemView14setDragEnabledEb(void* qthis, bool arg0);
  // proto:  QModelIndex QAbstractItemView::currentIndex();
extern void _ZNK17QAbstractItemView12currentIndexEv(void* qthis);
  // proto:  int QAbstractItemView::sizeHintForRow(int row);
extern void _ZNK17QAbstractItemView14sizeHintForRowEi(void* qthis, int32_t arg0);
  // proto:  void QAbstractItemView::QAbstractItemView(QWidget * parent);
extern void* dector_ZN17QAbstractItemViewC1EP7QWidget(void* arg0);
extern void _ZN17QAbstractItemViewC1EP7QWidget(void* qthis, void* arg0);
  // proto:  bool QAbstractItemView::showDropIndicator();
extern void _ZNK17QAbstractItemView17showDropIndicatorEv(void* qthis);
  // proto:  bool QAbstractItemView::hasAutoScroll();
extern void _ZNK17QAbstractItemView13hasAutoScrollEv(void* qthis);
  // proto:  void QAbstractItemView::QAbstractItemView(const QAbstractItemView & );
extern void* dector_ZN17QAbstractItemViewC1ERKS_(void* arg0);
extern void _ZN17QAbstractItemViewC1ERKS_(void* qthis, void* arg0);
  // proto:  void QAbstractItemView::selectAll();
extern void _ZN17QAbstractItemView9selectAllEv(void* qthis);
  // proto:  QAbstractItemDelegate * QAbstractItemView::itemDelegate();
extern void _ZNK17QAbstractItemView12itemDelegateEv(void* qthis);
  // proto:  void QAbstractItemView::edit(const QModelIndex & index);
extern void _ZN17QAbstractItemView4editERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QAbstractItemView::setAlternatingRowColors(bool enable);
extern void _ZN17QAbstractItemView23setAlternatingRowColorsEb(void* qthis, bool arg0);
  // proto:  int QAbstractItemView::sizeHintForColumn(int column);
extern void _ZNK17QAbstractItemView17sizeHintForColumnEi(void* qthis, int32_t arg0);
  // proto:  void QAbstractItemView::setIconSize(const QSize & size);
extern void _ZN17QAbstractItemView11setIconSizeERK5QSize(void* qthis, void* arg0);
  // proto:  void QAbstractItemView::closePersistentEditor(const QModelIndex & index);
extern void _ZN17QAbstractItemView21closePersistentEditorERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QAbstractItemView::setDragDropOverwriteMode(bool overwrite);
extern void _ZN17QAbstractItemView24setDragDropOverwriteModeEb(void* qthis, bool arg0);
  // proto:  void QAbstractItemView::clearSelection();
extern void _ZN17QAbstractItemView14clearSelectionEv(void* qthis);
  // proto:  void QAbstractItemView::scrollToTop();
extern void _ZN17QAbstractItemView11scrollToTopEv(void* qthis);
  // proto:  void QAbstractItemView::setSelectionModel(QItemSelectionModel * selectionModel);
extern void _ZN17QAbstractItemView17setSelectionModelEP19QItemSelectionModel(void* qthis, void* arg0);
  // proto:  void QAbstractItemView::setCurrentIndex(const QModelIndex & index);
extern void _ZN17QAbstractItemView15setCurrentIndexERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QAbstractItemView::setTabKeyNavigation(bool enable);
extern void _ZN17QAbstractItemView19setTabKeyNavigationEb(void* qthis, bool arg0);
  // proto:  void QAbstractItemView::setIndexWidget(const QModelIndex & index, QWidget * widget);
extern void _ZN17QAbstractItemView14setIndexWidgetERK11QModelIndexP7QWidget(void* qthis, void* arg0, void* arg1);
  // proto:  void QAbstractItemView::setAutoScroll(bool enable);
extern void _ZN17QAbstractItemView13setAutoScrollEb(void* qthis, bool arg0);
  // proto:  void QAbstractItemView::setModel(QAbstractItemModel * model);
extern void _ZN17QAbstractItemView8setModelEP18QAbstractItemModel(void* qthis, void* arg0);
  // proto:  void QAbstractItemView::update(const QModelIndex & index);
extern void _ZN17QAbstractItemView6updateERK11QModelIndex(void* qthis, void* arg0);
  // proto:  const QMetaObject * QAbstractItemView::metaObject();
extern void _ZNK17QAbstractItemView10metaObjectEv(void* qthis);
  // proto:  void QAbstractItemView::openPersistentEditor(const QModelIndex & index);
extern void _ZN17QAbstractItemView20openPersistentEditorERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QAbstractItemDelegate * QAbstractItemView::itemDelegateForRow(int row);
extern void _ZNK17QAbstractItemView18itemDelegateForRowEi(void* qthis, int32_t arg0);
  // proto:  bool QAbstractItemView::dragDropOverwriteMode();
extern void _ZNK17QAbstractItemView21dragDropOverwriteModeEv(void* qthis);
  // proto:  bool QAbstractItemView::tabKeyNavigation();
extern void _ZNK17QAbstractItemView16tabKeyNavigationEv(void* qthis);
  // proto:  int QAbstractItemView::autoScrollMargin();
extern void _ZNK17QAbstractItemView16autoScrollMarginEv(void* qthis);
  // proto:  bool QAbstractItemView::alternatingRowColors();
extern void _ZNK17QAbstractItemView20alternatingRowColorsEv(void* qthis);
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

// class sizeof(QAbstractItemView)=1
type QAbstractItemView struct {
  /*qbase*/ QAbstractScrollArea;
  qclsinst unsafe.Pointer /* *C.void */;
//  _iconSizeChanged QAbstractItemView_iconSizeChanged_signal;
//  _clicked QAbstractItemView_clicked_signal;
//  _viewportEntered QAbstractItemView_viewportEntered_signal;
//  _activated QAbstractItemView_activated_signal;
//  _pressed QAbstractItemView_pressed_signal;
//  _entered QAbstractItemView_entered_signal;
//  _doubleClicked QAbstractItemView_doubleClicked_signal;
}

  // proto:  QWidget * QAbstractItemView::indexWidget(const QModelIndex & index);
func (this *QAbstractItemView) indexWidget(args ...interface{}) () {
  // indexWidget(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView11indexWidgetERK11QModelIndex
    // invoke: QWidget * indexWidget(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK17QAbstractItemView11indexWidgetERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "indexWidget", args)
  }

}

  // proto:  void QAbstractItemView::scrollToBottom();
func (this *QAbstractItemView) scrollToBottom(args ...interface{}) () {
  // scrollToBottom()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView14scrollToBottomEv
    // invoke: void scrollToBottom()
    C._ZN17QAbstractItemView14scrollToBottomEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "scrollToBottom", args)
  }

}

  // proto:  void QAbstractItemView::setDropIndicatorShown(bool enable);
func (this *QAbstractItemView) setDropIndicatorShown(args ...interface{}) () {
  // setDropIndicatorShown(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView21setDropIndicatorShownEb
    // invoke: void setDropIndicatorShown(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN17QAbstractItemView21setDropIndicatorShownEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setDropIndicatorShown", args)
  }

}

  // proto:  QSize QAbstractItemView::sizeHintForIndex(const QModelIndex & index);
func (this *QAbstractItemView) sizeHintForIndex(args ...interface{}) () {
  // sizeHintForIndex(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView16sizeHintForIndexERK11QModelIndex
    // invoke: QSize sizeHintForIndex(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK17QAbstractItemView16sizeHintForIndexERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "sizeHintForIndex", args)
  }

}

  // proto:  void QAbstractItemView::setItemDelegateForColumn(int column, QAbstractItemDelegate * delegate);
func (this *QAbstractItemView) setItemDelegateForColumn(args ...interface{}) () {
  // setItemDelegateForColumn(int, class QAbstractItemDelegate *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QAbstractItemDelegate{}) // "QAbstractItemDelegate *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView24setItemDelegateForColumnEiP21QAbstractItemDelegate
    // invoke: void setItemDelegateForColumn(int, class QAbstractItemDelegate *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QAbstractItemDelegate).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN17QAbstractItemView24setItemDelegateForColumnEiP21QAbstractItemDelegate(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setItemDelegateForColumn", args)
  }

}

  // proto:  bool QAbstractItemView::dragEnabled();
func (this *QAbstractItemView) dragEnabled(args ...interface{}) () {
  // dragEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView11dragEnabledEv
    // invoke: bool dragEnabled()
    C._ZNK17QAbstractItemView11dragEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "dragEnabled", args)
  }

}

  // proto:  QAbstractItemDelegate * QAbstractItemView::itemDelegate(const QModelIndex & index);
func (this *QAbstractItemView) itemDelegate(args ...interface{}) () {
  // itemDelegate(const class QModelIndex &)
  // itemDelegate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView12itemDelegateERK11QModelIndex
    // invoke: QAbstractItemDelegate * itemDelegate(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK17QAbstractItemView12itemDelegateERK11QModelIndex(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK17QAbstractItemView12itemDelegateEv
    // invoke: QAbstractItemDelegate * itemDelegate()
    C._ZNK17QAbstractItemView12itemDelegateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "itemDelegate", args)
  }

}

  // proto:  void QAbstractItemView::keyboardSearch(const QString & search);
func (this *QAbstractItemView) keyboardSearch(args ...interface{}) () {
  // keyboardSearch(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView14keyboardSearchERK7QString
    // invoke: void keyboardSearch(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN17QAbstractItemView14keyboardSearchERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "keyboardSearch", args)
  }

}

  // proto:  QModelIndex QAbstractItemView::rootIndex();
func (this *QAbstractItemView) rootIndex(args ...interface{}) () {
  // rootIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView9rootIndexEv
    // invoke: QModelIndex rootIndex()
    C._ZNK17QAbstractItemView9rootIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "rootIndex", args)
  }

}

  // proto:  QAbstractItemDelegate * QAbstractItemView::itemDelegateForColumn(int column);
func (this *QAbstractItemView) itemDelegateForColumn(args ...interface{}) () {
  // itemDelegateForColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView21itemDelegateForColumnEi
    // invoke: QAbstractItemDelegate * itemDelegateForColumn(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK17QAbstractItemView21itemDelegateForColumnEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "itemDelegateForColumn", args)
  }

}

  // proto:  QItemSelectionModel * QAbstractItemView::selectionModel();
func (this *QAbstractItemView) selectionModel(args ...interface{}) () {
  // selectionModel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView14selectionModelEv
    // invoke: QItemSelectionModel * selectionModel()
    C._ZNK17QAbstractItemView14selectionModelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "selectionModel", args)
  }

}

  // proto:  void QAbstractItemView::reset();
func (this *QAbstractItemView) reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView5resetEv
    // invoke: void reset()
    C._ZN17QAbstractItemView5resetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "reset", args)
  }

}

  // proto:  void QAbstractItemView::setItemDelegateForRow(int row, QAbstractItemDelegate * delegate);
func (this *QAbstractItemView) setItemDelegateForRow(args ...interface{}) () {
  // setItemDelegateForRow(int, class QAbstractItemDelegate *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QAbstractItemDelegate{}) // "QAbstractItemDelegate *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView21setItemDelegateForRowEiP21QAbstractItemDelegate
    // invoke: void setItemDelegateForRow(int, class QAbstractItemDelegate *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QAbstractItemDelegate).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN17QAbstractItemView21setItemDelegateForRowEiP21QAbstractItemDelegate(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setItemDelegateForRow", args)
  }

}

  // proto:  void QAbstractItemView::setRootIndex(const QModelIndex & index);
func (this *QAbstractItemView) setRootIndex(args ...interface{}) () {
  // setRootIndex(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView12setRootIndexERK11QModelIndex
    // invoke: void setRootIndex(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN17QAbstractItemView12setRootIndexERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setRootIndex", args)
  }

}

  // proto:  void QAbstractItemView::setAutoScrollMargin(int margin);
func (this *QAbstractItemView) setAutoScrollMargin(args ...interface{}) () {
  // setAutoScrollMargin(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView19setAutoScrollMarginEi
    // invoke: void setAutoScrollMargin(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN17QAbstractItemView19setAutoScrollMarginEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setAutoScrollMargin", args)
  }

}

  // proto:  void QAbstractItemView::doItemsLayout();
func (this *QAbstractItemView) doItemsLayout(args ...interface{}) () {
  // doItemsLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView13doItemsLayoutEv
    // invoke: void doItemsLayout()
    C._ZN17QAbstractItemView13doItemsLayoutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "doItemsLayout", args)
  }

}

  // proto:  void QAbstractItemView::~QAbstractItemView();
func (this *QAbstractItemView) FreeQAbstractItemView(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractItemView", "~QAbstractItemView", args)
  }

}

  // proto:  QAbstractItemModel * QAbstractItemView::model();
func (this *QAbstractItemView) model(args ...interface{}) () {
  // model()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView5modelEv
    // invoke: QAbstractItemModel * model()
    C._ZNK17QAbstractItemView5modelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "model", args)
  }

}

  // proto:  QSize QAbstractItemView::iconSize();
func (this *QAbstractItemView) iconSize(args ...interface{}) () {
  // iconSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView8iconSizeEv
    // invoke: QSize iconSize()
    C._ZNK17QAbstractItemView8iconSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "iconSize", args)
  }

}

  // proto:  void QAbstractItemView::setItemDelegate(QAbstractItemDelegate * delegate);
func (this *QAbstractItemView) setItemDelegate(args ...interface{}) () {
  // setItemDelegate(class QAbstractItemDelegate *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemDelegate{}) // "QAbstractItemDelegate *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView15setItemDelegateEP21QAbstractItemDelegate
    // invoke: void setItemDelegate(class QAbstractItemDelegate *)
    var arg0 = args[0].(QAbstractItemDelegate).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN17QAbstractItemView15setItemDelegateEP21QAbstractItemDelegate(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setItemDelegate", args)
  }

}

  // proto:  void QAbstractItemView::setDragEnabled(bool enable);
func (this *QAbstractItemView) setDragEnabled(args ...interface{}) () {
  // setDragEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView14setDragEnabledEb
    // invoke: void setDragEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN17QAbstractItemView14setDragEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setDragEnabled", args)
  }

}

  // proto:  QModelIndex QAbstractItemView::currentIndex();
func (this *QAbstractItemView) currentIndex(args ...interface{}) () {
  // currentIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView12currentIndexEv
    // invoke: QModelIndex currentIndex()
    C._ZNK17QAbstractItemView12currentIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "currentIndex", args)
  }

}

  // proto:  int QAbstractItemView::sizeHintForRow(int row);
func (this *QAbstractItemView) sizeHintForRow(args ...interface{}) () {
  // sizeHintForRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView14sizeHintForRowEi
    // invoke: int sizeHintForRow(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK17QAbstractItemView14sizeHintForRowEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "sizeHintForRow", args)
  }

}

  // proto:  void QAbstractItemView::QAbstractItemView(QWidget * parent);
func NewQAbstractItemView(args ...interface{}) QAbstractItemView {
  return QAbstractItemView{}
}

  // proto:  bool QAbstractItemView::showDropIndicator();
func (this *QAbstractItemView) showDropIndicator(args ...interface{}) () {
  // showDropIndicator()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView17showDropIndicatorEv
    // invoke: bool showDropIndicator()
    C._ZNK17QAbstractItemView17showDropIndicatorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "showDropIndicator", args)
  }

}

  // proto:  bool QAbstractItemView::hasAutoScroll();
func (this *QAbstractItemView) hasAutoScroll(args ...interface{}) () {
  // hasAutoScroll()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView13hasAutoScrollEv
    // invoke: bool hasAutoScroll()
    C._ZNK17QAbstractItemView13hasAutoScrollEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "hasAutoScroll", args)
  }

}

  // proto:  void QAbstractItemView::selectAll();
func (this *QAbstractItemView) selectAll(args ...interface{}) () {
  // selectAll()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView9selectAllEv
    // invoke: void selectAll()
    C._ZN17QAbstractItemView9selectAllEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "selectAll", args)
  }

}

  // proto:  void QAbstractItemView::edit(const QModelIndex & index);
func (this *QAbstractItemView) edit(args ...interface{}) () {
  // edit(const class QModelIndex &, enum QAbstractItemView::EditTrigger, class QEvent *)
  // edit(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = qtrt.Int32Ty(false) // "QAbstractItemView::EditTrigger"
  vtys[0][2] = reflect.TypeOf(QEvent{}) // "QEvent *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView4editERK11QModelIndex
    // invoke: void edit(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN17QAbstractItemView4editERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "edit", args)
  }

}

  // proto:  void QAbstractItemView::setAlternatingRowColors(bool enable);
func (this *QAbstractItemView) setAlternatingRowColors(args ...interface{}) () {
  // setAlternatingRowColors(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView23setAlternatingRowColorsEb
    // invoke: void setAlternatingRowColors(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN17QAbstractItemView23setAlternatingRowColorsEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setAlternatingRowColors", args)
  }

}

  // proto:  int QAbstractItemView::sizeHintForColumn(int column);
func (this *QAbstractItemView) sizeHintForColumn(args ...interface{}) () {
  // sizeHintForColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView17sizeHintForColumnEi
    // invoke: int sizeHintForColumn(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK17QAbstractItemView17sizeHintForColumnEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "sizeHintForColumn", args)
  }

}

  // proto:  void QAbstractItemView::setIconSize(const QSize & size);
func (this *QAbstractItemView) setIconSize(args ...interface{}) () {
  // setIconSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView11setIconSizeERK5QSize
    // invoke: void setIconSize(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN17QAbstractItemView11setIconSizeERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setIconSize", args)
  }

}

  // proto:  void QAbstractItemView::closePersistentEditor(const QModelIndex & index);
func (this *QAbstractItemView) closePersistentEditor(args ...interface{}) () {
  // closePersistentEditor(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView21closePersistentEditorERK11QModelIndex
    // invoke: void closePersistentEditor(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN17QAbstractItemView21closePersistentEditorERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "closePersistentEditor", args)
  }

}

  // proto:  void QAbstractItemView::setDragDropOverwriteMode(bool overwrite);
func (this *QAbstractItemView) setDragDropOverwriteMode(args ...interface{}) () {
  // setDragDropOverwriteMode(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView24setDragDropOverwriteModeEb
    // invoke: void setDragDropOverwriteMode(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN17QAbstractItemView24setDragDropOverwriteModeEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setDragDropOverwriteMode", args)
  }

}

  // proto:  void QAbstractItemView::clearSelection();
func (this *QAbstractItemView) clearSelection(args ...interface{}) () {
  // clearSelection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView14clearSelectionEv
    // invoke: void clearSelection()
    C._ZN17QAbstractItemView14clearSelectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "clearSelection", args)
  }

}

  // proto:  void QAbstractItemView::scrollToTop();
func (this *QAbstractItemView) scrollToTop(args ...interface{}) () {
  // scrollToTop()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView11scrollToTopEv
    // invoke: void scrollToTop()
    C._ZN17QAbstractItemView11scrollToTopEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "scrollToTop", args)
  }

}

  // proto:  void QAbstractItemView::setSelectionModel(QItemSelectionModel * selectionModel);
func (this *QAbstractItemView) setSelectionModel(args ...interface{}) () {
  // setSelectionModel(class QItemSelectionModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QItemSelectionModel{}) // "QItemSelectionModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView17setSelectionModelEP19QItemSelectionModel
    // invoke: void setSelectionModel(class QItemSelectionModel *)
    var arg0 = args[0].(QItemSelectionModel).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN17QAbstractItemView17setSelectionModelEP19QItemSelectionModel(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setSelectionModel", args)
  }

}

  // proto:  void QAbstractItemView::setCurrentIndex(const QModelIndex & index);
func (this *QAbstractItemView) setCurrentIndex(args ...interface{}) () {
  // setCurrentIndex(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView15setCurrentIndexERK11QModelIndex
    // invoke: void setCurrentIndex(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN17QAbstractItemView15setCurrentIndexERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setCurrentIndex", args)
  }

}

  // proto:  void QAbstractItemView::setTabKeyNavigation(bool enable);
func (this *QAbstractItemView) setTabKeyNavigation(args ...interface{}) () {
  // setTabKeyNavigation(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView19setTabKeyNavigationEb
    // invoke: void setTabKeyNavigation(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN17QAbstractItemView19setTabKeyNavigationEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setTabKeyNavigation", args)
  }

}

  // proto:  void QAbstractItemView::setIndexWidget(const QModelIndex & index, QWidget * widget);
func (this *QAbstractItemView) setIndexWidget(args ...interface{}) () {
  // setIndexWidget(const class QModelIndex &, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView14setIndexWidgetERK11QModelIndexP7QWidget
    // invoke: void setIndexWidget(const class QModelIndex &, class QWidget *)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN17QAbstractItemView14setIndexWidgetERK11QModelIndexP7QWidget(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setIndexWidget", args)
  }

}

  // proto:  void QAbstractItemView::setAutoScroll(bool enable);
func (this *QAbstractItemView) setAutoScroll(args ...interface{}) () {
  // setAutoScroll(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView13setAutoScrollEb
    // invoke: void setAutoScroll(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN17QAbstractItemView13setAutoScrollEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setAutoScroll", args)
  }

}

  // proto:  void QAbstractItemView::setModel(QAbstractItemModel * model);
func (this *QAbstractItemView) setModel(args ...interface{}) () {
  // setModel(class QAbstractItemModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemModel{}) // "QAbstractItemModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView8setModelEP18QAbstractItemModel
    // invoke: void setModel(class QAbstractItemModel *)
    var arg0 = args[0].(QAbstractItemModel).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN17QAbstractItemView8setModelEP18QAbstractItemModel(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setModel", args)
  }

}

  // proto:  void QAbstractItemView::update(const QModelIndex & index);
func (this *QAbstractItemView) update(args ...interface{}) () {
  // update(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView6updateERK11QModelIndex
    // invoke: void update(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN17QAbstractItemView6updateERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "update", args)
  }

}

  // proto:  const QMetaObject * QAbstractItemView::metaObject();
func (this *QAbstractItemView) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK17QAbstractItemView10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "metaObject", args)
  }

}

  // proto:  void QAbstractItemView::openPersistentEditor(const QModelIndex & index);
func (this *QAbstractItemView) openPersistentEditor(args ...interface{}) () {
  // openPersistentEditor(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView20openPersistentEditorERK11QModelIndex
    // invoke: void openPersistentEditor(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN17QAbstractItemView20openPersistentEditorERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "openPersistentEditor", args)
  }

}

  // proto:  QAbstractItemDelegate * QAbstractItemView::itemDelegateForRow(int row);
func (this *QAbstractItemView) itemDelegateForRow(args ...interface{}) () {
  // itemDelegateForRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView18itemDelegateForRowEi
    // invoke: QAbstractItemDelegate * itemDelegateForRow(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK17QAbstractItemView18itemDelegateForRowEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "itemDelegateForRow", args)
  }

}

  // proto:  bool QAbstractItemView::dragDropOverwriteMode();
func (this *QAbstractItemView) dragDropOverwriteMode(args ...interface{}) () {
  // dragDropOverwriteMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView21dragDropOverwriteModeEv
    // invoke: bool dragDropOverwriteMode()
    C._ZNK17QAbstractItemView21dragDropOverwriteModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "dragDropOverwriteMode", args)
  }

}

  // proto:  bool QAbstractItemView::tabKeyNavigation();
func (this *QAbstractItemView) tabKeyNavigation(args ...interface{}) () {
  // tabKeyNavigation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView16tabKeyNavigationEv
    // invoke: bool tabKeyNavigation()
    C._ZNK17QAbstractItemView16tabKeyNavigationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "tabKeyNavigation", args)
  }

}

  // proto:  int QAbstractItemView::autoScrollMargin();
func (this *QAbstractItemView) autoScrollMargin(args ...interface{}) () {
  // autoScrollMargin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView16autoScrollMarginEv
    // invoke: int autoScrollMargin()
    C._ZNK17QAbstractItemView16autoScrollMarginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "autoScrollMargin", args)
  }

}

  // proto:  bool QAbstractItemView::alternatingRowColors();
func (this *QAbstractItemView) alternatingRowColors(args ...interface{}) () {
  // alternatingRowColors()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView20alternatingRowColorsEv
    // invoke: bool alternatingRowColors()
    C._ZNK17QAbstractItemView20alternatingRowColorsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "alternatingRowColors", args)
  }

}

// <= body block end

