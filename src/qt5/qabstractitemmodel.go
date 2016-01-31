package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QModelIndex QModelIndex::parent();
extern void C_ZNK11QModelIndex6parentEv(void* qthis); // 2
  // proto:  void * QModelIndex::internalPointer();
extern void C_ZNK11QModelIndex15internalPointerEv(void* qthis); // 2
  // proto:  quintptr QModelIndex::internalId();
extern void C_ZNK11QModelIndex10internalIdEv(void* qthis); // 2
  // proto:  int QModelIndex::column();
extern void C_ZNK11QModelIndex6columnEv(void* qthis); // 2
  // proto:  bool QModelIndex::isValid();
extern void C_ZNK11QModelIndex7isValidEv(void* qthis); // 2
  // proto:  QModelIndex QModelIndex::sibling(int row, int column);
extern void C_ZNK11QModelIndex7siblingEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  Qt::ItemFlags QModelIndex::flags();
extern void C_ZNK11QModelIndex5flagsEv(void* qthis); // 2
  // proto:  QModelIndex QModelIndex::child(int row, int column);
extern void C_ZNK11QModelIndex5childEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  const QAbstractItemModel * QModelIndex::model();
extern void C_ZNK11QModelIndex5modelEv(void* qthis); // 2
  // proto:  int QModelIndex::row();
extern void C_ZNK11QModelIndex3rowEv(void* qthis); // 2
  // proto:  QVariant QModelIndex::data(int role);
extern void C_ZNK11QModelIndex4dataEi(void* qthis, int32_t arg0); // 2
  // proto:  void QModelIndex::QModelIndex();
extern void C_ZN11QModelIndexC2Ev(void* qthis); // 1
  // proto:  void * QPersistentModelIndex::internalPointer();
extern void C_ZNK21QPersistentModelIndex15internalPointerEv(void* qthis); // 4
  // proto:  QModelIndex QPersistentModelIndex::parent();
extern void C_ZNK21QPersistentModelIndex6parentEv(void* qthis); // 4
  // proto:  int QPersistentModelIndex::column();
extern void C_ZNK21QPersistentModelIndex6columnEv(void* qthis); // 4
  // proto:  bool QPersistentModelIndex::isValid();
extern void C_ZNK21QPersistentModelIndex7isValidEv(void* qthis); // 4
  // proto:  void QPersistentModelIndex::~QPersistentModelIndex();
extern void C_ZN21QPersistentModelIndexD2Ev(void* qthis); // 4
  // proto:  void QPersistentModelIndex::QPersistentModelIndex();
extern void C_ZN21QPersistentModelIndexC2Ev(void* qthis); // 3
  // proto:  void QPersistentModelIndex::QPersistentModelIndex(const QPersistentModelIndex & other);
extern void C_ZN21QPersistentModelIndexC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  void QPersistentModelIndex::QPersistentModelIndex(const QModelIndex & index);
extern void C_ZN21QPersistentModelIndexC2ERK11QModelIndex(void* qthis, void* arg0); // 3
  // proto:  quintptr QPersistentModelIndex::internalId();
extern void C_ZNK21QPersistentModelIndex10internalIdEv(void* qthis); // 4
  // proto:  QModelIndex QPersistentModelIndex::sibling(int row, int column);
extern void C_ZNK21QPersistentModelIndex7siblingEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  Qt::ItemFlags QPersistentModelIndex::flags();
extern void C_ZNK21QPersistentModelIndex5flagsEv(void* qthis); // 4
  // proto:  void QPersistentModelIndex::swap(QPersistentModelIndex & other);
extern void C_ZN21QPersistentModelIndex4swapERS_(void* qthis, void* arg0); // 2
  // proto:  QModelIndex QPersistentModelIndex::child(int row, int column);
extern void C_ZNK21QPersistentModelIndex5childEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  const QAbstractItemModel * QPersistentModelIndex::model();
extern void C_ZNK21QPersistentModelIndex5modelEv(void* qthis); // 4
  // proto:  int QPersistentModelIndex::row();
