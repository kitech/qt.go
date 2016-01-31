package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtCore/qregexp.h
// dst-file: /src/core/qregexp.go
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
  // proto:  int QRegExp::pos(int nth);
extern void C_ZN7QRegExp3posEi(void* qthis, int32_t arg0); // 4
  // proto:  int QRegExp::matchedLength();
extern void C_ZNK7QRegExp13matchedLengthEv(void* qthis); // 4
  // proto:  void QRegExp::setMinimal(bool minimal);
extern void C_ZN7QRegExp10setMinimalEb(void* qthis, bool arg0); // 4
  // proto:  void QRegExp::QRegExp();
extern void* C_ZN7QRegExpC2Ev(); // 3
  // proto:  void QRegExp::QRegExp(const QRegExp & rx);
extern void* C_ZN7QRegExpC2ERKS_(void* arg0); // 3
  // proto: static QString QRegExp::escape(const QString & str);
extern void C_ZN7QRegExp6escapeERK7QString(void* arg0); // 4
  // proto:  void QRegExp::setPattern(const QString & pattern);
extern void C_ZN7QRegExp10setPatternERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QRegExp::pattern();
extern void C_ZNK7QRegExp7patternEv(void* qthis); // 4
  // proto:  QRegExp::PatternSyntax QRegExp::patternSyntax();
extern void C_ZNK7QRegExp13patternSyntaxEv(void* qthis); // 4
  // proto:  bool QRegExp::isEmpty();
extern void C_ZNK7QRegExp7isEmptyEv(void* qthis); // 4
  // proto:  void QRegExp::swap(QRegExp & other);
extern void C_ZN7QRegExp4swapERS_(void* qthis, void* arg0); // 2
  // proto:  Qt::CaseSensitivity QRegExp::caseSensitivity();
extern void C_ZNK7QRegExp15caseSensitivityEv(void* qthis); // 4
  // proto:  QString QRegExp::errorString();
extern void C_ZN7QRegExp11errorStringEv(void* qthis); // 4
  // proto:  QStringList QRegExp::capturedTexts();
extern void C_ZN7QRegExp13capturedTextsEv(void* qthis); // 4
  // proto:  bool QRegExp::isValid();
extern void C_ZNK7QRegExp7isValidEv(void* qthis); // 4
  // proto:  bool QRegExp::exactMatch(const QString & str);
extern void C_ZNK7QRegExp10exactMatchERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QRegExp::cap(int nth);
extern void C_ZN7QRegExp3capEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QRegExp::isMinimal();
extern void C_ZNK7QRegExp9isMinimalEv(void* qthis); // 4
  // proto:  void QRegExp::~QRegExp();
extern void C_ZN7QRegExpD2Ev(void* qthis); // 4
  // proto:  int QRegExp::captureCount();
extern void C_ZNK7QRegExp12captureCountEv(void* qthis); // 4
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

// class sizeof(QRegExp)=8
type QRegExp struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// pos(int)
func (this *QRegExp) pos(args ...interface{}) () {
  // pos(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegExp3posEi
    // invoke: int pos(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN7QRegExp3posEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegExp", "pos", args)
  }

}

// matchedLength()
func (this *QRegExp) matchedLength(args ...interface{}) () {
  // matchedLength()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegExp13matchedLengthEv
    // invoke: int matchedLength()
    var ret = C.C_ZNK7QRegExp13matchedLengthEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegExp", "matchedLength", args)
  }

}

// setMinimal(_Bool)
func (this *QRegExp) setMinimal(args ...interface{}) () {
  // setMinimal(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegExp10setMinimalEb
    // invoke: void setMinimal(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QRegExp10setMinimalEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegExp", "setMinimal", args)
  }

}

// QRegExp()
func NewQRegExp(args ...interface{}) *QRegExp {
  // QRegExp()
  // QRegExp(const class QRegExp &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegExpC1Ev
    // invoke: void QRegExp()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QRegExpC2Ev()
    return &QRegExp{qclsinst:qthis}
  case 1:
    // invoke: _ZN7QRegExpC1ERKS_
    // invoke: void QRegExp(const class QRegExp &)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QRegExpC2ERKS_(arg0)
    return &QRegExp{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QRegExp", "QRegExp", args)
  }

  return nil // QRegExp{qclsinst:qthis}
}

