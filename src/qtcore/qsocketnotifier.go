package qtcore
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
  // proto:  void QSocketNotifier::setEnabled(bool );
extern void C_ZN15QSocketNotifier10setEnabledEb(void* qthis, bool arg0); // 4
  // proto:  qintptr QSocketNotifier::socket();
extern int32_t C_ZNK15QSocketNotifier6socketEv(void* qthis); // 4
  // proto:  void QSocketNotifier::~QSocketNotifier();
extern void C_ZN15QSocketNotifierD2Ev(void* qthis); // 4
  // proto:  bool QSocketNotifier::isEnabled();
extern bool C_ZNK15QSocketNotifier9isEnabledEv(void* qthis); // 4
  // proto:  const QMetaObject * QSocketNotifier::metaObject();
extern void C_ZNK15QSocketNotifier10metaObjectEv(void* qthis); // 4
  // proto:  QSocketNotifier::Type QSocketNotifier::type();
extern void C_ZNK15QSocketNotifier4typeEv(void* qthis); // 4
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

// class sizeof(QSocketNotifier)=1
type QSocketNotifier struct {
  /*qbase*/ QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _activated QSocketNotifier_activated_signal;
}

// setEnabled(_Bool)
func (this *QSocketNotifier) SetEnabled(args ...interface{}) () {
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
    C.C_ZN15QSocketNotifier10setEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSocketNotifier", "setEnabled", args)
  }

  return
}

// socket()
func (this *QSocketNotifier) Socket(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QSocketNotifier6socketEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "qintptr"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSocketNotifier", "socket", args)
  }

  return
}

// ~QSocketNotifier()
func (this *QSocketNotifier) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN15QSocketNotifierD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QSocketNotifier", "~QSocketNotifier", args)
  }

  return
}

// isEnabled()
func (this *QSocketNotifier) IsEnabled(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QSocketNotifier9isEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSocketNotifier", "isEnabled", args)
  }

  return
}

// metaObject()
func (this *QSocketNotifier) MetaObject(args ...interface{}) () {
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
    C.C_ZNK15QSocketNotifier10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSocketNotifier", "metaObject", args)
  }

  return
}

// type()
func (this *QSocketNotifier) Type_(args ...interface{}) () {
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
    C.C_ZNK15QSocketNotifier4typeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSocketNotifier", "type", args)
  }

  return
}

// <= body block end

