package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtCore/qstorageinfo.h
// dst-file: /src/core/qstorageinfo.go
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
  // proto:  bool QStorageInfo::isReady();
extern void C_ZNK12QStorageInfo7isReadyEv(void* qthis); // 4
  // proto: static QList<QStorageInfo> QStorageInfo::mountedVolumes();
extern void C_ZN12QStorageInfo14mountedVolumesEv(); // 4
  // proto:  QString QStorageInfo::name();
extern void C_ZNK12QStorageInfo4nameEv(void* qthis); // 4
  // proto:  bool QStorageInfo::isRoot();
extern void C_ZNK12QStorageInfo6isRootEv(void* qthis); // 2
  // proto:  bool QStorageInfo::isReadOnly();
extern void C_ZNK12QStorageInfo10isReadOnlyEv(void* qthis); // 4
  // proto:  bool QStorageInfo::isValid();
extern void C_ZNK12QStorageInfo7isValidEv(void* qthis); // 4
  // proto:  QByteArray QStorageInfo::fileSystemType();
extern void C_ZNK12QStorageInfo14fileSystemTypeEv(void* qthis); // 4
  // proto:  QString QStorageInfo::rootPath();
extern void C_ZNK12QStorageInfo8rootPathEv(void* qthis); // 4
  // proto:  void QStorageInfo::refresh();
extern void C_ZN12QStorageInfo7refreshEv(void* qthis); // 4
  // proto:  qint64 QStorageInfo::bytesAvailable();
extern void C_ZNK12QStorageInfo14bytesAvailableEv(void* qthis); // 4
  // proto:  qint64 QStorageInfo::bytesFree();
extern void C_ZNK12QStorageInfo9bytesFreeEv(void* qthis); // 4
  // proto:  qint64 QStorageInfo::bytesTotal();
extern void C_ZNK12QStorageInfo10bytesTotalEv(void* qthis); // 4
  // proto:  QByteArray QStorageInfo::device();
extern void C_ZNK12QStorageInfo6deviceEv(void* qthis); // 4
  // proto:  void QStorageInfo::swap(QStorageInfo & other);
extern void C_ZN12QStorageInfo4swapERS_(void* qthis, void* arg0); // 2
  // proto:  void QStorageInfo::QStorageInfo(const QDir & dir);
extern void C_ZN12QStorageInfoC2ERK4QDir(void* qthis, void* arg0); // 3
  // proto:  void QStorageInfo::QStorageInfo(const QStorageInfo & other);
extern void C_ZN12QStorageInfoC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  void QStorageInfo::QStorageInfo();
extern void C_ZN12QStorageInfoC2Ev(void* qthis); // 3
  // proto:  void QStorageInfo::QStorageInfo(const QString & path);
extern void C_ZN12QStorageInfoC2ERK7QString(void* qthis, void* arg0); // 3
  // proto:  QString QStorageInfo::displayName();
extern void C_ZNK12QStorageInfo11displayNameEv(void* qthis); // 4
  // proto:  void QStorageInfo::setPath(const QString & path);
extern void C_ZN12QStorageInfo7setPathERK7QString(void* qthis, void* arg0); // 4
  // proto: static QStorageInfo QStorageInfo::root();
extern void C_ZN12QStorageInfo4rootEv(); // 4
  // proto:  void QStorageInfo::~QStorageInfo();
extern void C_ZN12QStorageInfoD2Ev(void* qthis); // 4
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

// class sizeof(QStorageInfo)=1
type QStorageInfo struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// isReady()
func (this *QStorageInfo) isReady(args ...interface{}) () {
  // isReady()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo7isReadyEv
    // invoke: bool isReady()
    C.C_ZNK12QStorageInfo7isReadyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "isReady", args)
  }

}

// mountedVolumes()
func (this *QStorageInfo) mountedVolumes_s(args ...interface{}) () {
  // mountedVolumes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QStorageInfo14mountedVolumesEv
    // invoke: QList<QStorageInfo> mountedVolumes()
    C.C_ZN12QStorageInfo14mountedVolumesEv()
  default:
    qtrt.ErrorResolve("QStorageInfo", "mountedVolumes", args)
  }

}

// name()
func (this *QStorageInfo) name(args ...interface{}) () {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo4nameEv
    // invoke: QString name()
    C.C_ZNK12QStorageInfo4nameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "name", args)
  }

}

// isRoot()
func (this *QStorageInfo) isRoot(args ...interface{}) () {
  // isRoot()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo6isRootEv
    // invoke: bool isRoot()
    C.C_ZNK12QStorageInfo6isRootEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "isRoot", args)
  }

}

// isReadOnly()
func (this *QStorageInfo) isReadOnly(args ...interface{}) () {
  // isReadOnly()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo10isReadOnlyEv
    // invoke: bool isReadOnly()
    C.C_ZNK12QStorageInfo10isReadOnlyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "isReadOnly", args)
  }

}

