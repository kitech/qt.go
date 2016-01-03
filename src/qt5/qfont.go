package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
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
  // proto:  void QFont::setWordSpacing(qreal spacing);
extern void _ZN5QFont14setWordSpacingEd(void* qthis, double arg0);
  // proto:  QString QFont::rawName();
extern void _ZNK5QFont7rawNameEv(void* qthis);
  // proto:  void QFont::setRawMode(bool );
extern void _ZN5QFont10setRawModeEb(void* qthis, bool arg0);
  // proto:  void QFont::setStyleName(const QString & );
extern void _ZN5QFont12setStyleNameERK7QString(void* qthis, void* arg0);
  // proto:  QFont QFont::resolve(const QFont & );
extern void _ZNK5QFont7resolveERKS_(void* qthis, void* arg0);
  // proto:  bool QFont::strikeOut();
extern void _ZNK5QFont9strikeOutEv(void* qthis);
  // proto:  int QFont::pixelSize();
extern void _ZNK5QFont9pixelSizeEv(void* qthis);
  // proto:  void QFont::setWeight(int );
extern void _ZN5QFont9setWeightEi(void* qthis, int32_t arg0);
  // proto:  int QFont::weight();
extern void _ZNK5QFont6weightEv(void* qthis);
  // proto: static void QFont::insertSubstitutions(const QString & , const QStringList & );
extern void _ZN5QFont19insertSubstitutionsERK7QStringRK11QStringList(void* arg0, void* arg1);
  // proto:  bool QFont::kerning();
extern void _ZNK5QFont7kerningEv(void* qthis);
  // proto: static QStringList QFont::substitutions();
extern void _ZN5QFont13substitutionsEv();
  // proto:  bool QFont::italic();
extern void demth_ZNK5QFont6italicEv(void* qthis);
  // proto:  void QFont::setUnderline(bool );
extern void _ZN5QFont12setUnderlineEb(void* qthis, bool arg0);
  // proto:  qreal QFont::letterSpacing();
extern void _ZNK5QFont13letterSpacingEv(void* qthis);
  // proto:  void QFont::setPointSize(int );
extern void _ZN5QFont12setPointSizeEi(void* qthis, int32_t arg0);
  // proto:  void QFont::QFont(const QString & family, int pointSize, int weight, bool italic);
extern void* dector_ZN5QFontC1ERK7QStringiib(void* arg0, int32_t arg1, int32_t arg2, bool arg3);
extern void _ZN5QFontC1ERK7QStringiib(void* qthis, void* arg0, int32_t arg1, int32_t arg2, bool arg3);
  // proto:  void QFont::setOverline(bool );
extern void _ZN5QFont11setOverlineEb(void* qthis, bool arg0);
  // proto:  QString QFont::family();
extern void _ZNK5QFont6familyEv(void* qthis);
  // proto:  QString QFont::lastResortFamily();
extern void _ZNK5QFont16lastResortFamilyEv(void* qthis);
  // proto:  void QFont::setItalic(bool b);
extern void demth_ZN5QFont9setItalicEb(void* qthis, bool arg0);
  // proto:  void QFont::setFamily(const QString & );
extern void _ZN5QFont9setFamilyERK7QString(void* qthis, void* arg0);
  // proto:  void QFont::QFont(const QFont & );
extern void* dector_ZN5QFontC1ERKS_(void* arg0);
extern void _ZN5QFontC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QFont::overline();
extern void _ZNK5QFont8overlineEv(void* qthis);
  // proto:  void QFont::~QFont();
extern void _ZN5QFontD0Ev(void* qthis);
  // proto:  void QFont::resolve(uint mask);
extern void demth_ZN5QFont7resolveEj(void* qthis, int32_t arg0);
  // proto:  void QFont::setBold(bool );
extern void demth_ZN5QFont7setBoldEb(void* qthis, bool arg0);
  // proto: static void QFont::cacheStatistics();
extern void _ZN5QFont15cacheStatisticsEv();
  // proto:  void QFont::setPointSizeF(qreal );
extern void _ZN5QFont13setPointSizeFEd(void* qthis, double arg0);
  // proto: static QStringList QFont::substitutes(const QString & );
