package qtgui
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto: static QVector3D QVector3D::crossProduct(const QVector3D & v1, const QVector3D & v2);
extern void* C_ZN9QVector3D12crossProductERKS_S1_(void* arg0, void* arg1); // 4
  // proto:  QVector2D QVector3D::toVector2D();
extern void* C_ZNK9QVector3D10toVector2DEv(void* qthis); // 4
  // proto:  float QVector3D::distanceToPlane(const QVector3D & plane, const QVector3D & normal);
extern float C_ZNK9QVector3D15distanceToPlaneERKS_S1_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  float QVector3D::distanceToPlane(const QVector3D & plane1, const QVector3D & plane2, const QVector3D & plane3);
extern float C_ZNK9QVector3D15distanceToPlaneERKS_S1_S1_(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  QVector4D QVector3D::toVector4D();
extern void* C_ZNK9QVector3D10toVector4DEv(void* qthis); // 4
  // proto: static float QVector3D::dotProduct(const QVector3D & v1, const QVector3D & v2);
extern float C_ZN9QVector3D10dotProductERKS_S1_(void* arg0, void* arg1); // 4
  // proto:  void QVector3D::normalize();
extern void C_ZN9QVector3D9normalizeEv(void* qthis); // 4
  // proto:  QPointF QVector3D::toPointF();
extern void* C_ZNK9QVector3D8toPointFEv(void* qthis); // 2
  // proto:  float QVector3D::lengthSquared();
extern float C_ZNK9QVector3D13lengthSquaredEv(void* qthis); // 4
  // proto: static QVector3D QVector3D::normal(const QVector3D & v1, const QVector3D & v2, const QVector3D & v3);
extern void* C_ZN9QVector3D6normalERKS_S1_S1_(void* arg0, void* arg1, void* arg2); // 4
  // proto: static QVector3D QVector3D::normal(const QVector3D & v1, const QVector3D & v2);
extern void* C_ZN9QVector3D6normalERKS_S1_(void* arg0, void* arg1); // 4
  // proto:  void QVector3D::QVector3D(const QPointF & point);
extern void* C_ZN9QVector3DC2ERK7QPointF(void* arg0); // 1
  // proto:  void QVector3D::QVector3D();
extern void* C_ZN9QVector3DC2Ev(); // 1
  // proto:  void QVector3D::QVector3D(const QVector2D & vector);
extern void* C_ZN9QVector3DC2ERK9QVector2D(void* arg0); // 3
  // proto:  void QVector3D::QVector3D(const QPoint & point);
extern void* C_ZN9QVector3DC2ERK6QPoint(void* arg0); // 1
  // proto:  void QVector3D::QVector3D(float xpos, float ypos, float zpos);
extern void* C_ZN9QVector3DC2Efff(float arg0, float arg1, float arg2); // 1
  // proto:  void QVector3D::QVector3D(const QVector4D & vector);
extern void* C_ZN9QVector3DC2ERK9QVector4D(void* arg0); // 3
  // proto:  void QVector3D::QVector3D(const QVector2D & vector, float zpos);
extern void* C_ZN9QVector3DC2ERK9QVector2Df(void* arg0, float arg1); // 3
  // proto:  QVector3D QVector3D::normalized();
extern void* C_ZNK9QVector3D10normalizedEv(void* qthis); // 4
  // proto:  float QVector3D::x();
extern void C_ZNK9QVector3D1xEv(void* qthis); // 2
  // proto:  void QVector3D::setX(float x);
extern void C_ZN9QVector3D4setXEf(void* qthis, float arg0); // 2
  // proto:  void QVector3D::setY(float y);
extern void C_ZN9QVector3D4setYEf(void* qthis, float arg0); // 2
  // proto:  void QVector3D::setZ(float z);
extern void C_ZN9QVector3D4setZEf(void* qthis, float arg0); // 2
  // proto:  QPoint QVector3D::toPoint();
extern void* C_ZNK9QVector3D7toPointEv(void* qthis); // 2
  // proto:  QVector3D QVector3D::unproject(const QMatrix4x4 & modelView, const QMatrix4x4 & projection, const QRect & viewport);
extern void* C_ZNK9QVector3D9unprojectERK10QMatrix4x4S2_RK5QRect(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  float QVector3D::distanceToLine(const QVector3D & point, const QVector3D & direction);
extern float C_ZNK9QVector3D14distanceToLineERKS_S1_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QVector3D QVector3D::project(const QMatrix4x4 & modelView, const QMatrix4x4 & projection, const QRect & viewport);
extern void* C_ZNK9QVector3D7projectERK10QMatrix4x4S2_RK5QRect(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  bool QVector3D::isNull();
extern bool C_ZNK9QVector3D6isNullEv(void* qthis); // 2
  // proto:  float QVector3D::length();
extern float C_ZNK9QVector3D6lengthEv(void* qthis); // 4
  // proto:  float QVector3D::distanceToPoint(const QVector3D & point);
extern float C_ZNK9QVector3D15distanceToPointERKS_(void* qthis, void* arg0); // 4
  // proto:  float QVector3D::y();
extern void C_ZNK9QVector3D1yEv(void* qthis); // 2
  // proto:  float QVector3D::z();
extern float C_ZNK9QVector3D1zEv(void* qthis); // 2
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
}

// class sizeof(QVector3D)=12
type QVector3D struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// crossProduct(const class QVector3D &, const class QVector3D &)
func (this *QVector3D) Crossproduct_S(args ...interface{}) (ret interface{}) {
  // crossProduct(const class QVector3D &, const class QVector3D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[0][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector3D12crossProductERKS_S1_
    // invoke: QVector3D crossProduct(const class QVector3D &, const class QVector3D &)
    var arg0 = args[0].(*QVector3D).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVector3D).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN9QVector3D12crossProductERKS_S1_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVector3D{}) // "QVector3D"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector3D", "crossProduct", args)
  }

  return
}

