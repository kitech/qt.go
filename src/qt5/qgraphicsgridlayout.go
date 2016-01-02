package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtWidgets/qgraphicsgridlayout.h
// dst-file: /src/widgets/qgraphicsgridlayout.go
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
  // proto:  void QGraphicsGridLayout::setRowPreferredHeight(int row, qreal height);
extern void _ZN19QGraphicsGridLayout21setRowPreferredHeightEid(void* qthis, int arg0, double arg1);
  // proto:  int QGraphicsGridLayout::columnCount();
extern void _ZNK19QGraphicsGridLayout11columnCountEv(void* qthis);
  // proto:  QGraphicsLayoutItem * QGraphicsGridLayout::itemAt(int index);
extern void _ZNK19QGraphicsGridLayout6itemAtEi(void* qthis, int arg0);
  // proto:  int QGraphicsGridLayout::count();
extern void _ZNK19QGraphicsGridLayout5countEv(void* qthis);
  // proto:  void QGraphicsGridLayout::setColumnFixedWidth(int column, qreal width);
extern void _ZN19QGraphicsGridLayout19setColumnFixedWidthEid(void* qthis, int arg0, double arg1);
  // proto:  void QGraphicsGridLayout::setColumnMaximumWidth(int column, qreal width);
extern void _ZN19QGraphicsGridLayout21setColumnMaximumWidthEid(void* qthis, int arg0, double arg1);
  // proto:  int QGraphicsGridLayout::rowStretchFactor(int row);
extern void _ZNK19QGraphicsGridLayout16rowStretchFactorEi(void* qthis, int arg0);
  // proto:  qreal QGraphicsGridLayout::verticalSpacing();
extern void _ZNK19QGraphicsGridLayout15verticalSpacingEv(void* qthis);
  // proto:  int QGraphicsGridLayout::columnStretchFactor(int column);
extern void _ZNK19QGraphicsGridLayout19columnStretchFactorEi(void* qthis, int arg0);
  // proto:  void QGraphicsGridLayout::setRowMaximumHeight(int row, qreal height);
extern void _ZN19QGraphicsGridLayout19setRowMaximumHeightEid(void* qthis, int arg0, double arg1);
  // proto:  void QGraphicsGridLayout::removeItem(QGraphicsLayoutItem * item);
extern void _ZN19QGraphicsGridLayout10removeItemEP19QGraphicsLayoutItem(void* qthis, void* arg0);
  // proto:  void QGraphicsGridLayout::~QGraphicsGridLayout();
extern void _ZN19QGraphicsGridLayoutD0Ev(void* qthis);
  // proto:  qreal QGraphicsGridLayout::rowMinimumHeight(int row);
extern void _ZNK19QGraphicsGridLayout16rowMinimumHeightEi(void* qthis, int arg0);
  // proto:  qreal QGraphicsGridLayout::rowMaximumHeight(int row);
extern void _ZNK19QGraphicsGridLayout16rowMaximumHeightEi(void* qthis, int arg0);
  // proto:  void QGraphicsGridLayout::QGraphicsGridLayout(QGraphicsLayoutItem * parent);
extern void* dector_ZN19QGraphicsGridLayoutC1EP19QGraphicsLayoutItem(void* arg0);
extern void _ZN19QGraphicsGridLayoutC1EP19QGraphicsLayoutItem(void* qthis, void* arg0);
  // proto:  void QGraphicsGridLayout::setColumnSpacing(int column, qreal spacing);
extern void _ZN19QGraphicsGridLayout16setColumnSpacingEid(void* qthis, int arg0, double arg1);
  // proto:  qreal QGraphicsGridLayout::rowSpacing(int row);
extern void _ZNK19QGraphicsGridLayout10rowSpacingEi(void* qthis, int arg0);
  // proto:  qreal QGraphicsGridLayout::columnMaximumWidth(int column);
extern void _ZNK19QGraphicsGridLayout18columnMaximumWidthEi(void* qthis, int arg0);
  // proto:  void QGraphicsGridLayout::setRowFixedHeight(int row, qreal height);
extern void _ZN19QGraphicsGridLayout17setRowFixedHeightEid(void* qthis, int arg0, double arg1);
  // proto:  qreal QGraphicsGridLayout::rowPreferredHeight(int row);
extern void _ZNK19QGraphicsGridLayout18rowPreferredHeightEi(void* qthis, int arg0);
  // proto:  QGraphicsLayoutItem * QGraphicsGridLayout::itemAt(int row, int column);
extern void _ZNK19QGraphicsGridLayout6itemAtEii(void* qthis, int arg0, int arg1);
  // proto:  void QGraphicsGridLayout::setVerticalSpacing(qreal spacing);
