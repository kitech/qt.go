package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtCore/qshareddata.h
// dst-file: /src/core/qshareddata.go
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
  // proto:  void QSharedData::QSharedData();
extern void* C_ZN11QSharedDataC2Ev(); // 1
  // proto:  void QSharedData::QSharedData(const QSharedData & );
extern void* C_ZN11QSharedDataC2ERKS_(void* arg0); // 1
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

// class sizeof(QSharedData)=1
type QSharedData struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// QSharedData()
func NewQSharedData(args ...interface{}) *QSharedData {
  // QSharedData()
  // QSharedData(const class QSharedData &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QSharedData{}) // "const QSharedData &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QSharedDataC1Ev
    // invoke: void QSharedData()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QSharedDataC2Ev()
    return &QSharedData{Qclsinst:qthis}
  case 1:
    // invoke: _ZN11QSharedDataC1ERKS_
    // invoke: void QSharedData(const class QSharedData &)
    var arg0 = args[0].(*QSharedData).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QSharedDataC2ERKS_(arg0)
    return &QSharedData{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QSharedData", "QSharedData", args)
  }

  return nil // QSharedData{Qclsinst:qthis}
}

// <= body block end

