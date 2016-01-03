package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
// src-file: /QtGui/qmatrix4x4.h
// dst-file: /src/gui/qmatrix4x4.go
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
  // proto:  QTransform QMatrix4x4::toTransform();
extern void _ZNK10QMatrix4x411toTransformEv(void* qthis);
  // proto:  void QMatrix4x4::scale(const QVector3D & vector);
extern void _ZN10QMatrix4x45scaleERK9QVector3D(void* qthis, void* arg0);
  // proto:  void QMatrix4x4::translate(float x, float y, float z);
extern void _ZN10QMatrix4x49translateEfff(void* qthis, float arg0, float arg1, float arg2);
  // proto:  const float * QMatrix4x4::constData();
extern void demth_ZNK10QMatrix4x49constDataEv(void* qthis);
  // proto:  float * QMatrix4x4::data();
extern void demth_ZN10QMatrix4x44dataEv(void* qthis);
  // proto:  QMatrix4x4 QMatrix4x4::inverted(bool * invertible);
extern void _ZNK10QMatrix4x48invertedEPb(void* qthis, bool* arg0);
  // proto:  QVector3D QMatrix4x4::mapVector(const QVector3D & vector);
extern void demth_ZNK10QMatrix4x49mapVectorERK9QVector3D(void* qthis, void* arg0);
  // proto:  void QMatrix4x4::ortho(float left, float right, float bottom, float top, float nearPlane, float farPlane);
extern void _ZN10QMatrix4x45orthoEffffff(void* qthis, float arg0, float arg1, float arg2, float arg3, float arg4, float arg5);
  // proto:  void QMatrix4x4::QMatrix4x4();
extern void* dector_ZN10QMatrix4x4C1Ev();
extern void demth_ZN10QMatrix4x4C1Ev(void* qthis);
  // proto:  QMatrix QMatrix4x4::toAffine();
extern void _ZNK10QMatrix4x48toAffineEv(void* qthis);
  // proto:  QRectF QMatrix4x4::mapRect(const QRectF & rect);
extern void _ZNK10QMatrix4x47mapRectERK6QRectF(void* qthis, void* arg0);
  // proto:  void QMatrix4x4::setColumn(int index, const QVector4D & value);
extern void demth_ZN10QMatrix4x49setColumnEiRK9QVector4D(void* qthis, int32_t arg0, void* arg1);
  // proto:  bool QMatrix4x4::isIdentity();
extern void demth_ZNK10QMatrix4x410isIdentityEv(void* qthis);
  // proto:  QVector4D QMatrix4x4::column(int index);
extern void demth_ZNK10QMatrix4x46columnEi(void* qthis, int32_t arg0);
  // proto:  void QMatrix4x4::setRow(int index, const QVector4D & value);
extern void demth_ZN10QMatrix4x46setRowEiRK9QVector4D(void* qthis, int32_t arg0, void* arg1);
  // proto:  void QMatrix4x4::flipCoordinates();
extern void _ZN10QMatrix4x415flipCoordinatesEv(void* qthis);
  // proto:  QMatrix3x3 QMatrix4x4::normalMatrix();
extern void _ZNK10QMatrix4x412normalMatrixEv(void* qthis);
  // proto:  void QMatrix4x4::viewport(float left, float bottom, float width, float height, float nearPlane, float farPlane);
extern void _ZN10QMatrix4x48viewportEffffff(void* qthis, float arg0, float arg1, float arg2, float arg3, float arg4, float arg5);
  // proto:  void QMatrix4x4::copyDataTo(float * values);
extern void _ZNK10QMatrix4x410copyDataToEPf(void* qthis, float* arg0);
  // proto:  void QMatrix4x4::QMatrix4x4(const QTransform & transform);
extern void* dector_ZN10QMatrix4x4C1ERK10QTransform(void* arg0);
extern void _ZN10QMatrix4x4C1ERK10QTransform(void* qthis, void* arg0);
  // proto:  bool QMatrix4x4::isAffine();
extern void demth_ZNK10QMatrix4x48isAffineEv(void* qthis);
  // proto:  void QMatrix4x4::ortho(const QRect & rect);
