package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtCore/qsortfilterproxymodel.h
// dst-file: /src/core/qsortfilterproxymodel.go
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
  // proto:  int QSortFilterProxyModel::sortColumn();
extern void C_ZNK21QSortFilterProxyModel10sortColumnEv(void* qthis); // 4
  // proto:  void QSortFilterProxyModel::invalidate();
extern void C_ZN21QSortFilterProxyModel10invalidateEv(void* qthis); // 4
  // proto:  void QSortFilterProxyModel::setDynamicSortFilter(bool enable);
extern void C_ZN21QSortFilterProxyModel20setDynamicSortFilterEb(void* qthis, bool arg0); // 4
  // proto:  bool QSortFilterProxyModel::canFetchMore(const QModelIndex & parent);
extern void C_ZNK21QSortFilterProxyModel12canFetchMoreERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QItemSelection QSortFilterProxyModel::mapSelectionToSource(const QItemSelection & proxySelection);
extern void C_ZNK21QSortFilterProxyModel20mapSelectionToSourceERK14QItemSelection(void* qthis, void* arg0); // 4
  // proto:  void QSortFilterProxyModel::setSourceModel(QAbstractItemModel * sourceModel);
extern void C_ZN21QSortFilterProxyModel14setSourceModelEP18QAbstractItemModel(void* qthis, void* arg0); // 4
  // proto:  QModelIndex QSortFilterProxyModel::sibling(int row, int column, const QModelIndex & idx);
extern void C_ZNK21QSortFilterProxyModel7siblingEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  Qt::CaseSensitivity QSortFilterProxyModel::filterCaseSensitivity();
extern void C_ZNK21QSortFilterProxyModel21filterCaseSensitivityEv(void* qthis); // 4
  // proto:  QModelIndex QSortFilterProxyModel::parent(const QModelIndex & child);
extern void C_ZNK21QSortFilterProxyModel6parentERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QModelIndex QSortFilterProxyModel::index(int row, int column, const QModelIndex & parent);
extern void C_ZNK21QSortFilterProxyModel5indexEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QSortFilterProxyModel::setFilterKeyColumn(int column);
extern void C_ZN21QSortFilterProxyModel18setFilterKeyColumnEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QSortFilterProxyModel::span(const QModelIndex & index);
extern void C_ZNK21QSortFilterProxyModel4spanERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QModelIndex QSortFilterProxyModel::buddy(const QModelIndex & index);
extern void C_ZNK21QSortFilterProxyModel5buddyERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QSortFilterProxyModel::hasChildren(const QModelIndex & parent);
extern void C_ZNK21QSortFilterProxyModel11hasChildrenERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QModelIndex QSortFilterProxyModel::mapFromSource(const QModelIndex & sourceIndex);
extern void C_ZNK21QSortFilterProxyModel13mapFromSourceERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  int QSortFilterProxyModel::filterRole();
extern void C_ZNK21QSortFilterProxyModel10filterRoleEv(void* qthis); // 4
  // proto:  int QSortFilterProxyModel::sortRole();
extern void C_ZNK21QSortFilterProxyModel8sortRoleEv(void* qthis); // 4
  // proto:  Qt::SortOrder QSortFilterProxyModel::sortOrder();
extern void C_ZNK21QSortFilterProxyModel9sortOrderEv(void* qthis); // 4
  // proto:  bool QSortFilterProxyModel::setData(const QModelIndex & index, const QVariant & value, int role);
