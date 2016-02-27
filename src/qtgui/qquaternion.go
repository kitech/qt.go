package qtgui
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
import "runtime"
import "qtrt"
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto: static QQuaternion QQuaternion::rotationTo(const QVector3D & from, const QVector3D & to);
extern void* C_ZN11QQuaternion10rotationToERK9QVector3DS2_(void* arg0, void* arg1); // 4
  // proto: static QQuaternion QQuaternion::fromDirection(const QVector3D & direction, const QVector3D & up);
extern void* C_ZN11QQuaternion13fromDirectionERK9QVector3DS2_(void* arg0, void* arg1); // 4
  // proto:  QVector4D QQuaternion::toVector4D();
extern void* C_ZNK11QQuaternion10toVector4DEv(void* qthis); // 2
  // proto:  void QQuaternion::QQuaternion(float scalar, float xpos, float ypos, float zpos);
extern void* C_ZN11QQuaternionC2Effff(float arg0, float arg1, float arg2, float arg3); // 1
  // proto:  void QQuaternion::QQuaternion(const QVector4D & vector);
extern void* C_ZN11QQuaternionC2ERK9QVector4D(void* arg0); // 1
  // proto:  void QQuaternion::QQuaternion(float scalar, const QVector3D & vector);
extern void* C_ZN11QQuaternionC2EfRK9QVector3D(float arg0, void* arg1); // 1
  // proto:  void QQuaternion::QQuaternion();
extern void* C_ZN11QQuaternionC2Ev(); // 1
  // proto:  QVector3D QQuaternion::rotatedVector(const QVector3D & vector);
extern void* C_ZNK11QQuaternion13rotatedVectorERK9QVector3D(void* qthis, void* arg0); // 4
  // proto:  QQuaternion QQuaternion::inverted();
extern void* C_ZNK11QQuaternion8invertedEv(void* qthis); // 2
  // proto:  float QQuaternion::scalar();
extern float C_ZNK11QQuaternion6scalarEv(void* qthis); // 2
  // proto:  QQuaternion QQuaternion::conjugate();
extern void* C_ZNK11QQuaternion9conjugateEv(void* qthis); // 2
  // proto: static float QQuaternion::dotProduct(const QQuaternion & q1, const QQuaternion & q2);
extern float C_ZN11QQuaternion10dotProductERKS_S1_(void* arg0, void* arg1); // 2
  // proto:  void QQuaternion::getAxes(QVector3D * xAxis, QVector3D * yAxis, QVector3D * zAxis);
extern void C_ZNK11QQuaternion7getAxesEP9QVector3DS1_S1_(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QQuaternion::normalize();
extern void C_ZN11QQuaternion9normalizeEv(void* qthis); // 4
  // proto: static QQuaternion QQuaternion::slerp(const QQuaternion & q1, const QQuaternion & q2, float t);
extern void* C_ZN11QQuaternion5slerpERKS_S1_f(void* arg0, void* arg1, float arg2); // 4
  // proto:  void QQuaternion::setVector(const QVector3D & vector);
extern void C_ZN11QQuaternion9setVectorERK9QVector3D(void* qthis, void* arg0); // 2
  // proto:  void QQuaternion::setVector(float x, float y, float z);
extern void C_ZN11QQuaternion9setVectorEfff(void* qthis, float arg0, float arg1, float arg2); // 2
  // proto: static QQuaternion QQuaternion::fromAxisAndAngle(const QVector3D & axis, float angle);
extern void* C_ZN11QQuaternion16fromAxisAndAngleERK9QVector3Df(void* arg0, float arg1); // 4
  // proto: static QQuaternion QQuaternion::fromAxisAndAngle(float x, float y, float z, float angle);
extern void* C_ZN11QQuaternion16fromAxisAndAngleEffff(float arg0, float arg1, float arg2, float arg3); // 4
  // proto:  QQuaternion QQuaternion::conjugated();
extern void* C_ZNK11QQuaternion10conjugatedEv(void* qthis); // 2
  // proto:  void QQuaternion::getAxisAndAngle(float * x, float * y, float * z, float * angle);
extern void C_ZNK11QQuaternion15getAxisAndAngleEPfS0_S0_S0_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3); // 4
  // proto:  void QQuaternion::getAxisAndAngle(QVector3D * axis, float * angle);
extern void C_ZNK11QQuaternion15getAxisAndAngleEP9QVector3DPf(void* qthis, void* arg0, void* arg1); // 2
  // proto: static QQuaternion QQuaternion::fromEulerAngles(const QVector3D & eulerAngles);
extern void* C_ZN11QQuaternion15fromEulerAnglesERK9QVector3D(void* arg0); // 2
  // proto: static QQuaternion QQuaternion::fromEulerAngles(float pitch, float yaw, float roll);
