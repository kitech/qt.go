package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtCore/qchar.h
// dst-file: /src/core/qchar.go
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
  // proto:  ushort QLatin1Char::unicode();
extern void demth_ZNK11QLatin1Char7unicodeEv(void* qthis);
  // proto:  void QLatin1Char::QLatin1Char(char c);
extern void* dector_ZN11QLatin1CharC1Ec(unsigned char arg0);
extern void _ZN11QLatin1CharC1Ec(void* qthis, unsigned char arg0);
  // proto:  char QLatin1Char::toLatin1();
extern void demth_ZNK11QLatin1Char8toLatin1Ev(void* qthis);
  // proto: static uint QChar::toUpper(uint ucs4);
extern void _ZN5QChar7toUpperEj(int32_t arg0);
  // proto:  bool QChar::hasMirrored();
extern void demth_ZNK5QChar11hasMirroredEv(void* qthis);
  // proto: static ushort QChar::lowSurrogate(uint ucs4);
extern void demth_ZN5QChar12lowSurrogateEj(int32_t arg0);
  // proto: static bool QChar::isSymbol(uint ucs4);
extern void _ZN5QChar8isSymbolEj(int32_t arg0);
  // proto:  uchar QChar::cell();
extern void demth_ZNK5QChar4cellEv(void* qthis);
  // proto: static uint QChar::surrogateToUcs4(QChar high, QChar low);
extern void demth_ZN5QChar15surrogateToUcs4ES_S_(void* arg0, void* arg1);
  // proto: static bool QChar::isTitleCase(uint ucs4);
extern void demth_ZN5QChar11isTitleCaseEj(int32_t arg0);
  // proto:  bool QChar::isNull();
extern void demth_ZNK5QChar6isNullEv(void* qthis);
  // proto:  int QChar::digitValue();
extern void demth_ZNK5QChar10digitValueEv(void* qthis);
  // proto:  QChar QChar::toTitleCase();
extern void demth_ZNK5QChar11toTitleCaseEv(void* qthis);
  // proto: static bool QChar::isLower(uint ucs4);
extern void demth_ZN5QChar7isLowerEj(int32_t arg0);
  // proto:  bool QChar::isLowSurrogate();
extern void demth_ZNK5QChar14isLowSurrogateEv(void* qthis);
  // proto: static bool QChar::isPrint(uint ucs4);
extern void _ZN5QChar7isPrintEj(int32_t arg0);
  // proto:  bool QChar::isSymbol();
extern void demth_ZNK5QChar8isSymbolEv(void* qthis);
  // proto:  bool QChar::isLower();
extern void demth_ZNK5QChar7isLowerEv(void* qthis);
  // proto:  QChar QChar::mirroredChar();
extern void demth_ZNK5QChar12mirroredCharEv(void* qthis);
  // proto:  uchar QChar::row();
extern void demth_ZNK5QChar3rowEv(void* qthis);
  // proto:  bool QChar::isDigit();
extern void demth_ZNK5QChar7isDigitEv(void* qthis);
  // proto: static uint QChar::toTitleCase(uint ucs4);
extern void _ZN5QChar11toTitleCaseEj(int32_t arg0);
  // proto:  bool QChar::isSurrogate();
extern void demth_ZNK5QChar11isSurrogateEv(void* qthis);
  // proto: static bool QChar::hasMirrored(uint ucs4);
extern void _ZN5QChar11hasMirroredEj(int32_t arg0);
  // proto:  bool QChar::isNumber();
extern void demth_ZNK5QChar8isNumberEv(void* qthis);
  // proto: static bool QChar::isHighSurrogate(uint ucs4);
extern void demth_ZN5QChar15isHighSurrogateEj(int32_t arg0);
  // proto: static uint QChar::toCaseFolded(uint ucs4);
extern void _ZN5QChar12toCaseFoldedEj(int32_t arg0);
  // proto:  bool QChar::isMark();
extern void demth_ZNK5QChar6isMarkEv(void* qthis);
  // proto: static uint QChar::surrogateToUcs4(ushort high, ushort low);
extern void demth_ZN5QChar15surrogateToUcs4Ett(int16_t arg0, int16_t arg1);
  // proto: static uint QChar::toLower(uint ucs4);