extern void _ZN19QGraphicsGridLayout18setVerticalSpacingEd(void* qthis, double arg0);
  // proto:  void QGraphicsGridLayout::setGeometry(const QRectF & rect);
extern void _ZN19QGraphicsGridLayout11setGeometryERK6QRectF(void* qthis, void* arg0);
  // proto:  int QGraphicsGridLayout::rowCount();
extern void _ZNK19QGraphicsGridLayout8rowCountEv(void* qthis);
  // proto:  void QGraphicsGridLayout::setSpacing(qreal spacing);
extern void _ZN19QGraphicsGridLayout10setSpacingEd(void* qthis, double arg0);
  // proto:  void QGraphicsGridLayout::setRowStretchFactor(int row, int stretch);
extern void _ZN19QGraphicsGridLayout19setRowStretchFactorEii(void* qthis, int arg0, int arg1);
  // proto:  qreal QGraphicsGridLayout::columnMinimumWidth(int column);
extern void _ZNK19QGraphicsGridLayout18columnMinimumWidthEi(void* qthis, int arg0);
  // proto:  void QGraphicsGridLayout::setColumnMinimumWidth(int column, qreal width);
extern void _ZN19QGraphicsGridLayout21setColumnMinimumWidthEid(void* qthis, int arg0, double arg1);
  // proto:  void QGraphicsGridLayout::setRowMinimumHeight(int row, qreal height);
extern void _ZN19QGraphicsGridLayout19setRowMinimumHeightEid(void* qthis, int arg0, double arg1);
  // proto:  void QGraphicsGridLayout::setHorizontalSpacing(qreal spacing);
extern void _ZN19QGraphicsGridLayout20setHorizontalSpacingEd(void* qthis, double arg0);
  // proto:  qreal QGraphicsGridLayout::horizontalSpacing();
extern void _ZNK19QGraphicsGridLayout17horizontalSpacingEv(void* qthis);
  // proto:  void QGraphicsGridLayout::setColumnStretchFactor(int column, int stretch);
extern void _ZN19QGraphicsGridLayout22setColumnStretchFactorEii(void* qthis, int arg0, int arg1);
  // proto:  void QGraphicsGridLayout::invalidate();
extern void _ZN19QGraphicsGridLayout10invalidateEv(void* qthis);
  // proto:  qreal QGraphicsGridLayout::columnPreferredWidth(int column);
extern void _ZNK19QGraphicsGridLayout20columnPreferredWidthEi(void* qthis, int arg0);
  // proto:  void QGraphicsGridLayout::setColumnPreferredWidth(int column, qreal width);
extern void _ZN19QGraphicsGridLayout23setColumnPreferredWidthEid(void* qthis, int arg0, double arg1);
  // proto:  qreal QGraphicsGridLayout::columnSpacing(int column);
extern void _ZNK19QGraphicsGridLayout13columnSpacingEi(void* qthis, int arg0);
  // proto:  void QGraphicsGridLayout::QGraphicsGridLayout(const QGraphicsGridLayout & );
extern void* dector_ZN19QGraphicsGridLayoutC1ERKS_(void* arg0);
extern void _ZN19QGraphicsGridLayoutC1ERKS_(void* qthis, void* arg0);
  // proto:  void QGraphicsGridLayout::setRowSpacing(int row, qreal spacing);
extern void _ZN19QGraphicsGridLayout13setRowSpacingEid(void* qthis, int arg0, double arg1);
  // proto:  void QGraphicsGridLayout::removeAt(int index);
extern void _ZN19QGraphicsGridLayout8removeAtEi(void* qthis, int arg0);
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

// class sizeof(QGraphicsGridLayout)=1
type QGraphicsGridLayout struct {
  /*qbase*/ QGraphicsLayout;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QGraphicsGridLayout::setRowPreferredHeight(int row, qreal height);
func (this *QGraphicsGridLayout) setRowPreferredHeight(args ...interface{}) () {
  // setRowPreferredHeight(int, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout21setRowPreferredHeightEid
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setRowPreferredHeight", args)
  }

}

  // proto:  int QGraphicsGridLayout::columnCount();
func (this *QGraphicsGridLayout) columnCount(args ...interface{}) () {
  // columnCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout11columnCountEv
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "columnCount", args)
  }

}

  // proto:  QGraphicsLayoutItem * QGraphicsGridLayout::itemAt(int index);
func (this *QGraphicsGridLayout) itemAt(args ...interface{}) () {
  // itemAt(int)
  // itemAt(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout6itemAtEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK19QGraphicsGridLayout6itemAtEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "itemAt", args)
  }

}

  // proto:  int QGraphicsGridLayout::count();
func (this *QGraphicsGridLayout) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout5countEv
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "count", args)
  }

}

  // proto:  void QGraphicsGridLayout::setColumnFixedWidth(int column, qreal width);
