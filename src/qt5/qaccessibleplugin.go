package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
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
  // proto:  void QAccessiblePlugin::QAccessiblePlugin(QObject * parent);
extern void* dector_ZN17QAccessiblePluginC1EP7QObject(void* arg0);
extern void _ZN17QAccessiblePluginC1EP7QObject(void* qthis, void* arg0);
  // proto:  const QMetaObject * QAccessiblePlugin::metaObject();
extern void _ZNK17QAccessiblePlugin10metaObjectEv(void* qthis);
  // proto:  void QAccessiblePlugin::~QAccessiblePlugin();
extern void _ZN17QAccessiblePluginD0Ev(void* qthis);
  // proto:  QAccessibleInterface * QAccessiblePlugin::create(const QString & key, QObject * object);
extern void _ZN17QAccessiblePlugin6createERK7QStringP7QObject(void* qthis, void* arg0, void* arg1);
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

// class sizeof(QAccessiblePlugin)=1
type QAccessiblePlugin struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QAccessiblePlugin::QAccessiblePlugin(QObject * parent);
func NewQAccessiblePlugin(args ...interface{}) QAccessiblePlugin {
  return QAccessiblePlugin{}
}

  // proto:  const QMetaObject * QAccessiblePlugin::metaObject();
func (this *QAccessiblePlugin) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAccessiblePlugin10metaObjectEv
  default:
    qtrt.ErrorResolve("QAccessiblePlugin", "metaObject", args)
  }

}

  // proto:  void QAccessiblePlugin::~QAccessiblePlugin();
func (this *QAccessiblePlugin) FreeQAccessiblePlugin(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessiblePlugin", "~QAccessiblePlugin", args)
  }

}

  // proto:  QAccessibleInterface * QAccessiblePlugin::create(const QString & key, QObject * object);
func (this *QAccessiblePlugin) create(args ...interface{}) () {
  // create(const class QString &, class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAccessiblePlugin6createERK7QStringP7QObject
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QAccessiblePlugin", "create", args)
  }

}

// <= body block end

