package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
  // proto:  const QMetaObject * QScrollBar::metaObject();
extern void C_ZNK10QScrollBar10metaObjectEv(void* qthis); // 4
  // proto:  QSize QScrollBar::sizeHint();
extern void* C_ZNK10QScrollBar8sizeHintEv(void* qthis); // 4
  // proto:  void QScrollBar::QScrollBar(QWidget * parent);
extern void* C_ZN10QScrollBarC2EP7QWidget(void* arg0); // 3
  // proto:  void QScrollBar::~QScrollBar();
extern void C_ZN10QScrollBarD2Ev(void* qthis); // 4
  // proto:  bool QScrollBar::event(QEvent * event);
extern bool C_ZN10QScrollBar5eventEP6QEvent(void* qthis, void* arg0); // 4
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
}

// class sizeof(QScrollBar)=1
type QScrollBar struct {
  /*qbase*/ QAbstractSlider;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// metaObject()
func (this *QScrollBar) Metaobject(args ...interface{}) () {
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
    C.C_ZNK10QScrollBar10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QScrollBar", "metaObject", args)
  }

  return
}

// sizeHint()
func (this *QScrollBar) Sizehint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QScrollBar8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScrollBar", "sizeHint", args)
  }

  return
}

// QScrollBar(class QWidget *)
func NewQScrollBar(args ...interface{}) *QScrollBar {
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QScrollBarC2EP7QWidget(arg0)
    return &QScrollBar{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QScrollBar", "QScrollBar", args)
  }

  return nil // QScrollBar{Qclsinst:qthis}
}

// ~QScrollBar()
func (this *QScrollBar) Freeqscrollbar(args ...interface{}) () {
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
    C.C_ZN10QScrollBarD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QScrollBar", "~QScrollBar", args)
  }

  return
}

// event(class QEvent *)
func (this *QScrollBar) Event(args ...interface{}) (ret interface{}) {
  // event(class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QScrollBar5eventEP6QEvent
    // invoke: bool event(class QEvent *)
    var arg0 = args[0].(*qtcore.QEvent).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN10QScrollBar5eventEP6QEvent(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScrollBar", "event", args)
  }

  return
}

// <= body block end

