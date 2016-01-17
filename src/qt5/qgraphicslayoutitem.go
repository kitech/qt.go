package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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
extern void _ZN19QGraphicsLayoutItem15setMaximumWidthEd(void* qthis, double arg0); // 4
  // proto:  QSizeF QGraphicsLayoutItem::preferredSize();
extern void _ZNK19QGraphicsLayoutItem13preferredSizeEv(void* qthis); // 4
  // proto:  void QGraphicsLayoutItem::setParentLayoutItem(QGraphicsLayoutItem * parent);
extern void _ZN19QGraphicsLayoutItem19setParentLayoutItemEPS_(void* qthis, void* arg0); // 4
  // proto:  QGraphicsItem * QGraphicsLayoutItem::graphicsItem();
extern void _ZNK19QGraphicsLayoutItem12graphicsItemEv(void* qthis); // 4
  // proto:  qreal QGraphicsLayoutItem::minimumHeight();
extern void _ZNK19QGraphicsLayoutItem13minimumHeightEv(void* qthis); // 2
  // proto:  void QGraphicsLayoutItem::setMaximumSize(const QSizeF & size);
extern void _ZN19QGraphicsLayoutItem14setMaximumSizeERK6QSizeF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsLayoutItem::setMaximumSize(qreal w, qreal h);
extern void _ZN19QGraphicsLayoutItem14setMaximumSizeEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QGraphicsLayoutItem::setGeometry(const QRectF & rect);
extern void _ZN19QGraphicsLayoutItem11setGeometryERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsLayoutItem::setMinimumHeight(qreal height);
extern void _ZN19QGraphicsLayoutItem16setMinimumHeightEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsLayoutItem::~QGraphicsLayoutItem();
extern void _ZN19QGraphicsLayoutItemD2Ev(void* qthis); // 4
  // proto:  void QGraphicsLayoutItem::setPreferredSize(const QSizeF & size);
extern void _ZN19QGraphicsLayoutItem16setPreferredSizeERK6QSizeF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsLayoutItem::setPreferredSize(qreal w, qreal h);
extern void _ZN19QGraphicsLayoutItem16setPreferredSizeEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QGraphicsLayoutItem::QGraphicsLayoutItem(QGraphicsLayoutItem * parent, bool isLayout);
extern void _ZN19QGraphicsLayoutItemC2EPS_b(void* qthis, void* arg0, bool arg1); // 3
  // proto:  QGraphicsLayoutItem * QGraphicsLayoutItem::parentLayoutItem();
extern void _ZNK19QGraphicsLayoutItem16parentLayoutItemEv(void* qthis); // 4
  // proto:  void QGraphicsLayoutItem::setPreferredHeight(qreal height);
extern void _ZN19QGraphicsLayoutItem18setPreferredHeightEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsLayoutItem::updateGeometry();
extern void _ZN19QGraphicsLayoutItem14updateGeometryEv(void* qthis); // 4
  // proto:  QSizePolicy QGraphicsLayoutItem::sizePolicy();
extern void _ZNK19QGraphicsLayoutItem10sizePolicyEv(void* qthis); // 4
  // proto:  qreal QGraphicsLayoutItem::minimumWidth();
extern void _ZNK19QGraphicsLayoutItem12minimumWidthEv(void* qthis); // 2
  // proto:  qreal QGraphicsLayoutItem::preferredHeight();
extern void _ZNK19QGraphicsLayoutItem15preferredHeightEv(void* qthis); // 2
  // proto:  qreal QGraphicsLayoutItem::maximumHeight();
extern void _ZNK19QGraphicsLayoutItem13maximumHeightEv(void* qthis); // 2
  // proto:  qreal QGraphicsLayoutItem::preferredWidth();
extern void _ZNK19QGraphicsLayoutItem14preferredWidthEv(void* qthis); // 2
  // proto:  void QGraphicsLayoutItem::setMinimumSize(qreal w, qreal h);
extern void _ZN19QGraphicsLayoutItem14setMinimumSizeEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QGraphicsLayoutItem::setMinimumSize(const QSizeF & size);
extern void _ZN19QGraphicsLayoutItem14setMinimumSizeERK6QSizeF(void* qthis, void* arg0); // 4
  // proto:  qreal QGraphicsLayoutItem::maximumWidth();
extern void _ZNK19QGraphicsLayoutItem12maximumWidthEv(void* qthis); // 2
  // proto:  void QGraphicsLayoutItem::setSizePolicy(const QSizePolicy & policy);
