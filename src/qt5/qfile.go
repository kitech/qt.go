package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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
extern void C_ZN5QFile6renameERK7QStringS2_(void* arg0, void* arg1); // 4
  // proto:  bool QFile::rename(const QString & newName);
extern void C_ZN5QFile6renameERK7QString(void* qthis, void* arg0); // 4
  // proto: static bool QFile::exists(const QString & fileName);
extern void C_ZN5QFile6existsERK7QString(void* arg0); // 4
  // proto:  bool QFile::exists();
extern void C_ZNK5QFile6existsEv(void* qthis); // 4
  // proto:  QString QFile::symLinkTarget();
extern void C_ZNK5QFile13symLinkTargetEv(void* qthis); // 2
  // proto: static QString QFile::symLinkTarget(const QString & fileName);
extern void C_ZN5QFile13symLinkTargetERK7QString(void* arg0); // 2
  // proto:  qint64 QFile::size();
extern void C_ZNK5QFile4sizeEv(void* qthis); // 4
  // proto: static QByteArray QFile::encodeName(const QString & fileName);
extern void C_ZN5QFile10encodeNameERK7QString(void* arg0); // 2
  // proto: static QString QFile::decodeName(const QByteArray & localFileName);
extern void C_ZN5QFile10decodeNameERK10QByteArray(void* arg0); // 2
  // proto: static QString QFile::decodeName(const char * localFileName);
extern void C_ZN5QFile10decodeNameEPKc(unsigned char* arg0); // 2
  // proto:  void QFile::QFile(const QString & name, QObject * parent);
extern void* C_ZN5QFileC2ERK7QStringP7QObject(void* arg0, void* arg1); // 3
  // proto:  void QFile::QFile();
extern void* C_ZN5QFileC2Ev(); // 3
  // proto:  void QFile::QFile(QObject * parent);
extern void* C_ZN5QFileC2EP7QObject(void* arg0); // 3
  // proto:  void QFile::QFile(const QString & name);
extern void* C_ZN5QFileC2ERK7QString(void* arg0); // 3
  // proto:  QString QFile::fileName();
extern void C_ZNK5QFile8fileNameEv(void* qthis); // 4
  // proto:  bool QFile::link(const QString & newName);
extern void C_ZN5QFile4linkERK7QString(void* qthis, void* arg0); // 4
  // proto: static bool QFile::link(const QString & oldname, const QString & newName);
extern void C_ZN5QFile4linkERK7QStringS2_(void* arg0, void* arg1); // 4
  // proto: static bool QFile::copy(const QString & fileName, const QString & newName);
extern void C_ZN5QFile4copyERK7QStringS2_(void* arg0, void* arg1); // 4
  // proto:  bool QFile::copy(const QString & newName);
extern void C_ZN5QFile4copyERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QFile::resize(qint64 sz);
extern void C_ZN5QFile6resizeEx(void* qthis, int64_t arg0); // 4
  // proto: static bool QFile::resize(const QString & filename, qint64 sz);
extern void C_ZN5QFile6resizeERK7QStringx(void* arg0, int64_t arg1); // 4
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
extern void C_ZN5QFile6removeERK7QString(void* arg0); // 4
  // proto:  bool QFile::remove();
extern void C_ZN5QFile6removeEv(void* qthis); // 4
  // proto: static QString QFile::readLink(const QString & fileName);
extern void C_ZN5QFile8readLinkERK7QString(void* arg0); // 4
  // proto:  QString QFile::readLink();
extern void C_ZNK5QFile8readLinkEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// rename(const class QString &, const class QString &)
func (this *QFile) rename_s(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN5QFile6renameERK7QStringS2_(arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN5QFile6renameERK7QString
    // invoke: bool rename(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN5QFile6renameERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFile", "rename", args)
  }

}

// exists(const class QString &)
func (this *QFile) exists_s(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN5QFile6existsERK7QString(arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK5QFile6existsEv
    // invoke: bool exists()
    var ret = C.C_ZNK5QFile6existsEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFile", "exists", args)
  }

}

