package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
  // proto:  void QAbstractItemView::keyboardSearch(const QString & search);
extern void C_ZN17QAbstractItemView14keyboardSearchERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QAbstractItemView::setDragEnabled(bool enable);
extern void C_ZN17QAbstractItemView14setDragEnabledEb(void* qthis, bool arg0); // 4
  // proto:  bool QAbstractItemView::dragDropOverwriteMode();
extern bool C_ZNK17QAbstractItemView21dragDropOverwriteModeEv(void* qthis); // 4
  // proto:  void QAbstractItemView::openPersistentEditor(const QModelIndex & index);
extern void C_ZN17QAbstractItemView20openPersistentEditorERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QAbstractItemView::SelectionMode QAbstractItemView::selectionMode();
extern void C_ZNK17QAbstractItemView13selectionModeEv(void* qthis); // 4
  // proto:  int QAbstractItemView::autoScrollMargin();
extern int32_t C_ZNK17QAbstractItemView16autoScrollMarginEv(void* qthis); // 4
  // proto:  QModelIndex QAbstractItemView::rootIndex();
extern void* C_ZNK17QAbstractItemView9rootIndexEv(void* qthis); // 4
  // proto:  void QAbstractItemView::setSelectionModel(QItemSelectionModel * selectionModel);
extern void C_ZN17QAbstractItemView17setSelectionModelEP19QItemSelectionModel(void* qthis, void* arg0); // 4
  // proto:  bool QAbstractItemView::tabKeyNavigation();
extern bool C_ZNK17QAbstractItemView16tabKeyNavigationEv(void* qthis); // 4
  // proto:  const QMetaObject * QAbstractItemView::metaObject();
extern void C_ZNK17QAbstractItemView10metaObjectEv(void* qthis); // 4
  // proto:  void QAbstractItemView::edit(const QModelIndex & index);
extern void C_ZN17QAbstractItemView4editERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QAbstractItemView::setIconSize(const QSize & size);
extern void C_ZN17QAbstractItemView11setIconSizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  QModelIndex QAbstractItemView::currentIndex();
extern void* C_ZNK17QAbstractItemView12currentIndexEv(void* qthis); // 4
  // proto:  void QAbstractItemView::closePersistentEditor(const QModelIndex & index);
extern void C_ZN17QAbstractItemView21closePersistentEditorERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QAbstractItemView::ScrollMode QAbstractItemView::verticalScrollMode();
extern void C_ZNK17QAbstractItemView18verticalScrollModeEv(void* qthis); // 4
  // proto:  void QAbstractItemView::doItemsLayout();
extern void C_ZN17QAbstractItemView13doItemsLayoutEv(void* qthis); // 4
  // proto:  void QAbstractItemView::scrollToBottom();
extern void C_ZN17QAbstractItemView14scrollToBottomEv(void* qthis); // 4
  // proto:  void QAbstractItemView::setItemDelegateForRow(int row, QAbstractItemDelegate * delegate);
