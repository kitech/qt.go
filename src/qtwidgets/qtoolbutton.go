package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
  // proto:  void QToolButton::showMenu();
extern void C_ZN11QToolButton8showMenuEv(void* qthis); // 4
  // proto:  void QToolButton::QToolButton(QWidget * parent);
extern void* C_ZN11QToolButtonC2EP7QWidget(void* arg0); // 3
  // proto:  void QToolButton::setMenu(QMenu * menu);
extern void C_ZN11QToolButton7setMenuEP5QMenu(void* qthis, void* arg0); // 4
  // proto:  void QToolButton::~QToolButton();
extern void C_ZN11QToolButtonD2Ev(void* qthis); // 4
  // proto:  void QToolButton::setAutoRaise(bool enable);
extern void C_ZN11QToolButton12setAutoRaiseEb(void* qthis, bool arg0); // 4
  // proto:  QMenu * QToolButton::menu();
extern void* C_ZNK11QToolButton4menuEv(void* qthis); // 4
  // proto:  bool QToolButton::autoRaise();
extern bool C_ZNK11QToolButton9autoRaiseEv(void* qthis); // 4
  // proto:  QAction * QToolButton::defaultAction();
extern void* C_ZNK11QToolButton13defaultActionEv(void* qthis); // 4
  // proto:  void QToolButton::setDefaultAction(QAction * );
extern void C_ZN11QToolButton16setDefaultActionEP7QAction(void* qthis, void* arg0); // 4
  // proto:  QSize QToolButton::sizeHint();
extern void* C_ZNK11QToolButton8sizeHintEv(void* qthis); // 4
  // proto:  const QMetaObject * QToolButton::metaObject();
extern void C_ZNK11QToolButton10metaObjectEv(void* qthis); // 4
  // proto:  QSize QToolButton::minimumSizeHint();
extern void* C_ZNK11QToolButton15minimumSizeHintEv(void* qthis); // 4
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
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QToolButton)=1
type QToolButton struct {
  /*qbase*/ QAbstractButton;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _triggered QToolButton_triggered_signal;
}

// showMenu()
func (this *QToolButton) ShowMenu(args ...interface{}) () {
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
    C.C_ZN11QToolButton8showMenuEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QToolButton", "showMenu", args)
  }

  return
}

// QToolButton(class QWidget *)
func GcfreeQToolButton(this *QToolButton) {
  qtrt.UniverseFree(this)
}
func NewQToolButton(args ...interface{}) *QToolButton {
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QToolButtonC2EP7QWidget(arg0)
    this := &QToolButton{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQToolButton)
    return this
  default:
    qtrt.ErrorResolve("QToolButton", "QToolButton", args)
  }

  return nil // QToolButton{Qclsinst:qthis}
}

// setMenu(class QMenu *)
func (this *QToolButton) SetMenu(args ...interface{}) () {
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
    var arg0 = args[0].(*QMenu).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QToolButton7setMenuEP5QMenu(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolButton", "setMenu", args)
  }

  return
}

// ~QToolButton()
func (this *QToolButton) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN11QToolButtonD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QToolButton", "~QToolButton", args)
  }

  return
}

// setAutoRaise(_Bool)
func (this *QToolButton) SetAutoRaise(args ...interface{}) () {
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
    C.C_ZN11QToolButton12setAutoRaiseEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolButton", "setAutoRaise", args)
  }

  return
}

// menu()
func (this *QToolButton) Menu(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QToolButton4menuEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMenu{}) // "QMenu *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolButton", "menu", args)
  }

  return
}

// autoRaise()
func (this *QToolButton) AutoRaise(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QToolButton9autoRaiseEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolButton", "autoRaise", args)
  }

  return
}

// defaultAction()
func (this *QToolButton) DefaultAction(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QToolButton13defaultActionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolButton", "defaultAction", args)
  }

  return
}

// setDefaultAction(class QAction *)
func (this *QToolButton) SetDefaultAction(args ...interface{}) () {
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
    var arg0 = args[0].(*QAction).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QToolButton16setDefaultActionEP7QAction(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolButton", "setDefaultAction", args)
  }

  return
}

// sizeHint()
func (this *QToolButton) SizeHint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QToolButton8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolButton", "sizeHint", args)
  }

  return
}

// metaObject()
func (this *QToolButton) MetaObject(args ...interface{}) () {
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
    C.C_ZNK11QToolButton10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QToolButton", "metaObject", args)
  }

  return
}

// minimumSizeHint()
func (this *QToolButton) MinimumSizeHint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QToolButton15minimumSizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolButton", "minimumSizeHint", args)
  }

  return
}

// arrowType()
func (this *QToolButton) ArrowType(args ...interface{}) () {
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
    C.C_ZNK11QToolButton9arrowTypeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QToolButton", "arrowType", args)
  }

  return
}

// toolButtonStyle()
func (this *QToolButton) ToolButtonStyle(args ...interface{}) () {
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
    C.C_ZNK11QToolButton15toolButtonStyleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QToolButton", "toolButtonStyle", args)
  }

  return
}

// popupMode()
func (this *QToolButton) PopupMode(args ...interface{}) () {
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
    C.C_ZNK11QToolButton9popupModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QToolButton", "popupMode", args)
  }

  return
}

// <= body block end

