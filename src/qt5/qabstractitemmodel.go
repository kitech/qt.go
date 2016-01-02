package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtCore/qabstractitemmodel.h
// dst-file: /src/core/qabstractitemmodel.go
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
  // proto:  void QModelIndex::QModelIndex(int arow, int acolumn, void * ptr, const QAbstractItemModel * amodel);
extern void* dector_ZN11QModelIndexC1EiiPvPK18QAbstractItemModel(int arg0, int arg1, void* arg2, void* arg3);
extern void demth_ZN11QModelIndexC1EiiPvPK18QAbstractItemModel(void* qthis, int arg0, int arg1, void* arg2, void* arg3);
  // proto:  void QModelIndex::QModelIndex();
extern void* dector_ZN11QModelIndexC1Ev();
extern void _ZN11QModelIndexC1Ev(void* qthis);
  // proto:  int QModelIndex::column();
extern void _ZNK11QModelIndex6columnEv(void* qthis);
  // proto:  quintptr QModelIndex::internalId();
extern void _ZNK11QModelIndex10internalIdEv(void* qthis);
  // proto:  QModelIndex QModelIndex::child(int row, int column);
extern void demth_ZNK11QModelIndex5childEii(void* qthis, int arg0, int arg1);
  // proto:  void * QModelIndex::internalPointer();
extern void demth_ZNK11QModelIndex15internalPointerEv(void* qthis);
  // proto:  void QModelIndex::QModelIndex(int arow, int acolumn, quintptr id, const QAbstractItemModel * amodel);
extern void* dector_ZN11QModelIndexC1EiiiPK18QAbstractItemModel(int arg0, int arg1, int arg2, void* arg3);
extern void _ZN11QModelIndexC1EiiiPK18QAbstractItemModel(void* qthis, int arg0, int arg1, int arg2, void* arg3);
  // proto:  bool QModelIndex::isValid();
extern void _ZNK11QModelIndex7isValidEv(void* qthis);
  // proto:  QModelIndex QModelIndex::parent();
extern void demth_ZNK11QModelIndex6parentEv(void* qthis);
  // proto:  QModelIndex QModelIndex::sibling(int row, int column);
extern void demth_ZNK11QModelIndex7siblingEii(void* qthis, int arg0, int arg1);
  // proto:  const QAbstractItemModel * QModelIndex::model();
extern void _ZNK11QModelIndex5modelEv(void* qthis);
  // proto:  QVariant QModelIndex::data(int role);
extern void demth_ZNK11QModelIndex4dataEi(void* qthis, int arg0);
  // proto:  int QModelIndex::row();
extern void _ZNK11QModelIndex3rowEv(void* qthis);
  // proto:  QModelIndex QPersistentModelIndex::sibling(int row, int column);
extern void _ZNK21QPersistentModelIndex7siblingEii(void* qthis, int arg0, int arg1);
  // proto:  QVariant QPersistentModelIndex::data(int role);
extern void _ZNK21QPersistentModelIndex4dataEi(void* qthis, int arg0);
  // proto:  QModelIndex QPersistentModelIndex::parent();
extern void _ZNK21QPersistentModelIndex6parentEv(void* qthis);
  // proto:  void QPersistentModelIndex::QPersistentModelIndex(const QPersistentModelIndex & other);
extern void* dector_ZN21QPersistentModelIndexC1ERKS_(void* arg0);
extern void _ZN21QPersistentModelIndexC1ERKS_(void* qthis, void* arg0);
  // proto:  void * QPersistentModelIndex::internalPointer();
extern void _ZNK21QPersistentModelIndex15internalPointerEv(void* qthis);
  // proto:  int QPersistentModelIndex::row();
extern void _ZNK21QPersistentModelIndex3rowEv(void* qthis);
  // proto:  quintptr QPersistentModelIndex::internalId();
extern void _ZNK21QPersistentModelIndex10internalIdEv(void* qthis);
  // proto:  const QAbstractItemModel * QPersistentModelIndex::model();
extern void _ZNK21QPersistentModelIndex5modelEv(void* qthis);
  // proto:  void QPersistentModelIndex::~QPersistentModelIndex();
extern void _ZN21QPersistentModelIndexD0Ev(void* qthis);
  // proto:  void QPersistentModelIndex::QPersistentModelIndex(const QModelIndex & index);
