package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtWidgets/qgraphicslayoutitem.h
// dst-file: /src/widgets/qgraphicslayoutitem.go
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
  // proto:  void QGraphicsLayoutItem::setMaximumWidth(qreal width);
extern void C_ZN19QGraphicsLayoutItem15setMaximumWidthEd(void* qthis, double arg0); // 4
  // proto:  QSizeF QGraphicsLayoutItem::preferredSize();
extern void* C_ZNK19QGraphicsLayoutItem13preferredSizeEv(void* qthis); // 4
  // proto:  void QGraphicsLayoutItem::setParentLayoutItem(QGraphicsLayoutItem * parent);
extern void C_ZN19QGraphicsLayoutItem19setParentLayoutItemEPS_(void* qthis, void* arg0); // 4
  // proto:  QGraphicsItem * QGraphicsLayoutItem::graphicsItem();
extern void C_ZNK19QGraphicsLayoutItem12graphicsItemEv(void* qthis); // 4
  // proto:  qreal QGraphicsLayoutItem::minimumHeight();
extern double C_ZNK19QGraphicsLayoutItem13minimumHeightEv(void* qthis); // 2
  // proto:  void QGraphicsLayoutItem::setMaximumSize(const QSizeF & size);
extern void C_ZN19QGraphicsLayoutItem14setMaximumSizeERK6QSizeF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsLayoutItem::setMaximumSize(qreal w, qreal h);
extern void C_ZN19QGraphicsLayoutItem14setMaximumSizeEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QGraphicsLayoutItem::setGeometry(const QRectF & rect);
extern void C_ZN19QGraphicsLayoutItem11setGeometryERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsLayoutItem::setMinimumHeight(qreal height);
extern void C_ZN19QGraphicsLayoutItem16setMinimumHeightEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsLayoutItem::~QGraphicsLayoutItem();
extern void C_ZN19QGraphicsLayoutItemD2Ev(void* qthis); // 4
  // proto:  void QGraphicsLayoutItem::setPreferredSize(const QSizeF & size);
extern void C_ZN19QGraphicsLayoutItem16setPreferredSizeERK6QSizeF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsLayoutItem::setPreferredSize(qreal w, qreal h);
extern void C_ZN19QGraphicsLayoutItem16setPreferredSizeEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QGraphicsLayoutItem::QGraphicsLayoutItem(QGraphicsLayoutItem * parent, bool isLayout);
extern void* C_ZN19QGraphicsLayoutItemC2EPS_b(void* arg0, bool arg1); // 3
  // proto:  QGraphicsLayoutItem * QGraphicsLayoutItem::parentLayoutItem();
extern void C_ZNK19QGraphicsLayoutItem16parentLayoutItemEv(void* qthis); // 4
  // proto:  void QGraphicsLayoutItem::setPreferredHeight(qreal height);
extern void C_ZN19QGraphicsLayoutItem18setPreferredHeightEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsLayoutItem::updateGeometry();
extern void C_ZN19QGraphicsLayoutItem14updateGeometryEv(void* qthis); // 4
  // proto:  QSizePolicy QGraphicsLayoutItem::sizePolicy();
extern void* C_ZNK19QGraphicsLayoutItem10sizePolicyEv(void* qthis); // 4
  // proto:  qreal QGraphicsLayoutItem::minimumWidth();
extern double C_ZNK19QGraphicsLayoutItem12minimumWidthEv(void* qthis); // 2
  // proto:  qreal QGraphicsLayoutItem::preferredHeight();
extern double C_ZNK19QGraphicsLayoutItem15preferredHeightEv(void* qthis); // 2
  // proto:  qreal QGraphicsLayoutItem::maximumHeight();
extern double C_ZNK19QGraphicsLayoutItem13maximumHeightEv(void* qthis); // 2
  // proto:  qreal QGraphicsLayoutItem::preferredWidth();
extern double C_ZNK19QGraphicsLayoutItem14preferredWidthEv(void* qthis); // 2
  // proto:  void QGraphicsLayoutItem::setMinimumSize(qreal w, qreal h);
