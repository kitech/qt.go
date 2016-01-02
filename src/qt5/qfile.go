package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
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
  // proto:  QString QFile::symLinkTarget();
extern void demth_ZNK5QFile13symLinkTargetEv(void* qthis);
  // proto:  void QFile::QFile();
extern void* dector_ZN5QFileC1Ev();
extern void _ZN5QFileC1Ev(void* qthis);
  // proto:  void QFile::QFile(QObject * parent);
extern void* dector_ZN5QFileC1EP7QObject(void* arg0);
extern void _ZN5QFileC1EP7QObject(void* qthis, void* arg0);
  // proto: static bool QFile::link(const QString & oldname, const QString & newName);
extern void _ZN5QFile4linkERK7QStringS2_(void* arg0, void* arg1);
  // proto: static bool QFile::rename(const QString & oldName, const QString & newName);
extern void _ZN5QFile6renameERK7QStringS2_(void* arg0, void* arg1);
  // proto:  bool QFile::link(const QString & newName);
extern void _ZN5QFile4linkERK7QString(void* qthis, void* arg0);
  // proto: static bool QFile::resize(const QString & filename, qint64 sz);
extern void _ZN5QFile6resizeERK7QStringx(void* arg0, long long arg1);
  // proto: static bool QFile::exists(const QString & fileName);
extern void _ZN5QFile6existsERK7QString(void* arg0);
  // proto:  void QFile::~QFile();
extern void _ZN5QFileD0Ev(void* qthis);
  // proto: static bool QFile::copy(const QString & fileName, const QString & newName);
extern void _ZN5QFile4copyERK7QStringS2_(void* arg0, void* arg1);
  // proto: static QString QFile::readLink(const QString & fileName);
extern void _ZN5QFile8readLinkERK7QString(void* arg0);
  // proto:  bool QFile::exists();
extern void _ZNK5QFile6existsEv(void* qthis);
  // proto:  qint64 QFile::size();
extern void _ZNK5QFile4sizeEv(void* qthis);
  // proto:  bool QFile::resize(qint64 sz);
extern void _ZN5QFile6resizeEx(void* qthis, long long arg0);
  // proto:  void QFile::QFile(const QFile & );
extern void* dector_ZN5QFileC1ERKS_(void* arg0);
extern void _ZN5QFileC1ERKS_(void* qthis, void* arg0);
  // proto:  void QFile::setFileName(const QString & name);
extern void _ZN5QFile11setFileNameERK7QString(void* qthis, void* arg0);
  // proto:  bool QFile::remove();
extern void _ZN5QFile6removeEv(void* qthis);
  // proto:  bool QFile::copy(const QString & newName);
extern void _ZN5QFile4copyERK7QString(void* qthis, void* arg0);
  // proto: static QByteArray QFile::encodeName(const QString & fileName);
extern void demth_ZN5QFile10encodeNameERK7QString(void* arg0);
  // proto: static QString QFile::decodeName(const QByteArray & localFileName);
extern void _ZN5QFile10decodeNameERK10QByteArray(void* arg0);
  // proto:  bool QFile::rename(const QString & newName);
extern void _ZN5QFile6renameERK7QString(void* qthis, void* arg0);
  // proto:  QString QFile::fileName();
extern void _ZNK5QFile8fileNameEv(void* qthis);
  // proto: static QString QFile::decodeName(const char * localFileName);
extern void demth_ZN5QFile10decodeNameEPKc(char* arg0);
  // proto:  const QMetaObject * QFile::metaObject();
extern void _ZNK5QFile10metaObjectEv(void* qthis);
  // proto:  void QFile::QFile(const QString & name, QObject * parent);
extern void* dector_ZN5QFileC1ERK7QStringP7QObject(void* arg0, void* arg1);
extern void _ZN5QFileC1ERK7QStringP7QObject(void* qthis, void* arg0, void* arg1);
  // proto: static QString QFile::symLinkTarget(const QString & fileName);
extern void demth_ZN5QFile13symLinkTargetERK7QString(void* arg0);
  // proto: static bool QFile::remove(const QString & fileName);
extern void _ZN5QFile6removeERK7QString(void* arg0);
  // proto:  void QFile::QFile(const QString & name);
