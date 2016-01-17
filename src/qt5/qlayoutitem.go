package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtWidgets/qlayoutitem.h
// dst-file: /src/widgets/qlayoutitem.go
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
  // proto:  void QLayoutItem::invalidate();
extern void _ZN11QLayoutItem10invalidateEv(void* qthis); // 4
  // proto:  QWidget * QLayoutItem::widget();
extern void _ZN11QLayoutItem6widgetEv(void* qthis); // 4
  // proto:  QLayout * QLayoutItem::layout();
extern void _ZN11QLayoutItem6layoutEv(void* qthis); // 4
  // proto:  void QLayoutItem::~QLayoutItem();
extern void _ZN11QLayoutItemD2Ev(void* qthis); // 4
  // proto:  QSpacerItem * QLayoutItem::spacerItem();
extern void _ZN11QLayoutItem10spacerItemEv(void* qthis); // 4
  // proto:  QSizePolicy::ControlTypes QLayoutItem::controlTypes();
extern void _ZNK11QLayoutItem12controlTypesEv(void* qthis); // 4
  // proto:  int QLayoutItem::heightForWidth(int );
extern void _ZNK11QLayoutItem14heightForWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QLayoutItem::hasHeightForWidth();
extern void _ZNK11QLayoutItem17hasHeightForWidthEv(void* qthis); // 4
  // proto:  int QLayoutItem::minimumHeightForWidth(int );
extern void _ZNK11QLayoutItem21minimumHeightForWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  Qt::Alignment QLayoutItem::alignment();
extern void _ZNK11QLayoutItem9alignmentEv(void* qthis); // 2
  // proto:  Qt::Orientations QSpacerItem::expandingDirections();
extern void _ZNK11QSpacerItem19expandingDirectionsEv(void* qthis); // 4
  // proto:  QSizePolicy QSpacerItem::sizePolicy();
extern void _ZNK11QSpacerItem10sizePolicyEv(void* qthis); // 2
  // proto:  QSize QSpacerItem::sizeHint();
extern void _ZNK11QSpacerItem8sizeHintEv(void* qthis); // 4
  // proto:  QRect QSpacerItem::geometry();
extern void _ZNK11QSpacerItem8geometryEv(void* qthis); // 4
  // proto:  void QSpacerItem::setGeometry(const QRect & );
extern void _ZN11QSpacerItem11setGeometryERK5QRect(void* qthis, void* arg0); // 4
  // proto:  QSize QSpacerItem::minimumSize();
extern void _ZNK11QSpacerItem11minimumSizeEv(void* qthis); // 4
  // proto:  QSpacerItem * QSpacerItem::spacerItem();
extern void _ZN11QSpacerItem10spacerItemEv(void* qthis); // 4
  // proto:  bool QSpacerItem::isEmpty();
extern void _ZNK11QSpacerItem7isEmptyEv(void* qthis); // 4
  // proto:  void QSpacerItem::~QSpacerItem();
extern void _ZN11QSpacerItemD2Ev(void* qthis); // 4
  // proto:  QSize QSpacerItem::maximumSize();
extern void _ZNK11QSpacerItem11maximumSizeEv(void* qthis); // 4
  // proto:  Qt::Orientations QWidgetItem::expandingDirections();
extern void _ZNK11QWidgetItem19expandingDirectionsEv(void* qthis); // 4
  // proto:  QWidget * QWidgetItem::widget();
extern void _ZN11QWidgetItem6widgetEv(void* qthis); // 4
  // proto:  void QWidgetItem::setGeometry(const QRect & );
extern void _ZN11QWidgetItem11setGeometryERK5QRect(void* qthis, void* arg0); // 4
  // proto:  QSize QWidgetItem::sizeHint();
extern void _ZNK11QWidgetItem8sizeHintEv(void* qthis); // 4
  // proto:  bool QWidgetItem::hasHeightForWidth();
extern void _ZNK11QWidgetItem17hasHeightForWidthEv(void* qthis); // 4
  // proto:  QRect QWidgetItem::geometry();
extern void _ZNK11QWidgetItem8geometryEv(void* qthis); // 4
  // proto:  void QWidgetItem::~QWidgetItem();
extern void _ZN11QWidgetItemD2Ev(void* qthis); // 4
  // proto:  void QWidgetItem::QWidgetItem(QWidget * w);
extern void _ZN11QWidgetItemC2EP7QWidget(void* qthis, void* arg0); // 1
  // proto:  QSize QWidgetItem::minimumSize();
extern void _ZNK11QWidgetItem11minimumSizeEv(void* qthis); // 4
  // proto:  QSizePolicy::ControlTypes QWidgetItem::controlTypes();
