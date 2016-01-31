package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  bool QDir::rename(const QString & oldName, const QString & newName);
extern void C_ZN4QDir6renameERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QDir::setNameFilters(const QStringList & nameFilters);
extern void C_ZN4QDir14setNameFiltersERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  void QDir::~QDir();
extern void C_ZN4QDirD2Ev(void* qthis); // 4
  // proto:  bool QDir::exists();
extern void C_ZNK4QDir6existsEv(void* qthis); // 4
  // proto:  bool QDir::exists(const QString & name);
extern void C_ZNK4QDir6existsERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QDir::relativeFilePath(const QString & fileName);
extern void C_ZNK4QDir16relativeFilePathERK7QString(void* qthis, void* arg0); // 4
  // proto: static QString QDir::currentPath();
extern void C_ZN4QDir11currentPathEv(); // 4
  // proto:  bool QDir::isRelative();
extern void C_ZNK4QDir10isRelativeEv(void* qthis); // 4
  // proto:  bool QDir::cd(const QString & dirName);
extern void C_ZN4QDir2cdERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QDir::QDir(const QDir & );
extern void* C_ZN4QDirC2ERKS_(void* arg0); // 3
  // proto:  void QDir::QDir(const QString & path);
extern void* C_ZN4QDirC2ERK7QString(void* arg0); // 3
  // proto: static bool QDir::isAbsolutePath(const QString & path);
extern void C_ZN4QDir14isAbsolutePathERK7QString(void* arg0); // 2
  // proto:  bool QDir::remove(const QString & fileName);
extern void C_ZN4QDir6removeERK7QString(void* qthis, void* arg0); // 4
  // proto: static QString QDir::tempPath();
extern void C_ZN4QDir8tempPathEv(); // 4
  // proto: static QString QDir::homePath();
extern void C_ZN4QDir8homePathEv(); // 4
  // proto: static QDir QDir::home();
extern void C_ZN4QDir4homeEv(); // 2
  // proto: static QString QDir::cleanPath(const QString & path);
extern void C_ZN4QDir9cleanPathERK7QString(void* arg0); // 4
  // proto: static bool QDir::setCurrent(const QString & path);
extern void C_ZN4QDir10setCurrentERK7QString(void* arg0); // 4
  // proto: static QString QDir::toNativeSeparators(const QString & pathName);
extern void C_ZN4QDir18toNativeSeparatorsERK7QString(void* arg0); // 4
  // proto:  Filters QDir::filter();
extern void C_ZNK4QDir6filterEv(void* qthis); // 4
  // proto: static QStringList QDir::searchPaths(const QString & prefix);
extern void C_ZN4QDir11searchPathsERK7QString(void* arg0); // 4
  // proto:  bool QDir::isRoot();
extern void C_ZNK4QDir6isRootEv(void* qthis); // 4
  // proto:  QString QDir::filePath(const QString & fileName);
extern void C_ZNK4QDir8filePathERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QDir::mkdir(const QString & dirName);
extern void C_ZNK4QDir5mkdirERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QDir::absolutePath();
extern void C_ZNK4QDir12absolutePathEv(void* qthis); // 4
  // proto: static void QDir::addResourceSearchPath(const QString & path);
extern void C_ZN4QDir21addResourceSearchPathERK7QString(void* arg0); // 4
  // proto: static QDir QDir::current();
extern void C_ZN4QDir7currentEv(); // 2
  // proto: static QString QDir::fromNativeSeparators(const QString & pathName);
extern void C_ZN4QDir20fromNativeSeparatorsERK7QString(void* arg0); // 4
  // proto: static bool QDir::isRelativePath(const QString & path);
extern void C_ZN4QDir14isRelativePathERK7QString(void* arg0); // 4
  // proto:  bool QDir::rmdir(const QString & dirName);
extern void C_ZNK4QDir5rmdirERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QDir::canonicalPath();
extern void C_ZNK4QDir13canonicalPathEv(void* qthis); // 4
  // proto:  bool QDir::removeRecursively();
