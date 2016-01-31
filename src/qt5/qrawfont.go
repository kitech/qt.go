package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtGui/qrawfont.h
// dst-file: /src/gui/qrawfont.go
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
  // proto:  QRectF QRawFont::boundingRect(quint32 glyphIndex);
extern void C_ZNK8QRawFont12boundingRectEj(void* qthis, int32_t arg0); // 4
  // proto:  qreal QRawFont::underlinePosition();
extern void C_ZNK8QRawFont17underlinePositionEv(void* qthis); // 4
  // proto:  qreal QRawFont::ascent();
extern void C_ZNK8QRawFont6ascentEv(void* qthis); // 4
  // proto:  qreal QRawFont::lineThickness();
extern void C_ZNK8QRawFont13lineThicknessEv(void* qthis); // 4
  // proto:  QFont::Style QRawFont::style();
extern void C_ZNK8QRawFont5styleEv(void* qthis); // 4
  // proto:  qreal QRawFont::descent();
extern void C_ZNK8QRawFont7descentEv(void* qthis); // 4
  // proto:  void QRawFont::QRawFont();
extern void C_ZN8QRawFontC2Ev(void* qthis); // 3
  // proto:  void QRawFont::QRawFont(const QRawFont & other);
extern void C_ZN8QRawFontC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  qreal QRawFont::leading();
extern void C_ZNK8QRawFont7leadingEv(void* qthis); // 4
  // proto:  bool QRawFont::advancesForGlyphIndexes(const quint32 * glyphIndexes, QPointF * advances, int numGlyphs);
extern void C_ZNK8QRawFont23advancesForGlyphIndexesEPKjP7QPointFi(void* qthis, int32_t* arg0, void* arg1, int32_t arg2); // 4
  // proto:  QByteArray QRawFont::fontTable(const char * tagName);
extern void C_ZNK8QRawFont9fontTableEPKc(void* qthis, unsigned char* arg0); // 4
  // proto:  void QRawFont::swap(QRawFont & other);
extern void C_ZN8QRawFont4swapERS_(void* qthis, void* arg0); // 2
  // proto:  qreal QRawFont::pixelSize();
extern void C_ZNK8QRawFont9pixelSizeEv(void* qthis); // 4
  // proto:  bool QRawFont::glyphIndexesForChars(const QChar * chars, int numChars, quint32 * glyphIndexes, int * numGlyphs);
extern void C_ZNK8QRawFont20glyphIndexesForCharsEPK5QChariPjPi(void* qthis, void* arg0, int32_t arg1, int32_t* arg2, int32_t* arg3); // 4
  // proto:  bool QRawFont::supportsCharacter(QChar character);
extern void C_ZNK8QRawFont17supportsCharacterE5QChar(void* qthis, void* arg0); // 4
  // proto:  bool QRawFont::supportsCharacter(uint ucs4);
extern void C_ZNK8QRawFont17supportsCharacterEj(void* qthis, int32_t arg0); // 4
  // proto:  QVector<quint32> QRawFont::glyphIndexesForString(const QString & text);
extern void C_ZNK8QRawFont21glyphIndexesForStringERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QRawFont::isValid();
extern void C_ZNK8QRawFont7isValidEv(void* qthis); // 4
  // proto:  QString QRawFont::familyName();
extern void C_ZNK8QRawFont10familyNameEv(void* qthis); // 4
  // proto:  int QRawFont::weight();
extern void C_ZNK8QRawFont6weightEv(void* qthis); // 4
  // proto:  QFont::HintingPreference QRawFont::hintingPreference();
extern void C_ZNK8QRawFont17hintingPreferenceEv(void* qthis); // 4
  // proto:  QPainterPath QRawFont::pathForGlyph(quint32 glyphIndex);
extern void C_ZNK8QRawFont12pathForGlyphEj(void* qthis, int32_t arg0); // 4
  // proto:  void QRawFont::setPixelSize(qreal pixelSize);
extern void C_ZN8QRawFont12setPixelSizeEd(void* qthis, double arg0); // 4
  // proto:  void QRawFont::~QRawFont();
extern void C_ZN8QRawFontD2Ev(void* qthis); // 4
  // proto:  qreal QRawFont::maxCharWidth();
extern void C_ZNK8QRawFont12maxCharWidthEv(void* qthis); // 4
  // proto:  qreal QRawFont::averageCharWidth();
extern void C_ZNK8QRawFont16averageCharWidthEv(void* qthis); // 4
  // proto:  qreal QRawFont::unitsPerEm();
extern void C_ZNK8QRawFont10unitsPerEmEv(void* qthis); // 4
  // proto:  qreal QRawFont::xHeight();
extern void C_ZNK8QRawFont7xHeightEv(void* qthis); // 4
  // proto:  QString QRawFont::styleName();
