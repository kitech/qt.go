package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtGui/qstylehints.h
// dst-file: /src/gui/qstylehints.go
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
  // proto:  bool QStyleHints::setFocusOnTouchRelease();
extern void C_ZNK11QStyleHints22setFocusOnTouchReleaseEv(void* qthis); // 4
  // proto:  int QStyleHints::startDragDistance();
extern void C_ZNK11QStyleHints17startDragDistanceEv(void* qthis); // 4
  // proto:  int QStyleHints::keyboardAutoRepeatRate();
extern void C_ZNK11QStyleHints22keyboardAutoRepeatRateEv(void* qthis); // 4
  // proto:  int QStyleHints::passwordMaskDelay();
extern void C_ZNK11QStyleHints17passwordMaskDelayEv(void* qthis); // 4
  // proto:  void QStyleHints::setMouseDoubleClickInterval(int mouseDoubleClickInterval);
extern void C_ZN11QStyleHints27setMouseDoubleClickIntervalEi(void* qthis, int32_t arg0); // 4
  // proto:  int QStyleHints::cursorFlashTime();
extern void C_ZNK11QStyleHints15cursorFlashTimeEv(void* qthis); // 4
  // proto:  int QStyleHints::mousePressAndHoldInterval();
extern void C_ZNK11QStyleHints25mousePressAndHoldIntervalEv(void* qthis); // 4
  // proto:  int QStyleHints::startDragVelocity();
extern void C_ZNK11QStyleHints17startDragVelocityEv(void* qthis); // 4
  // proto:  bool QStyleHints::useRtlExtensions();
extern void C_ZNK11QStyleHints16useRtlExtensionsEv(void* qthis); // 4
  // proto:  void QStyleHints::setKeyboardInputInterval(int keyboardInputInterval);
extern void C_ZN11QStyleHints24setKeyboardInputIntervalEi(void* qthis, int32_t arg0); // 4
  // proto:  Qt::TabFocusBehavior QStyleHints::tabFocusBehavior();
extern void C_ZNK11QStyleHints16tabFocusBehaviorEv(void* qthis); // 4
  // proto:  void QStyleHints::setStartDragTime(int startDragTime);
extern void C_ZN11QStyleHints16setStartDragTimeEi(void* qthis, int32_t arg0); // 4
  // proto:  void QStyleHints::setCursorFlashTime(int cursorFlashTime);
extern void C_ZN11QStyleHints18setCursorFlashTimeEi(void* qthis, int32_t arg0); // 4
  // proto:  qreal QStyleHints::fontSmoothingGamma();
extern void C_ZNK11QStyleHints18fontSmoothingGammaEv(void* qthis); // 4
  // proto:  void QStyleHints::setStartDragDistance(int startDragDistance);
extern void C_ZN11QStyleHints20setStartDragDistanceEi(void* qthis, int32_t arg0); // 4
  // proto:  const QMetaObject * QStyleHints::metaObject();
extern void C_ZNK11QStyleHints10metaObjectEv(void* qthis); // 4
  // proto:  int QStyleHints::keyboardInputInterval();
extern void C_ZNK11QStyleHints21keyboardInputIntervalEv(void* qthis); // 4
  // proto:  bool QStyleHints::showIsFullScreen();
extern void C_ZNK11QStyleHints16showIsFullScreenEv(void* qthis); // 4
  // proto:  int QStyleHints::startDragTime();
extern void C_ZNK11QStyleHints13startDragTimeEv(void* qthis); // 4
  // proto:  bool QStyleHints::singleClickActivation();
extern void C_ZNK11QStyleHints21singleClickActivationEv(void* qthis); // 4
  // proto:  QChar QStyleHints::passwordMaskCharacter();
extern void C_ZNK11QStyleHints21passwordMaskCharacterEv(void* qthis); // 4
  // proto:  int QStyleHints::mouseDoubleClickInterval();
extern void C_ZNK11QStyleHints24mouseDoubleClickIntervalEv(void* qthis); // 4
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

// class sizeof(QStyleHints)=1
type QStyleHints struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _startDragDistanceChanged QStyleHints_startDragDistanceChanged_signal;
//  _startDragTimeChanged QStyleHints_startDragTimeChanged_signal;
//  _mouseDoubleClickIntervalChanged QStyleHints_mouseDoubleClickIntervalChanged_signal;
//  _cursorFlashTimeChanged QStyleHints_cursorFlashTimeChanged_signal;
//  _keyboardInputIntervalChanged QStyleHints_keyboardInputIntervalChanged_signal;
}