extern void C_ZN21QSortFilterProxyModel7setDataERK11QModelIndexRK8QVarianti(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  void QSortFilterProxyModel::setFilterWildcard(const QString & pattern);
extern void C_ZN21QSortFilterProxyModel17setFilterWildcardERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QSortFilterProxyModel::setSortLocaleAware(bool on);
extern void C_ZN21QSortFilterProxyModel18setSortLocaleAwareEb(void* qthis, bool arg0); // 4
  // proto:  int QSortFilterProxyModel::columnCount(const QModelIndex & parent);
extern void C_ZNK21QSortFilterProxyModel11columnCountERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QSortFilterProxyModel::~QSortFilterProxyModel();
extern void C_ZN21QSortFilterProxyModelD2Ev(void* qthis); // 4
  // proto:  QStringList QSortFilterProxyModel::mimeTypes();
extern void C_ZNK21QSortFilterProxyModel9mimeTypesEv(void* qthis); // 4
  // proto:  QRegExp QSortFilterProxyModel::filterRegExp();
extern void C_ZNK21QSortFilterProxyModel12filterRegExpEv(void* qthis); // 4
  // proto:  int QSortFilterProxyModel::rowCount(const QModelIndex & parent);
extern void C_ZNK21QSortFilterProxyModel8rowCountERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QSortFilterProxyModel::insertRows(int row, int count, const QModelIndex & parent);
extern void C_ZN21QSortFilterProxyModel10insertRowsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QSortFilterProxyModel::setSortRole(int role);
extern void C_ZN21QSortFilterProxyModel11setSortRoleEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QSortFilterProxyModel::dynamicSortFilter();
extern void C_ZNK21QSortFilterProxyModel17dynamicSortFilterEv(void* qthis); // 4
  // proto:  void QSortFilterProxyModel::setFilterFixedString(const QString & pattern);
extern void C_ZN21QSortFilterProxyModel20setFilterFixedStringERK7QString(void* qthis, void* arg0); // 4
  // proto:  QVariant QSortFilterProxyModel::data(const QModelIndex & index, int role);
extern void C_ZNK21QSortFilterProxyModel4dataERK11QModelIndexi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  int QSortFilterProxyModel::filterKeyColumn();
extern void C_ZNK21QSortFilterProxyModel15filterKeyColumnEv(void* qthis); // 4
  // proto:  void QSortFilterProxyModel::setFilterRegExp(const QString & pattern);
extern void C_ZN21QSortFilterProxyModel15setFilterRegExpERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QSortFilterProxyModel::setFilterRegExp(const QRegExp & regExp);
extern void C_ZN21QSortFilterProxyModel15setFilterRegExpERK7QRegExp(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QSortFilterProxyModel::metaObject();
extern void C_ZNK21QSortFilterProxyModel10metaObjectEv(void* qthis); // 4
  // proto:  QItemSelection QSortFilterProxyModel::mapSelectionFromSource(const QItemSelection & sourceSelection);
extern void C_ZNK21QSortFilterProxyModel22mapSelectionFromSourceERK14QItemSelection(void* qthis, void* arg0); // 4
  // proto:  void QSortFilterProxyModel::clear();
extern void C_ZN21QSortFilterProxyModel5clearEv(void* qthis); // 4
  // proto:  Qt::CaseSensitivity QSortFilterProxyModel::sortCaseSensitivity();
extern void C_ZNK21QSortFilterProxyModel19sortCaseSensitivityEv(void* qthis); // 4
  // proto:  QModelIndex QSortFilterProxyModel::mapToSource(const QModelIndex & proxyIndex);
extern void C_ZNK21QSortFilterProxyModel11mapToSourceERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  Qt::ItemFlags QSortFilterProxyModel::flags(const QModelIndex & index);
extern void C_ZNK21QSortFilterProxyModel5flagsERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QSortFilterProxyModel::removeRows(int row, int count, const QModelIndex & parent);
extern void C_ZN21QSortFilterProxyModel10removeRowsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  bool QSortFilterProxyModel::insertColumns(int column, int count, const QModelIndex & parent);
extern void C_ZN21QSortFilterProxyModel13insertColumnsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  bool QSortFilterProxyModel::isSortLocaleAware();
extern void C_ZNK21QSortFilterProxyModel17isSortLocaleAwareEv(void* qthis); // 4
  // proto:  void QSortFilterProxyModel::fetchMore(const QModelIndex & parent);
extern void C_ZN21QSortFilterProxyModel9fetchMoreERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QSortFilterProxyModel::setFilterRole(int role);
extern void C_ZN21QSortFilterProxyModel13setFilterRoleEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QSortFilterProxyModel::removeColumns(int column, int count, const QModelIndex & parent);
extern void C_ZN21QSortFilterProxyModel13removeColumnsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QSortFilterProxyModel::QSortFilterProxyModel(QObject * parent);
extern void* C_ZN21QSortFilterProxyModelC2EP7QObject(void* arg0); // 3
  // proto:  Qt::DropActions QSortFilterProxyModel::supportedDropActions();
extern void C_ZNK21QSortFilterProxyModel20supportedDropActionsEv(void* qthis); // 4
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

// class sizeof(QSortFilterProxyModel)=1
type QSortFilterProxyModel struct {
  /*qbase*/ QAbstractProxyModel;
  qclsinst unsafe.Pointer /* *C.void */;
}

// sortColumn()
func (this *QSortFilterProxyModel) sortColumn(args ...interface{}) () {
  // sortColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel10sortColumnEv
    // invoke: int sortColumn()
    var ret = C.C_ZNK21QSortFilterProxyModel10sortColumnEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "sortColumn", args)
  }

}

