package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qlockfile.h
// dst-file: /src/core/qlockfile.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
// class sizeof(QLockFile)=1
type QLockFile struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


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
  default:
    qtrt.ErrorResolve("QLockFile", "removeStaleLockFile", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QLockFile", "staleLockTime", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QLockFile", "isLocked", args)
 }

}


func NewQLockFile(args ...interface{})() {
}


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
  default:
    qtrt.ErrorResolve("QLockFile", "unlock", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QLockFile", "tryLock", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QLockFile", "lock", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QLockFile", "setStaleLockTime", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QLockFile", "getLockInfo", args)
 }

}

// <= body block end

