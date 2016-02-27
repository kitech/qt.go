package qtcore
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtCore/qbitarray.h
// dst-file: /src/core/qbitarray.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "runtime"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  int QBitArray::size();
extern int32_t C_ZNK9QBitArray4sizeEv(void* qthis); // 2
  // proto:  bool QBitArray::testBit(int i);
extern bool C_ZNK9QBitArray7testBitEi(void* qthis, int32_t arg0); // 2
  // proto:  bool QBitArray::at(int i);
extern bool C_ZNK9QBitArray2atEi(void* qthis, int32_t arg0); // 2
  // proto:  bool QBitArray::toggleBit(int i);
extern bool C_ZN9QBitArray9toggleBitEi(void* qthis, int32_t arg0); // 2
  // proto:  void QBitArray::fill(bool val, int first, int last);
extern void C_ZN9QBitArray4fillEbii(void* qthis, bool arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  bool QBitArray::fill(bool val, int size);
extern bool C_ZN9QBitArray4fillEbi(void* qthis, bool arg0, int32_t arg1); // 2
  // proto:  void QBitArray::clearBit(int i);
extern void C_ZN9QBitArray8clearBitEi(void* qthis, int32_t arg0); // 2
  // proto:  void QBitArray::QBitArray();
extern void* C_ZN9QBitArrayC2Ev(); // 1
  // proto:  void QBitArray::QBitArray(int size, bool val);
extern void* C_ZN9QBitArrayC2Eib(int32_t arg0, bool arg1); // 3
  // proto:  void QBitArray::QBitArray(const QBitArray & other);
extern void* C_ZN9QBitArrayC2ERKS_(void* arg0); // 1
  // proto:  void QBitArray::swap(QBitArray & other);
extern void C_ZN9QBitArray4swapERS_(void* qthis, void* arg0); // 2
  // proto:  bool QBitArray::isEmpty();
extern bool C_ZNK9QBitArray7isEmptyEv(void* qthis); // 2
  // proto:  void QBitArray::setBit(int i, bool val);
extern void C_ZN9QBitArray6setBitEib(void* qthis, int32_t arg0, bool arg1); // 2
  // proto:  void QBitArray::setBit(int i);
extern void C_ZN9QBitArray6setBitEi(void* qthis, int32_t arg0); // 2
  // proto:  bool QBitArray::isDetached();
extern bool C_ZNK9QBitArray10isDetachedEv(void* qthis); // 2
  // proto:  void QBitArray::truncate(int pos);
extern void C_ZN9QBitArray8truncateEi(void* qthis, int32_t arg0); // 2
  // proto:  void QBitArray::detach();
extern void C_ZN9QBitArray6detachEv(void* qthis); // 2
  // proto:  void QBitArray::resize(int size);
extern void C_ZN9QBitArray6resizeEi(void* qthis, int32_t arg0); // 4
  // proto:  int QBitArray::count(bool on);
extern int32_t C_ZNK9QBitArray5countEb(void* qthis, bool arg0); // 4
  // proto:  int QBitArray::count();
extern int32_t C_ZNK9QBitArray5countEv(void* qthis); // 2
  // proto:  void QBitArray::clear();
extern void C_ZN9QBitArray5clearEv(void* qthis); // 2
  // proto:  bool QBitArray::isNull();
extern bool C_ZNK9QBitArray6isNullEv(void* qthis); // 2
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QBitRef)=16
type QBitRef struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QBitArray)=8
type QBitArray struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// size()
func (this *QBitArray) Size(args ...interface{}) (ret interface{}) {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QBitArray4sizeEv
    // invoke: int size()
    var ret0 = C.C_ZNK9QBitArray4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBitArray", "size", args)
  }

  return
}

// testBit(int)
func (this *QBitArray) TestBit(args ...interface{}) (ret interface{}) {
  // testBit(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QBitArray7testBitEi
    // invoke: bool testBit(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QBitArray7testBitEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBitArray", "testBit", args)
  }

  return
}

// at(int)
func (this *QBitArray) At(args ...interface{}) (ret interface{}) {
  // at(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QBitArray2atEi
    // invoke: bool at(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QBitArray2atEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBitArray", "at", args)
  }

  return
}

// toggleBit(int)
func (this *QBitArray) ToggleBit(args ...interface{}) (ret interface{}) {
  // toggleBit(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArray9toggleBitEi
    // invoke: bool toggleBit(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QBitArray9toggleBitEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBitArray", "toggleBit", args)
  }

  return
}

