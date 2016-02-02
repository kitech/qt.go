package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtWidgets/qsizegrip.h
// dst-file: /src/widgets/qsizegrip.go
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
  // proto:  void QSizeGrip::QSizeGrip(QWidget * parent);
extern void* C_ZN9QSizeGripC2EP7QWidget(void* arg0); // 3
  // proto:  QSize QSizeGrip::sizeHint();
extern void* C_ZNK9QSizeGrip8sizeHintEv(void* qthis); // 4
  // proto:  const QMetaObject * QSizeGrip::metaObject();
extern void C_ZNK9QSizeGrip10metaObjectEv(void* qthis); // 4
  // proto:  void QSizeGrip::setVisible(bool );
extern void C_ZN9QSizeGrip10setVisibleEb(void* qthis, bool arg0); // 4
  // proto:  void QSizeGrip::~QSizeGrip();
extern void C_ZN9QSizeGripD2Ev(void* qthis); // 4
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

// class sizeof(QSizeGrip)=1
type QSizeGrip struct {
  /*qbase*/ QWidget;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// QSizeGrip(class QWidget *)
func NewQSizeGrip(args ...interface{}) *QSizeGrip {
  // QSizeGrip(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSizeGripC1EP7QWidget
    // invoke: void QSizeGrip(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QSizeGripC2EP7QWidget(arg0)
    return &QSizeGrip{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QSizeGrip", "QSizeGrip", args)
  }

  return nil // QSizeGrip{Qclsinst:qthis}
}

// sizeHint()
func (this *QSizeGrip) Sizehint(args ...interface{}) (ret interface{}) {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSizeGrip8sizeHintEv
    // invoke: QSize sizeHint()
    var ret0 = C.C_ZNK9QSizeGrip8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSizeGrip", "sizeHint", args)
  }

  return
}

// metaObject()
func (this *QSizeGrip) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSizeGrip10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK9QSizeGrip10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSizeGrip", "metaObject", args)
  }

  return
}

// setVisible(_Bool)
func (this *QSizeGrip) Setvisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSizeGrip10setVisibleEb
    // invoke: void setVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QSizeGrip10setVisibleEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSizeGrip", "setVisible", args)
  }

  return
}

// ~QSizeGrip()
func (this *QSizeGrip) Freeqsizegrip(args ...interface{}) () {
  // ~QSizeGrip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSizeGripD0Ev
    // invoke: void ~QSizeGrip()
    C.C_ZN9QSizeGripD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSizeGrip", "~QSizeGrip", args)
  }

  return
}

// <= body block end

