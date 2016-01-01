package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qpainterpath.h
// dst-file: /src/gui/qpainterpath.go
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
// class sizeof(QPainterPath)=1
type QPainterPath struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QPainterPathStroker)=1
type QPainterPathStroker struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QPainterPath) setElementPositionAt(args ...interface{}) () {
  // setElementPositionAt(int, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath20setElementPositionAtEidd
  default:
    qtrt.ErrorResolve("QPainterPath", "setElementPositionAt", args)
  }

}


func (this *QPainterPath) toFillPolygon(args ...interface{}) () {
  // toFillPolygon(const class QMatrix &)
  // toFillPolygon(const class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix{}) // "const QMatrix &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath13toFillPolygonERK7QMatrix
  case 1:
    // invoke: _ZNK12QPainterPath13toFillPolygonERK10QTransform
  default:
    qtrt.ErrorResolve("QPainterPath", "toFillPolygon", args)
  }

}


func (this *QPainterPath) translated(args ...interface{}) () {
  // translated(qreal, qreal)
  // translated(const class QPointF &)
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
    // invoke: _ZNK12QPainterPath10translatedEdd
  case 1:
    // invoke: _ZNK12QPainterPath10translatedERK7QPointF
  default:
    qtrt.ErrorResolve("QPainterPath", "translated", args)
  }

}


func (this *QPainterPath) FreeQPainterPath(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPainterPath", "~QPainterPath", args)
  }

}


func (this *QPainterPath) toSubpathPolygons(args ...interface{}) () {
  // toSubpathPolygons(const class QTransform &)
  // toSubpathPolygons(const class QMatrix &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QMatrix{}) // "const QMatrix &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath17toSubpathPolygonsERK10QTransform
  case 1:
    // invoke: _ZNK12QPainterPath17toSubpathPolygonsERK7QMatrix
  default:
    qtrt.ErrorResolve("QPainterPath", "toSubpathPolygons", args)
  }

}


func (this *QPainterPath) controlPointRect(args ...interface{}) () {
  // controlPointRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath16controlPointRectEv
  default:
    qtrt.ErrorResolve("QPainterPath", "controlPointRect", args)
  }

}


func (this *QPainterPath) toFillPolygons(args ...interface{}) () {
  // toFillPolygons(const class QMatrix &)
  // toFillPolygons(const class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix{}) // "const QMatrix &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath14toFillPolygonsERK7QMatrix
  case 1:
    // invoke: _ZNK12QPainterPath14toFillPolygonsERK10QTransform
  default:
    qtrt.ErrorResolve("QPainterPath", "toFillPolygons", args)
  }

}


func (this *QPainterPath) quadTo(args ...interface{}) () {
  // quadTo(const class QPointF &, const class QPointF &)
  // quadTo(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath6quadToERK7QPointFS2_
  case 1:
    // invoke: _ZN12QPainterPath6quadToEdddd
  default:
    qtrt.ErrorResolve("QPainterPath", "quadTo", args)
  }

}


func (this *QPainterPath) arcTo(args ...interface{}) () {
  // arcTo(qreal, qreal, qreal, qreal, qreal, qreal)
  // arcTo(const class QRectF &, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][5] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath5arcToEdddddd
  case 1:
    // invoke: _ZN12QPainterPath5arcToERK6QRectFdd
  default:
    qtrt.ErrorResolve("QPainterPath", "arcTo", args)
  }

}


func (this *QPainterPath) addRect(args ...interface{}) () {
  // addRect(const class QRectF &)
  // addRect(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath7addRectERK6QRectF
  case 1:
    // invoke: _ZN12QPainterPath7addRectEdddd
  default:
    qtrt.ErrorResolve("QPainterPath", "addRect", args)
  }

}


func (this *QPainterPath) addRoundRect(args ...interface{}) () {
  // addRoundRect(const class QRectF &, int, int)
  // addRoundRect(const class QRectF &, int)
  // addRoundRect(qreal, qreal, qreal, qreal, int)
  // addRoundRect(qreal, qreal, qreal, qreal, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][4] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][4] = qtrt.Int32Ty(false) // "int"
  vtys[3][5] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath12addRoundRectERK6QRectFii
  case 1:
    // invoke: _ZN12QPainterPath12addRoundRectERK6QRectFi
  case 2:
    // invoke: _ZN12QPainterPath12addRoundRectEddddi
  case 3:
    // invoke: _ZN12QPainterPath12addRoundRectEddddii
  default:
    qtrt.ErrorResolve("QPainterPath", "addRoundRect", args)
  }

}


func (this *QPainterPath) addText(args ...interface{}) () {
  // addText(qreal, qreal, const class QFont &, const class QString &)
  // addText(const class QPointF &, const class QFont &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = reflect.TypeOf(QFont{}) // "const QFont &"
  vtys[0][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1][1] = reflect.TypeOf(QFont{}) // "const QFont &"
  vtys[1][2] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath7addTextEddRK5QFontRK7QString
  case 1:
    // invoke: _ZN12QPainterPath7addTextERK7QPointFRK5QFontRK7QString
  default:
    qtrt.ErrorResolve("QPainterPath", "addText", args)
  }

}


