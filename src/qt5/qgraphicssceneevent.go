package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtWidgets/qgraphicssceneevent.h
// dst-file: /src/widgets/qgraphicssceneevent.go
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
  // proto:  QPointF QGraphicsSceneMoveEvent::newPos();
extern void _ZNK23QGraphicsSceneMoveEvent6newPosEv(void* qthis);
  // proto:  QPointF QGraphicsSceneMoveEvent::oldPos();
extern void _ZNK23QGraphicsSceneMoveEvent6oldPosEv(void* qthis);
  // proto:  void QGraphicsSceneMoveEvent::~QGraphicsSceneMoveEvent();
extern void _ZN23QGraphicsSceneMoveEventD0Ev(void* qthis);
  // proto:  void QGraphicsSceneMoveEvent::QGraphicsSceneMoveEvent(const QGraphicsSceneMoveEvent & );
extern void* dector_ZN23QGraphicsSceneMoveEventC1ERKS_(void* arg0);
extern void _ZN23QGraphicsSceneMoveEventC1ERKS_(void* qthis, void* arg0);
  // proto:  void QGraphicsSceneMoveEvent::setNewPos(const QPointF & pos);
extern void _ZN23QGraphicsSceneMoveEvent9setNewPosERK7QPointF(void* qthis, void* arg0);
  // proto:  void QGraphicsSceneMoveEvent::QGraphicsSceneMoveEvent();
extern void* dector_ZN23QGraphicsSceneMoveEventC1Ev();
extern void _ZN23QGraphicsSceneMoveEventC1Ev(void* qthis);
  // proto:  void QGraphicsSceneMoveEvent::setOldPos(const QPointF & pos);
extern void _ZN23QGraphicsSceneMoveEvent9setOldPosERK7QPointF(void* qthis, void* arg0);
  // proto:  void QGraphicsSceneContextMenuEvent::QGraphicsSceneContextMenuEvent(const QGraphicsSceneContextMenuEvent & );
extern void* dector_ZN30QGraphicsSceneContextMenuEventC1ERKS_(void* arg0);
extern void _ZN30QGraphicsSceneContextMenuEventC1ERKS_(void* qthis, void* arg0);
  // proto:  QPointF QGraphicsSceneContextMenuEvent::scenePos();
extern void _ZNK30QGraphicsSceneContextMenuEvent8scenePosEv(void* qthis);
  // proto:  void QGraphicsSceneContextMenuEvent::~QGraphicsSceneContextMenuEvent();
extern void _ZN30QGraphicsSceneContextMenuEventD0Ev(void* qthis);
  // proto:  QPointF QGraphicsSceneContextMenuEvent::pos();
extern void _ZNK30QGraphicsSceneContextMenuEvent3posEv(void* qthis);
  // proto:  void QGraphicsSceneContextMenuEvent::setScreenPos(const QPoint & pos);
extern void _ZN30QGraphicsSceneContextMenuEvent12setScreenPosERK6QPoint(void* qthis, void* arg0);
  // proto:  void QGraphicsSceneContextMenuEvent::setPos(const QPointF & pos);
extern void _ZN30QGraphicsSceneContextMenuEvent6setPosERK7QPointF(void* qthis, void* arg0);
  // proto:  QPoint QGraphicsSceneContextMenuEvent::screenPos();
extern void _ZNK30QGraphicsSceneContextMenuEvent9screenPosEv(void* qthis);
  // proto:  void QGraphicsSceneContextMenuEvent::setScenePos(const QPointF & pos);
extern void _ZN30QGraphicsSceneContextMenuEvent11setScenePosERK7QPointF(void* qthis, void* arg0);
  // proto:  QPoint QGraphicsSceneMouseEvent::screenPos();
extern void _ZNK24QGraphicsSceneMouseEvent9screenPosEv(void* qthis);
  // proto:  void QGraphicsSceneMouseEvent::QGraphicsSceneMouseEvent(const QGraphicsSceneMouseEvent & );
extern void* dector_ZN24QGraphicsSceneMouseEventC1ERKS_(void* arg0);
extern void _ZN24QGraphicsSceneMouseEventC1ERKS_(void* qthis, void* arg0);
  // proto:  QPointF QGraphicsSceneMouseEvent::lastScenePos();
