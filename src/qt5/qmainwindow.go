package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtWidgets/qmainwindow.h
// dst-file: /src/widgets/qmainwindow.go
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
  // proto:  QStatusBar * QMainWindow::statusBar();
extern void _ZNK11QMainWindow9statusBarEv(void* qthis);
  // proto:  void QMainWindow::setAnimated(bool enabled);
extern void _ZN11QMainWindow11setAnimatedEb(void* qthis, bool arg0);
  // proto:  void QMainWindow::setDockNestingEnabled(bool enabled);
extern void _ZN11QMainWindow21setDockNestingEnabledEb(void* qthis, bool arg0);
  // proto:  bool QMainWindow::unifiedTitleAndToolBarOnMac();
extern void _ZNK11QMainWindow27unifiedTitleAndToolBarOnMacEv(void* qthis);
  // proto:  QWidget * QMainWindow::menuWidget();
extern void _ZNK11QMainWindow10menuWidgetEv(void* qthis);
  // proto:  void QMainWindow::tabifyDockWidget(QDockWidget * first, QDockWidget * second);
extern void _ZN11QMainWindow16tabifyDockWidgetEP11QDockWidgetS1_(void* qthis, void* arg0, void* arg1);
  // proto:  void QMainWindow::setDocumentMode(bool enabled);
extern void _ZN11QMainWindow15setDocumentModeEb(void* qthis, bool arg0);
  // proto:  QWidget * QMainWindow::centralWidget();
extern void _ZNK11QMainWindow13centralWidgetEv(void* qthis);
  // proto:  void QMainWindow::removeDockWidget(QDockWidget * dockwidget);
extern void _ZN11QMainWindow16removeDockWidgetEP11QDockWidget(void* qthis, void* arg0);
  // proto:  void QMainWindow::QMainWindow(const QMainWindow & );
extern void* dector_ZN11QMainWindowC1ERKS_(void* arg0);
extern void _ZN11QMainWindowC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QMainWindow::isAnimated();
extern void _ZNK11QMainWindow10isAnimatedEv(void* qthis);
  // proto:  QToolBar * QMainWindow::addToolBar(const QString & title);
extern void _ZN11QMainWindow10addToolBarERK7QString(void* qthis, void* arg0);
  // proto:  void QMainWindow::setIconSize(const QSize & iconSize);
extern void _ZN11QMainWindow11setIconSizeERK5QSize(void* qthis, void* arg0);
  // proto:  QByteArray QMainWindow::saveState(int version);
extern void _ZNK11QMainWindow9saveStateEi(void* qthis, int arg0);
  // proto:  bool QMainWindow::restoreState(const QByteArray & state, int version);
extern void _ZN11QMainWindow12restoreStateERK10QByteArrayi(void* qthis, void* arg0, int arg1);
  // proto:  void QMainWindow::insertToolBar(QToolBar * before, QToolBar * toolbar);
extern void _ZN11QMainWindow13insertToolBarEP8QToolBarS1_(void* qthis, void* arg0, void* arg1);
  // proto:  QMenu * QMainWindow::createPopupMenu();
extern void _ZN11QMainWindow15createPopupMenuEv(void* qthis);
  // proto:  void QMainWindow::setUnifiedTitleAndToolBarOnMac(bool set);
extern void _ZN11QMainWindow30setUnifiedTitleAndToolBarOnMacEb(void* qthis, bool arg0);
  // proto:  void QMainWindow::addToolBar(QToolBar * toolbar);
extern void _ZN11QMainWindow10addToolBarEP8QToolBar(void* qthis, void* arg0);
  // proto:  void QMainWindow::removeToolBarBreak(QToolBar * before);
extern void _ZN11QMainWindow18removeToolBarBreakEP8QToolBar(void* qthis, void* arg0);
  // proto:  bool QMainWindow::toolBarBreak(QToolBar * toolbar);
extern void _ZNK11QMainWindow12toolBarBreakEP8QToolBar(void* qthis, void* arg0);
  // proto:  bool QMainWindow::restoreDockWidget(QDockWidget * dockwidget);