// symLinkTarget()
func (this *QFile) symLinkTarget(args ...interface{}) () {
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
    var ret = C.C_ZNK5QFile13symLinkTargetEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN5QFile13symLinkTargetERK7QString
    // invoke: QString symLinkTarget(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN5QFile13symLinkTargetERK7QString(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFile", "symLinkTarget", args)
  }

}

// size()
func (this *QFile) size(args ...interface{}) () {
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
    var ret = C.C_ZNK5QFile4sizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFile", "size", args)
  }

}

// encodeName(const class QString &)
func (this *QFile) encodeName_s(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN5QFile10encodeNameERK7QString(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFile", "encodeName", args)
  }

}

// decodeName(const class QByteArray &)
func (this *QFile) decodeName_s(args ...interface{}) () {
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
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN5QFile10decodeNameERK10QByteArray(arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN5QFile10decodeNameEPKc
    // invoke: QString decodeName(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN5QFile10decodeNameEPKc(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFile", "decodeName", args)
  }

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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QFileC2ERK7QStringP7QObject(arg0, arg1)
    return &QFile{qclsinst:qthis}
  case 1:
    // invoke: _ZN5QFileC1Ev
    // invoke: void QFile()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QFileC2Ev()
    return &QFile{qclsinst:qthis}
  case 2:
    // invoke: _ZN5QFileC1EP7QObject
    // invoke: void QFile(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QFileC2EP7QObject(arg0)
    return &QFile{qclsinst:qthis}
  case 3:
    // invoke: _ZN5QFileC1ERK7QString
    // invoke: void QFile(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QFileC2ERK7QString(arg0)
    return &QFile{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QFile", "QFile", args)
  }

  return nil // QFile{qclsinst:qthis}
}

// fileName()
func (this *QFile) fileName(args ...interface{}) () {
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
    var ret = C.C_ZNK5QFile8fileNameEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFile", "fileName", args)
  }

}

// link(const class QString &)
func (this *QFile) link(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN5QFile4linkERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN5QFile4linkERK7QStringS2_
    // invoke: bool link(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN5QFile4linkERK7QStringS2_(arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFile", "link", args)
  }

}

// copy(const class QString &, const class QString &)
func (this *QFile) copy_s(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN5QFile4copyERK7QStringS2_(arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN5QFile4copyERK7QString
    // invoke: bool copy(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN5QFile4copyERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFile", "copy", args)
  }

}

// resize(qint64)
func (this *QFile) resize(args ...interface{}) () {
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
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN5QFile6resizeEx(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN5QFile6resizeERK7QStringx
    // invoke: bool resize(const class QString &, qint64)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int64_t(args[1].(int64))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN5QFile6resizeERK7QStringx(arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFile", "resize", args)
  }

}

// permissions()
func (this *QFile) permissions(args ...interface{}) () {
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
    C.C_ZNK5QFile11permissionsEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QFile11permissionsERK7QString
    // invoke: Permissions permissions(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QFile11permissionsERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QFile", "permissions", args)
  }

}

// metaObject()
func (this *QFile) metaObject(args ...interface{}) () {
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
    C.C_ZNK5QFile10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFile", "metaObject", args)
  }

}

// ~QFile()
func (this *QFile) FreeQFile(args ...interface{}) () {
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
    C.C_ZN5QFileD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFile", "~QFile", args)
  }

}

// setFileName(const class QString &)
func (this *QFile) setFileName(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QFile11setFileNameERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFile", "setFileName", args)
  }

}

// remove(const class QString &)
func (this *QFile) remove_s(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN5QFile6removeERK7QString(arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN5QFile6removeEv
    // invoke: bool remove()
    var ret = C.C_ZN5QFile6removeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFile", "remove", args)
  }

}

// readLink(const class QString &)
func (this *QFile) readLink_s(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN5QFile8readLinkERK7QString(arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK5QFile8readLinkEv
    // invoke: QString readLink()
    var ret = C.C_ZNK5QFile8readLinkEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFile", "readLink", args)
  }

}

// <= body block end

