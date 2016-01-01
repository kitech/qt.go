package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qmatrix4x4.h
// dst-file: /src/gui/qmatrix4x4.go
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
// class sizeof(QMatrix4x4)=68
type QMatrix4x4 struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QMatrix4x4) toTransform(args ...interface{}) () {
  // toTransform()
  // toTransform(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QMatrix4x411toTransformEv
  case 1:
    // invoke: _ZNK10QMatrix4x411toTransformEf
  default:
    qtrt.ErrorResolve("QMatrix4x4", "toTransform", args)
  }

}


func (this *QMatrix4x4) scale(args ...interface{}) () {
  // scale(const class QVector3D &)
  // scale(float, float, float)
  // scale(float, float)
  // scale(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.FloatTy(false) // "float"
  vtys[1][1] = qtrt.FloatTy(false) // "float"
  vtys[1][2] = qtrt.FloatTy(false) // "float"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.FloatTy(false) // "float"
  vtys[2][1] = qtrt.FloatTy(false) // "float"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QMatrix4x45scaleERK9QVector3D
  case 1:
    // invoke: _ZN10QMatrix4x45scaleEfff
  case 2:
    // invoke: _ZN10QMatrix4x45scaleEff
  case 3:
    // invoke: _ZN10QMatrix4x45scaleEf
  default:
    qtrt.ErrorResolve("QMatrix4x4", "scale", args)
  }

}


func (this *QMatrix4x4) translate(args ...interface{}) () {
  // translate(float, float, float)
  // translate(const class QVector3D &)
  // translate(float, float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"
  vtys[0][1] = qtrt.FloatTy(false) // "float"
  vtys[0][2] = qtrt.FloatTy(false) // "float"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.FloatTy(false) // "float"
  vtys[2][1] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QMatrix4x49translateEfff
  case 1:
    // invoke: _ZN10QMatrix4x49translateERK9QVector3D
  case 2:
    // invoke: _ZN10QMatrix4x49translateEff
  default:
    qtrt.ErrorResolve("QMatrix4x4", "translate", args)
  }

}


func (this *QMatrix4x4) constData(args ...interface{}) () {
  // constData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QMatrix4x49constDataEv
  default:
    qtrt.ErrorResolve("QMatrix4x4", "constData", args)
  }

}


func (this *QMatrix4x4) data(args ...interface{}) () {
  // data()
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QMatrix4x44dataEv
  case 1:
    // invoke: _ZNK10QMatrix4x44dataEv
  default:
    qtrt.ErrorResolve("QMatrix4x4", "data", args)
  }

}


func (this *QMatrix4x4) inverted(args ...interface{}) () {
  // inverted(_Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QMatrix4x48invertedEPb
  default:
    qtrt.ErrorResolve("QMatrix4x4", "inverted", args)
  }

}


func (this *QMatrix4x4) mapVector(args ...interface{}) () {
  // mapVector(const class QVector3D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QMatrix4x49mapVectorERK9QVector3D
  default:
    qtrt.ErrorResolve("QMatrix4x4", "mapVector", args)
  }

}


func (this *QMatrix4x4) ortho(args ...interface{}) () {
  // ortho(float, float, float, float, float, float)
  // ortho(const class QRect &)
  // ortho(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"
  vtys[0][1] = qtrt.FloatTy(false) // "float"
  vtys[0][2] = qtrt.FloatTy(false) // "float"
  vtys[0][3] = qtrt.FloatTy(false) // "float"
  vtys[0][4] = qtrt.FloatTy(false) // "float"
  vtys[0][5] = qtrt.FloatTy(false) // "float"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QMatrix4x45orthoEffffff
  case 1:
    // invoke: _ZN10QMatrix4x45orthoERK5QRect
  case 2:
    // invoke: _ZN10QMatrix4x45orthoERK6QRectF
  default:
    qtrt.ErrorResolve("QMatrix4x4", "ortho", args)
  }

}


func NewQMatrix4x4(args ...interface{}) QMatrix4x4 {
  return QMatrix4x4{}
}


func (this *QMatrix4x4) toAffine(args ...interface{}) () {
  // toAffine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QMatrix4x48toAffineEv
  default:
    qtrt.ErrorResolve("QMatrix4x4", "toAffine", args)
  }

}


func (this *QMatrix4x4) mapRect(args ...interface{}) () {
  // mapRect(const class QRectF &)
  // mapRect(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QMatrix4x47mapRectERK6QRectF
  case 1:
    // invoke: _ZNK10QMatrix4x47mapRectERK5QRect
  default:
    qtrt.ErrorResolve("QMatrix4x4", "mapRect", args)
  }

}


func (this *QMatrix4x4) setColumn(args ...interface{}) () {
  // setColumn(int, const class QVector4D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QVector4D{}) // "const QVector4D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QMatrix4x49setColumnEiRK9QVector4D
  default:
    qtrt.ErrorResolve("QMatrix4x4", "setColumn", args)
  }

}


func (this *QMatrix4x4) isIdentity(args ...interface{}) () {
  // isIdentity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QMatrix4x410isIdentityEv
  default:
    qtrt.ErrorResolve("QMatrix4x4", "isIdentity", args)
  }

}


func (this *QMatrix4x4) column(args ...interface{}) () {
  // column(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QMatrix4x46columnEi
  default:
    qtrt.ErrorResolve("QMatrix4x4", "column", args)
  }

}


func (this *QMatrix4x4) setRow(args ...interface{}) () {
  // setRow(int, const class QVector4D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QVector4D{}) // "const QVector4D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QMatrix4x46setRowEiRK9QVector4D
  default:
    qtrt.ErrorResolve("QMatrix4x4", "setRow", args)
  }

}


