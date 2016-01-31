package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  int QGraphicsGridLayout::columnCount();
extern void C_ZNK19QGraphicsGridLayout11columnCountEv(void* qthis); // 4
  // proto:  void QGraphicsGridLayout::invalidate();
extern void C_ZN19QGraphicsGridLayout10invalidateEv(void* qthis); // 4
  // proto:  qreal QGraphicsGridLayout::rowSpacing(int row);
extern void C_ZNK19QGraphicsGridLayout10rowSpacingEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGraphicsGridLayout::setRowFixedHeight(int row, qreal height);
extern void C_ZN19QGraphicsGridLayout17setRowFixedHeightEid(void* qthis, int32_t arg0, double arg1); // 4
  // proto:  Qt::Alignment QGraphicsGridLayout::columnAlignment(int column);
extern void C_ZNK19QGraphicsGridLayout15columnAlignmentEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGraphicsGridLayout::removeAt(int index);
extern void C_ZN19QGraphicsGridLayout8removeAtEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGraphicsGridLayout::setColumnSpacing(int column, qreal spacing);
extern void C_ZN19QGraphicsGridLayout16setColumnSpacingEid(void* qthis, int32_t arg0, double arg1); // 4
  // proto:  qreal QGraphicsGridLayout::verticalSpacing();
extern void C_ZNK19QGraphicsGridLayout15verticalSpacingEv(void* qthis); // 4
  // proto:  qreal QGraphicsGridLayout::rowMaximumHeight(int row);
extern void C_ZNK19QGraphicsGridLayout16rowMaximumHeightEi(void* qthis, int32_t arg0); // 4
  // proto:  qreal QGraphicsGridLayout::rowPreferredHeight(int row);
extern void C_ZNK19QGraphicsGridLayout18rowPreferredHeightEi(void* qthis, int32_t arg0); // 4
  // proto:  qreal QGraphicsGridLayout::columnMaximumWidth(int column);
extern void C_ZNK19QGraphicsGridLayout18columnMaximumWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGraphicsGridLayout::QGraphicsGridLayout(QGraphicsLayoutItem * parent);
extern void* C_ZN19QGraphicsGridLayoutC2EP19QGraphicsLayoutItem(void* arg0); // 3
  // proto:  void QGraphicsGridLayout::setRowStretchFactor(int row, int stretch);
extern void C_ZN19QGraphicsGridLayout19setRowStretchFactorEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  Qt::Alignment QGraphicsGridLayout::alignment(QGraphicsLayoutItem * item);
extern void C_ZNK19QGraphicsGridLayout9alignmentEP19QGraphicsLayoutItem(void* qthis, void* arg0); // 4
  // proto:  Qt::Alignment QGraphicsGridLayout::rowAlignment(int row);
extern void C_ZNK19QGraphicsGridLayout12rowAlignmentEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGraphicsGridLayout::setSpacing(qreal spacing);
extern void C_ZN19QGraphicsGridLayout10setSpacingEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsGridLayout::setGeometry(const QRectF & rect);
extern void C_ZN19QGraphicsGridLayout11setGeometryERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsGridLayout::setColumnStretchFactor(int column, int stretch);
extern void C_ZN19QGraphicsGridLayout22setColumnStretchFactorEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QGraphicsGridLayout::setVerticalSpacing(qreal spacing);
extern void C_ZN19QGraphicsGridLayout18setVerticalSpacingEd(void* qthis, double arg0); // 4
  // proto:  qreal QGraphicsGridLayout::rowMinimumHeight(int row);
extern void C_ZNK19QGraphicsGridLayout16rowMinimumHeightEi(void* qthis, int32_t arg0); // 4
  // proto:  qreal QGraphicsGridLayout::horizontalSpacing();
extern void C_ZNK19QGraphicsGridLayout17horizontalSpacingEv(void* qthis); // 4
  // proto:  int QGraphicsGridLayout::rowStretchFactor(int row);
extern void C_ZNK19QGraphicsGridLayout16rowStretchFactorEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGraphicsGridLayout::setColumnPreferredWidth(int column, qreal width);
extern void C_ZN19QGraphicsGridLayout23setColumnPreferredWidthEid(void* qthis, int32_t arg0, double arg1); // 4
  // proto:  qreal QGraphicsGridLayout::columnSpacing(int column);
