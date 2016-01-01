package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtWidgets/qmainwindow.h
// dst-file: /src/widgets/qmainwindow.go
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
// class sizeof(QMainWindow)=1
type QMainWindow struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
//  _toolButtonStyleChanged QMainWindow_toolButtonStyleChanged_signal;
//  _iconSizeChanged QMainWindow_iconSizeChanged_signal;
}


func (this *QMainWindow) statusBar(args ...interface{}) () {
  // statusBar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QMainWindow9statusBarEv
  default:
    qtrt.ErrorResolve("QMainWindow", "statusBar", args)
 }

}


func (this *QMainWindow) setAnimated(args ...interface{}) () {
  // setAnimated(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QMainWindow11setAnimatedEb
  default:
    qtrt.ErrorResolve("QMainWindow", "setAnimated", args)
 }

}


func (this *QMainWindow) setDockNestingEnabled(args ...interface{}) () {
  // setDockNestingEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QMainWindow21setDockNestingEnabledEb
  default:
    qtrt.ErrorResolve("QMainWindow", "setDockNestingEnabled", args)
 }

}


func (this *QMainWindow) unifiedTitleAndToolBarOnMac(args ...interface{}) () {
  // unifiedTitleAndToolBarOnMac()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QMainWindow27unifiedTitleAndToolBarOnMacEv
  default:
    qtrt.ErrorResolve("QMainWindow", "unifiedTitleAndToolBarOnMac", args)
 }

}


func (this *QMainWindow) menuWidget(args ...interface{}) () {
  // menuWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QMainWindow10menuWidgetEv
  default:
    qtrt.ErrorResolve("QMainWindow", "menuWidget", args)
 }

}


func (this *QMainWindow) tabifyDockWidget(args ...interface{}) () {
  // tabifyDockWidget(class QDockWidget *, class QDockWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDockWidget{}) // "QDockWidget *"
  vtys[0][1] = reflect.TypeOf(QDockWidget{}) // "QDockWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QMainWindow16tabifyDockWidgetEP11QDockWidgetS1_
  default:
    qtrt.ErrorResolve("QMainWindow", "tabifyDockWidget", args)
 }

}


func (this *QMainWindow) setDocumentMode(args ...interface{}) () {
  // setDocumentMode(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QMainWindow15setDocumentModeEb
  default:
    qtrt.ErrorResolve("QMainWindow", "setDocumentMode", args)
 }

}


func (this *QMainWindow) centralWidget(args ...interface{}) () {
  // centralWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QMainWindow13centralWidgetEv
  default:
    qtrt.ErrorResolve("QMainWindow", "centralWidget", args)
 }

}


func (this *QMainWindow) removeDockWidget(args ...interface{}) () {
  // removeDockWidget(class QDockWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDockWidget{}) // "QDockWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QMainWindow16removeDockWidgetEP11QDockWidget
  default:
    qtrt.ErrorResolve("QMainWindow", "removeDockWidget", args)
 }

}


func NewQMainWindow(args ...interface{})() {
}


func (this *QMainWindow) isAnimated(args ...interface{}) () {
  // isAnimated()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QMainWindow10isAnimatedEv
  default:
    qtrt.ErrorResolve("QMainWindow", "isAnimated", args)
 }

}


func (this *QMainWindow) addToolBar(args ...interface{}) () {
  // addToolBar(const class QString &)
  // addToolBar(class QToolBar *)
  // addToolBar(Qt::ToolBarArea, class QToolBar *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QToolBar{}) // "QToolBar *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "Qt::ToolBarArea"
  vtys[2][1] = reflect.TypeOf(QToolBar{}) // "QToolBar *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QMainWindow10addToolBarERK7QString
  case 1:
    // invoke: _ZN11QMainWindow10addToolBarEP8QToolBar
  case 2:
    // invoke: _ZN11QMainWindow10addToolBarEN2Qt11ToolBarAreaEP8QToolBar
  default:
    qtrt.ErrorResolve("QMainWindow", "addToolBar", args)
 }

}


func (this *QMainWindow) setIconSize(args ...interface{}) () {
  // setIconSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QMainWindow11setIconSizeERK5QSize
  default:
    qtrt.ErrorResolve("QMainWindow", "setIconSize", args)
 }

}


func (this *QMainWindow) saveState(args ...interface{}) () {
  // saveState(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QMainWindow9saveStateEi
  default:
    qtrt.ErrorResolve("QMainWindow", "saveState", args)
 }

}


func (this *QMainWindow) restoreState(args ...interface{}) () {
  // restoreState(const class QByteArray &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QMainWindow12restoreStateERK10QByteArrayi
  default:
    qtrt.ErrorResolve("QMainWindow", "restoreState", args)
 }

}


func (this *QMainWindow) insertToolBar(args ...interface{}) () {
  // insertToolBar(class QToolBar *, class QToolBar *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QToolBar{}) // "QToolBar *"
  vtys[0][1] = reflect.TypeOf(QToolBar{}) // "QToolBar *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QMainWindow13insertToolBarEP8QToolBarS1_
  default:
    qtrt.ErrorResolve("QMainWindow", "insertToolBar", args)
 }

}


func (this *QMainWindow) createPopupMenu(args ...interface{}) () {
  // createPopupMenu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QMainWindow15createPopupMenuEv
  default:
    qtrt.ErrorResolve("QMainWindow", "createPopupMenu", args)
 }

}


