package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtWidgets/qtoolbar.h
// dst-file: /src/widgets/qtoolbar.go
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
  // proto:  QAction * QToolBar::addAction(const QIcon & icon, const QString & text, const QObject * receiver, const char * member);
extern void _ZN8QToolBar9addActionERK5QIconRK7QStringPK7QObjectPKc(void* qthis, void* arg0, void* arg1, void* arg2, unsigned char* arg3); // 4
  // proto:  QAction * QToolBar::addAction(const QString & text, const QObject * receiver, const char * member);
extern void _ZN8QToolBar9addActionERK7QStringPK7QObjectPKc(void* qthis, void* arg0, void* arg1, unsigned char* arg2); // 4
  // proto:  QAction * QToolBar::addAction(const QIcon & icon, const QString & text);
extern void _ZN8QToolBar9addActionERK5QIconRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QAction * QToolBar::addAction(const QString & text);
extern void _ZN8QToolBar9addActionERK7QString(void* qthis, void* arg0); // 4
  // proto:  Qt::Orientation QToolBar::orientation();
extern void _ZNK8QToolBar11orientationEv(void* qthis); // 4
  // proto:  QAction * QToolBar::actionAt(int x, int y);
extern void _ZNK8QToolBar8actionAtEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QAction * QToolBar::actionAt(const QPoint & p);
extern void _ZNK8QToolBar8actionAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  bool QToolBar::isFloating();
extern void _ZNK8QToolBar10isFloatingEv(void* qthis); // 4
  // proto:  QRect QToolBar::actionGeometry(QAction * action);
extern void _ZNK8QToolBar14actionGeometryEP7QAction(void* qthis, void* arg0); // 4
  // proto:  bool QToolBar::isFloatable();
extern void _ZNK8QToolBar11isFloatableEv(void* qthis); // 4
  // proto:  QAction * QToolBar::insertSeparator(QAction * before);
extern void _ZN8QToolBar15insertSeparatorEP7QAction(void* qthis, void* arg0); // 4
  // proto:  QAction * QToolBar::toggleViewAction();
extern void _ZNK8QToolBar16toggleViewActionEv(void* qthis); // 4
  // proto:  void QToolBar::setMovable(bool movable);
extern void _ZN8QToolBar10setMovableEb(void* qthis, bool arg0); // 4
  // proto:  bool QToolBar::isMovable();
extern void _ZNK8QToolBar9isMovableEv(void* qthis); // 4
  // proto:  void QToolBar::QToolBar(QWidget * parent);
extern void _ZN8QToolBarC2EP7QWidget(void* qthis, void* arg0); // 3
  // proto:  void QToolBar::QToolBar(const QString & title, QWidget * parent);
extern void _ZN8QToolBarC2ERK7QStringP7QWidget(void* qthis, void* arg0, void* arg1); // 3
  // proto:  QAction * QToolBar::addSeparator();
extern void _ZN8QToolBar12addSeparatorEv(void* qthis); // 4
  // proto:  QSize QToolBar::iconSize();
extern void _ZNK8QToolBar8iconSizeEv(void* qthis); // 4
  // proto:  QAction * QToolBar::insertWidget(QAction * before, QWidget * widget);
