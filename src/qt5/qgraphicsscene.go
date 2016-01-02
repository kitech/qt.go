package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtWidgets/qgraphicsscene.h
// dst-file: /src/widgets/qgraphicsscene.go
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
  // proto:  void QGraphicsScene::setForegroundBrush(const QBrush & brush);
extern void _ZN14QGraphicsScene18setForegroundBrushERK6QBrush(void* qthis, void* arg0);
  // proto:  void QGraphicsScene::setSceneRect(const QRectF & rect);
extern void _ZN14QGraphicsScene12setSceneRectERK6QRectF(void* qthis, void* arg0);
  // proto:  bool QGraphicsScene::isActive();
extern void _ZNK14QGraphicsScene8isActiveEv(void* qthis);
  // proto:  bool QGraphicsScene::hasFocus();
extern void _ZNK14QGraphicsScene8hasFocusEv(void* qthis);
  // proto:  QRectF QGraphicsScene::itemsBoundingRect();
extern void _ZNK14QGraphicsScene17itemsBoundingRectEv(void* qthis);
  // proto:  bool QGraphicsScene::sendEvent(QGraphicsItem * item, QEvent * event);
extern void _ZN14QGraphicsScene9sendEventEP13QGraphicsItemP6QEvent(void* qthis, void* arg0, void* arg1);
  // proto:  qreal QGraphicsScene::minimumRenderSize();
extern void _ZNK14QGraphicsScene17minimumRenderSizeEv(void* qthis);
  // proto:  QPainterPath QGraphicsScene::selectionArea();
extern void _ZNK14QGraphicsScene13selectionAreaEv(void* qthis);
  // proto:  void QGraphicsScene::update(const QRectF & rect);
extern void _ZN14QGraphicsScene6updateERK6QRectF(void* qthis, void* arg0);
  // proto:  QGraphicsPolygonItem * QGraphicsScene::addPolygon(const QPolygonF & polygon, const QPen & pen, const QBrush & brush);
