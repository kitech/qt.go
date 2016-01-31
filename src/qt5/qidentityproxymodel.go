package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  int QIdentityProxyModel::columnCount(const QModelIndex & parent);
extern void C_ZNK19QIdentityProxyModel11columnCountERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QItemSelection QIdentityProxyModel::mapSelectionToSource(const QItemSelection & selection);
extern void C_ZNK19QIdentityProxyModel20mapSelectionToSourceERK14QItemSelection(void* qthis, void* arg0); // 4
  // proto:  void QIdentityProxyModel::setSourceModel(QAbstractItemModel * sourceModel);
extern void C_ZN19QIdentityProxyModel14setSourceModelEP18QAbstractItemModel(void* qthis, void* arg0); // 4
  // proto:  QModelIndex QIdentityProxyModel::sibling(int row, int column, const QModelIndex & idx);
extern void C_ZNK19QIdentityProxyModel7siblingEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  QModelIndex QIdentityProxyModel::index(int row, int column, const QModelIndex & parent);
extern void C_ZNK19QIdentityProxyModel5indexEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  QModelIndex QIdentityProxyModel::mapFromSource(const QModelIndex & sourceIndex);
extern void C_ZNK19QIdentityProxyModel13mapFromSourceERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QIdentityProxyModel::removeRows(int row, int count, const QModelIndex & parent);
extern void C_ZN19QIdentityProxyModel10removeRowsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QIdentityProxyModel::~QIdentityProxyModel();
extern void C_ZN19QIdentityProxyModelD2Ev(void* qthis); // 4
  // proto:  QModelIndex QIdentityProxyModel::parent(const QModelIndex & child);
extern void C_ZNK19QIdentityProxyModel6parentERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QItemSelection QIdentityProxyModel::mapSelectionFromSource(const QItemSelection & selection);
extern void C_ZNK19QIdentityProxyModel22mapSelectionFromSourceERK14QItemSelection(void* qthis, void* arg0); // 4
  // proto:  bool QIdentityProxyModel::insertRows(int row, int count, const QModelIndex & parent);
extern void C_ZN19QIdentityProxyModel10insertRowsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QIdentityProxyModel::QIdentityProxyModel(QObject * parent);
extern void C_ZN19QIdentityProxyModelC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  const QMetaObject * QIdentityProxyModel::metaObject();
extern void C_ZNK19QIdentityProxyModel10metaObjectEv(void* qthis); // 4
  // proto:  int QIdentityProxyModel::rowCount(const QModelIndex & parent);
extern void C_ZNK19QIdentityProxyModel8rowCountERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QModelIndex QIdentityProxyModel::mapToSource(const QModelIndex & proxyIndex);
extern void C_ZNK19QIdentityProxyModel11mapToSourceERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QIdentityProxyModel::insertColumns(int column, int count, const QModelIndex & parent);
extern void C_ZN19QIdentityProxyModel13insertColumnsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  bool QIdentityProxyModel::removeColumns(int column, int count, const QModelIndex & parent);
extern void C_ZN19QIdentityProxyModel13removeColumnsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// columnCount(const class QModelIndex &)
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
    C.C_ZNK19QIdentityProxyModel11columnCountERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "columnCount", args)
  }

}

// mapSelectionToSource(const class QItemSelection &)
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
    C.C_ZNK19QIdentityProxyModel20mapSelectionToSourceERK14QItemSelection(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "mapSelectionToSource", args)
  }

}

// setSourceModel(class QAbstractItemModel *)
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
    C.C_ZN19QIdentityProxyModel14setSourceModelEP18QAbstractItemModel(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "setSourceModel", args)
  }

}

// sibling(int, int, const class QModelIndex &)
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
    C.C_ZNK19QIdentityProxyModel7siblingEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "sibling", args)
  }

}

// index(int, int, const class QModelIndex &)
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
    C.C_ZNK19QIdentityProxyModel5indexEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "index", args)
  }

}

// mapFromSource(const class QModelIndex &)
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
    C.C_ZNK19QIdentityProxyModel13mapFromSourceERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "mapFromSource", args)
  }

}

// removeRows(int, int, const class QModelIndex &)
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
    C.C_ZN19QIdentityProxyModel10removeRowsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "removeRows", args)
  }

}

// ~QIdentityProxyModel()
func (this *QIdentityProxyModel) FreeQIdentityProxyModel(args ...interface{}) () {
  // ~QIdentityProxyModel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QIdentityProxyModelD0Ev
    // invoke: void ~QIdentityProxyModel()
    C.C_ZN19QIdentityProxyModelD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "~QIdentityProxyModel", args)
  }

}

// parent(const class QModelIndex &)
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
    C.C_ZNK19QIdentityProxyModel6parentERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "parent", args)
  }

}

// mapSelectionFromSource(const class QItemSelection &)
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
    C.C_ZNK19QIdentityProxyModel22mapSelectionFromSourceERK14QItemSelection(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "mapSelectionFromSource", args)
  }

}

// insertRows(int, int, const class QModelIndex &)
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
    C.C_ZN19QIdentityProxyModel10insertRowsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "insertRows", args)
  }

}

// QIdentityProxyModel(class QObject *)
func NewQIdentityProxyModel(args ...interface{}) QIdentityProxyModel {
  // QIdentityProxyModel(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QIdentityProxyModelC1EP7QObject
    // invoke: void QIdentityProxyModel(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN19QIdentityProxyModelC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "QIdentityProxyModel", args)
  }

  return QIdentityProxyModel{}
}

// metaObject()
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
    C.C_ZNK19QIdentityProxyModel10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "metaObject", args)
  }

}

// rowCount(const class QModelIndex &)
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
    C.C_ZNK19QIdentityProxyModel8rowCountERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "rowCount", args)
  }

}

// mapToSource(const class QModelIndex &)
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
    C.C_ZNK19QIdentityProxyModel11mapToSourceERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "mapToSource", args)
  }

}

// insertColumns(int, int, const class QModelIndex &)
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
    C.C_ZN19QIdentityProxyModel13insertColumnsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "insertColumns", args)
  }

}

// removeColumns(int, int, const class QModelIndex &)
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
    C.C_ZN19QIdentityProxyModel13removeColumnsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "removeColumns", args)
  }

}

// <= body block end

