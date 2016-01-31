package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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
extern void C_ZNK10QEventLoop9isRunningEv(void* qthis); // 4
  // proto:  void QEventLoop::QEventLoop(QObject * parent);
extern void C_ZN10QEventLoopC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  void QEventLoop::exit(int returnCode);
extern void C_ZN10QEventLoop4exitEi(void* qthis, int32_t arg0); // 4
  // proto:  void QEventLoop::wakeUp();
extern void C_ZN10QEventLoop6wakeUpEv(void* qthis); // 4
  // proto:  void QEventLoop::~QEventLoop();
extern void C_ZN10QEventLoopD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QEventLoop::metaObject();
extern void C_ZNK10QEventLoop10metaObjectEv(void* qthis); // 4
  // proto:  bool QEventLoop::event(QEvent * event);
extern void C_ZN10QEventLoop5eventEP6QEvent(void* qthis, void* arg0); // 4
  // proto:  void QEventLoopLocker::~QEventLoopLocker();
extern void C_ZN16QEventLoopLockerD2Ev(void* qthis); // 4
  // proto:  void QEventLoopLocker::QEventLoopLocker();
extern void C_ZN16QEventLoopLockerC2Ev(void* qthis); // 3
  // proto:  void QEventLoopLocker::QEventLoopLocker(QThread * thread);
extern void C_ZN16QEventLoopLockerC2EP7QThread(void* qthis, void* arg0); // 3
  // proto:  void QEventLoopLocker::QEventLoopLocker(QEventLoop * loop);
extern void C_ZN16QEventLoopLockerC2EP10QEventLoop(void* qthis, void* arg0); // 3
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

// class sizeof(QEventLoop)=1
type QEventLoop struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QEventLoopLocker)=8
type QEventLoopLocker struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// quit()
func (this *QEventLoop) quit(args ...interface{}) () {
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
    C.C_ZN10QEventLoop4quitEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEventLoop", "quit", args)
  }

}

// isRunning()
func (this *QEventLoop) isRunning(args ...interface{}) () {
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
    C.C_ZNK10QEventLoop9isRunningEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEventLoop", "isRunning", args)
  }

}

// QEventLoop(class QObject *)
func NewQEventLoop(args ...interface{}) QEventLoop {
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
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN10QEventLoopC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QEventLoop", "QEventLoop", args)
  }

  return QEventLoop{}
}

// exit(int)
func (this *QEventLoop) exit(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QEventLoop4exitEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QEventLoop", "exit", args)
  }

}

// wakeUp()
func (this *QEventLoop) wakeUp(args ...interface{}) () {
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
    C.C_ZN10QEventLoop6wakeUpEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEventLoop", "wakeUp", args)
  }

}

// ~QEventLoop()
func (this *QEventLoop) FreeQEventLoop(args ...interface{}) () {
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
    C.C_ZN10QEventLoopD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEventLoop", "~QEventLoop", args)
  }

}

// metaObject()
func (this *QEventLoop) metaObject(args ...interface{}) () {
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
    C.C_ZNK10QEventLoop10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEventLoop", "metaObject", args)
  }

}

// event(class QEvent *)
func (this *QEventLoop) event(args ...interface{}) () {
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
    var arg0 = args[0].(QEvent).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QEventLoop5eventEP6QEvent(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QEventLoop", "event", args)
  }

}

// ~QEventLoopLocker()
func (this *QEventLoopLocker) FreeQEventLoopLocker(args ...interface{}) () {
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
    C.C_ZN16QEventLoopLockerD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEventLoopLocker", "~QEventLoopLocker", args)
  }

}

// QEventLoopLocker()
func NewQEventLoopLocker(args ...interface{}) QEventLoopLocker {
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
    C.C_ZN16QEventLoopLockerC2Ev(qthis)
  case 1:
    // invoke: _ZN16QEventLoopLockerC1EP7QThread
    // invoke: void QEventLoopLocker(class QThread *)
    var arg0 = args[0].(QThread).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN16QEventLoopLockerC2EP7QThread(qthis, arg0)
  case 2:
    // invoke: _ZN16QEventLoopLockerC1EP10QEventLoop
    // invoke: void QEventLoopLocker(class QEventLoop *)
    var arg0 = args[0].(QEventLoop).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN16QEventLoopLockerC2EP10QEventLoop(qthis, arg0)
  default:
    qtrt.ErrorResolve("QEventLoopLocker", "QEventLoopLocker", args)
  }

  return QEventLoopLocker{}
}

// <= body block end

