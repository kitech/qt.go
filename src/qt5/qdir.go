package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtCore/qdir.h
// dst-file: /src/core/qdir.go
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
  // proto: static void QDir::addResourceSearchPath(const QString & path);
extern void _ZN4QDir21addResourceSearchPathERK7QString(void* arg0);
  // proto: static bool QDir::isAbsolutePath(const QString & path);
extern void demth_ZN4QDir14isAbsolutePathERK7QString(void* arg0);
  // proto:  QString QDir::relativeFilePath(const QString & fileName);
extern void _ZNK4QDir16relativeFilePathERK7QString(void* qthis, void* arg0);
  // proto: static QString QDir::tempPath();
extern void _ZN4QDir8tempPathEv();
  // proto:  void QDir::QDir(const QString & path);
extern void* dector_ZN4QDirC1ERK7QString(void* arg0);
extern void _ZN4QDirC1ERK7QString(void* qthis, void* arg0);
  // proto:  QStringList QDir::nameFilters();
extern void _ZNK4QDir11nameFiltersEv(void* qthis);
  // proto:  bool QDir::cd(const QString & dirName);
extern void _ZN4QDir2cdERK7QString(void* qthis, void* arg0);
  // proto: static QFileInfoList QDir::drives();
extern void _ZN4QDir6drivesEv();
  // proto: static bool QDir::setCurrent(const QString & path);
extern void _ZN4QDir10setCurrentERK7QString(void* arg0);
  // proto:  bool QDir::rmdir(const QString & dirName);
extern void _ZNK4QDir5rmdirERK7QString(void* qthis, void* arg0);
  // proto:  bool QDir::cdUp();
extern void _ZN4QDir4cdUpEv(void* qthis);
  // proto:  QString QDir::absolutePath();
extern void _ZNK4QDir12absolutePathEv(void* qthis);
  // proto: static void QDir::setSearchPaths(const QString & prefix, const QStringList & searchPaths);
extern void _ZN4QDir14setSearchPathsERK7QStringRK11QStringList(void* arg0, void* arg1);
  // proto:  QString QDir::absoluteFilePath(const QString & fileName);
extern void _ZNK4QDir16absoluteFilePathERK7QString(void* qthis, void* arg0);
  // proto:  bool QDir::rename(const QString & oldName, const QString & newName);
extern void _ZN4QDir6renameERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto: static bool QDir::match(const QString & filter, const QString & fileName);
extern void _ZN4QDir5matchERK7QStringS2_(void* arg0, void* arg1);
  // proto:  void QDir::refresh();
extern void _ZNK4QDir7refreshEv(void* qthis);
  // proto:  bool QDir::mkdir(const QString & dirName);
extern void _ZNK4QDir5mkdirERK7QString(void* qthis, void* arg0);
  // proto:  uint QDir::count();
extern void _ZNK4QDir5countEv(void* qthis);
  // proto: static QDir QDir::root();
extern void demth_ZN4QDir4rootEv();
  // proto: static QStringList QDir::nameFiltersFromString(const QString & nameFilter);
extern void _ZN4QDir21nameFiltersFromStringERK7QString(void* arg0);
  // proto:  QString QDir::filePath(const QString & fileName);
extern void _ZNK4QDir8filePathERK7QString(void* qthis, void* arg0);
  // proto:  bool QDir::rmpath(const QString & dirPath);
extern void _ZNK4QDir6rmpathERK7QString(void* qthis, void* arg0);
  // proto:  QString QDir::path();
extern void _ZNK4QDir4pathEv(void* qthis);
  // proto: static QString QDir::toNativeSeparators(const QString & pathName);
extern void _ZN4QDir18toNativeSeparatorsERK7QString(void* arg0);
  // proto: static QString QDir::cleanPath(const QString & path);
extern void _ZN4QDir9cleanPathERK7QString(void* arg0);
  // proto:  bool QDir::exists();
extern void _ZNK4QDir6existsEv(void* qthis);
  // proto:  bool QDir::remove(const QString & fileName);
extern void _ZN4QDir6removeERK7QString(void* qthis, void* arg0);
  // proto:  void QDir::~QDir();
extern void _ZN4QDirD0Ev(void* qthis);
  // proto: static QString QDir::rootPath();
extern void _ZN4QDir8rootPathEv();
  // proto: static QDir QDir::current();
extern void demth_ZN4QDir7currentEv();
  // proto: static bool QDir::match(const QStringList & filters, const QString & fileName);
