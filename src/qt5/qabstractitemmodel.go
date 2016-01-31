package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
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
extern void* C_ZNK11QModelIndex6parentEv(void* qthis); // 2
  // proto:  void * QModelIndex::internalPointer();
extern void C_ZNK11QModelIndex15internalPointerEv(void* qthis); // 2
  // proto:  quintptr QModelIndex::internalId();
extern int32_t C_ZNK11QModelIndex10internalIdEv(void* qthis); // 2
  // proto:  int QModelIndex::column();
extern int32_t C_ZNK11QModelIndex6columnEv(void* qthis); // 2
  // proto:  bool QModelIndex::isValid();
extern bool C_ZNK11QModelIndex7isValidEv(void* qthis); // 2
  // proto:  QModelIndex QModelIndex::sibling(int row, int column);
extern void* C_ZNK11QModelIndex7siblingEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  Qt::ItemFlags QModelIndex::flags();
extern void C_ZNK11QModelIndex5flagsEv(void* qthis); // 2
  // proto:  QModelIndex QModelIndex::child(int row, int column);
extern void* C_ZNK11QModelIndex5childEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  const QAbstractItemModel * QModelIndex::model();
extern void C_ZNK11QModelIndex5modelEv(void* qthis); // 2
  // proto:  int QModelIndex::row();
extern int32_t C_ZNK11QModelIndex3rowEv(void* qthis); // 2
  // proto:  QVariant QModelIndex::data(int role);
extern void* C_ZNK11QModelIndex4dataEi(void* qthis, int32_t arg0); // 2
  // proto:  void QModelIndex::QModelIndex();
extern void* C_ZN11QModelIndexC2Ev(); // 1
  // proto:  void * QPersistentModelIndex::internalPointer();
extern void C_ZNK21QPersistentModelIndex15internalPointerEv(void* qthis); // 4
  // proto:  QModelIndex QPersistentModelIndex::parent();
extern void* C_ZNK21QPersistentModelIndex6parentEv(void* qthis); // 4
  // proto:  int QPersistentModelIndex::column();
extern int32_t C_ZNK21QPersistentModelIndex6columnEv(void* qthis); // 4
  // proto:  bool QPersistentModelIndex::isValid();
extern bool C_ZNK21QPersistentModelIndex7isValidEv(void* qthis); // 4
  // proto:  void QPersistentModelIndex::~QPersistentModelIndex();
extern void C_ZN21QPersistentModelIndexD2Ev(void* qthis); // 4
  // proto:  void QPersistentModelIndex::QPersistentModelIndex();
extern void* C_ZN21QPersistentModelIndexC2Ev(); // 3
  // proto:  void QPersistentModelIndex::QPersistentModelIndex(const QPersistentModelIndex & other);
extern void* C_ZN21QPersistentModelIndexC2ERKS_(void* arg0); // 3
  // proto:  void QPersistentModelIndex::QPersistentModelIndex(const QModelIndex & index);
extern void* C_ZN21QPersistentModelIndexC2ERK11QModelIndex(void* arg0); // 3
  // proto:  quintptr QPersistentModelIndex::internalId();
extern int32_t C_ZNK21QPersistentModelIndex10internalIdEv(void* qthis); // 4
  // proto:  QModelIndex QPersistentModelIndex::sibling(int row, int column);
extern void* C_ZNK21QPersistentModelIndex7siblingEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  Qt::ItemFlags QPersistentModelIndex::flags();
extern void C_ZNK21QPersistentModelIndex5flagsEv(void* qthis); // 4
  // proto:  void QPersistentModelIndex::swap(QPersistentModelIndex & other);
extern void C_ZN21QPersistentModelIndex4swapERS_(void* qthis, void* arg0); // 2
  // proto:  QModelIndex QPersistentModelIndex::child(int row, int column);
extern void* C_ZNK21QPersistentModelIndex5childEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  const QAbstractItemModel * QPersistentModelIndex::model();
extern void C_ZNK21QPersistentModelIndex5modelEv(void* qthis); // 4
  // proto:  int QPersistentModelIndex::row();
