package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtWidgets/qradiobutton.h
// dst-file: /src/widgets/qradiobutton.go
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
  // proto:  const QMetaObject * QRadioButton::metaObject();
extern void C_ZNK12QRadioButton10metaObjectEv(void* qthis); // 4
  // proto:  QSize QRadioButton::minimumSizeHint();
extern void* C_ZNK12QRadioButton15minimumSizeHintEv(void* qthis); // 4
  // proto:  QSize QRadioButton::sizeHint();
extern void* C_ZNK12QRadioButton8sizeHintEv(void* qthis); // 4
  // proto:  void QRadioButton::~QRadioButton();
extern void C_ZN12QRadioButtonD2Ev(void* qthis); // 4
  // proto:  void QRadioButton::QRadioButton(QWidget * parent);
extern void* C_ZN12QRadioButtonC2EP7QWidget(void* arg0); // 3
  // proto:  void QRadioButton::QRadioButton(const QString & text, QWidget * parent);
extern void* C_ZN12QRadioButtonC2ERK7QStringP7QWidget(void* arg0, void* arg1); // 3
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

// class sizeof(QRadioButton)=1
type QRadioButton struct {
  /*qbase*/ QAbstractButton;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// metaObject()
func (this *QRadioButton) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QRadioButton10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK12QRadioButton10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QRadioButton", "metaObject", args)
  }

  return
}

// minimumSizeHint()
func (this *QRadioButton) Minimumsizehint(args ...interface{}) (ret interface{}) {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QRadioButton15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    var ret0 = C.C_ZNK12QRadioButton15minimumSizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRadioButton", "minimumSizeHint", args)
  }

  return
}

// sizeHint()
func (this *QRadioButton) Sizehint(args ...interface{}) (ret interface{}) {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QRadioButton8sizeHintEv
    // invoke: QSize sizeHint()
    var ret0 = C.C_ZNK12QRadioButton8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRadioButton", "sizeHint", args)
  }

  return
}

// ~QRadioButton()
func (this *QRadioButton) Freeqradiobutton(args ...interface{}) () {
  // ~QRadioButton()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QRadioButtonD0Ev
    // invoke: void ~QRadioButton()
    C.C_ZN12QRadioButtonD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QRadioButton", "~QRadioButton", args)
  }

  return
}

// QRadioButton(class QWidget *)
func NewQRadioButton(args ...interface{}) *QRadioButton {
  // QRadioButton(class QWidget *)
  // QRadioButton(const class QString &, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QRadioButtonC1EP7QWidget
    // invoke: void QRadioButton(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QRadioButtonC2EP7QWidget(arg0)
    return &QRadioButton{Qclsinst:qthis}
  case 1:
    // invoke: _ZN12QRadioButtonC1ERK7QStringP7QWidget
    // invoke: void QRadioButton(const class QString &, class QWidget *)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QRadioButtonC2ERK7QStringP7QWidget(arg0, arg1)
    return &QRadioButton{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QRadioButton", "QRadioButton", args)
  }

  return nil // QRadioButton{Qclsinst:qthis}
}

// <= body block end

