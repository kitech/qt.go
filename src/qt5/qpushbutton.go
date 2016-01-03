package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtWidgets/qpushbutton.h
// dst-file: /src/widgets/qpushbutton.go
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
  // proto:  void QPushButton::setMenu(QMenu * menu);
extern void _ZN11QPushButton7setMenuEP5QMenu(void* qthis, void* arg0);
  // proto:  void QPushButton::setFlat(bool );
extern void _ZN11QPushButton7setFlatEb(void* qthis, bool arg0);
  // proto:  void QPushButton::setAutoDefault(bool );
extern void _ZN11QPushButton14setAutoDefaultEb(void* qthis, bool arg0);
  // proto:  QSize QPushButton::minimumSizeHint();
extern void _ZNK11QPushButton15minimumSizeHintEv(void* qthis);
  // proto:  void QPushButton::setDefault(bool );
extern void _ZN11QPushButton10setDefaultEb(void* qthis, bool arg0);
  // proto:  void QPushButton::~QPushButton();
extern void _ZN11QPushButtonD0Ev(void* qthis);
  // proto:  void QPushButton::QPushButton(const QPushButton & );
extern void* dector_ZN11QPushButtonC1ERKS_(void* arg0);
extern void _ZN11QPushButtonC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QPushButton::isDefault();
extern void _ZNK11QPushButton9isDefaultEv(void* qthis);
  // proto:  void QPushButton::QPushButton(const QIcon & icon, const QString & text, QWidget * parent);
extern void* dector_ZN11QPushButtonC1ERK5QIconRK7QStringP7QWidget(void* arg0, void* arg1, void* arg2);
extern void _ZN11QPushButtonC1ERK5QIconRK7QStringP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  bool QPushButton::autoDefault();
extern void _ZNK11QPushButton11autoDefaultEv(void* qthis);
  // proto:  QSize QPushButton::sizeHint();
extern void _ZNK11QPushButton8sizeHintEv(void* qthis);
  // proto:  const QMetaObject * QPushButton::metaObject();
extern void _ZNK11QPushButton10metaObjectEv(void* qthis);
  // proto:  QMenu * QPushButton::menu();
extern void _ZNK11QPushButton4menuEv(void* qthis);
  // proto:  void QPushButton::QPushButton(QWidget * parent);
extern void* dector_ZN11QPushButtonC1EP7QWidget(void* arg0);
extern void _ZN11QPushButtonC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QPushButton::showMenu();
extern void _ZN11QPushButton8showMenuEv(void* qthis);
  // proto:  void QPushButton::QPushButton(const QString & text, QWidget * parent);
extern void* dector_ZN11QPushButtonC1ERK7QStringP7QWidget(void* arg0, void* arg1);
extern void _ZN11QPushButtonC1ERK7QStringP7QWidget(void* qthis, void* arg0, void* arg1);
  // proto:  bool QPushButton::isFlat();
extern void _ZNK11QPushButton6isFlatEv(void* qthis);
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

// class sizeof(QPushButton)=1
type QPushButton struct {
  /*qbase*/ QAbstractButton;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  void QPushButton::setMenu(QMenu * menu);
func (this *QPushButton) setMenu(args ...interface{}) () {
  // setMenu(class QMenu *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMenu{}) // "QMenu *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPushButton7setMenuEP5QMenu
    // invoke: void setMenu(class QMenu *)
    var arg0 = args[0].(QMenu).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QPushButton7setMenuEP5QMenu(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPushButton", "setMenu", args)
  }

}

  // proto:  void QPushButton::setFlat(bool );
func (this *QPushButton) setFlat(args ...interface{}) () {
  // setFlat(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPushButton7setFlatEb
    // invoke: void setFlat(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN11QPushButton7setFlatEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPushButton", "setFlat", args)
  }

}

  // proto:  void QPushButton::setAutoDefault(bool );
func (this *QPushButton) setAutoDefault(args ...interface{}) () {
  // setAutoDefault(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPushButton14setAutoDefaultEb
    // invoke: void setAutoDefault(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN11QPushButton14setAutoDefaultEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPushButton", "setAutoDefault", args)
  }

}

  // proto:  QSize QPushButton::minimumSizeHint();
func (this *QPushButton) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPushButton15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    C._ZNK11QPushButton15minimumSizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPushButton", "minimumSizeHint", args)
  }

}

  // proto:  void QPushButton::setDefault(bool );
func (this *QPushButton) setDefault(args ...interface{}) () {
  // setDefault(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPushButton10setDefaultEb
    // invoke: void setDefault(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN11QPushButton10setDefaultEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPushButton", "setDefault", args)
  }

}

  // proto:  void QPushButton::~QPushButton();
func (this *QPushButton) FreeQPushButton(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPushButton", "~QPushButton", args)
  }

}

  // proto:  void QPushButton::QPushButton(const QPushButton & );
func NewQPushButton(args ...interface{}) QPushButton {
  return QPushButton{}
}

  // proto:  bool QPushButton::isDefault();
func (this *QPushButton) isDefault(args ...interface{}) () {
  // isDefault()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPushButton9isDefaultEv
    // invoke: bool isDefault()
    C._ZNK11QPushButton9isDefaultEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPushButton", "isDefault", args)
  }

}

  // proto:  bool QPushButton::autoDefault();
func (this *QPushButton) autoDefault(args ...interface{}) () {
  // autoDefault()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPushButton11autoDefaultEv
    // invoke: bool autoDefault()
    C._ZNK11QPushButton11autoDefaultEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPushButton", "autoDefault", args)
  }

}

  // proto:  QSize QPushButton::sizeHint();
func (this *QPushButton) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPushButton8sizeHintEv
    // invoke: QSize sizeHint()
    C._ZNK11QPushButton8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPushButton", "sizeHint", args)
  }

}

  // proto:  const QMetaObject * QPushButton::metaObject();
func (this *QPushButton) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPushButton10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK11QPushButton10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPushButton", "metaObject", args)
  }

}

  // proto:  QMenu * QPushButton::menu();
func (this *QPushButton) menu(args ...interface{}) () {
  // menu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPushButton4menuEv
    // invoke: QMenu * menu()
    C._ZNK11QPushButton4menuEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPushButton", "menu", args)
  }

}

  // proto:  void QPushButton::showMenu();
func (this *QPushButton) showMenu(args ...interface{}) () {
  // showMenu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPushButton8showMenuEv
    // invoke: void showMenu()
    C._ZN11QPushButton8showMenuEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPushButton", "showMenu", args)
  }

}

  // proto:  bool QPushButton::isFlat();
func (this *QPushButton) isFlat(args ...interface{}) () {
  // isFlat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPushButton6isFlatEv
    // invoke: bool isFlat()
    C._ZNK11QPushButton6isFlatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPushButton", "isFlat", args)
  }

}

// <= body block end

