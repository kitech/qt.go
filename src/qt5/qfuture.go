package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtCore/qfuture.h
// dst-file: /src/core/qfuture.go
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
  // proto:  QString QFuture<void>::progressText();
extern void C_ZNK7QFutureIvE12progressTextEv(void* qthis); // 2
  // proto:  int QFuture<void>::resultCount();
extern void C_ZNK7QFutureIvE11resultCountEv(void* qthis); // 2
  // proto:  void QFuture<void>::~QFuture();
extern void C_ZN7QFutureIvED2Ev(void* qthis); // 4
  // proto:  void QFuture<void>::resume();
extern void C_ZN7QFutureIvE6resumeEv(void* qthis); // 2
  // proto:  bool QFuture<void>::isFinished();
extern void C_ZNK7QFutureIvE10isFinishedEv(void* qthis); // 2
  // proto:  void QFuture<void>::waitForFinished();
extern void C_ZN7QFutureIvE15waitForFinishedEv(void* qthis); // 2
  // proto:  bool QFuture<void>::isCanceled();
extern void C_ZNK7QFutureIvE10isCanceledEv(void* qthis); // 2
  // proto:  int QFuture<void>::progressMaximum();
extern void C_ZNK7QFutureIvE15progressMaximumEv(void* qthis); // 2
  // proto:  int QFuture<void>::progressMinimum();
extern void C_ZNK7QFutureIvE15progressMinimumEv(void* qthis); // 2
  // proto:  void QFuture<void>::pause();
extern void C_ZN7QFutureIvE5pauseEv(void* qthis); // 2
  // proto:  bool QFuture<void>::isStarted();
extern void C_ZNK7QFutureIvE9isStartedEv(void* qthis); // 2
  // proto:  bool QFuture<void>::isPaused();
extern void C_ZNK7QFutureIvE8isPausedEv(void* qthis); // 2
  // proto:  void QFuture<void>::cancel();
extern void C_ZN7QFutureIvE6cancelEv(void* qthis); // 2
  // proto:  bool QFuture<void>::isRunning();
extern void C_ZNK7QFutureIvE9isRunningEv(void* qthis); // 2
  // proto:  void QFuture<void>::QFuture(QFutureInterfaceBase * p);
extern void* C_ZN7QFutureIvEC2EP20QFutureInterfaceBase(void* arg0); // 1
  // proto:  void QFuture<void>::QFuture();
extern void* C_ZN7QFutureIvEC2Ev(); // 1
  // proto:  int QFuture<void>::progressValue();
extern void C_ZNK7QFutureIvE13progressValueEv(void* qthis); // 2
  // proto:  void QFuture<void>::togglePaused();
extern void C_ZN7QFutureIvE12togglePausedEv(void* qthis); // 2
  // proto:  void QFuture<void>::setPaused(bool paused);
extern void C_ZN7QFutureIvE9setPausedEb(void* qthis, bool arg0); // 2
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

// class sizeof(QFuture<void>)=16
type QFutureLvoidG struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// progressText()
func (this *QFutureLvoidG) progressText(args ...interface{}) () {
  // progressText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QFutureIvE12progressTextEv
    // invoke: QString progressText()
    var ret = C.C_ZNK7QFutureIvE12progressTextEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFuture<void>", "progressText", args)
  }

}

// resultCount()
func (this *QFutureLvoidG) resultCount(args ...interface{}) () {
  // resultCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QFutureIvE11resultCountEv
    // invoke: int resultCount()
    var ret = C.C_ZNK7QFutureIvE11resultCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFuture<void>", "resultCount", args)
  }

}

// ~QFuture()
func (this *QFutureLvoidG) FreeQFutureLvoidG(args ...interface{}) () {
  // ~QFuture()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QFutureIvED0Ev
    // invoke: void ~QFuture()
    C.C_ZN7QFutureIvED2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFuture<void>", "~QFuture", args)
  }

}

// resume()
func (this *QFutureLvoidG) resume(args ...interface{}) () {
  // resume()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QFutureIvE6resumeEv
    // invoke: void resume()
    C.C_ZN7QFutureIvE6resumeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFuture<void>", "resume", args)
  }

}

// isFinished()
func (this *QFutureLvoidG) isFinished(args ...interface{}) () {
  // isFinished()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QFutureIvE10isFinishedEv
    // invoke: bool isFinished()
    var ret = C.C_ZNK7QFutureIvE10isFinishedEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFuture<void>", "isFinished", args)
  }

}

