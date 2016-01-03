package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
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
  // proto:  void QEventTransition::QEventTransition(const QEventTransition & );
extern void* dector_ZN16QEventTransitionC1ERKS_(void* arg0);
extern void _ZN16QEventTransitionC1ERKS_(void* qthis, void* arg0);
  // proto:  void QEventTransition::~QEventTransition();
extern void _ZN16QEventTransitionD0Ev(void* qthis);
  // proto:  void QEventTransition::setEventSource(QObject * object);
extern void _ZN16QEventTransition14setEventSourceEP7QObject(void* qthis, void* arg0);
  // proto:  void QEventTransition::QEventTransition(QState * sourceState);
extern void* dector_ZN16QEventTransitionC1EP6QState(void* arg0);
extern void _ZN16QEventTransitionC1EP6QState(void* qthis, void* arg0);
  // proto:  const QMetaObject * QEventTransition::metaObject();
extern void _ZNK16QEventTransition10metaObjectEv(void* qthis);
  // proto:  QObject * QEventTransition::eventSource();
extern void _ZNK16QEventTransition11eventSourceEv(void* qthis);
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
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  void QEventTransition::QEventTransition(const QEventTransition & );
func NewQEventTransition(args ...interface{}) QEventTransition {
  return QEventTransition{}
}

  // proto:  void QEventTransition::~QEventTransition();
func (this *QEventTransition) FreeQEventTransition(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QEventTransition", "~QEventTransition", args)
  }

}

  // proto:  void QEventTransition::setEventSource(QObject * object);
func (this *QEventTransition) setEventSource(args ...interface{}) () {
  // setEventSource(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QEventTransition14setEventSourceEP7QObject
    // invoke: void setEventSource(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN16QEventTransition14setEventSourceEP7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QEventTransition", "setEventSource", args)
  }

}

  // proto:  const QMetaObject * QEventTransition::metaObject();
func (this *QEventTransition) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QEventTransition10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK16QEventTransition10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEventTransition", "metaObject", args)
  }

}

  // proto:  QObject * QEventTransition::eventSource();
func (this *QEventTransition) eventSource(args ...interface{}) () {
  // eventSource()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QEventTransition11eventSourceEv
    // invoke: QObject * eventSource()
    C._ZNK16QEventTransition11eventSourceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEventTransition", "eventSource", args)
  }

}

// <= body block end

