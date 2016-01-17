package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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
extern void _ZNK14QItemSelection8containsERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QModelIndexList QItemSelection::indexes();
extern void _ZNK14QItemSelection7indexesEv(void* qthis); // 4
  // proto:  void QItemSelection::QItemSelection(const QModelIndex & topLeft, const QModelIndex & bottomRight);
extern void _ZN14QItemSelectionC2ERK11QModelIndexS2_(void* qthis, void* arg0, void* arg1); // 3
  // proto:  void QItemSelection::QItemSelection();
extern void _ZN14QItemSelectionC2Ev(void* qthis); // 1
  // proto: static void QItemSelection::split(const QItemSelectionRange & range, const QItemSelectionRange & other, QItemSelection * result);
extern void _ZN14QItemSelection5splitERK19QItemSelectionRangeS2_PS_(void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QItemSelection::select(const QModelIndex & topLeft, const QModelIndex & bottomRight);
extern void _ZN14QItemSelection6selectERK11QModelIndexS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  const QPersistentModelIndex & QItemSelectionRange::topLeft();
extern void _ZNK19QItemSelectionRange7topLeftEv(void* qthis); // 2
  // proto:  int QItemSelectionRange::right();
extern void _ZNK19QItemSelectionRange5rightEv(void* qthis); // 2
  // proto:  QModelIndex QItemSelectionRange::parent();
extern void _ZNK19QItemSelectionRange6parentEv(void* qthis); // 2
  // proto:  int QItemSelectionRange::bottom();
extern void _ZNK19QItemSelectionRange6bottomEv(void* qthis); // 2
  // proto:  bool QItemSelectionRange::isValid();
extern void _ZNK19QItemSelectionRange7isValidEv(void* qthis); // 2
  // proto:  int QItemSelectionRange::top();
extern void _ZNK19QItemSelectionRange3topEv(void* qthis); // 2
  // proto:  bool QItemSelectionRange::contains(const QModelIndex & index);
extern void _ZNK19QItemSelectionRange8containsERK11QModelIndex(void* qthis, void* arg0); // 2
  // proto:  bool QItemSelectionRange::contains(int row, int column, const QModelIndex & parentIndex);
extern void _ZNK19QItemSelectionRange8containsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 2
  // proto:  void QItemSelectionRange::QItemSelectionRange(const QModelIndex & index);
extern void _ZN19QItemSelectionRangeC2ERK11QModelIndex(void* qthis, void* arg0); // 1
  // proto:  void QItemSelectionRange::QItemSelectionRange();
extern void _ZN19QItemSelectionRangeC2Ev(void* qthis); // 1
  // proto:  void QItemSelectionRange::QItemSelectionRange(const QModelIndex & topLeft, const QModelIndex & bottomRight);
extern void _ZN19QItemSelectionRangeC2ERK11QModelIndexS2_(void* qthis, void* arg0, void* arg1); // 1
  // proto:  void QItemSelectionRange::QItemSelectionRange(const QItemSelectionRange & other);
extern void _ZN19QItemSelectionRangeC2ERKS_(void* qthis, void* arg0); // 1
  // proto:  QModelIndexList QItemSelectionRange::indexes();
extern void _ZNK19QItemSelectionRange7indexesEv(void* qthis); // 4
  // proto:  int QItemSelectionRange::width();
extern void _ZNK19QItemSelectionRange5widthEv(void* qthis); // 2
  // proto:  const QPersistentModelIndex & QItemSelectionRange::bottomRight();
extern void _ZNK19QItemSelectionRange11bottomRightEv(void* qthis); // 2
  // proto:  bool QItemSelectionRange::isEmpty();
extern void _ZNK19QItemSelectionRange7isEmptyEv(void* qthis); // 4
  // proto:  bool QItemSelectionRange::intersects(const QItemSelectionRange & other);
extern void _ZNK19QItemSelectionRange10intersectsERKS_(void* qthis, void* arg0); // 4
  // proto:  const QAbstractItemModel * QItemSelectionRange::model();
extern void _ZNK19QItemSelectionRange5modelEv(void* qthis); // 2
  // proto:  int QItemSelectionRange::height();
extern void _ZNK19QItemSelectionRange6heightEv(void* qthis); // 2
  // proto:  QItemSelectionRange QItemSelectionRange::intersected(const QItemSelectionRange & other);
extern void _ZNK19QItemSelectionRange11intersectedERKS_(void* qthis, void* arg0); // 4
  // proto:  int QItemSelectionRange::left();
extern void _ZNK19QItemSelectionRange4leftEv(void* qthis); // 2
  // proto:  const QItemSelection QItemSelectionModel::selection();
extern void _ZNK19QItemSelectionModel9selectionEv(void* qthis); // 4
  // proto:  bool QItemSelectionModel::isColumnSelected(int column, const QModelIndex & parent);
extern void _ZNK19QItemSelectionModel16isColumnSelectedEiRK11QModelIndex(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QModelIndexList QItemSelectionModel::selectedColumns(int row);
extern void _ZNK19QItemSelectionModel15selectedColumnsEi(void* qthis, int32_t arg0); // 4
  // proto:  void QItemSelectionModel::clearSelection();
extern void _ZN19QItemSelectionModel14clearSelectionEv(void* qthis); // 4
  // proto:  void QItemSelectionModel::~QItemSelectionModel();
extern void _ZN19QItemSelectionModelD2Ev(void* qthis); // 4
  // proto:  bool QItemSelectionModel::isSelected(const QModelIndex & index);
extern void _ZNK19QItemSelectionModel10isSelectedERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QItemSelectionModel::rowIntersectsSelection(int row, const QModelIndex & parent);
extern void _ZNK19QItemSelectionModel22rowIntersectsSelectionEiRK11QModelIndex(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QItemSelectionModel::QItemSelectionModel(QAbstractItemModel * model, QObject * parent);
extern void _ZN19QItemSelectionModelC2EP18QAbstractItemModelP7QObject(void* qthis, void* arg0, void* arg1); // 3
  // proto:  void QItemSelectionModel::QItemSelectionModel(QAbstractItemModel * model);
extern void _ZN19QItemSelectionModelC2EP18QAbstractItemModel(void* qthis, void* arg0); // 3
  // proto:  void QItemSelectionModel::clearCurrentIndex();
extern void _ZN19QItemSelectionModel17clearCurrentIndexEv(void* qthis); // 4
  // proto:  void QItemSelectionModel::setModel(QAbstractItemModel * model);
extern void _ZN19QItemSelectionModel8setModelEP18QAbstractItemModel(void* qthis, void* arg0); // 4
  // proto:  bool QItemSelectionModel::isRowSelected(int row, const QModelIndex & parent);
extern void _ZNK19QItemSelectionModel13isRowSelectedEiRK11QModelIndex(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QItemSelectionModel::reset();
extern void _ZN19QItemSelectionModel5resetEv(void* qthis); // 4
  // proto:  const QMetaObject * QItemSelectionModel::metaObject();
extern void _ZNK19QItemSelectionModel10metaObjectEv(void* qthis); // 4
  // proto:  void QItemSelectionModel::clear();
extern void _ZN19QItemSelectionModel5clearEv(void* qthis); // 4
  // proto:  QModelIndexList QItemSelectionModel::selectedIndexes();
extern void _ZNK19QItemSelectionModel15selectedIndexesEv(void* qthis); // 4
  // proto:  bool QItemSelectionModel::columnIntersectsSelection(int column, const QModelIndex & parent);
extern void _ZNK19QItemSelectionModel25columnIntersectsSelectionEiRK11QModelIndex(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QModelIndex QItemSelectionModel::currentIndex();
extern void _ZNK19QItemSelectionModel12currentIndexEv(void* qthis); // 4
  // proto:  QModelIndexList QItemSelectionModel::selectedRows(int column);
extern void _ZNK19QItemSelectionModel12selectedRowsEi(void* qthis, int32_t arg0); // 4
  // proto:  QAbstractItemModel * QItemSelectionModel::model();
extern void _ZN19QItemSelectionModel5modelEv(void* qthis); // 4
  // proto:  bool QItemSelectionModel::hasSelection();
extern void _ZNK19QItemSelectionModel12hasSelectionEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QItemSelectionRange)=16
type QItemSelectionRange struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QItemSelectionModel)=1
type QItemSelectionModel struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _currentRowChanged QItemSelectionModel_currentRowChanged_signal;
//  _currentColumnChanged QItemSelectionModel_currentColumnChanged_signal;
//  _modelChanged QItemSelectionModel_modelChanged_signal;
//  _selectionChanged QItemSelectionModel_selectionChanged_signal;
//  _currentChanged QItemSelectionModel_currentChanged_signal;
}

// contains(const class QModelIndex &)
func (this *QItemSelection) contains(args ...interface{}) () {
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK14QItemSelection8containsERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QItemSelection", "contains", args)
  }

}

// indexes()
func (this *QItemSelection) indexes(args ...interface{}) () {
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
    C._ZNK14QItemSelection7indexesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelection", "indexes", args)
  }

}