extern void* dector_ZN21QPersistentModelIndexC1ERK11QModelIndex(void* arg0);
extern void _ZN21QPersistentModelIndexC1ERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QPersistentModelIndex::QPersistentModelIndex();
extern void* dector_ZN21QPersistentModelIndexC1Ev();
extern void _ZN21QPersistentModelIndexC1Ev(void* qthis);
  // proto:  int QPersistentModelIndex::column();
extern void _ZNK21QPersistentModelIndex6columnEv(void* qthis);
  // proto:  void QPersistentModelIndex::swap(QPersistentModelIndex & other);
extern void demth_ZN21QPersistentModelIndex4swapERS_(void* qthis, void* arg0);
  // proto:  QModelIndex QPersistentModelIndex::child(int row, int column);
extern void _ZNK21QPersistentModelIndex5childEii(void* qthis, int arg0, int arg1);
  // proto:  bool QPersistentModelIndex::isValid();
extern void _ZNK21QPersistentModelIndex7isValidEv(void* qthis);
  // proto:  void QAbstractTableModel::QAbstractTableModel(const QAbstractTableModel & );
extern void* dector_ZN19QAbstractTableModelC1ERKS_(void* arg0);
extern void _ZN19QAbstractTableModelC1ERKS_(void* qthis, void* arg0);
  // proto:  void QAbstractTableModel::QAbstractTableModel(QObject * parent);
extern void* dector_ZN19QAbstractTableModelC1EP7QObject(void* arg0);
extern void _ZN19QAbstractTableModelC1EP7QObject(void* qthis, void* arg0);
  // proto:  QModelIndex QAbstractTableModel::index(int row, int column, const QModelIndex & parent);
extern void _ZNK19QAbstractTableModel5indexEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  QModelIndex QAbstractTableModel::sibling(int row, int column, const QModelIndex & idx);
extern void _ZNK19QAbstractTableModel7siblingEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  void QAbstractTableModel::~QAbstractTableModel();
extern void _ZN19QAbstractTableModelD0Ev(void* qthis);
  // proto:  const QMetaObject * QAbstractTableModel::metaObject();
extern void _ZNK19QAbstractTableModel10metaObjectEv(void* qthis);
  // proto:  bool QAbstractItemModel::removeColumns(int column, int count, const QModelIndex & parent);
extern void _ZN18QAbstractItemModel13removeColumnsEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  void QAbstractItemModel::~QAbstractItemModel();
extern void _ZN18QAbstractItemModelD0Ev(void* qthis);
  // proto:  bool QAbstractItemModel::canFetchMore(const QModelIndex & parent);
extern void _ZNK18QAbstractItemModel12canFetchMoreERK11QModelIndex(void* qthis, void* arg0);
  // proto:  bool QAbstractItemModel::submit();
extern void _ZN18QAbstractItemModel6submitEv(void* qthis);
  // proto:  bool QAbstractItemModel::hasIndex(int row, int column, const QModelIndex & parent);
extern void _ZNK18QAbstractItemModel8hasIndexEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  QHash<int, QByteArray> QAbstractItemModel::roleNames();
extern void _ZNK18QAbstractItemModel9roleNamesEv(void* qthis);
  // proto:  bool QAbstractItemModel::moveColumn(const QModelIndex & sourceParent, int sourceColumn, const QModelIndex & destinationParent, int destinationChild);
extern void demth_ZN18QAbstractItemModel10moveColumnERK11QModelIndexiS2_i(void* qthis, void* arg0, int arg1, void* arg2, int arg3);
  // proto:  void QAbstractItemModel::fetchMore(const QModelIndex & parent);
extern void _ZN18QAbstractItemModel9fetchMoreERK11QModelIndex(void* qthis, void* arg0);
  // proto:  bool QAbstractItemModel::insertRows(int row, int count, const QModelIndex & parent);
extern void _ZN18QAbstractItemModel10insertRowsEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  QSize QAbstractItemModel::span(const QModelIndex & index);
extern void _ZNK18QAbstractItemModel4spanERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QAbstractItemModel::QAbstractItemModel(QObject * parent);
extern void* dector_ZN18QAbstractItemModelC1EP7QObject(void* arg0);
extern void _ZN18QAbstractItemModelC1EP7QObject(void* qthis, void* arg0);
  // proto:  QModelIndex QAbstractItemModel::index(int row, int column, const QModelIndex & parent);
