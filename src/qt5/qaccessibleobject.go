package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtGui/qaccessibleobject.h
// dst-file: /src/gui/qaccessibleobject.go
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
  // proto:  void QAccessibleApplication::QAccessibleApplication();
extern void* dector_ZN22QAccessibleApplicationC1Ev();
extern void _ZN22QAccessibleApplicationC1Ev(void* qthis);
  // proto:  QWindow * QAccessibleApplication::window();
extern void _ZNK22QAccessibleApplication6windowEv(void* qthis);
  // proto:  QAccessibleInterface * QAccessibleApplication::child(int index);
extern void _ZNK22QAccessibleApplication5childEi(void* qthis, int arg0);
  // proto:  int QAccessibleApplication::childCount();
extern void _ZNK22QAccessibleApplication10childCountEv(void* qthis);
  // proto:  QAccessibleInterface * QAccessibleApplication::parent();
extern void _ZNK22QAccessibleApplication6parentEv(void* qthis);
  // proto:  QAccessibleInterface * QAccessibleApplication::focusChild();
extern void _ZNK22QAccessibleApplication10focusChildEv(void* qthis);
  // proto:  int QAccessibleApplication::indexOfChild(const QAccessibleInterface * );
extern void _ZNK22QAccessibleApplication12indexOfChildEPK20QAccessibleInterface(void* qthis, void* arg0);
  // proto:  void QAccessibleObject::QAccessibleObject(QObject * object);
extern void* dector_ZN17QAccessibleObjectC1EP7QObject(void* arg0);
extern void _ZN17QAccessibleObjectC1EP7QObject(void* qthis, void* arg0);
  // proto:  QObject * QAccessibleObject::object();
extern void _ZNK17QAccessibleObject6objectEv(void* qthis);
  // proto:  QRect QAccessibleObject::rect();
extern void _ZNK17QAccessibleObject4rectEv(void* qthis);
  // proto:  QAccessibleInterface * QAccessibleObject::childAt(int x, int y);
extern void _ZNK17QAccessibleObject7childAtEii(void* qthis, int arg0, int arg1);
  // proto:  void QAccessibleObject::QAccessibleObject(const QAccessibleObject & );
extern void* dector_ZN17QAccessibleObjectC1ERKS_(void* arg0);
extern void _ZN17QAccessibleObjectC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QAccessibleObject::isValid();
extern void _ZNK17QAccessibleObject7isValidEv(void* qthis);
  // proto:  void QAccessibleObject::~QAccessibleObject();
extern void _ZN17QAccessibleObjectD0Ev(void* qthis);
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

// class sizeof(QAccessibleApplication)=16
type QAccessibleApplication struct {
  /*qbase*/ QAccessibleObject;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleObject)=16
type QAccessibleObject struct {
  /*qbase*/ QAccessibleInterface;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QAccessibleApplication::QAccessibleApplication();
func NewQAccessibleApplication(args ...interface{}) QAccessibleApplication {
  return QAccessibleApplication{}
}

  // proto:  QWindow * QAccessibleApplication::window();
func (this *QAccessibleApplication) window(args ...interface{}) () {
  // window()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QAccessibleApplication6windowEv
    // invoke: QWindow * window()
    C._ZNK22QAccessibleApplication6windowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleApplication", "window", args)
  }

}

  // proto:  QAccessibleInterface * QAccessibleApplication::child(int index);
func (this *QAccessibleApplication) child(args ...interface{}) () {
  // child(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QAccessibleApplication5childEi
    // invoke: QAccessibleInterface * child(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK22QAccessibleApplication5childEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAccessibleApplication", "child", args)
  }

}

  // proto:  int QAccessibleApplication::childCount();
func (this *QAccessibleApplication) childCount(args ...interface{}) () {
  // childCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QAccessibleApplication10childCountEv
    // invoke: int childCount()
    C._ZNK22QAccessibleApplication10childCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleApplication", "childCount", args)
  }

}

  // proto:  QAccessibleInterface * QAccessibleApplication::parent();
func (this *QAccessibleApplication) parent(args ...interface{}) () {
  // parent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QAccessibleApplication6parentEv
    // invoke: QAccessibleInterface * parent()
    C._ZNK22QAccessibleApplication6parentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleApplication", "parent", args)
  }

}

  // proto:  QAccessibleInterface * QAccessibleApplication::focusChild();
func (this *QAccessibleApplication) focusChild(args ...interface{}) () {
  // focusChild()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QAccessibleApplication10focusChildEv
    // invoke: QAccessibleInterface * focusChild()
    C._ZNK22QAccessibleApplication10focusChildEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleApplication", "focusChild", args)
  }

}

  // proto:  int QAccessibleApplication::indexOfChild(const QAccessibleInterface * );
func (this *QAccessibleApplication) indexOfChild(args ...interface{}) () {
  // indexOfChild(const class QAccessibleInterface *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAccessibleInterface{}) // "const QAccessibleInterface *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QAccessibleApplication12indexOfChildEPK20QAccessibleInterface
    // invoke: int indexOfChild(const class QAccessibleInterface *)
    var arg0 = args[0].(QAccessibleInterface).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK22QAccessibleApplication12indexOfChildEPK20QAccessibleInterface(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAccessibleApplication", "indexOfChild", args)
  }

}

  // proto:  void QAccessibleObject::QAccessibleObject(QObject * object);
func NewQAccessibleObject(args ...interface{}) QAccessibleObject {
  return QAccessibleObject{}
}

  // proto:  QObject * QAccessibleObject::object();
func (this *QAccessibleObject) object(args ...interface{}) () {
  // object()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAccessibleObject6objectEv
    // invoke: QObject * object()
    C._ZNK17QAccessibleObject6objectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleObject", "object", args)
  }

}

  // proto:  QRect QAccessibleObject::rect();
func (this *QAccessibleObject) rect(args ...interface{}) () {
  // rect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAccessibleObject4rectEv
    // invoke: QRect rect()
    C._ZNK17QAccessibleObject4rectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleObject", "rect", args)
  }

}

  // proto:  QAccessibleInterface * QAccessibleObject::childAt(int x, int y);
func (this *QAccessibleObject) childAt(args ...interface{}) () {
  // childAt(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAccessibleObject7childAtEii
    // invoke: QAccessibleInterface * childAt(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK17QAccessibleObject7childAtEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAccessibleObject", "childAt", args)
  }

}

  // proto:  bool QAccessibleObject::isValid();
func (this *QAccessibleObject) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAccessibleObject7isValidEv
    // invoke: bool isValid()
    C._ZNK17QAccessibleObject7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleObject", "isValid", args)
  }

}

  // proto:  void QAccessibleObject::~QAccessibleObject();
func (this *QAccessibleObject) FreeQAccessibleObject(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleObject", "~QAccessibleObject", args)
  }

}

// <= body block end

