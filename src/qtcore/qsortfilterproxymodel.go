package qtcore
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
import "runtime"
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
extern int32_t C_ZNK21QSortFilterProxyModel10sortColumnEv(void* qthis); // 4
  // proto:  void QSortFilterProxyModel::invalidate();
extern void C_ZN21QSortFilterProxyModel10invalidateEv(void* qthis); // 4
  // proto:  void QSortFilterProxyModel::setDynamicSortFilter(bool enable);
extern void C_ZN21QSortFilterProxyModel20setDynamicSortFilterEb(void* qthis, bool arg0); // 4
  // proto:  bool QSortFilterProxyModel::canFetchMore(const QModelIndex & parent);
extern bool C_ZNK21QSortFilterProxyModel12canFetchMoreERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QItemSelection QSortFilterProxyModel::mapSelectionToSource(const QItemSelection & proxySelection);
extern void* C_ZNK21QSortFilterProxyModel20mapSelectionToSourceERK14QItemSelection(void* qthis, void* arg0); // 4
  // proto:  void QSortFilterProxyModel::setSourceModel(QAbstractItemModel * sourceModel);
extern void C_ZN21QSortFilterProxyModel14setSourceModelEP18QAbstractItemModel(void* qthis, void* arg0); // 4
  // proto:  QModelIndex QSortFilterProxyModel::sibling(int row, int column, const QModelIndex & idx);
extern void* C_ZNK21QSortFilterProxyModel7siblingEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  Qt::CaseSensitivity QSortFilterProxyModel::filterCaseSensitivity();
extern void C_ZNK21QSortFilterProxyModel21filterCaseSensitivityEv(void* qthis); // 4
  // proto:  QModelIndex QSortFilterProxyModel::parent(const QModelIndex & child);
extern void* C_ZNK21QSortFilterProxyModel6parentERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QModelIndex QSortFilterProxyModel::index(int row, int column, const QModelIndex & parent);
extern void* C_ZNK21QSortFilterProxyModel5indexEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QSortFilterProxyModel::setFilterKeyColumn(int column);
extern void C_ZN21QSortFilterProxyModel18setFilterKeyColumnEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QSortFilterProxyModel::span(const QModelIndex & index);
extern void* C_ZNK21QSortFilterProxyModel4spanERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QModelIndex QSortFilterProxyModel::buddy(const QModelIndex & index);
extern void* C_ZNK21QSortFilterProxyModel5buddyERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QSortFilterProxyModel::hasChildren(const QModelIndex & parent);
extern bool C_ZNK21QSortFilterProxyModel11hasChildrenERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QModelIndex QSortFilterProxyModel::mapFromSource(const QModelIndex & sourceIndex);
extern void* C_ZNK21QSortFilterProxyModel13mapFromSourceERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  int QSortFilterProxyModel::filterRole();
extern int32_t C_ZNK21QSortFilterProxyModel10filterRoleEv(void* qthis); // 4
  // proto:  int QSortFilterProxyModel::sortRole();
extern int32_t C_ZNK21QSortFilterProxyModel8sortRoleEv(void* qthis); // 4
  // proto:  Qt::SortOrder QSortFilterProxyModel::sortOrder();
extern void C_ZNK21QSortFilterProxyModel9sortOrderEv(void* qthis); // 4
  // proto:  bool QSortFilterProxyModel::setData(const QModelIndex & index, const QVariant & value, int role);
