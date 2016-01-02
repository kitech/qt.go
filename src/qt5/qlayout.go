package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtWidgets/qlayout.h
// dst-file: /src/widgets/qlayout.go
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
  // proto:  void QLayout::setContentsMargins(int left, int top, int right, int bottom);
extern void _ZN7QLayout18setContentsMarginsEiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  int QLayout::spacing();
extern void _ZNK7QLayout7spacingEv(void* qthis);
  // proto:  void QLayout::QLayout();
extern void* dector_ZN7QLayoutC1Ev();
extern void _ZN7QLayoutC1Ev(void* qthis);
  // proto:  QRect QLayout::geometry();
extern void _ZNK7QLayout8geometryEv(void* qthis);
  // proto:  int QLayout::count();
extern void _ZNK7QLayout5countEv(void* qthis);
  // proto:  QSize QLayout::maximumSize();
extern void _ZNK7QLayout11maximumSizeEv(void* qthis);
  // proto:  void QLayout::setMenuBar(QWidget * w);
extern void _ZN7QLayout10setMenuBarEP7QWidget(void* qthis, void* arg0);
  // proto:  int QLayout::indexOf(QWidget * );
extern void _ZNK7QLayout7indexOfEP7QWidget(void* qthis, void* arg0);
  // proto:  void QLayout::setEnabled(bool );
extern void _ZN7QLayout10setEnabledEb(void* qthis, bool arg0);
  // proto:  QSize QLayout::minimumSize();
extern void _ZNK7QLayout11minimumSizeEv(void* qthis);
  // proto:  QLayoutItem * QLayout::takeAt(int index);
extern void _ZN7QLayout6takeAtEi(void* qthis, int arg0);
  // proto:  QSize QLayout::totalMaximumSize();
extern void _ZNK7QLayout16totalMaximumSizeEv(void* qthis);
  // proto:  void QLayout::invalidate();
extern void _ZN7QLayout10invalidateEv(void* qthis);
  // proto:  void QLayout::update();
extern void _ZN7QLayout6updateEv(void* qthis);
  // proto:  QRect QLayout::contentsRect();
extern void _ZNK7QLayout12contentsRectEv(void* qthis);
  // proto:  QSize QLayout::totalSizeHint();
extern void _ZNK7QLayout13totalSizeHintEv(void* qthis);
  // proto:  void QLayout::QLayout(QWidget * parent);
extern void* dector_ZN7QLayoutC1EP7QWidget(void* arg0);
extern void _ZN7QLayoutC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QLayout::addItem(QLayoutItem * );
extern void _ZN7QLayout7addItemEP11QLayoutItem(void* qthis, void* arg0);
  // proto:  int QLayout::totalHeightForWidth(int w);
extern void _ZNK7QLayout19totalHeightForWidthEi(void* qthis, int arg0);
  // proto:  void QLayout::setMargin(int );
extern void _ZN7QLayout9setMarginEi(void* qthis, int arg0);
  // proto:  bool QLayout::isEmpty();
extern void _ZNK7QLayout7isEmptyEv(void* qthis);
  // proto:  void QLayout::addWidget(QWidget * w);
extern void _ZN7QLayout9addWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  void QLayout::getContentsMargins(int * left, int * top, int * right, int * bottom);
extern void _ZNK7QLayout18getContentsMarginsEPiS0_S0_S0_(void* qthis, int* arg0, int* arg1, int* arg2, int* arg3);
  // proto:  QLayout * QLayout::layout();
extern void _ZN7QLayout6layoutEv(void* qthis);
  // proto:  bool QLayout::activate();
extern void _ZN7QLayout8activateEv(void* qthis);
  // proto:  bool QLayout::isEnabled();
extern void _ZNK7QLayout9isEnabledEv(void* qthis);
  // proto:  void QLayout::~QLayout();
extern void _ZN7QLayoutD0Ev(void* qthis);
  // proto:  int QLayout::margin();
extern void _ZNK7QLayout6marginEv(void* qthis);
  // proto:  void QLayout::setSpacing(int );
extern void _ZN7QLayout10setSpacingEi(void* qthis, int arg0);
  // proto:  QWidget * QLayout::menuBar();
extern void _ZNK7QLayout7menuBarEv(void* qthis);
  // proto:  void QLayout::QLayout(const QLayout & );