extern void* C_ZN11QQuaternion15fromEulerAnglesEfff(float arg0, float arg1, float arg2); // 4
  // proto:  float QQuaternion::lengthSquared();
extern float C_ZNK11QQuaternion13lengthSquaredEv(void* qthis); // 4
  // proto: static QQuaternion QQuaternion::nlerp(const QQuaternion & q1, const QQuaternion & q2, float t);
extern void* C_ZN11QQuaternion5nlerpERKS_S1_f(void* arg0, void* arg1, float arg2); // 4
  // proto:  QQuaternion QQuaternion::normalized();
extern void* C_ZNK11QQuaternion10normalizedEv(void* qthis); // 4
  // proto:  float QQuaternion::x();
extern void C_ZNK11QQuaternion1xEv(void* qthis); // 2
  // proto:  void QQuaternion::setX(float x);
extern void C_ZN11QQuaternion4setXEf(void* qthis, float arg0); // 2
  // proto:  void QQuaternion::setY(float y);
extern void C_ZN11QQuaternion4setYEf(void* qthis, float arg0); // 2
  // proto:  void QQuaternion::setZ(float z);
extern void C_ZN11QQuaternion4setZEf(void* qthis, float arg0); // 2
  // proto:  void QQuaternion::getEulerAngles(float * pitch, float * yaw, float * roll);
extern void C_ZNK11QQuaternion14getEulerAnglesEPfS0_S0_(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto: static QQuaternion QQuaternion::fromAxes(const QVector3D & xAxis, const QVector3D & yAxis, const QVector3D & zAxis);
extern void* C_ZN11QQuaternion8fromAxesERK9QVector3DS2_S2_(void* arg0, void* arg1, void* arg2); // 4
  // proto:  QVector3D QQuaternion::toEulerAngles();
extern void* C_ZNK11QQuaternion13toEulerAnglesEv(void* qthis); // 2
  // proto:  void QQuaternion::setScalar(float scalar);
extern void C_ZN11QQuaternion9setScalarEf(void* qthis, float arg0); // 2
  // proto:  bool QQuaternion::isNull();
extern bool C_ZNK11QQuaternion6isNullEv(void* qthis); // 2
  // proto:  float QQuaternion::length();
extern float C_ZNK11QQuaternion6lengthEv(void* qthis); // 4
  // proto:  QVector3D QQuaternion::vector();
extern void* C_ZNK11QQuaternion6vectorEv(void* qthis); // 2
  // proto:  bool QQuaternion::isIdentity();
extern bool C_ZNK11QQuaternion10isIdentityEv(void* qthis); // 2
  // proto:  float QQuaternion::y();
extern void C_ZNK11QQuaternion1yEv(void* qthis); // 2
  // proto:  float QQuaternion::z();
extern float C_ZNK11QQuaternion1zEv(void* qthis); // 2
  // proto:  QMatrix3x3 QQuaternion::toRotationMatrix();
extern void C_ZNK11QQuaternion16toRotationMatrixEv(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QQuaternion)=16
type QQuaternion struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// rotationTo(const class QVector3D &, const class QVector3D &)
func (this *QQuaternion) RotationTo_s(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QVector3D).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVector3D).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN11QQuaternion10rotationToERK9QVector3DS2_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QQuaternion{}) // "QQuaternion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QQuaternion", "rotationTo", args)
  }

  return
}

// fromDirection(const class QVector3D &, const class QVector3D &)
func (this *QQuaternion) FromDirection_s(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QVector3D).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVector3D).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN11QQuaternion13fromDirectionERK9QVector3DS2_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QQuaternion{}) // "QQuaternion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QQuaternion", "fromDirection", args)
  }

  return
}

// toVector4D()
func (this *QQuaternion) ToVector4D(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QQuaternion10toVector4DEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVector4D{}) // "QVector4D"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QQuaternion", "toVector4D", args)
  }

  return
}

