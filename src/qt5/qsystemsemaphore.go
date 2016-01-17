package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtCore/qsystemsemaphore.h
// dst-file: /src/core/qsystemsemaphore.go
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
  // proto:  QString QSystemSemaphore::errorString();
extern void _ZNK16QSystemSemaphore11errorStringEv(void* qthis); // 4
  // proto:  bool QSystemSemaphore::acquire();
extern void _ZN16QSystemSemaphore7acquireEv(void* qthis); // 4
  // proto:  void QSystemSemaphore::~QSystemSemaphore();
extern void _ZN16QSystemSemaphoreD2Ev(void* qthis); // 4
  // proto:  QString QSystemSemaphore::key();
extern void _ZNK16QSystemSemaphore3keyEv(void* qthis); // 4
  // proto:  QSystemSemaphore::SystemSemaphoreError QSystemSemaphore::error();
extern void _ZNK16QSystemSemaphore5errorEv(void* qthis); // 4
  // proto:  bool QSystemSemaphore::release(int n);
extern void _ZN16QSystemSemaphore7releaseEi(void* qthis, int32_t arg0); // 4
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

// class sizeof(QSystemSemaphore)=1
type QSystemSemaphore struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// errorString()
func (this *QSystemSemaphore) errorString(args ...interface{}) () {
  // errorString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QSystemSemaphore11errorStringEv
    // invoke: QString errorString()
    C._ZNK16QSystemSemaphore11errorStringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSystemSemaphore", "errorString", args)
  }

}

// acquire()
func (this *QSystemSemaphore) acquire(args ...interface{}) () {
  // acquire()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QSystemSemaphore7acquireEv
    // invoke: bool acquire()
    C._ZN16QSystemSemaphore7acquireEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSystemSemaphore", "acquire", args)
  }

}

// ~QSystemSemaphore()
func (this *QSystemSemaphore) FreeQSystemSemaphore(args ...interface{}) () {
  // ~QSystemSemaphore()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QSystemSemaphoreD0Ev
    // invoke: void ~QSystemSemaphore()
    C._ZN16QSystemSemaphoreD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSystemSemaphore", "~QSystemSemaphore", args)
  }

}

// key()
func (this *QSystemSemaphore) key(args ...interface{}) () {
  // key()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QSystemSemaphore3keyEv
    // invoke: QString key()
    C._ZNK16QSystemSemaphore3keyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSystemSemaphore", "key", args)
  }

}

// error()
func (this *QSystemSemaphore) error(args ...interface{}) () {
  // error()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QSystemSemaphore5errorEv
    // invoke: QSystemSemaphore::SystemSemaphoreError error()
    C._ZNK16QSystemSemaphore5errorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSystemSemaphore", "error", args)
  }

}

// release(int)
func (this *QSystemSemaphore) release(args ...interface{}) () {
  // release(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QSystemSemaphore7releaseEi
    // invoke: bool release(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN16QSystemSemaphore7releaseEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSystemSemaphore", "release", args)
  }

}

// <= body block end

