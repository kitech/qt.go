package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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
extern void C_ZN9QCollator20setIgnorePunctuationEb(void* qthis, bool arg0); // 4
  // proto:  int QCollator::compare(const QChar * s1, int len1, const QChar * s2, int len2);
extern int32_t C_ZNK9QCollator7compareEPK5QChariS2_i(void* qthis, void* arg0, int32_t arg1, void* arg2, int32_t arg3); // 4
  // proto:  int QCollator::compare(const QString & s1, const QString & s2);
extern int32_t C_ZNK9QCollator7compareERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QCollator::setLocale(const QLocale & locale);
extern void C_ZN9QCollator9setLocaleERK7QLocale(void* qthis, void* arg0); // 4
  // proto:  QLocale QCollator::locale();
extern void* C_ZNK9QCollator6localeEv(void* qthis); // 4
  // proto:  bool QCollator::numericMode();
extern bool C_ZNK9QCollator11numericModeEv(void* qthis); // 4
  // proto:  void QCollator::setNumericMode(bool on);
extern void C_ZN9QCollator14setNumericModeEb(void* qthis, bool arg0); // 4
  // proto:  QCollatorSortKey QCollator::sortKey(const QString & string);
extern void* C_ZNK9QCollator7sortKeyERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QCollator::~QCollator();
extern void C_ZN9QCollatorD2Ev(void* qthis); // 4
  // proto:  bool QCollator::ignorePunctuation();
extern bool C_ZNK9QCollator17ignorePunctuationEv(void* qthis); // 4
  // proto:  void QCollator::swap(QCollator & other);
extern void C_ZN9QCollator4swapERS_(void* qthis, void* arg0); // 2
  // proto:  void QCollator::QCollator(const QLocale & locale);
extern void* C_ZN9QCollatorC2ERK7QLocale(void* arg0); // 3
  // proto:  void QCollator::QCollator(const QCollator & );
extern void* C_ZN9QCollatorC2ERKS_(void* arg0); // 3
  // proto:  Qt::CaseSensitivity QCollator::caseSensitivity();
extern void C_ZNK9QCollator15caseSensitivityEv(void* qthis); // 4
  // proto:  void QCollatorSortKey::QCollatorSortKey(const QCollatorSortKey & other);
extern void* C_ZN16QCollatorSortKeyC2ERKS_(void* arg0); // 3
  // proto:  int QCollatorSortKey::compare(const QCollatorSortKey & key);
extern int32_t C_ZNK16QCollatorSortKey7compareERKS_(void* qthis, void* arg0); // 4
  // proto:  void QCollatorSortKey::swap(QCollatorSortKey & other);
extern void C_ZN16QCollatorSortKey4swapERS_(void* qthis, void* arg0); // 2
  // proto:  void QCollatorSortKey::~QCollatorSortKey();
extern void C_ZN16QCollatorSortKeyD2Ev(void* qthis); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QCollatorSortKey)=1
type QCollatorSortKey struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// setIgnorePunctuation(_Bool)
func (this *QCollator) Setignorepunctuation(args ...interface{}) () {
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
    C.C_ZN9QCollator20setIgnorePunctuationEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCollator", "setIgnorePunctuation", args)
  }

  return
}

// compare(const class QChar *, int, const class QChar *, int)
func (this *QCollator) Compare(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QChar).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK9QCollator7compareEPK5QChariS2_i(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK9QCollator7compareERK7QStringS2_
    // invoke: int compare(const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK9QCollator7compareERK7QStringS2_(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCollator", "compare", args)
  }

  return
}

// setLocale(const class QLocale &)
func (this *QCollator) Setlocale(args ...interface{}) () {
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
    var arg0 = args[0].(*QLocale).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QCollator9setLocaleERK7QLocale(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCollator", "setLocale", args)
  }

  return
}

// locale()
func (this *QCollator) Locale(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QCollator6localeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLocale{}) // "QLocale"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCollator", "locale", args)
  }

  return
}

// numericMode()
func (this *QCollator) Numericmode(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QCollator11numericModeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCollator", "numericMode", args)
  }

  return
}

// setNumericMode(_Bool)
func (this *QCollator) Setnumericmode(args ...interface{}) () {
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
    C.C_ZN9QCollator14setNumericModeEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCollator", "setNumericMode", args)
  }

  return
}

// sortKey(const class QString &)
func (this *QCollator) Sortkey(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QCollator7sortKeyERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QCollatorSortKey{}) // "QCollatorSortKey"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCollator", "sortKey", args)
  }

  return
}

// ~QCollator()
func (this *QCollator) Freeqcollator(args ...interface{}) () {
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
    C.C_ZN9QCollatorD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCollator", "~QCollator", args)
  }

  return
}

// ignorePunctuation()
func (this *QCollator) Ignorepunctuation(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QCollator17ignorePunctuationEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCollator", "ignorePunctuation", args)
  }

  return
}

// swap(class QCollator &)
func (this *QCollator) Swap(args ...interface{}) () {
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
    var arg0 = args[0].(*QCollator).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QCollator4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCollator", "swap", args)
  }

  return
}

// QCollator(const class QLocale &)
func NewQCollator(args ...interface{}) *QCollator {
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
    var arg0 = args[0].(*QLocale).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QCollatorC2ERK7QLocale(arg0)
    return &QCollator{Qclsinst:qthis}
  case 1:
    // invoke: _ZN9QCollatorC1ERKS_
    // invoke: void QCollator(const class QCollator &)
    var arg0 = args[0].(*QCollator).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QCollatorC2ERKS_(arg0)
    return &QCollator{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QCollator", "QCollator", args)
  }

  return nil // QCollator{Qclsinst:qthis}
}

// caseSensitivity()
func (this *QCollator) Casesensitivity(args ...interface{}) () {
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
    C.C_ZNK9QCollator15caseSensitivityEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCollator", "caseSensitivity", args)
  }

  return
}

// QCollatorSortKey(const class QCollatorSortKey &)
func NewQCollatorSortKey(args ...interface{}) *QCollatorSortKey {
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
    var arg0 = args[0].(*QCollatorSortKey).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QCollatorSortKeyC2ERKS_(arg0)
    return &QCollatorSortKey{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QCollatorSortKey", "QCollatorSortKey", args)
  }

  return nil // QCollatorSortKey{Qclsinst:qthis}
}

// compare(const class QCollatorSortKey &)
func (this *QCollatorSortKey) Compare(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QCollatorSortKey).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK16QCollatorSortKey7compareERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCollatorSortKey", "compare", args)
  }

  return
}

// swap(class QCollatorSortKey &)
func (this *QCollatorSortKey) Swap(args ...interface{}) () {
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
    var arg0 = args[0].(*QCollatorSortKey).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QCollatorSortKey4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCollatorSortKey", "swap", args)
  }

  return
}

// ~QCollatorSortKey()
func (this *QCollatorSortKey) Freeqcollatorsortkey(args ...interface{}) () {
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
    C.C_ZN16QCollatorSortKeyD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCollatorSortKey", "~QCollatorSortKey", args)
  }

  return
}

// <= body block end