extern void _ZNK18QAbstractItemModel5indexEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  bool QAbstractItemModel::insertRow(int row, const QModelIndex & parent);
extern void demth_ZN18QAbstractItemModel9insertRowEiRK11QModelIndex(void* qthis, int arg0, void* arg1);
  // proto:  void QAbstractItemModel::QAbstractItemModel(const QAbstractItemModel & );
extern void* dector_ZN18QAbstractItemModelC1ERKS_(void* arg0);
extern void _ZN18QAbstractItemModelC1ERKS_(void* qthis, void* arg0);
  // proto:  const QMetaObject * QAbstractItemModel::metaObject();
extern void _ZNK18QAbstractItemModel10metaObjectEv(void* qthis);
  // proto:  bool QAbstractItemModel::removeRow(int row, const QModelIndex & parent);
extern void demth_ZN18QAbstractItemModel9removeRowEiRK11QModelIndex(void* qthis, int arg0, void* arg1);
  // proto:  bool QAbstractItemModel::setData(const QModelIndex & index, const QVariant & value, int role);
extern void _ZN18QAbstractItemModel7setDataERK11QModelIndexRK8QVarianti(void* qthis, void* arg0, void* arg1, int arg2);
  // proto:  int QAbstractItemModel::rowCount(const QModelIndex & parent);
extern void _ZNK18QAbstractItemModel8rowCountERK11QModelIndex(void* qthis, void* arg0);
  // proto:  bool QAbstractItemModel::removeRows(int row, int count, const QModelIndex & parent);
extern void _ZN18QAbstractItemModel10removeRowsEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  bool QAbstractItemModel::hasChildren(const QModelIndex & parent);
extern void _ZNK18QAbstractItemModel11hasChildrenERK11QModelIndex(void* qthis, void* arg0);
  // proto:  bool QAbstractItemModel::moveRow(const QModelIndex & sourceParent, int sourceRow, const QModelIndex & destinationParent, int destinationChild);
extern void demth_ZN18QAbstractItemModel7moveRowERK11QModelIndexiS2_i(void* qthis, void* arg0, int arg1, void* arg2, int arg3);
  // proto:  void QAbstractItemModel::revert();
extern void _ZN18QAbstractItemModel6revertEv(void* qthis);
  // proto:  bool QAbstractItemModel::removeColumn(int column, const QModelIndex & parent);
extern void demth_ZN18QAbstractItemModel12removeColumnEiRK11QModelIndex(void* qthis, int arg0, void* arg1);
  // proto:  bool QAbstractItemModel::insertColumns(int column, int count, const QModelIndex & parent);
extern void _ZN18QAbstractItemModel13insertColumnsEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  bool QAbstractItemModel::insertColumn(int column, const QModelIndex & parent);
extern void demth_ZN18QAbstractItemModel12insertColumnEiRK11QModelIndex(void* qthis, int arg0, void* arg1);
  // proto:  bool QAbstractItemModel::moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild);
extern void _ZN18QAbstractItemModel11moveColumnsERK11QModelIndexiiS2_i(void* qthis, void* arg0, int arg1, int arg2, void* arg3, int arg4);
  // proto:  QMap<int, QVariant> QAbstractItemModel::itemData(const QModelIndex & index);
extern void _ZNK18QAbstractItemModel8itemDataERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QStringList QAbstractItemModel::mimeTypes();
extern void _ZNK18QAbstractItemModel9mimeTypesEv(void* qthis);
  // proto:  QModelIndex QAbstractItemModel::parent(const QModelIndex & child);
extern void _ZNK18QAbstractItemModel6parentERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QModelIndex QAbstractItemModel::buddy(const QModelIndex & index);
extern void _ZNK18QAbstractItemModel5buddyERK11QModelIndex(void* qthis, void* arg0);
  // proto:  int QAbstractItemModel::columnCount(const QModelIndex & parent);
