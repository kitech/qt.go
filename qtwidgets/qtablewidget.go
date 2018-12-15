package qtwidgets

// /usr/include/qt/QtWidgets/qtablewidget.h
// #include <qtablewidget.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 45
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

// bool event(QEvent *)
func (this *QTableWidget) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// QStringList mimeTypes()
func (this *QTableWidget) InheritMimeTypes(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "mimeTypes", f)
}

// bool dropMimeData(int, int, const QMimeData *, Qt::DropAction)
func (this *QTableWidget) InheritDropMimeData(f func(row int, column int, data *qtcore.QMimeData /*777 const QMimeData **/, action int) bool) {
	qtrt.SetAllInheritCallback(this, "dropMimeData", f)
}

// Qt::DropActions supportedDropActions()
func (this *QTableWidget) InheritSupportedDropActions(f func() int) {
	qtrt.SetAllInheritCallback(this, "supportedDropActions", f)
}

// QModelIndex indexFromItem(QTableWidgetItem *)
func (this *QTableWidget) InheritIndexFromItem(f func(item *QTableWidgetItem /*777 QTableWidgetItem **/) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "indexFromItem", f)
}

// QTableWidgetItem * itemFromIndex(const QModelIndex &)
func (this *QTableWidget) InheritItemFromIndex(f func(index *qtcore.QModelIndex) unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "itemFromIndex", f)
}

// void dropEvent(QDropEvent *)
func (this *QTableWidget) InheritDropEvent(f func(event *qtgui.QDropEvent /*777 QDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dropEvent", f)
}

/*

 */
type QTableWidget struct {
	*QTableView
}
type QTableWidget_ITF interface {
	QTableView_ITF
	QTableWidget_PTR() *QTableWidget
}

func (ptr *QTableWidget) QTableWidget_PTR() *QTableWidget { return ptr }

func (this *QTableWidget) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QTableView.GetCthis()
	}
}
func (this *QTableWidget) SetCthis(cthis unsafe.Pointer) {
	this.QTableView = NewQTableViewFromPointer(cthis)
}
func NewQTableWidgetFromPointer(cthis unsafe.Pointer) *QTableWidget {
	bcthis0 := NewQTableViewFromPointer(cthis)
	return &QTableWidget{bcthis0}
}
func (*QTableWidget) NewFromPointer(cthis unsafe.Pointer) *QTableWidget {
	return NewQTableWidgetFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:218
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QTableWidget) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtablewidget.h:224
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTableWidget(QWidget *)

