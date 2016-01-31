package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QGraphicsSceneMoveEvent::setNewPos(const QPointF & pos);
extern void C_ZN23QGraphicsSceneMoveEvent9setNewPosERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QPointF QGraphicsSceneMoveEvent::newPos();
extern void C_ZNK23QGraphicsSceneMoveEvent6newPosEv(void* qthis); // 4
  // proto:  void QGraphicsSceneMoveEvent::QGraphicsSceneMoveEvent();
extern void* C_ZN23QGraphicsSceneMoveEventC2Ev(); // 3
  // proto:  void QGraphicsSceneMoveEvent::~QGraphicsSceneMoveEvent();
extern void C_ZN23QGraphicsSceneMoveEventD2Ev(void* qthis); // 4
  // proto:  QPointF QGraphicsSceneMoveEvent::oldPos();
extern void C_ZNK23QGraphicsSceneMoveEvent6oldPosEv(void* qthis); // 4
  // proto:  void QGraphicsSceneMoveEvent::setOldPos(const QPointF & pos);
extern void C_ZN23QGraphicsSceneMoveEvent9setOldPosERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  Qt::KeyboardModifiers QGraphicsSceneContextMenuEvent::modifiers();
extern void C_ZNK30QGraphicsSceneContextMenuEvent9modifiersEv(void* qthis); // 4
  // proto:  void QGraphicsSceneContextMenuEvent::setScreenPos(const QPoint & pos);
extern void C_ZN30QGraphicsSceneContextMenuEvent12setScreenPosERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsSceneContextMenuEvent::setPos(const QPointF & pos);
extern void C_ZN30QGraphicsSceneContextMenuEvent6setPosERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QPointF QGraphicsSceneContextMenuEvent::pos();
extern void C_ZNK30QGraphicsSceneContextMenuEvent3posEv(void* qthis); // 4
  // proto:  void QGraphicsSceneContextMenuEvent::setScenePos(const QPointF & pos);
extern void C_ZN30QGraphicsSceneContextMenuEvent11setScenePosERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QGraphicsSceneContextMenuEvent::Reason QGraphicsSceneContextMenuEvent::reason();
extern void C_ZNK30QGraphicsSceneContextMenuEvent6reasonEv(void* qthis); // 4
  // proto:  QPoint QGraphicsSceneContextMenuEvent::screenPos();
extern void C_ZNK30QGraphicsSceneContextMenuEvent9screenPosEv(void* qthis); // 4
  // proto:  QPointF QGraphicsSceneContextMenuEvent::scenePos();
extern void C_ZNK30QGraphicsSceneContextMenuEvent8scenePosEv(void* qthis); // 4
  // proto:  void QGraphicsSceneContextMenuEvent::~QGraphicsSceneContextMenuEvent();
extern void C_ZN30QGraphicsSceneContextMenuEventD2Ev(void* qthis); // 4
  // proto:  void QGraphicsSceneMouseEvent::setScreenPos(const QPoint & pos);
extern void C_ZN24QGraphicsSceneMouseEvent12setScreenPosERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QPointF QGraphicsSceneMouseEvent::lastPos();
extern void C_ZNK24QGraphicsSceneMouseEvent7lastPosEv(void* qthis); // 4
  // proto:  QPointF QGraphicsSceneMouseEvent::pos();
extern void C_ZNK24QGraphicsSceneMouseEvent3posEv(void* qthis); // 4
  // proto:  QPointF QGraphicsSceneMouseEvent::lastScenePos();
extern void C_ZNK24QGraphicsSceneMouseEvent12lastScenePosEv(void* qthis); // 4
  // proto:  Qt::KeyboardModifiers QGraphicsSceneMouseEvent::modifiers();
extern void C_ZNK24QGraphicsSceneMouseEvent9modifiersEv(void* qthis); // 4
  // proto:  void QGraphicsSceneMouseEvent::setLastPos(const QPointF & pos);
extern void C_ZN24QGraphicsSceneMouseEvent10setLastPosERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsSceneMouseEvent::~QGraphicsSceneMouseEvent();
extern void C_ZN24QGraphicsSceneMouseEventD2Ev(void* qthis); // 4
  // proto:  Qt::MouseButtons QGraphicsSceneMouseEvent::buttons();
extern void C_ZNK24QGraphicsSceneMouseEvent7buttonsEv(void* qthis); // 4
  // proto:  Qt::MouseEventSource QGraphicsSceneMouseEvent::source();
extern void C_ZNK24QGraphicsSceneMouseEvent6sourceEv(void* qthis); // 4
  // proto:  QPointF QGraphicsSceneMouseEvent::scenePos();
extern void C_ZNK24QGraphicsSceneMouseEvent8scenePosEv(void* qthis); // 4
  // proto:  void QGraphicsSceneMouseEvent::setScenePos(const QPointF & pos);
extern void C_ZN24QGraphicsSceneMouseEvent11setScenePosERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsSceneMouseEvent::setLastScreenPos(const QPoint & pos);
extern void C_ZN24QGraphicsSceneMouseEvent16setLastScreenPosERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QPoint QGraphicsSceneMouseEvent::screenPos();
extern void C_ZNK24QGraphicsSceneMouseEvent9screenPosEv(void* qthis); // 4
  // proto:  QPoint QGraphicsSceneMouseEvent::lastScreenPos();
