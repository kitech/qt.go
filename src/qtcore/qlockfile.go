package qtcore
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
  // proto:  bool QLockFile::getLockInfo(qint64 * pid, QString * hostname, QString * appname);
extern bool C_ZNK9QLockFile11getLockInfoEPxP7QStringS2_(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QLockFile::setStaleLockTime(int );
extern void C_ZN9QLockFile16setStaleLockTimeEi(void* qthis, int32_t arg0); // 4
  // proto:  void QLockFile::QLockFile(const QString & fileName);
extern void* C_ZN9QLockFileC2ERK7QString(void* arg0); // 3
  // proto:  bool QLockFile::lock();
extern bool C_ZN9QLockFile4lockEv(void* qthis); // 4
  // proto:  void QLockFile::~QLockFile();
extern void C_ZN9QLockFileD2Ev(void* qthis); // 4
  // proto:  bool QLockFile::removeStaleLockFile();
extern bool C_ZN9QLockFile19removeStaleLockFileEv(void* qthis); // 4
  // proto:  bool QLockFile::isLocked();
extern bool C_ZNK9QLockFile8isLockedEv(void* qthis); // 4
  // proto:  int QLockFile::staleLockTime();
extern int32_t C_ZNK9QLockFile13staleLockTimeEv(void* qthis); // 4
  // proto:  bool QLockFile::tryLock(int timeout);
extern bool C_ZN9QLockFile7tryLockEi(void* qthis, int32_t arg0); // 4
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
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QLockFile)=1
type QLockFile struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// getLockInfo(qint64 *, class QString *, class QString *)
func (this *QLockFile) GetLockInfo(args ...interface{}) (ret interface{}) {
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
    var arg0 = (unsafe.Pointer)(args[0].(*int64))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QString).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK9QLockFile11getLockInfoEPxP7QStringS2_(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLockFile", "getLockInfo", args)
  }

  return
}

// setStaleLockTime(int)
func (this *QLockFile) SetStaleLockTime(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QLockFile16setStaleLockTimeEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLockFile", "setStaleLockTime", args)
  }

  return
}

// QLockFile(const class QString &)
func GcfreeQLockFile(this *QLockFile) {
  qtrt.UniverseFree(this)
}
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QLockFileC2ERK7QString(arg0)
    this := &QLockFile{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQLockFile)
    return this
  default:
    qtrt.ErrorResolve("QLockFile", "QLockFile", args)
  }

  return nil // QLockFile{Qclsinst:qthis}
}

// lock()
func (this *QLockFile) Lock(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN9QLockFile4lockEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLockFile", "lock", args)
  }

  return
}

// ~QLockFile()
func (this *QLockFile) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN9QLockFileD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QLockFile", "~QLockFile", args)
  }

  return
}

// removeStaleLockFile()
func (this *QLockFile) RemoveStaleLockFile(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN9QLockFile19removeStaleLockFileEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLockFile", "removeStaleLockFile", args)
  }

  return
}

// isLocked()
func (this *QLockFile) IsLocked(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QLockFile8isLockedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLockFile", "isLocked", args)
  }

  return
}

// staleLockTime()
func (this *QLockFile) StaleLockTime(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QLockFile13staleLockTimeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLockFile", "staleLockTime", args)
  }

  return
}

// tryLock(int)
func (this *QLockFile) TryLock(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QLockFile7tryLockEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLockFile", "tryLock", args)
  }

  return
}

// unlock()
func (this *QLockFile) Unlock(args ...interface{}) () {
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
    C.C_ZN9QLockFile6unlockEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLockFile", "unlock", args)
  }

  return
}

// error()
func (this *QLockFile) Error(args ...interface{}) () {
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
    C.C_ZNK9QLockFile5errorEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLockFile", "error", args)
  }

  return
}

// <= body block end

