package qtcore
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
extern void* C_ZNK19QAbstractTransition7machineEv(void* qthis); // 4
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
extern void* C_ZNK19QAbstractTransition11sourceStateEv(void* qthis); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
//  _targetStateChanged QAbstractTransition_targetStateChanged_signal;
//  _targetStatesChanged QAbstractTransition_targetStatesChanged_signal;
//  _triggered QAbstractTransition_triggered_signal;
}

// transitionType()
func (this *QAbstractTransition) Transitiontype(args ...interface{}) () {
  // transitionType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractTransition14transitionTypeEv
    // invoke: QAbstractTransition::TransitionType transitionType()
    C.C_ZNK19QAbstractTransition14transitionTypeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractTransition", "transitionType", args)
  }

  return
}

// targetStates()
func (this *QAbstractTransition) Targetstates(args ...interface{}) () {
  // targetStates()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractTransition12targetStatesEv
    // invoke: QList<QAbstractState *> targetStates()
    C.C_ZNK19QAbstractTransition12targetStatesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractTransition", "targetStates", args)
  }

  return
}

// removeAnimation(class QAbstractAnimation *)
func (this *QAbstractTransition) Removeanimation(args ...interface{}) () {
  // removeAnimation(class QAbstractAnimation *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractAnimation{}) // "QAbstractAnimation *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractTransition15removeAnimationEP18QAbstractAnimation
    // invoke: void removeAnimation(class QAbstractAnimation *)
    var arg0 = args[0].(*QAbstractAnimation).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QAbstractTransition15removeAnimationEP18QAbstractAnimation(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractTransition", "removeAnimation", args)
  }

  return
}

// machine()
func (this *QAbstractTransition) Machine(args ...interface{}) (ret interface{}) {
  // machine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractTransition7machineEv
    // invoke: QStateMachine * machine()
    var ret0 = C.C_ZNK19QAbstractTransition7machineEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStateMachine{}) // "QStateMachine *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractTransition", "machine", args)
  }

  return
}

// targetState()
func (this *QAbstractTransition) Targetstate(args ...interface{}) () {
  // targetState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractTransition11targetStateEv
    // invoke: QAbstractState * targetState()
    C.C_ZNK19QAbstractTransition11targetStateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractTransition", "targetState", args)
  }

  return
}

// animations()
func (this *QAbstractTransition) Animations(args ...interface{}) () {
  // animations()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractTransition10animationsEv
    // invoke: QList<QAbstractAnimation *> animations()
    C.C_ZNK19QAbstractTransition10animationsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractTransition", "animations", args)
  }

  return
}

// metaObject()
func (this *QAbstractTransition) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractTransition10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK19QAbstractTransition10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractTransition", "metaObject", args)
  }

  return
}

// setTargetState(class QAbstractState *)
func (this *QAbstractTransition) Settargetstate(args ...interface{}) () {
  // setTargetState(class QAbstractState *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractState{}) // "QAbstractState *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractTransition14setTargetStateEP14QAbstractState
    // invoke: void setTargetState(class QAbstractState *)
    var arg0 = args[0].(*QAbstractState).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QAbstractTransition14setTargetStateEP14QAbstractState(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractTransition", "setTargetState", args)
  }

  return
}

// ~QAbstractTransition()
func (this *QAbstractTransition) Freeqabstracttransition(args ...interface{}) () {
  // ~QAbstractTransition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractTransitionD0Ev
    // invoke: void ~QAbstractTransition()
    C.C_ZN19QAbstractTransitionD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractTransition", "~QAbstractTransition", args)
  }

  return
}

// addAnimation(class QAbstractAnimation *)
func (this *QAbstractTransition) Addanimation(args ...interface{}) () {
  // addAnimation(class QAbstractAnimation *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractAnimation{}) // "QAbstractAnimation *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractTransition12addAnimationEP18QAbstractAnimation
    // invoke: void addAnimation(class QAbstractAnimation *)
    var arg0 = args[0].(*QAbstractAnimation).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QAbstractTransition12addAnimationEP18QAbstractAnimation(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractTransition", "addAnimation", args)
  }

  return
}

// QAbstractTransition(class QState *)
func NewQAbstractTransition(args ...interface{}) *QAbstractTransition {
  // QAbstractTransition(class QState *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QState{}) // "QState *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractTransitionC1EP6QState
    // invoke: void QAbstractTransition(class QState *)
    var arg0 = args[0].(*QState).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QAbstractTransitionC2EP6QState(arg0)
    return &QAbstractTransition{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QAbstractTransition", "QAbstractTransition", args)
  }

  return nil // QAbstractTransition{Qclsinst:qthis}
}

// sourceState()
func (this *QAbstractTransition) Sourcestate(args ...interface{}) (ret interface{}) {
  // sourceState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractTransition11sourceStateEv
    // invoke: QState * sourceState()
    var ret0 = C.C_ZNK19QAbstractTransition11sourceStateEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QState{}) // "QState *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractTransition", "sourceState", args)
  }

  return
}

// <= body block end

