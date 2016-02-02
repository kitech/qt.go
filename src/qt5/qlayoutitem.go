package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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
extern void C_ZN11QLayoutItem10invalidateEv(void* qthis); // 4
  // proto:  QWidget * QLayoutItem::widget();
extern void* C_ZN11QLayoutItem6widgetEv(void* qthis); // 4
  // proto:  QLayout * QLayoutItem::layout();
extern void* C_ZN11QLayoutItem6layoutEv(void* qthis); // 4
  // proto:  void QLayoutItem::~QLayoutItem();
extern void C_ZN11QLayoutItemD2Ev(void* qthis); // 4
  // proto:  QSpacerItem * QLayoutItem::spacerItem();
extern void* C_ZN11QLayoutItem10spacerItemEv(void* qthis); // 4
  // proto:  QSizePolicy::ControlTypes QLayoutItem::controlTypes();
extern void C_ZNK11QLayoutItem12controlTypesEv(void* qthis); // 4
  // proto:  int QLayoutItem::heightForWidth(int );
extern int32_t C_ZNK11QLayoutItem14heightForWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QLayoutItem::hasHeightForWidth();
extern bool C_ZNK11QLayoutItem17hasHeightForWidthEv(void* qthis); // 4
  // proto:  int QLayoutItem::minimumHeightForWidth(int );
extern int32_t C_ZNK11QLayoutItem21minimumHeightForWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  Qt::Alignment QLayoutItem::alignment();
extern void C_ZNK11QLayoutItem9alignmentEv(void* qthis); // 2
  // proto:  Qt::Orientations QSpacerItem::expandingDirections();
extern void C_ZNK11QSpacerItem19expandingDirectionsEv(void* qthis); // 4
  // proto:  QSizePolicy QSpacerItem::sizePolicy();
extern void* C_ZNK11QSpacerItem10sizePolicyEv(void* qthis); // 2
  // proto:  QSize QSpacerItem::sizeHint();
extern void* C_ZNK11QSpacerItem8sizeHintEv(void* qthis); // 4
  // proto:  QRect QSpacerItem::geometry();
extern void* C_ZNK11QSpacerItem8geometryEv(void* qthis); // 4
  // proto:  void QSpacerItem::setGeometry(const QRect & );
extern void C_ZN11QSpacerItem11setGeometryERK5QRect(void* qthis, void* arg0); // 4
  // proto:  QSize QSpacerItem::minimumSize();
extern void* C_ZNK11QSpacerItem11minimumSizeEv(void* qthis); // 4
  // proto:  QSpacerItem * QSpacerItem::spacerItem();
extern void* C_ZN11QSpacerItem10spacerItemEv(void* qthis); // 4
  // proto:  bool QSpacerItem::isEmpty();
extern bool C_ZNK11QSpacerItem7isEmptyEv(void* qthis); // 4
  // proto:  void QSpacerItem::~QSpacerItem();
extern void C_ZN11QSpacerItemD2Ev(void* qthis); // 4
  // proto:  QSize QSpacerItem::maximumSize();
extern void* C_ZNK11QSpacerItem11maximumSizeEv(void* qthis); // 4
  // proto:  Qt::Orientations QWidgetItem::expandingDirections();
extern void C_ZNK11QWidgetItem19expandingDirectionsEv(void* qthis); // 4
  // proto:  QWidget * QWidgetItem::widget();
extern void* C_ZN11QWidgetItem6widgetEv(void* qthis); // 4
  // proto:  void QWidgetItem::setGeometry(const QRect & );
extern void C_ZN11QWidgetItem11setGeometryERK5QRect(void* qthis, void* arg0); // 4
  // proto:  QSize QWidgetItem::sizeHint();
extern void* C_ZNK11QWidgetItem8sizeHintEv(void* qthis); // 4
  // proto:  bool QWidgetItem::hasHeightForWidth();
extern bool C_ZNK11QWidgetItem17hasHeightForWidthEv(void* qthis); // 4
  // proto:  QRect QWidgetItem::geometry();
extern void* C_ZNK11QWidgetItem8geometryEv(void* qthis); // 4
  // proto:  void QWidgetItem::~QWidgetItem();
extern void C_ZN11QWidgetItemD2Ev(void* qthis); // 4
  // proto:  void QWidgetItem::QWidgetItem(QWidget * w);
extern void* C_ZN11QWidgetItemC2EP7QWidget(void* arg0); // 1
  // proto:  QSize QWidgetItem::minimumSize();
extern void* C_ZNK11QWidgetItem11minimumSizeEv(void* qthis); // 4
  // proto:  QSizePolicy::ControlTypes QWidgetItem::controlTypes();
extern void C_ZNK11QWidgetItem12controlTypesEv(void* qthis); // 4
  // proto:  int QWidgetItem::heightForWidth(int );