extern void _ZN11QMainWindow17restoreDockWidgetEP11QDockWidget(void* qthis, void* arg0);
  // proto:  QMenuBar * QMainWindow::menuBar();
extern void _ZNK11QMainWindow7menuBarEv(void* qthis);
  // proto:  void QMainWindow::setStatusBar(QStatusBar * statusbar);
extern void _ZN11QMainWindow12setStatusBarEP10QStatusBar(void* qthis, void* arg0);
  // proto:  void QMainWindow::~QMainWindow();
extern void _ZN11QMainWindowD0Ev(void* qthis);
  // proto:  bool QMainWindow::isSeparator(const QPoint & pos);
extern void _ZNK11QMainWindow11isSeparatorERK6QPoint(void* qthis, void* arg0);
  // proto:  QSize QMainWindow::iconSize();
extern void _ZNK11QMainWindow8iconSizeEv(void* qthis);
  // proto:  const QMetaObject * QMainWindow::metaObject();
extern void _ZNK11QMainWindow10metaObjectEv(void* qthis);
  // proto:  void QMainWindow::insertToolBarBreak(QToolBar * before);
extern void _ZN11QMainWindow18insertToolBarBreakEP8QToolBar(void* qthis, void* arg0);
  // proto:  QWidget * QMainWindow::takeCentralWidget();
extern void _ZN11QMainWindow17takeCentralWidgetEv(void* qthis);
  // proto:  bool QMainWindow::isDockNestingEnabled();
extern void _ZNK11QMainWindow20isDockNestingEnabledEv(void* qthis);
  // proto:  bool QMainWindow::documentMode();
extern void _ZNK11QMainWindow12documentModeEv(void* qthis);
  // proto:  void QMainWindow::setMenuWidget(QWidget * menubar);
extern void _ZN11QMainWindow13setMenuWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  void QMainWindow::removeToolBar(QToolBar * toolbar);
extern void _ZN11QMainWindow13removeToolBarEP8QToolBar(void* qthis, void* arg0);
  // proto:  void QMainWindow::setCentralWidget(QWidget * widget);
extern void _ZN11QMainWindow16setCentralWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  void QMainWindow::setMenuBar(QMenuBar * menubar);
extern void _ZN11QMainWindow10setMenuBarEP8QMenuBar(void* qthis, void* arg0);
  // proto:  QList<QDockWidget *> QMainWindow::tabifiedDockWidgets(QDockWidget * dockwidget);
extern void _ZNK11QMainWindow19tabifiedDockWidgetsEP11QDockWidget(void* qthis, void* arg0);
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

// class sizeof(QMainWindow)=1
type QMainWindow struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
//  _toolButtonStyleChanged QMainWindow_toolButtonStyleChanged_signal;
//  _iconSizeChanged QMainWindow_iconSizeChanged_signal;
}

  // proto:  QStatusBar * QMainWindow::statusBar();
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
    // invoke: QStatusBar * statusBar()
    C._ZNK11QMainWindow9statusBarEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMainWindow", "statusBar", args)
  }

}

  // proto:  void QMainWindow::setAnimated(bool enabled);
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
    // invoke: void setAnimated(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN11QMainWindow11setAnimatedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "setAnimated", args)
  }

}

  // proto:  void QMainWindow::setDockNestingEnabled(bool enabled);
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
    // invoke: void setDockNestingEnabled(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN11QMainWindow21setDockNestingEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "setDockNestingEnabled", args)
  }

}

  // proto:  bool QMainWindow::unifiedTitleAndToolBarOnMac();
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
    // invoke: bool unifiedTitleAndToolBarOnMac()
    C._ZNK11QMainWindow27unifiedTitleAndToolBarOnMacEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMainWindow", "unifiedTitleAndToolBarOnMac", args)
  }

}

  // proto:  QWidget * QMainWindow::menuWidget();
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
    // invoke: QWidget * menuWidget()
    C._ZNK11QMainWindow10menuWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMainWindow", "menuWidget", args)
  }

}

  // proto:  void QMainWindow::tabifyDockWidget(QDockWidget * first, QDockWidget * second);
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
    // invoke: void tabifyDockWidget(class QDockWidget *, class QDockWidget *)
    var arg0 = args[0].(QDockWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QDockWidget).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN11QMainWindow16tabifyDockWidgetEP11QDockWidgetS1_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMainWindow", "tabifyDockWidget", args)
  }

}

  // proto:  void QMainWindow::setDocumentMode(bool enabled);
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
    // invoke: void setDocumentMode(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN11QMainWindow15setDocumentModeEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "setDocumentMode", args)
  }

}

  // proto:  QWidget * QMainWindow::centralWidget();
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
    // invoke: QWidget * centralWidget()
    C._ZNK11QMainWindow13centralWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMainWindow", "centralWidget", args)
  }

}

  // proto:  void QMainWindow::removeDockWidget(QDockWidget * dockwidget);
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
    // invoke: void removeDockWidget(class QDockWidget *)
    var arg0 = args[0].(QDockWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QMainWindow16removeDockWidgetEP11QDockWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "removeDockWidget", args)
  }

}

  // proto:  void QMainWindow::QMainWindow(const QMainWindow & );
