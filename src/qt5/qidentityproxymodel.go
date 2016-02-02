package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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
extern int32_t C_ZNK19QIdentityProxyModel11columnCountERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QItemSelection QIdentityProxyModel::mapSelectionToSource(const QItemSelection & selection);
extern void* C_ZNK19QIdentityProxyModel20mapSelectionToSourceERK14QItemSelection(void* qthis, void* arg0); // 4
  // proto:  void QIdentityProxyModel::setSourceModel(QAbstractItemModel * sourceModel);
extern void C_ZN19QIdentityProxyModel14setSourceModelEP18QAbstractItemModel(void* qthis, void* arg0); // 4
  // proto:  QModelIndex QIdentityProxyModel::sibling(int row, int column, const QModelIndex & idx);
extern void* C_ZNK19QIdentityProxyModel7siblingEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  QModelIndex QIdentityProxyModel::index(int row, int column, const QModelIndex & parent);
extern void* C_ZNK19QIdentityProxyModel5indexEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  QModelIndex QIdentityProxyModel::mapFromSource(const QModelIndex & sourceIndex);
extern void* C_ZNK19QIdentityProxyModel13mapFromSourceERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QIdentityProxyModel::removeRows(int row, int count, const QModelIndex & parent);
extern bool C_ZN19QIdentityProxyModel10removeRowsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QIdentityProxyModel::~QIdentityProxyModel();
extern void C_ZN19QIdentityProxyModelD2Ev(void* qthis); // 4
  // proto:  QModelIndex QIdentityProxyModel::parent(const QModelIndex & child);
extern void* C_ZNK19QIdentityProxyModel6parentERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QItemSelection QIdentityProxyModel::mapSelectionFromSource(const QItemSelection & selection);
extern void* C_ZNK19QIdentityProxyModel22mapSelectionFromSourceERK14QItemSelection(void* qthis, void* arg0); // 4
  // proto:  bool QIdentityProxyModel::insertRows(int row, int count, const QModelIndex & parent);
extern bool C_ZN19QIdentityProxyModel10insertRowsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QIdentityProxyModel::QIdentityProxyModel(QObject * parent);
extern void* C_ZN19QIdentityProxyModelC2EP7QObject(void* arg0); // 3
  // proto:  const QMetaObject * QIdentityProxyModel::metaObject();
extern void C_ZNK19QIdentityProxyModel10metaObjectEv(void* qthis); // 4
  // proto:  int QIdentityProxyModel::rowCount(const QModelIndex & parent);
extern int32_t C_ZNK19QIdentityProxyModel8rowCountERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QModelIndex QIdentityProxyModel::mapToSource(const QModelIndex & proxyIndex);
extern void* C_ZNK19QIdentityProxyModel11mapToSourceERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QIdentityProxyModel::insertColumns(int column, int count, const QModelIndex & parent);
extern bool C_ZN19QIdentityProxyModel13insertColumnsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  bool QIdentityProxyModel::removeColumns(int column, int count, const QModelIndex & parent);
extern bool C_ZN19QIdentityProxyModel13removeColumnsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// columnCount(const class QModelIndex &)
func (this *QIdentityProxyModel) Columncount(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QIdentityProxyModel11columnCountERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "columnCount", args)
  }

  return
}

// mapSelectionToSource(const class QItemSelection &)
func (this *QIdentityProxyModel) Mapselectiontosource(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QItemSelection).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QIdentityProxyModel20mapSelectionToSourceERK14QItemSelection(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QItemSelection{}) // "QItemSelection"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "mapSelectionToSource", args)
  }

  return
}

// setSourceModel(class QAbstractItemModel *)
func (this *QIdentityProxyModel) Setsourcemodel(args ...interface{}) () {
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
    var arg0 = args[0].(*QAbstractItemModel).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QIdentityProxyModel14setSourceModelEP18QAbstractItemModel(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "setSourceModel", args)
  }

  return
}

// sibling(int, int, const class QModelIndex &)
func (this *QIdentityProxyModel) Sibling(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK19QIdentityProxyModel7siblingEiiRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "sibling", args)
  }

  return
}

// index(int, int, const class QModelIndex &)
func (this *QIdentityProxyModel) Index(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK19QIdentityProxyModel5indexEiiRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "index", args)
  }

  return
}

// mapFromSource(const class QModelIndex &)
func (this *QIdentityProxyModel) Mapfromsource(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QIdentityProxyModel13mapFromSourceERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "mapFromSource", args)
  }

  return
}

// removeRows(int, int, const class QModelIndex &)
func (this *QIdentityProxyModel) Removerows(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN19QIdentityProxyModel10removeRowsEiiRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "removeRows", args)
  }

  return
}

// ~QIdentityProxyModel()
func (this *QIdentityProxyModel) Freeqidentityproxymodel(args ...interface{}) () {
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
    C.C_ZN19QIdentityProxyModelD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "~QIdentityProxyModel", args)
  }

  return
}

// parent(const class QModelIndex &)
func (this *QIdentityProxyModel) Parent(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QIdentityProxyModel6parentERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "parent", args)
  }

  return
}

// mapSelectionFromSource(const class QItemSelection &)
func (this *QIdentityProxyModel) Mapselectionfromsource(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QItemSelection).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QIdentityProxyModel22mapSelectionFromSourceERK14QItemSelection(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QItemSelection{}) // "QItemSelection"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "mapSelectionFromSource", args)
  }

  return
}

// insertRows(int, int, const class QModelIndex &)
func (this *QIdentityProxyModel) Insertrows(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN19QIdentityProxyModel10insertRowsEiiRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "insertRows", args)
  }

  return
}

// QIdentityProxyModel(class QObject *)
func NewQIdentityProxyModel(args ...interface{}) *QIdentityProxyModel {
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
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QIdentityProxyModelC2EP7QObject(arg0)
    return &QIdentityProxyModel{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "QIdentityProxyModel", args)
  }

  return nil // QIdentityProxyModel{Qclsinst:qthis}
}

// metaObject()
func (this *QIdentityProxyModel) Metaobject(args ...interface{}) () {
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
    C.C_ZNK19QIdentityProxyModel10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "metaObject", args)
  }

  return
}

// rowCount(const class QModelIndex &)
func (this *QIdentityProxyModel) Rowcount(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QIdentityProxyModel8rowCountERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "rowCount", args)
  }

  return
}

// mapToSource(const class QModelIndex &)
func (this *QIdentityProxyModel) Maptosource(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QIdentityProxyModel11mapToSourceERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "mapToSource", args)
  }

  return
}

// insertColumns(int, int, const class QModelIndex &)
func (this *QIdentityProxyModel) Insertcolumns(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN19QIdentityProxyModel13insertColumnsEiiRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "insertColumns", args)
  }

  return
}

// removeColumns(int, int, const class QModelIndex &)
func (this *QIdentityProxyModel) Removecolumns(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN19QIdentityProxyModel13removeColumnsEiiRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIdentityProxyModel", "removeColumns", args)
  }

  return
}

// <= body block end

