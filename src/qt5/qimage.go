package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qimage.h
// dst-file: /src/gui/qimage.go
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
// class sizeof(QImage)=32
type QImage struct {
  /*qbase*/ QPaintDevice;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QImage) copy(args ...interface{}) () {
  // copy(const class QRect &)
  // copy(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage4copyERK5QRect
  case 1:
    // invoke: _ZNK6QImage4copyEiiii
  default:
    qtrt.ErrorResolve("QImage", "copy", args)
  }

}


func (this *QImage) trueMatrix_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QImage", "trueMatrix", args)
  }

}


func (this *QImage) bits(args ...interface{}) () {
  // bits()
  // bits()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage4bitsEv
  case 1:
    // invoke: _ZNK6QImage4bitsEv
  default:
    qtrt.ErrorResolve("QImage", "bits", args)
  }

}


func (this *QImage) setAlphaChannel(args ...interface{}) () {
  // setAlphaChannel(const class QImage &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QImage{}) // "const QImage &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage15setAlphaChannelERKS_
  default:
    qtrt.ErrorResolve("QImage", "setAlphaChannel", args)
  }

}


func (this *QImage) text(args ...interface{}) () {
  // text(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage4textERK7QString
  default:
    qtrt.ErrorResolve("QImage", "text", args)
  }

}


func (this *QImage) rect(args ...interface{}) () {
  // rect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage4rectEv
  default:
    qtrt.ErrorResolve("QImage", "rect", args)
  }

}


func NewQImage(args ...interface{}) QImage {
  return QImage{}
}


func (this *QImage) createHeuristicMask(args ...interface{}) () {
  // createHeuristicMask(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage19createHeuristicMaskEb
  default:
    qtrt.ErrorResolve("QImage", "createHeuristicMask", args)
  }

}


func (this *QImage) constBits(args ...interface{}) () {
  // constBits()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage9constBitsEv
  default:
    qtrt.ErrorResolve("QImage", "constBits", args)
  }

}


func (this *QImage) fromData_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QImage", "fromData", args)
  }

}


func (this *QImage) isDetached(args ...interface{}) () {
  // isDetached()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage10isDetachedEv
  default:
    qtrt.ErrorResolve("QImage", "isDetached", args)
  }

}


func (this *QImage) setOffset(args ...interface{}) () {
  // setOffset(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage9setOffsetERK6QPoint
  default:
    qtrt.ErrorResolve("QImage", "setOffset", args)
  }

}


func (this *QImage) isGrayscale(args ...interface{}) () {
  // isGrayscale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage11isGrayscaleEv
  default:
    qtrt.ErrorResolve("QImage", "isGrayscale", args)
  }

}


func (this *QImage) save(args ...interface{}) () {
  // save(class QIODevice *, const char *, int)
  // save(const class QString &, const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = qtrt.ByteTy(true) // "const char *"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage4saveEP9QIODevicePKci
  case 1:
    // invoke: _ZNK6QImage4saveERK7QStringPKci
  default:
    qtrt.ErrorResolve("QImage", "save", args)
  }

}


func (this *QImage) depth(args ...interface{}) () {
  // depth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage5depthEv
  default:
    qtrt.ErrorResolve("QImage", "depth", args)
  }

}


func (this *QImage) alphaChannel(args ...interface{}) () {
  // alphaChannel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage12alphaChannelEv
  default:
    qtrt.ErrorResolve("QImage", "alphaChannel", args)
  }

}


func (this *QImage) hasAlphaChannel(args ...interface{}) () {
  // hasAlphaChannel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage15hasAlphaChannelEv
  default:
    qtrt.ErrorResolve("QImage", "hasAlphaChannel", args)
  }

}


func (this *QImage) loadFromData(args ...interface{}) () {
  // loadFromData(const uchar *, int, const char *)
  // loadFromData(const class QByteArray &, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const uchar *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage12loadFromDataEPKhiPKc
  case 1:
    // invoke: _ZN6QImage12loadFromDataERK10QByteArrayPKc
  default:
    qtrt.ErrorResolve("QImage", "loadFromData", args)
  }

}