extern void _ZN14QGraphicsScene10addPolygonERK9QPolygonFRK4QPenRK6QBrush(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  QGraphicsLineItem * QGraphicsScene::addLine(const QLineF & line, const QPen & pen);
extern void _ZN14QGraphicsScene7addLineERK6QLineFRK4QPen(void* qthis, void* arg0, void* arg1);
  // proto:  QPalette QGraphicsScene::palette();
extern void _ZNK14QGraphicsScene7paletteEv(void* qthis);
  // proto:  bool QGraphicsScene::isSortCacheEnabled();
extern void _ZNK14QGraphicsScene18isSortCacheEnabledEv(void* qthis);
  // proto:  void QGraphicsScene::QGraphicsScene(const QRectF & sceneRect, QObject * parent);
extern void* dector_ZN14QGraphicsSceneC1ERK6QRectFP7QObject(void* arg0, void* arg1);
extern void _ZN14QGraphicsSceneC1ERK6QRectFP7QObject(void* qthis, void* arg0, void* arg1);
  // proto:  void QGraphicsScene::QGraphicsScene(QObject * parent);
extern void* dector_ZN14QGraphicsSceneC1EP7QObject(void* arg0);
extern void _ZN14QGraphicsSceneC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QGraphicsScene::clearFocus();
extern void _ZN14QGraphicsScene10clearFocusEv(void* qthis);
  // proto:  const QMetaObject * QGraphicsScene::metaObject();
extern void _ZNK14QGraphicsScene10metaObjectEv(void* qthis);
  // proto:  QGraphicsSimpleTextItem * QGraphicsScene::addSimpleText(const QString & text, const QFont & font);
extern void _ZN14QGraphicsScene13addSimpleTextERK7QStringRK5QFont(void* qthis, void* arg0, void* arg1);
  // proto:  QGraphicsLineItem * QGraphicsScene::addLine(qreal x1, qreal y1, qreal x2, qreal y2, const QPen & pen);
extern void demth_ZN14QGraphicsScene7addLineEddddRK4QPen(void* qthis, double arg0, double arg1, double arg2, double arg3, void* arg4);
  // proto:  void QGraphicsScene::setBspTreeDepth(int depth);
extern void _ZN14QGraphicsScene15setBspTreeDepthEi(void* qthis, int arg0);
  // proto:  QRectF QGraphicsScene::sceneRect();
extern void _ZNK14QGraphicsScene9sceneRectEv(void* qthis);
  // proto:  QGraphicsWidget * QGraphicsScene::activeWindow();
extern void _ZNK14QGraphicsScene12activeWindowEv(void* qthis);
  // proto:  QBrush QGraphicsScene::backgroundBrush();
extern void _ZNK14QGraphicsScene15backgroundBrushEv(void* qthis);
  // proto:  QGraphicsItem * QGraphicsScene::itemAt(qreal x, qreal y, const QTransform & deviceTransform);
extern void demth_ZNK14QGraphicsScene6itemAtEddRK10QTransform(void* qthis, double arg0, double arg1, void* arg2);
  // proto:  void QGraphicsScene::advance();
extern void _ZN14QGraphicsScene7advanceEv(void* qthis);
  // proto:  void QGraphicsScene::setStickyFocus(bool enabled);
extern void _ZN14QGraphicsScene14setStickyFocusEb(void* qthis, bool arg0);
  // proto:  QList<QGraphicsItem *> QGraphicsScene::selectedItems();
extern void _ZNK14QGraphicsScene13selectedItemsEv(void* qthis);
  // proto:  void QGraphicsScene::clear();
extern void _ZN14QGraphicsScene5clearEv(void* qthis);
  // proto:  void QGraphicsScene::setActivePanel(QGraphicsItem * item);
extern void _ZN14QGraphicsScene14setActivePanelEP13QGraphicsItem(void* qthis, void* arg0);
  // proto:  QGraphicsPixmapItem * QGraphicsScene::addPixmap(const QPixmap & pixmap);
extern void _ZN14QGraphicsScene9addPixmapERK7QPixmap(void* qthis, void* arg0);
  // proto:  QBrush QGraphicsScene::foregroundBrush();
extern void _ZNK14QGraphicsScene15foregroundBrushEv(void* qthis);
  // proto:  QList<QGraphicsView *> QGraphicsScene::views();
extern void _ZNK14QGraphicsScene5viewsEv(void* qthis);
  // proto:  void QGraphicsScene::~QGraphicsScene();
extern void _ZN14QGraphicsSceneD0Ev(void* qthis);
  // proto:  QGraphicsRectItem * QGraphicsScene::addRect(qreal x, qreal y, qreal w, qreal h, const QPen & pen, const QBrush & brush);
extern void demth_ZN14QGraphicsScene7addRectEddddRK4QPenRK6QBrush(void* qthis, double arg0, double arg1, double arg2, double arg3, void* arg4, void* arg5);
  // proto:  int QGraphicsScene::bspTreeDepth();
extern void _ZNK14QGraphicsScene12bspTreeDepthEv(void* qthis);
  // proto:  void QGraphicsScene::setSceneRect(qreal x, qreal y, qreal w, qreal h);
extern void demth_ZN14QGraphicsScene12setSceneRectEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  void QGraphicsScene::setStyle(QStyle * style);
extern void _ZN14QGraphicsScene8setStyleEP6QStyle(void* qthis, void* arg0);
  // proto:  void QGraphicsScene::setPalette(const QPalette & palette);
extern void _ZN14QGraphicsScene10setPaletteERK8QPalette(void* qthis, void* arg0);
  // proto:  void QGraphicsScene::setMinimumRenderSize(qreal minSize);
extern void _ZN14QGraphicsScene20setMinimumRenderSizeEd(void* qthis, double arg0);
  // proto:  void QGraphicsScene::QGraphicsScene(qreal x, qreal y, qreal width, qreal height, QObject * parent);
extern void* dector_ZN14QGraphicsSceneC1EddddP7QObject(double arg0, double arg1, double arg2, double arg3, void* arg4);
extern void _ZN14QGraphicsSceneC1EddddP7QObject(void* qthis, double arg0, double arg1, double arg2, double arg3, void* arg4);
  // proto:  QGraphicsItem * QGraphicsScene::mouseGrabberItem();
extern void _ZNK14QGraphicsScene16mouseGrabberItemEv(void* qthis);
  // proto:  QGraphicsRectItem * QGraphicsScene::addRect(const QRectF & rect, const QPen & pen, const QBrush & brush);
extern void _ZN14QGraphicsScene7addRectERK6QRectFRK4QPenRK6QBrush(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  QGraphicsEllipseItem * QGraphicsScene::addEllipse(const QRectF & rect, const QPen & pen, const QBrush & brush);
extern void _ZN14QGraphicsScene10addEllipseERK6QRectFRK4QPenRK6QBrush(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  qreal QGraphicsScene::height();
extern void demth_ZNK14QGraphicsScene6heightEv(void* qthis);
  // proto:  void QGraphicsScene::setSelectionArea(const QPainterPath & path, const QTransform & deviceTransform);
extern void _ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathRK10QTransform(void* qthis, void* arg0, void* arg1);
  // proto:  QFont QGraphicsScene::font();
extern void _ZNK14QGraphicsScene4fontEv(void* qthis);
  // proto:  void QGraphicsScene::clearSelection();
extern void _ZN14QGraphicsScene14clearSelectionEv(void* qthis);
  // proto:  void QGraphicsScene::QGraphicsScene(const QGraphicsScene & );
extern void* dector_ZN14QGraphicsSceneC1ERKS_(void* arg0);
extern void _ZN14QGraphicsSceneC1ERKS_(void* qthis, void* arg0);
  // proto:  void QGraphicsScene::removeItem(QGraphicsItem * item);
extern void _ZN14QGraphicsScene10removeItemEP13QGraphicsItem(void* qthis, void* arg0);
  // proto:  QGraphicsEllipseItem * QGraphicsScene::addEllipse(qreal x, qreal y, qreal w, qreal h, const QPen & pen, const QBrush & brush);
extern void demth_ZN14QGraphicsScene10addEllipseEddddRK4QPenRK6QBrush(void* qthis, double arg0, double arg1, double arg2, double arg3, void* arg4, void* arg5);
  // proto:  void QGraphicsScene::setActiveWindow(QGraphicsWidget * widget);
extern void _ZN14QGraphicsScene15setActiveWindowEP15QGraphicsWidget(void* qthis, void* arg0);
  // proto:  QGraphicsItem * QGraphicsScene::focusItem();
extern void _ZNK14QGraphicsScene9focusItemEv(void* qthis);
  // proto:  QGraphicsTextItem * QGraphicsScene::addText(const QString & text, const QFont & font);
extern void _ZN14QGraphicsScene7addTextERK7QStringRK5QFont(void* qthis, void* arg0, void* arg1);
  // proto:  void QGraphicsScene::setSortCacheEnabled(bool enabled);
extern void _ZN14QGraphicsScene19setSortCacheEnabledEb(void* qthis, bool arg0);
  // proto:  QGraphicsItem * QGraphicsScene::itemAt(const QPointF & pos, const QTransform & deviceTransform);
extern void _ZNK14QGraphicsScene6itemAtERK7QPointFRK10QTransform(void* qthis, void* arg0, void* arg1);
  // proto:  void QGraphicsScene::destroyItemGroup(QGraphicsItemGroup * group);
extern void _ZN14QGraphicsScene16destroyItemGroupEP18QGraphicsItemGroup(void* qthis, void* arg0);
  // proto:  qreal QGraphicsScene::width();
extern void demth_ZNK14QGraphicsScene5widthEv(void* qthis);
  // proto:  void QGraphicsScene::update(qreal x, qreal y, qreal w, qreal h);
extern void demth_ZN14QGraphicsScene6updateEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  void QGraphicsScene::addItem(QGraphicsItem * item);
extern void _ZN14QGraphicsScene7addItemEP13QGraphicsItem(void* qthis, void* arg0);
  // proto:  void QGraphicsScene::setBackgroundBrush(const QBrush & brush);
extern void _ZN14QGraphicsScene18setBackgroundBrushERK6QBrush(void* qthis, void* arg0);
  // proto:  QGraphicsItem * QGraphicsScene::activePanel();
extern void _ZNK14QGraphicsScene11activePanelEv(void* qthis);
  // proto:  QStyle * QGraphicsScene::style();
extern void _ZNK14QGraphicsScene5styleEv(void* qthis);
  // proto:  void QGraphicsScene::setFont(const QFont & font);
extern void _ZN14QGraphicsScene7setFontERK5QFont(void* qthis, void* arg0);
  // proto:  QGraphicsPathItem * QGraphicsScene::addPath(const QPainterPath & path, const QPen & pen, const QBrush & brush);
extern void _ZN14QGraphicsScene7addPathERK12QPainterPathRK4QPenRK6QBrush(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  bool QGraphicsScene::stickyFocus();
extern void _ZNK14QGraphicsScene11stickyFocusEv(void* qthis);
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

// class sizeof(QGraphicsScene)=1
type QGraphicsScene struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _changed QGraphicsScene_changed_signal;
//  _sceneRectChanged QGraphicsScene_sceneRectChanged_signal;
//  _selectionChanged QGraphicsScene_selectionChanged_signal;
//  _focusItemChanged QGraphicsScene_focusItemChanged_signal;
}

  // proto:  void QGraphicsScene::setForegroundBrush(const QBrush & brush);
func (this *QGraphicsScene) setForegroundBrush(args ...interface{}) () {
  // setForegroundBrush(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene18setForegroundBrushERK6QBrush
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setForegroundBrush", args)
  }

}

  // proto:  void QGraphicsScene::setSceneRect(const QRectF & rect);
func (this *QGraphicsScene) setSceneRect(args ...interface{}) () {
  // setSceneRect(const class QRectF &)
  // setSceneRect(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene12setSceneRectERK6QRectF
  case 1:
    // invoke: _ZN14QGraphicsScene12setSceneRectEdddd
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setSceneRect", args)
  }

}

  // proto:  bool QGraphicsScene::isActive();
func (this *QGraphicsScene) isActive(args ...interface{}) () {
  // isActive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene8isActiveEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "isActive", args)
  }

}

  // proto:  bool QGraphicsScene::hasFocus();
