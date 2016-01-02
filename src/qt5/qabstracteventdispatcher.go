package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
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
  // proto:  bool QAbstractEventDispatcher::hasPendingEvents();
extern void _ZN24QAbstractEventDispatcher16hasPendingEventsEv(void* qthis);
  // proto:  void QAbstractEventDispatcher::QAbstractEventDispatcher(QObject * parent);
extern void* dector_ZN24QAbstractEventDispatcherC1EP7QObject(void* arg0);
extern void _ZN24QAbstractEventDispatcherC1EP7QObject(void* qthis, void* arg0);
  // proto: static QAbstractEventDispatcher * QAbstractEventDispatcher::instance(QThread * thread);
extern void _ZN24QAbstractEventDispatcher8instanceEP7QThread(void* arg0);
  // proto:  bool QAbstractEventDispatcher::filterNativeEvent(const QByteArray & eventType, void * message, long * result);
extern void _ZN24QAbstractEventDispatcher17filterNativeEventERK10QByteArrayPvPl(void* qthis, void* arg0, void* arg1, long* arg2);
  // proto:  void QAbstractEventDispatcher::~QAbstractEventDispatcher();
extern void _ZN24QAbstractEventDispatcherD0Ev(void* qthis);
  // proto:  void QAbstractEventDispatcher::interrupt();
extern void _ZN24QAbstractEventDispatcher9interruptEv(void* qthis);
  // proto:  void QAbstractEventDispatcher::registerSocketNotifier(QSocketNotifier * notifier);
extern void _ZN24QAbstractEventDispatcher22registerSocketNotifierEP15QSocketNotifier(void* qthis, void* arg0);
  // proto:  void QAbstractEventDispatcher::installNativeEventFilter(QAbstractNativeEventFilter * filterObj);
extern void _ZN24QAbstractEventDispatcher24installNativeEventFilterEP26QAbstractNativeEventFilter(void* qthis, void* arg0);
  // proto:  void QAbstractEventDispatcher::removeNativeEventFilter(QAbstractNativeEventFilter * filterObj);
extern void _ZN24QAbstractEventDispatcher23removeNativeEventFilterEP26QAbstractNativeEventFilter(void* qthis, void* arg0);
  // proto:  void QAbstractEventDispatcher::flush();
extern void _ZN24QAbstractEventDispatcher5flushEv(void* qthis);
  // proto:  void QAbstractEventDispatcher::unregisterSocketNotifier(QSocketNotifier * notifier);
extern void _ZN24QAbstractEventDispatcher24unregisterSocketNotifierEP15QSocketNotifier(void* qthis, void* arg0);
  // proto:  void QAbstractEventDispatcher::wakeUp();
extern void _ZN24QAbstractEventDispatcher6wakeUpEv(void* qthis);
  // proto:  const QMetaObject * QAbstractEventDispatcher::metaObject();
extern void _ZNK24QAbstractEventDispatcher10metaObjectEv(void* qthis);
  // proto:  bool QAbstractEventDispatcher::unregisterTimers(QObject * object);
extern void _ZN24QAbstractEventDispatcher16unregisterTimersEP7QObject(void* qthis, void* arg0);
  // proto:  int QAbstractEventDispatcher::remainingTime(int timerId);
extern void _ZN24QAbstractEventDispatcher13remainingTimeEi(void* qthis, int arg0);
  // proto:  void QAbstractEventDispatcher::startingUp();
extern void _ZN24QAbstractEventDispatcher10startingUpEv(void* qthis);
  // proto:  void QAbstractEventDispatcher::closingDown();
extern void _ZN24QAbstractEventDispatcher11closingDownEv(void* qthis);
  // proto:  bool QAbstractEventDispatcher::unregisterTimer(int timerId);
extern void _ZN24QAbstractEventDispatcher15unregisterTimerEi(void* qthis, int arg0);
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

// class sizeof(QAbstractEventDispatcher)=1
type QAbstractEventDispatcher struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _aboutToBlock QAbstractEventDispatcher_aboutToBlock_signal;
//  _awake QAbstractEventDispatcher_awake_signal;
}

  // proto:  bool QAbstractEventDispatcher::hasPendingEvents();
func (this *QAbstractEventDispatcher) hasPendingEvents(args ...interface{}) () {
  // hasPendingEvents()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractEventDispatcher16hasPendingEventsEv
    // invoke: bool hasPendingEvents()
    C._ZN24QAbstractEventDispatcher16hasPendingEventsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "hasPendingEvents", args)
  }

}

  // proto:  void QAbstractEventDispatcher::QAbstractEventDispatcher(QObject * parent);
func NewQAbstractEventDispatcher(args ...interface{}) QAbstractEventDispatcher {
  return QAbstractEventDispatcher{}
}

  // proto: static QAbstractEventDispatcher * QAbstractEventDispatcher::instance(QThread * thread);
func (this *QAbstractEventDispatcher) instance_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "instance", args)
  }

}

  // proto:  bool QAbstractEventDispatcher::filterNativeEvent(const QByteArray & eventType, void * message, long * result);
func (this *QAbstractEventDispatcher) filterNativeEvent(args ...interface{}) () {
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
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    C._ZN24QAbstractEventDispatcher17filterNativeEventERK10QByteArrayPvPl(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "filterNativeEvent", args)
  }

}

  // proto:  void QAbstractEventDispatcher::~QAbstractEventDispatcher();
