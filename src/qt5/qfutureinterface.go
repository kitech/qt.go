package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QFutureInterfaceBase::waitForResume();
extern void C_ZN20QFutureInterfaceBase13waitForResumeEv(void* qthis); // 4
  // proto:  QString QFutureInterfaceBase::progressText();
extern void* C_ZNK20QFutureInterfaceBase12progressTextEv(void* qthis); // 4
  // proto:  void QFutureInterfaceBase::setThrottled(bool enable);
extern void C_ZN20QFutureInterfaceBase12setThrottledEb(void* qthis, bool arg0); // 4
  // proto:  void QFutureInterfaceBase::reportCanceled();
extern void C_ZN20QFutureInterfaceBase14reportCanceledEv(void* qthis); // 4
  // proto:  void QFutureInterfaceBase::setFilterMode(bool enable);
extern void C_ZN20QFutureInterfaceBase13setFilterModeEb(void* qthis, bool arg0); // 4
  // proto:  int QFutureInterfaceBase::progressMaximum();
extern int32_t C_ZNK20QFutureInterfaceBase15progressMaximumEv(void* qthis); // 4
  // proto:  void QFutureInterfaceBase::cancel();
extern void C_ZN20QFutureInterfaceBase6cancelEv(void* qthis); // 4
  // proto:  void QFutureInterfaceBase::reportStarted();
extern void C_ZN20QFutureInterfaceBase13reportStartedEv(void* qthis); // 4
  // proto:  void QFutureInterfaceBase::togglePaused();
extern void C_ZN20QFutureInterfaceBase12togglePausedEv(void* qthis); // 4
  // proto:  bool QFutureInterfaceBase::isProgressUpdateNeeded();
extern bool C_ZNK20QFutureInterfaceBase22isProgressUpdateNeededEv(void* qthis); // 4
  // proto:  int QFutureInterfaceBase::progressMinimum();
extern int32_t C_ZNK20QFutureInterfaceBase15progressMinimumEv(void* qthis); // 4
  // proto:  void QFutureInterfaceBase::waitForFinished();
extern void C_ZN20QFutureInterfaceBase15waitForFinishedEv(void* qthis); // 4
  // proto:  void QFutureInterfaceBase::QFutureInterfaceBase(const QFutureInterfaceBase & other);
extern void* C_ZN20QFutureInterfaceBaseC2ERKS_(void* arg0); // 3
  // proto:  bool QFutureInterfaceBase::isStarted();
extern bool C_ZNK20QFutureInterfaceBase9isStartedEv(void* qthis); // 4
  // proto:  QtPrivate::ExceptionStore & QFutureInterfaceBase::exceptionStore();
extern void C_ZN20QFutureInterfaceBase14exceptionStoreEv(void* qthis); // 4
  // proto:  QMutex * QFutureInterfaceBase::mutex();
extern void* C_ZNK20QFutureInterfaceBase5mutexEv(void* qthis); // 4
  // proto:  void QFutureInterfaceBase::setExpectedResultCount(int resultCount);
extern void C_ZN20QFutureInterfaceBase22setExpectedResultCountEi(void* qthis, int32_t arg0); // 4
  // proto:  void QFutureInterfaceBase::setRunnable(QRunnable * runnable);
extern void C_ZN20QFutureInterfaceBase11setRunnableEP9QRunnable(void* qthis, void* arg0); // 4
  // proto:  bool QFutureInterfaceBase::isThrottled();
extern bool C_ZNK20QFutureInterfaceBase11isThrottledEv(void* qthis); // 4
  // proto:  bool QFutureInterfaceBase::waitForNextResult();
extern bool C_ZN20QFutureInterfaceBase17waitForNextResultEv(void* qthis); // 4
  // proto:  bool QFutureInterfaceBase::isFinished();
extern bool C_ZNK20QFutureInterfaceBase10isFinishedEv(void* qthis); // 4
  // proto:  void QFutureInterfaceBase::reportFinished();
extern void C_ZN20QFutureInterfaceBase14reportFinishedEv(void* qthis); // 4
  // proto:  void QFutureInterfaceBase::waitForResult(int resultIndex);
extern void C_ZN20QFutureInterfaceBase13waitForResultEi(void* qthis, int32_t arg0); // 4
  // proto:  int QFutureInterfaceBase::expectedResultCount();
extern int32_t C_ZN20QFutureInterfaceBase19expectedResultCountEv(void* qthis); // 4
  // proto:  void QFutureInterfaceBase::setProgressValueAndText(int progressValue, const QString & progressText);
