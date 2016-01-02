package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
  // proto:  void QMutexLocker::QMutexLocker(QBasicMutex * m);
extern void* dector_ZN12QMutexLockerC1EP11QBasicMutex(void* arg0);
extern void demth_ZN12QMutexLockerC1EP11QBasicMutex(void* qthis, void* arg0);
  // proto:  QMutex * QMutexLocker::mutex();
extern void demth_ZNK12QMutexLocker5mutexEv(void* qthis);
  // proto:  void QMutexLocker::QMutexLocker(const QMutexLocker & );
extern void* dector_ZN12QMutexLockerC1ERKS_(void* arg0);
extern void _ZN12QMutexLockerC1ERKS_(void* qthis, void* arg0);
  // proto:  void QMutexLocker::relock();
extern void demth_ZN12QMutexLocker6relockEv(void* qthis);
  // proto:  void QMutexLocker::unlock();
extern void demth_ZN12QMutexLocker6unlockEv(void* qthis);
  // proto:  void QMutexLocker::~QMutexLocker();
extern void demth_ZN12QMutexLockerD0Ev(void* qthis);
  // proto:  void QBasicMutex::lock();
extern void demth_ZN11QBasicMutex4lockEv(void* qthis);
  // proto:  bool QBasicMutex::tryLock();
extern void _ZN11QBasicMutex7tryLockEv(void* qthis);
  // proto:  bool QBasicMutex::isRecursive();
extern void _ZN11QBasicMutex11isRecursiveEv(void* qthis);
  // proto:  void QBasicMutex::unlock();
extern void demth_ZN11QBasicMutex6unlockEv(void* qthis);
  // proto:  void QMutex::~QMutex();
extern void _ZN6QMutexD0Ev(void* qthis);
  // proto:  bool QMutex::tryLock(int timeout);
extern void _ZN6QMutex7tryLockEi(void* qthis, int arg0);
  // proto:  void QMutex::QMutex(const QMutex & );
extern void* dector_ZN6QMutexC1ERKS_(void* arg0);
extern void _ZN6QMutexC1ERKS_(void* qthis, void* arg0);
  // proto:  void QMutex::lock();
extern void _ZN6QMutex4lockEv(void* qthis);
  // proto:  void QMutex::unlock();
extern void _ZN6QMutex6unlockEv(void* qthis);
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
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QBasicMutex)=1
type QBasicMutex struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QMutex)=1
type QMutex struct {
  /*qbase*/ QBasicMutex;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QMutexLocker::QMutexLocker(QBasicMutex * m);
func NewQMutexLocker(args ...interface{}) QMutexLocker {
  return QMutexLocker{}
}

  // proto:  QMutex * QMutexLocker::mutex();
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
  default:
    qtrt.ErrorResolve("QMutexLocker", "mutex", args)
  }

}

  // proto:  void QMutexLocker::relock();
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
  default:
    qtrt.ErrorResolve("QMutexLocker", "relock", args)
  }

}

  // proto:  void QMutexLocker::unlock();
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
  default:
    qtrt.ErrorResolve("QMutexLocker", "unlock", args)
  }

}

  // proto:  void QMutexLocker::~QMutexLocker();
func (this *QMutexLocker) FreeQMutexLocker(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMutexLocker", "~QMutexLocker", args)
  }

}

  // proto:  void QBasicMutex::lock();
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
  default:
    qtrt.ErrorResolve("QBasicMutex", "lock", args)
  }

}

  // proto:  bool QBasicMutex::tryLock();
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
  default:
    qtrt.ErrorResolve("QBasicMutex", "tryLock", args)
  }

}

  // proto:  bool QBasicMutex::isRecursive();
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
  default:
    qtrt.ErrorResolve("QBasicMutex", "isRecursive", args)
  }

}

  // proto:  void QBasicMutex::unlock();
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
  default:
    qtrt.ErrorResolve("QBasicMutex", "unlock", args)
  }

}

  // proto:  void QMutex::~QMutex();
func (this *QMutex) FreeQMutex(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMutex", "~QMutex", args)
  }

}

  // proto:  bool QMutex::tryLock(int timeout);
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
  default:
    qtrt.ErrorResolve("QMutex", "tryLock", args)
  }

}

  // proto:  void QMutex::QMutex(const QMutex & );
func NewQMutex(args ...interface{}) QMutex {
  return QMutex{}
}

  // proto:  void QMutex::lock();
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
  default:
    qtrt.ErrorResolve("QMutex", "lock", args)
  }

}

  // proto:  void QMutex::unlock();
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
  default:
    qtrt.ErrorResolve("QMutex", "unlock", args)
  }

}

// <= body block end