extern void _ZNK24QGraphicsSceneMouseEvent12lastScenePosEv(void* qthis);
  // proto:  void QGraphicsSceneMouseEvent::~QGraphicsSceneMouseEvent();
extern void _ZN24QGraphicsSceneMouseEventD0Ev(void* qthis);
  // proto:  QPointF QGraphicsSceneMouseEvent::pos();
extern void _ZNK24QGraphicsSceneMouseEvent3posEv(void* qthis);
  // proto:  void QGraphicsSceneMouseEvent::setLastPos(const QPointF & pos);
extern void _ZN24QGraphicsSceneMouseEvent10setLastPosERK7QPointF(void* qthis, void* arg0);
  // proto:  void QGraphicsSceneMouseEvent::setLastScenePos(const QPointF & pos);
extern void _ZN24QGraphicsSceneMouseEvent15setLastScenePosERK7QPointF(void* qthis, void* arg0);
  // proto:  QPoint QGraphicsSceneMouseEvent::lastScreenPos();
extern void _ZNK24QGraphicsSceneMouseEvent13lastScreenPosEv(void* qthis);
  // proto:  void QGraphicsSceneMouseEvent::setScreenPos(const QPoint & pos);
extern void _ZN24QGraphicsSceneMouseEvent12setScreenPosERK6QPoint(void* qthis, void* arg0);
  // proto:  void QGraphicsSceneMouseEvent::setLastScreenPos(const QPoint & pos);
extern void _ZN24QGraphicsSceneMouseEvent16setLastScreenPosERK6QPoint(void* qthis, void* arg0);
  // proto:  void QGraphicsSceneMouseEvent::setScenePos(const QPointF & pos);
extern void _ZN24QGraphicsSceneMouseEvent11setScenePosERK7QPointF(void* qthis, void* arg0);
  // proto:  QPointF QGraphicsSceneMouseEvent::lastPos();
extern void _ZNK24QGraphicsSceneMouseEvent7lastPosEv(void* qthis);
  // proto:  QPointF QGraphicsSceneMouseEvent::scenePos();
extern void _ZNK24QGraphicsSceneMouseEvent8scenePosEv(void* qthis);
  // proto:  void QGraphicsSceneMouseEvent::setPos(const QPointF & pos);
extern void _ZN24QGraphicsSceneMouseEvent6setPosERK7QPointF(void* qthis, void* arg0);
  // proto:  void QGraphicsSceneHelpEvent::setScenePos(const QPointF & pos);
extern void _ZN23QGraphicsSceneHelpEvent11setScenePosERK7QPointF(void* qthis, void* arg0);
  // proto:  QPoint QGraphicsSceneHelpEvent::screenPos();
extern void _ZNK23QGraphicsSceneHelpEvent9screenPosEv(void* qthis);
  // proto:  void QGraphicsSceneHelpEvent::~QGraphicsSceneHelpEvent();
extern void _ZN23QGraphicsSceneHelpEventD0Ev(void* qthis);
  // proto:  void QGraphicsSceneHelpEvent::QGraphicsSceneHelpEvent(const QGraphicsSceneHelpEvent & );
extern void* dector_ZN23QGraphicsSceneHelpEventC1ERKS_(void* arg0);
extern void _ZN23QGraphicsSceneHelpEventC1ERKS_(void* qthis, void* arg0);
  // proto:  void QGraphicsSceneHelpEvent::setScreenPos(const QPoint & pos);
extern void _ZN23QGraphicsSceneHelpEvent12setScreenPosERK6QPoint(void* qthis, void* arg0);
  // proto:  QPointF QGraphicsSceneHelpEvent::scenePos();
extern void _ZNK23QGraphicsSceneHelpEvent8scenePosEv(void* qthis);
  // proto:  QPointF QGraphicsSceneHoverEvent::scenePos();
extern void _ZNK24QGraphicsSceneHoverEvent8scenePosEv(void* qthis);
  // proto:  void QGraphicsSceneHoverEvent::setLastPos(const QPointF & pos);
extern void _ZN24QGraphicsSceneHoverEvent10setLastPosERK7QPointF(void* qthis, void* arg0);
  // proto:  QPointF QGraphicsSceneHoverEvent::lastPos();
