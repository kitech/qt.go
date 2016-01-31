package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtWidgets/qmdisubwindow.h
// dst-file: /src/widgets/qmdisubwindow.go
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
  // proto:  void QMdiSubWindow::setKeyboardPageStep(int step);
extern void C_ZN13QMdiSubWindow19setKeyboardPageStepEi(void* qthis, int32_t arg0); // 4
  // proto:  void QMdiSubWindow::~QMdiSubWindow();
extern void C_ZN13QMdiSubWindowD2Ev(void* qthis); // 4
  // proto:  void QMdiSubWindow::setWidget(QWidget * widget);
extern void C_ZN13QMdiSubWindow9setWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  QWidget * QMdiSubWindow::maximizedButtonsWidget();
extern void C_ZNK13QMdiSubWindow22maximizedButtonsWidgetEv(void* qthis); // 4
  // proto:  QMdiArea * QMdiSubWindow::mdiArea();
extern void C_ZNK13QMdiSubWindow7mdiAreaEv(void* qthis); // 4
  // proto:  void QMdiSubWindow::setSystemMenu(QMenu * systemMenu);
extern void C_ZN13QMdiSubWindow13setSystemMenuEP5QMenu(void* qthis, void* arg0); // 4
  // proto:  int QMdiSubWindow::keyboardPageStep();
extern void C_ZNK13QMdiSubWindow16keyboardPageStepEv(void* qthis); // 4
  // proto:  void QMdiSubWindow::setKeyboardSingleStep(int step);
extern void C_ZN13QMdiSubWindow21setKeyboardSingleStepEi(void* qthis, int32_t arg0); // 4
  // proto:  void QMdiSubWindow::showShaded();
extern void C_ZN13QMdiSubWindow10showShadedEv(void* qthis); // 4
  // proto:  QWidget * QMdiSubWindow::widget();
extern void C_ZNK13QMdiSubWindow6widgetEv(void* qthis); // 4
  // proto:  int QMdiSubWindow::keyboardSingleStep();
extern void C_ZNK13QMdiSubWindow18keyboardSingleStepEv(void* qthis); // 4
  // proto:  QMenu * QMdiSubWindow::systemMenu();
extern void C_ZNK13QMdiSubWindow10systemMenuEv(void* qthis); // 4
  // proto:  const QMetaObject * QMdiSubWindow::metaObject();
extern void C_ZNK13QMdiSubWindow10metaObjectEv(void* qthis); // 4
  // proto:  bool QMdiSubWindow::isShaded();
extern void C_ZNK13QMdiSubWindow8isShadedEv(void* qthis); // 4
  // proto:  QSize QMdiSubWindow::sizeHint();
extern void C_ZNK13QMdiSubWindow8sizeHintEv(void* qthis); // 4
  // proto:  QSize QMdiSubWindow::minimumSizeHint();
extern void C_ZNK13QMdiSubWindow15minimumSizeHintEv(void* qthis); // 4
  // proto:  void QMdiSubWindow::showSystemMenu();
extern void C_ZN13QMdiSubWindow14showSystemMenuEv(void* qthis); // 4
  // proto:  QWidget * QMdiSubWindow::maximizedSystemMenuIconWidget();
extern void C_ZNK13QMdiSubWindow29maximizedSystemMenuIconWidgetEv(void* qthis); // 4
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

// class sizeof(QMdiSubWindow)=1
type QMdiSubWindow struct {
  /*qbase*/ QWidget;
  qclsinst unsafe.Pointer /* *C.void */;
//  _aboutToActivate QMdiSubWindow_aboutToActivate_signal;
//  _windowStateChanged QMdiSubWindow_windowStateChanged_signal;
}

// setKeyboardPageStep(int)
func (this *QMdiSubWindow) setKeyboardPageStep(args ...interface{}) () {
  // setKeyboardPageStep(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QMdiSubWindow19setKeyboardPageStepEi
    // invoke: void setKeyboardPageStep(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN13QMdiSubWindow19setKeyboardPageStepEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "setKeyboardPageStep", args)
  }

}

// ~QMdiSubWindow()
func (this *QMdiSubWindow) FreeQMdiSubWindow(args ...interface{}) () {
  // ~QMdiSubWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QMdiSubWindowD0Ev
    // invoke: void ~QMdiSubWindow()
    C.C_ZN13QMdiSubWindowD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "~QMdiSubWindow", args)
  }

}

// setWidget(class QWidget *)
func (this *QMdiSubWindow) setWidget(args ...interface{}) () {
  // setWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QMdiSubWindow9setWidgetEP7QWidget
    // invoke: void setWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QMdiSubWindow9setWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "setWidget", args)
  }

}