extern void C_ZN4QDir17removeRecursivelyEv(void* qthis); // 4
  // proto:  void QDir::setPath(const QString & path);
extern void C_ZN4QDir7setPathERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QDir::dirName();
extern void C_ZNK4QDir7dirNameEv(void* qthis); // 4
  // proto: static bool QDir::match(const QStringList & filters, const QString & fileName);
extern void C_ZN4QDir5matchERK11QStringListRK7QString(void* arg0, void* arg1); // 4
  // proto: static bool QDir::match(const QString & filter, const QString & fileName);
extern void C_ZN4QDir5matchERK7QStringS2_(void* arg0, void* arg1); // 4
  // proto:  void QDir::swap(QDir & other);
extern void C_ZN4QDir4swapERS_(void* qthis, void* arg0); // 2
  // proto: static void QDir::setSearchPaths(const QString & prefix, const QStringList & searchPaths);
extern void C_ZN4QDir14setSearchPathsERK7QStringRK11QStringList(void* arg0, void* arg1); // 4
  // proto:  QString QDir::absoluteFilePath(const QString & fileName);
extern void C_ZNK4QDir16absoluteFilePathERK7QString(void* qthis, void* arg0); // 4
  // proto: static QFileInfoList QDir::drives();
extern void C_ZN4QDir6drivesEv(); // 4
  // proto:  bool QDir::rmpath(const QString & dirPath);
extern void C_ZNK4QDir6rmpathERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QDir::path();
extern void C_ZNK4QDir4pathEv(void* qthis); // 4
  // proto: static void QDir::addSearchPath(const QString & prefix, const QString & path);
extern void C_ZN4QDir13addSearchPathERK7QStringS2_(void* arg0, void* arg1); // 4
  // proto:  uint QDir::count();
extern void C_ZNK4QDir5countEv(void* qthis); // 4
  // proto:  bool QDir::cdUp();
extern void C_ZN4QDir4cdUpEv(void* qthis); // 4
  // proto:  SortFlags QDir::sorting();
extern void C_ZNK4QDir7sortingEv(void* qthis); // 4
  // proto:  QStringList QDir::nameFilters();
extern void C_ZNK4QDir11nameFiltersEv(void* qthis); // 4
  // proto: static QDir QDir::temp();
extern void C_ZN4QDir4tempEv(); // 2
  // proto:  bool QDir::mkpath(const QString & dirPath);
extern void C_ZNK4QDir6mkpathERK7QString(void* qthis, void* arg0); // 4
  // proto: static QString QDir::rootPath();
extern void C_ZN4QDir8rootPathEv(); // 4
  // proto:  void QDir::refresh();
extern void C_ZNK4QDir7refreshEv(void* qthis); // 4
  // proto:  bool QDir::makeAbsolute();
extern void C_ZN4QDir12makeAbsoluteEv(void* qthis); // 4
  // proto: static QStringList QDir::nameFiltersFromString(const QString & nameFilter);
extern void C_ZN4QDir21nameFiltersFromStringERK7QString(void* arg0); // 4
  // proto: static QChar QDir::separator();
extern void C_ZN4QDir9separatorEv(); // 4
  // proto:  bool QDir::isReadable();
extern void C_ZNK4QDir10isReadableEv(void* qthis); // 4
  // proto: static QDir QDir::root();
extern void C_ZN4QDir4rootEv(); // 2
  // proto:  bool QDir::isAbsolute();
extern void C_ZNK4QDir10isAbsoluteEv(void* qthis); // 2
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// rename(const class QString &, const class QString &)
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
    // invoke: bool rename(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN4QDir6renameERK7QStringS2_(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "rename", args)
  }

}

// setNameFilters(const class QStringList &)
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
    // invoke: void setNameFilters(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN4QDir14setNameFiltersERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDir", "setNameFilters", args)
  }

}

