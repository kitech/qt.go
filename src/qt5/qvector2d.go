package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtGui/qvector2d.h
// dst-file: /src/gui/qvector2d.go
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
  // proto:  QVector4D QVector2D::toVector4D();
extern void* C_ZNK9QVector2D10toVector4DEv(void* qthis); // 4
  // proto:  void QVector2D::normalize();
extern void C_ZN9QVector2D9normalizeEv(void* qthis); // 4
  // proto: static float QVector2D::dotProduct(const QVector2D & v1, const QVector2D & v2);
extern float C_ZN9QVector2D10dotProductERKS_S1_(void* arg0, void* arg1); // 4
  // proto:  void QVector2D::QVector2D(const QPoint & point);
extern void* C_ZN9QVector2DC2ERK6QPoint(void* arg0); // 1
  // proto:  void QVector2D::QVector2D(float xpos, float ypos);
extern void* C_ZN9QVector2DC2Eff(float arg0, float arg1); // 1
  // proto:  void QVector2D::QVector2D(const QVector4D & vector);
extern void* C_ZN9QVector2DC2ERK9QVector4D(void* arg0); // 3
  // proto:  void QVector2D::QVector2D();
extern void* C_ZN9QVector2DC2Ev(); // 1
  // proto:  void QVector2D::QVector2D(const QVector3D & vector);
extern void* C_ZN9QVector2DC2ERK9QVector3D(void* arg0); // 3
  // proto:  void QVector2D::QVector2D(const QPointF & point);
extern void* C_ZN9QVector2DC2ERK7QPointF(void* arg0); // 1
  // proto:  QPointF QVector2D::toPointF();
extern void* C_ZNK9QVector2D8toPointFEv(void* qthis); // 2
  // proto:  QVector3D QVector2D::toVector3D();
extern void* C_ZNK9QVector2D10toVector3DEv(void* qthis); // 4
  // proto:  float QVector2D::lengthSquared();
extern float C_ZNK9QVector2D13lengthSquaredEv(void* qthis); // 4
  // proto:  QVector2D QVector2D::normalized();
extern void* C_ZNK9QVector2D10normalizedEv(void* qthis); // 4
  // proto:  void QVector2D::setX(float x);
extern void C_ZN9QVector2D4setXEf(void* qthis, float arg0); // 2
  // proto:  void QVector2D::setY(float y);
extern void C_ZN9QVector2D4setYEf(void* qthis, float arg0); // 2
  // proto:  QPoint QVector2D::toPoint();
extern void* C_ZNK9QVector2D7toPointEv(void* qthis); // 2
  // proto:  float QVector2D::distanceToLine(const QVector2D & point, const QVector2D & direction);
extern float C_ZNK9QVector2D14distanceToLineERKS_S1_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QVector2D::isNull();
extern bool C_ZNK9QVector2D6isNullEv(void* qthis); // 2
  // proto:  float QVector2D::length();
extern float C_ZNK9QVector2D6lengthEv(void* qthis); // 4
  // proto:  float QVector2D::distanceToPoint(const QVector2D & point);
extern float C_ZNK9QVector2D15distanceToPointERKS_(void* qthis, void* arg0); // 4
  // proto:  float QVector2D::y();
extern void C_ZNK9QVector2D1yEv(void* qthis); // 2
  // proto:  float QVector2D::x();
extern void C_ZNK9QVector2D1xEv(void* qthis); // 2
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

// class sizeof(QVector2D)=8
type QVector2D struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// toVector4D()
func (this *QVector2D) Tovector4D(args ...interface{}) (ret interface{}) {
  // toVector4D()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D10toVector4DEv
    // invoke: QVector4D toVector4D()
    var ret0 = C.C_ZNK9QVector2D10toVector4DEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVector4D{}) // "QVector4D"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector2D", "toVector4D", args)
  }

  return
}

// normalize()
func (this *QVector2D) Normalize(args ...interface{}) () {
  // normalize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector2D9normalizeEv
    // invoke: void normalize()
    C.C_ZN9QVector2D9normalizeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QVector2D", "normalize", args)
  }

  return
}

// dotProduct(const class QVector2D &, const class QVector2D &)
func (this *QVector2D) Dotproduct_S(args ...interface{}) (ret interface{}) {
  // dotProduct(const class QVector2D &, const class QVector2D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"
  vtys[0][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector2D10dotProductERKS_S1_
    // invoke: float dotProduct(const class QVector2D &, const class QVector2D &)
    var arg0 = args[0].(*QVector2D).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVector2D).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN9QVector2D10dotProductERKS_S1_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector2D", "dotProduct", args)
  }

  return
}