extern bool C_ZN21QSortFilterProxyModel7setDataERK11QModelIndexRK8QVarianti(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  void QSortFilterProxyModel::setFilterWildcard(const QString & pattern);
extern void C_ZN21QSortFilterProxyModel17setFilterWildcardERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QSortFilterProxyModel::setSortLocaleAware(bool on);
extern void C_ZN21QSortFilterProxyModel18setSortLocaleAwareEb(void* qthis, bool arg0); // 4
  // proto:  int QSortFilterProxyModel::columnCount(const QModelIndex & parent);
extern int32_t C_ZNK21QSortFilterProxyModel11columnCountERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QSortFilterProxyModel::~QSortFilterProxyModel();
extern void C_ZN21QSortFilterProxyModelD2Ev(void* qthis); // 4
  // proto:  QStringList QSortFilterProxyModel::mimeTypes();
extern void C_ZNK21QSortFilterProxyModel9mimeTypesEv(void* qthis); // 4
  // proto:  QRegExp QSortFilterProxyModel::filterRegExp();
extern void* C_ZNK21QSortFilterProxyModel12filterRegExpEv(void* qthis); // 4
  // proto:  int QSortFilterProxyModel::rowCount(const QModelIndex & parent);
extern int32_t C_ZNK21QSortFilterProxyModel8rowCountERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QSortFilterProxyModel::insertRows(int row, int count, const QModelIndex & parent);
extern bool C_ZN21QSortFilterProxyModel10insertRowsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QSortFilterProxyModel::setSortRole(int role);
extern void C_ZN21QSortFilterProxyModel11setSortRoleEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QSortFilterProxyModel::dynamicSortFilter();
extern bool C_ZNK21QSortFilterProxyModel17dynamicSortFilterEv(void* qthis); // 4
  // proto:  void QSortFilterProxyModel::setFilterFixedString(const QString & pattern);
extern void C_ZN21QSortFilterProxyModel20setFilterFixedStringERK7QString(void* qthis, void* arg0); // 4
  // proto:  QVariant QSortFilterProxyModel::data(const QModelIndex & index, int role);
extern void* C_ZNK21QSortFilterProxyModel4dataERK11QModelIndexi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  int QSortFilterProxyModel::filterKeyColumn();
extern int32_t C_ZNK21QSortFilterProxyModel15filterKeyColumnEv(void* qthis); // 4
  // proto:  void QSortFilterProxyModel::setFilterRegExp(const QString & pattern);
extern void C_ZN21QSortFilterProxyModel15setFilterRegExpERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QSortFilterProxyModel::setFilterRegExp(const QRegExp & regExp);
extern void C_ZN21QSortFilterProxyModel15setFilterRegExpERK7QRegExp(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QSortFilterProxyModel::metaObject();
extern void C_ZNK21QSortFilterProxyModel10metaObjectEv(void* qthis); // 4
  // proto:  QItemSelection QSortFilterProxyModel::mapSelectionFromSource(const QItemSelection & sourceSelection);
extern void* C_ZNK21QSortFilterProxyModel22mapSelectionFromSourceERK14QItemSelection(void* qthis, void* arg0); // 4
  // proto:  void QSortFilterProxyModel::clear();
extern void C_ZN21QSortFilterProxyModel5clearEv(void* qthis); // 4
  // proto:  Qt::CaseSensitivity QSortFilterProxyModel::sortCaseSensitivity();
extern void C_ZNK21QSortFilterProxyModel19sortCaseSensitivityEv(void* qthis); // 4
  // proto:  QModelIndex QSortFilterProxyModel::mapToSource(const QModelIndex & proxyIndex);
extern void* C_ZNK21QSortFilterProxyModel11mapToSourceERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  Qt::ItemFlags QSortFilterProxyModel::flags(const QModelIndex & index);
extern void C_ZNK21QSortFilterProxyModel5flagsERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QSortFilterProxyModel::removeRows(int row, int count, const QModelIndex & parent);
extern bool C_ZN21QSortFilterProxyModel10removeRowsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  bool QSortFilterProxyModel::insertColumns(int column, int count, const QModelIndex & parent);
extern bool C_ZN21QSortFilterProxyModel13insertColumnsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  bool QSortFilterProxyModel::isSortLocaleAware();
extern bool C_ZNK21QSortFilterProxyModel17isSortLocaleAwareEv(void* qthis); // 4
  // proto:  void QSortFilterProxyModel::fetchMore(const QModelIndex & parent);
extern void C_ZN21QSortFilterProxyModel9fetchMoreERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QSortFilterProxyModel::setFilterRole(int role);
extern void C_ZN21QSortFilterProxyModel13setFilterRoleEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QSortFilterProxyModel::removeColumns(int column, int count, const QModelIndex & parent);
extern bool C_ZN21QSortFilterProxyModel13removeColumnsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
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
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QSortFilterProxyModel)=1
type QSortFilterProxyModel struct {
  /*qbase*/ QAbstractProxyModel;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// sortColumn()
func (this *QSortFilterProxyModel) SortColumn(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK21QSortFilterProxyModel10sortColumnEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "sortColumn", args)
  }

  return
}

// invalidate()
func (this *QSortFilterProxyModel) Invalidate(args ...interface{}) () {
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
    C.C_ZN21QSortFilterProxyModel10invalidateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "invalidate", args)
  }

  return
}

// setDynamicSortFilter(_Bool)
func (this *QSortFilterProxyModel) SetDynamicSortFilter(args ...interface{}) () {
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
    C.C_ZN21QSortFilterProxyModel20setDynamicSortFilterEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setDynamicSortFilter", args)
  }

  return
}

// canFetchMore(const class QModelIndex &)
func (this *QSortFilterProxyModel) CanFetchMore(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK21QSortFilterProxyModel12canFetchMoreERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "canFetchMore", args)
  }

  return
}

// mapSelectionToSource(const class QItemSelection &)
func (this *QSortFilterProxyModel) MapSelectionToSource(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QItemSelection).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK21QSortFilterProxyModel20mapSelectionToSourceERK14QItemSelection(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QItemSelection{}) // "QItemSelection"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "mapSelectionToSource", args)
  }

  return
}

