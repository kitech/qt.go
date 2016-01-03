package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtCore/qscopedpointer.h
// dst-file: /src/core/qscopedpointer.go
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
  // proto: static void QScopedPointerPodDeleter::cleanup(void * pointer);
extern void demth_ZN24QScopedPointerPodDeleter7cleanupEPv(void* arg0);
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

// class sizeof(QScopedPointerPodDeleter)=1
type QScopedPointerPodDeleter struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto: static void QScopedPointerPodDeleter::cleanup(void * pointer);
func (this *QScopedPointerPodDeleter) cleanup_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QScopedPointerPodDeleter", "cleanup", args)
  }

}

// <= body block end