extern void* dector_ZN5QFileC1ERK7QString(void* arg0);
extern void _ZN5QFileC1ERK7QString(void* qthis, void* arg0);
  // proto:  QString QFile::readLink();
extern void _ZNK5QFile8readLinkEv(void* qthis);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QString QFile::symLinkTarget();
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
  case 1:
    // invoke: _ZN5QFile13symLinkTargetERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFile", "symLinkTarget", args)
  }

}

  // proto:  void QFile::QFile();
func NewQFile(args ...interface{}) QFile {
  return QFile{}
}

  // proto: static bool QFile::link(const QString & oldname, const QString & newName);
func (this *QFile) link_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFile", "link", args)
  }

}

  // proto: static bool QFile::rename(const QString & oldName, const QString & newName);
func (this *QFile) rename_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFile", "rename", args)
  }

}

  // proto:  bool QFile::link(const QString & newName);
func (this *QFile) link(args ...interface{}) () {
  // link(const class QString &, const class QString &)
  // link(const class QString &)
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
    // invoke: _ZN5QFile4linkERK7QStringS2_
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN5QFile4linkERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFile", "link", args)
  }

}

  // proto: static bool QFile::resize(const QString & filename, qint64 sz);
func (this *QFile) resize_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFile", "resize", args)
  }

}

  // proto: static bool QFile::exists(const QString & fileName);
func (this *QFile) exists_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFile", "exists", args)
  }

}

  // proto:  void QFile::~QFile();
func (this *QFile) FreeQFile(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFile", "~QFile", args)
  }

}

  // proto: static bool QFile::copy(const QString & fileName, const QString & newName);
func (this *QFile) copy_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFile", "copy", args)
  }

}

  // proto: static QString QFile::readLink(const QString & fileName);
func (this *QFile) readLink_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFile", "readLink", args)
  }

}

  // proto:  bool QFile::exists();
func (this *QFile) exists(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK5QFile6existsEv
  default:
    qtrt.ErrorResolve("QFile", "exists", args)
  }

}

  // proto:  qint64 QFile::size();
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
  default:
    qtrt.ErrorResolve("QFile", "size", args)
  }

}

  // proto:  bool QFile::resize(qint64 sz);
func (this *QFile) resize(args ...interface{}) () {
  // resize(const class QString &, qint64)
  // resize(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.Int64Ty(false) // "qint64"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFile6resizeERK7QStringx
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int64_t(args[1].(int64))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN5QFile6resizeEx
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFile", "resize", args)
  }

}

  // proto:  void QFile::setFileName(const QString & name);
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFile", "setFileName", args)
  }

}

  // proto:  bool QFile::remove();
func (this *QFile) remove(args ...interface{}) () {
  // remove()
  // remove(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFile6removeEv
  case 1:
    // invoke: _ZN5QFile6removeERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFile", "remove", args)
  }

}

  // proto:  bool QFile::copy(const QString & newName);
func (this *QFile) copy(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN5QFile4copyERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFile", "copy", args)
  }

}

  // proto: static QByteArray QFile::encodeName(const QString & fileName);
func (this *QFile) encodeName_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFile", "encodeName", args)
  }

}

  // proto: static QString QFile::decodeName(const QByteArray & localFileName);
func (this *QFile) decodeName_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFile", "decodeName", args)
  }

}

  // proto:  bool QFile::rename(const QString & newName);
func (this *QFile) rename(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN5QFile6renameERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFile", "rename", args)
  }

}

  // proto:  QString QFile::fileName();
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
  default:
    qtrt.ErrorResolve("QFile", "fileName", args)
  }

}

  // proto:  const QMetaObject * QFile::metaObject();
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
  default:
    qtrt.ErrorResolve("QFile", "metaObject", args)
  }

}

  // proto: static QString QFile::symLinkTarget(const QString & fileName);
func (this *QFile) symLinkTarget_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFile", "symLinkTarget", args)
  }

}

  // proto: static bool QFile::remove(const QString & fileName);
func (this *QFile) remove_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFile", "remove", args)
  }

}

  // proto:  QString QFile::readLink();
func (this *QFile) readLink(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK5QFile8readLinkEv
  default:
    qtrt.ErrorResolve("QFile", "readLink", args)
  }

}

// <= body block end

