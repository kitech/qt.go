//  header block begin
// /usr/include/qt/QtWidgets/qtableview.h
// #include <qtableview.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
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
	*QAbstractItemView
}

func (this *QTableView) GetCthis() unsafe.Pointer {
	return this.QAbstractItemView.GetCthis()
}

// /usr/include/qt/QtWidgets/qtableview.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QTableView) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:65
// index:0
// void QTableView(class QWidget *)
func NewQTableView(parent unsafe.Pointer) *QTableView {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableViewC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQTableViewFromPointer(cthis)
	return gothis
}
func NewQTableViewFromPointer(cthis unsafe.Pointer) *QTableView {
	bcthis0 := NewQAbstractItemViewFromPointer(cthis)
	return &QTableView{bcthis0}
}

// /usr/include/qt/QtWidgets/qtableview.h:146
// index:1
// void QTableView(class QTableViewPrivate &, class QWidget *)
func NewQTableView_1(arg0 unsafe.Pointer, parent unsafe.Pointer) *QTableView {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableViewC2ER17QTableViewPrivateP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, arg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQTableViewFromPointer(cthis)
	return gothis
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
	// 0: (, model QAbstractItemModel *), (model)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView8setModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_VOID, this.GetCthis(), model)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:69
// index:0
// virtual
// void setRootIndex(const class QModelIndex &)
func (this *QTableView) SetRootIndex(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView12setRootIndexERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:70
// index:0
// virtual
// void setSelectionModel(class QItemSelectionModel *)
func (this *QTableView) SetSelectionModel(selectionModel unsafe.Pointer) {
	// 0: (, selectionModel QItemSelectionModel *), (selectionModel)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView17setSelectionModelEP19QItemSelectionModel", ffiqt.FFI_TYPE_VOID, this.GetCthis(), selectionModel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:71
// index:0
// virtual
// void doItemsLayout()
func (this *QTableView) DoItemsLayout() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView13doItemsLayoutEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:73
// index:0
// QHeaderView * horizontalHeader()
func (this *QTableView) HorizontalHeader() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView16horizontalHeaderEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:74
// index:0
// QHeaderView * verticalHeader()
func (this *QTableView) VerticalHeader() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView14verticalHeaderEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:75
// index:0
// void setHorizontalHeader(class QHeaderView *)
func (this *QTableView) SetHorizontalHeader(header unsafe.Pointer) {
	// 0: (, header QHeaderView *), (header)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView19setHorizontalHeaderEP11QHeaderView", ffiqt.FFI_TYPE_VOID, this.GetCthis(), header)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:76
// index:0
// void setVerticalHeader(class QHeaderView *)
func (this *QTableView) SetVerticalHeader(header unsafe.Pointer) {
	// 0: (, header QHeaderView *), (header)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView17setVerticalHeaderEP11QHeaderView", ffiqt.FFI_TYPE_VOID, this.GetCthis(), header)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:78
// index:0
// int rowViewportPosition(int)
func (this *QTableView) RowViewportPosition(row int) {
	// 0: (, row int), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView19rowViewportPositionEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:79
// index:0
// int rowAt(int)
func (this *QTableView) RowAt(y int) {
	// 0: (, y int), (&y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView5rowAtEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:81
// index:0
// void setRowHeight(int, int)
func (this *QTableView) SetRowHeight(row int, height int) {
	// 0: (, row int, height int), (&row, &height)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView12setRowHeightEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &height)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:82
// index:0
// int rowHeight(int)
func (this *QTableView) RowHeight(row int) {
	// 0: (, row int), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView9rowHeightEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:84
// index:0
// int columnViewportPosition(int)
func (this *QTableView) ColumnViewportPosition(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView22columnViewportPositionEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:85
// index:0
// int columnAt(int)
func (this *QTableView) ColumnAt(x int) {
	// 0: (, x int), (&x)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView8columnAtEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:87
// index:0
// void setColumnWidth(int, int)
func (this *QTableView) SetColumnWidth(column int, width int) {
	// 0: (, column int, width int), (&column, &width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView14setColumnWidthEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:88
// index:0
// int columnWidth(int)
func (this *QTableView) ColumnWidth(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView11columnWidthEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:90
// index:0
// bool isRowHidden(int)
func (this *QTableView) IsRowHidden(row int) {
	// 0: (, row int), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView11isRowHiddenEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:91
// index:0
// void setRowHidden(int, _Bool)
func (this *QTableView) SetRowHidden(row int, hide bool) {
	// 0: (, row int, hide bool), (&row, &hide)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView12setRowHiddenEib", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:93
// index:0
// bool isColumnHidden(int)
func (this *QTableView) IsColumnHidden(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView14isColumnHiddenEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:94
// index:0
// void setColumnHidden(int, _Bool)
func (this *QTableView) SetColumnHidden(column int, hide bool) {
	// 0: (, column int, hide bool), (&column, &hide)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView15setColumnHiddenEib", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, &hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:96
// index:0
// void setSortingEnabled(_Bool)
func (this *QTableView) SetSortingEnabled(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView17setSortingEnabledEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:97
// index:0
// bool isSortingEnabled()
func (this *QTableView) IsSortingEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView16isSortingEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:99
// index:0
// bool showGrid()
func (this *QTableView) ShowGrid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView8showGridEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:101
// index:0
// Qt::PenStyle gridStyle()
func (this *QTableView) GridStyle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView9gridStyleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:102
// index:0
// void setGridStyle(Qt::PenStyle)
func (this *QTableView) SetGridStyle(style int) {
	// 0: (, style Qt::PenStyle), (&style)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView12setGridStyleEN2Qt8PenStyleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:104
// index:0
// void setWordWrap(_Bool)
func (this *QTableView) SetWordWrap(on bool) {
	// 0: (, on bool), (&on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView11setWordWrapEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:105
// index:0
// bool wordWrap()
func (this *QTableView) WordWrap() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView8wordWrapEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:108
// index:0
// void setCornerButtonEnabled(_Bool)
func (this *QTableView) SetCornerButtonEnabled(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView22setCornerButtonEnabledEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:109
// index:0
// bool isCornerButtonEnabled()
func (this *QTableView) IsCornerButtonEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView21isCornerButtonEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:112
// index:0
// virtual
// QRect visualRect(const class QModelIndex &)
func (this *QTableView) VisualRect(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView10visualRectERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:113
// index:0
// virtual
// void scrollTo(const class QModelIndex &, enum QAbstractItemView::ScrollHint)
func (this *QTableView) ScrollTo(index unsafe.Pointer, hint int) {
	// 0: (, index const QModelIndex &, hint QAbstractItemView::ScrollHint), (index, &hint)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index, &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:114
// index:0
// virtual
// QModelIndex indexAt(const class QPoint &)
func (this *QTableView) IndexAt(p unsafe.Pointer) {
	// 0: (, p const QPoint &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView7indexAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:116
// index:0
// void setSpan(int, int, int, int)
func (this *QTableView) SetSpan(row int, column int, rowSpan int, columnSpan int) {
	// 0: (, row int, column int, rowSpan int, columnSpan int), (&row, &column, &rowSpan, &columnSpan)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView7setSpanEiiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &column, &rowSpan, &columnSpan)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:117
// index:0
// int rowSpan(int, int)
func (this *QTableView) RowSpan(row int, column int) {
	// 0: (, row int, column int), (&row, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView7rowSpanEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:118
// index:0
// int columnSpan(int, int)
func (this *QTableView) ColumnSpan(row int, column int) {
	// 0: (, row int, column int), (&row, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView10columnSpanEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:119
// index:0
// void clearSpans()
func (this *QTableView) ClearSpans() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView10clearSpansEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:121
// index:0
// void sortByColumn(int, Qt::SortOrder)
func (this *QTableView) SortByColumn(column int, order int) {
	// 0: (, column int, order Qt::SortOrder), (&column, &order)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView12sortByColumnEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:134
// index:1
// void sortByColumn(int)
func (this *QTableView) SortByColumn_1(column int) {
	// 1: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView12sortByColumnEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:124
// index:0
// void selectRow(int)
func (this *QTableView) SelectRow(row int) {
	// 0: (, row int), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView9selectRowEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:125
// index:0
// void selectColumn(int)
func (this *QTableView) SelectColumn(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView12selectColumnEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:126
// index:0
// void hideRow(int)
func (this *QTableView) HideRow(row int) {
	// 0: (, row int), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView7hideRowEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:127
// index:0
// void hideColumn(int)
func (this *QTableView) HideColumn(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView10hideColumnEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:128
// index:0
// void showRow(int)
func (this *QTableView) ShowRow(row int) {
	// 0: (, row int), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView7showRowEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:129
// index:0
// void showColumn(int)
func (this *QTableView) ShowColumn(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView10showColumnEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:130
// index:0
// void resizeRowToContents(int)
func (this *QTableView) ResizeRowToContents(row int) {
	// 0: (, row int), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView19resizeRowToContentsEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:131
// index:0
// void resizeRowsToContents()
func (this *QTableView) ResizeRowsToContents() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView20resizeRowsToContentsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:132
// index:0
// void resizeColumnToContents(int)
func (this *QTableView) ResizeColumnToContents(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView22resizeColumnToContentsEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:133
// index:0
// void resizeColumnsToContents()
func (this *QTableView) ResizeColumnsToContents() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView23resizeColumnsToContentsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:135
// index:0
// void setShowGrid(_Bool)
func (this *QTableView) SetShowGrid(show bool) {
	// 0: (, show bool), (&show)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView11setShowGridEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &show)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:138
// index:0
// void rowMoved(int, int, int)
func (this *QTableView) RowMoved(row int, oldIndex int, newIndex int) {
	// 0: (, row int, oldIndex int, newIndex int), (&row, &oldIndex, &newIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView8rowMovedEiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &oldIndex, &newIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:139
// index:0
// void columnMoved(int, int, int)
func (this *QTableView) ColumnMoved(column int, oldIndex int, newIndex int) {
	// 0: (, column int, oldIndex int, newIndex int), (&column, &oldIndex, &newIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView11columnMovedEiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, &oldIndex, &newIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:140
// index:0
// void rowResized(int, int, int)
func (this *QTableView) RowResized(row int, oldHeight int, newHeight int) {
	// 0: (, row int, oldHeight int, newHeight int), (&row, &oldHeight, &newHeight)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView10rowResizedEiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &oldHeight, &newHeight)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:141
// index:0
// void columnResized(int, int, int)
func (this *QTableView) ColumnResized(column int, oldWidth int, newWidth int) {
	// 0: (, column int, oldWidth int, newWidth int), (&column, &oldWidth, &newWidth)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView13columnResizedEiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, &oldWidth, &newWidth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:142
// index:0
// void rowCountChanged(int, int)
func (this *QTableView) RowCountChanged(oldCount int, newCount int) {
	// 0: (, oldCount int, newCount int), (&oldCount, &newCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView15rowCountChangedEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &oldCount, &newCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:143
// index:0
// void columnCountChanged(int, int)
func (this *QTableView) ColumnCountChanged(oldCount int, newCount int) {
	// 0: (, oldCount int, newCount int), (&oldCount, &newCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView18columnCountChangedEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &oldCount, &newCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:147
// index:0
// virtual
// void scrollContentsBy(int, int)
func (this *QTableView) ScrollContentsBy(dx int, dy int) {
	// 0: (, dx int, dy int), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView16scrollContentsByEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:149
// index:0
// virtual
// QStyleOptionViewItem viewOptions()
func (this *QTableView) ViewOptions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView11viewOptionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:150
// index:0
// virtual
// void paintEvent(class QPaintEvent *)
func (this *QTableView) PaintEvent(e unsafe.Pointer) {
	// 0: (, e QPaintEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:152
// index:0
// virtual
// void timerEvent(class QTimerEvent *)
func (this *QTableView) TimerEvent(event unsafe.Pointer) {
	// 0: (, event QTimerEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:154
// index:0
// virtual
// int horizontalOffset()
func (this *QTableView) HorizontalOffset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView16horizontalOffsetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:155
// index:0
// virtual
// int verticalOffset()
func (this *QTableView) VerticalOffset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView14verticalOffsetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:156
// index:0
// virtual
// QModelIndex moveCursor(enum QAbstractItemView::CursorAction, Qt::KeyboardModifiers)
func (this *QTableView) MoveCursor(cursorAction int, modifiers int) {
	// 0: (, cursorAction QAbstractItemView::CursorAction, QFlags<Qt::KeyboardModifier> modifiers), (&cursorAction, &modifiers)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView10moveCursorEN17QAbstractItemView12CursorActionE6QFlagsIN2Qt16KeyboardModifierEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &cursorAction, &modifiers)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:158
// index:0
// virtual
// void setSelection(const class QRect &, class QItemSelectionModel::SelectionFlags)
func (this *QTableView) SetSelection(rect unsafe.Pointer, command int) {
	// 0: (, rect const QRect &, QFlags<QItemSelectionModel::SelectionFlag> command), (rect, &command)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &command)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:159
// index:0
// virtual
// QRegion visualRegionForSelection(const class QItemSelection &)
func (this *QTableView) VisualRegionForSelection(selection unsafe.Pointer) {
	// 0: (, selection const QItemSelection &), (selection)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView24visualRegionForSelectionERK14QItemSelection", ffiqt.FFI_TYPE_VOID, this.GetCthis(), selection)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:160
// index:0
// virtual
// QModelIndexList selectedIndexes()
func (this *QTableView) SelectedIndexes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView15selectedIndexesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:162
// index:0
// virtual
// void updateGeometries()
func (this *QTableView) UpdateGeometries() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView16updateGeometriesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:164
// index:0
// virtual
// QSize viewportSizeHint()
func (this *QTableView) ViewportSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView16viewportSizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:166
// index:0
// virtual
// int sizeHintForRow(int)
func (this *QTableView) SizeHintForRow(row int) {
	// 0: (, row int), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView14sizeHintForRowEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:167
// index:0
// virtual
// int sizeHintForColumn(int)
func (this *QTableView) SizeHintForColumn(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView17sizeHintForColumnEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:169
// index:0
// virtual
// void verticalScrollbarAction(int)
func (this *QTableView) VerticalScrollbarAction(action int) {
	// 0: (, action int), (&action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView23verticalScrollbarActionEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:170
// index:0
// virtual
// void horizontalScrollbarAction(int)
func (this *QTableView) HorizontalScrollbarAction(action int) {
	// 0: (, action int), (&action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView25horizontalScrollbarActionEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:172
// index:0
// virtual
// bool isIndexHidden(const class QModelIndex &)
func (this *QTableView) IsIndexHidden(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView13isIndexHiddenERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:174
// index:0
// virtual
// void selectionChanged(const class QItemSelection &, const class QItemSelection &)
func (this *QTableView) SelectionChanged(selected unsafe.Pointer, deselected unsafe.Pointer) {
	// 0: (, selected const QItemSelection &, deselected const QItemSelection &), (selected, deselected)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView16selectionChangedERK14QItemSelectionS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), selected, deselected)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:176
// index:0
// virtual
// void currentChanged(const class QModelIndex &, const class QModelIndex &)
func (this *QTableView) CurrentChanged(current unsafe.Pointer, previous unsafe.Pointer) {
	// 0: (, current const QModelIndex &, previous const QModelIndex &), (current, previous)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView14currentChangedERK11QModelIndexS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), current, previous)
	gopp.ErrPrint(err, rv)
}

//  body block end