extern void C_ZNK24QGraphicsSceneMouseEvent13lastScreenPosEv(void* qthis); // 4
  // proto:  void QGraphicsSceneMouseEvent::setPos(const QPointF & pos);
extern void C_ZN24QGraphicsSceneMouseEvent6setPosERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  Qt::MouseButton QGraphicsSceneMouseEvent::button();
extern void C_ZNK24QGraphicsSceneMouseEvent6buttonEv(void* qthis); // 4
  // proto:  Qt::MouseEventFlags QGraphicsSceneMouseEvent::flags();
extern void C_ZNK24QGraphicsSceneMouseEvent5flagsEv(void* qthis); // 4
  // proto:  void QGraphicsSceneMouseEvent::setLastScenePos(const QPointF & pos);
extern void C_ZN24QGraphicsSceneMouseEvent15setLastScenePosERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QPoint QGraphicsSceneHelpEvent::screenPos();
extern void C_ZNK23QGraphicsSceneHelpEvent9screenPosEv(void* qthis); // 4
  // proto:  void QGraphicsSceneHelpEvent::setScreenPos(const QPoint & pos);
extern void C_ZN23QGraphicsSceneHelpEvent12setScreenPosERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsSceneHelpEvent::~QGraphicsSceneHelpEvent();
extern void C_ZN23QGraphicsSceneHelpEventD2Ev(void* qthis); // 4
  // proto:  void QGraphicsSceneHelpEvent::setScenePos(const QPointF & pos);
extern void C_ZN23QGraphicsSceneHelpEvent11setScenePosERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QPointF QGraphicsSceneHelpEvent::scenePos();
extern void C_ZNK23QGraphicsSceneHelpEvent8scenePosEv(void* qthis); // 4
  // proto:  Qt::KeyboardModifiers QGraphicsSceneHoverEvent::modifiers();
extern void C_ZNK24QGraphicsSceneHoverEvent9modifiersEv(void* qthis); // 4
  // proto:  void QGraphicsSceneHoverEvent::setLastPos(const QPointF & pos);
extern void C_ZN24QGraphicsSceneHoverEvent10setLastPosERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsSceneHoverEvent::setScreenPos(const QPoint & pos);
extern void C_ZN24QGraphicsSceneHoverEvent12setScreenPosERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QPointF QGraphicsSceneHoverEvent::lastPos();
extern void C_ZNK24QGraphicsSceneHoverEvent7lastPosEv(void* qthis); // 4
  // proto:  QPoint QGraphicsSceneHoverEvent::lastScreenPos();
extern void C_ZNK24QGraphicsSceneHoverEvent13lastScreenPosEv(void* qthis); // 4
  // proto:  QPointF QGraphicsSceneHoverEvent::pos();
extern void C_ZNK24QGraphicsSceneHoverEvent3posEv(void* qthis); // 4
  // proto:  void QGraphicsSceneHoverEvent::setScenePos(const QPointF & pos);
extern void C_ZN24QGraphicsSceneHoverEvent11setScenePosERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsSceneHoverEvent::setPos(const QPointF & pos);
extern void C_ZN24QGraphicsSceneHoverEvent6setPosERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QPoint QGraphicsSceneHoverEvent::screenPos();
extern void C_ZNK24QGraphicsSceneHoverEvent9screenPosEv(void* qthis); // 4
  // proto:  void QGraphicsSceneHoverEvent::setLastScenePos(const QPointF & pos);
extern void C_ZN24QGraphicsSceneHoverEvent15setLastScenePosERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QPointF QGraphicsSceneHoverEvent::lastScenePos();
extern void C_ZNK24QGraphicsSceneHoverEvent12lastScenePosEv(void* qthis); // 4
  // proto:  void QGraphicsSceneHoverEvent::~QGraphicsSceneHoverEvent();
extern void C_ZN24QGraphicsSceneHoverEventD2Ev(void* qthis); // 4
  // proto:  void QGraphicsSceneHoverEvent::setLastScreenPos(const QPoint & pos);
extern void C_ZN24QGraphicsSceneHoverEvent16setLastScreenPosERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QPointF QGraphicsSceneHoverEvent::scenePos();
extern void C_ZNK24QGraphicsSceneHoverEvent8scenePosEv(void* qthis); // 4
  // proto:  Qt::KeyboardModifiers QGraphicsSceneWheelEvent::modifiers();
extern void C_ZNK24QGraphicsSceneWheelEvent9modifiersEv(void* qthis); // 4
  // proto:  void QGraphicsSceneWheelEvent::setDelta(int delta);
extern void C_ZN24QGraphicsSceneWheelEvent8setDeltaEi(void* qthis, int32_t arg0); // 4
  // proto:  Qt::Orientation QGraphicsSceneWheelEvent::orientation();
extern void C_ZNK24QGraphicsSceneWheelEvent11orientationEv(void* qthis); // 4
  // proto:  void QGraphicsSceneWheelEvent::setScreenPos(const QPoint & pos);
