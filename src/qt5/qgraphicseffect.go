package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qgraphicseffect.h
// dst-file: /src/widgets/qgraphicseffect.go
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
// class sizeof(QGraphicsColorizeEffect)=1
type QGraphicsColorizeEffect struct {
  /*qbase*/ QGraphicsEffect;
  qclsinst uint64 /* *mut c_void*/;
//  _strengthChanged QGraphicsColorizeEffect_strengthChanged_signal;
//  _colorChanged QGraphicsColorizeEffect_colorChanged_signal;
}

// class sizeof(QGraphicsEffect)=1
type QGraphicsEffect struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _enabledChanged QGraphicsEffect_enabledChanged_signal;
}

// class sizeof(QGraphicsDropShadowEffect)=1
type QGraphicsDropShadowEffect struct {
  /*qbase*/ QGraphicsEffect;
  qclsinst uint64 /* *mut c_void*/;
//  _colorChanged QGraphicsDropShadowEffect_colorChanged_signal;
//  _offsetChanged QGraphicsDropShadowEffect_offsetChanged_signal;
//  _blurRadiusChanged QGraphicsDropShadowEffect_blurRadiusChanged_signal;
}

// class sizeof(QGraphicsOpacityEffect)=1
type QGraphicsOpacityEffect struct {
  /*qbase*/ QGraphicsEffect;
  qclsinst uint64 /* *mut c_void*/;
//  _opacityMaskChanged QGraphicsOpacityEffect_opacityMaskChanged_signal;
//  _opacityChanged QGraphicsOpacityEffect_opacityChanged_signal;
}

// class sizeof(QGraphicsBlurEffect)=1
type QGraphicsBlurEffect struct {
  /*qbase*/ QGraphicsEffect;
  qclsinst uint64 /* *mut c_void*/;
//  _blurHintsChanged QGraphicsBlurEffect_blurHintsChanged_signal;
//  _blurRadiusChanged QGraphicsBlurEffect_blurRadiusChanged_signal;
}


func (this *QGraphicsColorizeEffect) setColor(args ...interface{}) () {
  // setColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QGraphicsColorizeEffect8setColorERK6QColor
  default:
    qtrt.ErrorResolve("QGraphicsColorizeEffect", "setColor", args)
  }

}


func NewQGraphicsColorizeEffect(args ...interface{}) QGraphicsColorizeEffect {
  return QGraphicsColorizeEffect{}
}


func (this *QGraphicsColorizeEffect) setStrength(args ...interface{}) () {
  // setStrength(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QGraphicsColorizeEffect11setStrengthEd
  default:
    qtrt.ErrorResolve("QGraphicsColorizeEffect", "setStrength", args)
  }

}


func (this *QGraphicsColorizeEffect) strength(args ...interface{}) () {
  // strength()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsColorizeEffect8strengthEv
  default:
    qtrt.ErrorResolve("QGraphicsColorizeEffect", "strength", args)
  }

}


func (this *QGraphicsColorizeEffect) FreeQGraphicsColorizeEffect(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsColorizeEffect", "~QGraphicsColorizeEffect", args)
  }

}


func (this *QGraphicsColorizeEffect) color(args ...interface{}) () {
  // color()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsColorizeEffect5colorEv
  default:
    qtrt.ErrorResolve("QGraphicsColorizeEffect", "color", args)
  }

}


func (this *QGraphicsColorizeEffect) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsColorizeEffect10metaObjectEv
  default:
    qtrt.ErrorResolve("QGraphicsColorizeEffect", "metaObject", args)
  }

}


func (this *QGraphicsEffect) boundingRectFor(args ...interface{}) () {
  // boundingRectFor(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsEffect15boundingRectForERK6QRectF
  default:
    qtrt.ErrorResolve("QGraphicsEffect", "boundingRectFor", args)
  }

}


func (this *QGraphicsEffect) source(args ...interface{}) () {
  // source()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsEffect6sourceEv
  default:
    qtrt.ErrorResolve("QGraphicsEffect", "source", args)
  }

}


func (this *QGraphicsEffect) update(args ...interface{}) () {
  // update()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsEffect6updateEv
  default:
    qtrt.ErrorResolve("QGraphicsEffect", "update", args)
  }

}


func (this *QGraphicsEffect) setEnabled(args ...interface{}) () {
  // setEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsEffect10setEnabledEb
  default:
    qtrt.ErrorResolve("QGraphicsEffect", "setEnabled", args)
  }

}


func (this *QGraphicsEffect) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsEffect10metaObjectEv
  default:
    qtrt.ErrorResolve("QGraphicsEffect", "metaObject", args)
  }

}


func (this *QGraphicsEffect) isEnabled(args ...interface{}) () {
  // isEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsEffect9isEnabledEv
  default:
    qtrt.ErrorResolve("QGraphicsEffect", "isEnabled", args)
  }

}


func (this *QGraphicsEffect) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsEffect12boundingRectEv
  default:
    qtrt.ErrorResolve("QGraphicsEffect", "boundingRect", args)
  }

}


func (this *QGraphicsEffect) FreeQGraphicsEffect(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsEffect", "~QGraphicsEffect", args)
  }

}


func NewQGraphicsEffect(args ...interface{}) QGraphicsEffect {
  return QGraphicsEffect{}
}


func (this *QGraphicsDropShadowEffect) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QGraphicsDropShadowEffect10metaObjectEv
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "metaObject", args)
  }

}


func (this *QGraphicsDropShadowEffect) boundingRectFor(args ...interface{}) () {
  // boundingRectFor(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QGraphicsDropShadowEffect15boundingRectForERK6QRectF
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "boundingRectFor", args)
  }

}


