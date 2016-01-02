package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
  // proto:  void QSharedData::QSharedData();
extern void* dector_ZN11QSharedDataC1Ev();
extern void demth_ZN11QSharedDataC1Ev(void* qthis);
  // proto:  void QSharedData::QSharedData(const QSharedData & );
extern void* dector_ZN11QSharedDataC1ERKS_(void* arg0);
extern void demth_ZN11QSharedDataC1ERKS_(void* qthis, void* arg0);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QSharedData::QSharedData();
func NewQSharedData(args ...interface{}) QSharedData {
  return QSharedData{}
}

// <= body block end

