package qtwidgets
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
  // proto:  QGraphicsRectItem * QGraphicsScene::addRect(qreal x, qreal y, qreal w, qreal h, const QPen & pen, const QBrush & brush);
extern void C_ZN14QGraphicsScene7addRectEddddRK4QPenRK6QBrush(void* qthis, double arg0, double arg1, double arg2, double arg3, void* arg4, void* arg5); // 2
  // proto:  QGraphicsRectItem * QGraphicsScene::addRect(const QRectF & rect, const QPen & pen, const QBrush & brush);
extern void C_ZN14QGraphicsScene7addRectERK6QRectFRK4QPenRK6QBrush(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  bool QGraphicsScene::hasFocus();
extern bool C_ZNK14QGraphicsScene8hasFocusEv(void* qthis); // 4
  // proto:  void QGraphicsScene::QGraphicsScene(QObject * parent);
extern void* C_ZN14QGraphicsSceneC2EP7QObject(void* arg0); // 3
  // proto:  void QGraphicsScene::QGraphicsScene(const QRectF & sceneRect, QObject * parent);
extern void* C_ZN14QGraphicsSceneC2ERK6QRectFP7QObject(void* arg0, void* arg1); // 3
  // proto:  void QGraphicsScene::QGraphicsScene(qreal x, qreal y, qreal width, qreal height, QObject * parent);
extern void* C_ZN14QGraphicsSceneC2EddddP7QObject(double arg0, double arg1, double arg2, double arg3, void* arg4); // 3
  // proto:  void QGraphicsScene::setStyle(QStyle * style);
extern void C_ZN14QGraphicsScene8setStyleEP6QStyle(void* qthis, void* arg0); // 4
  // proto:  QPainterPath QGraphicsScene::selectionArea();
extern void* C_ZNK14QGraphicsScene13selectionAreaEv(void* qthis); // 4
  // proto:  QFont QGraphicsScene::font();
extern void* C_ZNK14QGraphicsScene4fontEv(void* qthis); // 4
  // proto:  QGraphicsEllipseItem * QGraphicsScene::addEllipse(qreal x, qreal y, qreal w, qreal h, const QPen & pen, const QBrush & brush);
extern void C_ZN14QGraphicsScene10addEllipseEddddRK4QPenRK6QBrush(void* qthis, double arg0, double arg1, double arg2, double arg3, void* arg4, void* arg5); // 2
  // proto:  QGraphicsEllipseItem * QGraphicsScene::addEllipse(const QRectF & rect, const QPen & pen, const QBrush & brush);
extern void C_ZN14QGraphicsScene10addEllipseERK6QRectFRK4QPenRK6QBrush(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  QGraphicsPolygonItem * QGraphicsScene::addPolygon(const QPolygonF & polygon, const QPen & pen, const QBrush & brush);
extern void C_ZN14QGraphicsScene10addPolygonERK9QPolygonFRK4QPenRK6QBrush(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QGraphicsScene::setBspTreeDepth(int depth);
extern void C_ZN14QGraphicsScene15setBspTreeDepthEi(void* qthis, int32_t arg0); // 4
  // proto:  QList<QGraphicsView *> QGraphicsScene::views();
extern void C_ZNK14QGraphicsScene5viewsEv(void* qthis); // 4
  // proto:  void QGraphicsScene::setBackgroundBrush(const QBrush & brush);
extern void C_ZN14QGraphicsScene18setBackgroundBrushERK6QBrush(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsScene::removeItem(QGraphicsItem * item);
extern void C_ZN14QGraphicsScene10removeItemEP13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  int QGraphicsScene::bspTreeDepth();
extern int32_t C_ZNK14QGraphicsScene12bspTreeDepthEv(void* qthis); // 4
  // proto:  bool QGraphicsScene::isSortCacheEnabled();
extern bool C_ZNK14QGraphicsScene18isSortCacheEnabledEv(void* qthis); // 4
  // proto:  const QMetaObject * QGraphicsScene::metaObject();
extern void C_ZNK14QGraphicsScene10metaObjectEv(void* qthis); // 4
  // proto:  QRectF QGraphicsScene::sceneRect();
extern void* C_ZNK14QGraphicsScene9sceneRectEv(void* qthis); // 4
  // proto:  QGraphicsItem * QGraphicsScene::itemAt(const QPointF & pos, const QTransform & deviceTransform);
extern void C_ZNK14QGraphicsScene6itemAtERK7QPointFRK10QTransform(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QGraphicsItem * QGraphicsScene::itemAt(qreal x, qreal y, const QTransform & deviceTransform);
extern void C_ZNK14QGraphicsScene6itemAtEddRK10QTransform(void* qthis, double arg0, double arg1, void* arg2); // 2
  // proto:  void QGraphicsScene::clearFocus();
extern void C_ZN14QGraphicsScene10clearFocusEv(void* qthis); // 4
  // proto:  QGraphicsPixmapItem * QGraphicsScene::addPixmap(const QPixmap & pixmap);
extern void C_ZN14QGraphicsScene9addPixmapERK7QPixmap(void* qthis, void* arg0); // 4
  // proto:  bool QGraphicsScene::stickyFocus();
extern bool C_ZNK14QGraphicsScene11stickyFocusEv(void* qthis); // 4
  // proto:  QPalette QGraphicsScene::palette();
extern void* C_ZNK14QGraphicsScene7paletteEv(void* qthis); // 4
  // proto:  void QGraphicsScene::setMinimumRenderSize(qreal minSize);
extern void C_ZN14QGraphicsScene20setMinimumRenderSizeEd(void* qthis, double arg0); // 4
  // proto:  QGraphicsWidget * QGraphicsScene::activeWindow();
extern void C_ZNK14QGraphicsScene12activeWindowEv(void* qthis); // 4
  // proto:  void QGraphicsScene::~QGraphicsScene();
extern void C_ZN14QGraphicsSceneD2Ev(void* qthis); // 4
  // proto:  void QGraphicsScene::addItem(QGraphicsItem * item);
extern void C_ZN14QGraphicsScene7addItemEP13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  QStyle * QGraphicsScene::style();
extern void* C_ZNK14QGraphicsScene5styleEv(void* qthis); // 4
  // proto:  void QGraphicsScene::setActivePanel(QGraphicsItem * item);
extern void C_ZN14QGraphicsScene14setActivePanelEP13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  bool QGraphicsScene::isActive();
extern bool C_ZNK14QGraphicsScene8isActiveEv(void* qthis); // 4
  // proto:  void QGraphicsScene::advance();
extern void C_ZN14QGraphicsScene7advanceEv(void* qthis); // 4
  // proto:  QRectF QGraphicsScene::itemsBoundingRect();
extern void* C_ZNK14QGraphicsScene17itemsBoundingRectEv(void* qthis); // 4
  // proto:  void QGraphicsScene::setSceneRect(const QRectF & rect);
extern void C_ZN14QGraphicsScene12setSceneRectERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsScene::setSceneRect(qreal x, qreal y, qreal w, qreal h);
extern void C_ZN14QGraphicsScene12setSceneRectEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  QGraphicsScene::ItemIndexMethod QGraphicsScene::itemIndexMethod();
extern void C_ZNK14QGraphicsScene15itemIndexMethodEv(void* qthis); // 4
  // proto:  QGraphicsLineItem * QGraphicsScene::addLine(const QLineF & line, const QPen & pen);
extern void C_ZN14QGraphicsScene7addLineERK6QLineFRK4QPen(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QGraphicsLineItem * QGraphicsScene::addLine(qreal x1, qreal y1, qreal x2, qreal y2, const QPen & pen);
extern void C_ZN14QGraphicsScene7addLineEddddRK4QPen(void* qthis, double arg0, double arg1, double arg2, double arg3, void* arg4); // 2
  // proto:  qreal QGraphicsScene::minimumRenderSize();
extern double C_ZNK14QGraphicsScene17minimumRenderSizeEv(void* qthis); // 4
  // proto:  bool QGraphicsScene::sendEvent(QGraphicsItem * item, QEvent * event);
extern bool C_ZN14QGraphicsScene9sendEventEP13QGraphicsItemP6QEvent(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QGraphicsScene::setPalette(const QPalette & palette);
extern void C_ZN14QGraphicsScene10setPaletteERK8QPalette(void* qthis, void* arg0); // 4
  // proto:  QGraphicsItem * QGraphicsScene::mouseGrabberItem();
extern void C_ZNK14QGraphicsScene16mouseGrabberItemEv(void* qthis); // 4
  // proto:  void QGraphicsScene::setSelectionArea(const QPainterPath & path, const QTransform & deviceTransform);
extern void C_ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathRK10QTransform(void* qthis, void* arg0, void* arg1); // 4
  // proto:  qreal QGraphicsScene::width();
extern double C_ZNK14QGraphicsScene5widthEv(void* qthis); // 2
  // proto:  QBrush QGraphicsScene::foregroundBrush();
extern void* C_ZNK14QGraphicsScene15foregroundBrushEv(void* qthis); // 4
  // proto:  void QGraphicsScene::setForegroundBrush(const QBrush & brush);
extern void C_ZN14QGraphicsScene18setForegroundBrushERK6QBrush(void* qthis, void* arg0); // 4
  // proto:  QGraphicsPathItem * QGraphicsScene::addPath(const QPainterPath & path, const QPen & pen, const QBrush & brush);
extern void C_ZN14QGraphicsScene7addPathERK12QPainterPathRK4QPenRK6QBrush(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  QList<QGraphicsItem *> QGraphicsScene::selectedItems();
extern void C_ZNK14QGraphicsScene13selectedItemsEv(void* qthis); // 4
  // proto:  void QGraphicsScene::setActiveWindow(QGraphicsWidget * widget);
extern void C_ZN14QGraphicsScene15setActiveWindowEP15QGraphicsWidget(void* qthis, void* arg0); // 4
  // proto:  QBrush QGraphicsScene::backgroundBrush();
extern void* C_ZNK14QGraphicsScene15backgroundBrushEv(void* qthis); // 4
  // proto:  QGraphicsSimpleTextItem * QGraphicsScene::addSimpleText(const QString & text, const QFont & font);
extern void C_ZN14QGraphicsScene13addSimpleTextERK7QStringRK5QFont(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QGraphicsScene::clearSelection();
extern void C_ZN14QGraphicsScene14clearSelectionEv(void* qthis); // 4
  // proto:  QGraphicsItem * QGraphicsScene::activePanel();
extern void C_ZNK14QGraphicsScene11activePanelEv(void* qthis); // 4
  // proto:  QGraphicsItem * QGraphicsScene::focusItem();
extern void C_ZNK14QGraphicsScene9focusItemEv(void* qthis); // 4
  // proto:  void QGraphicsScene::destroyItemGroup(QGraphicsItemGroup * group);
extern void C_ZN14QGraphicsScene16destroyItemGroupEP18QGraphicsItemGroup(void* qthis, void* arg0); // 4
  // proto:  QGraphicsTextItem * QGraphicsScene::addText(const QString & text, const QFont & font);
extern void C_ZN14QGraphicsScene7addTextERK7QStringRK5QFont(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QGraphicsScene::update(const QRectF & rect);
extern void C_ZN14QGraphicsScene6updateERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsScene::update(qreal x, qreal y, qreal w, qreal h);
extern void C_ZN14QGraphicsScene6updateEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  void QGraphicsScene::setSortCacheEnabled(bool enabled);
extern void C_ZN14QGraphicsScene19setSortCacheEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QGraphicsScene::setStickyFocus(bool enabled);
extern void C_ZN14QGraphicsScene14setStickyFocusEb(void* qthis, bool arg0); // 4
  // proto:  void QGraphicsScene::clear();
extern void C_ZN14QGraphicsScene5clearEv(void* qthis); // 4
  // proto:  qreal QGraphicsScene::height();
extern double C_ZNK14QGraphicsScene6heightEv(void* qthis); // 2
  // proto:  void QGraphicsScene::setFont(const QFont & font);
extern void C_ZN14QGraphicsScene7setFontERK5QFont(void* qthis, void* arg0); // 4
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
}

// class sizeof(QGraphicsScene)=1
type QGraphicsScene struct {
  /*qbase*/ qtcore.QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _changed QGraphicsScene_changed_signal;
//  _sceneRectChanged QGraphicsScene_sceneRectChanged_signal;
//  _selectionChanged QGraphicsScene_selectionChanged_signal;
//  _focusItemChanged QGraphicsScene_focusItemChanged_signal;
}

// addRect(qreal, qreal, qreal, qreal, const class QPen &, const class QBrush &)
func (this *QGraphicsScene) Addrect(args ...interface{}) () {
  // addRect(qreal, qreal, qreal, qreal, const class QPen &, const class QBrush &)
  // addRect(const class QRectF &, const class QPen &, const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][4] = reflect.TypeOf(qtgui.QPen{}) // "const QPen &"
  vtys[0][5] = reflect.TypeOf(qtgui.QBrush{}) // "const QBrush &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[1][1] = reflect.TypeOf(qtgui.QPen{}) // "const QPen &"
  vtys[1][2] = reflect.TypeOf(qtgui.QBrush{}) // "const QBrush &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7addRectEddddRK4QPenRK6QBrush
    // invoke: QGraphicsRectItem * addRect(qreal, qreal, qreal, qreal, const class QPen &, const class QBrush &)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(*qtgui.QPen).Qclsinst
    if false {fmt.Println(arg4)}
    var arg5 = args[5].(*qtgui.QBrush).Qclsinst
    if false {fmt.Println(arg5)}
    C.C_ZN14QGraphicsScene7addRectEddddRK4QPenRK6QBrush(this.Qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  case 1:
    // invoke: _ZN14QGraphicsScene7addRectERK6QRectFRK4QPenRK6QBrush
    // invoke: QGraphicsRectItem * addRect(const class QRectF &, const class QPen &, const class QBrush &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtgui.QPen).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtgui.QBrush).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN14QGraphicsScene7addRectERK6QRectFRK4QPenRK6QBrush(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addRect", args)
  }

  return
}

// hasFocus()
func (this *QGraphicsScene) Hasfocus(args ...interface{}) (ret interface{}) {
  // hasFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene8hasFocusEv
    // invoke: bool hasFocus()
    var ret0 = C.C_ZNK14QGraphicsScene8hasFocusEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsScene", "hasFocus", args)
  }

  return
}

// QGraphicsScene(class QObject *)
func NewQGraphicsScene(args ...interface{}) *QGraphicsScene {
  // QGraphicsScene(class QObject *)
  // QGraphicsScene(const class QRectF &, class QObject *)
  // QGraphicsScene(qreal, qreal, qreal, qreal, class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[1][1] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][4] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsSceneC1EP7QObject
    // invoke: void QGraphicsScene(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QGraphicsSceneC2EP7QObject(arg0)
    return &QGraphicsScene{Qclsinst:qthis}
  case 1:
    // invoke: _ZN14QGraphicsSceneC1ERK6QRectFP7QObject
    // invoke: void QGraphicsScene(const class QRectF &, class QObject *)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QGraphicsSceneC2ERK6QRectFP7QObject(arg0, arg1)
    return &QGraphicsScene{Qclsinst:qthis}
  case 2:
    // invoke: _ZN14QGraphicsSceneC1EddddP7QObject
    // invoke: void QGraphicsScene(qreal, qreal, qreal, qreal, class QObject *)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg4)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QGraphicsSceneC2EddddP7QObject(arg0, arg1, arg2, arg3, arg4)
    return &QGraphicsScene{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsScene", "QGraphicsScene", args)
  }

  return nil // QGraphicsScene{Qclsinst:qthis}
}

// setStyle(class QStyle *)
func (this *QGraphicsScene) Setstyle(args ...interface{}) () {
  // setStyle(class QStyle *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStyle{}) // "QStyle *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene8setStyleEP6QStyle
    // invoke: void setStyle(class QStyle *)
    var arg0 = args[0].(*QStyle).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QGraphicsScene8setStyleEP6QStyle(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setStyle", args)
  }

  return
}

// selectionArea()
func (this *QGraphicsScene) Selectionarea(args ...interface{}) (ret interface{}) {
  // selectionArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene13selectionAreaEv
    // invoke: QPainterPath selectionArea()
    var ret0 = C.C_ZNK14QGraphicsScene13selectionAreaEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsScene", "selectionArea", args)
  }

  return
}

// font()
func (this *QGraphicsScene) Font(args ...interface{}) (ret interface{}) {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene4fontEv
    // invoke: QFont font()
    var ret0 = C.C_ZNK14QGraphicsScene4fontEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QFont{}) // "QFont"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsScene", "font", args)
  }

  return
}

// addEllipse(qreal, qreal, qreal, qreal, const class QPen &, const class QBrush &)
func (this *QGraphicsScene) Addellipse(args ...interface{}) () {
  // addEllipse(qreal, qreal, qreal, qreal, const class QPen &, const class QBrush &)
  // addEllipse(const class QRectF &, const class QPen &, const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][4] = reflect.TypeOf(qtgui.QPen{}) // "const QPen &"
  vtys[0][5] = reflect.TypeOf(qtgui.QBrush{}) // "const QBrush &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[1][1] = reflect.TypeOf(qtgui.QPen{}) // "const QPen &"
  vtys[1][2] = reflect.TypeOf(qtgui.QBrush{}) // "const QBrush &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene10addEllipseEddddRK4QPenRK6QBrush
    // invoke: QGraphicsEllipseItem * addEllipse(qreal, qreal, qreal, qreal, const class QPen &, const class QBrush &)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(*qtgui.QPen).Qclsinst
    if false {fmt.Println(arg4)}
    var arg5 = args[5].(*qtgui.QBrush).Qclsinst
    if false {fmt.Println(arg5)}
    C.C_ZN14QGraphicsScene10addEllipseEddddRK4QPenRK6QBrush(this.Qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  case 1:
    // invoke: _ZN14QGraphicsScene10addEllipseERK6QRectFRK4QPenRK6QBrush
    // invoke: QGraphicsEllipseItem * addEllipse(const class QRectF &, const class QPen &, const class QBrush &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtgui.QPen).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtgui.QBrush).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN14QGraphicsScene10addEllipseERK6QRectFRK4QPenRK6QBrush(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addEllipse", args)
  }

  return
}

// addPolygon(const class QPolygonF &, const class QPen &, const class QBrush &)
func (this *QGraphicsScene) Addpolygon(args ...interface{}) () {
  // addPolygon(const class QPolygonF &, const class QPen &, const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QPolygonF{}) // "const QPolygonF &"
  vtys[0][1] = reflect.TypeOf(qtgui.QPen{}) // "const QPen &"
  vtys[0][2] = reflect.TypeOf(qtgui.QBrush{}) // "const QBrush &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene10addPolygonERK9QPolygonFRK4QPenRK6QBrush
    // invoke: QGraphicsPolygonItem * addPolygon(const class QPolygonF &, const class QPen &, const class QBrush &)
    var arg0 = args[0].(*qtgui.QPolygonF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtgui.QPen).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtgui.QBrush).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN14QGraphicsScene10addPolygonERK9QPolygonFRK4QPenRK6QBrush(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addPolygon", args)
  }

  return
}

// setBspTreeDepth(int)
func (this *QGraphicsScene) Setbsptreedepth(args ...interface{}) () {
  // setBspTreeDepth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene15setBspTreeDepthEi
    // invoke: void setBspTreeDepth(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QGraphicsScene15setBspTreeDepthEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setBspTreeDepth", args)
  }

  return
}

// views()
func (this *QGraphicsScene) Views(args ...interface{}) () {
  // views()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene5viewsEv
    // invoke: QList<QGraphicsView *> views()
    C.C_ZNK14QGraphicsScene5viewsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "views", args)
  }

  return
}

// setBackgroundBrush(const class QBrush &)
func (this *QGraphicsScene) Setbackgroundbrush(args ...interface{}) () {
  // setBackgroundBrush(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QBrush{}) // "const QBrush &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene18setBackgroundBrushERK6QBrush
    // invoke: void setBackgroundBrush(const class QBrush &)
    var arg0 = args[0].(*qtgui.QBrush).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QGraphicsScene18setBackgroundBrushERK6QBrush(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setBackgroundBrush", args)
  }

  return
}

// removeItem(class QGraphicsItem *)
func (this *QGraphicsScene) Removeitem(args ...interface{}) () {
  // removeItem(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene10removeItemEP13QGraphicsItem
    // invoke: void removeItem(class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QGraphicsScene10removeItemEP13QGraphicsItem(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "removeItem", args)
  }

  return
}

