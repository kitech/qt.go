package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
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
extern void _ZNK11QLatin1Char7unicodeEv(void* qthis);
  // proto:  void QLatin1Char::QLatin1Char(char c);
extern void* dector_ZN11QLatin1CharC1Ec(char arg0);
extern void _ZN11QLatin1CharC1Ec(void* qthis, char arg0);
  // proto:  char QLatin1Char::toLatin1();
extern void _ZNK11QLatin1Char8toLatin1Ev(void* qthis);
  // proto: static uint QChar::toUpper(uint ucs4);
extern void _ZN5QChar7toUpperEj(unsigned int arg0);
  // proto:  bool QChar::hasMirrored();
extern void demth_ZNK5QChar11hasMirroredEv(void* qthis);
  // proto: static ushort QChar::lowSurrogate(uint ucs4);
extern void demth_ZN5QChar12lowSurrogateEj(unsigned int arg0);
  // proto: static bool QChar::isSymbol(uint ucs4);
extern void _ZN5QChar8isSymbolEj(unsigned int arg0);
  // proto:  uchar QChar::cell();
extern void _ZNK5QChar4cellEv(void* qthis);
  // proto: static uint QChar::surrogateToUcs4(QChar high, QChar low);
extern void demth_ZN5QChar15surrogateToUcs4ES_S_(void* arg0, void* arg1);
  // proto: static bool QChar::isTitleCase(uint ucs4);
extern void demth_ZN5QChar11isTitleCaseEj(unsigned int arg0);
  // proto:  bool QChar::isNull();
extern void _ZNK5QChar6isNullEv(void* qthis);
  // proto:  int QChar::digitValue();
extern void demth_ZNK5QChar10digitValueEv(void* qthis);
  // proto:  QChar QChar::toTitleCase();
extern void demth_ZNK5QChar11toTitleCaseEv(void* qthis);
  // proto: static bool QChar::isLower(uint ucs4);
extern void demth_ZN5QChar7isLowerEj(unsigned int arg0);
  // proto:  bool QChar::isLowSurrogate();
extern void _ZNK5QChar14isLowSurrogateEv(void* qthis);
  // proto: static bool QChar::isPrint(uint ucs4);
extern void _ZN5QChar7isPrintEj(unsigned int arg0);
  // proto:  bool QChar::isSymbol();
extern void demth_ZNK5QChar8isSymbolEv(void* qthis);
  // proto:  bool QChar::isLower();
extern void _ZNK5QChar7isLowerEv(void* qthis);
  // proto:  QChar QChar::mirroredChar();
extern void demth_ZNK5QChar12mirroredCharEv(void* qthis);
  // proto:  uchar QChar::row();
extern void _ZNK5QChar3rowEv(void* qthis);
  // proto:  bool QChar::isDigit();
extern void _ZNK5QChar7isDigitEv(void* qthis);
  // proto: static uint QChar::toTitleCase(uint ucs4);
extern void _ZN5QChar11toTitleCaseEj(unsigned int arg0);
  // proto:  bool QChar::isSurrogate();
extern void _ZNK5QChar11isSurrogateEv(void* qthis);
  // proto: static bool QChar::hasMirrored(uint ucs4);
extern void _ZN5QChar11hasMirroredEj(unsigned int arg0);
  // proto:  bool QChar::isNumber();
extern void _ZNK5QChar8isNumberEv(void* qthis);
  // proto: static bool QChar::isHighSurrogate(uint ucs4);
extern void demth_ZN5QChar15isHighSurrogateEj(unsigned int arg0);
  // proto: static uint QChar::toCaseFolded(uint ucs4);
extern void _ZN5QChar12toCaseFoldedEj(unsigned int arg0);
  // proto:  bool QChar::isMark();
extern void demth_ZNK5QChar6isMarkEv(void* qthis);
  // proto: static uint QChar::surrogateToUcs4(ushort high, ushort low);
extern void demth_ZN5QChar15surrogateToUcs4Ett(unsigned short arg0, unsigned short arg1);
  // proto: static uint QChar::toLower(uint ucs4);
extern void _ZN5QChar7toLowerEj(unsigned int arg0);
  // proto: static uint QChar::mirroredChar(uint ucs4);
