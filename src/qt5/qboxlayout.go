package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtWidgets/qboxlayout.h
// dst-file: /src/widgets/qboxlayout.go
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
  // proto:  void QHBoxLayout::QHBoxLayout(QWidget * parent);
extern void* C_ZN11QHBoxLayoutC2EP7QWidget(void* arg0); // 3
  // proto:  void QHBoxLayout::QHBoxLayout();
extern void* C_ZN11QHBoxLayoutC2Ev(); // 3
  // proto:  const QMetaObject * QHBoxLayout::metaObject();
extern void C_ZNK11QHBoxLayout10metaObjectEv(void* qthis); // 4
  // proto:  void QHBoxLayout::~QHBoxLayout();
extern void C_ZN11QHBoxLayoutD2Ev(void* qthis); // 4
  // proto:  void QBoxLayout::invalidate();
extern void C_ZN10QBoxLayout10invalidateEv(void* qthis); // 4
  // proto:  void QBoxLayout::setStretch(int index, int stretch);
extern void C_ZN10QBoxLayout10setStretchEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QBoxLayout::insertLayout(int index, QLayout * layout, int stretch);
extern void C_ZN10QBoxLayout12insertLayoutEiP7QLayouti(void* qthis, int32_t arg0, void* arg1, int32_t arg2); // 4
  // proto:  void QBoxLayout::insertStretch(int index, int stretch);
extern void C_ZN10QBoxLayout13insertStretchEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QBoxLayout::addSpacerItem(QSpacerItem * spacerItem);
extern void C_ZN10QBoxLayout13addSpacerItemEP11QSpacerItem(void* qthis, void* arg0); // 4
  // proto:  Qt::Orientations QBoxLayout::expandingDirections();
extern void C_ZNK10QBoxLayout19expandingDirectionsEv(void* qthis); // 4
  // proto:  void QBoxLayout::insertSpacing(int index, int size);
extern void C_ZN10QBoxLayout13insertSpacingEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  int QBoxLayout::stretch(int index);
extern int32_t C_ZNK10QBoxLayout7stretchEi(void* qthis, int32_t arg0); // 4
  // proto:  int QBoxLayout::minimumHeightForWidth(int );
extern int32_t C_ZNK10QBoxLayout21minimumHeightForWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  const QMetaObject * QBoxLayout::metaObject();
extern void C_ZNK10QBoxLayout10metaObjectEv(void* qthis); // 4
  // proto:  void QBoxLayout::addStretch(int stretch);
extern void C_ZN10QBoxLayout10addStretchEi(void* qthis, int32_t arg0); // 4
  // proto:  void QBoxLayout::addStrut(int );
extern void C_ZN10QBoxLayout8addStrutEi(void* qthis, int32_t arg0); // 4
  // proto:  QBoxLayout::Direction QBoxLayout::direction();
extern void C_ZNK10QBoxLayout9directionEv(void* qthis); // 4
  // proto:  void QBoxLayout::addItem(QLayoutItem * );
extern void C_ZN10QBoxLayout7addItemEP11QLayoutItem(void* qthis, void* arg0); // 4
  // proto:  QLayoutItem * QBoxLayout::takeAt(int );
extern void* C_ZN10QBoxLayout6takeAtEi(void* qthis, int32_t arg0); // 4
  // proto:  void QBoxLayout::insertItem(int index, QLayoutItem * );