extern void _ZNK11QWidgetItem12controlTypesEv(void* qthis); // 4
  // proto:  int QWidgetItem::heightForWidth(int );
extern void _ZNK11QWidgetItem14heightForWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QWidgetItem::isEmpty();
extern void _ZNK11QWidgetItem7isEmptyEv(void* qthis); // 4
  // proto:  QSize QWidgetItem::maximumSize();
extern void _ZNK11QWidgetItem11maximumSizeEv(void* qthis); // 4
  // proto:  void QWidgetItemV2::~QWidgetItemV2();
extern void _ZN13QWidgetItemV2D2Ev(void* qthis); // 4
  // proto:  QSize QWidgetItemV2::sizeHint();
extern void _ZNK13QWidgetItemV28sizeHintEv(void* qthis); // 4
  // proto:  void QWidgetItemV2::QWidgetItemV2(QWidget * widget);
extern void _ZN13QWidgetItemV2C2EP7QWidget(void* qthis, void* arg0); // 3
  // proto:  QSize QWidgetItemV2::minimumSize();
extern void _ZNK13QWidgetItemV211minimumSizeEv(void* qthis); // 4
  // proto:  int QWidgetItemV2::heightForWidth(int width);
extern void _ZNK13QWidgetItemV214heightForWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QWidgetItemV2::maximumSize();
extern void _ZNK13QWidgetItemV211maximumSizeEv(void* qthis); // 4
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

// class sizeof(QLayoutItem)=1
type QLayoutItem struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QSpacerItem)=1
type QSpacerItem struct {
  /*qbase*/ QLayoutItem;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QWidgetItem)=1
type QWidgetItem struct {
  /*qbase*/ QLayoutItem;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QWidgetItemV2)=1
type QWidgetItemV2 struct {
  /*qbase*/ QWidgetItem;
  qclsinst unsafe.Pointer /* *C.void */;
}

// invalidate()
func (this *QLayoutItem) invalidate(args ...interface{}) () {
  // invalidate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QLayoutItem10invalidateEv
    // invoke: void invalidate()
    C._ZN11QLayoutItem10invalidateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayoutItem", "invalidate", args)
  }

}

// widget()
func (this *QLayoutItem) widget(args ...interface{}) () {
  // widget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QLayoutItem6widgetEv
    // invoke: QWidget * widget()
    C._ZN11QLayoutItem6widgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayoutItem", "widget", args)
  }

}

// layout()
func (this *QLayoutItem) layout(args ...interface{}) () {
  // layout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QLayoutItem6layoutEv
    // invoke: QLayout * layout()
    C._ZN11QLayoutItem6layoutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayoutItem", "layout", args)
  }

}

// ~QLayoutItem()
func (this *QLayoutItem) FreeQLayoutItem(args ...interface{}) () {
  // ~QLayoutItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QLayoutItemD0Ev
    // invoke: void ~QLayoutItem()
    C._ZN11QLayoutItemD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayoutItem", "~QLayoutItem", args)
  }

}

// spacerItem()
func (this *QLayoutItem) spacerItem(args ...interface{}) () {
  // spacerItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QLayoutItem10spacerItemEv
    // invoke: QSpacerItem * spacerItem()
    C._ZN11QLayoutItem10spacerItemEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayoutItem", "spacerItem", args)
  }

}

// controlTypes()
func (this *QLayoutItem) controlTypes(args ...interface{}) () {
  // controlTypes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QLayoutItem12controlTypesEv
    // invoke: QSizePolicy::ControlTypes controlTypes()
    C._ZNK11QLayoutItem12controlTypesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayoutItem", "controlTypes", args)
  }

}

// heightForWidth(int)
func (this *QLayoutItem) heightForWidth(args ...interface{}) () {
  // heightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QLayoutItem14heightForWidthEi
    // invoke: int heightForWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK11QLayoutItem14heightForWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLayoutItem", "heightForWidth", args)
  }

}

// hasHeightForWidth()
func (this *QLayoutItem) hasHeightForWidth(args ...interface{}) () {
  // hasHeightForWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QLayoutItem17hasHeightForWidthEv
    // invoke: bool hasHeightForWidth()
    C._ZNK11QLayoutItem17hasHeightForWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayoutItem", "hasHeightForWidth", args)
  }

}

// minimumHeightForWidth(int)
func (this *QLayoutItem) minimumHeightForWidth(args ...interface{}) () {
  // minimumHeightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QLayoutItem21minimumHeightForWidthEi
    // invoke: int minimumHeightForWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK11QLayoutItem21minimumHeightForWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLayoutItem", "minimumHeightForWidth", args)
  }

}