extern int32_t C_ZNK21QPersistentModelIndex3rowEv(void* qthis); // 4
  // proto:  QVariant QPersistentModelIndex::data(int role);
extern void* C_ZNK21QPersistentModelIndex4dataEi(void* qthis, int32_t arg0); // 4
  // proto:  QModelIndex QAbstractTableModel::index(int row, int column, const QModelIndex & parent);
extern void* C_ZNK19QAbstractTableModel5indexEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  QModelIndex QAbstractTableModel::sibling(int row, int column, const QModelIndex & idx);
extern void* C_ZNK19QAbstractTableModel7siblingEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QAbstractTableModel::QAbstractTableModel(QObject * parent);
extern void* C_ZN19QAbstractTableModelC2EP7QObject(void* arg0); // 3
  // proto:  Qt::ItemFlags QAbstractTableModel::flags(const QModelIndex & index);
extern void C_ZNK19QAbstractTableModel5flagsERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QAbstractTableModel::~QAbstractTableModel();
extern void C_ZN19QAbstractTableModelD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QAbstractTableModel::metaObject();
extern void C_ZNK19QAbstractTableModel10metaObjectEv(void* qthis); // 4
  // proto:  Qt::DropActions QAbstractItemModel::supportedDragActions();
extern void C_ZNK18QAbstractItemModel20supportedDragActionsEv(void* qthis); // 4
  // proto:  QModelIndex QAbstractItemModel::buddy(const QModelIndex & index);
extern void* C_ZNK18QAbstractItemModel5buddyERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QAbstractItemModel::moveRow(const QModelIndex & sourceParent, int sourceRow, const QModelIndex & destinationParent, int destinationChild);
extern bool C_ZN18QAbstractItemModel7moveRowERK11QModelIndexiS2_i(void* qthis, void* arg0, int32_t arg1, void* arg2, int32_t arg3); // 2
  // proto:  bool QAbstractItemModel::removeColumn(int column, const QModelIndex & parent);
