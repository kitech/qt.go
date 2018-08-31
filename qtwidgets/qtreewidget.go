package qtwidgets

// /usr/include/qt/QtWidgets/qtreewidget.h
// #include <qtreewidget.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 69
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
func (this *QTreeWidget) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// QStringList mimeTypes()
func (this *QTreeWidget) InheritMimeTypes(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "mimeTypes", f)
}

// bool dropMimeData(QTreeWidgetItem *, int, const QMimeData *, Qt::DropAction)
func (this *QTreeWidget) InheritDropMimeData(f func(parent *QTreeWidgetItem /*777 QTreeWidgetItem **/, index int, data *qtcore.QMimeData /*777 const QMimeData **/, action int) bool) {
	qtrt.SetAllInheritCallback(this, "dropMimeData", f)
}

// Qt::DropActions supportedDropActions()
func (this *QTreeWidget) InheritSupportedDropActions(f func() int) {
	qtrt.SetAllInheritCallback(this, "supportedDropActions", f)
}

// QModelIndex indexFromItem(const QTreeWidgetItem *, int)
func (this *QTreeWidget) InheritIndexFromItem(f func(item *QTreeWidgetItem /*777 const QTreeWidgetItem **/, column int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "indexFromItem", f)
}

// QTreeWidgetItem * itemFromIndex(const QModelIndex &)
func (this *QTreeWidget) InheritItemFromIndex(f func(index *qtcore.QModelIndex) unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "itemFromIndex", f)
}

// void dropEvent(QDropEvent *)
func (this *QTreeWidget) InheritDropEvent(f func(event *qtgui.QDropEvent /*777 QDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dropEvent", f)
}

/*

 */
type QTreeWidget struct {
	*QTreeView
}
type QTreeWidget_ITF interface {
	QTreeView_ITF
	QTreeWidget_PTR() *QTreeWidget
}

func (ptr *QTreeWidget) QTreeWidget_PTR() *QTreeWidget { return ptr }

func (this *QTreeWidget) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QTreeView.GetCthis()
	}
}
func (this *QTreeWidget) SetCthis(cthis unsafe.Pointer) {
	this.QTreeView = NewQTreeViewFromPointer(cthis)
}
func NewQTreeWidgetFromPointer(cthis unsafe.Pointer) *QTreeWidget {
	bcthis0 := NewQTreeViewFromPointer(cthis)
	return &QTreeWidget{bcthis0}
}
func (*QTreeWidget) NewFromPointer(cthis unsafe.Pointer) *QTreeWidget {
	return NewQTreeWidgetFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:257
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QTreeWidget) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:264
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTreeWidget(QWidget *)