func (this *QMainWindow) setUnifiedTitleAndToolBarOnMac(args ...interface{}) () {
  // setUnifiedTitleAndToolBarOnMac(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QMainWindow30setUnifiedTitleAndToolBarOnMacEb
  default:
    qtrt.ErrorResolve("QMainWindow", "setUnifiedTitleAndToolBarOnMac", args)
 }

}


func (this *QMainWindow) removeToolBarBreak(args ...interface{}) () {
  // removeToolBarBreak(class QToolBar *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QToolBar{}) // "QToolBar *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QMainWindow18removeToolBarBreakEP8QToolBar
  default:
    qtrt.ErrorResolve("QMainWindow", "removeToolBarBreak", args)
 }

}


func (this *QMainWindow) toolBarBreak(args ...interface{}) () {
  // toolBarBreak(class QToolBar *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QToolBar{}) // "QToolBar *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QMainWindow12toolBarBreakEP8QToolBar
  default:
    qtrt.ErrorResolve("QMainWindow", "toolBarBreak", args)
 }

}


func (this *QMainWindow) restoreDockWidget(args ...interface{}) () {
  // restoreDockWidget(class QDockWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDockWidget{}) // "QDockWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QMainWindow17restoreDockWidgetEP11QDockWidget
  default:
    qtrt.ErrorResolve("QMainWindow", "restoreDockWidget", args)
 }

}


func (this *QMainWindow) menuBar(args ...interface{}) () {
  // menuBar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QMainWindow7menuBarEv
  default:
    qtrt.ErrorResolve("QMainWindow", "menuBar", args)
 }

}


func (this *QMainWindow) setStatusBar(args ...interface{}) () {
  // setStatusBar(class QStatusBar *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStatusBar{}) // "QStatusBar *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QMainWindow12setStatusBarEP10QStatusBar
  default:
    qtrt.ErrorResolve("QMainWindow", "setStatusBar", args)
 }

}


func (this *QMainWindow) FreeQMainWindow(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QMainWindow", "~QMainWindow", args)
 }

}


func (this *QMainWindow) isSeparator(args ...interface{}) () {
  // isSeparator(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QMainWindow11isSeparatorERK6QPoint
  default:
    qtrt.ErrorResolve("QMainWindow", "isSeparator", args)
 }

}


func (this *QMainWindow) iconSize(args ...interface{}) () {
  // iconSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QMainWindow8iconSizeEv
  default:
    qtrt.ErrorResolve("QMainWindow", "iconSize", args)
 }

}


func (this *QMainWindow) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QMainWindow10metaObjectEv
  default:
    qtrt.ErrorResolve("QMainWindow", "metaObject", args)
 }

}


func (this *QMainWindow) insertToolBarBreak(args ...interface{}) () {
  // insertToolBarBreak(class QToolBar *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QToolBar{}) // "QToolBar *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QMainWindow18insertToolBarBreakEP8QToolBar
  default:
    qtrt.ErrorResolve("QMainWindow", "insertToolBarBreak", args)
 }

}


func (this *QMainWindow) takeCentralWidget(args ...interface{}) () {
  // takeCentralWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QMainWindow17takeCentralWidgetEv
  default:
    qtrt.ErrorResolve("QMainWindow", "takeCentralWidget", args)
 }

}


func (this *QMainWindow) isDockNestingEnabled(args ...interface{}) () {
  // isDockNestingEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QMainWindow20isDockNestingEnabledEv
  default:
    qtrt.ErrorResolve("QMainWindow", "isDockNestingEnabled", args)
 }

}


func (this *QMainWindow) documentMode(args ...interface{}) () {
  // documentMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QMainWindow12documentModeEv
  default:
    qtrt.ErrorResolve("QMainWindow", "documentMode", args)
 }

}


func (this *QMainWindow) setMenuWidget(args ...interface{}) () {
  // setMenuWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QMainWindow13setMenuWidgetEP7QWidget
  default:
    qtrt.ErrorResolve("QMainWindow", "setMenuWidget", args)
 }

}


func (this *QMainWindow) removeToolBar(args ...interface{}) () {
  // removeToolBar(class QToolBar *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QToolBar{}) // "QToolBar *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QMainWindow13removeToolBarEP8QToolBar
  default:
    qtrt.ErrorResolve("QMainWindow", "removeToolBar", args)
 }

}


func (this *QMainWindow) setCentralWidget(args ...interface{}) () {
  // setCentralWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QMainWindow16setCentralWidgetEP7QWidget
  default:
    qtrt.ErrorResolve("QMainWindow", "setCentralWidget", args)
 }

}


func (this *QMainWindow) setMenuBar(args ...interface{}) () {
  // setMenuBar(class QMenuBar *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMenuBar{}) // "QMenuBar *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QMainWindow10setMenuBarEP8QMenuBar
  default:
    qtrt.ErrorResolve("QMainWindow", "setMenuBar", args)
 }

}


func (this *QMainWindow) tabifiedDockWidgets(args ...interface{}) () {
  // tabifiedDockWidgets(class QDockWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDockWidget{}) // "QDockWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QMainWindow19tabifiedDockWidgetsEP11QDockWidget
  default:
    qtrt.ErrorResolve("QMainWindow", "tabifiedDockWidgets", args)
 }

}

// <= body block end