extern void _ZN10QMatrix4x45orthoERK5QRect(void* qthis, void* arg0);
  // proto:  void QMatrix4x4::rotate(const QQuaternion & quaternion);
extern void _ZN10QMatrix4x46rotateERK11QQuaternion(void* qthis, void* arg0);
  // proto:  void QMatrix4x4::QMatrix4x4(const QMatrix & matrix);
extern void* dector_ZN10QMatrix4x4C1ERK7QMatrix(void* arg0);
extern void _ZN10QMatrix4x4C1ERK7QMatrix(void* qthis, void* arg0);
  // proto:  void QMatrix4x4::perspective(float verticalAngle, float aspectRatio, float nearPlane, float farPlane);
extern void _ZN10QMatrix4x411perspectiveEffff(void* qthis, float arg0, float arg1, float arg2, float arg3);
  // proto:  void QMatrix4x4::translate(const QVector3D & vector);
extern void _ZN10QMatrix4x49translateERK9QVector3D(void* qthis, void* arg0);
  // proto:  double QMatrix4x4::determinant();
extern void _ZNK10QMatrix4x411determinantEv(void* qthis);
  // proto:  void QMatrix4x4::scale(float x, float y, float z);
extern void _ZN10QMatrix4x45scaleEfff(void* qthis, float arg0, float arg1, float arg2);
  // proto:  void QMatrix4x4::frustum(float left, float right, float bottom, float top, float nearPlane, float farPlane);
extern void _ZN10QMatrix4x47frustumEffffff(void* qthis, float arg0, float arg1, float arg2, float arg3, float arg4, float arg5);
  // proto:  QPoint QMatrix4x4::map(const QPoint & point);
extern void demth_ZNK10QMatrix4x43mapERK6QPoint(void* qthis, void* arg0);
  // proto:  void QMatrix4x4::QMatrix4x4(int );
extern void* dector_ZN10QMatrix4x4C1Ei(int32_t arg0);
extern void _ZN10QMatrix4x4C1Ei(void* qthis, int32_t arg0);
  // proto:  void QMatrix4x4::optimize();
extern void _ZN10QMatrix4x48optimizeEv(void* qthis);
  // proto:  void QMatrix4x4::QMatrix4x4(const float * values);
extern void* dector_ZN10QMatrix4x4C1EPKf(float* arg0);
extern void _ZN10QMatrix4x4C1EPKf(void* qthis, float* arg0);
  // proto:  void QMatrix4x4::translate(float x, float y);
extern void _ZN10QMatrix4x49translateEff(void* qthis, float arg0, float arg1);
  // proto:  void QMatrix4x4::setToIdentity();
extern void demth_ZN10QMatrix4x413setToIdentityEv(void* qthis);
  // proto:  QRect QMatrix4x4::mapRect(const QRect & rect);
extern void _ZNK10QMatrix4x47mapRectERK5QRect(void* qthis, void* arg0);
  // proto:  void QMatrix4x4::scale(float x, float y);
extern void _ZN10QMatrix4x45scaleEff(void* qthis, float arg0, float arg1);
  // proto:  void QMatrix4x4::QMatrix4x4(float m11, float m12, float m13, float m14, float m21, float m22, float m23, float m24, float m31, float m32, float m33, float m34, float m41, float m42, float m43, float m44);
extern void* dector_ZN10QMatrix4x4C1Effffffffffffffff(float arg0, float arg1, float arg2, float arg3, float arg4, float arg5, float arg6, float arg7, float arg8, float arg9, float arg10, float arg11, float arg12, float arg13, float arg14, float arg15);
extern void demth_ZN10QMatrix4x4C1Effffffffffffffff(void* qthis, float arg0, float arg1, float arg2, float arg3, float arg4, float arg5, float arg6, float arg7, float arg8, float arg9, float arg10, float arg11, float arg12, float arg13, float arg14, float arg15);
  // proto:  QVector3D QMatrix4x4::map(const QVector3D & point);
extern void demth_ZNK10QMatrix4x43mapERK9QVector3D(void* qthis, void* arg0);
  // proto:  void QMatrix4x4::lookAt(const QVector3D & eye, const QVector3D & center, const QVector3D & up);
