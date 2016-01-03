package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
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
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  const QMetaObject * QSlider::metaObject();
extern void _ZNK7QSlider10metaObjectEv(void* qthis);
  // proto:  bool QSlider::event(QEvent * event);
extern void _ZN7QSlider5eventEP6QEvent(void* qthis, void* arg0);
  // proto:  int QSlider::tickInterval();
extern void _ZNK7QSlider12tickIntervalEv(void* qthis);
  // proto:  QSize QSlider::sizeHint();
extern void _ZNK7QSlider8sizeHintEv(void* qthis);
  // proto:  void QSlider::setTickInterval(int ti);
extern void _ZN7QSlider15setTickIntervalEi(void* qthis, int32_t arg0);
  // proto:  void QSlider::QSlider(const QSlider & );
extern void* dector_ZN7QSliderC1ERKS_(void* arg0);
extern void _ZN7QSliderC1ERKS_(void* qthis, void* arg0);
  // proto:  void QSlider::~QSlider();
extern void _ZN7QSliderD0Ev(void* qthis);
  // proto:  void QSlider::QSlider(QWidget * parent);
extern void* dector_ZN7QSliderC1EP7QWidget(void* arg0);
extern void _ZN7QSliderC1EP7QWidget(void* qthis, void* arg0);
  // proto:  QSize QSlider::minimumSizeHint();
extern void _ZNK7QSlider15minimumSizeHintEv(void* qthis);
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

  // proto:  const QMetaObject * QSlider::metaObject();
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
    C._ZNK7QSlider10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSlider", "metaObject", args)
  }

}

  // proto:  bool QSlider::event(QEvent * event);
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
    C._ZN7QSlider5eventEP6QEvent(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSlider", "event", args)
  }

}

  // proto:  int QSlider::tickInterval();
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
    C._ZNK7QSlider12tickIntervalEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSlider", "tickInterval", args)
  }

}

  // proto:  QSize QSlider::sizeHint();
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
    C._ZNK7QSlider8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSlider", "sizeHint", args)
  }

}

  // proto:  void QSlider::setTickInterval(int ti);
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
    C._ZN7QSlider15setTickIntervalEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSlider", "setTickInterval", args)
  }

}

  // proto:  void QSlider::QSlider(const QSlider & );
func NewQSlider(args ...interface{}) QSlider {
  return QSlider{}
}

  // proto:  void QSlider::~QSlider();
func (this *QSlider) FreeQSlider(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSlider", "~QSlider", args)
  }

}

  // proto:  QSize QSlider::minimumSizeHint();
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
    C._ZNK7QSlider15minimumSizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSlider", "minimumSizeHint", args)
  }

}

// <= body block end

