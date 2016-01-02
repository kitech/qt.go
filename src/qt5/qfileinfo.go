package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtCore/qfileinfo.h
// dst-file: /src/core/qfileinfo.go
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
  // proto:  bool QFileInfo::isHidden();
extern void _ZNK9QFileInfo8isHiddenEv(void* qthis);
  // proto:  void QFileInfo::QFileInfo();
extern void* dector_ZN9QFileInfoC1Ev();
extern void _ZN9QFileInfoC1Ev(void* qthis);
  // proto:  QString QFileInfo::completeSuffix();
extern void _ZNK9QFileInfo14completeSuffixEv(void* qthis);
  // proto:  QString QFileInfo::canonicalPath();
extern void _ZNK9QFileInfo13canonicalPathEv(void* qthis);
  // proto: static bool QFileInfo::exists(const QString & file);
extern void _ZN9QFileInfo6existsERK7QString(void* arg0);
  // proto:  bool QFileInfo::makeAbsolute();
extern void _ZN9QFileInfo12makeAbsoluteEv(void* qthis);
  // proto:  bool QFileInfo::isRoot();
extern void _ZNK9QFileInfo6isRootEv(void* qthis);
  // proto:  QString QFileInfo::canonicalFilePath();
extern void _ZNK9QFileInfo17canonicalFilePathEv(void* qthis);
  // proto:  bool QFileInfo::isDir();
extern void _ZNK9QFileInfo5isDirEv(void* qthis);
  // proto:  void QFileInfo::QFileInfo(const QString & file);
extern void* dector_ZN9QFileInfoC1ERK7QString(void* arg0);
extern void _ZN9QFileInfoC1ERK7QString(void* qthis, void* arg0);
  // proto:  QString QFileInfo::symLinkTarget();
extern void demth_ZNK9QFileInfo13symLinkTargetEv(void* qthis);
  // proto:  void QFileInfo::setFile(const QString & file);
extern void _ZN9QFileInfo7setFileERK7QString(void* qthis, void* arg0);
  // proto:  bool QFileInfo::isFile();
extern void _ZNK9QFileInfo6isFileEv(void* qthis);
  // proto:  void QFileInfo::QFileInfo(const QFile & file);
extern void* dector_ZN9QFileInfoC1ERK5QFile(void* arg0);
extern void _ZN9QFileInfoC1ERK5QFile(void* qthis, void* arg0);
  // proto:  void QFileInfo::setFile(const QFile & file);
extern void _ZN9QFileInfo7setFileERK5QFile(void* qthis, void* arg0);
  // proto:  uint QFileInfo::ownerId();
extern void _ZNK9QFileInfo7ownerIdEv(void* qthis);
  // proto:  QString QFileInfo::readLink();
extern void _ZNK9QFileInfo8readLinkEv(void* qthis);
  // proto:  QString QFileInfo::filePath();
extern void _ZNK9QFileInfo8filePathEv(void* qthis);
  // proto:  void QFileInfo::QFileInfo(const QFileInfo & fileinfo);
extern void* dector_ZN9QFileInfoC1ERKS_(void* arg0);
extern void _ZN9QFileInfoC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QFileInfo::isSymLink();
extern void _ZNK9QFileInfo9isSymLinkEv(void* qthis);
  // proto:  QDateTime QFileInfo::lastRead();
extern void _ZNK9QFileInfo8lastReadEv(void* qthis);
  // proto:  void QFileInfo::refresh();
extern void _ZN9QFileInfo7refreshEv(void* qthis);
  // proto:  void QFileInfo::QFileInfo(const QDir & dir, const QString & file);
extern void* dector_ZN9QFileInfoC1ERK4QDirRK7QString(void* arg0, void* arg1);
extern void _ZN9QFileInfoC1ERK4QDirRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  QString QFileInfo::path();
extern void _ZNK9QFileInfo4pathEv(void* qthis);
  // proto:  QDir QFileInfo::absoluteDir();
