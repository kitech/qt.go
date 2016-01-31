package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtCore/qregularexpression.h
// dst-file: /src/core/qregularexpression.go
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
  // proto:  QRegularExpressionMatch QRegularExpressionMatchIterator::peekNext();
extern void C_ZNK31QRegularExpressionMatchIterator8peekNextEv(void* qthis); // 4
  // proto:  bool QRegularExpressionMatchIterator::hasNext();
extern void C_ZNK31QRegularExpressionMatchIterator7hasNextEv(void* qthis); // 4
  // proto:  bool QRegularExpressionMatchIterator::isValid();
extern void C_ZNK31QRegularExpressionMatchIterator7isValidEv(void* qthis); // 4
  // proto:  QRegularExpression QRegularExpressionMatchIterator::regularExpression();
extern void C_ZNK31QRegularExpressionMatchIterator17regularExpressionEv(void* qthis); // 4
  // proto:  QRegularExpressionMatch QRegularExpressionMatchIterator::next();
extern void C_ZN31QRegularExpressionMatchIterator4nextEv(void* qthis); // 4
  // proto:  void QRegularExpressionMatchIterator::swap(QRegularExpressionMatchIterator & other);
extern void C_ZN31QRegularExpressionMatchIterator4swapERS_(void* qthis, void* arg0); // 2
  // proto:  QRegularExpression::MatchType QRegularExpressionMatchIterator::matchType();
extern void C_ZNK31QRegularExpressionMatchIterator9matchTypeEv(void* qthis); // 4
  // proto:  QRegularExpression::MatchOptions QRegularExpressionMatchIterator::matchOptions();
extern void C_ZNK31QRegularExpressionMatchIterator12matchOptionsEv(void* qthis); // 4
  // proto:  void QRegularExpressionMatchIterator::~QRegularExpressionMatchIterator();
extern void C_ZN31QRegularExpressionMatchIteratorD2Ev(void* qthis); // 4
  // proto:  void QRegularExpressionMatchIterator::QRegularExpressionMatchIterator(const QRegularExpressionMatchIterator & iterator);
extern void* C_ZN31QRegularExpressionMatchIteratorC2ERKS_(void* arg0); // 3
  // proto:  void QRegularExpressionMatchIterator::QRegularExpressionMatchIterator();
extern void* C_ZN31QRegularExpressionMatchIteratorC2Ev(); // 3
  // proto:  bool QRegularExpression::isValid();
extern void C_ZNK18QRegularExpression7isValidEv(void* qthis); // 4
  // proto:  QString QRegularExpression::errorString();
extern void C_ZNK18QRegularExpression11errorStringEv(void* qthis); // 4
  // proto:  QString QRegularExpression::pattern();
extern void C_ZNK18QRegularExpression7patternEv(void* qthis); // 4
  // proto:  void QRegularExpression::QRegularExpression(const QRegularExpression & re);
extern void* C_ZN18QRegularExpressionC2ERKS_(void* arg0); // 3
  // proto:  void QRegularExpression::QRegularExpression();
extern void* C_ZN18QRegularExpressionC2Ev(); // 3
  // proto:  void QRegularExpression::~QRegularExpression();
extern void C_ZN18QRegularExpressionD2Ev(void* qthis); // 4
  // proto:  void QRegularExpression::setPattern(const QString & pattern);
extern void C_ZN18QRegularExpression10setPatternERK7QString(void* qthis, void* arg0); // 4
  // proto:  int QRegularExpression::patternErrorOffset();
extern void C_ZNK18QRegularExpression18patternErrorOffsetEv(void* qthis); // 4
  // proto:  PatternOptions QRegularExpression::patternOptions();
extern void C_ZNK18QRegularExpression14patternOptionsEv(void* qthis); // 4
  // proto:  void QRegularExpression::swap(QRegularExpression & other);
extern void C_ZN18QRegularExpression4swapERS_(void* qthis, void* arg0); // 2
  // proto: static QString QRegularExpression::escape(const QString & str);
extern void C_ZN18QRegularExpression6escapeERK7QString(void* arg0); // 4
  // proto:  QStringList QRegularExpression::namedCaptureGroups();
extern void C_ZNK18QRegularExpression18namedCaptureGroupsEv(void* qthis); // 4
  // proto:  void QRegularExpression::optimize();