extern void C_ZN24QGraphicsSceneWheelEvent12setScreenPosERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsSceneWheelEvent::setPos(const QPointF & pos);
extern void C_ZN24QGraphicsSceneWheelEvent6setPosERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QPointF QGraphicsSceneWheelEvent::pos();
extern void C_ZNK24QGraphicsSceneWheelEvent3posEv(void* qthis); // 4
  // proto:  void QGraphicsSceneWheelEvent::setScenePos(const QPointF & pos);
extern void C_ZN24QGraphicsSceneWheelEvent11setScenePosERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  Qt::MouseButtons QGraphicsSceneWheelEvent::buttons();
extern void C_ZNK24QGraphicsSceneWheelEvent7buttonsEv(void* qthis); // 4
  // proto:  QPoint QGraphicsSceneWheelEvent::screenPos();
extern void C_ZNK24QGraphicsSceneWheelEvent9screenPosEv(void* qthis); // 4
  // proto:  int QGraphicsSceneWheelEvent::delta();
extern void C_ZNK24QGraphicsSceneWheelEvent5deltaEv(void* qthis); // 4
  // proto:  void QGraphicsSceneWheelEvent::~QGraphicsSceneWheelEvent();
extern void C_ZN24QGraphicsSceneWheelEventD2Ev(void* qthis); // 4
  // proto:  QPointF QGraphicsSceneWheelEvent::scenePos();
extern void C_ZNK24QGraphicsSceneWheelEvent8scenePosEv(void* qthis); // 4
  // proto:  void QGraphicsSceneDragDropEvent::setMimeData(const QMimeData * data);
extern void C_ZN27QGraphicsSceneDragDropEvent11setMimeDataEPK9QMimeData(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsSceneDragDropEvent::setScreenPos(const QPoint & pos);
extern void C_ZN27QGraphicsSceneDragDropEvent12setScreenPosERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QPointF QGraphicsSceneDragDropEvent::pos();
extern void C_ZNK27QGraphicsSceneDragDropEvent3posEv(void* qthis); // 4
  // proto:  void QGraphicsSceneDragDropEvent::setSource(QWidget * source);
extern void C_ZN27QGraphicsSceneDragDropEvent9setSourceEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsSceneDragDropEvent::acceptProposedAction();
extern void C_ZN27QGraphicsSceneDragDropEvent20acceptProposedActionEv(void* qthis); // 4
  // proto:  Qt::MouseButtons QGraphicsSceneDragDropEvent::buttons();
extern void C_ZNK27QGraphicsSceneDragDropEvent7buttonsEv(void* qthis); // 4
  // proto:  QWidget * QGraphicsSceneDragDropEvent::source();
extern void C_ZNK27QGraphicsSceneDragDropEvent6sourceEv(void* qthis); // 4
  // proto:  QPointF QGraphicsSceneDragDropEvent::scenePos();
extern void C_ZNK27QGraphicsSceneDragDropEvent8scenePosEv(void* qthis); // 4
  // proto:  Qt::DropAction QGraphicsSceneDragDropEvent::proposedAction();
extern void C_ZNK27QGraphicsSceneDragDropEvent14proposedActionEv(void* qthis); // 4
  // proto:  void QGraphicsSceneDragDropEvent::setScenePos(const QPointF & pos);
extern void C_ZN27QGraphicsSceneDragDropEvent11setScenePosERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  Qt::DropActions QGraphicsSceneDragDropEvent::possibleActions();
extern void C_ZNK27QGraphicsSceneDragDropEvent15possibleActionsEv(void* qthis); // 4
  // proto:  Qt::DropAction QGraphicsSceneDragDropEvent::dropAction();
extern void C_ZNK27QGraphicsSceneDragDropEvent10dropActionEv(void* qthis); // 4
  // proto:  void QGraphicsSceneDragDropEvent::~QGraphicsSceneDragDropEvent();
extern void C_ZN27QGraphicsSceneDragDropEventD2Ev(void* qthis); // 4
  // proto:  QPoint QGraphicsSceneDragDropEvent::screenPos();
extern void C_ZNK27QGraphicsSceneDragDropEvent9screenPosEv(void* qthis); // 4
  // proto:  const QMimeData * QGraphicsSceneDragDropEvent::mimeData();
extern void C_ZNK27QGraphicsSceneDragDropEvent8mimeDataEv(void* qthis); // 4
  // proto:  Qt::KeyboardModifiers QGraphicsSceneDragDropEvent::modifiers();
extern void C_ZNK27QGraphicsSceneDragDropEvent9modifiersEv(void* qthis); // 4
  // proto:  void QGraphicsSceneDragDropEvent::setPos(const QPointF & pos);
extern void C_ZN27QGraphicsSceneDragDropEvent6setPosERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsSceneEvent::~QGraphicsSceneEvent();
extern void C_ZN19QGraphicsSceneEventD2Ev(void* qthis); // 4
  // proto:  QWidget * QGraphicsSceneEvent::widget();
extern void C_ZNK19QGraphicsSceneEvent6widgetEv(void* qthis); // 4
  // proto:  void QGraphicsSceneEvent::setWidget(QWidget * widget);
extern void C_ZN19QGraphicsSceneEvent9setWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  QSizeF QGraphicsSceneResizeEvent::newSize();
extern void C_ZNK25QGraphicsSceneResizeEvent7newSizeEv(void* qthis); // 4
  // proto:  QSizeF QGraphicsSceneResizeEvent::oldSize();
extern void C_ZNK25QGraphicsSceneResizeEvent7oldSizeEv(void* qthis); // 4
  // proto:  void QGraphicsSceneResizeEvent::setOldSize(const QSizeF & size);
extern void C_ZN25QGraphicsSceneResizeEvent10setOldSizeERK6QSizeF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsSceneResizeEvent::setNewSize(const QSizeF & size);
extern void C_ZN25QGraphicsSceneResizeEvent10setNewSizeERK6QSizeF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsSceneResizeEvent::~QGraphicsSceneResizeEvent();
extern void C_ZN25QGraphicsSceneResizeEventD2Ev(void* qthis); // 4
  // proto:  void QGraphicsSceneResizeEvent::QGraphicsSceneResizeEvent();
extern void* C_ZN25QGraphicsSceneResizeEventC2Ev(); // 3
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGraphicsSceneContextMenuEvent)=1
type QGraphicsSceneContextMenuEvent struct {
  /*qbase*/ QGraphicsSceneEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGraphicsSceneMouseEvent)=1
type QGraphicsSceneMouseEvent struct {
  /*qbase*/ QGraphicsSceneEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGraphicsSceneHelpEvent)=1
type QGraphicsSceneHelpEvent struct {
  /*qbase*/ QGraphicsSceneEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGraphicsSceneHoverEvent)=1
type QGraphicsSceneHoverEvent struct {
  /*qbase*/ QGraphicsSceneEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGraphicsSceneWheelEvent)=1
type QGraphicsSceneWheelEvent struct {
  /*qbase*/ QGraphicsSceneEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGraphicsSceneDragDropEvent)=1
type QGraphicsSceneDragDropEvent struct {
  /*qbase*/ QGraphicsSceneEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGraphicsSceneEvent)=1
type QGraphicsSceneEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGraphicsSceneResizeEvent)=1
type QGraphicsSceneResizeEvent struct {
  /*qbase*/ QGraphicsSceneEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// setNewPos(const class QPointF &)
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
    // invoke: void setNewPos(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN23QGraphicsSceneMoveEvent9setNewPosERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneMoveEvent", "setNewPos", args)
  }

}

// newPos()
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
    // invoke: QPointF newPos()
    var ret = C.C_ZNK23QGraphicsSceneMoveEvent6newPosEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneMoveEvent", "newPos", args)
  }

}

// QGraphicsSceneMoveEvent()
func NewQGraphicsSceneMoveEvent(args ...interface{}) *QGraphicsSceneMoveEvent {
  // QGraphicsSceneMoveEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QGraphicsSceneMoveEventC1Ev
    // invoke: void QGraphicsSceneMoveEvent()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN23QGraphicsSceneMoveEventC2Ev()
    return &QGraphicsSceneMoveEvent{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsSceneMoveEvent", "QGraphicsSceneMoveEvent", args)
  }

  return nil // QGraphicsSceneMoveEvent{qclsinst:qthis}
}

// ~QGraphicsSceneMoveEvent()
func (this *QGraphicsSceneMoveEvent) FreeQGraphicsSceneMoveEvent(args ...interface{}) () {
  // ~QGraphicsSceneMoveEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QGraphicsSceneMoveEventD0Ev
    // invoke: void ~QGraphicsSceneMoveEvent()
    C.C_ZN23QGraphicsSceneMoveEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSceneMoveEvent", "~QGraphicsSceneMoveEvent", args)
  }

}

// oldPos()
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
    // invoke: QPointF oldPos()
    var ret = C.C_ZNK23QGraphicsSceneMoveEvent6oldPosEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneMoveEvent", "oldPos", args)
  }

}

// setOldPos(const class QPointF &)
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
    // invoke: void setOldPos(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN23QGraphicsSceneMoveEvent9setOldPosERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneMoveEvent", "setOldPos", args)
  }

}

// modifiers()
func (this *QGraphicsSceneContextMenuEvent) modifiers(args ...interface{}) () {
  // modifiers()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK30QGraphicsSceneContextMenuEvent9modifiersEv
    // invoke: Qt::KeyboardModifiers modifiers()
    C.C_ZNK30QGraphicsSceneContextMenuEvent9modifiersEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSceneContextMenuEvent", "modifiers", args)
  }

}

// setScreenPos(const class QPoint &)
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
    // invoke: void setScreenPos(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN30QGraphicsSceneContextMenuEvent12setScreenPosERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneContextMenuEvent", "setScreenPos", args)
  }

}

// setPos(const class QPointF &)
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
    // invoke: void setPos(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN30QGraphicsSceneContextMenuEvent6setPosERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneContextMenuEvent", "setPos", args)
  }

}

// pos()
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
    // invoke: QPointF pos()
    var ret = C.C_ZNK30QGraphicsSceneContextMenuEvent3posEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneContextMenuEvent", "pos", args)
  }

}

