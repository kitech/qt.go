package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QDir QFileInfo::absoluteDir();
extern void C_ZNK9QFileInfo11absoluteDirEv(void* qthis); // 4
  // proto:  QString QFileInfo::suffix();
extern void C_ZNK9QFileInfo6suffixEv(void* qthis); // 4
  // proto:  void QFileInfo::QFileInfo(const QFileInfo & fileinfo);
extern void* C_ZN9QFileInfoC2ERKS_(void* arg0); // 3
  // proto:  void QFileInfo::QFileInfo();
extern void* C_ZN9QFileInfoC2Ev(); // 3
  // proto:  void QFileInfo::QFileInfo(const QString & file);
extern void* C_ZN9QFileInfoC2ERK7QString(void* arg0); // 3
  // proto:  void QFileInfo::QFileInfo(const QFile & file);
extern void* C_ZN9QFileInfoC2ERK5QFile(void* arg0); // 3
  // proto:  void QFileInfo::QFileInfo(const QDir & dir, const QString & file);
extern void* C_ZN9QFileInfoC2ERK4QDirRK7QString(void* arg0, void* arg1); // 3
  // proto:  QString QFileInfo::symLinkTarget();
extern void C_ZNK9QFileInfo13symLinkTargetEv(void* qthis); // 2
  // proto:  bool QFileInfo::isRelative();
extern void C_ZNK9QFileInfo10isRelativeEv(void* qthis); // 4
  // proto:  QString QFileInfo::completeBaseName();
extern void C_ZNK9QFileInfo16completeBaseNameEv(void* qthis); // 4
  // proto:  QString QFileInfo::canonicalFilePath();
extern void C_ZNK9QFileInfo17canonicalFilePathEv(void* qthis); // 4
  // proto:  bool QFileInfo::isAbsolute();
extern void C_ZNK9QFileInfo10isAbsoluteEv(void* qthis); // 2
  // proto:  QString QFileInfo::owner();
extern void C_ZNK9QFileInfo5ownerEv(void* qthis); // 4
  // proto:  bool QFileInfo::isExecutable();
extern void C_ZNK9QFileInfo12isExecutableEv(void* qthis); // 4
  // proto:  qint64 QFileInfo::size();
extern void C_ZNK9QFileInfo4sizeEv(void* qthis); // 4
  // proto:  QString QFileInfo::bundleName();
extern void C_ZNK9QFileInfo10bundleNameEv(void* qthis); // 4
  // proto:  QString QFileInfo::group();
extern void C_ZNK9QFileInfo5groupEv(void* qthis); // 4
  // proto: static bool QFileInfo::exists(const QString & file);
extern void C_ZN9QFileInfo6existsERK7QString(void* arg0); // 4
  // proto:  bool QFileInfo::exists();
extern void C_ZNK9QFileInfo6existsEv(void* qthis); // 4
  // proto:  bool QFileInfo::isWritable();
extern void C_ZNK9QFileInfo10isWritableEv(void* qthis); // 4
  // proto:  QString QFileInfo::filePath();
extern void C_ZNK9QFileInfo8filePathEv(void* qthis); // 4
  // proto:  QString QFileInfo::absolutePath();
extern void C_ZNK9QFileInfo12absolutePathEv(void* qthis); // 4
  // proto:  void QFileInfo::swap(QFileInfo & other);
extern void C_ZN9QFileInfo4swapERS_(void* qthis, void* arg0); // 2
  // proto:  QString QFileInfo::canonicalPath();
extern void C_ZNK9QFileInfo13canonicalPathEv(void* qthis); // 4
  // proto:  bool QFileInfo::isBundle();
extern void C_ZNK9QFileInfo8isBundleEv(void* qthis); // 4
  // proto:  bool QFileInfo::isHidden();
extern void C_ZNK9QFileInfo8isHiddenEv(void* qthis); // 4
  // proto:  bool QFileInfo::isDir();
extern void C_ZNK9QFileInfo5isDirEv(void* qthis); // 4
  // proto:  QDateTime QFileInfo::lastRead();
extern void C_ZNK9QFileInfo8lastReadEv(void* qthis); // 4
  // proto:  bool QFileInfo::isRoot();
extern void C_ZNK9QFileInfo6isRootEv(void* qthis); // 4
  // proto:  QString QFileInfo::absoluteFilePath();
extern void C_ZNK9QFileInfo16absoluteFilePathEv(void* qthis); // 4
  // proto:  QString QFileInfo::fileName();
