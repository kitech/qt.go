package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto: static QQuaternion QQuaternion::rotationTo(const QVector3D & from, const QVector3D & to);
extern void C_ZN11QQuaternion10rotationToERK9QVector3DS2_(void* arg0, void* arg1); // 4
  // proto: static QQuaternion QQuaternion::fromDirection(const QVector3D & direction, const QVector3D & up);
extern void C_ZN11QQuaternion13fromDirectionERK9QVector3DS2_(void* arg0, void* arg1); // 4
  // proto:  QVector4D QQuaternion::toVector4D();
extern void C_ZNK11QQuaternion10toVector4DEv(void* qthis); // 2
  // proto:  void QQuaternion::QQuaternion(float scalar, float xpos, float ypos, float zpos);
extern void* C_ZN11QQuaternionC2Effff(float arg0, float arg1, float arg2, float arg3); // 1
  // proto:  void QQuaternion::QQuaternion(const QVector4D & vector);
extern void* C_ZN11QQuaternionC2ERK9QVector4D(void* arg0); // 1
  // proto:  void QQuaternion::QQuaternion(float scalar, const QVector3D & vector);
extern void* C_ZN11QQuaternionC2EfRK9QVector3D(float arg0, void* arg1); // 1
  // proto:  void QQuaternion::QQuaternion();
extern void* C_ZN11QQuaternionC2Ev(); // 1
  // proto:  QVector3D QQuaternion::rotatedVector(const QVector3D & vector);
extern void C_ZNK11QQuaternion13rotatedVectorERK9QVector3D(void* qthis, void* arg0); // 4
  // proto:  QQuaternion QQuaternion::inverted();
extern void C_ZNK11QQuaternion8invertedEv(void* qthis); // 2
  // proto:  float QQuaternion::scalar();
extern void C_ZNK11QQuaternion6scalarEv(void* qthis); // 2
  // proto:  QQuaternion QQuaternion::conjugate();
extern void C_ZNK11QQuaternion9conjugateEv(void* qthis); // 2
  // proto: static float QQuaternion::dotProduct(const QQuaternion & q1, const QQuaternion & q2);
extern void C_ZN11QQuaternion10dotProductERKS_S1_(void* arg0, void* arg1); // 2
  // proto:  void QQuaternion::getAxes(QVector3D * xAxis, QVector3D * yAxis, QVector3D * zAxis);
