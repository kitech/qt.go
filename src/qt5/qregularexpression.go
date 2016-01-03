package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
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
  // proto:  bool QRegularExpressionMatchIterator::hasNext();
extern void _ZNK31QRegularExpressionMatchIterator7hasNextEv(void* qthis);
  // proto:  bool QRegularExpressionMatchIterator::isValid();
extern void _ZNK31QRegularExpressionMatchIterator7isValidEv(void* qthis);
  // proto:  QRegularExpressionMatch QRegularExpressionMatchIterator::peekNext();
extern void _ZNK31QRegularExpressionMatchIterator8peekNextEv(void* qthis);
  // proto:  void QRegularExpressionMatchIterator::QRegularExpressionMatchIterator();
extern void* dector_ZN31QRegularExpressionMatchIteratorC1Ev();
extern void _ZN31QRegularExpressionMatchIteratorC1Ev(void* qthis);
  // proto:  QRegularExpression QRegularExpressionMatchIterator::regularExpression();
extern void _ZNK31QRegularExpressionMatchIterator17regularExpressionEv(void* qthis);
  // proto:  void QRegularExpressionMatchIterator::QRegularExpressionMatchIterator(const QRegularExpressionMatchIterator & iterator);
extern void* dector_ZN31QRegularExpressionMatchIteratorC1ERKS_(void* arg0);
extern void _ZN31QRegularExpressionMatchIteratorC1ERKS_(void* qthis, void* arg0);
  // proto:  void QRegularExpressionMatchIterator::~QRegularExpressionMatchIterator();
extern void _ZN31QRegularExpressionMatchIteratorD0Ev(void* qthis);
  // proto:  QRegularExpressionMatch QRegularExpressionMatchIterator::next();
extern void _ZN31QRegularExpressionMatchIterator4nextEv(void* qthis);
  // proto:  void QRegularExpressionMatchIterator::swap(QRegularExpressionMatchIterator & other);
extern void demth_ZN31QRegularExpressionMatchIterator4swapERS_(void* qthis, void* arg0);
  // proto:  int QRegularExpression::patternErrorOffset();
extern void _ZNK18QRegularExpression18patternErrorOffsetEv(void* qthis);
  // proto:  QString QRegularExpression::pattern();
extern void _ZNK18QRegularExpression7patternEv(void* qthis);
  // proto:  void QRegularExpression::~QRegularExpression();
extern void _ZN18QRegularExpressionD0Ev(void* qthis);
  // proto:  void QRegularExpression::optimize();
extern void _ZNK18QRegularExpression8optimizeEv(void* qthis);
  // proto: static QString QRegularExpression::escape(const QString & str);
extern void _ZN18QRegularExpression6escapeERK7QString(void* arg0);
  // proto:  void QRegularExpression::QRegularExpression();
extern void* dector_ZN18QRegularExpressionC1Ev();
extern void _ZN18QRegularExpressionC1Ev(void* qthis);
  // proto:  void QRegularExpression::swap(QRegularExpression & other);
extern void demth_ZN18QRegularExpression4swapERS_(void* qthis, void* arg0);
  // proto:  QString QRegularExpression::errorString();
extern void _ZNK18QRegularExpression11errorStringEv(void* qthis);
  // proto:  bool QRegularExpression::isValid();
extern void _ZNK18QRegularExpression7isValidEv(void* qthis);
  // proto:  void QRegularExpression::QRegularExpression(const QRegularExpression & re);
extern void* dector_ZN18QRegularExpressionC1ERKS_(void* arg0);
extern void _ZN18QRegularExpressionC1ERKS_(void* qthis, void* arg0);
  // proto:  QStringList QRegularExpression::namedCaptureGroups();
extern void _ZNK18QRegularExpression18namedCaptureGroupsEv(void* qthis);
  // proto:  int QRegularExpression::captureCount();
extern void _ZNK18QRegularExpression12captureCountEv(void* qthis);
  // proto:  void QRegularExpression::setPattern(const QString & pattern);
extern void _ZN18QRegularExpression10setPatternERK7QString(void* qthis, void* arg0);
  // proto:  int QRegularExpressionMatch::lastCapturedIndex();
extern void _ZNK23QRegularExpressionMatch17lastCapturedIndexEv(void* qthis);
  // proto:  void QRegularExpressionMatch::QRegularExpressionMatch();
extern void* dector_ZN23QRegularExpressionMatchC1Ev();
extern void _ZN23QRegularExpressionMatchC1Ev(void* qthis);
  // proto:  bool QRegularExpressionMatch::isValid();
extern void _ZNK23QRegularExpressionMatch7isValidEv(void* qthis);
  // proto:  int QRegularExpressionMatch::capturedLength(int nth);
extern void _ZNK23QRegularExpressionMatch14capturedLengthEi(void* qthis, int32_t arg0);
  // proto:  int QRegularExpressionMatch::capturedLength(const QString & name);
extern void _ZNK23QRegularExpressionMatch14capturedLengthERK7QString(void* qthis, void* arg0);
  // proto:  QStringRef QRegularExpressionMatch::capturedRef(int nth);
