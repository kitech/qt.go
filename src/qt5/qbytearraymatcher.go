package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtCore/qbytearraymatcher.h
// dst-file: /src/core/qbytearraymatcher.go
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
  // proto:  QByteArray QByteArrayMatcher::pattern();
extern void demth_ZNK17QByteArrayMatcher7patternEv(void* qthis);
  // proto:  int QByteArrayMatcher::indexIn(const char * str, int len, int from);
extern void _ZNK17QByteArrayMatcher7indexInEPKcii(void* qthis, char* arg0, int arg1, int arg2);
  // proto:  void QByteArrayMatcher::setPattern(const QByteArray & pattern);
extern void _ZN17QByteArrayMatcher10setPatternERK10QByteArray(void* qthis, void* arg0);
  // proto:  void QByteArrayMatcher::QByteArrayMatcher();
extern void* dector_ZN17QByteArrayMatcherC1Ev();
extern void _ZN17QByteArrayMatcherC1Ev(void* qthis);
  // proto:  void QByteArrayMatcher::QByteArrayMatcher(const char * pattern, int length);
extern void* dector_ZN17QByteArrayMatcherC1EPKci(char* arg0, int arg1);
extern void _ZN17QByteArrayMatcherC1EPKci(void* qthis, char* arg0, int arg1);
  // proto:  void QByteArrayMatcher::QByteArrayMatcher(const QByteArray & pattern);
extern void* dector_ZN17QByteArrayMatcherC1ERK10QByteArray(void* arg0);
extern void _ZN17QByteArrayMatcherC1ERK10QByteArray(void* qthis, void* arg0);
  // proto:  int QByteArrayMatcher::indexIn(const QByteArray & ba, int from);
extern void _ZNK17QByteArrayMatcher7indexInERK10QByteArrayi(void* qthis, void* arg0, int arg1);
  // proto:  void QByteArrayMatcher::QByteArrayMatcher(const QByteArrayMatcher & other);
extern void* dector_ZN17QByteArrayMatcherC1ERKS_(void* arg0);
extern void _ZN17QByteArrayMatcherC1ERKS_(void* qthis, void* arg0);
  // proto:  void QByteArrayMatcher::~QByteArrayMatcher();
extern void _ZN17QByteArrayMatcherD0Ev(void* qthis);
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

// class sizeof(QByteArrayMatcher)=1040
type QByteArrayMatcher struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QByteArray QByteArrayMatcher::pattern();
func (this *QByteArrayMatcher) pattern(args ...interface{}) () {
  // pattern()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QByteArrayMatcher7patternEv
    // invoke: QByteArray pattern()
    C.demth_ZNK17QByteArrayMatcher7patternEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArrayMatcher", "pattern", args)
  }

}

  // proto:  int QByteArrayMatcher::indexIn(const char * str, int len, int from);
func (this *QByteArrayMatcher) indexIn(args ...interface{}) () {
  // indexIn(const char *, int, int)
  // indexIn(const class QByteArray &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QByteArrayMatcher7indexInEPKcii
    // invoke: int indexIn(const char *, int, int)
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZNK17QByteArrayMatcher7indexInEPKcii(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZNK17QByteArrayMatcher7indexInERK10QByteArrayi
    // invoke: int indexIn(const class QByteArray &, int)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK17QByteArrayMatcher7indexInERK10QByteArrayi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QByteArrayMatcher", "indexIn", args)
  }

}

  // proto:  void QByteArrayMatcher::setPattern(const QByteArray & pattern);
func (this *QByteArrayMatcher) setPattern(args ...interface{}) () {
  // setPattern(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QByteArrayMatcher10setPatternERK10QByteArray
    // invoke: void setPattern(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN17QByteArrayMatcher10setPatternERK10QByteArray(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArrayMatcher", "setPattern", args)
  }

}

  // proto:  void QByteArrayMatcher::QByteArrayMatcher();
func NewQByteArrayMatcher(args ...interface{}) QByteArrayMatcher {
  return QByteArrayMatcher{}
}

  // proto:  void QByteArrayMatcher::~QByteArrayMatcher();
func (this *QByteArrayMatcher) FreeQByteArrayMatcher(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QByteArrayMatcher", "~QByteArrayMatcher", args)
  }

}

// <= body block end

