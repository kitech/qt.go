package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qrawfont.h
// dst-file: /src/gui/qrawfont.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
// class sizeof(QRawFont)=1
type QRawFont struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


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
  default:
    qtrt.ErrorResolve("QRawFont", "averageCharWidth", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QRawFont", "ascent", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QRawFont", "leading", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QRawFont", "lineThickness", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QRawFont", "isValid", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QRawFont", "boundingRect", args)
  }

}


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
  case 1:
    // invoke: _ZNK8QRawFont17supportsCharacterE5QChar
  default:
    qtrt.ErrorResolve("QRawFont", "supportsCharacter", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QRawFont", "swap", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QRawFont", "descent", args)
  }

}


func NewQRawFont(args ...interface{}) QRawFont {
  return QRawFont{}
}


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
  default:
    qtrt.ErrorResolve("QRawFont", "setPixelSize", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QRawFont", "glyphIndexesForChars", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QRawFont", "styleName", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QRawFont", "underlinePosition", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QRawFont", "unitsPerEm", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QRawFont", "familyName", args)
  }

}


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
    // invoke: _ZNK8QRawFont23advancesForGlyphIndexesERK7QVectorIjE6QFlagsINS_10LayoutFlagEE
  case 1:
    // invoke: _ZNK8QRawFont23advancesForGlyphIndexesEPKjP7QPointFi
  case 2:
    // invoke: _ZNK8QRawFont23advancesForGlyphIndexesERK7QVectorIjE
  case 3:
    // invoke: _ZNK8QRawFont23advancesForGlyphIndexesEPKjP7QPointFi6QFlagsINS_10LayoutFlagEE
  default:
    qtrt.ErrorResolve("QRawFont", "advancesForGlyphIndexes", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QRawFont", "pixelSize", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QRawFont", "weight", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QRawFont", "xHeight", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QRawFont", "pathForGlyph", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QRawFont", "fontTable", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QRawFont", "maxCharWidth", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QRawFont", "glyphIndexesForString", args)
  }

}

// <= body block end

