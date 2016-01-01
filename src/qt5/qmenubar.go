package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtWidgets/qmenubar.h
// dst-file: /src/widgets/qmenubar.go
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
// class sizeof(QMenuBar)=1
type QMenuBar struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
//  _hovered QMenuBar_hovered_signal;
//  _triggered QMenuBar_triggered_signal;
}


func (this *QMenuBar) addAction(args ...interface{}) () {
  // addAction(const class QString &)
  // addAction(const class QString &, const class QObject *, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[1][2] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN8QMenuBar9addActionERK7QString
  case 1:
    // invoke: _ZN8QMenuBar9addActionERK7QStringPK7QObjectPKc
  default:
    qtrt.ErrorResolve("QMenuBar", "addAction", args)
 }

}


func (this *QMenuBar) platformMenuBar(args ...interface{}) () {
  // platformMenuBar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN8QMenuBar15platformMenuBarEv
  default:
    qtrt.ErrorResolve("QMenuBar", "platformMenuBar", args)
 }

}


func (this *QMenuBar) setNativeMenuBar(args ...interface{}) () {
  // setNativeMenuBar(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN8QMenuBar16setNativeMenuBarEb
  default:
    qtrt.ErrorResolve("QMenuBar", "setNativeMenuBar", args)
 }

}


func (this *QMenuBar) FreeQMenuBar(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QMenuBar", "~QMenuBar", args)
 }

}


func (this *QMenuBar) addMenu(args ...interface{}) () {
  // addMenu(class QMenu *)
  // addMenu(const class QIcon &, const class QString &)
  // addMenu(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMenu{}) // "QMenu *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN8QMenuBar7addMenuEP5QMenu
  case 1:
    // invoke: _ZN8QMenuBar7addMenuERK5QIconRK7QString
  case 2:
    // invoke: _ZN8QMenuBar7addMenuERK7QString
  default:
    qtrt.ErrorResolve("QMenuBar", "addMenu", args)
 }

}


func (this *QMenuBar) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QMenuBar8sizeHintEv
  default:
    qtrt.ErrorResolve("QMenuBar", "sizeHint", args)
 }

}


func (this *QMenuBar) actionAt(args ...interface{}) () {
  // actionAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QMenuBar8actionAtERK6QPoint
  default:
    qtrt.ErrorResolve("QMenuBar", "actionAt", args)
 }

}


func (this *QMenuBar) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QMenuBar10metaObjectEv
  default:
    qtrt.ErrorResolve("QMenuBar", "metaObject", args)
 }

}


func (this *QMenuBar) isNativeMenuBar(args ...interface{}) () {
  // isNativeMenuBar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QMenuBar15isNativeMenuBarEv
  default:
    qtrt.ErrorResolve("QMenuBar", "isNativeMenuBar", args)
 }

}


func (this *QMenuBar) insertSeparator(args ...interface{}) () {
  // insertSeparator(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN8QMenuBar15insertSeparatorEP7QAction
  default:
    qtrt.ErrorResolve("QMenuBar", "insertSeparator", args)
 }

}


func (this *QMenuBar) addSeparator(args ...interface{}) () {
  // addSeparator()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN8QMenuBar12addSeparatorEv
  default:
    qtrt.ErrorResolve("QMenuBar", "addSeparator", args)
 }

}


func (this *QMenuBar) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QMenuBar15minimumSizeHintEv
  default:
    qtrt.ErrorResolve("QMenuBar", "minimumSizeHint", args)
 }

}


func (this *QMenuBar) isDefaultUp(args ...interface{}) () {
  // isDefaultUp()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QMenuBar11isDefaultUpEv
  default:
    qtrt.ErrorResolve("QMenuBar", "isDefaultUp", args)
 }

}


func NewQMenuBar(args ...interface{})() {
}


func (this *QMenuBar) setActiveAction(args ...interface{}) () {
  // setActiveAction(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN8QMenuBar15setActiveActionEP7QAction
  default:
    qtrt.ErrorResolve("QMenuBar", "setActiveAction", args)
 }

}


func (this *QMenuBar) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN8QMenuBar5clearEv
  default:
    qtrt.ErrorResolve("QMenuBar", "clear", args)
 }

}


func (this *QMenuBar) activeAction(args ...interface{}) () {
  // activeAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QMenuBar12activeActionEv
  default:
    qtrt.ErrorResolve("QMenuBar", "activeAction", args)
 }

}


func (this *QMenuBar) actionGeometry(args ...interface{}) () {
  // actionGeometry(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QMenuBar14actionGeometryEP7QAction
  default:
    qtrt.ErrorResolve("QMenuBar", "actionGeometry", args)
 }

}


func (this *QMenuBar) insertMenu(args ...interface{}) () {
  // insertMenu(class QAction *, class QMenu *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"
  vtys[0][1] = reflect.TypeOf(QMenu{}) // "QMenu *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN8QMenuBar10insertMenuEP7QActionP5QMenu
  default:
    qtrt.ErrorResolve("QMenuBar", "insertMenu", args)
 }

}


func (this *QMenuBar) setDefaultUp(args ...interface{}) () {
  // setDefaultUp(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN8QMenuBar12setDefaultUpEb
  default:
    qtrt.ErrorResolve("QMenuBar", "setDefaultUp", args)
 }

}


func (this *QMenuBar) setVisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN8QMenuBar10setVisibleEb
  default:
    qtrt.ErrorResolve("QMenuBar", "setVisible", args)
 }

}


func (this *QMenuBar) heightForWidth(args ...interface{}) () {
  // heightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QMenuBar14heightForWidthEi
  default:
    qtrt.ErrorResolve("QMenuBar", "heightForWidth", args)
 }

}

// <= body block end

