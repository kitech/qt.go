package qtcore
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
// src-file: /QtCore/qabstractproxymodel.h
// dst-file: /src/core/qabstractproxymodel.go
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
  // proto:  QItemSelection QAbstractProxyModel::mapSelectionToSource(const QItemSelection & selection);
extern void* C_ZNK19QAbstractProxyModel20mapSelectionToSourceERK14QItemSelection(void* qthis, void* arg0); // 4
  // proto:  Qt::DropActions QAbstractProxyModel::supportedDragActions();
extern void C_ZNK19QAbstractProxyModel20supportedDragActionsEv(void* qthis); // 4
  // proto:  void QAbstractProxyModel::setSourceModel(QAbstractItemModel * sourceModel);
extern void C_ZN19QAbstractProxyModel14setSourceModelEP18QAbstractItemModel(void* qthis, void* arg0); // 4
  // proto:  QModelIndex QAbstractProxyModel::sibling(int row, int column, const QModelIndex & idx);
extern void* C_ZNK19QAbstractProxyModel7siblingEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QAbstractProxyModel::QAbstractProxyModel(QObject * parent);
extern void* C_ZN19QAbstractProxyModelC2EP7QObject(void* arg0); // 3
  // proto:  QSize QAbstractProxyModel::span(const QModelIndex & index);
extern void* C_ZNK19QAbstractProxyModel4spanERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QModelIndex QAbstractProxyModel::buddy(const QModelIndex & index);
extern void* C_ZNK19QAbstractProxyModel5buddyERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QAbstractProxyModel::hasChildren(const QModelIndex & parent);
extern bool C_ZNK19QAbstractProxyModel11hasChildrenERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QAbstractProxyModel::submit();
extern bool C_ZN19QAbstractProxyModel6submitEv(void* qthis); // 4
  // proto:  bool QAbstractProxyModel::setData(const QModelIndex & index, const QVariant & value, int role);
extern bool C_ZN19QAbstractProxyModel7setDataERK11QModelIndexRK8QVarianti(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  bool QAbstractProxyModel::canFetchMore(const QModelIndex & parent);
extern bool C_ZNK19QAbstractProxyModel12canFetchMoreERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QMap<int, QVariant> QAbstractProxyModel::itemData(const QModelIndex & index);
extern void C_ZNK19QAbstractProxyModel8itemDataERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QItemSelection QAbstractProxyModel::mapSelectionFromSource(const QItemSelection & selection);
extern void* C_ZNK19QAbstractProxyModel22mapSelectionFromSourceERK14QItemSelection(void* qthis, void* arg0); // 4
  // proto:  QVariant QAbstractProxyModel::data(const QModelIndex & proxyIndex, int role);
extern void* C_ZNK19QAbstractProxyModel4dataERK11QModelIndexi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QStringList QAbstractProxyModel::mimeTypes();
extern void C_ZNK19QAbstractProxyModel9mimeTypesEv(void* qthis); // 4
  // proto:  const QMetaObject * QAbstractProxyModel::metaObject();
extern void C_ZNK19QAbstractProxyModel10metaObjectEv(void* qthis); // 4
  // proto:  void QAbstractProxyModel::revert();
extern void C_ZN19QAbstractProxyModel6revertEv(void* qthis); // 4
  // proto:  void QAbstractProxyModel::~QAbstractProxyModel();
extern void C_ZN19QAbstractProxyModelD2Ev(void* qthis); // 4
  // proto:  Qt::ItemFlags QAbstractProxyModel::flags(const QModelIndex & index);
extern void C_ZNK19QAbstractProxyModel5flagsERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QAbstractItemModel * QAbstractProxyModel::sourceModel();
extern void C_ZNK19QAbstractProxyModel11sourceModelEv(void* qthis); // 4
  // proto:  void QAbstractProxyModel::fetchMore(const QModelIndex & parent);
extern void C_ZN19QAbstractProxyModel9fetchMoreERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  Qt::DropActions QAbstractProxyModel::supportedDropActions();
extern void C_ZNK19QAbstractProxyModel20supportedDropActionsEv(void* qthis); // 4
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

// class sizeof(QAbstractProxyModel)=1
type QAbstractProxyModel struct {
  /*qbase*/ QAbstractItemModel;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _sourceModelChanged QAbstractProxyModel_sourceModelChanged_signal;
}

// mapSelectionToSource(const class QItemSelection &)
func (this *QAbstractProxyModel) Mapselectiontosource(args ...interface{}) (ret interface{}) {
  // mapSelectionToSource(const class QItemSelection &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QItemSelection{}) // "const QItemSelection &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractProxyModel20mapSelectionToSourceERK14QItemSelection
    // invoke: QItemSelection mapSelectionToSource(const class QItemSelection &)
    var arg0 = args[0].(*QItemSelection).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QAbstractProxyModel20mapSelectionToSourceERK14QItemSelection(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QItemSelection{}) // "QItemSelection"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "mapSelectionToSource", args)
  }

  return
}