// ~QDir()
func (this *QDir) FreeQDir(args ...interface{}) () {
  // ~QDir()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDirD0Ev
    // invoke: void ~QDir()
    C.C_ZN4QDirD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDir", "~QDir", args)
  }

}

// exists()
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
    // invoke: bool exists()
    var ret = C.C_ZNK4QDir6existsEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK4QDir6existsERK7QString
    // invoke: bool exists(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK4QDir6existsERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "exists", args)
  }

}

// relativeFilePath(const class QString &)
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
    // invoke: QString relativeFilePath(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK4QDir16relativeFilePathERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "relativeFilePath", args)
  }

}

// currentPath()
func (this *QDir) currentPath_s(args ...interface{}) () {
  // currentPath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir11currentPathEv
    // invoke: QString currentPath()
    var ret = C.C_ZN4QDir11currentPathEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "currentPath", args)
  }

}

// isRelative()
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
    // invoke: bool isRelative()
    var ret = C.C_ZNK4QDir10isRelativeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "isRelative", args)
  }

}

// cd(const class QString &)
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
    // invoke: bool cd(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN4QDir2cdERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "cd", args)
  }

}

// QDir(const class QDir &)
func NewQDir(args ...interface{}) *QDir {
  // QDir(const class QDir &)
  // QDir(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDir{}) // "const QDir &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDirC1ERKS_
    // invoke: void QDir(const class QDir &)
    var arg0 = args[0].(QDir).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN4QDirC2ERKS_(arg0)
    return &QDir{qclsinst:qthis}
  case 1:
    // invoke: _ZN4QDirC1ERK7QString
    // invoke: void QDir(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN4QDirC2ERK7QString(arg0)
    return &QDir{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QDir", "QDir", args)
  }

  return nil // QDir{qclsinst:qthis}
}

// isAbsolutePath(const class QString &)
func (this *QDir) isAbsolutePath_s(args ...interface{}) () {
  // isAbsolutePath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir14isAbsolutePathERK7QString
    // invoke: bool isAbsolutePath(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN4QDir14isAbsolutePathERK7QString(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "isAbsolutePath", args)
  }

}

// remove(const class QString &)
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
    // invoke: bool remove(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN4QDir6removeERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "remove", args)
  }

}

// tempPath()
func (this *QDir) tempPath_s(args ...interface{}) () {
  // tempPath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir8tempPathEv
    // invoke: QString tempPath()
    var ret = C.C_ZN4QDir8tempPathEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "tempPath", args)
  }

}

// homePath()
func (this *QDir) homePath_s(args ...interface{}) () {
  // homePath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir8homePathEv
    // invoke: QString homePath()
    var ret = C.C_ZN4QDir8homePathEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "homePath", args)
  }

}

// home()
func (this *QDir) home_s(args ...interface{}) () {
  // home()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir4homeEv
    // invoke: QDir home()
    var ret = C.C_ZN4QDir4homeEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "home", args)
  }

}

// cleanPath(const class QString &)
func (this *QDir) cleanPath_s(args ...interface{}) () {
  // cleanPath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir9cleanPathERK7QString
    // invoke: QString cleanPath(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN4QDir9cleanPathERK7QString(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "cleanPath", args)
  }

}

// setCurrent(const class QString &)
func (this *QDir) setCurrent_s(args ...interface{}) () {
  // setCurrent(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir10setCurrentERK7QString
    // invoke: bool setCurrent(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN4QDir10setCurrentERK7QString(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "setCurrent", args)
  }

}

// toNativeSeparators(const class QString &)
func (this *QDir) toNativeSeparators_s(args ...interface{}) () {
  // toNativeSeparators(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir18toNativeSeparatorsERK7QString
    // invoke: QString toNativeSeparators(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN4QDir18toNativeSeparatorsERK7QString(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "toNativeSeparators", args)
  }

}

// filter()
func (this *QDir) filter(args ...interface{}) () {
  // filter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QDir6filterEv
    // invoke: Filters filter()
    C.C_ZNK4QDir6filterEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDir", "filter", args)
  }

}

