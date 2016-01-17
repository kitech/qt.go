package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtGui/qvector4d.h
// dst-file: /src/gui/qvector4d.go
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
  // proto:  QVector2D QVector4D::toVector2D();
extern void _ZNK9QVector4D10toVector2DEv(void* qthis); // 4
  // proto:  QVector3D QVector4D::toVector3D();
extern void _ZNK9QVector4D10toVector3DEv(void* qthis); // 4
  // proto:  void QVector4D::QVector4D();
extern void _ZN9QVector4DC2Ev(void* qthis); // 1
  // proto:  void QVector4D::QVector4D(const QVector2D & vector);
extern void _ZN9QVector4DC2ERK9QVector2D(void* qthis, void* arg0); // 3
  // proto:  void QVector4D::QVector4D(float xpos, float ypos, float zpos, float wpos);
extern void _ZN9QVector4DC2Effff(void* qthis, float arg0, float arg1, float arg2, float arg3); // 1
  // proto:  void QVector4D::QVector4D(const QPointF & point);
extern void _ZN9QVector4DC2ERK7QPointF(void* qthis, void* arg0); // 1
  // proto:  void QVector4D::QVector4D(const QVector2D & vector, float zpos, float wpos);
extern void _ZN9QVector4DC2ERK9QVector2Dff(void* qthis, void* arg0, float arg1, float arg2); // 3
  // proto:  void QVector4D::QVector4D(const QVector3D & vector, float wpos);
extern void _ZN9QVector4DC2ERK9QVector3Df(void* qthis, void* arg0, float arg1); // 3
  // proto:  void QVector4D::QVector4D(const QVector3D & vector);
extern void _ZN9QVector4DC2ERK9QVector3D(void* qthis, void* arg0); // 3
  // proto:  void QVector4D::QVector4D(const QPoint & point);
extern void _ZN9QVector4DC2ERK6QPoint(void* qthis, void* arg0); // 1
  // proto: static float QVector4D::dotProduct(const QVector4D & v1, const QVector4D & v2);
extern void _ZN9QVector4D10dotProductERKS_S1_(void* arg0, void* arg1); // 4
  // proto:  void QVector4D::normalize();
extern void _ZN9QVector4D9normalizeEv(void* qthis); // 4
  // proto:  QPointF QVector4D::toPointF();
extern void _ZNK9QVector4D8toPointFEv(void* qthis); // 2
  // proto:  float QVector4D::lengthSquared();
extern void _ZNK9QVector4D13lengthSquaredEv(void* qthis); // 4
  // proto:  void QVector4D::setW(float w);
extern void _ZN9QVector4D4setWEf(void* qthis, float arg0); // 2
  // proto:  QVector4D QVector4D::normalized();
extern void _ZNK9QVector4D10normalizedEv(void* qthis); // 4
  // proto:  void QVector4D::setX(float x);
extern void _ZN9QVector4D4setXEf(void* qthis, float arg0); // 2
  // proto:  void QVector4D::setY(float y);
extern void _ZN9QVector4D4setYEf(void* qthis, float arg0); // 2
  // proto:  void QVector4D::setZ(float z);
extern void _ZN9QVector4D4setZEf(void* qthis, float arg0); // 2
  // proto:  QPoint QVector4D::toPoint();
extern void _ZNK9QVector4D7toPointEv(void* qthis); // 2
  // proto:  QVector2D QVector4D::toVector2DAffine();
extern void _ZNK9QVector4D16toVector2DAffineEv(void* qthis); // 4
  // proto:  QVector3D QVector4D::toVector3DAffine();
extern void _ZNK9QVector4D16toVector3DAffineEv(void* qthis); // 4
  // proto:  bool QVector4D::isNull();
extern void _ZNK9QVector4D6isNullEv(void* qthis); // 2
  // proto:  float QVector4D::length();
extern void _ZNK9QVector4D6lengthEv(void* qthis); // 4
  // proto:  float QVector4D::w();
extern void _ZNK9QVector4D1wEv(void* qthis); // 2
  // proto:  float QVector4D::y();
extern void _ZNK9QVector4D1yEv(void* qthis); // 2
  // proto:  float QVector4D::x();
extern void _ZNK9QVector4D1xEv(void* qthis); // 2
  // proto:  float QVector4D::z();