extern void C_ZNK19QGraphicsGridLayout13columnSpacingEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGraphicsGridLayout::setColumnMinimumWidth(int column, qreal width);
extern void C_ZN19QGraphicsGridLayout21setColumnMinimumWidthEid(void* qthis, int32_t arg0, double arg1); // 4
  // proto:  int QGraphicsGridLayout::columnStretchFactor(int column);
extern void C_ZNK19QGraphicsGridLayout19columnStretchFactorEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGraphicsGridLayout::setRowSpacing(int row, qreal spacing);
extern void C_ZN19QGraphicsGridLayout13setRowSpacingEid(void* qthis, int32_t arg0, double arg1); // 4
  // proto:  int QGraphicsGridLayout::rowCount();
extern void C_ZNK19QGraphicsGridLayout8rowCountEv(void* qthis); // 4
  // proto:  void QGraphicsGridLayout::removeItem(QGraphicsLayoutItem * item);
extern void C_ZN19QGraphicsGridLayout10removeItemEP19QGraphicsLayoutItem(void* qthis, void* arg0); // 4
  // proto:  int QGraphicsGridLayout::count();
extern void C_ZNK19QGraphicsGridLayout5countEv(void* qthis); // 4
  // proto:  void QGraphicsGridLayout::~QGraphicsGridLayout();
extern void C_ZN19QGraphicsGridLayoutD2Ev(void* qthis); // 4
  // proto:  void QGraphicsGridLayout::setColumnFixedWidth(int column, qreal width);
extern void C_ZN19QGraphicsGridLayout19setColumnFixedWidthEid(void* qthis, int32_t arg0, double arg1); // 4
  // proto:  QGraphicsLayoutItem * QGraphicsGridLayout::itemAt(int row, int column);
extern void C_ZNK19QGraphicsGridLayout6itemAtEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QGraphicsLayoutItem * QGraphicsGridLayout::itemAt(int index);
extern void C_ZNK19QGraphicsGridLayout6itemAtEi(void* qthis, int32_t arg0); // 4
  // proto:  qreal QGraphicsGridLayout::columnMinimumWidth(int column);
extern void C_ZNK19QGraphicsGridLayout18columnMinimumWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGraphicsGridLayout::setHorizontalSpacing(qreal spacing);
extern void C_ZN19QGraphicsGridLayout20setHorizontalSpacingEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsGridLayout::setRowMaximumHeight(int row, qreal height);
extern void C_ZN19QGraphicsGridLayout19setRowMaximumHeightEid(void* qthis, int32_t arg0, double arg1); // 4
  // proto:  qreal QGraphicsGridLayout::columnPreferredWidth(int column);
extern void C_ZNK19QGraphicsGridLayout20columnPreferredWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGraphicsGridLayout::setColumnMaximumWidth(int column, qreal width);
extern void C_ZN19QGraphicsGridLayout21setColumnMaximumWidthEid(void* qthis, int32_t arg0, double arg1); // 4
  // proto:  void QGraphicsGridLayout::setRowMinimumHeight(int row, qreal height);
extern void C_ZN19QGraphicsGridLayout19setRowMinimumHeightEid(void* qthis, int32_t arg0, double arg1); // 4
  // proto:  void QGraphicsGridLayout::setRowPreferredHeight(int row, qreal height);
extern void C_ZN19QGraphicsGridLayout21setRowPreferredHeightEid(void* qthis, int32_t arg0, double arg1); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// columnCount()
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
    // invoke: int columnCount()
    var ret = C.C_ZNK19QGraphicsGridLayout11columnCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "columnCount", args)
  }

}

// invalidate()
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
    // invoke: void invalidate()
    C.C_ZN19QGraphicsGridLayout10invalidateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "invalidate", args)
  }

}

// rowSpacing(int)
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
    // invoke: qreal rowSpacing(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK19QGraphicsGridLayout10rowSpacingEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "rowSpacing", args)
  }

}

// setRowFixedHeight(int, qreal)
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
    // invoke: void setRowFixedHeight(int, qreal)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsGridLayout17setRowFixedHeightEid(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setRowFixedHeight", args)
  }

}

