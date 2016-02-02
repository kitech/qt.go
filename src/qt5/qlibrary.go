package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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
extern bool C_ZN8QLibrary4loadEv(void* qthis); // 4
  // proto: static QFunctionPointer QLibrary::resolve(const QString & fileName, const QString & version, const char * symbol);
extern void C_ZN8QLibrary7resolveERK7QStringS2_PKc(void* arg0, void* arg1, void* arg2); // 4
  // proto:  QFunctionPointer QLibrary::resolve(const char * symbol);
extern void C_ZN8QLibrary7resolveEPKc(void* qthis, void* arg0); // 4
  // proto: static QFunctionPointer QLibrary::resolve(const QString & fileName, int verNum, const char * symbol);
extern void C_ZN8QLibrary7resolveERK7QStringiPKc(void* arg0, int32_t arg1, void* arg2); // 4
  // proto: static QFunctionPointer QLibrary::resolve(const QString & fileName, const char * symbol);
extern void C_ZN8QLibrary7resolveERK7QStringPKc(void* arg0, void* arg1); // 4
  // proto:  const QMetaObject * QLibrary::metaObject();
extern void C_ZNK8QLibrary10metaObjectEv(void* qthis); // 4
  // proto:  QString QLibrary::errorString();
extern void* C_ZNK8QLibrary11errorStringEv(void* qthis); // 4
  // proto:  bool QLibrary::unload();
extern bool C_ZN8QLibrary6unloadEv(void* qthis); // 4
  // proto:  void QLibrary::setFileNameAndVersion(const QString & fileName, const QString & version);
extern void C_ZN8QLibrary21setFileNameAndVersionERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QLibrary::setFileNameAndVersion(const QString & fileName, int verNum);
extern void C_ZN8QLibrary21setFileNameAndVersionERK7QStringi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QLibrary::setFileName(const QString & fileName);
extern void C_ZN8QLibrary11setFileNameERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QLibrary::isLoaded();
extern bool C_ZNK8QLibrary8isLoadedEv(void* qthis); // 4
  // proto:  void QLibrary::~QLibrary();
extern void C_ZN8QLibraryD2Ev(void* qthis); // 4
  // proto: static bool QLibrary::isLibrary(const QString & fileName);
extern bool C_ZN8QLibrary9isLibraryERK7QString(void* arg0); // 4
  // proto:  void QLibrary::QLibrary(QObject * parent);
extern void* C_ZN8QLibraryC2EP7QObject(void* arg0); // 3
  // proto:  void QLibrary::QLibrary(const QString & fileName, int verNum, QObject * parent);
extern void* C_ZN8QLibraryC2ERK7QStringiP7QObject(void* arg0, int32_t arg1, void* arg2); // 3
  // proto:  void QLibrary::QLibrary(const QString & fileName, const QString & version, QObject * parent);
extern void* C_ZN8QLibraryC2ERK7QStringS2_P7QObject(void* arg0, void* arg1, void* arg2); // 3
  // proto:  void QLibrary::QLibrary(const QString & fileName, QObject * parent);
extern void* C_ZN8QLibraryC2ERK7QStringP7QObject(void* arg0, void* arg1); // 3
  // proto:  QString QLibrary::fileName();
extern void* C_ZNK8QLibrary8fileNameEv(void* qthis); // 4
  // proto:  LoadHints QLibrary::loadHints();
extern void C_ZNK8QLibrary9loadHintsEv(void* qthis); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// load()
func (this *QLibrary) Load(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN8QLibrary4loadEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLibrary", "load", args)
  }

  return
}

