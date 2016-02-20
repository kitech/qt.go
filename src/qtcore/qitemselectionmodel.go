package qtcore
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
// src-file: /QtCore/qitemselectionmodel.h
// dst-file: /src/core/qitemselectionmodel.go
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
  // proto:  bool QItemSelection::contains(const QModelIndex & index);
extern bool C_ZNK14QItemSelection8containsERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QModelIndexList QItemSelection::indexes();
extern void C_ZNK14QItemSelection7indexesEv(void* qthis); // 4
  // proto:  void QItemSelection::QItemSelection(const QModelIndex & topLeft, const QModelIndex & bottomRight);
extern void* C_ZN14QItemSelectionC2ERK11QModelIndexS2_(void* arg0, void* arg1); // 3
  // proto:  void QItemSelection::QItemSelection();
extern void* C_ZN14QItemSelectionC2Ev(); // 1
  // proto: static void QItemSelection::split(const QItemSelectionRange & range, const QItemSelectionRange & other, QItemSelection * result);
extern void C_ZN14QItemSelection5splitERK19QItemSelectionRangeS2_PS_(void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QItemSelection::select(const QModelIndex & topLeft, const QModelIndex & bottomRight);
extern void C_ZN14QItemSelection6selectERK11QModelIndexS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  const QPersistentModelIndex & QItemSelectionRange::topLeft();
extern void* C_ZNK19QItemSelectionRange7topLeftEv(void* qthis); // 2
  // proto:  int QItemSelectionRange::right();
extern int32_t C_ZNK19QItemSelectionRange5rightEv(void* qthis); // 2
  // proto:  QModelIndex QItemSelectionRange::parent();
extern void* C_ZNK19QItemSelectionRange6parentEv(void* qthis); // 2
  // proto:  int QItemSelectionRange::bottom();
extern int32_t C_ZNK19QItemSelectionRange6bottomEv(void* qthis); // 2
  // proto:  bool QItemSelectionRange::isValid();
extern bool C_ZNK19QItemSelectionRange7isValidEv(void* qthis); // 2
  // proto:  int QItemSelectionRange::top();
extern int32_t C_ZNK19QItemSelectionRange3topEv(void* qthis); // 2
  // proto:  bool QItemSelectionRange::contains(const QModelIndex & index);
extern bool C_ZNK19QItemSelectionRange8containsERK11QModelIndex(void* qthis, void* arg0); // 2
  // proto:  bool QItemSelectionRange::contains(int row, int column, const QModelIndex & parentIndex);
extern bool C_ZNK19QItemSelectionRange8containsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 2
  // proto:  void QItemSelectionRange::QItemSelectionRange(const QModelIndex & index);
extern void* C_ZN19QItemSelectionRangeC2ERK11QModelIndex(void* arg0); // 1
  // proto:  void QItemSelectionRange::QItemSelectionRange();
extern void* C_ZN19QItemSelectionRangeC2Ev(); // 1
  // proto:  void QItemSelectionRange::QItemSelectionRange(const QModelIndex & topLeft, const QModelIndex & bottomRight);
extern void* C_ZN19QItemSelectionRangeC2ERK11QModelIndexS2_(void* arg0, void* arg1); // 1
  // proto:  void QItemSelectionRange::QItemSelectionRange(const QItemSelectionRange & other);
extern void* C_ZN19QItemSelectionRangeC2ERKS_(void* arg0); // 1
  // proto:  QModelIndexList QItemSelectionRange::indexes();
extern void C_ZNK19QItemSelectionRange7indexesEv(void* qthis); // 4
  // proto:  int QItemSelectionRange::width();
extern int32_t C_ZNK19QItemSelectionRange5widthEv(void* qthis); // 2
  // proto:  const QPersistentModelIndex & QItemSelectionRange::bottomRight();
extern void* C_ZNK19QItemSelectionRange11bottomRightEv(void* qthis); // 2
  // proto:  bool QItemSelectionRange::isEmpty();
extern bool C_ZNK19QItemSelectionRange7isEmptyEv(void* qthis); // 4
  // proto:  bool QItemSelectionRange::intersects(const QItemSelectionRange & other);
extern bool C_ZNK19QItemSelectionRange10intersectsERKS_(void* qthis, void* arg0); // 4
  // proto:  const QAbstractItemModel * QItemSelectionRange::model();
extern void C_ZNK19QItemSelectionRange5modelEv(void* qthis); // 2
  // proto:  int QItemSelectionRange::height();
extern int32_t C_ZNK19QItemSelectionRange6heightEv(void* qthis); // 2
  // proto:  QItemSelectionRange QItemSelectionRange::intersected(const QItemSelectionRange & other);
extern void* C_ZNK19QItemSelectionRange11intersectedERKS_(void* qthis, void* arg0); // 4
  // proto:  int QItemSelectionRange::left();
extern int32_t C_ZNK19QItemSelectionRange4leftEv(void* qthis); // 2
  // proto:  const QItemSelection QItemSelectionModel::selection();
extern void* C_ZNK19QItemSelectionModel9selectionEv(void* qthis); // 4
  // proto:  bool QItemSelectionModel::isColumnSelected(int column, const QModelIndex & parent);
extern bool C_ZNK19QItemSelectionModel16isColumnSelectedEiRK11QModelIndex(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QModelIndexList QItemSelectionModel::selectedColumns(int row);
extern void C_ZNK19QItemSelectionModel15selectedColumnsEi(void* qthis, int32_t arg0); // 4
  // proto:  void QItemSelectionModel::clearSelection();
extern void C_ZN19QItemSelectionModel14clearSelectionEv(void* qthis); // 4
  // proto:  void QItemSelectionModel::~QItemSelectionModel();
extern void C_ZN19QItemSelectionModelD2Ev(void* qthis); // 4
  // proto:  bool QItemSelectionModel::isSelected(const QModelIndex & index);
extern bool C_ZNK19QItemSelectionModel10isSelectedERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QItemSelectionModel::rowIntersectsSelection(int row, const QModelIndex & parent);
extern bool C_ZNK19QItemSelectionModel22rowIntersectsSelectionEiRK11QModelIndex(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QItemSelectionModel::QItemSelectionModel(QAbstractItemModel * model, QObject * parent);
extern void* C_ZN19QItemSelectionModelC2EP18QAbstractItemModelP7QObject(void* arg0, void* arg1); // 3
  // proto:  void QItemSelectionModel::QItemSelectionModel(QAbstractItemModel * model);
extern void* C_ZN19QItemSelectionModelC2EP18QAbstractItemModel(void* arg0); // 3
  // proto:  void QItemSelectionModel::clearCurrentIndex();
extern void C_ZN19QItemSelectionModel17clearCurrentIndexEv(void* qthis); // 4
  // proto:  void QItemSelectionModel::setModel(QAbstractItemModel * model);
extern void C_ZN19QItemSelectionModel8setModelEP18QAbstractItemModel(void* qthis, void* arg0); // 4
  // proto:  bool QItemSelectionModel::isRowSelected(int row, const QModelIndex & parent);
extern bool C_ZNK19QItemSelectionModel13isRowSelectedEiRK11QModelIndex(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QItemSelectionModel::reset();
extern void C_ZN19QItemSelectionModel5resetEv(void* qthis); // 4
  // proto:  const QMetaObject * QItemSelectionModel::metaObject();
extern void C_ZNK19QItemSelectionModel10metaObjectEv(void* qthis); // 4
  // proto:  void QItemSelectionModel::clear();
extern void C_ZN19QItemSelectionModel5clearEv(void* qthis); // 4
  // proto:  QModelIndexList QItemSelectionModel::selectedIndexes();
extern void C_ZNK19QItemSelectionModel15selectedIndexesEv(void* qthis); // 4
  // proto:  bool QItemSelectionModel::columnIntersectsSelection(int column, const QModelIndex & parent);
extern bool C_ZNK19QItemSelectionModel25columnIntersectsSelectionEiRK11QModelIndex(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QModelIndex QItemSelectionModel::currentIndex();
extern void* C_ZNK19QItemSelectionModel12currentIndexEv(void* qthis); // 4
  // proto:  QModelIndexList QItemSelectionModel::selectedRows(int column);
extern void C_ZNK19QItemSelectionModel12selectedRowsEi(void* qthis, int32_t arg0); // 4
  // proto:  QAbstractItemModel * QItemSelectionModel::model();
extern void C_ZN19QItemSelectionModel5modelEv(void* qthis); // 4
  // proto:  bool QItemSelectionModel::hasSelection();
extern bool C_ZNK19QItemSelectionModel12hasSelectionEv(void* qthis); // 4
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

// class sizeof(QItemSelection)=1
type QItemSelection struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QItemSelectionRange)=16
type QItemSelectionRange struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QItemSelectionModel)=1
type QItemSelectionModel struct {
  /*qbase*/ QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _currentRowChanged QItemSelectionModel_currentRowChanged_signal;
//  _currentColumnChanged QItemSelectionModel_currentColumnChanged_signal;
//  _modelChanged QItemSelectionModel_modelChanged_signal;
//  _selectionChanged QItemSelectionModel_selectionChanged_signal;
//  _currentChanged QItemSelectionModel_currentChanged_signal;
}

// contains(const class QModelIndex &)
func (this *QItemSelection) Contains(args ...interface{}) (ret interface{}) {
  // contains(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QItemSelection8containsERK11QModelIndex
    // invoke: bool contains(const class QModelIndex &)
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK14QItemSelection8containsERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemSelection", "contains", args)
  }

  return
}

// indexes()
func (this *QItemSelection) Indexes(args ...interface{}) () {
  // indexes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QItemSelection7indexesEv
    // invoke: QModelIndexList indexes()
    C.C_ZNK14QItemSelection7indexesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelection", "indexes", args)
  }

  return
}

// QItemSelection(const class QModelIndex &, const class QModelIndex &)
func NewQItemSelection(args ...interface{}) *QItemSelection {
  // QItemSelection(const class QModelIndex &, const class QModelIndex &)
  // QItemSelection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QItemSelectionC1ERK11QModelIndexS2_
    // invoke: void QItemSelection(const class QModelIndex &, const class QModelIndex &)
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QItemSelectionC2ERK11QModelIndexS2_(arg0, arg1)
    return &QItemSelection{Qclsinst:qthis}
  case 1:
    // invoke: _ZN14QItemSelectionC1Ev
    // invoke: void QItemSelection()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QItemSelectionC2Ev()
    return &QItemSelection{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QItemSelection", "QItemSelection", args)
  }

  return nil // QItemSelection{Qclsinst:qthis}
}

// split(const class QItemSelectionRange &, const class QItemSelectionRange &, class QItemSelection *)
func (this *QItemSelection) Split_S(args ...interface{}) () {
  // split(const class QItemSelectionRange &, const class QItemSelectionRange &, class QItemSelection *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QItemSelectionRange{}) // "const QItemSelectionRange &"
  vtys[0][1] = reflect.TypeOf(QItemSelectionRange{}) // "const QItemSelectionRange &"
  vtys[0][2] = reflect.TypeOf(QItemSelection{}) // "QItemSelection *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QItemSelection5splitERK19QItemSelectionRangeS2_PS_
    // invoke: void split(const class QItemSelectionRange &, const class QItemSelectionRange &, class QItemSelection *)
    var arg0 = args[0].(*QItemSelectionRange).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QItemSelectionRange).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QItemSelection).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN14QItemSelection5splitERK19QItemSelectionRangeS2_PS_(arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QItemSelection", "split", args)
  }

  return
}

