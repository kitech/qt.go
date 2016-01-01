package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qrunnable.h
// dst-file: /src/core/qrunnable.go
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
// class sizeof(QRunnable)=16
type QRunnable struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QRunnable) FreeQRunnable(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QRunnable", "~QRunnable", args)
  }

}


func (this *QRunnable) setAutoDelete(args ...interface{}) () {
  // setAutoDelete(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QRunnable13setAutoDeleteEb
  default:
    qtrt.ErrorResolve("QRunnable", "setAutoDelete", args)
  }

}


func NewQRunnable(args ...interface{}) QRunnable {
  return QRunnable{}
}


func (this *QRunnable) run(args ...interface{}) () {
  // run()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QRunnable3runEv
  default:
    qtrt.ErrorResolve("QRunnable", "run", args)
  }

}


func (this *QRunnable) autoDelete(args ...interface{}) () {
  // autoDelete()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QRunnable10autoDeleteEv
  default:
    qtrt.ErrorResolve("QRunnable", "autoDelete", args)
  }

}

// <= body block end

