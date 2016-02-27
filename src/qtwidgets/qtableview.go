package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtWidgets/qtableview.h
// dst-file: /src/widgets/qtableview.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "runtime"
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
  // proto:  bool QTableView::showGrid();
extern bool C_ZNK10QTableView8showGridEv(void* qthis); // 4
  // proto:  void QTableView::setRowHeight(int row, int height);
extern void C_ZN10QTableView12setRowHeightEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  bool QTableView::isSortingEnabled();
extern bool C_ZNK10QTableView16isSortingEnabledEv(void* qthis); // 4
  // proto:  int QTableView::rowSpan(int row, int column);
extern int32_t C_ZNK10QTableView7rowSpanEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QTableView::setShowGrid(bool show);
extern void C_ZN10QTableView11setShowGridEb(void* qthis, bool arg0); // 4
  // proto:  void QTableView::showRow(int row);
extern void C_ZN10QTableView7showRowEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTableView::setColumnWidth(int column, int width);
extern void C_ZN10QTableView14setColumnWidthEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QTableView::sortByColumn(int column);
extern void C_ZN10QTableView12sortByColumnEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTableView::setHorizontalHeader(QHeaderView * header);
extern void C_ZN10QTableView19setHorizontalHeaderEP11QHeaderView(void* qthis, void* arg0); // 4
  // proto:  void QTableView::resizeColumnToContents(int column);
extern void C_ZN10QTableView22resizeColumnToContentsEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTableView::resizeColumnsToContents();
extern void C_ZN10QTableView23resizeColumnsToContentsEv(void* qthis); // 4
  // proto:  int QTableView::columnAt(int x);
extern int32_t C_ZNK10QTableView8columnAtEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTableView::rowViewportPosition(int row);
extern int32_t C_ZNK10QTableView19rowViewportPositionEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QTableView::isCornerButtonEnabled();
extern bool C_ZNK10QTableView21isCornerButtonEnabledEv(void* qthis); // 4
  // proto:  void QTableView::hideColumn(int column);
extern void C_ZN10QTableView10hideColumnEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTableView::setColumnHidden(int column, bool hide);
extern void C_ZN10QTableView15setColumnHiddenEib(void* qthis, int32_t arg0, bool arg1); // 4
  // proto:  void QTableView::setSpan(int row, int column, int rowSpan, int columnSpan);
extern void C_ZN10QTableView7setSpanEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  QHeaderView * QTableView::verticalHeader();
extern void* C_ZNK10QTableView14verticalHeaderEv(void* qthis); // 4
  // proto:  bool QTableView::isColumnHidden(int column);
extern bool C_ZNK10QTableView14isColumnHiddenEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTableView::~QTableView();
extern void C_ZN10QTableViewD2Ev(void* qthis); // 4
  // proto:  void QTableView::QTableView(QWidget * parent);
extern void* C_ZN10QTableViewC2EP7QWidget(void* arg0); // 3
  // proto:  QRect QTableView::visualRect(const QModelIndex & index);
extern void* C_ZNK10QTableView10visualRectERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QTableView::setWordWrap(bool on);
extern void C_ZN10QTableView11setWordWrapEb(void* qthis, bool arg0); // 4
  // proto:  QModelIndex QTableView::indexAt(const QPoint & p);
extern void* C_ZNK10QTableView7indexAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  bool QTableView::isRowHidden(int row);
extern bool C_ZNK10QTableView11isRowHiddenEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTableView::clearSpans();
extern void C_ZN10QTableView10clearSpansEv(void* qthis); // 4
  // proto:  int QTableView::columnViewportPosition(int column);
extern int32_t C_ZNK10QTableView22columnViewportPositionEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTableView::rowAt(int y);
extern int32_t C_ZNK10QTableView5rowAtEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTableView::setRootIndex(const QModelIndex & index);
extern void C_ZN10QTableView12setRootIndexERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QHeaderView * QTableView::horizontalHeader();
extern void* C_ZNK10QTableView16horizontalHeaderEv(void* qthis); // 4
  // proto:  int QTableView::columnWidth(int column);
extern int32_t C_ZNK10QTableView11columnWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTableView::setRowHidden(int row, bool hide);
extern void C_ZN10QTableView12setRowHiddenEib(void* qthis, int32_t arg0, bool arg1); // 4
  // proto:  void QTableView::doItemsLayout();
extern void C_ZN10QTableView13doItemsLayoutEv(void* qthis); // 4
  // proto:  int QTableView::columnSpan(int row, int column);
