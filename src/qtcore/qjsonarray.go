package qtcore
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
extern int32_t C_ZNK10QJsonArray4sizeEv(void* qthis); // 4
  // proto:  QJsonArray::iterator QJsonArray::end();
extern void C_ZN10QJsonArray3endEv(void* qthis); // 2
  // proto:  void QJsonArray::pop_front();
extern void C_ZN10QJsonArray9pop_frontEv(void* qthis); // 2
  // proto:  bool QJsonArray::isEmpty();
extern bool C_ZNK10QJsonArray7isEmptyEv(void* qthis); // 4
  // proto:  void QJsonArray::pop_back();
extern void C_ZN10QJsonArray8pop_backEv(void* qthis); // 2
  // proto:  bool QJsonArray::empty();
extern bool C_ZNK10QJsonArray5emptyEv(void* qthis); // 2
  // proto:  QJsonArray::iterator QJsonArray::begin();
extern void C_ZN10QJsonArray5beginEv(void* qthis); // 2
  // proto:  QJsonValue QJsonArray::takeAt(int i);
extern void C_ZN10QJsonArray6takeAtEi(void* qthis, int32_t arg0); // 4
  // proto:  void QJsonArray::removeFirst();
extern void C_ZN10QJsonArray11removeFirstEv(void* qthis); // 2
  // proto: static QJsonArray QJsonArray::fromStringList(const QStringList & list);
extern void C_ZN10QJsonArray14fromStringListERK11QStringList(void* arg0); // 4
  // proto:  int QJsonArray::count();
extern int32_t C_ZNK10QJsonArray5countEv(void* qthis); // 2
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// QJsonArray()
func NewQJsonArray(args ...interface{}) *QJsonArray {
  // QJsonArray()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QJsonArrayC1Ev
    // invoke: void QJsonArray()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QJsonArrayC2Ev()
    return &QJsonArray{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QJsonArray", "QJsonArray", args)
  }

  return nil // QJsonArray{Qclsinst:qthis}
}

// constBegin()
func (this *QJsonArray) Constbegin(args ...interface{}) () {
  // constBegin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonArray10constBeginEv
    // invoke: QJsonArray::const_iterator constBegin()
    C.C_ZNK10QJsonArray10constBeginEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "constBegin", args)
  }

  return
}

// removeAt(int)
func (this *QJsonArray) Removeat(args ...interface{}) () {
  // removeAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QJsonArray8removeAtEi
    // invoke: void removeAt(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QJsonArray8removeAtEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QJsonArray", "removeAt", args)
  }

  return
}

// at(int)
func (this *QJsonArray) At(args ...interface{}) () {
  // at(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonArray2atEi
    // invoke: QJsonValue at(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK10QJsonArray2atEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QJsonArray", "at", args)
  }

  return
}

// size()
func (this *QJsonArray) Size(args ...interface{}) (ret interface{}) {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonArray4sizeEv
    // invoke: int size()
    var ret0 = C.C_ZNK10QJsonArray4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonArray", "size", args)
  }

  return
}

// end()
func (this *QJsonArray) End(args ...interface{}) () {
  // end()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QJsonArray3endEv
    // invoke: QJsonArray::iterator end()
    C.C_ZN10QJsonArray3endEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "end", args)
  }

  return
}

// pop_front()
func (this *QJsonArray) Pop_Front(args ...interface{}) () {
  // pop_front()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QJsonArray9pop_frontEv
    // invoke: void pop_front()
    C.C_ZN10QJsonArray9pop_frontEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "pop_front", args)
  }

  return
}

// isEmpty()
func (this *QJsonArray) Isempty(args ...interface{}) (ret interface{}) {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonArray7isEmptyEv
    // invoke: bool isEmpty()
    var ret0 = C.C_ZNK10QJsonArray7isEmptyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonArray", "isEmpty", args)
  }

  return
}

// pop_back()
func (this *QJsonArray) Pop_Back(args ...interface{}) () {
  // pop_back()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QJsonArray8pop_backEv
    // invoke: void pop_back()
    C.C_ZN10QJsonArray8pop_backEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "pop_back", args)
  }

  return
}

// empty()
func (this *QJsonArray) Empty(args ...interface{}) (ret interface{}) {
  // empty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonArray5emptyEv
    // invoke: bool empty()
    var ret0 = C.C_ZNK10QJsonArray5emptyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonArray", "empty", args)
  }

  return
}

// begin()
func (this *QJsonArray) Begin(args ...interface{}) () {
  // begin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QJsonArray5beginEv
    // invoke: QJsonArray::iterator begin()
    C.C_ZN10QJsonArray5beginEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "begin", args)
  }

  return
}

// takeAt(int)
func (this *QJsonArray) Takeat(args ...interface{}) () {
  // takeAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QJsonArray6takeAtEi
    // invoke: QJsonValue takeAt(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QJsonArray6takeAtEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QJsonArray", "takeAt", args)
  }

  return
}

// removeFirst()
func (this *QJsonArray) Removefirst(args ...interface{}) () {
  // removeFirst()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QJsonArray11removeFirstEv
    // invoke: void removeFirst()
    C.C_ZN10QJsonArray11removeFirstEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "removeFirst", args)
  }

  return
}

// fromStringList(const class QStringList &)
func (this *QJsonArray) Fromstringlist_S(args ...interface{}) () {
  // fromStringList(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QJsonArray14fromStringListERK11QStringList
    // invoke: QJsonArray fromStringList(const class QStringList &)
    var arg0 = args[0].(*QStringList).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QJsonArray14fromStringListERK11QStringList(arg0)
  default:
    qtrt.ErrorResolve("QJsonArray", "fromStringList", args)
  }

  return
}

// count()
func (this *QJsonArray) Count(args ...interface{}) (ret interface{}) {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonArray5countEv
    // invoke: int count()
    var ret0 = C.C_ZNK10QJsonArray5countEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonArray", "count", args)
  }

  return
}

// last()
func (this *QJsonArray) Last(args ...interface{}) () {
  // last()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonArray4lastEv
    // invoke: QJsonValue last()
    C.C_ZNK10QJsonArray4lastEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "last", args)
  }

  return
}

// toVariantList()
func (this *QJsonArray) Tovariantlist(args ...interface{}) () {
  // toVariantList()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonArray13toVariantListEv
    // invoke: QVariantList toVariantList()
    C.C_ZNK10QJsonArray13toVariantListEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "toVariantList", args)
  }

  return
}

// removeLast()
func (this *QJsonArray) Removelast(args ...interface{}) () {
  // removeLast()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QJsonArray10removeLastEv
    // invoke: void removeLast()
    C.C_ZN10QJsonArray10removeLastEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "removeLast", args)
  }

  return
}

// constEnd()
func (this *QJsonArray) Constend(args ...interface{}) () {
  // constEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonArray8constEndEv
    // invoke: QJsonArray::const_iterator constEnd()
    C.C_ZNK10QJsonArray8constEndEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "constEnd", args)
  }

  return
}

// ~QJsonArray()
func (this *QJsonArray) Freeqjsonarray(args ...interface{}) () {
  // ~QJsonArray()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QJsonArrayD0Ev
    // invoke: void ~QJsonArray()
    C.C_ZN10QJsonArrayD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "~QJsonArray", args)
  }

  return
}

// first()
func (this *QJsonArray) First(args ...interface{}) () {
  // first()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonArray5firstEv
    // invoke: QJsonValue first()
    C.C_ZNK10QJsonArray5firstEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonArray", "first", args)
  }

  return
}

// <= body block end