extern void _ZNK18QAbstractItemModel11columnCountERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QVariant QAbstractItemModel::data(const QModelIndex & index, int role);
extern void _ZNK18QAbstractItemModel4dataERK11QModelIndexi(void* qthis, void* arg0, int arg1);
  // proto:  QModelIndex QAbstractItemModel::sibling(int row, int column, const QModelIndex & idx);
extern void _ZNK18QAbstractItemModel7siblingEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  bool QAbstractItemModel::moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild);
extern void _ZN18QAbstractItemModel8moveRowsERK11QModelIndexiiS2_i(void* qthis, void* arg0, int arg1, int arg2, void* arg3, int arg4);
  // proto:  QModelIndex QAbstractListModel::sibling(int row, int column, const QModelIndex & idx);
extern void _ZNK18QAbstractListModel7siblingEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  void QAbstractListModel::QAbstractListModel(QObject * parent);
extern void* dector_ZN18QAbstractListModelC1EP7QObject(void* arg0);
extern void _ZN18QAbstractListModelC1EP7QObject(void* qthis, void* arg0);
  // proto:  const QMetaObject * QAbstractListModel::metaObject();
extern void _ZNK18QAbstractListModel10metaObjectEv(void* qthis);
  // proto:  void QAbstractListModel::QAbstractListModel(const QAbstractListModel & );
extern void* dector_ZN18QAbstractListModelC1ERKS_(void* arg0);
extern void _ZN18QAbstractListModelC1ERKS_(void* qthis, void* arg0);
  // proto:  QModelIndex QAbstractListModel::index(int row, int column, const QModelIndex & parent);
extern void _ZNK18QAbstractListModel5indexEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  void QAbstractListModel::~QAbstractListModel();
extern void _ZN18QAbstractListModelD0Ev(void* qthis);
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

// class sizeof(QModelIndex)=24
type QModelIndex struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QPersistentModelIndex)=8
type QPersistentModelIndex struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAbstractTableModel)=1
type QAbstractTableModel struct {
  /*qbase*/ QAbstractItemModel;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAbstractItemModel)=1
type QAbstractItemModel struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _modelReset QAbstractItemModel_modelReset_signal;
//  _columnsAboutToBeMoved QAbstractItemModel_columnsAboutToBeMoved_signal;
//  _columnsInserted QAbstractItemModel_columnsInserted_signal;
//  _rowsAboutToBeRemoved QAbstractItemModel_rowsAboutToBeRemoved_signal;
//  _columnsAboutToBeRemoved QAbstractItemModel_columnsAboutToBeRemoved_signal;
//  _rowsAboutToBeMoved QAbstractItemModel_rowsAboutToBeMoved_signal;
//  _layoutChanged QAbstractItemModel_layoutChanged_signal;
//  _columnsRemoved QAbstractItemModel_columnsRemoved_signal;
//  _rowsInserted QAbstractItemModel_rowsInserted_signal;
//  _columnsAboutToBeInserted QAbstractItemModel_columnsAboutToBeInserted_signal;
//  _layoutAboutToBeChanged QAbstractItemModel_layoutAboutToBeChanged_signal;
//  _rowsRemoved QAbstractItemModel_rowsRemoved_signal;
//  _rowsMoved QAbstractItemModel_rowsMoved_signal;
//  _headerDataChanged QAbstractItemModel_headerDataChanged_signal;
//  _columnsMoved QAbstractItemModel_columnsMoved_signal;
//  _rowsAboutToBeInserted QAbstractItemModel_rowsAboutToBeInserted_signal;
//  _modelAboutToBeReset QAbstractItemModel_modelAboutToBeReset_signal;
//  _dataChanged QAbstractItemModel_dataChanged_signal;
}

// class sizeof(QAbstractListModel)=1
type QAbstractListModel struct {
  /*qbase*/ QAbstractItemModel;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QModelIndex::QModelIndex(int arow, int acolumn, void * ptr, const QAbstractItemModel * amodel);
func NewQModelIndex(args ...interface{}) QModelIndex {
  return QModelIndex{}
}

  // proto:  int QModelIndex::column();
func (this *QModelIndex) column(args ...interface{}) () {
  // column()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QModelIndex6columnEv
  default:
    qtrt.ErrorResolve("QModelIndex", "column", args)
  }

}

  // proto:  quintptr QModelIndex::internalId();
