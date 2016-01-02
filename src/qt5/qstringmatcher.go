package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
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
  // proto:  void QStringMatcher::QStringMatcher();
extern void* dector_ZN14QStringMatcherC1Ev();
extern void _ZN14QStringMatcherC1Ev(void* qthis);
  // proto:  void QStringMatcher::QStringMatcher(const QStringMatcher & other);
extern void* dector_ZN14QStringMatcherC1ERKS_(void* arg0);
extern void _ZN14QStringMatcherC1ERKS_(void* qthis, void* arg0);
  // proto:  int QStringMatcher::indexIn(const QChar * str, int length, int from);
extern void _ZNK14QStringMatcher7indexInEPK5QCharii(void* qthis, void* arg0, int arg1, int arg2);
  // proto:  void QStringMatcher::setPattern(const QString & pattern);
extern void _ZN14QStringMatcher10setPatternERK7QString(void* qthis, void* arg0);
  // proto:  QString QStringMatcher::pattern();
extern void _ZNK14QStringMatcher7patternEv(void* qthis);
  // proto:  void QStringMatcher::~QStringMatcher();
extern void _ZN14QStringMatcherD0Ev(void* qthis);
  // proto:  int QStringMatcher::indexIn(const QString & str, int from);
extern void _ZNK14QStringMatcher7indexInERK7QStringi(void* qthis, void* arg0, int arg1);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QStringMatcher::QStringMatcher();
func NewQStringMatcher(args ...interface{}) QStringMatcher {
  return QStringMatcher{}
}

  // proto:  int QStringMatcher::indexIn(const QChar * str, int length, int from);
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
    C._ZNK14QStringMatcher7indexInEPK5QCharii(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZNK14QStringMatcher7indexInERK7QStringi
    // invoke: int indexIn(const class QString &, int)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK14QStringMatcher7indexInERK7QStringi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStringMatcher", "indexIn", args)
  }

}

  // proto:  void QStringMatcher::setPattern(const QString & pattern);
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
    C._ZN14QStringMatcher10setPatternERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStringMatcher", "setPattern", args)
  }

}

  // proto:  QString QStringMatcher::pattern();
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
    C._ZNK14QStringMatcher7patternEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringMatcher", "pattern", args)
  }

}

  // proto:  void QStringMatcher::~QStringMatcher();
func (this *QStringMatcher) FreeQStringMatcher(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStringMatcher", "~QStringMatcher", args)
  }

}

// <= body block end

