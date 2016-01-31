package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtWidgets/qgraphicswidget.h
// dst-file: /src/widgets/qgraphicswidget.go
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
  // proto:  QRectF QGraphicsWidget::boundingRect();
extern void C_ZNK15QGraphicsWidget12boundingRectEv(void* qthis); // 4
  // proto:  void QGraphicsWidget::setLayout(QGraphicsLayout * layout);
extern void C_ZN15QGraphicsWidget9setLayoutEP15QGraphicsLayout(void* qthis, void* arg0); // 4
  // proto:  QList<QAction *> QGraphicsWidget::actions();
extern void C_ZNK15QGraphicsWidget7actionsEv(void* qthis); // 4
  // proto:  void QGraphicsWidget::setStyle(QStyle * style);
extern void C_ZN15QGraphicsWidget8setStyleEP6QStyle(void* qthis, void* arg0); // 4
  // proto:  QPainterPath QGraphicsWidget::shape();
extern void C_ZNK15QGraphicsWidget5shapeEv(void* qthis); // 4
  // proto:  QRectF QGraphicsWidget::windowFrameGeometry();
extern void C_ZNK15QGraphicsWidget19windowFrameGeometryEv(void* qthis); // 4
  // proto:  QGraphicsWidget * QGraphicsWidget::focusWidget();
extern void C_ZNK15QGraphicsWidget11focusWidgetEv(void* qthis); // 4
  // proto:  QFont QGraphicsWidget::font();
extern void C_ZNK15QGraphicsWidget4fontEv(void* qthis); // 4
  // proto:  QStyle * QGraphicsWidget::style();
extern void C_ZNK15QGraphicsWidget5styleEv(void* qthis); // 4
  // proto:  QGraphicsLayout * QGraphicsWidget::layout();
extern void C_ZNK15QGraphicsWidget6layoutEv(void* qthis); // 4
  // proto:  void QGraphicsWidget::unsetLayoutDirection();
extern void C_ZN15QGraphicsWidget20unsetLayoutDirectionEv(void* qthis); // 4
  // proto:  void QGraphicsWidget::insertAction(QAction * before, QAction * action);
extern void C_ZN15QGraphicsWidget12insertActionEP7QActionS1_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QGraphicsWidget::unsetWindowFrameMargins();
extern void C_ZN15QGraphicsWidget23unsetWindowFrameMarginsEv(void* qthis); // 4
  // proto:  void QGraphicsWidget::setShortcutAutoRepeat(int id, bool enabled);
extern void C_ZN15QGraphicsWidget21setShortcutAutoRepeatEib(void* qthis, int32_t arg0, bool arg1); // 4
  // proto:  void QGraphicsWidget::setGeometry(const QRectF & rect);
extern void C_ZN15QGraphicsWidget11setGeometryERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsWidget::setGeometry(qreal x, qreal y, qreal w, qreal h);
extern void C_ZN15QGraphicsWidget11setGeometryEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  Qt::WindowType QGraphicsWidget::windowType();
extern void C_ZNK15QGraphicsWidget10windowTypeEv(void* qthis); // 4
  // proto:  void QGraphicsWidget::removeAction(QAction * action);