// columnAlignment(int)
func (this *QGraphicsGridLayout) columnAlignment(args ...interface{}) () {
  // columnAlignment(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout15columnAlignmentEi
    // invoke: Qt::Alignment columnAlignment(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK19QGraphicsGridLayout15columnAlignmentEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "columnAlignment", args)
  }

}

// removeAt(int)
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
    // invoke: void removeAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsGridLayout8removeAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "removeAt", args)
  }

}

// setColumnSpacing(int, qreal)
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
    // invoke: void setColumnSpacing(int, qreal)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsGridLayout16setColumnSpacingEid(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setColumnSpacing", args)
  }

}

// verticalSpacing()
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
    // invoke: qreal verticalSpacing()
    var ret = C.C_ZNK19QGraphicsGridLayout15verticalSpacingEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "verticalSpacing", args)
  }

}

// rowMaximumHeight(int)
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
    // invoke: qreal rowMaximumHeight(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK19QGraphicsGridLayout16rowMaximumHeightEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "rowMaximumHeight", args)
  }

}

// rowPreferredHeight(int)
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
    // invoke: qreal rowPreferredHeight(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK19QGraphicsGridLayout18rowPreferredHeightEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "rowPreferredHeight", args)
  }

}

// columnMaximumWidth(int)
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
    // invoke: qreal columnMaximumWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK19QGraphicsGridLayout18columnMaximumWidthEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "columnMaximumWidth", args)
  }

}

// QGraphicsGridLayout(class QGraphicsLayoutItem *)
func NewQGraphicsGridLayout(args ...interface{}) *QGraphicsGridLayout {
  // QGraphicsGridLayout(class QGraphicsLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsLayoutItem{}) // "QGraphicsLayoutItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayoutC1EP19QGraphicsLayoutItem
    // invoke: void QGraphicsGridLayout(class QGraphicsLayoutItem *)
    var arg0 = args[0].(QGraphicsLayoutItem).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QGraphicsGridLayoutC2EP19QGraphicsLayoutItem(arg0)
    return &QGraphicsGridLayout{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "QGraphicsGridLayout", args)
  }

  return nil // QGraphicsGridLayout{qclsinst:qthis}
}

// setRowStretchFactor(int, int)
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
    // invoke: void setRowStretchFactor(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsGridLayout19setRowStretchFactorEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setRowStretchFactor", args)
  }

}

// alignment(class QGraphicsLayoutItem *)
func (this *QGraphicsGridLayout) alignment(args ...interface{}) () {
  // alignment(class QGraphicsLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsLayoutItem{}) // "QGraphicsLayoutItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout9alignmentEP19QGraphicsLayoutItem
    // invoke: Qt::Alignment alignment(class QGraphicsLayoutItem *)
    var arg0 = args[0].(QGraphicsLayoutItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK19QGraphicsGridLayout9alignmentEP19QGraphicsLayoutItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "alignment", args)
  }

}

// rowAlignment(int)
func (this *QGraphicsGridLayout) rowAlignment(args ...interface{}) () {
  // rowAlignment(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout12rowAlignmentEi
    // invoke: Qt::Alignment rowAlignment(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK19QGraphicsGridLayout12rowAlignmentEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "rowAlignment", args)
  }

}

// setSpacing(qreal)
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
    // invoke: void setSpacing(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsGridLayout10setSpacingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setSpacing", args)
  }

}

// setGeometry(const class QRectF &)
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
    // invoke: void setGeometry(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsGridLayout11setGeometryERK6QRectF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setGeometry", args)
  }

}

// setColumnStretchFactor(int, int)
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
    // invoke: void setColumnStretchFactor(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsGridLayout22setColumnStretchFactorEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setColumnStretchFactor", args)
  }

}

// setVerticalSpacing(qreal)
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
    // invoke: void setVerticalSpacing(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsGridLayout18setVerticalSpacingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setVerticalSpacing", args)
  }

}

// rowMinimumHeight(int)
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
    // invoke: qreal rowMinimumHeight(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK19QGraphicsGridLayout16rowMinimumHeightEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "rowMinimumHeight", args)
  }

}

// horizontalSpacing()
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
    // invoke: qreal horizontalSpacing()
    var ret = C.C_ZNK19QGraphicsGridLayout17horizontalSpacingEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "horizontalSpacing", args)
  }

}