// QQuaternion(float, float, float, float)
func GcfreeQQuaternion(this *QQuaternion) {
  qtrt.UniverseFree(this)
}
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
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(qtrt.PrimConv(args[2], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(qtrt.PrimConv(args[3], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QQuaternionC2Effff(arg0, arg1, arg2, arg3)
    this := &QQuaternion{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQQuaternion)
    return this
  case 1:
    // invoke: _ZN11QQuaternionC1ERK9QVector4D
    // invoke: void QQuaternion(const class QVector4D &)
    var arg0 = args[0].(*QVector4D).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QQuaternionC2ERK9QVector4D(arg0)
    this := &QQuaternion{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQQuaternion)
    return this
  case 2:
    // invoke: _ZN11QQuaternionC1EfRK9QVector3D
    // invoke: void QQuaternion(float, const class QVector3D &)
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVector3D).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QQuaternionC2EfRK9QVector3D(arg0, arg1)
    this := &QQuaternion{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQQuaternion)
    return this
  case 3:
    // invoke: _ZN11QQuaternionC1Ev
    // invoke: void QQuaternion()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QQuaternionC2Ev()
    this := &QQuaternion{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQQuaternion)
    return this
  default:
    qtrt.ErrorResolve("QQuaternion", "QQuaternion", args)
  }

  return nil // QQuaternion{Qclsinst:qthis}
}

// rotatedVector(const class QVector3D &)
func (this *QQuaternion) RotatedVector(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QVector3D).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QQuaternion13rotatedVectorERK9QVector3D(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVector3D{}) // "QVector3D"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QQuaternion", "rotatedVector", args)
  }

  return
}

// inverted()
func (this *QQuaternion) Inverted(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QQuaternion8invertedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QQuaternion{}) // "QQuaternion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QQuaternion", "inverted", args)
  }

  return
}

// scalar()
func (this *QQuaternion) Scalar(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QQuaternion6scalarEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QQuaternion", "scalar", args)
  }

  return
}

// conjugate()
func (this *QQuaternion) Conjugate(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QQuaternion9conjugateEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QQuaternion{}) // "QQuaternion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QQuaternion", "conjugate", args)
  }

  return
}

// dotProduct(const class QQuaternion &, const class QQuaternion &)
func (this *QQuaternion) DotProduct_s(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QQuaternion).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QQuaternion).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN11QQuaternion10dotProductERKS_S1_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QQuaternion", "dotProduct", args)
  }

  return
}

// getAxes(class QVector3D *, class QVector3D *, class QVector3D *)
func (this *QQuaternion) GetAxes(args ...interface{}) () {
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
    var arg0 = args[0].(*QVector3D).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVector3D).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QVector3D).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZNK11QQuaternion7getAxesEP9QVector3DS1_S1_(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QQuaternion", "getAxes", args)
  }

  return
}

// normalize()
func (this *QQuaternion) Normalize(args ...interface{}) () {
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
    C.C_ZN11QQuaternion9normalizeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QQuaternion", "normalize", args)
  }

  return
}

// slerp(const class QQuaternion &, const class QQuaternion &, float)
func (this *QQuaternion) Slerp_s(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QQuaternion).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QQuaternion).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.float(qtrt.PrimConv(args[2], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN11QQuaternion5slerpERKS_S1_f(arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QQuaternion{}) // "QQuaternion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QQuaternion", "slerp", args)
  }

  return
}

// setVector(const class QVector3D &)
func (this *QQuaternion) SetVector(args ...interface{}) () {
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
    var arg0 = args[0].(*QVector3D).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QQuaternion9setVectorERK9QVector3D(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN11QQuaternion9setVectorEfff
    // invoke: void setVector(float, float, float)
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(qtrt.PrimConv(args[2], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg2)}
    C.C_ZN11QQuaternion9setVectorEfff(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QQuaternion", "setVector", args)
  }

  return
}

// fromAxisAndAngle(const class QVector3D &, float)
func (this *QQuaternion) FromAxisAndAngle_s(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QVector3D).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN11QQuaternion16fromAxisAndAngleERK9QVector3Df(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QQuaternion{}) // "QQuaternion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN11QQuaternion16fromAxisAndAngleEffff
    // invoke: QQuaternion fromAxisAndAngle(float, float, float, float)
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(qtrt.PrimConv(args[2], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(qtrt.PrimConv(args[3], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZN11QQuaternion16fromAxisAndAngleEffff(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QQuaternion{}) // "QQuaternion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QQuaternion", "fromAxisAndAngle", args)
  }

  return
}

// conjugated()
func (this *QQuaternion) Conjugated(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QQuaternion10conjugatedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QQuaternion{}) // "QQuaternion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QQuaternion", "conjugated", args)
  }

  return
}

// getAxisAndAngle(float *, float *, float *, float *)
func (this *QQuaternion) GetAxisAndAngle(args ...interface{}) () {
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
    var arg0 = (unsafe.Pointer)(args[0].(*float32))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*float32))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*float32))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*float32))
    if false {fmt.Println(arg3)}
    C.C_ZNK11QQuaternion15getAxisAndAngleEPfS0_S0_S0_(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZNK11QQuaternion15getAxisAndAngleEP9QVector3DPf
    // invoke: void getAxisAndAngle(class QVector3D *, float *)
    var arg0 = args[0].(*QVector3D).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*float32))
    if false {fmt.Println(arg1)}
    C.C_ZNK11QQuaternion15getAxisAndAngleEP9QVector3DPf(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QQuaternion", "getAxisAndAngle", args)
  }

  return
}

