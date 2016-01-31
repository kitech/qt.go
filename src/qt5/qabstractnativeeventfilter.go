package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtCore/qabstractnativeeventfilter.h
// dst-file: /src/core/qabstractnativeeventfilter.go
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
  // proto:  void QAbstractNativeEventFilter::~QAbstractNativeEventFilter();
extern void C_ZN26QAbstractNativeEventFilterD2Ev(void* qthis); // 4
  // proto:  void QAbstractNativeEventFilter::QAbstractNativeEventFilter();
extern void C_ZN26QAbstractNativeEventFilterC2Ev(void* qthis); // 3
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

// class sizeof(QAbstractNativeEventFilter)=16
type QAbstractNativeEventFilter struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// ~QAbstractNativeEventFilter()
func (this *QAbstractNativeEventFilter) FreeQAbstractNativeEventFilter(args ...interface{}) () {
  // ~QAbstractNativeEventFilter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAbstractNativeEventFilterD0Ev
    // invoke: void ~QAbstractNativeEventFilter()
    C.C_ZN26QAbstractNativeEventFilterD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractNativeEventFilter", "~QAbstractNativeEventFilter", args)
  }

}

// QAbstractNativeEventFilter()
func NewQAbstractNativeEventFilter(args ...interface{}) QAbstractNativeEventFilter {
  // QAbstractNativeEventFilter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAbstractNativeEventFilterC1Ev
    // invoke: void QAbstractNativeEventFilter()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN26QAbstractNativeEventFilterC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QAbstractNativeEventFilter", "QAbstractNativeEventFilter", args)
  }

  return QAbstractNativeEventFilter{}
}

// <= body block end

