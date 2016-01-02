package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
  // proto:  void QSortFilterProxyModel::setFilterRegExp(const QString & pattern);
extern void _ZN21QSortFilterProxyModel15setFilterRegExpERK7QString(void* qthis, void* arg0);
  // proto:  int QSortFilterProxyModel::rowCount(const QModelIndex & parent);
extern void _ZNK21QSortFilterProxyModel8rowCountERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QModelIndex QSortFilterProxyModel::sibling(int row, int column, const QModelIndex & idx);
extern void _ZNK21QSortFilterProxyModel7siblingEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  QSize QSortFilterProxyModel::span(const QModelIndex & index);
extern void _ZNK21QSortFilterProxyModel4spanERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QModelIndex QSortFilterProxyModel::mapFromSource(const QModelIndex & sourceIndex);
extern void _ZNK21QSortFilterProxyModel13mapFromSourceERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QSortFilterProxyModel::setFilterWildcard(const QString & pattern);
extern void _ZN21QSortFilterProxyModel17setFilterWildcardERK7QString(void* qthis, void* arg0);
  // proto:  bool QSortFilterProxyModel::hasChildren(const QModelIndex & parent);
extern void _ZNK21QSortFilterProxyModel11hasChildrenERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QSortFilterProxyModel::setFilterFixedString(const QString & pattern);
extern void _ZN21QSortFilterProxyModel20setFilterFixedStringERK7QString(void* qthis, void* arg0);
  // proto:  bool QSortFilterProxyModel::setData(const QModelIndex & index, const QVariant & value, int role);
extern void _ZN21QSortFilterProxyModel7setDataERK11QModelIndexRK8QVarianti(void* qthis, void* arg0, void* arg1, int arg2);
  // proto:  void QSortFilterProxyModel::setSortRole(int role);
extern void _ZN21QSortFilterProxyModel11setSortRoleEi(void* qthis, int arg0);
  // proto:  QVariant QSortFilterProxyModel::data(const QModelIndex & index, int role);
extern void _ZNK21QSortFilterProxyModel4dataERK11QModelIndexi(void* qthis, void* arg0, int arg1);
  // proto:  void QSortFilterProxyModel::invalidate();
extern void _ZN21QSortFilterProxyModel10invalidateEv(void* qthis);
  // proto:  int QSortFilterProxyModel::sortColumn();
extern void _ZNK21QSortFilterProxyModel10sortColumnEv(void* qthis);
  // proto:  bool QSortFilterProxyModel::insertRows(int row, int count, const QModelIndex & parent);
extern void _ZN21QSortFilterProxyModel10insertRowsEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  int QSortFilterProxyModel::filterKeyColumn();
extern void _ZNK21QSortFilterProxyModel15filterKeyColumnEv(void* qthis);
  // proto:  bool QSortFilterProxyModel::canFetchMore(const QModelIndex & parent);
extern void _ZNK21QSortFilterProxyModel12canFetchMoreERK11QModelIndex(void* qthis, void* arg0);
  // proto:  bool QSortFilterProxyModel::isSortLocaleAware();
extern void _ZNK21QSortFilterProxyModel17isSortLocaleAwareEv(void* qthis);
  // proto:  void QSortFilterProxyModel::fetchMore(const QModelIndex & parent);
extern void _ZN21QSortFilterProxyModel9fetchMoreERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QItemSelection QSortFilterProxyModel::mapSelectionFromSource(const QItemSelection & sourceSelection);
extern void _ZNK21QSortFilterProxyModel22mapSelectionFromSourceERK14QItemSelection(void* qthis, void* arg0);
  // proto:  QItemSelection QSortFilterProxyModel::mapSelectionToSource(const QItemSelection & proxySelection);
extern void _ZNK21QSortFilterProxyModel20mapSelectionToSourceERK14QItemSelection(void* qthis, void* arg0);
  // proto:  QStringList QSortFilterProxyModel::mimeTypes();
extern void _ZNK21QSortFilterProxyModel9mimeTypesEv(void* qthis);
  // proto:  QModelIndex QSortFilterProxyModel::buddy(const QModelIndex & index);
extern void _ZNK21QSortFilterProxyModel5buddyERK11QModelIndex(void* qthis, void* arg0);
  // proto:  int QSortFilterProxyModel::filterRole();
extern void _ZNK21QSortFilterProxyModel10filterRoleEv(void* qthis);
  // proto:  void QSortFilterProxyModel::clear();
extern void _ZN21QSortFilterProxyModel5clearEv(void* qthis);
  // proto:  void QSortFilterProxyModel::setFilterKeyColumn(int column);