extern int32_t C_ZNK10QTableView10columnSpanEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QTableView::setModel(QAbstractItemModel * model);
extern void C_ZN10QTableView8setModelEP18QAbstractItemModel(void* qthis, void* arg0); // 4
  // proto:  void QTableView::setSelectionModel(QItemSelectionModel * selectionModel);
extern void C_ZN10QTableView17setSelectionModelEP19QItemSelectionModel(void* qthis, void* arg0); // 4
  // proto:  int QTableView::rowHeight(int row);
extern int32_t C_ZNK10QTableView9rowHeightEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTableView::hideRow(int row);
extern void C_ZN10QTableView7hideRowEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTableView::selectRow(int row);
extern void C_ZN10QTableView9selectRowEi(void* qthis, int32_t arg0); // 4
  // proto:  Qt::PenStyle QTableView::gridStyle();
extern void C_ZNK10QTableView9gridStyleEv(void* qthis); // 4
  // proto:  void QTableView::setCornerButtonEnabled(bool enable);
extern void C_ZN10QTableView22setCornerButtonEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QTableView::setVerticalHeader(QHeaderView * header);
extern void C_ZN10QTableView17setVerticalHeaderEP11QHeaderView(void* qthis, void* arg0); // 4
  // proto:  void QTableView::showColumn(int column);
extern void C_ZN10QTableView10showColumnEi(void* qthis, int32_t arg0); // 4
  // proto:  const QMetaObject * QTableView::metaObject();
extern void C_ZNK10QTableView10metaObjectEv(void* qthis); // 4
  // proto:  void QTableView::selectColumn(int column);
extern void C_ZN10QTableView12selectColumnEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTableView::setSortingEnabled(bool enable);
extern void C_ZN10QTableView17setSortingEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QTableView::resizeRowToContents(int row);
extern void C_ZN10QTableView19resizeRowToContentsEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTableView::resizeRowsToContents();
extern void C_ZN10QTableView20resizeRowsToContentsEv(void* qthis); // 4
  // proto:  bool QTableView::wordWrap();
extern bool C_ZNK10QTableView8wordWrapEv(void* qthis); // 4
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
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QTableView)=1
type QTableView struct {
  /*qbase*/ QAbstractItemView;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// showGrid()
func (this *QTableView) ShowGrid(args ...interface{}) (ret interface{}) {
  // showGrid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView8showGridEv
    // invoke: bool showGrid()
    var ret0 = C.C_ZNK10QTableView8showGridEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableView", "showGrid", args)
  }

  return
}

// setRowHeight(int, int)
func (this *QTableView) SetRowHeight(args ...interface{}) () {
  // setRowHeight(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView12setRowHeightEii
    // invoke: void setRowHeight(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN10QTableView12setRowHeightEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTableView", "setRowHeight", args)
  }

  return
}

// isSortingEnabled()
func (this *QTableView) IsSortingEnabled(args ...interface{}) (ret interface{}) {
  // isSortingEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView16isSortingEnabledEv
    // invoke: bool isSortingEnabled()
    var ret0 = C.C_ZNK10QTableView16isSortingEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableView", "isSortingEnabled", args)
  }

  return
}

// rowSpan(int, int)
func (this *QTableView) RowSpan(args ...interface{}) (ret interface{}) {
  // rowSpan(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView7rowSpanEii
    // invoke: int rowSpan(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QTableView7rowSpanEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableView", "rowSpan", args)
  }

  return
}

// setShowGrid(_Bool)
func (this *QTableView) SetShowGrid(args ...interface{}) () {
  // setShowGrid(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView11setShowGridEb
    // invoke: void setShowGrid(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTableView11setShowGridEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "setShowGrid", args)
  }

  return
}

// showRow(int)
func (this *QTableView) ShowRow(args ...interface{}) () {
  // showRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView7showRowEi
    // invoke: void showRow(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTableView7showRowEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "showRow", args)
  }

  return
}

// setColumnWidth(int, int)
func (this *QTableView) SetColumnWidth(args ...interface{}) () {
  // setColumnWidth(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView14setColumnWidthEii
    // invoke: void setColumnWidth(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN10QTableView14setColumnWidthEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTableView", "setColumnWidth", args)
  }

  return
}

// sortByColumn(int)
func (this *QTableView) SortByColumn(args ...interface{}) () {
  // sortByColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView12sortByColumnEi
    // invoke: void sortByColumn(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTableView12sortByColumnEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "sortByColumn", args)
  }

  return
}

