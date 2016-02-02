package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtWidgets/qabstractbutton.h
// dst-file: /src/widgets/qabstractbutton.go
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
  // proto:  void QAbstractButton::QAbstractButton(QWidget * parent);
extern void* C_ZN15QAbstractButtonC2EP7QWidget(void* arg0); // 3
  // proto:  QIcon QAbstractButton::icon();
extern void* C_ZNK15QAbstractButton4iconEv(void* qthis); // 4
  // proto:  void QAbstractButton::setAutoExclusive(bool );
extern void C_ZN15QAbstractButton16setAutoExclusiveEb(void* qthis, bool arg0); // 4
  // proto:  void QAbstractButton::setChecked(bool );
extern void C_ZN15QAbstractButton10setCheckedEb(void* qthis, bool arg0); // 4
  // proto:  void QAbstractButton::setShortcut(const QKeySequence & key);
extern void C_ZN15QAbstractButton11setShortcutERK12QKeySequence(void* qthis, void* arg0); // 4
  // proto:  bool QAbstractButton::isCheckable();
extern bool C_ZNK15QAbstractButton11isCheckableEv(void* qthis); // 4
  // proto:  void QAbstractButton::setAutoRepeatInterval(int );
extern void C_ZN15QAbstractButton21setAutoRepeatIntervalEi(void* qthis, int32_t arg0); // 4
  // proto:  void QAbstractButton::animateClick(int msec);
extern void C_ZN15QAbstractButton12animateClickEi(void* qthis, int32_t arg0); // 4
  // proto:  QButtonGroup * QAbstractButton::group();
extern void* C_ZNK15QAbstractButton5groupEv(void* qthis); // 4
  // proto:  void QAbstractButton::setAutoRepeatDelay(int );
extern void C_ZN15QAbstractButton18setAutoRepeatDelayEi(void* qthis, int32_t arg0); // 4
  // proto:  QKeySequence QAbstractButton::shortcut();
extern void* C_ZNK15QAbstractButton8shortcutEv(void* qthis); // 4
  // proto:  void QAbstractButton::click();
extern void C_ZN15QAbstractButton5clickEv(void* qthis); // 4
  // proto:  int QAbstractButton::autoRepeatInterval();
extern int32_t C_ZNK15QAbstractButton18autoRepeatIntervalEv(void* qthis); // 4
  // proto:  bool QAbstractButton::isChecked();
extern bool C_ZNK15QAbstractButton9isCheckedEv(void* qthis); // 4
  // proto:  QString QAbstractButton::text();
extern void* C_ZNK15QAbstractButton4textEv(void* qthis); // 4
  // proto:  void QAbstractButton::setCheckable(bool );
extern void C_ZN15QAbstractButton12setCheckableEb(void* qthis, bool arg0); // 4
  // proto:  bool QAbstractButton::isDown();
extern bool C_ZNK15QAbstractButton6isDownEv(void* qthis); // 4
  // proto:  void QAbstractButton::setIcon(const QIcon & icon);
extern void C_ZN15QAbstractButton7setIconERK5QIcon(void* qthis, void* arg0); // 4
  // proto:  void QAbstractButton::~QAbstractButton();
extern void C_ZN15QAbstractButtonD2Ev(void* qthis); // 4
  // proto:  QSize QAbstractButton::iconSize();
extern void* C_ZNK15QAbstractButton8iconSizeEv(void* qthis); // 4
  // proto:  void QAbstractButton::setAutoRepeat(bool );
extern void C_ZN15QAbstractButton13setAutoRepeatEb(void* qthis, bool arg0); // 4
  // proto:  bool QAbstractButton::autoExclusive();
extern bool C_ZNK15QAbstractButton13autoExclusiveEv(void* qthis); // 4
  // proto:  const QMetaObject * QAbstractButton::metaObject();
extern void C_ZNK15QAbstractButton10metaObjectEv(void* qthis); // 4
  // proto:  void QAbstractButton::setText(const QString & text);
extern void C_ZN15QAbstractButton7setTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QAbstractButton::setIconSize(const QSize & size);
extern void C_ZN15QAbstractButton11setIconSizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  void QAbstractButton::toggle();
extern void C_ZN15QAbstractButton6toggleEv(void* qthis); // 4
  // proto:  void QAbstractButton::setDown(bool );
extern void C_ZN15QAbstractButton7setDownEb(void* qthis, bool arg0); // 4
  // proto:  int QAbstractButton::autoRepeatDelay();
extern int32_t C_ZNK15QAbstractButton15autoRepeatDelayEv(void* qthis); // 4
  // proto:  bool QAbstractButton::autoRepeat();
