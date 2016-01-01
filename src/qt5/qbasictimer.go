package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qbasictimer.h
// dst-file: /src/core/qbasictimer.go
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
// class sizeof(QBasicTimer)=4
type QBasicTimer struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QBasicTimer) FreeQBasicTimer(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QBasicTimer", "~QBasicTimer", args)
  }

}


func (this *QBasicTimer) stop(args ...interface{}) () {
  // stop()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QBasicTimer4stopEv
  default:
    qtrt.ErrorResolve("QBasicTimer", "stop", args)
  }

}


func (this *QBasicTimer) timerId(args ...interface{}) () {
  // timerId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QBasicTimer7timerIdEv
  default:
    qtrt.ErrorResolve("QBasicTimer", "timerId", args)
  }

}


func (this *QBasicTimer) isActive(args ...interface{}) () {
  // isActive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QBasicTimer8isActiveEv
  default:
    qtrt.ErrorResolve("QBasicTimer", "isActive", args)
  }

}


func NewQBasicTimer(args ...interface{}) QBasicTimer {
  return QBasicTimer{}
}


func (this *QBasicTimer) start(args ...interface{}) () {
  // start(int, Qt::TimerType, class QObject *)
  // start(int, class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "Qt::TimerType"
  vtys[0][2] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QBasicTimer5startEiN2Qt9TimerTypeEP7QObject
  case 1:
    // invoke: _ZN11QBasicTimer5startEiP7QObject
  default:
    qtrt.ErrorResolve("QBasicTimer", "start", args)
  }

}

// <= body block end

