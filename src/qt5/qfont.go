package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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
extern bool C_ZNK5QFont8isCopyOfERKS_(void* qthis, void* arg0); // 4
  // proto:  qreal QFont::wordSpacing();
extern double C_ZNK5QFont11wordSpacingEv(void* qthis); // 4
  // proto:  QFont::SpacingType QFont::letterSpacingType();
extern void C_ZNK5QFont17letterSpacingTypeEv(void* qthis); // 4
  // proto:  QString QFont::lastResortFamily();
extern void* C_ZNK5QFont16lastResortFamilyEv(void* qthis); // 4
  // proto:  void QFont::setStretch(int );
extern void C_ZN5QFont10setStretchEi(void* qthis, int32_t arg0); // 4
  // proto:  int QFont::weight();
extern int32_t C_ZNK5QFont6weightEv(void* qthis); // 4
  // proto:  QFont::StyleStrategy QFont::styleStrategy();
extern void C_ZNK5QFont13styleStrategyEv(void* qthis); // 4
  // proto:  bool QFont::fromString(const QString & );
extern bool C_ZN5QFont10fromStringERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QFont::setPointSizeF(qreal );
extern void C_ZN5QFont13setPointSizeFEd(void* qthis, double arg0); // 4
  // proto: static void QFont::insertSubstitutions(const QString & , const QStringList & );
extern void C_ZN5QFont19insertSubstitutionsERK7QStringRK11QStringList(void* arg0, void* arg1); // 4
  // proto:  bool QFont::rawMode();
extern bool C_ZNK5QFont7rawModeEv(void* qthis); // 4
  // proto:  void QFont::setItalic(bool b);
extern void C_ZN5QFont9setItalicEb(void* qthis, bool arg0); // 2
  // proto:  QString QFont::toString();
extern void* C_ZNK5QFont8toStringEv(void* qthis); // 4
  // proto:  bool QFont::underline();
extern bool C_ZNK5QFont9underlineEv(void* qthis); // 4
  // proto:  void QFont::setStyleName(const QString & );
extern void C_ZN5QFont12setStyleNameERK7QString(void* qthis, void* arg0); // 4
  // proto: static void QFont::removeSubstitutions(const QString & );
extern void C_ZN5QFont19removeSubstitutionsERK7QString(void* arg0); // 4
  // proto:  void QFont::setPointSize(int );
extern void C_ZN5QFont12setPointSizeEi(void* qthis, int32_t arg0); // 4
  // proto:  QFont::Style QFont::style();
extern void C_ZNK5QFont5styleEv(void* qthis); // 4
  // proto:  QString QFont::family();
extern void* C_ZNK5QFont6familyEv(void* qthis); // 4
  // proto:  void QFont::setBold(bool );
extern void C_ZN5QFont7setBoldEb(void* qthis, bool arg0); // 2
  // proto:  QString QFont::key();
extern void* C_ZNK5QFont3keyEv(void* qthis); // 4
  // proto:  int QFont::stretch();
extern int32_t C_ZNK5QFont7stretchEv(void* qthis); // 4
  // proto:  qreal QFont::pointSizeF();
extern double C_ZNK5QFont10pointSizeFEv(void* qthis); // 4
  // proto:  int QFont::pointSize();
extern int32_t C_ZNK5QFont9pointSizeEv(void* qthis); // 4
  // proto:  bool QFont::fixedPitch();
extern bool C_ZNK5QFont10fixedPitchEv(void* qthis); // 4
  // proto: static QStringList QFont::substitutions();
extern void C_ZN5QFont13substitutionsEv(); // 4
  // proto: static void QFont::cacheStatistics();
extern void C_ZN5QFont15cacheStatisticsEv(); // 4
  // proto:  void QFont::~QFont();
extern void C_ZN5QFontD2Ev(void* qthis); // 4
  // proto:  bool QFont::kerning();
extern bool C_ZNK5QFont7kerningEv(void* qthis); // 4
  // proto: static QStringList QFont::substitutes(const QString & );
extern void C_ZN5QFont11substitutesERK7QString(void* arg0); // 4
  // proto:  void QFont::setRawMode(bool );
extern void C_ZN5QFont10setRawModeEb(void* qthis, bool arg0); // 4
  // proto:  bool QFont::italic();
extern bool C_ZNK5QFont6italicEv(void* qthis); // 2
  // proto:  int QFont::pixelSize();
extern int32_t C_ZNK5QFont9pixelSizeEv(void* qthis); // 4
  // proto:  void QFont::swap(QFont & other);
extern void C_ZN5QFont4swapERS_(void* qthis, void* arg0); // 2
  // proto:  QString QFont::rawName();
