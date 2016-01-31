package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QLatin1Char::QLatin1Char(char c);
extern void C_ZN11QLatin1CharC2Ec(void* qthis, unsigned char arg0); // 1
  // proto:  char QLatin1Char::toLatin1();
extern void C_ZNK11QLatin1Char8toLatin1Ev(void* qthis); // 2
  // proto:  ushort QLatin1Char::unicode();
extern void C_ZNK11QLatin1Char7unicodeEv(void* qthis); // 2
  // proto: static bool QChar::isLetterOrNumber(uint ucs4);
extern void C_ZN5QChar16isLetterOrNumberEj(int32_t arg0); // 2
  // proto:  bool QChar::isLetterOrNumber();
extern void C_ZNK5QChar16isLetterOrNumberEv(void* qthis); // 2
  // proto:  QChar::UnicodeVersion QChar::unicodeVersion();
extern void C_ZNK5QChar14unicodeVersionEv(void* qthis); // 2
  // proto: static QChar::UnicodeVersion QChar::unicodeVersion(uint ucs4);
extern void C_ZN5QChar14unicodeVersionEj(int32_t arg0); // 4
  // proto: static bool QChar::hasMirrored(uint ucs4);
extern void C_ZN5QChar11hasMirroredEj(int32_t arg0); // 4
  // proto:  bool QChar::hasMirrored();
extern void C_ZNK5QChar11hasMirroredEv(void* qthis); // 2
  // proto:  char QChar::toLatin1();
extern void C_ZNK5QChar8toLatin1Ev(void* qthis); // 2
  // proto:  bool QChar::isTitleCase();
extern void C_ZNK5QChar11isTitleCaseEv(void* qthis); // 2
  // proto: static bool QChar::isTitleCase(uint ucs4);
extern void C_ZN5QChar11isTitleCaseEj(int32_t arg0); // 2
  // proto: static QChar::UnicodeVersion QChar::currentUnicodeVersion();
extern void C_ZN5QChar21currentUnicodeVersionEv(); // 4
  // proto:  void QChar::setCell(uchar cell);
extern void C_ZN5QChar7setCellEh(void* qthis, unsigned char arg0); // 2
  // proto:  ushort & QChar::unicode();
extern void C_ZN5QChar7unicodeEv(void* qthis); // 2
  // proto: static ushort QChar::lowSurrogate(uint ucs4);
extern void C_ZN5QChar12lowSurrogateEj(int32_t arg0); // 2
  // proto: static bool QChar::isSpace(uint ucs4);
extern void C_ZN5QChar7isSpaceEj(int32_t arg0); // 2
  // proto:  bool QChar::isSpace();
extern void C_ZNK5QChar7isSpaceEv(void* qthis); // 2
  // proto:  uchar QChar::row();
extern void C_ZNK5QChar3rowEv(void* qthis); // 2
  // proto:  QChar::Category QChar::category();
extern void C_ZNK5QChar8categoryEv(void* qthis); // 2
  // proto: static QChar::Category QChar::category(uint ucs4);
extern void C_ZN5QChar8categoryEj(int32_t arg0); // 4
  // proto:  QChar::Decomposition QChar::decompositionTag();
extern void C_ZNK5QChar16decompositionTagEv(void* qthis); // 2
  // proto: static QChar::Decomposition QChar::decompositionTag(uint ucs4);
extern void C_ZN5QChar16decompositionTagEj(int32_t arg0); // 4
  // proto:  bool QChar::isHighSurrogate();
extern void C_ZNK5QChar15isHighSurrogateEv(void* qthis); // 2
  // proto: static bool QChar::isHighSurrogate(uint ucs4);
extern void C_ZN5QChar15isHighSurrogateEj(int32_t arg0); // 2
  // proto: static QChar::Script QChar::script(uint ucs4);
extern void C_ZN5QChar6scriptEj(int32_t arg0); // 4
  // proto:  QChar::Script QChar::script();
