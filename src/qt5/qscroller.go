package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qscroller.h
// dst-file: /src/widgets/qscroller.go
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
// class sizeof(QScroller)=1
type QScroller struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _stateChanged QScroller_stateChanged_signal;
//  _scrollerPropertiesChanged QScroller_scrollerPropertiesChanged_signal;
}


func (this *QScroller) FreeQScroller(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QScroller", "~QScroller", args)
  }

}


func (this *QScroller) setSnapPositionsY(args ...interface{}) () {
  // setSnapPositionsY(qreal, qreal)
  // setSnapPositionsY(const QList<qreal> &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  // vtys[1][0] = reflect.TypeOf(QList<double>{}) // "const QList<qreal> &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QScroller17setSnapPositionsYEdd
  case 1:
    // invoke: _ZN9QScroller17setSnapPositionsYERK5QListIdE
  default:
    qtrt.ErrorResolve("QScroller", "setSnapPositionsY", args)
  }

}


func (this *QScroller) finalPosition(args ...interface{}) () {
  // finalPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QScroller13finalPositionEv
  default:
    qtrt.ErrorResolve("QScroller", "finalPosition", args)
  }

}


func (this *QScroller) activeScrollers_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QScroller", "activeScrollers", args)
  }

}


func (this *QScroller) setScrollerProperties(args ...interface{}) () {
  // setScrollerProperties(const class QScrollerProperties &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QScrollerProperties{}) // "const QScrollerProperties &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QScroller21setScrollerPropertiesERK19QScrollerProperties
  default:
    qtrt.ErrorResolve("QScroller", "setScrollerProperties", args)
  }

}


func (this *QScroller) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QScroller10metaObjectEv
  default:
    qtrt.ErrorResolve("QScroller", "metaObject", args)
  }

}


func (this *QScroller) velocity(args ...interface{}) () {
  // velocity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QScroller8velocityEv
  default:
    qtrt.ErrorResolve("QScroller", "velocity", args)
  }

}


func (this *QScroller) resendPrepareEvent(args ...interface{}) () {
  // resendPrepareEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QScroller18resendPrepareEventEv
  default:
    qtrt.ErrorResolve("QScroller", "resendPrepareEvent", args)
  }

}


func (this *QScroller) ensureVisible(args ...interface{}) () {
  // ensureVisible(const class QRectF &, qreal, qreal)
  // ensureVisible(const class QRectF &, qreal, qreal, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QScroller13ensureVisibleERK6QRectFdd
  case 1:
    // invoke: _ZN9QScroller13ensureVisibleERK6QRectFddi
  default:
    qtrt.ErrorResolve("QScroller", "ensureVisible", args)
  }

}


func (this *QScroller) setSnapPositionsX(args ...interface{}) () {
  // setSnapPositionsX(qreal, qreal)
  // setSnapPositionsX(const QList<qreal> &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  // vtys[1][0] = reflect.TypeOf(QList<double>{}) // "const QList<qreal> &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QScroller17setSnapPositionsXEdd
  case 1:
    // invoke: _ZN9QScroller17setSnapPositionsXERK5QListIdE
  default:
    qtrt.ErrorResolve("QScroller", "setSnapPositionsX", args)
  }

}


func (this *QScroller) hasScroller_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QScroller", "hasScroller", args)
  }

}


func NewQScroller(args ...interface{}) QScroller {
  return QScroller{}
}


func (this *QScroller) scrollTo(args ...interface{}) () {
  // scrollTo(const class QPointF &, int)
  // scrollTo(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QScroller8scrollToERK7QPointFi
  case 1:
    // invoke: _ZN9QScroller8scrollToERK7QPointF
  default:
    qtrt.ErrorResolve("QScroller", "scrollTo", args)
  }

}


func (this *QScroller) stop(args ...interface{}) () {
  // stop()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QScroller4stopEv
  default:
    qtrt.ErrorResolve("QScroller", "stop", args)
  }

}


func (this *QScroller) ungrabGesture_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QScroller", "ungrabGesture", args)
  }

}


func (this *QScroller) scrollerProperties(args ...interface{}) () {
  // scrollerProperties()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QScroller18scrollerPropertiesEv
  default:
    qtrt.ErrorResolve("QScroller", "scrollerProperties", args)
  }

}


func (this *QScroller) scroller_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QScroller", "scroller", args)
  }

}


func (this *QScroller) target(args ...interface{}) () {
  // target()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QScroller6targetEv
  default:
    qtrt.ErrorResolve("QScroller", "target", args)
  }

}


func (this *QScroller) pixelPerMeter(args ...interface{}) () {
  // pixelPerMeter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QScroller13pixelPerMeterEv
  default:
    qtrt.ErrorResolve("QScroller", "pixelPerMeter", args)
  }

}

// <= body block end