func (this *QModelIndex) internalId(args ...interface{}) () {
  // internalId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QModelIndex10internalIdEv
  default:
    qtrt.ErrorResolve("QModelIndex", "internalId", args)
  }

}

  // proto:  QModelIndex QModelIndex::child(int row, int column);
func (this *QModelIndex) child(args ...interface{}) () {
  // child(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QModelIndex5childEii
  default:
    qtrt.ErrorResolve("QModelIndex", "child", args)
  }

}

  // proto:  void * QModelIndex::internalPointer();
func (this *QModelIndex) internalPointer(args ...interface{}) () {
  // internalPointer()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QModelIndex15internalPointerEv
  default:
    qtrt.ErrorResolve("QModelIndex", "internalPointer", args)
  }

}

  // proto:  bool QModelIndex::isValid();
func (this *QModelIndex) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QModelIndex7isValidEv
  default:
    qtrt.ErrorResolve("QModelIndex", "isValid", args)
  }

}

  // proto:  QModelIndex QModelIndex::parent();
func (this *QModelIndex) parent(args ...interface{}) () {
  // parent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QModelIndex6parentEv
  default:
    qtrt.ErrorResolve("QModelIndex", "parent", args)
  }

}

  // proto:  QModelIndex QModelIndex::sibling(int row, int column);
func (this *QModelIndex) sibling(args ...interface{}) () {
  // sibling(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QModelIndex7siblingEii
  default:
    qtrt.ErrorResolve("QModelIndex", "sibling", args)
  }

}

  // proto:  const QAbstractItemModel * QModelIndex::model();
func (this *QModelIndex) model(args ...interface{}) () {
  // model()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QModelIndex5modelEv
  default:
    qtrt.ErrorResolve("QModelIndex", "model", args)
  }

}

  // proto:  QVariant QModelIndex::data(int role);
func (this *QModelIndex) data(args ...interface{}) () {
  // data(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QModelIndex4dataEi
  default:
    qtrt.ErrorResolve("QModelIndex", "data", args)
  }

}

  // proto:  int QModelIndex::row();
func (this *QModelIndex) row(args ...interface{}) () {
  // row()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QModelIndex3rowEv
  default:
    qtrt.ErrorResolve("QModelIndex", "row", args)
  }

}

  // proto:  QModelIndex QPersistentModelIndex::sibling(int row, int column);
func (this *QPersistentModelIndex) sibling(args ...interface{}) () {
  // sibling(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QPersistentModelIndex7siblingEii
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "sibling", args)
  }

}

  // proto:  QVariant QPersistentModelIndex::data(int role);
func (this *QPersistentModelIndex) data(args ...interface{}) () {
  // data(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QPersistentModelIndex4dataEi
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "data", args)
  }

}

  // proto:  QModelIndex QPersistentModelIndex::parent();
func (this *QPersistentModelIndex) parent(args ...interface{}) () {
  // parent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QPersistentModelIndex6parentEv
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "parent", args)
  }

}

  // proto:  void QPersistentModelIndex::QPersistentModelIndex(const QPersistentModelIndex & other);
func NewQPersistentModelIndex(args ...interface{}) QPersistentModelIndex {
  return QPersistentModelIndex{}
}

  // proto:  void * QPersistentModelIndex::internalPointer();
func (this *QPersistentModelIndex) internalPointer(args ...interface{}) () {
  // internalPointer()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QPersistentModelIndex15internalPointerEv
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "internalPointer", args)
  }

}

  // proto:  int QPersistentModelIndex::row();
func (this *QPersistentModelIndex) row(args ...interface{}) () {
  // row()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QPersistentModelIndex3rowEv
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "row", args)
  }

}

  // proto:  quintptr QPersistentModelIndex::internalId();
func (this *QPersistentModelIndex) internalId(args ...interface{}) () {
  // internalId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QPersistentModelIndex10internalIdEv
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "internalId", args)
  }

}

  // proto:  const QAbstractItemModel * QPersistentModelIndex::model();
func (this *QPersistentModelIndex) model(args ...interface{}) () {
  // model()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QPersistentModelIndex5modelEv
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "model", args)
  }

}

  // proto:  void QPersistentModelIndex::~QPersistentModelIndex();