extern void _ZN4QDir5matchERK11QStringListRK7QString(void* arg0, void* arg1);
  // proto: static QString QDir::fromNativeSeparators(const QString & pathName);
extern void _ZN4QDir20fromNativeSeparatorsERK7QString(void* arg0);
  // proto: static QDir QDir::home();
extern void demth_ZN4QDir4homeEv();
  // proto:  void QDir::setNameFilters(const QStringList & nameFilters);
extern void _ZN4QDir14setNameFiltersERK11QStringList(void* qthis, void* arg0);
  // proto: static QChar QDir::separator();
extern void _ZN4QDir9separatorEv();
  // proto:  void QDir::swap(QDir & other);
extern void demth_ZN4QDir4swapERS_(void* qthis, void* arg0);
  // proto: static QDir QDir::temp();
extern void demth_ZN4QDir4tempEv();
  // proto:  bool QDir::exists(const QString & name);
extern void _ZNK4QDir6existsERK7QString(void* qthis, void* arg0);
  // proto:  bool QDir::mkpath(const QString & dirPath);
extern void _ZNK4QDir6mkpathERK7QString(void* qthis, void* arg0);
  // proto: static void QDir::addSearchPath(const QString & prefix, const QString & path);
extern void _ZN4QDir13addSearchPathERK7QStringS2_(void* arg0, void* arg1);
  // proto:  QString QDir::dirName();
extern void _ZNK4QDir7dirNameEv(void* qthis);
  // proto: static QStringList QDir::searchPaths(const QString & prefix);
extern void _ZN4QDir11searchPathsERK7QString(void* arg0);
  // proto:  bool QDir::makeAbsolute();
extern void _ZN4QDir12makeAbsoluteEv(void* qthis);
  // proto:  QString QDir::canonicalPath();
extern void _ZNK4QDir13canonicalPathEv(void* qthis);
  // proto:  bool QDir::isReadable();
extern void _ZNK4QDir10isReadableEv(void* qthis);
  // proto:  bool QDir::isRelative();
extern void _ZNK4QDir10isRelativeEv(void* qthis);
  // proto: static QString QDir::currentPath();
extern void _ZN4QDir11currentPathEv();
  // proto:  bool QDir::isRoot();
extern void _ZNK4QDir6isRootEv(void* qthis);
  // proto:  bool QDir::removeRecursively();
extern void _ZN4QDir17removeRecursivelyEv(void* qthis);
  // proto:  bool QDir::isAbsolute();
extern void demth_ZNK4QDir10isAbsoluteEv(void* qthis);
  // proto: static QString QDir::homePath();
extern void _ZN4QDir8homePathEv();
  // proto:  void QDir::QDir(const QDir & );
extern void* dector_ZN4QDirC1ERKS_(void* arg0);
extern void _ZN4QDirC1ERKS_(void* qthis, void* arg0);
  // proto:  void QDir::setPath(const QString & path);
extern void _ZN4QDir7setPathERK7QString(void* qthis, void* arg0);
  // proto: static bool QDir::isRelativePath(const QString & path);
extern void _ZN4QDir14isRelativePathERK7QString(void* arg0);
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

// class sizeof(QDir)=1
type QDir struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto: static void QDir::addResourceSearchPath(const QString & path);
func (this *QDir) addResourceSearchPath_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDir", "addResourceSearchPath", args)
  }

}

  // proto: static bool QDir::isAbsolutePath(const QString & path);
func (this *QDir) isAbsolutePath_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDir", "isAbsolutePath", args)
  }

}

  // proto:  QString QDir::relativeFilePath(const QString & fileName);
func (this *QDir) relativeFilePath(args ...interface{}) () {
  // relativeFilePath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QDir16relativeFilePathERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QDir", "relativeFilePath", args)
  }

}

  // proto: static QString QDir::tempPath();
func (this *QDir) tempPath_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDir", "tempPath", args)
  }

}

  // proto:  void QDir::QDir(const QString & path);
func NewQDir(args ...interface{}) QDir {
  return QDir{}
}

  // proto:  QStringList QDir::nameFilters();
func (this *QDir) nameFilters(args ...interface{}) () {
  // nameFilters()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QDir11nameFiltersEv
  default:
    qtrt.ErrorResolve("QDir", "nameFilters", args)
  }

}

  // proto:  bool QDir::cd(const QString & dirName);
