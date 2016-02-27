package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
  // proto:  bool QCheckBox::isTristate();
extern bool C_ZNK9QCheckBox10isTristateEv(void* qthis); // 4
  // proto:  void QCheckBox::~QCheckBox();
extern void C_ZN9QCheckBoxD2Ev(void* qthis); // 4
  // proto:  void QCheckBox::setTristate(bool y);
extern void C_ZN9QCheckBox11setTristateEb(void* qthis, bool arg0); // 4
  // proto:  void QCheckBox::QCheckBox(const QString & text, QWidget * parent);
extern void* C_ZN9QCheckBoxC2ERK7QStringP7QWidget(void* arg0, void* arg1); // 3
  // proto:  void QCheckBox::QCheckBox(QWidget * parent);
extern void* C_ZN9QCheckBoxC2EP7QWidget(void* arg0); // 3
  // proto:  QSize QCheckBox::sizeHint();
extern void* C_ZNK9QCheckBox8sizeHintEv(void* qthis); // 4
  // proto:  const QMetaObject * QCheckBox::metaObject();
extern void C_ZNK9QCheckBox10metaObjectEv(void* qthis); // 4
  // proto:  QSize QCheckBox::minimumSizeHint();
extern void* C_ZNK9QCheckBox15minimumSizeHintEv(void* qthis); // 4
  // proto:  Qt::CheckState QCheckBox::checkState();
extern void C_ZNK9QCheckBox10checkStateEv(void* qthis); // 4
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

// class sizeof(QCheckBox)=1
type QCheckBox struct {
  /*qbase*/ QAbstractButton;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _stateChanged QCheckBox_stateChanged_signal;
}

// isTristate()
func (this *QCheckBox) IsTristate(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QCheckBox10isTristateEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCheckBox", "isTristate", args)
  }

  return
}

// ~QCheckBox()
func (this *QCheckBox) Free(args ...interface{}) () {
  // ~QCheckBox()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QCheckBoxD0Ev
    // invoke: void ~QCheckBox()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN9QCheckBoxD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QCheckBox", "~QCheckBox", args)
  }

  return
}

// setTristate(_Bool)
func (this *QCheckBox) SetTristate(args ...interface{}) () {
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QCheckBox11setTristateEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCheckBox", "setTristate", args)
  }

  return
}

// QCheckBox(const class QString &, class QWidget *)
func GcfreeQCheckBox(this *QCheckBox) {
  qtrt.UniverseFree(this)
}
func NewQCheckBox(args ...interface{}) *QCheckBox {
  // QCheckBox(const class QString &, class QWidget *)
  // QCheckBox(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QCheckBoxC1ERK7QStringP7QWidget
    // invoke: void QCheckBox(const class QString &, class QWidget *)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QCheckBoxC2ERK7QStringP7QWidget(arg0, arg1)
    this := &QCheckBox{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQCheckBox)
    return this
  case 1:
    // invoke: _ZN9QCheckBoxC1EP7QWidget
    // invoke: void QCheckBox(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QCheckBoxC2EP7QWidget(arg0)
    this := &QCheckBox{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQCheckBox)
    return this
  default:
    qtrt.ErrorResolve("QCheckBox", "QCheckBox", args)
  }

  return nil // QCheckBox{Qclsinst:qthis}
}

// sizeHint()
func (this *QCheckBox) SizeHint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QCheckBox8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCheckBox", "sizeHint", args)
  }

  return
}

// metaObject()
func (this *QCheckBox) MetaObject(args ...interface{}) () {
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
    C.C_ZNK9QCheckBox10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCheckBox", "metaObject", args)
  }

  return
}

// minimumSizeHint()
func (this *QCheckBox) MinimumSizeHint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QCheckBox15minimumSizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCheckBox", "minimumSizeHint", args)
  }

  return
}

// checkState()
func (this *QCheckBox) CheckState(args ...interface{}) () {
  // checkState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QCheckBox10checkStateEv
    // invoke: Qt::CheckState checkState()
    C.C_ZNK9QCheckBox10checkStateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCheckBox", "checkState", args)
  }

  return
}

// <= body block end

