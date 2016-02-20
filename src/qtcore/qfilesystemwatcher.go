package qtcore
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
// src-file: /QtCore/qfilesystemwatcher.h
// dst-file: /src/core/qfilesystemwatcher.go
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
  // proto:  QStringList QFileSystemWatcher::files();
extern void C_ZNK18QFileSystemWatcher5filesEv(void* qthis); // 4
  // proto:  const QMetaObject * QFileSystemWatcher::metaObject();
extern void C_ZNK18QFileSystemWatcher10metaObjectEv(void* qthis); // 4
  // proto:  void QFileSystemWatcher::QFileSystemWatcher(QObject * parent);
extern void* C_ZN18QFileSystemWatcherC2EP7QObject(void* arg0); // 3
  // proto:  void QFileSystemWatcher::QFileSystemWatcher(const QStringList & paths, QObject * parent);
extern void* C_ZN18QFileSystemWatcherC2ERK11QStringListP7QObject(void* arg0, void* arg1); // 3
  // proto:  QStringList QFileSystemWatcher::removePaths(const QStringList & files);
extern void C_ZN18QFileSystemWatcher11removePathsERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  QStringList QFileSystemWatcher::directories();
extern void C_ZNK18QFileSystemWatcher11directoriesEv(void* qthis); // 4
  // proto:  bool QFileSystemWatcher::addPath(const QString & file);
extern bool C_ZN18QFileSystemWatcher7addPathERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QFileSystemWatcher::~QFileSystemWatcher();
extern void C_ZN18QFileSystemWatcherD2Ev(void* qthis); // 4
  // proto:  QStringList QFileSystemWatcher::addPaths(const QStringList & files);
extern void C_ZN18QFileSystemWatcher8addPathsERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  bool QFileSystemWatcher::removePath(const QString & file);
extern bool C_ZN18QFileSystemWatcher10removePathERK7QString(void* qthis, void* arg0); // 4
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

// class sizeof(QFileSystemWatcher)=1
type QFileSystemWatcher struct {
  /*qbase*/ QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _directoryChanged QFileSystemWatcher_directoryChanged_signal;
//  _fileChanged QFileSystemWatcher_fileChanged_signal;
}

// files()
func (this *QFileSystemWatcher) Files(args ...interface{}) () {
  // files()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFileSystemWatcher5filesEv
    // invoke: QStringList files()
    C.C_ZNK18QFileSystemWatcher5filesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileSystemWatcher", "files", args)
  }

  return
}

// metaObject()
func (this *QFileSystemWatcher) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFileSystemWatcher10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK18QFileSystemWatcher10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileSystemWatcher", "metaObject", args)
  }

  return
}

// QFileSystemWatcher(class QObject *)
func NewQFileSystemWatcher(args ...interface{}) *QFileSystemWatcher {
  // QFileSystemWatcher(class QObject *)
  // QFileSystemWatcher(const class QStringList &, class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"
  vtys[1][1] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFileSystemWatcherC1EP7QObject
    // invoke: void QFileSystemWatcher(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QFileSystemWatcherC2EP7QObject(arg0)
    return &QFileSystemWatcher{Qclsinst:qthis}
  case 1:
    // invoke: _ZN18QFileSystemWatcherC1ERK11QStringListP7QObject
    // invoke: void QFileSystemWatcher(const class QStringList &, class QObject *)
    var arg0 = args[0].(*QStringList).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QObject).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QFileSystemWatcherC2ERK11QStringListP7QObject(arg0, arg1)
    return &QFileSystemWatcher{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QFileSystemWatcher", "QFileSystemWatcher", args)
  }

  return nil // QFileSystemWatcher{Qclsinst:qthis}
}

// removePaths(const class QStringList &)
func (this *QFileSystemWatcher) Removepaths(args ...interface{}) () {
  // removePaths(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFileSystemWatcher11removePathsERK11QStringList
    // invoke: QStringList removePaths(const class QStringList &)
    var arg0 = args[0].(*QStringList).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QFileSystemWatcher11removePathsERK11QStringList(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileSystemWatcher", "removePaths", args)
  }

  return
}

// directories()
func (this *QFileSystemWatcher) Directories(args ...interface{}) () {
  // directories()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFileSystemWatcher11directoriesEv
    // invoke: QStringList directories()
    C.C_ZNK18QFileSystemWatcher11directoriesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileSystemWatcher", "directories", args)
  }

  return
}

// addPath(const class QString &)
func (this *QFileSystemWatcher) Addpath(args ...interface{}) (ret interface{}) {
  // addPath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFileSystemWatcher7addPathERK7QString
    // invoke: bool addPath(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN18QFileSystemWatcher7addPathERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileSystemWatcher", "addPath", args)
  }

  return
}

// ~QFileSystemWatcher()
func (this *QFileSystemWatcher) Freeqfilesystemwatcher(args ...interface{}) () {
  // ~QFileSystemWatcher()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFileSystemWatcherD0Ev
    // invoke: void ~QFileSystemWatcher()
    C.C_ZN18QFileSystemWatcherD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileSystemWatcher", "~QFileSystemWatcher", args)
  }

  return
}

// addPaths(const class QStringList &)
func (this *QFileSystemWatcher) Addpaths(args ...interface{}) () {
  // addPaths(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFileSystemWatcher8addPathsERK11QStringList
    // invoke: QStringList addPaths(const class QStringList &)
    var arg0 = args[0].(*QStringList).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QFileSystemWatcher8addPathsERK11QStringList(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileSystemWatcher", "addPaths", args)
  }

  return
}

// removePath(const class QString &)
func (this *QFileSystemWatcher) Removepath(args ...interface{}) (ret interface{}) {
  // removePath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFileSystemWatcher10removePathERK7QString
    // invoke: bool removePath(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN18QFileSystemWatcher10removePathERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileSystemWatcher", "removePath", args)
  }

  return
}

// <= body block end

