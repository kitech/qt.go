package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtCore/qmimedatabase.h
// dst-file: /src/core/qmimedatabase.go
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
  // proto:  QMimeType QMimeDatabase::mimeTypeForName(const QString & nameOrAlias);
extern void C_ZNK13QMimeDatabase15mimeTypeForNameERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QMimeDatabase::QMimeDatabase();
extern void* C_ZN13QMimeDatabaseC2Ev(); // 3
  // proto:  void QMimeDatabase::~QMimeDatabase();
extern void C_ZN13QMimeDatabaseD2Ev(void* qthis); // 4
  // proto:  QList<QMimeType> QMimeDatabase::allMimeTypes();
extern void C_ZNK13QMimeDatabase12allMimeTypesEv(void* qthis); // 4
  // proto:  QMimeType QMimeDatabase::mimeTypeForData(QIODevice * device);
extern void C_ZNK13QMimeDatabase15mimeTypeForDataEP9QIODevice(void* qthis, void* arg0); // 4
  // proto:  QMimeType QMimeDatabase::mimeTypeForData(const QByteArray & data);
extern void C_ZNK13QMimeDatabase15mimeTypeForDataERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  QMimeType QMimeDatabase::mimeTypeForUrl(const QUrl & url);
extern void C_ZNK13QMimeDatabase14mimeTypeForUrlERK4QUrl(void* qthis, void* arg0); // 4
  // proto:  QMimeType QMimeDatabase::mimeTypeForFileNameAndData(const QString & fileName, const QByteArray & data);
extern void C_ZNK13QMimeDatabase26mimeTypeForFileNameAndDataERK7QStringRK10QByteArray(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QMimeType QMimeDatabase::mimeTypeForFileNameAndData(const QString & fileName, QIODevice * device);
extern void C_ZNK13QMimeDatabase26mimeTypeForFileNameAndDataERK7QStringP9QIODevice(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QList<QMimeType> QMimeDatabase::mimeTypesForFileName(const QString & fileName);
extern void C_ZNK13QMimeDatabase20mimeTypesForFileNameERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QMimeDatabase::suffixForFileName(const QString & fileName);
extern void C_ZNK13QMimeDatabase17suffixForFileNameERK7QString(void* qthis, void* arg0); // 4
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

// class sizeof(QMimeDatabase)=8
type QMimeDatabase struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// mimeTypeForName(const class QString &)
func (this *QMimeDatabase) mimeTypeForName(args ...interface{}) () {
  // mimeTypeForName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMimeDatabase15mimeTypeForNameERK7QString
    // invoke: QMimeType mimeTypeForName(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QMimeDatabase15mimeTypeForNameERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMimeDatabase", "mimeTypeForName", args)
  }

}

// QMimeDatabase()
func NewQMimeDatabase(args ...interface{}) *QMimeDatabase {
  // QMimeDatabase()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QMimeDatabaseC1Ev
    // invoke: void QMimeDatabase()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QMimeDatabaseC2Ev()
    return &QMimeDatabase{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QMimeDatabase", "QMimeDatabase", args)
  }

  return nil // QMimeDatabase{qclsinst:qthis}
}

// ~QMimeDatabase()
func (this *QMimeDatabase) FreeQMimeDatabase(args ...interface{}) () {
  // ~QMimeDatabase()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QMimeDatabaseD0Ev
    // invoke: void ~QMimeDatabase()
    C.C_ZN13QMimeDatabaseD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeDatabase", "~QMimeDatabase", args)
  }

}

// allMimeTypes()
func (this *QMimeDatabase) allMimeTypes(args ...interface{}) () {
  // allMimeTypes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMimeDatabase12allMimeTypesEv
    // invoke: QList<QMimeType> allMimeTypes()
    C.C_ZNK13QMimeDatabase12allMimeTypesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeDatabase", "allMimeTypes", args)
  }

}

// mimeTypeForData(class QIODevice *)
func (this *QMimeDatabase) mimeTypeForData(args ...interface{}) () {
  // mimeTypeForData(class QIODevice *)
  // mimeTypeForData(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMimeDatabase15mimeTypeForDataEP9QIODevice
    // invoke: QMimeType mimeTypeForData(class QIODevice *)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QMimeDatabase15mimeTypeForDataEP9QIODevice(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK13QMimeDatabase15mimeTypeForDataERK10QByteArray
    // invoke: QMimeType mimeTypeForData(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QMimeDatabase15mimeTypeForDataERK10QByteArray(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMimeDatabase", "mimeTypeForData", args)
  }

}

// mimeTypeForUrl(const class QUrl &)
func (this *QMimeDatabase) mimeTypeForUrl(args ...interface{}) () {
  // mimeTypeForUrl(const class QUrl &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QUrl{}) // "const QUrl &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMimeDatabase14mimeTypeForUrlERK4QUrl
    // invoke: QMimeType mimeTypeForUrl(const class QUrl &)
    var arg0 = args[0].(QUrl).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QMimeDatabase14mimeTypeForUrlERK4QUrl(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMimeDatabase", "mimeTypeForUrl", args)
  }

}

// mimeTypeForFileNameAndData(const class QString &, const class QByteArray &)
func (this *QMimeDatabase) mimeTypeForFileNameAndData(args ...interface{}) () {
  // mimeTypeForFileNameAndData(const class QString &, const class QByteArray &)
  // mimeTypeForFileNameAndData(const class QString &, class QIODevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QIODevice{}) // "QIODevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMimeDatabase26mimeTypeForFileNameAndDataERK7QStringRK10QByteArray
    // invoke: QMimeType mimeTypeForFileNameAndData(const class QString &, const class QByteArray &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QByteArray).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK13QMimeDatabase26mimeTypeForFileNameAndDataERK7QStringRK10QByteArray(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK13QMimeDatabase26mimeTypeForFileNameAndDataERK7QStringP9QIODevice
    // invoke: QMimeType mimeTypeForFileNameAndData(const class QString &, class QIODevice *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QIODevice).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK13QMimeDatabase26mimeTypeForFileNameAndDataERK7QStringP9QIODevice(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMimeDatabase", "mimeTypeForFileNameAndData", args)
  }

}

// mimeTypesForFileName(const class QString &)
func (this *QMimeDatabase) mimeTypesForFileName(args ...interface{}) () {
  // mimeTypesForFileName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMimeDatabase20mimeTypesForFileNameERK7QString
    // invoke: QList<QMimeType> mimeTypesForFileName(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK13QMimeDatabase20mimeTypesForFileNameERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMimeDatabase", "mimeTypesForFileName", args)
  }

}

// suffixForFileName(const class QString &)
func (this *QMimeDatabase) suffixForFileName(args ...interface{}) () {
  // suffixForFileName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMimeDatabase17suffixForFileNameERK7QString
    // invoke: QString suffixForFileName(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QMimeDatabase17suffixForFileNameERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMimeDatabase", "suffixForFileName", args)
  }

}

// <= body block end

