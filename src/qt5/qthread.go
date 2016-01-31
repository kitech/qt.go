package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QAbstractEventDispatcher * QThread::eventDispatcher();
extern void C_ZNK7QThread15eventDispatcherEv(void* qthis); // 4
  // proto: static QThread * QThread::currentThread();
extern void C_ZN7QThread13currentThreadEv(); // 4
  // proto:  void QThread::terminate();
extern void C_ZN7QThread9terminateEv(void* qthis); // 4
  // proto: static void QThread::sleep(unsigned long );
extern void C_ZN7QThread5sleepEm(int32_t arg0); // 4
  // proto: static void QThread::msleep(unsigned long );
extern void C_ZN7QThread6msleepEm(int32_t arg0); // 4
  // proto:  uint QThread::stackSize();
extern void C_ZNK7QThread9stackSizeEv(void* qthis); // 4
  // proto:  bool QThread::event(QEvent * event);
extern void C_ZN7QThread5eventEP6QEvent(void* qthis, void* arg0); // 4
  // proto: static void QThread::usleep(unsigned long );
extern void C_ZN7QThread6usleepEm(int32_t arg0); // 4
  // proto:  void QThread::quit();
extern void C_ZN7QThread4quitEv(void* qthis); // 4
  // proto:  int QThread::loopLevel();
extern void C_ZNK7QThread9loopLevelEv(void* qthis); // 4
  // proto:  QThread::Priority QThread::priority();
extern void C_ZNK7QThread8priorityEv(void* qthis); // 4
  // proto:  bool QThread::isFinished();
extern void C_ZNK7QThread10isFinishedEv(void* qthis); // 4
  // proto:  void QThread::exit(int retcode);
extern void C_ZN7QThread4exitEi(void* qthis, int32_t arg0); // 4
  // proto:  void QThread::setEventDispatcher(QAbstractEventDispatcher * eventDispatcher);
extern void C_ZN7QThread18setEventDispatcherEP24QAbstractEventDispatcher(void* qthis, void* arg0); // 4
  // proto:  void QThread::setStackSize(uint stackSize);
extern void C_ZN7QThread12setStackSizeEj(void* qthis, int32_t arg0); // 4
  // proto:  void QThread::~QThread();
extern void C_ZN7QThreadD2Ev(void* qthis); // 4
  // proto:  bool QThread::isRunning();
extern void C_ZNK7QThread9isRunningEv(void* qthis); // 4
  // proto:  bool QThread::isInterruptionRequested();
extern void C_ZNK7QThread23isInterruptionRequestedEv(void* qthis); // 4
  // proto:  void QThread::requestInterruption();
extern void C_ZN7QThread19requestInterruptionEv(void* qthis); // 4
  // proto:  void QThread::QThread(QObject * parent);
extern void C_ZN7QThreadC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  bool QThread::wait(unsigned long time);
extern void C_ZN7QThread4waitEm(void* qthis, int32_t arg0); // 4
  // proto:  const QMetaObject * QThread::metaObject();
extern void C_ZNK7QThread10metaObjectEv(void* qthis); // 4
  // proto: static int QThread::idealThreadCount();
extern void C_ZN7QThread16idealThreadCountEv(); // 4
  // proto: static Qt::HANDLE QThread::currentThreadId();
extern void C_ZN7QThread15currentThreadIdEv(); // 4
  // proto: static void QThread::yieldCurrentThread();
extern void C_ZN7QThread18yieldCurrentThreadEv(); // 4
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

// eventDispatcher()
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
    C.C_ZNK7QThread15eventDispatcherEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "eventDispatcher", args)
  }

}

// currentThread()
func (this *QThread) currentThread_s(args ...interface{}) () {
  // currentThread()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QThread13currentThreadEv
    // invoke: QThread * currentThread()
    C.C_ZN7QThread13currentThreadEv()
  default:
    qtrt.ErrorResolve("QThread", "currentThread", args)
  }

}

// terminate()
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
    C.C_ZN7QThread9terminateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "terminate", args)
  }

}

// sleep(unsigned long)
func (this *QThread) sleep_s(args ...interface{}) () {
  // sleep(unsigned long)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "unsigned long"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QThread5sleepEm
    // invoke: void sleep(unsigned long)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QThread5sleepEm(arg0)
  default:
    qtrt.ErrorResolve("QThread", "sleep", args)
  }

}

