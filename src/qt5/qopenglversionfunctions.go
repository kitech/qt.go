package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtGui/qopenglversionfunctions.h
// dst-file: /src/gui/qopenglversionfunctions.go
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
  // proto:  void QOpenGLVersionFunctionsBackend::QOpenGLVersionFunctionsBackend(QOpenGLContext * ctx);
extern void* C_ZN30QOpenGLVersionFunctionsBackendC2EP14QOpenGLContext(void* arg0); // 1
  // proto:  QAbstractOpenGLFunctionsPrivate * QAbstractOpenGLFunctions::d_func();
extern void C_ZN24QAbstractOpenGLFunctions6d_funcEv(void* qthis); // 4
  // proto:  void QAbstractOpenGLFunctions::~QAbstractOpenGLFunctions();
extern void C_ZN24QAbstractOpenGLFunctionsD2Ev(void* qthis); // 4
  // proto:  bool QAbstractOpenGLFunctions::initializeOpenGLFunctions();
extern bool C_ZN24QAbstractOpenGLFunctions25initializeOpenGLFunctionsEv(void* qthis); // 4
  // proto:  void QOpenGLVersionStatus::QOpenGLVersionStatus();
extern void* C_ZN20QOpenGLVersionStatusC2Ev(); // 1
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

// class sizeof(QOpenGLVersionFunctionsBackend)=1
type QOpenGLVersionFunctionsBackend struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAbstractOpenGLFunctions)=16
type QAbstractOpenGLFunctions struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QOpenGLVersionStatus)=1
type QOpenGLVersionStatus struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// QOpenGLVersionFunctionsBackend(class QOpenGLContext *)
func NewQOpenGLVersionFunctionsBackend(args ...interface{}) *QOpenGLVersionFunctionsBackend {
  // QOpenGLVersionFunctionsBackend(class QOpenGLContext *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QOpenGLContext{}) // "QOpenGLContext *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN30QOpenGLVersionFunctionsBackendC1EP14QOpenGLContext
    // invoke: void QOpenGLVersionFunctionsBackend(class QOpenGLContext *)
    var arg0 = args[0].(QOpenGLContext).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN30QOpenGLVersionFunctionsBackendC2EP14QOpenGLContext(arg0)
    return &QOpenGLVersionFunctionsBackend{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QOpenGLVersionFunctionsBackend", "QOpenGLVersionFunctionsBackend", args)
  }

  return nil // QOpenGLVersionFunctionsBackend{qclsinst:qthis}
}

// d_func()
func (this *QAbstractOpenGLFunctions) D_Func(args ...interface{}) () {
  // d_func()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractOpenGLFunctions6d_funcEv
    // invoke: QAbstractOpenGLFunctionsPrivate * d_func()
    C.C_ZN24QAbstractOpenGLFunctions6d_funcEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractOpenGLFunctions", "d_func", args)
  }

  return
}

// ~QAbstractOpenGLFunctions()
func (this *QAbstractOpenGLFunctions) Freeqabstractopenglfunctions(args ...interface{}) () {
  // ~QAbstractOpenGLFunctions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractOpenGLFunctionsD0Ev
    // invoke: void ~QAbstractOpenGLFunctions()
    C.C_ZN24QAbstractOpenGLFunctionsD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractOpenGLFunctions", "~QAbstractOpenGLFunctions", args)
  }

  return
}

// initializeOpenGLFunctions()
func (this *QAbstractOpenGLFunctions) Initializeopenglfunctions(args ...interface{}) (ret interface{}) {
  // initializeOpenGLFunctions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractOpenGLFunctions25initializeOpenGLFunctionsEv
    // invoke: bool initializeOpenGLFunctions()
    var ret0 = C.C_ZN24QAbstractOpenGLFunctions25initializeOpenGLFunctionsEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAbstractOpenGLFunctions", "initializeOpenGLFunctions", args)
  }

  return
}

// QOpenGLVersionStatus()
func NewQOpenGLVersionStatus(args ...interface{}) *QOpenGLVersionStatus {
  // QOpenGLVersionStatus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLVersionStatusC1Ev
    // invoke: void QOpenGLVersionStatus()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QOpenGLVersionStatusC2Ev()
    return &QOpenGLVersionStatus{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QOpenGLVersionStatus", "QOpenGLVersionStatus", args)
  }

  return nil // QOpenGLVersionStatus{qclsinst:qthis}
}

// <= body block end