extern void C_ZNK18QRegularExpression8optimizeEv(void* qthis); // 4
  // proto:  int QRegularExpression::captureCount();
extern void C_ZNK18QRegularExpression12captureCountEv(void* qthis); // 4
  // proto:  bool QRegularExpressionMatch::hasPartialMatch();
extern void C_ZNK23QRegularExpressionMatch15hasPartialMatchEv(void* qthis); // 4
  // proto:  QRegularExpression QRegularExpressionMatch::regularExpression();
extern void C_ZNK23QRegularExpressionMatch17regularExpressionEv(void* qthis); // 4
  // proto:  void QRegularExpressionMatch::QRegularExpressionMatch(const QRegularExpressionMatch & match);
extern void* C_ZN23QRegularExpressionMatchC2ERKS_(void* arg0); // 3
  // proto:  void QRegularExpressionMatch::QRegularExpressionMatch();
extern void* C_ZN23QRegularExpressionMatchC2Ev(); // 3
  // proto:  QStringList QRegularExpressionMatch::capturedTexts();
extern void C_ZNK23QRegularExpressionMatch13capturedTextsEv(void* qthis); // 4
  // proto:  bool QRegularExpressionMatch::isValid();
extern void C_ZNK23QRegularExpressionMatch7isValidEv(void* qthis); // 4
  // proto:  bool QRegularExpressionMatch::hasMatch();
extern void C_ZNK23QRegularExpressionMatch8hasMatchEv(void* qthis); // 4
  // proto:  void QRegularExpressionMatch::~QRegularExpressionMatch();
extern void C_ZN23QRegularExpressionMatchD2Ev(void* qthis); // 4
  // proto:  QRegularExpression::MatchOptions QRegularExpressionMatch::matchOptions();
extern void C_ZNK23QRegularExpressionMatch12matchOptionsEv(void* qthis); // 4
  // proto:  int QRegularExpressionMatch::capturedStart(const QString & name);
extern void C_ZNK23QRegularExpressionMatch13capturedStartERK7QString(void* qthis, void* arg0); // 4
  // proto:  int QRegularExpressionMatch::capturedStart(int nth);
extern void C_ZNK23QRegularExpressionMatch13capturedStartEi(void* qthis, int32_t arg0); // 4
  // proto:  QStringRef QRegularExpressionMatch::capturedRef(int nth);
extern void C_ZNK23QRegularExpressionMatch11capturedRefEi(void* qthis, int32_t arg0); // 4
  // proto:  QStringRef QRegularExpressionMatch::capturedRef(const QString & name);
extern void C_ZNK23QRegularExpressionMatch11capturedRefERK7QString(void* qthis, void* arg0); // 4
  // proto:  int QRegularExpressionMatch::lastCapturedIndex();
extern void C_ZNK23QRegularExpressionMatch17lastCapturedIndexEv(void* qthis); // 4
  // proto:  QString QRegularExpressionMatch::captured(int nth);
extern void C_ZNK23QRegularExpressionMatch8capturedEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QRegularExpressionMatch::captured(const QString & name);
extern void C_ZNK23QRegularExpressionMatch8capturedERK7QString(void* qthis, void* arg0); // 4
  // proto:  QRegularExpression::MatchType QRegularExpressionMatch::matchType();
extern void C_ZNK23QRegularExpressionMatch9matchTypeEv(void* qthis); // 4
  // proto:  int QRegularExpressionMatch::capturedLength(int nth);
extern void C_ZNK23QRegularExpressionMatch14capturedLengthEi(void* qthis, int32_t arg0); // 4
  // proto:  int QRegularExpressionMatch::capturedLength(const QString & name);
extern void C_ZNK23QRegularExpressionMatch14capturedLengthERK7QString(void* qthis, void* arg0); // 4
  // proto:  int QRegularExpressionMatch::capturedEnd(const QString & name);
extern void C_ZNK23QRegularExpressionMatch11capturedEndERK7QString(void* qthis, void* arg0); // 4
  // proto:  int QRegularExpressionMatch::capturedEnd(int nth);