extern void* dector_ZN7QLayoutC1ERKS_(void* arg0);
extern void _ZN7QLayoutC1ERKS_(void* qthis, void* arg0);
  // proto:  const QMetaObject * QLayout::metaObject();
extern void _ZNK7QLayout10metaObjectEv(void* qthis);
  // proto:  QLayoutItem * QLayout::itemAt(int index);
extern void _ZNK7QLayout6itemAtEi(void* qthis, int arg0);
  // proto:  QWidget * QLayout::parentWidget();
extern void _ZNK7QLayout12parentWidgetEv(void* qthis);
  // proto:  void QLayout::removeWidget(QWidget * w);
extern void _ZN7QLayout12removeWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  void QLayout::removeItem(QLayoutItem * );
extern void _ZN7QLayout10removeItemEP11QLayoutItem(void* qthis, void* arg0);
  // proto:  QMargins QLayout::contentsMargins();
extern void _ZNK7QLayout15contentsMarginsEv(void* qthis);
  // proto:  QSize QLayout::totalMinimumSize();
extern void _ZNK7QLayout16totalMinimumSizeEv(void* qthis);
  // proto:  void QLayout::setGeometry(const QRect & );
extern void _ZN7QLayout11setGeometryERK5QRect(void* qthis, void* arg0);
  // proto: static QSize QLayout::closestAcceptableSize(const QWidget * w, const QSize & s);
extern void _ZN7QLayout21closestAcceptableSizeEPK7QWidgetRK5QSize(void* arg0, void* arg1);
  // proto:  void QLayout::setContentsMargins(const QMargins & margins);
extern void _ZN7QLayout18setContentsMarginsERK8QMargins(void* qthis, void* arg0);
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

// class sizeof(QLayout)=1
type QLayout struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QLayout::setContentsMargins(int left, int top, int right, int bottom);
func (this *QLayout) setContentsMargins(args ...interface{}) () {
  // setContentsMargins(int, int, int, int)
  // setContentsMargins(const class QMargins &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QMargins{}) // "const QMargins &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout18setContentsMarginsEiiii
    // invoke: void setContentsMargins(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C._ZN7QLayout18setContentsMarginsEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN7QLayout18setContentsMarginsERK8QMargins
    // invoke: void setContentsMargins(const class QMargins &)
    var arg0 = args[0].(QMargins).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QLayout18setContentsMarginsERK8QMargins(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLayout", "setContentsMargins", args)
  }

}

  // proto:  int QLayout::spacing();
func (this *QLayout) spacing(args ...interface{}) () {
  // spacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout7spacingEv
    // invoke: int spacing()
    C._ZNK7QLayout7spacingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayout", "spacing", args)
  }

}

  // proto:  void QLayout::QLayout();
func NewQLayout(args ...interface{}) QLayout {
  return QLayout{}
}

  // proto:  QRect QLayout::geometry();
func (this *QLayout) geometry(args ...interface{}) () {
  // geometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout8geometryEv
    // invoke: QRect geometry()
    C._ZNK7QLayout8geometryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayout", "geometry", args)
  }

}

  // proto:  int QLayout::count();
func (this *QLayout) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout5countEv
    // invoke: int count()
    C._ZNK7QLayout5countEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayout", "count", args)
  }

}

  // proto:  QSize QLayout::maximumSize();
func (this *QLayout) maximumSize(args ...interface{}) () {
  // maximumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout11maximumSizeEv
    // invoke: QSize maximumSize()
    C._ZNK7QLayout11maximumSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayout", "maximumSize", args)
  }

}

  // proto:  void QLayout::setMenuBar(QWidget * w);
func (this *QLayout) setMenuBar(args ...interface{}) () {
  // setMenuBar(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout10setMenuBarEP7QWidget
    // invoke: void setMenuBar(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QLayout10setMenuBarEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLayout", "setMenuBar", args)
  }

}

  // proto:  int QLayout::indexOf(QWidget * );
func (this *QLayout) indexOf(args ...interface{}) () {
  // indexOf(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout7indexOfEP7QWidget
    // invoke: int indexOf(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QLayout7indexOfEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLayout", "indexOf", args)
  }

}

  // proto:  void QLayout::setEnabled(bool );