// QItemSelection(const class QModelIndex &, const class QModelIndex &)
func NewQItemSelection(args ...interface{}) QItemSelection {
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QModelIndex).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN14QItemSelectionC2ERK11QModelIndexS2_(qthis, arg0, arg1)
  case 1:
    // invoke: _ZN14QItemSelectionC1Ev
    // invoke: void QItemSelection()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN14QItemSelectionC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QItemSelection", "QItemSelection", args)
  }

  return QItemSelection{}
}

// split(const class QItemSelectionRange &, const class QItemSelectionRange &, class QItemSelection *)
func (this *QItemSelection) split_s(args ...interface{}) () {
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
    var arg0 = args[0].(QItemSelectionRange).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QItemSelectionRange).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QItemSelection).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN14QItemSelection5splitERK19QItemSelectionRangeS2_PS_(arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QItemSelection", "split", args)
  }

}

// select(const class QModelIndex &, const class QModelIndex &)
func (this *QItemSelection) select_(args ...interface{}) () {
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QModelIndex).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN14QItemSelection6selectERK11QModelIndexS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QItemSelection", "select", args)
  }

}

// topLeft()
func (this *QItemSelectionRange) topLeft(args ...interface{}) () {
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
    C._ZNK19QItemSelectionRange7topLeftEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "topLeft", args)
  }

}