// escape(const class QString &)
func (this *QRegExp) escape_s(args ...interface{}) () {
  // escape(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegExp6escapeERK7QString
    // invoke: QString escape(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN7QRegExp6escapeERK7QString(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegExp", "escape", args)
  }

}

// setPattern(const class QString &)
func (this *QRegExp) setPattern(args ...interface{}) () {
  // setPattern(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegExp10setPatternERK7QString
    // invoke: void setPattern(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QRegExp10setPatternERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegExp", "setPattern", args)
  }

}

// pattern()
func (this *QRegExp) pattern(args ...interface{}) () {
  // pattern()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegExp7patternEv
    // invoke: QString pattern()
    var ret = C.C_ZNK7QRegExp7patternEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegExp", "pattern", args)
  }

}

// patternSyntax()
func (this *QRegExp) patternSyntax(args ...interface{}) () {
  // patternSyntax()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegExp13patternSyntaxEv
    // invoke: QRegExp::PatternSyntax patternSyntax()
    C.C_ZNK7QRegExp13patternSyntaxEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegExp", "patternSyntax", args)
  }

}

// isEmpty()
func (this *QRegExp) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegExp7isEmptyEv
    // invoke: bool isEmpty()
    var ret = C.C_ZNK7QRegExp7isEmptyEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegExp", "isEmpty", args)
  }

}

// swap(class QRegExp &)
func (this *QRegExp) swap(args ...interface{}) () {
  // swap(class QRegExp &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegExp{}) // "QRegExp &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegExp4swapERS_
    // invoke: void swap(class QRegExp &)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QRegExp4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegExp", "swap", args)
  }

}

// caseSensitivity()
func (this *QRegExp) caseSensitivity(args ...interface{}) () {
  // caseSensitivity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegExp15caseSensitivityEv
    // invoke: Qt::CaseSensitivity caseSensitivity()
    C.C_ZNK7QRegExp15caseSensitivityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegExp", "caseSensitivity", args)
  }

}

// errorString()
func (this *QRegExp) errorString(args ...interface{}) () {
  // errorString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegExp11errorStringEv
    // invoke: QString errorString()
    var ret = C.C_ZN7QRegExp11errorStringEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegExp", "errorString", args)
  }

}

// capturedTexts()
func (this *QRegExp) capturedTexts(args ...interface{}) () {
  // capturedTexts()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegExp13capturedTextsEv
    // invoke: QStringList capturedTexts()
    C.C_ZN7QRegExp13capturedTextsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegExp", "capturedTexts", args)
  }

}

// isValid()
func (this *QRegExp) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegExp7isValidEv
    // invoke: bool isValid()
    var ret = C.C_ZNK7QRegExp7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegExp", "isValid", args)
  }

}

// exactMatch(const class QString &)
func (this *QRegExp) exactMatch(args ...interface{}) () {
  // exactMatch(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegExp10exactMatchERK7QString
    // invoke: bool exactMatch(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK7QRegExp10exactMatchERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegExp", "exactMatch", args)
  }

}

// cap(int)
func (this *QRegExp) cap(args ...interface{}) () {
  // cap(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegExp3capEi
    // invoke: QString cap(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN7QRegExp3capEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegExp", "cap", args)
  }

}

// isMinimal()
func (this *QRegExp) isMinimal(args ...interface{}) () {
  // isMinimal()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegExp9isMinimalEv
    // invoke: bool isMinimal()
    var ret = C.C_ZNK7QRegExp9isMinimalEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegExp", "isMinimal", args)
  }

}

// ~QRegExp()
func (this *QRegExp) FreeQRegExp(args ...interface{}) () {
  // ~QRegExp()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegExpD0Ev
    // invoke: void ~QRegExp()
    C.C_ZN7QRegExpD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegExp", "~QRegExp", args)
  }

}

// captureCount()
func (this *QRegExp) captureCount(args ...interface{}) () {
  // captureCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegExp12captureCountEv
    // invoke: int captureCount()
    var ret = C.C_ZNK7QRegExp12captureCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegExp", "captureCount", args)
  }

}

// <= body block end