extern void C_ZNK5QChar6scriptEv(void* qthis); // 2
  // proto: static QChar QChar::fromLatin1(char c);
extern void C_ZN5QChar10fromLatin1Ec(unsigned char arg0); // 2
  // proto: static QChar::JoiningType QChar::joiningType(uint ucs4);
extern void C_ZN5QChar11joiningTypeEj(int32_t arg0); // 4
  // proto:  QChar::JoiningType QChar::joiningType();
extern void C_ZNK5QChar11joiningTypeEv(void* qthis); // 2
  // proto:  uchar QChar::cell();
extern void C_ZNK5QChar4cellEv(void* qthis); // 2
  // proto:  int QChar::digitValue();
extern void C_ZNK5QChar10digitValueEv(void* qthis); // 2
  // proto: static int QChar::digitValue(uint ucs4);
extern void C_ZN5QChar10digitValueEj(int32_t arg0); // 4
  // proto: static bool QChar::requiresSurrogates(uint ucs4);
extern void C_ZN5QChar18requiresSurrogatesEj(int32_t arg0); // 2
  // proto:  QString QChar::decomposition();
extern void C_ZNK5QChar13decompositionEv(void* qthis); // 4
  // proto: static QString QChar::decomposition(uint ucs4);
extern void C_ZN5QChar13decompositionEj(int32_t arg0); // 4
  // proto: static ushort QChar::highSurrogate(uint ucs4);
extern void C_ZN5QChar13highSurrogateEj(int32_t arg0); // 2
  // proto:  bool QChar::isSurrogate();
extern void C_ZNK5QChar11isSurrogateEv(void* qthis); // 2
  // proto: static bool QChar::isSurrogate(uint ucs4);
extern void C_ZN5QChar11isSurrogateEj(int32_t arg0); // 2
  // proto: static QChar::Direction QChar::direction(uint ucs4);
extern void C_ZN5QChar9directionEj(int32_t arg0); // 4
  // proto:  QChar::Direction QChar::direction();
extern void C_ZNK5QChar9directionEv(void* qthis); // 2
  // proto: static uint QChar::mirroredChar(uint ucs4);
extern void C_ZN5QChar12mirroredCharEj(int32_t arg0); // 4
  // proto:  QChar QChar::mirroredChar();
extern void C_ZNK5QChar12mirroredCharEv(void* qthis); // 2
  // proto:  unsigned char QChar::combiningClass();
extern void C_ZNK5QChar14combiningClassEv(void* qthis); // 2
  // proto: static unsigned char QChar::combiningClass(uint ucs4);
extern void C_ZN5QChar14combiningClassEj(int32_t arg0); // 4
  // proto: static bool QChar::isMark(uint ucs4);
extern void C_ZN5QChar6isMarkEj(int32_t arg0); // 4
  // proto:  bool QChar::isMark();
extern void C_ZNK5QChar6isMarkEv(void* qthis); // 2
  // proto: static bool QChar::isLower(uint ucs4);
extern void C_ZN5QChar7isLowerEj(int32_t arg0); // 2
  // proto:  bool QChar::isLower();
extern void C_ZNK5QChar7isLowerEv(void* qthis); // 2
  // proto: static bool QChar::isNumber(uint ucs4);
extern void C_ZN5QChar8isNumberEj(int32_t arg0); // 2
  // proto:  bool QChar::isNumber();
extern void C_ZNK5QChar8isNumberEv(void* qthis); // 2
  // proto: static uint QChar::toLower(uint ucs4);
extern void C_ZN5QChar7toLowerEj(int32_t arg0); // 4
  // proto:  QChar QChar::toLower();
extern void C_ZNK5QChar7toLowerEv(void* qthis); // 2
  // proto:  QChar QChar::toTitleCase();
extern void C_ZNK5QChar11toTitleCaseEv(void* qthis); // 2
  // proto: static uint QChar::toTitleCase(uint ucs4);
