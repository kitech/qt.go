package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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
extern void* C_ZN21QObjectCleanupHandler3addEP7QObject(void* qthis, void* arg0); // 4
  // proto:  bool QObjectCleanupHandler::isEmpty();
extern bool C_ZNK21QObjectCleanupHandler7isEmptyEv(void* qthis); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// metaObject()
func (this *QObjectCleanupHandler) Metaobject(args ...interface{}) () {
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
    C.C_ZNK21QObjectCleanupHandler10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QObjectCleanupHandler", "metaObject", args)
  }

  return
}

// clear()
func (this *QObjectCleanupHandler) Clear(args ...interface{}) () {
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
    C.C_ZN21QObjectCleanupHandler5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QObjectCleanupHandler", "clear", args)
  }

  return
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
    return &QObjectCleanupHandler{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QObjectCleanupHandler", "QObjectCleanupHandler", args)
  }

  return nil // QObjectCleanupHandler{Qclsinst:qthis}
}

// add(class QObject *)
func (this *QObjectCleanupHandler) Add(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN21QObjectCleanupHandler3addEP7QObject(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QObject{}) // "QObject *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QObjectCleanupHandler", "add", args)
  }

  return
}

// isEmpty()
func (this *QObjectCleanupHandler) Isempty(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK21QObjectCleanupHandler7isEmptyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QObjectCleanupHandler", "isEmpty", args)
  }

  return
}

// remove(class QObject *)
func (this *QObjectCleanupHandler) Remove(args ...interface{}) () {
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
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN21QObjectCleanupHandler6removeEP7QObject(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QObjectCleanupHandler", "remove", args)
  }

  return
}

// ~QObjectCleanupHandler()
func (this *QObjectCleanupHandler) Freeqobjectcleanuphandler(args ...interface{}) () {
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
    C.C_ZN21QObjectCleanupHandlerD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QObjectCleanupHandler", "~QObjectCleanupHandler", args)
  }

  return
}

// <= body block end

