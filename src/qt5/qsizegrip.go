package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:13 2016
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
extern void _ZN9QSizeGripC2EP7QWidget(void* qthis, void* arg0); // 3
  // proto:  QSize QSizeGrip::sizeHint();
extern void _ZNK9QSizeGrip8sizeHintEv(void* qthis); // 4
  // proto:  const QMetaObject * QSizeGrip::metaObject();
extern void _ZNK9QSizeGrip10metaObjectEv(void* qthis); // 4
  // proto:  void QSizeGrip::setVisible(bool );
extern void _ZN9QSizeGrip10setVisibleEb(void* qthis, bool arg0); // 4
  // proto:  void QSizeGrip::~QSizeGrip();
extern void _ZN9QSizeGripD2Ev(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// QSizeGrip(class QWidget *)
func NewQSizeGrip(args ...interface{}) QSizeGrip {
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
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN9QSizeGripC2EP7QWidget(qthis, arg0)
  default:
    qtrt.ErrorResolve("QSizeGrip", "QSizeGrip", args)
  }

  return QSizeGrip{}
}

// sizeHint()
func (this *QSizeGrip) sizeHint(args ...interface{}) () {
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
    C._ZNK9QSizeGrip8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSizeGrip", "sizeHint", args)
  }

}

// metaObject()
func (this *QSizeGrip) metaObject(args ...interface{}) () {
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
    C._ZNK9QSizeGrip10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSizeGrip", "metaObject", args)
  }

}

// setVisible(_Bool)
func (this *QSizeGrip) setVisible(args ...interface{}) () {
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
    C._ZN9QSizeGrip10setVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSizeGrip", "setVisible", args)
  }

}

// ~QSizeGrip()
func (this *QSizeGrip) FreeQSizeGrip(args ...interface{}) () {
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
    C._ZN9QSizeGripD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSizeGrip", "~QSizeGrip", args)
  }

}

// <= body block end

