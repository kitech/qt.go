package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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
extern void* C_ZNK15QGraphicsWidget12boundingRectEv(void* qthis); // 4
  // proto:  void QGraphicsWidget::setLayout(QGraphicsLayout * layout);
extern void C_ZN15QGraphicsWidget9setLayoutEP15QGraphicsLayout(void* qthis, void* arg0); // 4
  // proto:  QList<QAction *> QGraphicsWidget::actions();
extern void C_ZNK15QGraphicsWidget7actionsEv(void* qthis); // 4
  // proto:  void QGraphicsWidget::setStyle(QStyle * style);
extern void C_ZN15QGraphicsWidget8setStyleEP6QStyle(void* qthis, void* arg0); // 4
  // proto:  QPainterPath QGraphicsWidget::shape();
extern void* C_ZNK15QGraphicsWidget5shapeEv(void* qthis); // 4
  // proto:  QRectF QGraphicsWidget::windowFrameGeometry();
extern void* C_ZNK15QGraphicsWidget19windowFrameGeometryEv(void* qthis); // 4
  // proto:  QGraphicsWidget * QGraphicsWidget::focusWidget();
extern void C_ZNK15QGraphicsWidget11focusWidgetEv(void* qthis); // 4
  // proto:  QFont QGraphicsWidget::font();
extern void* C_ZNK15QGraphicsWidget4fontEv(void* qthis); // 4
  // proto:  QStyle * QGraphicsWidget::style();
extern void* C_ZNK15QGraphicsWidget5styleEv(void* qthis); // 4
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
extern bool C_ZN15QGraphicsWidget5closeEv(void* qthis); // 4
  // proto:  QPalette QGraphicsWidget::palette();
extern void* C_ZNK15QGraphicsWidget7paletteEv(void* qthis); // 4
  // proto:  QString QGraphicsWidget::windowTitle();
extern void* C_ZNK15QGraphicsWidget11windowTitleEv(void* qthis); // 4
  // proto:  Qt::WindowFlags QGraphicsWidget::windowFlags();
extern void C_ZNK15QGraphicsWidget11windowFlagsEv(void* qthis); // 4
  // proto:  void QGraphicsWidget::setContentsMargins(qreal left, qreal top, qreal right, qreal bottom);
extern void C_ZN15QGraphicsWidget18setContentsMarginsEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 4
  // proto:  void QGraphicsWidget::adjustSize();
extern void C_ZN15QGraphicsWidget10adjustSizeEv(void* qthis); // 4
  // proto:  QRectF QGraphicsWidget::windowFrameRect();
extern void* C_ZNK15QGraphicsWidget15windowFrameRectEv(void* qthis); // 4
  // proto:  void QGraphicsWidget::getWindowFrameMargins(qreal * left, qreal * top, qreal * right, qreal * bottom);
extern void C_ZNK15QGraphicsWidget21getWindowFrameMarginsEPdS0_S0_S0_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3); // 4
  // proto:  void QGraphicsWidget::addAction(QAction * action);