extern void C_ZN20QFutureInterfaceBase23setProgressValueAndTextEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  int QFutureInterfaceBase::resultCount();
extern int32_t C_ZNK20QFutureInterfaceBase11resultCountEv(void* qthis); // 4
  // proto:  void QFutureInterfaceBase::setProgressValue(int progressValue);
extern void C_ZN20QFutureInterfaceBase16setProgressValueEi(void* qthis, int32_t arg0); // 4
  // proto:  void QFutureInterfaceBase::~QFutureInterfaceBase();
extern void C_ZN20QFutureInterfaceBaseD2Ev(void* qthis); // 4
  // proto:  bool QFutureInterfaceBase::isRunning();
extern bool C_ZNK20QFutureInterfaceBase9isRunningEv(void* qthis); // 4
  // proto:  void QFutureInterfaceBase::setProgressRange(int minimum, int maximum);
extern void C_ZN20QFutureInterfaceBase16setProgressRangeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QtPrivate::ResultStoreBase & QFutureInterfaceBase::resultStoreBase();
extern void C_ZN20QFutureInterfaceBase15resultStoreBaseEv(void* qthis); // 4
  // proto:  bool QFutureInterfaceBase::isCanceled();
extern bool C_ZNK20QFutureInterfaceBase10isCanceledEv(void* qthis); // 4
  // proto:  bool QFutureInterfaceBase::isResultReadyAt(int index);
extern bool C_ZNK20QFutureInterfaceBase15isResultReadyAtEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QFutureInterfaceBase::isPaused();
extern bool C_ZNK20QFutureInterfaceBase8isPausedEv(void* qthis); // 4
  // proto:  void QFutureInterfaceBase::setThreadPool(QThreadPool * pool);
extern void C_ZN20QFutureInterfaceBase13setThreadPoolEP11QThreadPool(void* qthis, void* arg0); // 4
  // proto:  void QFutureInterfaceBase::setPaused(bool paused);
extern void C_ZN20QFutureInterfaceBase9setPausedEb(void* qthis, bool arg0); // 4
  // proto:  int QFutureInterfaceBase::progressValue();
extern int32_t C_ZNK20QFutureInterfaceBase13progressValueEv(void* qthis); // 4
  // proto:  void QFutureInterfaceBase::reportResultsReady(int beginIndex, int endIndex);
extern void C_ZN20QFutureInterfaceBase18reportResultsReadyEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QFutureInterface<void>::reportFinished(const void * );
extern void C_ZN16QFutureInterfaceIvE14reportFinishedEPKv(void* qthis, void* arg0); // 2
  // proto:  void QFutureInterface<void>::reportResult(const void * , int );
extern void C_ZN16QFutureInterfaceIvE12reportResultEPKvi(void* qthis, void* arg0, int32_t arg1); // 2
  // proto:  QFuture<void> QFutureInterface<void>::future();
extern void C_ZN16QFutureInterfaceIvE6futureEv(void* qthis); // 2
  // proto: static QFutureInterface<void> QFutureInterface<void>::canceledResult();
extern void C_ZN16QFutureInterfaceIvE14canceledResultEv(); // 2
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QFutureInterface<void>)=16
type QFutureInterfaceLvoidG struct {
  /*qbase*/ QFutureInterfaceBase;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// waitForResume()
func (this *QFutureInterfaceBase) Waitforresume(args ...interface{}) () {
  // waitForResume()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase13waitForResumeEv
    // invoke: void waitForResume()
    C.C_ZN20QFutureInterfaceBase13waitForResumeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "waitForResume", args)
  }

  return
}

// progressText()
func (this *QFutureInterfaceBase) Progresstext(args ...interface{}) (ret interface{}) {
  // progressText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QFutureInterfaceBase12progressTextEv
    // invoke: QString progressText()
    var ret0 = C.C_ZNK20QFutureInterfaceBase12progressTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "progressText", args)
  }

  return
}

// setThrottled(_Bool)
func (this *QFutureInterfaceBase) Setthrottled(args ...interface{}) () {
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
    // invoke: void setThrottled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN20QFutureInterfaceBase12setThrottledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "setThrottled", args)
  }

  return
}

// reportCanceled()
func (this *QFutureInterfaceBase) Reportcanceled(args ...interface{}) () {
  // reportCanceled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase14reportCanceledEv
    // invoke: void reportCanceled()
    C.C_ZN20QFutureInterfaceBase14reportCanceledEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "reportCanceled", args)
  }

  return
}

