package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtCore/qresource.h
// dst-file: /src/core/qresource.go
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
  // proto: static QStringList QResource::searchPaths();
extern void C_ZN9QResource11searchPathsEv(); // 4
  // proto:  bool QResource::isCompressed();
extern void C_ZNK9QResource12isCompressedEv(void* qthis); // 4
  // proto:  void QResource::setLocale(const QLocale & locale);
extern void C_ZN9QResource9setLocaleERK7QLocale(void* qthis, void* arg0); // 4
  // proto:  QLocale QResource::locale();
extern void C_ZNK9QResource6localeEv(void* qthis); // 4
  // proto:  QString QResource::absoluteFilePath();
extern void C_ZNK9QResource16absoluteFilePathEv(void* qthis); // 4
  // proto:  void QResource::~QResource();
extern void C_ZN9QResourceD2Ev(void* qthis); // 4
  // proto:  void QResource::QResource(const QString & file, const QLocale & locale);
extern void C_ZN9QResourceC2ERK7QStringRK7QLocale(void* qthis, void* arg0, void* arg1); // 3
  // proto:  void QResource::setFileName(const QString & file);
extern void C_ZN9QResource11setFileNameERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QResource::fileName();
extern void C_ZNK9QResource8fileNameEv(void* qthis); // 4
  // proto: static bool QResource::unregisterResource(const QString & rccFilename, const QString & resourceRoot);
extern void C_ZN9QResource18unregisterResourceERK7QStringS2_(void* arg0, void* arg1); // 4
  // proto: static bool QResource::unregisterResource(const uchar * rccData, const QString & resourceRoot);
extern void C_ZN9QResource18unregisterResourceEPKhRK7QString(unsigned char* arg0, void* arg1); // 4
  // proto: static bool QResource::registerResource(const uchar * rccData, const QString & resourceRoot);
extern void C_ZN9QResource16registerResourceEPKhRK7QString(unsigned char* arg0, void* arg1); // 4
  // proto: static bool QResource::registerResource(const QString & rccFilename, const QString & resourceRoot);
extern void C_ZN9QResource16registerResourceERK7QStringS2_(void* arg0, void* arg1); // 4
  // proto:  bool QResource::isValid();
extern void C_ZNK9QResource7isValidEv(void* qthis); // 4
  // proto: static void QResource::addSearchPath(const QString & path);
extern void C_ZN9QResource13addSearchPathERK7QString(void* arg0); // 4
  // proto:  const uchar * QResource::data();
extern void C_ZNK9QResource4dataEv(void* qthis); // 4
  // proto:  qint64 QResource::size();
extern void C_ZNK9QResource4sizeEv(void* qthis); // 4
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

// class sizeof(QResource)=1
type QResource struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// searchPaths()
func (this *QResource) searchPaths_s(args ...interface{}) () {
  // searchPaths()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QResource11searchPathsEv
    // invoke: QStringList searchPaths()
    C.C_ZN9QResource11searchPathsEv()
  default:
    qtrt.ErrorResolve("QResource", "searchPaths", args)
  }

}

// isCompressed()
func (this *QResource) isCompressed(args ...interface{}) () {
  // isCompressed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QResource12isCompressedEv
    // invoke: bool isCompressed()
    C.C_ZNK9QResource12isCompressedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QResource", "isCompressed", args)
  }

}

// setLocale(const class QLocale &)
func (this *QResource) setLocale(args ...interface{}) () {
  // setLocale(const class QLocale &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLocale{}) // "const QLocale &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QResource9setLocaleERK7QLocale
    // invoke: void setLocale(const class QLocale &)
    var arg0 = args[0].(QLocale).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QResource9setLocaleERK7QLocale(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QResource", "setLocale", args)
  }

}

// locale()
func (this *QResource) locale(args ...interface{}) () {
  // locale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QResource6localeEv
    // invoke: QLocale locale()
    C.C_ZNK9QResource6localeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QResource", "locale", args)
  }

}