extern void _ZN21QSortFilterProxyModel18setFilterKeyColumnEi(void* qthis, int arg0);
  // proto:  const QMetaObject * QSortFilterProxyModel::metaObject();
extern void _ZNK21QSortFilterProxyModel10metaObjectEv(void* qthis);
  // proto:  int QSortFilterProxyModel::sortRole();
extern void _ZNK21QSortFilterProxyModel8sortRoleEv(void* qthis);
  // proto:  void QSortFilterProxyModel::setSortLocaleAware(bool on);
extern void _ZN21QSortFilterProxyModel18setSortLocaleAwareEb(void* qthis, bool arg0);
  // proto:  QModelIndex QSortFilterProxyModel::mapToSource(const QModelIndex & proxyIndex);
extern void _ZNK21QSortFilterProxyModel11mapToSourceERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QSortFilterProxyModel::setSourceModel(QAbstractItemModel * sourceModel);
extern void _ZN21QSortFilterProxyModel14setSourceModelEP18QAbstractItemModel(void* qthis, void* arg0);
  // proto:  void QSortFilterProxyModel::QSortFilterProxyModel(const QSortFilterProxyModel & );
extern void* dector_ZN21QSortFilterProxyModelC1ERKS_(void* arg0);
extern void _ZN21QSortFilterProxyModelC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QSortFilterProxyModel::removeColumns(int column, int count, const QModelIndex & parent);
extern void _ZN21QSortFilterProxyModel13removeColumnsEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  void QSortFilterProxyModel::~QSortFilterProxyModel();
extern void _ZN21QSortFilterProxyModelD0Ev(void* qthis);
  // proto:  bool QSortFilterProxyModel::dynamicSortFilter();
extern void _ZNK21QSortFilterProxyModel17dynamicSortFilterEv(void* qthis);
  // proto:  bool QSortFilterProxyModel::insertColumns(int column, int count, const QModelIndex & parent);
extern void _ZN21QSortFilterProxyModel13insertColumnsEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  int QSortFilterProxyModel::columnCount(const QModelIndex & parent);
extern void _ZNK21QSortFilterProxyModel11columnCountERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QSortFilterProxyModel::setFilterRegExp(const QRegExp & regExp);
extern void _ZN21QSortFilterProxyModel15setFilterRegExpERK7QRegExp(void* qthis, void* arg0);
  // proto:  QModelIndex QSortFilterProxyModel::parent(const QModelIndex & child);
extern void _ZNK21QSortFilterProxyModel6parentERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QRegExp QSortFilterProxyModel::filterRegExp();
extern void _ZNK21QSortFilterProxyModel12filterRegExpEv(void* qthis);
  // proto:  void QSortFilterProxyModel::setFilterRole(int role);
extern void _ZN21QSortFilterProxyModel13setFilterRoleEi(void* qthis, int arg0);
  // proto:  void QSortFilterProxyModel::QSortFilterProxyModel(QObject * parent);
extern void* dector_ZN21QSortFilterProxyModelC1EP7QObject(void* arg0);
extern void _ZN21QSortFilterProxyModelC1EP7QObject(void* qthis, void* arg0);
  // proto:  bool QSortFilterProxyModel::removeRows(int row, int count, const QModelIndex & parent);
