package qtwidgets
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
// src-file: /QtWidgets/qgraphicslinearlayout.h
// dst-file: /src/widgets/qgraphicslinearlayout.go
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
  // proto:  void QGraphicsLinearLayout::invalidate();
extern void C_ZN21QGraphicsLinearLayout10invalidateEv(void* qthis); // 4
  // proto:  Qt::Orientation QGraphicsLinearLayout::orientation();
extern void C_ZNK21QGraphicsLinearLayout11orientationEv(void* qthis); // 4
  // proto:  void QGraphicsLinearLayout::dump(int indent);
extern void C_ZNK21QGraphicsLinearLayout4dumpEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGraphicsLinearLayout::~QGraphicsLinearLayout();
extern void C_ZN21QGraphicsLinearLayoutD2Ev(void* qthis); // 4
  // proto:  void QGraphicsLinearLayout::insertStretch(int index, int stretch);
extern void C_ZN21QGraphicsLinearLayout13insertStretchEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  int QGraphicsLinearLayout::stretchFactor(QGraphicsLayoutItem * item);
extern int32_t C_ZNK21QGraphicsLinearLayout13stretchFactorEP19QGraphicsLayoutItem(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsLinearLayout::setItemSpacing(int index, qreal spacing);
extern void C_ZN21QGraphicsLinearLayout14setItemSpacingEid(void* qthis, int32_t arg0, double arg1); // 4
  // proto:  Qt::Alignment QGraphicsLinearLayout::alignment(QGraphicsLayoutItem * item);
extern void C_ZNK21QGraphicsLinearLayout9alignmentEP19QGraphicsLayoutItem(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsLinearLayout::setSpacing(qreal spacing);
extern void C_ZN21QGraphicsLinearLayout10setSpacingEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsLinearLayout::setStretchFactor(QGraphicsLayoutItem * item, int stretch);
extern void C_ZN21QGraphicsLinearLayout16setStretchFactorEP19QGraphicsLayoutItemi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QGraphicsLinearLayout::QGraphicsLinearLayout(QGraphicsLayoutItem * parent);
extern void* C_ZN21QGraphicsLinearLayoutC2EP19QGraphicsLayoutItem(void* arg0); // 3
  // proto:  void QGraphicsLinearLayout::addStretch(int stretch);
extern void C_ZN21QGraphicsLinearLayout10addStretchEi(void* qthis, int32_t arg0); // 2
  // proto:  void QGraphicsLinearLayout::setGeometry(const QRectF & rect);
extern void C_ZN21QGraphicsLinearLayout11setGeometryERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsLinearLayout::insertItem(int index, QGraphicsLayoutItem * item);
extern void C_ZN21QGraphicsLinearLayout10insertItemEiP19QGraphicsLayoutItem(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QGraphicsLinearLayout::addItem(QGraphicsLayoutItem * item);
extern void C_ZN21QGraphicsLinearLayout7addItemEP19QGraphicsLayoutItem(void* qthis, void* arg0); // 2
  // proto:  qreal QGraphicsLinearLayout::spacing();
extern double C_ZNK21QGraphicsLinearLayout7spacingEv(void* qthis); // 4
  // proto:  void QGraphicsLinearLayout::removeAt(int index);
extern void C_ZN21QGraphicsLinearLayout8removeAtEi(void* qthis, int32_t arg0); // 4
  // proto:  int QGraphicsLinearLayout::count();
extern int32_t C_ZNK21QGraphicsLinearLayout5countEv(void* qthis); // 4
  // proto:  QGraphicsLayoutItem * QGraphicsLinearLayout::itemAt(int index);
extern void C_ZNK21QGraphicsLinearLayout6itemAtEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGraphicsLinearLayout::removeItem(QGraphicsLayoutItem * item);
extern void C_ZN21QGraphicsLinearLayout10removeItemEP19QGraphicsLayoutItem(void* qthis, void* arg0); // 4
  // proto:  qreal QGraphicsLinearLayout::itemSpacing(int index);
extern double C_ZNK21QGraphicsLinearLayout11itemSpacingEi(void* qthis, int32_t arg0); // 4
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

// class sizeof(QGraphicsLinearLayout)=1
type QGraphicsLinearLayout struct {
  /*qbase*/ QGraphicsLayout;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// invalidate()
func (this *QGraphicsLinearLayout) Invalidate(args ...interface{}) () {
  // invalidate()
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
    // invoke: _ZN21QGraphicsLinearLayout10invalidateEv
    // invoke: void invalidate()
    C.C_ZN21QGraphicsLinearLayout10invalidateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "invalidate", args)
  }

  return
}

// orientation()
func (this *QGraphicsLinearLayout) Orientation(args ...interface{}) () {
  // orientation()
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
    // invoke: _ZNK21QGraphicsLinearLayout11orientationEv
    // invoke: Qt::Orientation orientation()
    C.C_ZNK21QGraphicsLinearLayout11orientationEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "orientation", args)
  }

  return
}

