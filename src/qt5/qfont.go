package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtGui/qfont.h
// dst-file: /src/gui/qfont.go
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
  // proto:  bool QFont::isCopyOf(const QFont & );
extern void C_ZNK5QFont8isCopyOfERKS_(void* qthis, void* arg0); // 4
  // proto:  qreal QFont::wordSpacing();
extern void C_ZNK5QFont11wordSpacingEv(void* qthis); // 4
  // proto:  QFont::SpacingType QFont::letterSpacingType();
extern void C_ZNK5QFont17letterSpacingTypeEv(void* qthis); // 4
  // proto:  QString QFont::lastResortFamily();
extern void C_ZNK5QFont16lastResortFamilyEv(void* qthis); // 4
  // proto:  void QFont::setStretch(int );
extern void C_ZN5QFont10setStretchEi(void* qthis, int32_t arg0); // 4
  // proto:  int QFont::weight();
extern void C_ZNK5QFont6weightEv(void* qthis); // 4
  // proto:  QFont::StyleStrategy QFont::styleStrategy();
extern void C_ZNK5QFont13styleStrategyEv(void* qthis); // 4
  // proto:  bool QFont::fromString(const QString & );
extern void C_ZN5QFont10fromStringERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QFont::setPointSizeF(qreal );
extern void C_ZN5QFont13setPointSizeFEd(void* qthis, double arg0); // 4
  // proto: static void QFont::insertSubstitutions(const QString & , const QStringList & );
extern void C_ZN5QFont19insertSubstitutionsERK7QStringRK11QStringList(void* arg0, void* arg1); // 4
  // proto:  bool QFont::rawMode();
extern void C_ZNK5QFont7rawModeEv(void* qthis); // 4
  // proto:  void QFont::setItalic(bool b);
extern void C_ZN5QFont9setItalicEb(void* qthis, bool arg0); // 2
  // proto:  QString QFont::toString();
extern void C_ZNK5QFont8toStringEv(void* qthis); // 4
  // proto:  bool QFont::underline();
extern void C_ZNK5QFont9underlineEv(void* qthis); // 4
  // proto:  void QFont::setStyleName(const QString & );
extern void C_ZN5QFont12setStyleNameERK7QString(void* qthis, void* arg0); // 4
  // proto: static void QFont::removeSubstitutions(const QString & );
extern void C_ZN5QFont19removeSubstitutionsERK7QString(void* arg0); // 4
  // proto:  void QFont::setPointSize(int );
extern void C_ZN5QFont12setPointSizeEi(void* qthis, int32_t arg0); // 4
  // proto:  QFont::Style QFont::style();
extern void C_ZNK5QFont5styleEv(void* qthis); // 4
  // proto:  QString QFont::family();
extern void C_ZNK5QFont6familyEv(void* qthis); // 4
  // proto:  void QFont::setBold(bool );
extern void C_ZN5QFont7setBoldEb(void* qthis, bool arg0); // 2
  // proto:  QString QFont::key();
extern void C_ZNK5QFont3keyEv(void* qthis); // 4
  // proto:  int QFont::stretch();
extern void C_ZNK5QFont7stretchEv(void* qthis); // 4
  // proto:  qreal QFont::pointSizeF();
extern void C_ZNK5QFont10pointSizeFEv(void* qthis); // 4
  // proto:  int QFont::pointSize();
extern void C_ZNK5QFont9pointSizeEv(void* qthis); // 4
  // proto:  bool QFont::fixedPitch();
extern void C_ZNK5QFont10fixedPitchEv(void* qthis); // 4
  // proto: static QStringList QFont::substitutions();
extern void C_ZN5QFont13substitutionsEv(); // 4
  // proto: static void QFont::cacheStatistics();
extern void C_ZN5QFont15cacheStatisticsEv(); // 4
  // proto:  void QFont::~QFont();
extern void C_ZN5QFontD2Ev(void* qthis); // 4
  // proto:  bool QFont::kerning();
extern void C_ZNK5QFont7kerningEv(void* qthis); // 4
  // proto: static QStringList QFont::substitutes(const QString & );
extern void C_ZN5QFont11substitutesERK7QString(void* arg0); // 4
  // proto:  void QFont::setRawMode(bool );
extern void C_ZN5QFont10setRawModeEb(void* qthis, bool arg0); // 4
  // proto:  bool QFont::italic();
extern void C_ZNK5QFont6italicEv(void* qthis); // 2
  // proto:  int QFont::pixelSize();
