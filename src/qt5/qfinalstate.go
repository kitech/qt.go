package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtCore/qfinalstate.h
// dst-file: /src/core/qfinalstate.go
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
  // proto:  void QFinalState::QFinalState(QState * parent);
extern void* dector_ZN11QFinalStateC1EP6QState(void* arg0);
extern void _ZN11QFinalStateC1EP6QState(void* qthis, void* arg0);
  // proto:  void QFinalState::QFinalState(const QFinalState & );
extern void* dector_ZN11QFinalStateC1ERKS_(void* arg0);
extern void _ZN11QFinalStateC1ERKS_(void* qthis, void* arg0);
  // proto:  void QFinalState::~QFinalState();
extern void _ZN11QFinalStateD0Ev(void* qthis);
  // proto:  const QMetaObject * QFinalState::metaObject();
extern void _ZNK11QFinalState10metaObjectEv(void* qthis);
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

// class sizeof(QFinalState)=1
type QFinalState struct {
  /*qbase*/ QAbstractState;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QFinalState::QFinalState(QState * parent);
func NewQFinalState(args ...interface{}) QFinalState {
  return QFinalState{}
}

  // proto:  void QFinalState::~QFinalState();
func (this *QFinalState) FreeQFinalState(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFinalState", "~QFinalState", args)
  }

}

  // proto:  const QMetaObject * QFinalState::metaObject();
func (this *QFinalState) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFinalState10metaObjectEv
  default:
    qtrt.ErrorResolve("QFinalState", "metaObject", args)
  }

}

// <= body block end

