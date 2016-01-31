package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtWidgets/qaccessiblewidget.h
// dst-file: /src/widgets/qaccessiblewidget.go
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
  // proto:  QColor QAccessibleWidget::foregroundColor();
extern void* C_ZNK17QAccessibleWidget15foregroundColorEv(void* qthis); // 4
  // proto:  int QAccessibleWidget::indexOfChild(const QAccessibleInterface * child);
extern int32_t C_ZNK17QAccessibleWidget12indexOfChildEPK20QAccessibleInterface(void* qthis, void* arg0); // 4
  // proto:  QWindow * QAccessibleWidget::window();
extern void* C_ZNK17QAccessibleWidget6windowEv(void* qthis); // 4
  // proto:  QAccessible::State QAccessibleWidget::state();
extern void C_ZNK17QAccessibleWidget5stateEv(void* qthis); // 4
  // proto:  QAccessible::Role QAccessibleWidget::role();
extern void C_ZNK17QAccessibleWidget4roleEv(void* qthis); // 4
  // proto:  QColor QAccessibleWidget::backgroundColor();
extern void* C_ZNK17QAccessibleWidget15backgroundColorEv(void* qthis); // 4
  // proto:  int QAccessibleWidget::childCount();
extern int32_t C_ZNK17QAccessibleWidget10childCountEv(void* qthis); // 4
  // proto:  QStringList QAccessibleWidget::keyBindingsForAction(const QString & actionName);
extern void C_ZNK17QAccessibleWidget20keyBindingsForActionERK7QString(void* qthis, void* arg0); // 4
  // proto:  QAccessibleInterface * QAccessibleWidget::parent();
extern void* C_ZNK17QAccessibleWidget6parentEv(void* qthis); // 4
  // proto:  bool QAccessibleWidget::isValid();
extern bool C_ZNK17QAccessibleWidget7isValidEv(void* qthis); // 4
  // proto:  QAccessibleInterface * QAccessibleWidget::focusChild();
extern void* C_ZNK17QAccessibleWidget10focusChildEv(void* qthis); // 4
  // proto:  QAccessibleInterface * QAccessibleWidget::child(int index);
extern void* C_ZNK17QAccessibleWidget5childEi(void* qthis, int32_t arg0); // 4
  // proto:  QRect QAccessibleWidget::rect();
extern void* C_ZNK17QAccessibleWidget4rectEv(void* qthis); // 4
  // proto:  void QAccessibleWidget::doAction(const QString & actionName);
extern void C_ZN17QAccessibleWidget8doActionERK7QString(void* qthis, void* arg0); // 4
  // proto:  QStringList QAccessibleWidget::actionNames();
extern void C_ZNK17QAccessibleWidget11actionNamesEv(void* qthis); // 4
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

// class sizeof(QAccessibleWidget)=32
type QAccessibleWidget struct {
  /*qbase*/ QAccessibleObject;
  qclsinst unsafe.Pointer /* *C.void */;
}

// foregroundColor()
func (this *QAccessibleWidget) Foregroundcolor(args ...interface{}) (ret interface{}) {
  // foregroundColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAccessibleWidget15foregroundColorEv
    // invoke: QColor foregroundColor()
    var ret0 = C.C_ZNK17QAccessibleWidget15foregroundColorEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "foregroundColor", args)
  }

  return
}

// indexOfChild(const class QAccessibleInterface *)
func (this *QAccessibleWidget) Indexofchild(args ...interface{}) (ret interface{}) {
  // indexOfChild(const class QAccessibleInterface *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAccessibleInterface{}) // "const QAccessibleInterface *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAccessibleWidget12indexOfChildEPK20QAccessibleInterface
    // invoke: int indexOfChild(const class QAccessibleInterface *)
    var arg0 = args[0].(QAccessibleInterface).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK17QAccessibleWidget12indexOfChildEPK20QAccessibleInterface(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "indexOfChild", args)
  }

  return
}