// setFilterMode(_Bool)
func (this *QFutureInterfaceBase) Setfiltermode(args ...interface{}) () {
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
    // invoke: void setFilterMode(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN20QFutureInterfaceBase13setFilterModeEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "setFilterMode", args)
  }

  return
}

// progressMaximum()
func (this *QFutureInterfaceBase) Progressmaximum(args ...interface{}) (ret interface{}) {
  // progressMaximum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QFutureInterfaceBase15progressMaximumEv
    // invoke: int progressMaximum()
    var ret0 = C.C_ZNK20QFutureInterfaceBase15progressMaximumEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "progressMaximum", args)
  }

  return
}

// cancel()
func (this *QFutureInterfaceBase) Cancel(args ...interface{}) () {
  // cancel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase6cancelEv
    // invoke: void cancel()
    C.C_ZN20QFutureInterfaceBase6cancelEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "cancel", args)
  }

  return
}

// reportStarted()
func (this *QFutureInterfaceBase) Reportstarted(args ...interface{}) () {
  // reportStarted()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase13reportStartedEv
    // invoke: void reportStarted()
    C.C_ZN20QFutureInterfaceBase13reportStartedEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "reportStarted", args)
  }

  return
}

// togglePaused()
func (this *QFutureInterfaceBase) Togglepaused(args ...interface{}) () {
  // togglePaused()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase12togglePausedEv
    // invoke: void togglePaused()
    C.C_ZN20QFutureInterfaceBase12togglePausedEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "togglePaused", args)
  }

  return
}

// isProgressUpdateNeeded()
func (this *QFutureInterfaceBase) Isprogressupdateneeded(args ...interface{}) (ret interface{}) {
  // isProgressUpdateNeeded()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QFutureInterfaceBase22isProgressUpdateNeededEv
    // invoke: bool isProgressUpdateNeeded()
    var ret0 = C.C_ZNK20QFutureInterfaceBase22isProgressUpdateNeededEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "isProgressUpdateNeeded", args)
  }

  return
}

// progressMinimum()
func (this *QFutureInterfaceBase) Progressminimum(args ...interface{}) (ret interface{}) {
  // progressMinimum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QFutureInterfaceBase15progressMinimumEv
    // invoke: int progressMinimum()
    var ret0 = C.C_ZNK20QFutureInterfaceBase15progressMinimumEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "progressMinimum", args)
  }

  return
}

// waitForFinished()
func (this *QFutureInterfaceBase) Waitforfinished(args ...interface{}) () {
  // waitForFinished()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase15waitForFinishedEv
    // invoke: void waitForFinished()
    C.C_ZN20QFutureInterfaceBase15waitForFinishedEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "waitForFinished", args)
  }

  return
}

// QFutureInterfaceBase(const class QFutureInterfaceBase &)
func NewQFutureInterfaceBase(args ...interface{}) *QFutureInterfaceBase {
  // QFutureInterfaceBase(const class QFutureInterfaceBase &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFutureInterfaceBase{}) // "const QFutureInterfaceBase &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBaseC1ERKS_
    // invoke: void QFutureInterfaceBase(const class QFutureInterfaceBase &)
    var arg0 = args[0].(*QFutureInterfaceBase).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QFutureInterfaceBaseC2ERKS_(arg0)
    return &QFutureInterfaceBase{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "QFutureInterfaceBase", args)
  }

  return nil // QFutureInterfaceBase{Qclsinst:qthis}
}

// isStarted()
func (this *QFutureInterfaceBase) Isstarted(args ...interface{}) (ret interface{}) {
  // isStarted()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QFutureInterfaceBase9isStartedEv
    // invoke: bool isStarted()
    var ret0 = C.C_ZNK20QFutureInterfaceBase9isStartedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "isStarted", args)
  }

  return
}

// exceptionStore()
func (this *QFutureInterfaceBase) Exceptionstore(args ...interface{}) () {
  // exceptionStore()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase14exceptionStoreEv
    // invoke: QtPrivate::ExceptionStore & exceptionStore()
    C.C_ZN20QFutureInterfaceBase14exceptionStoreEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "exceptionStore", args)
  }

  return
}

// mutex()
func (this *QFutureInterfaceBase) Mutex(args ...interface{}) (ret interface{}) {
  // mutex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QFutureInterfaceBase5mutexEv
    // invoke: QMutex * mutex()
    var ret0 = C.C_ZNK20QFutureInterfaceBase5mutexEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMutex{}) // "QMutex *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "mutex", args)
  }

  return
}

