package qtgui
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QRectF QRawFont::boundingRect(quint32 glyphIndex);
extern void* C_ZNK8QRawFont12boundingRectEj(void* qthis, int32_t arg0); // 4
  // proto:  qreal QRawFont::underlinePosition();
extern double C_ZNK8QRawFont17underlinePositionEv(void* qthis); // 4
  // proto:  qreal QRawFont::ascent();
extern double C_ZNK8QRawFont6ascentEv(void* qthis); // 4
  // proto:  qreal QRawFont::lineThickness();
extern double C_ZNK8QRawFont13lineThicknessEv(void* qthis); // 4
  // proto:  QFont::Style QRawFont::style();
extern void C_ZNK8QRawFont5styleEv(void* qthis); // 4
  // proto:  qreal QRawFont::descent();
extern double C_ZNK8QRawFont7descentEv(void* qthis); // 4
  // proto:  void QRawFont::QRawFont();
extern void* C_ZN8QRawFontC2Ev(); // 3
  // proto:  void QRawFont::QRawFont(const QRawFont & other);
extern void* C_ZN8QRawFontC2ERKS_(void* arg0); // 3
  // proto:  qreal QRawFont::leading();
extern double C_ZNK8QRawFont7leadingEv(void* qthis); // 4
  // proto:  bool QRawFont::advancesForGlyphIndexes(const quint32 * glyphIndexes, QPointF * advances, int numGlyphs);
extern bool C_ZNK8QRawFont23advancesForGlyphIndexesEPKjP7QPointFi(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  QByteArray QRawFont::fontTable(const char * tagName);
extern void* C_ZNK8QRawFont9fontTableEPKc(void* qthis, void* arg0); // 4
  // proto:  void QRawFont::swap(QRawFont & other);
extern void C_ZN8QRawFont4swapERS_(void* qthis, void* arg0); // 2
  // proto:  qreal QRawFont::pixelSize();
extern double C_ZNK8QRawFont9pixelSizeEv(void* qthis); // 4
  // proto:  bool QRawFont::glyphIndexesForChars(const QChar * chars, int numChars, quint32 * glyphIndexes, int * numGlyphs);
extern bool C_ZNK8QRawFont20glyphIndexesForCharsEPK5QChariPjPi(void* qthis, void* arg0, int32_t arg1, void* arg2, void* arg3); // 4
  // proto:  bool QRawFont::supportsCharacter(QChar character);
extern bool C_ZNK8QRawFont17supportsCharacterE5QChar(void* qthis, void* arg0); // 4
  // proto:  bool QRawFont::supportsCharacter(uint ucs4);
extern bool C_ZNK8QRawFont17supportsCharacterEj(void* qthis, int32_t arg0); // 4
  // proto:  QVector<quint32> QRawFont::glyphIndexesForString(const QString & text);
extern void C_ZNK8QRawFont21glyphIndexesForStringERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QRawFont::isValid();
extern bool C_ZNK8QRawFont7isValidEv(void* qthis); // 4
  // proto:  QString QRawFont::familyName();
extern void* C_ZNK8QRawFont10familyNameEv(void* qthis); // 4
  // proto:  int QRawFont::weight();
extern int32_t C_ZNK8QRawFont6weightEv(void* qthis); // 4
  // proto:  QFont::HintingPreference QRawFont::hintingPreference();
extern void C_ZNK8QRawFont17hintingPreferenceEv(void* qthis); // 4
  // proto:  QPainterPath QRawFont::pathForGlyph(quint32 glyphIndex);
extern void* C_ZNK8QRawFont12pathForGlyphEj(void* qthis, int32_t arg0); // 4
  // proto:  void QRawFont::setPixelSize(qreal pixelSize);
extern void C_ZN8QRawFont12setPixelSizeEd(void* qthis, double arg0); // 4
  // proto:  void QRawFont::~QRawFont();
extern void C_ZN8QRawFontD2Ev(void* qthis); // 4
  // proto:  qreal QRawFont::maxCharWidth();
extern double C_ZNK8QRawFont12maxCharWidthEv(void* qthis); // 4
  // proto:  qreal QRawFont::averageCharWidth();
extern double C_ZNK8QRawFont16averageCharWidthEv(void* qthis); // 4
  // proto:  qreal QRawFont::unitsPerEm();
extern double C_ZNK8QRawFont10unitsPerEmEv(void* qthis); // 4
  // proto:  qreal QRawFont::xHeight();
extern double C_ZNK8QRawFont7xHeightEv(void* qthis); // 4
  // proto:  QString QRawFont::styleName();
extern void* C_ZNK8QRawFont9styleNameEv(void* qthis); // 4
  // proto:  QList<QFontDatabase::WritingSystem> QRawFont::supportedWritingSystems();
extern void C_ZNK8QRawFont23supportedWritingSystemsEv(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QRawFont)=1
type QRawFont struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// boundingRect(quint32)
func (this *QRawFont) Boundingrect(args ...interface{}) (ret interface{}) {
  // boundingRect(quint32)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "quint32"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont12boundingRectEj
    // invoke: QRectF boundingRect(quint32)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK8QRawFont12boundingRectEj(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRawFont", "boundingRect", args)
  }

  return
}

