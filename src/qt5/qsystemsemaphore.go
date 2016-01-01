package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qsystemsemaphore.h
// dst-file: /src/core/qsystemsemaphore.go
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
// class sizeof(QSystemSemaphore)=1
type QSystemSemaphore struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QSystemSemaphore) key(args ...interface{}) () {
  // key()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QSystemSemaphore3keyEv
  default:
    qtrt.ErrorResolve("QSystemSemaphore", "key", args)
  }

}


func (this *QSystemSemaphore) release(args ...interface{}) () {
  // release(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QSystemSemaphore7releaseEi
  default:
    qtrt.ErrorResolve("QSystemSemaphore", "release", args)
  }

}


func (this *QSystemSemaphore) errorString(args ...interface{}) () {
  // errorString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QSystemSemaphore11errorStringEv
  default:
    qtrt.ErrorResolve("QSystemSemaphore", "errorString", args)
  }

}


func NewQSystemSemaphore(args ...interface{}) QSystemSemaphore {
  return QSystemSemaphore{}
}


func (this *QSystemSemaphore) acquire(args ...interface{}) () {
  // acquire()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QSystemSemaphore7acquireEv
  default:
    qtrt.ErrorResolve("QSystemSemaphore", "acquire", args)
  }

}


func (this *QSystemSemaphore) FreeQSystemSemaphore(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSystemSemaphore", "~QSystemSemaphore", args)
  }

}

// <= body block end