extern void C_ZNK5QFont9pixelSizeEv(void* qthis); // 4
  // proto:  void QFont::swap(QFont & other);
extern void C_ZN5QFont4swapERS_(void* qthis, void* arg0); // 2
  // proto:  QString QFont::rawName();
extern void C_ZNK5QFont7rawNameEv(void* qthis); // 4
  // proto: static void QFont::insertSubstitution(const QString & , const QString & );
extern void C_ZN5QFont18insertSubstitutionERK7QStringS2_(void* arg0, void* arg1); // 4
  // proto:  bool QFont::strikeOut();
extern void C_ZNK5QFont9strikeOutEv(void* qthis); // 4
  // proto:  bool QFont::bold();
extern void C_ZNK5QFont4boldEv(void* qthis); // 2
  // proto:  QFont::Capitalization QFont::capitalization();
extern void C_ZNK5QFont14capitalizationEv(void* qthis); // 4
  // proto:  QFont::StyleHint QFont::styleHint();
extern void C_ZNK5QFont9styleHintEv(void* qthis); // 4
  // proto:  void QFont::setKerning(bool );
extern void C_ZN5QFont10setKerningEb(void* qthis, bool arg0); // 4
  // proto:  void QFont::QFont(const QFont & );
extern void* C_ZN5QFontC2ERKS_(void* arg0); // 3
  // proto:  void QFont::QFont();
extern void* C_ZN5QFontC2Ev(); // 3
  // proto:  void QFont::QFont(const QString & family, int pointSize, int weight, bool italic);
extern void* C_ZN5QFontC2ERK7QStringiib(void* arg0, int32_t arg1, int32_t arg2, bool arg3); // 3
  // proto:  void QFont::QFont(const QFont & , QPaintDevice * pd);
extern void* C_ZN5QFontC2ERKS_P12QPaintDevice(void* arg0, void* arg1); // 3
  // proto: static QString QFont::substitute(const QString & );
extern void C_ZN5QFont10substituteERK7QString(void* arg0); // 4
  // proto:  QFont::HintingPreference QFont::hintingPreference();
extern void C_ZNK5QFont17hintingPreferenceEv(void* qthis); // 4
  // proto:  void QFont::setRawName(const QString & );
extern void C_ZN5QFont10setRawNameERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QFont::setOverline(bool );
extern void C_ZN5QFont11setOverlineEb(void* qthis, bool arg0); // 4
  // proto:  void QFont::setFixedPitch(bool );
extern void C_ZN5QFont13setFixedPitchEb(void* qthis, bool arg0); // 4
  // proto:  qreal QFont::letterSpacing();
extern void C_ZNK5QFont13letterSpacingEv(void* qthis); // 4
  // proto: static void QFont::initialize();
extern void C_ZN5QFont10initializeEv(); // 4
  // proto:  void QFont::setPixelSize(int );
extern void C_ZN5QFont12setPixelSizeEi(void* qthis, int32_t arg0); // 4
  // proto:  void QFont::resolve(uint mask);
extern void C_ZN5QFont7resolveEj(void* qthis, int32_t arg0); // 2
  // proto:  uint QFont::resolve();
extern void C_ZNK5QFont7resolveEv(void* qthis); // 2
  // proto:  QFont QFont::resolve(const QFont & );
extern void C_ZNK5QFont7resolveERKS_(void* qthis, void* arg0); // 4
  // proto:  bool QFont::exactMatch();
extern void C_ZNK5QFont10exactMatchEv(void* qthis); // 4
  // proto:  void QFont::setWordSpacing(qreal spacing);
extern void C_ZN5QFont14setWordSpacingEd(void* qthis, double arg0); // 4
  // proto:  void QFont::setStrikeOut(bool );
extern void C_ZN5QFont12setStrikeOutEb(void* qthis, bool arg0); // 4
  // proto:  void QFont::setWeight(int );
extern void C_ZN5QFont9setWeightEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QFont::lastResortFont();
extern void C_ZNK5QFont14lastResortFontEv(void* qthis); // 4
  // proto:  bool QFont::overline();
extern void C_ZNK5QFont8overlineEv(void* qthis); // 4
  // proto:  QString QFont::defaultFamily();
extern void C_ZNK5QFont13defaultFamilyEv(void* qthis); // 4
  // proto:  void QFont::setFamily(const QString & );
extern void C_ZN5QFont9setFamilyERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QFont::setUnderline(bool );
extern void C_ZN5QFont12setUnderlineEb(void* qthis, bool arg0); // 4
  // proto:  QString QFont::styleName();
