package qtcore
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
extern bool C_ZNK12QStorageInfo7isReadyEv(void* qthis); // 4
  // proto: static QList<QStorageInfo> QStorageInfo::mountedVolumes();
extern void C_ZN12QStorageInfo14mountedVolumesEv(); // 4
  // proto:  QString QStorageInfo::name();
extern void* C_ZNK12QStorageInfo4nameEv(void* qthis); // 4
  // proto:  bool QStorageInfo::isRoot();
extern bool C_ZNK12QStorageInfo6isRootEv(void* qthis); // 2
  // proto:  bool QStorageInfo::isReadOnly();
extern bool C_ZNK12QStorageInfo10isReadOnlyEv(void* qthis); // 4
  // proto:  bool QStorageInfo::isValid();
extern bool C_ZNK12QStorageInfo7isValidEv(void* qthis); // 4
  // proto:  QByteArray QStorageInfo::fileSystemType();
extern void* C_ZNK12QStorageInfo14fileSystemTypeEv(void* qthis); // 4
  // proto:  QString QStorageInfo::rootPath();
extern void* C_ZNK12QStorageInfo8rootPathEv(void* qthis); // 4
  // proto:  void QStorageInfo::refresh();
extern void C_ZN12QStorageInfo7refreshEv(void* qthis); // 4
  // proto:  QByteArray QStorageInfo::device();
extern void* C_ZNK12QStorageInfo6deviceEv(void* qthis); // 4
  // proto:  qint64 QStorageInfo::bytesFree();
extern int64_t C_ZNK12QStorageInfo9bytesFreeEv(void* qthis); // 4
  // proto:  qint64 QStorageInfo::bytesTotal();
extern int64_t C_ZNK12QStorageInfo10bytesTotalEv(void* qthis); // 4
  // proto:  qint64 QStorageInfo::bytesAvailable();
extern int64_t C_ZNK12QStorageInfo14bytesAvailableEv(void* qthis); // 4
  // proto:  void QStorageInfo::swap(QStorageInfo & other);
extern void C_ZN12QStorageInfo4swapERS_(void* qthis, void* arg0); // 2
  // proto:  int QStorageInfo::blockSize();
extern int32_t C_ZNK12QStorageInfo9blockSizeEv(void* qthis); // 4
  // proto:  void QStorageInfo::QStorageInfo(const QDir & dir);
extern void* C_ZN12QStorageInfoC2ERK4QDir(void* arg0); // 3
  // proto:  void QStorageInfo::QStorageInfo(const QStorageInfo & other);
extern void* C_ZN12QStorageInfoC2ERKS_(void* arg0); // 3
  // proto:  void QStorageInfo::QStorageInfo();
extern void* C_ZN12QStorageInfoC2Ev(); // 3
  // proto:  void QStorageInfo::QStorageInfo(const QString & path);
extern void* C_ZN12QStorageInfoC2ERK7QString(void* arg0); // 3
  // proto:  QString QStorageInfo::displayName();
extern void* C_ZNK12QStorageInfo11displayNameEv(void* qthis); // 4
  // proto:  void QStorageInfo::setPath(const QString & path);
extern void C_ZN12QStorageInfo7setPathERK7QString(void* qthis, void* arg0); // 4
  // proto: static QStorageInfo QStorageInfo::root();
extern void* C_ZN12QStorageInfo4rootEv(); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// isReady()
func (this *QStorageInfo) Isready(args ...interface{}) (ret interface{}) {
  // isReady()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo7isReadyEv
    // invoke: bool isReady()
    var ret0 = C.C_ZNK12QStorageInfo7isReadyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStorageInfo", "isReady", args)
  }

  return
}

// mountedVolumes()
func (this *QStorageInfo) Mountedvolumes_S(args ...interface{}) () {
  // mountedVolumes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

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

  return
}

// name()
func (this *QStorageInfo) Name(args ...interface{}) (ret interface{}) {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo4nameEv
    // invoke: QString name()
    var ret0 = C.C_ZNK12QStorageInfo4nameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStorageInfo", "name", args)
  }

  return
}

// isRoot()
func (this *QStorageInfo) Isroot(args ...interface{}) (ret interface{}) {
  // isRoot()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo6isRootEv
    // invoke: bool isRoot()
    var ret0 = C.C_ZNK12QStorageInfo6isRootEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStorageInfo", "isRoot", args)
  }

  return
}

// isReadOnly()
func (this *QStorageInfo) Isreadonly(args ...interface{}) (ret interface{}) {
  // isReadOnly()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo10isReadOnlyEv
    // invoke: bool isReadOnly()
    var ret0 = C.C_ZNK12QStorageInfo10isReadOnlyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStorageInfo", "isReadOnly", args)
  }

  return
}

// isValid()
func (this *QStorageInfo) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK12QStorageInfo7isValidEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStorageInfo", "isValid", args)
  }

  return
}

// fileSystemType()
func (this *QStorageInfo) Filesystemtype(args ...interface{}) (ret interface{}) {
  // fileSystemType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo14fileSystemTypeEv
    // invoke: QByteArray fileSystemType()
    var ret0 = C.C_ZNK12QStorageInfo14fileSystemTypeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStorageInfo", "fileSystemType", args)
  }

  return
}

