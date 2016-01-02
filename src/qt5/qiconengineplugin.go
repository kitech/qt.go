package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
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
  // proto:  const QMetaObject * QIconEnginePlugin::metaObject();
extern void _ZNK17QIconEnginePlugin10metaObjectEv(void* qthis);
  // proto:  void QIconEnginePlugin::QIconEnginePlugin(QObject * parent);
extern void* dector_ZN17QIconEnginePluginC1EP7QObject(void* arg0);
extern void _ZN17QIconEnginePluginC1EP7QObject(void* qthis, void* arg0);
  // proto:  QIconEngine * QIconEnginePlugin::create(const QString & filename);
extern void _ZN17QIconEnginePlugin6createERK7QString(void* qthis, void* arg0);
  // proto:  void QIconEnginePlugin::~QIconEnginePlugin();
extern void _ZN17QIconEnginePluginD0Ev(void* qthis);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  const QMetaObject * QIconEnginePlugin::metaObject();
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
  default:
    qtrt.ErrorResolve("QIconEnginePlugin", "metaObject", args)
  }

}

  // proto:  void QIconEnginePlugin::QIconEnginePlugin(QObject * parent);
func NewQIconEnginePlugin(args ...interface{}) QIconEnginePlugin {
  return QIconEnginePlugin{}
}

  // proto:  QIconEngine * QIconEnginePlugin::create(const QString & filename);
func (this *QIconEnginePlugin) create(args ...interface{}) () {
  // create(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QIconEnginePlugin6createERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QIconEnginePlugin", "create", args)
  }

}

  // proto:  void QIconEnginePlugin::~QIconEnginePlugin();
func (this *QIconEnginePlugin) FreeQIconEnginePlugin(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QIconEnginePlugin", "~QIconEnginePlugin", args)
  }

}

// <= body block end

