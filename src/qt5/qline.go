package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qline.h
// dst-file: /src/core/qline.go
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
// class sizeof(QLine)=16
type QLine struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QLineF)=32
type QLineF struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QLine) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QLine6isNullEv
  default:
    qtrt.ErrorResolve("QLine", "isNull", args)
  }

}


func (this *QLine) translated(args ...interface{}) () {
  // translated(const class QPoint &)
  // translated(int, int)
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
    // invoke: _ZNK5QLine10translatedERK6QPoint
  case 1:
    // invoke: _ZNK5QLine10translatedEii
  default:
    qtrt.ErrorResolve("QLine", "translated", args)
  }

}


func (this *QLine) setP2(args ...interface{}) () {
  // setP2(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QLine5setP2ERK6QPoint
  default:
    qtrt.ErrorResolve("QLine", "setP2", args)
  }

}


func (this *QLine) x2(args ...interface{}) () {
  // x2()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QLine2x2Ev
  default:
    qtrt.ErrorResolve("QLine", "x2", args)
  }

}


func NewQLine(args ...interface{}) QLine {
  return QLine{}
}


func (this *QLine) setP1(args ...interface{}) () {
  // setP1(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QLine5setP1ERK6QPoint
  default:
    qtrt.ErrorResolve("QLine", "setP1", args)
  }

}


func (this *QLine) translate(args ...interface{}) () {
  // translate(const class QPoint &)
  // translate(int, int)
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
    // invoke: _ZN5QLine9translateERK6QPoint
  case 1:
    // invoke: _ZN5QLine9translateEii
  default:
    qtrt.ErrorResolve("QLine", "translate", args)
  }

}


func (this *QLine) dx(args ...interface{}) () {
  // dx()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QLine2dxEv
  default:
    qtrt.ErrorResolve("QLine", "dx", args)
  }

}


func (this *QLine) y2(args ...interface{}) () {
  // y2()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QLine2y2Ev
  default:
    qtrt.ErrorResolve("QLine", "y2", args)
  }

}


func (this *QLine) dy(args ...interface{}) () {
  // dy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QLine2dyEv
  default:
    qtrt.ErrorResolve("QLine", "dy", args)
  }

}


func (this *QLine) y1(args ...interface{}) () {
  // y1()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QLine2y1Ev
  default:
    qtrt.ErrorResolve("QLine", "y1", args)
  }

}


func (this *QLine) p1(args ...interface{}) () {
  // p1()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QLine2p1Ev
  default:
    qtrt.ErrorResolve("QLine", "p1", args)
  }

}


func (this *QLine) p2(args ...interface{}) () {
  // p2()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QLine2p2Ev
  default:
    qtrt.ErrorResolve("QLine", "p2", args)
  }

}


func (this *QLine) setPoints(args ...interface{}) () {
  // setPoints(const class QPoint &, const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[0][1] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QLine9setPointsERK6QPointS2_
  default:
    qtrt.ErrorResolve("QLine", "setPoints", args)
  }

}


func (this *QLine) setLine(args ...interface{}) () {
  // setLine(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QLine7setLineEiiii
  default:
    qtrt.ErrorResolve("QLine", "setLine", args)
  }

}


func (this *QLine) x1(args ...interface{}) () {
  // x1()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QLine2x1Ev
  default:
    qtrt.ErrorResolve("QLine", "x1", args)
  }

}


func (this *QLineF) translate(args ...interface{}) () {
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
    // invoke: _ZN6QLineF9translateEdd
  case 1:
    // invoke: _ZN6QLineF9translateERK7QPointF
  default:
    qtrt.ErrorResolve("QLineF", "translate", args)
  }

}


func (this *QLineF) setPoints(args ...interface{}) () {
  // setPoints(const class QPointF &, const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLineF9setPointsERK7QPointFS2_
  default:
    qtrt.ErrorResolve("QLineF", "setPoints", args)
  }

}


func (this *QLineF) setP2(args ...interface{}) () {
  // setP2(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLineF5setP2ERK7QPointF
  default:
    qtrt.ErrorResolve("QLineF", "setP2", args)
  }

}


