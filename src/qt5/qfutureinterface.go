package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtCore/qfutureinterface.h
// dst-file: /src/core/qfutureinterface.go
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
  // proto:  int QFutureInterfaceBase::progressMinimum();
extern void _ZNK20QFutureInterfaceBase15progressMinimumEv(void* qthis);
  // proto:  bool QFutureInterfaceBase::isStarted();
extern void _ZNK20QFutureInterfaceBase9isStartedEv(void* qthis);
  // proto:  QMutex * QFutureInterfaceBase::mutex();
extern void _ZNK20QFutureInterfaceBase5mutexEv(void* qthis);
  // proto:  bool QFutureInterfaceBase::isResultReadyAt(int index);
extern void _ZNK20QFutureInterfaceBase15isResultReadyAtEi(void* qthis, int arg0);
  // proto:  void QFutureInterfaceBase::setPaused(bool paused);
extern void _ZN20QFutureInterfaceBase9setPausedEb(void* qthis, bool arg0);
  // proto:  int QFutureInterfaceBase::expectedResultCount();
extern void _ZN20QFutureInterfaceBase19expectedResultCountEv(void* qthis);
  // proto:  void QFutureInterfaceBase::waitForFinished();
extern void _ZN20QFutureInterfaceBase15waitForFinishedEv(void* qthis);
  // proto:  bool QFutureInterfaceBase::isRunning();
extern void _ZNK20QFutureInterfaceBase9isRunningEv(void* qthis);
  // proto:  void QFutureInterfaceBase::~QFutureInterfaceBase();
extern void _ZN20QFutureInterfaceBaseD0Ev(void* qthis);
  // proto:  void QFutureInterfaceBase::cancel();
extern void _ZN20QFutureInterfaceBase6cancelEv(void* qthis);
  // proto:  void QFutureInterfaceBase::reportStarted();
extern void _ZN20QFutureInterfaceBase13reportStartedEv(void* qthis);
  // proto:  void QFutureInterfaceBase::setRunnable(QRunnable * runnable);
extern void _ZN20QFutureInterfaceBase11setRunnableEP9QRunnable(void* qthis, void* arg0);
  // proto:  bool QFutureInterfaceBase::isCanceled();
extern void _ZNK20QFutureInterfaceBase10isCanceledEv(void* qthis);
  // proto:  QString QFutureInterfaceBase::progressText();
extern void _ZNK20QFutureInterfaceBase12progressTextEv(void* qthis);
  // proto:  bool QFutureInterfaceBase::isProgressUpdateNeeded();
extern void _ZNK20QFutureInterfaceBase22isProgressUpdateNeededEv(void* qthis);
  // proto:  void QFutureInterfaceBase::setExpectedResultCount(int resultCount);
extern void _ZN20QFutureInterfaceBase22setExpectedResultCountEi(void* qthis, int arg0);
  // proto:  void QFutureInterfaceBase::reportResultsReady(int beginIndex, int endIndex);
extern void _ZN20QFutureInterfaceBase18reportResultsReadyEii(void* qthis, int arg0, int arg1);
  // proto:  void QFutureInterfaceBase::reportCanceled();
extern void _ZN20QFutureInterfaceBase14reportCanceledEv(void* qthis);
  // proto:  int QFutureInterfaceBase::resultCount();
extern void _ZNK20QFutureInterfaceBase11resultCountEv(void* qthis);
  // proto:  void QFutureInterfaceBase::QFutureInterfaceBase(const QFutureInterfaceBase & other);
extern void* dector_ZN20QFutureInterfaceBaseC1ERKS_(void* arg0);
extern void _ZN20QFutureInterfaceBaseC1ERKS_(void* qthis, void* arg0);
  // proto:  int QFutureInterfaceBase::progressValue();
extern void _ZNK20QFutureInterfaceBase13progressValueEv(void* qthis);
  // proto:  bool QFutureInterfaceBase::isThrottled();
extern void _ZNK20QFutureInterfaceBase11isThrottledEv(void* qthis);
  // proto:  void QFutureInterfaceBase::setProgressRange(int minimum, int maximum);
extern void _ZN20QFutureInterfaceBase16setProgressRangeEii(void* qthis, int arg0, int arg1);
  // proto:  void QFutureInterfaceBase::setThrottled(bool enable);
extern void _ZN20QFutureInterfaceBase12setThrottledEb(void* qthis, bool arg0);
  // proto:  void QFutureInterfaceBase::setProgressValueAndText(int progressValue, const QString & progressText);
extern void _ZN20QFutureInterfaceBase23setProgressValueAndTextEiRK7QString(void* qthis, int arg0, void* arg1);
  // proto:  void QFutureInterfaceBase::togglePaused();
extern void _ZN20QFutureInterfaceBase12togglePausedEv(void* qthis);
  // proto:  void QFutureInterfaceBase::waitForResult(int resultIndex);
