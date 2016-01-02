package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
  // proto:  void QHBoxLayout::QHBoxLayout(QWidget * parent);
extern void* dector_ZN11QHBoxLayoutC1EP7QWidget(void* arg0);
extern void _ZN11QHBoxLayoutC1EP7QWidget(void* qthis, void* arg0);
  // proto:  const QMetaObject * QHBoxLayout::metaObject();
extern void _ZNK11QHBoxLayout10metaObjectEv(void* qthis);
  // proto:  void QHBoxLayout::~QHBoxLayout();
extern void _ZN11QHBoxLayoutD0Ev(void* qthis);
  // proto:  void QHBoxLayout::QHBoxLayout(const QHBoxLayout & );
extern void* dector_ZN11QHBoxLayoutC1ERKS_(void* arg0);
extern void _ZN11QHBoxLayoutC1ERKS_(void* qthis, void* arg0);
  // proto:  void QHBoxLayout::QHBoxLayout();
extern void* dector_ZN11QHBoxLayoutC1Ev();
extern void _ZN11QHBoxLayoutC1Ev(void* qthis);
  // proto:  int QBoxLayout::spacing();
extern void _ZNK10QBoxLayout7spacingEv(void* qthis);
  // proto:  bool QBoxLayout::hasHeightForWidth();
extern void _ZNK10QBoxLayout17hasHeightForWidthEv(void* qthis);
  // proto:  void QBoxLayout::addItem(QLayoutItem * );
extern void _ZN10QBoxLayout7addItemEP11QLayoutItem(void* qthis, void* arg0);
  // proto:  QSize QBoxLayout::sizeHint();
extern void _ZNK10QBoxLayout8sizeHintEv(void* qthis);
  // proto:  void QBoxLayout::~QBoxLayout();
extern void _ZN10QBoxLayoutD0Ev(void* qthis);
  // proto:  void QBoxLayout::insertSpacing(int index, int size);
extern void _ZN10QBoxLayout13insertSpacingEii(void* qthis, int arg0, int arg1);
  // proto:  void QBoxLayout::setStretch(int index, int stretch);
extern void _ZN10QBoxLayout10setStretchEii(void* qthis, int arg0, int arg1);
  // proto:  void QBoxLayout::QBoxLayout(const QBoxLayout & );
extern void* dector_ZN10QBoxLayoutC1ERKS_(void* arg0);
extern void _ZN10QBoxLayoutC1ERKS_(void* qthis, void* arg0);
  // proto:  void QBoxLayout::insertStretch(int index, int stretch);
extern void _ZN10QBoxLayout13insertStretchEii(void* qthis, int arg0, int arg1);
  // proto:  void QBoxLayout::addLayout(QLayout * layout, int stretch);
extern void _ZN10QBoxLayout9addLayoutEP7QLayouti(void* qthis, void* arg0, int arg1);
  // proto:  bool QBoxLayout::setStretchFactor(QWidget * w, int stretch);
extern void _ZN10QBoxLayout16setStretchFactorEP7QWidgeti(void* qthis, void* arg0, int arg1);
  // proto:  void QBoxLayout::invalidate();
extern void _ZN10QBoxLayout10invalidateEv(void* qthis);
  // proto:  void QBoxLayout::setGeometry(const QRect & );
extern void _ZN10QBoxLayout11setGeometryERK5QRect(void* qthis, void* arg0);
  // proto:  void QBoxLayout::addStretch(int stretch);
extern void _ZN10QBoxLayout10addStretchEi(void* qthis, int arg0);
  // proto:  void QBoxLayout::insertLayout(int index, QLayout * layout, int stretch);
extern void _ZN10QBoxLayout12insertLayoutEiP7QLayouti(void* qthis, int arg0, void* arg1, int arg2);
  // proto:  bool QBoxLayout::setStretchFactor(QLayout * l, int stretch);
extern void _ZN10QBoxLayout16setStretchFactorEP7QLayouti(void* qthis, void* arg0, int arg1);
  // proto:  int QBoxLayout::count();
extern void _ZNK10QBoxLayout5countEv(void* qthis);
  // proto:  QLayoutItem * QBoxLayout::itemAt(int );
extern void _ZNK10QBoxLayout6itemAtEi(void* qthis, int arg0);
  // proto:  const QMetaObject * QBoxLayout::metaObject();
extern void _ZNK10QBoxLayout10metaObjectEv(void* qthis);
  // proto:  void QBoxLayout::insertSpacerItem(int index, QSpacerItem * spacerItem);
extern void _ZN10QBoxLayout16insertSpacerItemEiP11QSpacerItem(void* qthis, int arg0, void* arg1);
  // proto:  int QBoxLayout::heightForWidth(int );
extern void _ZNK10QBoxLayout14heightForWidthEi(void* qthis, int arg0);
  // proto:  void QBoxLayout::addStrut(int );
