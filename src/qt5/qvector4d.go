package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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
extern void* C_ZNK9QVector4D10toVector2DEv(void* qthis); // 4
  // proto:  QVector3D QVector4D::toVector3D();
extern void* C_ZNK9QVector4D10toVector3DEv(void* qthis); // 4
  // proto:  void QVector4D::QVector4D();
extern void* C_ZN9QVector4DC2Ev(); // 1
  // proto:  void QVector4D::QVector4D(const QVector2D & vector);
extern void* C_ZN9QVector4DC2ERK9QVector2D(void* arg0); // 3
  // proto:  void QVector4D::QVector4D(float xpos, float ypos, float zpos, float wpos);
extern void* C_ZN9QVector4DC2Effff(float arg0, float arg1, float arg2, float arg3); // 1
  // proto:  void QVector4D::QVector4D(const QPointF & point);
extern void* C_ZN9QVector4DC2ERK7QPointF(void* arg0); // 1
  // proto:  void QVector4D::QVector4D(const QVector2D & vector, float zpos, float wpos);
extern void* C_ZN9QVector4DC2ERK9QVector2Dff(void* arg0, float arg1, float arg2); // 3
  // proto:  void QVector4D::QVector4D(const QVector3D & vector, float wpos);
extern void* C_ZN9QVector4DC2ERK9QVector3Df(void* arg0, float arg1); // 3
  // proto:  void QVector4D::QVector4D(const QVector3D & vector);
extern void* C_ZN9QVector4DC2ERK9QVector3D(void* arg0); // 3
  // proto:  void QVector4D::QVector4D(const QPoint & point);
extern void* C_ZN9QVector4DC2ERK6QPoint(void* arg0); // 1
  // proto: static float QVector4D::dotProduct(const QVector4D & v1, const QVector4D & v2);
extern float C_ZN9QVector4D10dotProductERKS_S1_(void* arg0, void* arg1); // 4
  // proto:  void QVector4D::normalize();
extern void C_ZN9QVector4D9normalizeEv(void* qthis); // 4
  // proto:  QPointF QVector4D::toPointF();
extern void* C_ZNK9QVector4D8toPointFEv(void* qthis); // 2
  // proto:  float QVector4D::lengthSquared();
extern float C_ZNK9QVector4D13lengthSquaredEv(void* qthis); // 4
  // proto:  void QVector4D::setW(float w);
extern void C_ZN9QVector4D4setWEf(void* qthis, float arg0); // 2
  // proto:  QVector4D QVector4D::normalized();
extern void* C_ZNK9QVector4D10normalizedEv(void* qthis); // 4
  // proto:  void QVector4D::setX(float x);
extern void C_ZN9QVector4D4setXEf(void* qthis, float arg0); // 2
  // proto:  void QVector4D::setY(float y);
extern void C_ZN9QVector4D4setYEf(void* qthis, float arg0); // 2
  // proto:  void QVector4D::setZ(float z);
extern void C_ZN9QVector4D4setZEf(void* qthis, float arg0); // 2
  // proto:  QPoint QVector4D::toPoint();
extern void* C_ZNK9QVector4D7toPointEv(void* qthis); // 2
  // proto:  QVector2D QVector4D::toVector2DAffine();
extern void* C_ZNK9QVector4D16toVector2DAffineEv(void* qthis); // 4
  // proto:  QVector3D QVector4D::toVector3DAffine();
extern void* C_ZNK9QVector4D16toVector3DAffineEv(void* qthis); // 4
  // proto:  bool QVector4D::isNull();
extern bool C_ZNK9QVector4D6isNullEv(void* qthis); // 2
  // proto:  float QVector4D::length();
extern float C_ZNK9QVector4D6lengthEv(void* qthis); // 4
  // proto:  float QVector4D::w();
extern float C_ZNK9QVector4D1wEv(void* qthis); // 2
  // proto:  float QVector4D::y();
extern void C_ZNK9QVector4D1yEv(void* qthis); // 2
  // proto:  float QVector4D::x();
extern void C_ZNK9QVector4D1xEv(void* qthis); // 2
  // proto:  float QVector4D::z();
extern float C_ZNK9QVector4D1zEv(void* qthis); // 2
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// toVector2D()
func (this *QVector4D) Tovector2D(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QVector4D10toVector2DEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVector2D{}) // "QVector2D"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector4D", "toVector2D", args)
  }

  return
}

// toVector3D()
func (this *QVector4D) Tovector3D(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QVector4D10toVector3DEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVector3D{}) // "QVector3D"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector4D", "toVector3D", args)
  }

  return
}

