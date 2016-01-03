package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtCore/qstatemachine.h
// dst-file: /src/core/qstatemachine.go
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
  // proto:  QList<QAbstractAnimation *> QStateMachine::defaultAnimations();
extern void _ZNK13QStateMachine17defaultAnimationsEv(void* qthis);
  // proto:  int QStateMachine::postDelayedEvent(QEvent * event, int delay);
extern void _ZN13QStateMachine16postDelayedEventEP6QEventi(void* qthis, void* arg0, int32_t arg1);
  // proto:  QSet<QAbstractState *> QStateMachine::configuration();
extern void _ZNK13QStateMachine13configurationEv(void* qthis);
  // proto:  void QStateMachine::setRunning(bool running);
extern void _ZN13QStateMachine10setRunningEb(void* qthis, bool arg0);
  // proto:  void QStateMachine::addDefaultAnimation(QAbstractAnimation * animation);
extern void _ZN13QStateMachine19addDefaultAnimationEP18QAbstractAnimation(void* qthis, void* arg0);
  // proto:  void QStateMachine::removeDefaultAnimation(QAbstractAnimation * animation);
extern void _ZN13QStateMachine22removeDefaultAnimationEP18QAbstractAnimation(void* qthis, void* arg0);
  // proto:  void QStateMachine::setAnimated(bool enabled);
extern void _ZN13QStateMachine11setAnimatedEb(void* qthis, bool arg0);
  // proto:  void QStateMachine::QStateMachine(QObject * parent);
extern void* dector_ZN13QStateMachineC1EP7QObject(void* arg0);
extern void _ZN13QStateMachineC1EP7QObject(void* qthis, void* arg0);
  // proto:  QString QStateMachine::errorString();
extern void _ZNK13QStateMachine11errorStringEv(void* qthis);
  // proto:  bool QStateMachine::isRunning();
extern void _ZNK13QStateMachine9isRunningEv(void* qthis);
  // proto:  bool QStateMachine::cancelDelayedEvent(int id);
extern void _ZN13QStateMachine18cancelDelayedEventEi(void* qthis, int32_t arg0);
  // proto:  void QStateMachine::~QStateMachine();
extern void _ZN13QStateMachineD0Ev(void* qthis);
  // proto:  const QMetaObject * QStateMachine::metaObject();
extern void _ZNK13QStateMachine10metaObjectEv(void* qthis);
  // proto:  void QStateMachine::addState(QAbstractState * state);
extern void _ZN13QStateMachine8addStateEP14QAbstractState(void* qthis, void* arg0);
  // proto:  void QStateMachine::clearError();
extern void _ZN13QStateMachine10clearErrorEv(void* qthis);
  // proto:  void QStateMachine::removeState(QAbstractState * state);
extern void _ZN13QStateMachine11removeStateEP14QAbstractState(void* qthis, void* arg0);
  // proto:  void QStateMachine::stop();
extern void _ZN13QStateMachine4stopEv(void* qthis);
  // proto:  bool QStateMachine::isAnimated();
extern void _ZNK13QStateMachine10isAnimatedEv(void* qthis);
  // proto:  void QStateMachine::start();
extern void _ZN13QStateMachine5startEv(void* qthis);
  // proto:  bool QStateMachine::eventFilter(QObject * watched, QEvent * event);
extern void _ZN13QStateMachine11eventFilterEP7QObjectP6QEvent(void* qthis, void* arg0, void* arg1);
  // proto:  void QStateMachine::QStateMachine(const QStateMachine & );
extern void* dector_ZN13QStateMachineC1ERKS_(void* arg0);
extern void _ZN13QStateMachineC1ERKS_(void* qthis, void* arg0);
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