extern void C_ZN17QAbstractItemView21setItemDelegateForRowEiP21QAbstractItemDelegate(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QAbstractItemView::setTabKeyNavigation(bool enable);
extern void C_ZN17QAbstractItemView19setTabKeyNavigationEb(void* qthis, bool arg0); // 4
  // proto:  QSize QAbstractItemView::iconSize();
extern void* C_ZNK17QAbstractItemView8iconSizeEv(void* qthis); // 4
  // proto:  void QAbstractItemView::selectAll();
extern void C_ZN17QAbstractItemView9selectAllEv(void* qthis); // 4
  // proto:  void QAbstractItemView::setAutoScrollMargin(int margin);
extern void C_ZN17QAbstractItemView19setAutoScrollMarginEi(void* qthis, int32_t arg0); // 4
  // proto:  QAbstractItemDelegate * QAbstractItemView::itemDelegateForColumn(int column);
extern void C_ZNK17QAbstractItemView21itemDelegateForColumnEi(void* qthis, int32_t arg0); // 4
  // proto:  void QAbstractItemView::setDragDropOverwriteMode(bool overwrite);
extern void C_ZN17QAbstractItemView24setDragDropOverwriteModeEb(void* qthis, bool arg0); // 4
  // proto:  QAbstractItemView::DragDropMode QAbstractItemView::dragDropMode();
extern void C_ZNK17QAbstractItemView12dragDropModeEv(void* qthis); // 4
  // proto:  QAbstractItemView::ScrollMode QAbstractItemView::horizontalScrollMode();
extern void C_ZNK17QAbstractItemView20horizontalScrollModeEv(void* qthis); // 4
  // proto:  bool QAbstractItemView::hasAutoScroll();
extern bool C_ZNK17QAbstractItemView13hasAutoScrollEv(void* qthis); // 4
  // proto:  QSize QAbstractItemView::sizeHintForIndex(const QModelIndex & index);
extern void* C_ZNK17QAbstractItemView16sizeHintForIndexERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QAbstractItemView::setRootIndex(const QModelIndex & index);
extern void C_ZN17QAbstractItemView12setRootIndexERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  EditTriggers QAbstractItemView::editTriggers();
extern void C_ZNK17QAbstractItemView12editTriggersEv(void* qthis); // 4
  // proto:  bool QAbstractItemView::showDropIndicator();
extern bool C_ZNK17QAbstractItemView17showDropIndicatorEv(void* qthis); // 4
  // proto:  Qt::TextElideMode QAbstractItemView::textElideMode();
extern void C_ZNK17QAbstractItemView13textElideModeEv(void* qthis); // 4
  // proto:  void QAbstractItemView::reset();
extern void C_ZN17QAbstractItemView5resetEv(void* qthis); // 4
  // proto:  bool QAbstractItemView::dragEnabled();
extern bool C_ZNK17QAbstractItemView11dragEnabledEv(void* qthis); // 4
  // proto:  void QAbstractItemView::setIndexWidget(const QModelIndex & index, QWidget * widget);
extern void C_ZN17QAbstractItemView14setIndexWidgetERK11QModelIndexP7QWidget(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QAbstractItemView::scrollToTop();
extern void C_ZN17QAbstractItemView11scrollToTopEv(void* qthis); // 4
  // proto:  void QAbstractItemView::setItemDelegate(QAbstractItemDelegate * delegate);
extern void C_ZN17QAbstractItemView15setItemDelegateEP21QAbstractItemDelegate(void* qthis, void* arg0); // 4
  // proto:  QAbstractItemDelegate * QAbstractItemView::itemDelegateForRow(int row);
extern void C_ZNK17QAbstractItemView18itemDelegateForRowEi(void* qthis, int32_t arg0); // 4
  // proto:  Qt::DropAction QAbstractItemView::defaultDropAction();
extern void C_ZNK17QAbstractItemView17defaultDropActionEv(void* qthis); // 4
  // proto:  void QAbstractItemView::setCurrentIndex(const QModelIndex & index);
extern void C_ZN17QAbstractItemView15setCurrentIndexERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QAbstractItemView::QAbstractItemView(QWidget * parent);
extern void* C_ZN17QAbstractItemViewC2EP7QWidget(void* arg0); // 3
  // proto:  void QAbstractItemView::setDropIndicatorShown(bool enable);
extern void C_ZN17QAbstractItemView21setDropIndicatorShownEb(void* qthis, bool arg0); // 4
  // proto:  QItemSelectionModel * QAbstractItemView::selectionModel();
extern void* C_ZNK17QAbstractItemView14selectionModelEv(void* qthis); // 4
  // proto:  void QAbstractItemView::setAlternatingRowColors(bool enable);
extern void C_ZN17QAbstractItemView23setAlternatingRowColorsEb(void* qthis, bool arg0); // 4
  // proto:  void QAbstractItemView::~QAbstractItemView();
extern void C_ZN17QAbstractItemViewD2Ev(void* qthis); // 4
  // proto:  int QAbstractItemView::sizeHintForColumn(int column);
extern int32_t C_ZNK17QAbstractItemView17sizeHintForColumnEi(void* qthis, int32_t arg0); // 4
  // proto:  void QAbstractItemView::setItemDelegateForColumn(int column, QAbstractItemDelegate * delegate);
extern void C_ZN17QAbstractItemView24setItemDelegateForColumnEiP21QAbstractItemDelegate(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QAbstractItemView::clearSelection();
extern void C_ZN17QAbstractItemView14clearSelectionEv(void* qthis); // 4
  // proto:  bool QAbstractItemView::alternatingRowColors();
extern bool C_ZNK17QAbstractItemView20alternatingRowColorsEv(void* qthis); // 4
  // proto:  void QAbstractItemView::update(const QModelIndex & index);
extern void C_ZN17QAbstractItemView6updateERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QAbstractItemView::setModel(QAbstractItemModel * model);
extern void C_ZN17QAbstractItemView8setModelEP18QAbstractItemModel(void* qthis, void* arg0); // 4
  // proto:  QWidget * QAbstractItemView::indexWidget(const QModelIndex & index);
extern void* C_ZNK17QAbstractItemView11indexWidgetERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QAbstractItemDelegate * QAbstractItemView::itemDelegate();
extern void C_ZNK17QAbstractItemView12itemDelegateEv(void* qthis); // 4
  // proto:  QAbstractItemDelegate * QAbstractItemView::itemDelegate(const QModelIndex & index);
extern void C_ZNK17QAbstractItemView12itemDelegateERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  int QAbstractItemView::sizeHintForRow(int row);
extern int32_t C_ZNK17QAbstractItemView14sizeHintForRowEi(void* qthis, int32_t arg0); // 4
  // proto:  QAbstractItemView::SelectionBehavior QAbstractItemView::selectionBehavior();
extern void C_ZNK17QAbstractItemView17selectionBehaviorEv(void* qthis); // 4
  // proto:  QAbstractItemModel * QAbstractItemView::model();
extern void C_ZNK17QAbstractItemView5modelEv(void* qthis); // 4
  // proto:  void QAbstractItemView::setAutoScroll(bool enable);
extern void C_ZN17QAbstractItemView13setAutoScrollEb(void* qthis, bool arg0); // 4
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

// class sizeof(QAbstractItemView)=1
type QAbstractItemView struct {
  /*qbase*/ QAbstractScrollArea;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _iconSizeChanged QAbstractItemView_iconSizeChanged_signal;
//  _clicked QAbstractItemView_clicked_signal;
//  _viewportEntered QAbstractItemView_viewportEntered_signal;
//  _activated QAbstractItemView_activated_signal;
//  _pressed QAbstractItemView_pressed_signal;
//  _entered QAbstractItemView_entered_signal;
//  _doubleClicked QAbstractItemView_doubleClicked_signal;
}

// keyboardSearch(const class QString &)
func (this *QAbstractItemView) KeyboardSearch(args ...interface{}) () {
  // keyboardSearch(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView14keyboardSearchERK7QString
    // invoke: void keyboardSearch(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QAbstractItemView14keyboardSearchERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "keyboardSearch", args)
  }

  return
}

// setDragEnabled(_Bool)
func (this *QAbstractItemView) SetDragEnabled(args ...interface{}) () {
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
    C.C_ZN17QAbstractItemView14setDragEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setDragEnabled", args)
  }

  return
}

// dragDropOverwriteMode()
func (this *QAbstractItemView) DragDropOverwriteMode(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK17QAbstractItemView21dragDropOverwriteModeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractItemView", "dragDropOverwriteMode", args)
  }

  return
}