// alignment()
func (this *QLayoutItem) alignment(args ...interface{}) () {
  // alignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QLayoutItem9alignmentEv
    // invoke: Qt::Alignment alignment()
    C._ZNK11QLayoutItem9alignmentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayoutItem", "alignment", args)
  }

}

// expandingDirections()
func (this *QSpacerItem) expandingDirections(args ...interface{}) () {
  // expandingDirections()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QSpacerItem19expandingDirectionsEv
    // invoke: Qt::Orientations expandingDirections()
    C._ZNK11QSpacerItem19expandingDirectionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSpacerItem", "expandingDirections", args)
  }

}

// sizePolicy()
func (this *QSpacerItem) sizePolicy(args ...interface{}) () {
  // sizePolicy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QSpacerItem10sizePolicyEv
    // invoke: QSizePolicy sizePolicy()
    C._ZNK11QSpacerItem10sizePolicyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSpacerItem", "sizePolicy", args)
  }

}

// sizeHint()
func (this *QSpacerItem) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QSpacerItem8sizeHintEv
    // invoke: QSize sizeHint()
    C._ZNK11QSpacerItem8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSpacerItem", "sizeHint", args)
  }

}

// geometry()
func (this *QSpacerItem) geometry(args ...interface{}) () {
  // geometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QSpacerItem8geometryEv
    // invoke: QRect geometry()
    C._ZNK11QSpacerItem8geometryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSpacerItem", "geometry", args)
  }

}

// setGeometry(const class QRect &)
func (this *QSpacerItem) setGeometry(args ...interface{}) () {
  // setGeometry(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QSpacerItem11setGeometryERK5QRect
    // invoke: void setGeometry(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QSpacerItem11setGeometryERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSpacerItem", "setGeometry", args)
  }

}

// minimumSize()
func (this *QSpacerItem) minimumSize(args ...interface{}) () {
  // minimumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QSpacerItem11minimumSizeEv
    // invoke: QSize minimumSize()
    C._ZNK11QSpacerItem11minimumSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSpacerItem", "minimumSize", args)
  }

}

// spacerItem()
func (this *QSpacerItem) spacerItem(args ...interface{}) () {
  // spacerItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QSpacerItem10spacerItemEv
    // invoke: QSpacerItem * spacerItem()
    C._ZN11QSpacerItem10spacerItemEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSpacerItem", "spacerItem", args)
  }

}

// isEmpty()
func (this *QSpacerItem) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QSpacerItem7isEmptyEv
    // invoke: bool isEmpty()
    C._ZNK11QSpacerItem7isEmptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSpacerItem", "isEmpty", args)
  }

}

// ~QSpacerItem()
func (this *QSpacerItem) FreeQSpacerItem(args ...interface{}) () {
  // ~QSpacerItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QSpacerItemD0Ev
    // invoke: void ~QSpacerItem()
    C._ZN11QSpacerItemD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSpacerItem", "~QSpacerItem", args)
  }

}

// maximumSize()
func (this *QSpacerItem) maximumSize(args ...interface{}) () {
  // maximumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QSpacerItem11maximumSizeEv
    // invoke: QSize maximumSize()
    C._ZNK11QSpacerItem11maximumSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSpacerItem", "maximumSize", args)
  }

}

// expandingDirections()
func (this *QWidgetItem) expandingDirections(args ...interface{}) () {
  // expandingDirections()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWidgetItem19expandingDirectionsEv
    // invoke: Qt::Orientations expandingDirections()
    C._ZNK11QWidgetItem19expandingDirectionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidgetItem", "expandingDirections", args)
  }

}

// widget()
func (this *QWidgetItem) widget(args ...interface{}) () {
  // widget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QWidgetItem6widgetEv
    // invoke: QWidget * widget()
    C._ZN11QWidgetItem6widgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidgetItem", "widget", args)
  }

}

// setGeometry(const class QRect &)
func (this *QWidgetItem) setGeometry(args ...interface{}) () {
  // setGeometry(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QWidgetItem11setGeometryERK5QRect
    // invoke: void setGeometry(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QWidgetItem11setGeometryERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidgetItem", "setGeometry", args)
  }

}

// sizeHint()
func (this *QWidgetItem) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWidgetItem8sizeHintEv
    // invoke: QSize sizeHint()
    C._ZNK11QWidgetItem8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidgetItem", "sizeHint", args)
  }

}

// hasHeightForWidth()
func (this *QWidgetItem) hasHeightForWidth(args ...interface{}) () {
  // hasHeightForWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWidgetItem17hasHeightForWidthEv
    // invoke: bool hasHeightForWidth()
    C._ZNK11QWidgetItem17hasHeightForWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidgetItem", "hasHeightForWidth", args)
  }

}

