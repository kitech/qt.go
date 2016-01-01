package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtWidgets/qabstractbutton.h
// dst-file: /src/widgets/qabstractbutton.go
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
// class sizeof(QAbstractButton)=1
type QAbstractButton struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
//  _released QAbstractButton_released_signal;
//  _clicked QAbstractButton_clicked_signal;
//  _pressed QAbstractButton_pressed_signal;
//  _toggled QAbstractButton_toggled_signal;
}


func (this *QAbstractButton) iconSize(args ...interface{}) () {
  // iconSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton8iconSizeEv
  default:
    qtrt.ErrorResolve("QAbstractButton", "iconSize", args)
 }

}


func (this *QAbstractButton) click(args ...interface{}) () {
  // click()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton5clickEv
  default:
    qtrt.ErrorResolve("QAbstractButton", "click", args)
 }

}


func (this *QAbstractButton) FreeQAbstractButton(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractButton", "~QAbstractButton", args)
 }

}


func (this *QAbstractButton) setChecked(args ...interface{}) () {
  // setChecked(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton10setCheckedEb
  default:
    qtrt.ErrorResolve("QAbstractButton", "setChecked", args)
 }

}


func (this *QAbstractButton) shortcut(args ...interface{}) () {
  // shortcut()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton8shortcutEv
  default:
    qtrt.ErrorResolve("QAbstractButton", "shortcut", args)
 }

}


func (this *QAbstractButton) group(args ...interface{}) () {
  // group()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton5groupEv
  default:
    qtrt.ErrorResolve("QAbstractButton", "group", args)
 }

}


func (this *QAbstractButton) isCheckable(args ...interface{}) () {
  // isCheckable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton11isCheckableEv
  default:
    qtrt.ErrorResolve("QAbstractButton", "isCheckable", args)
 }

}


func NewQAbstractButton(args ...interface{})() {
}


func (this *QAbstractButton) isDown(args ...interface{}) () {
  // isDown()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton6isDownEv
  default:
    qtrt.ErrorResolve("QAbstractButton", "isDown", args)
 }

}


func (this *QAbstractButton) setAutoExclusive(args ...interface{}) () {
  // setAutoExclusive(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton16setAutoExclusiveEb
  default:
    qtrt.ErrorResolve("QAbstractButton", "setAutoExclusive", args)
 }

}


func (this *QAbstractButton) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton10metaObjectEv
  default:
    qtrt.ErrorResolve("QAbstractButton", "metaObject", args)
 }

}


func (this *QAbstractButton) isChecked(args ...interface{}) () {
  // isChecked()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton9isCheckedEv
  default:
    qtrt.ErrorResolve("QAbstractButton", "isChecked", args)
 }

}


func (this *QAbstractButton) setAutoRepeatDelay(args ...interface{}) () {
  // setAutoRepeatDelay(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton18setAutoRepeatDelayEi
  default:
    qtrt.ErrorResolve("QAbstractButton", "setAutoRepeatDelay", args)
 }

}


func (this *QAbstractButton) autoRepeatDelay(args ...interface{}) () {
  // autoRepeatDelay()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton15autoRepeatDelayEv
  default:
    qtrt.ErrorResolve("QAbstractButton", "autoRepeatDelay", args)
 }

}


func (this *QAbstractButton) autoExclusive(args ...interface{}) () {
  // autoExclusive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton13autoExclusiveEv
  default:
    qtrt.ErrorResolve("QAbstractButton", "autoExclusive", args)
 }

}


func (this *QAbstractButton) toggle(args ...interface{}) () {
  // toggle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton6toggleEv
  default:
    qtrt.ErrorResolve("QAbstractButton", "toggle", args)
 }

}


func (this *QAbstractButton) setIcon(args ...interface{}) () {
  // setIcon(const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton7setIconERK5QIcon
  default:
    qtrt.ErrorResolve("QAbstractButton", "setIcon", args)
 }

}


func (this *QAbstractButton) setAutoRepeatInterval(args ...interface{}) () {
  // setAutoRepeatInterval(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton21setAutoRepeatIntervalEi
  default:
    qtrt.ErrorResolve("QAbstractButton", "setAutoRepeatInterval", args)
 }

}


func (this *QAbstractButton) setAutoRepeat(args ...interface{}) () {
  // setAutoRepeat(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton13setAutoRepeatEb
  default:
    qtrt.ErrorResolve("QAbstractButton", "setAutoRepeat", args)
 }

}


func (this *QAbstractButton) animateClick(args ...interface{}) () {
  // animateClick(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton12animateClickEi
  default:
    qtrt.ErrorResolve("QAbstractButton", "animateClick", args)
 }

}


func (this *QAbstractButton) setDown(args ...interface{}) () {
  // setDown(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton7setDownEb
  default:
    qtrt.ErrorResolve("QAbstractButton", "setDown", args)
 }

}


func (this *QAbstractButton) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton4textEv
  default:
    qtrt.ErrorResolve("QAbstractButton", "text", args)
 }

}


func (this *QAbstractButton) setShortcut(args ...interface{}) () {
  // setShortcut(const class QKeySequence &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QKeySequence{}) // "const QKeySequence &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton11setShortcutERK12QKeySequence
  default:
    qtrt.ErrorResolve("QAbstractButton", "setShortcut", args)
 }

}


func (this *QAbstractButton) setCheckable(args ...interface{}) () {
  // setCheckable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton12setCheckableEb
  default:
    qtrt.ErrorResolve("QAbstractButton", "setCheckable", args)
 }

}


func (this *QAbstractButton) icon(args ...interface{}) () {
  // icon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton4iconEv
  default:
    qtrt.ErrorResolve("QAbstractButton", "icon", args)
 }

}


func (this *QAbstractButton) setText(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton7setTextERK7QString
  default:
    qtrt.ErrorResolve("QAbstractButton", "setText", args)
 }

}


func (this *QAbstractButton) autoRepeatInterval(args ...interface{}) () {
  // autoRepeatInterval()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton18autoRepeatIntervalEv
  default:
    qtrt.ErrorResolve("QAbstractButton", "autoRepeatInterval", args)
 }

}


func (this *QAbstractButton) autoRepeat(args ...interface{}) () {
  // autoRepeat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton10autoRepeatEv
  default:
    qtrt.ErrorResolve("QAbstractButton", "autoRepeat", args)
 }

}


func (this *QAbstractButton) setIconSize(args ...interface{}) () {
  // setIconSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton11setIconSizeERK5QSize
  default:
    qtrt.ErrorResolve("QAbstractButton", "setIconSize", args)
 }

}

// <= body block end

