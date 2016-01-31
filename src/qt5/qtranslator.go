package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtCore/qtranslator.h
// dst-file: /src/core/qtranslator.go
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
  // proto:  bool QTranslator::load(const QString & filename, const QString & directory, const QString & search_delimiters, const QString & suffix);
extern void C_ZN11QTranslator4loadERK7QStringS2_S2_S2_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3); // 4
  // proto:  bool QTranslator::load(const QLocale & locale, const QString & filename, const QString & prefix, const QString & directory, const QString & suffix);
extern void C_ZN11QTranslator4loadERK7QLocaleRK7QStringS5_S5_S5_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3, void* arg4); // 4
  // proto:  bool QTranslator::load(const uchar * data, int len, const QString & directory);
extern void C_ZN11QTranslator4loadEPKhiRK7QString(void* qthis, unsigned char* arg0, int32_t arg1, void* arg2); // 4
  // proto:  const QMetaObject * QTranslator::metaObject();
extern void C_ZNK11QTranslator10metaObjectEv(void* qthis); // 4
  // proto:  void QTranslator::~QTranslator();
extern void C_ZN11QTranslatorD2Ev(void* qthis); // 4
  // proto:  bool QTranslator::isEmpty();
extern void C_ZNK11QTranslator7isEmptyEv(void* qthis); // 4
  // proto:  void QTranslator::QTranslator(QObject * parent);
extern void* C_ZN11QTranslatorC2EP7QObject(void* arg0); // 3
  // proto:  QString QTranslator::translate(const char * context, const char * sourceText, const char * disambiguation, int n);
extern void C_ZNK11QTranslator9translateEPKcS1_S1_i(void* qthis, unsigned char* arg0, unsigned char* arg1, unsigned char* arg2, int32_t arg3); // 4
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

// class sizeof(QTranslator)=1
type QTranslator struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
}

// load(const class QString &, const class QString &, const class QString &, const class QString &)
func (this *QTranslator) load(args ...interface{}) () {
  // load(const class QString &, const class QString &, const class QString &, const class QString &)
  // load(const class QLocale &, const class QString &, const class QString &, const class QString &, const class QString &)
  // load(const uchar *, int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QLocale{}) // "const QLocale &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][4] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const uchar *"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTranslator4loadERK7QStringS2_S2_S2_
    // invoke: bool load(const class QString &, const class QString &, const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QString).qclsinst
    if false {fmt.Println(arg3)}
    var ret = C.C_ZN11QTranslator4loadERK7QStringS2_S2_S2_(this.qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN11QTranslator4loadERK7QLocaleRK7QStringS5_S5_S5_
    // invoke: bool load(const class QLocale &, const class QString &, const class QString &, const class QString &, const class QString &)
    var arg0 = args[0].(QLocale).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QString).qclsinst
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QString).qclsinst
    if false {fmt.Println(arg4)}
    var ret = C.C_ZN11QTranslator4loadERK7QLocaleRK7QStringS5_S5_S5_(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret)}
  case 2:
    // invoke: _ZN11QTranslator4loadEPKhiRK7QString
    // invoke: bool load(const uchar *, int, const class QString &)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var ret = C.C_ZN11QTranslator4loadEPKhiRK7QString(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTranslator", "load", args)
  }

}

// metaObject()
func (this *QTranslator) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTranslator10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK11QTranslator10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTranslator", "metaObject", args)
  }

}

// ~QTranslator()
func (this *QTranslator) FreeQTranslator(args ...interface{}) () {
  // ~QTranslator()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTranslatorD0Ev
    // invoke: void ~QTranslator()
    C.C_ZN11QTranslatorD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTranslator", "~QTranslator", args)
  }

}

// isEmpty()
func (this *QTranslator) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTranslator7isEmptyEv
    // invoke: bool isEmpty()
    var ret = C.C_ZNK11QTranslator7isEmptyEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTranslator", "isEmpty", args)
  }

}

// QTranslator(class QObject *)
func NewQTranslator(args ...interface{}) *QTranslator {
  // QTranslator(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTranslatorC1EP7QObject
    // invoke: void QTranslator(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QTranslatorC2EP7QObject(arg0)
    return &QTranslator{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTranslator", "QTranslator", args)
  }

  return nil // QTranslator{qclsinst:qthis}
}

// translate(const char *, const char *, const char *, int)
func (this *QTranslator) translate(args ...interface{}) () {
  // translate(const char *, const char *, const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[0][2] = qtrt.ByteTy(true) // "const char *"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTranslator9translateEPKcS1_S1_i
    // invoke: QString translate(const char *, const char *, const char *, int)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[2].([]byte)).Pointer()))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var ret = C.C_ZNK11QTranslator9translateEPKcS1_S1_i(this.qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTranslator", "translate", args)
  }

}

// <= body block end

