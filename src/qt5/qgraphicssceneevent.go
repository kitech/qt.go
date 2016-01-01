package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qgraphicssceneevent.h
// dst-file: /src/widgets/qgraphicssceneevent.go
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
// class sizeof(QGraphicsSceneMoveEvent)=1
type QGraphicsSceneMoveEvent struct {
  /*qbase*/ QGraphicsSceneEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QGraphicsSceneContextMenuEvent)=1
type QGraphicsSceneContextMenuEvent struct {
  /*qbase*/ QGraphicsSceneEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QGraphicsSceneMouseEvent)=1
type QGraphicsSceneMouseEvent struct {
  /*qbase*/ QGraphicsSceneEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QGraphicsSceneHelpEvent)=1
type QGraphicsSceneHelpEvent struct {
  /*qbase*/ QGraphicsSceneEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QGraphicsSceneHoverEvent)=1
type QGraphicsSceneHoverEvent struct {
  /*qbase*/ QGraphicsSceneEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QGraphicsSceneWheelEvent)=1
type QGraphicsSceneWheelEvent struct {
  /*qbase*/ QGraphicsSceneEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QGraphicsSceneDragDropEvent)=1
type QGraphicsSceneDragDropEvent struct {
  /*qbase*/ QGraphicsSceneEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QGraphicsSceneEvent)=1
type QGraphicsSceneEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QGraphicsSceneResizeEvent)=1
type QGraphicsSceneResizeEvent struct {
  /*qbase*/ QGraphicsSceneEvent;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QGraphicsSceneMoveEvent) newPos(args ...interface{}) () {
  // newPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsSceneMoveEvent6newPosEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneMoveEvent", "newPos", args)
  }

}


func (this *QGraphicsSceneMoveEvent) oldPos(args ...interface{}) () {
  // oldPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsSceneMoveEvent6oldPosEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneMoveEvent", "oldPos", args)
  }

}


func (this *QGraphicsSceneMoveEvent) FreeQGraphicsSceneMoveEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsSceneMoveEvent", "~QGraphicsSceneMoveEvent", args)
  }

}


func NewQGraphicsSceneMoveEvent(args ...interface{}) QGraphicsSceneMoveEvent {
  return QGraphicsSceneMoveEvent{}
}


func (this *QGraphicsSceneMoveEvent) setNewPos(args ...interface{}) () {
  // setNewPos(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QGraphicsSceneMoveEvent9setNewPosERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsSceneMoveEvent", "setNewPos", args)
  }

}


func (this *QGraphicsSceneMoveEvent) setOldPos(args ...interface{}) () {
  // setOldPos(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QGraphicsSceneMoveEvent9setOldPosERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsSceneMoveEvent", "setOldPos", args)
  }

}


func NewQGraphicsSceneContextMenuEvent(args ...interface{}) QGraphicsSceneContextMenuEvent {
  return QGraphicsSceneContextMenuEvent{}
}


func (this *QGraphicsSceneContextMenuEvent) scenePos(args ...interface{}) () {
  // scenePos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK30QGraphicsSceneContextMenuEvent8scenePosEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneContextMenuEvent", "scenePos", args)
  }

}


func (this *QGraphicsSceneContextMenuEvent) FreeQGraphicsSceneContextMenuEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsSceneContextMenuEvent", "~QGraphicsSceneContextMenuEvent", args)
  }

}


func (this *QGraphicsSceneContextMenuEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK30QGraphicsSceneContextMenuEvent3posEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneContextMenuEvent", "pos", args)
  }

}


func (this *QGraphicsSceneContextMenuEvent) setScreenPos(args ...interface{}) () {
  // setScreenPos(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN30QGraphicsSceneContextMenuEvent12setScreenPosERK6QPoint
  default:
    qtrt.ErrorResolve("QGraphicsSceneContextMenuEvent", "setScreenPos", args)
  }

}


func (this *QGraphicsSceneContextMenuEvent) setPos(args ...interface{}) () {
  // setPos(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN30QGraphicsSceneContextMenuEvent6setPosERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsSceneContextMenuEvent", "setPos", args)
  }

}


func (this *QGraphicsSceneContextMenuEvent) screenPos(args ...interface{}) () {
  // screenPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK30QGraphicsSceneContextMenuEvent9screenPosEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneContextMenuEvent", "screenPos", args)
  }

}


func (this *QGraphicsSceneContextMenuEvent) setScenePos(args ...interface{}) () {
  // setScenePos(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN30QGraphicsSceneContextMenuEvent11setScenePosERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsSceneContextMenuEvent", "setScenePos", args)
  }

}


