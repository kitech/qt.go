package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtCore/qstringlistmodel.h
// dst-file: /src/core/qstringlistmodel.go
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
  // proto:  void QStringListModel::QStringListModel(const QStringList & strings, QObject * parent);
extern void* dector_ZN16QStringListModelC1ERK11QStringListP7QObject(void* arg0, void* arg1);
extern void _ZN16QStringListModelC1ERK11QStringListP7QObject(void* qthis, void* arg0, void* arg1);
  // proto:  bool QStringListModel::insertRows(int row, int count, const QModelIndex & parent);
extern void _ZN16QStringListModel10insertRowsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2);
  // proto:  QVariant QStringListModel::data(const QModelIndex & index, int role);
extern void _ZNK16QStringListModel4dataERK11QModelIndexi(void* qthis, void* arg0, int32_t arg1);
  // proto:  QStringList QStringListModel::stringList();
extern void _ZNK16QStringListModel10stringListEv(void* qthis);
  // proto:  const QMetaObject * QStringListModel::metaObject();
extern void _ZNK16QStringListModel10metaObjectEv(void* qthis);
  // proto:  bool QStringListModel::removeRows(int row, int count, const QModelIndex & parent);
extern void _ZN16QStringListModel10removeRowsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2);
  // proto:  void QStringListModel::QStringListModel(QObject * parent);
extern void* dector_ZN16QStringListModelC1EP7QObject(void* arg0);
extern void _ZN16QStringListModelC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QStringListModel::QStringListModel(const QStringListModel & );
extern void* dector_ZN16QStringListModelC1ERKS_(void* arg0);
extern void _ZN16QStringListModelC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QStringListModel::setData(const QModelIndex & index, const QVariant & value, int role);
extern void _ZN16QStringListModel7setDataERK11QModelIndexRK8QVarianti(void* qthis, void* arg0, void* arg1, int32_t arg2);
  // proto:  int QStringListModel::rowCount(const QModelIndex & parent);
extern void _ZNK16QStringListModel8rowCountERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QModelIndex QStringListModel::sibling(int row, int column, const QModelIndex & idx);
extern void _ZNK16QStringListModel7siblingEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2);
  // proto:  void QStringListModel::setStringList(const QStringList & strings);
extern void _ZN16QStringListModel13setStringListERK11QStringList(void* qthis, void* arg0);
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

// class sizeof(QStringListModel)=1
type QStringListModel struct {
  /*qbase*/ QAbstractListModel;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  void QStringListModel::QStringListModel(const QStringList & strings, QObject * parent);
func NewQStringListModel(args ...interface{}) QStringListModel {
  return QStringListModel{}
}

  // proto:  bool QStringListModel::insertRows(int row, int count, const QModelIndex & parent);
func (this *QStringListModel) insertRows(args ...interface{}) () {
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
    // invoke: _ZN16QStringListModel10insertRowsEiiRK11QModelIndex
    // invoke: bool insertRows(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN16QStringListModel10insertRowsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStringListModel", "insertRows", args)
  }

}

  // proto:  QVariant QStringListModel::data(const QModelIndex & index, int role);
func (this *QStringListModel) data(args ...interface{}) () {
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
    // invoke: _ZNK16QStringListModel4dataERK11QModelIndexi
    // invoke: QVariant data(const class QModelIndex &, int)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK16QStringListModel4dataERK11QModelIndexi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStringListModel", "data", args)
  }

}

  // proto:  QStringList QStringListModel::stringList();
func (this *QStringListModel) stringList(args ...interface{}) () {
  // stringList()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QStringListModel10stringListEv
    // invoke: QStringList stringList()
    C._ZNK16QStringListModel10stringListEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringListModel", "stringList", args)
  }

}

  // proto:  const QMetaObject * QStringListModel::metaObject();
func (this *QStringListModel) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QStringListModel10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK16QStringListModel10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringListModel", "metaObject", args)
  }

}

  // proto:  bool QStringListModel::removeRows(int row, int count, const QModelIndex & parent);
func (this *QStringListModel) removeRows(args ...interface{}) () {
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
    // invoke: _ZN16QStringListModel10removeRowsEiiRK11QModelIndex
    // invoke: bool removeRows(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN16QStringListModel10removeRowsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStringListModel", "removeRows", args)
  }

}

  // proto:  bool QStringListModel::setData(const QModelIndex & index, const QVariant & value, int role);
func (this *QStringListModel) setData(args ...interface{}) () {
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
    // invoke: _ZN16QStringListModel7setDataERK11QModelIndexRK8QVarianti
    // invoke: bool setData(const class QModelIndex &, const class QVariant &, int)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN16QStringListModel7setDataERK11QModelIndexRK8QVarianti(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStringListModel", "setData", args)
  }

}

  // proto:  int QStringListModel::rowCount(const QModelIndex & parent);
func (this *QStringListModel) rowCount(args ...interface{}) () {
  // rowCount(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QStringListModel8rowCountERK11QModelIndex
    // invoke: int rowCount(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK16QStringListModel8rowCountERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStringListModel", "rowCount", args)
  }

}

  // proto:  QModelIndex QStringListModel::sibling(int row, int column, const QModelIndex & idx);
func (this *QStringListModel) sibling(args ...interface{}) () {
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
    // invoke: _ZNK16QStringListModel7siblingEiiRK11QModelIndex
    // invoke: QModelIndex sibling(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C._ZNK16QStringListModel7siblingEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStringListModel", "sibling", args)
  }

}

  // proto:  void QStringListModel::setStringList(const QStringList & strings);
func (this *QStringListModel) setStringList(args ...interface{}) () {
  // setStringList(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QStringListModel13setStringListERK11QStringList
    // invoke: void setStringList(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN16QStringListModel13setStringListERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStringListModel", "setStringList", args)
  }

}

// <= body block end

