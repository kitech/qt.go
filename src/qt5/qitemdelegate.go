package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qitemdelegate.h
// dst-file: /src/widgets/qitemdelegate.go
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
// class sizeof(QItemDelegate)=1
type QItemDelegate struct {
  /*qbase*/ QAbstractItemDelegate;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QItemDelegate) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QItemDelegate10metaObjectEv
  default:
    qtrt.ErrorResolve("QItemDelegate", "metaObject", args)
  }

}


func (this *QItemDelegate) sizeHint(args ...interface{}) () {
  // sizeHint(const class QStyleOptionViewItem &, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStyleOptionViewItem{}) // "const QStyleOptionViewItem &"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex
  default:
    qtrt.ErrorResolve("QItemDelegate", "sizeHint", args)
  }

}


func (this *QItemDelegate) setModelData(args ...interface{}) () {
  // setModelData(class QWidget *, class QAbstractItemModel *, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(QAbstractItemModel{}) // "QAbstractItemModel *"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex
  default:
    qtrt.ErrorResolve("QItemDelegate", "setModelData", args)
  }

}


func (this *QItemDelegate) setItemEditorFactory(args ...interface{}) () {
  // setItemEditorFactory(class QItemEditorFactory *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QItemEditorFactory{}) // "QItemEditorFactory *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QItemDelegate20setItemEditorFactoryEP18QItemEditorFactory
  default:
    qtrt.ErrorResolve("QItemDelegate", "setItemEditorFactory", args)
  }

}


func (this *QItemDelegate) setClipping(args ...interface{}) () {
  // setClipping(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QItemDelegate11setClippingEb
  default:
    qtrt.ErrorResolve("QItemDelegate", "setClipping", args)
  }

}


func (this *QItemDelegate) updateEditorGeometry(args ...interface{}) () {
  // updateEditorGeometry(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionViewItem{}) // "const QStyleOptionViewItem &"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex
  default:
    qtrt.ErrorResolve("QItemDelegate", "updateEditorGeometry", args)
  }

}


func (this *QItemDelegate) itemEditorFactory(args ...interface{}) () {
  // itemEditorFactory()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QItemDelegate17itemEditorFactoryEv
  default:
    qtrt.ErrorResolve("QItemDelegate", "itemEditorFactory", args)
  }

}


func (this *QItemDelegate) createEditor(args ...interface{}) () {
  // createEditor(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionViewItem{}) // "const QStyleOptionViewItem &"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex
  default:
    qtrt.ErrorResolve("QItemDelegate", "createEditor", args)
  }

}


func NewQItemDelegate(args ...interface{}) QItemDelegate {
  return QItemDelegate{}
}


func (this *QItemDelegate) hasClipping(args ...interface{}) () {
  // hasClipping()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QItemDelegate11hasClippingEv
  default:
    qtrt.ErrorResolve("QItemDelegate", "hasClipping", args)
  }

}


func (this *QItemDelegate) FreeQItemDelegate(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QItemDelegate", "~QItemDelegate", args)
  }

}


func (this *QItemDelegate) paint(args ...interface{}) () {
  // paint(class QPainter *, const class QStyleOptionViewItem &, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionViewItem{}) // "const QStyleOptionViewItem &"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex
  default:
    qtrt.ErrorResolve("QItemDelegate", "paint", args)
  }

}


func (this *QItemDelegate) setEditorData(args ...interface{}) () {
  // setEditorData(class QWidget *, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex
  default:
    qtrt.ErrorResolve("QItemDelegate", "setEditorData", args)
  }

}

// <= body block end

