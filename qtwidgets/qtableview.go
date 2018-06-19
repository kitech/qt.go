package qtwidgets

// /usr/include/qt/QtWidgets/qtableview.h
// #include <qtableview.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// void rowMoved(int, int, int)
func (this *QTableView) InheritRowMoved(f func(row int, oldIndex int, newIndex int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "rowMoved", f)
}

// void columnMoved(int, int, int)
func (this *QTableView) InheritColumnMoved(f func(column int, oldIndex int, newIndex int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "columnMoved", f)
}

// void rowResized(int, int, int)
func (this *QTableView) InheritRowResized(f func(row int, oldHeight int, newHeight int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "rowResized", f)
}

// void columnResized(int, int, int)
func (this *QTableView) InheritColumnResized(f func(column int, oldWidth int, newWidth int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "columnResized", f)
}

// void rowCountChanged(int, int)
func (this *QTableView) InheritRowCountChanged(f func(oldCount int, newCount int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "rowCountChanged", f)
}

// void columnCountChanged(int, int)
func (this *QTableView) InheritColumnCountChanged(f func(oldCount int, newCount int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "columnCountChanged", f)
}

// void scrollContentsBy(int, int)
func (this *QTableView) InheritScrollContentsBy(f func(dx int, dy int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "scrollContentsBy", f)
}

// QStyleOptionViewItem viewOptions()
func (this *QTableView) InheritViewOptions(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "viewOptions", f)
}