extern void _ZNK23QRegularExpressionMatch11capturedRefEi(void* qthis, int32_t arg0);
  // proto:  int QRegularExpressionMatch::capturedEnd(const QString & name);
extern void _ZNK23QRegularExpressionMatch11capturedEndERK7QString(void* qthis, void* arg0);
  // proto:  QString QRegularExpressionMatch::captured(const QString & name);
extern void _ZNK23QRegularExpressionMatch8capturedERK7QString(void* qthis, void* arg0);
  // proto:  QStringList QRegularExpressionMatch::capturedTexts();
extern void _ZNK23QRegularExpressionMatch13capturedTextsEv(void* qthis);
  // proto:  void QRegularExpressionMatch::QRegularExpressionMatch(const QRegularExpressionMatch & match);
extern void* dector_ZN23QRegularExpressionMatchC1ERKS_(void* arg0);
extern void _ZN23QRegularExpressionMatchC1ERKS_(void* qthis, void* arg0);
  // proto:  void QRegularExpressionMatch::swap(QRegularExpressionMatch & other);
extern void demth_ZN23QRegularExpressionMatch4swapERS_(void* qthis, void* arg0);
  // proto:  void QRegularExpressionMatch::~QRegularExpressionMatch();
extern void _ZN23QRegularExpressionMatchD0Ev(void* qthis);
  // proto:  int QRegularExpressionMatch::capturedEnd(int nth);
extern void _ZNK23QRegularExpressionMatch11capturedEndEi(void* qthis, int32_t arg0);
  // proto:  QStringRef QRegularExpressionMatch::capturedRef(const QString & name);
extern void _ZNK23QRegularExpressionMatch11capturedRefERK7QString(void* qthis, void* arg0);
  // proto:  bool QRegularExpressionMatch::hasMatch();
extern void _ZNK23QRegularExpressionMatch8hasMatchEv(void* qthis);
  // proto:  int QRegularExpressionMatch::capturedStart(const QString & name);
extern void _ZNK23QRegularExpressionMatch13capturedStartERK7QString(void* qthis, void* arg0);
  // proto:  QRegularExpression QRegularExpressionMatch::regularExpression();
extern void _ZNK23QRegularExpressionMatch17regularExpressionEv(void* qthis);
  // proto:  QString QRegularExpressionMatch::captured(int nth);
extern void _ZNK23QRegularExpressionMatch8capturedEi(void* qthis, int32_t arg0);
  // proto:  int QRegularExpressionMatch::capturedStart(int nth);
extern void _ZNK23QRegularExpressionMatch13capturedStartEi(void* qthis, int32_t arg0);
  // proto:  bool QRegularExpressionMatch::hasPartialMatch();
extern void _ZNK23QRegularExpressionMatch15hasPartialMatchEv(void* qthis);
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

  // proto:  bool QRegularExpressionMatchIterator::hasNext();
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
    C._ZNK31QRegularExpressionMatchIterator7hasNextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatchIterator", "hasNext", args)
  }

}

  // proto:  bool QRegularExpressionMatchIterator::isValid();
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
    C._ZNK31QRegularExpressionMatchIterator7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatchIterator", "isValid", args)
  }

}

  // proto:  QRegularExpressionMatch QRegularExpressionMatchIterator::peekNext();
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
    C._ZNK31QRegularExpressionMatchIterator8peekNextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatchIterator", "peekNext", args)
  }

}

  // proto:  void QRegularExpressionMatchIterator::QRegularExpressionMatchIterator();
func NewQRegularExpressionMatchIterator(args ...interface{}) QRegularExpressionMatchIterator {
  return QRegularExpressionMatchIterator{}
}

  // proto:  QRegularExpression QRegularExpressionMatchIterator::regularExpression();
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
    C._ZNK31QRegularExpressionMatchIterator17regularExpressionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatchIterator", "regularExpression", args)
  }

}

  // proto:  void QRegularExpressionMatchIterator::~QRegularExpressionMatchIterator();
func (this *QRegularExpressionMatchIterator) FreeQRegularExpressionMatchIterator(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QRegularExpressionMatchIterator", "~QRegularExpressionMatchIterator", args)
  }

}

  // proto:  QRegularExpressionMatch QRegularExpressionMatchIterator::next();
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
    C._ZN31QRegularExpressionMatchIterator4nextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatchIterator", "next", args)
  }

}

  // proto:  void QRegularExpressionMatchIterator::swap(QRegularExpressionMatchIterator & other);
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
    C.demth_ZN31QRegularExpressionMatchIterator4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatchIterator", "swap", args)
  }

}

  // proto:  int QRegularExpression::patternErrorOffset();
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
    C._ZNK18QRegularExpression18patternErrorOffsetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpression", "patternErrorOffset", args)
  }

}

  // proto:  QString QRegularExpression::pattern();
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
    C._ZNK18QRegularExpression7patternEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpression", "pattern", args)
  }

}

  // proto:  void QRegularExpression::~QRegularExpression();