// supportedDragActions()
func (this *QAbstractProxyModel) Supporteddragactions(args ...interface{}) () {
  // supportedDragActions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractProxyModel20supportedDragActionsEv
    // invoke: Qt::DropActions supportedDragActions()
    C.C_ZNK19QAbstractProxyModel20supportedDragActionsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "supportedDragActions", args)
  }

  return
}

// setSourceModel(class QAbstractItemModel *)
func (this *QAbstractProxyModel) Setsourcemodel(args ...interface{}) () {
  // setSourceModel(class QAbstractItemModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemModel{}) // "QAbstractItemModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractProxyModel14setSourceModelEP18QAbstractItemModel
    // invoke: void setSourceModel(class QAbstractItemModel *)
    var arg0 = args[0].(*QAbstractItemModel).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QAbstractProxyModel14setSourceModelEP18QAbstractItemModel(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "setSourceModel", args)
  }

  return
}

// sibling(int, int, const class QModelIndex &)
func (this *QAbstractProxyModel) Sibling(args ...interface{}) (ret interface{}) {
  // sibling(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractProxyModel7siblingEiiRK11QModelIndex
    // invoke: QModelIndex sibling(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK19QAbstractProxyModel7siblingEiiRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "sibling", args)
  }

  return
}

// QAbstractProxyModel(class QObject *)
func NewQAbstractProxyModel(args ...interface{}) *QAbstractProxyModel {
  // QAbstractProxyModel(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractProxyModelC1EP7QObject
    // invoke: void QAbstractProxyModel(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QAbstractProxyModelC2EP7QObject(arg0)
    return &QAbstractProxyModel{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "QAbstractProxyModel", args)
  }

  return nil // QAbstractProxyModel{Qclsinst:qthis}
}

// span(const class QModelIndex &)
func (this *QAbstractProxyModel) Span(args ...interface{}) (ret interface{}) {
  // span(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractProxyModel4spanERK11QModelIndex
    // invoke: QSize span(const class QModelIndex &)
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QAbstractProxyModel4spanERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "span", args)
  }

  return
}

// buddy(const class QModelIndex &)
func (this *QAbstractProxyModel) Buddy(args ...interface{}) (ret interface{}) {
  // buddy(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractProxyModel5buddyERK11QModelIndex
    // invoke: QModelIndex buddy(const class QModelIndex &)
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QAbstractProxyModel5buddyERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "buddy", args)
  }

  return
}

// hasChildren(const class QModelIndex &)
func (this *QAbstractProxyModel) Haschildren(args ...interface{}) (ret interface{}) {
  // hasChildren(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractProxyModel11hasChildrenERK11QModelIndex
    // invoke: bool hasChildren(const class QModelIndex &)
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QAbstractProxyModel11hasChildrenERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "hasChildren", args)
  }

  return
}

// submit()
func (this *QAbstractProxyModel) Submit(args ...interface{}) (ret interface{}) {
  // submit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractProxyModel6submitEv
    // invoke: bool submit()
    var ret0 = C.C_ZN19QAbstractProxyModel6submitEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "submit", args)
  }

  return
}

// setData(const class QModelIndex &, const class QVariant &, int)
func (this *QAbstractProxyModel) Setdata(args ...interface{}) (ret interface{}) {
  // setData(const class QModelIndex &, const class QVariant &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractProxyModel7setDataERK11QModelIndexRK8QVarianti
    // invoke: bool setData(const class QModelIndex &, const class QVariant &, int)
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVariant).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN19QAbstractProxyModel7setDataERK11QModelIndexRK8QVarianti(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "setData", args)
  }

  return
}

