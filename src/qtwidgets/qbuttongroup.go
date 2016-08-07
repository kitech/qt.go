package qtwidgets
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
import "qtcore"
import "qtgui"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  bool QButtonGroup::exclusive();
extern bool C_ZNK12QButtonGroup9exclusiveEv(void* qthis); // 4
  // proto:  void QButtonGroup::QButtonGroup(QObject * parent);
extern void* C_ZN12QButtonGroupC2EP7QObject(void* arg0); // 3
  // proto:  void QButtonGroup::addButton(QAbstractButton * , int id);
extern void C_ZN12QButtonGroup9addButtonEP15QAbstractButtoni(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QButtonGroup::~QButtonGroup();
extern void C_ZN12QButtonGroupD2Ev(void* qthis); // 4
  // proto:  void QButtonGroup::removeButton(QAbstractButton * );
extern void C_ZN12QButtonGroup12removeButtonEP15QAbstractButton(void* qthis, void* arg0); // 4
  // proto:  int QButtonGroup::id(QAbstractButton * button);
extern int32_t C_ZNK12QButtonGroup2idEP15QAbstractButton(void* qthis, void* arg0); // 4
  // proto:  QList<QAbstractButton *> QButtonGroup::buttons();
extern void C_ZNK12QButtonGroup7buttonsEv(void* qthis); // 4
  // proto:  QAbstractButton * QButtonGroup::checkedButton();
extern void C_ZNK12QButtonGroup13checkedButtonEv(void* qthis); // 4
  // proto:  int QButtonGroup::checkedId();
extern int32_t C_ZNK12QButtonGroup9checkedIdEv(void* qthis); // 4
  // proto:  const QMetaObject * QButtonGroup::metaObject();
extern void C_ZNK12QButtonGroup10metaObjectEv(void* qthis); // 4
  // proto:  void QButtonGroup::setExclusive(bool );
extern void C_ZN12QButtonGroup12setExclusiveEb(void* qthis, bool arg0); // 4
  // proto:  QAbstractButton * QButtonGroup::button(int id);
extern void C_ZNK12QButtonGroup6buttonEi(void* qthis, int32_t arg0); // 4
  // proto:  void QButtonGroup::setId(QAbstractButton * button, int id);
extern void C_ZN12QButtonGroup5setIdEP15QAbstractButtoni(void* qthis, void* arg0, int32_t arg1); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QButtonGroup)=1
type QButtonGroup struct {
  /*qbase*/ qtcore.QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _buttonToggled QButtonGroup_buttonToggled_signal;
//  _buttonClicked QButtonGroup_buttonClicked_signal;
//  _buttonReleased QButtonGroup_buttonReleased_signal;
//  _buttonPressed QButtonGroup_buttonPressed_signal;
}

// exclusive()
func (this *QButtonGroup) Exclusive(args ...interface{}) (ret interface{}) {
  // exclusive()
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
    // invoke: _ZNK12QButtonGroup9exclusiveEv
    // invoke: bool exclusive()
    var ret0 = C.C_ZNK12QButtonGroup9exclusiveEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QButtonGroup", "exclusive", args)
  }

  return
}

// QButtonGroup(class QObject *)
func NewQButtonGroup(args ...interface{}) *QButtonGroup {
  // QButtonGroup(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QButtonGroupC1EP7QObject
    // invoke: void QButtonGroup(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QButtonGroupC2EP7QObject(arg0)
    return &QButtonGroup{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QButtonGroup", "QButtonGroup", args)
  }

  return nil // QButtonGroup{Qclsinst:qthis}
}

// addButton(class QAbstractButton *, int)
func (this *QButtonGroup) Addbutton(args ...interface{}) () {
  // addButton(class QAbstractButton *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractButton{}) // "QAbstractButton *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QButtonGroup9addButtonEP15QAbstractButtoni
    // invoke: void addButton(class QAbstractButton *, int)
    var arg0 = args[0].(*QAbstractButton).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN12QButtonGroup9addButtonEP15QAbstractButtoni(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QButtonGroup", "addButton", args)
  }

  return
}