func (this *QGraphicsScene) hasFocus(args ...interface{}) () {
  // hasFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene8hasFocusEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "hasFocus", args)
  }

}

  // proto:  QRectF QGraphicsScene::itemsBoundingRect();
func (this *QGraphicsScene) itemsBoundingRect(args ...interface{}) () {
  // itemsBoundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene17itemsBoundingRectEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "itemsBoundingRect", args)
  }

}

  // proto:  bool QGraphicsScene::sendEvent(QGraphicsItem * item, QEvent * event);
func (this *QGraphicsScene) sendEvent(args ...interface{}) () {
  // sendEvent(class QGraphicsItem *, class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  vtys[0][1] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene9sendEventEP13QGraphicsItemP6QEvent
  default:
    qtrt.ErrorResolve("QGraphicsScene", "sendEvent", args)
  }

}

  // proto:  qreal QGraphicsScene::minimumRenderSize();
func (this *QGraphicsScene) minimumRenderSize(args ...interface{}) () {
  // minimumRenderSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene17minimumRenderSizeEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "minimumRenderSize", args)
  }

}

  // proto:  QPainterPath QGraphicsScene::selectionArea();
func (this *QGraphicsScene) selectionArea(args ...interface{}) () {
  // selectionArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene13selectionAreaEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "selectionArea", args)
  }

}

  // proto:  void QGraphicsScene::update(const QRectF & rect);
