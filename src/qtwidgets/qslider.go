package qtwidgets
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
// src-file: /QtWidgets/qslider.h
// dst-file: /src/widgets/qslider.go
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
  // proto:  bool QSlider::event(QEvent * event);
extern bool C_ZN7QSlider5eventEP6QEvent(void* qthis, void* arg0); // 4
  // proto:  void QSlider::~QSlider();
extern void C_ZN7QSliderD2Ev(void* qthis); // 4
  // proto:  QSize QSlider::sizeHint();
extern void* C_ZNK7QSlider8sizeHintEv(void* qthis); // 4
  // proto:  int QSlider::tickInterval();
extern int32_t C_ZNK7QSlider12tickIntervalEv(void* qthis); // 4
  // proto:  const QMetaObject * QSlider::metaObject();
extern void C_ZNK7QSlider10metaObjectEv(void* qthis); // 4
  // proto:  QSize QSlider::minimumSizeHint();
extern void* C_ZNK7QSlider15minimumSizeHintEv(void* qthis); // 4
  // proto:  void QSlider::QSlider(QWidget * parent);
extern void* C_ZN7QSliderC2EP7QWidget(void* arg0); // 3
  // proto:  QSlider::TickPosition QSlider::tickPosition();
extern void C_ZNK7QSlider12tickPositionEv(void* qthis); // 4
  // proto:  void QSlider::setTickInterval(int ti);
extern void C_ZN7QSlider15setTickIntervalEi(void* qthis, int32_t arg0); // 4
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

// class sizeof(QSlider)=1
type QSlider struct {
  /*qbase*/ QAbstractSlider;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// event(class QEvent *)
func (this *QSlider) Event(args ...interface{}) (ret interface{}) {
  // event(class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QEvent{}) // "QEvent *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QSlider5eventEP6QEvent
    // invoke: bool event(class QEvent *)
    var arg0 = args[0].(*qtcore.QEvent).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN7QSlider5eventEP6QEvent(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSlider", "event", args)
  }

  return
}

// ~QSlider()
func (this *QSlider) Freeqslider(args ...interface{}) () {
  // ~QSlider()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QSliderD0Ev
    // invoke: void ~QSlider()
    C.C_ZN7QSliderD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSlider", "~QSlider", args)
  }

  return
}

// sizeHint()
func (this *QSlider) Sizehint(args ...interface{}) (ret interface{}) {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QSlider8sizeHintEv
    // invoke: QSize sizeHint()
    var ret0 = C.C_ZNK7QSlider8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSlider", "sizeHint", args)
  }

  return
}

// tickInterval()
func (this *QSlider) Tickinterval(args ...interface{}) (ret interface{}) {
  // tickInterval()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QSlider12tickIntervalEv
    // invoke: int tickInterval()
    var ret0 = C.C_ZNK7QSlider12tickIntervalEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSlider", "tickInterval", args)
  }

  return
}

// metaObject()
func (this *QSlider) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QSlider10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK7QSlider10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSlider", "metaObject", args)
  }

  return
}

// minimumSizeHint()
func (this *QSlider) Minimumsizehint(args ...interface{}) (ret interface{}) {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QSlider15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    var ret0 = C.C_ZNK7QSlider15minimumSizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSlider", "minimumSizeHint", args)
  }

  return
}

// QSlider(class QWidget *)
func NewQSlider(args ...interface{}) *QSlider {
  // QSlider(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QSliderC1EP7QWidget
    // invoke: void QSlider(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QSliderC2EP7QWidget(arg0)
    return &QSlider{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QSlider", "QSlider", args)
  }

  return nil // QSlider{Qclsinst:qthis}
}

// tickPosition()
func (this *QSlider) Tickposition(args ...interface{}) () {
  // tickPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QSlider12tickPositionEv
    // invoke: QSlider::TickPosition tickPosition()
    C.C_ZNK7QSlider12tickPositionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSlider", "tickPosition", args)
  }

  return
}

// setTickInterval(int)
func (this *QSlider) Settickinterval(args ...interface{}) () {
  // setTickInterval(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QSlider15setTickIntervalEi
    // invoke: void setTickInterval(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QSlider15setTickIntervalEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSlider", "setTickInterval", args)
  }

  return
}

// <= body block end

