package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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
extern void* C_ZN11QLatin1CharC2Ec(unsigned char arg0); // 1
  // proto:  char QLatin1Char::toLatin1();
extern unsigned char C_ZNK11QLatin1Char8toLatin1Ev(void* qthis); // 2
  // proto:  ushort QLatin1Char::unicode();
extern int16_t C_ZNK11QLatin1Char7unicodeEv(void* qthis); // 2
  // proto: static bool QChar::isLetterOrNumber(uint ucs4);
extern bool C_ZN5QChar16isLetterOrNumberEj(int32_t arg0); // 2
  // proto:  bool QChar::isLetterOrNumber();
extern bool C_ZNK5QChar16isLetterOrNumberEv(void* qthis); // 2
  // proto:  QChar::UnicodeVersion QChar::unicodeVersion();
extern void C_ZNK5QChar14unicodeVersionEv(void* qthis); // 2
  // proto: static QChar::UnicodeVersion QChar::unicodeVersion(uint ucs4);
extern void C_ZN5QChar14unicodeVersionEj(int32_t arg0); // 4
  // proto: static bool QChar::hasMirrored(uint ucs4);
extern bool C_ZN5QChar11hasMirroredEj(int32_t arg0); // 4
  // proto:  bool QChar::hasMirrored();
extern bool C_ZNK5QChar11hasMirroredEv(void* qthis); // 2
  // proto:  char QChar::toLatin1();
extern unsigned char C_ZNK5QChar8toLatin1Ev(void* qthis); // 2
  // proto:  bool QChar::isTitleCase();
extern bool C_ZNK5QChar11isTitleCaseEv(void* qthis); // 2
  // proto: static bool QChar::isTitleCase(uint ucs4);
extern bool C_ZN5QChar11isTitleCaseEj(int32_t arg0); // 2
  // proto: static QChar::UnicodeVersion QChar::currentUnicodeVersion();
extern void C_ZN5QChar21currentUnicodeVersionEv(); // 4
  // proto:  void QChar::setCell(uchar cell);
extern void C_ZN5QChar7setCellEh(void* qthis, unsigned char arg0); // 2
  // proto:  ushort & QChar::unicode();
extern void C_ZN5QChar7unicodeEv(void* qthis); // 2
  // proto: static ushort QChar::lowSurrogate(uint ucs4);
extern int16_t C_ZN5QChar12lowSurrogateEj(int32_t arg0); // 2
  // proto: static bool QChar::isSpace(uint ucs4);
extern bool C_ZN5QChar7isSpaceEj(int32_t arg0); // 2
  // proto:  bool QChar::isSpace();
extern bool C_ZNK5QChar7isSpaceEv(void* qthis); // 2
  // proto:  uchar QChar::row();
extern unsigned char C_ZNK5QChar3rowEv(void* qthis); // 2
  // proto:  QChar::Category QChar::category();
extern void C_ZNK5QChar8categoryEv(void* qthis); // 2
  // proto: static QChar::Category QChar::category(uint ucs4);
extern void C_ZN5QChar8categoryEj(int32_t arg0); // 4
  // proto:  QChar::Decomposition QChar::decompositionTag();
extern void C_ZNK5QChar16decompositionTagEv(void* qthis); // 2
  // proto: static QChar::Decomposition QChar::decompositionTag(uint ucs4);
extern void C_ZN5QChar16decompositionTagEj(int32_t arg0); // 4
  // proto:  bool QChar::isHighSurrogate();
extern bool C_ZNK5QChar15isHighSurrogateEv(void* qthis); // 2
  // proto: static bool QChar::isHighSurrogate(uint ucs4);
extern bool C_ZN5QChar15isHighSurrogateEj(int32_t arg0); // 2
  // proto: static QChar::Script QChar::script(uint ucs4);
extern void C_ZN5QChar6scriptEj(int32_t arg0); // 4
  // proto:  QChar::Script QChar::script();
extern void C_ZNK5QChar6scriptEv(void* qthis); // 2
  // proto: static QChar QChar::fromLatin1(char c);
extern void* C_ZN5QChar10fromLatin1Ec(unsigned char arg0); // 2
  // proto: static QChar::JoiningType QChar::joiningType(uint ucs4);
extern void C_ZN5QChar11joiningTypeEj(int32_t arg0); // 4
  // proto:  QChar::JoiningType QChar::joiningType();
extern void C_ZNK5QChar11joiningTypeEv(void* qthis); // 2
  // proto:  uchar QChar::cell();