extern void C_ZN15QGraphicsWidget12removeActionEP7QAction(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QGraphicsWidget::metaObject();
extern void C_ZNK15QGraphicsWidget10metaObjectEv(void* qthis); // 4
  // proto:  Qt::FocusPolicy QGraphicsWidget::focusPolicy();
extern void C_ZNK15QGraphicsWidget11focusPolicyEv(void* qthis); // 4
  // proto:  void QGraphicsWidget::paintWindowFrame(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void C_ZN15QGraphicsWidget16paintWindowFrameEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  bool QGraphicsWidget::close();
extern void C_ZN15QGraphicsWidget5closeEv(void* qthis); // 4
  // proto:  QPalette QGraphicsWidget::palette();
extern void C_ZNK15QGraphicsWidget7paletteEv(void* qthis); // 4
  // proto:  QString QGraphicsWidget::windowTitle();
extern void C_ZNK15QGraphicsWidget11windowTitleEv(void* qthis); // 4
  // proto:  Qt::WindowFlags QGraphicsWidget::windowFlags();
extern void C_ZNK15QGraphicsWidget11windowFlagsEv(void* qthis); // 4
  // proto:  void QGraphicsWidget::setContentsMargins(qreal left, qreal top, qreal right, qreal bottom);
extern void C_ZN15QGraphicsWidget18setContentsMarginsEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 4
  // proto:  void QGraphicsWidget::adjustSize();
extern void C_ZN15QGraphicsWidget10adjustSizeEv(void* qthis); // 4
  // proto:  QRectF QGraphicsWidget::windowFrameRect();
extern void C_ZNK15QGraphicsWidget15windowFrameRectEv(void* qthis); // 4
  // proto:  void QGraphicsWidget::getWindowFrameMargins(qreal * left, qreal * top, qreal * right, qreal * bottom);
extern void C_ZNK15QGraphicsWidget21getWindowFrameMarginsEPdS0_S0_S0_(void* qthis, double* arg0, double* arg1, double* arg2, double* arg3); // 4
  // proto:  void QGraphicsWidget::addAction(QAction * action);
extern void C_ZN15QGraphicsWidget9addActionEP7QAction(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsWidget::setAutoFillBackground(bool enabled);
extern void C_ZN15QGraphicsWidget21setAutoFillBackgroundEb(void* qthis, bool arg0); // 4
  // proto:  void QGraphicsWidget::setPalette(const QPalette & palette);
extern void C_ZN15QGraphicsWidget10setPaletteERK8QPalette(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsWidget::~QGraphicsWidget();
extern void C_ZN15QGraphicsWidgetD2Ev(void* qthis); // 4
  // proto:  bool QGraphicsWidget::isActiveWindow();
extern void C_ZNK15QGraphicsWidget14isActiveWindowEv(void* qthis); // 4
  // proto:  QSizeF QGraphicsWidget::size();
extern void C_ZNK15QGraphicsWidget4sizeEv(void* qthis); // 4
  // proto:  bool QGraphicsWidget::autoFillBackground();
extern void C_ZNK15QGraphicsWidget18autoFillBackgroundEv(void* qthis); // 4
  // proto:  Qt::LayoutDirection QGraphicsWidget::layoutDirection();
extern void C_ZNK15QGraphicsWidget15layoutDirectionEv(void* qthis); // 4
  // proto:  void QGraphicsWidget::setShortcutEnabled(int id, bool enabled);
extern void C_ZN15QGraphicsWidget18setShortcutEnabledEib(void* qthis, int32_t arg0, bool arg1); // 4
  // proto:  void QGraphicsWidget::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void C_ZN15QGraphicsWidget5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  QRectF QGraphicsWidget::rect();
extern void C_ZNK15QGraphicsWidget4rectEv(void* qthis); // 2
  // proto:  int QGraphicsWidget::type();
extern void C_ZNK15QGraphicsWidget4typeEv(void* qthis); // 4
  // proto:  void QGraphicsWidget::resize(qreal w, qreal h);
extern void C_ZN15QGraphicsWidget6resizeEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QGraphicsWidget::resize(const QSizeF & size);
extern void C_ZN15QGraphicsWidget6resizeERK6QSizeF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsWidget::releaseShortcut(int id);
extern void C_ZN15QGraphicsWidget15releaseShortcutEi(void* qthis, int32_t arg0); // 4
  // proto: static void QGraphicsWidget::setTabOrder(QGraphicsWidget * first, QGraphicsWidget * second);
extern void C_ZN15QGraphicsWidget11setTabOrderEPS_S0_(void* arg0, void* arg1); // 4
  // proto:  void QGraphicsWidget::setWindowTitle(const QString & title);
extern void C_ZN15QGraphicsWidget14setWindowTitleERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsWidget::setWindowFrameMargins(qreal left, qreal top, qreal right, qreal bottom);
extern void C_ZN15QGraphicsWidget21setWindowFrameMarginsEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 4
  // proto:  void QGraphicsWidget::getContentsMargins(qreal * left, qreal * top, qreal * right, qreal * bottom);
extern void C_ZNK15QGraphicsWidget18getContentsMarginsEPdS0_S0_S0_(void* qthis, double* arg0, double* arg1, double* arg2, double* arg3); // 4
  // proto:  void QGraphicsWidget::setFont(const QFont & font);
extern void C_ZN15QGraphicsWidget7setFontERK5QFont(void* qthis, void* arg0); // 4
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

// class sizeof(QGraphicsWidget)=1
type QGraphicsWidget struct {
  /*qbase*/ QGraphicsObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _layoutChanged QGraphicsWidget_layoutChanged_signal;
//  _geometryChanged QGraphicsWidget_geometryChanged_signal;
}

// boundingRect()
func (this *QGraphicsWidget) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget12boundingRectEv
    // invoke: QRectF boundingRect()
    var ret = C.C_ZNK15QGraphicsWidget12boundingRectEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "boundingRect", args)
  }

}