extern void _ZNK24QGraphicsSceneHoverEvent7lastPosEv(void* qthis);
  // proto:  void QGraphicsSceneHoverEvent::QGraphicsSceneHoverEvent(const QGraphicsSceneHoverEvent & );
extern void* dector_ZN24QGraphicsSceneHoverEventC1ERKS_(void* arg0);
extern void _ZN24QGraphicsSceneHoverEventC1ERKS_(void* qthis, void* arg0);
  // proto:  QPointF QGraphicsSceneHoverEvent::lastScenePos();
extern void _ZNK24QGraphicsSceneHoverEvent12lastScenePosEv(void* qthis);
  // proto:  void QGraphicsSceneHoverEvent::setLastScreenPos(const QPoint & pos);
extern void _ZN24QGraphicsSceneHoverEvent16setLastScreenPosERK6QPoint(void* qthis, void* arg0);
  // proto:  void QGraphicsSceneHoverEvent::setScenePos(const QPointF & pos);
extern void _ZN24QGraphicsSceneHoverEvent11setScenePosERK7QPointF(void* qthis, void* arg0);
  // proto:  void QGraphicsSceneHoverEvent::setPos(const QPointF & pos);
extern void _ZN24QGraphicsSceneHoverEvent6setPosERK7QPointF(void* qthis, void* arg0);
  // proto:  QPoint QGraphicsSceneHoverEvent::screenPos();
extern void _ZNK24QGraphicsSceneHoverEvent9screenPosEv(void* qthis);
  // proto:  QPoint QGraphicsSceneHoverEvent::lastScreenPos();
extern void _ZNK24QGraphicsSceneHoverEvent13lastScreenPosEv(void* qthis);
  // proto:  void QGraphicsSceneHoverEvent::setLastScenePos(const QPointF & pos);
extern void _ZN24QGraphicsSceneHoverEvent15setLastScenePosERK7QPointF(void* qthis, void* arg0);
  // proto:  QPointF QGraphicsSceneHoverEvent::pos();
extern void _ZNK24QGraphicsSceneHoverEvent3posEv(void* qthis);
  // proto:  void QGraphicsSceneHoverEvent::setScreenPos(const QPoint & pos);
extern void _ZN24QGraphicsSceneHoverEvent12setScreenPosERK6QPoint(void* qthis, void* arg0);
  // proto:  void QGraphicsSceneHoverEvent::~QGraphicsSceneHoverEvent();
extern void _ZN24QGraphicsSceneHoverEventD0Ev(void* qthis);
  // proto:  QPointF QGraphicsSceneWheelEvent::pos();
extern void _ZNK24QGraphicsSceneWheelEvent3posEv(void* qthis);
  // proto:  void QGraphicsSceneWheelEvent::~QGraphicsSceneWheelEvent();
extern void _ZN24QGraphicsSceneWheelEventD0Ev(void* qthis);
  // proto:  void QGraphicsSceneWheelEvent::setDelta(int delta);
extern void _ZN24QGraphicsSceneWheelEvent8setDeltaEi(void* qthis, int arg0);
  // proto:  void QGraphicsSceneWheelEvent::setScenePos(const QPointF & pos);
extern void _ZN24QGraphicsSceneWheelEvent11setScenePosERK7QPointF(void* qthis, void* arg0);
  // proto:  void QGraphicsSceneWheelEvent::QGraphicsSceneWheelEvent(const QGraphicsSceneWheelEvent & );
extern void* dector_ZN24QGraphicsSceneWheelEventC1ERKS_(void* arg0);
extern void _ZN24QGraphicsSceneWheelEventC1ERKS_(void* qthis, void* arg0);
  // proto:  void QGraphicsSceneWheelEvent::setPos(const QPointF & pos);
extern void _ZN24QGraphicsSceneWheelEvent6setPosERK7QPointF(void* qthis, void* arg0);
  // proto:  void QGraphicsSceneWheelEvent::setScreenPos(const QPoint & pos);
extern void _ZN24QGraphicsSceneWheelEvent12setScreenPosERK6QPoint(void* qthis, void* arg0);
  // proto:  int QGraphicsSceneWheelEvent::delta();
extern void _ZNK24QGraphicsSceneWheelEvent5deltaEv(void* qthis);
  // proto:  QPointF QGraphicsSceneWheelEvent::scenePos();
