package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtWidgets/qgraphicstransform.h
// dst-file: /src/widgets/qgraphicstransform.go
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
import "qtgui"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QVector3D QGraphicsRotation::origin();
extern void* C_ZNK17QGraphicsRotation6originEv(void* qthis); // 4
  // proto:  const QMetaObject * QGraphicsRotation::metaObject();
extern void C_ZNK17QGraphicsRotation10metaObjectEv(void* qthis); // 4
  // proto:  void QGraphicsRotation::applyTo(QMatrix4x4 * matrix);
extern void C_ZNK17QGraphicsRotation7applyToEP10QMatrix4x4(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsRotation::setOrigin(const QVector3D & point);
extern void C_ZN17QGraphicsRotation9setOriginERK9QVector3D(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsRotation::QGraphicsRotation(QObject * parent);
extern void* C_ZN17QGraphicsRotationC2EP7QObject(void* arg0); // 3
  // proto:  void QGraphicsRotation::setAngle(qreal );
extern void C_ZN17QGraphicsRotation8setAngleEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsRotation::~QGraphicsRotation();
extern void C_ZN17QGraphicsRotationD2Ev(void* qthis); // 4
  // proto:  qreal QGraphicsRotation::angle();
extern double C_ZNK17QGraphicsRotation5angleEv(void* qthis); // 4
  // proto:  QVector3D QGraphicsRotation::axis();
extern void* C_ZNK17QGraphicsRotation4axisEv(void* qthis); // 4
  // proto:  void QGraphicsRotation::setAxis(const QVector3D & axis);
extern void C_ZN17QGraphicsRotation7setAxisERK9QVector3D(void* qthis, void* arg0); // 4
  // proto:  QVector3D QGraphicsScale::origin();
extern void* C_ZNK14QGraphicsScale6originEv(void* qthis); // 4
  // proto:  qreal QGraphicsScale::zScale();
extern double C_ZNK14QGraphicsScale6zScaleEv(void* qthis); // 4
  // proto:  void QGraphicsScale::QGraphicsScale(QObject * parent);
extern void* C_ZN14QGraphicsScaleC2EP7QObject(void* arg0); // 3
  // proto:  qreal QGraphicsScale::yScale();
extern double C_ZNK14QGraphicsScale6yScaleEv(void* qthis); // 4
  // proto:  void QGraphicsScale::~QGraphicsScale();
extern void C_ZN14QGraphicsScaleD2Ev(void* qthis); // 4
  // proto:  void QGraphicsScale::setZScale(qreal );
extern void C_ZN14QGraphicsScale9setZScaleEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsScale::setYScale(qreal );
extern void C_ZN14QGraphicsScale9setYScaleEd(void* qthis, double arg0); // 4
  // proto:  qreal QGraphicsScale::xScale();
extern double C_ZNK14QGraphicsScale6xScaleEv(void* qthis); // 4
  // proto:  const QMetaObject * QGraphicsScale::metaObject();
extern void C_ZNK14QGraphicsScale10metaObjectEv(void* qthis); // 4
  // proto:  void QGraphicsScale::applyTo(QMatrix4x4 * matrix);
extern void C_ZNK14QGraphicsScale7applyToEP10QMatrix4x4(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsScale::setOrigin(const QVector3D & point);
extern void C_ZN14QGraphicsScale9setOriginERK9QVector3D(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsScale::setXScale(qreal );
extern void C_ZN14QGraphicsScale9setXScaleEd(void* qthis, double arg0); // 4
  // proto:  const QMetaObject * QGraphicsTransform::metaObject();
extern void C_ZNK18QGraphicsTransform10metaObjectEv(void* qthis); // 4
  // proto:  void QGraphicsTransform::QGraphicsTransform(QObject * parent);
extern void* C_ZN18QGraphicsTransformC2EP7QObject(void* arg0); // 3
  // proto:  void QGraphicsTransform::~QGraphicsTransform();
extern void C_ZN18QGraphicsTransformD2Ev(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QGraphicsRotation)=1
type QGraphicsRotation struct {
  /*qbase*/ QGraphicsTransform;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _originChanged QGraphicsRotation_originChanged_signal;
//  _axisChanged QGraphicsRotation_axisChanged_signal;
//  _angleChanged QGraphicsRotation_angleChanged_signal;
}

// class sizeof(QGraphicsScale)=1
type QGraphicsScale struct {
  /*qbase*/ QGraphicsTransform;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _yScaleChanged QGraphicsScale_yScaleChanged_signal;
//  _xScaleChanged QGraphicsScale_xScaleChanged_signal;
//  _zScaleChanged QGraphicsScale_zScaleChanged_signal;
//  _scaleChanged QGraphicsScale_scaleChanged_signal;
//  _originChanged QGraphicsScale_originChanged_signal;
}

// class sizeof(QGraphicsTransform)=1
type QGraphicsTransform struct {
  /*qbase*/ qtcore.QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// origin()
func (this *QGraphicsRotation) Origin(args ...interface{}) (ret interface{}) {
  // origin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRotation6originEv
    // invoke: QVector3D origin()
    var ret0 = C.C_ZNK17QGraphicsRotation6originEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QVector3D{}) // "QVector3D"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "origin", args)
  }

  return
}

// metaObject()
func (this *QGraphicsRotation) MetaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRotation10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK17QGraphicsRotation10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "metaObject", args)
  }

  return
}

