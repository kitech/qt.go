package qtwidgets
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
// src-file: /QtWidgets/qpushbutton.h
// dst-file: /src/widgets/qpushbutton.go
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
import "qtgui"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QPushButton::showMenu();
extern void C_ZN11QPushButton8showMenuEv(void* qthis); // 4
  // proto:  void QPushButton::setMenu(QMenu * menu);
extern void C_ZN11QPushButton7setMenuEP5QMenu(void* qthis, void* arg0); // 4
  // proto:  void QPushButton::~QPushButton();
extern void C_ZN11QPushButtonD2Ev(void* qthis); // 4
  // proto:  void QPushButton::setAutoDefault(bool );
extern void C_ZN11QPushButton14setAutoDefaultEb(void* qthis, bool arg0); // 4
  // proto:  void QPushButton::setFlat(bool );
extern void C_ZN11QPushButton7setFlatEb(void* qthis, bool arg0); // 4
  // proto:  QMenu * QPushButton::menu();
extern void* C_ZNK11QPushButton4menuEv(void* qthis); // 4
  // proto:  void QPushButton::QPushButton(const QString & text, QWidget * parent);
extern void* C_ZN11QPushButtonC2ERK7QStringP7QWidget(void* arg0, void* arg1); // 3
  // proto:  void QPushButton::QPushButton(const QIcon & icon, const QString & text, QWidget * parent);
extern void* C_ZN11QPushButtonC2ERK5QIconRK7QStringP7QWidget(void* arg0, void* arg1, void* arg2); // 3
  // proto:  void QPushButton::QPushButton(QWidget * parent);
extern void* C_ZN11QPushButtonC2EP7QWidget(void* arg0); // 3
  // proto:  bool QPushButton::isFlat();
extern bool C_ZNK11QPushButton6isFlatEv(void* qthis); // 4
  // proto:  bool QPushButton::isDefault();
extern bool C_ZNK11QPushButton9isDefaultEv(void* qthis); // 4
  // proto:  void QPushButton::setDefault(bool );
extern void C_ZN11QPushButton10setDefaultEb(void* qthis, bool arg0); // 4
  // proto:  const QMetaObject * QPushButton::metaObject();
extern void C_ZNK11QPushButton10metaObjectEv(void* qthis); // 4
  // proto:  QSize QPushButton::sizeHint();
extern void* C_ZNK11QPushButton8sizeHintEv(void* qthis); // 4
  // proto:  QSize QPushButton::minimumSizeHint();
extern void* C_ZNK11QPushButton15minimumSizeHintEv(void* qthis); // 4
  // proto:  bool QPushButton::autoDefault();
extern bool C_ZNK11QPushButton11autoDefaultEv(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QPushButton)=1
type QPushButton struct {
  /*qbase*/ QAbstractButton;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// showMenu()
func (this *QPushButton) Showmenu(args ...interface{}) () {
  // showMenu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPushButton8showMenuEv
    // invoke: void showMenu()
    C.C_ZN11QPushButton8showMenuEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPushButton", "showMenu", args)
  }

  return
}

// setMenu(class QMenu *)
func (this *QPushButton) Setmenu(args ...interface{}) () {
  // setMenu(class QMenu *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMenu{}) // "QMenu *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPushButton7setMenuEP5QMenu
    // invoke: void setMenu(class QMenu *)
    var arg0 = args[0].(*QMenu).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QPushButton7setMenuEP5QMenu(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPushButton", "setMenu", args)
  }

  return
}

// ~QPushButton()
func (this *QPushButton) Freeqpushbutton(args ...interface{}) () {
  // ~QPushButton()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPushButtonD0Ev
    // invoke: void ~QPushButton()
    C.C_ZN11QPushButtonD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPushButton", "~QPushButton", args)
  }

  return
}

// setAutoDefault(_Bool)
func (this *QPushButton) Setautodefault(args ...interface{}) () {
  // setAutoDefault(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPushButton14setAutoDefaultEb
    // invoke: void setAutoDefault(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QPushButton14setAutoDefaultEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPushButton", "setAutoDefault", args)
  }

  return
}

