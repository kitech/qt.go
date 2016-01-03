package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
// src-file: /QtCore/qthread.h
// dst-file: /src/core/qthread.go
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
  // proto:  void QThread::QThread(QObject * parent);
extern void* dector_ZN7QThreadC1EP7QObject(void* arg0);
extern void _ZN7QThreadC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QThread::setEventDispatcher(QAbstractEventDispatcher * eventDispatcher);
extern void _ZN7QThread18setEventDispatcherEP24QAbstractEventDispatcher(void* qthis, void* arg0);
  // proto:  const QMetaObject * QThread::metaObject();
extern void _ZNK7QThread10metaObjectEv(void* qthis);
  // proto: static void QThread::yieldCurrentThread();
extern void _ZN7QThread18yieldCurrentThreadEv();
  // proto:  bool QThread::isInterruptionRequested();
extern void _ZNK7QThread23isInterruptionRequestedEv(void* qthis);
  // proto: static void QThread::msleep(unsigned long );
extern void _ZN7QThread6msleepEm(int32_t arg0);
  // proto:  void QThread::requestInterruption();
extern void _ZN7QThread19requestInterruptionEv(void* qthis);
  // proto:  void QThread::exit(int retcode);
extern void _ZN7QThread4exitEi(void* qthis, int32_t arg0);
  // proto:  bool QThread::event(QEvent * event);
extern void _ZN7QThread5eventEP6QEvent(void* qthis, void* arg0);
  // proto:  uint QThread::stackSize();
extern void _ZNK7QThread9stackSizeEv(void* qthis);
  // proto:  QAbstractEventDispatcher * QThread::eventDispatcher();
extern void _ZNK7QThread15eventDispatcherEv(void* qthis);
  // proto:  void QThread::setStackSize(uint stackSize);
extern void _ZN7QThread12setStackSizeEj(void* qthis, int32_t arg0);
  // proto:  bool QThread::isFinished();
extern void _ZNK7QThread10isFinishedEv(void* qthis);
  // proto: static void QThread::sleep(unsigned long );
extern void _ZN7QThread5sleepEm(int32_t arg0);
  // proto: static void QThread::usleep(unsigned long );
extern void _ZN7QThread6usleepEm(int32_t arg0);
  // proto: static int QThread::idealThreadCount();
extern void _ZN7QThread16idealThreadCountEv();
  // proto:  bool QThread::wait(unsigned long time);
extern void _ZN7QThread4waitEm(void* qthis, int32_t arg0);
  // proto: static QThread * QThread::currentThread();
extern void _ZN7QThread13currentThreadEv();
  // proto:  bool QThread::isRunning();
extern void _ZNK7QThread9isRunningEv(void* qthis);
  // proto:  void QThread::terminate();
extern void _ZN7QThread9terminateEv(void* qthis);
  // proto:  void QThread::~QThread();
extern void _ZN7QThreadD0Ev(void* qthis);
  // proto:  void QThread::quit();
extern void _ZN7QThread4quitEv(void* qthis);
  // proto:  int QThread::loopLevel();
extern void _ZNK7QThread9loopLevelEv(void* qthis);
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

// class sizeof(QThread)=1
type QThread struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _started QThread_started_signal;
//  _finished QThread_finished_signal;
}

  // proto:  void QThread::QThread(QObject * parent);
func NewQThread(args ...interface{}) QThread {
  return QThread{}
}

  // proto:  void QThread::setEventDispatcher(QAbstractEventDispatcher * eventDispatcher);
func (this *QThread) setEventDispatcher(args ...interface{}) () {
  // setEventDispatcher(class QAbstractEventDispatcher *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractEventDispatcher{}) // "QAbstractEventDispatcher *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QThread18setEventDispatcherEP24QAbstractEventDispatcher
    // invoke: void setEventDispatcher(class QAbstractEventDispatcher *)
    var arg0 = args[0].(QAbstractEventDispatcher).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QThread18setEventDispatcherEP24QAbstractEventDispatcher(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QThread", "setEventDispatcher", args)
  }

}

  // proto:  const QMetaObject * QThread::metaObject();
func (this *QThread) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QThread10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK7QThread10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "metaObject", args)
  }

}

  // proto: static void QThread::yieldCurrentThread();
func (this *QThread) yieldCurrentThread_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QThread", "yieldCurrentThread", args)
  }

}

  // proto:  bool QThread::isInterruptionRequested();