// bspTreeDepth()
func (this *QGraphicsScene) Bsptreedepth(args ...interface{}) (ret interface{}) {
  // bspTreeDepth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene12bspTreeDepthEv
    // invoke: int bspTreeDepth()
    var ret0 = C.C_ZNK14QGraphicsScene12bspTreeDepthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsScene", "bspTreeDepth", args)
  }

  return
}

// isSortCacheEnabled()
func (this *QGraphicsScene) Issortcacheenabled(args ...interface{}) (ret interface{}) {
  // isSortCacheEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene18isSortCacheEnabledEv
    // invoke: bool isSortCacheEnabled()
    var ret0 = C.C_ZNK14QGraphicsScene18isSortCacheEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsScene", "isSortCacheEnabled", args)
  }

  return
}

// metaObject()
func (this *QGraphicsScene) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK14QGraphicsScene10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "metaObject", args)
  }

  return
}

// sceneRect()
func (this *QGraphicsScene) Scenerect(args ...interface{}) (ret interface{}) {
  // sceneRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene9sceneRectEv
    // invoke: QRectF sceneRect()
    var ret0 = C.C_ZNK14QGraphicsScene9sceneRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsScene", "sceneRect", args)
  }

  return
}

// itemAt(const class QPointF &, const class QTransform &)
func (this *QGraphicsScene) Itemat(args ...interface{}) () {
  // itemAt(const class QPointF &, const class QTransform &)
  // itemAt(qreal, qreal, const class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[0][1] = reflect.TypeOf(qtgui.QTransform{}) // "const QTransform &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = reflect.TypeOf(qtgui.QTransform{}) // "const QTransform &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene6itemAtERK7QPointFRK10QTransform
    // invoke: QGraphicsItem * itemAt(const class QPointF &, const class QTransform &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtgui.QTransform).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZNK14QGraphicsScene6itemAtERK7QPointFRK10QTransform(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK14QGraphicsScene6itemAtEddRK10QTransform
    // invoke: QGraphicsItem * itemAt(qreal, qreal, const class QTransform &)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtgui.QTransform).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZNK14QGraphicsScene6itemAtEddRK10QTransform(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "itemAt", args)
  }

  return
}

