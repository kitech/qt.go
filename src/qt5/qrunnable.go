package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
// src-file: /QtCore/qrunnable.h
// dst-file: /src/core/qrunnable.go
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
  // proto:  void QRunnable::~QRunnable();
extern void _ZN9QRunnableD0Ev(void* qthis);
  // proto:  void QRunnable::setAutoDelete(bool _autoDelete);
extern void demth_ZN9QRunnable13setAutoDeleteEb(void* qthis, bool arg0);
  // proto:  void QRunnable::QRunnable();
extern void* dector_ZN9QRunnableC1Ev();
extern void _ZN9QRunnableC1Ev(void* qthis);
  // proto:  bool QRunnable::autoDelete();
extern void demth_ZNK9QRunnable10autoDeleteEv(void* qthis);
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

// class sizeof(QRunnable)=16
type QRunnable struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  void QRunnable::~QRunnable();
func (this *QRunnable) FreeQRunnable(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QRunnable", "~QRunnable", args)
  }

}

  // proto:  void QRunnable::setAutoDelete(bool _autoDelete);
func (this *QRunnable) setAutoDelete(args ...interface{}) () {
  // setAutoDelete(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QRunnable13setAutoDeleteEb
    // invoke: void setAutoDelete(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.demth_ZN9QRunnable13setAutoDeleteEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRunnable", "setAutoDelete", args)
  }

}

  // proto:  void QRunnable::QRunnable();
func NewQRunnable(args ...interface{}) QRunnable {
  return QRunnable{}
}

  // proto:  bool QRunnable::autoDelete();
func (this *QRunnable) autoDelete(args ...interface{}) () {
  // autoDelete()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QRunnable10autoDeleteEv
    // invoke: bool autoDelete()
    C.demth_ZNK9QRunnable10autoDeleteEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRunnable", "autoDelete", args)
  }

}

// <= body block end