// ~QButtonGroup()
func (this *QButtonGroup) Freeqbuttongroup(args ...interface{}) () {
  // ~QButtonGroup()
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
    // invoke: _ZN12QButtonGroupD0Ev
    // invoke: void ~QButtonGroup()
    C.C_ZN12QButtonGroupD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QButtonGroup", "~QButtonGroup", args)
  }

  return
}

// removeButton(class QAbstractButton *)
func (this *QButtonGroup) Removebutton(args ...interface{}) () {
  // removeButton(class QAbstractButton *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractButton{}) // "QAbstractButton *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QButtonGroup12removeButtonEP15QAbstractButton
    // invoke: void removeButton(class QAbstractButton *)
    var arg0 = args[0].(*QAbstractButton).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QButtonGroup12removeButtonEP15QAbstractButton(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QButtonGroup", "removeButton", args)
  }

  return
}

// id(class QAbstractButton *)
func (this *QButtonGroup) Id(args ...interface{}) (ret interface{}) {
  // id(class QAbstractButton *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractButton{}) // "QAbstractButton *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QButtonGroup2idEP15QAbstractButton
    // invoke: int id(class QAbstractButton *)
    var arg0 = args[0].(*QAbstractButton).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QButtonGroup2idEP15QAbstractButton(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QButtonGroup", "id", args)
  }

  return
}

// buttons()
func (this *QButtonGroup) Buttons(args ...interface{}) () {
  // buttons()
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
    // invoke: _ZNK12QButtonGroup7buttonsEv
    // invoke: QList<QAbstractButton *> buttons()
    C.C_ZNK12QButtonGroup7buttonsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QButtonGroup", "buttons", args)
  }

  return
}

// checkedButton()
func (this *QButtonGroup) Checkedbutton(args ...interface{}) () {
  // checkedButton()
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
    // invoke: _ZNK12QButtonGroup13checkedButtonEv
    // invoke: QAbstractButton * checkedButton()
    C.C_ZNK12QButtonGroup13checkedButtonEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QButtonGroup", "checkedButton", args)
  }

  return
}

// checkedId()
func (this *QButtonGroup) Checkedid(args ...interface{}) (ret interface{}) {
  // checkedId()
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
    // invoke: _ZNK12QButtonGroup9checkedIdEv
    // invoke: int checkedId()
    var ret0 = C.C_ZNK12QButtonGroup9checkedIdEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QButtonGroup", "checkedId", args)
  }

  return
}

// metaObject()
func (this *QButtonGroup) Metaobject(args ...interface{}) () {
  // metaObject()
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
    // invoke: _ZNK12QButtonGroup10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK12QButtonGroup10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QButtonGroup", "metaObject", args)
  }

  return
}

// setExclusive(_Bool)
func (this *QButtonGroup) Setexclusive(args ...interface{}) () {
  // setExclusive(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QButtonGroup12setExclusiveEb
    // invoke: void setExclusive(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN12QButtonGroup12setExclusiveEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QButtonGroup", "setExclusive", args)
  }

  return
}

// button(int)
func (this *QButtonGroup) Button(args ...interface{}) () {
  // button(int)
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
    // invoke: _ZNK12QButtonGroup6buttonEi
    // invoke: QAbstractButton * button(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK12QButtonGroup6buttonEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QButtonGroup", "button", args)
  }

  return
}

// setId(class QAbstractButton *, int)
func (this *QButtonGroup) Setid(args ...interface{}) () {
  // setId(class QAbstractButton *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractButton{}) // "QAbstractButton *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QButtonGroup5setIdEP15QAbstractButtoni
    // invoke: void setId(class QAbstractButton *, int)
    var arg0 = args[0].(*QAbstractButton).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN12QButtonGroup5setIdEP15QAbstractButtoni(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QButtonGroup", "setId", args)
  }

  return
}

// <= body block end