// clearFocus()
func (this *QGraphicsScene) Clearfocus(args ...interface{}) () {
  // clearFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene10clearFocusEv
    // invoke: void clearFocus()
    C.C_ZN14QGraphicsScene10clearFocusEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "clearFocus", args)
  }

  return
}

// addPixmap(const class QPixmap &)
func (this *QGraphicsScene) Addpixmap(args ...interface{}) () {
  // addPixmap(const class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QPixmap{}) // "const QPixmap &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene9addPixmapERK7QPixmap
    // invoke: QGraphicsPixmapItem * addPixmap(const class QPixmap &)
    var arg0 = args[0].(*qtgui.QPixmap).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QGraphicsScene9addPixmapERK7QPixmap(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addPixmap", args)
  }

  return
}

// stickyFocus()
func (this *QGraphicsScene) Stickyfocus(args ...interface{}) (ret interface{}) {
  // stickyFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene11stickyFocusEv
    // invoke: bool stickyFocus()
    var ret0 = C.C_ZNK14QGraphicsScene11stickyFocusEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsScene", "stickyFocus", args)
  }

  return
}

// palette()
func (this *QGraphicsScene) Palette(args ...interface{}) (ret interface{}) {
  // palette()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene7paletteEv
    // invoke: QPalette palette()
    var ret0 = C.C_ZNK14QGraphicsScene7paletteEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QPalette{}) // "QPalette"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsScene", "palette", args)
  }

  return
}

