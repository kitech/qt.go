package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
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
  // proto:  QVector3D QGraphicsRotation::origin();
extern void _ZNK17QGraphicsRotation6originEv(void* qthis);
  // proto:  void QGraphicsRotation::setAngle(qreal );
extern void _ZN17QGraphicsRotation8setAngleEd(void* qthis, double arg0);
  // proto:  void QGraphicsRotation::QGraphicsRotation(QObject * parent);
extern void* dector_ZN17QGraphicsRotationC1EP7QObject(void* arg0);
extern void _ZN17QGraphicsRotationC1EP7QObject(void* qthis, void* arg0);
  // proto:  const QMetaObject * QGraphicsRotation::metaObject();
extern void _ZNK17QGraphicsRotation10metaObjectEv(void* qthis);
  // proto:  void QGraphicsRotation::~QGraphicsRotation();
extern void _ZN17QGraphicsRotationD0Ev(void* qthis);
  // proto:  void QGraphicsRotation::setOrigin(const QVector3D & point);
extern void _ZN17QGraphicsRotation9setOriginERK9QVector3D(void* qthis, void* arg0);
  // proto:  QVector3D QGraphicsRotation::axis();
extern void _ZNK17QGraphicsRotation4axisEv(void* qthis);
  // proto:  void QGraphicsRotation::applyTo(QMatrix4x4 * matrix);
extern void _ZNK17QGraphicsRotation7applyToEP10QMatrix4x4(void* qthis, void* arg0);
  // proto:  void QGraphicsRotation::setAxis(const QVector3D & axis);
extern void _ZN17QGraphicsRotation7setAxisERK9QVector3D(void* qthis, void* arg0);
  // proto:  qreal QGraphicsRotation::angle();
extern void _ZNK17QGraphicsRotation5angleEv(void* qthis);
  // proto:  void QGraphicsScale::applyTo(QMatrix4x4 * matrix);
extern void _ZNK14QGraphicsScale7applyToEP10QMatrix4x4(void* qthis, void* arg0);
  // proto:  qreal QGraphicsScale::zScale();
extern void _ZNK14QGraphicsScale6zScaleEv(void* qthis);
  // proto:  qreal QGraphicsScale::xScale();
extern void _ZNK14QGraphicsScale6xScaleEv(void* qthis);
  // proto:  qreal QGraphicsScale::yScale();
extern void _ZNK14QGraphicsScale6yScaleEv(void* qthis);
  // proto:  void QGraphicsScale::setOrigin(const QVector3D & point);
extern void _ZN14QGraphicsScale9setOriginERK9QVector3D(void* qthis, void* arg0);
  // proto:  void QGraphicsScale::setYScale(qreal );
extern void _ZN14QGraphicsScale9setYScaleEd(void* qthis, double arg0);
  // proto:  QVector3D QGraphicsScale::origin();
extern void _ZNK14QGraphicsScale6originEv(void* qthis);
  // proto:  void QGraphicsScale::setZScale(qreal );
extern void _ZN14QGraphicsScale9setZScaleEd(void* qthis, double arg0);
  // proto:  void QGraphicsScale::setXScale(qreal );
extern void _ZN14QGraphicsScale9setXScaleEd(void* qthis, double arg0);
  // proto:  const QMetaObject * QGraphicsScale::metaObject();
extern void _ZNK14QGraphicsScale10metaObjectEv(void* qthis);
  // proto:  void QGraphicsScale::QGraphicsScale(QObject * parent);
extern void* dector_ZN14QGraphicsScaleC1EP7QObject(void* arg0);
extern void _ZN14QGraphicsScaleC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QGraphicsScale::~QGraphicsScale();
extern void _ZN14QGraphicsScaleD0Ev(void* qthis);
  // proto:  void QGraphicsTransform::applyTo(QMatrix4x4 * matrix);
extern void _ZNK18QGraphicsTransform7applyToEP10QMatrix4x4(void* qthis, void* arg0);
  // proto:  void QGraphicsTransform::~QGraphicsTransform();
extern void _ZN18QGraphicsTransformD0Ev(void* qthis);
  // proto:  void QGraphicsTransform::QGraphicsTransform(QObject * parent);
extern void* dector_ZN18QGraphicsTransformC1EP7QObject(void* arg0);
extern void _ZN18QGraphicsTransformC1EP7QObject(void* qthis, void* arg0);
  // proto:  const QMetaObject * QGraphicsTransform::metaObject();
