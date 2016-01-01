package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qbrush.h
// dst-file: /src/gui/qbrush.go
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
// class sizeof(QRadialGradient)=1
type QRadialGradient struct {
  /*qbase*/ QGradient;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QConicalGradient)=1
type QConicalGradient struct {
  /*qbase*/ QGradient;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QBrush)=1
type QBrush struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QGradient)=1
type QGradient struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QBrushData)=1
type QBrushData struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QLinearGradient)=1
type QLinearGradient struct {
  /*qbase*/ QGradient;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQRadialGradient(args ...interface{}) QRadialGradient {
  return QRadialGradient{}
}


func (this *QRadialGradient) setFocalPoint(args ...interface{}) () {
  // setFocalPoint(qreal, qreal)
  // setFocalPoint(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QRadialGradient13setFocalPointEdd
  case 1:
    // invoke: _ZN15QRadialGradient13setFocalPointERK7QPointF
  default:
    qtrt.ErrorResolve("QRadialGradient", "setFocalPoint", args)
  }

}


func (this *QRadialGradient) radius(args ...interface{}) () {
  // radius()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QRadialGradient6radiusEv
  default:
    qtrt.ErrorResolve("QRadialGradient", "radius", args)
  }

}


func (this *QRadialGradient) centerRadius(args ...interface{}) () {
  // centerRadius()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QRadialGradient12centerRadiusEv
  default:
    qtrt.ErrorResolve("QRadialGradient", "centerRadius", args)
  }

}


func (this *QRadialGradient) focalPoint(args ...interface{}) () {
  // focalPoint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QRadialGradient10focalPointEv
  default:
    qtrt.ErrorResolve("QRadialGradient", "focalPoint", args)
  }

}


func (this *QRadialGradient) focalRadius(args ...interface{}) () {
  // focalRadius()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QRadialGradient11focalRadiusEv
  default:
    qtrt.ErrorResolve("QRadialGradient", "focalRadius", args)
  }

}


func (this *QRadialGradient) center(args ...interface{}) () {
  // center()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QRadialGradient6centerEv
  default:
    qtrt.ErrorResolve("QRadialGradient", "center", args)
  }

}


func (this *QRadialGradient) setCenter(args ...interface{}) () {
  // setCenter(const class QPointF &)
  // setCenter(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QRadialGradient9setCenterERK7QPointF
  case 1:
    // invoke: _ZN15QRadialGradient9setCenterEdd
  default:
    qtrt.ErrorResolve("QRadialGradient", "setCenter", args)
  }

}


func (this *QRadialGradient) setCenterRadius(args ...interface{}) () {
  // setCenterRadius(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QRadialGradient15setCenterRadiusEd
  default:
    qtrt.ErrorResolve("QRadialGradient", "setCenterRadius", args)
  }

}


func (this *QRadialGradient) setFocalRadius(args ...interface{}) () {
  // setFocalRadius(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QRadialGradient14setFocalRadiusEd
  default:
    qtrt.ErrorResolve("QRadialGradient", "setFocalRadius", args)
  }

}


func (this *QRadialGradient) setRadius(args ...interface{}) () {
  // setRadius(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QRadialGradient9setRadiusEd
  default:
    qtrt.ErrorResolve("QRadialGradient", "setRadius", args)
  }

}


func (this *QConicalGradient) angle(args ...interface{}) () {
  // angle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QConicalGradient5angleEv
  default:
    qtrt.ErrorResolve("QConicalGradient", "angle", args)
  }

}


func (this *QConicalGradient) center(args ...interface{}) () {
  // center()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QConicalGradient6centerEv
  default:
    qtrt.ErrorResolve("QConicalGradient", "center", args)
  }

}


func NewQConicalGradient(args ...interface{}) QConicalGradient {
  return QConicalGradient{}
}


func (this *QConicalGradient) setAngle(args ...interface{}) () {
  // setAngle(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QConicalGradient8setAngleEd
  default:
    qtrt.ErrorResolve("QConicalGradient", "setAngle", args)
  }

}


func (this *QConicalGradient) setCenter(args ...interface{}) () {
  // setCenter(qreal, qreal)
  // setCenter(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QConicalGradient9setCenterEdd
  case 1:
    // invoke: _ZN16QConicalGradient9setCenterERK7QPointF
  default:
    qtrt.ErrorResolve("QConicalGradient", "setCenter", args)
  }

}


func NewQBrush(args ...interface{}) QBrush {
  return QBrush{}
}


func (this *QBrush) setTexture(args ...interface{}) () {
  // setTexture(const class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QBrush10setTextureERK7QPixmap
  default:
    qtrt.ErrorResolve("QBrush", "setTexture", args)
  }

}