// right()
func (this *QItemSelectionRange) right(args ...interface{}) () {
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
    C._ZNK19QItemSelectionRange5rightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "right", args)
  }

}

// parent()
func (this *QItemSelectionRange) parent(args ...interface{}) () {
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
    C._ZNK19QItemSelectionRange6parentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "parent", args)
  }

}

// bottom()
func (this *QItemSelectionRange) bottom(args ...interface{}) () {
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
    C._ZNK19QItemSelectionRange6bottomEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "bottom", args)
  }

}

// isValid()
func (this *QItemSelectionRange) isValid(args ...interface{}) () {
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
    C._ZNK19QItemSelectionRange7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "isValid", args)
  }

}

// top()
func (this *QItemSelectionRange) top(args ...interface{}) () {
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
    C._ZNK19QItemSelectionRange3topEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "top", args)
  }

}

// contains(const class QModelIndex &)
func (this *QItemSelectionRange) contains(args ...interface{}) () {
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK19QItemSelectionRange8containsERK11QModelIndex(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK19QItemSelectionRange8containsEiiRK11QModelIndex
    // invoke: bool contains(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C._ZNK19QItemSelectionRange8containsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "contains", args)
  }

}

// QItemSelectionRange(const class QModelIndex &)
func NewQItemSelectionRange(args ...interface{}) QItemSelectionRange {
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN19QItemSelectionRangeC2ERK11QModelIndex(qthis, arg0)
  case 1:
    // invoke: _ZN19QItemSelectionRangeC1Ev
    // invoke: void QItemSelectionRange()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN19QItemSelectionRangeC2Ev(qthis)
  case 2:
    // invoke: _ZN19QItemSelectionRangeC1ERK11QModelIndexS2_
    // invoke: void QItemSelectionRange(const class QModelIndex &, const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QModelIndex).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN19QItemSelectionRangeC2ERK11QModelIndexS2_(qthis, arg0, arg1)
  case 3:
    // invoke: _ZN19QItemSelectionRangeC1ERKS_
    // invoke: void QItemSelectionRange(const class QItemSelectionRange &)
    var arg0 = args[0].(QItemSelectionRange).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN19QItemSelectionRangeC2ERKS_(qthis, arg0)
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "QItemSelectionRange", args)
  }

  return QItemSelectionRange{}
}