extern void C_ZNK23QRegularExpressionMatch11capturedEndEi(void* qthis, int32_t arg0); // 4
  // proto:  void QRegularExpressionMatch::swap(QRegularExpressionMatch & other);
extern void C_ZN23QRegularExpressionMatch4swapERS_(void* qthis, void* arg0); // 2
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

// class sizeof(QRegularExpressionMatchIterator)=1
type QRegularExpressionMatchIterator struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QRegularExpression)=1
type QRegularExpression struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QRegularExpressionMatch)=1
type QRegularExpressionMatch struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// peekNext()
func (this *QRegularExpressionMatchIterator) peekNext(args ...interface{}) () {
  // peekNext()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK31QRegularExpressionMatchIterator8peekNextEv
    // invoke: QRegularExpressionMatch peekNext()
    var ret = C.C_ZNK31QRegularExpressionMatchIterator8peekNextEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegularExpressionMatchIterator", "peekNext", args)
  }

}

// hasNext()
func (this *QRegularExpressionMatchIterator) hasNext(args ...interface{}) () {
  // hasNext()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK31QRegularExpressionMatchIterator7hasNextEv
    // invoke: bool hasNext()
    var ret = C.C_ZNK31QRegularExpressionMatchIterator7hasNextEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegularExpressionMatchIterator", "hasNext", args)
  }

}

// isValid()
func (this *QRegularExpressionMatchIterator) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK31QRegularExpressionMatchIterator7isValidEv
    // invoke: bool isValid()
    var ret = C.C_ZNK31QRegularExpressionMatchIterator7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegularExpressionMatchIterator", "isValid", args)
  }

}

// regularExpression()
func (this *QRegularExpressionMatchIterator) regularExpression(args ...interface{}) () {
  // regularExpression()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK31QRegularExpressionMatchIterator17regularExpressionEv
    // invoke: QRegularExpression regularExpression()
    var ret = C.C_ZNK31QRegularExpressionMatchIterator17regularExpressionEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegularExpressionMatchIterator", "regularExpression", args)
  }

}

// next()
func (this *QRegularExpressionMatchIterator) next(args ...interface{}) () {
  // next()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN31QRegularExpressionMatchIterator4nextEv
    // invoke: QRegularExpressionMatch next()
    var ret = C.C_ZN31QRegularExpressionMatchIterator4nextEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegularExpressionMatchIterator", "next", args)
  }

}

// swap(class QRegularExpressionMatchIterator &)
func (this *QRegularExpressionMatchIterator) swap(args ...interface{}) () {
  // swap(class QRegularExpressionMatchIterator &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegularExpressionMatchIterator{}) // "QRegularExpressionMatchIterator &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN31QRegularExpressionMatchIterator4swapERS_
    // invoke: void swap(class QRegularExpressionMatchIterator &)
    var arg0 = args[0].(QRegularExpressionMatchIterator).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN31QRegularExpressionMatchIterator4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatchIterator", "swap", args)
  }

}

// matchType()
func (this *QRegularExpressionMatchIterator) matchType(args ...interface{}) () {
  // matchType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK31QRegularExpressionMatchIterator9matchTypeEv
    // invoke: QRegularExpression::MatchType matchType()
    C.C_ZNK31QRegularExpressionMatchIterator9matchTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatchIterator", "matchType", args)
  }

}

// matchOptions()
func (this *QRegularExpressionMatchIterator) matchOptions(args ...interface{}) () {
  // matchOptions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK31QRegularExpressionMatchIterator12matchOptionsEv
    // invoke: QRegularExpression::MatchOptions matchOptions()
    C.C_ZNK31QRegularExpressionMatchIterator12matchOptionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatchIterator", "matchOptions", args)
  }

}

// ~QRegularExpressionMatchIterator()
func (this *QRegularExpressionMatchIterator) FreeQRegularExpressionMatchIterator(args ...interface{}) () {
  // ~QRegularExpressionMatchIterator()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN31QRegularExpressionMatchIteratorD0Ev
    // invoke: void ~QRegularExpressionMatchIterator()
    C.C_ZN31QRegularExpressionMatchIteratorD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatchIterator", "~QRegularExpressionMatchIterator", args)
  }

}