func (this *QGraphicsSceneMouseEvent) screenPos(args ...interface{}) () {
  // screenPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QGraphicsSceneMouseEvent9screenPosEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "screenPos", args)
  }

}


func NewQGraphicsSceneMouseEvent(args ...interface{}) QGraphicsSceneMouseEvent {
  return QGraphicsSceneMouseEvent{}
}


func (this *QGraphicsSceneMouseEvent) lastScenePos(args ...interface{}) () {
  // lastScenePos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QGraphicsSceneMouseEvent12lastScenePosEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "lastScenePos", args)
  }

}


func (this *QGraphicsSceneMouseEvent) FreeQGraphicsSceneMouseEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "~QGraphicsSceneMouseEvent", args)
  }

}


func (this *QGraphicsSceneMouseEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QGraphicsSceneMouseEvent3posEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "pos", args)
  }

}


func (this *QGraphicsSceneMouseEvent) setLastPos(args ...interface{}) () {
  // setLastPos(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QGraphicsSceneMouseEvent10setLastPosERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "setLastPos", args)
  }

}


func (this *QGraphicsSceneMouseEvent) setLastScenePos(args ...interface{}) () {
  // setLastScenePos(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QGraphicsSceneMouseEvent15setLastScenePosERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "setLastScenePos", args)
  }

}


func (this *QGraphicsSceneMouseEvent) lastScreenPos(args ...interface{}) () {
  // lastScreenPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QGraphicsSceneMouseEvent13lastScreenPosEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "lastScreenPos", args)
  }

}


func (this *QGraphicsSceneMouseEvent) setScreenPos(args ...interface{}) () {
  // setScreenPos(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QGraphicsSceneMouseEvent12setScreenPosERK6QPoint
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "setScreenPos", args)
  }

}


func (this *QGraphicsSceneMouseEvent) setLastScreenPos(args ...interface{}) () {
  // setLastScreenPos(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QGraphicsSceneMouseEvent16setLastScreenPosERK6QPoint
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "setLastScreenPos", args)
  }

}


func (this *QGraphicsSceneMouseEvent) setScenePos(args ...interface{}) () {
  // setScenePos(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QGraphicsSceneMouseEvent11setScenePosERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "setScenePos", args)
  }

}


func (this *QGraphicsSceneMouseEvent) lastPos(args ...interface{}) () {
  // lastPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QGraphicsSceneMouseEvent7lastPosEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "lastPos", args)
  }

}


func (this *QGraphicsSceneMouseEvent) scenePos(args ...interface{}) () {
  // scenePos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QGraphicsSceneMouseEvent8scenePosEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "scenePos", args)
  }

}


func (this *QGraphicsSceneMouseEvent) setPos(args ...interface{}) () {
  // setPos(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QGraphicsSceneMouseEvent6setPosERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "setPos", args)
  }

}


func (this *QGraphicsSceneHelpEvent) setScenePos(args ...interface{}) () {
  // setScenePos(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QGraphicsSceneHelpEvent11setScenePosERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsSceneHelpEvent", "setScenePos", args)
  }

}


func (this *QGraphicsSceneHelpEvent) screenPos(args ...interface{}) () {
  // screenPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsSceneHelpEvent9screenPosEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneHelpEvent", "screenPos", args)
  }

}


func (this *QGraphicsSceneHelpEvent) FreeQGraphicsSceneHelpEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsSceneHelpEvent", "~QGraphicsSceneHelpEvent", args)
  }

}


func NewQGraphicsSceneHelpEvent(args ...interface{}) QGraphicsSceneHelpEvent {
  return QGraphicsSceneHelpEvent{}
}


func (this *QGraphicsSceneHelpEvent) setScreenPos(args ...interface{}) () {
  // setScreenPos(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QGraphicsSceneHelpEvent12setScreenPosERK6QPoint
  default:
    qtrt.ErrorResolve("QGraphicsSceneHelpEvent", "setScreenPos", args)
  }

}


func (this *QGraphicsSceneHelpEvent) scenePos(args ...interface{}) () {
  // scenePos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsSceneHelpEvent8scenePosEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneHelpEvent", "scenePos", args)
  }

}


func (this *QGraphicsSceneHoverEvent) scenePos(args ...interface{}) () {
  // scenePos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QGraphicsSceneHoverEvent8scenePosEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "scenePos", args)
  }

}


func (this *QGraphicsSceneHoverEvent) setLastPos(args ...interface{}) () {
  // setLastPos(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QGraphicsSceneHoverEvent10setLastPosERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "setLastPos", args)
  }

}


func (this *QGraphicsSceneHoverEvent) lastPos(args ...interface{}) () {
  // lastPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QGraphicsSceneHoverEvent7lastPosEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "lastPos", args)
  }

}