extern void C_ZNK11QQuaternion7getAxesEP9QVector3DS1_S1_(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QQuaternion::normalize();
extern void C_ZN11QQuaternion9normalizeEv(void* qthis); // 4
  // proto: static QQuaternion QQuaternion::slerp(const QQuaternion & q1, const QQuaternion & q2, float t);
extern void C_ZN11QQuaternion5slerpERKS_S1_f(void* arg0, void* arg1, float arg2); // 4
  // proto:  void QQuaternion::setVector(const QVector3D & vector);
extern void C_ZN11QQuaternion9setVectorERK9QVector3D(void* qthis, void* arg0); // 2
  // proto:  void QQuaternion::setVector(float x, float y, float z);
extern void C_ZN11QQuaternion9setVectorEfff(void* qthis, float arg0, float arg1, float arg2); // 2
  // proto: static QQuaternion QQuaternion::fromAxisAndAngle(const QVector3D & axis, float angle);
extern void C_ZN11QQuaternion16fromAxisAndAngleERK9QVector3Df(void* arg0, float arg1); // 4
  // proto: static QQuaternion QQuaternion::fromAxisAndAngle(float x, float y, float z, float angle);
extern void C_ZN11QQuaternion16fromAxisAndAngleEffff(float arg0, float arg1, float arg2, float arg3); // 4
  // proto:  QQuaternion QQuaternion::conjugated();
extern void C_ZNK11QQuaternion10conjugatedEv(void* qthis); // 2
  // proto:  void QQuaternion::getAxisAndAngle(float * x, float * y, float * z, float * angle);
extern void C_ZNK11QQuaternion15getAxisAndAngleEPfS0_S0_S0_(void* qthis, float* arg0, float* arg1, float* arg2, float* arg3); // 4
  // proto:  void QQuaternion::getAxisAndAngle(QVector3D * axis, float * angle);
extern void C_ZNK11QQuaternion15getAxisAndAngleEP9QVector3DPf(void* qthis, void* arg0, float* arg1); // 2
  // proto: static QQuaternion QQuaternion::fromEulerAngles(const QVector3D & eulerAngles);
extern void C_ZN11QQuaternion15fromEulerAnglesERK9QVector3D(void* arg0); // 2
  // proto: static QQuaternion QQuaternion::fromEulerAngles(float pitch, float yaw, float roll);
extern void C_ZN11QQuaternion15fromEulerAnglesEfff(float arg0, float arg1, float arg2); // 4
  // proto:  float QQuaternion::lengthSquared();
extern void C_ZNK11QQuaternion13lengthSquaredEv(void* qthis); // 4
  // proto: static QQuaternion QQuaternion::nlerp(const QQuaternion & q1, const QQuaternion & q2, float t);
extern void C_ZN11QQuaternion5nlerpERKS_S1_f(void* arg0, void* arg1, float arg2); // 4
  // proto:  QQuaternion QQuaternion::normalized();
extern void C_ZNK11QQuaternion10normalizedEv(void* qthis); // 4
  // proto:  float QQuaternion::x();
extern void C_ZNK11QQuaternion1xEv(void* qthis); // 2
  // proto:  void QQuaternion::setX(float x);
extern void C_ZN11QQuaternion4setXEf(void* qthis, float arg0); // 2
  // proto:  void QQuaternion::setY(float y);
extern void C_ZN11QQuaternion4setYEf(void* qthis, float arg0); // 2
  // proto:  void QQuaternion::setZ(float z);
extern void C_ZN11QQuaternion4setZEf(void* qthis, float arg0); // 2
  // proto:  void QQuaternion::getEulerAngles(float * pitch, float * yaw, float * roll);
extern void C_ZNK11QQuaternion14getEulerAnglesEPfS0_S0_(void* qthis, float* arg0, float* arg1, float* arg2); // 4
  // proto: static QQuaternion QQuaternion::fromAxes(const QVector3D & xAxis, const QVector3D & yAxis, const QVector3D & zAxis);
extern void C_ZN11QQuaternion8fromAxesERK9QVector3DS2_S2_(void* arg0, void* arg1, void* arg2); // 4
  // proto:  QVector3D QQuaternion::toEulerAngles();
extern void C_ZNK11QQuaternion13toEulerAnglesEv(void* qthis); // 2
  // proto:  void QQuaternion::setScalar(float scalar);
extern void C_ZN11QQuaternion9setScalarEf(void* qthis, float arg0); // 2
  // proto:  bool QQuaternion::isNull();
extern void C_ZNK11QQuaternion6isNullEv(void* qthis); // 2
  // proto:  float QQuaternion::length();
extern void C_ZNK11QQuaternion6lengthEv(void* qthis); // 4
  // proto:  QVector3D QQuaternion::vector();
extern void C_ZNK11QQuaternion6vectorEv(void* qthis); // 2
  // proto:  bool QQuaternion::isIdentity();
extern void C_ZNK11QQuaternion10isIdentityEv(void* qthis); // 2
  // proto:  float QQuaternion::y();
extern void C_ZNK11QQuaternion1yEv(void* qthis); // 2
  // proto:  float QQuaternion::z();
extern void C_ZNK11QQuaternion1zEv(void* qthis); // 2
  // proto:  QMatrix3x3 QQuaternion::toRotationMatrix();
extern void C_ZNK11QQuaternion16toRotationMatrixEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// rotationTo(const class QVector3D &, const class QVector3D &)
func (this *QQuaternion) rotationTo_s(args ...interface{}) () {
  // rotationTo(const class QVector3D &, const class QVector3D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[0][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QQuaternion10rotationToERK9QVector3DS2_
    // invoke: QQuaternion rotationTo(const class QVector3D &, const class QVector3D &)
    var arg0 = args[0].(QVector3D).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector3D).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN11QQuaternion10rotationToERK9QVector3DS2_(arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QQuaternion", "rotationTo", args)
  }

}

// fromDirection(const class QVector3D &, const class QVector3D &)
func (this *QQuaternion) fromDirection_s(args ...interface{}) () {
  // fromDirection(const class QVector3D &, const class QVector3D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[0][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QQuaternion13fromDirectionERK9QVector3DS2_
    // invoke: QQuaternion fromDirection(const class QVector3D &, const class QVector3D &)
    var arg0 = args[0].(QVector3D).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector3D).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN11QQuaternion13fromDirectionERK9QVector3DS2_(arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QQuaternion", "fromDirection", args)
  }

}

