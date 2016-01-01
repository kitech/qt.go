package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qfontmetrics.h
// dst-file: /src/gui/qfontmetrics.go
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
// class sizeof(QFontMetrics)=1
type QFontMetrics struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QFontMetricsF)=1
type QFontMetricsF struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QFontMetrics) maxWidth(args ...interface{}) () {
  // maxWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics8maxWidthEv
  default:
    qtrt.ErrorResolve("QFontMetrics", "maxWidth", args)
 }

}


func (this *QFontMetrics) FreeQFontMetrics(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QFontMetrics", "~QFontMetrics", args)
 }

}


func (this *QFontMetrics) lineWidth(args ...interface{}) () {
  // lineWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics9lineWidthEv
  default:
    qtrt.ErrorResolve("QFontMetrics", "lineWidth", args)
 }

}


func (this *QFontMetrics) boundingRect(args ...interface{}) () {
  // boundingRect(const class QRect &, int, const class QString &, int, int *)
  // boundingRect(const class QString &)
  // boundingRect(class QChar)
  // boundingRect(int, int, int, int, int, const class QString &, int, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(true) // "int *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = qtrt.Int32Ty(false) // "int"
  vtys[3][3] = qtrt.Int32Ty(false) // "int"
  vtys[3][4] = qtrt.Int32Ty(false) // "int"
  vtys[3][5] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3][6] = qtrt.Int32Ty(false) // "int"
  vtys[3][7] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics12boundingRectERK5QRectiRK7QStringiPi
  case 1:
    // invoke: _ZNK12QFontMetrics12boundingRectERK7QString
  case 2:
    // invoke: _ZNK12QFontMetrics12boundingRectE5QChar
  case 3:
    // invoke: _ZNK12QFontMetrics12boundingRectEiiiiiRK7QStringiPi
  default:
    qtrt.ErrorResolve("QFontMetrics", "boundingRect", args)
 }

}


func (this *QFontMetrics) minLeftBearing(args ...interface{}) () {
  // minLeftBearing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics14minLeftBearingEv
  default:
    qtrt.ErrorResolve("QFontMetrics", "minLeftBearing", args)
 }

}


func (this *QFontMetrics) rightBearing(args ...interface{}) () {
  // rightBearing(class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics12rightBearingE5QChar
  default:
    qtrt.ErrorResolve("QFontMetrics", "rightBearing", args)
 }

}


func (this *QFontMetrics) ascent(args ...interface{}) () {
  // ascent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics6ascentEv
  default:
    qtrt.ErrorResolve("QFontMetrics", "ascent", args)
 }

}


func (this *QFontMetrics) size(args ...interface{}) () {
  // size(int, const class QString &, int, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics4sizeEiRK7QStringiPi
  default:
    qtrt.ErrorResolve("QFontMetrics", "size", args)
 }

}


func (this *QFontMetrics) overlinePos(args ...interface{}) () {
  // overlinePos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics11overlinePosEv
  default:
    qtrt.ErrorResolve("QFontMetrics", "overlinePos", args)
 }

}


func (this *QFontMetrics) leading(args ...interface{}) () {
  // leading()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics7leadingEv
  default:
    qtrt.ErrorResolve("QFontMetrics", "leading", args)
 }

}


func (this *QFontMetrics) tightBoundingRect(args ...interface{}) () {
  // tightBoundingRect(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics17tightBoundingRectERK7QString
  default:
    qtrt.ErrorResolve("QFontMetrics", "tightBoundingRect", args)
 }

}


func (this *QFontMetrics) averageCharWidth(args ...interface{}) () {
  // averageCharWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics16averageCharWidthEv
  default:
    qtrt.ErrorResolve("QFontMetrics", "averageCharWidth", args)
 }

}


func (this *QFontMetrics) underlinePos(args ...interface{}) () {
  // underlinePos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics12underlinePosEv
  default:
    qtrt.ErrorResolve("QFontMetrics", "underlinePos", args)
 }

}


func (this *QFontMetrics) inFont(args ...interface{}) () {
  // inFont(class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics6inFontE5QChar
  default:
    qtrt.ErrorResolve("QFontMetrics", "inFont", args)
 }

}


func (this *QFontMetrics) height(args ...interface{}) () {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics6heightEv
  default:
    qtrt.ErrorResolve("QFontMetrics", "height", args)
 }

}