extern void C_ZNK8QRawFont9styleNameEv(void* qthis); // 4
  // proto:  QList<QFontDatabase::WritingSystem> QRawFont::supportedWritingSystems();
extern void C_ZNK8QRawFont23supportedWritingSystemsEv(void* qthis); // 4
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

// class sizeof(QRawFont)=1
type QRawFont struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// boundingRect(quint32)
func (this *QRawFont) boundingRect(args ...interface{}) () {
  // boundingRect(quint32)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "quint32"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont12boundingRectEj
    // invoke: QRectF boundingRect(quint32)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK8QRawFont12boundingRectEj(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRawFont", "boundingRect", args)
  }

}

// underlinePosition()
func (this *QRawFont) underlinePosition(args ...interface{}) () {
  // underlinePosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont17underlinePositionEv
    // invoke: qreal underlinePosition()
    C.C_ZNK8QRawFont17underlinePositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "underlinePosition", args)
  }

}

// ascent()
func (this *QRawFont) ascent(args ...interface{}) () {
  // ascent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont6ascentEv
    // invoke: qreal ascent()
    C.C_ZNK8QRawFont6ascentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "ascent", args)
  }

}

// lineThickness()
func (this *QRawFont) lineThickness(args ...interface{}) () {
  // lineThickness()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont13lineThicknessEv
    // invoke: qreal lineThickness()
    C.C_ZNK8QRawFont13lineThicknessEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "lineThickness", args)
  }

}

// style()
func (this *QRawFont) style(args ...interface{}) () {
  // style()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont5styleEv
    // invoke: QFont::Style style()
    C.C_ZNK8QRawFont5styleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "style", args)
  }

}

// descent()
func (this *QRawFont) descent(args ...interface{}) () {
  // descent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont7descentEv
    // invoke: qreal descent()
    C.C_ZNK8QRawFont7descentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "descent", args)
  }

}

// QRawFont()
func NewQRawFont(args ...interface{}) QRawFont {
  // QRawFont()
  // QRawFont(const class QRawFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRawFont{}) // "const QRawFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QRawFontC1Ev
    // invoke: void QRawFont()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN8QRawFontC2Ev(qthis)
  case 1:
    // invoke: _ZN8QRawFontC1ERKS_
    // invoke: void QRawFont(const class QRawFont &)
    var arg0 = args[0].(QRawFont).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN8QRawFontC2ERKS_(qthis, arg0)
  default:
    qtrt.ErrorResolve("QRawFont", "QRawFont", args)
  }

  return QRawFont{}
}

// leading()
func (this *QRawFont) leading(args ...interface{}) () {
  // leading()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont7leadingEv
    // invoke: qreal leading()
    C.C_ZNK8QRawFont7leadingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "leading", args)
  }

}

// advancesForGlyphIndexes(const quint32 *, class QPointF *, int)
func (this *QRawFont) advancesForGlyphIndexes(args ...interface{}) () {
  // advancesForGlyphIndexes(const quint32 *, class QPointF *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "const quint32 *"
  vtys[0][1] = reflect.TypeOf(QPointF{}) // "QPointF *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont23advancesForGlyphIndexesEPKjP7QPointFi
    // invoke: bool advancesForGlyphIndexes(const quint32 *, class QPointF *, int)
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPointF).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C.C_ZNK8QRawFont23advancesForGlyphIndexesEPKjP7QPointFi(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QRawFont", "advancesForGlyphIndexes", args)
  }

}

// fontTable(const char *)
func (this *QRawFont) fontTable(args ...interface{}) () {
  // fontTable(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont9fontTableEPKc
    // invoke: QByteArray fontTable(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C.C_ZNK8QRawFont9fontTableEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRawFont", "fontTable", args)
  }

}

// swap(class QRawFont &)
func (this *QRawFont) swap(args ...interface{}) () {
  // swap(class QRawFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRawFont{}) // "QRawFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QRawFont4swapERS_
    // invoke: void swap(class QRawFont &)
    var arg0 = args[0].(QRawFont).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QRawFont4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRawFont", "swap", args)
  }

}

// pixelSize()
func (this *QRawFont) pixelSize(args ...interface{}) () {
  // pixelSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont9pixelSizeEv
    // invoke: qreal pixelSize()
    C.C_ZNK8QRawFont9pixelSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "pixelSize", args)
  }

}

// glyphIndexesForChars(const class QChar *, int, quint32 *, int *)
func (this *QRawFont) glyphIndexesForChars(args ...interface{}) () {
  // glyphIndexesForChars(const class QChar *, int, quint32 *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(true) // "quint32 *"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont20glyphIndexesForCharsEPK5QChariPjPi
    // invoke: bool glyphIndexesForChars(const class QChar *, int, quint32 *, int *)
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int32_t)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    C.C_ZNK8QRawFont20glyphIndexesForCharsEPK5QChariPjPi(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QRawFont", "glyphIndexesForChars", args)
  }

}

