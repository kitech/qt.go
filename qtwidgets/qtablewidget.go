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
// bool event(class QEvent *)
func (this *QTableWidget) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// bool dropMimeData(int, int, const class QMimeData *, Qt::DropAction)
func (this *QTableWidget) InheritDropMimeData(f func(row int, column int, data *qtcore.QMimeData /*777 const QMimeData **/, action int) bool) {
	qtrt.SetAllInheritCallback(this, "dropMimeData", f)
}

// Qt::DropActions supportedDropActions()
func (this *QTableWidget) InheritSupportedDropActions(f func() int) {
	qtrt.SetAllInheritCallback(this, "supportedDropActions", f)
}

// QModelIndex indexFromItem(class QTableWidgetItem *)
func (this *QTableWidget) InheritIndexFromItem(f func(item *QTableWidgetItem /*777 QTableWidgetItem **/) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "indexFromItem", f)
}

// QTableWidgetItem * itemFromIndex(const class QModelIndex &)
func (this *QTableWidget) InheritItemFromIndex(f func(index *qtcore.QModelIndex) unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "itemFromIndex", f)
}

// void dropEvent(class QDropEvent *)
func (this *QTableWidget) InheritDropEvent(f func(event *qtgui.QDropEvent /*777 QDropEvent **/)) {
	qtrt.SetAllInheritCallback(this, "dropEvent", f)
}

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QTableWidget) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:222
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTableWidget(QWidget *)
func NewQTableWidget(parent *QWidget /*777 QWidget **/) *QTableWidget {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidgetC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTableWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qtablewidget.h:223
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTableWidget(int, int, QWidget *)
func NewQTableWidget_1(rows int, columns int, parent *QWidget /*777 QWidget **/) *QTableWidget {
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidgetC2EiiP7QWidget", qtrt.FFI_TYPE_POINTER, rows, columns, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQTableWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qtablewidget.h:224
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTableWidget()
func DeleteQTableWidget(this *QTableWidget) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidgetD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:226
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowCount(int)
func (this *QTableWidget) SetRowCount(rows int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget11setRowCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rows)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:227
// index:0
// Public Visibility=Default Availability=Available
// [4] int rowCount()
func (this *QTableWidget) RowCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget8rowCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:229
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnCount(int)
func (this *QTableWidget) SetColumnCount(columns int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget14setColumnCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), columns)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:230
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnCount()
func (this *QTableWidget) ColumnCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget11columnCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:232
// index:0
// Public Visibility=Default Availability=Available
// [4] int row(const QTableWidgetItem *)
func (this *QTableWidget) Row(item *QTableWidgetItem /*777 const QTableWidgetItem **/) int {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget3rowEPK16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:233
// index:0
// Public Visibility=Default Availability=Available
// [4] int column(const QTableWidgetItem *)
func (this *QTableWidget) Column(item *QTableWidgetItem /*777 const QTableWidgetItem **/) int {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget6columnEPK16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:235
// index:0
// Public Visibility=Default Availability=Available
// [8] QTableWidgetItem * item(int, int)
func (this *QTableWidget) Item(row int, column int) *QTableWidgetItem /*777 QTableWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget4itemEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:236
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItem(int, int, QTableWidgetItem *)
func (this *QTableWidget) SetItem(row int, column int, item *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg2 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget7setItemEiiP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:237
// index:0
// Public Visibility=Default Availability=Available
// [8] QTableWidgetItem * takeItem(int, int)
func (this *QTableWidget) TakeItem(row int, column int) *QTableWidgetItem /*777 QTableWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget8takeItemEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:239
// index:0
// Public Visibility=Default Availability=Available
// [8] QTableWidgetItem * verticalHeaderItem(int)
func (this *QTableWidget) VerticalHeaderItem(row int) *QTableWidgetItem /*777 QTableWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget18verticalHeaderItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:240
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalHeaderItem(int, QTableWidgetItem *)
func (this *QTableWidget) SetVerticalHeaderItem(row int, item *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg1 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget21setVerticalHeaderItemEiP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:241
// index:0
// Public Visibility=Default Availability=Available
// [8] QTableWidgetItem * takeVerticalHeaderItem(int)
func (this *QTableWidget) TakeVerticalHeaderItem(row int) *QTableWidgetItem /*777 QTableWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget22takeVerticalHeaderItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:243
// index:0
// Public Visibility=Default Availability=Available
// [8] QTableWidgetItem * horizontalHeaderItem(int)
func (this *QTableWidget) HorizontalHeaderItem(column int) *QTableWidgetItem /*777 QTableWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget20horizontalHeaderItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:244
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHorizontalHeaderItem(int, QTableWidgetItem *)
func (this *QTableWidget) SetHorizontalHeaderItem(column int, item *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg1 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget23setHorizontalHeaderItemEiP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:245
// index:0
// Public Visibility=Default Availability=Available
// [8] QTableWidgetItem * takeHorizontalHeaderItem(int)
func (this *QTableWidget) TakeHorizontalHeaderItem(column int) *QTableWidgetItem /*777 QTableWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget24takeHorizontalHeaderItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:246
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalHeaderLabels(const QStringList &)
func (this *QTableWidget) SetVerticalHeaderLabels(labels *qtcore.QStringList) {
	var convArg0 = labels.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget23setVerticalHeaderLabelsERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:247
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHorizontalHeaderLabels(const QStringList &)
func (this *QTableWidget) SetHorizontalHeaderLabels(labels *qtcore.QStringList) {
	var convArg0 = labels.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget25setHorizontalHeaderLabelsERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:249
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentRow()
func (this *QTableWidget) CurrentRow() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget10currentRowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:250
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentColumn()
func (this *QTableWidget) CurrentColumn() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget13currentColumnEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:251
// index:0
// Public Visibility=Default Availability=Available
// [8] QTableWidgetItem * currentItem()
func (this *QTableWidget) CurrentItem() *QTableWidgetItem /*777 QTableWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget11currentItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:252
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentItem(QTableWidgetItem *)
func (this *QTableWidget) SetCurrentItem(item *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget14setCurrentItemEP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:253
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setCurrentItem(QTableWidgetItem *, QItemSelectionModel::SelectionFlags)
func (this *QTableWidget) SetCurrentItem_1(item *QTableWidgetItem /*777 QTableWidgetItem **/, command int) {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget14setCurrentItemEP16QTableWidgetItem6QFlagsIN19QItemSelectionModel13SelectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, command)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:254
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentCell(int, int)
func (this *QTableWidget) SetCurrentCell(row int, column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget14setCurrentCellEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:255
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setCurrentCell(int, int, QItemSelectionModel::SelectionFlags)
func (this *QTableWidget) SetCurrentCell_1(row int, column int, command int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget14setCurrentCellEii6QFlagsIN19QItemSelectionModel13SelectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, command)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:257
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sortItems(int, Qt::SortOrder)
func (this *QTableWidget) SortItems(column int, order int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget9sortItemsEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:258
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSortingEnabled(_Bool)
func (this *QTableWidget) SetSortingEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget17setSortingEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:259
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSortingEnabled()
func (this *QTableWidget) IsSortingEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget16isSortingEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtablewidget.h:261
// index:0
// Public Visibility=Default Availability=Available
// [-2] void editItem(QTableWidgetItem *)
func (this *QTableWidget) EditItem(item *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget8editItemEP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:262
// index:0
// Public Visibility=Default Availability=Available
// [-2] void openPersistentEditor(QTableWidgetItem *)
func (this *QTableWidget) OpenPersistentEditor(item *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget20openPersistentEditorEP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:263
// index:0
// Public Visibility=Default Availability=Available
// [-2] void closePersistentEditor(QTableWidgetItem *)
func (this *QTableWidget) ClosePersistentEditor(item *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget21closePersistentEditorEP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:265
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPersistentEditorOpen(QTableWidgetItem *)
func (this *QTableWidget) IsPersistentEditorOpen(item *QTableWidgetItem /*777 QTableWidgetItem **/) bool {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget22isPersistentEditorOpenEP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtablewidget.h:267
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * cellWidget(int, int)
func (this *QTableWidget) CellWidget(row int, column int) *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget10cellWidgetEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:268
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCellWidget(int, int, QWidget *)
func (this *QTableWidget) SetCellWidget(row int, column int, widget *QWidget /*777 QWidget **/) {
	var convArg2 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget13setCellWidgetEiiP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:269
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void removeCellWidget(int, int)
func (this *QTableWidget) RemoveCellWidget(row int, column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget16removeCellWidgetEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:271
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isItemSelected(const QTableWidgetItem *)
func (this *QTableWidget) IsItemSelected(item *QTableWidgetItem /*777 const QTableWidgetItem **/) bool {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget14isItemSelectedEPK16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtablewidget.h:272
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemSelected(const QTableWidgetItem *, _Bool)
func (this *QTableWidget) SetItemSelected(item *QTableWidgetItem /*777 const QTableWidgetItem **/, select_ bool) {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget15setItemSelectedEPK16QTableWidgetItemb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, select_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:273
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRangeSelected(const QTableWidgetSelectionRange &, _Bool)
func (this *QTableWidget) SetRangeSelected(range_ *QTableWidgetSelectionRange, select_ bool) {
	var convArg0 = range_.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget16setRangeSelectedERK26QTableWidgetSelectionRangeb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, select_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:279
// index:0
// Public Visibility=Default Availability=Available
// [4] int visualRow(int)
func (this *QTableWidget) VisualRow(logicalRow int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget9visualRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalRow)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:280
// index:0
// Public Visibility=Default Availability=Available
// [4] int visualColumn(int)
func (this *QTableWidget) VisualColumn(logicalColumn int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget12visualColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalColumn)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:282
// index:0
// Public Visibility=Default Availability=Available
// [8] QTableWidgetItem * itemAt(const QPoint &)
func (this *QTableWidget) ItemAt(p *qtcore.QPoint) *QTableWidgetItem /*777 QTableWidgetItem **/ {
	var convArg0 = p.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget6itemAtERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:283
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QTableWidgetItem * itemAt(int, int)
func (this *QTableWidget) ItemAt_1(x int, y int) *QTableWidgetItem /*777 QTableWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget6itemAtEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:284
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect visualItemRect(const QTableWidgetItem *)
func (this *QTableWidget) VisualItemRect(item *QTableWidgetItem /*777 const QTableWidgetItem **/) *qtcore.QRect /*123*/ {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget14visualItemRectEPK16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:286
// index:0
// Public Visibility=Default Availability=Available
// [8] const QTableWidgetItem * itemPrototype()
func (this *QTableWidget) ItemPrototype() *QTableWidgetItem /*777 const QTableWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget13itemPrototypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:287
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemPrototype(const QTableWidgetItem *)
func (this *QTableWidget) SetItemPrototype(item *QTableWidgetItem /*777 const QTableWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget16setItemPrototypeEPK16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:290
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scrollToItem(const QTableWidgetItem *, QAbstractItemView::ScrollHint)
func (this *QTableWidget) ScrollToItem(item *QTableWidgetItem /*777 const QTableWidgetItem **/, hint int) {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget12scrollToItemEPK16QTableWidgetItemN17QAbstractItemView10ScrollHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:291
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertRow(int)
func (this *QTableWidget) InsertRow(row int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget9insertRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:292
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertColumn(int)
func (this *QTableWidget) InsertColumn(column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget12insertColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:293
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeRow(int)
func (this *QTableWidget) RemoveRow(row int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget9removeRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:294
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeColumn(int)
func (this *QTableWidget) RemoveColumn(column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget12removeColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:295
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QTableWidget) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:296
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearContents()
func (this *QTableWidget) ClearContents() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget13clearContentsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:299
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemPressed(QTableWidgetItem *)
func (this *QTableWidget) ItemPressed(item *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget11itemPressedEP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:300
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemClicked(QTableWidgetItem *)
func (this *QTableWidget) ItemClicked(item *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget11itemClickedEP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:301
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemDoubleClicked(QTableWidgetItem *)
func (this *QTableWidget) ItemDoubleClicked(item *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget17itemDoubleClickedEP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:303
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemActivated(QTableWidgetItem *)
func (this *QTableWidget) ItemActivated(item *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget13itemActivatedEP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:304
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemEntered(QTableWidgetItem *)
func (this *QTableWidget) ItemEntered(item *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget11itemEnteredEP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:305
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemChanged(QTableWidgetItem *)
func (this *QTableWidget) ItemChanged(item *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget11itemChangedEP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:307
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentItemChanged(QTableWidgetItem *, QTableWidgetItem *)
func (this *QTableWidget) CurrentItemChanged(current *QTableWidgetItem /*777 QTableWidgetItem **/, previous *QTableWidgetItem /*777 QTableWidgetItem **/) {
	var convArg0 = current.GetCthis()
	var convArg1 = previous.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget18currentItemChangedEP16QTableWidgetItemS1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:308
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemSelectionChanged()
func (this *QTableWidget) ItemSelectionChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget20itemSelectionChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:310
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cellPressed(int, int)
func (this *QTableWidget) CellPressed(row int, column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget11cellPressedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:311
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cellClicked(int, int)
func (this *QTableWidget) CellClicked(row int, column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget11cellClickedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:312
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cellDoubleClicked(int, int)
func (this *QTableWidget) CellDoubleClicked(row int, column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget17cellDoubleClickedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:314
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cellActivated(int, int)
func (this *QTableWidget) CellActivated(row int, column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget13cellActivatedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:315
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cellEntered(int, int)
func (this *QTableWidget) CellEntered(row int, column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget11cellEnteredEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:316
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cellChanged(int, int)
func (this *QTableWidget) CellChanged(row int, column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget11cellChangedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:318
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentCellChanged(int, int, int, int)
func (this *QTableWidget) CurrentCellChanged(currentRow int, currentColumn int, previousRow int, previousColumn int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget18currentCellChangedEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), currentRow, currentColumn, previousRow, previousColumn)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:321
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QTableWidget) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtablewidget.h:328
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool dropMimeData(int, int, const QMimeData *, Qt::DropAction)
func (this *QTableWidget) DropMimeData(row int, column int, data *qtcore.QMimeData /*777 const QMimeData **/, action int) bool {
	var convArg2 = data.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget12dropMimeDataEiiPK9QMimeDataN2Qt10DropActionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2, action)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtablewidget.h:329
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] Qt::DropActions supportedDropActions()
func (this *QTableWidget) SupportedDropActions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget20supportedDropActionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:338
// index:0
// Protected Visibility=Default Availability=Available
// [24] QModelIndex indexFromItem(QTableWidgetItem *)
func (this *QTableWidget) IndexFromItem(item *QTableWidgetItem /*777 QTableWidgetItem **/) *qtcore.QModelIndex /*123*/ {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget13indexFromItemEP16QTableWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:339
// index:0
// Protected Visibility=Default Availability=Available
// [8] QTableWidgetItem * itemFromIndex(const QModelIndex &)
func (this *QTableWidget) ItemFromIndex(index *qtcore.QModelIndex) *QTableWidgetItem /*777 QTableWidgetItem **/ {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTableWidget13itemFromIndexERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:343
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)
func (this *QTableWidget) DropEvent(event *qtgui.QDropEvent /*777 QDropEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTableWidget9dropEventEP10QDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