func (this *QRegularExpression) FreeQRegularExpression(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QRegularExpression", "~QRegularExpression", args)
  }

}

  // proto:  void QRegularExpression::optimize();
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
    C._ZNK18QRegularExpression8optimizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpression", "optimize", args)
  }

}

  // proto: static QString QRegularExpression::escape(const QString & str);
func (this *QRegularExpression) escape_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QRegularExpression", "escape", args)
  }

}

  // proto:  void QRegularExpression::QRegularExpression();
func NewQRegularExpression(args ...interface{}) QRegularExpression {
  return QRegularExpression{}
}

  // proto:  void QRegularExpression::swap(QRegularExpression & other);
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
    C.demth_ZN18QRegularExpression4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegularExpression", "swap", args)
  }

}

  // proto:  QString QRegularExpression::errorString();
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
    C._ZNK18QRegularExpression11errorStringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpression", "errorString", args)
  }

}

  // proto:  bool QRegularExpression::isValid();
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
    C._ZNK18QRegularExpression7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpression", "isValid", args)
  }

}

  // proto:  QStringList QRegularExpression::namedCaptureGroups();
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
    C._ZNK18QRegularExpression18namedCaptureGroupsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpression", "namedCaptureGroups", args)
  }

}

  // proto:  int QRegularExpression::captureCount();
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
    C._ZNK18QRegularExpression12captureCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpression", "captureCount", args)
  }

}

  // proto:  void QRegularExpression::setPattern(const QString & pattern);
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
    C._ZN18QRegularExpression10setPatternERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegularExpression", "setPattern", args)
  }

}

  // proto:  int QRegularExpressionMatch::lastCapturedIndex();
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
    C._ZNK23QRegularExpressionMatch17lastCapturedIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "lastCapturedIndex", args)
  }

}

  // proto:  void QRegularExpressionMatch::QRegularExpressionMatch();
func NewQRegularExpressionMatch(args ...interface{}) QRegularExpressionMatch {
  return QRegularExpressionMatch{}
}

  // proto:  bool QRegularExpressionMatch::isValid();
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
    C._ZNK23QRegularExpressionMatch7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "isValid", args)
  }

}

  // proto:  int QRegularExpressionMatch::capturedLength(int nth);
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
    C._ZNK23QRegularExpressionMatch14capturedLengthEi(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK23QRegularExpressionMatch14capturedLengthERK7QString
    // invoke: int capturedLength(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK23QRegularExpressionMatch14capturedLengthERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "capturedLength", args)
  }

}

  // proto:  QStringRef QRegularExpressionMatch::capturedRef(int nth);
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
    C._ZNK23QRegularExpressionMatch11capturedRefEi(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK23QRegularExpressionMatch11capturedRefERK7QString
    // invoke: QStringRef capturedRef(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK23QRegularExpressionMatch11capturedRefERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "capturedRef", args)
  }

}

  // proto:  int QRegularExpressionMatch::capturedEnd(const QString & name);
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
    C._ZNK23QRegularExpressionMatch11capturedEndERK7QString(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK23QRegularExpressionMatch11capturedEndEi
    // invoke: int capturedEnd(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK23QRegularExpressionMatch11capturedEndEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "capturedEnd", args)
  }

}

  // proto:  QString QRegularExpressionMatch::captured(const QString & name);
func (this *QRegularExpressionMatch) captured(args ...interface{}) () {
  // captured(const class QString &)
  // captured(int)
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
    // invoke: _ZNK23QRegularExpressionMatch8capturedERK7QString
    // invoke: QString captured(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK23QRegularExpressionMatch8capturedERK7QString(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK23QRegularExpressionMatch8capturedEi
    // invoke: QString captured(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK23QRegularExpressionMatch8capturedEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "captured", args)
  }

}

  // proto:  QStringList QRegularExpressionMatch::capturedTexts();
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
    C._ZNK23QRegularExpressionMatch13capturedTextsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "capturedTexts", args)
  }

}

  // proto:  void QRegularExpressionMatch::swap(QRegularExpressionMatch & other);
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
    C.demth_ZN23QRegularExpressionMatch4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "swap", args)
  }

}

  // proto:  void QRegularExpressionMatch::~QRegularExpressionMatch();
func (this *QRegularExpressionMatch) FreeQRegularExpressionMatch(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "~QRegularExpressionMatch", args)
  }

}

  // proto:  bool QRegularExpressionMatch::hasMatch();
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
    C._ZNK23QRegularExpressionMatch8hasMatchEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "hasMatch", args)
  }

}

  // proto:  int QRegularExpressionMatch::capturedStart(const QString & name);
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
    C._ZNK23QRegularExpressionMatch13capturedStartERK7QString(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK23QRegularExpressionMatch13capturedStartEi
    // invoke: int capturedStart(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK23QRegularExpressionMatch13capturedStartEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "capturedStart", args)
  }

}

  // proto:  QRegularExpression QRegularExpressionMatch::regularExpression();
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
    C._ZNK23QRegularExpressionMatch17regularExpressionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "regularExpression", args)
  }

}

  // proto:  bool QRegularExpressionMatch::hasPartialMatch();
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
    C._ZNK23QRegularExpressionMatch15hasPartialMatchEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "hasPartialMatch", args)
  }

}

// <= body block end

