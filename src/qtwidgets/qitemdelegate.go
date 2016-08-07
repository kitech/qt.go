package qtwidgets
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
  // proto:  QItemEditorFactory * QItemDelegate::itemEditorFactory();
extern void* C_ZNK13QItemDelegate17itemEditorFactoryEv(void* qthis); // 4
  // proto:  QWidget * QItemDelegate::createEditor(QWidget * parent, const QStyleOptionViewItem & option, const QModelIndex & index);
extern void* C_ZNK13QItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QItemDelegate::updateEditorGeometry(QWidget * editor, const QStyleOptionViewItem & option, const QModelIndex & index);
extern void C_ZNK13QItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QItemDelegate::paint(QPainter * painter, const QStyleOptionViewItem & option, const QModelIndex & index);
extern void C_ZNK13QItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QItemDelegate::setModelData(QWidget * editor, QAbstractItemModel * model, const QModelIndex & index);
extern void C_ZNK13QItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QItemDelegate::~QItemDelegate();
extern void C_ZN13QItemDelegateD2Ev(void* qthis); // 4
  // proto:  bool QItemDelegate::hasClipping();
extern bool C_ZNK13QItemDelegate11hasClippingEv(void* qthis); // 4
  // proto:  void QItemDelegate::setEditorData(QWidget * editor, const QModelIndex & index);
extern void C_ZNK13QItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QSize QItemDelegate::sizeHint(const QStyleOptionViewItem & option, const QModelIndex & index);
extern void* C_ZNK13QItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QItemDelegate::setClipping(bool clip);
extern void C_ZN13QItemDelegate11setClippingEb(void* qthis, bool arg0); // 4
  // proto:  const QMetaObject * QItemDelegate::metaObject();
extern void C_ZNK13QItemDelegate10metaObjectEv(void* qthis); // 4
  // proto:  void QItemDelegate::setItemEditorFactory(QItemEditorFactory * factory);
extern void C_ZN13QItemDelegate20setItemEditorFactoryEP18QItemEditorFactory(void* qthis, void* arg0); // 4
  // proto:  void QItemDelegate::QItemDelegate(QObject * parent);
extern void* C_ZN13QItemDelegateC2EP7QObject(void* arg0); // 3
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

// class sizeof(QItemDelegate)=1
type QItemDelegate struct {
  /*qbase*/ QAbstractItemDelegate;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// itemEditorFactory()
func (this *QItemDelegate) Itemeditorfactory(args ...interface{}) (ret interface{}) {
  // itemEditorFactory()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QItemDelegate17itemEditorFactoryEv
    // invoke: QItemEditorFactory * itemEditorFactory()
    var ret0 = C.C_ZNK13QItemDelegate17itemEditorFactoryEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QItemEditorFactory{}) // "QItemEditorFactory *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemDelegate", "itemEditorFactory", args)
  }

  return
}

// createEditor(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QItemDelegate) Createeditor(args ...interface{}) (ret interface{}) {
  // createEditor(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionViewItem{}) // "const QStyleOptionViewItem &"
  vtys[0][2] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex
    // invoke: QWidget * createEditor(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStyleOptionViewItem).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK13QItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemDelegate", "createEditor", args)
  }

  return
}

// updateEditorGeometry(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QItemDelegate) Updateeditorgeometry(args ...interface{}) () {
  // updateEditorGeometry(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionViewItem{}) // "const QStyleOptionViewItem &"
  vtys[0][2] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex
    // invoke: void updateEditorGeometry(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStyleOptionViewItem).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZNK13QItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QItemDelegate", "updateEditorGeometry", args)
  }

  return
}

// paint(class QPainter *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QItemDelegate) Paint(args ...interface{}) () {
  // paint(class QPainter *, const class QStyleOptionViewItem &, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionViewItem{}) // "const QStyleOptionViewItem &"
  vtys[0][2] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex
    // invoke: void paint(class QPainter *, const class QStyleOptionViewItem &, const class QModelIndex &)
    var arg0 = args[0].(*qtgui.QPainter).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStyleOptionViewItem).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZNK13QItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QItemDelegate", "paint", args)
  }

  return
}

