package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtWidgets/qmdiarea.h
// dst-file: /src/widgets/qmdiarea.go
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
  // proto:  void QMdiArea::tileSubWindows();
extern void C_ZN8QMdiArea14tileSubWindowsEv(void* qthis); // 4
  // proto:  void QMdiArea::closeAllSubWindows();
extern void C_ZN8QMdiArea18closeAllSubWindowsEv(void* qthis); // 4
  // proto:  void QMdiArea::cascadeSubWindows();
extern void C_ZN8QMdiArea17cascadeSubWindowsEv(void* qthis); // 4
  // proto:  QTabWidget::TabPosition QMdiArea::tabPosition();
extern void C_ZNK8QMdiArea11tabPositionEv(void* qthis); // 4
  // proto:  QMdiArea::WindowOrder QMdiArea::activationOrder();
extern void C_ZNK8QMdiArea15activationOrderEv(void* qthis); // 4
  // proto:  void QMdiArea::setTabsMovable(bool movable);
extern void C_ZN8QMdiArea14setTabsMovableEb(void* qthis, bool arg0); // 4
  // proto:  bool QMdiArea::tabsClosable();
extern bool C_ZNK8QMdiArea12tabsClosableEv(void* qthis); // 4
  // proto:  void QMdiArea::closeActiveSubWindow();
extern void C_ZN8QMdiArea20closeActiveSubWindowEv(void* qthis); // 4
  // proto:  void QMdiArea::setDocumentMode(bool enabled);
extern void C_ZN8QMdiArea15setDocumentModeEb(void* qthis, bool arg0); // 4
  // proto:  QTabWidget::TabShape QMdiArea::tabShape();
extern void C_ZNK8QMdiArea8tabShapeEv(void* qthis); // 4
  // proto:  void QMdiArea::setBackground(const QBrush & background);
extern void C_ZN8QMdiArea13setBackgroundERK6QBrush(void* qthis, void* arg0); // 4
  // proto:  void QMdiArea::~QMdiArea();
extern void C_ZN8QMdiAreaD2Ev(void* qthis); // 4
  // proto:  bool QMdiArea::documentMode();
extern bool C_ZNK8QMdiArea12documentModeEv(void* qthis); // 4
  // proto:  bool QMdiArea::tabsMovable();
extern bool C_ZNK8QMdiArea11tabsMovableEv(void* qthis); // 4
  // proto:  void QMdiArea::setTabsClosable(bool closable);
extern void C_ZN8QMdiArea15setTabsClosableEb(void* qthis, bool arg0); // 4
  // proto:  void QMdiArea::QMdiArea(QWidget * parent);
extern void* C_ZN8QMdiAreaC2EP7QWidget(void* arg0); // 3
  // proto:  QBrush QMdiArea::background();
extern void* C_ZNK8QMdiArea10backgroundEv(void* qthis); // 4
  // proto:  QSize QMdiArea::sizeHint();
extern void* C_ZNK8QMdiArea8sizeHintEv(void* qthis); // 4
  // proto:  void QMdiArea::removeSubWindow(QWidget * widget);
extern void C_ZN8QMdiArea15removeSubWindowEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QMdiArea::setActiveSubWindow(QMdiSubWindow * window);
extern void C_ZN8QMdiArea18setActiveSubWindowEP13QMdiSubWindow(void* qthis, void* arg0); // 4
  // proto:  QMdiSubWindow * QMdiArea::currentSubWindow();
extern void* C_ZNK8QMdiArea16currentSubWindowEv(void* qthis); // 4
  // proto:  QMdiSubWindow * QMdiArea::activeSubWindow();
extern void* C_ZNK8QMdiArea15activeSubWindowEv(void* qthis); // 4
  // proto:  void QMdiArea::activateNextSubWindow();
extern void C_ZN8QMdiArea21activateNextSubWindowEv(void* qthis); // 4
  // proto:  const QMetaObject * QMdiArea::metaObject();
extern void C_ZNK8QMdiArea10metaObjectEv(void* qthis); // 4
  // proto:  QSize QMdiArea::minimumSizeHint();
extern void* C_ZNK8QMdiArea15minimumSizeHintEv(void* qthis); // 4
  // proto:  QMdiArea::ViewMode QMdiArea::viewMode();
extern void C_ZNK8QMdiArea8viewModeEv(void* qthis); // 4
  // proto:  void QMdiArea::activatePreviousSubWindow();
extern void C_ZN8QMdiArea25activatePreviousSubWindowEv(void* qthis); // 4
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

// class sizeof(QMdiArea)=1
type QMdiArea struct {
  /*qbase*/ QAbstractScrollArea;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _subWindowActivated QMdiArea_subWindowActivated_signal;
}

