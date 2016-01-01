package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qfutureinterface.h
// dst-file: /src/core/qfutureinterface.go
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
// class sizeof(QFutureInterfaceBase)=16
type QFutureInterfaceBase struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QFutureInterfaceBase) progressMinimum(args ...interface{}) () {
  // progressMinimum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QFutureInterfaceBase15progressMinimumEv
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "progressMinimum", args)
  }

}


func (this *QFutureInterfaceBase) isStarted(args ...interface{}) () {
  // isStarted()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QFutureInterfaceBase9isStartedEv
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "isStarted", args)
  }

}


func (this *QFutureInterfaceBase) mutex(args ...interface{}) () {
  // mutex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QFutureInterfaceBase5mutexEv
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "mutex", args)
  }

}


func (this *QFutureInterfaceBase) isResultReadyAt(args ...interface{}) () {
  // isResultReadyAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QFutureInterfaceBase15isResultReadyAtEi
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "isResultReadyAt", args)
  }

}


func (this *QFutureInterfaceBase) setPaused(args ...interface{}) () {
  // setPaused(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase9setPausedEb
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "setPaused", args)
  }

}


func (this *QFutureInterfaceBase) expectedResultCount(args ...interface{}) () {
  // expectedResultCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase19expectedResultCountEv
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "expectedResultCount", args)
  }

}


func (this *QFutureInterfaceBase) waitForFinished(args ...interface{}) () {
  // waitForFinished()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase15waitForFinishedEv
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "waitForFinished", args)
  }

}


func (this *QFutureInterfaceBase) isRunning(args ...interface{}) () {
  // isRunning()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QFutureInterfaceBase9isRunningEv
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "isRunning", args)
  }

}


func (this *QFutureInterfaceBase) FreeQFutureInterfaceBase(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "~QFutureInterfaceBase", args)
  }

}


func (this *QFutureInterfaceBase) cancel(args ...interface{}) () {
  // cancel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase6cancelEv
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "cancel", args)
  }

}


func (this *QFutureInterfaceBase) reportStarted(args ...interface{}) () {
  // reportStarted()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase13reportStartedEv
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "reportStarted", args)
  }

}


func (this *QFutureInterfaceBase) setRunnable(args ...interface{}) () {
  // setRunnable(class QRunnable *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRunnable{}) // "QRunnable *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase11setRunnableEP9QRunnable
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "setRunnable", args)
  }

}


func (this *QFutureInterfaceBase) isCanceled(args ...interface{}) () {
  // isCanceled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QFutureInterfaceBase10isCanceledEv
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "isCanceled", args)
  }

}


func (this *QFutureInterfaceBase) progressText(args ...interface{}) () {
  // progressText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QFutureInterfaceBase12progressTextEv
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "progressText", args)
  }

}


func (this *QFutureInterfaceBase) isProgressUpdateNeeded(args ...interface{}) () {
  // isProgressUpdateNeeded()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QFutureInterfaceBase22isProgressUpdateNeededEv
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "isProgressUpdateNeeded", args)
  }

}


func (this *QFutureInterfaceBase) setExpectedResultCount(args ...interface{}) () {
  // setExpectedResultCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase22setExpectedResultCountEi
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "setExpectedResultCount", args)
  }

}


func (this *QFutureInterfaceBase) reportResultsReady(args ...interface{}) () {
  // reportResultsReady(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase18reportResultsReadyEii
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "reportResultsReady", args)
  }

}


func (this *QFutureInterfaceBase) reportCanceled(args ...interface{}) () {
  // reportCanceled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase14reportCanceledEv
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "reportCanceled", args)
  }

}


func (this *QFutureInterfaceBase) resultCount(args ...interface{}) () {
  // resultCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QFutureInterfaceBase11resultCountEv
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "resultCount", args)
  }

}


func NewQFutureInterfaceBase(args ...interface{}) QFutureInterfaceBase {
  return QFutureInterfaceBase{}
}


func (this *QFutureInterfaceBase) progressValue(args ...interface{}) () {
  // progressValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QFutureInterfaceBase13progressValueEv
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "progressValue", args)
  }

}


func (this *QFutureInterfaceBase) isThrottled(args ...interface{}) () {
  // isThrottled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QFutureInterfaceBase11isThrottledEv
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "isThrottled", args)
  }

}


func (this *QFutureInterfaceBase) setProgressRange(args ...interface{}) () {
  // setProgressRange(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase16setProgressRangeEii
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "setProgressRange", args)
  }

}


func (this *QFutureInterfaceBase) setThrottled(args ...interface{}) () {
  // setThrottled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase12setThrottledEb
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "setThrottled", args)
  }

}


func (this *QFutureInterfaceBase) setProgressValueAndText(args ...interface{}) () {
  // setProgressValueAndText(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase23setProgressValueAndTextEiRK7QString
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "setProgressValueAndText", args)
  }

}


func (this *QFutureInterfaceBase) togglePaused(args ...interface{}) () {
  // togglePaused()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase12togglePausedEv
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "togglePaused", args)
  }

}


func (this *QFutureInterfaceBase) waitForResult(args ...interface{}) () {
  // waitForResult(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase13waitForResultEi
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "waitForResult", args)
  }

}


func (this *QFutureInterfaceBase) isPaused(args ...interface{}) () {
  // isPaused()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QFutureInterfaceBase8isPausedEv
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "isPaused", args)
  }

}


func (this *QFutureInterfaceBase) waitForNextResult(args ...interface{}) () {
  // waitForNextResult()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase17waitForNextResultEv
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "waitForNextResult", args)
  }

}


func (this *QFutureInterfaceBase) reportFinished(args ...interface{}) () {
  // reportFinished()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase14reportFinishedEv
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "reportFinished", args)
  }

}


func (this *QFutureInterfaceBase) setFilterMode(args ...interface{}) () {
  // setFilterMode(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase13setFilterModeEb
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "setFilterMode", args)
  }

}


func (this *QFutureInterfaceBase) progressMaximum(args ...interface{}) () {
  // progressMaximum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QFutureInterfaceBase15progressMaximumEv
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "progressMaximum", args)
  }

}


func (this *QFutureInterfaceBase) setThreadPool(args ...interface{}) () {
  // setThreadPool(class QThreadPool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QThreadPool{}) // "QThreadPool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase13setThreadPoolEP11QThreadPool
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "setThreadPool", args)
  }

}


func (this *QFutureInterfaceBase) waitForResume(args ...interface{}) () {
  // waitForResume()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase13waitForResumeEv
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "waitForResume", args)
  }

}


func (this *QFutureInterfaceBase) setProgressValue(args ...interface{}) () {
  // setProgressValue(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase16setProgressValueEi
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "setProgressValue", args)
  }

}


func (this *QFutureInterfaceBase) isFinished(args ...interface{}) () {
  // isFinished()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QFutureInterfaceBase10isFinishedEv
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "isFinished", args)
  }

}

// <= body block end