// geometry()
func (this *QWidgetItem) geometry(args ...interface{}) () {
  // geometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWidgetItem8geometryEv
    // invoke: QRect geometry()
    C._ZNK11QWidgetItem8geometryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidgetItem", "geometry", args)
  }

}

// ~QWidgetItem()
func (this *QWidgetItem) FreeQWidgetItem(args ...interface{}) () {
  // ~QWidgetItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QWidgetItemD0Ev
    // invoke: void ~QWidgetItem()
    C._ZN11QWidgetItemD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidgetItem", "~QWidgetItem", args)
  }

}

// QWidgetItem(class QWidget *)
func NewQWidgetItem(args ...interface{}) QWidgetItem {
  // QWidgetItem(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QWidgetItemC1EP7QWidget
    // invoke: void QWidgetItem(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN11QWidgetItemC2EP7QWidget(qthis, arg0)
  default:
    qtrt.ErrorResolve("QWidgetItem", "QWidgetItem", args)
  }

  return QWidgetItem{}
}

// minimumSize()
func (this *QWidgetItem) minimumSize(args ...interface{}) () {
  // minimumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWidgetItem11minimumSizeEv
    // invoke: QSize minimumSize()
    C._ZNK11QWidgetItem11minimumSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidgetItem", "minimumSize", args)
  }

}

// controlTypes()
func (this *QWidgetItem) controlTypes(args ...interface{}) () {
  // controlTypes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWidgetItem12controlTypesEv
    // invoke: QSizePolicy::ControlTypes controlTypes()
    C._ZNK11QWidgetItem12controlTypesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidgetItem", "controlTypes", args)
  }

}

// heightForWidth(int)
func (this *QWidgetItem) heightForWidth(args ...interface{}) () {
  // heightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWidgetItem14heightForWidthEi
    // invoke: int heightForWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK11QWidgetItem14heightForWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidgetItem", "heightForWidth", args)
  }

}

// isEmpty()
func (this *QWidgetItem) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWidgetItem7isEmptyEv
    // invoke: bool isEmpty()
    C._ZNK11QWidgetItem7isEmptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidgetItem", "isEmpty", args)
  }

}

// maximumSize()
func (this *QWidgetItem) maximumSize(args ...interface{}) () {
  // maximumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWidgetItem11maximumSizeEv
    // invoke: QSize maximumSize()
    C._ZNK11QWidgetItem11maximumSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidgetItem", "maximumSize", args)
  }

}

// ~QWidgetItemV2()
func (this *QWidgetItemV2) FreeQWidgetItemV2(args ...interface{}) () {
  // ~QWidgetItemV2()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QWidgetItemV2D0Ev
    // invoke: void ~QWidgetItemV2()
    C._ZN13QWidgetItemV2D2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidgetItemV2", "~QWidgetItemV2", args)
  }

}

// sizeHint()
func (this *QWidgetItemV2) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QWidgetItemV28sizeHintEv
    // invoke: QSize sizeHint()
    C._ZNK13QWidgetItemV28sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidgetItemV2", "sizeHint", args)
  }

}

// QWidgetItemV2(class QWidget *)
func NewQWidgetItemV2(args ...interface{}) QWidgetItemV2 {
  // QWidgetItemV2(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QWidgetItemV2C1EP7QWidget
    // invoke: void QWidgetItemV2(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN13QWidgetItemV2C2EP7QWidget(qthis, arg0)
  default:
    qtrt.ErrorResolve("QWidgetItemV2", "QWidgetItemV2", args)
  }

  return QWidgetItemV2{}
}

// minimumSize()
func (this *QWidgetItemV2) minimumSize(args ...interface{}) () {
  // minimumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QWidgetItemV211minimumSizeEv
    // invoke: QSize minimumSize()
    C._ZNK13QWidgetItemV211minimumSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidgetItemV2", "minimumSize", args)
  }

}

// heightForWidth(int)
func (this *QWidgetItemV2) heightForWidth(args ...interface{}) () {
  // heightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QWidgetItemV214heightForWidthEi
    // invoke: int heightForWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK13QWidgetItemV214heightForWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidgetItemV2", "heightForWidth", args)
  }

}

// maximumSize()
func (this *QWidgetItemV2) maximumSize(args ...interface{}) () {
  // maximumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QWidgetItemV211maximumSizeEv
    // invoke: QSize maximumSize()
    C._ZNK13QWidgetItemV211maximumSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidgetItemV2", "maximumSize", args)
  }

}

// <= body block end

