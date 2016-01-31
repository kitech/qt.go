package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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
extern void C_ZNK17QAccessibleWidget15foregroundColorEv(void* qthis); // 4
  // proto:  int QAccessibleWidget::indexOfChild(const QAccessibleInterface * child);
extern void C_ZNK17QAccessibleWidget12indexOfChildEPK20QAccessibleInterface(void* qthis, void* arg0); // 4
  // proto:  QWindow * QAccessibleWidget::window();
extern void C_ZNK17QAccessibleWidget6windowEv(void* qthis); // 4
  // proto:  QAccessible::State QAccessibleWidget::state();
extern void C_ZNK17QAccessibleWidget5stateEv(void* qthis); // 4
  // proto:  QAccessible::Role QAccessibleWidget::role();
extern void C_ZNK17QAccessibleWidget4roleEv(void* qthis); // 4
  // proto:  QColor QAccessibleWidget::backgroundColor();
extern void C_ZNK17QAccessibleWidget15backgroundColorEv(void* qthis); // 4
  // proto:  int QAccessibleWidget::childCount();
extern void C_ZNK17QAccessibleWidget10childCountEv(void* qthis); // 4
  // proto:  QStringList QAccessibleWidget::keyBindingsForAction(const QString & actionName);
extern void C_ZNK17QAccessibleWidget20keyBindingsForActionERK7QString(void* qthis, void* arg0); // 4
  // proto:  QAccessibleInterface * QAccessibleWidget::parent();
extern void C_ZNK17QAccessibleWidget6parentEv(void* qthis); // 4
  // proto:  bool QAccessibleWidget::isValid();
extern void C_ZNK17QAccessibleWidget7isValidEv(void* qthis); // 4
  // proto:  QAccessibleInterface * QAccessibleWidget::focusChild();
extern void C_ZNK17QAccessibleWidget10focusChildEv(void* qthis); // 4
  // proto:  QAccessibleInterface * QAccessibleWidget::child(int index);
extern void C_ZNK17QAccessibleWidget5childEi(void* qthis, int32_t arg0); // 4
  // proto:  QRect QAccessibleWidget::rect();
extern void C_ZNK17QAccessibleWidget4rectEv(void* qthis); // 4
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
func (this *QAccessibleWidget) foregroundColor(args ...interface{}) () {
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
    C.C_ZNK17QAccessibleWidget15foregroundColorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "foregroundColor", args)
  }

}

// indexOfChild(const class QAccessibleInterface *)
func (this *QAccessibleWidget) indexOfChild(args ...interface{}) () {
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
    C.C_ZNK17QAccessibleWidget12indexOfChildEPK20QAccessibleInterface(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "indexOfChild", args)
  }

}

// window()
func (this *QAccessibleWidget) window(args ...interface{}) () {
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
    C.C_ZNK17QAccessibleWidget6windowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "window", args)
  }

}

// state()
func (this *QAccessibleWidget) state(args ...interface{}) () {
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

}

// role()
func (this *QAccessibleWidget) role(args ...interface{}) () {
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

}

// backgroundColor()
func (this *QAccessibleWidget) backgroundColor(args ...interface{}) () {
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
    C.C_ZNK17QAccessibleWidget15backgroundColorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "backgroundColor", args)
  }

}

// childCount()
func (this *QAccessibleWidget) childCount(args ...interface{}) () {
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
    C.C_ZNK17QAccessibleWidget10childCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "childCount", args)
  }

}

// keyBindingsForAction(const class QString &)
func (this *QAccessibleWidget) keyBindingsForAction(args ...interface{}) () {
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

}

// parent()
func (this *QAccessibleWidget) parent(args ...interface{}) () {
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
    C.C_ZNK17QAccessibleWidget6parentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "parent", args)
  }

}

// isValid()
func (this *QAccessibleWidget) isValid(args ...interface{}) () {
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
    C.C_ZNK17QAccessibleWidget7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "isValid", args)
  }

}

// focusChild()
func (this *QAccessibleWidget) focusChild(args ...interface{}) () {
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
    C.C_ZNK17QAccessibleWidget10focusChildEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "focusChild", args)
  }

}

// child(int)
func (this *QAccessibleWidget) child(args ...interface{}) () {
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
    C.C_ZNK17QAccessibleWidget5childEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "child", args)
  }

}

// rect()
func (this *QAccessibleWidget) rect(args ...interface{}) () {
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
    C.C_ZNK17QAccessibleWidget4rectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "rect", args)
  }

}

// doAction(const class QString &)
func (this *QAccessibleWidget) doAction(args ...interface{}) () {
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

}

// actionNames()
func (this *QAccessibleWidget) actionNames(args ...interface{}) () {
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

}

// <= body block end