func (this *QPersistentModelIndex) FreeQPersistentModelIndex(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "~QPersistentModelIndex", args)
  }

}

  // proto:  int QPersistentModelIndex::column();
func (this *QPersistentModelIndex) column(args ...interface{}) () {
  // column()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QPersistentModelIndex6columnEv
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "column", args)
  }

}

  // proto:  void QPersistentModelIndex::swap(QPersistentModelIndex & other);
func (this *QPersistentModelIndex) swap(args ...interface{}) () {
  // swap(class QPersistentModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPersistentModelIndex{}) // "QPersistentModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QPersistentModelIndex4swapERS_
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "swap", args)
  }

}

  // proto:  QModelIndex QPersistentModelIndex::child(int row, int column);
func (this *QPersistentModelIndex) child(args ...interface{}) () {
  // child(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QPersistentModelIndex5childEii
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "child", args)
  }

}

  // proto:  bool QPersistentModelIndex::isValid();
func (this *QPersistentModelIndex) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QPersistentModelIndex7isValidEv
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "isValid", args)
  }

}

  // proto:  void QAbstractTableModel::QAbstractTableModel(const QAbstractTableModel & );
func NewQAbstractTableModel(args ...interface{}) QAbstractTableModel {
  return QAbstractTableModel{}
}

  // proto:  QModelIndex QAbstractTableModel::index(int row, int column, const QModelIndex & parent);
func (this *QAbstractTableModel) index(args ...interface{}) () {
  // index(int, int, const class QModelIndex &)
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
    // invoke: _ZNK19QAbstractTableModel5indexEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractTableModel", "index", args)
  }

}

  // proto:  QModelIndex QAbstractTableModel::sibling(int row, int column, const QModelIndex & idx);
func (this *QAbstractTableModel) sibling(args ...interface{}) () {
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
    // invoke: _ZNK19QAbstractTableModel7siblingEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractTableModel", "sibling", args)
  }

}

  // proto:  void QAbstractTableModel::~QAbstractTableModel();
func (this *QAbstractTableModel) FreeQAbstractTableModel(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractTableModel", "~QAbstractTableModel", args)
  }

}

  // proto:  const QMetaObject * QAbstractTableModel::metaObject();
func (this *QAbstractTableModel) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractTableModel10metaObjectEv
  default:
    qtrt.ErrorResolve("QAbstractTableModel", "metaObject", args)
  }

}

  // proto:  bool QAbstractItemModel::removeColumns(int column, int count, const QModelIndex & parent);
func (this *QAbstractItemModel) removeColumns(args ...interface{}) () {
  // removeColumns(int, int, const class QModelIndex &)
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
    // invoke: _ZN18QAbstractItemModel13removeColumnsEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "removeColumns", args)
  }

}

  // proto:  void QAbstractItemModel::~QAbstractItemModel();
func (this *QAbstractItemModel) FreeQAbstractItemModel(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "~QAbstractItemModel", args)
  }

}

  // proto:  bool QAbstractItemModel::canFetchMore(const QModelIndex & parent);
func (this *QAbstractItemModel) canFetchMore(args ...interface{}) () {
  // canFetchMore(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel12canFetchMoreERK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "canFetchMore", args)
  }

}

  // proto:  bool QAbstractItemModel::submit();
func (this *QAbstractItemModel) submit(args ...interface{}) () {
  // submit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModel6submitEv
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "submit", args)
  }

}

  // proto:  bool QAbstractItemModel::hasIndex(int row, int column, const QModelIndex & parent);
func (this *QAbstractItemModel) hasIndex(args ...interface{}) () {
  // hasIndex(int, int, const class QModelIndex &)
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
    // invoke: _ZNK18QAbstractItemModel8hasIndexEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "hasIndex", args)
  }

}

  // proto:  QHash<int, QByteArray> QAbstractItemModel::roleNames();
func (this *QAbstractItemModel) roleNames(args ...interface{}) () {
  // roleNames()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel9roleNamesEv
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "roleNames", args)
  }

}

  // proto:  bool QAbstractItemModel::moveColumn(const QModelIndex & sourceParent, int sourceColumn, const QModelIndex & destinationParent, int destinationChild);
