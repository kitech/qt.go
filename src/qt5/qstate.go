package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtCore/qstate.h
// dst-file: /src/core/qstate.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QState::QState(QState * parent);
extern void* C_ZN6QStateC2EPS_(void* arg0); // 3
  // proto:  QList<QAbstractTransition *> QState::transitions();
extern void C_ZNK6QState11transitionsEv(void* qthis); // 4
  // proto:  QAbstractState * QState::errorState();
extern void C_ZNK6QState10errorStateEv(void* qthis); // 4
  // proto:  void QState::removeTransition(QAbstractTransition * transition);
extern void C_ZN6QState16removeTransitionEP19QAbstractTransition(void* qthis, void* arg0); // 4
  // proto:  QAbstractTransition * QState::addTransition(QAbstractState * target);
extern void C_ZN6QState13addTransitionEP14QAbstractState(void* qthis, void* arg0); // 4
  // proto:  void QState::addTransition(QAbstractTransition * transition);
extern void C_ZN6QState13addTransitionEP19QAbstractTransition(void* qthis, void* arg0); // 4
  // proto:  QSignalTransition * QState::addTransition(const QObject * sender, const char * signal, QAbstractState * target);
extern void* C_ZN6QState13addTransitionEPK7QObjectPKcP14QAbstractState(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QState::setErrorState(QAbstractState * state);
extern void C_ZN6QState13setErrorStateEP14QAbstractState(void* qthis, void* arg0); // 4
  // proto:  QAbstractState * QState::initialState();
extern void C_ZNK6QState12initialStateEv(void* qthis); // 4
  // proto:  QState::ChildMode QState::childMode();
extern void C_ZNK6QState9childModeEv(void* qthis); // 4
  // proto:  void QState::assignProperty(QObject * object, const char * name, const QVariant & value);
extern void C_ZN6QState14assignPropertyEP7QObjectPKcRK8QVariant(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QState::~QState();
extern void C_ZN6QStateD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QState::metaObject();
extern void C_ZNK6QState10metaObjectEv(void* qthis); // 4
  // proto:  void QState::setInitialState(QAbstractState * state);
extern void C_ZN6QState15setInitialStateEP14QAbstractState(void* qthis, void* arg0); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QState)=1
type QState struct {
  /*qbase*/ QAbstractState;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _childModeChanged QState_childModeChanged_signal;
//  _errorStateChanged QState_errorStateChanged_signal;
//  _finished QState_finished_signal;
//  _propertiesAssigned QState_propertiesAssigned_signal;
//  _initialStateChanged QState_initialStateChanged_signal;
}

// QState(class QState *)
func NewQState(args ...interface{}) *QState {
  // QState(class QState *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QState{}) // "QState *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QStateC1EPS_
    // invoke: void QState(class QState *)
    var arg0 = args[0].(*QState).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QStateC2EPS_(arg0)
    return &QState{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QState", "QState", args)
  }

  return nil // QState{Qclsinst:qthis}
}

// transitions()
func (this *QState) Transitions(args ...interface{}) () {
  // transitions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QState11transitionsEv
    // invoke: QList<QAbstractTransition *> transitions()
    C.C_ZNK6QState11transitionsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QState", "transitions", args)
  }

  return
}

// errorState()
func (this *QState) Errorstate(args ...interface{}) () {
  // errorState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QState10errorStateEv
    // invoke: QAbstractState * errorState()
    C.C_ZNK6QState10errorStateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QState", "errorState", args)
  }

  return
}

// removeTransition(class QAbstractTransition *)
func (this *QState) Removetransition(args ...interface{}) () {
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
    // invoke: void removeTransition(class QAbstractTransition *)
    var arg0 = args[0].(*QAbstractTransition).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QState16removeTransitionEP19QAbstractTransition(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QState", "removeTransition", args)
  }

  return
}

// addTransition(class QAbstractState *)
func (this *QState) Addtransition(args ...interface{}) () {
  // addTransition(class QAbstractState *)
  // addTransition(class QAbstractTransition *)
  // addTransition(const class QObject *, const char *, class QAbstractState *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractState{}) // "QAbstractState *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QAbstractTransition{}) // "QAbstractTransition *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[2][1] = qtrt.ByteTy(true) // "const char *"
  vtys[2][2] = reflect.TypeOf(QAbstractState{}) // "QAbstractState *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QState13addTransitionEP14QAbstractState
    // invoke: QAbstractTransition * addTransition(class QAbstractState *)
    var arg0 = args[0].(*QAbstractState).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QState13addTransitionEP14QAbstractState(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN6QState13addTransitionEP19QAbstractTransition
    // invoke: void addTransition(class QAbstractTransition *)
    var arg0 = args[0].(*QAbstractTransition).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QState13addTransitionEP19QAbstractTransition(this.Qclsinst, arg0)
  case 2:
    // invoke: _ZN6QState13addTransitionEPK7QObjectPKcP14QAbstractState
    // invoke: QSignalTransition * addTransition(const class QObject *, const char *, class QAbstractState *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[2][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    var arg2 = args[2].(*QAbstractState).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN6QState13addTransitionEPK7QObjectPKcP14QAbstractState(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
  default:
    qtrt.ErrorResolve("QState", "addTransition", args)
  }

  return
}

// setErrorState(class QAbstractState *)
func (this *QState) Seterrorstate(args ...interface{}) () {
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
    // invoke: void setErrorState(class QAbstractState *)
    var arg0 = args[0].(*QAbstractState).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QState13setErrorStateEP14QAbstractState(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QState", "setErrorState", args)
  }

  return
}

// initialState()
func (this *QState) Initialstate(args ...interface{}) () {
  // initialState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QState12initialStateEv
    // invoke: QAbstractState * initialState()
    C.C_ZNK6QState12initialStateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QState", "initialState", args)
  }

  return
}

// childMode()
func (this *QState) Childmode(args ...interface{}) () {
  // childMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QState9childModeEv
    // invoke: QState::ChildMode childMode()
    C.C_ZNK6QState9childModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QState", "childMode", args)
  }

  return
}

// assignProperty(class QObject *, const char *, const class QVariant &)
func (this *QState) Assignproperty(args ...interface{}) () {
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
    // invoke: void assignProperty(class QObject *, const char *, const class QVariant &)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[0][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    var arg2 = args[2].(*QVariant).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN6QState14assignPropertyEP7QObjectPKcRK8QVariant(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QState", "assignProperty", args)
  }

  return
}

// ~QState()
func (this *QState) Freeqstate(args ...interface{}) () {
  // ~QState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QStateD0Ev
    // invoke: void ~QState()
    C.C_ZN6QStateD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QState", "~QState", args)
  }

  return
}

// metaObject()
func (this *QState) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QState10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK6QState10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QState", "metaObject", args)
  }

  return
}

// setInitialState(class QAbstractState *)
func (this *QState) Setinitialstate(args ...interface{}) () {
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
    // invoke: void setInitialState(class QAbstractState *)
    var arg0 = args[0].(*QAbstractState).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QState15setInitialStateEP14QAbstractState(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QState", "setInitialState", args)
  }

  return
}

// <= body block end