extern void _ZNK9QFileInfo11absoluteDirEv(void* qthis);
  // proto:  bool QFileInfo::isBundle();
extern void _ZNK9QFileInfo8isBundleEv(void* qthis);
  // proto:  void QFileInfo::setFile(const QDir & dir, const QString & file);
extern void _ZN9QFileInfo7setFileERK4QDirRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  bool QFileInfo::isRelative();
extern void _ZNK9QFileInfo10isRelativeEv(void* qthis);
  // proto:  void QFileInfo::setCaching(bool on);
extern void _ZN9QFileInfo10setCachingEb(void* qthis, bool arg0);
  // proto:  QDateTime QFileInfo::created();
extern void _ZNK9QFileInfo7createdEv(void* qthis);
  // proto:  bool QFileInfo::caching();
extern void _ZNK9QFileInfo7cachingEv(void* qthis);
  // proto:  void QFileInfo::~QFileInfo();
extern void _ZN9QFileInfoD0Ev(void* qthis);
  // proto:  QString QFileInfo::completeBaseName();
extern void _ZNK9QFileInfo16completeBaseNameEv(void* qthis);
  // proto:  QString QFileInfo::baseName();
extern void _ZNK9QFileInfo8baseNameEv(void* qthis);
  // proto:  bool QFileInfo::isExecutable();
extern void _ZNK9QFileInfo12isExecutableEv(void* qthis);
  // proto:  QString QFileInfo::bundleName();
extern void _ZNK9QFileInfo10bundleNameEv(void* qthis);
  // proto:  uint QFileInfo::groupId();
extern void _ZNK9QFileInfo7groupIdEv(void* qthis);
  // proto:  QString QFileInfo::fileName();
extern void _ZNK9QFileInfo8fileNameEv(void* qthis);
  // proto:  qint64 QFileInfo::size();
extern void _ZNK9QFileInfo4sizeEv(void* qthis);
  // proto:  QString QFileInfo::absoluteFilePath();
extern void _ZNK9QFileInfo16absoluteFilePathEv(void* qthis);
  // proto:  QString QFileInfo::suffix();
extern void _ZNK9QFileInfo6suffixEv(void* qthis);
  // proto:  QString QFileInfo::group();
extern void _ZNK9QFileInfo5groupEv(void* qthis);
  // proto:  bool QFileInfo::isAbsolute();
extern void demth_ZNK9QFileInfo10isAbsoluteEv(void* qthis);
  // proto:  bool QFileInfo::isNativePath();
extern void _ZNK9QFileInfo12isNativePathEv(void* qthis);
  // proto:  bool QFileInfo::isWritable();
extern void _ZNK9QFileInfo10isWritableEv(void* qthis);
  // proto:  QString QFileInfo::owner();
extern void _ZNK9QFileInfo5ownerEv(void* qthis);
  // proto:  bool QFileInfo::isReadable();
extern void _ZNK9QFileInfo10isReadableEv(void* qthis);
  // proto:  QDir QFileInfo::dir();
extern void _ZNK9QFileInfo3dirEv(void* qthis);
  // proto:  void QFileInfo::swap(QFileInfo & other);
extern void demth_ZN9QFileInfo4swapERS_(void* qthis, void* arg0);
  // proto:  bool QFileInfo::exists();
extern void _ZNK9QFileInfo6existsEv(void* qthis);
  // proto:  QDateTime QFileInfo::lastModified();
extern void _ZNK9QFileInfo12lastModifiedEv(void* qthis);
  // proto:  QString QFileInfo::absolutePath();
extern void _ZNK9QFileInfo12absolutePathEv(void* qthis);
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

// class sizeof(QFileInfo)=1
type QFileInfo struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  bool QFileInfo::isHidden();
func (this *QFileInfo) isHidden(args ...interface{}) () {
  // isHidden()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo8isHiddenEv
  default:
    qtrt.ErrorResolve("QFileInfo", "isHidden", args)
  }

}

  // proto:  void QFileInfo::QFileInfo();