extern void _ZN5QFont11substitutesERK7QString(void* arg0);
  // proto:  qreal QFont::wordSpacing();
extern void _ZNK5QFont11wordSpacingEv(void* qthis);
  // proto:  QString QFont::toString();
extern void _ZNK5QFont8toStringEv(void* qthis);
  // proto:  qreal QFont::pointSizeF();
extern void _ZNK5QFont10pointSizeFEv(void* qthis);
  // proto: static void QFont::insertSubstitution(const QString & , const QString & );
extern void _ZN5QFont18insertSubstitutionERK7QStringS2_(void* arg0, void* arg1);
  // proto:  void QFont::setStretch(int );
extern void _ZN5QFont10setStretchEi(void* qthis, int32_t arg0);
  // proto:  QString QFont::styleName();
extern void _ZNK5QFont9styleNameEv(void* qthis);
  // proto:  void QFont::QFont();
extern void* dector_ZN5QFontC1Ev();
extern void _ZN5QFontC1Ev(void* qthis);
  // proto:  bool QFont::rawMode();
extern void _ZNK5QFont7rawModeEv(void* qthis);
  // proto:  bool QFont::fromString(const QString & );
extern void _ZN5QFont10fromStringERK7QString(void* qthis, void* arg0);
  // proto:  bool QFont::underline();
extern void _ZNK5QFont9underlineEv(void* qthis);
  // proto:  bool QFont::isCopyOf(const QFont & );
extern void _ZNK5QFont8isCopyOfERKS_(void* qthis, void* arg0);
  // proto:  int QFont::pointSize();
extern void _ZNK5QFont9pointSizeEv(void* qthis);
  // proto:  void QFont::setKerning(bool );
extern void _ZN5QFont10setKerningEb(void* qthis, bool arg0);
  // proto:  bool QFont::bold();
extern void demth_ZNK5QFont4boldEv(void* qthis);
  // proto:  bool QFont::fixedPitch();
extern void _ZNK5QFont10fixedPitchEv(void* qthis);
  // proto:  void QFont::QFont(const QFont & , QPaintDevice * pd);
extern void* dector_ZN5QFontC1ERKS_P12QPaintDevice(void* arg0, void* arg1);
extern void _ZN5QFontC1ERKS_P12QPaintDevice(void* qthis, void* arg0, void* arg1);
  // proto: static QString QFont::substitute(const QString & );
extern void _ZN5QFont10substituteERK7QString(void* arg0);
  // proto:  void QFont::setFixedPitch(bool );
extern void _ZN5QFont13setFixedPitchEb(void* qthis, bool arg0);
  // proto: static void QFont::removeSubstitutions(const QString & );
extern void _ZN5QFont19removeSubstitutionsERK7QString(void* arg0);
  // proto:  void QFont::setPixelSize(int );
extern void _ZN5QFont12setPixelSizeEi(void* qthis, int32_t arg0);
  // proto: static void QFont::initialize();
extern void _ZN5QFont10initializeEv();
  // proto:  QString QFont::key();
extern void _ZNK5QFont3keyEv(void* qthis);
  // proto:  QString QFont::lastResortFont();
extern void _ZNK5QFont14lastResortFontEv(void* qthis);
  // proto:  void QFont::swap(QFont & other);
extern void demth_ZN5QFont4swapERS_(void* qthis, void* arg0);
  // proto:  QString QFont::defaultFamily();
extern void _ZNK5QFont13defaultFamilyEv(void* qthis);
  // proto:  void QFont::setStrikeOut(bool );
extern void _ZN5QFont12setStrikeOutEb(void* qthis, bool arg0);
  // proto:  uint QFont::resolve();
extern void demth_ZNK5QFont7resolveEv(void* qthis);
  // proto: static void QFont::cleanup();
extern void _ZN5QFont7cleanupEv();
  // proto:  bool QFont::exactMatch();
extern void _ZNK5QFont10exactMatchEv(void* qthis);
  // proto:  int QFont::stretch();
extern void _ZNK5QFont7stretchEv(void* qthis);
  // proto:  void QFont::setRawName(const QString & );