extern void* C_ZNK5QFont7rawNameEv(void* qthis); // 4
  // proto: static void QFont::insertSubstitution(const QString & , const QString & );
extern void C_ZN5QFont18insertSubstitutionERK7QStringS2_(void* arg0, void* arg1); // 4
  // proto:  bool QFont::strikeOut();
extern bool C_ZNK5QFont9strikeOutEv(void* qthis); // 4
  // proto:  bool QFont::bold();
extern bool C_ZNK5QFont4boldEv(void* qthis); // 2
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
extern void* C_ZN5QFont10substituteERK7QString(void* arg0); // 4
  // proto:  QFont::HintingPreference QFont::hintingPreference();
extern void C_ZNK5QFont17hintingPreferenceEv(void* qthis); // 4
  // proto:  void QFont::setRawName(const QString & );
extern void C_ZN5QFont10setRawNameERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QFont::setOverline(bool );
extern void C_ZN5QFont11setOverlineEb(void* qthis, bool arg0); // 4
  // proto:  void QFont::setFixedPitch(bool );
extern void C_ZN5QFont13setFixedPitchEb(void* qthis, bool arg0); // 4
  // proto:  qreal QFont::letterSpacing();
extern double C_ZNK5QFont13letterSpacingEv(void* qthis); // 4
  // proto: static void QFont::initialize();
extern void C_ZN5QFont10initializeEv(); // 4
  // proto:  void QFont::setPixelSize(int );
extern void C_ZN5QFont12setPixelSizeEi(void* qthis, int32_t arg0); // 4
  // proto:  void QFont::resolve(uint mask);
extern void C_ZN5QFont7resolveEj(void* qthis, int32_t arg0); // 2
  // proto:  uint QFont::resolve();
extern int32_t C_ZNK5QFont7resolveEv(void* qthis); // 2
  // proto:  QFont QFont::resolve(const QFont & );
extern void* C_ZNK5QFont7resolveERKS_(void* qthis, void* arg0); // 4
  // proto:  bool QFont::exactMatch();
extern bool C_ZNK5QFont10exactMatchEv(void* qthis); // 4
  // proto:  void QFont::setWordSpacing(qreal spacing);
extern void C_ZN5QFont14setWordSpacingEd(void* qthis, double arg0); // 4
  // proto:  void QFont::setStrikeOut(bool );
extern void C_ZN5QFont12setStrikeOutEb(void* qthis, bool arg0); // 4
  // proto:  void QFont::setWeight(int );
extern void C_ZN5QFont9setWeightEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QFont::lastResortFont();
extern void* C_ZNK5QFont14lastResortFontEv(void* qthis); // 4
  // proto:  bool QFont::overline();
extern bool C_ZNK5QFont8overlineEv(void* qthis); // 4
  // proto:  QString QFont::defaultFamily();
extern void* C_ZNK5QFont13defaultFamilyEv(void* qthis); // 4
  // proto:  void QFont::setFamily(const QString & );
extern void C_ZN5QFont9setFamilyERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QFont::setUnderline(bool );
extern void C_ZN5QFont12setUnderlineEb(void* qthis, bool arg0); // 4
  // proto:  QString QFont::styleName();
extern void* C_ZNK5QFont9styleNameEv(void* qthis); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// isCopyOf(const class QFont &)
func (this *QFont) Iscopyof(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QFont).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK5QFont8isCopyOfERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFont", "isCopyOf", args)
  }

  return
}

// wordSpacing()
func (this *QFont) Wordspacing(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QFont11wordSpacingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFont", "wordSpacing", args)
  }

  return
}

// letterSpacingType()
func (this *QFont) Letterspacingtype(args ...interface{}) () {
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
    C.C_ZNK5QFont17letterSpacingTypeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "letterSpacingType", args)
  }

  return
}

// lastResortFamily()
func (this *QFont) Lastresortfamily(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QFont16lastResortFamilyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFont", "lastResortFamily", args)
  }

  return
}

// setStretch(int)
func (this *QFont) Setstretch(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont10setStretchEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setStretch", args)
  }

  return
}

// weight()
func (this *QFont) Weight(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QFont6weightEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFont", "weight", args)
  }

  return
}

// styleStrategy()
func (this *QFont) Stylestrategy(args ...interface{}) () {
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
    C.C_ZNK5QFont13styleStrategyEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "styleStrategy", args)
  }

  return
}

// fromString(const class QString &)
func (this *QFont) Fromstring(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QFont10fromStringERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFont", "fromString", args)
  }

  return
}

// setPointSizeF(qreal)
func (this *QFont) Setpointsizef(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont13setPointSizeFEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setPointSizeF", args)
  }

  return
}