func (this *QImage) colorCount(args ...interface{}) () {
  // colorCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage10colorCountEv
  default:
    qtrt.ErrorResolve("QImage", "colorCount", args)
  }

}


func (this *QImage) allGray(args ...interface{}) () {
  // allGray()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage7allGrayEv
  default:
    qtrt.ErrorResolve("QImage", "allGray", args)
  }

}


func (this *QImage) setColorCount(args ...interface{}) () {
  // setColorCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage13setColorCountEi
  default:
    qtrt.ErrorResolve("QImage", "setColorCount", args)
  }

}


func (this *QImage) pixel(args ...interface{}) () {
  // pixel(const class QPoint &)
  // pixel(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage5pixelERK6QPoint
  case 1:
    // invoke: _ZNK6QImage5pixelEii
  default:
    qtrt.ErrorResolve("QImage", "pixel", args)
  }

}


func (this *QImage) setDevicePixelRatio(args ...interface{}) () {
  // setDevicePixelRatio(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage19setDevicePixelRatioEd
  default:
    qtrt.ErrorResolve("QImage", "setDevicePixelRatio", args)
  }

}


func (this *QImage) setText(args ...interface{}) () {
  // setText(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage7setTextERK7QStringS2_
  default:
    qtrt.ErrorResolve("QImage", "setText", args)
  }

}


func (this *QImage) color(args ...interface{}) () {
  // color(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage5colorEi
  default:
    qtrt.ErrorResolve("QImage", "color", args)
  }

}


func (this *QImage) setPixel(args ...interface{}) () {
  // setPixel(const class QPoint &, uint)
  // setPixel(int, int, uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[0][1] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage8setPixelERK6QPointj
  case 1:
    // invoke: _ZN6QImage8setPixelEiij
  default:
    qtrt.ErrorResolve("QImage", "setPixel", args)
  }

}


func (this *QImage) offset(args ...interface{}) () {
  // offset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage6offsetEv
  default:
    qtrt.ErrorResolve("QImage", "offset", args)
  }

}


func (this *QImage) constScanLine(args ...interface{}) () {
  // constScanLine(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage13constScanLineEi
  default:
    qtrt.ErrorResolve("QImage", "constScanLine", args)
  }

}


func (this *QImage) textKeys(args ...interface{}) () {
  // textKeys()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage8textKeysEv
  default:
    qtrt.ErrorResolve("QImage", "textKeys", args)
  }

}


func (this *QImage) dotsPerMeterY(args ...interface{}) () {
  // dotsPerMeterY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage13dotsPerMeterYEv
  default:
    qtrt.ErrorResolve("QImage", "dotsPerMeterY", args)
  }

}


func (this *QImage) fill(args ...interface{}) () {
  // fill(uint)
  // fill(const class QColor &)
  // fill(Qt::GlobalColor)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "Qt::GlobalColor"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage4fillEj
  case 1:
    // invoke: _ZN6QImage4fillERK6QColor
  case 2:
    // invoke: _ZN6QImage4fillEN2Qt11GlobalColorE
  default:
    qtrt.ErrorResolve("QImage", "fill", args)
  }

}


func (this *QImage) pixelFormat(args ...interface{}) () {
  // pixelFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage11pixelFormatEv
  default:
    qtrt.ErrorResolve("QImage", "pixelFormat", args)
  }

}


func (this *QImage) dotsPerMeterX(args ...interface{}) () {
  // dotsPerMeterX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage13dotsPerMeterXEv
  default:
    qtrt.ErrorResolve("QImage", "dotsPerMeterX", args)
  }

}


func (this *QImage) setDotsPerMeterY(args ...interface{}) () {
  // setDotsPerMeterY(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage16setDotsPerMeterYEi
  default:
    qtrt.ErrorResolve("QImage", "setDotsPerMeterY", args)
  }

}


func (this *QImage) bitPlaneCount(args ...interface{}) () {
  // bitPlaneCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage13bitPlaneCountEv
  default:
    qtrt.ErrorResolve("QImage", "bitPlaneCount", args)
  }

}