// setFocusOnTouchRelease()
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
    // invoke: bool setFocusOnTouchRelease()
    C.C_ZNK11QStyleHints22setFocusOnTouchReleaseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "setFocusOnTouchRelease", args)
  }

}

// startDragDistance()
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
    // invoke: int startDragDistance()
    C.C_ZNK11QStyleHints17startDragDistanceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "startDragDistance", args)
  }

}

// keyboardAutoRepeatRate()
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
    // invoke: int keyboardAutoRepeatRate()
    C.C_ZNK11QStyleHints22keyboardAutoRepeatRateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "keyboardAutoRepeatRate", args)
  }

}

// passwordMaskDelay()
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
    // invoke: int passwordMaskDelay()
    C.C_ZNK11QStyleHints17passwordMaskDelayEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "passwordMaskDelay", args)
  }

}

// setMouseDoubleClickInterval(int)
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
    // invoke: void setMouseDoubleClickInterval(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QStyleHints27setMouseDoubleClickIntervalEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStyleHints", "setMouseDoubleClickInterval", args)
  }

}

// cursorFlashTime()
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
    // invoke: int cursorFlashTime()
    C.C_ZNK11QStyleHints15cursorFlashTimeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "cursorFlashTime", args)
  }

}

// mousePressAndHoldInterval()
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
    // invoke: int mousePressAndHoldInterval()
    C.C_ZNK11QStyleHints25mousePressAndHoldIntervalEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "mousePressAndHoldInterval", args)
  }

}

// startDragVelocity()
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
    // invoke: int startDragVelocity()
    C.C_ZNK11QStyleHints17startDragVelocityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "startDragVelocity", args)
  }

}

// useRtlExtensions()
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
    // invoke: bool useRtlExtensions()
    C.C_ZNK11QStyleHints16useRtlExtensionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "useRtlExtensions", args)
  }

}

// setKeyboardInputInterval(int)
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
    // invoke: void setKeyboardInputInterval(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QStyleHints24setKeyboardInputIntervalEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStyleHints", "setKeyboardInputInterval", args)
  }

}

// tabFocusBehavior()
func (this *QStyleHints) tabFocusBehavior(args ...interface{}) () {
  // tabFocusBehavior()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QStyleHints16tabFocusBehaviorEv
    // invoke: Qt::TabFocusBehavior tabFocusBehavior()
    C.C_ZNK11QStyleHints16tabFocusBehaviorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "tabFocusBehavior", args)
  }

}

// setStartDragTime(int)
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
    // invoke: void setStartDragTime(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QStyleHints16setStartDragTimeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStyleHints", "setStartDragTime", args)
  }

}

// setCursorFlashTime(int)
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
    // invoke: void setCursorFlashTime(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QStyleHints18setCursorFlashTimeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStyleHints", "setCursorFlashTime", args)
  }

}

// fontSmoothingGamma()
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
    // invoke: qreal fontSmoothingGamma()
    C.C_ZNK11QStyleHints18fontSmoothingGammaEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "fontSmoothingGamma", args)
  }

}

// setStartDragDistance(int)
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
    // invoke: void setStartDragDistance(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QStyleHints20setStartDragDistanceEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStyleHints", "setStartDragDistance", args)
  }

}

// metaObject()
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
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK11QStyleHints10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "metaObject", args)
  }

}

// keyboardInputInterval()
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
    // invoke: int keyboardInputInterval()
    C.C_ZNK11QStyleHints21keyboardInputIntervalEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "keyboardInputInterval", args)
  }

}

// showIsFullScreen()
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
    // invoke: bool showIsFullScreen()
    C.C_ZNK11QStyleHints16showIsFullScreenEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "showIsFullScreen", args)
  }

}

// startDragTime()
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
    // invoke: int startDragTime()
    C.C_ZNK11QStyleHints13startDragTimeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "startDragTime", args)
  }

}

// singleClickActivation()
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
    // invoke: bool singleClickActivation()
    C.C_ZNK11QStyleHints21singleClickActivationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "singleClickActivation", args)
  }

}

// passwordMaskCharacter()
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
    // invoke: QChar passwordMaskCharacter()
    C.C_ZNK11QStyleHints21passwordMaskCharacterEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "passwordMaskCharacter", args)
  }

}

// mouseDoubleClickInterval()
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
    // invoke: int mouseDoubleClickInterval()
    C.C_ZNK11QStyleHints24mouseDoubleClickIntervalEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "mouseDoubleClickInterval", args)
  }

}

// <= body block end