extern void _ZNK24QGraphicsSceneWheelEvent8scenePosEv(void* qthis);
  // proto:  QPoint QGraphicsSceneWheelEvent::screenPos();
extern void _ZNK24QGraphicsSceneWheelEvent9screenPosEv(void* qthis);
  // proto:  QWidget * QGraphicsSceneDragDropEvent::source();
extern void _ZNK27QGraphicsSceneDragDropEvent6sourceEv(void* qthis);
  // proto:  QPointF QGraphicsSceneDragDropEvent::scenePos();
extern void _ZNK27QGraphicsSceneDragDropEvent8scenePosEv(void* qthis);
  // proto:  void QGraphicsSceneDragDropEvent::setPos(const QPointF & pos);
extern void _ZN27QGraphicsSceneDragDropEvent6setPosERK7QPointF(void* qthis, void* arg0);
  // proto:  void QGraphicsSceneDragDropEvent::QGraphicsSceneDragDropEvent(const QGraphicsSceneDragDropEvent & );
extern void* dector_ZN27QGraphicsSceneDragDropEventC1ERKS_(void* arg0);
extern void _ZN27QGraphicsSceneDragDropEventC1ERKS_(void* qthis, void* arg0);
  // proto:  void QGraphicsSceneDragDropEvent::setScreenPos(const QPoint & pos);
extern void _ZN27QGraphicsSceneDragDropEvent12setScreenPosERK6QPoint(void* qthis, void* arg0);
  // proto:  QPointF QGraphicsSceneDragDropEvent::pos();
extern void _ZNK27QGraphicsSceneDragDropEvent3posEv(void* qthis);
  // proto:  QPoint QGraphicsSceneDragDropEvent::screenPos();
extern void _ZNK27QGraphicsSceneDragDropEvent9screenPosEv(void* qthis);
  // proto:  const QMimeData * QGraphicsSceneDragDropEvent::mimeData();
extern void _ZNK27QGraphicsSceneDragDropEvent8mimeDataEv(void* qthis);
  // proto:  void QGraphicsSceneDragDropEvent::~QGraphicsSceneDragDropEvent();
extern void _ZN27QGraphicsSceneDragDropEventD0Ev(void* qthis);
  // proto:  void QGraphicsSceneDragDropEvent::setMimeData(const QMimeData * data);
extern void _ZN27QGraphicsSceneDragDropEvent11setMimeDataEPK9QMimeData(void* qthis, void* arg0);
  // proto:  void QGraphicsSceneDragDropEvent::setSource(QWidget * source);
extern void _ZN27QGraphicsSceneDragDropEvent9setSourceEP7QWidget(void* qthis, void* arg0);
  // proto:  void QGraphicsSceneDragDropEvent::setScenePos(const QPointF & pos);
extern void _ZN27QGraphicsSceneDragDropEvent11setScenePosERK7QPointF(void* qthis, void* arg0);
  // proto:  void QGraphicsSceneDragDropEvent::acceptProposedAction();
extern void _ZN27QGraphicsSceneDragDropEvent20acceptProposedActionEv(void* qthis);
  // proto:  void QGraphicsSceneEvent::QGraphicsSceneEvent(const QGraphicsSceneEvent & );
extern void* dector_ZN19QGraphicsSceneEventC1ERKS_(void* arg0);
extern void _ZN19QGraphicsSceneEventC1ERKS_(void* qthis, void* arg0);
  // proto:  QWidget * QGraphicsSceneEvent::widget();
extern void _ZNK19QGraphicsSceneEvent6widgetEv(void* qthis);
  // proto:  void QGraphicsSceneEvent::setWidget(QWidget * widget);
extern void _ZN19QGraphicsSceneEvent9setWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  void QGraphicsSceneEvent::~QGraphicsSceneEvent();
extern void _ZN19QGraphicsSceneEventD0Ev(void* qthis);
  // proto:  QSizeF QGraphicsSceneResizeEvent::newSize();
extern void _ZNK25QGraphicsSceneResizeEvent7newSizeEv(void* qthis);
  // proto:  QSizeF QGraphicsSceneResizeEvent::oldSize();