// resolve(const class QString &, const class QString &, const char *)
func (this *QLibrary) Resolve_S(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    argif2, free2 := qtrt.HandyConvert2c(args[2], vtys[0][2])
    var arg2 = argif2.(unsafe.Pointer)
    if false {fmt.Println(argif2, arg2)}
    if free2 {defer C.free(arg2)}
    C.C_ZN8QLibrary7resolveERK7QStringS2_PKc(arg0, arg1, arg2)
  case 1:
    // invoke: _ZN8QLibrary7resolveEPKc
    // invoke: QFunctionPointer resolve(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    C.C_ZN8QLibrary7resolveEPKc(this.Qclsinst, arg0)
  case 2:
    // invoke: _ZN8QLibrary7resolveERK7QStringiPKc
    // invoke: QFunctionPointer resolve(const class QString &, int, const char *)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    argif2, free2 := qtrt.HandyConvert2c(args[2], vtys[2][2])
    var arg2 = argif2.(unsafe.Pointer)
    if false {fmt.Println(argif2, arg2)}
    if free2 {defer C.free(arg2)}
    C.C_ZN8QLibrary7resolveERK7QStringiPKc(arg0, arg1, arg2)
  case 3:
    // invoke: _ZN8QLibrary7resolveERK7QStringPKc
    // invoke: QFunctionPointer resolve(const class QString &, const char *)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[3][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    C.C_ZN8QLibrary7resolveERK7QStringPKc(arg0, arg1)
  default:
    qtrt.ErrorResolve("QLibrary", "resolve", args)
  }

  return
}

// metaObject()
func (this *QLibrary) Metaobject(args ...interface{}) () {
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
    C.C_ZNK8QLibrary10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLibrary", "metaObject", args)
  }

  return
}

// errorString()
func (this *QLibrary) Errorstring(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QLibrary11errorStringEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLibrary", "errorString", args)
  }

  return
}

// unload()
func (this *QLibrary) Unload(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN8QLibrary6unloadEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLibrary", "unload", args)
  }

  return
}

// setFileNameAndVersion(const class QString &, const class QString &)
func (this *QLibrary) Setfilenameandversion(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QLibrary21setFileNameAndVersionERK7QStringS2_(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QLibrary21setFileNameAndVersionERK7QStringi
    // invoke: void setFileNameAndVersion(const class QString &, int)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN8QLibrary21setFileNameAndVersionERK7QStringi(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLibrary", "setFileNameAndVersion", args)
  }

  return
}

// setFileName(const class QString &)
func (this *QLibrary) Setfilename(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QLibrary11setFileNameERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLibrary", "setFileName", args)
  }

  return
}

// isLoaded()
func (this *QLibrary) Isloaded(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QLibrary8isLoadedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLibrary", "isLoaded", args)
  }

  return
}

// ~QLibrary()
func (this *QLibrary) Freeqlibrary(args ...interface{}) () {
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
    C.C_ZN8QLibraryD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLibrary", "~QLibrary", args)
  }

  return
}

// isLibrary(const class QString &)
func (this *QLibrary) Islibrary_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN8QLibrary9isLibraryERK7QString(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLibrary", "isLibrary", args)
  }

  return
}

// QLibrary(class QObject *)
func NewQLibrary(args ...interface{}) *QLibrary {
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
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QLibraryC2EP7QObject(arg0)
    return &QLibrary{Qclsinst:qthis}
  case 1:
    // invoke: _ZN8QLibraryC1ERK7QStringiP7QObject
    // invoke: void QLibrary(const class QString &, int, class QObject *)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QObject).Qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QLibraryC2ERK7QStringiP7QObject(arg0, arg1, arg2)
    return &QLibrary{Qclsinst:qthis}
  case 2:
    // invoke: _ZN8QLibraryC1ERK7QStringS2_P7QObject
    // invoke: void QLibrary(const class QString &, const class QString &, class QObject *)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QObject).Qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QLibraryC2ERK7QStringS2_P7QObject(arg0, arg1, arg2)
    return &QLibrary{Qclsinst:qthis}
  case 3:
    // invoke: _ZN8QLibraryC1ERK7QStringP7QObject
    // invoke: void QLibrary(const class QString &, class QObject *)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QObject).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QLibraryC2ERK7QStringP7QObject(arg0, arg1)
    return &QLibrary{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QLibrary", "QLibrary", args)
  }

  return nil // QLibrary{Qclsinst:qthis}
}

// fileName()
func (this *QLibrary) Filename(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QLibrary8fileNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLibrary", "fileName", args)
  }

  return
}

// loadHints()
func (this *QLibrary) Loadhints(args ...interface{}) () {
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
    C.C_ZNK8QLibrary9loadHintsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLibrary", "loadHints", args)
  }

  return
}

// <= body block end

