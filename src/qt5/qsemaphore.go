package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qsemaphore.h
// dst-file: /src/core/qsemaphore.go
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
// class sizeof(QSemaphore)=8
type QSemaphore struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QSemaphore) acquire(args ...interface{}) () {
  // acquire(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN10QSemaphore7acquireEi
  default:
    qtrt.ErrorResolve("QSemaphore", "acquire", args)
 }

}


func (this *QSemaphore) release(args ...interface{}) () {
  // release(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN10QSemaphore7releaseEi
  default:
    qtrt.ErrorResolve("QSemaphore", "release", args)
 }

}


func (this *QSemaphore) available(args ...interface{}) () {
  // available()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QSemaphore9availableEv
  default:
    qtrt.ErrorResolve("QSemaphore", "available", args)
 }

}


func (this *QSemaphore) tryAcquire(args ...interface{}) () {
  // tryAcquire(int, int)
  // tryAcquire(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN10QSemaphore10tryAcquireEii
  case 1:
    // invoke: _ZN10QSemaphore10tryAcquireEi
  default:
    qtrt.ErrorResolve("QSemaphore", "tryAcquire", args)
 }

}


func NewQSemaphore(args ...interface{})() {
}


func (this *QSemaphore) FreeQSemaphore(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QSemaphore", "~QSemaphore", args)
 }

}

// <= body block end