// dump(int)
func (this *QGraphicsLinearLayout) Dump(args ...interface{}) () {
  // dump(int)
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
    // invoke: _ZNK21QGraphicsLinearLayout4dumpEi
    // invoke: void dump(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK21QGraphicsLinearLayout4dumpEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "dump", args)
  }

  return
}

// ~QGraphicsLinearLayout()
func (this *QGraphicsLinearLayout) Freeqgraphicslinearlayout(args ...interface{}) () {
  // ~QGraphicsLinearLayout()
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
    // invoke: _ZN21QGraphicsLinearLayoutD0Ev
    // invoke: void ~QGraphicsLinearLayout()
    C.C_ZN21QGraphicsLinearLayoutD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "~QGraphicsLinearLayout", args)
  }

  return
}

// insertStretch(int, int)
func (this *QGraphicsLinearLayout) Insertstretch(args ...interface{}) () {
  // insertStretch(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout13insertStretchEii
    // invoke: void insertStretch(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN21QGraphicsLinearLayout13insertStretchEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "insertStretch", args)
  }

  return
}

// stretchFactor(class QGraphicsLayoutItem *)
func (this *QGraphicsLinearLayout) Stretchfactor(args ...interface{}) (ret interface{}) {
  // stretchFactor(class QGraphicsLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsLayoutItem{}) // "QGraphicsLayoutItem *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QGraphicsLinearLayout13stretchFactorEP19QGraphicsLayoutItem
    // invoke: int stretchFactor(class QGraphicsLayoutItem *)
    var arg0 = args[0].(*QGraphicsLayoutItem).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK21QGraphicsLinearLayout13stretchFactorEP19QGraphicsLayoutItem(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "stretchFactor", args)
  }

  return
}

// setItemSpacing(int, qreal)
func (this *QGraphicsLinearLayout) Setitemspacing(args ...interface{}) () {
  // setItemSpacing(int, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout14setItemSpacingEid
    // invoke: void setItemSpacing(int, qreal)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN21QGraphicsLinearLayout14setItemSpacingEid(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "setItemSpacing", args)
  }

  return
}

// alignment(class QGraphicsLayoutItem *)
func (this *QGraphicsLinearLayout) Alignment(args ...interface{}) () {
  // alignment(class QGraphicsLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsLayoutItem{}) // "QGraphicsLayoutItem *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QGraphicsLinearLayout9alignmentEP19QGraphicsLayoutItem
    // invoke: Qt::Alignment alignment(class QGraphicsLayoutItem *)
    var arg0 = args[0].(*QGraphicsLayoutItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK21QGraphicsLinearLayout9alignmentEP19QGraphicsLayoutItem(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "alignment", args)
  }

  return
}

// setSpacing(qreal)
func (this *QGraphicsLinearLayout) Setspacing(args ...interface{}) () {
  // setSpacing(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout10setSpacingEd
    // invoke: void setSpacing(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN21QGraphicsLinearLayout10setSpacingEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "setSpacing", args)
  }

  return
}

// setStretchFactor(class QGraphicsLayoutItem *, int)
func (this *QGraphicsLinearLayout) Setstretchfactor(args ...interface{}) () {
  // setStretchFactor(class QGraphicsLayoutItem *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsLayoutItem{}) // "QGraphicsLayoutItem *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout16setStretchFactorEP19QGraphicsLayoutItemi
    // invoke: void setStretchFactor(class QGraphicsLayoutItem *, int)
    var arg0 = args[0].(*QGraphicsLayoutItem).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN21QGraphicsLinearLayout16setStretchFactorEP19QGraphicsLayoutItemi(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "setStretchFactor", args)
  }

  return
}

// QGraphicsLinearLayout(class QGraphicsLayoutItem *)
func NewQGraphicsLinearLayout(args ...interface{}) *QGraphicsLinearLayout {
  // QGraphicsLinearLayout(class QGraphicsLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsLayoutItem{}) // "QGraphicsLayoutItem *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayoutC1EP19QGraphicsLayoutItem
    // invoke: void QGraphicsLinearLayout(class QGraphicsLayoutItem *)
    var arg0 = args[0].(*QGraphicsLayoutItem).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN21QGraphicsLinearLayoutC2EP19QGraphicsLayoutItem(arg0)
    return &QGraphicsLinearLayout{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "QGraphicsLinearLayout", args)
  }

  return nil // QGraphicsLinearLayout{Qclsinst:qthis}
}

