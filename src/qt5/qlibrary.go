package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtCore/qlibrary.h
// dst-file: /src/core/qlibrary.go
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
  // proto:  bool QLibrary::load();
extern void _ZN8QLibrary4loadEv(void* qthis); // 4
  // proto: static QFunctionPointer QLibrary::resolve(const QString & fileName, const QString & version, const char * symbol);
extern void _ZN8QLibrary7resolveERK7QStringS2_PKc(void* arg0, void* arg1, unsigned char* arg2); // 4
  // proto:  QFunctionPointer QLibrary::resolve(const char * symbol);
extern void _ZN8QLibrary7resolveEPKc(void* qthis, unsigned char* arg0); // 4
  // proto: static QFunctionPointer QLibrary::resolve(const QString & fileName, int verNum, const char * symbol);
extern void _ZN8QLibrary7resolveERK7QStringiPKc(void* arg0, int32_t arg1, unsigned char* arg2); // 4
  // proto: static QFunctionPointer QLibrary::resolve(const QString & fileName, const char * symbol);
extern void _ZN8QLibrary7resolveERK7QStringPKc(void* arg0, unsigned char* arg1); // 4
  // proto:  const QMetaObject * QLibrary::metaObject();
extern void _ZNK8QLibrary10metaObjectEv(void* qthis); // 4
  // proto:  QString QLibrary::errorString();
extern void _ZNK8QLibrary11errorStringEv(void* qthis); // 4
  // proto:  bool QLibrary::unload();
extern void _ZN8QLibrary6unloadEv(void* qthis); // 4
  // proto:  void QLibrary::setFileNameAndVersion(const QString & fileName, const QString & version);
extern void _ZN8QLibrary21setFileNameAndVersionERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QLibrary::setFileNameAndVersion(const QString & fileName, int verNum);
extern void _ZN8QLibrary21setFileNameAndVersionERK7QStringi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QLibrary::setFileName(const QString & fileName);
extern void _ZN8QLibrary11setFileNameERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QLibrary::isLoaded();
extern void _ZNK8QLibrary8isLoadedEv(void* qthis); // 4
  // proto:  void QLibrary::~QLibrary();
extern void _ZN8QLibraryD2Ev(void* qthis); // 4
  // proto: static bool QLibrary::isLibrary(const QString & fileName);
extern void _ZN8QLibrary9isLibraryERK7QString(void* arg0); // 4
  // proto:  void QLibrary::QLibrary(QObject * parent);
extern void _ZN8QLibraryC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  void QLibrary::QLibrary(const QString & fileName, int verNum, QObject * parent);
extern void _ZN8QLibraryC2ERK7QStringiP7QObject(void* qthis, void* arg0, int32_t arg1, void* arg2); // 3
  // proto:  void QLibrary::QLibrary(const QString & fileName, const QString & version, QObject * parent);
extern void _ZN8QLibraryC2ERK7QStringS2_P7QObject(void* qthis, void* arg0, void* arg1, void* arg2); // 3
  // proto:  void QLibrary::QLibrary(const QString & fileName, QObject * parent);
extern void _ZN8QLibraryC2ERK7QStringP7QObject(void* qthis, void* arg0, void* arg1); // 3
  // proto:  QString QLibrary::fileName();
extern void _ZNK8QLibrary8fileNameEv(void* qthis); // 4
  // proto:  LoadHints QLibrary::loadHints();
extern void _ZNK8QLibrary9loadHintsEv(void* qthis); // 4
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

// class sizeof(QLibrary)=1
type QLibrary struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
}

// load()
func (this *QLibrary) load(args ...interface{}) () {
  // load()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QLibrary4loadEv
    // invoke: bool load()
    C._ZN8QLibrary4loadEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLibrary", "load", args)
  }

}

// resolve(const class QString &, const class QString &, const char *)
func (this *QLibrary) resolve_s(args ...interface{}) () {
  // resolve(const class QString &, const class QString &, const char *)
  // resolve(const char *)
  // resolve(const class QString &, int, const char *)
  // resolve(const class QString &, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.ByteTy(true) // "const char *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QLibrary7resolveERK7QStringS2_PKc
    // invoke: QFunctionPointer resolve(const class QString &, const class QString &, const char *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[2].([]byte)).Pointer()))
    if false {fmt.Println(arg2)}
    C._ZN8QLibrary7resolveERK7QStringS2_PKc(arg0, arg1, arg2)
  case 1:
    // invoke: _ZN8QLibrary7resolveEPKc
    // invoke: QFunctionPointer resolve(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C._ZN8QLibrary7resolveEPKc(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN8QLibrary7resolveERK7QStringiPKc
    // invoke: QFunctionPointer resolve(const class QString &, int, const char *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[2].([]byte)).Pointer()))
    if false {fmt.Println(arg2)}
    C._ZN8QLibrary7resolveERK7QStringiPKc(arg0, arg1, arg2)
  case 3:
    // invoke: _ZN8QLibrary7resolveERK7QStringPKc
    // invoke: QFunctionPointer resolve(const class QString &, const char *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    C._ZN8QLibrary7resolveERK7QStringPKc(arg0, arg1)
  default:
    qtrt.ErrorResolve("QLibrary", "resolve", args)
  }

}