// window()
func (this *QAccessibleWidget) Window(args ...interface{}) (ret interface{}) {
  // window()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAccessibleWidget6windowEv
    // invoke: QWindow * window()
    var ret0 = C.C_ZNK17QAccessibleWidget6windowEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWindow{}) // "QWindow *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "window", args)
  }

  return
}

// state()
func (this *QAccessibleWidget) State(args ...interface{}) () {
  // state()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAccessibleWidget5stateEv
    // invoke: QAccessible::State state()
    C.C_ZNK17QAccessibleWidget5stateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "state", args)
  }

  return
}

// role()
func (this *QAccessibleWidget) Role(args ...interface{}) () {
  // role()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAccessibleWidget4roleEv
    // invoke: QAccessible::Role role()
    C.C_ZNK17QAccessibleWidget4roleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "role", args)
  }

  return
}

// backgroundColor()
func (this *QAccessibleWidget) Backgroundcolor(args ...interface{}) (ret interface{}) {
  // backgroundColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAccessibleWidget15backgroundColorEv
    // invoke: QColor backgroundColor()
    var ret0 = C.C_ZNK17QAccessibleWidget15backgroundColorEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "backgroundColor", args)
  }

  return
}

// childCount()
func (this *QAccessibleWidget) Childcount(args ...interface{}) (ret interface{}) {
  // childCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAccessibleWidget10childCountEv
    // invoke: int childCount()
    var ret0 = C.C_ZNK17QAccessibleWidget10childCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "childCount", args)
  }

  return
}

// keyBindingsForAction(const class QString &)
func (this *QAccessibleWidget) Keybindingsforaction(args ...interface{}) () {
  // keyBindingsForAction(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAccessibleWidget20keyBindingsForActionERK7QString
    // invoke: QStringList keyBindingsForAction(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK17QAccessibleWidget20keyBindingsForActionERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "keyBindingsForAction", args)
  }

  return
}

// parent()
func (this *QAccessibleWidget) Parent(args ...interface{}) (ret interface{}) {
  // parent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAccessibleWidget6parentEv
    // invoke: QAccessibleInterface * parent()
    var ret0 = C.C_ZNK17QAccessibleWidget6parentEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAccessibleInterface{}) // "QAccessibleInterface *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "parent", args)
  }

  return
}

// isValid()
func (this *QAccessibleWidget) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAccessibleWidget7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK17QAccessibleWidget7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "isValid", args)
  }

  return
}

// focusChild()
func (this *QAccessibleWidget) Focuschild(args ...interface{}) (ret interface{}) {
  // focusChild()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAccessibleWidget10focusChildEv
    // invoke: QAccessibleInterface * focusChild()
    var ret0 = C.C_ZNK17QAccessibleWidget10focusChildEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAccessibleInterface{}) // "QAccessibleInterface *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "focusChild", args)
  }

  return
}

// child(int)
func (this *QAccessibleWidget) Child(args ...interface{}) (ret interface{}) {
  // child(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAccessibleWidget5childEi
    // invoke: QAccessibleInterface * child(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK17QAccessibleWidget5childEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAccessibleInterface{}) // "QAccessibleInterface *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "child", args)
  }

  return
}

// rect()
func (this *QAccessibleWidget) Rect(args ...interface{}) (ret interface{}) {
  // rect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAccessibleWidget4rectEv
    // invoke: QRect rect()
    var ret0 = C.C_ZNK17QAccessibleWidget4rectEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "rect", args)
  }

  return
}

// doAction(const class QString &)
func (this *QAccessibleWidget) Doaction(args ...interface{}) () {
  // doAction(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAccessibleWidget8doActionERK7QString
    // invoke: void doAction(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QAccessibleWidget8doActionERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "doAction", args)
  }

  return
}

// actionNames()
func (this *QAccessibleWidget) Actionnames(args ...interface{}) () {
  // actionNames()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAccessibleWidget11actionNamesEv
    // invoke: QStringList actionNames()
    C.C_ZNK17QAccessibleWidget11actionNamesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "actionNames", args)
  }

  return
}

// <= body block end

