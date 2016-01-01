package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qeventloop.h
// dst-file: /src/core/qeventloop.go
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
// class sizeof(QEventLoop)=1
type QEventLoop struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QEventLoopLocker)=8
type QEventLoopLocker struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QEventLoop) exit(args ...interface{}) () {
  // exit(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QEventLoop4exitEi
  default:
    qtrt.ErrorResolve("QEventLoop", "exit", args)
  }

}


func (this *QEventLoop) quit(args ...interface{}) () {
  // quit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QEventLoop4quitEv
  default:
    qtrt.ErrorResolve("QEventLoop", "quit", args)
  }

}


func NewQEventLoop(args ...interface{}) QEventLoop {
  return QEventLoop{}
}


func (this *QEventLoop) isRunning(args ...interface{}) () {
  // isRunning()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QEventLoop9isRunningEv
  default:
    qtrt.ErrorResolve("QEventLoop", "isRunning", args)
  }

}


func (this *QEventLoop) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QEventLoop10metaObjectEv
  default:
    qtrt.ErrorResolve("QEventLoop", "metaObject", args)
  }

}


func (this *QEventLoop) wakeUp(args ...interface{}) () {
  // wakeUp()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QEventLoop6wakeUpEv
  default:
    qtrt.ErrorResolve("QEventLoop", "wakeUp", args)
  }

}


func (this *QEventLoop) FreeQEventLoop(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QEventLoop", "~QEventLoop", args)
  }

}


func (this *QEventLoop) event(args ...interface{}) () {
  // event(class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QEventLoop5eventEP6QEvent
  default:
    qtrt.ErrorResolve("QEventLoop", "event", args)
  }

}


func NewQEventLoopLocker(args ...interface{}) QEventLoopLocker {
  return QEventLoopLocker{}
}


func (this *QEventLoopLocker) FreeQEventLoopLocker(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QEventLoopLocker", "~QEventLoopLocker", args)
  }

}

// <= body block end