extern void C_ZN19QGraphicsLayoutItem14setMinimumSizeEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QGraphicsLayoutItem::setMinimumSize(const QSizeF & size);
extern void C_ZN19QGraphicsLayoutItem14setMinimumSizeERK6QSizeF(void* qthis, void* arg0); // 4
  // proto:  qreal QGraphicsLayoutItem::maximumWidth();
extern double C_ZNK19QGraphicsLayoutItem12maximumWidthEv(void* qthis); // 2
  // proto:  void QGraphicsLayoutItem::setSizePolicy(const QSizePolicy & policy);
extern void C_ZN19QGraphicsLayoutItem13setSizePolicyERK11QSizePolicy(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsLayoutItem::setMinimumWidth(qreal width);
extern void C_ZN19QGraphicsLayoutItem15setMinimumWidthEd(void* qthis, double arg0); // 4
  // proto:  QRectF QGraphicsLayoutItem::geometry();
extern void* C_ZNK19QGraphicsLayoutItem8geometryEv(void* qthis); // 4
  // proto:  void QGraphicsLayoutItem::getContentsMargins(qreal * left, qreal * top, qreal * right, qreal * bottom);
extern void C_ZNK19QGraphicsLayoutItem18getContentsMarginsEPdS0_S0_S0_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3); // 4
  // proto:  QSizeF QGraphicsLayoutItem::minimumSize();
extern void* C_ZNK19QGraphicsLayoutItem11minimumSizeEv(void* qthis); // 4
  // proto:  QRectF QGraphicsLayoutItem::contentsRect();
extern void* C_ZNK19QGraphicsLayoutItem12contentsRectEv(void* qthis); // 4
  // proto:  bool QGraphicsLayoutItem::isLayout();
extern bool C_ZNK19QGraphicsLayoutItem8isLayoutEv(void* qthis); // 4
  // proto:  void QGraphicsLayoutItem::setPreferredWidth(qreal width);
extern void C_ZN19QGraphicsLayoutItem17setPreferredWidthEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsLayoutItem::setMaximumHeight(qreal height);
extern void C_ZN19QGraphicsLayoutItem16setMaximumHeightEd(void* qthis, double arg0); // 4
  // proto:  QSizeF QGraphicsLayoutItem::maximumSize();
extern void* C_ZNK19QGraphicsLayoutItem11maximumSizeEv(void* qthis); // 4
  // proto:  bool QGraphicsLayoutItem::ownedByLayout();
extern bool C_ZNK19QGraphicsLayoutItem13ownedByLayoutEv(void* qthis); // 4
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

// class sizeof(QGraphicsLayoutItem)=1
type QGraphicsLayoutItem struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// setMaximumWidth(qreal)
func (this *QGraphicsLayoutItem) Setmaximumwidth(args ...interface{}) () {
  // setMaximumWidth(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItem15setMaximumWidthEd
    // invoke: void setMaximumWidth(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsLayoutItem15setMaximumWidthEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setMaximumWidth", args)
  }

  return
}

// preferredSize()
func (this *QGraphicsLayoutItem) Preferredsize(args ...interface{}) (ret interface{}) {
  // preferredSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem13preferredSizeEv
    // invoke: QSizeF preferredSize()
    var ret0 = C.C_ZNK19QGraphicsLayoutItem13preferredSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSizeF{}) // "QSizeF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "preferredSize", args)
  }

  return
}

// setParentLayoutItem(class QGraphicsLayoutItem *)
func (this *QGraphicsLayoutItem) Setparentlayoutitem(args ...interface{}) () {
  // setParentLayoutItem(class QGraphicsLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsLayoutItem{}) // "QGraphicsLayoutItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItem19setParentLayoutItemEPS_
    // invoke: void setParentLayoutItem(class QGraphicsLayoutItem *)
    var arg0 = args[0].(QGraphicsLayoutItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsLayoutItem19setParentLayoutItemEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setParentLayoutItem", args)
  }

  return
}