// setModelData(class QWidget *, class QAbstractItemModel *, const class QModelIndex &)
func (this *QItemDelegate) Setmodeldata(args ...interface{}) () {
  // setModelData(class QWidget *, class QAbstractItemModel *, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(qtcore.QAbstractItemModel{}) // "QAbstractItemModel *"
  vtys[0][2] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex
    // invoke: void setModelData(class QWidget *, class QAbstractItemModel *, const class QModelIndex &)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QAbstractItemModel).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZNK13QItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QItemDelegate", "setModelData", args)
  }

  return
}

// ~QItemDelegate()
func (this *QItemDelegate) Freeqitemdelegate(args ...interface{}) () {
  // ~QItemDelegate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QItemDelegateD0Ev
    // invoke: void ~QItemDelegate()
    C.C_ZN13QItemDelegateD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QItemDelegate", "~QItemDelegate", args)
  }

  return
}

// hasClipping()
func (this *QItemDelegate) Hasclipping(args ...interface{}) (ret interface{}) {
  // hasClipping()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QItemDelegate11hasClippingEv
    // invoke: bool hasClipping()
    var ret0 = C.C_ZNK13QItemDelegate11hasClippingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemDelegate", "hasClipping", args)
  }

  return
}

// setEditorData(class QWidget *, const class QModelIndex &)
func (this *QItemDelegate) Seteditordata(args ...interface{}) () {
  // setEditorData(class QWidget *, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex
    // invoke: void setEditorData(class QWidget *, const class QModelIndex &)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZNK13QItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QItemDelegate", "setEditorData", args)
  }

  return
}

// sizeHint(const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QItemDelegate) Sizehint(args ...interface{}) (ret interface{}) {
  // sizeHint(const class QStyleOptionViewItem &, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStyleOptionViewItem{}) // "const QStyleOptionViewItem &"
  vtys[0][1] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex
    // invoke: QSize sizeHint(const class QStyleOptionViewItem &, const class QModelIndex &)
    var arg0 = args[0].(*QStyleOptionViewItem).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemDelegate", "sizeHint", args)
  }

  return
}

// setClipping(_Bool)
func (this *QItemDelegate) Setclipping(args ...interface{}) () {
  // setClipping(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QItemDelegate11setClippingEb
    // invoke: void setClipping(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QItemDelegate11setClippingEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QItemDelegate", "setClipping", args)
  }

  return
}

// metaObject()
func (this *QItemDelegate) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QItemDelegate10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK13QItemDelegate10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QItemDelegate", "metaObject", args)
  }

  return
}

// setItemEditorFactory(class QItemEditorFactory *)
func (this *QItemDelegate) Setitemeditorfactory(args ...interface{}) () {
  // setItemEditorFactory(class QItemEditorFactory *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QItemEditorFactory{}) // "QItemEditorFactory *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QItemDelegate20setItemEditorFactoryEP18QItemEditorFactory
    // invoke: void setItemEditorFactory(class QItemEditorFactory *)
    var arg0 = args[0].(*QItemEditorFactory).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QItemDelegate20setItemEditorFactoryEP18QItemEditorFactory(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QItemDelegate", "setItemEditorFactory", args)
  }

  return
}

// QItemDelegate(class QObject *)
func NewQItemDelegate(args ...interface{}) *QItemDelegate {
  // QItemDelegate(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QItemDelegateC1EP7QObject
    // invoke: void QItemDelegate(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QItemDelegateC2EP7QObject(arg0)
    return &QItemDelegate{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QItemDelegate", "QItemDelegate", args)
  }

  return nil // QItemDelegate{Qclsinst:qthis}
}

// <= body block end

