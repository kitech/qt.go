package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qpaintengine.h
// dst-file: /src/gui/qpaintengine.go
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
// class sizeof(QTextItem)=1
type QTextItem struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QPaintEngineState)=1
type QPaintEngineState struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QPaintEngine)=1
type QPaintEngine struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QTextItem) descent(args ...interface{}) () {
  // descent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextItem7descentEv
  default:
    qtrt.ErrorResolve("QTextItem", "descent", args)
  }

}


func (this *QTextItem) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextItem5widthEv
  default:
    qtrt.ErrorResolve("QTextItem", "width", args)
  }

}


func (this *QTextItem) font(args ...interface{}) () {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextItem4fontEv
  default:
    qtrt.ErrorResolve("QTextItem", "font", args)
  }

}


func (this *QTextItem) ascent(args ...interface{}) () {
  // ascent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextItem6ascentEv
  default:
    qtrt.ErrorResolve("QTextItem", "ascent", args)
  }

}


func (this *QTextItem) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextItem4textEv
  default:
    qtrt.ErrorResolve("QTextItem", "text", args)
  }

}


func (this *QPaintEngineState) opacity(args ...interface{}) () {
  // opacity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState7opacityEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "opacity", args)
  }

}


func (this *QPaintEngineState) matrix(args ...interface{}) () {
  // matrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState6matrixEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "matrix", args)
  }

}


func (this *QPaintEngineState) painter(args ...interface{}) () {
  // painter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState7painterEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "painter", args)
  }

}


func (this *QPaintEngineState) transform(args ...interface{}) () {
  // transform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState9transformEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "transform", args)
  }

}


func (this *QPaintEngineState) brushOrigin(args ...interface{}) () {
  // brushOrigin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState11brushOriginEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "brushOrigin", args)
  }

}


func (this *QPaintEngineState) penNeedsResolving(args ...interface{}) () {
  // penNeedsResolving()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState17penNeedsResolvingEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "penNeedsResolving", args)
  }

}


func (this *QPaintEngineState) isClipEnabled(args ...interface{}) () {
  // isClipEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState13isClipEnabledEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "isClipEnabled", args)
  }

}


func (this *QPaintEngineState) font(args ...interface{}) () {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState4fontEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "font", args)
  }

}


func (this *QPaintEngineState) brushNeedsResolving(args ...interface{}) () {
  // brushNeedsResolving()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState19brushNeedsResolvingEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "brushNeedsResolving", args)
  }

}


func (this *QPaintEngineState) clipRegion(args ...interface{}) () {
  // clipRegion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState10clipRegionEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "clipRegion", args)
  }

}


func (this *QPaintEngineState) clipPath(args ...interface{}) () {
  // clipPath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState8clipPathEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "clipPath", args)
  }

}


func (this *QPaintEngineState) brush(args ...interface{}) () {
  // brush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState5brushEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "brush", args)
  }

}


func (this *QPaintEngineState) pen(args ...interface{}) () {
  // pen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState3penEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "pen", args)
  }

}


func (this *QPaintEngineState) backgroundBrush(args ...interface{}) () {
  // backgroundBrush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPaintEngineState15backgroundBrushEv
  default:
    qtrt.ErrorResolve("QPaintEngineState", "backgroundBrush", args)
  }

}


func (this *QPaintEngine) drawEllipse(args ...interface{}) () {
  // drawEllipse(const class QRect &)
  // drawEllipse(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine11drawEllipseERK5QRect
  case 1:
    // invoke: _ZN12QPaintEngine11drawEllipseERK6QRectF
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawEllipse", args)
  }

}


func NewQPaintEngine(args ...interface{}) QPaintEngine {
  return QPaintEngine{}
}


func (this *QPaintEngine) isActive(args ...interface{}) () {
  // isActive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintEngine8isActiveEv
  default:
    qtrt.ErrorResolve("QPaintEngine", "isActive", args)
  }

}


func (this *QPaintEngine) drawPoints(args ...interface{}) () {
  // drawPoints(const class QPointF *, int)
  // drawPoints(const class QPoint *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine10drawPointsEPK7QPointFi
  case 1:
    // invoke: _ZN12QPaintEngine10drawPointsEPK6QPointi
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawPoints", args)
  }

}


func (this *QPaintEngine) coordinateOffset(args ...interface{}) () {
  // coordinateOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintEngine16coordinateOffsetEv
  default:
    qtrt.ErrorResolve("QPaintEngine", "coordinateOffset", args)
  }

}


func (this *QPaintEngine) setPaintDevice(args ...interface{}) () {
  // setPaintDevice(class QPaintDevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPaintDevice{}) // "QPaintDevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine14setPaintDeviceEP12QPaintDevice
  default:
    qtrt.ErrorResolve("QPaintEngine", "setPaintDevice", args)
  }

}


func (this *QPaintEngine) setSystemRect(args ...interface{}) () {
  // setSystemRect(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine13setSystemRectERK5QRect
  default:
    qtrt.ErrorResolve("QPaintEngine", "setSystemRect", args)
  }

}


