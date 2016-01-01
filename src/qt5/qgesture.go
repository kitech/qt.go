package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qgesture.h
// dst-file: /src/widgets/qgesture.go
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
// class sizeof(QSwipeGesture)=1
type QSwipeGesture struct {
  /*qbase*/ QGesture;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QGesture)=1
type QGesture struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QGestureEvent)=1
type QGestureEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QPanGesture)=1
type QPanGesture struct {
  /*qbase*/ QGesture;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTapAndHoldGesture)=1
type QTapAndHoldGesture struct {
  /*qbase*/ QGesture;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTapGesture)=1
type QTapGesture struct {
  /*qbase*/ QGesture;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QPinchGesture)=1
type QPinchGesture struct {
  /*qbase*/ QGesture;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QSwipeGesture) FreeQSwipeGesture(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSwipeGesture", "~QSwipeGesture", args)
  }

}


func (this *QSwipeGesture) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QSwipeGesture10metaObjectEv
  default:
    qtrt.ErrorResolve("QSwipeGesture", "metaObject", args)
  }

}


func (this *QSwipeGesture) setSwipeAngle(args ...interface{}) () {
  // setSwipeAngle(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSwipeGesture13setSwipeAngleEd
  default:
    qtrt.ErrorResolve("QSwipeGesture", "setSwipeAngle", args)
  }

}


func NewQSwipeGesture(args ...interface{}) QSwipeGesture {
  return QSwipeGesture{}
}


func (this *QSwipeGesture) swipeAngle(args ...interface{}) () {
  // swipeAngle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QSwipeGesture10swipeAngleEv
  default:
    qtrt.ErrorResolve("QSwipeGesture", "swipeAngle", args)
  }

}


func (this *QGesture) hotSpot(args ...interface{}) () {
  // hotSpot()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QGesture7hotSpotEv
  default:
    qtrt.ErrorResolve("QGesture", "hotSpot", args)
  }

}


func (this *QGesture) hasHotSpot(args ...interface{}) () {
  // hasHotSpot()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QGesture10hasHotSpotEv
  default:
    qtrt.ErrorResolve("QGesture", "hasHotSpot", args)
  }

}


func (this *QGesture) unsetHotSpot(args ...interface{}) () {
  // unsetHotSpot()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QGesture12unsetHotSpotEv
  default:
    qtrt.ErrorResolve("QGesture", "unsetHotSpot", args)
  }

}


func (this *QGesture) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QGesture10metaObjectEv
  default:
    qtrt.ErrorResolve("QGesture", "metaObject", args)
  }

}


func NewQGesture(args ...interface{}) QGesture {
  return QGesture{}
}


func (this *QGesture) setHotSpot(args ...interface{}) () {
  // setHotSpot(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QGesture10setHotSpotERK7QPointF
  default:
    qtrt.ErrorResolve("QGesture", "setHotSpot", args)
  }

}


func (this *QGesture) FreeQGesture(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGesture", "~QGesture", args)
  }

}


func (this *QGestureEvent) isAccepted(args ...interface{}) () {
  // isAccepted(class QGesture *)
  // isAccepted(Qt::GestureType)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGesture{}) // "QGesture *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "Qt::GestureType"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGestureEvent10isAcceptedEP8QGesture
  case 1:
    // invoke: _ZNK13QGestureEvent10isAcceptedEN2Qt11GestureTypeE
  default:
    qtrt.ErrorResolve("QGestureEvent", "isAccepted", args)
  }

}


func (this *QGestureEvent) widget(args ...interface{}) () {
  // widget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGestureEvent6widgetEv
  default:
    qtrt.ErrorResolve("QGestureEvent", "widget", args)
  }

}


func (this *QGestureEvent) ignore(args ...interface{}) () {
  // ignore(class QGesture *)
  // ignore(Qt::GestureType)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGesture{}) // "QGesture *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "Qt::GestureType"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGestureEvent6ignoreEP8QGesture
  case 1:
    // invoke: _ZN13QGestureEvent6ignoreEN2Qt11GestureTypeE
  default:
    qtrt.ErrorResolve("QGestureEvent", "ignore", args)
  }

}