func (this *QLayout) setEnabled(args ...interface{}) () {
  // setEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout10setEnabledEb
    // invoke: void setEnabled(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN7QLayout10setEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLayout", "setEnabled", args)
  }

}

  // proto:  QSize QLayout::minimumSize();
func (this *QLayout) minimumSize(args ...interface{}) () {
  // minimumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout11minimumSizeEv
    // invoke: QSize minimumSize()
    C._ZNK7QLayout11minimumSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayout", "minimumSize", args)
  }

}

  // proto:  QLayoutItem * QLayout::takeAt(int index);
func (this *QLayout) takeAt(args ...interface{}) () {
  // takeAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout6takeAtEi
    // invoke: QLayoutItem * takeAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QLayout6takeAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLayout", "takeAt", args)
  }

}

  // proto:  QSize QLayout::totalMaximumSize();
func (this *QLayout) totalMaximumSize(args ...interface{}) () {
  // totalMaximumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout16totalMaximumSizeEv
    // invoke: QSize totalMaximumSize()
    C._ZNK7QLayout16totalMaximumSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayout", "totalMaximumSize", args)
  }

}

  // proto:  void QLayout::invalidate();
func (this *QLayout) invalidate(args ...interface{}) () {
  // invalidate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout10invalidateEv
    // invoke: void invalidate()
    C._ZN7QLayout10invalidateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayout", "invalidate", args)
  }

}

  // proto:  void QLayout::update();
func (this *QLayout) update(args ...interface{}) () {
  // update()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout6updateEv
    // invoke: void update()
    C._ZN7QLayout6updateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayout", "update", args)
  }

}

  // proto:  QRect QLayout::contentsRect();
func (this *QLayout) contentsRect(args ...interface{}) () {
  // contentsRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout12contentsRectEv
    // invoke: QRect contentsRect()
    C._ZNK7QLayout12contentsRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayout", "contentsRect", args)
  }

}

  // proto:  QSize QLayout::totalSizeHint();
func (this *QLayout) totalSizeHint(args ...interface{}) () {
  // totalSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout13totalSizeHintEv
    // invoke: QSize totalSizeHint()
    C._ZNK7QLayout13totalSizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayout", "totalSizeHint", args)
  }

}

  // proto:  void QLayout::addItem(QLayoutItem * );
func (this *QLayout) addItem(args ...interface{}) () {
  // addItem(class QLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLayoutItem{}) // "QLayoutItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout7addItemEP11QLayoutItem
    // invoke: void addItem(class QLayoutItem *)
    var arg0 = args[0].(QLayoutItem).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QLayout7addItemEP11QLayoutItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLayout", "addItem", args)
  }

}

  // proto:  int QLayout::totalHeightForWidth(int w);
func (this *QLayout) totalHeightForWidth(args ...interface{}) () {
  // totalHeightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout19totalHeightForWidthEi
    // invoke: int totalHeightForWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK7QLayout19totalHeightForWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLayout", "totalHeightForWidth", args)
  }

}

  // proto:  void QLayout::setMargin(int );
func (this *QLayout) setMargin(args ...interface{}) () {
  // setMargin(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout9setMarginEi
    // invoke: void setMargin(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QLayout9setMarginEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLayout", "setMargin", args)
  }

}

  // proto:  bool QLayout::isEmpty();
func (this *QLayout) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout7isEmptyEv
    // invoke: bool isEmpty()
    C._ZNK7QLayout7isEmptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayout", "isEmpty", args)
  }

}

  // proto:  void QLayout::addWidget(QWidget * w);
func (this *QLayout) addWidget(args ...interface{}) () {
  // addWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout9addWidgetEP7QWidget
    // invoke: void addWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QLayout9addWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLayout", "addWidget", args)
  }

}

  // proto:  void QLayout::getContentsMargins(int * left, int * top, int * right, int * bottom);
func (this *QLayout) getContentsMargins(args ...interface{}) () {
  // getContentsMargins(int *, int *, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "int *"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout18getContentsMarginsEPiS0_S0_S0_
    // invoke: void getContentsMargins(int *, int *, int *, int *)
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int32_t)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    C._ZNK7QLayout18getContentsMarginsEPiS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QLayout", "getContentsMargins", args)
  }

}

  // proto:  QLayout * QLayout::layout();