extern int32_t C_ZNK11QWidgetItem14heightForWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QWidgetItem::isEmpty();
extern bool C_ZNK11QWidgetItem7isEmptyEv(void* qthis); // 4
  // proto:  QSize QWidgetItem::maximumSize();
extern void* C_ZNK11QWidgetItem11maximumSizeEv(void* qthis); // 4
  // proto:  void QWidgetItemV2::~QWidgetItemV2();
extern void C_ZN13QWidgetItemV2D2Ev(void* qthis); // 4
  // proto:  QSize QWidgetItemV2::sizeHint();
extern void* C_ZNK13QWidgetItemV28sizeHintEv(void* qthis); // 4
  // proto:  void QWidgetItemV2::QWidgetItemV2(QWidget * widget);
extern void* C_ZN13QWidgetItemV2C2EP7QWidget(void* arg0); // 3
  // proto:  QSize QWidgetItemV2::minimumSize();
extern void* C_ZNK13QWidgetItemV211minimumSizeEv(void* qthis); // 4
  // proto:  int QWidgetItemV2::heightForWidth(int width);
extern int32_t C_ZNK13QWidgetItemV214heightForWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QWidgetItemV2::maximumSize();
extern void* C_ZNK13QWidgetItemV211maximumSizeEv(void* qthis); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QSpacerItem)=1
type QSpacerItem struct {
  /*qbase*/ QLayoutItem;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QWidgetItem)=1
type QWidgetItem struct {
  /*qbase*/ QLayoutItem;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QWidgetItemV2)=1
type QWidgetItemV2 struct {
  /*qbase*/ QWidgetItem;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// invalidate()
func (this *QLayoutItem) Invalidate(args ...interface{}) () {
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
    C.C_ZN11QLayoutItem10invalidateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLayoutItem", "invalidate", args)
  }

  return
}

// widget()
func (this *QLayoutItem) Widget(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN11QLayoutItem6widgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLayoutItem", "widget", args)
  }

  return
}

// layout()
func (this *QLayoutItem) Layout(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN11QLayoutItem6layoutEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLayout{}) // "QLayout *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLayoutItem", "layout", args)
  }

  return
}

// ~QLayoutItem()
func (this *QLayoutItem) Freeqlayoutitem(args ...interface{}) () {
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
    C.C_ZN11QLayoutItemD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLayoutItem", "~QLayoutItem", args)
  }

  return
}

// spacerItem()
func (this *QLayoutItem) Spaceritem(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN11QLayoutItem10spacerItemEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSpacerItem{}) // "QSpacerItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLayoutItem", "spacerItem", args)
  }

  return
}

// controlTypes()
func (this *QLayoutItem) Controltypes(args ...interface{}) () {
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
    C.C_ZNK11QLayoutItem12controlTypesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLayoutItem", "controlTypes", args)
  }

  return
}

// heightForWidth(int)
func (this *QLayoutItem) Heightforwidth(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QLayoutItem14heightForWidthEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLayoutItem", "heightForWidth", args)
  }

  return
}

// hasHeightForWidth()
func (this *QLayoutItem) Hasheightforwidth(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QLayoutItem17hasHeightForWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLayoutItem", "hasHeightForWidth", args)
  }

  return
}

// minimumHeightForWidth(int)
func (this *QLayoutItem) Minimumheightforwidth(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QLayoutItem21minimumHeightForWidthEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLayoutItem", "minimumHeightForWidth", args)
  }

  return
}

// alignment()
func (this *QLayoutItem) Alignment(args ...interface{}) () {
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
    C.C_ZNK11QLayoutItem9alignmentEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLayoutItem", "alignment", args)
  }

  return
}

// expandingDirections()
func (this *QSpacerItem) Expandingdirections(args ...interface{}) () {
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
    C.C_ZNK11QSpacerItem19expandingDirectionsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSpacerItem", "expandingDirections", args)
  }

  return
}

// sizePolicy()
func (this *QSpacerItem) Sizepolicy(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QSpacerItem10sizePolicyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSizePolicy{}) // "QSizePolicy"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSpacerItem", "sizePolicy", args)
  }

  return
}

// sizeHint()
func (this *QSpacerItem) Sizehint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QSpacerItem8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSpacerItem", "sizeHint", args)
  }

  return
}

// geometry()
func (this *QSpacerItem) Geometry(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QSpacerItem8geometryEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSpacerItem", "geometry", args)
  }

  return
}

// setGeometry(const class QRect &)
func (this *QSpacerItem) Setgeometry(args ...interface{}) () {
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
    var arg0 = args[0].(*QRect).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QSpacerItem11setGeometryERK5QRect(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSpacerItem", "setGeometry", args)
  }

  return
}

// minimumSize()
func (this *QSpacerItem) Minimumsize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QSpacerItem11minimumSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSpacerItem", "minimumSize", args)
  }

  return
}