func NewQGraphicsSceneHoverEvent(args ...interface{}) QGraphicsSceneHoverEvent {
  return QGraphicsSceneHoverEvent{}
}


func (this *QGraphicsSceneHoverEvent) lastScenePos(args ...interface{}) () {
  // lastScenePos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QGraphicsSceneHoverEvent12lastScenePosEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "lastScenePos", args)
  }

}


func (this *QGraphicsSceneHoverEvent) setLastScreenPos(args ...interface{}) () {
  // setLastScreenPos(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QGraphicsSceneHoverEvent16setLastScreenPosERK6QPoint
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "setLastScreenPos", args)
  }

}


func (this *QGraphicsSceneHoverEvent) setScenePos(args ...interface{}) () {
  // setScenePos(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QGraphicsSceneHoverEvent11setScenePosERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "setScenePos", args)
  }

}


func (this *QGraphicsSceneHoverEvent) setPos(args ...interface{}) () {
  // setPos(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QGraphicsSceneHoverEvent6setPosERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "setPos", args)
  }

}


func (this *QGraphicsSceneHoverEvent) screenPos(args ...interface{}) () {
  // screenPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QGraphicsSceneHoverEvent9screenPosEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "screenPos", args)
  }

}


func (this *QGraphicsSceneHoverEvent) lastScreenPos(args ...interface{}) () {
  // lastScreenPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QGraphicsSceneHoverEvent13lastScreenPosEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "lastScreenPos", args)
  }

}


func (this *QGraphicsSceneHoverEvent) setLastScenePos(args ...interface{}) () {
  // setLastScenePos(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QGraphicsSceneHoverEvent15setLastScenePosERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "setLastScenePos", args)
  }

}


func (this *QGraphicsSceneHoverEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QGraphicsSceneHoverEvent3posEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "pos", args)
  }

}


func (this *QGraphicsSceneHoverEvent) setScreenPos(args ...interface{}) () {
  // setScreenPos(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QGraphicsSceneHoverEvent12setScreenPosERK6QPoint
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "setScreenPos", args)
  }

}


func (this *QGraphicsSceneHoverEvent) FreeQGraphicsSceneHoverEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "~QGraphicsSceneHoverEvent", args)
  }

}


func (this *QGraphicsSceneWheelEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QGraphicsSceneWheelEvent3posEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneWheelEvent", "pos", args)
  }

}


func (this *QGraphicsSceneWheelEvent) FreeQGraphicsSceneWheelEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsSceneWheelEvent", "~QGraphicsSceneWheelEvent", args)
  }

}


func (this *QGraphicsSceneWheelEvent) setDelta(args ...interface{}) () {
  // setDelta(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QGraphicsSceneWheelEvent8setDeltaEi
  default:
    qtrt.ErrorResolve("QGraphicsSceneWheelEvent", "setDelta", args)
  }

}


func (this *QGraphicsSceneWheelEvent) setScenePos(args ...interface{}) () {
  // setScenePos(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QGraphicsSceneWheelEvent11setScenePosERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsSceneWheelEvent", "setScenePos", args)
  }

}


func NewQGraphicsSceneWheelEvent(args ...interface{}) QGraphicsSceneWheelEvent {
  return QGraphicsSceneWheelEvent{}
}


func (this *QGraphicsSceneWheelEvent) setPos(args ...interface{}) () {
  // setPos(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QGraphicsSceneWheelEvent6setPosERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsSceneWheelEvent", "setPos", args)
  }

}


func (this *QGraphicsSceneWheelEvent) setScreenPos(args ...interface{}) () {
  // setScreenPos(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QGraphicsSceneWheelEvent12setScreenPosERK6QPoint
  default:
    qtrt.ErrorResolve("QGraphicsSceneWheelEvent", "setScreenPos", args)
  }

}


func (this *QGraphicsSceneWheelEvent) delta(args ...interface{}) () {
  // delta()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QGraphicsSceneWheelEvent5deltaEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneWheelEvent", "delta", args)
  }

}


func (this *QGraphicsSceneWheelEvent) scenePos(args ...interface{}) () {
  // scenePos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QGraphicsSceneWheelEvent8scenePosEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneWheelEvent", "scenePos", args)
  }

}


func (this *QGraphicsSceneWheelEvent) screenPos(args ...interface{}) () {
  // screenPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QGraphicsSceneWheelEvent9screenPosEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneWheelEvent", "screenPos", args)
  }

}


func (this *QGraphicsSceneDragDropEvent) source(args ...interface{}) () {
  // source()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QGraphicsSceneDragDropEvent6sourceEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "source", args)
  }

}