// addStretch(int)
func (this *QGraphicsLinearLayout) Addstretch(args ...interface{}) () {
  // addStretch(int)
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
    // invoke: _ZN21QGraphicsLinearLayout10addStretchEi
    // invoke: void addStretch(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN21QGraphicsLinearLayout10addStretchEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "addStretch", args)
  }

  return
}

// setGeometry(const class QRectF &)
func (this *QGraphicsLinearLayout) Setgeometry(args ...interface{}) () {
  // setGeometry(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout11setGeometryERK6QRectF
    // invoke: void setGeometry(const class QRectF &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN21QGraphicsLinearLayout11setGeometryERK6QRectF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "setGeometry", args)
  }

  return
}

// insertItem(int, class QGraphicsLayoutItem *)
func (this *QGraphicsLinearLayout) Insertitem(args ...interface{}) () {
  // insertItem(int, class QGraphicsLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QGraphicsLayoutItem{}) // "QGraphicsLayoutItem *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout10insertItemEiP19QGraphicsLayoutItem
    // invoke: void insertItem(int, class QGraphicsLayoutItem *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QGraphicsLayoutItem).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN21QGraphicsLinearLayout10insertItemEiP19QGraphicsLayoutItem(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "insertItem", args)
  }

  return
}

// addItem(class QGraphicsLayoutItem *)
func (this *QGraphicsLinearLayout) Additem(args ...interface{}) () {
  // addItem(class QGraphicsLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsLayoutItem{}) // "QGraphicsLayoutItem *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout7addItemEP19QGraphicsLayoutItem
    // invoke: void addItem(class QGraphicsLayoutItem *)
    var arg0 = args[0].(*QGraphicsLayoutItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN21QGraphicsLinearLayout7addItemEP19QGraphicsLayoutItem(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "addItem", args)
  }

  return
}

// spacing()
func (this *QGraphicsLinearLayout) Spacing(args ...interface{}) (ret interface{}) {
  // spacing()
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
    // invoke: _ZNK21QGraphicsLinearLayout7spacingEv
    // invoke: qreal spacing()
    var ret0 = C.C_ZNK21QGraphicsLinearLayout7spacingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "spacing", args)
  }

  return
}

// removeAt(int)
func (this *QGraphicsLinearLayout) Removeat(args ...interface{}) () {
  // removeAt(int)
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
    // invoke: _ZN21QGraphicsLinearLayout8removeAtEi
    // invoke: void removeAt(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN21QGraphicsLinearLayout8removeAtEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "removeAt", args)
  }

  return
}

// count()
func (this *QGraphicsLinearLayout) Count(args ...interface{}) (ret interface{}) {
  // count()
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
    // invoke: _ZNK21QGraphicsLinearLayout5countEv
    // invoke: int count()
    var ret0 = C.C_ZNK21QGraphicsLinearLayout5countEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "count", args)
  }

  return
}

// itemAt(int)
func (this *QGraphicsLinearLayout) Itemat(args ...interface{}) () {
  // itemAt(int)
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
    // invoke: _ZNK21QGraphicsLinearLayout6itemAtEi
    // invoke: QGraphicsLayoutItem * itemAt(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK21QGraphicsLinearLayout6itemAtEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "itemAt", args)
  }

  return
}

// removeItem(class QGraphicsLayoutItem *)
func (this *QGraphicsLinearLayout) Removeitem(args ...interface{}) () {
  // removeItem(class QGraphicsLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsLayoutItem{}) // "QGraphicsLayoutItem *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout10removeItemEP19QGraphicsLayoutItem
    // invoke: void removeItem(class QGraphicsLayoutItem *)
    var arg0 = args[0].(*QGraphicsLayoutItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN21QGraphicsLinearLayout10removeItemEP19QGraphicsLayoutItem(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "removeItem", args)
  }

  return
}

// itemSpacing(int)
func (this *QGraphicsLinearLayout) Itemspacing(args ...interface{}) (ret interface{}) {
  // itemSpacing(int)
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
    // invoke: _ZNK21QGraphicsLinearLayout11itemSpacingEi
    // invoke: qreal itemSpacing(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK21QGraphicsLinearLayout11itemSpacingEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "itemSpacing", args)
  }

  return
}

// <= body block end