extern void _ZN5QChar7toLowerEj(int32_t arg0);
  // proto: static uint QChar::mirroredChar(uint ucs4);
extern void _ZN5QChar12mirroredCharEj(int32_t arg0);
  // proto:  void QChar::setRow(uchar row);
extern void demth_ZN5QChar6setRowEh(void* qthis, unsigned char arg0);
  // proto: static QString QChar::decomposition(uint ucs4);
extern void _ZN5QChar13decompositionEj(int32_t arg0);
  // proto: static int QChar::digitValue(uint ucs4);
extern void _ZN5QChar10digitValueEj(int32_t arg0);
  // proto:  void QChar::setCell(uchar cell);
extern void demth_ZN5QChar7setCellEh(void* qthis, unsigned char arg0);
  // proto: static bool QChar::isUpper(uint ucs4);
extern void demth_ZN5QChar7isUpperEj(int32_t arg0);
  // proto:  void QChar::QChar(uchar c, uchar r);
extern void* dector_ZN5QCharC1Ehh(unsigned char arg0, unsigned char arg1);
extern void _ZN5QCharC1Ehh(void* qthis, unsigned char arg0, unsigned char arg1);
  // proto:  QChar QChar::toCaseFolded();
extern void demth_ZNK5QChar12toCaseFoldedEv(void* qthis);
  // proto:  bool QChar::isPrint();
extern void demth_ZNK5QChar7isPrintEv(void* qthis);
  // proto:  void QChar::QChar(char c);
extern void* dector_ZN5QCharC1Ec(unsigned char arg0);
extern void _ZN5QCharC1Ec(void* qthis, unsigned char arg0);
  // proto:  bool QChar::isPunct();
extern void demth_ZNK5QChar7isPunctEv(void* qthis);
  // proto:  QString QChar::decomposition();
extern void _ZNK5QChar13decompositionEv(void* qthis);
  // proto:  void QChar::QChar(uint rc);
extern void* dector_ZN5QCharC1Ej(int32_t arg0);
extern void _ZN5QCharC1Ej(void* qthis, int32_t arg0);
  // proto:  void QChar::QChar(int rc);
extern void* dector_ZN5QCharC1Ei(int32_t arg0);
extern void _ZN5QCharC1Ei(void* qthis, int32_t arg0);
  // proto:  bool QChar::isSpace();
extern void demth_ZNK5QChar7isSpaceEv(void* qthis);
  // proto:  void QChar::QChar(short rc);
extern void* dector_ZN5QCharC1Es(int16_t arg0);
extern void _ZN5QCharC1Es(void* qthis, int16_t arg0);
  // proto:  void QChar::QChar();
extern void* dector_ZN5QCharC1Ev();
extern void _ZN5QCharC1Ev(void* qthis);
  // proto:  void QChar::QChar(ushort rc);
extern void* dector_ZN5QCharC1Et(int16_t arg0);
extern void _ZN5QCharC1Et(void* qthis, int16_t arg0);
  // proto:  bool QChar::isUpper();
extern void demth_ZNK5QChar7isUpperEv(void* qthis);
  // proto: static unsigned char QChar::combiningClass(uint ucs4);
extern void _ZN5QChar14combiningClassEj(int32_t arg0);
  // proto:  bool QChar::isNonCharacter();
extern void demth_ZNK5QChar14isNonCharacterEv(void* qthis);
  // proto: static bool QChar::isLetterOrNumber(uint ucs4);
extern void demth_ZN5QChar16isLetterOrNumberEj(int32_t arg0);
  // proto: static bool QChar::isDigit(uint ucs4);
extern void demth_ZN5QChar7isDigitEj(int32_t arg0);
  // proto: static bool QChar::isPunct(uint ucs4);
extern void _ZN5QChar7isPunctEj(int32_t arg0);
  // proto:  bool QChar::isTitleCase();
extern void demth_ZNK5QChar11isTitleCaseEv(void* qthis);
  // proto:  bool QChar::isLetter();
extern void demth_ZNK5QChar8isLetterEv(void* qthis);
  // proto:  unsigned char QChar::combiningClass();
extern void demth_ZNK5QChar14combiningClassEv(void* qthis);
  // proto:  bool QChar::isHighSurrogate();
