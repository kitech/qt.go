package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
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
  // proto:  qint64 QStorageInfo::bytesFree();
extern void _ZNK12QStorageInfo9bytesFreeEv(void* qthis);
  // proto:  void QStorageInfo::QStorageInfo(const QStorageInfo & other);
extern void* dector_ZN12QStorageInfoC1ERKS_(void* arg0);
extern void _ZN12QStorageInfoC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QStorageInfo::isRoot();
extern void demth_ZNK12QStorageInfo6isRootEv(void* qthis);
  // proto:  bool QStorageInfo::isReadOnly();
extern void _ZNK12QStorageInfo10isReadOnlyEv(void* qthis);
  // proto:  QByteArray QStorageInfo::fileSystemType();
extern void _ZNK12QStorageInfo14fileSystemTypeEv(void* qthis);
  // proto:  void QStorageInfo::setPath(const QString & path);
extern void _ZN12QStorageInfo7setPathERK7QString(void* qthis, void* arg0);
  // proto: static QList<QStorageInfo> QStorageInfo::mountedVolumes();
extern void _ZN12QStorageInfo14mountedVolumesEv();
  // proto:  QString QStorageInfo::name();
extern void _ZNK12QStorageInfo4nameEv(void* qthis);
  // proto:  void QStorageInfo::refresh();
extern void _ZN12QStorageInfo7refreshEv(void* qthis);
  // proto:  bool QStorageInfo::isValid();
extern void _ZNK12QStorageInfo7isValidEv(void* qthis);
  // proto:  bool QStorageInfo::isReady();
extern void _ZNK12QStorageInfo7isReadyEv(void* qthis);
  // proto:  qint64 QStorageInfo::bytesTotal();
extern void _ZNK12QStorageInfo10bytesTotalEv(void* qthis);
  // proto:  QString QStorageInfo::rootPath();
extern void _ZNK12QStorageInfo8rootPathEv(void* qthis);
  // proto:  void QStorageInfo::~QStorageInfo();
extern void _ZN12QStorageInfoD0Ev(void* qthis);
  // proto:  qint64 QStorageInfo::bytesAvailable();
extern void _ZNK12QStorageInfo14bytesAvailableEv(void* qthis);
  // proto:  void QStorageInfo::QStorageInfo();
extern void* dector_ZN12QStorageInfoC1Ev();
extern void _ZN12QStorageInfoC1Ev(void* qthis);
  // proto:  void QStorageInfo::QStorageInfo(const QDir & dir);
extern void* dector_ZN12QStorageInfoC1ERK4QDir(void* arg0);
extern void _ZN12QStorageInfoC1ERK4QDir(void* qthis, void* arg0);
  // proto: static QStorageInfo QStorageInfo::root();
extern void _ZN12QStorageInfo4rootEv();
  // proto:  void QStorageInfo::QStorageInfo(const QString & path);
extern void* dector_ZN12QStorageInfoC1ERK7QString(void* arg0);
extern void _ZN12QStorageInfoC1ERK7QString(void* qthis, void* arg0);
  // proto:  QByteArray QStorageInfo::device();
extern void _ZNK12QStorageInfo6deviceEv(void* qthis);
  // proto:  QString QStorageInfo::displayName();
extern void _ZNK12QStorageInfo11displayNameEv(void* qthis);
  // proto:  void QStorageInfo::swap(QStorageInfo & other);
extern void demth_ZN12QStorageInfo4swapERS_(void* qthis, void* arg0);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  qint64 QStorageInfo::bytesFree();
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
    C._ZNK12QStorageInfo9bytesFreeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "bytesFree", args)
  }

}

  // proto:  void QStorageInfo::QStorageInfo(const QStorageInfo & other);
func NewQStorageInfo(args ...interface{}) QStorageInfo {
  return QStorageInfo{}
}

  // proto:  bool QStorageInfo::isRoot();
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
    C.demth_ZNK12QStorageInfo6isRootEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "isRoot", args)
  }

}

  // proto:  bool QStorageInfo::isReadOnly();
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
    C._ZNK12QStorageInfo10isReadOnlyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "isReadOnly", args)
  }

}

  // proto:  QByteArray QStorageInfo::fileSystemType();
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
    C._ZNK12QStorageInfo14fileSystemTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "fileSystemType", args)
  }

}

  // proto:  void QStorageInfo::setPath(const QString & path);
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
    C._ZN12QStorageInfo7setPathERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStorageInfo", "setPath", args)
  }

}

  // proto: static QList<QStorageInfo> QStorageInfo::mountedVolumes();
func (this *QStorageInfo) mountedVolumes_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStorageInfo", "mountedVolumes", args)
  }

}

  // proto:  QString QStorageInfo::name();
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
    C._ZNK12QStorageInfo4nameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "name", args)
  }

}

  // proto:  void QStorageInfo::refresh();
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
    C._ZN12QStorageInfo7refreshEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "refresh", args)
  }

}

  // proto:  bool QStorageInfo::isValid();
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
    C._ZNK12QStorageInfo7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "isValid", args)
  }

}

  // proto:  bool QStorageInfo::isReady();
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
    C._ZNK12QStorageInfo7isReadyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "isReady", args)
  }

}

  // proto:  qint64 QStorageInfo::bytesTotal();
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
    C._ZNK12QStorageInfo10bytesTotalEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "bytesTotal", args)
  }

}

  // proto:  QString QStorageInfo::rootPath();
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
    C._ZNK12QStorageInfo8rootPathEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "rootPath", args)
  }

}

  // proto:  void QStorageInfo::~QStorageInfo();
func (this *QStorageInfo) FreeQStorageInfo(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStorageInfo", "~QStorageInfo", args)
  }

}

  // proto:  qint64 QStorageInfo::bytesAvailable();
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
    C._ZNK12QStorageInfo14bytesAvailableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "bytesAvailable", args)
  }

}

  // proto: static QStorageInfo QStorageInfo::root();
func (this *QStorageInfo) root_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStorageInfo", "root", args)
  }

}

  // proto:  QByteArray QStorageInfo::device();
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
    C._ZNK12QStorageInfo6deviceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "device", args)
  }

}

  // proto:  QString QStorageInfo::displayName();
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
    C._ZNK12QStorageInfo11displayNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "displayName", args)
  }

}

  // proto:  void QStorageInfo::swap(QStorageInfo & other);
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
    C.demth_ZN12QStorageInfo4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStorageInfo", "swap", args)
  }

}

// <= body block end

