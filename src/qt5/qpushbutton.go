package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QPushButton::showMenu();
extern void C_ZN11QPushButton8showMenuEv(void* qthis); // 4
  // proto:  void QPushButton::setMenu(QMenu * menu);
extern void C_ZN11QPushButton7setMenuEP5QMenu(void* qthis, void* arg0); // 4
  // proto:  void QPushButton::~QPushButton();
extern void C_ZN11QPushButtonD2Ev(void* qthis); // 4
  // proto:  void QPushButton::setAutoDefault(bool );
extern void C_ZN11QPushButton14setAutoDefaultEb(void* qthis, bool arg0); // 4
  // proto:  void QPushButton::setFlat(bool );
extern void C_ZN11QPushButton7setFlatEb(void* qthis, bool arg0); // 4
  // proto:  QMenu * QPushButton::menu();
extern void C_ZNK11QPushButton4menuEv(void* qthis); // 4
  // proto:  void QPushButton::QPushButton(const QString & text, QWidget * parent);
extern void* C_ZN11QPushButtonC2ERK7QStringP7QWidget(void* arg0, void* arg1); // 3
  // proto:  void QPushButton::QPushButton(const QIcon & icon, const QString & text, QWidget * parent);
extern void* C_ZN11QPushButtonC2ERK5QIconRK7QStringP7QWidget(void* arg0, void* arg1, void* arg2); // 3
  // proto:  void QPushButton::QPushButton(QWidget * parent);
extern void* C_ZN11QPushButtonC2EP7QWidget(void* arg0); // 3
  // proto:  bool QPushButton::isFlat();
extern void C_ZNK11QPushButton6isFlatEv(void* qthis); // 4
  // proto:  bool QPushButton::isDefault();
extern void C_ZNK11QPushButton9isDefaultEv(void* qthis); // 4
  // proto:  void QPushButton::setDefault(bool );
extern void C_ZN11QPushButton10setDefaultEb(void* qthis, bool arg0); // 4
  // proto:  const QMetaObject * QPushButton::metaObject();
extern void C_ZNK11QPushButton10metaObjectEv(void* qthis); // 4
  // proto:  QSize QPushButton::sizeHint();
extern void C_ZNK11QPushButton8sizeHintEv(void* qthis); // 4
  // proto:  QSize QPushButton::minimumSizeHint();
extern void C_ZNK11QPushButton15minimumSizeHintEv(void* qthis); // 4
  // proto:  bool QPushButton::autoDefault();
extern void C_ZNK11QPushButton11autoDefaultEv(void* qthis); // 4
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

// showMenu()
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
    C.C_ZN11QPushButton8showMenuEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPushButton", "showMenu", args)
  }

}

// setMenu(class QMenu *)
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
    C.C_ZN11QPushButton7setMenuEP5QMenu(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPushButton", "setMenu", args)
  }

}

// ~QPushButton()
func (this *QPushButton) FreeQPushButton(args ...interface{}) () {
  // ~QPushButton()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPushButtonD0Ev
    // invoke: void ~QPushButton()
    C.C_ZN11QPushButtonD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPushButton", "~QPushButton", args)
  }

}

// setAutoDefault(_Bool)
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
    C.C_ZN11QPushButton14setAutoDefaultEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPushButton", "setAutoDefault", args)
  }

}

// setFlat(_Bool)
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
    C.C_ZN11QPushButton7setFlatEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPushButton", "setFlat", args)
  }

}

// menu()
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
    var ret = C.C_ZNK11QPushButton4menuEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPushButton", "menu", args)
  }

}

// QPushButton(const class QString &, class QWidget *)
func NewQPushButton(args ...interface{}) *QPushButton {
  // QPushButton(const class QString &, class QWidget *)
  // QPushButton(const class QIcon &, const class QString &, class QWidget *)
  // QPushButton(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][2] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPushButtonC1ERK7QStringP7QWidget
    // invoke: void QPushButton(const class QString &, class QWidget *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QPushButtonC2ERK7QStringP7QWidget(arg0, arg1)
    return &QPushButton{qclsinst:qthis}
  case 1:
    // invoke: _ZN11QPushButtonC1ERK5QIconRK7QStringP7QWidget
    // invoke: void QPushButton(const class QIcon &, const class QString &, class QWidget *)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QWidget).qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QPushButtonC2ERK5QIconRK7QStringP7QWidget(arg0, arg1, arg2)
    return &QPushButton{qclsinst:qthis}
  case 2:
    // invoke: _ZN11QPushButtonC1EP7QWidget
    // invoke: void QPushButton(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QPushButtonC2EP7QWidget(arg0)
    return &QPushButton{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QPushButton", "QPushButton", args)
  }

  return nil // QPushButton{qclsinst:qthis}
}

// isFlat()
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
    var ret = C.C_ZNK11QPushButton6isFlatEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPushButton", "isFlat", args)
  }

}

// isDefault()
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
    var ret = C.C_ZNK11QPushButton9isDefaultEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPushButton", "isDefault", args)
  }

}

// setDefault(_Bool)
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
    C.C_ZN11QPushButton10setDefaultEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPushButton", "setDefault", args)
  }

}

// metaObject()
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
    C.C_ZNK11QPushButton10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPushButton", "metaObject", args)
  }

}

// sizeHint()
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
    var ret = C.C_ZNK11QPushButton8sizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPushButton", "sizeHint", args)
  }

}

// minimumSizeHint()
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
    var ret = C.C_ZNK11QPushButton15minimumSizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPushButton", "minimumSizeHint", args)
  }

}

// autoDefault()
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
    var ret = C.C_ZNK11QPushButton11autoDefaultEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPushButton", "autoDefault", args)
  }

}

// <= body block end

