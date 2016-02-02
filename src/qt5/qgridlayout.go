package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtWidgets/qgridlayout.h
// dst-file: /src/widgets/qgridlayout.go
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
  // proto:  int QGridLayout::columnCount();
extern int32_t C_ZNK11QGridLayout11columnCountEv(void* qthis); // 4
  // proto:  void QGridLayout::invalidate();
extern void C_ZN11QGridLayout10invalidateEv(void* qthis); // 4
  // proto:  void QGridLayout::~QGridLayout();
extern void C_ZN11QGridLayoutD2Ev(void* qthis); // 4
  // proto:  QSize QGridLayout::minimumSize();
extern void* C_ZNK11QGridLayout11minimumSizeEv(void* qthis); // 4
  // proto:  int QGridLayout::columnStretch(int column);
extern int32_t C_ZNK11QGridLayout13columnStretchEi(void* qthis, int32_t arg0); // 4
  // proto:  int QGridLayout::verticalSpacing();
extern int32_t C_ZNK11QGridLayout15verticalSpacingEv(void* qthis); // 4
  // proto:  QLayoutItem * QGridLayout::itemAtPosition(int row, int column);
extern void* C_ZNK11QGridLayout14itemAtPositionEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QGridLayout::setRowMinimumHeight(int row, int minSize);
extern void C_ZN11QGridLayout19setRowMinimumHeightEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QGridLayout::setVerticalSpacing(int spacing);
extern void C_ZN11QGridLayout18setVerticalSpacingEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGridLayout::getItemPosition(int idx, int * row, int * column, int * rowSpan, int * columnSpan);
extern void C_ZNK11QGridLayout15getItemPositionEiPiS0_S0_S0_(void* qthis, int32_t arg0, void* arg1, void* arg2, void* arg3, void* arg4); // 4
  // proto:  int QGridLayout::rowStretch(int row);
extern int32_t C_ZNK11QGridLayout10rowStretchEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGridLayout::setColumnStretch(int column, int stretch);
extern void C_ZN11QGridLayout16setColumnStretchEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QGridLayout::QGridLayout(QWidget * parent);
extern void* C_ZN11QGridLayoutC2EP7QWidget(void* arg0); // 3
  // proto:  void QGridLayout::QGridLayout();
extern void* C_ZN11QGridLayoutC2Ev(); // 3
  // proto:  int QGridLayout::rowMinimumHeight(int row);
extern int32_t C_ZNK11QGridLayout16rowMinimumHeightEi(void* qthis, int32_t arg0); // 4
  // proto:  QLayoutItem * QGridLayout::takeAt(int index);
extern void* C_ZN11QGridLayout6takeAtEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGridLayout::setRowStretch(int row, int stretch);
extern void C_ZN11QGridLayout13setRowStretchEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  int QGridLayout::minimumHeightForWidth(int );
extern int32_t C_ZNK11QGridLayout21minimumHeightForWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGridLayout::setSpacing(int spacing);
extern void C_ZN11QGridLayout10setSpacingEi(void* qthis, int32_t arg0); // 4
  // proto:  int QGridLayout::horizontalSpacing();
extern int32_t C_ZNK11QGridLayout17horizontalSpacingEv(void* qthis); // 4
  // proto:  void QGridLayout::setGeometry(const QRect & );
extern void C_ZN11QGridLayout11setGeometryERK5QRect(void* qthis, void* arg0); // 4
  // proto:  int QGridLayout::spacing();
extern int32_t C_ZNK11QGridLayout7spacingEv(void* qthis); // 4
  // proto:  int QGridLayout::rowCount();
extern int32_t C_ZNK11QGridLayout8rowCountEv(void* qthis); // 4
  // proto:  QRect QGridLayout::cellRect(int row, int column);
extern void* C_ZNK11QGridLayout8cellRectEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QGridLayout::addWidget(QWidget * w);
extern void C_ZN11QGridLayout9addWidgetEP7QWidget(void* qthis, void* arg0); // 2
  // proto:  const QMetaObject * QGridLayout::metaObject();
extern void C_ZNK11QGridLayout10metaObjectEv(void* qthis); // 4
  // proto:  int QGridLayout::count();
extern int32_t C_ZNK11QGridLayout5countEv(void* qthis); // 4
  // proto:  QSize QGridLayout::sizeHint();
extern void* C_ZNK11QGridLayout8sizeHintEv(void* qthis); // 4
  // proto:  bool QGridLayout::hasHeightForWidth();