// absoluteFilePath()
func (this *QResource) absoluteFilePath(args ...interface{}) () {
  // absoluteFilePath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QResource16absoluteFilePathEv
    // invoke: QString absoluteFilePath()
    C.C_ZNK9QResource16absoluteFilePathEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QResource", "absoluteFilePath", args)
  }

}

// ~QResource()
func (this *QResource) FreeQResource(args ...interface{}) () {
  // ~QResource()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QResourceD0Ev
    // invoke: void ~QResource()
    C.C_ZN9QResourceD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QResource", "~QResource", args)
  }

}

// QResource(const class QString &, const class QLocale &)
func NewQResource(args ...interface{}) QResource {
  // QResource(const class QString &, const class QLocale &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QLocale{}) // "const QLocale &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QResourceC1ERK7QStringRK7QLocale
    // invoke: void QResource(const class QString &, const class QLocale &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QLocale).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN9QResourceC2ERK7QStringRK7QLocale(qthis, arg0, arg1)
  default:
    qtrt.ErrorResolve("QResource", "QResource", args)
  }

  return QResource{}
}

// setFileName(const class QString &)
func (this *QResource) setFileName(args ...interface{}) () {
  // setFileName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QResource11setFileNameERK7QString
    // invoke: void setFileName(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QResource11setFileNameERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QResource", "setFileName", args)
  }

}

// fileName()
func (this *QResource) fileName(args ...interface{}) () {
  // fileName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QResource8fileNameEv
    // invoke: QString fileName()
    C.C_ZNK9QResource8fileNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QResource", "fileName", args)
  }

}

// unregisterResource(const class QString &, const class QString &)
func (this *QResource) unregisterResource_s(args ...interface{}) () {
  // unregisterResource(const class QString &, const class QString &)
  // unregisterResource(const uchar *, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const uchar *"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QResource18unregisterResourceERK7QStringS2_
    // invoke: bool unregisterResource(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN9QResource18unregisterResourceERK7QStringS2_(arg0, arg1)
  case 1:
    // invoke: _ZN9QResource18unregisterResourceEPKhRK7QString
    // invoke: bool unregisterResource(const uchar *, const class QString &)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN9QResource18unregisterResourceEPKhRK7QString(arg0, arg1)
  default:
    qtrt.ErrorResolve("QResource", "unregisterResource", args)
  }

}

// registerResource(const uchar *, const class QString &)
func (this *QResource) registerResource_s(args ...interface{}) () {
  // registerResource(const uchar *, const class QString &)
  // registerResource(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const uchar *"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QResource16registerResourceEPKhRK7QString
    // invoke: bool registerResource(const uchar *, const class QString &)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN9QResource16registerResourceEPKhRK7QString(arg0, arg1)
  case 1:
    // invoke: _ZN9QResource16registerResourceERK7QStringS2_
    // invoke: bool registerResource(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN9QResource16registerResourceERK7QStringS2_(arg0, arg1)
  default:
    qtrt.ErrorResolve("QResource", "registerResource", args)
  }

}

// isValid()
func (this *QResource) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QResource7isValidEv
    // invoke: bool isValid()
    C.C_ZNK9QResource7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QResource", "isValid", args)
  }

}

// addSearchPath(const class QString &)
func (this *QResource) addSearchPath_s(args ...interface{}) () {
  // addSearchPath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QResource13addSearchPathERK7QString
    // invoke: void addSearchPath(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QResource13addSearchPathERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QResource", "addSearchPath", args)
  }

}

// data()
func (this *QResource) data(args ...interface{}) () {
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QResource4dataEv
    // invoke: const uchar * data()
    C.C_ZNK9QResource4dataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QResource", "data", args)
  }

}

// size()
func (this *QResource) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QResource4sizeEv
    // invoke: qint64 size()
    C.C_ZNK9QResource4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QResource", "size", args)
  }

}

// <= body block end