extern void C_ZN15QGraphicsWidget9addActionEP7QAction(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsWidget::setAutoFillBackground(bool enabled);
extern void C_ZN15QGraphicsWidget21setAutoFillBackgroundEb(void* qthis, bool arg0); // 4
  // proto:  void QGraphicsWidget::setPalette(const QPalette & palette);
extern void C_ZN15QGraphicsWidget10setPaletteERK8QPalette(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsWidget::~QGraphicsWidget();
extern void C_ZN15QGraphicsWidgetD2Ev(void* qthis); // 4
  // proto:  bool QGraphicsWidget::isActiveWindow();
extern bool C_ZNK15QGraphicsWidget14isActiveWindowEv(void* qthis); // 4
  // proto:  QSizeF QGraphicsWidget::size();
extern void* C_ZNK15QGraphicsWidget4sizeEv(void* qthis); // 4
  // proto:  bool QGraphicsWidget::autoFillBackground();
extern bool C_ZNK15QGraphicsWidget18autoFillBackgroundEv(void* qthis); // 4
  // proto:  Qt::LayoutDirection QGraphicsWidget::layoutDirection();
extern void C_ZNK15QGraphicsWidget15layoutDirectionEv(void* qthis); // 4
  // proto:  void QGraphicsWidget::setShortcutEnabled(int id, bool enabled);
extern void C_ZN15QGraphicsWidget18setShortcutEnabledEib(void* qthis, int32_t arg0, bool arg1); // 4
  // proto:  void QGraphicsWidget::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void C_ZN15QGraphicsWidget5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  QRectF QGraphicsWidget::rect();
extern void* C_ZNK15QGraphicsWidget4rectEv(void* qthis); // 2
  // proto:  int QGraphicsWidget::type();
extern int32_t C_ZNK15QGraphicsWidget4typeEv(void* qthis); // 4
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
extern void C_ZNK15QGraphicsWidget18getContentsMarginsEPdS0_S0_S0_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
//  _layoutChanged QGraphicsWidget_layoutChanged_signal;
//  _geometryChanged QGraphicsWidget_geometryChanged_signal;
}

// boundingRect()
func (this *QGraphicsWidget) Boundingrect(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QGraphicsWidget12boundingRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "boundingRect", args)
  }

  return
}

// setLayout(class QGraphicsLayout *)
func (this *QGraphicsWidget) Setlayout(args ...interface{}) () {
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
    var arg0 = args[0].(*QGraphicsLayout).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QGraphicsWidget9setLayoutEP15QGraphicsLayout(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setLayout", args)
  }

  return
}

// actions()
func (this *QGraphicsWidget) Actions(args ...interface{}) () {
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
    C.C_ZNK15QGraphicsWidget7actionsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "actions", args)
  }

  return
}

// setStyle(class QStyle *)
func (this *QGraphicsWidget) Setstyle(args ...interface{}) () {
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
    var arg0 = args[0].(*QStyle).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QGraphicsWidget8setStyleEP6QStyle(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setStyle", args)
  }

  return
}

// shape()
func (this *QGraphicsWidget) Shape(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QGraphicsWidget5shapeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "shape", args)
  }

  return
}

// windowFrameGeometry()
func (this *QGraphicsWidget) Windowframegeometry(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QGraphicsWidget19windowFrameGeometryEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "windowFrameGeometry", args)
  }

  return
}

// focusWidget()
func (this *QGraphicsWidget) Focuswidget(args ...interface{}) () {
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
    C.C_ZNK15QGraphicsWidget11focusWidgetEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "focusWidget", args)
  }

  return
}

// font()
func (this *QGraphicsWidget) Font(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QGraphicsWidget4fontEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QFont{}) // "QFont"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "font", args)
  }

  return
}

// style()
func (this *QGraphicsWidget) Style(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QGraphicsWidget5styleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStyle{}) // "QStyle *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "style", args)
  }

  return
}

// layout()
func (this *QGraphicsWidget) Layout(args ...interface{}) () {
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
    C.C_ZNK15QGraphicsWidget6layoutEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "layout", args)
  }

  return
}

// unsetLayoutDirection()
func (this *QGraphicsWidget) Unsetlayoutdirection(args ...interface{}) () {
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
    C.C_ZN15QGraphicsWidget20unsetLayoutDirectionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "unsetLayoutDirection", args)
  }

  return
}

// insertAction(class QAction *, class QAction *)
func (this *QGraphicsWidget) Insertaction(args ...interface{}) () {
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
    var arg0 = args[0].(*QAction).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QAction).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN15QGraphicsWidget12insertActionEP7QActionS1_(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "insertAction", args)
  }

  return
}

// unsetWindowFrameMargins()
func (this *QGraphicsWidget) Unsetwindowframemargins(args ...interface{}) () {
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
    C.C_ZN15QGraphicsWidget23unsetWindowFrameMarginsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "unsetWindowFrameMargins", args)
  }

  return
}

// setShortcutAutoRepeat(int, _Bool)
func (this *QGraphicsWidget) Setshortcutautorepeat(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN15QGraphicsWidget21setShortcutAutoRepeatEib(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setShortcutAutoRepeat", args)
  }

  return
}