extern bool C_ZNK15QAbstractButton10autoRepeatEv(void* qthis); // 4
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

// class sizeof(QAbstractButton)=1
type QAbstractButton struct {
  /*qbase*/ QWidget;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _released QAbstractButton_released_signal;
//  _clicked QAbstractButton_clicked_signal;
//  _pressed QAbstractButton_pressed_signal;
//  _toggled QAbstractButton_toggled_signal;
}

// QAbstractButton(class QWidget *)
func NewQAbstractButton(args ...interface{}) *QAbstractButton {
  // QAbstractButton(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButtonC1EP7QWidget
    // invoke: void QAbstractButton(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QAbstractButtonC2EP7QWidget(arg0)
    return &QAbstractButton{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QAbstractButton", "QAbstractButton", args)
  }

  return nil // QAbstractButton{Qclsinst:qthis}
}

// icon()
func (this *QAbstractButton) Icon(args ...interface{}) (ret interface{}) {
  // icon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton4iconEv
    // invoke: QIcon icon()
    var ret0 = C.C_ZNK15QAbstractButton4iconEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QIcon{}) // "QIcon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractButton", "icon", args)
  }

  return
}

// setAutoExclusive(_Bool)
func (this *QAbstractButton) Setautoexclusive(args ...interface{}) () {
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
    // invoke: void setAutoExclusive(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractButton16setAutoExclusiveEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractButton", "setAutoExclusive", args)
  }

  return
}

// setChecked(_Bool)
func (this *QAbstractButton) Setchecked(args ...interface{}) () {
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
    // invoke: void setChecked(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractButton10setCheckedEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractButton", "setChecked", args)
  }

  return
}

// setShortcut(const class QKeySequence &)
func (this *QAbstractButton) Setshortcut(args ...interface{}) () {
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
    // invoke: void setShortcut(const class QKeySequence &)
    var arg0 = args[0].(*QKeySequence).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractButton11setShortcutERK12QKeySequence(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractButton", "setShortcut", args)
  }

  return
}

// isCheckable()
func (this *QAbstractButton) Ischeckable(args ...interface{}) (ret interface{}) {
  // isCheckable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton11isCheckableEv
    // invoke: bool isCheckable()
    var ret0 = C.C_ZNK15QAbstractButton11isCheckableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractButton", "isCheckable", args)
  }

  return
}

// setAutoRepeatInterval(int)
func (this *QAbstractButton) Setautorepeatinterval(args ...interface{}) () {
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
    // invoke: void setAutoRepeatInterval(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractButton21setAutoRepeatIntervalEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractButton", "setAutoRepeatInterval", args)
  }

  return
}

// animateClick(int)
func (this *QAbstractButton) Animateclick(args ...interface{}) () {
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
    // invoke: void animateClick(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractButton12animateClickEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractButton", "animateClick", args)
  }

  return
}

// group()
func (this *QAbstractButton) Group(args ...interface{}) (ret interface{}) {
  // group()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton5groupEv
    // invoke: QButtonGroup * group()
    var ret0 = C.C_ZNK15QAbstractButton5groupEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QButtonGroup{}) // "QButtonGroup *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractButton", "group", args)
  }

  return
}

// setAutoRepeatDelay(int)
func (this *QAbstractButton) Setautorepeatdelay(args ...interface{}) () {
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
    // invoke: void setAutoRepeatDelay(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractButton18setAutoRepeatDelayEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractButton", "setAutoRepeatDelay", args)
  }

  return
}

// shortcut()
func (this *QAbstractButton) Shortcut(args ...interface{}) (ret interface{}) {
  // shortcut()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton8shortcutEv
    // invoke: QKeySequence shortcut()
    var ret0 = C.C_ZNK15QAbstractButton8shortcutEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QKeySequence{}) // "QKeySequence"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractButton", "shortcut", args)
  }

  return
}

// click()
func (this *QAbstractButton) Click(args ...interface{}) () {
  // click()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton5clickEv
    // invoke: void click()
    C.C_ZN15QAbstractButton5clickEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractButton", "click", args)
  }

  return
}

// autoRepeatInterval()
func (this *QAbstractButton) Autorepeatinterval(args ...interface{}) (ret interface{}) {
  // autoRepeatInterval()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton18autoRepeatIntervalEv
    // invoke: int autoRepeatInterval()
    var ret0 = C.C_ZNK15QAbstractButton18autoRepeatIntervalEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractButton", "autoRepeatInterval", args)
  }

  return
}