extern void _ZNK25QGraphicsSceneResizeEvent7oldSizeEv(void* qthis);
  // proto:  void QGraphicsSceneResizeEvent::~QGraphicsSceneResizeEvent();
extern void _ZN25QGraphicsSceneResizeEventD0Ev(void* qthis);
  // proto:  void QGraphicsSceneResizeEvent::setNewSize(const QSizeF & size);
extern void _ZN25QGraphicsSceneResizeEvent10setNewSizeERK6QSizeF(void* qthis, void* arg0);
  // proto:  void QGraphicsSceneResizeEvent::QGraphicsSceneResizeEvent();
extern void* dector_ZN25QGraphicsSceneResizeEventC1Ev();
extern void _ZN25QGraphicsSceneResizeEventC1Ev(void* qthis);
  // proto:  void QGraphicsSceneResizeEvent::setOldSize(const QSizeF & size);
extern void _ZN25QGraphicsSceneResizeEvent10setOldSizeERK6QSizeF(void* qthis, void* arg0);
  // proto:  void QGraphicsSceneResizeEvent::QGraphicsSceneResizeEvent(const QGraphicsSceneResizeEvent & );
extern void* dector_ZN25QGraphicsSceneResizeEventC1ERKS_(void* arg0);
extern void _ZN25QGraphicsSceneResizeEventC1ERKS_(void* qthis, void* arg0);
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

  // proto:  QPointF QGraphicsSceneMoveEvent::newPos();
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

  // proto:  QPointF QGraphicsSceneMoveEvent::oldPos();
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

  // proto:  void QGraphicsSceneMoveEvent::~QGraphicsSceneMoveEvent();
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

  // proto:  void QGraphicsSceneMoveEvent::QGraphicsSceneMoveEvent(const QGraphicsSceneMoveEvent & );
func NewQGraphicsSceneMoveEvent(args ...interface{}) QGraphicsSceneMoveEvent {
  return QGraphicsSceneMoveEvent{}
}

  // proto:  void QGraphicsSceneMoveEvent::setNewPos(const QPointF & pos);
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
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneMoveEvent", "setNewPos", args)
  }

}

  // proto:  void QGraphicsSceneMoveEvent::setOldPos(const QPointF & pos);
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
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneMoveEvent", "setOldPos", args)
  }

}

  // proto:  void QGraphicsSceneContextMenuEvent::QGraphicsSceneContextMenuEvent(const QGraphicsSceneContextMenuEvent & );
func NewQGraphicsSceneContextMenuEvent(args ...interface{}) QGraphicsSceneContextMenuEvent {
  return QGraphicsSceneContextMenuEvent{}
}

  // proto:  QPointF QGraphicsSceneContextMenuEvent::scenePos();
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

  // proto:  void QGraphicsSceneContextMenuEvent::~QGraphicsSceneContextMenuEvent();
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

  // proto:  QPointF QGraphicsSceneContextMenuEvent::pos();
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

  // proto:  void QGraphicsSceneContextMenuEvent::setScreenPos(const QPoint & pos);
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
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneContextMenuEvent", "setScreenPos", args)
  }

}

  // proto:  void QGraphicsSceneContextMenuEvent::setPos(const QPointF & pos);
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
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneContextMenuEvent", "setPos", args)
  }

}

  // proto:  QPoint QGraphicsSceneContextMenuEvent::screenPos();
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

  // proto:  void QGraphicsSceneContextMenuEvent::setScenePos(const QPointF & pos);
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
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneContextMenuEvent", "setScenePos", args)
  }

}

  // proto:  QPoint QGraphicsSceneMouseEvent::screenPos();
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

  // proto:  void QGraphicsSceneMouseEvent::QGraphicsSceneMouseEvent(const QGraphicsSceneMouseEvent & );
