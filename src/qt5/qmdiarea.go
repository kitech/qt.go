package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
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
  // proto:  void QMdiArea::activateNextSubWindow();
extern void _ZN8QMdiArea21activateNextSubWindowEv(void* qthis);
  // proto:  void QMdiArea::setBackground(const QBrush & background);
extern void _ZN8QMdiArea13setBackgroundERK6QBrush(void* qthis, void* arg0);
  // proto:  void QMdiArea::~QMdiArea();
extern void _ZN8QMdiAreaD0Ev(void* qthis);
  // proto:  void QMdiArea::removeSubWindow(QWidget * widget);
extern void _ZN8QMdiArea15removeSubWindowEP7QWidget(void* qthis, void* arg0);
  // proto:  void QMdiArea::setTabsClosable(bool closable);
extern void _ZN8QMdiArea15setTabsClosableEb(void* qthis, bool arg0);
  // proto:  QMdiSubWindow * QMdiArea::currentSubWindow();
extern void _ZNK8QMdiArea16currentSubWindowEv(void* qthis);
  // proto:  bool QMdiArea::tabsMovable();
extern void _ZNK8QMdiArea11tabsMovableEv(void* qthis);
  // proto:  void QMdiArea::activatePreviousSubWindow();
extern void _ZN8QMdiArea25activatePreviousSubWindowEv(void* qthis);
  // proto:  void QMdiArea::setDocumentMode(bool enabled);
extern void _ZN8QMdiArea15setDocumentModeEb(void* qthis, bool arg0);
  // proto:  bool QMdiArea::documentMode();
extern void _ZNK8QMdiArea12documentModeEv(void* qthis);
  // proto:  void QMdiArea::setActiveSubWindow(QMdiSubWindow * window);
extern void _ZN8QMdiArea18setActiveSubWindowEP13QMdiSubWindow(void* qthis, void* arg0);
  // proto:  QMdiSubWindow * QMdiArea::activeSubWindow();
extern void _ZNK8QMdiArea15activeSubWindowEv(void* qthis);
  // proto:  void QMdiArea::setTabsMovable(bool movable);
extern void _ZN8QMdiArea14setTabsMovableEb(void* qthis, bool arg0);
  // proto:  const QMetaObject * QMdiArea::metaObject();
extern void _ZNK8QMdiArea10metaObjectEv(void* qthis);
  // proto:  void QMdiArea::QMdiArea(QWidget * parent);
extern void* dector_ZN8QMdiAreaC1EP7QWidget(void* arg0);
extern void _ZN8QMdiAreaC1EP7QWidget(void* qthis, void* arg0);
  // proto:  QSize QMdiArea::sizeHint();
extern void _ZNK8QMdiArea8sizeHintEv(void* qthis);
  // proto:  void QMdiArea::closeAllSubWindows();
extern void _ZN8QMdiArea18closeAllSubWindowsEv(void* qthis);
  // proto:  void QMdiArea::QMdiArea(const QMdiArea & );
extern void* dector_ZN8QMdiAreaC1ERKS_(void* arg0);
extern void _ZN8QMdiAreaC1ERKS_(void* qthis, void* arg0);
  // proto:  void QMdiArea::cascadeSubWindows();
extern void _ZN8QMdiArea17cascadeSubWindowsEv(void* qthis);
  // proto:  void QMdiArea::closeActiveSubWindow();
extern void _ZN8QMdiArea20closeActiveSubWindowEv(void* qthis);
  // proto:  QBrush QMdiArea::background();
extern void _ZNK8QMdiArea10backgroundEv(void* qthis);
  // proto:  void QMdiArea::tileSubWindows();
extern void _ZN8QMdiArea14tileSubWindowsEv(void* qthis);
  // proto:  bool QMdiArea::tabsClosable();
extern void _ZNK8QMdiArea12tabsClosableEv(void* qthis);
  // proto:  QSize QMdiArea::minimumSizeHint();
extern void _ZNK8QMdiArea15minimumSizeHintEv(void* qthis);
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
  qclsinst uint64 /* *mut c_void*/;
//  _subWindowActivated QMdiArea_subWindowActivated_signal;
}

  // proto:  void QMdiArea::activateNextSubWindow();
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

  // proto:  void QMdiArea::setBackground(const QBrush & background);
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

  // proto:  void QMdiArea::~QMdiArea();
func (this *QMdiArea) FreeQMdiArea(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMdiArea", "~QMdiArea", args)
  }

}

  // proto:  void QMdiArea::removeSubWindow(QWidget * widget);
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

  // proto:  void QMdiArea::setTabsClosable(bool closable);
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN8QMdiArea15setTabsClosableEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMdiArea", "setTabsClosable", args)
  }

}

  // proto:  QMdiSubWindow * QMdiArea::currentSubWindow();
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

  // proto:  bool QMdiArea::tabsMovable();
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

  // proto:  void QMdiArea::activatePreviousSubWindow();
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

  // proto:  void QMdiArea::setDocumentMode(bool enabled);
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN8QMdiArea15setDocumentModeEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMdiArea", "setDocumentMode", args)
  }

}

  // proto:  bool QMdiArea::documentMode();
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

  // proto:  void QMdiArea::setActiveSubWindow(QMdiSubWindow * window);
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

  // proto:  QMdiSubWindow * QMdiArea::activeSubWindow();
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

  // proto:  void QMdiArea::setTabsMovable(bool movable);
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN8QMdiArea14setTabsMovableEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMdiArea", "setTabsMovable", args)
  }

}

  // proto:  const QMetaObject * QMdiArea::metaObject();
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

  // proto:  void QMdiArea::QMdiArea(QWidget * parent);
func NewQMdiArea(args ...interface{}) QMdiArea {
  return QMdiArea{}
}

  // proto:  QSize QMdiArea::sizeHint();
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

  // proto:  void QMdiArea::closeAllSubWindows();
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

  // proto:  void QMdiArea::cascadeSubWindows();
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

  // proto:  void QMdiArea::closeActiveSubWindow();
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

  // proto:  QBrush QMdiArea::background();
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

  // proto:  void QMdiArea::tileSubWindows();
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

  // proto:  bool QMdiArea::tabsClosable();
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

  // proto:  QSize QMdiArea::minimumSizeHint();
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

// <= body block end