// metaObject()
func (this *QLibrary) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QLibrary10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK8QLibrary10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLibrary", "metaObject", args)
  }

}

// errorString()
func (this *QLibrary) errorString(args ...interface{}) () {
  // errorString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QLibrary11errorStringEv
    // invoke: QString errorString()
    C._ZNK8QLibrary11errorStringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLibrary", "errorString", args)
  }

}

// unload()
func (this *QLibrary) unload(args ...interface{}) () {
  // unload()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QLibrary6unloadEv
    // invoke: bool unload()
    C._ZN8QLibrary6unloadEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLibrary", "unload", args)
  }

}

// setFileNameAndVersion(const class QString &, const class QString &)
func (this *QLibrary) setFileNameAndVersion(args ...interface{}) () {
  // setFileNameAndVersion(const class QString &, const class QString &)
  // setFileNameAndVersion(const class QString &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QLibrary21setFileNameAndVersionERK7QStringS2_
    // invoke: void setFileNameAndVersion(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN8QLibrary21setFileNameAndVersionERK7QStringS2_(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QLibrary21setFileNameAndVersionERK7QStringi
    // invoke: void setFileNameAndVersion(const class QString &, int)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN8QLibrary21setFileNameAndVersionERK7QStringi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLibrary", "setFileNameAndVersion", args)
  }

}

// setFileName(const class QString &)
func (this *QLibrary) setFileName(args ...interface{}) () {
  // setFileName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QLibrary11setFileNameERK7QString
    // invoke: void setFileName(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QLibrary11setFileNameERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLibrary", "setFileName", args)
  }

}

// isLoaded()
func (this *QLibrary) isLoaded(args ...interface{}) () {
  // isLoaded()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QLibrary8isLoadedEv
    // invoke: bool isLoaded()
    C._ZNK8QLibrary8isLoadedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLibrary", "isLoaded", args)
  }

}

// ~QLibrary()
func (this *QLibrary) FreeQLibrary(args ...interface{}) () {
  // ~QLibrary()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QLibraryD0Ev
    // invoke: void ~QLibrary()
    C._ZN8QLibraryD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLibrary", "~QLibrary", args)
  }

}

// isLibrary(const class QString &)
func (this *QLibrary) isLibrary_s(args ...interface{}) () {
  // isLibrary(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QLibrary9isLibraryERK7QString
    // invoke: bool isLibrary(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QLibrary9isLibraryERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QLibrary", "isLibrary", args)
  }

}

// QLibrary(class QObject *)
func NewQLibrary(args ...interface{}) QLibrary {
  // QLibrary(class QObject *)
  // QLibrary(const class QString &, int, class QObject *)
  // QLibrary(const class QString &, const class QString &, class QObject *)
  // QLibrary(const class QString &, class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][2] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3][1] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QLibraryC1EP7QObject
    // invoke: void QLibrary(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QLibraryC2EP7QObject(qthis, arg0)
  case 1:
    // invoke: _ZN8QLibraryC1ERK7QStringiP7QObject
    // invoke: void QLibrary(const class QString &, int, class QObject *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QObject).qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QLibraryC2ERK7QStringiP7QObject(qthis, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN8QLibraryC1ERK7QStringS2_P7QObject
    // invoke: void QLibrary(const class QString &, const class QString &, class QObject *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QObject).qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QLibraryC2ERK7QStringS2_P7QObject(qthis, arg0, arg1, arg2)
  case 3:
    // invoke: _ZN8QLibraryC1ERK7QStringP7QObject
    // invoke: void QLibrary(const class QString &, class QObject *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QLibraryC2ERK7QStringP7QObject(qthis, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLibrary", "QLibrary", args)
  }

  return QLibrary{}
}

// fileName()
func (this *QLibrary) fileName(args ...interface{}) () {
  // fileName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QLibrary8fileNameEv
    // invoke: QString fileName()
    C._ZNK8QLibrary8fileNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLibrary", "fileName", args)
  }

}

// loadHints()
func (this *QLibrary) loadHints(args ...interface{}) () {
  // loadHints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QLibrary9loadHintsEv
    // invoke: LoadHints loadHints()
    C._ZNK8QLibrary9loadHintsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLibrary", "loadHints", args)
  }

}

// <= body block end

