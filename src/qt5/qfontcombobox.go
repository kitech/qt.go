package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtWidgets/qfontcombobox.h
// dst-file: /src/widgets/qfontcombobox.go
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
// class sizeof(QFontComboBox)=1
type QFontComboBox struct {
  /*qbase*/ QComboBox;
  qclsinst uint64 /* *mut c_void*/;
//  _currentFontChanged QFontComboBox_currentFontChanged_signal;
}


func NewQFontComboBox(args ...interface{})() {
}


func (this *QFontComboBox) FreeQFontComboBox(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QFontComboBox", "~QFontComboBox", args)
 }

}


func (this *QFontComboBox) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QFontComboBox10metaObjectEv
  default:
    qtrt.ErrorResolve("QFontComboBox", "metaObject", args)
 }

}


func (this *QFontComboBox) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QFontComboBox8sizeHintEv
  default:
    qtrt.ErrorResolve("QFontComboBox", "sizeHint", args)
 }

}


func (this *QFontComboBox) currentFont(args ...interface{}) () {
  // currentFont()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QFontComboBox11currentFontEv
  default:
    qtrt.ErrorResolve("QFontComboBox", "currentFont", args)
 }

}


func (this *QFontComboBox) setCurrentFont(args ...interface{}) () {
  // setCurrentFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QFontComboBox14setCurrentFontERK5QFont
  default:
    qtrt.ErrorResolve("QFontComboBox", "setCurrentFont", args)
 }

}

// <= body block end

