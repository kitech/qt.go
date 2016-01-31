package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto: static void QStandardPaths::setTestModeEnabled(bool testMode);
extern void C_ZN14QStandardPaths18setTestModeEnabledEb(bool arg0); // 4
  // proto: static void QStandardPaths::enableTestMode(bool testMode);
extern void C_ZN14QStandardPaths14enableTestModeEb(bool arg0); // 4
  // proto: static QString QStandardPaths::findExecutable(const QString & executableName, const QStringList & paths);
extern void C_ZN14QStandardPaths14findExecutableERK7QStringRK11QStringList(void* arg0, void* arg1); // 4
  // proto: static bool QStandardPaths::isTestModeEnabled();
extern void C_ZN14QStandardPaths17isTestModeEnabledEv(); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// setTestModeEnabled(_Bool)
func (this *QStandardPaths) setTestModeEnabled_s(args ...interface{}) () {
  // setTestModeEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStandardPaths18setTestModeEnabledEb
    // invoke: void setTestModeEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN14QStandardPaths18setTestModeEnabledEb(arg0)
  default:
    qtrt.ErrorResolve("QStandardPaths", "setTestModeEnabled", args)
  }

}

// enableTestMode(_Bool)
func (this *QStandardPaths) enableTestMode_s(args ...interface{}) () {
  // enableTestMode(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStandardPaths14enableTestModeEb
    // invoke: void enableTestMode(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN14QStandardPaths14enableTestModeEb(arg0)
  default:
    qtrt.ErrorResolve("QStandardPaths", "enableTestMode", args)
  }

}

// findExecutable(const class QString &, const class QStringList &)
func (this *QStandardPaths) findExecutable_s(args ...interface{}) () {
  // findExecutable(const class QString &, const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStandardPaths14findExecutableERK7QStringRK11QStringList
    // invoke: QString findExecutable(const class QString &, const class QStringList &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStringList).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN14QStandardPaths14findExecutableERK7QStringRK11QStringList(arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QStandardPaths", "findExecutable", args)
  }

}

// isTestModeEnabled()
func (this *QStandardPaths) isTestModeEnabled_s(args ...interface{}) () {
  // isTestModeEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStandardPaths17isTestModeEnabledEv
    // invoke: bool isTestModeEnabled()
    var ret = C.C_ZN14QStandardPaths17isTestModeEnabledEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QStandardPaths", "isTestModeEnabled", args)
  }

}

// <= body block end

