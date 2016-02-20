package qtcore
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
extern bool C_ZN13QPluginLoader4loadEv(void* qthis); // 4
  // proto:  QString QPluginLoader::errorString();
extern void* C_ZNK13QPluginLoader11errorStringEv(void* qthis); // 4
  // proto:  bool QPluginLoader::unload();
extern bool C_ZN13QPluginLoader6unloadEv(void* qthis); // 4
  // proto:  void QPluginLoader::QPluginLoader(QObject * parent);
extern void* C_ZN13QPluginLoaderC2EP7QObject(void* arg0); // 3
  // proto:  void QPluginLoader::QPluginLoader(const QString & fileName, QObject * parent);
extern void* C_ZN13QPluginLoaderC2ERK7QStringP7QObject(void* arg0, void* arg1); // 3
  // proto:  void QPluginLoader::~QPluginLoader();
extern void C_ZN13QPluginLoaderD2Ev(void* qthis); // 4
  // proto:  void QPluginLoader::setFileName(const QString & fileName);
extern void C_ZN13QPluginLoader11setFileNameERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QPluginLoader::isLoaded();
extern bool C_ZNK13QPluginLoader8isLoadedEv(void* qthis); // 4
  // proto:  QObject * QPluginLoader::instance();
extern void* C_ZN13QPluginLoader8instanceEv(void* qthis); // 4
  // proto:  QString QPluginLoader::fileName();
extern void* C_ZNK13QPluginLoader8fileNameEv(void* qthis); // 4
  // proto:  QLibrary::LoadHints QPluginLoader::loadHints();
extern void C_ZNK13QPluginLoader9loadHintsEv(void* qthis); // 4
  // proto:  const QMetaObject * QPluginLoader::metaObject();
extern void C_ZNK13QPluginLoader10metaObjectEv(void* qthis); // 4
  // proto:  QJsonObject QPluginLoader::metaData();
extern void C_ZNK13QPluginLoader8metaDataEv(void* qthis); // 4
  // proto: static QObjectList QPluginLoader::staticInstances();
extern void C_ZN13QPluginLoader15staticInstancesEv(); // 4
  // proto: static QVector<QStaticPlugin> QPluginLoader::staticPlugins();
extern void C_ZN13QPluginLoader13staticPluginsEv(); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// load()
func (this *QPluginLoader) Load(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN13QPluginLoader4loadEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPluginLoader", "load", args)
  }

  return
}

// errorString()
func (this *QPluginLoader) Errorstring(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QPluginLoader11errorStringEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPluginLoader", "errorString", args)
  }

  return
}

// unload()
func (this *QPluginLoader) Unload(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN13QPluginLoader6unloadEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPluginLoader", "unload", args)
  }

  return
}

// QPluginLoader(class QObject *)
func NewQPluginLoader(args ...interface{}) *QPluginLoader {
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
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QPluginLoaderC2EP7QObject(arg0)
    return &QPluginLoader{Qclsinst:qthis}
  case 1:
    // invoke: _ZN13QPluginLoaderC1ERK7QStringP7QObject
    // invoke: void QPluginLoader(const class QString &, class QObject *)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QObject).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QPluginLoaderC2ERK7QStringP7QObject(arg0, arg1)
    return &QPluginLoader{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QPluginLoader", "QPluginLoader", args)
  }

  return nil // QPluginLoader{Qclsinst:qthis}
}

// ~QPluginLoader()
func (this *QPluginLoader) Freeqpluginloader(args ...interface{}) () {
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
    C.C_ZN13QPluginLoaderD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPluginLoader", "~QPluginLoader", args)
  }

  return
}

// setFileName(const class QString &)
func (this *QPluginLoader) Setfilename(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QPluginLoader11setFileNameERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPluginLoader", "setFileName", args)
  }

  return
}

// isLoaded()
func (this *QPluginLoader) Isloaded(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QPluginLoader8isLoadedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPluginLoader", "isLoaded", args)
  }

  return
}

// instance()
func (this *QPluginLoader) Instance(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN13QPluginLoader8instanceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QObject{}) // "QObject *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPluginLoader", "instance", args)
  }

  return
}

// fileName()
func (this *QPluginLoader) Filename(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QPluginLoader8fileNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPluginLoader", "fileName", args)
  }

  return
}

// loadHints()
func (this *QPluginLoader) Loadhints(args ...interface{}) () {
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
    C.C_ZNK13QPluginLoader9loadHintsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPluginLoader", "loadHints", args)
  }

  return
}

// metaObject()
func (this *QPluginLoader) Metaobject(args ...interface{}) () {
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
    C.C_ZNK13QPluginLoader10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPluginLoader", "metaObject", args)
  }

  return
}

// metaData()
func (this *QPluginLoader) Metadata(args ...interface{}) () {
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
    C.C_ZNK13QPluginLoader8metaDataEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPluginLoader", "metaData", args)
  }

  return
}

// staticInstances()
func (this *QPluginLoader) Staticinstances_S(args ...interface{}) () {
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
    C.C_ZN13QPluginLoader15staticInstancesEv()
  default:
    qtrt.ErrorResolve("QPluginLoader", "staticInstances", args)
  }

  return
}

// staticPlugins()
func (this *QPluginLoader) Staticplugins_S(args ...interface{}) () {
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
    C.C_ZN13QPluginLoader13staticPluginsEv()
  default:
    qtrt.ErrorResolve("QPluginLoader", "staticPlugins", args)
  }

  return
}

// <= body block end

