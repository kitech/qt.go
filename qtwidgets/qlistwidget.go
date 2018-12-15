package qtwidgets

// /usr/include/qt/QtWidgets/qlistwidget.h
// #include <qlistwidget.h>
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
func (this *QListWidget) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// QStringList mimeTypes()
func (this *QListWidget) InheritMimeTypes(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "mimeTypes", f)
}

// bool dropMimeData(int, const QMimeData *, Qt::DropAction)
func (this *QListWidget) InheritDropMimeData(f func(index int, data *qtcore.QMimeData /*777 const QMimeData **/, action int) bool) {
	qtrt.SetAllInheritCallback(this, "dropMimeData", f)
}

// Qt::DropActions supportedDropActions()
func (this *QListWidget) InheritSupportedDropActions(f func() int) {
	qtrt.SetAllInheritCallback(this, "supportedDropActions", f)
}

// QModelIndex indexFromItem(QListWidgetItem *)
func (this *QListWidget) InheritIndexFromItem(f func(item *QListWidgetItem /*777 QListWidgetItem **/) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "indexFromItem", f)
}

// QListWidgetItem * itemFromIndex(const QModelIndex &)
func (this *QListWidget) InheritItemFromIndex(f func(index *qtcore.QModelIndex) unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "itemFromIndex", f)
}

/*

 */
type QListWidget struct {
	*QListView
}
type QListWidget_ITF interface {
	QListView_ITF
	QListWidget_PTR() *QListWidget
}

func (ptr *QListWidget) QListWidget_PTR() *QListWidget { return ptr }

func (this *QListWidget) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QListView.GetCthis()
	}
}
func (this *QListWidget) SetCthis(cthis unsafe.Pointer) {
	this.QListView = NewQListViewFromPointer(cthis)
}
func NewQListWidgetFromPointer(cthis unsafe.Pointer) *QListWidget {
	bcthis0 := NewQListViewFromPointer(cthis)
	return &QListWidget{bcthis0}
}
func (*QListWidget) NewFromPointer(cthis unsafe.Pointer) *QListWidget {
	return NewQListWidgetFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:200
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QListWidget) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlistwidget.h:208
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QListWidget(QWidget *)