// openPersistentEditor(const class QModelIndex &)
func (this *QAbstractItemView) OpenPersistentEditor(args ...interface{}) () {
  // openPersistentEditor(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView20openPersistentEditorERK11QModelIndex
    // invoke: void openPersistentEditor(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QAbstractItemView20openPersistentEditorERK11QModelIndex(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "openPersistentEditor", args)
  }

  return
}

// selectionMode()
func (this *QAbstractItemView) SelectionMode(args ...interface{}) () {
  // selectionMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView13selectionModeEv
    // invoke: QAbstractItemView::SelectionMode selectionMode()
    C.C_ZNK17QAbstractItemView13selectionModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "selectionMode", args)
  }

  return
}

// autoScrollMargin()
func (this *QAbstractItemView) AutoScrollMargin(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK17QAbstractItemView16autoScrollMarginEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractItemView", "autoScrollMargin", args)
  }

  return
}

// rootIndex()
func (this *QAbstractItemView) RootIndex(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK17QAbstractItemView9rootIndexEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractItemView", "rootIndex", args)
  }

  return
}

// setSelectionModel(class QItemSelectionModel *)
func (this *QAbstractItemView) SetSelectionModel(args ...interface{}) () {
  // setSelectionModel(class QItemSelectionModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QItemSelectionModel{}) // "QItemSelectionModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView17setSelectionModelEP19QItemSelectionModel
    // invoke: void setSelectionModel(class QItemSelectionModel *)
    var arg0 = args[0].(*qtcore.QItemSelectionModel).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QAbstractItemView17setSelectionModelEP19QItemSelectionModel(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setSelectionModel", args)
  }

  return
}