// underlinePosition()
func (this *QRawFont) Underlineposition(args ...interface{}) (ret interface{}) {
  // underlinePosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont17underlinePositionEv
    // invoke: qreal underlinePosition()
    var ret0 = C.C_ZNK8QRawFont17underlinePositionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRawFont", "underlinePosition", args)
  }

  return
}

// ascent()
func (this *QRawFont) Ascent(args ...interface{}) (ret interface{}) {
  // ascent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont6ascentEv
    // invoke: qreal ascent()
    var ret0 = C.C_ZNK8QRawFont6ascentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRawFont", "ascent", args)
  }

  return
}

// lineThickness()
func (this *QRawFont) Linethickness(args ...interface{}) (ret interface{}) {
  // lineThickness()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont13lineThicknessEv
    // invoke: qreal lineThickness()
    var ret0 = C.C_ZNK8QRawFont13lineThicknessEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRawFont", "lineThickness", args)
  }

  return
}

// style()
func (this *QRawFont) Style(args ...interface{}) () {
  // style()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont5styleEv
    // invoke: QFont::Style style()
    C.C_ZNK8QRawFont5styleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "style", args)
  }

  return
}

// descent()
func (this *QRawFont) Descent(args ...interface{}) (ret interface{}) {
  // descent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont7descentEv
    // invoke: qreal descent()
    var ret0 = C.C_ZNK8QRawFont7descentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRawFont", "descent", args)
  }

  return
}

// QRawFont()
func NewQRawFont(args ...interface{}) *QRawFont {
  // QRawFont()
  // QRawFont(const class QRawFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRawFont{}) // "const QRawFont &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QRawFontC1Ev
    // invoke: void QRawFont()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QRawFontC2Ev()
    return &QRawFont{Qclsinst:qthis}
  case 1:
    // invoke: _ZN8QRawFontC1ERKS_
    // invoke: void QRawFont(const class QRawFont &)
    var arg0 = args[0].(*QRawFont).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QRawFontC2ERKS_(arg0)
    return &QRawFont{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QRawFont", "QRawFont", args)
  }

  return nil // QRawFont{Qclsinst:qthis}
}

// leading()
func (this *QRawFont) Leading(args ...interface{}) (ret interface{}) {
  // leading()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont7leadingEv
    // invoke: qreal leading()
    var ret0 = C.C_ZNK8QRawFont7leadingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRawFont", "leading", args)
  }

  return
}

// advancesForGlyphIndexes(const quint32 *, class QPointF *, int)
func (this *QRawFont) Advancesforglyphindexes(args ...interface{}) (ret interface{}) {
  // advancesForGlyphIndexes(const quint32 *, class QPointF *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "const quint32 *"
  vtys[0][1] = reflect.TypeOf(qtcore.QPointF{}) // "QPointF *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont23advancesForGlyphIndexesEPKjP7QPointFi
    // invoke: bool advancesForGlyphIndexes(const quint32 *, class QPointF *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK8QRawFont23advancesForGlyphIndexesEPKjP7QPointFi(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRawFont", "advancesForGlyphIndexes", args)
  }

  return
}

// fontTable(const char *)
func (this *QRawFont) Fonttable(args ...interface{}) (ret interface{}) {
  // fontTable(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont9fontTableEPKc
    // invoke: QByteArray fontTable(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZNK8QRawFont9fontTableEPKc(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRawFont", "fontTable", args)
  }

  return
}

// swap(class QRawFont &)
func (this *QRawFont) Swap(args ...interface{}) () {
  // swap(class QRawFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRawFont{}) // "QRawFont &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QRawFont4swapERS_
    // invoke: void swap(class QRawFont &)
    var arg0 = args[0].(*QRawFont).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QRawFont4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRawFont", "swap", args)
  }

  return
}

// pixelSize()
func (this *QRawFont) Pixelsize(args ...interface{}) (ret interface{}) {
  // pixelSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont9pixelSizeEv
    // invoke: qreal pixelSize()
    var ret0 = C.C_ZNK8QRawFont9pixelSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRawFont", "pixelSize", args)
  }

  return
}

