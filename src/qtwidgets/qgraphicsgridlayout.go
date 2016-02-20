package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
import "qtcore"
import "qtgui"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  int QGraphicsGridLayout::columnCount();
extern int32_t C_ZNK19QGraphicsGridLayout11columnCountEv(void* qthis); // 4
  // proto:  void QGraphicsGridLayout::invalidate();
extern void C_ZN19QGraphicsGridLayout10invalidateEv(void* qthis); // 4
  // proto:  qreal QGraphicsGridLayout::rowSpacing(int row);
extern double C_ZNK19QGraphicsGridLayout10rowSpacingEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGraphicsGridLayout::setRowFixedHeight(int row, qreal height);
extern void C_ZN19QGraphicsGridLayout17setRowFixedHeightEid(void* qthis, int32_t arg0, double arg1); // 4
  // proto:  Qt::Alignment QGraphicsGridLayout::columnAlignment(int column);
extern void C_ZNK19QGraphicsGridLayout15columnAlignmentEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGraphicsGridLayout::removeAt(int index);
extern void C_ZN19QGraphicsGridLayout8removeAtEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGraphicsGridLayout::setColumnSpacing(int column, qreal spacing);
extern void C_ZN19QGraphicsGridLayout16setColumnSpacingEid(void* qthis, int32_t arg0, double arg1); // 4
  // proto:  qreal QGraphicsGridLayout::verticalSpacing();
extern double C_ZNK19QGraphicsGridLayout15verticalSpacingEv(void* qthis); // 4
  // proto:  qreal QGraphicsGridLayout::rowMaximumHeight(int row);
extern double C_ZNK19QGraphicsGridLayout16rowMaximumHeightEi(void* qthis, int32_t arg0); // 4
  // proto:  qreal QGraphicsGridLayout::rowPreferredHeight(int row);
extern double C_ZNK19QGraphicsGridLayout18rowPreferredHeightEi(void* qthis, int32_t arg0); // 4
  // proto:  qreal QGraphicsGridLayout::columnMaximumWidth(int column);
extern double C_ZNK19QGraphicsGridLayout18columnMaximumWidthEi(void* qthis, int32_t arg0); // 4
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
extern double C_ZNK19QGraphicsGridLayout16rowMinimumHeightEi(void* qthis, int32_t arg0); // 4
  // proto:  qreal QGraphicsGridLayout::horizontalSpacing();
extern double C_ZNK19QGraphicsGridLayout17horizontalSpacingEv(void* qthis); // 4
  // proto:  int QGraphicsGridLayout::rowStretchFactor(int row);
extern int32_t C_ZNK19QGraphicsGridLayout16rowStretchFactorEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGraphicsGridLayout::setColumnPreferredWidth(int column, qreal width);
extern void C_ZN19QGraphicsGridLayout23setColumnPreferredWidthEid(void* qthis, int32_t arg0, double arg1); // 4
  // proto:  qreal QGraphicsGridLayout::columnSpacing(int column);
extern double C_ZNK19QGraphicsGridLayout13columnSpacingEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGraphicsGridLayout::setColumnMinimumWidth(int column, qreal width);
extern void C_ZN19QGraphicsGridLayout21setColumnMinimumWidthEid(void* qthis, int32_t arg0, double arg1); // 4
  // proto:  int QGraphicsGridLayout::columnStretchFactor(int column);
extern int32_t C_ZNK19QGraphicsGridLayout19columnStretchFactorEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGraphicsGridLayout::setRowSpacing(int row, qreal spacing);
extern void C_ZN19QGraphicsGridLayout13setRowSpacingEid(void* qthis, int32_t arg0, double arg1); // 4
  // proto:  int QGraphicsGridLayout::rowCount();
extern int32_t C_ZNK19QGraphicsGridLayout8rowCountEv(void* qthis); // 4
  // proto:  void QGraphicsGridLayout::removeItem(QGraphicsLayoutItem * item);
extern void C_ZN19QGraphicsGridLayout10removeItemEP19QGraphicsLayoutItem(void* qthis, void* arg0); // 4
  // proto:  int QGraphicsGridLayout::count();
extern int32_t C_ZNK19QGraphicsGridLayout5countEv(void* qthis); // 4
  // proto:  void QGraphicsGridLayout::~QGraphicsGridLayout();
extern void C_ZN19QGraphicsGridLayoutD2Ev(void* qthis); // 4
  // proto:  void QGraphicsGridLayout::setColumnFixedWidth(int column, qreal width);
extern void C_ZN19QGraphicsGridLayout19setColumnFixedWidthEid(void* qthis, int32_t arg0, double arg1); // 4
  // proto:  QGraphicsLayoutItem * QGraphicsGridLayout::itemAt(int row, int column);
extern void C_ZNK19QGraphicsGridLayout6itemAtEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QGraphicsLayoutItem * QGraphicsGridLayout::itemAt(int index);
extern void C_ZNK19QGraphicsGridLayout6itemAtEi(void* qthis, int32_t arg0); // 4
  // proto:  qreal QGraphicsGridLayout::columnMinimumWidth(int column);
