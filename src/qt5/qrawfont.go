package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
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
  // proto:  qreal QRawFont::averageCharWidth();
extern void _ZNK8QRawFont16averageCharWidthEv(void* qthis);
  // proto:  qreal QRawFont::ascent();
extern void _ZNK8QRawFont6ascentEv(void* qthis);
  // proto:  qreal QRawFont::leading();
extern void _ZNK8QRawFont7leadingEv(void* qthis);
  // proto:  qreal QRawFont::lineThickness();
extern void _ZNK8QRawFont13lineThicknessEv(void* qthis);
  // proto:  bool QRawFont::isValid();
extern void _ZNK8QRawFont7isValidEv(void* qthis);
  // proto:  QRectF QRawFont::boundingRect(quint32 glyphIndex);
extern void _ZNK8QRawFont12boundingRectEj(void* qthis, unsigned int arg0);
  // proto:  bool QRawFont::supportsCharacter(uint ucs4);
extern void _ZNK8QRawFont17supportsCharacterEj(void* qthis, unsigned int arg0);
  // proto:  void QRawFont::swap(QRawFont & other);
extern void _ZN8QRawFont4swapERS_(void* qthis, void* arg0);
  // proto:  qreal QRawFont::descent();
extern void _ZNK8QRawFont7descentEv(void* qthis);
  // proto:  void QRawFont::QRawFont();
extern void* dector_ZN8QRawFontC1Ev();
extern void _ZN8QRawFontC1Ev(void* qthis);
  // proto:  void QRawFont::setPixelSize(qreal pixelSize);
extern void _ZN8QRawFont12setPixelSizeEd(void* qthis, double arg0);
  // proto:  bool QRawFont::glyphIndexesForChars(const QChar * chars, int numChars, quint32 * glyphIndexes, int * numGlyphs);
extern void _ZNK8QRawFont20glyphIndexesForCharsEPK5QChariPjPi(void* qthis, void* arg0, int arg1, unsigned int* arg2, int* arg3);
  // proto:  QString QRawFont::styleName();
extern void _ZNK8QRawFont9styleNameEv(void* qthis);
  // proto:  qreal QRawFont::underlinePosition();
extern void _ZNK8QRawFont17underlinePositionEv(void* qthis);
  // proto:  qreal QRawFont::unitsPerEm();
extern void _ZNK8QRawFont10unitsPerEmEv(void* qthis);
  // proto:  bool QRawFont::supportsCharacter(QChar character);
extern void _ZNK8QRawFont17supportsCharacterE5QChar(void* qthis, void* arg0);
  // proto:  QString QRawFont::familyName();
extern void _ZNK8QRawFont10familyNameEv(void* qthis);
  // proto:  bool QRawFont::advancesForGlyphIndexes(const quint32 * glyphIndexes, QPointF * advances, int numGlyphs);
extern void _ZNK8QRawFont23advancesForGlyphIndexesEPKjP7QPointFi(void* qthis, unsigned int* arg0, void* arg1, int arg2);
  // proto:  qreal QRawFont::pixelSize();
extern void _ZNK8QRawFont9pixelSizeEv(void* qthis);
  // proto:  int QRawFont::weight();
extern void _ZNK8QRawFont6weightEv(void* qthis);
  // proto:  void QRawFont::QRawFont(const QRawFont & other);
extern void* dector_ZN8QRawFontC1ERKS_(void* arg0);
extern void _ZN8QRawFontC1ERKS_(void* qthis, void* arg0);
  // proto:  qreal QRawFont::xHeight();
extern void _ZNK8QRawFont7xHeightEv(void* qthis);
  // proto:  void QRawFont::~QRawFont();
extern void _ZN8QRawFontD0Ev(void* qthis);
  // proto:  QPainterPath QRawFont::pathForGlyph(quint32 glyphIndex);
extern void _ZNK8QRawFont12pathForGlyphEj(void* qthis, unsigned int arg0);
  // proto:  QByteArray QRawFont::fontTable(const char * tagName);
extern void _ZNK8QRawFont9fontTableEPKc(void* qthis, char* arg0);
  // proto:  qreal QRawFont::maxCharWidth();
extern void _ZNK8QRawFont12maxCharWidthEv(void* qthis);
  // proto:  QVector<quint32> QRawFont::glyphIndexesForString(const QString & text);
extern void _ZNK8QRawFont21glyphIndexesForStringERK7QString(void* qthis, void* arg0);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  qreal QRawFont::averageCharWidth();
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
    C._ZNK8QRawFont16averageCharWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "averageCharWidth", args)
  }

}

  // proto:  qreal QRawFont::ascent();
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
    C._ZNK8QRawFont6ascentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "ascent", args)
  }

}

  // proto:  qreal QRawFont::leading();
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
    C._ZNK8QRawFont7leadingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "leading", args)
  }

}

  // proto:  qreal QRawFont::lineThickness();
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
    C._ZNK8QRawFont13lineThicknessEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "lineThickness", args)
  }

}

  // proto:  bool QRawFont::isValid();
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
    C._ZNK8QRawFont7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "isValid", args)
  }

}

  // proto:  QRectF QRawFont::boundingRect(quint32 glyphIndex);
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
    C._ZNK8QRawFont12boundingRectEj(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRawFont", "boundingRect", args)
  }

}

  // proto:  bool QRawFont::supportsCharacter(uint ucs4);