// tabKeyNavigation()
func (this *QAbstractItemView) TabKeyNavigation(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK17QAbstractItemView16tabKeyNavigationEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractItemView", "tabKeyNavigation", args)
  }

  return
}

// metaObject()
func (this *QAbstractItemView) MetaObject(args ...interface{}) () {
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
    C.C_ZNK17QAbstractItemView10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "metaObject", args)
  }

  return
}

// edit(const class QModelIndex &)
func (this *QAbstractItemView) Edit(args ...interface{}) () {
  // edit(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView4editERK11QModelIndex
    // invoke: void edit(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QAbstractItemView4editERK11QModelIndex(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "edit", args)
  }

  return
}

// setIconSize(const class QSize &)
func (this *QAbstractItemView) SetIconSize(args ...interface{}) () {
  // setIconSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView11setIconSizeERK5QSize
    // invoke: void setIconSize(const class QSize &)
    var arg0 = args[0].(*qtcore.QSize).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QAbstractItemView11setIconSizeERK5QSize(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setIconSize", args)
  }

  return
}

// currentIndex()
func (this *QAbstractItemView) CurrentIndex(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK17QAbstractItemView12currentIndexEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractItemView", "currentIndex", args)
  }

  return
}

// closePersistentEditor(const class QModelIndex &)
func (this *QAbstractItemView) ClosePersistentEditor(args ...interface{}) () {
  // closePersistentEditor(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView21closePersistentEditorERK11QModelIndex
    // invoke: void closePersistentEditor(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QAbstractItemView21closePersistentEditorERK11QModelIndex(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "closePersistentEditor", args)
  }

  return
}

// verticalScrollMode()
func (this *QAbstractItemView) VerticalScrollMode(args ...interface{}) () {
  // verticalScrollMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView18verticalScrollModeEv
    // invoke: QAbstractItemView::ScrollMode verticalScrollMode()
    C.C_ZNK17QAbstractItemView18verticalScrollModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "verticalScrollMode", args)
  }

  return
}

// doItemsLayout()
func (this *QAbstractItemView) DoItemsLayout(args ...interface{}) () {
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
    C.C_ZN17QAbstractItemView13doItemsLayoutEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "doItemsLayout", args)
  }

  return
}

// scrollToBottom()
func (this *QAbstractItemView) ScrollToBottom(args ...interface{}) () {
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
    C.C_ZN17QAbstractItemView14scrollToBottomEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "scrollToBottom", args)
  }

  return
}