func (this *QGraphicsScene) update(args ...interface{}) () {
  // update(const class QRectF &)
  // update(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene6updateERK6QRectF
  case 1:
    // invoke: _ZN14QGraphicsScene6updateEdddd
  default:
    qtrt.ErrorResolve("QGraphicsScene", "update", args)
  }

}

  // proto:  QGraphicsPolygonItem * QGraphicsScene::addPolygon(const QPolygonF & polygon, const QPen & pen, const QBrush & brush);
func (this *QGraphicsScene) addPolygon(args ...interface{}) () {
  // addPolygon(const class QPolygonF &, const class QPen &, const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[0][1] = reflect.TypeOf(QPen{}) // "const QPen &"
  vtys[0][2] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene10addPolygonERK9QPolygonFRK4QPenRK6QBrush
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addPolygon", args)
  }

}

  // proto:  QGraphicsLineItem * QGraphicsScene::addLine(const QLineF & line, const QPen & pen);
func (this *QGraphicsScene) addLine(args ...interface{}) () {
  // addLine(const class QLineF &, const class QPen &)
  // addLine(qreal, qreal, qreal, qreal, const class QPen &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLineF{}) // "const QLineF &"
  vtys[0][1] = reflect.TypeOf(QPen{}) // "const QPen &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][4] = reflect.TypeOf(QPen{}) // "const QPen &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7addLineERK6QLineFRK4QPen
  case 1:
    // invoke: _ZN14QGraphicsScene7addLineEddddRK4QPen
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addLine", args)
  }

}

  // proto:  QPalette QGraphicsScene::palette();