func (this *QImage) detach(args ...interface{}) () {
  // detach()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage6detachEv
  default:
    qtrt.ErrorResolve("QImage", "detach", args)
  }

}


func (this *QImage) paintEngine(args ...interface{}) () {
  // paintEngine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage11paintEngineEv
  default:
    qtrt.ErrorResolve("QImage", "paintEngine", args)
  }

}


func (this *QImage) swap(args ...interface{}) () {
  // swap(class QImage &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QImage{}) // "QImage &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage4swapERS_
  default:
    qtrt.ErrorResolve("QImage", "swap", args)
  }

}


func (this *QImage) devicePixelRatio(args ...interface{}) () {
  // devicePixelRatio()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage16devicePixelRatioEv
  default:
    qtrt.ErrorResolve("QImage", "devicePixelRatio", args)
  }

}


func (this *QImage) devType(args ...interface{}) () {
  // devType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage7devTypeEv
  default:
    qtrt.ErrorResolve("QImage", "devType", args)
  }

}


func (this *QImage) valid(args ...interface{}) () {
  // valid(const class QPoint &)
  // valid(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage5validERK6QPoint
  case 1:
    // invoke: _ZNK6QImage5validEii
  default:
    qtrt.ErrorResolve("QImage", "valid", args)
  }

}


func (this *QImage) pixelIndex(args ...interface{}) () {
  // pixelIndex(const class QPoint &)
  // pixelIndex(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage10pixelIndexERK6QPoint
  case 1:
    // invoke: _ZNK6QImage10pixelIndexEii
  default:
    qtrt.ErrorResolve("QImage", "pixelIndex", args)
  }

}


func (this *QImage) setDotsPerMeterX(args ...interface{}) () {
  // setDotsPerMeterX(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage16setDotsPerMeterXEi
  default:
    qtrt.ErrorResolve("QImage", "setDotsPerMeterX", args)
  }

}


func (this *QImage) load(args ...interface{}) () {
  // load(const class QString &, const char *)
  // load(class QIODevice *, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"
  vtys[1][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage4loadERK7QStringPKc
  case 1:
    // invoke: _ZN6QImage4loadEP9QIODevicePKc
  default:
    qtrt.ErrorResolve("QImage", "load", args)
  }

}


func (this *QImage) colorTable(args ...interface{}) () {
  // colorTable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage10colorTableEv
  default:
    qtrt.ErrorResolve("QImage", "colorTable", args)
  }

}


func (this *QImage) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage4sizeEv
  default:
    qtrt.ErrorResolve("QImage", "size", args)
  }

}


func (this *QImage) height(args ...interface{}) () {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage6heightEv
  default:
    qtrt.ErrorResolve("QImage", "height", args)
  }

}


func (this *QImage) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage5widthEv
  default:
    qtrt.ErrorResolve("QImage", "width", args)
  }

}


func (this *QImage) scanLine(args ...interface{}) () {
  // scanLine(int)
  // scanLine(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage8scanLineEi
  case 1:
    // invoke: _ZN6QImage8scanLineEi
  default:
    qtrt.ErrorResolve("QImage", "scanLine", args)
  }

}


func (this *QImage) bytesPerLine(args ...interface{}) () {
  // bytesPerLine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage12bytesPerLineEv
  default:
    qtrt.ErrorResolve("QImage", "bytesPerLine", args)
  }

}


func (this *QImage) cacheKey(args ...interface{}) () {
  // cacheKey()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage8cacheKeyEv
  default:
    qtrt.ErrorResolve("QImage", "cacheKey", args)
  }

}


func (this *QImage) FreeQImage(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QImage", "~QImage", args)
  }

}


func (this *QImage) setColor(args ...interface{}) () {
  // setColor(int, QRgb)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "QRgb"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage8setColorEij
  default:
    qtrt.ErrorResolve("QImage", "setColor", args)
  }

}


func (this *QImage) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage6isNullEv
  default:
    qtrt.ErrorResolve("QImage", "isNull", args)
  }

}


func (this *QImage) byteCount(args ...interface{}) () {
  // byteCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage9byteCountEv
  default:
    qtrt.ErrorResolve("QImage", "byteCount", args)
  }

}

// <= body block end