// graphicsItem()
func (this *QGraphicsLayoutItem) Graphicsitem(args ...interface{}) () {
  // graphicsItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem12graphicsItemEv
    // invoke: QGraphicsItem * graphicsItem()
    C.C_ZNK19QGraphicsLayoutItem12graphicsItemEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "graphicsItem", args)
  }

  return
}

// minimumHeight()
func (this *QGraphicsLayoutItem) Minimumheight(args ...interface{}) (ret interface{}) {
  // minimumHeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem13minimumHeightEv
    // invoke: qreal minimumHeight()
    var ret0 = C.C_ZNK19QGraphicsLayoutItem13minimumHeightEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "minimumHeight", args)
  }

  return
}

// setMaximumSize(const class QSizeF &)
func (this *QGraphicsLayoutItem) Setmaximumsize(args ...interface{}) () {
  // setMaximumSize(const class QSizeF &)
  // setMaximumSize(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItem14setMaximumSizeERK6QSizeF
    // invoke: void setMaximumSize(const class QSizeF &)
    var arg0 = args[0].(QSizeF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsLayoutItem14setMaximumSizeERK6QSizeF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN19QGraphicsLayoutItem14setMaximumSizeEdd
    // invoke: void setMaximumSize(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsLayoutItem14setMaximumSizeEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setMaximumSize", args)
  }

  return
}

// setGeometry(const class QRectF &)
func (this *QGraphicsLayoutItem) Setgeometry(args ...interface{}) () {
  // setGeometry(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItem11setGeometryERK6QRectF
    // invoke: void setGeometry(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsLayoutItem11setGeometryERK6QRectF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setGeometry", args)
  }

  return
}

// setMinimumHeight(qreal)
func (this *QGraphicsLayoutItem) Setminimumheight(args ...interface{}) () {
  // setMinimumHeight(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItem16setMinimumHeightEd
    // invoke: void setMinimumHeight(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsLayoutItem16setMinimumHeightEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setMinimumHeight", args)
  }

  return
}

// ~QGraphicsLayoutItem()
func (this *QGraphicsLayoutItem) Freeqgraphicslayoutitem(args ...interface{}) () {
  // ~QGraphicsLayoutItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItemD0Ev
    // invoke: void ~QGraphicsLayoutItem()
    C.C_ZN19QGraphicsLayoutItemD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "~QGraphicsLayoutItem", args)
  }

  return
}

// setPreferredSize(const class QSizeF &)
func (this *QGraphicsLayoutItem) Setpreferredsize(args ...interface{}) () {
  // setPreferredSize(const class QSizeF &)
  // setPreferredSize(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItem16setPreferredSizeERK6QSizeF
    // invoke: void setPreferredSize(const class QSizeF &)
    var arg0 = args[0].(QSizeF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsLayoutItem16setPreferredSizeERK6QSizeF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN19QGraphicsLayoutItem16setPreferredSizeEdd
    // invoke: void setPreferredSize(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsLayoutItem16setPreferredSizeEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setPreferredSize", args)
  }

  return
}

// QGraphicsLayoutItem(class QGraphicsLayoutItem *, _Bool)
func NewQGraphicsLayoutItem(args ...interface{}) *QGraphicsLayoutItem {
  // QGraphicsLayoutItem(class QGraphicsLayoutItem *, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsLayoutItem{}) // "QGraphicsLayoutItem *"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItemC1EPS_b
    // invoke: void QGraphicsLayoutItem(class QGraphicsLayoutItem *, _Bool)
    var arg0 = args[0].(QGraphicsLayoutItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QGraphicsLayoutItemC2EPS_b(arg0, arg1)
    return &QGraphicsLayoutItem{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "QGraphicsLayoutItem", args)
  }

  return nil // QGraphicsLayoutItem{qclsinst:qthis}
}

// parentLayoutItem()
func (this *QGraphicsLayoutItem) Parentlayoutitem(args ...interface{}) () {
  // parentLayoutItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem16parentLayoutItemEv
    // invoke: QGraphicsLayoutItem * parentLayoutItem()
    C.C_ZNK19QGraphicsLayoutItem16parentLayoutItemEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "parentLayoutItem", args)
  }

  return
}