extern unsigned char C_ZNK5QChar4cellEv(void* qthis); // 2
  // proto:  int QChar::digitValue();
extern int32_t C_ZNK5QChar10digitValueEv(void* qthis); // 2
  // proto: static int QChar::digitValue(uint ucs4);
extern int32_t C_ZN5QChar10digitValueEj(int32_t arg0); // 4
  // proto: static bool QChar::requiresSurrogates(uint ucs4);
extern bool C_ZN5QChar18requiresSurrogatesEj(int32_t arg0); // 2
  // proto:  QString QChar::decomposition();
extern void* C_ZNK5QChar13decompositionEv(void* qthis); // 4
  // proto: static QString QChar::decomposition(uint ucs4);
extern void* C_ZN5QChar13decompositionEj(int32_t arg0); // 4
  // proto: static ushort QChar::highSurrogate(uint ucs4);
extern int16_t C_ZN5QChar13highSurrogateEj(int32_t arg0); // 2
  // proto:  bool QChar::isSurrogate();
extern bool C_ZNK5QChar11isSurrogateEv(void* qthis); // 2
  // proto: static bool QChar::isSurrogate(uint ucs4);
extern bool C_ZN5QChar11isSurrogateEj(int32_t arg0); // 2
  // proto: static QChar::Direction QChar::direction(uint ucs4);
extern void C_ZN5QChar9directionEj(int32_t arg0); // 4
  // proto:  QChar::Direction QChar::direction();
extern void C_ZNK5QChar9directionEv(void* qthis); // 2
  // proto: static uint QChar::mirroredChar(uint ucs4);
extern int32_t C_ZN5QChar12mirroredCharEj(int32_t arg0); // 4
  // proto:  QChar QChar::mirroredChar();
extern void* C_ZNK5QChar12mirroredCharEv(void* qthis); // 2
  // proto:  unsigned char QChar::combiningClass();
extern unsigned char C_ZNK5QChar14combiningClassEv(void* qthis); // 2
  // proto: static unsigned char QChar::combiningClass(uint ucs4);
extern unsigned char C_ZN5QChar14combiningClassEj(int32_t arg0); // 4
  // proto: static bool QChar::isMark(uint ucs4);
extern bool C_ZN5QChar6isMarkEj(int32_t arg0); // 4
  // proto:  bool QChar::isMark();
extern bool C_ZNK5QChar6isMarkEv(void* qthis); // 2
  // proto: static bool QChar::isLower(uint ucs4);
extern bool C_ZN5QChar7isLowerEj(int32_t arg0); // 2
  // proto:  bool QChar::isLower();
extern bool C_ZNK5QChar7isLowerEv(void* qthis); // 2
  // proto: static bool QChar::isNumber(uint ucs4);
extern bool C_ZN5QChar8isNumberEj(int32_t arg0); // 2
  // proto:  bool QChar::isNumber();
extern bool C_ZNK5QChar8isNumberEv(void* qthis); // 2
  // proto: static uint QChar::toLower(uint ucs4);
extern int32_t C_ZN5QChar7toLowerEj(int32_t arg0); // 4
  // proto:  QChar QChar::toLower();
extern void* C_ZNK5QChar7toLowerEv(void* qthis); // 2
  // proto:  QChar QChar::toTitleCase();
extern void* C_ZNK5QChar11toTitleCaseEv(void* qthis); // 2
  // proto: static uint QChar::toTitleCase(uint ucs4);
extern int32_t C_ZN5QChar11toTitleCaseEj(int32_t arg0); // 4
  // proto: static bool QChar::isLowSurrogate(uint ucs4);
extern bool C_ZN5QChar14isLowSurrogateEj(int32_t arg0); // 2
  // proto:  bool QChar::isLowSurrogate();
extern bool C_ZNK5QChar14isLowSurrogateEv(void* qthis); // 2
  // proto:  bool QChar::isUpper();
extern bool C_ZNK5QChar7isUpperEv(void* qthis); // 2
  // proto: static bool QChar::isUpper(uint ucs4);
extern bool C_ZN5QChar7isUpperEj(int32_t arg0); // 2
  // proto:  void QChar::QChar(char c);
extern void* C_ZN5QCharC2Ec(unsigned char arg0); // 1
  // proto:  void QChar::QChar(uint rc);
extern void* C_ZN5QCharC2Ej(int32_t arg0); // 1
  // proto:  void QChar::QChar(int rc);
