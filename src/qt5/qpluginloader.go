package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  bool QPluginLoader::load();
extern void _ZN13QPluginLoader4loadEv(void* qthis); // 4
  // proto:  QString QPluginLoader::errorString();
extern void _ZNK13QPluginLoader11errorStringEv(void* qthis); // 4
  // proto:  bool QPluginLoader::unload();
extern void _ZN13QPluginLoader6unloadEv(void* qthis); // 4
  // proto:  void QPluginLoader::QPluginLoader(QObject * parent);
extern void _ZN13QPluginLoaderC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  void QPluginLoader::QPluginLoader(const QString & fileName, QObject * parent);
extern void _ZN13QPluginLoaderC2ERK7QStringP7QObject(void* qthis, void* arg0, void* arg1); // 3
  // proto:  void QPluginLoader::~QPluginLoader();
extern void _ZN13QPluginLoaderD2Ev(void* qthis); // 4
  // proto:  void QPluginLoader::setFileName(const QString & fileName);
extern void _ZN13QPluginLoader11setFileNameERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QPluginLoader::isLoaded();
extern void _ZNK13QPluginLoader8isLoadedEv(void* qthis); // 4
  // proto:  QObject * QPluginLoader::instance();
extern void _ZN13QPluginLoader8instanceEv(void* qthis); // 4
  // proto:  QString QPluginLoader::fileName();
extern void _ZNK13QPluginLoader8fileNameEv(void* qthis); // 4
  // proto:  QLibrary::LoadHints QPluginLoader::loadHints();
extern void _ZNK13QPluginLoader9loadHintsEv(void* qthis); // 4
  // proto:  const QMetaObject * QPluginLoader::metaObject();
extern void _ZNK13QPluginLoader10metaObjectEv(void* qthis); // 4
  // proto:  QJsonObject QPluginLoader::metaData();
extern void _ZNK13QPluginLoader8metaDataEv(void* qthis); // 4
  // proto: static QObjectList QPluginLoader::staticInstances();
extern void _ZN13QPluginLoader15staticInstancesEv(); // 4
  // proto: static QVector<QStaticPlugin> QPluginLoader::staticPlugins();
extern void _ZN13QPluginLoader13staticPluginsEv(); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// load()
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
    // invoke: bool load()
    C._ZN13QPluginLoader4loadEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPluginLoader", "load", args)
  }

}

// errorString()
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
    // invoke: QString errorString()
    C._ZNK13QPluginLoader11errorStringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPluginLoader", "errorString", args)
  }

}

// unload()
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
    // invoke: bool unload()
    C._ZN13QPluginLoader6unloadEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPluginLoader", "unload", args)
  }

}

// QPluginLoader(class QObject *)
func NewQPluginLoader(args ...interface{}) QPluginLoader {
  // QPluginLoader(class QObject *)
  // QPluginLoader(const class QString &, class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPluginLoaderC1EP7QObject
    // invoke: void QPluginLoader(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN13QPluginLoaderC2EP7QObject(qthis, arg0)
  case 1:
    // invoke: _ZN13QPluginLoaderC1ERK7QStringP7QObject
    // invoke: void QPluginLoader(const class QString &, class QObject *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN13QPluginLoaderC2ERK7QStringP7QObject(qthis, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPluginLoader", "QPluginLoader", args)
  }

  return QPluginLoader{}
}

// ~QPluginLoader()
func (this *QPluginLoader) FreeQPluginLoader(args ...interface{}) () {
  // ~QPluginLoader()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPluginLoaderD0Ev
    // invoke: void ~QPluginLoader()
    C._ZN13QPluginLoaderD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPluginLoader", "~QPluginLoader", args)
  }

}

// setFileName(const class QString &)
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
    // invoke: void setFileName(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QPluginLoader11setFileNameERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPluginLoader", "setFileName", args)
  }

}

// isLoaded()
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
    // invoke: bool isLoaded()
    C._ZNK13QPluginLoader8isLoadedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPluginLoader", "isLoaded", args)
  }

}

// instance()
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
    // invoke: QObject * instance()
    C._ZN13QPluginLoader8instanceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPluginLoader", "instance", args)
  }

}

// fileName()
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
    // invoke: QString fileName()
    C._ZNK13QPluginLoader8fileNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPluginLoader", "fileName", args)
  }

}

// loadHints()
func (this *QPluginLoader) loadHints(args ...interface{}) () {
  // loadHints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPluginLoader9loadHintsEv
    // invoke: QLibrary::LoadHints loadHints()
    C._ZNK13QPluginLoader9loadHintsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPluginLoader", "loadHints", args)
  }

}

// metaObject()
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
    // invoke: const QMetaObject * metaObject()
    C._ZNK13QPluginLoader10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPluginLoader", "metaObject", args)
  }

}

// metaData()
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
    // invoke: QJsonObject metaData()
    C._ZNK13QPluginLoader8metaDataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPluginLoader", "metaData", args)
  }

}

// staticInstances()
func (this *QPluginLoader) staticInstances_s(args ...interface{}) () {
  // staticInstances()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPluginLoader15staticInstancesEv
    // invoke: QObjectList staticInstances()
    C._ZN13QPluginLoader15staticInstancesEv()
  default:
    qtrt.ErrorResolve("QPluginLoader", "staticInstances", args)
  }

}

// staticPlugins()
func (this *QPluginLoader) staticPlugins_s(args ...interface{}) () {
  // staticPlugins()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPluginLoader13staticPluginsEv
    // invoke: QVector<QStaticPlugin> staticPlugins()
    C._ZN13QPluginLoader13staticPluginsEv()
  default:
    qtrt.ErrorResolve("QPluginLoader", "staticPlugins", args)
  }

}

// <= body block end

