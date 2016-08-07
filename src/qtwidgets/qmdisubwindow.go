package qtwidgets
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
  // proto:  void QMdiSubWindow::setKeyboardPageStep(int step);
extern void C_ZN13QMdiSubWindow19setKeyboardPageStepEi(void* qthis, int32_t arg0); // 4
  // proto:  void QMdiSubWindow::~QMdiSubWindow();
extern void C_ZN13QMdiSubWindowD2Ev(void* qthis); // 4
  // proto:  void QMdiSubWindow::setWidget(QWidget * widget);
extern void C_ZN13QMdiSubWindow9setWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  QWidget * QMdiSubWindow::maximizedButtonsWidget();
extern void* C_ZNK13QMdiSubWindow22maximizedButtonsWidgetEv(void* qthis); // 4
  // proto:  QMdiArea * QMdiSubWindow::mdiArea();
extern void* C_ZNK13QMdiSubWindow7mdiAreaEv(void* qthis); // 4
  // proto:  void QMdiSubWindow::setSystemMenu(QMenu * systemMenu);
extern void C_ZN13QMdiSubWindow13setSystemMenuEP5QMenu(void* qthis, void* arg0); // 4
  // proto:  int QMdiSubWindow::keyboardPageStep();
extern int32_t C_ZNK13QMdiSubWindow16keyboardPageStepEv(void* qthis); // 4
  // proto:  void QMdiSubWindow::setKeyboardSingleStep(int step);
extern void C_ZN13QMdiSubWindow21setKeyboardSingleStepEi(void* qthis, int32_t arg0); // 4
  // proto:  void QMdiSubWindow::showShaded();
extern void C_ZN13QMdiSubWindow10showShadedEv(void* qthis); // 4
  // proto:  QWidget * QMdiSubWindow::widget();
extern void* C_ZNK13QMdiSubWindow6widgetEv(void* qthis); // 4
  // proto:  int QMdiSubWindow::keyboardSingleStep();
extern int32_t C_ZNK13QMdiSubWindow18keyboardSingleStepEv(void* qthis); // 4
  // proto:  QMenu * QMdiSubWindow::systemMenu();
extern void* C_ZNK13QMdiSubWindow10systemMenuEv(void* qthis); // 4
  // proto:  const QMetaObject * QMdiSubWindow::metaObject();
extern void C_ZNK13QMdiSubWindow10metaObjectEv(void* qthis); // 4
  // proto:  bool QMdiSubWindow::isShaded();
extern bool C_ZNK13QMdiSubWindow8isShadedEv(void* qthis); // 4
  // proto:  QSize QMdiSubWindow::sizeHint();
extern void* C_ZNK13QMdiSubWindow8sizeHintEv(void* qthis); // 4
  // proto:  QSize QMdiSubWindow::minimumSizeHint();
extern void* C_ZNK13QMdiSubWindow15minimumSizeHintEv(void* qthis); // 4
  // proto:  void QMdiSubWindow::showSystemMenu();
extern void C_ZN13QMdiSubWindow14showSystemMenuEv(void* qthis); // 4
  // proto:  QWidget * QMdiSubWindow::maximizedSystemMenuIconWidget();
extern void* C_ZNK13QMdiSubWindow29maximizedSystemMenuIconWidgetEv(void* qthis); // 4
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
}

// class sizeof(QMdiSubWindow)=1
type QMdiSubWindow struct {
  /*qbase*/ QWidget;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _aboutToActivate QMdiSubWindow_aboutToActivate_signal;
//  _windowStateChanged QMdiSubWindow_windowStateChanged_signal;
}

// setKeyboardPageStep(int)
func (this *QMdiSubWindow) Setkeyboardpagestep(args ...interface{}) () {
  // setKeyboardPageStep(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QMdiSubWindow19setKeyboardPageStepEi
    // invoke: void setKeyboardPageStep(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN13QMdiSubWindow19setKeyboardPageStepEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "setKeyboardPageStep", args)
  }

  return
}

// ~QMdiSubWindow()
func (this *QMdiSubWindow) Freeqmdisubwindow(args ...interface{}) () {
  // ~QMdiSubWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QMdiSubWindowD0Ev
    // invoke: void ~QMdiSubWindow()
    C.C_ZN13QMdiSubWindowD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "~QMdiSubWindow", args)
  }

  return
}

// setWidget(class QWidget *)
func (this *QMdiSubWindow) Setwidget(args ...interface{}) () {
  // setWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QMdiSubWindow9setWidgetEP7QWidget
    // invoke: void setWidget(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QMdiSubWindow9setWidgetEP7QWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "setWidget", args)
  }

  return
}

// maximizedButtonsWidget()
func (this *QMdiSubWindow) Maximizedbuttonswidget(args ...interface{}) (ret interface{}) {
  // maximizedButtonsWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow22maximizedButtonsWidgetEv
    // invoke: QWidget * maximizedButtonsWidget()
    var ret0 = C.C_ZNK13QMdiSubWindow22maximizedButtonsWidgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "maximizedButtonsWidget", args)
  }

  return
}

