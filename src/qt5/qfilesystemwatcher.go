package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
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
  // proto:  void QFileSystemWatcher::~QFileSystemWatcher();
extern void _ZN18QFileSystemWatcherD0Ev(void* qthis);
  // proto:  bool QFileSystemWatcher::removePath(const QString & file);
extern void _ZN18QFileSystemWatcher10removePathERK7QString(void* qthis, void* arg0);
  // proto:  QStringList QFileSystemWatcher::directories();
extern void _ZNK18QFileSystemWatcher11directoriesEv(void* qthis);
  // proto:  QStringList QFileSystemWatcher::files();
extern void _ZNK18QFileSystemWatcher5filesEv(void* qthis);
  // proto:  QStringList QFileSystemWatcher::addPaths(const QStringList & files);
extern void _ZN18QFileSystemWatcher8addPathsERK11QStringList(void* qthis, void* arg0);
  // proto:  QStringList QFileSystemWatcher::removePaths(const QStringList & files);
extern void _ZN18QFileSystemWatcher11removePathsERK11QStringList(void* qthis, void* arg0);
  // proto:  void QFileSystemWatcher::QFileSystemWatcher(QObject * parent);
extern void* dector_ZN18QFileSystemWatcherC1EP7QObject(void* arg0);
extern void _ZN18QFileSystemWatcherC1EP7QObject(void* qthis, void* arg0);
  // proto:  bool QFileSystemWatcher::addPath(const QString & file);
extern void _ZN18QFileSystemWatcher7addPathERK7QString(void* qthis, void* arg0);
  // proto:  const QMetaObject * QFileSystemWatcher::metaObject();
extern void _ZNK18QFileSystemWatcher10metaObjectEv(void* qthis);
  // proto:  void QFileSystemWatcher::QFileSystemWatcher(const QStringList & paths, QObject * parent);
extern void* dector_ZN18QFileSystemWatcherC1ERK11QStringListP7QObject(void* arg0, void* arg1);
extern void _ZN18QFileSystemWatcherC1ERK11QStringListP7QObject(void* qthis, void* arg0, void* arg1);
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
  qclsinst unsafe.Pointer /* *C.void */;
//  _directoryChanged QFileSystemWatcher_directoryChanged_signal;
//  _fileChanged QFileSystemWatcher_fileChanged_signal;
}

  // proto:  void QFileSystemWatcher::~QFileSystemWatcher();
func (this *QFileSystemWatcher) FreeQFileSystemWatcher(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFileSystemWatcher", "~QFileSystemWatcher", args)
  }

}

  // proto:  bool QFileSystemWatcher::removePath(const QString & file);
func (this *QFileSystemWatcher) removePath(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN18QFileSystemWatcher10removePathERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileSystemWatcher", "removePath", args)
  }

}

  // proto:  QStringList QFileSystemWatcher::directories();
func (this *QFileSystemWatcher) directories(args ...interface{}) () {
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
    C._ZNK18QFileSystemWatcher11directoriesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileSystemWatcher", "directories", args)
  }

}

  // proto:  QStringList QFileSystemWatcher::files();
func (this *QFileSystemWatcher) files(args ...interface{}) () {
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
    C._ZNK18QFileSystemWatcher5filesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileSystemWatcher", "files", args)
  }

}

  // proto:  QStringList QFileSystemWatcher::addPaths(const QStringList & files);
func (this *QFileSystemWatcher) addPaths(args ...interface{}) () {
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
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN18QFileSystemWatcher8addPathsERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileSystemWatcher", "addPaths", args)
  }

}

  // proto:  QStringList QFileSystemWatcher::removePaths(const QStringList & files);
func (this *QFileSystemWatcher) removePaths(args ...interface{}) () {
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
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN18QFileSystemWatcher11removePathsERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileSystemWatcher", "removePaths", args)
  }

}

  // proto:  void QFileSystemWatcher::QFileSystemWatcher(QObject * parent);
func NewQFileSystemWatcher(args ...interface{}) QFileSystemWatcher {
  return QFileSystemWatcher{}
}

  // proto:  bool QFileSystemWatcher::addPath(const QString & file);
func (this *QFileSystemWatcher) addPath(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN18QFileSystemWatcher7addPathERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileSystemWatcher", "addPath", args)
  }

}

  // proto:  const QMetaObject * QFileSystemWatcher::metaObject();
func (this *QFileSystemWatcher) metaObject(args ...interface{}) () {
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
    C._ZNK18QFileSystemWatcher10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileSystemWatcher", "metaObject", args)
  }

}

// <= body block end

