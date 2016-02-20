package qtcore
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
// src-file: /QtCore/qfile.h
// dst-file: /src/core/qfile.go
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
  // proto: static bool QFile::rename(const QString & oldName, const QString & newName);
extern bool C_ZN5QFile6renameERK7QStringS2_(void* arg0, void* arg1); // 4
  // proto:  bool QFile::rename(const QString & newName);
extern bool C_ZN5QFile6renameERK7QString(void* qthis, void* arg0); // 4
  // proto: static bool QFile::exists(const QString & fileName);
extern bool C_ZN5QFile6existsERK7QString(void* arg0); // 4
  // proto:  bool QFile::exists();
extern bool C_ZNK5QFile6existsEv(void* qthis); // 4
  // proto:  QString QFile::symLinkTarget();
extern void* C_ZNK5QFile13symLinkTargetEv(void* qthis); // 2
  // proto: static QString QFile::symLinkTarget(const QString & fileName);
extern void* C_ZN5QFile13symLinkTargetERK7QString(void* arg0); // 2
  // proto:  qint64 QFile::size();
extern int64_t C_ZNK5QFile4sizeEv(void* qthis); // 4
  // proto: static QByteArray QFile::encodeName(const QString & fileName);
extern void* C_ZN5QFile10encodeNameERK7QString(void* arg0); // 2
  // proto: static QString QFile::decodeName(const QByteArray & localFileName);
extern void* C_ZN5QFile10decodeNameERK10QByteArray(void* arg0); // 2
  // proto: static QString QFile::decodeName(const char * localFileName);
extern void* C_ZN5QFile10decodeNameEPKc(void* arg0); // 2
  // proto:  void QFile::QFile(const QString & name, QObject * parent);
extern void* C_ZN5QFileC2ERK7QStringP7QObject(void* arg0, void* arg1); // 3
  // proto:  void QFile::QFile();
extern void* C_ZN5QFileC2Ev(); // 3
  // proto:  void QFile::QFile(QObject * parent);
extern void* C_ZN5QFileC2EP7QObject(void* arg0); // 3
  // proto:  void QFile::QFile(const QString & name);
extern void* C_ZN5QFileC2ERK7QString(void* arg0); // 3
  // proto:  QString QFile::fileName();
extern void* C_ZNK5QFile8fileNameEv(void* qthis); // 4
  // proto:  bool QFile::link(const QString & newName);
extern bool C_ZN5QFile4linkERK7QString(void* qthis, void* arg0); // 4
  // proto: static bool QFile::link(const QString & oldname, const QString & newName);
extern bool C_ZN5QFile4linkERK7QStringS2_(void* arg0, void* arg1); // 4
  // proto: static bool QFile::copy(const QString & fileName, const QString & newName);
extern bool C_ZN5QFile4copyERK7QStringS2_(void* arg0, void* arg1); // 4
  // proto:  bool QFile::copy(const QString & newName);
extern bool C_ZN5QFile4copyERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QFile::resize(qint64 sz);
extern bool C_ZN5QFile6resizeEx(void* qthis, int64_t arg0); // 4
  // proto: static bool QFile::resize(const QString & filename, qint64 sz);
extern bool C_ZN5QFile6resizeERK7QStringx(void* arg0, int64_t arg1); // 4
  // proto:  Permissions QFile::permissions();
extern void C_ZNK5QFile11permissionsEv(void* qthis); // 4
  // proto: static Permissions QFile::permissions(const QString & filename);
extern void C_ZN5QFile11permissionsERK7QString(void* arg0); // 4
  // proto:  const QMetaObject * QFile::metaObject();
extern void C_ZNK5QFile10metaObjectEv(void* qthis); // 4
  // proto:  void QFile::~QFile();
extern void C_ZN5QFileD2Ev(void* qthis); // 4
  // proto:  void QFile::setFileName(const QString & name);
extern void C_ZN5QFile11setFileNameERK7QString(void* qthis, void* arg0); // 4
  // proto: static bool QFile::remove(const QString & fileName);
extern bool C_ZN5QFile6removeERK7QString(void* arg0); // 4
  // proto:  bool QFile::remove();
extern bool C_ZN5QFile6removeEv(void* qthis); // 4
  // proto: static QString QFile::readLink(const QString & fileName);
extern void* C_ZN5QFile8readLinkERK7QString(void* arg0); // 4
  // proto:  QString QFile::readLink();
extern void* C_ZNK5QFile8readLinkEv(void* qthis); // 4
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

// class sizeof(QFile)=1
type QFile struct {
  /*qbase*/ QFileDevice;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// rename(const class QString &, const class QString &)
func (this *QFile) Rename_S(args ...interface{}) (ret interface{}) {
  // rename(const class QString &, const class QString &)
  // rename(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFile6renameERK7QStringS2_
    // invoke: bool rename(const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN5QFile6renameERK7QStringS2_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN5QFile6renameERK7QString
    // invoke: bool rename(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QFile6renameERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFile", "rename", args)
  }

  return
}

// exists(const class QString &)
func (this *QFile) Exists_S(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZN5QFile6existsERK7QString
    // invoke: bool exists(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QFile6existsERK7QString(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK5QFile6existsEv
    // invoke: bool exists()
    var ret0 = C.C_ZNK5QFile6existsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFile", "exists", args)
  }

  return
}

// symLinkTarget()
func (this *QFile) Symlinktarget(args ...interface{}) (ret interface{}) {
  // symLinkTarget()
  // symLinkTarget(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFile13symLinkTargetEv
    // invoke: QString symLinkTarget()
    var ret0 = C.C_ZNK5QFile13symLinkTargetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN5QFile13symLinkTargetERK7QString
    // invoke: QString symLinkTarget(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QFile13symLinkTargetERK7QString(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFile", "symLinkTarget", args)
  }

  return
}

