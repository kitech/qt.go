package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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
extern void _ZNK19QAbstractProxyModel20mapSelectionToSourceERK14QItemSelection(void* qthis, void* arg0); // 4
  // proto:  Qt::DropActions QAbstractProxyModel::supportedDragActions();
extern void _ZNK19QAbstractProxyModel20supportedDragActionsEv(void* qthis); // 4
  // proto:  void QAbstractProxyModel::setSourceModel(QAbstractItemModel * sourceModel);
extern void _ZN19QAbstractProxyModel14setSourceModelEP18QAbstractItemModel(void* qthis, void* arg0); // 4
  // proto:  QModelIndex QAbstractProxyModel::sibling(int row, int column, const QModelIndex & idx);
extern void _ZNK19QAbstractProxyModel7siblingEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QAbstractProxyModel::QAbstractProxyModel(QObject * parent);
extern void _ZN19QAbstractProxyModelC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  QSize QAbstractProxyModel::span(const QModelIndex & index);
extern void _ZNK19QAbstractProxyModel4spanERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QModelIndex QAbstractProxyModel::buddy(const QModelIndex & index);
extern void _ZNK19QAbstractProxyModel5buddyERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QAbstractProxyModel::hasChildren(const QModelIndex & parent);
extern void _ZNK19QAbstractProxyModel11hasChildrenERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QAbstractProxyModel::submit();
extern void _ZN19QAbstractProxyModel6submitEv(void* qthis); // 4
  // proto:  bool QAbstractProxyModel::setData(const QModelIndex & index, const QVariant & value, int role);
extern void _ZN19QAbstractProxyModel7setDataERK11QModelIndexRK8QVarianti(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  bool QAbstractProxyModel::canFetchMore(const QModelIndex & parent);
extern void _ZNK19QAbstractProxyModel12canFetchMoreERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QMap<int, QVariant> QAbstractProxyModel::itemData(const QModelIndex & index);
extern void _ZNK19QAbstractProxyModel8itemDataERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QItemSelection QAbstractProxyModel::mapSelectionFromSource(const QItemSelection & selection);
extern void _ZNK19QAbstractProxyModel22mapSelectionFromSourceERK14QItemSelection(void* qthis, void* arg0); // 4
  // proto:  QVariant QAbstractProxyModel::data(const QModelIndex & proxyIndex, int role);
extern void _ZNK19QAbstractProxyModel4dataERK11QModelIndexi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QStringList QAbstractProxyModel::mimeTypes();
extern void _ZNK19QAbstractProxyModel9mimeTypesEv(void* qthis); // 4
  // proto:  const QMetaObject * QAbstractProxyModel::metaObject();
extern void _ZNK19QAbstractProxyModel10metaObjectEv(void* qthis); // 4
  // proto:  void QAbstractProxyModel::revert();
extern void _ZN19QAbstractProxyModel6revertEv(void* qthis); // 4
  // proto:  void QAbstractProxyModel::~QAbstractProxyModel();
extern void _ZN19QAbstractProxyModelD2Ev(void* qthis); // 4
  // proto:  Qt::ItemFlags QAbstractProxyModel::flags(const QModelIndex & index);
extern void _ZNK19QAbstractProxyModel5flagsERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QAbstractItemModel * QAbstractProxyModel::sourceModel();
extern void _ZNK19QAbstractProxyModel11sourceModelEv(void* qthis); // 4
  // proto:  void QAbstractProxyModel::fetchMore(const QModelIndex & parent);
extern void _ZN19QAbstractProxyModel9fetchMoreERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  Qt::DropActions QAbstractProxyModel::supportedDropActions();
extern void _ZNK19QAbstractProxyModel20supportedDropActionsEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
//  _sourceModelChanged QAbstractProxyModel_sourceModelChanged_signal;
}

// mapSelectionToSource(const class QItemSelection &)
func (this *QAbstractProxyModel) mapSelectionToSource(args ...interface{}) () {
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
    var arg0 = args[0].(QItemSelection).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK19QAbstractProxyModel20mapSelectionToSourceERK14QItemSelection(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "mapSelectionToSource", args)
  }

}

// supportedDragActions()
func (this *QAbstractProxyModel) supportedDragActions(args ...interface{}) () {
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
    C._ZNK19QAbstractProxyModel20supportedDragActionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "supportedDragActions", args)
  }

}