extern void C_ZNK21QPersistentModelIndex3rowEv(void* qthis); // 4
  // proto:  QVariant QPersistentModelIndex::data(int role);
extern void C_ZNK21QPersistentModelIndex4dataEi(void* qthis, int32_t arg0); // 4
  // proto:  QModelIndex QAbstractTableModel::index(int row, int column, const QModelIndex & parent);
extern void C_ZNK19QAbstractTableModel5indexEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  QModelIndex QAbstractTableModel::sibling(int row, int column, const QModelIndex & idx);
extern void C_ZNK19QAbstractTableModel7siblingEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QAbstractTableModel::QAbstractTableModel(QObject * parent);
extern void C_ZN19QAbstractTableModelC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  Qt::ItemFlags QAbstractTableModel::flags(const QModelIndex & index);
extern void C_ZNK19QAbstractTableModel5flagsERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QAbstractTableModel::~QAbstractTableModel();
extern void C_ZN19QAbstractTableModelD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QAbstractTableModel::metaObject();
extern void C_ZNK19QAbstractTableModel10metaObjectEv(void* qthis); // 4
  // proto:  Qt::DropActions QAbstractItemModel::supportedDragActions();
extern void C_ZNK18QAbstractItemModel20supportedDragActionsEv(void* qthis); // 4
  // proto:  QModelIndex QAbstractItemModel::buddy(const QModelIndex & index);
extern void C_ZNK18QAbstractItemModel5buddyERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QAbstractItemModel::moveRow(const QModelIndex & sourceParent, int sourceRow, const QModelIndex & destinationParent, int destinationChild);
extern void C_ZN18QAbstractItemModel7moveRowERK11QModelIndexiS2_i(void* qthis, void* arg0, int32_t arg1, void* arg2, int32_t arg3); // 2
  // proto:  bool QAbstractItemModel::removeColumn(int column, const QModelIndex & parent);
