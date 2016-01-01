package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qabstracteventdispatcher.h
// dst-file: /src/core/qabstracteventdispatcher.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
// class sizeof(QAbstractEventDispatcher)=1
type QAbstractEventDispatcher struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _aboutToBlock QAbstractEventDispatcher_aboutToBlock_signal;
//  _awake QAbstractEventDispatcher_awake_signal;
}


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
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "hasPendingEvents", args)
  }

}


func NewQAbstractEventDispatcher(args ...interface{}) QAbstractEventDispatcher {
  return QAbstractEventDispatcher{}
}


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
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "filterNativeEvent", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "interrupt", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "registerSocketNotifier", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "installNativeEventFilter", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "removeNativeEventFilter", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "flush", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "unregisterSocketNotifier", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "wakeUp", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "metaObject", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "unregisterTimers", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "remainingTime", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "startingUp", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "closingDown", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "unregisterTimer", args)
  }

}

// <= body block end