// maximizedButtonsWidget()
func (this *QMdiSubWindow) maximizedButtonsWidget(args ...interface{}) () {
  // maximizedButtonsWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow22maximizedButtonsWidgetEv
    // invoke: QWidget * maximizedButtonsWidget()
    C.C_ZNK13QMdiSubWindow22maximizedButtonsWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "maximizedButtonsWidget", args)
  }

}

// mdiArea()
func (this *QMdiSubWindow) mdiArea(args ...interface{}) () {
  // mdiArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow7mdiAreaEv
    // invoke: QMdiArea * mdiArea()
    C.C_ZNK13QMdiSubWindow7mdiAreaEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "mdiArea", args)
  }

}

// setSystemMenu(class QMenu *)
func (this *QMdiSubWindow) setSystemMenu(args ...interface{}) () {
  // setSystemMenu(class QMenu *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMenu{}) // "QMenu *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QMdiSubWindow13setSystemMenuEP5QMenu
    // invoke: void setSystemMenu(class QMenu *)
    var arg0 = args[0].(QMenu).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QMdiSubWindow13setSystemMenuEP5QMenu(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "setSystemMenu", args)
  }

}

// keyboardPageStep()
func (this *QMdiSubWindow) keyboardPageStep(args ...interface{}) () {
  // keyboardPageStep()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow16keyboardPageStepEv
    // invoke: int keyboardPageStep()
    C.C_ZNK13QMdiSubWindow16keyboardPageStepEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "keyboardPageStep", args)
  }

}

// setKeyboardSingleStep(int)
func (this *QMdiSubWindow) setKeyboardSingleStep(args ...interface{}) () {
  // setKeyboardSingleStep(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QMdiSubWindow21setKeyboardSingleStepEi
    // invoke: void setKeyboardSingleStep(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN13QMdiSubWindow21setKeyboardSingleStepEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "setKeyboardSingleStep", args)
  }

}

// showShaded()
func (this *QMdiSubWindow) showShaded(args ...interface{}) () {
  // showShaded()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QMdiSubWindow10showShadedEv
    // invoke: void showShaded()
    C.C_ZN13QMdiSubWindow10showShadedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "showShaded", args)
  }

}

// widget()
func (this *QMdiSubWindow) widget(args ...interface{}) () {
  // widget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow6widgetEv
    // invoke: QWidget * widget()
    C.C_ZNK13QMdiSubWindow6widgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "widget", args)
  }

}

// keyboardSingleStep()
func (this *QMdiSubWindow) keyboardSingleStep(args ...interface{}) () {
  // keyboardSingleStep()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow18keyboardSingleStepEv
    // invoke: int keyboardSingleStep()
    C.C_ZNK13QMdiSubWindow18keyboardSingleStepEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "keyboardSingleStep", args)
  }

}

// systemMenu()
func (this *QMdiSubWindow) systemMenu(args ...interface{}) () {
  // systemMenu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow10systemMenuEv
    // invoke: QMenu * systemMenu()
    C.C_ZNK13QMdiSubWindow10systemMenuEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "systemMenu", args)
  }

}

// metaObject()
func (this *QMdiSubWindow) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK13QMdiSubWindow10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "metaObject", args)
  }

}

// isShaded()
func (this *QMdiSubWindow) isShaded(args ...interface{}) () {
  // isShaded()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow8isShadedEv
    // invoke: bool isShaded()
    C.C_ZNK13QMdiSubWindow8isShadedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "isShaded", args)
  }

}

// sizeHint()
func (this *QMdiSubWindow) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow8sizeHintEv
    // invoke: QSize sizeHint()
    C.C_ZNK13QMdiSubWindow8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "sizeHint", args)
  }

}

// minimumSizeHint()
func (this *QMdiSubWindow) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    C.C_ZNK13QMdiSubWindow15minimumSizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "minimumSizeHint", args)
  }

}

// showSystemMenu()
func (this *QMdiSubWindow) showSystemMenu(args ...interface{}) () {
  // showSystemMenu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QMdiSubWindow14showSystemMenuEv
    // invoke: void showSystemMenu()
    C.C_ZN13QMdiSubWindow14showSystemMenuEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "showSystemMenu", args)
  }

}

// maximizedSystemMenuIconWidget()
func (this *QMdiSubWindow) maximizedSystemMenuIconWidget(args ...interface{}) () {
  // maximizedSystemMenuIconWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow29maximizedSystemMenuIconWidgetEv
    // invoke: QWidget * maximizedSystemMenuIconWidget()
    C.C_ZNK13QMdiSubWindow29maximizedSystemMenuIconWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "maximizedSystemMenuIconWidget", args)
  }

}

// <= body block end