func (this *QPainterPath) intersects(args ...interface{}) () {
  // intersects(const class QRectF &)
  // intersects(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath10intersectsERK6QRectF
  case 1:
    // invoke: _ZNK12QPainterPath10intersectsERKS_
  default:
    qtrt.ErrorResolve("QPainterPath", "intersects", args)
  }

}


func (this *QPainterPath) contains(args ...interface{}) () {
  // contains(const class QPointF &)
  // contains(const class QRectF &)
  // contains(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath8containsERK7QPointF
  case 1:
    // invoke: _ZNK12QPainterPath8containsERK6QRectF
  case 2:
    // invoke: _ZNK12QPainterPath8containsERKS_
  default:
    qtrt.ErrorResolve("QPainterPath", "contains", args)
  }

}


func (this *QPainterPath) addEllipse(args ...interface{}) () {
  // addEllipse(const class QPointF &, qreal, qreal)
  // addEllipse(qreal, qreal, qreal, qreal)
  // addEllipse(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath10addEllipseERK7QPointFdd
  case 1:
    // invoke: _ZN12QPainterPath10addEllipseEdddd
  case 2:
    // invoke: _ZN12QPainterPath10addEllipseERK6QRectF
  default:
    qtrt.ErrorResolve("QPainterPath", "addEllipse", args)
  }

}


func (this *QPainterPath) lineTo(args ...interface{}) () {
  // lineTo(qreal, qreal)
  // lineTo(const class QPointF &)
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
    // invoke: _ZN12QPainterPath6lineToEdd
  case 1:
    // invoke: _ZN12QPainterPath6lineToERK7QPointF
  default:
    qtrt.ErrorResolve("QPainterPath", "lineTo", args)
  }

}


func (this *QPainterPath) cubicTo(args ...interface{}) () {
  // cubicTo(const class QPointF &, const class QPointF &, const class QPointF &)
  // cubicTo(qreal, qreal, qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][2] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][5] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath7cubicToERK7QPointFS2_S2_
  case 1:
    // invoke: _ZN12QPainterPath7cubicToEdddddd
  default:
    qtrt.ErrorResolve("QPainterPath", "cubicTo", args)
  }

}


func (this *QPainterPath) slopeAtPercent(args ...interface{}) () {
  // slopeAtPercent(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath14slopeAtPercentEd
  default:
    qtrt.ErrorResolve("QPainterPath", "slopeAtPercent", args)
  }

}


func NewQPainterPath(args ...interface{}) QPainterPath {
  return QPainterPath{}
}


func (this *QPainterPath) intersected(args ...interface{}) () {
  // intersected(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath11intersectedERKS_
  default:
    qtrt.ErrorResolve("QPainterPath", "intersected", args)
  }

}


func (this *QPainterPath) translate(args ...interface{}) () {
  // translate(qreal, qreal)
  // translate(const class QPointF &)
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
    // invoke: _ZN12QPainterPath9translateEdd
  case 1:
    // invoke: _ZN12QPainterPath9translateERK7QPointF
  default:
    qtrt.ErrorResolve("QPainterPath", "translate", args)
  }

}


func (this *QPainterPath) addPolygon(args ...interface{}) () {
  // addPolygon(const class QPolygonF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath10addPolygonERK9QPolygonF
  default:
    qtrt.ErrorResolve("QPainterPath", "addPolygon", args)
  }

}


func (this *QPainterPath) addPath(args ...interface{}) () {
  // addPath(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath7addPathERKS_
  default:
    qtrt.ErrorResolve("QPainterPath", "addPath", args)
  }

}


func (this *QPainterPath) elementCount(args ...interface{}) () {
  // elementCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath12elementCountEv
  default:
    qtrt.ErrorResolve("QPainterPath", "elementCount", args)
  }

}


func (this *QPainterPath) simplified(args ...interface{}) () {
  // simplified()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath10simplifiedEv
  default:
    qtrt.ErrorResolve("QPainterPath", "simplified", args)
  }

}


func (this *QPainterPath) length(args ...interface{}) () {
  // length()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath6lengthEv
  default:
    qtrt.ErrorResolve("QPainterPath", "length", args)
  }

}


func (this *QPainterPath) connectPath(args ...interface{}) () {
  // connectPath(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath11connectPathERKS_
  default:
    qtrt.ErrorResolve("QPainterPath", "connectPath", args)
  }

}


func (this *QPainterPath) addRegion(args ...interface{}) () {
  // addRegion(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath9addRegionERK7QRegion
  default:
    qtrt.ErrorResolve("QPainterPath", "addRegion", args)
  }

}


func (this *QPainterPath) currentPosition(args ...interface{}) () {
  // currentPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath15currentPositionEv
  default:
    qtrt.ErrorResolve("QPainterPath", "currentPosition", args)
  }

}


func (this *QPainterPath) toReversed(args ...interface{}) () {
  // toReversed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath10toReversedEv
  default:
    qtrt.ErrorResolve("QPainterPath", "toReversed", args)
  }

}