extern void C_ZN10QBoxLayout10insertItemEiP11QLayoutItem(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QBoxLayout::setSpacing(int spacing);
extern void C_ZN10QBoxLayout10setSpacingEi(void* qthis, int32_t arg0); // 4
  // proto:  void QBoxLayout::setGeometry(const QRect & );
extern void C_ZN10QBoxLayout11setGeometryERK5QRect(void* qthis, void* arg0); // 4
  // proto:  int QBoxLayout::spacing();
extern int32_t C_ZNK10QBoxLayout7spacingEv(void* qthis); // 4
  // proto:  void QBoxLayout::~QBoxLayout();
extern void C_ZN10QBoxLayoutD2Ev(void* qthis); // 4
  // proto:  int QBoxLayout::count();
extern int32_t C_ZNK10QBoxLayout5countEv(void* qthis); // 4
  // proto:  void QBoxLayout::addSpacing(int size);
extern void C_ZN10QBoxLayout10addSpacingEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QBoxLayout::sizeHint();
extern void* C_ZNK10QBoxLayout8sizeHintEv(void* qthis); // 4
  // proto:  bool QBoxLayout::hasHeightForWidth();
extern bool C_ZNK10QBoxLayout17hasHeightForWidthEv(void* qthis); // 4
  // proto:  QLayoutItem * QBoxLayout::itemAt(int );
extern void* C_ZNK10QBoxLayout6itemAtEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QBoxLayout::minimumSize();
extern void* C_ZNK10QBoxLayout11minimumSizeEv(void* qthis); // 4
  // proto:  int QBoxLayout::heightForWidth(int );
extern int32_t C_ZNK10QBoxLayout14heightForWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QBoxLayout::insertSpacerItem(int index, QSpacerItem * spacerItem);
extern void C_ZN10QBoxLayout16insertSpacerItemEiP11QSpacerItem(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QBoxLayout::addLayout(QLayout * layout, int stretch);
extern void C_ZN10QBoxLayout9addLayoutEP7QLayouti(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  bool QBoxLayout::setStretchFactor(QLayout * l, int stretch);
extern bool C_ZN10QBoxLayout16setStretchFactorEP7QLayouti(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  bool QBoxLayout::setStretchFactor(QWidget * w, int stretch);
extern bool C_ZN10QBoxLayout16setStretchFactorEP7QWidgeti(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QSize QBoxLayout::maximumSize();
extern void* C_ZNK10QBoxLayout11maximumSizeEv(void* qthis); // 4
  // proto:  void QVBoxLayout::QVBoxLayout();
extern void* C_ZN11QVBoxLayoutC2Ev(); // 3
  // proto:  void QVBoxLayout::QVBoxLayout(QWidget * parent);
extern void* C_ZN11QVBoxLayoutC2EP7QWidget(void* arg0); // 3
  // proto:  void QVBoxLayout::~QVBoxLayout();
extern void C_ZN11QVBoxLayoutD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QVBoxLayout::metaObject();
extern void C_ZNK11QVBoxLayout10metaObjectEv(void* qthis); // 4
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

// class sizeof(QHBoxLayout)=1
type QHBoxLayout struct {
  /*qbase*/ QBoxLayout;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QBoxLayout)=1
type QBoxLayout struct {
  /*qbase*/ QLayout;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QVBoxLayout)=1
type QVBoxLayout struct {
  /*qbase*/ QBoxLayout;
  qclsinst unsafe.Pointer /* *C.void */;
}

// QHBoxLayout(class QWidget *)
func NewQHBoxLayout(args ...interface{}) *QHBoxLayout {
  // QHBoxLayout(class QWidget *)
  // QHBoxLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHBoxLayoutC1EP7QWidget
    // invoke: void QHBoxLayout(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QHBoxLayoutC2EP7QWidget(arg0)
    return &QHBoxLayout{qclsinst:qthis}
  case 1:
    // invoke: _ZN11QHBoxLayoutC1Ev
    // invoke: void QHBoxLayout()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QHBoxLayoutC2Ev()
    return &QHBoxLayout{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QHBoxLayout", "QHBoxLayout", args)
  }

  return nil // QHBoxLayout{qclsinst:qthis}
}

// metaObject()
func (this *QHBoxLayout) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHBoxLayout10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK11QHBoxLayout10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHBoxLayout", "metaObject", args)
  }

  return
}

// ~QHBoxLayout()
func (this *QHBoxLayout) Freeqhboxlayout(args ...interface{}) () {
  // ~QHBoxLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHBoxLayoutD0Ev
    // invoke: void ~QHBoxLayout()
    C.C_ZN11QHBoxLayoutD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHBoxLayout", "~QHBoxLayout", args)
  }

  return
}

// invalidate()
func (this *QBoxLayout) Invalidate(args ...interface{}) () {
  // invalidate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QBoxLayout10invalidateEv
    // invoke: void invalidate()
    C.C_ZN10QBoxLayout10invalidateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBoxLayout", "invalidate", args)
  }

  return
}

// setStretch(int, int)
func (this *QBoxLayout) Setstretch(args ...interface{}) () {
  // setStretch(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QBoxLayout10setStretchEii
    // invoke: void setStretch(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN10QBoxLayout10setStretchEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QBoxLayout", "setStretch", args)
  }

  return
}

