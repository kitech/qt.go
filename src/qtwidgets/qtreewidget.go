package qtwidgets

// /usr/include/qt/QtWidgets/qtreewidget.h
// #include <qtreewidget.h>
// #include <QtWidgets>

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
// Public virtual
// const QMetaObject * metaObject()
func (this *QTreeWidget) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:264
// index:0
// Public
// void QTreeWidget(class QWidget *)
func NewQTreeWidget(parent *QWidget /*444 QWidget **/) *QTreeWidget {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidgetC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
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
func (this *QTreeWidget) ColumnCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget11columnCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtreewidget.h:268
// index:0
// Public
// void setColumnCount(int)
func (this *QTreeWidget) SetColumnCount(columns int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget14setColumnCountEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), columns)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:270
// index:0
// Public
// QTreeWidgetItem * invisibleRootItem()
func (this *QTreeWidget) InvisibleRootItem() *QTreeWidgetItem /*444 QTreeWidgetItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget17invisibleRootItemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:271
// index:0
// Public
// QTreeWidgetItem * topLevelItem(int)
func (this *QTreeWidget) TopLevelItem(index int) *QTreeWidgetItem /*444 QTreeWidgetItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget12topLevelItemEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:272
// index:0
// Public
// int topLevelItemCount()
func (this *QTreeWidget) TopLevelItemCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget17topLevelItemCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtreewidget.h:273
// index:0
// Public
// void insertTopLevelItem(int, class QTreeWidgetItem *)
func (this *QTreeWidget) InsertTopLevelItem(index int, item *QTreeWidgetItem /*444 QTreeWidgetItem **/) {
	var convArg1 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget18insertTopLevelItemEiP15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:274
// index:0
// Public
// void addTopLevelItem(class QTreeWidgetItem *)
func (this *QTreeWidget) AddTopLevelItem(item *QTreeWidgetItem /*444 QTreeWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget15addTopLevelItemEP15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:275
// index:0
// Public
// QTreeWidgetItem * takeTopLevelItem(int)
func (this *QTreeWidget) TakeTopLevelItem(index int) *QTreeWidgetItem /*444 QTreeWidgetItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget16takeTopLevelItemEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:276
// index:0
// Public
// int indexOfTopLevelItem(class QTreeWidgetItem *)
func (this *QTreeWidget) IndexOfTopLevelItem(item *QTreeWidgetItem /*444 QTreeWidgetItem **/) int {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget19indexOfTopLevelItemEP15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtreewidget.h:281
// index:0
// Public
// QTreeWidgetItem * headerItem()
func (this *QTreeWidget) HeaderItem() *QTreeWidgetItem /*444 QTreeWidgetItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget10headerItemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:282
// index:0
// Public
// void setHeaderItem(class QTreeWidgetItem *)
func (this *QTreeWidget) SetHeaderItem(item *QTreeWidgetItem /*444 QTreeWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget13setHeaderItemEP15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
func (this *QTreeWidget) CurrentItem() *QTreeWidgetItem /*444 QTreeWidgetItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget11currentItemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:287
// index:0
// Public
// int currentColumn()
func (this *QTreeWidget) CurrentColumn() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget13currentColumnEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtreewidget.h:288
// index:0
// Public
// void setCurrentItem(class QTreeWidgetItem *)
func (this *QTreeWidget) SetCurrentItem(item *QTreeWidgetItem /*444 QTreeWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:289
// index:1
// Public
// void setCurrentItem(class QTreeWidgetItem *, int)
func (this *QTreeWidget) SetCurrentItem_1(item *QTreeWidgetItem /*444 QTreeWidgetItem **/, column int) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:290
// index:2
// Public
// void setCurrentItem(class QTreeWidgetItem *, int, class QItemSelectionModel::SelectionFlags)
func (this *QTreeWidget) SetCurrentItem_2(item *QTreeWidgetItem /*444 QTreeWidgetItem **/, column int, command int) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItemi6QFlagsIN19QItemSelectionModel13SelectionFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column, command)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:292
// index:0
// Public
// QTreeWidgetItem * itemAt(const class QPoint &)
func (this *QTreeWidget) ItemAt(p *qtcore.QPoint) *QTreeWidgetItem /*444 QTreeWidgetItem **/ {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget6itemAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:293
// index:1
// Public inline
// QTreeWidgetItem * itemAt(int, int)
func (this *QTreeWidget) ItemAt_1(x int, y int) *QTreeWidgetItem /*444 QTreeWidgetItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget6itemAtEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:294
// index:0
// Public
// QRect visualItemRect(const class QTreeWidgetItem *)
func (this *QTreeWidget) VisualItemRect(item *QTreeWidgetItem /*444 const QTreeWidgetItem **/) *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget14visualItemRectEPK15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:296
// index:0
// Public
// int sortColumn()
func (this *QTreeWidget) SortColumn() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget10sortColumnEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtreewidget.h:297
// index:0
// Public
// void sortItems(int, Qt::SortOrder)
func (this *QTreeWidget) SortItems(column int, order int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget9sortItemsEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:299
// index:0
// Public
// void editItem(class QTreeWidgetItem *, int)
func (this *QTreeWidget) EditItem(item *QTreeWidgetItem /*444 QTreeWidgetItem **/, column int) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget8editItemEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:300
// index:0
// Public
// void openPersistentEditor(class QTreeWidgetItem *, int)
func (this *QTreeWidget) OpenPersistentEditor(item *QTreeWidgetItem /*444 QTreeWidgetItem **/, column int) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget20openPersistentEditorEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:301
// index:0
// Public
// void closePersistentEditor(class QTreeWidgetItem *, int)
func (this *QTreeWidget) ClosePersistentEditor(item *QTreeWidgetItem /*444 QTreeWidgetItem **/, column int) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget21closePersistentEditorEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:303
// index:0
// Public
// bool isPersistentEditorOpen(class QTreeWidgetItem *, int)
func (this *QTreeWidget) IsPersistentEditorOpen(item *QTreeWidgetItem /*444 QTreeWidgetItem **/, column int) bool {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget22isPersistentEditorOpenEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:305
// index:0
// Public
// QWidget * itemWidget(class QTreeWidgetItem *, int)
func (this *QTreeWidget) ItemWidget(item *QTreeWidgetItem /*444 QTreeWidgetItem **/, column int) *QWidget /*444 QWidget **/ {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget10itemWidgetEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:306
// index:0
// Public
// void setItemWidget(class QTreeWidgetItem *, int, class QWidget *)
func (this *QTreeWidget) SetItemWidget(item *QTreeWidgetItem /*444 QTreeWidgetItem **/, column int, widget *QWidget /*444 QWidget **/) {
	var convArg0 = item.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget13setItemWidgetEP15QTreeWidgetItemiP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:307
// index:0
// Public inline
// void removeItemWidget(class QTreeWidgetItem *, int)
func (this *QTreeWidget) RemoveItemWidget(item *QTreeWidgetItem /*444 QTreeWidgetItem **/, column int) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget16removeItemWidgetEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:309
// index:0
// Public
// bool isItemSelected(const class QTreeWidgetItem *)
func (this *QTreeWidget) IsItemSelected(item *QTreeWidgetItem /*444 const QTreeWidgetItem **/) bool {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget14isItemSelectedEPK15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:310
// index:0
// Public
// void setItemSelected(const class QTreeWidgetItem *, _Bool)
func (this *QTreeWidget) SetItemSelected(item *QTreeWidgetItem /*444 const QTreeWidgetItem **/, select_ bool) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget15setItemSelectedEPK15QTreeWidgetItemb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, select_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:315
// index:0
// Public
// bool isItemHidden(const class QTreeWidgetItem *)
func (this *QTreeWidget) IsItemHidden(item *QTreeWidgetItem /*444 const QTreeWidgetItem **/) bool {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget12isItemHiddenEPK15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:316
// index:0
// Public
// void setItemHidden(const class QTreeWidgetItem *, _Bool)
func (this *QTreeWidget) SetItemHidden(item *QTreeWidgetItem /*444 const QTreeWidgetItem **/, hide bool) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget13setItemHiddenEPK15QTreeWidgetItemb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:318
// index:0
// Public
// bool isItemExpanded(const class QTreeWidgetItem *)
func (this *QTreeWidget) IsItemExpanded(item *QTreeWidgetItem /*444 const QTreeWidgetItem **/) bool {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget14isItemExpandedEPK15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:319
// index:0
// Public
// void setItemExpanded(const class QTreeWidgetItem *, _Bool)
func (this *QTreeWidget) SetItemExpanded(item *QTreeWidgetItem /*444 const QTreeWidgetItem **/, expand bool) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget15setItemExpandedEPK15QTreeWidgetItemb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, expand)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:321
// index:0
// Public
// bool isFirstItemColumnSpanned(const class QTreeWidgetItem *)
func (this *QTreeWidget) IsFirstItemColumnSpanned(item *QTreeWidgetItem /*444 const QTreeWidgetItem **/) bool {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget24isFirstItemColumnSpannedEPK15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:322
// index:0
// Public
// void setFirstItemColumnSpanned(const class QTreeWidgetItem *, _Bool)
func (this *QTreeWidget) SetFirstItemColumnSpanned(item *QTreeWidgetItem /*444 const QTreeWidgetItem **/, span bool) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget25setFirstItemColumnSpannedEPK15QTreeWidgetItemb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, span)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:324
// index:0
// Public
// QTreeWidgetItem * itemAbove(const class QTreeWidgetItem *)
func (this *QTreeWidget) ItemAbove(item *QTreeWidgetItem /*444 const QTreeWidgetItem **/) *QTreeWidgetItem /*444 QTreeWidgetItem **/ {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget9itemAboveEPK15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:325
// index:0
// Public
// QTreeWidgetItem * itemBelow(const class QTreeWidgetItem *)
func (this *QTreeWidget) ItemBelow(item *QTreeWidgetItem /*444 const QTreeWidgetItem **/) *QTreeWidgetItem /*444 QTreeWidgetItem **/ {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget9itemBelowEPK15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:327
// index:0
// Public virtual
// void setSelectionModel(class QItemSelectionModel *)
func (this *QTreeWidget) SetSelectionModel(selectionModel *qtcore.QItemSelectionModel /*444 QItemSelectionModel **/) {
	var convArg0 = selectionModel.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget17setSelectionModelEP19QItemSelectionModel", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:330
// index:0
// Public
// void scrollToItem(const class QTreeWidgetItem *, class QAbstractItemView::ScrollHint)
func (this *QTreeWidget) ScrollToItem(item *QTreeWidgetItem /*444 const QTreeWidgetItem **/, hint int) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget12scrollToItemEPK15QTreeWidgetItemN17QAbstractItemView10ScrollHintE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:332
// index:0
// Public
// void expandItem(const class QTreeWidgetItem *)
func (this *QTreeWidget) ExpandItem(item *QTreeWidgetItem /*444 const QTreeWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget10expandItemEPK15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:333
// index:0
// Public
// void collapseItem(const class QTreeWidgetItem *)
func (this *QTreeWidget) CollapseItem(item *QTreeWidgetItem /*444 const QTreeWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget12collapseItemEPK15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
func (this *QTreeWidget) ItemPressed(item *QTreeWidgetItem /*444 QTreeWidgetItem **/, column int) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget11itemPressedEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:338
// index:0
// Public
// void itemClicked(class QTreeWidgetItem *, int)
func (this *QTreeWidget) ItemClicked(item *QTreeWidgetItem /*444 QTreeWidgetItem **/, column int) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget11itemClickedEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:339
// index:0
// Public
// void itemDoubleClicked(class QTreeWidgetItem *, int)
func (this *QTreeWidget) ItemDoubleClicked(item *QTreeWidgetItem /*444 QTreeWidgetItem **/, column int) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget17itemDoubleClickedEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:340
// index:0
// Public
// void itemActivated(class QTreeWidgetItem *, int)
func (this *QTreeWidget) ItemActivated(item *QTreeWidgetItem /*444 QTreeWidgetItem **/, column int) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget13itemActivatedEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:341
// index:0
// Public
// void itemEntered(class QTreeWidgetItem *, int)
func (this *QTreeWidget) ItemEntered(item *QTreeWidgetItem /*444 QTreeWidgetItem **/, column int) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget11itemEnteredEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:342
// index:0
// Public
// void itemChanged(class QTreeWidgetItem *, int)
func (this *QTreeWidget) ItemChanged(item *QTreeWidgetItem /*444 QTreeWidgetItem **/, column int) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget11itemChangedEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:343
// index:0
// Public
// void itemExpanded(class QTreeWidgetItem *)
func (this *QTreeWidget) ItemExpanded(item *QTreeWidgetItem /*444 QTreeWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget12itemExpandedEP15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:344
// index:0
// Public
// void itemCollapsed(class QTreeWidgetItem *)
func (this *QTreeWidget) ItemCollapsed(item *QTreeWidgetItem /*444 QTreeWidgetItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget13itemCollapsedEP15QTreeWidgetItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:345
// index:0
// Public
// void currentItemChanged(class QTreeWidgetItem *, class QTreeWidgetItem *)
func (this *QTreeWidget) CurrentItemChanged(current *QTreeWidgetItem /*444 QTreeWidgetItem **/, previous *QTreeWidgetItem /*444 QTreeWidgetItem **/) {
	var convArg0 = current.GetCthis()
	var convArg1 = previous.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget18currentItemChangedEP15QTreeWidgetItemS1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
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
func (this *QTreeWidget) Event(e *qtcore.QEvent /*444 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:356
// index:0
// Protected virtual
// bool dropMimeData(class QTreeWidgetItem *, int, const class QMimeData *, Qt::DropAction)
func (this *QTreeWidget) DropMimeData(parent *QTreeWidgetItem /*444 QTreeWidgetItem **/, index int, data *qtcore.QMimeData /*444 const QMimeData **/, action int) bool {
	var convArg0 = parent.GetCthis()
	var convArg2 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget12dropMimeDataEP15QTreeWidgetItemiPK9QMimeDataN2Qt10DropActionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, index, convArg2, action)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:358
// index:0
// Protected virtual
// Qt::DropActions supportedDropActions()
func (this *QTreeWidget) SupportedDropActions() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget20supportedDropActionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:367
// index:0
// Protected
// QModelIndex indexFromItem(const class QTreeWidgetItem *, int)
func (this *QTreeWidget) IndexFromItem(item *QTreeWidgetItem /*444 const QTreeWidgetItem **/, column int) *qtcore.QModelIndex /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget13indexFromItemEPK15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, column)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:368
// index:1
// Protected
// QModelIndex indexFromItem(class QTreeWidgetItem *, int)
func (this *QTreeWidget) IndexFromItem_1(item *QTreeWidgetItem /*444 QTreeWidgetItem **/, column int) *qtcore.QModelIndex /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget13indexFromItemEP15QTreeWidgetItemi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, column)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:369
// index:0
// Protected
// QTreeWidgetItem * itemFromIndex(const class QModelIndex &)
func (this *QTreeWidget) ItemFromIndex(index *qtcore.QModelIndex) *QTreeWidgetItem /*444 QTreeWidgetItem **/ {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTreeWidget13itemFromIndexERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:373
// index:0
// Protected virtual
// void dropEvent(class QDropEvent *)
func (this *QTreeWidget) DropEvent(event *qtgui.QDropEvent /*444 QDropEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTreeWidget9dropEventEP10QDropEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