extern void demth_ZNK5QChar15isHighSurrogateEv(void* qthis);
  // proto: static ushort QChar::highSurrogate(uint ucs4);
extern void demth_ZN5QChar13highSurrogateEj(int32_t arg0);
  // proto: static bool QChar::requiresSurrogates(uint ucs4);
extern void demth_ZN5QChar18requiresSurrogatesEj(int32_t arg0);
  // proto:  bool QChar::isLetterOrNumber();
extern void demth_ZNK5QChar16isLetterOrNumberEv(void* qthis);
  // proto:  ushort & QChar::unicode();
extern void demth_ZN5QChar7unicodeEv(void* qthis);
  // proto: static bool QChar::isLowSurrogate(uint ucs4);
extern void demth_ZN5QChar14isLowSurrogateEj(int32_t arg0);
  // proto: static bool QChar::isNumber(uint ucs4);
extern void demth_ZN5QChar8isNumberEj(int32_t arg0);
  // proto:  QChar QChar::toLower();
extern void demth_ZNK5QChar7toLowerEv(void* qthis);
  // proto:  void QChar::QChar(uchar c);
extern void* dector_ZN5QCharC1Eh(unsigned char arg0);
extern void _ZN5QCharC1Eh(void* qthis, unsigned char arg0);
  // proto: static bool QChar::isLetter(uint ucs4);
extern void demth_ZN5QChar8isLetterEj(int32_t arg0);
  // proto:  QChar QChar::toUpper();
extern void demth_ZNK5QChar7toUpperEv(void* qthis);
  // proto: static bool QChar::isSpace(uint ucs4);
extern void demth_ZN5QChar7isSpaceEj(int32_t arg0);
  // proto: static QChar QChar::fromLatin1(char c);
extern void demth_ZN5QChar10fromLatin1Ec(unsigned char arg0);
  // proto: static bool QChar::isSurrogate(uint ucs4);
extern void demth_ZN5QChar11isSurrogateEj(int32_t arg0);
  // proto: static bool QChar::isMark(uint ucs4);
extern void _ZN5QChar6isMarkEj(int32_t arg0);
  // proto: static bool QChar::isNonCharacter(uint ucs4);
extern void demth_ZN5QChar14isNonCharacterEj(int32_t arg0);
  // proto:  char QChar::toLatin1();
extern void demth_ZNK5QChar8toLatin1Ev(void* qthis);
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

// class sizeof(QLatin1Char)=1
type QLatin1Char struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QChar)=2
type QChar struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  ushort QLatin1Char::unicode();
func (this *QLatin1Char) unicode(args ...interface{}) () {
  // unicode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QLatin1Char7unicodeEv
    // invoke: ushort unicode()
    C.demth_ZNK11QLatin1Char7unicodeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLatin1Char", "unicode", args)
  }

}

  // proto:  void QLatin1Char::QLatin1Char(char c);
func NewQLatin1Char(args ...interface{}) QLatin1Char {
  return QLatin1Char{}
}

  // proto:  char QLatin1Char::toLatin1();
func (this *QLatin1Char) toLatin1(args ...interface{}) () {
  // toLatin1()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QLatin1Char8toLatin1Ev
    // invoke: char toLatin1()
    C.demth_ZNK11QLatin1Char8toLatin1Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLatin1Char", "toLatin1", args)
  }

}

  // proto: static uint QChar::toUpper(uint ucs4);
func (this *QChar) toUpper_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "toUpper", args)
  }

}

  // proto:  bool QChar::hasMirrored();
func (this *QChar) hasMirrored(args ...interface{}) () {
  // hasMirrored()
  // hasMirrored(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar11hasMirroredEv
    // invoke: bool hasMirrored()
    C.demth_ZNK5QChar11hasMirroredEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar11hasMirroredEj
    // invoke: bool hasMirrored(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN5QChar11hasMirroredEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "hasMirrored", args)
  }

}

  // proto: static ushort QChar::lowSurrogate(uint ucs4);
func (this *QChar) lowSurrogate_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "lowSurrogate", args)
  }

}

  // proto: static bool QChar::isSymbol(uint ucs4);
func (this *QChar) isSymbol_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isSymbol", args)
  }

}

  // proto:  uchar QChar::cell();