func (this *QGestureEvent) accept(args ...interface{}) () {
  // accept(class QGesture *)
  // accept(Qt::GestureType)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGesture{}) // "QGesture *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "Qt::GestureType"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGestureEvent6acceptEP8QGesture
  case 1:
    // invoke: _ZN13QGestureEvent6acceptEN2Qt11GestureTypeE
  default:
    qtrt.ErrorResolve("QGestureEvent", "accept", args)
  }

}


func (this *QGestureEvent) activeGestures(args ...interface{}) () {
  // activeGestures()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGestureEvent14activeGesturesEv
  default:
    qtrt.ErrorResolve("QGestureEvent", "activeGestures", args)
  }

}


func (this *QGestureEvent) gestures(args ...interface{}) () {
  // gestures()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGestureEvent8gesturesEv
  default:
    qtrt.ErrorResolve("QGestureEvent", "gestures", args)
  }

}


func (this *QGestureEvent) setAccepted(args ...interface{}) () {
  // setAccepted(Qt::GestureType, _Bool)
  // setAccepted(class QGesture *, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "Qt::GestureType"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QGesture{}) // "QGesture *"
  vtys[1][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGestureEvent11setAcceptedEN2Qt11GestureTypeEb
  case 1:
    // invoke: _ZN13QGestureEvent11setAcceptedEP8QGestureb
  default:
    qtrt.ErrorResolve("QGestureEvent", "setAccepted", args)
  }

}


func (this *QGestureEvent) setWidget(args ...interface{}) () {
  // setWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGestureEvent9setWidgetEP7QWidget
  default:
    qtrt.ErrorResolve("QGestureEvent", "setWidget", args)
  }

}


func (this *QGestureEvent) FreeQGestureEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGestureEvent", "~QGestureEvent", args)
  }

}


func (this *QGestureEvent) canceledGestures(args ...interface{}) () {
  // canceledGestures()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGestureEvent16canceledGesturesEv
  default:
    qtrt.ErrorResolve("QGestureEvent", "canceledGestures", args)
  }

}


func (this *QGestureEvent) mapToGraphicsScene(args ...interface{}) () {
  // mapToGraphicsScene(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGestureEvent18mapToGraphicsSceneERK7QPointF
  default:
    qtrt.ErrorResolve("QGestureEvent", "mapToGraphicsScene", args)
  }

}


func (this *QPanGesture) offset(args ...interface{}) () {
  // offset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPanGesture6offsetEv
  default:
    qtrt.ErrorResolve("QPanGesture", "offset", args)
  }

}


func (this *QPanGesture) delta(args ...interface{}) () {
  // delta()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPanGesture5deltaEv
  default:
    qtrt.ErrorResolve("QPanGesture", "delta", args)
  }

}


func (this *QPanGesture) setOffset(args ...interface{}) () {
  // setOffset(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPanGesture9setOffsetERK7QPointF
  default:
    qtrt.ErrorResolve("QPanGesture", "setOffset", args)
  }

}


func (this *QPanGesture) acceleration(args ...interface{}) () {
  // acceleration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPanGesture12accelerationEv
  default:
    qtrt.ErrorResolve("QPanGesture", "acceleration", args)
  }

}


func (this *QPanGesture) FreeQPanGesture(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPanGesture", "~QPanGesture", args)
  }

}


func NewQPanGesture(args ...interface{}) QPanGesture {
  return QPanGesture{}
}


func (this *QPanGesture) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPanGesture10metaObjectEv
  default:
    qtrt.ErrorResolve("QPanGesture", "metaObject", args)
  }

}


