package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtWidgets/qscrollbar.h
// dst-file: /src/widgets/qscrollbar.go
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
  // proto:  const QMetaObject * QScrollBar::metaObject();
extern void C_ZNK10QScrollBar10metaObjectEv(void* qthis); // 4
  // proto:  QSize QScrollBar::sizeHint();
extern void C_ZNK10QScrollBar8sizeHintEv(void* qthis); // 4
  // proto:  void QScrollBar::QScrollBar(QWidget * parent);
extern void C_ZN10QScrollBarC2EP7QWidget(void* qthis, void* arg0); // 3
  // proto:  void QScrollBar::~QScrollBar();
extern void C_ZN10QScrollBarD2Ev(void* qthis); // 4
  // proto:  bool QScrollBar::event(QEvent * event);
extern void C_ZN10QScrollBar5eventEP6QEvent(void* qthis, void* arg0); // 4
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

// class sizeof(QScrollBar)=1
type QScrollBar struct {
  /*qbase*/ QAbstractSlider;
  qclsinst unsafe.Pointer /* *C.void */;
}

// metaObject()
func (this *QScrollBar) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QScrollBar10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK10QScrollBar10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScrollBar", "metaObject", args)
  }

}

// sizeHint()
func (this *QScrollBar) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QScrollBar8sizeHintEv
    // invoke: QSize sizeHint()
    C.C_ZNK10QScrollBar8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScrollBar", "sizeHint", args)
  }

}

// QScrollBar(class QWidget *)
func NewQScrollBar(args ...interface{}) QScrollBar {
  // QScrollBar(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QScrollBarC1EP7QWidget
    // invoke: void QScrollBar(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN10QScrollBarC2EP7QWidget(qthis, arg0)
  default:
    qtrt.ErrorResolve("QScrollBar", "QScrollBar", args)
  }

  return QScrollBar{}
}

// ~QScrollBar()
func (this *QScrollBar) FreeQScrollBar(args ...interface{}) () {
  // ~QScrollBar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QScrollBarD0Ev
    // invoke: void ~QScrollBar()
    C.C_ZN10QScrollBarD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScrollBar", "~QScrollBar", args)
  }

}

// event(class QEvent *)
func (this *QScrollBar) event(args ...interface{}) () {
  // event(class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QScrollBar5eventEP6QEvent
    // invoke: bool event(class QEvent *)
    var arg0 = args[0].(QEvent).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QScrollBar5eventEP6QEvent(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QScrollBar", "event", args)
  }

}

// <= body block end