// setExpectedResultCount(int)
func (this *QFutureInterfaceBase) Setexpectedresultcount(args ...interface{}) () {
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
    // invoke: void setExpectedResultCount(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN20QFutureInterfaceBase22setExpectedResultCountEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "setExpectedResultCount", args)
  }

  return
}

// setRunnable(class QRunnable *)
func (this *QFutureInterfaceBase) Setrunnable(args ...interface{}) () {
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
    // invoke: void setRunnable(class QRunnable *)
    var arg0 = args[0].(*QRunnable).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN20QFutureInterfaceBase11setRunnableEP9QRunnable(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "setRunnable", args)
  }

  return
}

// isThrottled()
func (this *QFutureInterfaceBase) Isthrottled(args ...interface{}) (ret interface{}) {
  // isThrottled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QFutureInterfaceBase11isThrottledEv
    // invoke: bool isThrottled()
    var ret0 = C.C_ZNK20QFutureInterfaceBase11isThrottledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "isThrottled", args)
  }

  return
}

// waitForNextResult()
func (this *QFutureInterfaceBase) Waitfornextresult(args ...interface{}) (ret interface{}) {
  // waitForNextResult()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase17waitForNextResultEv
    // invoke: bool waitForNextResult()
    var ret0 = C.C_ZN20QFutureInterfaceBase17waitForNextResultEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "waitForNextResult", args)
  }

  return
}

// isFinished()
func (this *QFutureInterfaceBase) Isfinished(args ...interface{}) (ret interface{}) {
  // isFinished()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QFutureInterfaceBase10isFinishedEv
    // invoke: bool isFinished()
    var ret0 = C.C_ZNK20QFutureInterfaceBase10isFinishedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "isFinished", args)
  }

  return
}

// reportFinished()
func (this *QFutureInterfaceBase) Reportfinished(args ...interface{}) () {
  // reportFinished()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase14reportFinishedEv
    // invoke: void reportFinished()
    C.C_ZN20QFutureInterfaceBase14reportFinishedEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "reportFinished", args)
  }

  return
}

// waitForResult(int)
func (this *QFutureInterfaceBase) Waitforresult(args ...interface{}) () {
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
    // invoke: void waitForResult(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN20QFutureInterfaceBase13waitForResultEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "waitForResult", args)
  }

  return
}

// expectedResultCount()
func (this *QFutureInterfaceBase) Expectedresultcount(args ...interface{}) (ret interface{}) {
  // expectedResultCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase19expectedResultCountEv
    // invoke: int expectedResultCount()
    var ret0 = C.C_ZN20QFutureInterfaceBase19expectedResultCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "expectedResultCount", args)
  }

  return
}

// setProgressValueAndText(int, const class QString &)
func (this *QFutureInterfaceBase) Setprogressvalueandtext(args ...interface{}) () {
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
    // invoke: void setProgressValueAndText(int, const class QString &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QFutureInterfaceBase23setProgressValueAndTextEiRK7QString(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "setProgressValueAndText", args)
  }

  return
}

// resultCount()
func (this *QFutureInterfaceBase) Resultcount(args ...interface{}) (ret interface{}) {
  // resultCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QFutureInterfaceBase11resultCountEv
    // invoke: int resultCount()
    var ret0 = C.C_ZNK20QFutureInterfaceBase11resultCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "resultCount", args)
  }

  return
}

// setProgressValue(int)
func (this *QFutureInterfaceBase) Setprogressvalue(args ...interface{}) () {
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
    // invoke: void setProgressValue(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN20QFutureInterfaceBase16setProgressValueEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "setProgressValue", args)
  }

  return
}

// ~QFutureInterfaceBase()
func (this *QFutureInterfaceBase) Freeqfutureinterfacebase(args ...interface{}) () {
  // ~QFutureInterfaceBase()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBaseD0Ev
    // invoke: void ~QFutureInterfaceBase()
    C.C_ZN20QFutureInterfaceBaseD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "~QFutureInterfaceBase", args)
  }

  return
}

// isRunning()
func (this *QFutureInterfaceBase) Isrunning(args ...interface{}) (ret interface{}) {
  // isRunning()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QFutureInterfaceBase9isRunningEv
    // invoke: bool isRunning()
    var ret0 = C.C_ZNK20QFutureInterfaceBase9isRunningEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "isRunning", args)
  }

  return
}

