package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtCore/qsignaltransition.h
// dst-file: /src/core/qsignaltransition.go
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
  // proto:  QByteArray QSignalTransition::signal();
extern void _ZNK17QSignalTransition6signalEv(void* qthis); // 4
  // proto:  void QSignalTransition::~QSignalTransition();
extern void _ZN17QSignalTransitionD2Ev(void* qthis); // 4
  // proto:  void QSignalTransition::QSignalTransition(QState * sourceState);
extern void _ZN17QSignalTransitionC2EP6QState(void* qthis, void* arg0); // 3
  // proto:  void QSignalTransition::QSignalTransition(const QObject * sender, const char * signal, QState * sourceState);
extern void _ZN17QSignalTransitionC2EPK7QObjectPKcP6QState(void* qthis, void* arg0, unsigned char* arg1, void* arg2); // 3
  // proto:  const QMetaObject * QSignalTransition::metaObject();
extern void _ZNK17QSignalTransition10metaObjectEv(void* qthis); // 4
  // proto:  void QSignalTransition::setSenderObject(const QObject * sender);
extern void _ZN17QSignalTransition15setSenderObjectEPK7QObject(void* qthis, void* arg0); // 4
  // proto:  QObject * QSignalTransition::senderObject();
extern void _ZNK17QSignalTransition12senderObjectEv(void* qthis); // 4
  // proto:  void QSignalTransition::setSignal(const QByteArray & signal);
extern void _ZN17QSignalTransition9setSignalERK10QByteArray(void* qthis, void* arg0); // 4
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

// class sizeof(QSignalTransition)=1
type QSignalTransition struct {
  /*qbase*/ QAbstractTransition;
  qclsinst unsafe.Pointer /* *C.void */;
//  _senderObjectChanged QSignalTransition_senderObjectChanged_signal;
//  _signalChanged QSignalTransition_signalChanged_signal;
}

// signal()
func (this *QSignalTransition) signal(args ...interface{}) () {
  // signal()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QSignalTransition6signalEv
    // invoke: QByteArray signal()
    C._ZNK17QSignalTransition6signalEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSignalTransition", "signal", args)
  }

}

// ~QSignalTransition()
func (this *QSignalTransition) FreeQSignalTransition(args ...interface{}) () {
  // ~QSignalTransition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QSignalTransitionD0Ev
    // invoke: void ~QSignalTransition()
    C._ZN17QSignalTransitionD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSignalTransition", "~QSignalTransition", args)
  }

}

// QSignalTransition(class QState *)
func NewQSignalTransition(args ...interface{}) QSignalTransition {
  // QSignalTransition(class QState *)
  // QSignalTransition(const class QObject *, const char *, class QState *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QState{}) // "QState *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[1][1] = qtrt.ByteTy(true) // "const char *"
  vtys[1][2] = reflect.TypeOf(QState{}) // "QState *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QSignalTransitionC1EP6QState
    // invoke: void QSignalTransition(class QState *)
    var arg0 = args[0].(QState).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN17QSignalTransitionC2EP6QState(qthis, arg0)
  case 1:
    // invoke: _ZN17QSignalTransitionC1EPK7QObjectPKcP6QState
    // invoke: void QSignalTransition(const class QObject *, const char *, class QState *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QState).qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN17QSignalTransitionC2EPK7QObjectPKcP6QState(qthis, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QSignalTransition", "QSignalTransition", args)
  }

  return QSignalTransition{}
}

// metaObject()
func (this *QSignalTransition) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QSignalTransition10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK17QSignalTransition10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSignalTransition", "metaObject", args)
  }

}

// setSenderObject(const class QObject *)
func (this *QSignalTransition) setSenderObject(args ...interface{}) () {
  // setSenderObject(const class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "const QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QSignalTransition15setSenderObjectEPK7QObject
    // invoke: void setSenderObject(const class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN17QSignalTransition15setSenderObjectEPK7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSignalTransition", "setSenderObject", args)
  }

}

// senderObject()
func (this *QSignalTransition) senderObject(args ...interface{}) () {
  // senderObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QSignalTransition12senderObjectEv
    // invoke: QObject * senderObject()
    C._ZNK17QSignalTransition12senderObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSignalTransition", "senderObject", args)
  }

}

// setSignal(const class QByteArray &)
func (this *QSignalTransition) setSignal(args ...interface{}) () {
  // setSignal(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QSignalTransition9setSignalERK10QByteArray
    // invoke: void setSignal(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN17QSignalTransition9setSignalERK10QByteArray(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSignalTransition", "setSignal", args)
  }

}

// <= body block end