// setHorizontalHeader(class QHeaderView *)
func (this *QTableView) SetHorizontalHeader(args ...interface{}) () {
  // setHorizontalHeader(class QHeaderView *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QHeaderView{}) // "QHeaderView *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView19setHorizontalHeaderEP11QHeaderView
    // invoke: void setHorizontalHeader(class QHeaderView *)
    var arg0 = args[0].(*QHeaderView).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QTableView19setHorizontalHeaderEP11QHeaderView(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "setHorizontalHeader", args)
  }

  return
}

// resizeColumnToContents(int)
func (this *QTableView) ResizeColumnToContents(args ...interface{}) () {
  // resizeColumnToContents(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView22resizeColumnToContentsEi
    // invoke: void resizeColumnToContents(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTableView22resizeColumnToContentsEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "resizeColumnToContents", args)
  }

  return
}

// resizeColumnsToContents()
func (this *QTableView) ResizeColumnsToContents(args ...interface{}) () {
  // resizeColumnsToContents()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView23resizeColumnsToContentsEv
    // invoke: void resizeColumnsToContents()
    C.C_ZN10QTableView23resizeColumnsToContentsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTableView", "resizeColumnsToContents", args)
  }

  return
}

// columnAt(int)
func (this *QTableView) ColumnAt(args ...interface{}) (ret interface{}) {
  // columnAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView8columnAtEi
    // invoke: int columnAt(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTableView8columnAtEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableView", "columnAt", args)
  }

  return
}

// rowViewportPosition(int)
func (this *QTableView) RowViewportPosition(args ...interface{}) (ret interface{}) {
  // rowViewportPosition(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView19rowViewportPositionEi
    // invoke: int rowViewportPosition(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTableView19rowViewportPositionEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableView", "rowViewportPosition", args)
  }

  return
}

// isCornerButtonEnabled()
func (this *QTableView) IsCornerButtonEnabled(args ...interface{}) (ret interface{}) {
  // isCornerButtonEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView21isCornerButtonEnabledEv
    // invoke: bool isCornerButtonEnabled()
    var ret0 = C.C_ZNK10QTableView21isCornerButtonEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableView", "isCornerButtonEnabled", args)
  }

  return
}

// hideColumn(int)
func (this *QTableView) HideColumn(args ...interface{}) () {
  // hideColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView10hideColumnEi
    // invoke: void hideColumn(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTableView10hideColumnEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "hideColumn", args)
  }

  return
}

// setColumnHidden(int, _Bool)
func (this *QTableView) SetColumnHidden(args ...interface{}) () {
  // setColumnHidden(int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView15setColumnHiddenEib
    // invoke: void setColumnHidden(int, _Bool)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN10QTableView15setColumnHiddenEib(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTableView", "setColumnHidden", args)
  }

  return
}

// setSpan(int, int, int, int)
func (this *QTableView) SetSpan(args ...interface{}) () {
  // setSpan(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView7setSpanEiiii
    // invoke: void setSpan(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN10QTableView7setSpanEiiii(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QTableView", "setSpan", args)
  }

  return
}

// verticalHeader()
func (this *QTableView) VerticalHeader(args ...interface{}) (ret interface{}) {
  // verticalHeader()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView14verticalHeaderEv
    // invoke: QHeaderView * verticalHeader()
    var ret0 = C.C_ZNK10QTableView14verticalHeaderEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QHeaderView{}) // "QHeaderView *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableView", "verticalHeader", args)
  }

  return
}

// isColumnHidden(int)
func (this *QTableView) IsColumnHidden(args ...interface{}) (ret interface{}) {
  // isColumnHidden(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView14isColumnHiddenEi
    // invoke: bool isColumnHidden(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTableView14isColumnHiddenEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableView", "isColumnHidden", args)
  }

  return
}

// ~QTableView()
func (this *QTableView) Free(args ...interface{}) () {
  // ~QTableView()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableViewD0Ev
    // invoke: void ~QTableView()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN10QTableViewD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QTableView", "~QTableView", args)
  }

  return
}

// QTableView(class QWidget *)
func GcfreeQTableView(this *QTableView) {
  qtrt.UniverseFree(this)
}
func NewQTableView(args ...interface{}) *QTableView {
  // QTableView(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableViewC1EP7QWidget
    // invoke: void QTableView(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QTableViewC2EP7QWidget(arg0)
    this := &QTableView{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQTableView)
    return this
  default:
    qtrt.ErrorResolve("QTableView", "QTableView", args)
  }

  return nil // QTableView{Qclsinst:qthis}
}

