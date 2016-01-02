package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
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
  // proto:  int QAccessibleWidget::childCount();
extern void _ZNK17QAccessibleWidget10childCountEv(void* qthis);
  // proto:  QAccessibleInterface * QAccessibleWidget::child(int index);
extern void _ZNK17QAccessibleWidget5childEi(void* qthis, int arg0);
  // proto:  QWindow * QAccessibleWidget::window();
extern void _ZNK17QAccessibleWidget6windowEv(void* qthis);
  // proto:  QRect QAccessibleWidget::rect();
extern void _ZNK17QAccessibleWidget4rectEv(void* qthis);
  // proto:  QColor QAccessibleWidget::foregroundColor();
extern void _ZNK17QAccessibleWidget15foregroundColorEv(void* qthis);
  // proto:  bool QAccessibleWidget::isValid();
extern void _ZNK17QAccessibleWidget7isValidEv(void* qthis);
  // proto:  QAccessibleInterface * QAccessibleWidget::focusChild();
extern void _ZNK17QAccessibleWidget10focusChildEv(void* qthis);
  // proto:  void QAccessibleWidget::QAccessibleWidget(const QAccessibleWidget & );
extern void* dector_ZN17QAccessibleWidgetC1ERKS_(void* arg0);
extern void _ZN17QAccessibleWidgetC1ERKS_(void* qthis, void* arg0);
  // proto:  QColor QAccessibleWidget::backgroundColor();
extern void _ZNK17QAccessibleWidget15backgroundColorEv(void* qthis);
  // proto:  void QAccessibleWidget::~QAccessibleWidget();
extern void _ZN17QAccessibleWidgetD0Ev(void* qthis);
  // proto:  QStringList QAccessibleWidget::actionNames();
extern void _ZNK17QAccessibleWidget11actionNamesEv(void* qthis);
  // proto:  QAccessibleInterface * QAccessibleWidget::parent();
extern void _ZNK17QAccessibleWidget6parentEv(void* qthis);
  // proto:  void QAccessibleWidget::doAction(const QString & actionName);
extern void _ZN17QAccessibleWidget8doActionERK7QString(void* qthis, void* arg0);
  // proto:  QStringList QAccessibleWidget::keyBindingsForAction(const QString & actionName);
extern void _ZNK17QAccessibleWidget20keyBindingsForActionERK7QString(void* qthis, void* arg0);
  // proto:  int QAccessibleWidget::indexOfChild(const QAccessibleInterface * child);
extern void _ZNK17QAccessibleWidget12indexOfChildEPK20QAccessibleInterface(void* qthis, void* arg0);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  int QAccessibleWidget::childCount();
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
    C._ZNK17QAccessibleWidget10childCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "childCount", args)
  }

}

  // proto:  QAccessibleInterface * QAccessibleWidget::child(int index);
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
    C._ZNK17QAccessibleWidget5childEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "child", args)
  }

}

  // proto:  QWindow * QAccessibleWidget::window();
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
    C._ZNK17QAccessibleWidget6windowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "window", args)
  }

}

  // proto:  QRect QAccessibleWidget::rect();
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
    C._ZNK17QAccessibleWidget4rectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "rect", args)
  }

}

  // proto:  QColor QAccessibleWidget::foregroundColor();
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
    C._ZNK17QAccessibleWidget15foregroundColorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "foregroundColor", args)
  }

}

  // proto:  bool QAccessibleWidget::isValid();
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
    C._ZNK17QAccessibleWidget7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "isValid", args)
  }

}

  // proto:  QAccessibleInterface * QAccessibleWidget::focusChild();
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
    C._ZNK17QAccessibleWidget10focusChildEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "focusChild", args)
  }

}

  // proto:  void QAccessibleWidget::QAccessibleWidget(const QAccessibleWidget & );
func NewQAccessibleWidget(args ...interface{}) QAccessibleWidget {
  return QAccessibleWidget{}
}

  // proto:  QColor QAccessibleWidget::backgroundColor();
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
    C._ZNK17QAccessibleWidget15backgroundColorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "backgroundColor", args)
  }

}

  // proto:  void QAccessibleWidget::~QAccessibleWidget();
func (this *QAccessibleWidget) FreeQAccessibleWidget(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "~QAccessibleWidget", args)
  }

}

  // proto:  QStringList QAccessibleWidget::actionNames();
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
    C._ZNK17QAccessibleWidget11actionNamesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "actionNames", args)
  }

}

  // proto:  QAccessibleInterface * QAccessibleWidget::parent();
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
    C._ZNK17QAccessibleWidget6parentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "parent", args)
  }

}

  // proto:  void QAccessibleWidget::doAction(const QString & actionName);
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
    C._ZN17QAccessibleWidget8doActionERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "doAction", args)
  }

}

  // proto:  QStringList QAccessibleWidget::keyBindingsForAction(const QString & actionName);
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
    C._ZNK17QAccessibleWidget20keyBindingsForActionERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "keyBindingsForAction", args)
  }

}

  // proto:  int QAccessibleWidget::indexOfChild(const QAccessibleInterface * child);
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
    C._ZNK17QAccessibleWidget12indexOfChildEPK20QAccessibleInterface(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAccessibleWidget", "indexOfChild", args)
  }

}

// <= body block end