func NewQFileInfo(args ...interface{}) QFileInfo {
  return QFileInfo{}
}

  // proto:  QString QFileInfo::completeSuffix();
func (this *QFileInfo) completeSuffix(args ...interface{}) () {
  // completeSuffix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo14completeSuffixEv
  default:
    qtrt.ErrorResolve("QFileInfo", "completeSuffix", args)
  }

}

  // proto:  QString QFileInfo::canonicalPath();
func (this *QFileInfo) canonicalPath(args ...interface{}) () {
  // canonicalPath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo13canonicalPathEv
  default:
    qtrt.ErrorResolve("QFileInfo", "canonicalPath", args)
  }

}

  // proto: static bool QFileInfo::exists(const QString & file);
func (this *QFileInfo) exists_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFileInfo", "exists", args)
  }

}

  // proto:  bool QFileInfo::makeAbsolute();
func (this *QFileInfo) makeAbsolute(args ...interface{}) () {
  // makeAbsolute()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QFileInfo12makeAbsoluteEv
  default:
    qtrt.ErrorResolve("QFileInfo", "makeAbsolute", args)
  }

}

  // proto:  bool QFileInfo::isRoot();
func (this *QFileInfo) isRoot(args ...interface{}) () {
  // isRoot()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo6isRootEv
  default:
    qtrt.ErrorResolve("QFileInfo", "isRoot", args)
  }

}

  // proto:  QString QFileInfo::canonicalFilePath();
func (this *QFileInfo) canonicalFilePath(args ...interface{}) () {
  // canonicalFilePath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo17canonicalFilePathEv
  default:
    qtrt.ErrorResolve("QFileInfo", "canonicalFilePath", args)
  }

}

  // proto:  bool QFileInfo::isDir();
func (this *QFileInfo) isDir(args ...interface{}) () {
  // isDir()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo5isDirEv
  default:
    qtrt.ErrorResolve("QFileInfo", "isDir", args)
  }

}

  // proto:  QString QFileInfo::symLinkTarget();
func (this *QFileInfo) symLinkTarget(args ...interface{}) () {
  // symLinkTarget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo13symLinkTargetEv
  default:
    qtrt.ErrorResolve("QFileInfo", "symLinkTarget", args)
  }

}

  // proto:  void QFileInfo::setFile(const QString & file);
func (this *QFileInfo) setFile(args ...interface{}) () {
  // setFile(const class QString &)
  // setFile(const class QFile &)
  // setFile(const class QDir &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QFile{}) // "const QFile &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QDir{}) // "const QDir &"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QFileInfo7setFileERK7QString
  case 1:
    // invoke: _ZN9QFileInfo7setFileERK5QFile
  case 2:
    // invoke: _ZN9QFileInfo7setFileERK4QDirRK7QString
  default:
    qtrt.ErrorResolve("QFileInfo", "setFile", args)
  }

}

  // proto:  bool QFileInfo::isFile();
func (this *QFileInfo) isFile(args ...interface{}) () {
  // isFile()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo6isFileEv
  default:
    qtrt.ErrorResolve("QFileInfo", "isFile", args)
  }

}

  // proto:  uint QFileInfo::ownerId();
func (this *QFileInfo) ownerId(args ...interface{}) () {
  // ownerId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo7ownerIdEv
  default:
    qtrt.ErrorResolve("QFileInfo", "ownerId", args)
  }

}

  // proto:  QString QFileInfo::readLink();
func (this *QFileInfo) readLink(args ...interface{}) () {
  // readLink()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo8readLinkEv
  default:
    qtrt.ErrorResolve("QFileInfo", "readLink", args)
  }

}

  // proto:  QString QFileInfo::filePath();
func (this *QFileInfo) filePath(args ...interface{}) () {
  // filePath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo8filePathEv
  default:
    qtrt.ErrorResolve("QFileInfo", "filePath", args)
  }

}

  // proto:  bool QFileInfo::isSymLink();
