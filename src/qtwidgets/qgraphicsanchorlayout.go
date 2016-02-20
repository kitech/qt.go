package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
// src-file: /QtWidgets/qgraphicsanchorlayout.h
// dst-file: /src/widgets/qgraphicsanchorlayout.go
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
  // proto:  void QGraphicsAnchorLayout::setSpacing(qreal spacing);
extern void C_ZN21QGraphicsAnchorLayout10setSpacingEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsAnchorLayout::~QGraphicsAnchorLayout();
extern void C_ZN21QGraphicsAnchorLayoutD2Ev(void* qthis); // 4
  // proto:  QGraphicsLayoutItem * QGraphicsAnchorLayout::itemAt(int index);
extern void C_ZNK21QGraphicsAnchorLayout6itemAtEi(void* qthis, int32_t arg0); // 4
  // proto:  qreal QGraphicsAnchorLayout::horizontalSpacing();
extern double C_ZNK21QGraphicsAnchorLayout17horizontalSpacingEv(void* qthis); // 4
  // proto:  void QGraphicsAnchorLayout::setGeometry(const QRectF & rect);
extern void C_ZN21QGraphicsAnchorLayout11setGeometryERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsAnchorLayout::invalidate();
extern void C_ZN21QGraphicsAnchorLayout10invalidateEv(void* qthis); // 4
  // proto:  int QGraphicsAnchorLayout::count();
extern int32_t C_ZNK21QGraphicsAnchorLayout5countEv(void* qthis); // 4
  // proto:  void QGraphicsAnchorLayout::setVerticalSpacing(qreal spacing);
extern void C_ZN21QGraphicsAnchorLayout18setVerticalSpacingEd(void* qthis, double arg0); // 4
  // proto:  qreal QGraphicsAnchorLayout::verticalSpacing();
extern double C_ZNK21QGraphicsAnchorLayout15verticalSpacingEv(void* qthis); // 4
  // proto:  void QGraphicsAnchorLayout::removeAt(int index);
extern void C_ZN21QGraphicsAnchorLayout8removeAtEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGraphicsAnchorLayout::QGraphicsAnchorLayout(QGraphicsLayoutItem * parent);
extern void* C_ZN21QGraphicsAnchorLayoutC2EP19QGraphicsLayoutItem(void* arg0); // 3
  // proto:  void QGraphicsAnchorLayout::setHorizontalSpacing(qreal spacing);
extern void C_ZN21QGraphicsAnchorLayout20setHorizontalSpacingEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsAnchor::setSpacing(qreal spacing);
extern void C_ZN15QGraphicsAnchor10setSpacingEd(void* qthis, double arg0); // 4
  // proto:  QSizePolicy::Policy QGraphicsAnchor::sizePolicy();
extern void C_ZNK15QGraphicsAnchor10sizePolicyEv(void* qthis); // 4
  // proto:  const QMetaObject * QGraphicsAnchor::metaObject();
extern void C_ZNK15QGraphicsAnchor10metaObjectEv(void* qthis); // 4
  // proto:  qreal QGraphicsAnchor::spacing();
extern double C_ZNK15QGraphicsAnchor7spacingEv(void* qthis); // 4
  // proto:  void QGraphicsAnchor::unsetSpacing();
extern void C_ZN15QGraphicsAnchor12unsetSpacingEv(void* qthis); // 4
  // proto:  void QGraphicsAnchor::~QGraphicsAnchor();
extern void C_ZN15QGraphicsAnchorD2Ev(void* qthis); // 4
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

// class sizeof(QGraphicsAnchorLayout)=1
type QGraphicsAnchorLayout struct {
  /*qbase*/ QGraphicsLayout;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGraphicsAnchor)=1
type QGraphicsAnchor struct {
  /*qbase*/ qtcore.QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// setSpacing(qreal)
func (this *QGraphicsAnchorLayout) Setspacing(args ...interface{}) () {
  // setSpacing(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsAnchorLayout10setSpacingEd
    // invoke: void setSpacing(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN21QGraphicsAnchorLayout10setSpacingEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "setSpacing", args)
  }

  return
}

// ~QGraphicsAnchorLayout()
func (this *QGraphicsAnchorLayout) Freeqgraphicsanchorlayout(args ...interface{}) () {
  // ~QGraphicsAnchorLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsAnchorLayoutD0Ev
    // invoke: void ~QGraphicsAnchorLayout()
    C.C_ZN21QGraphicsAnchorLayoutD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "~QGraphicsAnchorLayout", args)
  }

  return
}

// itemAt(int)
func (this *QGraphicsAnchorLayout) Itemat(args ...interface{}) () {
  // itemAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QGraphicsAnchorLayout6itemAtEi
    // invoke: QGraphicsLayoutItem * itemAt(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK21QGraphicsAnchorLayout6itemAtEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "itemAt", args)
  }

  return
}

// horizontalSpacing()
func (this *QGraphicsAnchorLayout) Horizontalspacing(args ...interface{}) (ret interface{}) {
  // horizontalSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QGraphicsAnchorLayout17horizontalSpacingEv
    // invoke: qreal horizontalSpacing()
    var ret0 = C.C_ZNK21QGraphicsAnchorLayout17horizontalSpacingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "horizontalSpacing", args)
  }

  return
}

