package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
  // proto:  QStringList QAbstractProxyModel::mimeTypes();
extern void _ZNK19QAbstractProxyModel9mimeTypesEv(void* qthis);
  // proto:  void QAbstractProxyModel::revert();
extern void _ZN19QAbstractProxyModel6revertEv(void* qthis);
  // proto:  QItemSelection QAbstractProxyModel::mapSelectionToSource(const QItemSelection & selection);
extern void _ZNK19QAbstractProxyModel20mapSelectionToSourceERK14QItemSelection(void* qthis, void* arg0);
  // proto:  QItemSelection QAbstractProxyModel::mapSelectionFromSource(const QItemSelection & selection);
extern void _ZNK19QAbstractProxyModel22mapSelectionFromSourceERK14QItemSelection(void* qthis, void* arg0);
  // proto:  QVariant QAbstractProxyModel::data(const QModelIndex & proxyIndex, int role);
extern void _ZNK19QAbstractProxyModel4dataERK11QModelIndexi(void* qthis, void* arg0, int arg1);
  // proto:  bool QAbstractProxyModel::submit();
extern void _ZN19QAbstractProxyModel6submitEv(void* qthis);
  // proto:  QSize QAbstractProxyModel::span(const QModelIndex & index);
extern void _ZNK19QAbstractProxyModel4spanERK11QModelIndex(void* qthis, void* arg0);
  // proto:  bool QAbstractProxyModel::canFetchMore(const QModelIndex & parent);
extern void _ZNK19QAbstractProxyModel12canFetchMoreERK11QModelIndex(void* qthis, void* arg0);
  // proto:  const QMetaObject * QAbstractProxyModel::metaObject();
extern void _ZNK19QAbstractProxyModel10metaObjectEv(void* qthis);
  // proto:  void QAbstractProxyModel::QAbstractProxyModel(const QAbstractProxyModel & );
extern void* dector_ZN19QAbstractProxyModelC1ERKS_(void* arg0);
extern void _ZN19QAbstractProxyModelC1ERKS_(void* qthis, void* arg0);
  // proto:  QModelIndex QAbstractProxyModel::mapToSource(const QModelIndex & proxyIndex);
extern void _ZNK19QAbstractProxyModel11mapToSourceERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QModelIndex QAbstractProxyModel::mapFromSource(const QModelIndex & sourceIndex);
extern void _ZNK19QAbstractProxyModel13mapFromSourceERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QMap<int, QVariant> QAbstractProxyModel::itemData(const QModelIndex & index);
extern void _ZNK19QAbstractProxyModel8itemDataERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QModelIndex QAbstractProxyModel::buddy(const QModelIndex & index);
extern void _ZNK19QAbstractProxyModel5buddyERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QAbstractProxyModel::setSourceModel(QAbstractItemModel * sourceModel);
extern void _ZN19QAbstractProxyModel14setSourceModelEP18QAbstractItemModel(void* qthis, void* arg0);
  // proto:  QModelIndex QAbstractProxyModel::sibling(int row, int column, const QModelIndex & idx);
extern void _ZNK19QAbstractProxyModel7siblingEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  bool QAbstractProxyModel::setData(const QModelIndex & index, const QVariant & value, int role);
extern void _ZN19QAbstractProxyModel7setDataERK11QModelIndexRK8QVarianti(void* qthis, void* arg0, void* arg1, int arg2);
  // proto:  void QAbstractProxyModel::fetchMore(const QModelIndex & parent);
extern void _ZN19QAbstractProxyModel9fetchMoreERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QAbstractProxyModel::~QAbstractProxyModel();
extern void _ZN19QAbstractProxyModelD0Ev(void* qthis);
  // proto:  void QAbstractProxyModel::QAbstractProxyModel(QObject * parent);
extern void* dector_ZN19QAbstractProxyModelC1EP7QObject(void* arg0);
extern void _ZN19QAbstractProxyModelC1EP7QObject(void* qthis, void* arg0);
  // proto:  bool QAbstractProxyModel::hasChildren(const QModelIndex & parent);