extern void C_ZNK5QFont9styleNameEv(void* qthis); // 4
  // proto: static void QFont::cleanup();
extern void C_ZN5QFont7cleanupEv(); // 4
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

// class sizeof(QFont)=1
type QFont struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// isCopyOf(const class QFont &)
func (this *QFont) isCopyOf(args ...interface{}) () {
  // isCopyOf(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont8isCopyOfERKS_
    // invoke: bool isCopyOf(const class QFont &)
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK5QFont8isCopyOfERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "isCopyOf", args)
  }

}

// wordSpacing()
func (this *QFont) wordSpacing(args ...interface{}) () {
  // wordSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont11wordSpacingEv
    // invoke: qreal wordSpacing()
    var ret = C.C_ZNK5QFont11wordSpacingEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "wordSpacing", args)
  }

}

// letterSpacingType()
func (this *QFont) letterSpacingType(args ...interface{}) () {
  // letterSpacingType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont17letterSpacingTypeEv
    // invoke: QFont::SpacingType letterSpacingType()
    C.C_ZNK5QFont17letterSpacingTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "letterSpacingType", args)
  }

}

// lastResortFamily()
func (this *QFont) lastResortFamily(args ...interface{}) () {
  // lastResortFamily()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont16lastResortFamilyEv
    // invoke: QString lastResortFamily()
    var ret = C.C_ZNK5QFont16lastResortFamilyEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "lastResortFamily", args)
  }

}

// setStretch(int)
func (this *QFont) setStretch(args ...interface{}) () {
  // setStretch(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont10setStretchEi
    // invoke: void setStretch(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont10setStretchEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setStretch", args)
  }

}

// weight()
func (this *QFont) weight(args ...interface{}) () {
  // weight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont6weightEv
    // invoke: int weight()
    var ret = C.C_ZNK5QFont6weightEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "weight", args)
  }

}

// styleStrategy()
func (this *QFont) styleStrategy(args ...interface{}) () {
  // styleStrategy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont13styleStrategyEv
    // invoke: QFont::StyleStrategy styleStrategy()
    C.C_ZNK5QFont13styleStrategyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "styleStrategy", args)
  }

}

// fromString(const class QString &)
func (this *QFont) fromString(args ...interface{}) () {
  // fromString(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont10fromStringERK7QString
    // invoke: bool fromString(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN5QFont10fromStringERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "fromString", args)
  }

}

// setPointSizeF(qreal)
func (this *QFont) setPointSizeF(args ...interface{}) () {
  // setPointSizeF(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont13setPointSizeFEd
    // invoke: void setPointSizeF(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont13setPointSizeFEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setPointSizeF", args)
  }

}

// insertSubstitutions(const class QString &, const class QStringList &)
func (this *QFont) insertSubstitutions_s(args ...interface{}) () {
  // insertSubstitutions(const class QString &, const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont19insertSubstitutionsERK7QStringRK11QStringList
    // invoke: void insertSubstitutions(const class QString &, const class QStringList &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStringList).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN5QFont19insertSubstitutionsERK7QStringRK11QStringList(arg0, arg1)
  default:
    qtrt.ErrorResolve("QFont", "insertSubstitutions", args)
  }

}

// rawMode()
func (this *QFont) rawMode(args ...interface{}) () {
  // rawMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont7rawModeEv
    // invoke: bool rawMode()
    var ret = C.C_ZNK5QFont7rawModeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "rawMode", args)
  }

}

// setItalic(_Bool)
func (this *QFont) setItalic(args ...interface{}) () {
  // setItalic(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont9setItalicEb
    // invoke: void setItalic(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont9setItalicEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setItalic", args)
  }

}

// toString()
func (this *QFont) toString(args ...interface{}) () {
  // toString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont8toStringEv
    // invoke: QString toString()
    var ret = C.C_ZNK5QFont8toStringEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "toString", args)
  }

}

// underline()
func (this *QFont) underline(args ...interface{}) () {
  // underline()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont9underlineEv
    // invoke: bool underline()
    var ret = C.C_ZNK5QFont9underlineEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "underline", args)
  }

}

// setStyleName(const class QString &)
func (this *QFont) setStyleName(args ...interface{}) () {
  // setStyleName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont12setStyleNameERK7QString
    // invoke: void setStyleName(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont12setStyleNameERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setStyleName", args)
  }

}

