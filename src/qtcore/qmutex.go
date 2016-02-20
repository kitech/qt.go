package qtcore
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
// src-file: /QtCore/qmutex.h
// dst-file: /src/core/qmutex.go
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
  // proto:  void QMutexLocker::~QMutexLocker();
extern void C_ZN12QMutexLockerD2Ev(void* qthis); // 2
  // proto:  void QMutexLocker::unlock();
extern void C_ZN12QMutexLocker6unlockEv(void* qthis); // 2
  // proto:  void QMutexLocker::QMutexLocker(QBasicMutex * m);
extern void* C_ZN12QMutexLockerC2EP11QBasicMutex(void* arg0); // 1
  // proto:  QMutex * QMutexLocker::mutex();
extern void* C_ZNK12QMutexLocker5mutexEv(void* qthis); // 2
  // proto:  void QMutexLocker::relock();
extern void C_ZN12QMutexLocker6relockEv(void* qthis); // 2
  // proto:  void QBasicMutex::lock();
extern void C_ZN11QBasicMutex4lockEv(void* qthis); // 2
  // proto:  void QBasicMutex::unlock();
extern void C_ZN11QBasicMutex6unlockEv(void* qthis); // 2
  // proto:  bool QBasicMutex::tryLock();
extern bool C_ZN11QBasicMutex7tryLockEv(void* qthis); // 2
  // proto:  bool QBasicMutex::isRecursive();
extern bool C_ZN11QBasicMutex11isRecursiveEv(void* qthis); // 4
  // proto:  void QMutex::lock();
extern void C_ZN6QMutex4lockEv(void* qthis); // 4
  // proto:  void QMutex::~QMutex();
extern void C_ZN6QMutexD2Ev(void* qthis); // 4
  // proto:  void QMutex::unlock();
extern void C_ZN6QMutex6unlockEv(void* qthis); // 4
  // proto:  bool QMutex::tryLock(int timeout);
extern bool C_ZN6QMutex7tryLockEi(void* qthis, int32_t arg0); // 4
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

// class sizeof(QMutexLocker)=4
type QMutexLocker struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QBasicMutex)=1
type QBasicMutex struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QMutex)=1
type QMutex struct {
  /*qbase*/ QBasicMutex;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// ~QMutexLocker()
func (this *QMutexLocker) Freeqmutexlocker(args ...interface{}) () {
  // ~QMutexLocker()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QMutexLockerD0Ev
    // invoke: void ~QMutexLocker()
    C.C_ZN12QMutexLockerD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMutexLocker", "~QMutexLocker", args)
  }

  return
}

// unlock()
func (this *QMutexLocker) Unlock(args ...interface{}) () {
  // unlock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QMutexLocker6unlockEv
    // invoke: void unlock()
    C.C_ZN12QMutexLocker6unlockEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMutexLocker", "unlock", args)
  }

  return
}

// QMutexLocker(class QBasicMutex *)
func NewQMutexLocker(args ...interface{}) *QMutexLocker {
  // QMutexLocker(class QBasicMutex *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBasicMutex{}) // "QBasicMutex *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QMutexLockerC1EP11QBasicMutex
    // invoke: void QMutexLocker(class QBasicMutex *)
    var arg0 = args[0].(*QBasicMutex).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QMutexLockerC2EP11QBasicMutex(arg0)
    return &QMutexLocker{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QMutexLocker", "QMutexLocker", args)
  }

  return nil // QMutexLocker{Qclsinst:qthis}
}

// mutex()
func (this *QMutexLocker) Mutex(args ...interface{}) (ret interface{}) {
  // mutex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QMutexLocker5mutexEv
    // invoke: QMutex * mutex()
    var ret0 = C.C_ZNK12QMutexLocker5mutexEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMutex{}) // "QMutex *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMutexLocker", "mutex", args)
  }

  return
}

// relock()
func (this *QMutexLocker) Relock(args ...interface{}) () {
  // relock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QMutexLocker6relockEv
    // invoke: void relock()
    C.C_ZN12QMutexLocker6relockEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMutexLocker", "relock", args)
  }

  return
}

// lock()
func (this *QBasicMutex) Lock(args ...interface{}) () {
  // lock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QBasicMutex4lockEv
    // invoke: void lock()
    C.C_ZN11QBasicMutex4lockEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QBasicMutex", "lock", args)
  }

  return
}

// unlock()
func (this *QBasicMutex) Unlock(args ...interface{}) () {
  // unlock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QBasicMutex6unlockEv
    // invoke: void unlock()
    C.C_ZN11QBasicMutex6unlockEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QBasicMutex", "unlock", args)
  }

  return
}

// tryLock()
func (this *QBasicMutex) Trylock(args ...interface{}) (ret interface{}) {
  // tryLock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QBasicMutex7tryLockEv
    // invoke: bool tryLock()
    var ret0 = C.C_ZN11QBasicMutex7tryLockEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBasicMutex", "tryLock", args)
  }

  return
}

// isRecursive()
func (this *QBasicMutex) Isrecursive(args ...interface{}) (ret interface{}) {
  // isRecursive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QBasicMutex11isRecursiveEv
    // invoke: bool isRecursive()
    var ret0 = C.C_ZN11QBasicMutex11isRecursiveEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBasicMutex", "isRecursive", args)
  }

  return
}

// lock()
func (this *QMutex) Lock(args ...interface{}) () {
  // lock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QMutex4lockEv
    // invoke: void lock()
    C.C_ZN6QMutex4lockEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMutex", "lock", args)
  }

  return
}

// ~QMutex()
func (this *QMutex) Freeqmutex(args ...interface{}) () {
  // ~QMutex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QMutexD0Ev
    // invoke: void ~QMutex()
    C.C_ZN6QMutexD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMutex", "~QMutex", args)
  }

  return
}

// unlock()
func (this *QMutex) Unlock(args ...interface{}) () {
  // unlock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QMutex6unlockEv
    // invoke: void unlock()
    C.C_ZN6QMutex6unlockEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMutex", "unlock", args)
  }

  return
}

// tryLock(int)
func (this *QMutex) Trylock(args ...interface{}) (ret interface{}) {
  // tryLock(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QMutex7tryLockEi
    // invoke: bool tryLock(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN6QMutex7tryLockEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMutex", "tryLock", args)
  }

  return
}

// <= body block end

