package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
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
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QAbstractState * QState::errorState();
extern void _ZNK6QState10errorStateEv(void* qthis);
  // proto:  QAbstractState * QState::initialState();
extern void _ZNK6QState12initialStateEv(void* qthis);
  // proto:  void QState::~QState();
extern void _ZN6QStateD0Ev(void* qthis);
  // proto:  void QState::assignProperty(QObject * object, const char * name, const QVariant & value);
extern void _ZN6QState14assignPropertyEP7QObjectPKcRK8QVariant(void* qthis, void* arg0, unsigned char* arg1, void* arg2);
  // proto:  void QState::QState(const QState & );
extern void* dector_ZN6QStateC1ERKS_(void* arg0);
extern void _ZN6QStateC1ERKS_(void* qthis, void* arg0);
  // proto:  void QState::QState(QState * parent);
extern void* dector_ZN6QStateC1EPS_(void* arg0);
extern void _ZN6QStateC1EPS_(void* qthis, void* arg0);
  // proto:  const QMetaObject * QState::metaObject();
extern void _ZNK6QState10metaObjectEv(void* qthis);
  // proto:  void QState::setErrorState(QAbstractState * state);
extern void _ZN6QState13setErrorStateEP14QAbstractState(void* qthis, void* arg0);
  // proto:  void QState::addTransition(QAbstractTransition * transition);
extern void _ZN6QState13addTransitionEP19QAbstractTransition(void* qthis, void* arg0);
  // proto:  void QState::removeTransition(QAbstractTransition * transition);
extern void _ZN6QState16removeTransitionEP19QAbstractTransition(void* qthis, void* arg0);
  // proto:  QSignalTransition * QState::addTransition(const QObject * sender, const char * signal, QAbstractState * target);
extern void _ZN6QState13addTransitionEPK7QObjectPKcP14QAbstractState(void* qthis, void* arg0, unsigned char* arg1, void* arg2);
  // proto:  QAbstractTransition * QState::addTransition(QAbstractState * target);
extern void _ZN6QState13addTransitionEP14QAbstractState(void* qthis, void* arg0);
  // proto:  QList<QAbstractTransition *> QState::transitions();
extern void _ZNK6QState11transitionsEv(void* qthis);
  // proto:  void QState::setInitialState(QAbstractState * state);
extern void _ZN6QState15setInitialStateEP14QAbstractState(void* qthis, void* arg0);
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
  qclsinst unsafe.Pointer /* *C.void */;
//  _childModeChanged QState_childModeChanged_signal;
//  _errorStateChanged QState_errorStateChanged_signal;
//  _finished QState_finished_signal;
//  _propertiesAssigned QState_propertiesAssigned_signal;
//  _initialStateChanged QState_initialStateChanged_signal;
}

  // proto:  QAbstractState * QState::errorState();
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
    // invoke: QAbstractState * errorState()
    C._ZNK6QState10errorStateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QState", "errorState", args)
  }

}

  // proto:  QAbstractState * QState::initialState();
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
    // invoke: QAbstractState * initialState()
    C._ZNK6QState12initialStateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QState", "initialState", args)
  }

}

  // proto:  void QState::~QState();
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

  // proto:  void QState::assignProperty(QObject * object, const char * name, const QVariant & value);
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
    // invoke: void assignProperty(class QObject *, const char *, const class QVariant &)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QVariant).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN6QState14assignPropertyEP7QObjectPKcRK8QVariant(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QState", "assignProperty", args)
  }

}

  // proto:  void QState::QState(const QState & );
func NewQState(args ...interface{}) QState {
  return QState{}
}

  // proto:  const QMetaObject * QState::metaObject();
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
    // invoke: const QMetaObject * metaObject()
    C._ZNK6QState10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QState", "metaObject", args)
  }

}

  // proto:  void QState::setErrorState(QAbstractState * state);
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
    // invoke: void setErrorState(class QAbstractState *)
    var arg0 = args[0].(QAbstractState).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN6QState13setErrorStateEP14QAbstractState(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QState", "setErrorState", args)
  }

}

  // proto:  void QState::addTransition(QAbstractTransition * transition);
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
    // invoke: void addTransition(class QAbstractTransition *)
    var arg0 = args[0].(QAbstractTransition).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN6QState13addTransitionEP19QAbstractTransition(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN6QState13addTransitionEPK7QObjectPKcP14QAbstractState
    // invoke: QSignalTransition * addTransition(const class QObject *, const char *, class QAbstractState *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QAbstractState).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN6QState13addTransitionEPK7QObjectPKcP14QAbstractState(this.qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN6QState13addTransitionEP14QAbstractState
    // invoke: QAbstractTransition * addTransition(class QAbstractState *)
    var arg0 = args[0].(QAbstractState).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN6QState13addTransitionEP14QAbstractState(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QState", "addTransition", args)
  }

}

  // proto:  void QState::removeTransition(QAbstractTransition * transition);
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
    // invoke: void removeTransition(class QAbstractTransition *)
    var arg0 = args[0].(QAbstractTransition).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN6QState16removeTransitionEP19QAbstractTransition(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QState", "removeTransition", args)
  }

}

  // proto:  QList<QAbstractTransition *> QState::transitions();
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
    // invoke: QList<QAbstractTransition *> transitions()
    C._ZNK6QState11transitionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QState", "transitions", args)
  }

}

  // proto:  void QState::setInitialState(QAbstractState * state);
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
    // invoke: void setInitialState(class QAbstractState *)
    var arg0 = args[0].(QAbstractState).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN6QState15setInitialStateEP14QAbstractState(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QState", "setInitialState", args)
  }

}

// <= body block end