// select(const class QModelIndex &, const class QModelIndex &)
func (this *QItemSelection) Select_(args ...interface{}) () {
  // select(const class QModelIndex &, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QItemSelection6selectERK11QModelIndexS2_
    // invoke: void select(const class QModelIndex &, const class QModelIndex &)
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN14QItemSelection6selectERK11QModelIndexS2_(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QItemSelection", "select", args)
  }

  return
}

// topLeft()
func (this *QItemSelectionRange) Topleft(args ...interface{}) (ret interface{}) {
  // topLeft()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange7topLeftEv
    // invoke: const QPersistentModelIndex & topLeft()
    var ret0 = C.C_ZNK19QItemSelectionRange7topLeftEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPersistentModelIndex{}) // "const QPersistentModelIndex &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "topLeft", args)
  }

  return
}

// right()
func (this *QItemSelectionRange) Right(args ...interface{}) (ret interface{}) {
  // right()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange5rightEv
    // invoke: int right()
    var ret0 = C.C_ZNK19QItemSelectionRange5rightEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "right", args)
  }

  return
}

// parent()
func (this *QItemSelectionRange) Parent(args ...interface{}) (ret interface{}) {
  // parent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange6parentEv
    // invoke: QModelIndex parent()
    var ret0 = C.C_ZNK19QItemSelectionRange6parentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "parent", args)
  }

  return
}