extern void _ZN19QGraphicsLayoutItem13setSizePolicyERK11QSizePolicy(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsLayoutItem::setMinimumWidth(qreal width);
extern void _ZN19QGraphicsLayoutItem15setMinimumWidthEd(void* qthis, double arg0); // 4
  // proto:  QRectF QGraphicsLayoutItem::geometry();
extern void _ZNK19QGraphicsLayoutItem8geometryEv(void* qthis); // 4
  // proto:  void QGraphicsLayoutItem::getContentsMargins(qreal * left, qreal * top, qreal * right, qreal * bottom);
extern void _ZNK19QGraphicsLayoutItem18getContentsMarginsEPdS0_S0_S0_(void* qthis, double* arg0, double* arg1, double* arg2, double* arg3); // 4
  // proto:  QSizeF QGraphicsLayoutItem::minimumSize();
extern void _ZNK19QGraphicsLayoutItem11minimumSizeEv(void* qthis); // 4
  // proto:  QRectF QGraphicsLayoutItem::contentsRect();
extern void _ZNK19QGraphicsLayoutItem12contentsRectEv(void* qthis); // 4
  // proto:  bool QGraphicsLayoutItem::isLayout();
extern void _ZNK19QGraphicsLayoutItem8isLayoutEv(void* qthis); // 4
  // proto:  void QGraphicsLayoutItem::setPreferredWidth(qreal width);
extern void _ZN19QGraphicsLayoutItem17setPreferredWidthEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsLayoutItem::setMaximumHeight(qreal height);
extern void _ZN19QGraphicsLayoutItem16setMaximumHeightEd(void* qthis, double arg0); // 4
  // proto:  QSizeF QGraphicsLayoutItem::maximumSize();
extern void _ZNK19QGraphicsLayoutItem11maximumSizeEv(void* qthis); // 4
  // proto:  bool QGraphicsLayoutItem::ownedByLayout();
extern void _ZNK19QGraphicsLayoutItem13ownedByLayoutEv(void* qthis); // 4
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
func (this *QGraphicsLayoutItem) setMaximumWidth(args ...interface{}) () {
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
    C._ZN19QGraphicsLayoutItem15setMaximumWidthEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setMaximumWidth", args)
  }

}

// preferredSize()
func (this *QGraphicsLayoutItem) preferredSize(args ...interface{}) () {
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
    C._ZNK19QGraphicsLayoutItem13preferredSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "preferredSize", args)
  }

}

// setParentLayoutItem(class QGraphicsLayoutItem *)
func (this *QGraphicsLayoutItem) setParentLayoutItem(args ...interface{}) () {
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
    C._ZN19QGraphicsLayoutItem19setParentLayoutItemEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setParentLayoutItem", args)
  }

}

// graphicsItem()
func (this *QGraphicsLayoutItem) graphicsItem(args ...interface{}) () {
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
    C._ZNK19QGraphicsLayoutItem12graphicsItemEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "graphicsItem", args)
  }

}

// minimumHeight()
func (this *QGraphicsLayoutItem) minimumHeight(args ...interface{}) () {
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
    C._ZNK19QGraphicsLayoutItem13minimumHeightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "minimumHeight", args)
  }

}

// setMaximumSize(const class QSizeF &)
func (this *QGraphicsLayoutItem) setMaximumSize(args ...interface{}) () {
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
    C._ZN19QGraphicsLayoutItem14setMaximumSizeERK6QSizeF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN19QGraphicsLayoutItem14setMaximumSizeEdd
    // invoke: void setMaximumSize(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C._ZN19QGraphicsLayoutItem14setMaximumSizeEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setMaximumSize", args)
  }

}

// setGeometry(const class QRectF &)
func (this *QGraphicsLayoutItem) setGeometry(args ...interface{}) () {
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
    C._ZN19QGraphicsLayoutItem11setGeometryERK6QRectF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setGeometry", args)
  }

}

// setMinimumHeight(qreal)
func (this *QGraphicsLayoutItem) setMinimumHeight(args ...interface{}) () {
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
    C._ZN19QGraphicsLayoutItem16setMinimumHeightEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setMinimumHeight", args)
  }

}

// ~QGraphicsLayoutItem()
func (this *QGraphicsLayoutItem) FreeQGraphicsLayoutItem(args ...interface{}) () {
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
    C._ZN19QGraphicsLayoutItemD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "~QGraphicsLayoutItem", args)
  }

}

// setPreferredSize(const class QSizeF &)
func (this *QGraphicsLayoutItem) setPreferredSize(args ...interface{}) () {
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
    C._ZN19QGraphicsLayoutItem16setPreferredSizeERK6QSizeF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN19QGraphicsLayoutItem16setPreferredSizeEdd
    // invoke: void setPreferredSize(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C._ZN19QGraphicsLayoutItem16setPreferredSizeEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setPreferredSize", args)
  }

}

// QGraphicsLayoutItem(class QGraphicsLayoutItem *, _Bool)
func NewQGraphicsLayoutItem(args ...interface{}) QGraphicsLayoutItem {
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
    C._ZN19QGraphicsLayoutItemC2EPS_b(qthis, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "QGraphicsLayoutItem", args)
  }

  return QGraphicsLayoutItem{}
}

// parentLayoutItem()
func (this *QGraphicsLayoutItem) parentLayoutItem(args ...interface{}) () {
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
    C._ZNK19QGraphicsLayoutItem16parentLayoutItemEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "parentLayoutItem", args)
  }

}

// setPreferredHeight(qreal)
func (this *QGraphicsLayoutItem) setPreferredHeight(args ...interface{}) () {
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
    C._ZN19QGraphicsLayoutItem18setPreferredHeightEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setPreferredHeight", args)
  }

}

