package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtCore/qreadwritelock.h
// dst-file: /src/core/qreadwritelock.go
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
  // proto:  void QWriteLocker::QWriteLocker(QReadWriteLock * readWriteLock);
extern void C_ZN12QWriteLockerC2EP14QReadWriteLock(void* qthis, void* arg0); // 1
  // proto:  QReadWriteLock * QWriteLocker::readWriteLock();
extern void C_ZNK12QWriteLocker13readWriteLockEv(void* qthis); // 2
  // proto:  void QWriteLocker::~QWriteLocker();
extern void C_ZN12QWriteLockerD2Ev(void* qthis); // 2
  // proto:  void QWriteLocker::unlock();
extern void C_ZN12QWriteLocker6unlockEv(void* qthis); // 2
  // proto:  void QWriteLocker::relock();
extern void C_ZN12QWriteLocker6relockEv(void* qthis); // 2
  // proto:  bool QReadWriteLock::tryLockForWrite();
extern void C_ZN14QReadWriteLock15tryLockForWriteEv(void* qthis); // 4
  // proto:  bool QReadWriteLock::tryLockForWrite(int timeout);
extern void C_ZN14QReadWriteLock15tryLockForWriteEi(void* qthis, int32_t arg0); // 4
  // proto:  void QReadWriteLock::~QReadWriteLock();
extern void C_ZN14QReadWriteLockD2Ev(void* qthis); // 4
  // proto:  bool QReadWriteLock::tryLockForRead(int timeout);
extern void C_ZN14QReadWriteLock14tryLockForReadEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QReadWriteLock::tryLockForRead();
extern void C_ZN14QReadWriteLock14tryLockForReadEv(void* qthis); // 4
  // proto:  void QReadWriteLock::unlock();
extern void C_ZN14QReadWriteLock6unlockEv(void* qthis); // 4
  // proto:  void QReadWriteLock::lockForWrite();
extern void C_ZN14QReadWriteLock12lockForWriteEv(void* qthis); // 4
  // proto:  void QReadWriteLock::lockForRead();
extern void C_ZN14QReadWriteLock11lockForReadEv(void* qthis); // 4
  // proto:  QReadWriteLock * QReadLocker::readWriteLock();
extern void C_ZNK11QReadLocker13readWriteLockEv(void* qthis); // 2
  // proto:  void QReadLocker::~QReadLocker();
extern void C_ZN11QReadLockerD2Ev(void* qthis); // 2
  // proto:  void QReadLocker::unlock();
extern void C_ZN11QReadLocker6unlockEv(void* qthis); // 2
  // proto:  void QReadLocker::relock();
extern void C_ZN11QReadLocker6relockEv(void* qthis); // 2
  // proto:  void QReadLocker::QReadLocker(QReadWriteLock * readWriteLock);
extern void C_ZN11QReadLockerC2EP14QReadWriteLock(void* qthis, void* arg0); // 1
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

// class sizeof(QWriteLocker)=4
type QWriteLocker struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QReadWriteLock)=8
type QReadWriteLock struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QReadLocker)=4
type QReadLocker struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// QWriteLocker(class QReadWriteLock *)
func NewQWriteLocker(args ...interface{}) QWriteLocker {
  // QWriteLocker(class QReadWriteLock *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QReadWriteLock{}) // "QReadWriteLock *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QWriteLockerC1EP14QReadWriteLock
    // invoke: void QWriteLocker(class QReadWriteLock *)
    var arg0 = args[0].(QReadWriteLock).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN12QWriteLockerC2EP14QReadWriteLock(qthis, arg0)
  default:
    qtrt.ErrorResolve("QWriteLocker", "QWriteLocker", args)
  }

  return QWriteLocker{}
}

// readWriteLock()
func (this *QWriteLocker) readWriteLock(args ...interface{}) () {
  // readWriteLock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QWriteLocker13readWriteLockEv
    // invoke: QReadWriteLock * readWriteLock()
    C.C_ZNK12QWriteLocker13readWriteLockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWriteLocker", "readWriteLock", args)
  }

}

// ~QWriteLocker()
func (this *QWriteLocker) FreeQWriteLocker(args ...interface{}) () {
  // ~QWriteLocker()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QWriteLockerD0Ev
    // invoke: void ~QWriteLocker()
    C.C_ZN12QWriteLockerD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWriteLocker", "~QWriteLocker", args)
  }

}

// unlock()
func (this *QWriteLocker) unlock(args ...interface{}) () {
  // unlock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QWriteLocker6unlockEv
    // invoke: void unlock()
    C.C_ZN12QWriteLocker6unlockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWriteLocker", "unlock", args)
  }

}