// canFetchMore(const class QModelIndex &)
func (this *QAbstractProxyModel) Canfetchmore(args ...interface{}) (ret interface{}) {
  // canFetchMore(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractProxyModel12canFetchMoreERK11QModelIndex
    // invoke: bool canFetchMore(const class QModelIndex &)
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QAbstractProxyModel12canFetchMoreERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "canFetchMore", args)
  }

  return
}

// itemData(const class QModelIndex &)
func (this *QAbstractProxyModel) Itemdata(args ...interface{}) () {
  // itemData(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractProxyModel8itemDataERK11QModelIndex
    // invoke: QMap<int, QVariant> itemData(const class QModelIndex &)
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK19QAbstractProxyModel8itemDataERK11QModelIndex(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "itemData", args)
  }

  return
}

// mapSelectionFromSource(const class QItemSelection &)
func (this *QAbstractProxyModel) Mapselectionfromsource(args ...interface{}) (ret interface{}) {
  // mapSelectionFromSource(const class QItemSelection &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QItemSelection{}) // "const QItemSelection &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractProxyModel22mapSelectionFromSourceERK14QItemSelection
    // invoke: QItemSelection mapSelectionFromSource(const class QItemSelection &)
    var arg0 = args[0].(*QItemSelection).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QAbstractProxyModel22mapSelectionFromSourceERK14QItemSelection(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QItemSelection{}) // "QItemSelection"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "mapSelectionFromSource", args)
  }

  return
}

// data(const class QModelIndex &, int)
func (this *QAbstractProxyModel) Data(args ...interface{}) (ret interface{}) {
  // data(const class QModelIndex &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractProxyModel4dataERK11QModelIndexi
    // invoke: QVariant data(const class QModelIndex &, int)
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK19QAbstractProxyModel4dataERK11QModelIndexi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVariant{}) // "QVariant"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "data", args)
  }

  return
}

// mimeTypes()
func (this *QAbstractProxyModel) Mimetypes(args ...interface{}) () {
  // mimeTypes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractProxyModel9mimeTypesEv
    // invoke: QStringList mimeTypes()
    C.C_ZNK19QAbstractProxyModel9mimeTypesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "mimeTypes", args)
  }

  return
}

// metaObject()
func (this *QAbstractProxyModel) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractProxyModel10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK19QAbstractProxyModel10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "metaObject", args)
  }

  return
}

// revert()
func (this *QAbstractProxyModel) Revert(args ...interface{}) () {
  // revert()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractProxyModel6revertEv
    // invoke: void revert()
    C.C_ZN19QAbstractProxyModel6revertEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "revert", args)
  }

  return
}

// ~QAbstractProxyModel()
func (this *QAbstractProxyModel) Freeqabstractproxymodel(args ...interface{}) () {
  // ~QAbstractProxyModel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractProxyModelD0Ev
    // invoke: void ~QAbstractProxyModel()
    C.C_ZN19QAbstractProxyModelD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "~QAbstractProxyModel", args)
  }

  return
}

// flags(const class QModelIndex &)
func (this *QAbstractProxyModel) Flags(args ...interface{}) () {
  // flags(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractProxyModel5flagsERK11QModelIndex
    // invoke: Qt::ItemFlags flags(const class QModelIndex &)
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK19QAbstractProxyModel5flagsERK11QModelIndex(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "flags", args)
  }

  return
}

// sourceModel()
func (this *QAbstractProxyModel) Sourcemodel(args ...interface{}) () {
  // sourceModel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractProxyModel11sourceModelEv
    // invoke: QAbstractItemModel * sourceModel()
    C.C_ZNK19QAbstractProxyModel11sourceModelEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "sourceModel", args)
  }

  return
}

// fetchMore(const class QModelIndex &)
func (this *QAbstractProxyModel) Fetchmore(args ...interface{}) () {
  // fetchMore(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractProxyModel9fetchMoreERK11QModelIndex
    // invoke: void fetchMore(const class QModelIndex &)
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QAbstractProxyModel9fetchMoreERK11QModelIndex(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "fetchMore", args)
  }

  return
}

// supportedDropActions()
func (this *QAbstractProxyModel) Supporteddropactions(args ...interface{}) () {
  // supportedDropActions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractProxyModel20supportedDropActionsEv
    // invoke: Qt::DropActions supportedDropActions()
    C.C_ZNK19QAbstractProxyModel20supportedDropActionsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "supportedDropActions", args)
  }

  return
}

// <= body block end