extern void* C_ZN5QCharC2Ei(int32_t arg0); // 1
  // proto:  void QChar::QChar(uchar c);
extern void* C_ZN5QCharC2Eh(unsigned char arg0); // 1
  // proto:  void QChar::QChar(short rc);
extern void* C_ZN5QCharC2Es(int16_t arg0); // 1
  // proto:  void QChar::QChar(uchar c, uchar r);
extern void* C_ZN5QCharC2Ehh(unsigned char arg0, unsigned char arg1); // 1
  // proto:  void QChar::QChar();
extern void* C_ZN5QCharC2Ev(); // 1
  // proto:  void QChar::QChar(ushort rc);
extern void* C_ZN5QCharC2Et(int16_t arg0); // 1
  // proto: static uint QChar::toCaseFolded(uint ucs4);
extern int32_t C_ZN5QChar12toCaseFoldedEj(int32_t arg0); // 4
  // proto:  QChar QChar::toCaseFolded();
extern void* C_ZNK5QChar12toCaseFoldedEv(void* qthis); // 2
  // proto:  bool QChar::isLetter();
extern bool C_ZNK5QChar8isLetterEv(void* qthis); // 2
  // proto: static bool QChar::isLetter(uint ucs4);
extern bool C_ZN5QChar8isLetterEj(int32_t arg0); // 2
  // proto: static bool QChar::isNonCharacter(uint ucs4);
extern bool C_ZN5QChar14isNonCharacterEj(int32_t arg0); // 2
  // proto:  bool QChar::isNonCharacter();
extern bool C_ZNK5QChar14isNonCharacterEv(void* qthis); // 2
  // proto:  bool QChar::isNull();
extern bool C_ZNK5QChar6isNullEv(void* qthis); // 2
  // proto:  bool QChar::isPrint();
extern bool C_ZNK5QChar7isPrintEv(void* qthis); // 2
  // proto: static bool QChar::isPrint(uint ucs4);
extern bool C_ZN5QChar7isPrintEj(int32_t arg0); // 4
  // proto:  QChar QChar::toUpper();
extern void* C_ZNK5QChar7toUpperEv(void* qthis); // 2
  // proto: static uint QChar::toUpper(uint ucs4);
extern int32_t C_ZN5QChar7toUpperEj(int32_t arg0); // 4
  // proto: static uint QChar::surrogateToUcs4(ushort high, ushort low);
extern int32_t C_ZN5QChar15surrogateToUcs4Ett(int16_t arg0, int16_t arg1); // 2
  // proto: static uint QChar::surrogateToUcs4(QChar high, QChar low);
extern int32_t C_ZN5QChar15surrogateToUcs4ES_S_(void* arg0, void* arg1); // 2
  // proto: static bool QChar::isDigit(uint ucs4);
extern bool C_ZN5QChar7isDigitEj(int32_t arg0); // 2
  // proto:  bool QChar::isDigit();
extern bool C_ZNK5QChar7isDigitEv(void* qthis); // 2
  // proto:  void QChar::setRow(uchar row);
extern void C_ZN5QChar6setRowEh(void* qthis, unsigned char arg0); // 2
  // proto:  bool QChar::isPunct();
extern bool C_ZNK5QChar7isPunctEv(void* qthis); // 2
  // proto: static bool QChar::isPunct(uint ucs4);
extern bool C_ZN5QChar7isPunctEj(int32_t arg0); // 4
  // proto: static bool QChar::isSymbol(uint ucs4);
extern bool C_ZN5QChar8isSymbolEj(int32_t arg0); // 4
  // proto:  bool QChar::isSymbol();
extern bool C_ZNK5QChar8isSymbolEv(void* qthis); // 2
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QChar)=2
type QChar struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// QLatin1Char(char)
func NewQLatin1Char(args ...interface{}) *QLatin1Char {
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
    qthis = C.C_ZN11QLatin1CharC2Ec(arg0)
    return &QLatin1Char{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QLatin1Char", "QLatin1Char", args)
  }

  return nil // QLatin1Char{Qclsinst:qthis}
}

// toLatin1()
func (this *QLatin1Char) Tolatin1(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QLatin1Char8toLatin1Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "char"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLatin1Char", "toLatin1", args)
  }

  return
}

// unicode()
func (this *QLatin1Char) Unicode(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QLatin1Char7unicodeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int16Ty(false) // "ushort"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLatin1Char", "unicode", args)
  }

  return
}

