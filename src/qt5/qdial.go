package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtWidgets/qdial.h
// dst-file: /src/widgets/qdial.go
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
  // proto:  void QDial::setNotchTarget(double target);
extern void C_ZN5QDial14setNotchTargetEd(void* qthis, double arg0); // 4
  // proto:  bool QDial::notchesVisible();
extern bool C_ZNK5QDial14notchesVisibleEv(void* qthis); // 4
  // proto:  void QDial::~QDial();
extern void C_ZN5QDialD2Ev(void* qthis); // 4
  // proto:  void QDial::QDial(QWidget * parent);
extern void* C_ZN5QDialC2EP7QWidget(void* arg0); // 3
  // proto:  void QDial::setWrapping(bool on);
extern void C_ZN5QDial11setWrappingEb(void* qthis, bool arg0); // 4
  // proto:  void QDial::setNotchesVisible(bool visible);
extern void C_ZN5QDial17setNotchesVisibleEb(void* qthis, bool arg0); // 4
  // proto:  int QDial::notchSize();
extern int32_t C_ZNK5QDial9notchSizeEv(void* qthis); // 4
  // proto:  bool QDial::wrapping();
extern bool C_ZNK5QDial8wrappingEv(void* qthis); // 4
  // proto:  QSize QDial::sizeHint();
extern void* C_ZNK5QDial8sizeHintEv(void* qthis); // 4
  // proto:  qreal QDial::notchTarget();
extern double C_ZNK5QDial11notchTargetEv(void* qthis); // 4
  // proto:  const QMetaObject * QDial::metaObject();
extern void C_ZNK5QDial10metaObjectEv(void* qthis); // 4
  // proto:  QSize QDial::minimumSizeHint();
extern void* C_ZNK5QDial15minimumSizeHintEv(void* qthis); // 4
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

// class sizeof(QDial)=1
type QDial struct {
  /*qbase*/ QAbstractSlider;
  qclsinst unsafe.Pointer /* *C.void */;
}

// setNotchTarget(double)
func (this *QDial) Setnotchtarget(args ...interface{}) () {
  // setNotchTarget(double)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "double"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QDial14setNotchTargetEd
    // invoke: void setNotchTarget(double)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN5QDial14setNotchTargetEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDial", "setNotchTarget", args)
  }

  return
}

// notchesVisible()
func (this *QDial) Notchesvisible(args ...interface{}) (ret interface{}) {
  // notchesVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDial14notchesVisibleEv
    // invoke: bool notchesVisible()
    var ret0 = C.C_ZNK5QDial14notchesVisibleEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDial", "notchesVisible", args)
  }

  return
}

// ~QDial()
func (this *QDial) Freeqdial(args ...interface{}) () {
  // ~QDial()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QDialD0Ev
    // invoke: void ~QDial()
    C.C_ZN5QDialD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDial", "~QDial", args)
  }

  return
}

// QDial(class QWidget *)
func NewQDial(args ...interface{}) *QDial {
  // QDial(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QDialC1EP7QWidget
    // invoke: void QDial(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QDialC2EP7QWidget(arg0)
    return &QDial{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QDial", "QDial", args)
  }

  return nil // QDial{qclsinst:qthis}
}

// setWrapping(_Bool)
func (this *QDial) Setwrapping(args ...interface{}) () {
  // setWrapping(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QDial11setWrappingEb
    // invoke: void setWrapping(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN5QDial11setWrappingEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDial", "setWrapping", args)
  }

  return
}

// setNotchesVisible(_Bool)
func (this *QDial) Setnotchesvisible(args ...interface{}) () {
  // setNotchesVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QDial17setNotchesVisibleEb
    // invoke: void setNotchesVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN5QDial17setNotchesVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDial", "setNotchesVisible", args)
  }

  return
}

// notchSize()
func (this *QDial) Notchsize(args ...interface{}) (ret interface{}) {
  // notchSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDial9notchSizeEv
    // invoke: int notchSize()
    var ret0 = C.C_ZNK5QDial9notchSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDial", "notchSize", args)
  }

  return
}

// wrapping()
func (this *QDial) Wrapping(args ...interface{}) (ret interface{}) {
  // wrapping()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDial8wrappingEv
    // invoke: bool wrapping()
    var ret0 = C.C_ZNK5QDial8wrappingEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDial", "wrapping", args)
  }

  return
}

// sizeHint()
func (this *QDial) Sizehint(args ...interface{}) (ret interface{}) {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDial8sizeHintEv
    // invoke: QSize sizeHint()
    var ret0 = C.C_ZNK5QDial8sizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDial", "sizeHint", args)
  }

  return
}

// notchTarget()
func (this *QDial) Notchtarget(args ...interface{}) (ret interface{}) {
  // notchTarget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDial11notchTargetEv
    // invoke: qreal notchTarget()
    var ret0 = C.C_ZNK5QDial11notchTargetEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDial", "notchTarget", args)
  }

  return
}

// metaObject()
func (this *QDial) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDial10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK5QDial10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDial", "metaObject", args)
  }

  return
}

// minimumSizeHint()
func (this *QDial) Minimumsizehint(args ...interface{}) (ret interface{}) {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDial15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    var ret0 = C.C_ZNK5QDial15minimumSizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDial", "minimumSizeHint", args)
  }

  return
}

// <= body block end

