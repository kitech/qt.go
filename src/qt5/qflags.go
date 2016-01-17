package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtCore/qflags.h
// dst-file: /src/core/qflags.go
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
  // proto:  void QIncompatibleFlag::QIncompatibleFlag(int i);
extern void _ZN17QIncompatibleFlagC2Ei(void* qthis, int32_t arg0); // 1
  // proto:  void QFlag::QFlag(ushort ai);
extern void _ZN5QFlagC2Et(void* qthis, int16_t arg0); // 1
  // proto:  void QFlag::QFlag(int ai);
extern void _ZN5QFlagC2Ei(void* qthis, int32_t arg0); // 1
  // proto:  void QFlag::QFlag(short ai);
extern void _ZN5QFlagC2Es(void* qthis, int16_t arg0); // 1
  // proto:  void QFlag::QFlag(uint ai);
extern void _ZN5QFlagC2Ej(void* qthis, int32_t arg0); // 1
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

// class sizeof(QIncompatibleFlag)=4
type QIncompatibleFlag struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QFlag)=4
type QFlag struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// QIncompatibleFlag(int)
func NewQIncompatibleFlag(args ...interface{}) QIncompatibleFlag {
  // QIncompatibleFlag(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QIncompatibleFlagC1Ei
    // invoke: void QIncompatibleFlag(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN17QIncompatibleFlagC2Ei(qthis, arg0)
  default:
    qtrt.ErrorResolve("QIncompatibleFlag", "QIncompatibleFlag", args)
  }

  return QIncompatibleFlag{}
}

// QFlag(ushort)
func NewQFlag(args ...interface{}) QFlag {
  // QFlag(ushort)
  // QFlag(int)
  // QFlag(short)
  // QFlag(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int16Ty(false) // "ushort"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int16Ty(false) // "short"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFlagC1Et
    // invoke: void QFlag(ushort)
    var arg0 = C.int16_t(args[0].(int16))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN5QFlagC2Et(qthis, arg0)
  case 1:
    // invoke: _ZN5QFlagC1Ei
    // invoke: void QFlag(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN5QFlagC2Ei(qthis, arg0)
  case 2:
    // invoke: _ZN5QFlagC1Es
    // invoke: void QFlag(short)
    var arg0 = C.int16_t(args[0].(int16))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN5QFlagC2Es(qthis, arg0)
  case 3:
    // invoke: _ZN5QFlagC1Ej
    // invoke: void QFlag(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN5QFlagC2Ej(qthis, arg0)
  default:
    qtrt.ErrorResolve("QFlag", "QFlag", args)
  }

  return QFlag{}
}

// <= body block end