// toVector2D()
func (this *QVector3D) Tovector2D(args ...interface{}) (ret interface{}) {
  // toVector2D()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D10toVector2DEv
    // invoke: QVector2D toVector2D()
    var ret0 = C.C_ZNK9QVector3D10toVector2DEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVector2D{}) // "QVector2D"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector3D", "toVector2D", args)
  }

  return
}

// distanceToPlane(const class QVector3D &, const class QVector3D &)
func (this *QVector3D) Distancetoplane(args ...interface{}) (ret interface{}) {
  // distanceToPlane(const class QVector3D &, const class QVector3D &)
  // distanceToPlane(const class QVector3D &, const class QVector3D &, const class QVector3D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[0][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[1][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[1][2] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D15distanceToPlaneERKS_S1_
    // invoke: float distanceToPlane(const class QVector3D &, const class QVector3D &)
    var arg0 = args[0].(*QVector3D).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVector3D).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK9QVector3D15distanceToPlaneERKS_S1_(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK9QVector3D15distanceToPlaneERKS_S1_S1_
    // invoke: float distanceToPlane(const class QVector3D &, const class QVector3D &, const class QVector3D &)
    var arg0 = args[0].(*QVector3D).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVector3D).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QVector3D).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK9QVector3D15distanceToPlaneERKS_S1_S1_(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector3D", "distanceToPlane", args)
  }

  return
}

// toVector4D()
func (this *QVector3D) Tovector4D(args ...interface{}) (ret interface{}) {
  // toVector4D()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D10toVector4DEv
    // invoke: QVector4D toVector4D()
    var ret0 = C.C_ZNK9QVector3D10toVector4DEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVector4D{}) // "QVector4D"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector3D", "toVector4D", args)
  }

  return
}

// dotProduct(const class QVector3D &, const class QVector3D &)
func (this *QVector3D) Dotproduct_S(args ...interface{}) (ret interface{}) {
  // dotProduct(const class QVector3D &, const class QVector3D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[0][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector3D10dotProductERKS_S1_
    // invoke: float dotProduct(const class QVector3D &, const class QVector3D &)
    var arg0 = args[0].(*QVector3D).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVector3D).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN9QVector3D10dotProductERKS_S1_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector3D", "dotProduct", args)
  }

  return
}

// normalize()
func (this *QVector3D) Normalize(args ...interface{}) () {
  // normalize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector3D9normalizeEv
    // invoke: void normalize()
    C.C_ZN9QVector3D9normalizeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QVector3D", "normalize", args)
  }

  return
}

// toPointF()
func (this *QVector3D) Topointf(args ...interface{}) (ret interface{}) {
  // toPointF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D8toPointFEv
    // invoke: QPointF toPointF()
    var ret0 = C.C_ZNK9QVector3D8toPointFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector3D", "toPointF", args)
  }

  return
}

// lengthSquared()
func (this *QVector3D) Lengthsquared(args ...interface{}) (ret interface{}) {
  // lengthSquared()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D13lengthSquaredEv
    // invoke: float lengthSquared()
    var ret0 = C.C_ZNK9QVector3D13lengthSquaredEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector3D", "lengthSquared", args)
  }

  return
}