// applyTo(class QMatrix4x4 *)
func (this *QGraphicsRotation) ApplyTo(args ...interface{}) () {
  // applyTo(class QMatrix4x4 *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QMatrix4x4{}) // "QMatrix4x4 *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRotation7applyToEP10QMatrix4x4
    // invoke: void applyTo(class QMatrix4x4 *)
    var arg0 = args[0].(*qtgui.QMatrix4x4).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK17QGraphicsRotation7applyToEP10QMatrix4x4(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "applyTo", args)
  }

  return
}

// setOrigin(const class QVector3D &)
func (this *QGraphicsRotation) SetOrigin(args ...interface{}) () {
  // setOrigin(const class QVector3D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QVector3D{}) // "const QVector3D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsRotation9setOriginERK9QVector3D
    // invoke: void setOrigin(const class QVector3D &)
    var arg0 = args[0].(*qtgui.QVector3D).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsRotation9setOriginERK9QVector3D(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "setOrigin", args)
  }

  return
}

// QGraphicsRotation(class QObject *)
func GcfreeQGraphicsRotation(this *QGraphicsRotation) {
  qtrt.UniverseFree(this)
}
func NewQGraphicsRotation(args ...interface{}) *QGraphicsRotation {
  // QGraphicsRotation(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsRotationC1EP7QObject
    // invoke: void QGraphicsRotation(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QGraphicsRotationC2EP7QObject(arg0)
    this := &QGraphicsRotation{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQGraphicsRotation)
    return this
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "QGraphicsRotation", args)
  }

  return nil // QGraphicsRotation{Qclsinst:qthis}
}

// setAngle(qreal)
func (this *QGraphicsRotation) SetAngle(args ...interface{}) () {
  // setAngle(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsRotation8setAngleEd
    // invoke: void setAngle(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsRotation8setAngleEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "setAngle", args)
  }

  return
}

// ~QGraphicsRotation()
func (this *QGraphicsRotation) Free(args ...interface{}) () {
  // ~QGraphicsRotation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsRotationD0Ev
    // invoke: void ~QGraphicsRotation()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN17QGraphicsRotationD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "~QGraphicsRotation", args)
  }

  return
}

// angle()
func (this *QGraphicsRotation) Angle(args ...interface{}) (ret interface{}) {
  // angle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRotation5angleEv
    // invoke: qreal angle()
    var ret0 = C.C_ZNK17QGraphicsRotation5angleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "angle", args)
  }

  return
}

// axis()
func (this *QGraphicsRotation) Axis(args ...interface{}) (ret interface{}) {
  // axis()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRotation4axisEv
    // invoke: QVector3D axis()
    var ret0 = C.C_ZNK17QGraphicsRotation4axisEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QVector3D{}) // "QVector3D"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "axis", args)
  }

  return
}

// setAxis(const class QVector3D &)
func (this *QGraphicsRotation) SetAxis(args ...interface{}) () {
  // setAxis(const class QVector3D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QVector3D{}) // "const QVector3D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsRotation7setAxisERK9QVector3D
    // invoke: void setAxis(const class QVector3D &)
    var arg0 = args[0].(*qtgui.QVector3D).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsRotation7setAxisERK9QVector3D(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "setAxis", args)
  }

  return
}

// origin()
func (this *QGraphicsScale) Origin(args ...interface{}) (ret interface{}) {
  // origin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScale6originEv
    // invoke: QVector3D origin()
    var ret0 = C.C_ZNK14QGraphicsScale6originEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QVector3D{}) // "QVector3D"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsScale", "origin", args)
  }

  return
}

// zScale()
func (this *QGraphicsScale) ZScale(args ...interface{}) (ret interface{}) {
  // zScale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScale6zScaleEv
    // invoke: qreal zScale()
    var ret0 = C.C_ZNK14QGraphicsScale6zScaleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsScale", "zScale", args)
  }

  return
}

