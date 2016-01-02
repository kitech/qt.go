package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
  // proto:  void QGraphicsLayoutItem::setSizePolicy(const QSizePolicy & policy);
extern void _ZN19QGraphicsLayoutItem13setSizePolicyERK11QSizePolicy(void* qthis, void* arg0);
  // proto:  QGraphicsLayoutItem * QGraphicsLayoutItem::parentLayoutItem();
extern void _ZNK19QGraphicsLayoutItem16parentLayoutItemEv(void* qthis);
  // proto:  qreal QGraphicsLayoutItem::minimumWidth();
extern void demth_ZNK19QGraphicsLayoutItem12minimumWidthEv(void* qthis);
  // proto:  QGraphicsItem * QGraphicsLayoutItem::graphicsItem();
extern void _ZNK19QGraphicsLayoutItem12graphicsItemEv(void* qthis);
  // proto:  qreal QGraphicsLayoutItem::preferredWidth();
extern void demth_ZNK19QGraphicsLayoutItem14preferredWidthEv(void* qthis);
  // proto:  bool QGraphicsLayoutItem::ownedByLayout();
extern void _ZNK19QGraphicsLayoutItem13ownedByLayoutEv(void* qthis);
  // proto:  QSizeF QGraphicsLayoutItem::preferredSize();
extern void _ZNK19QGraphicsLayoutItem13preferredSizeEv(void* qthis);
  // proto:  QRectF QGraphicsLayoutItem::geometry();
extern void _ZNK19QGraphicsLayoutItem8geometryEv(void* qthis);
  // proto:  void QGraphicsLayoutItem::QGraphicsLayoutItem(QGraphicsLayoutItem * parent, bool isLayout);
extern void* dector_ZN19QGraphicsLayoutItemC1EPS_b(void* arg0, bool arg1);
extern void _ZN19QGraphicsLayoutItemC1EPS_b(void* qthis, void* arg0, bool arg1);
  // proto:  qreal QGraphicsLayoutItem::minimumHeight();
extern void demth_ZNK19QGraphicsLayoutItem13minimumHeightEv(void* qthis);
  // proto:  qreal QGraphicsLayoutItem::preferredHeight();
extern void demth_ZNK19QGraphicsLayoutItem15preferredHeightEv(void* qthis);
  // proto:  QSizeF QGraphicsLayoutItem::maximumSize();
extern void _ZNK19QGraphicsLayoutItem11maximumSizeEv(void* qthis);
  // proto:  QSizePolicy QGraphicsLayoutItem::sizePolicy();
extern void _ZNK19QGraphicsLayoutItem10sizePolicyEv(void* qthis);
  // proto:  qreal QGraphicsLayoutItem::maximumHeight();
extern void demth_ZNK19QGraphicsLayoutItem13maximumHeightEv(void* qthis);
  // proto:  void QGraphicsLayoutItem::setGeometry(const QRectF & rect);
extern void _ZN19QGraphicsLayoutItem11setGeometryERK6QRectF(void* qthis, void* arg0);
  // proto:  void QGraphicsLayoutItem::setPreferredWidth(qreal width);
extern void _ZN19QGraphicsLayoutItem17setPreferredWidthEd(void* qthis, double arg0);
  // proto:  void QGraphicsLayoutItem::setMaximumSize(const QSizeF & size);
extern void _ZN19QGraphicsLayoutItem14setMaximumSizeERK6QSizeF(void* qthis, void* arg0);
  // proto:  qreal QGraphicsLayoutItem::maximumWidth();
extern void demth_ZNK19QGraphicsLayoutItem12maximumWidthEv(void* qthis);
  // proto:  void QGraphicsLayoutItem::setMinimumSize(qreal w, qreal h);
extern void demth_ZN19QGraphicsLayoutItem14setMinimumSizeEdd(void* qthis, double arg0, double arg1);
  // proto:  void QGraphicsLayoutItem::setMaximumHeight(qreal height);
extern void _ZN19QGraphicsLayoutItem16setMaximumHeightEd(void* qthis, double arg0);
  // proto:  void QGraphicsLayoutItem::setMinimumSize(const QSizeF & size);
extern void _ZN19QGraphicsLayoutItem14setMinimumSizeERK6QSizeF(void* qthis, void* arg0);
  // proto:  void QGraphicsLayoutItem::setPreferredSize(const QSizeF & size);
