package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
import "runtime"
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
  // proto:  void QMainWindow::setMenuWidget(QWidget * menubar);
extern void C_ZN11QMainWindow13setMenuWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QMainWindow::removeToolBar(QToolBar * toolbar);
extern void C_ZN11QMainWindow13removeToolBarEP8QToolBar(void* qthis, void* arg0); // 4
  // proto:  Qt::DockWidgetArea QMainWindow::dockWidgetArea(QDockWidget * dockwidget);
extern void C_ZNK11QMainWindow14dockWidgetAreaEP11QDockWidget(void* qthis, void* arg0); // 4
  // proto:  void QMainWindow::setMenuBar(QMenuBar * menubar);
extern void C_ZN11QMainWindow10setMenuBarEP8QMenuBar(void* qthis, void* arg0); // 4
  // proto:  QByteArray QMainWindow::saveState(int version);
extern void* C_ZNK11QMainWindow9saveStateEi(void* qthis, int32_t arg0); // 4
  // proto:  void QMainWindow::insertToolBar(QToolBar * before, QToolBar * toolbar);
extern void C_ZN11QMainWindow13insertToolBarEP8QToolBarS1_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QList<QDockWidget *> QMainWindow::tabifiedDockWidgets(QDockWidget * dockwidget);
extern void C_ZNK11QMainWindow19tabifiedDockWidgetsEP11QDockWidget(void* qthis, void* arg0); // 4
  // proto:  void QMainWindow::addToolBar(QToolBar * toolbar);
extern void C_ZN11QMainWindow10addToolBarEP8QToolBar(void* qthis, void* arg0); // 4
  // proto:  QToolBar * QMainWindow::addToolBar(const QString & title);
extern void* C_ZN11QMainWindow10addToolBarERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QMainWindow::isAnimated();
extern bool C_ZNK11QMainWindow10isAnimatedEv(void* qthis); // 4
  // proto:  void QMainWindow::setUnifiedTitleAndToolBarOnMac(bool set);
extern void C_ZN11QMainWindow30setUnifiedTitleAndToolBarOnMacEb(void* qthis, bool arg0); // 4
  // proto:  void QMainWindow::setCentralWidget(QWidget * widget);
extern void C_ZN11QMainWindow16setCentralWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QMainWindow::removeDockWidget(QDockWidget * dockwidget);
extern void C_ZN11QMainWindow16removeDockWidgetEP11QDockWidget(void* qthis, void* arg0); // 4
  // proto:  Qt::ToolBarArea QMainWindow::toolBarArea(QToolBar * toolbar);
extern void C_ZNK11QMainWindow11toolBarAreaEP8QToolBar(void* qthis, void* arg0); // 4
  // proto:  bool QMainWindow::unifiedTitleAndToolBarOnMac();
extern bool C_ZNK11QMainWindow27unifiedTitleAndToolBarOnMacEv(void* qthis); // 4
  // proto:  QTabWidget::TabShape QMainWindow::tabShape();
extern void C_ZNK11QMainWindow8tabShapeEv(void* qthis); // 4
  // proto:  QSize QMainWindow::iconSize();
extern void* C_ZNK11QMainWindow8iconSizeEv(void* qthis); // 4
  // proto:  void QMainWindow::tabifyDockWidget(QDockWidget * first, QDockWidget * second);
extern void C_ZN11QMainWindow16tabifyDockWidgetEP11QDockWidgetS1_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QMainWindow::toolBarBreak(QToolBar * toolbar);
extern bool C_ZNK11QMainWindow12toolBarBreakEP8QToolBar(void* qthis, void* arg0); // 4
  // proto:  bool QMainWindow::documentMode();
extern bool C_ZNK11QMainWindow12documentModeEv(void* qthis); // 4
  // proto:  void QMainWindow::insertToolBarBreak(QToolBar * before);
extern void C_ZN11QMainWindow18insertToolBarBreakEP8QToolBar(void* qthis, void* arg0); // 4
  // proto:  bool QMainWindow::restoreState(const QByteArray & state, int version);
extern bool C_ZN11QMainWindow12restoreStateERK10QByteArrayi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  bool QMainWindow::restoreDockWidget(QDockWidget * dockwidget);
extern bool C_ZN11QMainWindow17restoreDockWidgetEP11QDockWidget(void* qthis, void* arg0); // 4
  // proto:  QStatusBar * QMainWindow::statusBar();
extern void* C_ZNK11QMainWindow9statusBarEv(void* qthis); // 4
  // proto:  DockOptions QMainWindow::dockOptions();
extern void C_ZNK11QMainWindow11dockOptionsEv(void* qthis); // 4
  // proto:  void QMainWindow::removeToolBarBreak(QToolBar * before);
