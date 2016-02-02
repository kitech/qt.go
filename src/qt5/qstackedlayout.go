package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtWidgets/qstackedlayout.h
// dst-file: /src/widgets/qstackedlayout.go
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
  // proto:  int QStackedLayout::addWidget(QWidget * w);
extern int32_t C_ZN14QStackedLayout9addWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QStackedLayout::setCurrentIndex(int index);
extern void C_ZN14QStackedLayout15setCurrentIndexEi(void* qthis, int32_t arg0); // 4
  // proto:  QWidget * QStackedLayout::currentWidget();
extern void* C_ZNK14QStackedLayout13currentWidgetEv(void* qthis); // 4
  // proto:  void QStackedLayout::setGeometry(const QRect & rect);
extern void C_ZN14QStackedLayout11setGeometryERK5QRect(void* qthis, void* arg0); // 4
  // proto:  void QStackedLayout::~QStackedLayout();
extern void C_ZN14QStackedLayoutD2Ev(void* qthis); // 4
  // proto:  void QStackedLayout::setCurrentWidget(QWidget * w);
extern void C_ZN14QStackedLayout16setCurrentWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  QWidget * QStackedLayout::widget(int );
extern void* C_ZNK14QStackedLayout6widgetEi(void* qthis, int32_t arg0); // 4
  // proto:  void QStackedLayout::addItem(QLayoutItem * item);
extern void C_ZN14QStackedLayout7addItemEP11QLayoutItem(void* qthis, void* arg0); // 4
  // proto:  QLayoutItem * QStackedLayout::takeAt(int );
extern void* C_ZN14QStackedLayout6takeAtEi(void* qthis, int32_t arg0); // 4
  // proto:  int QStackedLayout::insertWidget(int index, QWidget * w);
extern int32_t C_ZN14QStackedLayout12insertWidgetEiP7QWidget(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QStackedLayout::QStackedLayout();
extern void* C_ZN14QStackedLayoutC2Ev(); // 3
  // proto:  void QStackedLayout::QStackedLayout(QLayout * parentLayout);
extern void* C_ZN14QStackedLayoutC2EP7QLayout(void* arg0); // 3
  // proto:  void QStackedLayout::QStackedLayout(QWidget * parent);
extern void* C_ZN14QStackedLayoutC2EP7QWidget(void* arg0); // 3
  // proto:  QSize QStackedLayout::sizeHint();
extern void* C_ZNK14QStackedLayout8sizeHintEv(void* qthis); // 4
  // proto:  QStackedLayout::StackingMode QStackedLayout::stackingMode();
extern void C_ZNK14QStackedLayout12stackingModeEv(void* qthis); // 4
  // proto:  int QStackedLayout::count();
extern int32_t C_ZNK14QStackedLayout5countEv(void* qthis); // 4
  // proto:  const QMetaObject * QStackedLayout::metaObject();
extern void C_ZNK14QStackedLayout10metaObjectEv(void* qthis); // 4
  // proto:  bool QStackedLayout::hasHeightForWidth();
extern bool C_ZNK14QStackedLayout17hasHeightForWidthEv(void* qthis); // 4
  // proto:  QLayoutItem * QStackedLayout::itemAt(int );
extern void* C_ZNK14QStackedLayout6itemAtEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QStackedLayout::minimumSize();
extern void* C_ZNK14QStackedLayout11minimumSizeEv(void* qthis); // 4
  // proto:  int QStackedLayout::heightForWidth(int width);
extern int32_t C_ZNK14QStackedLayout14heightForWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  int QStackedLayout::currentIndex();
extern int32_t C_ZNK14QStackedLayout12currentIndexEv(void* qthis); // 4
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

// class sizeof(QStackedLayout)=1
type QStackedLayout struct {
  /*qbase*/ QLayout;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _widgetRemoved QStackedLayout_widgetRemoved_signal;
//  _currentChanged QStackedLayout_currentChanged_signal;
}

// addWidget(class QWidget *)
func (this *QStackedLayout) Addwidget(args ...interface{}) (ret interface{}) {
  // addWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedLayout9addWidgetEP7QWidget
    // invoke: int addWidget(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN14QStackedLayout9addWidgetEP7QWidget(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStackedLayout", "addWidget", args)
  }

  return
}

// setCurrentIndex(int)
func (this *QStackedLayout) Setcurrentindex(args ...interface{}) () {
  // setCurrentIndex(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedLayout15setCurrentIndexEi
    // invoke: void setCurrentIndex(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QStackedLayout15setCurrentIndexEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedLayout", "setCurrentIndex", args)
  }

  return
}

// currentWidget()
func (this *QStackedLayout) Currentwidget(args ...interface{}) (ret interface{}) {
  // currentWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedLayout13currentWidgetEv
    // invoke: QWidget * currentWidget()
    var ret0 = C.C_ZNK14QStackedLayout13currentWidgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStackedLayout", "currentWidget", args)
  }

  return
}

// setGeometry(const class QRect &)
func (this *QStackedLayout) Setgeometry(args ...interface{}) () {
  // setGeometry(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedLayout11setGeometryERK5QRect
    // invoke: void setGeometry(const class QRect &)
    var arg0 = args[0].(*QRect).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QStackedLayout11setGeometryERK5QRect(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedLayout", "setGeometry", args)
  }

  return
}

// ~QStackedLayout()
func (this *QStackedLayout) Freeqstackedlayout(args ...interface{}) () {
  // ~QStackedLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedLayoutD0Ev
    // invoke: void ~QStackedLayout()
    C.C_ZN14QStackedLayoutD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStackedLayout", "~QStackedLayout", args)
  }

  return
}