// toVector4D()
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
    // invoke: QVector4D toVector4D()
    var ret = C.C_ZNK11QQuaternion10toVector4DEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QQuaternion", "toVector4D", args)
  }

}

// QQuaternion(float, float, float, float)
func NewQQuaternion(args ...interface{}) *QQuaternion {
  // QQuaternion(float, float, float, float)
  // QQuaternion(const class QVector4D &)
  // QQuaternion(float, const class QVector3D &)
  // QQuaternion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"
  vtys[0][1] = qtrt.FloatTy(false) // "float"
  vtys[0][2] = qtrt.FloatTy(false) // "float"
  vtys[0][3] = qtrt.FloatTy(false) // "float"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QVector4D{}) // "const QVector4D &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.FloatTy(false) // "float"
  vtys[2][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[3] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QQuaternionC1Effff
    // invoke: void QQuaternion(float, float, float, float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(args[3].(float32))
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QQuaternionC2Effff(arg0, arg1, arg2, arg3)
    return &QQuaternion{qclsinst:qthis}
  case 1:
    // invoke: _ZN11QQuaternionC1ERK9QVector4D
    // invoke: void QQuaternion(const class QVector4D &)
    var arg0 = args[0].(QVector4D).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QQuaternionC2ERK9QVector4D(arg0)
    return &QQuaternion{qclsinst:qthis}
  case 2:
    // invoke: _ZN11QQuaternionC1EfRK9QVector3D
    // invoke: void QQuaternion(float, const class QVector3D &)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector3D).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QQuaternionC2EfRK9QVector3D(arg0, arg1)
    return &QQuaternion{qclsinst:qthis}
  case 3:
    // invoke: _ZN11QQuaternionC1Ev
    // invoke: void QQuaternion()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QQuaternionC2Ev()
    return &QQuaternion{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QQuaternion", "QQuaternion", args)
  }

  return nil // QQuaternion{qclsinst:qthis}
}

// rotatedVector(const class QVector3D &)
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
    // invoke: QVector3D rotatedVector(const class QVector3D &)
    var arg0 = args[0].(QVector3D).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK11QQuaternion13rotatedVectorERK9QVector3D(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QQuaternion", "rotatedVector", args)
  }

}

// inverted()
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
    // invoke: QQuaternion inverted()
    var ret = C.C_ZNK11QQuaternion8invertedEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QQuaternion", "inverted", args)
  }

}

// scalar()
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
    // invoke: float scalar()
    var ret = C.C_ZNK11QQuaternion6scalarEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QQuaternion", "scalar", args)
  }

}

// conjugate()
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
    // invoke: QQuaternion conjugate()
    var ret = C.C_ZNK11QQuaternion9conjugateEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QQuaternion", "conjugate", args)
  }

}

// dotProduct(const class QQuaternion &, const class QQuaternion &)
func (this *QQuaternion) dotProduct_s(args ...interface{}) () {
  // dotProduct(const class QQuaternion &, const class QQuaternion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QQuaternion{}) // "const QQuaternion &"
  vtys[0][1] = reflect.TypeOf(QQuaternion{}) // "const QQuaternion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QQuaternion10dotProductERKS_S1_
    // invoke: float dotProduct(const class QQuaternion &, const class QQuaternion &)
    var arg0 = args[0].(QQuaternion).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QQuaternion).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN11QQuaternion10dotProductERKS_S1_(arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QQuaternion", "dotProduct", args)
  }

}

// getAxes(class QVector3D *, class QVector3D *, class QVector3D *)
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
    // invoke: void getAxes(class QVector3D *, class QVector3D *, class QVector3D *)
    var arg0 = args[0].(QVector3D).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector3D).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QVector3D).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZNK11QQuaternion7getAxesEP9QVector3DS1_S1_(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QQuaternion", "getAxes", args)
  }

}

// normalize()
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
    // invoke: void normalize()
    C.C_ZN11QQuaternion9normalizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QQuaternion", "normalize", args)
  }

}

