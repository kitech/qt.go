package qtgui
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QAccessibleApplication::QAccessibleApplication();
extern void* C_ZN22QAccessibleApplicationC2Ev(); // 3
  // proto:  QAccessibleInterface * QAccessibleApplication::parent();
extern void* C_ZNK22QAccessibleApplication6parentEv(void* qthis); // 4
  // proto:  int QAccessibleApplication::indexOfChild(const QAccessibleInterface * );
extern int32_t C_ZNK22QAccessibleApplication12indexOfChildEPK20QAccessibleInterface(void* qthis, void* arg0); // 4
  // proto:  QAccessible::State QAccessibleApplication::state();
extern void C_ZNK22QAccessibleApplication5stateEv(void* qthis); // 4
  // proto:  QAccessibleInterface * QAccessibleApplication::focusChild();
extern void* C_ZNK22QAccessibleApplication10focusChildEv(void* qthis); // 4
  // proto:  QWindow * QAccessibleApplication::window();
extern void* C_ZNK22QAccessibleApplication6windowEv(void* qthis); // 4
  // proto:  QAccessible::Role QAccessibleApplication::role();
extern void C_ZNK22QAccessibleApplication4roleEv(void* qthis); // 4
  // proto:  QAccessibleInterface * QAccessibleApplication::child(int index);
extern void* C_ZNK22QAccessibleApplication5childEi(void* qthis, int32_t arg0); // 4
  // proto:  int QAccessibleApplication::childCount();
extern int32_t C_ZNK22QAccessibleApplication10childCountEv(void* qthis); // 4
  // proto:  QObject * QAccessibleObject::object();
extern void* C_ZNK17QAccessibleObject6objectEv(void* qthis); // 4
  // proto:  bool QAccessibleObject::isValid();
extern bool C_ZNK17QAccessibleObject7isValidEv(void* qthis); // 4
  // proto:  QAccessibleInterface * QAccessibleObject::childAt(int x, int y);
extern void* C_ZNK17QAccessibleObject7childAtEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QRect QAccessibleObject::rect();
extern void* C_ZNK17QAccessibleObject4rectEv(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QAccessibleApplication)=16
type QAccessibleApplication struct {
  /*qbase*/ QAccessibleObject;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleObject)=16
type QAccessibleObject struct {
  /*qbase*/ QAccessibleInterface;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// QAccessibleApplication()
func NewQAccessibleApplication(args ...interface{}) *QAccessibleApplication {
  // QAccessibleApplication()
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
    // invoke: _ZN22QAccessibleApplicationC1Ev
    // invoke: void QAccessibleApplication()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN22QAccessibleApplicationC2Ev()
    return &QAccessibleApplication{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QAccessibleApplication", "QAccessibleApplication", args)
  }

  return nil // QAccessibleApplication{Qclsinst:qthis}
}

// parent()
func (this *QAccessibleApplication) Parent(args ...interface{}) (ret interface{}) {
  // parent()
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
    // invoke: _ZNK22QAccessibleApplication6parentEv
    // invoke: QAccessibleInterface * parent()
    var ret0 = C.C_ZNK22QAccessibleApplication6parentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAccessibleInterface{}) // "QAccessibleInterface *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleApplication", "parent", args)
  }

  return
}

// indexOfChild(const class QAccessibleInterface *)
func (this *QAccessibleApplication) Indexofchild(args ...interface{}) (ret interface{}) {
  // indexOfChild(const class QAccessibleInterface *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAccessibleInterface{}) // "const QAccessibleInterface *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QAccessibleApplication12indexOfChildEPK20QAccessibleInterface
    // invoke: int indexOfChild(const class QAccessibleInterface *)
    var arg0 = args[0].(*QAccessibleInterface).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK22QAccessibleApplication12indexOfChildEPK20QAccessibleInterface(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleApplication", "indexOfChild", args)
  }

  return
}

// state()
func (this *QAccessibleApplication) State(args ...interface{}) () {
  // state()
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
    // invoke: _ZNK22QAccessibleApplication5stateEv
    // invoke: QAccessible::State state()
    C.C_ZNK22QAccessibleApplication5stateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleApplication", "state", args)
  }

  return
}

// focusChild()
func (this *QAccessibleApplication) Focuschild(args ...interface{}) (ret interface{}) {
  // focusChild()
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
    // invoke: _ZNK22QAccessibleApplication10focusChildEv
    // invoke: QAccessibleInterface * focusChild()
    var ret0 = C.C_ZNK22QAccessibleApplication10focusChildEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAccessibleInterface{}) // "QAccessibleInterface *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleApplication", "focusChild", args)
  }

  return
}

// window()
func (this *QAccessibleApplication) Window(args ...interface{}) (ret interface{}) {
  // window()
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
    // invoke: _ZNK22QAccessibleApplication6windowEv
    // invoke: QWindow * window()
    var ret0 = C.C_ZNK22QAccessibleApplication6windowEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWindow{}) // "QWindow *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleApplication", "window", args)
  }

  return
}

// role()
func (this *QAccessibleApplication) Role(args ...interface{}) () {
  // role()
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
    // invoke: _ZNK22QAccessibleApplication4roleEv
    // invoke: QAccessible::Role role()
    C.C_ZNK22QAccessibleApplication4roleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleApplication", "role", args)
  }

  return
}

// child(int)
func (this *QAccessibleApplication) Child(args ...interface{}) (ret interface{}) {
  // child(int)
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
    // invoke: _ZNK22QAccessibleApplication5childEi
    // invoke: QAccessibleInterface * child(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK22QAccessibleApplication5childEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAccessibleInterface{}) // "QAccessibleInterface *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleApplication", "child", args)
  }

  return
}

// childCount()
func (this *QAccessibleApplication) Childcount(args ...interface{}) (ret interface{}) {
  // childCount()
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
    // invoke: _ZNK22QAccessibleApplication10childCountEv
    // invoke: int childCount()
    var ret0 = C.C_ZNK22QAccessibleApplication10childCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleApplication", "childCount", args)
  }

  return
}

// object()
func (this *QAccessibleObject) Object(args ...interface{}) (ret interface{}) {
  // object()
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
    // invoke: _ZNK17QAccessibleObject6objectEv
    // invoke: QObject * object()
    var ret0 = C.C_ZNK17QAccessibleObject6objectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleObject", "object", args)
  }

  return
}

// isValid()
func (this *QAccessibleObject) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
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
    // invoke: _ZNK17QAccessibleObject7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK17QAccessibleObject7isValidEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleObject", "isValid", args)
  }

  return
}

// childAt(int, int)
func (this *QAccessibleObject) Childat(args ...interface{}) (ret interface{}) {
  // childAt(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QAccessibleObject7childAtEii
    // invoke: QAccessibleInterface * childAt(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK17QAccessibleObject7childAtEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAccessibleInterface{}) // "QAccessibleInterface *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleObject", "childAt", args)
  }

  return
}

// rect()
func (this *QAccessibleObject) Rect(args ...interface{}) (ret interface{}) {
  // rect()
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
    // invoke: _ZNK17QAccessibleObject4rectEv
    // invoke: QRect rect()
    var ret0 = C.C_ZNK17QAccessibleObject4rectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleObject", "rect", args)
  }

  return
}

// <= body block end