extern bool C_ZN18QAbstractItemModel12removeColumnEiRK11QModelIndex(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  const QMetaObject * QAbstractItemModel::metaObject();
extern void C_ZNK18QAbstractItemModel10metaObjectEv(void* qthis); // 4
  // proto:  bool QAbstractItemModel::insertColumns(int column, int count, const QModelIndex & parent);
extern bool C_ZN18QAbstractItemModel13insertColumnsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QAbstractItemModel::revert();
extern void C_ZN18QAbstractItemModel6revertEv(void* qthis); // 4
  // proto:  bool QAbstractItemModel::insertRow(int row, const QModelIndex & parent);
extern bool C_ZN18QAbstractItemModel9insertRowEiRK11QModelIndex(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  void QAbstractItemModel::fetchMore(const QModelIndex & parent);
extern void C_ZN18QAbstractItemModel9fetchMoreERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  Qt::DropActions QAbstractItemModel::supportedDropActions();
extern void C_ZNK18QAbstractItemModel20supportedDropActionsEv(void* qthis); // 4
  // proto:  bool QAbstractItemModel::moveColumn(const QModelIndex & sourceParent, int sourceColumn, const QModelIndex & destinationParent, int destinationChild);
extern bool C_ZN18QAbstractItemModel10moveColumnERK11QModelIndexiS2_i(void* qthis, void* arg0, int32_t arg1, void* arg2, int32_t arg3); // 2
  // proto:  QMap<int, QVariant> QAbstractItemModel::itemData(const QModelIndex & index);
extern void C_ZNK18QAbstractItemModel8itemDataERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QAbstractItemModel::removeRow(int row, const QModelIndex & parent);
extern bool C_ZN18QAbstractItemModel9removeRowEiRK11QModelIndex(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  QStringList QAbstractItemModel::mimeTypes();
extern void C_ZNK18QAbstractItemModel9mimeTypesEv(void* qthis); // 4
  // proto:  void QAbstractItemModel::~QAbstractItemModel();
extern void C_ZN18QAbstractItemModelD2Ev(void* qthis); // 4
  // proto:  QHash<int, QByteArray> QAbstractItemModel::roleNames();
extern void C_ZNK18QAbstractItemModel9roleNamesEv(void* qthis); // 4
  // proto:  bool QAbstractItemModel::removeColumns(int column, int count, const QModelIndex & parent);
extern bool C_ZN18QAbstractItemModel13removeColumnsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QAbstractItemModel::QAbstractItemModel(QObject * parent);
extern void* C_ZN18QAbstractItemModelC2EP7QObject(void* arg0); // 3
  // proto:  QSize QAbstractItemModel::span(const QModelIndex & index);
extern void* C_ZNK18QAbstractItemModel4spanERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QAbstractItemModel::submit();
extern bool C_ZN18QAbstractItemModel6submitEv(void* qthis); // 4
  // proto:  bool QAbstractItemModel::removeRows(int row, int count, const QModelIndex & parent);
extern bool C_ZN18QAbstractItemModel10removeRowsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  bool QAbstractItemModel::canFetchMore(const QModelIndex & parent);
extern bool C_ZNK18QAbstractItemModel12canFetchMoreERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QAbstractItemModel::moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild);
extern bool C_ZN18QAbstractItemModel11moveColumnsERK11QModelIndexiiS2_i(void* qthis, void* arg0, int32_t arg1, int32_t arg2, void* arg3, int32_t arg4); // 4
  // proto:  QModelIndex QAbstractItemModel::sibling(int row, int column, const QModelIndex & idx);
extern void* C_ZNK18QAbstractItemModel7siblingEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  bool QAbstractItemModel::moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild);
extern bool C_ZN18QAbstractItemModel8moveRowsERK11QModelIndexiiS2_i(void* qthis, void* arg0, int32_t arg1, int32_t arg2, void* arg3, int32_t arg4); // 4
  // proto:  bool QAbstractItemModel::insertColumn(int column, const QModelIndex & parent);
extern bool C_ZN18QAbstractItemModel12insertColumnEiRK11QModelIndex(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  bool QAbstractItemModel::hasChildren(const QModelIndex & parent);
extern bool C_ZNK18QAbstractItemModel11hasChildrenERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QAbstractItemModel::hasIndex(int row, int column, const QModelIndex & parent);
extern bool C_ZNK18QAbstractItemModel8hasIndexEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  bool QAbstractItemModel::setData(const QModelIndex & index, const QVariant & value, int role);
extern bool C_ZN18QAbstractItemModel7setDataERK11QModelIndexRK8QVarianti(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  bool QAbstractItemModel::insertRows(int row, int count, const QModelIndex & parent);
extern bool C_ZN18QAbstractItemModel10insertRowsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  Qt::ItemFlags QAbstractItemModel::flags(const QModelIndex & index);
extern void C_ZNK18QAbstractItemModel5flagsERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QModelIndex QAbstractListModel::index(int row, int column, const QModelIndex & parent);
extern void* C_ZNK18QAbstractListModel5indexEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  const QMetaObject * QAbstractListModel::metaObject();
extern void C_ZNK18QAbstractListModel10metaObjectEv(void* qthis); // 4
  // proto:  void QAbstractListModel::QAbstractListModel(QObject * parent);
extern void* C_ZN18QAbstractListModelC2EP7QObject(void* arg0); // 3
  // proto:  QModelIndex QAbstractListModel::sibling(int row, int column, const QModelIndex & idx);
extern void* C_ZNK18QAbstractListModel7siblingEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
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
func (this *QModelIndex) Parent(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QModelIndex6parentEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QModelIndex", "parent", args)
  }

  return
}

// internalPointer()
func (this *QModelIndex) Internalpointer(args ...interface{}) () {
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

  return
}

// internalId()
func (this *QModelIndex) Internalid(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QModelIndex10internalIdEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "quintptr"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QModelIndex", "internalId", args)
  }

  return
}

// column()
func (this *QModelIndex) Column(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QModelIndex6columnEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QModelIndex", "column", args)
  }

  return
}

// isValid()
func (this *QModelIndex) Isvalid(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QModelIndex7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QModelIndex", "isValid", args)
  }

  return
}