func (this *QGraphicsScene) palette(args ...interface{}) () {
  // palette()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene7paletteEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "palette", args)
  }

}

  // proto:  bool QGraphicsScene::isSortCacheEnabled();
func (this *QGraphicsScene) isSortCacheEnabled(args ...interface{}) () {
  // isSortCacheEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene18isSortCacheEnabledEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "isSortCacheEnabled", args)
  }

}

  // proto:  void QGraphicsScene::QGraphicsScene(const QRectF & sceneRect, QObject * parent);
func NewQGraphicsScene(args ...interface{}) QGraphicsScene {
  return QGraphicsScene{}
}

  // proto:  void QGraphicsScene::clearFocus();
func (this *QGraphicsScene) clearFocus(args ...interface{}) () {
  // clearFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene10clearFocusEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "clearFocus", args)
  }

}

  // proto:  const QMetaObject * QGraphicsScene::metaObject();
func (this *QGraphicsScene) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene10metaObjectEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "metaObject", args)
  }

}

  // proto:  QGraphicsSimpleTextItem * QGraphicsScene::addSimpleText(const QString & text, const QFont & font);
func (this *QGraphicsScene) addSimpleText(args ...interface{}) () {
  // addSimpleText(const class QString &, const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene13addSimpleTextERK7QStringRK5QFont
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addSimpleText", args)
  }

}

  // proto:  void QGraphicsScene::setBspTreeDepth(int depth);
func (this *QGraphicsScene) setBspTreeDepth(args ...interface{}) () {
  // setBspTreeDepth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene15setBspTreeDepthEi
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setBspTreeDepth", args)
  }

}

  // proto:  QRectF QGraphicsScene::sceneRect();
func (this *QGraphicsScene) sceneRect(args ...interface{}) () {
  // sceneRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene9sceneRectEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "sceneRect", args)
  }

}

  // proto:  QGraphicsWidget * QGraphicsScene::activeWindow();
func (this *QGraphicsScene) activeWindow(args ...interface{}) () {
  // activeWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene12activeWindowEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "activeWindow", args)
  }

}

  // proto:  QBrush QGraphicsScene::backgroundBrush();