func NewQGraphicsDropShadowEffect(args ...interface{}) QGraphicsDropShadowEffect {
  return QGraphicsDropShadowEffect{}
}


func (this *QGraphicsDropShadowEffect) offset(args ...interface{}) () {
  // offset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QGraphicsDropShadowEffect6offsetEv
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "offset", args)
  }

}


func (this *QGraphicsDropShadowEffect) setYOffset(args ...interface{}) () {
  // setYOffset(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QGraphicsDropShadowEffect10setYOffsetEd
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "setYOffset", args)
  }

}


func (this *QGraphicsDropShadowEffect) xOffset(args ...interface{}) () {
  // xOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QGraphicsDropShadowEffect7xOffsetEv
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "xOffset", args)
  }

}


func (this *QGraphicsDropShadowEffect) blurRadius(args ...interface{}) () {
  // blurRadius()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QGraphicsDropShadowEffect10blurRadiusEv
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "blurRadius", args)
  }

}


func (this *QGraphicsDropShadowEffect) color(args ...interface{}) () {
  // color()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QGraphicsDropShadowEffect5colorEv
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "color", args)
  }

}


func (this *QGraphicsDropShadowEffect) setColor(args ...interface{}) () {
  // setColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QGraphicsDropShadowEffect8setColorERK6QColor
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "setColor", args)
  }

}


func (this *QGraphicsDropShadowEffect) setOffset(args ...interface{}) () {
  // setOffset(const class QPointF &)
  // setOffset(qreal, qreal)
  // setOffset(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QGraphicsDropShadowEffect9setOffsetERK7QPointF
  case 1:
    // invoke: _ZN25QGraphicsDropShadowEffect9setOffsetEdd
  case 2:
    // invoke: _ZN25QGraphicsDropShadowEffect9setOffsetEd
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "setOffset", args)
  }

}


func (this *QGraphicsDropShadowEffect) yOffset(args ...interface{}) () {
  // yOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QGraphicsDropShadowEffect7yOffsetEv
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "yOffset", args)
  }

}


func (this *QGraphicsDropShadowEffect) setXOffset(args ...interface{}) () {
  // setXOffset(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QGraphicsDropShadowEffect10setXOffsetEd
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "setXOffset", args)
  }

}


func (this *QGraphicsDropShadowEffect) setBlurRadius(args ...interface{}) () {
  // setBlurRadius(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QGraphicsDropShadowEffect13setBlurRadiusEd
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "setBlurRadius", args)
  }

}


func (this *QGraphicsDropShadowEffect) FreeQGraphicsDropShadowEffect(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "~QGraphicsDropShadowEffect", args)
  }

}


func NewQGraphicsOpacityEffect(args ...interface{}) QGraphicsOpacityEffect {
  return QGraphicsOpacityEffect{}
}


func (this *QGraphicsOpacityEffect) FreeQGraphicsOpacityEffect(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsOpacityEffect", "~QGraphicsOpacityEffect", args)
  }

}


func (this *QGraphicsOpacityEffect) setOpacityMask(args ...interface{}) () {
  // setOpacityMask(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsOpacityEffect14setOpacityMaskERK6QBrush
  default:
    qtrt.ErrorResolve("QGraphicsOpacityEffect", "setOpacityMask", args)
  }

}


func (this *QGraphicsOpacityEffect) opacityMask(args ...interface{}) () {
  // opacityMask()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsOpacityEffect11opacityMaskEv
  default:
    qtrt.ErrorResolve("QGraphicsOpacityEffect", "opacityMask", args)
  }

}


func (this *QGraphicsOpacityEffect) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsOpacityEffect10metaObjectEv
  default:
    qtrt.ErrorResolve("QGraphicsOpacityEffect", "metaObject", args)
  }

}


func (this *QGraphicsOpacityEffect) opacity(args ...interface{}) () {
  // opacity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsOpacityEffect7opacityEv
  default:
    qtrt.ErrorResolve("QGraphicsOpacityEffect", "opacity", args)
  }

}


func (this *QGraphicsOpacityEffect) setOpacity(args ...interface{}) () {
  // setOpacity(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsOpacityEffect10setOpacityEd
  default:
    qtrt.ErrorResolve("QGraphicsOpacityEffect", "setOpacity", args)
  }

}


func (this *QGraphicsBlurEffect) blurRadius(args ...interface{}) () {
  // blurRadius()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsBlurEffect10blurRadiusEv
  default:
    qtrt.ErrorResolve("QGraphicsBlurEffect", "blurRadius", args)
  }

}


func (this *QGraphicsBlurEffect) setBlurRadius(args ...interface{}) () {
  // setBlurRadius(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsBlurEffect13setBlurRadiusEd
  default:
    qtrt.ErrorResolve("QGraphicsBlurEffect", "setBlurRadius", args)
  }

}


func (this *QGraphicsBlurEffect) FreeQGraphicsBlurEffect(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsBlurEffect", "~QGraphicsBlurEffect", args)
  }

}


func NewQGraphicsBlurEffect(args ...interface{}) QGraphicsBlurEffect {
  return QGraphicsBlurEffect{}
}


func (this *QGraphicsBlurEffect) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsBlurEffect10metaObjectEv
  default:
    qtrt.ErrorResolve("QGraphicsBlurEffect", "metaObject", args)
  }

}


func (this *QGraphicsBlurEffect) boundingRectFor(args ...interface{}) () {
  // boundingRectFor(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsBlurEffect15boundingRectForERK6QRectF
  default:
    qtrt.ErrorResolve("QGraphicsBlurEffect", "boundingRectFor", args)
  }

}

// <= body block end