extern void C_ZN18QAbstractItemModel12removeColumnEiRK11QModelIndex(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  const QMetaObject * QAbstractItemModel::metaObject();
extern void C_ZNK18QAbstractItemModel10metaObjectEv(void* qthis); // 4
  // proto:  bool QAbstractItemModel::insertColumns(int column, int count, const QModelIndex & parent);
extern void C_ZN18QAbstractItemModel13insertColumnsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QAbstractItemModel::revert();
extern void C_ZN18QAbstractItemModel6revertEv(void* qthis); // 4
  // proto:  bool QAbstractItemModel::insertRow(int row, const QModelIndex & parent);
extern void C_ZN18QAbstractItemModel9insertRowEiRK11QModelIndex(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  void QAbstractItemModel::fetchMore(const QModelIndex & parent);
extern void C_ZN18QAbstractItemModel9fetchMoreERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  Qt::DropActions QAbstractItemModel::supportedDropActions();
extern void C_ZNK18QAbstractItemModel20supportedDropActionsEv(void* qthis); // 4
  // proto:  bool QAbstractItemModel::moveColumn(const QModelIndex & sourceParent, int sourceColumn, const QModelIndex & destinationParent, int destinationChild);
extern void C_ZN18QAbstractItemModel10moveColumnERK11QModelIndexiS2_i(void* qthis, void* arg0, int32_t arg1, void* arg2, int32_t arg3); // 2
  // proto:  QMap<int, QVariant> QAbstractItemModel::itemData(const QModelIndex & index);
extern void C_ZNK18QAbstractItemModel8itemDataERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QAbstractItemModel::removeRow(int row, const QModelIndex & parent);
extern void C_ZN18QAbstractItemModel9removeRowEiRK11QModelIndex(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  QStringList QAbstractItemModel::mimeTypes();
extern void C_ZNK18QAbstractItemModel9mimeTypesEv(void* qthis); // 4
  // proto:  void QAbstractItemModel::~QAbstractItemModel();
extern void C_ZN18QAbstractItemModelD2Ev(void* qthis); // 4
  // proto:  QHash<int, QByteArray> QAbstractItemModel::roleNames();
extern void C_ZNK18QAbstractItemModel9roleNamesEv(void* qthis); // 4
  // proto:  bool QAbstractItemModel::removeColumns(int column, int count, const QModelIndex & parent);
extern void C_ZN18QAbstractItemModel13removeColumnsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QAbstractItemModel::QAbstractItemModel(QObject * parent);
extern void C_ZN18QAbstractItemModelC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  QSize QAbstractItemModel::span(const QModelIndex & index);
extern void C_ZNK18QAbstractItemModel4spanERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QAbstractItemModel::submit();
extern void C_ZN18QAbstractItemModel6submitEv(void* qthis); // 4
  // proto:  bool QAbstractItemModel::removeRows(int row, int count, const QModelIndex & parent);
extern void C_ZN18QAbstractItemModel10removeRowsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  bool QAbstractItemModel::canFetchMore(const QModelIndex & parent);
extern void C_ZNK18QAbstractItemModel12canFetchMoreERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QAbstractItemModel::moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild);
extern void C_ZN18QAbstractItemModel11moveColumnsERK11QModelIndexiiS2_i(void* qthis, void* arg0, int32_t arg1, int32_t arg2, void* arg3, int32_t arg4); // 4
  // proto:  QModelIndex QAbstractItemModel::sibling(int row, int column, const QModelIndex & idx);
extern void C_ZNK18QAbstractItemModel7siblingEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  bool QAbstractItemModel::moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild);
extern void C_ZN18QAbstractItemModel8moveRowsERK11QModelIndexiiS2_i(void* qthis, void* arg0, int32_t arg1, int32_t arg2, void* arg3, int32_t arg4); // 4
  // proto:  bool QAbstractItemModel::insertColumn(int column, const QModelIndex & parent);
extern void C_ZN18QAbstractItemModel12insertColumnEiRK11QModelIndex(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  bool QAbstractItemModel::hasChildren(const QModelIndex & parent);
extern void C_ZNK18QAbstractItemModel11hasChildrenERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QAbstractItemModel::hasIndex(int row, int column, const QModelIndex & parent);
extern void C_ZNK18QAbstractItemModel8hasIndexEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  bool QAbstractItemModel::setData(const QModelIndex & index, const QVariant & value, int role);
extern void C_ZN18QAbstractItemModel7setDataERK11QModelIndexRK8QVarianti(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  bool QAbstractItemModel::insertRows(int row, int count, const QModelIndex & parent);
extern void C_ZN18QAbstractItemModel10insertRowsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  Qt::ItemFlags QAbstractItemModel::flags(const QModelIndex & index);
extern void C_ZNK18QAbstractItemModel5flagsERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QModelIndex QAbstractListModel::index(int row, int column, const QModelIndex & parent);
extern void C_ZNK18QAbstractListModel5indexEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  const QMetaObject * QAbstractListModel::metaObject();
extern void C_ZNK18QAbstractListModel10metaObjectEv(void* qthis); // 4
  // proto:  void QAbstractListModel::QAbstractListModel(QObject * parent);
extern void C_ZN18QAbstractListModelC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  QModelIndex QAbstractListModel::sibling(int row, int column, const QModelIndex & idx);
extern void C_ZNK18QAbstractListModel7siblingEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  Qt::ItemFlags QAbstractListModel::flags(const QModelIndex & index);
extern void C_ZNK18QAbstractListModel5flagsERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QAbstractListModel::~QAbstractListModel();
extern void C_ZN18QAbstractListModelD2Ev(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QPersistentModelIndex)=8
type QPersistentModelIndex struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAbstractTableModel)=1
type QAbstractTableModel struct {
  /*qbase*/ QAbstractItemModel;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAbstractItemModel)=1
type QAbstractItemModel struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// parent()
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
    // invoke: QModelIndex parent()
    C.C_ZNK11QModelIndex6parentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QModelIndex", "parent", args)
  }

}