func (this *QThread) isInterruptionRequested(args ...interface{}) () {
  // isInterruptionRequested()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QThread23isInterruptionRequestedEv
    // invoke: bool isInterruptionRequested()
    C._ZNK7QThread23isInterruptionRequestedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "isInterruptionRequested", args)
  }

}

  // proto: static void QThread::msleep(unsigned long );
func (this *QThread) msleep_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QThread", "msleep", args)
  }

}

  // proto:  void QThread::requestInterruption();
func (this *QThread) requestInterruption(args ...interface{}) () {
  // requestInterruption()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QThread19requestInterruptionEv
    // invoke: void requestInterruption()
    C._ZN7QThread19requestInterruptionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "requestInterruption", args)
  }

}

  // proto:  void QThread::exit(int retcode);
func (this *QThread) exit(args ...interface{}) () {
  // exit(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QThread4exitEi
    // invoke: void exit(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QThread4exitEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QThread", "exit", args)
  }

}

  // proto:  bool QThread::event(QEvent * event);
func (this *QThread) event(args ...interface{}) () {
  // event(class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QThread5eventEP6QEvent
    // invoke: bool event(class QEvent *)
    var arg0 = args[0].(QEvent).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QThread5eventEP6QEvent(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QThread", "event", args)
  }

}

  // proto:  uint QThread::stackSize();
func (this *QThread) stackSize(args ...interface{}) () {
  // stackSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QThread9stackSizeEv
    // invoke: uint stackSize()
    C._ZNK7QThread9stackSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "stackSize", args)
  }

}

  // proto:  QAbstractEventDispatcher * QThread::eventDispatcher();
func (this *QThread) eventDispatcher(args ...interface{}) () {
  // eventDispatcher()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QThread15eventDispatcherEv
    // invoke: QAbstractEventDispatcher * eventDispatcher()
    C._ZNK7QThread15eventDispatcherEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "eventDispatcher", args)
  }

}

  // proto:  void QThread::setStackSize(uint stackSize);
func (this *QThread) setStackSize(args ...interface{}) () {
  // setStackSize(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QThread12setStackSizeEj
    // invoke: void setStackSize(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QThread12setStackSizeEj(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QThread", "setStackSize", args)
  }

}

  // proto:  bool QThread::isFinished();
func (this *QThread) isFinished(args ...interface{}) () {
  // isFinished()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QThread10isFinishedEv
    // invoke: bool isFinished()
    C._ZNK7QThread10isFinishedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "isFinished", args)
  }

}

  // proto: static void QThread::sleep(unsigned long );
func (this *QThread) sleep_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QThread", "sleep", args)
  }

}

  // proto: static void QThread::usleep(unsigned long );
func (this *QThread) usleep_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QThread", "usleep", args)
  }

}

  // proto: static int QThread::idealThreadCount();
func (this *QThread) idealThreadCount_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QThread", "idealThreadCount", args)
  }

}

  // proto:  bool QThread::wait(unsigned long time);
func (this *QThread) wait(args ...interface{}) () {
  // wait(unsigned long)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "unsigned long"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QThread4waitEm
    // invoke: bool wait(unsigned long)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QThread4waitEm(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QThread", "wait", args)
  }

}

  // proto: static QThread * QThread::currentThread();
func (this *QThread) currentThread_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QThread", "currentThread", args)
  }

}

  // proto:  bool QThread::isRunning();
func (this *QThread) isRunning(args ...interface{}) () {
  // isRunning()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QThread9isRunningEv
    // invoke: bool isRunning()
    C._ZNK7QThread9isRunningEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "isRunning", args)
  }

}

  // proto:  void QThread::terminate();
func (this *QThread) terminate(args ...interface{}) () {
  // terminate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QThread9terminateEv
    // invoke: void terminate()
    C._ZN7QThread9terminateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "terminate", args)
  }

}

  // proto:  void QThread::~QThread();
func (this *QThread) FreeQThread(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QThread", "~QThread", args)
  }

}

  // proto:  void QThread::quit();
func (this *QThread) quit(args ...interface{}) () {
  // quit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QThread4quitEv
    // invoke: void quit()
    C._ZN7QThread4quitEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "quit", args)
  }

}

  // proto:  int QThread::loopLevel();
func (this *QThread) loopLevel(args ...interface{}) () {
  // loopLevel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QThread9loopLevelEv
    // invoke: int loopLevel()
    C._ZNK7QThread9loopLevelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "loopLevel", args)
  }

}

// <= body block end