// waitForFinished()
func (this *QFutureLvoidG) waitForFinished(args ...interface{}) () {
  // waitForFinished()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QFutureIvE15waitForFinishedEv
    // invoke: void waitForFinished()
    C.C_ZN7QFutureIvE15waitForFinishedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFuture<void>", "waitForFinished", args)
  }

}

// isCanceled()
func (this *QFutureLvoidG) isCanceled(args ...interface{}) () {
  // isCanceled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QFutureIvE10isCanceledEv
    // invoke: bool isCanceled()
    var ret = C.C_ZNK7QFutureIvE10isCanceledEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFuture<void>", "isCanceled", args)
  }

}

// progressMaximum()
func (this *QFutureLvoidG) progressMaximum(args ...interface{}) () {
  // progressMaximum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QFutureIvE15progressMaximumEv
    // invoke: int progressMaximum()
    var ret = C.C_ZNK7QFutureIvE15progressMaximumEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFuture<void>", "progressMaximum", args)
  }

}

// progressMinimum()
func (this *QFutureLvoidG) progressMinimum(args ...interface{}) () {
  // progressMinimum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QFutureIvE15progressMinimumEv
    // invoke: int progressMinimum()
    var ret = C.C_ZNK7QFutureIvE15progressMinimumEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFuture<void>", "progressMinimum", args)
  }

}

// pause()
func (this *QFutureLvoidG) pause(args ...interface{}) () {
  // pause()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QFutureIvE5pauseEv
    // invoke: void pause()
    C.C_ZN7QFutureIvE5pauseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFuture<void>", "pause", args)
  }

}

// isStarted()
func (this *QFutureLvoidG) isStarted(args ...interface{}) () {
  // isStarted()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QFutureIvE9isStartedEv
    // invoke: bool isStarted()
    var ret = C.C_ZNK7QFutureIvE9isStartedEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFuture<void>", "isStarted", args)
  }

}

// isPaused()
func (this *QFutureLvoidG) isPaused(args ...interface{}) () {
  // isPaused()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QFutureIvE8isPausedEv
    // invoke: bool isPaused()
    var ret = C.C_ZNK7QFutureIvE8isPausedEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFuture<void>", "isPaused", args)
  }

}

// cancel()
func (this *QFutureLvoidG) cancel(args ...interface{}) () {
  // cancel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QFutureIvE6cancelEv
    // invoke: void cancel()
    C.C_ZN7QFutureIvE6cancelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFuture<void>", "cancel", args)
  }

}

// isRunning()
func (this *QFutureLvoidG) isRunning(args ...interface{}) () {
  // isRunning()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QFutureIvE9isRunningEv
    // invoke: bool isRunning()
    var ret = C.C_ZNK7QFutureIvE9isRunningEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFuture<void>", "isRunning", args)
  }

}

// QFuture(class QFutureInterfaceBase *)
func NewQFutureLvoidG(args ...interface{}) *QFutureLvoidG {
  // QFuture(class QFutureInterfaceBase *)
  // QFuture()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFutureInterfaceBase{}) // "QFutureInterfaceBase *"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QFutureIvEC1EP20QFutureInterfaceBase
    // invoke: void QFuture(class QFutureInterfaceBase *)
    var arg0 = args[0].(QFutureInterfaceBase).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QFutureIvEC2EP20QFutureInterfaceBase(arg0)
    return &QFutureLvoidG{qclsinst:qthis}
  case 1:
    // invoke: _ZN7QFutureIvEC1Ev
    // invoke: void QFuture()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QFutureIvEC2Ev()
    return &QFutureLvoidG{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QFuture<void>", "QFuture", args)
  }

  return nil // QFutureLvoidG{qclsinst:qthis}
}

// progressValue()
func (this *QFutureLvoidG) progressValue(args ...interface{}) () {
  // progressValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QFutureIvE13progressValueEv
    // invoke: int progressValue()
    var ret = C.C_ZNK7QFutureIvE13progressValueEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFuture<void>", "progressValue", args)
  }

}

// togglePaused()
func (this *QFutureLvoidG) togglePaused(args ...interface{}) () {
  // togglePaused()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QFutureIvE12togglePausedEv
    // invoke: void togglePaused()
    C.C_ZN7QFutureIvE12togglePausedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFuture<void>", "togglePaused", args)
  }

}

// setPaused(_Bool)
func (this *QFutureLvoidG) setPaused(args ...interface{}) () {
  // setPaused(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QFutureIvE9setPausedEb
    // invoke: void setPaused(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QFutureIvE9setPausedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFuture<void>", "setPaused", args)
  }

}

// <= body block end

