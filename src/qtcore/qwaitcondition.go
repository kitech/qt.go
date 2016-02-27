package qtcore
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtCore/qwaitcondition.h
// dst-file: /src/core/qwaitcondition.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "runtime"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QWaitCondition::QWaitCondition();
extern void* C_ZN14QWaitConditionC2Ev(); // 3
  // proto:  void QWaitCondition::wakeAll();
extern void C_ZN14QWaitCondition7wakeAllEv(void* qthis); // 4
  // proto:  void QWaitCondition::wakeOne();
extern void C_ZN14QWaitCondition7wakeOneEv(void* qthis); // 4
  // proto:  void QWaitCondition::~QWaitCondition();
extern void C_ZN14QWaitConditionD2Ev(void* qthis); // 4
  // proto:  bool QWaitCondition::wait(QReadWriteLock * lockedReadWriteLock, unsigned long time);
extern bool C_ZN14QWaitCondition4waitEP14QReadWriteLockm(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  bool QWaitCondition::wait(QMutex * lockedMutex, unsigned long time);
extern bool C_ZN14QWaitCondition4waitEP6QMutexm(void* qthis, void* arg0, int32_t arg1); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QWaitCondition)=8
type QWaitCondition struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// QWaitCondition()
func GcfreeQWaitCondition(this *QWaitCondition) {
  qtrt.UniverseFree(this)
}
func NewQWaitCondition(args ...interface{}) *QWaitCondition {
  // QWaitCondition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QWaitConditionC1Ev
    // invoke: void QWaitCondition()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QWaitConditionC2Ev()
    this := &QWaitCondition{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQWaitCondition)
    return this
  default:
    qtrt.ErrorResolve("QWaitCondition", "QWaitCondition", args)
  }

  return nil // QWaitCondition{Qclsinst:qthis}
}

// wakeAll()
func (this *QWaitCondition) WakeAll(args ...interface{}) () {
  // wakeAll()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QWaitCondition7wakeAllEv
    // invoke: void wakeAll()
    C.C_ZN14QWaitCondition7wakeAllEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWaitCondition", "wakeAll", args)
  }

  return
}

// wakeOne()
func (this *QWaitCondition) WakeOne(args ...interface{}) () {
  // wakeOne()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QWaitCondition7wakeOneEv
    // invoke: void wakeOne()
    C.C_ZN14QWaitCondition7wakeOneEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWaitCondition", "wakeOne", args)
  }

  return
}

// ~QWaitCondition()
func (this *QWaitCondition) Free(args ...interface{}) () {
  // ~QWaitCondition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QWaitConditionD0Ev
    // invoke: void ~QWaitCondition()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN14QWaitConditionD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QWaitCondition", "~QWaitCondition", args)
  }

  return
}

// wait(class QReadWriteLock *, unsigned long)
func (this *QWaitCondition) Wait(args ...interface{}) (ret interface{}) {
  // wait(class QReadWriteLock *, unsigned long)
  // wait(class QMutex *, unsigned long)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QReadWriteLock{}) // "QReadWriteLock *"
  vtys[0][1] = qtrt.Int32Ty(false) // "unsigned long"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QMutex{}) // "QMutex *"
  vtys[1][1] = qtrt.Int32Ty(false) // "unsigned long"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QWaitCondition4waitEP14QReadWriteLockm
    // invoke: bool wait(class QReadWriteLock *, unsigned long)
    var arg0 = args[0].(*QReadWriteLock).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN14QWaitCondition4waitEP14QReadWriteLockm(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN14QWaitCondition4waitEP6QMutexm
    // invoke: bool wait(class QMutex *, unsigned long)
    var arg0 = args[0].(*QMutex).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN14QWaitCondition4waitEP6QMutexm(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWaitCondition", "wait", args)
  }

  return
}

// <= body block end