// setLayout(class QGraphicsLayout *)
func (this *QGraphicsWidget) setLayout(args ...interface{}) () {
  // setLayout(class QGraphicsLayout *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsLayout{}) // "QGraphicsLayout *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget9setLayoutEP15QGraphicsLayout
    // invoke: void setLayout(class QGraphicsLayout *)
    var arg0 = args[0].(QGraphicsLayout).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QGraphicsWidget9setLayoutEP15QGraphicsLayout(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setLayout", args)
  }

}

// actions()
func (this *QGraphicsWidget) actions(args ...interface{}) () {
  // actions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget7actionsEv
    // invoke: QList<QAction *> actions()
    C.C_ZNK15QGraphicsWidget7actionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "actions", args)
  }

}

// setStyle(class QStyle *)
func (this *QGraphicsWidget) setStyle(args ...interface{}) () {
  // setStyle(class QStyle *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStyle{}) // "QStyle *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget8setStyleEP6QStyle
    // invoke: void setStyle(class QStyle *)
    var arg0 = args[0].(QStyle).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QGraphicsWidget8setStyleEP6QStyle(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setStyle", args)
  }

}

// shape()
func (this *QGraphicsWidget) shape(args ...interface{}) () {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget5shapeEv
    // invoke: QPainterPath shape()
    var ret = C.C_ZNK15QGraphicsWidget5shapeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "shape", args)
  }

}

// windowFrameGeometry()
func (this *QGraphicsWidget) windowFrameGeometry(args ...interface{}) () {
  // windowFrameGeometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget19windowFrameGeometryEv
    // invoke: QRectF windowFrameGeometry()
    var ret = C.C_ZNK15QGraphicsWidget19windowFrameGeometryEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "windowFrameGeometry", args)
  }

}

// focusWidget()
func (this *QGraphicsWidget) focusWidget(args ...interface{}) () {
  // focusWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget11focusWidgetEv
    // invoke: QGraphicsWidget * focusWidget()
    C.C_ZNK15QGraphicsWidget11focusWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "focusWidget", args)
  }

}

// font()
func (this *QGraphicsWidget) font(args ...interface{}) () {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget4fontEv
    // invoke: QFont font()
    var ret = C.C_ZNK15QGraphicsWidget4fontEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "font", args)
  }

}

// style()
func (this *QGraphicsWidget) style(args ...interface{}) () {
  // style()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget5styleEv
    // invoke: QStyle * style()
    var ret = C.C_ZNK15QGraphicsWidget5styleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "style", args)
  }

}

// layout()
func (this *QGraphicsWidget) layout(args ...interface{}) () {
  // layout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget6layoutEv
    // invoke: QGraphicsLayout * layout()
    C.C_ZNK15QGraphicsWidget6layoutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "layout", args)
  }

}

// unsetLayoutDirection()
func (this *QGraphicsWidget) unsetLayoutDirection(args ...interface{}) () {
  // unsetLayoutDirection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget20unsetLayoutDirectionEv
    // invoke: void unsetLayoutDirection()
    C.C_ZN15QGraphicsWidget20unsetLayoutDirectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "unsetLayoutDirection", args)
  }

}

// insertAction(class QAction *, class QAction *)
func (this *QGraphicsWidget) insertAction(args ...interface{}) () {
  // insertAction(class QAction *, class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"
  vtys[0][1] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget12insertActionEP7QActionS1_
    // invoke: void insertAction(class QAction *, class QAction *)
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QAction).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN15QGraphicsWidget12insertActionEP7QActionS1_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "insertAction", args)
  }

}

