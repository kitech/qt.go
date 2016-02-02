package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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
extern void* C_ZNK9QFileInfo11absoluteDirEv(void* qthis); // 4
  // proto:  QString QFileInfo::suffix();
extern void* C_ZNK9QFileInfo6suffixEv(void* qthis); // 4
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
extern void* C_ZNK9QFileInfo13symLinkTargetEv(void* qthis); // 2
  // proto:  bool QFileInfo::isRelative();
extern bool C_ZNK9QFileInfo10isRelativeEv(void* qthis); // 4
  // proto:  QString QFileInfo::completeBaseName();
extern void* C_ZNK9QFileInfo16completeBaseNameEv(void* qthis); // 4
  // proto:  QString QFileInfo::canonicalFilePath();
extern void* C_ZNK9QFileInfo17canonicalFilePathEv(void* qthis); // 4
  // proto:  bool QFileInfo::isAbsolute();
extern bool C_ZNK9QFileInfo10isAbsoluteEv(void* qthis); // 2
  // proto:  QString QFileInfo::owner();
extern void* C_ZNK9QFileInfo5ownerEv(void* qthis); // 4
  // proto:  bool QFileInfo::isExecutable();
extern bool C_ZNK9QFileInfo12isExecutableEv(void* qthis); // 4
  // proto:  qint64 QFileInfo::size();
extern int64_t C_ZNK9QFileInfo4sizeEv(void* qthis); // 4
  // proto:  QString QFileInfo::bundleName();
extern void* C_ZNK9QFileInfo10bundleNameEv(void* qthis); // 4
  // proto:  QString QFileInfo::group();
extern void* C_ZNK9QFileInfo5groupEv(void* qthis); // 4
  // proto: static bool QFileInfo::exists(const QString & file);
extern bool C_ZN9QFileInfo6existsERK7QString(void* arg0); // 4
  // proto:  bool QFileInfo::exists();
extern bool C_ZNK9QFileInfo6existsEv(void* qthis); // 4
  // proto:  bool QFileInfo::isWritable();
extern bool C_ZNK9QFileInfo10isWritableEv(void* qthis); // 4
  // proto:  QString QFileInfo::filePath();
extern void* C_ZNK9QFileInfo8filePathEv(void* qthis); // 4
  // proto:  QString QFileInfo::absolutePath();
extern void* C_ZNK9QFileInfo12absolutePathEv(void* qthis); // 4
  // proto:  void QFileInfo::swap(QFileInfo & other);
extern void C_ZN9QFileInfo4swapERS_(void* qthis, void* arg0); // 2
  // proto:  QString QFileInfo::canonicalPath();
extern void* C_ZNK9QFileInfo13canonicalPathEv(void* qthis); // 4
  // proto:  bool QFileInfo::isBundle();
extern bool C_ZNK9QFileInfo8isBundleEv(void* qthis); // 4
  // proto:  bool QFileInfo::isHidden();
extern bool C_ZNK9QFileInfo8isHiddenEv(void* qthis); // 4
  // proto:  bool QFileInfo::isDir();
extern bool C_ZNK9QFileInfo5isDirEv(void* qthis); // 4
  // proto:  QDateTime QFileInfo::lastRead();
extern void* C_ZNK9QFileInfo8lastReadEv(void* qthis); // 4
  // proto:  bool QFileInfo::isRoot();
extern bool C_ZNK9QFileInfo6isRootEv(void* qthis); // 4
  // proto:  QString QFileInfo::absoluteFilePath();
extern void* C_ZNK9QFileInfo16absoluteFilePathEv(void* qthis); // 4
  // proto:  QString QFileInfo::fileName();
extern void* C_ZNK9QFileInfo8fileNameEv(void* qthis); // 4
  // proto:  void QFileInfo::setCaching(bool on);
extern void C_ZN9QFileInfo10setCachingEb(void* qthis, bool arg0); // 4
  // proto:  QString QFileInfo::completeSuffix();
extern void* C_ZNK9QFileInfo14completeSuffixEv(void* qthis); // 4
  // proto:  QString QFileInfo::path();
extern void* C_ZNK9QFileInfo4pathEv(void* qthis); // 4
  // proto:  void QFileInfo::~QFileInfo();
extern void C_ZN9QFileInfoD2Ev(void* qthis); // 4
  // proto:  uint QFileInfo::groupId();
extern int32_t C_ZNK9QFileInfo7groupIdEv(void* qthis); // 4
  // proto:  QFile::Permissions QFileInfo::permissions();
extern void C_ZNK9QFileInfo11permissionsEv(void* qthis); // 4
  // proto:  bool QFileInfo::isNativePath();
extern bool C_ZNK9QFileInfo12isNativePathEv(void* qthis); // 4
  // proto:  QDateTime QFileInfo::created();
extern void* C_ZNK9QFileInfo7createdEv(void* qthis); // 4
  // proto:  bool QFileInfo::isSymLink();
extern bool C_ZNK9QFileInfo9isSymLinkEv(void* qthis); // 4
  // proto:  QDateTime QFileInfo::lastModified();
