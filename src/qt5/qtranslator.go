package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
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
  // proto:  const QMetaObject * QTranslator::metaObject();
extern void _ZNK11QTranslator10metaObjectEv(void* qthis);
  // proto:  void QTranslator::QTranslator(QObject * parent);
extern void* dector_ZN11QTranslatorC1EP7QObject(void* arg0);
extern void _ZN11QTranslatorC1EP7QObject(void* qthis, void* arg0);
  // proto:  bool QTranslator::isEmpty();
extern void _ZNK11QTranslator7isEmptyEv(void* qthis);
  // proto:  void QTranslator::QTranslator(const QTranslator & );
extern void* dector_ZN11QTranslatorC1ERKS_(void* arg0);
extern void _ZN11QTranslatorC1ERKS_(void* qthis, void* arg0);
  // proto:  void QTranslator::~QTranslator();
extern void _ZN11QTranslatorD0Ev(void* qthis);
  // proto:  bool QTranslator::load(const QString & filename, const QString & directory, const QString & search_delimiters, const QString & suffix);
extern void _ZN11QTranslator4loadERK7QStringS2_S2_S2_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3);
  // proto:  QString QTranslator::translate(const char * context, const char * sourceText, const char * disambiguation, int n);
extern void _ZNK11QTranslator9translateEPKcS1_S1_i(void* qthis, char* arg0, char* arg1, char* arg2, int arg3);
  // proto:  bool QTranslator::load(const uchar * data, int len, const QString & directory);
extern void _ZN11QTranslator4loadEPKhiRK7QString(void* qthis, unsigned char* arg0, int arg1, void* arg2);
  // proto:  bool QTranslator::load(const QLocale & locale, const QString & filename, const QString & prefix, const QString & directory, const QString & suffix);
extern void _ZN11QTranslator4loadERK7QLocaleRK7QStringS5_S5_S5_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3, void* arg4);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  const QMetaObject * QTranslator::metaObject();
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
  default:
    qtrt.ErrorResolve("QTranslator", "metaObject", args)
  }

}

  // proto:  void QTranslator::QTranslator(QObject * parent);
func NewQTranslator(args ...interface{}) QTranslator {
  return QTranslator{}
}

  // proto:  bool QTranslator::isEmpty();
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
  default:
    qtrt.ErrorResolve("QTranslator", "isEmpty", args)
  }

}

  // proto:  void QTranslator::~QTranslator();
func (this *QTranslator) FreeQTranslator(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTranslator", "~QTranslator", args)
  }

}

  // proto:  bool QTranslator::load(const QString & filename, const QString & directory, const QString & search_delimiters, const QString & suffix);
func (this *QTranslator) load(args ...interface{}) () {
  // load(const class QString &, const class QString &, const class QString &, const class QString &)
  // load(const uchar *, int, const class QString &)
  // load(const class QLocale &, const class QString &, const class QString &, const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const uchar *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QLocale{}) // "const QLocale &"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][4] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTranslator4loadERK7QStringS2_S2_S2_
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QString).qclsinst
    if false {fmt.Println(arg3)}
  case 1:
    // invoke: _ZN11QTranslator4loadEPKhiRK7QString
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
  case 2:
    // invoke: _ZN11QTranslator4loadERK7QLocaleRK7QStringS5_S5_S5_
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
  default:
    qtrt.ErrorResolve("QTranslator", "load", args)
  }

}

  // proto:  QString QTranslator::translate(const char * context, const char * sourceText, const char * disambiguation, int n);
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
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = C.CString(args[1].(string))
    if false {fmt.Println(arg1)}
    var arg2 = C.CString(args[2].(string))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QTranslator", "translate", args)
  }

}

// <= body block end