func (this *QLineF) translated(args ...interface{}) () {
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
    // invoke: _ZNK6QLineF10translatedEdd
  case 1:
    // invoke: _ZNK6QLineF10translatedERK7QPointF
  default:
    qtrt.ErrorResolve("QLineF", "translated", args)
  }

}


func (this *QLineF) setLength(args ...interface{}) () {
  // setLength(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLineF9setLengthEd
  default:
    qtrt.ErrorResolve("QLineF", "setLength", args)
  }

}


func (this *QLineF) x1(args ...interface{}) () {
  // x1()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF2x1Ev
  default:
    qtrt.ErrorResolve("QLineF", "x1", args)
  }

}


func (this *QLineF) angle(args ...interface{}) () {
  // angle()
  // angle(const class QLineF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QLineF{}) // "const QLineF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF5angleEv
  case 1:
    // invoke: _ZNK6QLineF5angleERKS_
  default:
    qtrt.ErrorResolve("QLineF", "angle", args)
  }

}


func NewQLineF(args ...interface{}) QLineF {
  return QLineF{}
}


func (this *QLineF) length(args ...interface{}) () {
  // length()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF6lengthEv
  default:
    qtrt.ErrorResolve("QLineF", "length", args)
  }

}


func (this *QLineF) setAngle(args ...interface{}) () {
  // setAngle(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLineF8setAngleEd
  default:
    qtrt.ErrorResolve("QLineF", "setAngle", args)
  }

}


func (this *QLineF) x2(args ...interface{}) () {
  // x2()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF2x2Ev
  default:
    qtrt.ErrorResolve("QLineF", "x2", args)
  }

}


func (this *QLineF) dx(args ...interface{}) () {
  // dx()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF2dxEv
  default:
    qtrt.ErrorResolve("QLineF", "dx", args)
  }

}


func (this *QLineF) p1(args ...interface{}) () {
  // p1()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF2p1Ev
  default:
    qtrt.ErrorResolve("QLineF", "p1", args)
  }

}


func (this *QLineF) normalVector(args ...interface{}) () {
  // normalVector()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF12normalVectorEv
  default:
    qtrt.ErrorResolve("QLineF", "normalVector", args)
  }

}


func (this *QLineF) toLine(args ...interface{}) () {
  // toLine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF6toLineEv
  default:
    qtrt.ErrorResolve("QLineF", "toLine", args)
  }

}


func (this *QLineF) pointAt(args ...interface{}) () {
  // pointAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF7pointAtEd
  default:
    qtrt.ErrorResolve("QLineF", "pointAt", args)
  }

}


func (this *QLineF) p2(args ...interface{}) () {
  // p2()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF2p2Ev
  default:
    qtrt.ErrorResolve("QLineF", "p2", args)
  }

}


func (this *QLineF) y2(args ...interface{}) () {
  // y2()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF2y2Ev
  default:
    qtrt.ErrorResolve("QLineF", "y2", args)
  }

}


func (this *QLineF) dy(args ...interface{}) () {
  // dy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF2dyEv
  default:
    qtrt.ErrorResolve("QLineF", "dy", args)
  }

}


func (this *QLineF) unitVector(args ...interface{}) () {
  // unitVector()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF10unitVectorEv
  default:
    qtrt.ErrorResolve("QLineF", "unitVector", args)
  }

}


func (this *QLineF) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF6isNullEv
  default:
    qtrt.ErrorResolve("QLineF", "isNull", args)
  }

}


func (this *QLineF) y1(args ...interface{}) () {
  // y1()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF2y1Ev
  default:
    qtrt.ErrorResolve("QLineF", "y1", args)
  }

}


func (this *QLineF) angleTo(args ...interface{}) () {
  // angleTo(const class QLineF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLineF{}) // "const QLineF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLineF7angleToERKS_
  default:
    qtrt.ErrorResolve("QLineF", "angleTo", args)
  }

}


func (this *QLineF) setLine(args ...interface{}) () {
  // setLine(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLineF7setLineEdddd
  default:
    qtrt.ErrorResolve("QLineF", "setLine", args)
  }

}


func (this *QLineF) fromPolar_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QLineF", "fromPolar", args)
  }

}


func (this *QLineF) setP1(args ...interface{}) () {
  // setP1(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLineF5setP1ERK7QPointF
  default:
    qtrt.ErrorResolve("QLineF", "setP1", args)
  }

}

// <= body block end

