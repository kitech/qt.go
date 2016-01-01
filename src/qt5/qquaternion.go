package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qquaternion.h
// dst-file: /src/gui/qquaternion.go
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
// class sizeof(QQuaternion)=16
type QQuaternion struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QQuaternion) getAxisAndAngle(args ...interface{}) () {
  // getAxisAndAngle(float *, float *, float *, float *)
  // getAxisAndAngle(class QVector3D *, float *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(true) // "float *"
  vtys[0][1] = qtrt.FloatTy(true) // "float *"
  vtys[0][2] = qtrt.FloatTy(true) // "float *"
  vtys[0][3] = qtrt.FloatTy(true) // "float *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QVector3D{}) // "QVector3D *"
  vtys[1][1] = qtrt.FloatTy(true) // "float *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QQuaternion15getAxisAndAngleEPfS0_S0_S0_
  case 1:
    // invoke: _ZNK11QQuaternion15getAxisAndAngleEP9QVector3DPf
  default:
    qtrt.ErrorResolve("QQuaternion", "getAxisAndAngle", args)
 }

}


func (this *QQuaternion) scalar(args ...interface{}) () {
  // scalar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QQuaternion6scalarEv
  default:
    qtrt.ErrorResolve("QQuaternion", "scalar", args)
 }

}


func (this *QQuaternion) setX(args ...interface{}) () {
  // setX(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QQuaternion4setXEf
  default:
    qtrt.ErrorResolve("QQuaternion", "setX", args)
 }

}


func (this *QQuaternion) setVector(args ...interface{}) () {
  // setVector(const class QVector3D &)
  // setVector(float, float, float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.FloatTy(false) // "float"
  vtys[1][1] = qtrt.FloatTy(false) // "float"
  vtys[1][2] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QQuaternion9setVectorERK9QVector3D
  case 1:
    // invoke: _ZN11QQuaternion9setVectorEfff
  default:
    qtrt.ErrorResolve("QQuaternion", "setVector", args)
 }

}


func NewQQuaternion(args ...interface{})() {
}


func (this *QQuaternion) rotationTo_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QQuaternion", "rotationTo", args)
 }

}


func (this *QQuaternion) getEulerAngles(args ...interface{}) () {
  // getEulerAngles(float *, float *, float *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(true) // "float *"
  vtys[0][1] = qtrt.FloatTy(true) // "float *"
  vtys[0][2] = qtrt.FloatTy(true) // "float *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QQuaternion14getEulerAnglesEPfS0_S0_
  default:
    qtrt.ErrorResolve("QQuaternion", "getEulerAngles", args)
 }

}


func (this *QQuaternion) setY(args ...interface{}) () {
  // setY(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QQuaternion4setYEf
  default:
    qtrt.ErrorResolve("QQuaternion", "setY", args)
 }

}


func (this *QQuaternion) toEulerAngles(args ...interface{}) () {
  // toEulerAngles()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QQuaternion13toEulerAnglesEv
  default:
    qtrt.ErrorResolve("QQuaternion", "toEulerAngles", args)
 }

}


func (this *QQuaternion) inverted(args ...interface{}) () {
  // inverted()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QQuaternion8invertedEv
  default:
    qtrt.ErrorResolve("QQuaternion", "inverted", args)
 }

}


func (this *QQuaternion) setZ(args ...interface{}) () {
  // setZ(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QQuaternion4setZEf
  default:
    qtrt.ErrorResolve("QQuaternion", "setZ", args)
 }

}


func (this *QQuaternion) getAxes(args ...interface{}) () {
  // getAxes(class QVector3D *, class QVector3D *, class QVector3D *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector3D{}) // "QVector3D *"
  vtys[0][1] = reflect.TypeOf(QVector3D{}) // "QVector3D *"
  vtys[0][2] = reflect.TypeOf(QVector3D{}) // "QVector3D *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QQuaternion7getAxesEP9QVector3DS1_S1_
  default:
    qtrt.ErrorResolve("QQuaternion", "getAxes", args)
 }

}


func (this *QQuaternion) nlerp_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QQuaternion", "nlerp", args)
 }

}