// insertSubstitutions(const class QString &, const class QStringList &)
func (this *QFont) Insertsubstitutions_S(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStringList).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN5QFont19insertSubstitutionsERK7QStringRK11QStringList(arg0, arg1)
  default:
    qtrt.ErrorResolve("QFont", "insertSubstitutions", args)
  }

  return
}

// rawMode()
func (this *QFont) Rawmode(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QFont7rawModeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFont", "rawMode", args)
  }

  return
}

// setItalic(_Bool)
func (this *QFont) Setitalic(args ...interface{}) () {
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
    C.C_ZN5QFont9setItalicEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setItalic", args)
  }

  return
}

// toString()
func (this *QFont) Tostring(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QFont8toStringEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFont", "toString", args)
  }

  return
}

// underline()
func (this *QFont) Underline(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QFont9underlineEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFont", "underline", args)
  }

  return
}

// setStyleName(const class QString &)
func (this *QFont) Setstylename(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont12setStyleNameERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setStyleName", args)
  }

  return
}

// removeSubstitutions(const class QString &)
func (this *QFont) Removesubstitutions_S(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont19removeSubstitutionsERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QFont", "removeSubstitutions", args)
  }

  return
}

// setPointSize(int)
func (this *QFont) Setpointsize(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont12setPointSizeEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setPointSize", args)
  }

  return
}

// style()
func (this *QFont) Style(args ...interface{}) () {
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
    C.C_ZNK5QFont5styleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "style", args)
  }

  return
}

// family()
func (this *QFont) Family(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QFont6familyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFont", "family", args)
  }

  return
}

// setBold(_Bool)
func (this *QFont) Setbold(args ...interface{}) () {
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
    C.C_ZN5QFont7setBoldEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setBold", args)
  }

  return
}

// key()
func (this *QFont) Key(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QFont3keyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFont", "key", args)
  }

  return
}

// stretch()
func (this *QFont) Stretch(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QFont7stretchEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFont", "stretch", args)
  }

  return
}

// pointSizeF()
func (this *QFont) Pointsizef(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QFont10pointSizeFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFont", "pointSizeF", args)
  }

  return
}

// pointSize()
func (this *QFont) Pointsize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QFont9pointSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFont", "pointSize", args)
  }

  return
}

// fixedPitch()
func (this *QFont) Fixedpitch(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QFont10fixedPitchEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFont", "fixedPitch", args)
  }

  return
}

// substitutions()
func (this *QFont) Substitutions_S(args ...interface{}) () {
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

  return
}

// cacheStatistics()
func (this *QFont) Cachestatistics_S(args ...interface{}) () {
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

  return
}

// ~QFont()
func (this *QFont) Freeqfont(args ...interface{}) () {
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
    C.C_ZN5QFontD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "~QFont", args)
  }

  return
}

// kerning()
func (this *QFont) Kerning(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QFont7kerningEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFont", "kerning", args)
  }

  return
}

// substitutes(const class QString &)
func (this *QFont) Substitutes_S(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont11substitutesERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QFont", "substitutes", args)
  }

  return
}

// setRawMode(_Bool)
func (this *QFont) Setrawmode(args ...interface{}) () {
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
    C.C_ZN5QFont10setRawModeEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setRawMode", args)
  }

  return
}

// italic()
func (this *QFont) Italic(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QFont6italicEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFont", "italic", args)
  }

  return
}

// pixelSize()
func (this *QFont) Pixelsize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QFont9pixelSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFont", "pixelSize", args)
  }

  return
}

// swap(class QFont &)
func (this *QFont) Swap(args ...interface{}) () {
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
    var arg0 = args[0].(*QFont).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "swap", args)
  }

  return
}

// rawName()
func (this *QFont) Rawname(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QFont7rawNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFont", "rawName", args)
  }

  return
}

// insertSubstitution(const class QString &, const class QString &)
func (this *QFont) Insertsubstitution_S(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN5QFont18insertSubstitutionERK7QStringS2_(arg0, arg1)
  default:
    qtrt.ErrorResolve("QFont", "insertSubstitution", args)
  }

  return
}

// strikeOut()
func (this *QFont) Strikeout(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QFont9strikeOutEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFont", "strikeOut", args)
  }

  return
}

// bold()
func (this *QFont) Bold(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QFont4boldEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFont", "bold", args)
  }

  return
}

// capitalization()
func (this *QFont) Capitalization(args ...interface{}) () {
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
    C.C_ZNK5QFont14capitalizationEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "capitalization", args)
  }

  return
}

// styleHint()
func (this *QFont) Stylehint(args ...interface{}) () {
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
    C.C_ZNK5QFont9styleHintEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "styleHint", args)
  }

  return
}