// removeSubstitutions(const class QString &)
func (this *QFont) removeSubstitutions_s(args ...interface{}) () {
  // removeSubstitutions(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont19removeSubstitutionsERK7QString
    // invoke: void removeSubstitutions(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont19removeSubstitutionsERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QFont", "removeSubstitutions", args)
  }

}

// setPointSize(int)
func (this *QFont) setPointSize(args ...interface{}) () {
  // setPointSize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont12setPointSizeEi
    // invoke: void setPointSize(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont12setPointSizeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setPointSize", args)
  }

}

// style()
func (this *QFont) style(args ...interface{}) () {
  // style()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont5styleEv
    // invoke: QFont::Style style()
    C.C_ZNK5QFont5styleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "style", args)
  }

}

// family()
func (this *QFont) family(args ...interface{}) () {
  // family()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont6familyEv
    // invoke: QString family()
    var ret = C.C_ZNK5QFont6familyEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "family", args)
  }

}

// setBold(_Bool)
func (this *QFont) setBold(args ...interface{}) () {
  // setBold(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont7setBoldEb
    // invoke: void setBold(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont7setBoldEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setBold", args)
  }

}

// key()
func (this *QFont) key(args ...interface{}) () {
  // key()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont3keyEv
    // invoke: QString key()
    var ret = C.C_ZNK5QFont3keyEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "key", args)
  }

}

// stretch()
func (this *QFont) stretch(args ...interface{}) () {
  // stretch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont7stretchEv
    // invoke: int stretch()
    var ret = C.C_ZNK5QFont7stretchEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "stretch", args)
  }

}

// pointSizeF()
func (this *QFont) pointSizeF(args ...interface{}) () {
  // pointSizeF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont10pointSizeFEv
    // invoke: qreal pointSizeF()
    var ret = C.C_ZNK5QFont10pointSizeFEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "pointSizeF", args)
  }

}

// pointSize()
func (this *QFont) pointSize(args ...interface{}) () {
  // pointSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont9pointSizeEv
    // invoke: int pointSize()
    var ret = C.C_ZNK5QFont9pointSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "pointSize", args)
  }

}

// fixedPitch()
func (this *QFont) fixedPitch(args ...interface{}) () {
  // fixedPitch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont10fixedPitchEv
    // invoke: bool fixedPitch()
    var ret = C.C_ZNK5QFont10fixedPitchEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "fixedPitch", args)
  }

}

// substitutions()
func (this *QFont) substitutions_s(args ...interface{}) () {
  // substitutions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont13substitutionsEv
    // invoke: QStringList substitutions()
    C.C_ZN5QFont13substitutionsEv()
  default:
    qtrt.ErrorResolve("QFont", "substitutions", args)
  }

}

// cacheStatistics()
func (this *QFont) cacheStatistics_s(args ...interface{}) () {
  // cacheStatistics()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont15cacheStatisticsEv
    // invoke: void cacheStatistics()
    C.C_ZN5QFont15cacheStatisticsEv()
  default:
    qtrt.ErrorResolve("QFont", "cacheStatistics", args)
  }

}

// ~QFont()
func (this *QFont) FreeQFont(args ...interface{}) () {
  // ~QFont()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFontD0Ev
    // invoke: void ~QFont()
    C.C_ZN5QFontD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "~QFont", args)
  }

}

// kerning()
func (this *QFont) kerning(args ...interface{}) () {
  // kerning()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont7kerningEv
    // invoke: bool kerning()
    var ret = C.C_ZNK5QFont7kerningEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "kerning", args)
  }

}

// substitutes(const class QString &)
func (this *QFont) substitutes_s(args ...interface{}) () {
  // substitutes(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont11substitutesERK7QString
    // invoke: QStringList substitutes(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont11substitutesERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QFont", "substitutes", args)
  }

}

// setRawMode(_Bool)
func (this *QFont) setRawMode(args ...interface{}) () {
  // setRawMode(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont10setRawModeEb
    // invoke: void setRawMode(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont10setRawModeEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setRawMode", args)
  }

}

// italic()
func (this *QFont) italic(args ...interface{}) () {
  // italic()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont6italicEv
    // invoke: bool italic()
    var ret = C.C_ZNK5QFont6italicEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "italic", args)
  }

}

// pixelSize()
func (this *QFont) pixelSize(args ...interface{}) () {
  // pixelSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont9pixelSizeEv
    // invoke: int pixelSize()
    var ret = C.C_ZNK5QFont9pixelSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "pixelSize", args)
  }

}