// unsetWindowFrameMargins()
func (this *QGraphicsWidget) unsetWindowFrameMargins(args ...interface{}) () {
  // unsetWindowFrameMargins()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget23unsetWindowFrameMarginsEv
    // invoke: void unsetWindowFrameMargins()
    C.C_ZN15QGraphicsWidget23unsetWindowFrameMarginsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "unsetWindowFrameMargins", args)
  }

}

// setShortcutAutoRepeat(int, _Bool)
func (this *QGraphicsWidget) setShortcutAutoRepeat(args ...interface{}) () {
  // setShortcutAutoRepeat(int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget21setShortcutAutoRepeatEib
    // invoke: void setShortcutAutoRepeat(int, _Bool)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN15QGraphicsWidget21setShortcutAutoRepeatEib(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setShortcutAutoRepeat", args)
  }

}

// setGeometry(const class QRectF &)
func (this *QGraphicsWidget) setGeometry(args ...interface{}) () {
  // setGeometry(const class QRectF &)
  // setGeometry(qreal, qreal, qreal, qreal)
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
    // invoke: _ZN15QGraphicsWidget11setGeometryERK6QRectF
    // invoke: void setGeometry(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QGraphicsWidget11setGeometryERK6QRectF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN15QGraphicsWidget11setGeometryEdddd
    // invoke: void setGeometry(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN15QGraphicsWidget11setGeometryEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setGeometry", args)
  }

}

// windowType()
func (this *QGraphicsWidget) windowType(args ...interface{}) () {
  // windowType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget10windowTypeEv
    // invoke: Qt::WindowType windowType()
    C.C_ZNK15QGraphicsWidget10windowTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "windowType", args)
  }

}

// removeAction(class QAction *)
func (this *QGraphicsWidget) removeAction(args ...interface{}) () {
  // removeAction(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget12removeActionEP7QAction
    // invoke: void removeAction(class QAction *)
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QGraphicsWidget12removeActionEP7QAction(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "removeAction", args)
  }

}

// metaObject()
func (this *QGraphicsWidget) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK15QGraphicsWidget10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "metaObject", args)
  }

}

// focusPolicy()
func (this *QGraphicsWidget) focusPolicy(args ...interface{}) () {
  // focusPolicy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget11focusPolicyEv
    // invoke: Qt::FocusPolicy focusPolicy()
    C.C_ZNK15QGraphicsWidget11focusPolicyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "focusPolicy", args)
  }

}

// paintWindowFrame(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsWidget) paintWindowFrame(args ...interface{}) () {
  // paintWindowFrame(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionGraphicsItem{}) // "const QStyleOptionGraphicsItem *"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget16paintWindowFrameEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget
    // invoke: void paintWindowFrame(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
    var arg0 = args[0].(QPainter).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStyleOptionGraphicsItem).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QWidget).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN15QGraphicsWidget16paintWindowFrameEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "paintWindowFrame", args)
  }

}

// close()
func (this *QGraphicsWidget) close(args ...interface{}) () {
  // close()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget5closeEv
    // invoke: bool close()
    var ret = C.C_ZN15QGraphicsWidget5closeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "close", args)
  }

}

// palette()
func (this *QGraphicsWidget) palette(args ...interface{}) () {
  // palette()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget7paletteEv
    // invoke: QPalette palette()
    var ret = C.C_ZNK15QGraphicsWidget7paletteEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "palette", args)
  }

}

// windowTitle()
func (this *QGraphicsWidget) windowTitle(args ...interface{}) () {
  // windowTitle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget11windowTitleEv
    // invoke: QString windowTitle()
    var ret = C.C_ZNK15QGraphicsWidget11windowTitleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "windowTitle", args)
  }

}

// windowFlags()
func (this *QGraphicsWidget) windowFlags(args ...interface{}) () {
  // windowFlags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget11windowFlagsEv
    // invoke: Qt::WindowFlags windowFlags()
    C.C_ZNK15QGraphicsWidget11windowFlagsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "windowFlags", args)
  }

}

// setContentsMargins(qreal, qreal, qreal, qreal)
func (this *QGraphicsWidget) setContentsMargins(args ...interface{}) () {
  // setContentsMargins(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget18setContentsMarginsEdddd
    // invoke: void setContentsMargins(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN15QGraphicsWidget18setContentsMarginsEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setContentsMargins", args)
  }

}

