//  header block begin
// /usr/include/qt/QtWidgets/qtableview.h
// #include <qtableview.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QTableView struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qtableview.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QTableView) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:65
// index:0
// void QTableView(class QWidget *)
func NewQTableView(parent unsafe.Pointer) *QTableView {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableViewC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QTableView{cthis}
}

// /usr/include/qt/QtWidgets/qtableview.h:66
// index:0
// virtual
// void ~QTableView()
func DeleteQTableView(*QTableView) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableViewD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:68
// index:0
// virtual
// void setModel(class QAbstractItemModel *)
func (this *QTableView) SetModel(model unsafe.Pointer) {
	// 0: (, QAbstractItemModel * model), (model)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView8setModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_VOID, this.cthis, model)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:69
// index:0
// virtual
// void setRootIndex(const class QModelIndex &)
func (this *QTableView) SetRootIndex(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView12setRootIndexERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:70
// index:0
// virtual
// void setSelectionModel(class QItemSelectionModel *)
func (this *QTableView) SetSelectionModel(selectionModel unsafe.Pointer) {
	// 0: (, QItemSelectionModel * selectionModel), (selectionModel)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView17setSelectionModelEP19QItemSelectionModel", ffiqt.FFI_TYPE_VOID, this.cthis, selectionModel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:71
// index:0
// virtual
// void doItemsLayout()
func (this *QTableView) DoItemsLayout() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView13doItemsLayoutEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:73
// index:0
// QHeaderView * horizontalHeader()
func (this *QTableView) HorizontalHeader() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView16horizontalHeaderEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:74
// index:0
// QHeaderView * verticalHeader()
func (this *QTableView) VerticalHeader() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView14verticalHeaderEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:75
// index:0
// void setHorizontalHeader(class QHeaderView *)
func (this *QTableView) SetHorizontalHeader(header unsafe.Pointer) {
	// 0: (, QHeaderView * header), (header)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView19setHorizontalHeaderEP11QHeaderView", ffiqt.FFI_TYPE_VOID, this.cthis, header)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:76
// index:0
// void setVerticalHeader(class QHeaderView *)
func (this *QTableView) SetVerticalHeader(header unsafe.Pointer) {
	// 0: (, QHeaderView * header), (header)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView17setVerticalHeaderEP11QHeaderView", ffiqt.FFI_TYPE_VOID, this.cthis, header)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:78
// index:0
// int rowViewportPosition(int)
func (this *QTableView) RowViewportPosition(row int) {
	// 0: (, int row), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView19rowViewportPositionEi", ffiqt.FFI_TYPE_VOID, this.cthis, &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:79
// index:0
// int rowAt(int)
func (this *QTableView) RowAt(y int) {
	// 0: (, int y), (&y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView5rowAtEi", ffiqt.FFI_TYPE_VOID, this.cthis, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:81
// index:0
// void setRowHeight(int, int)
func (this *QTableView) SetRowHeight(row int, height int) {
	// 0: (, int row, int height), (&row, &height)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView12setRowHeightEii", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &height)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:82
// index:0
// int rowHeight(int)
func (this *QTableView) RowHeight(row int) {
	// 0: (, int row), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView9rowHeightEi", ffiqt.FFI_TYPE_VOID, this.cthis, &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:84
// index:0
// int columnViewportPosition(int)
func (this *QTableView) ColumnViewportPosition(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView22columnViewportPositionEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:85
// index:0
// int columnAt(int)
func (this *QTableView) ColumnAt(x int) {
	// 0: (, int x), (&x)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView8columnAtEi", ffiqt.FFI_TYPE_VOID, this.cthis, &x)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:87
// index:0
// void setColumnWidth(int, int)
func (this *QTableView) SetColumnWidth(column int, width int) {
	// 0: (, int column, int width), (&column, &width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView14setColumnWidthEii", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:88
// index:0
// int columnWidth(int)
func (this *QTableView) ColumnWidth(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView11columnWidthEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:90
// index:0
// bool isRowHidden(int)
func (this *QTableView) IsRowHidden(row int) {
	// 0: (, int row), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView11isRowHiddenEi", ffiqt.FFI_TYPE_VOID, this.cthis, &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:91
// index:0
// void setRowHidden(int, _Bool)
func (this *QTableView) SetRowHidden(row int, hide bool) {
	// 0: (, int row, bool hide), (&row, &hide)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView12setRowHiddenEib", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:93
// index:0
// bool isColumnHidden(int)
func (this *QTableView) IsColumnHidden(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView14isColumnHiddenEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:94
// index:0
// void setColumnHidden(int, _Bool)
func (this *QTableView) SetColumnHidden(column int, hide bool) {
	// 0: (, int column, bool hide), (&column, &hide)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView15setColumnHiddenEib", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:96
// index:0
// void setSortingEnabled(_Bool)
func (this *QTableView) SetSortingEnabled(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView17setSortingEnabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:97
// index:0
// bool isSortingEnabled()
func (this *QTableView) IsSortingEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView16isSortingEnabledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:99
// index:0
// bool showGrid()
func (this *QTableView) ShowGrid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView8showGridEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:101
// index:0
// Qt::PenStyle gridStyle()
func (this *QTableView) GridStyle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView9gridStyleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:102
// index:0
// void setGridStyle(Qt::PenStyle)
func (this *QTableView) SetGridStyle(style int) {
	// 0: (, Qt::PenStyle style), (&style)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView12setGridStyleEN2Qt8PenStyleE", ffiqt.FFI_TYPE_VOID, this.cthis, &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:104
// index:0
// void setWordWrap(_Bool)
func (this *QTableView) SetWordWrap(on bool) {
	// 0: (, bool on), (&on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView11setWordWrapEb", ffiqt.FFI_TYPE_VOID, this.cthis, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:105
// index:0
// bool wordWrap()
func (this *QTableView) WordWrap() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView8wordWrapEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:108
// index:0
// void setCornerButtonEnabled(_Bool)
func (this *QTableView) SetCornerButtonEnabled(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView22setCornerButtonEnabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:109
// index:0
// bool isCornerButtonEnabled()
func (this *QTableView) IsCornerButtonEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView21isCornerButtonEnabledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:112
// index:0
// virtual
// QRect visualRect(const class QModelIndex &)
func (this *QTableView) VisualRect(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView10visualRectERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:113
// index:0
// virtual
// void scrollTo(const class QModelIndex &, enum QAbstractItemView::ScrollHint)
func (this *QTableView) ScrollTo(index unsafe.Pointer, hint int) {
	// 0: (, const QModelIndex & index, QAbstractItemView::ScrollHint hint), (index, &hint)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE", ffiqt.FFI_TYPE_VOID, this.cthis, index, &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:114
// index:0
// virtual
// QModelIndex indexAt(const class QPoint &)
func (this *QTableView) IndexAt(p unsafe.Pointer) {
	// 0: (, const QPoint & p), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView7indexAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:116
// index:0
// void setSpan(int, int, int, int)
func (this *QTableView) SetSpan(row int, column int, rowSpan int, columnSpan int) {
	// 0: (, int row, int column, int rowSpan, int columnSpan), (&row, &column, &rowSpan, &columnSpan)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView7setSpanEiiii", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column, &rowSpan, &columnSpan)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:117
// index:0
// int rowSpan(int, int)
func (this *QTableView) RowSpan(row int, column int) {
	// 0: (, int row, int column), (&row, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView7rowSpanEii", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:118
// index:0
// int columnSpan(int, int)
func (this *QTableView) ColumnSpan(row int, column int) {
	// 0: (, int row, int column), (&row, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView10columnSpanEii", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:119
// index:0
// void clearSpans()
func (this *QTableView) ClearSpans() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView10clearSpansEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:121
// index:0
// void sortByColumn(int, Qt::SortOrder)
func (this *QTableView) SortByColumn(column int, order int) {
	// 0: (, int column, Qt::SortOrder order), (&column, &order)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView12sortByColumnEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:134
// index:1
// void sortByColumn(int)
func (this *QTableView) SortByColumn_1(column int) {
	// 1: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView12sortByColumnEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:124
// index:0
// void selectRow(int)
func (this *QTableView) SelectRow(row int) {
	// 0: (, int row), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView9selectRowEi", ffiqt.FFI_TYPE_VOID, this.cthis, &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:125
// index:0
// void selectColumn(int)
func (this *QTableView) SelectColumn(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView12selectColumnEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:126
// index:0
// void hideRow(int)
func (this *QTableView) HideRow(row int) {
	// 0: (, int row), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView7hideRowEi", ffiqt.FFI_TYPE_VOID, this.cthis, &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:127
// index:0
// void hideColumn(int)
func (this *QTableView) HideColumn(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView10hideColumnEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:128
// index:0
// void showRow(int)
func (this *QTableView) ShowRow(row int) {
	// 0: (, int row), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView7showRowEi", ffiqt.FFI_TYPE_VOID, this.cthis, &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:129
// index:0
// void showColumn(int)
func (this *QTableView) ShowColumn(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView10showColumnEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:130
// index:0
// void resizeRowToContents(int)
func (this *QTableView) ResizeRowToContents(row int) {
	// 0: (, int row), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView19resizeRowToContentsEi", ffiqt.FFI_TYPE_VOID, this.cthis, &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:131
// index:0
// void resizeRowsToContents()
func (this *QTableView) ResizeRowsToContents() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView20resizeRowsToContentsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:132
// index:0
// void resizeColumnToContents(int)
func (this *QTableView) ResizeColumnToContents(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView22resizeColumnToContentsEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:133
// index:0
// void resizeColumnsToContents()
func (this *QTableView) ResizeColumnsToContents() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView23resizeColumnsToContentsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:135
// index:0
// void setShowGrid(_Bool)
func (this *QTableView) SetShowGrid(show bool) {
	// 0: (, bool show), (&show)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView11setShowGridEb", ffiqt.FFI_TYPE_VOID, this.cthis, &show)
	gopp.ErrPrint(err, rv)
}

//  body block end
