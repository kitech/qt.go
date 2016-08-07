package qtwidgets
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
  // proto:  QAction * QToolBar::addAction(const QIcon & icon, const QString & text, const QObject * receiver, const char * member);
extern void* C_ZN8QToolBar9addActionERK5QIconRK7QStringPK7QObjectPKc(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3); // 4
  // proto:  QAction * QToolBar::addAction(const QString & text, const QObject * receiver, const char * member);
extern void* C_ZN8QToolBar9addActionERK7QStringPK7QObjectPKc(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  QAction * QToolBar::addAction(const QIcon & icon, const QString & text);
extern void* C_ZN8QToolBar9addActionERK5QIconRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QAction * QToolBar::addAction(const QString & text);
extern void* C_ZN8QToolBar9addActionERK7QString(void* qthis, void* arg0); // 4
  // proto:  Qt::Orientation QToolBar::orientation();
extern void C_ZNK8QToolBar11orientationEv(void* qthis); // 4
  // proto:  QAction * QToolBar::actionAt(int x, int y);
extern void* C_ZNK8QToolBar8actionAtEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QAction * QToolBar::actionAt(const QPoint & p);
extern void* C_ZNK8QToolBar8actionAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  bool QToolBar::isFloating();
extern bool C_ZNK8QToolBar10isFloatingEv(void* qthis); // 4
  // proto:  QRect QToolBar::actionGeometry(QAction * action);
extern void* C_ZNK8QToolBar14actionGeometryEP7QAction(void* qthis, void* arg0); // 4
  // proto:  bool QToolBar::isFloatable();
extern bool C_ZNK8QToolBar11isFloatableEv(void* qthis); // 4
  // proto:  QAction * QToolBar::insertSeparator(QAction * before);
extern void* C_ZN8QToolBar15insertSeparatorEP7QAction(void* qthis, void* arg0); // 4
  // proto:  QAction * QToolBar::toggleViewAction();
extern void* C_ZNK8QToolBar16toggleViewActionEv(void* qthis); // 4
  // proto:  void QToolBar::setMovable(bool movable);
extern void C_ZN8QToolBar10setMovableEb(void* qthis, bool arg0); // 4
  // proto:  bool QToolBar::isMovable();
extern bool C_ZNK8QToolBar9isMovableEv(void* qthis); // 4
  // proto:  void QToolBar::QToolBar(QWidget * parent);
extern void* C_ZN8QToolBarC2EP7QWidget(void* arg0); // 3
  // proto:  void QToolBar::QToolBar(const QString & title, QWidget * parent);
extern void* C_ZN8QToolBarC2ERK7QStringP7QWidget(void* arg0, void* arg1); // 3
  // proto:  QAction * QToolBar::addSeparator();
extern void* C_ZN8QToolBar12addSeparatorEv(void* qthis); // 4
  // proto:  QSize QToolBar::iconSize();
extern void* C_ZNK8QToolBar8iconSizeEv(void* qthis); // 4
  // proto:  QAction * QToolBar::insertWidget(QAction * before, QWidget * widget);
extern void* C_ZN8QToolBar12insertWidgetEP7QActionP7QWidget(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QToolBar::~QToolBar();
extern void C_ZN8QToolBarD2Ev(void* qthis); // 4
  // proto:  QWidget * QToolBar::widgetForAction(QAction * action);
extern void* C_ZNK8QToolBar15widgetForActionEP7QAction(void* qthis, void* arg0); // 4
  // proto:  QAction * QToolBar::addWidget(QWidget * widget);
extern void* C_ZN8QToolBar9addWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QToolBar::metaObject();
extern void C_ZNK8QToolBar10metaObjectEv(void* qthis); // 4
  // proto:  void QToolBar::setFloatable(bool floatable);
extern void C_ZN8QToolBar12setFloatableEb(void* qthis, bool arg0); // 4
  // proto:  Qt::ToolBarAreas QToolBar::allowedAreas();
extern void C_ZNK8QToolBar12allowedAreasEv(void* qthis); // 4
  // proto:  void QToolBar::clear();
extern void C_ZN8QToolBar5clearEv(void* qthis); // 4
  // proto:  void QToolBar::setIconSize(const QSize & iconSize);
extern void C_ZN8QToolBar11setIconSizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  Qt::ToolButtonStyle QToolBar::toolButtonStyle();
extern void C_ZNK8QToolBar15toolButtonStyleEv(void* qthis); // 4
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

// class sizeof(QToolBar)=1
type QToolBar struct {
  /*qbase*/ QWidget;
  Qclsinst unsafe.Pointer /* *C.void */;
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
func (this *QToolBar) Addaction(args ...interface{}) (ret interface{}) {
  // addAction(const class QIcon &, const class QString &, const class QObject *, const char *)
  // addAction(const class QString &, const class QObject *, const char *)
  // addAction(const class QIcon &, const class QString &)
  // addAction(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QIcon{}) // "const QIcon &"
  vtys[0][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(qtcore.QObject{}) // "const QObject *"
  vtys[0][3] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(qtcore.QObject{}) // "const QObject *"
  vtys[1][2] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtgui.QIcon{}) // "const QIcon &"
  vtys[2][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBar9addActionERK5QIconRK7QStringPK7QObjectPKc
    // invoke: QAction * addAction(const class QIcon &, const class QString &, const class QObject *, const char *)
    var arg0 = args[0].(*qtgui.QIcon).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg2)}
    argif3, free3 := qtrt.HandyConvert2c(args[3], vtys[0][3])
    var arg3 = argif3.(unsafe.Pointer)
    if false {fmt.Println(argif3, arg3)}
    if free3 {defer C.free(arg3)}
    var ret0 = C.C_ZN8QToolBar9addActionERK5QIconRK7QStringPK7QObjectPKc(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN8QToolBar9addActionERK7QStringPK7QObjectPKc
    // invoke: QAction * addAction(const class QString &, const class QObject *, const char *)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg1)}
    argif2, free2 := qtrt.HandyConvert2c(args[2], vtys[1][2])
    var arg2 = argif2.(unsafe.Pointer)
    if false {fmt.Println(argif2, arg2)}
    if free2 {defer C.free(arg2)}
    var ret0 = C.C_ZN8QToolBar9addActionERK7QStringPK7QObjectPKc(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZN8QToolBar9addActionERK5QIconRK7QString
    // invoke: QAction * addAction(const class QIcon &, const class QString &)
    var arg0 = args[0].(*qtgui.QIcon).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN8QToolBar9addActionERK5QIconRK7QString(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZN8QToolBar9addActionERK7QString
    // invoke: QAction * addAction(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN8QToolBar9addActionERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolBar", "addAction", args)
  }

  return
}

// orientation()
func (this *QToolBar) Orientation(args ...interface{}) () {
  // orientation()
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
    // invoke: _ZNK8QToolBar11orientationEv
    // invoke: Qt::Orientation orientation()
    C.C_ZNK8QToolBar11orientationEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QToolBar", "orientation", args)
  }

  return
}

// actionAt(int, int)
func (this *QToolBar) Actionat(args ...interface{}) (ret interface{}) {
  // actionAt(int, int)
  // actionAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBar8actionAtEii
    // invoke: QAction * actionAt(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK8QToolBar8actionAtEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK8QToolBar8actionAtERK6QPoint
    // invoke: QAction * actionAt(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK8QToolBar8actionAtERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolBar", "actionAt", args)
  }

  return
}

// isFloating()
func (this *QToolBar) Isfloating(args ...interface{}) (ret interface{}) {
  // isFloating()
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
    // invoke: _ZNK8QToolBar10isFloatingEv
    // invoke: bool isFloating()
    var ret0 = C.C_ZNK8QToolBar10isFloatingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolBar", "isFloating", args)
  }

  return
}

// actionGeometry(class QAction *)
func (this *QToolBar) Actiongeometry(args ...interface{}) (ret interface{}) {
  // actionGeometry(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBar14actionGeometryEP7QAction
    // invoke: QRect actionGeometry(class QAction *)
    var arg0 = args[0].(*QAction).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK8QToolBar14actionGeometryEP7QAction(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolBar", "actionGeometry", args)
  }

  return
}

// isFloatable()
func (this *QToolBar) Isfloatable(args ...interface{}) (ret interface{}) {
  // isFloatable()
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
    // invoke: _ZNK8QToolBar11isFloatableEv
    // invoke: bool isFloatable()
    var ret0 = C.C_ZNK8QToolBar11isFloatableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolBar", "isFloatable", args)
  }

  return
}

