package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QVector4D QMatrix4x4::row(int index);
extern void _ZNK10QMatrix4x43rowEi(void* qthis, int32_t arg0); // 2
  // proto:  QVector3D QMatrix4x4::mapVector(const QVector3D & vector);
extern void _ZNK10QMatrix4x49mapVectorERK9QVector3D(void* qthis, void* arg0); // 2
  // proto:  QRectF QMatrix4x4::mapRect(const QRectF & rect);
extern void _ZNK10QMatrix4x47mapRectERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  QRect QMatrix4x4::mapRect(const QRect & rect);
extern void _ZNK10QMatrix4x47mapRectERK5QRect(void* qthis, void* arg0); // 4
  // proto:  QMatrix4x4 QMatrix4x4::inverted(bool * invertible);
extern void _ZNK10QMatrix4x48invertedEPb(void* qthis, bool* arg0); // 4
  // proto:  void QMatrix4x4::lookAt(const QVector3D & eye, const QVector3D & center, const QVector3D & up);
extern void _ZN10QMatrix4x46lookAtERK9QVector3DS2_S2_(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QMatrix4x4::perspective(float verticalAngle, float aspectRatio, float nearPlane, float farPlane);
extern void _ZN10QMatrix4x411perspectiveEffff(void* qthis, float arg0, float arg1, float arg2, float arg3); // 4
  // proto:  void QMatrix4x4::fill(float value);
extern void _ZN10QMatrix4x44fillEf(void* qthis, float arg0); // 2
  // proto:  void QMatrix4x4::flipCoordinates();
extern void _ZN10QMatrix4x415flipCoordinatesEv(void* qthis); // 4
  // proto:  void QMatrix4x4::scale(float x, float y);
extern void _ZN10QMatrix4x45scaleEff(void* qthis, float arg0, float arg1); // 4
  // proto:  void QMatrix4x4::scale(float factor);
extern void _ZN10QMatrix4x45scaleEf(void* qthis, float arg0); // 4
  // proto:  void QMatrix4x4::scale(const QVector3D & vector);
extern void _ZN10QMatrix4x45scaleERK9QVector3D(void* qthis, void* arg0); // 4
  // proto:  void QMatrix4x4::scale(float x, float y, float z);
extern void _ZN10QMatrix4x45scaleEfff(void* qthis, float arg0, float arg1, float arg2); // 4
  // proto:  const float * QMatrix4x4::constData();
extern void _ZNK10QMatrix4x49constDataEv(void* qthis); // 2
  // proto:  void QMatrix4x4::copyDataTo(float * values);
extern void _ZNK10QMatrix4x410copyDataToEPf(void* qthis, float* arg0); // 4
  // proto:  bool QMatrix4x4::isIdentity();
extern void _ZNK10QMatrix4x410isIdentityEv(void* qthis); // 2
  // proto:  void QMatrix4x4::translate(const QVector3D & vector);
extern void _ZN10QMatrix4x49translateERK9QVector3D(void* qthis, void* arg0); // 4
  // proto:  void QMatrix4x4::translate(float x, float y, float z);
extern void _ZN10QMatrix4x49translateEfff(void* qthis, float arg0, float arg1, float arg2); // 4
  // proto:  void QMatrix4x4::translate(float x, float y);
extern void _ZN10QMatrix4x49translateEff(void* qthis, float arg0, float arg1); // 4
  // proto:  void QMatrix4x4::setToIdentity();
extern void _ZN10QMatrix4x413setToIdentityEv(void* qthis); // 2
  // proto:  QPoint QMatrix4x4::map(const QPoint & point);
extern void _ZNK10QMatrix4x43mapERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  QPointF QMatrix4x4::map(const QPointF & point);
extern void _ZNK10QMatrix4x43mapERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  QVector4D QMatrix4x4::map(const QVector4D & point);
extern void _ZNK10QMatrix4x43mapERK9QVector4D(void* qthis, void* arg0); // 2
  // proto:  QVector3D QMatrix4x4::map(const QVector3D & point);
extern void _ZNK10QMatrix4x43mapERK9QVector3D(void* qthis, void* arg0); // 2
  // proto:  double QMatrix4x4::determinant();
extern void _ZNK10QMatrix4x411determinantEv(void* qthis); // 4
  // proto:  void QMatrix4x4::frustum(float left, float right, float bottom, float top, float nearPlane, float farPlane);
extern void _ZN10QMatrix4x47frustumEffffff(void* qthis, float arg0, float arg1, float arg2, float arg3, float arg4, float arg5); // 4
  // proto:  void QMatrix4x4::setColumn(int index, const QVector4D & value);
extern void _ZN10QMatrix4x49setColumnEiRK9QVector4D(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  QMatrix3x3 QMatrix4x4::normalMatrix();
extern void _ZNK10QMatrix4x412normalMatrixEv(void* qthis); // 4
  // proto:  QTransform QMatrix4x4::toTransform(float distanceToPlane);
extern void _ZNK10QMatrix4x411toTransformEf(void* qthis, float arg0); // 4
  // proto:  QTransform QMatrix4x4::toTransform();
extern void _ZNK10QMatrix4x411toTransformEv(void* qthis); // 4
  // proto:  QMatrix QMatrix4x4::toAffine();
extern void _ZNK10QMatrix4x48toAffineEv(void* qthis); // 4
  // proto:  float * QMatrix4x4::data();
extern void _ZN10QMatrix4x44dataEv(void* qthis); // 2
  // proto:  void QMatrix4x4::optimize();
extern void _ZN10QMatrix4x48optimizeEv(void* qthis); // 4
  // proto:  void QMatrix4x4::viewport(float left, float bottom, float width, float height, float nearPlane, float farPlane);
extern void _ZN10QMatrix4x48viewportEffffff(void* qthis, float arg0, float arg1, float arg2, float arg3, float arg4, float arg5); // 4
  // proto:  void QMatrix4x4::viewport(const QRectF & rect);
extern void _ZN10QMatrix4x48viewportERK6QRectF(void* qthis, void* arg0); // 2
  // proto:  QMatrix4x4 QMatrix4x4::transposed();
extern void _ZNK10QMatrix4x410transposedEv(void* qthis); // 4
  // proto:  void QMatrix4x4::rotate(float angle, const QVector3D & vector);
extern void _ZN10QMatrix4x46rotateEfRK9QVector3D(void* qthis, float arg0, void* arg1); // 4
  // proto:  void QMatrix4x4::rotate(const QQuaternion & quaternion);
extern void _ZN10QMatrix4x46rotateERK11QQuaternion(void* qthis, void* arg0); // 4
  // proto:  void QMatrix4x4::rotate(float angle, float x, float y, float z);
extern void _ZN10QMatrix4x46rotateEffff(void* qthis, float arg0, float arg1, float arg2, float arg3); // 4
  // proto:  QVector4D QMatrix4x4::column(int index);
extern void _ZNK10QMatrix4x46columnEi(void* qthis, int32_t arg0); // 2
  // proto:  void QMatrix4x4::ortho(float left, float right, float bottom, float top, float nearPlane, float farPlane);
extern void _ZN10QMatrix4x45orthoEffffff(void* qthis, float arg0, float arg1, float arg2, float arg3, float arg4, float arg5); // 4
  // proto:  void QMatrix4x4::ortho(const QRect & rect);
extern void _ZN10QMatrix4x45orthoERK5QRect(void* qthis, void* arg0); // 4
  // proto:  void QMatrix4x4::ortho(const QRectF & rect);
extern void _ZN10QMatrix4x45orthoERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QMatrix4x4::QMatrix4x4(const QTransform & transform);
extern void _ZN10QMatrix4x4C2ERK10QTransform(void* qthis, void* arg0); // 3
  // proto:  void QMatrix4x4::QMatrix4x4(const float * values, int cols, int rows);
extern void _ZN10QMatrix4x4C2EPKfii(void* qthis, float* arg0, int32_t arg1, int32_t arg2); // 3
  // proto:  void QMatrix4x4::QMatrix4x4(const float * values);
extern void _ZN10QMatrix4x4C2EPKf(void* qthis, float* arg0); // 3
  // proto:  void QMatrix4x4::QMatrix4x4(const QMatrix & matrix);
extern void _ZN10QMatrix4x4C2ERK7QMatrix(void* qthis, void* arg0); // 3
  // proto:  void QMatrix4x4::QMatrix4x4();
extern void _ZN10QMatrix4x4C2Ev(void* qthis); // 1
  // proto:  void QMatrix4x4::QMatrix4x4(float m11, float m12, float m13, float m14, float m21, float m22, float m23, float m24, float m31, float m32, float m33, float m34, float m41, float m42, float m43, float m44);
extern void _ZN10QMatrix4x4C2Effffffffffffffff(void* qthis, float arg0, float arg1, float arg2, float arg3, float arg4, float arg5, float arg6, float arg7, float arg8, float arg9, float arg10, float arg11, float arg12, float arg13, float arg14, float arg15); // 1
  // proto:  void QMatrix4x4::setRow(int index, const QVector4D & value);
extern void _ZN10QMatrix4x46setRowEiRK9QVector4D(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  bool QMatrix4x4::isAffine();
extern void _ZNK10QMatrix4x48isAffineEv(void* qthis); // 2
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

// row(int)
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
    C._ZNK10QMatrix4x43rowEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "row", args)
  }

}

