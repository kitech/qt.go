package qtgui
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QAbstractOpenGLFunctionsPrivate * QAbstractOpenGLFunctions::d_func();
extern void C_ZN24QAbstractOpenGLFunctions6d_funcEv(void* qthis); // 4
  // proto:  void QAbstractOpenGLFunctions::~QAbstractOpenGLFunctions();
extern void C_ZN24QAbstractOpenGLFunctionsD2Ev(void* qthis); // 4
  // proto:  bool QAbstractOpenGLFunctions::initializeOpenGLFunctions();
extern bool C_ZN24QAbstractOpenGLFunctions25initializeOpenGLFunctionsEv(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QAbstractOpenGLFunctions)=16
type QAbstractOpenGLFunctions struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// d_func()
func (this *QAbstractOpenGLFunctions) D_Func(args ...interface{}) () {
  // d_func()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractOpenGLFunctions6d_funcEv
    // invoke: QAbstractOpenGLFunctionsPrivate * d_func()
    C.C_ZN24QAbstractOpenGLFunctions6d_funcEv(this.Qclsinst)
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
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractOpenGLFunctionsD0Ev
    // invoke: void ~QAbstractOpenGLFunctions()
    C.C_ZN24QAbstractOpenGLFunctionsD2Ev(this.Qclsinst)
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
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractOpenGLFunctions25initializeOpenGLFunctionsEv
    // invoke: bool initializeOpenGLFunctions()
    var ret0 = C.C_ZN24QAbstractOpenGLFunctions25initializeOpenGLFunctionsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractOpenGLFunctions", "initializeOpenGLFunctions", args)
  }

  return
}

// <= body block end