extern void _ZNK19QAbstractProxyModel11hasChildrenERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QAbstractItemModel * QAbstractProxyModel::sourceModel();
extern void _ZNK19QAbstractProxyModel11sourceModelEv(void* qthis);
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
  qclsinst uint64 /* *mut c_void*/;
//  _sourceModelChanged QAbstractProxyModel_sourceModelChanged_signal;
}

  // proto:  QStringList QAbstractProxyModel::mimeTypes();
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
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "mimeTypes", args)
  }

}

  // proto:  void QAbstractProxyModel::revert();
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
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "revert", args)
  }

}

  // proto:  QItemSelection QAbstractProxyModel::mapSelectionToSource(const QItemSelection & selection);
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
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "mapSelectionToSource", args)
  }

}

  // proto:  QItemSelection QAbstractProxyModel::mapSelectionFromSource(const QItemSelection & selection);
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
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "mapSelectionFromSource", args)
  }

}

  // proto:  QVariant QAbstractProxyModel::data(const QModelIndex & proxyIndex, int role);
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
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "data", args)
  }

}

  // proto:  bool QAbstractProxyModel::submit();
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
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "submit", args)
  }

}

  // proto:  QSize QAbstractProxyModel::span(const QModelIndex & index);
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
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "span", args)
  }

}

  // proto:  bool QAbstractProxyModel::canFetchMore(const QModelIndex & parent);
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
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "canFetchMore", args)
  }

}

  // proto:  const QMetaObject * QAbstractProxyModel::metaObject();
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
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "metaObject", args)
  }

}

  // proto:  void QAbstractProxyModel::QAbstractProxyModel(const QAbstractProxyModel & );
func NewQAbstractProxyModel(args ...interface{}) QAbstractProxyModel {
  return QAbstractProxyModel{}
}

  // proto:  QModelIndex QAbstractProxyModel::mapToSource(const QModelIndex & proxyIndex);
func (this *QAbstractProxyModel) mapToSource(args ...interface{}) () {
  // mapToSource(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractProxyModel11mapToSourceERK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "mapToSource", args)
  }

}

  // proto:  QModelIndex QAbstractProxyModel::mapFromSource(const QModelIndex & sourceIndex);
func (this *QAbstractProxyModel) mapFromSource(args ...interface{}) () {
  // mapFromSource(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractProxyModel13mapFromSourceERK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "mapFromSource", args)
  }

}

  // proto:  QMap<int, QVariant> QAbstractProxyModel::itemData(const QModelIndex & index);
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
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "itemData", args)
  }

}

  // proto:  QModelIndex QAbstractProxyModel::buddy(const QModelIndex & index);
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
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "buddy", args)
  }

}

  // proto:  void QAbstractProxyModel::setSourceModel(QAbstractItemModel * sourceModel);
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
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "setSourceModel", args)
  }

}

  // proto:  QModelIndex QAbstractProxyModel::sibling(int row, int column, const QModelIndex & idx);
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
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "sibling", args)
  }

}

  // proto:  bool QAbstractProxyModel::setData(const QModelIndex & index, const QVariant & value, int role);
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
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "setData", args)
  }

}

  // proto:  void QAbstractProxyModel::fetchMore(const QModelIndex & parent);
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
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "fetchMore", args)
  }

}

  // proto:  void QAbstractProxyModel::~QAbstractProxyModel();
func (this *QAbstractProxyModel) FreeQAbstractProxyModel(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "~QAbstractProxyModel", args)
  }

}

  // proto:  bool QAbstractProxyModel::hasChildren(const QModelIndex & parent);
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
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "hasChildren", args)
  }

}

  // proto:  QAbstractItemModel * QAbstractProxyModel::sourceModel();
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
  default:
    qtrt.ErrorResolve("QAbstractProxyModel", "sourceModel", args)
  }

}

// <= body block end

