package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtGui/qiconengineplugin.h
// dst-file: /src/gui/qiconengineplugin.go
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
  // proto:  const QMetaObject * QIconEnginePlugin::metaObject();
extern void _ZNK17QIconEnginePlugin10metaObjectEv(void* qthis); // 4
  // proto:  void QIconEnginePlugin::QIconEnginePlugin(QObject * parent);
extern void _ZN17QIconEnginePluginC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  void QIconEnginePlugin::~QIconEnginePlugin();
extern void _ZN17QIconEnginePluginD2Ev(void* qthis); // 4
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

// class sizeof(QIconEnginePlugin)=1
type QIconEnginePlugin struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
}

// metaObject()
func (this *QIconEnginePlugin) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QIconEnginePlugin10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK17QIconEnginePlugin10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIconEnginePlugin", "metaObject", args)
  }

}

// QIconEnginePlugin(class QObject *)
func NewQIconEnginePlugin(args ...interface{}) QIconEnginePlugin {
  // QIconEnginePlugin(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QIconEnginePluginC1EP7QObject
    // invoke: void QIconEnginePlugin(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN17QIconEnginePluginC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QIconEnginePlugin", "QIconEnginePlugin", args)
  }

  return QIconEnginePlugin{}
}

// ~QIconEnginePlugin()
func (this *QIconEnginePlugin) FreeQIconEnginePlugin(args ...interface{}) () {
  // ~QIconEnginePlugin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QIconEnginePluginD0Ev
    // invoke: void ~QIconEnginePlugin()
    C._ZN17QIconEnginePluginD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIconEnginePlugin", "~QIconEnginePlugin", args)
  }

}

// <= body block end