// internalPointer()
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
    // invoke: void * internalPointer()
    C.C_ZNK11QModelIndex15internalPointerEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QModelIndex", "internalPointer", args)
  }

}

// internalId()
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
    // invoke: quintptr internalId()
    C.C_ZNK11QModelIndex10internalIdEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QModelIndex", "internalId", args)
  }

}

// column()
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
    // invoke: int column()
    C.C_ZNK11QModelIndex6columnEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QModelIndex", "column", args)
  }

}

// isValid()
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
    // invoke: bool isValid()
    C.C_ZNK11QModelIndex7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QModelIndex", "isValid", args)
  }

}

// sibling(int, int)
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
    // invoke: QModelIndex sibling(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK11QModelIndex7siblingEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QModelIndex", "sibling", args)
  }

}

// flags()
func (this *QModelIndex) flags(args ...interface{}) () {
  // flags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QModelIndex5flagsEv
    // invoke: Qt::ItemFlags flags()
    C.C_ZNK11QModelIndex5flagsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QModelIndex", "flags", args)
  }

}

// child(int, int)
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
    // invoke: QModelIndex child(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK11QModelIndex5childEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QModelIndex", "child", args)
  }

}

// model()
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
    // invoke: const QAbstractItemModel * model()
    C.C_ZNK11QModelIndex5modelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QModelIndex", "model", args)
  }

}

// row()
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
    // invoke: int row()
    C.C_ZNK11QModelIndex3rowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QModelIndex", "row", args)
  }

}

// data(int)
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
    // invoke: QVariant data(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QModelIndex4dataEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QModelIndex", "data", args)
  }

}

// QModelIndex()
func NewQModelIndex(args ...interface{}) QModelIndex {
  // QModelIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QModelIndexC1Ev
    // invoke: void QModelIndex()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QModelIndexC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QModelIndex", "QModelIndex", args)
  }

  return QModelIndex{}
}

// internalPointer()
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
    // invoke: void * internalPointer()
    C.C_ZNK21QPersistentModelIndex15internalPointerEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "internalPointer", args)
  }

}

// parent()
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
    // invoke: QModelIndex parent()
    C.C_ZNK21QPersistentModelIndex6parentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "parent", args)
  }

}

// column()
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
    // invoke: int column()
    C.C_ZNK21QPersistentModelIndex6columnEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "column", args)
  }

}

// isValid()
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
    // invoke: bool isValid()
    C.C_ZNK21QPersistentModelIndex7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "isValid", args)
  }

}

// ~QPersistentModelIndex()
func (this *QPersistentModelIndex) FreeQPersistentModelIndex(args ...interface{}) () {
  // ~QPersistentModelIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QPersistentModelIndexD0Ev
    // invoke: void ~QPersistentModelIndex()
    C.C_ZN21QPersistentModelIndexD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "~QPersistentModelIndex", args)
  }

}

// QPersistentModelIndex()
func NewQPersistentModelIndex(args ...interface{}) QPersistentModelIndex {
  // QPersistentModelIndex()
  // QPersistentModelIndex(const class QPersistentModelIndex &)
  // QPersistentModelIndex(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPersistentModelIndex{}) // "const QPersistentModelIndex &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QPersistentModelIndexC1Ev
    // invoke: void QPersistentModelIndex()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN21QPersistentModelIndexC2Ev(qthis)
  case 1:
    // invoke: _ZN21QPersistentModelIndexC1ERKS_
    // invoke: void QPersistentModelIndex(const class QPersistentModelIndex &)
    var arg0 = args[0].(QPersistentModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN21QPersistentModelIndexC2ERKS_(qthis, arg0)
  case 2:
    // invoke: _ZN21QPersistentModelIndexC1ERK11QModelIndex
    // invoke: void QPersistentModelIndex(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN21QPersistentModelIndexC2ERK11QModelIndex(qthis, arg0)
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "QPersistentModelIndex", args)
  }

  return QPersistentModelIndex{}
}