// sibling(int, int)
func (this *QModelIndex) Sibling(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QModelIndex7siblingEii(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QModelIndex", "sibling", args)
  }

  return
}

// flags()
func (this *QModelIndex) Flags(args ...interface{}) () {
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

  return
}

// child(int, int)
func (this *QModelIndex) Child(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QModelIndex5childEii(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QModelIndex", "child", args)
  }

  return
}

// model()
func (this *QModelIndex) Model(args ...interface{}) () {
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

  return
}

// row()
func (this *QModelIndex) Row(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QModelIndex3rowEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QModelIndex", "row", args)
  }

  return
}

// data(int)
func (this *QModelIndex) Data(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QModelIndex4dataEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVariant{}) // "QVariant"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QModelIndex", "data", args)
  }

  return
}

// QModelIndex()
func NewQModelIndex(args ...interface{}) *QModelIndex {
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
    qthis = C.C_ZN11QModelIndexC2Ev()
    return &QModelIndex{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QModelIndex", "QModelIndex", args)
  }

  return nil // QModelIndex{qclsinst:qthis}
}

// internalPointer()
func (this *QPersistentModelIndex) Internalpointer(args ...interface{}) () {
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

  return
}

// parent()
func (this *QPersistentModelIndex) Parent(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK21QPersistentModelIndex6parentEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "parent", args)
  }

  return
}

// column()
func (this *QPersistentModelIndex) Column(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK21QPersistentModelIndex6columnEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "column", args)
  }

  return
}

// isValid()
func (this *QPersistentModelIndex) Isvalid(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK21QPersistentModelIndex7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "isValid", args)
  }

  return
}

// ~QPersistentModelIndex()
func (this *QPersistentModelIndex) Freeqpersistentmodelindex(args ...interface{}) () {
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

  return
}

// QPersistentModelIndex()
func NewQPersistentModelIndex(args ...interface{}) *QPersistentModelIndex {
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
    qthis = C.C_ZN21QPersistentModelIndexC2Ev()
    return &QPersistentModelIndex{qclsinst:qthis}
  case 1:
    // invoke: _ZN21QPersistentModelIndexC1ERKS_
    // invoke: void QPersistentModelIndex(const class QPersistentModelIndex &)
    var arg0 = args[0].(QPersistentModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN21QPersistentModelIndexC2ERKS_(arg0)
    return &QPersistentModelIndex{qclsinst:qthis}
  case 2:
    // invoke: _ZN21QPersistentModelIndexC1ERK11QModelIndex
    // invoke: void QPersistentModelIndex(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN21QPersistentModelIndexC2ERK11QModelIndex(arg0)
    return &QPersistentModelIndex{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "QPersistentModelIndex", args)
  }

  return nil // QPersistentModelIndex{qclsinst:qthis}
}

// internalId()
func (this *QPersistentModelIndex) Internalid(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK21QPersistentModelIndex10internalIdEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "quintptr"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "internalId", args)
  }

  return
}

// sibling(int, int)
func (this *QPersistentModelIndex) Sibling(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK21QPersistentModelIndex7siblingEii(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "sibling", args)
  }

  return
}

// flags()
func (this *QPersistentModelIndex) Flags(args ...interface{}) () {
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

  return
}

// swap(class QPersistentModelIndex &)
func (this *QPersistentModelIndex) Swap(args ...interface{}) () {
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

  return
}

// child(int, int)
func (this *QPersistentModelIndex) Child(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK21QPersistentModelIndex5childEii(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "child", args)
  }

  return
}

// model()
func (this *QPersistentModelIndex) Model(args ...interface{}) () {
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

  return
}

// row()
func (this *QPersistentModelIndex) Row(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK21QPersistentModelIndex3rowEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "row", args)
  }

  return
}

// data(int)
func (this *QPersistentModelIndex) Data(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK21QPersistentModelIndex4dataEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVariant{}) // "QVariant"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "data", args)
  }

  return
}