// setPreferredHeight(qreal)
func (this *QGraphicsLayoutItem) Setpreferredheight(args ...interface{}) () {
  // setPreferredHeight(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItem18setPreferredHeightEd
    // invoke: void setPreferredHeight(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsLayoutItem18setPreferredHeightEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setPreferredHeight", args)
  }

  return
}

// updateGeometry()
func (this *QGraphicsLayoutItem) Updategeometry(args ...interface{}) () {
  // updateGeometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItem14updateGeometryEv
    // invoke: void updateGeometry()
    C.C_ZN19QGraphicsLayoutItem14updateGeometryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "updateGeometry", args)
  }

  return
}

// sizePolicy()
func (this *QGraphicsLayoutItem) Sizepolicy(args ...interface{}) (ret interface{}) {
  // sizePolicy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem10sizePolicyEv
    // invoke: QSizePolicy sizePolicy()
    var ret0 = C.C_ZNK19QGraphicsLayoutItem10sizePolicyEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSizePolicy{}) // "QSizePolicy"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "sizePolicy", args)
  }

  return
}

// minimumWidth()
func (this *QGraphicsLayoutItem) Minimumwidth(args ...interface{}) (ret interface{}) {
  // minimumWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem12minimumWidthEv
    // invoke: qreal minimumWidth()
    var ret0 = C.C_ZNK19QGraphicsLayoutItem12minimumWidthEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "minimumWidth", args)
  }

  return
}

// preferredHeight()
func (this *QGraphicsLayoutItem) Preferredheight(args ...interface{}) (ret interface{}) {
  // preferredHeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem15preferredHeightEv
    // invoke: qreal preferredHeight()
    var ret0 = C.C_ZNK19QGraphicsLayoutItem15preferredHeightEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "preferredHeight", args)
  }

  return
}

// maximumHeight()
func (this *QGraphicsLayoutItem) Maximumheight(args ...interface{}) (ret interface{}) {
  // maximumHeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem13maximumHeightEv
    // invoke: qreal maximumHeight()
    var ret0 = C.C_ZNK19QGraphicsLayoutItem13maximumHeightEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "maximumHeight", args)
  }

  return
}

// preferredWidth()
func (this *QGraphicsLayoutItem) Preferredwidth(args ...interface{}) (ret interface{}) {
  // preferredWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem14preferredWidthEv
    // invoke: qreal preferredWidth()
    var ret0 = C.C_ZNK19QGraphicsLayoutItem14preferredWidthEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "preferredWidth", args)
  }

  return
}

// setMinimumSize(qreal, qreal)
func (this *QGraphicsLayoutItem) Setminimumsize(args ...interface{}) () {
  // setMinimumSize(qreal, qreal)
  // setMinimumSize(const class QSizeF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItem14setMinimumSizeEdd
    // invoke: void setMinimumSize(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsLayoutItem14setMinimumSizeEdd(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN19QGraphicsLayoutItem14setMinimumSizeERK6QSizeF
    // invoke: void setMinimumSize(const class QSizeF &)
    var arg0 = args[0].(QSizeF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsLayoutItem14setMinimumSizeERK6QSizeF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setMinimumSize", args)
  }

  return
}

// maximumWidth()
func (this *QGraphicsLayoutItem) Maximumwidth(args ...interface{}) (ret interface{}) {
  // maximumWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem12maximumWidthEv
    // invoke: qreal maximumWidth()
    var ret0 = C.C_ZNK19QGraphicsLayoutItem12maximumWidthEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "maximumWidth", args)
  }

  return
}

// setSizePolicy(const class QSizePolicy &)
func (this *QGraphicsLayoutItem) Setsizepolicy(args ...interface{}) () {
  // setSizePolicy(const class QSizePolicy &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSizePolicy{}) // "const QSizePolicy &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItem13setSizePolicyERK11QSizePolicy
    // invoke: void setSizePolicy(const class QSizePolicy &)
    var arg0 = args[0].(QSizePolicy).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsLayoutItem13setSizePolicyERK11QSizePolicy(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setSizePolicy", args)
  }

  return
}