func (this *QMatrix4x4) flipCoordinates(args ...interface{}) () {
  // flipCoordinates()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QMatrix4x415flipCoordinatesEv
  default:
    qtrt.ErrorResolve("QMatrix4x4", "flipCoordinates", args)
  }

}


func (this *QMatrix4x4) normalMatrix(args ...interface{}) () {
  // normalMatrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QMatrix4x412normalMatrixEv
  default:
    qtrt.ErrorResolve("QMatrix4x4", "normalMatrix", args)
  }

}


func (this *QMatrix4x4) viewport(args ...interface{}) () {
  // viewport(float, float, float, float, float, float)
  // viewport(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"
  vtys[0][1] = qtrt.FloatTy(false) // "float"
  vtys[0][2] = qtrt.FloatTy(false) // "float"
  vtys[0][3] = qtrt.FloatTy(false) // "float"
  vtys[0][4] = qtrt.FloatTy(false) // "float"
  vtys[0][5] = qtrt.FloatTy(false) // "float"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QMatrix4x48viewportEffffff
  case 1:
    // invoke: _ZN10QMatrix4x48viewportERK6QRectF
  default:
    qtrt.ErrorResolve("QMatrix4x4", "viewport", args)
  }

}


func (this *QMatrix4x4) copyDataTo(args ...interface{}) () {
  // copyDataTo(float *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(true) // "float *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QMatrix4x410copyDataToEPf
  default:
    qtrt.ErrorResolve("QMatrix4x4", "copyDataTo", args)
  }

}


func (this *QMatrix4x4) isAffine(args ...interface{}) () {
  // isAffine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QMatrix4x48isAffineEv
  default:
    qtrt.ErrorResolve("QMatrix4x4", "isAffine", args)
  }

}


func (this *QMatrix4x4) rotate(args ...interface{}) () {
  // rotate(const class QQuaternion &)
  // rotate(float, float, float, float)
  // rotate(float, const class QVector3D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QQuaternion{}) // "const QQuaternion &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.FloatTy(false) // "float"
  vtys[1][1] = qtrt.FloatTy(false) // "float"
  vtys[1][2] = qtrt.FloatTy(false) // "float"
  vtys[1][3] = qtrt.FloatTy(false) // "float"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.FloatTy(false) // "float"
  vtys[2][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QMatrix4x46rotateERK11QQuaternion
  case 1:
    // invoke: _ZN10QMatrix4x46rotateEffff
  case 2:
    // invoke: _ZN10QMatrix4x46rotateEfRK9QVector3D
  default:
    qtrt.ErrorResolve("QMatrix4x4", "rotate", args)
  }

}


func (this *QMatrix4x4) perspective(args ...interface{}) () {
  // perspective(float, float, float, float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"
  vtys[0][1] = qtrt.FloatTy(false) // "float"
  vtys[0][2] = qtrt.FloatTy(false) // "float"
  vtys[0][3] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QMatrix4x411perspectiveEffff
  default:
    qtrt.ErrorResolve("QMatrix4x4", "perspective", args)
  }

}


func (this *QMatrix4x4) determinant(args ...interface{}) () {
  // determinant()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QMatrix4x411determinantEv
  default:
    qtrt.ErrorResolve("QMatrix4x4", "determinant", args)
  }

}


func (this *QMatrix4x4) frustum(args ...interface{}) () {
  // frustum(float, float, float, float, float, float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"
  vtys[0][1] = qtrt.FloatTy(false) // "float"
  vtys[0][2] = qtrt.FloatTy(false) // "float"
  vtys[0][3] = qtrt.FloatTy(false) // "float"
  vtys[0][4] = qtrt.FloatTy(false) // "float"
  vtys[0][5] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QMatrix4x47frustumEffffff
  default:
    qtrt.ErrorResolve("QMatrix4x4", "frustum", args)
  }

}


func (this *QMatrix4x4) map_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMatrix4x4", "map", args)
  }

}


func (this *QMatrix4x4) optimize(args ...interface{}) () {
  // optimize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QMatrix4x48optimizeEv
  default:
    qtrt.ErrorResolve("QMatrix4x4", "optimize", args)
  }

}


func (this *QMatrix4x4) setToIdentity(args ...interface{}) () {
  // setToIdentity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QMatrix4x413setToIdentityEv
  default:
    qtrt.ErrorResolve("QMatrix4x4", "setToIdentity", args)
  }

}


func (this *QMatrix4x4) lookAt(args ...interface{}) () {
  // lookAt(const class QVector3D &, const class QVector3D &, const class QVector3D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[0][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[0][2] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QMatrix4x46lookAtERK9QVector3DS2_S2_
  default:
    qtrt.ErrorResolve("QMatrix4x4", "lookAt", args)
  }

}


func (this *QMatrix4x4) fill(args ...interface{}) () {
  // fill(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QMatrix4x44fillEf
  default:
    qtrt.ErrorResolve("QMatrix4x4", "fill", args)
  }

}


func (this *QMatrix4x4) transposed(args ...interface{}) () {
  // transposed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QMatrix4x410transposedEv
  default:
    qtrt.ErrorResolve("QMatrix4x4", "transposed", args)
  }

}


func (this *QMatrix4x4) row(args ...interface{}) () {
  // row(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QMatrix4x43rowEi
  default:
    qtrt.ErrorResolve("QMatrix4x4", "row", args)
  }

}

// <= body block end