extern void _ZN20QFutureInterfaceBase13waitForResultEi(void* qthis, int arg0);
  // proto:  bool QFutureInterfaceBase::isPaused();
extern void _ZNK20QFutureInterfaceBase8isPausedEv(void* qthis);
  // proto:  bool QFutureInterfaceBase::waitForNextResult();
extern void _ZN20QFutureInterfaceBase17waitForNextResultEv(void* qthis);
  // proto:  void QFutureInterfaceBase::reportFinished();
extern void _ZN20QFutureInterfaceBase14reportFinishedEv(void* qthis);
  // proto:  void QFutureInterfaceBase::setFilterMode(bool enable);
extern void _ZN20QFutureInterfaceBase13setFilterModeEb(void* qthis, bool arg0);
  // proto:  int QFutureInterfaceBase::progressMaximum();
extern void _ZNK20QFutureInterfaceBase15progressMaximumEv(void* qthis);
  // proto:  void QFutureInterfaceBase::setThreadPool(QThreadPool * pool);
extern void _ZN20QFutureInterfaceBase13setThreadPoolEP11QThreadPool(void* qthis, void* arg0);
  // proto:  void QFutureInterfaceBase::waitForResume();
extern void _ZN20QFutureInterfaceBase13waitForResumeEv(void* qthis);
  // proto:  void QFutureInterfaceBase::setProgressValue(int progressValue);
extern void _ZN20QFutureInterfaceBase16setProgressValueEi(void* qthis, int arg0);
  // proto:  bool QFutureInterfaceBase::isFinished();
extern void _ZNK20QFutureInterfaceBase10isFinishedEv(void* qthis);
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

// class sizeof(QFutureInterfaceBase)=16
type QFutureInterfaceBase struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  int QFutureInterfaceBase::progressMinimum();
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

  // proto:  bool QFutureInterfaceBase::isStarted();
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

  // proto:  QMutex * QFutureInterfaceBase::mutex();
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

  // proto:  bool QFutureInterfaceBase::isResultReadyAt(int index);
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

  // proto:  void QFutureInterfaceBase::setPaused(bool paused);
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

  // proto:  int QFutureInterfaceBase::expectedResultCount();
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

  // proto:  void QFutureInterfaceBase::waitForFinished();
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

  // proto:  bool QFutureInterfaceBase::isRunning();
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

  // proto:  void QFutureInterfaceBase::~QFutureInterfaceBase();
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

  // proto:  void QFutureInterfaceBase::cancel();
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

  // proto:  void QFutureInterfaceBase::reportStarted();
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

  // proto:  void QFutureInterfaceBase::setRunnable(QRunnable * runnable);
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

  // proto:  bool QFutureInterfaceBase::isCanceled();
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

  // proto:  QString QFutureInterfaceBase::progressText();
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

  // proto:  bool QFutureInterfaceBase::isProgressUpdateNeeded();
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

  // proto:  void QFutureInterfaceBase::setExpectedResultCount(int resultCount);
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

  // proto:  void QFutureInterfaceBase::reportResultsReady(int beginIndex, int endIndex);
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

  // proto:  void QFutureInterfaceBase::reportCanceled();
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

  // proto:  int QFutureInterfaceBase::resultCount();
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

  // proto:  void QFutureInterfaceBase::QFutureInterfaceBase(const QFutureInterfaceBase & other);
func NewQFutureInterfaceBase(args ...interface{}) QFutureInterfaceBase {
  return QFutureInterfaceBase{}
}

  // proto:  int QFutureInterfaceBase::progressValue();
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

  // proto:  bool QFutureInterfaceBase::isThrottled();
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

  // proto:  void QFutureInterfaceBase::setProgressRange(int minimum, int maximum);
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

  // proto:  void QFutureInterfaceBase::setThrottled(bool enable);
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

  // proto:  void QFutureInterfaceBase::setProgressValueAndText(int progressValue, const QString & progressText);
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

  // proto:  void QFutureInterfaceBase::togglePaused();
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

  // proto:  void QFutureInterfaceBase::waitForResult(int resultIndex);
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

  // proto:  bool QFutureInterfaceBase::isPaused();
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

  // proto:  bool QFutureInterfaceBase::waitForNextResult();
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

  // proto:  void QFutureInterfaceBase::reportFinished();
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

  // proto:  void QFutureInterfaceBase::setFilterMode(bool enable);
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

  // proto:  int QFutureInterfaceBase::progressMaximum();
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

  // proto:  void QFutureInterfaceBase::setThreadPool(QThreadPool * pool);
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

  // proto:  void QFutureInterfaceBase::waitForResume();
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

  // proto:  void QFutureInterfaceBase::setProgressValue(int progressValue);
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

  // proto:  bool QFutureInterfaceBase::isFinished();
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

