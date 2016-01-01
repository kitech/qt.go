package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qmutex.h
// dst-file: /src/core/qmutex.go
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


func NewQMutexLocker(args ...interface{})() {
}


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


func NewQMutex(args ...interface{})() {
}


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