extern bool C_ZNK11QGridLayout17hasHeightForWidthEv(void* qthis); // 4
  // proto:  QLayoutItem * QGridLayout::itemAt(int index);
extern void* C_ZNK11QGridLayout6itemAtEi(void* qthis, int32_t arg0); // 4
  // proto:  int QGridLayout::columnMinimumWidth(int column);
extern int32_t C_ZNK11QGridLayout18columnMinimumWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  Qt::Corner QGridLayout::originCorner();
extern void C_ZNK11QGridLayout12originCornerEv(void* qthis); // 4
  // proto:  void QGridLayout::setColumnMinimumWidth(int column, int minSize);
extern void C_ZN11QGridLayout21setColumnMinimumWidthEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QGridLayout::setHorizontalSpacing(int spacing);
extern void C_ZN11QGridLayout20setHorizontalSpacingEi(void* qthis, int32_t arg0); // 4
  // proto:  int QGridLayout::heightForWidth(int );
extern int32_t C_ZNK11QGridLayout14heightForWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  Qt::Orientations QGridLayout::expandingDirections();
extern void C_ZNK11QGridLayout19expandingDirectionsEv(void* qthis); // 4
  // proto:  QSize QGridLayout::maximumSize();
extern void* C_ZNK11QGridLayout11maximumSizeEv(void* qthis); // 4
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

// class sizeof(QGridLayout)=1
type QGridLayout struct {
  /*qbase*/ QLayout;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// columnCount()
func (this *QGridLayout) Columncount(args ...interface{}) (ret interface{}) {
  // columnCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout11columnCountEv
    // invoke: int columnCount()
    var ret0 = C.C_ZNK11QGridLayout11columnCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGridLayout", "columnCount", args)
  }

  return
}

// invalidate()
func (this *QGridLayout) Invalidate(args ...interface{}) () {
  // invalidate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout10invalidateEv
    // invoke: void invalidate()
    C.C_ZN11QGridLayout10invalidateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGridLayout", "invalidate", args)
  }

  return
}

// ~QGridLayout()
func (this *QGridLayout) Freeqgridlayout(args ...interface{}) () {
  // ~QGridLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayoutD0Ev
    // invoke: void ~QGridLayout()
    C.C_ZN11QGridLayoutD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGridLayout", "~QGridLayout", args)
  }

  return
}

// minimumSize()
func (this *QGridLayout) Minimumsize(args ...interface{}) (ret interface{}) {
  // minimumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout11minimumSizeEv
    // invoke: QSize minimumSize()
    var ret0 = C.C_ZNK11QGridLayout11minimumSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGridLayout", "minimumSize", args)
  }

  return
}

// columnStretch(int)
func (this *QGridLayout) Columnstretch(args ...interface{}) (ret interface{}) {
  // columnStretch(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout13columnStretchEi
    // invoke: int columnStretch(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QGridLayout13columnStretchEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGridLayout", "columnStretch", args)
  }

  return
}

// verticalSpacing()
func (this *QGridLayout) Verticalspacing(args ...interface{}) (ret interface{}) {
  // verticalSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout15verticalSpacingEv
    // invoke: int verticalSpacing()
    var ret0 = C.C_ZNK11QGridLayout15verticalSpacingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGridLayout", "verticalSpacing", args)
  }

  return
}

// itemAtPosition(int, int)
func (this *QGridLayout) Itematposition(args ...interface{}) (ret interface{}) {
  // itemAtPosition(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout14itemAtPositionEii
    // invoke: QLayoutItem * itemAtPosition(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK11QGridLayout14itemAtPositionEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLayoutItem{}) // "QLayoutItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGridLayout", "itemAtPosition", args)
  }

  return
}

// setRowMinimumHeight(int, int)
func (this *QGridLayout) Setrowminimumheight(args ...interface{}) () {
  // setRowMinimumHeight(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout19setRowMinimumHeightEii
    // invoke: void setRowMinimumHeight(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN11QGridLayout19setRowMinimumHeightEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGridLayout", "setRowMinimumHeight", args)
  }

  return
}

// setVerticalSpacing(int)
func (this *QGridLayout) Setverticalspacing(args ...interface{}) () {
  // setVerticalSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout18setVerticalSpacingEi
    // invoke: void setVerticalSpacing(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QGridLayout18setVerticalSpacingEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGridLayout", "setVerticalSpacing", args)
  }

  return
}