func (this *QGraphicsGridLayout) setColumnFixedWidth(args ...interface{}) () {
  // setColumnFixedWidth(int, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout19setColumnFixedWidthEid
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setColumnFixedWidth", args)
  }

}

  // proto:  void QGraphicsGridLayout::setColumnMaximumWidth(int column, qreal width);
func (this *QGraphicsGridLayout) setColumnMaximumWidth(args ...interface{}) () {
  // setColumnMaximumWidth(int, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout21setColumnMaximumWidthEid
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setColumnMaximumWidth", args)
  }

}

  // proto:  int QGraphicsGridLayout::rowStretchFactor(int row);
func (this *QGraphicsGridLayout) rowStretchFactor(args ...interface{}) () {
  // rowStretchFactor(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout16rowStretchFactorEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "rowStretchFactor", args)
  }

}

  // proto:  qreal QGraphicsGridLayout::verticalSpacing();
func (this *QGraphicsGridLayout) verticalSpacing(args ...interface{}) () {
  // verticalSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout15verticalSpacingEv
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "verticalSpacing", args)
  }

}

  // proto:  int QGraphicsGridLayout::columnStretchFactor(int column);
func (this *QGraphicsGridLayout) columnStretchFactor(args ...interface{}) () {
  // columnStretchFactor(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout19columnStretchFactorEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "columnStretchFactor", args)
  }

}

  // proto:  void QGraphicsGridLayout::setRowMaximumHeight(int row, qreal height);
func (this *QGraphicsGridLayout) setRowMaximumHeight(args ...interface{}) () {
  // setRowMaximumHeight(int, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout19setRowMaximumHeightEid
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setRowMaximumHeight", args)
  }

}

  // proto:  void QGraphicsGridLayout::removeItem(QGraphicsLayoutItem * item);
func (this *QGraphicsGridLayout) removeItem(args ...interface{}) () {
  // removeItem(class QGraphicsLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsLayoutItem{}) // "QGraphicsLayoutItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout10removeItemEP19QGraphicsLayoutItem
    var arg0 = args[0].(QGraphicsLayoutItem).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "removeItem", args)
  }

}

  // proto:  void QGraphicsGridLayout::~QGraphicsGridLayout();
func (this *QGraphicsGridLayout) FreeQGraphicsGridLayout(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "~QGraphicsGridLayout", args)
  }

}

  // proto:  qreal QGraphicsGridLayout::rowMinimumHeight(int row);
func (this *QGraphicsGridLayout) rowMinimumHeight(args ...interface{}) () {
  // rowMinimumHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout16rowMinimumHeightEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "rowMinimumHeight", args)
  }

}

  // proto:  qreal QGraphicsGridLayout::rowMaximumHeight(int row);
func (this *QGraphicsGridLayout) rowMaximumHeight(args ...interface{}) () {
  // rowMaximumHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout16rowMaximumHeightEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "rowMaximumHeight", args)
  }

}

  // proto:  void QGraphicsGridLayout::QGraphicsGridLayout(QGraphicsLayoutItem * parent);
func NewQGraphicsGridLayout(args ...interface{}) QGraphicsGridLayout {
  return QGraphicsGridLayout{}
}

  // proto:  void QGraphicsGridLayout::setColumnSpacing(int column, qreal spacing);
func (this *QGraphicsGridLayout) setColumnSpacing(args ...interface{}) () {
  // setColumnSpacing(int, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout16setColumnSpacingEid
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setColumnSpacing", args)
  }

}

  // proto:  qreal QGraphicsGridLayout::rowSpacing(int row);
func (this *QGraphicsGridLayout) rowSpacing(args ...interface{}) () {
  // rowSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout10rowSpacingEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "rowSpacing", args)
  }

}

  // proto:  qreal QGraphicsGridLayout::columnMaximumWidth(int column);
func (this *QGraphicsGridLayout) columnMaximumWidth(args ...interface{}) () {
  // columnMaximumWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout18columnMaximumWidthEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "columnMaximumWidth", args)
  }

}

  // proto:  void QGraphicsGridLayout::setRowFixedHeight(int row, qreal height);
func (this *QGraphicsGridLayout) setRowFixedHeight(args ...interface{}) () {
  // setRowFixedHeight(int, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout17setRowFixedHeightEid
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setRowFixedHeight", args)
  }

}

  // proto:  qreal QGraphicsGridLayout::rowPreferredHeight(int row);
func (this *QGraphicsGridLayout) rowPreferredHeight(args ...interface{}) () {
  // rowPreferredHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout18rowPreferredHeightEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "rowPreferredHeight", args)
  }

}

  // proto:  void QGraphicsGridLayout::setVerticalSpacing(qreal spacing);
