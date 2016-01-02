package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
  // proto:  qreal QGraphicsLinearLayout::spacing();
extern void _ZNK21QGraphicsLinearLayout7spacingEv(void* qthis);
  // proto:  void QGraphicsLinearLayout::QGraphicsLinearLayout(QGraphicsLayoutItem * parent);
extern void* dector_ZN21QGraphicsLinearLayoutC1EP19QGraphicsLayoutItem(void* arg0);
extern void _ZN21QGraphicsLinearLayoutC1EP19QGraphicsLayoutItem(void* qthis, void* arg0);
  // proto:  void QGraphicsLinearLayout::QGraphicsLinearLayout(const QGraphicsLinearLayout & );
extern void* dector_ZN21QGraphicsLinearLayoutC1ERKS_(void* arg0);
extern void _ZN21QGraphicsLinearLayoutC1ERKS_(void* qthis, void* arg0);
  // proto:  QGraphicsLayoutItem * QGraphicsLinearLayout::itemAt(int index);
extern void _ZNK21QGraphicsLinearLayout6itemAtEi(void* qthis, int arg0);
  // proto:  void QGraphicsLinearLayout::invalidate();
extern void _ZN21QGraphicsLinearLayout10invalidateEv(void* qthis);
  // proto:  void QGraphicsLinearLayout::setGeometry(const QRectF & rect);
extern void _ZN21QGraphicsLinearLayout11setGeometryERK6QRectF(void* qthis, void* arg0);
  // proto:  void QGraphicsLinearLayout::addStretch(int stretch);
extern void demth_ZN21QGraphicsLinearLayout10addStretchEi(void* qthis, int arg0);
  // proto:  int QGraphicsLinearLayout::count();
extern void _ZNK21QGraphicsLinearLayout5countEv(void* qthis);
  // proto:  void QGraphicsLinearLayout::setSpacing(qreal spacing);
extern void _ZN21QGraphicsLinearLayout10setSpacingEd(void* qthis, double arg0);
  // proto:  void QGraphicsLinearLayout::insertItem(int index, QGraphicsLayoutItem * item);
extern void _ZN21QGraphicsLinearLayout10insertItemEiP19QGraphicsLayoutItem(void* qthis, int arg0, void* arg1);
  // proto:  void QGraphicsLinearLayout::~QGraphicsLinearLayout();
extern void _ZN21QGraphicsLinearLayoutD0Ev(void* qthis);
  // proto:  void QGraphicsLinearLayout::dump(int indent);
extern void _ZNK21QGraphicsLinearLayout4dumpEi(void* qthis, int arg0);
  // proto:  void QGraphicsLinearLayout::setStretchFactor(QGraphicsLayoutItem * item, int stretch);
extern void _ZN21QGraphicsLinearLayout16setStretchFactorEP19QGraphicsLayoutItemi(void* qthis, void* arg0, int arg1);
  // proto:  void QGraphicsLinearLayout::addItem(QGraphicsLayoutItem * item);
extern void demth_ZN21QGraphicsLinearLayout7addItemEP19QGraphicsLayoutItem(void* qthis, void* arg0);
  // proto:  qreal QGraphicsLinearLayout::itemSpacing(int index);
extern void _ZNK21QGraphicsLinearLayout11itemSpacingEi(void* qthis, int arg0);
  // proto:  void QGraphicsLinearLayout::removeAt(int index);
extern void _ZN21QGraphicsLinearLayout8removeAtEi(void* qthis, int arg0);
  // proto:  void QGraphicsLinearLayout::insertStretch(int index, int stretch);
extern void _ZN21QGraphicsLinearLayout13insertStretchEii(void* qthis, int arg0, int arg1);
  // proto:  void QGraphicsLinearLayout::setItemSpacing(int index, qreal spacing);
extern void _ZN21QGraphicsLinearLayout14setItemSpacingEid(void* qthis, int arg0, double arg1);
  // proto:  void QGraphicsLinearLayout::removeItem(QGraphicsLayoutItem * item);
extern void _ZN21QGraphicsLinearLayout10removeItemEP19QGraphicsLayoutItem(void* qthis, void* arg0);
  // proto:  int QGraphicsLinearLayout::stretchFactor(QGraphicsLayoutItem * item);
extern void _ZNK21QGraphicsLinearLayout13stretchFactorEP19QGraphicsLayoutItem(void* qthis, void* arg0);
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

// class sizeof(QGraphicsLinearLayout)=1
type QGraphicsLinearLayout struct {
  /*qbase*/ QGraphicsLayout;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  qreal QGraphicsLinearLayout::spacing();
func (this *QGraphicsLinearLayout) spacing(args ...interface{}) () {
  // spacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QGraphicsLinearLayout7spacingEv
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "spacing", args)
  }

}

  // proto:  void QGraphicsLinearLayout::QGraphicsLinearLayout(QGraphicsLayoutItem * parent);
func NewQGraphicsLinearLayout(args ...interface{}) QGraphicsLinearLayout {
  return QGraphicsLinearLayout{}
}

  // proto:  QGraphicsLayoutItem * QGraphicsLinearLayout::itemAt(int index);