// setMinimumWidth(qreal)
func (this *QGraphicsLayoutItem) Setminimumwidth(args ...interface{}) () {
  // setMinimumWidth(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItem15setMinimumWidthEd
    // invoke: void setMinimumWidth(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsLayoutItem15setMinimumWidthEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setMinimumWidth", args)
  }

  return
}

// geometry()
func (this *QGraphicsLayoutItem) Geometry(args ...interface{}) (ret interface{}) {
  // geometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem8geometryEv
    // invoke: QRectF geometry()
    var ret0 = C.C_ZNK19QGraphicsLayoutItem8geometryEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "geometry", args)
  }

  return
}

// getContentsMargins(qreal *, qreal *, qreal *, qreal *)
func (this *QGraphicsLayoutItem) Getcontentsmargins(args ...interface{}) () {
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
    // invoke: _ZNK19QGraphicsLayoutItem18getContentsMarginsEPdS0_S0_S0_
    // invoke: void getContentsMargins(qreal *, qreal *, qreal *, qreal *)
    var arg0 = (unsafe.Pointer)(args[0].(*float64))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*float64))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*float64))
    if false {fmt.Println(arg3)}
    C.C_ZNK19QGraphicsLayoutItem18getContentsMarginsEPdS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "getContentsMargins", args)
  }

  return
}

// minimumSize()
func (this *QGraphicsLayoutItem) Minimumsize(args ...interface{}) (ret interface{}) {
  // minimumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem11minimumSizeEv
    // invoke: QSizeF minimumSize()
    var ret0 = C.C_ZNK19QGraphicsLayoutItem11minimumSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSizeF{}) // "QSizeF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "minimumSize", args)
  }

  return
}

// contentsRect()
func (this *QGraphicsLayoutItem) Contentsrect(args ...interface{}) (ret interface{}) {
  // contentsRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem12contentsRectEv
    // invoke: QRectF contentsRect()
    var ret0 = C.C_ZNK19QGraphicsLayoutItem12contentsRectEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "contentsRect", args)
  }

  return
}

// isLayout()
func (this *QGraphicsLayoutItem) Islayout(args ...interface{}) (ret interface{}) {
  // isLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem8isLayoutEv
    // invoke: bool isLayout()
    var ret0 = C.C_ZNK19QGraphicsLayoutItem8isLayoutEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "isLayout", args)
  }

  return
}

// setPreferredWidth(qreal)
func (this *QGraphicsLayoutItem) Setpreferredwidth(args ...interface{}) () {
  // setPreferredWidth(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItem17setPreferredWidthEd
    // invoke: void setPreferredWidth(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsLayoutItem17setPreferredWidthEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setPreferredWidth", args)
  }

  return
}

// setMaximumHeight(qreal)
func (this *QGraphicsLayoutItem) Setmaximumheight(args ...interface{}) () {
  // setMaximumHeight(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItem16setMaximumHeightEd
    // invoke: void setMaximumHeight(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsLayoutItem16setMaximumHeightEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setMaximumHeight", args)
  }

  return
}

// maximumSize()
func (this *QGraphicsLayoutItem) Maximumsize(args ...interface{}) (ret interface{}) {
  // maximumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem11maximumSizeEv
    // invoke: QSizeF maximumSize()
    var ret0 = C.C_ZNK19QGraphicsLayoutItem11maximumSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSizeF{}) // "QSizeF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "maximumSize", args)
  }

  return
}

// ownedByLayout()
func (this *QGraphicsLayoutItem) Ownedbylayout(args ...interface{}) (ret interface{}) {
  // ownedByLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem13ownedByLayoutEv
    // invoke: bool ownedByLayout()
    var ret0 = C.C_ZNK19QGraphicsLayoutItem13ownedByLayoutEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "ownedByLayout", args)
  }

  return
}

// <= body block end