// fromEulerAngles(const class QVector3D &)
func (this *QQuaternion) FromEulerAngles_s(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QVector3D).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN11QQuaternion15fromEulerAnglesERK9QVector3D(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QQuaternion{}) // "QQuaternion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN11QQuaternion15fromEulerAnglesEfff
    // invoke: QQuaternion fromEulerAngles(float, float, float)
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(qtrt.PrimConv(args[2], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN11QQuaternion15fromEulerAnglesEfff(arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QQuaternion{}) // "QQuaternion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QQuaternion", "fromEulerAngles", args)
  }

  return
}

// lengthSquared()
func (this *QQuaternion) LengthSquared(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QQuaternion13lengthSquaredEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QQuaternion", "lengthSquared", args)
  }

  return
}

// nlerp(const class QQuaternion &, const class QQuaternion &, float)
func (this *QQuaternion) Nlerp_s(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QQuaternion).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QQuaternion).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.float(qtrt.PrimConv(args[2], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN11QQuaternion5nlerpERKS_S1_f(arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QQuaternion{}) // "QQuaternion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QQuaternion", "nlerp", args)
  }

  return
}

// normalized()
func (this *QQuaternion) Normalized(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QQuaternion10normalizedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QQuaternion{}) // "QQuaternion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QQuaternion", "normalized", args)
  }

  return
}

// x()
func (this *QQuaternion) X(args ...interface{}) () {
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
    C.C_ZNK11QQuaternion1xEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QQuaternion", "x", args)
  }

  return
}

// setX(float)
func (this *QQuaternion) SetX(args ...interface{}) () {
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
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QQuaternion4setXEf(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QQuaternion", "setX", args)
  }

  return
}

// setY(float)
func (this *QQuaternion) SetY(args ...interface{}) () {
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
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QQuaternion4setYEf(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QQuaternion", "setY", args)
  }

  return
}

// setZ(float)
func (this *QQuaternion) SetZ(args ...interface{}) () {
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
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QQuaternion4setZEf(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QQuaternion", "setZ", args)
  }

  return
}

// getEulerAngles(float *, float *, float *)
func (this *QQuaternion) GetEulerAngles(args ...interface{}) () {
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
    var arg0 = (unsafe.Pointer)(args[0].(*float32))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*float32))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*float32))
    if false {fmt.Println(arg2)}
    C.C_ZNK11QQuaternion14getEulerAnglesEPfS0_S0_(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QQuaternion", "getEulerAngles", args)
  }

  return
}

// fromAxes(const class QVector3D &, const class QVector3D &, const class QVector3D &)
func (this *QQuaternion) FromAxes_s(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QVector3D).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVector3D).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QVector3D).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN11QQuaternion8fromAxesERK9QVector3DS2_S2_(arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QQuaternion{}) // "QQuaternion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QQuaternion", "fromAxes", args)
  }

  return
}

// toEulerAngles()
func (this *QQuaternion) ToEulerAngles(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QQuaternion13toEulerAnglesEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVector3D{}) // "QVector3D"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QQuaternion", "toEulerAngles", args)
  }

  return
}

// setScalar(float)
func (this *QQuaternion) SetScalar(args ...interface{}) () {
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
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QQuaternion9setScalarEf(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QQuaternion", "setScalar", args)
  }

  return
}

// isNull()
func (this *QQuaternion) IsNull(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QQuaternion6isNullEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QQuaternion", "isNull", args)
  }

  return
}

// length()
func (this *QQuaternion) Length(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QQuaternion6lengthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QQuaternion", "length", args)
  }

  return
}

// vector()
func (this *QQuaternion) Vector(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QQuaternion6vectorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVector3D{}) // "QVector3D"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QQuaternion", "vector", args)
  }

  return
}

// isIdentity()
func (this *QQuaternion) IsIdentity(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QQuaternion10isIdentityEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QQuaternion", "isIdentity", args)
  }

  return
}

// y()
func (this *QQuaternion) Y(args ...interface{}) () {
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
    C.C_ZNK11QQuaternion1yEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QQuaternion", "y", args)
  }

  return
}

// z()
func (this *QQuaternion) Z(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QQuaternion1zEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QQuaternion", "z", args)
  }

  return
}

// toRotationMatrix()
func (this *QQuaternion) ToRotationMatrix(args ...interface{}) () {
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
    C.C_ZNK11QQuaternion16toRotationMatrixEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QQuaternion", "toRotationMatrix", args)
  }

  return
}

// <= body block end