func (this *QPaintEngine) FreeQPaintEngine(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPaintEngine", "~QPaintEngine", args)
  }

}


func (this *QPaintEngine) end(args ...interface{}) () {
  // end()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine3endEv
  default:
    qtrt.ErrorResolve("QPaintEngine", "end", args)
  }

}


func (this *QPaintEngine) drawTiledPixmap(args ...interface{}) () {
  // drawTiledPixmap(const class QRectF &, const class QPixmap &, const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[0][2] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine15drawTiledPixmapERK6QRectFRK7QPixmapRK7QPointF
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawTiledPixmap", args)
  }

}


func (this *QPaintEngine) setActive(args ...interface{}) () {
  // setActive(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine9setActiveEb
  default:
    qtrt.ErrorResolve("QPaintEngine", "setActive", args)
  }

}


func (this *QPaintEngine) drawPixmap(args ...interface{}) () {
  // drawPixmap(const class QRectF &, const class QPixmap &, const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[0][2] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine10drawPixmapERK6QRectFRK7QPixmapS2_
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawPixmap", args)
  }

}


func (this *QPaintEngine) drawLines(args ...interface{}) () {
  // drawLines(const class QLine *, int)
  // drawLines(const class QLineF *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLine{}) // "const QLine *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QLineF{}) // "const QLineF *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine9drawLinesEPK5QLinei
  case 1:
    // invoke: _ZN12QPaintEngine9drawLinesEPK6QLineFi
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawLines", args)
  }

}


func (this *QPaintEngine) drawPath(args ...interface{}) () {
  // drawPath(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine8drawPathERK12QPainterPath
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawPath", args)
  }

}


func (this *QPaintEngine) updateState(args ...interface{}) () {
  // updateState(const class QPaintEngineState &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPaintEngineState{}) // "const QPaintEngineState &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine11updateStateERK17QPaintEngineState
  default:
    qtrt.ErrorResolve("QPaintEngine", "updateState", args)
  }

}


func (this *QPaintEngine) begin(args ...interface{}) () {
  // begin(class QPaintDevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPaintDevice{}) // "QPaintDevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine5beginEP12QPaintDevice
  default:
    qtrt.ErrorResolve("QPaintEngine", "begin", args)
  }

}


func (this *QPaintEngine) systemRect(args ...interface{}) () {
  // systemRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintEngine10systemRectEv
  default:
    qtrt.ErrorResolve("QPaintEngine", "systemRect", args)
  }

}


func (this *QPaintEngine) drawRects(args ...interface{}) () {
  // drawRects(const class QRectF *, int)
  // drawRects(const class QRect *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine9drawRectsEPK6QRectFi
  case 1:
    // invoke: _ZN12QPaintEngine9drawRectsEPK5QRecti
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawRects", args)
  }

}


func (this *QPaintEngine) setSystemClip(args ...interface{}) () {
  // setSystemClip(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine13setSystemClipERK7QRegion
  default:
    qtrt.ErrorResolve("QPaintEngine", "setSystemClip", args)
  }

}


func (this *QPaintEngine) systemClip(args ...interface{}) () {
  // systemClip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintEngine10systemClipEv
  default:
    qtrt.ErrorResolve("QPaintEngine", "systemClip", args)
  }

}


func (this *QPaintEngine) paintDevice(args ...interface{}) () {
  // paintDevice()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintEngine11paintDeviceEv
  default:
    qtrt.ErrorResolve("QPaintEngine", "paintDevice", args)
  }

}


func (this *QPaintEngine) syncState(args ...interface{}) () {
  // syncState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine9syncStateEv
  default:
    qtrt.ErrorResolve("QPaintEngine", "syncState", args)
  }

}


func (this *QPaintEngine) painter(args ...interface{}) () {
  // painter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintEngine7painterEv
  default:
    qtrt.ErrorResolve("QPaintEngine", "painter", args)
  }

}


func (this *QPaintEngine) drawTextItem(args ...interface{}) () {
  // drawTextItem(const class QPointF &, const class QTextItem &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][1] = reflect.TypeOf(QTextItem{}) // "const QTextItem &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine12drawTextItemERK7QPointFRK9QTextItem
  default:
    qtrt.ErrorResolve("QPaintEngine", "drawTextItem", args)
  }

}


func (this *QPaintEngine) fix_neg_rect(args ...interface{}) () {
  // fix_neg_rect(int *, int *, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "int *"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintEngine12fix_neg_rectEPiS0_S0_S0_
  default:
    qtrt.ErrorResolve("QPaintEngine", "fix_neg_rect", args)
  }

}


func (this *QPaintEngine) isExtended(args ...interface{}) () {
  // isExtended()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintEngine10isExtendedEv
  default:
    qtrt.ErrorResolve("QPaintEngine", "isExtended", args)
  }

}

// <= body block end

