package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtCore/qthreadpool.h
// dst-file: /src/core/qthreadpool.go
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
  // proto:  void QThreadPool::setExpiryTimeout(int expiryTimeout);
extern void C_ZN11QThreadPool16setExpiryTimeoutEi(void* qthis, int32_t arg0); // 4
  // proto:  void QThreadPool::reserveThread();
extern void C_ZN11QThreadPool13reserveThreadEv(void* qthis); // 4
  // proto:  void QThreadPool::cancel(QRunnable * runnable);
extern void C_ZN11QThreadPool6cancelEP9QRunnable(void* qthis, void* arg0); // 4
  // proto:  void QThreadPool::QThreadPool(QObject * parent);
extern void C_ZN11QThreadPoolC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  bool QThreadPool::waitForDone(int msecs);
extern void C_ZN11QThreadPool11waitForDoneEi(void* qthis, int32_t arg0); // 4
  // proto:  void QThreadPool::start(QRunnable * runnable, int priority);
extern void C_ZN11QThreadPool5startEP9QRunnablei(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  int QThreadPool::activeThreadCount();
extern void C_ZNK11QThreadPool17activeThreadCountEv(void* qthis); // 4
  // proto:  void QThreadPool::setMaxThreadCount(int maxThreadCount);
extern void C_ZN11QThreadPool17setMaxThreadCountEi(void* qthis, int32_t arg0); // 4
  // proto:  void QThreadPool::~QThreadPool();
extern void C_ZN11QThreadPoolD2Ev(void* qthis); // 4
  // proto: static QThreadPool * QThreadPool::globalInstance();
extern void C_ZN11QThreadPool14globalInstanceEv(); // 4
  // proto:  const QMetaObject * QThreadPool::metaObject();
extern void C_ZNK11QThreadPool10metaObjectEv(void* qthis); // 4
  // proto:  void QThreadPool::clear();
extern void C_ZN11QThreadPool5clearEv(void* qthis); // 4
  // proto:  void QThreadPool::releaseThread();
extern void C_ZN11QThreadPool13releaseThreadEv(void* qthis); // 4
  // proto:  bool QThreadPool::tryStart(QRunnable * runnable);
extern void C_ZN11QThreadPool8tryStartEP9QRunnable(void* qthis, void* arg0); // 4
  // proto:  int QThreadPool::maxThreadCount();
extern void C_ZNK11QThreadPool14maxThreadCountEv(void* qthis); // 4
  // proto:  int QThreadPool::expiryTimeout();
extern void C_ZNK11QThreadPool13expiryTimeoutEv(void* qthis); // 4
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

// class sizeof(QThreadPool)=1
type QThreadPool struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
}

// setExpiryTimeout(int)
func (this *QThreadPool) setExpiryTimeout(args ...interface{}) () {
  // setExpiryTimeout(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QThreadPool16setExpiryTimeoutEi
    // invoke: void setExpiryTimeout(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QThreadPool16setExpiryTimeoutEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QThreadPool", "setExpiryTimeout", args)
  }

}

// reserveThread()
func (this *QThreadPool) reserveThread(args ...interface{}) () {
  // reserveThread()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QThreadPool13reserveThreadEv
    // invoke: void reserveThread()
    C.C_ZN11QThreadPool13reserveThreadEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThreadPool", "reserveThread", args)
  }

}

// cancel(class QRunnable *)
func (this *QThreadPool) cancel(args ...interface{}) () {
  // cancel(class QRunnable *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRunnable{}) // "QRunnable *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QThreadPool6cancelEP9QRunnable
    // invoke: void cancel(class QRunnable *)
    var arg0 = args[0].(QRunnable).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QThreadPool6cancelEP9QRunnable(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QThreadPool", "cancel", args)
  }

}

// QThreadPool(class QObject *)
func NewQThreadPool(args ...interface{}) QThreadPool {
  // QThreadPool(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QThreadPoolC1EP7QObject
    // invoke: void QThreadPool(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QThreadPoolC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QThreadPool", "QThreadPool", args)
  }

  return QThreadPool{}
}