// index(int, int, const class QModelIndex &)
func (this *QAbstractTableModel) Index(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK19QAbstractTableModel5indexEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAbstractTableModel", "index", args)
  }

  return
}

// sibling(int, int, const class QModelIndex &)
func (this *QAbstractTableModel) Sibling(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK19QAbstractTableModel7siblingEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAbstractTableModel", "sibling", args)
  }

  return
}

// QAbstractTableModel(class QObject *)
func NewQAbstractTableModel(args ...interface{}) *QAbstractTableModel {
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
    qthis = C.C_ZN19QAbstractTableModelC2EP7QObject(arg0)
    return &QAbstractTableModel{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QAbstractTableModel", "QAbstractTableModel", args)
  }

  return nil // QAbstractTableModel{qclsinst:qthis}
}

// flags(const class QModelIndex &)
func (this *QAbstractTableModel) Flags(args ...interface{}) () {
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

  return
}

// ~QAbstractTableModel()
func (this *QAbstractTableModel) Freeqabstracttablemodel(args ...interface{}) () {
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

  return
}

// metaObject()
func (this *QAbstractTableModel) Metaobject(args ...interface{}) () {
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

  return
}

// supportedDragActions()
func (this *QAbstractItemModel) Supporteddragactions(args ...interface{}) () {
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

  return
}

// buddy(const class QModelIndex &)
func (this *QAbstractItemModel) Buddy(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK18QAbstractItemModel5buddyERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "buddy", args)
  }

  return
}

// moveRow(const class QModelIndex &, int, const class QModelIndex &, int)
func (this *QAbstractItemModel) Moverow(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN18QAbstractItemModel7moveRowERK11QModelIndexiS2_i(this.qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "moveRow", args)
  }

  return
}

// removeColumn(int, const class QModelIndex &)
func (this *QAbstractItemModel) Removecolumn(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN18QAbstractItemModel12removeColumnEiRK11QModelIndex(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "removeColumn", args)
  }

  return
}

// metaObject()
func (this *QAbstractItemModel) Metaobject(args ...interface{}) () {
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

  return
}

// insertColumns(int, int, const class QModelIndex &)
func (this *QAbstractItemModel) Insertcolumns(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN18QAbstractItemModel13insertColumnsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "insertColumns", args)
  }

  return
}

// revert()
func (this *QAbstractItemModel) Revert(args ...interface{}) () {
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

  return
}

// insertRow(int, const class QModelIndex &)
func (this *QAbstractItemModel) Insertrow(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN18QAbstractItemModel9insertRowEiRK11QModelIndex(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "insertRow", args)
  }

  return
}

// fetchMore(const class QModelIndex &)
func (this *QAbstractItemModel) Fetchmore(args ...interface{}) () {
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

  return
}

// supportedDropActions()
func (this *QAbstractItemModel) Supporteddropactions(args ...interface{}) () {
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

  return
}

// moveColumn(const class QModelIndex &, int, const class QModelIndex &, int)
func (this *QAbstractItemModel) Movecolumn(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN18QAbstractItemModel10moveColumnERK11QModelIndexiS2_i(this.qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "moveColumn", args)
  }

  return
}

// itemData(const class QModelIndex &)
func (this *QAbstractItemModel) Itemdata(args ...interface{}) () {
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

  return
}

// removeRow(int, const class QModelIndex &)
func (this *QAbstractItemModel) Removerow(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN18QAbstractItemModel9removeRowEiRK11QModelIndex(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "removeRow", args)
  }

  return
}

// mimeTypes()
func (this *QAbstractItemModel) Mimetypes(args ...interface{}) () {
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

  return
}

// ~QAbstractItemModel()
func (this *QAbstractItemModel) Freeqabstractitemmodel(args ...interface{}) () {
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

  return
}

// roleNames()
func (this *QAbstractItemModel) Rolenames(args ...interface{}) () {
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

  return
}

// removeColumns(int, int, const class QModelIndex &)
func (this *QAbstractItemModel) Removecolumns(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN18QAbstractItemModel13removeColumnsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "removeColumns", args)
  }

  return
}