extern void C_ZN11QMainWindow18removeToolBarBreakEP8QToolBar(void* qthis, void* arg0); // 4
  // proto:  void QMainWindow::setAnimated(bool enabled);
extern void C_ZN11QMainWindow11setAnimatedEb(void* qthis, bool arg0); // 4
  // proto:  const QMetaObject * QMainWindow::metaObject();
extern void C_ZNK11QMainWindow10metaObjectEv(void* qthis); // 4
  // proto:  void QMainWindow::setDocumentMode(bool enabled);
extern void C_ZN11QMainWindow15setDocumentModeEb(void* qthis, bool arg0); // 4
  // proto:  QWidget * QMainWindow::takeCentralWidget();
extern void* C_ZN11QMainWindow17takeCentralWidgetEv(void* qthis); // 4
  // proto:  QMenu * QMainWindow::createPopupMenu();
extern void* C_ZN11QMainWindow15createPopupMenuEv(void* qthis); // 4
  // proto:  void QMainWindow::setDockNestingEnabled(bool enabled);
extern void C_ZN11QMainWindow21setDockNestingEnabledEb(void* qthis, bool arg0); // 4
  // proto:  QWidget * QMainWindow::centralWidget();
extern void* C_ZNK11QMainWindow13centralWidgetEv(void* qthis); // 4
  // proto:  void QMainWindow::setIconSize(const QSize & iconSize);
extern void C_ZN11QMainWindow11setIconSizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  void QMainWindow::setStatusBar(QStatusBar * statusbar);
extern void C_ZN11QMainWindow12setStatusBarEP10QStatusBar(void* qthis, void* arg0); // 4
  // proto:  bool QMainWindow::isSeparator(const QPoint & pos);
extern bool C_ZNK11QMainWindow11isSeparatorERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QWidget * QMainWindow::menuWidget();
extern void* C_ZNK11QMainWindow10menuWidgetEv(void* qthis); // 4
  // proto:  Qt::ToolButtonStyle QMainWindow::toolButtonStyle();
extern void C_ZNK11QMainWindow15toolButtonStyleEv(void* qthis); // 4
  // proto:  void QMainWindow::~QMainWindow();
extern void C_ZN11QMainWindowD2Ev(void* qthis); // 4
  // proto:  bool QMainWindow::isDockNestingEnabled();
extern bool C_ZNK11QMainWindow20isDockNestingEnabledEv(void* qthis); // 4
  // proto:  QMenuBar * QMainWindow::menuBar();
extern void* C_ZNK11QMainWindow7menuBarEv(void* qthis); // 4
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
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QMainWindow)=1
type QMainWindow struct {
  /*qbase*/ QWidget;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _toolButtonStyleChanged QMainWindow_toolButtonStyleChanged_signal;
//  _iconSizeChanged QMainWindow_iconSizeChanged_signal;
}

// setMenuWidget(class QWidget *)
func (this *QMainWindow) SetMenuWidget(args ...interface{}) () {
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMainWindow13setMenuWidgetEP7QWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "setMenuWidget", args)
  }

  return
}

// removeToolBar(class QToolBar *)
func (this *QMainWindow) RemoveToolBar(args ...interface{}) () {
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
    var arg0 = args[0].(*QToolBar).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMainWindow13removeToolBarEP8QToolBar(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "removeToolBar", args)
  }

  return
}

// dockWidgetArea(class QDockWidget *)
func (this *QMainWindow) DockWidgetArea(args ...interface{}) () {
  // dockWidgetArea(class QDockWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDockWidget{}) // "QDockWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMainWindow14dockWidgetAreaEP11QDockWidget
    // invoke: Qt::DockWidgetArea dockWidgetArea(class QDockWidget *)
    var arg0 = args[0].(*QDockWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK11QMainWindow14dockWidgetAreaEP11QDockWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "dockWidgetArea", args)
  }

  return
}

// setMenuBar(class QMenuBar *)
func (this *QMainWindow) SetMenuBar(args ...interface{}) () {
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
    var arg0 = args[0].(*QMenuBar).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMainWindow10setMenuBarEP8QMenuBar(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "setMenuBar", args)
  }

  return
}

// saveState(int)
func (this *QMainWindow) SaveState(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QMainWindow9saveStateEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMainWindow", "saveState", args)
  }

  return
}

// insertToolBar(class QToolBar *, class QToolBar *)
func (this *QMainWindow) InsertToolBar(args ...interface{}) () {
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
    var arg0 = args[0].(*QToolBar).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QToolBar).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QMainWindow13insertToolBarEP8QToolBarS1_(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMainWindow", "insertToolBar", args)
  }

  return
}