// searchPaths(const class QString &)
func (this *QDir) searchPaths_s(args ...interface{}) () {
  // searchPaths(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir11searchPathsERK7QString
    // invoke: QStringList searchPaths(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN4QDir11searchPathsERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QDir", "searchPaths", args)
  }

}

// isRoot()
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
    // invoke: bool isRoot()
    var ret = C.C_ZNK4QDir6isRootEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "isRoot", args)
  }

}

// filePath(const class QString &)
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
    // invoke: QString filePath(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK4QDir8filePathERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "filePath", args)
  }

}

// mkdir(const class QString &)
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
    // invoke: bool mkdir(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK4QDir5mkdirERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "mkdir", args)
  }

}

// absolutePath()
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
    // invoke: QString absolutePath()
    var ret = C.C_ZNK4QDir12absolutePathEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "absolutePath", args)
  }

}

// addResourceSearchPath(const class QString &)
func (this *QDir) addResourceSearchPath_s(args ...interface{}) () {
  // addResourceSearchPath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir21addResourceSearchPathERK7QString
    // invoke: void addResourceSearchPath(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN4QDir21addResourceSearchPathERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QDir", "addResourceSearchPath", args)
  }

}

// current()
func (this *QDir) current_s(args ...interface{}) () {
  // current()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir7currentEv
    // invoke: QDir current()
    var ret = C.C_ZN4QDir7currentEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "current", args)
  }

}

// fromNativeSeparators(const class QString &)
func (this *QDir) fromNativeSeparators_s(args ...interface{}) () {
  // fromNativeSeparators(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir20fromNativeSeparatorsERK7QString
    // invoke: QString fromNativeSeparators(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN4QDir20fromNativeSeparatorsERK7QString(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "fromNativeSeparators", args)
  }

}

// isRelativePath(const class QString &)
func (this *QDir) isRelativePath_s(args ...interface{}) () {
  // isRelativePath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir14isRelativePathERK7QString
    // invoke: bool isRelativePath(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN4QDir14isRelativePathERK7QString(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "isRelativePath", args)
  }

}

// rmdir(const class QString &)
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
    // invoke: bool rmdir(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK4QDir5rmdirERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "rmdir", args)
  }

}

// canonicalPath()
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
    // invoke: QString canonicalPath()
    var ret = C.C_ZNK4QDir13canonicalPathEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "canonicalPath", args)
  }

}

// removeRecursively()
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
    // invoke: bool removeRecursively()
    var ret = C.C_ZN4QDir17removeRecursivelyEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "removeRecursively", args)
  }

}

// setPath(const class QString &)
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
    // invoke: void setPath(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN4QDir7setPathERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDir", "setPath", args)
  }

}

// dirName()
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
    // invoke: QString dirName()
    var ret = C.C_ZNK4QDir7dirNameEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "dirName", args)
  }

}

// match(const class QStringList &, const class QString &)
func (this *QDir) match_s(args ...interface{}) () {
  // match(const class QStringList &, const class QString &)
  // match(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir5matchERK11QStringListRK7QString
    // invoke: bool match(const class QStringList &, const class QString &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN4QDir5matchERK11QStringListRK7QString(arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN4QDir5matchERK7QStringS2_
    // invoke: bool match(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN4QDir5matchERK7QStringS2_(arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "match", args)
  }

}

// swap(class QDir &)
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
    // invoke: void swap(class QDir &)
    var arg0 = args[0].(QDir).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN4QDir4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDir", "swap", args)
  }

}

// setSearchPaths(const class QString &, const class QStringList &)
func (this *QDir) setSearchPaths_s(args ...interface{}) () {
  // setSearchPaths(const class QString &, const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir14setSearchPathsERK7QStringRK11QStringList
    // invoke: void setSearchPaths(const class QString &, const class QStringList &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStringList).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN4QDir14setSearchPathsERK7QStringRK11QStringList(arg0, arg1)
  default:
    qtrt.ErrorResolve("QDir", "setSearchPaths", args)
  }

}