func (this *QPanGesture) setAcceleration(args ...interface{}) () {
  // setAcceleration(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPanGesture15setAccelerationEd
  default:
    qtrt.ErrorResolve("QPanGesture", "setAcceleration", args)
  }

}


func (this *QPanGesture) lastOffset(args ...interface{}) () {
  // lastOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPanGesture10lastOffsetEv
  default:
    qtrt.ErrorResolve("QPanGesture", "lastOffset", args)
  }

}


func (this *QPanGesture) setLastOffset(args ...interface{}) () {
  // setLastOffset(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPanGesture13setLastOffsetERK7QPointF
  default:
    qtrt.ErrorResolve("QPanGesture", "setLastOffset", args)
  }

}


func NewQTapAndHoldGesture(args ...interface{}) QTapAndHoldGesture {
  return QTapAndHoldGesture{}
}


func (this *QTapAndHoldGesture) FreeQTapAndHoldGesture(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTapAndHoldGesture", "~QTapAndHoldGesture", args)
  }

}


func (this *QTapAndHoldGesture) position(args ...interface{}) () {
  // position()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QTapAndHoldGesture8positionEv
  default:
    qtrt.ErrorResolve("QTapAndHoldGesture", "position", args)
  }

}


func (this *QTapAndHoldGesture) setTimeout_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTapAndHoldGesture", "setTimeout", args)
  }

}


func (this *QTapAndHoldGesture) timeout_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTapAndHoldGesture", "timeout", args)
  }

}


func (this *QTapAndHoldGesture) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QTapAndHoldGesture10metaObjectEv
  default:
    qtrt.ErrorResolve("QTapAndHoldGesture", "metaObject", args)
  }

}


func (this *QTapAndHoldGesture) setPosition(args ...interface{}) () {
  // setPosition(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QTapAndHoldGesture11setPositionERK7QPointF
  default:
    qtrt.ErrorResolve("QTapAndHoldGesture", "setPosition", args)
  }

}


func (this *QTapGesture) position(args ...interface{}) () {
  // position()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTapGesture8positionEv
  default:
    qtrt.ErrorResolve("QTapGesture", "position", args)
  }

}


func (this *QTapGesture) setPosition(args ...interface{}) () {
  // setPosition(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTapGesture11setPositionERK7QPointF
  default:
    qtrt.ErrorResolve("QTapGesture", "setPosition", args)
  }

}


func NewQTapGesture(args ...interface{}) QTapGesture {
  return QTapGesture{}
}


func (this *QTapGesture) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTapGesture10metaObjectEv
  default:
    qtrt.ErrorResolve("QTapGesture", "metaObject", args)
  }

}


func (this *QTapGesture) FreeQTapGesture(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTapGesture", "~QTapGesture", args)
  }

}


func (this *QPinchGesture) setRotationAngle(args ...interface{}) () {
  // setRotationAngle(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPinchGesture16setRotationAngleEd
  default:
    qtrt.ErrorResolve("QPinchGesture", "setRotationAngle", args)
  }

}


func (this *QPinchGesture) lastScaleFactor(args ...interface{}) () {
  // lastScaleFactor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPinchGesture15lastScaleFactorEv
  default:
    qtrt.ErrorResolve("QPinchGesture", "lastScaleFactor", args)
  }

}


func (this *QPinchGesture) lastRotationAngle(args ...interface{}) () {
  // lastRotationAngle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPinchGesture17lastRotationAngleEv
  default:
    qtrt.ErrorResolve("QPinchGesture", "lastRotationAngle", args)
  }

}


func (this *QPinchGesture) startCenterPoint(args ...interface{}) () {
  // startCenterPoint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPinchGesture16startCenterPointEv
  default:
    qtrt.ErrorResolve("QPinchGesture", "startCenterPoint", args)
  }

}


func (this *QPinchGesture) rotationAngle(args ...interface{}) () {
  // rotationAngle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPinchGesture13rotationAngleEv
  default:
    qtrt.ErrorResolve("QPinchGesture", "rotationAngle", args)
  }

}


