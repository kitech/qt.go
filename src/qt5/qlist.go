package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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
extern void _ZN9QListData6insertEi(void* qthis, int32_t arg0); // 4
  // proto:  void ** QListData::begin();
extern void _ZNK9QListData5beginEv(void* qthis); // 2
  // proto:  void ** QListData::end();
extern void _ZNK9QListData3endEv(void* qthis); // 2
  // proto:  void QListData::realloc(int alloc);
extern void _ZN9QListData7reallocEi(void* qthis, int32_t arg0); // 4
  // proto:  void QListData::move(int from, int to);
extern void _ZN9QListData4moveEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QListData::dispose();
extern void _ZN9QListData7disposeEv(void* qthis); // 2
  // proto:  void QListData::remove(int i);
extern void _ZN9QListData6removeEi(void* qthis, int32_t arg0); // 4
  // proto:  void QListData::remove(int i, int n);
extern void _ZN9QListData6removeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void ** QListData::prepend();
extern void _ZN9QListData7prependEv(void* qthis); // 4
  // proto:  void ** QListData::erase(void ** xi);
extern void _ZN9QListData5eraseEPPv(void* qthis, void* arg0); // 4
  // proto:  bool QListData::isEmpty();
extern void _ZNK9QListData7isEmptyEv(void* qthis); // 2
  // proto:  void ** QListData::at(int i);
extern void _ZNK9QListData2atEi(void* qthis, int32_t arg0); // 2
  // proto:  QListData::Data * QListData::detach_grow(int * i, int n);
extern void _ZN9QListData11detach_growEPii(void* qthis, int32_t* arg0, int32_t arg1); // 4
  // proto:  QListData::Data * QListData::detach(int alloc);
extern void _ZN9QListData6detachEi(void* qthis, int32_t arg0); // 4
  // proto:  void ** QListData::append(int n);
extern void _ZN9QListData6appendEi(void* qthis, int32_t arg0); // 4
  // proto:  void ** QListData::append(const QListData & l);
extern void _ZN9QListData6appendERKS_(void* qthis, void* arg0); // 4
  // proto:  void ** QListData::append();
extern void _ZN9QListData6appendEv(void* qthis); // 4
  // proto:  int QListData::size();
extern void _ZNK9QListData4sizeEv(void* qthis); // 2
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

// class sizeof(QListData)=8
type QListData struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// insert(int)
func (this *QListData) insert(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QListData6insertEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListData", "insert", args)
  }

}

// begin()
func (this *QListData) begin(args ...interface{}) () {
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
    C._ZNK9QListData5beginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListData", "begin", args)
  }

}

// end()
func (this *QListData) end(args ...interface{}) () {
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
    C._ZNK9QListData3endEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListData", "end", args)
  }

}

// realloc(int)
func (this *QListData) realloc(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QListData7reallocEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListData", "realloc", args)
  }

}

// move(int, int)
func (this *QListData) move_(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN9QListData4moveEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QListData", "move", args)
  }

}

// dispose()
func (this *QListData) dispose(args ...interface{}) () {
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
    C._ZN9QListData7disposeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListData", "dispose", args)
  }

}

// remove(int)
func (this *QListData) remove(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QListData6removeEi(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN9QListData6removeEii
    // invoke: void remove(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN9QListData6removeEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QListData", "remove", args)
  }

}

// prepend()
func (this *QListData) prepend(args ...interface{}) () {
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
    C._ZN9QListData7prependEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListData", "prepend", args)
  }

}

// erase(void **)
func (this *QListData) erase(args ...interface{}) () {
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
    C._ZN9QListData5eraseEPPv(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListData", "erase", args)
  }

}

// isEmpty()
func (this *QListData) isEmpty(args ...interface{}) () {
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
    C._ZNK9QListData7isEmptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListData", "isEmpty", args)
  }

}

// at(int)
func (this *QListData) at(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK9QListData2atEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListData", "at", args)
  }

}

// detach_grow(int *, int)
func (this *QListData) detach_grow(args ...interface{}) () {
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
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN9QListData11detach_growEPii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QListData", "detach_grow", args)
  }

}

// detach(int)
func (this *QListData) detach(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QListData6detachEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListData", "detach", args)
  }

}

// append(int)
func (this *QListData) append(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QListData6appendEi(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN9QListData6appendERKS_
    // invoke: void ** append(const struct QListData &)
    var arg0 = args[0].(QListData).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QListData6appendERKS_(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN9QListData6appendEv
    // invoke: void ** append()
    C._ZN9QListData6appendEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListData", "append", args)
  }

}

// size()
func (this *QListData) size(args ...interface{}) () {
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
    C._ZNK9QListData4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListData", "size", args)
  }

}

// <= body block end