// bottom()
func (this *QItemSelectionRange) Bottom(args ...interface{}) (ret interface{}) {
  // bottom()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange6bottomEv
    // invoke: int bottom()
    var ret0 = C.C_ZNK19QItemSelectionRange6bottomEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "bottom", args)
  }

  return
}

// isValid()
func (this *QItemSelectionRange) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK19QItemSelectionRange7isValidEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "isValid", args)
  }

  return
}

// top()
func (this *QItemSelectionRange) Top(args ...interface{}) (ret interface{}) {
  // top()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange3topEv
    // invoke: int top()
    var ret0 = C.C_ZNK19QItemSelectionRange3topEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "top", args)
  }

  return
}

// contains(const class QModelIndex &)
func (this *QItemSelectionRange) Contains(args ...interface{}) (ret interface{}) {
  // contains(const class QModelIndex &)
  // contains(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange8containsERK11QModelIndex
    // invoke: bool contains(const class QModelIndex &)
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QItemSelectionRange8containsERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK19QItemSelectionRange8containsEiiRK11QModelIndex
    // invoke: bool contains(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK19QItemSelectionRange8containsEiiRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "contains", args)
  }

  return
}

// QItemSelectionRange(const class QModelIndex &)
func NewQItemSelectionRange(args ...interface{}) *QItemSelectionRange {
  // QItemSelectionRange(const class QModelIndex &)
  // QItemSelectionRange()
  // QItemSelectionRange(const class QModelIndex &, const class QModelIndex &)
  // QItemSelectionRange(const class QItemSelectionRange &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[2][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QItemSelectionRange{}) // "const QItemSelectionRange &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QItemSelectionRangeC1ERK11QModelIndex
    // invoke: void QItemSelectionRange(const class QModelIndex &)
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QItemSelectionRangeC2ERK11QModelIndex(arg0)
    return &QItemSelectionRange{Qclsinst:qthis}
  case 1:
    // invoke: _ZN19QItemSelectionRangeC1Ev
    // invoke: void QItemSelectionRange()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QItemSelectionRangeC2Ev()
    return &QItemSelectionRange{Qclsinst:qthis}
  case 2:
    // invoke: _ZN19QItemSelectionRangeC1ERK11QModelIndexS2_
    // invoke: void QItemSelectionRange(const class QModelIndex &, const class QModelIndex &)
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QItemSelectionRangeC2ERK11QModelIndexS2_(arg0, arg1)
    return &QItemSelectionRange{Qclsinst:qthis}
  case 3:
    // invoke: _ZN19QItemSelectionRangeC1ERKS_
    // invoke: void QItemSelectionRange(const class QItemSelectionRange &)
    var arg0 = args[0].(*QItemSelectionRange).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QItemSelectionRangeC2ERKS_(arg0)
    return &QItemSelectionRange{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "QItemSelectionRange", args)
  }

  return nil // QItemSelectionRange{Qclsinst:qthis}
}