// slerp(const class QQuaternion &, const class QQuaternion &, float)
func (this *QQuaternion) slerp_s(args ...interface{}) () {
  // slerp(const class QQuaternion &, const class QQuaternion &, float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QQuaternion{}) // "const QQuaternion &"
  vtys[0][1] = reflect.TypeOf(QQuaternion{}) // "const QQuaternion &"
  vtys[0][2] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QQuaternion5slerpERKS_S1_f
    // invoke: QQuaternion slerp(const class QQuaternion &, const class QQuaternion &, float)
    var arg0 = args[0].(QQuaternion).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QQuaternion).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    var ret = C.C_ZN11QQuaternion5slerpERKS_S1_f(arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QQuaternion", "slerp", args)
  }

}

// setVector(const class QVector3D &)
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
    // invoke: void setVector(const class QVector3D &)
    var arg0 = args[0].(QVector3D).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QQuaternion9setVectorERK9QVector3D(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN11QQuaternion9setVectorEfff
    // invoke: void setVector(float, float, float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    C.C_ZN11QQuaternion9setVectorEfff(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QQuaternion", "setVector", args)
  }

}

// fromAxisAndAngle(const class QVector3D &, float)
func (this *QQuaternion) fromAxisAndAngle_s(args ...interface{}) () {
  // fromAxisAndAngle(const class QVector3D &, float)
  // fromAxisAndAngle(float, float, float, float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[0][1] = qtrt.FloatTy(false) // "float"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.FloatTy(false) // "float"
  vtys[1][1] = qtrt.FloatTy(false) // "float"
  vtys[1][2] = qtrt.FloatTy(false) // "float"
  vtys[1][3] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QQuaternion16fromAxisAndAngleERK9QVector3Df
    // invoke: QQuaternion fromAxisAndAngle(const class QVector3D &, float)
    var arg0 = args[0].(QVector3D).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN11QQuaternion16fromAxisAndAngleERK9QVector3Df(arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN11QQuaternion16fromAxisAndAngleEffff
    // invoke: QQuaternion fromAxisAndAngle(float, float, float, float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(args[3].(float32))
    if false {fmt.Println(arg3)}
    var ret = C.C_ZN11QQuaternion16fromAxisAndAngleEffff(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QQuaternion", "fromAxisAndAngle", args)
  }

}

// conjugated()
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
    // invoke: QQuaternion conjugated()
    var ret = C.C_ZNK11QQuaternion10conjugatedEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QQuaternion", "conjugated", args)
  }

}

// getAxisAndAngle(float *, float *, float *, float *)
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
    // invoke: void getAxisAndAngle(float *, float *, float *, float *)
    var arg0 = (*C.float)(args[0].(*float32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.float)(args[1].(*float32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.float)(args[2].(*float32))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.float)(args[3].(*float32))
    if false {fmt.Println(arg3)}
    C.C_ZNK11QQuaternion15getAxisAndAngleEPfS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZNK11QQuaternion15getAxisAndAngleEP9QVector3DPf
    // invoke: void getAxisAndAngle(class QVector3D *, float *)
    var arg0 = args[0].(QVector3D).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.float)(args[1].(*float32))
    if false {fmt.Println(arg1)}
    C.C_ZNK11QQuaternion15getAxisAndAngleEP9QVector3DPf(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QQuaternion", "getAxisAndAngle", args)
  }

}

// fromEulerAngles(const class QVector3D &)
func (this *QQuaternion) fromEulerAngles_s(args ...interface{}) () {
  // fromEulerAngles(const class QVector3D &)
  // fromEulerAngles(float, float, float)
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
    // invoke: _ZN11QQuaternion15fromEulerAnglesERK9QVector3D
    // invoke: QQuaternion fromEulerAngles(const class QVector3D &)
    var arg0 = args[0].(QVector3D).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN11QQuaternion15fromEulerAnglesERK9QVector3D(arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN11QQuaternion15fromEulerAnglesEfff
    // invoke: QQuaternion fromEulerAngles(float, float, float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    var ret = C.C_ZN11QQuaternion15fromEulerAnglesEfff(arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QQuaternion", "fromEulerAngles", args)
  }

}

// lengthSquared()
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
    // invoke: float lengthSquared()
    var ret = C.C_ZNK11QQuaternion13lengthSquaredEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QQuaternion", "lengthSquared", args)
  }

}