// setGeometry(const class QRectF &)
func (this *QGraphicsWidget) Setgeometry(args ...interface{}) () {
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
    var arg0 = args[0].(*QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QGraphicsWidget11setGeometryERK6QRectF(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN15QGraphicsWidget11setGeometryEdddd
    // invoke: void setGeometry(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN15QGraphicsWidget11setGeometryEdddd(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setGeometry", args)
  }

  return
}

// windowType()
func (this *QGraphicsWidget) Windowtype(args ...interface{}) () {
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
    C.C_ZNK15QGraphicsWidget10windowTypeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "windowType", args)
  }

  return
}

// removeAction(class QAction *)
func (this *QGraphicsWidget) Removeaction(args ...interface{}) () {
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
    var arg0 = args[0].(*QAction).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QGraphicsWidget12removeActionEP7QAction(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "removeAction", args)
  }

  return
}

// metaObject()
func (this *QGraphicsWidget) Metaobject(args ...interface{}) () {
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
    C.C_ZNK15QGraphicsWidget10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "metaObject", args)
  }

  return
}

// focusPolicy()
func (this *QGraphicsWidget) Focuspolicy(args ...interface{}) () {
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
    C.C_ZNK15QGraphicsWidget11focusPolicyEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "focusPolicy", args)
  }

  return
}

// paintWindowFrame(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsWidget) Paintwindowframe(args ...interface{}) () {
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
    var arg0 = args[0].(*QPainter).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStyleOptionGraphicsItem).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QWidget).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN15QGraphicsWidget16paintWindowFrameEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "paintWindowFrame", args)
  }

  return
}

// close()
func (this *QGraphicsWidget) Close(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN15QGraphicsWidget5closeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "close", args)
  }

  return
}

// palette()
func (this *QGraphicsWidget) Palette(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QGraphicsWidget7paletteEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPalette{}) // "QPalette"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "palette", args)
  }

  return
}

// windowTitle()
func (this *QGraphicsWidget) Windowtitle(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QGraphicsWidget11windowTitleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "windowTitle", args)
  }

  return
}

// windowFlags()
func (this *QGraphicsWidget) Windowflags(args ...interface{}) () {
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
    C.C_ZNK15QGraphicsWidget11windowFlagsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "windowFlags", args)
  }

  return
}

// setContentsMargins(qreal, qreal, qreal, qreal)
func (this *QGraphicsWidget) Setcontentsmargins(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN15QGraphicsWidget18setContentsMarginsEdddd(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setContentsMargins", args)
  }

  return
}

// adjustSize()
func (this *QGraphicsWidget) Adjustsize(args ...interface{}) () {
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
    C.C_ZN15QGraphicsWidget10adjustSizeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "adjustSize", args)
  }

  return
}

// windowFrameRect()
func (this *QGraphicsWidget) Windowframerect(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QGraphicsWidget15windowFrameRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "windowFrameRect", args)
  }

  return
}

// getWindowFrameMargins(qreal *, qreal *, qreal *, qreal *)
func (this *QGraphicsWidget) Getwindowframemargins(args ...interface{}) () {
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
    var arg0 = (unsafe.Pointer)(args[0].(*float64))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*float64))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*float64))
    if false {fmt.Println(arg3)}
    C.C_ZNK15QGraphicsWidget21getWindowFrameMarginsEPdS0_S0_S0_(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "getWindowFrameMargins", args)
  }

  return
}

// addAction(class QAction *)
func (this *QGraphicsWidget) Addaction(args ...interface{}) () {
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
    var arg0 = args[0].(*QAction).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QGraphicsWidget9addActionEP7QAction(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "addAction", args)
  }

  return
}

// setAutoFillBackground(_Bool)
func (this *QGraphicsWidget) Setautofillbackground(args ...interface{}) () {
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
    C.C_ZN15QGraphicsWidget21setAutoFillBackgroundEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setAutoFillBackground", args)
  }

  return
}