// supportsCharacter(class QChar)
func (this *QRawFont) supportsCharacter(args ...interface{}) () {
  // supportsCharacter(class QChar)
  // supportsCharacter(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont17supportsCharacterE5QChar
    // invoke: bool supportsCharacter(class QChar)
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK8QRawFont17supportsCharacterE5QChar(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK8QRawFont17supportsCharacterEj
    // invoke: bool supportsCharacter(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK8QRawFont17supportsCharacterEj(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRawFont", "supportsCharacter", args)
  }

}

// glyphIndexesForString(const class QString &)
func (this *QRawFont) glyphIndexesForString(args ...interface{}) () {
  // glyphIndexesForString(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont21glyphIndexesForStringERK7QString
    // invoke: QVector<quint32> glyphIndexesForString(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK8QRawFont21glyphIndexesForStringERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRawFont", "glyphIndexesForString", args)
  }

}

// isValid()
func (this *QRawFont) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont7isValidEv
    // invoke: bool isValid()
    C.C_ZNK8QRawFont7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "isValid", args)
  }

}

// familyName()
func (this *QRawFont) familyName(args ...interface{}) () {
  // familyName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont10familyNameEv
    // invoke: QString familyName()
    C.C_ZNK8QRawFont10familyNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "familyName", args)
  }

}

// weight()
func (this *QRawFont) weight(args ...interface{}) () {
  // weight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont6weightEv
    // invoke: int weight()
    C.C_ZNK8QRawFont6weightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "weight", args)
  }

}

// hintingPreference()
func (this *QRawFont) hintingPreference(args ...interface{}) () {
  // hintingPreference()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont17hintingPreferenceEv
    // invoke: QFont::HintingPreference hintingPreference()
    C.C_ZNK8QRawFont17hintingPreferenceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "hintingPreference", args)
  }

}

// pathForGlyph(quint32)
func (this *QRawFont) pathForGlyph(args ...interface{}) () {
  // pathForGlyph(quint32)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "quint32"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont12pathForGlyphEj
    // invoke: QPainterPath pathForGlyph(quint32)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK8QRawFont12pathForGlyphEj(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRawFont", "pathForGlyph", args)
  }

}

// setPixelSize(qreal)
func (this *QRawFont) setPixelSize(args ...interface{}) () {
  // setPixelSize(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QRawFont12setPixelSizeEd
    // invoke: void setPixelSize(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN8QRawFont12setPixelSizeEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRawFont", "setPixelSize", args)
  }

}

// ~QRawFont()
func (this *QRawFont) FreeQRawFont(args ...interface{}) () {
  // ~QRawFont()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QRawFontD0Ev
    // invoke: void ~QRawFont()
    C.C_ZN8QRawFontD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "~QRawFont", args)
  }

}

// maxCharWidth()
func (this *QRawFont) maxCharWidth(args ...interface{}) () {
  // maxCharWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont12maxCharWidthEv
    // invoke: qreal maxCharWidth()
    C.C_ZNK8QRawFont12maxCharWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "maxCharWidth", args)
  }

}

// averageCharWidth()
func (this *QRawFont) averageCharWidth(args ...interface{}) () {
  // averageCharWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont16averageCharWidthEv
    // invoke: qreal averageCharWidth()
    C.C_ZNK8QRawFont16averageCharWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "averageCharWidth", args)
  }

}

// unitsPerEm()
func (this *QRawFont) unitsPerEm(args ...interface{}) () {
  // unitsPerEm()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont10unitsPerEmEv
    // invoke: qreal unitsPerEm()
    C.C_ZNK8QRawFont10unitsPerEmEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "unitsPerEm", args)
  }

}

// xHeight()
func (this *QRawFont) xHeight(args ...interface{}) () {
  // xHeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont7xHeightEv
    // invoke: qreal xHeight()
    C.C_ZNK8QRawFont7xHeightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "xHeight", args)
  }

}

// styleName()
func (this *QRawFont) styleName(args ...interface{}) () {
  // styleName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont9styleNameEv
    // invoke: QString styleName()
    C.C_ZNK8QRawFont9styleNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "styleName", args)
  }

}

// supportedWritingSystems()
func (this *QRawFont) supportedWritingSystems(args ...interface{}) () {
  // supportedWritingSystems()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont23supportedWritingSystemsEv
    // invoke: QList<QFontDatabase::WritingSystem> supportedWritingSystems()
    C.C_ZNK8QRawFont23supportedWritingSystemsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "supportedWritingSystems", args)
  }

}

// <= body block end