// setKerning(_Bool)
func (this *QFont) Setkerning(args ...interface{}) () {
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
    C.C_ZN5QFont10setKerningEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setKerning", args)
  }

  return
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
    var arg0 = args[0].(*QFont).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QFontC2ERKS_(arg0)
    return &QFont{Qclsinst:qthis}
  case 1:
    // invoke: _ZN5QFontC1Ev
    // invoke: void QFont()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QFontC2Ev()
    return &QFont{Qclsinst:qthis}
  case 2:
    // invoke: _ZN5QFontC1ERK7QStringiib
    // invoke: void QFont(const class QString &, int, int, _Bool)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.bool(args[3].(bool))
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QFontC2ERK7QStringiib(arg0, arg1, arg2, arg3)
    return &QFont{Qclsinst:qthis}
  case 3:
    // invoke: _ZN5QFontC1ERKS_P12QPaintDevice
    // invoke: void QFont(const class QFont &, class QPaintDevice *)
    var arg0 = args[0].(*QFont).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPaintDevice).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QFontC2ERKS_P12QPaintDevice(arg0, arg1)
    return &QFont{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QFont", "QFont", args)
  }

  return nil // QFont{Qclsinst:qthis}
}

// substitute(const class QString &)
func (this *QFont) Substitute_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QFont10substituteERK7QString(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFont", "substitute", args)
  }

  return
}

// hintingPreference()
func (this *QFont) Hintingpreference(args ...interface{}) () {
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
    C.C_ZNK5QFont17hintingPreferenceEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "hintingPreference", args)
  }

  return
}

// setRawName(const class QString &)
func (this *QFont) Setrawname(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont10setRawNameERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setRawName", args)
  }

  return
}

// setOverline(_Bool)
func (this *QFont) Setoverline(args ...interface{}) () {
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
    C.C_ZN5QFont11setOverlineEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setOverline", args)
  }

  return
}

// setFixedPitch(_Bool)
func (this *QFont) Setfixedpitch(args ...interface{}) () {
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
    C.C_ZN5QFont13setFixedPitchEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setFixedPitch", args)
  }

  return
}

// letterSpacing()
func (this *QFont) Letterspacing(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QFont13letterSpacingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFont", "letterSpacing", args)
  }

  return
}

// initialize()
func (this *QFont) Initialize_S(args ...interface{}) () {
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

  return
}

// setPixelSize(int)
func (this *QFont) Setpixelsize(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont12setPixelSizeEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setPixelSize", args)
  }

  return
}

// resolve(uint)
func (this *QFont) Resolve(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont7resolveEj(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZNK5QFont7resolveEv
    // invoke: uint resolve()
    var ret0 = C.C_ZNK5QFont7resolveEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
  case 2:
    // invoke: _ZNK5QFont7resolveERKS_
    // invoke: QFont resolve(const class QFont &)
    var arg0 = args[0].(*QFont).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK5QFont7resolveERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
  default:
    qtrt.ErrorResolve("QFont", "resolve", args)
  }

  return
}

// exactMatch()
func (this *QFont) Exactmatch(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QFont10exactMatchEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFont", "exactMatch", args)
  }

  return
}

// setWordSpacing(qreal)
func (this *QFont) Setwordspacing(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont14setWordSpacingEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setWordSpacing", args)
  }

  return
}

// setStrikeOut(_Bool)
func (this *QFont) Setstrikeout(args ...interface{}) () {
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
    C.C_ZN5QFont12setStrikeOutEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setStrikeOut", args)
  }

  return
}

// setWeight(int)
func (this *QFont) Setweight(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont9setWeightEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setWeight", args)
  }

  return
}

// lastResortFont()
func (this *QFont) Lastresortfont(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QFont14lastResortFontEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFont", "lastResortFont", args)
  }

  return
}

// overline()
func (this *QFont) Overline(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QFont8overlineEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFont", "overline", args)
  }

  return
}

// defaultFamily()
func (this *QFont) Defaultfamily(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QFont13defaultFamilyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFont", "defaultFamily", args)
  }

  return
}

// setFamily(const class QString &)
func (this *QFont) Setfamily(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QFont9setFamilyERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setFamily", args)
  }

  return
}

// setUnderline(_Bool)
func (this *QFont) Setunderline(args ...interface{}) () {
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
    C.C_ZN5QFont12setUnderlineEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setUnderline", args)
  }

  return
}

// styleName()
func (this *QFont) Stylename(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QFont9styleNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFont", "styleName", args)
  }

  return
}

// cleanup()
func (this *QFont) Cleanup_S(args ...interface{}) () {
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

  return
}

// <= body block end

