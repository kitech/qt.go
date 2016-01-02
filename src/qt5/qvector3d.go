package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtGui/qvector3d.h
// dst-file: /src/gui/qvector3d.go
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
  // proto:  float QVector3D::x();
extern void _ZNK9QVector3D1xEv(void* qthis);
  // proto:  QPoint QVector3D::toPoint();
extern void _ZNK9QVector3D7toPointEv(void* qthis);
  // proto:  float QVector3D::distanceToLine(const QVector3D & point, const QVector3D & direction);
extern void _ZNK9QVector3D14distanceToLineERKS_S1_(void* qthis, void* arg0, void* arg1);
  // proto:  float QVector3D::y();
extern void _ZNK9QVector3D1yEv(void* qthis);
  // proto: static QVector3D QVector3D::normal(const QVector3D & v1, const QVector3D & v2);
extern void _ZN9QVector3D6normalERKS_S1_(void* arg0, void* arg1);
  // proto:  QPointF QVector3D::toPointF();
extern void _ZNK9QVector3D8toPointFEv(void* qthis);
  // proto:  void QVector3D::normalize();
extern void _ZN9QVector3D9normalizeEv(void* qthis);
  // proto:  QVector3D QVector3D::unproject(const QMatrix4x4 & modelView, const QMatrix4x4 & projection, const QRect & viewport);
