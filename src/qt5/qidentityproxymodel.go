package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtCore/qidentityproxymodel.h
// dst-file: /src/core/qidentityproxymodel.go
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
  // proto:  bool QIdentityProxyModel::removeRows(int row, int count, const QModelIndex & parent);
extern void _ZN19QIdentityProxyModel10removeRowsEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  bool QIdentityProxyModel::removeColumns(int column, int count, const QModelIndex & parent);
extern void _ZN19QIdentityProxyModel13removeColumnsEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  QItemSelection QIdentityProxyModel::mapSelectionFromSource(const QItemSelection & selection);
extern void _ZNK19QIdentityProxyModel22mapSelectionFromSourceERK14QItemSelection(void* qthis, void* arg0);
  // proto:  QModelIndex QIdentityProxyModel::index(int row, int column, const QModelIndex & parent);
extern void _ZNK19QIdentityProxyModel5indexEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  bool QIdentityProxyModel::insertRows(int row, int count, const QModelIndex & parent);
extern void _ZN19QIdentityProxyModel10insertRowsEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  bool QIdentityProxyModel::insertColumns(int column, int count, const QModelIndex & parent);
extern void _ZN19QIdentityProxyModel13insertColumnsEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  QModelIndex QIdentityProxyModel::sibling(int row, int column, const QModelIndex & idx);
extern void _ZNK19QIdentityProxyModel7siblingEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  const QMetaObject * QIdentityProxyModel::metaObject();
extern void _ZNK19QIdentityProxyModel10metaObjectEv(void* qthis);
  // proto:  void QIdentityProxyModel::~QIdentityProxyModel();
extern void _ZN19QIdentityProxyModelD0Ev(void* qthis);
  // proto:  QModelIndex QIdentityProxyModel::parent(const QModelIndex & child);
extern void _ZNK19QIdentityProxyModel6parentERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QIdentityProxyModel::setSourceModel(QAbstractItemModel * sourceModel);
extern void _ZN19QIdentityProxyModel14setSourceModelEP18QAbstractItemModel(void* qthis, void* arg0);
  // proto:  QModelIndex QIdentityProxyModel::mapToSource(const QModelIndex & proxyIndex);
extern void _ZNK19QIdentityProxyModel11mapToSourceERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QModelIndex QIdentityProxyModel::mapFromSource(const QModelIndex & sourceIndex);
extern void _ZNK19QIdentityProxyModel13mapFromSourceERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QIdentityProxyModel::QIdentityProxyModel(QObject * parent);
extern void* dector_ZN19QIdentityProxyModelC1EP7QObject(void* arg0);
extern void _ZN19QIdentityProxyModelC1EP7QObject(void* qthis, void* arg0);
  // proto:  int QIdentityProxyModel::columnCount(const QModelIndex & parent);
extern void _ZNK19QIdentityProxyModel11columnCountERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QIdentityProxyModel::QIdentityProxyModel(const QIdentityProxyModel & );
extern void* dector_ZN19QIdentityProxyModelC1ERKS_(void* arg0);
extern void _ZN19QIdentityProxyModelC1ERKS_(void* qthis, void* arg0);
  // proto:  QItemSelection QIdentityProxyModel::mapSelectionToSource(const QItemSelection & selection);
extern void _ZNK19QIdentityProxyModel20mapSelectionToSourceERK14QItemSelection(void* qthis, void* arg0);
  // proto:  int QIdentityProxyModel::rowCount(const QModelIndex & parent);
extern void _ZNK19QIdentityProxyModel8rowCountERK11QModelIndex(void* qthis, void* arg0);
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

// class sizeof(QIdentityProxyModel)=1
type QIdentityProxyModel struct {
  /*qbase*/ QAbstractProxyModel;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  bool QIdentityProxyModel::removeRows(int row, int count, const QModelIndex & parent);
func (this *QIdentityProxyModel) removeRows(args ...interface{}) () {
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
    // invoke: _ZN19QIdentityProxyModel10removeRowsEiiRK11QModelIndex
    // invoke: bool removeRows(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN19QIdentityProxyModel10removeRowsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "removeRows", args)
  }

}

  // proto:  bool QIdentityProxyModel::removeColumns(int column, int count, const QModelIndex & parent);
func (this *QIdentityProxyModel) removeColumns(args ...interface{}) () {
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
    // invoke: _ZN19QIdentityProxyModel13removeColumnsEiiRK11QModelIndex
    // invoke: bool removeColumns(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN19QIdentityProxyModel13removeColumnsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "removeColumns", args)
  }

}

  // proto:  QItemSelection QIdentityProxyModel::mapSelectionFromSource(const QItemSelection & selection);
func (this *QIdentityProxyModel) mapSelectionFromSource(args ...interface{}) () {
  // mapSelectionFromSource(const class QItemSelection &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QItemSelection{}) // "const QItemSelection &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QIdentityProxyModel22mapSelectionFromSourceERK14QItemSelection
    // invoke: QItemSelection mapSelectionFromSource(const class QItemSelection &)
    var arg0 = args[0].(QItemSelection).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK19QIdentityProxyModel22mapSelectionFromSourceERK14QItemSelection(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "mapSelectionFromSource", args)
  }

}

  // proto:  QModelIndex QIdentityProxyModel::index(int row, int column, const QModelIndex & parent);
