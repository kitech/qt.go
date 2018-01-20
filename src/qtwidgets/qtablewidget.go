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
	*QTableView
}

func (this *QTableWidget) GetCthis() unsafe.Pointer {
	return this.QTableView.GetCthis()
}
func NewQTableWidgetFromPointer(cthis unsafe.Pointer) *QTableWidget {
	bcthis0 := NewQTableViewFromPointer(cthis)
	return &QTableWidget{bcthis0}
}

// /usr/include/qt/QtWidgets/qtablewidget.h:216
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QTableWidget) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:222
// index:0
// Public
// void QTableWidget(class QWidget *)
func NewQTableWidget(parent unsafe.Pointer) *QTableWidget {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidgetC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQTableWidgetFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtablewidget.h:223
// index:1
// Public
// void QTableWidget(int, int, class QWidget *)
func NewQTableWidget_1(rows int, columns int, parent unsafe.Pointer) *QTableWidget {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidgetC2EiiP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, &rows, &columns, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQTableWidgetFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtablewidget.h:224
// index:0
// Public virtual
// void ~QTableWidget()
func DeleteQTableWidget(*QTableWidget) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidgetD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:226
// index:0
// Public
// void setRowCount(int)
func (this *QTableWidget) SetRowCount(rows int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget11setRowCountEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &rows)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:227
// index:0
// Public
// int rowCount()
func (this *QTableWidget) RowCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget8rowCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:229
// index:0
// Public
// void setColumnCount(int)
func (this *QTableWidget) SetColumnCount(columns int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget14setColumnCountEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &columns)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:230
// index:0
// Public
// int columnCount()
func (this *QTableWidget) ColumnCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget11columnCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:232
// index:0
// Public
// int row(const class QTableWidgetItem *)
func (this *QTableWidget) Row(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget3rowEPK16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:233
// index:0
// Public
// int column(const class QTableWidgetItem *)
func (this *QTableWidget) Column(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget6columnEPK16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:235
// index:0
// Public
// QTableWidgetItem * item(int, int)
func (this *QTableWidget) Item(row int, column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget4itemEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:236
// index:0
// Public
// void setItem(int, int, class QTableWidgetItem *)
func (this *QTableWidget) SetItem(row int, column int, item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget7setItemEiiP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:237
// index:0
// Public
// QTableWidgetItem * takeItem(int, int)
func (this *QTableWidget) TakeItem(row int, column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget8takeItemEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:239
// index:0
// Public
// QTableWidgetItem * verticalHeaderItem(int)
func (this *QTableWidget) VerticalHeaderItem(row int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget18verticalHeaderItemEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:240
// index:0
// Public
// void setVerticalHeaderItem(int, class QTableWidgetItem *)
func (this *QTableWidget) SetVerticalHeaderItem(row int, item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget21setVerticalHeaderItemEiP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:241
// index:0
// Public
// QTableWidgetItem * takeVerticalHeaderItem(int)
func (this *QTableWidget) TakeVerticalHeaderItem(row int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget22takeVerticalHeaderItemEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:243
// index:0
// Public
// QTableWidgetItem * horizontalHeaderItem(int)
func (this *QTableWidget) HorizontalHeaderItem(column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget20horizontalHeaderItemEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:244
// index:0
// Public
// void setHorizontalHeaderItem(int, class QTableWidgetItem *)
func (this *QTableWidget) SetHorizontalHeaderItem(column int, item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget23setHorizontalHeaderItemEiP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:245
// index:0
// Public
// QTableWidgetItem * takeHorizontalHeaderItem(int)
func (this *QTableWidget) TakeHorizontalHeaderItem(column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget24takeHorizontalHeaderItemEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:246
// index:0
// Public
// void setVerticalHeaderLabels(const class QStringList &)
func (this *QTableWidget) SetVerticalHeaderLabels(labels *qtcore.QStringList) {
	var convArg0 = labels.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget23setVerticalHeaderLabelsERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:247
// index:0
// Public
// void setHorizontalHeaderLabels(const class QStringList &)
func (this *QTableWidget) SetHorizontalHeaderLabels(labels *qtcore.QStringList) {
	var convArg0 = labels.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget25setHorizontalHeaderLabelsERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:249
// index:0
// Public
// int currentRow()
func (this *QTableWidget) CurrentRow() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget10currentRowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:250
// index:0
// Public
// int currentColumn()
func (this *QTableWidget) CurrentColumn() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget13currentColumnEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:251
// index:0
// Public
// QTableWidgetItem * currentItem()
func (this *QTableWidget) CurrentItem() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget11currentItemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:252
// index:0
// Public
// void setCurrentItem(class QTableWidgetItem *)
func (this *QTableWidget) SetCurrentItem(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget14setCurrentItemEP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:254
// index:0
// Public
// void setCurrentCell(int, int)
func (this *QTableWidget) SetCurrentCell(row int, column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget14setCurrentCellEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:257
// index:0
// Public
// void sortItems(int, Qt::SortOrder)
func (this *QTableWidget) SortItems(column int, order int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget9sortItemsEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:258
// index:0
// Public
// void setSortingEnabled(_Bool)
func (this *QTableWidget) SetSortingEnabled(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget17setSortingEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:259
// index:0
// Public
// bool isSortingEnabled()
func (this *QTableWidget) IsSortingEnabled() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget16isSortingEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:261
// index:0
// Public
// void editItem(class QTableWidgetItem *)
func (this *QTableWidget) EditItem(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget8editItemEP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:262
// index:0
// Public
// void openPersistentEditor(class QTableWidgetItem *)
func (this *QTableWidget) OpenPersistentEditor(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget20openPersistentEditorEP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:263
// index:0
// Public
// void closePersistentEditor(class QTableWidgetItem *)
func (this *QTableWidget) ClosePersistentEditor(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget21closePersistentEditorEP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:265
// index:0
// Public
// bool isPersistentEditorOpen(class QTableWidgetItem *)
func (this *QTableWidget) IsPersistentEditorOpen(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget22isPersistentEditorOpenEP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:267
// index:0
// Public
// QWidget * cellWidget(int, int)
func (this *QTableWidget) CellWidget(row int, column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget10cellWidgetEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:268
// index:0
// Public
// void setCellWidget(int, int, class QWidget *)
func (this *QTableWidget) SetCellWidget(row int, column int, widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget13setCellWidgetEiiP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:269
// index:0
// Public inline
// void removeCellWidget(int, int)
func (this *QTableWidget) RemoveCellWidget(row int, column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget16removeCellWidgetEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:271
// index:0
// Public
// bool isItemSelected(const class QTableWidgetItem *)
func (this *QTableWidget) IsItemSelected(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget14isItemSelectedEPK16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:272
// index:0
// Public
// void setItemSelected(const class QTableWidgetItem *, _Bool)
func (this *QTableWidget) SetItemSelected(item unsafe.Pointer, select_ bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget15setItemSelectedEPK16QTableWidgetItemb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &select_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:273
// index:0
// Public
// void setRangeSelected(const class QTableWidgetSelectionRange &, _Bool)
func (this *QTableWidget) SetRangeSelected(range_ *QTableWidgetSelectionRange, select_ bool) {
	var convArg0 = range_.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget16setRangeSelectedERK26QTableWidgetSelectionRangeb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &select_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:275
// index:0
// Public
// QList<QTableWidgetSelectionRange> selectedRanges()
func (this *QTableWidget) SelectedRanges() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget14selectedRangesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:276
// index:0
// Public
// QList<QTableWidgetItem *> selectedItems()
func (this *QTableWidget) SelectedItems() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget13selectedItemsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:279
// index:0
// Public
// int visualRow(int)
func (this *QTableWidget) VisualRow(logicalRow int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget9visualRowEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &logicalRow)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:280
// index:0
// Public
// int visualColumn(int)
func (this *QTableWidget) VisualColumn(logicalColumn int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget12visualColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &logicalColumn)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:282
// index:0
// Public
// QTableWidgetItem * itemAt(const class QPoint &)
func (this *QTableWidget) ItemAt(p *qtcore.QPoint) interface{} {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget6itemAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:283
// index:1
// Public inline
// QTableWidgetItem * itemAt(int, int)
func (this *QTableWidget) ItemAt_1(x int, y int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget6itemAtEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:284
// index:0
// Public
// QRect visualItemRect(const class QTableWidgetItem *)
func (this *QTableWidget) VisualItemRect(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget14visualItemRectEPK16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:286
// index:0
// Public
// const QTableWidgetItem * itemPrototype()
func (this *QTableWidget) ItemPrototype() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget13itemPrototypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:287
// index:0
// Public
// void setItemPrototype(const class QTableWidgetItem *)
func (this *QTableWidget) SetItemPrototype(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget16setItemPrototypeEPK16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:290
// index:0
// Public
// void scrollToItem(const class QTableWidgetItem *, class QAbstractItemView::ScrollHint)
func (this *QTableWidget) ScrollToItem(item unsafe.Pointer, hint int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget12scrollToItemEPK16QTableWidgetItemN17QAbstractItemView10ScrollHintE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:291
// index:0
// Public
// void insertRow(int)
func (this *QTableWidget) InsertRow(row int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget9insertRowEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:292
// index:0
// Public
// void insertColumn(int)
func (this *QTableWidget) InsertColumn(column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget12insertColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:293
// index:0
// Public
// void removeRow(int)
func (this *QTableWidget) RemoveRow(row int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget9removeRowEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:294
// index:0
// Public
// void removeColumn(int)
func (this *QTableWidget) RemoveColumn(column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget12removeColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:295
// index:0
// Public
// void clear()
func (this *QTableWidget) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:296
// index:0
// Public
// void clearContents()
func (this *QTableWidget) ClearContents() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget13clearContentsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:299
// index:0
// Public
// void itemPressed(class QTableWidgetItem *)
func (this *QTableWidget) ItemPressed(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget11itemPressedEP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:300
// index:0
// Public
// void itemClicked(class QTableWidgetItem *)
func (this *QTableWidget) ItemClicked(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget11itemClickedEP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:301
// index:0
// Public
// void itemDoubleClicked(class QTableWidgetItem *)
func (this *QTableWidget) ItemDoubleClicked(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget17itemDoubleClickedEP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:303
// index:0
// Public
// void itemActivated(class QTableWidgetItem *)
func (this *QTableWidget) ItemActivated(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget13itemActivatedEP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:304
// index:0
// Public
// void itemEntered(class QTableWidgetItem *)
func (this *QTableWidget) ItemEntered(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget11itemEnteredEP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:305
// index:0
// Public
// void itemChanged(class QTableWidgetItem *)
func (this *QTableWidget) ItemChanged(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget11itemChangedEP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:307
// index:0
// Public
// void currentItemChanged(class QTableWidgetItem *, class QTableWidgetItem *)
func (this *QTableWidget) CurrentItemChanged(current unsafe.Pointer, previous unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget18currentItemChangedEP16QTableWidgetItemS1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), current, previous)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:308
// index:0
// Public
// void itemSelectionChanged()
func (this *QTableWidget) ItemSelectionChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget20itemSelectionChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:310
// index:0
// Public
// void cellPressed(int, int)
func (this *QTableWidget) CellPressed(row int, column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget11cellPressedEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:311
// index:0
// Public
// void cellClicked(int, int)
func (this *QTableWidget) CellClicked(row int, column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget11cellClickedEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:312
// index:0
// Public
// void cellDoubleClicked(int, int)
func (this *QTableWidget) CellDoubleClicked(row int, column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget17cellDoubleClickedEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:314
// index:0
// Public
// void cellActivated(int, int)
func (this *QTableWidget) CellActivated(row int, column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget13cellActivatedEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:315
// index:0
// Public
// void cellEntered(int, int)
func (this *QTableWidget) CellEntered(row int, column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget11cellEnteredEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:316
// index:0
// Public
// void cellChanged(int, int)
func (this *QTableWidget) CellChanged(row int, column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget11cellChangedEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:318
// index:0
// Public
// void currentCellChanged(int, int, int, int)
func (this *QTableWidget) CurrentCellChanged(currentRow int, currentColumn int, previousRow int, previousColumn int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget18currentCellChangedEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &currentRow, &currentColumn, &previousRow, &previousColumn)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:321
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QTableWidget) Event(e unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:322
// index:0
// Protected virtual
// QStringList mimeTypes()
func (this *QTableWidget) MimeTypes() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget9mimeTypesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:328
// index:0
// Protected virtual
// bool dropMimeData(int, int, const class QMimeData *, Qt::DropAction)
func (this *QTableWidget) DropMimeData(row int, column int, data unsafe.Pointer, action int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget12dropMimeDataEiiPK9QMimeDataN2Qt10DropActionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column, data, &action)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:329
// index:0
// Protected virtual
// Qt::DropActions supportedDropActions()
func (this *QTableWidget) SupportedDropActions() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget20supportedDropActionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:336
// index:0
// Protected
// QList<QTableWidgetItem *> items(const class QMimeData *)
func (this *QTableWidget) Items(data unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget5itemsEPK9QMimeData", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), data)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:338
// index:0
// Protected
// QModelIndex indexFromItem(class QTableWidgetItem *)
func (this *QTableWidget) IndexFromItem(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget13indexFromItemEP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:339
// index:0
// Protected
// QTableWidgetItem * itemFromIndex(const class QModelIndex &)
func (this *QTableWidget) ItemFromIndex(index *qtcore.QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget13itemFromIndexERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:343
// index:0
// Protected virtual
// void dropEvent(class QDropEvent *)
func (this *QTableWidget) DropEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget9dropEventEP10QDropEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

//  body block end