// setFlat(_Bool)
func (this *QPushButton) Setflat(args ...interface{}) () {
  // setFlat(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPushButton7setFlatEb
    // invoke: void setFlat(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QPushButton7setFlatEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPushButton", "setFlat", args)
  }

  return
}

// menu()
func (this *QPushButton) Menu(args ...interface{}) (ret interface{}) {
  // menu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPushButton4menuEv
    // invoke: QMenu * menu()
    var ret0 = C.C_ZNK11QPushButton4menuEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMenu{}) // "QMenu *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPushButton", "menu", args)
  }

  return
}

// QPushButton(const class QString &, class QWidget *)
func NewQPushButton(args ...interface{}) *QPushButton {
  // QPushButton(const class QString &, class QWidget *)
  // QPushButton(const class QIcon &, const class QString &, class QWidget *)
  // QPushButton(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtgui.QIcon{}) // "const QIcon &"
  vtys[1][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][2] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPushButtonC1ERK7QStringP7QWidget
    // invoke: void QPushButton(const class QString &, class QWidget *)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QPushButtonC2ERK7QStringP7QWidget(arg0, arg1)
    return &QPushButton{Qclsinst:qthis}
  case 1:
    // invoke: _ZN11QPushButtonC1ERK5QIconRK7QStringP7QWidget
    // invoke: void QPushButton(const class QIcon &, const class QString &, class QWidget *)
    var arg0 = args[0].(*qtgui.QIcon).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QWidget).Qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QPushButtonC2ERK5QIconRK7QStringP7QWidget(arg0, arg1, arg2)
    return &QPushButton{Qclsinst:qthis}
  case 2:
    // invoke: _ZN11QPushButtonC1EP7QWidget
    // invoke: void QPushButton(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QPushButtonC2EP7QWidget(arg0)
    return &QPushButton{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QPushButton", "QPushButton", args)
  }

  return nil // QPushButton{Qclsinst:qthis}
}

// isFlat()
func (this *QPushButton) Isflat(args ...interface{}) (ret interface{}) {
  // isFlat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPushButton6isFlatEv
    // invoke: bool isFlat()
    var ret0 = C.C_ZNK11QPushButton6isFlatEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPushButton", "isFlat", args)
  }

  return
}

// isDefault()
func (this *QPushButton) Isdefault(args ...interface{}) (ret interface{}) {
  // isDefault()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPushButton9isDefaultEv
    // invoke: bool isDefault()
    var ret0 = C.C_ZNK11QPushButton9isDefaultEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPushButton", "isDefault", args)
  }

  return
}

// setDefault(_Bool)
func (this *QPushButton) Setdefault(args ...interface{}) () {
  // setDefault(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPushButton10setDefaultEb
    // invoke: void setDefault(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QPushButton10setDefaultEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPushButton", "setDefault", args)
  }

  return
}

// metaObject()
func (this *QPushButton) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPushButton10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK11QPushButton10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPushButton", "metaObject", args)
  }

  return
}

// sizeHint()
func (this *QPushButton) Sizehint(args ...interface{}) (ret interface{}) {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPushButton8sizeHintEv
    // invoke: QSize sizeHint()
    var ret0 = C.C_ZNK11QPushButton8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPushButton", "sizeHint", args)
  }

  return
}

// minimumSizeHint()
func (this *QPushButton) Minimumsizehint(args ...interface{}) (ret interface{}) {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPushButton15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    var ret0 = C.C_ZNK11QPushButton15minimumSizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPushButton", "minimumSizeHint", args)
  }

  return
}

// autoDefault()
func (this *QPushButton) Autodefault(args ...interface{}) (ret interface{}) {
  // autoDefault()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPushButton11autoDefaultEv
    // invoke: bool autoDefault()
    var ret0 = C.C_ZNK11QPushButton11autoDefaultEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPushButton", "autoDefault", args)
  }

  return
}

// <= body block end