// insertSeparator(class QAction *)
func (this *QToolBar) Insertseparator(args ...interface{}) (ret interface{}) {
  // insertSeparator(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBar15insertSeparatorEP7QAction
    // invoke: QAction * insertSeparator(class QAction *)
    var arg0 = args[0].(*QAction).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN8QToolBar15insertSeparatorEP7QAction(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolBar", "insertSeparator", args)
  }

  return
}

// toggleViewAction()
func (this *QToolBar) Toggleviewaction(args ...interface{}) (ret interface{}) {
  // toggleViewAction()
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
    // invoke: _ZNK8QToolBar16toggleViewActionEv
    // invoke: QAction * toggleViewAction()
    var ret0 = C.C_ZNK8QToolBar16toggleViewActionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolBar", "toggleViewAction", args)
  }

  return
}

// setMovable(_Bool)
func (this *QToolBar) Setmovable(args ...interface{}) () {
  // setMovable(_Bool)
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
    // invoke: _ZN8QToolBar10setMovableEb
    // invoke: void setMovable(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN8QToolBar10setMovableEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolBar", "setMovable", args)
  }

  return
}

// isMovable()
func (this *QToolBar) Ismovable(args ...interface{}) (ret interface{}) {
  // isMovable()
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
    // invoke: _ZNK8QToolBar9isMovableEv
    // invoke: bool isMovable()
    var ret0 = C.C_ZNK8QToolBar9isMovableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolBar", "isMovable", args)
  }

  return
}