// class sizeof(QStateMachine)=1
type QStateMachine struct {
  /*qbase*/ QState;
  qclsinst unsafe.Pointer /* *C.void */;
//  _started QStateMachine_started_signal;
//  _runningChanged QStateMachine_runningChanged_signal;
//  _stopped QStateMachine_stopped_signal;
}

  // proto:  QList<QAbstractAnimation *> QStateMachine::defaultAnimations();
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
    // invoke: QList<QAbstractAnimation *> defaultAnimations()
    C._ZNK13QStateMachine17defaultAnimationsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStateMachine", "defaultAnimations", args)
  }

}

  // proto:  int QStateMachine::postDelayedEvent(QEvent * event, int delay);
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
    // invoke: int postDelayedEvent(class QEvent *, int)
    var arg0 = args[0].(QEvent).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN13QStateMachine16postDelayedEventEP6QEventi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStateMachine", "postDelayedEvent", args)
  }

}

  // proto:  QSet<QAbstractState *> QStateMachine::configuration();
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
    // invoke: QSet<QAbstractState *> configuration()
    C._ZNK13QStateMachine13configurationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStateMachine", "configuration", args)
  }

}

  // proto:  void QStateMachine::setRunning(bool running);
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
    // invoke: void setRunning(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN13QStateMachine10setRunningEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStateMachine", "setRunning", args)
  }

}

  // proto:  void QStateMachine::addDefaultAnimation(QAbstractAnimation * animation);
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
    // invoke: void addDefaultAnimation(class QAbstractAnimation *)
    var arg0 = args[0].(QAbstractAnimation).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QStateMachine19addDefaultAnimationEP18QAbstractAnimation(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStateMachine", "addDefaultAnimation", args)
  }

}

  // proto:  void QStateMachine::removeDefaultAnimation(QAbstractAnimation * animation);
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
    // invoke: void removeDefaultAnimation(class QAbstractAnimation *)
    var arg0 = args[0].(QAbstractAnimation).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QStateMachine22removeDefaultAnimationEP18QAbstractAnimation(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStateMachine", "removeDefaultAnimation", args)
  }

}

  // proto:  void QStateMachine::setAnimated(bool enabled);
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
    // invoke: void setAnimated(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN13QStateMachine11setAnimatedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStateMachine", "setAnimated", args)
  }

}

  // proto:  void QStateMachine::QStateMachine(QObject * parent);
func NewQStateMachine(args ...interface{}) QStateMachine {
  return QStateMachine{}
}

  // proto:  QString QStateMachine::errorString();
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
    // invoke: QString errorString()
    C._ZNK13QStateMachine11errorStringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStateMachine", "errorString", args)
  }

}

  // proto:  bool QStateMachine::isRunning();
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
    // invoke: bool isRunning()
    C._ZNK13QStateMachine9isRunningEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStateMachine", "isRunning", args)
  }

}

  // proto:  bool QStateMachine::cancelDelayedEvent(int id);
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
    // invoke: bool cancelDelayedEvent(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN13QStateMachine18cancelDelayedEventEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStateMachine", "cancelDelayedEvent", args)
  }

}

  // proto:  void QStateMachine::~QStateMachine();
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

  // proto:  const QMetaObject * QStateMachine::metaObject();
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
    // invoke: const QMetaObject * metaObject()
    C._ZNK13QStateMachine10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStateMachine", "metaObject", args)
  }

}

  // proto:  void QStateMachine::addState(QAbstractState * state);
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
    // invoke: void addState(class QAbstractState *)
    var arg0 = args[0].(QAbstractState).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QStateMachine8addStateEP14QAbstractState(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStateMachine", "addState", args)
  }

}

  // proto:  void QStateMachine::clearError();
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
    // invoke: void clearError()
    C._ZN13QStateMachine10clearErrorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStateMachine", "clearError", args)
  }

}

  // proto:  void QStateMachine::removeState(QAbstractState * state);
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
    // invoke: void removeState(class QAbstractState *)
    var arg0 = args[0].(QAbstractState).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QStateMachine11removeStateEP14QAbstractState(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStateMachine", "removeState", args)
  }

}

  // proto:  void QStateMachine::stop();
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
    // invoke: void stop()
    C._ZN13QStateMachine4stopEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStateMachine", "stop", args)
  }

}

  // proto:  bool QStateMachine::isAnimated();
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
    // invoke: bool isAnimated()
    C._ZNK13QStateMachine10isAnimatedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStateMachine", "isAnimated", args)
  }

}

  // proto:  void QStateMachine::start();
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
    // invoke: void start()
    C._ZN13QStateMachine5startEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStateMachine", "start", args)
  }

}

  // proto:  bool QStateMachine::eventFilter(QObject * watched, QEvent * event);
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
    // invoke: bool eventFilter(class QObject *, class QEvent *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QEvent).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN13QStateMachine11eventFilterEP7QObjectP6QEvent(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStateMachine", "eventFilter", args)
  }

}

// <= body block end

