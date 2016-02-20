package qtgui
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  bool QStyleHints::setFocusOnTouchRelease();
extern bool C_ZNK11QStyleHints22setFocusOnTouchReleaseEv(void* qthis); // 4
  // proto:  int QStyleHints::startDragDistance();
extern int32_t C_ZNK11QStyleHints17startDragDistanceEv(void* qthis); // 4
  // proto:  int QStyleHints::keyboardAutoRepeatRate();
extern int32_t C_ZNK11QStyleHints22keyboardAutoRepeatRateEv(void* qthis); // 4
  // proto:  int QStyleHints::passwordMaskDelay();
extern int32_t C_ZNK11QStyleHints17passwordMaskDelayEv(void* qthis); // 4
  // proto:  void QStyleHints::setMouseDoubleClickInterval(int mouseDoubleClickInterval);
extern void C_ZN11QStyleHints27setMouseDoubleClickIntervalEi(void* qthis, int32_t arg0); // 4
  // proto:  int QStyleHints::cursorFlashTime();
extern int32_t C_ZNK11QStyleHints15cursorFlashTimeEv(void* qthis); // 4
  // proto:  int QStyleHints::mousePressAndHoldInterval();
extern int32_t C_ZNK11QStyleHints25mousePressAndHoldIntervalEv(void* qthis); // 4
  // proto:  int QStyleHints::startDragVelocity();
extern int32_t C_ZNK11QStyleHints17startDragVelocityEv(void* qthis); // 4
  // proto:  bool QStyleHints::useRtlExtensions();
extern bool C_ZNK11QStyleHints16useRtlExtensionsEv(void* qthis); // 4
  // proto:  void QStyleHints::setKeyboardInputInterval(int keyboardInputInterval);
extern void C_ZN11QStyleHints24setKeyboardInputIntervalEi(void* qthis, int32_t arg0); // 4
  // proto:  Qt::TabFocusBehavior QStyleHints::tabFocusBehavior();
extern void C_ZNK11QStyleHints16tabFocusBehaviorEv(void* qthis); // 4
  // proto:  void QStyleHints::setStartDragTime(int startDragTime);
extern void C_ZN11QStyleHints16setStartDragTimeEi(void* qthis, int32_t arg0); // 4
  // proto:  void QStyleHints::setCursorFlashTime(int cursorFlashTime);
extern void C_ZN11QStyleHints18setCursorFlashTimeEi(void* qthis, int32_t arg0); // 4
  // proto:  qreal QStyleHints::fontSmoothingGamma();
extern double C_ZNK11QStyleHints18fontSmoothingGammaEv(void* qthis); // 4
  // proto:  void QStyleHints::setStartDragDistance(int startDragDistance);
extern void C_ZN11QStyleHints20setStartDragDistanceEi(void* qthis, int32_t arg0); // 4
  // proto:  const QMetaObject * QStyleHints::metaObject();
extern void C_ZNK11QStyleHints10metaObjectEv(void* qthis); // 4
  // proto:  int QStyleHints::keyboardInputInterval();
extern int32_t C_ZNK11QStyleHints21keyboardInputIntervalEv(void* qthis); // 4
  // proto:  bool QStyleHints::showIsFullScreen();
extern bool C_ZNK11QStyleHints16showIsFullScreenEv(void* qthis); // 4
  // proto:  int QStyleHints::startDragTime();
extern int32_t C_ZNK11QStyleHints13startDragTimeEv(void* qthis); // 4
  // proto:  bool QStyleHints::singleClickActivation();
extern bool C_ZNK11QStyleHints21singleClickActivationEv(void* qthis); // 4
  // proto:  QChar QStyleHints::passwordMaskCharacter();
extern void* C_ZNK11QStyleHints21passwordMaskCharacterEv(void* qthis); // 4
  // proto:  int QStyleHints::mouseDoubleClickInterval();
extern int32_t C_ZNK11QStyleHints24mouseDoubleClickIntervalEv(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QStyleHints)=1
type QStyleHints struct {
  /*qbase*/ qtcore.QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _startDragDistanceChanged QStyleHints_startDragDistanceChanged_signal;
//  _startDragTimeChanged QStyleHints_startDragTimeChanged_signal;
//  _mouseDoubleClickIntervalChanged QStyleHints_mouseDoubleClickIntervalChanged_signal;
//  _cursorFlashTimeChanged QStyleHints_cursorFlashTimeChanged_signal;
//  _keyboardInputIntervalChanged QStyleHints_keyboardInputIntervalChanged_signal;
}