// invalidate()
func (this *QSortFilterProxyModel) invalidate(args ...interface{}) () {
  // invalidate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel10invalidateEv
    // invoke: void invalidate()
    C.C_ZN21QSortFilterProxyModel10invalidateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "invalidate", args)
  }

}

// setDynamicSortFilter(_Bool)
func (this *QSortFilterProxyModel) setDynamicSortFilter(args ...interface{}) () {
  // setDynamicSortFilter(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel20setDynamicSortFilterEb
    // invoke: void setDynamicSortFilter(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN21QSortFilterProxyModel20setDynamicSortFilterEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setDynamicSortFilter", args)
  }

}

// canFetchMore(const class QModelIndex &)
func (this *QSortFilterProxyModel) canFetchMore(args ...interface{}) () {
  // canFetchMore(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel12canFetchMoreERK11QModelIndex
    // invoke: bool canFetchMore(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK21QSortFilterProxyModel12canFetchMoreERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "canFetchMore", args)
  }

}

// mapSelectionToSource(const class QItemSelection &)
func (this *QSortFilterProxyModel) mapSelectionToSource(args ...interface{}) () {
  // mapSelectionToSource(const class QItemSelection &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QItemSelection{}) // "const QItemSelection &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel20mapSelectionToSourceERK14QItemSelection
    // invoke: QItemSelection mapSelectionToSource(const class QItemSelection &)
    var arg0 = args[0].(QItemSelection).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK21QSortFilterProxyModel20mapSelectionToSourceERK14QItemSelection(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "mapSelectionToSource", args)
  }

}

// setSourceModel(class QAbstractItemModel *)
func (this *QSortFilterProxyModel) setSourceModel(args ...interface{}) () {
  // setSourceModel(class QAbstractItemModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemModel{}) // "QAbstractItemModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel14setSourceModelEP18QAbstractItemModel
    // invoke: void setSourceModel(class QAbstractItemModel *)
    var arg0 = args[0].(QAbstractItemModel).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN21QSortFilterProxyModel14setSourceModelEP18QAbstractItemModel(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setSourceModel", args)
  }

}

// sibling(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) sibling(args ...interface{}) () {
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
    // invoke: _ZNK21QSortFilterProxyModel7siblingEiiRK11QModelIndex
    // invoke: QModelIndex sibling(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    var ret = C.C_ZNK21QSortFilterProxyModel7siblingEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "sibling", args)
  }

}

// filterCaseSensitivity()
func (this *QSortFilterProxyModel) filterCaseSensitivity(args ...interface{}) () {
  // filterCaseSensitivity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel21filterCaseSensitivityEv
    // invoke: Qt::CaseSensitivity filterCaseSensitivity()
    C.C_ZNK21QSortFilterProxyModel21filterCaseSensitivityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "filterCaseSensitivity", args)
  }

}

// parent(const class QModelIndex &)
func (this *QSortFilterProxyModel) parent(args ...interface{}) () {
  // parent(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel6parentERK11QModelIndex
    // invoke: QModelIndex parent(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK21QSortFilterProxyModel6parentERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "parent", args)
  }

}

// index(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) index(args ...interface{}) () {
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
    // invoke: _ZNK21QSortFilterProxyModel5indexEiiRK11QModelIndex
    // invoke: QModelIndex index(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    var ret = C.C_ZNK21QSortFilterProxyModel5indexEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "index", args)
  }

}