// QRegularExpressionMatchIterator(const class QRegularExpressionMatchIterator &)
func NewQRegularExpressionMatchIterator(args ...interface{}) *QRegularExpressionMatchIterator {
  // QRegularExpressionMatchIterator(const class QRegularExpressionMatchIterator &)
  // QRegularExpressionMatchIterator()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegularExpressionMatchIterator{}) // "const QRegularExpressionMatchIterator &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN31QRegularExpressionMatchIteratorC1ERKS_
    // invoke: void QRegularExpressionMatchIterator(const class QRegularExpressionMatchIterator &)
    var arg0 = args[0].(QRegularExpressionMatchIterator).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN31QRegularExpressionMatchIteratorC2ERKS_(arg0)
    return &QRegularExpressionMatchIterator{qclsinst:qthis}
  case 1:
    // invoke: _ZN31QRegularExpressionMatchIteratorC1Ev
    // invoke: void QRegularExpressionMatchIterator()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN31QRegularExpressionMatchIteratorC2Ev()
    return &QRegularExpressionMatchIterator{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QRegularExpressionMatchIterator", "QRegularExpressionMatchIterator", args)
  }

  return nil // QRegularExpressionMatchIterator{qclsinst:qthis}
}

// isValid()
func (this *QRegularExpression) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QRegularExpression7isValidEv
    // invoke: bool isValid()
    var ret = C.C_ZNK18QRegularExpression7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegularExpression", "isValid", args)
  }

}

// errorString()
func (this *QRegularExpression) errorString(args ...interface{}) () {
  // errorString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QRegularExpression11errorStringEv
    // invoke: QString errorString()
    var ret = C.C_ZNK18QRegularExpression11errorStringEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegularExpression", "errorString", args)
  }

}

// pattern()
func (this *QRegularExpression) pattern(args ...interface{}) () {
  // pattern()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QRegularExpression7patternEv
    // invoke: QString pattern()
    var ret = C.C_ZNK18QRegularExpression7patternEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegularExpression", "pattern", args)
  }

}

// QRegularExpression(const class QRegularExpression &)
func NewQRegularExpression(args ...interface{}) *QRegularExpression {
  // QRegularExpression(const class QRegularExpression &)
  // QRegularExpression()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QRegularExpressionC1ERKS_
    // invoke: void QRegularExpression(const class QRegularExpression &)
    var arg0 = args[0].(QRegularExpression).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QRegularExpressionC2ERKS_(arg0)
    return &QRegularExpression{qclsinst:qthis}
  case 1:
    // invoke: _ZN18QRegularExpressionC1Ev
    // invoke: void QRegularExpression()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QRegularExpressionC2Ev()
    return &QRegularExpression{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QRegularExpression", "QRegularExpression", args)
  }

  return nil // QRegularExpression{qclsinst:qthis}
}

// ~QRegularExpression()
func (this *QRegularExpression) FreeQRegularExpression(args ...interface{}) () {
  // ~QRegularExpression()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QRegularExpressionD0Ev
    // invoke: void ~QRegularExpression()
    C.C_ZN18QRegularExpressionD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpression", "~QRegularExpression", args)
  }

}

// setPattern(const class QString &)
func (this *QRegularExpression) setPattern(args ...interface{}) () {
  // setPattern(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QRegularExpression10setPatternERK7QString
    // invoke: void setPattern(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QRegularExpression10setPatternERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegularExpression", "setPattern", args)
  }

}

// patternErrorOffset()
func (this *QRegularExpression) patternErrorOffset(args ...interface{}) () {
  // patternErrorOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QRegularExpression18patternErrorOffsetEv
    // invoke: int patternErrorOffset()
    var ret = C.C_ZNK18QRegularExpression18patternErrorOffsetEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegularExpression", "patternErrorOffset", args)
  }

}

// patternOptions()
func (this *QRegularExpression) patternOptions(args ...interface{}) () {
  // patternOptions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QRegularExpression14patternOptionsEv
    // invoke: PatternOptions patternOptions()
    C.C_ZNK18QRegularExpression14patternOptionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpression", "patternOptions", args)
  }

}

// swap(class QRegularExpression &)
func (this *QRegularExpression) swap(args ...interface{}) () {
  // swap(class QRegularExpression &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegularExpression{}) // "QRegularExpression &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QRegularExpression4swapERS_
    // invoke: void swap(class QRegularExpression &)
    var arg0 = args[0].(QRegularExpression).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QRegularExpression4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegularExpression", "swap", args)
  }

}