func (this *QGraphicsScene) backgroundBrush(args ...interface{}) () {
  // backgroundBrush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene15backgroundBrushEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "backgroundBrush", args)
  }

}

  // proto:  QGraphicsItem * QGraphicsScene::itemAt(qreal x, qreal y, const QTransform & deviceTransform);
func (this *QGraphicsScene) itemAt(args ...interface{}) () {
  // itemAt(qreal, qreal, const class QTransform &)
  // itemAt(const class QPointF &, const class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1][1] = reflect.TypeOf(QTransform{}) // "const QTransform &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene6itemAtEddRK10QTransform
  case 1:
    // invoke: _ZNK14QGraphicsScene6itemAtERK7QPointFRK10QTransform
  default:
    qtrt.ErrorResolve("QGraphicsScene", "itemAt", args)
  }

}

  // proto:  void QGraphicsScene::advance();
func (this *QGraphicsScene) advance(args ...interface{}) () {
  // advance()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7advanceEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "advance", args)
  }

}

  // proto:  void QGraphicsScene::setStickyFocus(bool enabled);
func (this *QGraphicsScene) setStickyFocus(args ...interface{}) () {
  // setStickyFocus(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene14setStickyFocusEb
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setStickyFocus", args)
  }

}

  // proto:  QList<QGraphicsItem *> QGraphicsScene::selectedItems();
func (this *QGraphicsScene) selectedItems(args ...interface{}) () {
  // selectedItems()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene13selectedItemsEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "selectedItems", args)
  }

}

  // proto:  void QGraphicsScene::clear();
func (this *QGraphicsScene) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene5clearEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "clear", args)
  }

}

  // proto:  void QGraphicsScene::setActivePanel(QGraphicsItem * item);
func (this *QGraphicsScene) setActivePanel(args ...interface{}) () {
  // setActivePanel(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene14setActivePanelEP13QGraphicsItem
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setActivePanel", args)
  }

}

  // proto:  QGraphicsPixmapItem * QGraphicsScene::addPixmap(const QPixmap & pixmap);
func (this *QGraphicsScene) addPixmap(args ...interface{}) () {
  // addPixmap(const class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene9addPixmapERK7QPixmap
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addPixmap", args)
  }

}

  // proto:  QBrush QGraphicsScene::foregroundBrush();
func (this *QGraphicsScene) foregroundBrush(args ...interface{}) () {
  // foregroundBrush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene15foregroundBrushEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "foregroundBrush", args)
  }

}

  // proto:  QList<QGraphicsView *> QGraphicsScene::views();
func (this *QGraphicsScene) views(args ...interface{}) () {
  // views()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene5viewsEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "views", args)
  }

}

  // proto:  void QGraphicsScene::~QGraphicsScene();
func (this *QGraphicsScene) FreeQGraphicsScene(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsScene", "~QGraphicsScene", args)
  }

}

  // proto:  QGraphicsRectItem * QGraphicsScene::addRect(qreal x, qreal y, qreal w, qreal h, const QPen & pen, const QBrush & brush);
func (this *QGraphicsScene) addRect(args ...interface{}) () {
  // addRect(qreal, qreal, qreal, qreal, const class QPen &, const class QBrush &)
  // addRect(const class QRectF &, const class QPen &, const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][4] = reflect.TypeOf(QPen{}) // "const QPen &"
  vtys[0][5] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1][1] = reflect.TypeOf(QPen{}) // "const QPen &"
  vtys[1][2] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7addRectEddddRK4QPenRK6QBrush
  case 1:
    // invoke: _ZN14QGraphicsScene7addRectERK6QRectFRK4QPenRK6QBrush
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addRect", args)
  }

}

  // proto:  int QGraphicsScene::bspTreeDepth();