// swap(class QFont &)
func (this *QFont) swap(args ...interface{}) () {
  // swap(class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont4swapERS_
    // invoke: void swap(class QFont &)
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "swap", args)
  }

}

// rawName()
func (this *QFont) rawName(args ...interface{}) () {
  // rawName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont7rawNameEv
    // invoke: QString rawName()
    var ret = C.C_ZNK5QFont7rawNameEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "rawName", args)
  }

}

// insertSubstitution(const class QString &, const class QString &)
func (this *QFont) insertSubstitution_s(args ...interface{}) () {
  // insertSubstitution(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont18insertSubstitutionERK7QStringS2_
    // invoke: void insertSubstitution(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN5QFont18insertSubstitutionERK7QStringS2_(arg0, arg1)
  default:
    qtrt.ErrorResolve("QFont", "insertSubstitution", args)
  }

}

// strikeOut()
func (this *QFont) strikeOut(args ...interface{}) () {
  // strikeOut()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont9strikeOutEv
    // invoke: bool strikeOut()
    var ret = C.C_ZNK5QFont9strikeOutEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "strikeOut", args)
  }

}

// bold()
func (this *QFont) bold(args ...interface{}) () {
  // bold()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont4boldEv
    // invoke: bool bold()
    var ret = C.C_ZNK5QFont4boldEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "bold", args)
  }

}

// capitalization()
func (this *QFont) capitalization(args ...interface{}) () {
  // capitalization()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont14capitalizationEv
    // invoke: QFont::Capitalization capitalization()
    C.C_ZNK5QFont14capitalizationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "capitalization", args)
  }

}

// styleHint()
func (this *QFont) styleHint(args ...interface{}) () {
  // styleHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont9styleHintEv
    // invoke: QFont::StyleHint styleHint()
    C.C_ZNK5QFont9styleHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "styleHint", args)
  }

}

// setKerning(_Bool)
func (this *QFont) setKerning(args ...interface{}) () {
  // setKerning(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont10setKerningEb
    // invoke: void setKerning(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont10setKerningEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setKerning", args)
  }

}

// QFont(const class QFont &)
func NewQFont(args ...interface{}) *QFont {
  // QFont(const class QFont &)
  // QFont()
  // QFont(const class QString &, int, int, _Bool)
  // QFont(const class QFont &, class QPaintDevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[2][3] = qtrt.BoolTy(false) // "bool"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QFont{}) // "const QFont &"
  vtys[3][1] = reflect.TypeOf(QPaintDevice{}) // "QPaintDevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFontC1ERKS_
    // invoke: void QFont(const class QFont &)
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QFontC2ERKS_(arg0)
    return &QFont{qclsinst:qthis}
  case 1:
    // invoke: _ZN5QFontC1Ev
    // invoke: void QFont()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QFontC2Ev()
    return &QFont{qclsinst:qthis}
  case 2:
    // invoke: _ZN5QFontC1ERK7QStringiib
    // invoke: void QFont(const class QString &, int, int, _Bool)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.bool(args[3].(bool))
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QFontC2ERK7QStringiib(arg0, arg1, arg2, arg3)
    return &QFont{qclsinst:qthis}
  case 3:
    // invoke: _ZN5QFontC1ERKS_P12QPaintDevice
    // invoke: void QFont(const class QFont &, class QPaintDevice *)
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPaintDevice).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QFontC2ERKS_P12QPaintDevice(arg0, arg1)
    return &QFont{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QFont", "QFont", args)
  }

  return nil // QFont{qclsinst:qthis}
}

// substitute(const class QString &)
func (this *QFont) substitute_s(args ...interface{}) () {
  // substitute(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont10substituteERK7QString
    // invoke: QString substitute(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN5QFont10substituteERK7QString(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "substitute", args)
  }

}

// hintingPreference()
func (this *QFont) hintingPreference(args ...interface{}) () {
  // hintingPreference()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont17hintingPreferenceEv
    // invoke: QFont::HintingPreference hintingPreference()
    C.C_ZNK5QFont17hintingPreferenceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "hintingPreference", args)
  }

}

// setRawName(const class QString &)
func (this *QFont) setRawName(args ...interface{}) () {
  // setRawName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont10setRawNameERK7QString
    // invoke: void setRawName(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont10setRawNameERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setRawName", args)
  }

}

// setOverline(_Bool)
func (this *QFont) setOverline(args ...interface{}) () {
  // setOverline(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont11setOverlineEb
    // invoke: void setOverline(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont11setOverlineEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setOverline", args)
  }

}