extern void _ZNK9QVector4D1zEv(void* qthis); // 2
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

// class sizeof(QVector4D)=16
type QVector4D struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// toVector2D()
func (this *QVector4D) toVector2D(args ...interface{}) () {
  // toVector2D()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector4D10toVector2DEv
    // invoke: QVector2D toVector2D()
    C._ZNK9QVector4D10toVector2DEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVector4D", "toVector2D", args)
  }

}

// toVector3D()
func (this *QVector4D) toVector3D(args ...interface{}) () {
  // toVector3D()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector4D10toVector3DEv
    // invoke: QVector3D toVector3D()
    C._ZNK9QVector4D10toVector3DEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVector4D", "toVector3D", args)
  }

}

// QVector4D()
func NewQVector4D(args ...interface{}) QVector4D {
  // QVector4D()
  // QVector4D(const class QVector2D &)
  // QVector4D(float, float, float, float)
  // QVector4D(const class QPointF &)
  // QVector4D(const class QVector2D &, float, float)
  // QVector4D(const class QVector3D &, float)
  // QVector4D(const class QVector3D &)
  // QVector4D(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.FloatTy(false) // "float"
  vtys[2][1] = qtrt.FloatTy(false) // "float"
  vtys[2][2] = qtrt.FloatTy(false) // "float"
  vtys[2][3] = qtrt.FloatTy(false) // "float"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"
  vtys[4][1] = qtrt.FloatTy(false) // "float"
  vtys[4][2] = qtrt.FloatTy(false) // "float"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[5][1] = qtrt.FloatTy(false) // "float"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector4DC1Ev
    // invoke: void QVector4D()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN9QVector4DC2Ev(qthis)
  case 1:
    // invoke: _ZN9QVector4DC1ERK9QVector2D
    // invoke: void QVector4D(const class QVector2D &)
    var arg0 = args[0].(QVector2D).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN9QVector4DC2ERK9QVector2D(qthis, arg0)
  case 2:
    // invoke: _ZN9QVector4DC1Effff
    // invoke: void QVector4D(float, float, float, float)
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
    C._ZN9QVector4DC2Effff(qthis, arg0, arg1, arg2, arg3)
  case 3:
    // invoke: _ZN9QVector4DC1ERK7QPointF
    // invoke: void QVector4D(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN9QVector4DC2ERK7QPointF(qthis, arg0)
  case 4:
    // invoke: _ZN9QVector4DC1ERK9QVector2Dff
    // invoke: void QVector4D(const class QVector2D &, float, float)
    var arg0 = args[0].(QVector2D).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN9QVector4DC2ERK9QVector2Dff(qthis, arg0, arg1, arg2)
  case 5:
    // invoke: _ZN9QVector4DC1ERK9QVector3Df
    // invoke: void QVector4D(const class QVector3D &, float)
    var arg0 = args[0].(QVector3D).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN9QVector4DC2ERK9QVector3Df(qthis, arg0, arg1)
  case 6:
    // invoke: _ZN9QVector4DC1ERK9QVector3D
    // invoke: void QVector4D(const class QVector3D &)
    var arg0 = args[0].(QVector3D).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN9QVector4DC2ERK9QVector3D(qthis, arg0)
  case 7:
    // invoke: _ZN9QVector4DC1ERK6QPoint
    // invoke: void QVector4D(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN9QVector4DC2ERK6QPoint(qthis, arg0)
  default:
    qtrt.ErrorResolve("QVector4D", "QVector4D", args)
  }

  return QVector4D{}
}

// dotProduct(const class QVector4D &, const class QVector4D &)
func (this *QVector4D) dotProduct_s(args ...interface{}) () {
  // dotProduct(const class QVector4D &, const class QVector4D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector4D{}) // "const QVector4D &"
  vtys[0][1] = reflect.TypeOf(QVector4D{}) // "const QVector4D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector4D10dotProductERKS_S1_
    // invoke: float dotProduct(const class QVector4D &, const class QVector4D &)
    var arg0 = args[0].(QVector4D).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector4D).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN9QVector4D10dotProductERKS_S1_(arg0, arg1)
  default:
    qtrt.ErrorResolve("QVector4D", "dotProduct", args)
  }

}