// indexes()
func (this *QItemSelectionRange) indexes(args ...interface{}) () {
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
    C._ZNK19QItemSelectionRange7indexesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "indexes", args)
  }

}

// width()
func (this *QItemSelectionRange) width(args ...interface{}) () {
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
    C._ZNK19QItemSelectionRange5widthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "width", args)
  }

}

// bottomRight()
func (this *QItemSelectionRange) bottomRight(args ...interface{}) () {
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
    C._ZNK19QItemSelectionRange11bottomRightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "bottomRight", args)
  }

}

// isEmpty()
func (this *QItemSelectionRange) isEmpty(args ...interface{}) () {
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
    C._ZNK19QItemSelectionRange7isEmptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "isEmpty", args)
  }

}

// intersects(const class QItemSelectionRange &)
func (this *QItemSelectionRange) intersects(args ...interface{}) () {
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
    var arg0 = args[0].(QItemSelectionRange).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK19QItemSelectionRange10intersectsERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "intersects", args)
  }

}

// model()
func (this *QItemSelectionRange) model(args ...interface{}) () {
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
    C._ZNK19QItemSelectionRange5modelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "model", args)
  }

}

// height()
func (this *QItemSelectionRange) height(args ...interface{}) () {
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
    C._ZNK19QItemSelectionRange6heightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "height", args)
  }

}

// intersected(const class QItemSelectionRange &)
func (this *QItemSelectionRange) intersected(args ...interface{}) () {
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
    var arg0 = args[0].(QItemSelectionRange).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK19QItemSelectionRange11intersectedERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "intersected", args)
  }

}

// left()
func (this *QItemSelectionRange) left(args ...interface{}) () {
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
    C._ZNK19QItemSelectionRange4leftEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "left", args)
  }

}

// selection()
func (this *QItemSelectionModel) selection(args ...interface{}) () {
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
    C._ZNK19QItemSelectionModel9selectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "selection", args)
  }

}

// isColumnSelected(int, const class QModelIndex &)
func (this *QItemSelectionModel) isColumnSelected(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QModelIndex).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK19QItemSelectionModel16isColumnSelectedEiRK11QModelIndex(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "isColumnSelected", args)
  }

}

// selectedColumns(int)
func (this *QItemSelectionModel) selectedColumns(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK19QItemSelectionModel15selectedColumnsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "selectedColumns", args)
  }

}

// clearSelection()
func (this *QItemSelectionModel) clearSelection(args ...interface{}) () {
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
    C._ZN19QItemSelectionModel14clearSelectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "clearSelection", args)
  }

}

// ~QItemSelectionModel()
func (this *QItemSelectionModel) FreeQItemSelectionModel(args ...interface{}) () {
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
    C._ZN19QItemSelectionModelD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "~QItemSelectionModel", args)
  }

}