func (this *QFontMetrics) width(args ...interface{}) () {
  // width(class QChar)
  // width(const class QString &, int, int)
  // width(const class QString &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics5widthE5QChar
  case 1:
    // invoke: _ZNK12QFontMetrics5widthERK7QStringii
  case 2:
    // invoke: _ZNK12QFontMetrics5widthERK7QStringi
  default:
    qtrt.ErrorResolve("QFontMetrics", "width", args)
 }

}


func (this *QFontMetrics) xHeight(args ...interface{}) () {
  // xHeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics7xHeightEv
  default:
    qtrt.ErrorResolve("QFontMetrics", "xHeight", args)
 }

}


func (this *QFontMetrics) strikeOutPos(args ...interface{}) () {
  // strikeOutPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics12strikeOutPosEv
  default:
    qtrt.ErrorResolve("QFontMetrics", "strikeOutPos", args)
 }

}


func (this *QFontMetrics) lineSpacing(args ...interface{}) () {
  // lineSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics11lineSpacingEv
  default:
    qtrt.ErrorResolve("QFontMetrics", "lineSpacing", args)
 }

}


func NewQFontMetrics(args ...interface{})() {
}


func (this *QFontMetrics) minRightBearing(args ...interface{}) () {
  // minRightBearing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics15minRightBearingEv
  default:
    qtrt.ErrorResolve("QFontMetrics", "minRightBearing", args)
 }

}


func (this *QFontMetrics) swap(args ...interface{}) () {
  // swap(class QFontMetrics &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFontMetrics{}) // "QFontMetrics &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QFontMetrics4swapERS_
  default:
    qtrt.ErrorResolve("QFontMetrics", "swap", args)
 }

}


func (this *QFontMetrics) charWidth(args ...interface{}) () {
  // charWidth(const class QString &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics9charWidthERK7QStringi
  default:
    qtrt.ErrorResolve("QFontMetrics", "charWidth", args)
 }

}


func (this *QFontMetrics) leftBearing(args ...interface{}) () {
  // leftBearing(class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics11leftBearingE5QChar
  default:
    qtrt.ErrorResolve("QFontMetrics", "leftBearing", args)
 }

}


func (this *QFontMetrics) inFontUcs4(args ...interface{}) () {
  // inFontUcs4(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics10inFontUcs4Ej
  default:
    qtrt.ErrorResolve("QFontMetrics", "inFontUcs4", args)
 }

}


func (this *QFontMetrics) descent(args ...interface{}) () {
  // descent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics7descentEv
  default:
    qtrt.ErrorResolve("QFontMetrics", "descent", args)
 }

}


func (this *QFontMetricsF) inFont(args ...interface{}) () {
  // inFont(class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF6inFontE5QChar
  default:
    qtrt.ErrorResolve("QFontMetricsF", "inFont", args)
 }

}


func (this *QFontMetricsF) size(args ...interface{}) () {
  // size(int, const class QString &, int, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF4sizeEiRK7QStringiPi
  default:
    qtrt.ErrorResolve("QFontMetricsF", "size", args)
 }

}


func (this *QFontMetricsF) minRightBearing(args ...interface{}) () {
  // minRightBearing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF15minRightBearingEv
  default:
    qtrt.ErrorResolve("QFontMetricsF", "minRightBearing", args)
 }

}


func NewQFontMetricsF(args ...interface{})() {
}


func (this *QFontMetricsF) xHeight(args ...interface{}) () {
  // xHeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF7xHeightEv
  default:
    qtrt.ErrorResolve("QFontMetricsF", "xHeight", args)
 }

}


func (this *QFontMetricsF) width(args ...interface{}) () {
  // width(class QChar)
  // width(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF5widthE5QChar
  case 1:
    // invoke: _ZNK13QFontMetricsF5widthERK7QString
  default:
    qtrt.ErrorResolve("QFontMetricsF", "width", args)
 }

}


func (this *QFontMetricsF) FreeQFontMetricsF(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QFontMetricsF", "~QFontMetricsF", args)
 }

}


