package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
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
  // proto:  void QStyleHints::setMouseDoubleClickInterval(int mouseDoubleClickInterval);
extern void _ZN11QStyleHints27setMouseDoubleClickIntervalEi(void* qthis, int arg0);
  // proto:  int QStyleHints::mousePressAndHoldInterval();
extern void _ZNK11QStyleHints25mousePressAndHoldIntervalEv(void* qthis);
  // proto:  int QStyleHints::passwordMaskDelay();
extern void _ZNK11QStyleHints17passwordMaskDelayEv(void* qthis);
  // proto:  const QMetaObject * QStyleHints::metaObject();
extern void _ZNK11QStyleHints10metaObjectEv(void* qthis);
  // proto:  void QStyleHints::setKeyboardInputInterval(int keyboardInputInterval);
extern void _ZN11QStyleHints24setKeyboardInputIntervalEi(void* qthis, int arg0);
  // proto:  void QStyleHints::QStyleHints();
extern void* dector_ZN11QStyleHintsC1Ev();
extern void _ZN11QStyleHintsC1Ev(void* qthis);
  // proto:  bool QStyleHints::showIsFullScreen();
extern void _ZNK11QStyleHints16showIsFullScreenEv(void* qthis);
  // proto:  bool QStyleHints::useRtlExtensions();
extern void _ZNK11QStyleHints16useRtlExtensionsEv(void* qthis);
  // proto:  void QStyleHints::setStartDragDistance(int startDragDistance);
extern void _ZN11QStyleHints20setStartDragDistanceEi(void* qthis, int arg0);
  // proto:  bool QStyleHints::setFocusOnTouchRelease();
extern void _ZNK11QStyleHints22setFocusOnTouchReleaseEv(void* qthis);
  // proto:  int QStyleHints::startDragVelocity();
extern void _ZNK11QStyleHints17startDragVelocityEv(void* qthis);
  // proto:  int QStyleHints::startDragTime();
extern void _ZNK11QStyleHints13startDragTimeEv(void* qthis);
  // proto:  int QStyleHints::keyboardInputInterval();
extern void _ZNK11QStyleHints21keyboardInputIntervalEv(void* qthis);
  // proto:  void QStyleHints::setStartDragTime(int startDragTime);
extern void _ZN11QStyleHints16setStartDragTimeEi(void* qthis, int arg0);
  // proto:  void QStyleHints::setCursorFlashTime(int cursorFlashTime);
extern void _ZN11QStyleHints18setCursorFlashTimeEi(void* qthis, int arg0);
  // proto:  int QStyleHints::cursorFlashTime();
extern void _ZNK11QStyleHints15cursorFlashTimeEv(void* qthis);
  // proto:  QChar QStyleHints::passwordMaskCharacter();
extern void _ZNK11QStyleHints21passwordMaskCharacterEv(void* qthis);
  // proto:  int QStyleHints::keyboardAutoRepeatRate();
extern void _ZNK11QStyleHints22keyboardAutoRepeatRateEv(void* qthis);
  // proto:  int QStyleHints::startDragDistance();
extern void _ZNK11QStyleHints17startDragDistanceEv(void* qthis);
  // proto:  qreal QStyleHints::fontSmoothingGamma();
extern void _ZNK11QStyleHints18fontSmoothingGammaEv(void* qthis);
  // proto:  bool QStyleHints::singleClickActivation();
extern void _ZNK11QStyleHints21singleClickActivationEv(void* qthis);
  // proto:  int QStyleHints::mouseDoubleClickInterval();
extern void _ZNK11QStyleHints24mouseDoubleClickIntervalEv(void* qthis);
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
  qclsinst uint64 /* *mut c_void*/;