// setScenePos(const class QPointF &)
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
    // invoke: void setScenePos(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN30QGraphicsSceneContextMenuEvent11setScenePosERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneContextMenuEvent", "setScenePos", args)
  }

}

// reason()
func (this *QGraphicsSceneContextMenuEvent) reason(args ...interface{}) () {
  // reason()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK30QGraphicsSceneContextMenuEvent6reasonEv
    // invoke: QGraphicsSceneContextMenuEvent::Reason reason()
    C.C_ZNK30QGraphicsSceneContextMenuEvent6reasonEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSceneContextMenuEvent", "reason", args)
  }

}

// screenPos()
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
    // invoke: QPoint screenPos()
    var ret = C.C_ZNK30QGraphicsSceneContextMenuEvent9screenPosEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneContextMenuEvent", "screenPos", args)
  }

}

// scenePos()
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
    // invoke: QPointF scenePos()
    var ret = C.C_ZNK30QGraphicsSceneContextMenuEvent8scenePosEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneContextMenuEvent", "scenePos", args)
  }

}

// ~QGraphicsSceneContextMenuEvent()
func (this *QGraphicsSceneContextMenuEvent) FreeQGraphicsSceneContextMenuEvent(args ...interface{}) () {
  // ~QGraphicsSceneContextMenuEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN30QGraphicsSceneContextMenuEventD0Ev
    // invoke: void ~QGraphicsSceneContextMenuEvent()
    C.C_ZN30QGraphicsSceneContextMenuEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSceneContextMenuEvent", "~QGraphicsSceneContextMenuEvent", args)
  }

}

// setScreenPos(const class QPoint &)
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
    // invoke: void setScreenPos(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN24QGraphicsSceneMouseEvent12setScreenPosERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "setScreenPos", args)
  }

}

// lastPos()
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
    // invoke: QPointF lastPos()
    var ret = C.C_ZNK24QGraphicsSceneMouseEvent7lastPosEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "lastPos", args)
  }

}

// pos()
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
    // invoke: QPointF pos()
    var ret = C.C_ZNK24QGraphicsSceneMouseEvent3posEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "pos", args)
  }

}

// lastScenePos()
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
    // invoke: QPointF lastScenePos()
    var ret = C.C_ZNK24QGraphicsSceneMouseEvent12lastScenePosEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "lastScenePos", args)
  }

}

// modifiers()
func (this *QGraphicsSceneMouseEvent) modifiers(args ...interface{}) () {
  // modifiers()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QGraphicsSceneMouseEvent9modifiersEv
    // invoke: Qt::KeyboardModifiers modifiers()
    C.C_ZNK24QGraphicsSceneMouseEvent9modifiersEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "modifiers", args)
  }

}

// setLastPos(const class QPointF &)
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
    // invoke: void setLastPos(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN24QGraphicsSceneMouseEvent10setLastPosERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "setLastPos", args)
  }

}

// ~QGraphicsSceneMouseEvent()
func (this *QGraphicsSceneMouseEvent) FreeQGraphicsSceneMouseEvent(args ...interface{}) () {
  // ~QGraphicsSceneMouseEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QGraphicsSceneMouseEventD0Ev
    // invoke: void ~QGraphicsSceneMouseEvent()
    C.C_ZN24QGraphicsSceneMouseEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "~QGraphicsSceneMouseEvent", args)
  }

}

// buttons()
func (this *QGraphicsSceneMouseEvent) buttons(args ...interface{}) () {
  // buttons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QGraphicsSceneMouseEvent7buttonsEv
    // invoke: Qt::MouseButtons buttons()
    C.C_ZNK24QGraphicsSceneMouseEvent7buttonsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "buttons", args)
  }

}

// source()
func (this *QGraphicsSceneMouseEvent) source(args ...interface{}) () {
  // source()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QGraphicsSceneMouseEvent6sourceEv
    // invoke: Qt::MouseEventSource source()
    C.C_ZNK24QGraphicsSceneMouseEvent6sourceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "source", args)
  }

}

// scenePos()
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
    // invoke: QPointF scenePos()
    var ret = C.C_ZNK24QGraphicsSceneMouseEvent8scenePosEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "scenePos", args)
  }

}

// setScenePos(const class QPointF &)
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
    // invoke: void setScenePos(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN24QGraphicsSceneMouseEvent11setScenePosERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "setScenePos", args)
  }

}

// setLastScreenPos(const class QPoint &)
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
    // invoke: void setLastScreenPos(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN24QGraphicsSceneMouseEvent16setLastScreenPosERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "setLastScreenPos", args)
  }

}

// screenPos()
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
    // invoke: QPoint screenPos()
    var ret = C.C_ZNK24QGraphicsSceneMouseEvent9screenPosEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "screenPos", args)
  }

}

// lastScreenPos()
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
    // invoke: QPoint lastScreenPos()
    var ret = C.C_ZNK24QGraphicsSceneMouseEvent13lastScreenPosEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "lastScreenPos", args)
  }

}

// setPos(const class QPointF &)
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
    // invoke: void setPos(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN24QGraphicsSceneMouseEvent6setPosERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "setPos", args)
  }

}

