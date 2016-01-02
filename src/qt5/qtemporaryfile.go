package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtCore/qtemporaryfile.h
// dst-file: /src/core/qtemporaryfile.go
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
  // proto:  bool QTemporaryFile::autoRemove();
extern void _ZNK14QTemporaryFile10autoRemoveEv(void* qthis);
  // proto: static QTemporaryFile * QTemporaryFile::createLocalFile(QFile & file);
extern void demth_ZN14QTemporaryFile15createLocalFileER5QFile(void* arg0);
  // proto:  void QTemporaryFile::QTemporaryFile(const QString & templateName);
extern void* dector_ZN14QTemporaryFileC1ERK7QString(void* arg0);
extern void _ZN14QTemporaryFileC1ERK7QString(void* qthis, void* arg0);
  // proto:  void QTemporaryFile::QTemporaryFile();
extern void* dector_ZN14QTemporaryFileC1Ev();
extern void _ZN14QTemporaryFileC1Ev(void* qthis);
  // proto:  void QTemporaryFile::QTemporaryFile(QObject * parent);
extern void* dector_ZN14QTemporaryFileC1EP7QObject(void* arg0);
extern void _ZN14QTemporaryFileC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QTemporaryFile::~QTemporaryFile();
extern void _ZN14QTemporaryFileD0Ev(void* qthis);
  // proto:  const QMetaObject * QTemporaryFile::metaObject();
extern void _ZNK14QTemporaryFile10metaObjectEv(void* qthis);
  // proto:  void QTemporaryFile::setAutoRemove(bool b);
extern void _ZN14QTemporaryFile13setAutoRemoveEb(void* qthis, bool arg0);
  // proto:  QString QTemporaryFile::fileName();
extern void _ZNK14QTemporaryFile8fileNameEv(void* qthis);
  // proto:  QString QTemporaryFile::fileTemplate();
extern void _ZNK14QTemporaryFile12fileTemplateEv(void* qthis);
  // proto: static QTemporaryFile * QTemporaryFile::createNativeFile(const QString & fileName);
extern void demth_ZN14QTemporaryFile16createNativeFileERK7QString(void* arg0);
  // proto:  bool QTemporaryFile::open();
extern void _ZN14QTemporaryFile4openEv(void* qthis);
  // proto: static QTemporaryFile * QTemporaryFile::createLocalFile(const QString & fileName);
extern void demth_ZN14QTemporaryFile15createLocalFileERK7QString(void* arg0);
  // proto: static QTemporaryFile * QTemporaryFile::createNativeFile(QFile & file);
extern void _ZN14QTemporaryFile16createNativeFileER5QFile(void* arg0);
  // proto:  void QTemporaryFile::setFileTemplate(const QString & name);
extern void _ZN14QTemporaryFile15setFileTemplateERK7QString(void* qthis, void* arg0);
  // proto:  void QTemporaryFile::QTemporaryFile(const QTemporaryFile & );
extern void* dector_ZN14QTemporaryFileC1ERKS_(void* arg0);
extern void _ZN14QTemporaryFileC1ERKS_(void* qthis, void* arg0);
  // proto:  void QTemporaryFile::QTemporaryFile(const QString & templateName, QObject * parent);
extern void* dector_ZN14QTemporaryFileC1ERK7QStringP7QObject(void* arg0, void* arg1);
extern void _ZN14QTemporaryFileC1ERK7QStringP7QObject(void* qthis, void* arg0, void* arg1);
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

// class sizeof(QTemporaryFile)=1
type QTemporaryFile struct {
  /*qbase*/ QFile;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  bool QTemporaryFile::autoRemove();
func (this *QTemporaryFile) autoRemove(args ...interface{}) () {
  // autoRemove()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTemporaryFile10autoRemoveEv
  default:
    qtrt.ErrorResolve("QTemporaryFile", "autoRemove", args)
  }

}

  // proto: static QTemporaryFile * QTemporaryFile::createLocalFile(QFile & file);
func (this *QTemporaryFile) createLocalFile_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTemporaryFile", "createLocalFile", args)
  }

}

  // proto:  void QTemporaryFile::QTemporaryFile(const QString & templateName);
func NewQTemporaryFile(args ...interface{}) QTemporaryFile {
  return QTemporaryFile{}
}

  // proto:  void QTemporaryFile::~QTemporaryFile();
func (this *QTemporaryFile) FreeQTemporaryFile(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTemporaryFile", "~QTemporaryFile", args)
  }

}

  // proto:  const QMetaObject * QTemporaryFile::metaObject();
func (this *QTemporaryFile) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTemporaryFile10metaObjectEv
  default:
    qtrt.ErrorResolve("QTemporaryFile", "metaObject", args)
  }

}

  // proto:  void QTemporaryFile::setAutoRemove(bool b);
func (this *QTemporaryFile) setAutoRemove(args ...interface{}) () {
  // setAutoRemove(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QTemporaryFile13setAutoRemoveEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTemporaryFile", "setAutoRemove", args)
  }

}

  // proto:  QString QTemporaryFile::fileName();
func (this *QTemporaryFile) fileName(args ...interface{}) () {
  // fileName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTemporaryFile8fileNameEv
  default:
    qtrt.ErrorResolve("QTemporaryFile", "fileName", args)
  }

}

  // proto:  QString QTemporaryFile::fileTemplate();
func (this *QTemporaryFile) fileTemplate(args ...interface{}) () {
  // fileTemplate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTemporaryFile12fileTemplateEv
  default:
    qtrt.ErrorResolve("QTemporaryFile", "fileTemplate", args)
  }

}

  // proto: static QTemporaryFile * QTemporaryFile::createNativeFile(const QString & fileName);
func (this *QTemporaryFile) createNativeFile_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTemporaryFile", "createNativeFile", args)
  }

}

  // proto:  bool QTemporaryFile::open();
func (this *QTemporaryFile) open(args ...interface{}) () {
  // open()
  // open(OpenMode)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int64Ty(false) // "OpenMode"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QTemporaryFile4openEv
  case 1:
    // invoke: _ZN14QTemporaryFile4openE6QFlagsIN9QIODevice12OpenModeFlagEE
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTemporaryFile", "open", args)
  }

}

  // proto:  void QTemporaryFile::setFileTemplate(const QString & name);
func (this *QTemporaryFile) setFileTemplate(args ...interface{}) () {
  // setFileTemplate(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QTemporaryFile15setFileTemplateERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTemporaryFile", "setFileTemplate", args)
  }

}

// <= body block end

