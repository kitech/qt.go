package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
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
  // proto:  QStyle * QStylePlugin::create(const QString & key);
extern void _ZN12QStylePlugin6createERK7QString(void* qthis, void* arg0);
  // proto:  const QMetaObject * QStylePlugin::metaObject();
extern void _ZNK12QStylePlugin10metaObjectEv(void* qthis);
  // proto:  void QStylePlugin::QStylePlugin(QObject * parent);
extern void* dector_ZN12QStylePluginC1EP7QObject(void* arg0);
extern void _ZN12QStylePluginC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QStylePlugin::~QStylePlugin();
extern void _ZN12QStylePluginD0Ev(void* qthis);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QStyle * QStylePlugin::create(const QString & key);
func (this *QStylePlugin) create(args ...interface{}) () {
  // create(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QStylePlugin6createERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QStylePlugin", "create", args)
  }

}

  // proto:  const QMetaObject * QStylePlugin::metaObject();
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
  default:
    qtrt.ErrorResolve("QStylePlugin", "metaObject", args)
  }

}

  // proto:  void QStylePlugin::QStylePlugin(QObject * parent);
func NewQStylePlugin(args ...interface{}) QStylePlugin {
  return QStylePlugin{}
}

  // proto:  void QStylePlugin::~QStylePlugin();
func (this *QStylePlugin) FreeQStylePlugin(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStylePlugin", "~QStylePlugin", args)
  }

}

// <= body block end

