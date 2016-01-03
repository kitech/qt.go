package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
// src-file: /QtWidgets/qdockwidget.h
// dst-file: /src/widgets/qdockwidget.go
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
  // proto:  QWidget * QDockWidget::widget();
extern void _ZNK11QDockWidget6widgetEv(void* qthis);
  // proto:  void QDockWidget::setFloating(bool floating);
extern void _ZN11QDockWidget11setFloatingEb(void* qthis, bool arg0);
  // proto:  QWidget * QDockWidget::titleBarWidget();
extern void _ZNK11QDockWidget14titleBarWidgetEv(void* qthis);
  // proto:  void QDockWidget::~QDockWidget();
extern void _ZN11QDockWidgetD0Ev(void* qthis);
  // proto:  void QDockWidget::setWidget(QWidget * widget);
extern void _ZN11QDockWidget9setWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  bool QDockWidget::isFloating();
extern void demth_ZNK11QDockWidget10isFloatingEv(void* qthis);
  // proto:  QAction * QDockWidget::toggleViewAction();
extern void _ZNK11QDockWidget16toggleViewActionEv(void* qthis);
  // proto:  void QDockWidget::QDockWidget(const QDockWidget & );
extern void* dector_ZN11QDockWidgetC1ERKS_(void* arg0);
extern void _ZN11QDockWidgetC1ERKS_(void* qthis, void* arg0);
  // proto:  void QDockWidget::setTitleBarWidget(QWidget * widget);
extern void _ZN11QDockWidget17setTitleBarWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  const QMetaObject * QDockWidget::metaObject();
extern void _ZNK11QDockWidget10metaObjectEv(void* qthis);
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

// class sizeof(QDockWidget)=1
type QDockWidget struct {
  /*qbase*/ QWidget;
  qclsinst unsafe.Pointer /* *C.void */;
//  _featuresChanged QDockWidget_featuresChanged_signal;
//  _visibilityChanged QDockWidget_visibilityChanged_signal;
//  _topLevelChanged QDockWidget_topLevelChanged_signal;
//  _allowedAreasChanged QDockWidget_allowedAreasChanged_signal;
//  _dockLocationChanged QDockWidget_dockLocationChanged_signal;
}

  // proto:  QWidget * QDockWidget::widget();
func (this *QDockWidget) widget(args ...interface{}) () {
  // widget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QDockWidget6widgetEv
    // invoke: QWidget * widget()
    C._ZNK11QDockWidget6widgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDockWidget", "widget", args)
  }

}

  // proto:  void QDockWidget::setFloating(bool floating);
func (this *QDockWidget) setFloating(args ...interface{}) () {
  // setFloating(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QDockWidget11setFloatingEb
    // invoke: void setFloating(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN11QDockWidget11setFloatingEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDockWidget", "setFloating", args)
  }

}

  // proto:  QWidget * QDockWidget::titleBarWidget();
func (this *QDockWidget) titleBarWidget(args ...interface{}) () {
  // titleBarWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QDockWidget14titleBarWidgetEv
    // invoke: QWidget * titleBarWidget()
    C._ZNK11QDockWidget14titleBarWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDockWidget", "titleBarWidget", args)
  }

}

  // proto:  void QDockWidget::~QDockWidget();
func (this *QDockWidget) FreeQDockWidget(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDockWidget", "~QDockWidget", args)
  }

}

  // proto:  void QDockWidget::setWidget(QWidget * widget);
func (this *QDockWidget) setWidget(args ...interface{}) () {
  // setWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QDockWidget9setWidgetEP7QWidget
    // invoke: void setWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QDockWidget9setWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDockWidget", "setWidget", args)
  }

}

  // proto:  bool QDockWidget::isFloating();
func (this *QDockWidget) isFloating(args ...interface{}) () {
  // isFloating()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QDockWidget10isFloatingEv
    // invoke: bool isFloating()
    C.demth_ZNK11QDockWidget10isFloatingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDockWidget", "isFloating", args)
  }

}

  // proto:  QAction * QDockWidget::toggleViewAction();
func (this *QDockWidget) toggleViewAction(args ...interface{}) () {
  // toggleViewAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QDockWidget16toggleViewActionEv
    // invoke: QAction * toggleViewAction()
    C._ZNK11QDockWidget16toggleViewActionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDockWidget", "toggleViewAction", args)
  }

}

  // proto:  void QDockWidget::QDockWidget(const QDockWidget & );
func NewQDockWidget(args ...interface{}) QDockWidget {
  return QDockWidget{}
}

  // proto:  void QDockWidget::setTitleBarWidget(QWidget * widget);
func (this *QDockWidget) setTitleBarWidget(args ...interface{}) () {
  // setTitleBarWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QDockWidget17setTitleBarWidgetEP7QWidget
    // invoke: void setTitleBarWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QDockWidget17setTitleBarWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDockWidget", "setTitleBarWidget", args)
  }

}

  // proto:  const QMetaObject * QDockWidget::metaObject();
func (this *QDockWidget) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QDockWidget10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK11QDockWidget10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDockWidget", "metaObject", args)
  }

}

// <= body block end

