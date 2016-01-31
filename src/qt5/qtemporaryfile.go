package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  bool QTemporaryFile::autoRemove();
extern void C_ZNK14QTemporaryFile10autoRemoveEv(void* qthis); // 4
  // proto:  const QMetaObject * QTemporaryFile::metaObject();
extern void C_ZNK14QTemporaryFile10metaObjectEv(void* qthis); // 4
  // proto: static QTemporaryFile * QTemporaryFile::createNativeFile(const QString & fileName);
extern void C_ZN14QTemporaryFile16createNativeFileERK7QString(void* arg0); // 2
  // proto: static QTemporaryFile * QTemporaryFile::createNativeFile(QFile & file);
extern void C_ZN14QTemporaryFile16createNativeFileER5QFile(void* arg0); // 4
  // proto:  void QTemporaryFile::~QTemporaryFile();
extern void C_ZN14QTemporaryFileD2Ev(void* qthis); // 4
  // proto: static QTemporaryFile * QTemporaryFile::createLocalFile(QFile & file);
extern void C_ZN14QTemporaryFile15createLocalFileER5QFile(void* arg0); // 2
  // proto: static QTemporaryFile * QTemporaryFile::createLocalFile(const QString & fileName);
extern void C_ZN14QTemporaryFile15createLocalFileERK7QString(void* arg0); // 2
  // proto:  void QTemporaryFile::setAutoRemove(bool b);
extern void C_ZN14QTemporaryFile13setAutoRemoveEb(void* qthis, bool arg0); // 4
  // proto:  void QTemporaryFile::QTemporaryFile(QObject * parent);
extern void* C_ZN14QTemporaryFileC2EP7QObject(void* arg0); // 3
  // proto:  void QTemporaryFile::QTemporaryFile(const QString & templateName);
extern void* C_ZN14QTemporaryFileC2ERK7QString(void* arg0); // 3
  // proto:  void QTemporaryFile::QTemporaryFile(const QString & templateName, QObject * parent);
extern void* C_ZN14QTemporaryFileC2ERK7QStringP7QObject(void* arg0, void* arg1); // 3
  // proto:  void QTemporaryFile::QTemporaryFile();
extern void* C_ZN14QTemporaryFileC2Ev(); // 3
  // proto:  QString QTemporaryFile::fileName();
extern void C_ZNK14QTemporaryFile8fileNameEv(void* qthis); // 4
  // proto:  bool QTemporaryFile::open();
extern void C_ZN14QTemporaryFile4openEv(void* qthis); // 2
  // proto:  void QTemporaryFile::setFileTemplate(const QString & name);
extern void C_ZN14QTemporaryFile15setFileTemplateERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QTemporaryFile::fileTemplate();
extern void C_ZNK14QTemporaryFile12fileTemplateEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// autoRemove()
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
    // invoke: bool autoRemove()
    var ret = C.C_ZNK14QTemporaryFile10autoRemoveEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTemporaryFile", "autoRemove", args)
  }

}

// metaObject()
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
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK14QTemporaryFile10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTemporaryFile", "metaObject", args)
  }

}

// createNativeFile(const class QString &)
func (this *QTemporaryFile) createNativeFile_s(args ...interface{}) () {
  // createNativeFile(const class QString &)
  // createNativeFile(class QFile &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QFile{}) // "QFile &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QTemporaryFile16createNativeFileERK7QString
    // invoke: QTemporaryFile * createNativeFile(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN14QTemporaryFile16createNativeFileERK7QString(arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN14QTemporaryFile16createNativeFileER5QFile
    // invoke: QTemporaryFile * createNativeFile(class QFile &)
    var arg0 = args[0].(QFile).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN14QTemporaryFile16createNativeFileER5QFile(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTemporaryFile", "createNativeFile", args)
  }

}

// ~QTemporaryFile()
func (this *QTemporaryFile) FreeQTemporaryFile(args ...interface{}) () {
  // ~QTemporaryFile()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QTemporaryFileD0Ev
    // invoke: void ~QTemporaryFile()
    C.C_ZN14QTemporaryFileD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTemporaryFile", "~QTemporaryFile", args)
  }

}

// createLocalFile(class QFile &)
func (this *QTemporaryFile) createLocalFile_s(args ...interface{}) () {
  // createLocalFile(class QFile &)
  // createLocalFile(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFile{}) // "QFile &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QTemporaryFile15createLocalFileER5QFile
    // invoke: QTemporaryFile * createLocalFile(class QFile &)
    var arg0 = args[0].(QFile).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN14QTemporaryFile15createLocalFileER5QFile(arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN14QTemporaryFile15createLocalFileERK7QString
    // invoke: QTemporaryFile * createLocalFile(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN14QTemporaryFile15createLocalFileERK7QString(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTemporaryFile", "createLocalFile", args)
  }

}

// setAutoRemove(_Bool)
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
    // invoke: void setAutoRemove(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN14QTemporaryFile13setAutoRemoveEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTemporaryFile", "setAutoRemove", args)
  }

}

// QTemporaryFile(class QObject *)
func NewQTemporaryFile(args ...interface{}) *QTemporaryFile {
  // QTemporaryFile(class QObject *)
  // QTemporaryFile(const class QString &)
  // QTemporaryFile(const class QString &, class QObject *)
  // QTemporaryFile()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][1] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[3] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QTemporaryFileC1EP7QObject
    // invoke: void QTemporaryFile(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QTemporaryFileC2EP7QObject(arg0)
    return &QTemporaryFile{qclsinst:qthis}
  case 1:
    // invoke: _ZN14QTemporaryFileC1ERK7QString
    // invoke: void QTemporaryFile(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QTemporaryFileC2ERK7QString(arg0)
    return &QTemporaryFile{qclsinst:qthis}
  case 2:
    // invoke: _ZN14QTemporaryFileC1ERK7QStringP7QObject
    // invoke: void QTemporaryFile(const class QString &, class QObject *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QTemporaryFileC2ERK7QStringP7QObject(arg0, arg1)
    return &QTemporaryFile{qclsinst:qthis}
  case 3:
    // invoke: _ZN14QTemporaryFileC1Ev
    // invoke: void QTemporaryFile()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QTemporaryFileC2Ev()
    return &QTemporaryFile{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTemporaryFile", "QTemporaryFile", args)
  }

  return nil // QTemporaryFile{qclsinst:qthis}
}

// fileName()
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
    // invoke: QString fileName()
    var ret = C.C_ZNK14QTemporaryFile8fileNameEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTemporaryFile", "fileName", args)
  }

}

// open()
func (this *QTemporaryFile) open(args ...interface{}) () {
  // open()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QTemporaryFile4openEv
    // invoke: bool open()
    var ret = C.C_ZN14QTemporaryFile4openEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTemporaryFile", "open", args)
  }

}

// setFileTemplate(const class QString &)
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
    // invoke: void setFileTemplate(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QTemporaryFile15setFileTemplateERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTemporaryFile", "setFileTemplate", args)
  }

}

// fileTemplate()
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
    // invoke: QString fileTemplate()
    var ret = C.C_ZNK14QTemporaryFile12fileTemplateEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTemporaryFile", "fileTemplate", args)
  }

}

// <= body block end