// fill(_Bool, int, int)
func (this *QBitArray) Fill(args ...interface{}) () {
  // fill(_Bool, int, int)
  // fill(_Bool, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.BoolTy(false) // "bool"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArray4fillEbii
    // invoke: void fill(_Bool, int, int)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN9QBitArray4fillEbii(this.Qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN9QBitArray4fillEbi
    // invoke: bool fill(_Bool, int)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN9QBitArray4fillEbi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
  default:
    qtrt.ErrorResolve("QBitArray", "fill", args)
  }

  return
}

// clearBit(int)
func (this *QBitArray) ClearBit(args ...interface{}) () {
  // clearBit(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArray8clearBitEi
    // invoke: void clearBit(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QBitArray8clearBitEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBitArray", "clearBit", args)
  }

  return
}

// QBitArray()
func GcfreeQBitArray(this *QBitArray) {
  qtrt.UniverseFree(this)
}
func NewQBitArray(args ...interface{}) *QBitArray {
  // QBitArray()
  // QBitArray(int, _Bool)
  // QBitArray(const class QBitArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.BoolTy(false) // "bool"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QBitArray{}) // "const QBitArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArrayC1Ev
    // invoke: void QBitArray()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QBitArrayC2Ev()
    this := &QBitArray{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQBitArray)
    return this
  case 1:
    // invoke: _ZN9QBitArrayC1Eib
    // invoke: void QBitArray(int, _Bool)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QBitArrayC2Eib(arg0, arg1)
    this := &QBitArray{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQBitArray)
    return this
  case 2:
    // invoke: _ZN9QBitArrayC1ERKS_
    // invoke: void QBitArray(const class QBitArray &)
    var arg0 = args[0].(*QBitArray).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QBitArrayC2ERKS_(arg0)
    this := &QBitArray{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQBitArray)
    return this
  default:
    qtrt.ErrorResolve("QBitArray", "QBitArray", args)
  }

  return nil // QBitArray{Qclsinst:qthis}
}

// swap(class QBitArray &)
func (this *QBitArray) Swap(args ...interface{}) () {
  // swap(class QBitArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBitArray{}) // "QBitArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArray4swapERS_
    // invoke: void swap(class QBitArray &)
    var arg0 = args[0].(*QBitArray).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QBitArray4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBitArray", "swap", args)
  }

  return
}

// isEmpty()
func (this *QBitArray) IsEmpty(args ...interface{}) (ret interface{}) {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QBitArray7isEmptyEv
    // invoke: bool isEmpty()
    var ret0 = C.C_ZNK9QBitArray7isEmptyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBitArray", "isEmpty", args)
  }

  return
}

// setBit(int, _Bool)
func (this *QBitArray) SetBit(args ...interface{}) () {
  // setBit(int, _Bool)
  // setBit(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArray6setBitEib
    // invoke: void setBit(int, _Bool)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN9QBitArray6setBitEib(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN9QBitArray6setBitEi
    // invoke: void setBit(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QBitArray6setBitEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBitArray", "setBit", args)
  }

  return
}

// isDetached()
func (this *QBitArray) IsDetached(args ...interface{}) (ret interface{}) {
  // isDetached()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QBitArray10isDetachedEv
    // invoke: bool isDetached()
    var ret0 = C.C_ZNK9QBitArray10isDetachedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBitArray", "isDetached", args)
  }

  return
}

// truncate(int)
func (this *QBitArray) Truncate(args ...interface{}) () {
  // truncate(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArray8truncateEi
    // invoke: void truncate(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QBitArray8truncateEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBitArray", "truncate", args)
  }

  return
}

// detach()
func (this *QBitArray) Detach(args ...interface{}) () {
  // detach()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArray6detachEv
    // invoke: void detach()
    C.C_ZN9QBitArray6detachEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QBitArray", "detach", args)
  }

  return
}

// resize(int)
func (this *QBitArray) Resize(args ...interface{}) () {
  // resize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArray6resizeEi
    // invoke: void resize(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QBitArray6resizeEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBitArray", "resize", args)
  }

  return
}

// count(_Bool)
func (this *QBitArray) Count(args ...interface{}) (ret interface{}) {
  // count(_Bool)
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QBitArray5countEb
    // invoke: int count(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QBitArray5countEb(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK9QBitArray5countEv
    // invoke: int count()
    var ret0 = C.C_ZNK9QBitArray5countEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBitArray", "count", args)
  }

  return
}

// clear()
func (this *QBitArray) Clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArray5clearEv
    // invoke: void clear()
    C.C_ZN9QBitArray5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QBitArray", "clear", args)
  }

  return
}

// isNull()
func (this *QBitArray) IsNull(args ...interface{}) (ret interface{}) {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QBitArray6isNullEv
    // invoke: bool isNull()
    var ret0 = C.C_ZNK9QBitArray6isNullEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBitArray", "isNull", args)
  }

  return
}

// <= body block end