// internalId()
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
    // invoke: quintptr internalId()
    C.C_ZNK21QPersistentModelIndex10internalIdEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "internalId", args)
  }

}

// sibling(int, int)
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
    // invoke: QModelIndex sibling(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK21QPersistentModelIndex7siblingEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "sibling", args)
  }

}

// flags()
func (this *QPersistentModelIndex) flags(args ...interface{}) () {
  // flags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QPersistentModelIndex5flagsEv
    // invoke: Qt::ItemFlags flags()
    C.C_ZNK21QPersistentModelIndex5flagsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "flags", args)
  }

}

// swap(class QPersistentModelIndex &)
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
    // invoke: void swap(class QPersistentModelIndex &)
    var arg0 = args[0].(QPersistentModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN21QPersistentModelIndex4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "swap", args)
  }

}

// child(int, int)
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
    // invoke: QModelIndex child(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK21QPersistentModelIndex5childEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "child", args)
  }

}

// model()
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
    // invoke: const QAbstractItemModel * model()
    C.C_ZNK21QPersistentModelIndex5modelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "model", args)
  }

}

// row()
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
    // invoke: int row()
    C.C_ZNK21QPersistentModelIndex3rowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "row", args)
  }

}

// data(int)
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
    // invoke: QVariant data(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK21QPersistentModelIndex4dataEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "data", args)
  }

}

// index(int, int, const class QModelIndex &)
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
    // invoke: QModelIndex index(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZNK19QAbstractTableModel5indexEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QAbstractTableModel", "index", args)
  }

}

// sibling(int, int, const class QModelIndex &)
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
    // invoke: QModelIndex sibling(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZNK19QAbstractTableModel7siblingEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QAbstractTableModel", "sibling", args)
  }

}

// QAbstractTableModel(class QObject *)
func NewQAbstractTableModel(args ...interface{}) QAbstractTableModel {
  // QAbstractTableModel(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractTableModelC1EP7QObject
    // invoke: void QAbstractTableModel(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN19QAbstractTableModelC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QAbstractTableModel", "QAbstractTableModel", args)
  }

  return QAbstractTableModel{}
}

// flags(const class QModelIndex &)
func (this *QAbstractTableModel) flags(args ...interface{}) () {
  // flags(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractTableModel5flagsERK11QModelIndex
    // invoke: Qt::ItemFlags flags(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK19QAbstractTableModel5flagsERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractTableModel", "flags", args)
  }

}

// ~QAbstractTableModel()
func (this *QAbstractTableModel) FreeQAbstractTableModel(args ...interface{}) () {
  // ~QAbstractTableModel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractTableModelD0Ev
    // invoke: void ~QAbstractTableModel()
    C.C_ZN19QAbstractTableModelD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractTableModel", "~QAbstractTableModel", args)
  }

}

// metaObject()
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
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK19QAbstractTableModel10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractTableModel", "metaObject", args)
  }

}

// supportedDragActions()
func (this *QAbstractItemModel) supportedDragActions(args ...interface{}) () {
  // supportedDragActions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel20supportedDragActionsEv
    // invoke: Qt::DropActions supportedDragActions()
    C.C_ZNK18QAbstractItemModel20supportedDragActionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "supportedDragActions", args)
  }

}

// buddy(const class QModelIndex &)
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
    // invoke: QModelIndex buddy(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK18QAbstractItemModel5buddyERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "buddy", args)
  }

}

// moveRow(const class QModelIndex &, int, const class QModelIndex &, int)
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
    // invoke: bool moveRow(const class QModelIndex &, int, const class QModelIndex &, int)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN18QAbstractItemModel7moveRowERK11QModelIndexiS2_i(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "moveRow", args)
  }

}

