package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
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
  // proto:  QVector4D QVector4D::normalized();
extern void _ZNK9QVector4D10normalizedEv(void* qthis);
  // proto:  void QVector4D::setW(float w);
extern void _ZN9QVector4D4setWEf(void* qthis, float arg0);
  // proto:  void QVector4D::QVector4D(const QVector2D & vector, float zpos, float wpos);
extern void* dector_ZN9QVector4DC1ERK9QVector2Dff(void* arg0, float arg1, float arg2);
extern void _ZN9QVector4DC1ERK9QVector2Dff(void* qthis, void* arg0, float arg1, float arg2);
  // proto:  QPointF QVector4D::toPointF();
extern void _ZNK9QVector4D8toPointFEv(void* qthis);
  // proto:  float QVector4D::y();
extern void _ZNK9QVector4D1yEv(void* qthis);
  // proto:  QVector2D QVector4D::toVector2D();
extern void _ZNK9QVector4D10toVector2DEv(void* qthis);
  // proto:  void QVector4D::setZ(float z);
extern void _ZN9QVector4D4setZEf(void* qthis, float arg0);
  // proto:  void QVector4D::QVector4D(const QVector2D & vector);
extern void* dector_ZN9QVector4DC1ERK9QVector2D(void* arg0);
extern void _ZN9QVector4DC1ERK9QVector2D(void* qthis, void* arg0);
  // proto:  void QVector4D::normalize();
extern void _ZN9QVector4D9normalizeEv(void* qthis);
  // proto:  void QVector4D::QVector4D(float xpos, float ypos, float zpos, float wpos);
extern void* dector_ZN9QVector4DC1Effff(float arg0, float arg1, float arg2, float arg3);
extern void _ZN9QVector4DC1Effff(void* qthis, float arg0, float arg1, float arg2, float arg3);
  // proto:  void QVector4D::QVector4D(const QVector3D & vector, float wpos);
extern void* dector_ZN9QVector4DC1ERK9QVector3Df(void* arg0, float arg1);
extern void _ZN9QVector4DC1ERK9QVector3Df(void* qthis, void* arg0, float arg1);
  // proto:  void QVector4D::QVector4D(const QPointF & point);
extern void* dector_ZN9QVector4DC1ERK7QPointF(void* arg0);
extern void _ZN9QVector4DC1ERK7QPointF(void* qthis, void* arg0);
  // proto:  float QVector4D::z();
extern void _ZNK9QVector4D1zEv(void* qthis);
  // proto:  void QVector4D::QVector4D();
extern void* dector_ZN9QVector4DC1Ev();
extern void _ZN9QVector4DC1Ev(void* qthis);
  // proto:  void QVector4D::setX(float x);
extern void _ZN9QVector4D4setXEf(void* qthis, float arg0);
  // proto:  void QVector4D::setY(float y);
extern void _ZN9QVector4D4setYEf(void* qthis, float arg0);
  // proto:  void QVector4D::QVector4D(const QPoint & point);
extern void* dector_ZN9QVector4DC1ERK6QPoint(void* arg0);
extern void _ZN9QVector4DC1ERK6QPoint(void* qthis, void* arg0);
  // proto:  QVector3D QVector4D::toVector3D();
extern void _ZNK9QVector4D10toVector3DEv(void* qthis);
  // proto:  float QVector4D::x();
extern void _ZNK9QVector4D1xEv(void* qthis);
  // proto:  QVector2D QVector4D::toVector2DAffine();
extern void _ZNK9QVector4D16toVector2DAffineEv(void* qthis);
  // proto:  float QVector4D::length();
extern void _ZNK9QVector4D6lengthEv(void* qthis);
  // proto:  void QVector4D::QVector4D(const QVector3D & vector);
extern void* dector_ZN9QVector4DC1ERK9QVector3D(void* arg0);
extern void _ZN9QVector4DC1ERK9QVector3D(void* qthis, void* arg0);
  // proto: static float QVector4D::dotProduct(const QVector4D & v1, const QVector4D & v2);
