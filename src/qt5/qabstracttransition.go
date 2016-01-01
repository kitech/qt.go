package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qabstracttransition.h
// dst-file: /src/core/qabstracttransition.go
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
// class sizeof(QAbstractTransition)=1
type QAbstractTransition struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _targetStateChanged QAbstractTransition_targetStateChanged_signal;
//  _targetStatesChanged QAbstractTransition_targetStatesChanged_signal;
//  _triggered QAbstractTransition_triggered_signal;
}


func (this *QAbstractTransition) targetState(args ...interface{}) () {
  // targetState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractTransition11targetStateEv
  default:
    qtrt.ErrorResolve("QAbstractTransition", "targetState", args)
 }

}


func (this *QAbstractTransition) targetStates(args ...interface{}) () {
  // targetStates()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractTransition12targetStatesEv
  default:
    qtrt.ErrorResolve("QAbstractTransition", "targetStates", args)
 }

}


func (this *QAbstractTransition) sourceState(args ...interface{}) () {
  // sourceState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractTransition11sourceStateEv
  default:
    qtrt.ErrorResolve("QAbstractTransition", "sourceState", args)
 }

}


func (this *QAbstractTransition) FreeQAbstractTransition(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractTransition", "~QAbstractTransition", args)
 }

}


func NewQAbstractTransition(args ...interface{})() {
}


func (this *QAbstractTransition) setTargetState(args ...interface{}) () {
  // setTargetState(class QAbstractState *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractState{}) // "QAbstractState *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractTransition14setTargetStateEP14QAbstractState
  default:
    qtrt.ErrorResolve("QAbstractTransition", "setTargetState", args)
 }

}


func (this *QAbstractTransition) addAnimation(args ...interface{}) () {
  // addAnimation(class QAbstractAnimation *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractAnimation{}) // "QAbstractAnimation *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractTransition12addAnimationEP18QAbstractAnimation
  default:
    qtrt.ErrorResolve("QAbstractTransition", "addAnimation", args)
 }

}


func (this *QAbstractTransition) animations(args ...interface{}) () {
  // animations()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractTransition10animationsEv
  default:
    qtrt.ErrorResolve("QAbstractTransition", "animations", args)
 }

}


func (this *QAbstractTransition) removeAnimation(args ...interface{}) () {
  // removeAnimation(class QAbstractAnimation *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractAnimation{}) // "QAbstractAnimation *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractTransition15removeAnimationEP18QAbstractAnimation
  default:
    qtrt.ErrorResolve("QAbstractTransition", "removeAnimation", args)
 }

}


func (this *QAbstractTransition) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractTransition10metaObjectEv
  default:
    qtrt.ErrorResolve("QAbstractTransition", "metaObject", args)
 }

}


func (this *QAbstractTransition) machine(args ...interface{}) () {
  // machine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractTransition7machineEv
  default:
    qtrt.ErrorResolve("QAbstractTransition", "machine", args)
 }

}

// <= body block end

