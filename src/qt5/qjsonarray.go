package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QJsonArray::QJsonArray();
extern void* C_ZN10QJsonArrayC2Ev(); // 3
  // proto:  QJsonArray::const_iterator QJsonArray::constBegin();
extern void C_ZNK10QJsonArray10constBeginEv(void* qthis); // 2
  // proto:  void QJsonArray::removeAt(int i);
extern void C_ZN10QJsonArray8removeAtEi(void* qthis, int32_t arg0); // 4
  // proto:  QJsonValue QJsonArray::at(int i);
extern void C_ZNK10QJsonArray2atEi(void* qthis, int32_t arg0); // 4
  // proto:  int QJsonArray::size();
extern void C_ZNK10QJsonArray4sizeEv(void* qthis); // 4
  // proto:  QJsonArray::iterator QJsonArray::end();
extern void C_ZN10QJsonArray3endEv(void* qthis); // 2
  // proto:  void QJsonArray::pop_front();
extern void C_ZN10QJsonArray9pop_frontEv(void* qthis); // 2
  // proto:  bool QJsonArray::isEmpty();
extern void C_ZNK10QJsonArray7isEmptyEv(void* qthis); // 4
  // proto:  void QJsonArray::pop_back();
extern void C_ZN10QJsonArray8pop_backEv(void* qthis); // 2
  // proto:  bool QJsonArray::empty();
extern void C_ZNK10QJsonArray5emptyEv(void* qthis); // 2
  // proto:  QJsonArray::iterator QJsonArray::begin();
extern void C_ZN10QJsonArray5beginEv(void* qthis); // 2
  // proto:  QJsonValue QJsonArray::takeAt(int i);
extern void C_ZN10QJsonArray6takeAtEi(void* qthis, int32_t arg0); // 4
  // proto:  void QJsonArray::removeFirst();
extern void C_ZN10QJsonArray11removeFirstEv(void* qthis); // 2
  // proto: static QJsonArray QJsonArray::fromStringList(const QStringList & list);
extern void C_ZN10QJsonArray14fromStringListERK11QStringList(void* arg0); // 4
  // proto:  int QJsonArray::count();
extern void C_ZNK10QJsonArray5countEv(void* qthis); // 2
  // proto:  QJsonValue QJsonArray::last();
extern void C_ZNK10QJsonArray4lastEv(void* qthis); // 4
  // proto:  QVariantList QJsonArray::toVariantList();
extern void C_ZNK10QJsonArray13toVariantListEv(void* qthis); // 4
  // proto:  void QJsonArray::removeLast();
extern void C_ZN10QJsonArray10removeLastEv(void* qthis); // 2
  // proto:  QJsonArray::const_iterator QJsonArray::constEnd();
extern void C_ZNK10QJsonArray8constEndEv(void* qthis); // 2
  // proto:  void QJsonArray::~QJsonArray();
extern void C_ZN10QJsonArrayD2Ev(void* qthis); // 4
  // proto:  QJsonValue QJsonArray::first();
extern void C_ZNK10QJsonArray5firstEv(void* qthis); // 4
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

// QJsonArray()
func NewQJsonArray(args ...interface{}) *QJsonArray {
  // QJsonArray()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QJsonArrayC1Ev
    // invoke: void QJsonArray()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QJsonArrayC2Ev()
    return &QJsonArray{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QJsonArray", "QJsonArray", args)
  }

  return nil // QJsonArray{qclsinst:qthis}
}

// constBegin()
func (this *QJsonArray) constBegin(args ...interface{}) () {
  // constBegin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonArray10constBeginEv
    // invoke: QJsonArray::const_iterator constBegin()
    C.C_ZNK10QJsonArray10constBeginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "constBegin", args)
  }

}

// removeAt(int)
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
    C.C_ZN10QJsonArray8removeAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QJsonArray", "removeAt", args)
  }

}

// at(int)
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
    C.C_ZNK10QJsonArray2atEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QJsonArray", "at", args)
  }

}

// size()
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
    var ret = C.C_ZNK10QJsonArray4sizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonArray", "size", args)
  }

}

// end()
func (this *QJsonArray) end(args ...interface{}) () {
  // end()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QJsonArray3endEv
    // invoke: QJsonArray::iterator end()
    C.C_ZN10QJsonArray3endEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "end", args)
  }

}

// pop_front()
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
    C.C_ZN10QJsonArray9pop_frontEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "pop_front", args)
  }

}

// isEmpty()
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
    var ret = C.C_ZNK10QJsonArray7isEmptyEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonArray", "isEmpty", args)
  }

}

// pop_back()
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
    C.C_ZN10QJsonArray8pop_backEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "pop_back", args)
  }

}

// empty()
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
    var ret = C.C_ZNK10QJsonArray5emptyEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonArray", "empty", args)
  }

}

// begin()
func (this *QJsonArray) begin(args ...interface{}) () {
  // begin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QJsonArray5beginEv
    // invoke: QJsonArray::iterator begin()
    C.C_ZN10QJsonArray5beginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "begin", args)
  }

}

// takeAt(int)
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
    C.C_ZN10QJsonArray6takeAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QJsonArray", "takeAt", args)
  }

}

// removeFirst()
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
    C.C_ZN10QJsonArray11removeFirstEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "removeFirst", args)
  }

}

// fromStringList(const class QStringList &)
func (this *QJsonArray) fromStringList_s(args ...interface{}) () {
  // fromStringList(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QJsonArray14fromStringListERK11QStringList
    // invoke: QJsonArray fromStringList(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QJsonArray14fromStringListERK11QStringList(arg0)
  default:
    qtrt.ErrorResolve("QJsonArray", "fromStringList", args)
  }

}

// count()
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
    var ret = C.C_ZNK10QJsonArray5countEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonArray", "count", args)
  }

}

// last()
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
    C.C_ZNK10QJsonArray4lastEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "last", args)
  }

}

// toVariantList()
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
    C.C_ZNK10QJsonArray13toVariantListEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "toVariantList", args)
  }

}

// removeLast()
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
    C.C_ZN10QJsonArray10removeLastEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "removeLast", args)
  }

}

// constEnd()
func (this *QJsonArray) constEnd(args ...interface{}) () {
  // constEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonArray8constEndEv
    // invoke: QJsonArray::const_iterator constEnd()
    C.C_ZNK10QJsonArray8constEndEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "constEnd", args)
  }

}

// ~QJsonArray()
func (this *QJsonArray) FreeQJsonArray(args ...interface{}) () {
  // ~QJsonArray()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QJsonArrayD0Ev
    // invoke: void ~QJsonArray()
    C.C_ZN10QJsonArrayD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "~QJsonArray", args)
  }

}

// first()
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
    C.C_ZNK10QJsonArray5firstEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "first", args)
  }

}

// <= body block end