extern void _ZN5QChar12mirroredCharEj(unsigned int arg0);
  // proto:  void QChar::setRow(uchar row);
extern void demth_ZN5QChar6setRowEh(void* qthis, unsigned char arg0);
  // proto: static QString QChar::decomposition(uint ucs4);
extern void _ZN5QChar13decompositionEj(unsigned int arg0);
  // proto: static int QChar::digitValue(uint ucs4);
extern void _ZN5QChar10digitValueEj(unsigned int arg0);
  // proto:  void QChar::setCell(uchar cell);
extern void demth_ZN5QChar7setCellEh(void* qthis, unsigned char arg0);
  // proto: static bool QChar::isUpper(uint ucs4);
extern void demth_ZN5QChar7isUpperEj(unsigned int arg0);
  // proto:  void QChar::QChar(uchar c, uchar r);
extern void* dector_ZN5QCharC1Ehh(unsigned char arg0, unsigned char arg1);
extern void _ZN5QCharC1Ehh(void* qthis, unsigned char arg0, unsigned char arg1);
  // proto:  QChar QChar::toCaseFolded();
extern void demth_ZNK5QChar12toCaseFoldedEv(void* qthis);
  // proto:  bool QChar::isPrint();
extern void demth_ZNK5QChar7isPrintEv(void* qthis);
  // proto:  void QChar::QChar(char c);
extern void* dector_ZN5QCharC1Ec(char arg0);
extern void _ZN5QCharC1Ec(void* qthis, char arg0);
  // proto:  bool QChar::isPunct();
extern void demth_ZNK5QChar7isPunctEv(void* qthis);
  // proto:  QString QChar::decomposition();
extern void _ZNK5QChar13decompositionEv(void* qthis);
  // proto:  void QChar::QChar(uint rc);
extern void* dector_ZN5QCharC1Ej(unsigned int arg0);
extern void _ZN5QCharC1Ej(void* qthis, unsigned int arg0);
  // proto:  void QChar::QChar(int rc);
extern void* dector_ZN5QCharC1Ei(int arg0);
extern void _ZN5QCharC1Ei(void* qthis, int arg0);
  // proto:  bool QChar::isSpace();
extern void _ZNK5QChar7isSpaceEv(void* qthis);
  // proto:  void QChar::QChar(short rc);
extern void* dector_ZN5QCharC1Es(short arg0);
extern void _ZN5QCharC1Es(void* qthis, short arg0);
  // proto:  void QChar::QChar();
extern void* dector_ZN5QCharC1Ev();
extern void _ZN5QCharC1Ev(void* qthis);
  // proto:  void QChar::QChar(ushort rc);
extern void* dector_ZN5QCharC1Et(unsigned short arg0);
extern void _ZN5QCharC1Et(void* qthis, unsigned short arg0);
  // proto:  bool QChar::isUpper();
extern void _ZNK5QChar7isUpperEv(void* qthis);
  // proto: static unsigned char QChar::combiningClass(uint ucs4);
extern void _ZN5QChar14combiningClassEj(unsigned int arg0);
  // proto:  bool QChar::isNonCharacter();
extern void _ZNK5QChar14isNonCharacterEv(void* qthis);
  // proto: static bool QChar::isLetterOrNumber(uint ucs4);
extern void demth_ZN5QChar16isLetterOrNumberEj(unsigned int arg0);
  // proto: static bool QChar::isDigit(uint ucs4);
extern void demth_ZN5QChar7isDigitEj(unsigned int arg0);
  // proto: static bool QChar::isPunct(uint ucs4);
extern void _ZN5QChar7isPunctEj(unsigned int arg0);
  // proto:  bool QChar::isTitleCase();
extern void _ZNK5QChar11isTitleCaseEv(void* qthis);
  // proto:  bool QChar::isLetter();
extern void _ZNK5QChar8isLetterEv(void* qthis);
  // proto:  unsigned char QChar::combiningClass();
extern void demth_ZNK5QChar14combiningClassEv(void* qthis);
  // proto:  bool QChar::isHighSurrogate();
extern void _ZNK5QChar15isHighSurrogateEv(void* qthis);
  // proto: static ushort QChar::highSurrogate(uint ucs4);