func (this *QFileInfo) isSymLink(args ...interface{}) () {
  // isSymLink()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo9isSymLinkEv
  default:
    qtrt.ErrorResolve("QFileInfo", "isSymLink", args)
  }

}

  // proto:  QDateTime QFileInfo::lastRead();
func (this *QFileInfo) lastRead(args ...interface{}) () {
  // lastRead()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo8lastReadEv
  default:
    qtrt.ErrorResolve("QFileInfo", "lastRead", args)
  }

}

  // proto:  void QFileInfo::refresh();
func (this *QFileInfo) refresh(args ...interface{}) () {
  // refresh()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QFileInfo7refreshEv
  default:
    qtrt.ErrorResolve("QFileInfo", "refresh", args)
  }

}

  // proto:  QString QFileInfo::path();
func (this *QFileInfo) path(args ...interface{}) () {
  // path()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo4pathEv
  default:
    qtrt.ErrorResolve("QFileInfo", "path", args)
  }

}

  // proto:  QDir QFileInfo::absoluteDir();
func (this *QFileInfo) absoluteDir(args ...interface{}) () {
  // absoluteDir()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo11absoluteDirEv
  default:
    qtrt.ErrorResolve("QFileInfo", "absoluteDir", args)
  }

}

  // proto:  bool QFileInfo::isBundle();
func (this *QFileInfo) isBundle(args ...interface{}) () {
  // isBundle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo8isBundleEv
  default:
    qtrt.ErrorResolve("QFileInfo", "isBundle", args)
  }

}

  // proto:  bool QFileInfo::isRelative();
func (this *QFileInfo) isRelative(args ...interface{}) () {
  // isRelative()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo10isRelativeEv
  default:
    qtrt.ErrorResolve("QFileInfo", "isRelative", args)
  }

}

  // proto:  void QFileInfo::setCaching(bool on);
func (this *QFileInfo) setCaching(args ...interface{}) () {
  // setCaching(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QFileInfo10setCachingEb
  default:
    qtrt.ErrorResolve("QFileInfo", "setCaching", args)
  }

}

  // proto:  QDateTime QFileInfo::created();
func (this *QFileInfo) created(args ...interface{}) () {
  // created()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo7createdEv
  default:
    qtrt.ErrorResolve("QFileInfo", "created", args)
  }

}

  // proto:  bool QFileInfo::caching();
func (this *QFileInfo) caching(args ...interface{}) () {
  // caching()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo7cachingEv
  default:
    qtrt.ErrorResolve("QFileInfo", "caching", args)
  }

}

  // proto:  void QFileInfo::~QFileInfo();
func (this *QFileInfo) FreeQFileInfo(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFileInfo", "~QFileInfo", args)
  }

}

  // proto:  QString QFileInfo::completeBaseName();
func (this *QFileInfo) completeBaseName(args ...interface{}) () {
  // completeBaseName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo16completeBaseNameEv
  default:
    qtrt.ErrorResolve("QFileInfo", "completeBaseName", args)
  }

}

  // proto:  QString QFileInfo::baseName();
func (this *QFileInfo) baseName(args ...interface{}) () {
  // baseName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo8baseNameEv
  default:
    qtrt.ErrorResolve("QFileInfo", "baseName", args)
  }

}

  // proto:  bool QFileInfo::isExecutable();
func (this *QFileInfo) isExecutable(args ...interface{}) () {
  // isExecutable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo12isExecutableEv
  default:
    qtrt.ErrorResolve("QFileInfo", "isExecutable", args)
  }

}

  // proto:  QString QFileInfo::bundleName();
func (this *QFileInfo) bundleName(args ...interface{}) () {
  // bundleName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo10bundleNameEv
  default:
    qtrt.ErrorResolve("QFileInfo", "bundleName", args)
  }

}

  // proto:  uint QFileInfo::groupId();
func (this *QFileInfo) groupId(args ...interface{}) () {
  // groupId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo7groupIdEv
  default:
    qtrt.ErrorResolve("QFileInfo", "groupId", args)
  }

}

  // proto:  QString QFileInfo::fileName();
