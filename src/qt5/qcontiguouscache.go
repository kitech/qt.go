package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtCore/qcontiguouscache.h
// dst-file: /src/core/qcontiguouscache.go
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
  // proto: static QContiguousCacheData * QContiguousCacheData::allocateData(int size, int alignment);
extern void C_ZN20QContiguousCacheData12allocateDataEii(int32_t arg0, int32_t arg1); // 4
  // proto: static void QContiguousCacheData::freeData(QContiguousCacheData * data);
extern void C_ZN20QContiguousCacheData8freeDataEPS_(void* arg0); // 4
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

// class sizeof(QContiguousCacheData)=1
type QContiguousCacheData struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// allocateData(int, int)
func (this *QContiguousCacheData) allocateData_s(args ...interface{}) () {
  // allocateData(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QContiguousCacheData12allocateDataEii
    // invoke: QContiguousCacheData * allocateData(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN20QContiguousCacheData12allocateDataEii(arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QContiguousCacheData", "allocateData", args)
  }

}

// freeData(struct QContiguousCacheData *)
func (this *QContiguousCacheData) freeData_s(args ...interface{}) () {
  // freeData(struct QContiguousCacheData *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QContiguousCacheData{}) // "QContiguousCacheData *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QContiguousCacheData8freeDataEPS_
    // invoke: void freeData(struct QContiguousCacheData *)
    var arg0 = args[0].(QContiguousCacheData).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN20QContiguousCacheData8freeDataEPS_(arg0)
  default:
    qtrt.ErrorResolve("QContiguousCacheData", "freeData", args)
  }

}

// <= body block end

