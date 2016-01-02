package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtWidgets/qitemdelegate.h
// dst-file: /src/widgets/qitemdelegate.go
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
  // proto:  const QMetaObject * QItemDelegate::metaObject();
extern void _ZNK13QItemDelegate10metaObjectEv(void* qthis);
  // proto:  QSize QItemDelegate::sizeHint(const QStyleOptionViewItem & option, const QModelIndex & index);
extern void _ZNK13QItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex(void* qthis, void* arg0, void* arg1);
  // proto:  void QItemDelegate::setModelData(QWidget * editor, QAbstractItemModel * model, const QModelIndex & index);
extern void _ZNK13QItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QItemDelegate::setItemEditorFactory(QItemEditorFactory * factory);
extern void _ZN13QItemDelegate20setItemEditorFactoryEP18QItemEditorFactory(void* qthis, void* arg0);
  // proto:  void QItemDelegate::setClipping(bool clip);
extern void _ZN13QItemDelegate11setClippingEb(void* qthis, bool arg0);
  // proto:  void QItemDelegate::updateEditorGeometry(QWidget * editor, const QStyleOptionViewItem & option, const QModelIndex & index);
extern void _ZNK13QItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  QItemEditorFactory * QItemDelegate::itemEditorFactory();
extern void _ZNK13QItemDelegate17itemEditorFactoryEv(void* qthis);
  // proto:  QWidget * QItemDelegate::createEditor(QWidget * parent, const QStyleOptionViewItem & option, const QModelIndex & index);
extern void _ZNK13QItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QItemDelegate::QItemDelegate(QObject * parent);
extern void* dector_ZN13QItemDelegateC1EP7QObject(void* arg0);
extern void _ZN13QItemDelegateC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QItemDelegate::QItemDelegate(const QItemDelegate & );
extern void* dector_ZN13QItemDelegateC1ERKS_(void* arg0);
extern void _ZN13QItemDelegateC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QItemDelegate::hasClipping();
extern void _ZNK13QItemDelegate11hasClippingEv(void* qthis);
  // proto:  void QItemDelegate::~QItemDelegate();
extern void _ZN13QItemDelegateD0Ev(void* qthis);
  // proto:  void QItemDelegate::paint(QPainter * painter, const QStyleOptionViewItem & option, const QModelIndex & index);
extern void _ZNK13QItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QItemDelegate::setEditorData(QWidget * editor, const QModelIndex & index);
extern void _ZNK13QItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex(void* qthis, void* arg0, void* arg1);
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

// class sizeof(QItemDelegate)=1
type QItemDelegate struct {
  /*qbase*/ QAbstractItemDelegate;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  const QMetaObject * QItemDelegate::metaObject();
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

  // proto:  QSize QItemDelegate::sizeHint(const QStyleOptionViewItem & option, const QModelIndex & index);
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

  // proto:  void QItemDelegate::setModelData(QWidget * editor, QAbstractItemModel * model, const QModelIndex & index);
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

  // proto:  void QItemDelegate::setItemEditorFactory(QItemEditorFactory * factory);
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

  // proto:  void QItemDelegate::setClipping(bool clip);
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

  // proto:  void QItemDelegate::updateEditorGeometry(QWidget * editor, const QStyleOptionViewItem & option, const QModelIndex & index);
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

  // proto:  QItemEditorFactory * QItemDelegate::itemEditorFactory();
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

  // proto:  QWidget * QItemDelegate::createEditor(QWidget * parent, const QStyleOptionViewItem & option, const QModelIndex & index);
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

  // proto:  void QItemDelegate::QItemDelegate(QObject * parent);
func NewQItemDelegate(args ...interface{}) QItemDelegate {
  return QItemDelegate{}
}

  // proto:  bool QItemDelegate::hasClipping();
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

  // proto:  void QItemDelegate::~QItemDelegate();
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

  // proto:  void QItemDelegate::paint(QPainter * painter, const QStyleOptionViewItem & option, const QModelIndex & index);
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

  // proto:  void QItemDelegate::setEditorData(QWidget * editor, const QModelIndex & index);
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

