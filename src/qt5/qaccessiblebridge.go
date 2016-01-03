package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
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
  // proto:  void QAccessibleBridgePlugin::QAccessibleBridgePlugin(QObject * parent);
extern void* dector_ZN23QAccessibleBridgePluginC1EP7QObject(void* arg0);
extern void _ZN23QAccessibleBridgePluginC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QAccessibleBridgePlugin::~QAccessibleBridgePlugin();
extern void _ZN23QAccessibleBridgePluginD0Ev(void* qthis);
  // proto:  const QMetaObject * QAccessibleBridgePlugin::metaObject();
extern void _ZNK23QAccessibleBridgePlugin10metaObjectEv(void* qthis);
  // proto:  void QAccessibleBridge::~QAccessibleBridge();
extern void _ZN17QAccessibleBridgeD0Ev(void* qthis);
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

// class sizeof(QAccessibleBridgePlugin)=1
type QAccessibleBridgePlugin struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleBridge)=8
type QAccessibleBridge struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  void QAccessibleBridgePlugin::QAccessibleBridgePlugin(QObject * parent);
func NewQAccessibleBridgePlugin(args ...interface{}) QAccessibleBridgePlugin {
  return QAccessibleBridgePlugin{}
}

  // proto:  void QAccessibleBridgePlugin::~QAccessibleBridgePlugin();
func (this *QAccessibleBridgePlugin) FreeQAccessibleBridgePlugin(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleBridgePlugin", "~QAccessibleBridgePlugin", args)
  }

}

  // proto:  const QMetaObject * QAccessibleBridgePlugin::metaObject();
func (this *QAccessibleBridgePlugin) metaObject(args ...interface{}) () {
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
    C._ZNK23QAccessibleBridgePlugin10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleBridgePlugin", "metaObject", args)
  }

}

  // proto:  void QAccessibleBridge::~QAccessibleBridge();
func (this *QAccessibleBridge) FreeQAccessibleBridge(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleBridge", "~QAccessibleBridge", args)
  }

}

// <= body block end

