package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:13 2016
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
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QMdiArea::tileSubWindows();
extern void _ZN8QMdiArea14tileSubWindowsEv(void* qthis); // 4
  // proto:  void QMdiArea::closeAllSubWindows();
extern void _ZN8QMdiArea18closeAllSubWindowsEv(void* qthis); // 4
  // proto:  void QMdiArea::cascadeSubWindows();
extern void _ZN8QMdiArea17cascadeSubWindowsEv(void* qthis); // 4
  // proto:  QTabWidget::TabPosition QMdiArea::tabPosition();
extern void _ZNK8QMdiArea11tabPositionEv(void* qthis); // 4
  // proto:  QMdiArea::WindowOrder QMdiArea::activationOrder();
extern void _ZNK8QMdiArea15activationOrderEv(void* qthis); // 4
  // proto:  void QMdiArea::setTabsMovable(bool movable);
extern void _ZN8QMdiArea14setTabsMovableEb(void* qthis, bool arg0); // 4
  // proto:  bool QMdiArea::tabsClosable();
extern void _ZNK8QMdiArea12tabsClosableEv(void* qthis); // 4
  // proto:  void QMdiArea::closeActiveSubWindow();
extern void _ZN8QMdiArea20closeActiveSubWindowEv(void* qthis); // 4
  // proto:  void QMdiArea::setDocumentMode(bool enabled);
extern void _ZN8QMdiArea15setDocumentModeEb(void* qthis, bool arg0); // 4
  // proto:  QTabWidget::TabShape QMdiArea::tabShape();
extern void _ZNK8QMdiArea8tabShapeEv(void* qthis); // 4
  // proto:  void QMdiArea::setBackground(const QBrush & background);
extern void _ZN8QMdiArea13setBackgroundERK6QBrush(void* qthis, void* arg0); // 4
  // proto:  void QMdiArea::~QMdiArea();
extern void _ZN8QMdiAreaD2Ev(void* qthis); // 4
  // proto:  bool QMdiArea::documentMode();
extern void _ZNK8QMdiArea12documentModeEv(void* qthis); // 4
  // proto:  bool QMdiArea::tabsMovable();
extern void _ZNK8QMdiArea11tabsMovableEv(void* qthis); // 4
  // proto:  void QMdiArea::setTabsClosable(bool closable);
extern void _ZN8QMdiArea15setTabsClosableEb(void* qthis, bool arg0); // 4
  // proto:  void QMdiArea::QMdiArea(QWidget * parent);
extern void _ZN8QMdiAreaC2EP7QWidget(void* qthis, void* arg0); // 3
  // proto:  QBrush QMdiArea::background();
extern void _ZNK8QMdiArea10backgroundEv(void* qthis); // 4
  // proto:  QSize QMdiArea::sizeHint();
extern void _ZNK8QMdiArea8sizeHintEv(void* qthis); // 4
  // proto:  void QMdiArea::removeSubWindow(QWidget * widget);
extern void _ZN8QMdiArea15removeSubWindowEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QMdiArea::setActiveSubWindow(QMdiSubWindow * window);
extern void _ZN8QMdiArea18setActiveSubWindowEP13QMdiSubWindow(void* qthis, void* arg0); // 4
  // proto:  QMdiSubWindow * QMdiArea::currentSubWindow();
extern void _ZNK8QMdiArea16currentSubWindowEv(void* qthis); // 4
  // proto:  QMdiSubWindow * QMdiArea::activeSubWindow();
extern void _ZNK8QMdiArea15activeSubWindowEv(void* qthis); // 4
  // proto:  void QMdiArea::activateNextSubWindow();
extern void _ZN8QMdiArea21activateNextSubWindowEv(void* qthis); // 4
  // proto:  const QMetaObject * QMdiArea::metaObject();
extern void _ZNK8QMdiArea10metaObjectEv(void* qthis); // 4
  // proto:  QSize QMdiArea::minimumSizeHint();
extern void _ZNK8QMdiArea15minimumSizeHintEv(void* qthis); // 4
  // proto:  QMdiArea::ViewMode QMdiArea::viewMode();
extern void _ZNK8QMdiArea8viewModeEv(void* qthis); // 4
  // proto:  void QMdiArea::activatePreviousSubWindow();
extern void _ZN8QMdiArea25activatePreviousSubWindowEv(void* qthis); // 4
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

// class sizeof(QMdiArea)=1
type QMdiArea struct {
  /*qbase*/ QAbstractScrollArea;
  qclsinst unsafe.Pointer /* *C.void */;
//  _subWindowActivated QMdiArea_subWindowActivated_signal;
}

// tileSubWindows()
func (this *QMdiArea) tileSubWindows(args ...interface{}) () {
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
    C._ZN8QMdiArea14tileSubWindowsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "tileSubWindows", args)
  }

}

// closeAllSubWindows()
func (this *QMdiArea) closeAllSubWindows(args ...interface{}) () {
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
    C._ZN8QMdiArea18closeAllSubWindowsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "closeAllSubWindows", args)
  }

}

// cascadeSubWindows()
func (this *QMdiArea) cascadeSubWindows(args ...interface{}) () {
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
    C._ZN8QMdiArea17cascadeSubWindowsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "cascadeSubWindows", args)
  }

}