// tabifiedDockWidgets(class QDockWidget *)
func (this *QMainWindow) TabifiedDockWidgets(args ...interface{}) () {
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
    var arg0 = args[0].(*QDockWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK11QMainWindow19tabifiedDockWidgetsEP11QDockWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "tabifiedDockWidgets", args)
  }

  return
}

// addToolBar(class QToolBar *)
func (this *QMainWindow) AddToolBar(args ...interface{}) () {
  // addToolBar(class QToolBar *)
  // addToolBar(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QToolBar{}) // "QToolBar *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMainWindow10addToolBarEP8QToolBar
    // invoke: void addToolBar(class QToolBar *)
    var arg0 = args[0].(*QToolBar).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMainWindow10addToolBarEP8QToolBar(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN11QMainWindow10addToolBarERK7QString
    // invoke: QToolBar * addToolBar(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN11QMainWindow10addToolBarERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
  default:
    qtrt.ErrorResolve("QMainWindow", "addToolBar", args)
  }

  return
}

// isAnimated()
func (this *QMainWindow) IsAnimated(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QMainWindow10isAnimatedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMainWindow", "isAnimated", args)
  }

  return
}

// setUnifiedTitleAndToolBarOnMac(_Bool)
func (this *QMainWindow) SetUnifiedTitleAndToolBarOnMac(args ...interface{}) () {
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QMainWindow30setUnifiedTitleAndToolBarOnMacEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "setUnifiedTitleAndToolBarOnMac", args)
  }

  return
}

// setCentralWidget(class QWidget *)
func (this *QMainWindow) SetCentralWidget(args ...interface{}) () {
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMainWindow16setCentralWidgetEP7QWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "setCentralWidget", args)
  }

  return
}

// removeDockWidget(class QDockWidget *)
func (this *QMainWindow) RemoveDockWidget(args ...interface{}) () {
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
    var arg0 = args[0].(*QDockWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMainWindow16removeDockWidgetEP11QDockWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "removeDockWidget", args)
  }

  return
}

// toolBarArea(class QToolBar *)
func (this *QMainWindow) ToolBarArea(args ...interface{}) () {
  // toolBarArea(class QToolBar *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QToolBar{}) // "QToolBar *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMainWindow11toolBarAreaEP8QToolBar
    // invoke: Qt::ToolBarArea toolBarArea(class QToolBar *)
    var arg0 = args[0].(*QToolBar).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK11QMainWindow11toolBarAreaEP8QToolBar(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "toolBarArea", args)
  }

  return
}

// unifiedTitleAndToolBarOnMac()
func (this *QMainWindow) UnifiedTitleAndToolBarOnMac(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QMainWindow27unifiedTitleAndToolBarOnMacEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMainWindow", "unifiedTitleAndToolBarOnMac", args)
  }

  return
}

// tabShape()
func (this *QMainWindow) TabShape(args ...interface{}) () {
  // tabShape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMainWindow8tabShapeEv
    // invoke: QTabWidget::TabShape tabShape()
    C.C_ZNK11QMainWindow8tabShapeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMainWindow", "tabShape", args)
  }

  return
}

// iconSize()
func (this *QMainWindow) IconSize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QMainWindow8iconSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMainWindow", "iconSize", args)
  }

  return
}

// tabifyDockWidget(class QDockWidget *, class QDockWidget *)
func (this *QMainWindow) TabifyDockWidget(args ...interface{}) () {
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
    var arg0 = args[0].(*QDockWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QDockWidget).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QMainWindow16tabifyDockWidgetEP11QDockWidgetS1_(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMainWindow", "tabifyDockWidget", args)
  }

  return
}

// toolBarBreak(class QToolBar *)
func (this *QMainWindow) ToolBarBreak(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QToolBar).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QMainWindow12toolBarBreakEP8QToolBar(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMainWindow", "toolBarBreak", args)
  }

  return
}

// documentMode()
func (this *QMainWindow) DocumentMode(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QMainWindow12documentModeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMainWindow", "documentMode", args)
  }

  return
}

// insertToolBarBreak(class QToolBar *)
func (this *QMainWindow) InsertToolBarBreak(args ...interface{}) () {
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
    var arg0 = args[0].(*QToolBar).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMainWindow18insertToolBarBreakEP8QToolBar(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "insertToolBarBreak", args)
  }

  return
}

// restoreState(const class QByteArray &, int)
func (this *QMainWindow) RestoreState(args ...interface{}) (ret interface{}) {
  // restoreState(const class QByteArray &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QByteArray{}) // "const QByteArray &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMainWindow12restoreStateERK10QByteArrayi
    // invoke: bool restoreState(const class QByteArray &, int)
    var arg0 = args[0].(*qtcore.QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN11QMainWindow12restoreStateERK10QByteArrayi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMainWindow", "restoreState", args)
  }

  return
}