func (this *QRawFont) supportsCharacter(args ...interface{}) () {
  // supportsCharacter(uint)
  // supportsCharacter(class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont17supportsCharacterEj
    // invoke: bool supportsCharacter(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK8QRawFont17supportsCharacterEj(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK8QRawFont17supportsCharacterE5QChar
    // invoke: bool supportsCharacter(class QChar)
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK8QRawFont17supportsCharacterE5QChar(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRawFont", "supportsCharacter", args)
  }

}

  // proto:  void QRawFont::swap(QRawFont & other);
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
    C._ZN8QRawFont4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRawFont", "swap", args)
  }

}

  // proto:  qreal QRawFont::descent();
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
    C._ZNK8QRawFont7descentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "descent", args)
  }

}

  // proto:  void QRawFont::QRawFont();
func NewQRawFont(args ...interface{}) QRawFont {
  return QRawFont{}
}

  // proto:  void QRawFont::setPixelSize(qreal pixelSize);
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
    C._ZN8QRawFont12setPixelSizeEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRawFont", "setPixelSize", args)
  }

}

  // proto:  bool QRawFont::glyphIndexesForChars(const QChar * chars, int numChars, quint32 * glyphIndexes, int * numGlyphs);
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
    C._ZNK8QRawFont20glyphIndexesForCharsEPK5QChariPjPi(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QRawFont", "glyphIndexesForChars", args)
  }

}

  // proto:  QString QRawFont::styleName();
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
    C._ZNK8QRawFont9styleNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "styleName", args)
  }

}

  // proto:  qreal QRawFont::underlinePosition();
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
    C._ZNK8QRawFont17underlinePositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "underlinePosition", args)
  }

}

  // proto:  qreal QRawFont::unitsPerEm();
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
    C._ZNK8QRawFont10unitsPerEmEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "unitsPerEm", args)
  }

}

  // proto:  QString QRawFont::familyName();
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
    C._ZNK8QRawFont10familyNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "familyName", args)
  }

}

  // proto:  bool QRawFont::advancesForGlyphIndexes(const quint32 * glyphIndexes, QPointF * advances, int numGlyphs);
func (this *QRawFont) advancesForGlyphIndexes(args ...interface{}) () {
  // advancesForGlyphIndexes(const QVector<quint32> &, LayoutFlags)
  // advancesForGlyphIndexes(const quint32 *, class QPointF *, int)
  // advancesForGlyphIndexes(const QVector<quint32> &)
  // advancesForGlyphIndexes(const quint32 *, class QPointF *, int, LayoutFlags)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  // vtys[0][0] = reflect.TypeOf(QVector<unsigned int>{}) // "const QVector<quint32> &"
  vtys[0][1] = qtrt.Int64Ty(false) // "LayoutFlags"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(true) // "const quint32 *"
  vtys[1][1] = reflect.TypeOf(QPointF{}) // "QPointF *"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  // vtys[2][0] = reflect.TypeOf(QVector<unsigned int>{}) // "const QVector<quint32> &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(true) // "const quint32 *"
  vtys[3][1] = reflect.TypeOf(QPointF{}) // "QPointF *"
  vtys[3][2] = qtrt.Int32Ty(false) // "int"
  vtys[3][3] = qtrt.Int64Ty(false) // "LayoutFlags"

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
    C._ZNK8QRawFont23advancesForGlyphIndexesEPKjP7QPointFi(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QRawFont", "advancesForGlyphIndexes", args)
  }

}

  // proto:  qreal QRawFont::pixelSize();
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
    C._ZNK8QRawFont9pixelSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "pixelSize", args)
  }

}

  // proto:  int QRawFont::weight();
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
    C._ZNK8QRawFont6weightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "weight", args)
  }

}

  // proto:  qreal QRawFont::xHeight();
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
    C._ZNK8QRawFont7xHeightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "xHeight", args)
  }

}

  // proto:  void QRawFont::~QRawFont();
func (this *QRawFont) FreeQRawFont(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QRawFont", "~QRawFont", args)
  }

}

  // proto:  QPainterPath QRawFont::pathForGlyph(quint32 glyphIndex);
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
    C._ZNK8QRawFont12pathForGlyphEj(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRawFont", "pathForGlyph", args)
  }

}

  // proto:  QByteArray QRawFont::fontTable(const char * tagName);
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
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    C._ZNK8QRawFont9fontTableEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRawFont", "fontTable", args)
  }

}

  // proto:  qreal QRawFont::maxCharWidth();
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
    C._ZNK8QRawFont12maxCharWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "maxCharWidth", args)
  }

}

  // proto:  QVector<quint32> QRawFont::glyphIndexesForString(const QString & text);
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
    C._ZNK8QRawFont21glyphIndexesForStringERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRawFont", "glyphIndexesForString", args)
  }

}

// <= body block end

