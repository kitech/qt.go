package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
// src-file: /QtCore/qobjectcleanuphandler.h
// dst-file: /src/core/qobjectcleanuphandler.go
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
  // proto:  void QObjectCleanupHandler::clear();
extern void _ZN21QObjectCleanupHandler5clearEv(void* qthis);
  // proto:  bool QObjectCleanupHandler::isEmpty();
extern void _ZNK21QObjectCleanupHandler7isEmptyEv(void* qthis);
  // proto:  void QObjectCleanupHandler::~QObjectCleanupHandler();
extern void _ZN21QObjectCleanupHandlerD0Ev(void* qthis);
  // proto:  const QMetaObject * QObjectCleanupHandler::metaObject();
extern void _ZNK21QObjectCleanupHandler10metaObjectEv(void* qthis);
  // proto:  void QObjectCleanupHandler::remove(QObject * object);
extern void _ZN21QObjectCleanupHandler6removeEP7QObject(void* qthis, void* arg0);
  // proto:  QObject * QObjectCleanupHandler::add(QObject * object);
extern void _ZN21QObjectCleanupHandler3addEP7QObject(void* qthis, void* arg0);
  // proto:  void QObjectCleanupHandler::QObjectCleanupHandler();
extern void* dector_ZN21QObjectCleanupHandlerC1Ev();
extern void _ZN21QObjectCleanupHandlerC1Ev(void* qthis);
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

// class sizeof(QObjectCleanupHandler)=1
type QObjectCleanupHandler struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  void QObjectCleanupHandler::clear();
func (this *QObjectCleanupHandler) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QObjectCleanupHandler5clearEv
    // invoke: void clear()
    C._ZN21QObjectCleanupHandler5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObjectCleanupHandler", "clear", args)
  }

}

  // proto:  bool QObjectCleanupHandler::isEmpty();
func (this *QObjectCleanupHandler) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QObjectCleanupHandler7isEmptyEv
    // invoke: bool isEmpty()
    C._ZNK21QObjectCleanupHandler7isEmptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObjectCleanupHandler", "isEmpty", args)
  }

}

  // proto:  void QObjectCleanupHandler::~QObjectCleanupHandler();
func (this *QObjectCleanupHandler) FreeQObjectCleanupHandler(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QObjectCleanupHandler", "~QObjectCleanupHandler", args)
  }

}

  // proto:  const QMetaObject * QObjectCleanupHandler::metaObject();
func (this *QObjectCleanupHandler) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QObjectCleanupHandler10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK21QObjectCleanupHandler10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObjectCleanupHandler", "metaObject", args)
  }

}

  // proto:  void QObjectCleanupHandler::remove(QObject * object);
func (this *QObjectCleanupHandler) remove(args ...interface{}) () {
  // remove(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QObjectCleanupHandler6removeEP7QObject
    // invoke: void remove(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN21QObjectCleanupHandler6removeEP7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QObjectCleanupHandler", "remove", args)
  }

}

  // proto:  QObject * QObjectCleanupHandler::add(QObject * object);
func (this *QObjectCleanupHandler) add(args ...interface{}) () {
  // add(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QObjectCleanupHandler3addEP7QObject
    // invoke: QObject * add(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN21QObjectCleanupHandler3addEP7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QObjectCleanupHandler", "add", args)
  }

}

  // proto:  void QObjectCleanupHandler::QObjectCleanupHandler();
func NewQObjectCleanupHandler(args ...interface{}) QObjectCleanupHandler {
  return QObjectCleanupHandler{}
}

// <= body block end