extern void C_ZN5QChar11toTitleCaseEj(int32_t arg0); // 4
  // proto: static bool QChar::isLowSurrogate(uint ucs4);
extern void C_ZN5QChar14isLowSurrogateEj(int32_t arg0); // 2
  // proto:  bool QChar::isLowSurrogate();
extern void C_ZNK5QChar14isLowSurrogateEv(void* qthis); // 2
  // proto:  bool QChar::isUpper();
extern void C_ZNK5QChar7isUpperEv(void* qthis); // 2
  // proto: static bool QChar::isUpper(uint ucs4);
extern void C_ZN5QChar7isUpperEj(int32_t arg0); // 2
  // proto:  void QChar::QChar(char c);
extern void C_ZN5QCharC2Ec(void* qthis, unsigned char arg0); // 1
  // proto:  void QChar::QChar(uint rc);
extern void C_ZN5QCharC2Ej(void* qthis, int32_t arg0); // 1
  // proto:  void QChar::QChar(int rc);
extern void C_ZN5QCharC2Ei(void* qthis, int32_t arg0); // 1
  // proto:  void QChar::QChar(uchar c);
extern void C_ZN5QCharC2Eh(void* qthis, unsigned char arg0); // 1
  // proto:  void QChar::QChar(short rc);
extern void C_ZN5QCharC2Es(void* qthis, int16_t arg0); // 1
  // proto:  void QChar::QChar(uchar c, uchar r);
extern void C_ZN5QCharC2Ehh(void* qthis, unsigned char arg0, unsigned char arg1); // 1
  // proto:  void QChar::QChar();
extern void C_ZN5QCharC2Ev(void* qthis); // 1
  // proto:  void QChar::QChar(ushort rc);
extern void C_ZN5QCharC2Et(void* qthis, int16_t arg0); // 1
  // proto: static uint QChar::toCaseFolded(uint ucs4);
extern void C_ZN5QChar12toCaseFoldedEj(int32_t arg0); // 4
  // proto:  QChar QChar::toCaseFolded();
extern void C_ZNK5QChar12toCaseFoldedEv(void* qthis); // 2
  // proto:  bool QChar::isLetter();
extern void C_ZNK5QChar8isLetterEv(void* qthis); // 2
  // proto: static bool QChar::isLetter(uint ucs4);
extern void C_ZN5QChar8isLetterEj(int32_t arg0); // 2
  // proto: static bool QChar::isNonCharacter(uint ucs4);
extern void C_ZN5QChar14isNonCharacterEj(int32_t arg0); // 2
  // proto:  bool QChar::isNonCharacter();
extern void C_ZNK5QChar14isNonCharacterEv(void* qthis); // 2
  // proto:  bool QChar::isNull();
extern void C_ZNK5QChar6isNullEv(void* qthis); // 2
  // proto:  bool QChar::isPrint();
extern void C_ZNK5QChar7isPrintEv(void* qthis); // 2
  // proto: static bool QChar::isPrint(uint ucs4);
extern void C_ZN5QChar7isPrintEj(int32_t arg0); // 4
  // proto:  QChar QChar::toUpper();
extern void C_ZNK5QChar7toUpperEv(void* qthis); // 2
  // proto: static uint QChar::toUpper(uint ucs4);
extern void C_ZN5QChar7toUpperEj(int32_t arg0); // 4
  // proto: static uint QChar::surrogateToUcs4(ushort high, ushort low);
extern void C_ZN5QChar15surrogateToUcs4Ett(int16_t arg0, int16_t arg1); // 2
  // proto: static uint QChar::surrogateToUcs4(QChar high, QChar low);
extern void C_ZN5QChar15surrogateToUcs4ES_S_(void* arg0, void* arg1); // 2
  // proto: static bool QChar::isDigit(uint ucs4);