extern void _ZN8QToolBar12insertWidgetEP7QActionP7QWidget(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QToolBar::~QToolBar();
extern void _ZN8QToolBarD2Ev(void* qthis); // 4
  // proto:  QWidget * QToolBar::widgetForAction(QAction * action);
extern void _ZNK8QToolBar15widgetForActionEP7QAction(void* qthis, void* arg0); // 4
  // proto:  QAction * QToolBar::addWidget(QWidget * widget);
extern void _ZN8QToolBar9addWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QToolBar::metaObject();
extern void _ZNK8QToolBar10metaObjectEv(void* qthis); // 4
  // proto:  void QToolBar::setFloatable(bool floatable);
extern void _ZN8QToolBar12setFloatableEb(void* qthis, bool arg0); // 4
  // proto:  Qt::ToolBarAreas QToolBar::allowedAreas();
extern void _ZNK8QToolBar12allowedAreasEv(void* qthis); // 4
  // proto:  void QToolBar::clear();
extern void _ZN8QToolBar5clearEv(void* qthis); // 4
  // proto:  void QToolBar::setIconSize(const QSize & iconSize);
extern void _ZN8QToolBar11setIconSizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  Qt::ToolButtonStyle QToolBar::toolButtonStyle();
extern void _ZNK8QToolBar15toolButtonStyleEv(void* qthis); // 4
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

// class sizeof(QToolBar)=1
type QToolBar struct {
  /*qbase*/ QWidget;
  qclsinst unsafe.Pointer /* *C.void */;
//  _iconSizeChanged QToolBar_iconSizeChanged_signal;
//  _allowedAreasChanged QToolBar_allowedAreasChanged_signal;
//  _movableChanged QToolBar_movableChanged_signal;
//  _toolButtonStyleChanged QToolBar_toolButtonStyleChanged_signal;
//  _topLevelChanged QToolBar_topLevelChanged_signal;
//  _actionTriggered QToolBar_actionTriggered_signal;
//  _orientationChanged QToolBar_orientationChanged_signal;
//  _visibilityChanged QToolBar_visibilityChanged_signal;
}

// addAction(const class QIcon &, const class QString &, const class QObject *, const char *)
func (this *QToolBar) addAction(args ...interface{}) () {
  // addAction(const class QIcon &, const class QString &, const class QObject *, const char *)
  // addAction(const class QString &, const class QObject *, const char *)
  // addAction(const class QIcon &, const class QString &)
  // addAction(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[0][3] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[1][2] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBar9addActionERK5QIconRK7QStringPK7QObjectPKc
    // invoke: QAction * addAction(const class QIcon &, const class QString &, const class QObject *, const char *)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QObject).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[3].([]byte)).Pointer()))
    if false {fmt.Println(arg3)}
    C._ZN8QToolBar9addActionERK5QIconRK7QStringPK7QObjectPKc(this.qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN8QToolBar9addActionERK7QStringPK7QObjectPKc
    // invoke: QAction * addAction(const class QString &, const class QObject *, const char *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[2].([]byte)).Pointer()))
    if false {fmt.Println(arg2)}
    C._ZN8QToolBar9addActionERK7QStringPK7QObjectPKc(this.qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN8QToolBar9addActionERK5QIconRK7QString
    // invoke: QAction * addAction(const class QIcon &, const class QString &)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN8QToolBar9addActionERK5QIconRK7QString(this.qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZN8QToolBar9addActionERK7QString
    // invoke: QAction * addAction(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QToolBar9addActionERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolBar", "addAction", args)
  }

}

// orientation()
func (this *QToolBar) orientation(args ...interface{}) () {
  // orientation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBar11orientationEv
    // invoke: Qt::Orientation orientation()
    C._ZNK8QToolBar11orientationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolBar", "orientation", args)
  }

}

// actionAt(int, int)
func (this *QToolBar) actionAt(args ...interface{}) () {
  // actionAt(int, int)
  // actionAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBar8actionAtEii
    // invoke: QAction * actionAt(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK8QToolBar8actionAtEii(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK8QToolBar8actionAtERK6QPoint
    // invoke: QAction * actionAt(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK8QToolBar8actionAtERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolBar", "actionAt", args)
  }

}

// isFloating()
func (this *QToolBar) isFloating(args ...interface{}) () {
  // isFloating()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBar10isFloatingEv
    // invoke: bool isFloating()
    C._ZNK8QToolBar10isFloatingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolBar", "isFloating", args)
  }

}

// actionGeometry(class QAction *)
func (this *QToolBar) actionGeometry(args ...interface{}) () {
  // actionGeometry(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBar14actionGeometryEP7QAction
    // invoke: QRect actionGeometry(class QAction *)
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK8QToolBar14actionGeometryEP7QAction(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolBar", "actionGeometry", args)
  }

}

// isFloatable()
func (this *QToolBar) isFloatable(args ...interface{}) () {
  // isFloatable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBar11isFloatableEv
    // invoke: bool isFloatable()
    C._ZNK8QToolBar11isFloatableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolBar", "isFloatable", args)
  }

}

// insertSeparator(class QAction *)
func (this *QToolBar) insertSeparator(args ...interface{}) () {
  // insertSeparator(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBar15insertSeparatorEP7QAction
    // invoke: QAction * insertSeparator(class QAction *)
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QToolBar15insertSeparatorEP7QAction(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolBar", "insertSeparator", args)
  }

}

// toggleViewAction()
func (this *QToolBar) toggleViewAction(args ...interface{}) () {
  // toggleViewAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBar16toggleViewActionEv
    // invoke: QAction * toggleViewAction()
    C._ZNK8QToolBar16toggleViewActionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolBar", "toggleViewAction", args)
  }

}