/*
Constructs an empty QListWidget with the given parent.
*/
func (*QListWidget) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QListWidget {
	return NewQListWidget(parent)
}
func NewQListWidget(parent QWidget_ITF /*777 QWidget **/) *QListWidget {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidgetC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQListWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QListWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qlistwidget.h:208
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QListWidget(QWidget *)

/*
Constructs an empty QListWidget with the given parent.
*/
func (*QListWidget) NewForInheritp() *QListWidget {
	return NewQListWidgetp()
}
func NewQListWidgetp() *QListWidget {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidgetC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQListWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QListWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qlistwidget.h:209
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QListWidget()

/*

 */
func DeleteQListWidget(this *QListWidget) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidgetD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:211
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setSelectionModel(QItemSelectionModel *)

/*
Reimplemented from QAbstractItemView::setSelectionModel().
*/
func (this *QListWidget) SetSelectionModel(selectionModel qtcore.QItemSelectionModel_ITF /*777 QItemSelectionModel **/) {
	var convArg0 unsafe.Pointer
	if selectionModel != nil && selectionModel.QItemSelectionModel_PTR() != nil {
		convArg0 = selectionModel.QItemSelectionModel_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget17setSelectionModelEP19QItemSelectionModel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:213
// index:0
// Public Visibility=Default Availability=Available
// [8] QListWidgetItem * item(int) const

/*
Returns the item that occupies the given row in the list if one has been set; otherwise returns 0.

See also row().
*/
func (this *QListWidget) Item(row int) *QListWidgetItem /*777 QListWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget4itemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQListWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlistwidget.h:214
// index:0
// Public Visibility=Default Availability=Available
// [4] int row(const QListWidgetItem *) const

/*
Returns the row containing the given item.

See also item().
*/
func (this *QListWidget) Row(item QListWidgetItem_ITF /*777 const QListWidgetItem **/) int {
	var convArg0 unsafe.Pointer
	if item != nil && item.QListWidgetItem_PTR() != nil {
		convArg0 = item.QListWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget3rowEPK15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlistwidget.h:215
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertItem(int, QListWidgetItem *)

/*
Inserts the item at the position in the list given by row.

See also addItem().
*/
func (this *QListWidget) InsertItem(row int, item QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg1 unsafe.Pointer
	if item != nil && item.QListWidgetItem_PTR() != nil {
		convArg1 = item.QListWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget10insertItemEiP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:216
// index:1
// Public Visibility=Default Availability=Available
// [-2] void insertItem(int, const QString &)

/*
Inserts the item at the position in the list given by row.

See also addItem().
*/
func (this *QListWidget) InsertItem1(row int, label string) {
	var tmpArg1 = qtcore.NewQString5(label)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget10insertItemEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:217
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertItems(int, const QStringList &)

/*
Inserts items from the list of labels into the list, starting at the given row.

See also insertItem() and addItem().
*/
func (this *QListWidget) InsertItems(row int, labels qtcore.QStringList_ITF) {
	var convArg1 unsafe.Pointer
	if labels != nil && labels.QStringList_PTR() != nil {
		convArg1 = labels.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget11insertItemsEiRK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:218
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void addItem(const QString &)

/*
Inserts an item with the text label at the end of the list widget.
*/
func (this *QListWidget) AddItem(label string) {
	var tmpArg0 = qtcore.NewQString5(label)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget7addItemERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:219
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void addItem(QListWidgetItem *)

/*
Inserts an item with the text label at the end of the list widget.
*/
func (this *QListWidget) AddItem1(item QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QListWidgetItem_PTR() != nil {
		convArg0 = item.QListWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget7addItemEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:220
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void addItems(const QStringList &)

/*
Inserts items with the text labels at the end of the list widget.

See also insertItems().
*/
func (this *QListWidget) AddItems(labels qtcore.QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if labels != nil && labels.QStringList_PTR() != nil {
		convArg0 = labels.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget8addItemsERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:221
// index:0
// Public Visibility=Default Availability=Available
// [8] QListWidgetItem * takeItem(int)

/*
Removes and returns the item from the given row in the list widget; otherwise returns 0.

Items removed from a list widget will not be managed by Qt, and will need to be deleted manually.

See also insertItem() and addItem().
*/
func (this *QListWidget) TakeItem(row int) *QListWidgetItem /*777 QListWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget8takeItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQListWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlistwidget.h:222
// index:0
// Public Visibility=Default Availability=Available
// [4] int count() const

/*

 */
func (this *QListWidget) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlistwidget.h:224
// index:0
// Public Visibility=Default Availability=Available
// [8] QListWidgetItem * currentItem() const

/*
Returns the current item.

See also setCurrentItem().
*/
func (this *QListWidget) CurrentItem() *QListWidgetItem /*777 QListWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget11currentItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQListWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlistwidget.h:225
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentItem(QListWidgetItem *)

/*
Sets the current item to item.

Unless the selection mode is NoSelection, the item is also selected.

See also currentItem().
*/
func (this *QListWidget) SetCurrentItem(item QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QListWidgetItem_PTR() != nil {
		convArg0 = item.QListWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget14setCurrentItemEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:226
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setCurrentItem(QListWidgetItem *, QItemSelectionModel::SelectionFlags)

/*
Sets the current item to item.

Unless the selection mode is NoSelection, the item is also selected.

See also currentItem().
*/
func (this *QListWidget) SetCurrentItem1(item QListWidgetItem_ITF /*777 QListWidgetItem **/, command int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QListWidgetItem_PTR() != nil {
		convArg0 = item.QListWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget14setCurrentItemEP15QListWidgetItem6QFlagsIN19QItemSelectionModel13SelectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, command)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:228
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentRow() const

/*

 */
func (this *QListWidget) CurrentRow() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget10currentRowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlistwidget.h:229
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentRow(int)

/*

 */
func (this *QListWidget) SetCurrentRow(row int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget13setCurrentRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:230
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setCurrentRow(int, QItemSelectionModel::SelectionFlags)

/*

 */
func (this *QListWidget) SetCurrentRow1(row int, command int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget13setCurrentRowEi6QFlagsIN19QItemSelectionModel13SelectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, command)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:232
// index:0
// Public Visibility=Default Availability=Available
// [8] QListWidgetItem * itemAt(const QPoint &) const

/*
Returns a pointer to the item at the coordinates p. The coordinates are relative to the list widget's viewport().
*/
func (this *QListWidget) ItemAt(p qtcore.QPoint_ITF) *QListWidgetItem /*777 QListWidgetItem **/ {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget6itemAtERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQListWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlistwidget.h:233
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QListWidgetItem * itemAt(int, int) const

/*
Returns a pointer to the item at the coordinates p. The coordinates are relative to the list widget's viewport().
*/
func (this *QListWidget) ItemAt1(x int, y int) *QListWidgetItem /*777 QListWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget6itemAtEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQListWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlistwidget.h:234
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect visualItemRect(const QListWidgetItem *) const

/*
Returns the rectangle on the viewport occupied by the item at item.
*/
func (this *QListWidget) VisualItemRect(item QListWidgetItem_ITF /*777 const QListWidgetItem **/) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QListWidgetItem_PTR() != nil {
		convArg0 = item.QListWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget14visualItemRectEPK15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:236
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sortItems(Qt::SortOrder)

/*
Sorts all the items in the list widget according to the specified order.
*/
func (this *QListWidget) SortItems(order int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget9sortItemsEN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:236
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sortItems(Qt::SortOrder)

/*
Sorts all the items in the list widget according to the specified order.
*/
func (this *QListWidget) SortItemsp() {
	// arg: 0, Qt::SortOrder=Elaborated, Qt::SortOrder=Enum, , Invalid
	order := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget9sortItemsEN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:237
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSortingEnabled(bool)

/*

 */
func (this *QListWidget) SetSortingEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget17setSortingEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:238
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSortingEnabled() const

/*

 */
func (this *QListWidget) IsSortingEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget16isSortingEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistwidget.h:240
// index:0
// Public Visibility=Default Availability=Available
// [-2] void editItem(QListWidgetItem *)

/*
Starts editing the item if it is editable.
*/
func (this *QListWidget) EditItem(item QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QListWidgetItem_PTR() != nil {
		convArg0 = item.QListWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget8editItemEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:241
// index:0
// Public Visibility=Default Availability=Available
// [-2] void openPersistentEditor(QListWidgetItem *)

/*
Opens an editor for the given item. The editor remains open after editing.

See also closePersistentEditor().
*/
func (this *QListWidget) OpenPersistentEditor(item QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QListWidgetItem_PTR() != nil {
		convArg0 = item.QListWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget20openPersistentEditorEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:242
// index:0
// Public Visibility=Default Availability=Available
// [-2] void closePersistentEditor(QListWidgetItem *)

/*
Closes the persistent editor for the given item.

See also openPersistentEditor().
*/
func (this *QListWidget) ClosePersistentEditor(item QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QListWidgetItem_PTR() != nil {
		convArg0 = item.QListWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget21closePersistentEditorEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:244
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * itemWidget(QListWidgetItem *) const

/*
Returns the widget displayed in the given item.

This function was introduced in  Qt 4.1.

See also setItemWidget() and removeItemWidget().
*/
func (this *QListWidget) ItemWidget(item QListWidgetItem_ITF /*777 QListWidgetItem **/) *QWidget /*777 QWidget **/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QListWidgetItem_PTR() != nil {
		convArg0 = item.QListWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget10itemWidgetEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlistwidget.h:245
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemWidget(QListWidgetItem *, QWidget *)

/*
Sets the widget to be displayed in the given item.

This function should only be used to display static content in the place of a list widget item. If you want to display custom dynamic content or implement a custom editor widget, use QListView and subclass QItemDelegate instead.

This function was introduced in  Qt 4.1.

See also itemWidget(), removeItemWidget(), and Delegate Classes.
*/
func (this *QListWidget) SetItemWidget(item QListWidgetItem_ITF /*777 QListWidgetItem **/, widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QListWidgetItem_PTR() != nil {
		convArg0 = item.QListWidgetItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg1 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget13setItemWidgetEP15QListWidgetItemP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:246
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void removeItemWidget(QListWidgetItem *)

/*
Removes the widget set on the given item.

To remove an item (row) from the list entirely, either delete the item or use takeItem().

This function was introduced in  Qt 4.3.

See also itemWidget() and setItemWidget().
*/
func (this *QListWidget) RemoveItemWidget(item QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QListWidgetItem_PTR() != nil {
		convArg0 = item.QListWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget16removeItemWidgetEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:248
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isItemSelected(const QListWidgetItem *) const

/*

 */
func (this *QListWidget) IsItemSelected(item QListWidgetItem_ITF /*777 const QListWidgetItem **/) bool {
	var convArg0 unsafe.Pointer
	if item != nil && item.QListWidgetItem_PTR() != nil {
		convArg0 = item.QListWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget14isItemSelectedEPK15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistwidget.h:249
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemSelected(const QListWidgetItem *, bool)

/*

 */
func (this *QListWidget) SetItemSelected(item QListWidgetItem_ITF /*777 const QListWidgetItem **/, select_ bool) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QListWidgetItem_PTR() != nil {
		convArg0 = item.QListWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget15setItemSelectedEPK15QListWidgetItemb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, select_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:253
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isItemHidden(const QListWidgetItem *) const

/*

 */
func (this *QListWidget) IsItemHidden(item QListWidgetItem_ITF /*777 const QListWidgetItem **/) bool {
	var convArg0 unsafe.Pointer
	if item != nil && item.QListWidgetItem_PTR() != nil {
		convArg0 = item.QListWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget12isItemHiddenEPK15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistwidget.h:254
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemHidden(const QListWidgetItem *, bool)

/*

 */
func (this *QListWidget) SetItemHidden(item QListWidgetItem_ITF /*777 const QListWidgetItem **/, hide bool) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QListWidgetItem_PTR() != nil {
		convArg0 = item.QListWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget13setItemHiddenEPK15QListWidgetItemb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hide)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:256
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)

/*
Reimplemented from QWidget::dropEvent().
*/
func (this *QListWidget) DropEvent(event qtgui.QDropEvent_ITF /*777 QDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QDropEvent_PTR() != nil {
		convArg0 = event.QDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget9dropEventEP10QDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:259
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scrollToItem(const QListWidgetItem *, QAbstractItemView::ScrollHint)

/*
Scrolls the view if necessary to ensure that the item is visible.

hint specifies where the item should be located after the operation.
*/
func (this *QListWidget) ScrollToItem(item QListWidgetItem_ITF /*777 const QListWidgetItem **/, hint int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QListWidgetItem_PTR() != nil {
		convArg0 = item.QListWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget12scrollToItemEPK15QListWidgetItemN17QAbstractItemView10ScrollHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:259
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scrollToItem(const QListWidgetItem *, QAbstractItemView::ScrollHint)

/*
Scrolls the view if necessary to ensure that the item is visible.

hint specifies where the item should be located after the operation.
*/
func (this *QListWidget) ScrollToItemp(item QListWidgetItem_ITF /*777 const QListWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QListWidgetItem_PTR() != nil {
		convArg0 = item.QListWidgetItem_PTR().GetCthis()
	}
	// arg: 1, QAbstractItemView::ScrollHint=Elaborated, QAbstractItemView::ScrollHint=Enum, , Invalid
	hint := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget12scrollToItemEPK15QListWidgetItemN17QAbstractItemView10ScrollHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:260
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Removes all items and selections in the view.

Warning: All items will be permanently deleted.
*/
func (this *QListWidget) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:263
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemPressed(QListWidgetItem *)

/*
This signal is emitted with the specified item when a mouse button is pressed on an item in the widget.

See also itemClicked() and itemDoubleClicked().
*/
func (this *QListWidget) ItemPressed(item QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QListWidgetItem_PTR() != nil {
		convArg0 = item.QListWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget11itemPressedEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:264
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemClicked(QListWidgetItem *)

/*
This signal is emitted with the specified item when a mouse button is clicked on an item in the widget.

See also itemPressed() and itemDoubleClicked().
*/
func (this *QListWidget) ItemClicked(item QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QListWidgetItem_PTR() != nil {
		convArg0 = item.QListWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget11itemClickedEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:265
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemDoubleClicked(QListWidgetItem *)

/*
This signal is emitted with the specified item when a mouse button is double clicked on an item in the widget.

See also itemClicked() and itemPressed().
*/
func (this *QListWidget) ItemDoubleClicked(item QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QListWidgetItem_PTR() != nil {
		convArg0 = item.QListWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget17itemDoubleClickedEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:266
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemActivated(QListWidgetItem *)

/*
This signal is emitted when the item is activated. The item is activated when the user clicks or double clicks on it, depending on the system configuration. It is also activated when the user presses the activation key (on Windows and X11 this is the Return key, on Mac OS X it is Command+O).
*/
func (this *QListWidget) ItemActivated(item QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QListWidgetItem_PTR() != nil {
		convArg0 = item.QListWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget13itemActivatedEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:267
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemEntered(QListWidgetItem *)

/*
This signal is emitted when the mouse cursor enters an item. The item is the item entered. This signal is only emitted when mouseTracking is turned on, or when a mouse button is pressed while moving into an item.

See also QWidget::setMouseTracking().
*/
func (this *QListWidget) ItemEntered(item QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QListWidgetItem_PTR() != nil {
		convArg0 = item.QListWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget11itemEnteredEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:268
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemChanged(QListWidgetItem *)

/*
This signal is emitted whenever the data of item has changed.
*/
func (this *QListWidget) ItemChanged(item QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QListWidgetItem_PTR() != nil {
		convArg0 = item.QListWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget11itemChangedEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:270
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentItemChanged(QListWidgetItem *, QListWidgetItem *)

/*
This signal is emitted whenever the current item changes.

previous is the item that previously had the focus; current is the new current item.
*/
func (this *QListWidget) CurrentItemChanged(current QListWidgetItem_ITF /*777 QListWidgetItem **/, previous QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if current != nil && current.QListWidgetItem_PTR() != nil {
		convArg0 = current.QListWidgetItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if previous != nil && previous.QListWidgetItem_PTR() != nil {
		convArg1 = previous.QListWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget18currentItemChangedEP15QListWidgetItemS1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:271
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentTextChanged(const QString &)

/*
This signal is emitted whenever the current item changes.

currentText is the text data in the current item. If there is no current item, the currentText is invalid.
*/
func (this *QListWidget) CurrentTextChanged(currentText string) {
	var tmpArg0 = qtcore.NewQString5(currentText)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget18currentTextChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:272
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentRowChanged(int)

/*
This signal is emitted whenever the current item changes.

currentRow is the row of the current item. If there is no current item, the currentRow is -1.

Note: Notifier signal for property currentRow.
*/
func (this *QListWidget) CurrentRowChanged(currentRow int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget17currentRowChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), currentRow)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:274
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemSelectionChanged()

/*
This signal is emitted whenever the selection changes.

See also selectedItems(), QListWidgetItem::isSelected(), and currentItemChanged().
*/
func (this *QListWidget) ItemSelectionChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget20itemSelectionChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:277
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QListWidget) Event(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistwidget.h:278
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QStringList mimeTypes() const

/*
Returns a list of MIME types that can be used to describe a list of listwidget items.

See also mimeData().
*/
func (this *QListWidget) MimeTypes() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget9mimeTypesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:285
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool dropMimeData(int, const QMimeData *, Qt::DropAction)

/*
Handles data supplied by an external drag and drop operation that ended with the given action in the given index. Returns true if data and action can be handled by the model; otherwise returns false.

See also supportedDropActions().
*/
func (this *QListWidget) DropMimeData(index int, data qtcore.QMimeData_ITF /*777 const QMimeData **/, action int) bool {
	var convArg1 unsafe.Pointer
	if data != nil && data.QMimeData_PTR() != nil {
		convArg1 = data.QMimeData_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget12dropMimeDataEiPK9QMimeDataN2Qt10DropActionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, action)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistwidget.h:286
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] Qt::DropActions supportedDropActions() const

/*
Returns the drop actions supported by this view.

See also Qt::DropActions.
*/
func (this *QListWidget) SupportedDropActions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget20supportedDropActionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:290
// index:0
// Protected Visibility=Default Availability=Available
// [24] QModelIndex indexFromItem(QListWidgetItem *) const

/*
Returns the QModelIndex associated with the given item.
*/
func (this *QListWidget) IndexFromItem(item QListWidgetItem_ITF /*777 QListWidgetItem **/) *qtcore.QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QListWidgetItem_PTR() != nil {
		convArg0 = item.QListWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget13indexFromItemEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:291
// index:0
// Protected Visibility=Default Availability=Available
// [8] QListWidgetItem * itemFromIndex(const QModelIndex &) const

/*
Returns a pointer to the QListWidgetItem associated with the given index.
*/
func (this *QListWidget) ItemFromIndex(index qtcore.QModelIndex_ITF) *QListWidgetItem /*777 QListWidgetItem **/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget13itemFromIndexERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQListWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
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
