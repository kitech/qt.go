package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qreadwritelock.h
// dst-file: /src/core/qreadwritelock.go
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
// class sizeof(QWriteLocker)=4
type QWriteLocker struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QReadWriteLock)=8
type QReadWriteLock struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QReadLocker)=4
type QReadLocker struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


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
  default:
    qtrt.ErrorResolve("QWriteLocker", "readWriteLock", args)
  }

}


func NewQWriteLocker(args ...interface{}) QWriteLocker {
  return QWriteLocker{}
}


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
  default:
    qtrt.ErrorResolve("QWriteLocker", "unlock", args)
  }

}


func (this *QWriteLocker) FreeQWriteLocker(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWriteLocker", "~QWriteLocker", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWriteLocker", "relock", args)
  }

}


func (this *QReadWriteLock) FreeQReadWriteLock(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QReadWriteLock", "~QReadWriteLock", args)
  }

}


func NewQReadWriteLock(args ...interface{}) QReadWriteLock {
  return QReadWriteLock{}
}


func (this *QReadWriteLock) tryLockForRead(args ...interface{}) () {
  // tryLockForRead()
  // tryLockForRead(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QReadWriteLock14tryLockForReadEv
  case 1:
    // invoke: _ZN14QReadWriteLock14tryLockForReadEi
  default:
    qtrt.ErrorResolve("QReadWriteLock", "tryLockForRead", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QReadWriteLock", "lockForWrite", args)
  }

}


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
  case 1:
    // invoke: _ZN14QReadWriteLock15tryLockForWriteEi
  default:
    qtrt.ErrorResolve("QReadWriteLock", "tryLockForWrite", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QReadWriteLock", "unlock", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QReadWriteLock", "lockForRead", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QReadLocker", "readWriteLock", args)
  }

}


func (this *QReadLocker) FreeQReadLocker(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QReadLocker", "~QReadLocker", args)
  }

}


func NewQReadLocker(args ...interface{}) QReadLocker {
  return QReadLocker{}
}


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
  default:
    qtrt.ErrorResolve("QReadLocker", "relock", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QReadLocker", "unlock", args)
  }

}

// <= body block end