extern void C_ZN5QChar7isDigitEj(int32_t arg0); // 2
  // proto:  bool QChar::isDigit();
extern void C_ZNK5QChar7isDigitEv(void* qthis); // 2
  // proto:  void QChar::setRow(uchar row);
extern void C_ZN5QChar6setRowEh(void* qthis, unsigned char arg0); // 2
  // proto:  bool QChar::isPunct();
extern void C_ZNK5QChar7isPunctEv(void* qthis); // 2
  // proto: static bool QChar::isPunct(uint ucs4);
extern void C_ZN5QChar7isPunctEj(int32_t arg0); // 4
  // proto: static bool QChar::isSymbol(uint ucs4);
extern void C_ZN5QChar8isSymbolEj(int32_t arg0); // 4
  // proto:  bool QChar::isSymbol();
extern void C_ZNK5QChar8isSymbolEv(void* qthis); // 2
  // proto:  QChar::Joining QChar::joining();
extern void C_ZNK5QChar7joiningEv(void* qthis); // 2
  // proto: static QChar::Joining QChar::joining(uint ucs4);
extern void C_ZN5QChar7joiningEj(int32_t arg0); // 4
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

// QLatin1Char(char)
func NewQLatin1Char(args ...interface{}) QLatin1Char {
  // QLatin1Char(char)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "char"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QLatin1CharC1Ec
    // invoke: void QLatin1Char(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QLatin1CharC2Ec(qthis, arg0)
  default:
    qtrt.ErrorResolve("QLatin1Char", "QLatin1Char", args)
  }

  return QLatin1Char{}
}

// toLatin1()
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
    C.C_ZNK11QLatin1Char8toLatin1Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLatin1Char", "toLatin1", args)
  }

}

// unicode()
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
    C.C_ZNK11QLatin1Char7unicodeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLatin1Char", "unicode", args)
  }

}

// isLetterOrNumber(uint)
func (this *QChar) isLetterOrNumber_s(args ...interface{}) () {
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
    C.C_ZN5QChar16isLetterOrNumberEj(arg0)
  case 1:
    // invoke: _ZNK5QChar16isLetterOrNumberEv
    // invoke: bool isLetterOrNumber()
    C.C_ZNK5QChar16isLetterOrNumberEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "isLetterOrNumber", args)
  }

}

// unicodeVersion()
func (this *QChar) unicodeVersion(args ...interface{}) () {
  // unicodeVersion()
  // unicodeVersion(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar14unicodeVersionEv
    // invoke: QChar::UnicodeVersion unicodeVersion()
    C.C_ZNK5QChar14unicodeVersionEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar14unicodeVersionEj
    // invoke: QChar::UnicodeVersion unicodeVersion(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar14unicodeVersionEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "unicodeVersion", args)
  }

}

// hasMirrored(uint)
func (this *QChar) hasMirrored_s(args ...interface{}) () {
  // hasMirrored(uint)
  // hasMirrored()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar11hasMirroredEj
    // invoke: bool hasMirrored(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar11hasMirroredEj(arg0)
  case 1:
    // invoke: _ZNK5QChar11hasMirroredEv
    // invoke: bool hasMirrored()
    C.C_ZNK5QChar11hasMirroredEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "hasMirrored", args)
  }

}

// toLatin1()
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
    C.C_ZNK5QChar8toLatin1Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "toLatin1", args)
  }

}

// isTitleCase()
func (this *QChar) isTitleCase(args ...interface{}) () {
  // isTitleCase()
  // isTitleCase(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar11isTitleCaseEv
    // invoke: bool isTitleCase()
    C.C_ZNK5QChar11isTitleCaseEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar11isTitleCaseEj
    // invoke: bool isTitleCase(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar11isTitleCaseEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "isTitleCase", args)
  }

}