extern void _ZNK18QGraphicsTransform10metaObjectEv(void* qthis);
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

// class sizeof(QGraphicsRotation)=1
type QGraphicsRotation struct {
  /*qbase*/ QGraphicsTransform;
  qclsinst uint64 /* *mut c_void*/;
//  _originChanged QGraphicsRotation_originChanged_signal;
//  _axisChanged QGraphicsRotation_axisChanged_signal;
//  _angleChanged QGraphicsRotation_angleChanged_signal;
}

// class sizeof(QGraphicsScale)=1
type QGraphicsScale struct {
  /*qbase*/ QGraphicsTransform;
  qclsinst uint64 /* *mut c_void*/;
//  _yScaleChanged QGraphicsScale_yScaleChanged_signal;
//  _xScaleChanged QGraphicsScale_xScaleChanged_signal;
//  _zScaleChanged QGraphicsScale_zScaleChanged_signal;
//  _scaleChanged QGraphicsScale_scaleChanged_signal;
//  _originChanged QGraphicsScale_originChanged_signal;
}

// class sizeof(QGraphicsTransform)=1
type QGraphicsTransform struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QVector3D QGraphicsRotation::origin();
func (this *QGraphicsRotation) origin(args ...interface{}) () {
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
    C._ZNK17QGraphicsRotation6originEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "origin", args)
  }

}

  // proto:  void QGraphicsRotation::setAngle(qreal );
func (this *QGraphicsRotation) setAngle(args ...interface{}) () {
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN17QGraphicsRotation8setAngleEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "setAngle", args)
  }

}

  // proto:  void QGraphicsRotation::QGraphicsRotation(QObject * parent);
func NewQGraphicsRotation(args ...interface{}) QGraphicsRotation {
  return QGraphicsRotation{}
}

  // proto:  const QMetaObject * QGraphicsRotation::metaObject();
func (this *QGraphicsRotation) metaObject(args ...interface{}) () {
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
    C._ZNK17QGraphicsRotation10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "metaObject", args)
  }

}

  // proto:  void QGraphicsRotation::~QGraphicsRotation();
func (this *QGraphicsRotation) FreeQGraphicsRotation(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "~QGraphicsRotation", args)
  }

}

  // proto:  void QGraphicsRotation::setOrigin(const QVector3D & point);
func (this *QGraphicsRotation) setOrigin(args ...interface{}) () {
  // setOrigin(const class QVector3D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsRotation9setOriginERK9QVector3D
    // invoke: void setOrigin(const class QVector3D &)
    var arg0 = args[0].(QVector3D).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN17QGraphicsRotation9setOriginERK9QVector3D(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "setOrigin", args)
  }

}

  // proto:  QVector3D QGraphicsRotation::axis();
func (this *QGraphicsRotation) axis(args ...interface{}) () {
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
    C._ZNK17QGraphicsRotation4axisEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "axis", args)
  }

}

  // proto:  void QGraphicsRotation::applyTo(QMatrix4x4 * matrix);
func (this *QGraphicsRotation) applyTo(args ...interface{}) () {
  // applyTo(class QMatrix4x4 *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix4x4{}) // "QMatrix4x4 *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRotation7applyToEP10QMatrix4x4
    // invoke: void applyTo(class QMatrix4x4 *)
    var arg0 = args[0].(QMatrix4x4).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK17QGraphicsRotation7applyToEP10QMatrix4x4(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "applyTo", args)
  }

}

  // proto:  void QGraphicsRotation::setAxis(const QVector3D & axis);
func (this *QGraphicsRotation) setAxis(args ...interface{}) () {
  // setAxis(Qt::Axis)
  // setAxis(const class QVector3D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "Qt::Axis"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsRotation7setAxisERK9QVector3D
    // invoke: void setAxis(const class QVector3D &)
    var arg0 = args[0].(QVector3D).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN17QGraphicsRotation7setAxisERK9QVector3D(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "setAxis", args)
  }

}

  // proto:  qreal QGraphicsRotation::angle();
func (this *QGraphicsRotation) angle(args ...interface{}) () {
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
    C._ZNK17QGraphicsRotation5angleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "angle", args)
  }

}

  // proto:  void QGraphicsScale::applyTo(QMatrix4x4 * matrix);
