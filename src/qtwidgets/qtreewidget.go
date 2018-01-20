//  header block begin
// /usr/include/qt/QtWidgets/qtreewidget.h
// #include <qtreewidget.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 68
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
type QTreeWidget struct {
	*QTreeView
}

func (this *QTreeWidget) GetCthis() unsafe.Pointer {
	return this.QTreeView.GetCthis()
}

// /usr/include/qt/QtWidgets/qtreewidget.h:257
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QTreeWidget) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:264
// index:0
// void QTreeWidget(class QWidget *)
func NewQTreeWidget(parent unsafe.Pointer) *QTreeWidget {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidgetC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQTreeWidgetFromPointer(cthis)
	return gothis
}
func NewQTreeWidgetFromPointer(cthis unsafe.Pointer) *QTreeWidget {
	bcthis0 := NewQTreeViewFromPointer(cthis)
	return &QTreeWidget{bcthis0}
}

// /usr/include/qt/QtWidgets/qtreewidget.h:265
// index:0
// virtual
// void ~QTreeWidget()
func DeleteQTreeWidget(*QTreeWidget) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidgetD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:267
// index:0
// int columnCount()
func (this *QTreeWidget) ColumnCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget11columnCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:268
// index:0
// void setColumnCount(int)
func (this *QTreeWidget) SetColumnCount(columns int) {
	// 0: (, columns int), (&columns)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget14setColumnCountEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &columns)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:270
// index:0
// QTreeWidgetItem * invisibleRootItem()
func (this *QTreeWidget) InvisibleRootItem() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget17invisibleRootItemEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:271
// index:0
// QTreeWidgetItem * topLevelItem(int)
func (this *QTreeWidget) TopLevelItem(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget12topLevelItemEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:272
// index:0
// int topLevelItemCount()
func (this *QTreeWidget) TopLevelItemCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget17topLevelItemCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:273
// index:0
// void insertTopLevelItem(int, class QTreeWidgetItem *)
func (this *QTreeWidget) InsertTopLevelItem(index int, item unsafe.Pointer) {
	// 0: (, index int, item QTreeWidgetItem *), (&index, item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget18insertTopLevelItemEiP15QTreeWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:274
// index:0
// void addTopLevelItem(class QTreeWidgetItem *)
func (this *QTreeWidget) AddTopLevelItem(item unsafe.Pointer) {
	// 0: (, item QTreeWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget15addTopLevelItemEP15QTreeWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:275
// index:0
// QTreeWidgetItem * takeTopLevelItem(int)
func (this *QTreeWidget) TakeTopLevelItem(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget16takeTopLevelItemEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:276
// index:0
// int indexOfTopLevelItem(class QTreeWidgetItem *)
func (this *QTreeWidget) IndexOfTopLevelItem(item unsafe.Pointer) {
	// 0: (, item QTreeWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget19indexOfTopLevelItemEP15QTreeWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:281
// index:0
// QTreeWidgetItem * headerItem()
func (this *QTreeWidget) HeaderItem() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget10headerItemEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:282
// index:0
// void setHeaderItem(class QTreeWidgetItem *)
func (this *QTreeWidget) SetHeaderItem(item unsafe.Pointer) {
	// 0: (, item QTreeWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget13setHeaderItemEP15QTreeWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:283
// index:0
// void setHeaderLabels(const class QStringList &)
func (this *QTreeWidget) SetHeaderLabels(labels unsafe.Pointer) {
	// 0: (, labels const QStringList &), (labels)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget15setHeaderLabelsERK11QStringList", ffiqt.FFI_TYPE_VOID, this.GetCthis(), labels)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:284
// index:0
// inline
// void setHeaderLabel(const class QString &)
func (this *QTreeWidget) SetHeaderLabel(label unsafe.Pointer) {
	// 0: (, label const QString &), (label)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget14setHeaderLabelERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), label)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:286
// index:0
// QTreeWidgetItem * currentItem()
func (this *QTreeWidget) CurrentItem() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget11currentItemEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:287
// index:0
// int currentColumn()
func (this *QTreeWidget) CurrentColumn() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget13currentColumnEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:288
// index:0
// void setCurrentItem(class QTreeWidgetItem *)
func (this *QTreeWidget) SetCurrentItem(item unsafe.Pointer) {
	// 0: (, item QTreeWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:289
// index:1
// void setCurrentItem(class QTreeWidgetItem *, int)
func (this *QTreeWidget) SetCurrentItem_1(item unsafe.Pointer, column int) {
	// 1: (, item QTreeWidgetItem *, column int), (item, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:290
// index:2
// void setCurrentItem(class QTreeWidgetItem *, int, class QItemSelectionModel::SelectionFlags)
func (this *QTreeWidget) SetCurrentItem_2(item unsafe.Pointer, column int, command int) {
	// 2: (, item QTreeWidgetItem *, column int, QFlags<QItemSelectionModel::SelectionFlag> command), (item, &column, &command)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItemi6QFlagsIN19QItemSelectionModel13SelectionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &column, &command)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:292
// index:0
// QTreeWidgetItem * itemAt(const class QPoint &)
func (this *QTreeWidget) ItemAt(p unsafe.Pointer) {
	// 0: (, p const QPoint &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget6itemAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:293
// index:1
// inline
// QTreeWidgetItem * itemAt(int, int)
func (this *QTreeWidget) ItemAt_1(x int, y int) {
	// 1: (, x int, y int), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget6itemAtEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:294
// index:0
// QRect visualItemRect(const class QTreeWidgetItem *)
func (this *QTreeWidget) VisualItemRect(item unsafe.Pointer) {
	// 0: (, item const QTreeWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget14visualItemRectEPK15QTreeWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:296
// index:0
// int sortColumn()
func (this *QTreeWidget) SortColumn() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget10sortColumnEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:297
// index:0
// void sortItems(int, Qt::SortOrder)
func (this *QTreeWidget) SortItems(column int, order int) {
	// 0: (, column int, order Qt::SortOrder), (&column, &order)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget9sortItemsEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:299
// index:0
// void editItem(class QTreeWidgetItem *, int)
func (this *QTreeWidget) EditItem(item unsafe.Pointer, column int) {
	// 0: (, item QTreeWidgetItem *, column int), (item, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget8editItemEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:300
// index:0
// void openPersistentEditor(class QTreeWidgetItem *, int)
func (this *QTreeWidget) OpenPersistentEditor(item unsafe.Pointer, column int) {
	// 0: (, item QTreeWidgetItem *, column int), (item, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget20openPersistentEditorEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:301
// index:0
// void closePersistentEditor(class QTreeWidgetItem *, int)
func (this *QTreeWidget) ClosePersistentEditor(item unsafe.Pointer, column int) {
	// 0: (, item QTreeWidgetItem *, column int), (item, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget21closePersistentEditorEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:303
// index:0
// bool isPersistentEditorOpen(class QTreeWidgetItem *, int)
func (this *QTreeWidget) IsPersistentEditorOpen(item unsafe.Pointer, column int) {
	// 0: (, item QTreeWidgetItem *, column int), (item, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget22isPersistentEditorOpenEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:305
// index:0
// QWidget * itemWidget(class QTreeWidgetItem *, int)
func (this *QTreeWidget) ItemWidget(item unsafe.Pointer, column int) {
	// 0: (, item QTreeWidgetItem *, column int), (item, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget10itemWidgetEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:306
// index:0
// void setItemWidget(class QTreeWidgetItem *, int, class QWidget *)
func (this *QTreeWidget) SetItemWidget(item unsafe.Pointer, column int, widget unsafe.Pointer) {
	// 0: (, item QTreeWidgetItem *, column int, widget QWidget *), (item, &column, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget13setItemWidgetEP15QTreeWidgetItemiP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &column, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:307
// index:0
// inline
// void removeItemWidget(class QTreeWidgetItem *, int)
func (this *QTreeWidget) RemoveItemWidget(item unsafe.Pointer, column int) {
	// 0: (, item QTreeWidgetItem *, column int), (item, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget16removeItemWidgetEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:309
// index:0
// bool isItemSelected(const class QTreeWidgetItem *)
func (this *QTreeWidget) IsItemSelected(item unsafe.Pointer) {
	// 0: (, item const QTreeWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget14isItemSelectedEPK15QTreeWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:310
// index:0
// void setItemSelected(const class QTreeWidgetItem *, _Bool)
func (this *QTreeWidget) SetItemSelected(item unsafe.Pointer, select_ bool) {
	// 0: (, item const QTreeWidgetItem *, select bool), (item, &select_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget15setItemSelectedEPK15QTreeWidgetItemb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &select_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:311
// index:0
// QList<QTreeWidgetItem *> selectedItems()
func (this *QTreeWidget) SelectedItems() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget13selectedItemsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:312
// index:0
// QList<QTreeWidgetItem *> findItems(const class QString &, Qt::MatchFlags, int)
func (this *QTreeWidget) FindItems(text unsafe.Pointer, flags int, column int) {
	// 0: (, text const QString &, QFlags<Qt::MatchFlag> flags, column int), (text, &flags, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget9findItemsERK7QString6QFlagsIN2Qt9MatchFlagEEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text, &flags, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:315
// index:0
// bool isItemHidden(const class QTreeWidgetItem *)
func (this *QTreeWidget) IsItemHidden(item unsafe.Pointer) {
	// 0: (, item const QTreeWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget12isItemHiddenEPK15QTreeWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:316
// index:0
// void setItemHidden(const class QTreeWidgetItem *, _Bool)
func (this *QTreeWidget) SetItemHidden(item unsafe.Pointer, hide bool) {
	// 0: (, item const QTreeWidgetItem *, hide bool), (item, &hide)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget13setItemHiddenEPK15QTreeWidgetItemb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:318
// index:0
// bool isItemExpanded(const class QTreeWidgetItem *)
func (this *QTreeWidget) IsItemExpanded(item unsafe.Pointer) {
	// 0: (, item const QTreeWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget14isItemExpandedEPK15QTreeWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:319
// index:0
// void setItemExpanded(const class QTreeWidgetItem *, _Bool)
func (this *QTreeWidget) SetItemExpanded(item unsafe.Pointer, expand bool) {
	// 0: (, item const QTreeWidgetItem *, expand bool), (item, &expand)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget15setItemExpandedEPK15QTreeWidgetItemb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &expand)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:321
// index:0
// bool isFirstItemColumnSpanned(const class QTreeWidgetItem *)
func (this *QTreeWidget) IsFirstItemColumnSpanned(item unsafe.Pointer) {
	// 0: (, item const QTreeWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget24isFirstItemColumnSpannedEPK15QTreeWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:322
// index:0
// void setFirstItemColumnSpanned(const class QTreeWidgetItem *, _Bool)
func (this *QTreeWidget) SetFirstItemColumnSpanned(item unsafe.Pointer, span bool) {
	// 0: (, item const QTreeWidgetItem *, span bool), (item, &span)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget25setFirstItemColumnSpannedEPK15QTreeWidgetItemb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &span)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:324
// index:0
// QTreeWidgetItem * itemAbove(const class QTreeWidgetItem *)
func (this *QTreeWidget) ItemAbove(item unsafe.Pointer) {
	// 0: (, item const QTreeWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget9itemAboveEPK15QTreeWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:325
// index:0
// QTreeWidgetItem * itemBelow(const class QTreeWidgetItem *)
func (this *QTreeWidget) ItemBelow(item unsafe.Pointer) {
	// 0: (, item const QTreeWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget9itemBelowEPK15QTreeWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:327
// index:0
// virtual
// void setSelectionModel(class QItemSelectionModel *)
func (this *QTreeWidget) SetSelectionModel(selectionModel unsafe.Pointer) {
	// 0: (, selectionModel QItemSelectionModel *), (selectionModel)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget17setSelectionModelEP19QItemSelectionModel", ffiqt.FFI_TYPE_VOID, this.GetCthis(), selectionModel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:330
// index:0
// void scrollToItem(const class QTreeWidgetItem *, class QAbstractItemView::ScrollHint)
func (this *QTreeWidget) ScrollToItem(item unsafe.Pointer, hint int) {
	// 0: (, item const QTreeWidgetItem *, hint QAbstractItemView::ScrollHint), (item, &hint)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget12scrollToItemEPK15QTreeWidgetItemN17QAbstractItemView10ScrollHintE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:332
// index:0
// void expandItem(const class QTreeWidgetItem *)
func (this *QTreeWidget) ExpandItem(item unsafe.Pointer) {
	// 0: (, item const QTreeWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget10expandItemEPK15QTreeWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:333
// index:0
// void collapseItem(const class QTreeWidgetItem *)
func (this *QTreeWidget) CollapseItem(item unsafe.Pointer) {
	// 0: (, item const QTreeWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget12collapseItemEPK15QTreeWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:334
// index:0
// void clear()
func (this *QTreeWidget) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget5clearEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:337
// index:0
// void itemPressed(class QTreeWidgetItem *, int)
func (this *QTreeWidget) ItemPressed(item unsafe.Pointer, column int) {
	// 0: (, item QTreeWidgetItem *, column int), (item, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget11itemPressedEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:338
// index:0
// void itemClicked(class QTreeWidgetItem *, int)
func (this *QTreeWidget) ItemClicked(item unsafe.Pointer, column int) {
	// 0: (, item QTreeWidgetItem *, column int), (item, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget11itemClickedEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:339
// index:0
// void itemDoubleClicked(class QTreeWidgetItem *, int)
func (this *QTreeWidget) ItemDoubleClicked(item unsafe.Pointer, column int) {
	// 0: (, item QTreeWidgetItem *, column int), (item, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget17itemDoubleClickedEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:340
// index:0
// void itemActivated(class QTreeWidgetItem *, int)
func (this *QTreeWidget) ItemActivated(item unsafe.Pointer, column int) {
	// 0: (, item QTreeWidgetItem *, column int), (item, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget13itemActivatedEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:341
// index:0
// void itemEntered(class QTreeWidgetItem *, int)
func (this *QTreeWidget) ItemEntered(item unsafe.Pointer, column int) {
	// 0: (, item QTreeWidgetItem *, column int), (item, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget11itemEnteredEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:342
// index:0
// void itemChanged(class QTreeWidgetItem *, int)
func (this *QTreeWidget) ItemChanged(item unsafe.Pointer, column int) {
	// 0: (, item QTreeWidgetItem *, column int), (item, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget11itemChangedEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:343
// index:0
// void itemExpanded(class QTreeWidgetItem *)
func (this *QTreeWidget) ItemExpanded(item unsafe.Pointer) {
	// 0: (, item QTreeWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget12itemExpandedEP15QTreeWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:344
// index:0
// void itemCollapsed(class QTreeWidgetItem *)
func (this *QTreeWidget) ItemCollapsed(item unsafe.Pointer) {
	// 0: (, item QTreeWidgetItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget13itemCollapsedEP15QTreeWidgetItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:345
// index:0
// void currentItemChanged(class QTreeWidgetItem *, class QTreeWidgetItem *)
func (this *QTreeWidget) CurrentItemChanged(current unsafe.Pointer, previous unsafe.Pointer) {
	// 0: (, current QTreeWidgetItem *, previous QTreeWidgetItem *), (current, previous)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget18currentItemChangedEP15QTreeWidgetItemS1_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), current, previous)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:346
// index:0
// void itemSelectionChanged()
func (this *QTreeWidget) ItemSelectionChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget20itemSelectionChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:349
// index:0
// virtual
// bool event(class QEvent *)
func (this *QTreeWidget) Event(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:350
// index:0
// virtual
// QStringList mimeTypes()
func (this *QTreeWidget) MimeTypes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget9mimeTypesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:356
// index:0
// virtual
// bool dropMimeData(class QTreeWidgetItem *, int, const class QMimeData *, Qt::DropAction)
func (this *QTreeWidget) DropMimeData(parent unsafe.Pointer, index int, data unsafe.Pointer, action int) {
	// 0: (, parent QTreeWidgetItem *, index int, data const QMimeData *, action Qt::DropAction), (parent, &index, data, &action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget12dropMimeDataEP15QTreeWidgetItemiPK9QMimeDataN2Qt10DropActionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent, &index, data, &action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:358
// index:0
// virtual
// Qt::DropActions supportedDropActions()
func (this *QTreeWidget) SupportedDropActions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget20supportedDropActionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:365
// index:0
// QList<QTreeWidgetItem *> items(const class QMimeData *)
func (this *QTreeWidget) Items(data unsafe.Pointer) {
	// 0: (, data const QMimeData *), (data)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget5itemsEPK9QMimeData", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:367
// index:0
// QModelIndex indexFromItem(const class QTreeWidgetItem *, int)
func (this *QTreeWidget) IndexFromItem(item unsafe.Pointer, column int) {
	// 0: (, item const QTreeWidgetItem *, column int), (item, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget13indexFromItemEPK15QTreeWidgetItemi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:368
// index:1
// QModelIndex indexFromItem(class QTreeWidgetItem *, int)
func (this *QTreeWidget) IndexFromItem_1(item unsafe.Pointer, column int) {
	// 1: (, item QTreeWidgetItem *, column int), (item, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget13indexFromItemEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:369
// index:0
// QTreeWidgetItem * itemFromIndex(const class QModelIndex &)
func (this *QTreeWidget) ItemFromIndex(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget13itemFromIndexERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:373
// index:0
// virtual
// void dropEvent(class QDropEvent *)
func (this *QTreeWidget) DropEvent(event unsafe.Pointer) {
	// 0: (, event QDropEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget9dropEventEP10QDropEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

//  body block end
