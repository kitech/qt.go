package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qvector3d.h
// dst-file: /src/gui/qvector3d.go
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
// class sizeof(QVector3D)=12
type QVector3D struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QVector3D) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D1xEv
  default:
    qtrt.ErrorResolve("QVector3D", "x", args)
  }

}


func (this *QVector3D) toPoint(args ...interface{}) () {
  // toPoint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D7toPointEv
  default:
    qtrt.ErrorResolve("QVector3D", "toPoint", args)
  }

}


func (this *QVector3D) distanceToLine(args ...interface{}) () {
  // distanceToLine(const class QVector3D &, const class QVector3D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[0][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D14distanceToLineERKS_S1_
  default:
    qtrt.ErrorResolve("QVector3D", "distanceToLine", args)
  }

}


func (this *QVector3D) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D1yEv
  default:
    qtrt.ErrorResolve("QVector3D", "y", args)
  }

}


func (this *QVector3D) normal_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QVector3D", "normal", args)
  }

}


func (this *QVector3D) toPointF(args ...interface{}) () {
  // toPointF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D8toPointFEv
  default:
    qtrt.ErrorResolve("QVector3D", "toPointF", args)
  }

}


func (this *QVector3D) normalize(args ...interface{}) () {
  // normalize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector3D9normalizeEv
  default:
    qtrt.ErrorResolve("QVector3D", "normalize", args)
  }

}


func (this *QVector3D) unproject(args ...interface{}) () {
  // unproject(const class QMatrix4x4 &, const class QMatrix4x4 &, const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix4x4{}) // "const QMatrix4x4 &"
  vtys[0][1] = reflect.TypeOf(QMatrix4x4{}) // "const QMatrix4x4 &"
  vtys[0][2] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D9unprojectERK10QMatrix4x4S2_RK5QRect
  default:
    qtrt.ErrorResolve("QVector3D", "unproject", args)
  }

}


func (this *QVector3D) setY(args ...interface{}) () {
  // setY(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector3D4setYEf
  default:
    qtrt.ErrorResolve("QVector3D", "setY", args)
  }

}


func (this *QVector3D) project(args ...interface{}) () {
  // project(const class QMatrix4x4 &, const class QMatrix4x4 &, const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix4x4{}) // "const QMatrix4x4 &"
  vtys[0][1] = reflect.TypeOf(QMatrix4x4{}) // "const QMatrix4x4 &"
  vtys[0][2] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D7projectERK10QMatrix4x4S2_RK5QRect
  default:
    qtrt.ErrorResolve("QVector3D", "project", args)
  }

}


func (this *QVector3D) crossProduct_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QVector3D", "crossProduct", args)
  }

}


func (this *QVector3D) z(args ...interface{}) () {
  // z()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D1zEv
  default:
    qtrt.ErrorResolve("QVector3D", "z", args)
  }

}


func (this *QVector3D) setZ(args ...interface{}) () {
  // setZ(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector3D4setZEf
  default:
    qtrt.ErrorResolve("QVector3D", "setZ", args)
  }

}


func (this *QVector3D) toVector2D(args ...interface{}) () {
  // toVector2D()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D10toVector2DEv
  default:
    qtrt.ErrorResolve("QVector3D", "toVector2D", args)
  }

}


func (this *QVector3D) distanceToPlane(args ...interface{}) () {
  // distanceToPlane(const class QVector3D &, const class QVector3D &, const class QVector3D &)
  // distanceToPlane(const class QVector3D &, const class QVector3D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[0][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[0][2] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[1][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D15distanceToPlaneERKS_S1_S1_
  case 1:
    // invoke: _ZNK9QVector3D15distanceToPlaneERKS_S1_
  default:
    qtrt.ErrorResolve("QVector3D", "distanceToPlane", args)
  }

}


func NewQVector3D(args ...interface{}) QVector3D {
  return QVector3D{}
}


func (this *QVector3D) distanceToPoint(args ...interface{}) () {
  // distanceToPoint(const class QVector3D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D15distanceToPointERKS_
  default:
    qtrt.ErrorResolve("QVector3D", "distanceToPoint", args)
  }

}


func (this *QVector3D) lengthSquared(args ...interface{}) () {
  // lengthSquared()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D13lengthSquaredEv
  default:
    qtrt.ErrorResolve("QVector3D", "lengthSquared", args)
  }

}


func (this *QVector3D) normalized(args ...interface{}) () {
  // normalized()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D10normalizedEv
  default:
    qtrt.ErrorResolve("QVector3D", "normalized", args)
  }

}


func (this *QVector3D) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D6isNullEv
  default:
    qtrt.ErrorResolve("QVector3D", "isNull", args)
  }

}


func (this *QVector3D) length(args ...interface{}) () {
  // length()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D6lengthEv
  default:
    qtrt.ErrorResolve("QVector3D", "length", args)
  }

}


func (this *QVector3D) dotProduct_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QVector3D", "dotProduct", args)
  }

}


func (this *QVector3D) toVector4D(args ...interface{}) () {
  // toVector4D()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D10toVector4DEv
  default:
    qtrt.ErrorResolve("QVector3D", "toVector4D", args)
  }

}


func (this *QVector3D) setX(args ...interface{}) () {
  // setX(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector3D4setXEf
  default:
    qtrt.ErrorResolve("QVector3D", "setX", args)
  }

}

// <= body block end