// restoreDockWidget(class QDockWidget *)
func (this *QMainWindow) RestoreDockWidget(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QDockWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN11QMainWindow17restoreDockWidgetEP11QDockWidget(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMainWindow", "restoreDockWidget", args)
  }

  return
}

// statusBar()
func (this *QMainWindow) StatusBar(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QMainWindow9statusBarEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStatusBar{}) // "QStatusBar *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMainWindow", "statusBar", args)
  }

  return
}

// dockOptions()
func (this *QMainWindow) DockOptions(args ...interface{}) () {
  // dockOptions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMainWindow11dockOptionsEv
    // invoke: DockOptions dockOptions()
    C.C_ZNK11QMainWindow11dockOptionsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMainWindow", "dockOptions", args)
  }

  return
}

// removeToolBarBreak(class QToolBar *)
func (this *QMainWindow) RemoveToolBarBreak(args ...interface{}) () {
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
    var arg0 = args[0].(*QToolBar).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMainWindow18removeToolBarBreakEP8QToolBar(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "removeToolBarBreak", args)
  }

  return
}

// setAnimated(_Bool)
func (this *QMainWindow) SetAnimated(args ...interface{}) () {
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QMainWindow11setAnimatedEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "setAnimated", args)
  }

  return
}

// metaObject()
func (this *QMainWindow) MetaObject(args ...interface{}) () {
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
    C.C_ZNK11QMainWindow10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMainWindow", "metaObject", args)
  }

  return
}

// setDocumentMode(_Bool)
func (this *QMainWindow) SetDocumentMode(args ...interface{}) () {
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QMainWindow15setDocumentModeEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "setDocumentMode", args)
  }

  return
}

// takeCentralWidget()
func (this *QMainWindow) TakeCentralWidget(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN11QMainWindow17takeCentralWidgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMainWindow", "takeCentralWidget", args)
  }

  return
}

// createPopupMenu()
func (this *QMainWindow) CreatePopupMenu(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN11QMainWindow15createPopupMenuEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMenu{}) // "QMenu *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMainWindow", "createPopupMenu", args)
  }

  return
}

// setDockNestingEnabled(_Bool)
func (this *QMainWindow) SetDockNestingEnabled(args ...interface{}) () {
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QMainWindow21setDockNestingEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "setDockNestingEnabled", args)
  }

  return
}

// centralWidget()
func (this *QMainWindow) CentralWidget(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QMainWindow13centralWidgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMainWindow", "centralWidget", args)
  }

  return
}

// setIconSize(const class QSize &)
func (this *QMainWindow) SetIconSize(args ...interface{}) () {
  // setIconSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMainWindow11setIconSizeERK5QSize
    // invoke: void setIconSize(const class QSize &)
    var arg0 = args[0].(*qtcore.QSize).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMainWindow11setIconSizeERK5QSize(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "setIconSize", args)
  }

  return
}

// setStatusBar(class QStatusBar *)
func (this *QMainWindow) SetStatusBar(args ...interface{}) () {
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
    var arg0 = args[0].(*QStatusBar).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMainWindow12setStatusBarEP10QStatusBar(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMainWindow", "setStatusBar", args)
  }

  return
}

// isSeparator(const class QPoint &)
func (this *QMainWindow) IsSeparator(args ...interface{}) (ret interface{}) {
  // isSeparator(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMainWindow11isSeparatorERK6QPoint
    // invoke: bool isSeparator(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QMainWindow11isSeparatorERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMainWindow", "isSeparator", args)
  }

  return
}

// menuWidget()
func (this *QMainWindow) MenuWidget(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QMainWindow10menuWidgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMainWindow", "menuWidget", args)
  }

  return
}

// toolButtonStyle()
func (this *QMainWindow) ToolButtonStyle(args ...interface{}) () {
  // toolButtonStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMainWindow15toolButtonStyleEv
    // invoke: Qt::ToolButtonStyle toolButtonStyle()
    C.C_ZNK11QMainWindow15toolButtonStyleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMainWindow", "toolButtonStyle", args)
  }

  return
}

// ~QMainWindow()
func (this *QMainWindow) Free(args ...interface{}) () {
  // ~QMainWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMainWindowD0Ev
    // invoke: void ~QMainWindow()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN11QMainWindowD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QMainWindow", "~QMainWindow", args)
  }

  return
}

// isDockNestingEnabled()
func (this *QMainWindow) IsDockNestingEnabled(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QMainWindow20isDockNestingEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMainWindow", "isDockNestingEnabled", args)
  }

  return
}

// menuBar()
func (this *QMainWindow) MenuBar(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QMainWindow7menuBarEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMenuBar{}) // "QMenuBar *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMainWindow", "menuBar", args)
  }

  return
}

// <= body block end