// setMinimumRenderSize(qreal)
func (this *QGraphicsScene) Setminimumrendersize(args ...interface{}) () {
  // setMinimumRenderSize(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene20setMinimumRenderSizeEd
    // invoke: void setMinimumRenderSize(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN14QGraphicsScene20setMinimumRenderSizeEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setMinimumRenderSize", args)
  }

  return
}

// activeWindow()
func (this *QGraphicsScene) Activewindow(args ...interface{}) () {
  // activeWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene12activeWindowEv
    // invoke: QGraphicsWidget * activeWindow()
    C.C_ZNK14QGraphicsScene12activeWindowEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "activeWindow", args)
  }

  return
}

// ~QGraphicsScene()
func (this *QGraphicsScene) Freeqgraphicsscene(args ...interface{}) () {
  // ~QGraphicsScene()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsSceneD0Ev
    // invoke: void ~QGraphicsScene()
    C.C_ZN14QGraphicsSceneD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "~QGraphicsScene", args)
  }

  return
}

// addItem(class QGraphicsItem *)
func (this *QGraphicsScene) Additem(args ...interface{}) () {
  // addItem(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7addItemEP13QGraphicsItem
    // invoke: void addItem(class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QGraphicsScene7addItemEP13QGraphicsItem(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addItem", args)
  }

  return
}

