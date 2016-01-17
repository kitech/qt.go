package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:13 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QGraphicsLayout::updateGeometry();
extern void _ZN15QGraphicsLayout14updateGeometryEv(void* qthis); // 4
  // proto:  void QGraphicsLayout::activate();
extern void _ZN15QGraphicsLayout8activateEv(void* qthis); // 4
  // proto:  bool QGraphicsLayout::isActivated();
extern void _ZNK15QGraphicsLayout11isActivatedEv(void* qthis); // 4
  // proto:  void QGraphicsLayout::invalidate();
extern void _ZN15QGraphicsLayout10invalidateEv(void* qthis); // 4
  // proto:  void QGraphicsLayout::getContentsMargins(qreal * left, qreal * top, qreal * right, qreal * bottom);
extern void _ZNK15QGraphicsLayout18getContentsMarginsEPdS0_S0_S0_(void* qthis, double* arg0, double* arg1, double* arg2, double* arg3); // 4
  // proto:  void QGraphicsLayout::setContentsMargins(qreal left, qreal top, qreal right, qreal bottom);
extern void _ZN15QGraphicsLayout18setContentsMarginsEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 4
  // proto:  void QGraphicsLayout::~QGraphicsLayout();
extern void _ZN15QGraphicsLayoutD2Ev(void* qthis); // 4
  // proto: static void QGraphicsLayout::setInstantInvalidatePropagation(bool enable);
extern void _ZN15QGraphicsLayout31setInstantInvalidatePropagationEb(bool arg0); // 4
  // proto: static bool QGraphicsLayout::instantInvalidatePropagation();
extern void _ZN15QGraphicsLayout28instantInvalidatePropagationEv(); // 4
  // proto:  void QGraphicsLayout::widgetEvent(QEvent * e);
extern void _ZN15QGraphicsLayout11widgetEventEP6QEvent(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsLayout::QGraphicsLayout(QGraphicsLayoutItem * parent);
extern void _ZN15QGraphicsLayoutC2EP19QGraphicsLayoutItem(void* qthis, void* arg0); // 3
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// updateGeometry()
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
    // invoke: void updateGeometry()
    C._ZN15QGraphicsLayout14updateGeometryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "updateGeometry", args)
  }

}

// activate()
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
    // invoke: void activate()
    C._ZN15QGraphicsLayout8activateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "activate", args)
  }

}

// isActivated()
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
    // invoke: bool isActivated()
    C._ZNK15QGraphicsLayout11isActivatedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "isActivated", args)
  }

}

// invalidate()
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
    // invoke: void invalidate()
    C._ZN15QGraphicsLayout10invalidateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "invalidate", args)
  }

}

// getContentsMargins(qreal *, qreal *, qreal *, qreal *)
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
    // invoke: void getContentsMargins(qreal *, qreal *, qreal *, qreal *)
    var arg0 = (*C.double)(args[0].(*float64))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.double)(args[1].(*float64))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.double)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.double)(args[3].(*float64))
    if false {fmt.Println(arg3)}
    C._ZNK15QGraphicsLayout18getContentsMarginsEPdS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "getContentsMargins", args)
  }

}

// setContentsMargins(qreal, qreal, qreal, qreal)
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
    // invoke: void setContentsMargins(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    C._ZN15QGraphicsLayout18setContentsMarginsEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "setContentsMargins", args)
  }

}

// ~QGraphicsLayout()
func (this *QGraphicsLayout) FreeQGraphicsLayout(args ...interface{}) () {
  // ~QGraphicsLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsLayoutD0Ev
    // invoke: void ~QGraphicsLayout()
    C._ZN15QGraphicsLayoutD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "~QGraphicsLayout", args)
  }

}

// setInstantInvalidatePropagation(_Bool)
func (this *QGraphicsLayout) setInstantInvalidatePropagation_s(args ...interface{}) () {
  // setInstantInvalidatePropagation(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsLayout31setInstantInvalidatePropagationEb
    // invoke: void setInstantInvalidatePropagation(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN15QGraphicsLayout31setInstantInvalidatePropagationEb(arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "setInstantInvalidatePropagation", args)
  }

}

// instantInvalidatePropagation()
func (this *QGraphicsLayout) instantInvalidatePropagation_s(args ...interface{}) () {
  // instantInvalidatePropagation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsLayout28instantInvalidatePropagationEv
    // invoke: bool instantInvalidatePropagation()
    C._ZN15QGraphicsLayout28instantInvalidatePropagationEv()
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "instantInvalidatePropagation", args)
  }

}

// widgetEvent(class QEvent *)
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
    // invoke: void widgetEvent(class QEvent *)
    var arg0 = args[0].(QEvent).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QGraphicsLayout11widgetEventEP6QEvent(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "widgetEvent", args)
  }

}

// QGraphicsLayout(class QGraphicsLayoutItem *)
func NewQGraphicsLayout(args ...interface{}) QGraphicsLayout {
  // QGraphicsLayout(class QGraphicsLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsLayoutItem{}) // "QGraphicsLayoutItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsLayoutC1EP19QGraphicsLayoutItem
    // invoke: void QGraphicsLayout(class QGraphicsLayoutItem *)
    var arg0 = args[0].(QGraphicsLayoutItem).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN15QGraphicsLayoutC2EP19QGraphicsLayoutItem(qthis, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "QGraphicsLayout", args)
  }

  return QGraphicsLayout{}
}

// <= body block end