// size()
func (this *QFile) Size(args ...interface{}) (ret interface{}) {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFile4sizeEv
    // invoke: qint64 size()
    var ret0 = C.C_ZNK5QFile4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFile", "size", args)
  }

  return
}

// encodeName(const class QString &)
func (this *QFile) Encodename_S(args ...interface{}) (ret interface{}) {
  // encodeName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFile10encodeNameERK7QString
    // invoke: QByteArray encodeName(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QFile10encodeNameERK7QString(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFile", "encodeName", args)
  }

  return
}

// decodeName(const class QByteArray &)
func (this *QFile) Decodename_S(args ...interface{}) (ret interface{}) {
  // decodeName(const class QByteArray &)
  // decodeName(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFile10decodeNameERK10QByteArray
    // invoke: QString decodeName(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QFile10decodeNameERK10QByteArray(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN5QFile10decodeNameEPKc
    // invoke: QString decodeName(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZN5QFile10decodeNameEPKc(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFile", "decodeName", args)
  }

  return
}

// QFile(const class QString &, class QObject *)
func NewQFile(args ...interface{}) *QFile {
  // QFile(const class QString &, class QObject *)
  // QFile()
  // QFile(class QObject *)
  // QFile(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFileC1ERK7QStringP7QObject
    // invoke: void QFile(const class QString &, class QObject *)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QObject).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QFileC2ERK7QStringP7QObject(arg0, arg1)
    return &QFile{Qclsinst:qthis}
  case 1:
    // invoke: _ZN5QFileC1Ev
    // invoke: void QFile()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QFileC2Ev()
    return &QFile{Qclsinst:qthis}
  case 2:
    // invoke: _ZN5QFileC1EP7QObject
    // invoke: void QFile(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QFileC2EP7QObject(arg0)
    return &QFile{Qclsinst:qthis}
  case 3:
    // invoke: _ZN5QFileC1ERK7QString
    // invoke: void QFile(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QFileC2ERK7QString(arg0)
    return &QFile{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QFile", "QFile", args)
  }

  return nil // QFile{Qclsinst:qthis}
}

// fileName()
func (this *QFile) Filename(args ...interface{}) (ret interface{}) {
  // fileName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFile8fileNameEv
    // invoke: QString fileName()
    var ret0 = C.C_ZNK5QFile8fileNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFile", "fileName", args)
  }

  return
}

// link(const class QString &)
func (this *QFile) Link(args ...interface{}) (ret interface{}) {
  // link(const class QString &)
  // link(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFile4linkERK7QString
    // invoke: bool link(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QFile4linkERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN5QFile4linkERK7QStringS2_
    // invoke: bool link(const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN5QFile4linkERK7QStringS2_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFile", "link", args)
  }

  return
}

// copy(const class QString &, const class QString &)
func (this *QFile) Copy_S(args ...interface{}) (ret interface{}) {
  // copy(const class QString &, const class QString &)
  // copy(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFile4copyERK7QStringS2_
    // invoke: bool copy(const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN5QFile4copyERK7QStringS2_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN5QFile4copyERK7QString
    // invoke: bool copy(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QFile4copyERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFile", "copy", args)
  }

  return
}

// resize(qint64)
func (this *QFile) Resize(args ...interface{}) (ret interface{}) {
  // resize(qint64)
  // resize(const class QString &, qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFile6resizeEx
    // invoke: bool resize(qint64)
    var arg0 = C.int64_t(qtrt.PrimConv(args[0], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QFile6resizeEx(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN5QFile6resizeERK7QStringx
    // invoke: bool resize(const class QString &, qint64)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int64_t(qtrt.PrimConv(args[1], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN5QFile6resizeERK7QStringx(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFile", "resize", args)
  }

  return
}

// permissions()
func (this *QFile) Permissions(args ...interface{}) () {
  // permissions()
  // permissions(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFile11permissionsEv
    // invoke: Permissions permissions()
    C.C_ZNK5QFile11permissionsEv(this.Qclsinst)
  case 1:
    // invoke: _ZN5QFile11permissionsERK7QString
    // invoke: Permissions permissions(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QFile11permissionsERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QFile", "permissions", args)
  }

  return
}

// metaObject()
func (this *QFile) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFile10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK5QFile10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFile", "metaObject", args)
  }

  return
}

// ~QFile()
func (this *QFile) Freeqfile(args ...interface{}) () {
  // ~QFile()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFileD0Ev
    // invoke: void ~QFile()
    C.C_ZN5QFileD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFile", "~QFile", args)
  }

  return
}

// setFileName(const class QString &)
func (this *QFile) Setfilename(args ...interface{}) () {
  // setFileName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFile11setFileNameERK7QString
    // invoke: void setFileName(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QFile11setFileNameERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFile", "setFileName", args)
  }

  return
}

// remove(const class QString &)
func (this *QFile) Remove_S(args ...interface{}) (ret interface{}) {
  // remove(const class QString &)
  // remove()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFile6removeERK7QString
    // invoke: bool remove(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QFile6removeERK7QString(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN5QFile6removeEv
    // invoke: bool remove()
    var ret0 = C.C_ZN5QFile6removeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFile", "remove", args)
  }

  return
}

// readLink(const class QString &)
func (this *QFile) Readlink_S(args ...interface{}) (ret interface{}) {
  // readLink(const class QString &)
  // readLink()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFile8readLinkERK7QString
    // invoke: QString readLink(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QFile8readLinkERK7QString(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK5QFile8readLinkEv
    // invoke: QString readLink()
    var ret0 = C.C_ZNK5QFile8readLinkEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFile", "readLink", args)
  }

  return
}

// <= body block end