extern void _ZNK9QVector3D9unprojectERK10QMatrix4x4S2_RK5QRect(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QVector3D::setY(float y);
extern void _ZN9QVector3D4setYEf(void* qthis, float arg0);
  // proto:  QVector3D QVector3D::project(const QMatrix4x4 & modelView, const QMatrix4x4 & projection, const QRect & viewport);
extern void _ZNK9QVector3D7projectERK10QMatrix4x4S2_RK5QRect(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto: static QVector3D QVector3D::crossProduct(const QVector3D & v1, const QVector3D & v2);
extern void _ZN9QVector3D12crossProductERKS_S1_(void* arg0, void* arg1);
  // proto:  float QVector3D::z();
extern void _ZNK9QVector3D1zEv(void* qthis);
  // proto:  void QVector3D::setZ(float z);
extern void _ZN9QVector3D4setZEf(void* qthis, float arg0);
  // proto:  QVector2D QVector3D::toVector2D();
extern void _ZNK9QVector3D10toVector2DEv(void* qthis);
  // proto:  float QVector3D::distanceToPlane(const QVector3D & plane1, const QVector3D & plane2, const QVector3D & plane3);
extern void _ZNK9QVector3D15distanceToPlaneERKS_S1_S1_(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QVector3D::QVector3D(const QPointF & point);
extern void* dector_ZN9QVector3DC1ERK7QPointF(void* arg0);
extern void _ZN9QVector3DC1ERK7QPointF(void* qthis, void* arg0);
  // proto:  float QVector3D::distanceToPoint(const QVector3D & point);
extern void _ZNK9QVector3D15distanceToPointERKS_(void* qthis, void* arg0);
  // proto:  void QVector3D::QVector3D(const QVector2D & vector);
extern void* dector_ZN9QVector3DC1ERK9QVector2D(void* arg0);
extern void _ZN9QVector3DC1ERK9QVector2D(void* qthis, void* arg0);
  // proto:  void QVector3D::QVector3D(const QPoint & point);
extern void* dector_ZN9QVector3DC1ERK6QPoint(void* arg0);
extern void _ZN9QVector3DC1ERK6QPoint(void* qthis, void* arg0);
  // proto:  float QVector3D::lengthSquared();
extern void _ZNK9QVector3D13lengthSquaredEv(void* qthis);
  // proto: static QVector3D QVector3D::normal(const QVector3D & v1, const QVector3D & v2, const QVector3D & v3);
extern void _ZN9QVector3D6normalERKS_S1_S1_(void* arg0, void* arg1, void* arg2);
  // proto:  float QVector3D::distanceToPlane(const QVector3D & plane, const QVector3D & normal);
extern void _ZNK9QVector3D15distanceToPlaneERKS_S1_(void* qthis, void* arg0, void* arg1);
  // proto:  QVector3D QVector3D::normalized();
extern void _ZNK9QVector3D10normalizedEv(void* qthis);
  // proto:  void QVector3D::QVector3D(const QVector4D & vector);
extern void* dector_ZN9QVector3DC1ERK9QVector4D(void* arg0);
extern void _ZN9QVector3DC1ERK9QVector4D(void* qthis, void* arg0);
  // proto:  bool QVector3D::isNull();
extern void _ZNK9QVector3D6isNullEv(void* qthis);
  // proto:  float QVector3D::length();
extern void _ZNK9QVector3D6lengthEv(void* qthis);
  // proto: static float QVector3D::dotProduct(const QVector3D & v1, const QVector3D & v2);
extern void _ZN9QVector3D10dotProductERKS_S1_(void* arg0, void* arg1);
  // proto:  void QVector3D::QVector3D();
extern void* dector_ZN9QVector3DC1Ev();
extern void _ZN9QVector3DC1Ev(void* qthis);
  // proto:  void QVector3D::QVector3D(float xpos, float ypos, float zpos);
extern void* dector_ZN9QVector3DC1Efff(float arg0, float arg1, float arg2);
extern void _ZN9QVector3DC1Efff(void* qthis, float arg0, float arg1, float arg2);
  // proto:  QVector4D QVector3D::toVector4D();
extern void _ZNK9QVector3D10toVector4DEv(void* qthis);
  // proto:  void QVector3D::QVector3D(const QVector2D & vector, float zpos);
extern void* dector_ZN9QVector3DC1ERK9QVector2Df(void* arg0, float arg1);
extern void _ZN9QVector3DC1ERK9QVector2Df(void* qthis, void* arg0, float arg1);
  // proto:  void QVector3D::setX(float x);
extern void _ZN9QVector3D4setXEf(void* qthis, float arg0);
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

// class sizeof(QVector3D)=12
type QVector3D struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  float QVector3D::x();
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

  // proto:  QPoint QVector3D::toPoint();
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

  // proto:  float QVector3D::distanceToLine(const QVector3D & point, const QVector3D & direction);
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

  // proto:  float QVector3D::y();
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

  // proto: static QVector3D QVector3D::normal(const QVector3D & v1, const QVector3D & v2);
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

  // proto:  QPointF QVector3D::toPointF();
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

  // proto:  void QVector3D::normalize();
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

  // proto:  QVector3D QVector3D::unproject(const QMatrix4x4 & modelView, const QMatrix4x4 & projection, const QRect & viewport);
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

  // proto:  void QVector3D::setY(float y);
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

  // proto:  QVector3D QVector3D::project(const QMatrix4x4 & modelView, const QMatrix4x4 & projection, const QRect & viewport);
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

  // proto: static QVector3D QVector3D::crossProduct(const QVector3D & v1, const QVector3D & v2);
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

  // proto:  float QVector3D::z();
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

  // proto:  void QVector3D::setZ(float z);
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

  // proto:  QVector2D QVector3D::toVector2D();
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

  // proto:  float QVector3D::distanceToPlane(const QVector3D & plane1, const QVector3D & plane2, const QVector3D & plane3);
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

  // proto:  void QVector3D::QVector3D(const QPointF & point);
func NewQVector3D(args ...interface{}) QVector3D {
  return QVector3D{}
}

  // proto:  float QVector3D::distanceToPoint(const QVector3D & point);
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

  // proto:  float QVector3D::lengthSquared();
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

  // proto:  QVector3D QVector3D::normalized();
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

  // proto:  bool QVector3D::isNull();
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

  // proto:  float QVector3D::length();
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

  // proto: static float QVector3D::dotProduct(const QVector3D & v1, const QVector3D & v2);
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

  // proto:  QVector4D QVector3D::toVector4D();
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

  // proto:  void QVector3D::setX(float x);
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