func (this *QAbstractEventDispatcher) FreeQAbstractEventDispatcher(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "~QAbstractEventDispatcher", args)
  }

}

  // proto:  void QAbstractEventDispatcher::interrupt();
func (this *QAbstractEventDispatcher) interrupt(args ...interface{}) () {
  // interrupt()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractEventDispatcher9interruptEv
    // invoke: void interrupt()
    C._ZN24QAbstractEventDispatcher9interruptEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "interrupt", args)
  }

}

  // proto:  void QAbstractEventDispatcher::registerSocketNotifier(QSocketNotifier * notifier);
func (this *QAbstractEventDispatcher) registerSocketNotifier(args ...interface{}) () {
  // registerSocketNotifier(class QSocketNotifier *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSocketNotifier{}) // "QSocketNotifier *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractEventDispatcher22registerSocketNotifierEP15QSocketNotifier
    // invoke: void registerSocketNotifier(class QSocketNotifier *)
    var arg0 = args[0].(QSocketNotifier).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN24QAbstractEventDispatcher22registerSocketNotifierEP15QSocketNotifier(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "registerSocketNotifier", args)
  }

}

  // proto:  void QAbstractEventDispatcher::installNativeEventFilter(QAbstractNativeEventFilter * filterObj);
func (this *QAbstractEventDispatcher) installNativeEventFilter(args ...interface{}) () {
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
    var arg0 = args[0].(QAbstractNativeEventFilter).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN24QAbstractEventDispatcher24installNativeEventFilterEP26QAbstractNativeEventFilter(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "installNativeEventFilter", args)
  }

}

  // proto:  void QAbstractEventDispatcher::removeNativeEventFilter(QAbstractNativeEventFilter * filterObj);
func (this *QAbstractEventDispatcher) removeNativeEventFilter(args ...interface{}) () {
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
    var arg0 = args[0].(QAbstractNativeEventFilter).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN24QAbstractEventDispatcher23removeNativeEventFilterEP26QAbstractNativeEventFilter(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "removeNativeEventFilter", args)
  }

}

  // proto:  void QAbstractEventDispatcher::flush();
func (this *QAbstractEventDispatcher) flush(args ...interface{}) () {
  // flush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractEventDispatcher5flushEv
    // invoke: void flush()
    C._ZN24QAbstractEventDispatcher5flushEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "flush", args)
  }

}

  // proto:  void QAbstractEventDispatcher::unregisterSocketNotifier(QSocketNotifier * notifier);
func (this *QAbstractEventDispatcher) unregisterSocketNotifier(args ...interface{}) () {
  // unregisterSocketNotifier(class QSocketNotifier *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSocketNotifier{}) // "QSocketNotifier *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractEventDispatcher24unregisterSocketNotifierEP15QSocketNotifier
    // invoke: void unregisterSocketNotifier(class QSocketNotifier *)
    var arg0 = args[0].(QSocketNotifier).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN24QAbstractEventDispatcher24unregisterSocketNotifierEP15QSocketNotifier(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "unregisterSocketNotifier", args)
  }

}

  // proto:  void QAbstractEventDispatcher::wakeUp();
func (this *QAbstractEventDispatcher) wakeUp(args ...interface{}) () {
  // wakeUp()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractEventDispatcher6wakeUpEv
    // invoke: void wakeUp()
    C._ZN24QAbstractEventDispatcher6wakeUpEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "wakeUp", args)
  }

}

  // proto:  const QMetaObject * QAbstractEventDispatcher::metaObject();
func (this *QAbstractEventDispatcher) metaObject(args ...interface{}) () {
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
    C._ZNK24QAbstractEventDispatcher10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "metaObject", args)
  }

}

  // proto:  bool QAbstractEventDispatcher::unregisterTimers(QObject * object);
func (this *QAbstractEventDispatcher) unregisterTimers(args ...interface{}) () {
  // unregisterTimers(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractEventDispatcher16unregisterTimersEP7QObject
    // invoke: bool unregisterTimers(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN24QAbstractEventDispatcher16unregisterTimersEP7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "unregisterTimers", args)
  }

}

  // proto:  int QAbstractEventDispatcher::remainingTime(int timerId);
func (this *QAbstractEventDispatcher) remainingTime(args ...interface{}) () {
  // remainingTime(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractEventDispatcher13remainingTimeEi
    // invoke: int remainingTime(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN24QAbstractEventDispatcher13remainingTimeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "remainingTime", args)
  }

}

  // proto:  void QAbstractEventDispatcher::startingUp();
func (this *QAbstractEventDispatcher) startingUp(args ...interface{}) () {
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
    C._ZN24QAbstractEventDispatcher10startingUpEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "startingUp", args)
  }

}

  // proto:  void QAbstractEventDispatcher::closingDown();
func (this *QAbstractEventDispatcher) closingDown(args ...interface{}) () {
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
    C._ZN24QAbstractEventDispatcher11closingDownEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "closingDown", args)
  }

}

  // proto:  bool QAbstractEventDispatcher::unregisterTimer(int timerId);
func (this *QAbstractEventDispatcher) unregisterTimer(args ...interface{}) () {
  // unregisterTimer(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractEventDispatcher15unregisterTimerEi
    // invoke: bool unregisterTimer(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN24QAbstractEventDispatcher15unregisterTimerEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "unregisterTimer", args)
  }

}

// <= body block end