// currentUnicodeVersion()
func (this *QChar) currentUnicodeVersion_s(args ...interface{}) () {
  // currentUnicodeVersion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar21currentUnicodeVersionEv
    // invoke: QChar::UnicodeVersion currentUnicodeVersion()
    C.C_ZN5QChar21currentUnicodeVersionEv()
  default:
    qtrt.ErrorResolve("QChar", "currentUnicodeVersion", args)
  }

}

// setCell(uchar)
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
    C.C_ZN5QChar7setCellEh(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QChar", "setCell", args)
  }

}

// unicode()
func (this *QChar) unicode(args ...interface{}) () {
  // unicode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar7unicodeEv
    // invoke: ushort & unicode()
    C.C_ZN5QChar7unicodeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "unicode", args)
  }

}

// lowSurrogate(uint)
func (this *QChar) lowSurrogate_s(args ...interface{}) () {
  // lowSurrogate(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar12lowSurrogateEj
    // invoke: ushort lowSurrogate(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar12lowSurrogateEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "lowSurrogate", args)
  }

}

// isSpace(uint)
func (this *QChar) isSpace_s(args ...interface{}) () {
  // isSpace(uint)
  // isSpace()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar7isSpaceEj
    // invoke: bool isSpace(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar7isSpaceEj(arg0)
  case 1:
    // invoke: _ZNK5QChar7isSpaceEv
    // invoke: bool isSpace()
    C.C_ZNK5QChar7isSpaceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "isSpace", args)
  }

}

// row()
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
    C.C_ZNK5QChar3rowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "row", args)
  }

}

// category()
func (this *QChar) category(args ...interface{}) () {
  // category()
  // category(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar8categoryEv
    // invoke: QChar::Category category()
    C.C_ZNK5QChar8categoryEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar8categoryEj
    // invoke: QChar::Category category(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar8categoryEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "category", args)
  }

}

// decompositionTag()
func (this *QChar) decompositionTag(args ...interface{}) () {
  // decompositionTag()
  // decompositionTag(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar16decompositionTagEv
    // invoke: QChar::Decomposition decompositionTag()
    C.C_ZNK5QChar16decompositionTagEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar16decompositionTagEj
    // invoke: QChar::Decomposition decompositionTag(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar16decompositionTagEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "decompositionTag", args)
  }

}

// isHighSurrogate()
func (this *QChar) isHighSurrogate(args ...interface{}) () {
  // isHighSurrogate()
  // isHighSurrogate(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar15isHighSurrogateEv
    // invoke: bool isHighSurrogate()
    C.C_ZNK5QChar15isHighSurrogateEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar15isHighSurrogateEj
    // invoke: bool isHighSurrogate(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar15isHighSurrogateEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "isHighSurrogate", args)
  }

}

// script(uint)
func (this *QChar) script_s(args ...interface{}) () {
  // script(uint)
  // script()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar6scriptEj
    // invoke: QChar::Script script(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar6scriptEj(arg0)
  case 1:
    // invoke: _ZNK5QChar6scriptEv
    // invoke: QChar::Script script()
    C.C_ZNK5QChar6scriptEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "script", args)
  }

}

// fromLatin1(char)
func (this *QChar) fromLatin1_s(args ...interface{}) () {
  // fromLatin1(char)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "char"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar10fromLatin1Ec
    // invoke: QChar fromLatin1(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar10fromLatin1Ec(arg0)
  default:
    qtrt.ErrorResolve("QChar", "fromLatin1", args)
  }

}

// joiningType(uint)
func (this *QChar) joiningType_s(args ...interface{}) () {
  // joiningType(uint)
  // joiningType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar11joiningTypeEj
    // invoke: QChar::JoiningType joiningType(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar11joiningTypeEj(arg0)
  case 1:
    // invoke: _ZNK5QChar11joiningTypeEv
    // invoke: QChar::JoiningType joiningType()
    C.C_ZNK5QChar11joiningTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "joiningType", args)
  }

}