// isChecked()
func (this *QAbstractButton) Ischecked(args ...interface{}) (ret interface{}) {
  // isChecked()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton9isCheckedEv
    // invoke: bool isChecked()
    var ret0 = C.C_ZNK15QAbstractButton9isCheckedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractButton", "isChecked", args)
  }

  return
}

// text()
func (this *QAbstractButton) Text(args ...interface{}) (ret interface{}) {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton4textEv
    // invoke: QString text()
    var ret0 = C.C_ZNK15QAbstractButton4textEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractButton", "text", args)
  }

  return
}

// setCheckable(_Bool)
func (this *QAbstractButton) Setcheckable(args ...interface{}) () {
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
    // invoke: void setCheckable(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractButton12setCheckableEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractButton", "setCheckable", args)
  }

  return
}

// isDown()
func (this *QAbstractButton) Isdown(args ...interface{}) (ret interface{}) {
  // isDown()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton6isDownEv
    // invoke: bool isDown()
    var ret0 = C.C_ZNK15QAbstractButton6isDownEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractButton", "isDown", args)
  }

  return
}

// setIcon(const class QIcon &)
func (this *QAbstractButton) Seticon(args ...interface{}) () {
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
    // invoke: void setIcon(const class QIcon &)
    var arg0 = args[0].(*QIcon).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractButton7setIconERK5QIcon(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractButton", "setIcon", args)
  }

  return
}

// ~QAbstractButton()
func (this *QAbstractButton) Freeqabstractbutton(args ...interface{}) () {
  // ~QAbstractButton()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButtonD0Ev
    // invoke: void ~QAbstractButton()
    C.C_ZN15QAbstractButtonD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractButton", "~QAbstractButton", args)
  }

  return
}

// iconSize()
func (this *QAbstractButton) Iconsize(args ...interface{}) (ret interface{}) {
  // iconSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton8iconSizeEv
    // invoke: QSize iconSize()
    var ret0 = C.C_ZNK15QAbstractButton8iconSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractButton", "iconSize", args)
  }

  return
}

// setAutoRepeat(_Bool)
func (this *QAbstractButton) Setautorepeat(args ...interface{}) () {
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
    // invoke: void setAutoRepeat(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractButton13setAutoRepeatEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractButton", "setAutoRepeat", args)
  }

  return
}

// autoExclusive()
func (this *QAbstractButton) Autoexclusive(args ...interface{}) (ret interface{}) {
  // autoExclusive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton13autoExclusiveEv
    // invoke: bool autoExclusive()
    var ret0 = C.C_ZNK15QAbstractButton13autoExclusiveEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractButton", "autoExclusive", args)
  }

  return
}

// metaObject()
func (this *QAbstractButton) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK15QAbstractButton10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractButton", "metaObject", args)
  }

  return
}

// setText(const class QString &)
func (this *QAbstractButton) Settext(args ...interface{}) () {
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
    // invoke: void setText(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractButton7setTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractButton", "setText", args)
  }

  return
}

// setIconSize(const class QSize &)
func (this *QAbstractButton) Seticonsize(args ...interface{}) () {
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
    // invoke: void setIconSize(const class QSize &)
    var arg0 = args[0].(*QSize).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractButton11setIconSizeERK5QSize(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractButton", "setIconSize", args)
  }

  return
}

// toggle()
func (this *QAbstractButton) Toggle(args ...interface{}) () {
  // toggle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton6toggleEv
    // invoke: void toggle()
    C.C_ZN15QAbstractButton6toggleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractButton", "toggle", args)
  }

  return
}

// setDown(_Bool)
func (this *QAbstractButton) Setdown(args ...interface{}) () {
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
    // invoke: void setDown(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractButton7setDownEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractButton", "setDown", args)
  }

  return
}

// autoRepeatDelay()
func (this *QAbstractButton) Autorepeatdelay(args ...interface{}) (ret interface{}) {
  // autoRepeatDelay()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton15autoRepeatDelayEv
    // invoke: int autoRepeatDelay()
    var ret0 = C.C_ZNK15QAbstractButton15autoRepeatDelayEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractButton", "autoRepeatDelay", args)
  }

  return
}

// autoRepeat()
func (this *QAbstractButton) Autorepeat(args ...interface{}) (ret interface{}) {
  // autoRepeat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton10autoRepeatEv
    // invoke: bool autoRepeat()
    var ret0 = C.C_ZNK15QAbstractButton10autoRepeatEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractButton", "autoRepeat", args)
  }

  return
}

// <= body block end