// QVector2D(const class QPoint &)
func NewQVector2D(args ...interface{}) *QVector2D {
  // QVector2D(const class QPoint &)
  // QVector2D(float, float)
  // QVector2D(const class QVector4D &)
  // QVector2D()
  // QVector2D(const class QVector3D &)
  // QVector2D(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.FloatTy(false) // "float"
  vtys[1][1] = qtrt.FloatTy(false) // "float"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QVector4D{}) // "const QVector4D &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector2DC1ERK6QPoint
    // invoke: void QVector2D(const class QPoint &)
    var arg0 = args[0].(*QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QVector2DC2ERK6QPoint(arg0)
    return &QVector2D{Qclsinst:qthis}
  case 1:
    // invoke: _ZN9QVector2DC1Eff
    // invoke: void QVector2D(float, float)
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QVector2DC2Eff(arg0, arg1)
    return &QVector2D{Qclsinst:qthis}
  case 2:
    // invoke: _ZN9QVector2DC1ERK9QVector4D
    // invoke: void QVector2D(const class QVector4D &)
    var arg0 = args[0].(*QVector4D).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QVector2DC2ERK9QVector4D(arg0)
    return &QVector2D{Qclsinst:qthis}
  case 3:
    // invoke: _ZN9QVector2DC1Ev
    // invoke: void QVector2D()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QVector2DC2Ev()
    return &QVector2D{Qclsinst:qthis}
  case 4:
    // invoke: _ZN9QVector2DC1ERK9QVector3D
    // invoke: void QVector2D(const class QVector3D &)
    var arg0 = args[0].(*QVector3D).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QVector2DC2ERK9QVector3D(arg0)
    return &QVector2D{Qclsinst:qthis}
  case 5:
    // invoke: _ZN9QVector2DC1ERK7QPointF
    // invoke: void QVector2D(const class QPointF &)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QVector2DC2ERK7QPointF(arg0)
    return &QVector2D{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QVector2D", "QVector2D", args)
  }

  return nil // QVector2D{Qclsinst:qthis}
}

// toPointF()
func (this *QVector2D) Topointf(args ...interface{}) (ret interface{}) {
  // toPointF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D8toPointFEv
    // invoke: QPointF toPointF()
    var ret0 = C.C_ZNK9QVector2D8toPointFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector2D", "toPointF", args)
  }

  return
}

// toVector3D()
func (this *QVector2D) Tovector3D(args ...interface{}) (ret interface{}) {
  // toVector3D()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D10toVector3DEv
    // invoke: QVector3D toVector3D()
    var ret0 = C.C_ZNK9QVector2D10toVector3DEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVector3D{}) // "QVector3D"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector2D", "toVector3D", args)
  }

  return
}

// lengthSquared()
func (this *QVector2D) Lengthsquared(args ...interface{}) (ret interface{}) {
  // lengthSquared()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D13lengthSquaredEv
    // invoke: float lengthSquared()
    var ret0 = C.C_ZNK9QVector2D13lengthSquaredEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector2D", "lengthSquared", args)
  }

  return
}

// normalized()
func (this *QVector2D) Normalized(args ...interface{}) (ret interface{}) {
  // normalized()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D10normalizedEv
    // invoke: QVector2D normalized()
    var ret0 = C.C_ZNK9QVector2D10normalizedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVector2D{}) // "QVector2D"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector2D", "normalized", args)
  }

  return
}

// setX(float)
func (this *QVector2D) Setx(args ...interface{}) () {
  // setX(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector2D4setXEf
    // invoke: void setX(float)
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QVector2D4setXEf(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVector2D", "setX", args)
  }

  return
}

// setY(float)
func (this *QVector2D) Sety(args ...interface{}) () {
  // setY(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector2D4setYEf
    // invoke: void setY(float)
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QVector2D4setYEf(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVector2D", "setY", args)
  }

  return
}

// toPoint()
func (this *QVector2D) Topoint(args ...interface{}) (ret interface{}) {
  // toPoint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D7toPointEv
    // invoke: QPoint toPoint()
    var ret0 = C.C_ZNK9QVector2D7toPointEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector2D", "toPoint", args)
  }

  return
}

// distanceToLine(const class QVector2D &, const class QVector2D &)
func (this *QVector2D) Distancetoline(args ...interface{}) (ret interface{}) {
  // distanceToLine(const class QVector2D &, const class QVector2D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"
  vtys[0][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D14distanceToLineERKS_S1_
    // invoke: float distanceToLine(const class QVector2D &, const class QVector2D &)
    var arg0 = args[0].(*QVector2D).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVector2D).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK9QVector2D14distanceToLineERKS_S1_(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector2D", "distanceToLine", args)
  }

  return
}

// isNull()
func (this *QVector2D) Isnull(args ...interface{}) (ret interface{}) {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D6isNullEv
    // invoke: bool isNull()
    var ret0 = C.C_ZNK9QVector2D6isNullEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector2D", "isNull", args)
  }

  return
}

// length()
func (this *QVector2D) Length(args ...interface{}) (ret interface{}) {
  // length()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D6lengthEv
    // invoke: float length()
    var ret0 = C.C_ZNK9QVector2D6lengthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector2D", "length", args)
  }

  return
}

// distanceToPoint(const class QVector2D &)
func (this *QVector2D) Distancetopoint(args ...interface{}) (ret interface{}) {
  // distanceToPoint(const class QVector2D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D15distanceToPointERKS_
    // invoke: float distanceToPoint(const class QVector2D &)
    var arg0 = args[0].(*QVector2D).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QVector2D15distanceToPointERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector2D", "distanceToPoint", args)
  }

  return
}

// y()
func (this *QVector2D) Y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D1yEv
    // invoke: float y()
    C.C_ZNK9QVector2D1yEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QVector2D", "y", args)
  }

  return
}

// x()
func (this *QVector2D) X(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D1xEv
    // invoke: float x()
    C.C_ZNK9QVector2D1xEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QVector2D", "x", args)
  }

  return
}

// <= body block end