// tileSubWindows()
func (this *QMdiArea) TileSubWindows(args ...interface{}) () {
  // tileSubWindows()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMdiArea14tileSubWindowsEv
    // invoke: void tileSubWindows()
    C.C_ZN8QMdiArea14tileSubWindowsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "tileSubWindows", args)
  }

  return
}

// closeAllSubWindows()
func (this *QMdiArea) CloseAllSubWindows(args ...interface{}) () {
  // closeAllSubWindows()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMdiArea18closeAllSubWindowsEv
    // invoke: void closeAllSubWindows()
    C.C_ZN8QMdiArea18closeAllSubWindowsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "closeAllSubWindows", args)
  }

  return
}

// cascadeSubWindows()
func (this *QMdiArea) CascadeSubWindows(args ...interface{}) () {
  // cascadeSubWindows()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMdiArea17cascadeSubWindowsEv
    // invoke: void cascadeSubWindows()
    C.C_ZN8QMdiArea17cascadeSubWindowsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "cascadeSubWindows", args)
  }

  return
}

// tabPosition()
func (this *QMdiArea) TabPosition(args ...interface{}) () {
  // tabPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMdiArea11tabPositionEv
    // invoke: QTabWidget::TabPosition tabPosition()
    C.C_ZNK8QMdiArea11tabPositionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "tabPosition", args)
  }

  return
}

// activationOrder()
func (this *QMdiArea) ActivationOrder(args ...interface{}) () {
  // activationOrder()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMdiArea15activationOrderEv
    // invoke: QMdiArea::WindowOrder activationOrder()
    C.C_ZNK8QMdiArea15activationOrderEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "activationOrder", args)
  }

  return
}

// setTabsMovable(_Bool)
func (this *QMdiArea) SetTabsMovable(args ...interface{}) () {
  // setTabsMovable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMdiArea14setTabsMovableEb
    // invoke: void setTabsMovable(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN8QMdiArea14setTabsMovableEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMdiArea", "setTabsMovable", args)
  }

  return
}

// tabsClosable()
func (this *QMdiArea) TabsClosable(args ...interface{}) (ret interface{}) {
  // tabsClosable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMdiArea12tabsClosableEv
    // invoke: bool tabsClosable()
    var ret0 = C.C_ZNK8QMdiArea12tabsClosableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMdiArea", "tabsClosable", args)
  }

  return
}

// closeActiveSubWindow()
func (this *QMdiArea) CloseActiveSubWindow(args ...interface{}) () {
  // closeActiveSubWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMdiArea20closeActiveSubWindowEv
    // invoke: void closeActiveSubWindow()
    C.C_ZN8QMdiArea20closeActiveSubWindowEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "closeActiveSubWindow", args)
  }

  return
}

// setDocumentMode(_Bool)
func (this *QMdiArea) SetDocumentMode(args ...interface{}) () {
  // setDocumentMode(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMdiArea15setDocumentModeEb
    // invoke: void setDocumentMode(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN8QMdiArea15setDocumentModeEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMdiArea", "setDocumentMode", args)
  }

  return
}

// tabShape()
func (this *QMdiArea) TabShape(args ...interface{}) () {
  // tabShape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMdiArea8tabShapeEv
    // invoke: QTabWidget::TabShape tabShape()
    C.C_ZNK8QMdiArea8tabShapeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "tabShape", args)
  }

  return
}

// setBackground(const class QBrush &)
func (this *QMdiArea) SetBackground(args ...interface{}) () {
  // setBackground(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMdiArea13setBackgroundERK6QBrush
    // invoke: void setBackground(const class QBrush &)
    var arg0 = args[0].(*qtgui.QBrush).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QMdiArea13setBackgroundERK6QBrush(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMdiArea", "setBackground", args)
  }

  return
}

// ~QMdiArea()
func (this *QMdiArea) Free(args ...interface{}) () {
  // ~QMdiArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMdiAreaD0Ev
    // invoke: void ~QMdiArea()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN8QMdiAreaD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QMdiArea", "~QMdiArea", args)
  }

  return
}

// documentMode()
func (this *QMdiArea) DocumentMode(args ...interface{}) (ret interface{}) {
  // documentMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMdiArea12documentModeEv
    // invoke: bool documentMode()
    var ret0 = C.C_ZNK8QMdiArea12documentModeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMdiArea", "documentMode", args)
  }

  return
}

// tabsMovable()
func (this *QMdiArea) TabsMovable(args ...interface{}) (ret interface{}) {
  // tabsMovable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMdiArea11tabsMovableEv
    // invoke: bool tabsMovable()
    var ret0 = C.C_ZNK8QMdiArea11tabsMovableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMdiArea", "tabsMovable", args)
  }

  return
}

