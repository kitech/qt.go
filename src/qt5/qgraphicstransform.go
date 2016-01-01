package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtWidgets/qgraphicstransform.h
// dst-file: /src/widgets/qgraphicstransform.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
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
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "origin", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "setAngle", args)
 }

}


func NewQGraphicsRotation(args ...interface{})() {
}


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
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "metaObject", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "setOrigin", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "axis", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "applyTo", args)
 }

}


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
    // invoke: _ZN17QGraphicsRotation7setAxisEN2Qt4AxisE
  case 1:
    // invoke: _ZN17QGraphicsRotation7setAxisERK9QVector3D
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "setAxis", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "angle", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsScale", "applyTo", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsScale", "zScale", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsScale", "xScale", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsScale", "yScale", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsScale", "setOrigin", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsScale", "setYScale", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsScale", "origin", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsScale", "setZScale", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsScale", "setXScale", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsScale", "metaObject", args)
 }

}


func NewQGraphicsScale(args ...interface{})() {
}


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
  default:
    qtrt.ErrorResolve("QGraphicsTransform", "applyTo", args)
 }

}


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


func NewQGraphicsTransform(args ...interface{})() {
}


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
  default:
    qtrt.ErrorResolve("QGraphicsTransform", "metaObject", args)
 }

}

// <= body block end