extern void _ZN10QMatrix4x46lookAtERK9QVector3DS2_S2_(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QMatrix4x4::ortho(const QRectF & rect);
extern void _ZN10QMatrix4x45orthoERK6QRectF(void* qthis, void* arg0);
  // proto:  void QMatrix4x4::viewport(const QRectF & rect);
extern void demth_ZN10QMatrix4x48viewportERK6QRectF(void* qthis, void* arg0);
  // proto:  void QMatrix4x4::rotate(float angle, float x, float y, float z);
extern void _ZN10QMatrix4x46rotateEffff(void* qthis, float arg0, float arg1, float arg2, float arg3);
  // proto:  void QMatrix4x4::fill(float value);
extern void demth_ZN10QMatrix4x44fillEf(void* qthis, float arg0);
  // proto:  void QMatrix4x4::QMatrix4x4(const float * values, int cols, int rows);
extern void* dector_ZN10QMatrix4x4C1EPKfii(float* arg0, int32_t arg1, int32_t arg2);
extern void _ZN10QMatrix4x4C1EPKfii(void* qthis, float* arg0, int32_t arg1, int32_t arg2);
  // proto:  QTransform QMatrix4x4::toTransform(float distanceToPlane);
extern void _ZNK10QMatrix4x411toTransformEf(void* qthis, float arg0);
  // proto:  QMatrix4x4 QMatrix4x4::transposed();
extern void _ZNK10QMatrix4x410transposedEv(void* qthis);
  // proto:  QPointF QMatrix4x4::map(const QPointF & point);
extern void demth_ZNK10QMatrix4x43mapERK7QPointF(void* qthis, void* arg0);
  // proto:  void QMatrix4x4::scale(float factor);
extern void _ZN10QMatrix4x45scaleEf(void* qthis, float arg0);
  // proto:  QVector4D QMatrix4x4::row(int index);
extern void demth_ZNK10QMatrix4x43rowEi(void* qthis, int32_t arg0);
  // proto:  void QMatrix4x4::rotate(float angle, const QVector3D & vector);
extern void _ZN10QMatrix4x46rotateEfRK9QVector3D(void* qthis, float arg0, void* arg1);
  // proto:  QVector4D QMatrix4x4::map(const QVector4D & point);
extern void demth_ZNK10QMatrix4x43mapERK9QVector4D(void* qthis, void* arg0);
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

// class sizeof(QMatrix4x4)=68
type QMatrix4x4 struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  QTransform QMatrix4x4::toTransform();
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
    // invoke: QTransform toTransform()
    C._ZNK10QMatrix4x411toTransformEv(this.qclsinst)
  case 1:
    // invoke: _ZNK10QMatrix4x411toTransformEf
    // invoke: QTransform toTransform(float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    C._ZNK10QMatrix4x411toTransformEf(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "toTransform", args)
  }

}

  // proto:  void QMatrix4x4::scale(const QVector3D & vector);
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
    // invoke: void scale(const class QVector3D &)
    var arg0 = args[0].(QVector3D).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QMatrix4x45scaleERK9QVector3D(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN10QMatrix4x45scaleEfff
    // invoke: void scale(float, float, float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    C._ZN10QMatrix4x45scaleEfff(this.qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN10QMatrix4x45scaleEff
    // invoke: void scale(float, float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    C._ZN10QMatrix4x45scaleEff(this.qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZN10QMatrix4x45scaleEf
    // invoke: void scale(float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    C._ZN10QMatrix4x45scaleEf(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "scale", args)
  }

}

  // proto:  void QMatrix4x4::translate(float x, float y, float z);
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
    // invoke: void translate(float, float, float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    C._ZN10QMatrix4x49translateEfff(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN10QMatrix4x49translateERK9QVector3D
    // invoke: void translate(const class QVector3D &)
    var arg0 = args[0].(QVector3D).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QMatrix4x49translateERK9QVector3D(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN10QMatrix4x49translateEff
    // invoke: void translate(float, float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    C._ZN10QMatrix4x49translateEff(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "translate", args)
  }

}

  // proto:  const float * QMatrix4x4::constData();
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
    // invoke: const float * constData()
    C.demth_ZNK10QMatrix4x49constDataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "constData", args)
  }

}

  // proto:  float * QMatrix4x4::data();
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
    // invoke: float * data()
    C.demth_ZN10QMatrix4x44dataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "data", args)
  }

}

  // proto:  QMatrix4x4 QMatrix4x4::inverted(bool * invertible);
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
    // invoke: QMatrix4x4 inverted(_Bool *)
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    C._ZNK10QMatrix4x48invertedEPb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "inverted", args)
  }

}

  // proto:  QVector3D QMatrix4x4::mapVector(const QVector3D & vector);
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
    // invoke: QVector3D mapVector(const class QVector3D &)
    var arg0 = args[0].(QVector3D).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZNK10QMatrix4x49mapVectorERK9QVector3D(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "mapVector", args)
  }

}

  // proto:  void QMatrix4x4::ortho(float left, float right, float bottom, float top, float nearPlane, float farPlane);
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
    // invoke: void ortho(float, float, float, float, float, float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(args[3].(float32))
    if false {fmt.Println(arg3)}
    var arg4 = C.float(args[4].(float32))
    if false {fmt.Println(arg4)}
    var arg5 = C.float(args[5].(float32))
    if false {fmt.Println(arg5)}
    C._ZN10QMatrix4x45orthoEffffff(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  case 1:
    // invoke: _ZN10QMatrix4x45orthoERK5QRect
    // invoke: void ortho(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QMatrix4x45orthoERK5QRect(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN10QMatrix4x45orthoERK6QRectF
    // invoke: void ortho(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QMatrix4x45orthoERK6QRectF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "ortho", args)
  }

}

  // proto:  void QMatrix4x4::QMatrix4x4();
