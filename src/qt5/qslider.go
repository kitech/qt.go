package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
extern void _ZN7QSlider15setTickIntervalEi(void* qthis, int arg0);
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
  qclsinst uint64 /* *mut c_void*/;
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
  default:
    qtrt.ErrorResolve("QSlider", "minimumSizeHint", args)
  }

}

// <= body block end