// setPalette(const class QPalette &)
func (this *QGraphicsWidget) Setpalette(args ...interface{}) () {
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
    var arg0 = args[0].(*QPalette).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QGraphicsWidget10setPaletteERK8QPalette(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setPalette", args)
  }

  return
}

// ~QGraphicsWidget()
func (this *QGraphicsWidget) Freeqgraphicswidget(args ...interface{}) () {
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
    C.C_ZN15QGraphicsWidgetD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "~QGraphicsWidget", args)
  }

  return
}

// isActiveWindow()
func (this *QGraphicsWidget) Isactivewindow(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QGraphicsWidget14isActiveWindowEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "isActiveWindow", args)
  }

  return
}

// size()
func (this *QGraphicsWidget) Size(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QGraphicsWidget4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSizeF{}) // "QSizeF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "size", args)
  }

  return
}

// autoFillBackground()
func (this *QGraphicsWidget) Autofillbackground(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QGraphicsWidget18autoFillBackgroundEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "autoFillBackground", args)
  }

  return
}

// layoutDirection()
func (this *QGraphicsWidget) Layoutdirection(args ...interface{}) () {
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
    C.C_ZNK15QGraphicsWidget15layoutDirectionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "layoutDirection", args)
  }

  return
}

// setShortcutEnabled(int, _Bool)
func (this *QGraphicsWidget) Setshortcutenabled(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN15QGraphicsWidget18setShortcutEnabledEib(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setShortcutEnabled", args)
  }

  return
}

// paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsWidget) Paint(args ...interface{}) () {
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
    var arg0 = args[0].(*QPainter).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStyleOptionGraphicsItem).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QWidget).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN15QGraphicsWidget5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "paint", args)
  }

  return
}

// rect()
func (this *QGraphicsWidget) Rect(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QGraphicsWidget4rectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "rect", args)
  }

  return
}

// type()
func (this *QGraphicsWidget) Type_(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QGraphicsWidget4typeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "type", args)
  }

  return
}

// resize(qreal, qreal)
func (this *QGraphicsWidget) Resize(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN15QGraphicsWidget6resizeEdd(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN15QGraphicsWidget6resizeERK6QSizeF
    // invoke: void resize(const class QSizeF &)
    var arg0 = args[0].(*QSizeF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QGraphicsWidget6resizeERK6QSizeF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "resize", args)
  }

  return
}

// releaseShortcut(int)
func (this *QGraphicsWidget) Releaseshortcut(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QGraphicsWidget15releaseShortcutEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "releaseShortcut", args)
  }

  return
}

// setTabOrder(class QGraphicsWidget *, class QGraphicsWidget *)
func (this *QGraphicsWidget) Settaborder_S(args ...interface{}) () {
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
    var arg0 = args[0].(*QGraphicsWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QGraphicsWidget).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN15QGraphicsWidget11setTabOrderEPS_S0_(arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setTabOrder", args)
  }

  return
}

// setWindowTitle(const class QString &)
func (this *QGraphicsWidget) Setwindowtitle(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QGraphicsWidget14setWindowTitleERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setWindowTitle", args)
  }

  return
}

// setWindowFrameMargins(qreal, qreal, qreal, qreal)
func (this *QGraphicsWidget) Setwindowframemargins(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN15QGraphicsWidget21setWindowFrameMarginsEdddd(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setWindowFrameMargins", args)
  }

  return
}

// getContentsMargins(qreal *, qreal *, qreal *, qreal *)
func (this *QGraphicsWidget) Getcontentsmargins(args ...interface{}) () {
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
    var arg0 = (unsafe.Pointer)(args[0].(*float64))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*float64))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*float64))
    if false {fmt.Println(arg3)}
    C.C_ZNK15QGraphicsWidget18getContentsMarginsEPdS0_S0_S0_(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "getContentsMargins", args)
  }

  return
}

// setFont(const class QFont &)
func (this *QGraphicsWidget) Setfont(args ...interface{}) () {
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
    var arg0 = args[0].(*QFont).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QGraphicsWidget7setFontERK5QFont(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setFont", args)
  }

  return
}

// <= body block end