// button()
func (this *QGraphicsSceneMouseEvent) button(args ...interface{}) () {
  // button()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QGraphicsSceneMouseEvent6buttonEv
    // invoke: Qt::MouseButton button()
    C.C_ZNK24QGraphicsSceneMouseEvent6buttonEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "button", args)
  }

}

// flags()
func (this *QGraphicsSceneMouseEvent) flags(args ...interface{}) () {
  // flags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QGraphicsSceneMouseEvent5flagsEv
    // invoke: Qt::MouseEventFlags flags()
    C.C_ZNK24QGraphicsSceneMouseEvent5flagsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "flags", args)
  }

}

// setLastScenePos(const class QPointF &)
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
    // invoke: void setLastScenePos(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN24QGraphicsSceneMouseEvent15setLastScenePosERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneMouseEvent", "setLastScenePos", args)
  }

}

// screenPos()
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
    // invoke: QPoint screenPos()
    var ret = C.C_ZNK23QGraphicsSceneHelpEvent9screenPosEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneHelpEvent", "screenPos", args)
  }

}

// setScreenPos(const class QPoint &)
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
    // invoke: void setScreenPos(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN23QGraphicsSceneHelpEvent12setScreenPosERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneHelpEvent", "setScreenPos", args)
  }

}

// ~QGraphicsSceneHelpEvent()
func (this *QGraphicsSceneHelpEvent) FreeQGraphicsSceneHelpEvent(args ...interface{}) () {
  // ~QGraphicsSceneHelpEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QGraphicsSceneHelpEventD0Ev
    // invoke: void ~QGraphicsSceneHelpEvent()
    C.C_ZN23QGraphicsSceneHelpEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSceneHelpEvent", "~QGraphicsSceneHelpEvent", args)
  }

}

// setScenePos(const class QPointF &)
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
    // invoke: void setScenePos(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN23QGraphicsSceneHelpEvent11setScenePosERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneHelpEvent", "setScenePos", args)
  }

}

// scenePos()
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
    // invoke: QPointF scenePos()
    var ret = C.C_ZNK23QGraphicsSceneHelpEvent8scenePosEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneHelpEvent", "scenePos", args)
  }

}

// modifiers()
func (this *QGraphicsSceneHoverEvent) modifiers(args ...interface{}) () {
  // modifiers()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QGraphicsSceneHoverEvent9modifiersEv
    // invoke: Qt::KeyboardModifiers modifiers()
    C.C_ZNK24QGraphicsSceneHoverEvent9modifiersEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "modifiers", args)
  }

}

// setLastPos(const class QPointF &)
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
    // invoke: void setLastPos(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN24QGraphicsSceneHoverEvent10setLastPosERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "setLastPos", args)
  }

}

// setScreenPos(const class QPoint &)
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
    // invoke: void setScreenPos(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN24QGraphicsSceneHoverEvent12setScreenPosERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "setScreenPos", args)
  }

}

// lastPos()
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
    // invoke: QPointF lastPos()
    var ret = C.C_ZNK24QGraphicsSceneHoverEvent7lastPosEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "lastPos", args)
  }

}

// lastScreenPos()
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
    // invoke: QPoint lastScreenPos()
    var ret = C.C_ZNK24QGraphicsSceneHoverEvent13lastScreenPosEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "lastScreenPos", args)
  }

}

// pos()
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
    // invoke: QPointF pos()
    var ret = C.C_ZNK24QGraphicsSceneHoverEvent3posEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "pos", args)
  }

}

// setScenePos(const class QPointF &)
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
    // invoke: void setScenePos(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN24QGraphicsSceneHoverEvent11setScenePosERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "setScenePos", args)
  }

}

// setPos(const class QPointF &)
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
    // invoke: void setPos(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN24QGraphicsSceneHoverEvent6setPosERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "setPos", args)
  }

}

// screenPos()
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
    // invoke: QPoint screenPos()
    var ret = C.C_ZNK24QGraphicsSceneHoverEvent9screenPosEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "screenPos", args)
  }

}

// setLastScenePos(const class QPointF &)
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
    // invoke: void setLastScenePos(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN24QGraphicsSceneHoverEvent15setLastScenePosERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "setLastScenePos", args)
  }

}

// lastScenePos()
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
    // invoke: QPointF lastScenePos()
    var ret = C.C_ZNK24QGraphicsSceneHoverEvent12lastScenePosEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "lastScenePos", args)
  }

}

// ~QGraphicsSceneHoverEvent()
func (this *QGraphicsSceneHoverEvent) FreeQGraphicsSceneHoverEvent(args ...interface{}) () {
  // ~QGraphicsSceneHoverEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QGraphicsSceneHoverEventD0Ev
    // invoke: void ~QGraphicsSceneHoverEvent()
    C.C_ZN24QGraphicsSceneHoverEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "~QGraphicsSceneHoverEvent", args)
  }

}

// setLastScreenPos(const class QPoint &)
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
    // invoke: void setLastScreenPos(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN24QGraphicsSceneHoverEvent16setLastScreenPosERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "setLastScreenPos", args)
  }

}

// scenePos()
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
    // invoke: QPointF scenePos()
    var ret = C.C_ZNK24QGraphicsSceneHoverEvent8scenePosEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneHoverEvent", "scenePos", args)
  }

}