// setMovable(_Bool)
func (this *QToolBar) setMovable(args ...interface{}) () {
  // setMovable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBar10setMovableEb
    // invoke: void setMovable(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN8QToolBar10setMovableEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolBar", "setMovable", args)
  }

}

// isMovable()
func (this *QToolBar) isMovable(args ...interface{}) () {
  // isMovable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBar9isMovableEv
    // invoke: bool isMovable()
    C._ZNK8QToolBar9isMovableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolBar", "isMovable", args)
  }

}

// QToolBar(class QWidget *)
func NewQToolBar(args ...interface{}) QToolBar {
  // QToolBar(class QWidget *)
  // QToolBar(const class QString &, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBarC1EP7QWidget
    // invoke: void QToolBar(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QToolBarC2EP7QWidget(qthis, arg0)
  case 1:
    // invoke: _ZN8QToolBarC1ERK7QStringP7QWidget
    // invoke: void QToolBar(const class QString &, class QWidget *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QToolBarC2ERK7QStringP7QWidget(qthis, arg0, arg1)
  default:
    qtrt.ErrorResolve("QToolBar", "QToolBar", args)
  }

  return QToolBar{}
}

// addSeparator()
func (this *QToolBar) addSeparator(args ...interface{}) () {
  // addSeparator()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBar12addSeparatorEv
    // invoke: QAction * addSeparator()
    C._ZN8QToolBar12addSeparatorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolBar", "addSeparator", args)
  }

}

// iconSize()
func (this *QToolBar) iconSize(args ...interface{}) () {
  // iconSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBar8iconSizeEv
    // invoke: QSize iconSize()
    C._ZNK8QToolBar8iconSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolBar", "iconSize", args)
  }

}

// insertWidget(class QAction *, class QWidget *)
func (this *QToolBar) insertWidget(args ...interface{}) () {
  // insertWidget(class QAction *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBar12insertWidgetEP7QActionP7QWidget
    // invoke: QAction * insertWidget(class QAction *, class QWidget *)
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN8QToolBar12insertWidgetEP7QActionP7QWidget(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QToolBar", "insertWidget", args)
  }

}

// ~QToolBar()
func (this *QToolBar) FreeQToolBar(args ...interface{}) () {
  // ~QToolBar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBarD0Ev
    // invoke: void ~QToolBar()
    C._ZN8QToolBarD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolBar", "~QToolBar", args)
  }

}

// widgetForAction(class QAction *)
func (this *QToolBar) widgetForAction(args ...interface{}) () {
  // widgetForAction(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBar15widgetForActionEP7QAction
    // invoke: QWidget * widgetForAction(class QAction *)
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK8QToolBar15widgetForActionEP7QAction(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolBar", "widgetForAction", args)
  }

}

// addWidget(class QWidget *)
func (this *QToolBar) addWidget(args ...interface{}) () {
  // addWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBar9addWidgetEP7QWidget
    // invoke: QAction * addWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QToolBar9addWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolBar", "addWidget", args)
  }

}

// metaObject()
func (this *QToolBar) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBar10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK8QToolBar10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolBar", "metaObject", args)
  }

}

// setFloatable(_Bool)
func (this *QToolBar) setFloatable(args ...interface{}) () {
  // setFloatable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBar12setFloatableEb
    // invoke: void setFloatable(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN8QToolBar12setFloatableEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolBar", "setFloatable", args)
  }

}

// allowedAreas()
func (this *QToolBar) allowedAreas(args ...interface{}) () {
  // allowedAreas()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBar12allowedAreasEv
    // invoke: Qt::ToolBarAreas allowedAreas()
    C._ZNK8QToolBar12allowedAreasEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolBar", "allowedAreas", args)
  }

}

// clear()
func (this *QToolBar) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBar5clearEv
    // invoke: void clear()
    C._ZN8QToolBar5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolBar", "clear", args)
  }

}

// setIconSize(const class QSize &)
func (this *QToolBar) setIconSize(args ...interface{}) () {
  // setIconSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBar11setIconSizeERK5QSize
    // invoke: void setIconSize(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QToolBar11setIconSizeERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolBar", "setIconSize", args)
  }

}

// toolButtonStyle()
func (this *QToolBar) toolButtonStyle(args ...interface{}) () {
  // toolButtonStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBar15toolButtonStyleEv
    // invoke: Qt::ToolButtonStyle toolButtonStyle()
    C._ZNK8QToolBar15toolButtonStyleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolBar", "toolButtonStyle", args)
  }

}

// <= body block end