func (this *QDir) cd(args ...interface{}) () {
  // cd(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir2cdERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QDir", "cd", args)
  }

}

  // proto: static QFileInfoList QDir::drives();
func (this *QDir) drives_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDir", "drives", args)
  }

}

  // proto: static bool QDir::setCurrent(const QString & path);
func (this *QDir) setCurrent_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDir", "setCurrent", args)
  }

}

  // proto:  bool QDir::rmdir(const QString & dirName);
func (this *QDir) rmdir(args ...interface{}) () {
  // rmdir(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QDir5rmdirERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QDir", "rmdir", args)
  }

}

  // proto:  bool QDir::cdUp();
func (this *QDir) cdUp(args ...interface{}) () {
  // cdUp()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir4cdUpEv
  default:
    qtrt.ErrorResolve("QDir", "cdUp", args)
  }

}

  // proto:  QString QDir::absolutePath();
func (this *QDir) absolutePath(args ...interface{}) () {
  // absolutePath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QDir12absolutePathEv
  default:
    qtrt.ErrorResolve("QDir", "absolutePath", args)
  }

}

  // proto: static void QDir::setSearchPaths(const QString & prefix, const QStringList & searchPaths);
func (this *QDir) setSearchPaths_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDir", "setSearchPaths", args)
  }

}

  // proto:  QString QDir::absoluteFilePath(const QString & fileName);
func (this *QDir) absoluteFilePath(args ...interface{}) () {
  // absoluteFilePath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QDir16absoluteFilePathERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QDir", "absoluteFilePath", args)
  }

}

  // proto:  bool QDir::rename(const QString & oldName, const QString & newName);
func (this *QDir) rename(args ...interface{}) () {
  // rename(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir6renameERK7QStringS2_
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QDir", "rename", args)
  }

}

  // proto: static bool QDir::match(const QString & filter, const QString & fileName);
func (this *QDir) match_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDir", "match", args)
  }

}

  // proto:  void QDir::refresh();
func (this *QDir) refresh(args ...interface{}) () {
  // refresh()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QDir7refreshEv
  default:
    qtrt.ErrorResolve("QDir", "refresh", args)
  }

}

  // proto:  bool QDir::mkdir(const QString & dirName);
func (this *QDir) mkdir(args ...interface{}) () {
  // mkdir(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QDir5mkdirERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QDir", "mkdir", args)
  }

}

  // proto:  uint QDir::count();
func (this *QDir) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QDir5countEv
  default:
    qtrt.ErrorResolve("QDir", "count", args)
  }

}

  // proto: static QDir QDir::root();
func (this *QDir) root_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDir", "root", args)
  }

}

  // proto: static QStringList QDir::nameFiltersFromString(const QString & nameFilter);
func (this *QDir) nameFiltersFromString_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDir", "nameFiltersFromString", args)
  }

}

  // proto:  QString QDir::filePath(const QString & fileName);
func (this *QDir) filePath(args ...interface{}) () {
  // filePath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QDir8filePathERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QDir", "filePath", args)
  }

}

  // proto:  bool QDir::rmpath(const QString & dirPath);
func (this *QDir) rmpath(args ...interface{}) () {
  // rmpath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QDir6rmpathERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QDir", "rmpath", args)
  }

}

  // proto:  QString QDir::path();
func (this *QDir) path(args ...interface{}) () {
  // path()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QDir4pathEv
  default:
    qtrt.ErrorResolve("QDir", "path", args)
  }

}

  // proto: static QString QDir::toNativeSeparators(const QString & pathName);
func (this *QDir) toNativeSeparators_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDir", "toNativeSeparators", args)
  }

}

  // proto: static QString QDir::cleanPath(const QString & path);
func (this *QDir) cleanPath_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDir", "cleanPath", args)
  }

}

  // proto:  bool QDir::exists();
func (this *QDir) exists(args ...interface{}) () {
  // exists()
  // exists(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QDir6existsEv
  case 1:
    // invoke: _ZNK4QDir6existsERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QDir", "exists", args)
  }

}

  // proto:  bool QDir::remove(const QString & fileName);
func (this *QDir) remove(args ...interface{}) () {
  // remove(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir6removeERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QDir", "remove", args)
  }

}

  // proto:  void QDir::~QDir();
func (this *QDir) FreeQDir(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDir", "~QDir", args)
  }

}

  // proto: static QString QDir::rootPath();