func NewQMatrix4x4(args ...interface{}) QMatrix4x4 {
  return QMatrix4x4{}
}

  // proto:  QMatrix QMatrix4x4::toAffine();
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
    // invoke: QMatrix toAffine()
    C._ZNK10QMatrix4x48toAffineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "toAffine", args)
  }

}

  // proto:  QRectF QMatrix4x4::mapRect(const QRectF & rect);
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
    // invoke: QRectF mapRect(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK10QMatrix4x47mapRectERK6QRectF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK10QMatrix4x47mapRectERK5QRect
    // invoke: QRect mapRect(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK10QMatrix4x47mapRectERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "mapRect", args)
  }

}

  // proto:  void QMatrix4x4::setColumn(int index, const QVector4D & value);
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
    // invoke: void setColumn(int, const class QVector4D &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector4D).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZN10QMatrix4x49setColumnEiRK9QVector4D(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "setColumn", args)
  }

}

  // proto:  bool QMatrix4x4::isIdentity();
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
    // invoke: bool isIdentity()
    C.demth_ZNK10QMatrix4x410isIdentityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "isIdentity", args)
  }

}

  // proto:  QVector4D QMatrix4x4::column(int index);
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
    // invoke: QVector4D column(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.demth_ZNK10QMatrix4x46columnEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "column", args)
  }

}

  // proto:  void QMatrix4x4::setRow(int index, const QVector4D & value);
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
    // invoke: void setRow(int, const class QVector4D &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector4D).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZN10QMatrix4x46setRowEiRK9QVector4D(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "setRow", args)
  }

}

  // proto:  void QMatrix4x4::flipCoordinates();
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
    // invoke: void flipCoordinates()
    C._ZN10QMatrix4x415flipCoordinatesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "flipCoordinates", args)
  }

}

  // proto:  QMatrix3x3 QMatrix4x4::normalMatrix();
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
    // invoke: QMatrix3x3 normalMatrix()
    C._ZNK10QMatrix4x412normalMatrixEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "normalMatrix", args)
  }

}

  // proto:  void QMatrix4x4::viewport(float left, float bottom, float width, float height, float nearPlane, float farPlane);
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
    // invoke: void viewport(float, float, float, float, float, float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(args[3].(float32))
    if false {fmt.Println(arg3)}
    var arg4 = C.float(args[4].(float32))
    if false {fmt.Println(arg4)}
    var arg5 = C.float(args[5].(float32))
    if false {fmt.Println(arg5)}
    C._ZN10QMatrix4x48viewportEffffff(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  case 1:
    // invoke: _ZN10QMatrix4x48viewportERK6QRectF
    // invoke: void viewport(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN10QMatrix4x48viewportERK6QRectF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "viewport", args)
  }

}

  // proto:  void QMatrix4x4::copyDataTo(float * values);
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
    // invoke: void copyDataTo(float *)
    var arg0 = (*C.float)(args[0].(*float32))
    if false {fmt.Println(arg0)}
    C._ZNK10QMatrix4x410copyDataToEPf(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "copyDataTo", args)
  }

}

  // proto:  bool QMatrix4x4::isAffine();
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
    // invoke: bool isAffine()
    C.demth_ZNK10QMatrix4x48isAffineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "isAffine", args)
  }

}

  // proto:  void QMatrix4x4::rotate(const QQuaternion & quaternion);
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
    // invoke: void rotate(const class QQuaternion &)
    var arg0 = args[0].(QQuaternion).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QMatrix4x46rotateERK11QQuaternion(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN10QMatrix4x46rotateEffff
    // invoke: void rotate(float, float, float, float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(args[3].(float32))
    if false {fmt.Println(arg3)}
    C._ZN10QMatrix4x46rotateEffff(this.qclsinst, arg0, arg1, arg2, arg3)
  case 2:
    // invoke: _ZN10QMatrix4x46rotateEfRK9QVector3D
    // invoke: void rotate(float, const class QVector3D &)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector3D).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN10QMatrix4x46rotateEfRK9QVector3D(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "rotate", args)
  }

}

  // proto:  void QMatrix4x4::perspective(float verticalAngle, float aspectRatio, float nearPlane, float farPlane);
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
    // invoke: void perspective(float, float, float, float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(args[3].(float32))
    if false {fmt.Println(arg3)}
    C._ZN10QMatrix4x411perspectiveEffff(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "perspective", args)
  }

}

  // proto:  double QMatrix4x4::determinant();
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
    // invoke: double determinant()
    C._ZNK10QMatrix4x411determinantEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "determinant", args)
  }

}

  // proto:  void QMatrix4x4::frustum(float left, float right, float bottom, float top, float nearPlane, float farPlane);
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
    // invoke: void frustum(float, float, float, float, float, float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(args[3].(float32))
    if false {fmt.Println(arg3)}
    var arg4 = C.float(args[4].(float32))
    if false {fmt.Println(arg4)}
    var arg5 = C.float(args[5].(float32))
    if false {fmt.Println(arg5)}
    C._ZN10QMatrix4x47frustumEffffff(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "frustum", args)
  }

}

  // proto:  QPoint QMatrix4x4::map(const QPoint & point);
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

  // proto:  void QMatrix4x4::optimize();
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
    // invoke: void optimize()
    C._ZN10QMatrix4x48optimizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "optimize", args)
  }

}

  // proto:  void QMatrix4x4::setToIdentity();
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
    // invoke: void setToIdentity()
    C.demth_ZN10QMatrix4x413setToIdentityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "setToIdentity", args)
  }

}

  // proto:  void QMatrix4x4::lookAt(const QVector3D & eye, const QVector3D & center, const QVector3D & up);
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
    // invoke: void lookAt(const class QVector3D &, const class QVector3D &, const class QVector3D &)
    var arg0 = args[0].(QVector3D).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector3D).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QVector3D).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN10QMatrix4x46lookAtERK9QVector3DS2_S2_(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "lookAt", args)
  }

}

  // proto:  void QMatrix4x4::fill(float value);
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
    // invoke: void fill(float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    C.demth_ZN10QMatrix4x44fillEf(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "fill", args)
  }

}

  // proto:  QMatrix4x4 QMatrix4x4::transposed();
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
    // invoke: QMatrix4x4 transposed()
    C._ZNK10QMatrix4x410transposedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "transposed", args)
  }

}

  // proto:  QVector4D QMatrix4x4::row(int index);
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
    // invoke: QVector4D row(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.demth_ZNK10QMatrix4x43rowEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "row", args)
  }

}

// <= body block end