func (this *QFontMetricsF) boundingRect(args ...interface{}) () {
  // boundingRect(const class QRectF &, int, const class QString &, int, int *)
  // boundingRect(const class QString &)
  // boundingRect(class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(true) // "int *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF12boundingRectERK6QRectFiRK7QStringiPi
  case 1:
    // invoke: _ZNK13QFontMetricsF12boundingRectERK7QString
  case 2:
    // invoke: _ZNK13QFontMetricsF12boundingRectE5QChar
  default:
    qtrt.ErrorResolve("QFontMetricsF", "boundingRect", args)
 }

}


func (this *QFontMetricsF) swap(args ...interface{}) () {
  // swap(class QFontMetricsF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFontMetricsF{}) // "QFontMetricsF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QFontMetricsF4swapERS_
  default:
    qtrt.ErrorResolve("QFontMetricsF", "swap", args)
 }

}


func (this *QFontMetricsF) tightBoundingRect(args ...interface{}) () {
  // tightBoundingRect(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF17tightBoundingRectERK7QString
  default:
    qtrt.ErrorResolve("QFontMetricsF", "tightBoundingRect", args)
 }

}


func (this *QFontMetricsF) leftBearing(args ...interface{}) () {
  // leftBearing(class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF11leftBearingE5QChar
  default:
    qtrt.ErrorResolve("QFontMetricsF", "leftBearing", args)
 }

}


func (this *QFontMetricsF) rightBearing(args ...interface{}) () {
  // rightBearing(class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF12rightBearingE5QChar
  default:
    qtrt.ErrorResolve("QFontMetricsF", "rightBearing", args)
 }

}


func (this *QFontMetricsF) overlinePos(args ...interface{}) () {
  // overlinePos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF11overlinePosEv
  default:
    qtrt.ErrorResolve("QFontMetricsF", "overlinePos", args)
 }

}


func (this *QFontMetricsF) height(args ...interface{}) () {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF6heightEv
  default:
    qtrt.ErrorResolve("QFontMetricsF", "height", args)
 }

}


func (this *QFontMetricsF) descent(args ...interface{}) () {
  // descent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF7descentEv
  default:
    qtrt.ErrorResolve("QFontMetricsF", "descent", args)
 }

}


func (this *QFontMetricsF) lineWidth(args ...interface{}) () {
  // lineWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF9lineWidthEv
  default:
    qtrt.ErrorResolve("QFontMetricsF", "lineWidth", args)
 }

}


func (this *QFontMetricsF) strikeOutPos(args ...interface{}) () {
  // strikeOutPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF12strikeOutPosEv
  default:
    qtrt.ErrorResolve("QFontMetricsF", "strikeOutPos", args)
 }

}


func (this *QFontMetricsF) lineSpacing(args ...interface{}) () {
  // lineSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF11lineSpacingEv
  default:
    qtrt.ErrorResolve("QFontMetricsF", "lineSpacing", args)
 }

}


func (this *QFontMetricsF) averageCharWidth(args ...interface{}) () {
  // averageCharWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF16averageCharWidthEv
  default:
    qtrt.ErrorResolve("QFontMetricsF", "averageCharWidth", args)
 }

}


func (this *QFontMetricsF) leading(args ...interface{}) () {
  // leading()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF7leadingEv
  default:
    qtrt.ErrorResolve("QFontMetricsF", "leading", args)
 }

}


func (this *QFontMetricsF) inFontUcs4(args ...interface{}) () {
  // inFontUcs4(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF10inFontUcs4Ej
  default:
    qtrt.ErrorResolve("QFontMetricsF", "inFontUcs4", args)
 }

}


func (this *QFontMetricsF) minLeftBearing(args ...interface{}) () {
  // minLeftBearing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF14minLeftBearingEv
  default:
    qtrt.ErrorResolve("QFontMetricsF", "minLeftBearing", args)
 }

}


func (this *QFontMetricsF) ascent(args ...interface{}) () {
  // ascent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF6ascentEv
  default:
    qtrt.ErrorResolve("QFontMetricsF", "ascent", args)
 }

}


func (this *QFontMetricsF) maxWidth(args ...interface{}) () {
  // maxWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF8maxWidthEv
  default:
    qtrt.ErrorResolve("QFontMetricsF", "maxWidth", args)
 }

}


func (this *QFontMetricsF) underlinePos(args ...interface{}) () {
  // underlinePos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF12underlinePosEv
  default:
    qtrt.ErrorResolve("QFontMetricsF", "underlinePos", args)
 }

}

// <= body block end