// adjustSize()
func (this *QGraphicsWidget) adjustSize(args ...interface{}) () {
  // adjustSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget10adjustSizeEv
    // invoke: void adjustSize()
    C.C_ZN15QGraphicsWidget10adjustSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "adjustSize", args)
  }

}

// windowFrameRect()
func (this *QGraphicsWidget) windowFrameRect(args ...interface{}) () {
  // windowFrameRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget15windowFrameRectEv
    // invoke: QRectF windowFrameRect()
    var ret = C.C_ZNK15QGraphicsWidget15windowFrameRectEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "windowFrameRect", args)
  }

}

// getWindowFrameMargins(qreal *, qreal *, qreal *, qreal *)
func (this *QGraphicsWidget) getWindowFrameMargins(args ...interface{}) () {
  // getWindowFrameMargins(qreal *, qreal *, qreal *, qreal *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][1] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][2] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][3] = qtrt.DoubleTy(true) // "qreal *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget21getWindowFrameMarginsEPdS0_S0_S0_
    // invoke: void getWindowFrameMargins(qreal *, qreal *, qreal *, qreal *)
    var arg0 = (*C.double)(args[0].(*float64))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.double)(args[1].(*float64))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.double)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.double)(args[3].(*float64))
    if false {fmt.Println(arg3)}
    C.C_ZNK15QGraphicsWidget21getWindowFrameMarginsEPdS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "getWindowFrameMargins", args)
  }

}

// addAction(class QAction *)
func (this *QGraphicsWidget) addAction(args ...interface{}) () {
  // addAction(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget9addActionEP7QAction
    // invoke: void addAction(class QAction *)
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QGraphicsWidget9addActionEP7QAction(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "addAction", args)
  }

}

// setAutoFillBackground(_Bool)
func (this *QGraphicsWidget) setAutoFillBackground(args ...interface{}) () {
  // setAutoFillBackground(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget21setAutoFillBackgroundEb
    // invoke: void setAutoFillBackground(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QGraphicsWidget21setAutoFillBackgroundEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setAutoFillBackground", args)
  }

}

// setPalette(const class QPalette &)
func (this *QGraphicsWidget) setPalette(args ...interface{}) () {
  // setPalette(const class QPalette &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPalette{}) // "const QPalette &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget10setPaletteERK8QPalette
    // invoke: void setPalette(const class QPalette &)
    var arg0 = args[0].(QPalette).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QGraphicsWidget10setPaletteERK8QPalette(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setPalette", args)
  }

}

// ~QGraphicsWidget()
func (this *QGraphicsWidget) FreeQGraphicsWidget(args ...interface{}) () {
  // ~QGraphicsWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidgetD0Ev
    // invoke: void ~QGraphicsWidget()
    C.C_ZN15QGraphicsWidgetD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "~QGraphicsWidget", args)
  }

}

// isActiveWindow()
func (this *QGraphicsWidget) isActiveWindow(args ...interface{}) () {
  // isActiveWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget14isActiveWindowEv
    // invoke: bool isActiveWindow()
    var ret = C.C_ZNK15QGraphicsWidget14isActiveWindowEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "isActiveWindow", args)
  }

}

// size()
func (this *QGraphicsWidget) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget4sizeEv
    // invoke: QSizeF size()
    var ret = C.C_ZNK15QGraphicsWidget4sizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "size", args)
  }

}

// autoFillBackground()
func (this *QGraphicsWidget) autoFillBackground(args ...interface{}) () {
  // autoFillBackground()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget18autoFillBackgroundEv
    // invoke: bool autoFillBackground()
    var ret = C.C_ZNK15QGraphicsWidget18autoFillBackgroundEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "autoFillBackground", args)
  }

}

// layoutDirection()
func (this *QGraphicsWidget) layoutDirection(args ...interface{}) () {
  // layoutDirection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget15layoutDirectionEv
    // invoke: Qt::LayoutDirection layoutDirection()
    C.C_ZNK15QGraphicsWidget15layoutDirectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "layoutDirection", args)
  }

}

// setShortcutEnabled(int, _Bool)
func (this *QGraphicsWidget) setShortcutEnabled(args ...interface{}) () {
  // setShortcutEnabled(int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget18setShortcutEnabledEib
    // invoke: void setShortcutEnabled(int, _Bool)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN15QGraphicsWidget18setShortcutEnabledEib(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setShortcutEnabled", args)
  }

}

// paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsWidget) paint(args ...interface{}) () {
  // paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionGraphicsItem{}) // "const QStyleOptionGraphicsItem *"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget
    // invoke: void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
    var arg0 = args[0].(QPainter).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStyleOptionGraphicsItem).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QWidget).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN15QGraphicsWidget5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "paint", args)
  }

}

// rect()
func (this *QGraphicsWidget) rect(args ...interface{}) () {
  // rect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget4rectEv
    // invoke: QRectF rect()
    var ret = C.C_ZNK15QGraphicsWidget4rectEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "rect", args)
  }

}

// type()
func (this *QGraphicsWidget) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget4typeEv
    // invoke: int type()
    var ret = C.C_ZNK15QGraphicsWidget4typeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "type", args)
  }

}

// resize(qreal, qreal)
func (this *QGraphicsWidget) resize(args ...interface{}) () {
  // resize(qreal, qreal)
  // resize(const class QSizeF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget6resizeEdd
    // invoke: void resize(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN15QGraphicsWidget6resizeEdd(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN15QGraphicsWidget6resizeERK6QSizeF
    // invoke: void resize(const class QSizeF &)
    var arg0 = args[0].(QSizeF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QGraphicsWidget6resizeERK6QSizeF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "resize", args)
  }

}

// releaseShortcut(int)
func (this *QGraphicsWidget) releaseShortcut(args ...interface{}) () {
  // releaseShortcut(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget15releaseShortcutEi
    // invoke: void releaseShortcut(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QGraphicsWidget15releaseShortcutEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "releaseShortcut", args)
  }

}

// setTabOrder(class QGraphicsWidget *, class QGraphicsWidget *)
func (this *QGraphicsWidget) setTabOrder_s(args ...interface{}) () {
  // setTabOrder(class QGraphicsWidget *, class QGraphicsWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsWidget{}) // "QGraphicsWidget *"
  vtys[0][1] = reflect.TypeOf(QGraphicsWidget{}) // "QGraphicsWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget11setTabOrderEPS_S0_
    // invoke: void setTabOrder(class QGraphicsWidget *, class QGraphicsWidget *)
    var arg0 = args[0].(QGraphicsWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QGraphicsWidget).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN15QGraphicsWidget11setTabOrderEPS_S0_(arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setTabOrder", args)
  }

}

// setWindowTitle(const class QString &)
func (this *QGraphicsWidget) setWindowTitle(args ...interface{}) () {
  // setWindowTitle(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget14setWindowTitleERK7QString
    // invoke: void setWindowTitle(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QGraphicsWidget14setWindowTitleERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setWindowTitle", args)
  }

}

// setWindowFrameMargins(qreal, qreal, qreal, qreal)
func (this *QGraphicsWidget) setWindowFrameMargins(args ...interface{}) () {
  // setWindowFrameMargins(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget21setWindowFrameMarginsEdddd
    // invoke: void setWindowFrameMargins(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN15QGraphicsWidget21setWindowFrameMarginsEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setWindowFrameMargins", args)
  }

}

// getContentsMargins(qreal *, qreal *, qreal *, qreal *)
func (this *QGraphicsWidget) getContentsMargins(args ...interface{}) () {
  // getContentsMargins(qreal *, qreal *, qreal *, qreal *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][1] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][2] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][3] = qtrt.DoubleTy(true) // "qreal *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget18getContentsMarginsEPdS0_S0_S0_
    // invoke: void getContentsMargins(qreal *, qreal *, qreal *, qreal *)
    var arg0 = (*C.double)(args[0].(*float64))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.double)(args[1].(*float64))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.double)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.double)(args[3].(*float64))
    if false {fmt.Println(arg3)}
    C.C_ZNK15QGraphicsWidget18getContentsMarginsEPdS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "getContentsMargins", args)
  }

}

// setFont(const class QFont &)
func (this *QGraphicsWidget) setFont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget7setFontERK5QFont
    // invoke: void setFont(const class QFont &)
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QGraphicsWidget7setFontERK5QFont(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setFont", args)
  }

}

// <= body block end

