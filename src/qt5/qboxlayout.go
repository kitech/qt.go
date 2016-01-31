package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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
extern void C_ZNK10QBoxLayout7stretchEi(void* qthis, int32_t arg0); // 4
  // proto:  int QBoxLayout::minimumHeightForWidth(int );
extern void C_ZNK10QBoxLayout21minimumHeightForWidthEi(void* qthis, int32_t arg0); // 4
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
extern void C_ZN10QBoxLayout6takeAtEi(void* qthis, int32_t arg0); // 4
  // proto:  void QBoxLayout::insertItem(int index, QLayoutItem * );
extern void C_ZN10QBoxLayout10insertItemEiP11QLayoutItem(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QBoxLayout::setSpacing(int spacing);
extern void C_ZN10QBoxLayout10setSpacingEi(void* qthis, int32_t arg0); // 4
  // proto:  void QBoxLayout::setGeometry(const QRect & );
extern void C_ZN10QBoxLayout11setGeometryERK5QRect(void* qthis, void* arg0); // 4
  // proto:  int QBoxLayout::spacing();
extern void C_ZNK10QBoxLayout7spacingEv(void* qthis); // 4
  // proto:  void QBoxLayout::~QBoxLayout();
extern void C_ZN10QBoxLayoutD2Ev(void* qthis); // 4
  // proto:  int QBoxLayout::count();
extern void C_ZNK10QBoxLayout5countEv(void* qthis); // 4
  // proto:  void QBoxLayout::addSpacing(int size);
extern void C_ZN10QBoxLayout10addSpacingEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QBoxLayout::sizeHint();
extern void C_ZNK10QBoxLayout8sizeHintEv(void* qthis); // 4
  // proto:  bool QBoxLayout::hasHeightForWidth();
extern void C_ZNK10QBoxLayout17hasHeightForWidthEv(void* qthis); // 4
  // proto:  QLayoutItem * QBoxLayout::itemAt(int );
extern void C_ZNK10QBoxLayout6itemAtEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QBoxLayout::minimumSize();
extern void C_ZNK10QBoxLayout11minimumSizeEv(void* qthis); // 4
  // proto:  int QBoxLayout::heightForWidth(int );
extern void C_ZNK10QBoxLayout14heightForWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QBoxLayout::insertSpacerItem(int index, QSpacerItem * spacerItem);
extern void C_ZN10QBoxLayout16insertSpacerItemEiP11QSpacerItem(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QBoxLayout::addLayout(QLayout * layout, int stretch);
extern void C_ZN10QBoxLayout9addLayoutEP7QLayouti(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  bool QBoxLayout::setStretchFactor(QLayout * l, int stretch);
extern void C_ZN10QBoxLayout16setStretchFactorEP7QLayouti(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  bool QBoxLayout::setStretchFactor(QWidget * w, int stretch);
extern void C_ZN10QBoxLayout16setStretchFactorEP7QWidgeti(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QSize QBoxLayout::maximumSize();
extern void C_ZNK10QBoxLayout11maximumSizeEv(void* qthis); // 4
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
func (this *QHBoxLayout) metaObject(args ...interface{}) () {
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

}

// ~QHBoxLayout()
func (this *QHBoxLayout) FreeQHBoxLayout(args ...interface{}) () {
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

}

// invalidate()
func (this *QBoxLayout) invalidate(args ...interface{}) () {
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

}

// setStretch(int, int)
func (this *QBoxLayout) setStretch(args ...interface{}) () {
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

}

// insertLayout(int, class QLayout *, int)
func (this *QBoxLayout) insertLayout(args ...interface{}) () {
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

}

// insertStretch(int, int)
func (this *QBoxLayout) insertStretch(args ...interface{}) () {
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

}

// addSpacerItem(class QSpacerItem *)
func (this *QBoxLayout) addSpacerItem(args ...interface{}) () {
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

}

// expandingDirections()
func (this *QBoxLayout) expandingDirections(args ...interface{}) () {
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

}

// insertSpacing(int, int)
func (this *QBoxLayout) insertSpacing(args ...interface{}) () {
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

}

// stretch(int)
func (this *QBoxLayout) stretch(args ...interface{}) () {
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
    var ret = C.C_ZNK10QBoxLayout7stretchEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QBoxLayout", "stretch", args)
  }

}

// minimumHeightForWidth(int)
func (this *QBoxLayout) minimumHeightForWidth(args ...interface{}) () {
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
    var ret = C.C_ZNK10QBoxLayout21minimumHeightForWidthEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QBoxLayout", "minimumHeightForWidth", args)
  }

}

// metaObject()
func (this *QBoxLayout) metaObject(args ...interface{}) () {
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

}

// addStretch(int)
func (this *QBoxLayout) addStretch(args ...interface{}) () {
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

}

// addStrut(int)
func (this *QBoxLayout) addStrut(args ...interface{}) () {
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

}

// direction()
func (this *QBoxLayout) direction(args ...interface{}) () {
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

}

// addItem(class QLayoutItem *)
func (this *QBoxLayout) addItem(args ...interface{}) () {
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

}

// takeAt(int)
func (this *QBoxLayout) takeAt(args ...interface{}) () {
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
    var ret = C.C_ZN10QBoxLayout6takeAtEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QBoxLayout", "takeAt", args)
  }

}

// insertItem(int, class QLayoutItem *)
func (this *QBoxLayout) insertItem(args ...interface{}) () {
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

}

// setSpacing(int)
func (this *QBoxLayout) setSpacing(args ...interface{}) () {
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

}

// setGeometry(const class QRect &)
func (this *QBoxLayout) setGeometry(args ...interface{}) () {
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

}

// spacing()
func (this *QBoxLayout) spacing(args ...interface{}) () {
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
    var ret = C.C_ZNK10QBoxLayout7spacingEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QBoxLayout", "spacing", args)
  }

}

// ~QBoxLayout()
func (this *QBoxLayout) FreeQBoxLayout(args ...interface{}) () {
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

}

// count()
func (this *QBoxLayout) count(args ...interface{}) () {
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
    var ret = C.C_ZNK10QBoxLayout5countEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QBoxLayout", "count", args)
  }

}

