package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto: static QArrayData * QArrayData::sharedNull();
extern void _ZN10QArrayData10sharedNullEv();
  // proto:  void * QArrayData::data();
extern void _ZN10QArrayData4dataEv(void* qthis);
  // proto: static void QArrayData::deallocate(QArrayData * data, int objectSize, int alignment);
extern void _ZN10QArrayData10deallocateEPS_ii(void* arg0, int arg1, int arg2);
  // proto:  bool QArrayData::isMutable();
extern void _ZNK10QArrayData9isMutableEv(void* qthis);
  // proto:  int QArrayData::detachCapacity(int newSize);
extern void _ZNK10QArrayData14detachCapacityEi(void* qthis, int arg0);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto: static QArrayData * QArrayData::sharedNull();
func (this *QArrayData) sharedNull_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QArrayData", "sharedNull", args)
  }

}

  // proto:  void * QArrayData::data();
func (this *QArrayData) data(args ...interface{}) () {
  // data()
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QArrayData4dataEv
  case 1:
    // invoke: _ZN10QArrayData4dataEv
  default:
    qtrt.ErrorResolve("QArrayData", "data", args)
  }

}

  // proto: static void QArrayData::deallocate(QArrayData * data, int objectSize, int alignment);
func (this *QArrayData) deallocate_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QArrayData", "deallocate", args)
  }

}

  // proto:  bool QArrayData::isMutable();
func (this *QArrayData) isMutable(args ...interface{}) () {
  // isMutable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QArrayData9isMutableEv
  default:
    qtrt.ErrorResolve("QArrayData", "isMutable", args)
  }

}

  // proto:  int QArrayData::detachCapacity(int newSize);
func (this *QArrayData) detachCapacity(args ...interface{}) () {
  // detachCapacity(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QArrayData14detachCapacityEi
  default:
    qtrt.ErrorResolve("QArrayData", "detachCapacity", args)
  }

}

// <= body block end

