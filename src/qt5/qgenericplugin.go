package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtGui/qgenericplugin.h
// dst-file: /src/gui/qgenericplugin.go
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
  // proto:  void QGenericPlugin::QGenericPlugin(QObject * parent);
extern void* dector_ZN14QGenericPluginC1EP7QObject(void* arg0);
extern void _ZN14QGenericPluginC1EP7QObject(void* qthis, void* arg0);
  // proto:  QObject * QGenericPlugin::create(const QString & name, const QString & spec);
extern void _ZN14QGenericPlugin6createERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  void QGenericPlugin::~QGenericPlugin();
extern void _ZN14QGenericPluginD0Ev(void* qthis);
  // proto:  const QMetaObject * QGenericPlugin::metaObject();
extern void _ZNK14QGenericPlugin10metaObjectEv(void* qthis);
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

// class sizeof(QGenericPlugin)=1
type QGenericPlugin struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QGenericPlugin::QGenericPlugin(QObject * parent);
func NewQGenericPlugin(args ...interface{}) QGenericPlugin {
  return QGenericPlugin{}
}

  // proto:  QObject * QGenericPlugin::create(const QString & name, const QString & spec);
func (this *QGenericPlugin) create(args ...interface{}) () {
  // create(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGenericPlugin6createERK7QStringS2_
  default:
    qtrt.ErrorResolve("QGenericPlugin", "create", args)
  }

}

  // proto:  void QGenericPlugin::~QGenericPlugin();
func (this *QGenericPlugin) FreeQGenericPlugin(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGenericPlugin", "~QGenericPlugin", args)
  }

}

  // proto:  const QMetaObject * QGenericPlugin::metaObject();
func (this *QGenericPlugin) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGenericPlugin10metaObjectEv
  default:
    qtrt.ErrorResolve("QGenericPlugin", "metaObject", args)
  }

}

// <= body block end

