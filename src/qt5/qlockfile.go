package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
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
  // proto:  bool QLockFile::removeStaleLockFile();
extern void _ZN9QLockFile19removeStaleLockFileEv(void* qthis);
  // proto:  int QLockFile::staleLockTime();
extern void _ZNK9QLockFile13staleLockTimeEv(void* qthis);
  // proto:  bool QLockFile::isLocked();
extern void _ZNK9QLockFile8isLockedEv(void* qthis);
  // proto:  void QLockFile::QLockFile(const QLockFile & );
extern void* dector_ZN9QLockFileC1ERKS_(void* arg0);
extern void _ZN9QLockFileC1ERKS_(void* qthis, void* arg0);
  // proto:  void QLockFile::~QLockFile();
extern void _ZN9QLockFileD0Ev(void* qthis);
  // proto:  void QLockFile::unlock();
extern void _ZN9QLockFile6unlockEv(void* qthis);
  // proto:  bool QLockFile::tryLock(int timeout);
extern void _ZN9QLockFile7tryLockEi(void* qthis, int arg0);
  // proto:  bool QLockFile::lock();
extern void _ZN9QLockFile4lockEv(void* qthis);
  // proto:  void QLockFile::setStaleLockTime(int );
extern void _ZN9QLockFile16setStaleLockTimeEi(void* qthis, int arg0);
  // proto:  bool QLockFile::getLockInfo(qint64 * pid, QString * hostname, QString * appname);
extern void _ZNK9QLockFile11getLockInfoEPxP7QStringS2_(void* qthis, long long* arg0, void* arg1, void* arg2);
  // proto:  void QLockFile::QLockFile(const QString & fileName);
extern void* dector_ZN9QLockFileC1ERK7QString(void* arg0);
extern void _ZN9QLockFileC1ERK7QString(void* qthis, void* arg0);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  bool QLockFile::removeStaleLockFile();
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
    C._ZN9QLockFile19removeStaleLockFileEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLockFile", "removeStaleLockFile", args)
  }

}

  // proto:  int QLockFile::staleLockTime();
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
    C._ZNK9QLockFile13staleLockTimeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLockFile", "staleLockTime", args)
  }

}

  // proto:  bool QLockFile::isLocked();
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
    C._ZNK9QLockFile8isLockedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLockFile", "isLocked", args)
  }

}

  // proto:  void QLockFile::QLockFile(const QLockFile & );
func NewQLockFile(args ...interface{}) QLockFile {
  return QLockFile{}
}

  // proto:  void QLockFile::~QLockFile();
func (this *QLockFile) FreeQLockFile(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QLockFile", "~QLockFile", args)
  }

}

  // proto:  void QLockFile::unlock();
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
    C._ZN9QLockFile6unlockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLockFile", "unlock", args)
  }

}

  // proto:  bool QLockFile::tryLock(int timeout);
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
    C._ZN9QLockFile7tryLockEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLockFile", "tryLock", args)
  }

}

  // proto:  bool QLockFile::lock();
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
    C._ZN9QLockFile4lockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLockFile", "lock", args)
  }

}

  // proto:  void QLockFile::setStaleLockTime(int );
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
    C._ZN9QLockFile16setStaleLockTimeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLockFile", "setStaleLockTime", args)
  }

}

  // proto:  bool QLockFile::getLockInfo(qint64 * pid, QString * hostname, QString * appname);
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
    C._ZNK9QLockFile11getLockInfoEPxP7QStringS2_(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QLockFile", "getLockInfo", args)
  }

}

// <= body block end