// setSourceModel(class QAbstractItemModel *)
func (this *QSortFilterProxyModel) SetSourceModel(args ...interface{}) () {
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
    var arg0 = args[0].(*QAbstractItemModel).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN21QSortFilterProxyModel14setSourceModelEP18QAbstractItemModel(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setSourceModel", args)
  }

  return
}

// sibling(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) Sibling(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK21QSortFilterProxyModel7siblingEiiRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "sibling", args)
  }

  return
}

// filterCaseSensitivity()
func (this *QSortFilterProxyModel) FilterCaseSensitivity(args ...interface{}) () {
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
    C.C_ZNK21QSortFilterProxyModel21filterCaseSensitivityEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "filterCaseSensitivity", args)
  }

  return
}

// parent(const class QModelIndex &)
func (this *QSortFilterProxyModel) Parent(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK21QSortFilterProxyModel6parentERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "parent", args)
  }

  return
}

// index(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) Index(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK21QSortFilterProxyModel5indexEiiRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "index", args)
  }

  return
}

// setFilterKeyColumn(int)
func (this *QSortFilterProxyModel) SetFilterKeyColumn(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN21QSortFilterProxyModel18setFilterKeyColumnEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setFilterKeyColumn", args)
  }

  return
}

// span(const class QModelIndex &)
func (this *QSortFilterProxyModel) Span(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK21QSortFilterProxyModel4spanERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "span", args)
  }

  return
}

// buddy(const class QModelIndex &)
func (this *QSortFilterProxyModel) Buddy(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK21QSortFilterProxyModel5buddyERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "buddy", args)
  }

  return
}

// hasChildren(const class QModelIndex &)
func (this *QSortFilterProxyModel) HasChildren(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK21QSortFilterProxyModel11hasChildrenERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "hasChildren", args)
  }

  return
}

// mapFromSource(const class QModelIndex &)
func (this *QSortFilterProxyModel) MapFromSource(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK21QSortFilterProxyModel13mapFromSourceERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "mapFromSource", args)
  }

  return
}

// filterRole()
func (this *QSortFilterProxyModel) FilterRole(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK21QSortFilterProxyModel10filterRoleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "filterRole", args)
  }

  return
}

// sortRole()
func (this *QSortFilterProxyModel) SortRole(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK21QSortFilterProxyModel8sortRoleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "sortRole", args)
  }

  return
}

// sortOrder()
func (this *QSortFilterProxyModel) SortOrder(args ...interface{}) () {
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
    C.C_ZNK21QSortFilterProxyModel9sortOrderEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "sortOrder", args)
  }

  return
}

// setData(const class QModelIndex &, const class QVariant &, int)
func (this *QSortFilterProxyModel) SetData(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVariant).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN21QSortFilterProxyModel7setDataERK11QModelIndexRK8QVarianti(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setData", args)
  }

  return
}

// setFilterWildcard(const class QString &)
func (this *QSortFilterProxyModel) SetFilterWildcard(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN21QSortFilterProxyModel17setFilterWildcardERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setFilterWildcard", args)
  }

  return
}

// setSortLocaleAware(_Bool)
func (this *QSortFilterProxyModel) SetSortLocaleAware(args ...interface{}) () {
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
    C.C_ZN21QSortFilterProxyModel18setSortLocaleAwareEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setSortLocaleAware", args)
  }

  return
}

// columnCount(const class QModelIndex &)
func (this *QSortFilterProxyModel) ColumnCount(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK21QSortFilterProxyModel11columnCountERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "columnCount", args)
  }

  return
}

// ~QSortFilterProxyModel()
func (this *QSortFilterProxyModel) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN21QSortFilterProxyModelD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "~QSortFilterProxyModel", args)
  }

  return
}

// mimeTypes()
func (this *QSortFilterProxyModel) MimeTypes(args ...interface{}) () {
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
    C.C_ZNK21QSortFilterProxyModel9mimeTypesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "mimeTypes", args)
  }

  return
}

// filterRegExp()
func (this *QSortFilterProxyModel) FilterRegExp(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK21QSortFilterProxyModel12filterRegExpEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRegExp{}) // "QRegExp"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "filterRegExp", args)
  }

  return
}

// rowCount(const class QModelIndex &)
func (this *QSortFilterProxyModel) RowCount(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK21QSortFilterProxyModel8rowCountERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "rowCount", args)
  }

  return
}

// insertRows(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) InsertRows(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN21QSortFilterProxyModel10insertRowsEiiRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "insertRows", args)
  }

  return
}

// setSortRole(int)
func (this *QSortFilterProxyModel) SetSortRole(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN21QSortFilterProxyModel11setSortRoleEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setSortRole", args)
  }

  return
}