// setItemDelegateForRow(int, class QAbstractItemDelegate *)
func (this *QAbstractItemView) SetItemDelegateForRow(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QAbstractItemDelegate).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN17QAbstractItemView21setItemDelegateForRowEiP21QAbstractItemDelegate(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setItemDelegateForRow", args)
  }

  return
}

// setTabKeyNavigation(_Bool)
func (this *QAbstractItemView) SetTabKeyNavigation(args ...interface{}) () {
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
    C.C_ZN17QAbstractItemView19setTabKeyNavigationEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setTabKeyNavigation", args)
  }

  return
}

// iconSize()
func (this *QAbstractItemView) IconSize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK17QAbstractItemView8iconSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractItemView", "iconSize", args)
  }

  return
}

// selectAll()
func (this *QAbstractItemView) SelectAll(args ...interface{}) () {
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
    C.C_ZN17QAbstractItemView9selectAllEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "selectAll", args)
  }

  return
}

// setAutoScrollMargin(int)
func (this *QAbstractItemView) SetAutoScrollMargin(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN17QAbstractItemView19setAutoScrollMarginEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setAutoScrollMargin", args)
  }

  return
}

// itemDelegateForColumn(int)
func (this *QAbstractItemView) ItemDelegateForColumn(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK17QAbstractItemView21itemDelegateForColumnEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "itemDelegateForColumn", args)
  }

  return
}

// setDragDropOverwriteMode(_Bool)
func (this *QAbstractItemView) SetDragDropOverwriteMode(args ...interface{}) () {
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
    C.C_ZN17QAbstractItemView24setDragDropOverwriteModeEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setDragDropOverwriteMode", args)
  }

  return
}

// dragDropMode()
func (this *QAbstractItemView) DragDropMode(args ...interface{}) () {
  // dragDropMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView12dragDropModeEv
    // invoke: QAbstractItemView::DragDropMode dragDropMode()
    C.C_ZNK17QAbstractItemView12dragDropModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "dragDropMode", args)
  }

  return
}

// horizontalScrollMode()
func (this *QAbstractItemView) HorizontalScrollMode(args ...interface{}) () {
  // horizontalScrollMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView20horizontalScrollModeEv
    // invoke: QAbstractItemView::ScrollMode horizontalScrollMode()
    C.C_ZNK17QAbstractItemView20horizontalScrollModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "horizontalScrollMode", args)
  }

  return
}

// hasAutoScroll()
func (this *QAbstractItemView) HasAutoScroll(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK17QAbstractItemView13hasAutoScrollEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractItemView", "hasAutoScroll", args)
  }

  return
}

// sizeHintForIndex(const class QModelIndex &)
func (this *QAbstractItemView) SizeHintForIndex(args ...interface{}) (ret interface{}) {
  // sizeHintForIndex(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView16sizeHintForIndexERK11QModelIndex
    // invoke: QSize sizeHintForIndex(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK17QAbstractItemView16sizeHintForIndexERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractItemView", "sizeHintForIndex", args)
  }

  return
}

// setRootIndex(const class QModelIndex &)
func (this *QAbstractItemView) SetRootIndex(args ...interface{}) () {
  // setRootIndex(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView12setRootIndexERK11QModelIndex
    // invoke: void setRootIndex(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QAbstractItemView12setRootIndexERK11QModelIndex(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setRootIndex", args)
  }

  return
}

// editTriggers()
func (this *QAbstractItemView) EditTriggers(args ...interface{}) () {
  // editTriggers()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView12editTriggersEv
    // invoke: EditTriggers editTriggers()
    C.C_ZNK17QAbstractItemView12editTriggersEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "editTriggers", args)
  }

  return
}

// showDropIndicator()
func (this *QAbstractItemView) ShowDropIndicator(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK17QAbstractItemView17showDropIndicatorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractItemView", "showDropIndicator", args)
  }

  return
}

// textElideMode()
func (this *QAbstractItemView) TextElideMode(args ...interface{}) () {
  // textElideMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView13textElideModeEv
    // invoke: Qt::TextElideMode textElideMode()
    C.C_ZNK17QAbstractItemView13textElideModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "textElideMode", args)
  }

  return
}

