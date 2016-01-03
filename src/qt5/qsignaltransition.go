package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
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
  // proto:  void QSignalTransition::setSenderObject(const QObject * sender);
extern void _ZN17QSignalTransition15setSenderObjectEPK7QObject(void* qthis, void* arg0);
  // proto:  QByteArray QSignalTransition::signal();
extern void _ZNK17QSignalTransition6signalEv(void* qthis);
  // proto:  void QSignalTransition::~QSignalTransition();
extern void _ZN17QSignalTransitionD0Ev(void* qthis);
  // proto:  void QSignalTransition::QSignalTransition(const QSignalTransition & );
extern void* dector_ZN17QSignalTransitionC1ERKS_(void* arg0);
extern void _ZN17QSignalTransitionC1ERKS_(void* qthis, void* arg0);
  // proto:  void QSignalTransition::QSignalTransition(const QObject * sender, const char * signal, QState * sourceState);
extern void* dector_ZN17QSignalTransitionC1EPK7QObjectPKcP6QState(void* arg0, unsigned char* arg1, void* arg2);
extern void _ZN17QSignalTransitionC1EPK7QObjectPKcP6QState(void* qthis, void* arg0, unsigned char* arg1, void* arg2);
  // proto:  QObject * QSignalTransition::senderObject();
extern void _ZNK17QSignalTransition12senderObjectEv(void* qthis);
  // proto:  void QSignalTransition::QSignalTransition(QState * sourceState);
extern void* dector_ZN17QSignalTransitionC1EP6QState(void* arg0);
extern void _ZN17QSignalTransitionC1EP6QState(void* qthis, void* arg0);
  // proto:  void QSignalTransition::setSignal(const QByteArray & signal);
extern void _ZN17QSignalTransition9setSignalERK10QByteArray(void* qthis, void* arg0);
  // proto:  const QMetaObject * QSignalTransition::metaObject();
extern void _ZNK17QSignalTransition10metaObjectEv(void* qthis);
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

  // proto:  void QSignalTransition::setSenderObject(const QObject * sender);
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

  // proto:  QByteArray QSignalTransition::signal();
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

  // proto:  void QSignalTransition::~QSignalTransition();
func (this *QSignalTransition) FreeQSignalTransition(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSignalTransition", "~QSignalTransition", args)
  }

}

  // proto:  void QSignalTransition::QSignalTransition(const QSignalTransition & );
func NewQSignalTransition(args ...interface{}) QSignalTransition {
  return QSignalTransition{}
}

  // proto:  QObject * QSignalTransition::senderObject();
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

  // proto:  void QSignalTransition::setSignal(const QByteArray & signal);
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

  // proto:  const QMetaObject * QSignalTransition::metaObject();
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

// <= body block end

