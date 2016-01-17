package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtCore/qcollator.h
// dst-file: /src/core/qcollator.go
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
  // proto:  void QCollator::setIgnorePunctuation(bool on);
extern void _ZN9QCollator20setIgnorePunctuationEb(void* qthis, bool arg0); // 4
  // proto:  int QCollator::compare(const QChar * s1, int len1, const QChar * s2, int len2);
extern void _ZNK9QCollator7compareEPK5QChariS2_i(void* qthis, void* arg0, int32_t arg1, void* arg2, int32_t arg3); // 4
  // proto:  int QCollator::compare(const QString & s1, const QString & s2);
extern void _ZNK9QCollator7compareERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QCollator::setLocale(const QLocale & locale);
extern void _ZN9QCollator9setLocaleERK7QLocale(void* qthis, void* arg0); // 4
  // proto:  QLocale QCollator::locale();
extern void _ZNK9QCollator6localeEv(void* qthis); // 4
  // proto:  bool QCollator::numericMode();
extern void _ZNK9QCollator11numericModeEv(void* qthis); // 4
  // proto:  void QCollator::setNumericMode(bool on);
extern void _ZN9QCollator14setNumericModeEb(void* qthis, bool arg0); // 4
  // proto:  QCollatorSortKey QCollator::sortKey(const QString & string);
extern void _ZNK9QCollator7sortKeyERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QCollator::~QCollator();
extern void _ZN9QCollatorD2Ev(void* qthis); // 4
  // proto:  bool QCollator::ignorePunctuation();
extern void _ZNK9QCollator17ignorePunctuationEv(void* qthis); // 4
  // proto:  void QCollator::swap(QCollator & other);
extern void _ZN9QCollator4swapERS_(void* qthis, void* arg0); // 2
  // proto:  void QCollator::QCollator(const QLocale & locale);
extern void _ZN9QCollatorC2ERK7QLocale(void* qthis, void* arg0); // 3
  // proto:  void QCollator::QCollator(const QCollator & );
extern void _ZN9QCollatorC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  Qt::CaseSensitivity QCollator::caseSensitivity();
extern void _ZNK9QCollator15caseSensitivityEv(void* qthis); // 4
  // proto:  void QCollatorSortKey::QCollatorSortKey(const QCollatorSortKey & other);
extern void _ZN16QCollatorSortKeyC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  int QCollatorSortKey::compare(const QCollatorSortKey & key);
extern void _ZNK16QCollatorSortKey7compareERKS_(void* qthis, void* arg0); // 4
  // proto:  void QCollatorSortKey::swap(QCollatorSortKey & other);
extern void _ZN16QCollatorSortKey4swapERS_(void* qthis, void* arg0); // 2
  // proto:  void QCollatorSortKey::~QCollatorSortKey();
extern void _ZN16QCollatorSortKeyD2Ev(void* qthis); // 4
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

// class sizeof(QCollator)=8
type QCollator struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QCollatorSortKey)=1
type QCollatorSortKey struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// setIgnorePunctuation(_Bool)
func (this *QCollator) setIgnorePunctuation(args ...interface{}) () {
  // setIgnorePunctuation(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QCollator20setIgnorePunctuationEb
    // invoke: void setIgnorePunctuation(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QCollator20setIgnorePunctuationEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCollator", "setIgnorePunctuation", args)
  }

}

// compare(const class QChar *, int, const class QChar *, int)
func (this *QCollator) compare(args ...interface{}) () {
  // compare(const class QChar *, int, const class QChar *, int)
  // compare(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QCollator7compareEPK5QChariS2_i
    // invoke: int compare(const class QChar *, int, const class QChar *, int)
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QChar).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C._ZNK9QCollator7compareEPK5QChariS2_i(this.qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZNK9QCollator7compareERK7QStringS2_
    // invoke: int compare(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK9QCollator7compareERK7QStringS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QCollator", "compare", args)
  }

}

// setLocale(const class QLocale &)
func (this *QCollator) setLocale(args ...interface{}) () {
  // setLocale(const class QLocale &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLocale{}) // "const QLocale &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QCollator9setLocaleERK7QLocale
    // invoke: void setLocale(const class QLocale &)
    var arg0 = args[0].(QLocale).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QCollator9setLocaleERK7QLocale(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCollator", "setLocale", args)
  }

}

// locale()
func (this *QCollator) locale(args ...interface{}) () {
  // locale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QCollator6localeEv
    // invoke: QLocale locale()
    C._ZNK9QCollator6localeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCollator", "locale", args)
  }

}

// numericMode()
func (this *QCollator) numericMode(args ...interface{}) () {
  // numericMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QCollator11numericModeEv
    // invoke: bool numericMode()
    C._ZNK9QCollator11numericModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCollator", "numericMode", args)
  }

}