func (this *QAbstractItemModel) moveColumn(args ...interface{}) () {
  // moveColumn(const class QModelIndex &, int, const class QModelIndex &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModel10moveColumnERK11QModelIndexiS2_i
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "moveColumn", args)
  }

}

  // proto:  void QAbstractItemModel::fetchMore(const QModelIndex & parent);
func (this *QAbstractItemModel) fetchMore(args ...interface{}) () {
  // fetchMore(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModel9fetchMoreERK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "fetchMore", args)
  }

}

  // proto:  bool QAbstractItemModel::insertRows(int row, int count, const QModelIndex & parent);
func (this *QAbstractItemModel) insertRows(args ...interface{}) () {
  // insertRows(int, int, const class QModelIndex &)
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
    // invoke: _ZN18QAbstractItemModel10insertRowsEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "insertRows", args)
  }

}

  // proto:  QSize QAbstractItemModel::span(const QModelIndex & index);
func (this *QAbstractItemModel) span(args ...interface{}) () {
  // span(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel4spanERK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "span", args)
  }

}

  // proto:  void QAbstractItemModel::QAbstractItemModel(QObject * parent);
func NewQAbstractItemModel(args ...interface{}) QAbstractItemModel {
  return QAbstractItemModel{}
}

  // proto:  QModelIndex QAbstractItemModel::index(int row, int column, const QModelIndex & parent);
func (this *QAbstractItemModel) index(args ...interface{}) () {
  // index(int, int, const class QModelIndex &)
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
    // invoke: _ZNK18QAbstractItemModel5indexEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "index", args)
  }

}

  // proto:  bool QAbstractItemModel::insertRow(int row, const QModelIndex & parent);
func (this *QAbstractItemModel) insertRow(args ...interface{}) () {
  // insertRow(int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModel9insertRowEiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "insertRow", args)
  }

}

  // proto:  const QMetaObject * QAbstractItemModel::metaObject();
func (this *QAbstractItemModel) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel10metaObjectEv
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "metaObject", args)
  }

}

  // proto:  bool QAbstractItemModel::removeRow(int row, const QModelIndex & parent);
func (this *QAbstractItemModel) removeRow(args ...interface{}) () {
  // removeRow(int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModel9removeRowEiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "removeRow", args)
  }

}

  // proto:  bool QAbstractItemModel::setData(const QModelIndex & index, const QVariant & value, int role);
func (this *QAbstractItemModel) setData(args ...interface{}) () {
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
    // invoke: _ZN18QAbstractItemModel7setDataERK11QModelIndexRK8QVarianti
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "setData", args)
  }

}

  // proto:  int QAbstractItemModel::rowCount(const QModelIndex & parent);
func (this *QAbstractItemModel) rowCount(args ...interface{}) () {
  // rowCount(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel8rowCountERK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "rowCount", args)
  }

}

  // proto:  bool QAbstractItemModel::removeRows(int row, int count, const QModelIndex & parent);
func (this *QAbstractItemModel) removeRows(args ...interface{}) () {
  // removeRows(int, int, const class QModelIndex &)
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
    // invoke: _ZN18QAbstractItemModel10removeRowsEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "removeRows", args)
  }

}

  // proto:  bool QAbstractItemModel::hasChildren(const QModelIndex & parent);
func (this *QAbstractItemModel) hasChildren(args ...interface{}) () {
  // hasChildren(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel11hasChildrenERK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "hasChildren", args)
  }

}

  // proto:  bool QAbstractItemModel::moveRow(const QModelIndex & sourceParent, int sourceRow, const QModelIndex & destinationParent, int destinationChild);
func (this *QAbstractItemModel) moveRow(args ...interface{}) () {
  // moveRow(const class QModelIndex &, int, const class QModelIndex &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModel7moveRowERK11QModelIndexiS2_i
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "moveRow", args)
  }

}

  // proto:  void QAbstractItemModel::revert();
func (this *QAbstractItemModel) revert(args ...interface{}) () {
  // revert()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModel6revertEv
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "revert", args)
  }

}

  // proto:  bool QAbstractItemModel::removeColumn(int column, const QModelIndex & parent);
func (this *QAbstractItemModel) removeColumn(args ...interface{}) () {
  // removeColumn(int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModel12removeColumnEiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "removeColumn", args)
  }

}

  // proto:  bool QAbstractItemModel::insertColumns(int column, int count, const QModelIndex & parent);
