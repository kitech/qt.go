package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtWidgets/qbuttongroup.h
// dst-file: /src/widgets/qbuttongroup.go
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
  // proto:  bool QButtonGroup::exclusive();
extern void _ZNK12QButtonGroup9exclusiveEv(void* qthis); // 4
  // proto:  void QButtonGroup::QButtonGroup(QObject * parent);
extern void _ZN12QButtonGroupC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  void QButtonGroup::addButton(QAbstractButton * , int id);
extern void _ZN12QButtonGroup9addButtonEP15QAbstractButtoni(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QButtonGroup::~QButtonGroup();
extern void _ZN12QButtonGroupD2Ev(void* qthis); // 4
  // proto:  void QButtonGroup::removeButton(QAbstractButton * );
extern void _ZN12QButtonGroup12removeButtonEP15QAbstractButton(void* qthis, void* arg0); // 4
  // proto:  int QButtonGroup::id(QAbstractButton * button);
extern void _ZNK12QButtonGroup2idEP15QAbstractButton(void* qthis, void* arg0); // 4
  // proto:  QList<QAbstractButton *> QButtonGroup::buttons();
extern void _ZNK12QButtonGroup7buttonsEv(void* qthis); // 4
  // proto:  QAbstractButton * QButtonGroup::checkedButton();
extern void _ZNK12QButtonGroup13checkedButtonEv(void* qthis); // 4
  // proto:  int QButtonGroup::checkedId();
extern void _ZNK12QButtonGroup9checkedIdEv(void* qthis); // 4
  // proto:  const QMetaObject * QButtonGroup::metaObject();
extern void _ZNK12QButtonGroup10metaObjectEv(void* qthis); // 4
  // proto:  void QButtonGroup::setExclusive(bool );
extern void _ZN12QButtonGroup12setExclusiveEb(void* qthis, bool arg0); // 4
  // proto:  QAbstractButton * QButtonGroup::button(int id);
extern void _ZNK12QButtonGroup6buttonEi(void* qthis, int32_t arg0); // 4
  // proto:  void QButtonGroup::setId(QAbstractButton * button, int id);
extern void _ZN12QButtonGroup5setIdEP15QAbstractButtoni(void* qthis, void* arg0, int32_t arg1); // 4
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

// class sizeof(QButtonGroup)=1
type QButtonGroup struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _buttonToggled QButtonGroup_buttonToggled_signal;
//  _buttonClicked QButtonGroup_buttonClicked_signal;
//  _buttonReleased QButtonGroup_buttonReleased_signal;
//  _buttonPressed QButtonGroup_buttonPressed_signal;
}

// exclusive()
func (this *QButtonGroup) exclusive(args ...interface{}) () {
  // exclusive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QButtonGroup9exclusiveEv
    // invoke: bool exclusive()
    C._ZNK12QButtonGroup9exclusiveEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QButtonGroup", "exclusive", args)
  }

}

// QButtonGroup(class QObject *)
func NewQButtonGroup(args ...interface{}) QButtonGroup {
  // QButtonGroup(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QButtonGroupC1EP7QObject
    // invoke: void QButtonGroup(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN12QButtonGroupC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QButtonGroup", "QButtonGroup", args)
  }

  return QButtonGroup{}
}

// addButton(class QAbstractButton *, int)
func (this *QButtonGroup) addButton(args ...interface{}) () {
  // addButton(class QAbstractButton *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractButton{}) // "QAbstractButton *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QButtonGroup9addButtonEP15QAbstractButtoni
    // invoke: void addButton(class QAbstractButton *, int)
    var arg0 = args[0].(QAbstractButton).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN12QButtonGroup9addButtonEP15QAbstractButtoni(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QButtonGroup", "addButton", args)
  }

}

// ~QButtonGroup()
func (this *QButtonGroup) FreeQButtonGroup(args ...interface{}) () {
  // ~QButtonGroup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QButtonGroupD0Ev
    // invoke: void ~QButtonGroup()
    C._ZN12QButtonGroupD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QButtonGroup", "~QButtonGroup", args)
  }

}

// removeButton(class QAbstractButton *)
func (this *QButtonGroup) removeButton(args ...interface{}) () {
  // removeButton(class QAbstractButton *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractButton{}) // "QAbstractButton *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QButtonGroup12removeButtonEP15QAbstractButton
    // invoke: void removeButton(class QAbstractButton *)
    var arg0 = args[0].(QAbstractButton).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN12QButtonGroup12removeButtonEP15QAbstractButton(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QButtonGroup", "removeButton", args)
  }

}

// id(class QAbstractButton *)
func (this *QButtonGroup) id(args ...interface{}) () {
  // id(class QAbstractButton *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractButton{}) // "QAbstractButton *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QButtonGroup2idEP15QAbstractButton
    // invoke: int id(class QAbstractButton *)
    var arg0 = args[0].(QAbstractButton).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK12QButtonGroup2idEP15QAbstractButton(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QButtonGroup", "id", args)
  }

}

// buttons()
func (this *QButtonGroup) buttons(args ...interface{}) () {
  // buttons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QButtonGroup7buttonsEv
    // invoke: QList<QAbstractButton *> buttons()
    C._ZNK12QButtonGroup7buttonsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QButtonGroup", "buttons", args)
  }

}

// checkedButton()
func (this *QButtonGroup) checkedButton(args ...interface{}) () {
  // checkedButton()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QButtonGroup13checkedButtonEv
    // invoke: QAbstractButton * checkedButton()
    C._ZNK12QButtonGroup13checkedButtonEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QButtonGroup", "checkedButton", args)
  }

}

// checkedId()
func (this *QButtonGroup) checkedId(args ...interface{}) () {
  // checkedId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QButtonGroup9checkedIdEv
    // invoke: int checkedId()
    C._ZNK12QButtonGroup9checkedIdEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QButtonGroup", "checkedId", args)
  }

}

// metaObject()
func (this *QButtonGroup) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QButtonGroup10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK12QButtonGroup10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QButtonGroup", "metaObject", args)
  }

}

// setExclusive(_Bool)
func (this *QButtonGroup) setExclusive(args ...interface{}) () {
  // setExclusive(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QButtonGroup12setExclusiveEb
    // invoke: void setExclusive(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN12QButtonGroup12setExclusiveEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QButtonGroup", "setExclusive", args)
  }

}

// button(int)
func (this *QButtonGroup) button(args ...interface{}) () {
  // button(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QButtonGroup6buttonEi
    // invoke: QAbstractButton * button(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK12QButtonGroup6buttonEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QButtonGroup", "button", args)
  }

}

// setId(class QAbstractButton *, int)
func (this *QButtonGroup) setId(args ...interface{}) () {
  // setId(class QAbstractButton *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractButton{}) // "QAbstractButton *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QButtonGroup5setIdEP15QAbstractButtoni
    // invoke: void setId(class QAbstractButton *, int)
    var arg0 = args[0].(QAbstractButton).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN12QButtonGroup5setIdEP15QAbstractButtoni(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QButtonGroup", "setId", args)
  }

}

// <= body block end

