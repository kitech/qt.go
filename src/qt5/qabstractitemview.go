package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtWidgets/qabstractitemview.h
// dst-file: /src/widgets/qabstractitemview.go
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
// class sizeof(QAbstractItemView)=1
type QAbstractItemView struct {
  /*qbase*/ QAbstractScrollArea;
  qclsinst uint64 /* *mut c_void*/;
//  _iconSizeChanged QAbstractItemView_iconSizeChanged_signal;
//  _clicked QAbstractItemView_clicked_signal;
//  _viewportEntered QAbstractItemView_viewportEntered_signal;
//  _activated QAbstractItemView_activated_signal;
//  _pressed QAbstractItemView_pressed_signal;
//  _entered QAbstractItemView_entered_signal;
//  _doubleClicked QAbstractItemView_doubleClicked_signal;
}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "indexWidget", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "scrollToBottom", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setDropIndicatorShown", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "sizeHintForIndex", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setItemDelegateForColumn", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "dragEnabled", args)
 }

}


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
  case 1:
    // invoke: _ZNK17QAbstractItemView12itemDelegateEv
  default:
    qtrt.ErrorResolve("QAbstractItemView", "itemDelegate", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "keyboardSearch", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "rootIndex", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "itemDelegateForColumn", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "selectionModel", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "reset", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setItemDelegateForRow", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setRootIndex", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setAutoScrollMargin", args)
 }

}


func (this *QAbstractItemView) visualRect(args ...interface{}) () {
  // visualRect(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView10visualRectERK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemView", "visualRect", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "doItemsLayout", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "model", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "iconSize", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setItemDelegate", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setDragEnabled", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "currentIndex", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "sizeHintForRow", args)
 }

}


func NewQAbstractItemView(args ...interface{})() {
}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "showDropIndicator", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "hasAutoScroll", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "selectAll", args)
 }

}


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
    // invoke: _ZN17QAbstractItemView4editERK11QModelIndexNS_11EditTriggerEP6QEvent
  case 1:
    // invoke: _ZN17QAbstractItemView4editERK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemView", "edit", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setAlternatingRowColors", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "sizeHintForColumn", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setIconSize", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "closePersistentEditor", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setDragDropOverwriteMode", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "clearSelection", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "scrollToTop", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setSelectionModel", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setCurrentIndex", args)
 }

}


func (this *QAbstractItemView) indexAt(args ...interface{}) () {
  // indexAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QAbstractItemView7indexAtERK6QPoint
  default:
    qtrt.ErrorResolve("QAbstractItemView", "indexAt", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setTabKeyNavigation", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setIndexWidget", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setAutoScroll", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "setModel", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "update", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "metaObject", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "openPersistentEditor", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "itemDelegateForRow", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "dragDropOverwriteMode", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "tabKeyNavigation", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "autoScrollMargin", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractItemView", "alternatingRowColors", args)
 }

}

// <= body block end