// normal(const class QVector3D &, const class QVector3D &, const class QVector3D &)
func (this *QVector3D) Normal_S(args ...interface{}) (ret interface{}) {
  // normal(const class QVector3D &, const class QVector3D &, const class QVector3D &)
  // normal(const class QVector3D &, const class QVector3D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[0][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[0][2] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[1][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector3D6normalERKS_S1_S1_
    // invoke: QVector3D normal(const class QVector3D &, const class QVector3D &, const class QVector3D &)
    var arg0 = args[0].(*QVector3D).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVector3D).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QVector3D).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN9QVector3D6normalERKS_S1_S1_(arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVector3D{}) // "QVector3D"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN9QVector3D6normalERKS_S1_
    // invoke: QVector3D normal(const class QVector3D &, const class QVector3D &)
    var arg0 = args[0].(*QVector3D).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVector3D).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN9QVector3D6normalERKS_S1_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVector3D{}) // "QVector3D"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector3D", "normal", args)
  }

  return
}

// QVector3D(const class QPointF &)
func NewQVector3D(args ...interface{}) *QVector3D {
  // QVector3D(const class QPointF &)
  // QVector3D()
  // QVector3D(const class QVector2D &)
  // QVector3D(const class QPoint &)
  // QVector3D(float, float, float)
  // QVector3D(const class QVector4D &)
  // QVector3D(const class QVector2D &, float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.FloatTy(false) // "float"
  vtys[4][1] = qtrt.FloatTy(false) // "float"
  vtys[4][2] = qtrt.FloatTy(false) // "float"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QVector4D{}) // "const QVector4D &"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"
  vtys[6][1] = qtrt.FloatTy(false) // "float"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector3DC1ERK7QPointF
    // invoke: void QVector3D(const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QVector3DC2ERK7QPointF(arg0)
    return &QVector3D{Qclsinst:qthis}
  case 1:
    // invoke: _ZN9QVector3DC1Ev
    // invoke: void QVector3D()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QVector3DC2Ev()
    return &QVector3D{Qclsinst:qthis}
  case 2:
    // invoke: _ZN9QVector3DC1ERK9QVector2D
    // invoke: void QVector3D(const class QVector2D &)
    var arg0 = args[0].(*QVector2D).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QVector3DC2ERK9QVector2D(arg0)
    return &QVector3D{Qclsinst:qthis}
  case 3:
    // invoke: _ZN9QVector3DC1ERK6QPoint
    // invoke: void QVector3D(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QVector3DC2ERK6QPoint(arg0)
    return &QVector3D{Qclsinst:qthis}
  case 4:
    // invoke: _ZN9QVector3DC1Efff
    // invoke: void QVector3D(float, float, float)
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(qtrt.PrimConv(args[2], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QVector3DC2Efff(arg0, arg1, arg2)
    return &QVector3D{Qclsinst:qthis}
  case 5:
    // invoke: _ZN9QVector3DC1ERK9QVector4D
    // invoke: void QVector3D(const class QVector4D &)
    var arg0 = args[0].(*QVector4D).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QVector3DC2ERK9QVector4D(arg0)
    return &QVector3D{Qclsinst:qthis}
  case 6:
    // invoke: _ZN9QVector3DC1ERK9QVector2Df
    // invoke: void QVector3D(const class QVector2D &, float)
    var arg0 = args[0].(*QVector2D).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QVector3DC2ERK9QVector2Df(arg0, arg1)
    return &QVector3D{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QVector3D", "QVector3D", args)
  }

  return nil // QVector3D{Qclsinst:qthis}
}

// normalized()
func (this *QVector3D) Normalized(args ...interface{}) (ret interface{}) {
  // normalized()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D10normalizedEv
    // invoke: QVector3D normalized()
    var ret0 = C.C_ZNK9QVector3D10normalizedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVector3D{}) // "QVector3D"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector3D", "normalized", args)
  }

  return
}

// x()
func (this *QVector3D) X(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D1xEv
    // invoke: float x()
    C.C_ZNK9QVector3D1xEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QVector3D", "x", args)
  }

  return
}

// setX(float)
func (this *QVector3D) Setx(args ...interface{}) () {
  // setX(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector3D4setXEf
    // invoke: void setX(float)
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QVector3D4setXEf(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVector3D", "setX", args)
  }

  return
}

