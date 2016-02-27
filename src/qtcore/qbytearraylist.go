package qtcore
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtCore/qbytearraylist.h
// dst-file: /src/core/qbytearraylist.go
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
  // proto:  QByteArray QListSpecialMethods<QByteArray>::join();
extern void* C_ZNK19QListSpecialMethodsI10QByteArrayE4joinEv(void* qthis); // 2
  // proto:  QByteArray QListSpecialMethods<QByteArray>::join(const QByteArray & sep);
extern void* C_ZNK19QListSpecialMethodsI10QByteArrayE4joinERKS0_(void* qthis, void* arg0); // 2
  // proto:  QByteArray QListSpecialMethods<QByteArray>::join(char sep);
extern void* C_ZNK19QListSpecialMethodsI10QByteArrayE4joinEc(void* qthis, unsigned char arg0); // 2
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

// class sizeof(QListSpecialMethods<QByteArray>)=1
type QListSpecialMethodsLQByteArrayG struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// join()
func (this *QListSpecialMethodsLQByteArrayG) Join(args ...interface{}) (ret interface{}) {
  // join()
  // join(const class QByteArray &)
  // join(char)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(false) // "char"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QListSpecialMethodsI10QByteArrayE4joinEv
    // invoke: QByteArray join()
    var ret0 = C.C_ZNK19QListSpecialMethodsI10QByteArrayE4joinEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK19QListSpecialMethodsI10QByteArrayE4joinERKS0_
    // invoke: QByteArray join(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QListSpecialMethodsI10QByteArrayE4joinERKS0_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK19QListSpecialMethodsI10QByteArrayE4joinEc
    // invoke: QByteArray join(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QListSpecialMethodsI10QByteArrayE4joinEc(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListSpecialMethods<QByteArray>", "join", args)
  }

  return
}

// <= body block end