func NewQGraphicsSceneMouseEvent(args ...interface{}) QGraphicsSceneMouseEvent {
  return QGraphicsSceneMouseEvent{}
}

  // proto:  QPointF QGraphicsSceneMouseEvent::lastScenePos();
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

  // proto:  void QGraphicsSceneMouseEvent::~QGraphicsSceneMouseEvent();
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

  // proto:  QPointF QGraphicsSceneMouseEvent::pos();
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

  // proto:  void QGraphicsSceneMouseEvent::setLastPos(const QPointF & pos);
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
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "setLastPos", args)
  }

}

  // proto:  void QGraphicsSceneMouseEvent::setLastScenePos(const QPointF & pos);
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
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "setLastScenePos", args)
  }

}

  // proto:  QPoint QGraphicsSceneMouseEvent::lastScreenPos();
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

  // proto:  void QGraphicsSceneMouseEvent::setScreenPos(const QPoint & pos);
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
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "setScreenPos", args)
  }

}

  // proto:  void QGraphicsSceneMouseEvent::setLastScreenPos(const QPoint & pos);
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
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "setLastScreenPos", args)
  }

}

  // proto:  void QGraphicsSceneMouseEvent::setScenePos(const QPointF & pos);
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
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "setScenePos", args)
  }

}

  // proto:  QPointF QGraphicsSceneMouseEvent::lastPos();
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

  // proto:  QPointF QGraphicsSceneMouseEvent::scenePos();
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

  // proto:  void QGraphicsSceneMouseEvent::setPos(const QPointF & pos);
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
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "setPos", args)
  }

}

  // proto:  void QGraphicsSceneHelpEvent::setScenePos(const QPointF & pos);
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
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneHelpEvent", "setScenePos", args)
  }

}

  // proto:  QPoint QGraphicsSceneHelpEvent::screenPos();
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

  // proto:  void QGraphicsSceneHelpEvent::~QGraphicsSceneHelpEvent();
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

  // proto:  void QGraphicsSceneHelpEvent::QGraphicsSceneHelpEvent(const QGraphicsSceneHelpEvent & );
func NewQGraphicsSceneHelpEvent(args ...interface{}) QGraphicsSceneHelpEvent {
  return QGraphicsSceneHelpEvent{}
}

  // proto:  void QGraphicsSceneHelpEvent::setScreenPos(const QPoint & pos);
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
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneHelpEvent", "setScreenPos", args)
  }

}

  // proto:  QPointF QGraphicsSceneHelpEvent::scenePos();
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

  // proto:  QPointF QGraphicsSceneHoverEvent::scenePos();
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

  // proto:  void QGraphicsSceneHoverEvent::setLastPos(const QPointF & pos);
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
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "setLastPos", args)
  }

}

  // proto:  QPointF QGraphicsSceneHoverEvent::lastPos();
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

  // proto:  void QGraphicsSceneHoverEvent::QGraphicsSceneHoverEvent(const QGraphicsSceneHoverEvent & );
func NewQGraphicsSceneHoverEvent(args ...interface{}) QGraphicsSceneHoverEvent {
  return QGraphicsSceneHoverEvent{}
}

  // proto:  QPointF QGraphicsSceneHoverEvent::lastScenePos();
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

  // proto:  void QGraphicsSceneHoverEvent::setLastScreenPos(const QPoint & pos);
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
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "setLastScreenPos", args)
  }

}

  // proto:  void QGraphicsSceneHoverEvent::setScenePos(const QPointF & pos);
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
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "setScenePos", args)
  }

}

  // proto:  void QGraphicsSceneHoverEvent::setPos(const QPointF & pos);
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
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "setPos", args)
  }

}

  // proto:  QPoint QGraphicsSceneHoverEvent::screenPos();
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

  // proto:  QPoint QGraphicsSceneHoverEvent::lastScreenPos();
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

  // proto:  void QGraphicsSceneHoverEvent::setLastScenePos(const QPointF & pos);
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
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "setLastScenePos", args)
  }

}

  // proto:  QPointF QGraphicsSceneHoverEvent::pos();
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

  // proto:  void QGraphicsSceneHoverEvent::setScreenPos(const QPoint & pos);
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
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "setScreenPos", args)
  }

}

  // proto:  void QGraphicsSceneHoverEvent::~QGraphicsSceneHoverEvent();
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

  // proto:  QPointF QGraphicsSceneWheelEvent::pos();
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

  // proto:  void QGraphicsSceneWheelEvent::~QGraphicsSceneWheelEvent();
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

  // proto:  void QGraphicsSceneWheelEvent::setDelta(int delta);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneWheelEvent", "setDelta", args)
  }

}

  // proto:  void QGraphicsSceneWheelEvent::setScenePos(const QPointF & pos);
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
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneWheelEvent", "setScenePos", args)
  }

}

  // proto:  void QGraphicsSceneWheelEvent::QGraphicsSceneWheelEvent(const QGraphicsSceneWheelEvent & );
