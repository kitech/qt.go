package qtcore
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtCore/qabstracteventdispatcher.h
// dst-file: /src/core/qabstracteventdispatcher.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "runtime"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QAbstractEventDispatcher::installNativeEventFilter(QAbstractNativeEventFilter * filterObj);
extern void C_ZN24QAbstractEventDispatcher24installNativeEventFilterEP26QAbstractNativeEventFilter(void* qthis, void* arg0); // 4
  // proto:  void QAbstractEventDispatcher::~QAbstractEventDispatcher();
extern void C_ZN24QAbstractEventDispatcherD2Ev(void* qthis); // 4
  // proto:  void QAbstractEventDispatcher::closingDown();
extern void C_ZN24QAbstractEventDispatcher11closingDownEv(void* qthis); // 4
  // proto:  void QAbstractEventDispatcher::startingUp();
extern void C_ZN24QAbstractEventDispatcher10startingUpEv(void* qthis); // 4
  // proto: static QAbstractEventDispatcher * QAbstractEventDispatcher::instance(QThread * thread);
extern void C_ZN24QAbstractEventDispatcher8instanceEP7QThread(void* arg0); // 4
  // proto:  void QAbstractEventDispatcher::removeNativeEventFilter(QAbstractNativeEventFilter * filterObj);
extern void C_ZN24QAbstractEventDispatcher23removeNativeEventFilterEP26QAbstractNativeEventFilter(void* qthis, void* arg0); // 4
  // proto:  bool QAbstractEventDispatcher::filterNativeEvent(const QByteArray & eventType, void * message, long * result);
extern bool C_ZN24QAbstractEventDispatcher17filterNativeEventERK10QByteArrayPvPl(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QAbstractEventDispatcher::QAbstractEventDispatcher(QObject * parent);
extern void* C_ZN24QAbstractEventDispatcherC2EP7QObject(void* arg0); // 3
  // proto:  const QMetaObject * QAbstractEventDispatcher::metaObject();
extern void C_ZNK24QAbstractEventDispatcher10metaObjectEv(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QAbstractEventDispatcher)=1
type QAbstractEventDispatcher struct {
  /*qbase*/ QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _aboutToBlock QAbstractEventDispatcher_aboutToBlock_signal;
//  _awake QAbstractEventDispatcher_awake_signal;
}

// installNativeEventFilter(class QAbstractNativeEventFilter *)
func (this *QAbstractEventDispatcher) InstallNativeEventFilter(args ...interface{}) () {
  // installNativeEventFilter(class QAbstractNativeEventFilter *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractNativeEventFilter{}) // "QAbstractNativeEventFilter *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractEventDispatcher24installNativeEventFilterEP26QAbstractNativeEventFilter
    // invoke: void installNativeEventFilter(class QAbstractNativeEventFilter *)
    var arg0 = args[0].(*QAbstractNativeEventFilter).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN24QAbstractEventDispatcher24installNativeEventFilterEP26QAbstractNativeEventFilter(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "installNativeEventFilter", args)
  }

  return
}

// ~QAbstractEventDispatcher()
func (this *QAbstractEventDispatcher) Free(args ...interface{}) () {
  // ~QAbstractEventDispatcher()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractEventDispatcherD0Ev
    // invoke: void ~QAbstractEventDispatcher()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN24QAbstractEventDispatcherD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "~QAbstractEventDispatcher", args)
  }

  return
}

// closingDown()
func (this *QAbstractEventDispatcher) ClosingDown(args ...interface{}) () {
  // closingDown()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractEventDispatcher11closingDownEv
    // invoke: void closingDown()
    C.C_ZN24QAbstractEventDispatcher11closingDownEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "closingDown", args)
  }

  return
}

// startingUp()
func (this *QAbstractEventDispatcher) StartingUp(args ...interface{}) () {
  // startingUp()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractEventDispatcher10startingUpEv
    // invoke: void startingUp()
    C.C_ZN24QAbstractEventDispatcher10startingUpEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "startingUp", args)
  }

  return
}

// instance(class QThread *)
func (this *QAbstractEventDispatcher) Instance_s(args ...interface{}) () {
  // instance(class QThread *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QThread{}) // "QThread *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractEventDispatcher8instanceEP7QThread
    // invoke: QAbstractEventDispatcher * instance(class QThread *)
    var arg0 = args[0].(*QThread).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN24QAbstractEventDispatcher8instanceEP7QThread(arg0)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "instance", args)
  }

  return
}

// removeNativeEventFilter(class QAbstractNativeEventFilter *)
func (this *QAbstractEventDispatcher) RemoveNativeEventFilter(args ...interface{}) () {
  // removeNativeEventFilter(class QAbstractNativeEventFilter *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractNativeEventFilter{}) // "QAbstractNativeEventFilter *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractEventDispatcher23removeNativeEventFilterEP26QAbstractNativeEventFilter
    // invoke: void removeNativeEventFilter(class QAbstractNativeEventFilter *)
    var arg0 = args[0].(*QAbstractNativeEventFilter).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN24QAbstractEventDispatcher23removeNativeEventFilterEP26QAbstractNativeEventFilter(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "removeNativeEventFilter", args)
  }

  return
}

// filterNativeEvent(const class QByteArray &, void *, long *)
func (this *QAbstractEventDispatcher) FilterNativeEvent(args ...interface{}) (ret interface{}) {
  // filterNativeEvent(const class QByteArray &, void *, long *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[0][1] = qtrt.VoidpTy() // "void *"
  vtys[0][2] = qtrt.Int32Ty(true) // "long *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractEventDispatcher17filterNativeEventERK10QByteArrayPvPl
    // invoke: bool filterNativeEvent(const class QByteArray &, void *, long *)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN24QAbstractEventDispatcher17filterNativeEventERK10QByteArrayPvPl(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "filterNativeEvent", args)
  }

  return
}

// QAbstractEventDispatcher(class QObject *)
func GcfreeQAbstractEventDispatcher(this *QAbstractEventDispatcher) {
  qtrt.UniverseFree(this)
}
func NewQAbstractEventDispatcher(args ...interface{}) *QAbstractEventDispatcher {
  // QAbstractEventDispatcher(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractEventDispatcherC1EP7QObject
    // invoke: void QAbstractEventDispatcher(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN24QAbstractEventDispatcherC2EP7QObject(arg0)
    this := &QAbstractEventDispatcher{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQAbstractEventDispatcher)
    return this
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "QAbstractEventDispatcher", args)
  }

  return nil // QAbstractEventDispatcher{Qclsinst:qthis}
}

// metaObject()
func (this *QAbstractEventDispatcher) MetaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QAbstractEventDispatcher10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK24QAbstractEventDispatcher10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "metaObject", args)
  }

  return
}

// <= body block end