// waitForDone(int)
func (this *QThreadPool) waitForDone(args ...interface{}) () {
  // waitForDone(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QThreadPool11waitForDoneEi
    // invoke: bool waitForDone(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QThreadPool11waitForDoneEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QThreadPool", "waitForDone", args)
  }

}

// start(class QRunnable *, int)
func (this *QThreadPool) start(args ...interface{}) () {
  // start(class QRunnable *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRunnable{}) // "QRunnable *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QThreadPool5startEP9QRunnablei
    // invoke: void start(class QRunnable *, int)
    var arg0 = args[0].(QRunnable).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN11QThreadPool5startEP9QRunnablei(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QThreadPool", "start", args)
  }

}

// activeThreadCount()
func (this *QThreadPool) activeThreadCount(args ...interface{}) () {
  // activeThreadCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QThreadPool17activeThreadCountEv
    // invoke: int activeThreadCount()
    C.C_ZNK11QThreadPool17activeThreadCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThreadPool", "activeThreadCount", args)
  }

}

// setMaxThreadCount(int)
func (this *QThreadPool) setMaxThreadCount(args ...interface{}) () {
  // setMaxThreadCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QThreadPool17setMaxThreadCountEi
    // invoke: void setMaxThreadCount(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QThreadPool17setMaxThreadCountEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QThreadPool", "setMaxThreadCount", args)
  }

}

// ~QThreadPool()
func (this *QThreadPool) FreeQThreadPool(args ...interface{}) () {
  // ~QThreadPool()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QThreadPoolD0Ev
    // invoke: void ~QThreadPool()
    C.C_ZN11QThreadPoolD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThreadPool", "~QThreadPool", args)
  }

}

// globalInstance()
func (this *QThreadPool) globalInstance_s(args ...interface{}) () {
  // globalInstance()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QThreadPool14globalInstanceEv
    // invoke: QThreadPool * globalInstance()
    C.C_ZN11QThreadPool14globalInstanceEv()
  default:
    qtrt.ErrorResolve("QThreadPool", "globalInstance", args)
  }

}

// metaObject()
func (this *QThreadPool) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QThreadPool10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK11QThreadPool10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThreadPool", "metaObject", args)
  }

}

// clear()
func (this *QThreadPool) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QThreadPool5clearEv
    // invoke: void clear()
    C.C_ZN11QThreadPool5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThreadPool", "clear", args)
  }

}

// releaseThread()
func (this *QThreadPool) releaseThread(args ...interface{}) () {
  // releaseThread()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QThreadPool13releaseThreadEv
    // invoke: void releaseThread()
    C.C_ZN11QThreadPool13releaseThreadEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThreadPool", "releaseThread", args)
  }

}

// tryStart(class QRunnable *)
func (this *QThreadPool) tryStart(args ...interface{}) () {
  // tryStart(class QRunnable *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRunnable{}) // "QRunnable *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QThreadPool8tryStartEP9QRunnable
    // invoke: bool tryStart(class QRunnable *)
    var arg0 = args[0].(QRunnable).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QThreadPool8tryStartEP9QRunnable(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QThreadPool", "tryStart", args)
  }

}

// maxThreadCount()
func (this *QThreadPool) maxThreadCount(args ...interface{}) () {
  // maxThreadCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QThreadPool14maxThreadCountEv
    // invoke: int maxThreadCount()
    C.C_ZNK11QThreadPool14maxThreadCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThreadPool", "maxThreadCount", args)
  }

}

// expiryTimeout()
func (this *QThreadPool) expiryTimeout(args ...interface{}) () {
  // expiryTimeout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QThreadPool13expiryTimeoutEv
    // invoke: int expiryTimeout()
    C.C_ZNK11QThreadPool13expiryTimeoutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThreadPool", "expiryTimeout", args)
  }

}

// <= body block end

