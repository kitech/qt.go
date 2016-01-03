package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtCore/qjsonarray.h
// dst-file: /src/core/qjsonarray.go
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
  // proto:  QJsonValue QJsonArray::first();
extern void _ZNK10QJsonArray5firstEv(void* qthis);
  // proto:  bool QJsonArray::empty();
extern void demth_ZNK10QJsonArray5emptyEv(void* qthis);
  // proto:  QJsonValue QJsonArray::takeAt(int i);
extern void _ZN10QJsonArray6takeAtEi(void* qthis, int32_t arg0);
  // proto:  void QJsonArray::removeLast();
extern void demth_ZN10QJsonArray10removeLastEv(void* qthis);
  // proto:  void QJsonArray::pop_front();
extern void demth_ZN10QJsonArray9pop_frontEv(void* qthis);
  // proto:  QVariantList QJsonArray::toVariantList();
extern void _ZNK10QJsonArray13toVariantListEv(void* qthis);
  // proto:  void QJsonArray::~QJsonArray();
extern void _ZN10QJsonArrayD0Ev(void* qthis);
  // proto:  int QJsonArray::size();
extern void _ZNK10QJsonArray4sizeEv(void* qthis);
  // proto:  int QJsonArray::count();
extern void demth_ZNK10QJsonArray5countEv(void* qthis);
  // proto:  void QJsonArray::QJsonArray();
extern void* dector_ZN10QJsonArrayC1Ev();
extern void _ZN10QJsonArrayC1Ev(void* qthis);
  // proto:  QJsonValue QJsonArray::at(int i);
extern void _ZNK10QJsonArray2atEi(void* qthis, int32_t arg0);
  // proto:  void QJsonArray::pop_back();
extern void demth_ZN10QJsonArray8pop_backEv(void* qthis);
  // proto:  bool QJsonArray::isEmpty();
extern void _ZNK10QJsonArray7isEmptyEv(void* qthis);
  // proto: static QJsonArray QJsonArray::fromStringList(const QStringList & list);
extern void _ZN10QJsonArray14fromStringListERK11QStringList(void* arg0);
  // proto:  QJsonValue QJsonArray::last();
extern void _ZNK10QJsonArray4lastEv(void* qthis);
  // proto:  void QJsonArray::removeFirst();
extern void demth_ZN10QJsonArray11removeFirstEv(void* qthis);
  // proto:  void QJsonArray::removeAt(int i);
extern void _ZN10QJsonArray8removeAtEi(void* qthis, int32_t arg0);
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

// class sizeof(QJsonArray)=16
type QJsonArray struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  QJsonValue QJsonArray::first();
func (this *QJsonArray) first(args ...interface{}) () {
  // first()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonArray5firstEv
    // invoke: QJsonValue first()
    C._ZNK10QJsonArray5firstEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "first", args)
  }

}

  // proto:  bool QJsonArray::empty();
func (this *QJsonArray) empty(args ...interface{}) () {
  // empty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonArray5emptyEv
    // invoke: bool empty()
    C.demth_ZNK10QJsonArray5emptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "empty", args)
  }

}

  // proto:  QJsonValue QJsonArray::takeAt(int i);
func (this *QJsonArray) takeAt(args ...interface{}) () {
  // takeAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QJsonArray6takeAtEi
    // invoke: QJsonValue takeAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QJsonArray6takeAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QJsonArray", "takeAt", args)
  }

}

  // proto:  void QJsonArray::removeLast();
func (this *QJsonArray) removeLast(args ...interface{}) () {
  // removeLast()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QJsonArray10removeLastEv
    // invoke: void removeLast()
    C.demth_ZN10QJsonArray10removeLastEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "removeLast", args)
  }

}

  // proto:  void QJsonArray::pop_front();
func (this *QJsonArray) pop_front(args ...interface{}) () {
  // pop_front()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QJsonArray9pop_frontEv
    // invoke: void pop_front()
    C.demth_ZN10QJsonArray9pop_frontEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "pop_front", args)
  }

}

  // proto:  QVariantList QJsonArray::toVariantList();
func (this *QJsonArray) toVariantList(args ...interface{}) () {
  // toVariantList()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonArray13toVariantListEv
    // invoke: QVariantList toVariantList()
    C._ZNK10QJsonArray13toVariantListEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "toVariantList", args)
  }

}

  // proto:  void QJsonArray::~QJsonArray();
func (this *QJsonArray) FreeQJsonArray(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QJsonArray", "~QJsonArray", args)
  }

}

  // proto:  int QJsonArray::size();
func (this *QJsonArray) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonArray4sizeEv
    // invoke: int size()
    C._ZNK10QJsonArray4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "size", args)
  }

}

  // proto:  int QJsonArray::count();
func (this *QJsonArray) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonArray5countEv
    // invoke: int count()
    C.demth_ZNK10QJsonArray5countEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "count", args)
  }

}

  // proto:  void QJsonArray::QJsonArray();
func NewQJsonArray(args ...interface{}) QJsonArray {
  return QJsonArray{}
}

  // proto:  QJsonValue QJsonArray::at(int i);
func (this *QJsonArray) at(args ...interface{}) () {
  // at(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonArray2atEi
    // invoke: QJsonValue at(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK10QJsonArray2atEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QJsonArray", "at", args)
  }

}

  // proto:  void QJsonArray::pop_back();
func (this *QJsonArray) pop_back(args ...interface{}) () {
  // pop_back()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QJsonArray8pop_backEv
    // invoke: void pop_back()
    C.demth_ZN10QJsonArray8pop_backEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "pop_back", args)
  }

}

  // proto:  bool QJsonArray::isEmpty();
func (this *QJsonArray) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonArray7isEmptyEv
    // invoke: bool isEmpty()
    C._ZNK10QJsonArray7isEmptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "isEmpty", args)
  }

}

  // proto: static QJsonArray QJsonArray::fromStringList(const QStringList & list);
func (this *QJsonArray) fromStringList_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QJsonArray", "fromStringList", args)
  }

}

  // proto:  QJsonValue QJsonArray::last();
func (this *QJsonArray) last(args ...interface{}) () {
  // last()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonArray4lastEv
    // invoke: QJsonValue last()
    C._ZNK10QJsonArray4lastEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "last", args)
  }

}

  // proto:  void QJsonArray::removeFirst();
func (this *QJsonArray) removeFirst(args ...interface{}) () {
  // removeFirst()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QJsonArray11removeFirstEv
    // invoke: void removeFirst()
    C.demth_ZN10QJsonArray11removeFirstEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "removeFirst", args)
  }

}

  // proto:  void QJsonArray::removeAt(int i);
func (this *QJsonArray) removeAt(args ...interface{}) () {
  // removeAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QJsonArray8removeAtEi
    // invoke: void removeAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QJsonArray8removeAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QJsonArray", "removeAt", args)
  }

}

// <= body block end

