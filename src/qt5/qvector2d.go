package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qvector2d.h
// dst-file: /src/gui/qvector2d.go
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
// class sizeof(QVector2D)=8
type QVector2D struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QVector2D) toPointF(args ...interface{}) () {
  // toPointF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D8toPointFEv
  default:
    qtrt.ErrorResolve("QVector2D", "toPointF", args)
  }

}


func (this *QVector2D) setX(args ...interface{}) () {
  // setX(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector2D4setXEf
  default:
    qtrt.ErrorResolve("QVector2D", "setX", args)
  }

}


func NewQVector2D(args ...interface{}) QVector2D {
  return QVector2D{}
}


func (this *QVector2D) toPoint(args ...interface{}) () {
  // toPoint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D7toPointEv
  default:
    qtrt.ErrorResolve("QVector2D", "toPoint", args)
  }

}


func (this *QVector2D) length(args ...interface{}) () {
  // length()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D6lengthEv
  default:
    qtrt.ErrorResolve("QVector2D", "length", args)
  }

}


func (this *QVector2D) setY(args ...interface{}) () {
  // setY(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector2D4setYEf
  default:
    qtrt.ErrorResolve("QVector2D", "setY", args)
  }

}


func (this *QVector2D) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D6isNullEv
  default:
    qtrt.ErrorResolve("QVector2D", "isNull", args)
  }

}


func (this *QVector2D) distanceToLine(args ...interface{}) () {
  // distanceToLine(const class QVector2D &, const class QVector2D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"
  vtys[0][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D14distanceToLineERKS_S1_
  default:
    qtrt.ErrorResolve("QVector2D", "distanceToLine", args)
  }

}


func (this *QVector2D) toVector3D(args ...interface{}) () {
  // toVector3D()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D10toVector3DEv
  default:
    qtrt.ErrorResolve("QVector2D", "toVector3D", args)
  }

}


func (this *QVector2D) lengthSquared(args ...interface{}) () {
  // lengthSquared()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D13lengthSquaredEv
  default:
    qtrt.ErrorResolve("QVector2D", "lengthSquared", args)
  }

}


func (this *QVector2D) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D1yEv
  default:
    qtrt.ErrorResolve("QVector2D", "y", args)
  }

}


func (this *QVector2D) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D1xEv
  default:
    qtrt.ErrorResolve("QVector2D", "x", args)
  }

}


func (this *QVector2D) distanceToPoint(args ...interface{}) () {
  // distanceToPoint(const class QVector2D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D15distanceToPointERKS_
  default:
    qtrt.ErrorResolve("QVector2D", "distanceToPoint", args)
  }

}


func (this *QVector2D) toVector4D(args ...interface{}) () {
  // toVector4D()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D10toVector4DEv
  default:
    qtrt.ErrorResolve("QVector2D", "toVector4D", args)
  }

}


func (this *QVector2D) normalized(args ...interface{}) () {
  // normalized()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D10normalizedEv
  default:
    qtrt.ErrorResolve("QVector2D", "normalized", args)
  }

}


func (this *QVector2D) normalize(args ...interface{}) () {
  // normalize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector2D9normalizeEv
  default:
    qtrt.ErrorResolve("QVector2D", "normalize", args)
  }

}


func (this *QVector2D) dotProduct_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QVector2D", "dotProduct", args)
  }

}

// <= body block end