// style()
func (this *QGraphicsScene) Style(args ...interface{}) (ret interface{}) {
  // style()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene5styleEv
    // invoke: QStyle * style()
    var ret0 = C.C_ZNK14QGraphicsScene5styleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStyle{}) // "QStyle *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsScene", "style", args)
  }

  return
}

// setActivePanel(class QGraphicsItem *)
func (this *QGraphicsScene) Setactivepanel(args ...interface{}) () {
  // setActivePanel(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene14setActivePanelEP13QGraphicsItem
    // invoke: void setActivePanel(class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QGraphicsScene14setActivePanelEP13QGraphicsItem(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setActivePanel", args)
  }

  return
}

// isActive()
func (this *QGraphicsScene) Isactive(args ...interface{}) (ret interface{}) {
  // isActive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene8isActiveEv
    // invoke: bool isActive()
    var ret0 = C.C_ZNK14QGraphicsScene8isActiveEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsScene", "isActive", args)
  }

  return
}

// advance()
func (this *QGraphicsScene) Advance(args ...interface{}) () {
  // advance()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7advanceEv
    // invoke: void advance()
    C.C_ZN14QGraphicsScene7advanceEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "advance", args)
  }

  return
}

// itemsBoundingRect()
func (this *QGraphicsScene) Itemsboundingrect(args ...interface{}) (ret interface{}) {
  // itemsBoundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene17itemsBoundingRectEv
    // invoke: QRectF itemsBoundingRect()
    var ret0 = C.C_ZNK14QGraphicsScene17itemsBoundingRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsScene", "itemsBoundingRect", args)
  }

  return
}