/*
Creates a new table view with the given parent.
*/
func (*QTableWidget) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QTableWidget {
	return NewQTableWidget(parent)
}
func NewQTableWidget(parent QWidget_ITF /*777 QWidget **/) *QTableWidget {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidgetC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTableWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTableWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qtablewidget.h:224
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTableWidget(QWidget *)

/*
Creates a new table view with the given parent.
*/
func (*QTableWidget) NewForInheritp() *QTableWidget {
	return NewQTableWidgetp()
}
func NewQTableWidgetp() *QTableWidget {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidgetC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTableWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTableWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qtablewidget.h:225
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTableWidget(int, int, QWidget *)

/*
Creates a new table view with the given parent.
*/
func (*QTableWidget) NewForInherit1(rows int, columns int, parent QWidget_ITF /*777 QWidget **/) *QTableWidget {
	return NewQTableWidget1(rows, columns, parent)
}
func NewQTableWidget1(rows int, columns int, parent QWidget_ITF /*777 QWidget **/) *QTableWidget {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg2 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidgetC2EiiP7QWidget", qtrt.FFI_TYPE_POINTER, rows, columns, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTableWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTableWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qtablewidget.h:225
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTableWidget(int, int, QWidget *)

/*
Creates a new table view with the given parent.
*/
func (*QTableWidget) NewForInherit1p(rows int, columns int) *QTableWidget {
	return NewQTableWidget1p(rows, columns)
}
func NewQTableWidget1p(rows int, columns int) *QTableWidget {
	// arg: 2, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidgetC2EiiP7QWidget", qtrt.FFI_TYPE_POINTER, rows, columns, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTableWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTableWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qtablewidget.h:226
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTableWidget()

/*

 */
func DeleteQTableWidget(this *QTableWidget) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidgetD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:228
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowCount(int)

/*
Sets the number of rows in this table's model to rows. If this is less than rowCount(), the data in the unwanted rows is discarded.

Note: Setter function for property rowCount.

See also rowCount() and setColumnCount().
*/
func (this *QTableWidget) SetRowCount(rows int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget11setRowCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rows)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:229
// index:0
// Public Visibility=Default Availability=Available
// [4] int rowCount() const

/*
Returns the number of rows.

Note: Getter function for property rowCount.

See also setRowCount().
*/
func (this *QTableWidget) RowCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget8rowCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:231
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnCount(int)

/*
Sets the number of columns in this table's model to columns. If this is less than columnCount(), the data in the unwanted columns is discarded.

Note: Setter function for property columnCount.

See also columnCount() and setRowCount().
*/
func (this *QTableWidget) SetColumnCount(columns int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget14setColumnCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), columns)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:232
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnCount() const

/*
Returns the number of columns.

Note: Getter function for property columnCount.

See also setColumnCount().
*/
func (this *QTableWidget) ColumnCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget11columnCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:234
// index:0
// Public Visibility=Default Availability=Available
// [4] int row(const QTableWidgetItem *) const

/*
Returns the row for the item.
*/
func (this *QTableWidget) Row(item QTableWidgetItem_ITF /*777 const QTableWidgetItem **/) int {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTableWidgetItem_PTR() != nil {
		convArg0 = item.QTableWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget3rowEPK16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:235
// index:0
// Public Visibility=Default Availability=Available
// [4] int column(const QTableWidgetItem *) const

/*
Returns the column for the item.
*/
func (this *QTableWidget) Column(item QTableWidgetItem_ITF /*777 const QTableWidgetItem **/) int {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTableWidgetItem_PTR() != nil {
		convArg0 = item.QTableWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget6columnEPK16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:237
// index:0
// Public Visibility=Default Availability=Available
// [8] QTableWidgetItem * item(int, int) const

/*
Returns the item for the given row and column if one has been set; otherwise returns 0.

See also setItem().
*/
func (this *QTableWidget) Item(row int, column int) *QTableWidgetItem /*777 QTableWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget4itemEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtablewidget.h:238
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItem(int, int, QTableWidgetItem *)

/*
Sets the item for the given row and column to item.

The table takes ownership of the item.

Note that if sorting is enabled (see sortingEnabled) and column is the current sort column, the row will be moved to the sorted position determined by item.

If you want to set several items of a particular row (say, by calling setItem() in a loop), you may want to turn off sorting before doing so, and turn it back on afterwards; this will allow you to use the same row argument for all items in the same row (i.e. setItem() will not move the row).

See also item() and takeItem().
*/
func (this *QTableWidget) SetItem(row int, column int, item QTableWidgetItem_ITF /*777 QTableWidgetItem **/) {
	var convArg2 unsafe.Pointer
	if item != nil && item.QTableWidgetItem_PTR() != nil {
		convArg2 = item.QTableWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget7setItemEiiP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:239
// index:0
// Public Visibility=Default Availability=Available
// [8] QTableWidgetItem * takeItem(int, int)

/*
Removes the item at row and column from the table without deleting it.
*/
func (this *QTableWidget) TakeItem(row int, column int) *QTableWidgetItem /*777 QTableWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget8takeItemEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtablewidget.h:241
// index:0
// Public Visibility=Default Availability=Available
// [8] QTableWidgetItem * verticalHeaderItem(int) const

/*
Returns the vertical header item for row row.

See also setVerticalHeaderItem().
*/
func (this *QTableWidget) VerticalHeaderItem(row int) *QTableWidgetItem /*777 QTableWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget18verticalHeaderItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtablewidget.h:242
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalHeaderItem(int, QTableWidgetItem *)

/*
Sets the vertical header item for row row to item.

See also verticalHeaderItem().
*/
func (this *QTableWidget) SetVerticalHeaderItem(row int, item QTableWidgetItem_ITF /*777 QTableWidgetItem **/) {
	var convArg1 unsafe.Pointer
	if item != nil && item.QTableWidgetItem_PTR() != nil {
		convArg1 = item.QTableWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget21setVerticalHeaderItemEiP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:243
// index:0
// Public Visibility=Default Availability=Available
// [8] QTableWidgetItem * takeVerticalHeaderItem(int)

/*
Removes the vertical header item at row from the header without deleting it.

This function was introduced in  Qt 4.1.
*/
func (this *QTableWidget) TakeVerticalHeaderItem(row int) *QTableWidgetItem /*777 QTableWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget22takeVerticalHeaderItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtablewidget.h:245
// index:0
// Public Visibility=Default Availability=Available
// [8] QTableWidgetItem * horizontalHeaderItem(int) const

/*
Returns the horizontal header item for column, column, if one has been set; otherwise returns 0.

See also setHorizontalHeaderItem().
*/
func (this *QTableWidget) HorizontalHeaderItem(column int) *QTableWidgetItem /*777 QTableWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget20horizontalHeaderItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtablewidget.h:246
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHorizontalHeaderItem(int, QTableWidgetItem *)

/*
Sets the horizontal header item for column column to item. If necessary, the column count is increased to fit the item. The previous header item (if there was one) is deleted.

See also horizontalHeaderItem().
*/
func (this *QTableWidget) SetHorizontalHeaderItem(column int, item QTableWidgetItem_ITF /*777 QTableWidgetItem **/) {
	var convArg1 unsafe.Pointer
	if item != nil && item.QTableWidgetItem_PTR() != nil {
		convArg1 = item.QTableWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget23setHorizontalHeaderItemEiP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:247
// index:0
// Public Visibility=Default Availability=Available
// [8] QTableWidgetItem * takeHorizontalHeaderItem(int)

/*
Removes the horizontal header item at column from the header without deleting it.

This function was introduced in  Qt 4.1.
*/
func (this *QTableWidget) TakeHorizontalHeaderItem(column int) *QTableWidgetItem /*777 QTableWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget24takeHorizontalHeaderItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtablewidget.h:248
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalHeaderLabels(const QStringList &)

/*
Sets the vertical header labels using labels.
*/
func (this *QTableWidget) SetVerticalHeaderLabels(labels qtcore.QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if labels != nil && labels.QStringList_PTR() != nil {
		convArg0 = labels.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget23setVerticalHeaderLabelsERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:249
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHorizontalHeaderLabels(const QStringList &)

/*
Sets the horizontal header labels using labels.
*/
func (this *QTableWidget) SetHorizontalHeaderLabels(labels qtcore.QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if labels != nil && labels.QStringList_PTR() != nil {
		convArg0 = labels.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget25setHorizontalHeaderLabelsERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:251
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentRow() const

/*
Returns the row of the current item.

See also currentColumn() and setCurrentCell().
*/
func (this *QTableWidget) CurrentRow() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget10currentRowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:252
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentColumn() const

/*
Returns the column of the current item.

See also currentRow() and setCurrentCell().
*/
func (this *QTableWidget) CurrentColumn() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget13currentColumnEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:253
// index:0
// Public Visibility=Default Availability=Available
// [8] QTableWidgetItem * currentItem() const

/*
Returns the current item.

See also setCurrentItem().
*/
func (this *QTableWidget) CurrentItem() *QTableWidgetItem /*777 QTableWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget11currentItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtablewidget.h:254
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentItem(QTableWidgetItem *)

/*
Sets the current item to item.

Unless the selection mode is NoSelection, the item is also selected.

See also currentItem() and setCurrentCell().
*/
func (this *QTableWidget) SetCurrentItem(item QTableWidgetItem_ITF /*777 QTableWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTableWidgetItem_PTR() != nil {
		convArg0 = item.QTableWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget14setCurrentItemEP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:255
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setCurrentItem(QTableWidgetItem *, QItemSelectionModel::SelectionFlags)

/*
Sets the current item to item.

Unless the selection mode is NoSelection, the item is also selected.

See also currentItem() and setCurrentCell().
*/
func (this *QTableWidget) SetCurrentItem1(item QTableWidgetItem_ITF /*777 QTableWidgetItem **/, command int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTableWidgetItem_PTR() != nil {
		convArg0 = item.QTableWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget14setCurrentItemEP16QTableWidgetItem6QFlagsIN19QItemSelectionModel13SelectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, command)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:256
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentCell(int, int)

/*
Sets the current cell to be the cell at position (row, column).

Depending on the current selection mode, the cell may also be selected.

This function was introduced in  Qt 4.1.

See also setCurrentItem(), currentRow(), and currentColumn().
*/
func (this *QTableWidget) SetCurrentCell(row int, column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget14setCurrentCellEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:257
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setCurrentCell(int, int, QItemSelectionModel::SelectionFlags)

/*
Sets the current cell to be the cell at position (row, column).

Depending on the current selection mode, the cell may also be selected.

This function was introduced in  Qt 4.1.

See also setCurrentItem(), currentRow(), and currentColumn().
*/
func (this *QTableWidget) SetCurrentCell1(row int, column int, command int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget14setCurrentCellEii6QFlagsIN19QItemSelectionModel13SelectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, command)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:259
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sortItems(int, Qt::SortOrder)

/*
Sorts all the rows in the table widget based on column and order.
*/
func (this *QTableWidget) SortItems(column int, order int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget9sortItemsEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:259
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sortItems(int, Qt::SortOrder)

/*
Sorts all the rows in the table widget based on column and order.
*/
func (this *QTableWidget) SortItemsp(column int) {
	// arg: 1, Qt::SortOrder=Elaborated, Qt::SortOrder=Enum, , Invalid
	order := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget9sortItemsEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:260
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSortingEnabled(bool)

/*

 */
func (this *QTableWidget) SetSortingEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget17setSortingEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:261
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSortingEnabled() const

/*

 */
func (this *QTableWidget) IsSortingEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget16isSortingEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtablewidget.h:263
// index:0
// Public Visibility=Default Availability=Available
// [-2] void editItem(QTableWidgetItem *)

/*
Starts editing the item if it is editable.
*/
func (this *QTableWidget) EditItem(item QTableWidgetItem_ITF /*777 QTableWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTableWidgetItem_PTR() != nil {
		convArg0 = item.QTableWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget8editItemEP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:264
// index:0
// Public Visibility=Default Availability=Available
// [-2] void openPersistentEditor(QTableWidgetItem *)

/*
Opens an editor for the give item. The editor remains open after editing.

See also closePersistentEditor().
*/
func (this *QTableWidget) OpenPersistentEditor(item QTableWidgetItem_ITF /*777 QTableWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTableWidgetItem_PTR() != nil {
		convArg0 = item.QTableWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget20openPersistentEditorEP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:265
// index:0
// Public Visibility=Default Availability=Available
// [-2] void closePersistentEditor(QTableWidgetItem *)

/*
Closes the persistent editor for item.

See also openPersistentEditor().
*/
func (this *QTableWidget) ClosePersistentEditor(item QTableWidgetItem_ITF /*777 QTableWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTableWidgetItem_PTR() != nil {
		convArg0 = item.QTableWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget21closePersistentEditorEP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:267
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * cellWidget(int, int) const

/*
Returns the widget displayed in the cell in the given row and column.

Note: The table takes ownership of the widget.

This function was introduced in  Qt 4.1.

See also setCellWidget().
*/
func (this *QTableWidget) CellWidget(row int, column int) *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget10cellWidgetEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtablewidget.h:268
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCellWidget(int, int, QWidget *)

/*
Sets the given widget to be displayed in the cell in the given row and column, passing the ownership of the widget to the table.

If cell widget A is replaced with cell widget B, cell widget A will be deleted. For example, in the code snippet below, the QLineEdit object will be deleted.


  setCellWidget(row, column, new QLineEdit);
  ...
  setCellWidget(row, column, new QTextEdit);



This function was introduced in  Qt 4.1.

See also cellWidget().
*/
func (this *QTableWidget) SetCellWidget(row int, column int, widget QWidget_ITF /*777 QWidget **/) {
	var convArg2 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg2 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget13setCellWidgetEiiP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:269
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void removeCellWidget(int, int)

/*
Removes the widget set on the cell indicated by row and column.

This function was introduced in  Qt 4.3.
*/
func (this *QTableWidget) RemoveCellWidget(row int, column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget16removeCellWidgetEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:271
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isItemSelected(const QTableWidgetItem *) const

/*

 */
func (this *QTableWidget) IsItemSelected(item QTableWidgetItem_ITF /*777 const QTableWidgetItem **/) bool {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTableWidgetItem_PTR() != nil {
		convArg0 = item.QTableWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget14isItemSelectedEPK16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtablewidget.h:272
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemSelected(const QTableWidgetItem *, bool)

/*

 */
func (this *QTableWidget) SetItemSelected(item QTableWidgetItem_ITF /*777 const QTableWidgetItem **/, select_ bool) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTableWidgetItem_PTR() != nil {
		convArg0 = item.QTableWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget15setItemSelectedEPK16QTableWidgetItemb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, select_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:273
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRangeSelected(const QTableWidgetSelectionRange &, bool)

/*
Selects or deselects the range depending on select.
*/
func (this *QTableWidget) SetRangeSelected(range_ QTableWidgetSelectionRange_ITF, select_ bool) {
	var convArg0 unsafe.Pointer
	if range_ != nil && range_.QTableWidgetSelectionRange_PTR() != nil {
		convArg0 = range_.QTableWidgetSelectionRange_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget16setRangeSelectedERK26QTableWidgetSelectionRangeb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, select_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:279
// index:0
// Public Visibility=Default Availability=Available
// [4] int visualRow(int) const

/*
Returns the visual row of the given logicalRow.
*/
func (this *QTableWidget) VisualRow(logicalRow int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget9visualRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalRow)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:280
// index:0
// Public Visibility=Default Availability=Available
// [4] int visualColumn(int) const

/*
Returns the visual column of the given logicalColumn.
*/
func (this *QTableWidget) VisualColumn(logicalColumn int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget12visualColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalColumn)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:282
// index:0
// Public Visibility=Default Availability=Available
// [8] QTableWidgetItem * itemAt(const QPoint &) const

/*
Returns a pointer to the item at the given point, or returns 0 if point is not covered by an item in the table widget.

See also item().
*/
func (this *QTableWidget) ItemAt(p qtcore.QPoint_ITF) *QTableWidgetItem /*777 QTableWidgetItem **/ {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget6itemAtERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtablewidget.h:283
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QTableWidgetItem * itemAt(int, int) const

/*
Returns a pointer to the item at the given point, or returns 0 if point is not covered by an item in the table widget.

See also item().
*/
func (this *QTableWidget) ItemAt1(x int, y int) *QTableWidgetItem /*777 QTableWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget6itemAtEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtablewidget.h:284
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect visualItemRect(const QTableWidgetItem *) const

/*
Returns the rectangle on the viewport occupied by the item at item.
*/
func (this *QTableWidget) VisualItemRect(item QTableWidgetItem_ITF /*777 const QTableWidgetItem **/) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTableWidgetItem_PTR() != nil {
		convArg0 = item.QTableWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget14visualItemRectEPK16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:286
// index:0
// Public Visibility=Default Availability=Available
// [8] const QTableWidgetItem * itemPrototype() const

/*
Returns the item prototype used by the table.

See also setItemPrototype().
*/
func (this *QTableWidget) ItemPrototype() *QTableWidgetItem /*777 const QTableWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget13itemPrototypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtablewidget.h:287
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemPrototype(const QTableWidgetItem *)

/*
Sets the item prototype for the table to the specified item.

The table widget will use the item prototype clone function when it needs to create a new table item. For example when the user is editing in an empty cell. This is useful when you have a QTableWidgetItem subclass and want to make sure that QTableWidget creates instances of your subclass.

The table takes ownership of the prototype.

See also itemPrototype().
*/
func (this *QTableWidget) SetItemPrototype(item QTableWidgetItem_ITF /*777 const QTableWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTableWidgetItem_PTR() != nil {
		convArg0 = item.QTableWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget16setItemPrototypeEPK16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:290
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scrollToItem(const QTableWidgetItem *, QAbstractItemView::ScrollHint)

/*
Scrolls the view if necessary to ensure that the item is visible. The hint parameter specifies more precisely where the item should be located after the operation.
*/
func (this *QTableWidget) ScrollToItem(item QTableWidgetItem_ITF /*777 const QTableWidgetItem **/, hint int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTableWidgetItem_PTR() != nil {
		convArg0 = item.QTableWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget12scrollToItemEPK16QTableWidgetItemN17QAbstractItemView10ScrollHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:290
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scrollToItem(const QTableWidgetItem *, QAbstractItemView::ScrollHint)

/*
Scrolls the view if necessary to ensure that the item is visible. The hint parameter specifies more precisely where the item should be located after the operation.
*/
func (this *QTableWidget) ScrollToItemp(item QTableWidgetItem_ITF /*777 const QTableWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTableWidgetItem_PTR() != nil {
		convArg0 = item.QTableWidgetItem_PTR().GetCthis()
	}
	// arg: 1, QAbstractItemView::ScrollHint=Elaborated, QAbstractItemView::ScrollHint=Enum, , Invalid
	hint := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget12scrollToItemEPK16QTableWidgetItemN17QAbstractItemView10ScrollHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:291
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertRow(int)

/*
Inserts an empty row into the table at row.
*/
func (this *QTableWidget) InsertRow(row int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget9insertRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:292
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertColumn(int)

/*
Inserts an empty column into the table at column.
*/
func (this *QTableWidget) InsertColumn(column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget12insertColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:293
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeRow(int)

/*
Removes the row row and all its items from the table.
*/
func (this *QTableWidget) RemoveRow(row int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget9removeRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:294
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeColumn(int)

/*
Removes the column column and all its items from the table.
*/
func (this *QTableWidget) RemoveColumn(column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget12removeColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:295
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Removes all items in the view. This will also remove all selections and headers. If you don't want to remove the headers, use QTableWidget::clearContents(). The table dimensions stay the same.
*/
func (this *QTableWidget) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:296
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearContents()

/*
Removes all items not in the headers from the view. This will also remove all selections. The table dimensions stay the same.

This function was introduced in  Qt 4.2.
*/
func (this *QTableWidget) ClearContents() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget13clearContentsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:299
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemPressed(QTableWidgetItem *)

/*
This signal is emitted whenever an item in the table is pressed. The item specified is the item that was pressed.
*/
func (this *QTableWidget) ItemPressed(item QTableWidgetItem_ITF /*777 QTableWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTableWidgetItem_PTR() != nil {
		convArg0 = item.QTableWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget11itemPressedEP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:300
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemClicked(QTableWidgetItem *)

/*
This signal is emitted whenever an item in the table is clicked. The item specified is the item that was clicked.
*/
func (this *QTableWidget) ItemClicked(item QTableWidgetItem_ITF /*777 QTableWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTableWidgetItem_PTR() != nil {
		convArg0 = item.QTableWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget11itemClickedEP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:301
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemDoubleClicked(QTableWidgetItem *)

/*
This signal is emitted whenever an item in the table is double clicked. The item specified is the item that was double clicked.
*/
func (this *QTableWidget) ItemDoubleClicked(item QTableWidgetItem_ITF /*777 QTableWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTableWidgetItem_PTR() != nil {
		convArg0 = item.QTableWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget17itemDoubleClickedEP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:303
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemActivated(QTableWidgetItem *)

/*
This signal is emitted when the specified item has been activated
*/
func (this *QTableWidget) ItemActivated(item QTableWidgetItem_ITF /*777 QTableWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTableWidgetItem_PTR() != nil {
		convArg0 = item.QTableWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget13itemActivatedEP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:304
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemEntered(QTableWidgetItem *)

/*
This signal is emitted when the mouse cursor enters an item. The item is the item entered.

This signal is only emitted when mouseTracking is turned on, or when a mouse button is pressed while moving into an item.
*/
func (this *QTableWidget) ItemEntered(item QTableWidgetItem_ITF /*777 QTableWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTableWidgetItem_PTR() != nil {
		convArg0 = item.QTableWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget11itemEnteredEP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:305
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemChanged(QTableWidgetItem *)

/*
This signal is emitted whenever the data of item has changed.
*/
func (this *QTableWidget) ItemChanged(item QTableWidgetItem_ITF /*777 QTableWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTableWidgetItem_PTR() != nil {
		convArg0 = item.QTableWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget11itemChangedEP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:307
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentItemChanged(QTableWidgetItem *, QTableWidgetItem *)

/*
This signal is emitted whenever the current item changes. The previous item is the item that previously had the focus, current is the new current item.
*/
func (this *QTableWidget) CurrentItemChanged(current QTableWidgetItem_ITF /*777 QTableWidgetItem **/, previous QTableWidgetItem_ITF /*777 QTableWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if current != nil && current.QTableWidgetItem_PTR() != nil {
		convArg0 = current.QTableWidgetItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if previous != nil && previous.QTableWidgetItem_PTR() != nil {
		convArg1 = previous.QTableWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget18currentItemChangedEP16QTableWidgetItemS1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:308
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemSelectionChanged()

/*
This signal is emitted whenever the selection changes.

See also selectedItems() and QTableWidgetItem::isSelected().
*/
func (this *QTableWidget) ItemSelectionChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget20itemSelectionChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:310
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cellPressed(int, int)

/*
This signal is emitted whenever a cell in the table is pressed. The row and column specified is the cell that was pressed.

This function was introduced in  Qt 4.1.
*/
func (this *QTableWidget) CellPressed(row int, column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget11cellPressedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:311
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cellClicked(int, int)

/*
This signal is emitted whenever a cell in the table is clicked. The row and column specified is the cell that was clicked.

This function was introduced in  Qt 4.1.
*/
func (this *QTableWidget) CellClicked(row int, column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget11cellClickedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:312
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cellDoubleClicked(int, int)

/*
This signal is emitted whenever a cell in the table is double clicked. The row and column specified is the cell that was double clicked.

This function was introduced in  Qt 4.1.
*/
func (this *QTableWidget) CellDoubleClicked(row int, column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget17cellDoubleClickedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:314
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cellActivated(int, int)

/*
This signal is emitted when the cell specified by row and column has been activated

This function was introduced in  Qt 4.1.
*/
func (this *QTableWidget) CellActivated(row int, column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget13cellActivatedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:315
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cellEntered(int, int)

/*
This signal is emitted when the mouse cursor enters a cell. The cell is specified by row and column.

This signal is only emitted when mouseTracking is turned on, or when a mouse button is pressed while moving into an item.

This function was introduced in  Qt 4.1.
*/
func (this *QTableWidget) CellEntered(row int, column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget11cellEnteredEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:316
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cellChanged(int, int)

/*
This signal is emitted whenever the data of the item in the cell specified by row and column has changed.

This function was introduced in  Qt 4.1.
*/
func (this *QTableWidget) CellChanged(row int, column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget11cellChangedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:318
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentCellChanged(int, int, int, int)

/*
This signal is emitted whenever the current cell changes. The cell specified by previousRow and previousColumn is the cell that previously had the focus, the cell specified by currentRow and currentColumn is the new current cell.

This function was introduced in  Qt 4.1.
*/
func (this *QTableWidget) CurrentCellChanged(currentRow int, currentColumn int, previousRow int, previousColumn int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget18currentCellChangedEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), currentRow, currentColumn, previousRow, previousColumn)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:321
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QTableWidget) Event(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtablewidget.h:322
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QStringList mimeTypes() const

/*
Returns a list of MIME types that can be used to describe a list of tablewidget items.

See also mimeData().
*/
func (this *QTableWidget) MimeTypes() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget9mimeTypesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:328
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool dropMimeData(int, int, const QMimeData *, Qt::DropAction)

/*
Handles the data supplied by a drag and drop operation that ended with the given action in the given row and column. Returns true if the data and action can be handled by the model; otherwise returns false.

See also supportedDropActions().
*/
func (this *QTableWidget) DropMimeData(row int, column int, data qtcore.QMimeData_ITF /*777 const QMimeData **/, action int) bool {
	var convArg2 unsafe.Pointer
	if data != nil && data.QMimeData_PTR() != nil {
		convArg2 = data.QMimeData_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget12dropMimeDataEiiPK9QMimeDataN2Qt10DropActionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2, action)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtablewidget.h:329
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] Qt::DropActions supportedDropActions() const

/*
Returns the drop actions supported by this view.

See also Qt::DropActions.
*/
func (this *QTableWidget) SupportedDropActions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget20supportedDropActionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:332
// index:0
// Protected Visibility=Default Availability=Available
// [24] QModelIndex indexFromItem(QTableWidgetItem *) const

/*
Returns the QModelIndex associated with the given item.
*/
func (this *QTableWidget) IndexFromItem(item QTableWidgetItem_ITF /*777 QTableWidgetItem **/) *qtcore.QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTableWidgetItem_PTR() != nil {
		convArg0 = item.QTableWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget13indexFromItemEP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:333
// index:0
// Protected Visibility=Default Availability=Available
// [8] QTableWidgetItem * itemFromIndex(const QModelIndex &) const

/*
Returns a pointer to the QTableWidgetItem associated with the given index.
*/
func (this *QTableWidget) ItemFromIndex(index qtcore.QModelIndex_ITF) *QTableWidgetItem /*777 QTableWidgetItem **/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget13itemFromIndexERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtablewidget.h:335
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)

/*
Reimplemented from QWidget::dropEvent().
*/
func (this *QTableWidget) DropEvent(event qtgui.QDropEvent_ITF /*777 QDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QDropEvent_PTR() != nil {
		convArg0 = event.QDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget9dropEventEP10QDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
