package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qstylehints.h
// dst-file: /src/gui/qstylehints.go
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
// class sizeof(QStyleHints)=1
type QStyleHints struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _startDragDistanceChanged QStyleHints_startDragDistanceChanged_signal;
//  _startDragTimeChanged QStyleHints_startDragTimeChanged_signal;
//  _mouseDoubleClickIntervalChanged QStyleHints_mouseDoubleClickIntervalChanged_signal;
//  _cursorFlashTimeChanged QStyleHints_cursorFlashTimeChanged_signal;
//  _keyboardInputIntervalChanged QStyleHints_keyboardInputIntervalChanged_signal;
}


func (this *QStyleHints) setMouseDoubleClickInterval(args ...interface{}) () {
  // setMouseDoubleClickInterval(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QStyleHints27setMouseDoubleClickIntervalEi
  default:
    qtrt.ErrorResolve("QStyleHints", "setMouseDoubleClickInterval", args)
 }

}


func (this *QStyleHints) mousePressAndHoldInterval(args ...interface{}) () {
  // mousePressAndHoldInterval()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QStyleHints25mousePressAndHoldIntervalEv
  default:
    qtrt.ErrorResolve("QStyleHints", "mousePressAndHoldInterval", args)
 }

}


func (this *QStyleHints) passwordMaskDelay(args ...interface{}) () {
  // passwordMaskDelay()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QStyleHints17passwordMaskDelayEv
  default:
    qtrt.ErrorResolve("QStyleHints", "passwordMaskDelay", args)
 }

}


func (this *QStyleHints) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QStyleHints10metaObjectEv
  default:
    qtrt.ErrorResolve("QStyleHints", "metaObject", args)
 }

}


func (this *QStyleHints) setKeyboardInputInterval(args ...interface{}) () {
  // setKeyboardInputInterval(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QStyleHints24setKeyboardInputIntervalEi
  default:
    qtrt.ErrorResolve("QStyleHints", "setKeyboardInputInterval", args)
 }

}


func NewQStyleHints(args ...interface{})() {
}


func (this *QStyleHints) showIsFullScreen(args ...interface{}) () {
  // showIsFullScreen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QStyleHints16showIsFullScreenEv
  default:
    qtrt.ErrorResolve("QStyleHints", "showIsFullScreen", args)
 }

}


func (this *QStyleHints) useRtlExtensions(args ...interface{}) () {
  // useRtlExtensions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QStyleHints16useRtlExtensionsEv
  default:
    qtrt.ErrorResolve("QStyleHints", "useRtlExtensions", args)
 }

}


func (this *QStyleHints) setStartDragDistance(args ...interface{}) () {
  // setStartDragDistance(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QStyleHints20setStartDragDistanceEi
  default:
    qtrt.ErrorResolve("QStyleHints", "setStartDragDistance", args)
 }

}


func (this *QStyleHints) setFocusOnTouchRelease(args ...interface{}) () {
  // setFocusOnTouchRelease()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QStyleHints22setFocusOnTouchReleaseEv
  default:
    qtrt.ErrorResolve("QStyleHints", "setFocusOnTouchRelease", args)
 }

}


func (this *QStyleHints) startDragVelocity(args ...interface{}) () {
  // startDragVelocity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QStyleHints17startDragVelocityEv
  default:
    qtrt.ErrorResolve("QStyleHints", "startDragVelocity", args)
 }

}


func (this *QStyleHints) startDragTime(args ...interface{}) () {
  // startDragTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QStyleHints13startDragTimeEv
  default:
    qtrt.ErrorResolve("QStyleHints", "startDragTime", args)
 }

}


func (this *QStyleHints) keyboardInputInterval(args ...interface{}) () {
  // keyboardInputInterval()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QStyleHints21keyboardInputIntervalEv
  default:
    qtrt.ErrorResolve("QStyleHints", "keyboardInputInterval", args)
 }

}


func (this *QStyleHints) setStartDragTime(args ...interface{}) () {
  // setStartDragTime(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QStyleHints16setStartDragTimeEi
  default:
    qtrt.ErrorResolve("QStyleHints", "setStartDragTime", args)
 }

}


func (this *QStyleHints) setCursorFlashTime(args ...interface{}) () {
  // setCursorFlashTime(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QStyleHints18setCursorFlashTimeEi
  default:
    qtrt.ErrorResolve("QStyleHints", "setCursorFlashTime", args)
 }

}


func (this *QStyleHints) cursorFlashTime(args ...interface{}) () {
  // cursorFlashTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QStyleHints15cursorFlashTimeEv
  default:
    qtrt.ErrorResolve("QStyleHints", "cursorFlashTime", args)
 }

}


func (this *QStyleHints) passwordMaskCharacter(args ...interface{}) () {
  // passwordMaskCharacter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QStyleHints21passwordMaskCharacterEv
  default:
    qtrt.ErrorResolve("QStyleHints", "passwordMaskCharacter", args)
 }

}


func (this *QStyleHints) keyboardAutoRepeatRate(args ...interface{}) () {
  // keyboardAutoRepeatRate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QStyleHints22keyboardAutoRepeatRateEv
  default:
    qtrt.ErrorResolve("QStyleHints", "keyboardAutoRepeatRate", args)
 }

}


func (this *QStyleHints) startDragDistance(args ...interface{}) () {
  // startDragDistance()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QStyleHints17startDragDistanceEv
  default:
    qtrt.ErrorResolve("QStyleHints", "startDragDistance", args)
 }

}


func (this *QStyleHints) fontSmoothingGamma(args ...interface{}) () {
  // fontSmoothingGamma()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QStyleHints18fontSmoothingGammaEv
  default:
    qtrt.ErrorResolve("QStyleHints", "fontSmoothingGamma", args)
 }

}


func (this *QStyleHints) singleClickActivation(args ...interface{}) () {
  // singleClickActivation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QStyleHints21singleClickActivationEv
  default:
    qtrt.ErrorResolve("QStyleHints", "singleClickActivation", args)
 }

}


func (this *QStyleHints) mouseDoubleClickInterval(args ...interface{}) () {
  // mouseDoubleClickInterval()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QStyleHints24mouseDoubleClickIntervalEv
  default:
    qtrt.ErrorResolve("QStyleHints", "mouseDoubleClickInterval", args)
 }

}

// <= body block end