// setFilterKeyColumn(int)
func (this *QSortFilterProxyModel) setFilterKeyColumn(args ...interface{}) () {
  // setFilterKeyColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel18setFilterKeyColumnEi
    // invoke: void setFilterKeyColumn(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN21QSortFilterProxyModel18setFilterKeyColumnEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setFilterKeyColumn", args)
  }

}

// span(const class QModelIndex &)
func (this *QSortFilterProxyModel) span(args ...interface{}) () {
  // span(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel4spanERK11QModelIndex
    // invoke: QSize span(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK21QSortFilterProxyModel4spanERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "span", args)
  }

}

// buddy(const class QModelIndex &)
func (this *QSortFilterProxyModel) buddy(args ...interface{}) () {
  // buddy(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel5buddyERK11QModelIndex
    // invoke: QModelIndex buddy(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK21QSortFilterProxyModel5buddyERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "buddy", args)
  }

}

// hasChildren(const class QModelIndex &)
func (this *QSortFilterProxyModel) hasChildren(args ...interface{}) () {
  // hasChildren(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel11hasChildrenERK11QModelIndex
    // invoke: bool hasChildren(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK21QSortFilterProxyModel11hasChildrenERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "hasChildren", args)
  }

}

// mapFromSource(const class QModelIndex &)
func (this *QSortFilterProxyModel) mapFromSource(args ...interface{}) () {
  // mapFromSource(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel13mapFromSourceERK11QModelIndex
    // invoke: QModelIndex mapFromSource(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK21QSortFilterProxyModel13mapFromSourceERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "mapFromSource", args)
  }

}

// filterRole()
func (this *QSortFilterProxyModel) filterRole(args ...interface{}) () {
  // filterRole()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel10filterRoleEv
    // invoke: int filterRole()
    var ret = C.C_ZNK21QSortFilterProxyModel10filterRoleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "filterRole", args)
  }

}

// sortRole()
func (this *QSortFilterProxyModel) sortRole(args ...interface{}) () {
  // sortRole()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel8sortRoleEv
    // invoke: int sortRole()
    var ret = C.C_ZNK21QSortFilterProxyModel8sortRoleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "sortRole", args)
  }

}

// sortOrder()
func (this *QSortFilterProxyModel) sortOrder(args ...interface{}) () {
  // sortOrder()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel9sortOrderEv
    // invoke: Qt::SortOrder sortOrder()
    C.C_ZNK21QSortFilterProxyModel9sortOrderEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "sortOrder", args)
  }

}

// setData(const class QModelIndex &, const class QVariant &, int)
func (this *QSortFilterProxyModel) setData(args ...interface{}) () {
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
    // invoke: _ZN21QSortFilterProxyModel7setDataERK11QModelIndexRK8QVarianti
    // invoke: bool setData(const class QModelIndex &, const class QVariant &, int)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var ret = C.C_ZN21QSortFilterProxyModel7setDataERK11QModelIndexRK8QVarianti(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setData", args)
  }

}

// setFilterWildcard(const class QString &)
func (this *QSortFilterProxyModel) setFilterWildcard(args ...interface{}) () {
  // setFilterWildcard(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel17setFilterWildcardERK7QString
    // invoke: void setFilterWildcard(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN21QSortFilterProxyModel17setFilterWildcardERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setFilterWildcard", args)
  }

}

// setSortLocaleAware(_Bool)
func (this *QSortFilterProxyModel) setSortLocaleAware(args ...interface{}) () {
  // setSortLocaleAware(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel18setSortLocaleAwareEb
    // invoke: void setSortLocaleAware(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN21QSortFilterProxyModel18setSortLocaleAwareEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setSortLocaleAware", args)
  }

}

// columnCount(const class QModelIndex &)
func (this *QSortFilterProxyModel) columnCount(args ...interface{}) () {
  // columnCount(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel11columnCountERK11QModelIndex
    // invoke: int columnCount(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK21QSortFilterProxyModel11columnCountERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "columnCount", args)
  }

}

// ~QSortFilterProxyModel()
func (this *QSortFilterProxyModel) FreeQSortFilterProxyModel(args ...interface{}) () {
  // ~QSortFilterProxyModel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModelD0Ev
    // invoke: void ~QSortFilterProxyModel()
    C.C_ZN21QSortFilterProxyModelD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "~QSortFilterProxyModel", args)
  }

}