// mapVector(const class QVector3D &)
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
    C._ZNK10QMatrix4x49mapVectorERK9QVector3D(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "mapVector", args)
  }

}

// mapRect(const class QRectF &)
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

// inverted(_Bool *)
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

// lookAt(const class QVector3D &, const class QVector3D &, const class QVector3D &)
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

// perspective(float, float, float, float)
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

// fill(float)
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
    C._ZN10QMatrix4x44fillEf(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "fill", args)
  }

}

// flipCoordinates()
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

// scale(float, float)
func (this *QMatrix4x4) scale(args ...interface{}) () {
  // scale(float, float)
  // scale(float)
  // scale(const class QVector3D &)
  // scale(float, float, float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"
  vtys[0][1] = qtrt.FloatTy(false) // "float"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.FloatTy(false) // "float"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.FloatTy(false) // "float"
  vtys[3][1] = qtrt.FloatTy(false) // "float"
  vtys[3][2] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QMatrix4x45scaleEff
    // invoke: void scale(float, float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    C._ZN10QMatrix4x45scaleEff(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN10QMatrix4x45scaleEf
    // invoke: void scale(float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    C._ZN10QMatrix4x45scaleEf(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN10QMatrix4x45scaleERK9QVector3D
    // invoke: void scale(const class QVector3D &)
    var arg0 = args[0].(QVector3D).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QMatrix4x45scaleERK9QVector3D(this.qclsinst, arg0)
  case 3:
    // invoke: _ZN10QMatrix4x45scaleEfff
    // invoke: void scale(float, float, float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    C._ZN10QMatrix4x45scaleEfff(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "scale", args)
  }

}