// nlerp(const class QQuaternion &, const class QQuaternion &, float)
func (this *QQuaternion) nlerp_s(args ...interface{}) () {
  // nlerp(const class QQuaternion &, const class QQuaternion &, float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QQuaternion{}) // "const QQuaternion &"
  vtys[0][1] = reflect.TypeOf(QQuaternion{}) // "const QQuaternion &"
  vtys[0][2] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QQuaternion5nlerpERKS_S1_f
    // invoke: QQuaternion nlerp(const class QQuaternion &, const class QQuaternion &, float)
    var arg0 = args[0].(QQuaternion).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QQuaternion).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    var ret = C.C_ZN11QQuaternion5nlerpERKS_S1_f(arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QQuaternion", "nlerp", args)
  }

}

// normalized()
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
    // invoke: QQuaternion normalized()
    var ret = C.C_ZNK11QQuaternion10normalizedEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QQuaternion", "normalized", args)
  }

}

// x()
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
    // invoke: float x()
    C.C_ZNK11QQuaternion1xEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QQuaternion", "x", args)
  }

}

// setX(float)
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
    // invoke: void setX(float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QQuaternion4setXEf(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QQuaternion", "setX", args)
  }

}

// setY(float)
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
    // invoke: void setY(float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QQuaternion4setYEf(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QQuaternion", "setY", args)
  }

}

// setZ(float)
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
    // invoke: void setZ(float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QQuaternion4setZEf(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QQuaternion", "setZ", args)
  }

}

// getEulerAngles(float *, float *, float *)
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
    // invoke: void getEulerAngles(float *, float *, float *)
    var arg0 = (*C.float)(args[0].(*float32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.float)(args[1].(*float32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.float)(args[2].(*float32))
    if false {fmt.Println(arg2)}
    C.C_ZNK11QQuaternion14getEulerAnglesEPfS0_S0_(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QQuaternion", "getEulerAngles", args)
  }

}

// fromAxes(const class QVector3D &, const class QVector3D &, const class QVector3D &)
func (this *QQuaternion) fromAxes_s(args ...interface{}) () {
  // fromAxes(const class QVector3D &, const class QVector3D &, const class QVector3D &)
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
    // invoke: _ZN11QQuaternion8fromAxesERK9QVector3DS2_S2_
    // invoke: QQuaternion fromAxes(const class QVector3D &, const class QVector3D &, const class QVector3D &)
    var arg0 = args[0].(QVector3D).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector3D).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QVector3D).qclsinst
    if false {fmt.Println(arg2)}
    var ret = C.C_ZN11QQuaternion8fromAxesERK9QVector3DS2_S2_(arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QQuaternion", "fromAxes", args)
  }

}

// toEulerAngles()
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
    // invoke: QVector3D toEulerAngles()
    var ret = C.C_ZNK11QQuaternion13toEulerAnglesEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QQuaternion", "toEulerAngles", args)
  }

}

// setScalar(float)
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
    // invoke: void setScalar(float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QQuaternion9setScalarEf(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QQuaternion", "setScalar", args)
  }

}

// isNull()
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
    // invoke: bool isNull()
    var ret = C.C_ZNK11QQuaternion6isNullEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QQuaternion", "isNull", args)
  }

}

// length()
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
    // invoke: float length()
    var ret = C.C_ZNK11QQuaternion6lengthEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QQuaternion", "length", args)
  }

}

// vector()
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
    // invoke: QVector3D vector()
    var ret = C.C_ZNK11QQuaternion6vectorEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QQuaternion", "vector", args)
  }

}

// isIdentity()
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
    // invoke: bool isIdentity()
    var ret = C.C_ZNK11QQuaternion10isIdentityEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QQuaternion", "isIdentity", args)
  }

}

// y()
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
    // invoke: float y()
    C.C_ZNK11QQuaternion1yEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QQuaternion", "y", args)
  }

}

// z()
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
    // invoke: float z()
    var ret = C.C_ZNK11QQuaternion1zEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QQuaternion", "z", args)
  }

}

// toRotationMatrix()
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
    // invoke: QMatrix3x3 toRotationMatrix()
    C.C_ZNK11QQuaternion16toRotationMatrixEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QQuaternion", "toRotationMatrix", args)
  }

}

// <= body block end