// rootPath()
func (this *QStorageInfo) Rootpath(args ...interface{}) (ret interface{}) {
  // rootPath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo8rootPathEv
    // invoke: QString rootPath()
    var ret0 = C.C_ZNK12QStorageInfo8rootPathEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStorageInfo", "rootPath", args)
  }

  return
}

// refresh()
func (this *QStorageInfo) Refresh(args ...interface{}) () {
  // refresh()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QStorageInfo7refreshEv
    // invoke: void refresh()
    C.C_ZN12QStorageInfo7refreshEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "refresh", args)
  }

  return
}

// device()
func (this *QStorageInfo) Device(args ...interface{}) (ret interface{}) {
  // device()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo6deviceEv
    // invoke: QByteArray device()
    var ret0 = C.C_ZNK12QStorageInfo6deviceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStorageInfo", "device", args)
  }

  return
}

// bytesFree()
func (this *QStorageInfo) Bytesfree(args ...interface{}) (ret interface{}) {
  // bytesFree()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo9bytesFreeEv
    // invoke: qint64 bytesFree()
    var ret0 = C.C_ZNK12QStorageInfo9bytesFreeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStorageInfo", "bytesFree", args)
  }

  return
}

// bytesTotal()
func (this *QStorageInfo) Bytestotal(args ...interface{}) (ret interface{}) {
  // bytesTotal()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo10bytesTotalEv
    // invoke: qint64 bytesTotal()
    var ret0 = C.C_ZNK12QStorageInfo10bytesTotalEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStorageInfo", "bytesTotal", args)
  }

  return
}

// bytesAvailable()
func (this *QStorageInfo) Bytesavailable(args ...interface{}) (ret interface{}) {
  // bytesAvailable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo14bytesAvailableEv
    // invoke: qint64 bytesAvailable()
    var ret0 = C.C_ZNK12QStorageInfo14bytesAvailableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStorageInfo", "bytesAvailable", args)
  }

  return
}

// swap(class QStorageInfo &)
func (this *QStorageInfo) Swap(args ...interface{}) () {
  // swap(class QStorageInfo &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStorageInfo{}) // "QStorageInfo &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QStorageInfo4swapERS_
    // invoke: void swap(class QStorageInfo &)
    var arg0 = args[0].(*QStorageInfo).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QStorageInfo4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStorageInfo", "swap", args)
  }

  return
}

// blockSize()
func (this *QStorageInfo) Blocksize(args ...interface{}) (ret interface{}) {
  // blockSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo9blockSizeEv
    // invoke: int blockSize()
    var ret0 = C.C_ZNK12QStorageInfo9blockSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStorageInfo", "blockSize", args)
  }

  return
}

// QStorageInfo(const class QDir &)
func NewQStorageInfo(args ...interface{}) *QStorageInfo {
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
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QStorageInfoC1ERK4QDir
    // invoke: void QStorageInfo(const class QDir &)
    var arg0 = args[0].(*QDir).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QStorageInfoC2ERK4QDir(arg0)
    return &QStorageInfo{Qclsinst:qthis}
  case 1:
    // invoke: _ZN12QStorageInfoC1ERKS_
    // invoke: void QStorageInfo(const class QStorageInfo &)
    var arg0 = args[0].(*QStorageInfo).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QStorageInfoC2ERKS_(arg0)
    return &QStorageInfo{Qclsinst:qthis}
  case 2:
    // invoke: _ZN12QStorageInfoC1Ev
    // invoke: void QStorageInfo()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QStorageInfoC2Ev()
    return &QStorageInfo{Qclsinst:qthis}
  case 3:
    // invoke: _ZN12QStorageInfoC1ERK7QString
    // invoke: void QStorageInfo(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QStorageInfoC2ERK7QString(arg0)
    return &QStorageInfo{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStorageInfo", "QStorageInfo", args)
  }

  return nil // QStorageInfo{Qclsinst:qthis}
}

// displayName()
func (this *QStorageInfo) Displayname(args ...interface{}) (ret interface{}) {
  // displayName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo11displayNameEv
    // invoke: QString displayName()
    var ret0 = C.C_ZNK12QStorageInfo11displayNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStorageInfo", "displayName", args)
  }

  return
}

// setPath(const class QString &)
func (this *QStorageInfo) Setpath(args ...interface{}) () {
  // setPath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QStorageInfo7setPathERK7QString
    // invoke: void setPath(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QStorageInfo7setPathERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStorageInfo", "setPath", args)
  }

  return
}

// root()
func (this *QStorageInfo) Root_S(args ...interface{}) (ret interface{}) {
  // root()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QStorageInfo4rootEv
    // invoke: QStorageInfo root()
    var ret0 = C.C_ZN12QStorageInfo4rootEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStorageInfo{}) // "QStorageInfo"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStorageInfo", "root", args)
  }

  return
}

// ~QStorageInfo()
func (this *QStorageInfo) Freeqstorageinfo(args ...interface{}) () {
  // ~QStorageInfo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QStorageInfoD0Ev
    // invoke: void ~QStorageInfo()
    C.C_ZN12QStorageInfoD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStorageInfo", "~QStorageInfo", args)
  }

  return
}

// <= body block end