// indexes()
func (this *QItemSelectionRange) Indexes(args ...interface{}) () {
  // indexes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange7indexesEv
    // invoke: QModelIndexList indexes()
    C.C_ZNK19QItemSelectionRange7indexesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "indexes", args)
  }

  return
}

// width()
func (this *QItemSelectionRange) Width(args ...interface{}) (ret interface{}) {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange5widthEv
    // invoke: int width()
    var ret0 = C.C_ZNK19QItemSelectionRange5widthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "width", args)
  }

  return
}

// bottomRight()
func (this *QItemSelectionRange) Bottomright(args ...interface{}) (ret interface{}) {
  // bottomRight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange11bottomRightEv
    // invoke: const QPersistentModelIndex & bottomRight()
    var ret0 = C.C_ZNK19QItemSelectionRange11bottomRightEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPersistentModelIndex{}) // "const QPersistentModelIndex &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "bottomRight", args)
  }

  return
}

// isEmpty()
func (this *QItemSelectionRange) Isempty(args ...interface{}) (ret interface{}) {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange7isEmptyEv
    // invoke: bool isEmpty()
    var ret0 = C.C_ZNK19QItemSelectionRange7isEmptyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "isEmpty", args)
  }

  return
}

// intersects(const class QItemSelectionRange &)
func (this *QItemSelectionRange) Intersects(args ...interface{}) (ret interface{}) {
  // intersects(const class QItemSelectionRange &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QItemSelectionRange{}) // "const QItemSelectionRange &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange10intersectsERKS_
    // invoke: bool intersects(const class QItemSelectionRange &)
    var arg0 = args[0].(*QItemSelectionRange).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QItemSelectionRange10intersectsERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "intersects", args)
  }

  return
}

// model()
func (this *QItemSelectionRange) Model(args ...interface{}) () {
  // model()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange5modelEv
    // invoke: const QAbstractItemModel * model()
    C.C_ZNK19QItemSelectionRange5modelEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "model", args)
  }

  return
}

