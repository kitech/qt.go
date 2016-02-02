package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QList<QAbstractAnimation *> QStateMachine::defaultAnimations();
extern void C_ZNK13QStateMachine17defaultAnimationsEv(void* qthis); // 4
  // proto:  QState::RestorePolicy QStateMachine::globalRestorePolicy();
extern void C_ZNK13QStateMachine19globalRestorePolicyEv(void* qthis); // 4
  // proto:  void QStateMachine::addDefaultAnimation(QAbstractAnimation * animation);
extern void C_ZN13QStateMachine19addDefaultAnimationEP18QAbstractAnimation(void* qthis, void* arg0); // 4
  // proto:  void QStateMachine::clearError();
extern void C_ZN13QStateMachine10clearErrorEv(void* qthis); // 4
  // proto:  bool QStateMachine::isAnimated();
extern bool C_ZNK13QStateMachine10isAnimatedEv(void* qthis); // 4
  // proto:  void QStateMachine::addState(QAbstractState * state);
extern void C_ZN13QStateMachine8addStateEP14QAbstractState(void* qthis, void* arg0); // 4
  // proto:  bool QStateMachine::eventFilter(QObject * watched, QEvent * event);
extern bool C_ZN13QStateMachine11eventFilterEP7QObjectP6QEvent(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QStateMachine::start();
extern void C_ZN13QStateMachine5startEv(void* qthis); // 4
  // proto:  bool QStateMachine::cancelDelayedEvent(int id);
extern bool C_ZN13QStateMachine18cancelDelayedEventEi(void* qthis, int32_t arg0); // 4
  // proto:  void QStateMachine::setRunning(bool running);
extern void C_ZN13QStateMachine10setRunningEb(void* qthis, bool arg0); // 4
  // proto:  void QStateMachine::removeState(QAbstractState * state);
extern void C_ZN13QStateMachine11removeStateEP14QAbstractState(void* qthis, void* arg0); // 4
  // proto:  QString QStateMachine::errorString();
extern void* C_ZNK13QStateMachine11errorStringEv(void* qthis); // 4
  // proto:  void QStateMachine::stop();
extern void C_ZN13QStateMachine4stopEv(void* qthis); // 4
  // proto:  void QStateMachine::setAnimated(bool enabled);
extern void C_ZN13QStateMachine11setAnimatedEb(void* qthis, bool arg0); // 4
  // proto:  const QMetaObject * QStateMachine::metaObject();
extern void C_ZNK13QStateMachine10metaObjectEv(void* qthis); // 4
  // proto:  QSet<QAbstractState *> QStateMachine::configuration();
extern void C_ZNK13QStateMachine13configurationEv(void* qthis); // 4
  // proto:  bool QStateMachine::isRunning();
extern bool C_ZNK13QStateMachine9isRunningEv(void* qthis); // 4
  // proto:  void QStateMachine::~QStateMachine();
extern void C_ZN13QStateMachineD2Ev(void* qthis); // 4
  // proto:  void QStateMachine::QStateMachine(QObject * parent);
extern void* C_ZN13QStateMachineC2EP7QObject(void* arg0); // 3
  // proto:  int QStateMachine::postDelayedEvent(QEvent * event, int delay);
extern int32_t C_ZN13QStateMachine16postDelayedEventEP6QEventi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QStateMachine::Error QStateMachine::error();
extern void C_ZNK13QStateMachine5errorEv(void* qthis); // 4
  // proto:  void QStateMachine::removeDefaultAnimation(QAbstractAnimation * animation);
extern void C_ZN13QStateMachine22removeDefaultAnimationEP18QAbstractAnimation(void* qthis, void* arg0); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
//  _started QStateMachine_started_signal;
//  _runningChanged QStateMachine_runningChanged_signal;
//  _stopped QStateMachine_stopped_signal;
}

// defaultAnimations()
func (this *QStateMachine) Defaultanimations(args ...interface{}) () {
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
    C.C_ZNK13QStateMachine17defaultAnimationsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStateMachine", "defaultAnimations", args)
  }

  return
}

// globalRestorePolicy()
func (this *QStateMachine) Globalrestorepolicy(args ...interface{}) () {
  // globalRestorePolicy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStateMachine19globalRestorePolicyEv
    // invoke: QState::RestorePolicy globalRestorePolicy()
    C.C_ZNK13QStateMachine19globalRestorePolicyEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStateMachine", "globalRestorePolicy", args)
  }

  return
}

// addDefaultAnimation(class QAbstractAnimation *)
func (this *QStateMachine) Adddefaultanimation(args ...interface{}) () {
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
    var arg0 = args[0].(*QAbstractAnimation).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QStateMachine19addDefaultAnimationEP18QAbstractAnimation(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStateMachine", "addDefaultAnimation", args)
  }

  return
}

