package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qstatemachine.h
// dst-file: /src/core/qstatemachine.go
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
// class sizeof(QStateMachine)=1
type QStateMachine struct {
  /*qbase*/ QState;
  qclsinst uint64 /* *mut c_void*/;
//  _started QStateMachine_started_signal;
//  _runningChanged QStateMachine_runningChanged_signal;
//  _stopped QStateMachine_stopped_signal;
}


func (this *QStateMachine) defaultAnimations(args ...interface{}) () {
  // defaultAnimations()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStateMachine17defaultAnimationsEv
  default:
    qtrt.ErrorResolve("QStateMachine", "defaultAnimations", args)
  }

}


func (this *QStateMachine) postDelayedEvent(args ...interface{}) () {
  // postDelayedEvent(class QEvent *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QEvent{}) // "QEvent *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStateMachine16postDelayedEventEP6QEventi
  default:
    qtrt.ErrorResolve("QStateMachine", "postDelayedEvent", args)
  }

}


func (this *QStateMachine) configuration(args ...interface{}) () {
  // configuration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStateMachine13configurationEv
  default:
    qtrt.ErrorResolve("QStateMachine", "configuration", args)
  }

}


func (this *QStateMachine) setRunning(args ...interface{}) () {
  // setRunning(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStateMachine10setRunningEb
  default:
    qtrt.ErrorResolve("QStateMachine", "setRunning", args)
  }

}


func (this *QStateMachine) addDefaultAnimation(args ...interface{}) () {
  // addDefaultAnimation(class QAbstractAnimation *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractAnimation{}) // "QAbstractAnimation *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStateMachine19addDefaultAnimationEP18QAbstractAnimation
  default:
    qtrt.ErrorResolve("QStateMachine", "addDefaultAnimation", args)
  }

}


func (this *QStateMachine) removeDefaultAnimation(args ...interface{}) () {
  // removeDefaultAnimation(class QAbstractAnimation *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractAnimation{}) // "QAbstractAnimation *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStateMachine22removeDefaultAnimationEP18QAbstractAnimation
  default:
    qtrt.ErrorResolve("QStateMachine", "removeDefaultAnimation", args)
  }

}


func (this *QStateMachine) setAnimated(args ...interface{}) () {
  // setAnimated(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStateMachine11setAnimatedEb
  default:
    qtrt.ErrorResolve("QStateMachine", "setAnimated", args)
  }

}


func NewQStateMachine(args ...interface{}) QStateMachine {
  return QStateMachine{}
}


func (this *QStateMachine) errorString(args ...interface{}) () {
  // errorString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStateMachine11errorStringEv
  default:
    qtrt.ErrorResolve("QStateMachine", "errorString", args)
  }

}


func (this *QStateMachine) isRunning(args ...interface{}) () {
  // isRunning()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStateMachine9isRunningEv
  default:
    qtrt.ErrorResolve("QStateMachine", "isRunning", args)
  }

}


func (this *QStateMachine) cancelDelayedEvent(args ...interface{}) () {
  // cancelDelayedEvent(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStateMachine18cancelDelayedEventEi
  default:
    qtrt.ErrorResolve("QStateMachine", "cancelDelayedEvent", args)
  }

}


func (this *QStateMachine) FreeQStateMachine(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStateMachine", "~QStateMachine", args)
  }

}


func (this *QStateMachine) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStateMachine10metaObjectEv
  default:
    qtrt.ErrorResolve("QStateMachine", "metaObject", args)
  }

}


func (this *QStateMachine) addState(args ...interface{}) () {
  // addState(class QAbstractState *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractState{}) // "QAbstractState *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStateMachine8addStateEP14QAbstractState
  default:
    qtrt.ErrorResolve("QStateMachine", "addState", args)
  }

}


func (this *QStateMachine) clearError(args ...interface{}) () {
  // clearError()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStateMachine10clearErrorEv
  default:
    qtrt.ErrorResolve("QStateMachine", "clearError", args)
  }

}


func (this *QStateMachine) removeState(args ...interface{}) () {
  // removeState(class QAbstractState *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractState{}) // "QAbstractState *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStateMachine11removeStateEP14QAbstractState
  default:
    qtrt.ErrorResolve("QStateMachine", "removeState", args)
  }

}


func (this *QStateMachine) stop(args ...interface{}) () {
  // stop()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStateMachine4stopEv
  default:
    qtrt.ErrorResolve("QStateMachine", "stop", args)
  }

}


func (this *QStateMachine) isAnimated(args ...interface{}) () {
  // isAnimated()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStateMachine10isAnimatedEv
  default:
    qtrt.ErrorResolve("QStateMachine", "isAnimated", args)
  }

}


func (this *QStateMachine) start(args ...interface{}) () {
  // start()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStateMachine5startEv
  default:
    qtrt.ErrorResolve("QStateMachine", "start", args)
  }

}


func (this *QStateMachine) eventFilter(args ...interface{}) () {
  // eventFilter(class QObject *, class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[0][1] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStateMachine11eventFilterEP7QObjectP6QEvent
  default:
    qtrt.ErrorResolve("QStateMachine", "eventFilter", args)
  }

}

// <= body block end