// setProgressRange(int, int)
func (this *QFutureInterfaceBase) Setprogressrange(args ...interface{}) () {
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
    // invoke: void setProgressRange(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN20QFutureInterfaceBase16setProgressRangeEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "setProgressRange", args)
  }

  return
}

// resultStoreBase()
func (this *QFutureInterfaceBase) Resultstorebase(args ...interface{}) () {
  // resultStoreBase()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QFutureInterfaceBase15resultStoreBaseEv
    // invoke: QtPrivate::ResultStoreBase & resultStoreBase()
    C.C_ZN20QFutureInterfaceBase15resultStoreBaseEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "resultStoreBase", args)
  }

  return
}

// isCanceled()
func (this *QFutureInterfaceBase) Iscanceled(args ...interface{}) (ret interface{}) {
  // isCanceled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QFutureInterfaceBase10isCanceledEv
    // invoke: bool isCanceled()
    var ret0 = C.C_ZNK20QFutureInterfaceBase10isCanceledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "isCanceled", args)
  }

  return
}

// isResultReadyAt(int)
func (this *QFutureInterfaceBase) Isresultreadyat(args ...interface{}) (ret interface{}) {
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
    // invoke: bool isResultReadyAt(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK20QFutureInterfaceBase15isResultReadyAtEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "isResultReadyAt", args)
  }

  return
}

// isPaused()
func (this *QFutureInterfaceBase) Ispaused(args ...interface{}) (ret interface{}) {
  // isPaused()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QFutureInterfaceBase8isPausedEv
    // invoke: bool isPaused()
    var ret0 = C.C_ZNK20QFutureInterfaceBase8isPausedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "isPaused", args)
  }

  return
}

// setThreadPool(class QThreadPool *)
func (this *QFutureInterfaceBase) Setthreadpool(args ...interface{}) () {
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
    // invoke: void setThreadPool(class QThreadPool *)
    var arg0 = args[0].(*QThreadPool).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN20QFutureInterfaceBase13setThreadPoolEP11QThreadPool(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "setThreadPool", args)
  }

  return
}

// setPaused(_Bool)
func (this *QFutureInterfaceBase) Setpaused(args ...interface{}) () {
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
    // invoke: void setPaused(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN20QFutureInterfaceBase9setPausedEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "setPaused", args)
  }

  return
}

// progressValue()
func (this *QFutureInterfaceBase) Progressvalue(args ...interface{}) (ret interface{}) {
  // progressValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QFutureInterfaceBase13progressValueEv
    // invoke: int progressValue()
    var ret0 = C.C_ZNK20QFutureInterfaceBase13progressValueEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "progressValue", args)
  }

  return
}

// reportResultsReady(int, int)
func (this *QFutureInterfaceBase) Reportresultsready(args ...interface{}) () {
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
    // invoke: void reportResultsReady(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN20QFutureInterfaceBase18reportResultsReadyEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QFutureInterfaceBase", "reportResultsReady", args)
  }

  return
}

// reportFinished(const void *)
func (this *QFutureInterfaceLvoidG) Reportfinished(args ...interface{}) () {
  // reportFinished(const void *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.VoidpTy() // "const void *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QFutureInterfaceIvE14reportFinishedEPKv
    // invoke: void reportFinished(const void *)
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
    C.C_ZN16QFutureInterfaceIvE14reportFinishedEPKv(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFutureInterface<void>", "reportFinished", args)
  }

  return
}

// reportResult(const void *, int)
func (this *QFutureInterfaceLvoidG) Reportresult(args ...interface{}) () {
  // reportResult(const void *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.VoidpTy() // "const void *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QFutureInterfaceIvE12reportResultEPKvi
    // invoke: void reportResult(const void *, int)
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN16QFutureInterfaceIvE12reportResultEPKvi(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QFutureInterface<void>", "reportResult", args)
  }

  return
}

// future()
func (this *QFutureInterfaceLvoidG) Future(args ...interface{}) () {
  // future()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QFutureInterfaceIvE6futureEv
    // invoke: QFuture<void> future()
    C.C_ZN16QFutureInterfaceIvE6futureEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFutureInterface<void>", "future", args)
  }

  return
}

// canceledResult()
func (this *QFutureInterfaceLvoidG) Canceledresult_S(args ...interface{}) () {
  // canceledResult()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QFutureInterfaceIvE14canceledResultEv
    // invoke: QFutureInterface<void> canceledResult()
    C.C_ZN16QFutureInterfaceIvE14canceledResultEv()
  default:
    qtrt.ErrorResolve("QFutureInterface<void>", "canceledResult", args)
  }

  return
}

// <= body block end