extern double C_ZNK19QGraphicsGridLayout18columnMinimumWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGraphicsGridLayout::setHorizontalSpacing(qreal spacing);
extern void C_ZN19QGraphicsGridLayout20setHorizontalSpacingEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsGridLayout::setRowMaximumHeight(int row, qreal height);
extern void C_ZN19QGraphicsGridLayout19setRowMaximumHeightEid(void* qthis, int32_t arg0, double arg1); // 4
  // proto:  qreal QGraphicsGridLayout::columnPreferredWidth(int column);
extern double C_ZNK19QGraphicsGridLayout20columnPreferredWidthEi(void* qthis, int32_t arg0); // 4
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
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QGraphicsGridLayout)=1
type QGraphicsGridLayout struct {
  /*qbase*/ QGraphicsLayout;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// columnCount()
func (this *QGraphicsGridLayout) Columncount(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK19QGraphicsGridLayout11columnCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "columnCount", args)
  }

  return
}

// invalidate()
func (this *QGraphicsGridLayout) Invalidate(args ...interface{}) () {
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
    C.C_ZN19QGraphicsGridLayout10invalidateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "invalidate", args)
  }

  return
}

// rowSpacing(int)
func (this *QGraphicsGridLayout) Rowspacing(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QGraphicsGridLayout10rowSpacingEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "rowSpacing", args)
  }

  return
}

// setRowFixedHeight(int, qreal)
func (this *QGraphicsGridLayout) Setrowfixedheight(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsGridLayout17setRowFixedHeightEid(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setRowFixedHeight", args)
  }

  return
}

// columnAlignment(int)
func (this *QGraphicsGridLayout) Columnalignment(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK19QGraphicsGridLayout15columnAlignmentEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "columnAlignment", args)
  }

  return
}

// removeAt(int)
func (this *QGraphicsGridLayout) Removeat(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsGridLayout8removeAtEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "removeAt", args)
  }

  return
}

// setColumnSpacing(int, qreal)
func (this *QGraphicsGridLayout) Setcolumnspacing(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsGridLayout16setColumnSpacingEid(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setColumnSpacing", args)
  }

  return
}

// verticalSpacing()
func (this *QGraphicsGridLayout) Verticalspacing(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK19QGraphicsGridLayout15verticalSpacingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "verticalSpacing", args)
  }

  return
}

// rowMaximumHeight(int)
func (this *QGraphicsGridLayout) Rowmaximumheight(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QGraphicsGridLayout16rowMaximumHeightEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "rowMaximumHeight", args)
  }

  return
}

// rowPreferredHeight(int)
func (this *QGraphicsGridLayout) Rowpreferredheight(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QGraphicsGridLayout18rowPreferredHeightEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "rowPreferredHeight", args)
  }

  return
}

// columnMaximumWidth(int)
func (this *QGraphicsGridLayout) Columnmaximumwidth(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QGraphicsGridLayout18columnMaximumWidthEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "columnMaximumWidth", args)
  }

  return
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
    var arg0 = args[0].(*QGraphicsLayoutItem).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QGraphicsGridLayoutC2EP19QGraphicsLayoutItem(arg0)
    return &QGraphicsGridLayout{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "QGraphicsGridLayout", args)
  }

  return nil // QGraphicsGridLayout{Qclsinst:qthis}
}

// setRowStretchFactor(int, int)
func (this *QGraphicsGridLayout) Setrowstretchfactor(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsGridLayout19setRowStretchFactorEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setRowStretchFactor", args)
  }

  return
}

// alignment(class QGraphicsLayoutItem *)
func (this *QGraphicsGridLayout) Alignment(args ...interface{}) () {
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
    var arg0 = args[0].(*QGraphicsLayoutItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK19QGraphicsGridLayout9alignmentEP19QGraphicsLayoutItem(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "alignment", args)
  }

  return
}

// rowAlignment(int)
func (this *QGraphicsGridLayout) Rowalignment(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK19QGraphicsGridLayout12rowAlignmentEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "rowAlignment", args)
  }

  return
}

// setSpacing(qreal)
func (this *QGraphicsGridLayout) Setspacing(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsGridLayout10setSpacingEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setSpacing", args)
  }

  return
}

// setGeometry(const class QRectF &)
func (this *QGraphicsGridLayout) Setgeometry(args ...interface{}) () {
  // setGeometry(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout11setGeometryERK6QRectF
    // invoke: void setGeometry(const class QRectF &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsGridLayout11setGeometryERK6QRectF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setGeometry", args)
  }

  return
}

// setColumnStretchFactor(int, int)
func (this *QGraphicsGridLayout) Setcolumnstretchfactor(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsGridLayout22setColumnStretchFactorEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setColumnStretchFactor", args)
  }

  return
}