// updateGeometry()
func (this *QGraphicsLayoutItem) updateGeometry(args ...interface{}) () {
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
    C._ZN19QGraphicsLayoutItem14updateGeometryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "updateGeometry", args)
  }

}

// sizePolicy()
func (this *QGraphicsLayoutItem) sizePolicy(args ...interface{}) () {
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
    C._ZNK19QGraphicsLayoutItem10sizePolicyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "sizePolicy", args)
  }

}

// minimumWidth()
func (this *QGraphicsLayoutItem) minimumWidth(args ...interface{}) () {
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
    C._ZNK19QGraphicsLayoutItem12minimumWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "minimumWidth", args)
  }

}

// preferredHeight()
func (this *QGraphicsLayoutItem) preferredHeight(args ...interface{}) () {
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
    C._ZNK19QGraphicsLayoutItem15preferredHeightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "preferredHeight", args)
  }

}

// maximumHeight()
func (this *QGraphicsLayoutItem) maximumHeight(args ...interface{}) () {
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
    C._ZNK19QGraphicsLayoutItem13maximumHeightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "maximumHeight", args)
  }

}

// preferredWidth()
func (this *QGraphicsLayoutItem) preferredWidth(args ...interface{}) () {
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
    C._ZNK19QGraphicsLayoutItem14preferredWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "preferredWidth", args)
  }

}

// setMinimumSize(qreal, qreal)
func (this *QGraphicsLayoutItem) setMinimumSize(args ...interface{}) () {
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
    C._ZN19QGraphicsLayoutItem14setMinimumSizeEdd(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN19QGraphicsLayoutItem14setMinimumSizeERK6QSizeF
    // invoke: void setMinimumSize(const class QSizeF &)
    var arg0 = args[0].(QSizeF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN19QGraphicsLayoutItem14setMinimumSizeERK6QSizeF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setMinimumSize", args)
  }

}

// maximumWidth()
func (this *QGraphicsLayoutItem) maximumWidth(args ...interface{}) () {
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
    C._ZNK19QGraphicsLayoutItem12maximumWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "maximumWidth", args)
  }

}

// setSizePolicy(const class QSizePolicy &)
func (this *QGraphicsLayoutItem) setSizePolicy(args ...interface{}) () {
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
    C._ZN19QGraphicsLayoutItem13setSizePolicyERK11QSizePolicy(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setSizePolicy", args)
  }

}

// setMinimumWidth(qreal)
func (this *QGraphicsLayoutItem) setMinimumWidth(args ...interface{}) () {
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
    C._ZN19QGraphicsLayoutItem15setMinimumWidthEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setMinimumWidth", args)
  }

}

// geometry()
func (this *QGraphicsLayoutItem) geometry(args ...interface{}) () {
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
    C._ZNK19QGraphicsLayoutItem8geometryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "geometry", args)
  }

}

// getContentsMargins(qreal *, qreal *, qreal *, qreal *)
func (this *QGraphicsLayoutItem) getContentsMargins(args ...interface{}) () {
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
    var arg0 = (*C.double)(args[0].(*float64))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.double)(args[1].(*float64))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.double)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.double)(args[3].(*float64))
    if false {fmt.Println(arg3)}
    C._ZNK19QGraphicsLayoutItem18getContentsMarginsEPdS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "getContentsMargins", args)
  }

}

// minimumSize()
func (this *QGraphicsLayoutItem) minimumSize(args ...interface{}) () {
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
    C._ZNK19QGraphicsLayoutItem11minimumSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "minimumSize", args)
  }

}

// contentsRect()
func (this *QGraphicsLayoutItem) contentsRect(args ...interface{}) () {
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
    C._ZNK19QGraphicsLayoutItem12contentsRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "contentsRect", args)
  }

}

// isLayout()
func (this *QGraphicsLayoutItem) isLayout(args ...interface{}) () {
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
    C._ZNK19QGraphicsLayoutItem8isLayoutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "isLayout", args)
  }

}

// setPreferredWidth(qreal)
func (this *QGraphicsLayoutItem) setPreferredWidth(args ...interface{}) () {
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
    C._ZN19QGraphicsLayoutItem17setPreferredWidthEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setPreferredWidth", args)
  }

}

// setMaximumHeight(qreal)
func (this *QGraphicsLayoutItem) setMaximumHeight(args ...interface{}) () {
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
    C._ZN19QGraphicsLayoutItem16setMaximumHeightEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setMaximumHeight", args)
  }

}

// maximumSize()
func (this *QGraphicsLayoutItem) maximumSize(args ...interface{}) () {
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
    C._ZNK19QGraphicsLayoutItem11maximumSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "maximumSize", args)
  }

}

// ownedByLayout()
func (this *QGraphicsLayoutItem) ownedByLayout(args ...interface{}) () {
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
    C._ZNK19QGraphicsLayoutItem13ownedByLayoutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "ownedByLayout", args)
  }

}

// <= body block end

