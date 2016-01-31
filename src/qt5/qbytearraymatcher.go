package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  int QByteArrayMatcher::indexIn(const char * str, int len, int from);
extern int32_t C_ZNK17QByteArrayMatcher7indexInEPKcii(void* qthis, void* arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  int QByteArrayMatcher::indexIn(const QByteArray & ba, int from);
extern int32_t C_ZNK17QByteArrayMatcher7indexInERK10QByteArrayi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QByteArray QByteArrayMatcher::pattern();
extern void* C_ZNK17QByteArrayMatcher7patternEv(void* qthis); // 2
  // proto:  void QByteArrayMatcher::QByteArrayMatcher(const QByteArrayMatcher & other);
extern void* C_ZN17QByteArrayMatcherC2ERKS_(void* arg0); // 3
  // proto:  void QByteArrayMatcher::QByteArrayMatcher();
extern void* C_ZN17QByteArrayMatcherC2Ev(); // 3
  // proto:  void QByteArrayMatcher::QByteArrayMatcher(const char * pattern, int length);
extern void* C_ZN17QByteArrayMatcherC2EPKci(void* arg0, int32_t arg1); // 3
  // proto:  void QByteArrayMatcher::QByteArrayMatcher(const QByteArray & pattern);
extern void* C_ZN17QByteArrayMatcherC2ERK10QByteArray(void* arg0); // 3
  // proto:  void QByteArrayMatcher::setPattern(const QByteArray & pattern);
extern void C_ZN17QByteArrayMatcher10setPatternERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  void QByteArrayMatcher::~QByteArrayMatcher();
extern void C_ZN17QByteArrayMatcherD2Ev(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// indexIn(const char *, int, int)
func (this *QByteArrayMatcher) Indexin(args ...interface{}) (ret interface{}) {
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
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK17QByteArrayMatcher7indexInEPKcii(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZNK17QByteArrayMatcher7indexInERK10QByteArrayi
    // invoke: int indexIn(const class QByteArray &, int)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK17QByteArrayMatcher7indexInERK10QByteArrayi(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QByteArrayMatcher", "indexIn", args)
  }

  return
}

// pattern()
func (this *QByteArrayMatcher) Pattern(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK17QByteArrayMatcher7patternEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QByteArrayMatcher", "pattern", args)
  }

  return
}

// QByteArrayMatcher(const class QByteArrayMatcher &)
func NewQByteArrayMatcher(args ...interface{}) *QByteArrayMatcher {
  // QByteArrayMatcher(const class QByteArrayMatcher &)
  // QByteArrayMatcher()
  // QByteArrayMatcher(const char *, int)
  // QByteArrayMatcher(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArrayMatcher{}) // "const QByteArrayMatcher &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QByteArrayMatcherC1ERKS_
    // invoke: void QByteArrayMatcher(const class QByteArrayMatcher &)
    var arg0 = args[0].(QByteArrayMatcher).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QByteArrayMatcherC2ERKS_(arg0)
    return &QByteArrayMatcher{qclsinst:qthis}
  case 1:
    // invoke: _ZN17QByteArrayMatcherC1Ev
    // invoke: void QByteArrayMatcher()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QByteArrayMatcherC2Ev()
    return &QByteArrayMatcher{qclsinst:qthis}
  case 2:
    // invoke: _ZN17QByteArrayMatcherC1EPKci
    // invoke: void QByteArrayMatcher(const char *, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[2][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QByteArrayMatcherC2EPKci(arg0, arg1)
    return &QByteArrayMatcher{qclsinst:qthis}
  case 3:
    // invoke: _ZN17QByteArrayMatcherC1ERK10QByteArray
    // invoke: void QByteArrayMatcher(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QByteArrayMatcherC2ERK10QByteArray(arg0)
    return &QByteArrayMatcher{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QByteArrayMatcher", "QByteArrayMatcher", args)
  }

  return nil // QByteArrayMatcher{qclsinst:qthis}
}

// setPattern(const class QByteArray &)
func (this *QByteArrayMatcher) Setpattern(args ...interface{}) () {
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
    C.C_ZN17QByteArrayMatcher10setPatternERK10QByteArray(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArrayMatcher", "setPattern", args)
  }

  return
}

// ~QByteArrayMatcher()
func (this *QByteArrayMatcher) Freeqbytearraymatcher(args ...interface{}) () {
  // ~QByteArrayMatcher()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QByteArrayMatcherD0Ev
    // invoke: void ~QByteArrayMatcher()
    C.C_ZN17QByteArrayMatcherD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArrayMatcher", "~QByteArrayMatcher", args)
  }

  return
}

// <= body block end

