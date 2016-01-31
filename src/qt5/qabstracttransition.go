package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtCore/qabstracttransition.h
// dst-file: /src/core/qabstracttransition.go
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
  // proto:  QAbstractTransition::TransitionType QAbstractTransition::transitionType();
extern void C_ZNK19QAbstractTransition14transitionTypeEv(void* qthis); // 4
  // proto:  QList<QAbstractState *> QAbstractTransition::targetStates();
extern void C_ZNK19QAbstractTransition12targetStatesEv(void* qthis); // 4
  // proto:  void QAbstractTransition::removeAnimation(QAbstractAnimation * animation);
extern void C_ZN19QAbstractTransition15removeAnimationEP18QAbstractAnimation(void* qthis, void* arg0); // 4
  // proto:  QStateMachine * QAbstractTransition::machine();
extern void C_ZNK19QAbstractTransition7machineEv(void* qthis); // 4
  // proto:  QAbstractState * QAbstractTransition::targetState();
extern void C_ZNK19QAbstractTransition11targetStateEv(void* qthis); // 4
  // proto:  QList<QAbstractAnimation *> QAbstractTransition::animations();
extern void C_ZNK19QAbstractTransition10animationsEv(void* qthis); // 4
  // proto:  const QMetaObject * QAbstractTransition::metaObject();
extern void C_ZNK19QAbstractTransition10metaObjectEv(void* qthis); // 4
  // proto:  void QAbstractTransition::setTargetState(QAbstractState * target);
extern void C_ZN19QAbstractTransition14setTargetStateEP14QAbstractState(void* qthis, void* arg0); // 4
  // proto:  void QAbstractTransition::~QAbstractTransition();
extern void C_ZN19QAbstractTransitionD2Ev(void* qthis); // 4
  // proto:  void QAbstractTransition::addAnimation(QAbstractAnimation * animation);
extern void C_ZN19QAbstractTransition12addAnimationEP18QAbstractAnimation(void* qthis, void* arg0); // 4
  // proto:  void QAbstractTransition::QAbstractTransition(QState * sourceState);
extern void* C_ZN19QAbstractTransitionC2EP6QState(void* arg0); // 3
  // proto:  QState * QAbstractTransition::sourceState();
extern void C_ZNK19QAbstractTransition11sourceStateEv(void* qthis); // 4
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

// class sizeof(QAbstractTransition)=1
type QAbstractTransition struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _targetStateChanged QAbstractTransition_targetStateChanged_signal;
//  _targetStatesChanged QAbstractTransition_targetStatesChanged_signal;
//  _triggered QAbstractTransition_triggered_signal;
}

// transitionType()
func (this *QAbstractTransition) transitionType(args ...interface{}) () {
  // transitionType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractTransition14transitionTypeEv
    // invoke: QAbstractTransition::TransitionType transitionType()
    C.C_ZNK19QAbstractTransition14transitionTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractTransition", "transitionType", args)
  }

}

// targetStates()
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
    // invoke: QList<QAbstractState *> targetStates()
    C.C_ZNK19QAbstractTransition12targetStatesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractTransition", "targetStates", args)
  }

}

// removeAnimation(class QAbstractAnimation *)
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
    // invoke: void removeAnimation(class QAbstractAnimation *)
    var arg0 = args[0].(QAbstractAnimation).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QAbstractTransition15removeAnimationEP18QAbstractAnimation(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractTransition", "removeAnimation", args)
  }

}

// machine()
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
    // invoke: QStateMachine * machine()
    var ret = C.C_ZNK19QAbstractTransition7machineEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAbstractTransition", "machine", args)
  }

}

// targetState()
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
    // invoke: QAbstractState * targetState()
    C.C_ZNK19QAbstractTransition11targetStateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractTransition", "targetState", args)
  }

}

// animations()
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
    // invoke: QList<QAbstractAnimation *> animations()
    C.C_ZNK19QAbstractTransition10animationsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractTransition", "animations", args)
  }

}

// metaObject()
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
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK19QAbstractTransition10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractTransition", "metaObject", args)
  }

}

// setTargetState(class QAbstractState *)
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
    // invoke: void setTargetState(class QAbstractState *)
    var arg0 = args[0].(QAbstractState).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QAbstractTransition14setTargetStateEP14QAbstractState(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractTransition", "setTargetState", args)
  }

}

// ~QAbstractTransition()
func (this *QAbstractTransition) FreeQAbstractTransition(args ...interface{}) () {
  // ~QAbstractTransition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractTransitionD0Ev
    // invoke: void ~QAbstractTransition()
    C.C_ZN19QAbstractTransitionD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractTransition", "~QAbstractTransition", args)
  }

}

// addAnimation(class QAbstractAnimation *)
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
    // invoke: void addAnimation(class QAbstractAnimation *)
    var arg0 = args[0].(QAbstractAnimation).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QAbstractTransition12addAnimationEP18QAbstractAnimation(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractTransition", "addAnimation", args)
  }

}

// QAbstractTransition(class QState *)
func NewQAbstractTransition(args ...interface{}) *QAbstractTransition {
  // QAbstractTransition(class QState *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QState{}) // "QState *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractTransitionC1EP6QState
    // invoke: void QAbstractTransition(class QState *)
    var arg0 = args[0].(QState).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QAbstractTransitionC2EP6QState(arg0)
    return &QAbstractTransition{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QAbstractTransition", "QAbstractTransition", args)
  }

  return nil // QAbstractTransition{qclsinst:qthis}
}

// sourceState()
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
    // invoke: QState * sourceState()
    var ret = C.C_ZNK19QAbstractTransition11sourceStateEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAbstractTransition", "sourceState", args)
  }

}

// <= body block end

