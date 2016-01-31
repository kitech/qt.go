package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtCore/qthreadstorage.h
// dst-file: /src/core/qthreadstorage.go
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
  // proto:  void QThreadStorageData::~QThreadStorageData();
extern void C_ZN18QThreadStorageDataD2Ev(void* qthis); // 4
  // proto:  void ** QThreadStorageData::set(void * p);
extern void C_ZN18QThreadStorageData3setEPv(void* qthis, void* arg0); // 4
  // proto: static void QThreadStorageData::finish(void ** );
extern void C_ZN18QThreadStorageData6finishEPPv(void* arg0); // 4
  // proto:  void ** QThreadStorageData::get();
extern void C_ZNK18QThreadStorageData3getEv(void* qthis); // 4
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

// class sizeof(QThreadStorageData)=4
type QThreadStorageData struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// ~QThreadStorageData()
func (this *QThreadStorageData) Freeqthreadstoragedata(args ...interface{}) () {
  // ~QThreadStorageData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QThreadStorageDataD0Ev
    // invoke: void ~QThreadStorageData()
    C.C_ZN18QThreadStorageDataD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThreadStorageData", "~QThreadStorageData", args)
  }

  return
}

// set(void *)
func (this *QThreadStorageData) Set(args ...interface{}) () {
  // set(void *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.VoidpTy() // "void *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QThreadStorageData3setEPv
    // invoke: void ** set(void *)
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
    C.C_ZN18QThreadStorageData3setEPv(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QThreadStorageData", "set", args)
  }

  return
}

// finish(void **)
func (this *QThreadStorageData) Finish_S(args ...interface{}) () {
  // finish(void **)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.VoidpTy() // "void **"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QThreadStorageData6finishEPPv
    // invoke: void finish(void **)
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
    C.C_ZN18QThreadStorageData6finishEPPv(arg0)
  default:
    qtrt.ErrorResolve("QThreadStorageData", "finish", args)
  }

  return
}

// get()
func (this *QThreadStorageData) Get(args ...interface{}) () {
  // get()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QThreadStorageData3getEv
    // invoke: void ** get()
    C.C_ZNK18QThreadStorageData3getEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThreadStorageData", "get", args)
  }

  return
}

// <= body block end

