package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtCore/qsemaphore.h
// dst-file: /src/core/qsemaphore.go
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
  // proto:  void QSemaphore::acquire(int n);
extern void _ZN10QSemaphore7acquireEi(void* qthis, int arg0);
  // proto:  void QSemaphore::release(int n);
extern void _ZN10QSemaphore7releaseEi(void* qthis, int arg0);
  // proto:  int QSemaphore::available();
extern void _ZNK10QSemaphore9availableEv(void* qthis);
  // proto:  bool QSemaphore::tryAcquire(int n, int timeout);
extern void _ZN10QSemaphore10tryAcquireEii(void* qthis, int arg0, int arg1);
  // proto:  void QSemaphore::QSemaphore(const QSemaphore & );
extern void* dector_ZN10QSemaphoreC1ERKS_(void* arg0);
extern void _ZN10QSemaphoreC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QSemaphore::tryAcquire(int n);
extern void _ZN10QSemaphore10tryAcquireEi(void* qthis, int arg0);
  // proto:  void QSemaphore::QSemaphore(int n);
extern void* dector_ZN10QSemaphoreC1Ei(int arg0);
extern void _ZN10QSemaphoreC1Ei(void* qthis, int arg0);
  // proto:  void QSemaphore::~QSemaphore();
extern void _ZN10QSemaphoreD0Ev(void* qthis);
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

// class sizeof(QSemaphore)=8
type QSemaphore struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QSemaphore::acquire(int n);
func (this *QSemaphore) acquire(args ...interface{}) () {
  // acquire(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QSemaphore7acquireEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QSemaphore", "acquire", args)
  }

}

  // proto:  void QSemaphore::release(int n);
func (this *QSemaphore) release(args ...interface{}) () {
  // release(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QSemaphore7releaseEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QSemaphore", "release", args)
  }

}

  // proto:  int QSemaphore::available();
func (this *QSemaphore) available(args ...interface{}) () {
  // available()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QSemaphore9availableEv
  default:
    qtrt.ErrorResolve("QSemaphore", "available", args)
  }

}

  // proto:  bool QSemaphore::tryAcquire(int n, int timeout);
func (this *QSemaphore) tryAcquire(args ...interface{}) () {
  // tryAcquire(int, int)
  // tryAcquire(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QSemaphore10tryAcquireEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN10QSemaphore10tryAcquireEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QSemaphore", "tryAcquire", args)
  }

}

  // proto:  void QSemaphore::QSemaphore(const QSemaphore & );
func NewQSemaphore(args ...interface{}) QSemaphore {
  return QSemaphore{}
}

  // proto:  void QSemaphore::~QSemaphore();
func (this *QSemaphore) FreeQSemaphore(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSemaphore", "~QSemaphore", args)
  }

}

// <= body block end