extern void _ZN21QSortFilterProxyModel10removeRowsEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  QModelIndex QSortFilterProxyModel::index(int row, int column, const QModelIndex & parent);
extern void _ZNK21QSortFilterProxyModel5indexEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  void QSortFilterProxyModel::setDynamicSortFilter(bool enable);
extern void _ZN21QSortFilterProxyModel20setDynamicSortFilterEb(void* qthis, bool arg0);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QSortFilterProxyModel::setFilterRegExp(const QString & pattern);
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
  case 1:
    // invoke: _ZN21QSortFilterProxyModel15setFilterRegExpERK7QRegExp
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setFilterRegExp", args)
  }

}

  // proto:  int QSortFilterProxyModel::rowCount(const QModelIndex & parent);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "rowCount", args)
  }

}

  // proto:  QModelIndex QSortFilterProxyModel::sibling(int row, int column, const QModelIndex & idx);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "sibling", args)
  }

}

  // proto:  QSize QSortFilterProxyModel::span(const QModelIndex & index);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "span", args)
  }

}

  // proto:  QModelIndex QSortFilterProxyModel::mapFromSource(const QModelIndex & sourceIndex);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "mapFromSource", args)
  }

}

  // proto:  void QSortFilterProxyModel::setFilterWildcard(const QString & pattern);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setFilterWildcard", args)
  }

}

  // proto:  bool QSortFilterProxyModel::hasChildren(const QModelIndex & parent);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "hasChildren", args)
  }

}

  // proto:  void QSortFilterProxyModel::setFilterFixedString(const QString & pattern);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setFilterFixedString", args)
  }

}

  // proto:  bool QSortFilterProxyModel::setData(const QModelIndex & index, const QVariant & value, int role);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setData", args)
  }

}

  // proto:  void QSortFilterProxyModel::setSortRole(int role);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setSortRole", args)
  }

}

  // proto:  QVariant QSortFilterProxyModel::data(const QModelIndex & index, int role);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "data", args)
  }

}

  // proto:  void QSortFilterProxyModel::invalidate();
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "invalidate", args)
  }

}

  // proto:  int QSortFilterProxyModel::sortColumn();
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "sortColumn", args)
  }

}

  // proto:  bool QSortFilterProxyModel::insertRows(int row, int count, const QModelIndex & parent);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "insertRows", args)
  }

}

  // proto:  int QSortFilterProxyModel::filterKeyColumn();
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "filterKeyColumn", args)
  }

}

  // proto:  bool QSortFilterProxyModel::canFetchMore(const QModelIndex & parent);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "canFetchMore", args)
  }

}

  // proto:  bool QSortFilterProxyModel::isSortLocaleAware();
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "isSortLocaleAware", args)
  }

}

  // proto:  void QSortFilterProxyModel::fetchMore(const QModelIndex & parent);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "fetchMore", args)
  }

}

  // proto:  QItemSelection QSortFilterProxyModel::mapSelectionFromSource(const QItemSelection & sourceSelection);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "mapSelectionFromSource", args)
  }

}

  // proto:  QItemSelection QSortFilterProxyModel::mapSelectionToSource(const QItemSelection & proxySelection);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "mapSelectionToSource", args)
  }

}

  // proto:  QStringList QSortFilterProxyModel::mimeTypes();
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "mimeTypes", args)
  }

}

  // proto:  QModelIndex QSortFilterProxyModel::buddy(const QModelIndex & index);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "buddy", args)
  }

}

  // proto:  int QSortFilterProxyModel::filterRole();
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "filterRole", args)
  }

}

  // proto:  void QSortFilterProxyModel::clear();
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "clear", args)
  }

}

  // proto:  void QSortFilterProxyModel::setFilterKeyColumn(int column);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setFilterKeyColumn", args)
  }

}

  // proto:  const QMetaObject * QSortFilterProxyModel::metaObject();
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "metaObject", args)
  }

}

  // proto:  int QSortFilterProxyModel::sortRole();
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "sortRole", args)
  }

}

  // proto:  void QSortFilterProxyModel::setSortLocaleAware(bool on);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setSortLocaleAware", args)
  }

}

  // proto:  QModelIndex QSortFilterProxyModel::mapToSource(const QModelIndex & proxyIndex);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "mapToSource", args)
  }

}

  // proto:  void QSortFilterProxyModel::setSourceModel(QAbstractItemModel * sourceModel);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setSourceModel", args)
  }

}

  // proto:  void QSortFilterProxyModel::QSortFilterProxyModel(const QSortFilterProxyModel & );
func NewQSortFilterProxyModel(args ...interface{}) QSortFilterProxyModel {
  return QSortFilterProxyModel{}
}

  // proto:  bool QSortFilterProxyModel::removeColumns(int column, int count, const QModelIndex & parent);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "removeColumns", args)
  }

}

  // proto:  void QSortFilterProxyModel::~QSortFilterProxyModel();
func (this *QSortFilterProxyModel) FreeQSortFilterProxyModel(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "~QSortFilterProxyModel", args)
  }

}

  // proto:  bool QSortFilterProxyModel::dynamicSortFilter();
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "dynamicSortFilter", args)
  }

}

  // proto:  bool QSortFilterProxyModel::insertColumns(int column, int count, const QModelIndex & parent);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "insertColumns", args)
  }

}

  // proto:  int QSortFilterProxyModel::columnCount(const QModelIndex & parent);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "columnCount", args)
  }

}

  // proto:  QModelIndex QSortFilterProxyModel::parent(const QModelIndex & child);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "parent", args)
  }

}

  // proto:  QRegExp QSortFilterProxyModel::filterRegExp();
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "filterRegExp", args)
  }

}

  // proto:  void QSortFilterProxyModel::setFilterRole(int role);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setFilterRole", args)
  }

}

  // proto:  bool QSortFilterProxyModel::removeRows(int row, int count, const QModelIndex & parent);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "removeRows", args)
  }

}

  // proto:  QModelIndex QSortFilterProxyModel::index(int row, int column, const QModelIndex & parent);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "index", args)
  }

}

  // proto:  void QSortFilterProxyModel::setDynamicSortFilter(bool enable);
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
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setDynamicSortFilter", args)
  }

}

// <= body block end

