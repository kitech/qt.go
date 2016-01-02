package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
  // proto:  QList<QAbstractButton *> QDialogButtonBox::buttons();
extern void _ZNK16QDialogButtonBox7buttonsEv(void* qthis);
  // proto:  void QDialogButtonBox::setCenterButtons(bool center);
extern void _ZN16QDialogButtonBox16setCenterButtonsEb(void* qthis, bool arg0);
  // proto:  bool QDialogButtonBox::centerButtons();
extern void _ZNK16QDialogButtonBox13centerButtonsEv(void* qthis);
  // proto:  const QMetaObject * QDialogButtonBox::metaObject();
extern void _ZNK16QDialogButtonBox10metaObjectEv(void* qthis);
  // proto:  void QDialogButtonBox::removeButton(QAbstractButton * button);
extern void _ZN16QDialogButtonBox12removeButtonEP15QAbstractButton(void* qthis, void* arg0);
  // proto:  void QDialogButtonBox::~QDialogButtonBox();
extern void _ZN16QDialogButtonBoxD0Ev(void* qthis);
  // proto:  void QDialogButtonBox::QDialogButtonBox(QWidget * parent);
extern void* dector_ZN16QDialogButtonBoxC1EP7QWidget(void* arg0);
extern void _ZN16QDialogButtonBoxC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QDialogButtonBox::QDialogButtonBox(const QDialogButtonBox & );
extern void* dector_ZN16QDialogButtonBoxC1ERKS_(void* arg0);
extern void _ZN16QDialogButtonBoxC1ERKS_(void* qthis, void* arg0);
  // proto:  void QDialogButtonBox::clear();
extern void _ZN16QDialogButtonBox5clearEv(void* qthis);
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

// class sizeof(QDialogButtonBox)=1
type QDialogButtonBox struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
//  _helpRequested QDialogButtonBox_helpRequested_signal;
//  _accepted QDialogButtonBox_accepted_signal;
//  _clicked QDialogButtonBox_clicked_signal;
//  _rejected QDialogButtonBox_rejected_signal;
}

  // proto:  QList<QAbstractButton *> QDialogButtonBox::buttons();
func (this *QDialogButtonBox) buttons(args ...interface{}) () {
  // buttons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QDialogButtonBox7buttonsEv
  default:
    qtrt.ErrorResolve("QDialogButtonBox", "buttons", args)
  }

}

  // proto:  void QDialogButtonBox::setCenterButtons(bool center);
func (this *QDialogButtonBox) setCenterButtons(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QDialogButtonBox", "setCenterButtons", args)
  }

}

  // proto:  bool QDialogButtonBox::centerButtons();
func (this *QDialogButtonBox) centerButtons(args ...interface{}) () {
  // centerButtons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QDialogButtonBox13centerButtonsEv
  default:
    qtrt.ErrorResolve("QDialogButtonBox", "centerButtons", args)
  }

}

  // proto:  const QMetaObject * QDialogButtonBox::metaObject();
func (this *QDialogButtonBox) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QDialogButtonBox10metaObjectEv
  default:
    qtrt.ErrorResolve("QDialogButtonBox", "metaObject", args)
  }

}

  // proto:  void QDialogButtonBox::removeButton(QAbstractButton * button);
func (this *QDialogButtonBox) removeButton(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QDialogButtonBox", "removeButton", args)
  }

}

  // proto:  void QDialogButtonBox::~QDialogButtonBox();
func (this *QDialogButtonBox) FreeQDialogButtonBox(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDialogButtonBox", "~QDialogButtonBox", args)
  }

}

  // proto:  void QDialogButtonBox::QDialogButtonBox(QWidget * parent);
func NewQDialogButtonBox(args ...interface{}) QDialogButtonBox {
  return QDialogButtonBox{}
}

  // proto:  void QDialogButtonBox::clear();
func (this *QDialogButtonBox) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QDialogButtonBox5clearEv
  default:
    qtrt.ErrorResolve("QDialogButtonBox", "clear", args)
  }

}

// <= body block end

