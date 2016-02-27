package qtcore
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtCore/qlist.h
// dst-file: /src/core/qlist.go
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
  // proto:  void ** QListData::insert(int i);
extern void* C_ZN9QListData6insertEi(void* qthis, int32_t arg0); // 4
  // proto:  void ** QListData::begin();
extern void* C_ZNK9QListData5beginEv(void* qthis); // 2
  // proto:  void ** QListData::end();
extern void* C_ZNK9QListData3endEv(void* qthis); // 2
  // proto:  void QListData::realloc(int alloc);
extern void C_ZN9QListData7reallocEi(void* qthis, int32_t arg0); // 4
  // proto:  void QListData::move(int from, int to);
extern void C_ZN9QListData4moveEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QListData::dispose();
extern void C_ZN9QListData7disposeEv(void* qthis); // 2
  // proto:  void QListData::remove(int i);
extern void C_ZN9QListData6removeEi(void* qthis, int32_t arg0); // 4
  // proto:  void QListData::remove(int i, int n);
extern void C_ZN9QListData6removeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void ** QListData::prepend();
extern void* C_ZN9QListData7prependEv(void* qthis); // 4
  // proto:  void ** QListData::erase(void ** xi);
extern void* C_ZN9QListData5eraseEPPv(void* qthis, void* arg0); // 4
  // proto:  bool QListData::isEmpty();
extern bool C_ZNK9QListData7isEmptyEv(void* qthis); // 2
  // proto:  void ** QListData::at(int i);
extern void* C_ZNK9QListData2atEi(void* qthis, int32_t arg0); // 2
  // proto:  QListData::Data * QListData::detach_grow(int * i, int n);
extern void C_ZN9QListData11detach_growEPii(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QListData::Data * QListData::detach(int alloc);
extern void C_ZN9QListData6detachEi(void* qthis, int32_t arg0); // 4
  // proto:  void ** QListData::append(int n);
extern void* C_ZN9QListData6appendEi(void* qthis, int32_t arg0); // 4
  // proto:  void ** QListData::append(const QListData & l);
extern void* C_ZN9QListData6appendERKS_(void* qthis, void* arg0); // 4
  // proto:  void ** QListData::append();
extern void* C_ZN9QListData6appendEv(void* qthis); // 4
  // proto:  int QListData::size();
extern int32_t C_ZNK9QListData4sizeEv(void* qthis); // 2
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

// class sizeof(QListData)=8
type QListData struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// insert(int)
func (this *QListData) Insert(args ...interface{}) (ret interface{}) {
  // insert(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListData6insertEi
    // invoke: void ** insert(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QListData6insertEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.VoidpTy() // "void **"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListData", "insert", args)
  }

  return
}

// begin()
func (this *QListData) Begin(args ...interface{}) (ret interface{}) {
  // begin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListData5beginEv
    // invoke: void ** begin()
    var ret0 = C.C_ZNK9QListData5beginEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.VoidpTy() // "void **"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListData", "begin", args)
  }

  return
}

// end()
func (this *QListData) End(args ...interface{}) (ret interface{}) {
  // end()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListData3endEv
    // invoke: void ** end()
    var ret0 = C.C_ZNK9QListData3endEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.VoidpTy() // "void **"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListData", "end", args)
  }

  return
}

// realloc(int)
func (this *QListData) Realloc(args ...interface{}) () {
  // realloc(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListData7reallocEi
    // invoke: void realloc(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QListData7reallocEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListData", "realloc", args)
  }

  return
}

// move(int, int)
func (this *QListData) Move_(args ...interface{}) () {
  // move(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListData4moveEii
    // invoke: void move(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN9QListData4moveEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QListData", "move", args)
  }

  return
}

// dispose()
func (this *QListData) Dispose(args ...interface{}) () {
  // dispose()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListData7disposeEv
    // invoke: void dispose()
    C.C_ZN9QListData7disposeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QListData", "dispose", args)
  }

  return
}

// remove(int)
func (this *QListData) Remove(args ...interface{}) () {
  // remove(int)
  // remove(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListData6removeEi
    // invoke: void remove(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QListData6removeEi(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN9QListData6removeEii
    // invoke: void remove(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN9QListData6removeEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QListData", "remove", args)
  }

  return
}

// prepend()
func (this *QListData) Prepend(args ...interface{}) (ret interface{}) {
  // prepend()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListData7prependEv
    // invoke: void ** prepend()
    var ret0 = C.C_ZN9QListData7prependEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.VoidpTy() // "void **"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListData", "prepend", args)
  }

  return
}

// erase(void **)
func (this *QListData) Erase(args ...interface{}) (ret interface{}) {
  // erase(void **)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.VoidpTy() // "void **"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListData5eraseEPPv
    // invoke: void ** erase(void **)
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QListData5eraseEPPv(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.VoidpTy() // "void **"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListData", "erase", args)
  }

  return
}

// isEmpty()
func (this *QListData) IsEmpty(args ...interface{}) (ret interface{}) {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListData7isEmptyEv
    // invoke: bool isEmpty()
    var ret0 = C.C_ZNK9QListData7isEmptyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListData", "isEmpty", args)
  }

  return
}

// at(int)
func (this *QListData) At(args ...interface{}) (ret interface{}) {
  // at(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListData2atEi
    // invoke: void ** at(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QListData2atEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.VoidpTy() // "void **"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListData", "at", args)
  }

  return
}

// detach_grow(int *, int)
func (this *QListData) Detach_grow(args ...interface{}) () {
  // detach_grow(int *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "int *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListData11detach_growEPii
    // invoke: QListData::Data * detach_grow(int *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN9QListData11detach_growEPii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QListData", "detach_grow", args)
  }

  return
}

// detach(int)
func (this *QListData) Detach(args ...interface{}) () {
  // detach(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListData6detachEi
    // invoke: QListData::Data * detach(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QListData6detachEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListData", "detach", args)
  }

  return
}

// append(int)
func (this *QListData) Append(args ...interface{}) (ret interface{}) {
  // append(int)
  // append(const struct QListData &)
  // append()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QListData{}) // "const QListData &"
  vtys[2] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListData6appendEi
    // invoke: void ** append(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QListData6appendEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.VoidpTy() // "void **"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN9QListData6appendERKS_
    // invoke: void ** append(const struct QListData &)
    var arg0 = args[0].(*QListData).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QListData6appendERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.VoidpTy() // "void **"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZN9QListData6appendEv
    // invoke: void ** append()
    var ret0 = C.C_ZN9QListData6appendEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.VoidpTy() // "void **"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListData", "append", args)
  }

  return
}

// size()
func (this *QListData) Size(args ...interface{}) (ret interface{}) {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListData4sizeEv
    // invoke: int size()
    var ret0 = C.C_ZNK9QListData4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListData", "size", args)
  }

  return
}

// <= body block end

