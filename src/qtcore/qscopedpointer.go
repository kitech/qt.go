package qtcore
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
// src-file: /QtCore/qscopedpointer.h
// dst-file: /src/core/qscopedpointer.go
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
  // proto: static void QScopedPointerPodDeleter::cleanup(void * pointer);
extern void C_ZN24QScopedPointerPodDeleter7cleanupEPv(void* arg0); // 2
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

// class sizeof(QScopedPointerPodDeleter)=1
type QScopedPointerPodDeleter struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// cleanup(void *)
func (this *QScopedPointerPodDeleter) Cleanup_S(args ...interface{}) () {
  // cleanup(void *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.VoidpTy() // "void *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QScopedPointerPodDeleter7cleanupEPv
    // invoke: void cleanup(void *)
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
    C.C_ZN24QScopedPointerPodDeleter7cleanupEPv(arg0)
  default:
    qtrt.ErrorResolve("QScopedPointerPodDeleter", "cleanup", args)
  }

  return
}

// <= body block end

