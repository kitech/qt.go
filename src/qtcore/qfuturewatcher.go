package qtcore
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtCore/qfuturewatcher.h
// dst-file: /src/core/qfuturewatcher.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "runtime"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  bool QFutureWatcherBase::isStarted();
extern bool C_ZNK18QFutureWatcherBase9isStartedEv(void* qthis); // 4
  // proto:  void QFutureWatcherBase::cancel();
extern void C_ZN18QFutureWatcherBase6cancelEv(void* qthis); // 4
  // proto:  bool QFutureWatcherBase::event(QEvent * event);
extern bool C_ZN18QFutureWatcherBase5eventEP6QEvent(void* qthis, void* arg0); // 4
  // proto:  void QFutureWatcherBase::togglePaused();
extern void C_ZN18QFutureWatcherBase12togglePausedEv(void* qthis); // 4
  // proto:  int QFutureWatcherBase::progressMinimum();
extern int32_t C_ZNK18QFutureWatcherBase15progressMinimumEv(void* qthis); // 4
  // proto:  void QFutureWatcherBase::waitForFinished();
extern void C_ZN18QFutureWatcherBase15waitForFinishedEv(void* qthis); // 4
  // proto:  void QFutureWatcherBase::resume();
extern void C_ZN18QFutureWatcherBase6resumeEv(void* qthis); // 4
  // proto:  int QFutureWatcherBase::progressMaximum();
extern int32_t C_ZNK18QFutureWatcherBase15progressMaximumEv(void* qthis); // 4
  // proto:  bool QFutureWatcherBase::isFinished();
extern bool C_ZNK18QFutureWatcherBase10isFinishedEv(void* qthis); // 4
  // proto:  QString QFutureWatcherBase::progressText();
extern void* C_ZNK18QFutureWatcherBase12progressTextEv(void* qthis); // 4
  // proto:  void QFutureWatcherBase::setPendingResultsLimit(int limit);
extern void C_ZN18QFutureWatcherBase22setPendingResultsLimitEi(void* qthis, int32_t arg0); // 4
  // proto:  void QFutureWatcherBase::pause();
extern void C_ZN18QFutureWatcherBase5pauseEv(void* qthis); // 4
  // proto:  const QMetaObject * QFutureWatcherBase::metaObject();
extern void C_ZNK18QFutureWatcherBase10metaObjectEv(void* qthis); // 4
  // proto:  void QFutureWatcherBase::QFutureWatcherBase(QObject * parent);
extern void* C_ZN18QFutureWatcherBaseC2EP7QObject(void* arg0); // 3
  // proto:  bool QFutureWatcherBase::isRunning();
extern bool C_ZNK18QFutureWatcherBase9isRunningEv(void* qthis); // 4
  // proto:  bool QFutureWatcherBase::isCanceled();
extern bool C_ZNK18QFutureWatcherBase10isCanceledEv(void* qthis); // 4
  // proto:  bool QFutureWatcherBase::isPaused();
extern bool C_ZNK18QFutureWatcherBase8isPausedEv(void* qthis); // 4
  // proto:  int QFutureWatcherBase::progressValue();
extern int32_t C_ZNK18QFutureWatcherBase13progressValueEv(void* qthis); // 4
  // proto:  void QFutureWatcherBase::setPaused(bool paused);
extern void C_ZN18QFutureWatcherBase9setPausedEb(void* qthis, bool arg0); // 4
  // proto:  void QFutureWatcher<void>::~QFutureWatcher();
extern void C_ZN14QFutureWatcherIvED2Ev(void* qthis); // 4
  // proto:  void QFutureWatcher<void>::QFutureWatcher(QObject * _parent);
extern void* C_ZN14QFutureWatcherIvEC2EP7QObject(void* arg0); // 1
  // proto:  QFuture<void> QFutureWatcher<void>::future();
