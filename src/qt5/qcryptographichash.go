package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtCore/qcryptographichash.h
// dst-file: /src/core/qcryptographichash.go
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
  // proto:  void QCryptographicHash::reset();
extern void C_ZN18QCryptographicHash5resetEv(void* qthis); // 4
  // proto:  void QCryptographicHash::addData(const char * data, int length);
extern void C_ZN18QCryptographicHash7addDataEPKci(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  bool QCryptographicHash::addData(QIODevice * device);
extern bool C_ZN18QCryptographicHash7addDataEP9QIODevice(void* qthis, void* arg0); // 4
  // proto:  void QCryptographicHash::addData(const QByteArray & data);
extern void C_ZN18QCryptographicHash7addDataERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  QByteArray QCryptographicHash::result();
extern void* C_ZNK18QCryptographicHash6resultEv(void* qthis); // 4
  // proto:  void QCryptographicHash::~QCryptographicHash();
extern void C_ZN18QCryptographicHashD2Ev(void* qthis); // 4
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

// class sizeof(QCryptographicHash)=8
type QCryptographicHash struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// reset()
func (this *QCryptographicHash) Reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCryptographicHash5resetEv
    // invoke: void reset()
    C.C_ZN18QCryptographicHash5resetEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCryptographicHash", "reset", args)
  }

  return
}

// addData(const char *, int)
func (this *QCryptographicHash) Adddata(args ...interface{}) () {
  // addData(const char *, int)
  // addData(class QIODevice *)
  // addData(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCryptographicHash7addDataEPKci
    // invoke: void addData(const char *, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN18QCryptographicHash7addDataEPKci(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN18QCryptographicHash7addDataEP9QIODevice
    // invoke: bool addData(class QIODevice *)
    var arg0 = args[0].(*QIODevice).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN18QCryptographicHash7addDataEP9QIODevice(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
  case 2:
    // invoke: _ZN18QCryptographicHash7addDataERK10QByteArray
    // invoke: void addData(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QCryptographicHash7addDataERK10QByteArray(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCryptographicHash", "addData", args)
  }

  return
}

// result()
func (this *QCryptographicHash) Result(args ...interface{}) (ret interface{}) {
  // result()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QCryptographicHash6resultEv
    // invoke: QByteArray result()
    var ret0 = C.C_ZNK18QCryptographicHash6resultEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCryptographicHash", "result", args)
  }

  return
}

// ~QCryptographicHash()
func (this *QCryptographicHash) Freeqcryptographichash(args ...interface{}) () {
  // ~QCryptographicHash()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCryptographicHashD0Ev
    // invoke: void ~QCryptographicHash()
    C.C_ZN18QCryptographicHashD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCryptographicHash", "~QCryptographicHash", args)
  }

  return
}

// <= body block end