extern void demth_ZN5QChar13highSurrogateEj(unsigned int arg0);
  // proto: static bool QChar::requiresSurrogates(uint ucs4);
extern void demth_ZN5QChar18requiresSurrogatesEj(unsigned int arg0);
  // proto:  bool QChar::isLetterOrNumber();
extern void _ZNK5QChar16isLetterOrNumberEv(void* qthis);
  // proto:  ushort & QChar::unicode();
extern void demth_ZN5QChar7unicodeEv(void* qthis);
  // proto: static bool QChar::isLowSurrogate(uint ucs4);
extern void demth_ZN5QChar14isLowSurrogateEj(unsigned int arg0);
  // proto: static bool QChar::isNumber(uint ucs4);
extern void demth_ZN5QChar8isNumberEj(unsigned int arg0);
  // proto:  QChar QChar::toLower();
extern void demth_ZNK5QChar7toLowerEv(void* qthis);
  // proto:  void QChar::QChar(uchar c);
extern void* dector_ZN5QCharC1Eh(unsigned char arg0);
extern void _ZN5QCharC1Eh(void* qthis, unsigned char arg0);
  // proto: static bool QChar::isLetter(uint ucs4);
extern void demth_ZN5QChar8isLetterEj(unsigned int arg0);
  // proto:  QChar QChar::toUpper();
extern void demth_ZNK5QChar7toUpperEv(void* qthis);
  // proto: static bool QChar::isSpace(uint ucs4);
extern void demth_ZN5QChar7isSpaceEj(unsigned int arg0);
  // proto: static QChar QChar::fromLatin1(char c);
extern void _ZN5QChar10fromLatin1Ec(char arg0);
  // proto: static bool QChar::isSurrogate(uint ucs4);
extern void demth_ZN5QChar11isSurrogateEj(unsigned int arg0);
  // proto: static bool QChar::isMark(uint ucs4);
extern void _ZN5QChar6isMarkEj(unsigned int arg0);
  // proto: static bool QChar::isNonCharacter(uint ucs4);
extern void demth_ZN5QChar14isNonCharacterEj(unsigned int arg0);
  // proto:  char QChar::toLatin1();
extern void _ZNK5QChar8toLatin1Ev(void* qthis);
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
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QChar)=2
type QChar struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
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
  case 1:
    // invoke: _ZN5QChar11hasMirroredEj
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
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
  case 1:
    // invoke: _ZN5QChar10digitValueEj
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
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
  case 1:
    // invoke: _ZN5QChar11toTitleCaseEj
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
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
  case 1:
    // invoke: _ZN5QChar14isLowSurrogateEj
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK5QChar8isSymbolEv
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK5QChar7isLowerEv
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
  case 1:
    // invoke: _ZN5QChar12mirroredCharEj
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
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
  case 1:
    // invoke: _ZN5QChar7isDigitEj
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
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
  case 1:
    // invoke: _ZN5QChar11isSurrogateEj
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
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
  case 1:
    // invoke: _ZN5QChar8isNumberEj
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
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
  case 1:
    // invoke: _ZN5QChar6isMarkEj
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
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
    var arg0 = C.char(args[0].(byte))
    if false {fmt.Println(arg0)}
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
    var arg0 = C.char(args[0].(byte))
    if false {fmt.Println(arg0)}
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK5QChar12toCaseFoldedEv
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK5QChar7isPrintEv
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
  case 1:
    // invoke: _ZN5QChar7isPunctEj
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK5QChar13decompositionEv
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
  case 1:
    // invoke: _ZN5QChar7isSpaceEj
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK5QChar7isUpperEv
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
  case 1:
    // invoke: _ZN5QChar14isNonCharacterEj
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK5QChar11isTitleCaseEv
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
  case 1:
    // invoke: _ZN5QChar8isLetterEj
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK5QChar14combiningClassEv
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK5QChar15isHighSurrogateEv
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK5QChar16isLetterOrNumberEv
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
    // invoke: _ZNK5QChar7unicodeEv
  case 1:
    // invoke: _ZN5QChar7unicodeEv
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK5QChar7toLowerEv
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK5QChar7toUpperEv
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
  default:
    qtrt.ErrorResolve("QChar", "toLatin1", args)
  }

}

// <= body block end