// constData()
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
    C._ZNK10QMatrix4x49constDataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "constData", args)
  }

}

// copyDataTo(float *)
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

// isIdentity()
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
    C._ZNK10QMatrix4x410isIdentityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "isIdentity", args)
  }

}

// translate(const class QVector3D &)
func (this *QMatrix4x4) translate(args ...interface{}) () {
  // translate(const class QVector3D &)
  // translate(float, float, float)
  // translate(float, float)
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

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QMatrix4x49translateERK9QVector3D
    // invoke: void translate(const class QVector3D &)
    var arg0 = args[0].(QVector3D).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QMatrix4x49translateERK9QVector3D(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN10QMatrix4x49translateEfff
    // invoke: void translate(float, float, float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    C._ZN10QMatrix4x49translateEfff(this.qclsinst, arg0, arg1, arg2)
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

// setToIdentity()
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
    C._ZN10QMatrix4x413setToIdentityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "setToIdentity", args)
  }

}

// map(const class QPoint &)
func (this *QMatrix4x4) map_(args ...interface{}) () {
  // map(const class QPoint &)
  // map(const class QPointF &)
  // map(const class QVector4D &)
  // map(const class QVector3D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QVector4D{}) // "const QVector4D &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QMatrix4x43mapERK6QPoint
    // invoke: QPoint map(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK10QMatrix4x43mapERK6QPoint(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK10QMatrix4x43mapERK7QPointF
    // invoke: QPointF map(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK10QMatrix4x43mapERK7QPointF(this.qclsinst, arg0)
  case 2:
    // invoke: _ZNK10QMatrix4x43mapERK9QVector4D
    // invoke: QVector4D map(const class QVector4D &)
    var arg0 = args[0].(QVector4D).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK10QMatrix4x43mapERK9QVector4D(this.qclsinst, arg0)
  case 3:
    // invoke: _ZNK10QMatrix4x43mapERK9QVector3D
    // invoke: QVector3D map(const class QVector3D &)
    var arg0 = args[0].(QVector3D).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK10QMatrix4x43mapERK9QVector3D(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "map", args)
  }

}

