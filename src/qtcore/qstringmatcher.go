package qtcore
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
import "runtime"
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
extern int32_t C_ZNK14QStringMatcher7indexInEPK5QCharii(void* qthis, void* arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  int QStringMatcher::indexIn(const QString & str, int from);
extern int32_t C_ZNK14QStringMatcher7indexInERK7QStringi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QString QStringMatcher::pattern();
extern void* C_ZNK14QStringMatcher7patternEv(void* qthis); // 4
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
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QStringMatcher)=1048
type QStringMatcher struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// QStringMatcher(const class QStringMatcher &)
func GcfreeQStringMatcher(this *QStringMatcher) {
  qtrt.UniverseFree(this)
}
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
    var arg0 = args[0].(*QStringMatcher).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QStringMatcherC2ERKS_(arg0)
    this := &QStringMatcher{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQStringMatcher)
    return this
  case 1:
    // invoke: _ZN14QStringMatcherC1Ev
    // invoke: void QStringMatcher()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QStringMatcherC2Ev()
    this := &QStringMatcher{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQStringMatcher)
    return this
  default:
    qtrt.ErrorResolve("QStringMatcher", "QStringMatcher", args)
  }

  return nil // QStringMatcher{Qclsinst:qthis}
}

// setPattern(const class QString &)
func (this *QStringMatcher) SetPattern(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QStringMatcher10setPatternERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStringMatcher", "setPattern", args)
  }

  return
}

// indexIn(const class QChar *, int, int)
func (this *QStringMatcher) IndexIn(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK14QStringMatcher7indexInEPK5QCharii(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK14QStringMatcher7indexInERK7QStringi
    // invoke: int indexIn(const class QString &, int)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK14QStringMatcher7indexInERK7QStringi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringMatcher", "indexIn", args)
  }

  return
}

// pattern()
func (this *QStringMatcher) Pattern(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QStringMatcher7patternEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringMatcher", "pattern", args)
  }

  return
}

// ~QStringMatcher()
func (this *QStringMatcher) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN14QStringMatcherD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QStringMatcher", "~QStringMatcher", args)
  }

  return
}

// caseSensitivity()
func (this *QStringMatcher) CaseSensitivity(args ...interface{}) () {
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
    C.C_ZNK14QStringMatcher15caseSensitivityEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStringMatcher", "caseSensitivity", args)
  }

  return
}

// <= body block end

