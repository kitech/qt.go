package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtCore/qhash.h
// dst-file: /src/core/qhash.go
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
  // proto:  void QHashData::rehash(int hint);
extern void C_ZN9QHashData6rehashEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QHashData::willGrow();
extern bool C_ZN9QHashData8willGrowEv(void* qthis); // 2
  // proto:  void * QHashData::allocateNode(int nodeAlign);
extern void* C_ZN9QHashData12allocateNodeEi(void* qthis, int32_t arg0); // 4
  // proto:  void QHashData::freeNode(void * node);
extern void C_ZN9QHashData8freeNodeEPv(void* qthis, void* arg0); // 4
  // proto:  QHashData::Node * QHashData::firstNode();
extern void C_ZN9QHashData9firstNodeEv(void* qthis); // 2
  // proto:  void QHashData::hasShrunk();
extern void C_ZN9QHashData9hasShrunkEv(void* qthis); // 2
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

// class sizeof(QHashDummyValue)=1
type QHashDummyValue struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QHashData)=1
type QHashData struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// rehash(int)
func (this *QHashData) Rehash(args ...interface{}) () {
  // rehash(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QHashData6rehashEi
    // invoke: void rehash(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QHashData6rehashEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHashData", "rehash", args)
  }

  return
}

// willGrow()
func (this *QHashData) Willgrow(args ...interface{}) (ret interface{}) {
  // willGrow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QHashData8willGrowEv
    // invoke: bool willGrow()
    var ret0 = C.C_ZN9QHashData8willGrowEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHashData", "willGrow", args)
  }

  return
}

// allocateNode(int)
func (this *QHashData) Allocatenode(args ...interface{}) (ret interface{}) {
  // allocateNode(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QHashData12allocateNodeEi
    // invoke: void * allocateNode(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QHashData12allocateNodeEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.VoidpTy() // "void *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHashData", "allocateNode", args)
  }

  return
}

// freeNode(void *)
func (this *QHashData) Freenode(args ...interface{}) () {
  // freeNode(void *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.VoidpTy() // "void *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QHashData8freeNodeEPv
    // invoke: void freeNode(void *)
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
    C.C_ZN9QHashData8freeNodeEPv(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHashData", "freeNode", args)
  }

  return
}

// firstNode()
func (this *QHashData) Firstnode(args ...interface{}) () {
  // firstNode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QHashData9firstNodeEv
    // invoke: QHashData::Node * firstNode()
    C.C_ZN9QHashData9firstNodeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHashData", "firstNode", args)
  }

  return
}

// hasShrunk()
func (this *QHashData) Hasshrunk(args ...interface{}) () {
  // hasShrunk()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QHashData9hasShrunkEv
    // invoke: void hasShrunk()
    C.C_ZN9QHashData9hasShrunkEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHashData", "hasShrunk", args)
  }

  return
}

// <= body block end