func (this *QQuaternion) isIdentity(args ...interface{}) () {
  // isIdentity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QQuaternion10isIdentityEv
  default:
    qtrt.ErrorResolve("QQuaternion", "isIdentity", args)
 }

}


func (this *QQuaternion) fromAxes_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QQuaternion", "fromAxes", args)
 }

}


func (this *QQuaternion) slerp_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QQuaternion", "slerp", args)
 }

}


func (this *QQuaternion) fromDirection_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QQuaternion", "fromDirection", args)
 }

}


func (this *QQuaternion) normalized(args ...interface{}) () {
  // normalized()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QQuaternion10normalizedEv
  default:
    qtrt.ErrorResolve("QQuaternion", "normalized", args)
 }

}


func (this *QQuaternion) toVector4D(args ...interface{}) () {
  // toVector4D()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QQuaternion10toVector4DEv
  default:
    qtrt.ErrorResolve("QQuaternion", "toVector4D", args)
 }

}


func (this *QQuaternion) fromEulerAngles_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QQuaternion", "fromEulerAngles", args)
 }

}


func (this *QQuaternion) conjugate(args ...interface{}) () {
  // conjugate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QQuaternion9conjugateEv
  default:
    qtrt.ErrorResolve("QQuaternion", "conjugate", args)
 }

}


func (this *QQuaternion) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QQuaternion6isNullEv
  default:
    qtrt.ErrorResolve("QQuaternion", "isNull", args)
 }

}


func (this *QQuaternion) toRotationMatrix(args ...interface{}) () {
  // toRotationMatrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QQuaternion16toRotationMatrixEv
  default:
    qtrt.ErrorResolve("QQuaternion", "toRotationMatrix", args)
 }

}


func (this *QQuaternion) rotatedVector(args ...interface{}) () {
  // rotatedVector(const class QVector3D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QQuaternion13rotatedVectorERK9QVector3D
  default:
    qtrt.ErrorResolve("QQuaternion", "rotatedVector", args)
 }

}


func (this *QQuaternion) lengthSquared(args ...interface{}) () {
  // lengthSquared()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QQuaternion13lengthSquaredEv
  default:
    qtrt.ErrorResolve("QQuaternion", "lengthSquared", args)
 }

}


func (this *QQuaternion) setScalar(args ...interface{}) () {
  // setScalar(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QQuaternion9setScalarEf
  default:
    qtrt.ErrorResolve("QQuaternion", "setScalar", args)
 }

}


func (this *QQuaternion) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QQuaternion1yEv
  default:
    qtrt.ErrorResolve("QQuaternion", "y", args)
 }

}


func (this *QQuaternion) vector(args ...interface{}) () {
  // vector()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QQuaternion6vectorEv
  default:
    qtrt.ErrorResolve("QQuaternion", "vector", args)
 }

}


func (this *QQuaternion) dotProduct_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QQuaternion", "dotProduct", args)
 }

}


func (this *QQuaternion) fromAxisAndAngle_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QQuaternion", "fromAxisAndAngle", args)
 }

}


func (this *QQuaternion) length(args ...interface{}) () {
  // length()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QQuaternion6lengthEv
  default:
    qtrt.ErrorResolve("QQuaternion", "length", args)
 }

}


func (this *QQuaternion) normalize(args ...interface{}) () {
  // normalize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QQuaternion9normalizeEv
  default:
    qtrt.ErrorResolve("QQuaternion", "normalize", args)
 }

}


func (this *QQuaternion) conjugated(args ...interface{}) () {
  // conjugated()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QQuaternion10conjugatedEv
  default:
    qtrt.ErrorResolve("QQuaternion", "conjugated", args)
 }

}


func (this *QQuaternion) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QQuaternion1xEv
  default:
    qtrt.ErrorResolve("QQuaternion", "x", args)
 }

}


func (this *QQuaternion) z(args ...interface{}) () {
  // z()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QQuaternion1zEv
  default:
    qtrt.ErrorResolve("QQuaternion", "z", args)
 }

}

// <= body block end