// getItemPosition(int, int *, int *, int *, int *)
func (this *QGridLayout) Getitemposition(args ...interface{}) () {
  // getItemPosition(int, int *, int *, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"
  vtys[0][4] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout15getItemPositionEiPiS0_S0_S0_
    // invoke: void getItemPosition(int, int *, int *, int *, int *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    var arg4 = (unsafe.Pointer)(args[4].(*int32))
    if false {fmt.Println(arg4)}
    C.C_ZNK11QGridLayout15getItemPositionEiPiS0_S0_S0_(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
  default:
    qtrt.ErrorResolve("QGridLayout", "getItemPosition", args)
  }

  return
}

// rowStretch(int)
func (this *QGridLayout) Rowstretch(args ...interface{}) (ret interface{}) {
  // rowStretch(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout10rowStretchEi
    // invoke: int rowStretch(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QGridLayout10rowStretchEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGridLayout", "rowStretch", args)
  }

  return
}

// setColumnStretch(int, int)
func (this *QGridLayout) Setcolumnstretch(args ...interface{}) () {
  // setColumnStretch(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout16setColumnStretchEii
    // invoke: void setColumnStretch(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN11QGridLayout16setColumnStretchEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGridLayout", "setColumnStretch", args)
  }

  return
}

// QGridLayout(class QWidget *)
func NewQGridLayout(args ...interface{}) *QGridLayout {
  // QGridLayout(class QWidget *)
  // QGridLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayoutC1EP7QWidget
    // invoke: void QGridLayout(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QGridLayoutC2EP7QWidget(arg0)
    return &QGridLayout{Qclsinst:qthis}
  case 1:
    // invoke: _ZN11QGridLayoutC1Ev
    // invoke: void QGridLayout()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QGridLayoutC2Ev()
    return &QGridLayout{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGridLayout", "QGridLayout", args)
  }

  return nil // QGridLayout{Qclsinst:qthis}
}

// rowMinimumHeight(int)
func (this *QGridLayout) Rowminimumheight(args ...interface{}) (ret interface{}) {
  // rowMinimumHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout16rowMinimumHeightEi
    // invoke: int rowMinimumHeight(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QGridLayout16rowMinimumHeightEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGridLayout", "rowMinimumHeight", args)
  }

  return
}

// takeAt(int)
func (this *QGridLayout) Takeat(args ...interface{}) (ret interface{}) {
  // takeAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout6takeAtEi
    // invoke: QLayoutItem * takeAt(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN11QGridLayout6takeAtEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLayoutItem{}) // "QLayoutItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGridLayout", "takeAt", args)
  }

  return
}

// setRowStretch(int, int)
func (this *QGridLayout) Setrowstretch(args ...interface{}) () {
  // setRowStretch(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout13setRowStretchEii
    // invoke: void setRowStretch(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN11QGridLayout13setRowStretchEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGridLayout", "setRowStretch", args)
  }

  return
}

// minimumHeightForWidth(int)
func (this *QGridLayout) Minimumheightforwidth(args ...interface{}) (ret interface{}) {
  // minimumHeightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout21minimumHeightForWidthEi
    // invoke: int minimumHeightForWidth(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QGridLayout21minimumHeightForWidthEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGridLayout", "minimumHeightForWidth", args)
  }

  return
}

// setSpacing(int)
func (this *QGridLayout) Setspacing(args ...interface{}) () {
  // setSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout10setSpacingEi
    // invoke: void setSpacing(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QGridLayout10setSpacingEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGridLayout", "setSpacing", args)
  }

  return
}

// horizontalSpacing()
func (this *QGridLayout) Horizontalspacing(args ...interface{}) (ret interface{}) {
  // horizontalSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout17horizontalSpacingEv
    // invoke: int horizontalSpacing()
    var ret0 = C.C_ZNK11QGridLayout17horizontalSpacingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGridLayout", "horizontalSpacing", args)
  }

  return
}

// setGeometry(const class QRect &)
func (this *QGridLayout) Setgeometry(args ...interface{}) () {
  // setGeometry(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout11setGeometryERK5QRect
    // invoke: void setGeometry(const class QRect &)
    var arg0 = args[0].(*QRect).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QGridLayout11setGeometryERK5QRect(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGridLayout", "setGeometry", args)
  }

  return
}

// spacing()
func (this *QGridLayout) Spacing(args ...interface{}) (ret interface{}) {
  // spacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout7spacingEv
    // invoke: int spacing()
    var ret0 = C.C_ZNK11QGridLayout7spacingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGridLayout", "spacing", args)
  }

  return
}