func (this *QGraphicsScale) applyTo(args ...interface{}) () {
  // applyTo(class QMatrix4x4 *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix4x4{}) // "QMatrix4x4 *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScale7applyToEP10QMatrix4x4
    // invoke: void applyTo(class QMatrix4x4 *)
    var arg0 = args[0].(QMatrix4x4).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK14QGraphicsScale7applyToEP10QMatrix4x4(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "applyTo", args)
  }

}

  // proto:  qreal QGraphicsScale::zScale();
func (this *QGraphicsScale) zScale(args ...interface{}) () {
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
    C._ZNK14QGraphicsScale6zScaleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "zScale", args)
  }

}

  // proto:  qreal QGraphicsScale::xScale();
func (this *QGraphicsScale) xScale(args ...interface{}) () {
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
    C._ZNK14QGraphicsScale6xScaleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "xScale", args)
  }

}

  // proto:  qreal QGraphicsScale::yScale();
func (this *QGraphicsScale) yScale(args ...interface{}) () {
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
    C._ZNK14QGraphicsScale6yScaleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "yScale", args)
  }

}

  // proto:  void QGraphicsScale::setOrigin(const QVector3D & point);
func (this *QGraphicsScale) setOrigin(args ...interface{}) () {
  // setOrigin(const class QVector3D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScale9setOriginERK9QVector3D
    // invoke: void setOrigin(const class QVector3D &)
    var arg0 = args[0].(QVector3D).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QGraphicsScale9setOriginERK9QVector3D(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "setOrigin", args)
  }

}

  // proto:  void QGraphicsScale::setYScale(qreal );
func (this *QGraphicsScale) setYScale(args ...interface{}) () {
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN14QGraphicsScale9setYScaleEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "setYScale", args)
  }

}

  // proto:  QVector3D QGraphicsScale::origin();
func (this *QGraphicsScale) origin(args ...interface{}) () {
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
    C._ZNK14QGraphicsScale6originEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "origin", args)
  }

}

  // proto:  void QGraphicsScale::setZScale(qreal );
func (this *QGraphicsScale) setZScale(args ...interface{}) () {
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN14QGraphicsScale9setZScaleEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "setZScale", args)
  }

}

  // proto:  void QGraphicsScale::setXScale(qreal );
func (this *QGraphicsScale) setXScale(args ...interface{}) () {
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN14QGraphicsScale9setXScaleEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "setXScale", args)
  }

}

  // proto:  const QMetaObject * QGraphicsScale::metaObject();
func (this *QGraphicsScale) metaObject(args ...interface{}) () {
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
    C._ZNK14QGraphicsScale10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "metaObject", args)
  }

}

  // proto:  void QGraphicsScale::QGraphicsScale(QObject * parent);
func NewQGraphicsScale(args ...interface{}) QGraphicsScale {
  return QGraphicsScale{}
}

  // proto:  void QGraphicsScale::~QGraphicsScale();
func (this *QGraphicsScale) FreeQGraphicsScale(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsScale", "~QGraphicsScale", args)
  }

}

  // proto:  void QGraphicsTransform::applyTo(QMatrix4x4 * matrix);
func (this *QGraphicsTransform) applyTo(args ...interface{}) () {
  // applyTo(class QMatrix4x4 *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix4x4{}) // "QMatrix4x4 *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QGraphicsTransform7applyToEP10QMatrix4x4
    // invoke: void applyTo(class QMatrix4x4 *)
    var arg0 = args[0].(QMatrix4x4).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK18QGraphicsTransform7applyToEP10QMatrix4x4(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsTransform", "applyTo", args)
  }

}

  // proto:  void QGraphicsTransform::~QGraphicsTransform();
func (this *QGraphicsTransform) FreeQGraphicsTransform(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsTransform", "~QGraphicsTransform", args)
  }

}

  // proto:  void QGraphicsTransform::QGraphicsTransform(QObject * parent);
func NewQGraphicsTransform(args ...interface{}) QGraphicsTransform {
  return QGraphicsTransform{}
}

  // proto:  const QMetaObject * QGraphicsTransform::metaObject();
func (this *QGraphicsTransform) metaObject(args ...interface{}) () {
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
    C._ZNK18QGraphicsTransform10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsTransform", "metaObject", args)
  }

}

// <= body block end

