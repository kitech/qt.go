//  header block begin
// /usr/include/qt/QtWidgets/qlistwidget.h
// #include <qlistwidget.h>
// #include <QtWidgets>
package qtwidgets

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

// /usr/include/qt/QtWidgets/qlistwidget.h:199
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QListWidget) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:207
// index:0
// void QListWidget(class QWidget *)
func NewQListWidget(parent unsafe.Pointer) *QListWidget {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidgetC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQListWidgetFromPointer(cthis)
	return gothis
}
func NewQListWidgetFromPointer(cthis unsafe.Pointer) *QListWidget {
	bcthis0 := NewQListViewFromPointer(cthis)
	return &QListWidget{bcthis0}
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
	// 0: (, selectionModel QItemSelectionModel *), (selectionModel)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget17setSelectionModelEP19QItemSelectionModel", ffiqt.FFI_TYPE_VOID, this.GetCthis(), selectionModel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:212
// index:0
// QListWidgetItem * item(int)
func (this *QListWidget) Item(row int) {
	// 0: (, row int), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget4itemEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:213
// index:0
// int row(const class QListWidgetItem *)
func (this *QListWidget) Row(item unsafe.Pointer) {
	// 0: (, item const QListWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget3rowEPK15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:214
// index:0
// void insertItem(int, class QListWidgetItem *)
func (this *QListWidget) InsertItem(row int, item unsafe.Pointer) {
	// 0: (, row int, item QListWidgetItem *), (&row, item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget10insertItemEiP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:215
// index:1
// void insertItem(int, const class QString &)
func (this *QListWidget) InsertItem_1(row int, label unsafe.Pointer) {
	// 1: (, row int, label const QString &), (&row, label)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget10insertItemEiRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, label)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:216
// index:0
// void insertItems(int, const class QStringList &)
func (this *QListWidget) InsertItems(row int, labels unsafe.Pointer) {
	// 0: (, row int, labels const QStringList &), (&row, labels)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget11insertItemsEiRK11QStringList", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, labels)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:217
// index:0
// inline
// void addItem(const class QString &)
func (this *QListWidget) AddItem(label unsafe.Pointer) {
	// 0: (, label const QString &), (label)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget7addItemERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), label)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:218
// index:1
// inline
// void addItem(class QListWidgetItem *)
func (this *QListWidget) AddItem_1(item unsafe.Pointer) {
	// 1: (, item QListWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget7addItemEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:219
// index:0
// inline
// void addItems(const class QStringList &)
func (this *QListWidget) AddItems(labels unsafe.Pointer) {
	// 0: (, labels const QStringList &), (labels)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget8addItemsERK11QStringList", ffiqt.FFI_TYPE_VOID, this.GetCthis(), labels)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:220
// index:0
// QListWidgetItem * takeItem(int)
func (this *QListWidget) TakeItem(row int) {
	// 0: (, row int), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget8takeItemEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:221
// index:0
// int count()
func (this *QListWidget) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget5countEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:223
// index:0
// QListWidgetItem * currentItem()
func (this *QListWidget) CurrentItem() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget11currentItemEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:224
// index:0
// void setCurrentItem(class QListWidgetItem *)
func (this *QListWidget) SetCurrentItem(item unsafe.Pointer) {
	// 0: (, item QListWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget14setCurrentItemEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:225
// index:1
// void setCurrentItem(class QListWidgetItem *, class QItemSelectionModel::SelectionFlags)
func (this *QListWidget) SetCurrentItem_1(item unsafe.Pointer, command int) {
	// 1: (, item QListWidgetItem *, QFlags<QItemSelectionModel::SelectionFlag> command), (item, &command)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget14setCurrentItemEP15QListWidgetItem6QFlagsIN19QItemSelectionModel13SelectionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &command)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:227
// index:0
// int currentRow()
func (this *QListWidget) CurrentRow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget10currentRowEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:228
// index:0
// void setCurrentRow(int)
func (this *QListWidget) SetCurrentRow(row int) {
	// 0: (, row int), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget13setCurrentRowEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:229
// index:1
// void setCurrentRow(int, class QItemSelectionModel::SelectionFlags)
func (this *QListWidget) SetCurrentRow_1(row int, command int) {
	// 1: (, row int, QFlags<QItemSelectionModel::SelectionFlag> command), (&row, &command)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget13setCurrentRowEi6QFlagsIN19QItemSelectionModel13SelectionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &command)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:231
// index:0
// QListWidgetItem * itemAt(const class QPoint &)
func (this *QListWidget) ItemAt(p unsafe.Pointer) {
	// 0: (, p const QPoint &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget6itemAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:232
// index:1
// inline
// QListWidgetItem * itemAt(int, int)
func (this *QListWidget) ItemAt_1(x int, y int) {
	// 1: (, x int, y int), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget6itemAtEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:233
// index:0
// QRect visualItemRect(const class QListWidgetItem *)
func (this *QListWidget) VisualItemRect(item unsafe.Pointer) {
	// 0: (, item const QListWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget14visualItemRectEPK15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:235
// index:0
// void sortItems(Qt::SortOrder)
func (this *QListWidget) SortItems(order int) {
	// 0: (, order Qt::SortOrder), (&order)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget9sortItemsEN2Qt9SortOrderE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:236
// index:0
// void setSortingEnabled(_Bool)
func (this *QListWidget) SetSortingEnabled(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget17setSortingEnabledEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:237
// index:0
// bool isSortingEnabled()
func (this *QListWidget) IsSortingEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget16isSortingEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:239
// index:0
// void editItem(class QListWidgetItem *)
func (this *QListWidget) EditItem(item unsafe.Pointer) {
	// 0: (, item QListWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget8editItemEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:240
// index:0
// void openPersistentEditor(class QListWidgetItem *)
func (this *QListWidget) OpenPersistentEditor(item unsafe.Pointer) {
	// 0: (, item QListWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget20openPersistentEditorEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:241
// index:0
// void closePersistentEditor(class QListWidgetItem *)
func (this *QListWidget) ClosePersistentEditor(item unsafe.Pointer) {
	// 0: (, item QListWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget21closePersistentEditorEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:243
// index:0
// bool isPersistentEditorOpen(class QListWidgetItem *)
func (this *QListWidget) IsPersistentEditorOpen(item unsafe.Pointer) {
	// 0: (, item QListWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget22isPersistentEditorOpenEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:245
// index:0
// QWidget * itemWidget(class QListWidgetItem *)
func (this *QListWidget) ItemWidget(item unsafe.Pointer) {
	// 0: (, item QListWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget10itemWidgetEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:246
// index:0
// void setItemWidget(class QListWidgetItem *, class QWidget *)
func (this *QListWidget) SetItemWidget(item unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, item QListWidgetItem *, widget QWidget *), (item, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget13setItemWidgetEP15QListWidgetItemP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:247
// index:0
// inline
// void removeItemWidget(class QListWidgetItem *)
func (this *QListWidget) RemoveItemWidget(item unsafe.Pointer) {
	// 0: (, item QListWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget16removeItemWidgetEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:249
// index:0
// bool isItemSelected(const class QListWidgetItem *)
func (this *QListWidget) IsItemSelected(item unsafe.Pointer) {
	// 0: (, item const QListWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget14isItemSelectedEPK15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:250
// index:0
// void setItemSelected(const class QListWidgetItem *, _Bool)
func (this *QListWidget) SetItemSelected(item unsafe.Pointer, select_ bool) {
	// 0: (, item const QListWidgetItem *, select bool), (item, &select_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget15setItemSelectedEPK15QListWidgetItemb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &select_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:251
// index:0
// QList<QListWidgetItem *> selectedItems()
func (this *QListWidget) SelectedItems() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget13selectedItemsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:252
// index:0
// QList<QListWidgetItem *> findItems(const class QString &, Qt::MatchFlags)
func (this *QListWidget) FindItems(text unsafe.Pointer, flags int) {
	// 0: (, text const QString &, QFlags<Qt::MatchFlag> flags), (text, &flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget9findItemsERK7QString6QFlagsIN2Qt9MatchFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text, &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:254
// index:0
// bool isItemHidden(const class QListWidgetItem *)
func (this *QListWidget) IsItemHidden(item unsafe.Pointer) {
	// 0: (, item const QListWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget12isItemHiddenEPK15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:255
// index:0
// void setItemHidden(const class QListWidgetItem *, _Bool)
func (this *QListWidget) SetItemHidden(item unsafe.Pointer, hide bool) {
	// 0: (, item const QListWidgetItem *, hide bool), (item, &hide)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget13setItemHiddenEPK15QListWidgetItemb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:260
// index:0
// virtual
// void dropEvent(class QDropEvent *)
func (this *QListWidget) DropEvent(event unsafe.Pointer) {
	// 0: (, event QDropEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget9dropEventEP10QDropEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:263
// index:0
// void scrollToItem(const class QListWidgetItem *, class QAbstractItemView::ScrollHint)
func (this *QListWidget) ScrollToItem(item unsafe.Pointer, hint int) {
	// 0: (, item const QListWidgetItem *, hint QAbstractItemView::ScrollHint), (item, &hint)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget12scrollToItemEPK15QListWidgetItemN17QAbstractItemView10ScrollHintE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:264
// index:0
// void clear()
func (this *QListWidget) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget5clearEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:267
// index:0
// void itemPressed(class QListWidgetItem *)
func (this *QListWidget) ItemPressed(item unsafe.Pointer) {
	// 0: (, item QListWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget11itemPressedEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:268
// index:0
// void itemClicked(class QListWidgetItem *)
func (this *QListWidget) ItemClicked(item unsafe.Pointer) {
	// 0: (, item QListWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget11itemClickedEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:269
// index:0
// void itemDoubleClicked(class QListWidgetItem *)
func (this *QListWidget) ItemDoubleClicked(item unsafe.Pointer) {
	// 0: (, item QListWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget17itemDoubleClickedEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:270
// index:0
// void itemActivated(class QListWidgetItem *)
func (this *QListWidget) ItemActivated(item unsafe.Pointer) {
	// 0: (, item QListWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget13itemActivatedEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:271
// index:0
// void itemEntered(class QListWidgetItem *)
func (this *QListWidget) ItemEntered(item unsafe.Pointer) {
	// 0: (, item QListWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget11itemEnteredEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:272
// index:0
// void itemChanged(class QListWidgetItem *)
func (this *QListWidget) ItemChanged(item unsafe.Pointer) {
	// 0: (, item QListWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget11itemChangedEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:274
// index:0
// void currentItemChanged(class QListWidgetItem *, class QListWidgetItem *)
func (this *QListWidget) CurrentItemChanged(current unsafe.Pointer, previous unsafe.Pointer) {
	// 0: (, current QListWidgetItem *, previous QListWidgetItem *), (current, previous)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget18currentItemChangedEP15QListWidgetItemS1_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), current, previous)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:275
// index:0
// void currentTextChanged(const class QString &)
func (this *QListWidget) CurrentTextChanged(currentText unsafe.Pointer) {
	// 0: (, currentText const QString &), (currentText)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget18currentTextChangedERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), currentText)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:276
// index:0
// void currentRowChanged(int)
func (this *QListWidget) CurrentRowChanged(currentRow int) {
	// 0: (, currentRow int), (&currentRow)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget17currentRowChangedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &currentRow)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:278
// index:0
// void itemSelectionChanged()
func (this *QListWidget) ItemSelectionChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget20itemSelectionChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:281
// index:0
// virtual
// bool event(class QEvent *)
func (this *QListWidget) Event(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:282
// index:0
// virtual
// QStringList mimeTypes()
func (this *QListWidget) MimeTypes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget9mimeTypesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:289
// index:0
// virtual
// bool dropMimeData(int, const class QMimeData *, Qt::DropAction)
func (this *QListWidget) DropMimeData(index int, data unsafe.Pointer, action int) {
	// 0: (, index int, data const QMimeData *, action Qt::DropAction), (&index, data, &action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QListWidget12dropMimeDataEiPK9QMimeDataN2Qt10DropActionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, data, &action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:290
// index:0
// virtual
// Qt::DropActions supportedDropActions()
func (this *QListWidget) SupportedDropActions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget20supportedDropActionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:298
// index:0
// QList<QListWidgetItem *> items(const class QMimeData *)
func (this *QListWidget) Items(data unsafe.Pointer) {
	// 0: (, data const QMimeData *), (data)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget5itemsEPK9QMimeData", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:300
// index:0
// QModelIndex indexFromItem(class QListWidgetItem *)
func (this *QListWidget) IndexFromItem(item unsafe.Pointer) {
	// 0: (, item QListWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget13indexFromItemEP15QListWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:301
// index:0
// QListWidgetItem * itemFromIndex(const class QModelIndex &)
func (this *QListWidget) ItemFromIndex(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QListWidget13itemFromIndexERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

//  body block end
