package qtcore
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtCore/qeventloop.h
// dst-file: /src/core/qeventloop.go
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
  // proto:  void QEventLoop::quit();
extern void C_ZN10QEventLoop4quitEv(void* qthis); // 4
  // proto:  bool QEventLoop::isRunning();
extern bool C_ZNK10QEventLoop9isRunningEv(void* qthis); // 4
  // proto:  void QEventLoop::QEventLoop(QObject * parent);
extern void* C_ZN10QEventLoopC2EP7QObject(void* arg0); // 3
  // proto:  void QEventLoop::exit(int returnCode);
extern void C_ZN10QEventLoop4exitEi(void* qthis, int32_t arg0); // 4
  // proto:  void QEventLoop::wakeUp();
extern void C_ZN10QEventLoop6wakeUpEv(void* qthis); // 4
  // proto:  void QEventLoop::~QEventLoop();
extern void C_ZN10QEventLoopD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QEventLoop::metaObject();
extern void C_ZNK10QEventLoop10metaObjectEv(void* qthis); // 4
  // proto:  bool QEventLoop::event(QEvent * event);
extern bool C_ZN10QEventLoop5eventEP6QEvent(void* qthis, void* arg0); // 4
  // proto:  void QEventLoopLocker::~QEventLoopLocker();
extern void C_ZN16QEventLoopLockerD2Ev(void* qthis); // 4
  // proto:  void QEventLoopLocker::QEventLoopLocker();
extern void* C_ZN16QEventLoopLockerC2Ev(); // 3
  // proto:  void QEventLoopLocker::QEventLoopLocker(QThread * thread);
extern void* C_ZN16QEventLoopLockerC2EP7QThread(void* arg0); // 3
  // proto:  void QEventLoopLocker::QEventLoopLocker(QEventLoop * loop);
extern void* C_ZN16QEventLoopLockerC2EP10QEventLoop(void* arg0); // 3
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

// class sizeof(QEventLoop)=1
type QEventLoop struct {
  /*qbase*/ QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QEventLoopLocker)=8
type QEventLoopLocker struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// quit()
func (this *QEventLoop) Quit(args ...interface{}) () {
  // quit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QEventLoop4quitEv
    // invoke: void quit()
    C.C_ZN10QEventLoop4quitEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QEventLoop", "quit", args)
  }

  return
}

// isRunning()
func (this *QEventLoop) IsRunning(args ...interface{}) (ret interface{}) {
  // isRunning()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QEventLoop9isRunningEv
    // invoke: bool isRunning()
    var ret0 = C.C_ZNK10QEventLoop9isRunningEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QEventLoop", "isRunning", args)
  }

  return
}

// QEventLoop(class QObject *)
func GcfreeQEventLoop(this *QEventLoop) {
  qtrt.UniverseFree(this)
}
func NewQEventLoop(args ...interface{}) *QEventLoop {
  // QEventLoop(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QEventLoopC1EP7QObject
    // invoke: void QEventLoop(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QEventLoopC2EP7QObject(arg0)
    this := &QEventLoop{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQEventLoop)
    return this
  default:
    qtrt.ErrorResolve("QEventLoop", "QEventLoop", args)
  }

  return nil // QEventLoop{Qclsinst:qthis}
}

// exit(int)
func (this *QEventLoop) Exit(args ...interface{}) () {
  // exit(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QEventLoop4exitEi
    // invoke: void exit(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QEventLoop4exitEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QEventLoop", "exit", args)
  }

  return
}

// wakeUp()
func (this *QEventLoop) WakeUp(args ...interface{}) () {
  // wakeUp()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QEventLoop6wakeUpEv
    // invoke: void wakeUp()
    C.C_ZN10QEventLoop6wakeUpEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QEventLoop", "wakeUp", args)
  }

  return
}

// ~QEventLoop()
func (this *QEventLoop) Free(args ...interface{}) () {
  // ~QEventLoop()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QEventLoopD0Ev
    // invoke: void ~QEventLoop()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN10QEventLoopD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QEventLoop", "~QEventLoop", args)
  }

  return
}

// metaObject()
func (this *QEventLoop) MetaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QEventLoop10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK10QEventLoop10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QEventLoop", "metaObject", args)
  }

  return
}

// event(class QEvent *)
func (this *QEventLoop) Event(args ...interface{}) (ret interface{}) {
  // event(class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QEventLoop5eventEP6QEvent
    // invoke: bool event(class QEvent *)
    var arg0 = args[0].(*QEvent).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN10QEventLoop5eventEP6QEvent(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QEventLoop", "event", args)
  }

  return
}

// ~QEventLoopLocker()
func (this *QEventLoopLocker) Free(args ...interface{}) () {
  // ~QEventLoopLocker()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QEventLoopLockerD0Ev
    // invoke: void ~QEventLoopLocker()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN16QEventLoopLockerD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QEventLoopLocker", "~QEventLoopLocker", args)
  }

  return
}

// QEventLoopLocker()
func GcfreeQEventLoopLocker(this *QEventLoopLocker) {
  qtrt.UniverseFree(this)
}
func NewQEventLoopLocker(args ...interface{}) *QEventLoopLocker {
  // QEventLoopLocker()
  // QEventLoopLocker(class QThread *)
  // QEventLoopLocker(class QEventLoop *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QThread{}) // "QThread *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QEventLoop{}) // "QEventLoop *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QEventLoopLockerC1Ev
    // invoke: void QEventLoopLocker()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QEventLoopLockerC2Ev()
    this := &QEventLoopLocker{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQEventLoopLocker)
    return this
  case 1:
    // invoke: _ZN16QEventLoopLockerC1EP7QThread
    // invoke: void QEventLoopLocker(class QThread *)
    var arg0 = args[0].(*QThread).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QEventLoopLockerC2EP7QThread(arg0)
    this := &QEventLoopLocker{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQEventLoopLocker)
    return this
  case 2:
    // invoke: _ZN16QEventLoopLockerC1EP10QEventLoop
    // invoke: void QEventLoopLocker(class QEventLoop *)
    var arg0 = args[0].(*QEventLoop).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QEventLoopLockerC2EP10QEventLoop(arg0)
    this := &QEventLoopLocker{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQEventLoopLocker)
    return this
  default:
    qtrt.ErrorResolve("QEventLoopLocker", "QEventLoopLocker", args)
  }

  return nil // QEventLoopLocker{Qclsinst:qthis}
}

// <= body block end