// setTabsClosable(_Bool)
func (this *QMdiArea) SetTabsClosable(args ...interface{}) () {
  // setTabsClosable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMdiArea15setTabsClosableEb
    // invoke: void setTabsClosable(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN8QMdiArea15setTabsClosableEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMdiArea", "setTabsClosable", args)
  }

  return
}

// QMdiArea(class QWidget *)
func GcfreeQMdiArea(this *QMdiArea) {
  qtrt.UniverseFree(this)
}
func NewQMdiArea(args ...interface{}) *QMdiArea {
  // QMdiArea(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMdiAreaC1EP7QWidget
    // invoke: void QMdiArea(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QMdiAreaC2EP7QWidget(arg0)
    this := &QMdiArea{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQMdiArea)
    return this
  default:
    qtrt.ErrorResolve("QMdiArea", "QMdiArea", args)
  }

  return nil // QMdiArea{Qclsinst:qthis}
}

// background()
func (this *QMdiArea) Background(args ...interface{}) (ret interface{}) {
  // background()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMdiArea10backgroundEv
    // invoke: QBrush background()
    var ret0 = C.C_ZNK8QMdiArea10backgroundEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QBrush{}) // "QBrush"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMdiArea", "background", args)
  }

  return
}

// sizeHint()
func (this *QMdiArea) SizeHint(args ...interface{}) (ret interface{}) {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMdiArea8sizeHintEv
    // invoke: QSize sizeHint()
    var ret0 = C.C_ZNK8QMdiArea8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMdiArea", "sizeHint", args)
  }

  return
}

// removeSubWindow(class QWidget *)
func (this *QMdiArea) RemoveSubWindow(args ...interface{}) () {
  // removeSubWindow(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMdiArea15removeSubWindowEP7QWidget
    // invoke: void removeSubWindow(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QMdiArea15removeSubWindowEP7QWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMdiArea", "removeSubWindow", args)
  }

  return
}

// setActiveSubWindow(class QMdiSubWindow *)
func (this *QMdiArea) SetActiveSubWindow(args ...interface{}) () {
  // setActiveSubWindow(class QMdiSubWindow *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMdiSubWindow{}) // "QMdiSubWindow *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMdiArea18setActiveSubWindowEP13QMdiSubWindow
    // invoke: void setActiveSubWindow(class QMdiSubWindow *)
    var arg0 = args[0].(*QMdiSubWindow).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QMdiArea18setActiveSubWindowEP13QMdiSubWindow(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMdiArea", "setActiveSubWindow", args)
  }

  return
}

// currentSubWindow()
func (this *QMdiArea) CurrentSubWindow(args ...interface{}) (ret interface{}) {
  // currentSubWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMdiArea16currentSubWindowEv
    // invoke: QMdiSubWindow * currentSubWindow()
    var ret0 = C.C_ZNK8QMdiArea16currentSubWindowEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMdiSubWindow{}) // "QMdiSubWindow *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMdiArea", "currentSubWindow", args)
  }

  return
}

// activeSubWindow()
func (this *QMdiArea) ActiveSubWindow(args ...interface{}) (ret interface{}) {
  // activeSubWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMdiArea15activeSubWindowEv
    // invoke: QMdiSubWindow * activeSubWindow()
    var ret0 = C.C_ZNK8QMdiArea15activeSubWindowEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMdiSubWindow{}) // "QMdiSubWindow *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMdiArea", "activeSubWindow", args)
  }

  return
}

// activateNextSubWindow()
func (this *QMdiArea) ActivateNextSubWindow(args ...interface{}) () {
  // activateNextSubWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMdiArea21activateNextSubWindowEv
    // invoke: void activateNextSubWindow()
    C.C_ZN8QMdiArea21activateNextSubWindowEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "activateNextSubWindow", args)
  }

  return
}

// metaObject()
func (this *QMdiArea) MetaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMdiArea10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK8QMdiArea10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "metaObject", args)
  }

  return
}

// minimumSizeHint()
func (this *QMdiArea) MinimumSizeHint(args ...interface{}) (ret interface{}) {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMdiArea15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    var ret0 = C.C_ZNK8QMdiArea15minimumSizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMdiArea", "minimumSizeHint", args)
  }

  return
}

// viewMode()
func (this *QMdiArea) ViewMode(args ...interface{}) () {
  // viewMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMdiArea8viewModeEv
    // invoke: QMdiArea::ViewMode viewMode()
    C.C_ZNK8QMdiArea8viewModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "viewMode", args)
  }

  return
}

// activatePreviousSubWindow()
func (this *QMdiArea) ActivatePreviousSubWindow(args ...interface{}) () {
  // activatePreviousSubWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMdiArea25activatePreviousSubWindowEv
    // invoke: void activatePreviousSubWindow()
    C.C_ZN8QMdiArea25activatePreviousSubWindowEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "activatePreviousSubWindow", args)
  }

  return
}

// <= body block end