// mimeTypes()
func (this *QSortFilterProxyModel) mimeTypes(args ...interface{}) () {
  // mimeTypes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel9mimeTypesEv
    // invoke: QStringList mimeTypes()
    C.C_ZNK21QSortFilterProxyModel9mimeTypesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "mimeTypes", args)
  }

}

// filterRegExp()
func (this *QSortFilterProxyModel) filterRegExp(args ...interface{}) () {
  // filterRegExp()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel12filterRegExpEv
    // invoke: QRegExp filterRegExp()
    var ret = C.C_ZNK21QSortFilterProxyModel12filterRegExpEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "filterRegExp", args)
  }

}

// rowCount(const class QModelIndex &)
func (this *QSortFilterProxyModel) rowCount(args ...interface{}) () {
  // rowCount(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel8rowCountERK11QModelIndex
    // invoke: int rowCount(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK21QSortFilterProxyModel8rowCountERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "rowCount", args)
  }

}

// insertRows(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) insertRows(args ...interface{}) () {
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
    // invoke: _ZN21QSortFilterProxyModel10insertRowsEiiRK11QModelIndex
    // invoke: bool insertRows(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    var ret = C.C_ZN21QSortFilterProxyModel10insertRowsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "insertRows", args)
  }

}

// setSortRole(int)
func (this *QSortFilterProxyModel) setSortRole(args ...interface{}) () {
  // setSortRole(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel11setSortRoleEi
    // invoke: void setSortRole(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN21QSortFilterProxyModel11setSortRoleEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setSortRole", args)
  }

}

// dynamicSortFilter()
func (this *QSortFilterProxyModel) dynamicSortFilter(args ...interface{}) () {
  // dynamicSortFilter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel17dynamicSortFilterEv
    // invoke: bool dynamicSortFilter()
    var ret = C.C_ZNK21QSortFilterProxyModel17dynamicSortFilterEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "dynamicSortFilter", args)
  }

}

// setFilterFixedString(const class QString &)
func (this *QSortFilterProxyModel) setFilterFixedString(args ...interface{}) () {
  // setFilterFixedString(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel20setFilterFixedStringERK7QString
    // invoke: void setFilterFixedString(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN21QSortFilterProxyModel20setFilterFixedStringERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setFilterFixedString", args)
  }

}

// data(const class QModelIndex &, int)
func (this *QSortFilterProxyModel) data(args ...interface{}) () {
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
    // invoke: _ZNK21QSortFilterProxyModel4dataERK11QModelIndexi
    // invoke: QVariant data(const class QModelIndex &, int)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK21QSortFilterProxyModel4dataERK11QModelIndexi(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "data", args)
  }

}

// filterKeyColumn()
func (this *QSortFilterProxyModel) filterKeyColumn(args ...interface{}) () {
  // filterKeyColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel15filterKeyColumnEv
    // invoke: int filterKeyColumn()
    var ret = C.C_ZNK21QSortFilterProxyModel15filterKeyColumnEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "filterKeyColumn", args)
  }

}

// setFilterRegExp(const class QString &)
func (this *QSortFilterProxyModel) setFilterRegExp(args ...interface{}) () {
  // setFilterRegExp(const class QString &)
  // setFilterRegExp(const class QRegExp &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel15setFilterRegExpERK7QString
    // invoke: void setFilterRegExp(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN21QSortFilterProxyModel15setFilterRegExpERK7QString(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN21QSortFilterProxyModel15setFilterRegExpERK7QRegExp
    // invoke: void setFilterRegExp(const class QRegExp &)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN21QSortFilterProxyModel15setFilterRegExpERK7QRegExp(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setFilterRegExp", args)
  }

}

// metaObject()
func (this *QSortFilterProxyModel) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK21QSortFilterProxyModel10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "metaObject", args)
  }

}

// mapSelectionFromSource(const class QItemSelection &)
func (this *QSortFilterProxyModel) mapSelectionFromSource(args ...interface{}) () {
  // mapSelectionFromSource(const class QItemSelection &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QItemSelection{}) // "const QItemSelection &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel22mapSelectionFromSourceERK14QItemSelection
    // invoke: QItemSelection mapSelectionFromSource(const class QItemSelection &)
    var arg0 = args[0].(QItemSelection).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK21QSortFilterProxyModel22mapSelectionFromSourceERK14QItemSelection(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "mapSelectionFromSource", args)
  }

}