func (this *QPainterPath) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath12boundingRectEv
  default:
    qtrt.ErrorResolve("QPainterPath", "boundingRect", args)
  }

}


func (this *QPainterPath) swap(args ...interface{}) () {
  // swap(class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath4swapERS_
  default:
    qtrt.ErrorResolve("QPainterPath", "swap", args)
  }

}


func (this *QPainterPath) moveTo(args ...interface{}) () {
  // moveTo(qreal, qreal)
  // moveTo(const class QPointF &)
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
    // invoke: _ZN12QPainterPath6moveToEdd
  case 1:
    // invoke: _ZN12QPainterPath6moveToERK7QPointF
  default:
    qtrt.ErrorResolve("QPainterPath", "moveTo", args)
  }

}


func (this *QPainterPath) subtracted(args ...interface{}) () {
  // subtracted(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath10subtractedERKS_
  default:
    qtrt.ErrorResolve("QPainterPath", "subtracted", args)
  }

}


func (this *QPainterPath) pointAtPercent(args ...interface{}) () {
  // pointAtPercent(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath14pointAtPercentEd
  default:
    qtrt.ErrorResolve("QPainterPath", "pointAtPercent", args)
  }

}


func (this *QPainterPath) percentAtLength(args ...interface{}) () {
  // percentAtLength(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath15percentAtLengthEd
  default:
    qtrt.ErrorResolve("QPainterPath", "percentAtLength", args)
  }

}


func (this *QPainterPath) subtractedInverted(args ...interface{}) () {
  // subtractedInverted(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath18subtractedInvertedERKS_
  default:
    qtrt.ErrorResolve("QPainterPath", "subtractedInverted", args)
  }

}


func (this *QPainterPath) arcMoveTo(args ...interface{}) () {
  // arcMoveTo(qreal, qreal, qreal, qreal, qreal)
  // arcMoveTo(const class QRectF &, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath9arcMoveToEddddd
  case 1:
    // invoke: _ZN12QPainterPath9arcMoveToERK6QRectFd
  default:
    qtrt.ErrorResolve("QPainterPath", "arcMoveTo", args)
  }

}


func (this *QPainterPath) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath7isEmptyEv
  default:
    qtrt.ErrorResolve("QPainterPath", "isEmpty", args)
  }

}


func (this *QPainterPath) united(args ...interface{}) () {
  // united(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath6unitedERKS_
  default:
    qtrt.ErrorResolve("QPainterPath", "united", args)
  }

}


func (this *QPainterPath) angleAtPercent(args ...interface{}) () {
  // angleAtPercent(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPainterPath14angleAtPercentEd
  default:
    qtrt.ErrorResolve("QPainterPath", "angleAtPercent", args)
  }

}


func (this *QPainterPath) closeSubpath(args ...interface{}) () {
  // closeSubpath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPainterPath12closeSubpathEv
  default:
    qtrt.ErrorResolve("QPainterPath", "closeSubpath", args)
  }

}


func (this *QPainterPathStroker) curveThreshold(args ...interface{}) () {
  // curveThreshold()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QPainterPathStroker14curveThresholdEv
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "curveThreshold", args)
  }

}


func (this *QPainterPathStroker) setWidth(args ...interface{}) () {
  // setWidth(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QPainterPathStroker8setWidthEd
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "setWidth", args)
  }

}


func (this *QPainterPathStroker) FreeQPainterPathStroker(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "~QPainterPathStroker", args)
  }

}


func (this *QPainterPathStroker) setMiterLimit(args ...interface{}) () {
  // setMiterLimit(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QPainterPathStroker13setMiterLimitEd
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "setMiterLimit", args)
  }

}


func NewQPainterPathStroker(args ...interface{}) QPainterPathStroker {
  return QPainterPathStroker{}
}


func (this *QPainterPathStroker) setCurveThreshold(args ...interface{}) () {
  // setCurveThreshold(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QPainterPathStroker17setCurveThresholdEd
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "setCurveThreshold", args)
  }

}


func (this *QPainterPathStroker) dashPattern(args ...interface{}) () {
  // dashPattern()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QPainterPathStroker11dashPatternEv
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "dashPattern", args)
  }

}


func (this *QPainterPathStroker) dashOffset(args ...interface{}) () {
  // dashOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QPainterPathStroker10dashOffsetEv
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "dashOffset", args)
  }

}


func (this *QPainterPathStroker) createStroke(args ...interface{}) () {
  // createStroke(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QPainterPathStroker12createStrokeERK12QPainterPath
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "createStroke", args)
  }

}


func (this *QPainterPathStroker) setDashOffset(args ...interface{}) () {
  // setDashOffset(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QPainterPathStroker13setDashOffsetEd
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "setDashOffset", args)
  }

}


func (this *QPainterPathStroker) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QPainterPathStroker5widthEv
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "width", args)
  }

}


func (this *QPainterPathStroker) miterLimit(args ...interface{}) () {
  // miterLimit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QPainterPathStroker10miterLimitEv
  default:
    qtrt.ErrorResolve("QPainterPathStroker", "miterLimit", args)
  }

}

// <= body block end