extern void C_ZNK14QFutureWatcherIvE6futureEv(void* qthis); // 2
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QFutureWatcherBase)=1
type QFutureWatcherBase struct {
  /*qbase*/ QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
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

// class sizeof(QFutureWatcher<void>)=1
type QFutureWatcherLvoidG struct {
  /*qbase*/ QFutureWatcherBase;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// isStarted()
func (this *QFutureWatcherBase) IsStarted(args ...interface{}) (ret interface{}) {
  // isStarted()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase9isStartedEv
    // invoke: bool isStarted()
    var ret0 = C.C_ZNK18QFutureWatcherBase9isStartedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "isStarted", args)
  }

  return
}

// cancel()
func (this *QFutureWatcherBase) Cancel(args ...interface{}) () {
  // cancel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase6cancelEv
    // invoke: void cancel()
    C.C_ZN18QFutureWatcherBase6cancelEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "cancel", args)
  }

  return
}

// event(class QEvent *)
func (this *QFutureWatcherBase) Event(args ...interface{}) (ret interface{}) {
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
    // invoke: bool event(class QEvent *)
    var arg0 = args[0].(*QEvent).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN18QFutureWatcherBase5eventEP6QEvent(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "event", args)
  }

  return
}

// togglePaused()
func (this *QFutureWatcherBase) TogglePaused(args ...interface{}) () {
  // togglePaused()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase12togglePausedEv
    // invoke: void togglePaused()
    C.C_ZN18QFutureWatcherBase12togglePausedEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "togglePaused", args)
  }

  return
}

// progressMinimum()
func (this *QFutureWatcherBase) ProgressMinimum(args ...interface{}) (ret interface{}) {
  // progressMinimum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase15progressMinimumEv
    // invoke: int progressMinimum()
    var ret0 = C.C_ZNK18QFutureWatcherBase15progressMinimumEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "progressMinimum", args)
  }

  return
}

// waitForFinished()
func (this *QFutureWatcherBase) WaitForFinished(args ...interface{}) () {
  // waitForFinished()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase15waitForFinishedEv
    // invoke: void waitForFinished()
    C.C_ZN18QFutureWatcherBase15waitForFinishedEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "waitForFinished", args)
  }

  return
}

// resume()
func (this *QFutureWatcherBase) Resume(args ...interface{}) () {
  // resume()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase6resumeEv
    // invoke: void resume()
    C.C_ZN18QFutureWatcherBase6resumeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "resume", args)
  }

  return
}

// progressMaximum()
func (this *QFutureWatcherBase) ProgressMaximum(args ...interface{}) (ret interface{}) {
  // progressMaximum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase15progressMaximumEv
    // invoke: int progressMaximum()
    var ret0 = C.C_ZNK18QFutureWatcherBase15progressMaximumEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "progressMaximum", args)
  }

  return
}

// isFinished()
func (this *QFutureWatcherBase) IsFinished(args ...interface{}) (ret interface{}) {
  // isFinished()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase10isFinishedEv
    // invoke: bool isFinished()
    var ret0 = C.C_ZNK18QFutureWatcherBase10isFinishedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "isFinished", args)
  }

  return
}

// progressText()
func (this *QFutureWatcherBase) ProgressText(args ...interface{}) (ret interface{}) {
  // progressText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase12progressTextEv
    // invoke: QString progressText()
    var ret0 = C.C_ZNK18QFutureWatcherBase12progressTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "progressText", args)
  }

  return
}

// setPendingResultsLimit(int)
func (this *QFutureWatcherBase) SetPendingResultsLimit(args ...interface{}) () {
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
    // invoke: void setPendingResultsLimit(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN18QFutureWatcherBase22setPendingResultsLimitEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "setPendingResultsLimit", args)
  }

  return
}

// pause()
func (this *QFutureWatcherBase) Pause(args ...interface{}) () {
  // pause()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase5pauseEv
    // invoke: void pause()
    C.C_ZN18QFutureWatcherBase5pauseEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "pause", args)
  }

  return
}