// isLetterOrNumber(uint)
func (this *QChar) Isletterornumber_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar16isLetterOrNumberEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK5QChar16isLetterOrNumberEv
    // invoke: bool isLetterOrNumber()
    var ret0 = C.C_ZNK5QChar16isLetterOrNumberEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "isLetterOrNumber", args)
  }

  return
}

// unicodeVersion()
func (this *QChar) Unicodeversion(args ...interface{}) () {
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
    C.C_ZNK5QChar14unicodeVersionEv(this.Qclsinst)
  case 1:
    // invoke: _ZN5QChar14unicodeVersionEj
    // invoke: QChar::UnicodeVersion unicodeVersion(uint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar14unicodeVersionEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "unicodeVersion", args)
  }

  return
}

// hasMirrored(uint)
func (this *QChar) Hasmirrored_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar11hasMirroredEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK5QChar11hasMirroredEv
    // invoke: bool hasMirrored()
    var ret0 = C.C_ZNK5QChar11hasMirroredEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "hasMirrored", args)
  }

  return
}

// toLatin1()
func (this *QChar) Tolatin1(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QChar8toLatin1Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "char"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "toLatin1", args)
  }

  return
}

// isTitleCase()
func (this *QChar) Istitlecase(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QChar11isTitleCaseEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN5QChar11isTitleCaseEj
    // invoke: bool isTitleCase(uint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar11isTitleCaseEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "isTitleCase", args)
  }

  return
}

// currentUnicodeVersion()
func (this *QChar) Currentunicodeversion_S(args ...interface{}) () {
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

  return
}

// setCell(uchar)
func (this *QChar) Setcell(args ...interface{}) () {
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
    C.C_ZN5QChar7setCellEh(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QChar", "setCell", args)
  }

  return
}

// unicode()
func (this *QChar) Unicode(args ...interface{}) () {
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
    C.C_ZN5QChar7unicodeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "unicode", args)
  }

  return
}

// lowSurrogate(uint)
func (this *QChar) Lowsurrogate_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar12lowSurrogateEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int16Ty(false) // "ushort"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "lowSurrogate", args)
  }

  return
}

// isSpace(uint)
func (this *QChar) Isspace_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar7isSpaceEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK5QChar7isSpaceEv
    // invoke: bool isSpace()
    var ret0 = C.C_ZNK5QChar7isSpaceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "isSpace", args)
  }

  return
}

// row()
func (this *QChar) Row(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QChar3rowEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "uchar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "row", args)
  }

  return
}

// category()
func (this *QChar) Category(args ...interface{}) () {
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
    C.C_ZNK5QChar8categoryEv(this.Qclsinst)
  case 1:
    // invoke: _ZN5QChar8categoryEj
    // invoke: QChar::Category category(uint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar8categoryEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "category", args)
  }

  return
}

// decompositionTag()
func (this *QChar) Decompositiontag(args ...interface{}) () {
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
    C.C_ZNK5QChar16decompositionTagEv(this.Qclsinst)
  case 1:
    // invoke: _ZN5QChar16decompositionTagEj
    // invoke: QChar::Decomposition decompositionTag(uint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar16decompositionTagEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "decompositionTag", args)
  }

  return
}

// isHighSurrogate()
func (this *QChar) Ishighsurrogate(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QChar15isHighSurrogateEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN5QChar15isHighSurrogateEj
    // invoke: bool isHighSurrogate(uint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar15isHighSurrogateEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "isHighSurrogate", args)
  }

  return
}

// script(uint)
func (this *QChar) Script_S(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar6scriptEj(arg0)
  case 1:
    // invoke: _ZNK5QChar6scriptEv
    // invoke: QChar::Script script()
    C.C_ZNK5QChar6scriptEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "script", args)
  }

  return
}

// fromLatin1(char)
func (this *QChar) Fromlatin1_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN5QChar10fromLatin1Ec(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "QChar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "fromLatin1", args)
  }

  return
}

// joiningType(uint)
func (this *QChar) Joiningtype_S(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar11joiningTypeEj(arg0)
  case 1:
    // invoke: _ZNK5QChar11joiningTypeEv
    // invoke: QChar::JoiningType joiningType()
    C.C_ZNK5QChar11joiningTypeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "joiningType", args)
  }

  return
}

// cell()
func (this *QChar) Cell(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QChar4cellEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "uchar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "cell", args)
  }

  return
}

// digitValue()
func (this *QChar) Digitvalue(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QChar10digitValueEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN5QChar10digitValueEj
    // invoke: int digitValue(uint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar10digitValueEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "digitValue", args)
  }

  return
}

