package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
  // proto:  void QGraphicsLayout::updateGeometry();
extern void C_ZN15QGraphicsLayout14updateGeometryEv(void* qthis); // 4
  // proto:  void QGraphicsLayout::activate();
extern void C_ZN15QGraphicsLayout8activateEv(void* qthis); // 4
  // proto:  bool QGraphicsLayout::isActivated();
extern bool C_ZNK15QGraphicsLayout11isActivatedEv(void* qthis); // 4
  // proto:  void QGraphicsLayout::invalidate();
extern void C_ZN15QGraphicsLayout10invalidateEv(void* qthis); // 4
  // proto:  void QGraphicsLayout::getContentsMargins(qreal * left, qreal * top, qreal * right, qreal * bottom);
extern void C_ZNK15QGraphicsLayout18getContentsMarginsEPdS0_S0_S0_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3); // 4
  // proto:  void QGraphicsLayout::setContentsMargins(qreal left, qreal top, qreal right, qreal bottom);
extern void C_ZN15QGraphicsLayout18setContentsMarginsEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 4
  // proto:  void QGraphicsLayout::~QGraphicsLayout();
extern void C_ZN15QGraphicsLayoutD2Ev(void* qthis); // 4
  // proto: static void QGraphicsLayout::setInstantInvalidatePropagation(bool enable);
extern void C_ZN15QGraphicsLayout31setInstantInvalidatePropagationEb(bool arg0); // 4
  // proto: static bool QGraphicsLayout::instantInvalidatePropagation();
extern bool C_ZN15QGraphicsLayout28instantInvalidatePropagationEv(); // 4
  // proto:  void QGraphicsLayout::widgetEvent(QEvent * e);
extern void C_ZN15QGraphicsLayout11widgetEventEP6QEvent(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsLayout::QGraphicsLayout(QGraphicsLayoutItem * parent);
extern void* C_ZN15QGraphicsLayoutC2EP19QGraphicsLayoutItem(void* arg0); // 3
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

// class sizeof(QGraphicsLayout)=1
type QGraphicsLayout struct {
  /*qbase*/ QGraphicsLayoutItem;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// updateGeometry()
func (this *QGraphicsLayout) Updategeometry(args ...interface{}) () {
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
    C.C_ZN15QGraphicsLayout14updateGeometryEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "updateGeometry", args)
  }

  return
}

// activate()
func (this *QGraphicsLayout) Activate(args ...interface{}) () {
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
    C.C_ZN15QGraphicsLayout8activateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "activate", args)
  }

  return
}

// isActivated()
func (this *QGraphicsLayout) Isactivated(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QGraphicsLayout11isActivatedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "isActivated", args)
  }

  return
}

// invalidate()
func (this *QGraphicsLayout) Invalidate(args ...interface{}) () {
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
    C.C_ZN15QGraphicsLayout10invalidateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "invalidate", args)
  }

  return
}

// getContentsMargins(qreal *, qreal *, qreal *, qreal *)
func (this *QGraphicsLayout) Getcontentsmargins(args ...interface{}) () {
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
    var arg0 = (unsafe.Pointer)(args[0].(*float64))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*float64))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*float64))
    if false {fmt.Println(arg3)}
    C.C_ZNK15QGraphicsLayout18getContentsMarginsEPdS0_S0_S0_(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "getContentsMargins", args)
  }

  return
}

// setContentsMargins(qreal, qreal, qreal, qreal)
func (this *QGraphicsLayout) Setcontentsmargins(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN15QGraphicsLayout18setContentsMarginsEdddd(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "setContentsMargins", args)
  }

  return
}

// ~QGraphicsLayout()
func (this *QGraphicsLayout) Freeqgraphicslayout(args ...interface{}) () {
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
    C.C_ZN15QGraphicsLayoutD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "~QGraphicsLayout", args)
  }

  return
}

// setInstantInvalidatePropagation(_Bool)
func (this *QGraphicsLayout) Setinstantinvalidatepropagation_S(args ...interface{}) () {
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
    C.C_ZN15QGraphicsLayout31setInstantInvalidatePropagationEb(arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "setInstantInvalidatePropagation", args)
  }

  return
}

// instantInvalidatePropagation()
func (this *QGraphicsLayout) Instantinvalidatepropagation_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN15QGraphicsLayout28instantInvalidatePropagationEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "instantInvalidatePropagation", args)
  }

  return
}

// widgetEvent(class QEvent *)
func (this *QGraphicsLayout) Widgetevent(args ...interface{}) () {
  // widgetEvent(class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsLayout11widgetEventEP6QEvent
    // invoke: void widgetEvent(class QEvent *)
    var arg0 = args[0].(*qtcore.QEvent).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QGraphicsLayout11widgetEventEP6QEvent(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "widgetEvent", args)
  }

  return
}

// QGraphicsLayout(class QGraphicsLayoutItem *)
func NewQGraphicsLayout(args ...interface{}) *QGraphicsLayout {
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
    var arg0 = args[0].(*QGraphicsLayoutItem).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QGraphicsLayoutC2EP19QGraphicsLayoutItem(arg0)
    return &QGraphicsLayout{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "QGraphicsLayout", args)
  }

  return nil // QGraphicsLayout{Qclsinst:qthis}
}

// <= body block end