// clear()
func (this *QSortFilterProxyModel) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel5clearEv
    // invoke: void clear()
    C.C_ZN21QSortFilterProxyModel5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "clear", args)
  }

}

// sortCaseSensitivity()
func (this *QSortFilterProxyModel) sortCaseSensitivity(args ...interface{}) () {
  // sortCaseSensitivity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel19sortCaseSensitivityEv
    // invoke: Qt::CaseSensitivity sortCaseSensitivity()
    C.C_ZNK21QSortFilterProxyModel19sortCaseSensitivityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "sortCaseSensitivity", args)
  }

}

// mapToSource(const class QModelIndex &)
func (this *QSortFilterProxyModel) mapToSource(args ...interface{}) () {
  // mapToSource(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel11mapToSourceERK11QModelIndex
    // invoke: QModelIndex mapToSource(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK21QSortFilterProxyModel11mapToSourceERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "mapToSource", args)
  }

}

// flags(const class QModelIndex &)
func (this *QSortFilterProxyModel) flags(args ...interface{}) () {
  // flags(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel5flagsERK11QModelIndex
    // invoke: Qt::ItemFlags flags(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK21QSortFilterProxyModel5flagsERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "flags", args)
  }

}

// removeRows(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) removeRows(args ...interface{}) () {
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
    // invoke: _ZN21QSortFilterProxyModel10removeRowsEiiRK11QModelIndex
    // invoke: bool removeRows(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    var ret = C.C_ZN21QSortFilterProxyModel10removeRowsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "removeRows", args)
  }

}

// insertColumns(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) insertColumns(args ...interface{}) () {
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
    // invoke: _ZN21QSortFilterProxyModel13insertColumnsEiiRK11QModelIndex
    // invoke: bool insertColumns(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    var ret = C.C_ZN21QSortFilterProxyModel13insertColumnsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "insertColumns", args)
  }

}

// isSortLocaleAware()
func (this *QSortFilterProxyModel) isSortLocaleAware(args ...interface{}) () {
  // isSortLocaleAware()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel17isSortLocaleAwareEv
    // invoke: bool isSortLocaleAware()
    var ret = C.C_ZNK21QSortFilterProxyModel17isSortLocaleAwareEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "isSortLocaleAware", args)
  }

}

// fetchMore(const class QModelIndex &)
func (this *QSortFilterProxyModel) fetchMore(args ...interface{}) () {
  // fetchMore(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel9fetchMoreERK11QModelIndex
    // invoke: void fetchMore(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN21QSortFilterProxyModel9fetchMoreERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "fetchMore", args)
  }

}

// setFilterRole(int)
func (this *QSortFilterProxyModel) setFilterRole(args ...interface{}) () {
  // setFilterRole(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel13setFilterRoleEi
    // invoke: void setFilterRole(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN21QSortFilterProxyModel13setFilterRoleEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setFilterRole", args)
  }

}

// removeColumns(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) removeColumns(args ...interface{}) () {
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
    // invoke: _ZN21QSortFilterProxyModel13removeColumnsEiiRK11QModelIndex
    // invoke: bool removeColumns(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    var ret = C.C_ZN21QSortFilterProxyModel13removeColumnsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "removeColumns", args)
  }

}

// QSortFilterProxyModel(class QObject *)
func NewQSortFilterProxyModel(args ...interface{}) *QSortFilterProxyModel {
  // QSortFilterProxyModel(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModelC1EP7QObject
    // invoke: void QSortFilterProxyModel(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN21QSortFilterProxyModelC2EP7QObject(arg0)
    return &QSortFilterProxyModel{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "QSortFilterProxyModel", args)
  }

  return nil // QSortFilterProxyModel{qclsinst:qthis}
}

// supportedDropActions()
func (this *QSortFilterProxyModel) supportedDropActions(args ...interface{}) () {
  // supportedDropActions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel20supportedDropActionsEv
    // invoke: Qt::DropActions supportedDropActions()
    C.C_ZNK21QSortFilterProxyModel20supportedDropActionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "supportedDropActions", args)
  }

}

// <= body block end