// setNumericMode(_Bool)
func (this *QCollator) setNumericMode(args ...interface{}) () {
  // setNumericMode(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QCollator14setNumericModeEb
    // invoke: void setNumericMode(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QCollator14setNumericModeEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCollator", "setNumericMode", args)
  }

}

// sortKey(const class QString &)
func (this *QCollator) sortKey(args ...interface{}) () {
  // sortKey(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QCollator7sortKeyERK7QString
    // invoke: QCollatorSortKey sortKey(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK9QCollator7sortKeyERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCollator", "sortKey", args)
  }

}

// ~QCollator()
func (this *QCollator) FreeQCollator(args ...interface{}) () {
  // ~QCollator()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QCollatorD0Ev
    // invoke: void ~QCollator()
    C._ZN9QCollatorD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCollator", "~QCollator", args)
  }

}

// ignorePunctuation()
func (this *QCollator) ignorePunctuation(args ...interface{}) () {
  // ignorePunctuation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QCollator17ignorePunctuationEv
    // invoke: bool ignorePunctuation()
    C._ZNK9QCollator17ignorePunctuationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCollator", "ignorePunctuation", args)
  }

}

// swap(class QCollator &)
func (this *QCollator) swap(args ...interface{}) () {
  // swap(class QCollator &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCollator{}) // "QCollator &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QCollator4swapERS_
    // invoke: void swap(class QCollator &)
    var arg0 = args[0].(QCollator).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QCollator4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCollator", "swap", args)
  }

}

// QCollator(const class QLocale &)
func NewQCollator(args ...interface{}) QCollator {
  // QCollator(const class QLocale &)
  // QCollator(const class QCollator &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLocale{}) // "const QLocale &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QCollator{}) // "const QCollator &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QCollatorC1ERK7QLocale
    // invoke: void QCollator(const class QLocale &)
    var arg0 = args[0].(QLocale).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN9QCollatorC2ERK7QLocale(qthis, arg0)
  case 1:
    // invoke: _ZN9QCollatorC1ERKS_
    // invoke: void QCollator(const class QCollator &)
    var arg0 = args[0].(QCollator).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN9QCollatorC2ERKS_(qthis, arg0)
  default:
    qtrt.ErrorResolve("QCollator", "QCollator", args)
  }

  return QCollator{}
}

// caseSensitivity()
func (this *QCollator) caseSensitivity(args ...interface{}) () {
  // caseSensitivity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QCollator15caseSensitivityEv
    // invoke: Qt::CaseSensitivity caseSensitivity()
    C._ZNK9QCollator15caseSensitivityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCollator", "caseSensitivity", args)
  }

}

// QCollatorSortKey(const class QCollatorSortKey &)
func NewQCollatorSortKey(args ...interface{}) QCollatorSortKey {
  // QCollatorSortKey(const class QCollatorSortKey &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCollatorSortKey{}) // "const QCollatorSortKey &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCollatorSortKeyC1ERKS_
    // invoke: void QCollatorSortKey(const class QCollatorSortKey &)
    var arg0 = args[0].(QCollatorSortKey).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN16QCollatorSortKeyC2ERKS_(qthis, arg0)
  default:
    qtrt.ErrorResolve("QCollatorSortKey", "QCollatorSortKey", args)
  }

  return QCollatorSortKey{}
}

// compare(const class QCollatorSortKey &)
func (this *QCollatorSortKey) compare(args ...interface{}) () {
  // compare(const class QCollatorSortKey &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCollatorSortKey{}) // "const QCollatorSortKey &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QCollatorSortKey7compareERKS_
    // invoke: int compare(const class QCollatorSortKey &)
    var arg0 = args[0].(QCollatorSortKey).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK16QCollatorSortKey7compareERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCollatorSortKey", "compare", args)
  }

}

// swap(class QCollatorSortKey &)
func (this *QCollatorSortKey) swap(args ...interface{}) () {
  // swap(class QCollatorSortKey &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCollatorSortKey{}) // "QCollatorSortKey &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCollatorSortKey4swapERS_
    // invoke: void swap(class QCollatorSortKey &)
    var arg0 = args[0].(QCollatorSortKey).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN16QCollatorSortKey4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCollatorSortKey", "swap", args)
  }

}

// ~QCollatorSortKey()
func (this *QCollatorSortKey) FreeQCollatorSortKey(args ...interface{}) () {
  // ~QCollatorSortKey()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCollatorSortKeyD0Ev
    // invoke: void ~QCollatorSortKey()
    C._ZN16QCollatorSortKeyD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCollatorSortKey", "~QCollatorSortKey", args)
  }

}

// <= body block end