extern void _ZN5QFont10setRawNameERK7QString(void* qthis, void* arg0);
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

  // proto:  void QFont::setWordSpacing(qreal spacing);
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
    C._ZN5QFont14setWordSpacingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setWordSpacing", args)
  }

}

  // proto:  QString QFont::rawName();
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
    C._ZNK5QFont7rawNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "rawName", args)
  }

}

  // proto:  void QFont::setRawMode(bool );
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
    C._ZN5QFont10setRawModeEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setRawMode", args)
  }

}

  // proto:  void QFont::setStyleName(const QString & );
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
    C._ZN5QFont12setStyleNameERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setStyleName", args)
  }

}

  // proto:  QFont QFont::resolve(const QFont & );
func (this *QFont) resolve(args ...interface{}) () {
  // resolve(const class QFont &)
  // resolve(uint)
  // resolve()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"
  vtys[2] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QFont7resolveERKS_
    // invoke: QFont resolve(const class QFont &)
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK5QFont7resolveERKS_(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN5QFont7resolveEj
    // invoke: void resolve(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.demth_ZN5QFont7resolveEj(this.qclsinst, arg0)
  case 2:
    // invoke: _ZNK5QFont7resolveEv
    // invoke: uint resolve()
    C.demth_ZNK5QFont7resolveEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "resolve", args)
  }

}

  // proto:  bool QFont::strikeOut();
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
    C._ZNK5QFont9strikeOutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "strikeOut", args)
  }

}

  // proto:  int QFont::pixelSize();
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
    C._ZNK5QFont9pixelSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "pixelSize", args)
  }

}

  // proto:  void QFont::setWeight(int );
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
    C._ZN5QFont9setWeightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setWeight", args)
  }

}

  // proto:  int QFont::weight();
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
    C._ZNK5QFont6weightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "weight", args)
  }

}

  // proto: static void QFont::insertSubstitutions(const QString & , const QStringList & );
func (this *QFont) insertSubstitutions_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFont", "insertSubstitutions", args)
  }

}

  // proto:  bool QFont::kerning();
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
    C._ZNK5QFont7kerningEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "kerning", args)
  }

}

  // proto: static QStringList QFont::substitutions();
func (this *QFont) substitutions_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFont", "substitutions", args)
  }

}

  // proto:  bool QFont::italic();
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
    C.demth_ZNK5QFont6italicEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "italic", args)
  }

}

  // proto:  void QFont::setUnderline(bool );
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
    C._ZN5QFont12setUnderlineEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setUnderline", args)
  }

}

  // proto:  qreal QFont::letterSpacing();
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
    C._ZNK5QFont13letterSpacingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "letterSpacing", args)
  }

}

  // proto:  void QFont::setPointSize(int );
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
    C._ZN5QFont12setPointSizeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setPointSize", args)
  }

}

  // proto:  void QFont::QFont(const QString & family, int pointSize, int weight, bool italic);
func NewQFont(args ...interface{}) QFont {
  return QFont{}
}

  // proto:  void QFont::setOverline(bool );
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
    C._ZN5QFont11setOverlineEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setOverline", args)
  }

}

  // proto:  QString QFont::family();
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
    C._ZNK5QFont6familyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "family", args)
  }

}

  // proto:  QString QFont::lastResortFamily();
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
    C._ZNK5QFont16lastResortFamilyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "lastResortFamily", args)
  }

}

  // proto:  void QFont::setItalic(bool b);
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
    C.demth_ZN5QFont9setItalicEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setItalic", args)
  }

}

  // proto:  void QFont::setFamily(const QString & );
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
    C._ZN5QFont9setFamilyERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setFamily", args)
  }

}

  // proto:  bool QFont::overline();
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
    C._ZNK5QFont8overlineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "overline", args)
  }

}

  // proto:  void QFont::~QFont();
func (this *QFont) FreeQFont(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFont", "~QFont", args)
  }

}

  // proto:  void QFont::setBold(bool );
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
    C.demth_ZN5QFont7setBoldEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setBold", args)
  }

}

  // proto: static void QFont::cacheStatistics();