//  _startDragDistanceChanged QStyleHints_startDragDistanceChanged_signal;
//  _startDragTimeChanged QStyleHints_startDragTimeChanged_signal;
//  _mouseDoubleClickIntervalChanged QStyleHints_mouseDoubleClickIntervalChanged_signal;
//  _cursorFlashTimeChanged QStyleHints_cursorFlashTimeChanged_signal;
//  _keyboardInputIntervalChanged QStyleHints_keyboardInputIntervalChanged_signal;
}

  // proto:  void QStyleHints::setMouseDoubleClickInterval(int mouseDoubleClickInterval);
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
    C._ZN11QStyleHints27setMouseDoubleClickIntervalEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStyleHints", "setMouseDoubleClickInterval", args)
  }

}

  // proto:  int QStyleHints::mousePressAndHoldInterval();
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
    C._ZNK11QStyleHints25mousePressAndHoldIntervalEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "mousePressAndHoldInterval", args)
  }

}

  // proto:  int QStyleHints::passwordMaskDelay();
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
    C._ZNK11QStyleHints17passwordMaskDelayEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "passwordMaskDelay", args)
  }

}

  // proto:  const QMetaObject * QStyleHints::metaObject();
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
    C._ZNK11QStyleHints10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "metaObject", args)
  }

}

  // proto:  void QStyleHints::setKeyboardInputInterval(int keyboardInputInterval);
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
    C._ZN11QStyleHints24setKeyboardInputIntervalEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStyleHints", "setKeyboardInputInterval", args)
  }

}

  // proto:  void QStyleHints::QStyleHints();
func NewQStyleHints(args ...interface{}) QStyleHints {
  return QStyleHints{}
}

  // proto:  bool QStyleHints::showIsFullScreen();
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
    C._ZNK11QStyleHints16showIsFullScreenEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "showIsFullScreen", args)
  }

}

  // proto:  bool QStyleHints::useRtlExtensions();
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
    C._ZNK11QStyleHints16useRtlExtensionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "useRtlExtensions", args)
  }

}

  // proto:  void QStyleHints::setStartDragDistance(int startDragDistance);
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
    C._ZN11QStyleHints20setStartDragDistanceEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStyleHints", "setStartDragDistance", args)
  }

}

  // proto:  bool QStyleHints::setFocusOnTouchRelease();
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
    C._ZNK11QStyleHints22setFocusOnTouchReleaseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "setFocusOnTouchRelease", args)
  }

}

  // proto:  int QStyleHints::startDragVelocity();
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
    C._ZNK11QStyleHints17startDragVelocityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "startDragVelocity", args)
  }

}

  // proto:  int QStyleHints::startDragTime();
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
    C._ZNK11QStyleHints13startDragTimeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "startDragTime", args)
  }

}

  // proto:  int QStyleHints::keyboardInputInterval();
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
    C._ZNK11QStyleHints21keyboardInputIntervalEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "keyboardInputInterval", args)
  }

}

  // proto:  void QStyleHints::setStartDragTime(int startDragTime);
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
    C._ZN11QStyleHints16setStartDragTimeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStyleHints", "setStartDragTime", args)
  }

}

  // proto:  void QStyleHints::setCursorFlashTime(int cursorFlashTime);
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
    C._ZN11QStyleHints18setCursorFlashTimeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStyleHints", "setCursorFlashTime", args)
  }

}

  // proto:  int QStyleHints::cursorFlashTime();
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
    C._ZNK11QStyleHints15cursorFlashTimeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "cursorFlashTime", args)
  }

}

  // proto:  QChar QStyleHints::passwordMaskCharacter();
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
    C._ZNK11QStyleHints21passwordMaskCharacterEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "passwordMaskCharacter", args)
  }

}

  // proto:  int QStyleHints::keyboardAutoRepeatRate();
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
    C._ZNK11QStyleHints22keyboardAutoRepeatRateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "keyboardAutoRepeatRate", args)
  }

}

  // proto:  int QStyleHints::startDragDistance();
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
    C._ZNK11QStyleHints17startDragDistanceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "startDragDistance", args)
  }

}

  // proto:  qreal QStyleHints::fontSmoothingGamma();
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
    C._ZNK11QStyleHints18fontSmoothingGammaEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "fontSmoothingGamma", args)
  }

}

  // proto:  bool QStyleHints::singleClickActivation();
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
    C._ZNK11QStyleHints21singleClickActivationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "singleClickActivation", args)
  }

}

  // proto:  int QStyleHints::mouseDoubleClickInterval();
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
    C._ZNK11QStyleHints24mouseDoubleClickIntervalEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "mouseDoubleClickInterval", args)
  }

}

// <= body block end