extern void* C_ZNK9QFileInfo12lastModifiedEv(void* qthis); // 4
  // proto:  QString QFileInfo::baseName();
extern void* C_ZNK9QFileInfo8baseNameEv(void* qthis); // 4
  // proto:  void QFileInfo::refresh();
extern void C_ZN9QFileInfo7refreshEv(void* qthis); // 4
  // proto:  QString QFileInfo::readLink();
extern void* C_ZNK9QFileInfo8readLinkEv(void* qthis); // 4
  // proto:  void QFileInfo::setFile(const QDir & dir, const QString & file);
extern void C_ZN9QFileInfo7setFileERK4QDirRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QFileInfo::setFile(const QString & file);
extern void C_ZN9QFileInfo7setFileERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QFileInfo::setFile(const QFile & file);
extern void C_ZN9QFileInfo7setFileERK5QFile(void* qthis, void* arg0); // 4
  // proto:  bool QFileInfo::isReadable();
extern bool C_ZNK9QFileInfo10isReadableEv(void* qthis); // 4
  // proto:  bool QFileInfo::caching();
extern bool C_ZNK9QFileInfo7cachingEv(void* qthis); // 4
  // proto:  uint QFileInfo::ownerId();
extern int32_t C_ZNK9QFileInfo7ownerIdEv(void* qthis); // 4
  // proto:  bool QFileInfo::isFile();
extern bool C_ZNK9QFileInfo6isFileEv(void* qthis); // 4
  // proto:  bool QFileInfo::makeAbsolute();
extern bool C_ZN9QFileInfo12makeAbsoluteEv(void* qthis); // 4
  // proto:  QDir QFileInfo::dir();
extern void* C_ZNK9QFileInfo3dirEv(void* qthis); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// absoluteDir()
func (this *QFileInfo) Absolutedir(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo11absoluteDirEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDir{}) // "QDir"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "absoluteDir", args)
  }

  return
}

// suffix()
func (this *QFileInfo) Suffix(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo6suffixEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "suffix", args)
  }

  return
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
    var arg0 = args[0].(*QFileInfo).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QFileInfoC2ERKS_(arg0)
    return &QFileInfo{Qclsinst:qthis}
  case 1:
    // invoke: _ZN9QFileInfoC1Ev
    // invoke: void QFileInfo()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QFileInfoC2Ev()
    return &QFileInfo{Qclsinst:qthis}
  case 2:
    // invoke: _ZN9QFileInfoC1ERK7QString
    // invoke: void QFileInfo(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QFileInfoC2ERK7QString(arg0)
    return &QFileInfo{Qclsinst:qthis}
  case 3:
    // invoke: _ZN9QFileInfoC1ERK5QFile
    // invoke: void QFileInfo(const class QFile &)
    var arg0 = args[0].(*QFile).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QFileInfoC2ERK5QFile(arg0)
    return &QFileInfo{Qclsinst:qthis}
  case 4:
    // invoke: _ZN9QFileInfoC1ERK4QDirRK7QString
    // invoke: void QFileInfo(const class QDir &, const class QString &)
    var arg0 = args[0].(*QDir).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QFileInfoC2ERK4QDirRK7QString(arg0, arg1)
    return &QFileInfo{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QFileInfo", "QFileInfo", args)
  }

  return nil // QFileInfo{Qclsinst:qthis}
}

// symLinkTarget()
func (this *QFileInfo) Symlinktarget(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo13symLinkTargetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "symLinkTarget", args)
  }

  return
}

// isRelative()
func (this *QFileInfo) Isrelative(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo10isRelativeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "isRelative", args)
  }

  return
}

// completeBaseName()
func (this *QFileInfo) Completebasename(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo16completeBaseNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "completeBaseName", args)
  }

  return
}

// canonicalFilePath()
func (this *QFileInfo) Canonicalfilepath(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo17canonicalFilePathEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "canonicalFilePath", args)
  }

  return
}

// isAbsolute()
func (this *QFileInfo) Isabsolute(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo10isAbsoluteEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "isAbsolute", args)
  }

  return
}

// owner()
func (this *QFileInfo) Owner(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo5ownerEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "owner", args)
  }

  return
}

// isExecutable()
func (this *QFileInfo) Isexecutable(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo12isExecutableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "isExecutable", args)
  }

  return
}

// size()
func (this *QFileInfo) Size(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "size", args)
  }

  return
}

// bundleName()
func (this *QFileInfo) Bundlename(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo10bundleNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "bundleName", args)
  }

  return
}

// group()
func (this *QFileInfo) Group(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo5groupEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "group", args)
  }

  return
}

// exists(const class QString &)
func (this *QFileInfo) Exists_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QFileInfo6existsERK7QString(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK9QFileInfo6existsEv
    // invoke: bool exists()
    var ret0 = C.C_ZNK9QFileInfo6existsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "exists", args)
  }

  return
}

// isWritable()
func (this *QFileInfo) Iswritable(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo10isWritableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "isWritable", args)
  }

  return
}