func (this *QFont) cacheStatistics_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFont", "cacheStatistics", args)
  }

}

  // proto:  void QFont::setPointSizeF(qreal );
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
    C._ZN5QFont13setPointSizeFEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setPointSizeF", args)
  }

}

  // proto: static QStringList QFont::substitutes(const QString & );
func (this *QFont) substitutes_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFont", "substitutes", args)
  }

}

  // proto:  qreal QFont::wordSpacing();
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
    C._ZNK5QFont11wordSpacingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "wordSpacing", args)
  }

}

  // proto:  QString QFont::toString();
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
    C._ZNK5QFont8toStringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "toString", args)
  }

}

  // proto:  qreal QFont::pointSizeF();
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
    C._ZNK5QFont10pointSizeFEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "pointSizeF", args)
  }

}

  // proto: static void QFont::insertSubstitution(const QString & , const QString & );
func (this *QFont) insertSubstitution_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFont", "insertSubstitution", args)
  }

}

  // proto:  void QFont::setStretch(int );
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
    C._ZN5QFont10setStretchEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setStretch", args)
  }

}

  // proto:  QString QFont::styleName();
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
    C._ZNK5QFont9styleNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "styleName", args)
  }

}

  // proto:  bool QFont::rawMode();
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
    C._ZNK5QFont7rawModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "rawMode", args)
  }

}

  // proto:  bool QFont::fromString(const QString & );
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
    C._ZN5QFont10fromStringERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "fromString", args)
  }

}

  // proto:  bool QFont::underline();
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
    C._ZNK5QFont9underlineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "underline", args)
  }

}

  // proto:  bool QFont::isCopyOf(const QFont & );
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
    C._ZNK5QFont8isCopyOfERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "isCopyOf", args)
  }

}

  // proto:  int QFont::pointSize();
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
    C._ZNK5QFont9pointSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "pointSize", args)
  }

}

  // proto:  void QFont::setKerning(bool );
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
    C._ZN5QFont10setKerningEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setKerning", args)
  }

}

  // proto:  bool QFont::bold();
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
    C.demth_ZNK5QFont4boldEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "bold", args)
  }

}

  // proto:  bool QFont::fixedPitch();
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
    C._ZNK5QFont10fixedPitchEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "fixedPitch", args)
  }

}

  // proto: static QString QFont::substitute(const QString & );
func (this *QFont) substitute_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFont", "substitute", args)
  }

}

  // proto:  void QFont::setFixedPitch(bool );
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
    C._ZN5QFont13setFixedPitchEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setFixedPitch", args)
  }

}

  // proto: static void QFont::removeSubstitutions(const QString & );
func (this *QFont) removeSubstitutions_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFont", "removeSubstitutions", args)
  }

}

  // proto:  void QFont::setPixelSize(int );
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
    C._ZN5QFont12setPixelSizeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setPixelSize", args)
  }

}

  // proto: static void QFont::initialize();
func (this *QFont) initialize_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFont", "initialize", args)
  }

}

  // proto:  QString QFont::key();
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
    C._ZNK5QFont3keyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "key", args)
  }

}

  // proto:  QString QFont::lastResortFont();
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
    C._ZNK5QFont14lastResortFontEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "lastResortFont", args)
  }

}

  // proto:  void QFont::swap(QFont & other);
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
    C.demth_ZN5QFont4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "swap", args)
  }

}

  // proto:  QString QFont::defaultFamily();
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
    C._ZNK5QFont13defaultFamilyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "defaultFamily", args)
  }

}

  // proto:  void QFont::setStrikeOut(bool );
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
    C._ZN5QFont12setStrikeOutEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setStrikeOut", args)
  }

}

  // proto: static void QFont::cleanup();
func (this *QFont) cleanup_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFont", "cleanup", args)
  }

}

  // proto:  bool QFont::exactMatch();
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
    C._ZNK5QFont10exactMatchEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "exactMatch", args)
  }

}

  // proto:  int QFont::stretch();
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
    C._ZNK5QFont7stretchEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFont", "stretch", args)
  }

}

  // proto:  void QFont::setRawName(const QString & );
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
    C._ZN5QFont10setRawNameERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFont", "setRawName", args)
  }

}

// <= body block end