// cell()
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
    C.C_ZNK5QChar4cellEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "cell", args)
  }

}

// digitValue()
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
    C.C_ZNK5QChar10digitValueEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar10digitValueEj
    // invoke: int digitValue(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar10digitValueEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "digitValue", args)
  }

}

// requiresSurrogates(uint)
func (this *QChar) requiresSurrogates_s(args ...interface{}) () {
  // requiresSurrogates(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar18requiresSurrogatesEj
    // invoke: bool requiresSurrogates(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar18requiresSurrogatesEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "requiresSurrogates", args)
  }

}

// decomposition()
func (this *QChar) decomposition(args ...interface{}) () {
  // decomposition()
  // decomposition(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar13decompositionEv
    // invoke: QString decomposition()
    C.C_ZNK5QChar13decompositionEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar13decompositionEj
    // invoke: QString decomposition(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar13decompositionEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "decomposition", args)
  }

}

// highSurrogate(uint)
func (this *QChar) highSurrogate_s(args ...interface{}) () {
  // highSurrogate(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar13highSurrogateEj
    // invoke: ushort highSurrogate(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar13highSurrogateEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "highSurrogate", args)
  }

}

// isSurrogate()
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
    C.C_ZNK5QChar11isSurrogateEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar11isSurrogateEj
    // invoke: bool isSurrogate(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar11isSurrogateEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "isSurrogate", args)
  }

}

// direction(uint)
func (this *QChar) direction_s(args ...interface{}) () {
  // direction(uint)
  // direction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar9directionEj
    // invoke: QChar::Direction direction(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar9directionEj(arg0)
  case 1:
    // invoke: _ZNK5QChar9directionEv
    // invoke: QChar::Direction direction()
    C.C_ZNK5QChar9directionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "direction", args)
  }

}

// mirroredChar(uint)
func (this *QChar) mirroredChar_s(args ...interface{}) () {
  // mirroredChar(uint)
  // mirroredChar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar12mirroredCharEj
    // invoke: uint mirroredChar(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar12mirroredCharEj(arg0)
  case 1:
    // invoke: _ZNK5QChar12mirroredCharEv
    // invoke: QChar mirroredChar()
    C.C_ZNK5QChar12mirroredCharEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "mirroredChar", args)
  }

}

// combiningClass()
func (this *QChar) combiningClass(args ...interface{}) () {
  // combiningClass()
  // combiningClass(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar14combiningClassEv
    // invoke: unsigned char combiningClass()
    C.C_ZNK5QChar14combiningClassEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar14combiningClassEj
    // invoke: unsigned char combiningClass(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar14combiningClassEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "combiningClass", args)
  }

}

// isMark(uint)
func (this *QChar) isMark_s(args ...interface{}) () {
  // isMark(uint)
  // isMark()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar6isMarkEj
    // invoke: bool isMark(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar6isMarkEj(arg0)
  case 1:
    // invoke: _ZNK5QChar6isMarkEv
    // invoke: bool isMark()
    C.C_ZNK5QChar6isMarkEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "isMark", args)
  }

}

// isLower(uint)
func (this *QChar) isLower_s(args ...interface{}) () {
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
    C.C_ZN5QChar7isLowerEj(arg0)
  case 1:
    // invoke: _ZNK5QChar7isLowerEv
    // invoke: bool isLower()
    C.C_ZNK5QChar7isLowerEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "isLower", args)
  }

}

// isNumber(uint)
func (this *QChar) isNumber_s(args ...interface{}) () {
  // isNumber(uint)
  // isNumber()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar8isNumberEj
    // invoke: bool isNumber(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar8isNumberEj(arg0)
  case 1:
    // invoke: _ZNK5QChar8isNumberEv
    // invoke: bool isNumber()
    C.C_ZNK5QChar8isNumberEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "isNumber", args)
  }

}

