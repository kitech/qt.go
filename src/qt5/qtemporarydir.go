package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtCore/qtemporarydir.h
// dst-file: /src/core/qtemporarydir.go
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
  // proto:  void QTemporaryDir::~QTemporaryDir();
extern void C_ZN13QTemporaryDirD2Ev(void* qthis); // 4
  // proto:  void QTemporaryDir::setAutoRemove(bool b);
extern void C_ZN13QTemporaryDir13setAutoRemoveEb(void* qthis, bool arg0); // 4
  // proto:  void QTemporaryDir::QTemporaryDir();
extern void* C_ZN13QTemporaryDirC2Ev(); // 3
  // proto:  void QTemporaryDir::QTemporaryDir(const QString & templateName);
extern void* C_ZN13QTemporaryDirC2ERK7QString(void* arg0); // 3
  // proto:  bool QTemporaryDir::isValid();
extern void C_ZNK13QTemporaryDir7isValidEv(void* qthis); // 4
  // proto:  bool QTemporaryDir::remove();
extern void C_ZN13QTemporaryDir6removeEv(void* qthis); // 4
  // proto:  bool QTemporaryDir::autoRemove();
extern void C_ZNK13QTemporaryDir10autoRemoveEv(void* qthis); // 4
  // proto:  QString QTemporaryDir::path();
extern void C_ZNK13QTemporaryDir4pathEv(void* qthis); // 4
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

// class sizeof(QTemporaryDir)=1
type QTemporaryDir struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// ~QTemporaryDir()
func (this *QTemporaryDir) FreeQTemporaryDir(args ...interface{}) () {
  // ~QTemporaryDir()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTemporaryDirD0Ev
    // invoke: void ~QTemporaryDir()
    C.C_ZN13QTemporaryDirD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTemporaryDir", "~QTemporaryDir", args)
  }

}

// setAutoRemove(_Bool)
func (this *QTemporaryDir) setAutoRemove(args ...interface{}) () {
  // setAutoRemove(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTemporaryDir13setAutoRemoveEb
    // invoke: void setAutoRemove(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QTemporaryDir13setAutoRemoveEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTemporaryDir", "setAutoRemove", args)
  }

}

// QTemporaryDir()
func NewQTemporaryDir(args ...interface{}) *QTemporaryDir {
  // QTemporaryDir()
  // QTemporaryDir(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTemporaryDirC1Ev
    // invoke: void QTemporaryDir()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QTemporaryDirC2Ev()
    return &QTemporaryDir{qclsinst:qthis}
  case 1:
    // invoke: _ZN13QTemporaryDirC1ERK7QString
    // invoke: void QTemporaryDir(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QTemporaryDirC2ERK7QString(arg0)
    return &QTemporaryDir{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTemporaryDir", "QTemporaryDir", args)
  }

  return nil // QTemporaryDir{qclsinst:qthis}
}

// isValid()
func (this *QTemporaryDir) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTemporaryDir7isValidEv
    // invoke: bool isValid()
    var ret = C.C_ZNK13QTemporaryDir7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTemporaryDir", "isValid", args)
  }

}

// remove()
func (this *QTemporaryDir) remove(args ...interface{}) () {
  // remove()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTemporaryDir6removeEv
    // invoke: bool remove()
    var ret = C.C_ZN13QTemporaryDir6removeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTemporaryDir", "remove", args)
  }

}

// autoRemove()
func (this *QTemporaryDir) autoRemove(args ...interface{}) () {
  // autoRemove()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTemporaryDir10autoRemoveEv
    // invoke: bool autoRemove()
    var ret = C.C_ZNK13QTemporaryDir10autoRemoveEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTemporaryDir", "autoRemove", args)
  }

}

// path()
func (this *QTemporaryDir) path(args ...interface{}) () {
  // path()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTemporaryDir4pathEv
    // invoke: QString path()
    var ret = C.C_ZNK13QTemporaryDir4pathEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTemporaryDir", "path", args)
  }

}

// <= body block end