func (this *QChar) cell(args ...interface{}) () {
  // cell()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar4cellEv
    // invoke: uchar cell()
    C.demth_ZNK5QChar4cellEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "cell", args)
  }

}

  // proto: static uint QChar::surrogateToUcs4(QChar high, QChar low);
func (this *QChar) surrogateToUcs4_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "surrogateToUcs4", args)
  }

}

  // proto: static bool QChar::isTitleCase(uint ucs4);
func (this *QChar) isTitleCase_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isTitleCase", args)
  }

}

  // proto:  bool QChar::isNull();
func (this *QChar) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar6isNullEv
    // invoke: bool isNull()
    C.demth_ZNK5QChar6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "isNull", args)
  }

}

  // proto:  int QChar::digitValue();
func (this *QChar) digitValue(args ...interface{}) () {
  // digitValue()
  // digitValue(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar10digitValueEv
    // invoke: int digitValue()
    C.demth_ZNK5QChar10digitValueEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar10digitValueEj
    // invoke: int digitValue(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN5QChar10digitValueEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "digitValue", args)
  }

}

  // proto:  QChar QChar::toTitleCase();
func (this *QChar) toTitleCase(args ...interface{}) () {
  // toTitleCase()
  // toTitleCase(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar11toTitleCaseEv
    // invoke: QChar toTitleCase()
    C.demth_ZNK5QChar11toTitleCaseEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar11toTitleCaseEj
    // invoke: uint toTitleCase(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN5QChar11toTitleCaseEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "toTitleCase", args)
  }

}

  // proto: static bool QChar::isLower(uint ucs4);
func (this *QChar) isLower_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isLower", args)
  }

}

  // proto:  bool QChar::isLowSurrogate();
func (this *QChar) isLowSurrogate(args ...interface{}) () {
  // isLowSurrogate()
  // isLowSurrogate(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar14isLowSurrogateEv
    // invoke: bool isLowSurrogate()
    C.demth_ZNK5QChar14isLowSurrogateEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar14isLowSurrogateEj
    // invoke: bool isLowSurrogate(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.demth_ZN5QChar14isLowSurrogateEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "isLowSurrogate", args)
  }

}

  // proto: static bool QChar::isPrint(uint ucs4);
func (this *QChar) isPrint_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isPrint", args)
  }

}

  // proto:  bool QChar::isSymbol();
func (this *QChar) isSymbol(args ...interface{}) () {
  // isSymbol(uint)
  // isSymbol()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar8isSymbolEj
    // invoke: bool isSymbol(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN5QChar8isSymbolEj(arg0)
  case 1:
    // invoke: _ZNK5QChar8isSymbolEv
    // invoke: bool isSymbol()
    C.demth_ZNK5QChar8isSymbolEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "isSymbol", args)
  }

}

  // proto:  bool QChar::isLower();
func (this *QChar) isLower(args ...interface{}) () {
  // isLower(uint)
  // isLower()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar7isLowerEj
    // invoke: bool isLower(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.demth_ZN5QChar7isLowerEj(arg0)
  case 1:
    // invoke: _ZNK5QChar7isLowerEv
    // invoke: bool isLower()
    C.demth_ZNK5QChar7isLowerEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "isLower", args)
  }

}

  // proto:  QChar QChar::mirroredChar();
func (this *QChar) mirroredChar(args ...interface{}) () {
  // mirroredChar()
  // mirroredChar(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar12mirroredCharEv
    // invoke: QChar mirroredChar()
    C.demth_ZNK5QChar12mirroredCharEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar12mirroredCharEj
    // invoke: uint mirroredChar(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN5QChar12mirroredCharEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "mirroredChar", args)
  }

}

  // proto:  uchar QChar::row();
func (this *QChar) row(args ...interface{}) () {
  // row()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar3rowEv
    // invoke: uchar row()
    C.demth_ZNK5QChar3rowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "row", args)
  }

}

  // proto:  bool QChar::isDigit();
func (this *QChar) isDigit(args ...interface{}) () {
  // isDigit()
  // isDigit(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar7isDigitEv
    // invoke: bool isDigit()
    C.demth_ZNK5QChar7isDigitEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar7isDigitEj
    // invoke: bool isDigit(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.demth_ZN5QChar7isDigitEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "isDigit", args)
  }

}

  // proto: static uint QChar::toTitleCase(uint ucs4);