// spacerItem()
func (this *QSpacerItem) Spaceritem(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN11QSpacerItem10spacerItemEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSpacerItem{}) // "QSpacerItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSpacerItem", "spacerItem", args)
  }

  return
}

// isEmpty()
func (this *QSpacerItem) Isempty(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QSpacerItem7isEmptyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSpacerItem", "isEmpty", args)
  }

  return
}

// ~QSpacerItem()
func (this *QSpacerItem) Freeqspaceritem(args ...interface{}) () {
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
    C.C_ZN11QSpacerItemD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSpacerItem", "~QSpacerItem", args)
  }

  return
}

// maximumSize()
func (this *QSpacerItem) Maximumsize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QSpacerItem11maximumSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSpacerItem", "maximumSize", args)
  }

  return
}

// expandingDirections()
func (this *QWidgetItem) Expandingdirections(args ...interface{}) () {
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
    C.C_ZNK11QWidgetItem19expandingDirectionsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidgetItem", "expandingDirections", args)
  }

  return
}

// widget()
func (this *QWidgetItem) Widget(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN11QWidgetItem6widgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidgetItem", "widget", args)
  }

  return
}

// setGeometry(const class QRect &)
func (this *QWidgetItem) Setgeometry(args ...interface{}) () {
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
    var arg0 = args[0].(*QRect).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QWidgetItem11setGeometryERK5QRect(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidgetItem", "setGeometry", args)
  }

  return
}

// sizeHint()
func (this *QWidgetItem) Sizehint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QWidgetItem8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidgetItem", "sizeHint", args)
  }

  return
}

// hasHeightForWidth()
func (this *QWidgetItem) Hasheightforwidth(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QWidgetItem17hasHeightForWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidgetItem", "hasHeightForWidth", args)
  }

  return
}

// geometry()
func (this *QWidgetItem) Geometry(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QWidgetItem8geometryEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidgetItem", "geometry", args)
  }

  return
}

// ~QWidgetItem()
func (this *QWidgetItem) Freeqwidgetitem(args ...interface{}) () {
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
    C.C_ZN11QWidgetItemD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidgetItem", "~QWidgetItem", args)
  }

  return
}

// QWidgetItem(class QWidget *)
func NewQWidgetItem(args ...interface{}) *QWidgetItem {
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QWidgetItemC2EP7QWidget(arg0)
    return &QWidgetItem{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QWidgetItem", "QWidgetItem", args)
  }

  return nil // QWidgetItem{Qclsinst:qthis}
}

// minimumSize()
func (this *QWidgetItem) Minimumsize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QWidgetItem11minimumSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidgetItem", "minimumSize", args)
  }

  return
}

// controlTypes()
func (this *QWidgetItem) Controltypes(args ...interface{}) () {
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
    C.C_ZNK11QWidgetItem12controlTypesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidgetItem", "controlTypes", args)
  }

  return
}

// heightForWidth(int)
func (this *QWidgetItem) Heightforwidth(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QWidgetItem14heightForWidthEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidgetItem", "heightForWidth", args)
  }

  return
}

// isEmpty()
func (this *QWidgetItem) Isempty(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QWidgetItem7isEmptyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidgetItem", "isEmpty", args)
  }

  return
}

// maximumSize()
func (this *QWidgetItem) Maximumsize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QWidgetItem11maximumSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidgetItem", "maximumSize", args)
  }

  return
}

// ~QWidgetItemV2()
func (this *QWidgetItemV2) Freeqwidgetitemv2(args ...interface{}) () {
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
    C.C_ZN13QWidgetItemV2D2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidgetItemV2", "~QWidgetItemV2", args)
  }

  return
}

// sizeHint()
func (this *QWidgetItemV2) Sizehint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QWidgetItemV28sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidgetItemV2", "sizeHint", args)
  }

  return
}

// QWidgetItemV2(class QWidget *)
func NewQWidgetItemV2(args ...interface{}) *QWidgetItemV2 {
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QWidgetItemV2C2EP7QWidget(arg0)
    return &QWidgetItemV2{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QWidgetItemV2", "QWidgetItemV2", args)
  }

  return nil // QWidgetItemV2{Qclsinst:qthis}
}

// minimumSize()
func (this *QWidgetItemV2) Minimumsize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QWidgetItemV211minimumSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidgetItemV2", "minimumSize", args)
  }

  return
}

// heightForWidth(int)
func (this *QWidgetItemV2) Heightforwidth(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QWidgetItemV214heightForWidthEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidgetItemV2", "heightForWidth", args)
  }

  return
}

// maximumSize()
func (this *QWidgetItemV2) Maximumsize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QWidgetItemV211maximumSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidgetItemV2", "maximumSize", args)
  }

  return
}

// <= body block end