// toLower(uint)
func (this *QChar) toLower_s(args ...interface{}) () {
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
    C.C_ZN5QChar7toLowerEj(arg0)
  case 1:
    // invoke: _ZNK5QChar7toLowerEv
    // invoke: QChar toLower()
    C.C_ZNK5QChar7toLowerEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "toLower", args)
  }

}

// toTitleCase()
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
    C.C_ZNK5QChar11toTitleCaseEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar11toTitleCaseEj
    // invoke: uint toTitleCase(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar11toTitleCaseEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "toTitleCase", args)
  }

}

// isLowSurrogate(uint)
func (this *QChar) isLowSurrogate_s(args ...interface{}) () {
  // isLowSurrogate(uint)
  // isLowSurrogate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar14isLowSurrogateEj
    // invoke: bool isLowSurrogate(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar14isLowSurrogateEj(arg0)
  case 1:
    // invoke: _ZNK5QChar14isLowSurrogateEv
    // invoke: bool isLowSurrogate()
    C.C_ZNK5QChar14isLowSurrogateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "isLowSurrogate", args)
  }

}

// isUpper()
func (this *QChar) isUpper(args ...interface{}) () {
  // isUpper()
  // isUpper(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar7isUpperEv
    // invoke: bool isUpper()
    C.C_ZNK5QChar7isUpperEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar7isUpperEj
    // invoke: bool isUpper(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar7isUpperEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "isUpper", args)
  }

}

// QChar(char)
func NewQChar(args ...interface{}) QChar {
  // QChar(char)
  // QChar(uint)
  // QChar(int)
  // QChar(uchar)
  // QChar(short)
  // QChar(uchar, uchar)
  // QChar()
  // QChar(ushort)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "char"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.ByteTy(false) // "uchar"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int16Ty(false) // "short"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.ByteTy(false) // "uchar"
  vtys[5][1] = qtrt.ByteTy(false) // "uchar"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int16Ty(false) // "ushort"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QCharC1Ec
    // invoke: void QChar(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN5QCharC2Ec(qthis, arg0)
  case 1:
    // invoke: _ZN5QCharC1Ej
    // invoke: void QChar(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN5QCharC2Ej(qthis, arg0)
  case 2:
    // invoke: _ZN5QCharC1Ei
    // invoke: void QChar(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN5QCharC2Ei(qthis, arg0)
  case 3:
    // invoke: _ZN5QCharC1Eh
    // invoke: void QChar(uchar)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN5QCharC2Eh(qthis, arg0)
  case 4:
    // invoke: _ZN5QCharC1Es
    // invoke: void QChar(short)
    var arg0 = C.int16_t(args[0].(int16))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN5QCharC2Es(qthis, arg0)
  case 5:
    // invoke: _ZN5QCharC1Ehh
    // invoke: void QChar(uchar, uchar)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN5QCharC2Ehh(qthis, arg0, arg1)
  case 6:
    // invoke: _ZN5QCharC1Ev
    // invoke: void QChar()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN5QCharC2Ev(qthis)
  case 7:
    // invoke: _ZN5QCharC1Et
    // invoke: void QChar(ushort)
    var arg0 = C.int16_t(args[0].(int16))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN5QCharC2Et(qthis, arg0)
  default:
    qtrt.ErrorResolve("QChar", "QChar", args)
  }

  return QChar{}
}

// toCaseFolded(uint)
func (this *QChar) toCaseFolded_s(args ...interface{}) () {
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
    C.C_ZN5QChar12toCaseFoldedEj(arg0)
  case 1:
    // invoke: _ZNK5QChar12toCaseFoldedEv
    // invoke: QChar toCaseFolded()
    C.C_ZNK5QChar12toCaseFoldedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "toCaseFolded", args)
  }

}

// isLetter()
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
    C.C_ZNK5QChar8isLetterEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar8isLetterEj
    // invoke: bool isLetter(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar8isLetterEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "isLetter", args)
  }

}