func NewQGraphicsSceneWheelEvent(args ...interface{}) QGraphicsSceneWheelEvent {
  return QGraphicsSceneWheelEvent{}
}

  // proto:  void QGraphicsSceneWheelEvent::setPos(const QPointF & pos);
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
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneWheelEvent", "setPos", args)
  }

}

  // proto:  void QGraphicsSceneWheelEvent::setScreenPos(const QPoint & pos);
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
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneWheelEvent", "setScreenPos", args)
  }

}

  // proto:  int QGraphicsSceneWheelEvent::delta();
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

  // proto:  QPointF QGraphicsSceneWheelEvent::scenePos();
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

  // proto:  QPoint QGraphicsSceneWheelEvent::screenPos();
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

  // proto:  QWidget * QGraphicsSceneDragDropEvent::source();
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

  // proto:  QPointF QGraphicsSceneDragDropEvent::scenePos();
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

  // proto:  void QGraphicsSceneDragDropEvent::setPos(const QPointF & pos);
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
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "setPos", args)
  }

}

  // proto:  void QGraphicsSceneDragDropEvent::QGraphicsSceneDragDropEvent(const QGraphicsSceneDragDropEvent & );
func NewQGraphicsSceneDragDropEvent(args ...interface{}) QGraphicsSceneDragDropEvent {
  return QGraphicsSceneDragDropEvent{}
}

  // proto:  void QGraphicsSceneDragDropEvent::setScreenPos(const QPoint & pos);
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
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "setScreenPos", args)
  }

}

  // proto:  QPointF QGraphicsSceneDragDropEvent::pos();
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

  // proto:  QPoint QGraphicsSceneDragDropEvent::screenPos();
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

  // proto:  const QMimeData * QGraphicsSceneDragDropEvent::mimeData();
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

  // proto:  void QGraphicsSceneDragDropEvent::~QGraphicsSceneDragDropEvent();
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

  // proto:  void QGraphicsSceneDragDropEvent::setMimeData(const QMimeData * data);
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
    var arg0 = args[0].(QMimeData).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "setMimeData", args)
  }

}

  // proto:  void QGraphicsSceneDragDropEvent::setSource(QWidget * source);
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
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "setSource", args)
  }

}

  // proto:  void QGraphicsSceneDragDropEvent::setScenePos(const QPointF & pos);
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
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "setScenePos", args)
  }

}

  // proto:  void QGraphicsSceneDragDropEvent::acceptProposedAction();
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

  // proto:  void QGraphicsSceneEvent::QGraphicsSceneEvent(const QGraphicsSceneEvent & );
func NewQGraphicsSceneEvent(args ...interface{}) QGraphicsSceneEvent {
  return QGraphicsSceneEvent{}
}

  // proto:  QWidget * QGraphicsSceneEvent::widget();
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

  // proto:  void QGraphicsSceneEvent::setWidget(QWidget * widget);
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
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneEvent", "setWidget", args)
  }

}

  // proto:  void QGraphicsSceneEvent::~QGraphicsSceneEvent();
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

  // proto:  QSizeF QGraphicsSceneResizeEvent::newSize();
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

  // proto:  QSizeF QGraphicsSceneResizeEvent::oldSize();
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

  // proto:  void QGraphicsSceneResizeEvent::~QGraphicsSceneResizeEvent();
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

  // proto:  void QGraphicsSceneResizeEvent::setNewSize(const QSizeF & size);
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
    var arg0 = args[0].(QSizeF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneResizeEvent", "setNewSize", args)
  }

}

  // proto:  void QGraphicsSceneResizeEvent::QGraphicsSceneResizeEvent();
func NewQGraphicsSceneResizeEvent(args ...interface{}) QGraphicsSceneResizeEvent {
  return QGraphicsSceneResizeEvent{}
}

  // proto:  void QGraphicsSceneResizeEvent::setOldSize(const QSizeF & size);
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
    var arg0 = args[0].(QSizeF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneResizeEvent", "setOldSize", args)
  }

}

// <= body block end