// filePath()
func (this *QFileInfo) Filepath(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo8filePathEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "filePath", args)
  }

  return
}

// absolutePath()
func (this *QFileInfo) Absolutepath(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo12absolutePathEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "absolutePath", args)
  }

  return
}

// swap(class QFileInfo &)
func (this *QFileInfo) Swap(args ...interface{}) () {
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
    var arg0 = args[0].(*QFileInfo).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QFileInfo4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileInfo", "swap", args)
  }

  return
}

// canonicalPath()
func (this *QFileInfo) Canonicalpath(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo13canonicalPathEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "canonicalPath", args)
  }

  return
}

// isBundle()
func (this *QFileInfo) Isbundle(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo8isBundleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "isBundle", args)
  }

  return
}

// isHidden()
func (this *QFileInfo) Ishidden(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo8isHiddenEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "isHidden", args)
  }

  return
}

// isDir()
func (this *QFileInfo) Isdir(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo5isDirEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "isDir", args)
  }

  return
}

// lastRead()
func (this *QFileInfo) Lastread(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo8lastReadEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDateTime{}) // "QDateTime"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "lastRead", args)
  }

  return
}

// isRoot()
func (this *QFileInfo) Isroot(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo6isRootEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "isRoot", args)
  }

  return
}

// absoluteFilePath()
func (this *QFileInfo) Absolutefilepath(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo16absoluteFilePathEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "absoluteFilePath", args)
  }

  return
}

// fileName()
func (this *QFileInfo) Filename(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo8fileNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "fileName", args)
  }

  return
}

// setCaching(_Bool)
func (this *QFileInfo) Setcaching(args ...interface{}) () {
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
    C.C_ZN9QFileInfo10setCachingEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileInfo", "setCaching", args)
  }

  return
}

// completeSuffix()
func (this *QFileInfo) Completesuffix(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo14completeSuffixEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "completeSuffix", args)
  }

  return
}

// path()
func (this *QFileInfo) Path(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo4pathEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "path", args)
  }

  return
}

// ~QFileInfo()
func (this *QFileInfo) Freeqfileinfo(args ...interface{}) () {
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
    C.C_ZN9QFileInfoD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileInfo", "~QFileInfo", args)
  }

  return
}

// groupId()
func (this *QFileInfo) Groupid(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo7groupIdEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "uint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "groupId", args)
  }

  return
}

// permissions()
func (this *QFileInfo) Permissions(args ...interface{}) () {
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
    C.C_ZNK9QFileInfo11permissionsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileInfo", "permissions", args)
  }

  return
}

// isNativePath()
func (this *QFileInfo) Isnativepath(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo12isNativePathEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "isNativePath", args)
  }

  return
}

// created()
func (this *QFileInfo) Created(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo7createdEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDateTime{}) // "QDateTime"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "created", args)
  }

  return
}

// isSymLink()
func (this *QFileInfo) Issymlink(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo9isSymLinkEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "isSymLink", args)
  }

  return
}

// lastModified()
func (this *QFileInfo) Lastmodified(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo12lastModifiedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDateTime{}) // "QDateTime"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "lastModified", args)
  }

  return
}

// baseName()
func (this *QFileInfo) Basename(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo8baseNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "baseName", args)
  }

  return
}

// refresh()
func (this *QFileInfo) Refresh(args ...interface{}) () {
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
    C.C_ZN9QFileInfo7refreshEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileInfo", "refresh", args)
  }

  return
}

// readLink()
func (this *QFileInfo) Readlink(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo8readLinkEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "readLink", args)
  }

  return
}

// setFile(const class QDir &, const class QString &)
func (this *QFileInfo) Setfile(args ...interface{}) () {
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
    var arg0 = args[0].(*QDir).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN9QFileInfo7setFileERK4QDirRK7QString(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN9QFileInfo7setFileERK7QString
    // invoke: void setFile(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QFileInfo7setFileERK7QString(this.Qclsinst, arg0)
  case 2:
    // invoke: _ZN9QFileInfo7setFileERK5QFile
    // invoke: void setFile(const class QFile &)
    var arg0 = args[0].(*QFile).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QFileInfo7setFileERK5QFile(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileInfo", "setFile", args)
  }

  return
}

// isReadable()
func (this *QFileInfo) Isreadable(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo10isReadableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "isReadable", args)
  }

  return
}

// caching()
func (this *QFileInfo) Caching(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo7cachingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "caching", args)
  }

  return
}

// ownerId()
func (this *QFileInfo) Ownerid(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo7ownerIdEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "uint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "ownerId", args)
  }

  return
}

// isFile()
func (this *QFileInfo) Isfile(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo6isFileEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "isFile", args)
  }

  return
}

// makeAbsolute()
func (this *QFileInfo) Makeabsolute(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN9QFileInfo12makeAbsoluteEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "makeAbsolute", args)
  }

  return
}

// dir()
func (this *QFileInfo) Dir(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QFileInfo3dirEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDir{}) // "QDir"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileInfo", "dir", args)
  }

  return
}

// <= body block end