func (this *QGraphicsScene) bspTreeDepth(args ...interface{}) () {
  // bspTreeDepth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene12bspTreeDepthEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "bspTreeDepth", args)
  }

}

  // proto:  void QGraphicsScene::setStyle(QStyle * style);
func (this *QGraphicsScene) setStyle(args ...interface{}) () {
  // setStyle(class QStyle *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStyle{}) // "QStyle *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene8setStyleEP6QStyle
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setStyle", args)
  }

}

  // proto:  void QGraphicsScene::setPalette(const QPalette & palette);
func (this *QGraphicsScene) setPalette(args ...interface{}) () {
  // setPalette(const class QPalette &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPalette{}) // "const QPalette &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene10setPaletteERK8QPalette
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setPalette", args)
  }

}

  // proto:  void QGraphicsScene::setMinimumRenderSize(qreal minSize);
func (this *QGraphicsScene) setMinimumRenderSize(args ...interface{}) () {
  // setMinimumRenderSize(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene20setMinimumRenderSizeEd
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setMinimumRenderSize", args)
  }

}

  // proto:  QGraphicsItem * QGraphicsScene::mouseGrabberItem();
func (this *QGraphicsScene) mouseGrabberItem(args ...interface{}) () {
  // mouseGrabberItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene16mouseGrabberItemEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "mouseGrabberItem", args)
  }

}

  // proto:  QGraphicsEllipseItem * QGraphicsScene::addEllipse(const QRectF & rect, const QPen & pen, const QBrush & brush);
func (this *QGraphicsScene) addEllipse(args ...interface{}) () {
  // addEllipse(const class QRectF &, const class QPen &, const class QBrush &)
  // addEllipse(qreal, qreal, qreal, qreal, const class QPen &, const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][1] = reflect.TypeOf(QPen{}) // "const QPen &"
  vtys[0][2] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][4] = reflect.TypeOf(QPen{}) // "const QPen &"
  vtys[1][5] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene10addEllipseERK6QRectFRK4QPenRK6QBrush
  case 1:
    // invoke: _ZN14QGraphicsScene10addEllipseEddddRK4QPenRK6QBrush
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addEllipse", args)
  }

}

  // proto:  qreal QGraphicsScene::height();
func (this *QGraphicsScene) height(args ...interface{}) () {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene6heightEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "height", args)
  }

}

  // proto:  void QGraphicsScene::setSelectionArea(const QPainterPath & path, const QTransform & deviceTransform);
func (this *QGraphicsScene) setSelectionArea(args ...interface{}) () {
  // setSelectionArea(const class QPainterPath &, Qt::ItemSelectionMode, const class QTransform &)
  // setSelectionArea(const class QPainterPath &, const class QTransform &)
  // setSelectionArea(const class QPainterPath &, Qt::ItemSelectionOperation, Qt::ItemSelectionMode, const class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[0][1] = qtrt.Int32Ty(false) // "Qt::ItemSelectionMode"
  vtys[0][2] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[1][1] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[2][1] = qtrt.Int32Ty(false) // "Qt::ItemSelectionOperation"
  vtys[2][2] = qtrt.Int32Ty(false) // "Qt::ItemSelectionMode"
  vtys[2][3] = reflect.TypeOf(QTransform{}) // "const QTransform &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathN2Qt17ItemSelectionModeERK10QTransform
  case 1:
    // invoke: _ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathRK10QTransform
  case 2:
    // invoke: _ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathN2Qt22ItemSelectionOperationENS3_17ItemSelectionModeERK10QTransform
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setSelectionArea", args)
  }

}

  // proto:  QFont QGraphicsScene::font();
func (this *QGraphicsScene) font(args ...interface{}) () {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene4fontEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "font", args)
  }

}

  // proto:  void QGraphicsScene::clearSelection();
func (this *QGraphicsScene) clearSelection(args ...interface{}) () {
  // clearSelection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene14clearSelectionEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "clearSelection", args)
  }

}

  // proto:  void QGraphicsScene::removeItem(QGraphicsItem * item);