extern void C_ZNK9QFileInfo8fileNameEv(void* qthis); // 4
  // proto:  void QFileInfo::setCaching(bool on);
extern void C_ZN9QFileInfo10setCachingEb(void* qthis, bool arg0); // 4
  // proto:  QString QFileInfo::completeSuffix();
extern void C_ZNK9QFileInfo14completeSuffixEv(void* qthis); // 4
  // proto:  QString QFileInfo::path();
extern void C_ZNK9QFileInfo4pathEv(void* qthis); // 4
  // proto:  void QFileInfo::~QFileInfo();
extern void C_ZN9QFileInfoD2Ev(void* qthis); // 4
  // proto:  uint QFileInfo::groupId();
extern void C_ZNK9QFileInfo7groupIdEv(void* qthis); // 4
  // proto:  QFile::Permissions QFileInfo::permissions();
extern void C_ZNK9QFileInfo11permissionsEv(void* qthis); // 4
  // proto:  bool QFileInfo::isNativePath();
extern void C_ZNK9QFileInfo12isNativePathEv(void* qthis); // 4
  // proto:  QDateTime QFileInfo::created();
extern void C_ZNK9QFileInfo7createdEv(void* qthis); // 4
  // proto:  bool QFileInfo::isSymLink();
extern void C_ZNK9QFileInfo9isSymLinkEv(void* qthis); // 4
  // proto:  QDateTime QFileInfo::lastModified();
extern void C_ZNK9QFileInfo12lastModifiedEv(void* qthis); // 4
  // proto:  QString QFileInfo::baseName();
extern void C_ZNK9QFileInfo8baseNameEv(void* qthis); // 4
  // proto:  void QFileInfo::refresh();
extern void C_ZN9QFileInfo7refreshEv(void* qthis); // 4
  // proto:  QString QFileInfo::readLink();
extern void C_ZNK9QFileInfo8readLinkEv(void* qthis); // 4
  // proto:  void QFileInfo::setFile(const QDir & dir, const QString & file);
extern void C_ZN9QFileInfo7setFileERK4QDirRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QFileInfo::setFile(const QString & file);
extern void C_ZN9QFileInfo7setFileERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QFileInfo::setFile(const QFile & file);
extern void C_ZN9QFileInfo7setFileERK5QFile(void* qthis, void* arg0); // 4
  // proto:  bool QFileInfo::isReadable();
extern void C_ZNK9QFileInfo10isReadableEv(void* qthis); // 4
  // proto:  bool QFileInfo::caching();
extern void C_ZNK9QFileInfo7cachingEv(void* qthis); // 4
  // proto:  uint QFileInfo::ownerId();
extern void C_ZNK9QFileInfo7ownerIdEv(void* qthis); // 4
  // proto:  bool QFileInfo::isFile();
extern void C_ZNK9QFileInfo6isFileEv(void* qthis); // 4
  // proto:  bool QFileInfo::makeAbsolute();
extern void C_ZN9QFileInfo12makeAbsoluteEv(void* qthis); // 4
  // proto:  QDir QFileInfo::dir();
extern void C_ZNK9QFileInfo3dirEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// absoluteDir()
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
    // invoke: QDir absoluteDir()
    var ret = C.C_ZNK9QFileInfo11absoluteDirEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "absoluteDir", args)
  }

}

// suffix()
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
    // invoke: QString suffix()
    var ret = C.C_ZNK9QFileInfo6suffixEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "suffix", args)
  }

}

// QFileInfo(const class QFileInfo &)
func NewQFileInfo(args ...interface{}) *QFileInfo {
  // QFileInfo(const class QFileInfo &)
  // QFileInfo()
  // QFileInfo(const class QString &)
  // QFileInfo(const class QFile &)
  // QFileInfo(const class QDir &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFileInfo{}) // "const QFileInfo &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QFile{}) // "const QFile &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QDir{}) // "const QDir &"
  vtys[4][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QFileInfoC1ERKS_
    // invoke: void QFileInfo(const class QFileInfo &)
    var arg0 = args[0].(QFileInfo).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QFileInfoC2ERKS_(arg0)
    return &QFileInfo{qclsinst:qthis}
  case 1:
    // invoke: _ZN9QFileInfoC1Ev
    // invoke: void QFileInfo()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QFileInfoC2Ev()
    return &QFileInfo{qclsinst:qthis}
  case 2:
    // invoke: _ZN9QFileInfoC1ERK7QString
    // invoke: void QFileInfo(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QFileInfoC2ERK7QString(arg0)
    return &QFileInfo{qclsinst:qthis}
  case 3:
    // invoke: _ZN9QFileInfoC1ERK5QFile
    // invoke: void QFileInfo(const class QFile &)
    var arg0 = args[0].(QFile).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QFileInfoC2ERK5QFile(arg0)
    return &QFileInfo{qclsinst:qthis}
  case 4:
    // invoke: _ZN9QFileInfoC1ERK4QDirRK7QString
    // invoke: void QFileInfo(const class QDir &, const class QString &)
    var arg0 = args[0].(QDir).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QFileInfoC2ERK4QDirRK7QString(arg0, arg1)
    return &QFileInfo{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QFileInfo", "QFileInfo", args)
  }

  return nil // QFileInfo{qclsinst:qthis}
}

// symLinkTarget()
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
    // invoke: QString symLinkTarget()
    var ret = C.C_ZNK9QFileInfo13symLinkTargetEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "symLinkTarget", args)
  }

}