// normalize()
func (this *QVector4D) normalize(args ...interface{}) () {
  // normalize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector4D9normalizeEv
    // invoke: void normalize()
    C._ZN9QVector4D9normalizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVector4D", "normalize", args)
  }

}

// toPointF()
func (this *QVector4D) toPointF(args ...interface{}) () {
  // toPointF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector4D8toPointFEv
    // invoke: QPointF toPointF()
    C._ZNK9QVector4D8toPointFEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVector4D", "toPointF", args)
  }

}

// lengthSquared()
func (this *QVector4D) lengthSquared(args ...interface{}) () {
  // lengthSquared()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector4D13lengthSquaredEv
    // invoke: float lengthSquared()
    C._ZNK9QVector4D13lengthSquaredEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVector4D", "lengthSquared", args)
  }

}

// setW(float)
func (this *QVector4D) setW(args ...interface{}) () {
  // setW(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector4D4setWEf
    // invoke: void setW(float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    C._ZN9QVector4D4setWEf(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVector4D", "setW", args)
  }

}

// normalized()
func (this *QVector4D) normalized(args ...interface{}) () {
  // normalized()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector4D10normalizedEv
    // invoke: QVector4D normalized()
    C._ZNK9QVector4D10normalizedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVector4D", "normalized", args)
  }

}

// setX(float)
func (this *QVector4D) setX(args ...interface{}) () {
  // setX(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector4D4setXEf
    // invoke: void setX(float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    C._ZN9QVector4D4setXEf(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVector4D", "setX", args)
  }

}

// setY(float)
func (this *QVector4D) setY(args ...interface{}) () {
  // setY(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector4D4setYEf
    // invoke: void setY(float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    C._ZN9QVector4D4setYEf(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVector4D", "setY", args)
  }

}

// setZ(float)
func (this *QVector4D) setZ(args ...interface{}) () {
  // setZ(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector4D4setZEf
    // invoke: void setZ(float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    C._ZN9QVector4D4setZEf(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVector4D", "setZ", args)
  }

}

// toPoint()
func (this *QVector4D) toPoint(args ...interface{}) () {
  // toPoint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector4D7toPointEv
    // invoke: QPoint toPoint()
    C._ZNK9QVector4D7toPointEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVector4D", "toPoint", args)
  }

}

// toVector2DAffine()
func (this *QVector4D) toVector2DAffine(args ...interface{}) () {
  // toVector2DAffine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector4D16toVector2DAffineEv
    // invoke: QVector2D toVector2DAffine()
    C._ZNK9QVector4D16toVector2DAffineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVector4D", "toVector2DAffine", args)
  }

}

// toVector3DAffine()
func (this *QVector4D) toVector3DAffine(args ...interface{}) () {
  // toVector3DAffine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector4D16toVector3DAffineEv
    // invoke: QVector3D toVector3DAffine()
    C._ZNK9QVector4D16toVector3DAffineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVector4D", "toVector3DAffine", args)
  }

}

// isNull()
func (this *QVector4D) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector4D6isNullEv
    // invoke: bool isNull()
    C._ZNK9QVector4D6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVector4D", "isNull", args)
  }

}

// length()
func (this *QVector4D) length(args ...interface{}) () {
  // length()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector4D6lengthEv
    // invoke: float length()
    C._ZNK9QVector4D6lengthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVector4D", "length", args)
  }

}

// w()
func (this *QVector4D) w(args ...interface{}) () {
  // w()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector4D1wEv
    // invoke: float w()
    C._ZNK9QVector4D1wEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVector4D", "w", args)
  }

}

// y()
func (this *QVector4D) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector4D1yEv
    // invoke: float y()
    C._ZNK9QVector4D1yEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVector4D", "y", args)
  }

}

// x()
func (this *QVector4D) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector4D1xEv
    // invoke: float x()
    C._ZNK9QVector4D1xEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVector4D", "x", args)
  }

}

// z()
func (this *QVector4D) z(args ...interface{}) () {
  // z()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector4D1zEv
    // invoke: float z()
    C._ZNK9QVector4D1zEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVector4D", "z", args)
  }

}

// <= body block end

