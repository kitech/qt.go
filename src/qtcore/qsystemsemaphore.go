package qtcore
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
// src-file: /QtCore/qsystemsemaphore.h
// dst-file: /src/core/qsystemsemaphore.go
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
  // proto:  QString QSystemSemaphore::errorString();
extern void* C_ZNK16QSystemSemaphore11errorStringEv(void* qthis); // 4
  // proto:  bool QSystemSemaphore::acquire();
extern bool C_ZN16QSystemSemaphore7acquireEv(void* qthis); // 4
  // proto:  void QSystemSemaphore::~QSystemSemaphore();
extern void C_ZN16QSystemSemaphoreD2Ev(void* qthis); // 4
  // proto:  QString QSystemSemaphore::key();
extern void* C_ZNK16QSystemSemaphore3keyEv(void* qthis); // 4
  // proto:  QSystemSemaphore::SystemSemaphoreError QSystemSemaphore::error();
extern void C_ZNK16QSystemSemaphore5errorEv(void* qthis); // 4
  // proto:  bool QSystemSemaphore::release(int n);
extern bool C_ZN16QSystemSemaphore7releaseEi(void* qthis, int32_t arg0); // 4
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

// class sizeof(QSystemSemaphore)=1
type QSystemSemaphore struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// errorString()
func (this *QSystemSemaphore) Errorstring(args ...interface{}) (ret interface{}) {
  // errorString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QSystemSemaphore11errorStringEv
    // invoke: QString errorString()
    var ret0 = C.C_ZNK16QSystemSemaphore11errorStringEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSystemSemaphore", "errorString", args)
  }

  return
}

// acquire()
func (this *QSystemSemaphore) Acquire(args ...interface{}) (ret interface{}) {
  // acquire()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QSystemSemaphore7acquireEv
    // invoke: bool acquire()
    var ret0 = C.C_ZN16QSystemSemaphore7acquireEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSystemSemaphore", "acquire", args)
  }

  return
}

// ~QSystemSemaphore()
func (this *QSystemSemaphore) Freeqsystemsemaphore(args ...interface{}) () {
  // ~QSystemSemaphore()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QSystemSemaphoreD0Ev
    // invoke: void ~QSystemSemaphore()
    C.C_ZN16QSystemSemaphoreD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSystemSemaphore", "~QSystemSemaphore", args)
  }

  return
}

// key()
func (this *QSystemSemaphore) Key(args ...interface{}) (ret interface{}) {
  // key()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QSystemSemaphore3keyEv
    // invoke: QString key()
    var ret0 = C.C_ZNK16QSystemSemaphore3keyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSystemSemaphore", "key", args)
  }

  return
}

// error()
func (this *QSystemSemaphore) Error(args ...interface{}) () {
  // error()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QSystemSemaphore5errorEv
    // invoke: QSystemSemaphore::SystemSemaphoreError error()
    C.C_ZNK16QSystemSemaphore5errorEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSystemSemaphore", "error", args)
  }

  return
}

// release(int)
func (this *QSystemSemaphore) Release(args ...interface{}) (ret interface{}) {
  // release(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QSystemSemaphore7releaseEi
    // invoke: bool release(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN16QSystemSemaphore7releaseEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSystemSemaphore", "release", args)
  }

  return
}

// <= body block end