// QToolBar(class QWidget *)
func NewQToolBar(args ...interface{}) *QToolBar {
  // QToolBar(class QWidget *)
  // QToolBar(const class QString &, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBarC1EP7QWidget
    // invoke: void QToolBar(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QToolBarC2EP7QWidget(arg0)
    return &QToolBar{Qclsinst:qthis}
  case 1:
    // invoke: _ZN8QToolBarC1ERK7QStringP7QWidget
    // invoke: void QToolBar(const class QString &, class QWidget *)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QToolBarC2ERK7QStringP7QWidget(arg0, arg1)
    return &QToolBar{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QToolBar", "QToolBar", args)
  }

  return nil // QToolBar{Qclsinst:qthis}
}

// addSeparator()
func (this *QToolBar) Addseparator(args ...interface{}) (ret interface{}) {
  // addSeparator()
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
    // invoke: _ZN8QToolBar12addSeparatorEv
    // invoke: QAction * addSeparator()
    var ret0 = C.C_ZN8QToolBar12addSeparatorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolBar", "addSeparator", args)
  }

  return
}

// iconSize()
func (this *QToolBar) Iconsize(args ...interface{}) (ret interface{}) {
  // iconSize()
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
    // invoke: _ZNK8QToolBar8iconSizeEv
    // invoke: QSize iconSize()
    var ret0 = C.C_ZNK8QToolBar8iconSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolBar", "iconSize", args)
  }

  return
}

// insertWidget(class QAction *, class QWidget *)
func (this *QToolBar) Insertwidget(args ...interface{}) (ret interface{}) {
  // insertWidget(class QAction *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBar12insertWidgetEP7QActionP7QWidget
    // invoke: QAction * insertWidget(class QAction *, class QWidget *)
    var arg0 = args[0].(*QAction).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN8QToolBar12insertWidgetEP7QActionP7QWidget(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolBar", "insertWidget", args)
  }

  return
}

// ~QToolBar()
func (this *QToolBar) Freeqtoolbar(args ...interface{}) () {
  // ~QToolBar()
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
    // invoke: _ZN8QToolBarD0Ev
    // invoke: void ~QToolBar()
    C.C_ZN8QToolBarD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QToolBar", "~QToolBar", args)
  }

  return
}

// widgetForAction(class QAction *)
func (this *QToolBar) Widgetforaction(args ...interface{}) (ret interface{}) {
  // widgetForAction(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBar15widgetForActionEP7QAction
    // invoke: QWidget * widgetForAction(class QAction *)
    var arg0 = args[0].(*QAction).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK8QToolBar15widgetForActionEP7QAction(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolBar", "widgetForAction", args)
  }

  return
}

// addWidget(class QWidget *)
func (this *QToolBar) Addwidget(args ...interface{}) (ret interface{}) {
  // addWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBar9addWidgetEP7QWidget
    // invoke: QAction * addWidget(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN8QToolBar9addWidgetEP7QWidget(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolBar", "addWidget", args)
  }

  return
}

// metaObject()
func (this *QToolBar) Metaobject(args ...interface{}) () {
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
    // invoke: _ZNK8QToolBar10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK8QToolBar10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QToolBar", "metaObject", args)
  }

  return
}

// setFloatable(_Bool)
func (this *QToolBar) Setfloatable(args ...interface{}) () {
  // setFloatable(_Bool)
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
    // invoke: _ZN8QToolBar12setFloatableEb
    // invoke: void setFloatable(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN8QToolBar12setFloatableEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolBar", "setFloatable", args)
  }

  return
}

// allowedAreas()
func (this *QToolBar) Allowedareas(args ...interface{}) () {
  // allowedAreas()
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
    // invoke: _ZNK8QToolBar12allowedAreasEv
    // invoke: Qt::ToolBarAreas allowedAreas()
    C.C_ZNK8QToolBar12allowedAreasEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QToolBar", "allowedAreas", args)
  }

  return
}

// clear()
func (this *QToolBar) Clear(args ...interface{}) () {
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
    // invoke: _ZN8QToolBar5clearEv
    // invoke: void clear()
    C.C_ZN8QToolBar5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QToolBar", "clear", args)
  }

  return
}

// setIconSize(const class QSize &)
func (this *QToolBar) Seticonsize(args ...interface{}) () {
  // setIconSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QSize{}) // "const QSize &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBar11setIconSizeERK5QSize
    // invoke: void setIconSize(const class QSize &)
    var arg0 = args[0].(*qtcore.QSize).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QToolBar11setIconSizeERK5QSize(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolBar", "setIconSize", args)
  }

  return
}

// toolButtonStyle()
func (this *QToolBar) Toolbuttonstyle(args ...interface{}) () {
  // toolButtonStyle()
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
    // invoke: _ZNK8QToolBar15toolButtonStyleEv
    // invoke: Qt::ToolButtonStyle toolButtonStyle()
    C.C_ZNK8QToolBar15toolButtonStyleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QToolBar", "toolButtonStyle", args)
  }

  return
}

// <= body block end

