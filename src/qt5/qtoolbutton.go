package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
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
  // proto:  void QToolButton::setAutoRaise(bool enable);
extern void _ZN11QToolButton12setAutoRaiseEb(void* qthis, bool arg0);
  // proto:  QAction * QToolButton::defaultAction();
extern void _ZNK11QToolButton13defaultActionEv(void* qthis);
  // proto:  const QMetaObject * QToolButton::metaObject();
extern void _ZNK11QToolButton10metaObjectEv(void* qthis);
  // proto:  QSize QToolButton::minimumSizeHint();
extern void _ZNK11QToolButton15minimumSizeHintEv(void* qthis);
  // proto:  void QToolButton::~QToolButton();
extern void _ZN11QToolButtonD0Ev(void* qthis);
  // proto:  void QToolButton::showMenu();
extern void _ZN11QToolButton8showMenuEv(void* qthis);
  // proto:  QSize QToolButton::sizeHint();
extern void _ZNK11QToolButton8sizeHintEv(void* qthis);
  // proto:  void QToolButton::QToolButton(const QToolButton & );
extern void* dector_ZN11QToolButtonC1ERKS_(void* arg0);
extern void _ZN11QToolButtonC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QToolButton::autoRaise();
extern void _ZNK11QToolButton9autoRaiseEv(void* qthis);
  // proto:  QMenu * QToolButton::menu();
extern void _ZNK11QToolButton4menuEv(void* qthis);
  // proto:  void QToolButton::setMenu(QMenu * menu);
extern void _ZN11QToolButton7setMenuEP5QMenu(void* qthis, void* arg0);
  // proto:  void QToolButton::QToolButton(QWidget * parent);
extern void* dector_ZN11QToolButtonC1EP7QWidget(void* arg0);
extern void _ZN11QToolButtonC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QToolButton::setDefaultAction(QAction * );
extern void _ZN11QToolButton16setDefaultActionEP7QAction(void* qthis, void* arg0);
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
  qclsinst uint64 /* *mut c_void*/;
//  _triggered QToolButton_triggered_signal;
}

  // proto:  void QToolButton::setAutoRaise(bool enable);
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QToolButton", "setAutoRaise", args)
  }

}

  // proto:  QAction * QToolButton::defaultAction();
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
  default:
    qtrt.ErrorResolve("QToolButton", "defaultAction", args)
  }

}

  // proto:  const QMetaObject * QToolButton::metaObject();
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
  default:
    qtrt.ErrorResolve("QToolButton", "metaObject", args)
  }

}

  // proto:  QSize QToolButton::minimumSizeHint();
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
  default:
    qtrt.ErrorResolve("QToolButton", "minimumSizeHint", args)
  }

}

  // proto:  void QToolButton::~QToolButton();
func (this *QToolButton) FreeQToolButton(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QToolButton", "~QToolButton", args)
  }

}

  // proto:  void QToolButton::showMenu();
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
  default:
    qtrt.ErrorResolve("QToolButton", "showMenu", args)
  }

}

  // proto:  QSize QToolButton::sizeHint();
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
  default:
    qtrt.ErrorResolve("QToolButton", "sizeHint", args)
  }

}

  // proto:  void QToolButton::QToolButton(const QToolButton & );
func NewQToolButton(args ...interface{}) QToolButton {
  return QToolButton{}
}

  // proto:  bool QToolButton::autoRaise();
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
  default:
    qtrt.ErrorResolve("QToolButton", "autoRaise", args)
  }

}

  // proto:  QMenu * QToolButton::menu();
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
  default:
    qtrt.ErrorResolve("QToolButton", "menu", args)
  }

}

  // proto:  void QToolButton::setMenu(QMenu * menu);
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
    var arg0 = args[0].(QMenu).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QToolButton", "setMenu", args)
  }

}

  // proto:  void QToolButton::setDefaultAction(QAction * );
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
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QToolButton", "setDefaultAction", args)
  }

}

// <= body block end

