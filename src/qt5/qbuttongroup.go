package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qbuttongroup.h
// dst-file: /src/widgets/qbuttongroup.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
// class sizeof(QButtonGroup)=1
type QButtonGroup struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _buttonToggled QButtonGroup_buttonToggled_signal;
//  _buttonClicked QButtonGroup_buttonClicked_signal;
//  _buttonReleased QButtonGroup_buttonReleased_signal;
//  _buttonPressed QButtonGroup_buttonPressed_signal;
}


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
  default:
    qtrt.ErrorResolve("QButtonGroup", "addButton", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QButtonGroup", "buttons", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QButtonGroup", "id", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QButtonGroup", "removeButton", args)
  }

}


func NewQButtonGroup(args ...interface{}) QButtonGroup {
  return QButtonGroup{}
}


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
  default:
    qtrt.ErrorResolve("QButtonGroup", "metaObject", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QButtonGroup", "button", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QButtonGroup", "checkedId", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QButtonGroup", "checkedButton", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QButtonGroup", "setExclusive", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QButtonGroup", "setId", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QButtonGroup", "exclusive", args)
  }

}

// <= body block end

