package qtcore
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
// src-file: /QtCore/qeventtransition.h
// dst-file: /src/core/qeventtransition.go
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
  // proto:  const QMetaObject * QEventTransition::metaObject();
extern void C_ZNK16QEventTransition10metaObjectEv(void* qthis); // 4
  // proto:  void QEventTransition::setEventSource(QObject * object);
extern void C_ZN16QEventTransition14setEventSourceEP7QObject(void* qthis, void* arg0); // 4
  // proto:  void QEventTransition::QEventTransition(QState * sourceState);
extern void* C_ZN16QEventTransitionC2EP6QState(void* arg0); // 3
  // proto:  QObject * QEventTransition::eventSource();
extern void* C_ZNK16QEventTransition11eventSourceEv(void* qthis); // 4
  // proto:  QEvent::Type QEventTransition::eventType();
extern void C_ZNK16QEventTransition9eventTypeEv(void* qthis); // 4
  // proto:  void QEventTransition::~QEventTransition();
extern void C_ZN16QEventTransitionD2Ev(void* qthis); // 4
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

// class sizeof(QEventTransition)=1
type QEventTransition struct {
  /*qbase*/ QAbstractTransition;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// metaObject()
func (this *QEventTransition) Metaobject(args ...interface{}) () {
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
    // invoke: _ZNK16QEventTransition10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK16QEventTransition10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QEventTransition", "metaObject", args)
  }

  return
}

// setEventSource(class QObject *)
func (this *QEventTransition) Seteventsource(args ...interface{}) () {
  // setEventSource(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QEventTransition14setEventSourceEP7QObject
    // invoke: void setEventSource(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QEventTransition14setEventSourceEP7QObject(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QEventTransition", "setEventSource", args)
  }

  return
}

// QEventTransition(class QState *)
func NewQEventTransition(args ...interface{}) *QEventTransition {
  // QEventTransition(class QState *)
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
    // invoke: _ZN16QEventTransitionC1EP6QState
    // invoke: void QEventTransition(class QState *)
    var arg0 = args[0].(*QState).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QEventTransitionC2EP6QState(arg0)
    return &QEventTransition{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QEventTransition", "QEventTransition", args)
  }

  return nil // QEventTransition{Qclsinst:qthis}
}

// eventSource()
func (this *QEventTransition) Eventsource(args ...interface{}) (ret interface{}) {
  // eventSource()
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
    // invoke: _ZNK16QEventTransition11eventSourceEv
    // invoke: QObject * eventSource()
    var ret0 = C.C_ZNK16QEventTransition11eventSourceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QObject{}) // "QObject *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QEventTransition", "eventSource", args)
  }

  return
}

// eventType()
func (this *QEventTransition) Eventtype(args ...interface{}) () {
  // eventType()
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
    // invoke: _ZNK16QEventTransition9eventTypeEv
    // invoke: QEvent::Type eventType()
    C.C_ZNK16QEventTransition9eventTypeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QEventTransition", "eventType", args)
  }

  return
}

// ~QEventTransition()
func (this *QEventTransition) Freeqeventtransition(args ...interface{}) () {
  // ~QEventTransition()
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
    // invoke: _ZN16QEventTransitionD0Ev
    // invoke: void ~QEventTransition()
    C.C_ZN16QEventTransitionD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QEventTransition", "~QEventTransition", args)
  }

  return
}

// <= body block end