func (this *QChar) toTitleCase_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "toTitleCase", args)
  }

}

  // proto:  bool QChar::isSurrogate();
func (this *QChar) isSurrogate(args ...interface{}) () {
  // isSurrogate()
  // isSurrogate(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar11isSurrogateEv
    // invoke: bool isSurrogate()
    C.demth_ZNK5QChar11isSurrogateEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar11isSurrogateEj
    // invoke: bool isSurrogate(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.demth_ZN5QChar11isSurrogateEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "isSurrogate", args)
  }

}

  // proto: static bool QChar::hasMirrored(uint ucs4);
func (this *QChar) hasMirrored_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "hasMirrored", args)
  }

}

  // proto:  bool QChar::isNumber();
func (this *QChar) isNumber(args ...interface{}) () {
  // isNumber()
  // isNumber(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar8isNumberEv
    // invoke: bool isNumber()
    C.demth_ZNK5QChar8isNumberEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar8isNumberEj
    // invoke: bool isNumber(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.demth_ZN5QChar8isNumberEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "isNumber", args)
  }

}

  // proto: static bool QChar::isHighSurrogate(uint ucs4);
func (this *QChar) isHighSurrogate_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isHighSurrogate", args)
  }

}

  // proto: static uint QChar::toCaseFolded(uint ucs4);
func (this *QChar) toCaseFolded_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "toCaseFolded", args)
  }

}

  // proto:  bool QChar::isMark();
func (this *QChar) isMark(args ...interface{}) () {
  // isMark()
  // isMark(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar6isMarkEv
    // invoke: bool isMark()
    C.demth_ZNK5QChar6isMarkEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar6isMarkEj
    // invoke: bool isMark(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN5QChar6isMarkEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "isMark", args)
  }

}

  // proto: static uint QChar::toLower(uint ucs4);
func (this *QChar) toLower_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "toLower", args)
  }

}

  // proto: static uint QChar::mirroredChar(uint ucs4);
func (this *QChar) mirroredChar_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "mirroredChar", args)
  }

}

  // proto:  void QChar::setRow(uchar row);
func (this *QChar) setRow(args ...interface{}) () {
  // setRow(uchar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "uchar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar6setRowEh
    // invoke: void setRow(uchar)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    C.demth_ZN5QChar6setRowEh(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QChar", "setRow", args)
  }

}

  // proto: static QString QChar::decomposition(uint ucs4);
func (this *QChar) decomposition_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "decomposition", args)
  }

}

  // proto: static int QChar::digitValue(uint ucs4);
func (this *QChar) digitValue_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "digitValue", args)
  }

}

  // proto:  void QChar::setCell(uchar cell);
func (this *QChar) setCell(args ...interface{}) () {
  // setCell(uchar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "uchar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar7setCellEh
    // invoke: void setCell(uchar)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    C.demth_ZN5QChar7setCellEh(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QChar", "setCell", args)
  }

}

  // proto: static bool QChar::isUpper(uint ucs4);
func (this *QChar) isUpper_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isUpper", args)
  }

}

  // proto:  void QChar::QChar(uchar c, uchar r);
func NewQChar(args ...interface{}) QChar {
  return QChar{}
}

  // proto:  QChar QChar::toCaseFolded();
func (this *QChar) toCaseFolded(args ...interface{}) () {
  // toCaseFolded(uint)
  // toCaseFolded()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar12toCaseFoldedEj
    // invoke: uint toCaseFolded(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN5QChar12toCaseFoldedEj(arg0)
  case 1:
    // invoke: _ZNK5QChar12toCaseFoldedEv
    // invoke: QChar toCaseFolded()
    C.demth_ZNK5QChar12toCaseFoldedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "toCaseFolded", args)
  }

}

  // proto:  bool QChar::isPrint();
func (this *QChar) isPrint(args ...interface{}) () {
  // isPrint(uint)
  // isPrint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar7isPrintEj
    // invoke: bool isPrint(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN5QChar7isPrintEj(arg0)
  case 1:
    // invoke: _ZNK5QChar7isPrintEv
    // invoke: bool isPrint()
    C.demth_ZNK5QChar7isPrintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "isPrint", args)
  }

}

  // proto:  bool QChar::isPunct();