// metaObject()
func (this *QFutureWatcherBase) MetaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK18QFutureWatcherBase10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "metaObject", args)
  }

  return
}

// QFutureWatcherBase(class QObject *)
func GcfreeQFutureWatcherBase(this *QFutureWatcherBase) {
  qtrt.UniverseFree(this)
}
func NewQFutureWatcherBase(args ...interface{}) *QFutureWatcherBase {
  // QFutureWatcherBase(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBaseC1EP7QObject
    // invoke: void QFutureWatcherBase(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QFutureWatcherBaseC2EP7QObject(arg0)
    this := &QFutureWatcherBase{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQFutureWatcherBase)
    return this
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "QFutureWatcherBase", args)
  }

  return nil // QFutureWatcherBase{Qclsinst:qthis}
}

// isRunning()
func (this *QFutureWatcherBase) IsRunning(args ...interface{}) (ret interface{}) {
  // isRunning()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase9isRunningEv
    // invoke: bool isRunning()
    var ret0 = C.C_ZNK18QFutureWatcherBase9isRunningEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "isRunning", args)
  }

  return
}

// isCanceled()
func (this *QFutureWatcherBase) IsCanceled(args ...interface{}) (ret interface{}) {
  // isCanceled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase10isCanceledEv
    // invoke: bool isCanceled()
    var ret0 = C.C_ZNK18QFutureWatcherBase10isCanceledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "isCanceled", args)
  }

  return
}

// isPaused()
func (this *QFutureWatcherBase) IsPaused(args ...interface{}) (ret interface{}) {
  // isPaused()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase8isPausedEv
    // invoke: bool isPaused()
    var ret0 = C.C_ZNK18QFutureWatcherBase8isPausedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "isPaused", args)
  }

  return
}

// progressValue()
func (this *QFutureWatcherBase) ProgressValue(args ...interface{}) (ret interface{}) {
  // progressValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase13progressValueEv
    // invoke: int progressValue()
    var ret0 = C.C_ZNK18QFutureWatcherBase13progressValueEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "progressValue", args)
  }

  return
}

// setPaused(_Bool)
func (this *QFutureWatcherBase) SetPaused(args ...interface{}) () {
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
    // invoke: void setPaused(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN18QFutureWatcherBase9setPausedEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "setPaused", args)
  }

  return
}

// ~QFutureWatcher()
func (this *QFutureWatcherLvoidG) Free(args ...interface{}) () {
  // ~QFutureWatcher()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QFutureWatcherIvED0Ev
    // invoke: void ~QFutureWatcher()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN14QFutureWatcherIvED2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QFutureWatcher<void>", "~QFutureWatcher", args)
  }

  return
}

// QFutureWatcher(class QObject *)
func GcfreeQFutureWatcherLvoidG(this *QFutureWatcherLvoidG) {
  qtrt.UniverseFree(this)
}
func NewQFutureWatcherLvoidG(args ...interface{}) *QFutureWatcherLvoidG {
  // QFutureWatcher(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QFutureWatcherIvEC1EP7QObject
    // invoke: void QFutureWatcher(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QFutureWatcherIvEC2EP7QObject(arg0)
    this := &QFutureWatcherLvoidG{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQFutureWatcherLvoidG)
    return this
  default:
    qtrt.ErrorResolve("QFutureWatcher<void>", "QFutureWatcher", args)
  }

  return nil // QFutureWatcherLvoidG{Qclsinst:qthis}
}

// future()
func (this *QFutureWatcherLvoidG) Future(args ...interface{}) () {
  // future()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QFutureWatcherIvE6futureEv
    // invoke: QFuture<void> future()
    C.C_ZNK14QFutureWatcherIvE6futureEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFutureWatcher<void>", "future", args)
  }

  return
}

// <= body block end