func (this *QGraphicsLinearLayout) itemAt(args ...interface{}) () {
  // itemAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QGraphicsLinearLayout6itemAtEi
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "itemAt", args)
  }

}

  // proto:  void QGraphicsLinearLayout::invalidate();
func (this *QGraphicsLinearLayout) invalidate(args ...interface{}) () {
  // invalidate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout10invalidateEv
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "invalidate", args)
  }

}

  // proto:  void QGraphicsLinearLayout::setGeometry(const QRectF & rect);
func (this *QGraphicsLinearLayout) setGeometry(args ...interface{}) () {
  // setGeometry(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout11setGeometryERK6QRectF
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "setGeometry", args)
  }

}

  // proto:  void QGraphicsLinearLayout::addStretch(int stretch);
func (this *QGraphicsLinearLayout) addStretch(args ...interface{}) () {
  // addStretch(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout10addStretchEi
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "addStretch", args)
  }

}

  // proto:  int QGraphicsLinearLayout::count();
func (this *QGraphicsLinearLayout) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QGraphicsLinearLayout5countEv
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "count", args)
  }

}

  // proto:  void QGraphicsLinearLayout::setSpacing(qreal spacing);
func (this *QGraphicsLinearLayout) setSpacing(args ...interface{}) () {
  // setSpacing(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout10setSpacingEd
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "setSpacing", args)
  }

}

  // proto:  void QGraphicsLinearLayout::insertItem(int index, QGraphicsLayoutItem * item);
func (this *QGraphicsLinearLayout) insertItem(args ...interface{}) () {
  // insertItem(int, class QGraphicsLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QGraphicsLayoutItem{}) // "QGraphicsLayoutItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout10insertItemEiP19QGraphicsLayoutItem
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "insertItem", args)
  }

}

  // proto:  void QGraphicsLinearLayout::~QGraphicsLinearLayout();
func (this *QGraphicsLinearLayout) FreeQGraphicsLinearLayout(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "~QGraphicsLinearLayout", args)
  }

}

  // proto:  void QGraphicsLinearLayout::dump(int indent);
func (this *QGraphicsLinearLayout) dump(args ...interface{}) () {
  // dump(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QGraphicsLinearLayout4dumpEi
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "dump", args)
  }

}

  // proto:  void QGraphicsLinearLayout::setStretchFactor(QGraphicsLayoutItem * item, int stretch);
func (this *QGraphicsLinearLayout) setStretchFactor(args ...interface{}) () {
  // setStretchFactor(class QGraphicsLayoutItem *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsLayoutItem{}) // "QGraphicsLayoutItem *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout16setStretchFactorEP19QGraphicsLayoutItemi
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "setStretchFactor", args)
  }

}

  // proto:  void QGraphicsLinearLayout::addItem(QGraphicsLayoutItem * item);
func (this *QGraphicsLinearLayout) addItem(args ...interface{}) () {
  // addItem(class QGraphicsLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsLayoutItem{}) // "QGraphicsLayoutItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout7addItemEP19QGraphicsLayoutItem
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "addItem", args)
  }

}

  // proto:  qreal QGraphicsLinearLayout::itemSpacing(int index);
func (this *QGraphicsLinearLayout) itemSpacing(args ...interface{}) () {
  // itemSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QGraphicsLinearLayout11itemSpacingEi
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "itemSpacing", args)
  }

}

  // proto:  void QGraphicsLinearLayout::removeAt(int index);
func (this *QGraphicsLinearLayout) removeAt(args ...interface{}) () {
  // removeAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout8removeAtEi
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "removeAt", args)
  }

}

  // proto:  void QGraphicsLinearLayout::insertStretch(int index, int stretch);
func (this *QGraphicsLinearLayout) insertStretch(args ...interface{}) () {
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
    // invoke: _ZN21QGraphicsLinearLayout13insertStretchEii
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "insertStretch", args)
  }

}

  // proto:  void QGraphicsLinearLayout::setItemSpacing(int index, qreal spacing);
func (this *QGraphicsLinearLayout) setItemSpacing(args ...interface{}) () {
  // setItemSpacing(int, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout14setItemSpacingEid
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "setItemSpacing", args)
  }

}

  // proto:  void QGraphicsLinearLayout::removeItem(QGraphicsLayoutItem * item);
func (this *QGraphicsLinearLayout) removeItem(args ...interface{}) () {
  // removeItem(class QGraphicsLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsLayoutItem{}) // "QGraphicsLayoutItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout10removeItemEP19QGraphicsLayoutItem
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "removeItem", args)
  }

}

  // proto:  int QGraphicsLinearLayout::stretchFactor(QGraphicsLayoutItem * item);
func (this *QGraphicsLinearLayout) stretchFactor(args ...interface{}) () {
  // stretchFactor(class QGraphicsLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsLayoutItem{}) // "QGraphicsLayoutItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QGraphicsLinearLayout13stretchFactorEP19QGraphicsLayoutItem
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "stretchFactor", args)
  }

}

// <= body block end