// reset()
func (this *QAbstractItemView) Reset(args ...interface{}) () {
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
    C.C_ZN17QAbstractItemView5resetEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "reset", args)
  }

  return
}

// dragEnabled()
func (this *QAbstractItemView) DragEnabled(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK17QAbstractItemView11dragEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractItemView", "dragEnabled", args)
  }

  return
}

// setIndexWidget(const class QModelIndex &, class QWidget *)
func (this *QAbstractItemView) SetIndexWidget(args ...interface{}) () {
  // setIndexWidget(const class QModelIndex &, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView14setIndexWidgetERK11QModelIndexP7QWidget
    // invoke: void setIndexWidget(const class QModelIndex &, class QWidget *)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN17QAbstractItemView14setIndexWidgetERK11QModelIndexP7QWidget(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setIndexWidget", args)
  }

  return
}

// scrollToTop()
func (this *QAbstractItemView) ScrollToTop(args ...interface{}) () {
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
    C.C_ZN17QAbstractItemView11scrollToTopEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "scrollToTop", args)
  }

  return
}

// setItemDelegate(class QAbstractItemDelegate *)
func (this *QAbstractItemView) SetItemDelegate(args ...interface{}) () {
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
    var arg0 = args[0].(*QAbstractItemDelegate).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QAbstractItemView15setItemDelegateEP21QAbstractItemDelegate(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setItemDelegate", args)
  }

  return
}

// itemDelegateForRow(int)
func (this *QAbstractItemView) ItemDelegateForRow(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK17QAbstractItemView18itemDelegateForRowEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "itemDelegateForRow", args)
  }

  return
}

// defaultDropAction()
func (this *QAbstractItemView) DefaultDropAction(args ...interface{}) () {
  // defaultDropAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView17defaultDropActionEv
    // invoke: Qt::DropAction defaultDropAction()
    C.C_ZNK17QAbstractItemView17defaultDropActionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "defaultDropAction", args)
  }

  return
}

// setCurrentIndex(const class QModelIndex &)
func (this *QAbstractItemView) SetCurrentIndex(args ...interface{}) () {
  // setCurrentIndex(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView15setCurrentIndexERK11QModelIndex
    // invoke: void setCurrentIndex(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QAbstractItemView15setCurrentIndexERK11QModelIndex(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setCurrentIndex", args)
  }

  return
}

// QAbstractItemView(class QWidget *)
func GcfreeQAbstractItemView(this *QAbstractItemView) {
  qtrt.UniverseFree(this)
}
func NewQAbstractItemView(args ...interface{}) *QAbstractItemView {
  // QAbstractItemView(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemViewC1EP7QWidget
    // invoke: void QAbstractItemView(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QAbstractItemViewC2EP7QWidget(arg0)
    this := &QAbstractItemView{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQAbstractItemView)
    return this
  default:
    qtrt.ErrorResolve("QAbstractItemView", "QAbstractItemView", args)
  }

  return nil // QAbstractItemView{Qclsinst:qthis}
}

// setDropIndicatorShown(_Bool)
func (this *QAbstractItemView) SetDropIndicatorShown(args ...interface{}) () {
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
    C.C_ZN17QAbstractItemView21setDropIndicatorShownEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setDropIndicatorShown", args)
  }

  return
}

// selectionModel()
func (this *QAbstractItemView) SelectionModel(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK17QAbstractItemView14selectionModelEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QItemSelectionModel{}) // "QItemSelectionModel *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractItemView", "selectionModel", args)
  }

  return
}

// setAlternatingRowColors(_Bool)
func (this *QAbstractItemView) SetAlternatingRowColors(args ...interface{}) () {
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
    C.C_ZN17QAbstractItemView23setAlternatingRowColorsEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setAlternatingRowColors", args)
  }

  return
}

