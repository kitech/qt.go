package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
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
  // proto:  void QButtonGroup::addButton(QAbstractButton * , int id);
extern void _ZN12QButtonGroup9addButtonEP15QAbstractButtoni(void* qthis, void* arg0, int32_t arg1);
  // proto:  QList<QAbstractButton *> QButtonGroup::buttons();
extern void _ZNK12QButtonGroup7buttonsEv(void* qthis);
  // proto:  void QButtonGroup::~QButtonGroup();
extern void _ZN12QButtonGroupD0Ev(void* qthis);
  // proto:  int QButtonGroup::id(QAbstractButton * button);
extern void _ZNK12QButtonGroup2idEP15QAbstractButton(void* qthis, void* arg0);
  // proto:  void QButtonGroup::removeButton(QAbstractButton * );
extern void _ZN12QButtonGroup12removeButtonEP15QAbstractButton(void* qthis, void* arg0);
  // proto:  void QButtonGroup::QButtonGroup(const QButtonGroup & );
extern void* dector_ZN12QButtonGroupC1ERKS_(void* arg0);
extern void _ZN12QButtonGroupC1ERKS_(void* qthis, void* arg0);
  // proto:  const QMetaObject * QButtonGroup::metaObject();
extern void _ZNK12QButtonGroup10metaObjectEv(void* qthis);
  // proto:  void QButtonGroup::QButtonGroup(QObject * parent);
extern void* dector_ZN12QButtonGroupC1EP7QObject(void* arg0);
extern void _ZN12QButtonGroupC1EP7QObject(void* qthis, void* arg0);
  // proto:  QAbstractButton * QButtonGroup::button(int id);
extern void _ZNK12QButtonGroup6buttonEi(void* qthis, int32_t arg0);
  // proto:  int QButtonGroup::checkedId();
extern void _ZNK12QButtonGroup9checkedIdEv(void* qthis);
  // proto:  QAbstractButton * QButtonGroup::checkedButton();
extern void _ZNK12QButtonGroup13checkedButtonEv(void* qthis);
  // proto:  void QButtonGroup::setExclusive(bool );
extern void _ZN12QButtonGroup12setExclusiveEb(void* qthis, bool arg0);
  // proto:  void QButtonGroup::setId(QAbstractButton * button, int id);
extern void _ZN12QButtonGroup5setIdEP15QAbstractButtoni(void* qthis, void* arg0, int32_t arg1);
  // proto:  bool QButtonGroup::exclusive();
extern void _ZNK12QButtonGroup9exclusiveEv(void* qthis);
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

  // proto:  void QButtonGroup::addButton(QAbstractButton * , int id);
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

  // proto:  QList<QAbstractButton *> QButtonGroup::buttons();
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

  // proto:  void QButtonGroup::~QButtonGroup();
func (this *QButtonGroup) FreeQButtonGroup(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QButtonGroup", "~QButtonGroup", args)
  }

}

  // proto:  int QButtonGroup::id(QAbstractButton * button);
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

  // proto:  void QButtonGroup::removeButton(QAbstractButton * );
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

  // proto:  void QButtonGroup::QButtonGroup(const QButtonGroup & );
func NewQButtonGroup(args ...interface{}) QButtonGroup {
  return QButtonGroup{}
}

  // proto:  const QMetaObject * QButtonGroup::metaObject();
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

  // proto:  QAbstractButton * QButtonGroup::button(int id);
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

  // proto:  int QButtonGroup::checkedId();
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

  // proto:  QAbstractButton * QButtonGroup::checkedButton();
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

  // proto:  void QButtonGroup::setExclusive(bool );
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

  // proto:  void QButtonGroup::setId(QAbstractButton * button, int id);
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

  // proto:  bool QButtonGroup::exclusive();
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

// <= body block end