// setFixedPitch(_Bool)
func (this *QFont) setFixedPitch(args ...interface{}) () {
  // setFixedPitch(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont13setFixedPitchEb
    // invoke: void setFixedPitch(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont13setFixedPitchEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setFixedPitch", args)
  }

}

// letterSpacing()
func (this *QFont) letterSpacing(args ...interface{}) () {
  // letterSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont13letterSpacingEv
    // invoke: qreal letterSpacing()
    var ret = C.C_ZNK5QFont13letterSpacingEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "letterSpacing", args)
  }

}

// initialize()
func (this *QFont) initialize_s(args ...interface{}) () {
  // initialize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont10initializeEv
    // invoke: void initialize()
    C.C_ZN5QFont10initializeEv()
  default:
    qtrt.ErrorResolve("QFont", "initialize", args)
  }

}

// setPixelSize(int)
func (this *QFont) setPixelSize(args ...interface{}) () {
  // setPixelSize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont12setPixelSizeEi
    // invoke: void setPixelSize(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont12setPixelSizeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setPixelSize", args)
  }

}

// resolve(uint)
func (this *QFont) resolve(args ...interface{}) () {
  // resolve(uint)
  // resolve()
  // resolve(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont7resolveEj
    // invoke: void resolve(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont7resolveEj(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK5QFont7resolveEv
    // invoke: uint resolve()
    var ret = C.C_ZNK5QFont7resolveEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  case 2:
    // invoke: _ZNK5QFont7resolveERKS_
    // invoke: QFont resolve(const class QFont &)
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK5QFont7resolveERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "resolve", args)
  }

}

// exactMatch()
func (this *QFont) exactMatch(args ...interface{}) () {
  // exactMatch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont10exactMatchEv
    // invoke: bool exactMatch()
    var ret = C.C_ZNK5QFont10exactMatchEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "exactMatch", args)
  }

}

// setWordSpacing(qreal)
func (this *QFont) setWordSpacing(args ...interface{}) () {
  // setWordSpacing(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont14setWordSpacingEd
    // invoke: void setWordSpacing(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont14setWordSpacingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setWordSpacing", args)
  }

}

// setStrikeOut(_Bool)
func (this *QFont) setStrikeOut(args ...interface{}) () {
  // setStrikeOut(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont12setStrikeOutEb
    // invoke: void setStrikeOut(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont12setStrikeOutEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setStrikeOut", args)
  }

}

// setWeight(int)
func (this *QFont) setWeight(args ...interface{}) () {
  // setWeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont9setWeightEi
    // invoke: void setWeight(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont9setWeightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setWeight", args)
  }

}

// lastResortFont()
func (this *QFont) lastResortFont(args ...interface{}) () {
  // lastResortFont()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont14lastResortFontEv
    // invoke: QString lastResortFont()
    var ret = C.C_ZNK5QFont14lastResortFontEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "lastResortFont", args)
  }

}

// overline()
func (this *QFont) overline(args ...interface{}) () {
  // overline()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont8overlineEv
    // invoke: bool overline()
    var ret = C.C_ZNK5QFont8overlineEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "overline", args)
  }

}

// defaultFamily()
func (this *QFont) defaultFamily(args ...interface{}) () {
  // defaultFamily()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont13defaultFamilyEv
    // invoke: QString defaultFamily()
    var ret = C.C_ZNK5QFont13defaultFamilyEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "defaultFamily", args)
  }

}

// setFamily(const class QString &)
func (this *QFont) setFamily(args ...interface{}) () {
  // setFamily(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont9setFamilyERK7QString
    // invoke: void setFamily(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont9setFamilyERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setFamily", args)
  }

}

// setUnderline(_Bool)
func (this *QFont) setUnderline(args ...interface{}) () {
  // setUnderline(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont12setUnderlineEb
    // invoke: void setUnderline(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont12setUnderlineEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setUnderline", args)
  }

}

// styleName()
func (this *QFont) styleName(args ...interface{}) () {
  // styleName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont9styleNameEv
    // invoke: QString styleName()
    var ret = C.C_ZNK5QFont9styleNameEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFont", "styleName", args)
  }

}

// cleanup()
func (this *QFont) cleanup_s(args ...interface{}) () {
  // cleanup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFont7cleanupEv
    // invoke: void cleanup()
    C.C_ZN5QFont7cleanupEv()
  default:
    qtrt.ErrorResolve("QFont", "cleanup", args)
  }

}

// <= body block end

