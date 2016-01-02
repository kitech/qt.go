package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtGui/qquaternion.h
// dst-file: /src/gui/qquaternion.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QQuaternion::getAxisAndAngle(float * x, float * y, float * z, float * angle);
extern void _ZNK11QQuaternion15getAxisAndAngleEPfS0_S0_S0_(void* qthis, float* arg0, float* arg1, float* arg2, float* arg3);
  // proto:  float QQuaternion::scalar();
extern void _ZNK11QQuaternion6scalarEv(void* qthis);
  // proto:  void QQuaternion::setX(float x);
extern void _ZN11QQuaternion4setXEf(void* qthis, float arg0);
  // proto:  void QQuaternion::setVector(const QVector3D & vector);
extern void _ZN11QQuaternion9setVectorERK9QVector3D(void* qthis, void* arg0);
  // proto:  void QQuaternion::QQuaternion(const QVector4D & vector);
extern void* dector_ZN11QQuaternionC1ERK9QVector4D(void* arg0);
extern void _ZN11QQuaternionC1ERK9QVector4D(void* qthis, void* arg0);
  // proto: static QQuaternion QQuaternion::rotationTo(const QVector3D & from, const QVector3D & to);
extern void _ZN11QQuaternion10rotationToERK9QVector3DS2_(void* arg0, void* arg1);
  // proto:  void QQuaternion::getEulerAngles(float * pitch, float * yaw, float * roll);
extern void _ZNK11QQuaternion14getEulerAnglesEPfS0_S0_(void* qthis, float* arg0, float* arg1, float* arg2);
  // proto:  void QQuaternion::setY(float y);
extern void _ZN11QQuaternion4setYEf(void* qthis, float arg0);
  // proto:  QVector3D QQuaternion::toEulerAngles();
extern void demth_ZNK11QQuaternion13toEulerAnglesEv(void* qthis);
  // proto:  QQuaternion QQuaternion::inverted();
extern void demth_ZNK11QQuaternion8invertedEv(void* qthis);
  // proto:  void QQuaternion::setZ(float z);
extern void _ZN11QQuaternion4setZEf(void* qthis, float arg0);
  // proto:  void QQuaternion::getAxes(QVector3D * xAxis, QVector3D * yAxis, QVector3D * zAxis);
extern void _ZNK11QQuaternion7getAxesEP9QVector3DS1_S1_(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto: static QQuaternion QQuaternion::nlerp(const QQuaternion & q1, const QQuaternion & q2, float t);
extern void _ZN11QQuaternion5nlerpERKS_S1_f(void* arg0, void* arg1, float arg2);
  // proto:  bool QQuaternion::isIdentity();
extern void _ZNK11QQuaternion10isIdentityEv(void* qthis);
  // proto: static QQuaternion QQuaternion::fromAxes(const QVector3D & xAxis, const QVector3D & yAxis, const QVector3D & zAxis);
extern void _ZN11QQuaternion8fromAxesERK9QVector3DS2_S2_(void* arg0, void* arg1, void* arg2);
  // proto: static QQuaternion QQuaternion::slerp(const QQuaternion & q1, const QQuaternion & q2, float t);
extern void _ZN11QQuaternion5slerpERKS_S1_f(void* arg0, void* arg1, float arg2);
  // proto: static QQuaternion QQuaternion::fromDirection(const QVector3D & direction, const QVector3D & up);
extern void _ZN11QQuaternion13fromDirectionERK9QVector3DS2_(void* arg0, void* arg1);
  // proto:  void QQuaternion::QQuaternion();
extern void* dector_ZN11QQuaternionC1Ev();
extern void _ZN11QQuaternionC1Ev(void* qthis);
  // proto:  QQuaternion QQuaternion::normalized();
extern void _ZNK11QQuaternion10normalizedEv(void* qthis);
  // proto:  QVector4D QQuaternion::toVector4D();
extern void _ZNK11QQuaternion10toVector4DEv(void* qthis);
  // proto: static QQuaternion QQuaternion::fromEulerAngles(float pitch, float yaw, float roll);
extern void _ZN11QQuaternion15fromEulerAnglesEfff(float arg0, float arg1, float arg2);
  // proto:  QQuaternion QQuaternion::conjugate();
extern void _ZNK11QQuaternion9conjugateEv(void* qthis);
  // proto:  bool QQuaternion::isNull();
extern void _ZNK11QQuaternion6isNullEv(void* qthis);
  // proto:  void QQuaternion::getAxisAndAngle(QVector3D * axis, float * angle);
extern void demth_ZNK11QQuaternion15getAxisAndAngleEP9QVector3DPf(void* qthis, void* arg0, float* arg1);
  // proto:  QMatrix3x3 QQuaternion::toRotationMatrix();
extern void _ZNK11QQuaternion16toRotationMatrixEv(void* qthis);
  // proto: static QQuaternion QQuaternion::fromEulerAngles(const QVector3D & eulerAngles);
extern void demth_ZN11QQuaternion15fromEulerAnglesERK9QVector3D(void* arg0);
  // proto:  QVector3D QQuaternion::rotatedVector(const QVector3D & vector);
extern void _ZNK11QQuaternion13rotatedVectorERK9QVector3D(void* qthis, void* arg0);
  // proto:  float QQuaternion::lengthSquared();
extern void _ZNK11QQuaternion13lengthSquaredEv(void* qthis);
  // proto:  void QQuaternion::setScalar(float scalar);
extern void _ZN11QQuaternion9setScalarEf(void* qthis, float arg0);
  // proto:  float QQuaternion::y();
extern void _ZNK11QQuaternion1yEv(void* qthis);
  // proto:  QVector3D QQuaternion::vector();
extern void _ZNK11QQuaternion6vectorEv(void* qthis);
  // proto: static float QQuaternion::dotProduct(const QQuaternion & q1, const QQuaternion & q2);
extern void _ZN11QQuaternion10dotProductERKS_S1_(void* arg0, void* arg1);
  // proto:  void QQuaternion::setVector(float x, float y, float z);
extern void _ZN11QQuaternion9setVectorEfff(void* qthis, float arg0, float arg1, float arg2);
  // proto:  void QQuaternion::QQuaternion(float scalar, float xpos, float ypos, float zpos);
extern void* dector_ZN11QQuaternionC1Effff(float arg0, float arg1, float arg2, float arg3);
extern void _ZN11QQuaternionC1Effff(void* qthis, float arg0, float arg1, float arg2, float arg3);
  // proto: static QQuaternion QQuaternion::fromAxisAndAngle(const QVector3D & axis, float angle);
extern void _ZN11QQuaternion16fromAxisAndAngleERK9QVector3Df(void* arg0, float arg1);
  // proto:  void QQuaternion::QQuaternion(float scalar, const QVector3D & vector);
extern void* dector_ZN11QQuaternionC1EfRK9QVector3D(float arg0, void* arg1);
extern void _ZN11QQuaternionC1EfRK9QVector3D(void* qthis, float arg0, void* arg1);
  // proto:  float QQuaternion::length();
extern void _ZNK11QQuaternion6lengthEv(void* qthis);
  // proto:  void QQuaternion::normalize();
extern void _ZN11QQuaternion9normalizeEv(void* qthis);
  // proto:  QQuaternion QQuaternion::conjugated();
extern void _ZNK11QQuaternion10conjugatedEv(void* qthis);
  // proto: static QQuaternion QQuaternion::fromAxisAndAngle(float x, float y, float z, float angle);
extern void _ZN11QQuaternion16fromAxisAndAngleEffff(float arg0, float arg1, float arg2, float arg3);
  // proto:  float QQuaternion::x();
extern void _ZNK11QQuaternion1xEv(void* qthis);
  // proto:  float QQuaternion::z();
extern void _ZNK11QQuaternion1zEv(void* qthis);
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QQuaternion)=16
type QQuaternion struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QQuaternion::getAxisAndAngle(float * x, float * y, float * z, float * angle);
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

  // proto:  float QQuaternion::scalar();
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

  // proto:  void QQuaternion::setX(float x);
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

  // proto:  void QQuaternion::setVector(const QVector3D & vector);
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

  // proto:  void QQuaternion::QQuaternion(const QVector4D & vector);
