package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QString QIconEngine::iconName();
extern void* C_ZNK11QIconEngine8iconNameEv(void* qthis); // 4
  // proto:  bool QIconEngine::write(QDataStream & out);
extern bool C_ZNK11QIconEngine5writeER11QDataStream(void* qthis, void* arg0); // 4
  // proto:  bool QIconEngine::read(QDataStream & in);
extern bool C_ZN11QIconEngine4readER11QDataStream(void* qthis, void* arg0); // 4
  // proto:  void QIconEngine::virtual_hook(int id, void * data);
extern void C_ZN11QIconEngine12virtual_hookEiPv(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QString QIconEngine::key();
extern void* C_ZNK11QIconEngine3keyEv(void* qthis); // 4
  // proto:  void QIconEngine::~QIconEngine();
extern void C_ZN11QIconEngineD2Ev(void* qthis); // 4
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

// iconName()
func (this *QIconEngine) Iconname(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QIconEngine8iconNameEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QIconEngine", "iconName", args)
  }

  return
}

// write(class QDataStream &)
func (this *QIconEngine) Write(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QIconEngine5writeER11QDataStream(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QIconEngine", "write", args)
  }

  return
}

// read(class QDataStream &)
func (this *QIconEngine) Read(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN11QIconEngine4readER11QDataStream(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QIconEngine", "read", args)
  }

  return
}

// virtual_hook(int, void *)
func (this *QIconEngine) Virtual_Hook(args ...interface{}) () {
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
    C.C_ZN11QIconEngine12virtual_hookEiPv(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QIconEngine", "virtual_hook", args)
  }

  return
}

// key()
func (this *QIconEngine) Key(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QIconEngine3keyEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QIconEngine", "key", args)
  }

  return
}

// ~QIconEngine()
func (this *QIconEngine) Freeqiconengine(args ...interface{}) () {
  // ~QIconEngine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QIconEngineD0Ev
    // invoke: void ~QIconEngine()
    C.C_ZN11QIconEngineD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIconEngine", "~QIconEngine", args)
  }

  return
}

// <= body block end