// setSourceModel(class QAbstractItemModel *)
func (this *QAbstractProxyModel) setSourceModel(args ...interface{}) () {
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
    var arg0 = args[0].(QAbstractItemModel).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN19QAbstractProxyModel14setSourceModelEP18QAbstractItemModel(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "setSourceModel", args)
  }

}

// sibling(int, int, const class QModelIndex &)
func (this *QAbstractProxyModel) sibling(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C._ZNK19QAbstractProxyModel7siblingEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "sibling", args)
  }

}

// QAbstractProxyModel(class QObject *)
func NewQAbstractProxyModel(args ...interface{}) QAbstractProxyModel {
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
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN19QAbstractProxyModelC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "QAbstractProxyModel", args)
  }

  return QAbstractProxyModel{}
}

// span(const class QModelIndex &)
func (this *QAbstractProxyModel) span(args ...interface{}) () {
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK19QAbstractProxyModel4spanERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "span", args)
  }

}

// buddy(const class QModelIndex &)
func (this *QAbstractProxyModel) buddy(args ...interface{}) () {
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK19QAbstractProxyModel5buddyERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "buddy", args)
  }

}

// hasChildren(const class QModelIndex &)
func (this *QAbstractProxyModel) hasChildren(args ...interface{}) () {
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK19QAbstractProxyModel11hasChildrenERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "hasChildren", args)
  }

}

// submit()
func (this *QAbstractProxyModel) submit(args ...interface{}) () {
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
    C._ZN19QAbstractProxyModel6submitEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "submit", args)
  }

}

// setData(const class QModelIndex &, const class QVariant &, int)
func (this *QAbstractProxyModel) setData(args ...interface{}) () {
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN19QAbstractProxyModel7setDataERK11QModelIndexRK8QVarianti(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "setData", args)
  }

}

// canFetchMore(const class QModelIndex &)
func (this *QAbstractProxyModel) canFetchMore(args ...interface{}) () {
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK19QAbstractProxyModel12canFetchMoreERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "canFetchMore", args)
  }

}

// itemData(const class QModelIndex &)
func (this *QAbstractProxyModel) itemData(args ...interface{}) () {
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK19QAbstractProxyModel8itemDataERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "itemData", args)
  }

}

// mapSelectionFromSource(const class QItemSelection &)
func (this *QAbstractProxyModel) mapSelectionFromSource(args ...interface{}) () {
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
    var arg0 = args[0].(QItemSelection).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK19QAbstractProxyModel22mapSelectionFromSourceERK14QItemSelection(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "mapSelectionFromSource", args)
  }

}

// data(const class QModelIndex &, int)
func (this *QAbstractProxyModel) data(args ...interface{}) () {
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK19QAbstractProxyModel4dataERK11QModelIndexi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "data", args)
  }

}

// mimeTypes()
func (this *QAbstractProxyModel) mimeTypes(args ...interface{}) () {
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
    C._ZNK19QAbstractProxyModel9mimeTypesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "mimeTypes", args)
  }

}

// metaObject()
func (this *QAbstractProxyModel) metaObject(args ...interface{}) () {
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
    C._ZNK19QAbstractProxyModel10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "metaObject", args)
  }

}

// revert()
func (this *QAbstractProxyModel) revert(args ...interface{}) () {
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
    C._ZN19QAbstractProxyModel6revertEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "revert", args)
  }

}

// ~QAbstractProxyModel()
func (this *QAbstractProxyModel) FreeQAbstractProxyModel(args ...interface{}) () {
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
    C._ZN19QAbstractProxyModelD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "~QAbstractProxyModel", args)
  }

}

// flags(const class QModelIndex &)
func (this *QAbstractProxyModel) flags(args ...interface{}) () {
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK19QAbstractProxyModel5flagsERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "flags", args)
  }

}

// sourceModel()
func (this *QAbstractProxyModel) sourceModel(args ...interface{}) () {
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
    C._ZNK19QAbstractProxyModel11sourceModelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "sourceModel", args)
  }

}

// fetchMore(const class QModelIndex &)
func (this *QAbstractProxyModel) fetchMore(args ...interface{}) () {
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN19QAbstractProxyModel9fetchMoreERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "fetchMore", args)
  }

}

// supportedDropActions()
func (this *QAbstractProxyModel) supportedDropActions(args ...interface{}) () {
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
    C._ZNK19QAbstractProxyModel20supportedDropActionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "supportedDropActions", args)
  }

}

// <= body block end

