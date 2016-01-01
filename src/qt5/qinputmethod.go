package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qinputmethod.h
// dst-file: /src/gui/qinputmethod.go
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
// class sizeof(QInputMethod)=1
type QInputMethod struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _cursorRectangleChanged QInputMethod_cursorRectangleChanged_signal;
//  _localeChanged QInputMethod_localeChanged_signal;
//  _inputDirectionChanged QInputMethod_inputDirectionChanged_signal;
//  _animatingChanged QInputMethod_animatingChanged_signal;
//  _keyboardRectangleChanged QInputMethod_keyboardRectangleChanged_signal;
//  _visibleChanged QInputMethod_visibleChanged_signal;
}


func (this *QInputMethod) inputItemRectangle(args ...interface{}) () {
  // inputItemRectangle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QInputMethod18inputItemRectangleEv
  default:
    qtrt.ErrorResolve("QInputMethod", "inputItemRectangle", args)
 }

}


func (this *QInputMethod) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QInputMethod10metaObjectEv
  default:
    qtrt.ErrorResolve("QInputMethod", "metaObject", args)
 }

}


func (this *QInputMethod) inputItemTransform(args ...interface{}) () {
  // inputItemTransform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QInputMethod18inputItemTransformEv
  default:
    qtrt.ErrorResolve("QInputMethod", "inputItemTransform", args)
 }

}


func (this *QInputMethod) hide(args ...interface{}) () {
  // hide()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QInputMethod4hideEv
  default:
    qtrt.ErrorResolve("QInputMethod", "hide", args)
 }

}


func (this *QInputMethod) keyboardRectangle(args ...interface{}) () {
  // keyboardRectangle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QInputMethod17keyboardRectangleEv
  default:
    qtrt.ErrorResolve("QInputMethod", "keyboardRectangle", args)
 }

}


func (this *QInputMethod) show(args ...interface{}) () {
  // show()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QInputMethod4showEv
  default:
    qtrt.ErrorResolve("QInputMethod", "show", args)
 }

}


func NewQInputMethod(args ...interface{})() {
}


func (this *QInputMethod) isAnimating(args ...interface{}) () {
  // isAnimating()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QInputMethod11isAnimatingEv
  default:
    qtrt.ErrorResolve("QInputMethod", "isAnimating", args)
 }

}


func (this *QInputMethod) setVisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QInputMethod10setVisibleEb
  default:
    qtrt.ErrorResolve("QInputMethod", "setVisible", args)
 }

}


func (this *QInputMethod) setInputItemRectangle(args ...interface{}) () {
  // setInputItemRectangle(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QInputMethod21setInputItemRectangleERK6QRectF
  default:
    qtrt.ErrorResolve("QInputMethod", "setInputItemRectangle", args)
 }

}


func (this *QInputMethod) commit(args ...interface{}) () {
  // commit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QInputMethod6commitEv
  default:
    qtrt.ErrorResolve("QInputMethod", "commit", args)
 }

}


func (this *QInputMethod) setInputItemTransform(args ...interface{}) () {
  // setInputItemTransform(const class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QInputMethod21setInputItemTransformERK10QTransform
  default:
    qtrt.ErrorResolve("QInputMethod", "setInputItemTransform", args)
 }

}


func (this *QInputMethod) cursorRectangle(args ...interface{}) () {
  // cursorRectangle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QInputMethod15cursorRectangleEv
  default:
    qtrt.ErrorResolve("QInputMethod", "cursorRectangle", args)
 }

}


func (this *QInputMethod) isVisible(args ...interface{}) () {
  // isVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QInputMethod9isVisibleEv
  default:
    qtrt.ErrorResolve("QInputMethod", "isVisible", args)
 }

}


func (this *QInputMethod) FreeQInputMethod(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QInputMethod", "~QInputMethod", args)
 }

}


func (this *QInputMethod) locale(args ...interface{}) () {
  // locale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QInputMethod6localeEv
  default:
    qtrt.ErrorResolve("QInputMethod", "locale", args)
 }

}


func (this *QInputMethod) reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QInputMethod5resetEv
  default:
    qtrt.ErrorResolve("QInputMethod", "reset", args)
 }

}

// <= body block end

