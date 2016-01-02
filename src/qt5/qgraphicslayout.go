package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtWidgets/qgraphicslayout.h
// dst-file: /src/widgets/qgraphicslayout.go
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
  // proto:  void QGraphicsLayout::updateGeometry();
extern void _ZN15QGraphicsLayout14updateGeometryEv(void* qthis);
  // proto:  bool QGraphicsLayout::isActivated();
extern void _ZNK15QGraphicsLayout11isActivatedEv(void* qthis);
  // proto:  void QGraphicsLayout::invalidate();
extern void _ZN15QGraphicsLayout10invalidateEv(void* qthis);
  // proto:  void QGraphicsLayout::removeAt(int index);
extern void _ZN15QGraphicsLayout8removeAtEi(void* qthis, int arg0);
  // proto:  QGraphicsLayoutItem * QGraphicsLayout::itemAt(int i);
extern void _ZNK15QGraphicsLayout6itemAtEi(void* qthis, int arg0);
  // proto:  void QGraphicsLayout::QGraphicsLayout(const QGraphicsLayout & );
extern void* dector_ZN15QGraphicsLayoutC1ERKS_(void* arg0);
extern void _ZN15QGraphicsLayoutC1ERKS_(void* qthis, void* arg0);
  // proto:  void QGraphicsLayout::getContentsMargins(qreal * left, qreal * top, qreal * right, qreal * bottom);
extern void _ZNK15QGraphicsLayout18getContentsMarginsEPdS0_S0_S0_(void* qthis, double* arg0, double* arg1, double* arg2, double* arg3);
  // proto:  void QGraphicsLayout::setContentsMargins(qreal left, qreal top, qreal right, qreal bottom);
extern void _ZN15QGraphicsLayout18setContentsMarginsEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  void QGraphicsLayout::widgetEvent(QEvent * e);
extern void _ZN15QGraphicsLayout11widgetEventEP6QEvent(void* qthis, void* arg0);
  // proto: static bool QGraphicsLayout::instantInvalidatePropagation();
extern void _ZN15QGraphicsLayout28instantInvalidatePropagationEv();
  // proto: static void QGraphicsLayout::setInstantInvalidatePropagation(bool enable);
extern void _ZN15QGraphicsLayout31setInstantInvalidatePropagationEb(bool arg0);
  // proto:  void QGraphicsLayout::~QGraphicsLayout();
extern void _ZN15QGraphicsLayoutD0Ev(void* qthis);
  // proto:  void QGraphicsLayout::activate();
extern void _ZN15QGraphicsLayout8activateEv(void* qthis);
  // proto:  int QGraphicsLayout::count();
extern void _ZNK15QGraphicsLayout5countEv(void* qthis);
  // proto:  void QGraphicsLayout::QGraphicsLayout(QGraphicsLayoutItem * parent);
extern void* dector_ZN15QGraphicsLayoutC1EP19QGraphicsLayoutItem(void* arg0);
extern void _ZN15QGraphicsLayoutC1EP19QGraphicsLayoutItem(void* qthis, void* arg0);
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

// class sizeof(QGraphicsLayout)=1
type QGraphicsLayout struct {
  /*qbase*/ QGraphicsLayoutItem;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QGraphicsLayout::updateGeometry();
func (this *QGraphicsLayout) updateGeometry(args ...interface{}) () {
  // updateGeometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsLayout14updateGeometryEv
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "updateGeometry", args)
  }

}

  // proto:  bool QGraphicsLayout::isActivated();
func (this *QGraphicsLayout) isActivated(args ...interface{}) () {
  // isActivated()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsLayout11isActivatedEv
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "isActivated", args)
  }

}

  // proto:  void QGraphicsLayout::invalidate();
func (this *QGraphicsLayout) invalidate(args ...interface{}) () {
  // invalidate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsLayout10invalidateEv
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "invalidate", args)
  }

}

  // proto:  void QGraphicsLayout::removeAt(int index);
func (this *QGraphicsLayout) removeAt(args ...interface{}) () {
  // removeAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsLayout8removeAtEi
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "removeAt", args)
  }

}

  // proto:  QGraphicsLayoutItem * QGraphicsLayout::itemAt(int i);
func (this *QGraphicsLayout) itemAt(args ...interface{}) () {
  // itemAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsLayout6itemAtEi
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "itemAt", args)
  }

}

  // proto:  void QGraphicsLayout::QGraphicsLayout(const QGraphicsLayout & );
func NewQGraphicsLayout(args ...interface{}) QGraphicsLayout {
  return QGraphicsLayout{}
}

  // proto:  void QGraphicsLayout::getContentsMargins(qreal * left, qreal * top, qreal * right, qreal * bottom);
func (this *QGraphicsLayout) getContentsMargins(args ...interface{}) () {
  // getContentsMargins(qreal *, qreal *, qreal *, qreal *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][1] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][2] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][3] = qtrt.DoubleTy(true) // "qreal *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsLayout18getContentsMarginsEPdS0_S0_S0_
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "getContentsMargins", args)
  }

}

  // proto:  void QGraphicsLayout::setContentsMargins(qreal left, qreal top, qreal right, qreal bottom);
func (this *QGraphicsLayout) setContentsMargins(args ...interface{}) () {
  // setContentsMargins(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsLayout18setContentsMarginsEdddd
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "setContentsMargins", args)
  }

}

  // proto:  void QGraphicsLayout::widgetEvent(QEvent * e);
func (this *QGraphicsLayout) widgetEvent(args ...interface{}) () {
  // widgetEvent(class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsLayout11widgetEventEP6QEvent
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "widgetEvent", args)
  }

}

  // proto: static bool QGraphicsLayout::instantInvalidatePropagation();
func (this *QGraphicsLayout) instantInvalidatePropagation_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "instantInvalidatePropagation", args)
  }

}

  // proto: static void QGraphicsLayout::setInstantInvalidatePropagation(bool enable);
func (this *QGraphicsLayout) setInstantInvalidatePropagation_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "setInstantInvalidatePropagation", args)
  }

}

  // proto:  void QGraphicsLayout::~QGraphicsLayout();
func (this *QGraphicsLayout) FreeQGraphicsLayout(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "~QGraphicsLayout", args)
  }

}

  // proto:  void QGraphicsLayout::activate();
func (this *QGraphicsLayout) activate(args ...interface{}) () {
  // activate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsLayout8activateEv
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "activate", args)
  }

}

  // proto:  int QGraphicsLayout::count();
func (this *QGraphicsLayout) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsLayout5countEv
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "count", args)
  }

}

// <= body block end