// requiresSurrogates(uint)
func (this *QChar) Requiressurrogates_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar18requiresSurrogatesEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "requiresSurrogates", args)
  }

  return
}

// decomposition()
func (this *QChar) Decomposition(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QChar13decompositionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN5QChar13decompositionEj
    // invoke: QString decomposition(uint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar13decompositionEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "decomposition", args)
  }

  return
}

// highSurrogate(uint)
func (this *QChar) Highsurrogate_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar13highSurrogateEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int16Ty(false) // "ushort"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "highSurrogate", args)
  }

  return
}

// isSurrogate()
func (this *QChar) Issurrogate(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QChar11isSurrogateEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN5QChar11isSurrogateEj
    // invoke: bool isSurrogate(uint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar11isSurrogateEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "isSurrogate", args)
  }

  return
}

// direction(uint)
func (this *QChar) Direction_S(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar9directionEj(arg0)
  case 1:
    // invoke: _ZNK5QChar9directionEv
    // invoke: QChar::Direction direction()
    C.C_ZNK5QChar9directionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QChar", "direction", args)
  }

  return
}

// mirroredChar(uint)
func (this *QChar) Mirroredchar_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar12mirroredCharEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "uint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK5QChar12mirroredCharEv
    // invoke: QChar mirroredChar()
    var ret0 = C.C_ZNK5QChar12mirroredCharEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "QChar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "mirroredChar", args)
  }

  return
}

// combiningClass()
func (this *QChar) Combiningclass(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QChar14combiningClassEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "unsigned char"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN5QChar14combiningClassEj
    // invoke: unsigned char combiningClass(uint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar14combiningClassEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "unsigned char"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "combiningClass", args)
  }

  return
}

// isMark(uint)
func (this *QChar) Ismark_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar6isMarkEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK5QChar6isMarkEv
    // invoke: bool isMark()
    var ret0 = C.C_ZNK5QChar6isMarkEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "isMark", args)
  }

  return
}

// isLower(uint)
func (this *QChar) Islower_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar7isLowerEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK5QChar7isLowerEv
    // invoke: bool isLower()
    var ret0 = C.C_ZNK5QChar7isLowerEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "isLower", args)
  }

  return
}

// isNumber(uint)
func (this *QChar) Isnumber_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar8isNumberEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK5QChar8isNumberEv
    // invoke: bool isNumber()
    var ret0 = C.C_ZNK5QChar8isNumberEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "isNumber", args)
  }

  return
}

// toLower(uint)
func (this *QChar) Tolower_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar7toLowerEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "uint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK5QChar7toLowerEv
    // invoke: QChar toLower()
    var ret0 = C.C_ZNK5QChar7toLowerEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "QChar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "toLower", args)
  }

  return
}

// toTitleCase()
func (this *QChar) Totitlecase(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QChar11toTitleCaseEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "QChar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN5QChar11toTitleCaseEj
    // invoke: uint toTitleCase(uint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar11toTitleCaseEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "uint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "toTitleCase", args)
  }

  return
}

// isLowSurrogate(uint)
func (this *QChar) Islowsurrogate_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar14isLowSurrogateEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK5QChar14isLowSurrogateEv
    // invoke: bool isLowSurrogate()
    var ret0 = C.C_ZNK5QChar14isLowSurrogateEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "isLowSurrogate", args)
  }

  return
}

// isUpper()
func (this *QChar) Isupper(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QChar7isUpperEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN5QChar7isUpperEj
    // invoke: bool isUpper(uint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar7isUpperEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "isUpper", args)
  }

  return
}