// QVector4D()
func NewQVector4D(args ...interface{}) *QVector4D {
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
    qthis = C.C_ZN9QVector4DC2Ev()
    return &QVector4D{Qclsinst:qthis}
  case 1:
    // invoke: _ZN9QVector4DC1ERK9QVector2D
    // invoke: void QVector4D(const class QVector2D &)
    var arg0 = args[0].(*QVector2D).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QVector4DC2ERK9QVector2D(arg0)
    return &QVector4D{Qclsinst:qthis}
  case 2:
    // invoke: _ZN9QVector4DC1Effff
    // invoke: void QVector4D(float, float, float, float)
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
    qthis = C.C_ZN9QVector4DC2Effff(arg0, arg1, arg2, arg3)
    return &QVector4D{Qclsinst:qthis}
  case 3:
    // invoke: _ZN9QVector4DC1ERK7QPointF
    // invoke: void QVector4D(const class QPointF &)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QVector4DC2ERK7QPointF(arg0)
    return &QVector4D{Qclsinst:qthis}
  case 4:
    // invoke: _ZN9QVector4DC1ERK9QVector2Dff
    // invoke: void QVector4D(const class QVector2D &, float, float)
    var arg0 = args[0].(*QVector2D).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(qtrt.PrimConv(args[2], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QVector4DC2ERK9QVector2Dff(arg0, arg1, arg2)
    return &QVector4D{Qclsinst:qthis}
  case 5:
    // invoke: _ZN9QVector4DC1ERK9QVector3Df
    // invoke: void QVector4D(const class QVector3D &, float)
    var arg0 = args[0].(*QVector3D).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QVector4DC2ERK9QVector3Df(arg0, arg1)
    return &QVector4D{Qclsinst:qthis}
  case 6:
    // invoke: _ZN9QVector4DC1ERK9QVector3D
    // invoke: void QVector4D(const class QVector3D &)
    var arg0 = args[0].(*QVector3D).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QVector4DC2ERK9QVector3D(arg0)
    return &QVector4D{Qclsinst:qthis}
  case 7:
    // invoke: _ZN9QVector4DC1ERK6QPoint
    // invoke: void QVector4D(const class QPoint &)
    var arg0 = args[0].(*QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QVector4DC2ERK6QPoint(arg0)
    return &QVector4D{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QVector4D", "QVector4D", args)
  }

  return nil // QVector4D{Qclsinst:qthis}
}

// dotProduct(const class QVector4D &, const class QVector4D &)
func (this *QVector4D) Dotproduct_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QVector4D).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVector4D).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN9QVector4D10dotProductERKS_S1_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector4D", "dotProduct", args)
  }

  return
}

// normalize()
func (this *QVector4D) Normalize(args ...interface{}) () {
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
    C.C_ZN9QVector4D9normalizeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QVector4D", "normalize", args)
  }

  return
}

// toPointF()
func (this *QVector4D) Topointf(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QVector4D8toPointFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector4D", "toPointF", args)
  }

  return
}

// lengthSquared()
func (this *QVector4D) Lengthsquared(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QVector4D13lengthSquaredEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector4D", "lengthSquared", args)
  }

  return
}

// setW(float)
func (this *QVector4D) Setw(args ...interface{}) () {
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
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QVector4D4setWEf(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVector4D", "setW", args)
  }

  return
}

// normalized()
func (this *QVector4D) Normalized(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QVector4D10normalizedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVector4D{}) // "QVector4D"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector4D", "normalized", args)
  }

  return
}

// setX(float)
func (this *QVector4D) Setx(args ...interface{}) () {
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
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QVector4D4setXEf(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVector4D", "setX", args)
  }

  return
}

// setY(float)
func (this *QVector4D) Sety(args ...interface{}) () {
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
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QVector4D4setYEf(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVector4D", "setY", args)
  }

  return
}

// setZ(float)
func (this *QVector4D) Setz(args ...interface{}) () {
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
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QVector4D4setZEf(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVector4D", "setZ", args)
  }

  return
}

// toPoint()
func (this *QVector4D) Topoint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QVector4D7toPointEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector4D", "toPoint", args)
  }

  return
}

// toVector2DAffine()
func (this *QVector4D) Tovector2Daffine(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QVector4D16toVector2DAffineEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVector2D{}) // "QVector2D"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector4D", "toVector2DAffine", args)
  }

  return
}

// toVector3DAffine()
func (this *QVector4D) Tovector3Daffine(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QVector4D16toVector3DAffineEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVector3D{}) // "QVector3D"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector4D", "toVector3DAffine", args)
  }

  return
}

// isNull()
func (this *QVector4D) Isnull(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QVector4D6isNullEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector4D", "isNull", args)
  }

  return
}

// length()
func (this *QVector4D) Length(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QVector4D6lengthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector4D", "length", args)
  }

  return
}

// w()
func (this *QVector4D) W(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QVector4D1wEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector4D", "w", args)
  }

  return
}

// y()
func (this *QVector4D) Y(args ...interface{}) () {
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
    C.C_ZNK9QVector4D1yEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QVector4D", "y", args)
  }

  return
}

// x()
func (this *QVector4D) X(args ...interface{}) () {
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
    C.C_ZNK9QVector4D1xEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QVector4D", "x", args)
  }

  return
}

// z()
func (this *QVector4D) Z(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QVector4D1zEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector4D", "z", args)
  }

  return
}

// <= body block end

