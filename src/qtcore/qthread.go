package qtcore
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
extern void* C_ZN7QThread13currentThreadEv(); // 4
  // proto:  void QThread::terminate();
extern void C_ZN7QThread9terminateEv(void* qthis); // 4
  // proto: static void QThread::sleep(unsigned long );
extern void C_ZN7QThread5sleepEm(int32_t arg0); // 4
  // proto: static void QThread::msleep(unsigned long );
extern void C_ZN7QThread6msleepEm(int32_t arg0); // 4
  // proto:  uint QThread::stackSize();
extern int32_t C_ZNK7QThread9stackSizeEv(void* qthis); // 4
  // proto:  bool QThread::event(QEvent * event);
extern bool C_ZN7QThread5eventEP6QEvent(void* qthis, void* arg0); // 4
  // proto: static void QThread::usleep(unsigned long );
extern void C_ZN7QThread6usleepEm(int32_t arg0); // 4
  // proto:  void QThread::quit();
extern void C_ZN7QThread4quitEv(void* qthis); // 4
  // proto:  int QThread::loopLevel();
extern int32_t C_ZNK7QThread9loopLevelEv(void* qthis); // 4
  // proto:  QThread::Priority QThread::priority();
extern void C_ZNK7QThread8priorityEv(void* qthis); // 4
  // proto:  bool QThread::isFinished();
extern bool C_ZNK7QThread10isFinishedEv(void* qthis); // 4
  // proto:  void QThread::exit(int retcode);
extern void C_ZN7QThread4exitEi(void* qthis, int32_t arg0); // 4
  // proto:  void QThread::setEventDispatcher(QAbstractEventDispatcher * eventDispatcher);
extern void C_ZN7QThread18setEventDispatcherEP24QAbstractEventDispatcher(void* qthis, void* arg0); // 4
  // proto:  void QThread::setStackSize(uint stackSize);
extern void C_ZN7QThread12setStackSizeEj(void* qthis, int32_t arg0); // 4
  // proto:  void QThread::~QThread();
extern void C_ZN7QThreadD2Ev(void* qthis); // 4
  // proto:  bool QThread::isRunning();
extern bool C_ZNK7QThread9isRunningEv(void* qthis); // 4
  // proto:  bool QThread::isInterruptionRequested();
extern bool C_ZNK7QThread23isInterruptionRequestedEv(void* qthis); // 4
  // proto:  void QThread::requestInterruption();
extern void C_ZN7QThread19requestInterruptionEv(void* qthis); // 4
  // proto:  void QThread::QThread(QObject * parent);
extern void* C_ZN7QThreadC2EP7QObject(void* arg0); // 3
  // proto:  bool QThread::wait(unsigned long time);
extern bool C_ZN7QThread4waitEm(void* qthis, int32_t arg0); // 4
  // proto:  const QMetaObject * QThread::metaObject();
extern void C_ZNK7QThread10metaObjectEv(void* qthis); // 4
  // proto: static int QThread::idealThreadCount();
extern int32_t C_ZN7QThread16idealThreadCountEv(); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
//  _started QThread_started_signal;
//  _finished QThread_finished_signal;
}

// eventDispatcher()
func (this *QThread) Eventdispatcher(args ...interface{}) () {
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
    C.C_ZNK7QThread15eventDispatcherEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "eventDispatcher", args)
  }

  return
}

// currentThread()
func (this *QThread) Currentthread_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN7QThread13currentThreadEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QThread{}) // "QThread *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QThread", "currentThread", args)
  }

  return
}

// terminate()
func (this *QThread) Terminate(args ...interface{}) () {
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
    C.C_ZN7QThread9terminateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "terminate", args)
  }

  return
}

// sleep(unsigned long)
func (this *QThread) Sleep_S(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QThread5sleepEm(arg0)
  default:
    qtrt.ErrorResolve("QThread", "sleep", args)
  }

  return
}

// msleep(unsigned long)
func (this *QThread) Msleep_S(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QThread6msleepEm(arg0)
  default:
    qtrt.ErrorResolve("QThread", "msleep", args)
  }

  return
}

// stackSize()
func (this *QThread) Stacksize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QThread9stackSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "uint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QThread", "stackSize", args)
  }

  return
}

// event(class QEvent *)
func (this *QThread) Event(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QEvent).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN7QThread5eventEP6QEvent(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QThread", "event", args)
  }

  return
}

// usleep(unsigned long)
func (this *QThread) Usleep_S(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QThread6usleepEm(arg0)
  default:
    qtrt.ErrorResolve("QThread", "usleep", args)
  }

  return
}

// quit()
func (this *QThread) Quit(args ...interface{}) () {
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
    C.C_ZN7QThread4quitEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "quit", args)
  }

  return
}

// loopLevel()
func (this *QThread) Looplevel(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QThread9loopLevelEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QThread", "loopLevel", args)
  }

  return
}

// priority()
func (this *QThread) Priority(args ...interface{}) () {
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
    C.C_ZNK7QThread8priorityEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "priority", args)
  }

  return
}

// isFinished()
func (this *QThread) Isfinished(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QThread10isFinishedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QThread", "isFinished", args)
  }

  return
}

// exit(int)
func (this *QThread) Exit(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QThread4exitEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QThread", "exit", args)
  }

  return
}

// setEventDispatcher(class QAbstractEventDispatcher *)
func (this *QThread) Seteventdispatcher(args ...interface{}) () {
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
    var arg0 = args[0].(*QAbstractEventDispatcher).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QThread18setEventDispatcherEP24QAbstractEventDispatcher(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QThread", "setEventDispatcher", args)
  }

  return
}

// setStackSize(uint)
func (this *QThread) Setstacksize(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QThread12setStackSizeEj(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QThread", "setStackSize", args)
  }

  return
}

// ~QThread()
func (this *QThread) Freeqthread(args ...interface{}) () {
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
    C.C_ZN7QThreadD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "~QThread", args)
  }

  return
}

// isRunning()
func (this *QThread) Isrunning(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QThread9isRunningEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QThread", "isRunning", args)
  }

  return
}

// isInterruptionRequested()
func (this *QThread) Isinterruptionrequested(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QThread23isInterruptionRequestedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QThread", "isInterruptionRequested", args)
  }

  return
}

// requestInterruption()
func (this *QThread) Requestinterruption(args ...interface{}) () {
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
    C.C_ZN7QThread19requestInterruptionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "requestInterruption", args)
  }

  return
}

// QThread(class QObject *)
func NewQThread(args ...interface{}) *QThread {
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
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QThreadC2EP7QObject(arg0)
    return &QThread{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QThread", "QThread", args)
  }

  return nil // QThread{Qclsinst:qthis}
}

// wait(unsigned long)
func (this *QThread) Wait(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN7QThread4waitEm(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QThread", "wait", args)
  }

  return
}

// metaObject()
func (this *QThread) Metaobject(args ...interface{}) () {
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
    C.C_ZNK7QThread10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QThread", "metaObject", args)
  }

  return
}

// idealThreadCount()
func (this *QThread) Idealthreadcount_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN7QThread16idealThreadCountEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QThread", "idealThreadCount", args)
  }

  return
}

// currentThreadId()
func (this *QThread) Currentthreadid_S(args ...interface{}) () {
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

  return
}

// yieldCurrentThread()
func (this *QThread) Yieldcurrentthread_S(args ...interface{}) () {
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

  return
}

// <= body block end

