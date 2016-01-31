package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtWidgets/qtoolbutton.h
// dst-file: /src/widgets/qtoolbutton.go
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
  // proto:  void QToolButton::showMenu();
extern void C_ZN11QToolButton8showMenuEv(void* qthis); // 4
  // proto:  void QToolButton::QToolButton(QWidget * parent);
extern void C_ZN11QToolButtonC2EP7QWidget(void* qthis, void* arg0); // 3
  // proto:  void QToolButton::setMenu(QMenu * menu);
extern void C_ZN11QToolButton7setMenuEP5QMenu(void* qthis, void* arg0); // 4
  // proto:  void QToolButton::~QToolButton();
extern void C_ZN11QToolButtonD2Ev(void* qthis); // 4
  // proto:  void QToolButton::setAutoRaise(bool enable);
extern void C_ZN11QToolButton12setAutoRaiseEb(void* qthis, bool arg0); // 4
  // proto:  QMenu * QToolButton::menu();
extern void C_ZNK11QToolButton4menuEv(void* qthis); // 4
  // proto:  bool QToolButton::autoRaise();
extern void C_ZNK11QToolButton9autoRaiseEv(void* qthis); // 4
  // proto:  QAction * QToolButton::defaultAction();
extern void C_ZNK11QToolButton13defaultActionEv(void* qthis); // 4
  // proto:  void QToolButton::setDefaultAction(QAction * );
extern void C_ZN11QToolButton16setDefaultActionEP7QAction(void* qthis, void* arg0); // 4
  // proto:  QSize QToolButton::sizeHint();
extern void C_ZNK11QToolButton8sizeHintEv(void* qthis); // 4
  // proto:  const QMetaObject * QToolButton::metaObject();
extern void C_ZNK11QToolButton10metaObjectEv(void* qthis); // 4
  // proto:  QSize QToolButton::minimumSizeHint();
extern void C_ZNK11QToolButton15minimumSizeHintEv(void* qthis); // 4
  // proto:  Qt::ArrowType QToolButton::arrowType();
extern void C_ZNK11QToolButton9arrowTypeEv(void* qthis); // 4
  // proto:  Qt::ToolButtonStyle QToolButton::toolButtonStyle();
extern void C_ZNK11QToolButton15toolButtonStyleEv(void* qthis); // 4
  // proto:  QToolButton::ToolButtonPopupMode QToolButton::popupMode();
extern void C_ZNK11QToolButton9popupModeEv(void* qthis); // 4
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

// class sizeof(QToolButton)=1
type QToolButton struct {
  /*qbase*/ QAbstractButton;
  qclsinst unsafe.Pointer /* *C.void */;
//  _triggered QToolButton_triggered_signal;
}

// showMenu()
func (this *QToolButton) showMenu(args ...interface{}) () {
  // showMenu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QToolButton8showMenuEv
    // invoke: void showMenu()
    C.C_ZN11QToolButton8showMenuEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolButton", "showMenu", args)
  }

}

// QToolButton(class QWidget *)
func NewQToolButton(args ...interface{}) QToolButton {
  // QToolButton(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QToolButtonC1EP7QWidget
    // invoke: void QToolButton(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QToolButtonC2EP7QWidget(qthis, arg0)
  default:
    qtrt.ErrorResolve("QToolButton", "QToolButton", args)
  }

  return QToolButton{}
}

// setMenu(class QMenu *)
func (this *QToolButton) setMenu(args ...interface{}) () {
  // setMenu(class QMenu *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMenu{}) // "QMenu *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QToolButton7setMenuEP5QMenu
    // invoke: void setMenu(class QMenu *)
    var arg0 = args[0].(QMenu).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QToolButton7setMenuEP5QMenu(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolButton", "setMenu", args)
  }

}

// ~QToolButton()
func (this *QToolButton) FreeQToolButton(args ...interface{}) () {
  // ~QToolButton()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QToolButtonD0Ev
    // invoke: void ~QToolButton()
    C.C_ZN11QToolButtonD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolButton", "~QToolButton", args)
  }

}

// setAutoRaise(_Bool)
func (this *QToolButton) setAutoRaise(args ...interface{}) () {
  // setAutoRaise(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QToolButton12setAutoRaiseEb
    // invoke: void setAutoRaise(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QToolButton12setAutoRaiseEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolButton", "setAutoRaise", args)
  }

}

// menu()
func (this *QToolButton) menu(args ...interface{}) () {
  // menu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QToolButton4menuEv
    // invoke: QMenu * menu()
    C.C_ZNK11QToolButton4menuEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolButton", "menu", args)
  }

}

// autoRaise()
func (this *QToolButton) autoRaise(args ...interface{}) () {
  // autoRaise()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QToolButton9autoRaiseEv
    // invoke: bool autoRaise()
    C.C_ZNK11QToolButton9autoRaiseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolButton", "autoRaise", args)
  }

}

// defaultAction()
func (this *QToolButton) defaultAction(args ...interface{}) () {
  // defaultAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QToolButton13defaultActionEv
    // invoke: QAction * defaultAction()
    C.C_ZNK11QToolButton13defaultActionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolButton", "defaultAction", args)
  }

}

// setDefaultAction(class QAction *)
func (this *QToolButton) setDefaultAction(args ...interface{}) () {
  // setDefaultAction(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QToolButton16setDefaultActionEP7QAction
    // invoke: void setDefaultAction(class QAction *)
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QToolButton16setDefaultActionEP7QAction(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolButton", "setDefaultAction", args)
  }

}

// sizeHint()
func (this *QToolButton) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QToolButton8sizeHintEv
    // invoke: QSize sizeHint()
    C.C_ZNK11QToolButton8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolButton", "sizeHint", args)
  }

}

// metaObject()
func (this *QToolButton) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QToolButton10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK11QToolButton10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolButton", "metaObject", args)
  }

}

// minimumSizeHint()
func (this *QToolButton) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QToolButton15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    C.C_ZNK11QToolButton15minimumSizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolButton", "minimumSizeHint", args)
  }

}

// arrowType()
func (this *QToolButton) arrowType(args ...interface{}) () {
  // arrowType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QToolButton9arrowTypeEv
    // invoke: Qt::ArrowType arrowType()
    C.C_ZNK11QToolButton9arrowTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolButton", "arrowType", args)
  }

}

// toolButtonStyle()
func (this *QToolButton) toolButtonStyle(args ...interface{}) () {
  // toolButtonStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QToolButton15toolButtonStyleEv
    // invoke: Qt::ToolButtonStyle toolButtonStyle()
    C.C_ZNK11QToolButton15toolButtonStyleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolButton", "toolButtonStyle", args)
  }

}

// popupMode()
func (this *QToolButton) popupMode(args ...interface{}) () {
  // popupMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QToolButton9popupModeEv
    // invoke: QToolButton::ToolButtonPopupMode popupMode()
    C.C_ZNK11QToolButton9popupModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolButton", "popupMode", args)
  }

}

// <= body block end