// setY(float)
func (this *QVector3D) Sety(args ...interface{}) () {
  // setY(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector3D4setYEf
    // invoke: void setY(float)
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QVector3D4setYEf(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVector3D", "setY", args)
  }

  return
}

// setZ(float)
func (this *QVector3D) Setz(args ...interface{}) () {
  // setZ(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector3D4setZEf
    // invoke: void setZ(float)
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QVector3D4setZEf(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVector3D", "setZ", args)
  }

  return
}

// toPoint()
func (this *QVector3D) Topoint(args ...interface{}) (ret interface{}) {
  // toPoint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D7toPointEv
    // invoke: QPoint toPoint()
    var ret0 = C.C_ZNK9QVector3D7toPointEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector3D", "toPoint", args)
  }

  return
}

// unproject(const class QMatrix4x4 &, const class QMatrix4x4 &, const class QRect &)
func (this *QVector3D) Unproject(args ...interface{}) (ret interface{}) {
  // unproject(const class QMatrix4x4 &, const class QMatrix4x4 &, const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix4x4{}) // "const QMatrix4x4 &"
  vtys[0][1] = reflect.TypeOf(QMatrix4x4{}) // "const QMatrix4x4 &"
  vtys[0][2] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D9unprojectERK10QMatrix4x4S2_RK5QRect
    // invoke: QVector3D unproject(const class QMatrix4x4 &, const class QMatrix4x4 &, const class QRect &)
    var arg0 = args[0].(*QMatrix4x4).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QMatrix4x4).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK9QVector3D9unprojectERK10QMatrix4x4S2_RK5QRect(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVector3D{}) // "QVector3D"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector3D", "unproject", args)
  }

  return
}

// distanceToLine(const class QVector3D &, const class QVector3D &)
func (this *QVector3D) Distancetoline(args ...interface{}) (ret interface{}) {
  // distanceToLine(const class QVector3D &, const class QVector3D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[0][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D14distanceToLineERKS_S1_
    // invoke: float distanceToLine(const class QVector3D &, const class QVector3D &)
    var arg0 = args[0].(*QVector3D).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVector3D).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK9QVector3D14distanceToLineERKS_S1_(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector3D", "distanceToLine", args)
  }

  return
}

// project(const class QMatrix4x4 &, const class QMatrix4x4 &, const class QRect &)
func (this *QVector3D) Project(args ...interface{}) (ret interface{}) {
  // project(const class QMatrix4x4 &, const class QMatrix4x4 &, const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix4x4{}) // "const QMatrix4x4 &"
  vtys[0][1] = reflect.TypeOf(QMatrix4x4{}) // "const QMatrix4x4 &"
  vtys[0][2] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D7projectERK10QMatrix4x4S2_RK5QRect
    // invoke: QVector3D project(const class QMatrix4x4 &, const class QMatrix4x4 &, const class QRect &)
    var arg0 = args[0].(*QMatrix4x4).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QMatrix4x4).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK9QVector3D7projectERK10QMatrix4x4S2_RK5QRect(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVector3D{}) // "QVector3D"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector3D", "project", args)
  }

  return
}

// isNull()
func (this *QVector3D) Isnull(args ...interface{}) (ret interface{}) {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D6isNullEv
    // invoke: bool isNull()
    var ret0 = C.C_ZNK9QVector3D6isNullEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector3D", "isNull", args)
  }

  return
}

// length()
func (this *QVector3D) Length(args ...interface{}) (ret interface{}) {
  // length()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D6lengthEv
    // invoke: float length()
    var ret0 = C.C_ZNK9QVector3D6lengthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector3D", "length", args)
  }

  return
}

// distanceToPoint(const class QVector3D &)
func (this *QVector3D) Distancetopoint(args ...interface{}) (ret interface{}) {
  // distanceToPoint(const class QVector3D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D15distanceToPointERKS_
    // invoke: float distanceToPoint(const class QVector3D &)
    var arg0 = args[0].(*QVector3D).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QVector3D15distanceToPointERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector3D", "distanceToPoint", args)
  }

  return
}

// y()
func (this *QVector3D) Y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D1yEv
    // invoke: float y()
    C.C_ZNK9QVector3D1yEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QVector3D", "y", args)
  }

  return
}

// z()
func (this *QVector3D) Z(args ...interface{}) (ret interface{}) {
  // z()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector3D1zEv
    // invoke: float z()
    var ret0 = C.C_ZNK9QVector3D1zEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVector3D", "z", args)
  }

  return
}

// <= body block end

