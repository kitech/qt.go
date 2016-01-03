package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
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
  // proto:  bool QTemporaryDir::remove();
extern void _ZN13QTemporaryDir6removeEv(void* qthis);
  // proto:  bool QTemporaryDir::autoRemove();
extern void _ZNK13QTemporaryDir10autoRemoveEv(void* qthis);
  // proto:  bool QTemporaryDir::isValid();
extern void _ZNK13QTemporaryDir7isValidEv(void* qthis);
  // proto:  void QTemporaryDir::setAutoRemove(bool b);
extern void _ZN13QTemporaryDir13setAutoRemoveEb(void* qthis, bool arg0);
  // proto:  void QTemporaryDir::~QTemporaryDir();
extern void _ZN13QTemporaryDirD0Ev(void* qthis);
  // proto:  void QTemporaryDir::QTemporaryDir();
extern void* dector_ZN13QTemporaryDirC1Ev();
extern void _ZN13QTemporaryDirC1Ev(void* qthis);
  // proto:  void QTemporaryDir::QTemporaryDir(const QString & templateName);
extern void* dector_ZN13QTemporaryDirC1ERK7QString(void* arg0);
extern void _ZN13QTemporaryDirC1ERK7QString(void* qthis, void* arg0);
  // proto:  QString QTemporaryDir::path();
extern void _ZNK13QTemporaryDir4pathEv(void* qthis);
  // proto:  void QTemporaryDir::QTemporaryDir(const QTemporaryDir & );
extern void* dector_ZN13QTemporaryDirC1ERKS_(void* arg0);
extern void _ZN13QTemporaryDirC1ERKS_(void* qthis, void* arg0);
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

  // proto:  bool QTemporaryDir::remove();
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
    C._ZN13QTemporaryDir6removeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTemporaryDir", "remove", args)
  }

}

  // proto:  bool QTemporaryDir::autoRemove();
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
    C._ZNK13QTemporaryDir10autoRemoveEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTemporaryDir", "autoRemove", args)
  }

}

  // proto:  bool QTemporaryDir::isValid();
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
    C._ZNK13QTemporaryDir7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTemporaryDir", "isValid", args)
  }

}

  // proto:  void QTemporaryDir::setAutoRemove(bool b);
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
    C._ZN13QTemporaryDir13setAutoRemoveEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTemporaryDir", "setAutoRemove", args)
  }

}

  // proto:  void QTemporaryDir::~QTemporaryDir();
func (this *QTemporaryDir) FreeQTemporaryDir(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTemporaryDir", "~QTemporaryDir", args)
  }

}

  // proto:  void QTemporaryDir::QTemporaryDir();
func NewQTemporaryDir(args ...interface{}) QTemporaryDir {
  return QTemporaryDir{}
}

  // proto:  QString QTemporaryDir::path();
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
    C._ZNK13QTemporaryDir4pathEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTemporaryDir", "path", args)
  }

}

// <= body block end

