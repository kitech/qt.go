package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
  // proto:  void QDockWidget::setWidget(QWidget * widget);
extern void C_ZN11QDockWidget9setWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  bool QDockWidget::isFloating();
extern bool C_ZNK11QDockWidget10isFloatingEv(void* qthis); // 2
  // proto:  QAction * QDockWidget::toggleViewAction();
extern void* C_ZNK11QDockWidget16toggleViewActionEv(void* qthis); // 4
  // proto:  QWidget * QDockWidget::widget();
extern void* C_ZNK11QDockWidget6widgetEv(void* qthis); // 4
  // proto:  void QDockWidget::setTitleBarWidget(QWidget * widget);
extern void C_ZN11QDockWidget17setTitleBarWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QDockWidget::~QDockWidget();
extern void C_ZN11QDockWidgetD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QDockWidget::metaObject();
extern void C_ZNK11QDockWidget10metaObjectEv(void* qthis); // 4
  // proto:  Qt::DockWidgetAreas QDockWidget::allowedAreas();
extern void C_ZNK11QDockWidget12allowedAreasEv(void* qthis); // 4
  // proto:  QWidget * QDockWidget::titleBarWidget();
extern void* C_ZNK11QDockWidget14titleBarWidgetEv(void* qthis); // 4
  // proto:  void QDockWidget::setFloating(bool floating);
extern void C_ZN11QDockWidget11setFloatingEb(void* qthis, bool arg0); // 4
  // proto:  DockWidgetFeatures QDockWidget::features();
extern void C_ZNK11QDockWidget8featuresEv(void* qthis); // 4
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

// class sizeof(QDockWidget)=1
type QDockWidget struct {
  /*qbase*/ QWidget;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _featuresChanged QDockWidget_featuresChanged_signal;
//  _visibilityChanged QDockWidget_visibilityChanged_signal;
//  _topLevelChanged QDockWidget_topLevelChanged_signal;
//  _allowedAreasChanged QDockWidget_allowedAreasChanged_signal;
//  _dockLocationChanged QDockWidget_dockLocationChanged_signal;
}

// setWidget(class QWidget *)
func (this *QDockWidget) SetWidget(args ...interface{}) () {
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QDockWidget9setWidgetEP7QWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDockWidget", "setWidget", args)
  }

  return
}

// isFloating()
func (this *QDockWidget) IsFloating(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QDockWidget10isFloatingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDockWidget", "isFloating", args)
  }

  return
}

// toggleViewAction()
func (this *QDockWidget) ToggleViewAction(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QDockWidget16toggleViewActionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDockWidget", "toggleViewAction", args)
  }

  return
}

// widget()
func (this *QDockWidget) Widget(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QDockWidget6widgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDockWidget", "widget", args)
  }

  return
}

// setTitleBarWidget(class QWidget *)
func (this *QDockWidget) SetTitleBarWidget(args ...interface{}) () {
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QDockWidget17setTitleBarWidgetEP7QWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDockWidget", "setTitleBarWidget", args)
  }

  return
}

// ~QDockWidget()
func (this *QDockWidget) Free(args ...interface{}) () {
  // ~QDockWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QDockWidgetD0Ev
    // invoke: void ~QDockWidget()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN11QDockWidgetD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QDockWidget", "~QDockWidget", args)
  }

  return
}

// metaObject()
func (this *QDockWidget) MetaObject(args ...interface{}) () {
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
    C.C_ZNK11QDockWidget10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDockWidget", "metaObject", args)
  }

  return
}

// allowedAreas()
func (this *QDockWidget) AllowedAreas(args ...interface{}) () {
  // allowedAreas()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QDockWidget12allowedAreasEv
    // invoke: Qt::DockWidgetAreas allowedAreas()
    C.C_ZNK11QDockWidget12allowedAreasEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDockWidget", "allowedAreas", args)
  }

  return
}

// titleBarWidget()
func (this *QDockWidget) TitleBarWidget(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QDockWidget14titleBarWidgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDockWidget", "titleBarWidget", args)
  }

  return
}

// setFloating(_Bool)
func (this *QDockWidget) SetFloating(args ...interface{}) () {
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
    C.C_ZN11QDockWidget11setFloatingEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDockWidget", "setFloating", args)
  }

  return
}

// features()
func (this *QDockWidget) Features(args ...interface{}) () {
  // features()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QDockWidget8featuresEv
    // invoke: DockWidgetFeatures features()
    C.C_ZNK11QDockWidget8featuresEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDockWidget", "features", args)
  }

  return
}

// <= body block end

