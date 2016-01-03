package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
// src-file: /QtGui/qiconengine.h
// dst-file: /src/gui/qiconengine.go
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
  // proto:  bool QIconEngine::read(QDataStream & in);
extern void _ZN11QIconEngine4readER11QDataStream(void* qthis, void* arg0);
  // proto:  QString QIconEngine::iconName();
extern void _ZNK11QIconEngine8iconNameEv(void* qthis);
  // proto:  bool QIconEngine::write(QDataStream & out);
extern void _ZNK11QIconEngine5writeER11QDataStream(void* qthis, void* arg0);
  // proto:  void QIconEngine::virtual_hook(int id, void * data);
extern void _ZN11QIconEngine12virtual_hookEiPv(void* qthis, int32_t arg0, void* arg1);
  // proto:  QString QIconEngine::key();
extern void _ZNK11QIconEngine3keyEv(void* qthis);
  // proto:  void QIconEngine::~QIconEngine();
extern void _ZN11QIconEngineD0Ev(void* qthis);
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

// class sizeof(QIconEngine)=8
type QIconEngine struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  bool QIconEngine::read(QDataStream & in);
func (this *QIconEngine) read(args ...interface{}) () {
  // read(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDataStream{}) // "QDataStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QIconEngine4readER11QDataStream
    // invoke: bool read(class QDataStream &)
    var arg0 = args[0].(QDataStream).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QIconEngine4readER11QDataStream(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIconEngine", "read", args)
  }

}

  // proto:  QString QIconEngine::iconName();
func (this *QIconEngine) iconName(args ...interface{}) () {
  // iconName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QIconEngine8iconNameEv
    // invoke: QString iconName()
    C._ZNK11QIconEngine8iconNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIconEngine", "iconName", args)
  }

}

  // proto:  bool QIconEngine::write(QDataStream & out);
func (this *QIconEngine) write(args ...interface{}) () {
  // write(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDataStream{}) // "QDataStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QIconEngine5writeER11QDataStream
    // invoke: bool write(class QDataStream &)
    var arg0 = args[0].(QDataStream).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK11QIconEngine5writeER11QDataStream(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIconEngine", "write", args)
  }

}

  // proto:  void QIconEngine::virtual_hook(int id, void * data);
func (this *QIconEngine) virtual_hook(args ...interface{}) () {
  // virtual_hook(int, void *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.VoidpTy() // "void *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QIconEngine12virtual_hookEiPv
    // invoke: void virtual_hook(int, void *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    C._ZN11QIconEngine12virtual_hookEiPv(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QIconEngine", "virtual_hook", args)
  }

}

  // proto:  QString QIconEngine::key();
func (this *QIconEngine) key(args ...interface{}) () {
  // key()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QIconEngine3keyEv
    // invoke: QString key()
    C._ZNK11QIconEngine3keyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIconEngine", "key", args)
  }

}

  // proto:  void QIconEngine::~QIconEngine();
func (this *QIconEngine) FreeQIconEngine(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QIconEngine", "~QIconEngine", args)
  }

}

// <= body block end