func (this *QDir) rootPath_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDir", "rootPath", args)
  }

}

  // proto: static QDir QDir::current();
func (this *QDir) current_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDir", "current", args)
  }

}

  // proto: static QString QDir::fromNativeSeparators(const QString & pathName);
func (this *QDir) fromNativeSeparators_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDir", "fromNativeSeparators", args)
  }

}

  // proto: static QDir QDir::home();
func (this *QDir) home_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDir", "home", args)
  }

}

  // proto:  void QDir::setNameFilters(const QStringList & nameFilters);
func (this *QDir) setNameFilters(args ...interface{}) () {
  // setNameFilters(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir14setNameFiltersERK11QStringList
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QDir", "setNameFilters", args)
  }

}

  // proto: static QChar QDir::separator();
func (this *QDir) separator_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDir", "separator", args)
  }

}

  // proto:  void QDir::swap(QDir & other);
func (this *QDir) swap(args ...interface{}) () {
  // swap(class QDir &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDir{}) // "QDir &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir4swapERS_
    var arg0 = args[0].(QDir).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QDir", "swap", args)
  }

}

  // proto: static QDir QDir::temp();
func (this *QDir) temp_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDir", "temp", args)
  }

}

  // proto:  bool QDir::mkpath(const QString & dirPath);
func (this *QDir) mkpath(args ...interface{}) () {
  // mkpath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QDir6mkpathERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QDir", "mkpath", args)
  }

}

  // proto: static void QDir::addSearchPath(const QString & prefix, const QString & path);
func (this *QDir) addSearchPath_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDir", "addSearchPath", args)
  }

}

  // proto:  QString QDir::dirName();
func (this *QDir) dirName(args ...interface{}) () {
  // dirName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QDir7dirNameEv
  default:
    qtrt.ErrorResolve("QDir", "dirName", args)
  }

}

  // proto: static QStringList QDir::searchPaths(const QString & prefix);
func (this *QDir) searchPaths_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDir", "searchPaths", args)
  }

}

  // proto:  bool QDir::makeAbsolute();
func (this *QDir) makeAbsolute(args ...interface{}) () {
  // makeAbsolute()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir12makeAbsoluteEv
  default:
    qtrt.ErrorResolve("QDir", "makeAbsolute", args)
  }

}

  // proto:  QString QDir::canonicalPath();
func (this *QDir) canonicalPath(args ...interface{}) () {
  // canonicalPath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QDir13canonicalPathEv
  default:
    qtrt.ErrorResolve("QDir", "canonicalPath", args)
  }

}

  // proto:  bool QDir::isReadable();
func (this *QDir) isReadable(args ...interface{}) () {
  // isReadable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QDir10isReadableEv
  default:
    qtrt.ErrorResolve("QDir", "isReadable", args)
  }

}

  // proto:  bool QDir::isRelative();
func (this *QDir) isRelative(args ...interface{}) () {
  // isRelative()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QDir10isRelativeEv
  default:
    qtrt.ErrorResolve("QDir", "isRelative", args)
  }

}

  // proto: static QString QDir::currentPath();
func (this *QDir) currentPath_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDir", "currentPath", args)
  }

}

  // proto:  bool QDir::isRoot();
func (this *QDir) isRoot(args ...interface{}) () {
  // isRoot()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QDir6isRootEv
  default:
    qtrt.ErrorResolve("QDir", "isRoot", args)
  }

}

  // proto:  bool QDir::removeRecursively();
func (this *QDir) removeRecursively(args ...interface{}) () {
  // removeRecursively()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir17removeRecursivelyEv
  default:
    qtrt.ErrorResolve("QDir", "removeRecursively", args)
  }

}

  // proto:  bool QDir::isAbsolute();
func (this *QDir) isAbsolute(args ...interface{}) () {
  // isAbsolute()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QDir10isAbsoluteEv
  default:
    qtrt.ErrorResolve("QDir", "isAbsolute", args)
  }

}

  // proto: static QString QDir::homePath();
func (this *QDir) homePath_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDir", "homePath", args)
  }

}

  // proto:  void QDir::setPath(const QString & path);
func (this *QDir) setPath(args ...interface{}) () {
  // setPath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir7setPathERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QDir", "setPath", args)
  }

}

  // proto: static bool QDir::isRelativePath(const QString & path);
func (this *QDir) isRelativePath_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDir", "isRelativePath", args)
  }

}

// <= body block end