// setVerticalSpacing(qreal)
func (this *QGraphicsGridLayout) Setverticalspacing(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsGridLayout18setVerticalSpacingEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setVerticalSpacing", args)
  }

  return
}

// rowMinimumHeight(int)
func (this *QGraphicsGridLayout) Rowminimumheight(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QGraphicsGridLayout16rowMinimumHeightEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "rowMinimumHeight", args)
  }

  return
}

// horizontalSpacing()
func (this *QGraphicsGridLayout) Horizontalspacing(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK19QGraphicsGridLayout17horizontalSpacingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "horizontalSpacing", args)
  }

  return
}

// rowStretchFactor(int)
func (this *QGraphicsGridLayout) Rowstretchfactor(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QGraphicsGridLayout16rowStretchFactorEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "rowStretchFactor", args)
  }

  return
}

// setColumnPreferredWidth(int, qreal)
func (this *QGraphicsGridLayout) Setcolumnpreferredwidth(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsGridLayout23setColumnPreferredWidthEid(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setColumnPreferredWidth", args)
  }

  return
}

// columnSpacing(int)
func (this *QGraphicsGridLayout) Columnspacing(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QGraphicsGridLayout13columnSpacingEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "columnSpacing", args)
  }

  return
}

// setColumnMinimumWidth(int, qreal)
func (this *QGraphicsGridLayout) Setcolumnminimumwidth(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsGridLayout21setColumnMinimumWidthEid(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setColumnMinimumWidth", args)
  }

  return
}

// columnStretchFactor(int)
func (this *QGraphicsGridLayout) Columnstretchfactor(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QGraphicsGridLayout19columnStretchFactorEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "columnStretchFactor", args)
  }

  return
}

// setRowSpacing(int, qreal)
func (this *QGraphicsGridLayout) Setrowspacing(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsGridLayout13setRowSpacingEid(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setRowSpacing", args)
  }

  return
}

// rowCount()
func (this *QGraphicsGridLayout) Rowcount(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK19QGraphicsGridLayout8rowCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "rowCount", args)
  }

  return
}

// removeItem(class QGraphicsLayoutItem *)
func (this *QGraphicsGridLayout) Removeitem(args ...interface{}) () {
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
    var arg0 = args[0].(*QGraphicsLayoutItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsGridLayout10removeItemEP19QGraphicsLayoutItem(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "removeItem", args)
  }

  return
}

// count()
func (this *QGraphicsGridLayout) Count(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK19QGraphicsGridLayout5countEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "count", args)
  }

  return
}

// ~QGraphicsGridLayout()
func (this *QGraphicsGridLayout) Freeqgraphicsgridlayout(args ...interface{}) () {
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
    C.C_ZN19QGraphicsGridLayoutD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "~QGraphicsGridLayout", args)
  }

  return
}

// setColumnFixedWidth(int, qreal)
func (this *QGraphicsGridLayout) Setcolumnfixedwidth(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsGridLayout19setColumnFixedWidthEid(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setColumnFixedWidth", args)
  }

  return
}

// itemAt(int, int)
func (this *QGraphicsGridLayout) Itemat(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK19QGraphicsGridLayout6itemAtEii(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK19QGraphicsGridLayout6itemAtEi
    // invoke: QGraphicsLayoutItem * itemAt(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK19QGraphicsGridLayout6itemAtEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "itemAt", args)
  }

  return
}

// columnMinimumWidth(int)
func (this *QGraphicsGridLayout) Columnminimumwidth(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QGraphicsGridLayout18columnMinimumWidthEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "columnMinimumWidth", args)
  }

  return
}

// setHorizontalSpacing(qreal)
func (this *QGraphicsGridLayout) Sethorizontalspacing(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsGridLayout20setHorizontalSpacingEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setHorizontalSpacing", args)
  }

  return
}

// setRowMaximumHeight(int, qreal)
func (this *QGraphicsGridLayout) Setrowmaximumheight(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsGridLayout19setRowMaximumHeightEid(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setRowMaximumHeight", args)
  }

  return
}

// columnPreferredWidth(int)
func (this *QGraphicsGridLayout) Columnpreferredwidth(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QGraphicsGridLayout20columnPreferredWidthEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "columnPreferredWidth", args)
  }

  return
}

// setColumnMaximumWidth(int, qreal)
func (this *QGraphicsGridLayout) Setcolumnmaximumwidth(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsGridLayout21setColumnMaximumWidthEid(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setColumnMaximumWidth", args)
  }

  return
}

// setRowMinimumHeight(int, qreal)
func (this *QGraphicsGridLayout) Setrowminimumheight(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsGridLayout19setRowMinimumHeightEid(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setRowMinimumHeight", args)
  }

  return
}

// setRowPreferredHeight(int, qreal)
func (this *QGraphicsGridLayout) Setrowpreferredheight(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsGridLayout21setRowPreferredHeightEid(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setRowPreferredHeight", args)
  }

  return
}

// <= body block end

