//  header block begin
// /usr/include/qt/QtWidgets/qtreewidget.h
// #include <qtreewidget.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 67
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
func NewQTreeWidgetFromPointer(cthis unsafe.Pointer) *QTreeWidget {
	bcthis0 := NewQTreeViewFromPointer(cthis)
	return &QTreeWidget{bcthis0}
}

// /usr/include/qt/QtWidgets/qtreewidget.h:257
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QTreeWidget) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:264
// index:0
// Public
// void QTreeWidget(class QWidget *)
func NewQTreeWidget(parent unsafe.Pointer) *QTreeWidget {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidgetC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQTreeWidgetFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:265
// index:0
// Public virtual
// void ~QTreeWidget()
func DeleteQTreeWidget(*QTreeWidget) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidgetD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:267
// index:0
// Public
// int columnCount()
func (this *QTreeWidget) ColumnCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget11columnCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:268
// index:0
// Public
// void setColumnCount(int)
func (this *QTreeWidget) SetColumnCount(columns int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget14setColumnCountEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &columns)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:270
// index:0
// Public
// QTreeWidgetItem * invisibleRootItem()
func (this *QTreeWidget) InvisibleRootItem() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget17invisibleRootItemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:271
// index:0
// Public
// QTreeWidgetItem * topLevelItem(int)
func (this *QTreeWidget) TopLevelItem(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget12topLevelItemEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:272
// index:0
// Public
// int topLevelItemCount()
func (this *QTreeWidget) TopLevelItemCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget17topLevelItemCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:273
// index:0
// Public
// void insertTopLevelItem(int, class QTreeWidgetItem *)
func (this *QTreeWidget) InsertTopLevelItem(index int, item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget18insertTopLevelItemEiP15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:274
// index:0
// Public
// void addTopLevelItem(class QTreeWidgetItem *)
func (this *QTreeWidget) AddTopLevelItem(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget15addTopLevelItemEP15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:275
// index:0
// Public
// QTreeWidgetItem * takeTopLevelItem(int)
func (this *QTreeWidget) TakeTopLevelItem(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget16takeTopLevelItemEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:276
// index:0
// Public
// int indexOfTopLevelItem(class QTreeWidgetItem *)
func (this *QTreeWidget) IndexOfTopLevelItem(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget19indexOfTopLevelItemEP15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:281
// index:0
// Public
// QTreeWidgetItem * headerItem()
func (this *QTreeWidget) HeaderItem() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget10headerItemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:282
// index:0
// Public
// void setHeaderItem(class QTreeWidgetItem *)
func (this *QTreeWidget) SetHeaderItem(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget13setHeaderItemEP15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:283
// index:0
// Public
// void setHeaderLabels(const class QStringList &)
func (this *QTreeWidget) SetHeaderLabels(labels *qtcore.QStringList) {
	var convArg0 = labels.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget15setHeaderLabelsERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:284
// index:0
// Public inline
// void setHeaderLabel(const class QString &)
func (this *QTreeWidget) SetHeaderLabel(label *qtcore.QString) {
	var convArg0 = label.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget14setHeaderLabelERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:286
// index:0
// Public
// QTreeWidgetItem * currentItem()
func (this *QTreeWidget) CurrentItem() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget11currentItemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:287
// index:0
// Public
// int currentColumn()
func (this *QTreeWidget) CurrentColumn() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget13currentColumnEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:288
// index:0
// Public
// void setCurrentItem(class QTreeWidgetItem *)
func (this *QTreeWidget) SetCurrentItem(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:289
// index:1
// Public
// void setCurrentItem(class QTreeWidgetItem *, int)
func (this *QTreeWidget) SetCurrentItem_1(item unsafe.Pointer, column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:292
// index:0
// Public
// QTreeWidgetItem * itemAt(const class QPoint &)
func (this *QTreeWidget) ItemAt(p *qtcore.QPoint) interface{} {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget6itemAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:293
// index:1
// Public inline
// QTreeWidgetItem * itemAt(int, int)
func (this *QTreeWidget) ItemAt_1(x int, y int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget6itemAtEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:294
// index:0
// Public
// QRect visualItemRect(const class QTreeWidgetItem *)
func (this *QTreeWidget) VisualItemRect(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget14visualItemRectEPK15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:296
// index:0
// Public
// int sortColumn()
func (this *QTreeWidget) SortColumn() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget10sortColumnEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:297
// index:0
// Public
// void sortItems(int, Qt::SortOrder)
func (this *QTreeWidget) SortItems(column int, order int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget9sortItemsEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:299
// index:0
// Public
// void editItem(class QTreeWidgetItem *, int)
func (this *QTreeWidget) EditItem(item unsafe.Pointer, column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget8editItemEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:300
// index:0
// Public
// void openPersistentEditor(class QTreeWidgetItem *, int)
func (this *QTreeWidget) OpenPersistentEditor(item unsafe.Pointer, column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget20openPersistentEditorEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:301
// index:0
// Public
// void closePersistentEditor(class QTreeWidgetItem *, int)
func (this *QTreeWidget) ClosePersistentEditor(item unsafe.Pointer, column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget21closePersistentEditorEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:303
// index:0
// Public
// bool isPersistentEditorOpen(class QTreeWidgetItem *, int)
func (this *QTreeWidget) IsPersistentEditorOpen(item unsafe.Pointer, column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget22isPersistentEditorOpenEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:305
// index:0
// Public
// QWidget * itemWidget(class QTreeWidgetItem *, int)
func (this *QTreeWidget) ItemWidget(item unsafe.Pointer, column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget10itemWidgetEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:306
// index:0
// Public
// void setItemWidget(class QTreeWidgetItem *, int, class QWidget *)
func (this *QTreeWidget) SetItemWidget(item unsafe.Pointer, column int, widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget13setItemWidgetEP15QTreeWidgetItemiP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &column, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:307
// index:0
// Public inline
// void removeItemWidget(class QTreeWidgetItem *, int)
func (this *QTreeWidget) RemoveItemWidget(item unsafe.Pointer, column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget16removeItemWidgetEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:309
// index:0
// Public
// bool isItemSelected(const class QTreeWidgetItem *)
func (this *QTreeWidget) IsItemSelected(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget14isItemSelectedEPK15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:310
// index:0
// Public
// void setItemSelected(const class QTreeWidgetItem *, _Bool)
func (this *QTreeWidget) SetItemSelected(item unsafe.Pointer, select_ bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget15setItemSelectedEPK15QTreeWidgetItemb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &select_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:311
// index:0
// Public
// QList<QTreeWidgetItem *> selectedItems()
func (this *QTreeWidget) SelectedItems() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget13selectedItemsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:315
// index:0
// Public
// bool isItemHidden(const class QTreeWidgetItem *)
func (this *QTreeWidget) IsItemHidden(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget12isItemHiddenEPK15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:316
// index:0
// Public
// void setItemHidden(const class QTreeWidgetItem *, _Bool)
func (this *QTreeWidget) SetItemHidden(item unsafe.Pointer, hide bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget13setItemHiddenEPK15QTreeWidgetItemb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:318
// index:0
// Public
// bool isItemExpanded(const class QTreeWidgetItem *)
func (this *QTreeWidget) IsItemExpanded(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget14isItemExpandedEPK15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:319
// index:0
// Public
// void setItemExpanded(const class QTreeWidgetItem *, _Bool)
func (this *QTreeWidget) SetItemExpanded(item unsafe.Pointer, expand bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget15setItemExpandedEPK15QTreeWidgetItemb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &expand)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:321
// index:0
// Public
// bool isFirstItemColumnSpanned(const class QTreeWidgetItem *)
func (this *QTreeWidget) IsFirstItemColumnSpanned(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget24isFirstItemColumnSpannedEPK15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:322
// index:0
// Public
// void setFirstItemColumnSpanned(const class QTreeWidgetItem *, _Bool)
func (this *QTreeWidget) SetFirstItemColumnSpanned(item unsafe.Pointer, span bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget25setFirstItemColumnSpannedEPK15QTreeWidgetItemb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &span)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:324
// index:0
// Public
// QTreeWidgetItem * itemAbove(const class QTreeWidgetItem *)
func (this *QTreeWidget) ItemAbove(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget9itemAboveEPK15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:325
// index:0
// Public
// QTreeWidgetItem * itemBelow(const class QTreeWidgetItem *)
func (this *QTreeWidget) ItemBelow(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget9itemBelowEPK15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:327
// index:0
// Public virtual
// void setSelectionModel(class QItemSelectionModel *)
func (this *QTreeWidget) SetSelectionModel(selectionModel unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget17setSelectionModelEP19QItemSelectionModel", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), selectionModel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:330
// index:0
// Public
// void scrollToItem(const class QTreeWidgetItem *, class QAbstractItemView::ScrollHint)
func (this *QTreeWidget) ScrollToItem(item unsafe.Pointer, hint int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget12scrollToItemEPK15QTreeWidgetItemN17QAbstractItemView10ScrollHintE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:332
// index:0
// Public
// void expandItem(const class QTreeWidgetItem *)
func (this *QTreeWidget) ExpandItem(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget10expandItemEPK15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:333
// index:0
// Public
// void collapseItem(const class QTreeWidgetItem *)
func (this *QTreeWidget) CollapseItem(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget12collapseItemEPK15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:334
// index:0
// Public
// void clear()
func (this *QTreeWidget) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:337
// index:0
// Public
// void itemPressed(class QTreeWidgetItem *, int)
func (this *QTreeWidget) ItemPressed(item unsafe.Pointer, column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget11itemPressedEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:338
// index:0
// Public
// void itemClicked(class QTreeWidgetItem *, int)
func (this *QTreeWidget) ItemClicked(item unsafe.Pointer, column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget11itemClickedEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:339
// index:0
// Public
// void itemDoubleClicked(class QTreeWidgetItem *, int)
func (this *QTreeWidget) ItemDoubleClicked(item unsafe.Pointer, column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget17itemDoubleClickedEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:340
// index:0
// Public
// void itemActivated(class QTreeWidgetItem *, int)
func (this *QTreeWidget) ItemActivated(item unsafe.Pointer, column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget13itemActivatedEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:341
// index:0
// Public
// void itemEntered(class QTreeWidgetItem *, int)
func (this *QTreeWidget) ItemEntered(item unsafe.Pointer, column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget11itemEnteredEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:342
// index:0
// Public
// void itemChanged(class QTreeWidgetItem *, int)
func (this *QTreeWidget) ItemChanged(item unsafe.Pointer, column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget11itemChangedEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:343
// index:0
// Public
// void itemExpanded(class QTreeWidgetItem *)
func (this *QTreeWidget) ItemExpanded(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget12itemExpandedEP15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:344
// index:0
// Public
// void itemCollapsed(class QTreeWidgetItem *)
func (this *QTreeWidget) ItemCollapsed(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget13itemCollapsedEP15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:345
// index:0
// Public
// void currentItemChanged(class QTreeWidgetItem *, class QTreeWidgetItem *)
func (this *QTreeWidget) CurrentItemChanged(current unsafe.Pointer, previous unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget18currentItemChangedEP15QTreeWidgetItemS1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), current, previous)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:346
// index:0
// Public
// void itemSelectionChanged()
func (this *QTreeWidget) ItemSelectionChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget20itemSelectionChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:349
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QTreeWidget) Event(e unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:350
// index:0
// Protected virtual
// QStringList mimeTypes()
func (this *QTreeWidget) MimeTypes() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget9mimeTypesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:356
// index:0
// Protected virtual
// bool dropMimeData(class QTreeWidgetItem *, int, const class QMimeData *, Qt::DropAction)
func (this *QTreeWidget) DropMimeData(parent unsafe.Pointer, index int, data unsafe.Pointer, action int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget12dropMimeDataEP15QTreeWidgetItemiPK9QMimeDataN2Qt10DropActionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), parent, &index, data, &action)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:358
// index:0
// Protected virtual
// Qt::DropActions supportedDropActions()
func (this *QTreeWidget) SupportedDropActions() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget20supportedDropActionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:365
// index:0
// Protected
// QList<QTreeWidgetItem *> items(const class QMimeData *)
func (this *QTreeWidget) Items(data unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget5itemsEPK9QMimeData", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), data)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:367
// index:0
// Protected
// QModelIndex indexFromItem(const class QTreeWidgetItem *, int)
func (this *QTreeWidget) IndexFromItem(item unsafe.Pointer, column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget13indexFromItemEPK15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:368
// index:1
// Protected
// QModelIndex indexFromItem(class QTreeWidgetItem *, int)
func (this *QTreeWidget) IndexFromItem_1(item unsafe.Pointer, column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget13indexFromItemEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:369
// index:0
// Protected
// QTreeWidgetItem * itemFromIndex(const class QModelIndex &)
func (this *QTreeWidget) ItemFromIndex(index *qtcore.QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget13itemFromIndexERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:373
// index:0
// Protected virtual
// void dropEvent(class QDropEvent *)
func (this *QTreeWidget) DropEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget9dropEventEP10QDropEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

//  body block end
