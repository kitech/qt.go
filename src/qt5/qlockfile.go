package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtCore/qlockfile.h
// dst-file: /src/core/qlockfile.go
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
  // proto:  bool QLockFile::getLockInfo(qint64 * pid, QString * hostname, QString * appname);
extern void C_ZNK9QLockFile11getLockInfoEPxP7QStringS2_(void* qthis, int64_t* arg0, void* arg1, void* arg2); // 4
  // proto:  void QLockFile::setStaleLockTime(int );
extern void C_ZN9QLockFile16setStaleLockTimeEi(void* qthis, int32_t arg0); // 4
  // proto:  void QLockFile::QLockFile(const QString & fileName);
extern void* C_ZN9QLockFileC2ERK7QString(void* arg0); // 3
  // proto:  bool QLockFile::lock();
extern void C_ZN9QLockFile4lockEv(void* qthis); // 4
  // proto:  void QLockFile::~QLockFile();
extern void C_ZN9QLockFileD2Ev(void* qthis); // 4
  // proto:  bool QLockFile::removeStaleLockFile();
extern void C_ZN9QLockFile19removeStaleLockFileEv(void* qthis); // 4
  // proto:  bool QLockFile::isLocked();
extern void C_ZNK9QLockFile8isLockedEv(void* qthis); // 4
  // proto:  int QLockFile::staleLockTime();
extern void C_ZNK9QLockFile13staleLockTimeEv(void* qthis); // 4
  // proto:  bool QLockFile::tryLock(int timeout);
extern void C_ZN9QLockFile7tryLockEi(void* qthis, int32_t arg0); // 4
  // proto:  void QLockFile::unlock();
extern void C_ZN9QLockFile6unlockEv(void* qthis); // 4
  // proto:  QLockFile::LockError QLockFile::error();
extern void C_ZNK9QLockFile5errorEv(void* qthis); // 4
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

// class sizeof(QLockFile)=1
type QLockFile struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// getLockInfo(qint64 *, class QString *, class QString *)
func (this *QLockFile) getLockInfo(args ...interface{}) () {
  // getLockInfo(qint64 *, class QString *, class QString *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(true) // "qint64 *"
  vtys[0][1] = reflect.TypeOf(QString{}) // "QString *"
  vtys[0][2] = reflect.TypeOf(QString{}) // "QString *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLockFile11getLockInfoEPxP7QStringS2_
    // invoke: bool getLockInfo(qint64 *, class QString *, class QString *)
    var arg0 = (*C.int64_t)(args[0].(*int64))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var ret = C.C_ZNK9QLockFile11getLockInfoEPxP7QStringS2_(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QLockFile", "getLockInfo", args)
  }

}

// setStaleLockTime(int)
func (this *QLockFile) setStaleLockTime(args ...interface{}) () {
  // setStaleLockTime(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLockFile16setStaleLockTimeEi
    // invoke: void setStaleLockTime(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QLockFile16setStaleLockTimeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLockFile", "setStaleLockTime", args)
  }

}

// QLockFile(const class QString &)
func NewQLockFile(args ...interface{}) *QLockFile {
  // QLockFile(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLockFileC1ERK7QString
    // invoke: void QLockFile(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QLockFileC2ERK7QString(arg0)
    return &QLockFile{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QLockFile", "QLockFile", args)
  }

  return nil // QLockFile{qclsinst:qthis}
}

// lock()
func (this *QLockFile) lock(args ...interface{}) () {
  // lock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLockFile4lockEv
    // invoke: bool lock()
    var ret = C.C_ZN9QLockFile4lockEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QLockFile", "lock", args)
  }

}

// ~QLockFile()
func (this *QLockFile) FreeQLockFile(args ...interface{}) () {
  // ~QLockFile()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLockFileD0Ev
    // invoke: void ~QLockFile()
    C.C_ZN9QLockFileD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLockFile", "~QLockFile", args)
  }

}

// removeStaleLockFile()
func (this *QLockFile) removeStaleLockFile(args ...interface{}) () {
  // removeStaleLockFile()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLockFile19removeStaleLockFileEv
    // invoke: bool removeStaleLockFile()
    var ret = C.C_ZN9QLockFile19removeStaleLockFileEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QLockFile", "removeStaleLockFile", args)
  }

}

// isLocked()
func (this *QLockFile) isLocked(args ...interface{}) () {
  // isLocked()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLockFile8isLockedEv
    // invoke: bool isLocked()
    var ret = C.C_ZNK9QLockFile8isLockedEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QLockFile", "isLocked", args)
  }

}

// staleLockTime()
func (this *QLockFile) staleLockTime(args ...interface{}) () {
  // staleLockTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLockFile13staleLockTimeEv
    // invoke: int staleLockTime()
    var ret = C.C_ZNK9QLockFile13staleLockTimeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QLockFile", "staleLockTime", args)
  }

}

// tryLock(int)
func (this *QLockFile) tryLock(args ...interface{}) () {
  // tryLock(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLockFile7tryLockEi
    // invoke: bool tryLock(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN9QLockFile7tryLockEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QLockFile", "tryLock", args)
  }

}

// unlock()
func (this *QLockFile) unlock(args ...interface{}) () {
  // unlock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLockFile6unlockEv
    // invoke: void unlock()
    C.C_ZN9QLockFile6unlockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLockFile", "unlock", args)
  }

}

// error()
func (this *QLockFile) error(args ...interface{}) () {
  // error()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLockFile5errorEv
    // invoke: QLockFile::LockError error()
    C.C_ZNK9QLockFile5errorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLockFile", "error", args)
  }

}

// <= body block end