// setSceneRect(const class QRectF &)
func (this *QGraphicsScene) Setscenerect(args ...interface{}) () {
  // setSceneRect(const class QRectF &)
  // setSceneRect(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene12setSceneRectERK6QRectF
    // invoke: void setSceneRect(const class QRectF &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QGraphicsScene12setSceneRectERK6QRectF(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN14QGraphicsScene12setSceneRectEdddd
    // invoke: void setSceneRect(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN14QGraphicsScene12setSceneRectEdddd(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setSceneRect", args)
  }

  return
}

// itemIndexMethod()
func (this *QGraphicsScene) Itemindexmethod(args ...interface{}) () {
  // itemIndexMethod()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene15itemIndexMethodEv
    // invoke: QGraphicsScene::ItemIndexMethod itemIndexMethod()
    C.C_ZNK14QGraphicsScene15itemIndexMethodEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "itemIndexMethod", args)
  }

  return
}

// addLine(const class QLineF &, const class QPen &)
func (this *QGraphicsScene) Addline(args ...interface{}) () {
  // addLine(const class QLineF &, const class QPen &)
  // addLine(qreal, qreal, qreal, qreal, const class QPen &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QLineF{}) // "const QLineF &"
  vtys[0][1] = reflect.TypeOf(qtgui.QPen{}) // "const QPen &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][4] = reflect.TypeOf(qtgui.QPen{}) // "const QPen &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7addLineERK6QLineFRK4QPen
    // invoke: QGraphicsLineItem * addLine(const class QLineF &, const class QPen &)
    var arg0 = args[0].(*qtcore.QLineF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtgui.QPen).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN14QGraphicsScene7addLineERK6QLineFRK4QPen(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN14QGraphicsScene7addLineEddddRK4QPen
    // invoke: QGraphicsLineItem * addLine(qreal, qreal, qreal, qreal, const class QPen &)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(*qtgui.QPen).Qclsinst
    if false {fmt.Println(arg4)}
    C.C_ZN14QGraphicsScene7addLineEddddRK4QPen(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addLine", args)
  }

  return
}

// minimumRenderSize()
func (this *QGraphicsScene) Minimumrendersize(args ...interface{}) (ret interface{}) {
  // minimumRenderSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene17minimumRenderSizeEv
    // invoke: qreal minimumRenderSize()
    var ret0 = C.C_ZNK14QGraphicsScene17minimumRenderSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsScene", "minimumRenderSize", args)
  }

  return
}

