package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtCore/qatomic.h
// dst-file: /src/core/qatomic.go
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
  // proto:  void QAtomicInt::QAtomicInt(int value);
extern void* C_ZN10QAtomicIntC2Ei(int32_t arg0); // 1
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

// class sizeof(QAtomicInt)=1
type QAtomicInt struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// QAtomicInt(int)
func NewQAtomicInt(args ...interface{}) *QAtomicInt {
  // QAtomicInt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QAtomicIntC1Ei
    // invoke: void QAtomicInt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QAtomicIntC2Ei(arg0)
    return &QAtomicInt{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QAtomicInt", "QAtomicInt", args)
  }

  return nil // QAtomicInt{qclsinst:qthis}
}

// <= body block end

