package qtgui
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtGui/qaccessibleplugin.h
// dst-file: /src/gui/qaccessibleplugin.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "runtime"
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
  // proto:  void QAccessiblePlugin::QAccessiblePlugin(QObject * parent);
extern void* C_ZN17QAccessiblePluginC2EP7QObject(void* arg0); // 3
  // proto:  const QMetaObject * QAccessiblePlugin::metaObject();
extern void C_ZNK17QAccessiblePlugin10metaObjectEv(void* qthis); // 4
  // proto:  void QAccessiblePlugin::~QAccessiblePlugin();
extern void C_ZN17QAccessiblePluginD2Ev(void* qthis); // 4
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
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QAccessiblePlugin)=1
type QAccessiblePlugin struct {
  /*qbase*/ qtcore.QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// QAccessiblePlugin(class QObject *)
func GcfreeQAccessiblePlugin(this *QAccessiblePlugin) {
  qtrt.UniverseFree(this)
}
func NewQAccessiblePlugin(args ...interface{}) *QAccessiblePlugin {
  // QAccessiblePlugin(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAccessiblePluginC1EP7QObject
    // invoke: void QAccessiblePlugin(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QAccessiblePluginC2EP7QObject(arg0)
    this := &QAccessiblePlugin{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQAccessiblePlugin)
    return this
  default:
    qtrt.ErrorResolve("QAccessiblePlugin", "QAccessiblePlugin", args)
  }

  return nil // QAccessiblePlugin{Qclsinst:qthis}
}

// metaObject()
func (this *QAccessiblePlugin) MetaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAccessiblePlugin10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK17QAccessiblePlugin10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAccessiblePlugin", "metaObject", args)
  }

  return
}

// ~QAccessiblePlugin()
func (this *QAccessiblePlugin) Free(args ...interface{}) () {
  // ~QAccessiblePlugin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAccessiblePluginD0Ev
    // invoke: void ~QAccessiblePlugin()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN17QAccessiblePluginD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QAccessiblePlugin", "~QAccessiblePlugin", args)
  }

  return
}

// <= body block end

