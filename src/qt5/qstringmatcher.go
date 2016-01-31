package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtCore/qstringmatcher.h
// dst-file: /src/core/qstringmatcher.go
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
  // proto:  void QStringMatcher::QStringMatcher(const QStringMatcher & other);
extern void* C_ZN14QStringMatcherC2ERKS_(void* arg0); // 3
  // proto:  void QStringMatcher::QStringMatcher();
extern void* C_ZN14QStringMatcherC2Ev(); // 3
  // proto:  void QStringMatcher::setPattern(const QString & pattern);
extern void C_ZN14QStringMatcher10setPatternERK7QString(void* qthis, void* arg0); // 4
  // proto:  int QStringMatcher::indexIn(const QChar * str, int length, int from);
extern void C_ZNK14QStringMatcher7indexInEPK5QCharii(void* qthis, void* arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  int QStringMatcher::indexIn(const QString & str, int from);
extern void C_ZNK14QStringMatcher7indexInERK7QStringi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QString QStringMatcher::pattern();
extern void C_ZNK14QStringMatcher7patternEv(void* qthis); // 4
  // proto:  void QStringMatcher::~QStringMatcher();
extern void C_ZN14QStringMatcherD2Ev(void* qthis); // 4
  // proto:  Qt::CaseSensitivity QStringMatcher::caseSensitivity();
extern void C_ZNK14QStringMatcher15caseSensitivityEv(void* qthis); // 2
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

// class sizeof(QStringMatcher)=1048
type QStringMatcher struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// QStringMatcher(const class QStringMatcher &)
func NewQStringMatcher(args ...interface{}) *QStringMatcher {
  // QStringMatcher(const class QStringMatcher &)
  // QStringMatcher()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringMatcher{}) // "const QStringMatcher &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStringMatcherC1ERKS_
    // invoke: void QStringMatcher(const class QStringMatcher &)
    var arg0 = args[0].(QStringMatcher).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QStringMatcherC2ERKS_(arg0)
    return &QStringMatcher{qclsinst:qthis}
  case 1:
    // invoke: _ZN14QStringMatcherC1Ev
    // invoke: void QStringMatcher()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QStringMatcherC2Ev()
    return &QStringMatcher{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStringMatcher", "QStringMatcher", args)
  }

  return nil // QStringMatcher{qclsinst:qthis}
}

// setPattern(const class QString &)
func (this *QStringMatcher) setPattern(args ...interface{}) () {
  // setPattern(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStringMatcher10setPatternERK7QString
    // invoke: void setPattern(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QStringMatcher10setPatternERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStringMatcher", "setPattern", args)
  }

}

// indexIn(const class QChar *, int, int)
func (this *QStringMatcher) indexIn(args ...interface{}) () {
  // indexIn(const class QChar *, int, int)
  // indexIn(const class QString &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStringMatcher7indexInEPK5QCharii
    // invoke: int indexIn(const class QChar *, int, int)
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var ret = C.C_ZNK14QStringMatcher7indexInEPK5QCharii(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK14QStringMatcher7indexInERK7QStringi
    // invoke: int indexIn(const class QString &, int)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK14QStringMatcher7indexInERK7QStringi(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QStringMatcher", "indexIn", args)
  }

}

// pattern()
func (this *QStringMatcher) pattern(args ...interface{}) () {
  // pattern()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStringMatcher7patternEv
    // invoke: QString pattern()
    var ret = C.C_ZNK14QStringMatcher7patternEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QStringMatcher", "pattern", args)
  }

}

// ~QStringMatcher()
func (this *QStringMatcher) FreeQStringMatcher(args ...interface{}) () {
  // ~QStringMatcher()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStringMatcherD0Ev
    // invoke: void ~QStringMatcher()
    C.C_ZN14QStringMatcherD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringMatcher", "~QStringMatcher", args)
  }

}

// caseSensitivity()
func (this *QStringMatcher) caseSensitivity(args ...interface{}) () {
  // caseSensitivity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStringMatcher15caseSensitivityEv
    // invoke: Qt::CaseSensitivity caseSensitivity()
    C.C_ZNK14QStringMatcher15caseSensitivityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringMatcher", "caseSensitivity", args)
  }

}

// <= body block end