func (this *QPinchGesture) lastCenterPoint(args ...interface{}) () {
  // lastCenterPoint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPinchGesture15lastCenterPointEv
  default:
    qtrt.ErrorResolve("QPinchGesture", "lastCenterPoint", args)
  }

}


func NewQPinchGesture(args ...interface{}) QPinchGesture {
  return QPinchGesture{}
}


func (this *QPinchGesture) totalScaleFactor(args ...interface{}) () {
  // totalScaleFactor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPinchGesture16totalScaleFactorEv
  default:
    qtrt.ErrorResolve("QPinchGesture", "totalScaleFactor", args)
  }

}


func (this *QPinchGesture) setTotalScaleFactor(args ...interface{}) () {
  // setTotalScaleFactor(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPinchGesture19setTotalScaleFactorEd
  default:
    qtrt.ErrorResolve("QPinchGesture", "setTotalScaleFactor", args)
  }

}


func (this *QPinchGesture) totalRotationAngle(args ...interface{}) () {
  // totalRotationAngle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPinchGesture18totalRotationAngleEv
  default:
    qtrt.ErrorResolve("QPinchGesture", "totalRotationAngle", args)
  }

}


func (this *QPinchGesture) setLastScaleFactor(args ...interface{}) () {
  // setLastScaleFactor(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPinchGesture18setLastScaleFactorEd
  default:
    qtrt.ErrorResolve("QPinchGesture", "setLastScaleFactor", args)
  }

}


func (this *QPinchGesture) setLastCenterPoint(args ...interface{}) () {
  // setLastCenterPoint(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPinchGesture18setLastCenterPointERK7QPointF
  default:
    qtrt.ErrorResolve("QPinchGesture", "setLastCenterPoint", args)
  }

}


func (this *QPinchGesture) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPinchGesture10metaObjectEv
  default:
    qtrt.ErrorResolve("QPinchGesture", "metaObject", args)
  }

}


func (this *QPinchGesture) setLastRotationAngle(args ...interface{}) () {
  // setLastRotationAngle(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPinchGesture20setLastRotationAngleEd
  default:
    qtrt.ErrorResolve("QPinchGesture", "setLastRotationAngle", args)
  }

}


func (this *QPinchGesture) centerPoint(args ...interface{}) () {
  // centerPoint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPinchGesture11centerPointEv
  default:
    qtrt.ErrorResolve("QPinchGesture", "centerPoint", args)
  }

}


func (this *QPinchGesture) setCenterPoint(args ...interface{}) () {
  // setCenterPoint(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPinchGesture14setCenterPointERK7QPointF
  default:
    qtrt.ErrorResolve("QPinchGesture", "setCenterPoint", args)
  }

}


func (this *QPinchGesture) setTotalRotationAngle(args ...interface{}) () {
  // setTotalRotationAngle(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPinchGesture21setTotalRotationAngleEd
  default:
    qtrt.ErrorResolve("QPinchGesture", "setTotalRotationAngle", args)
  }

}


func (this *QPinchGesture) setScaleFactor(args ...interface{}) () {
  // setScaleFactor(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPinchGesture14setScaleFactorEd
  default:
    qtrt.ErrorResolve("QPinchGesture", "setScaleFactor", args)
  }

}


func (this *QPinchGesture) FreeQPinchGesture(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPinchGesture", "~QPinchGesture", args)
  }

}


func (this *QPinchGesture) setStartCenterPoint(args ...interface{}) () {
  // setStartCenterPoint(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPinchGesture19setStartCenterPointERK7QPointF
  default:
    qtrt.ErrorResolve("QPinchGesture", "setStartCenterPoint", args)
  }

}


func (this *QPinchGesture) scaleFactor(args ...interface{}) () {
  // scaleFactor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPinchGesture11scaleFactorEv
  default:
    qtrt.ErrorResolve("QPinchGesture", "scaleFactor", args)
  }

}

// <= body block end