// rowStretchFactor(int)
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
    // invoke: int rowStretchFactor(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK19QGraphicsGridLayout16rowStretchFactorEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "rowStretchFactor", args)
  }

}

// setColumnPreferredWidth(int, qreal)
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
    // invoke: void setColumnPreferredWidth(int, qreal)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsGridLayout23setColumnPreferredWidthEid(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setColumnPreferredWidth", args)
  }

}

// columnSpacing(int)
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
    // invoke: qreal columnSpacing(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK19QGraphicsGridLayout13columnSpacingEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "columnSpacing", args)
  }

}

// setColumnMinimumWidth(int, qreal)
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
    // invoke: void setColumnMinimumWidth(int, qreal)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsGridLayout21setColumnMinimumWidthEid(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setColumnMinimumWidth", args)
  }

}

// columnStretchFactor(int)
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
    // invoke: int columnStretchFactor(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK19QGraphicsGridLayout19columnStretchFactorEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "columnStretchFactor", args)
  }

}

// setRowSpacing(int, qreal)
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
    // invoke: void setRowSpacing(int, qreal)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsGridLayout13setRowSpacingEid(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setRowSpacing", args)
  }

}

// rowCount()
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
    // invoke: int rowCount()
    var ret = C.C_ZNK19QGraphicsGridLayout8rowCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "rowCount", args)
  }

}

// removeItem(class QGraphicsLayoutItem *)
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
    // invoke: void removeItem(class QGraphicsLayoutItem *)
    var arg0 = args[0].(QGraphicsLayoutItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsGridLayout10removeItemEP19QGraphicsLayoutItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "removeItem", args)
  }

}

// count()
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
    // invoke: int count()
    var ret = C.C_ZNK19QGraphicsGridLayout5countEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "count", args)
  }

}

// ~QGraphicsGridLayout()
func (this *QGraphicsGridLayout) FreeQGraphicsGridLayout(args ...interface{}) () {
  // ~QGraphicsGridLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayoutD0Ev
    // invoke: void ~QGraphicsGridLayout()
    C.C_ZN19QGraphicsGridLayoutD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "~QGraphicsGridLayout", args)
  }

}

// setColumnFixedWidth(int, qreal)
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
    // invoke: void setColumnFixedWidth(int, qreal)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsGridLayout19setColumnFixedWidthEid(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setColumnFixedWidth", args)
  }

}

// itemAt(int, int)
func (this *QGraphicsGridLayout) itemAt(args ...interface{}) () {
  // itemAt(int, int)
  // itemAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout6itemAtEii
    // invoke: QGraphicsLayoutItem * itemAt(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK19QGraphicsGridLayout6itemAtEii(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK19QGraphicsGridLayout6itemAtEi
    // invoke: QGraphicsLayoutItem * itemAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK19QGraphicsGridLayout6itemAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "itemAt", args)
  }

}

// columnMinimumWidth(int)
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
    // invoke: qreal columnMinimumWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK19QGraphicsGridLayout18columnMinimumWidthEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "columnMinimumWidth", args)
  }

}

// setHorizontalSpacing(qreal)
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
    // invoke: void setHorizontalSpacing(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsGridLayout20setHorizontalSpacingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setHorizontalSpacing", args)
  }

}

// setRowMaximumHeight(int, qreal)
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
    // invoke: void setRowMaximumHeight(int, qreal)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsGridLayout19setRowMaximumHeightEid(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setRowMaximumHeight", args)
  }

}

// columnPreferredWidth(int)
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
    // invoke: qreal columnPreferredWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK19QGraphicsGridLayout20columnPreferredWidthEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "columnPreferredWidth", args)
  }

}

// setColumnMaximumWidth(int, qreal)
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
    // invoke: void setColumnMaximumWidth(int, qreal)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsGridLayout21setColumnMaximumWidthEid(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setColumnMaximumWidth", args)
  }

}

// setRowMinimumHeight(int, qreal)
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
    // invoke: void setRowMinimumHeight(int, qreal)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsGridLayout19setRowMinimumHeightEid(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setRowMinimumHeight", args)
  }

}

// setRowPreferredHeight(int, qreal)
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
    // invoke: void setRowPreferredHeight(int, qreal)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsGridLayout21setRowPreferredHeightEid(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setRowPreferredHeight", args)
  }

}

// <= body block end