// isValid()
func (this *QStorageInfo) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo7isValidEv
    // invoke: bool isValid()
    C.C_ZNK12QStorageInfo7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "isValid", args)
  }

}

// fileSystemType()
func (this *QStorageInfo) fileSystemType(args ...interface{}) () {
  // fileSystemType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo14fileSystemTypeEv
    // invoke: QByteArray fileSystemType()
    C.C_ZNK12QStorageInfo14fileSystemTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "fileSystemType", args)
  }

}

// rootPath()
func (this *QStorageInfo) rootPath(args ...interface{}) () {
  // rootPath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo8rootPathEv
    // invoke: QString rootPath()
    C.C_ZNK12QStorageInfo8rootPathEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "rootPath", args)
  }

}

// refresh()
func (this *QStorageInfo) refresh(args ...interface{}) () {
  // refresh()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QStorageInfo7refreshEv
    // invoke: void refresh()
    C.C_ZN12QStorageInfo7refreshEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "refresh", args)
  }

}

// bytesAvailable()
func (this *QStorageInfo) bytesAvailable(args ...interface{}) () {
  // bytesAvailable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo14bytesAvailableEv
    // invoke: qint64 bytesAvailable()
    C.C_ZNK12QStorageInfo14bytesAvailableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "bytesAvailable", args)
  }

}

// bytesFree()
func (this *QStorageInfo) bytesFree(args ...interface{}) () {
  // bytesFree()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo9bytesFreeEv
    // invoke: qint64 bytesFree()
    C.C_ZNK12QStorageInfo9bytesFreeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "bytesFree", args)
  }

}

// bytesTotal()
func (this *QStorageInfo) bytesTotal(args ...interface{}) () {
  // bytesTotal()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo10bytesTotalEv
    // invoke: qint64 bytesTotal()
    C.C_ZNK12QStorageInfo10bytesTotalEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "bytesTotal", args)
  }

}

// device()
func (this *QStorageInfo) device(args ...interface{}) () {
  // device()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo6deviceEv
    // invoke: QByteArray device()
    C.C_ZNK12QStorageInfo6deviceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "device", args)
  }

}

// swap(class QStorageInfo &)
func (this *QStorageInfo) swap(args ...interface{}) () {
  // swap(class QStorageInfo &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStorageInfo{}) // "QStorageInfo &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QStorageInfo4swapERS_
    // invoke: void swap(class QStorageInfo &)
    var arg0 = args[0].(QStorageInfo).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QStorageInfo4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStorageInfo", "swap", args)
  }

}

// QStorageInfo(const class QDir &)
func NewQStorageInfo(args ...interface{}) QStorageInfo {
  // QStorageInfo(const class QDir &)
  // QStorageInfo(const class QStorageInfo &)
  // QStorageInfo()
  // QStorageInfo(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDir{}) // "const QDir &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QStorageInfo{}) // "const QStorageInfo &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QStorageInfoC1ERK4QDir
    // invoke: void QStorageInfo(const class QDir &)
    var arg0 = args[0].(QDir).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN12QStorageInfoC2ERK4QDir(qthis, arg0)
  case 1:
    // invoke: _ZN12QStorageInfoC1ERKS_
    // invoke: void QStorageInfo(const class QStorageInfo &)
    var arg0 = args[0].(QStorageInfo).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN12QStorageInfoC2ERKS_(qthis, arg0)
  case 2:
    // invoke: _ZN12QStorageInfoC1Ev
    // invoke: void QStorageInfo()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN12QStorageInfoC2Ev(qthis)
  case 3:
    // invoke: _ZN12QStorageInfoC1ERK7QString
    // invoke: void QStorageInfo(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN12QStorageInfoC2ERK7QString(qthis, arg0)
  default:
    qtrt.ErrorResolve("QStorageInfo", "QStorageInfo", args)
  }

  return QStorageInfo{}
}

// displayName()
func (this *QStorageInfo) displayName(args ...interface{}) () {
  // displayName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo11displayNameEv
    // invoke: QString displayName()
    C.C_ZNK12QStorageInfo11displayNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "displayName", args)
  }

}

// setPath(const class QString &)
func (this *QStorageInfo) setPath(args ...interface{}) () {
  // setPath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QStorageInfo7setPathERK7QString
    // invoke: void setPath(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QStorageInfo7setPathERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStorageInfo", "setPath", args)
  }

}

// root()
func (this *QStorageInfo) root_s(args ...interface{}) () {
  // root()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QStorageInfo4rootEv
    // invoke: QStorageInfo root()
    C.C_ZN12QStorageInfo4rootEv()
  default:
    qtrt.ErrorResolve("QStorageInfo", "root", args)
  }

}

// ~QStorageInfo()
func (this *QStorageInfo) FreeQStorageInfo(args ...interface{}) () {
  // ~QStorageInfo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QStorageInfoD0Ev
    // invoke: void ~QStorageInfo()
    C.C_ZN12QStorageInfoD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "~QStorageInfo", args)
  }

}

// <= body block end

