package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qpixmap.h
// dst-file: /src/gui/qpixmap.go
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
// class sizeof(QPixmap)=1
type QPixmap struct {
  /*qbase*/ QPaintDevice;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QPixmap) save(args ...interface{}) () {
  // save(const class QString &, const char *, int)
  // save(class QIODevice *, const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"
  vtys[1][1] = qtrt.ByteTy(true) // "const char *"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap4saveERK7QStringPKci
  case 1:
    // invoke: _ZNK7QPixmap4saveEP9QIODevicePKci
  default:
    qtrt.ErrorResolve("QPixmap", "save", args)
  }

}


func (this *QPixmap) swap(args ...interface{}) () {
  // swap(class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPixmap{}) // "QPixmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmap4swapERS_
  default:
    qtrt.ErrorResolve("QPixmap", "swap", args)
  }

}


func (this *QPixmap) isQBitmap(args ...interface{}) () {
  // isQBitmap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap9isQBitmapEv
  default:
    qtrt.ErrorResolve("QPixmap", "isQBitmap", args)
  }

}


func (this *QPixmap) devicePixelRatio(args ...interface{}) () {
  // devicePixelRatio()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap16devicePixelRatioEv
  default:
    qtrt.ErrorResolve("QPixmap", "devicePixelRatio", args)
  }

}


func NewQPixmap(args ...interface{}) QPixmap {
  return QPixmap{}
}


func (this *QPixmap) fill(args ...interface{}) () {
  // fill(const class QPaintDevice *, int, int)
  // fill(const class QColor &)
  // fill(const class QPaintDevice *, const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPaintDevice{}) // "const QPaintDevice *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPaintDevice{}) // "const QPaintDevice *"
  vtys[2][1] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmap4fillEPK12QPaintDeviceii
  case 1:
    // invoke: _ZN7QPixmap4fillERK6QColor
  case 2:
    // invoke: _ZN7QPixmap4fillEPK12QPaintDeviceRK6QPoint
  default:
    qtrt.ErrorResolve("QPixmap", "fill", args)
  }

}


func (this *QPixmap) devType(args ...interface{}) () {
  // devType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap7devTypeEv
  default:
    qtrt.ErrorResolve("QPixmap", "devType", args)
  }

}


func (this *QPixmap) scroll(args ...interface{}) () {
  // scroll(int, int, int, int, int, int, class QRegion *)
  // scroll(int, int, const class QRect &, class QRegion *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"
  vtys[0][5] = qtrt.Int32Ty(false) // "int"
  vtys[0][6] = reflect.TypeOf(QRegion{}) // "QRegion *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1][3] = reflect.TypeOf(QRegion{}) // "QRegion *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmap6scrollEiiiiiiP7QRegion
  case 1:
    // invoke: _ZN7QPixmap6scrollEiiRK5QRectP7QRegion
  default:
    qtrt.ErrorResolve("QPixmap", "scroll", args)
  }

}


func (this *QPixmap) copy(args ...interface{}) () {
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
    // invoke: _ZNK7QPixmap4copyERK5QRect
  case 1:
    // invoke: _ZNK7QPixmap4copyEiiii
  default:
    qtrt.ErrorResolve("QPixmap", "copy", args)
  }

}


func (this *QPixmap) trueMatrix_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPixmap", "trueMatrix", args)
  }

}


func (this *QPixmap) grabWindow_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPixmap", "grabWindow", args)
  }

}


func (this *QPixmap) isDetached(args ...interface{}) () {
  // isDetached()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap10isDetachedEv
  default:
    qtrt.ErrorResolve("QPixmap", "isDetached", args)
  }

}


func (this *QPixmap) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap6isNullEv
  default:
    qtrt.ErrorResolve("QPixmap", "isNull", args)
  }

}


func (this *QPixmap) defaultDepth_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPixmap", "defaultDepth", args)
  }

}


func (this *QPixmap) detach(args ...interface{}) () {
  // detach()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmap6detachEv
  default:
    qtrt.ErrorResolve("QPixmap", "detach", args)
  }

}


func (this *QPixmap) setMask(args ...interface{}) () {
  // setMask(const class QBitmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBitmap{}) // "const QBitmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmap7setMaskERK7QBitmap
  default:
    qtrt.ErrorResolve("QPixmap", "setMask", args)
  }

}


func (this *QPixmap) grabWidget_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPixmap", "grabWidget", args)
  }

}


func (this *QPixmap) setDevicePixelRatio(args ...interface{}) () {
  // setDevicePixelRatio(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmap19setDevicePixelRatioEd
  default:
    qtrt.ErrorResolve("QPixmap", "setDevicePixelRatio", args)
  }

}


func (this *QPixmap) cacheKey(args ...interface{}) () {
  // cacheKey()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap8cacheKeyEv
  default:
    qtrt.ErrorResolve("QPixmap", "cacheKey", args)
  }

}


func (this *QPixmap) createHeuristicMask(args ...interface{}) () {
  // createHeuristicMask(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap19createHeuristicMaskEb
  default:
    qtrt.ErrorResolve("QPixmap", "createHeuristicMask", args)
  }

}


func (this *QPixmap) depth(args ...interface{}) () {
  // depth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap5depthEv
  default:
    qtrt.ErrorResolve("QPixmap", "depth", args)
  }

}


func (this *QPixmap) toImage(args ...interface{}) () {
  // toImage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap7toImageEv
  default:
    qtrt.ErrorResolve("QPixmap", "toImage", args)
  }

}


func (this *QPixmap) handle(args ...interface{}) () {
  // handle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap6handleEv
  default:
    qtrt.ErrorResolve("QPixmap", "handle", args)
  }

}


func (this *QPixmap) hasAlphaChannel(args ...interface{}) () {
  // hasAlphaChannel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap15hasAlphaChannelEv
  default:
    qtrt.ErrorResolve("QPixmap", "hasAlphaChannel", args)
  }

}


func (this *QPixmap) rect(args ...interface{}) () {
  // rect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap4rectEv
  default:
    qtrt.ErrorResolve("QPixmap", "rect", args)
  }

}


func (this *QPixmap) mask(args ...interface{}) () {
  // mask()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap4maskEv
  default:
    qtrt.ErrorResolve("QPixmap", "mask", args)
  }

}


func (this *QPixmap) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap5widthEv
  default:
    qtrt.ErrorResolve("QPixmap", "width", args)
  }

}


func (this *QPixmap) paintEngine(args ...interface{}) () {
  // paintEngine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap11paintEngineEv
  default:
    qtrt.ErrorResolve("QPixmap", "paintEngine", args)
  }

}


func (this *QPixmap) FreeQPixmap(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPixmap", "~QPixmap", args)
  }

}


func (this *QPixmap) height(args ...interface{}) () {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap6heightEv
  default:
    qtrt.ErrorResolve("QPixmap", "height", args)
  }

}


func (this *QPixmap) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap4sizeEv
  default:
    qtrt.ErrorResolve("QPixmap", "size", args)
  }

}


func (this *QPixmap) hasAlpha(args ...interface{}) () {
  // hasAlpha()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap8hasAlphaEv
  default:
    qtrt.ErrorResolve("QPixmap", "hasAlpha", args)
  }

}

// <= body block end