// glyphIndexesForChars(const class QChar *, int, quint32 *, int *)
func (this *QRawFont) Glyphindexesforchars(args ...interface{}) (ret interface{}) {
  // glyphIndexesForChars(const class QChar *, int, quint32 *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QChar{}) // "const QChar *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(true) // "quint32 *"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont20glyphIndexesForCharsEPK5QChariPjPi
    // invoke: bool glyphIndexesForChars(const class QChar *, int, quint32 *, int *)
    var arg0 = args[0].(*qtcore.QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK8QRawFont20glyphIndexesForCharsEPK5QChariPjPi(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRawFont", "glyphIndexesForChars", args)
  }

  return
}

// supportsCharacter(class QChar)
func (this *QRawFont) Supportscharacter(args ...interface{}) (ret interface{}) {
  // supportsCharacter(class QChar)
  // supportsCharacter(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QChar{}) // "QChar"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont17supportsCharacterE5QChar
    // invoke: bool supportsCharacter(class QChar)
    var arg0 = args[0].(*qtcore.QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK8QRawFont17supportsCharacterE5QChar(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK8QRawFont17supportsCharacterEj
    // invoke: bool supportsCharacter(uint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK8QRawFont17supportsCharacterEj(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRawFont", "supportsCharacter", args)
  }

  return
}

// glyphIndexesForString(const class QString &)
func (this *QRawFont) Glyphindexesforstring(args ...interface{}) () {
  // glyphIndexesForString(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont21glyphIndexesForStringERK7QString
    // invoke: QVector<quint32> glyphIndexesForString(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK8QRawFont21glyphIndexesForStringERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRawFont", "glyphIndexesForString", args)
  }

  return
}

// isValid()
func (this *QRawFont) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK8QRawFont7isValidEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRawFont", "isValid", args)
  }

  return
}

// familyName()
func (this *QRawFont) Familyname(args ...interface{}) (ret interface{}) {
  // familyName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont10familyNameEv
    // invoke: QString familyName()
    var ret0 = C.C_ZNK8QRawFont10familyNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRawFont", "familyName", args)
  }

  return
}

// weight()
func (this *QRawFont) Weight(args ...interface{}) (ret interface{}) {
  // weight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont6weightEv
    // invoke: int weight()
    var ret0 = C.C_ZNK8QRawFont6weightEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRawFont", "weight", args)
  }

  return
}

// hintingPreference()
func (this *QRawFont) Hintingpreference(args ...interface{}) () {
  // hintingPreference()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont17hintingPreferenceEv
    // invoke: QFont::HintingPreference hintingPreference()
    C.C_ZNK8QRawFont17hintingPreferenceEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "hintingPreference", args)
  }

  return
}

// pathForGlyph(quint32)
func (this *QRawFont) Pathforglyph(args ...interface{}) (ret interface{}) {
  // pathForGlyph(quint32)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "quint32"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont12pathForGlyphEj
    // invoke: QPainterPath pathForGlyph(quint32)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK8QRawFont12pathForGlyphEj(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRawFont", "pathForGlyph", args)
  }

  return
}

// setPixelSize(qreal)
func (this *QRawFont) Setpixelsize(args ...interface{}) () {
  // setPixelSize(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QRawFont12setPixelSizeEd
    // invoke: void setPixelSize(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN8QRawFont12setPixelSizeEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRawFont", "setPixelSize", args)
  }

  return
}

// ~QRawFont()
func (this *QRawFont) Freeqrawfont(args ...interface{}) () {
  // ~QRawFont()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QRawFontD0Ev
    // invoke: void ~QRawFont()
    C.C_ZN8QRawFontD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "~QRawFont", args)
  }

  return
}

// maxCharWidth()
func (this *QRawFont) Maxcharwidth(args ...interface{}) (ret interface{}) {
  // maxCharWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont12maxCharWidthEv
    // invoke: qreal maxCharWidth()
    var ret0 = C.C_ZNK8QRawFont12maxCharWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRawFont", "maxCharWidth", args)
  }

  return
}

// averageCharWidth()
func (this *QRawFont) Averagecharwidth(args ...interface{}) (ret interface{}) {
  // averageCharWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont16averageCharWidthEv
    // invoke: qreal averageCharWidth()
    var ret0 = C.C_ZNK8QRawFont16averageCharWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRawFont", "averageCharWidth", args)
  }

  return
}

// unitsPerEm()
func (this *QRawFont) Unitsperem(args ...interface{}) (ret interface{}) {
  // unitsPerEm()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont10unitsPerEmEv
    // invoke: qreal unitsPerEm()
    var ret0 = C.C_ZNK8QRawFont10unitsPerEmEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRawFont", "unitsPerEm", args)
  }

  return
}

// xHeight()
func (this *QRawFont) Xheight(args ...interface{}) (ret interface{}) {
  // xHeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont7xHeightEv
    // invoke: qreal xHeight()
    var ret0 = C.C_ZNK8QRawFont7xHeightEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRawFont", "xHeight", args)
  }

  return
}

// styleName()
func (this *QRawFont) Stylename(args ...interface{}) (ret interface{}) {
  // styleName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont9styleNameEv
    // invoke: QString styleName()
    var ret0 = C.C_ZNK8QRawFont9styleNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRawFont", "styleName", args)
  }

  return
}

// supportedWritingSystems()
func (this *QRawFont) Supportedwritingsystems(args ...interface{}) () {
  // supportedWritingSystems()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QRawFont23supportedWritingSystemsEv
    // invoke: QList<QFontDatabase::WritingSystem> supportedWritingSystems()
    C.C_ZNK8QRawFont23supportedWritingSystemsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QRawFont", "supportedWritingSystems", args)
  }

  return
}

// <= body block end

