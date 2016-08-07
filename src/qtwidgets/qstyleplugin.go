package qtwidgets
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
import "qtcore"
import "qtgui"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QStylePlugin::~QStylePlugin();
extern void C_ZN12QStylePluginD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QStylePlugin::metaObject();
extern void C_ZNK12QStylePlugin10metaObjectEv(void* qthis); // 4
  // proto:  void QStylePlugin::QStylePlugin(QObject * parent);
extern void* C_ZN12QStylePluginC2EP7QObject(void* arg0); // 3
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QStylePlugin)=1
type QStylePlugin struct {
  /*qbase*/ qtcore.QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// ~QStylePlugin()
func (this *QStylePlugin) Freeqstyleplugin(args ...interface{}) () {
  // ~QStylePlugin()
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
    // invoke: _ZN12QStylePluginD0Ev
    // invoke: void ~QStylePlugin()
    C.C_ZN12QStylePluginD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStylePlugin", "~QStylePlugin", args)
  }

  return
}

// metaObject()
func (this *QStylePlugin) Metaobject(args ...interface{}) () {
  // metaObject()
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
    // invoke: _ZNK12QStylePlugin10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK12QStylePlugin10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStylePlugin", "metaObject", args)
  }

  return
}

// QStylePlugin(class QObject *)
func NewQStylePlugin(args ...interface{}) *QStylePlugin {
  // QStylePlugin(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QStylePluginC1EP7QObject
    // invoke: void QStylePlugin(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QStylePluginC2EP7QObject(arg0)
    return &QStylePlugin{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStylePlugin", "QStylePlugin", args)
  }

  return nil // QStylePlugin{Qclsinst:qthis}
}

// <= body block end