// isSelected(const class QModelIndex &)
func (this *QItemSelectionModel) isSelected(args ...interface{}) () {
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK19QItemSelectionModel10isSelectedERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "isSelected", args)
  }

}

// rowIntersectsSelection(int, const class QModelIndex &)
func (this *QItemSelectionModel) rowIntersectsSelection(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QModelIndex).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK19QItemSelectionModel22rowIntersectsSelectionEiRK11QModelIndex(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "rowIntersectsSelection", args)
  }

}

// QItemSelectionModel(class QAbstractItemModel *, class QObject *)
func NewQItemSelectionModel(args ...interface{}) QItemSelectionModel {
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
    var arg0 = args[0].(QAbstractItemModel).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN19QItemSelectionModelC2EP18QAbstractItemModelP7QObject(qthis, arg0, arg1)
  case 1:
    // invoke: _ZN19QItemSelectionModelC1EP18QAbstractItemModel
    // invoke: void QItemSelectionModel(class QAbstractItemModel *)
    var arg0 = args[0].(QAbstractItemModel).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN19QItemSelectionModelC2EP18QAbstractItemModel(qthis, arg0)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "QItemSelectionModel", args)
  }

  return QItemSelectionModel{}
}

// clearCurrentIndex()
func (this *QItemSelectionModel) clearCurrentIndex(args ...interface{}) () {
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
    C._ZN19QItemSelectionModel17clearCurrentIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "clearCurrentIndex", args)
  }

}

// setModel(class QAbstractItemModel *)
func (this *QItemSelectionModel) setModel(args ...interface{}) () {
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
    var arg0 = args[0].(QAbstractItemModel).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN19QItemSelectionModel8setModelEP18QAbstractItemModel(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "setModel", args)
  }

}

// isRowSelected(int, const class QModelIndex &)
func (this *QItemSelectionModel) isRowSelected(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QModelIndex).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK19QItemSelectionModel13isRowSelectedEiRK11QModelIndex(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "isRowSelected", args)
  }

}

// reset()
func (this *QItemSelectionModel) reset(args ...interface{}) () {
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
    C._ZN19QItemSelectionModel5resetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "reset", args)
  }

}

// metaObject()
func (this *QItemSelectionModel) metaObject(args ...interface{}) () {
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
    C._ZNK19QItemSelectionModel10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "metaObject", args)
  }

}

// clear()
func (this *QItemSelectionModel) clear(args ...interface{}) () {
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
    C._ZN19QItemSelectionModel5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "clear", args)
  }

}

// selectedIndexes()
func (this *QItemSelectionModel) selectedIndexes(args ...interface{}) () {
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
    C._ZNK19QItemSelectionModel15selectedIndexesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "selectedIndexes", args)
  }

}

// columnIntersectsSelection(int, const class QModelIndex &)
func (this *QItemSelectionModel) columnIntersectsSelection(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QModelIndex).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK19QItemSelectionModel25columnIntersectsSelectionEiRK11QModelIndex(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "columnIntersectsSelection", args)
  }

}

// currentIndex()
func (this *QItemSelectionModel) currentIndex(args ...interface{}) () {
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
    C._ZNK19QItemSelectionModel12currentIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "currentIndex", args)
  }

}

// selectedRows(int)
func (this *QItemSelectionModel) selectedRows(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK19QItemSelectionModel12selectedRowsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "selectedRows", args)
  }

}

// model()
func (this *QItemSelectionModel) model(args ...interface{}) () {
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
    C._ZN19QItemSelectionModel5modelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "model", args)
  }

}

// hasSelection()
func (this *QItemSelectionModel) hasSelection(args ...interface{}) () {
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
    C._ZNK19QItemSelectionModel12hasSelectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "hasSelection", args)
  }

}

// <= body block end