// void paintEvent(QPaintEvent *)
func (this *QTableView) InheritPaintEvent(f func(e *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void timerEvent(QTimerEvent *)
func (this *QTableView) InheritTimerEvent(f func(event *qtcore.QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

// int horizontalOffset()
func (this *QTableView) InheritHorizontalOffset(f func() int) {
	qtrt.SetAllInheritCallback(this, "horizontalOffset", f)
}

// int verticalOffset()
func (this *QTableView) InheritVerticalOffset(f func() int) {
	qtrt.SetAllInheritCallback(this, "verticalOffset", f)
}

// QModelIndex moveCursor(QAbstractItemView::CursorAction, Qt::KeyboardModifiers)
func (this *QTableView) InheritMoveCursor(f func(cursorAction int, modifiers int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "moveCursor", f)
}

// void setSelection(const QRect &, QItemSelectionModel::SelectionFlags)
func (this *QTableView) InheritSetSelection(f func(rect *qtcore.QRect, command int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setSelection", f)
}

// QRegion visualRegionForSelection(const QItemSelection &)
func (this *QTableView) InheritVisualRegionForSelection(f func(selection *qtcore.QItemSelection) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "visualRegionForSelection", f)
}

// QModelIndexList selectedIndexes()
func (this *QTableView) InheritSelectedIndexes(f func() *qtcore.QModelIndexList /*9999*/) {
	qtrt.SetAllInheritCallback(this, "selectedIndexes", f)
}

// void updateGeometries()
func (this *QTableView) InheritUpdateGeometries(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateGeometries", f)
}

// QSize viewportSizeHint()
func (this *QTableView) InheritViewportSizeHint(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "viewportSizeHint", f)
}

// int sizeHintForRow(int)
func (this *QTableView) InheritSizeHintForRow(f func(row int) int) {
	qtrt.SetAllInheritCallback(this, "sizeHintForRow", f)
}

// int sizeHintForColumn(int)
func (this *QTableView) InheritSizeHintForColumn(f func(column int) int) {
	qtrt.SetAllInheritCallback(this, "sizeHintForColumn", f)
}

// void verticalScrollbarAction(int)
func (this *QTableView) InheritVerticalScrollbarAction(f func(action int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "verticalScrollbarAction", f)
}

// void horizontalScrollbarAction(int)
func (this *QTableView) InheritHorizontalScrollbarAction(f func(action int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "horizontalScrollbarAction", f)
}

// bool isIndexHidden(const QModelIndex &)
func (this *QTableView) InheritIsIndexHidden(f func(index *qtcore.QModelIndex) bool) {
	qtrt.SetAllInheritCallback(this, "isIndexHidden", f)
}

// void selectionChanged(const QItemSelection &, const QItemSelection &)
func (this *QTableView) InheritSelectionChanged(f func(selected *qtcore.QItemSelection, deselected *qtcore.QItemSelection) /*void*/) {
	qtrt.SetAllInheritCallback(this, "selectionChanged", f)
}

// void currentChanged(const QModelIndex &, const QModelIndex &)
func (this *QTableView) InheritCurrentChanged(f func(current *qtcore.QModelIndex, previous *qtcore.QModelIndex) /*void*/) {
	qtrt.SetAllInheritCallback(this, "currentChanged", f)
}

/*

 */
type QTableView struct {
	*QAbstractItemView
}
type QTableView_ITF interface {
	QAbstractItemView_ITF
	QTableView_PTR() *QTableView
}

func (ptr *QTableView) QTableView_PTR() *QTableView { return ptr }

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QTableView) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtableview.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTableView(QWidget *)

/*
Constructs a table view with a parent to represent the data.

See also QAbstractItemModel.
*/
func NewQTableView(parent QWidget_ITF /*777 QWidget **/) *QTableView {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableViewC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTableViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTableView")
	return gothis
}

// /usr/include/qt/QtWidgets/qtableview.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTableView(QWidget *)

/*
Constructs a table view with a parent to represent the data.

See also QAbstractItemModel.
*/
func NewQTableView__() *QTableView {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableViewC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTableViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTableView")
	return gothis
}

// /usr/include/qt/QtWidgets/qtableview.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTableView()

/*

 */
func DeleteQTableView(this *QTableView) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableViewD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qtableview.h:68
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setModel(QAbstractItemModel *)

/*
Reimplemented from QAbstractItemView::setModel().
*/
func (this *QTableView) SetModel(model qtcore.QAbstractItemModel_ITF /*777 QAbstractItemModel **/) {
	var convArg0 unsafe.Pointer
	if model != nil && model.QAbstractItemModel_PTR() != nil {
		convArg0 = model.QAbstractItemModel_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView8setModelEP18QAbstractItemModel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setRootIndex(const QModelIndex &)

/*
Reimplemented from QAbstractItemView::setRootIndex().
*/
func (this *QTableView) SetRootIndex(index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView12setRootIndexERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:70
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setSelectionModel(QItemSelectionModel *)

/*
Reimplemented from QAbstractItemView::setSelectionModel().
*/
func (this *QTableView) SetSelectionModel(selectionModel qtcore.QItemSelectionModel_ITF /*777 QItemSelectionModel **/) {
	var convArg0 unsafe.Pointer
	if selectionModel != nil && selectionModel.QItemSelectionModel_PTR() != nil {
		convArg0 = selectionModel.QItemSelectionModel_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView17setSelectionModelEP19QItemSelectionModel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void doItemsLayout()

/*

 */
func (this *QTableView) DoItemsLayout() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView13doItemsLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] QHeaderView * horizontalHeader() const

/*
Returns the table view's horizontal header.

See also setHorizontalHeader(), verticalHeader(), and QAbstractItemModel::headerData().
*/
func (this *QTableView) HorizontalHeader() *QHeaderView /*777 QHeaderView **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView16horizontalHeaderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQHeaderViewFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtableview.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QHeaderView * verticalHeader() const

/*
Returns the table view's vertical header.

See also setVerticalHeader(), horizontalHeader(), and QAbstractItemModel::headerData().
*/
func (this *QTableView) VerticalHeader() *QHeaderView /*777 QHeaderView **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView14verticalHeaderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQHeaderViewFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtableview.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHorizontalHeader(QHeaderView *)

/*
Sets the widget to use for the horizontal header to header.

See also horizontalHeader() and setVerticalHeader().
*/
func (this *QTableView) SetHorizontalHeader(header QHeaderView_ITF /*777 QHeaderView **/) {
	var convArg0 unsafe.Pointer
	if header != nil && header.QHeaderView_PTR() != nil {
		convArg0 = header.QHeaderView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView19setHorizontalHeaderEP11QHeaderView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalHeader(QHeaderView *)

/*
Sets the widget to use for the vertical header to header.

See also verticalHeader() and setHorizontalHeader().
*/
func (this *QTableView) SetVerticalHeader(header QHeaderView_ITF /*777 QHeaderView **/) {
	var convArg0 unsafe.Pointer
	if header != nil && header.QHeaderView_PTR() != nil {
		convArg0 = header.QHeaderView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView17setVerticalHeaderEP11QHeaderView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:78
// index:0
// Public Visibility=Default Availability=Available
// [4] int rowViewportPosition(int) const

/*
Returns the y-coordinate in contents coordinates of the given row.
*/
func (this *QTableView) RowViewportPosition(row int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView19rowViewportPositionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtableview.h:79
// index:0
// Public Visibility=Default Availability=Available
// [4] int rowAt(int) const

/*
Returns the row in which the given y-coordinate, y, in contents coordinates is located.

Note: This function returns -1 if the given coordinate is not valid (has no row).

See also columnAt().
*/
func (this *QTableView) RowAt(y int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView5rowAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), y)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtableview.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowHeight(int, int)

/*
Sets the height of the given row to be height.

This function was introduced in  Qt 4.1.

See also rowHeight().
*/
func (this *QTableView) SetRowHeight(row int, height int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView12setRowHeightEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, height)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] int rowHeight(int) const

/*
Returns the height of the given row.

See also setRowHeight(), resizeRowToContents(), and columnWidth().
*/
func (this *QTableView) RowHeight(row int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView9rowHeightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtableview.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnViewportPosition(int) const

/*
Returns the x-coordinate in contents coordinates of the given column.
*/
func (this *QTableView) ColumnViewportPosition(column int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView22columnViewportPositionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtableview.h:85
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnAt(int) const

/*
Returns the column in which the given x-coordinate, x, in contents coordinates is located.

Note: This function returns -1 if the given coordinate is not valid (has no column).

See also rowAt().
*/
func (this *QTableView) ColumnAt(x int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView8columnAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtableview.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnWidth(int, int)

/*
Sets the width of the given column to be width.

This function was introduced in  Qt 4.1.

See also columnWidth().
*/
func (this *QTableView) SetColumnWidth(column int, width int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView14setColumnWidthEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:88
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnWidth(int) const

/*
Returns the width of the given column.

See also setColumnWidth(), resizeColumnToContents(), and rowHeight().
*/
func (this *QTableView) ColumnWidth(column int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView11columnWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtableview.h:90
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRowHidden(int) const

/*
Returns true if the given row is hidden; otherwise returns false.

See also isColumnHidden().
*/
func (this *QTableView) IsRowHidden(row int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView11isRowHiddenEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtableview.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowHidden(int, bool)

/*
If hide is true row will be hidden, otherwise it will be shown.

See also isRowHidden() and setColumnHidden().
*/
func (this *QTableView) SetRowHidden(row int, hide bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView12setRowHiddenEib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, hide)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:93
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isColumnHidden(int) const

/*
Returns true if the given column is hidden; otherwise returns false.

See also isRowHidden().
*/
func (this *QTableView) IsColumnHidden(column int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView14isColumnHiddenEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtableview.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnHidden(int, bool)

/*
If hide is true the given column will be hidden; otherwise it will be shown.

See also isColumnHidden() and setRowHidden().
*/
func (this *QTableView) SetColumnHidden(column int, hide bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView15setColumnHiddenEib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, hide)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSortingEnabled(bool)

/*
If enable is true, enables sorting for the table and immediately trigger a call to sortByColumn() with the current sort section and order

Note: Setter function for property sortingEnabled.

See also isSortingEnabled().
*/
func (this *QTableView) SetSortingEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView17setSortingEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:97
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSortingEnabled() const

/*

 */
func (this *QTableView) IsSortingEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView16isSortingEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtableview.h:99
// index:0
// Public Visibility=Default Availability=Available
// [1] bool showGrid() const

/*

 */
func (this *QTableView) ShowGrid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView8showGridEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtableview.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::PenStyle gridStyle() const

/*

 */
func (this *QTableView) GridStyle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView9gridStyleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGridStyle(Qt::PenStyle)

/*

 */
func (this *QTableView) SetGridStyle(style int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView12setGridStyleEN2Qt8PenStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), style)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWordWrap(bool)

/*

 */
func (this *QTableView) SetWordWrap(on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView11setWordWrapEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:105
// index:0
// Public Visibility=Default Availability=Available
// [1] bool wordWrap() const

/*

 */
func (this *QTableView) WordWrap() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView8wordWrapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtableview.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCornerButtonEnabled(bool)

/*

 */
func (this *QTableView) SetCornerButtonEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView22setCornerButtonEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:109
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCornerButtonEnabled() const

/*

 */
func (this *QTableView) IsCornerButtonEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView21isCornerButtonEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtableview.h:112
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QRect visualRect(const QModelIndex &) const

/*

 */
func (this *QTableView) VisualRect(index qtcore.QModelIndex_ITF) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView10visualRectERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qtableview.h:113
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void scrollTo(const QModelIndex &, QAbstractItemView::ScrollHint)

/*

 */
func (this *QTableView) ScrollTo(index qtcore.QModelIndex_ITF, hint int) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:113
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void scrollTo(const QModelIndex &, QAbstractItemView::ScrollHint)

/*

 */
func (this *QTableView) ScrollTo__(index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	// arg: 1, QAbstractItemView::ScrollHint=Enum, QAbstractItemView::ScrollHint=Enum, , Invalid
	hint := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:114
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex indexAt(const QPoint &) const

/*
Reimplemented from QAbstractItemView::indexAt().

Returns the index position of the model item corresponding to the table item at position pos in contents coordinates.
*/
func (this *QTableView) IndexAt(p qtcore.QPoint_ITF) *qtcore.QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView7indexAtERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qtableview.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpan(int, int, int, int)

/*
Sets the span of the table element at (row, column) to the number of rows and columns specified by (rowSpanCount, columnSpanCount).

This function was introduced in  Qt 4.2.

See also rowSpan() and columnSpan().
*/
func (this *QTableView) SetSpan(row int, column int, rowSpan int, columnSpan int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView7setSpanEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, rowSpan, columnSpan)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:117
// index:0
// Public Visibility=Default Availability=Available
// [4] int rowSpan(int, int) const

/*
Returns the row span of the table element at (row, column). The default is 1.

This function was introduced in  Qt 4.2.

See also setSpan() and columnSpan().
*/
func (this *QTableView) RowSpan(row int, column int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView7rowSpanEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtableview.h:118
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnSpan(int, int) const

/*
Returns the column span of the table element at (row, column). The default is 1.

This function was introduced in  Qt 4.2.

See also setSpan() and rowSpan().
*/
func (this *QTableView) ColumnSpan(row int, column int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView10columnSpanEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtableview.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearSpans()

/*
Removes all row and column spans in the table view.

This function was introduced in  Qt 4.4.

See also setSpan().
*/
func (this *QTableView) ClearSpans() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView10clearSpansEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sortByColumn(int, Qt::SortOrder)

/*
Sorts the model by the values in the given column in the given order.

This function was introduced in  Qt 4.2.

See also sortingEnabled.
*/
func (this *QTableView) SortByColumn(column int, order int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView12sortByColumnEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:134
// index:1
// Public Visibility=Default Availability=Available
// [-2] void sortByColumn(int)

/*
Sorts the model by the values in the given column in the given order.

This function was introduced in  Qt 4.2.

See also sortingEnabled.
*/
func (this *QTableView) SortByColumn_1(column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView12sortByColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectRow(int)

/*
Selects the given row in the table view if the current SelectionMode and SelectionBehavior allows rows to be selected.

See also selectColumn().
*/
func (this *QTableView) SelectRow(row int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView9selectRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectColumn(int)

/*
Selects the given column in the table view if the current SelectionMode and SelectionBehavior allows columns to be selected.

See also selectRow().
*/
func (this *QTableView) SelectColumn(column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView12selectColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void hideRow(int)

/*
Hide the given row.

See also showRow() and hideColumn().
*/
func (this *QTableView) HideRow(row int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView7hideRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void hideColumn(int)

/*
Hide the given column.

See also showColumn() and hideRow().
*/
func (this *QTableView) HideColumn(column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView10hideColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showRow(int)

/*
Show the given row.

See also hideRow() and showColumn().
*/
func (this *QTableView) ShowRow(row int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView7showRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showColumn(int)

/*
Show the given column.

See also hideColumn() and showRow().
*/
func (this *QTableView) ShowColumn(column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView10showColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resizeRowToContents(int)

/*
Resizes the given row based on the size hints of the delegate used to render each item in the row.

See also resizeRowsToContents(), sizeHintForRow(), and QHeaderView::resizeContentsPrecision().
*/
func (this *QTableView) ResizeRowToContents(row int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView19resizeRowToContentsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resizeRowsToContents()

/*
Resizes all rows based on the size hints of the delegate used to render each item in the rows.

See also resizeRowToContents(), sizeHintForRow(), and QHeaderView::resizeContentsPrecision().
*/
func (this *QTableView) ResizeRowsToContents() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView20resizeRowsToContentsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resizeColumnToContents(int)

/*
Resizes the given column based on the size hints of the delegate used to render each item in the column.

Note: Only visible columns will be resized. Reimplement sizeHintForColumn() to resize hidden columns as well.

See also resizeColumnsToContents(), sizeHintForColumn(), and QHeaderView::resizeContentsPrecision().
*/
func (this *QTableView) ResizeColumnToContents(column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView22resizeColumnToContentsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resizeColumnsToContents()

/*
Resizes all columns based on the size hints of the delegate used to render each item in the columns.

See also resizeColumnToContents(), sizeHintForColumn(), and QHeaderView::resizeContentsPrecision().
*/
func (this *QTableView) ResizeColumnsToContents() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView23resizeColumnsToContentsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setShowGrid(bool)

/*

 */
func (this *QTableView) SetShowGrid(show bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView11setShowGridEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), show)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:138
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void rowMoved(int, int, int)

/*
This slot is called to change the index of the given row in the table view. The old index is specified by oldIndex, and the new index by newIndex.

See also columnMoved().
*/
func (this *QTableView) RowMoved(row int, oldIndex int, newIndex int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView8rowMovedEiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, oldIndex, newIndex)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:139
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void columnMoved(int, int, int)

/*
This slot is called to change the index of the given column in the table view. The old index is specified by oldIndex, and the new index by newIndex.

See also rowMoved().
*/
func (this *QTableView) ColumnMoved(column int, oldIndex int, newIndex int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView11columnMovedEiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, oldIndex, newIndex)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:140
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void rowResized(int, int, int)

/*
This slot is called to change the height of the given row. The old height is specified by oldHeight, and the new height by newHeight.

See also columnResized().
*/
func (this *QTableView) RowResized(row int, oldHeight int, newHeight int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView10rowResizedEiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, oldHeight, newHeight)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:141
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void columnResized(int, int, int)

/*
This slot is called to change the width of the given column. The old width is specified by oldWidth, and the new width by newWidth.

See also rowResized().
*/
func (this *QTableView) ColumnResized(column int, oldWidth int, newWidth int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView13columnResizedEiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, oldWidth, newWidth)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:142
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void rowCountChanged(int, int)

/*
This slot is called whenever rows are added or deleted. The previous number of rows is specified by oldCount, and the new number of rows is specified by newCount.
*/
func (this *QTableView) RowCountChanged(oldCount int, newCount int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView15rowCountChangedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), oldCount, newCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:143
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void columnCountChanged(int, int)

/*
This slot is called whenever columns are added or deleted. The previous number of columns is specified by oldCount, and the new number of columns is specified by newCount.
*/
func (this *QTableView) ColumnCountChanged(oldCount int, newCount int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView18columnCountChangedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), oldCount, newCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:147
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void scrollContentsBy(int, int)

/*

 */
func (this *QTableView) ScrollContentsBy(dx int, dy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView16scrollContentsByEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:149
// index:0
// Protected virtual Visibility=Default Availability=Available
// [192] QStyleOptionViewItem viewOptions() const

/*
Reimplemented from QAbstractItemView::viewOptions().
*/
func (this *QTableView) ViewOptions() *QStyleOptionViewItem /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView11viewOptionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStyleOptionViewItemFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStyleOptionViewItem)
	return rv2
}

// /usr/include/qt/QtWidgets/qtableview.h:150
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().

Paints the table on receipt of the given paint event event.
*/
func (this *QTableView) PaintEvent(e qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QPaintEvent_PTR() != nil {
		convArg0 = e.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:152
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)

/*
Reimplemented from QObject::timerEvent().
*/
func (this *QTableView) TimerEvent(event qtcore.QTimerEvent_ITF /*777 QTimerEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QTimerEvent_PTR() != nil {
		convArg0 = event.QTimerEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:154
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int horizontalOffset() const

/*
Reimplemented from QAbstractItemView::horizontalOffset().

Returns the horizontal offset of the items in the table view.

Note that the table view uses the horizontal header section positions to determine the positions of columns in the view.

See also verticalOffset().
*/
func (this *QTableView) HorizontalOffset() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView16horizontalOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtableview.h:155
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int verticalOffset() const

/*
Reimplemented from QAbstractItemView::verticalOffset().

Returns the vertical offset of the items in the table view.

Note that the table view uses the vertical header section positions to determine the positions of rows in the view.

See also horizontalOffset().
*/
func (this *QTableView) VerticalOffset() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView14verticalOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtableview.h:156
// index:0
// Protected virtual Visibility=Default Availability=Available
// [24] QModelIndex moveCursor(QAbstractItemView::CursorAction, Qt::KeyboardModifiers)

/*
Reimplemented from QAbstractItemView::moveCursor().

Moves the cursor in accordance with the given cursorAction, using the information provided by the modifiers.

See also QAbstractItemView::CursorAction.
*/
func (this *QTableView) MoveCursor(cursorAction int, modifiers int) *qtcore.QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView10moveCursorEN17QAbstractItemView12CursorActionE6QFlagsIN2Qt16KeyboardModifierEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cursorAction, modifiers)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qtableview.h:158
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void setSelection(const QRect &, QItemSelectionModel::SelectionFlags)

/*
Reimplemented from QAbstractItemView::setSelection().

Selects the items within the given rect and in accordance with the specified selection flags.
*/
func (this *QTableView) SetSelection(rect qtcore.QRect_ITF, command int) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, command)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:159
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QRegion visualRegionForSelection(const QItemSelection &) const

/*

 */
func (this *QTableView) VisualRegionForSelection(selection qtcore.QItemSelection_ITF) *qtgui.QRegion /*123*/ {
	var convArg0 unsafe.Pointer
	if selection != nil && selection.QItemSelection_PTR() != nil {
		convArg0 = selection.QItemSelection_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView24visualRegionForSelectionERK14QItemSelection", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtWidgets/qtableview.h:160
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QModelIndexList selectedIndexes() const

/*
Reimplemented from QAbstractItemView::selectedIndexes().
*/
func (this *QTableView) SelectedIndexes() *qtcore.QModelIndexList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView15selectedIndexesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtWidgets/qtableview.h:162
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateGeometries()

/*
Reimplemented from QAbstractItemView::updateGeometries().
*/
func (this *QTableView) UpdateGeometries() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView16updateGeometriesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:164
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QSize viewportSizeHint() const

/*
Reimplemented from QAbstractScrollArea::viewportSizeHint().
*/
func (this *QTableView) ViewportSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView16viewportSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qtableview.h:166
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int sizeHintForRow(int) const

/*
Reimplemented from QAbstractItemView::sizeHintForRow().

Returns the size hint for the given row's height or -1 if there is no model.

If you need to set the height of a given row to a fixed value, call QHeaderView::resizeSection() on the table's vertical header.

If you reimplement this function in a subclass, note that the value you return is only used when resizeRowToContents() is called. In that case, if a larger row height is required by either the vertical header or the item delegate, that width will be used instead.

See also QWidget::sizeHint, verticalHeader(), and QHeaderView::resizeContentsPrecision().
*/
func (this *QTableView) SizeHintForRow(row int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView14sizeHintForRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtableview.h:167
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int sizeHintForColumn(int) const

/*
Reimplemented from QAbstractItemView::sizeHintForColumn().

Returns the size hint for the given column's width or -1 if there is no model.

If you need to set the width of a given column to a fixed value, call QHeaderView::resizeSection() on the table's horizontal header.

If you reimplement this function in a subclass, note that the value you return will be used when resizeColumnToContents() or QHeaderView::resizeSections() is called. If a larger column width is required by either the horizontal header or the item delegate, the larger width will be used instead.

See also QWidget::sizeHint, horizontalHeader(), and QHeaderView::resizeContentsPrecision().
*/
func (this *QTableView) SizeHintForColumn(column int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView17sizeHintForColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtableview.h:169
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void verticalScrollbarAction(int)

/*

 */
func (this *QTableView) VerticalScrollbarAction(action int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView23verticalScrollbarActionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:170
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void horizontalScrollbarAction(int)

/*

 */
func (this *QTableView) HorizontalScrollbarAction(action int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView25horizontalScrollbarActionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:172
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool isIndexHidden(const QModelIndex &) const

/*
Reimplemented from QAbstractItemView::isIndexHidden().
*/
func (this *QTableView) IsIndexHidden(index qtcore.QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView13isIndexHiddenERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtableview.h:174
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void selectionChanged(const QItemSelection &, const QItemSelection &)

/*
Reimplemented from QAbstractItemView::selectionChanged().
*/
func (this *QTableView) SelectionChanged(selected qtcore.QItemSelection_ITF, deselected qtcore.QItemSelection_ITF) {
	var convArg0 unsafe.Pointer
	if selected != nil && selected.QItemSelection_PTR() != nil {
		convArg0 = selected.QItemSelection_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if deselected != nil && deselected.QItemSelection_PTR() != nil {
		convArg1 = deselected.QItemSelection_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView16selectionChangedERK14QItemSelectionS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:176
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void currentChanged(const QModelIndex &, const QModelIndex &)

/*
Reimplemented from QAbstractItemView::currentChanged().
*/
func (this *QTableView) CurrentChanged(current qtcore.QModelIndex_ITF, previous qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if current != nil && current.QModelIndex_PTR() != nil {
		convArg0 = current.QModelIndex_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if previous != nil && previous.QModelIndex_PTR() != nil {
		convArg1 = previous.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView14currentChangedERK11QModelIndexS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