// setCurrentWidget(class QWidget *)
func (this *QStackedLayout) Setcurrentwidget(args ...interface{}) () {
  // setCurrentWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedLayout16setCurrentWidgetEP7QWidget
    // invoke: void setCurrentWidget(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QStackedLayout16setCurrentWidgetEP7QWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedLayout", "setCurrentWidget", args)
  }

  return
}

// widget(int)
func (this *QStackedLayout) Widget(args ...interface{}) (ret interface{}) {
  // widget(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedLayout6widgetEi
    // invoke: QWidget * widget(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK14QStackedLayout6widgetEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStackedLayout", "widget", args)
  }

  return
}

// addItem(class QLayoutItem *)
func (this *QStackedLayout) Additem(args ...interface{}) () {
  // addItem(class QLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLayoutItem{}) // "QLayoutItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedLayout7addItemEP11QLayoutItem
    // invoke: void addItem(class QLayoutItem *)
    var arg0 = args[0].(*QLayoutItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QStackedLayout7addItemEP11QLayoutItem(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedLayout", "addItem", args)
  }

  return
}

// takeAt(int)
func (this *QStackedLayout) Takeat(args ...interface{}) (ret interface{}) {
  // takeAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedLayout6takeAtEi
    // invoke: QLayoutItem * takeAt(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN14QStackedLayout6takeAtEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLayoutItem{}) // "QLayoutItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStackedLayout", "takeAt", args)
  }

  return
}

// insertWidget(int, class QWidget *)
func (this *QStackedLayout) Insertwidget(args ...interface{}) (ret interface{}) {
  // insertWidget(int, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedLayout12insertWidgetEiP7QWidget
    // invoke: int insertWidget(int, class QWidget *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN14QStackedLayout12insertWidgetEiP7QWidget(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStackedLayout", "insertWidget", args)
  }

  return
}

// QStackedLayout()
func NewQStackedLayout(args ...interface{}) *QStackedLayout {
  // QStackedLayout()
  // QStackedLayout(class QLayout *)
  // QStackedLayout(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QLayout{}) // "QLayout *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedLayoutC1Ev
    // invoke: void QStackedLayout()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QStackedLayoutC2Ev()
    return &QStackedLayout{Qclsinst:qthis}
  case 1:
    // invoke: _ZN14QStackedLayoutC1EP7QLayout
    // invoke: void QStackedLayout(class QLayout *)
    var arg0 = args[0].(*QLayout).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QStackedLayoutC2EP7QLayout(arg0)
    return &QStackedLayout{Qclsinst:qthis}
  case 2:
    // invoke: _ZN14QStackedLayoutC1EP7QWidget
    // invoke: void QStackedLayout(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QStackedLayoutC2EP7QWidget(arg0)
    return &QStackedLayout{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStackedLayout", "QStackedLayout", args)
  }

  return nil // QStackedLayout{Qclsinst:qthis}
}

// sizeHint()
func (this *QStackedLayout) Sizehint(args ...interface{}) (ret interface{}) {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedLayout8sizeHintEv
    // invoke: QSize sizeHint()
    var ret0 = C.C_ZNK14QStackedLayout8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStackedLayout", "sizeHint", args)
  }

  return
}

// stackingMode()
func (this *QStackedLayout) Stackingmode(args ...interface{}) () {
  // stackingMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedLayout12stackingModeEv
    // invoke: QStackedLayout::StackingMode stackingMode()
    C.C_ZNK14QStackedLayout12stackingModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStackedLayout", "stackingMode", args)
  }

  return
}

// count()
func (this *QStackedLayout) Count(args ...interface{}) (ret interface{}) {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedLayout5countEv
    // invoke: int count()
    var ret0 = C.C_ZNK14QStackedLayout5countEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStackedLayout", "count", args)
  }

  return
}

// metaObject()
func (this *QStackedLayout) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedLayout10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK14QStackedLayout10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStackedLayout", "metaObject", args)
  }

  return
}

// hasHeightForWidth()
func (this *QStackedLayout) Hasheightforwidth(args ...interface{}) (ret interface{}) {
  // hasHeightForWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedLayout17hasHeightForWidthEv
    // invoke: bool hasHeightForWidth()
    var ret0 = C.C_ZNK14QStackedLayout17hasHeightForWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStackedLayout", "hasHeightForWidth", args)
  }

  return
}

// itemAt(int)
func (this *QStackedLayout) Itemat(args ...interface{}) (ret interface{}) {
  // itemAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedLayout6itemAtEi
    // invoke: QLayoutItem * itemAt(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK14QStackedLayout6itemAtEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLayoutItem{}) // "QLayoutItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStackedLayout", "itemAt", args)
  }

  return
}

// minimumSize()
func (this *QStackedLayout) Minimumsize(args ...interface{}) (ret interface{}) {
  // minimumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedLayout11minimumSizeEv
    // invoke: QSize minimumSize()
    var ret0 = C.C_ZNK14QStackedLayout11minimumSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStackedLayout", "minimumSize", args)
  }

  return
}

// heightForWidth(int)
func (this *QStackedLayout) Heightforwidth(args ...interface{}) (ret interface{}) {
  // heightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedLayout14heightForWidthEi
    // invoke: int heightForWidth(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK14QStackedLayout14heightForWidthEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStackedLayout", "heightForWidth", args)
  }

  return
}

// currentIndex()
func (this *QStackedLayout) Currentindex(args ...interface{}) (ret interface{}) {
  // currentIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedLayout12currentIndexEv
    // invoke: int currentIndex()
    var ret0 = C.C_ZNK14QStackedLayout12currentIndexEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStackedLayout", "currentIndex", args)
  }

  return
}

// <= body block end

