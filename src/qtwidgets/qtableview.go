package qtwidgets

// /usr/include/qt/QtWidgets/qtableview.h
// #include <qtableview.h>
// #include <QtWidgets>

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
	if this == nil {
		return nil
	} else {
		return this.QAbstractItemView.GetCthis()
	}
}
func (this *QTableView) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractItemView = NewQAbstractItemViewFromPointer(cthis)
}
func NewQTableViewFromPointer(cthis unsafe.Pointer) *QTableView {
	bcthis0 := NewQAbstractItemViewFromPointer(cthis)
	return &QTableView{bcthis0}
}
func (*QTableView) NewFromPointer(cthis unsafe.Pointer) *QTableView {
	return NewQTableViewFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qtableview.h:55
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QTableView) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtableview.h:65
// index:0
// Public
// void QTableView(class QWidget *)
func NewQTableView(parent *QWidget /*444 QWidget **/) *QTableView {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableViewC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTableViewFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtableview.h:66
// index:0
// Public virtual
// void ~QTableView()
func DeleteQTableView(*QTableView) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableViewD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:68
// index:0
// Public virtual
// void setModel(class QAbstractItemModel *)
func (this *QTableView) SetModel(model *qtcore.QAbstractItemModel /*444 QAbstractItemModel **/) {
	var convArg0 = model.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView8setModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:69
// index:0
// Public virtual
// void setRootIndex(const class QModelIndex &)
func (this *QTableView) SetRootIndex(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView12setRootIndexERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:70
// index:0
// Public virtual
// void setSelectionModel(class QItemSelectionModel *)
func (this *QTableView) SetSelectionModel(selectionModel *qtcore.QItemSelectionModel /*444 QItemSelectionModel **/) {
	var convArg0 = selectionModel.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView17setSelectionModelEP19QItemSelectionModel", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:71
// index:0
// Public virtual
// void doItemsLayout()
func (this *QTableView) DoItemsLayout() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView13doItemsLayoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:73
// index:0
// Public
// QHeaderView * horizontalHeader()
func (this *QTableView) HorizontalHeader() *QHeaderView /*444 QHeaderView **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView16horizontalHeaderEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQHeaderViewFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtableview.h:74
// index:0
// Public
// QHeaderView * verticalHeader()
func (this *QTableView) VerticalHeader() *QHeaderView /*444 QHeaderView **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView14verticalHeaderEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQHeaderViewFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtableview.h:75
// index:0
// Public
// void setHorizontalHeader(class QHeaderView *)
func (this *QTableView) SetHorizontalHeader(header *QHeaderView /*444 QHeaderView **/) {
	var convArg0 = header.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView19setHorizontalHeaderEP11QHeaderView", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:76
// index:0
// Public
// void setVerticalHeader(class QHeaderView *)
func (this *QTableView) SetVerticalHeader(header *QHeaderView /*444 QHeaderView **/) {
	var convArg0 = header.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView17setVerticalHeaderEP11QHeaderView", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:78
// index:0
// Public
// int rowViewportPosition(int)
func (this *QTableView) RowViewportPosition(row int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView19rowViewportPositionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtableview.h:79
// index:0
// Public
// int rowAt(int)
func (this *QTableView) RowAt(y int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView5rowAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), y)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtableview.h:81
// index:0
// Public
// void setRowHeight(int, int)
func (this *QTableView) SetRowHeight(row int, height int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView12setRowHeightEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, height)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:82
// index:0
// Public
// int rowHeight(int)
func (this *QTableView) RowHeight(row int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView9rowHeightEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtableview.h:84
// index:0
// Public
// int columnViewportPosition(int)
func (this *QTableView) ColumnViewportPosition(column int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView22columnViewportPositionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtableview.h:85
// index:0
// Public
// int columnAt(int)
func (this *QTableView) ColumnAt(x int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView8columnAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtableview.h:87
// index:0
// Public
// void setColumnWidth(int, int)
func (this *QTableView) SetColumnWidth(column int, width int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView14setColumnWidthEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:88
// index:0
// Public
// int columnWidth(int)
func (this *QTableView) ColumnWidth(column int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView11columnWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtableview.h:90
// index:0
// Public
// bool isRowHidden(int)
func (this *QTableView) IsRowHidden(row int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView11isRowHiddenEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtableview.h:91
// index:0
// Public
// void setRowHidden(int, _Bool)
func (this *QTableView) SetRowHidden(row int, hide bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView12setRowHiddenEib", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:93
// index:0
// Public
// bool isColumnHidden(int)
func (this *QTableView) IsColumnHidden(column int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView14isColumnHiddenEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtableview.h:94
// index:0
// Public
// void setColumnHidden(int, _Bool)
func (this *QTableView) SetColumnHidden(column int, hide bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView15setColumnHiddenEib", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:96
// index:0
// Public
// void setSortingEnabled(_Bool)
func (this *QTableView) SetSortingEnabled(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView17setSortingEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:97
// index:0
// Public
// bool isSortingEnabled()
func (this *QTableView) IsSortingEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView16isSortingEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtableview.h:99
// index:0
// Public
// bool showGrid()
func (this *QTableView) ShowGrid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView8showGridEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtableview.h:101
// index:0
// Public
// Qt::PenStyle gridStyle()
func (this *QTableView) GridStyle() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView9gridStyleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:102
// index:0
// Public
// void setGridStyle(Qt::PenStyle)
func (this *QTableView) SetGridStyle(style int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView12setGridStyleEN2Qt8PenStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:104
// index:0
// Public
// void setWordWrap(_Bool)
func (this *QTableView) SetWordWrap(on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView11setWordWrapEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:105
// index:0
// Public
// bool wordWrap()
func (this *QTableView) WordWrap() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView8wordWrapEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtableview.h:108
// index:0
// Public
// void setCornerButtonEnabled(_Bool)
func (this *QTableView) SetCornerButtonEnabled(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView22setCornerButtonEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:109
// index:0
// Public
// bool isCornerButtonEnabled()
func (this *QTableView) IsCornerButtonEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView21isCornerButtonEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtableview.h:112
// index:0
// Public virtual
// QRect visualRect(const class QModelIndex &)
func (this *QTableView) VisualRect(index *qtcore.QModelIndex) *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView10visualRectERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtableview.h:113
// index:0
// Public virtual
// void scrollTo(const class QModelIndex &, enum QAbstractItemView::ScrollHint)
func (this *QTableView) ScrollTo(index *qtcore.QModelIndex, hint int) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:114
// index:0
// Public virtual
// QModelIndex indexAt(const class QPoint &)
func (this *QTableView) IndexAt(p *qtcore.QPoint) *qtcore.QModelIndex /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView7indexAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtableview.h:116
// index:0
// Public
// void setSpan(int, int, int, int)
func (this *QTableView) SetSpan(row int, column int, rowSpan int, columnSpan int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView7setSpanEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column, rowSpan, columnSpan)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:117
// index:0
// Public
// int rowSpan(int, int)
func (this *QTableView) RowSpan(row int, column int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView7rowSpanEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtableview.h:118
// index:0
// Public
// int columnSpan(int, int)
func (this *QTableView) ColumnSpan(row int, column int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView10columnSpanEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtableview.h:119
// index:0
// Public
// void clearSpans()
func (this *QTableView) ClearSpans() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView10clearSpansEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:121
// index:0
// Public
// void sortByColumn(int, Qt::SortOrder)
func (this *QTableView) SortByColumn(column int, order int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView12sortByColumnEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:134
// index:1
// Public
// void sortByColumn(int)
func (this *QTableView) SortByColumn_1(column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView12sortByColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:124
// index:0
// Public
// void selectRow(int)
func (this *QTableView) SelectRow(row int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView9selectRowEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:125
// index:0
// Public
// void selectColumn(int)
func (this *QTableView) SelectColumn(column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView12selectColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:126
// index:0
// Public
// void hideRow(int)
func (this *QTableView) HideRow(row int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView7hideRowEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:127
// index:0
// Public
// void hideColumn(int)
func (this *QTableView) HideColumn(column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView10hideColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:128
// index:0
// Public
// void showRow(int)
func (this *QTableView) ShowRow(row int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView7showRowEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:129
// index:0
// Public
// void showColumn(int)
func (this *QTableView) ShowColumn(column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView10showColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:130
// index:0
// Public
// void resizeRowToContents(int)
func (this *QTableView) ResizeRowToContents(row int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView19resizeRowToContentsEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:131
// index:0
// Public
// void resizeRowsToContents()
func (this *QTableView) ResizeRowsToContents() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView20resizeRowsToContentsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:132
// index:0
// Public
// void resizeColumnToContents(int)
func (this *QTableView) ResizeColumnToContents(column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView22resizeColumnToContentsEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:133
// index:0
// Public
// void resizeColumnsToContents()
func (this *QTableView) ResizeColumnsToContents() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView23resizeColumnsToContentsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:135
// index:0
// Public
// void setShowGrid(_Bool)
func (this *QTableView) SetShowGrid(show bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView11setShowGridEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), show)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:138
// index:0
// Protected
// void rowMoved(int, int, int)
func (this *QTableView) RowMoved(row int, oldIndex int, newIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView8rowMovedEiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, oldIndex, newIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:139
// index:0
// Protected
// void columnMoved(int, int, int)
func (this *QTableView) ColumnMoved(column int, oldIndex int, newIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView11columnMovedEiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, oldIndex, newIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:140
// index:0
// Protected
// void rowResized(int, int, int)
func (this *QTableView) RowResized(row int, oldHeight int, newHeight int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView10rowResizedEiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, oldHeight, newHeight)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:141
// index:0
// Protected
// void columnResized(int, int, int)
func (this *QTableView) ColumnResized(column int, oldWidth int, newWidth int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView13columnResizedEiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, oldWidth, newWidth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:142
// index:0
// Protected
// void rowCountChanged(int, int)
func (this *QTableView) RowCountChanged(oldCount int, newCount int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView15rowCountChangedEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), oldCount, newCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:143
// index:0
// Protected
// void columnCountChanged(int, int)
func (this *QTableView) ColumnCountChanged(oldCount int, newCount int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView18columnCountChangedEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), oldCount, newCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:147
// index:0
// Protected virtual
// void scrollContentsBy(int, int)
func (this *QTableView) ScrollContentsBy(dx int, dy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView16scrollContentsByEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:149
// index:0
// Protected virtual
// QStyleOptionViewItem viewOptions()
func (this *QTableView) ViewOptions() *QStyleOptionViewItem /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView11viewOptionsEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStyleOptionViewItemFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtableview.h:150
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QTableView) PaintEvent(e *qtgui.QPaintEvent /*444 QPaintEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:152
// index:0
// Protected virtual
// void timerEvent(class QTimerEvent *)
func (this *QTableView) TimerEvent(event *qtcore.QTimerEvent /*444 QTimerEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:154
// index:0
// Protected virtual
// int horizontalOffset()
func (this *QTableView) HorizontalOffset() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView16horizontalOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtableview.h:155
// index:0
// Protected virtual
// int verticalOffset()
func (this *QTableView) VerticalOffset() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView14verticalOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtableview.h:158
// index:0
// Protected virtual
// void setSelection(const class QRect &, class QItemSelectionModel::SelectionFlags)
func (this *QTableView) SetSelection(rect *qtcore.QRect, command int) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, command)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:159
// index:0
// Protected virtual
// QRegion visualRegionForSelection(const class QItemSelection &)
func (this *QTableView) VisualRegionForSelection(selection *qtcore.QItemSelection) *qtgui.QRegion /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = selection.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView24visualRegionForSelectionERK14QItemSelection", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtableview.h:162
// index:0
// Protected virtual
// void updateGeometries()
func (this *QTableView) UpdateGeometries() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView16updateGeometriesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:164
// index:0
// Protected virtual
// QSize viewportSizeHint()
func (this *QTableView) ViewportSizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView16viewportSizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtableview.h:166
// index:0
// Protected virtual
// int sizeHintForRow(int)
func (this *QTableView) SizeHintForRow(row int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView14sizeHintForRowEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtableview.h:167
// index:0
// Protected virtual
// int sizeHintForColumn(int)
func (this *QTableView) SizeHintForColumn(column int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView17sizeHintForColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtableview.h:169
// index:0
// Protected virtual
// void verticalScrollbarAction(int)
func (this *QTableView) VerticalScrollbarAction(action int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView23verticalScrollbarActionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:170
// index:0
// Protected virtual
// void horizontalScrollbarAction(int)
func (this *QTableView) HorizontalScrollbarAction(action int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView25horizontalScrollbarActionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:172
// index:0
// Protected virtual
// bool isIndexHidden(const class QModelIndex &)
func (this *QTableView) IsIndexHidden(index *qtcore.QModelIndex) bool {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTableView13isIndexHiddenERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtableview.h:174
// index:0
// Protected virtual
// void selectionChanged(const class QItemSelection &, const class QItemSelection &)
func (this *QTableView) SelectionChanged(selected *qtcore.QItemSelection, deselected *qtcore.QItemSelection) {
	var convArg0 = selected.GetCthis()
	var convArg1 = deselected.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView16selectionChangedERK14QItemSelectionS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:176
// index:0
// Protected virtual
// void currentChanged(const class QModelIndex &, const class QModelIndex &)
func (this *QTableView) CurrentChanged(current *qtcore.QModelIndex, previous *qtcore.QModelIndex) {
	var convArg0 = current.GetCthis()
	var convArg1 = previous.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTableView14currentChangedERK11QModelIndexS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

//  body block end
