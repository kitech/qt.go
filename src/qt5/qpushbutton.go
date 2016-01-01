package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtWidgets/qpushbutton.h
// dst-file: /src/widgets/qpushbutton.go
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
// class sizeof(QPushButton)=1
type QPushButton struct {
  /*qbase*/ QAbstractButton;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QPushButton) setMenu(args ...interface{}) () {
  // setMenu(class QMenu *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMenu{}) // "QMenu *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QPushButton7setMenuEP5QMenu
  default:
    qtrt.ErrorResolve("QPushButton", "setMenu", args)
 }

}


func (this *QPushButton) setFlat(args ...interface{}) () {
  // setFlat(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QPushButton7setFlatEb
  default:
    qtrt.ErrorResolve("QPushButton", "setFlat", args)
 }

}


func (this *QPushButton) setAutoDefault(args ...interface{}) () {
  // setAutoDefault(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QPushButton14setAutoDefaultEb
  default:
    qtrt.ErrorResolve("QPushButton", "setAutoDefault", args)
 }

}


func (this *QPushButton) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QPushButton15minimumSizeHintEv
  default:
    qtrt.ErrorResolve("QPushButton", "minimumSizeHint", args)
 }

}


func (this *QPushButton) setDefault(args ...interface{}) () {
  // setDefault(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QPushButton10setDefaultEb
  default:
    qtrt.ErrorResolve("QPushButton", "setDefault", args)
 }

}


func (this *QPushButton) FreeQPushButton(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QPushButton", "~QPushButton", args)
 }

}


func NewQPushButton(args ...interface{})() {
}


func (this *QPushButton) isDefault(args ...interface{}) () {
  // isDefault()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QPushButton9isDefaultEv
  default:
    qtrt.ErrorResolve("QPushButton", "isDefault", args)
 }

}


func (this *QPushButton) autoDefault(args ...interface{}) () {
  // autoDefault()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QPushButton11autoDefaultEv
  default:
    qtrt.ErrorResolve("QPushButton", "autoDefault", args)
 }

}


func (this *QPushButton) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QPushButton8sizeHintEv
  default:
    qtrt.ErrorResolve("QPushButton", "sizeHint", args)
 }

}


func (this *QPushButton) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QPushButton10metaObjectEv
  default:
    qtrt.ErrorResolve("QPushButton", "metaObject", args)
 }

}


func (this *QPushButton) menu(args ...interface{}) () {
  // menu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QPushButton4menuEv
  default:
    qtrt.ErrorResolve("QPushButton", "menu", args)
 }

}


func (this *QPushButton) showMenu(args ...interface{}) () {
  // showMenu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QPushButton8showMenuEv
  default:
    qtrt.ErrorResolve("QPushButton", "showMenu", args)
 }

}


func (this *QPushButton) isFlat(args ...interface{}) () {
  // isFlat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QPushButton6isFlatEv
  default:
    qtrt.ErrorResolve("QPushButton", "isFlat", args)
 }

}

// <= body block end