func (this *QChar) isPunct(args ...interface{}) () {
  // isPunct()
  // isPunct(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar7isPunctEv
    // invoke: bool isPunct()
    C.demth_ZNK5QChar7isPunctEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar7isPunctEj
    // invoke: bool isPunct(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN5QChar7isPunctEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "isPunct", args)
  }

}

  // proto:  QString QChar::decomposition();
func (this *QChar) decomposition(args ...interface{}) () {
  // decomposition(uint)
  // decomposition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar13decompositionEj
    // invoke: QString decomposition(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN5QChar13decompositionEj(arg0)
  case 1:
    // invoke: _ZNK5QChar13decompositionEv
    // invoke: QString decomposition()
    C._ZNK5QChar13decompositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "decomposition", args)
  }

}

  // proto:  bool QChar::isSpace();
func (this *QChar) isSpace(args ...interface{}) () {
  // isSpace()
  // isSpace(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar7isSpaceEv
    // invoke: bool isSpace()
    C.demth_ZNK5QChar7isSpaceEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar7isSpaceEj
    // invoke: bool isSpace(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.demth_ZN5QChar7isSpaceEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "isSpace", args)
  }

}

  // proto:  bool QChar::isUpper();
func (this *QChar) isUpper(args ...interface{}) () {
  // isUpper(uint)
  // isUpper()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar7isUpperEj
    // invoke: bool isUpper(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.demth_ZN5QChar7isUpperEj(arg0)
  case 1:
    // invoke: _ZNK5QChar7isUpperEv
    // invoke: bool isUpper()
    C.demth_ZNK5QChar7isUpperEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "isUpper", args)
  }

}

  // proto: static unsigned char QChar::combiningClass(uint ucs4);
func (this *QChar) combiningClass_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "combiningClass", args)
  }

}

  // proto:  bool QChar::isNonCharacter();
func (this *QChar) isNonCharacter(args ...interface{}) () {
  // isNonCharacter()
  // isNonCharacter(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar14isNonCharacterEv
    // invoke: bool isNonCharacter()
    C.demth_ZNK5QChar14isNonCharacterEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar14isNonCharacterEj
    // invoke: bool isNonCharacter(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.demth_ZN5QChar14isNonCharacterEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "isNonCharacter", args)
  }

}

  // proto: static bool QChar::isLetterOrNumber(uint ucs4);
func (this *QChar) isLetterOrNumber_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isLetterOrNumber", args)
  }

}

  // proto: static bool QChar::isDigit(uint ucs4);
func (this *QChar) isDigit_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isDigit", args)
  }

}

  // proto: static bool QChar::isPunct(uint ucs4);
func (this *QChar) isPunct_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isPunct", args)
  }

}

  // proto:  bool QChar::isTitleCase();
func (this *QChar) isTitleCase(args ...interface{}) () {
  // isTitleCase(uint)
  // isTitleCase()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar11isTitleCaseEj
    // invoke: bool isTitleCase(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.demth_ZN5QChar11isTitleCaseEj(arg0)
  case 1:
    // invoke: _ZNK5QChar11isTitleCaseEv
    // invoke: bool isTitleCase()
    C.demth_ZNK5QChar11isTitleCaseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "isTitleCase", args)
  }

}

  // proto:  bool QChar::isLetter();
func (this *QChar) isLetter(args ...interface{}) () {
  // isLetter()
  // isLetter(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar8isLetterEv
    // invoke: bool isLetter()
    C.demth_ZNK5QChar8isLetterEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar8isLetterEj
    // invoke: bool isLetter(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.demth_ZN5QChar8isLetterEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "isLetter", args)
  }

}

  // proto:  unsigned char QChar::combiningClass();
func (this *QChar) combiningClass(args ...interface{}) () {
  // combiningClass(uint)
  // combiningClass()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar14combiningClassEj
    // invoke: unsigned char combiningClass(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN5QChar14combiningClassEj(arg0)
  case 1:
    // invoke: _ZNK5QChar14combiningClassEv
    // invoke: unsigned char combiningClass()
    C.demth_ZNK5QChar14combiningClassEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "combiningClass", args)
  }

}

  // proto:  bool QChar::isHighSurrogate();