// addSpacing(int)
func (this *QBoxLayout) addSpacing(args ...interface{}) () {
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

}

// sizeHint()
func (this *QBoxLayout) sizeHint(args ...interface{}) () {
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
    var ret = C.C_ZNK10QBoxLayout8sizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QBoxLayout", "sizeHint", args)
  }

}

// hasHeightForWidth()
func (this *QBoxLayout) hasHeightForWidth(args ...interface{}) () {
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
    var ret = C.C_ZNK10QBoxLayout17hasHeightForWidthEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QBoxLayout", "hasHeightForWidth", args)
  }

}

// itemAt(int)
func (this *QBoxLayout) itemAt(args ...interface{}) () {
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
    var ret = C.C_ZNK10QBoxLayout6itemAtEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QBoxLayout", "itemAt", args)
  }

}

// minimumSize()
func (this *QBoxLayout) minimumSize(args ...interface{}) () {
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
    var ret = C.C_ZNK10QBoxLayout11minimumSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QBoxLayout", "minimumSize", args)
  }

}

// heightForWidth(int)
func (this *QBoxLayout) heightForWidth(args ...interface{}) () {
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
    var ret = C.C_ZNK10QBoxLayout14heightForWidthEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QBoxLayout", "heightForWidth", args)
  }

}

// insertSpacerItem(int, class QSpacerItem *)
func (this *QBoxLayout) insertSpacerItem(args ...interface{}) () {
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

}

// addLayout(class QLayout *, int)
func (this *QBoxLayout) addLayout(args ...interface{}) () {
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

}

// setStretchFactor(class QLayout *, int)
func (this *QBoxLayout) setStretchFactor(args ...interface{}) () {
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
    var ret = C.C_ZN10QBoxLayout16setStretchFactorEP7QLayouti(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN10QBoxLayout16setStretchFactorEP7QWidgeti
    // invoke: bool setStretchFactor(class QWidget *, int)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN10QBoxLayout16setStretchFactorEP7QWidgeti(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QBoxLayout", "setStretchFactor", args)
  }

}

// maximumSize()
func (this *QBoxLayout) maximumSize(args ...interface{}) () {
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
    var ret = C.C_ZNK10QBoxLayout11maximumSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QBoxLayout", "maximumSize", args)
  }

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
func (this *QVBoxLayout) FreeQVBoxLayout(args ...interface{}) () {
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

}

// metaObject()
func (this *QVBoxLayout) metaObject(args ...interface{}) () {
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

}

// <= body block end