// visualRect(const class QModelIndex &)
func (this *QTableView) VisualRect(args ...interface{}) (ret interface{}) {
  // visualRect(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView10visualRectERK11QModelIndex
    // invoke: QRect visualRect(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTableView10visualRectERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableView", "visualRect", args)
  }

  return
}

// setWordWrap(_Bool)
func (this *QTableView) SetWordWrap(args ...interface{}) () {
  // setWordWrap(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView11setWordWrapEb
    // invoke: void setWordWrap(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTableView11setWordWrapEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "setWordWrap", args)
  }

  return
}

// indexAt(const class QPoint &)
func (this *QTableView) IndexAt(args ...interface{}) (ret interface{}) {
  // indexAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView7indexAtERK6QPoint
    // invoke: QModelIndex indexAt(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTableView7indexAtERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableView", "indexAt", args)
  }

  return
}

// isRowHidden(int)
func (this *QTableView) IsRowHidden(args ...interface{}) (ret interface{}) {
  // isRowHidden(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView11isRowHiddenEi
    // invoke: bool isRowHidden(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTableView11isRowHiddenEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableView", "isRowHidden", args)
  }

  return
}

// clearSpans()
func (this *QTableView) ClearSpans(args ...interface{}) () {
  // clearSpans()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView10clearSpansEv
    // invoke: void clearSpans()
    C.C_ZN10QTableView10clearSpansEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTableView", "clearSpans", args)
  }

  return
}

// columnViewportPosition(int)
func (this *QTableView) ColumnViewportPosition(args ...interface{}) (ret interface{}) {
  // columnViewportPosition(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView22columnViewportPositionEi
    // invoke: int columnViewportPosition(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTableView22columnViewportPositionEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableView", "columnViewportPosition", args)
  }

  return
}

// rowAt(int)
func (this *QTableView) RowAt(args ...interface{}) (ret interface{}) {
  // rowAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView5rowAtEi
    // invoke: int rowAt(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTableView5rowAtEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableView", "rowAt", args)
  }

  return
}

// setRootIndex(const class QModelIndex &)
func (this *QTableView) SetRootIndex(args ...interface{}) () {
  // setRootIndex(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView12setRootIndexERK11QModelIndex
    // invoke: void setRootIndex(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QTableView12setRootIndexERK11QModelIndex(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "setRootIndex", args)
  }

  return
}

// horizontalHeader()
func (this *QTableView) HorizontalHeader(args ...interface{}) (ret interface{}) {
  // horizontalHeader()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView16horizontalHeaderEv
    // invoke: QHeaderView * horizontalHeader()
    var ret0 = C.C_ZNK10QTableView16horizontalHeaderEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QHeaderView{}) // "QHeaderView *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableView", "horizontalHeader", args)
  }

  return
}

// columnWidth(int)
func (this *QTableView) ColumnWidth(args ...interface{}) (ret interface{}) {
  // columnWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView11columnWidthEi
    // invoke: int columnWidth(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTableView11columnWidthEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableView", "columnWidth", args)
  }

  return
}

// setRowHidden(int, _Bool)
func (this *QTableView) SetRowHidden(args ...interface{}) () {
  // setRowHidden(int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView12setRowHiddenEib
    // invoke: void setRowHidden(int, _Bool)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN10QTableView12setRowHiddenEib(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTableView", "setRowHidden", args)
  }

  return
}

// doItemsLayout()
func (this *QTableView) DoItemsLayout(args ...interface{}) () {
  // doItemsLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView13doItemsLayoutEv
    // invoke: void doItemsLayout()
    C.C_ZN10QTableView13doItemsLayoutEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTableView", "doItemsLayout", args)
  }

  return
}

// columnSpan(int, int)
func (this *QTableView) ColumnSpan(args ...interface{}) (ret interface{}) {
  // columnSpan(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView10columnSpanEii
    // invoke: int columnSpan(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QTableView10columnSpanEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableView", "columnSpan", args)
  }

  return
}

// setModel(class QAbstractItemModel *)
func (this *QTableView) SetModel(args ...interface{}) () {
  // setModel(class QAbstractItemModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QAbstractItemModel{}) // "QAbstractItemModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView8setModelEP18QAbstractItemModel
    // invoke: void setModel(class QAbstractItemModel *)
    var arg0 = args[0].(*qtcore.QAbstractItemModel).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QTableView8setModelEP18QAbstractItemModel(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "setModel", args)
  }

  return
}