// modifiers()
func (this *QGraphicsSceneWheelEvent) modifiers(args ...interface{}) () {
  // modifiers()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QGraphicsSceneWheelEvent9modifiersEv
    // invoke: Qt::KeyboardModifiers modifiers()
    C.C_ZNK24QGraphicsSceneWheelEvent9modifiersEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSceneWheelEvent", "modifiers", args)
  }

}

// setDelta(int)
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
    // invoke: void setDelta(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN24QGraphicsSceneWheelEvent8setDeltaEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneWheelEvent", "setDelta", args)
  }

}

// orientation()
func (this *QGraphicsSceneWheelEvent) orientation(args ...interface{}) () {
  // orientation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QGraphicsSceneWheelEvent11orientationEv
    // invoke: Qt::Orientation orientation()
    C.C_ZNK24QGraphicsSceneWheelEvent11orientationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSceneWheelEvent", "orientation", args)
  }

}

// setScreenPos(const class QPoint &)
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
    // invoke: void setScreenPos(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN24QGraphicsSceneWheelEvent12setScreenPosERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneWheelEvent", "setScreenPos", args)
  }

}

// setPos(const class QPointF &)
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
    // invoke: void setPos(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN24QGraphicsSceneWheelEvent6setPosERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneWheelEvent", "setPos", args)
  }

}

// pos()
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
    // invoke: QPointF pos()
    var ret = C.C_ZNK24QGraphicsSceneWheelEvent3posEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneWheelEvent", "pos", args)
  }

}

// setScenePos(const class QPointF &)
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
    // invoke: void setScenePos(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN24QGraphicsSceneWheelEvent11setScenePosERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneWheelEvent", "setScenePos", args)
  }

}

// buttons()
func (this *QGraphicsSceneWheelEvent) buttons(args ...interface{}) () {
  // buttons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QGraphicsSceneWheelEvent7buttonsEv
    // invoke: Qt::MouseButtons buttons()
    C.C_ZNK24QGraphicsSceneWheelEvent7buttonsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSceneWheelEvent", "buttons", args)
  }

}

// screenPos()
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
    // invoke: QPoint screenPos()
    var ret = C.C_ZNK24QGraphicsSceneWheelEvent9screenPosEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneWheelEvent", "screenPos", args)
  }

}

// delta()
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
    // invoke: int delta()
    var ret = C.C_ZNK24QGraphicsSceneWheelEvent5deltaEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneWheelEvent", "delta", args)
  }

}

// ~QGraphicsSceneWheelEvent()
func (this *QGraphicsSceneWheelEvent) FreeQGraphicsSceneWheelEvent(args ...interface{}) () {
  // ~QGraphicsSceneWheelEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QGraphicsSceneWheelEventD0Ev
    // invoke: void ~QGraphicsSceneWheelEvent()
    C.C_ZN24QGraphicsSceneWheelEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSceneWheelEvent", "~QGraphicsSceneWheelEvent", args)
  }

}

// scenePos()
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
    // invoke: QPointF scenePos()
    var ret = C.C_ZNK24QGraphicsSceneWheelEvent8scenePosEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneWheelEvent", "scenePos", args)
  }

}

// setMimeData(const class QMimeData *)
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
    // invoke: void setMimeData(const class QMimeData *)
    var arg0 = args[0].(QMimeData).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN27QGraphicsSceneDragDropEvent11setMimeDataEPK9QMimeData(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "setMimeData", args)
  }

}

// setScreenPos(const class QPoint &)
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
    // invoke: void setScreenPos(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN27QGraphicsSceneDragDropEvent12setScreenPosERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "setScreenPos", args)
  }

}

// pos()
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
    // invoke: QPointF pos()
    var ret = C.C_ZNK27QGraphicsSceneDragDropEvent3posEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "pos", args)
  }

}

// setSource(class QWidget *)
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
    // invoke: void setSource(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN27QGraphicsSceneDragDropEvent9setSourceEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "setSource", args)
  }

}

// acceptProposedAction()
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
    // invoke: void acceptProposedAction()
    C.C_ZN27QGraphicsSceneDragDropEvent20acceptProposedActionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "acceptProposedAction", args)
  }

}

// buttons()
func (this *QGraphicsSceneDragDropEvent) buttons(args ...interface{}) () {
  // buttons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QGraphicsSceneDragDropEvent7buttonsEv
    // invoke: Qt::MouseButtons buttons()
    C.C_ZNK27QGraphicsSceneDragDropEvent7buttonsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "buttons", args)
  }

}

// source()
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
    // invoke: QWidget * source()
    var ret = C.C_ZNK27QGraphicsSceneDragDropEvent6sourceEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "source", args)
  }

}

// scenePos()
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
    // invoke: QPointF scenePos()
    var ret = C.C_ZNK27QGraphicsSceneDragDropEvent8scenePosEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "scenePos", args)
  }

}

// proposedAction()
func (this *QGraphicsSceneDragDropEvent) proposedAction(args ...interface{}) () {
  // proposedAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QGraphicsSceneDragDropEvent14proposedActionEv
    // invoke: Qt::DropAction proposedAction()
    C.C_ZNK27QGraphicsSceneDragDropEvent14proposedActionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "proposedAction", args)
  }

}