// removeColumn(int, const class QModelIndex &)
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
    // invoke: bool removeColumn(int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QModelIndex).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN18QAbstractItemModel12removeColumnEiRK11QModelIndex(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "removeColumn", args)
  }

}

// metaObject()
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
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK18QAbstractItemModel10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "metaObject", args)
  }

}

// insertColumns(int, int, const class QModelIndex &)
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
    // invoke: bool insertColumns(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN18QAbstractItemModel13insertColumnsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "insertColumns", args)
  }

}

// revert()
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
    // invoke: void revert()
    C.C_ZN18QAbstractItemModel6revertEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "revert", args)
  }

}

// insertRow(int, const class QModelIndex &)
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
    // invoke: bool insertRow(int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QModelIndex).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN18QAbstractItemModel9insertRowEiRK11QModelIndex(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "insertRow", args)
  }

}

// fetchMore(const class QModelIndex &)
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
    // invoke: void fetchMore(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QAbstractItemModel9fetchMoreERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "fetchMore", args)
  }

}

// supportedDropActions()
func (this *QAbstractItemModel) supportedDropActions(args ...interface{}) () {
  // supportedDropActions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel20supportedDropActionsEv
    // invoke: Qt::DropActions supportedDropActions()
    C.C_ZNK18QAbstractItemModel20supportedDropActionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "supportedDropActions", args)
  }

}

// moveColumn(const class QModelIndex &, int, const class QModelIndex &, int)
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
    // invoke: bool moveColumn(const class QModelIndex &, int, const class QModelIndex &, int)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN18QAbstractItemModel10moveColumnERK11QModelIndexiS2_i(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "moveColumn", args)
  }

}

// itemData(const class QModelIndex &)
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
    // invoke: QMap<int, QVariant> itemData(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK18QAbstractItemModel8itemDataERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "itemData", args)
  }

}

// removeRow(int, const class QModelIndex &)
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
    // invoke: bool removeRow(int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QModelIndex).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN18QAbstractItemModel9removeRowEiRK11QModelIndex(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "removeRow", args)
  }

}

// mimeTypes()
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
    // invoke: QStringList mimeTypes()
    C.C_ZNK18QAbstractItemModel9mimeTypesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "mimeTypes", args)
  }

}

// ~QAbstractItemModel()
func (this *QAbstractItemModel) FreeQAbstractItemModel(args ...interface{}) () {
  // ~QAbstractItemModel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModelD0Ev
    // invoke: void ~QAbstractItemModel()
    C.C_ZN18QAbstractItemModelD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "~QAbstractItemModel", args)
  }

}

// roleNames()
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
    // invoke: QHash<int, QByteArray> roleNames()
    C.C_ZNK18QAbstractItemModel9roleNamesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "roleNames", args)
  }

}

// removeColumns(int, int, const class QModelIndex &)
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
    // invoke: bool removeColumns(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN18QAbstractItemModel13removeColumnsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "removeColumns", args)
  }

}

// QAbstractItemModel(class QObject *)
func NewQAbstractItemModel(args ...interface{}) QAbstractItemModel {
  // QAbstractItemModel(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModelC1EP7QObject
    // invoke: void QAbstractItemModel(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN18QAbstractItemModelC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "QAbstractItemModel", args)
  }

  return QAbstractItemModel{}
}

// span(const class QModelIndex &)
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
    // invoke: QSize span(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK18QAbstractItemModel4spanERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "span", args)
  }

}

// submit()
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
    // invoke: bool submit()
    C.C_ZN18QAbstractItemModel6submitEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "submit", args)
  }

}

// removeRows(int, int, const class QModelIndex &)
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
    // invoke: bool removeRows(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN18QAbstractItemModel10removeRowsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "removeRows", args)
  }

}