// setSelectionModel(class QItemSelectionModel *)
func (this *QTableView) SetSelectionModel(args ...interface{}) () {
  // setSelectionModel(class QItemSelectionModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QItemSelectionModel{}) // "QItemSelectionModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView17setSelectionModelEP19QItemSelectionModel
    // invoke: void setSelectionModel(class QItemSelectionModel *)
    var arg0 = args[0].(*qtcore.QItemSelectionModel).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QTableView17setSelectionModelEP19QItemSelectionModel(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "setSelectionModel", args)
  }

  return
}

// rowHeight(int)
func (this *QTableView) RowHeight(args ...interface{}) (ret interface{}) {
  // rowHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView9rowHeightEi
    // invoke: int rowHeight(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTableView9rowHeightEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableView", "rowHeight", args)
  }

  return
}

// hideRow(int)
func (this *QTableView) HideRow(args ...interface{}) () {
  // hideRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView7hideRowEi
    // invoke: void hideRow(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTableView7hideRowEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "hideRow", args)
  }

  return
}

// selectRow(int)
func (this *QTableView) SelectRow(args ...interface{}) () {
  // selectRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView9selectRowEi
    // invoke: void selectRow(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTableView9selectRowEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "selectRow", args)
  }

  return
}

// gridStyle()
func (this *QTableView) GridStyle(args ...interface{}) () {
  // gridStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView9gridStyleEv
    // invoke: Qt::PenStyle gridStyle()
    C.C_ZNK10QTableView9gridStyleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTableView", "gridStyle", args)
  }

  return
}

// setCornerButtonEnabled(_Bool)
func (this *QTableView) SetCornerButtonEnabled(args ...interface{}) () {
  // setCornerButtonEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView22setCornerButtonEnabledEb
    // invoke: void setCornerButtonEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTableView22setCornerButtonEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "setCornerButtonEnabled", args)
  }

  return
}

// setVerticalHeader(class QHeaderView *)
func (this *QTableView) SetVerticalHeader(args ...interface{}) () {
  // setVerticalHeader(class QHeaderView *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QHeaderView{}) // "QHeaderView *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView17setVerticalHeaderEP11QHeaderView
    // invoke: void setVerticalHeader(class QHeaderView *)
    var arg0 = args[0].(*QHeaderView).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QTableView17setVerticalHeaderEP11QHeaderView(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "setVerticalHeader", args)
  }

  return
}

// showColumn(int)
func (this *QTableView) ShowColumn(args ...interface{}) () {
  // showColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView10showColumnEi
    // invoke: void showColumn(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTableView10showColumnEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "showColumn", args)
  }

  return
}

// metaObject()
func (this *QTableView) MetaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK10QTableView10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTableView", "metaObject", args)
  }

  return
}

// selectColumn(int)
func (this *QTableView) SelectColumn(args ...interface{}) () {
  // selectColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView12selectColumnEi
    // invoke: void selectColumn(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTableView12selectColumnEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "selectColumn", args)
  }

  return
}

// setSortingEnabled(_Bool)
func (this *QTableView) SetSortingEnabled(args ...interface{}) () {
  // setSortingEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView17setSortingEnabledEb
    // invoke: void setSortingEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTableView17setSortingEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "setSortingEnabled", args)
  }

  return
}

// resizeRowToContents(int)
func (this *QTableView) ResizeRowToContents(args ...interface{}) () {
  // resizeRowToContents(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView19resizeRowToContentsEi
    // invoke: void resizeRowToContents(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTableView19resizeRowToContentsEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "resizeRowToContents", args)
  }

  return
}

// resizeRowsToContents()
func (this *QTableView) ResizeRowsToContents(args ...interface{}) () {
  // resizeRowsToContents()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView20resizeRowsToContentsEv
    // invoke: void resizeRowsToContents()
    C.C_ZN10QTableView20resizeRowsToContentsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTableView", "resizeRowsToContents", args)
  }

  return
}

// wordWrap()
func (this *QTableView) WordWrap(args ...interface{}) (ret interface{}) {
  // wordWrap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView8wordWrapEv
    // invoke: bool wordWrap()
    var ret0 = C.C_ZNK10QTableView8wordWrapEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableView", "wordWrap", args)
  }

  return
}

// <= body block end

