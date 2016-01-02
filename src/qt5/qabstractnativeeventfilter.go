package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
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
  // proto:  bool QAbstractNativeEventFilter::nativeEventFilter(const QByteArray & eventType, void * message, long * result);
extern void _ZN26QAbstractNativeEventFilter17nativeEventFilterERK10QByteArrayPvPl(void* qthis, void* arg0, void* arg1, long* arg2);
  // proto:  void QAbstractNativeEventFilter::QAbstractNativeEventFilter();
extern void* dector_ZN26QAbstractNativeEventFilterC1Ev();
extern void _ZN26QAbstractNativeEventFilterC1Ev(void* qthis);
  // proto:  void QAbstractNativeEventFilter::QAbstractNativeEventFilter(const QAbstractNativeEventFilter & );
extern void* dector_ZN26QAbstractNativeEventFilterC1ERKS_(void* arg0);
extern void _ZN26QAbstractNativeEventFilterC1ERKS_(void* qthis, void* arg0);
  // proto:  void QAbstractNativeEventFilter::~QAbstractNativeEventFilter();
extern void _ZN26QAbstractNativeEventFilterD0Ev(void* qthis);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  bool QAbstractNativeEventFilter::nativeEventFilter(const QByteArray & eventType, void * message, long * result);
func (this *QAbstractNativeEventFilter) nativeEventFilter(args ...interface{}) () {
  // nativeEventFilter(const class QByteArray &, void *, long *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[0][1] = qtrt.VoidpTy() // "void *"
  vtys[0][2] = qtrt.Int32Ty(true) // "long *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAbstractNativeEventFilter17nativeEventFilterERK10QByteArrayPvPl
    // invoke: bool nativeEventFilter(const class QByteArray &, void *, long *)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    C._ZN26QAbstractNativeEventFilter17nativeEventFilterERK10QByteArrayPvPl(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QAbstractNativeEventFilter", "nativeEventFilter", args)
  }

}

  // proto:  void QAbstractNativeEventFilter::QAbstractNativeEventFilter();
func NewQAbstractNativeEventFilter(args ...interface{}) QAbstractNativeEventFilter {
  return QAbstractNativeEventFilter{}
}

  // proto:  void QAbstractNativeEventFilter::~QAbstractNativeEventFilter();
func (this *QAbstractNativeEventFilter) FreeQAbstractNativeEventFilter(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractNativeEventFilter", "~QAbstractNativeEventFilter", args)
  }

}

// <= body block end