// canFetchMore(const class QModelIndex &)
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
    // invoke: bool canFetchMore(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK18QAbstractItemModel12canFetchMoreERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "canFetchMore", args)
  }

}

// moveColumns(const class QModelIndex &, int, int, const class QModelIndex &, int)
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
    // invoke: bool moveColumns(const class QModelIndex &, int, int, const class QModelIndex &, int)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QModelIndex).qclsinst
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    C.C_ZN18QAbstractItemModel11moveColumnsERK11QModelIndexiiS2_i(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "moveColumns", args)
  }

}

// sibling(int, int, const class QModelIndex &)
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
    // invoke: QModelIndex sibling(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZNK18QAbstractItemModel7siblingEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "sibling", args)
  }

}

// moveRows(const class QModelIndex &, int, int, const class QModelIndex &, int)
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
    // invoke: bool moveRows(const class QModelIndex &, int, int, const class QModelIndex &, int)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QModelIndex).qclsinst
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    C.C_ZN18QAbstractItemModel8moveRowsERK11QModelIndexiiS2_i(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "moveRows", args)
  }

}

// insertColumn(int, const class QModelIndex &)
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
    // invoke: bool insertColumn(int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QModelIndex).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN18QAbstractItemModel12insertColumnEiRK11QModelIndex(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "insertColumn", args)
  }

}

// hasChildren(const class QModelIndex &)
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
    // invoke: bool hasChildren(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK18QAbstractItemModel11hasChildrenERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "hasChildren", args)
  }

}

// hasIndex(int, int, const class QModelIndex &)
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
    // invoke: bool hasIndex(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZNK18QAbstractItemModel8hasIndexEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "hasIndex", args)
  }

}

// setData(const class QModelIndex &, const class QVariant &, int)
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
    // invoke: bool setData(const class QModelIndex &, const class QVariant &, int)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN18QAbstractItemModel7setDataERK11QModelIndexRK8QVarianti(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "setData", args)
  }

}

// insertRows(int, int, const class QModelIndex &)
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
    // invoke: bool insertRows(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN18QAbstractItemModel10insertRowsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "insertRows", args)
  }

}

// flags(const class QModelIndex &)
func (this *QAbstractItemModel) flags(args ...interface{}) () {
  // flags(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel5flagsERK11QModelIndex
    // invoke: Qt::ItemFlags flags(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK18QAbstractItemModel5flagsERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "flags", args)
  }

}

// index(int, int, const class QModelIndex &)
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
    // invoke: QModelIndex index(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZNK18QAbstractListModel5indexEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QAbstractListModel", "index", args)
  }

}

// metaObject()
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
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK18QAbstractListModel10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractListModel", "metaObject", args)
  }

}

// QAbstractListModel(class QObject *)
func NewQAbstractListModel(args ...interface{}) QAbstractListModel {
  // QAbstractListModel(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractListModelC1EP7QObject
    // invoke: void QAbstractListModel(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN18QAbstractListModelC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QAbstractListModel", "QAbstractListModel", args)
  }

  return QAbstractListModel{}
}

// sibling(int, int, const class QModelIndex &)
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
    // invoke: QModelIndex sibling(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZNK18QAbstractListModel7siblingEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QAbstractListModel", "sibling", args)
  }

}

// flags(const class QModelIndex &)
func (this *QAbstractListModel) flags(args ...interface{}) () {
  // flags(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractListModel5flagsERK11QModelIndex
    // invoke: Qt::ItemFlags flags(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK18QAbstractListModel5flagsERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractListModel", "flags", args)
  }

}

// ~QAbstractListModel()
func (this *QAbstractListModel) FreeQAbstractListModel(args ...interface{}) () {
  // ~QAbstractListModel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractListModelD0Ev
    // invoke: void ~QAbstractListModel()
    C.C_ZN18QAbstractListModelD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractListModel", "~QAbstractListModel", args)
  }

}

// <= body block end