extern void _ZN10QBoxLayout8addStrutEi(void* qthis, int arg0);
  // proto:  QSize QBoxLayout::maximumSize();
extern void _ZNK10QBoxLayout11maximumSizeEv(void* qthis);
  // proto:  int QBoxLayout::stretch(int index);
extern void _ZNK10QBoxLayout7stretchEi(void* qthis, int arg0);
  // proto:  void QBoxLayout::addSpacerItem(QSpacerItem * spacerItem);
extern void _ZN10QBoxLayout13addSpacerItemEP11QSpacerItem(void* qthis, void* arg0);
  // proto:  int QBoxLayout::minimumHeightForWidth(int );
extern void _ZNK10QBoxLayout21minimumHeightForWidthEi(void* qthis, int arg0);
  // proto:  QSize QBoxLayout::minimumSize();
extern void _ZNK10QBoxLayout11minimumSizeEv(void* qthis);
  // proto:  void QBoxLayout::setSpacing(int spacing);
extern void _ZN10QBoxLayout10setSpacingEi(void* qthis, int arg0);
  // proto:  QLayoutItem * QBoxLayout::takeAt(int );
extern void _ZN10QBoxLayout6takeAtEi(void* qthis, int arg0);
  // proto:  void QBoxLayout::insertItem(int index, QLayoutItem * );
extern void _ZN10QBoxLayout10insertItemEiP11QLayoutItem(void* qthis, int arg0, void* arg1);
  // proto:  void QBoxLayout::addSpacing(int size);
extern void _ZN10QBoxLayout10addSpacingEi(void* qthis, int arg0);
  // proto:  void QVBoxLayout::QVBoxLayout();
extern void* dector_ZN11QVBoxLayoutC1Ev();
extern void _ZN11QVBoxLayoutC1Ev(void* qthis);
  // proto:  const QMetaObject * QVBoxLayout::metaObject();
extern void _ZNK11QVBoxLayout10metaObjectEv(void* qthis);
  // proto:  void QVBoxLayout::QVBoxLayout(const QVBoxLayout & );
extern void* dector_ZN11QVBoxLayoutC1ERKS_(void* arg0);
extern void _ZN11QVBoxLayoutC1ERKS_(void* qthis, void* arg0);
  // proto:  void QVBoxLayout::QVBoxLayout(QWidget * parent);
extern void* dector_ZN11QVBoxLayoutC1EP7QWidget(void* arg0);
extern void _ZN11QVBoxLayoutC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QVBoxLayout::~QVBoxLayout();
extern void _ZN11QVBoxLayoutD0Ev(void* qthis);
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
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QBoxLayout)=1
type QBoxLayout struct {
  /*qbase*/ QLayout;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QVBoxLayout)=1
type QVBoxLayout struct {
  /*qbase*/ QBoxLayout;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QHBoxLayout::QHBoxLayout(QWidget * parent);
func NewQHBoxLayout(args ...interface{}) QHBoxLayout {
  return QHBoxLayout{}
}

  // proto:  const QMetaObject * QHBoxLayout::metaObject();
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
  default:
    qtrt.ErrorResolve("QHBoxLayout", "metaObject", args)
  }

}

  // proto:  void QHBoxLayout::~QHBoxLayout();
func (this *QHBoxLayout) FreeQHBoxLayout(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QHBoxLayout", "~QHBoxLayout", args)
  }

}

  // proto:  int QBoxLayout::spacing();
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
  default:
    qtrt.ErrorResolve("QBoxLayout", "spacing", args)
  }

}

  // proto:  bool QBoxLayout::hasHeightForWidth();
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
  default:
    qtrt.ErrorResolve("QBoxLayout", "hasHeightForWidth", args)
  }

}

  // proto:  void QBoxLayout::addItem(QLayoutItem * );
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
  default:
    qtrt.ErrorResolve("QBoxLayout", "addItem", args)
  }

}

  // proto:  QSize QBoxLayout::sizeHint();
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
  default:
    qtrt.ErrorResolve("QBoxLayout", "sizeHint", args)
  }

}

  // proto:  void QBoxLayout::~QBoxLayout();
func (this *QBoxLayout) FreeQBoxLayout(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QBoxLayout", "~QBoxLayout", args)
  }

}

  // proto:  void QBoxLayout::insertSpacing(int index, int size);
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
  default:
    qtrt.ErrorResolve("QBoxLayout", "insertSpacing", args)
  }

}

  // proto:  void QBoxLayout::setStretch(int index, int stretch);
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
  default:
    qtrt.ErrorResolve("QBoxLayout", "setStretch", args)
  }

}

  // proto:  void QBoxLayout::QBoxLayout(const QBoxLayout & );