func (this *QBrush) setTextureImage(args ...interface{}) () {
  // setTextureImage(const class QImage &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QImage{}) // "const QImage &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QBrush15setTextureImageERK6QImage
  default:
    qtrt.ErrorResolve("QBrush", "setTextureImage", args)
  }

}


func (this *QBrush) texture(args ...interface{}) () {
  // texture()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QBrush7textureEv
  default:
    qtrt.ErrorResolve("QBrush", "texture", args)
  }

}


func (this *QBrush) transform(args ...interface{}) () {
  // transform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QBrush9transformEv
  default:
    qtrt.ErrorResolve("QBrush", "transform", args)
  }

}


func (this *QBrush) setTransform(args ...interface{}) () {
  // setTransform(const class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QBrush12setTransformERK10QTransform
  default:
    qtrt.ErrorResolve("QBrush", "setTransform", args)
  }

}


func (this *QBrush) isOpaque(args ...interface{}) () {
  // isOpaque()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QBrush8isOpaqueEv
  default:
    qtrt.ErrorResolve("QBrush", "isOpaque", args)
  }

}


func (this *QBrush) gradient(args ...interface{}) () {
  // gradient()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QBrush8gradientEv
  default:
    qtrt.ErrorResolve("QBrush", "gradient", args)
  }

}


func (this *QBrush) FreeQBrush(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QBrush", "~QBrush", args)
  }

}


func (this *QBrush) setMatrix(args ...interface{}) () {
  // setMatrix(const class QMatrix &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix{}) // "const QMatrix &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QBrush9setMatrixERK7QMatrix
  default:
    qtrt.ErrorResolve("QBrush", "setMatrix", args)
  }

}


func (this *QBrush) setColor(args ...interface{}) () {
  // setColor(const class QColor &)
  // setColor(Qt::GlobalColor)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "Qt::GlobalColor"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QBrush8setColorERK6QColor
  case 1:
    // invoke: _ZN6QBrush8setColorEN2Qt11GlobalColorE
  default:
    qtrt.ErrorResolve("QBrush", "setColor", args)
  }

}


func (this *QBrush) matrix(args ...interface{}) () {
  // matrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QBrush6matrixEv
  default:
    qtrt.ErrorResolve("QBrush", "matrix", args)
  }

}


func (this *QBrush) textureImage(args ...interface{}) () {
  // textureImage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QBrush12textureImageEv
  default:
    qtrt.ErrorResolve("QBrush", "textureImage", args)
  }

}


func (this *QBrush) isDetached(args ...interface{}) () {
  // isDetached()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QBrush10isDetachedEv
  default:
    qtrt.ErrorResolve("QBrush", "isDetached", args)
  }

}


func (this *QBrush) swap(args ...interface{}) () {
  // swap(class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QBrush4swapERS_
  default:
    qtrt.ErrorResolve("QBrush", "swap", args)
  }

}


func (this *QBrush) color(args ...interface{}) () {
  // color()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QBrush5colorEv
  default:
    qtrt.ErrorResolve("QBrush", "color", args)
  }

}


func (this *QGradient) setColorAt(args ...interface{}) () {
  // setColorAt(qreal, const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGradient10setColorAtEdRK6QColor
  default:
    qtrt.ErrorResolve("QGradient", "setColorAt", args)
  }

}


func (this *QGradient) stops(args ...interface{}) () {
  // stops()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGradient5stopsEv
  default:
    qtrt.ErrorResolve("QGradient", "stops", args)
  }

}


func NewQGradient(args ...interface{}) QGradient {
  return QGradient{}
}


func (this *QLinearGradient) setFinalStop(args ...interface{}) () {
  // setFinalStop(qreal, qreal)
  // setFinalStop(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QLinearGradient12setFinalStopEdd
  case 1:
    // invoke: _ZN15QLinearGradient12setFinalStopERK7QPointF
  default:
    qtrt.ErrorResolve("QLinearGradient", "setFinalStop", args)
  }

}


func (this *QLinearGradient) start(args ...interface{}) () {
  // start()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QLinearGradient5startEv
  default:
    qtrt.ErrorResolve("QLinearGradient", "start", args)
  }

}


func NewQLinearGradient(args ...interface{}) QLinearGradient {
  return QLinearGradient{}
}


func (this *QLinearGradient) setStart(args ...interface{}) () {
  // setStart(qreal, qreal)
  // setStart(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QLinearGradient8setStartEdd
  case 1:
    // invoke: _ZN15QLinearGradient8setStartERK7QPointF
  default:
    qtrt.ErrorResolve("QLinearGradient", "setStart", args)
  }

}


func (this *QLinearGradient) finalStop(args ...interface{}) () {
  // finalStop()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QLinearGradient9finalStopEv
  default:
    qtrt.ErrorResolve("QLinearGradient", "finalStop", args)
  }

}

// <= body block end