// tabPosition()
func (this *QMdiArea) tabPosition(args ...interface{}) () {
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
    C._ZNK8QMdiArea11tabPositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "tabPosition", args)
  }

}

// activationOrder()
func (this *QMdiArea) activationOrder(args ...interface{}) () {
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
    C._ZNK8QMdiArea15activationOrderEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "activationOrder", args)
  }

}

// setTabsMovable(_Bool)
func (this *QMdiArea) setTabsMovable(args ...interface{}) () {
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
    C._ZN8QMdiArea14setTabsMovableEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMdiArea", "setTabsMovable", args)
  }

}

// tabsClosable()
func (this *QMdiArea) tabsClosable(args ...interface{}) () {
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
    C._ZNK8QMdiArea12tabsClosableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "tabsClosable", args)
  }

}

// closeActiveSubWindow()
func (this *QMdiArea) closeActiveSubWindow(args ...interface{}) () {
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
    C._ZN8QMdiArea20closeActiveSubWindowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "closeActiveSubWindow", args)
  }

}

// setDocumentMode(_Bool)
func (this *QMdiArea) setDocumentMode(args ...interface{}) () {
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
    C._ZN8QMdiArea15setDocumentModeEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMdiArea", "setDocumentMode", args)
  }

}

// tabShape()
func (this *QMdiArea) tabShape(args ...interface{}) () {
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
    C._ZNK8QMdiArea8tabShapeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "tabShape", args)
  }

}

// setBackground(const class QBrush &)
func (this *QMdiArea) setBackground(args ...interface{}) () {
  // setBackground(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMdiArea13setBackgroundERK6QBrush
    // invoke: void setBackground(const class QBrush &)
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QMdiArea13setBackgroundERK6QBrush(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMdiArea", "setBackground", args)
  }

}

// ~QMdiArea()
func (this *QMdiArea) FreeQMdiArea(args ...interface{}) () {
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
    C._ZN8QMdiAreaD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "~QMdiArea", args)
  }

}

// documentMode()
func (this *QMdiArea) documentMode(args ...interface{}) () {
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
    C._ZNK8QMdiArea12documentModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "documentMode", args)
  }

}

// tabsMovable()
func (this *QMdiArea) tabsMovable(args ...interface{}) () {
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
    C._ZNK8QMdiArea11tabsMovableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "tabsMovable", args)
  }

}

// setTabsClosable(_Bool)
func (this *QMdiArea) setTabsClosable(args ...interface{}) () {
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
    C._ZN8QMdiArea15setTabsClosableEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMdiArea", "setTabsClosable", args)
  }

}

// QMdiArea(class QWidget *)
func NewQMdiArea(args ...interface{}) QMdiArea {
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
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QMdiAreaC2EP7QWidget(qthis, arg0)
  default:
    qtrt.ErrorResolve("QMdiArea", "QMdiArea", args)
  }

  return QMdiArea{}
}

// background()
func (this *QMdiArea) background(args ...interface{}) () {
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
    C._ZNK8QMdiArea10backgroundEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "background", args)
  }

}

// sizeHint()
func (this *QMdiArea) sizeHint(args ...interface{}) () {
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
    C._ZNK8QMdiArea8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "sizeHint", args)
  }

}

// removeSubWindow(class QWidget *)
func (this *QMdiArea) removeSubWindow(args ...interface{}) () {
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
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QMdiArea15removeSubWindowEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMdiArea", "removeSubWindow", args)
  }

}

// setActiveSubWindow(class QMdiSubWindow *)
func (this *QMdiArea) setActiveSubWindow(args ...interface{}) () {
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
    var arg0 = args[0].(QMdiSubWindow).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QMdiArea18setActiveSubWindowEP13QMdiSubWindow(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMdiArea", "setActiveSubWindow", args)
  }

}

// currentSubWindow()
func (this *QMdiArea) currentSubWindow(args ...interface{}) () {
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
    C._ZNK8QMdiArea16currentSubWindowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "currentSubWindow", args)
  }

}

// activeSubWindow()
func (this *QMdiArea) activeSubWindow(args ...interface{}) () {
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
    C._ZNK8QMdiArea15activeSubWindowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "activeSubWindow", args)
  }

}

// activateNextSubWindow()
func (this *QMdiArea) activateNextSubWindow(args ...interface{}) () {
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
    C._ZN8QMdiArea21activateNextSubWindowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "activateNextSubWindow", args)
  }

}

// metaObject()
func (this *QMdiArea) metaObject(args ...interface{}) () {
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
    C._ZNK8QMdiArea10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "metaObject", args)
  }

}

// minimumSizeHint()
func (this *QMdiArea) minimumSizeHint(args ...interface{}) () {
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
    C._ZNK8QMdiArea15minimumSizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "minimumSizeHint", args)
  }

}

// viewMode()
func (this *QMdiArea) viewMode(args ...interface{}) () {
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
    C._ZNK8QMdiArea8viewModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "viewMode", args)
  }

}

// activatePreviousSubWindow()
func (this *QMdiArea) activatePreviousSubWindow(args ...interface{}) () {
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
    C._ZN8QMdiArea25activatePreviousSubWindowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiArea", "activatePreviousSubWindow", args)
  }

}

// <= body block end