func NewQQuaternion(args ...interface{}) QQuaternion {
  return QQuaternion{}
}

  // proto: static QQuaternion QQuaternion::rotationTo(const QVector3D & from, const QVector3D & to);
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

  // proto:  void QQuaternion::getEulerAngles(float * pitch, float * yaw, float * roll);
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

  // proto:  void QQuaternion::setY(float y);
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

  // proto:  QVector3D QQuaternion::toEulerAngles();
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

  // proto:  QQuaternion QQuaternion::inverted();
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

  // proto:  void QQuaternion::setZ(float z);
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

  // proto:  void QQuaternion::getAxes(QVector3D * xAxis, QVector3D * yAxis, QVector3D * zAxis);
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

  // proto: static QQuaternion QQuaternion::nlerp(const QQuaternion & q1, const QQuaternion & q2, float t);
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

  // proto:  bool QQuaternion::isIdentity();
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

  // proto: static QQuaternion QQuaternion::fromAxes(const QVector3D & xAxis, const QVector3D & yAxis, const QVector3D & zAxis);
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

  // proto: static QQuaternion QQuaternion::slerp(const QQuaternion & q1, const QQuaternion & q2, float t);
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

  // proto: static QQuaternion QQuaternion::fromDirection(const QVector3D & direction, const QVector3D & up);
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

  // proto:  QQuaternion QQuaternion::normalized();
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

  // proto:  QVector4D QQuaternion::toVector4D();
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

  // proto: static QQuaternion QQuaternion::fromEulerAngles(float pitch, float yaw, float roll);
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

  // proto:  QQuaternion QQuaternion::conjugate();
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

  // proto:  bool QQuaternion::isNull();
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

  // proto:  QMatrix3x3 QQuaternion::toRotationMatrix();
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

  // proto:  QVector3D QQuaternion::rotatedVector(const QVector3D & vector);
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

  // proto:  float QQuaternion::lengthSquared();
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

  // proto:  void QQuaternion::setScalar(float scalar);
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

  // proto:  float QQuaternion::y();
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

  // proto:  QVector3D QQuaternion::vector();
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

  // proto: static float QQuaternion::dotProduct(const QQuaternion & q1, const QQuaternion & q2);
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

  // proto: static QQuaternion QQuaternion::fromAxisAndAngle(const QVector3D & axis, float angle);
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

  // proto:  float QQuaternion::length();
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

  // proto:  void QQuaternion::normalize();
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

  // proto:  QQuaternion QQuaternion::conjugated();
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

  // proto:  float QQuaternion::x();
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

  // proto:  float QQuaternion::z();
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