// QGraphicsScale(class QObject *)
func GcfreeQGraphicsScale(this *QGraphicsScale) {
  qtrt.UniverseFree(this)
}
func NewQGraphicsScale(args ...interface{}) *QGraphicsScale {
  // QGraphicsScale(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScaleC1EP7QObject
    // invoke: void QGraphicsScale(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QGraphicsScaleC2EP7QObject(arg0)
    this := &QGraphicsScale{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQGraphicsScale)
    return this
  default:
    qtrt.ErrorResolve("QGraphicsScale", "QGraphicsScale", args)
  }

  return nil // QGraphicsScale{Qclsinst:qthis}
}

// yScale()
func (this *QGraphicsScale) YScale(args ...interface{}) (ret interface{}) {
  // yScale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScale6yScaleEv
    // invoke: qreal yScale()
    var ret0 = C.C_ZNK14QGraphicsScale6yScaleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsScale", "yScale", args)
  }

  return
}

// ~QGraphicsScale()
func (this *QGraphicsScale) Free(args ...interface{}) () {
  // ~QGraphicsScale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScaleD0Ev
    // invoke: void ~QGraphicsScale()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN14QGraphicsScaleD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QGraphicsScale", "~QGraphicsScale", args)
  }

  return
}

// setZScale(qreal)
func (this *QGraphicsScale) SetZScale(args ...interface{}) () {
  // setZScale(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScale9setZScaleEd
    // invoke: void setZScale(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN14QGraphicsScale9setZScaleEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "setZScale", args)
  }

  return
}

// setYScale(qreal)
func (this *QGraphicsScale) SetYScale(args ...interface{}) () {
  // setYScale(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScale9setYScaleEd
    // invoke: void setYScale(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN14QGraphicsScale9setYScaleEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "setYScale", args)
  }

  return
}

// xScale()
func (this *QGraphicsScale) XScale(args ...interface{}) (ret interface{}) {
  // xScale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScale6xScaleEv
    // invoke: qreal xScale()
    var ret0 = C.C_ZNK14QGraphicsScale6xScaleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsScale", "xScale", args)
  }

  return
}

// metaObject()
func (this *QGraphicsScale) MetaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScale10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK14QGraphicsScale10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "metaObject", args)
  }

  return
}

// applyTo(class QMatrix4x4 *)
func (this *QGraphicsScale) ApplyTo(args ...interface{}) () {
  // applyTo(class QMatrix4x4 *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QMatrix4x4{}) // "QMatrix4x4 *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScale7applyToEP10QMatrix4x4
    // invoke: void applyTo(class QMatrix4x4 *)
    var arg0 = args[0].(*qtgui.QMatrix4x4).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK14QGraphicsScale7applyToEP10QMatrix4x4(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "applyTo", args)
  }

  return
}

// setOrigin(const class QVector3D &)
func (this *QGraphicsScale) SetOrigin(args ...interface{}) () {
  // setOrigin(const class QVector3D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QVector3D{}) // "const QVector3D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScale9setOriginERK9QVector3D
    // invoke: void setOrigin(const class QVector3D &)
    var arg0 = args[0].(*qtgui.QVector3D).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QGraphicsScale9setOriginERK9QVector3D(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "setOrigin", args)
  }

  return
}

// setXScale(qreal)
func (this *QGraphicsScale) SetXScale(args ...interface{}) () {
  // setXScale(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScale9setXScaleEd
    // invoke: void setXScale(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN14QGraphicsScale9setXScaleEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "setXScale", args)
  }

  return
}

// metaObject()
func (this *QGraphicsTransform) MetaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QGraphicsTransform10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK18QGraphicsTransform10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsTransform", "metaObject", args)
  }

  return
}

// QGraphicsTransform(class QObject *)
func GcfreeQGraphicsTransform(this *QGraphicsTransform) {
  qtrt.UniverseFree(this)
}
func NewQGraphicsTransform(args ...interface{}) *QGraphicsTransform {
  // QGraphicsTransform(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QGraphicsTransformC1EP7QObject
    // invoke: void QGraphicsTransform(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QGraphicsTransformC2EP7QObject(arg0)
    this := &QGraphicsTransform{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQGraphicsTransform)
    return this
  default:
    qtrt.ErrorResolve("QGraphicsTransform", "QGraphicsTransform", args)
  }

  return nil // QGraphicsTransform{Qclsinst:qthis}
}

// ~QGraphicsTransform()
func (this *QGraphicsTransform) Free(args ...interface{}) () {
  // ~QGraphicsTransform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QGraphicsTransformD0Ev
    // invoke: void ~QGraphicsTransform()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN18QGraphicsTransformD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QGraphicsTransform", "~QGraphicsTransform", args)
  }

  return
}

// <= body block end