// isNonCharacter(uint)
func (this *QChar) isNonCharacter_s(args ...interface{}) () {
  // isNonCharacter(uint)
  // isNonCharacter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar14isNonCharacterEj
    // invoke: bool isNonCharacter(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar14isNonCharacterEj(arg0)
  case 1:
    // invoke: _ZNK5QChar14isNonCharacterEv
    // invoke: bool isNonCharacter()
    C.C_ZNK5QChar14isNonCharacterEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "isNonCharacter", args)
  }

}

// isNull()
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
    C.C_ZNK5QChar6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "isNull", args)
  }

}

// isPrint()
func (this *QChar) isPrint(args ...interface{}) () {
  // isPrint()
  // isPrint(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar7isPrintEv
    // invoke: bool isPrint()
    C.C_ZNK5QChar7isPrintEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar7isPrintEj
    // invoke: bool isPrint(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar7isPrintEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "isPrint", args)
  }

}

// toUpper()
func (this *QChar) toUpper(args ...interface{}) () {
  // toUpper()
  // toUpper(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar7toUpperEv
    // invoke: QChar toUpper()
    C.C_ZNK5QChar7toUpperEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar7toUpperEj
    // invoke: uint toUpper(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar7toUpperEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "toUpper", args)
  }

}

// surrogateToUcs4(ushort, ushort)
func (this *QChar) surrogateToUcs4_s(args ...interface{}) () {
  // surrogateToUcs4(ushort, ushort)
  // surrogateToUcs4(class QChar, class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int16Ty(false) // "ushort"
  vtys[0][1] = qtrt.Int16Ty(false) // "ushort"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[1][1] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar15surrogateToUcs4Ett
    // invoke: uint surrogateToUcs4(ushort, ushort)
    var arg0 = C.int16_t(args[0].(int16))
    if false {fmt.Println(arg0)}
    var arg1 = C.int16_t(args[1].(int16))
    if false {fmt.Println(arg1)}
    C.C_ZN5QChar15surrogateToUcs4Ett(arg0, arg1)
  case 1:
    // invoke: _ZN5QChar15surrogateToUcs4ES_S_
    // invoke: uint surrogateToUcs4(class QChar, class QChar)
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QChar).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN5QChar15surrogateToUcs4ES_S_(arg0, arg1)
  default:
    qtrt.ErrorResolve("QChar", "surrogateToUcs4", args)
  }

}

// isDigit(uint)
func (this *QChar) isDigit_s(args ...interface{}) () {
  // isDigit(uint)
  // isDigit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar7isDigitEj
    // invoke: bool isDigit(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar7isDigitEj(arg0)
  case 1:
    // invoke: _ZNK5QChar7isDigitEv
    // invoke: bool isDigit()
    C.C_ZNK5QChar7isDigitEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "isDigit", args)
  }

}

// setRow(uchar)
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
    C.C_ZN5QChar6setRowEh(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QChar", "setRow", args)
  }

}

// isPunct()
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
    C.C_ZNK5QChar7isPunctEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar7isPunctEj
    // invoke: bool isPunct(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar7isPunctEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "isPunct", args)
  }

}

// isSymbol(uint)
func (this *QChar) isSymbol_s(args ...interface{}) () {
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
    C.C_ZN5QChar8isSymbolEj(arg0)
  case 1:
    // invoke: _ZNK5QChar8isSymbolEv
    // invoke: bool isSymbol()
    C.C_ZNK5QChar8isSymbolEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "isSymbol", args)
  }

}

// joining()
func (this *QChar) joining(args ...interface{}) () {
  // joining()
  // joining(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar7joiningEv
    // invoke: QChar::Joining joining()
    C.C_ZNK5QChar7joiningEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QChar7joiningEj
    // invoke: QChar::Joining joining(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar7joiningEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "joining", args)
  }

}

// <= body block end