func (this *QGraphicsScene) removeItem(args ...interface{}) () {
  // removeItem(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene10removeItemEP13QGraphicsItem
  default:
    qtrt.ErrorResolve("QGraphicsScene", "removeItem", args)
  }

}

  // proto:  void QGraphicsScene::setActiveWindow(QGraphicsWidget * widget);
func (this *QGraphicsScene) setActiveWindow(args ...interface{}) () {
  // setActiveWindow(class QGraphicsWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsWidget{}) // "QGraphicsWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene15setActiveWindowEP15QGraphicsWidget
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setActiveWindow", args)
  }

}

  // proto:  QGraphicsItem * QGraphicsScene::focusItem();
func (this *QGraphicsScene) focusItem(args ...interface{}) () {
  // focusItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene9focusItemEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "focusItem", args)
  }

}

  // proto:  QGraphicsTextItem * QGraphicsScene::addText(const QString & text, const QFont & font);
func (this *QGraphicsScene) addText(args ...interface{}) () {
  // addText(const class QString &, const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7addTextERK7QStringRK5QFont
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addText", args)
  }

}

  // proto:  void QGraphicsScene::setSortCacheEnabled(bool enabled);
func (this *QGraphicsScene) setSortCacheEnabled(args ...interface{}) () {
  // setSortCacheEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene19setSortCacheEnabledEb
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setSortCacheEnabled", args)
  }

}

  // proto:  void QGraphicsScene::destroyItemGroup(QGraphicsItemGroup * group);
func (this *QGraphicsScene) destroyItemGroup(args ...interface{}) () {
  // destroyItemGroup(class QGraphicsItemGroup *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItemGroup{}) // "QGraphicsItemGroup *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene16destroyItemGroupEP18QGraphicsItemGroup
  default:
    qtrt.ErrorResolve("QGraphicsScene", "destroyItemGroup", args)
  }

}

  // proto:  qreal QGraphicsScene::width();
func (this *QGraphicsScene) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene5widthEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "width", args)
  }

}

  // proto:  void QGraphicsScene::addItem(QGraphicsItem * item);
func (this *QGraphicsScene) addItem(args ...interface{}) () {
  // addItem(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7addItemEP13QGraphicsItem
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addItem", args)
  }

}

  // proto:  void QGraphicsScene::setBackgroundBrush(const QBrush & brush);
func (this *QGraphicsScene) setBackgroundBrush(args ...interface{}) () {
  // setBackgroundBrush(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene18setBackgroundBrushERK6QBrush
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setBackgroundBrush", args)
  }

}

  // proto:  QGraphicsItem * QGraphicsScene::activePanel();
func (this *QGraphicsScene) activePanel(args ...interface{}) () {
  // activePanel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene11activePanelEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "activePanel", args)
  }

}

  // proto:  QStyle * QGraphicsScene::style();
func (this *QGraphicsScene) style(args ...interface{}) () {
  // style()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene5styleEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "style", args)
  }

}

  // proto:  void QGraphicsScene::setFont(const QFont & font);
func (this *QGraphicsScene) setFont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7setFontERK5QFont
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setFont", args)
  }

}

  // proto:  QGraphicsPathItem * QGraphicsScene::addPath(const QPainterPath & path, const QPen & pen, const QBrush & brush);
func (this *QGraphicsScene) addPath(args ...interface{}) () {
  // addPath(const class QPainterPath &, const class QPen &, const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[0][1] = reflect.TypeOf(QPen{}) // "const QPen &"
  vtys[0][2] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7addPathERK12QPainterPathRK4QPenRK6QBrush
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addPath", args)
  }

}

  // proto:  bool QGraphicsScene::stickyFocus();
func (this *QGraphicsScene) stickyFocus(args ...interface{}) () {
  // stickyFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene11stickyFocusEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "stickyFocus", args)
  }

}

// <= body block end

