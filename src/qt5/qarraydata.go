package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtCore/qarraydata.h
// dst-file: /src/core/qarraydata.go
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
  // proto:  AllocationOptions QArrayData::detachFlags();
extern void C_ZNK10QArrayData11detachFlagsEv(void* qthis); // 2
  // proto:  AllocationOptions QArrayData::cloneFlags();
extern void C_ZNK10QArrayData10cloneFlagsEv(void* qthis); // 2
  // proto:  void * QArrayData::data();
extern void* C_ZN10QArrayData4dataEv(void* qthis); // 2
  // proto: static QArrayData * QArrayData::sharedNull();
extern void* C_ZN10QArrayData10sharedNullEv(); // 2
  // proto:  bool QArrayData::isMutable();
extern bool C_ZNK10QArrayData9isMutableEv(void* qthis); // 2
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

// class sizeof(QArrayData)=1
type QArrayData struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// detachFlags()
func (this *QArrayData) Detachflags(args ...interface{}) () {
  // detachFlags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QArrayData11detachFlagsEv
    // invoke: AllocationOptions detachFlags()
    C.C_ZNK10QArrayData11detachFlagsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QArrayData", "detachFlags", args)
  }

  return
}

// cloneFlags()
func (this *QArrayData) Cloneflags(args ...interface{}) () {
  // cloneFlags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QArrayData10cloneFlagsEv
    // invoke: AllocationOptions cloneFlags()
    C.C_ZNK10QArrayData10cloneFlagsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QArrayData", "cloneFlags", args)
  }

  return
}

// data()
func (this *QArrayData) Data(args ...interface{}) (ret interface{}) {
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QArrayData4dataEv
    // invoke: void * data()
    var ret0 = C.C_ZN10QArrayData4dataEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.VoidpTy() // "void *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QArrayData", "data", args)
  }

  return
}

// sharedNull()
func (this *QArrayData) Sharednull_S(args ...interface{}) (ret interface{}) {
  // sharedNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QArrayData10sharedNullEv
    // invoke: QArrayData * sharedNull()
    var ret0 = C.C_ZN10QArrayData10sharedNullEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QArrayData{}) // "QArrayData *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QArrayData", "sharedNull", args)
  }

  return
}

// isMutable()
func (this *QArrayData) Ismutable(args ...interface{}) (ret interface{}) {
  // isMutable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QArrayData9isMutableEv
    // invoke: bool isMutable()
    var ret0 = C.C_ZNK10QArrayData9isMutableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QArrayData", "isMutable", args)
  }

  return
}

// <= body block end