// absoluteFilePath(const class QString &)
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
    // invoke: QString absoluteFilePath(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK4QDir16absoluteFilePathERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "absoluteFilePath", args)
  }

}

// drives()
func (this *QDir) drives_s(args ...interface{}) () {
  // drives()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir6drivesEv
    // invoke: QFileInfoList drives()
    C.C_ZN4QDir6drivesEv()
  default:
    qtrt.ErrorResolve("QDir", "drives", args)
  }

}

// rmpath(const class QString &)
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
    // invoke: bool rmpath(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK4QDir6rmpathERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "rmpath", args)
  }

}

// path()
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
    // invoke: QString path()
    var ret = C.C_ZNK4QDir4pathEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "path", args)
  }

}

// addSearchPath(const class QString &, const class QString &)
func (this *QDir) addSearchPath_s(args ...interface{}) () {
  // addSearchPath(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir13addSearchPathERK7QStringS2_
    // invoke: void addSearchPath(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN4QDir13addSearchPathERK7QStringS2_(arg0, arg1)
  default:
    qtrt.ErrorResolve("QDir", "addSearchPath", args)
  }

}

// count()
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
    // invoke: uint count()
    var ret = C.C_ZNK4QDir5countEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "count", args)
  }

}

// cdUp()
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
    // invoke: bool cdUp()
    var ret = C.C_ZN4QDir4cdUpEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "cdUp", args)
  }

}

// sorting()
func (this *QDir) sorting(args ...interface{}) () {
  // sorting()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QDir7sortingEv
    // invoke: SortFlags sorting()
    C.C_ZNK4QDir7sortingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDir", "sorting", args)
  }

}

// nameFilters()
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
    // invoke: QStringList nameFilters()
    C.C_ZNK4QDir11nameFiltersEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDir", "nameFilters", args)
  }

}

// temp()
func (this *QDir) temp_s(args ...interface{}) () {
  // temp()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir4tempEv
    // invoke: QDir temp()
    var ret = C.C_ZN4QDir4tempEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "temp", args)
  }

}

// mkpath(const class QString &)
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
    // invoke: bool mkpath(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK4QDir6mkpathERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "mkpath", args)
  }

}

// rootPath()
func (this *QDir) rootPath_s(args ...interface{}) () {
  // rootPath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir8rootPathEv
    // invoke: QString rootPath()
    var ret = C.C_ZN4QDir8rootPathEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "rootPath", args)
  }

}

// refresh()
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
    // invoke: void refresh()
    C.C_ZNK4QDir7refreshEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDir", "refresh", args)
  }

}

// makeAbsolute()
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
    // invoke: bool makeAbsolute()
    var ret = C.C_ZN4QDir12makeAbsoluteEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "makeAbsolute", args)
  }

}

// nameFiltersFromString(const class QString &)
func (this *QDir) nameFiltersFromString_s(args ...interface{}) () {
  // nameFiltersFromString(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir21nameFiltersFromStringERK7QString
    // invoke: QStringList nameFiltersFromString(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN4QDir21nameFiltersFromStringERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QDir", "nameFiltersFromString", args)
  }

}

// separator()
func (this *QDir) separator_s(args ...interface{}) () {
  // separator()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir9separatorEv
    // invoke: QChar separator()
    var ret = C.C_ZN4QDir9separatorEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "separator", args)
  }

}

// isReadable()
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
    // invoke: bool isReadable()
    var ret = C.C_ZNK4QDir10isReadableEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "isReadable", args)
  }

}

// root()
func (this *QDir) root_s(args ...interface{}) () {
  // root()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QDir4rootEv
    // invoke: QDir root()
    var ret = C.C_ZN4QDir4rootEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "root", args)
  }

}

// isAbsolute()
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
    // invoke: bool isAbsolute()
    var ret = C.C_ZNK4QDir10isAbsoluteEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDir", "isAbsolute", args)
  }

}

// <= body block end

