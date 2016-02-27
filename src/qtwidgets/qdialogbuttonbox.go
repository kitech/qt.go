package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtWidgets/qdialogbuttonbox.h
// dst-file: /src/widgets/qdialogbuttonbox.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "runtime"
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
  // proto:  Qt::Orientation QDialogButtonBox::orientation();
extern void C_ZNK16QDialogButtonBox11orientationEv(void* qthis); // 4
  // proto:  bool QDialogButtonBox::centerButtons();
extern bool C_ZNK16QDialogButtonBox13centerButtonsEv(void* qthis); // 4
  // proto:  void QDialogButtonBox::setCenterButtons(bool center);
extern void C_ZN16QDialogButtonBox16setCenterButtonsEb(void* qthis, bool arg0); // 4
  // proto:  void QDialogButtonBox::removeButton(QAbstractButton * button);
extern void C_ZN16QDialogButtonBox12removeButtonEP15QAbstractButton(void* qthis, void* arg0); // 4
  // proto:  void QDialogButtonBox::QDialogButtonBox(QWidget * parent);
extern void* C_ZN16QDialogButtonBoxC2EP7QWidget(void* arg0); // 3
  // proto:  QDialogButtonBox::ButtonRole QDialogButtonBox::buttonRole(QAbstractButton * button);
extern void C_ZNK16QDialogButtonBox10buttonRoleEP15QAbstractButton(void* qthis, void* arg0); // 4
  // proto:  QList<QAbstractButton *> QDialogButtonBox::buttons();
extern void C_ZNK16QDialogButtonBox7buttonsEv(void* qthis); // 4
  // proto:  void QDialogButtonBox::~QDialogButtonBox();
extern void C_ZN16QDialogButtonBoxD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QDialogButtonBox::metaObject();
extern void C_ZNK16QDialogButtonBox10metaObjectEv(void* qthis); // 4
  // proto:  StandardButtons QDialogButtonBox::standardButtons();
extern void C_ZNK16QDialogButtonBox15standardButtonsEv(void* qthis); // 4
  // proto:  void QDialogButtonBox::clear();
extern void C_ZN16QDialogButtonBox5clearEv(void* qthis); // 4
  // proto:  QDialogButtonBox::StandardButton QDialogButtonBox::standardButton(QAbstractButton * button);
extern void C_ZNK16QDialogButtonBox14standardButtonEP15QAbstractButton(void* qthis, void* arg0); // 4
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
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QDialogButtonBox)=1
type QDialogButtonBox struct {
  /*qbase*/ QWidget;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _helpRequested QDialogButtonBox_helpRequested_signal;
//  _accepted QDialogButtonBox_accepted_signal;
//  _clicked QDialogButtonBox_clicked_signal;
//  _rejected QDialogButtonBox_rejected_signal;
}

// orientation()
func (this *QDialogButtonBox) Orientation(args ...interface{}) () {
  // orientation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QDialogButtonBox11orientationEv
    // invoke: Qt::Orientation orientation()
    C.C_ZNK16QDialogButtonBox11orientationEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDialogButtonBox", "orientation", args)
  }

  return
}

// centerButtons()
func (this *QDialogButtonBox) CenterButtons(args ...interface{}) (ret interface{}) {
  // centerButtons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QDialogButtonBox13centerButtonsEv
    // invoke: bool centerButtons()
    var ret0 = C.C_ZNK16QDialogButtonBox13centerButtonsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDialogButtonBox", "centerButtons", args)
  }

  return
}

// setCenterButtons(_Bool)
func (this *QDialogButtonBox) SetCenterButtons(args ...interface{}) () {
  // setCenterButtons(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QDialogButtonBox16setCenterButtonsEb
    // invoke: void setCenterButtons(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN16QDialogButtonBox16setCenterButtonsEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDialogButtonBox", "setCenterButtons", args)
  }

  return
}

// removeButton(class QAbstractButton *)
func (this *QDialogButtonBox) RemoveButton(args ...interface{}) () {
  // removeButton(class QAbstractButton *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractButton{}) // "QAbstractButton *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QDialogButtonBox12removeButtonEP15QAbstractButton
    // invoke: void removeButton(class QAbstractButton *)
    var arg0 = args[0].(*QAbstractButton).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QDialogButtonBox12removeButtonEP15QAbstractButton(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDialogButtonBox", "removeButton", args)
  }

  return
}

// QDialogButtonBox(class QWidget *)
func GcfreeQDialogButtonBox(this *QDialogButtonBox) {
  qtrt.UniverseFree(this)
}
func NewQDialogButtonBox(args ...interface{}) *QDialogButtonBox {
  // QDialogButtonBox(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QDialogButtonBoxC1EP7QWidget
    // invoke: void QDialogButtonBox(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QDialogButtonBoxC2EP7QWidget(arg0)
    this := &QDialogButtonBox{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQDialogButtonBox)
    return this
  default:
    qtrt.ErrorResolve("QDialogButtonBox", "QDialogButtonBox", args)
  }

  return nil // QDialogButtonBox{Qclsinst:qthis}
}

// buttonRole(class QAbstractButton *)
func (this *QDialogButtonBox) ButtonRole(args ...interface{}) () {
  // buttonRole(class QAbstractButton *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractButton{}) // "QAbstractButton *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QDialogButtonBox10buttonRoleEP15QAbstractButton
    // invoke: QDialogButtonBox::ButtonRole buttonRole(class QAbstractButton *)
    var arg0 = args[0].(*QAbstractButton).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK16QDialogButtonBox10buttonRoleEP15QAbstractButton(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDialogButtonBox", "buttonRole", args)
  }

  return
}

// buttons()
func (this *QDialogButtonBox) Buttons(args ...interface{}) () {
  // buttons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QDialogButtonBox7buttonsEv
    // invoke: QList<QAbstractButton *> buttons()
    C.C_ZNK16QDialogButtonBox7buttonsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDialogButtonBox", "buttons", args)
  }

  return
}

// ~QDialogButtonBox()
func (this *QDialogButtonBox) Free(args ...interface{}) () {
  // ~QDialogButtonBox()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QDialogButtonBoxD0Ev
    // invoke: void ~QDialogButtonBox()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN16QDialogButtonBoxD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QDialogButtonBox", "~QDialogButtonBox", args)
  }

  return
}

// metaObject()
func (this *QDialogButtonBox) MetaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QDialogButtonBox10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK16QDialogButtonBox10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDialogButtonBox", "metaObject", args)
  }

  return
}

// standardButtons()
func (this *QDialogButtonBox) StandardButtons(args ...interface{}) () {
  // standardButtons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QDialogButtonBox15standardButtonsEv
    // invoke: StandardButtons standardButtons()
    C.C_ZNK16QDialogButtonBox15standardButtonsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDialogButtonBox", "standardButtons", args)
  }

  return
}

// clear()
func (this *QDialogButtonBox) Clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QDialogButtonBox5clearEv
    // invoke: void clear()
    C.C_ZN16QDialogButtonBox5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDialogButtonBox", "clear", args)
  }

  return
}

// standardButton(class QAbstractButton *)
func (this *QDialogButtonBox) StandardButton(args ...interface{}) () {
  // standardButton(class QAbstractButton *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractButton{}) // "QAbstractButton *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QDialogButtonBox14standardButtonEP15QAbstractButton
    // invoke: QDialogButtonBox::StandardButton standardButton(class QAbstractButton *)
    var arg0 = args[0].(*QAbstractButton).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK16QDialogButtonBox14standardButtonEP15QAbstractButton(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDialogButtonBox", "standardButton", args)
  }

  return
}

// <= body block end

