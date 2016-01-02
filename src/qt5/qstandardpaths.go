package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtCore/qstandardpaths.h
// dst-file: /src/core/qstandardpaths.go
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
  // proto:  void QStandardPaths::QStandardPaths();
extern void* dector_ZN14QStandardPathsC1Ev();
extern void _ZN14QStandardPathsC1Ev(void* qthis);
  // proto:  void QStandardPaths::~QStandardPaths();
extern void _ZN14QStandardPathsD0Ev(void* qthis);
  // proto: static void QStandardPaths::setTestModeEnabled(bool testMode);
extern void _ZN14QStandardPaths18setTestModeEnabledEb(bool arg0);
  // proto: static QString QStandardPaths::findExecutable(const QString & executableName, const QStringList & paths);
extern void _ZN14QStandardPaths14findExecutableERK7QStringRK11QStringList(void* arg0, void* arg1);
  // proto: static void QStandardPaths::enableTestMode(bool testMode);
extern void _ZN14QStandardPaths14enableTestModeEb(bool arg0);
  // proto: static bool QStandardPaths::isTestModeEnabled();
extern void _ZN14QStandardPaths17isTestModeEnabledEv();
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

// class sizeof(QStandardPaths)=1
type QStandardPaths struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QStandardPaths::QStandardPaths();
func NewQStandardPaths(args ...interface{}) QStandardPaths {
  return QStandardPaths{}
}

  // proto:  void QStandardPaths::~QStandardPaths();
func (this *QStandardPaths) FreeQStandardPaths(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStandardPaths", "~QStandardPaths", args)
  }

}

  // proto: static void QStandardPaths::setTestModeEnabled(bool testMode);
func (this *QStandardPaths) setTestModeEnabled_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStandardPaths", "setTestModeEnabled", args)
  }

}

  // proto: static QString QStandardPaths::findExecutable(const QString & executableName, const QStringList & paths);
func (this *QStandardPaths) findExecutable_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStandardPaths", "findExecutable", args)
  }

}

  // proto: static void QStandardPaths::enableTestMode(bool testMode);
func (this *QStandardPaths) enableTestMode_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStandardPaths", "enableTestMode", args)
  }

}

  // proto: static bool QStandardPaths::isTestModeEnabled();
func (this *QStandardPaths) isTestModeEnabled_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStandardPaths", "isTestModeEnabled", args)
  }

}

// <= body block end

