package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  const QMetaObject * QObjectCleanupHandler::metaObject();
extern void C_ZNK21QObjectCleanupHandler10metaObjectEv(void* qthis); // 4
  // proto:  void QObjectCleanupHandler::clear();
extern void C_ZN21QObjectCleanupHandler5clearEv(void* qthis); // 4
  // proto:  void QObjectCleanupHandler::QObjectCleanupHandler();
extern void* C_ZN21QObjectCleanupHandlerC2Ev(); // 3
  // proto:  QObject * QObjectCleanupHandler::add(QObject * object);
extern void C_ZN21QObjectCleanupHandler3addEP7QObject(void* qthis, void* arg0); // 4
  // proto:  bool QObjectCleanupHandler::isEmpty();
extern void C_ZNK21QObjectCleanupHandler7isEmptyEv(void* qthis); // 4
  // proto:  void QObjectCleanupHandler::remove(QObject * object);
extern void C_ZN21QObjectCleanupHandler6removeEP7QObject(void* qthis, void* arg0); // 4
  // proto:  void QObjectCleanupHandler::~QObjectCleanupHandler();
extern void C_ZN21QObjectCleanupHandlerD2Ev(void* qthis); // 4
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

// metaObject()
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
    C.C_ZNK21QObjectCleanupHandler10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObjectCleanupHandler", "metaObject", args)
  }

}

// clear()
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
    C.C_ZN21QObjectCleanupHandler5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObjectCleanupHandler", "clear", args)
  }

}

// QObjectCleanupHandler()
func NewQObjectCleanupHandler(args ...interface{}) *QObjectCleanupHandler {
  // QObjectCleanupHandler()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QObjectCleanupHandlerC1Ev
    // invoke: void QObjectCleanupHandler()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN21QObjectCleanupHandlerC2Ev()
    return &QObjectCleanupHandler{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QObjectCleanupHandler", "QObjectCleanupHandler", args)
  }

  return nil // QObjectCleanupHandler{qclsinst:qthis}
}

// add(class QObject *)
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
    var ret = C.C_ZN21QObjectCleanupHandler3addEP7QObject(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QObjectCleanupHandler", "add", args)
  }

}

// isEmpty()
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
    var ret = C.C_ZNK21QObjectCleanupHandler7isEmptyEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QObjectCleanupHandler", "isEmpty", args)
  }

}

// remove(class QObject *)
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
    C.C_ZN21QObjectCleanupHandler6removeEP7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QObjectCleanupHandler", "remove", args)
  }

}

// ~QObjectCleanupHandler()
func (this *QObjectCleanupHandler) FreeQObjectCleanupHandler(args ...interface{}) () {
  // ~QObjectCleanupHandler()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QObjectCleanupHandlerD0Ev
    // invoke: void ~QObjectCleanupHandler()
    C.C_ZN21QObjectCleanupHandlerD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObjectCleanupHandler", "~QObjectCleanupHandler", args)
  }

}

// <= body block end