func (this *QIdentityProxyModel) index(args ...interface{}) () {
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
    // invoke: _ZNK19QIdentityProxyModel5indexEiiRK11QModelIndex
    // invoke: QModelIndex index(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C._ZNK19QIdentityProxyModel5indexEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "index", args)
  }

}

  // proto:  bool QIdentityProxyModel::insertRows(int row, int count, const QModelIndex & parent);
func (this *QIdentityProxyModel) insertRows(args ...interface{}) () {
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
    // invoke: _ZN19QIdentityProxyModel10insertRowsEiiRK11QModelIndex
    // invoke: bool insertRows(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN19QIdentityProxyModel10insertRowsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "insertRows", args)
  }

}

  // proto:  bool QIdentityProxyModel::insertColumns(int column, int count, const QModelIndex & parent);
func (this *QIdentityProxyModel) insertColumns(args ...interface{}) () {
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
    // invoke: _ZN19QIdentityProxyModel13insertColumnsEiiRK11QModelIndex
    // invoke: bool insertColumns(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN19QIdentityProxyModel13insertColumnsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "insertColumns", args)
  }

}

  // proto:  QModelIndex QIdentityProxyModel::sibling(int row, int column, const QModelIndex & idx);
func (this *QIdentityProxyModel) sibling(args ...interface{}) () {
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
    // invoke: _ZNK19QIdentityProxyModel7siblingEiiRK11QModelIndex
    // invoke: QModelIndex sibling(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C._ZNK19QIdentityProxyModel7siblingEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "sibling", args)
  }

}

  // proto:  const QMetaObject * QIdentityProxyModel::metaObject();
func (this *QIdentityProxyModel) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QIdentityProxyModel10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK19QIdentityProxyModel10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "metaObject", args)
  }

}

  // proto:  void QIdentityProxyModel::~QIdentityProxyModel();
func (this *QIdentityProxyModel) FreeQIdentityProxyModel(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "~QIdentityProxyModel", args)
  }

}

  // proto:  QModelIndex QIdentityProxyModel::parent(const QModelIndex & child);
func (this *QIdentityProxyModel) parent(args ...interface{}) () {
  // parent(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QIdentityProxyModel6parentERK11QModelIndex
    // invoke: QModelIndex parent(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK19QIdentityProxyModel6parentERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "parent", args)
  }

}

  // proto:  void QIdentityProxyModel::setSourceModel(QAbstractItemModel * sourceModel);
func (this *QIdentityProxyModel) setSourceModel(args ...interface{}) () {
  // setSourceModel(class QAbstractItemModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemModel{}) // "QAbstractItemModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QIdentityProxyModel14setSourceModelEP18QAbstractItemModel
    // invoke: void setSourceModel(class QAbstractItemModel *)
    var arg0 = args[0].(QAbstractItemModel).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN19QIdentityProxyModel14setSourceModelEP18QAbstractItemModel(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "setSourceModel", args)
  }

}

  // proto:  QModelIndex QIdentityProxyModel::mapToSource(const QModelIndex & proxyIndex);
func (this *QIdentityProxyModel) mapToSource(args ...interface{}) () {
  // mapToSource(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QIdentityProxyModel11mapToSourceERK11QModelIndex
    // invoke: QModelIndex mapToSource(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK19QIdentityProxyModel11mapToSourceERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "mapToSource", args)
  }

}

  // proto:  QModelIndex QIdentityProxyModel::mapFromSource(const QModelIndex & sourceIndex);
func (this *QIdentityProxyModel) mapFromSource(args ...interface{}) () {
  // mapFromSource(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QIdentityProxyModel13mapFromSourceERK11QModelIndex
    // invoke: QModelIndex mapFromSource(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK19QIdentityProxyModel13mapFromSourceERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "mapFromSource", args)
  }

}

  // proto:  void QIdentityProxyModel::QIdentityProxyModel(QObject * parent);
func NewQIdentityProxyModel(args ...interface{}) QIdentityProxyModel {
  return QIdentityProxyModel{}
}

  // proto:  int QIdentityProxyModel::columnCount(const QModelIndex & parent);
func (this *QIdentityProxyModel) columnCount(args ...interface{}) () {
  // columnCount(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QIdentityProxyModel11columnCountERK11QModelIndex
    // invoke: int columnCount(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK19QIdentityProxyModel11columnCountERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "columnCount", args)
  }

}

  // proto:  QItemSelection QIdentityProxyModel::mapSelectionToSource(const QItemSelection & selection);
func (this *QIdentityProxyModel) mapSelectionToSource(args ...interface{}) () {
  // mapSelectionToSource(const class QItemSelection &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QItemSelection{}) // "const QItemSelection &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QIdentityProxyModel20mapSelectionToSourceERK14QItemSelection
    // invoke: QItemSelection mapSelectionToSource(const class QItemSelection &)
    var arg0 = args[0].(QItemSelection).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK19QIdentityProxyModel20mapSelectionToSourceERK14QItemSelection(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "mapSelectionToSource", args)
  }

}

  // proto:  int QIdentityProxyModel::rowCount(const QModelIndex & parent);
func (this *QIdentityProxyModel) rowCount(args ...interface{}) () {
  // rowCount(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QIdentityProxyModel8rowCountERK11QModelIndex
    // invoke: int rowCount(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK19QIdentityProxyModel8rowCountERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "rowCount", args)
  }

}

// <= body block end

