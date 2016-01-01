package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qstate.h
// dst-file: /src/core/qstate.go
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
// class sizeof(QState)=1
type QState struct {
  /*qbase*/ QAbstractState;
  qclsinst uint64 /* *mut c_void*/;
//  _childModeChanged QState_childModeChanged_signal;
//  _errorStateChanged QState_errorStateChanged_signal;
//  _finished QState_finished_signal;
//  _propertiesAssigned QState_propertiesAssigned_signal;
//  _initialStateChanged QState_initialStateChanged_signal;
}


func (this *QState) errorState(args ...interface{}) () {
  // errorState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK6QState10errorStateEv
  default:
    qtrt.ErrorResolve("QState", "errorState", args)
 }

}


func (this *QState) initialState(args ...interface{}) () {
  // initialState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK6QState12initialStateEv
  default:
    qtrt.ErrorResolve("QState", "initialState", args)
 }

}


func (this *QState) FreeQState(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QState", "~QState", args)
 }

}


func (this *QState) assignProperty(args ...interface{}) () {
  // assignProperty(class QObject *, const char *, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[0][2] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN6QState14assignPropertyEP7QObjectPKcRK8QVariant
  default:
    qtrt.ErrorResolve("QState", "assignProperty", args)
 }

}


func NewQState(args ...interface{})() {
}


func (this *QState) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK6QState10metaObjectEv
  default:
    qtrt.ErrorResolve("QState", "metaObject", args)
 }

}


func (this *QState) setErrorState(args ...interface{}) () {
  // setErrorState(class QAbstractState *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractState{}) // "QAbstractState *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN6QState13setErrorStateEP14QAbstractState
  default:
    qtrt.ErrorResolve("QState", "setErrorState", args)
 }

}


func (this *QState) addTransition(args ...interface{}) () {
  // addTransition(class QAbstractTransition *)
  // addTransition(const class QObject *, const char *, class QAbstractState *)
  // addTransition(class QAbstractState *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractTransition{}) // "QAbstractTransition *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[1][1] = qtrt.ByteTy(true) // "const char *"
  vtys[1][2] = reflect.TypeOf(QAbstractState{}) // "QAbstractState *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QAbstractState{}) // "QAbstractState *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN6QState13addTransitionEP19QAbstractTransition
  case 1:
    // invoke: _ZN6QState13addTransitionEPK7QObjectPKcP14QAbstractState
  case 2:
    // invoke: _ZN6QState13addTransitionEP14QAbstractState
  default:
    qtrt.ErrorResolve("QState", "addTransition", args)
 }

}


func (this *QState) removeTransition(args ...interface{}) () {
  // removeTransition(class QAbstractTransition *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractTransition{}) // "QAbstractTransition *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN6QState16removeTransitionEP19QAbstractTransition
  default:
    qtrt.ErrorResolve("QState", "removeTransition", args)
 }

}


func (this *QState) transitions(args ...interface{}) () {
  // transitions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK6QState11transitionsEv
  default:
    qtrt.ErrorResolve("QState", "transitions", args)
 }

}


func (this *QState) setInitialState(args ...interface{}) () {
  // setInitialState(class QAbstractState *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractState{}) // "QAbstractState *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN6QState15setInitialStateEP14QAbstractState
  default:
    qtrt.ErrorResolve("QState", "setInitialState", args)
 }

}

// <= body block end

