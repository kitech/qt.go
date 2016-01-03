package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
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
  qclsinst unsafe.Pointer /* *C.void */;
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