// setFocusOnTouchRelease()
func (this *QStyleHints) Setfocusontouchrelease(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QStyleHints22setFocusOnTouchReleaseEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStyleHints", "setFocusOnTouchRelease", args)
  }

  return
}

// startDragDistance()
func (this *QStyleHints) Startdragdistance(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QStyleHints17startDragDistanceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStyleHints", "startDragDistance", args)
  }

  return
}

// keyboardAutoRepeatRate()
func (this *QStyleHints) Keyboardautorepeatrate(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QStyleHints22keyboardAutoRepeatRateEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStyleHints", "keyboardAutoRepeatRate", args)
  }

  return
}

// passwordMaskDelay()
func (this *QStyleHints) Passwordmaskdelay(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QStyleHints17passwordMaskDelayEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStyleHints", "passwordMaskDelay", args)
  }

  return
}

// setMouseDoubleClickInterval(int)
func (this *QStyleHints) Setmousedoubleclickinterval(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QStyleHints27setMouseDoubleClickIntervalEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStyleHints", "setMouseDoubleClickInterval", args)
  }

  return
}

// cursorFlashTime()
func (this *QStyleHints) Cursorflashtime(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QStyleHints15cursorFlashTimeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStyleHints", "cursorFlashTime", args)
  }

  return
}

// mousePressAndHoldInterval()
func (this *QStyleHints) Mousepressandholdinterval(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QStyleHints25mousePressAndHoldIntervalEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStyleHints", "mousePressAndHoldInterval", args)
  }

  return
}

// startDragVelocity()
func (this *QStyleHints) Startdragvelocity(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QStyleHints17startDragVelocityEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStyleHints", "startDragVelocity", args)
  }

  return
}

// useRtlExtensions()
func (this *QStyleHints) Usertlextensions(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QStyleHints16useRtlExtensionsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStyleHints", "useRtlExtensions", args)
  }

  return
}

// setKeyboardInputInterval(int)
func (this *QStyleHints) Setkeyboardinputinterval(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QStyleHints24setKeyboardInputIntervalEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStyleHints", "setKeyboardInputInterval", args)
  }

  return
}

// tabFocusBehavior()
func (this *QStyleHints) Tabfocusbehavior(args ...interface{}) () {
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
    C.C_ZNK11QStyleHints16tabFocusBehaviorEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "tabFocusBehavior", args)
  }

  return
}

// setStartDragTime(int)
func (this *QStyleHints) Setstartdragtime(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QStyleHints16setStartDragTimeEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStyleHints", "setStartDragTime", args)
  }

  return
}

// setCursorFlashTime(int)
func (this *QStyleHints) Setcursorflashtime(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QStyleHints18setCursorFlashTimeEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStyleHints", "setCursorFlashTime", args)
  }

  return
}

// fontSmoothingGamma()
func (this *QStyleHints) Fontsmoothinggamma(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QStyleHints18fontSmoothingGammaEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStyleHints", "fontSmoothingGamma", args)
  }

  return
}

// setStartDragDistance(int)
func (this *QStyleHints) Setstartdragdistance(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QStyleHints20setStartDragDistanceEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStyleHints", "setStartDragDistance", args)
  }

  return
}

// metaObject()
func (this *QStyleHints) Metaobject(args ...interface{}) () {
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
    C.C_ZNK11QStyleHints10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHints", "metaObject", args)
  }

  return
}

// keyboardInputInterval()
func (this *QStyleHints) Keyboardinputinterval(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QStyleHints21keyboardInputIntervalEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStyleHints", "keyboardInputInterval", args)
  }

  return
}

// showIsFullScreen()
func (this *QStyleHints) Showisfullscreen(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QStyleHints16showIsFullScreenEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStyleHints", "showIsFullScreen", args)
  }

  return
}

// startDragTime()
func (this *QStyleHints) Startdragtime(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QStyleHints13startDragTimeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStyleHints", "startDragTime", args)
  }

  return
}

// singleClickActivation()
func (this *QStyleHints) Singleclickactivation(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QStyleHints21singleClickActivationEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStyleHints", "singleClickActivation", args)
  }

  return
}

// passwordMaskCharacter()
func (this *QStyleHints) Passwordmaskcharacter(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QStyleHints21passwordMaskCharacterEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QChar{}) // "QChar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStyleHints", "passwordMaskCharacter", args)
  }

  return
}

// mouseDoubleClickInterval()
func (this *QStyleHints) Mousedoubleclickinterval(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QStyleHints24mouseDoubleClickIntervalEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStyleHints", "mouseDoubleClickInterval", args)
  }

  return
}

// <= body block end