// rowCount()
func (this *QGridLayout) Rowcount(args ...interface{}) (ret interface{}) {
  // rowCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout8rowCountEv
    // invoke: int rowCount()
    var ret0 = C.C_ZNK11QGridLayout8rowCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGridLayout", "rowCount", args)
  }

  return
}

// cellRect(int, int)
func (this *QGridLayout) Cellrect(args ...interface{}) (ret interface{}) {
  // cellRect(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout8cellRectEii
    // invoke: QRect cellRect(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK11QGridLayout8cellRectEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGridLayout", "cellRect", args)
  }

  return
}

// addWidget(class QWidget *)
func (this *QGridLayout) Addwidget(args ...interface{}) () {
  // addWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout9addWidgetEP7QWidget
    // invoke: void addWidget(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QGridLayout9addWidgetEP7QWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGridLayout", "addWidget", args)
  }

  return
}

// metaObject()
func (this *QGridLayout) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK11QGridLayout10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGridLayout", "metaObject", args)
  }

  return
}

// count()
func (this *QGridLayout) Count(args ...interface{}) (ret interface{}) {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout5countEv
    // invoke: int count()
    var ret0 = C.C_ZNK11QGridLayout5countEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGridLayout", "count", args)
  }

  return
}

// sizeHint()
func (this *QGridLayout) Sizehint(args ...interface{}) (ret interface{}) {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout8sizeHintEv
    // invoke: QSize sizeHint()
    var ret0 = C.C_ZNK11QGridLayout8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGridLayout", "sizeHint", args)
  }

  return
}

// hasHeightForWidth()
func (this *QGridLayout) Hasheightforwidth(args ...interface{}) (ret interface{}) {
  // hasHeightForWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout17hasHeightForWidthEv
    // invoke: bool hasHeightForWidth()
    var ret0 = C.C_ZNK11QGridLayout17hasHeightForWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGridLayout", "hasHeightForWidth", args)
  }

  return
}

// itemAt(int)
func (this *QGridLayout) Itemat(args ...interface{}) (ret interface{}) {
  // itemAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout6itemAtEi
    // invoke: QLayoutItem * itemAt(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QGridLayout6itemAtEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLayoutItem{}) // "QLayoutItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGridLayout", "itemAt", args)
  }

  return
}

// columnMinimumWidth(int)
func (this *QGridLayout) Columnminimumwidth(args ...interface{}) (ret interface{}) {
  // columnMinimumWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout18columnMinimumWidthEi
    // invoke: int columnMinimumWidth(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QGridLayout18columnMinimumWidthEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGridLayout", "columnMinimumWidth", args)
  }

  return
}

// originCorner()
func (this *QGridLayout) Origincorner(args ...interface{}) () {
  // originCorner()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout12originCornerEv
    // invoke: Qt::Corner originCorner()
    C.C_ZNK11QGridLayout12originCornerEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGridLayout", "originCorner", args)
  }

  return
}

// setColumnMinimumWidth(int, int)
func (this *QGridLayout) Setcolumnminimumwidth(args ...interface{}) () {
  // setColumnMinimumWidth(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout21setColumnMinimumWidthEii
    // invoke: void setColumnMinimumWidth(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN11QGridLayout21setColumnMinimumWidthEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGridLayout", "setColumnMinimumWidth", args)
  }

  return
}

// setHorizontalSpacing(int)
func (this *QGridLayout) Sethorizontalspacing(args ...interface{}) () {
  // setHorizontalSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout20setHorizontalSpacingEi
    // invoke: void setHorizontalSpacing(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QGridLayout20setHorizontalSpacingEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGridLayout", "setHorizontalSpacing", args)
  }

  return
}

// heightForWidth(int)
func (this *QGridLayout) Heightforwidth(args ...interface{}) (ret interface{}) {
  // heightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout14heightForWidthEi
    // invoke: int heightForWidth(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QGridLayout14heightForWidthEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGridLayout", "heightForWidth", args)
  }

  return
}

// expandingDirections()
func (this *QGridLayout) Expandingdirections(args ...interface{}) () {
  // expandingDirections()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout19expandingDirectionsEv
    // invoke: Qt::Orientations expandingDirections()
    C.C_ZNK11QGridLayout19expandingDirectionsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGridLayout", "expandingDirections", args)
  }

  return
}

// maximumSize()
func (this *QGridLayout) Maximumsize(args ...interface{}) (ret interface{}) {
  // maximumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout11maximumSizeEv
    // invoke: QSize maximumSize()
    var ret0 = C.C_ZNK11QGridLayout11maximumSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGridLayout", "maximumSize", args)
  }

  return
}

// <= body block end