// setGeometry(const class QRectF &)
func (this *QGraphicsAnchorLayout) Setgeometry(args ...interface{}) () {
  // setGeometry(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsAnchorLayout11setGeometryERK6QRectF
    // invoke: void setGeometry(const class QRectF &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN21QGraphicsAnchorLayout11setGeometryERK6QRectF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "setGeometry", args)
  }

  return
}

// invalidate()
func (this *QGraphicsAnchorLayout) Invalidate(args ...interface{}) () {
  // invalidate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsAnchorLayout10invalidateEv
    // invoke: void invalidate()
    C.C_ZN21QGraphicsAnchorLayout10invalidateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "invalidate", args)
  }

  return
}

// count()
func (this *QGraphicsAnchorLayout) Count(args ...interface{}) (ret interface{}) {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QGraphicsAnchorLayout5countEv
    // invoke: int count()
    var ret0 = C.C_ZNK21QGraphicsAnchorLayout5countEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "count", args)
  }

  return
}

// setVerticalSpacing(qreal)
func (this *QGraphicsAnchorLayout) Setverticalspacing(args ...interface{}) () {
  // setVerticalSpacing(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsAnchorLayout18setVerticalSpacingEd
    // invoke: void setVerticalSpacing(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN21QGraphicsAnchorLayout18setVerticalSpacingEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "setVerticalSpacing", args)
  }

  return
}

// verticalSpacing()
func (this *QGraphicsAnchorLayout) Verticalspacing(args ...interface{}) (ret interface{}) {
  // verticalSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QGraphicsAnchorLayout15verticalSpacingEv
    // invoke: qreal verticalSpacing()
    var ret0 = C.C_ZNK21QGraphicsAnchorLayout15verticalSpacingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "verticalSpacing", args)
  }

  return
}

// removeAt(int)
func (this *QGraphicsAnchorLayout) Removeat(args ...interface{}) () {
  // removeAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsAnchorLayout8removeAtEi
    // invoke: void removeAt(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN21QGraphicsAnchorLayout8removeAtEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "removeAt", args)
  }

  return
}

// QGraphicsAnchorLayout(class QGraphicsLayoutItem *)
func NewQGraphicsAnchorLayout(args ...interface{}) *QGraphicsAnchorLayout {
  // QGraphicsAnchorLayout(class QGraphicsLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsLayoutItem{}) // "QGraphicsLayoutItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsAnchorLayoutC1EP19QGraphicsLayoutItem
    // invoke: void QGraphicsAnchorLayout(class QGraphicsLayoutItem *)
    var arg0 = args[0].(*QGraphicsLayoutItem).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN21QGraphicsAnchorLayoutC2EP19QGraphicsLayoutItem(arg0)
    return &QGraphicsAnchorLayout{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "QGraphicsAnchorLayout", args)
  }

  return nil // QGraphicsAnchorLayout{Qclsinst:qthis}
}

// setHorizontalSpacing(qreal)
func (this *QGraphicsAnchorLayout) Sethorizontalspacing(args ...interface{}) () {
  // setHorizontalSpacing(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsAnchorLayout20setHorizontalSpacingEd
    // invoke: void setHorizontalSpacing(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN21QGraphicsAnchorLayout20setHorizontalSpacingEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "setHorizontalSpacing", args)
  }

  return
}

// setSpacing(qreal)
func (this *QGraphicsAnchor) Setspacing(args ...interface{}) () {
  // setSpacing(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsAnchor10setSpacingEd
    // invoke: void setSpacing(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN15QGraphicsAnchor10setSpacingEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsAnchor", "setSpacing", args)
  }

  return
}

// sizePolicy()
func (this *QGraphicsAnchor) Sizepolicy(args ...interface{}) () {
  // sizePolicy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsAnchor10sizePolicyEv
    // invoke: QSizePolicy::Policy sizePolicy()
    C.C_ZNK15QGraphicsAnchor10sizePolicyEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsAnchor", "sizePolicy", args)
  }

  return
}

// metaObject()
func (this *QGraphicsAnchor) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsAnchor10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK15QGraphicsAnchor10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsAnchor", "metaObject", args)
  }

  return
}

// spacing()
func (this *QGraphicsAnchor) Spacing(args ...interface{}) (ret interface{}) {
  // spacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsAnchor7spacingEv
    // invoke: qreal spacing()
    var ret0 = C.C_ZNK15QGraphicsAnchor7spacingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsAnchor", "spacing", args)
  }

  return
}

// unsetSpacing()
func (this *QGraphicsAnchor) Unsetspacing(args ...interface{}) () {
  // unsetSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsAnchor12unsetSpacingEv
    // invoke: void unsetSpacing()
    C.C_ZN15QGraphicsAnchor12unsetSpacingEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsAnchor", "unsetSpacing", args)
  }

  return
}

// ~QGraphicsAnchor()
func (this *QGraphicsAnchor) Freeqgraphicsanchor(args ...interface{}) () {
  // ~QGraphicsAnchor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsAnchorD0Ev
    // invoke: void ~QGraphicsAnchor()
    C.C_ZN15QGraphicsAnchorD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsAnchor", "~QGraphicsAnchor", args)
  }

  return
}

// <= body block end

