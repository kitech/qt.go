//  header block begin
// /usr/include/qt/QtWidgets/qlistwidget.h
// #include <qlistwidget.h>
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
type QListWidget struct {
	*QListView
}

func (this *QListWidget) GetCthis() unsafe.Pointer {
	return this.QListView.GetCthis()
}
func NewQListWidgetFromPointer(cthis unsafe.Pointer) *QListWidget {
	bcthis0 := NewQListViewFromPointer(cthis)
	return &QListWidget{bcthis0}
}

// /usr/include/qt/QtWidgets/qlistwidget.h:199
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QListWidget) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlistwidget.h:207
// index:0
// Public
// void QListWidget(class QWidget *)
func NewQListWidget(parent unsafe.Pointer) *QListWidget {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidgetC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQListWidgetFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qlistwidget.h:208
// index:0
// Public virtual
// void ~QListWidget()
func DeleteQListWidget(*QListWidget) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidgetD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:210
// index:0
// Public virtual
// void setSelectionModel(class QItemSelectionModel *)
func (this *QListWidget) SetSelectionModel(selectionModel unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget17setSelectionModelEP19QItemSelectionModel", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), selectionModel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:212
// index:0
// Public
// QListWidgetItem * item(int)
func (this *QListWidget) Item(row int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget4itemEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlistwidget.h:213
// index:0
// Public
// int row(const class QListWidgetItem *)
func (this *QListWidget) Row(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget3rowEPK15QListWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlistwidget.h:214
// index:0
// Public
// void insertItem(int, class QListWidgetItem *)
func (this *QListWidget) InsertItem(row int, item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget10insertItemEiP15QListWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:215
// index:1
// Public
// void insertItem(int, const class QString &)
func (this *QListWidget) InsertItem_1(row int, label *qtcore.QString) {
	var convArg1 = label.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget10insertItemEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:216
// index:0
// Public
// void insertItems(int, const class QStringList &)
func (this *QListWidget) InsertItems(row int, labels *qtcore.QStringList) {
	var convArg1 = labels.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget11insertItemsEiRK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:217
// index:0
// Public inline
// void addItem(const class QString &)
func (this *QListWidget) AddItem(label *qtcore.QString) {
	var convArg0 = label.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget7addItemERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:218
// index:1
// Public inline
// void addItem(class QListWidgetItem *)
func (this *QListWidget) AddItem_1(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget7addItemEP15QListWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:219
// index:0
// Public inline
// void addItems(const class QStringList &)
func (this *QListWidget) AddItems(labels *qtcore.QStringList) {
	var convArg0 = labels.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget8addItemsERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:220
// index:0
// Public
// QListWidgetItem * takeItem(int)
func (this *QListWidget) TakeItem(row int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget8takeItemEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlistwidget.h:221
// index:0
// Public
// int count()
func (this *QListWidget) Count() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlistwidget.h:223
// index:0
// Public
// QListWidgetItem * currentItem()
func (this *QListWidget) CurrentItem() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget11currentItemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlistwidget.h:224
// index:0
// Public
// void setCurrentItem(class QListWidgetItem *)
func (this *QListWidget) SetCurrentItem(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget14setCurrentItemEP15QListWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:227
// index:0
// Public
// int currentRow()
func (this *QListWidget) CurrentRow() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget10currentRowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlistwidget.h:228
// index:0
// Public
// void setCurrentRow(int)
func (this *QListWidget) SetCurrentRow(row int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget13setCurrentRowEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:231
// index:0
// Public
// QListWidgetItem * itemAt(const class QPoint &)
func (this *QListWidget) ItemAt(p *qtcore.QPoint) interface{} {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget6itemAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlistwidget.h:232
// index:1
// Public inline
// QListWidgetItem * itemAt(int, int)
func (this *QListWidget) ItemAt_1(x int, y int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget6itemAtEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlistwidget.h:233
// index:0
// Public
// QRect visualItemRect(const class QListWidgetItem *)
func (this *QListWidget) VisualItemRect(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget14visualItemRectEPK15QListWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlistwidget.h:235
// index:0
// Public
// void sortItems(Qt::SortOrder)
func (this *QListWidget) SortItems(order int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget9sortItemsEN2Qt9SortOrderE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:236
// index:0
// Public
// void setSortingEnabled(_Bool)
func (this *QListWidget) SetSortingEnabled(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget17setSortingEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:237
// index:0
// Public
// bool isSortingEnabled()
func (this *QListWidget) IsSortingEnabled() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget16isSortingEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlistwidget.h:239
// index:0
// Public
// void editItem(class QListWidgetItem *)
func (this *QListWidget) EditItem(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget8editItemEP15QListWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:240
// index:0
// Public
// void openPersistentEditor(class QListWidgetItem *)
func (this *QListWidget) OpenPersistentEditor(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget20openPersistentEditorEP15QListWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:241
// index:0
// Public
// void closePersistentEditor(class QListWidgetItem *)
func (this *QListWidget) ClosePersistentEditor(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget21closePersistentEditorEP15QListWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:243
// index:0
// Public
// bool isPersistentEditorOpen(class QListWidgetItem *)
func (this *QListWidget) IsPersistentEditorOpen(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget22isPersistentEditorOpenEP15QListWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlistwidget.h:245
// index:0
// Public
// QWidget * itemWidget(class QListWidgetItem *)
func (this *QListWidget) ItemWidget(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget10itemWidgetEP15QListWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlistwidget.h:246
// index:0
// Public
// void setItemWidget(class QListWidgetItem *, class QWidget *)
func (this *QListWidget) SetItemWidget(item unsafe.Pointer, widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget13setItemWidgetEP15QListWidgetItemP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:247
// index:0
// Public inline
// void removeItemWidget(class QListWidgetItem *)
func (this *QListWidget) RemoveItemWidget(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget16removeItemWidgetEP15QListWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:249
// index:0
// Public
// bool isItemSelected(const class QListWidgetItem *)
func (this *QListWidget) IsItemSelected(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget14isItemSelectedEPK15QListWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlistwidget.h:250
// index:0
// Public
// void setItemSelected(const class QListWidgetItem *, _Bool)
func (this *QListWidget) SetItemSelected(item unsafe.Pointer, select_ bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget15setItemSelectedEPK15QListWidgetItemb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &select_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:251
// index:0
// Public
// QList<QListWidgetItem *> selectedItems()
func (this *QListWidget) SelectedItems() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget13selectedItemsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlistwidget.h:254
// index:0
// Public
// bool isItemHidden(const class QListWidgetItem *)
func (this *QListWidget) IsItemHidden(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget12isItemHiddenEPK15QListWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlistwidget.h:255
// index:0
// Public
// void setItemHidden(const class QListWidgetItem *, _Bool)
func (this *QListWidget) SetItemHidden(item unsafe.Pointer, hide bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget13setItemHiddenEPK15QListWidgetItemb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:260
// index:0
// Public virtual
// void dropEvent(class QDropEvent *)
func (this *QListWidget) DropEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget9dropEventEP10QDropEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:263
// index:0
// Public
// void scrollToItem(const class QListWidgetItem *, class QAbstractItemView::ScrollHint)
func (this *QListWidget) ScrollToItem(item unsafe.Pointer, hint int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget12scrollToItemEPK15QListWidgetItemN17QAbstractItemView10ScrollHintE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:264
// index:0
// Public
// void clear()
func (this *QListWidget) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:267
// index:0
// Public
// void itemPressed(class QListWidgetItem *)
func (this *QListWidget) ItemPressed(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget11itemPressedEP15QListWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:268
// index:0
// Public
// void itemClicked(class QListWidgetItem *)
func (this *QListWidget) ItemClicked(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget11itemClickedEP15QListWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:269
// index:0
// Public
// void itemDoubleClicked(class QListWidgetItem *)
func (this *QListWidget) ItemDoubleClicked(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget17itemDoubleClickedEP15QListWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:270
// index:0
// Public
// void itemActivated(class QListWidgetItem *)
func (this *QListWidget) ItemActivated(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget13itemActivatedEP15QListWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:271
// index:0
// Public
// void itemEntered(class QListWidgetItem *)
func (this *QListWidget) ItemEntered(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget11itemEnteredEP15QListWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:272
// index:0
// Public
// void itemChanged(class QListWidgetItem *)
func (this *QListWidget) ItemChanged(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget11itemChangedEP15QListWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:274
// index:0
// Public
// void currentItemChanged(class QListWidgetItem *, class QListWidgetItem *)
func (this *QListWidget) CurrentItemChanged(current unsafe.Pointer, previous unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget18currentItemChangedEP15QListWidgetItemS1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), current, previous)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:275
// index:0
// Public
// void currentTextChanged(const class QString &)
func (this *QListWidget) CurrentTextChanged(currentText *qtcore.QString) {
	var convArg0 = currentText.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget18currentTextChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:276
// index:0
// Public
// void currentRowChanged(int)
func (this *QListWidget) CurrentRowChanged(currentRow int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget17currentRowChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &currentRow)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:278
// index:0
// Public
// void itemSelectionChanged()
func (this *QListWidget) ItemSelectionChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget20itemSelectionChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:281
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QListWidget) Event(e unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlistwidget.h:282
// index:0
// Protected virtual
// QStringList mimeTypes()
func (this *QListWidget) MimeTypes() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget9mimeTypesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlistwidget.h:289
// index:0
// Protected virtual
// bool dropMimeData(int, const class QMimeData *, Qt::DropAction)
func (this *QListWidget) DropMimeData(index int, data unsafe.Pointer, action int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget12dropMimeDataEiPK9QMimeDataN2Qt10DropActionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, data, &action)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlistwidget.h:290
// index:0
// Protected virtual
// Qt::DropActions supportedDropActions()
func (this *QListWidget) SupportedDropActions() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget20supportedDropActionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlistwidget.h:298
// index:0
// Protected
// QList<QListWidgetItem *> items(const class QMimeData *)
func (this *QListWidget) Items(data unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget5itemsEPK9QMimeData", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), data)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlistwidget.h:300
// index:0
// Protected
// QModelIndex indexFromItem(class QListWidgetItem *)
func (this *QListWidget) IndexFromItem(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget13indexFromItemEP15QListWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlistwidget.h:301
// index:0
// Protected
// QListWidgetItem * itemFromIndex(const class QModelIndex &)
func (this *QListWidget) ItemFromIndex(index *qtcore.QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget13itemFromIndexERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