// sendEvent(class QGraphicsItem *, class QEvent *)
func (this *QGraphicsScene) Sendevent(args ...interface{}) (ret interface{}) {
  // sendEvent(class QGraphicsItem *, class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  vtys[0][1] = reflect.TypeOf(qtcore.QEvent{}) // "QEvent *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene9sendEventEP13QGraphicsItemP6QEvent
    // invoke: bool sendEvent(class QGraphicsItem *, class QEvent *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QEvent).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN14QGraphicsScene9sendEventEP13QGraphicsItemP6QEvent(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsScene", "sendEvent", args)
  }

  return
}

// setPalette(const class QPalette &)
func (this *QGraphicsScene) Setpalette(args ...interface{}) () {
  // setPalette(const class QPalette &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QPalette{}) // "const QPalette &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene10setPaletteERK8QPalette
    // invoke: void setPalette(const class QPalette &)
    var arg0 = args[0].(*qtgui.QPalette).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QGraphicsScene10setPaletteERK8QPalette(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setPalette", args)
  }

  return
}

// mouseGrabberItem()
func (this *QGraphicsScene) Mousegrabberitem(args ...interface{}) () {
  // mouseGrabberItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene16mouseGrabberItemEv
    // invoke: QGraphicsItem * mouseGrabberItem()
    C.C_ZNK14QGraphicsScene16mouseGrabberItemEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "mouseGrabberItem", args)
  }

  return
}

// setSelectionArea(const class QPainterPath &, const class QTransform &)
func (this *QGraphicsScene) Setselectionarea(args ...interface{}) () {
  // setSelectionArea(const class QPainterPath &, const class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QPainterPath{}) // "const QPainterPath &"
  vtys[0][1] = reflect.TypeOf(qtgui.QTransform{}) // "const QTransform &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathRK10QTransform
    // invoke: void setSelectionArea(const class QPainterPath &, const class QTransform &)
    var arg0 = args[0].(*qtgui.QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtgui.QTransform).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathRK10QTransform(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setSelectionArea", args)
  }

  return
}

// width()
func (this *QGraphicsScene) Width(args ...interface{}) (ret interface{}) {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene5widthEv
    // invoke: qreal width()
    var ret0 = C.C_ZNK14QGraphicsScene5widthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsScene", "width", args)
  }

  return
}

// foregroundBrush()
func (this *QGraphicsScene) Foregroundbrush(args ...interface{}) (ret interface{}) {
  // foregroundBrush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene15foregroundBrushEv
    // invoke: QBrush foregroundBrush()
    var ret0 = C.C_ZNK14QGraphicsScene15foregroundBrushEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QBrush{}) // "QBrush"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsScene", "foregroundBrush", args)
  }

  return
}

// setForegroundBrush(const class QBrush &)
func (this *QGraphicsScene) Setforegroundbrush(args ...interface{}) () {
  // setForegroundBrush(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QBrush{}) // "const QBrush &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene18setForegroundBrushERK6QBrush
    // invoke: void setForegroundBrush(const class QBrush &)
    var arg0 = args[0].(*qtgui.QBrush).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QGraphicsScene18setForegroundBrushERK6QBrush(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setForegroundBrush", args)
  }

  return
}

// addPath(const class QPainterPath &, const class QPen &, const class QBrush &)
func (this *QGraphicsScene) Addpath(args ...interface{}) () {
  // addPath(const class QPainterPath &, const class QPen &, const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QPainterPath{}) // "const QPainterPath &"
  vtys[0][1] = reflect.TypeOf(qtgui.QPen{}) // "const QPen &"
  vtys[0][2] = reflect.TypeOf(qtgui.QBrush{}) // "const QBrush &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7addPathERK12QPainterPathRK4QPenRK6QBrush
    // invoke: QGraphicsPathItem * addPath(const class QPainterPath &, const class QPen &, const class QBrush &)
    var arg0 = args[0].(*qtgui.QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtgui.QPen).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtgui.QBrush).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN14QGraphicsScene7addPathERK12QPainterPathRK4QPenRK6QBrush(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addPath", args)
  }

  return
}

// selectedItems()
func (this *QGraphicsScene) Selecteditems(args ...interface{}) () {
  // selectedItems()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene13selectedItemsEv
    // invoke: QList<QGraphicsItem *> selectedItems()
    C.C_ZNK14QGraphicsScene13selectedItemsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "selectedItems", args)
  }

  return
}

// setActiveWindow(class QGraphicsWidget *)
func (this *QGraphicsScene) Setactivewindow(args ...interface{}) () {
  // setActiveWindow(class QGraphicsWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsWidget{}) // "QGraphicsWidget *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene15setActiveWindowEP15QGraphicsWidget
    // invoke: void setActiveWindow(class QGraphicsWidget *)
    var arg0 = args[0].(*QGraphicsWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QGraphicsScene15setActiveWindowEP15QGraphicsWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setActiveWindow", args)
  }

  return
}

