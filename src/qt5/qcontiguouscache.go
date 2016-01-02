package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtCore/qcontiguouscache.h
// dst-file: /src/core/qcontiguouscache.go
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
  // proto: static QContiguousCacheData * QContiguousCacheData::allocateData(int size, int alignment);
extern void _ZN20QContiguousCacheData12allocateDataEii(int arg0, int arg1);
  // proto: static void QContiguousCacheData::freeData(QContiguousCacheData * data);
extern void _ZN20QContiguousCacheData8freeDataEPS_(void* arg0);
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

// class sizeof(QContiguousCacheData)=1
type QContiguousCacheData struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto: static QContiguousCacheData * QContiguousCacheData::allocateData(int size, int alignment);
func (this *QContiguousCacheData) allocateData_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QContiguousCacheData", "allocateData", args)
  }

}

  // proto: static void QContiguousCacheData::freeData(QContiguousCacheData * data);
func (this *QContiguousCacheData) freeData_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QContiguousCacheData", "freeData", args)
  }

}

// <= body block end