// determinant()
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

// frustum(float, float, float, float, float, float)
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

// setColumn(int, const class QVector4D &)
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
    C._ZN10QMatrix4x49setColumnEiRK9QVector4D(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "setColumn", args)
  }

}

// normalMatrix()
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

// toTransform(float)
func (this *QMatrix4x4) toTransform(args ...interface{}) () {
  // toTransform(float)
  // toTransform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QMatrix4x411toTransformEf
    // invoke: QTransform toTransform(float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    C._ZNK10QMatrix4x411toTransformEf(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK10QMatrix4x411toTransformEv
    // invoke: QTransform toTransform()
    C._ZNK10QMatrix4x411toTransformEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "toTransform", args)
  }

}

// toAffine()
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

// data()
func (this *QMatrix4x4) data(args ...interface{}) () {
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QMatrix4x44dataEv
    // invoke: float * data()
    C._ZN10QMatrix4x44dataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "data", args)
  }

}

// optimize()
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

// viewport(float, float, float, float, float, float)
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
    C._ZN10QMatrix4x48viewportERK6QRectF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "viewport", args)
  }

}

// transposed()
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

// rotate(float, const class QVector3D &)
func (this *QMatrix4x4) rotate(args ...interface{}) () {
  // rotate(float, const class QVector3D &)
  // rotate(const class QQuaternion &)
  // rotate(float, float, float, float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"
  vtys[0][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QQuaternion{}) // "const QQuaternion &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.FloatTy(false) // "float"
  vtys[2][1] = qtrt.FloatTy(false) // "float"
  vtys[2][2] = qtrt.FloatTy(false) // "float"
  vtys[2][3] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QMatrix4x46rotateEfRK9QVector3D
    // invoke: void rotate(float, const class QVector3D &)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector3D).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN10QMatrix4x46rotateEfRK9QVector3D(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN10QMatrix4x46rotateERK11QQuaternion
    // invoke: void rotate(const class QQuaternion &)
    var arg0 = args[0].(QQuaternion).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QMatrix4x46rotateERK11QQuaternion(this.qclsinst, arg0)
  case 2:
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
  default:
    qtrt.ErrorResolve("QMatrix4x4", "rotate", args)
  }

}

// column(int)
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
    C._ZNK10QMatrix4x46columnEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "column", args)
  }

}

// ortho(float, float, float, float, float, float)
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

