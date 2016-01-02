package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtCore/qpluginloader.h
// dst-file: /src/core/qpluginloader.go
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
  // proto:  bool QPluginLoader::isLoaded();
extern void _ZNK13QPluginLoader8isLoadedEv(void* qthis);
  // proto:  bool QPluginLoader::unload();
extern void _ZN13QPluginLoader6unloadEv(void* qthis);
  // proto:  void QPluginLoader::QPluginLoader(const QString & fileName, QObject * parent);
extern void* dector_ZN13QPluginLoaderC1ERK7QStringP7QObject(void* arg0, void* arg1);
extern void _ZN13QPluginLoaderC1ERK7QStringP7QObject(void* qthis, void* arg0, void* arg1);
  // proto:  bool QPluginLoader::load();
extern void _ZN13QPluginLoader4loadEv(void* qthis);
  // proto:  const QMetaObject * QPluginLoader::metaObject();
extern void _ZNK13QPluginLoader10metaObjectEv(void* qthis);
  // proto:  QObject * QPluginLoader::instance();
extern void _ZN13QPluginLoader8instanceEv(void* qthis);
  // proto: static QVector<QStaticPlugin> QPluginLoader::staticPlugins();
extern void _ZN13QPluginLoader13staticPluginsEv();
  // proto: static QObjectList QPluginLoader::staticInstances();
extern void _ZN13QPluginLoader15staticInstancesEv();
  // proto:  QJsonObject QPluginLoader::metaData();
extern void _ZNK13QPluginLoader8metaDataEv(void* qthis);
  // proto:  QString QPluginLoader::errorString();
extern void _ZNK13QPluginLoader11errorStringEv(void* qthis);
  // proto:  void QPluginLoader::QPluginLoader(const QPluginLoader & );
extern void* dector_ZN13QPluginLoaderC1ERKS_(void* arg0);
extern void _ZN13QPluginLoaderC1ERKS_(void* qthis, void* arg0);
  // proto:  QString QPluginLoader::fileName();
extern void _ZNK13QPluginLoader8fileNameEv(void* qthis);
  // proto:  void QPluginLoader::setFileName(const QString & fileName);
extern void _ZN13QPluginLoader11setFileNameERK7QString(void* qthis, void* arg0);
  // proto:  void QPluginLoader::QPluginLoader(QObject * parent);
extern void* dector_ZN13QPluginLoaderC1EP7QObject(void* arg0);
extern void _ZN13QPluginLoaderC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QPluginLoader::~QPluginLoader();
extern void _ZN13QPluginLoaderD0Ev(void* qthis);
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

// class sizeof(QPluginLoader)=1
type QPluginLoader struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  bool QPluginLoader::isLoaded();
func (this *QPluginLoader) isLoaded(args ...interface{}) () {
  // isLoaded()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPluginLoader8isLoadedEv
  default:
    qtrt.ErrorResolve("QPluginLoader", "isLoaded", args)
  }

}

  // proto:  bool QPluginLoader::unload();
func (this *QPluginLoader) unload(args ...interface{}) () {
  // unload()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPluginLoader6unloadEv
  default:
    qtrt.ErrorResolve("QPluginLoader", "unload", args)
  }

}

  // proto:  void QPluginLoader::QPluginLoader(const QString & fileName, QObject * parent);
func NewQPluginLoader(args ...interface{}) QPluginLoader {
  return QPluginLoader{}
}

  // proto:  bool QPluginLoader::load();
func (this *QPluginLoader) load(args ...interface{}) () {
  // load()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPluginLoader4loadEv
  default:
    qtrt.ErrorResolve("QPluginLoader", "load", args)
  }

}

  // proto:  const QMetaObject * QPluginLoader::metaObject();
func (this *QPluginLoader) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPluginLoader10metaObjectEv
  default:
    qtrt.ErrorResolve("QPluginLoader", "metaObject", args)
  }

}

  // proto:  QObject * QPluginLoader::instance();
func (this *QPluginLoader) instance(args ...interface{}) () {
  // instance()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPluginLoader8instanceEv
  default:
    qtrt.ErrorResolve("QPluginLoader", "instance", args)
  }

}

  // proto: static QVector<QStaticPlugin> QPluginLoader::staticPlugins();
func (this *QPluginLoader) staticPlugins_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPluginLoader", "staticPlugins", args)
  }

}

  // proto: static QObjectList QPluginLoader::staticInstances();
func (this *QPluginLoader) staticInstances_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPluginLoader", "staticInstances", args)
  }

}

  // proto:  QJsonObject QPluginLoader::metaData();
func (this *QPluginLoader) metaData(args ...interface{}) () {
  // metaData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPluginLoader8metaDataEv
  default:
    qtrt.ErrorResolve("QPluginLoader", "metaData", args)
  }

}

  // proto:  QString QPluginLoader::errorString();
func (this *QPluginLoader) errorString(args ...interface{}) () {
  // errorString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPluginLoader11errorStringEv
  default:
    qtrt.ErrorResolve("QPluginLoader", "errorString", args)
  }

}

  // proto:  QString QPluginLoader::fileName();
func (this *QPluginLoader) fileName(args ...interface{}) () {
  // fileName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPluginLoader8fileNameEv
  default:
    qtrt.ErrorResolve("QPluginLoader", "fileName", args)
  }

}

  // proto:  void QPluginLoader::setFileName(const QString & fileName);
func (this *QPluginLoader) setFileName(args ...interface{}) () {
  // setFileName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPluginLoader11setFileNameERK7QString
  default:
    qtrt.ErrorResolve("QPluginLoader", "setFileName", args)
  }

}

  // proto:  void QPluginLoader::~QPluginLoader();
func (this *QPluginLoader) FreeQPluginLoader(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPluginLoader", "~QPluginLoader", args)
  }

}

// <= body block end