// isRelative()
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
    // invoke: bool isRelative()
    var ret = C.C_ZNK9QFileInfo10isRelativeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "isRelative", args)
  }

}

// completeBaseName()
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
    // invoke: QString completeBaseName()
    var ret = C.C_ZNK9QFileInfo16completeBaseNameEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "completeBaseName", args)
  }

}

// canonicalFilePath()
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
    // invoke: QString canonicalFilePath()
    var ret = C.C_ZNK9QFileInfo17canonicalFilePathEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "canonicalFilePath", args)
  }

}

// isAbsolute()
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
    // invoke: bool isAbsolute()
    var ret = C.C_ZNK9QFileInfo10isAbsoluteEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "isAbsolute", args)
  }

}

// owner()
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
    // invoke: QString owner()
    var ret = C.C_ZNK9QFileInfo5ownerEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "owner", args)
  }

}

// isExecutable()
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
    // invoke: bool isExecutable()
    var ret = C.C_ZNK9QFileInfo12isExecutableEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "isExecutable", args)
  }

}

// size()
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
    // invoke: qint64 size()
    var ret = C.C_ZNK9QFileInfo4sizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "size", args)
  }

}

// bundleName()
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
    // invoke: QString bundleName()
    var ret = C.C_ZNK9QFileInfo10bundleNameEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "bundleName", args)
  }

}

// group()
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
    // invoke: QString group()
    var ret = C.C_ZNK9QFileInfo5groupEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "group", args)
  }

}

// exists(const class QString &)
func (this *QFileInfo) exists_s(args ...interface{}) () {
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
    // invoke: bool exists(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN9QFileInfo6existsERK7QString(arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK9QFileInfo6existsEv
    // invoke: bool exists()
    var ret = C.C_ZNK9QFileInfo6existsEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "exists", args)
  }

}

// isWritable()
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
    // invoke: bool isWritable()
    var ret = C.C_ZNK9QFileInfo10isWritableEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "isWritable", args)
  }

}

// filePath()
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
    // invoke: QString filePath()
    var ret = C.C_ZNK9QFileInfo8filePathEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "filePath", args)
  }

}

// absolutePath()
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
    // invoke: QString absolutePath()
    var ret = C.C_ZNK9QFileInfo12absolutePathEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "absolutePath", args)
  }

}

// swap(class QFileInfo &)
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
    // invoke: void swap(class QFileInfo &)
    var arg0 = args[0].(QFileInfo).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QFileInfo4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileInfo", "swap", args)
  }

}

// canonicalPath()
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
    // invoke: QString canonicalPath()
    var ret = C.C_ZNK9QFileInfo13canonicalPathEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "canonicalPath", args)
  }

}

// isBundle()
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
    // invoke: bool isBundle()
    var ret = C.C_ZNK9QFileInfo8isBundleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "isBundle", args)
  }

}

// isHidden()
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
    // invoke: bool isHidden()
    var ret = C.C_ZNK9QFileInfo8isHiddenEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "isHidden", args)
  }

}

// isDir()
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
    // invoke: bool isDir()
    var ret = C.C_ZNK9QFileInfo5isDirEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "isDir", args)
  }

}

// lastRead()
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
    // invoke: QDateTime lastRead()
    var ret = C.C_ZNK9QFileInfo8lastReadEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "lastRead", args)
  }

}

// isRoot()
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
    // invoke: bool isRoot()
    var ret = C.C_ZNK9QFileInfo6isRootEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "isRoot", args)
  }

}