// insertLayout(int, class QLayout *, int)
func (this *QBoxLayout) Insertlayout(args ...interface{}) () {
  // insertLayout(int, class QLayout *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QLayout{}) // "QLayout *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QBoxLayout12insertLayoutEiP7QLayouti
    // invoke: void insertLayout(int, class QLayout *, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QLayout).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN10QBoxLayout12insertLayoutEiP7QLayouti(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QBoxLayout", "insertLayout", args)
  }

  return
}

// insertStretch(int, int)
func (this *QBoxLayout) Insertstretch(args ...interface{}) () {
  // insertStretch(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QBoxLayout13insertStretchEii
    // invoke: void insertStretch(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN10QBoxLayout13insertStretchEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QBoxLayout", "insertStretch", args)
  }

  return
}

// addSpacerItem(class QSpacerItem *)
func (this *QBoxLayout) Addspaceritem(args ...interface{}) () {
  // addSpacerItem(class QSpacerItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSpacerItem{}) // "QSpacerItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QBoxLayout13addSpacerItemEP11QSpacerItem
    // invoke: void addSpacerItem(class QSpacerItem *)
    var arg0 = args[0].(QSpacerItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QBoxLayout13addSpacerItemEP11QSpacerItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBoxLayout", "addSpacerItem", args)
  }

  return
}

// expandingDirections()
func (this *QBoxLayout) Expandingdirections(args ...interface{}) () {
  // expandingDirections()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QBoxLayout19expandingDirectionsEv
    // invoke: Qt::Orientations expandingDirections()
    C.C_ZNK10QBoxLayout19expandingDirectionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBoxLayout", "expandingDirections", args)
  }

  return
}

// insertSpacing(int, int)
func (this *QBoxLayout) Insertspacing(args ...interface{}) () {
  // insertSpacing(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QBoxLayout13insertSpacingEii
    // invoke: void insertSpacing(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN10QBoxLayout13insertSpacingEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QBoxLayout", "insertSpacing", args)
  }

  return
}

// stretch(int)
func (this *QBoxLayout) Stretch(args ...interface{}) (ret interface{}) {
  // stretch(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QBoxLayout7stretchEi
    // invoke: int stretch(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QBoxLayout7stretchEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QBoxLayout", "stretch", args)
  }

  return
}

// minimumHeightForWidth(int)
func (this *QBoxLayout) Minimumheightforwidth(args ...interface{}) (ret interface{}) {
  // minimumHeightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QBoxLayout21minimumHeightForWidthEi
    // invoke: int minimumHeightForWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QBoxLayout21minimumHeightForWidthEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QBoxLayout", "minimumHeightForWidth", args)
  }

  return
}

// metaObject()
func (this *QBoxLayout) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QBoxLayout10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK10QBoxLayout10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBoxLayout", "metaObject", args)
  }

  return
}

// addStretch(int)
func (this *QBoxLayout) Addstretch(args ...interface{}) () {
  // addStretch(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QBoxLayout10addStretchEi
    // invoke: void addStretch(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QBoxLayout10addStretchEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBoxLayout", "addStretch", args)
  }

  return
}

// addStrut(int)
func (this *QBoxLayout) Addstrut(args ...interface{}) () {
  // addStrut(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QBoxLayout8addStrutEi
    // invoke: void addStrut(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QBoxLayout8addStrutEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBoxLayout", "addStrut", args)
  }

  return
}

// direction()
func (this *QBoxLayout) Direction(args ...interface{}) () {
  // direction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QBoxLayout9directionEv
    // invoke: QBoxLayout::Direction direction()
    C.C_ZNK10QBoxLayout9directionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBoxLayout", "direction", args)
  }

  return
}