// escape(const class QString &)
func (this *QRegularExpression) escape_s(args ...interface{}) () {
  // escape(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QRegularExpression6escapeERK7QString
    // invoke: QString escape(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN18QRegularExpression6escapeERK7QString(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegularExpression", "escape", args)
  }

}

// namedCaptureGroups()
func (this *QRegularExpression) namedCaptureGroups(args ...interface{}) () {
  // namedCaptureGroups()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QRegularExpression18namedCaptureGroupsEv
    // invoke: QStringList namedCaptureGroups()
    C.C_ZNK18QRegularExpression18namedCaptureGroupsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpression", "namedCaptureGroups", args)
  }

}

// optimize()
func (this *QRegularExpression) optimize(args ...interface{}) () {
  // optimize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QRegularExpression8optimizeEv
    // invoke: void optimize()
    C.C_ZNK18QRegularExpression8optimizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpression", "optimize", args)
  }

}

// captureCount()
func (this *QRegularExpression) captureCount(args ...interface{}) () {
  // captureCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QRegularExpression12captureCountEv
    // invoke: int captureCount()
    var ret = C.C_ZNK18QRegularExpression12captureCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegularExpression", "captureCount", args)
  }

}

// hasPartialMatch()
func (this *QRegularExpressionMatch) hasPartialMatch(args ...interface{}) () {
  // hasPartialMatch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QRegularExpressionMatch15hasPartialMatchEv
    // invoke: bool hasPartialMatch()
    var ret = C.C_ZNK23QRegularExpressionMatch15hasPartialMatchEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "hasPartialMatch", args)
  }

}

// regularExpression()
func (this *QRegularExpressionMatch) regularExpression(args ...interface{}) () {
  // regularExpression()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QRegularExpressionMatch17regularExpressionEv
    // invoke: QRegularExpression regularExpression()
    var ret = C.C_ZNK23QRegularExpressionMatch17regularExpressionEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "regularExpression", args)
  }

}

// QRegularExpressionMatch(const class QRegularExpressionMatch &)
func NewQRegularExpressionMatch(args ...interface{}) *QRegularExpressionMatch {
  // QRegularExpressionMatch(const class QRegularExpressionMatch &)
  // QRegularExpressionMatch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegularExpressionMatch{}) // "const QRegularExpressionMatch &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QRegularExpressionMatchC1ERKS_
    // invoke: void QRegularExpressionMatch(const class QRegularExpressionMatch &)
    var arg0 = args[0].(QRegularExpressionMatch).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN23QRegularExpressionMatchC2ERKS_(arg0)
    return &QRegularExpressionMatch{qclsinst:qthis}
  case 1:
    // invoke: _ZN23QRegularExpressionMatchC1Ev
    // invoke: void QRegularExpressionMatch()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN23QRegularExpressionMatchC2Ev()
    return &QRegularExpressionMatch{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "QRegularExpressionMatch", args)
  }

  return nil // QRegularExpressionMatch{qclsinst:qthis}
}

// capturedTexts()
func (this *QRegularExpressionMatch) capturedTexts(args ...interface{}) () {
  // capturedTexts()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QRegularExpressionMatch13capturedTextsEv
    // invoke: QStringList capturedTexts()
    C.C_ZNK23QRegularExpressionMatch13capturedTextsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "capturedTexts", args)
  }

}

// isValid()
func (this *QRegularExpressionMatch) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QRegularExpressionMatch7isValidEv
    // invoke: bool isValid()
    var ret = C.C_ZNK23QRegularExpressionMatch7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "isValid", args)
  }

}

// hasMatch()
func (this *QRegularExpressionMatch) hasMatch(args ...interface{}) () {
  // hasMatch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QRegularExpressionMatch8hasMatchEv
    // invoke: bool hasMatch()
    var ret = C.C_ZNK23QRegularExpressionMatch8hasMatchEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "hasMatch", args)
  }

}