func (this *QGraphicsGridLayout) setVerticalSpacing(args ...interface{}) () {
  // setVerticalSpacing(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout18setVerticalSpacingEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setVerticalSpacing", args)
  }

}

  // proto:  void QGraphicsGridLayout::setGeometry(const QRectF & rect);
func (this *QGraphicsGridLayout) setGeometry(args ...interface{}) () {
  // setGeometry(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout11setGeometryERK6QRectF
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setGeometry", args)
  }

}

  // proto:  int QGraphicsGridLayout::rowCount();
func (this *QGraphicsGridLayout) rowCount(args ...interface{}) () {
  // rowCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout8rowCountEv
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "rowCount", args)
  }

}

  // proto:  void QGraphicsGridLayout::setSpacing(qreal spacing);
func (this *QGraphicsGridLayout) setSpacing(args ...interface{}) () {
  // setSpacing(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout10setSpacingEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setSpacing", args)
  }

}

  // proto:  void QGraphicsGridLayout::setRowStretchFactor(int row, int stretch);
func (this *QGraphicsGridLayout) setRowStretchFactor(args ...interface{}) () {
  // setRowStretchFactor(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout19setRowStretchFactorEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setRowStretchFactor", args)
  }

}

  // proto:  qreal QGraphicsGridLayout::columnMinimumWidth(int column);
func (this *QGraphicsGridLayout) columnMinimumWidth(args ...interface{}) () {
  // columnMinimumWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout18columnMinimumWidthEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "columnMinimumWidth", args)
  }

}

  // proto:  void QGraphicsGridLayout::setColumnMinimumWidth(int column, qreal width);
func (this *QGraphicsGridLayout) setColumnMinimumWidth(args ...interface{}) () {
  // setColumnMinimumWidth(int, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout21setColumnMinimumWidthEid
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setColumnMinimumWidth", args)
  }

}

  // proto:  void QGraphicsGridLayout::setRowMinimumHeight(int row, qreal height);
func (this *QGraphicsGridLayout) setRowMinimumHeight(args ...interface{}) () {
  // setRowMinimumHeight(int, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout19setRowMinimumHeightEid
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setRowMinimumHeight", args)
  }

}

  // proto:  void QGraphicsGridLayout::setHorizontalSpacing(qreal spacing);
func (this *QGraphicsGridLayout) setHorizontalSpacing(args ...interface{}) () {
  // setHorizontalSpacing(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout20setHorizontalSpacingEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setHorizontalSpacing", args)
  }

}

  // proto:  qreal QGraphicsGridLayout::horizontalSpacing();
func (this *QGraphicsGridLayout) horizontalSpacing(args ...interface{}) () {
  // horizontalSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout17horizontalSpacingEv
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "horizontalSpacing", args)
  }

}

  // proto:  void QGraphicsGridLayout::setColumnStretchFactor(int column, int stretch);
func (this *QGraphicsGridLayout) setColumnStretchFactor(args ...interface{}) () {
  // setColumnStretchFactor(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout22setColumnStretchFactorEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setColumnStretchFactor", args)
  }

}

  // proto:  void QGraphicsGridLayout::invalidate();
func (this *QGraphicsGridLayout) invalidate(args ...interface{}) () {
  // invalidate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout10invalidateEv
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "invalidate", args)
  }

}

  // proto:  qreal QGraphicsGridLayout::columnPreferredWidth(int column);
func (this *QGraphicsGridLayout) columnPreferredWidth(args ...interface{}) () {
  // columnPreferredWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout20columnPreferredWidthEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "columnPreferredWidth", args)
  }

}

  // proto:  void QGraphicsGridLayout::setColumnPreferredWidth(int column, qreal width);
func (this *QGraphicsGridLayout) setColumnPreferredWidth(args ...interface{}) () {
  // setColumnPreferredWidth(int, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout23setColumnPreferredWidthEid
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setColumnPreferredWidth", args)
  }

}

  // proto:  qreal QGraphicsGridLayout::columnSpacing(int column);
func (this *QGraphicsGridLayout) columnSpacing(args ...interface{}) () {
  // columnSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout13columnSpacingEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "columnSpacing", args)
  }

}

  // proto:  void QGraphicsGridLayout::setRowSpacing(int row, qreal spacing);
func (this *QGraphicsGridLayout) setRowSpacing(args ...interface{}) () {
  // setRowSpacing(int, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout13setRowSpacingEid
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setRowSpacing", args)
  }

}

  // proto:  void QGraphicsGridLayout::removeAt(int index);
func (this *QGraphicsGridLayout) removeAt(args ...interface{}) () {
  // removeAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout8removeAtEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "removeAt", args)
  }

}

// <= body block end