// addItem(class QLayoutItem *)
func (this *QBoxLayout) Additem(args ...interface{}) () {
  // addItem(class QLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLayoutItem{}) // "QLayoutItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QBoxLayout7addItemEP11QLayoutItem
    // invoke: void addItem(class QLayoutItem *)
    var arg0 = args[0].(QLayoutItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QBoxLayout7addItemEP11QLayoutItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBoxLayout", "addItem", args)
  }

  return
}

// takeAt(int)
func (this *QBoxLayout) Takeat(args ...interface{}) (ret interface{}) {
  // takeAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QBoxLayout6takeAtEi
    // invoke: QLayoutItem * takeAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN10QBoxLayout6takeAtEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLayoutItem{}) // "QLayoutItem *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QBoxLayout", "takeAt", args)
  }

  return
}

// insertItem(int, class QLayoutItem *)
func (this *QBoxLayout) Insertitem(args ...interface{}) () {
  // insertItem(int, class QLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QLayoutItem{}) // "QLayoutItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QBoxLayout10insertItemEiP11QLayoutItem
    // invoke: void insertItem(int, class QLayoutItem *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QLayoutItem).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN10QBoxLayout10insertItemEiP11QLayoutItem(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QBoxLayout", "insertItem", args)
  }

  return
}

// setSpacing(int)
func (this *QBoxLayout) Setspacing(args ...interface{}) () {
  // setSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QBoxLayout10setSpacingEi
    // invoke: void setSpacing(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QBoxLayout10setSpacingEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBoxLayout", "setSpacing", args)
  }

  return
}

// setGeometry(const class QRect &)
func (this *QBoxLayout) Setgeometry(args ...interface{}) () {
  // setGeometry(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QBoxLayout11setGeometryERK5QRect
    // invoke: void setGeometry(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QBoxLayout11setGeometryERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBoxLayout", "setGeometry", args)
  }

  return
}

// spacing()
func (this *QBoxLayout) Spacing(args ...interface{}) (ret interface{}) {
  // spacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QBoxLayout7spacingEv
    // invoke: int spacing()
    var ret0 = C.C_ZNK10QBoxLayout7spacingEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QBoxLayout", "spacing", args)
  }

  return
}

// ~QBoxLayout()
func (this *QBoxLayout) Freeqboxlayout(args ...interface{}) () {
  // ~QBoxLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QBoxLayoutD0Ev
    // invoke: void ~QBoxLayout()
    C.C_ZN10QBoxLayoutD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBoxLayout", "~QBoxLayout", args)
  }

  return
}

// count()
func (this *QBoxLayout) Count(args ...interface{}) (ret interface{}) {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QBoxLayout5countEv
    // invoke: int count()
    var ret0 = C.C_ZNK10QBoxLayout5countEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QBoxLayout", "count", args)
  }

  return
}

// addSpacing(int)
func (this *QBoxLayout) Addspacing(args ...interface{}) () {
  // addSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QBoxLayout10addSpacingEi
    // invoke: void addSpacing(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QBoxLayout10addSpacingEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBoxLayout", "addSpacing", args)
  }

  return
}

// sizeHint()
func (this *QBoxLayout) Sizehint(args ...interface{}) (ret interface{}) {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QBoxLayout8sizeHintEv
    // invoke: QSize sizeHint()
    var ret0 = C.C_ZNK10QBoxLayout8sizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QBoxLayout", "sizeHint", args)
  }

  return
}

// hasHeightForWidth()
func (this *QBoxLayout) Hasheightforwidth(args ...interface{}) (ret interface{}) {
  // hasHeightForWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QBoxLayout17hasHeightForWidthEv
    // invoke: bool hasHeightForWidth()
    var ret0 = C.C_ZNK10QBoxLayout17hasHeightForWidthEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QBoxLayout", "hasHeightForWidth", args)
  }

  return
}

// itemAt(int)
func (this *QBoxLayout) Itemat(args ...interface{}) (ret interface{}) {
  // itemAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QBoxLayout6itemAtEi
    // invoke: QLayoutItem * itemAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QBoxLayout6itemAtEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLayoutItem{}) // "QLayoutItem *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QBoxLayout", "itemAt", args)
  }

  return
}