// dynamicSortFilter()
func (this *QSortFilterProxyModel) DynamicSortFilter(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK21QSortFilterProxyModel17dynamicSortFilterEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "dynamicSortFilter", args)
  }

  return
}

// setFilterFixedString(const class QString &)
func (this *QSortFilterProxyModel) SetFilterFixedString(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN21QSortFilterProxyModel20setFilterFixedStringERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setFilterFixedString", args)
  }

  return
}

// data(const class QModelIndex &, int)
func (this *QSortFilterProxyModel) Data(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK21QSortFilterProxyModel4dataERK11QModelIndexi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVariant{}) // "QVariant"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "data", args)
  }

  return
}

// filterKeyColumn()
func (this *QSortFilterProxyModel) FilterKeyColumn(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK21QSortFilterProxyModel15filterKeyColumnEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "filterKeyColumn", args)
  }

  return
}

// setFilterRegExp(const class QString &)
func (this *QSortFilterProxyModel) SetFilterRegExp(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN21QSortFilterProxyModel15setFilterRegExpERK7QString(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN21QSortFilterProxyModel15setFilterRegExpERK7QRegExp
    // invoke: void setFilterRegExp(const class QRegExp &)
    var arg0 = args[0].(*QRegExp).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN21QSortFilterProxyModel15setFilterRegExpERK7QRegExp(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setFilterRegExp", args)
  }

  return
}

// metaObject()
func (this *QSortFilterProxyModel) MetaObject(args ...interface{}) () {
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
    C.C_ZNK21QSortFilterProxyModel10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "metaObject", args)
  }

  return
}

// mapSelectionFromSource(const class QItemSelection &)
func (this *QSortFilterProxyModel) MapSelectionFromSource(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QItemSelection).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK21QSortFilterProxyModel22mapSelectionFromSourceERK14QItemSelection(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QItemSelection{}) // "QItemSelection"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "mapSelectionFromSource", args)
  }

  return
}

// clear()
func (this *QSortFilterProxyModel) Clear(args ...interface{}) () {
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
    C.C_ZN21QSortFilterProxyModel5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "clear", args)
  }

  return
}

// sortCaseSensitivity()
func (this *QSortFilterProxyModel) SortCaseSensitivity(args ...interface{}) () {
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
    C.C_ZNK21QSortFilterProxyModel19sortCaseSensitivityEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "sortCaseSensitivity", args)
  }

  return
}

// mapToSource(const class QModelIndex &)
func (this *QSortFilterProxyModel) MapToSource(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK21QSortFilterProxyModel11mapToSourceERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "mapToSource", args)
  }

  return
}

// flags(const class QModelIndex &)
func (this *QSortFilterProxyModel) Flags(args ...interface{}) () {
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
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK21QSortFilterProxyModel5flagsERK11QModelIndex(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "flags", args)
  }

  return
}

// removeRows(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) RemoveRows(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN21QSortFilterProxyModel10removeRowsEiiRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "removeRows", args)
  }

  return
}

// insertColumns(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) InsertColumns(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN21QSortFilterProxyModel13insertColumnsEiiRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "insertColumns", args)
  }

  return
}

// isSortLocaleAware()
func (this *QSortFilterProxyModel) IsSortLocaleAware(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK21QSortFilterProxyModel17isSortLocaleAwareEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "isSortLocaleAware", args)
  }

  return
}

// fetchMore(const class QModelIndex &)
func (this *QSortFilterProxyModel) FetchMore(args ...interface{}) () {
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
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN21QSortFilterProxyModel9fetchMoreERK11QModelIndex(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "fetchMore", args)
  }

  return
}

// setFilterRole(int)
func (this *QSortFilterProxyModel) SetFilterRole(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN21QSortFilterProxyModel13setFilterRoleEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setFilterRole", args)
  }

  return
}

// removeColumns(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) RemoveColumns(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN21QSortFilterProxyModel13removeColumnsEiiRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "removeColumns", args)
  }

  return
}

// QSortFilterProxyModel(class QObject *)
func GcfreeQSortFilterProxyModel(this *QSortFilterProxyModel) {
  qtrt.UniverseFree(this)
}
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
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN21QSortFilterProxyModelC2EP7QObject(arg0)
    this := &QSortFilterProxyModel{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQSortFilterProxyModel)
    return this
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "QSortFilterProxyModel", args)
  }

  return nil // QSortFilterProxyModel{Qclsinst:qthis}
}

// supportedDropActions()
func (this *QSortFilterProxyModel) SupportedDropActions(args ...interface{}) () {
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
    C.C_ZNK21QSortFilterProxyModel20supportedDropActionsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "supportedDropActions", args)
  }

  return
}

// <= body block end