// QAbstractItemModel(class QObject *)
func NewQAbstractItemModel(args ...interface{}) *QAbstractItemModel {
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
    qthis = C.C_ZN18QAbstractItemModelC2EP7QObject(arg0)
    return &QAbstractItemModel{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "QAbstractItemModel", args)
  }

  return nil // QAbstractItemModel{qclsinst:qthis}
}

// span(const class QModelIndex &)
func (this *QAbstractItemModel) Span(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK18QAbstractItemModel4spanERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "span", args)
  }

  return
}

// submit()
func (this *QAbstractItemModel) Submit(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN18QAbstractItemModel6submitEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "submit", args)
  }

  return
}

// removeRows(int, int, const class QModelIndex &)
func (this *QAbstractItemModel) Removerows(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN18QAbstractItemModel10removeRowsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "removeRows", args)
  }

  return
}

// canFetchMore(const class QModelIndex &)
func (this *QAbstractItemModel) Canfetchmore(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK18QAbstractItemModel12canFetchMoreERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "canFetchMore", args)
  }

  return
}

// moveColumns(const class QModelIndex &, int, int, const class QModelIndex &, int)
func (this *QAbstractItemModel) Movecolumns(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN18QAbstractItemModel11moveColumnsERK11QModelIndexiiS2_i(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "moveColumns", args)
  }

  return
}

// sibling(int, int, const class QModelIndex &)
func (this *QAbstractItemModel) Sibling(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK18QAbstractItemModel7siblingEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "sibling", args)
  }

  return
}

// moveRows(const class QModelIndex &, int, int, const class QModelIndex &, int)
func (this *QAbstractItemModel) Moverows(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN18QAbstractItemModel8moveRowsERK11QModelIndexiiS2_i(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "moveRows", args)
  }

  return
}

// insertColumn(int, const class QModelIndex &)
func (this *QAbstractItemModel) Insertcolumn(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN18QAbstractItemModel12insertColumnEiRK11QModelIndex(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "insertColumn", args)
  }

  return
}

// hasChildren(const class QModelIndex &)
func (this *QAbstractItemModel) Haschildren(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK18QAbstractItemModel11hasChildrenERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "hasChildren", args)
  }

  return
}

// hasIndex(int, int, const class QModelIndex &)
func (this *QAbstractItemModel) Hasindex(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK18QAbstractItemModel8hasIndexEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "hasIndex", args)
  }

  return
}

// setData(const class QModelIndex &, const class QVariant &, int)
func (this *QAbstractItemModel) Setdata(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN18QAbstractItemModel7setDataERK11QModelIndexRK8QVarianti(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "setData", args)
  }

  return
}

// insertRows(int, int, const class QModelIndex &)
func (this *QAbstractItemModel) Insertrows(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN18QAbstractItemModel10insertRowsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "insertRows", args)
  }

  return
}

// flags(const class QModelIndex &)
func (this *QAbstractItemModel) Flags(args ...interface{}) () {
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

  return
}

// index(int, int, const class QModelIndex &)
func (this *QAbstractListModel) Index(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK18QAbstractListModel5indexEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAbstractListModel", "index", args)
  }

  return
}

// metaObject()
func (this *QAbstractListModel) Metaobject(args ...interface{}) () {
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

  return
}

// QAbstractListModel(class QObject *)
func NewQAbstractListModel(args ...interface{}) *QAbstractListModel {
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
    qthis = C.C_ZN18QAbstractListModelC2EP7QObject(arg0)
    return &QAbstractListModel{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QAbstractListModel", "QAbstractListModel", args)
  }

  return nil // QAbstractListModel{qclsinst:qthis}
}

// sibling(int, int, const class QModelIndex &)
func (this *QAbstractListModel) Sibling(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK18QAbstractListModel7siblingEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAbstractListModel", "sibling", args)
  }

  return
}

// flags(const class QModelIndex &)
func (this *QAbstractListModel) Flags(args ...interface{}) () {
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

  return
}

// ~QAbstractListModel()
func (this *QAbstractListModel) Freeqabstractlistmodel(args ...interface{}) () {
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

  return
}

// <= body block end