func NewQMainWindow(args ...interface{}) QMainWindow {
  return QMainWindow{}
}

  // proto:  bool QMainWindow::isAnimated();
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
    // invoke: bool isAnimated()
    C._ZNK11QMainWindow10isAnimatedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMainWindow", "isAnimated", args)
  }

}

  // proto:  QToolBar * QMainWindow::addToolBar(const QString & title);
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
    // invoke: QToolBar * addToolBar(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QMainWindow10addToolBarERK7QString(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN11QMainWindow10addToolBarEP8QToolBar
    // invoke: void addToolBar(class QToolBar *)
    var arg0 = args[0].(QToolBar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QMainWindow10addToolBarEP8QToolBar(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "addToolBar", args)
  }

}

  // proto:  void QMainWindow::setIconSize(const QSize & iconSize);
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
    // invoke: void setIconSize(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QMainWindow11setIconSizeERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "setIconSize", args)
  }

}

  // proto:  QByteArray QMainWindow::saveState(int version);
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
    // invoke: QByteArray saveState(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK11QMainWindow9saveStateEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "saveState", args)
  }

}

  // proto:  bool QMainWindow::restoreState(const QByteArray & state, int version);
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
    // invoke: bool restoreState(const class QByteArray &, int)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN11QMainWindow12restoreStateERK10QByteArrayi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMainWindow", "restoreState", args)
  }

}

  // proto:  void QMainWindow::insertToolBar(QToolBar * before, QToolBar * toolbar);
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
    // invoke: void insertToolBar(class QToolBar *, class QToolBar *)
    var arg0 = args[0].(QToolBar).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QToolBar).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN11QMainWindow13insertToolBarEP8QToolBarS1_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMainWindow", "insertToolBar", args)
  }

}

  // proto:  QMenu * QMainWindow::createPopupMenu();
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
    // invoke: QMenu * createPopupMenu()
    C._ZN11QMainWindow15createPopupMenuEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMainWindow", "createPopupMenu", args)
  }

}

  // proto:  void QMainWindow::setUnifiedTitleAndToolBarOnMac(bool set);
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
    // invoke: void setUnifiedTitleAndToolBarOnMac(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN11QMainWindow30setUnifiedTitleAndToolBarOnMacEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "setUnifiedTitleAndToolBarOnMac", args)
  }

}

  // proto:  void QMainWindow::removeToolBarBreak(QToolBar * before);
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
    // invoke: void removeToolBarBreak(class QToolBar *)
    var arg0 = args[0].(QToolBar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QMainWindow18removeToolBarBreakEP8QToolBar(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "removeToolBarBreak", args)
  }

}

  // proto:  bool QMainWindow::toolBarBreak(QToolBar * toolbar);
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
    // invoke: bool toolBarBreak(class QToolBar *)
    var arg0 = args[0].(QToolBar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK11QMainWindow12toolBarBreakEP8QToolBar(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "toolBarBreak", args)
  }

}

  // proto:  bool QMainWindow::restoreDockWidget(QDockWidget * dockwidget);
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
    // invoke: bool restoreDockWidget(class QDockWidget *)
    var arg0 = args[0].(QDockWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QMainWindow17restoreDockWidgetEP11QDockWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "restoreDockWidget", args)
  }

}

  // proto:  QMenuBar * QMainWindow::menuBar();
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
    // invoke: QMenuBar * menuBar()
    C._ZNK11QMainWindow7menuBarEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMainWindow", "menuBar", args)
  }

}

  // proto:  void QMainWindow::setStatusBar(QStatusBar * statusbar);
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
    // invoke: void setStatusBar(class QStatusBar *)
    var arg0 = args[0].(QStatusBar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QMainWindow12setStatusBarEP10QStatusBar(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "setStatusBar", args)
  }

}

  // proto:  void QMainWindow::~QMainWindow();
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

  // proto:  bool QMainWindow::isSeparator(const QPoint & pos);
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
    // invoke: bool isSeparator(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK11QMainWindow11isSeparatorERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "isSeparator", args)
  }

}

  // proto:  QSize QMainWindow::iconSize();
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
    // invoke: QSize iconSize()
    C._ZNK11QMainWindow8iconSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMainWindow", "iconSize", args)
  }

}

  // proto:  const QMetaObject * QMainWindow::metaObject();
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
    // invoke: const QMetaObject * metaObject()
    C._ZNK11QMainWindow10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMainWindow", "metaObject", args)
  }

}

  // proto:  void QMainWindow::insertToolBarBreak(QToolBar * before);
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
    // invoke: void insertToolBarBreak(class QToolBar *)
    var arg0 = args[0].(QToolBar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QMainWindow18insertToolBarBreakEP8QToolBar(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "insertToolBarBreak", args)
  }

}

  // proto:  QWidget * QMainWindow::takeCentralWidget();
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
    // invoke: QWidget * takeCentralWidget()
    C._ZN11QMainWindow17takeCentralWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMainWindow", "takeCentralWidget", args)
  }

}

  // proto:  bool QMainWindow::isDockNestingEnabled();
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
    // invoke: bool isDockNestingEnabled()
    C._ZNK11QMainWindow20isDockNestingEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMainWindow", "isDockNestingEnabled", args)
  }

}

  // proto:  bool QMainWindow::documentMode();
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
    // invoke: bool documentMode()
    C._ZNK11QMainWindow12documentModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMainWindow", "documentMode", args)
  }

}

  // proto:  void QMainWindow::setMenuWidget(QWidget * menubar);
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
    // invoke: void setMenuWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QMainWindow13setMenuWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "setMenuWidget", args)
  }

}

  // proto:  void QMainWindow::removeToolBar(QToolBar * toolbar);
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
    // invoke: void removeToolBar(class QToolBar *)
    var arg0 = args[0].(QToolBar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QMainWindow13removeToolBarEP8QToolBar(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "removeToolBar", args)
  }

}

  // proto:  void QMainWindow::setCentralWidget(QWidget * widget);
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
    // invoke: void setCentralWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QMainWindow16setCentralWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "setCentralWidget", args)
  }

}

  // proto:  void QMainWindow::setMenuBar(QMenuBar * menubar);
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
    // invoke: void setMenuBar(class QMenuBar *)
    var arg0 = args[0].(QMenuBar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QMainWindow10setMenuBarEP8QMenuBar(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "setMenuBar", args)
  }

}

  // proto:  QList<QDockWidget *> QMainWindow::tabifiedDockWidgets(QDockWidget * dockwidget);
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
    // invoke: QList<QDockWidget *> tabifiedDockWidgets(class QDockWidget *)
    var arg0 = args[0].(QDockWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK11QMainWindow19tabifiedDockWidgetsEP11QDockWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "tabifiedDockWidgets", args)
  }

}

// <= body block end