// QChar(char)
func NewQChar(args ...interface{}) *QChar {
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
    qthis = C.C_ZN5QCharC2Ec(arg0)
    return &QChar{Qclsinst:qthis}
  case 1:
    // invoke: _ZN5QCharC1Ej
    // invoke: void QChar(uint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QCharC2Ej(arg0)
    return &QChar{Qclsinst:qthis}
  case 2:
    // invoke: _ZN5QCharC1Ei
    // invoke: void QChar(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QCharC2Ei(arg0)
    return &QChar{Qclsinst:qthis}
  case 3:
    // invoke: _ZN5QCharC1Eh
    // invoke: void QChar(uchar)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QCharC2Eh(arg0)
    return &QChar{Qclsinst:qthis}
  case 4:
    // invoke: _ZN5QCharC1Es
    // invoke: void QChar(short)
    var arg0 = C.int16_t(qtrt.PrimConv(args[0], qtrt.Int16Ty(false)).(int16))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QCharC2Es(arg0)
    return &QChar{Qclsinst:qthis}
  case 5:
    // invoke: _ZN5QCharC1Ehh
    // invoke: void QChar(uchar, uchar)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QCharC2Ehh(arg0, arg1)
    return &QChar{Qclsinst:qthis}
  case 6:
    // invoke: _ZN5QCharC1Ev
    // invoke: void QChar()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QCharC2Ev()
    return &QChar{Qclsinst:qthis}
  case 7:
    // invoke: _ZN5QCharC1Et
    // invoke: void QChar(ushort)
    var arg0 = C.int16_t(qtrt.PrimConv(args[0], qtrt.Int16Ty(false)).(int16))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QCharC2Et(arg0)
    return &QChar{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QChar", "QChar", args)
  }

  return nil // QChar{Qclsinst:qthis}
}

// toCaseFolded(uint)
func (this *QChar) Tocasefolded_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar12toCaseFoldedEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "uint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK5QChar12toCaseFoldedEv
    // invoke: QChar toCaseFolded()
    var ret0 = C.C_ZNK5QChar12toCaseFoldedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "QChar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "toCaseFolded", args)
  }

  return
}

// isLetter()
func (this *QChar) Isletter(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QChar8isLetterEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN5QChar8isLetterEj
    // invoke: bool isLetter(uint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar8isLetterEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "isLetter", args)
  }

  return
}

// isNonCharacter(uint)
func (this *QChar) Isnoncharacter_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar14isNonCharacterEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK5QChar14isNonCharacterEv
    // invoke: bool isNonCharacter()
    var ret0 = C.C_ZNK5QChar14isNonCharacterEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "isNonCharacter", args)
  }

  return
}

// isNull()
func (this *QChar) Isnull(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QChar6isNullEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "isNull", args)
  }

  return
}

// isPrint()
func (this *QChar) Isprint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QChar7isPrintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN5QChar7isPrintEj
    // invoke: bool isPrint(uint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar7isPrintEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "isPrint", args)
  }

  return
}

// toUpper()
func (this *QChar) Toupper(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QChar7toUpperEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "QChar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN5QChar7toUpperEj
    // invoke: uint toUpper(uint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar7toUpperEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "uint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "toUpper", args)
  }

  return
}

// surrogateToUcs4(ushort, ushort)
func (this *QChar) Surrogatetoucs4_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int16_t(qtrt.PrimConv(args[0], qtrt.Int16Ty(false)).(int16))
    if false {fmt.Println(arg0)}
    var arg1 = C.int16_t(qtrt.PrimConv(args[1], qtrt.Int16Ty(false)).(int16))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN5QChar15surrogateToUcs4Ett(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "uint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN5QChar15surrogateToUcs4ES_S_
    // invoke: uint surrogateToUcs4(class QChar, class QChar)
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QChar).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN5QChar15surrogateToUcs4ES_S_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "uint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "surrogateToUcs4", args)
  }

  return
}

// isDigit(uint)
func (this *QChar) Isdigit_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar7isDigitEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK5QChar7isDigitEv
    // invoke: bool isDigit()
    var ret0 = C.C_ZNK5QChar7isDigitEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "isDigit", args)
  }

  return
}

// setRow(uchar)
func (this *QChar) Setrow(args ...interface{}) () {
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
    C.C_ZN5QChar6setRowEh(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QChar", "setRow", args)
  }

  return
}

// isPunct()
func (this *QChar) Ispunct(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QChar7isPunctEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN5QChar7isPunctEj
    // invoke: bool isPunct(uint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar7isPunctEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "isPunct", args)
  }

  return
}

// isSymbol(uint)
func (this *QChar) Issymbol_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QChar8isSymbolEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK5QChar8isSymbolEv
    // invoke: bool isSymbol()
    var ret0 = C.C_ZNK5QChar8isSymbolEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChar", "isSymbol", args)
  }

  return
}

// joining()
func (this *QChar) Joining(args ...interface{}) () {
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
    C.C_ZNK5QChar7joiningEv(this.Qclsinst)
  case 1:
    // invoke: _ZN5QChar7joiningEj
    // invoke: QChar::Joining joining(uint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QChar7joiningEj(arg0)
  default:
    qtrt.ErrorResolve("QChar", "joining", args)
  }

  return
}

// <= body block end