// clearError()
func (this *QStateMachine) Clearerror(args ...interface{}) () {
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
    C.C_ZN13QStateMachine10clearErrorEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStateMachine", "clearError", args)
  }

  return
}

// isAnimated()
func (this *QStateMachine) Isanimated(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStateMachine10isAnimatedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStateMachine", "isAnimated", args)
  }

  return
}

// addState(class QAbstractState *)
func (this *QStateMachine) Addstate(args ...interface{}) () {
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
    var arg0 = args[0].(*QAbstractState).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QStateMachine8addStateEP14QAbstractState(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStateMachine", "addState", args)
  }

  return
}

// eventFilter(class QObject *, class QEvent *)
func (this *QStateMachine) Eventfilter(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QEvent).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN13QStateMachine11eventFilterEP7QObjectP6QEvent(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStateMachine", "eventFilter", args)
  }

  return
}

// start()
func (this *QStateMachine) Start(args ...interface{}) () {
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
    C.C_ZN13QStateMachine5startEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStateMachine", "start", args)
  }

  return
}

// cancelDelayedEvent(int)
func (this *QStateMachine) Canceldelayedevent(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN13QStateMachine18cancelDelayedEventEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStateMachine", "cancelDelayedEvent", args)
  }

  return
}

// setRunning(_Bool)
func (this *QStateMachine) Setrunning(args ...interface{}) () {
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
    C.C_ZN13QStateMachine10setRunningEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStateMachine", "setRunning", args)
  }

  return
}

// removeState(class QAbstractState *)
func (this *QStateMachine) Removestate(args ...interface{}) () {
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
    var arg0 = args[0].(*QAbstractState).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QStateMachine11removeStateEP14QAbstractState(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStateMachine", "removeState", args)
  }

  return
}

// errorString()
func (this *QStateMachine) Errorstring(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStateMachine11errorStringEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStateMachine", "errorString", args)
  }

  return
}

// stop()
func (this *QStateMachine) Stop(args ...interface{}) () {
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
    C.C_ZN13QStateMachine4stopEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStateMachine", "stop", args)
  }

  return
}

// setAnimated(_Bool)
func (this *QStateMachine) Setanimated(args ...interface{}) () {
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
    C.C_ZN13QStateMachine11setAnimatedEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStateMachine", "setAnimated", args)
  }

  return
}

// metaObject()
func (this *QStateMachine) Metaobject(args ...interface{}) () {
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
    C.C_ZNK13QStateMachine10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStateMachine", "metaObject", args)
  }

  return
}

// configuration()
func (this *QStateMachine) Configuration(args ...interface{}) () {
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
    C.C_ZNK13QStateMachine13configurationEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStateMachine", "configuration", args)
  }

  return
}

// isRunning()
func (this *QStateMachine) Isrunning(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStateMachine9isRunningEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStateMachine", "isRunning", args)
  }

  return
}

// ~QStateMachine()
func (this *QStateMachine) Freeqstatemachine(args ...interface{}) () {
  // ~QStateMachine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStateMachineD0Ev
    // invoke: void ~QStateMachine()
    C.C_ZN13QStateMachineD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStateMachine", "~QStateMachine", args)
  }

  return
}

// QStateMachine(class QObject *)
func NewQStateMachine(args ...interface{}) *QStateMachine {
  // QStateMachine(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStateMachineC1EP7QObject
    // invoke: void QStateMachine(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QStateMachineC2EP7QObject(arg0)
    return &QStateMachine{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStateMachine", "QStateMachine", args)
  }

  return nil // QStateMachine{Qclsinst:qthis}
}

// postDelayedEvent(class QEvent *, int)
func (this *QStateMachine) Postdelayedevent(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QEvent).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN13QStateMachine16postDelayedEventEP6QEventi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStateMachine", "postDelayedEvent", args)
  }

  return
}

// error()
func (this *QStateMachine) Error(args ...interface{}) () {
  // error()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStateMachine5errorEv
    // invoke: QStateMachine::Error error()
    C.C_ZNK13QStateMachine5errorEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStateMachine", "error", args)
  }

  return
}

// removeDefaultAnimation(class QAbstractAnimation *)
func (this *QStateMachine) Removedefaultanimation(args ...interface{}) () {
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
    var arg0 = args[0].(*QAbstractAnimation).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QStateMachine22removeDefaultAnimationEP18QAbstractAnimation(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStateMachine", "removeDefaultAnimation", args)
  }

  return
}

// <= body block end