extern void _ZN19QGraphicsLayoutItem16setPreferredSizeERK6QSizeF(void* qthis, void* arg0);
  // proto:  void QGraphicsLayoutItem::getContentsMargins(qreal * left, qreal * top, qreal * right, qreal * bottom);
extern void _ZNK19QGraphicsLayoutItem18getContentsMarginsEPdS0_S0_S0_(void* qthis, double* arg0, double* arg1, double* arg2, double* arg3);
  // proto:  void QGraphicsLayoutItem::setParentLayoutItem(QGraphicsLayoutItem * parent);
extern void _ZN19QGraphicsLayoutItem19setParentLayoutItemEPS_(void* qthis, void* arg0);
  // proto:  void QGraphicsLayoutItem::setMinimumWidth(qreal width);
extern void _ZN19QGraphicsLayoutItem15setMinimumWidthEd(void* qthis, double arg0);
  // proto:  void QGraphicsLayoutItem::setMaximumWidth(qreal width);
extern void _ZN19QGraphicsLayoutItem15setMaximumWidthEd(void* qthis, double arg0);
  // proto:  void QGraphicsLayoutItem::updateGeometry();
extern void _ZN19QGraphicsLayoutItem14updateGeometryEv(void* qthis);
  // proto:  void QGraphicsLayoutItem::setPreferredHeight(qreal height);
extern void _ZN19QGraphicsLayoutItem18setPreferredHeightEd(void* qthis, double arg0);
  // proto:  QSizeF QGraphicsLayoutItem::minimumSize();
extern void _ZNK19QGraphicsLayoutItem11minimumSizeEv(void* qthis);
  // proto:  QRectF QGraphicsLayoutItem::contentsRect();
extern void _ZNK19QGraphicsLayoutItem12contentsRectEv(void* qthis);
  // proto:  bool QGraphicsLayoutItem::isLayout();
extern void _ZNK19QGraphicsLayoutItem8isLayoutEv(void* qthis);
  // proto:  void QGraphicsLayoutItem::setPreferredSize(qreal w, qreal h);
extern void demth_ZN19QGraphicsLayoutItem16setPreferredSizeEdd(void* qthis, double arg0, double arg1);
  // proto:  void QGraphicsLayoutItem::~QGraphicsLayoutItem();
extern void _ZN19QGraphicsLayoutItemD0Ev(void* qthis);
  // proto:  void QGraphicsLayoutItem::setMinimumHeight(qreal height);
extern void _ZN19QGraphicsLayoutItem16setMinimumHeightEd(void* qthis, double arg0);
  // proto:  void QGraphicsLayoutItem::setMaximumSize(qreal w, qreal h);
extern void demth_ZN19QGraphicsLayoutItem14setMaximumSizeEdd(void* qthis, double arg0, double arg1);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QGraphicsLayoutItem::setSizePolicy(const QSizePolicy & policy);
func (this *QGraphicsLayoutItem) setSizePolicy(args ...interface{}) () {
  // setSizePolicy(const class QSizePolicy &)
  // setSizePolicy(class QSizePolicy::Policy, class QSizePolicy::Policy, class QSizePolicy::ControlType)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSizePolicy{}) // "const QSizePolicy &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "QSizePolicy::Policy"
  vtys[1][1] = qtrt.Int32Ty(false) // "QSizePolicy::Policy"
  vtys[1][2] = qtrt.Int32Ty(false) // "QSizePolicy::ControlType"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItem13setSizePolicyERK11QSizePolicy
  case 1:
    // invoke: _ZN19QGraphicsLayoutItem13setSizePolicyEN11QSizePolicy6PolicyES1_NS0_11ControlTypeE
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setSizePolicy", args)
  }

}

  // proto:  QGraphicsLayoutItem * QGraphicsLayoutItem::parentLayoutItem();
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
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "parentLayoutItem", args)
  }

}

  // proto:  qreal QGraphicsLayoutItem::minimumWidth();
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
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "minimumWidth", args)
  }

}

  // proto:  QGraphicsItem * QGraphicsLayoutItem::graphicsItem();
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
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "graphicsItem", args)
  }

}

  // proto:  qreal QGraphicsLayoutItem::preferredWidth();
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
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "preferredWidth", args)
  }

}

  // proto:  bool QGraphicsLayoutItem::ownedByLayout();
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
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "ownedByLayout", args)
  }

}

  // proto:  QSizeF QGraphicsLayoutItem::preferredSize();
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
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "preferredSize", args)
  }

}

  // proto:  QRectF QGraphicsLayoutItem::geometry();
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
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "geometry", args)
  }

}

  // proto:  void QGraphicsLayoutItem::QGraphicsLayoutItem(QGraphicsLayoutItem * parent, bool isLayout);
