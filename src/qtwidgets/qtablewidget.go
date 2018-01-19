//  header block begin
// /usr/include/qt/QtWidgets/qtablewidget.h
// #include <qtablewidget.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 42
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
type QTableWidget struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qtablewidget.h:216
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QTableWidget) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:222
// index:0
// void QTableWidget(class QWidget *)
func NewQTableWidget(parent unsafe.Pointer) *QTableWidget {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidgetC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QTableWidget{cthis}
}

// /usr/include/qt/QtWidgets/qtablewidget.h:223
// index:1
// void QTableWidget(int, int, class QWidget *)
func NewQTableWidget_1(rows int, columns int, parent unsafe.Pointer) *QTableWidget {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidgetC2EiiP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, &rows, &columns, parent)
	gopp.ErrPrint(err, rv)
	return &QTableWidget{cthis}
}

// /usr/include/qt/QtWidgets/qtablewidget.h:224
// index:0
// virtual
// void ~QTableWidget()
func DeleteQTableWidget(*QTableWidget) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidgetD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:226
// index:0
// void setRowCount(int)
func (this *QTableWidget) SetRowCount(rows int) {
	// 0: (, int rows), (&rows)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget11setRowCountEi", ffiqt.FFI_TYPE_VOID, this.cthis, &rows)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:227
// index:0
// int rowCount()
func (this *QTableWidget) RowCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget8rowCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:229
// index:0
// void setColumnCount(int)
func (this *QTableWidget) SetColumnCount(columns int) {
	// 0: (, int columns), (&columns)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget14setColumnCountEi", ffiqt.FFI_TYPE_VOID, this.cthis, &columns)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:230
// index:0
// int columnCount()
func (this *QTableWidget) ColumnCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget11columnCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:232
// index:0
// int row(const class QTableWidgetItem *)
func (this *QTableWidget) Row(item unsafe.Pointer) {
	// 0: (, const QTableWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget3rowEPK16QTableWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:233
// index:0
// int column(const class QTableWidgetItem *)
func (this *QTableWidget) Column(item unsafe.Pointer) {
	// 0: (, const QTableWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget6columnEPK16QTableWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:235
// index:0
// QTableWidgetItem * item(int, int)
func (this *QTableWidget) Item(row int, column int) {
	// 0: (, int row, int column), (&row, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget4itemEii", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:236
// index:0
// void setItem(int, int, class QTableWidgetItem *)
func (this *QTableWidget) SetItem(row int, column int, item unsafe.Pointer) {
	// 0: (, int row, int column, QTableWidgetItem * item), (&row, &column, item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget7setItemEiiP16QTableWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:237
// index:0
// QTableWidgetItem * takeItem(int, int)
func (this *QTableWidget) TakeItem(row int, column int) {
	// 0: (, int row, int column), (&row, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget8takeItemEii", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:239
// index:0
// QTableWidgetItem * verticalHeaderItem(int)
func (this *QTableWidget) VerticalHeaderItem(row int) {
	// 0: (, int row), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget18verticalHeaderItemEi", ffiqt.FFI_TYPE_VOID, this.cthis, &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:240
// index:0
// void setVerticalHeaderItem(int, class QTableWidgetItem *)
func (this *QTableWidget) SetVerticalHeaderItem(row int, item unsafe.Pointer) {
	// 0: (, int row, QTableWidgetItem * item), (&row, item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget21setVerticalHeaderItemEiP16QTableWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, &row, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:241
// index:0
// QTableWidgetItem * takeVerticalHeaderItem(int)
func (this *QTableWidget) TakeVerticalHeaderItem(row int) {
	// 0: (, int row), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget22takeVerticalHeaderItemEi", ffiqt.FFI_TYPE_VOID, this.cthis, &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:243
// index:0
// QTableWidgetItem * horizontalHeaderItem(int)
func (this *QTableWidget) HorizontalHeaderItem(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget20horizontalHeaderItemEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:244
// index:0
// void setHorizontalHeaderItem(int, class QTableWidgetItem *)
func (this *QTableWidget) SetHorizontalHeaderItem(column int, item unsafe.Pointer) {
	// 0: (, int column, QTableWidgetItem * item), (&column, item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget23setHorizontalHeaderItemEiP16QTableWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, &column, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:245
// index:0
// QTableWidgetItem * takeHorizontalHeaderItem(int)
func (this *QTableWidget) TakeHorizontalHeaderItem(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget24takeHorizontalHeaderItemEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:246
// index:0
// void setVerticalHeaderLabels(const class QStringList &)
func (this *QTableWidget) SetVerticalHeaderLabels(labels unsafe.Pointer) {
	// 0: (, const QStringList & labels), (labels)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget23setVerticalHeaderLabelsERK11QStringList", ffiqt.FFI_TYPE_VOID, this.cthis, labels)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:247
// index:0
// void setHorizontalHeaderLabels(const class QStringList &)
func (this *QTableWidget) SetHorizontalHeaderLabels(labels unsafe.Pointer) {
	// 0: (, const QStringList & labels), (labels)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget25setHorizontalHeaderLabelsERK11QStringList", ffiqt.FFI_TYPE_VOID, this.cthis, labels)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:249
// index:0
// int currentRow()
func (this *QTableWidget) CurrentRow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget10currentRowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:250
// index:0
// int currentColumn()
func (this *QTableWidget) CurrentColumn() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget13currentColumnEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:251
// index:0
// QTableWidgetItem * currentItem()
func (this *QTableWidget) CurrentItem() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget11currentItemEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:252
// index:0
// void setCurrentItem(class QTableWidgetItem *)
func (this *QTableWidget) SetCurrentItem(item unsafe.Pointer) {
	// 0: (, QTableWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget14setCurrentItemEP16QTableWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:254
// index:0
// void setCurrentCell(int, int)
func (this *QTableWidget) SetCurrentCell(row int, column int) {
	// 0: (, int row, int column), (&row, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget14setCurrentCellEii", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:257
// index:0
// void sortItems(int, Qt::SortOrder)
func (this *QTableWidget) SortItems(column int, order int) {
	// 0: (, int column, Qt::SortOrder order), (&column, &order)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget9sortItemsEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:258
// index:0
// void setSortingEnabled(_Bool)
func (this *QTableWidget) SetSortingEnabled(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget17setSortingEnabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:259
// index:0
// bool isSortingEnabled()
func (this *QTableWidget) IsSortingEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget16isSortingEnabledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:261
// index:0
// void editItem(class QTableWidgetItem *)
func (this *QTableWidget) EditItem(item unsafe.Pointer) {
	// 0: (, QTableWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget8editItemEP16QTableWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:262
// index:0
// void openPersistentEditor(class QTableWidgetItem *)
func (this *QTableWidget) OpenPersistentEditor(item unsafe.Pointer) {
	// 0: (, QTableWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget20openPersistentEditorEP16QTableWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:263
// index:0
// void closePersistentEditor(class QTableWidgetItem *)
func (this *QTableWidget) ClosePersistentEditor(item unsafe.Pointer) {
	// 0: (, QTableWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget21closePersistentEditorEP16QTableWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:265
// index:0
// bool isPersistentEditorOpen(class QTableWidgetItem *)
func (this *QTableWidget) IsPersistentEditorOpen(item unsafe.Pointer) {
	// 0: (, QTableWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget22isPersistentEditorOpenEP16QTableWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:267
// index:0
// QWidget * cellWidget(int, int)
func (this *QTableWidget) CellWidget(row int, column int) {
	// 0: (, int row, int column), (&row, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget10cellWidgetEii", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:268
// index:0
// void setCellWidget(int, int, class QWidget *)
func (this *QTableWidget) SetCellWidget(row int, column int, widget unsafe.Pointer) {
	// 0: (, int row, int column, QWidget * widget), (&row, &column, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget13setCellWidgetEiiP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:269
// index:0
// inline
// void removeCellWidget(int, int)
func (this *QTableWidget) RemoveCellWidget(row int, column int) {
	// 0: (, int row, int column), (&row, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget16removeCellWidgetEii", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:271
// index:0
// bool isItemSelected(const class QTableWidgetItem *)
func (this *QTableWidget) IsItemSelected(item unsafe.Pointer) {
	// 0: (, const QTableWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget14isItemSelectedEPK16QTableWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:272
// index:0
// void setItemSelected(const class QTableWidgetItem *, _Bool)
func (this *QTableWidget) SetItemSelected(item unsafe.Pointer, select_ bool) {
	// 0: (, const QTableWidgetItem * item, bool select), (item, &select_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget15setItemSelectedEPK16QTableWidgetItemb", ffiqt.FFI_TYPE_VOID, this.cthis, item, &select_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:273
// index:0
// void setRangeSelected(const class QTableWidgetSelectionRange &, _Bool)
func (this *QTableWidget) SetRangeSelected(range_ unsafe.Pointer, select_ bool) {
	// 0: (, const QTableWidgetSelectionRange & range, bool select), (range_, &select_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget16setRangeSelectedERK26QTableWidgetSelectionRangeb", ffiqt.FFI_TYPE_VOID, this.cthis, range_, &select_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:275
// index:0
// QList<QTableWidgetSelectionRange> selectedRanges()
func (this *QTableWidget) SelectedRanges() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget14selectedRangesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:276
// index:0
// QList<QTableWidgetItem *> selectedItems()
func (this *QTableWidget) SelectedItems() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget13selectedItemsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:279
// index:0
// int visualRow(int)
func (this *QTableWidget) VisualRow(logicalRow int) {
	// 0: (, int logicalRow), (&logicalRow)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget9visualRowEi", ffiqt.FFI_TYPE_VOID, this.cthis, &logicalRow)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:280
// index:0
// int visualColumn(int)
func (this *QTableWidget) VisualColumn(logicalColumn int) {
	// 0: (, int logicalColumn), (&logicalColumn)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget12visualColumnEi", ffiqt.FFI_TYPE_VOID, this.cthis, &logicalColumn)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:282
// index:0
// QTableWidgetItem * itemAt(const class QPoint &)
func (this *QTableWidget) ItemAt(p unsafe.Pointer) {
	// 0: (, const QPoint & p), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget6itemAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:283
// index:1
// inline
// QTableWidgetItem * itemAt(int, int)
func (this *QTableWidget) ItemAt_1(x int, y int) {
	// 1: (, int x, int y), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget6itemAtEii", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:284
// index:0
// QRect visualItemRect(const class QTableWidgetItem *)
func (this *QTableWidget) VisualItemRect(item unsafe.Pointer) {
	// 0: (, const QTableWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget14visualItemRectEPK16QTableWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:286
// index:0
// const QTableWidgetItem * itemPrototype()
func (this *QTableWidget) ItemPrototype() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget13itemPrototypeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:287
// index:0
// void setItemPrototype(const class QTableWidgetItem *)
func (this *QTableWidget) SetItemPrototype(item unsafe.Pointer) {
	// 0: (, const QTableWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget16setItemPrototypeEPK16QTableWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:290
// index:0
// void scrollToItem(const class QTableWidgetItem *, class QAbstractItemView::ScrollHint)
func (this *QTableWidget) ScrollToItem(item unsafe.Pointer, hint int) {
	// 0: (, const QTableWidgetItem * item, QAbstractItemView::ScrollHint hint), (item, &hint)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget12scrollToItemEPK16QTableWidgetItemN17QAbstractItemView10ScrollHintE", ffiqt.FFI_TYPE_VOID, this.cthis, item, &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:291
// index:0
// void insertRow(int)
func (this *QTableWidget) InsertRow(row int) {
	// 0: (, int row), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget9insertRowEi", ffiqt.FFI_TYPE_VOID, this.cthis, &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:292
// index:0
// void insertColumn(int)
func (this *QTableWidget) InsertColumn(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget12insertColumnEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:293
// index:0
// void removeRow(int)
func (this *QTableWidget) RemoveRow(row int) {
	// 0: (, int row), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget9removeRowEi", ffiqt.FFI_TYPE_VOID, this.cthis, &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:294
// index:0
// void removeColumn(int)
func (this *QTableWidget) RemoveColumn(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget12removeColumnEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:295
// index:0
// void clear()
func (this *QTableWidget) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget5clearEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:296
// index:0
// void clearContents()
func (this *QTableWidget) ClearContents() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget13clearContentsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:299
// index:0
// void itemPressed(class QTableWidgetItem *)
func (this *QTableWidget) ItemPressed(item unsafe.Pointer) {
	// 0: (, QTableWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget11itemPressedEP16QTableWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:300
// index:0
// void itemClicked(class QTableWidgetItem *)
func (this *QTableWidget) ItemClicked(item unsafe.Pointer) {
	// 0: (, QTableWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget11itemClickedEP16QTableWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:301
// index:0
// void itemDoubleClicked(class QTableWidgetItem *)
func (this *QTableWidget) ItemDoubleClicked(item unsafe.Pointer) {
	// 0: (, QTableWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget17itemDoubleClickedEP16QTableWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:303
// index:0
// void itemActivated(class QTableWidgetItem *)
func (this *QTableWidget) ItemActivated(item unsafe.Pointer) {
	// 0: (, QTableWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget13itemActivatedEP16QTableWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:304
// index:0
// void itemEntered(class QTableWidgetItem *)
func (this *QTableWidget) ItemEntered(item unsafe.Pointer) {
	// 0: (, QTableWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget11itemEnteredEP16QTableWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:305
// index:0
// void itemChanged(class QTableWidgetItem *)
func (this *QTableWidget) ItemChanged(item unsafe.Pointer) {
	// 0: (, QTableWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget11itemChangedEP16QTableWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:307
// index:0
// void currentItemChanged(class QTableWidgetItem *, class QTableWidgetItem *)
func (this *QTableWidget) CurrentItemChanged(current unsafe.Pointer, previous unsafe.Pointer) {
	// 0: (, QTableWidgetItem * current, QTableWidgetItem * previous), (current, previous)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget18currentItemChangedEP16QTableWidgetItemS1_", ffiqt.FFI_TYPE_VOID, this.cthis, current, previous)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:308
// index:0
// void itemSelectionChanged()
func (this *QTableWidget) ItemSelectionChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget20itemSelectionChangedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:310
// index:0
// void cellPressed(int, int)
func (this *QTableWidget) CellPressed(row int, column int) {
	// 0: (, int row, int column), (&row, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget11cellPressedEii", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:311
// index:0
// void cellClicked(int, int)
func (this *QTableWidget) CellClicked(row int, column int) {
	// 0: (, int row, int column), (&row, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget11cellClickedEii", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:312
// index:0
// void cellDoubleClicked(int, int)
func (this *QTableWidget) CellDoubleClicked(row int, column int) {
	// 0: (, int row, int column), (&row, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget17cellDoubleClickedEii", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:314
// index:0
// void cellActivated(int, int)
func (this *QTableWidget) CellActivated(row int, column int) {
	// 0: (, int row, int column), (&row, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget13cellActivatedEii", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:315
// index:0
// void cellEntered(int, int)
func (this *QTableWidget) CellEntered(row int, column int) {
	// 0: (, int row, int column), (&row, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget11cellEnteredEii", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:316
// index:0
// void cellChanged(int, int)
func (this *QTableWidget) CellChanged(row int, column int) {
	// 0: (, int row, int column), (&row, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget11cellChangedEii", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:318
// index:0
// void currentCellChanged(int, int, int, int)
func (this *QTableWidget) CurrentCellChanged(currentRow int, currentColumn int, previousRow int, previousColumn int) {
	// 0: (, int currentRow, int currentColumn, int previousRow, int previousColumn), (&currentRow, &currentColumn, &previousRow, &previousColumn)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget18currentCellChangedEiiii", ffiqt.FFI_TYPE_VOID, this.cthis, &currentRow, &currentColumn, &previousRow, &previousColumn)
	gopp.ErrPrint(err, rv)
}

//  body block end
