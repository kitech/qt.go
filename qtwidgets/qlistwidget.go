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
// extern C begin: 43
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// bool event(class QEvent *)
func (this *QListWidget) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// QStringList mimeTypes()
func (this *QListWidget) InheritMimeTypes(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "mimeTypes", f)
}

// bool dropMimeData(int, const class QMimeData *, Qt::DropAction)
func (this *QListWidget) InheritDropMimeData(f func(index int, data *qtcore.QMimeData /*777 const QMimeData **/, action int) bool) {
	qtrt.SetAllInheritCallback(this, "dropMimeData", f)
}

// Qt::DropActions supportedDropActions()
func (this *QListWidget) InheritSupportedDropActions(f func() int) {
	qtrt.SetAllInheritCallback(this, "supportedDropActions", f)
}

// QModelIndex indexFromItem(class QListWidgetItem *)
func (this *QListWidget) InheritIndexFromItem(f func(item *QListWidgetItem /*777 QListWidgetItem **/) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "indexFromItem", f)
}

// QListWidgetItem * itemFromIndex(const class QModelIndex &)
func (this *QListWidget) InheritItemFromIndex(f func(index *qtcore.QModelIndex) unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "itemFromIndex", f)
}

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

// /usr/include/qt/QtWidgets/qlistwidget.h:199
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QListWidget) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlistwidget.h:207
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QListWidget(QWidget *)
func NewQListWidget(parent QWidget_ITF /*777 QWidget **/) *QListWidget {
	var convArg0 = parent.QWidget_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidgetC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQListWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qlistwidget.h:208
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QListWidget()
func DeleteQListWidget(this *QListWidget) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidgetD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:210
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setSelectionModel(QItemSelectionModel *)
func (this *QListWidget) SetSelectionModel(selectionModel qtcore.QItemSelectionModel_ITF /*777 QItemSelectionModel **/) {
	var convArg0 = selectionModel.QItemSelectionModel_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget17setSelectionModelEP19QItemSelectionModel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:212
// index:0
// Public Visibility=Default Availability=Available
// [8] QListWidgetItem * item(int)
func (this *QListWidget) Item(row int) *QListWidgetItem /*777 QListWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget4itemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQListWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlistwidget.h:213
// index:0
// Public Visibility=Default Availability=Available
// [4] int row(const QListWidgetItem *)
func (this *QListWidget) Row(item QListWidgetItem_ITF /*777 const QListWidgetItem **/) int {
	var convArg0 = item.QListWidgetItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget3rowEPK15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlistwidget.h:214
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertItem(int, QListWidgetItem *)
func (this *QListWidget) InsertItem(row int, item QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg1 = item.QListWidgetItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget10insertItemEiP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:215
// index:1
// Public Visibility=Default Availability=Available
// [-2] void insertItem(int, const QString &)
func (this *QListWidget) InsertItem_1(row int, label string) {
	var tmpArg1 = qtcore.NewQString_5(label)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget10insertItemEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:216
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertItems(int, const QStringList &)
func (this *QListWidget) InsertItems(row int, labels qtcore.QStringList_ITF) {
	var convArg1 = labels.QStringList_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget11insertItemsEiRK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:217
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void addItem(const QString &)
func (this *QListWidget) AddItem(label string) {
	var tmpArg0 = qtcore.NewQString_5(label)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget7addItemERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:218
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void addItem(QListWidgetItem *)
func (this *QListWidget) AddItem_1(item QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg0 = item.QListWidgetItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget7addItemEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:219
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void addItems(const QStringList &)
func (this *QListWidget) AddItems(labels qtcore.QStringList_ITF) {
	var convArg0 = labels.QStringList_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget8addItemsERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:220
// index:0
// Public Visibility=Default Availability=Available
// [8] QListWidgetItem * takeItem(int)
func (this *QListWidget) TakeItem(row int) *QListWidgetItem /*777 QListWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget8takeItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQListWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlistwidget.h:221
// index:0
// Public Visibility=Default Availability=Available
// [4] int count()
func (this *QListWidget) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlistwidget.h:223
// index:0
// Public Visibility=Default Availability=Available
// [8] QListWidgetItem * currentItem()
func (this *QListWidget) CurrentItem() *QListWidgetItem /*777 QListWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget11currentItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQListWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlistwidget.h:224
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentItem(QListWidgetItem *)
func (this *QListWidget) SetCurrentItem(item QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg0 = item.QListWidgetItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget14setCurrentItemEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:225
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setCurrentItem(QListWidgetItem *, QItemSelectionModel::SelectionFlags)
func (this *QListWidget) SetCurrentItem_1(item QListWidgetItem_ITF /*777 QListWidgetItem **/, command int) {
	var convArg0 = item.QListWidgetItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget14setCurrentItemEP15QListWidgetItem6QFlagsIN19QItemSelectionModel13SelectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, command)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:227
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentRow()
func (this *QListWidget) CurrentRow() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget10currentRowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlistwidget.h:228
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentRow(int)
func (this *QListWidget) SetCurrentRow(row int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget13setCurrentRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:229
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setCurrentRow(int, QItemSelectionModel::SelectionFlags)
func (this *QListWidget) SetCurrentRow_1(row int, command int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget13setCurrentRowEi6QFlagsIN19QItemSelectionModel13SelectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, command)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:231
// index:0
// Public Visibility=Default Availability=Available
// [8] QListWidgetItem * itemAt(const QPoint &)
func (this *QListWidget) ItemAt(p qtcore.QPoint_ITF) *QListWidgetItem /*777 QListWidgetItem **/ {
	var convArg0 = p.QPoint_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget6itemAtERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQListWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlistwidget.h:232
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QListWidgetItem * itemAt(int, int)
func (this *QListWidget) ItemAt_1(x int, y int) *QListWidgetItem /*777 QListWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget6itemAtEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQListWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlistwidget.h:233
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect visualItemRect(const QListWidgetItem *)
func (this *QListWidget) VisualItemRect(item QListWidgetItem_ITF /*777 const QListWidgetItem **/) *qtcore.QRect /*123*/ {
	var convArg0 = item.QListWidgetItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget14visualItemRectEPK15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:235
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sortItems(Qt::SortOrder)
func (this *QListWidget) SortItems(order int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget9sortItemsEN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:236
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSortingEnabled(_Bool)
func (this *QListWidget) SetSortingEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget17setSortingEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:237
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSortingEnabled()
func (this *QListWidget) IsSortingEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget16isSortingEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistwidget.h:239
// index:0
// Public Visibility=Default Availability=Available
// [-2] void editItem(QListWidgetItem *)
func (this *QListWidget) EditItem(item QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg0 = item.QListWidgetItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget8editItemEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:240
// index:0
// Public Visibility=Default Availability=Available
// [-2] void openPersistentEditor(QListWidgetItem *)
func (this *QListWidget) OpenPersistentEditor(item QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg0 = item.QListWidgetItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget20openPersistentEditorEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:241
// index:0
// Public Visibility=Default Availability=Available
// [-2] void closePersistentEditor(QListWidgetItem *)
func (this *QListWidget) ClosePersistentEditor(item QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg0 = item.QListWidgetItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget21closePersistentEditorEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:243
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPersistentEditorOpen(QListWidgetItem *)
func (this *QListWidget) IsPersistentEditorOpen(item QListWidgetItem_ITF /*777 QListWidgetItem **/) bool {
	var convArg0 = item.QListWidgetItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget22isPersistentEditorOpenEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistwidget.h:245
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * itemWidget(QListWidgetItem *)
func (this *QListWidget) ItemWidget(item QListWidgetItem_ITF /*777 QListWidgetItem **/) *QWidget /*777 QWidget **/ {
	var convArg0 = item.QListWidgetItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget10itemWidgetEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlistwidget.h:246
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemWidget(QListWidgetItem *, QWidget *)
func (this *QListWidget) SetItemWidget(item QListWidgetItem_ITF /*777 QListWidgetItem **/, widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 = item.QListWidgetItem_PTR().GetCthis()
	var convArg1 = widget.QWidget_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget13setItemWidgetEP15QListWidgetItemP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:247
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void removeItemWidget(QListWidgetItem *)
func (this *QListWidget) RemoveItemWidget(item QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg0 = item.QListWidgetItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget16removeItemWidgetEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:249
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isItemSelected(const QListWidgetItem *)
func (this *QListWidget) IsItemSelected(item QListWidgetItem_ITF /*777 const QListWidgetItem **/) bool {
	var convArg0 = item.QListWidgetItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget14isItemSelectedEPK15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistwidget.h:250
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemSelected(const QListWidgetItem *, _Bool)
func (this *QListWidget) SetItemSelected(item QListWidgetItem_ITF /*777 const QListWidgetItem **/, select_ bool) {
	var convArg0 = item.QListWidgetItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget15setItemSelectedEPK15QListWidgetItemb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, select_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:254
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isItemHidden(const QListWidgetItem *)
func (this *QListWidget) IsItemHidden(item QListWidgetItem_ITF /*777 const QListWidgetItem **/) bool {
	var convArg0 = item.QListWidgetItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget12isItemHiddenEPK15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistwidget.h:255
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemHidden(const QListWidgetItem *, _Bool)
func (this *QListWidget) SetItemHidden(item QListWidgetItem_ITF /*777 const QListWidgetItem **/, hide bool) {
	var convArg0 = item.QListWidgetItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget13setItemHiddenEPK15QListWidgetItemb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hide)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:260
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)
func (this *QListWidget) DropEvent(event qtgui.QDropEvent_ITF /*777 QDropEvent **/) {
	var convArg0 = event.QDropEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget9dropEventEP10QDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:263
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scrollToItem(const QListWidgetItem *, QAbstractItemView::ScrollHint)
func (this *QListWidget) ScrollToItem(item QListWidgetItem_ITF /*777 const QListWidgetItem **/, hint int) {
	var convArg0 = item.QListWidgetItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget12scrollToItemEPK15QListWidgetItemN17QAbstractItemView10ScrollHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:264
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QListWidget) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:267
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemPressed(QListWidgetItem *)
func (this *QListWidget) ItemPressed(item QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg0 = item.QListWidgetItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget11itemPressedEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:268
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemClicked(QListWidgetItem *)
func (this *QListWidget) ItemClicked(item QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg0 = item.QListWidgetItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget11itemClickedEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:269
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemDoubleClicked(QListWidgetItem *)
func (this *QListWidget) ItemDoubleClicked(item QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg0 = item.QListWidgetItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget17itemDoubleClickedEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:270
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemActivated(QListWidgetItem *)
func (this *QListWidget) ItemActivated(item QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg0 = item.QListWidgetItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget13itemActivatedEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:271
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemEntered(QListWidgetItem *)
func (this *QListWidget) ItemEntered(item QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg0 = item.QListWidgetItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget11itemEnteredEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:272
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemChanged(QListWidgetItem *)
func (this *QListWidget) ItemChanged(item QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg0 = item.QListWidgetItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget11itemChangedEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:274
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentItemChanged(QListWidgetItem *, QListWidgetItem *)
func (this *QListWidget) CurrentItemChanged(current QListWidgetItem_ITF /*777 QListWidgetItem **/, previous QListWidgetItem_ITF /*777 QListWidgetItem **/) {
	var convArg0 = current.QListWidgetItem_PTR().GetCthis()
	var convArg1 = previous.QListWidgetItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget18currentItemChangedEP15QListWidgetItemS1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:275
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentTextChanged(const QString &)
func (this *QListWidget) CurrentTextChanged(currentText string) {
	var tmpArg0 = qtcore.NewQString_5(currentText)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget18currentTextChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:276
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentRowChanged(int)
func (this *QListWidget) CurrentRowChanged(currentRow int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget17currentRowChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), currentRow)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:278
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemSelectionChanged()
func (this *QListWidget) ItemSelectionChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget20itemSelectionChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:281
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QListWidget) Event(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 = e.QEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistwidget.h:282
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QStringList mimeTypes()
func (this *QListWidget) MimeTypes() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget9mimeTypesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:289
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool dropMimeData(int, const QMimeData *, Qt::DropAction)
func (this *QListWidget) DropMimeData(index int, data qtcore.QMimeData_ITF /*777 const QMimeData **/, action int) bool {
	var convArg1 = data.QMimeData_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget12dropMimeDataEiPK9QMimeDataN2Qt10DropActionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, action)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistwidget.h:290
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] Qt::DropActions supportedDropActions()
func (this *QListWidget) SupportedDropActions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget20supportedDropActionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:300
// index:0
// Protected Visibility=Default Availability=Available
// [24] QModelIndex indexFromItem(QListWidgetItem *)
func (this *QListWidget) IndexFromItem(item QListWidgetItem_ITF /*777 QListWidgetItem **/) *qtcore.QModelIndex /*123*/ {
	var convArg0 = item.QListWidgetItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget13indexFromItemEP15QListWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:301
// index:0
// Protected Visibility=Default Availability=Available
// [8] QListWidgetItem * itemFromIndex(const QModelIndex &)
func (this *QListWidget) ItemFromIndex(index qtcore.QModelIndex_ITF) *QListWidgetItem /*777 QListWidgetItem **/ {
	var convArg0 = index.QModelIndex_PTR().GetCthis()
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
