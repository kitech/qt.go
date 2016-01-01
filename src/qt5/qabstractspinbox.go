package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qabstractspinbox.h
// dst-file: /src/widgets/qabstractspinbox.go
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
// class sizeof(QAbstractSpinBox)=1
type QAbstractSpinBox struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
//  _editingFinished QAbstractSpinBox_editingFinished_signal;
}


func (this *QAbstractSpinBox) stepBy(args ...interface{}) () {
  // stepBy(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox6stepByEi
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "stepBy", args)
  }

}


func (this *QAbstractSpinBox) setReadOnly(args ...interface{}) () {
  // setReadOnly(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox11setReadOnlyEb
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "setReadOnly", args)
  }

}


func (this *QAbstractSpinBox) setFrame(args ...interface{}) () {
  // setFrame(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox8setFrameEb
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "setFrame", args)
  }

}


func (this *QAbstractSpinBox) setSpecialValueText(args ...interface{}) () {
  // setSpecialValueText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox19setSpecialValueTextERK7QString
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "setSpecialValueText", args)
  }

}


func (this *QAbstractSpinBox) setAccelerated(args ...interface{}) () {
  // setAccelerated(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox14setAcceleratedEb
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "setAccelerated", args)
  }

}


func (this *QAbstractSpinBox) interpretText(args ...interface{}) () {
  // interpretText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox13interpretTextEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "interpretText", args)
  }

}


func (this *QAbstractSpinBox) event(args ...interface{}) () {
  // event(class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox5eventEP6QEvent
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "event", args)
  }

}


func (this *QAbstractSpinBox) keyboardTracking(args ...interface{}) () {
  // keyboardTracking()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox16keyboardTrackingEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "keyboardTracking", args)
  }

}


func NewQAbstractSpinBox(args ...interface{}) QAbstractSpinBox {
  return QAbstractSpinBox{}
}


func (this *QAbstractSpinBox) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox10metaObjectEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "metaObject", args)
  }

}


func (this *QAbstractSpinBox) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox8sizeHintEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "sizeHint", args)
  }

}


func (this *QAbstractSpinBox) FreeQAbstractSpinBox(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "~QAbstractSpinBox", args)
  }

}


func (this *QAbstractSpinBox) fixup(args ...interface{}) () {
  // fixup(class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox5fixupER7QString
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "fixup", args)
  }

}


func (this *QAbstractSpinBox) selectAll(args ...interface{}) () {
  // selectAll()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox9selectAllEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "selectAll", args)
  }

}


func (this *QAbstractSpinBox) stepDown(args ...interface{}) () {
  // stepDown()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox8stepDownEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "stepDown", args)
  }

}


func (this *QAbstractSpinBox) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox5clearEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "clear", args)
  }

}


func (this *QAbstractSpinBox) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox4textEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "text", args)
  }

}


func (this *QAbstractSpinBox) specialValueText(args ...interface{}) () {
  // specialValueText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox16specialValueTextEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "specialValueText", args)
  }

}


func (this *QAbstractSpinBox) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox15minimumSizeHintEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "minimumSizeHint", args)
  }

}


func (this *QAbstractSpinBox) wrapping(args ...interface{}) () {
  // wrapping()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox8wrappingEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "wrapping", args)
  }

}


func (this *QAbstractSpinBox) stepUp(args ...interface{}) () {
  // stepUp()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox6stepUpEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "stepUp", args)
  }

}


func (this *QAbstractSpinBox) setWrapping(args ...interface{}) () {
  // setWrapping(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox11setWrappingEb
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "setWrapping", args)
  }

}


func (this *QAbstractSpinBox) setKeyboardTracking(args ...interface{}) () {
  // setKeyboardTracking(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox19setKeyboardTrackingEb
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "setKeyboardTracking", args)
  }

}


func (this *QAbstractSpinBox) isAccelerated(args ...interface{}) () {
  // isAccelerated()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox13isAcceleratedEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "isAccelerated", args)
  }

}


func (this *QAbstractSpinBox) setGroupSeparatorShown(args ...interface{}) () {
  // setGroupSeparatorShown(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox22setGroupSeparatorShownEb
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "setGroupSeparatorShown", args)
  }

}


func (this *QAbstractSpinBox) isReadOnly(args ...interface{}) () {
  // isReadOnly()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox10isReadOnlyEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "isReadOnly", args)
  }

}


func (this *QAbstractSpinBox) hasAcceptableInput(args ...interface{}) () {
  // hasAcceptableInput()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox18hasAcceptableInputEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "hasAcceptableInput", args)
  }

}


func (this *QAbstractSpinBox) isGroupSeparatorShown(args ...interface{}) () {
  // isGroupSeparatorShown()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox21isGroupSeparatorShownEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "isGroupSeparatorShown", args)
  }

}


func (this *QAbstractSpinBox) hasFrame(args ...interface{}) () {
  // hasFrame()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox8hasFrameEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "hasFrame", args)
  }

}

// <= body block end

