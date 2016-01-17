package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:13 2016
// src-file: /QtWidgets/qstyleplugin.h
// dst-file: /src/widgets/qstyleplugin.go
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
  // proto:  void QStylePlugin::~QStylePlugin();
extern void _ZN12QStylePluginD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QStylePlugin::metaObject();
extern void _ZNK12QStylePlugin10metaObjectEv(void* qthis); // 4
  // proto:  void QStylePlugin::QStylePlugin(QObject * parent);
extern void _ZN12QStylePluginC2EP7QObject(void* qthis, void* arg0); // 3
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

// class sizeof(QStylePlugin)=1
type QStylePlugin struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
}

// ~QStylePlugin()
func (this *QStylePlugin) FreeQStylePlugin(args ...interface{}) () {
  // ~QStylePlugin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QStylePluginD0Ev
    // invoke: void ~QStylePlugin()
    C._ZN12QStylePluginD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStylePlugin", "~QStylePlugin", args)
  }

}

// metaObject()
func (this *QStylePlugin) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QStylePlugin10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK12QStylePlugin10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStylePlugin", "metaObject", args)
  }

}

// QStylePlugin(class QObject *)
func NewQStylePlugin(args ...interface{}) QStylePlugin {
  // QStylePlugin(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QStylePluginC1EP7QObject
    // invoke: void QStylePlugin(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN12QStylePluginC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QStylePlugin", "QStylePlugin", args)
  }

  return QStylePlugin{}
}

// <= body block end

