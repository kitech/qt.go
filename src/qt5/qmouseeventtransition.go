package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtWidgets/qmouseeventtransition.h
// dst-file: /src/widgets/qmouseeventtransition.go
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
  // proto:  const QMetaObject * QMouseEventTransition::metaObject();
extern void _ZNK21QMouseEventTransition10metaObjectEv(void* qthis);
  // proto:  void QMouseEventTransition::QMouseEventTransition(QState * sourceState);
extern void* dector_ZN21QMouseEventTransitionC1EP6QState(void* arg0);
extern void _ZN21QMouseEventTransitionC1EP6QState(void* qthis, void* arg0);
  // proto:  void QMouseEventTransition::QMouseEventTransition(const QMouseEventTransition & );
extern void* dector_ZN21QMouseEventTransitionC1ERKS_(void* arg0);
extern void _ZN21QMouseEventTransitionC1ERKS_(void* qthis, void* arg0);
  // proto:  void QMouseEventTransition::setHitTestPath(const QPainterPath & path);
extern void _ZN21QMouseEventTransition14setHitTestPathERK12QPainterPath(void* qthis, void* arg0);
  // proto:  void QMouseEventTransition::~QMouseEventTransition();
extern void _ZN21QMouseEventTransitionD0Ev(void* qthis);
  // proto:  QPainterPath QMouseEventTransition::hitTestPath();
extern void _ZNK21QMouseEventTransition11hitTestPathEv(void* qthis);
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

// class sizeof(QMouseEventTransition)=1
type QMouseEventTransition struct {
  /*qbase*/ QEventTransition;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  const QMetaObject * QMouseEventTransition::metaObject();
func (this *QMouseEventTransition) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QMouseEventTransition10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK21QMouseEventTransition10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEventTransition", "metaObject", args)
  }

}

  // proto:  void QMouseEventTransition::QMouseEventTransition(QState * sourceState);
func NewQMouseEventTransition(args ...interface{}) QMouseEventTransition {
  return QMouseEventTransition{}
}

  // proto:  void QMouseEventTransition::setHitTestPath(const QPainterPath & path);
func (this *QMouseEventTransition) setHitTestPath(args ...interface{}) () {
  // setHitTestPath(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QMouseEventTransition14setHitTestPathERK12QPainterPath
    // invoke: void setHitTestPath(const class QPainterPath &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN21QMouseEventTransition14setHitTestPathERK12QPainterPath(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMouseEventTransition", "setHitTestPath", args)
  }

}

  // proto:  void QMouseEventTransition::~QMouseEventTransition();
func (this *QMouseEventTransition) FreeQMouseEventTransition(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMouseEventTransition", "~QMouseEventTransition", args)
  }

}

  // proto:  QPainterPath QMouseEventTransition::hitTestPath();
func (this *QMouseEventTransition) hitTestPath(args ...interface{}) () {
  // hitTestPath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QMouseEventTransition11hitTestPathEv
    // invoke: QPainterPath hitTestPath()
    C._ZNK21QMouseEventTransition11hitTestPathEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEventTransition", "hitTestPath", args)
  }

}

// <= body block end