// setScenePos(const class QPointF &)
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
    // invoke: void setScenePos(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN27QGraphicsSceneDragDropEvent11setScenePosERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "setScenePos", args)
  }

}

// possibleActions()
func (this *QGraphicsSceneDragDropEvent) possibleActions(args ...interface{}) () {
  // possibleActions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QGraphicsSceneDragDropEvent15possibleActionsEv
    // invoke: Qt::DropActions possibleActions()
    C.C_ZNK27QGraphicsSceneDragDropEvent15possibleActionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "possibleActions", args)
  }

}

// dropAction()
func (this *QGraphicsSceneDragDropEvent) dropAction(args ...interface{}) () {
  // dropAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QGraphicsSceneDragDropEvent10dropActionEv
    // invoke: Qt::DropAction dropAction()
    C.C_ZNK27QGraphicsSceneDragDropEvent10dropActionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "dropAction", args)
  }

}

// ~QGraphicsSceneDragDropEvent()
func (this *QGraphicsSceneDragDropEvent) FreeQGraphicsSceneDragDropEvent(args ...interface{}) () {
  // ~QGraphicsSceneDragDropEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QGraphicsSceneDragDropEventD0Ev
    // invoke: void ~QGraphicsSceneDragDropEvent()
    C.C_ZN27QGraphicsSceneDragDropEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "~QGraphicsSceneDragDropEvent", args)
  }

}

// screenPos()
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
    // invoke: QPoint screenPos()
    var ret = C.C_ZNK27QGraphicsSceneDragDropEvent9screenPosEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "screenPos", args)
  }

}

// mimeData()
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
    // invoke: const QMimeData * mimeData()
    var ret = C.C_ZNK27QGraphicsSceneDragDropEvent8mimeDataEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "mimeData", args)
  }

}

// modifiers()
func (this *QGraphicsSceneDragDropEvent) modifiers(args ...interface{}) () {
  // modifiers()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QGraphicsSceneDragDropEvent9modifiersEv
    // invoke: Qt::KeyboardModifiers modifiers()
    C.C_ZNK27QGraphicsSceneDragDropEvent9modifiersEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "modifiers", args)
  }

}

// setPos(const class QPointF &)
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
    // invoke: void setPos(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN27QGraphicsSceneDragDropEvent6setPosERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneDragDropEvent", "setPos", args)
  }

}

// ~QGraphicsSceneEvent()
func (this *QGraphicsSceneEvent) FreeQGraphicsSceneEvent(args ...interface{}) () {
  // ~QGraphicsSceneEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsSceneEventD0Ev
    // invoke: void ~QGraphicsSceneEvent()
    C.C_ZN19QGraphicsSceneEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSceneEvent", "~QGraphicsSceneEvent", args)
  }

}

// widget()
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
    // invoke: QWidget * widget()
    var ret = C.C_ZNK19QGraphicsSceneEvent6widgetEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneEvent", "widget", args)
  }

}

// setWidget(class QWidget *)
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
    // invoke: void setWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsSceneEvent9setWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneEvent", "setWidget", args)
  }

}

// newSize()
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
    // invoke: QSizeF newSize()
    var ret = C.C_ZNK25QGraphicsSceneResizeEvent7newSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneResizeEvent", "newSize", args)
  }

}

// oldSize()
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
    // invoke: QSizeF oldSize()
    var ret = C.C_ZNK25QGraphicsSceneResizeEvent7oldSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsSceneResizeEvent", "oldSize", args)
  }

}

// setOldSize(const class QSizeF &)
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
    // invoke: void setOldSize(const class QSizeF &)
    var arg0 = args[0].(QSizeF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN25QGraphicsSceneResizeEvent10setOldSizeERK6QSizeF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneResizeEvent", "setOldSize", args)
  }

}

// setNewSize(const class QSizeF &)
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
    // invoke: void setNewSize(const class QSizeF &)
    var arg0 = args[0].(QSizeF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN25QGraphicsSceneResizeEvent10setNewSizeERK6QSizeF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSceneResizeEvent", "setNewSize", args)
  }

}

// ~QGraphicsSceneResizeEvent()
func (this *QGraphicsSceneResizeEvent) FreeQGraphicsSceneResizeEvent(args ...interface{}) () {
  // ~QGraphicsSceneResizeEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QGraphicsSceneResizeEventD0Ev
    // invoke: void ~QGraphicsSceneResizeEvent()
    C.C_ZN25QGraphicsSceneResizeEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSceneResizeEvent", "~QGraphicsSceneResizeEvent", args)
  }

}

// QGraphicsSceneResizeEvent()
func NewQGraphicsSceneResizeEvent(args ...interface{}) *QGraphicsSceneResizeEvent {
  // QGraphicsSceneResizeEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QGraphicsSceneResizeEventC1Ev
    // invoke: void QGraphicsSceneResizeEvent()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN25QGraphicsSceneResizeEventC2Ev()
    return &QGraphicsSceneResizeEvent{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsSceneResizeEvent", "QGraphicsSceneResizeEvent", args)
  }

  return nil // QGraphicsSceneResizeEvent{qclsinst:qthis}
}

// <= body block end

