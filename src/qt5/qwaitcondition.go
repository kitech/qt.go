package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtCore/qwaitcondition.h
// dst-file: /src/core/qwaitcondition.go
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
  // proto:  bool QWaitCondition::wait(QReadWriteLock * lockedReadWriteLock, unsigned long time);
extern void _ZN14QWaitCondition4waitEP14QReadWriteLockm(void* qthis, void* arg0, unsigned long arg1);
  // proto:  bool QWaitCondition::wait(QMutex * lockedMutex, unsigned long time);
extern void _ZN14QWaitCondition4waitEP6QMutexm(void* qthis, void* arg0, unsigned long arg1);
  // proto:  void QWaitCondition::wakeAll();
extern void _ZN14QWaitCondition7wakeAllEv(void* qthis);
  // proto:  void QWaitCondition::wakeOne();
extern void _ZN14QWaitCondition7wakeOneEv(void* qthis);
  // proto:  void QWaitCondition::QWaitCondition(const QWaitCondition & );
extern void* dector_ZN14QWaitConditionC1ERKS_(void* arg0);
extern void _ZN14QWaitConditionC1ERKS_(void* qthis, void* arg0);
  // proto:  void QWaitCondition::~QWaitCondition();
extern void _ZN14QWaitConditionD0Ev(void* qthis);
  // proto:  void QWaitCondition::QWaitCondition();
extern void* dector_ZN14QWaitConditionC1Ev();
extern void _ZN14QWaitConditionC1Ev(void* qthis);
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

// class sizeof(QWaitCondition)=8
type QWaitCondition struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  bool QWaitCondition::wait(QReadWriteLock * lockedReadWriteLock, unsigned long time);
func (this *QWaitCondition) wait(args ...interface{}) () {
  // wait(class QReadWriteLock *, unsigned long)
  // wait(class QMutex *, unsigned long)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QReadWriteLock{}) // "QReadWriteLock *"
  vtys[0][1] = qtrt.Int32Ty(false) // "unsigned long"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QMutex{}) // "QMutex *"
  vtys[1][1] = qtrt.Int32Ty(false) // "unsigned long"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QWaitCondition4waitEP14QReadWriteLockm
    var arg0 = args[0].(QReadWriteLock).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN14QWaitCondition4waitEP6QMutexm
    var arg0 = args[0].(QMutex).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QWaitCondition", "wait", args)
  }

}

  // proto:  void QWaitCondition::wakeAll();
func (this *QWaitCondition) wakeAll(args ...interface{}) () {
  // wakeAll()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QWaitCondition7wakeAllEv
  default:
    qtrt.ErrorResolve("QWaitCondition", "wakeAll", args)
  }

}

  // proto:  void QWaitCondition::wakeOne();
func (this *QWaitCondition) wakeOne(args ...interface{}) () {
  // wakeOne()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QWaitCondition7wakeOneEv
  default:
    qtrt.ErrorResolve("QWaitCondition", "wakeOne", args)
  }

}

  // proto:  void QWaitCondition::QWaitCondition(const QWaitCondition & );
func NewQWaitCondition(args ...interface{}) QWaitCondition {
  return QWaitCondition{}
}

  // proto:  void QWaitCondition::~QWaitCondition();
func (this *QWaitCondition) FreeQWaitCondition(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWaitCondition", "~QWaitCondition", args)
  }

}

// <= body block end

