package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtWidgets/qstyleditemdelegate.h
// dst-file: /src/widgets/qstyleditemdelegate.go
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
  // proto:  QItemEditorFactory * QStyledItemDelegate::itemEditorFactory();
extern void* C_ZNK19QStyledItemDelegate17itemEditorFactoryEv(void* qthis); // 4
  // proto:  QWidget * QStyledItemDelegate::createEditor(QWidget * parent, const QStyleOptionViewItem & option, const QModelIndex & index);
extern void* C_ZNK19QStyledItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QStyledItemDelegate::updateEditorGeometry(QWidget * editor, const QStyleOptionViewItem & option, const QModelIndex & index);
extern void C_ZNK19QStyledItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QStyledItemDelegate::QStyledItemDelegate(QObject * parent);
extern void* C_ZN19QStyledItemDelegateC2EP7QObject(void* arg0); // 3
  // proto:  void QStyledItemDelegate::paint(QPainter * painter, const QStyleOptionViewItem & option, const QModelIndex & index);
extern void C_ZNK19QStyledItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QStyledItemDelegate::setModelData(QWidget * editor, QAbstractItemModel * model, const QModelIndex & index);
extern void C_ZNK19QStyledItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  QString QStyledItemDelegate::displayText(const QVariant & value, const QLocale & locale);
extern void* C_ZNK19QStyledItemDelegate11displayTextERK8QVariantRK7QLocale(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QStyledItemDelegate::setEditorData(QWidget * editor, const QModelIndex & index);
extern void C_ZNK19QStyledItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QStyledItemDelegate::~QStyledItemDelegate();
extern void C_ZN19QStyledItemDelegateD2Ev(void* qthis); // 4
  // proto:  QSize QStyledItemDelegate::sizeHint(const QStyleOptionViewItem & option, const QModelIndex & index);
extern void* C_ZNK19QStyledItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex(void* qthis, void* arg0, void* arg1); // 4
  // proto:  const QMetaObject * QStyledItemDelegate::metaObject();
extern void C_ZNK19QStyledItemDelegate10metaObjectEv(void* qthis); // 4
  // proto:  void QStyledItemDelegate::setItemEditorFactory(QItemEditorFactory * factory);
extern void C_ZN19QStyledItemDelegate20setItemEditorFactoryEP18QItemEditorFactory(void* qthis, void* arg0); // 4
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

// class sizeof(QStyledItemDelegate)=1
type QStyledItemDelegate struct {
  /*qbase*/ QAbstractItemDelegate;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// itemEditorFactory()
func (this *QStyledItemDelegate) Itemeditorfactory(args ...interface{}) (ret interface{}) {
  // itemEditorFactory()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QStyledItemDelegate17itemEditorFactoryEv
    // invoke: QItemEditorFactory * itemEditorFactory()
    var ret0 = C.C_ZNK19QStyledItemDelegate17itemEditorFactoryEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QItemEditorFactory{}) // "QItemEditorFactory *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStyledItemDelegate", "itemEditorFactory", args)
  }

  return
}

// createEditor(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QStyledItemDelegate) Createeditor(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK19QStyledItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex
    // invoke: QWidget * createEditor(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStyleOptionViewItem).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK19QStyledItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStyledItemDelegate", "createEditor", args)
  }

  return
}

// updateEditorGeometry(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QStyledItemDelegate) Updateeditorgeometry(args ...interface{}) () {
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
    // invoke: _ZNK19QStyledItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex
    // invoke: void updateEditorGeometry(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStyleOptionViewItem).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZNK19QStyledItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStyledItemDelegate", "updateEditorGeometry", args)
  }

  return
}

// QStyledItemDelegate(class QObject *)
func NewQStyledItemDelegate(args ...interface{}) *QStyledItemDelegate {
  // QStyledItemDelegate(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QStyledItemDelegateC1EP7QObject
    // invoke: void QStyledItemDelegate(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QStyledItemDelegateC2EP7QObject(arg0)
    return &QStyledItemDelegate{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyledItemDelegate", "QStyledItemDelegate", args)
  }

  return nil // QStyledItemDelegate{Qclsinst:qthis}
}

// paint(class QPainter *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QStyledItemDelegate) Paint(args ...interface{}) () {
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
    // invoke: _ZNK19QStyledItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex
    // invoke: void paint(class QPainter *, const class QStyleOptionViewItem &, const class QModelIndex &)
    var arg0 = args[0].(*QPainter).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStyleOptionViewItem).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZNK19QStyledItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStyledItemDelegate", "paint", args)
  }

  return
}

// setModelData(class QWidget *, class QAbstractItemModel *, const class QModelIndex &)
func (this *QStyledItemDelegate) Setmodeldata(args ...interface{}) () {
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
    // invoke: _ZNK19QStyledItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex
    // invoke: void setModelData(class QWidget *, class QAbstractItemModel *, const class QModelIndex &)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QAbstractItemModel).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZNK19QStyledItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStyledItemDelegate", "setModelData", args)
  }

  return
}

// displayText(const class QVariant &, const class QLocale &)
func (this *QStyledItemDelegate) Displaytext(args ...interface{}) (ret interface{}) {
  // displayText(const class QVariant &, const class QLocale &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVariant{}) // "const QVariant &"
  vtys[0][1] = reflect.TypeOf(QLocale{}) // "const QLocale &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QStyledItemDelegate11displayTextERK8QVariantRK7QLocale
    // invoke: QString displayText(const class QVariant &, const class QLocale &)
    var arg0 = args[0].(*QVariant).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QLocale).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK19QStyledItemDelegate11displayTextERK8QVariantRK7QLocale(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStyledItemDelegate", "displayText", args)
  }

  return
}

// setEditorData(class QWidget *, const class QModelIndex &)
func (this *QStyledItemDelegate) Seteditordata(args ...interface{}) () {
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
    // invoke: _ZNK19QStyledItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex
    // invoke: void setEditorData(class QWidget *, const class QModelIndex &)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZNK19QStyledItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStyledItemDelegate", "setEditorData", args)
  }

  return
}

// ~QStyledItemDelegate()
func (this *QStyledItemDelegate) Freeqstyleditemdelegate(args ...interface{}) () {
  // ~QStyledItemDelegate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QStyledItemDelegateD0Ev
    // invoke: void ~QStyledItemDelegate()
    C.C_ZN19QStyledItemDelegateD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStyledItemDelegate", "~QStyledItemDelegate", args)
  }

  return
}

// sizeHint(const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QStyledItemDelegate) Sizehint(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK19QStyledItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex
    // invoke: QSize sizeHint(const class QStyleOptionViewItem &, const class QModelIndex &)
    var arg0 = args[0].(*QStyleOptionViewItem).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK19QStyledItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStyledItemDelegate", "sizeHint", args)
  }

  return
}

// metaObject()
func (this *QStyledItemDelegate) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QStyledItemDelegate10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK19QStyledItemDelegate10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStyledItemDelegate", "metaObject", args)
  }

  return
}

// setItemEditorFactory(class QItemEditorFactory *)
func (this *QStyledItemDelegate) Setitemeditorfactory(args ...interface{}) () {
  // setItemEditorFactory(class QItemEditorFactory *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QItemEditorFactory{}) // "QItemEditorFactory *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QStyledItemDelegate20setItemEditorFactoryEP18QItemEditorFactory
    // invoke: void setItemEditorFactory(class QItemEditorFactory *)
    var arg0 = args[0].(*QItemEditorFactory).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QStyledItemDelegate20setItemEditorFactoryEP18QItemEditorFactory(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStyledItemDelegate", "setItemEditorFactory", args)
  }

  return
}

// <= body block end