func (this *QGraphicsSceneDragDropEvent) scenePos(args ...interface{}) () {
  // scenePos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QGraphicsSceneDragDropEvent8scenePosEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "scenePos", args)
  }

}


func (this *QGraphicsSceneDragDropEvent) setPos(args ...interface{}) () {
  // setPos(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QGraphicsSceneDragDropEvent6setPosERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "setPos", args)
  }

}


func NewQGraphicsSceneDragDropEvent(args ...interface{}) QGraphicsSceneDragDropEvent {
  return QGraphicsSceneDragDropEvent{}
}


func (this *QGraphicsSceneDragDropEvent) setScreenPos(args ...interface{}) () {
  // setScreenPos(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QGraphicsSceneDragDropEvent12setScreenPosERK6QPoint
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "setScreenPos", args)
  }

}


func (this *QGraphicsSceneDragDropEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QGraphicsSceneDragDropEvent3posEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "pos", args)
  }

}


func (this *QGraphicsSceneDragDropEvent) screenPos(args ...interface{}) () {
  // screenPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QGraphicsSceneDragDropEvent9screenPosEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "screenPos", args)
  }

}


func (this *QGraphicsSceneDragDropEvent) mimeData(args ...interface{}) () {
  // mimeData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QGraphicsSceneDragDropEvent8mimeDataEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "mimeData", args)
  }

}


func (this *QGraphicsSceneDragDropEvent) FreeQGraphicsSceneDragDropEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "~QGraphicsSceneDragDropEvent", args)
  }

}


func (this *QGraphicsSceneDragDropEvent) setMimeData(args ...interface{}) () {
  // setMimeData(const class QMimeData *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMimeData{}) // "const QMimeData *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QGraphicsSceneDragDropEvent11setMimeDataEPK9QMimeData
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "setMimeData", args)
  }

}


func (this *QGraphicsSceneDragDropEvent) setSource(args ...interface{}) () {
  // setSource(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QGraphicsSceneDragDropEvent9setSourceEP7QWidget
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "setSource", args)
  }

}


func (this *QGraphicsSceneDragDropEvent) setScenePos(args ...interface{}) () {
  // setScenePos(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QGraphicsSceneDragDropEvent11setScenePosERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "setScenePos", args)
  }

}


func (this *QGraphicsSceneDragDropEvent) acceptProposedAction(args ...interface{}) () {
  // acceptProposedAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QGraphicsSceneDragDropEvent20acceptProposedActionEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "acceptProposedAction", args)
  }

}


func NewQGraphicsSceneEvent(args ...interface{}) QGraphicsSceneEvent {
  return QGraphicsSceneEvent{}
}


func (this *QGraphicsSceneEvent) widget(args ...interface{}) () {
  // widget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsSceneEvent6widgetEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneEvent", "widget", args)
  }

}


func (this *QGraphicsSceneEvent) setWidget(args ...interface{}) () {
  // setWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsSceneEvent9setWidgetEP7QWidget
  default:
    qtrt.ErrorResolve("QGraphicsSceneEvent", "setWidget", args)
  }

}


func (this *QGraphicsSceneEvent) FreeQGraphicsSceneEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsSceneEvent", "~QGraphicsSceneEvent", args)
  }

}


func (this *QGraphicsSceneResizeEvent) newSize(args ...interface{}) () {
  // newSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QGraphicsSceneResizeEvent7newSizeEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneResizeEvent", "newSize", args)
  }

}


func (this *QGraphicsSceneResizeEvent) oldSize(args ...interface{}) () {
  // oldSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QGraphicsSceneResizeEvent7oldSizeEv
  default:
    qtrt.ErrorResolve("QGraphicsSceneResizeEvent", "oldSize", args)
  }

}


func (this *QGraphicsSceneResizeEvent) FreeQGraphicsSceneResizeEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsSceneResizeEvent", "~QGraphicsSceneResizeEvent", args)
  }

}


func (this *QGraphicsSceneResizeEvent) setNewSize(args ...interface{}) () {
  // setNewSize(const class QSizeF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QGraphicsSceneResizeEvent10setNewSizeERK6QSizeF
  default:
    qtrt.ErrorResolve("QGraphicsSceneResizeEvent", "setNewSize", args)
  }

}


func NewQGraphicsSceneResizeEvent(args ...interface{}) QGraphicsSceneResizeEvent {
  return QGraphicsSceneResizeEvent{}
}


func (this *QGraphicsSceneResizeEvent) setOldSize(args ...interface{}) () {
  // setOldSize(const class QSizeF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QGraphicsSceneResizeEvent10setOldSizeERK6QSizeF
  default:
    qtrt.ErrorResolve("QGraphicsSceneResizeEvent", "setOldSize", args)
  }

}

// <= body block end

