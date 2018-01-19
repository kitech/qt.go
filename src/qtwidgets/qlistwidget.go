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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qlistwidget.h:199
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QListWidget) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:207
// index:0
// void QListWidget(class QWidget *)
func NewQListWidget(parent unsafe.Pointer) *QListWidget {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidgetC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QListWidget{cthis}
}

// /usr/include/qt/QtWidgets/qlistwidget.h:208
// index:0
// virtual
// void ~QListWidget()
func DeleteQListWidget(*QListWidget) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidgetD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:210
// index:0
// virtual
// void setSelectionModel(class QItemSelectionModel *)
func (this *QListWidget) SetSelectionModel(selectionModel unsafe.Pointer) {
	// 0: (, QItemSelectionModel * selectionModel), (selectionModel)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget17setSelectionModelEP19QItemSelectionModel", ffiqt.FFI_TYPE_VOID, this.cthis, selectionModel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:212
// index:0
// QListWidgetItem * item(int)
func (this *QListWidget) Item(row int) {
	// 0: (, int row), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget4itemEi", ffiqt.FFI_TYPE_VOID, this.cthis, &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:213
// index:0
// int row(const class QListWidgetItem *)
func (this *QListWidget) Row(item unsafe.Pointer) {
	// 0: (, const QListWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget3rowEPK15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:214
// index:0
// void insertItem(int, class QListWidgetItem *)
func (this *QListWidget) InsertItem(row int, item unsafe.Pointer) {
	// 0: (, int row, QListWidgetItem * item), (&row, item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget10insertItemEiP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, &row, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:215
// index:1
// void insertItem(int, const class QString &)
func (this *QListWidget) InsertItem_1(row int, label unsafe.Pointer) {
	// 1: (, int row, const QString & label), (&row, label)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget10insertItemEiRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, &row, label)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:216
// index:0
// void insertItems(int, const class QStringList &)
func (this *QListWidget) InsertItems(row int, labels unsafe.Pointer) {
	// 0: (, int row, const QStringList & labels), (&row, labels)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget11insertItemsEiRK11QStringList", ffiqt.FFI_TYPE_VOID, this.cthis, &row, labels)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:217
// index:0
// inline
// void addItem(const class QString &)
func (this *QListWidget) AddItem(label unsafe.Pointer) {
	// 0: (, const QString & label), (label)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget7addItemERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, label)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:218
// index:1
// inline
// void addItem(class QListWidgetItem *)
func (this *QListWidget) AddItem_1(item unsafe.Pointer) {
	// 1: (, QListWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget7addItemEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:219
// index:0
// inline
// void addItems(const class QStringList &)
func (this *QListWidget) AddItems(labels unsafe.Pointer) {
	// 0: (, const QStringList & labels), (labels)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget8addItemsERK11QStringList", ffiqt.FFI_TYPE_VOID, this.cthis, labels)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:220
// index:0
// QListWidgetItem * takeItem(int)
func (this *QListWidget) TakeItem(row int) {
	// 0: (, int row), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget8takeItemEi", ffiqt.FFI_TYPE_VOID, this.cthis, &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:221
// index:0
// int count()
func (this *QListWidget) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget5countEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:223
// index:0
// QListWidgetItem * currentItem()
func (this *QListWidget) CurrentItem() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget11currentItemEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:224
// index:0
// void setCurrentItem(class QListWidgetItem *)
func (this *QListWidget) SetCurrentItem(item unsafe.Pointer) {
	// 0: (, QListWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget14setCurrentItemEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:227
// index:0
// int currentRow()
func (this *QListWidget) CurrentRow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget10currentRowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:228
// index:0
// void setCurrentRow(int)
func (this *QListWidget) SetCurrentRow(row int) {
	// 0: (, int row), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget13setCurrentRowEi", ffiqt.FFI_TYPE_VOID, this.cthis, &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:231
// index:0
// QListWidgetItem * itemAt(const class QPoint &)
func (this *QListWidget) ItemAt(p unsafe.Pointer) {
	// 0: (, const QPoint & p), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget6itemAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:232
// index:1
// inline
// QListWidgetItem * itemAt(int, int)
func (this *QListWidget) ItemAt_1(x int, y int) {
	// 1: (, int x, int y), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget6itemAtEii", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:233
// index:0
// QRect visualItemRect(const class QListWidgetItem *)
func (this *QListWidget) VisualItemRect(item unsafe.Pointer) {
	// 0: (, const QListWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget14visualItemRectEPK15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:235
// index:0
// void sortItems(Qt::SortOrder)
func (this *QListWidget) SortItems(order int) {
	// 0: (, Qt::SortOrder order), (&order)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget9sortItemsEN2Qt9SortOrderE", ffiqt.FFI_TYPE_VOID, this.cthis, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:236
// index:0
// void setSortingEnabled(_Bool)
func (this *QListWidget) SetSortingEnabled(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget17setSortingEnabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:237
// index:0
// bool isSortingEnabled()
func (this *QListWidget) IsSortingEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget16isSortingEnabledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:239
// index:0
// void editItem(class QListWidgetItem *)
func (this *QListWidget) EditItem(item unsafe.Pointer) {
	// 0: (, QListWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget8editItemEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:240
// index:0
// void openPersistentEditor(class QListWidgetItem *)
func (this *QListWidget) OpenPersistentEditor(item unsafe.Pointer) {
	// 0: (, QListWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget20openPersistentEditorEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:241
// index:0
// void closePersistentEditor(class QListWidgetItem *)
func (this *QListWidget) ClosePersistentEditor(item unsafe.Pointer) {
	// 0: (, QListWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget21closePersistentEditorEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:243
// index:0
// bool isPersistentEditorOpen(class QListWidgetItem *)
func (this *QListWidget) IsPersistentEditorOpen(item unsafe.Pointer) {
	// 0: (, QListWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget22isPersistentEditorOpenEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:245
// index:0
// QWidget * itemWidget(class QListWidgetItem *)
func (this *QListWidget) ItemWidget(item unsafe.Pointer) {
	// 0: (, QListWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget10itemWidgetEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:246
// index:0
// void setItemWidget(class QListWidgetItem *, class QWidget *)
func (this *QListWidget) SetItemWidget(item unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, QListWidgetItem * item, QWidget * widget), (item, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget13setItemWidgetEP15QListWidgetItemP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, item, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:247
// index:0
// inline
// void removeItemWidget(class QListWidgetItem *)
func (this *QListWidget) RemoveItemWidget(item unsafe.Pointer) {
	// 0: (, QListWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget16removeItemWidgetEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:249
// index:0
// bool isItemSelected(const class QListWidgetItem *)
func (this *QListWidget) IsItemSelected(item unsafe.Pointer) {
	// 0: (, const QListWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget14isItemSelectedEPK15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:250
// index:0
// void setItemSelected(const class QListWidgetItem *, _Bool)
func (this *QListWidget) SetItemSelected(item unsafe.Pointer, select_ bool) {
	// 0: (, const QListWidgetItem * item, bool select), (item, &select_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget15setItemSelectedEPK15QListWidgetItemb", ffiqt.FFI_TYPE_VOID, this.cthis, item, &select_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:251
// index:0
// QList<QListWidgetItem *> selectedItems()
func (this *QListWidget) SelectedItems() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget13selectedItemsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:254
// index:0
// bool isItemHidden(const class QListWidgetItem *)
func (this *QListWidget) IsItemHidden(item unsafe.Pointer) {
	// 0: (, const QListWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget12isItemHiddenEPK15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:255
// index:0
// void setItemHidden(const class QListWidgetItem *, _Bool)
func (this *QListWidget) SetItemHidden(item unsafe.Pointer, hide bool) {
	// 0: (, const QListWidgetItem * item, bool hide), (item, &hide)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget13setItemHiddenEPK15QListWidgetItemb", ffiqt.FFI_TYPE_VOID, this.cthis, item, &hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:260
// index:0
// virtual
// void dropEvent(class QDropEvent *)
func (this *QListWidget) DropEvent(event unsafe.Pointer) {
	// 0: (, QDropEvent * event), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget9dropEventEP10QDropEvent", ffiqt.FFI_TYPE_VOID, this.cthis, event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:263
// index:0
// void scrollToItem(const class QListWidgetItem *, class QAbstractItemView::ScrollHint)
func (this *QListWidget) ScrollToItem(item unsafe.Pointer, hint int) {
	// 0: (, const QListWidgetItem * item, QAbstractItemView::ScrollHint hint), (item, &hint)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget12scrollToItemEPK15QListWidgetItemN17QAbstractItemView10ScrollHintE", ffiqt.FFI_TYPE_VOID, this.cthis, item, &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:264
// index:0
// void clear()
func (this *QListWidget) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget5clearEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:267
// index:0
// void itemPressed(class QListWidgetItem *)
func (this *QListWidget) ItemPressed(item unsafe.Pointer) {
	// 0: (, QListWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget11itemPressedEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:268
// index:0
// void itemClicked(class QListWidgetItem *)
func (this *QListWidget) ItemClicked(item unsafe.Pointer) {
	// 0: (, QListWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget11itemClickedEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:269
// index:0
// void itemDoubleClicked(class QListWidgetItem *)
func (this *QListWidget) ItemDoubleClicked(item unsafe.Pointer) {
	// 0: (, QListWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget17itemDoubleClickedEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:270
// index:0
// void itemActivated(class QListWidgetItem *)
func (this *QListWidget) ItemActivated(item unsafe.Pointer) {
	// 0: (, QListWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget13itemActivatedEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:271
// index:0
// void itemEntered(class QListWidgetItem *)
func (this *QListWidget) ItemEntered(item unsafe.Pointer) {
	// 0: (, QListWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget11itemEnteredEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:272
// index:0
// void itemChanged(class QListWidgetItem *)
func (this *QListWidget) ItemChanged(item unsafe.Pointer) {
	// 0: (, QListWidgetItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget11itemChangedEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:274
// index:0
// void currentItemChanged(class QListWidgetItem *, class QListWidgetItem *)
func (this *QListWidget) CurrentItemChanged(current unsafe.Pointer, previous unsafe.Pointer) {
	// 0: (, QListWidgetItem * current, QListWidgetItem * previous), (current, previous)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget18currentItemChangedEP15QListWidgetItemS1_", ffiqt.FFI_TYPE_VOID, this.cthis, current, previous)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:275
// index:0
// void currentTextChanged(const class QString &)
func (this *QListWidget) CurrentTextChanged(currentText unsafe.Pointer) {
	// 0: (, const QString & currentText), (currentText)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget18currentTextChangedERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, currentText)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:276
// index:0
// void currentRowChanged(int)
func (this *QListWidget) CurrentRowChanged(currentRow int) {
	// 0: (, int currentRow), (&currentRow)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget17currentRowChangedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &currentRow)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:278
// index:0
// void itemSelectionChanged()
func (this *QListWidget) ItemSelectionChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget20itemSelectionChangedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