// QMatrix4x4(const class QTransform &)
func NewQMatrix4x4(args ...interface{}) QMatrix4x4 {
  // QMatrix4x4(const class QTransform &)
  // QMatrix4x4(const float *, int, int)
  // QMatrix4x4(const float *)
  // QMatrix4x4(const class QMatrix &)
  // QMatrix4x4()
  // QMatrix4x4(float, float, float, float, float, float, float, float, float, float, float, float, float, float, float, float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.FloatTy(true) // "const float *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.FloatTy(true) // "const float *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QMatrix{}) // "const QMatrix &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.FloatTy(false) // "float"
  vtys[5][1] = qtrt.FloatTy(false) // "float"
  vtys[5][2] = qtrt.FloatTy(false) // "float"
  vtys[5][3] = qtrt.FloatTy(false) // "float"
  vtys[5][4] = qtrt.FloatTy(false) // "float"
  vtys[5][5] = qtrt.FloatTy(false) // "float"
  vtys[5][6] = qtrt.FloatTy(false) // "float"
  vtys[5][7] = qtrt.FloatTy(false) // "float"
  vtys[5][8] = qtrt.FloatTy(false) // "float"
  vtys[5][9] = qtrt.FloatTy(false) // "float"
  vtys[5][10] = qtrt.FloatTy(false) // "float"
  vtys[5][11] = qtrt.FloatTy(false) // "float"
  vtys[5][12] = qtrt.FloatTy(false) // "float"
  vtys[5][13] = qtrt.FloatTy(false) // "float"
  vtys[5][14] = qtrt.FloatTy(false) // "float"
  vtys[5][15] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QMatrix4x4C1ERK10QTransform
    // invoke: void QMatrix4x4(const class QTransform &)
    var arg0 = args[0].(QTransform).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN10QMatrix4x4C2ERK10QTransform(qthis, arg0)
  case 1:
    // invoke: _ZN10QMatrix4x4C1EPKfii
    // invoke: void QMatrix4x4(const float *, int, int)
    var arg0 = (*C.float)(args[0].(*float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN10QMatrix4x4C2EPKfii(qthis, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN10QMatrix4x4C1EPKf
    // invoke: void QMatrix4x4(const float *)
    var arg0 = (*C.float)(args[0].(*float32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN10QMatrix4x4C2EPKf(qthis, arg0)
  case 3:
    // invoke: _ZN10QMatrix4x4C1ERK7QMatrix
    // invoke: void QMatrix4x4(const class QMatrix &)
    var arg0 = args[0].(QMatrix).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN10QMatrix4x4C2ERK7QMatrix(qthis, arg0)
  case 4:
    // invoke: _ZN10QMatrix4x4C1Ev
    // invoke: void QMatrix4x4()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN10QMatrix4x4C2Ev(qthis)
  case 5:
    // invoke: _ZN10QMatrix4x4C1Effffffffffffffff
    // invoke: void QMatrix4x4(float, float, float, float, float, float, float, float, float, float, float, float, float, float, float, float)
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
    var arg6 = C.float(args[6].(float32))
    if false {fmt.Println(arg6)}
    var arg7 = C.float(args[7].(float32))
    if false {fmt.Println(arg7)}
    var arg8 = C.float(args[8].(float32))
    if false {fmt.Println(arg8)}
    var arg9 = C.float(args[9].(float32))
    if false {fmt.Println(arg9)}
    var arg10 = C.float(args[10].(float32))
    if false {fmt.Println(arg10)}
    var arg11 = C.float(args[11].(float32))
    if false {fmt.Println(arg11)}
    var arg12 = C.float(args[12].(float32))
    if false {fmt.Println(arg12)}
    var arg13 = C.float(args[13].(float32))
    if false {fmt.Println(arg13)}
    var arg14 = C.float(args[14].(float32))
    if false {fmt.Println(arg14)}
    var arg15 = C.float(args[15].(float32))
    if false {fmt.Println(arg15)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN10QMatrix4x4C2Effffffffffffffff(qthis, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "QMatrix4x4", args)
  }

  return QMatrix4x4{}
}

// setRow(int, const class QVector4D &)
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
    C._ZN10QMatrix4x46setRowEiRK9QVector4D(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "setRow", args)
  }

}

// isAffine()
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
    C._ZNK10QMatrix4x48isAffineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMatrix4x4", "isAffine", args)
  }

}

// <= body block end