// msleep(unsigned long)
func (this *QThread) msleep_s(args ...interface{}) () {
  // msleep(unsigned long)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "unsigned long"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QThread6msleepEm
    // invoke: void msleep(unsigned long)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QThread6msleepEm(arg0)
  default:
    qtrt.ErrorResolve("QThread", "msleep", args)
  }

}

// stackSize()
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
    C.C_ZNK7QThread9stackSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "stackSize", args)
  }

}

// event(class QEvent *)
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
    C.C_ZN7QThread5eventEP6QEvent(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QThread", "event", args)
  }

}

// usleep(unsigned long)
func (this *QThread) usleep_s(args ...interface{}) () {
  // usleep(unsigned long)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "unsigned long"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QThread6usleepEm
    // invoke: void usleep(unsigned long)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QThread6usleepEm(arg0)
  default:
    qtrt.ErrorResolve("QThread", "usleep", args)
  }

}

// quit()
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
    C.C_ZN7QThread4quitEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "quit", args)
  }

}

// loopLevel()
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
    C.C_ZNK7QThread9loopLevelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "loopLevel", args)
  }

}

// priority()
func (this *QThread) priority(args ...interface{}) () {
  // priority()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QThread8priorityEv
    // invoke: QThread::Priority priority()
    C.C_ZNK7QThread8priorityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "priority", args)
  }

}

// isFinished()
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
    C.C_ZNK7QThread10isFinishedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "isFinished", args)
  }

}

// exit(int)
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
    C.C_ZN7QThread4exitEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QThread", "exit", args)
  }

}

// setEventDispatcher(class QAbstractEventDispatcher *)
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
    C.C_ZN7QThread18setEventDispatcherEP24QAbstractEventDispatcher(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QThread", "setEventDispatcher", args)
  }

}

// setStackSize(uint)
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
    C.C_ZN7QThread12setStackSizeEj(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QThread", "setStackSize", args)
  }

}

// ~QThread()
func (this *QThread) FreeQThread(args ...interface{}) () {
  // ~QThread()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QThreadD0Ev
    // invoke: void ~QThread()
    C.C_ZN7QThreadD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "~QThread", args)
  }

}

// isRunning()
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
    C.C_ZNK7QThread9isRunningEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "isRunning", args)
  }

}

// isInterruptionRequested()
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
    C.C_ZNK7QThread23isInterruptionRequestedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "isInterruptionRequested", args)
  }

}

// requestInterruption()
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
    C.C_ZN7QThread19requestInterruptionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "requestInterruption", args)
  }

}

// QThread(class QObject *)
func NewQThread(args ...interface{}) QThread {
  // QThread(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QThreadC1EP7QObject
    // invoke: void QThread(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN7QThreadC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QThread", "QThread", args)
  }

  return QThread{}
}

// wait(unsigned long)
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
    C.C_ZN7QThread4waitEm(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QThread", "wait", args)
  }

}

// metaObject()
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
    C.C_ZNK7QThread10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "metaObject", args)
  }

}

// idealThreadCount()
func (this *QThread) idealThreadCount_s(args ...interface{}) () {
  // idealThreadCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QThread16idealThreadCountEv
    // invoke: int idealThreadCount()
    C.C_ZN7QThread16idealThreadCountEv()
  default:
    qtrt.ErrorResolve("QThread", "idealThreadCount", args)
  }

}

// currentThreadId()
func (this *QThread) currentThreadId_s(args ...interface{}) () {
  // currentThreadId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QThread15currentThreadIdEv
    // invoke: Qt::HANDLE currentThreadId()
    C.C_ZN7QThread15currentThreadIdEv()
  default:
    qtrt.ErrorResolve("QThread", "currentThreadId", args)
  }

}

// yieldCurrentThread()
func (this *QThread) yieldCurrentThread_s(args ...interface{}) () {
  // yieldCurrentThread()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QThread18yieldCurrentThreadEv
    // invoke: void yieldCurrentThread()
    C.C_ZN7QThread18yieldCurrentThreadEv()
  default:
    qtrt.ErrorResolve("QThread", "yieldCurrentThread", args)
  }

}

// <= body block end

