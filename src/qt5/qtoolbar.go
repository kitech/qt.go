package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtWidgets/qtoolbar.h
// dst-file: /src/widgets/qtoolbar.go
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
// class sizeof(QToolBar)=1
type QToolBar struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
//  _iconSizeChanged QToolBar_iconSizeChanged_signal;
//  _allowedAreasChanged QToolBar_allowedAreasChanged_signal;
//  _movableChanged QToolBar_movableChanged_signal;
//  _toolButtonStyleChanged QToolBar_toolButtonStyleChanged_signal;
//  _topLevelChanged QToolBar_topLevelChanged_signal;
//  _actionTriggered QToolBar_actionTriggered_signal;
//  _orientationChanged QToolBar_orientationChanged_signal;
//  _visibilityChanged QToolBar_visibilityChanged_signal;
}


func (this *QToolBar) addAction(args ...interface{}) () {
  // addAction(const class QIcon &, const class QString &, const class QObject *, const char *)
  // addAction(const class QString &)
  // addAction(const class QIcon &, const class QString &)
  // addAction(const class QString &, const class QObject *, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[0][3] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3][1] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[3][2] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN8QToolBar9addActionERK5QIconRK7QStringPK7QObjectPKc
  case 1:
    // invoke: _ZN8QToolBar9addActionERK7QString
  case 2:
    // invoke: _ZN8QToolBar9addActionERK5QIconRK7QString
  case 3:
    // invoke: _ZN8QToolBar9addActionERK7QStringPK7QObjectPKc
  default:
    qtrt.ErrorResolve("QToolBar", "addAction", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QToolBar", "isFloatable", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QToolBar", "iconSize", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QToolBar", "actionGeometry", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QToolBar", "widgetForAction", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QToolBar", "clear", args)
 }

}


func NewQToolBar(args ...interface{})() {
}


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
  default:
    qtrt.ErrorResolve("QToolBar", "setMovable", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QToolBar", "isMovable", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QToolBar", "setIconSize", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QToolBar", "addSeparator", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QToolBar", "setFloatable", args)
 }

}


func (this *QToolBar) actionAt(args ...interface{}) () {
  // actionAt(const class QPoint &)
  // actionAt(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBar8actionAtERK6QPoint
  case 1:
    // invoke: _ZNK8QToolBar8actionAtEii
  default:
    qtrt.ErrorResolve("QToolBar", "actionAt", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QToolBar", "isFloating", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QToolBar", "toggleViewAction", args)
 }

}


func (this *QToolBar) FreeQToolBar(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QToolBar", "~QToolBar", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QToolBar", "insertWidget", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QToolBar", "addWidget", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QToolBar", "metaObject", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QToolBar", "insertSeparator", args)
 }

}

// <= body block end