func NewQBoxLayout(args ...interface{}) QBoxLayout {
  return QBoxLayout{}
}

  // proto:  void QBoxLayout::insertStretch(int index, int stretch);
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
  default:
    qtrt.ErrorResolve("QBoxLayout", "insertStretch", args)
  }

}

  // proto:  void QBoxLayout::addLayout(QLayout * layout, int stretch);
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
  default:
    qtrt.ErrorResolve("QBoxLayout", "addLayout", args)
  }

}

  // proto:  bool QBoxLayout::setStretchFactor(QWidget * w, int stretch);
func (this *QBoxLayout) setStretchFactor(args ...interface{}) () {
  // setStretchFactor(class QWidget *, int)
  // setStretchFactor(class QLayout *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QLayout{}) // "QLayout *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QBoxLayout16setStretchFactorEP7QWidgeti
  case 1:
    // invoke: _ZN10QBoxLayout16setStretchFactorEP7QLayouti
  default:
    qtrt.ErrorResolve("QBoxLayout", "setStretchFactor", args)
  }

}

  // proto:  void QBoxLayout::invalidate();
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
  default:
    qtrt.ErrorResolve("QBoxLayout", "invalidate", args)
  }

}

  // proto:  void QBoxLayout::setGeometry(const QRect & );
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
  default:
    qtrt.ErrorResolve("QBoxLayout", "setGeometry", args)
  }

}

  // proto:  void QBoxLayout::addStretch(int stretch);
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
  default:
    qtrt.ErrorResolve("QBoxLayout", "addStretch", args)
  }

}

  // proto:  void QBoxLayout::insertLayout(int index, QLayout * layout, int stretch);
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
  default:
    qtrt.ErrorResolve("QBoxLayout", "insertLayout", args)
  }

}

  // proto:  int QBoxLayout::count();
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
  default:
    qtrt.ErrorResolve("QBoxLayout", "count", args)
  }

}

  // proto:  QLayoutItem * QBoxLayout::itemAt(int );
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
  default:
    qtrt.ErrorResolve("QBoxLayout", "itemAt", args)
  }

}

  // proto:  const QMetaObject * QBoxLayout::metaObject();
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
  default:
    qtrt.ErrorResolve("QBoxLayout", "metaObject", args)
  }

}

  // proto:  void QBoxLayout::insertSpacerItem(int index, QSpacerItem * spacerItem);
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
  default:
    qtrt.ErrorResolve("QBoxLayout", "insertSpacerItem", args)
  }

}

  // proto:  int QBoxLayout::heightForWidth(int );
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
  default:
    qtrt.ErrorResolve("QBoxLayout", "heightForWidth", args)
  }

}

  // proto:  void QBoxLayout::addStrut(int );
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
  default:
    qtrt.ErrorResolve("QBoxLayout", "addStrut", args)
  }

}

  // proto:  QSize QBoxLayout::maximumSize();
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
  default:
    qtrt.ErrorResolve("QBoxLayout", "maximumSize", args)
  }

}

  // proto:  int QBoxLayout::stretch(int index);
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
  default:
    qtrt.ErrorResolve("QBoxLayout", "stretch", args)
  }

}

  // proto:  void QBoxLayout::addSpacerItem(QSpacerItem * spacerItem);
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
  default:
    qtrt.ErrorResolve("QBoxLayout", "addSpacerItem", args)
  }

}

  // proto:  int QBoxLayout::minimumHeightForWidth(int );
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
  default:
    qtrt.ErrorResolve("QBoxLayout", "minimumHeightForWidth", args)
  }

}

  // proto:  QSize QBoxLayout::minimumSize();
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
  default:
    qtrt.ErrorResolve("QBoxLayout", "minimumSize", args)
  }

}

  // proto:  void QBoxLayout::setSpacing(int spacing);
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
  default:
    qtrt.ErrorResolve("QBoxLayout", "setSpacing", args)
  }

}

  // proto:  QLayoutItem * QBoxLayout::takeAt(int );
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
  default:
    qtrt.ErrorResolve("QBoxLayout", "takeAt", args)
  }

}

  // proto:  void QBoxLayout::insertItem(int index, QLayoutItem * );
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
  default:
    qtrt.ErrorResolve("QBoxLayout", "insertItem", args)
  }

}

  // proto:  void QBoxLayout::addSpacing(int size);
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
  default:
    qtrt.ErrorResolve("QBoxLayout", "addSpacing", args)
  }

}

  // proto:  void QVBoxLayout::QVBoxLayout();
func NewQVBoxLayout(args ...interface{}) QVBoxLayout {
  return QVBoxLayout{}
}

  // proto:  const QMetaObject * QVBoxLayout::metaObject();
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
  default:
    qtrt.ErrorResolve("QVBoxLayout", "metaObject", args)
  }

}

  // proto:  void QVBoxLayout::~QVBoxLayout();
func (this *QVBoxLayout) FreeQVBoxLayout(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QVBoxLayout", "~QVBoxLayout", args)
  }

}

// <= body block end

