package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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
extern void _ZN12QMutexLockerD2Ev(void* qthis); // 2
  // proto:  void QMutexLocker::unlock();
extern void _ZN12QMutexLocker6unlockEv(void* qthis); // 2
  // proto:  void QMutexLocker::QMutexLocker(QBasicMutex * m);
extern void _ZN12QMutexLockerC2EP11QBasicMutex(void* qthis, void* arg0); // 1
  // proto:  QMutex * QMutexLocker::mutex();
extern void _ZNK12QMutexLocker5mutexEv(void* qthis); // 2
  // proto:  void QMutexLocker::relock();
extern void _ZN12QMutexLocker6relockEv(void* qthis); // 2
  // proto:  void QBasicMutex::lock();
extern void _ZN11QBasicMutex4lockEv(void* qthis); // 2
  // proto:  void QBasicMutex::unlock();
extern void _ZN11QBasicMutex6unlockEv(void* qthis); // 2
  // proto:  bool QBasicMutex::tryLock();
extern void _ZN11QBasicMutex7tryLockEv(void* qthis); // 2
  // proto:  bool QBasicMutex::isRecursive();
extern void _ZN11QBasicMutex11isRecursiveEv(void* qthis); // 4
  // proto:  void QMutex::lock();
extern void _ZN6QMutex4lockEv(void* qthis); // 4
  // proto:  void QMutex::~QMutex();
extern void _ZN6QMutexD2Ev(void* qthis); // 4
  // proto:  void QMutex::unlock();
extern void _ZN6QMutex6unlockEv(void* qthis); // 4
  // proto:  bool QMutex::tryLock(int timeout);
extern void _ZN6QMutex7tryLockEi(void* qthis, int32_t arg0); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QBasicMutex)=1
type QBasicMutex struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QMutex)=1
type QMutex struct {
  /*qbase*/ QBasicMutex;
  qclsinst unsafe.Pointer /* *C.void */;
}

// ~QMutexLocker()
func (this *QMutexLocker) FreeQMutexLocker(args ...interface{}) () {
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
    C._ZN12QMutexLockerD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMutexLocker", "~QMutexLocker", args)
  }

}

// unlock()
func (this *QMutexLocker) unlock(args ...interface{}) () {
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
    C._ZN12QMutexLocker6unlockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMutexLocker", "unlock", args)
  }

}

// QMutexLocker(class QBasicMutex *)
func NewQMutexLocker(args ...interface{}) QMutexLocker {
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
    var arg0 = args[0].(QBasicMutex).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN12QMutexLockerC2EP11QBasicMutex(qthis, arg0)
  default:
    qtrt.ErrorResolve("QMutexLocker", "QMutexLocker", args)
  }

  return QMutexLocker{}
}

// mutex()
func (this *QMutexLocker) mutex(args ...interface{}) () {
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
    C._ZNK12QMutexLocker5mutexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMutexLocker", "mutex", args)
  }

}

// relock()
func (this *QMutexLocker) relock(args ...interface{}) () {
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
    C._ZN12QMutexLocker6relockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMutexLocker", "relock", args)
  }

}

// lock()
func (this *QBasicMutex) lock(args ...interface{}) () {
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
    C._ZN11QBasicMutex4lockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBasicMutex", "lock", args)
  }

}

// unlock()
func (this *QBasicMutex) unlock(args ...interface{}) () {
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
    C._ZN11QBasicMutex6unlockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBasicMutex", "unlock", args)
  }

}

// tryLock()
func (this *QBasicMutex) tryLock(args ...interface{}) () {
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
    C._ZN11QBasicMutex7tryLockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBasicMutex", "tryLock", args)
  }

}

// isRecursive()
func (this *QBasicMutex) isRecursive(args ...interface{}) () {
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
    C._ZN11QBasicMutex11isRecursiveEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBasicMutex", "isRecursive", args)
  }

}

// lock()
func (this *QMutex) lock(args ...interface{}) () {
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
    C._ZN6QMutex4lockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMutex", "lock", args)
  }

}

// ~QMutex()
func (this *QMutex) FreeQMutex(args ...interface{}) () {
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
    C._ZN6QMutexD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMutex", "~QMutex", args)
  }

}

// unlock()
func (this *QMutex) unlock(args ...interface{}) () {
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
    C._ZN6QMutex6unlockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMutex", "unlock", args)
  }

}

// tryLock(int)
func (this *QMutex) tryLock(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN6QMutex7tryLockEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMutex", "tryLock", args)
  }

}

// <= body block end