// ~QRegularExpressionMatch()
func (this *QRegularExpressionMatch) FreeQRegularExpressionMatch(args ...interface{}) () {
  // ~QRegularExpressionMatch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QRegularExpressionMatchD0Ev
    // invoke: void ~QRegularExpressionMatch()
    C.C_ZN23QRegularExpressionMatchD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "~QRegularExpressionMatch", args)
  }

}

// matchOptions()
func (this *QRegularExpressionMatch) matchOptions(args ...interface{}) () {
  // matchOptions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QRegularExpressionMatch12matchOptionsEv
    // invoke: QRegularExpression::MatchOptions matchOptions()
    C.C_ZNK23QRegularExpressionMatch12matchOptionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "matchOptions", args)
  }

}

// capturedStart(const class QString &)
func (this *QRegularExpressionMatch) capturedStart(args ...interface{}) () {
  // capturedStart(const class QString &)
  // capturedStart(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QRegularExpressionMatch13capturedStartERK7QString
    // invoke: int capturedStart(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK23QRegularExpressionMatch13capturedStartERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK23QRegularExpressionMatch13capturedStartEi
    // invoke: int capturedStart(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK23QRegularExpressionMatch13capturedStartEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "capturedStart", args)
  }

}

// capturedRef(int)
func (this *QRegularExpressionMatch) capturedRef(args ...interface{}) () {
  // capturedRef(int)
  // capturedRef(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QRegularExpressionMatch11capturedRefEi
    // invoke: QStringRef capturedRef(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK23QRegularExpressionMatch11capturedRefEi(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK23QRegularExpressionMatch11capturedRefERK7QString
    // invoke: QStringRef capturedRef(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK23QRegularExpressionMatch11capturedRefERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "capturedRef", args)
  }

}

// lastCapturedIndex()
func (this *QRegularExpressionMatch) lastCapturedIndex(args ...interface{}) () {
  // lastCapturedIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QRegularExpressionMatch17lastCapturedIndexEv
    // invoke: int lastCapturedIndex()
    var ret = C.C_ZNK23QRegularExpressionMatch17lastCapturedIndexEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "lastCapturedIndex", args)
  }

}

// captured(int)
func (this *QRegularExpressionMatch) captured(args ...interface{}) () {
  // captured(int)
  // captured(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QRegularExpressionMatch8capturedEi
    // invoke: QString captured(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK23QRegularExpressionMatch8capturedEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK23QRegularExpressionMatch8capturedERK7QString
    // invoke: QString captured(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK23QRegularExpressionMatch8capturedERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "captured", args)
  }

}

// matchType()
func (this *QRegularExpressionMatch) matchType(args ...interface{}) () {
  // matchType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QRegularExpressionMatch9matchTypeEv
    // invoke: QRegularExpression::MatchType matchType()
    C.C_ZNK23QRegularExpressionMatch9matchTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "matchType", args)
  }

}

// capturedLength(int)
func (this *QRegularExpressionMatch) capturedLength(args ...interface{}) () {
  // capturedLength(int)
  // capturedLength(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QRegularExpressionMatch14capturedLengthEi
    // invoke: int capturedLength(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK23QRegularExpressionMatch14capturedLengthEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK23QRegularExpressionMatch14capturedLengthERK7QString
    // invoke: int capturedLength(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK23QRegularExpressionMatch14capturedLengthERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "capturedLength", args)
  }

}

// capturedEnd(const class QString &)
func (this *QRegularExpressionMatch) capturedEnd(args ...interface{}) () {
  // capturedEnd(const class QString &)
  // capturedEnd(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QRegularExpressionMatch11capturedEndERK7QString
    // invoke: int capturedEnd(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK23QRegularExpressionMatch11capturedEndERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK23QRegularExpressionMatch11capturedEndEi
    // invoke: int capturedEnd(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK23QRegularExpressionMatch11capturedEndEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "capturedEnd", args)
  }

}

// swap(class QRegularExpressionMatch &)
func (this *QRegularExpressionMatch) swap(args ...interface{}) () {
  // swap(class QRegularExpressionMatch &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegularExpressionMatch{}) // "QRegularExpressionMatch &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QRegularExpressionMatch4swapERS_
    // invoke: void swap(class QRegularExpressionMatch &)
    var arg0 = args[0].(QRegularExpressionMatch).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN23QRegularExpressionMatch4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "swap", args)
  }

}

// <= body block end

