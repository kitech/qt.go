package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtWidgets/qcheckbox.h
// dst-file: /src/widgets/qcheckbox.go
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
  // proto:  const QMetaObject * QCheckBox::metaObject();
extern void _ZNK9QCheckBox10metaObjectEv(void* qthis);
  // proto:  QSize QCheckBox::minimumSizeHint();
extern void _ZNK9QCheckBox15minimumSizeHintEv(void* qthis);
  // proto:  void QCheckBox::~QCheckBox();
extern void _ZN9QCheckBoxD0Ev(void* qthis);
  // proto:  QSize QCheckBox::sizeHint();
extern void _ZNK9QCheckBox8sizeHintEv(void* qthis);
  // proto:  void QCheckBox::setTristate(bool y);
extern void _ZN9QCheckBox11setTristateEb(void* qthis, bool arg0);
  // proto:  void QCheckBox::QCheckBox(const QCheckBox & );
extern void* dector_ZN9QCheckBoxC1ERKS_(void* arg0);
extern void _ZN9QCheckBoxC1ERKS_(void* qthis, void* arg0);
  // proto:  void QCheckBox::QCheckBox(QWidget * parent);
extern void* dector_ZN9QCheckBoxC1EP7QWidget(void* arg0);
extern void _ZN9QCheckBoxC1EP7QWidget(void* qthis, void* arg0);
  // proto:  bool QCheckBox::isTristate();
extern void _ZNK9QCheckBox10isTristateEv(void* qthis);
  // proto:  void QCheckBox::QCheckBox(const QString & text, QWidget * parent);
extern void* dector_ZN9QCheckBoxC1ERK7QStringP7QWidget(void* arg0, void* arg1);
extern void _ZN9QCheckBoxC1ERK7QStringP7QWidget(void* qthis, void* arg0, void* arg1);
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

// class sizeof(QCheckBox)=1
type QCheckBox struct {
  /*qbase*/ QAbstractButton;
  qclsinst uint64 /* *mut c_void*/;
//  _stateChanged QCheckBox_stateChanged_signal;
}

  // proto:  const QMetaObject * QCheckBox::metaObject();
func (this *QCheckBox) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QCheckBox10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK9QCheckBox10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCheckBox", "metaObject", args)
  }

}

  // proto:  QSize QCheckBox::minimumSizeHint();
func (this *QCheckBox) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QCheckBox15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    C._ZNK9QCheckBox15minimumSizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCheckBox", "minimumSizeHint", args)
  }

}

  // proto:  void QCheckBox::~QCheckBox();
func (this *QCheckBox) FreeQCheckBox(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCheckBox", "~QCheckBox", args)
  }

}

  // proto:  QSize QCheckBox::sizeHint();
func (this *QCheckBox) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QCheckBox8sizeHintEv
    // invoke: QSize sizeHint()
    C._ZNK9QCheckBox8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCheckBox", "sizeHint", args)
  }

}

  // proto:  void QCheckBox::setTristate(bool y);
func (this *QCheckBox) setTristate(args ...interface{}) () {
  // setTristate(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QCheckBox11setTristateEb
    // invoke: void setTristate(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN9QCheckBox11setTristateEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCheckBox", "setTristate", args)
  }

}

  // proto:  void QCheckBox::QCheckBox(const QCheckBox & );
func NewQCheckBox(args ...interface{}) QCheckBox {
  return QCheckBox{}
}

  // proto:  bool QCheckBox::isTristate();
func (this *QCheckBox) isTristate(args ...interface{}) () {
  // isTristate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QCheckBox10isTristateEv
    // invoke: bool isTristate()
    C._ZNK9QCheckBox10isTristateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCheckBox", "isTristate", args)
  }

}

// <= body block end

