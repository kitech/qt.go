package qtwidgets

// /usr/include/qt/QtWidgets/qtablewidget.h
// #include <qtablewidget.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 43
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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

// /usr/include/qt/QtWidgets/qtablewidget.h:216
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QTableWidget) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:222
// index:0
// Public
// void QTableWidget(QWidget *)
func NewQTableWidget(parent *QWidget /*777 QWidget **/) *QTableWidget {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidgetC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTableWidgetFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtablewidget.h:223
// index:1
// Public
// void QTableWidget(int, int, QWidget *)
func NewQTableWidget_1(rows int, columns int, parent *QWidget /*777 QWidget **/) *QTableWidget {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidgetC2EiiP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, rows, columns, convArg2)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget11setRowCountEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), rows)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:227
// index:0
// Public
// int rowCount()
func (this *QTableWidget) RowCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget8rowCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:229
// index:0
// Public
// void setColumnCount(int)
func (this *QTableWidget) SetColumnCount(columns int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget14setColumnCountEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), columns)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:230
// index:0
// Public
// int columnCount()
func (this *QTableWidget) ColumnCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget11columnCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:232
// index:0
// Public
// int row(const QTableWidgetItem *)
func (this *QTableWidget) Row(item *QTableWidgetItem /*777 const QTableWidgetItem **/) int {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget3rowEPK16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:233
// index:0
// Public
// int column(const QTableWidgetItem *)
func (this *QTableWidget) Column(item *QTableWidgetItem /*777 const QTableWidgetItem **/) int {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget6columnEPK16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:235
// index:0
// Public
// QTableWidgetItem * item(int, int)
func (this *QTableWidget) Item(row int, column int) *QTableWidgetItem /*777 QTableWidgetItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget4itemEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:236
// index:0
// Public
// void setItem(int, int, QTableWidgetItem *)
func (this *QTableWidget) SetItem(row int, column int, item *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg2 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget7setItemEiiP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:237
// index:0
// Public
// QTableWidgetItem * takeItem(int, int)
func (this *QTableWidget) TakeItem(row int, column int) *QTableWidgetItem /*777 QTableWidgetItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget8takeItemEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:239
// index:0
// Public
// QTableWidgetItem * verticalHeaderItem(int)
func (this *QTableWidget) VerticalHeaderItem(row int) *QTableWidgetItem /*777 QTableWidgetItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget18verticalHeaderItemEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:240
// index:0
// Public
// void setVerticalHeaderItem(int, QTableWidgetItem *)
func (this *QTableWidget) SetVerticalHeaderItem(row int, item *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg1 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget21setVerticalHeaderItemEiP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:241
// index:0
// Public
// QTableWidgetItem * takeVerticalHeaderItem(int)
func (this *QTableWidget) TakeVerticalHeaderItem(row int) *QTableWidgetItem /*777 QTableWidgetItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget22takeVerticalHeaderItemEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:243
// index:0
// Public
// QTableWidgetItem * horizontalHeaderItem(int)
func (this *QTableWidget) HorizontalHeaderItem(column int) *QTableWidgetItem /*777 QTableWidgetItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget20horizontalHeaderItemEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:244
// index:0
// Public
// void setHorizontalHeaderItem(int, QTableWidgetItem *)
func (this *QTableWidget) SetHorizontalHeaderItem(column int, item *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg1 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget23setHorizontalHeaderItemEiP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:245
// index:0
// Public
// QTableWidgetItem * takeHorizontalHeaderItem(int)
func (this *QTableWidget) TakeHorizontalHeaderItem(column int) *QTableWidgetItem /*777 QTableWidgetItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget24takeHorizontalHeaderItemEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:246
// index:0
// Public
// void setVerticalHeaderLabels(const QStringList &)
func (this *QTableWidget) SetVerticalHeaderLabels(labels *qtcore.QStringList) {
	var convArg0 = labels.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget23setVerticalHeaderLabelsERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:247
// index:0
// Public
// void setHorizontalHeaderLabels(const QStringList &)
func (this *QTableWidget) SetHorizontalHeaderLabels(labels *qtcore.QStringList) {
	var convArg0 = labels.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget25setHorizontalHeaderLabelsERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:249
// index:0
// Public
// int currentRow()
func (this *QTableWidget) CurrentRow() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget10currentRowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:250
// index:0
// Public
// int currentColumn()
func (this *QTableWidget) CurrentColumn() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget13currentColumnEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:251
// index:0
// Public
// QTableWidgetItem * currentItem()
func (this *QTableWidget) CurrentItem() *QTableWidgetItem /*777 QTableWidgetItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget11currentItemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:252
// index:0
// Public
// void setCurrentItem(QTableWidgetItem *)
func (this *QTableWidget) SetCurrentItem(item *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget14setCurrentItemEP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:253
// index:1
// Public
// void setCurrentItem(QTableWidgetItem *, QItemSelectionModel::SelectionFlags)
func (this *QTableWidget) SetCurrentItem_1(item *QTableWidgetItem /*777 QTableWidgetItem **/, command int) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget14setCurrentItemEP16QTableWidgetItem6QFlagsIN19QItemSelectionModel13SelectionFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, command)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:254
// index:0
// Public
// void setCurrentCell(int, int)
func (this *QTableWidget) SetCurrentCell(row int, column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget14setCurrentCellEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:255
// index:1
// Public
// void setCurrentCell(int, int, QItemSelectionModel::SelectionFlags)
func (this *QTableWidget) SetCurrentCell_1(row int, column int, command int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget14setCurrentCellEii6QFlagsIN19QItemSelectionModel13SelectionFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column, command)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:257
// index:0
// Public
// void sortItems(int, Qt::SortOrder)
func (this *QTableWidget) SortItems(column int, order int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget9sortItemsEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:258
// index:0
// Public
// void setSortingEnabled(_Bool)
func (this *QTableWidget) SetSortingEnabled(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget17setSortingEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:259
// index:0
// Public
// bool isSortingEnabled()
func (this *QTableWidget) IsSortingEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget16isSortingEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtablewidget.h:261
// index:0
// Public
// void editItem(QTableWidgetItem *)
func (this *QTableWidget) EditItem(item *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget8editItemEP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:262
// index:0
// Public
// void openPersistentEditor(QTableWidgetItem *)
func (this *QTableWidget) OpenPersistentEditor(item *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget20openPersistentEditorEP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:263
// index:0
// Public
// void closePersistentEditor(QTableWidgetItem *)
func (this *QTableWidget) ClosePersistentEditor(item *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget21closePersistentEditorEP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:265
// index:0
// Public
// bool isPersistentEditorOpen(QTableWidgetItem *)
func (this *QTableWidget) IsPersistentEditorOpen(item *QTableWidgetItem /*777 QTableWidgetItem **/) bool {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget22isPersistentEditorOpenEP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtablewidget.h:267
// index:0
// Public
// QWidget * cellWidget(int, int)
func (this *QTableWidget) CellWidget(row int, column int) *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget10cellWidgetEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:268
// index:0
// Public
// void setCellWidget(int, int, QWidget *)
func (this *QTableWidget) SetCellWidget(row int, column int, widget *QWidget /*777 QWidget **/) {
	var convArg2 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget13setCellWidgetEiiP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:269
// index:0
// Public inline
// void removeCellWidget(int, int)
func (this *QTableWidget) RemoveCellWidget(row int, column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget16removeCellWidgetEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:271
// index:0
// Public
// bool isItemSelected(const QTableWidgetItem *)
func (this *QTableWidget) IsItemSelected(item *QTableWidgetItem /*777 const QTableWidgetItem **/) bool {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget14isItemSelectedEPK16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtablewidget.h:272
// index:0
// Public
// void setItemSelected(const QTableWidgetItem *, _Bool)
func (this *QTableWidget) SetItemSelected(item *QTableWidgetItem /*777 const QTableWidgetItem **/, select_ bool) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget15setItemSelectedEPK16QTableWidgetItemb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, select_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:273
// index:0
// Public
// void setRangeSelected(const QTableWidgetSelectionRange &, _Bool)
func (this *QTableWidget) SetRangeSelected(range_ *QTableWidgetSelectionRange, select_ bool) {
	var convArg0 = range_.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget16setRangeSelectedERK26QTableWidgetSelectionRangeb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, select_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:279
// index:0
// Public
// int visualRow(int)
func (this *QTableWidget) VisualRow(logicalRow int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget9visualRowEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), logicalRow)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:280
// index:0
// Public
// int visualColumn(int)
func (this *QTableWidget) VisualColumn(logicalColumn int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget12visualColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), logicalColumn)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:282
// index:0
// Public
// QTableWidgetItem * itemAt(const QPoint &)
func (this *QTableWidget) ItemAt(p *qtcore.QPoint) *QTableWidgetItem /*777 QTableWidgetItem **/ {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget6itemAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:283
// index:1
// Public inline
// QTableWidgetItem * itemAt(int, int)
func (this *QTableWidget) ItemAt_1(x int, y int) *QTableWidgetItem /*777 QTableWidgetItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget6itemAtEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:284
// index:0
// Public
// QRect visualItemRect(const QTableWidgetItem *)
func (this *QTableWidget) VisualItemRect(item *QTableWidgetItem /*777 const QTableWidgetItem **/) *qtcore.QRect /*123*/ {
	var convArg0 = item.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget14visualItemRectEPK16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:286
// index:0
// Public
// const QTableWidgetItem * itemPrototype()
func (this *QTableWidget) ItemPrototype() *QTableWidgetItem /*777 const QTableWidgetItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget13itemPrototypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:287
// index:0
// Public
// void setItemPrototype(const QTableWidgetItem *)
func (this *QTableWidget) SetItemPrototype(item *QTableWidgetItem /*777 const QTableWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget16setItemPrototypeEPK16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:290
// index:0
// Public
// void scrollToItem(const QTableWidgetItem *, QAbstractItemView::ScrollHint)
func (this *QTableWidget) ScrollToItem(item *QTableWidgetItem /*777 const QTableWidgetItem **/, hint int) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget12scrollToItemEPK16QTableWidgetItemN17QAbstractItemView10ScrollHintE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:291
// index:0
// Public
// void insertRow(int)
func (this *QTableWidget) InsertRow(row int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget9insertRowEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:292
// index:0
// Public
// void insertColumn(int)
func (this *QTableWidget) InsertColumn(column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget12insertColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:293
// index:0
// Public
// void removeRow(int)
func (this *QTableWidget) RemoveRow(row int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget9removeRowEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:294
// index:0
// Public
// void removeColumn(int)
func (this *QTableWidget) RemoveColumn(column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget12removeColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
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
// void itemPressed(QTableWidgetItem *)
func (this *QTableWidget) ItemPressed(item *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget11itemPressedEP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:300
// index:0
// Public
// void itemClicked(QTableWidgetItem *)
func (this *QTableWidget) ItemClicked(item *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget11itemClickedEP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:301
// index:0
// Public
// void itemDoubleClicked(QTableWidgetItem *)
func (this *QTableWidget) ItemDoubleClicked(item *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget17itemDoubleClickedEP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:303
// index:0
// Public
// void itemActivated(QTableWidgetItem *)
func (this *QTableWidget) ItemActivated(item *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget13itemActivatedEP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:304
// index:0
// Public
// void itemEntered(QTableWidgetItem *)
func (this *QTableWidget) ItemEntered(item *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget11itemEnteredEP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:305
// index:0
// Public
// void itemChanged(QTableWidgetItem *)
func (this *QTableWidget) ItemChanged(item *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget11itemChangedEP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:307
// index:0
// Public
// void currentItemChanged(QTableWidgetItem *, QTableWidgetItem *)
func (this *QTableWidget) CurrentItemChanged(current *QTableWidgetItem /*777 QTableWidgetItem **/, previous *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg0 = current.GetCthis()
	var convArg1 = previous.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget18currentItemChangedEP16QTableWidgetItemS1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget11cellPressedEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:311
// index:0
// Public
// void cellClicked(int, int)
func (this *QTableWidget) CellClicked(row int, column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget11cellClickedEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:312
// index:0
// Public
// void cellDoubleClicked(int, int)
func (this *QTableWidget) CellDoubleClicked(row int, column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget17cellDoubleClickedEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:314
// index:0
// Public
// void cellActivated(int, int)
func (this *QTableWidget) CellActivated(row int, column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget13cellActivatedEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:315
// index:0
// Public
// void cellEntered(int, int)
func (this *QTableWidget) CellEntered(row int, column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget11cellEnteredEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:316
// index:0
// Public
// void cellChanged(int, int)
func (this *QTableWidget) CellChanged(row int, column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget11cellChangedEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:318
// index:0
// Public
// void currentCellChanged(int, int, int, int)
func (this *QTableWidget) CurrentCellChanged(currentRow int, currentColumn int, previousRow int, previousColumn int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget18currentCellChangedEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), currentRow, currentColumn, previousRow, previousColumn)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:321
// index:0
// Protected virtual
// bool event(QEvent *)
func (this *QTableWidget) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtablewidget.h:328
// index:0
// Protected virtual
// bool dropMimeData(int, int, const QMimeData *, Qt::DropAction)
func (this *QTableWidget) DropMimeData(row int, column int, data *qtcore.QMimeData /*777 const QMimeData **/, action int) bool {
	var convArg2 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget12dropMimeDataEiiPK9QMimeDataN2Qt10DropActionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2, action)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtablewidget.h:329
// index:0
// Protected virtual
// Qt::DropActions supportedDropActions()
func (this *QTableWidget) SupportedDropActions() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget20supportedDropActionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:338
// index:0
// Protected
// QModelIndex indexFromItem(QTableWidgetItem *)
func (this *QTableWidget) IndexFromItem(item *QTableWidgetItem /*777 QTableWidgetItem **/) *qtcore.QModelIndex /*123*/ {
	var convArg0 = item.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget13indexFromItemEP16QTableWidgetItem", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:339
// index:0
// Protected
// QTableWidgetItem * itemFromIndex(const QModelIndex &)
func (this *QTableWidget) ItemFromIndex(index *qtcore.QModelIndex) *QTableWidgetItem /*777 QTableWidgetItem **/ {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTableWidget13itemFromIndexERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:343
// index:0
// Protected virtual
// void dropEvent(QDropEvent *)
func (this *QTableWidget) DropEvent(event *qtgui.QDropEvent /*777 QDropEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTableWidget9dropEventEP10QDropEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
