package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtCore/qsocketnotifier.h
// dst-file: /src/core/qsocketnotifier.go
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
  // proto:  void QSocketNotifier::QSocketNotifier(const QSocketNotifier & );
extern void* dector_ZN15QSocketNotifierC1ERKS_(void* arg0);
extern void _ZN15QSocketNotifierC1ERKS_(void* qthis, void* arg0);
  // proto:  qintptr QSocketNotifier::socket();
extern void _ZNK15QSocketNotifier6socketEv(void* qthis);
  // proto:  bool QSocketNotifier::isEnabled();
extern void _ZNK15QSocketNotifier9isEnabledEv(void* qthis);
  // proto:  void QSocketNotifier::setEnabled(bool );
extern void _ZN15QSocketNotifier10setEnabledEb(void* qthis, bool arg0);
  // proto:  const QMetaObject * QSocketNotifier::metaObject();
extern void _ZNK15QSocketNotifier10metaObjectEv(void* qthis);
  // proto:  void QSocketNotifier::~QSocketNotifier();
extern void _ZN15QSocketNotifierD0Ev(void* qthis);
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

// class sizeof(QSocketNotifier)=1
type QSocketNotifier struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _activated QSocketNotifier_activated_signal;
}

  // proto:  void QSocketNotifier::QSocketNotifier(const QSocketNotifier & );
func NewQSocketNotifier(args ...interface{}) QSocketNotifier {
  return QSocketNotifier{}
}

  // proto:  qintptr QSocketNotifier::socket();
func (this *QSocketNotifier) socket(args ...interface{}) () {
  // socket()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSocketNotifier6socketEv
    // invoke: qintptr socket()
    C._ZNK15QSocketNotifier6socketEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSocketNotifier", "socket", args)
  }

}

  // proto:  bool QSocketNotifier::isEnabled();
func (this *QSocketNotifier) isEnabled(args ...interface{}) () {
  // isEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSocketNotifier9isEnabledEv
    // invoke: bool isEnabled()
    C._ZNK15QSocketNotifier9isEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSocketNotifier", "isEnabled", args)
  }

}

  // proto:  void QSocketNotifier::setEnabled(bool );
func (this *QSocketNotifier) setEnabled(args ...interface{}) () {
  // setEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSocketNotifier10setEnabledEb
    // invoke: void setEnabled(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN15QSocketNotifier10setEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSocketNotifier", "setEnabled", args)
  }

}

  // proto:  const QMetaObject * QSocketNotifier::metaObject();
func (this *QSocketNotifier) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSocketNotifier10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK15QSocketNotifier10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSocketNotifier", "metaObject", args)
  }

}

  // proto:  void QSocketNotifier::~QSocketNotifier();
func (this *QSocketNotifier) FreeQSocketNotifier(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSocketNotifier", "~QSocketNotifier", args)
  }

}

// <= body block end

