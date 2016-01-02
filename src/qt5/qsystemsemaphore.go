package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
  // proto:  QString QSystemSemaphore::key();
extern void _ZNK16QSystemSemaphore3keyEv(void* qthis);
  // proto:  bool QSystemSemaphore::release(int n);
extern void _ZN16QSystemSemaphore7releaseEi(void* qthis, int arg0);
  // proto:  QString QSystemSemaphore::errorString();
extern void _ZNK16QSystemSemaphore11errorStringEv(void* qthis);
  // proto:  void QSystemSemaphore::QSystemSemaphore(const QSystemSemaphore & );
extern void* dector_ZN16QSystemSemaphoreC1ERKS_(void* arg0);
extern void _ZN16QSystemSemaphoreC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QSystemSemaphore::acquire();
extern void _ZN16QSystemSemaphore7acquireEv(void* qthis);
  // proto:  void QSystemSemaphore::~QSystemSemaphore();
extern void _ZN16QSystemSemaphoreD0Ev(void* qthis);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QString QSystemSemaphore::key();
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
  default:
    qtrt.ErrorResolve("QSystemSemaphore", "key", args)
  }

}

  // proto:  bool QSystemSemaphore::release(int n);
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
  default:
    qtrt.ErrorResolve("QSystemSemaphore", "release", args)
  }

}

  // proto:  QString QSystemSemaphore::errorString();
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
  default:
    qtrt.ErrorResolve("QSystemSemaphore", "errorString", args)
  }

}

  // proto:  void QSystemSemaphore::QSystemSemaphore(const QSystemSemaphore & );
func NewQSystemSemaphore(args ...interface{}) QSystemSemaphore {
  return QSystemSemaphore{}
}

  // proto:  bool QSystemSemaphore::acquire();
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
  default:
    qtrt.ErrorResolve("QSystemSemaphore", "acquire", args)
  }

}

  // proto:  void QSystemSemaphore::~QSystemSemaphore();
func (this *QSystemSemaphore) FreeQSystemSemaphore(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSystemSemaphore", "~QSystemSemaphore", args)
  }

}

// <= body block end