func (this *QAbstractItemModel) insertColumns(args ...interface{}) () {
  // insertColumns(int, int, const class QModelIndex &)
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
    // invoke: _ZN18QAbstractItemModel13insertColumnsEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "insertColumns", args)
  }

}

  // proto:  bool QAbstractItemModel::insertColumn(int column, const QModelIndex & parent);
func (this *QAbstractItemModel) insertColumn(args ...interface{}) () {
  // insertColumn(int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModel12insertColumnEiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "insertColumn", args)
  }

}

  // proto:  bool QAbstractItemModel::moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild);
func (this *QAbstractItemModel) moveColumns(args ...interface{}) () {
  // moveColumns(const class QModelIndex &, int, int, const class QModelIndex &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModel11moveColumnsERK11QModelIndexiiS2_i
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "moveColumns", args)
  }

}

  // proto:  QMap<int, QVariant> QAbstractItemModel::itemData(const QModelIndex & index);
func (this *QAbstractItemModel) itemData(args ...interface{}) () {
  // itemData(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel8itemDataERK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "itemData", args)
  }

}

  // proto:  QStringList QAbstractItemModel::mimeTypes();
func (this *QAbstractItemModel) mimeTypes(args ...interface{}) () {
  // mimeTypes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel9mimeTypesEv
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "mimeTypes", args)
  }

}

  // proto:  QModelIndex QAbstractItemModel::parent(const QModelIndex & child);
func (this *QAbstractItemModel) parent(args ...interface{}) () {
  // parent(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel6parentERK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "parent", args)
  }

}

  // proto:  QModelIndex QAbstractItemModel::buddy(const QModelIndex & index);
func (this *QAbstractItemModel) buddy(args ...interface{}) () {
  // buddy(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel5buddyERK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "buddy", args)
  }

}

  // proto:  int QAbstractItemModel::columnCount(const QModelIndex & parent);
func (this *QAbstractItemModel) columnCount(args ...interface{}) () {
  // columnCount(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel11columnCountERK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "columnCount", args)
  }

}

  // proto:  QVariant QAbstractItemModel::data(const QModelIndex & index, int role);
func (this *QAbstractItemModel) data(args ...interface{}) () {
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
    // invoke: _ZNK18QAbstractItemModel4dataERK11QModelIndexi
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "data", args)
  }

}

  // proto:  QModelIndex QAbstractItemModel::sibling(int row, int column, const QModelIndex & idx);
func (this *QAbstractItemModel) sibling(args ...interface{}) () {
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
    // invoke: _ZNK18QAbstractItemModel7siblingEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "sibling", args)
  }

}

  // proto:  bool QAbstractItemModel::moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild);
func (this *QAbstractItemModel) moveRows(args ...interface{}) () {
  // moveRows(const class QModelIndex &, int, int, const class QModelIndex &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModel8moveRowsERK11QModelIndexiiS2_i
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "moveRows", args)
  }

}

  // proto:  QModelIndex QAbstractListModel::sibling(int row, int column, const QModelIndex & idx);
func (this *QAbstractListModel) sibling(args ...interface{}) () {
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
    // invoke: _ZNK18QAbstractListModel7siblingEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractListModel", "sibling", args)
  }

}

  // proto:  void QAbstractListModel::QAbstractListModel(QObject * parent);
func NewQAbstractListModel(args ...interface{}) QAbstractListModel {
  return QAbstractListModel{}
}

  // proto:  const QMetaObject * QAbstractListModel::metaObject();
func (this *QAbstractListModel) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractListModel10metaObjectEv
  default:
    qtrt.ErrorResolve("QAbstractListModel", "metaObject", args)
  }

}

  // proto:  QModelIndex QAbstractListModel::index(int row, int column, const QModelIndex & parent);
func (this *QAbstractListModel) index(args ...interface{}) () {
  // index(int, int, const class QModelIndex &)
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
    // invoke: _ZNK18QAbstractListModel5indexEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractListModel", "index", args)
  }

}

  // proto:  void QAbstractListModel::~QAbstractListModel();
func (this *QAbstractListModel) FreeQAbstractListModel(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractListModel", "~QAbstractListModel", args)
  }

}

// <= body block end

