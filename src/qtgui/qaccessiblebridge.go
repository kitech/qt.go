package qtgui
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
// src-file: /QtGui/qaccessiblebridge.h
// dst-file: /src/gui/qaccessiblebridge.go
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
  // proto:  const QMetaObject * QAccessibleBridgePlugin::metaObject();
extern void C_ZNK23QAccessibleBridgePlugin10metaObjectEv(void* qthis); // 4
  // proto:  void QAccessibleBridgePlugin::QAccessibleBridgePlugin(QObject * parent);
extern void* C_ZN23QAccessibleBridgePluginC2EP7QObject(void* arg0); // 3
  // proto:  void QAccessibleBridgePlugin::~QAccessibleBridgePlugin();
extern void C_ZN23QAccessibleBridgePluginD2Ev(void* qthis); // 4
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

// class sizeof(QAccessibleBridgePlugin)=1
type QAccessibleBridgePlugin struct {
  /*qbase*/ qtcore.QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleBridge)=8
type QAccessibleBridge struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// metaObject()
func (this *QAccessibleBridgePlugin) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QAccessibleBridgePlugin10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK23QAccessibleBridgePlugin10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleBridgePlugin", "metaObject", args)
  }

  return
}

// QAccessibleBridgePlugin(class QObject *)
func NewQAccessibleBridgePlugin(args ...interface{}) *QAccessibleBridgePlugin {
  // QAccessibleBridgePlugin(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QAccessibleBridgePluginC1EP7QObject
    // invoke: void QAccessibleBridgePlugin(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN23QAccessibleBridgePluginC2EP7QObject(arg0)
    return &QAccessibleBridgePlugin{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QAccessibleBridgePlugin", "QAccessibleBridgePlugin", args)
  }

  return nil // QAccessibleBridgePlugin{Qclsinst:qthis}
}

// ~QAccessibleBridgePlugin()
func (this *QAccessibleBridgePlugin) Freeqaccessiblebridgeplugin(args ...interface{}) () {
  // ~QAccessibleBridgePlugin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QAccessibleBridgePluginD0Ev
    // invoke: void ~QAccessibleBridgePlugin()
    C.C_ZN23QAccessibleBridgePluginD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleBridgePlugin", "~QAccessibleBridgePlugin", args)
  }

  return
}

// <= body block end

