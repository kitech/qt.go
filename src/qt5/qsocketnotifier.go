package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QSocketNotifier::setEnabled(bool );
extern void _ZN15QSocketNotifier10setEnabledEb(void* qthis, bool arg0); // 4
  // proto:  qintptr QSocketNotifier::socket();
extern void _ZNK15QSocketNotifier6socketEv(void* qthis); // 4
  // proto:  void QSocketNotifier::~QSocketNotifier();
extern void _ZN15QSocketNotifierD2Ev(void* qthis); // 4
  // proto:  bool QSocketNotifier::isEnabled();
extern void _ZNK15QSocketNotifier9isEnabledEv(void* qthis); // 4
  // proto:  const QMetaObject * QSocketNotifier::metaObject();
extern void _ZNK15QSocketNotifier10metaObjectEv(void* qthis); // 4
  // proto:  QSocketNotifier::Type QSocketNotifier::type();
extern void _ZNK15QSocketNotifier4typeEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
//  _activated QSocketNotifier_activated_signal;
}

// setEnabled(_Bool)
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN15QSocketNotifier10setEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSocketNotifier", "setEnabled", args)
  }

}

// socket()
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

// ~QSocketNotifier()
func (this *QSocketNotifier) FreeQSocketNotifier(args ...interface{}) () {
  // ~QSocketNotifier()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSocketNotifierD0Ev
    // invoke: void ~QSocketNotifier()
    C._ZN15QSocketNotifierD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSocketNotifier", "~QSocketNotifier", args)
  }

}

// isEnabled()
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

// metaObject()
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

// type()
func (this *QSocketNotifier) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSocketNotifier4typeEv
    // invoke: QSocketNotifier::Type type()
    C._ZNK15QSocketNotifier4typeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSocketNotifier", "type", args)
  }

}

// <= body block end

