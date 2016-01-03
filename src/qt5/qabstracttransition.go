package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
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
  // proto:  QAbstractState * QAbstractTransition::targetState();
extern void _ZNK19QAbstractTransition11targetStateEv(void* qthis);
  // proto:  QList<QAbstractState *> QAbstractTransition::targetStates();
extern void _ZNK19QAbstractTransition12targetStatesEv(void* qthis);
  // proto:  QState * QAbstractTransition::sourceState();
extern void _ZNK19QAbstractTransition11sourceStateEv(void* qthis);
  // proto:  void QAbstractTransition::~QAbstractTransition();
extern void _ZN19QAbstractTransitionD0Ev(void* qthis);
  // proto:  void QAbstractTransition::QAbstractTransition(QState * sourceState);
extern void* dector_ZN19QAbstractTransitionC1EP6QState(void* arg0);
extern void _ZN19QAbstractTransitionC1EP6QState(void* qthis, void* arg0);
  // proto:  void QAbstractTransition::setTargetState(QAbstractState * target);
extern void _ZN19QAbstractTransition14setTargetStateEP14QAbstractState(void* qthis, void* arg0);
  // proto:  void QAbstractTransition::addAnimation(QAbstractAnimation * animation);
extern void _ZN19QAbstractTransition12addAnimationEP18QAbstractAnimation(void* qthis, void* arg0);
  // proto:  QList<QAbstractAnimation *> QAbstractTransition::animations();
extern void _ZNK19QAbstractTransition10animationsEv(void* qthis);
  // proto:  void QAbstractTransition::removeAnimation(QAbstractAnimation * animation);
extern void _ZN19QAbstractTransition15removeAnimationEP18QAbstractAnimation(void* qthis, void* arg0);
  // proto:  const QMetaObject * QAbstractTransition::metaObject();
extern void _ZNK19QAbstractTransition10metaObjectEv(void* qthis);
  // proto:  void QAbstractTransition::QAbstractTransition(const QAbstractTransition & );
extern void* dector_ZN19QAbstractTransitionC1ERKS_(void* arg0);
extern void _ZN19QAbstractTransitionC1ERKS_(void* qthis, void* arg0);
  // proto:  QStateMachine * QAbstractTransition::machine();
extern void _ZNK19QAbstractTransition7machineEv(void* qthis);
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

  // proto:  QAbstractState * QAbstractTransition::targetState();
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
    C._ZNK19QAbstractTransition11targetStateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractTransition", "targetState", args)
  }

}

  // proto:  QList<QAbstractState *> QAbstractTransition::targetStates();
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
    C._ZNK19QAbstractTransition12targetStatesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractTransition", "targetStates", args)
  }

}

  // proto:  QState * QAbstractTransition::sourceState();
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
    C._ZNK19QAbstractTransition11sourceStateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractTransition", "sourceState", args)
  }

}

  // proto:  void QAbstractTransition::~QAbstractTransition();
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

  // proto:  void QAbstractTransition::QAbstractTransition(QState * sourceState);
func NewQAbstractTransition(args ...interface{}) QAbstractTransition {
  return QAbstractTransition{}
}

  // proto:  void QAbstractTransition::setTargetState(QAbstractState * target);
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
    C._ZN19QAbstractTransition14setTargetStateEP14QAbstractState(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractTransition", "setTargetState", args)
  }

}

  // proto:  void QAbstractTransition::addAnimation(QAbstractAnimation * animation);
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
    C._ZN19QAbstractTransition12addAnimationEP18QAbstractAnimation(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractTransition", "addAnimation", args)
  }

}

  // proto:  QList<QAbstractAnimation *> QAbstractTransition::animations();
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
    C._ZNK19QAbstractTransition10animationsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractTransition", "animations", args)
  }

}

  // proto:  void QAbstractTransition::removeAnimation(QAbstractAnimation * animation);
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
    C._ZN19QAbstractTransition15removeAnimationEP18QAbstractAnimation(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractTransition", "removeAnimation", args)
  }

}

  // proto:  const QMetaObject * QAbstractTransition::metaObject();
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
    C._ZNK19QAbstractTransition10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractTransition", "metaObject", args)
  }

}

  // proto:  QStateMachine * QAbstractTransition::machine();
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
    C._ZNK19QAbstractTransition7machineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractTransition", "machine", args)
  }

}

// <= body block end