// backgroundBrush()
func (this *QGraphicsScene) Backgroundbrush(args ...interface{}) (ret interface{}) {
  // backgroundBrush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene15backgroundBrushEv
    // invoke: QBrush backgroundBrush()
    var ret0 = C.C_ZNK14QGraphicsScene15backgroundBrushEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QBrush{}) // "QBrush"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsScene", "backgroundBrush", args)
  }

  return
}

// addSimpleText(const class QString &, const class QFont &)
func (this *QGraphicsScene) Addsimpletext(args ...interface{}) () {
  // addSimpleText(const class QString &, const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(qtgui.QFont{}) // "const QFont &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene13addSimpleTextERK7QStringRK5QFont
    // invoke: QGraphicsSimpleTextItem * addSimpleText(const class QString &, const class QFont &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtgui.QFont).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN14QGraphicsScene13addSimpleTextERK7QStringRK5QFont(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addSimpleText", args)
  }

  return
}

// clearSelection()
func (this *QGraphicsScene) Clearselection(args ...interface{}) () {
  // clearSelection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene14clearSelectionEv
    // invoke: void clearSelection()
    C.C_ZN14QGraphicsScene14clearSelectionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "clearSelection", args)
  }

  return
}

// activePanel()
func (this *QGraphicsScene) Activepanel(args ...interface{}) () {
  // activePanel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene11activePanelEv
    // invoke: QGraphicsItem * activePanel()
    C.C_ZNK14QGraphicsScene11activePanelEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "activePanel", args)
  }

  return
}

// focusItem()
func (this *QGraphicsScene) Focusitem(args ...interface{}) () {
  // focusItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene9focusItemEv
    // invoke: QGraphicsItem * focusItem()
    C.C_ZNK14QGraphicsScene9focusItemEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "focusItem", args)
  }

  return
}

// destroyItemGroup(class QGraphicsItemGroup *)
func (this *QGraphicsScene) Destroyitemgroup(args ...interface{}) () {
  // destroyItemGroup(class QGraphicsItemGroup *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItemGroup{}) // "QGraphicsItemGroup *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene16destroyItemGroupEP18QGraphicsItemGroup
    // invoke: void destroyItemGroup(class QGraphicsItemGroup *)
    var arg0 = args[0].(*QGraphicsItemGroup).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QGraphicsScene16destroyItemGroupEP18QGraphicsItemGroup(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "destroyItemGroup", args)
  }

  return
}

// addText(const class QString &, const class QFont &)
func (this *QGraphicsScene) Addtext(args ...interface{}) () {
  // addText(const class QString &, const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(qtgui.QFont{}) // "const QFont &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7addTextERK7QStringRK5QFont
    // invoke: QGraphicsTextItem * addText(const class QString &, const class QFont &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtgui.QFont).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN14QGraphicsScene7addTextERK7QStringRK5QFont(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addText", args)
  }

  return
}

// update(const class QRectF &)
func (this *QGraphicsScene) Update(args ...interface{}) () {
  // update(const class QRectF &)
  // update(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene6updateERK6QRectF
    // invoke: void update(const class QRectF &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QGraphicsScene6updateERK6QRectF(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN14QGraphicsScene6updateEdddd
    // invoke: void update(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN14QGraphicsScene6updateEdddd(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "update", args)
  }

  return
}

// setSortCacheEnabled(_Bool)
func (this *QGraphicsScene) Setsortcacheenabled(args ...interface{}) () {
  // setSortCacheEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene19setSortCacheEnabledEb
    // invoke: void setSortCacheEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN14QGraphicsScene19setSortCacheEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setSortCacheEnabled", args)
  }

  return
}

// setStickyFocus(_Bool)
func (this *QGraphicsScene) Setstickyfocus(args ...interface{}) () {
  // setStickyFocus(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene14setStickyFocusEb
    // invoke: void setStickyFocus(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN14QGraphicsScene14setStickyFocusEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setStickyFocus", args)
  }

  return
}

// clear()
func (this *QGraphicsScene) Clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene5clearEv
    // invoke: void clear()
    C.C_ZN14QGraphicsScene5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "clear", args)
  }

  return
}

// height()
func (this *QGraphicsScene) Height(args ...interface{}) (ret interface{}) {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene6heightEv
    // invoke: qreal height()
    var ret0 = C.C_ZNK14QGraphicsScene6heightEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsScene", "height", args)
  }

  return
}

// setFont(const class QFont &)
func (this *QGraphicsScene) Setfont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QFont{}) // "const QFont &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7setFontERK5QFont
    // invoke: void setFont(const class QFont &)
    var arg0 = args[0].(*qtgui.QFont).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QGraphicsScene7setFontERK5QFont(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setFont", args)
  }

  return
}

// <= body block end

