package qtcore
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
extern void* C_ZN20QContiguousCacheData12allocateDataEii(int32_t arg0, int32_t arg1); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// allocateData(int, int)
func (this *QContiguousCacheData) Allocatedata_S(args ...interface{}) (ret interface{}) {
  // allocateData(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QContiguousCacheData12allocateDataEii
    // invoke: QContiguousCacheData * allocateData(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN20QContiguousCacheData12allocateDataEii(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QContiguousCacheData{}) // "QContiguousCacheData *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QContiguousCacheData", "allocateData", args)
  }

  return
}

// freeData(struct QContiguousCacheData *)
func (this *QContiguousCacheData) Freedata_S(args ...interface{}) () {
  // freeData(struct QContiguousCacheData *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QContiguousCacheData{}) // "QContiguousCacheData *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QContiguousCacheData8freeDataEPS_
    // invoke: void freeData(struct QContiguousCacheData *)
    var arg0 = args[0].(*QContiguousCacheData).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN20QContiguousCacheData8freeDataEPS_(arg0)
  default:
    qtrt.ErrorResolve("QContiguousCacheData", "freeData", args)
  }

  return
}

// <= body block end

