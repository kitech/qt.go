package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:13 2016
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
extern void _ZNK12QRadioButton10metaObjectEv(void* qthis); // 4
  // proto:  QSize QRadioButton::minimumSizeHint();
extern void _ZNK12QRadioButton15minimumSizeHintEv(void* qthis); // 4
  // proto:  QSize QRadioButton::sizeHint();
extern void _ZNK12QRadioButton8sizeHintEv(void* qthis); // 4
  // proto:  void QRadioButton::~QRadioButton();
extern void _ZN12QRadioButtonD2Ev(void* qthis); // 4
  // proto:  void QRadioButton::QRadioButton(QWidget * parent);
extern void _ZN12QRadioButtonC2EP7QWidget(void* qthis, void* arg0); // 3
  // proto:  void QRadioButton::QRadioButton(const QString & text, QWidget * parent);
extern void _ZN12QRadioButtonC2ERK7QStringP7QWidget(void* qthis, void* arg0, void* arg1); // 3
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// metaObject()
func (this *QRadioButton) metaObject(args ...interface{}) () {
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
    C._ZNK12QRadioButton10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRadioButton", "metaObject", args)
  }

}

// minimumSizeHint()
func (this *QRadioButton) minimumSizeHint(args ...interface{}) () {
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
    C._ZNK12QRadioButton15minimumSizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRadioButton", "minimumSizeHint", args)
  }

}

// sizeHint()
func (this *QRadioButton) sizeHint(args ...interface{}) () {
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
    C._ZNK12QRadioButton8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRadioButton", "sizeHint", args)
  }

}

// ~QRadioButton()
func (this *QRadioButton) FreeQRadioButton(args ...interface{}) () {
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
    C._ZN12QRadioButtonD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRadioButton", "~QRadioButton", args)
  }

}

// QRadioButton(class QWidget *)
func NewQRadioButton(args ...interface{}) QRadioButton {
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
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN12QRadioButtonC2EP7QWidget(qthis, arg0)
  case 1:
    // invoke: _ZN12QRadioButtonC1ERK7QStringP7QWidget
    // invoke: void QRadioButton(const class QString &, class QWidget *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN12QRadioButtonC2ERK7QStringP7QWidget(qthis, arg0, arg1)
  default:
    qtrt.ErrorResolve("QRadioButton", "QRadioButton", args)
  }

  return QRadioButton{}
}

// <= body block end