func NewQGraphicsLayoutItem(args ...interface{}) QGraphicsLayoutItem {
  return QGraphicsLayoutItem{}
}

  // proto:  qreal QGraphicsLayoutItem::minimumHeight();
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
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "minimumHeight", args)
  }

}

  // proto:  qreal QGraphicsLayoutItem::preferredHeight();
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
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "preferredHeight", args)
  }

}

  // proto:  QSizeF QGraphicsLayoutItem::maximumSize();
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
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "maximumSize", args)
  }

}

  // proto:  QSizePolicy QGraphicsLayoutItem::sizePolicy();
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
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "sizePolicy", args)
  }

}

  // proto:  qreal QGraphicsLayoutItem::maximumHeight();
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
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "maximumHeight", args)
  }

}

  // proto:  void QGraphicsLayoutItem::setGeometry(const QRectF & rect);
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
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setGeometry", args)
  }

}

  // proto:  void QGraphicsLayoutItem::setPreferredWidth(qreal width);
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
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setPreferredWidth", args)
  }

}

  // proto:  void QGraphicsLayoutItem::setMaximumSize(const QSizeF & size);
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
  case 1:
    // invoke: _ZN19QGraphicsLayoutItem14setMaximumSizeEdd
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setMaximumSize", args)
  }

}

  // proto:  qreal QGraphicsLayoutItem::maximumWidth();
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
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "maximumWidth", args)
  }

}

  // proto:  void QGraphicsLayoutItem::setMinimumSize(qreal w, qreal h);
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
  case 1:
    // invoke: _ZN19QGraphicsLayoutItem14setMinimumSizeERK6QSizeF
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setMinimumSize", args)
  }

}

  // proto:  void QGraphicsLayoutItem::setMaximumHeight(qreal height);
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
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setMaximumHeight", args)
  }

}

  // proto:  void QGraphicsLayoutItem::setPreferredSize(const QSizeF & size);
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
  case 1:
    // invoke: _ZN19QGraphicsLayoutItem16setPreferredSizeEdd
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setPreferredSize", args)
  }

}

  // proto:  void QGraphicsLayoutItem::getContentsMargins(qreal * left, qreal * top, qreal * right, qreal * bottom);
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
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "getContentsMargins", args)
  }

}

  // proto:  void QGraphicsLayoutItem::setParentLayoutItem(QGraphicsLayoutItem * parent);
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
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setParentLayoutItem", args)
  }

}

  // proto:  void QGraphicsLayoutItem::setMinimumWidth(qreal width);
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
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setMinimumWidth", args)
  }

}

  // proto:  void QGraphicsLayoutItem::setMaximumWidth(qreal width);
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
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setMaximumWidth", args)
  }

}

  // proto:  void QGraphicsLayoutItem::updateGeometry();
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
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "updateGeometry", args)
  }

}

  // proto:  void QGraphicsLayoutItem::setPreferredHeight(qreal height);
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
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setPreferredHeight", args)
  }

}

  // proto:  QSizeF QGraphicsLayoutItem::minimumSize();
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
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "minimumSize", args)
  }

}

  // proto:  QRectF QGraphicsLayoutItem::contentsRect();
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
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "contentsRect", args)
  }

}

  // proto:  bool QGraphicsLayoutItem::isLayout();
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
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "isLayout", args)
  }

}

  // proto:  void QGraphicsLayoutItem::~QGraphicsLayoutItem();
func (this *QGraphicsLayoutItem) FreeQGraphicsLayoutItem(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "~QGraphicsLayoutItem", args)
  }

}

  // proto:  void QGraphicsLayoutItem::setMinimumHeight(qreal height);
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
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setMinimumHeight", args)
  }

}

// <= body block end