// relock()
func (this *QWriteLocker) relock(args ...interface{}) () {
  // relock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QWriteLocker6relockEv
    // invoke: void relock()
    C.C_ZN12QWriteLocker6relockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWriteLocker", "relock", args)
  }

}

// tryLockForWrite()
func (this *QReadWriteLock) tryLockForWrite(args ...interface{}) () {
  // tryLockForWrite()
  // tryLockForWrite(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QReadWriteLock15tryLockForWriteEv
    // invoke: bool tryLockForWrite()
    C.C_ZN14QReadWriteLock15tryLockForWriteEv(this.qclsinst)
  case 1:
    // invoke: _ZN14QReadWriteLock15tryLockForWriteEi
    // invoke: bool tryLockForWrite(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QReadWriteLock15tryLockForWriteEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QReadWriteLock", "tryLockForWrite", args)
  }

}

// ~QReadWriteLock()
func (this *QReadWriteLock) FreeQReadWriteLock(args ...interface{}) () {
  // ~QReadWriteLock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QReadWriteLockD0Ev
    // invoke: void ~QReadWriteLock()
    C.C_ZN14QReadWriteLockD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QReadWriteLock", "~QReadWriteLock", args)
  }

}

// tryLockForRead(int)
func (this *QReadWriteLock) tryLockForRead(args ...interface{}) () {
  // tryLockForRead(int)
  // tryLockForRead()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QReadWriteLock14tryLockForReadEi
    // invoke: bool tryLockForRead(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QReadWriteLock14tryLockForReadEi(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN14QReadWriteLock14tryLockForReadEv
    // invoke: bool tryLockForRead()
    C.C_ZN14QReadWriteLock14tryLockForReadEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QReadWriteLock", "tryLockForRead", args)
  }

}

// unlock()
func (this *QReadWriteLock) unlock(args ...interface{}) () {
  // unlock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QReadWriteLock6unlockEv
    // invoke: void unlock()
    C.C_ZN14QReadWriteLock6unlockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QReadWriteLock", "unlock", args)
  }

}

// lockForWrite()
func (this *QReadWriteLock) lockForWrite(args ...interface{}) () {
  // lockForWrite()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QReadWriteLock12lockForWriteEv
    // invoke: void lockForWrite()
    C.C_ZN14QReadWriteLock12lockForWriteEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QReadWriteLock", "lockForWrite", args)
  }

}

// lockForRead()
func (this *QReadWriteLock) lockForRead(args ...interface{}) () {
  // lockForRead()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QReadWriteLock11lockForReadEv
    // invoke: void lockForRead()
    C.C_ZN14QReadWriteLock11lockForReadEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QReadWriteLock", "lockForRead", args)
  }

}

// readWriteLock()
func (this *QReadLocker) readWriteLock(args ...interface{}) () {
  // readWriteLock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QReadLocker13readWriteLockEv
    // invoke: QReadWriteLock * readWriteLock()
    C.C_ZNK11QReadLocker13readWriteLockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QReadLocker", "readWriteLock", args)
  }

}

// ~QReadLocker()
func (this *QReadLocker) FreeQReadLocker(args ...interface{}) () {
  // ~QReadLocker()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QReadLockerD0Ev
    // invoke: void ~QReadLocker()
    C.C_ZN11QReadLockerD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QReadLocker", "~QReadLocker", args)
  }

}

// unlock()
func (this *QReadLocker) unlock(args ...interface{}) () {
  // unlock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QReadLocker6unlockEv
    // invoke: void unlock()
    C.C_ZN11QReadLocker6unlockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QReadLocker", "unlock", args)
  }

}

// relock()
func (this *QReadLocker) relock(args ...interface{}) () {
  // relock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QReadLocker6relockEv
    // invoke: void relock()
    C.C_ZN11QReadLocker6relockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QReadLocker", "relock", args)
  }

}

// QReadLocker(class QReadWriteLock *)
func NewQReadLocker(args ...interface{}) QReadLocker {
  // QReadLocker(class QReadWriteLock *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QReadWriteLock{}) // "QReadWriteLock *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QReadLockerC1EP14QReadWriteLock
    // invoke: void QReadLocker(class QReadWriteLock *)
    var arg0 = args[0].(QReadWriteLock).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QReadLockerC2EP14QReadWriteLock(qthis, arg0)
  default:
    qtrt.ErrorResolve("QReadLocker", "QReadLocker", args)
  }

  return QReadLocker{}
}

// <= body block end

