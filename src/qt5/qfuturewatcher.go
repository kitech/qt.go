package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qfuturewatcher.h
// dst-file: /src/core/qfuturewatcher.go
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
// class sizeof(QFutureWatcherBase)=1
type QFutureWatcherBase struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _progressRangeChanged QFutureWatcherBase_progressRangeChanged_signal;
//  _resumed QFutureWatcherBase_resumed_signal;
//  _progressValueChanged QFutureWatcherBase_progressValueChanged_signal;
//  _started QFutureWatcherBase_started_signal;
//  _resultReadyAt QFutureWatcherBase_resultReadyAt_signal;
//  _resultsReadyAt QFutureWatcherBase_resultsReadyAt_signal;
//  _paused QFutureWatcherBase_paused_signal;
//  _canceled QFutureWatcherBase_canceled_signal;
//  _finished QFutureWatcherBase_finished_signal;
//  _progressTextChanged QFutureWatcherBase_progressTextChanged_signal;
}


func (this *QFutureWatcherBase) isRunning(args ...interface{}) () {
  // isRunning()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase9isRunningEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "isRunning", args)
 }

}


func (this *QFutureWatcherBase) setPaused(args ...interface{}) () {
  // setPaused(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase9setPausedEb
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "setPaused", args)
 }

}


func (this *QFutureWatcherBase) progressMinimum(args ...interface{}) () {
  // progressMinimum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase15progressMinimumEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "progressMinimum", args)
 }

}


func (this *QFutureWatcherBase) resume(args ...interface{}) () {
  // resume()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase6resumeEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "resume", args)
 }

}


func (this *QFutureWatcherBase) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase10metaObjectEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "metaObject", args)
 }

}


func (this *QFutureWatcherBase) isFinished(args ...interface{}) () {
  // isFinished()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase10isFinishedEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "isFinished", args)
 }

}


func (this *QFutureWatcherBase) progressMaximum(args ...interface{}) () {
  // progressMaximum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase15progressMaximumEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "progressMaximum", args)
 }

}


func (this *QFutureWatcherBase) event(args ...interface{}) () {
  // event(class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase5eventEP6QEvent
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "event", args)
 }

}


func (this *QFutureWatcherBase) isCanceled(args ...interface{}) () {
  // isCanceled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase10isCanceledEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "isCanceled", args)
 }

}


func (this *QFutureWatcherBase) progressValue(args ...interface{}) () {
  // progressValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase13progressValueEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "progressValue", args)
 }

}


func (this *QFutureWatcherBase) isStarted(args ...interface{}) () {
  // isStarted()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase9isStartedEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "isStarted", args)
 }

}


func (this *QFutureWatcherBase) setPendingResultsLimit(args ...interface{}) () {
  // setPendingResultsLimit(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase22setPendingResultsLimitEi
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "setPendingResultsLimit", args)
 }

}


func (this *QFutureWatcherBase) cancel(args ...interface{}) () {
  // cancel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase6cancelEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "cancel", args)
 }

}


func (this *QFutureWatcherBase) isPaused(args ...interface{}) () {
  // isPaused()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase8isPausedEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "isPaused", args)
 }

}


func (this *QFutureWatcherBase) pause(args ...interface{}) () {
  // pause()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase5pauseEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "pause", args)
 }

}


func (this *QFutureWatcherBase) progressText(args ...interface{}) () {
  // progressText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase12progressTextEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "progressText", args)
 }

}


func NewQFutureWatcherBase(args ...interface{})() {
}


func (this *QFutureWatcherBase) togglePaused(args ...interface{}) () {
  // togglePaused()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase12togglePausedEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "togglePaused", args)
 }

}


func (this *QFutureWatcherBase) waitForFinished(args ...interface{}) () {
  // waitForFinished()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase15waitForFinishedEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "waitForFinished", args)
 }

}

// <= body block end