// height()
func (this *QItemSelectionRange) Height(args ...interface{}) (ret interface{}) {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange6heightEv
    // invoke: int height()
    var ret0 = C.C_ZNK19QItemSelectionRange6heightEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "height", args)
  }

  return
}

// intersected(const class QItemSelectionRange &)
func (this *QItemSelectionRange) Intersected(args ...interface{}) (ret interface{}) {
  // intersected(const class QItemSelectionRange &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QItemSelectionRange{}) // "const QItemSelectionRange &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange11intersectedERKS_
    // invoke: QItemSelectionRange intersected(const class QItemSelectionRange &)
    var arg0 = args[0].(*QItemSelectionRange).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QItemSelectionRange11intersectedERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QItemSelectionRange{}) // "QItemSelectionRange"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "intersected", args)
  }

  return
}

// left()
func (this *QItemSelectionRange) Left(args ...interface{}) (ret interface{}) {
  // left()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange4leftEv
    // invoke: int left()
    var ret0 = C.C_ZNK19QItemSelectionRange4leftEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "left", args)
  }

  return
}

// selection()
func (this *QItemSelectionModel) Selection(args ...interface{}) (ret interface{}) {
  // selection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionModel9selectionEv
    // invoke: const QItemSelection selection()
    var ret0 = C.C_ZNK19QItemSelectionModel9selectionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QItemSelection{}) // "const QItemSelection"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "selection", args)
  }

  return
}

// isColumnSelected(int, const class QModelIndex &)
func (this *QItemSelectionModel) Iscolumnselected(args ...interface{}) (ret interface{}) {
  // isColumnSelected(int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionModel16isColumnSelectedEiRK11QModelIndex
    // invoke: bool isColumnSelected(int, const class QModelIndex &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK19QItemSelectionModel16isColumnSelectedEiRK11QModelIndex(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "isColumnSelected", args)
  }

  return
}

// selectedColumns(int)
func (this *QItemSelectionModel) Selectedcolumns(args ...interface{}) () {
  // selectedColumns(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionModel15selectedColumnsEi
    // invoke: QModelIndexList selectedColumns(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK19QItemSelectionModel15selectedColumnsEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "selectedColumns", args)
  }

  return
}

// clearSelection()
func (this *QItemSelectionModel) Clearselection(args ...interface{}) () {
  // clearSelection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QItemSelectionModel14clearSelectionEv
    // invoke: void clearSelection()
    C.C_ZN19QItemSelectionModel14clearSelectionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "clearSelection", args)
  }

  return
}

// ~QItemSelectionModel()
func (this *QItemSelectionModel) Freeqitemselectionmodel(args ...interface{}) () {
  // ~QItemSelectionModel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QItemSelectionModelD0Ev
    // invoke: void ~QItemSelectionModel()
    C.C_ZN19QItemSelectionModelD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "~QItemSelectionModel", args)
  }

  return
}

// isSelected(const class QModelIndex &)
func (this *QItemSelectionModel) Isselected(args ...interface{}) (ret interface{}) {
  // isSelected(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionModel10isSelectedERK11QModelIndex
    // invoke: bool isSelected(const class QModelIndex &)
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QItemSelectionModel10isSelectedERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "isSelected", args)
  }

  return
}

// rowIntersectsSelection(int, const class QModelIndex &)
func (this *QItemSelectionModel) Rowintersectsselection(args ...interface{}) (ret interface{}) {
  // rowIntersectsSelection(int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionModel22rowIntersectsSelectionEiRK11QModelIndex
    // invoke: bool rowIntersectsSelection(int, const class QModelIndex &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK19QItemSelectionModel22rowIntersectsSelectionEiRK11QModelIndex(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "rowIntersectsSelection", args)
  }

  return
}

// QItemSelectionModel(class QAbstractItemModel *, class QObject *)
func NewQItemSelectionModel(args ...interface{}) *QItemSelectionModel {
  // QItemSelectionModel(class QAbstractItemModel *, class QObject *)
  // QItemSelectionModel(class QAbstractItemModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemModel{}) // "QAbstractItemModel *"
  vtys[0][1] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QAbstractItemModel{}) // "QAbstractItemModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QItemSelectionModelC1EP18QAbstractItemModelP7QObject
    // invoke: void QItemSelectionModel(class QAbstractItemModel *, class QObject *)
    var arg0 = args[0].(*QAbstractItemModel).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QObject).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QItemSelectionModelC2EP18QAbstractItemModelP7QObject(arg0, arg1)
    return &QItemSelectionModel{Qclsinst:qthis}
  case 1:
    // invoke: _ZN19QItemSelectionModelC1EP18QAbstractItemModel
    // invoke: void QItemSelectionModel(class QAbstractItemModel *)
    var arg0 = args[0].(*QAbstractItemModel).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QItemSelectionModelC2EP18QAbstractItemModel(arg0)
    return &QItemSelectionModel{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "QItemSelectionModel", args)
  }

  return nil // QItemSelectionModel{Qclsinst:qthis}
}

