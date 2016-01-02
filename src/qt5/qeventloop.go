package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
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
  // proto:  void QEventLoop::exit(int returnCode);
extern void _ZN10QEventLoop4exitEi(void* qthis, int arg0);
  // proto:  void QEventLoop::quit();
extern void _ZN10QEventLoop4quitEv(void* qthis);
  // proto:  void QEventLoop::QEventLoop(QObject * parent);
extern void* dector_ZN10QEventLoopC1EP7QObject(void* arg0);
extern void _ZN10QEventLoopC1EP7QObject(void* qthis, void* arg0);
  // proto:  bool QEventLoop::isRunning();
extern void _ZNK10QEventLoop9isRunningEv(void* qthis);
  // proto:  const QMetaObject * QEventLoop::metaObject();
extern void _ZNK10QEventLoop10metaObjectEv(void* qthis);
  // proto:  void QEventLoop::wakeUp();
extern void _ZN10QEventLoop6wakeUpEv(void* qthis);
  // proto:  void QEventLoop::~QEventLoop();
extern void _ZN10QEventLoopD0Ev(void* qthis);
  // proto:  bool QEventLoop::event(QEvent * event);
extern void _ZN10QEventLoop5eventEP6QEvent(void* qthis, void* arg0);
  // proto:  void QEventLoopLocker::QEventLoopLocker(QThread * thread);
extern void* dector_ZN16QEventLoopLockerC1EP7QThread(void* arg0);
extern void _ZN16QEventLoopLockerC1EP7QThread(void* qthis, void* arg0);
  // proto:  void QEventLoopLocker::QEventLoopLocker(QEventLoop * loop);
extern void* dector_ZN16QEventLoopLockerC1EP10QEventLoop(void* arg0);
extern void _ZN16QEventLoopLockerC1EP10QEventLoop(void* qthis, void* arg0);
  // proto:  void QEventLoopLocker::QEventLoopLocker();
extern void* dector_ZN16QEventLoopLockerC1Ev();
extern void _ZN16QEventLoopLockerC1Ev(void* qthis);
  // proto:  void QEventLoopLocker::QEventLoopLocker(const QEventLoopLocker & );
extern void* dector_ZN16QEventLoopLockerC1ERKS_(void* arg0);
extern void _ZN16QEventLoopLockerC1ERKS_(void* qthis, void* arg0);
  // proto:  void QEventLoopLocker::~QEventLoopLocker();
extern void _ZN16QEventLoopLockerD0Ev(void* qthis);
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
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QEventLoopLocker)=8
type QEventLoopLocker struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QEventLoop::exit(int returnCode);
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
    C._ZN10QEventLoop4exitEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QEventLoop", "exit", args)
  }

}

  // proto:  void QEventLoop::quit();
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
    C._ZN10QEventLoop4quitEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEventLoop", "quit", args)
  }

}

  // proto:  void QEventLoop::QEventLoop(QObject * parent);
func NewQEventLoop(args ...interface{}) QEventLoop {
  return QEventLoop{}
}

  // proto:  bool QEventLoop::isRunning();
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
    C._ZNK10QEventLoop9isRunningEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEventLoop", "isRunning", args)
  }

}

  // proto:  const QMetaObject * QEventLoop::metaObject();
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
    C._ZNK10QEventLoop10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEventLoop", "metaObject", args)
  }

}

  // proto:  void QEventLoop::wakeUp();
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
    C._ZN10QEventLoop6wakeUpEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEventLoop", "wakeUp", args)
  }

}

  // proto:  void QEventLoop::~QEventLoop();
func (this *QEventLoop) FreeQEventLoop(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QEventLoop", "~QEventLoop", args)
  }

}

  // proto:  bool QEventLoop::event(QEvent * event);
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
    C._ZN10QEventLoop5eventEP6QEvent(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QEventLoop", "event", args)
  }

}

  // proto:  void QEventLoopLocker::QEventLoopLocker(QThread * thread);
func NewQEventLoopLocker(args ...interface{}) QEventLoopLocker {
  return QEventLoopLocker{}
}

  // proto:  void QEventLoopLocker::~QEventLoopLocker();
func (this *QEventLoopLocker) FreeQEventLoopLocker(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QEventLoopLocker", "~QEventLoopLocker", args)
  }

}

// <= body block end