extern void _ZN9QVector4D10dotProductERKS_S1_(void* arg0, void* arg1);
  // proto:  bool QVector4D::isNull();
extern void _ZNK9QVector4D6isNullEv(void* qthis);
  // proto:  float QVector4D::lengthSquared();
extern void _ZNK9QVector4D13lengthSquaredEv(void* qthis);
  // proto:  QVector3D QVector4D::toVector3DAffine();
extern void _ZNK9QVector4D16toVector3DAffineEv(void* qthis);
  // proto:  QPoint QVector4D::toPoint();
extern void _ZNK9QVector4D7toPointEv(void* qthis);
  // proto:  float QVector4D::w();
extern void _ZNK9QVector4D1wEv(void* qthis);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QVector4D QVector4D::normalized();
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
  default:
    qtrt.ErrorResolve("QVector4D", "normalized", args)
  }

}

  // proto:  void QVector4D::setW(float w);
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
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QVector4D", "setW", args)
  }

}

  // proto:  void QVector4D::QVector4D(const QVector2D & vector, float zpos, float wpos);
func NewQVector4D(args ...interface{}) QVector4D {
  return QVector4D{}
}

  // proto:  QPointF QVector4D::toPointF();
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
  default:
    qtrt.ErrorResolve("QVector4D", "toPointF", args)
  }

}

  // proto:  float QVector4D::y();
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
  default:
    qtrt.ErrorResolve("QVector4D", "y", args)
  }

}

  // proto:  QVector2D QVector4D::toVector2D();
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
  default:
    qtrt.ErrorResolve("QVector4D", "toVector2D", args)
  }

}

  // proto:  void QVector4D::setZ(float z);
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
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QVector4D", "setZ", args)
  }

}

  // proto:  void QVector4D::normalize();
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
  default:
    qtrt.ErrorResolve("QVector4D", "normalize", args)
  }

}

  // proto:  float QVector4D::z();
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
  default:
    qtrt.ErrorResolve("QVector4D", "z", args)
  }

}

  // proto:  void QVector4D::setX(float x);
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
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QVector4D", "setX", args)
  }

}

  // proto:  void QVector4D::setY(float y);
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
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QVector4D", "setY", args)
  }

}

  // proto:  QVector3D QVector4D::toVector3D();
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
  default:
    qtrt.ErrorResolve("QVector4D", "toVector3D", args)
  }

}

  // proto:  float QVector4D::x();
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
  default:
    qtrt.ErrorResolve("QVector4D", "x", args)
  }

}

  // proto:  QVector2D QVector4D::toVector2DAffine();
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
  default:
    qtrt.ErrorResolve("QVector4D", "toVector2DAffine", args)
  }

}

  // proto:  float QVector4D::length();
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
  default:
    qtrt.ErrorResolve("QVector4D", "length", args)
  }

}

  // proto: static float QVector4D::dotProduct(const QVector4D & v1, const QVector4D & v2);
func (this *QVector4D) dotProduct_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QVector4D", "dotProduct", args)
  }

}

  // proto:  bool QVector4D::isNull();
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
  default:
    qtrt.ErrorResolve("QVector4D", "isNull", args)
  }

}

  // proto:  float QVector4D::lengthSquared();
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
  default:
    qtrt.ErrorResolve("QVector4D", "lengthSquared", args)
  }

}

  // proto:  QVector3D QVector4D::toVector3DAffine();
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
  default:
    qtrt.ErrorResolve("QVector4D", "toVector3DAffine", args)
  }

}

  // proto:  QPoint QVector4D::toPoint();
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
  default:
    qtrt.ErrorResolve("QVector4D", "toPoint", args)
  }

}

  // proto:  float QVector4D::w();
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
  default:
    qtrt.ErrorResolve("QVector4D", "w", args)
  }

}

// <= body block end