// absoluteFilePath()
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
    // invoke: QString absoluteFilePath()
    var ret = C.C_ZNK9QFileInfo16absoluteFilePathEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "absoluteFilePath", args)
  }

}

// fileName()
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
    // invoke: QString fileName()
    var ret = C.C_ZNK9QFileInfo8fileNameEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "fileName", args)
  }

}

// setCaching(_Bool)
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
    // invoke: void setCaching(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QFileInfo10setCachingEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileInfo", "setCaching", args)
  }

}

// completeSuffix()
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
    // invoke: QString completeSuffix()
    var ret = C.C_ZNK9QFileInfo14completeSuffixEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "completeSuffix", args)
  }

}

// path()
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
    // invoke: QString path()
    var ret = C.C_ZNK9QFileInfo4pathEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "path", args)
  }

}

// ~QFileInfo()
func (this *QFileInfo) FreeQFileInfo(args ...interface{}) () {
  // ~QFileInfo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QFileInfoD0Ev
    // invoke: void ~QFileInfo()
    C.C_ZN9QFileInfoD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileInfo", "~QFileInfo", args)
  }

}

// groupId()
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
    // invoke: uint groupId()
    var ret = C.C_ZNK9QFileInfo7groupIdEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "groupId", args)
  }

}

// permissions()
func (this *QFileInfo) permissions(args ...interface{}) () {
  // permissions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo11permissionsEv
    // invoke: QFile::Permissions permissions()
    C.C_ZNK9QFileInfo11permissionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileInfo", "permissions", args)
  }

}

// isNativePath()
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
    // invoke: bool isNativePath()
    var ret = C.C_ZNK9QFileInfo12isNativePathEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "isNativePath", args)
  }

}

// created()
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
    // invoke: QDateTime created()
    var ret = C.C_ZNK9QFileInfo7createdEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "created", args)
  }

}

// isSymLink()
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
    // invoke: bool isSymLink()
    var ret = C.C_ZNK9QFileInfo9isSymLinkEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "isSymLink", args)
  }

}

// lastModified()
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
    // invoke: QDateTime lastModified()
    var ret = C.C_ZNK9QFileInfo12lastModifiedEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "lastModified", args)
  }

}

// baseName()
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
    // invoke: QString baseName()
    var ret = C.C_ZNK9QFileInfo8baseNameEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "baseName", args)
  }

}

// refresh()
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
    // invoke: void refresh()
    C.C_ZN9QFileInfo7refreshEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileInfo", "refresh", args)
  }

}

// readLink()
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
    // invoke: QString readLink()
    var ret = C.C_ZNK9QFileInfo8readLinkEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "readLink", args)
  }

}

// setFile(const class QDir &, const class QString &)
func (this *QFileInfo) setFile(args ...interface{}) () {
  // setFile(const class QDir &, const class QString &)
  // setFile(const class QString &)
  // setFile(const class QFile &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDir{}) // "const QDir &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QFile{}) // "const QFile &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QFileInfo7setFileERK4QDirRK7QString
    // invoke: void setFile(const class QDir &, const class QString &)
    var arg0 = args[0].(QDir).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN9QFileInfo7setFileERK4QDirRK7QString(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN9QFileInfo7setFileERK7QString
    // invoke: void setFile(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QFileInfo7setFileERK7QString(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN9QFileInfo7setFileERK5QFile
    // invoke: void setFile(const class QFile &)
    var arg0 = args[0].(QFile).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QFileInfo7setFileERK5QFile(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileInfo", "setFile", args)
  }

}

// isReadable()
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
    // invoke: bool isReadable()
    var ret = C.C_ZNK9QFileInfo10isReadableEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "isReadable", args)
  }

}

// caching()
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
    // invoke: bool caching()
    var ret = C.C_ZNK9QFileInfo7cachingEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "caching", args)
  }

}

// ownerId()
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
    // invoke: uint ownerId()
    var ret = C.C_ZNK9QFileInfo7ownerIdEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "ownerId", args)
  }

}

// isFile()
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
    // invoke: bool isFile()
    var ret = C.C_ZNK9QFileInfo6isFileEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "isFile", args)
  }

}

// makeAbsolute()
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
    // invoke: bool makeAbsolute()
    var ret = C.C_ZN9QFileInfo12makeAbsoluteEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "makeAbsolute", args)
  }

}

// dir()
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
    // invoke: QDir dir()
    var ret = C.C_ZNK9QFileInfo3dirEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileInfo", "dir", args)
  }

}

// <= body block end