func (this *QFileInfo) fileName(args ...interface{}) () {
  // fileName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo8fileNameEv
  default:
    qtrt.ErrorResolve("QFileInfo", "fileName", args)
  }

}

  // proto:  qint64 QFileInfo::size();
func (this *QFileInfo) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo4sizeEv
  default:
    qtrt.ErrorResolve("QFileInfo", "size", args)
  }

}

  // proto:  QString QFileInfo::absoluteFilePath();
func (this *QFileInfo) absoluteFilePath(args ...interface{}) () {
  // absoluteFilePath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo16absoluteFilePathEv
  default:
    qtrt.ErrorResolve("QFileInfo", "absoluteFilePath", args)
  }

}

  // proto:  QString QFileInfo::suffix();
func (this *QFileInfo) suffix(args ...interface{}) () {
  // suffix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo6suffixEv
  default:
    qtrt.ErrorResolve("QFileInfo", "suffix", args)
  }

}

  // proto:  QString QFileInfo::group();
func (this *QFileInfo) group(args ...interface{}) () {
  // group()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo5groupEv
  default:
    qtrt.ErrorResolve("QFileInfo", "group", args)
  }

}

  // proto:  bool QFileInfo::isAbsolute();
func (this *QFileInfo) isAbsolute(args ...interface{}) () {
  // isAbsolute()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo10isAbsoluteEv
  default:
    qtrt.ErrorResolve("QFileInfo", "isAbsolute", args)
  }

}

  // proto:  bool QFileInfo::isNativePath();
func (this *QFileInfo) isNativePath(args ...interface{}) () {
  // isNativePath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo12isNativePathEv
  default:
    qtrt.ErrorResolve("QFileInfo", "isNativePath", args)
  }

}

  // proto:  bool QFileInfo::isWritable();
func (this *QFileInfo) isWritable(args ...interface{}) () {
  // isWritable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo10isWritableEv
  default:
    qtrt.ErrorResolve("QFileInfo", "isWritable", args)
  }

}

  // proto:  QString QFileInfo::owner();
func (this *QFileInfo) owner(args ...interface{}) () {
  // owner()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo5ownerEv
  default:
    qtrt.ErrorResolve("QFileInfo", "owner", args)
  }

}

  // proto:  bool QFileInfo::isReadable();
func (this *QFileInfo) isReadable(args ...interface{}) () {
  // isReadable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo10isReadableEv
  default:
    qtrt.ErrorResolve("QFileInfo", "isReadable", args)
  }

}

  // proto:  QDir QFileInfo::dir();
func (this *QFileInfo) dir(args ...interface{}) () {
  // dir()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo3dirEv
  default:
    qtrt.ErrorResolve("QFileInfo", "dir", args)
  }

}

  // proto:  void QFileInfo::swap(QFileInfo & other);
func (this *QFileInfo) swap(args ...interface{}) () {
  // swap(class QFileInfo &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFileInfo{}) // "QFileInfo &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QFileInfo4swapERS_
  default:
    qtrt.ErrorResolve("QFileInfo", "swap", args)
  }

}

  // proto:  bool QFileInfo::exists();
func (this *QFileInfo) exists(args ...interface{}) () {
  // exists(const class QString &)
  // exists()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QFileInfo6existsERK7QString
  case 1:
    // invoke: _ZNK9QFileInfo6existsEv
  default:
    qtrt.ErrorResolve("QFileInfo", "exists", args)
  }

}

  // proto:  QDateTime QFileInfo::lastModified();
func (this *QFileInfo) lastModified(args ...interface{}) () {
  // lastModified()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo12lastModifiedEv
  default:
    qtrt.ErrorResolve("QFileInfo", "lastModified", args)
  }

}

  // proto:  QString QFileInfo::absolutePath();
func (this *QFileInfo) absolutePath(args ...interface{}) () {
  // absolutePath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo12absolutePathEv
  default:
    qtrt.ErrorResolve("QFileInfo", "absolutePath", args)
  }

}

// <= body block end