// ~QAbstractItemView()
func (this *QAbstractItemView) Free(args ...interface{}) () {
  // ~QAbstractItemView()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemViewD0Ev
    // invoke: void ~QAbstractItemView()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN17QAbstractItemViewD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QAbstractItemView", "~QAbstractItemView", args)
  }

  return
}

// sizeHintForColumn(int)
func (this *QAbstractItemView) SizeHintForColumn(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK17QAbstractItemView17sizeHintForColumnEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractItemView", "sizeHintForColumn", args)
  }

  return
}

// setItemDelegateForColumn(int, class QAbstractItemDelegate *)
func (this *QAbstractItemView) SetItemDelegateForColumn(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QAbstractItemDelegate).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN17QAbstractItemView24setItemDelegateForColumnEiP21QAbstractItemDelegate(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setItemDelegateForColumn", args)
  }

  return
}

// clearSelection()
func (this *QAbstractItemView) ClearSelection(args ...interface{}) () {
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
    C.C_ZN17QAbstractItemView14clearSelectionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "clearSelection", args)
  }

  return
}

// alternatingRowColors()
func (this *QAbstractItemView) AlternatingRowColors(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK17QAbstractItemView20alternatingRowColorsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractItemView", "alternatingRowColors", args)
  }

  return
}

// update(const class QModelIndex &)
func (this *QAbstractItemView) Update(args ...interface{}) () {
  // update(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView6updateERK11QModelIndex
    // invoke: void update(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QAbstractItemView6updateERK11QModelIndex(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "update", args)
  }

  return
}

// setModel(class QAbstractItemModel *)
func (this *QAbstractItemView) SetModel(args ...interface{}) () {
  // setModel(class QAbstractItemModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QAbstractItemModel{}) // "QAbstractItemModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractItemView8setModelEP18QAbstractItemModel
    // invoke: void setModel(class QAbstractItemModel *)
    var arg0 = args[0].(*qtcore.QAbstractItemModel).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QAbstractItemView8setModelEP18QAbstractItemModel(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setModel", args)
  }

  return
}

// indexWidget(const class QModelIndex &)
func (this *QAbstractItemView) IndexWidget(args ...interface{}) (ret interface{}) {
  // indexWidget(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView11indexWidgetERK11QModelIndex
    // invoke: QWidget * indexWidget(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK17QAbstractItemView11indexWidgetERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractItemView", "indexWidget", args)
  }

  return
}

// itemDelegate()
func (this *QAbstractItemView) ItemDelegate(args ...interface{}) () {
  // itemDelegate()
  // itemDelegate(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView12itemDelegateEv
    // invoke: QAbstractItemDelegate * itemDelegate()
    C.C_ZNK17QAbstractItemView12itemDelegateEv(this.Qclsinst)
  case 1:
    // invoke: _ZNK17QAbstractItemView12itemDelegateERK11QModelIndex
    // invoke: QAbstractItemDelegate * itemDelegate(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK17QAbstractItemView12itemDelegateERK11QModelIndex(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "itemDelegate", args)
  }

  return
}

// sizeHintForRow(int)
func (this *QAbstractItemView) SizeHintForRow(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK17QAbstractItemView14sizeHintForRowEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractItemView", "sizeHintForRow", args)
  }

  return
}

// selectionBehavior()
func (this *QAbstractItemView) SelectionBehavior(args ...interface{}) () {
  // selectionBehavior()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView17selectionBehaviorEv
    // invoke: QAbstractItemView::SelectionBehavior selectionBehavior()
    C.C_ZNK17QAbstractItemView17selectionBehaviorEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "selectionBehavior", args)
  }

  return
}

// model()
func (this *QAbstractItemView) Model(args ...interface{}) () {
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
    C.C_ZNK17QAbstractItemView5modelEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "model", args)
  }

  return
}

// setAutoScroll(_Bool)
func (this *QAbstractItemView) SetAutoScroll(args ...interface{}) () {
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
    C.C_ZN17QAbstractItemView13setAutoScrollEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setAutoScroll", args)
  }

  return
}

// <= body block end