// mdiArea()
func (this *QMdiSubWindow) Mdiarea(args ...interface{}) (ret interface{}) {
  // mdiArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow7mdiAreaEv
    // invoke: QMdiArea * mdiArea()
    var ret0 = C.C_ZNK13QMdiSubWindow7mdiAreaEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMdiArea{}) // "QMdiArea *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "mdiArea", args)
  }

  return
}

// setSystemMenu(class QMenu *)
func (this *QMdiSubWindow) Setsystemmenu(args ...interface{}) () {
  // setSystemMenu(class QMenu *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMenu{}) // "QMenu *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QMdiSubWindow13setSystemMenuEP5QMenu
    // invoke: void setSystemMenu(class QMenu *)
    var arg0 = args[0].(*QMenu).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QMdiSubWindow13setSystemMenuEP5QMenu(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "setSystemMenu", args)
  }

  return
}

// keyboardPageStep()
func (this *QMdiSubWindow) Keyboardpagestep(args ...interface{}) (ret interface{}) {
  // keyboardPageStep()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow16keyboardPageStepEv
    // invoke: int keyboardPageStep()
    var ret0 = C.C_ZNK13QMdiSubWindow16keyboardPageStepEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "keyboardPageStep", args)
  }

  return
}

// setKeyboardSingleStep(int)
func (this *QMdiSubWindow) Setkeyboardsinglestep(args ...interface{}) () {
  // setKeyboardSingleStep(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QMdiSubWindow21setKeyboardSingleStepEi
    // invoke: void setKeyboardSingleStep(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN13QMdiSubWindow21setKeyboardSingleStepEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "setKeyboardSingleStep", args)
  }

  return
}

// showShaded()
func (this *QMdiSubWindow) Showshaded(args ...interface{}) () {
  // showShaded()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QMdiSubWindow10showShadedEv
    // invoke: void showShaded()
    C.C_ZN13QMdiSubWindow10showShadedEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "showShaded", args)
  }

  return
}

// widget()
func (this *QMdiSubWindow) Widget(args ...interface{}) (ret interface{}) {
  // widget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow6widgetEv
    // invoke: QWidget * widget()
    var ret0 = C.C_ZNK13QMdiSubWindow6widgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "widget", args)
  }

  return
}

// keyboardSingleStep()
func (this *QMdiSubWindow) Keyboardsinglestep(args ...interface{}) (ret interface{}) {
  // keyboardSingleStep()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow18keyboardSingleStepEv
    // invoke: int keyboardSingleStep()
    var ret0 = C.C_ZNK13QMdiSubWindow18keyboardSingleStepEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "keyboardSingleStep", args)
  }

  return
}

// systemMenu()
func (this *QMdiSubWindow) Systemmenu(args ...interface{}) (ret interface{}) {
  // systemMenu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow10systemMenuEv
    // invoke: QMenu * systemMenu()
    var ret0 = C.C_ZNK13QMdiSubWindow10systemMenuEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMenu{}) // "QMenu *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "systemMenu", args)
  }

  return
}

// metaObject()
func (this *QMdiSubWindow) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK13QMdiSubWindow10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "metaObject", args)
  }

  return
}

// isShaded()
func (this *QMdiSubWindow) Isshaded(args ...interface{}) (ret interface{}) {
  // isShaded()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow8isShadedEv
    // invoke: bool isShaded()
    var ret0 = C.C_ZNK13QMdiSubWindow8isShadedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "isShaded", args)
  }

  return
}

// sizeHint()
func (this *QMdiSubWindow) Sizehint(args ...interface{}) (ret interface{}) {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow8sizeHintEv
    // invoke: QSize sizeHint()
    var ret0 = C.C_ZNK13QMdiSubWindow8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "sizeHint", args)
  }

  return
}

// minimumSizeHint()
func (this *QMdiSubWindow) Minimumsizehint(args ...interface{}) (ret interface{}) {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    var ret0 = C.C_ZNK13QMdiSubWindow15minimumSizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "minimumSizeHint", args)
  }

  return
}

// showSystemMenu()
func (this *QMdiSubWindow) Showsystemmenu(args ...interface{}) () {
  // showSystemMenu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QMdiSubWindow14showSystemMenuEv
    // invoke: void showSystemMenu()
    C.C_ZN13QMdiSubWindow14showSystemMenuEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "showSystemMenu", args)
  }

  return
}

// maximizedSystemMenuIconWidget()
func (this *QMdiSubWindow) Maximizedsystemmenuiconwidget(args ...interface{}) (ret interface{}) {
  // maximizedSystemMenuIconWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow29maximizedSystemMenuIconWidgetEv
    // invoke: QWidget * maximizedSystemMenuIconWidget()
    var ret0 = C.C_ZNK13QMdiSubWindow29maximizedSystemMenuIconWidgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "maximizedSystemMenuIconWidget", args)
  }

  return
}

// <= body block end

