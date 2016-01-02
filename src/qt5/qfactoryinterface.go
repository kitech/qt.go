package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtCore/qfactoryinterface.h
// dst-file: /src/core/qfactoryinterface.go
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
  // proto:  void QFactoryInterface::~QFactoryInterface();
extern void _ZN17QFactoryInterfaceD0Ev(void* qthis);
  // proto:  QStringList QFactoryInterface::keys();
extern void _ZNK17QFactoryInterface4keysEv(void* qthis);
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

// class sizeof(QFactoryInterface)=8
type QFactoryInterface struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QFactoryInterface::~QFactoryInterface();
func (this *QFactoryInterface) FreeQFactoryInterface(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFactoryInterface", "~QFactoryInterface", args)
  }

}

  // proto:  QStringList QFactoryInterface::keys();
func (this *QFactoryInterface) keys(args ...interface{}) () {
  // keys()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QFactoryInterface4keysEv
  default:
    qtrt.ErrorResolve("QFactoryInterface", "keys", args)
  }

}

// <= body block end

