package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  bool QSlider::event(QEvent * event);
extern void C_ZN7QSlider5eventEP6QEvent(void* qthis, void* arg0); // 4
  // proto:  void QSlider::~QSlider();
extern void C_ZN7QSliderD2Ev(void* qthis); // 4
  // proto:  QSize QSlider::sizeHint();
extern void C_ZNK7QSlider8sizeHintEv(void* qthis); // 4
  // proto:  int QSlider::tickInterval();
extern void C_ZNK7QSlider12tickIntervalEv(void* qthis); // 4
  // proto:  const QMetaObject * QSlider::metaObject();
extern void C_ZNK7QSlider10metaObjectEv(void* qthis); // 4
  // proto:  QSize QSlider::minimumSizeHint();
extern void C_ZNK7QSlider15minimumSizeHintEv(void* qthis); // 4
  // proto:  void QSlider::QSlider(QWidget * parent);
extern void C_ZN7QSliderC2EP7QWidget(void* qthis, void* arg0); // 3
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
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QSlider)=1
type QSlider struct {
  /*qbase*/ QAbstractSlider;
  qclsinst unsafe.Pointer /* *C.void */;
}

// event(class QEvent *)
func (this *QSlider) event(args ...interface{}) () {
  // event(class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QSlider5eventEP6QEvent
    // invoke: bool event(class QEvent *)
    var arg0 = args[0].(QEvent).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QSlider5eventEP6QEvent(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSlider", "event", args)
  }

}

// ~QSlider()
func (this *QSlider) FreeQSlider(args ...interface{}) () {
  // ~QSlider()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QSliderD0Ev
    // invoke: void ~QSlider()
    C.C_ZN7QSliderD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSlider", "~QSlider", args)
  }

}

// sizeHint()
func (this *QSlider) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QSlider8sizeHintEv
    // invoke: QSize sizeHint()
    C.C_ZNK7QSlider8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSlider", "sizeHint", args)
  }

}

// tickInterval()
func (this *QSlider) tickInterval(args ...interface{}) () {
  // tickInterval()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QSlider12tickIntervalEv
    // invoke: int tickInterval()
    C.C_ZNK7QSlider12tickIntervalEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSlider", "tickInterval", args)
  }

}

// metaObject()
func (this *QSlider) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QSlider10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK7QSlider10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSlider", "metaObject", args)
  }

}

// minimumSizeHint()
func (this *QSlider) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QSlider15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    C.C_ZNK7QSlider15minimumSizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSlider", "minimumSizeHint", args)
  }

}

// QSlider(class QWidget *)
func NewQSlider(args ...interface{}) QSlider {
  // QSlider(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QSliderC1EP7QWidget
    // invoke: void QSlider(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN7QSliderC2EP7QWidget(qthis, arg0)
  default:
    qtrt.ErrorResolve("QSlider", "QSlider", args)
  }

  return QSlider{}
}

// tickPosition()
func (this *QSlider) tickPosition(args ...interface{}) () {
  // tickPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QSlider12tickPositionEv
    // invoke: QSlider::TickPosition tickPosition()
    C.C_ZNK7QSlider12tickPositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSlider", "tickPosition", args)
  }

}

// setTickInterval(int)
func (this *QSlider) setTickInterval(args ...interface{}) () {
  // setTickInterval(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QSlider15setTickIntervalEi
    // invoke: void setTickInterval(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QSlider15setTickIntervalEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSlider", "setTickInterval", args)
  }

}

// <= body block end