// clearCurrentIndex()
func (this *QItemSelectionModel) Clearcurrentindex(args ...interface{}) () {
  // clearCurrentIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QItemSelectionModel17clearCurrentIndexEv
    // invoke: void clearCurrentIndex()
    C.C_ZN19QItemSelectionModel17clearCurrentIndexEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "clearCurrentIndex", args)
  }

  return
}

// setModel(class QAbstractItemModel *)
func (this *QItemSelectionModel) Setmodel(args ...interface{}) () {
  // setModel(class QAbstractItemModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemModel{}) // "QAbstractItemModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QItemSelectionModel8setModelEP18QAbstractItemModel
    // invoke: void setModel(class QAbstractItemModel *)
    var arg0 = args[0].(*QAbstractItemModel).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QItemSelectionModel8setModelEP18QAbstractItemModel(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "setModel", args)
  }

  return
}

// isRowSelected(int, const class QModelIndex &)
func (this *QItemSelectionModel) Isrowselected(args ...interface{}) (ret interface{}) {
  // isRowSelected(int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionModel13isRowSelectedEiRK11QModelIndex
    // invoke: bool isRowSelected(int, const class QModelIndex &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK19QItemSelectionModel13isRowSelectedEiRK11QModelIndex(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "isRowSelected", args)
  }

  return
}

// reset()
func (this *QItemSelectionModel) Reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QItemSelectionModel5resetEv
    // invoke: void reset()
    C.C_ZN19QItemSelectionModel5resetEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "reset", args)
  }

  return
}

// metaObject()
func (this *QItemSelectionModel) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionModel10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK19QItemSelectionModel10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "metaObject", args)
  }

  return
}

// clear()
func (this *QItemSelectionModel) Clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QItemSelectionModel5clearEv
    // invoke: void clear()
    C.C_ZN19QItemSelectionModel5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "clear", args)
  }

  return
}

// selectedIndexes()
func (this *QItemSelectionModel) Selectedindexes(args ...interface{}) () {
  // selectedIndexes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionModel15selectedIndexesEv
    // invoke: QModelIndexList selectedIndexes()
    C.C_ZNK19QItemSelectionModel15selectedIndexesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "selectedIndexes", args)
  }

  return
}

// columnIntersectsSelection(int, const class QModelIndex &)
func (this *QItemSelectionModel) Columnintersectsselection(args ...interface{}) (ret interface{}) {
  // columnIntersectsSelection(int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionModel25columnIntersectsSelectionEiRK11QModelIndex
    // invoke: bool columnIntersectsSelection(int, const class QModelIndex &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK19QItemSelectionModel25columnIntersectsSelectionEiRK11QModelIndex(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "columnIntersectsSelection", args)
  }

  return
}

// currentIndex()
func (this *QItemSelectionModel) Currentindex(args ...interface{}) (ret interface{}) {
  // currentIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionModel12currentIndexEv
    // invoke: QModelIndex currentIndex()
    var ret0 = C.C_ZNK19QItemSelectionModel12currentIndexEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "currentIndex", args)
  }

  return
}

// selectedRows(int)
func (this *QItemSelectionModel) Selectedrows(args ...interface{}) () {
  // selectedRows(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionModel12selectedRowsEi
    // invoke: QModelIndexList selectedRows(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK19QItemSelectionModel12selectedRowsEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "selectedRows", args)
  }

  return
}

// model()
func (this *QItemSelectionModel) Model(args ...interface{}) () {
  // model()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QItemSelectionModel5modelEv
    // invoke: QAbstractItemModel * model()
    C.C_ZN19QItemSelectionModel5modelEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "model", args)
  }

  return
}

// hasSelection()
func (this *QItemSelectionModel) Hasselection(args ...interface{}) (ret interface{}) {
  // hasSelection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionModel12hasSelectionEv
    // invoke: bool hasSelection()
    var ret0 = C.C_ZNK19QItemSelectionModel12hasSelectionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "hasSelection", args)
  }

  return
}

// <= body block end