// minimumSize()
func (this *QBoxLayout) Minimumsize(args ...interface{}) (ret interface{}) {
  // minimumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QBoxLayout11minimumSizeEv
    // invoke: QSize minimumSize()
    var ret0 = C.C_ZNK10QBoxLayout11minimumSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QBoxLayout", "minimumSize", args)
  }

  return
}

// heightForWidth(int)
func (this *QBoxLayout) Heightforwidth(args ...interface{}) (ret interface{}) {
  // heightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QBoxLayout14heightForWidthEi
    // invoke: int heightForWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QBoxLayout14heightForWidthEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QBoxLayout", "heightForWidth", args)
  }

  return
}

// insertSpacerItem(int, class QSpacerItem *)
func (this *QBoxLayout) Insertspaceritem(args ...interface{}) () {
  // insertSpacerItem(int, class QSpacerItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QSpacerItem{}) // "QSpacerItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QBoxLayout16insertSpacerItemEiP11QSpacerItem
    // invoke: void insertSpacerItem(int, class QSpacerItem *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QSpacerItem).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN10QBoxLayout16insertSpacerItemEiP11QSpacerItem(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QBoxLayout", "insertSpacerItem", args)
  }

  return
}

// addLayout(class QLayout *, int)
func (this *QBoxLayout) Addlayout(args ...interface{}) () {
  // addLayout(class QLayout *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLayout{}) // "QLayout *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QBoxLayout9addLayoutEP7QLayouti
    // invoke: void addLayout(class QLayout *, int)
    var arg0 = args[0].(QLayout).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN10QBoxLayout9addLayoutEP7QLayouti(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QBoxLayout", "addLayout", args)
  }

  return
}

// setStretchFactor(class QLayout *, int)
func (this *QBoxLayout) Setstretchfactor(args ...interface{}) (ret interface{}) {
  // setStretchFactor(class QLayout *, int)
  // setStretchFactor(class QWidget *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLayout{}) // "QLayout *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QBoxLayout16setStretchFactorEP7QLayouti
    // invoke: bool setStretchFactor(class QLayout *, int)
    var arg0 = args[0].(QLayout).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QBoxLayout16setStretchFactorEP7QLayouti(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZN10QBoxLayout16setStretchFactorEP7QWidgeti
    // invoke: bool setStretchFactor(class QWidget *, int)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QBoxLayout16setStretchFactorEP7QWidgeti(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QBoxLayout", "setStretchFactor", args)
  }

  return
}

// maximumSize()
func (this *QBoxLayout) Maximumsize(args ...interface{}) (ret interface{}) {
  // maximumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QBoxLayout11maximumSizeEv
    // invoke: QSize maximumSize()
    var ret0 = C.C_ZNK10QBoxLayout11maximumSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QBoxLayout", "maximumSize", args)
  }

  return
}

// QVBoxLayout()
func NewQVBoxLayout(args ...interface{}) *QVBoxLayout {
  // QVBoxLayout()
  // QVBoxLayout(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QVBoxLayoutC1Ev
    // invoke: void QVBoxLayout()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QVBoxLayoutC2Ev()
    return &QVBoxLayout{qclsinst:qthis}
  case 1:
    // invoke: _ZN11QVBoxLayoutC1EP7QWidget
    // invoke: void QVBoxLayout(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QVBoxLayoutC2EP7QWidget(arg0)
    return &QVBoxLayout{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QVBoxLayout", "QVBoxLayout", args)
  }

  return nil // QVBoxLayout{qclsinst:qthis}
}

// ~QVBoxLayout()
func (this *QVBoxLayout) Freeqvboxlayout(args ...interface{}) () {
  // ~QVBoxLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QVBoxLayoutD0Ev
    // invoke: void ~QVBoxLayout()
    C.C_ZN11QVBoxLayoutD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVBoxLayout", "~QVBoxLayout", args)
  }

  return
}

// metaObject()
func (this *QVBoxLayout) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QVBoxLayout10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK11QVBoxLayout10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVBoxLayout", "metaObject", args)
  }

  return
}

// <= body block end