func (this *QLayout) layout(args ...interface{}) () {
  // layout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout6layoutEv
    // invoke: QLayout * layout()
    C._ZN7QLayout6layoutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayout", "layout", args)
  }

}

  // proto:  bool QLayout::activate();
func (this *QLayout) activate(args ...interface{}) () {
  // activate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout8activateEv
    // invoke: bool activate()
    C._ZN7QLayout8activateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayout", "activate", args)
  }

}

  // proto:  bool QLayout::isEnabled();
func (this *QLayout) isEnabled(args ...interface{}) () {
  // isEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout9isEnabledEv
    // invoke: bool isEnabled()
    C._ZNK7QLayout9isEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayout", "isEnabled", args)
  }

}

  // proto:  void QLayout::~QLayout();
func (this *QLayout) FreeQLayout(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QLayout", "~QLayout", args)
  }

}

  // proto:  int QLayout::margin();
func (this *QLayout) margin(args ...interface{}) () {
  // margin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout6marginEv
    // invoke: int margin()
    C._ZNK7QLayout6marginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayout", "margin", args)
  }

}

  // proto:  void QLayout::setSpacing(int );
func (this *QLayout) setSpacing(args ...interface{}) () {
  // setSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout10setSpacingEi
    // invoke: void setSpacing(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QLayout10setSpacingEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLayout", "setSpacing", args)
  }

}

  // proto:  QWidget * QLayout::menuBar();
func (this *QLayout) menuBar(args ...interface{}) () {
  // menuBar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout7menuBarEv
    // invoke: QWidget * menuBar()
    C._ZNK7QLayout7menuBarEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayout", "menuBar", args)
  }

}

  // proto:  const QMetaObject * QLayout::metaObject();
func (this *QLayout) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK7QLayout10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayout", "metaObject", args)
  }

}

  // proto:  QLayoutItem * QLayout::itemAt(int index);
func (this *QLayout) itemAt(args ...interface{}) () {
  // itemAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout6itemAtEi
    // invoke: QLayoutItem * itemAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK7QLayout6itemAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLayout", "itemAt", args)
  }

}

  // proto:  QWidget * QLayout::parentWidget();
func (this *QLayout) parentWidget(args ...interface{}) () {
  // parentWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout12parentWidgetEv
    // invoke: QWidget * parentWidget()
    C._ZNK7QLayout12parentWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayout", "parentWidget", args)
  }

}

  // proto:  void QLayout::removeWidget(QWidget * w);
func (this *QLayout) removeWidget(args ...interface{}) () {
  // removeWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout12removeWidgetEP7QWidget
    // invoke: void removeWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QLayout12removeWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLayout", "removeWidget", args)
  }

}

  // proto:  void QLayout::removeItem(QLayoutItem * );
func (this *QLayout) removeItem(args ...interface{}) () {
  // removeItem(class QLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLayoutItem{}) // "QLayoutItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout10removeItemEP11QLayoutItem
    // invoke: void removeItem(class QLayoutItem *)
    var arg0 = args[0].(QLayoutItem).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QLayout10removeItemEP11QLayoutItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLayout", "removeItem", args)
  }

}

  // proto:  QMargins QLayout::contentsMargins();
func (this *QLayout) contentsMargins(args ...interface{}) () {
  // contentsMargins()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout15contentsMarginsEv
    // invoke: QMargins contentsMargins()
    C._ZNK7QLayout15contentsMarginsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayout", "contentsMargins", args)
  }

}

  // proto:  QSize QLayout::totalMinimumSize();
func (this *QLayout) totalMinimumSize(args ...interface{}) () {
  // totalMinimumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout16totalMinimumSizeEv
    // invoke: QSize totalMinimumSize()
    C._ZNK7QLayout16totalMinimumSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayout", "totalMinimumSize", args)
  }

}

  // proto:  void QLayout::setGeometry(const QRect & );
func (this *QLayout) setGeometry(args ...interface{}) () {
  // setGeometry(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout11setGeometryERK5QRect
    // invoke: void setGeometry(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QLayout11setGeometryERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLayout", "setGeometry", args)
  }

}

  // proto: static QSize QLayout::closestAcceptableSize(const QWidget * w, const QSize & s);
func (this *QLayout) closestAcceptableSize_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QLayout", "closestAcceptableSize", args)
  }

}

// <= body block end