/*
Constructs a tree widget with the given parent.
*/
func (*QTreeWidget) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QTreeWidget {
	return NewQTreeWidget(parent)
}
func NewQTreeWidget(parent QWidget_ITF /*777 QWidget **/) *QTreeWidget {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidgetC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTreeWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTreeWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:264
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTreeWidget(QWidget *)

/*
Constructs a tree widget with the given parent.
*/
func (*QTreeWidget) NewForInherit__() *QTreeWidget {
	return NewQTreeWidget__()
}
func NewQTreeWidget__() *QTreeWidget {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidgetC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTreeWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTreeWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:265
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTreeWidget()

/*

 */
func DeleteQTreeWidget(this *QTreeWidget) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidgetD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:267
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnCount() const

/*

 */
func (this *QTreeWidget) ColumnCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget11columnCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreewidget.h:268
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnCount(int)

/*

 */
func (this *QTreeWidget) SetColumnCount(columns int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget14setColumnCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), columns)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:270
// index:0
// Public Visibility=Default Availability=Available
// [8] QTreeWidgetItem * invisibleRootItem() const

/*
Returns the tree widget's invisible root item.

The invisible root item provides access to the tree widget's top-level items through the QTreeWidgetItem API, making it possible to write functions that can treat top-level items and their children in a uniform way; for example, recursive functions.

This function was introduced in  Qt 4.2.
*/
func (this *QTreeWidget) InvisibleRootItem() *QTreeWidgetItem /*777 QTreeWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget17invisibleRootItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:271
// index:0
// Public Visibility=Default Availability=Available
// [8] QTreeWidgetItem * topLevelItem(int) const

/*
Returns the top level item at the given index, or 0 if the item does not exist.

See also topLevelItemCount() and insertTopLevelItem().
*/
func (this *QTreeWidget) TopLevelItem(index int) *QTreeWidgetItem /*777 QTreeWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget12topLevelItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:272
// index:0
// Public Visibility=Default Availability=Available
// [4] int topLevelItemCount() const

/*

 */
func (this *QTreeWidget) TopLevelItemCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget17topLevelItemCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreewidget.h:273
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertTopLevelItem(int, QTreeWidgetItem *)

/*
Inserts the item at index in the top level in the view.

If the item has already been inserted somewhere else it won't be inserted.

See also addTopLevelItem() and columnCount().
*/
func (this *QTreeWidget) InsertTopLevelItem(index int, item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) {
	var convArg1 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg1 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget18insertTopLevelItemEiP15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:274
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addTopLevelItem(QTreeWidgetItem *)

/*
Appends the item as a top-level item in the widget.

This function was introduced in  Qt 4.1.

See also insertTopLevelItem().
*/
func (this *QTreeWidget) AddTopLevelItem(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget15addTopLevelItemEP15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:275
// index:0
// Public Visibility=Default Availability=Available
// [8] QTreeWidgetItem * takeTopLevelItem(int)

/*
Removes the top-level item at the given index in the tree and returns it, otherwise returns 0;

See also insertTopLevelItem(), topLevelItem(), and topLevelItemCount().
*/
func (this *QTreeWidget) TakeTopLevelItem(index int) *QTreeWidgetItem /*777 QTreeWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget16takeTopLevelItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:276
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOfTopLevelItem(QTreeWidgetItem *) const

/*
Returns the index of the given top-level item, or -1 if the item cannot be found.

See also sortItems() and topLevelItemCount().
*/
func (this *QTreeWidget) IndexOfTopLevelItem(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) int {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget19indexOfTopLevelItemEP15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreewidget.h:281
// index:0
// Public Visibility=Default Availability=Available
// [8] QTreeWidgetItem * headerItem() const

/*
Returns the item used for the tree widget's header.

See also setHeaderItem().
*/
func (this *QTreeWidget) HeaderItem() *QTreeWidgetItem /*777 QTreeWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget10headerItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:282
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHeaderItem(QTreeWidgetItem *)

/*
Sets the header item for the tree widget. The label for each column in the header is supplied by the corresponding label in the item.

The tree widget takes ownership of the item.

See also headerItem() and setHeaderLabels().
*/
func (this *QTreeWidget) SetHeaderItem(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget13setHeaderItemEP15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:283
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHeaderLabels(const QStringList &)

/*
Adds a column in the header for each item in the labels list, and sets the label for each column.

Note that setHeaderLabels() won't remove existing columns.

See also setHeaderItem() and setHeaderLabel().
*/
func (this *QTreeWidget) SetHeaderLabels(labels qtcore.QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if labels != nil && labels.QStringList_PTR() != nil {
		convArg0 = labels.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget15setHeaderLabelsERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:284
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setHeaderLabel(const QString &)

/*
Same as setHeaderLabels(QStringList(label)).

This function was introduced in  Qt 4.2.
*/
func (this *QTreeWidget) SetHeaderLabel(label string) {
	var tmpArg0 = qtcore.NewQString_5(label)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget14setHeaderLabelERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:286
// index:0
// Public Visibility=Default Availability=Available
// [8] QTreeWidgetItem * currentItem() const

/*
Returns the current item in the tree widget.

See also setCurrentItem() and currentItemChanged().
*/
func (this *QTreeWidget) CurrentItem() *QTreeWidgetItem /*777 QTreeWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget11currentItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:287
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentColumn() const

/*
Returns the current column in the tree widget.

This function was introduced in  Qt 4.1.

See also setCurrentItem() and columnCount().
*/
func (this *QTreeWidget) CurrentColumn() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget13currentColumnEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreewidget.h:288
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentItem(QTreeWidgetItem *)

/*
Sets the current item in the tree widget.

Unless the selection mode is NoSelection, the item is also selected.

See also currentItem() and currentItemChanged().
*/
func (this *QTreeWidget) SetCurrentItem(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:289
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setCurrentItem(QTreeWidgetItem *, int)

/*
Sets the current item in the tree widget.

Unless the selection mode is NoSelection, the item is also selected.

See also currentItem() and currentItemChanged().
*/
func (this *QTreeWidget) SetCurrentItem_1(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:290
// index:2
// Public Visibility=Default Availability=Available
// [-2] void setCurrentItem(QTreeWidgetItem *, int, QItemSelectionModel::SelectionFlags)

/*
Sets the current item in the tree widget.

Unless the selection mode is NoSelection, the item is also selected.

See also currentItem() and currentItemChanged().
*/
func (this *QTreeWidget) SetCurrentItem_2(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int, command int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItemi6QFlagsIN19QItemSelectionModel13SelectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column, command)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:292
// index:0
// Public Visibility=Default Availability=Available
// [8] QTreeWidgetItem * itemAt(const QPoint &) const

/*
Returns a pointer to the item at the coordinates p. The coordinates are relative to the tree widget's viewport().

See also visualItemRect().
*/
func (this *QTreeWidget) ItemAt(p qtcore.QPoint_ITF) *QTreeWidgetItem /*777 QTreeWidgetItem **/ {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget6itemAtERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:293
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QTreeWidgetItem * itemAt(int, int) const

/*
Returns a pointer to the item at the coordinates p. The coordinates are relative to the tree widget's viewport().

See also visualItemRect().
*/
func (this *QTreeWidget) ItemAt_1(x int, y int) *QTreeWidgetItem /*777 QTreeWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget6itemAtEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:294
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect visualItemRect(const QTreeWidgetItem *) const

/*
Returns the rectangle on the viewport occupied by the item at item.

See also itemAt().
*/
func (this *QTreeWidget) VisualItemRect(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget14visualItemRectEPK15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:296
// index:0
// Public Visibility=Default Availability=Available
// [4] int sortColumn() const

/*
Returns the column used to sort the contents of the widget.

This function was introduced in  Qt 4.1.

See also sortItems().
*/
func (this *QTreeWidget) SortColumn() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget10sortColumnEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreewidget.h:297
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sortItems(int, Qt::SortOrder)

/*
Sorts the items in the widget in the specified order by the values in the given column.

See also sortColumn().
*/
func (this *QTreeWidget) SortItems(column int, order int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget9sortItemsEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:299
// index:0
// Public Visibility=Default Availability=Available
// [-2] void editItem(QTreeWidgetItem *, int)

/*
Starts editing the item in the given column if it is editable.
*/
func (this *QTreeWidget) EditItem(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget8editItemEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:299
// index:0
// Public Visibility=Default Availability=Available
// [-2] void editItem(QTreeWidgetItem *, int)

/*
Starts editing the item in the given column if it is editable.
*/
func (this *QTreeWidget) EditItem__(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	column := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget8editItemEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:300
// index:0
// Public Visibility=Default Availability=Available
// [-2] void openPersistentEditor(QTreeWidgetItem *, int)

/*
Opens a persistent editor for the item in the given column.

See also closePersistentEditor() and isPersistentEditorOpen().
*/
func (this *QTreeWidget) OpenPersistentEditor(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget20openPersistentEditorEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:300
// index:0
// Public Visibility=Default Availability=Available
// [-2] void openPersistentEditor(QTreeWidgetItem *, int)

/*
Opens a persistent editor for the item in the given column.

See also closePersistentEditor() and isPersistentEditorOpen().
*/
func (this *QTreeWidget) OpenPersistentEditor__(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	column := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget20openPersistentEditorEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:301
// index:0
// Public Visibility=Default Availability=Available
// [-2] void closePersistentEditor(QTreeWidgetItem *, int)

/*
Closes the persistent editor for the item in the given column.

This function has no effect if no persistent editor is open for this combination of item and column.

See also openPersistentEditor() and isPersistentEditorOpen().
*/
func (this *QTreeWidget) ClosePersistentEditor(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget21closePersistentEditorEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:301
// index:0
// Public Visibility=Default Availability=Available
// [-2] void closePersistentEditor(QTreeWidgetItem *, int)

/*
Closes the persistent editor for the item in the given column.

This function has no effect if no persistent editor is open for this combination of item and column.

See also openPersistentEditor() and isPersistentEditorOpen().
*/
func (this *QTreeWidget) ClosePersistentEditor__(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	column := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget21closePersistentEditorEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:303
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPersistentEditorOpen(QTreeWidgetItem *, int) const

/*
Returns whether a persistent editor is open for item item in column column.

This function was introduced in  Qt 5.10.

See also openPersistentEditor() and closePersistentEditor().
*/
func (this *QTreeWidget) IsPersistentEditorOpen(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) bool {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget22isPersistentEditorOpenEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:303
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPersistentEditorOpen(QTreeWidgetItem *, int) const

/*
Returns whether a persistent editor is open for item item in column column.

This function was introduced in  Qt 5.10.

See also openPersistentEditor() and closePersistentEditor().
*/
func (this *QTreeWidget) IsPersistentEditorOpen__(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) bool {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	column := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget22isPersistentEditorOpenEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:305
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * itemWidget(QTreeWidgetItem *, int) const

/*
Returns the widget displayed in the cell specified by item and the given column.

This function was introduced in  Qt 4.1.

See also setItemWidget().
*/
func (this *QTreeWidget) ItemWidget(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) *QWidget /*777 QWidget **/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget10itemWidgetEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:306
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemWidget(QTreeWidgetItem *, int, QWidget *)

/*
Sets the given widget to be displayed in the cell specified by the given item and column.

The given widget's autoFillBackground property must be set to true, otherwise the widget's background will be transparent, showing both the model data and the tree widget item.

This function should only be used to display static content in the place of a tree widget item. If you want to display custom dynamic content or implement a custom editor widget, use QTreeView and subclass QItemDelegate instead.

This function cannot be called before the item hierarchy has been set up, i.e., the QTreeWidgetItem that will hold widget must have been added to the view before widget is set.

Note: The tree takes ownership of the widget.

This function was introduced in  Qt 4.1.

See also itemWidget() and Delegate Classes.
*/
func (this *QTreeWidget) SetItemWidget(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int, widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg2 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget13setItemWidgetEP15QTreeWidgetItemiP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:307
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void removeItemWidget(QTreeWidgetItem *, int)

/*
Removes the widget set in the given item in the given column.

This function was introduced in  Qt 4.3.
*/
func (this *QTreeWidget) RemoveItemWidget(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget16removeItemWidgetEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:309
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isItemSelected(const QTreeWidgetItem *) const

/*

 */
func (this *QTreeWidget) IsItemSelected(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/) bool {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget14isItemSelectedEPK15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:310
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemSelected(const QTreeWidgetItem *, bool)

/*

 */
func (this *QTreeWidget) SetItemSelected(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/, select_ bool) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget15setItemSelectedEPK15QTreeWidgetItemb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, select_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:315
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isItemHidden(const QTreeWidgetItem *) const

/*

 */
func (this *QTreeWidget) IsItemHidden(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/) bool {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget12isItemHiddenEPK15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:316
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemHidden(const QTreeWidgetItem *, bool)

/*

 */
func (this *QTreeWidget) SetItemHidden(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/, hide bool) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget13setItemHiddenEPK15QTreeWidgetItemb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hide)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:318
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isItemExpanded(const QTreeWidgetItem *) const

/*

 */
func (this *QTreeWidget) IsItemExpanded(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/) bool {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget14isItemExpandedEPK15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:319
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemExpanded(const QTreeWidgetItem *, bool)

/*

 */
func (this *QTreeWidget) SetItemExpanded(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/, expand bool) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget15setItemExpandedEPK15QTreeWidgetItemb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, expand)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:321
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFirstItemColumnSpanned(const QTreeWidgetItem *) const

/*
Returns true if the given item is set to show only one section over all columns; otherwise returns false.

This function was introduced in  Qt 4.3.

See also setFirstItemColumnSpanned().
*/
func (this *QTreeWidget) IsFirstItemColumnSpanned(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/) bool {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget24isFirstItemColumnSpannedEPK15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:322
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFirstItemColumnSpanned(const QTreeWidgetItem *, bool)

/*
Sets the given item to only show one section for all columns if span is true; otherwise the item will show one section per column.

This function was introduced in  Qt 4.3.

See also isFirstItemColumnSpanned().
*/
func (this *QTreeWidget) SetFirstItemColumnSpanned(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/, span bool) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget25setFirstItemColumnSpannedEPK15QTreeWidgetItemb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, span)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:324
// index:0
// Public Visibility=Default Availability=Available
// [8] QTreeWidgetItem * itemAbove(const QTreeWidgetItem *) const

/*
Returns the item above the given item.

This function was introduced in  Qt 4.3.
*/
func (this *QTreeWidget) ItemAbove(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/) *QTreeWidgetItem /*777 QTreeWidgetItem **/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget9itemAboveEPK15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:325
// index:0
// Public Visibility=Default Availability=Available
// [8] QTreeWidgetItem * itemBelow(const QTreeWidgetItem *) const

/*
Returns the item visually below the given item.

This function was introduced in  Qt 4.3.
*/
func (this *QTreeWidget) ItemBelow(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/) *QTreeWidgetItem /*777 QTreeWidgetItem **/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget9itemBelowEPK15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:327
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setSelectionModel(QItemSelectionModel *)

/*
Reimplemented from QAbstractItemView::setSelectionModel().
*/
func (this *QTreeWidget) SetSelectionModel(selectionModel qtcore.QItemSelectionModel_ITF /*777 QItemSelectionModel **/) {
	var convArg0 unsafe.Pointer
	if selectionModel != nil && selectionModel.QItemSelectionModel_PTR() != nil {
		convArg0 = selectionModel.QItemSelectionModel_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget17setSelectionModelEP19QItemSelectionModel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:330
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scrollToItem(const QTreeWidgetItem *, QAbstractItemView::ScrollHint)

/*
Ensures that the item is visible, scrolling the view if necessary using the specified hint.

See also currentItem(), itemAt(), and topLevelItem().
*/
func (this *QTreeWidget) ScrollToItem(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/, hint int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget12scrollToItemEPK15QTreeWidgetItemN17QAbstractItemView10ScrollHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:330
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scrollToItem(const QTreeWidgetItem *, QAbstractItemView::ScrollHint)

/*
Ensures that the item is visible, scrolling the view if necessary using the specified hint.

See also currentItem(), itemAt(), and topLevelItem().
*/
func (this *QTreeWidget) ScrollToItem__(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	// arg: 1, QAbstractItemView::ScrollHint=Elaborated, QAbstractItemView::ScrollHint=Enum, , Invalid
	hint := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget12scrollToItemEPK15QTreeWidgetItemN17QAbstractItemView10ScrollHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:332
// index:0
// Public Visibility=Default Availability=Available
// [-2] void expandItem(const QTreeWidgetItem *)

/*
Expands the item. This causes the tree containing the item's children to be expanded.

See also collapseItem(), currentItem(), itemAt(), topLevelItem(), and itemExpanded().
*/
func (this *QTreeWidget) ExpandItem(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget10expandItemEPK15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:333
// index:0
// Public Visibility=Default Availability=Available
// [-2] void collapseItem(const QTreeWidgetItem *)

/*
Closes the item. This causes the tree containing the item's children to be collapsed.

See also expandItem(), currentItem(), itemAt(), and topLevelItem().
*/
func (this *QTreeWidget) CollapseItem(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget12collapseItemEPK15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:334
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Clears the tree widget by removing all of its items and selections.

Note: Since each item is removed from the tree widget before being deleted, the return value of QTreeWidgetItem::treeWidget() will be invalid when called from an item's destructor.

See also takeTopLevelItem(), topLevelItemCount(), and columnCount().
*/
func (this *QTreeWidget) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:337
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemPressed(QTreeWidgetItem *, int)

/*
This signal is emitted when the user presses a mouse button inside the widget.

The specified item is the item that was clicked, or 0 if no item was clicked. The column is the item's column that was clicked, or -1 if no item was clicked.
*/
func (this *QTreeWidget) ItemPressed(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget11itemPressedEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:338
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemClicked(QTreeWidgetItem *, int)

/*
This signal is emitted when the user clicks inside the widget.

The specified item is the item that was clicked. The column is the item's column that was clicked. If no item was clicked, no signal will be emitted.
*/
func (this *QTreeWidget) ItemClicked(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget11itemClickedEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:339
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemDoubleClicked(QTreeWidgetItem *, int)

/*
This signal is emitted when the user double clicks inside the widget.

The specified item is the item that was clicked, or 0 if no item was clicked. The column is the item's column that was clicked. If no item was double clicked, no signal will be emitted.
*/
func (this *QTreeWidget) ItemDoubleClicked(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget17itemDoubleClickedEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:340
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemActivated(QTreeWidgetItem *, int)

/*
This signal is emitted when the user activates an item by single- or double-clicking (depending on the platform, i.e. on the QStyle::SH_ItemView_ActivateItemOnSingleClick style hint) or pressing a special key (e.g., Enter).

The specified item is the item that was clicked, or 0 if no item was clicked. The column is the item's column that was clicked, or -1 if no item was clicked.
*/
func (this *QTreeWidget) ItemActivated(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget13itemActivatedEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:341
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemEntered(QTreeWidgetItem *, int)

/*
This signal is emitted when the mouse cursor enters an item over the specified column. QTreeWidget mouse tracking needs to be enabled for this feature to work.
*/
func (this *QTreeWidget) ItemEntered(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget11itemEnteredEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:342
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemChanged(QTreeWidgetItem *, int)

/*
This signal is emitted when the contents of the column in the specified item changes.
*/
func (this *QTreeWidget) ItemChanged(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget11itemChangedEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:343
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemExpanded(QTreeWidgetItem *)

/*
This signal is emitted when the specified item is expanded so that all of its children are displayed.

Note: This signal will not be emitted if an item changes its state when expandAll() is invoked.

See also setItemExpanded(), QTreeWidgetItem::isExpanded(), itemCollapsed(), and expandItem().
*/
func (this *QTreeWidget) ItemExpanded(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget12itemExpandedEP15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:344
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemCollapsed(QTreeWidgetItem *)

/*
This signal is emitted when the specified item is collapsed so that none of its children are displayed.

Note: This signal will not be emitted if an item changes its state when collapseAll() is invoked.

See also QTreeWidgetItem::isExpanded(), itemExpanded(), and collapseItem().
*/
func (this *QTreeWidget) ItemCollapsed(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget13itemCollapsedEP15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:345
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentItemChanged(QTreeWidgetItem *, QTreeWidgetItem *)

/*
This signal is emitted when the current item changes. The current item is specified by current, and this replaces the previous current item.

See also setCurrentItem().
*/
func (this *QTreeWidget) CurrentItemChanged(current QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, previous QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if current != nil && current.QTreeWidgetItem_PTR() != nil {
		convArg0 = current.QTreeWidgetItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if previous != nil && previous.QTreeWidgetItem_PTR() != nil {
		convArg1 = previous.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget18currentItemChangedEP15QTreeWidgetItemS1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:346
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemSelectionChanged()

/*
This signal is emitted when the selection changes in the tree widget. The current selection can be found with selectedItems().
*/
func (this *QTreeWidget) ItemSelectionChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget20itemSelectionChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:349
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QTreeWidget) Event(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:350
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QStringList mimeTypes() const

/*
Returns a list of MIME types that can be used to describe a list of treewidget items.

See also mimeData().
*/
func (this *QTreeWidget) MimeTypes() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget9mimeTypesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:356
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool dropMimeData(QTreeWidgetItem *, int, const QMimeData *, Qt::DropAction)

/*
Handles the data supplied by a drag and drop operation that ended with the given action in the index in the given parent item.

The default implementation returns true if the drop was successfully handled by decoding the mime data and inserting it into the model; otherwise it returns false.

See also supportedDropActions().
*/
func (this *QTreeWidget) DropMimeData(parent QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, index int, data qtcore.QMimeData_ITF /*777 const QMimeData **/, action int) bool {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QTreeWidgetItem_PTR() != nil {
		convArg0 = parent.QTreeWidgetItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if data != nil && data.QMimeData_PTR() != nil {
		convArg2 = data.QMimeData_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget12dropMimeDataEP15QTreeWidgetItemiPK9QMimeDataN2Qt10DropActionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, index, convArg2, action)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:358
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] Qt::DropActions supportedDropActions() const

/*
Returns the drop actions supported by this view.

See also Qt::DropActions.
*/
func (this *QTreeWidget) SupportedDropActions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget20supportedDropActionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:367
// index:0
// Protected Visibility=Default Availability=Available
// [24] QModelIndex indexFromItem(const QTreeWidgetItem *, int) const

/*
Returns the QModelIndex associated with the given item in the given column.

Note: In Qt versions prior to 5.7, this function took a non-const item.

See also itemFromIndex() and topLevelItem().
*/
func (this *QTreeWidget) IndexFromItem(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/, column int) *qtcore.QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget13indexFromItemEPK15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:367
// index:0
// Protected Visibility=Default Availability=Available
// [24] QModelIndex indexFromItem(const QTreeWidgetItem *, int) const

/*
Returns the QModelIndex associated with the given item in the given column.

Note: In Qt versions prior to 5.7, this function took a non-const item.

See also itemFromIndex() and topLevelItem().
*/
func (this *QTreeWidget) IndexFromItem__(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/) *qtcore.QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	column := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget13indexFromItemEPK15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:368
// index:1
// Protected Visibility=Default Availability=Available
// [24] QModelIndex indexFromItem(QTreeWidgetItem *, int) const

/*
Returns the QModelIndex associated with the given item in the given column.

Note: In Qt versions prior to 5.7, this function took a non-const item.

See also itemFromIndex() and topLevelItem().
*/
func (this *QTreeWidget) IndexFromItem_1(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) *qtcore.QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget13indexFromItemEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:368
// index:1
// Protected Visibility=Default Availability=Available
// [24] QModelIndex indexFromItem(QTreeWidgetItem *, int) const

/*
Returns the QModelIndex associated with the given item in the given column.

Note: In Qt versions prior to 5.7, this function took a non-const item.

See also itemFromIndex() and topLevelItem().
*/
func (this *QTreeWidget) IndexFromItem_1_(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) *qtcore.QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	column := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget13indexFromItemEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:369
// index:0
// Protected Visibility=Default Availability=Available
// [8] QTreeWidgetItem * itemFromIndex(const QModelIndex &) const

/*
Returns a pointer to the QTreeWidgetItem associated with the given index.

See also indexFromItem().
*/
func (this *QTreeWidget) ItemFromIndex(index qtcore.QModelIndex_ITF) *QTreeWidgetItem /*777 QTreeWidgetItem **/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget13itemFromIndexERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:373
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)

/*
Reimplemented from QWidget::dropEvent().
*/
func (this *QTreeWidget) DropEvent(event qtgui.QDropEvent_ITF /*777 QDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QDropEvent_PTR() != nil {
		convArg0 = event.QDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget9dropEventEP10QDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