func (this *QChar) isHighSurrogate(args ...interface{}) () {
  // isHighSurrogate(uint)
  // isHighSurrogate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar15isHighSurrogateEj
    // invoke: bool isHighSurrogate(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.demth_ZN5QChar15isHighSurrogateEj(arg0)
  case 1:
    // invoke: _ZNK5QChar15isHighSurrogateEv
    // invoke: bool isHighSurrogate()
    C.demth_ZNK5QChar15isHighSurrogateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "isHighSurrogate", args)
  }

}

  // proto: static ushort QChar::highSurrogate(uint ucs4);
func (this *QChar) highSurrogate_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "highSurrogate", args)
  }

}

  // proto: static bool QChar::requiresSurrogates(uint ucs4);
func (this *QChar) requiresSurrogates_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "requiresSurrogates", args)
  }

}

  // proto:  bool QChar::isLetterOrNumber();
func (this *QChar) isLetterOrNumber(args ...interface{}) () {
  // isLetterOrNumber(uint)
  // isLetterOrNumber()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar16isLetterOrNumberEj
    // invoke: bool isLetterOrNumber(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.demth_ZN5QChar16isLetterOrNumberEj(arg0)
  case 1:
    // invoke: _ZNK5QChar16isLetterOrNumberEv
    // invoke: bool isLetterOrNumber()
    C.demth_ZNK5QChar16isLetterOrNumberEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "isLetterOrNumber", args)
  }

}

  // proto:  ushort & QChar::unicode();
func (this *QChar) unicode(args ...interface{}) () {
  // unicode()
  // unicode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar7unicodeEv
    // invoke: ushort & unicode()
    C.demth_ZN5QChar7unicodeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "unicode", args)
  }

}

  // proto: static bool QChar::isLowSurrogate(uint ucs4);
func (this *QChar) isLowSurrogate_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isLowSurrogate", args)
  }

}

  // proto: static bool QChar::isNumber(uint ucs4);
func (this *QChar) isNumber_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isNumber", args)
  }

}

  // proto:  QChar QChar::toLower();
func (this *QChar) toLower(args ...interface{}) () {
  // toLower(uint)
  // toLower()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar7toLowerEj
    // invoke: uint toLower(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN5QChar7toLowerEj(arg0)
  case 1:
    // invoke: _ZNK5QChar7toLowerEv
    // invoke: QChar toLower()
    C.demth_ZNK5QChar7toLowerEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "toLower", args)
  }

}

  // proto: static bool QChar::isLetter(uint ucs4);
func (this *QChar) isLetter_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isLetter", args)
  }

}

  // proto:  QChar QChar::toUpper();
func (this *QChar) toUpper(args ...interface{}) () {
  // toUpper(uint)
  // toUpper()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar7toUpperEj
    // invoke: uint toUpper(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN5QChar7toUpperEj(arg0)
  case 1:
    // invoke: _ZNK5QChar7toUpperEv
    // invoke: QChar toUpper()
    C.demth_ZNK5QChar7toUpperEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "toUpper", args)
  }

}

  // proto: static bool QChar::isSpace(uint ucs4);
func (this *QChar) isSpace_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isSpace", args)
  }

}

  // proto: static QChar QChar::fromLatin1(char c);
func (this *QChar) fromLatin1_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "fromLatin1", args)
  }

}

  // proto: static bool QChar::isSurrogate(uint ucs4);
func (this *QChar) isSurrogate_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isSurrogate", args)
  }

}

  // proto: static bool QChar::isMark(uint ucs4);
func (this *QChar) isMark_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isMark", args)
  }

}

  // proto: static bool QChar::isNonCharacter(uint ucs4);
func (this *QChar) isNonCharacter_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isNonCharacter", args)
  }

}

  // proto:  char QChar::toLatin1();
func (this *QChar) toLatin1(args ...interface{}) () {
  // toLatin1()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar8toLatin1Ev
    // invoke: char toLatin1()
    C.demth_ZNK5QChar8toLatin1Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "toLatin1", args)
  }

}

// <= body block end

